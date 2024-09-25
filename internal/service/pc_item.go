package service

import (
	"context"
	"pcbuilder/internal/storage"
)

type PCItemService struct {
	repository storage.PCItem
}

func NewPCItemService(repository storage.PCItem) *PCItemService {
	return &PCItemService{
		repository: repository,
	}
}

func (s *PCItemService) Create(ctx context.Context, item interface{}) (string, error) {
	return s.repository.Create(ctx, item)
}
