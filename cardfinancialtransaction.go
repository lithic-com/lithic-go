// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// CardFinancialTransactionService contains methods and other services that help
// with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardFinancialTransactionService] method instead.
type CardFinancialTransactionService struct {
	Options []option.RequestOption
}

// NewCardFinancialTransactionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCardFinancialTransactionService(opts ...option.RequestOption) (r *CardFinancialTransactionService) {
	r = &CardFinancialTransactionService{}
	r.Options = opts
	return
}

// Get the card financial transaction for the provided token.
func (r *CardFinancialTransactionService) Get(ctx context.Context, cardToken string, financialTransactionToken string, opts ...option.RequestOption) (res *FinancialTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/financial_transactions/%s", cardToken, financialTransactionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the financial transactions for a given card.
func (r *CardFinancialTransactionService) List(ctx context.Context, cardToken string, query CardFinancialTransactionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[FinancialTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("cards/%s/financial_transactions", cardToken)
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

// List the financial transactions for a given card.
func (r *CardFinancialTransactionService) ListAutoPaging(ctx context.Context, cardToken string, query CardFinancialTransactionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FinancialTransaction] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, cardToken, query, opts...))
}

type CardFinancialTransactionListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Financial Transaction category to be returned.
	Category param.Field[CardFinancialTransactionListParamsCategory] `query:"category"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Financial Transaction result to be returned.
	Result param.Field[CardFinancialTransactionListParamsResult] `query:"result"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Financial Transaction status to be returned.
	Status param.Field[CardFinancialTransactionListParamsStatus] `query:"status"`
}

// URLQuery serializes [CardFinancialTransactionListParams]'s query parameters as
// `url.Values`.
func (r CardFinancialTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Financial Transaction category to be returned.
type CardFinancialTransactionListParamsCategory string

const (
	CardFinancialTransactionListParamsCategoryCard     CardFinancialTransactionListParamsCategory = "CARD"
	CardFinancialTransactionListParamsCategoryTransfer CardFinancialTransactionListParamsCategory = "TRANSFER"
)

func (r CardFinancialTransactionListParamsCategory) IsKnown() bool {
	switch r {
	case CardFinancialTransactionListParamsCategoryCard, CardFinancialTransactionListParamsCategoryTransfer:
		return true
	}
	return false
}

// Financial Transaction result to be returned.
type CardFinancialTransactionListParamsResult string

const (
	CardFinancialTransactionListParamsResultApproved CardFinancialTransactionListParamsResult = "APPROVED"
	CardFinancialTransactionListParamsResultDeclined CardFinancialTransactionListParamsResult = "DECLINED"
)

func (r CardFinancialTransactionListParamsResult) IsKnown() bool {
	switch r {
	case CardFinancialTransactionListParamsResultApproved, CardFinancialTransactionListParamsResultDeclined:
		return true
	}
	return false
}

// Financial Transaction status to be returned.
type CardFinancialTransactionListParamsStatus string

const (
	CardFinancialTransactionListParamsStatusDeclined CardFinancialTransactionListParamsStatus = "DECLINED"
	CardFinancialTransactionListParamsStatusExpired  CardFinancialTransactionListParamsStatus = "EXPIRED"
	CardFinancialTransactionListParamsStatusPending  CardFinancialTransactionListParamsStatus = "PENDING"
	CardFinancialTransactionListParamsStatusReturned CardFinancialTransactionListParamsStatus = "RETURNED"
	CardFinancialTransactionListParamsStatusSettled  CardFinancialTransactionListParamsStatus = "SETTLED"
	CardFinancialTransactionListParamsStatusVoided   CardFinancialTransactionListParamsStatus = "VOIDED"
)

func (r CardFinancialTransactionListParamsStatus) IsKnown() bool {
	switch r {
	case CardFinancialTransactionListParamsStatusDeclined, CardFinancialTransactionListParamsStatusExpired, CardFinancialTransactionListParamsStatusPending, CardFinancialTransactionListParamsStatusReturned, CardFinancialTransactionListParamsStatusSettled, CardFinancialTransactionListParamsStatusVoided:
		return true
	}
	return false
}
