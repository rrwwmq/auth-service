package service

import (
	"fmt"

	"github.com/rrwwmq/auth-service/internal/domain"
	"github.com/rrwwmq/auth-service/internal/repository/postgres"
)

type AuthService struct {
	userRepo *postgres.UserRepo
}

func NewAuthService(userRepo *postgres.UserRepo) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(email, password string) error {
	_, err := s.userRepo.GetByEmail(email)
	if err == nil {
		return fmt.Errorf("email alredy exists")
	}

	hash, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash the password")
	}

	return s.userRepo.Create(domain.User{
		Email:        email,
		PasswordHash: hash,
	})
}

func (s *AuthService) Login(email, password string) error {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to find the email")
	}

	if !CheckPasswordHash(password, user.PasswordHash) {
		fmt.Println("invalid password")
		return fmt.Errorf("invalid password")
	}

	return nil
}
