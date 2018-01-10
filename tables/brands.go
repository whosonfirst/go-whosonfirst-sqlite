package tables

import (
       "errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
)

type BrandsTable struct {
	sqlite.Table
	name string
}

type BrandsRow struct {
	Id           int64
	Name         string
	Size         string
	IsCurrent    int
	LastModified int64
}

func NewBrandsTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewBrandsTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewBrandsTable() (sqlite.Table, error) {

	t := BrandsTable{
		name: "brands",
	}

	return &t, nil
}

func (t *BrandsTable) Name() string {
	return t.name
}

func (t *BrandsTable) Schema() string {

	sql := `CREATE TABLE %s (
	       id INTEGER NOT NULL,
	       name TEXT,
	       size TEXT,
	       is_current INTEGER,
	       lastmodified INTEGER
	);

	CREATE INDEX brands_by_name ON %s (name, size, is_current);
	CREATE INDEX brands_by_name_current ON %s (name, is_current);	
	CREATE INDEX brands_by_lastmod ON %s (lastmodified);
	CREATE INDEX brands_by_id ON %s (id);
	`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(sql, t.Name(), t.Name(), t.Name(), t.Name(), t.Name())
}

func (t *BrandsTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *BrandsTable) IndexFeature(db sqlite.Database, f geojson.Feature) error {

     return errors.New("Please write me")
}
