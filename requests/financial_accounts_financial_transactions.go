package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	"github.com/lithic-com/lithic-go/core/query"
)

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
	return query.Marshal(r)
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
