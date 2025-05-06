package models

import (
	"fmt"
	"log"
	"rest_api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, user_id)
	VALUES (?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: %v", err)
		return fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return fmt.Errorf("failed to execute query: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error retrieving last insert ID: %v", err)
		return fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)

	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) Update() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.ID)

	stmt.Close()
	return err
}

func (event *Event) Delete() error {
	query := `DELETE FROM events where id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	return err
}
