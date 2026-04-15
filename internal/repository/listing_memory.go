package repository

import "github.com/T-AKHMAD/marketplace-backend/internal/domain"

type ListingRepository struct {
	items  []domain.Listing
	nextID int64
}

func NewListingRepository() *ListingRepository {
	return &ListingRepository{
		items:  make([]domain.Listing, 0),
		nextID: 1,
	}
}

func (r *ListingRepository) Create(l domain.Listing) domain.Listing {
	l.ID = r.nextID
	r.nextID++

	r.items = append(r.items, l)

	return l
}

func (r *ListingRepository) List() []domain.Listing {
	return r.items
}

func (r *ListingRepository) GetByID(id int64) (domain.Listing, bool) {
	for _, item := range r.items {
		if item.ID == id {
			return item, true
		}
	}
	return domain.Listing{}, false
}
