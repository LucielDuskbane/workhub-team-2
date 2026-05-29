package services

import (
	"errors"
	"workhub/internal/dto"
	"workhub/internal/models"
	"workhub/internal/repositories"
	"workhub/internal/repositories/interfaces"
	"workhub/internal/utils"
)

type AuthService struct {
	userRepo interfaces.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repositories.NewUserRepository(),
	}
}

func (s *AuthService) Register(req dto.RegisterRequest) error {
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser.ID != 0 {
		return errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// if req.Role == "admin" {
	// 	return errors.New("cannot register as admin")
	// }

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}
	return s.userRepo.CreateUser(&user)
}

func (s *AuthService) Login(req dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
