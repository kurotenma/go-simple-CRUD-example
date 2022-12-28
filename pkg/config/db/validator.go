package dbConfigTools

import (
	"errors"
)

func ValidateDB(db PostgresDB) error {
	if db.DBHost == "" {
		return errors.New("DB_HOST is required")
	}
	if db.DBUser == "" {
		return errors.New("DB_USER is required")
	}
	if db.DBPass == "" {
		return errors.New("DB_PASS is required")
	}
	if db.DBPort == "" {
		return errors.New("DB_PORT is required")
	}
	if db.DBName == "" {
		return errors.New("DB_NAME is required")
	}
	return nil
}
