package services

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type AggregateBalanceService struct {
	Options []option.RequestOption
}

func NewAggregateBalanceService(opts ...option.RequestOption) (r *AggregateBalanceService) {
	r = &AggregateBalanceService{}
	r.Options = opts
	return
}

// Get the aggregated balance across all end-user accounts by financial account
// type
func (r *AggregateBalanceService) List(ctx context.Context, query *requests.AggregateBalanceListParams, opts ...option.RequestOption) (res *responses.SinglePage[responses.AggregateBalance], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "aggregate_balances"
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

// Get the aggregated balance across all end-user accounts by financial account
// type
func (r *AggregateBalanceService) ListAutoPager(ctx context.Context, query *requests.AggregateBalanceListParams, opts ...option.RequestOption) *responses.SinglePageAutoPager[responses.AggregateBalance] {
	return responses.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}
