package utils

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func EnsurePropertiesAny(body []byte, properties []string) error {

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

func EnsureProperties(body []byte, properties []string) error {

	for _, path := range properties {

		r := gjson.GetBytes(body, path)

		if !r.Exists() {
			msg := fmt.Sprintf("Feature is missing a %s property", path)
			return errors.New(msg)
		}
	}

	return nil
}

func Int64Property(body []byte, possible []string, d int64) int64 {

	for _, path := range possible {

		v := gjson.GetBytes(body, path)

		if v.Exists() {
			return v.Int()
		}
	}

	return d
}

func Int64PropertyArray(body []byte, possible []string) []int64 {

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

func StringProperty(body []byte, possible []string, d string) string {

	for _, path := range possible {

		v := gjson.GetBytes(body, path)

		if v.Exists() {
			return v.String()
		}
	}

	return d
}
