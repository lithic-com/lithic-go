// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// CardAggregateBalanceService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardAggregateBalanceService] method instead.
type CardAggregateBalanceService struct {
	Options []option.RequestOption
}

// NewCardAggregateBalanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCardAggregateBalanceService(opts ...option.RequestOption) (r *CardAggregateBalanceService) {
	r = &CardAggregateBalanceService{}
	r.Options = opts
	return
}

// Get the aggregated card balance across all end-user accounts.
func (r *CardAggregateBalanceService) List(ctx context.Context, query CardAggregateBalanceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CardAggregateBalanceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cards/aggregate_balances"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Get the aggregated card balance across all end-user accounts.
func (r *CardAggregateBalanceService) ListAutoPaging(ctx context.Context, query CardAggregateBalanceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CardAggregateBalanceListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Card Aggregate Balance across all end-user accounts
type CardAggregateBalanceListResponse struct {
	// Funds available for spend in the currency's smallest unit (e.g., cents for USD)
	AvailableAmount int64 `json:"available_amount,required"`
	// Date and time for when the balance was first created.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the local currency of the balance.
	Currency string `json:"currency,required"`
	// Globally unique identifier for the card that had its balance updated most
	// recently
	LastCardToken string `json:"last_card_token,required" format:"uuid"`
	// Globally unique identifier for the last transaction event that impacted this
	// balance
	LastTransactionEventToken string `json:"last_transaction_event_token,required" format:"uuid"`
	// Globally unique identifier for the last transaction that impacted this balance
	LastTransactionToken string `json:"last_transaction_token,required" format:"uuid"`
	// Funds not available for spend due to card authorizations or pending ACH release.
	// Shown in the currency's smallest unit (e.g., cents for USD)
	PendingAmount int64 `json:"pending_amount,required"`
	// The sum of available and pending balance in the currency's smallest unit (e.g.,
	// cents for USD)
	TotalAmount int64 `json:"total_amount,required"`
	// Date and time for when the balance was last updated.
	Updated time.Time                            `json:"updated,required" format:"date-time"`
	JSON    cardAggregateBalanceListResponseJSON `json:"-"`
}

// cardAggregateBalanceListResponseJSON contains the JSON metadata for the struct
// [CardAggregateBalanceListResponse]
type cardAggregateBalanceListResponseJSON struct {
	AvailableAmount           apijson.Field
	Created                   apijson.Field
	Currency                  apijson.Field
	LastCardToken             apijson.Field
	LastTransactionEventToken apijson.Field
	LastTransactionToken      apijson.Field
	PendingAmount             apijson.Field
	TotalAmount               apijson.Field
	Updated                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CardAggregateBalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAggregateBalanceListResponseJSON) RawJSON() string {
	return r.raw
}

type CardAggregateBalanceListParams struct {
	// Cardholder to retrieve aggregate balances for.
	AccountToken param.Field[string] `query:"account_token"`
	// Business to retrieve aggregate balances for.
	BusinessAccountToken param.Field[string] `query:"business_account_token"`
}

// URLQuery serializes [CardAggregateBalanceListParams]'s query parameters as
// `url.Values`.
func (r CardAggregateBalanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
