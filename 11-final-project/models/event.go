package models

import (
	"proj/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	insert into events(name, description, location, dateTime, user_id)
	values(?,?,?,?,?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	q := "select * from events"
	rows, err := db.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	q := `
	select * from events
	where id = ?
	`
	row := db.DB.QueryRow(q, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	q := `
	update events 
	set name = ?, description =?, location = ?, dateTime = ?
	where id = ?
	`

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}

func (e Event) Delete() error {
	q := "delete from events where id = ?"

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err

}

func (e Event) Register(userId int64) error {
	q := "insert into registrations(event_id, user_id) values(?,?)"
	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	q := "delete from registrations where event_id = ? and user_id = ?"

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err
}
