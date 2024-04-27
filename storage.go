package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(accountId int) error
	UpdateAccount(*Account) error
	GetAccountById(accountId int) (*Account, error)
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgreStore, error) {

	//TODO: move connStr to .env when change to use final variables so u dont have hardcoded all the important stuff
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &PostgreStore{
		db: db,
	}, nil

}

func (s *PostgreStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgreStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgreStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
func (s *PostgreStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgreStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgreStore) createAccountTable() error {
	query := `create table if not exists accounts (
		id serial primary key, 
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}
