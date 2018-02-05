package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-json"
	"github.com/whosonfirst/go-whosonfirst-json/utils"
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
	flag.Var(&props, "property", "...")

	flag.Parse()

	for _, path := range flag.Args() {

		d, err := json.LoadDocumentFromFile(path)

		if err != nil {
			log.Fatal(err)
		}

		err = utils.EnsureProperties(d.Bytes(), props)

		if err != nil {
			log.Fatal(err)
		}

		for _, p := range props {

			log.Println(path, p, utils.StringProperty(d.Bytes(), []string{p}, ""))
		}
	}
}
