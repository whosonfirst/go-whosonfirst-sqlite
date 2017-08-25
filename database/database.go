package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDatabase struct {
	conn *sql.DB
	dsn  string
}

func NewDB(dsn string) (*SQLiteDatabase, error) {

	conn, err := sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	db := SQLiteDatabase{
		conn: conn,
		dsn:  dsn,
	}

	return &db, err
}

func (db *SQLiteDatabase) Conn() (*sql.DB, error) {
	return db.conn, nil
}

func (db *SQLiteDatabase) DSN() string {
	return db.dsn
}

func (db *SQLiteDatabase) Close() error {
	return db.conn.Close()
}
