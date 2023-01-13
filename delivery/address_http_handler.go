package delivery

import (
	"net/http"
	"time"

	"github.com/w-woong/common/logger"
	"github.com/w-woong/partner/port"
)

type AddressHttpHandler struct {
	timeout time.Duration
	usc     port.AddressFinder
}

func NewAddressHttpHandler(timeout time.Duration, usc port.AddressFinder) *AddressHttpHandler {
	return &AddressHttpHandler{
		timeout: timeout,
		usc:     usc,
	}
}

func (d *AddressHttpHandler) HandleFindAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := d.usc.FindAddress(ctx, w); err != nil {
		logger.Error(err.Error())
		return
	}
}
