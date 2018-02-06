package properties

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-json"
	"strings"
)

func EnsurePropertiesAny(doc json.Document, properties []string) error {

     body := doc.Bytes()
	return EnsurePropertiesAnyBytes(body, properties)
}

func EnsurePropertiesAnyBytes(body []byte, properties []string) error {

	for _, path := range properties {

		r := gjson.GetBytes(body, path)

		if r.Exists() {
			return nil
		}
	}

	str_props := strings.Join(properties, ";")

	msg := fmt.Sprintf("Feature is missing any of the following properties: %s", str_props)
	return errors.New(msg)
}

func EnsureProperties(doc json.Document, properties []string) error {

	body := doc.Bytes()
	return EnsurePropertiesBytes(body, properties)
}

func EnsurePropertiesBytes(body []byte, properties []string) error {

	for _, path := range properties {

		r := gjson.GetBytes(body, path)

		if !r.Exists() {
			msg := fmt.Sprintf("Feature is missing a %s property", path)
			return errors.New(msg)
		}
	}

	return nil
}

func Int64Property(doc json.Document, possible []string, d int64) int64 {

	body := doc.Bytes()

	for _, path := range possible {

		v := gjson.GetBytes(body, path)

		if v.Exists() {
			return v.Int()
		}
	}

	return d
}

func Int64PropertyArray(doc json.Document, possible []string) []int64 {

	body := doc.Bytes()

	results := make([]int64, 0)

	for _, p := range possible {

		rsp := gjson.GetBytes(body, p)

		if rsp.Exists() {

			for _, id := range rsp.Array() {
				results = append(results, id.Int())
			}

			break
		}
	}

	return results
}

func StringProperty(doc json.Document, possible []string, d string) string {

	body := doc.Bytes()

	for _, path := range possible {

		v := gjson.GetBytes(body, path)

		if v.Exists() {
			return v.String()
		}
	}

	return d
}

func StringPropertyArray(doc json.Document, possible []string) []string {

	body := doc.Bytes()

	results := make([]string, 0)

	for _, p := range possible {

		rsp := gjson.GetBytes(body, p)

		if rsp.Exists() {

			for _, id := range rsp.Array() {
				results = append(results, id.String())
			}

			break
		}
	}

	return results
}
