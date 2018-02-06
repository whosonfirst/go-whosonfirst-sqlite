package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	wof_index "github.com/whosonfirst/go-whosonfirst-index"
	"github.com/whosonfirst/go-whosonfirst-index/utils"
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"github.com/whosonfirst/go-whosonfirst-sqlite/index"
	"github.com/whosonfirst/go-whosonfirst-sqlite/tables"
	"io"
	"os"
	"runtime"
	"strings"
)

// THIS IS A TOTAL HACK UNTIL WE CAN SORT THINGS OUT IN
// go-whosonfirst-index... (20180206/thisisaaronland)

type Closer struct {
	fh io.Reader
}

func (c Closer) Read(b []byte) (int, error) {
	return c.fh.Read(b)
}

func (c Closer) Close() error {
	return nil
}

func main() {

	valid_modes := strings.Join(wof_index.Modes(), ",")
	desc_modes := fmt.Sprintf("The mode to use importing data. Valid modes are: %s.", valid_modes)

	dsn := flag.String("dsn", ":memory:", "")
	driver := flag.String("driver", "sqlite3", "")

	mode := flag.String("mode", "files", desc_modes)

	all := flag.Bool("all", false, "Index all tables (except geometries which you need to specify explicitly)")
	ancestors := flag.Bool("ancestors", false, "Index the 'ancestors' tables")
	concordances := flag.Bool("concordances", false, "Index the 'concordances' tables")
	geojson := flag.Bool("geojson", false, "Index the 'geojson' table")
	geometries := flag.Bool("geometries", false, "Index the 'geometries' table (requires that libspatialite already be installed)")
	names := flag.Bool("names", false, "Index the 'names' table")
	spr := flag.Bool("spr", false, "Index the 'spr' table")
	live_hard := flag.Bool("live-hard-die-fast", false, "Enable various performance-related pragmas at the expense of possible (unlikely) database corruption")
	timings := flag.Bool("timings", false, "Display timings during and after indexing")
	var procs = flag.Int("processes", (runtime.NumCPU() * 2), "The number of concurrent processes to index data with")

	flag.Parse()

	runtime.GOMAXPROCS(*procs)

	logger := log.SimpleWOFLogger()

	stdout := io.Writer(os.Stdout)
	logger.AddLogger(stdout, "status")

	if *geometries && *driver != "spatialite" {
		logger.Fatal("you asked to index geometries but specified the '%s' driver instead of spatialite", *driver)
	}

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

	if *geojson || *all {

		gt, err := tables.NewGeoJSONTableWithDatabase(db)

		if err != nil {
			logger.Fatal("failed to create 'geojson' table because %s", err)
		}

		to_index = append(to_index, gt)
	}

	if *spr || *all {

		st, err := tables.NewSPRTableWithDatabase(db)

		if err != nil {
			logger.Fatal("failed to create 'spr' table because %s", err)
		}

		to_index = append(to_index, st)
	}

	if *names || *all {

		nm, err := tables.NewNamesTableWithDatabase(db)

		if err != nil {
			logger.Fatal("failed to create 'names' table because %s", err)
		}

		to_index = append(to_index, nm)
	}

	if *ancestors || *all {

		an, err := tables.NewAncestorsTableWithDatabase(db)

		if err != nil {
			logger.Fatal("failed to create 'ancestors' table because %s", err)
		}

		to_index = append(to_index, an)
	}

	if *concordances || *all {

		cn, err := tables.NewConcordancesTableWithDatabase(db)

		if err != nil {
			logger.Fatal("failed to create 'concordances' table because %s", err)
		}

		to_index = append(to_index, cn)
	}

	// see the way we don't check *all here - that's so people who don't have
	// spatialite installed can still use *all (20180122/thisisaaronland)

	if *geometries {

		gm, err := tables.NewGeometriesTableWithDatabase(db)

		if err != nil {
			logger.Fatal("failed to create 'geometries' table because %s", err)
		}

		to_index = append(to_index, gm)
	}

	if len(to_index) == 0 {
		logger.Fatal("You forgot to specify which (any) tables to index")
	}

	cb := func(ctx context.Context, fh io.Reader, args ...interface{}) (interface{}, error) {

		ok, err := utils.IsPrincipalWOFRecord(fh, ctx)

		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, nil
		}

		// HACK - see above
		closer := Closer{fh}

		return feature.LoadWOFFeatureFromReader(closer)
	}

	idx, err := index.NewSQLiteIndexer(db, to_index, cb)

	if err != nil {
		logger.Fatal("failed to create sqlite indexer because %s", err)
	}

	idx.Timings = *timings
	idx.Logger = logger

	err = idx.IndexPaths(*mode, flag.Args())

	if err != nil {
		logger.Fatal("Failed to index paths in %s mode because: %s", *mode, err)
	}

	os.Exit(0)
}
