package postgrp

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"scrapper/internal/core/post"
)

type Handler struct {
	postService post.Service
}

func NewHandler(service post.Service) *Handler {
	return &Handler{
		postService: service,
	}
}

func respond(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(data)
	w.WriteHeader(statusCode)
	w.Write(res)
}

func wrap(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	res := chi.URLParam(r, "prompt")
	if res == "" {
		respond(w, wrap(errors.New("no prompt")), http.StatusBadRequest)
		return
	}

	posts, err := h.postService.Query(r.Context(), res)
	if err != nil {
		respond(w, wrap(err), http.StatusBadRequest)
		return
	}

	respond(w, posts, http.StatusOK)
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	res := chi.URLParam(r, "id")
	if res == "" {
		respond(w, wrap(errors.New("invalid id")), http.StatusBadRequest)
		return
	}

	post, err := h.postService.GetOne(r.Context(), res)
	if err != nil {
		respond(w, wrap(err), http.StatusBadRequest)
		return
	}

	respond(w, post, http.StatusOK)
}
