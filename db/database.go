package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName    = "db/testDb"
	structure = "db/database/structure.sql"
	data      = "db/database/data.sql"
)

type Database struct{}

func (d *Database) Init() (*sql.DB, error) {
	dbPath := dbName + ".db"

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(-1)

	err = d.createTables(db)
	if err != nil {
		return nil, err
	}

	// Verileri yükle
	isEmpty, err := d.isTableEmpty(db, "construction_stages")
	if err != nil {
		return nil, err
	}
	if isEmpty {
		err = d.loadData(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func (d *Database) createTables(db *sql.DB) error {
	sql, err := ioutil.ReadFile(structure)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sql))
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) loadData(db *sql.DB) error {
	sql, err := ioutil.ReadFile(data)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sql))
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) isTableEmpty(db *sql.DB, tableName string) (bool, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s LIMIT 1", tableName)
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 0, nil
}
