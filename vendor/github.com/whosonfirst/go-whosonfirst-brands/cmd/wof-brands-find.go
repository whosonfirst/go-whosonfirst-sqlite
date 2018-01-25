package main

import (
	"flag"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-crawl"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {

	var procs = flag.Int("processes", runtime.NumCPU()*2, "The number of concurrent processes to use")
	var root = flag.String("root", "", "...")

	flag.Parse()

	names := flag.Args()

	runtime.GOMAXPROCS(*procs)

	data := filepath.Join(*root, "data")

	_, err := os.Stat(data)

	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	cb := func(path string, info os.FileInfo) error {

		if info.IsDir() {
			return nil
		}

		fh, err := os.Open(path)

		if err != nil {
			log.Printf("failed to open %s, because %s\n", path, err)
			return err
		}

		defer fh.Close()

		body, err := ioutil.ReadAll(fh)

		if err != nil {
			log.Printf("failed to read %s, because %s\n", path, err)
			return err
		}

		rsp := gjson.GetBytes(body, "wof:name")

		if !rsp.Exists() {
			return nil
		}

		for _, n := range names {

			if n == rsp.String() {
				fmt.Printf("%s %s\n", path, n)
			}
		}

		return nil
	}

	c := crawl.NewCrawler(data)
	c.Crawl(cb)

	os.Exit(0)
}
