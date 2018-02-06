package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-brands/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-crawl"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {

	var procs = flag.Int("processes", runtime.NumCPU()*2, "The number of concurrent processes to use")
	var root = flag.String("root", "", "...")

	flag.Parse()

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

		br, err := whosonfirst.LoadWOFBrandFromFile(path)

		if err != nil {
			log.Printf("failed to open %s, because %s\n", path, err)
			return err
		}

		is_current, err := br.IsCurrent()

		if err != nil {
			log.Printf("failed to is_current for %s, because %s\n", path, err)
			return err
		}

		log.Println(br.Id(), br.Name(), is_current, br.LastModified())
		return nil
	}

	c := crawl.NewCrawler(data)
	c.Crawl(cb)

	os.Exit(0)
}
