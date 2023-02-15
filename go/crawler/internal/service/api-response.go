package service

import (
	"context"
	apir "crawler/internal/core/api-response"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type APIResponseReceiver struct {
	client http.Client
}

func NewAPIResponseReceiver() apir.Receiver {
	return APIResponseReceiver{
		client: http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       3 * time.Second,
		},
	}
}

func (s APIResponseReceiver) Receive(ctx context.Context) (apir.APIResponse, error) {
	resp, err := s.client.Get("https://www.zakon.kz/api/today-news/?pn=1&pSize=20")
	if err != nil {
		return apir.APIResponse{}, err
	}
	if resp.StatusCode != 200 {
		return apir.APIResponse{}, errors.New("APIResopnseReceiver,")
	}
	defer resp.Body.Close()

	var res apir.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return apir.APIResponse{}, err
	}

	return res, nil
}
