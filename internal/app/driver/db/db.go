package db

import (
	"database/sql"
	"errors"
	"github.com/aririfani/auth-app/config"
	"github.com/aririfani/auth-app/internal/app/driver"
)

type DB struct {
	config config.Config
}

// IDB ...
type IDB interface {
	Manager(string) (*sql.DB, error)
}

// New ...
func New(callbacks ...Option) IDB {
	db := new(DB)
	for _, callback := range callbacks {
		callback(db)
	}
	return db
}

// Manager ...
func (db *DB) Manager(dialect string) (dbraw *sql.DB, err error) {
	switch dialect {
	case SqlLiteDialectParam:
		if !db.config.GetBool("database.sqlite.is_active") {
			return
		}
		dbraw, err = driver.OpenSqlite(db.config)
	default:
		err = errors.New("Undefined connection; ignore error if desired")
	}

	return
}
