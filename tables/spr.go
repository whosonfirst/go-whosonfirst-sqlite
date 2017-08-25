package tables

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
)

type SPRTable struct {
	sqlite.Table
	name string
}

type SPRRow struct {
     Id     int64		// properties.wof:id	INTEGER
     Name   string		// properties.wof:name  VARCHAR(255)
     Placetype	string		// properties.wof:placetype VARCHAR(64)
     Country	string		// properties.wof:country VARCHAR(2)
     IsCurrent	int64
     IsCeased	int64
     IsDeprecated int64
     IsSuperseded int64
     IsSuperseding int64
}

func NewSPRTable() (*SPRTable, error) {

	t := SPRTable{
		name: "spr",
	}

	return &t, nil
}

func (t *SPRTable) InitializeTable(db sqlite.Database) error {
	return nil
}

func (t *SPRTable) Name() string {
	return t.name
}

func (t *SPRTable) Schema() string {

	sql := `CREATE TABLE %s (
			id INTEGER NOT NULL PRIMARY KEY,
			name TEXT,
			placetype TEXT,
			country TEXT,
			is_current INTEGER,
			is_deprecated INTEGER,
			is_ceased INTEGER,
			is_superseded INTEGER,
			is_superseding INTEGER
	)`

	return fmt.Sprintf(sql, t.Name())
}

func (t *SPRTable) IndexFeature(db sqlite.Database, f geojson.Feature) error {
	return errors.New("Please implement me")
}
