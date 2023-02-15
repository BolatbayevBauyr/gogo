package post

import "context"

type Service interface {
	Query(ctx context.Context, words string) ([]Post, error)
	GetOne(ctx context.Context, id string) (Post, error)
}
