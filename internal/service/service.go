package service

import (
	"context"
	"pcbuilder/internal/domain"
	"pcbuilder/internal/storage"
)

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
}

type PCList interface {
}

type PCItem interface {
}

type Service struct {
	Authorization
	PCItem
	PCList
}

func NewService(repo *storage.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
