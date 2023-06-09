package storer

import (
	"database/sql"
	"fmt"

	"github.com/vicolby/meets/types"
)

type EventStorage interface {
	GetEvents() ([]*types.Event, error)
	GetEvent(id int) (*types.Event, error)
	AddEvent(event *types.Event) error
	DeleteEvent(id int) error
	UpdateEvent(event *types.Event) error
}

type EventPostgresStorage struct {
	db *sql.DB
}

func NewEventPostgresStorage(db *sql.DB) *EventPostgresStorage {
	return &EventPostgresStorage{db: db}
}

func (s *EventPostgresStorage) GetEvents() ([]*types.Event, error) {
	rows, err := s.db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	events := []*types.Event{}
	for rows.Next() {
		event, err := scanIntoEvent(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (s *EventPostgresStorage) GetEvent(id int) (*types.Event, error) {
	rows, err := s.db.Query("SELECT * FROM events WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoEvent(rows)
	}
	return nil, fmt.Errorf("event with id %d not found", id)
}

func (s *EventPostgresStorage) AddEvent(event *types.Event) error {
	_, err := s.db.Exec(
		"INSERT INTO events (name, description, owner_id) VALUES ($1, $2, $3)",
		event.Name,
		event.Description,
		event.OwnerID,
	)
	return err
}

func (s *EventPostgresStorage) DeleteEvent(id int) error {
	_, err := s.db.Exec("DELETE FROM events WHERE id = $1", id)
	return err
}

func (s *EventPostgresStorage) UpdateEvent(event *types.Event) error {
	_, err := s.db.Exec(
		"UPDATE events SET name = $1, description = $2, participants = &3 WHERE id = $4",
		event.Name,
		event.Description,
		event.Participants,
		event.ID,
	)
	return err
}

func scanIntoEvent(rows *sql.Rows) (*types.Event, error) {
	e := &types.Event{}

	if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.OwnerID, &e.Participants); err != nil {
		return nil, err
	}
	return e, nil
}
