package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/whosonfirst/go-whosonfirst-sqlite/assets/html"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type HTMLVars struct {
	Reports      []Report
	LastModified string
}

type Report struct {
	path             string
	include          bool
	Name             string    `json:"name"`
	NameCompressed   string    `json:"name_compressed"`
	Count            int       `json:"count"`
	Size             int64     `json:"size"`
	SizeCompressed   int64     `json:"size_compressed"`
	Sha256Compressed string    `json:"sha256_compressed"`
	LastUpdate       time.Time `json:"last_update"`
	LastModified     time.Time `json:"lastmodified"`
}

func NewReport(path string) Report {

	name := filepath.Base(path)
	name_compressed := fmt.Sprintf("%s.bz2", name)

	now := time.Now()

	r := Report{
		path:             path,
		include:          false,
		Name:             name,
		NameCompressed:   name_compressed,
		Count:            0,
		Size:             0,
		SizeCompressed:   0,
		Sha256Compressed: "",
		LastModified:     now,
		LastUpdate:       now,
	}

	return r
}

func (r Report) CountString() string {
	return humanize.Comma(int64(r.Count))
}

func (r Report) SizeString() string {
	return humanize.Bytes(uint64(r.Size))
}

func (r Report) SizeCompressedString() string {
	return humanize.Bytes(uint64(r.SizeCompressed))
}

func (r Report) LastModifiedString() string {
	return r.LastModified.Format(time.RFC3339)
}

func (r Report) LastUpdateString() string {
	return r.LastUpdate.Format(time.RFC3339)
}

func Compress(r Report, report_ch chan Report, done_ch chan bool, err_ch chan error) {

	defer func() {
		done_ch <- true
	}()

	abs_source, err := filepath.Abs(r.path)

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

	h, err := HashFile(dest)

	if err != nil {
		err_ch <- err
		return
	}

	r.Sha256Compressed = h

	report_ch <- r
	return
}

func HashFile(path string) (string, error) {

	fh, err := os.Open(path)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return "", err
	}

	defer fh.Close()

	hash := sha256.Sum256(body)
	enc := hex.EncodeToString(hash[:])

	return enc, nil
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

	row := conn.QueryRow("SELECT COUNT(id) FROM geojson")

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

	r.include = true
	r.Count = count

	row = conn.QueryRow("SELECT MAX(lastmodified) FROM geojson")

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

	var outdir = flag.String("outdir", "", "...")

	flag.Parse()

	// general setup - or things that we want to fail immediately
	// if they're going to fail at all

	if *outdir == "" {

		cwd, err := os.Getwd()

		if err != nil {
			log.Fatal()
		}

		*outdir = cwd
	} else {

		info, err := os.Stat(*outdir)

		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() {
			log.Fatal("outdir is not a directory")
		}
	}

	inventory_html := filepath.Join(*outdir, "inventory.html")
	inventory_json := filepath.Join(*outdir, "inventory.json")

	fh_html, err := os.Create(inventory_html)

	if err != nil {
		log.Fatal(err)
	}

	defer fh_html.Close()

	fh_json, err := os.Create(inventory_json)

	if err != nil {
		log.Fatal(err)
	}

	defer fh_json.Close()

	www_bytes, err := html.Asset("templates/html/inventory.html")

	if err != nil {
		log.Fatal(err)
	}

	www_template, err := template.New("www").Parse(string(www_bytes))

	if err != nil {
		log.Fatal(err)
	}

	// start working

	databases := flag.Args()
	count := len(databases)

	report_ch := make(chan Report)
	done_ch := make(chan bool)
	error_ch := make(chan error)

	// validate all the things

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

			if r.include {
				reports = append(reports, r)
			} else {
				os.Remove(r.path)
			}
		}
	}

	// compress and hash all the things

	// please for to make throttling work
	// (20171220/thisisaaronland)

	for _, r := range reports {

		go func(r Report) {
			Compress(r, report_ch, done_ch, error_ch)
		}(r)
	}

	remaining = len(reports)

	compressed := make([]Report, 0)

	for remaining > 0 {

		select {
		case <-done_ch:
			remaining -= 1
		case err := <-error_ch:
			log.Println("ERROR", err)
		case r := <-report_ch:
			compressed = append(compressed, r)
		}
	}

	// sort things by filename

	by_name := make(map[string]Report)
	names := make([]string, 0)

	for _, r := range compressed {
		n := r.Name
		n = strings.Replace(n, "-latest.db", "", -1)
		names = append(names, n)
		by_name[n] = r
	}

	sort.Strings(names)

	sorted := make([]Report, 0)

	for _, n := range names {
		sorted = append(sorted, by_name[n])
	}

	// finally write things to disk

	now := time.Now()
	lastmod := now.Format(time.RFC3339)

	vars := HTMLVars{
		Reports:      sorted,
		LastModified: lastmod,
	}

	err = www_template.Execute(fh_html, vars)

	if err != nil {
		log.Fatal(err)
	}

	enc, err := json.Marshal(sorted)

	if err != nil {
		log.Fatal(err)
	}

	fh_json.Write(enc)

	fh_html.Close()
	fh_json.Close()

	log.Println("DONE")
}
