package api_response

import "context"

type Receiver interface {
	Receive(ctx context.Context) (APIResponse, error)
}
