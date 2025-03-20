package repositories

import (
	"context"
	"database/sql"
	"server/models"
)

type UserRepository interface {
	CheckIfEmailExists(ctx context.Context, email string) (bool, error)
	AddUser(ctx context.Context, user *models.UserPayload) error
}

type DefaultUserRepository struct {
	db *sql.DB
}

func (r *DefaultUserRepository) CheckIfEmailExists(ctx context.Context, email string) (bool, error) {
	row := r.db.QueryRowContext(
		ctx,
		` SELECT COUNT(*) FROM users
            WHERE email = $1
		`,
		email)

	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *DefaultUserRepository) AddUser(ctx context.Context, user *models.UserPayload) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO users(email, password)
		VALUES(?, ?)`,
		user.Email,
		user.Password,
	)

	return err
}

func NewDefaultUserRepository(db *sql.DB) *DefaultUserRepository {
	return &DefaultUserRepository{
		db: db,
	}
}
