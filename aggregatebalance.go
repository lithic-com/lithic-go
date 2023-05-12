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
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// AggregateBalanceService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAggregateBalanceService] method
// instead.
type AggregateBalanceService struct {
	Options []option.RequestOption
}

// NewAggregateBalanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAggregateBalanceService(opts ...option.RequestOption) (r *AggregateBalanceService) {
	r = &AggregateBalanceService{}
	r.Options = opts
	return
}

// Get the aggregated balance across all end-user accounts by financial account
// type
func (r *AggregateBalanceService) List(ctx context.Context, query AggregateBalanceListParams, opts ...option.RequestOption) (res *shared.SinglePage[AggregateBalance], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "aggregate_balances"
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

// Get the aggregated balance across all end-user accounts by financial account
// type
func (r *AggregateBalanceService) ListAutoPaging(ctx context.Context, query AggregateBalanceListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[AggregateBalance] {
	return shared.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Aggregate Balance across all end-user accounts
type AggregateBalance struct {
	// Type of financial account
	FinancialAccountType AggregateBalanceFinancialAccountType `json:"financial_account_type,required"`
	// 3-digit alphabetic ISO 4217 code for the local currency of the balance.
	Currency string `json:"currency,required"`
	// Funds available for spend in the currency's smallest unit (e.g., cents for USD)
	AvailableAmount int64 `json:"available_amount,required"`
	// Funds not available for spend due to card authorizations or pending ACH release.
	// Shown in the currency's smallest unit (e.g., cents for USD)
	PendingAmount int64 `json:"pending_amount,required"`
	// The sum of available and pending balance in the currency's smallest unit (e.g.,
	// cents for USD)
	TotalAmount int64 `json:"total_amount,required"`
	// Date and time for when the balance was first created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Date and time for when the balance was last updated.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Globally unique identifier for the last transaction that impacted this balance
	LastTransactionToken string `json:"last_transaction_token,required" format:"uuid"`
	// Globally unique identifier for the last transaction event that impacted this
	// balance
	LastTransactionEventToken string `json:"last_transaction_event_token,required" format:"uuid"`
	// Globally unique identifier for the financial account that had its balance
	// updated most recently
	LastFinancialAccountToken string `json:"last_financial_account_token,required" format:"uuid"`
	JSON                      aggregateBalanceJSON
}

// aggregateBalanceJSON contains the JSON metadata for the struct
// [AggregateBalance]
type aggregateBalanceJSON struct {
	FinancialAccountType      apijson.Field
	Currency                  apijson.Field
	AvailableAmount           apijson.Field
	PendingAmount             apijson.Field
	TotalAmount               apijson.Field
	Created                   apijson.Field
	Updated                   apijson.Field
	LastTransactionToken      apijson.Field
	LastTransactionEventToken apijson.Field
	LastFinancialAccountToken apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AggregateBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AggregateBalanceFinancialAccountType string

const (
	AggregateBalanceFinancialAccountTypeIssuing AggregateBalanceFinancialAccountType = "ISSUING"
	AggregateBalanceFinancialAccountTypeReserve AggregateBalanceFinancialAccountType = "RESERVE"
)

type AggregateBalanceListParams struct {
	// Get the aggregate balance for a given Financial Account type.
	FinancialAccountType param.Field[AggregateBalanceListParamsFinancialAccountType] `query:"financial_account_type"`
}

// URLQuery serializes [AggregateBalanceListParams]'s query parameters as
// `url.Values`.
func (r AggregateBalanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AggregateBalanceListParamsFinancialAccountType string

const (
	AggregateBalanceListParamsFinancialAccountTypeIssuing AggregateBalanceListParamsFinancialAccountType = "ISSUING"
	AggregateBalanceListParamsFinancialAccountTypeReserve AggregateBalanceListParamsFinancialAccountType = "RESERVE"
)

type AggregateBalanceListResponse struct {
	Data []AggregateBalance `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    aggregateBalanceListResponseJSON
}

// aggregateBalanceListResponseJSON contains the JSON metadata for the struct
// [AggregateBalanceListResponse]
type aggregateBalanceListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AggregateBalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
