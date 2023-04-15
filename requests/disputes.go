package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type DisputeNewParams struct {
	// Amount to dispute
	Amount field.Field[int64] `json:"amount,required"`
	// Date the customer filed the dispute
	CustomerFiledDate field.Field[time.Time] `json:"customer_filed_date" format:"date-time"`
	// Reason for dispute
	Reason field.Field[DisputeNewParamsReason] `json:"reason,required"`
	// Transaction to dispute
	TransactionToken field.Field[string] `json:"transaction_token,required" format:"uuid"`
	// Customer description of dispute
	CustomerNote field.Field[string] `json:"customer_note"`
}

// MarshalJSON serializes DisputeNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r DisputeNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type DisputeNewParamsReason string

const (
	DisputeNewParamsReasonAtmCashMisdispense               DisputeNewParamsReason = "ATM_CASH_MISDISPENSE"
	DisputeNewParamsReasonCancelled                        DisputeNewParamsReason = "CANCELLED"
	DisputeNewParamsReasonDuplicated                       DisputeNewParamsReason = "DUPLICATED"
	DisputeNewParamsReasonFraudCardNotPresent              DisputeNewParamsReason = "FRAUD_CARD_NOT_PRESENT"
	DisputeNewParamsReasonFraudCardPresent                 DisputeNewParamsReason = "FRAUD_CARD_PRESENT"
	DisputeNewParamsReasonFraudOther                       DisputeNewParamsReason = "FRAUD_OTHER"
	DisputeNewParamsReasonGoodsServicesNotAsDescribed      DisputeNewParamsReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	DisputeNewParamsReasonGoodsServicesNotReceived         DisputeNewParamsReason = "GOODS_SERVICES_NOT_RECEIVED"
	DisputeNewParamsReasonIncorrectAmount                  DisputeNewParamsReason = "INCORRECT_AMOUNT"
	DisputeNewParamsReasonMissingAuth                      DisputeNewParamsReason = "MISSING_AUTH"
	DisputeNewParamsReasonOther                            DisputeNewParamsReason = "OTHER"
	DisputeNewParamsReasonProcessingError                  DisputeNewParamsReason = "PROCESSING_ERROR"
	DisputeNewParamsReasonRefundNotProcessed               DisputeNewParamsReason = "REFUND_NOT_PROCESSED"
	DisputeNewParamsReasonRecurringTransactionNotCancelled DisputeNewParamsReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
)

type DisputeUpdateParams struct {
	// Amount to dispute
	Amount field.Field[int64] `json:"amount"`
	// Date the customer filed the dispute
	CustomerFiledDate field.Field[time.Time] `json:"customer_filed_date" format:"date-time"`
	// Customer description of dispute
	CustomerNote field.Field[string] `json:"customer_note"`
	// Reason for dispute
	Reason field.Field[DisputeUpdateParamsReason] `json:"reason"`
}

// MarshalJSON serializes DisputeUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r DisputeUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type DisputeUpdateParamsReason string

const (
	DisputeUpdateParamsReasonAtmCashMisdispense               DisputeUpdateParamsReason = "ATM_CASH_MISDISPENSE"
	DisputeUpdateParamsReasonCancelled                        DisputeUpdateParamsReason = "CANCELLED"
	DisputeUpdateParamsReasonDuplicated                       DisputeUpdateParamsReason = "DUPLICATED"
	DisputeUpdateParamsReasonFraudCardNotPresent              DisputeUpdateParamsReason = "FRAUD_CARD_NOT_PRESENT"
	DisputeUpdateParamsReasonFraudCardPresent                 DisputeUpdateParamsReason = "FRAUD_CARD_PRESENT"
	DisputeUpdateParamsReasonFraudOther                       DisputeUpdateParamsReason = "FRAUD_OTHER"
	DisputeUpdateParamsReasonGoodsServicesNotAsDescribed      DisputeUpdateParamsReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	DisputeUpdateParamsReasonGoodsServicesNotReceived         DisputeUpdateParamsReason = "GOODS_SERVICES_NOT_RECEIVED"
	DisputeUpdateParamsReasonIncorrectAmount                  DisputeUpdateParamsReason = "INCORRECT_AMOUNT"
	DisputeUpdateParamsReasonMissingAuth                      DisputeUpdateParamsReason = "MISSING_AUTH"
	DisputeUpdateParamsReasonOther                            DisputeUpdateParamsReason = "OTHER"
	DisputeUpdateParamsReasonProcessingError                  DisputeUpdateParamsReason = "PROCESSING_ERROR"
	DisputeUpdateParamsReasonRefundNotProcessed               DisputeUpdateParamsReason = "REFUND_NOT_PROCESSED"
	DisputeUpdateParamsReasonRecurringTransactionNotCancelled DisputeUpdateParamsReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
)

type DisputeListParams struct {
	// List disputes of a given transaction token.
	TransactionToken field.Field[string] `query:"transaction_token" format:"uuid"`
	// List disputes of a specific status.
	Status field.Field[DisputeListParamsStatus] `query:"status"`
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
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

// URLQuery serializes DisputeListParams into a url.Values of the query parameters
// associated with this value
func (r DisputeListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type DisputeListParamsStatus string

const (
	DisputeListParamsStatusNew             DisputeListParamsStatus = "NEW"
	DisputeListParamsStatusPendingCustomer DisputeListParamsStatus = "PENDING_CUSTOMER"
	DisputeListParamsStatusSubmitted       DisputeListParamsStatus = "SUBMITTED"
	DisputeListParamsStatusRepresentment   DisputeListParamsStatus = "REPRESENTMENT"
	DisputeListParamsStatusPrearbitration  DisputeListParamsStatus = "PREARBITRATION"
	DisputeListParamsStatusArbitration     DisputeListParamsStatus = "ARBITRATION"
	DisputeListParamsStatusCaseWon         DisputeListParamsStatus = "CASE_WON"
	DisputeListParamsStatusCaseClosed      DisputeListParamsStatus = "CASE_CLOSED"
)

type DisputeListEvidencesParams struct {
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
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

// URLQuery serializes DisputeListEvidencesParams into a url.Values of the query
// parameters associated with this value
func (r DisputeListEvidencesParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
