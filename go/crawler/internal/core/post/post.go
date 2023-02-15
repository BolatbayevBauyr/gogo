package post

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        uuid.UUID
	URL       string
	Title     string
	Body      string
	CreatedAt time.Time
}
