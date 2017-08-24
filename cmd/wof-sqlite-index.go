package main

import (
	"context"
	"database/sql"
	_ "encoding/json"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-index"
	"github.com/whosonfirst/go-whosonfirst-index/utils"
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/go-whosonfirst-sqlite/schema"
	"io"
	"os"
)

func main() {

	mode := flag.String("mode", "files", "The mode to use importing data.")
	database := flag.String("database", ":memory:", "")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	// If you're reading this then that means it is probably still "too soon".
	// All (or most) of the SQL specific stuff will be moved in to proper
	// packages... but not today (20170824/thisisaaronland)

	db, err := sql.Open("sqlite3", *database)

	if err != nil {
		logger.Fatal("unable to create database (%s) because %s", *database, err)
	}

	defer db.Close()

	// For example... check to see that the database doesn't already exist and
	// stuff like that (20170824/thisisaaronland)

	sql := schema.WhosOnFirstSchema()

	_, err = db.Exec(sql)

	if err != nil {
		logger.Fatal("failed to load schema because %s", err)
	}

	cb := func(fh io.Reader, ctx context.Context, args ...interface{}) error {

		ok, err := utils.IsPrincipalWOFRecord(fh, ctx)

		if err != nil {
			return err
		}

		if !ok {
			return nil
		}

		f, err := feature.LoadWOFFeatureFromReader(fh)

		if err != nil {
			return err
		}

		str_id := f.Id()
		body := f.Bytes()

		if err != nil {
			return err
		}

		tx, err := db.Begin()

		if err != nil {
			return err
		}

		stmt, err := tx.Prepare("insert into whosonfirst(id, body) values(?, ?)")

		if err != nil {
			return err
		}

		defer stmt.Close()

		str_body := string(body)

		_, err = stmt.Exec(str_id, str_body)

		if err != nil {
			return err
		}

		tx.Commit()

		return nil
	}

	indexer, err := index.NewIndexer(*mode, cb)

	if err != nil {
		logger.Fatal("Failed to create new indexer because %s", err)
	}

	err = indexer.IndexPaths(flag.Args())

	if err != nil {
		logger.Fatal("Failed to index paths in %s mode because %s", *mode, err)
	}

	logger.Status("DONE INDEXING")

	stmt, err := db.Prepare("select body from whosonfirst LIMIT 1")

	if err != nil {
		logger.Fatal("Failed to prepare statement because %s", err)
	}

	defer stmt.Close()

	var body string
	row := stmt.QueryRow()

	row.Scan(&body)

	if err != nil {
		logger.Fatal("Failed to scane row because %s", err)
	}

	f, err := feature.LoadFeature([]byte(body))

	if err != nil {
		logger.Fatal("Failed to scane row because %s", err)
	}

	logger.Status("RETRIEVED %s (%s)", f.Id(), f.Name())
	os.Exit(0)
}
