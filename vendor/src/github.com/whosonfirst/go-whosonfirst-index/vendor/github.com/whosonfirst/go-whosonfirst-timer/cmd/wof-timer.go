package main

import (
	"github.com/whosonfirst/go-whosonfirst-timer"
	"log"
	"time"
)

func main() {

	tm, err := timer.NewDefaultTimer()

	if err != nil {
		log.Fatal(err)
	}

	defer tm.Stop()

	time.Sleep(1000)
}
