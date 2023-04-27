package services

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type TransferService struct {
	Options []option.RequestOption
}

func NewTransferService(opts ...option.RequestOption) (r *TransferService) {
	r = &TransferService{}
	r.Options = opts
	return
}

// Transfer funds between two financial accounts
func (r *TransferService) New(ctx context.Context, body requests.TransferNewParams, opts ...option.RequestOption) (res *responses.TransferCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "transfer"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}
