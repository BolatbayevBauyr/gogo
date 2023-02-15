package service

import (
	"context"
	"fmt"
	"testing"
)

func TestAPIResponseReceiver_Receive(t *testing.T) {
	ctx := context.Background()
	rec := NewAPIResponseReceiver()

	res, err := rec.Receive(ctx)
	if err != nil {
		t.Fatalf("Test failed : err : %v", err)
	}

	fmt.Println(res)
}