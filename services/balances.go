package services

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type BalanceService struct {
	Options []option.RequestOption
}

func NewBalanceService(opts ...option.RequestOption) (r *BalanceService) {
	r = &BalanceService{}
	r.Options = opts
	return
}

// Get the balances for a program or a given end-user account
func (r *BalanceService) List(ctx context.Context, query *requests.BalanceListParams, opts ...option.RequestOption) (res *responses.SinglePage[responses.Balance], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "balances"
	cfg, err := option.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get the balances for a program or a given end-user account
func (r *BalanceService) ListAutoPager(ctx context.Context, query *requests.BalanceListParams, opts ...option.RequestOption) *responses.SinglePageAutoPager[responses.Balance] {
	return responses.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}
