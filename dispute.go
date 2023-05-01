package lithic

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
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

type DisputeService struct {
	Options []option.RequestOption
}

func NewDisputeService(opts ...option.RequestOption) (r *DisputeService) {
	r = &DisputeService{}
	r.Options = opts
	return
}

// Initiate a dispute.
func (r *DisputeService) New(ctx context.Context, body DisputeNewParams, opts ...option.RequestOption) (res *Dispute, err error) {
	opts = append(r.Options[:], opts...)
	path := "disputes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get dispute.
func (r *DisputeService) Get(ctx context.Context, dispute_token string, opts ...option.RequestOption) (res *Dispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s", dispute_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update dispute. Can only be modified if status is `NEW`.
func (r *DisputeService) Update(ctx context.Context, dispute_token string, body DisputeUpdateParams, opts ...option.RequestOption) (res *Dispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s", dispute_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List disputes.
func (r *DisputeService) List(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) (res *shared.CursorPage[Dispute], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "disputes"
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

// List disputes.
func (r *DisputeService) ListAutoPaging(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[Dispute] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Withdraw dispute.
func (r *DisputeService) Delete(ctx context.Context, dispute_token string, opts ...option.RequestOption) (res *Dispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s", dispute_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Soft delete evidence for a dispute. Evidence will not be reviewed or submitted
// by Lithic after it is withdrawn.
func (r *DisputeService) DeleteEvidence(ctx context.Context, dispute_token string, evidence_token string, opts ...option.RequestOption) (res *DisputeEvidence, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s/evidences/%s", dispute_token, evidence_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Use this endpoint to upload evidences for the dispute. It will return a URL to
// upload your documents to. The URL will expire in 30 minutes.
//
// Uploaded documents must either be a `jpg`, `png` or `pdf` file, and each must be
// less than 5 GiB.
func (r *DisputeService) InitiateEvidenceUpload(ctx context.Context, dispute_token string, opts ...option.RequestOption) (res *DisputeEvidence, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s/evidences", dispute_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List evidence metadata for a dispute.
func (r *DisputeService) ListEvidences(ctx context.Context, dispute_token string, query DisputeListEvidencesParams, opts ...option.RequestOption) (res *shared.CursorPage[DisputeEvidence], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("disputes/%s/evidences", dispute_token)
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

// List evidence metadata for a dispute.
func (r *DisputeService) ListEvidencesAutoPaging(ctx context.Context, dispute_token string, query DisputeListEvidencesParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[DisputeEvidence] {
	return shared.NewCursorPageAutoPager(r.ListEvidences(ctx, dispute_token, query, opts...))
}

// Get a dispute's evidence metadata.
func (r *DisputeService) GetEvidence(ctx context.Context, dispute_token string, evidence_token string, opts ...option.RequestOption) (res *DisputeEvidence, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s/evidences/%s", dispute_token, evidence_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *DisputeService) UploadEvidence(ctx context.Context, disputeToken string, file io.Reader, opts ...option.RequestOption) (err error) {
	payload, err := r.InitiateEvidenceUpload(ctx, disputeToken, opts...)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	name := "anonymous_file"
	if nameable, ok := file.(interface{ Name() string }); ok {
		name = nameable.Name()
	}
	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		return err
	}
	io.Copy(part, file)

	req, err := http.NewRequestWithContext(ctx, "PUT", payload.UploadURL, body)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return

}

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
	DisputeResolutionReasonNoDisputeRights3DS            DisputeResolutionReason = "NO_DISPUTE_RIGHTS_3DS"
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
	// Timestamp of when dispute evidence was created.
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
	return apijson.MarshalRoot(r)
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
	return apijson.MarshalRoot(r)
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
	return apiquery.Marshal(r)
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
	return apiquery.Marshal(r)
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
