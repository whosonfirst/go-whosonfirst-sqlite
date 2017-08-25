package tables

import (
	"errors"
	_ "fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
)

type SPRTable struct {
	sqlite.Table
	name string
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
	return ""
}

func (t *SPRTable) IndexFeature(db sqlite.Database, f geojson.Feature) error {
	return errors.New("Please implement me")
}
