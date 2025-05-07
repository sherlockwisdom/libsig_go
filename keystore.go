package main

import (
	"database/sql"
	"fmt"

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

func (k *Keystore) Init(tableName string) error {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}

	k.connection = db

	sqlStmt := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s ( 
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	pnt TEXT NOT NULL UNIQUE, 
	pk BLOB NOT NULL, 
	_pk BLOB NOT NULL, 
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`, tableName)

	_, err = db.Exec(sqlStmt)

	if err != nil {
		return err
	}
	return err
}

func (k *Keystore) Store(keypair Keypairs) error {
	prKey, pubKey := keypair.PrivateKey.Bytes(), keypair.PrivateKey.PublicKey.Bytes()

	tx, err := k.connection.Begin()
	if err != nil {
		return err
	}

	sqlInsertStmt := fmt.Sprintf(`

	)`, tableName)

	stmt, err := tx.Prepare(sqlInsertStmt)
	if err != nil {
		return err
	}
}

func (k *Keystore) Close() {
	defer k.connection.Close()
}
