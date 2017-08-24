package sqlite

// this is here so we can pass both sql.Row and sql.Rows to the
// SQLiteQueryRowToPgisRowFunc below (20170824/thisisaaronland)

type SQLiteResultSet interface {
	Scan(dest ...interface{}) error
}

type SQLiteRow interface {
     // uhhh.... (20170824/thisisaaronland)
}

type SQLiteQueryRowFunc func(row SQLiteResultSet) (*SQLiteRow, error)
