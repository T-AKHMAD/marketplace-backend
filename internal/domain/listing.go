package domain

import "time"

type Listing struct {
	ID          int64
	Title       string
	Description string
	Price       int64
	CreatedAt   time.Time
}
