package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type FinancialAccountBalanceService struct {
	Options []option.RequestOption
}

func NewFinancialAccountBalanceService(opts ...option.RequestOption) (r *FinancialAccountBalanceService) {
	r = &FinancialAccountBalanceService{}
	r.Options = opts
	return
}

// Get the balances for a given financial account.
func (r *FinancialAccountBalanceService) List(ctx context.Context, financial_account_token string, query *requests.FinancialAccountBalanceListParams, opts ...option.RequestOption) (res *responses.SinglePage[responses.Balance], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("financial_accounts/%s/balances", financial_account_token)
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
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

// Get the balances for a given financial account.
func (r *FinancialAccountBalanceService) ListAutoPager(ctx context.Context, financial_account_token string, query *requests.FinancialAccountBalanceListParams, opts ...option.RequestOption) *responses.SinglePageAutoPager[responses.Balance] {
	return responses.NewSinglePageAutoPager(r.List(ctx, financial_account_token, query, opts...))
}
