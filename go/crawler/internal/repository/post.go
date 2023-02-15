package repository

import (
	"context"
	"crawler/internal/core/post"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PostRepository struct {
	db  *pgxpool.Pool
	log *log.Logger
}

func NewPostRepository(db *pgxpool.Pool, log *log.Logger) post.Repository {
	return &PostRepository{
		db:  db,
		log: log,
	}
}

func (p PostRepository) FetchOne(ctx context.Context, url string) (post.Post, error) {
	query := `select * from posts where url = $1`

	var res post.Post

	row := p.db.QueryRow(ctx, query, url)
	if err := row.Scan(&res.ID, &res.URL, &res.Title, &res.Body, &res.CreatedAt); err != nil {
		p.log.Printf("error in PostRepository.FetchOne() : %v", err)
		return post.Post{}, err
	}

	return res, nil
}

func (p PostRepository) Create(ctx context.Context, post post.Post) error {
	query := `insert into posts (url, title, body, created_at)
				values ($1, $2, $3, $4);`

	res, err := p.db.Exec(ctx, query, post.URL, post.Title, post.Body, post.CreatedAt)
	if err != nil {
		return err
	}

	p.log.Println(res.String())

	return nil
}
