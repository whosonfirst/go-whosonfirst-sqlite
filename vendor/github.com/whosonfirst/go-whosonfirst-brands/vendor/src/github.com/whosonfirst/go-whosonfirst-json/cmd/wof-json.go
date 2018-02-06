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
			log.Println(path, p, properties.StringProperty(doc, []string{p}, ""))
		}
	}
}
