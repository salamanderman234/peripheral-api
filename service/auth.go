package service

import (
	"context"

	"github.com/salamanderman234/peripheral-api/domain"
	"github.com/salamanderman234/peripheral-api/entity"
)

type authService struct {
	repository domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) domain.AuthService {
	return &authService{
		repository: repo,
	}
}

func (a *authService) Authenticate(ctx context.Context, cred entity.Credentials) (string, error) {
	return "", nil
}
func (a *authService) Register(ctx context.Context, user entity.User) (entity.User, bool, error) {
	return entity.User{}, false, nil
}