package database

import (
	"database/sql"
	"fmt"
)

type Database struct {
	DB *sql.DB
}

type DatabaseOptions struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Connect(options DatabaseOptions) {
	connectionAddr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", options.User, options.Password, options.Host, options.Port, options.Name)
	db, err := sql.Open(options.Driver, connectionAddr)
	if err != nil {
		panic(err)
	}
	d.DB = db
}
