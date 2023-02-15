package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"scrapper/internal/core/post"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) post.Repository {
	return &PostRepository{
		db: db,
	}
}

func (p PostRepository) Find(ctx context.Context, filter string) ([]post.Post, error) {
	query := `select * from posts where title like '%` + filter + `%'`

	res, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var posts []post.Post

	for res.Next() {
		var pt post.Post

		if err := res.Scan(&pt.ID, &pt.URL, &pt.Title, &pt.Body, &pt.CreatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, pt)

	}
	return posts, nil
}

func (p PostRepository) FindOneByID(ctx context.Context, id uuid.UUID) (post.Post, error) {
	query := `SELECT * FROM posts where id = $1`

	row := p.db.QueryRow(ctx, query, id)
	var result post.Post

	if err := row.Scan(&result.ID, &result.URL, &result.Title, &result.Body, &result.CreatedAt); err != nil {
		return post.Post{}, err
	}

	return result, nil
}
