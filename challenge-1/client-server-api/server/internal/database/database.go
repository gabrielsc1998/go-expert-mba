package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Connect() error {
	const BASE_DIR = "../data"
	exists, _ := os.Stat(BASE_DIR)
	if exists == nil {
		os.Mkdir(BASE_DIR, 0777)
	}
	db, err := gorm.Open(sqlite.Open(BASE_DIR+"/quotation.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	d.DB = db
	return nil
}
