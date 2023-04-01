package utils

import "database/sql"

func Transaction(tx *sql.Tx) {
	if err := recover(); err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
}
