package responses

import (
	"time"

	apijson "github.com/lithic-com/lithic-go/internal/json"
)

type Dispute struct {
	// Amount under dispute. May be different from the original transaction amount.
	Amount int64 `json:"amount,required"`
	// Date dispute entered arbitration.
	ArbitrationDate time.Time `json:"arbitration_date,required" format:"date-time"`
	// Timestamp of when first Dispute was reported.
	Created time.Time `json:"created,required" format:"date-time"`
	// Date that the dispute was filed by the customer making the dispute.
	CustomerFiledDate time.Time `json:"customer_filed_date,required" format:"date-time"`
	// End customer description of the reason for the dispute.
	CustomerNote string `json:"customer_note,required"`
	// Unique identifiers for the dispute from the network.
	NetworkClaimIDs []string `json:"network_claim_ids,required"`
	// Unique identifier for the dispute from the network. If there are multiple, this
	// will be the first claim id set by the network
	PrimaryClaimID string `json:"primary_claim_id,required"`
	// Date that the dispute was submitted to the network.
	NetworkFiledDate time.Time `json:"network_filed_date,required" format:"date-time"`
	// Network reason code used to file the dispute.
	NetworkReasonCode string `json:"network_reason_code,required"`
	// Date dispute entered pre-arbitration.
	PrearbitrationDate time.Time `json:"prearbitration_date,required" format:"date-time"`
	// Dispute reason:
	//
	//   - `ATM_CASH_MISDISPENSE`: ATM cash misdispense.
	//   - `CANCELLED`: Transaction was cancelled by the customer.
	//   - `DUPLICATED`: The transaction was a duplicate.
	//   - `FRAUD_CARD_NOT_PRESENT`: Fraudulent transaction, card not present.
	//   - `FRAUD_CARD_PRESENT`: Fraudulent transaction, card present.
	//   - `FRAUD_OTHER`: Fraudulent transaction, other types such as questionable
	//     merchant activity.
	//   - `GOODS_SERVICES_NOT_AS_DESCRIBED`: The goods or services were not as
	//     described.
	//   - `GOODS_SERVICES_NOT_RECEIVED`: The goods or services were not received.
	//   - `INCORRECT_AMOUNT`: The transaction amount was incorrect.
	//   - `MISSING_AUTH`: The transaction was missing authorization.
	//   - `OTHER`: Other reason.
	//   - `PROCESSING_ERROR`: Processing error.
	//   - `REFUND_NOT_PROCESSED`: The refund was not processed.
	//   - `RECURRING_TRANSACTION_NOT_CANCELLED`: The recurring transaction was not
	//     cancelled.
	Reason DisputeReason `json:"reason,required"`
	// Date the representment was received.
	RepresentmentDate time.Time `json:"representment_date,required" format:"date-time"`
	// Resolution amount net of network fees.
	ResolutionAmount int64 `json:"resolution_amount,required"`
	// Date that the dispute was resolved.
	ResolutionDate time.Time `json:"resolution_date,required" format:"date-time"`
	// Note by Dispute team on the case resolution.
	ResolutionNote string `json:"resolution_note,required"`
	// Reason for the dispute resolution:
	//
	// - `CASE_LOST`: This case was lost at final arbitration.
	// - `NETWORK_REJECTED`: Network rejected.
	// - `NO_DISPUTE_RIGHTS_3DS`: No dispute rights, 3DS.
	// - `NO_DISPUTE_RIGHTS_BELOW_THRESHOLD`: No dispute rights, below threshold.
	// - `NO_DISPUTE_RIGHTS_CONTACTLESS`: No dispute rights, contactless.
	// - `NO_DISPUTE_RIGHTS_HYBRID`: No dispute rights, hybrid.
	// - `NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS`: No dispute rights, max chargebacks.
	// - `NO_DISPUTE_RIGHTS_OTHER`: No dispute rights, other.
	// - `PAST_FILING_DATE`: Past filing date.
	// - `PREARBITRATION_REJECTED`: Prearbitration rejected.
	// - `PROCESSOR_REJECTED_OTHER`: Processor rejected, other.
	// - `REFUNDED`: Refunded.
	// - `REFUNDED_AFTER_CHARGEBACK`: Refunded after chargeback.
	// - `WITHDRAWN`: Withdrawn.
	// - `WON_ARBITRATION`: Won arbitration.
	// - `WON_FIRST_CHARGEBACK`: Won first chargeback.
	// - `WON_PREARBITRATION`: Won prearbitration.
	ResolutionReason DisputeResolutionReason `json:"resolution_reason,required"`
	// Status types:
	//
	//   - `NEW` - New dispute case is opened.
	//   - `PENDING_CUSTOMER` - Lithic is waiting for customer to provide more
	//     information.
	//   - `SUBMITTED` - Dispute is submitted to the card network.
	//   - `REPRESENTMENT` - Case has entered second presentment.
	//   - `PREARBITRATION` - Case has entered prearbitration.
	//   - `ARBITRATION` - Case has entered arbitration.
	//   - `CASE_WON` - Case was won and credit will be issued.
	//   - `CASE_CLOSED` - Case was lost or withdrawn.
	Status DisputeStatus `json:"status,required"`
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// The transaction that is being disputed. A transaction can only be disputed once
	// but may have multiple dispute cases.
	TransactionToken string `json:"transaction_token,required" format:"uuid"`
	JSON             DisputeJSON
}

type DisputeJSON struct {
	Amount             apijson.Metadata
	ArbitrationDate    apijson.Metadata
	Created            apijson.Metadata
	CustomerFiledDate  apijson.Metadata
	CustomerNote       apijson.Metadata
	NetworkClaimIDs    apijson.Metadata
	PrimaryClaimID     apijson.Metadata
	NetworkFiledDate   apijson.Metadata
	NetworkReasonCode  apijson.Metadata
	PrearbitrationDate apijson.Metadata
	Reason             apijson.Metadata
	RepresentmentDate  apijson.Metadata
	ResolutionAmount   apijson.Metadata
	ResolutionDate     apijson.Metadata
	ResolutionNote     apijson.Metadata
	ResolutionReason   apijson.Metadata
	Status             apijson.Metadata
	Token              apijson.Metadata
	TransactionToken   apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Dispute using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Dispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DisputeReason string

const (
	DisputeReasonAtmCashMisdispense               DisputeReason = "ATM_CASH_MISDISPENSE"
	DisputeReasonCancelled                        DisputeReason = "CANCELLED"
	DisputeReasonDuplicated                       DisputeReason = "DUPLICATED"
	DisputeReasonFraudCardNotPresent              DisputeReason = "FRAUD_CARD_NOT_PRESENT"
	DisputeReasonFraudCardPresent                 DisputeReason = "FRAUD_CARD_PRESENT"
	DisputeReasonFraudOther                       DisputeReason = "FRAUD_OTHER"
	DisputeReasonGoodsServicesNotAsDescribed      DisputeReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	DisputeReasonGoodsServicesNotReceived         DisputeReason = "GOODS_SERVICES_NOT_RECEIVED"
	DisputeReasonIncorrectAmount                  DisputeReason = "INCORRECT_AMOUNT"
	DisputeReasonMissingAuth                      DisputeReason = "MISSING_AUTH"
	DisputeReasonOther                            DisputeReason = "OTHER"
	DisputeReasonProcessingError                  DisputeReason = "PROCESSING_ERROR"
	DisputeReasonRefundNotProcessed               DisputeReason = "REFUND_NOT_PROCESSED"
	DisputeReasonRecurringTransactionNotCancelled DisputeReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
)

type DisputeResolutionReason string

const (
	DisputeResolutionReasonCaseLost                      DisputeResolutionReason = "CASE_LOST"
	DisputeResolutionReasonNetworkRejected               DisputeResolutionReason = "NETWORK_REJECTED"
	DisputeResolutionReasonNoDisputeRights_3Ds           DisputeResolutionReason = "NO_DISPUTE_RIGHTS_3DS"
	DisputeResolutionReasonNoDisputeRightsBelowThreshold DisputeResolutionReason = "NO_DISPUTE_RIGHTS_BELOW_THRESHOLD"
	DisputeResolutionReasonNoDisputeRightsContactless    DisputeResolutionReason = "NO_DISPUTE_RIGHTS_CONTACTLESS"
	DisputeResolutionReasonNoDisputeRightsHybrid         DisputeResolutionReason = "NO_DISPUTE_RIGHTS_HYBRID"
	DisputeResolutionReasonNoDisputeRightsMaxChargebacks DisputeResolutionReason = "NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS"
	DisputeResolutionReasonNoDisputeRightsOther          DisputeResolutionReason = "NO_DISPUTE_RIGHTS_OTHER"
	DisputeResolutionReasonPastFilingDate                DisputeResolutionReason = "PAST_FILING_DATE"
	DisputeResolutionReasonPrearbitrationRejected        DisputeResolutionReason = "PREARBITRATION_REJECTED"
	DisputeResolutionReasonProcessorRejectedOther        DisputeResolutionReason = "PROCESSOR_REJECTED_OTHER"
	DisputeResolutionReasonRefunded                      DisputeResolutionReason = "REFUNDED"
	DisputeResolutionReasonRefundedAfterChargeback       DisputeResolutionReason = "REFUNDED_AFTER_CHARGEBACK"
	DisputeResolutionReasonWithdrawn                     DisputeResolutionReason = "WITHDRAWN"
	DisputeResolutionReasonWonArbitration                DisputeResolutionReason = "WON_ARBITRATION"
	DisputeResolutionReasonWonFirstChargeback            DisputeResolutionReason = "WON_FIRST_CHARGEBACK"
	DisputeResolutionReasonWonPrearbitration             DisputeResolutionReason = "WON_PREARBITRATION"
)

type DisputeStatus string

const (
	DisputeStatusNew             DisputeStatus = "NEW"
	DisputeStatusPendingCustomer DisputeStatus = "PENDING_CUSTOMER"
	DisputeStatusSubmitted       DisputeStatus = "SUBMITTED"
	DisputeStatusRepresentment   DisputeStatus = "REPRESENTMENT"
	DisputeStatusPrearbitration  DisputeStatus = "PREARBITRATION"
	DisputeStatusArbitration     DisputeStatus = "ARBITRATION"
	DisputeStatusCaseWon         DisputeStatus = "CASE_WON"
	DisputeStatusCaseClosed      DisputeStatus = "CASE_CLOSED"
)

type DisputeEvidence struct {
	// Timestamp of when first Dispute was reported.
	Created time.Time `json:"created,required" format:"date-time"`
	// Dispute token evidence is attached to.
	DisputeToken string `json:"dispute_token,required" format:"uuid"`
	// URL to download evidence. Only shown when `upload_status` is `UPLOADED`.
	DownloadURL string `json:"download_url"`
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Upload status types:
	//
	// - `DELETED` - Evidence was deleted.
	// - `ERROR` - Evidence upload failed.
	// - `PENDING` - Evidence is pending upload.
	// - `REJECTED` - Evidence was rejected.
	// - `UPLOADED` - Evidence was uploaded.
	UploadStatus DisputeEvidenceUploadStatus `json:"upload_status,required"`
	// URL to upload evidence. Only shown when `upload_status` is `PENDING`.
	UploadURL string `json:"upload_url"`
	JSON      DisputeEvidenceJSON
}

type DisputeEvidenceJSON struct {
	Created      apijson.Metadata
	DisputeToken apijson.Metadata
	DownloadURL  apijson.Metadata
	Token        apijson.Metadata
	UploadStatus apijson.Metadata
	UploadURL    apijson.Metadata
	raw          string
	Extras       map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DisputeEvidence using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DisputeEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DisputeEvidenceUploadStatus string

const (
	DisputeEvidenceUploadStatusDeleted  DisputeEvidenceUploadStatus = "DELETED"
	DisputeEvidenceUploadStatusError    DisputeEvidenceUploadStatus = "ERROR"
	DisputeEvidenceUploadStatusPending  DisputeEvidenceUploadStatus = "PENDING"
	DisputeEvidenceUploadStatusRejected DisputeEvidenceUploadStatus = "REJECTED"
	DisputeEvidenceUploadStatusUploaded DisputeEvidenceUploadStatus = "UPLOADED"
)

type DisputeInitiateEvidenceUploadResponse struct {
	UploadURL string `json:"upload_url"`
	JSON      DisputeInitiateEvidenceUploadResponseJSON
}

type DisputeInitiateEvidenceUploadResponseJSON struct {
	UploadURL apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DisputeInitiateEvidenceUploadResponse using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DisputeInitiateEvidenceUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DisputeListResponse struct {
	Data []Dispute `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    DisputeListResponseJSON
}

type DisputeListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	raw     string
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DisputeListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DisputeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DisputeListEvidencesResponse struct {
	Data []DisputeEvidence `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    DisputeListEvidencesResponseJSON
}

type DisputeListEvidencesResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	raw     string
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DisputeListEvidencesResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DisputeListEvidencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
