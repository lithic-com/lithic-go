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
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
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
func (r *DisputeService) New(ctx context.Context, body DisputeNewParams, opts ...option.RequestOption) (res *DisputeNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/disputes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get dispute.
func (r *DisputeService) Get(ctx context.Context, disputeToken string, opts ...option.RequestOption) (res *DisputeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("v1/disputes/%s", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update dispute. Can only be modified if status is `NEW`.
func (r *DisputeService) Update(ctx context.Context, disputeToken string, body DisputeUpdateParams, opts ...option.RequestOption) (res *DisputeUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("v1/disputes/%s", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List disputes.
func (r *DisputeService) List(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) (res *pagination.CursorPage[DisputeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/disputes"
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
func (r *DisputeService) ListAutoPaging(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[DisputeListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Withdraw dispute.
func (r *DisputeService) Delete(ctx context.Context, disputeToken string, opts ...option.RequestOption) (res *DisputeDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("v1/disputes/%s", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Soft delete evidence for a dispute. Evidence will not be reviewed or submitted
// by Lithic after it is withdrawn.
func (r *DisputeService) DeleteEvidence(ctx context.Context, disputeToken string, evidenceToken string, opts ...option.RequestOption) (res *DisputeEvidence, err error) {
	opts = slices.Concat(r.Options, opts)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	if evidenceToken == "" {
		err = errors.New("missing required evidence_token parameter")
		return
	}
	path := fmt.Sprintf("v1/disputes/%s/evidences/%s", disputeToken, evidenceToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Use this endpoint to upload evidences for the dispute. It will return a URL to
// upload your documents to. The URL will expire in 30 minutes.
//
// Uploaded documents must either be a `jpg`, `png` or `pdf` file, and each must be
// less than 5 GiB.
func (r *DisputeService) InitiateEvidenceUpload(ctx context.Context, disputeToken string, body DisputeInitiateEvidenceUploadParams, opts ...option.RequestOption) (res *DisputeEvidence, err error) {
	opts = slices.Concat(r.Options, opts)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("v1/disputes/%s/evidences", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List evidence metadata for a dispute.
func (r *DisputeService) ListEvidences(ctx context.Context, disputeToken string, query DisputeListEvidencesParams, opts ...option.RequestOption) (res *pagination.CursorPage[DisputeEvidence], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("v1/disputes/%s/evidences", disputeToken)
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
	opts = slices.Concat(r.Options, opts)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	if evidenceToken == "" {
		err = errors.New("missing required evidence_token parameter")
		return
	}
	path := fmt.Sprintf("v1/disputes/%s/evidences/%s", disputeToken, evidenceToken)
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

// Dispute.
type DisputeNewResponse struct {
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
	Reason DisputeNewResponseReason `json:"reason,required"`
	// Date the representment was received.
	RepresentmentDate time.Time `json:"representment_date,required,nullable" format:"date-time"`
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
	ResolutionReason DisputeNewResponseResolutionReason `json:"resolution_reason,required,nullable"`
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
	Status DisputeNewResponseStatus `json:"status,required"`
	// The transaction that is being disputed. A transaction can only be disputed once
	// but may have multiple dispute cases.
	TransactionToken string                 `json:"transaction_token,required" format:"uuid"`
	JSON             disputeNewResponseJSON `json:"-"`
}

// disputeNewResponseJSON contains the JSON metadata for the struct
// [DisputeNewResponse]
type disputeNewResponseJSON struct {
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
	ResolutionDate     apijson.Field
	ResolutionNote     apijson.Field
	ResolutionReason   apijson.Field
	Status             apijson.Field
	TransactionToken   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DisputeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeNewResponseJSON) RawJSON() string {
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
type DisputeNewResponseReason string

const (
	DisputeNewResponseReasonAtmCashMisdispense               DisputeNewResponseReason = "ATM_CASH_MISDISPENSE"
	DisputeNewResponseReasonCancelled                        DisputeNewResponseReason = "CANCELLED"
	DisputeNewResponseReasonDuplicated                       DisputeNewResponseReason = "DUPLICATED"
	DisputeNewResponseReasonFraudCardNotPresent              DisputeNewResponseReason = "FRAUD_CARD_NOT_PRESENT"
	DisputeNewResponseReasonFraudCardPresent                 DisputeNewResponseReason = "FRAUD_CARD_PRESENT"
	DisputeNewResponseReasonFraudOther                       DisputeNewResponseReason = "FRAUD_OTHER"
	DisputeNewResponseReasonGoodsServicesNotAsDescribed      DisputeNewResponseReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	DisputeNewResponseReasonGoodsServicesNotReceived         DisputeNewResponseReason = "GOODS_SERVICES_NOT_RECEIVED"
	DisputeNewResponseReasonIncorrectAmount                  DisputeNewResponseReason = "INCORRECT_AMOUNT"
	DisputeNewResponseReasonMissingAuth                      DisputeNewResponseReason = "MISSING_AUTH"
	DisputeNewResponseReasonOther                            DisputeNewResponseReason = "OTHER"
	DisputeNewResponseReasonProcessingError                  DisputeNewResponseReason = "PROCESSING_ERROR"
	DisputeNewResponseReasonRecurringTransactionNotCancelled DisputeNewResponseReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	DisputeNewResponseReasonRefundNotProcessed               DisputeNewResponseReason = "REFUND_NOT_PROCESSED"
)

func (r DisputeNewResponseReason) IsKnown() bool {
	switch r {
	case DisputeNewResponseReasonAtmCashMisdispense, DisputeNewResponseReasonCancelled, DisputeNewResponseReasonDuplicated, DisputeNewResponseReasonFraudCardNotPresent, DisputeNewResponseReasonFraudCardPresent, DisputeNewResponseReasonFraudOther, DisputeNewResponseReasonGoodsServicesNotAsDescribed, DisputeNewResponseReasonGoodsServicesNotReceived, DisputeNewResponseReasonIncorrectAmount, DisputeNewResponseReasonMissingAuth, DisputeNewResponseReasonOther, DisputeNewResponseReasonProcessingError, DisputeNewResponseReasonRecurringTransactionNotCancelled, DisputeNewResponseReasonRefundNotProcessed:
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
type DisputeNewResponseResolutionReason string

const (
	DisputeNewResponseResolutionReasonCaseLost                      DisputeNewResponseResolutionReason = "CASE_LOST"
	DisputeNewResponseResolutionReasonNetworkRejected               DisputeNewResponseResolutionReason = "NETWORK_REJECTED"
	DisputeNewResponseResolutionReasonNoDisputeRights3DS            DisputeNewResponseResolutionReason = "NO_DISPUTE_RIGHTS_3DS"
	DisputeNewResponseResolutionReasonNoDisputeRightsBelowThreshold DisputeNewResponseResolutionReason = "NO_DISPUTE_RIGHTS_BELOW_THRESHOLD"
	DisputeNewResponseResolutionReasonNoDisputeRightsContactless    DisputeNewResponseResolutionReason = "NO_DISPUTE_RIGHTS_CONTACTLESS"
	DisputeNewResponseResolutionReasonNoDisputeRightsHybrid         DisputeNewResponseResolutionReason = "NO_DISPUTE_RIGHTS_HYBRID"
	DisputeNewResponseResolutionReasonNoDisputeRightsMaxChargebacks DisputeNewResponseResolutionReason = "NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS"
	DisputeNewResponseResolutionReasonNoDisputeRightsOther          DisputeNewResponseResolutionReason = "NO_DISPUTE_RIGHTS_OTHER"
	DisputeNewResponseResolutionReasonPastFilingDate                DisputeNewResponseResolutionReason = "PAST_FILING_DATE"
	DisputeNewResponseResolutionReasonPrearbitrationRejected        DisputeNewResponseResolutionReason = "PREARBITRATION_REJECTED"
	DisputeNewResponseResolutionReasonProcessorRejectedOther        DisputeNewResponseResolutionReason = "PROCESSOR_REJECTED_OTHER"
	DisputeNewResponseResolutionReasonRefunded                      DisputeNewResponseResolutionReason = "REFUNDED"
	DisputeNewResponseResolutionReasonRefundedAfterChargeback       DisputeNewResponseResolutionReason = "REFUNDED_AFTER_CHARGEBACK"
	DisputeNewResponseResolutionReasonWithdrawn                     DisputeNewResponseResolutionReason = "WITHDRAWN"
	DisputeNewResponseResolutionReasonWonArbitration                DisputeNewResponseResolutionReason = "WON_ARBITRATION"
	DisputeNewResponseResolutionReasonWonFirstChargeback            DisputeNewResponseResolutionReason = "WON_FIRST_CHARGEBACK"
	DisputeNewResponseResolutionReasonWonPrearbitration             DisputeNewResponseResolutionReason = "WON_PREARBITRATION"
)

func (r DisputeNewResponseResolutionReason) IsKnown() bool {
	switch r {
	case DisputeNewResponseResolutionReasonCaseLost, DisputeNewResponseResolutionReasonNetworkRejected, DisputeNewResponseResolutionReasonNoDisputeRights3DS, DisputeNewResponseResolutionReasonNoDisputeRightsBelowThreshold, DisputeNewResponseResolutionReasonNoDisputeRightsContactless, DisputeNewResponseResolutionReasonNoDisputeRightsHybrid, DisputeNewResponseResolutionReasonNoDisputeRightsMaxChargebacks, DisputeNewResponseResolutionReasonNoDisputeRightsOther, DisputeNewResponseResolutionReasonPastFilingDate, DisputeNewResponseResolutionReasonPrearbitrationRejected, DisputeNewResponseResolutionReasonProcessorRejectedOther, DisputeNewResponseResolutionReasonRefunded, DisputeNewResponseResolutionReasonRefundedAfterChargeback, DisputeNewResponseResolutionReasonWithdrawn, DisputeNewResponseResolutionReasonWonArbitration, DisputeNewResponseResolutionReasonWonFirstChargeback, DisputeNewResponseResolutionReasonWonPrearbitration:
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
type DisputeNewResponseStatus string

const (
	DisputeNewResponseStatusArbitration     DisputeNewResponseStatus = "ARBITRATION"
	DisputeNewResponseStatusCaseClosed      DisputeNewResponseStatus = "CASE_CLOSED"
	DisputeNewResponseStatusCaseWon         DisputeNewResponseStatus = "CASE_WON"
	DisputeNewResponseStatusNew             DisputeNewResponseStatus = "NEW"
	DisputeNewResponseStatusPendingCustomer DisputeNewResponseStatus = "PENDING_CUSTOMER"
	DisputeNewResponseStatusPrearbitration  DisputeNewResponseStatus = "PREARBITRATION"
	DisputeNewResponseStatusRepresentment   DisputeNewResponseStatus = "REPRESENTMENT"
	DisputeNewResponseStatusSubmitted       DisputeNewResponseStatus = "SUBMITTED"
)

func (r DisputeNewResponseStatus) IsKnown() bool {
	switch r {
	case DisputeNewResponseStatusArbitration, DisputeNewResponseStatusCaseClosed, DisputeNewResponseStatusCaseWon, DisputeNewResponseStatusNew, DisputeNewResponseStatusPendingCustomer, DisputeNewResponseStatusPrearbitration, DisputeNewResponseStatusRepresentment, DisputeNewResponseStatusSubmitted:
		return true
	}
	return false
}

// Dispute.
type DisputeGetResponse struct {
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
	Reason DisputeGetResponseReason `json:"reason,required"`
	// Date the representment was received.
	RepresentmentDate time.Time `json:"representment_date,required,nullable" format:"date-time"`
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
	ResolutionReason DisputeGetResponseResolutionReason `json:"resolution_reason,required,nullable"`
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
	Status DisputeGetResponseStatus `json:"status,required"`
	// The transaction that is being disputed. A transaction can only be disputed once
	// but may have multiple dispute cases.
	TransactionToken string                 `json:"transaction_token,required" format:"uuid"`
	JSON             disputeGetResponseJSON `json:"-"`
}

// disputeGetResponseJSON contains the JSON metadata for the struct
// [DisputeGetResponse]
type disputeGetResponseJSON struct {
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
	ResolutionDate     apijson.Field
	ResolutionNote     apijson.Field
	ResolutionReason   apijson.Field
	Status             apijson.Field
	TransactionToken   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DisputeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeGetResponseJSON) RawJSON() string {
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
type DisputeGetResponseReason string

const (
	DisputeGetResponseReasonAtmCashMisdispense               DisputeGetResponseReason = "ATM_CASH_MISDISPENSE"
	DisputeGetResponseReasonCancelled                        DisputeGetResponseReason = "CANCELLED"
	DisputeGetResponseReasonDuplicated                       DisputeGetResponseReason = "DUPLICATED"
	DisputeGetResponseReasonFraudCardNotPresent              DisputeGetResponseReason = "FRAUD_CARD_NOT_PRESENT"
	DisputeGetResponseReasonFraudCardPresent                 DisputeGetResponseReason = "FRAUD_CARD_PRESENT"
	DisputeGetResponseReasonFraudOther                       DisputeGetResponseReason = "FRAUD_OTHER"
	DisputeGetResponseReasonGoodsServicesNotAsDescribed      DisputeGetResponseReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	DisputeGetResponseReasonGoodsServicesNotReceived         DisputeGetResponseReason = "GOODS_SERVICES_NOT_RECEIVED"
	DisputeGetResponseReasonIncorrectAmount                  DisputeGetResponseReason = "INCORRECT_AMOUNT"
	DisputeGetResponseReasonMissingAuth                      DisputeGetResponseReason = "MISSING_AUTH"
	DisputeGetResponseReasonOther                            DisputeGetResponseReason = "OTHER"
	DisputeGetResponseReasonProcessingError                  DisputeGetResponseReason = "PROCESSING_ERROR"
	DisputeGetResponseReasonRecurringTransactionNotCancelled DisputeGetResponseReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	DisputeGetResponseReasonRefundNotProcessed               DisputeGetResponseReason = "REFUND_NOT_PROCESSED"
)

func (r DisputeGetResponseReason) IsKnown() bool {
	switch r {
	case DisputeGetResponseReasonAtmCashMisdispense, DisputeGetResponseReasonCancelled, DisputeGetResponseReasonDuplicated, DisputeGetResponseReasonFraudCardNotPresent, DisputeGetResponseReasonFraudCardPresent, DisputeGetResponseReasonFraudOther, DisputeGetResponseReasonGoodsServicesNotAsDescribed, DisputeGetResponseReasonGoodsServicesNotReceived, DisputeGetResponseReasonIncorrectAmount, DisputeGetResponseReasonMissingAuth, DisputeGetResponseReasonOther, DisputeGetResponseReasonProcessingError, DisputeGetResponseReasonRecurringTransactionNotCancelled, DisputeGetResponseReasonRefundNotProcessed:
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
type DisputeGetResponseResolutionReason string

const (
	DisputeGetResponseResolutionReasonCaseLost                      DisputeGetResponseResolutionReason = "CASE_LOST"
	DisputeGetResponseResolutionReasonNetworkRejected               DisputeGetResponseResolutionReason = "NETWORK_REJECTED"
	DisputeGetResponseResolutionReasonNoDisputeRights3DS            DisputeGetResponseResolutionReason = "NO_DISPUTE_RIGHTS_3DS"
	DisputeGetResponseResolutionReasonNoDisputeRightsBelowThreshold DisputeGetResponseResolutionReason = "NO_DISPUTE_RIGHTS_BELOW_THRESHOLD"
	DisputeGetResponseResolutionReasonNoDisputeRightsContactless    DisputeGetResponseResolutionReason = "NO_DISPUTE_RIGHTS_CONTACTLESS"
	DisputeGetResponseResolutionReasonNoDisputeRightsHybrid         DisputeGetResponseResolutionReason = "NO_DISPUTE_RIGHTS_HYBRID"
	DisputeGetResponseResolutionReasonNoDisputeRightsMaxChargebacks DisputeGetResponseResolutionReason = "NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS"
	DisputeGetResponseResolutionReasonNoDisputeRightsOther          DisputeGetResponseResolutionReason = "NO_DISPUTE_RIGHTS_OTHER"
	DisputeGetResponseResolutionReasonPastFilingDate                DisputeGetResponseResolutionReason = "PAST_FILING_DATE"
	DisputeGetResponseResolutionReasonPrearbitrationRejected        DisputeGetResponseResolutionReason = "PREARBITRATION_REJECTED"
	DisputeGetResponseResolutionReasonProcessorRejectedOther        DisputeGetResponseResolutionReason = "PROCESSOR_REJECTED_OTHER"
	DisputeGetResponseResolutionReasonRefunded                      DisputeGetResponseResolutionReason = "REFUNDED"
	DisputeGetResponseResolutionReasonRefundedAfterChargeback       DisputeGetResponseResolutionReason = "REFUNDED_AFTER_CHARGEBACK"
	DisputeGetResponseResolutionReasonWithdrawn                     DisputeGetResponseResolutionReason = "WITHDRAWN"
	DisputeGetResponseResolutionReasonWonArbitration                DisputeGetResponseResolutionReason = "WON_ARBITRATION"
	DisputeGetResponseResolutionReasonWonFirstChargeback            DisputeGetResponseResolutionReason = "WON_FIRST_CHARGEBACK"
	DisputeGetResponseResolutionReasonWonPrearbitration             DisputeGetResponseResolutionReason = "WON_PREARBITRATION"
)

func (r DisputeGetResponseResolutionReason) IsKnown() bool {
	switch r {
	case DisputeGetResponseResolutionReasonCaseLost, DisputeGetResponseResolutionReasonNetworkRejected, DisputeGetResponseResolutionReasonNoDisputeRights3DS, DisputeGetResponseResolutionReasonNoDisputeRightsBelowThreshold, DisputeGetResponseResolutionReasonNoDisputeRightsContactless, DisputeGetResponseResolutionReasonNoDisputeRightsHybrid, DisputeGetResponseResolutionReasonNoDisputeRightsMaxChargebacks, DisputeGetResponseResolutionReasonNoDisputeRightsOther, DisputeGetResponseResolutionReasonPastFilingDate, DisputeGetResponseResolutionReasonPrearbitrationRejected, DisputeGetResponseResolutionReasonProcessorRejectedOther, DisputeGetResponseResolutionReasonRefunded, DisputeGetResponseResolutionReasonRefundedAfterChargeback, DisputeGetResponseResolutionReasonWithdrawn, DisputeGetResponseResolutionReasonWonArbitration, DisputeGetResponseResolutionReasonWonFirstChargeback, DisputeGetResponseResolutionReasonWonPrearbitration:
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
type DisputeGetResponseStatus string

const (
	DisputeGetResponseStatusArbitration     DisputeGetResponseStatus = "ARBITRATION"
	DisputeGetResponseStatusCaseClosed      DisputeGetResponseStatus = "CASE_CLOSED"
	DisputeGetResponseStatusCaseWon         DisputeGetResponseStatus = "CASE_WON"
	DisputeGetResponseStatusNew             DisputeGetResponseStatus = "NEW"
	DisputeGetResponseStatusPendingCustomer DisputeGetResponseStatus = "PENDING_CUSTOMER"
	DisputeGetResponseStatusPrearbitration  DisputeGetResponseStatus = "PREARBITRATION"
	DisputeGetResponseStatusRepresentment   DisputeGetResponseStatus = "REPRESENTMENT"
	DisputeGetResponseStatusSubmitted       DisputeGetResponseStatus = "SUBMITTED"
)

func (r DisputeGetResponseStatus) IsKnown() bool {
	switch r {
	case DisputeGetResponseStatusArbitration, DisputeGetResponseStatusCaseClosed, DisputeGetResponseStatusCaseWon, DisputeGetResponseStatusNew, DisputeGetResponseStatusPendingCustomer, DisputeGetResponseStatusPrearbitration, DisputeGetResponseStatusRepresentment, DisputeGetResponseStatusSubmitted:
		return true
	}
	return false
}

// Dispute.
type DisputeUpdateResponse struct {
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
	Reason DisputeUpdateResponseReason `json:"reason,required"`
	// Date the representment was received.
	RepresentmentDate time.Time `json:"representment_date,required,nullable" format:"date-time"`
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
	ResolutionReason DisputeUpdateResponseResolutionReason `json:"resolution_reason,required,nullable"`
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
	Status DisputeUpdateResponseStatus `json:"status,required"`
	// The transaction that is being disputed. A transaction can only be disputed once
	// but may have multiple dispute cases.
	TransactionToken string                    `json:"transaction_token,required" format:"uuid"`
	JSON             disputeUpdateResponseJSON `json:"-"`
}

// disputeUpdateResponseJSON contains the JSON metadata for the struct
// [DisputeUpdateResponse]
type disputeUpdateResponseJSON struct {
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
	ResolutionDate     apijson.Field
	ResolutionNote     apijson.Field
	ResolutionReason   apijson.Field
	Status             apijson.Field
	TransactionToken   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DisputeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeUpdateResponseJSON) RawJSON() string {
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
type DisputeUpdateResponseReason string

const (
	DisputeUpdateResponseReasonAtmCashMisdispense               DisputeUpdateResponseReason = "ATM_CASH_MISDISPENSE"
	DisputeUpdateResponseReasonCancelled                        DisputeUpdateResponseReason = "CANCELLED"
	DisputeUpdateResponseReasonDuplicated                       DisputeUpdateResponseReason = "DUPLICATED"
	DisputeUpdateResponseReasonFraudCardNotPresent              DisputeUpdateResponseReason = "FRAUD_CARD_NOT_PRESENT"
	DisputeUpdateResponseReasonFraudCardPresent                 DisputeUpdateResponseReason = "FRAUD_CARD_PRESENT"
	DisputeUpdateResponseReasonFraudOther                       DisputeUpdateResponseReason = "FRAUD_OTHER"
	DisputeUpdateResponseReasonGoodsServicesNotAsDescribed      DisputeUpdateResponseReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	DisputeUpdateResponseReasonGoodsServicesNotReceived         DisputeUpdateResponseReason = "GOODS_SERVICES_NOT_RECEIVED"
	DisputeUpdateResponseReasonIncorrectAmount                  DisputeUpdateResponseReason = "INCORRECT_AMOUNT"
	DisputeUpdateResponseReasonMissingAuth                      DisputeUpdateResponseReason = "MISSING_AUTH"
	DisputeUpdateResponseReasonOther                            DisputeUpdateResponseReason = "OTHER"
	DisputeUpdateResponseReasonProcessingError                  DisputeUpdateResponseReason = "PROCESSING_ERROR"
	DisputeUpdateResponseReasonRecurringTransactionNotCancelled DisputeUpdateResponseReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	DisputeUpdateResponseReasonRefundNotProcessed               DisputeUpdateResponseReason = "REFUND_NOT_PROCESSED"
)

func (r DisputeUpdateResponseReason) IsKnown() bool {
	switch r {
	case DisputeUpdateResponseReasonAtmCashMisdispense, DisputeUpdateResponseReasonCancelled, DisputeUpdateResponseReasonDuplicated, DisputeUpdateResponseReasonFraudCardNotPresent, DisputeUpdateResponseReasonFraudCardPresent, DisputeUpdateResponseReasonFraudOther, DisputeUpdateResponseReasonGoodsServicesNotAsDescribed, DisputeUpdateResponseReasonGoodsServicesNotReceived, DisputeUpdateResponseReasonIncorrectAmount, DisputeUpdateResponseReasonMissingAuth, DisputeUpdateResponseReasonOther, DisputeUpdateResponseReasonProcessingError, DisputeUpdateResponseReasonRecurringTransactionNotCancelled, DisputeUpdateResponseReasonRefundNotProcessed:
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
type DisputeUpdateResponseResolutionReason string

const (
	DisputeUpdateResponseResolutionReasonCaseLost                      DisputeUpdateResponseResolutionReason = "CASE_LOST"
	DisputeUpdateResponseResolutionReasonNetworkRejected               DisputeUpdateResponseResolutionReason = "NETWORK_REJECTED"
	DisputeUpdateResponseResolutionReasonNoDisputeRights3DS            DisputeUpdateResponseResolutionReason = "NO_DISPUTE_RIGHTS_3DS"
	DisputeUpdateResponseResolutionReasonNoDisputeRightsBelowThreshold DisputeUpdateResponseResolutionReason = "NO_DISPUTE_RIGHTS_BELOW_THRESHOLD"
	DisputeUpdateResponseResolutionReasonNoDisputeRightsContactless    DisputeUpdateResponseResolutionReason = "NO_DISPUTE_RIGHTS_CONTACTLESS"
	DisputeUpdateResponseResolutionReasonNoDisputeRightsHybrid         DisputeUpdateResponseResolutionReason = "NO_DISPUTE_RIGHTS_HYBRID"
	DisputeUpdateResponseResolutionReasonNoDisputeRightsMaxChargebacks DisputeUpdateResponseResolutionReason = "NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS"
	DisputeUpdateResponseResolutionReasonNoDisputeRightsOther          DisputeUpdateResponseResolutionReason = "NO_DISPUTE_RIGHTS_OTHER"
	DisputeUpdateResponseResolutionReasonPastFilingDate                DisputeUpdateResponseResolutionReason = "PAST_FILING_DATE"
	DisputeUpdateResponseResolutionReasonPrearbitrationRejected        DisputeUpdateResponseResolutionReason = "PREARBITRATION_REJECTED"
	DisputeUpdateResponseResolutionReasonProcessorRejectedOther        DisputeUpdateResponseResolutionReason = "PROCESSOR_REJECTED_OTHER"
	DisputeUpdateResponseResolutionReasonRefunded                      DisputeUpdateResponseResolutionReason = "REFUNDED"
	DisputeUpdateResponseResolutionReasonRefundedAfterChargeback       DisputeUpdateResponseResolutionReason = "REFUNDED_AFTER_CHARGEBACK"
	DisputeUpdateResponseResolutionReasonWithdrawn                     DisputeUpdateResponseResolutionReason = "WITHDRAWN"
	DisputeUpdateResponseResolutionReasonWonArbitration                DisputeUpdateResponseResolutionReason = "WON_ARBITRATION"
	DisputeUpdateResponseResolutionReasonWonFirstChargeback            DisputeUpdateResponseResolutionReason = "WON_FIRST_CHARGEBACK"
	DisputeUpdateResponseResolutionReasonWonPrearbitration             DisputeUpdateResponseResolutionReason = "WON_PREARBITRATION"
)

func (r DisputeUpdateResponseResolutionReason) IsKnown() bool {
	switch r {
	case DisputeUpdateResponseResolutionReasonCaseLost, DisputeUpdateResponseResolutionReasonNetworkRejected, DisputeUpdateResponseResolutionReasonNoDisputeRights3DS, DisputeUpdateResponseResolutionReasonNoDisputeRightsBelowThreshold, DisputeUpdateResponseResolutionReasonNoDisputeRightsContactless, DisputeUpdateResponseResolutionReasonNoDisputeRightsHybrid, DisputeUpdateResponseResolutionReasonNoDisputeRightsMaxChargebacks, DisputeUpdateResponseResolutionReasonNoDisputeRightsOther, DisputeUpdateResponseResolutionReasonPastFilingDate, DisputeUpdateResponseResolutionReasonPrearbitrationRejected, DisputeUpdateResponseResolutionReasonProcessorRejectedOther, DisputeUpdateResponseResolutionReasonRefunded, DisputeUpdateResponseResolutionReasonRefundedAfterChargeback, DisputeUpdateResponseResolutionReasonWithdrawn, DisputeUpdateResponseResolutionReasonWonArbitration, DisputeUpdateResponseResolutionReasonWonFirstChargeback, DisputeUpdateResponseResolutionReasonWonPrearbitration:
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
type DisputeUpdateResponseStatus string

const (
	DisputeUpdateResponseStatusArbitration     DisputeUpdateResponseStatus = "ARBITRATION"
	DisputeUpdateResponseStatusCaseClosed      DisputeUpdateResponseStatus = "CASE_CLOSED"
	DisputeUpdateResponseStatusCaseWon         DisputeUpdateResponseStatus = "CASE_WON"
	DisputeUpdateResponseStatusNew             DisputeUpdateResponseStatus = "NEW"
	DisputeUpdateResponseStatusPendingCustomer DisputeUpdateResponseStatus = "PENDING_CUSTOMER"
	DisputeUpdateResponseStatusPrearbitration  DisputeUpdateResponseStatus = "PREARBITRATION"
	DisputeUpdateResponseStatusRepresentment   DisputeUpdateResponseStatus = "REPRESENTMENT"
	DisputeUpdateResponseStatusSubmitted       DisputeUpdateResponseStatus = "SUBMITTED"
)

func (r DisputeUpdateResponseStatus) IsKnown() bool {
	switch r {
	case DisputeUpdateResponseStatusArbitration, DisputeUpdateResponseStatusCaseClosed, DisputeUpdateResponseStatusCaseWon, DisputeUpdateResponseStatusNew, DisputeUpdateResponseStatusPendingCustomer, DisputeUpdateResponseStatusPrearbitration, DisputeUpdateResponseStatusRepresentment, DisputeUpdateResponseStatusSubmitted:
		return true
	}
	return false
}

// Dispute.
type DisputeListResponse struct {
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
	Reason DisputeListResponseReason `json:"reason,required"`
	// Date the representment was received.
	RepresentmentDate time.Time `json:"representment_date,required,nullable" format:"date-time"`
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
	ResolutionReason DisputeListResponseResolutionReason `json:"resolution_reason,required,nullable"`
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
	Status DisputeListResponseStatus `json:"status,required"`
	// The transaction that is being disputed. A transaction can only be disputed once
	// but may have multiple dispute cases.
	TransactionToken string                  `json:"transaction_token,required" format:"uuid"`
	JSON             disputeListResponseJSON `json:"-"`
}

// disputeListResponseJSON contains the JSON metadata for the struct
// [DisputeListResponse]
type disputeListResponseJSON struct {
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
	ResolutionDate     apijson.Field
	ResolutionNote     apijson.Field
	ResolutionReason   apijson.Field
	Status             apijson.Field
	TransactionToken   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DisputeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeListResponseJSON) RawJSON() string {
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
type DisputeListResponseReason string

const (
	DisputeListResponseReasonAtmCashMisdispense               DisputeListResponseReason = "ATM_CASH_MISDISPENSE"
	DisputeListResponseReasonCancelled                        DisputeListResponseReason = "CANCELLED"
	DisputeListResponseReasonDuplicated                       DisputeListResponseReason = "DUPLICATED"
	DisputeListResponseReasonFraudCardNotPresent              DisputeListResponseReason = "FRAUD_CARD_NOT_PRESENT"
	DisputeListResponseReasonFraudCardPresent                 DisputeListResponseReason = "FRAUD_CARD_PRESENT"
	DisputeListResponseReasonFraudOther                       DisputeListResponseReason = "FRAUD_OTHER"
	DisputeListResponseReasonGoodsServicesNotAsDescribed      DisputeListResponseReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	DisputeListResponseReasonGoodsServicesNotReceived         DisputeListResponseReason = "GOODS_SERVICES_NOT_RECEIVED"
	DisputeListResponseReasonIncorrectAmount                  DisputeListResponseReason = "INCORRECT_AMOUNT"
	DisputeListResponseReasonMissingAuth                      DisputeListResponseReason = "MISSING_AUTH"
	DisputeListResponseReasonOther                            DisputeListResponseReason = "OTHER"
	DisputeListResponseReasonProcessingError                  DisputeListResponseReason = "PROCESSING_ERROR"
	DisputeListResponseReasonRecurringTransactionNotCancelled DisputeListResponseReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	DisputeListResponseReasonRefundNotProcessed               DisputeListResponseReason = "REFUND_NOT_PROCESSED"
)

func (r DisputeListResponseReason) IsKnown() bool {
	switch r {
	case DisputeListResponseReasonAtmCashMisdispense, DisputeListResponseReasonCancelled, DisputeListResponseReasonDuplicated, DisputeListResponseReasonFraudCardNotPresent, DisputeListResponseReasonFraudCardPresent, DisputeListResponseReasonFraudOther, DisputeListResponseReasonGoodsServicesNotAsDescribed, DisputeListResponseReasonGoodsServicesNotReceived, DisputeListResponseReasonIncorrectAmount, DisputeListResponseReasonMissingAuth, DisputeListResponseReasonOther, DisputeListResponseReasonProcessingError, DisputeListResponseReasonRecurringTransactionNotCancelled, DisputeListResponseReasonRefundNotProcessed:
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
type DisputeListResponseResolutionReason string

const (
	DisputeListResponseResolutionReasonCaseLost                      DisputeListResponseResolutionReason = "CASE_LOST"
	DisputeListResponseResolutionReasonNetworkRejected               DisputeListResponseResolutionReason = "NETWORK_REJECTED"
	DisputeListResponseResolutionReasonNoDisputeRights3DS            DisputeListResponseResolutionReason = "NO_DISPUTE_RIGHTS_3DS"
	DisputeListResponseResolutionReasonNoDisputeRightsBelowThreshold DisputeListResponseResolutionReason = "NO_DISPUTE_RIGHTS_BELOW_THRESHOLD"
	DisputeListResponseResolutionReasonNoDisputeRightsContactless    DisputeListResponseResolutionReason = "NO_DISPUTE_RIGHTS_CONTACTLESS"
	DisputeListResponseResolutionReasonNoDisputeRightsHybrid         DisputeListResponseResolutionReason = "NO_DISPUTE_RIGHTS_HYBRID"
	DisputeListResponseResolutionReasonNoDisputeRightsMaxChargebacks DisputeListResponseResolutionReason = "NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS"
	DisputeListResponseResolutionReasonNoDisputeRightsOther          DisputeListResponseResolutionReason = "NO_DISPUTE_RIGHTS_OTHER"
	DisputeListResponseResolutionReasonPastFilingDate                DisputeListResponseResolutionReason = "PAST_FILING_DATE"
	DisputeListResponseResolutionReasonPrearbitrationRejected        DisputeListResponseResolutionReason = "PREARBITRATION_REJECTED"
	DisputeListResponseResolutionReasonProcessorRejectedOther        DisputeListResponseResolutionReason = "PROCESSOR_REJECTED_OTHER"
	DisputeListResponseResolutionReasonRefunded                      DisputeListResponseResolutionReason = "REFUNDED"
	DisputeListResponseResolutionReasonRefundedAfterChargeback       DisputeListResponseResolutionReason = "REFUNDED_AFTER_CHARGEBACK"
	DisputeListResponseResolutionReasonWithdrawn                     DisputeListResponseResolutionReason = "WITHDRAWN"
	DisputeListResponseResolutionReasonWonArbitration                DisputeListResponseResolutionReason = "WON_ARBITRATION"
	DisputeListResponseResolutionReasonWonFirstChargeback            DisputeListResponseResolutionReason = "WON_FIRST_CHARGEBACK"
	DisputeListResponseResolutionReasonWonPrearbitration             DisputeListResponseResolutionReason = "WON_PREARBITRATION"
)

func (r DisputeListResponseResolutionReason) IsKnown() bool {
	switch r {
	case DisputeListResponseResolutionReasonCaseLost, DisputeListResponseResolutionReasonNetworkRejected, DisputeListResponseResolutionReasonNoDisputeRights3DS, DisputeListResponseResolutionReasonNoDisputeRightsBelowThreshold, DisputeListResponseResolutionReasonNoDisputeRightsContactless, DisputeListResponseResolutionReasonNoDisputeRightsHybrid, DisputeListResponseResolutionReasonNoDisputeRightsMaxChargebacks, DisputeListResponseResolutionReasonNoDisputeRightsOther, DisputeListResponseResolutionReasonPastFilingDate, DisputeListResponseResolutionReasonPrearbitrationRejected, DisputeListResponseResolutionReasonProcessorRejectedOther, DisputeListResponseResolutionReasonRefunded, DisputeListResponseResolutionReasonRefundedAfterChargeback, DisputeListResponseResolutionReasonWithdrawn, DisputeListResponseResolutionReasonWonArbitration, DisputeListResponseResolutionReasonWonFirstChargeback, DisputeListResponseResolutionReasonWonPrearbitration:
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
type DisputeListResponseStatus string

const (
	DisputeListResponseStatusArbitration     DisputeListResponseStatus = "ARBITRATION"
	DisputeListResponseStatusCaseClosed      DisputeListResponseStatus = "CASE_CLOSED"
	DisputeListResponseStatusCaseWon         DisputeListResponseStatus = "CASE_WON"
	DisputeListResponseStatusNew             DisputeListResponseStatus = "NEW"
	DisputeListResponseStatusPendingCustomer DisputeListResponseStatus = "PENDING_CUSTOMER"
	DisputeListResponseStatusPrearbitration  DisputeListResponseStatus = "PREARBITRATION"
	DisputeListResponseStatusRepresentment   DisputeListResponseStatus = "REPRESENTMENT"
	DisputeListResponseStatusSubmitted       DisputeListResponseStatus = "SUBMITTED"
)

func (r DisputeListResponseStatus) IsKnown() bool {
	switch r {
	case DisputeListResponseStatusArbitration, DisputeListResponseStatusCaseClosed, DisputeListResponseStatusCaseWon, DisputeListResponseStatusNew, DisputeListResponseStatusPendingCustomer, DisputeListResponseStatusPrearbitration, DisputeListResponseStatusRepresentment, DisputeListResponseStatusSubmitted:
		return true
	}
	return false
}

// Dispute.
type DisputeDeleteResponse struct {
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
	Reason DisputeDeleteResponseReason `json:"reason,required"`
	// Date the representment was received.
	RepresentmentDate time.Time `json:"representment_date,required,nullable" format:"date-time"`
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
	ResolutionReason DisputeDeleteResponseResolutionReason `json:"resolution_reason,required,nullable"`
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
	Status DisputeDeleteResponseStatus `json:"status,required"`
	// The transaction that is being disputed. A transaction can only be disputed once
	// but may have multiple dispute cases.
	TransactionToken string                    `json:"transaction_token,required" format:"uuid"`
	JSON             disputeDeleteResponseJSON `json:"-"`
}

// disputeDeleteResponseJSON contains the JSON metadata for the struct
// [DisputeDeleteResponse]
type disputeDeleteResponseJSON struct {
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
	ResolutionDate     apijson.Field
	ResolutionNote     apijson.Field
	ResolutionReason   apijson.Field
	Status             apijson.Field
	TransactionToken   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DisputeDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeDeleteResponseJSON) RawJSON() string {
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
type DisputeDeleteResponseReason string

const (
	DisputeDeleteResponseReasonAtmCashMisdispense               DisputeDeleteResponseReason = "ATM_CASH_MISDISPENSE"
	DisputeDeleteResponseReasonCancelled                        DisputeDeleteResponseReason = "CANCELLED"
	DisputeDeleteResponseReasonDuplicated                       DisputeDeleteResponseReason = "DUPLICATED"
	DisputeDeleteResponseReasonFraudCardNotPresent              DisputeDeleteResponseReason = "FRAUD_CARD_NOT_PRESENT"
	DisputeDeleteResponseReasonFraudCardPresent                 DisputeDeleteResponseReason = "FRAUD_CARD_PRESENT"
	DisputeDeleteResponseReasonFraudOther                       DisputeDeleteResponseReason = "FRAUD_OTHER"
	DisputeDeleteResponseReasonGoodsServicesNotAsDescribed      DisputeDeleteResponseReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	DisputeDeleteResponseReasonGoodsServicesNotReceived         DisputeDeleteResponseReason = "GOODS_SERVICES_NOT_RECEIVED"
	DisputeDeleteResponseReasonIncorrectAmount                  DisputeDeleteResponseReason = "INCORRECT_AMOUNT"
	DisputeDeleteResponseReasonMissingAuth                      DisputeDeleteResponseReason = "MISSING_AUTH"
	DisputeDeleteResponseReasonOther                            DisputeDeleteResponseReason = "OTHER"
	DisputeDeleteResponseReasonProcessingError                  DisputeDeleteResponseReason = "PROCESSING_ERROR"
	DisputeDeleteResponseReasonRecurringTransactionNotCancelled DisputeDeleteResponseReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	DisputeDeleteResponseReasonRefundNotProcessed               DisputeDeleteResponseReason = "REFUND_NOT_PROCESSED"
)

func (r DisputeDeleteResponseReason) IsKnown() bool {
	switch r {
	case DisputeDeleteResponseReasonAtmCashMisdispense, DisputeDeleteResponseReasonCancelled, DisputeDeleteResponseReasonDuplicated, DisputeDeleteResponseReasonFraudCardNotPresent, DisputeDeleteResponseReasonFraudCardPresent, DisputeDeleteResponseReasonFraudOther, DisputeDeleteResponseReasonGoodsServicesNotAsDescribed, DisputeDeleteResponseReasonGoodsServicesNotReceived, DisputeDeleteResponseReasonIncorrectAmount, DisputeDeleteResponseReasonMissingAuth, DisputeDeleteResponseReasonOther, DisputeDeleteResponseReasonProcessingError, DisputeDeleteResponseReasonRecurringTransactionNotCancelled, DisputeDeleteResponseReasonRefundNotProcessed:
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
type DisputeDeleteResponseResolutionReason string

const (
	DisputeDeleteResponseResolutionReasonCaseLost                      DisputeDeleteResponseResolutionReason = "CASE_LOST"
	DisputeDeleteResponseResolutionReasonNetworkRejected               DisputeDeleteResponseResolutionReason = "NETWORK_REJECTED"
	DisputeDeleteResponseResolutionReasonNoDisputeRights3DS            DisputeDeleteResponseResolutionReason = "NO_DISPUTE_RIGHTS_3DS"
	DisputeDeleteResponseResolutionReasonNoDisputeRightsBelowThreshold DisputeDeleteResponseResolutionReason = "NO_DISPUTE_RIGHTS_BELOW_THRESHOLD"
	DisputeDeleteResponseResolutionReasonNoDisputeRightsContactless    DisputeDeleteResponseResolutionReason = "NO_DISPUTE_RIGHTS_CONTACTLESS"
	DisputeDeleteResponseResolutionReasonNoDisputeRightsHybrid         DisputeDeleteResponseResolutionReason = "NO_DISPUTE_RIGHTS_HYBRID"
	DisputeDeleteResponseResolutionReasonNoDisputeRightsMaxChargebacks DisputeDeleteResponseResolutionReason = "NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS"
	DisputeDeleteResponseResolutionReasonNoDisputeRightsOther          DisputeDeleteResponseResolutionReason = "NO_DISPUTE_RIGHTS_OTHER"
	DisputeDeleteResponseResolutionReasonPastFilingDate                DisputeDeleteResponseResolutionReason = "PAST_FILING_DATE"
	DisputeDeleteResponseResolutionReasonPrearbitrationRejected        DisputeDeleteResponseResolutionReason = "PREARBITRATION_REJECTED"
	DisputeDeleteResponseResolutionReasonProcessorRejectedOther        DisputeDeleteResponseResolutionReason = "PROCESSOR_REJECTED_OTHER"
	DisputeDeleteResponseResolutionReasonRefunded                      DisputeDeleteResponseResolutionReason = "REFUNDED"
	DisputeDeleteResponseResolutionReasonRefundedAfterChargeback       DisputeDeleteResponseResolutionReason = "REFUNDED_AFTER_CHARGEBACK"
	DisputeDeleteResponseResolutionReasonWithdrawn                     DisputeDeleteResponseResolutionReason = "WITHDRAWN"
	DisputeDeleteResponseResolutionReasonWonArbitration                DisputeDeleteResponseResolutionReason = "WON_ARBITRATION"
	DisputeDeleteResponseResolutionReasonWonFirstChargeback            DisputeDeleteResponseResolutionReason = "WON_FIRST_CHARGEBACK"
	DisputeDeleteResponseResolutionReasonWonPrearbitration             DisputeDeleteResponseResolutionReason = "WON_PREARBITRATION"
)

func (r DisputeDeleteResponseResolutionReason) IsKnown() bool {
	switch r {
	case DisputeDeleteResponseResolutionReasonCaseLost, DisputeDeleteResponseResolutionReasonNetworkRejected, DisputeDeleteResponseResolutionReasonNoDisputeRights3DS, DisputeDeleteResponseResolutionReasonNoDisputeRightsBelowThreshold, DisputeDeleteResponseResolutionReasonNoDisputeRightsContactless, DisputeDeleteResponseResolutionReasonNoDisputeRightsHybrid, DisputeDeleteResponseResolutionReasonNoDisputeRightsMaxChargebacks, DisputeDeleteResponseResolutionReasonNoDisputeRightsOther, DisputeDeleteResponseResolutionReasonPastFilingDate, DisputeDeleteResponseResolutionReasonPrearbitrationRejected, DisputeDeleteResponseResolutionReasonProcessorRejectedOther, DisputeDeleteResponseResolutionReasonRefunded, DisputeDeleteResponseResolutionReasonRefundedAfterChargeback, DisputeDeleteResponseResolutionReasonWithdrawn, DisputeDeleteResponseResolutionReasonWonArbitration, DisputeDeleteResponseResolutionReasonWonFirstChargeback, DisputeDeleteResponseResolutionReasonWonPrearbitration:
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
type DisputeDeleteResponseStatus string

const (
	DisputeDeleteResponseStatusArbitration     DisputeDeleteResponseStatus = "ARBITRATION"
	DisputeDeleteResponseStatusCaseClosed      DisputeDeleteResponseStatus = "CASE_CLOSED"
	DisputeDeleteResponseStatusCaseWon         DisputeDeleteResponseStatus = "CASE_WON"
	DisputeDeleteResponseStatusNew             DisputeDeleteResponseStatus = "NEW"
	DisputeDeleteResponseStatusPendingCustomer DisputeDeleteResponseStatus = "PENDING_CUSTOMER"
	DisputeDeleteResponseStatusPrearbitration  DisputeDeleteResponseStatus = "PREARBITRATION"
	DisputeDeleteResponseStatusRepresentment   DisputeDeleteResponseStatus = "REPRESENTMENT"
	DisputeDeleteResponseStatusSubmitted       DisputeDeleteResponseStatus = "SUBMITTED"
)

func (r DisputeDeleteResponseStatus) IsKnown() bool {
	switch r {
	case DisputeDeleteResponseStatusArbitration, DisputeDeleteResponseStatusCaseClosed, DisputeDeleteResponseStatusCaseWon, DisputeDeleteResponseStatusNew, DisputeDeleteResponseStatusPendingCustomer, DisputeDeleteResponseStatusPrearbitration, DisputeDeleteResponseStatusRepresentment, DisputeDeleteResponseStatusSubmitted:
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
