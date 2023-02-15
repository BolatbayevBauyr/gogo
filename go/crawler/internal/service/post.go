package service

import (
	"context"
	api_response "crawler/internal/core/api-response"
	"crawler/internal/core/post"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"strings"
)

type PostService struct {
	postRepository post.Repository
	log            *log.Logger
	client         *http.Client
}

func NewPostService(postRepository post.Repository, log *log.Logger) post.Service {
	return &PostService{
		postRepository: postRepository,
		log:            log,
		client:         &http.Client{},
	}
}

func (p PostService) Create(ctx context.Context, d api_response.Data) (post.Post, error) {
	var res post.Post

	res.CreatedAt = d.PublishedDate
	res.Title = d.PageTitle
	url := fmt.Sprintf("https://zakon.kz/%s", d.Alias)
	_, err := p.postRepository.FetchOne(ctx, url)
	if err != nil && err != pgx.ErrNoRows {
		return post.Post{}, err
	}
	if err == nil {
		return post.Post{}, nil
	}
	res.URL = url

	body, err := p.collectBody(ctx, url)
	if err != nil {
		return post.Post{}, err
	}
	res.Body = body

	if err := p.postRepository.Create(ctx, res); err != nil {
		return post.Post{}, err
	}

	return res, nil

}

func (p PostService) collectBody(ctx context.Context, url string) (string, error) {
	response, err := p.client.Get(url)
	p.log.Println(response.StatusCode)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	var res []string
	doc.Find(".content p").Each(func(i int, s *goquery.Selection) {
		p.log.Println(s.Text())
		if s.Parent().HasClass("content") {
			res = append(res, s.Text())
		}
	})

	return strings.Join(res, " "), nil
}
