package services

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type FinancialAccountService struct {
	Options               []option.RequestOption
	Balances              *FinancialAccountBalanceService
	FinancialTransactions *FinancialAccountFinancialTransactionService
}

func NewFinancialAccountService(opts ...option.RequestOption) (r *FinancialAccountService) {
	r = &FinancialAccountService{}
	r.Options = opts
	r.Balances = NewFinancialAccountBalanceService(opts...)
	r.FinancialTransactions = NewFinancialAccountFinancialTransactionService(opts...)
	return
}

// Retrieve information on your financial accounts including routing and account
// number.
func (r *FinancialAccountService) List(ctx context.Context, query requests.FinancialAccountListParams, opts ...option.RequestOption) (res *responses.SinglePage[responses.FinancialAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "financial_accounts"
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

// Retrieve information on your financial accounts including routing and account
// number.
func (r *FinancialAccountService) ListAutoPaging(ctx context.Context, query requests.FinancialAccountListParams, opts ...option.RequestOption) *responses.SinglePageAutoPager[responses.FinancialAccount] {
	return responses.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}
