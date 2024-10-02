package services

import (
	"database/sql"
	"go-app/api/repositories"
	"go-app/internal/models"
	"log"
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
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		log.Printf(("Error in UserService.GetUserByID: %v"), err)
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		log.Printf("Error in UserService.GetAllUsers: %v", err)
		return nil, err
	}
	return users, nil
}
