# go-whosonfirst-sqlite

Go package for working with Who's On First documents and SQLite.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

It's probably still too soon. Lots of things may still change.

## Tables

### geojson

_Please write me._

### names

_Please write me._

### spr

_Please write me._

## Interfaces

### Database

```
type Database interface {
     Conn() (*sql.DB, error)
     DSN() string
     Close() error
}
```

### Table

```
type Table interface {
     Name() string
     Schema() string
     InitializeTable(Database) error
     IndexFeature(Database, geojson.Feature) error
}
```

Where `geojson.Feature` is defined in the [go-whosonfirst-geojson-v2](https://github.com/whosonfirst/go-whosonfirst-geojson-v2#geojsonfeature) package.

## Tools

### wof-sqlite-index

```
./bin/wof-sqlite-index -h
Usage of ./bin/wof-sqlite-index:
  -all
    	Index all tables
  -dsn string
    	 (default ":memory:")
  -geojson
    	Index the 'geojson' table
  -mode string
    	The mode to use importing data. Valid modes are: directory,feature,feature-collection,geojson-ls,meta,path,repo (default "files")
  -names
    	Index the 'names' table
  -spr
    	Index the 'spr' table
```

For example:

```
./bin/wof-sqlite-index -dsn microhoods.db -all -mode meta /usr/local/data/whosonfirst-data/meta/wof-microhood-latest.csv
```

You can also use `wof-sqlite-index` in combination with the [go-whosonfirst-api](https://github.com/whosonfirst/go-whosonfirst-api) `wof-api` tool and populate your SQLite database from API results. For example, here's how you might index all the neighbourhoods in Montreal:

```
/usr/local/bin/wof-api -param method=whosonfirst.places.getDescendants -param id=101736545 \
-param placetype=neighbourhood -param api_key=mapzen-xxxxxx -geojson-ls | \
/usr/local/bin/wof-sqlite-index -dsn microhoods.db -all -mode geojson-ls STDIN
```