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
		fmt.Println("email already exists")
		return err
	}

	hash, err := HashPassword(password)
	if err != nil {
		fmt.Println("failed to hash the password")
		return err
	}

	return s.userRepo.Create(domain.User{
		Email:        email,
		PasswordHash: hash,
	})
}

func (s *AuthService) Login(email, password string) error {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		fmt.Println("failed to find the email")
		return err
	}

	if !CheckPasswordHash(password, user.PasswordHash) {
		fmt.Println("invalid password")
		return err
	}

	return nil
}
