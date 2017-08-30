package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type SQLiteDatabase struct {
	conn *sql.DB
	dsn  string
	mu   *sync.Mutex
}

func NewDB(dsn string) (*SQLiteDatabase, error) {

	conn, err := sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	mu := new(sync.Mutex)

	db := SQLiteDatabase{
		conn: conn,
		dsn:  dsn,
		mu:   mu,
	}

	return &db, err
}

func (db *SQLiteDatabase) Lock() {
	db.mu.Lock()
}

func (db *SQLiteDatabase) Unlock() {
	db.mu.Unlock()
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
