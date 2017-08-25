package utils

import (
	"github.com/whosonfirst/go-whosonfirst-sqlite"
)

func CreateTableIfNecessary(db sqlite.Database, t sqlite.Table) error {

	create := false

	if db.DSN() == ":memory:" {
		create = true
	}

	if create {

		sql := t.Schema()

		conn, err := db.Conn()

		if err != nil {
			return err
		}

		_, err = conn.Exec(sql)

		if err != nil {
			return err
		}
	}

	return nil
}
