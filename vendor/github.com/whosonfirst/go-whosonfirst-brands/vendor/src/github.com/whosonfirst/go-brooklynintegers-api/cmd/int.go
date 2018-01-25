package main

import (
	"bufio"
	"flag"
	"github.com/whosonfirst/go-brooklynintegers-api"
	"github.com/whosonfirst/go-writer-tts"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {

	var count = flag.Int("count", 1, "The number of Brooklyn Integers to mint.")
	var tts_speak = flag.Bool("tts", false, "Output integers to a text-to-speak engine.")
	var tts_engine = flag.String("tts-engine", "", "A valid go-writer-tts text-to-speak engine. Valid options are: osx.")

	flag.Parse()

	writers := []io.Writer{
		os.Stdout,
	}

	if *tts_speak {

		speaker, err := tts.NewSpeakerForEngine(*tts_engine)

		if err != nil {
			log.Fatal(err)
		}

		writers = append(writers, speaker)
	}

	multi := io.MultiWriter(writers...)
	writer := bufio.NewWriter(multi)

	client := api.NewAPIClient()

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	t := 10
	throttle := make(chan bool, t)

	for i := 0; i < t; i++ {

		throttle <- true
	}

	for i := 0; i < *count; i++ {

		wg.Add(1)

		go func(throttle chan bool) {

			<-throttle

			defer func() {
				throttle <- true
				wg.Done()
			}()

			i, err := client.CreateInteger()

			if err != nil {
				log.Fatal(err)
			}

			mu.Lock()

			str_i := strconv.Itoa(int(i))
			writer.WriteString(str_i + "\n")
			writer.Flush()

			mu.Unlock()

		}(throttle)

	}

	wg.Wait()
	os.Exit(0)
}
