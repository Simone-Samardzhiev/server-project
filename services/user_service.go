package services

import (
	"context"
	"net/http"
	"server/auth"
	"server/models"
	"server/repositories"
	"server/utils"
)

// UserService interface will manage users business logic.
type UserService interface {
	// AddUser will check the email, encrypt the password and add the user.
	AddUser(ctx context.Context, user *models.UserPayload) *utils.ErrorResponse
}

// DefaultUserService struct is the default implementation of [UserService].
type DefaultUserService struct {
	userRepository repositories.UserRepository
}

func (s *DefaultUserService) AddUser(ctx context.Context, user *models.UserPayload) *utils.ErrorResponse {
	result, err := s.userRepository.CheckIfEmailExists(ctx, user.Email)
	if err != nil {
		return utils.InternalServerError()
	}

	if result {
		return utils.NewErrorResponse(
			"The email is already taken",
			http.StatusConflict,
		)
	}

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return utils.InternalServerError()
	}

	err = s.userRepository.AddUser(ctx, user)
	if err != nil {
		return utils.InternalServerError()
	}

	return nil
}

func NewDefaultUserService(userRepository repositories.UserRepository) *DefaultUserService {
	return &DefaultUserService{
		userRepository: userRepository,
	}
}
