package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestSomething(t *testing.T) {
	ctx := context.Background()
	service := PostService{
		log:    log.New(os.Stdout, "SUKA : ", log.LstdFlags),
		client: &http.Client{},
	}

	res, err := service.collectBody(ctx, "https://www.zakon.kz/6384372-v-turkestanskoy-oblasti-otkrylas-tysyachnaya-po-schetu-shkola.html")
	if err != nil {
		t.Fatalf("err : %v", err)
	}

	t.Logf("res : %s", res)
}
