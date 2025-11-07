package services

import (
	"fmt"


	"github.com/diegoQuinas/gopenca/repository"
	"github.com/diegoQuinas/gopenca/models"


)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo} 
}

func (s *UserService) CreateUser(u models.User) (*models.User, error) {
	u.Email = strings.ToLower(u.Email)
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email address")
	}
	exists, _ := s.repo.EmailExists(u.Email) 
	if exists {
		return nil, fmt.Errorf("email already in use")
	}

	return s.repo.Create(u)


}
