package repositories

import (
	"context"
	"database/sql"
	"server/models"
)

// EventRepository will manage events data.
type EventRepository interface {
	// GetEvents will return all events.
	GetEvents(ctx context.Context) ([]models.Event, error)

	// CheckEventId will check if an event id exists. Returns true if the id exists.
	CheckEventId(ctx context.Context, eventId int) (bool, error)

	// CheckIfRegistrationExists will check if registration is already present.
	CheckIfRegistrationExists(ctx context.Context, userId, eventId int) (bool, error)

	// AddRegistration will add a new registration to an event
	AddRegistration(ctx context.Context, userId, eventId int) error
}

// DefaultEventRepository is the default implementation of [EventRepository].
type DefaultEventRepository struct {
	db *sql.DB
}

func (r *DefaultEventRepository) countEvents(ctx context.Context) (int, error) {
	row := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM events")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *DefaultEventRepository) GetEvents(ctx context.Context) ([]models.Event, error) {
	count, err := r.countEvents(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, title, date, address, image_url, description
		FROM events`,
	)

	if err != nil {
		return nil, err
	}

	result := make([]models.Event, 0, count)

	for rows.Next() {
		var event models.Event
		err = rows.Scan(&event.Id, &event.Title, &event.Date, &event.Address, &event.ImageURL, &event.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, event)
	}
	return result, nil
}

func (r *DefaultEventRepository) CheckEventId(ctx context.Context, eventId int) (bool, error) {
	row := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM events WHERE id = $1", eventId)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *DefaultEventRepository) CheckIfRegistrationExists(ctx context.Context, userId, eventId int) (bool, error) {
	row := r.db.QueryRowContext(
		ctx,
		`SELECT COUNT(*)
		FROM registrations
		WHERE user_id = $1 
  		AND event_id = $2`,
		userId,
		eventId)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *DefaultEventRepository) AddRegistration(ctx context.Context, userId, eventId int) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO registrations(user_id, event_id)
		VALUES ($1, $2)`,
		userId,
		eventId,
	)

	return err
}

// NewDefaultEventRepository will create new [DefaultEventRepository].
func NewDefaultEventRepository(db *sql.DB) *DefaultEventRepository {
	return &DefaultEventRepository{
		db: db,
	}
}
