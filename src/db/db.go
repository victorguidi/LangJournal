package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
	// Storage
}

// type Storage interface {
// 	GetBaseQuestions(language string) (interface{}, error)
// }

func New() (*DB, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &DB{
		db: db,
	}, nil

}

func (d *DB) GetBaseQuestions(language string) (interface{}, error) {
	var obj []string
	return obj, nil
}
