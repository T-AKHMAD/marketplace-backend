package service

import (
	"context"
	"log"
	"time"

	"marketplace/internal/domain"
	"marketplace/internal/repository"
)

type ListingService struct {
	pgRepo  *repository.ListingPostgresRepository
}

func NewListingService(	pg *repository.ListingPostgresRepository) *ListingService {
	return &ListingService{
		pgRepo:  pg,
	}
}

func (s *ListingService) Create(ctx context.Context, title, description string, price int64) (domain.Listing, error) {
	l := domain.Listing{
		Title:       title,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
	}
	
	l, err := s.pgRepo.Create(ctx, l)
	if err != nil {
		return domain.Listing{}, err
	}
	return l, nil
}

func (s *ListingService) List(ctx context.Context) ([]domain.Listing, error) {
	list, err := s.pgRepo.List(ctx)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return list, nil
}

func (s *ListingService) GetByID(ctx context.Context, id int64) (domain.Listing, error) {
	l, err := s.pgRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Listing{}, err
	}
	return l, nil
}
