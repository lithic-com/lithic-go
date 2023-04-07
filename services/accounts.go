package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type AccountService struct {
	Options []option.RequestOption
}

func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Get account configuration such as spend limits.
func (r *AccountService) Get(ctx context.Context, account_token string, opts ...option.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_token)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update account configuration such as spend limits and verification address. Can
// only be run on accounts that are part of the program managed by this API key.
//
// Accounts that are in the `PAUSED` state will not be able to transact or create
// new cards.
func (r *AccountService) Update(ctx context.Context, account_token string, body *requests.AccountUpdateParams, opts ...option.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_token)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List account configurations.
func (r *AccountService) List(ctx context.Context, query *requests.AccountListParams, opts ...option.RequestOption) (res *responses.Page[responses.Account], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return
	}
	err = cfg.Execute()
	if err != nil {
		return
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List account configurations.
func (r *AccountService) ListAutoPager(ctx context.Context, query *requests.AccountListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Account] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
