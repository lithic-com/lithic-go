// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// DisputeService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDisputeService] method instead.
type DisputeService struct {
	Options []option.RequestOption
}

// NewDisputeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
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
func (r *DisputeService) Get(ctx context.Context, disputeToken string, opts ...option.RequestOption) (res *Dispute, err error) {
	opts = append(r.Options[:], opts...)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("disputes/%s", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update dispute. Can only be modified if status is `NEW`.
func (r *DisputeService) Update(ctx context.Context, disputeToken string, body DisputeUpdateParams, opts ...option.RequestOption) (res *Dispute, err error) {
	opts = append(r.Options[:], opts...)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("disputes/%s", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List disputes.
func (r *DisputeService) List(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Dispute], err error) {
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
func (r *DisputeService) ListAutoPaging(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Dispute] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Withdraw dispute.
func (r *DisputeService) Delete(ctx context.Context, disputeToken string, opts ...option.RequestOption) (res *Dispute, err error) {
	opts = append(r.Options[:], opts...)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("disputes/%s", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Soft delete evidence for a dispute. Evidence will not be reviewed or submitted
// by Lithic after it is withdrawn.
func (r *DisputeService) DeleteEvidence(ctx context.Context, disputeToken string, evidenceToken string, opts ...option.RequestOption) (res *DisputeEvidence, err error) {
	opts = append(r.Options[:], opts...)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	if evidenceToken == "" {
		err = errors.New("missing required evidence_token parameter")
		return
	}
	path := fmt.Sprintf("disputes/%s/evidences/%s", disputeToken, evidenceToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Use this endpoint to upload evidences for the dispute. It will return a URL to
// upload your documents to. The URL will expire in 30 minutes.
//
// Uploaded documents must either be a `jpg`, `png` or `pdf` file, and each must be
// less than 5 GiB.
func (r *DisputeService) InitiateEvidenceUpload(ctx context.Context, disputeToken string, body DisputeInitiateEvidenceUploadParams, opts ...option.RequestOption) (res *DisputeEvidence, err error) {
	opts = append(r.Options[:], opts...)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("disputes/%s/evidences", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List evidence metadata for a dispute.
func (r *DisputeService) ListEvidences(ctx context.Context, disputeToken string, query DisputeListEvidencesParams, opts ...option.RequestOption) (res *pagination.CursorPage[DisputeEvidence], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("disputes/%s/evidences", disputeToken)
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
func (r *DisputeService) ListEvidencesAutoPaging(ctx context.Context, disputeToken string, query DisputeListEvidencesParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[DisputeEvidence] {
	return pagination.NewCursorPageAutoPager(r.ListEvidences(ctx, disputeToken, query, opts...))
}

// Get a dispute's evidence metadata.
func (r *DisputeService) GetEvidence(ctx context.Context, disputeToken string, evidenceToken string, opts ...option.RequestOption) (res *DisputeEvidence, err error) {
	opts = append(r.Options[:], opts...)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	if evidenceToken == "" {
		err = errors.New("missing required evidence_token parameter")
		return
	}
	path := fmt.Sprintf("disputes/%s/evidences/%s", disputeToken, evidenceToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *DisputeService) UploadEvidence(ctx context.Context, disputeToken string, file io.Reader, opts ...option.RequestOption) (err error) {
	payload, err := r.InitiateEvidenceUpload(ctx, disputeToken, DisputeInitiateEvidenceUploadParams{}, opts...)
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

// Dispute.
type Dispute struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount under dispute. May be different from the original transaction amount.
	Amount int64 `json:"amount,required"`
	// Date dispute entered arbitration.
	ArbitrationDate time.Time `json:"arbitration_date,required,nullable" format:"date-time"`
	// Timestamp of when first Dispute was reported.
	Created time.Time `json:"created,required" format:"date-time"`
	// Date that the dispute was filed by the customer making the dispute.
	CustomerFiledDate time.Time `json:"customer_filed_date,required,nullable" format:"date-time"`
	// End customer description of the reason for the dispute.
	CustomerNote string `json:"customer_note,required,nullable"`
	// Unique identifiers for the dispute from the network.
	NetworkClaimIDs []string `json:"network_claim_ids,required,nullable"`
	// Date that the dispute was submitted to the network.
	NetworkFiledDate time.Time `json:"network_filed_date,required,nullable" format:"date-time"`
	// Network reason code used to file the dispute.
	NetworkReasonCode string `json:"network_reason_code,required,nullable"`
	// Date dispute entered pre-arbitration.
	PrearbitrationDate time.Time `json:"prearbitration_date,required,nullable" format:"date-time"`
	// Unique identifier for the dispute from the network. If there are multiple, this
	// will be the first claim id set by the network
	PrimaryClaimID string `json:"primary_claim_id,required,nullable"`
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
	RepresentmentDate time.Time `json:"representment_date,required,nullable" format:"date-time"`
	// Resolution amount net of network fees.
	ResolutionAmount int64 `json:"resolution_amount,required,nullable"`
	// Date that the dispute was resolved.
	ResolutionDate time.Time `json:"resolution_date,required,nullable" format:"date-time"`
	// Note by Dispute team on the case resolution.
	ResolutionNote string `json:"resolution_note,required,nullable"`
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
	ResolutionReason DisputeResolutionReason `json:"resolution_reason,required,nullable"`
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
	// The transaction that is being disputed. A transaction can only be disputed once
	// but may have multiple dispute cases.
	TransactionToken string      `json:"transaction_token,required" format:"uuid"`
	JSON             disputeJSON `json:"-"`
}

// disputeJSON contains the JSON metadata for the struct [Dispute]
type disputeJSON struct {
	Token              apijson.Field
	Amount             apijson.Field
	ArbitrationDate    apijson.Field
	Created            apijson.Field
	CustomerFiledDate  apijson.Field
	CustomerNote       apijson.Field
	NetworkClaimIDs    apijson.Field
	NetworkFiledDate   apijson.Field
	NetworkReasonCode  apijson.Field
	PrearbitrationDate apijson.Field
	PrimaryClaimID     apijson.Field
	Reason             apijson.Field
	RepresentmentDate  apijson.Field
	ResolutionAmount   apijson.Field
	ResolutionDate     apijson.Field
	ResolutionNote     apijson.Field
	ResolutionReason   apijson.Field
	Status             apijson.Field
	TransactionToken   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Dispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeJSON) RawJSON() string {
	return r.raw
}

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
	DisputeReasonRecurringTransactionNotCancelled DisputeReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	DisputeReasonRefundNotProcessed               DisputeReason = "REFUND_NOT_PROCESSED"
)

func (r DisputeReason) IsKnown() bool {
	switch r {
	case DisputeReasonAtmCashMisdispense, DisputeReasonCancelled, DisputeReasonDuplicated, DisputeReasonFraudCardNotPresent, DisputeReasonFraudCardPresent, DisputeReasonFraudOther, DisputeReasonGoodsServicesNotAsDescribed, DisputeReasonGoodsServicesNotReceived, DisputeReasonIncorrectAmount, DisputeReasonMissingAuth, DisputeReasonOther, DisputeReasonProcessingError, DisputeReasonRecurringTransactionNotCancelled, DisputeReasonRefundNotProcessed:
		return true
	}
	return false
}

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

func (r DisputeResolutionReason) IsKnown() bool {
	switch r {
	case DisputeResolutionReasonCaseLost, DisputeResolutionReasonNetworkRejected, DisputeResolutionReasonNoDisputeRights3DS, DisputeResolutionReasonNoDisputeRightsBelowThreshold, DisputeResolutionReasonNoDisputeRightsContactless, DisputeResolutionReasonNoDisputeRightsHybrid, DisputeResolutionReasonNoDisputeRightsMaxChargebacks, DisputeResolutionReasonNoDisputeRightsOther, DisputeResolutionReasonPastFilingDate, DisputeResolutionReasonPrearbitrationRejected, DisputeResolutionReasonProcessorRejectedOther, DisputeResolutionReasonRefunded, DisputeResolutionReasonRefundedAfterChargeback, DisputeResolutionReasonWithdrawn, DisputeResolutionReasonWonArbitration, DisputeResolutionReasonWonFirstChargeback, DisputeResolutionReasonWonPrearbitration:
		return true
	}
	return false
}

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
type DisputeStatus string

const (
	DisputeStatusArbitration     DisputeStatus = "ARBITRATION"
	DisputeStatusCaseClosed      DisputeStatus = "CASE_CLOSED"
	DisputeStatusCaseWon         DisputeStatus = "CASE_WON"
	DisputeStatusNew             DisputeStatus = "NEW"
	DisputeStatusPendingCustomer DisputeStatus = "PENDING_CUSTOMER"
	DisputeStatusPrearbitration  DisputeStatus = "PREARBITRATION"
	DisputeStatusRepresentment   DisputeStatus = "REPRESENTMENT"
	DisputeStatusSubmitted       DisputeStatus = "SUBMITTED"
)

func (r DisputeStatus) IsKnown() bool {
	switch r {
	case DisputeStatusArbitration, DisputeStatusCaseClosed, DisputeStatusCaseWon, DisputeStatusNew, DisputeStatusPendingCustomer, DisputeStatusPrearbitration, DisputeStatusRepresentment, DisputeStatusSubmitted:
		return true
	}
	return false
}

// Dispute evidence.
type DisputeEvidence struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Timestamp of when dispute evidence was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Dispute token evidence is attached to.
	DisputeToken string `json:"dispute_token,required" format:"uuid"`
	// Upload status types:
	//
	// - `DELETED` - Evidence was deleted.
	// - `ERROR` - Evidence upload failed.
	// - `PENDING` - Evidence is pending upload.
	// - `REJECTED` - Evidence was rejected.
	// - `UPLOADED` - Evidence was uploaded.
	UploadStatus DisputeEvidenceUploadStatus `json:"upload_status,required"`
	// URL to download evidence. Only shown when `upload_status` is `UPLOADED`.
	DownloadURL string `json:"download_url"`
	// File name of evidence. Recommended to give the dispute evidence a human-readable
	// identifier.
	Filename string `json:"filename"`
	// URL to upload evidence. Only shown when `upload_status` is `PENDING`.
	UploadURL string              `json:"upload_url"`
	JSON      disputeEvidenceJSON `json:"-"`
}

// disputeEvidenceJSON contains the JSON metadata for the struct [DisputeEvidence]
type disputeEvidenceJSON struct {
	Token        apijson.Field
	Created      apijson.Field
	DisputeToken apijson.Field
	UploadStatus apijson.Field
	DownloadURL  apijson.Field
	Filename     apijson.Field
	UploadURL    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DisputeEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeEvidenceJSON) RawJSON() string {
	return r.raw
}

// Upload status types:
//
// - `DELETED` - Evidence was deleted.
// - `ERROR` - Evidence upload failed.
// - `PENDING` - Evidence is pending upload.
// - `REJECTED` - Evidence was rejected.
// - `UPLOADED` - Evidence was uploaded.
type DisputeEvidenceUploadStatus string

const (
	DisputeEvidenceUploadStatusDeleted  DisputeEvidenceUploadStatus = "DELETED"
	DisputeEvidenceUploadStatusError    DisputeEvidenceUploadStatus = "ERROR"
	DisputeEvidenceUploadStatusPending  DisputeEvidenceUploadStatus = "PENDING"
	DisputeEvidenceUploadStatusRejected DisputeEvidenceUploadStatus = "REJECTED"
	DisputeEvidenceUploadStatusUploaded DisputeEvidenceUploadStatus = "UPLOADED"
)

func (r DisputeEvidenceUploadStatus) IsKnown() bool {
	switch r {
	case DisputeEvidenceUploadStatusDeleted, DisputeEvidenceUploadStatusError, DisputeEvidenceUploadStatusPending, DisputeEvidenceUploadStatusRejected, DisputeEvidenceUploadStatusUploaded:
		return true
	}
	return false
}

type DisputeNewParams struct {
	// Amount to dispute
	Amount param.Field[int64] `json:"amount,required"`
	// Reason for dispute
	Reason param.Field[DisputeNewParamsReason] `json:"reason,required"`
	// Transaction to dispute
	TransactionToken param.Field[string] `json:"transaction_token,required" format:"uuid"`
	// Date the customer filed the dispute
	CustomerFiledDate param.Field[time.Time] `json:"customer_filed_date" format:"date-time"`
	// Customer description of dispute
	CustomerNote param.Field[string] `json:"customer_note"`
}

func (r DisputeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Reason for dispute
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
	DisputeNewParamsReasonRecurringTransactionNotCancelled DisputeNewParamsReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	DisputeNewParamsReasonRefundNotProcessed               DisputeNewParamsReason = "REFUND_NOT_PROCESSED"
)

func (r DisputeNewParamsReason) IsKnown() bool {
	switch r {
	case DisputeNewParamsReasonAtmCashMisdispense, DisputeNewParamsReasonCancelled, DisputeNewParamsReasonDuplicated, DisputeNewParamsReasonFraudCardNotPresent, DisputeNewParamsReasonFraudCardPresent, DisputeNewParamsReasonFraudOther, DisputeNewParamsReasonGoodsServicesNotAsDescribed, DisputeNewParamsReasonGoodsServicesNotReceived, DisputeNewParamsReasonIncorrectAmount, DisputeNewParamsReasonMissingAuth, DisputeNewParamsReasonOther, DisputeNewParamsReasonProcessingError, DisputeNewParamsReasonRecurringTransactionNotCancelled, DisputeNewParamsReasonRefundNotProcessed:
		return true
	}
	return false
}

type DisputeUpdateParams struct {
	// Amount to dispute
	Amount param.Field[int64] `json:"amount"`
	// Date the customer filed the dispute
	CustomerFiledDate param.Field[time.Time] `json:"customer_filed_date" format:"date-time"`
	// Customer description of dispute
	CustomerNote param.Field[string] `json:"customer_note"`
	// Reason for dispute
	Reason param.Field[DisputeUpdateParamsReason] `json:"reason"`
}

func (r DisputeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Reason for dispute
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
	DisputeUpdateParamsReasonRecurringTransactionNotCancelled DisputeUpdateParamsReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	DisputeUpdateParamsReasonRefundNotProcessed               DisputeUpdateParamsReason = "REFUND_NOT_PROCESSED"
)

func (r DisputeUpdateParamsReason) IsKnown() bool {
	switch r {
	case DisputeUpdateParamsReasonAtmCashMisdispense, DisputeUpdateParamsReasonCancelled, DisputeUpdateParamsReasonDuplicated, DisputeUpdateParamsReasonFraudCardNotPresent, DisputeUpdateParamsReasonFraudCardPresent, DisputeUpdateParamsReasonFraudOther, DisputeUpdateParamsReasonGoodsServicesNotAsDescribed, DisputeUpdateParamsReasonGoodsServicesNotReceived, DisputeUpdateParamsReasonIncorrectAmount, DisputeUpdateParamsReasonMissingAuth, DisputeUpdateParamsReasonOther, DisputeUpdateParamsReasonProcessingError, DisputeUpdateParamsReasonRecurringTransactionNotCancelled, DisputeUpdateParamsReasonRefundNotProcessed:
		return true
	}
	return false
}

type DisputeListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// List disputes of a specific status.
	Status param.Field[DisputeListParamsStatus] `query:"status"`
	// Transaction tokens to filter by.
	TransactionTokens param.Field[[]string] `query:"transaction_tokens" format:"uuid"`
}

// URLQuery serializes [DisputeListParams]'s query parameters as `url.Values`.
func (r DisputeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// List disputes of a specific status.
type DisputeListParamsStatus string

const (
	DisputeListParamsStatusArbitration     DisputeListParamsStatus = "ARBITRATION"
	DisputeListParamsStatusCaseClosed      DisputeListParamsStatus = "CASE_CLOSED"
	DisputeListParamsStatusCaseWon         DisputeListParamsStatus = "CASE_WON"
	DisputeListParamsStatusNew             DisputeListParamsStatus = "NEW"
	DisputeListParamsStatusPendingCustomer DisputeListParamsStatus = "PENDING_CUSTOMER"
	DisputeListParamsStatusPrearbitration  DisputeListParamsStatus = "PREARBITRATION"
	DisputeListParamsStatusRepresentment   DisputeListParamsStatus = "REPRESENTMENT"
	DisputeListParamsStatusSubmitted       DisputeListParamsStatus = "SUBMITTED"
)

func (r DisputeListParamsStatus) IsKnown() bool {
	switch r {
	case DisputeListParamsStatusArbitration, DisputeListParamsStatusCaseClosed, DisputeListParamsStatusCaseWon, DisputeListParamsStatusNew, DisputeListParamsStatusPendingCustomer, DisputeListParamsStatusPrearbitration, DisputeListParamsStatusRepresentment, DisputeListParamsStatusSubmitted:
		return true
	}
	return false
}

type DisputeInitiateEvidenceUploadParams struct {
	// Filename of the evidence.
	Filename param.Field[string] `json:"filename"`
}

func (r DisputeInitiateEvidenceUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DisputeListEvidencesParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [DisputeListEvidencesParams]'s query parameters as
// `url.Values`.
func (r DisputeListEvidencesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DisputeUploadEvidenceParams struct {
}
