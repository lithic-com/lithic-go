package services

import (
	"context"
	"fmt"

	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/pagination"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type AccountService struct {
	Options []options.RequestOption
}

func NewAccountService(opts ...options.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Get account configuration such as spend limits.
func (r *AccountService) Get(ctx context.Context, account_token string, opts ...options.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_token)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update account configuration such as spend limits and verification address. Can
// only be run on accounts that are part of the program managed by this API key.
//
// Accounts that are in the `PAUSED` state will not be able to transact or create
// new cards.
func (r *AccountService) Update(ctx context.Context, account_token string, body *requests.AccountUpdateParams, opts ...options.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_token)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List account configurations.
func (r *AccountService) List(ctx context.Context, query *requests.AccountListParams, opts ...options.RequestOption) (res *responses.AccountsPage, err error) {
	opts = append(r.Options, opts...)
	path := "accounts"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.AccountsPage{
		Page: &pagination.Page[responses.Account]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
