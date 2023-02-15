package post

import (
	"context"
	api_response "crawler/internal/core/api-response"
)

type Service interface {
	Create(ctx context.Context, d api_response.Data) (Post, error)
}
