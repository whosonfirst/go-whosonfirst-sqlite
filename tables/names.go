package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-names/tags"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
)

type NamesTable struct {
	sqlite.Table
	name string
}

type NamesRow struct {
	Id           int64
	Placetype    string
	Country      string
	Language     string
	ExtLang      string
	Script       string
	Region       string
	Variant      string
	Extension    string
	PrivateUse   string
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
	return fmt.Sprintf(`CREATE TABLE %s (
	       id INTEGER NOT NULL,
	       placetype TEXT,
	       country TEXT,
	       language TEXT,
	       extlang TEXT,
	       script TEXT,
	       region TEXT,
	       variant TEXT,
	       extension TEXT,
	       privateuse TEXT,
	       name TEXT,
	       lastmodified INTEGER
	)`, t.Name())
}

func (t *NamesTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *NamesTable) IndexFeature(db sqlite.Database, f geojson.Feature) error {

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	id := f.Id()

	sql := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, t.Name())

	stmt, err := tx.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	pt := f.Placetype()
	co := whosonfirst.Country(f)

	lastmod := whosonfirst.LastModified(f)
	names := whosonfirst.Names(f)

	for tag, names := range names {

		lt, err := tags.NewLangTag(tag)

		if err != nil {
			return err
		}

		for _, n := range names {

			if err != nil {
				return err
			}

			sql := fmt.Sprintf(`INSERT INTO %s (
	    			id, placetype, country,
				language, extlang,
				region, script, variant,
	    			extension, privateuse,
				name,
	    			lastmodified
			) VALUES (
	    		  	?, ?, ?,
				?, ?,
				?, ?, ?,
				?, ?,
				?,
				?
			)`, t.Name())

			stmt, err := tx.Prepare(sql)

			if err != nil {
				return err
			}

			defer stmt.Close()

			_, err = stmt.Exec(id, pt, co, lt.Language(), lt.ExtLang(), lt.Script(), lt.Region(), lt.Variant(), lt.Extension(), lt.PrivateUse(), n, lastmod)

			if err != nil {
				return err
			}

		}
	}

	return tx.Commit()
}
