package utils

import "database/sql"

func Transaction(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	// do something with tx

	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
