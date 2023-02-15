package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"scrapper/cmd/api/handlers"
	"syscall"
)

var (
	dsn = "postgres://postgres:1952@localhost:5432/postsdb"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic("SUKA")
	}

	log.Sugar().Infof("Starting application on port : %s", ":8080")
	if err := run(log.Sugar()); err != nil {
		log.Sugar().Errorf("ERROR : %v", err)
	}
}

func run(log *zap.SugaredLogger) error {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}

	router := handlers.API(db)

	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT)

	serverErrors := make(chan error, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			serverErrors <- err
		}
	}()

	select {
	case <-shutdown:
		if err := srv.Shutdown(ctx); err != nil {
			log.Errorf("server shutdown error : %v", err)
			return err
		}

	case err := <-serverErrors:
		log.Errorf("server error : %v", err)
		return err
	}

	return nil
}
