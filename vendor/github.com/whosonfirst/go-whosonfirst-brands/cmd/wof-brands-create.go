package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/tidwall/pretty"
	"github.com/whosonfirst/go-whosonfirst-brands"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var name = flag.String("name", "", "...")
	var root = flag.String("root", "", "...")

	flag.Parse()

	brand, err := brands.NewBrand(*name)

	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(brand)
	pretty_b := pretty.Pretty(b)

	if *root == "" {
		os.Stdout.Write(pretty_b)
		os.Exit(0)
	}

	abs_path, err := uri.Id2AbsPath(*root, brand.WOFBrandId)

	if err != nil {
		log.Fatal(err)
	}

	// https://github.com/whosonfirst/go-whosonfirst-brands/issues/1
	abs_path = strings.Replace(abs_path, ".geojson", ".json", 1)

	abs_root := filepath.Dir(abs_path)

	_, err = os.Stat(abs_root)

	if os.IsNotExist(err) {

		err = os.MkdirAll(abs_root, 0755)

		if err != nil {
			log.Fatal(err)
		}
	}

	fh, err := os.Create(abs_path)

	if err != nil {
		log.Fatal(err)
	}

	fh.Write(pretty_b)
	fh.Close()

	fmt.Println(abs_path)
}
