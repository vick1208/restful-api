package helper

import "database/sql"

func RollbackOrCommit(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRb := tx.Rollback()
		PanicIE(errRb)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIE(errCommit)
	}
}
