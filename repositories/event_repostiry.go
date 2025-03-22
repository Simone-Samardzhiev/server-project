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

// NewDefaultEventRepository will create new [DefaultEventRepository].
func NewDefaultEventRepository(db *sql.DB) *DefaultEventRepository {
	return &DefaultEventRepository{
		db: db,
	}
}
