package service

import (
	"github.com/goIdioms/rest/pkg/models"
	"github.com/goIdioms/rest/pkg/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generateHashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err.Error())
	}
	return string(hash)
}
