package storer

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func CreateConnection() (*sql.DB, error) {
	connStr := "user=postgres password=12345 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *PostgresStorage) Init() error {
	err := s.CreateUserTable()
	if err != nil {
		return err
	}

	err = s.CreateEventTable()
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL
	);`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateEventTable() error {
	query := `CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		owner_id INT NOT NULL
	);`

	_, err := s.db.Exec(query)
	return err
}
