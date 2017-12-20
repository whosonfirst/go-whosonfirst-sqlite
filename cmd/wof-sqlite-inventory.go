package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Report struct {
	Path           string
	Include        bool
	Count          int
	Size           int64
	SizeCompressed int64
	Sha1Sum        string
	LastUpdate     time.Time
	LastModified   time.Time
}

func NewReport(path string) Report {

	now := time.Now()

	r := Report{
		Path:           path,
		Include:        false,
		Count:          0,
		Size:           0,
		SizeCompressed: 0,
		Sha1Sum:        "",
		LastModified:   now,
		LastUpdate:     now,
	}

	return r
}

func Compress(r Report, report_ch chan Report, done_ch chan bool, err_ch chan error) {

	defer func() {
		done_ch <- true
	}()

	abs_source, err := filepath.Abs(r.Path)

	if err != nil {
		err_ch <- err
		return
	}

	abs_chroot := filepath.Dir(abs_source)

	fname := filepath.Base(abs_source)
	fname = fmt.Sprintf("%s.bz2", fname)

	dest := filepath.Join(abs_chroot, fname)

	info, err := os.Stat(dest)

	if os.IsNotExist(err) {

		bz := "bzip2"

		args := []string{
			"-f",
			"-k", // Keep (don't delete) input files during compression or decompression.
			abs_source,
		}

		log.Println(bz, strings.Join(args, " "))

		cmd := exec.Command(bz, args...)
		err = cmd.Run()

		if err != nil {
			err_ch <- err
			return
		}

		info, err = os.Stat(dest)

		if err != nil {
			err_ch <- err
			return
		}

	} else {

		if err != nil {
			err_ch <- err
			return
		}

	}

	r.SizeCompressed = info.Size()
	report_ch <- r
	return
}

func Hash(r Report, report_ch chan Report, done_ch chan bool, err_ch chan error) {

	defer func() {
		done_ch <- true
	}()

	path := fmt.Sprintf("%s.bz2", r.Path)

	fh, err := os.Open(path)

	if err != nil {
		err_ch <- err
		return
	}

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		err_ch <- err
		return
	}

	defer fh.Close()

	hash := sha1.Sum(body)
	enc := hex.EncodeToString(hash[:])

	r.Sha1Sum = enc

	report_ch <- r
	return
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

	r := NewReport(abs_path)

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

	reports := make([]Report, 0)

	for remaining > 0 {

		select {
		case <-done_ch:
			remaining -= 1
		case err := <-error_ch:
			log.Println("ERROR", err)
		case r := <-report_ch:

			if r.Include {
				reports = append(reports, r)
			} else {
				log.Println("REMOVE", r.Path)
				os.Remove(r.Path)
			}
		}
	}

	/*
		count_throttle := 8
		throttle_ch := make(chan bool, count_throttle)

		for i :=0; i < count_throttle; i ++ {
			throttle_ch <- true
		}
	*/

	for _, r := range reports {

		log.Println("WAITING FOR THROTTLE")

		// throttle_ch

		go func(r Report) {
			Compress(r, report_ch, done_ch, error_ch)
		}(r)
	}

	remaining = len(reports)

	for remaining > 0 {

		select {
		case <-done_ch:
			log.Println("DONE", remaining)
			remaining -= 1
		case err := <-error_ch:
			log.Println("ERROR", err)
		case r := <-report_ch:
			log.Println("REPORT", r)
		}
	}

	log.Println("DONE")
}
