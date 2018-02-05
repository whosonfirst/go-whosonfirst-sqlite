package json

import (
	gojson "encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type Document interface {
	Bytes() []byte
	String() string
}

type JSONDocument struct {
	Document
	body []byte
}

func (d *JSONDocument) Bytes() []byte {
	return d.body
}

func (d *JSONDocument) String() string {
	return string(d.Bytes())
}

func NewJSONDocument(body []byte) (Document, error) {

	d := JSONDocument{
		body: body,
	}

	return &d, nil
}

func LoadDocument(body []byte) (Document, error) {

	return NewJSONDocument(body)
}

func LoadDocumentFromReader(fh io.Reader) (Document, error) {

	body, err := UnmarshalDocumentFromReader(fh)

	if err != nil {
		return nil, err
	}

	return LoadDocument(body)
}

func LoadDocumentFromFile(path string) (Document, error) {

	body, err := UnmarshalDocumentFromFile(path)

	if err != nil {
		return nil, err
	}

	return LoadDocument(body)
}

func UnmarshalDocument(body []byte) ([]byte, error) {

	var stub interface{}
	err := gojson.Unmarshal(body, &stub)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func UnmarshalDocumentFromReader(fh io.Reader) ([]byte, error) {

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return nil, err
	}

	return UnmarshalDocument(body)
}

func UnmarshalDocumentFromFile(path string) ([]byte, error) {

	fh, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	return UnmarshalDocumentFromReader(fh)
}
