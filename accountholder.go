package lithic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// AccountHolderService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountHolderService] method
// instead.
type AccountHolderService struct {
	Options []option.RequestOption
}

// NewAccountHolderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountHolderService(opts ...option.RequestOption) (r *AccountHolderService) {
	r = &AccountHolderService{}
	r.Options = opts
	return
}

// Run an individual or business's information through the Customer Identification
// Program (CIP) and return an `account_token` if the status is accepted or pending
// (i.e., further action required). All calls to this endpoint will return an
// immediate response - though in some cases, the response may indicate the
// workflow is under review or further action will be needed to complete the
// account creation process. This endpoint can only be used on accounts that are
// part of the program the calling API key manages.
//
// Note: If you choose to set a timeout for this request, we recommend 5 minutes.
func (r *AccountHolderService) New(ctx context.Context, body AccountHolderNewParams, opts ...option.RequestOption) (res *AccountHolder, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_holders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an Individual or Business Account Holder and/or their KYC or KYB evaluation
// status.
func (r *AccountHolderService) Get(ctx context.Context, account_holder_token string, opts ...option.RequestOption) (res *AccountHolder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s", account_holder_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the information associated with a particular account holder.
func (r *AccountHolderService) Update(ctx context.Context, account_holder_token string, body AccountHolderUpdateParams, opts ...option.RequestOption) (res *AccountHolderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s", account_holder_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Create a webhook to receive KYC or KYB evaluation events.
//
// There are two types of account holder webhooks:
//
//   - `verification`: Webhook sent when the status of a KYC or KYB evaluation
//     changes from `PENDING_DOCUMENT` (KYC) or `PENDING` (KYB) to `ACCEPTED` or
//     `REJECTED`.
//   - `document_upload_front`/`document_upload_back`: Webhook sent when a document
//     upload fails.
//
// After a webhook has been created, this endpoint can be used to rotate a webhooks
// HMAC token or modify the registered URL. Only a single webhook is allowed per
// program. Since HMAC verification is available, the IP addresses from which
// KYC/KYB webhooks are sent are subject to change.
func (r *AccountHolderService) NewWebhook(ctx context.Context, body AccountHolderNewWebhookParams, opts ...option.RequestOption) (res *AccountHolderCreateWebhookResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "webhooks/account_holders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the status of account holder document uploads, or retrieve the upload
// URLs to process your image uploads.
//
// Note that this is not equivalent to checking the status of the KYC evaluation
// overall (a document may be successfully uploaded but not be sufficient for KYC
// to pass).
//
// In the event your upload URLs have expired, calling this endpoint will refresh
// them. Similarly, in the event a previous account holder document upload has
// failed, you can use this endpoint to get a new upload URL for the failed image
// upload.
//
// When a new document upload is generated for a failed attempt, the response will
// show an additional entry in the `required_document_uploads` list in a `PENDING`
// state for the corresponding `image_type`.
func (r *AccountHolderService) ListDocuments(ctx context.Context, account_holder_token string, opts ...option.RequestOption) (res *AccountHolderListDocumentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s/documents", account_holder_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resubmit a KYC submission. This endpoint should be used in cases where a KYC
// submission returned a `PENDING_RESUBMIT` result, meaning one or more critical
// KYC fields may have been mis-entered and the individual's identity has not yet
// been successfully verified. This step must be completed in order to proceed with
// the KYC evaluation.
//
// Two resubmission attempts are permitted via this endpoint before a `REJECTED`
// status is returned and the account creation process is ended.
func (r *AccountHolderService) Resubmit(ctx context.Context, account_holder_token string, body AccountHolderResubmitParams, opts ...option.RequestOption) (res *AccountHolder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s/resubmit", account_holder_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check the status of an account holder document upload, or retrieve the upload
// URLs to process your image uploads.
//
// Note that this is not equivalent to checking the status of the KYC evaluation
// overall (a document may be successfully uploaded but not be sufficient for KYC
// to pass).
//
// In the event your upload URLs have expired, calling this endpoint will refresh
// them. Similarly, in the event a document upload has failed, you can use this
// endpoint to get a new upload URL for the failed image upload.
//
// When a new account holder document upload is generated for a failed attempt, the
// response will show an additional entry in the `required_document_uploads` array
// in a `PENDING` state for the corresponding `image_type`.
func (r *AccountHolderService) GetDocument(ctx context.Context, account_holder_token string, document_token string, opts ...option.RequestOption) (res *AccountHolderDocument, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s/documents/%s", account_holder_token, document_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to identify which type of supported government-issued
// documentation you will upload for further verification. It will return two URLs
// to upload your document images to - one for the front image and one for the back
// image.
//
// This endpoint is only valid for evaluations in a `PENDING_DOCUMENT` state.
//
// Uploaded images must either be a `jpg` or `png` file, and each must be less than
// 15 MiB. Once both required uploads have been successfully completed, your
// document will be run through KYC verification.
//
// If you have registered a webhook, you will receive evaluation updates for any
// document submission evaluations, as well as for any failed document uploads.
//
// Two document submission attempts are permitted via this endpoint before a
// `REJECTED` status is returned and the account creation process is ended.
// Currently only one type of account holder document is supported per KYC
// verification.
func (r *AccountHolderService) UploadDocument(ctx context.Context, account_holder_token string, body AccountHolderUploadDocumentParams, opts ...option.RequestOption) (res *AccountHolderDocument, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s/documents", account_holder_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

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
	JSON          accountHolderJSON
}

// accountHolderJSON contains the JSON metadata for the struct [AccountHolder]
type accountHolderJSON struct {
	Token                apijson.Field
	AccountToken         apijson.Field
	BusinessAccountToken apijson.Field
	Status               apijson.Field
	StatusReasons        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Describes the document and the required document image uploads required to
// re-run KYC.
type AccountHolderDocument struct {
	// Globally unique identifier for the account holder.
	AccountHolderToken string `json:"account_holder_token" format:"uuid"`
	// Type of documentation to be submitted for verification.
	DocumentType            AccountHolderDocumentDocumentType              `json:"document_type"`
	RequiredDocumentUploads []AccountHolderDocumentRequiredDocumentUploads `json:"required_document_uploads"`
	// Globally unique identifier for the document.
	Token string `json:"token" format:"uuid"`
	JSON  accountHolderDocumentJSON
}

// accountHolderDocumentJSON contains the JSON metadata for the struct
// [AccountHolderDocument]
type accountHolderDocumentJSON struct {
	AccountHolderToken      apijson.Field
	DocumentType            apijson.Field
	RequiredDocumentUploads apijson.Field
	Token                   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountHolderDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHolderDocumentDocumentType string

const (
	AccountHolderDocumentDocumentTypeCommercialLicense AccountHolderDocumentDocumentType = "commercial_license"
	AccountHolderDocumentDocumentTypeDriversLicense    AccountHolderDocumentDocumentType = "drivers_license"
	AccountHolderDocumentDocumentTypePassport          AccountHolderDocumentDocumentType = "passport"
	AccountHolderDocumentDocumentTypePassportCard      AccountHolderDocumentDocumentType = "passport_card"
	AccountHolderDocumentDocumentTypeVisa              AccountHolderDocumentDocumentType = "visa"
)

// Represents a single image of the document to upload.
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
	JSON      accountHolderDocumentRequiredDocumentUploadsJSON
}

// accountHolderDocumentRequiredDocumentUploadsJSON contains the JSON metadata for
// the struct [AccountHolderDocumentRequiredDocumentUploads]
type accountHolderDocumentRequiredDocumentUploadsJSON struct {
	ImageType     apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	UploadURL     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountHolderDocumentRequiredDocumentUploads) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type KYBParam struct {
	// Information for business for which the account is being opened and KYB is being
	// run.
	BusinessEntity param.Field[KYBBusinessEntityParam] `json:"business_entity,required"`
	// List of all entities with >25% ownership in the company. If no entity or
	// individual owns >25% of the company, and the largest shareholder is an entity,
	// please identify them in this field. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background. If no business owner is an entity, pass in an
	// empty list. However, either this parameter or `beneficial_owner_individuals`
	// must be populated. on entities that should be included.
	BeneficialOwnerEntities param.Field[[]KYBBeneficialOwnerEntitiesParam] `json:"beneficial_owner_entities,required"`
	// List of all individuals with >25% ownership in the company. If no entity or
	// individual owns >25% of the company, and the largest shareholder is an
	// individual, please identify them in this field. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included. If no
	// individual is an entity, pass in an empty list. However, either this parameter
	// or `beneficial_owner_entities` must be populated.
	BeneficialOwnerIndividuals param.Field[[]KYBBeneficialOwnerIndividualsParam] `json:"beneficial_owner_individuals,required"`
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer, Managing Member, General Partner, President, Vice President, or
	// Treasurer). This can be an executive, or someone who will have program-wide
	// access to the cards that Lithic will provide. In some cases, this individual
	// could also be a beneficial owner listed above. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section II) for more background.
	ControlPerson param.Field[KYBControlPersonParam] `json:"control_person,required"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// business with a pass result.
	//
	// This field is required only if workflow type is `KYB_BYO`.
	KYBPassedTimestamp param.Field[string] `json:"kyb_passed_timestamp"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness param.Field[string] `json:"nature_of_business,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp param.Field[string] `json:"tos_timestamp,required"`
	// Company website URL.
	WebsiteURL param.Field[string] `json:"website_url,required"`
	// Specifies the type of KYB workflow to run.
	Workflow param.Field[KYBWorkflow] `json:"workflow,required"`
}

func (r KYBParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYBParam) implementsAccountHolderNewParams() {}

// Information for business for which the account is being opened and KYB is being
// run.
type KYBBusinessEntityParam struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName param.Field[string] `json:"dba_business_name"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID param.Field[string] `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName param.Field[string] `json:"legal_business_name,required"`
	// Parent company name (if applicable).
	ParentCompany param.Field[string] `json:"parent_company"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers param.Field[[]string] `json:"phone_numbers,required"`
}

func (r KYBBusinessEntityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBBeneficialOwnerEntitiesParam struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName param.Field[string] `json:"dba_business_name"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID param.Field[string] `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName param.Field[string] `json:"legal_business_name,required"`
	// Parent company name (if applicable).
	ParentCompany param.Field[string] `json:"parent_company"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers param.Field[[]string] `json:"phone_numbers,required"`
}

func (r KYBBeneficialOwnerEntitiesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBBeneficialOwnerIndividualsParam struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob param.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email param.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName param.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID param.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName param.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number,required"`
}

func (r KYBBeneficialOwnerIndividualsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An individual with significant responsibility for managing the legal entity
// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
// Officer, Managing Member, General Partner, President, Vice President, or
// Treasurer). This can be an executive, or someone who will have program-wide
// access to the cards that Lithic will provide. In some cases, this individual
// could also be a beneficial owner listed above. See
// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
// (Section II) for more background.
type KYBControlPersonParam struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob param.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email param.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName param.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID param.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName param.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number,required"`
}

func (r KYBControlPersonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBWorkflow string

const (
	KYBWorkflowKYBBasic KYBWorkflow = "KYB_BASIC"
	KYBWorkflowKYBByo   KYBWorkflow = "KYB_BYO"
)

type KYCParam struct {
	// Information on individual for whom the account is being opened and KYC is being
	// run.
	Individual param.Field[KYCIndividualParam] `json:"individual,required"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// individual with a pass result.
	//
	// This field is required only if workflow type is `KYC_BYO`.
	KYCPassedTimestamp param.Field[string] `json:"kyc_passed_timestamp"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp param.Field[string] `json:"tos_timestamp,required"`
	// Specifies the type of KYC workflow to run.
	Workflow param.Field[KYCWorkflow] `json:"workflow,required"`
}

func (r KYCParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYCParam) implementsAccountHolderNewParams() {}

// Information on individual for whom the account is being opened and KYC is being
// run.
type KYCIndividualParam struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob param.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email param.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName param.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID param.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName param.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number,required"`
}

func (r KYCIndividualParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYCWorkflow string

const (
	KYCWorkflowKYCAdvanced KYCWorkflow = "KYC_ADVANCED"
	KYCWorkflowKYCBasic    KYCWorkflow = "KYC_BASIC"
	KYCWorkflowKYCByo      KYCWorkflow = "KYC_BYO"
)

type KYCExemptParam struct {
	// Specifies the workflow type. This must be 'KYC_EXEMPT'
	Workflow param.Field[KYCExemptWorkflow] `json:"workflow,required"`
	// Specifies the type of KYC Exempt user
	KYCExemptionType param.Field[KYCExemptKYCExemptionType] `json:"kyc_exemption_type,required"`
	// The KYC Exempt user's first name
	FirstName param.Field[string] `json:"first_name,required"`
	// The KYC Exempt user's last name
	LastName param.Field[string] `json:"last_name,required"`
	// The KYC Exempt user's email
	Email param.Field[string] `json:"email,required"`
	// The KYC Exempt user's phone number
	PhoneNumber param.Field[string] `json:"phone_number,required"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken param.Field[string] `json:"business_account_token"`
	// KYC Exempt user's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[shared.AddressParam] `json:"address"`
}

func (r KYCExemptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYCExemptParam) implementsAccountHolderNewParams() {}

type KYCExemptWorkflow string

const (
	KYCExemptWorkflowKYCExempt KYCExemptWorkflow = "KYC_EXEMPT"
)

type KYCExemptKYCExemptionType string

const (
	KYCExemptKYCExemptionTypeAuthorizedUser  KYCExemptKYCExemptionType = "AUTHORIZED_USER"
	KYCExemptKYCExemptionTypePrepaidCardUser KYCExemptKYCExemptionType = "PREPAID_CARD_USER"
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
	JSON                 accountHolderUpdateResponseJSON
}

// accountHolderUpdateResponseJSON contains the JSON metadata for the struct
// [AccountHolderUpdateResponse]
type accountHolderUpdateResponseJSON struct {
	Token                apijson.Field
	Email                apijson.Field
	PhoneNumber          apijson.Field
	BusinessAccountToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountHolderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHolderListDocumentsResponse struct {
	Data []AccountHolderDocument `json:"data"`
	JSON accountHolderListDocumentsResponseJSON
}

// accountHolderListDocumentsResponseJSON contains the JSON metadata for the struct
// [AccountHolderListDocumentsResponse]
type accountHolderListDocumentsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderListDocumentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHolderCreateWebhookResponse struct {
	Data AccountHolderCreateWebhookResponseData `json:"data"`
	JSON accountHolderCreateWebhookResponseJSON
}

// accountHolderCreateWebhookResponseJSON contains the JSON metadata for the struct
// [AccountHolderCreateWebhookResponse]
type accountHolderCreateWebhookResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderCreateWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountHolderCreateWebhookResponseData struct {
	// Shared secret which can optionally be used to validate the authenticity of
	// incoming identity webhooks.
	HmacToken string `json:"hmac_token" format:"uuid"`
	JSON      accountHolderCreateWebhookResponseDataJSON
}

// accountHolderCreateWebhookResponseDataJSON contains the JSON metadata for the
// struct [AccountHolderCreateWebhookResponseData]
type accountHolderCreateWebhookResponseDataJSON struct {
	HmacToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderCreateWebhookResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This interface is a union satisfied by one of the following: [KYBParam],
// [KYCParam], [KYCExemptParam].
type AccountHolderNewParams interface {
	implementsAccountHolderNewParams()
}

type AccountHolderUpdateParams struct {
	// Account holder's email address. The primary purpose of this field is for
	// cardholder identification and verification during the digital wallet
	// tokenization process.
	Email param.Field[string] `json:"email"`
	// Account holder's phone number, entered in E.164 format. The primary purpose of
	// this field is for cardholder identification and verification during the digital
	// wallet tokenization process.
	PhoneNumber param.Field[string] `json:"phone_number"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken param.Field[string] `json:"business_account_token"`
}

func (r AccountHolderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHolderNewWebhookParams struct {
	// URL to receive webhook requests. Must be a valid HTTPS address.
	URL param.Field[string] `json:"url,required"`
}

func (r AccountHolderNewWebhookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHolderResubmitParams struct {
	Workflow param.Field[AccountHolderResubmitParamsWorkflow] `json:"workflow,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp param.Field[string] `json:"tos_timestamp,required"`
	// Information on individual for whom the account is being opened and KYC is being
	// re-run.
	Individual param.Field[AccountHolderResubmitParamsIndividual] `json:"individual,required"`
}

func (r AccountHolderResubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHolderResubmitParamsWorkflow string

const (
	AccountHolderResubmitParamsWorkflowKYCAdvanced AccountHolderResubmitParamsWorkflow = "KYC_ADVANCED"
)

// Information on individual for whom the account is being opened and KYC is being
// re-run.
type AccountHolderResubmitParamsIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob param.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email param.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName param.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID param.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName param.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number,required"`
}

type AccountHolderUploadDocumentParams struct {
	// Type of the document to upload.
	DocumentType param.Field[AccountHolderUploadDocumentParamsDocumentType] `json:"document_type,required"`
}

func (r AccountHolderUploadDocumentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHolderUploadDocumentParamsDocumentType string

const (
	AccountHolderUploadDocumentParamsDocumentTypeCommercialLicense AccountHolderUploadDocumentParamsDocumentType = "commercial_license"
	AccountHolderUploadDocumentParamsDocumentTypeDriversLicense    AccountHolderUploadDocumentParamsDocumentType = "drivers_license"
	AccountHolderUploadDocumentParamsDocumentTypePassport          AccountHolderUploadDocumentParamsDocumentType = "passport"
	AccountHolderUploadDocumentParamsDocumentTypePassportCard      AccountHolderUploadDocumentParamsDocumentType = "passport_card"
	AccountHolderUploadDocumentParamsDocumentTypeVisa              AccountHolderUploadDocumentParamsDocumentType = "visa"
)