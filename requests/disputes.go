package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/fields"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type DisputeNewParams struct {
	// Amount to dispute
	Amount fields.Field[int64] `json:"amount,required"`
	// Date the customer filed the dispute
	CustomerFiledDate fields.Field[time.Time] `json:"customer_filed_date" format:"date-time"`
	// Reason for dispute
	Reason fields.Field[DisputeNewParamsReason] `json:"reason,required"`
	// Transaction to dispute
	TransactionToken fields.Field[string] `json:"transaction_token,required" format:"uuid"`
	// Customer description of dispute
	CustomerNote fields.Field[string] `json:"customer_note"`
}

// MarshalJSON serializes DisputeNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DisputeNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DisputeNewParams) String() (result string) {
	return fmt.Sprintf("&DisputeNewParams{Amount:%s CustomerFiledDate:%s Reason:%s TransactionToken:%s CustomerNote:%s}", r.Amount, r.CustomerFiledDate, r.Reason, r.TransactionToken, r.CustomerNote)
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
	Amount fields.Field[int64] `json:"amount"`
	// Date the customer filed the dispute
	CustomerFiledDate fields.Field[time.Time] `json:"customer_filed_date" format:"date-time"`
	// Customer description of dispute
	CustomerNote fields.Field[string] `json:"customer_note"`
	// Reason for dispute
	Reason fields.Field[DisputeUpdateParamsReason] `json:"reason"`
}

// MarshalJSON serializes DisputeUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DisputeUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DisputeUpdateParams) String() (result string) {
	return fmt.Sprintf("&DisputeUpdateParams{Amount:%s CustomerFiledDate:%s CustomerNote:%s Reason:%s}", r.Amount, r.CustomerFiledDate, r.CustomerNote, r.Reason)
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
	TransactionToken fields.Field[string] `query:"transaction_token" format:"uuid"`
	// List disputes of a specific status.
	Status fields.Field[DisputeListParamsStatus] `query:"status"`
	// Page size (for pagination).
	PageSize fields.Field[int64] `query:"page_size"`
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin fields.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End fields.Field[time.Time] `query:"end" format:"date-time"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter fields.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore fields.Field[string] `query:"ending_before"`
}

// URLQuery serializes DisputeListParams into a url.Values of the query parameters
// associated with this value
func (r *DisputeListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DisputeListParams) String() (result string) {
	return fmt.Sprintf("&DisputeListParams{TransactionToken:%s Status:%s PageSize:%s Begin:%s End:%s StartingAfter:%s EndingBefore:%s}", r.TransactionToken, r.Status, r.PageSize, r.Begin, r.End, r.StartingAfter, r.EndingBefore)
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

type DisputesDeleteEvidenceParams struct {
	// The {dispute_token} parameter of
	// /disputes/{dispute_token}/evidences/{evidence_token}
	DisputeToken string `path:"dispute_token" json:"-"`
}

type DisputeListEvidencesParams struct {
	// Page size (for pagination).
	PageSize fields.Field[int64] `query:"page_size"`
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin fields.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End fields.Field[time.Time] `query:"end" format:"date-time"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter fields.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore fields.Field[string] `query:"ending_before"`
}

// URLQuery serializes DisputeListEvidencesParams into a url.Values of the query
// parameters associated with this value
func (r *DisputeListEvidencesParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DisputeListEvidencesParams) String() (result string) {
	return fmt.Sprintf("&DisputeListEvidencesParams{PageSize:%s Begin:%s End:%s StartingAfter:%s EndingBefore:%s}", r.PageSize, r.Begin, r.End, r.StartingAfter, r.EndingBefore)
}

type DisputesGetEvidenceParams struct {
	// The {dispute_token} parameter of
	// /disputes/{dispute_token}/evidences/{evidence_token}
	DisputeToken string `path:"dispute_token" json:"-"`
}
