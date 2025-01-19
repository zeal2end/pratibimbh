package service

import (
	"github.com/zeal2end/pratibimbh/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *repository.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*repository.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) GetUserByName(name string) (*repository.User, error) {
	return s.repo.GetUserByName(name)
}
