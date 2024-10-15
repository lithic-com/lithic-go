// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// BalanceService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options []option.RequestOption
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r *BalanceService) {
	r = &BalanceService{}
	r.Options = opts
	return
}

// Get the balances for a program or a given end-user account
func (r *BalanceService) List(ctx context.Context, query BalanceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Balance], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/balances"
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

// Get the balances for a program or a given end-user account
func (r *BalanceService) ListAutoPaging(ctx context.Context, query BalanceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Balance] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Balance
type Balance struct {
	// Funds available for spend in the currency's smallest unit (e.g., cents for USD)
	AvailableAmount int64 `json:"available_amount,required"`
	// Date and time for when the balance was first created.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the local currency of the balance.
	Currency string `json:"currency,required"`
	// Globally unique identifier for the financial account that holds this balance.
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Type of financial account.
	FinancialAccountType BalanceFinancialAccountType `json:"financial_account_type,required"`
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
	// Date and time for when the balance was last updated.
	Updated time.Time   `json:"updated,required" format:"date-time"`
	JSON    balanceJSON `json:"-"`
}

// balanceJSON contains the JSON metadata for the struct [Balance]
type balanceJSON struct {
	AvailableAmount           apijson.Field
	Created                   apijson.Field
	Currency                  apijson.Field
	FinancialAccountToken     apijson.Field
	FinancialAccountType      apijson.Field
	LastTransactionEventToken apijson.Field
	LastTransactionToken      apijson.Field
	PendingAmount             apijson.Field
	TotalAmount               apijson.Field
	Updated                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Balance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceJSON) RawJSON() string {
	return r.raw
}

// Type of financial account.
type BalanceFinancialAccountType string

const (
	BalanceFinancialAccountTypeIssuing   BalanceFinancialAccountType = "ISSUING"
	BalanceFinancialAccountTypeOperating BalanceFinancialAccountType = "OPERATING"
	BalanceFinancialAccountTypeReserve   BalanceFinancialAccountType = "RESERVE"
)

func (r BalanceFinancialAccountType) IsKnown() bool {
	switch r {
	case BalanceFinancialAccountTypeIssuing, BalanceFinancialAccountTypeOperating, BalanceFinancialAccountTypeReserve:
		return true
	}
	return false
}

type BalanceListParams struct {
	// List balances for all financial accounts of a given account_token.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// UTC date and time of the balances to retrieve. Defaults to latest available
	// balances
	BalanceDate param.Field[time.Time] `query:"balance_date" format:"date-time"`
	// List balances for a given Financial Account type.
	FinancialAccountType param.Field[BalanceListParamsFinancialAccountType] `query:"financial_account_type"`
}

// URLQuery serializes [BalanceListParams]'s query parameters as `url.Values`.
func (r BalanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// List balances for a given Financial Account type.
type BalanceListParamsFinancialAccountType string

const (
	BalanceListParamsFinancialAccountTypeIssuing   BalanceListParamsFinancialAccountType = "ISSUING"
	BalanceListParamsFinancialAccountTypeOperating BalanceListParamsFinancialAccountType = "OPERATING"
	BalanceListParamsFinancialAccountTypeReserve   BalanceListParamsFinancialAccountType = "RESERVE"
)

func (r BalanceListParamsFinancialAccountType) IsKnown() bool {
	switch r {
	case BalanceListParamsFinancialAccountTypeIssuing, BalanceListParamsFinancialAccountTypeOperating, BalanceListParamsFinancialAccountTypeReserve:
		return true
	}
	return false
}
