package main

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-index"
	"github.com/whosonfirst/go-whosonfirst-index/utils"
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"github.com/whosonfirst/go-whosonfirst-sqlite/tables"
	"io"
	"os"
)

func main() {

	mode := flag.String("mode", "files", "The mode to use importing data.")
	dsn := flag.String("dsn", ":memory:", "")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	db, err := database.NewDB(*dsn)

	if err != nil {
		logger.Fatal("unable to create database (%s) because %s", *dsn, err)
	}

	defer db.Close()

	gt, err := tables.NewGeoJSONTable()

	if err != nil {
		logger.Fatal("failed to create geojson table because %s", err)
	}

	err = gt.InitializeTable(db)

	if err != nil {
		logger.Fatal("failed to initialize geojson table because %s", err)
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

		return gt.IndexFeature(db, f)
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

	conn, err := db.Conn()

	if err != nil {
		logger.Fatal("Failed to get DB conn")
	}

	stmt, err := conn.Prepare("select body from whosonfirst LIMIT 1")

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
