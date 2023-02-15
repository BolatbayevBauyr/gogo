package post

import "context"

type Repository interface {
	FetchOne(ctx context.Context, url string) (Post, error)
	Create(ctx context.Context, p Post) error
}
