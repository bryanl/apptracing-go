package db

import "github.com/jmoiron/sqlx"

// DB wraps SQLx so it can be instrumented
type DB struct {
	*sqlx.DB
}

// Init connects the database
func Init(driverName, dataSourceName string) *DB {
	x := sqlx.MustConnect(driverName, dataSourceName)

	return &DB{
		DB: x,
	}
}
