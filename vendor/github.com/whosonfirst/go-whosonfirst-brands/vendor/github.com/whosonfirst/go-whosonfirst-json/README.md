# go-whosonfirst-json

There are many packages for working with JSON documents. This one is ours.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

This is work in progress.

This is mostly just a set of utility methods for loading JSON documents, testing for specific properties and retreiving specific properties as `string` or `int64` values. There is no reason you _couldn't_ use this package directly but generally it is expected to be wrapped by a purpose-specific JSON package, for Who's On First features or brands.

Testing and retrieving property is done using @tidwall 's `gjson` package under the hood. It is not currently possible to _assign_ values to JSON documents.

## Usage

```
package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-json"
	"github.com/whosonfirst/go-whosonfirst-json/properties"
	"log"
)

type PropertiesFlags []string

func (p *PropertiesFlags) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *PropertiesFlags) Set(path string) error {
	*p = append(*p, path)
	return nil
}

func main() {

	var props PropertiesFlags
	flag.Var(&props, "property", "A JSON property in dot-notation form to test for and display.")

	flag.Parse()

	for _, path := range flag.Args() {

		doc, err := json.LoadDocumentFromFile(path)

		if err != nil {
			log.Fatal(err)
		}

		err = properties.EnsureProperties(doc, props)

		if err != nil {
			log.Fatal(err)
		}

		for _, p := range props {
			log.Println(path, p, properties.StringProperty(doc, []string{p}, "some default value"))
		}
	}
}
```

## Tools

### wof-json

```
./bin/wof-json -property 'wof:brand_id' -property 'wof:brand_name' ../whosonfirst-brands/data/110/871/892/5/1108718925.json
2018/02/05 16:55:20 ../whosonfirst-brands/data/110/871/892/5/1108718925.json wof:brand_id 1108718925
2018/02/05 16:55:20 ../whosonfirst-brands/data/110/871/892/5/1108718925.json wof:brand_name Mapzen
```

## See also

* github.com/tidwall/gjson
