package main

import (
	"errors"
	"flag"
	"fmt"
	api "github.com/whosonfirst/go-brooklynintegers-api"
	log "github.com/whosonfirst/go-whosonfirst-log"
	pool "github.com/whosonfirst/go-whosonfirst-pool"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Proxy struct {
	logger  *log.WOFLogger
	client  *api.APIClient
	pool    *pool.LIFOPool
	minpool int64
	refill  chan bool
}

func NewProxy(min_pool int64, logger *log.WOFLogger) *Proxy {

	api_client := api.NewAPIClient()
	pool := pool.NewLIFOPool()

	// See notes in RefillPool() for details

	size := 10
	refill := make(chan bool, 10)

	for i := 0; i < size; i++ {
		refill <- true
	}

	// Possibly also keep global stats on number of fetches
	// and cache hits/misses/etc

	proxy := Proxy{
		logger:  logger,
		client:  api_client,
		pool:    pool,
		minpool: min_pool,
		refill:  refill,
	}

	return &proxy
}

func (p *Proxy) Init() {

	go p.RefillPool()

	go p.Status()
	go p.Monitor()
}

func (p *Proxy) Status() {

	for {
		select {
		case <-time.After(5 * time.Second):
			p.logger.Status("pool length: %d", p.pool.Length())
		}
	}
}

func (p *Proxy) Monitor() {

	for {
		select {
		case <-time.After(10 * time.Second):
			if p.pool.Length() < p.minpool {
				go p.RefillPool()
			}
		}

	}
}

func (p *Proxy) RefillPool() {

	// Remember there is a fixed size work queue of allowable times to try
	// and refill the pool simultaneously. First, we block until a slot opens
	// up.

	<-p.refill

	t1 := time.Now()

	// Figure out how many integers we need to get *at this moment* which when
	// the service is under heavy load is a misleading number at best. It might
	// be worth adjusting this by a factor of (n) depending on the current load.
	// But that also means tracking what we think the current load means so we
	// aren't going to do that now...

	todo := p.minpool - p.pool.Length()
	workers := int(p.minpool / 2)

	if workers == 0 {
		workers = 1
	}

	// Now we're going to set up two simultaneous queues. One (the work group) is
	// just there to keep track of all the requests for new integers we need to
	// make. The second (the throttle) is there to make sure we don't exhaust all
	// the filehandles or network connections.

	th := make(chan bool, workers)

	for i := 0; i < workers; i++ {
		th <- true
	}

	wg := new(sync.WaitGroup)

	p.logger.Debug("refill poll w/ %d integers and %d workers", todo, workers)

	success := 0
	failed := 0

	for j := 0; int64(j) < todo; j++ {

		// Wait for the throttle to open a slot. Also record whether
		// the operation was successful.

		rsp := <-th

		if rsp == true {
			success += 1
		} else {
			failed += 1
		}

		// First check that we still actually need to keep fetching integers

		if p.pool.Length() >= p.minpool {
			p.logger.Debug("pool is full (%d) stopping after %d iterations", p.pool.Length(), j)
			break
		}

		// Standard work group stuff

		wg.Add(1)

		// Sudo make me a sandwitch. Note the part where we ping the throttle with
		// the return value at the end both to signal an available slot and to record
		// whether the integer harvesting was successful.

		go func(pr *Proxy) {
			defer wg.Done()
			th <- pr.AddToPool()
		}(p)
	}

	// More standard work group stuff

	wg.Wait()

	// Again note the way we are freeing a spot in the refill queue

	p.refill <- true

	t2 := time.Since(t1)
	p.logger.Info("time to refill the pool with %d integers (success: %d failed: %d): %v (pool length is now %d)", todo, success, failed, t2, p.pool.Length())

}

func (p *Proxy) AddToPool() bool {

	i, err := p.GetInteger()

	if err != nil {
		return false
	}

	pi := pool.PoolInt{Int: i}

	p.pool.Push(pi)
	return true
}

func (p *Proxy) GetInteger() (int64, error) {

	i, err := p.client.CreateInteger()

	if err != nil {
		p.logger.Error("failed to create new integer, because %v", err)
		return 0, err
	}

	p.logger.Debug("got new integer %d", i)
	return i, nil
}

func (p *Proxy) Integer() (int64, error) {

	if p.pool.Length() == 0 {

		go p.RefillPool()

		p.logger.Warning("pool length is 0 so fetching integer from source")
		return p.GetInteger()
	}

	v, ok := p.pool.Pop()

	if !ok {
		p.logger.Error("failed to pop integer!")
		return 0, errors.New("Failed to pop")
	}

	i := v.IntValue()

	p.logger.Debug("return cached integer %d", i)

	return i, nil
}

func main() {

	var port = flag.Int("port", 8080, "Port to listen")
	var min = flag.Int("min", 5, "The minimum number of Brooklyn Integers to keep on hand at all times")
	var loglevel = flag.String("loglevel", "info", "Log level")
	var cors = flag.Bool("cors", false, "Enable CORS headers")

	flag.Parse()

	writer := io.MultiWriter(os.Stdout)

	logger := log.NewWOFLogger("[big-integer] ")
	logger.AddLogger(writer, *loglevel)

	proxy := NewProxy(int64(*min), logger)
	proxy.Init()

	handler := func(rsp http.ResponseWriter, r *http.Request) {

		i, err := proxy.Integer()

		if err != nil {
			msg := fmt.Sprintf("Failed to retrieve integer because %v", err)
			http.Error(rsp, msg, http.StatusBadRequest)
		}

		if *cors {
			rsp.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

		io.WriteString(rsp, strconv.Itoa(int(i)))
	}

	http.HandleFunc("/", handler)

	str_port := ":" + strconv.Itoa(*port)
	err := http.ListenAndServe(str_port, nil)

	if err != nil {
		logger.Fatal("Failed to start server, because %v\n", err)
	}

}
