package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-brands"
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"github.com/whosonfirst/go-whosonfirst-sqlite/index"
	"github.com/whosonfirst/go-whosonfirst-sqlite/tables"
	"io"
	"os"
	"runtime"
)

func main() {

	valid_modes := strings.Join(index.Modes(), ",")
	desc_modes := fmt.Sprintf("The mode to use importing data. Valid modes are: %s.", valid_modes)

	dsn := flag.String("dsn", ":memory:", "")
	driver := flag.String("driver", "sqlite3", "")

	mode := flag.String("mode", "files", desc_modes)

	all := flag.Bool("all", false, "Index all tables (except geometries which you need to specify explicitly)")
	live_hard := flag.Bool("live-hard-die-fast", false, "Enable various performance-related pragmas at the expense of possible (unlikely) database corruption")
	timings := flag.Bool("timings", false, "Display timings during and after indexing")
	var procs = flag.Int("processes", (runtime.NumCPU() * 2), "The number of concurrent processes to index data with")

	flag.Parse()

	runtime.GOMAXPROCS(*procs)

	logger := log.SimpleWOFLogger()

	stdout := io.Writer(os.Stdout)
	logger.AddLogger(stdout, "status")

	db, err := database.NewDBWithDriver(*driver, *dsn)

	if err != nil {
		logger.Fatal("unable to create database (%s) because %s", *dsn, err)
	}

	defer db.Close()

	if *live_hard {

		err = db.LiveHardDieFast()

		if err != nil {
			logger.Fatal("Unable to live hard and die fast so just dying fast instead, because %s", err)
		}
	}

	to_index := make([]sqlite.Table, 0)

	br, err := tables.NewBrandsTableWithDatabase(db)

	if err != nil {
		logger.Fatal("failed to create 'geojson' table because %s", err)
	}

	to_index = append(to_index, br)

	if len(to_index) == 0 {
		logger.Fatal("You forgot to specify which (any) tables to index")
	}

	cb := func(ctx context.Context, fh io.Reader, args ...interface{}) (interface{}, error) {

		path, err := index.PathForContext(ctx)

		if err != nil {
			return err
		}

		return whosonfirst.LoadWOFBrandFromReader(fh)
	}

	idx, err := index.NewSQLiteIndexer(db, tables, cb)

	if err != nil {
		logger.Fatal("failed to create sqlite indexer because %s", err)
	}

	idx.Timings = *timings
	idx.Logger = logger

	err = idx.IndexPaths(flag.Args())

	if err != nil {
		logger.Fatal("Failed to index paths in %s mode because: %s", *mode, err)
	}

	os.Exit(0)
}
