package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type AccountHolder struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token" format:"uuid"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token" format:"uuid"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken string `json:"business_account_token" format:"uuid"`
	// KYC and KYB evaluation states.
	//
	// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	// `ADVANCED` workflow.
	Status AccountHolderStatus `json:"status"`
	// Reason for the evaluation status.
	StatusReasons []AccountHolderStatusReasons `json:"status_reasons"`
	JSON          AccountHolderJSON
}

type AccountHolderJSON struct {
	Token                pjson.Metadata
	AccountToken         pjson.Metadata
	BusinessAccountToken pjson.Metadata
	Status               pjson.Metadata
	StatusReasons        pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountHolder using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountHolder) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountHolderStatus string

const (
	AccountHolderStatusAccepted        AccountHolderStatus = "ACCEPTED"
	AccountHolderStatusRejected        AccountHolderStatus = "REJECTED"
	AccountHolderStatusPendingResubmit AccountHolderStatus = "PENDING_RESUBMIT"
	AccountHolderStatusPendingDocument AccountHolderStatus = "PENDING_DOCUMENT"
)

type AccountHolderStatusReasons string

const (
	AccountHolderStatusReasonsAddressVerificationFailure  AccountHolderStatusReasons = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderStatusReasonsAgeThresholdFailure         AccountHolderStatusReasons = "AGE_THRESHOLD_FAILURE"
	AccountHolderStatusReasonsCompleteVerificationFailure AccountHolderStatusReasons = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderStatusReasonsDobVerificationFailure      AccountHolderStatusReasons = "DOB_VERIFICATION_FAILURE"
	AccountHolderStatusReasonsIDVerificationFailure       AccountHolderStatusReasons = "ID_VERIFICATION_FAILURE"
	AccountHolderStatusReasonsMaxDocumentAttempts         AccountHolderStatusReasons = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderStatusReasonsMaxResubmissionAttempts     AccountHolderStatusReasons = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderStatusReasonsNameVerificationFailure     AccountHolderStatusReasons = "NAME_VERIFICATION_FAILURE"
	AccountHolderStatusReasonsOtherVerificationFailure    AccountHolderStatusReasons = "OTHER_VERIFICATION_FAILURE"
	AccountHolderStatusReasonsRiskThresholdFailure        AccountHolderStatusReasons = "RISK_THRESHOLD_FAILURE"
	AccountHolderStatusReasonsWatchlistAlertFailure       AccountHolderStatusReasons = "WATCHLIST_ALERT_FAILURE"
)

type AccountHolderDocument struct {
	// Globally unique identifier for the account holder.
	AccountHolderToken string `json:"account_holder_token" format:"uuid"`
	// Type of documentation to be submitted for verification.
	DocumentType            AccountHolderDocumentDocumentType              `json:"document_type"`
	RequiredDocumentUploads []AccountHolderDocumentRequiredDocumentUploads `json:"required_document_uploads"`
	// Globally unique identifier for the document.
	Token string `json:"token" format:"uuid"`
	JSON  AccountHolderDocumentJSON
}

type AccountHolderDocumentJSON struct {
	AccountHolderToken      pjson.Metadata
	DocumentType            pjson.Metadata
	RequiredDocumentUploads pjson.Metadata
	Token                   pjson.Metadata
	Raw                     []byte
	Extras                  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountHolderDocument using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountHolderDocument) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountHolderDocumentDocumentType string

const (
	AccountHolderDocumentDocumentTypeCommercialLicense AccountHolderDocumentDocumentType = "commercial_license"
	AccountHolderDocumentDocumentTypeDriversLicense    AccountHolderDocumentDocumentType = "drivers_license"
	AccountHolderDocumentDocumentTypePassport          AccountHolderDocumentDocumentType = "passport"
	AccountHolderDocumentDocumentTypePassportCard      AccountHolderDocumentDocumentType = "passport_card"
	AccountHolderDocumentDocumentTypeVisa              AccountHolderDocumentDocumentType = "visa"
)

type AccountHolderDocumentRequiredDocumentUploads struct {
	// Type of image to upload.
	ImageType AccountHolderDocumentRequiredDocumentUploadsImageType `json:"image_type"`
	// Status of document image upload.
	Status        AccountHolderDocumentRequiredDocumentUploadsStatus          `json:"status"`
	StatusReasons []AccountHolderDocumentRequiredDocumentUploadsStatusReasons `json:"status_reasons"`
	// URL to upload document image to.
	//
	// Note that the upload URLs expire after 7 days. If an upload URL expires, you can
	// refresh the URLs by retrieving the document upload from
	// `GET /account_holders/{account_holder_token}/documents`.
	UploadURL string `json:"upload_url"`
	JSON      AccountHolderDocumentRequiredDocumentUploadsJSON
}

type AccountHolderDocumentRequiredDocumentUploadsJSON struct {
	ImageType     pjson.Metadata
	Status        pjson.Metadata
	StatusReasons pjson.Metadata
	UploadURL     pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// AccountHolderDocumentRequiredDocumentUploads using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AccountHolderDocumentRequiredDocumentUploads) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountHolderDocumentRequiredDocumentUploadsImageType string

const (
	AccountHolderDocumentRequiredDocumentUploadsImageTypeBack  AccountHolderDocumentRequiredDocumentUploadsImageType = "back"
	AccountHolderDocumentRequiredDocumentUploadsImageTypeFront AccountHolderDocumentRequiredDocumentUploadsImageType = "front"
)

type AccountHolderDocumentRequiredDocumentUploadsStatus string

const (
	AccountHolderDocumentRequiredDocumentUploadsStatusCompleted AccountHolderDocumentRequiredDocumentUploadsStatus = "COMPLETED"
	AccountHolderDocumentRequiredDocumentUploadsStatusFailed    AccountHolderDocumentRequiredDocumentUploadsStatus = "FAILED"
	AccountHolderDocumentRequiredDocumentUploadsStatusPending   AccountHolderDocumentRequiredDocumentUploadsStatus = "PENDING"
	AccountHolderDocumentRequiredDocumentUploadsStatusUploaded  AccountHolderDocumentRequiredDocumentUploadsStatus = "UPLOADED"
)

type AccountHolderDocumentRequiredDocumentUploadsStatusReasons string

const (
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonsBackImageBlurry  AccountHolderDocumentRequiredDocumentUploadsStatusReasons = "BACK_IMAGE_BLURRY"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonsFileSizeTooLarge AccountHolderDocumentRequiredDocumentUploadsStatusReasons = "FILE_SIZE_TOO_LARGE"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonsFrontImageBlurry AccountHolderDocumentRequiredDocumentUploadsStatusReasons = "FRONT_IMAGE_BLURRY"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonsFrontImageGlare  AccountHolderDocumentRequiredDocumentUploadsStatusReasons = "FRONT_IMAGE_GLARE"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonsInvalidFileType  AccountHolderDocumentRequiredDocumentUploadsStatusReasons = "INVALID_FILE_TYPE"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonsUnknownError     AccountHolderDocumentRequiredDocumentUploadsStatusReasons = "UNKNOWN_ERROR"
)

type AccountHolderUpdateResponse struct {
	// The token for the account holder that was updated
	Token string `json:"token"`
	// The newly updated email for the account holder
	Email string `json:"email"`
	// The newly updated phone_number for the account holder
	PhoneNumber string `json:"phone_number"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll businesses
	// with authorized users. Pass the account_token of the enrolled business
	// associated with the AUTHORIZED_USER in this field.
	BusinessAccountToken string `json:"business_account_token"`
	JSON                 AccountHolderUpdateResponseJSON
}

type AccountHolderUpdateResponseJSON struct {
	Token                pjson.Metadata
	Email                pjson.Metadata
	PhoneNumber          pjson.Metadata
	BusinessAccountToken pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountHolderUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountHolderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountHolderListDocumentsResponse struct {
	Data []AccountHolderDocument `json:"data"`
	JSON AccountHolderListDocumentsResponseJSON
}

type AccountHolderListDocumentsResponseJSON struct {
	Data   pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// AccountHolderListDocumentsResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AccountHolderListDocumentsResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountHolderCreateWebhookResponse struct {
	Data AccountHolderCreateWebhookResponseData `json:"data"`
	JSON AccountHolderCreateWebhookResponseJSON
}

type AccountHolderCreateWebhookResponseJSON struct {
	Data   pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// AccountHolderCreateWebhookResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AccountHolderCreateWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountHolderCreateWebhookResponseData struct {
	// Shared secret which can optionally be used to validate the authenticity of
	// incoming identity webhooks.
	HmacToken string `json:"hmac_token" format:"uuid"`
	JSON      AccountHolderCreateWebhookResponseDataJSON
}

type AccountHolderCreateWebhookResponseDataJSON struct {
	HmacToken pjson.Metadata
	Raw       []byte
	Extras    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// AccountHolderCreateWebhookResponseData using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AccountHolderCreateWebhookResponseData) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
