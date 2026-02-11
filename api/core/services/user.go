package services

import (
	"errors"

	models "tcc-test/api/core/models"
	utlis "tcc-test/api/utils"
)

type UserService struct {
	repo models.UserRepository
}

func NewUserService(repo models.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("user id is required")
	}
	return s.repo.GetOne(id)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("user id is required")
	}
	return s.repo.GetByUsername(username)
}

func (s *UserService) CreateUser(user *models.UserCreate) (*models.User, error) {
	if user == nil || user.Username == "" || user.Password == "" {
		return nil, errors.New("invalid user data")
	}
	hashedPassword, err := utlis.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	return s.repo.Create(*user)
}
