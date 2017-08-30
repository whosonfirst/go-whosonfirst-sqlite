package tables

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
)

type NamesTable struct {
	sqlite.Table
	name string
}

type NamesRow struct {
	Id           int64
	Language     string
	Name         string
	LastModified int64
}

func NewNamesTable() (*NamesTable, error) {

	t := NamesTable{
		name: "names",
	}

	return &t, nil
}

func (t *NamesTable) Name() string {
	return t.name
}

func (t *NamesTable) Schema() string {
	return fmt.Sprintf("CREATE TABLE %s (id INTEGER NOT NULL PRIMARY KEY, language TEXT, name TEXT, lastmodified INTEGER)", t.Name())
}

func (t *NamesTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *NamesTable) IndexFeature(db sqlite.Database, f geojson.Feature) error {

	return errors.New("please implement me")

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	id := f.Id()

	lastmod := whosonfirst.LastModified(f)

	language := "fix me"
	name := "fix me"

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	sql := fmt.Sprintf("INSERT INTO %s (id, language, name, lastmodified) values(?, ?, ?, ?)", t.Name())

	stmt, err := tx.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, language, name, lastmod)

	if err != nil {
		return err
	}

	return tx.Commit()
}
