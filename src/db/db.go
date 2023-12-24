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
}

type BaseQuestions struct {
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	IsCorrect bool   `json:"isCorrect"`
}

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

func (d *DB) GetBaseQuestions() ([]BaseQuestions, error) {

	questions, err := d.db.Query(`
    SELECT q.question, a.answer, a.is_correct FROM defaultquestions q
    LEFT JOIN Answers a ON q.id = a.question_id
  `)

	if err != nil {
		return nil, err
	}

	var baseQuestions []BaseQuestions

	for questions.Next() {
		var question BaseQuestions
		err = questions.Scan(&question.Question, &question.Answer, &question.IsCorrect)
		if err != nil {
			return nil, err
		}
		baseQuestions = append(baseQuestions, question)
	}

	return baseQuestions, nil
}
