package storer

import (
	"database/sql"
	"fmt"

	"github.com/vicolby/meets/types"
)

type UserStorage interface {
	GetUser(id int) (*types.User, error)
	AddUser(user *types.User) error
	DeleteUser(id int) error
	UpdateUser(user *types.User) error
}

type UserPostgresStorage struct {
	db *sql.DB
}

func NewUserPostgresStorage(db *sql.DB) *UserPostgresStorage {
	return &UserPostgresStorage{db: db}
}

func (s *UserPostgresStorage) GetUser(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoUser(rows)
	}
	return nil, fmt.Errorf("user with id %d not found", id)
}

func (s *UserPostgresStorage) AddUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (s *UserPostgresStorage) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (s *UserPostgresStorage) UpdateUser(user *types.User) error {
	_, err := s.db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
	return err
}

func scanIntoUser(rows *sql.Rows) (*types.User, error) {
	u := &types.User{}

	if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return u, nil
}
