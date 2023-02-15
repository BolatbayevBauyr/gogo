package service

import (
	"context"
	api_response "crawler/internal/core/api-response"
	"crawler/internal/core/post"
	"log"
	"time"
)

type Worker struct {
	postService         post.Service
	apiResponseReceiver api_response.Receiver
	log                 *log.Logger
}

func NewWorker(postService post.Service, receiver api_response.Receiver, log *log.Logger) *Worker {
	return &Worker{
		postService:         postService,
		apiResponseReceiver: receiver,
		log:                 log,
	}
}

func (w *Worker) Start(ctx context.Context, tickDuration time.Duration) {
	ticker := time.NewTicker(tickDuration)

	for {
		select {
		case <-ticker.C:
			resp, err := w.apiResponseReceiver.Receive(ctx)
			if err != nil {
				w.log.Printf("worker : err : %v", err)
			}

			for _, v := range resp.DataList {
				p, err := w.postService.Create(ctx, v)
				if err != nil {
					w.log.Printf("worker : err : %v", err)
				}
				p.Body = ""
				w.log.Printf("worker : post : %v", p)
			}
		case <-ctx.Done():
			break
		}
	}

}
