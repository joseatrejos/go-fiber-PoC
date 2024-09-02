package services

import (
	"go-fiber-PoC/data/models"
	"go-fiber-PoC/data/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.UserRepo.GetByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAll()
}
