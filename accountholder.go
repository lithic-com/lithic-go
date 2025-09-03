// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/lithic-com/lithic-go/shared"
	"github.com/tidwall/gjson"
)

// AccountHolderService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountHolderService] method instead.
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

// Create an account holder and initiate the appropriate onboarding workflow.
// Account holders and accounts have a 1:1 relationship. When an account holder is
// successfully created an associated account is also created. All calls to this
// endpoint will return a synchronous response. The response time will depend on
// the workflow. In some cases, the response may indicate the workflow is under
// review or further action will be needed to complete the account creation
// process. This endpoint can only be used on accounts that are part of the program
// that the calling API key manages.
//
// Note: If you choose to set a timeout for this request, we recommend 5 minutes.
func (r *AccountHolderService) New(ctx context.Context, body AccountHolderNewParams, opts ...option.RequestOption) (res *AccountHolderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/account_holders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an Individual or Business Account Holder and/or their KYC or KYB evaluation
// status.
func (r *AccountHolderService) Get(ctx context.Context, accountHolderToken string, opts ...option.RequestOption) (res *AccountHolder, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	path := fmt.Sprintf("v1/account_holders/%s", accountHolderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the information associated with a particular account holder (including
// business owners and control persons associated to a business account). If Lithic
// is performing KYB or KYC and additional verification is required we will run the
// individual's or business's updated information again and return whether the
// status is accepted or pending (i.e., further action required). All calls to this
// endpoint will return a synchronous response. The response time will depend on
// the workflow. In some cases, the response may indicate the workflow is under
// review or further action will be needed to complete the account creation
// process. This endpoint can only be used on existing accounts that are part of
// the program that the calling API key manages.
func (r *AccountHolderService) Update(ctx context.Context, accountHolderToken string, body AccountHolderUpdateParams, opts ...option.RequestOption) (res *AccountHolderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	path := fmt.Sprintf("v1/account_holders/%s", accountHolderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of individual or business account holders and their KYC or KYB
// evaluation status.
func (r *AccountHolderService) List(ctx context.Context, query AccountHolderListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccountHolder], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/account_holders"
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

// Get a list of individual or business account holders and their KYC or KYB
// evaluation status.
func (r *AccountHolderService) ListAutoPaging(ctx context.Context, query AccountHolderListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccountHolder] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
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
func (r *AccountHolderService) ListDocuments(ctx context.Context, accountHolderToken string, opts ...option.RequestOption) (res *AccountHolderListDocumentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	path := fmt.Sprintf("v1/account_holders/%s/documents", accountHolderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *AccountHolderService) GetDocument(ctx context.Context, accountHolderToken string, documentToken string, opts ...option.RequestOption) (res *shared.Document, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	if documentToken == "" {
		err = errors.New("missing required document_token parameter")
		return
	}
	path := fmt.Sprintf("v1/account_holders/%s/documents/%s", accountHolderToken, documentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Simulates a review for an account holder document upload.
func (r *AccountHolderService) SimulateEnrollmentDocumentReview(ctx context.Context, body AccountHolderSimulateEnrollmentDocumentReviewParams, opts ...option.RequestOption) (res *shared.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/account_holders/enrollment_document_review"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates an enrollment review for an account holder. This endpoint is only
// applicable for workflows that may required intervention such as `KYB_BASIC`.
func (r *AccountHolderService) SimulateEnrollmentReview(ctx context.Context, body AccountHolderSimulateEnrollmentReviewParams, opts ...option.RequestOption) (res *AccountHolderSimulateEnrollmentReviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/account_holders/enrollment_review"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
func (r *AccountHolderService) UploadDocument(ctx context.Context, accountHolderToken string, body AccountHolderUploadDocumentParams, opts ...option.RequestOption) (res *shared.Document, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	path := fmt.Sprintf("v1/account_holders/%s/documents", accountHolderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountHolder struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token,required" format:"uuid"`
	// Timestamp of when the account holder was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token" format:"uuid"`
	// Deprecated.
	//
	// Deprecated: deprecated
	BeneficialOwnerEntities []AccountHolderBeneficialOwnerEntity `json:"beneficial_owner_entities"`
	// Only present when user_type == "BUSINESS". You must submit a list of all direct
	// and indirect individuals with 25% or more ownership in the company. A maximum of
	// 4 beneficial owners can be submitted. If no individual owns 25% of the company
	// you do not need to send beneficial owner information. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included.
	BeneficialOwnerIndividuals []AccountHolderBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken string `json:"business_account_token" format:"uuid"`
	// Only present when user_type == "BUSINESS". Information about the business for
	// which the account is being opened and KYB is being run.
	BusinessEntity AccountHolderBusinessEntity `json:"business_entity"`
	// Only present when user_type == "BUSINESS". An individual with significant
	// responsibility for managing the legal entity (e.g., a Chief Executive Officer,
	// Chief Financial Officer, Chief Operating Officer, Managing Member, General
	// Partner, President, Vice President, or Treasurer). This can be an executive, or
	// someone who will have program-wide access to the cards that Lithic will provide.
	// In some cases, this individual could also be a beneficial owner listed above.
	ControlPerson AccountHolderControlPerson `json:"control_person"`
	// (Deprecated. Use control_person.email when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".) Primary email of
	// Account Holder.
	Email string `json:"email"`
	// The type of KYC exemption for a KYC-Exempt Account Holder.
	ExemptionType AccountHolderExemptionType `json:"exemption_type"`
	// Customer-provided token that indicates a relationship with an object outside of
	// the Lithic ecosystem.
	ExternalID string `json:"external_id" format:"string"`
	// Only present when user_type == "INDIVIDUAL". Information about the individual
	// for which the account is being opened and KYC is being run.
	Individual AccountHolderIndividual `json:"individual"`
	// Only present when user_type == "BUSINESS". User-submitted description of the
	// business.
	NatureOfBusiness string `json:"nature_of_business" format:"string"`
	// (Deprecated. Use control_person.phone_number when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".) Primary phone of
	// Account Holder, entered in E.164 format.
	PhoneNumber string `json:"phone_number"`
	// Only present for "KYB_BASIC" workflow. A list of documents required for the
	// account holder to be approved.
	RequiredDocuments []RequiredDocument `json:"required_documents"`
	// (Deprecated. Use verification_application.status instead)
	//
	// KYC and KYB evaluation states.
	//
	// Note:
	//
	// - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
	Status AccountHolderStatus `json:"status"`
	// (Deprecated. Use verification_application.status_reasons)
	//
	// Reason for the evaluation status.
	StatusReasons []AccountHolderStatusReason `json:"status_reasons"`
	// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
	// attribute will be present. If the type is "BUSINESS" then the "business_entity",
	// "control_person", "beneficial_owner_individuals", "nature_of_business", and
	// "website_url" attributes will be present.
	UserType AccountHolderUserType `json:"user_type"`
	// Information about the most recent identity verification attempt
	VerificationApplication AccountHolderVerificationApplication `json:"verification_application"`
	// Only present when user_type == "BUSINESS". Business's primary website.
	WebsiteURL string            `json:"website_url" format:"string"`
	JSON       accountHolderJSON `json:"-"`
}

// accountHolderJSON contains the JSON metadata for the struct [AccountHolder]
type accountHolderJSON struct {
	Token                      apijson.Field
	Created                    apijson.Field
	AccountToken               apijson.Field
	BeneficialOwnerEntities    apijson.Field
	BeneficialOwnerIndividuals apijson.Field
	BusinessAccountToken       apijson.Field
	BusinessEntity             apijson.Field
	ControlPerson              apijson.Field
	Email                      apijson.Field
	ExemptionType              apijson.Field
	ExternalID                 apijson.Field
	Individual                 apijson.Field
	NatureOfBusiness           apijson.Field
	PhoneNumber                apijson.Field
	RequiredDocuments          apijson.Field
	Status                     apijson.Field
	StatusReasons              apijson.Field
	UserType                   apijson.Field
	VerificationApplication    apijson.Field
	WebsiteURL                 apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderJSON) RawJSON() string {
	return r.raw
}

type AccountHolderBeneficialOwnerEntity struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address shared.Address `json:"address,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName string `json:"dba_business_name,required"`
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID string `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName string `json:"legal_business_name,required"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers []string `json:"phone_numbers,required"`
	// Parent company name (if applicable).
	ParentCompany string                                 `json:"parent_company"`
	JSON          accountHolderBeneficialOwnerEntityJSON `json:"-"`
}

// accountHolderBeneficialOwnerEntityJSON contains the JSON metadata for the struct
// [AccountHolderBeneficialOwnerEntity]
type accountHolderBeneficialOwnerEntityJSON struct {
	Address           apijson.Field
	DbaBusinessName   apijson.Field
	EntityToken       apijson.Field
	GovernmentID      apijson.Field
	LegalBusinessName apijson.Field
	PhoneNumbers      apijson.Field
	ParentCompany     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountHolderBeneficialOwnerEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderBeneficialOwnerEntityJSON) RawJSON() string {
	return r.raw
}

// Information about an individual associated with an account holder. A subset of
// the information provided via KYC. For example, we do not return the government
// id.
type AccountHolderBeneficialOwnerIndividual struct {
	// Individual's current address
	Address shared.Address `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob,required"`
	// Individual's email address.
	Email string `json:"email,required"`
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                                     `json:"phone_number,required"`
	JSON        accountHolderBeneficialOwnerIndividualJSON `json:"-"`
}

// accountHolderBeneficialOwnerIndividualJSON contains the JSON metadata for the
// struct [AccountHolderBeneficialOwnerIndividual]
type accountHolderBeneficialOwnerIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	EntityToken apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderBeneficialOwnerIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderBeneficialOwnerIndividualJSON) RawJSON() string {
	return r.raw
}

// Only present when user_type == "BUSINESS". Information about the business for
// which the account is being opened and KYB is being run.
type AccountHolderBusinessEntity struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address shared.Address `json:"address,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName string `json:"dba_business_name,required"`
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID string `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName string `json:"legal_business_name,required"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers []string `json:"phone_numbers,required"`
	// Parent company name (if applicable).
	ParentCompany string                          `json:"parent_company"`
	JSON          accountHolderBusinessEntityJSON `json:"-"`
}

// accountHolderBusinessEntityJSON contains the JSON metadata for the struct
// [AccountHolderBusinessEntity]
type accountHolderBusinessEntityJSON struct {
	Address           apijson.Field
	DbaBusinessName   apijson.Field
	EntityToken       apijson.Field
	GovernmentID      apijson.Field
	LegalBusinessName apijson.Field
	PhoneNumbers      apijson.Field
	ParentCompany     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountHolderBusinessEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderBusinessEntityJSON) RawJSON() string {
	return r.raw
}

// Only present when user_type == "BUSINESS". An individual with significant
// responsibility for managing the legal entity (e.g., a Chief Executive Officer,
// Chief Financial Officer, Chief Operating Officer, Managing Member, General
// Partner, President, Vice President, or Treasurer). This can be an executive, or
// someone who will have program-wide access to the cards that Lithic will provide.
// In some cases, this individual could also be a beneficial owner listed above.
type AccountHolderControlPerson struct {
	// Individual's current address
	Address shared.Address `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob,required"`
	// Individual's email address.
	Email string `json:"email,required"`
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                         `json:"phone_number,required"`
	JSON        accountHolderControlPersonJSON `json:"-"`
}

// accountHolderControlPersonJSON contains the JSON metadata for the struct
// [AccountHolderControlPerson]
type accountHolderControlPersonJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	EntityToken apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderControlPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderControlPersonJSON) RawJSON() string {
	return r.raw
}

// The type of KYC exemption for a KYC-Exempt Account Holder.
type AccountHolderExemptionType string

const (
	AccountHolderExemptionTypeAuthorizedUser  AccountHolderExemptionType = "AUTHORIZED_USER"
	AccountHolderExemptionTypePrepaidCardUser AccountHolderExemptionType = "PREPAID_CARD_USER"
)

func (r AccountHolderExemptionType) IsKnown() bool {
	switch r {
	case AccountHolderExemptionTypeAuthorizedUser, AccountHolderExemptionTypePrepaidCardUser:
		return true
	}
	return false
}

// Only present when user_type == "INDIVIDUAL". Information about the individual
// for which the account is being opened and KYC is being run.
type AccountHolderIndividual struct {
	// Individual's current address
	Address shared.Address `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob,required"`
	// Individual's email address.
	Email string `json:"email,required"`
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                      `json:"phone_number,required"`
	JSON        accountHolderIndividualJSON `json:"-"`
}

// accountHolderIndividualJSON contains the JSON metadata for the struct
// [AccountHolderIndividual]
type accountHolderIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	EntityToken apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderIndividualJSON) RawJSON() string {
	return r.raw
}

// (Deprecated. Use verification_application.status instead)
//
// KYC and KYB evaluation states.
//
// Note:
//
// - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
type AccountHolderStatus string

const (
	AccountHolderStatusAccepted        AccountHolderStatus = "ACCEPTED"
	AccountHolderStatusPendingReview   AccountHolderStatus = "PENDING_REVIEW"
	AccountHolderStatusPendingDocument AccountHolderStatus = "PENDING_DOCUMENT"
	AccountHolderStatusPendingResubmit AccountHolderStatus = "PENDING_RESUBMIT"
	AccountHolderStatusRejected        AccountHolderStatus = "REJECTED"
)

func (r AccountHolderStatus) IsKnown() bool {
	switch r {
	case AccountHolderStatusAccepted, AccountHolderStatusPendingReview, AccountHolderStatusPendingDocument, AccountHolderStatusPendingResubmit, AccountHolderStatusRejected:
		return true
	}
	return false
}

type AccountHolderStatusReason string

const (
	AccountHolderStatusReasonAddressVerificationFailure  AccountHolderStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderStatusReasonAgeThresholdFailure         AccountHolderStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderStatusReasonCompleteVerificationFailure AccountHolderStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderStatusReasonDobVerificationFailure      AccountHolderStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderStatusReasonIDVerificationFailure       AccountHolderStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderStatusReasonMaxDocumentAttempts         AccountHolderStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderStatusReasonMaxResubmissionAttempts     AccountHolderStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderStatusReasonNameVerificationFailure     AccountHolderStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderStatusReasonOtherVerificationFailure    AccountHolderStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderStatusReasonRiskThresholdFailure        AccountHolderStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderStatusReasonWatchlistAlertFailure       AccountHolderStatusReason = "WATCHLIST_ALERT_FAILURE"
)

func (r AccountHolderStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderStatusReasonAddressVerificationFailure, AccountHolderStatusReasonAgeThresholdFailure, AccountHolderStatusReasonCompleteVerificationFailure, AccountHolderStatusReasonDobVerificationFailure, AccountHolderStatusReasonIDVerificationFailure, AccountHolderStatusReasonMaxDocumentAttempts, AccountHolderStatusReasonMaxResubmissionAttempts, AccountHolderStatusReasonNameVerificationFailure, AccountHolderStatusReasonOtherVerificationFailure, AccountHolderStatusReasonRiskThresholdFailure, AccountHolderStatusReasonWatchlistAlertFailure:
		return true
	}
	return false
}

// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
// attribute will be present. If the type is "BUSINESS" then the "business_entity",
// "control_person", "beneficial_owner_individuals", "nature_of_business", and
// "website_url" attributes will be present.
type AccountHolderUserType string

const (
	AccountHolderUserTypeBusiness   AccountHolderUserType = "BUSINESS"
	AccountHolderUserTypeIndividual AccountHolderUserType = "INDIVIDUAL"
)

func (r AccountHolderUserType) IsKnown() bool {
	switch r {
	case AccountHolderUserTypeBusiness, AccountHolderUserTypeIndividual:
		return true
	}
	return false
}

// Information about the most recent identity verification attempt
type AccountHolderVerificationApplication struct {
	// Timestamp of when the application was created.
	Created time.Time `json:"created" format:"date-time"`
	// KYC and KYB evaluation states.
	//
	// Note:
	//
	// - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
	Status AccountHolderVerificationApplicationStatus `json:"status"`
	// Reason for the evaluation status.
	StatusReasons []AccountHolderVerificationApplicationStatusReason `json:"status_reasons"`
	// Timestamp of when the application was last updated.
	Updated time.Time                                `json:"updated" format:"date-time"`
	JSON    accountHolderVerificationApplicationJSON `json:"-"`
}

// accountHolderVerificationApplicationJSON contains the JSON metadata for the
// struct [AccountHolderVerificationApplication]
type accountHolderVerificationApplicationJSON struct {
	Created       apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	Updated       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountHolderVerificationApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderVerificationApplicationJSON) RawJSON() string {
	return r.raw
}

// KYC and KYB evaluation states.
//
// Note:
//
// - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
type AccountHolderVerificationApplicationStatus string

const (
	AccountHolderVerificationApplicationStatusAccepted        AccountHolderVerificationApplicationStatus = "ACCEPTED"
	AccountHolderVerificationApplicationStatusPendingReview   AccountHolderVerificationApplicationStatus = "PENDING_REVIEW"
	AccountHolderVerificationApplicationStatusPendingDocument AccountHolderVerificationApplicationStatus = "PENDING_DOCUMENT"
	AccountHolderVerificationApplicationStatusPendingResubmit AccountHolderVerificationApplicationStatus = "PENDING_RESUBMIT"
	AccountHolderVerificationApplicationStatusRejected        AccountHolderVerificationApplicationStatus = "REJECTED"
)

func (r AccountHolderVerificationApplicationStatus) IsKnown() bool {
	switch r {
	case AccountHolderVerificationApplicationStatusAccepted, AccountHolderVerificationApplicationStatusPendingReview, AccountHolderVerificationApplicationStatusPendingDocument, AccountHolderVerificationApplicationStatusPendingResubmit, AccountHolderVerificationApplicationStatusRejected:
		return true
	}
	return false
}

type AccountHolderVerificationApplicationStatusReason string

const (
	AccountHolderVerificationApplicationStatusReasonAddressVerificationFailure  AccountHolderVerificationApplicationStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderVerificationApplicationStatusReasonAgeThresholdFailure         AccountHolderVerificationApplicationStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderVerificationApplicationStatusReasonCompleteVerificationFailure AccountHolderVerificationApplicationStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderVerificationApplicationStatusReasonDobVerificationFailure      AccountHolderVerificationApplicationStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderVerificationApplicationStatusReasonIDVerificationFailure       AccountHolderVerificationApplicationStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderVerificationApplicationStatusReasonMaxDocumentAttempts         AccountHolderVerificationApplicationStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderVerificationApplicationStatusReasonMaxResubmissionAttempts     AccountHolderVerificationApplicationStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderVerificationApplicationStatusReasonNameVerificationFailure     AccountHolderVerificationApplicationStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderVerificationApplicationStatusReasonOtherVerificationFailure    AccountHolderVerificationApplicationStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderVerificationApplicationStatusReasonRiskThresholdFailure        AccountHolderVerificationApplicationStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderVerificationApplicationStatusReasonWatchlistAlertFailure       AccountHolderVerificationApplicationStatusReason = "WATCHLIST_ALERT_FAILURE"
)

func (r AccountHolderVerificationApplicationStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderVerificationApplicationStatusReasonAddressVerificationFailure, AccountHolderVerificationApplicationStatusReasonAgeThresholdFailure, AccountHolderVerificationApplicationStatusReasonCompleteVerificationFailure, AccountHolderVerificationApplicationStatusReasonDobVerificationFailure, AccountHolderVerificationApplicationStatusReasonIDVerificationFailure, AccountHolderVerificationApplicationStatusReasonMaxDocumentAttempts, AccountHolderVerificationApplicationStatusReasonMaxResubmissionAttempts, AccountHolderVerificationApplicationStatusReasonNameVerificationFailure, AccountHolderVerificationApplicationStatusReasonOtherVerificationFailure, AccountHolderVerificationApplicationStatusReasonRiskThresholdFailure, AccountHolderVerificationApplicationStatusReasonWatchlistAlertFailure:
		return true
	}
	return false
}

type AddressUpdateParam struct {
	// Valid deliverable address (no PO boxes).
	Address1 param.Field[string] `json:"address1"`
	// Unit or apartment number (if applicable).
	Address2 param.Field[string] `json:"address2"`
	// Name of city.
	City param.Field[string] `json:"city"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country param.Field[string] `json:"country"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode param.Field[string] `json:"postal_code"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State param.Field[string] `json:"state"`
}

func (r AddressUpdateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBParam struct {
	// You must submit a list of all direct and indirect individuals with 25% or more
	// ownership in the company. A maximum of 4 beneficial owners can be submitted. If
	// no individual owns 25% of the company you do not need to send beneficial owner
	// information. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included.
	BeneficialOwnerIndividuals param.Field[[]KYBBeneficialOwnerIndividualParam] `json:"beneficial_owner_individuals,required"`
	// Information for business for which the account is being opened and KYB is being
	// run.
	BusinessEntity param.Field[KYBBusinessEntityParam] `json:"business_entity,required"`
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer, Managing Member, General Partner, President, Vice President, or
	// Treasurer). This can be an executive, or someone who will have program-wide
	// access to the cards that Lithic will provide. In some cases, this individual
	// could also be a beneficial owner listed above. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section II) for more background.
	ControlPerson param.Field[KYBControlPersonParam] `json:"control_person,required"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness param.Field[string] `json:"nature_of_business,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp param.Field[string] `json:"tos_timestamp,required"`
	// Specifies the type of KYB workflow to run.
	Workflow param.Field[KYBWorkflow] `json:"workflow,required"`
	// Deprecated.
	//
	// Deprecated: deprecated
	BeneficialOwnerEntities param.Field[[]KYBBeneficialOwnerEntityParam] `json:"beneficial_owner_entities"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// An RFC 3339 timestamp indicating when precomputed KYB was completed on the
	// business with a pass result.
	//
	// This field is required only if workflow type is `KYB_BYO`.
	KYBPassedTimestamp param.Field[string] `json:"kyb_passed_timestamp"`
	// Company website URL.
	WebsiteURL param.Field[string] `json:"website_url"`
}

func (r KYBParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYBParam) implementsAccountHolderNewParamsBodyUnion() {}

// Individuals associated with a KYB application. Phone number is optional.
type KYBBeneficialOwnerIndividualParam struct {
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
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r KYBBeneficialOwnerIndividualParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information for business for which the account is being opened and KYB is being
// run.
type KYBBusinessEntityParam struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID param.Field[string] `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName param.Field[string] `json:"legal_business_name,required"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers param.Field[[]string] `json:"phone_numbers,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName param.Field[string] `json:"dba_business_name"`
	// Parent company name (if applicable).
	ParentCompany param.Field[string] `json:"parent_company"`
}

func (r KYBBusinessEntityParam) MarshalJSON() (data []byte, err error) {
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
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r KYBControlPersonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of KYB workflow to run.
type KYBWorkflow string

const (
	KYBWorkflowKYBBasic KYBWorkflow = "KYB_BASIC"
	KYBWorkflowKYBByo   KYBWorkflow = "KYB_BYO"
)

func (r KYBWorkflow) IsKnown() bool {
	switch r {
	case KYBWorkflowKYBBasic, KYBWorkflowKYBByo:
		return true
	}
	return false
}

type KYBBeneficialOwnerEntityParam struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID param.Field[string] `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName param.Field[string] `json:"legal_business_name,required"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers param.Field[[]string] `json:"phone_numbers,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName param.Field[string] `json:"dba_business_name"`
	// Parent company name (if applicable).
	ParentCompany param.Field[string] `json:"parent_company"`
}

func (r KYBBeneficialOwnerEntityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBBusinessEntity struct {
	// Business”s physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address KYBBusinessEntityAddress `json:"address,required"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID string `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName string `json:"legal_business_name,required"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers []string `json:"phone_numbers,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName string `json:"dba_business_name"`
	// Parent company name (if applicable).
	ParentCompany string                `json:"parent_company"`
	JSON          kybBusinessEntityJSON `json:"-"`
}

// kybBusinessEntityJSON contains the JSON metadata for the struct
// [KYBBusinessEntity]
type kybBusinessEntityJSON struct {
	Address           apijson.Field
	GovernmentID      apijson.Field
	LegalBusinessName apijson.Field
	PhoneNumbers      apijson.Field
	DbaBusinessName   apijson.Field
	ParentCompany     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *KYBBusinessEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kybBusinessEntityJSON) RawJSON() string {
	return r.raw
}

// Business”s physical address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable.
type KYBBusinessEntityAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                       `json:"address2"`
	JSON     kybBusinessEntityAddressJSON `json:"-"`
}

// kybBusinessEntityAddressJSON contains the JSON metadata for the struct
// [KYBBusinessEntityAddress]
type kybBusinessEntityAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KYBBusinessEntityAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kybBusinessEntityAddressJSON) RawJSON() string {
	return r.raw
}

type KYCParam struct {
	// Information on individual for whom the account is being opened and KYC is being
	// run.
	Individual param.Field[KYCIndividualParam] `json:"individual,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp param.Field[string] `json:"tos_timestamp,required"`
	// Specifies the type of KYC workflow to run.
	Workflow param.Field[KYCWorkflow] `json:"workflow,required"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// individual with a pass result.
	//
	// This field is required only if workflow type is `KYC_BYO`.
	KYCPassedTimestamp param.Field[string] `json:"kyc_passed_timestamp"`
}

func (r KYCParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYCParam) implementsAccountHolderNewParamsBodyUnion() {}

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

// Specifies the type of KYC workflow to run.
type KYCWorkflow string

const (
	KYCWorkflowKYCBasic KYCWorkflow = "KYC_BASIC"
	KYCWorkflowKYCByo   KYCWorkflow = "KYC_BYO"
)

func (r KYCWorkflow) IsKnown() bool {
	switch r {
	case KYCWorkflowKYCBasic, KYCWorkflowKYCByo:
		return true
	}
	return false
}

type KYCExemptParam struct {
	// KYC Exempt user's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// The KYC Exempt user's email
	Email param.Field[string] `json:"email,required"`
	// The KYC Exempt user's first name
	FirstName param.Field[string] `json:"first_name,required"`
	// Specifies the type of KYC Exempt user
	KYCExemptionType param.Field[KYCExemptKYCExemptionType] `json:"kyc_exemption_type,required"`
	// The KYC Exempt user's last name
	LastName param.Field[string] `json:"last_name,required"`
	// The KYC Exempt user's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number,required"`
	// Specifies the workflow type. This must be 'KYC_EXEMPT'
	Workflow param.Field[KYCExemptWorkflow] `json:"workflow,required"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken param.Field[string] `json:"business_account_token"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
}

func (r KYCExemptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYCExemptParam) implementsAccountHolderNewParamsBodyUnion() {}

// Specifies the type of KYC Exempt user
type KYCExemptKYCExemptionType string

const (
	KYCExemptKYCExemptionTypeAuthorizedUser  KYCExemptKYCExemptionType = "AUTHORIZED_USER"
	KYCExemptKYCExemptionTypePrepaidCardUser KYCExemptKYCExemptionType = "PREPAID_CARD_USER"
)

func (r KYCExemptKYCExemptionType) IsKnown() bool {
	switch r {
	case KYCExemptKYCExemptionTypeAuthorizedUser, KYCExemptKYCExemptionTypePrepaidCardUser:
		return true
	}
	return false
}

// Specifies the workflow type. This must be 'KYC_EXEMPT'
type KYCExemptWorkflow string

const (
	KYCExemptWorkflowKYCExempt KYCExemptWorkflow = "KYC_EXEMPT"
)

func (r KYCExemptWorkflow) IsKnown() bool {
	switch r {
	case KYCExemptWorkflowKYCExempt:
		return true
	}
	return false
}

type RequiredDocument struct {
	// Globally unique identifier for an entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// Provides the status reasons that will be satisfied by providing one of the valid
	// documents.
	StatusReasons []string `json:"status_reasons,required"`
	// A list of valid documents that will satisfy the KYC requirements for the
	// specified entity.
	ValidDocuments []string             `json:"valid_documents,required"`
	JSON           requiredDocumentJSON `json:"-"`
}

// requiredDocumentJSON contains the JSON metadata for the struct
// [RequiredDocument]
type requiredDocumentJSON struct {
	EntityToken    apijson.Field
	StatusReasons  apijson.Field
	ValidDocuments apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RequiredDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requiredDocumentJSON) RawJSON() string {
	return r.raw
}

type AccountHolderNewResponse struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token,required" format:"uuid"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token,required" format:"uuid"`
	// KYC and KYB evaluation states.
	//
	// Note:
	//
	// - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
	Status AccountHolderNewResponseStatus `json:"status,required"`
	// Reason for the evaluation status.
	StatusReasons []AccountHolderNewResponseStatusReason `json:"status_reasons,required"`
	// Timestamp of when the account holder was created.
	Created time.Time `json:"created" format:"date-time"`
	// Customer-provided token that indicates a relationship with an object outside of
	// the Lithic ecosystem.
	ExternalID string `json:"external_id" format:"string"`
	// Only present for "KYB_BASIC" workflow. A list of documents required for the
	// account holder to be approved.
	RequiredDocuments []RequiredDocument           `json:"required_documents"`
	JSON              accountHolderNewResponseJSON `json:"-"`
}

// accountHolderNewResponseJSON contains the JSON metadata for the struct
// [AccountHolderNewResponse]
type accountHolderNewResponseJSON struct {
	Token             apijson.Field
	AccountToken      apijson.Field
	Status            apijson.Field
	StatusReasons     apijson.Field
	Created           apijson.Field
	ExternalID        apijson.Field
	RequiredDocuments apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountHolderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderNewResponseJSON) RawJSON() string {
	return r.raw
}

// KYC and KYB evaluation states.
//
// Note:
//
// - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
type AccountHolderNewResponseStatus string

const (
	AccountHolderNewResponseStatusAccepted        AccountHolderNewResponseStatus = "ACCEPTED"
	AccountHolderNewResponseStatusPendingReview   AccountHolderNewResponseStatus = "PENDING_REVIEW"
	AccountHolderNewResponseStatusPendingDocument AccountHolderNewResponseStatus = "PENDING_DOCUMENT"
	AccountHolderNewResponseStatusPendingResubmit AccountHolderNewResponseStatus = "PENDING_RESUBMIT"
	AccountHolderNewResponseStatusRejected        AccountHolderNewResponseStatus = "REJECTED"
)

func (r AccountHolderNewResponseStatus) IsKnown() bool {
	switch r {
	case AccountHolderNewResponseStatusAccepted, AccountHolderNewResponseStatusPendingReview, AccountHolderNewResponseStatusPendingDocument, AccountHolderNewResponseStatusPendingResubmit, AccountHolderNewResponseStatusRejected:
		return true
	}
	return false
}

// Status Reasons for KYC/KYB enrollment states
type AccountHolderNewResponseStatusReason string

const (
	AccountHolderNewResponseStatusReasonAddressVerificationFailure                      AccountHolderNewResponseStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonAgeThresholdFailure                             AccountHolderNewResponseStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderNewResponseStatusReasonCompleteVerificationFailure                     AccountHolderNewResponseStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonDobVerificationFailure                          AccountHolderNewResponseStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonIDVerificationFailure                           AccountHolderNewResponseStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonMaxDocumentAttempts                             AccountHolderNewResponseStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderNewResponseStatusReasonMaxResubmissionAttempts                         AccountHolderNewResponseStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderNewResponseStatusReasonNameVerificationFailure                         AccountHolderNewResponseStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonOtherVerificationFailure                        AccountHolderNewResponseStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonRiskThresholdFailure                            AccountHolderNewResponseStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderNewResponseStatusReasonWatchlistAlertFailure                           AccountHolderNewResponseStatusReason = "WATCHLIST_ALERT_FAILURE"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntityIDVerificationFailure      AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_ID_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntityAddressVerificationFailure AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_ADDRESS_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntityNameVerificationFailure    AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_NAME_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_BUSINESS_OFFICERS_NOT_MATCHED"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntitySosFilingInactive          AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_FILING_INACTIVE"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntitySosNotMatched              AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_NOT_MATCHED"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntityCmraFailure                AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_CMRA_FAILURE"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntityWatchlistFailure           AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_WATCHLIST_FAILURE"
	AccountHolderNewResponseStatusReasonPrimaryBusinessEntityRegisteredAgentFailure     AccountHolderNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_REGISTERED_AGENT_FAILURE"
	AccountHolderNewResponseStatusReasonControlPersonBlocklistAlertFailure              AccountHolderNewResponseStatusReason = "CONTROL_PERSON_BLOCKLIST_ALERT_FAILURE"
	AccountHolderNewResponseStatusReasonControlPersonIDVerificationFailure              AccountHolderNewResponseStatusReason = "CONTROL_PERSON_ID_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonControlPersonDobVerificationFailure             AccountHolderNewResponseStatusReason = "CONTROL_PERSON_DOB_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonControlPersonNameVerificationFailure            AccountHolderNewResponseStatusReason = "CONTROL_PERSON_NAME_VERIFICATION_FAILURE"
)

func (r AccountHolderNewResponseStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderNewResponseStatusReasonAddressVerificationFailure, AccountHolderNewResponseStatusReasonAgeThresholdFailure, AccountHolderNewResponseStatusReasonCompleteVerificationFailure, AccountHolderNewResponseStatusReasonDobVerificationFailure, AccountHolderNewResponseStatusReasonIDVerificationFailure, AccountHolderNewResponseStatusReasonMaxDocumentAttempts, AccountHolderNewResponseStatusReasonMaxResubmissionAttempts, AccountHolderNewResponseStatusReasonNameVerificationFailure, AccountHolderNewResponseStatusReasonOtherVerificationFailure, AccountHolderNewResponseStatusReasonRiskThresholdFailure, AccountHolderNewResponseStatusReasonWatchlistAlertFailure, AccountHolderNewResponseStatusReasonPrimaryBusinessEntityIDVerificationFailure, AccountHolderNewResponseStatusReasonPrimaryBusinessEntityAddressVerificationFailure, AccountHolderNewResponseStatusReasonPrimaryBusinessEntityNameVerificationFailure, AccountHolderNewResponseStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched, AccountHolderNewResponseStatusReasonPrimaryBusinessEntitySosFilingInactive, AccountHolderNewResponseStatusReasonPrimaryBusinessEntitySosNotMatched, AccountHolderNewResponseStatusReasonPrimaryBusinessEntityCmraFailure, AccountHolderNewResponseStatusReasonPrimaryBusinessEntityWatchlistFailure, AccountHolderNewResponseStatusReasonPrimaryBusinessEntityRegisteredAgentFailure, AccountHolderNewResponseStatusReasonControlPersonBlocklistAlertFailure, AccountHolderNewResponseStatusReasonControlPersonIDVerificationFailure, AccountHolderNewResponseStatusReasonControlPersonDobVerificationFailure, AccountHolderNewResponseStatusReasonControlPersonNameVerificationFailure:
		return true
	}
	return false
}

type AccountHolderUpdateResponse struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token" format:"uuid"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token" format:"uuid"`
	// This field can have the runtime type of
	// [AccountHolderUpdateResponsePatchResponseAddress].
	Address interface{} `json:"address"`
	// This field can have the runtime type of [[]KYBBusinessEntity].
	BeneficialOwnerEntities interface{} `json:"beneficial_owner_entities"`
	// This field can have the runtime type of
	// [[]AccountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividual].
	BeneficialOwnerIndividuals interface{} `json:"beneficial_owner_individuals"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken string `json:"business_account_token" format:"uuid"`
	// Only present when user_type == "BUSINESS". Information about the business for
	// which the account is being opened and KYB is being run.
	BusinessEntity KYBBusinessEntity `json:"business_entity"`
	// This field can have the runtime type of
	// [AccountHolderUpdateResponseKYBKYCPatchResponseControlPerson].
	ControlPerson interface{} `json:"control_person"`
	// Timestamp of when the account holder was created.
	Created time.Time `json:"created" format:"date-time"`
	// (Deprecated. Use control_person.email when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".) Primary email of
	// Account Holder.
	Email string `json:"email"`
	// The type of KYC exemption for a KYC-Exempt Account Holder. "None" if the account
	// holder is not KYC-Exempt.
	ExemptionType AccountHolderUpdateResponseExemptionType `json:"exemption_type"`
	// Customer-provided token that indicates a relationship with an object outside of
	// the Lithic ecosystem.
	ExternalID string `json:"external_id" format:"string"`
	// The first name for the account holder
	FirstName string `json:"first_name"`
	// This field can have the runtime type of
	// [AccountHolderUpdateResponseKYBKYCPatchResponseIndividual].
	Individual interface{} `json:"individual"`
	// The last name for the account holder
	LastName string `json:"last_name"`
	// The legal business name for the account holder
	LegalBusinessName string `json:"legal_business_name"`
	// Only present when user_type == "BUSINESS". User-submitted description of the
	// business.
	NatureOfBusiness string `json:"nature_of_business" format:"string"`
	// (Deprecated. Use control_person.phone_number when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".) Primary phone of
	// Account Holder, entered in E.164 format.
	PhoneNumber string `json:"phone_number"`
	// This field can have the runtime type of [[]RequiredDocument].
	RequiredDocuments interface{} `json:"required_documents"`
	// (Deprecated. Use verification_application.status instead) KYC and KYB evaluation
	// states.
	//
	// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	// `ADVANCED` workflow.
	Status AccountHolderUpdateResponseStatus `json:"status"`
	// This field can have the runtime type of
	// [[]AccountHolderUpdateResponseKybkycPatchResponseStatusReason].
	StatusReasons interface{} `json:"status_reasons"`
	// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
	// attribute will be present.
	//
	// If the type is "BUSINESS" then the "business_entity", "control_person",
	// "beneficial_owner_individuals", "nature_of_business", and "website_url"
	// attributes will be present.
	UserType AccountHolderUpdateResponseUserType `json:"user_type"`
	// This field can have the runtime type of
	// [AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplication].
	VerificationApplication interface{} `json:"verification_application"`
	// Only present when user_type == "BUSINESS". Business's primary website.
	WebsiteURL string                          `json:"website_url" format:"string"`
	JSON       accountHolderUpdateResponseJSON `json:"-"`
	union      AccountHolderUpdateResponseUnion
}

// accountHolderUpdateResponseJSON contains the JSON metadata for the struct
// [AccountHolderUpdateResponse]
type accountHolderUpdateResponseJSON struct {
	Token                      apijson.Field
	AccountToken               apijson.Field
	Address                    apijson.Field
	BeneficialOwnerEntities    apijson.Field
	BeneficialOwnerIndividuals apijson.Field
	BusinessAccountToken       apijson.Field
	BusinessEntity             apijson.Field
	ControlPerson              apijson.Field
	Created                    apijson.Field
	Email                      apijson.Field
	ExemptionType              apijson.Field
	ExternalID                 apijson.Field
	FirstName                  apijson.Field
	Individual                 apijson.Field
	LastName                   apijson.Field
	LegalBusinessName          apijson.Field
	NatureOfBusiness           apijson.Field
	PhoneNumber                apijson.Field
	RequiredDocuments          apijson.Field
	Status                     apijson.Field
	StatusReasons              apijson.Field
	UserType                   apijson.Field
	VerificationApplication    apijson.Field
	WebsiteURL                 apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r accountHolderUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccountHolderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccountHolderUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountHolderUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountHolderUpdateResponseKYBKYCPatchResponse],
// [AccountHolderUpdateResponsePatchResponse].
func (r AccountHolderUpdateResponse) AsUnion() AccountHolderUpdateResponseUnion {
	return r.union
}

// Union satisfied by [AccountHolderUpdateResponseKYBKYCPatchResponse] or
// [AccountHolderUpdateResponsePatchResponse].
type AccountHolderUpdateResponseUnion interface {
	implementsAccountHolderUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountHolderUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountHolderUpdateResponseKYBKYCPatchResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountHolderUpdateResponsePatchResponse{}),
		},
	)
}

type AccountHolderUpdateResponseKYBKYCPatchResponse struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token" format:"uuid"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token" format:"uuid"`
	// Deprecated.
	BeneficialOwnerEntities []KYBBusinessEntity `json:"beneficial_owner_entities"`
	// Only present when user_type == "BUSINESS". You must submit a list of all direct
	// and indirect individuals with 25% or more ownership in the company. A maximum of
	// 4 beneficial owners can be submitted. If no individual owns 25% of the company
	// you do not need to send beneficial owner information. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included.
	BeneficialOwnerIndividuals []AccountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken string `json:"business_account_token" format:"uuid"`
	// Only present when user_type == "BUSINESS". Information about the business for
	// which the account is being opened and KYB is being run.
	BusinessEntity KYBBusinessEntity `json:"business_entity"`
	// Only present when user_type == "BUSINESS".
	//
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer,
	//
	// Managing Member, General Partner, President, Vice President, or Treasurer). This
	// can be an executive, or someone who will have program-wide access
	//
	// to the cards that Lithic will provide. In some cases, this individual could also
	// be a beneficial owner listed above.
	ControlPerson AccountHolderUpdateResponseKYBKYCPatchResponseControlPerson `json:"control_person"`
	// Timestamp of when the account holder was created.
	Created time.Time `json:"created" format:"date-time"`
	// (Deprecated. Use control_person.email when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".) Primary email of
	// Account Holder.
	Email string `json:"email"`
	// The type of KYC exemption for a KYC-Exempt Account Holder. "None" if the account
	// holder is not KYC-Exempt.
	ExemptionType AccountHolderUpdateResponseKYBKYCPatchResponseExemptionType `json:"exemption_type"`
	// Customer-provided token that indicates a relationship with an object outside of
	// the Lithic ecosystem.
	ExternalID string `json:"external_id" format:"string"`
	// Only present when user_type == "INDIVIDUAL". Information about the individual
	// for which the account is being opened and KYC is being run.
	Individual AccountHolderUpdateResponseKYBKYCPatchResponseIndividual `json:"individual"`
	// Only present when user_type == "BUSINESS". User-submitted description of the
	// business.
	NatureOfBusiness string `json:"nature_of_business" format:"string"`
	// (Deprecated. Use control_person.phone_number when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".) Primary phone of
	// Account Holder, entered in E.164 format.
	PhoneNumber string `json:"phone_number"`
	// Only present for "KYB_BASIC" and "KYC_ADVANCED" workflows. A list of documents
	// required for the account holder to be approved.
	RequiredDocuments []RequiredDocument `json:"required_documents"`
	// (Deprecated. Use verification_application.status instead) KYC and KYB evaluation
	// states.
	//
	// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	// `ADVANCED` workflow.
	Status AccountHolderUpdateResponseKYBKYCPatchResponseStatus `json:"status"`
	// (Deprecated. Use verification_application.status_reasons) Reason for the
	// evaluation status.
	StatusReasons []AccountHolderUpdateResponseKybkycPatchResponseStatusReason `json:"status_reasons"`
	// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
	// attribute will be present.
	//
	// If the type is "BUSINESS" then the "business_entity", "control_person",
	// "beneficial_owner_individuals", "nature_of_business", and "website_url"
	// attributes will be present.
	UserType AccountHolderUpdateResponseKYBKYCPatchResponseUserType `json:"user_type"`
	// Information about the most recent identity verification attempt
	VerificationApplication AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplication `json:"verification_application"`
	// Only present when user_type == "BUSINESS". Business's primary website.
	WebsiteURL string                                             `json:"website_url" format:"string"`
	JSON       accountHolderUpdateResponseKybkycPatchResponseJSON `json:"-"`
}

// accountHolderUpdateResponseKybkycPatchResponseJSON contains the JSON metadata
// for the struct [AccountHolderUpdateResponseKYBKYCPatchResponse]
type accountHolderUpdateResponseKybkycPatchResponseJSON struct {
	Token                      apijson.Field
	AccountToken               apijson.Field
	BeneficialOwnerEntities    apijson.Field
	BeneficialOwnerIndividuals apijson.Field
	BusinessAccountToken       apijson.Field
	BusinessEntity             apijson.Field
	ControlPerson              apijson.Field
	Created                    apijson.Field
	Email                      apijson.Field
	ExemptionType              apijson.Field
	ExternalID                 apijson.Field
	Individual                 apijson.Field
	NatureOfBusiness           apijson.Field
	PhoneNumber                apijson.Field
	RequiredDocuments          apijson.Field
	Status                     apijson.Field
	StatusReasons              apijson.Field
	UserType                   apijson.Field
	VerificationApplication    apijson.Field
	WebsiteURL                 apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountHolderUpdateResponseKYBKYCPatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseKybkycPatchResponseJSON) RawJSON() string {
	return r.raw
}

func (r AccountHolderUpdateResponseKYBKYCPatchResponse) implementsAccountHolderUpdateResponse() {}

type AccountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderUpdateResponseKYBKYCPatchResponseBeneficialOwnerIndividualsAddress `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email string `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                                                                      `json:"phone_number"`
	JSON        accountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividualJSON `json:"-"`
}

// accountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividualJSON
// contains the JSON metadata for the struct
// [AccountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividual]
type accountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividualJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderUpdateResponseKYBKYCPatchResponseBeneficialOwnerIndividualsAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                                                                              `json:"address2"`
	JSON     accountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividualsAddressJSON `json:"-"`
}

// accountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividualsAddressJSON
// contains the JSON metadata for the struct
// [AccountHolderUpdateResponseKYBKYCPatchResponseBeneficialOwnerIndividualsAddress]
type accountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividualsAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdateResponseKYBKYCPatchResponseBeneficialOwnerIndividualsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseKybkycPatchResponseBeneficialOwnerIndividualsAddressJSON) RawJSON() string {
	return r.raw
}

// Only present when user_type == "BUSINESS".
//
// An individual with significant responsibility for managing the legal entity
// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
// Officer,
//
// Managing Member, General Partner, President, Vice President, or Treasurer). This
// can be an executive, or someone who will have program-wide access
//
// to the cards that Lithic will provide. In some cases, this individual could also
// be a beneficial owner listed above.
type AccountHolderUpdateResponseKYBKYCPatchResponseControlPerson struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderUpdateResponseKYBKYCPatchResponseControlPersonAddress `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email string `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                                                          `json:"phone_number"`
	JSON        accountHolderUpdateResponseKybkycPatchResponseControlPersonJSON `json:"-"`
}

// accountHolderUpdateResponseKybkycPatchResponseControlPersonJSON contains the
// JSON metadata for the struct
// [AccountHolderUpdateResponseKYBKYCPatchResponseControlPerson]
type accountHolderUpdateResponseKybkycPatchResponseControlPersonJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdateResponseKYBKYCPatchResponseControlPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseKybkycPatchResponseControlPersonJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderUpdateResponseKYBKYCPatchResponseControlPersonAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                                                                 `json:"address2"`
	JSON     accountHolderUpdateResponseKybkycPatchResponseControlPersonAddressJSON `json:"-"`
}

// accountHolderUpdateResponseKybkycPatchResponseControlPersonAddressJSON contains
// the JSON metadata for the struct
// [AccountHolderUpdateResponseKYBKYCPatchResponseControlPersonAddress]
type accountHolderUpdateResponseKybkycPatchResponseControlPersonAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdateResponseKYBKYCPatchResponseControlPersonAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseKybkycPatchResponseControlPersonAddressJSON) RawJSON() string {
	return r.raw
}

// The type of KYC exemption for a KYC-Exempt Account Holder. "None" if the account
// holder is not KYC-Exempt.
type AccountHolderUpdateResponseKYBKYCPatchResponseExemptionType string

const (
	AccountHolderUpdateResponseKYBKYCPatchResponseExemptionTypeAuthorizedUser  AccountHolderUpdateResponseKYBKYCPatchResponseExemptionType = "AUTHORIZED_USER"
	AccountHolderUpdateResponseKYBKYCPatchResponseExemptionTypePrepaidCardUser AccountHolderUpdateResponseKYBKYCPatchResponseExemptionType = "PREPAID_CARD_USER"
)

func (r AccountHolderUpdateResponseKYBKYCPatchResponseExemptionType) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseKYBKYCPatchResponseExemptionTypeAuthorizedUser, AccountHolderUpdateResponseKYBKYCPatchResponseExemptionTypePrepaidCardUser:
		return true
	}
	return false
}

// Only present when user_type == "INDIVIDUAL". Information about the individual
// for which the account is being opened and KYC is being run.
type AccountHolderUpdateResponseKYBKYCPatchResponseIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderUpdateResponseKYBKYCPatchResponseIndividualAddress `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email string `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                                                       `json:"phone_number"`
	JSON        accountHolderUpdateResponseKybkycPatchResponseIndividualJSON `json:"-"`
}

// accountHolderUpdateResponseKybkycPatchResponseIndividualJSON contains the JSON
// metadata for the struct
// [AccountHolderUpdateResponseKYBKYCPatchResponseIndividual]
type accountHolderUpdateResponseKybkycPatchResponseIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdateResponseKYBKYCPatchResponseIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseKybkycPatchResponseIndividualJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderUpdateResponseKYBKYCPatchResponseIndividualAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                                                              `json:"address2"`
	JSON     accountHolderUpdateResponseKybkycPatchResponseIndividualAddressJSON `json:"-"`
}

// accountHolderUpdateResponseKybkycPatchResponseIndividualAddressJSON contains the
// JSON metadata for the struct
// [AccountHolderUpdateResponseKYBKYCPatchResponseIndividualAddress]
type accountHolderUpdateResponseKybkycPatchResponseIndividualAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdateResponseKYBKYCPatchResponseIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseKybkycPatchResponseIndividualAddressJSON) RawJSON() string {
	return r.raw
}

// (Deprecated. Use verification_application.status instead) KYC and KYB evaluation
// states.
//
// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
// `ADVANCED` workflow.
type AccountHolderUpdateResponseKYBKYCPatchResponseStatus string

const (
	AccountHolderUpdateResponseKYBKYCPatchResponseStatusAccepted        AccountHolderUpdateResponseKYBKYCPatchResponseStatus = "ACCEPTED"
	AccountHolderUpdateResponseKYBKYCPatchResponseStatusPendingDocument AccountHolderUpdateResponseKYBKYCPatchResponseStatus = "PENDING_DOCUMENT"
	AccountHolderUpdateResponseKYBKYCPatchResponseStatusPendingResubmit AccountHolderUpdateResponseKYBKYCPatchResponseStatus = "PENDING_RESUBMIT"
	AccountHolderUpdateResponseKYBKYCPatchResponseStatusRejected        AccountHolderUpdateResponseKYBKYCPatchResponseStatus = "REJECTED"
)

func (r AccountHolderUpdateResponseKYBKYCPatchResponseStatus) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseKYBKYCPatchResponseStatusAccepted, AccountHolderUpdateResponseKYBKYCPatchResponseStatusPendingDocument, AccountHolderUpdateResponseKYBKYCPatchResponseStatusPendingResubmit, AccountHolderUpdateResponseKYBKYCPatchResponseStatusRejected:
		return true
	}
	return false
}

// Status Reasons for KYC/KYB enrollment states
type AccountHolderUpdateResponseKybkycPatchResponseStatusReason string

const (
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonAddressVerificationFailure                      AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonAgeThresholdFailure                             AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonCompleteVerificationFailure                     AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonDobVerificationFailure                          AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonIDVerificationFailure                           AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonMaxDocumentAttempts                             AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonMaxResubmissionAttempts                         AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonNameVerificationFailure                         AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonOtherVerificationFailure                        AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonRiskThresholdFailure                            AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonWatchlistAlertFailure                           AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "WATCHLIST_ALERT_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityIDVerificationFailure      AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_ID_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityAddressVerificationFailure AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_ADDRESS_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityNameVerificationFailure    AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_NAME_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_BUSINESS_OFFICERS_NOT_MATCHED"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntitySosFilingInactive          AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_FILING_INACTIVE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntitySosNotMatched              AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_NOT_MATCHED"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityCmraFailure                AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_CMRA_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityWatchlistFailure           AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_WATCHLIST_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityRegisteredAgentFailure     AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_REGISTERED_AGENT_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonControlPersonBlocklistAlertFailure              AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "CONTROL_PERSON_BLOCKLIST_ALERT_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonControlPersonIDVerificationFailure              AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "CONTROL_PERSON_ID_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonControlPersonDobVerificationFailure             AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "CONTROL_PERSON_DOB_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseStatusReasonControlPersonNameVerificationFailure            AccountHolderUpdateResponseKybkycPatchResponseStatusReason = "CONTROL_PERSON_NAME_VERIFICATION_FAILURE"
)

func (r AccountHolderUpdateResponseKybkycPatchResponseStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseKybkycPatchResponseStatusReasonAddressVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonAgeThresholdFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonCompleteVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonDobVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonIDVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonMaxDocumentAttempts, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonMaxResubmissionAttempts, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonNameVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonOtherVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonRiskThresholdFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonWatchlistAlertFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityIDVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityAddressVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityNameVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntitySosFilingInactive, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntitySosNotMatched, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityCmraFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityWatchlistFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonPrimaryBusinessEntityRegisteredAgentFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonControlPersonBlocklistAlertFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonControlPersonIDVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonControlPersonDobVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseStatusReasonControlPersonNameVerificationFailure:
		return true
	}
	return false
}

// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
// attribute will be present.
//
// If the type is "BUSINESS" then the "business_entity", "control_person",
// "beneficial_owner_individuals", "nature_of_business", and "website_url"
// attributes will be present.
type AccountHolderUpdateResponseKYBKYCPatchResponseUserType string

const (
	AccountHolderUpdateResponseKYBKYCPatchResponseUserTypeBusiness   AccountHolderUpdateResponseKYBKYCPatchResponseUserType = "BUSINESS"
	AccountHolderUpdateResponseKYBKYCPatchResponseUserTypeIndividual AccountHolderUpdateResponseKYBKYCPatchResponseUserType = "INDIVIDUAL"
)

func (r AccountHolderUpdateResponseKYBKYCPatchResponseUserType) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseKYBKYCPatchResponseUserTypeBusiness, AccountHolderUpdateResponseKYBKYCPatchResponseUserTypeIndividual:
		return true
	}
	return false
}

// Information about the most recent identity verification attempt
type AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplication struct {
	// Timestamp of when the application was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// KYC and KYB evaluation states.
	//
	// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	// `ADVANCED` workflow.
	Status AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatus `json:"status,required"`
	// Reason for the evaluation status.
	StatusReasons []AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason `json:"status_reasons,required"`
	// Timestamp of when the application was last updated.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Timestamp of when the application passed the verification process. Only present
	// if `status` is `ACCEPTED`
	KyPassedAt time.Time                                                                 `json:"ky_passed_at" format:"date-time"`
	JSON       accountHolderUpdateResponseKybkycPatchResponseVerificationApplicationJSON `json:"-"`
}

// accountHolderUpdateResponseKybkycPatchResponseVerificationApplicationJSON
// contains the JSON metadata for the struct
// [AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplication]
type accountHolderUpdateResponseKybkycPatchResponseVerificationApplicationJSON struct {
	Created       apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	Updated       apijson.Field
	KyPassedAt    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseKybkycPatchResponseVerificationApplicationJSON) RawJSON() string {
	return r.raw
}

// KYC and KYB evaluation states.
//
// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
// `ADVANCED` workflow.
type AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatus string

const (
	AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatusAccepted        AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatus = "ACCEPTED"
	AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatusPendingDocument AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatus = "PENDING_DOCUMENT"
	AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatusPendingResubmit AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatus = "PENDING_RESUBMIT"
	AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatusRejected        AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatus = "REJECTED"
)

func (r AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatus) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatusAccepted, AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatusPendingDocument, AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatusPendingResubmit, AccountHolderUpdateResponseKYBKYCPatchResponseVerificationApplicationStatusRejected:
		return true
	}
	return false
}

// Status Reasons for KYC/KYB enrollment states
type AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason string

const (
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonAddressVerificationFailure                      AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonAgeThresholdFailure                             AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonCompleteVerificationFailure                     AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonDobVerificationFailure                          AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonIDVerificationFailure                           AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonMaxDocumentAttempts                             AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonMaxResubmissionAttempts                         AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonNameVerificationFailure                         AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonOtherVerificationFailure                        AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonRiskThresholdFailure                            AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonWatchlistAlertFailure                           AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "WATCHLIST_ALERT_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityIDVerificationFailure      AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_ID_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityAddressVerificationFailure AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_ADDRESS_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityNameVerificationFailure    AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_NAME_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_BUSINESS_OFFICERS_NOT_MATCHED"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntitySosFilingInactive          AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_FILING_INACTIVE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntitySosNotMatched              AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_NOT_MATCHED"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityCmraFailure                AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_CMRA_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityWatchlistFailure           AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_WATCHLIST_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityRegisteredAgentFailure     AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_REGISTERED_AGENT_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonControlPersonBlocklistAlertFailure              AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "CONTROL_PERSON_BLOCKLIST_ALERT_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonControlPersonIDVerificationFailure              AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "CONTROL_PERSON_ID_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonControlPersonDobVerificationFailure             AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "CONTROL_PERSON_DOB_VERIFICATION_FAILURE"
	AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonControlPersonNameVerificationFailure            AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason = "CONTROL_PERSON_NAME_VERIFICATION_FAILURE"
)

func (r AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonAddressVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonAgeThresholdFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonCompleteVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonDobVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonIDVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonMaxDocumentAttempts, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonMaxResubmissionAttempts, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonNameVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonOtherVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonRiskThresholdFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonWatchlistAlertFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityIDVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityAddressVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityNameVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntitySosFilingInactive, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntitySosNotMatched, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityCmraFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityWatchlistFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonPrimaryBusinessEntityRegisteredAgentFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonControlPersonBlocklistAlertFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonControlPersonIDVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonControlPersonDobVerificationFailure, AccountHolderUpdateResponseKybkycPatchResponseVerificationApplicationStatusReasonControlPersonNameVerificationFailure:
		return true
	}
	return false
}

type AccountHolderUpdateResponsePatchResponse struct {
	// The token for the account holder that was updated
	Token string `json:"token"`
	// The address for the account holder
	Address AccountHolderUpdateResponsePatchResponseAddress `json:"address"`
	// The token for the business account that the account holder is associated with
	BusinessAccountToken string `json:"business_account_token"`
	// The email for the account holder
	Email string `json:"email"`
	// The first name for the account holder
	FirstName string `json:"first_name"`
	// The last name for the account holder
	LastName string `json:"last_name"`
	// The legal business name for the account holder
	LegalBusinessName string `json:"legal_business_name"`
	// The phone_number for the account holder
	PhoneNumber string                                       `json:"phone_number"`
	JSON        accountHolderUpdateResponsePatchResponseJSON `json:"-"`
}

// accountHolderUpdateResponsePatchResponseJSON contains the JSON metadata for the
// struct [AccountHolderUpdateResponsePatchResponse]
type accountHolderUpdateResponsePatchResponseJSON struct {
	Token                apijson.Field
	Address              apijson.Field
	BusinessAccountToken apijson.Field
	Email                apijson.Field
	FirstName            apijson.Field
	LastName             apijson.Field
	LegalBusinessName    apijson.Field
	PhoneNumber          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountHolderUpdateResponsePatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponsePatchResponseJSON) RawJSON() string {
	return r.raw
}

func (r AccountHolderUpdateResponsePatchResponse) implementsAccountHolderUpdateResponse() {}

// The address for the account holder
type AccountHolderUpdateResponsePatchResponseAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                                              `json:"address2"`
	JSON     accountHolderUpdateResponsePatchResponseAddressJSON `json:"-"`
}

// accountHolderUpdateResponsePatchResponseAddressJSON contains the JSON metadata
// for the struct [AccountHolderUpdateResponsePatchResponseAddress]
type accountHolderUpdateResponsePatchResponseAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdateResponsePatchResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponsePatchResponseAddressJSON) RawJSON() string {
	return r.raw
}

// The type of KYC exemption for a KYC-Exempt Account Holder. "None" if the account
// holder is not KYC-Exempt.
type AccountHolderUpdateResponseExemptionType string

const (
	AccountHolderUpdateResponseExemptionTypeAuthorizedUser  AccountHolderUpdateResponseExemptionType = "AUTHORIZED_USER"
	AccountHolderUpdateResponseExemptionTypePrepaidCardUser AccountHolderUpdateResponseExemptionType = "PREPAID_CARD_USER"
)

func (r AccountHolderUpdateResponseExemptionType) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseExemptionTypeAuthorizedUser, AccountHolderUpdateResponseExemptionTypePrepaidCardUser:
		return true
	}
	return false
}

// (Deprecated. Use verification_application.status instead) KYC and KYB evaluation
// states.
//
// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
// `ADVANCED` workflow.
type AccountHolderUpdateResponseStatus string

const (
	AccountHolderUpdateResponseStatusAccepted        AccountHolderUpdateResponseStatus = "ACCEPTED"
	AccountHolderUpdateResponseStatusPendingDocument AccountHolderUpdateResponseStatus = "PENDING_DOCUMENT"
	AccountHolderUpdateResponseStatusPendingResubmit AccountHolderUpdateResponseStatus = "PENDING_RESUBMIT"
	AccountHolderUpdateResponseStatusRejected        AccountHolderUpdateResponseStatus = "REJECTED"
)

func (r AccountHolderUpdateResponseStatus) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseStatusAccepted, AccountHolderUpdateResponseStatusPendingDocument, AccountHolderUpdateResponseStatusPendingResubmit, AccountHolderUpdateResponseStatusRejected:
		return true
	}
	return false
}

// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
// attribute will be present.
//
// If the type is "BUSINESS" then the "business_entity", "control_person",
// "beneficial_owner_individuals", "nature_of_business", and "website_url"
// attributes will be present.
type AccountHolderUpdateResponseUserType string

const (
	AccountHolderUpdateResponseUserTypeBusiness   AccountHolderUpdateResponseUserType = "BUSINESS"
	AccountHolderUpdateResponseUserTypeIndividual AccountHolderUpdateResponseUserType = "INDIVIDUAL"
)

func (r AccountHolderUpdateResponseUserType) IsKnown() bool {
	switch r {
	case AccountHolderUpdateResponseUserTypeBusiness, AccountHolderUpdateResponseUserTypeIndividual:
		return true
	}
	return false
}

type AccountHolderListDocumentsResponse struct {
	Data []shared.Document                      `json:"data"`
	JSON accountHolderListDocumentsResponseJSON `json:"-"`
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

func (r accountHolderListDocumentsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHolderSimulateEnrollmentReviewResponse struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token" format:"uuid"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token" format:"uuid"`
	// Deprecated.
	BeneficialOwnerEntities []KYBBusinessEntity `json:"beneficial_owner_entities"`
	// Only present when user_type == "BUSINESS". You must submit a list of all direct
	// and indirect individuals with 25% or more ownership in the company. A maximum of
	// 4 beneficial owners can be submitted. If no individual owns 25% of the company
	// you do not need to send beneficial owner information. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included.
	BeneficialOwnerIndividuals []AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken string `json:"business_account_token" format:"uuid"`
	// Only present when user_type == "BUSINESS". Information about the business for
	// which the account is being opened and KYB is being run.
	BusinessEntity KYBBusinessEntity `json:"business_entity"`
	// Only present when user_type == "BUSINESS".
	//
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer,
	//
	// Managing Member, General Partner, President, Vice President, or Treasurer). This
	// can be an executive, or someone who will have program-wide access
	//
	// to the cards that Lithic will provide. In some cases, this individual could also
	// be a beneficial owner listed above.
	ControlPerson AccountHolderSimulateEnrollmentReviewResponseControlPerson `json:"control_person"`
	// Timestamp of when the account holder was created.
	Created time.Time `json:"created" format:"date-time"`
	// (Deprecated. Use control_person.email when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".) Primary email of
	// Account Holder.
	Email string `json:"email"`
	// The type of KYC exemption for a KYC-Exempt Account Holder. "None" if the account
	// holder is not KYC-Exempt.
	ExemptionType AccountHolderSimulateEnrollmentReviewResponseExemptionType `json:"exemption_type"`
	// Customer-provided token that indicates a relationship with an object outside of
	// the Lithic ecosystem.
	ExternalID string `json:"external_id" format:"string"`
	// Only present when user_type == "INDIVIDUAL". Information about the individual
	// for which the account is being opened and KYC is being run.
	Individual AccountHolderSimulateEnrollmentReviewResponseIndividual `json:"individual"`
	// Only present when user_type == "BUSINESS". User-submitted description of the
	// business.
	NatureOfBusiness string `json:"nature_of_business" format:"string"`
	// (Deprecated. Use control_person.phone_number when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".) Primary phone of
	// Account Holder, entered in E.164 format.
	PhoneNumber string `json:"phone_number"`
	// Only present for "KYB_BASIC" and "KYC_ADVANCED" workflows. A list of documents
	// required for the account holder to be approved.
	RequiredDocuments []RequiredDocument `json:"required_documents"`
	// (Deprecated. Use verification_application.status instead) KYC and KYB evaluation
	// states.
	//
	// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	// `ADVANCED` workflow.
	Status AccountHolderSimulateEnrollmentReviewResponseStatus `json:"status"`
	// (Deprecated. Use verification_application.status_reasons) Reason for the
	// evaluation status.
	StatusReasons []AccountHolderSimulateEnrollmentReviewResponseStatusReason `json:"status_reasons"`
	// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
	// attribute will be present.
	//
	// If the type is "BUSINESS" then the "business_entity", "control_person",
	// "beneficial_owner_individuals", "nature_of_business", and "website_url"
	// attributes will be present.
	UserType AccountHolderSimulateEnrollmentReviewResponseUserType `json:"user_type"`
	// Information about the most recent identity verification attempt
	VerificationApplication AccountHolderSimulateEnrollmentReviewResponseVerificationApplication `json:"verification_application"`
	// Only present when user_type == "BUSINESS". Business's primary website.
	WebsiteURL string                                            `json:"website_url" format:"string"`
	JSON       accountHolderSimulateEnrollmentReviewResponseJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseJSON contains the JSON metadata for
// the struct [AccountHolderSimulateEnrollmentReviewResponse]
type accountHolderSimulateEnrollmentReviewResponseJSON struct {
	Token                      apijson.Field
	AccountToken               apijson.Field
	BeneficialOwnerEntities    apijson.Field
	BeneficialOwnerIndividuals apijson.Field
	BusinessAccountToken       apijson.Field
	BusinessEntity             apijson.Field
	ControlPerson              apijson.Field
	Created                    apijson.Field
	Email                      apijson.Field
	ExemptionType              apijson.Field
	ExternalID                 apijson.Field
	Individual                 apijson.Field
	NatureOfBusiness           apijson.Field
	PhoneNumber                apijson.Field
	RequiredDocuments          apijson.Field
	Status                     apijson.Field
	StatusReasons              apijson.Field
	UserType                   apijson.Field
	VerificationApplication    apijson.Field
	WebsiteURL                 apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualsAddress `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email string `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                                                                     `json:"phone_number"`
	JSON        accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualJSON
// contains the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividual]
type accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualsAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                                                                             `json:"address2"`
	JSON     accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualsAddressJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualsAddressJSON
// contains the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualsAddress]
type accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualsAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividualsAddressJSON) RawJSON() string {
	return r.raw
}

// Only present when user_type == "BUSINESS".
//
// An individual with significant responsibility for managing the legal entity
// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
// Officer,
//
// Managing Member, General Partner, President, Vice President, or Treasurer). This
// can be an executive, or someone who will have program-wide access
//
// to the cards that Lithic will provide. In some cases, this individual could also
// be a beneficial owner listed above.
type AccountHolderSimulateEnrollmentReviewResponseControlPerson struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderSimulateEnrollmentReviewResponseControlPersonAddress `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email string `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                                                         `json:"phone_number"`
	JSON        accountHolderSimulateEnrollmentReviewResponseControlPersonJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseControlPersonJSON contains the JSON
// metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseControlPerson]
type accountHolderSimulateEnrollmentReviewResponseControlPersonJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseControlPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseControlPersonJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderSimulateEnrollmentReviewResponseControlPersonAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                                                                `json:"address2"`
	JSON     accountHolderSimulateEnrollmentReviewResponseControlPersonAddressJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseControlPersonAddressJSON contains
// the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseControlPersonAddress]
type accountHolderSimulateEnrollmentReviewResponseControlPersonAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseControlPersonAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseControlPersonAddressJSON) RawJSON() string {
	return r.raw
}

// The type of KYC exemption for a KYC-Exempt Account Holder. "None" if the account
// holder is not KYC-Exempt.
type AccountHolderSimulateEnrollmentReviewResponseExemptionType string

const (
	AccountHolderSimulateEnrollmentReviewResponseExemptionTypeAuthorizedUser  AccountHolderSimulateEnrollmentReviewResponseExemptionType = "AUTHORIZED_USER"
	AccountHolderSimulateEnrollmentReviewResponseExemptionTypePrepaidCardUser AccountHolderSimulateEnrollmentReviewResponseExemptionType = "PREPAID_CARD_USER"
)

func (r AccountHolderSimulateEnrollmentReviewResponseExemptionType) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewResponseExemptionTypeAuthorizedUser, AccountHolderSimulateEnrollmentReviewResponseExemptionTypePrepaidCardUser:
		return true
	}
	return false
}

// Only present when user_type == "INDIVIDUAL". Information about the individual
// for which the account is being opened and KYC is being run.
type AccountHolderSimulateEnrollmentReviewResponseIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderSimulateEnrollmentReviewResponseIndividualAddress `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email string `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                                                      `json:"phone_number"`
	JSON        accountHolderSimulateEnrollmentReviewResponseIndividualJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseIndividualJSON contains the JSON
// metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseIndividual]
type accountHolderSimulateEnrollmentReviewResponseIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseIndividualJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderSimulateEnrollmentReviewResponseIndividualAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                                                             `json:"address2"`
	JSON     accountHolderSimulateEnrollmentReviewResponseIndividualAddressJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseIndividualAddressJSON contains the
// JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseIndividualAddress]
type accountHolderSimulateEnrollmentReviewResponseIndividualAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseIndividualAddressJSON) RawJSON() string {
	return r.raw
}

// (Deprecated. Use verification_application.status instead) KYC and KYB evaluation
// states.
//
// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
// `ADVANCED` workflow.
type AccountHolderSimulateEnrollmentReviewResponseStatus string

const (
	AccountHolderSimulateEnrollmentReviewResponseStatusAccepted        AccountHolderSimulateEnrollmentReviewResponseStatus = "ACCEPTED"
	AccountHolderSimulateEnrollmentReviewResponseStatusPendingDocument AccountHolderSimulateEnrollmentReviewResponseStatus = "PENDING_DOCUMENT"
	AccountHolderSimulateEnrollmentReviewResponseStatusPendingResubmit AccountHolderSimulateEnrollmentReviewResponseStatus = "PENDING_RESUBMIT"
	AccountHolderSimulateEnrollmentReviewResponseStatusRejected        AccountHolderSimulateEnrollmentReviewResponseStatus = "REJECTED"
)

func (r AccountHolderSimulateEnrollmentReviewResponseStatus) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewResponseStatusAccepted, AccountHolderSimulateEnrollmentReviewResponseStatusPendingDocument, AccountHolderSimulateEnrollmentReviewResponseStatusPendingResubmit, AccountHolderSimulateEnrollmentReviewResponseStatusRejected:
		return true
	}
	return false
}

// Status Reasons for KYC/KYB enrollment states
type AccountHolderSimulateEnrollmentReviewResponseStatusReason string

const (
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonAddressVerificationFailure                      AccountHolderSimulateEnrollmentReviewResponseStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonAgeThresholdFailure                             AccountHolderSimulateEnrollmentReviewResponseStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonCompleteVerificationFailure                     AccountHolderSimulateEnrollmentReviewResponseStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonDobVerificationFailure                          AccountHolderSimulateEnrollmentReviewResponseStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonIDVerificationFailure                           AccountHolderSimulateEnrollmentReviewResponseStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonMaxDocumentAttempts                             AccountHolderSimulateEnrollmentReviewResponseStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonMaxResubmissionAttempts                         AccountHolderSimulateEnrollmentReviewResponseStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonNameVerificationFailure                         AccountHolderSimulateEnrollmentReviewResponseStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonOtherVerificationFailure                        AccountHolderSimulateEnrollmentReviewResponseStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonRiskThresholdFailure                            AccountHolderSimulateEnrollmentReviewResponseStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonWatchlistAlertFailure                           AccountHolderSimulateEnrollmentReviewResponseStatusReason = "WATCHLIST_ALERT_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityIDVerificationFailure      AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityAddressVerificationFailure AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_ADDRESS_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityNameVerificationFailure    AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_NAME_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_BUSINESS_OFFICERS_NOT_MATCHED"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntitySosFilingInactive          AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_FILING_INACTIVE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntitySosNotMatched              AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_NOT_MATCHED"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityCmraFailure                AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_CMRA_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityWatchlistFailure           AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_WATCHLIST_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityRegisteredAgentFailure     AccountHolderSimulateEnrollmentReviewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_REGISTERED_AGENT_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonControlPersonBlocklistAlertFailure              AccountHolderSimulateEnrollmentReviewResponseStatusReason = "CONTROL_PERSON_BLOCKLIST_ALERT_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonControlPersonIDVerificationFailure              AccountHolderSimulateEnrollmentReviewResponseStatusReason = "CONTROL_PERSON_ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonControlPersonDobVerificationFailure             AccountHolderSimulateEnrollmentReviewResponseStatusReason = "CONTROL_PERSON_DOB_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonControlPersonNameVerificationFailure            AccountHolderSimulateEnrollmentReviewResponseStatusReason = "CONTROL_PERSON_NAME_VERIFICATION_FAILURE"
)

func (r AccountHolderSimulateEnrollmentReviewResponseStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewResponseStatusReasonAddressVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonAgeThresholdFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonCompleteVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonDobVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonIDVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonMaxDocumentAttempts, AccountHolderSimulateEnrollmentReviewResponseStatusReasonMaxResubmissionAttempts, AccountHolderSimulateEnrollmentReviewResponseStatusReasonNameVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonOtherVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonRiskThresholdFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonWatchlistAlertFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityIDVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityAddressVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityNameVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntitySosFilingInactive, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntitySosNotMatched, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityCmraFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityWatchlistFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonPrimaryBusinessEntityRegisteredAgentFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonControlPersonBlocklistAlertFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonControlPersonIDVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonControlPersonDobVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonControlPersonNameVerificationFailure:
		return true
	}
	return false
}

// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
// attribute will be present.
//
// If the type is "BUSINESS" then the "business_entity", "control_person",
// "beneficial_owner_individuals", "nature_of_business", and "website_url"
// attributes will be present.
type AccountHolderSimulateEnrollmentReviewResponseUserType string

const (
	AccountHolderSimulateEnrollmentReviewResponseUserTypeBusiness   AccountHolderSimulateEnrollmentReviewResponseUserType = "BUSINESS"
	AccountHolderSimulateEnrollmentReviewResponseUserTypeIndividual AccountHolderSimulateEnrollmentReviewResponseUserType = "INDIVIDUAL"
)

func (r AccountHolderSimulateEnrollmentReviewResponseUserType) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewResponseUserTypeBusiness, AccountHolderSimulateEnrollmentReviewResponseUserTypeIndividual:
		return true
	}
	return false
}

// Information about the most recent identity verification attempt
type AccountHolderSimulateEnrollmentReviewResponseVerificationApplication struct {
	// Timestamp of when the application was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// KYC and KYB evaluation states.
	//
	// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	// `ADVANCED` workflow.
	Status AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatus `json:"status,required"`
	// Reason for the evaluation status.
	StatusReasons []AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason `json:"status_reasons,required"`
	// Timestamp of when the application was last updated.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Timestamp of when the application passed the verification process. Only present
	// if `status` is `ACCEPTED`
	KyPassedAt time.Time                                                                `json:"ky_passed_at" format:"date-time"`
	JSON       accountHolderSimulateEnrollmentReviewResponseVerificationApplicationJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseVerificationApplicationJSON
// contains the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseVerificationApplication]
type accountHolderSimulateEnrollmentReviewResponseVerificationApplicationJSON struct {
	Created       apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	Updated       apijson.Field
	KyPassedAt    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseVerificationApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseVerificationApplicationJSON) RawJSON() string {
	return r.raw
}

// KYC and KYB evaluation states.
//
// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
// `ADVANCED` workflow.
type AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatus string

const (
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusAccepted        AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatus = "ACCEPTED"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusPendingDocument AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatus = "PENDING_DOCUMENT"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusPendingResubmit AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatus = "PENDING_RESUBMIT"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusRejected        AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatus = "REJECTED"
)

func (r AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatus) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusAccepted, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusPendingDocument, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusPendingResubmit, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusRejected:
		return true
	}
	return false
}

// Status Reasons for KYC/KYB enrollment states
type AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason string

const (
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonAddressVerificationFailure                      AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonAgeThresholdFailure                             AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonCompleteVerificationFailure                     AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonDobVerificationFailure                          AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonIDVerificationFailure                           AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonMaxDocumentAttempts                             AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonMaxResubmissionAttempts                         AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonNameVerificationFailure                         AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonOtherVerificationFailure                        AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonRiskThresholdFailure                            AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonWatchlistAlertFailure                           AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "WATCHLIST_ALERT_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityIDVerificationFailure      AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityAddressVerificationFailure AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_ADDRESS_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityNameVerificationFailure    AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_NAME_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_BUSINESS_OFFICERS_NOT_MATCHED"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntitySosFilingInactive          AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_FILING_INACTIVE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntitySosNotMatched              AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_NOT_MATCHED"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityCmraFailure                AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_CMRA_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityWatchlistFailure           AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_WATCHLIST_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityRegisteredAgentFailure     AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "PRIMARY_BUSINESS_ENTITY_REGISTERED_AGENT_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonControlPersonBlocklistAlertFailure              AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "CONTROL_PERSON_BLOCKLIST_ALERT_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonControlPersonIDVerificationFailure              AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "CONTROL_PERSON_ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonControlPersonDobVerificationFailure             AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "CONTROL_PERSON_DOB_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonControlPersonNameVerificationFailure            AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "CONTROL_PERSON_NAME_VERIFICATION_FAILURE"
)

func (r AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonAddressVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonAgeThresholdFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonCompleteVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonDobVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonIDVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonMaxDocumentAttempts, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonMaxResubmissionAttempts, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonNameVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonOtherVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonRiskThresholdFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonWatchlistAlertFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityIDVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityAddressVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityNameVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntitySosFilingInactive, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntitySosNotMatched, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityCmraFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityWatchlistFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonPrimaryBusinessEntityRegisteredAgentFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonControlPersonBlocklistAlertFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonControlPersonIDVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonControlPersonDobVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonControlPersonNameVerificationFailure:
		return true
	}
	return false
}

type AccountHolderNewParams struct {
	Body AccountHolderNewParamsBodyUnion `json:"body,required"`
}

func (r AccountHolderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountHolderNewParamsBody struct {
	// KYC Exempt user's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address                    param.Field[shared.AddressParam] `json:"address"`
	BeneficialOwnerEntities    param.Field[interface{}]         `json:"beneficial_owner_entities"`
	BeneficialOwnerIndividuals param.Field[interface{}]         `json:"beneficial_owner_individuals"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken param.Field[string]      `json:"business_account_token"`
	BusinessEntity       param.Field[interface{}] `json:"business_entity"`
	ControlPerson        param.Field[interface{}] `json:"control_person"`
	// The KYC Exempt user's email
	Email param.Field[string] `json:"email"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// The KYC Exempt user's first name
	FirstName  param.Field[string]      `json:"first_name"`
	Individual param.Field[interface{}] `json:"individual"`
	// An RFC 3339 timestamp indicating when precomputed KYB was completed on the
	// business with a pass result.
	//
	// This field is required only if workflow type is `KYB_BYO`.
	KYBPassedTimestamp param.Field[string] `json:"kyb_passed_timestamp"`
	// Specifies the type of KYC Exempt user
	KYCExemptionType param.Field[AccountHolderNewParamsBodyKYCExemptionType] `json:"kyc_exemption_type"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// individual with a pass result.
	//
	// This field is required only if workflow type is `KYC_BYO`.
	KYCPassedTimestamp param.Field[string] `json:"kyc_passed_timestamp"`
	// The KYC Exempt user's last name
	LastName param.Field[string] `json:"last_name"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness param.Field[string] `json:"nature_of_business"`
	// The KYC Exempt user's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp param.Field[string] `json:"tos_timestamp"`
	// Company website URL.
	WebsiteURL param.Field[string] `json:"website_url"`
	// Specifies the type of KYB workflow to run.
	Workflow param.Field[AccountHolderNewParamsBodyWorkflow] `json:"workflow"`
}

func (r AccountHolderNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountHolderNewParamsBody) implementsAccountHolderNewParamsBodyUnion() {}

// Satisfied by [KYBParam], [AccountHolderNewParamsBodyKYBDelegated], [KYCParam],
// [KYCExemptParam], [AccountHolderNewParamsBody].
type AccountHolderNewParamsBodyUnion interface {
	implementsAccountHolderNewParamsBodyUnion()
}

type AccountHolderNewParamsBodyKYBDelegated struct {
	// Information for business for which the account is being opened.
	BusinessEntity param.Field[AccountHolderNewParamsBodyKYBDelegatedBusinessEntity] `json:"business_entity,required"`
	// You can submit a list of all direct and indirect individuals with 25% or more
	// ownership in the company. A maximum of 4 beneficial owners can be submitted. If
	// no individual owns 25% of the company you do not need to send beneficial owner
	// information. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included.
	BeneficialOwnerIndividuals param.Field[[]AccountHolderNewParamsBodyKYBDelegatedBeneficialOwnerIndividual] `json:"beneficial_owner_individuals"`
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer, Managing Member, General Partner, President, Vice President, or
	// Treasurer). This can be an executive, or someone who will have program-wide
	// access to the cards that Lithic will provide. In some cases, this individual
	// could also be a beneficial owner listed above. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section II) for more background.
	ControlPerson param.Field[AccountHolderNewParamsBodyKYBDelegatedControlPerson] `json:"control_person"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness param.Field[string] `json:"nature_of_business"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp param.Field[string] `json:"tos_timestamp"`
	// Company website URL.
	WebsiteURL param.Field[string] `json:"website_url"`
	// Specifies the type of KYB workflow to run.
	Workflow param.Field[AccountHolderNewParamsBodyKYBDelegatedWorkflow] `json:"workflow"`
}

func (r AccountHolderNewParamsBodyKYBDelegated) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountHolderNewParamsBodyKYBDelegated) implementsAccountHolderNewParamsBodyUnion() {}

// Information for business for which the account is being opened.
type AccountHolderNewParamsBodyKYBDelegatedBusinessEntity struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// Legal (formal) business name.
	LegalBusinessName param.Field[string] `json:"legal_business_name,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName param.Field[string] `json:"dba_business_name"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID param.Field[string] `json:"government_id"`
	// Parent company name (if applicable).
	ParentCompany param.Field[string] `json:"parent_company"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers param.Field[[]string] `json:"phone_numbers"`
}

func (r AccountHolderNewParamsBodyKYBDelegatedBusinessEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Individuals associated with a KYB application. Phone number is optional.
type AccountHolderNewParamsBodyKYBDelegatedBeneficialOwnerIndividual struct {
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
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r AccountHolderNewParamsBodyKYBDelegatedBeneficialOwnerIndividual) MarshalJSON() (data []byte, err error) {
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
type AccountHolderNewParamsBodyKYBDelegatedControlPerson struct {
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
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r AccountHolderNewParamsBodyKYBDelegatedControlPerson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of KYB workflow to run.
type AccountHolderNewParamsBodyKYBDelegatedWorkflow string

const (
	AccountHolderNewParamsBodyKYBDelegatedWorkflowKYBDelegated AccountHolderNewParamsBodyKYBDelegatedWorkflow = "KYB_DELEGATED"
)

func (r AccountHolderNewParamsBodyKYBDelegatedWorkflow) IsKnown() bool {
	switch r {
	case AccountHolderNewParamsBodyKYBDelegatedWorkflowKYBDelegated:
		return true
	}
	return false
}

// Specifies the type of KYC Exempt user
type AccountHolderNewParamsBodyKYCExemptionType string

const (
	AccountHolderNewParamsBodyKYCExemptionTypeAuthorizedUser  AccountHolderNewParamsBodyKYCExemptionType = "AUTHORIZED_USER"
	AccountHolderNewParamsBodyKYCExemptionTypePrepaidCardUser AccountHolderNewParamsBodyKYCExemptionType = "PREPAID_CARD_USER"
)

func (r AccountHolderNewParamsBodyKYCExemptionType) IsKnown() bool {
	switch r {
	case AccountHolderNewParamsBodyKYCExemptionTypeAuthorizedUser, AccountHolderNewParamsBodyKYCExemptionTypePrepaidCardUser:
		return true
	}
	return false
}

// Specifies the type of KYB workflow to run.
type AccountHolderNewParamsBodyWorkflow string

const (
	AccountHolderNewParamsBodyWorkflowKYBBasic     AccountHolderNewParamsBodyWorkflow = "KYB_BASIC"
	AccountHolderNewParamsBodyWorkflowKYBByo       AccountHolderNewParamsBodyWorkflow = "KYB_BYO"
	AccountHolderNewParamsBodyWorkflowKYBDelegated AccountHolderNewParamsBodyWorkflow = "KYB_DELEGATED"
	AccountHolderNewParamsBodyWorkflowKYCBasic     AccountHolderNewParamsBodyWorkflow = "KYC_BASIC"
	AccountHolderNewParamsBodyWorkflowKYCByo       AccountHolderNewParamsBodyWorkflow = "KYC_BYO"
	AccountHolderNewParamsBodyWorkflowKYCExempt    AccountHolderNewParamsBodyWorkflow = "KYC_EXEMPT"
)

func (r AccountHolderNewParamsBodyWorkflow) IsKnown() bool {
	switch r {
	case AccountHolderNewParamsBodyWorkflowKYBBasic, AccountHolderNewParamsBodyWorkflowKYBByo, AccountHolderNewParamsBodyWorkflowKYBDelegated, AccountHolderNewParamsBodyWorkflowKYCBasic, AccountHolderNewParamsBodyWorkflowKYCByo, AccountHolderNewParamsBodyWorkflowKYCExempt:
		return true
	}
	return false
}

type AccountHolderUpdateParams struct {
	// The KYB request payload for updating a business.
	Body AccountHolderUpdateParamsBodyUnion `json:"body,required"`
}

func (r AccountHolderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// The KYB request payload for updating a business.
type AccountHolderUpdateParamsBody struct {
	// Allowed for: KYC-Exempt, BYO-KYC, BYO-KYB.
	Address                    param.Field[AddressUpdateParam] `json:"address"`
	BeneficialOwnerEntities    param.Field[interface{}]        `json:"beneficial_owner_entities"`
	BeneficialOwnerIndividuals param.Field[interface{}]        `json:"beneficial_owner_individuals"`
	// Allowed for: KYC-Exempt, BYO-KYC. The token of the business account to which the
	// account holder is associated.
	BusinessAccountToken param.Field[string]      `json:"business_account_token"`
	BusinessEntity       param.Field[interface{}] `json:"business_entity"`
	ControlPerson        param.Field[interface{}] `json:"control_person"`
	// Allowed for all Account Holders. Account holder's email address. The primary
	// purpose of this field is for cardholder identification and verification during
	// the digital wallet tokenization process.
	Email param.Field[string] `json:"email"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// Allowed for KYC-Exempt, BYO-KYC. Account holder's first name.
	FirstName  param.Field[string]      `json:"first_name"`
	Individual param.Field[interface{}] `json:"individual"`
	// Allowed for KYC-Exempt, BYO-KYC. Account holder's last name.
	LastName param.Field[string] `json:"last_name"`
	// Allowed for BYO-KYB. Legal business name of the account holder.
	LegalBusinessName param.Field[string] `json:"legal_business_name"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness param.Field[string] `json:"nature_of_business"`
	// Allowed for all Account Holders. Account holder's phone number, entered in E.164
	// format. The primary purpose of this field is for cardholder identification and
	// verification during the digital wallet tokenization process.
	PhoneNumber param.Field[string] `json:"phone_number"`
	// Company website URL.
	WebsiteURL param.Field[string] `json:"website_url"`
}

func (r AccountHolderUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountHolderUpdateParamsBody) implementsAccountHolderUpdateParamsBodyUnion() {}

// The KYB request payload for updating a business.
//
// Satisfied by [AccountHolderUpdateParamsBodyKYBPatchRequest],
// [AccountHolderUpdateParamsBodyKYCPatchRequest],
// [AccountHolderUpdateParamsBodyPatchRequest], [AccountHolderUpdateParamsBody].
type AccountHolderUpdateParamsBodyUnion interface {
	implementsAccountHolderUpdateParamsBodyUnion()
}

// The KYB request payload for updating a business.
type AccountHolderUpdateParamsBodyKYBPatchRequest struct {
	// Deprecated.
	//
	// Deprecated: deprecated
	BeneficialOwnerEntities param.Field[[]AccountHolderUpdateParamsBodyKYBPatchRequestBeneficialOwnerEntity] `json:"beneficial_owner_entities"`
	// You must submit a list of all direct and indirect individuals with 25% or more
	// ownership in the company. A maximum of 4 beneficial owners can be submitted. If
	// no individual owns 25% of the company you do not need to send beneficial owner
	// information. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included.
	BeneficialOwnerIndividuals param.Field[[]AccountHolderUpdateParamsBodyKYBPatchRequestBeneficialOwnerIndividual] `json:"beneficial_owner_individuals"`
	// Information for business for which the account is being opened and KYB is being
	// run.
	BusinessEntity param.Field[AccountHolderUpdateParamsBodyKYBPatchRequestBusinessEntity] `json:"business_entity"`
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer, Managing Member, General Partner, President, Vice President, or
	// Treasurer). This can be an executive, or someone who will have program-wide
	// access to the cards that Lithic will provide. In some cases, this individual
	// could also be a beneficial owner listed above. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section II) for more background.
	ControlPerson param.Field[AccountHolderUpdateParamsBodyKYBPatchRequestControlPerson] `json:"control_person"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness param.Field[string] `json:"nature_of_business"`
	// Company website URL.
	WebsiteURL param.Field[string] `json:"website_url"`
}

func (r AccountHolderUpdateParamsBodyKYBPatchRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountHolderUpdateParamsBodyKYBPatchRequest) implementsAccountHolderUpdateParamsBodyUnion() {
}

type AccountHolderUpdateParamsBodyKYBPatchRequestBeneficialOwnerEntity struct {
	// Globally unique identifier for an entity.
	EntityToken param.Field[string] `json:"entity_token,required" format:"uuid"`
	// Business”s physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address param.Field[AddressUpdateParam] `json:"address"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName param.Field[string] `json:"dba_business_name"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID param.Field[string] `json:"government_id"`
	// Legal (formal) business name.
	LegalBusinessName param.Field[string] `json:"legal_business_name"`
	// Parent company name (if applicable).
	ParentCompany param.Field[string] `json:"parent_company"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers param.Field[[]string] `json:"phone_numbers"`
}

func (r AccountHolderUpdateParamsBodyKYBPatchRequestBeneficialOwnerEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Individuals associated with a KYB application. Phone number is optional.
type AccountHolderUpdateParamsBodyKYBPatchRequestBeneficialOwnerIndividual struct {
	// Globally unique identifier for an entity.
	EntityToken param.Field[string] `json:"entity_token,required" format:"uuid"`
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[AddressUpdateParam] `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob param.Field[string] `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email param.Field[string] `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName param.Field[string] `json:"first_name"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID param.Field[string] `json:"government_id"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName param.Field[string] `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r AccountHolderUpdateParamsBodyKYBPatchRequestBeneficialOwnerIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information for business for which the account is being opened and KYB is being
// run.
type AccountHolderUpdateParamsBodyKYBPatchRequestBusinessEntity struct {
	// Globally unique identifier for an entity.
	EntityToken param.Field[string] `json:"entity_token,required" format:"uuid"`
	// Business”s physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address param.Field[AddressUpdateParam] `json:"address"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName param.Field[string] `json:"dba_business_name"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID param.Field[string] `json:"government_id"`
	// Legal (formal) business name.
	LegalBusinessName param.Field[string] `json:"legal_business_name"`
	// Parent company name (if applicable).
	ParentCompany param.Field[string] `json:"parent_company"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers param.Field[[]string] `json:"phone_numbers"`
}

func (r AccountHolderUpdateParamsBodyKYBPatchRequestBusinessEntity) MarshalJSON() (data []byte, err error) {
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
type AccountHolderUpdateParamsBodyKYBPatchRequestControlPerson struct {
	// Globally unique identifier for an entity.
	EntityToken param.Field[string] `json:"entity_token,required" format:"uuid"`
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[AddressUpdateParam] `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob param.Field[string] `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email param.Field[string] `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName param.Field[string] `json:"first_name"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID param.Field[string] `json:"government_id"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName param.Field[string] `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r AccountHolderUpdateParamsBodyKYBPatchRequestControlPerson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The KYC request payload for updating an account holder.
type AccountHolderUpdateParamsBodyKYCPatchRequest struct {
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// Information on the individual for whom the account is being opened and KYC is
	// being run.
	Individual param.Field[AccountHolderUpdateParamsBodyKYCPatchRequestIndividual] `json:"individual"`
}

func (r AccountHolderUpdateParamsBodyKYCPatchRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountHolderUpdateParamsBodyKYCPatchRequest) implementsAccountHolderUpdateParamsBodyUnion() {
}

// Information on the individual for whom the account is being opened and KYC is
// being run.
type AccountHolderUpdateParamsBodyKYCPatchRequestIndividual struct {
	// Globally unique identifier for an entity.
	EntityToken param.Field[string] `json:"entity_token,required" format:"uuid"`
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[AddressUpdateParam] `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob param.Field[string] `json:"dob"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email param.Field[string] `json:"email"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName param.Field[string] `json:"first_name"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID param.Field[string] `json:"government_id"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName param.Field[string] `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r AccountHolderUpdateParamsBodyKYCPatchRequestIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The legacy request for updating an account holder.
type AccountHolderUpdateParamsBodyPatchRequest struct {
	// Allowed for: KYC-Exempt, BYO-KYC, BYO-KYB.
	Address param.Field[AddressUpdateParam] `json:"address"`
	// Allowed for: KYC-Exempt, BYO-KYC. The token of the business account to which the
	// account holder is associated.
	BusinessAccountToken param.Field[string] `json:"business_account_token"`
	// Allowed for all Account Holders. Account holder's email address. The primary
	// purpose of this field is for cardholder identification and verification during
	// the digital wallet tokenization process.
	Email param.Field[string] `json:"email"`
	// Allowed for KYC-Exempt, BYO-KYC. Account holder's first name.
	FirstName param.Field[string] `json:"first_name"`
	// Allowed for KYC-Exempt, BYO-KYC. Account holder's last name.
	LastName param.Field[string] `json:"last_name"`
	// Allowed for BYO-KYB. Legal business name of the account holder.
	LegalBusinessName param.Field[string] `json:"legal_business_name"`
	// Allowed for all Account Holders. Account holder's phone number, entered in E.164
	// format. The primary purpose of this field is for cardholder identification and
	// verification during the digital wallet tokenization process.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r AccountHolderUpdateParamsBodyPatchRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountHolderUpdateParamsBodyPatchRequest) implementsAccountHolderUpdateParamsBodyUnion() {}

type AccountHolderListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Email address of the account holder. The query must be an exact match, case
	// insensitive.
	Email param.Field[string] `query:"email"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// If applicable, represents the external_id associated with the account_holder.
	ExternalID param.Field[string] `query:"external_id" format:"uuid"`
	// (Individual Account Holders only) The first name of the account holder. The
	// query is case insensitive and supports partial matches.
	FirstName param.Field[string] `query:"first_name"`
	// (Individual Account Holders only) The last name of the account holder. The query
	// is case insensitive and supports partial matches.
	LastName param.Field[string] `query:"last_name"`
	// (Business Account Holders only) The legal business name of the account holder.
	// The query is case insensitive and supports partial matches.
	LegalBusinessName param.Field[string] `query:"legal_business_name"`
	// The number of account_holders to limit the response to.
	Limit param.Field[int64] `query:"limit"`
	// Phone number of the account holder. The query must be an exact match.
	PhoneNumber param.Field[string] `query:"phone_number"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [AccountHolderListParams]'s query parameters as
// `url.Values`.
func (r AccountHolderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountHolderSimulateEnrollmentDocumentReviewParams struct {
	// The account holder document upload which to perform the simulation upon.
	DocumentUploadToken param.Field[string] `json:"document_upload_token,required"`
	// An account holder document's upload status for use within the simulation.
	Status param.Field[AccountHolderSimulateEnrollmentDocumentReviewParamsStatus] `json:"status,required"`
	// A list of status reasons associated with a KYB account holder in PENDING_REVIEW
	AcceptedEntityStatusReasons param.Field[[]string] `json:"accepted_entity_status_reasons"`
	// Status reason that will be associated with the simulated account holder status.
	// Only required for a `REJECTED` status or `PARTIAL_APPROVAL` status.
	StatusReason param.Field[AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason] `json:"status_reason"`
}

func (r AccountHolderSimulateEnrollmentDocumentReviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An account holder document's upload status for use within the simulation.
type AccountHolderSimulateEnrollmentDocumentReviewParamsStatus string

const (
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusUploaded        AccountHolderSimulateEnrollmentDocumentReviewParamsStatus = "UPLOADED"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusAccepted        AccountHolderSimulateEnrollmentDocumentReviewParamsStatus = "ACCEPTED"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusRejected        AccountHolderSimulateEnrollmentDocumentReviewParamsStatus = "REJECTED"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusPartialApproval AccountHolderSimulateEnrollmentDocumentReviewParamsStatus = "PARTIAL_APPROVAL"
)

func (r AccountHolderSimulateEnrollmentDocumentReviewParamsStatus) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentDocumentReviewParamsStatusUploaded, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusAccepted, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusRejected, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusPartialApproval:
		return true
	}
	return false
}

// Status reason that will be associated with the simulated account holder status.
// Only required for a `REJECTED` status or `PARTIAL_APPROVAL` status.
type AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason string

const (
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentMissingRequiredData     AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "DOCUMENT_MISSING_REQUIRED_DATA"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentUploadTooBlurry         AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "DOCUMENT_UPLOAD_TOO_BLURRY"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonFileSizeTooLarge                AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "FILE_SIZE_TOO_LARGE"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidDocumentType             AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "INVALID_DOCUMENT_TYPE"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidDocumentUpload           AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "INVALID_DOCUMENT_UPLOAD"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidEntity                   AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "INVALID_ENTITY"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentExpired                 AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "DOCUMENT_EXPIRED"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentIssuedGreaterThan30Days AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "DOCUMENT_ISSUED_GREATER_THAN_30_DAYS"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentTypeNotSupported        AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "DOCUMENT_TYPE_NOT_SUPPORTED"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonUnknownFailureReason            AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "UNKNOWN_FAILURE_REASON"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonUnknownError                    AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "UNKNOWN_ERROR"
)

func (r AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentMissingRequiredData, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentUploadTooBlurry, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonFileSizeTooLarge, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidDocumentType, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidDocumentUpload, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidEntity, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentExpired, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentIssuedGreaterThan30Days, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentTypeNotSupported, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonUnknownFailureReason, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonUnknownError:
		return true
	}
	return false
}

type AccountHolderSimulateEnrollmentReviewParams struct {
	// The account holder which to perform the simulation upon.
	AccountHolderToken param.Field[string] `json:"account_holder_token"`
	// An account holder's status for use within the simulation.
	Status param.Field[AccountHolderSimulateEnrollmentReviewParamsStatus] `json:"status"`
	// Status reason that will be associated with the simulated account holder status.
	// Only required for a `REJECTED` status.
	StatusReasons param.Field[[]AccountHolderSimulateEnrollmentReviewParamsStatusReason] `json:"status_reasons"`
}

func (r AccountHolderSimulateEnrollmentReviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An account holder's status for use within the simulation.
type AccountHolderSimulateEnrollmentReviewParamsStatus string

const (
	AccountHolderSimulateEnrollmentReviewParamsStatusAccepted AccountHolderSimulateEnrollmentReviewParamsStatus = "ACCEPTED"
	AccountHolderSimulateEnrollmentReviewParamsStatusRejected AccountHolderSimulateEnrollmentReviewParamsStatus = "REJECTED"
)

func (r AccountHolderSimulateEnrollmentReviewParamsStatus) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewParamsStatusAccepted, AccountHolderSimulateEnrollmentReviewParamsStatusRejected:
		return true
	}
	return false
}

type AccountHolderSimulateEnrollmentReviewParamsStatusReason string

const (
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityIDVerificationFailure       AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityAddressVerificationFailure  AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_ADDRESS_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityNameVerificationFailure     AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_NAME_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched  AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_BUSINESS_OFFICERS_NOT_MATCHED"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntitySosFilingInactive           AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_FILING_INACTIVE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntitySosNotMatched               AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_NOT_MATCHED"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityCmraFailure                 AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_CMRA_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityWatchlistFailure            AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_WATCHLIST_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityRegisteredAgentFailure      AccountHolderSimulateEnrollmentReviewParamsStatusReason = "PRIMARY_BUSINESS_ENTITY_REGISTERED_AGENT_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonBlocklistAlertFailure               AccountHolderSimulateEnrollmentReviewParamsStatusReason = "CONTROL_PERSON_BLOCKLIST_ALERT_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonIDVerificationFailure               AccountHolderSimulateEnrollmentReviewParamsStatusReason = "CONTROL_PERSON_ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonDobVerificationFailure              AccountHolderSimulateEnrollmentReviewParamsStatusReason = "CONTROL_PERSON_DOB_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonNameVerificationFailure             AccountHolderSimulateEnrollmentReviewParamsStatusReason = "CONTROL_PERSON_NAME_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualDobVerificationFailure  AccountHolderSimulateEnrollmentReviewParamsStatusReason = "BENEFICIAL_OWNER_INDIVIDUAL_DOB_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualBlocklistAlertFailure   AccountHolderSimulateEnrollmentReviewParamsStatusReason = "BENEFICIAL_OWNER_INDIVIDUAL_BLOCKLIST_ALERT_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualIDVerificationFailure   AccountHolderSimulateEnrollmentReviewParamsStatusReason = "BENEFICIAL_OWNER_INDIVIDUAL_ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualNameVerificationFailure AccountHolderSimulateEnrollmentReviewParamsStatusReason = "BENEFICIAL_OWNER_INDIVIDUAL_NAME_VERIFICATION_FAILURE"
)

func (r AccountHolderSimulateEnrollmentReviewParamsStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityIDVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityAddressVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityNameVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntitySosFilingInactive, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntitySosNotMatched, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityCmraFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityWatchlistFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityRegisteredAgentFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonBlocklistAlertFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonIDVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonDobVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonNameVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualDobVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualBlocklistAlertFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualIDVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualNameVerificationFailure:
		return true
	}
	return false
}

type AccountHolderUploadDocumentParams struct {
	// The type of document to upload
	DocumentType param.Field[AccountHolderUploadDocumentParamsDocumentType] `json:"document_type,required"`
	// Globally unique identifier for the entity.
	EntityToken param.Field[string] `json:"entity_token,required" format:"uuid"`
}

func (r AccountHolderUploadDocumentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of document to upload
type AccountHolderUploadDocumentParamsDocumentType string

const (
	AccountHolderUploadDocumentParamsDocumentTypeEinLetter                 AccountHolderUploadDocumentParamsDocumentType = "EIN_LETTER"
	AccountHolderUploadDocumentParamsDocumentTypeTaxReturn                 AccountHolderUploadDocumentParamsDocumentType = "TAX_RETURN"
	AccountHolderUploadDocumentParamsDocumentTypeOperatingAgreement        AccountHolderUploadDocumentParamsDocumentType = "OPERATING_AGREEMENT"
	AccountHolderUploadDocumentParamsDocumentTypeCertificateOfFormation    AccountHolderUploadDocumentParamsDocumentType = "CERTIFICATE_OF_FORMATION"
	AccountHolderUploadDocumentParamsDocumentTypeDriversLicense            AccountHolderUploadDocumentParamsDocumentType = "DRIVERS_LICENSE"
	AccountHolderUploadDocumentParamsDocumentTypePassport                  AccountHolderUploadDocumentParamsDocumentType = "PASSPORT"
	AccountHolderUploadDocumentParamsDocumentTypePassportCard              AccountHolderUploadDocumentParamsDocumentType = "PASSPORT_CARD"
	AccountHolderUploadDocumentParamsDocumentTypeCertificateOfGoodStanding AccountHolderUploadDocumentParamsDocumentType = "CERTIFICATE_OF_GOOD_STANDING"
	AccountHolderUploadDocumentParamsDocumentTypeArticlesOfIncorporation   AccountHolderUploadDocumentParamsDocumentType = "ARTICLES_OF_INCORPORATION"
	AccountHolderUploadDocumentParamsDocumentTypeArticlesOfOrganization    AccountHolderUploadDocumentParamsDocumentType = "ARTICLES_OF_ORGANIZATION"
	AccountHolderUploadDocumentParamsDocumentTypeBylaws                    AccountHolderUploadDocumentParamsDocumentType = "BYLAWS"
	AccountHolderUploadDocumentParamsDocumentTypeGovernmentBusinessLicense AccountHolderUploadDocumentParamsDocumentType = "GOVERNMENT_BUSINESS_LICENSE"
	AccountHolderUploadDocumentParamsDocumentTypePartnershipAgreement      AccountHolderUploadDocumentParamsDocumentType = "PARTNERSHIP_AGREEMENT"
	AccountHolderUploadDocumentParamsDocumentTypeSs4Form                   AccountHolderUploadDocumentParamsDocumentType = "SS4_FORM"
	AccountHolderUploadDocumentParamsDocumentTypeBankStatement             AccountHolderUploadDocumentParamsDocumentType = "BANK_STATEMENT"
	AccountHolderUploadDocumentParamsDocumentTypeUtilityBillStatement      AccountHolderUploadDocumentParamsDocumentType = "UTILITY_BILL_STATEMENT"
	AccountHolderUploadDocumentParamsDocumentTypeSsnCard                   AccountHolderUploadDocumentParamsDocumentType = "SSN_CARD"
	AccountHolderUploadDocumentParamsDocumentTypeItinLetter                AccountHolderUploadDocumentParamsDocumentType = "ITIN_LETTER"
	AccountHolderUploadDocumentParamsDocumentTypeFincenBoiReport           AccountHolderUploadDocumentParamsDocumentType = "FINCEN_BOI_REPORT"
)

func (r AccountHolderUploadDocumentParamsDocumentType) IsKnown() bool {
	switch r {
	case AccountHolderUploadDocumentParamsDocumentTypeEinLetter, AccountHolderUploadDocumentParamsDocumentTypeTaxReturn, AccountHolderUploadDocumentParamsDocumentTypeOperatingAgreement, AccountHolderUploadDocumentParamsDocumentTypeCertificateOfFormation, AccountHolderUploadDocumentParamsDocumentTypeDriversLicense, AccountHolderUploadDocumentParamsDocumentTypePassport, AccountHolderUploadDocumentParamsDocumentTypePassportCard, AccountHolderUploadDocumentParamsDocumentTypeCertificateOfGoodStanding, AccountHolderUploadDocumentParamsDocumentTypeArticlesOfIncorporation, AccountHolderUploadDocumentParamsDocumentTypeArticlesOfOrganization, AccountHolderUploadDocumentParamsDocumentTypeBylaws, AccountHolderUploadDocumentParamsDocumentTypeGovernmentBusinessLicense, AccountHolderUploadDocumentParamsDocumentTypePartnershipAgreement, AccountHolderUploadDocumentParamsDocumentTypeSs4Form, AccountHolderUploadDocumentParamsDocumentTypeBankStatement, AccountHolderUploadDocumentParamsDocumentTypeUtilityBillStatement, AccountHolderUploadDocumentParamsDocumentTypeSsnCard, AccountHolderUploadDocumentParamsDocumentTypeItinLetter, AccountHolderUploadDocumentParamsDocumentTypeFincenBoiReport:
		return true
	}
	return false
}
