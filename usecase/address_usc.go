package usecase

import (
	"context"
	"net/http"

	"github.com/w-woong/partner/port"
)

type addressFinder struct {
	svc port.AddressSvc
}

func NewAddressFinder(svc port.AddressSvc) *addressFinder {
	return &addressFinder{
		svc: svc,
	}
}

func (u *addressFinder) FindAddress(ctx context.Context, w http.ResponseWriter) error {
	return u.svc.FindAddress(ctx, w)
}
