package post

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Find(ctx context.Context, query string) ([]Post, error)
	FindOneByID(ctx context.Context, id uuid.UUID) (Post, error)
}
