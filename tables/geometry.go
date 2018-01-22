package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/geometry"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	_ "log"
)

type GeometryTable struct {
	sqlite.Table
	name string
}

type GeometryRow struct {
	Id           int64
	Body         string
	LastModified int64
}

func NewGeometryTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewGeometryTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewGeometryTable() (sqlite.Table, error) {

	t := GeometryTable{
		name: "geometry_whosonfirst",
	}

	return &t, nil
}

func (t *GeometryTable) Name() string {
	return t.name
}

func (t *GeometryTable) Schema() string {

	// really this should probably be the SPR table + geom but
	// let's just get this working first and then make it fancy
	// (20180109/thisisaaronland)

	// https://www.gaia-gis.it/spatialite-1.0a/SpatiaLite-tutorial.html
	// http://www.gaia-gis.it/gaia-sins/spatialite-sql-4.3.0.html

	// Note the InitSpatialMetaData() command because this:
	// https://stackoverflow.com/questions/17761089/cannot-create-column-with-spatialite-unexpected-metadata-layout

	sql := `CREATE TABLE %s (
		id INTEGER NOT NULL PRIMARY KEY,
		placetype TEXT,
		lastmodified INTEGER
	);

	SELECT InitSpatialMetaData();
	SELECT AddGeometryColumn('%s', 'geom', 4326, 'GEOMETRY', 'XY');
	SELECT CreateSpatialIndex('%s', 'geom');

	CREATE INDEX geojson_by_lastmod ON %s (lastmodified);`

	return fmt.Sprintf(sql, t.Name(), t.Name(), t.Name(), t.Name())
}

func (t *GeometryTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *GeometryTable) IndexFeature(db sqlite.Database, f geojson.Feature) error {

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	str_id := f.Id()
	pt := f.Placetype()

	lastmod := whosonfirst.LastModified(f)

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	// apparently this really needs to be WKT or spatialite
	// will complain about stuff

	str_geom, err := geometry.ToString(f)

	if err != nil {
		return err
	}

	sql := fmt.Sprintf(`INSERT OR REPLACE INTO %s (
		id, placetype, geom, lastmodified
	) VALUES (
		?, ?, GeomFromGeoJSON('%s'), ?
	)`, t.Name(), str_geom)

	stmt, err := tx.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(str_id, pt, lastmod)

	if err != nil {
		return err
	}

	return tx.Commit()
}
