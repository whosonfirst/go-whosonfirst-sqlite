package utils

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"os"
)

func CreateTableIfNecessary(db sqlite.Database, t sqlite.Table) error {

	create := false

	if db.DSN() == ":memory:" {
		create = true
	} else {

		info, err := os.Stat(db.DSN())

		if info.IsDir() {
			return errors.New("path is a directory")
		}

		if os.IsNotExist(err) {
			create = true
		}
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
