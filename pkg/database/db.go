package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Rows interface {
    Close() error
    Scan(dest ...interface{}) error
    Next() bool
    Err() error
}

type DB struct {
    *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}