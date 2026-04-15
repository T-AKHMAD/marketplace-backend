package repository

import (
	"context"
	"github.com/T-AKHMAD/marketplace-backend/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ListingPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewListingPostgresRepository(pool *pgxpool.Pool) *ListingPostgresRepository {
	return &ListingPostgresRepository{
		pool: pool,
	}
}

func (r *ListingPostgresRepository) Create(ctx context.Context, l domain.Listing) (domain.Listing, error) {
	err := r.pool.QueryRow(
		ctx,
		`INSERT INTO listings (title, description, price, created_at)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		l.Title,
		l.Description,
		l.Price,
		l.CreatedAt,
	).Scan(&l.ID)

	if err != nil {
		return domain.Listing{}, err
	}

	return l, nil
}

func (r *ListingPostgresRepository) List(ctx context.Context) ([]domain.Listing, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, description, price, created_at
		 FROM listings;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []domain.Listing
	for rows.Next() {
		var l domain.Listing

		err := rows.Scan(
			&l.ID,
			&l.Title,
			&l.Description,
			&l.Price,
			&l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listings, nil
}

func (r *ListingPostgresRepository) GetByID(ctx context.Context, id int64) (domain.Listing, error) {
	var l domain.Listing
	err := r.pool.QueryRow(
		ctx,
		`SELECT id, title, description, price, created_at
		FROM listings
		WHERE id = $1;`,
		id,
	).Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.Listing{}, domain.ErrNotFound
		}
		return domain.Listing{}, err
	}

	return l, nil
}
