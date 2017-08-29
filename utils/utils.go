package utils

import (
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"os"
)

func CreateTableIfNecessary(db sqlite.Database, t sqlite.Table) error {

	create := false

	if db.DSN() == ":memory:" {
		create = true
	} else {

		_, err := os.Stat(db.DSN())

		if os.IsNotExist(err) {
			create = true
		} else {

			conn, err := db.Conn()

			if err != nil {
				return err
			}

			sql := "SELECT name FROM sqlite_master WHERE type='table'"

			rows, err := conn.Query(sql)

			if err != nil {
				return err
			}

			defer rows.Close()

			create = true

			for rows.Next() {

				var name string
				err := rows.Scan(&name)

				if err != nil {
					return err
				}

				if name == t.Name() {
					create = false
					break
				}
			}
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
