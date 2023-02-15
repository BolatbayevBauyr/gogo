package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"scrapper/cmd/api/handlers/postgrp"
	"scrapper/internal/repository"
	"scrapper/internal/service"
)

func API(db *pgxpool.Pool) http.Handler {
	mux := chi.NewMux()

	repo := repository.NewPostRepository(db)
	serv := service.NewPostService(repo)
	postHandler := postgrp.NewHandler(serv)

	mux.Get("/post/query/{prompt}", postHandler.Get)
	mux.Get("/post/{id}", postHandler.GetOne)

	return mux
}
