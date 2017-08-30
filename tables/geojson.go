package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
)

type GeoJSONTable struct {
	sqlite.Table
	name string
}

type GeoJSONRow struct {
	Id           int64
	Body         string
	LastModified int64
}

func NewGeoJSONTable() (*GeoJSONTable, error) {

	t := GeoJSONTable{
		name: "geojson",
	}

	return &t, nil
}

func (t *GeoJSONTable) Name() string {
	return t.name
}

func (t *GeoJSONTable) Schema() string {
	return fmt.Sprintf("CREATE TABLE %s (id INTEGER NOT NULL PRIMARY KEY, body TEXT, lastmodified INTEGER)", t.Name())
}

func (t *GeoJSONTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *GeoJSONTable) IndexFeature(db sqlite.Database, f geojson.Feature) error {

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	str_id := f.Id()
	body := f.Bytes()

	lastmod := whosonfirst.LastModified(f)

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	sql := fmt.Sprintf("INSERT INTO %s (id, body, lastmodified) values(?, ?, ?)", t.Name())

	stmt, err := tx.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	str_body := string(body)

	_, err = stmt.Exec(str_id, str_body, lastmod)

	if err != nil {
		return err
	}

	return tx.Commit()
}
