package main

import (
	"context"
	"crawler/internal/repository"
	"crawler/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	dsn = "postgres://postgres:1952@localhost:5432/postsdb"
)

func main() {
	log := log.New(os.Stdout, "Spyder : ", log.LstdFlags|log.Lshortfile)

	if err := run(log); err != nil {
		log.Fatalf("fatal error : %v \n", err)
	}
}

func run(log *log.Logger) error {
	ctx := context.Background()
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(ctx); err != nil {
		return err
	}

	postRepo := repository.NewPostRepository(db, log)
	postService := service.NewPostService(postRepo, log)
	receiver := service.NewAPIResponseReceiver()
	worker := service.NewWorker(postService, receiver, log)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		worker.Start(ctx, time.Minute)
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case <-shutdown:
		cancel()
		log.Println("Service stopped working")
	}

	return nil
}
