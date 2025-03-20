package services

import (
	"context"
	"net/http"
	"server/models"
	"server/repositories"
	"server/utils"
)

type UserService interface {
	AddUser(ctx context.Context, user *models.UserPayload) *utils.ErrorResponse
}

type DefaultUserService struct {
	userRepository repositories.UserRepository
}

func (s *DefaultUserService) AddUser(ctx context.Context, user *models.UserPayload) *utils.ErrorResponse {
	result, err := s.userRepository.CheckIfEmailExists(ctx, user.Email)
	if err != nil {
		return utils.InternalServerError()
	}

	if !result {
		return utils.NewErrorResponse(
			"The email is already taken",
			http.StatusConflict,
		)
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
