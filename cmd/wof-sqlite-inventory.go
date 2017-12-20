package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Report struct {
	Path         string
	Include      bool
	Count        int
	Size         int64
	LastUpdate   time.Time
	LastModified time.Time
}

func Inventory(path string, report_ch chan Report, done_ch chan bool, error_ch chan error) {

	defer func() {
		done_ch <- true
	}()

	abs_path, err := filepath.Abs(path)

	if err != nil {
		error_ch <- err
		return
	}

	db, err := database.NewDB(abs_path)

	if err != nil {
		error_ch <- err
		return
	}

	conn, err := db.Conn()

	if err != nil {
		error_ch <- err
		return
	}

	row := conn.QueryRow("SELECT COUNT(id) FROM spr")

	var count int
	err = row.Scan(&count)

	if err != nil {
		error_ch <- err
		return
	}

	now := time.Now()

	r := Report{
		Path:         abs_path,
		Include:      false,
		Count:        0,
		Size:         0,
		LastModified: now,
		LastUpdate:   now,
	}

	if count == 0 {
		report_ch <- r
		return
	}

	r.Include = true
	r.Count = count

	row = conn.QueryRow("SELECT MAX(lastmodified) FROM spr")

	var lastmod int64
	err = row.Scan(&lastmod)

	if err != nil {
		error_ch <- err
		return
	}

	info, err := os.Stat(path)

	if err != nil {
		error_ch <- err
		return
	}

	r.Size = info.Size()
	r.LastModified = info.ModTime()
	r.LastUpdate = time.Unix(lastmod, 0)

	report_ch <- r
}

func main() {

	flag.Parse()

	databases := flag.Args()
	count := len(databases)

	report_ch := make(chan Report)
	done_ch := make(chan bool)
	error_ch := make(chan error)

	for _, path := range databases {
		go Inventory(path, report_ch, done_ch, error_ch)
	}

	remaining := count

	// reports := make(map[string]Report)

	for remaining > 0 {

		select {
		case <-done_ch:
			remaining -= 1
		case err := <-error_ch:
			log.Println("ERROR", err)
		case r := <-report_ch:

			if r.Include {

				// do stuff here...
			} else {
				log.Println("REMOVE", r.Path)
				os.Remove(r.Path)
			}
		}
	}

	log.Println("DONE")
}
