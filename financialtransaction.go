// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// FinancialAccountFinancialTransactionService contains methods and other services
// that help with interacting with the lithic API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewFinancialAccountFinancialTransactionService] method instead.
type FinancialAccountFinancialTransactionService struct {
	Options []option.RequestOption
}

// NewFinancialAccountFinancialTransactionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewFinancialAccountFinancialTransactionService(opts ...option.RequestOption) (r *FinancialAccountFinancialTransactionService) {
	r = &FinancialAccountFinancialTransactionService{}
	r.Options = opts
	return
}

// Get the financial transaction for the provided token.
func (r *FinancialAccountFinancialTransactionService) Get(ctx context.Context, financialAccountToken string, financialTransactionToken string, opts ...option.RequestOption) (res *FinancialTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("financial_accounts/%s/financial_transactions/%s", financialAccountToken, financialTransactionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the financial transactions for a given financial account.
func (r *FinancialAccountFinancialTransactionService) List(ctx context.Context, financialAccountToken string, query FinancialTransactionListParams, opts ...option.RequestOption) (res *shared.SinglePage[FinancialTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("financial_accounts/%s/financial_transactions", financialAccountToken)
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

// List the financial transactions for a given financial account.
func (r *FinancialAccountFinancialTransactionService) ListAutoPaging(ctx context.Context, financialAccountToken string, query FinancialTransactionListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[FinancialTransaction] {
	return shared.NewSinglePageAutoPager(r.List(ctx, financialAccountToken, query, opts...))
}

type FinancialTransactionListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Financial Transaction category to be returned.
	Category param.Field[FinancialTransactionListParamsCategory] `query:"category"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Financial Transaction result to be returned.
	Result param.Field[FinancialTransactionListParamsResult] `query:"result"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Financial Transaction status to be returned.
	Status param.Field[FinancialTransactionListParamsStatus] `query:"status"`
}

// URLQuery serializes [FinancialTransactionListParams]'s query parameters as
// `url.Values`.
func (r FinancialTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Financial Transaction category to be returned.
type FinancialTransactionListParamsCategory string

const (
	FinancialTransactionListParamsCategoryACH      FinancialTransactionListParamsCategory = "ACH"
	FinancialTransactionListParamsCategoryCard     FinancialTransactionListParamsCategory = "CARD"
	FinancialTransactionListParamsCategoryTransfer FinancialTransactionListParamsCategory = "TRANSFER"
)

// Financial Transaction result to be returned.
type FinancialTransactionListParamsResult string

const (
	FinancialTransactionListParamsResultApproved FinancialTransactionListParamsResult = "APPROVED"
	FinancialTransactionListParamsResultDeclined FinancialTransactionListParamsResult = "DECLINED"
)

// Financial Transaction status to be returned.
type FinancialTransactionListParamsStatus string

const (
	FinancialTransactionListParamsStatusDeclined FinancialTransactionListParamsStatus = "DECLINED"
	FinancialTransactionListParamsStatusExpired  FinancialTransactionListParamsStatus = "EXPIRED"
	FinancialTransactionListParamsStatusPending  FinancialTransactionListParamsStatus = "PENDING"
	FinancialTransactionListParamsStatusReturned FinancialTransactionListParamsStatus = "RETURNED"
	FinancialTransactionListParamsStatusSettled  FinancialTransactionListParamsStatus = "SETTLED"
	FinancialTransactionListParamsStatusVoided   FinancialTransactionListParamsStatus = "VOIDED"
)
