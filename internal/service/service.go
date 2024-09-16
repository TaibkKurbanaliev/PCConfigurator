package service

import "pcbuilder/internal/storage"

type Service struct {
	// Service fields here
}

func NewService(repo *storage.Repository) *Service {
	return &Service{}
}
