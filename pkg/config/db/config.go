package dbConfigTools

import (
	"github.com/joho/godotenv"
	"os"
)

type PostgresDB struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
}

func InitDB() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	env := os.Getenv("ENV_MODE")
	if env != "DEVELOPMENT" {
		if err := godotenv.Load(".env"); err != nil {
			return err
		}
	} else {
		if err := godotenv.Load(".env"); err != nil {
			return err
		}
	}
	return nil
}

func LoadPostgresDB() (PostgresDB, error) {
	var db PostgresDB
	if err := InitDB(); err != nil {
		return db, err
	}
	db.DBHost = os.Getenv("DB_HOST")
	db.DBUser = os.Getenv("DB_USER")
	db.DBPass = os.Getenv("DB_PASS")
	db.DBPort = os.Getenv("DB_PORT")
	db.DBName = os.Getenv("DB_NAME")
	if err := ValidateDB(db); err != nil {
		return db, err
	}

	return db, nil
}

func (db *PostgresDB) Url() string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := "postgres://" +
		db.DBUser + ":" +
		db.DBPass + "@" +
		db.DBHost + ":" +
		db.DBPort + "/" +
		db.DBName
	return url
}
