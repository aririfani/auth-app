package db

import (
	"github.com/aririfani/auth-app/config"
)

const (
	// SqlLiteDialectParam ...
	SqlLiteDialectParam = "sqlite"
)

type Option func(*DB)

// WithConfig ....
func WithConfig(config config.Config) Option {
	return func(db *DB) {
		db.config = config
	}
}
