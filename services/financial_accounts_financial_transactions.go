package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type FinancialAccountFinancialTransactionService struct {
	Options []option.RequestOption
}

func NewFinancialAccountFinancialTransactionService(opts ...option.RequestOption) (r *FinancialAccountFinancialTransactionService) {
	r = &FinancialAccountFinancialTransactionService{}
	r.Options = opts
	return
}

// Get the financial transaction for the provided token.
func (r *FinancialAccountFinancialTransactionService) Get(ctx context.Context, financial_account_token string, financial_transaction_token string, opts ...option.RequestOption) (res *responses.FinancialTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("financial_accounts/%s/financial_transactions/%s", financial_account_token, financial_transaction_token)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the financial transactions for a given financial account.
func (r *FinancialAccountFinancialTransactionService) List(ctx context.Context, financial_account_token string, query requests.FinancialTransactionListParams, opts ...option.RequestOption) (res *responses.SinglePage[responses.FinancialTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("financial_accounts/%s/financial_transactions", financial_account_token)
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

// List the financial transactions for a given financial account.
func (r *FinancialAccountFinancialTransactionService) ListAutoPaging(ctx context.Context, financial_account_token string, query requests.FinancialTransactionListParams, opts ...option.RequestOption) *responses.SinglePageAutoPager[responses.FinancialTransaction] {
	return responses.NewSinglePageAutoPager(r.List(ctx, financial_account_token, query, opts...))
}
