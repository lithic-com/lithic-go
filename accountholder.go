// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/shared"
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

// Run an individual or business's information through the Customer Identification
// Program (CIP). All calls to this endpoint will return an immediate response -
// though in some cases, the response may indicate the enrollment is under review
// or further action will be needed to complete the account enrollment process.
// This endpoint can only be used on accounts that are part of the program that the
// calling API key manages.
//
// Note: If you choose to set a timeout for this request, we recommend 5 minutes.
func (r *AccountHolderService) New(ctx context.Context, body AccountHolderNewParams, opts ...option.RequestOption) (res *AccountHolderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_holders"
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
	path := fmt.Sprintf("account_holders/%s", accountHolderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the information associated with a particular account holder.
func (r *AccountHolderService) Update(ctx context.Context, accountHolderToken string, body AccountHolderUpdateParams, opts ...option.RequestOption) (res *AccountHolderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	path := fmt.Sprintf("account_holders/%s", accountHolderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of individual or business account holders and their KYC or KYB
// evaluation status.
func (r *AccountHolderService) List(ctx context.Context, query AccountHolderListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccountHolder], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_holders"
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
	path := fmt.Sprintf("account_holders/%s/documents", accountHolderToken)
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
func (r *AccountHolderService) Resubmit(ctx context.Context, accountHolderToken string, body AccountHolderResubmitParams, opts ...option.RequestOption) (res *AccountHolder, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	path := fmt.Sprintf("account_holders/%s/resubmit", accountHolderToken)
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
func (r *AccountHolderService) GetDocument(ctx context.Context, accountHolderToken string, documentToken string, opts ...option.RequestOption) (res *AccountHolderDocument, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	if documentToken == "" {
		err = errors.New("missing required document_token parameter")
		return
	}
	path := fmt.Sprintf("account_holders/%s/documents/%s", accountHolderToken, documentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Simulates a review for an account holder document upload.
func (r *AccountHolderService) SimulateEnrollmentDocumentReview(ctx context.Context, body AccountHolderSimulateEnrollmentDocumentReviewParams, opts ...option.RequestOption) (res *AccountHolderSimulateEnrollmentDocumentReviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/account_holders/enrollment_document_review"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates an enrollment review for an account holder. This endpoint is only
// applicable for workflows that may required intervention such as `KYB_BASIC` or
// `KYC_ADVANCED`.
func (r *AccountHolderService) SimulateEnrollmentReview(ctx context.Context, body AccountHolderSimulateEnrollmentReviewParams, opts ...option.RequestOption) (res *AccountHolderSimulateEnrollmentReviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/account_holders/enrollment_review"
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
func (r *AccountHolderService) UploadDocument(ctx context.Context, accountHolderToken string, body AccountHolderUploadDocumentParams, opts ...option.RequestOption) (res *AccountHolderDocument, err error) {
	opts = append(r.Options[:], opts...)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	path := fmt.Sprintf("account_holders/%s/documents", accountHolderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountHolder struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token,required" format:"uuid"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token" format:"uuid"`
	// Only present when user_type == "BUSINESS". List of all entities with >25%
	// ownership in the company.
	BeneficialOwnerEntities []AccountHolderBeneficialOwnerEntity `json:"beneficial_owner_entities"`
	// Only present when user_type == "BUSINESS". List of all individuals with >25%
	// ownership in the company.
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
	// Timestamp of when the account holder was created.
	Created time.Time `json:"created" format:"date-time"`
	// < Deprecated. Use control_person.email when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".
	//
	// > Primary email of Account Holder.
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
	// < Deprecated. Use control_person.phone_number when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".
	//
	// > Primary phone of Account Holder, entered in E.164 format.
	PhoneNumber string `json:"phone_number"`
	// Only present for "KYB_BASIC" and "KYC_ADVANCED" workflows. A list of documents
	// required for the account holder to be approved.
	RequiredDocuments []AccountHolderRequiredDocument `json:"required_documents"`
	// <Deprecated. Use verification_application.status instead>
	//
	// KYC and KYB evaluation states.
	//
	// Note:
	//
	//   - `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	//     `KYC_ADVANCED` workflow.
	//   - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
	Status AccountHolderStatus `json:"status"`
	// <Deprecated. Use verification_application.status_reasons> Reason for the
	// evaluation status.
	StatusReasons []AccountHolderStatusReason `json:"status_reasons"`
	// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
	// attribute will be present. If the type is "BUSINESS" then the "business_entity",
	// "control_person", "beneficial_owner_individuals", "beneficial_owner_entities",
	// "nature_of_business", and "website_url" attributes will be present.
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
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token" format:"uuid"`
	// Parent company name (if applicable).
	ParentCompany string                                 `json:"parent_company"`
	JSON          accountHolderBeneficialOwnerEntityJSON `json:"-"`
}

// accountHolderBeneficialOwnerEntityJSON contains the JSON metadata for the struct
// [AccountHolderBeneficialOwnerEntity]
type accountHolderBeneficialOwnerEntityJSON struct {
	Address           apijson.Field
	GovernmentID      apijson.Field
	LegalBusinessName apijson.Field
	PhoneNumbers      apijson.Field
	DbaBusinessName   apijson.Field
	EntityToken       apijson.Field
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
	Address shared.Address `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address.
	Email string `json:"email"`
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token" format:"uuid"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                                     `json:"phone_number"`
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
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token" format:"uuid"`
	// Parent company name (if applicable).
	ParentCompany string                          `json:"parent_company"`
	JSON          accountHolderBusinessEntityJSON `json:"-"`
}

// accountHolderBusinessEntityJSON contains the JSON metadata for the struct
// [AccountHolderBusinessEntity]
type accountHolderBusinessEntityJSON struct {
	Address           apijson.Field
	GovernmentID      apijson.Field
	LegalBusinessName apijson.Field
	PhoneNumbers      apijson.Field
	DbaBusinessName   apijson.Field
	EntityToken       apijson.Field
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
	Address shared.Address `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address.
	Email string `json:"email"`
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token" format:"uuid"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                         `json:"phone_number"`
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
	Address shared.Address `json:"address"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob string `json:"dob"`
	// Individual's email address.
	Email string `json:"email"`
	// Globally unique identifier for the entity.
	EntityToken string `json:"entity_token" format:"uuid"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName string `json:"first_name"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName string `json:"last_name"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber string                      `json:"phone_number"`
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

type AccountHolderRequiredDocument struct {
	// Globally unique identifier for an entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// rovides the status reasons that will be satisfied by providing one of the valid
	// documents.
	StatusReasons []string `json:"status_reasons,required"`
	// A list of valid documents that will satisfy the KYC requirements for the
	// specified entity.
	ValidDocuments []string                          `json:"valid_documents,required"`
	JSON           accountHolderRequiredDocumentJSON `json:"-"`
}

// accountHolderRequiredDocumentJSON contains the JSON metadata for the struct
// [AccountHolderRequiredDocument]
type accountHolderRequiredDocumentJSON struct {
	EntityToken    apijson.Field
	StatusReasons  apijson.Field
	ValidDocuments apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountHolderRequiredDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderRequiredDocumentJSON) RawJSON() string {
	return r.raw
}

// <Deprecated. Use verification_application.status instead>
//
// KYC and KYB evaluation states.
//
// Note:
//
//   - `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
//     `KYC_ADVANCED` workflow.
//   - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
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
// "control_person", "beneficial_owner_individuals", "beneficial_owner_entities",
// "nature_of_business", and "website_url" attributes will be present.
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
	//   - `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	//     `KYC_ADVANCED` workflow.
	//   - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
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
//   - `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
//     `KYC_ADVANCED` workflow.
//   - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
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

// Describes the document and the required document image uploads required to
// re-run KYC.
type AccountHolderDocument struct {
	// Globally unique identifier for the document.
	Token string `json:"token" format:"uuid"`
	// Globally unique identifier for the account holder.
	AccountHolderToken string `json:"account_holder_token" format:"uuid"`
	// Type of documentation to be submitted for verification.
	DocumentType AccountHolderDocumentDocumentType `json:"document_type"`
	// Globally unique identifier for the entity.
	EntityToken             string                                        `json:"entity_token" format:"uuid"`
	RequiredDocumentUploads []AccountHolderDocumentRequiredDocumentUpload `json:"required_document_uploads"`
	JSON                    accountHolderDocumentJSON                     `json:"-"`
}

// accountHolderDocumentJSON contains the JSON metadata for the struct
// [AccountHolderDocument]
type accountHolderDocumentJSON struct {
	Token                   apijson.Field
	AccountHolderToken      apijson.Field
	DocumentType            apijson.Field
	EntityToken             apijson.Field
	RequiredDocumentUploads apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountHolderDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderDocumentJSON) RawJSON() string {
	return r.raw
}

// Type of documentation to be submitted for verification.
type AccountHolderDocumentDocumentType string

const (
	AccountHolderDocumentDocumentTypeEinLetter                 AccountHolderDocumentDocumentType = "EIN_LETTER"
	AccountHolderDocumentDocumentTypeTaxReturn                 AccountHolderDocumentDocumentType = "TAX_RETURN"
	AccountHolderDocumentDocumentTypeOperatingAgreement        AccountHolderDocumentDocumentType = "OPERATING_AGREEMENT"
	AccountHolderDocumentDocumentTypeCertificateOfFormation    AccountHolderDocumentDocumentType = "CERTIFICATE_OF_FORMATION"
	AccountHolderDocumentDocumentTypeDriversLicense            AccountHolderDocumentDocumentType = "DRIVERS_LICENSE"
	AccountHolderDocumentDocumentTypePassport                  AccountHolderDocumentDocumentType = "PASSPORT"
	AccountHolderDocumentDocumentTypePassportCard              AccountHolderDocumentDocumentType = "PASSPORT_CARD"
	AccountHolderDocumentDocumentTypeCertificateOfGoodStanding AccountHolderDocumentDocumentType = "CERTIFICATE_OF_GOOD_STANDING"
	AccountHolderDocumentDocumentTypeArticlesOfIncorporation   AccountHolderDocumentDocumentType = "ARTICLES_OF_INCORPORATION"
	AccountHolderDocumentDocumentTypeArticlesOfOrganization    AccountHolderDocumentDocumentType = "ARTICLES_OF_ORGANIZATION"
	AccountHolderDocumentDocumentTypeBylaws                    AccountHolderDocumentDocumentType = "BYLAWS"
	AccountHolderDocumentDocumentTypeGovernmentBusinessLicense AccountHolderDocumentDocumentType = "GOVERNMENT_BUSINESS_LICENSE"
	AccountHolderDocumentDocumentTypePartnershipAgreement      AccountHolderDocumentDocumentType = "PARTNERSHIP_AGREEMENT"
	AccountHolderDocumentDocumentTypeSs4Form                   AccountHolderDocumentDocumentType = "SS4_FORM"
	AccountHolderDocumentDocumentTypeBankStatement             AccountHolderDocumentDocumentType = "BANK_STATEMENT"
	AccountHolderDocumentDocumentTypeUtilityBillStatement      AccountHolderDocumentDocumentType = "UTILITY_BILL_STATEMENT"
	AccountHolderDocumentDocumentTypeSsnCard                   AccountHolderDocumentDocumentType = "SSN_CARD"
	AccountHolderDocumentDocumentTypeItinLetter                AccountHolderDocumentDocumentType = "ITIN_LETTER"
)

func (r AccountHolderDocumentDocumentType) IsKnown() bool {
	switch r {
	case AccountHolderDocumentDocumentTypeEinLetter, AccountHolderDocumentDocumentTypeTaxReturn, AccountHolderDocumentDocumentTypeOperatingAgreement, AccountHolderDocumentDocumentTypeCertificateOfFormation, AccountHolderDocumentDocumentTypeDriversLicense, AccountHolderDocumentDocumentTypePassport, AccountHolderDocumentDocumentTypePassportCard, AccountHolderDocumentDocumentTypeCertificateOfGoodStanding, AccountHolderDocumentDocumentTypeArticlesOfIncorporation, AccountHolderDocumentDocumentTypeArticlesOfOrganization, AccountHolderDocumentDocumentTypeBylaws, AccountHolderDocumentDocumentTypeGovernmentBusinessLicense, AccountHolderDocumentDocumentTypePartnershipAgreement, AccountHolderDocumentDocumentTypeSs4Form, AccountHolderDocumentDocumentTypeBankStatement, AccountHolderDocumentDocumentTypeUtilityBillStatement, AccountHolderDocumentDocumentTypeSsnCard, AccountHolderDocumentDocumentTypeItinLetter:
		return true
	}
	return false
}

// Represents a single image of the document to upload.
type AccountHolderDocumentRequiredDocumentUpload struct {
	// Globally unique identifier for the document upload.
	Token string `json:"token" format:"uuid"`
	// Type of image to upload.
	ImageType AccountHolderDocumentRequiredDocumentUploadsImageType `json:"image_type"`
	// Status of document image upload.
	Status        AccountHolderDocumentRequiredDocumentUploadsStatus         `json:"status"`
	StatusReasons []AccountHolderDocumentRequiredDocumentUploadsStatusReason `json:"status_reasons"`
	// URL to upload document image to.
	//
	// Note that the upload URLs expire after 7 days. If an upload URL expires, you can
	// refresh the URLs by retrieving the document upload from
	// `GET /account_holders/{account_holder_token}/documents`.
	UploadURL string                                          `json:"upload_url"`
	JSON      accountHolderDocumentRequiredDocumentUploadJSON `json:"-"`
}

// accountHolderDocumentRequiredDocumentUploadJSON contains the JSON metadata for
// the struct [AccountHolderDocumentRequiredDocumentUpload]
type accountHolderDocumentRequiredDocumentUploadJSON struct {
	Token         apijson.Field
	ImageType     apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	UploadURL     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountHolderDocumentRequiredDocumentUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderDocumentRequiredDocumentUploadJSON) RawJSON() string {
	return r.raw
}

// Type of image to upload.
type AccountHolderDocumentRequiredDocumentUploadsImageType string

const (
	AccountHolderDocumentRequiredDocumentUploadsImageTypeBack  AccountHolderDocumentRequiredDocumentUploadsImageType = "back"
	AccountHolderDocumentRequiredDocumentUploadsImageTypeFront AccountHolderDocumentRequiredDocumentUploadsImageType = "front"
)

func (r AccountHolderDocumentRequiredDocumentUploadsImageType) IsKnown() bool {
	switch r {
	case AccountHolderDocumentRequiredDocumentUploadsImageTypeBack, AccountHolderDocumentRequiredDocumentUploadsImageTypeFront:
		return true
	}
	return false
}

// Status of document image upload.
type AccountHolderDocumentRequiredDocumentUploadsStatus string

const (
	AccountHolderDocumentRequiredDocumentUploadsStatusCompleted     AccountHolderDocumentRequiredDocumentUploadsStatus = "COMPLETED"
	AccountHolderDocumentRequiredDocumentUploadsStatusFailed        AccountHolderDocumentRequiredDocumentUploadsStatus = "FAILED"
	AccountHolderDocumentRequiredDocumentUploadsStatusPendingUpload AccountHolderDocumentRequiredDocumentUploadsStatus = "PENDING_UPLOAD"
	AccountHolderDocumentRequiredDocumentUploadsStatusUploaded      AccountHolderDocumentRequiredDocumentUploadsStatus = "UPLOADED"
)

func (r AccountHolderDocumentRequiredDocumentUploadsStatus) IsKnown() bool {
	switch r {
	case AccountHolderDocumentRequiredDocumentUploadsStatusCompleted, AccountHolderDocumentRequiredDocumentUploadsStatusFailed, AccountHolderDocumentRequiredDocumentUploadsStatusPendingUpload, AccountHolderDocumentRequiredDocumentUploadsStatusUploaded:
		return true
	}
	return false
}

// Reasons for document image upload status.
type AccountHolderDocumentRequiredDocumentUploadsStatusReason string

const (
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonBackImageBlurry  AccountHolderDocumentRequiredDocumentUploadsStatusReason = "BACK_IMAGE_BLURRY"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonFileSizeTooLarge AccountHolderDocumentRequiredDocumentUploadsStatusReason = "FILE_SIZE_TOO_LARGE"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonFrontImageBlurry AccountHolderDocumentRequiredDocumentUploadsStatusReason = "FRONT_IMAGE_BLURRY"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonFrontImageGlare  AccountHolderDocumentRequiredDocumentUploadsStatusReason = "FRONT_IMAGE_GLARE"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonInvalidFileType  AccountHolderDocumentRequiredDocumentUploadsStatusReason = "INVALID_FILE_TYPE"
	AccountHolderDocumentRequiredDocumentUploadsStatusReasonUnknownError     AccountHolderDocumentRequiredDocumentUploadsStatusReason = "UNKNOWN_ERROR"
)

func (r AccountHolderDocumentRequiredDocumentUploadsStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderDocumentRequiredDocumentUploadsStatusReasonBackImageBlurry, AccountHolderDocumentRequiredDocumentUploadsStatusReasonFileSizeTooLarge, AccountHolderDocumentRequiredDocumentUploadsStatusReasonFrontImageBlurry, AccountHolderDocumentRequiredDocumentUploadsStatusReasonFrontImageGlare, AccountHolderDocumentRequiredDocumentUploadsStatusReasonInvalidFileType, AccountHolderDocumentRequiredDocumentUploadsStatusReasonUnknownError:
		return true
	}
	return false
}

type KYBParam struct {
	// List of all entities with >25% ownership in the company. If no entity or
	// individual owns >25% of the company, and the largest shareholder is an entity,
	// please identify them in this field. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background. If no business owner is an entity, pass in an
	// empty list. However, either this parameter or `beneficial_owner_individuals`
	// must be populated. on entities that should be included.
	BeneficialOwnerEntities param.Field[[]KYBBeneficialOwnerEntityParam] `json:"beneficial_owner_entities,required"`
	// List of all direct and indirect individuals with >25% ownership in the company.
	// If no entity or individual owns >25% of the company, and the largest shareholder
	// is an individual, please identify them in this field. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included. If no
	// individual is an entity, pass in an empty list. However, either this parameter
	// or `beneficial_owner_entities` must be populated.
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
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
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
	// Globally unique identifier for the entity.
	EntityToken param.Field[string] `json:"entity_token" format:"uuid"`
	// Parent company name (if applicable).
	ParentCompany param.Field[string] `json:"parent_company"`
}

func (r KYBBeneficialOwnerEntityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	// Globally unique identifier for the entity.
	EntityToken param.Field[string] `json:"entity_token" format:"uuid"`
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
	KYCWorkflowKYCAdvanced KYCWorkflow = "KYC_ADVANCED"
	KYCWorkflowKYCBasic    KYCWorkflow = "KYC_BASIC"
	KYCWorkflowKYCByo      KYCWorkflow = "KYC_BYO"
)

func (r KYCWorkflow) IsKnown() bool {
	switch r {
	case KYCWorkflowKYCAdvanced, KYCWorkflowKYCBasic, KYCWorkflowKYCByo:
		return true
	}
	return false
}

type KYCExemptParam struct {
	// KYC Exempt user's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[shared.AddressParam] `json:"address,required"`
	// The KYC Exempt user's email
	Email param.Field[string] `json:"email,required"`
	// The KYC Exempt user's first name
	FirstName param.Field[string] `json:"first_name,required"`
	// Specifies the type of KYC Exempt user
	KYCExemptionType param.Field[KYCExemptKYCExemptionType] `json:"kyc_exemption_type,required"`
	// The KYC Exempt user's last name
	LastName param.Field[string] `json:"last_name,required"`
	// The KYC Exempt user's phone number
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

type AccountHolderNewResponse struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token,required" format:"uuid"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token,required" format:"uuid"`
	// KYC and KYB evaluation states.
	//
	// Note:
	//
	//   - `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	//     `KYC_ADVANCED` workflow.
	//   - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
	Status AccountHolderNewResponseStatus `json:"status,required"`
	// Reason for the evaluation status.
	StatusReasons []AccountHolderNewResponseStatusReason `json:"status_reasons,required"`
	// Timestamp of when the account holder was created.
	Created time.Time `json:"created" format:"date-time"`
	// Customer-provided token that indicates a relationship with an object outside of
	// the Lithic ecosystem.
	ExternalID string                       `json:"external_id" format:"string"`
	JSON       accountHolderNewResponseJSON `json:"-"`
}

// accountHolderNewResponseJSON contains the JSON metadata for the struct
// [AccountHolderNewResponse]
type accountHolderNewResponseJSON struct {
	Token         apijson.Field
	AccountToken  apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	Created       apijson.Field
	ExternalID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
//   - `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
//     `KYC_ADVANCED` workflow.
//   - `PENDING_REVIEW` is only applicable for the `KYB_BASIC` workflow.
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

type AccountHolderNewResponseStatusReason string

const (
	AccountHolderNewResponseStatusReasonAddressVerificationFailure  AccountHolderNewResponseStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonAgeThresholdFailure         AccountHolderNewResponseStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderNewResponseStatusReasonCompleteVerificationFailure AccountHolderNewResponseStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonDobVerificationFailure      AccountHolderNewResponseStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonIDVerificationFailure       AccountHolderNewResponseStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonMaxDocumentAttempts         AccountHolderNewResponseStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderNewResponseStatusReasonMaxResubmissionAttempts     AccountHolderNewResponseStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderNewResponseStatusReasonNameVerificationFailure     AccountHolderNewResponseStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonOtherVerificationFailure    AccountHolderNewResponseStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderNewResponseStatusReasonRiskThresholdFailure        AccountHolderNewResponseStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderNewResponseStatusReasonWatchlistAlertFailure       AccountHolderNewResponseStatusReason = "WATCHLIST_ALERT_FAILURE"
)

func (r AccountHolderNewResponseStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderNewResponseStatusReasonAddressVerificationFailure, AccountHolderNewResponseStatusReasonAgeThresholdFailure, AccountHolderNewResponseStatusReasonCompleteVerificationFailure, AccountHolderNewResponseStatusReasonDobVerificationFailure, AccountHolderNewResponseStatusReasonIDVerificationFailure, AccountHolderNewResponseStatusReasonMaxDocumentAttempts, AccountHolderNewResponseStatusReasonMaxResubmissionAttempts, AccountHolderNewResponseStatusReasonNameVerificationFailure, AccountHolderNewResponseStatusReasonOtherVerificationFailure, AccountHolderNewResponseStatusReasonRiskThresholdFailure, AccountHolderNewResponseStatusReasonWatchlistAlertFailure:
		return true
	}
	return false
}

type AccountHolderUpdateResponse struct {
	// The token for the account holder that was updated
	Token string `json:"token"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll businesses
	// with authorized users. Pass the account_token of the enrolled business
	// associated with the AUTHORIZED_USER in this field.
	BusinessAccountToken string `json:"business_account_token"`
	// The newly updated email for the account holder
	Email string `json:"email"`
	// The newly updated phone_number for the account holder
	PhoneNumber string                          `json:"phone_number"`
	JSON        accountHolderUpdateResponseJSON `json:"-"`
}

// accountHolderUpdateResponseJSON contains the JSON metadata for the struct
// [AccountHolderUpdateResponse]
type accountHolderUpdateResponseJSON struct {
	Token                apijson.Field
	BusinessAccountToken apijson.Field
	Email                apijson.Field
	PhoneNumber          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountHolderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHolderListDocumentsResponse struct {
	Data []AccountHolderDocument                `json:"data"`
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

// A document to be submitted for account holder verification.
type AccountHolderSimulateEnrollmentDocumentReviewResponse struct {
	// Globally unique identifier for the document.
	Token string `json:"token" format:"uuid"`
	// Globally unique identifier for the account holder.
	AccountHolderToken string `json:"account_holder_token" format:"uuid"`
	// Type of documentation to be submitted for verification.
	DocumentType AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentType `json:"document_type"`
	// List of required document images to upload.
	RequiredDocumentUploads []AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUpload `json:"required_document_uploads"`
	JSON                    accountHolderSimulateEnrollmentDocumentReviewResponseJSON                     `json:"-"`
}

// accountHolderSimulateEnrollmentDocumentReviewResponseJSON contains the JSON
// metadata for the struct [AccountHolderSimulateEnrollmentDocumentReviewResponse]
type accountHolderSimulateEnrollmentDocumentReviewResponseJSON struct {
	Token                   apijson.Field
	AccountHolderToken      apijson.Field
	DocumentType            apijson.Field
	RequiredDocumentUploads apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentDocumentReviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentDocumentReviewResponseJSON) RawJSON() string {
	return r.raw
}

// Type of documentation to be submitted for verification.
type AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentType string

const (
	AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypeCommercialLicense AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentType = "commercial_license"
	AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypeDriversLicense    AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentType = "drivers_license"
	AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypePassport          AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentType = "passport"
	AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypePassportCard      AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentType = "passport_card"
	AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypeVisa              AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentType = "visa"
)

func (r AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentType) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypeCommercialLicense, AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypeDriversLicense, AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypePassport, AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypePassportCard, AccountHolderSimulateEnrollmentDocumentReviewResponseDocumentTypeVisa:
		return true
	}
	return false
}

// Represents a single image of the document to upload.
type AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUpload struct {
	// Type of image to upload.
	ImageType AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageType `json:"image_type"`
	// Status of document image upload.
	Status AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatus `json:"status"`
	// Reasons for document image upload status.
	StatusReasons []AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason `json:"status_reasons"`
	// URL to upload document image to.
	//
	// Note that the upload URLs expire after 7 days. If an upload URL expires, you can
	// refresh the URLs by retrieving the document upload from
	// `GET /account_holders/{account_holder_token}/documents`.
	UploadURL string                                                                          `json:"upload_url"`
	JSON      accountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadJSON `json:"-"`
}

// accountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadJSON
// contains the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUpload]
type accountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadJSON struct {
	ImageType     apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	UploadURL     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadJSON) RawJSON() string {
	return r.raw
}

// Type of image to upload.
type AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageType string

const (
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageTypeBack  AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageType = "back"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageTypeFront AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageType = "front"
)

func (r AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageType) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageTypeBack, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsImageTypeFront:
		return true
	}
	return false
}

// Status of document image upload.
type AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatus string

const (
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusCompleted AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatus = "COMPLETED"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusFailed    AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatus = "FAILED"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusPending   AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatus = "PENDING"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusUploaded  AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatus = "UPLOADED"
)

func (r AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatus) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusCompleted, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusFailed, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusPending, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusUploaded:
		return true
	}
	return false
}

type AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason string

const (
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonBackImageBlurry  AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason = "BACK_IMAGE_BLURRY"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonFileSizeTooLarge AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason = "FILE_SIZE_TOO_LARGE"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonFrontImageBlurry AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason = "FRONT_IMAGE_BLURRY"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonFrontImageGlare  AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason = "FRONT_IMAGE_GLARE"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonInvalidFileType  AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason = "INVALID_FILE_TYPE"
	AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonUnknownError     AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason = "UNKNOWN_ERROR"
)

func (r AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonBackImageBlurry, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonFileSizeTooLarge, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonFrontImageBlurry, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonFrontImageGlare, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonInvalidFileType, AccountHolderSimulateEnrollmentDocumentReviewResponseRequiredDocumentUploadsStatusReasonUnknownError:
		return true
	}
	return false
}

type AccountHolderSimulateEnrollmentReviewResponse struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token" format:"uuid"`
	// Globally unique identifier for the account.
	AccountToken string `json:"account_token" format:"uuid"`
	// Only present when user_type == "BUSINESS". List of all entities with >25%
	// ownership in the company.
	BeneficialOwnerEntities []AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntity `json:"beneficial_owner_entities"`
	// Only present when user_type == "BUSINESS". List of all individuals with >25%
	// ownership in the company.
	BeneficialOwnerIndividuals []AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken string `json:"business_account_token" format:"uuid"`
	// Only present when user_type == "BUSINESS". Information about the business for
	// which the account is being opened and KYB is being run.
	BusinessEntity AccountHolderSimulateEnrollmentReviewResponseBusinessEntity `json:"business_entity"`
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
	// < Deprecated. Use control_person.email when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".
	//
	// > Primary email of Account Holder.
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
	// < Deprecated. Use control_person.phone_number when user_type == "BUSINESS". Use
	// individual.phone_number when user_type == "INDIVIDUAL".
	//
	// > Primary phone of Account Holder, entered in E.164 format.
	PhoneNumber string `json:"phone_number"`
	// Only present for "KYB_BASIC" and "KYC_ADVANCED" workflows. A list of documents
	// required for the account holder to be approved.
	RequiredDocuments []AccountHolderSimulateEnrollmentReviewResponseRequiredDocument `json:"required_documents"`
	// <Deprecated. Use verification_application.status instead>
	//
	// KYC and KYB evaluation states.
	//
	// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	// `ADVANCED` workflow.
	Status AccountHolderSimulateEnrollmentReviewResponseStatus `json:"status"`
	// <Deprecated. Use verification_application.status_reasons> Reason for the
	// evaluation status.
	StatusReasons []AccountHolderSimulateEnrollmentReviewResponseStatusReason `json:"status_reasons"`
	// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
	// attribute will be present.
	//
	// If the type is "BUSINESS" then the "business_entity", "control_person",
	// "beneficial_owner_individuals", "beneficial_owner_entities",
	//
	// "nature_of_business", and "website_url" attributes will be present.
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

type AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntity struct {
	// Business”s physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntitiesAddress `json:"address,required"`
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
	ParentCompany string                                                                 `json:"parent_company"`
	JSON          accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntityJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntityJSON contains
// the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntity]
type accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntityJSON struct {
	Address           apijson.Field
	GovernmentID      apijson.Field
	LegalBusinessName apijson.Field
	PhoneNumbers      apijson.Field
	DbaBusinessName   apijson.Field
	ParentCompany     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntityJSON) RawJSON() string {
	return r.raw
}

// Business”s physical address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable.
type AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntitiesAddress struct {
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
	Address2 string                                                                          `json:"address2"`
	JSON     accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntitiesAddressJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntitiesAddressJSON
// contains the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntitiesAddress]
type accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntitiesAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntitiesAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseBeneficialOwnerEntitiesAddressJSON) RawJSON() string {
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

// Only present when user_type == "BUSINESS". Information about the business for
// which the account is being opened and KYB is being run.
type AccountHolderSimulateEnrollmentReviewResponseBusinessEntity struct {
	// Business”s physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address AccountHolderSimulateEnrollmentReviewResponseBusinessEntityAddress `json:"address,required"`
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
	ParentCompany string                                                          `json:"parent_company"`
	JSON          accountHolderSimulateEnrollmentReviewResponseBusinessEntityJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseBusinessEntityJSON contains the
// JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseBusinessEntity]
type accountHolderSimulateEnrollmentReviewResponseBusinessEntityJSON struct {
	Address           apijson.Field
	GovernmentID      apijson.Field
	LegalBusinessName apijson.Field
	PhoneNumbers      apijson.Field
	DbaBusinessName   apijson.Field
	ParentCompany     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseBusinessEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseBusinessEntityJSON) RawJSON() string {
	return r.raw
}

// Business”s physical address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable.
type AccountHolderSimulateEnrollmentReviewResponseBusinessEntityAddress struct {
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
	JSON     accountHolderSimulateEnrollmentReviewResponseBusinessEntityAddressJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseBusinessEntityAddressJSON contains
// the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseBusinessEntityAddress]
type accountHolderSimulateEnrollmentReviewResponseBusinessEntityAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseBusinessEntityAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseBusinessEntityAddressJSON) RawJSON() string {
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

type AccountHolderSimulateEnrollmentReviewResponseRequiredDocument struct {
	// Globally unique identifier for an entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// rovides the status reasons that will be satisfied by providing one of the valid
	// documents.
	StatusReasons []string `json:"status_reasons,required"`
	// A list of valid documents that will satisfy the KYC requirements for the
	// specified entity.
	ValidDocuments []string                                                          `json:"valid_documents,required"`
	JSON           accountHolderSimulateEnrollmentReviewResponseRequiredDocumentJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseRequiredDocumentJSON contains the
// JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseRequiredDocument]
type accountHolderSimulateEnrollmentReviewResponseRequiredDocumentJSON struct {
	EntityToken    apijson.Field
	StatusReasons  apijson.Field
	ValidDocuments apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountHolderSimulateEnrollmentReviewResponseRequiredDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderSimulateEnrollmentReviewResponseRequiredDocumentJSON) RawJSON() string {
	return r.raw
}

// <Deprecated. Use verification_application.status instead>
//
// KYC and KYB evaluation states.
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
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonAddressVerificationFailure  AccountHolderSimulateEnrollmentReviewResponseStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonAgeThresholdFailure         AccountHolderSimulateEnrollmentReviewResponseStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonCompleteVerificationFailure AccountHolderSimulateEnrollmentReviewResponseStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonDobVerificationFailure      AccountHolderSimulateEnrollmentReviewResponseStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonIDVerificationFailure       AccountHolderSimulateEnrollmentReviewResponseStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonMaxDocumentAttempts         AccountHolderSimulateEnrollmentReviewResponseStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonMaxResubmissionAttempts     AccountHolderSimulateEnrollmentReviewResponseStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonNameVerificationFailure     AccountHolderSimulateEnrollmentReviewResponseStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonOtherVerificationFailure    AccountHolderSimulateEnrollmentReviewResponseStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonRiskThresholdFailure        AccountHolderSimulateEnrollmentReviewResponseStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseStatusReasonWatchlistAlertFailure       AccountHolderSimulateEnrollmentReviewResponseStatusReason = "WATCHLIST_ALERT_FAILURE"
)

func (r AccountHolderSimulateEnrollmentReviewResponseStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewResponseStatusReasonAddressVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonAgeThresholdFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonCompleteVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonDobVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonIDVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonMaxDocumentAttempts, AccountHolderSimulateEnrollmentReviewResponseStatusReasonMaxResubmissionAttempts, AccountHolderSimulateEnrollmentReviewResponseStatusReasonNameVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonOtherVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonRiskThresholdFailure, AccountHolderSimulateEnrollmentReviewResponseStatusReasonWatchlistAlertFailure:
		return true
	}
	return false
}

// The type of Account Holder. If the type is "INDIVIDUAL", the "individual"
// attribute will be present.
//
// If the type is "BUSINESS" then the "business_entity", "control_person",
// "beneficial_owner_individuals", "beneficial_owner_entities",
//
// "nature_of_business", and "website_url" attributes will be present.
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
	Created time.Time `json:"created" format:"date-time"`
	// KYC and KYB evaluation states.
	//
	// Note: `PENDING_RESUBMIT` and `PENDING_DOCUMENT` are only applicable for the
	// `ADVANCED` workflow.
	Status AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatus `json:"status"`
	// Reason for the evaluation status.
	StatusReasons []AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason `json:"status_reasons"`
	// Timestamp of when the application was last updated.
	Updated time.Time                                                                `json:"updated" format:"date-time"`
	JSON    accountHolderSimulateEnrollmentReviewResponseVerificationApplicationJSON `json:"-"`
}

// accountHolderSimulateEnrollmentReviewResponseVerificationApplicationJSON
// contains the JSON metadata for the struct
// [AccountHolderSimulateEnrollmentReviewResponseVerificationApplication]
type accountHolderSimulateEnrollmentReviewResponseVerificationApplicationJSON struct {
	Created       apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	Updated       apijson.Field
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
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonAddressVerificationFailure  AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonAgeThresholdFailure         AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonCompleteVerificationFailure AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonDobVerificationFailure      AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonIDVerificationFailure       AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonMaxDocumentAttempts         AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonMaxResubmissionAttempts     AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonNameVerificationFailure     AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonOtherVerificationFailure    AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonRiskThresholdFailure        AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonWatchlistAlertFailure       AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason = "WATCHLIST_ALERT_FAILURE"
)

func (r AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonAddressVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonAgeThresholdFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonCompleteVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonDobVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonIDVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonMaxDocumentAttempts, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonMaxResubmissionAttempts, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonNameVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonOtherVerificationFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonRiskThresholdFailure, AccountHolderSimulateEnrollmentReviewResponseVerificationApplicationStatusReasonWatchlistAlertFailure:
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
	BeneficialOwnerEntities    param.Field[interface{}] `json:"beneficial_owner_entities,required"`
	BeneficialOwnerIndividuals param.Field[interface{}] `json:"beneficial_owner_individuals,required"`
	BusinessEntity             param.Field[interface{}] `json:"business_entity,required"`
	ControlPerson              param.Field[interface{}] `json:"control_person,required"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID param.Field[string] `json:"external_id"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// business with a pass result.
	//
	// This field is required only if workflow type is `KYB_BYO`.
	KYBPassedTimestamp param.Field[string] `json:"kyb_passed_timestamp"`
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
	Workflow   param.Field[AccountHolderNewParamsBodyWorkflow] `json:"workflow,required"`
	Individual param.Field[interface{}]                        `json:"individual,required"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// individual with a pass result.
	//
	// This field is required only if workflow type is `KYC_BYO`.
	KYCPassedTimestamp param.Field[string] `json:"kyc_passed_timestamp"`
	// KYC Exempt user's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[shared.AddressParam] `json:"address"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken param.Field[string] `json:"business_account_token"`
	// The KYC Exempt user's email
	Email param.Field[string] `json:"email"`
	// The KYC Exempt user's first name
	FirstName param.Field[string] `json:"first_name"`
	// Specifies the type of KYC Exempt user
	KYCExemptionType param.Field[AccountHolderNewParamsBodyKYCExemptionType] `json:"kyc_exemption_type"`
	// The KYC Exempt user's last name
	LastName param.Field[string] `json:"last_name"`
	// The KYC Exempt user's phone number
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r AccountHolderNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountHolderNewParamsBody) implementsAccountHolderNewParamsBodyUnion() {}

// Satisfied by [KYBParam], [KYCParam], [KYCExemptParam],
// [AccountHolderNewParamsBody].
type AccountHolderNewParamsBodyUnion interface {
	implementsAccountHolderNewParamsBodyUnion()
}

// Specifies the type of KYB workflow to run.
type AccountHolderNewParamsBodyWorkflow string

const (
	AccountHolderNewParamsBodyWorkflowKYBBasic    AccountHolderNewParamsBodyWorkflow = "KYB_BASIC"
	AccountHolderNewParamsBodyWorkflowKYBByo      AccountHolderNewParamsBodyWorkflow = "KYB_BYO"
	AccountHolderNewParamsBodyWorkflowKYCAdvanced AccountHolderNewParamsBodyWorkflow = "KYC_ADVANCED"
	AccountHolderNewParamsBodyWorkflowKYCBasic    AccountHolderNewParamsBodyWorkflow = "KYC_BASIC"
	AccountHolderNewParamsBodyWorkflowKYCByo      AccountHolderNewParamsBodyWorkflow = "KYC_BYO"
	AccountHolderNewParamsBodyWorkflowKYCExempt   AccountHolderNewParamsBodyWorkflow = "KYC_EXEMPT"
)

func (r AccountHolderNewParamsBodyWorkflow) IsKnown() bool {
	switch r {
	case AccountHolderNewParamsBodyWorkflowKYBBasic, AccountHolderNewParamsBodyWorkflowKYBByo, AccountHolderNewParamsBodyWorkflowKYCAdvanced, AccountHolderNewParamsBodyWorkflowKYCBasic, AccountHolderNewParamsBodyWorkflowKYCByo, AccountHolderNewParamsBodyWorkflowKYCExempt:
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

type AccountHolderUpdateParams struct {
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken param.Field[string] `json:"business_account_token"`
	// Account holder's email address. The primary purpose of this field is for
	// cardholder identification and verification during the digital wallet
	// tokenization process.
	Email param.Field[string] `json:"email"`
	// Account holder's phone number, entered in E.164 format. The primary purpose of
	// this field is for cardholder identification and verification during the digital
	// wallet tokenization process.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r AccountHolderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHolderListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// If applicable, represents the external_id associated with the account_holder.
	ExternalID param.Field[string] `query:"external_id" format:"uuid"`
	// The number of account_holders to limit the response to.
	Limit param.Field[int64] `query:"limit"`
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

type AccountHolderResubmitParams struct {
	// Information on individual for whom the account is being opened and KYC is being
	// re-run.
	Individual param.Field[AccountHolderResubmitParamsIndividual] `json:"individual,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp param.Field[string]                              `json:"tos_timestamp,required"`
	Workflow     param.Field[AccountHolderResubmitParamsWorkflow] `json:"workflow,required"`
}

func (r AccountHolderResubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

func (r AccountHolderResubmitParamsIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHolderResubmitParamsWorkflow string

const (
	AccountHolderResubmitParamsWorkflowKYCAdvanced AccountHolderResubmitParamsWorkflow = "KYC_ADVANCED"
)

func (r AccountHolderResubmitParamsWorkflow) IsKnown() bool {
	switch r {
	case AccountHolderResubmitParamsWorkflowKYCAdvanced:
		return true
	}
	return false
}

type AccountHolderSimulateEnrollmentDocumentReviewParams struct {
	// The account holder document upload which to perform the simulation upon.
	DocumentUploadToken param.Field[string] `json:"document_upload_token"`
	// An account holder document's upload status for use within the simulation.
	Status param.Field[AccountHolderSimulateEnrollmentDocumentReviewParamsStatus] `json:"status"`
	// Status reason that will be associated with the simulated account holder status.
	// Only required for a `REJECTED` status.
	StatusReasons param.Field[[]AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason] `json:"status_reasons"`
}

func (r AccountHolderSimulateEnrollmentDocumentReviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An account holder document's upload status for use within the simulation.
type AccountHolderSimulateEnrollmentDocumentReviewParamsStatus string

const (
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusUploaded AccountHolderSimulateEnrollmentDocumentReviewParamsStatus = "UPLOADED"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusAccepted AccountHolderSimulateEnrollmentDocumentReviewParamsStatus = "ACCEPTED"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusRejected AccountHolderSimulateEnrollmentDocumentReviewParamsStatus = "REJECTED"
)

func (r AccountHolderSimulateEnrollmentDocumentReviewParamsStatus) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentDocumentReviewParamsStatusUploaded, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusAccepted, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusRejected:
		return true
	}
	return false
}

type AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason string

const (
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentMissingRequiredData AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "DOCUMENT_MISSING_REQUIRED_DATA"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentUploadTooBlurry     AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "DOCUMENT_UPLOAD_TOO_BLURRY"
	AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidDocumentType         AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason = "INVALID_DOCUMENT_TYPE"
)

func (r AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentMissingRequiredData, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentUploadTooBlurry, AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidDocumentType:
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
	case AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityIDVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityAddressVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityNameVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonBlocklistAlertFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonIDVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonDobVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonControlPersonNameVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualDobVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualBlocklistAlertFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualIDVerificationFailure, AccountHolderSimulateEnrollmentReviewParamsStatusReasonBeneficialOwnerIndividualNameVerificationFailure:
		return true
	}
	return false
}

type AccountHolderUploadDocumentParams struct {
	// The type of document to upload
	DocumentType param.Field[AccountHolderUploadDocumentParamsDocumentType] `json:"document_type"`
	// Globally unique identifier for the entity.
	EntityToken param.Field[string] `json:"entity_token" format:"uuid"`
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
)

func (r AccountHolderUploadDocumentParamsDocumentType) IsKnown() bool {
	switch r {
	case AccountHolderUploadDocumentParamsDocumentTypeEinLetter, AccountHolderUploadDocumentParamsDocumentTypeTaxReturn, AccountHolderUploadDocumentParamsDocumentTypeOperatingAgreement, AccountHolderUploadDocumentParamsDocumentTypeCertificateOfFormation, AccountHolderUploadDocumentParamsDocumentTypeDriversLicense, AccountHolderUploadDocumentParamsDocumentTypePassport, AccountHolderUploadDocumentParamsDocumentTypePassportCard, AccountHolderUploadDocumentParamsDocumentTypeCertificateOfGoodStanding, AccountHolderUploadDocumentParamsDocumentTypeArticlesOfIncorporation, AccountHolderUploadDocumentParamsDocumentTypeArticlesOfOrganization, AccountHolderUploadDocumentParamsDocumentTypeBylaws, AccountHolderUploadDocumentParamsDocumentTypeGovernmentBusinessLicense, AccountHolderUploadDocumentParamsDocumentTypePartnershipAgreement, AccountHolderUploadDocumentParamsDocumentTypeSs4Form, AccountHolderUploadDocumentParamsDocumentTypeBankStatement, AccountHolderUploadDocumentParamsDocumentTypeUtilityBillStatement, AccountHolderUploadDocumentParamsDocumentTypeSsnCard, AccountHolderUploadDocumentParamsDocumentTypeItinLetter:
		return true
	}
	return false
}
