package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// https://github.com/mattn/go-sqlite3/blob/v1.14.28/_example/simple/simple.go

type KeystoreInterface interface {
	Init() error
	Close()
}

type Keystore struct {
	connection *sql.DB
}

func (k *Keystore) Init(filename string) error {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}

	k.connection = db

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS _crypto ( 
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	pnt TEXT NOT NULL UNIQUE, 
	prKey BLOB NOT NULL, 
	pubKey BLOB NOT NULL, 
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		return err
	}
	return err
}

func (k *Keystore) Store(prKey []byte, pubKey []byte, pnt string) error {
	tx, err := k.connection.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO _crypto (pnt, prKey, pubKey) values(?, ?, ?)`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(pnt, prKey, pubKey)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (k *Keystore) Fetch(pnt string) ([]byte, []byte, error) {
	stmt, err := k.connection.Prepare("select prKey, pubKey from _crypto where pnt = ?")
	var prKey, pubKey []byte

	if err != nil {
		return prKey, pubKey, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(pnt).Scan(&prKey, &pubKey)
	if err != nil {
		return prKey, pubKey, err
	}
	return prKey, pubKey, nil
}

func (k *Keystore) Close() {
	defer k.connection.Close()
}
