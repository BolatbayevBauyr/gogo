package service

import (
	"context"
	"github.com/google/uuid"
	"log"
	"scrapper/internal/core/post"
	"strings"
)

type PostService struct {
	postRepository post.Repository
}

func NewPostService(repository post.Repository) post.Service {
	return &PostService{
		postRepository: repository,
	}
}

func (p PostService) Query(ctx context.Context, words string) ([]post.Post, error) {
	words = strings.Trim(words, " ")
	arr := strings.Split(words, "_")

	m := map[uuid.UUID]bool{}
	var all []post.Post
	for _, v := range arr {
		single, err := p.postRepository.Find(ctx, v)
		if err != nil {
			return nil, err
		}
		all = append(all, single...)
	}
	var result []post.Post
	for i := range all {
		if _, ok := m[all[i].ID]; !ok {
			m[all[i].ID] = true
			result = append(result, all[i])
		}
	}

	return result, nil
}

func (p PostService) GetOne(ctx context.Context, id string) (post.Post, error) {
	uid, err := uuid.Parse(id)
	log.Println(err)
	if err != nil {
		return post.Post{}, err
	}

	res, err := p.postRepository.FindOneByID(ctx, uid)
	log.Println(err)
	if err != nil {
		return post.Post{}, err
	}

	return res, nil
}
