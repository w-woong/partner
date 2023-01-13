package port

import (
	"context"
	"net/http"
)

type AddressSvc interface {
	FindAddress(ctx context.Context, w http.ResponseWriter) error
}

type AddressFinder interface {
	FindAddress(ctx context.Context, w http.ResponseWriter) error
}
