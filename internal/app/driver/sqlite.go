package driver

import (
	"database/sql"
	"errors"
	"github.com/aririfani/auth-app/config"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteOption struct {
	DBName string
}

func OpenSqlite(cfg config.Config) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", cfg.GetString("database.sqlite.db_file"))

	if err != nil {
		return
	}

	migrate(db)

	if db == nil {
		err = errors.New("error db")
		return
	}

	return
}

func migrate(db *sql.DB) {
	sqlQuery := `CREATE TABLE IF NOT EXISTS users(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username VARCHAR NOT NULL,
		phone VARCHAR ,
		password VARCHAR,
		registered_at TIMESTAMP,
		role VARCHAR
    );`

	_, err := db.Exec(sqlQuery)

	if err != nil {
		panic(err)
	}
}
