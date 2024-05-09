// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"fmt"
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

// CardBalanceService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardBalanceService] method instead.
type CardBalanceService struct {
	Options []option.RequestOption
}

// NewCardBalanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardBalanceService(opts ...option.RequestOption) (r *CardBalanceService) {
	r = &CardBalanceService{}
	r.Options = opts
	return
}

// Get the balances for a given card.
func (r *CardBalanceService) List(ctx context.Context, cardToken string, query CardBalanceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CardBalanceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("cards/%s/balances", cardToken)
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

// Get the balances for a given card.
func (r *CardBalanceService) ListAutoPaging(ctx context.Context, cardToken string, query CardBalanceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CardBalanceListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, cardToken, query, opts...))
}

// Balance of a Financial Account
type CardBalanceListResponse struct {
	// Globally unique identifier for the financial account that holds this balance.
	Token string `json:"token,required" format:"uuid"`
	// Funds available for spend in the currency's smallest unit (e.g., cents for USD)
	AvailableAmount int64 `json:"available_amount,required"`
	// Date and time for when the balance was first created.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the local currency of the balance.
	Currency string `json:"currency,required"`
	// Globally unique identifier for the last financial transaction event that
	// impacted this balance.
	LastTransactionEventToken string `json:"last_transaction_event_token,required" format:"uuid"`
	// Globally unique identifier for the last financial transaction that impacted this
	// balance.
	LastTransactionToken string `json:"last_transaction_token,required" format:"uuid"`
	// Funds not available for spend due to card authorizations or pending ACH release.
	// Shown in the currency's smallest unit (e.g., cents for USD).
	PendingAmount int64 `json:"pending_amount,required"`
	// The sum of available and pending balance in the currency's smallest unit (e.g.,
	// cents for USD).
	TotalAmount int64 `json:"total_amount,required"`
	// Type of financial account.
	Type CardBalanceListResponseType `json:"type,required"`
	// Date and time for when the balance was last updated.
	Updated time.Time                   `json:"updated,required" format:"date-time"`
	JSON    cardBalanceListResponseJSON `json:"-"`
}

// cardBalanceListResponseJSON contains the JSON metadata for the struct
// [CardBalanceListResponse]
type cardBalanceListResponseJSON struct {
	Token                     apijson.Field
	AvailableAmount           apijson.Field
	Created                   apijson.Field
	Currency                  apijson.Field
	LastTransactionEventToken apijson.Field
	LastTransactionToken      apijson.Field
	PendingAmount             apijson.Field
	TotalAmount               apijson.Field
	Type                      apijson.Field
	Updated                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CardBalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardBalanceListResponseJSON) RawJSON() string {
	return r.raw
}

// Type of financial account.
type CardBalanceListResponseType string

const (
	CardBalanceListResponseTypeIssuing   CardBalanceListResponseType = "ISSUING"
	CardBalanceListResponseTypeOperating CardBalanceListResponseType = "OPERATING"
	CardBalanceListResponseTypeReserve   CardBalanceListResponseType = "RESERVE"
)

func (r CardBalanceListResponseType) IsKnown() bool {
	switch r {
	case CardBalanceListResponseTypeIssuing, CardBalanceListResponseTypeOperating, CardBalanceListResponseTypeReserve:
		return true
	}
	return false
}

type CardBalanceListParams struct {
	// UTC date of the balance to retrieve. Defaults to latest available balance
	BalanceDate param.Field[time.Time] `query:"balance_date" format:"date-time"`
	// Balance after a given financial event occured. For example, passing the
	// event_token of a $5 CARD_CLEARING financial event will return a balance
	// decreased by $5
	LastTransactionEventToken param.Field[string] `query:"last_transaction_event_token" format:"uuid"`
}

// URLQuery serializes [CardBalanceListParams]'s query parameters as `url.Values`.
func (r CardBalanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
