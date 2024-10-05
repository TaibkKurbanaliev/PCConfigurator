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
