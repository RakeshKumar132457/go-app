package services

import (
	"database/sql"
	"go-app/api/repositories"
	"go-app/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(id int64) (*models.User, error) {
	return s.userRepo.GetByID(id)
}
