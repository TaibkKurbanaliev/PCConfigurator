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

func (s *PCItemService) GetByID(ctx context.Context, id string, itemType string) (interface{}, error) {
	return s.repository.GetByID(ctx, id, itemType)
}

func (s *PCItemService) GetAllItems(ctx context.Context, itemType string) ([]interface{}, error) {
	return s.repository.GetAllItems(ctx, itemType)
}

func (s *PCItemService) UpdateByID(ctx context.Context, id string, item interface{}) (int64, error) {
	return s.repository.UpdateByID(ctx, id, item)
}
