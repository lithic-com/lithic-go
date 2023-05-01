package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/field"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
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
func (r *FinancialAccountFinancialTransactionService) Get(ctx context.Context, financial_account_token string, financial_transaction_token string, opts ...option.RequestOption) (res *FinancialTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("financial_accounts/%s/financial_transactions/%s", financial_account_token, financial_transaction_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the financial transactions for a given financial account.
func (r *FinancialAccountFinancialTransactionService) List(ctx context.Context, financial_account_token string, query FinancialTransactionListParams, opts ...option.RequestOption) (res *shared.SinglePage[FinancialTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("financial_accounts/%s/financial_transactions", financial_account_token)
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
func (r *FinancialAccountFinancialTransactionService) ListAutoPaging(ctx context.Context, financial_account_token string, query FinancialTransactionListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[FinancialTransaction] {
	return shared.NewSinglePageAutoPager(r.List(ctx, financial_account_token, query, opts...))
}

type FinancialTransactionListParams struct {
	// Financial Transaction category to be returned.
	Category field.Field[FinancialTransactionListParamsCategory] `query:"category"`
	// Financial Transaction status to be returned.
	Status field.Field[FinancialTransactionListParamsStatus] `query:"status"`
	// Financial Transaction result to be returned.
	Result field.Field[FinancialTransactionListParamsResult] `query:"result"`
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter field.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore field.Field[string] `query:"ending_before"`
}

// URLQuery serializes FinancialTransactionListParams into a url.Values of the
// query parameters associated with this value
func (r FinancialTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type FinancialTransactionListParamsCategory string

const (
	FinancialTransactionListParamsCategoryACH      FinancialTransactionListParamsCategory = "ACH"
	FinancialTransactionListParamsCategoryCard     FinancialTransactionListParamsCategory = "CARD"
	FinancialTransactionListParamsCategoryTransfer FinancialTransactionListParamsCategory = "TRANSFER"
)

type FinancialTransactionListParamsStatus string

const (
	FinancialTransactionListParamsStatusDeclined FinancialTransactionListParamsStatus = "DECLINED"
	FinancialTransactionListParamsStatusExpired  FinancialTransactionListParamsStatus = "EXPIRED"
	FinancialTransactionListParamsStatusPending  FinancialTransactionListParamsStatus = "PENDING"
	FinancialTransactionListParamsStatusSettled  FinancialTransactionListParamsStatus = "SETTLED"
	FinancialTransactionListParamsStatusVoided   FinancialTransactionListParamsStatus = "VOIDED"
)

type FinancialTransactionListParamsResult string

const (
	FinancialTransactionListParamsResultApproved FinancialTransactionListParamsResult = "APPROVED"
	FinancialTransactionListParamsResultDeclined FinancialTransactionListParamsResult = "DECLINED"
)

type FinancialTransactionListResponse struct {
	Data []FinancialTransaction `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    FinancialTransactionListResponseJSON
}

type FinancialTransactionListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	raw     string
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// FinancialTransactionListResponse using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *FinancialTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
