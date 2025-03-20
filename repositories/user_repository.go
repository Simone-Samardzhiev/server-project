package repositories

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"server/models"
)

// UserRepository interface will manage user data.
type UserRepository interface {
	// CheckIfEmailExists will check if the email is already in use.
	CheckIfEmailExists(ctx context.Context, email string) (bool, error)

	// AddUser will save user information.
	AddUser(ctx context.Context, user *models.UserPayload) error
}

// DefaultUserRepository struct is default implementation of [UserRepository].
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
		VALUES($1, $2)`,
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
