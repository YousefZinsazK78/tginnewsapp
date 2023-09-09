package database

import "database/sql"

type Database struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) Database {
	return Database{
		db,
	}
}
