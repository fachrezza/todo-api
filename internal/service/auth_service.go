package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/fachrezza/todo-api/config"
	"github.com/fachrezza/todo-api/internal/dto"
	"github.com/fachrezza/todo-api/internal/model"
	"github.com/fachrezza/todo-api/internal/repository"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) error {

	// cek email sudah ada atau belum
	user, _ := s.repo.GetByEmail(ctx, req.Email)

	if user != nil {
		return errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	newUser := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(ctx, &newUser)
}

func (s *AuthService) Login(
	ctx context.Context,
	req dto.LoginRequest,
) (string, error) {

	user, err := s.repo.GetByEmail(ctx, req.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := config.GenerateToken(user.ID.String())

	if err != nil {
		return "", err
	}

	return token, nil
}