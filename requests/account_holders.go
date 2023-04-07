package requests

import (
	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
)

// 🚧 Unions are still being implemented.
type AccountHolderNewParams struct{}

func (r *AccountHolderNewParams) MarshalJSON() ([]byte, error) { return nil, nil }

type AccountHolderNewParamsBusinessEntity struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address field.Field[Address] `json:"address,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName field.Field[string] `json:"dba_business_name"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID field.Field[string] `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName field.Field[string] `json:"legal_business_name,required"`
	// Parent company name (if applicable).
	ParentCompany field.Field[string] `json:"parent_company"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers field.Field[[]string] `json:"phone_numbers,required"`
}

type AccountHolderNewParamsBeneficialOwnerEntities struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address field.Field[Address] `json:"address,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName field.Field[string] `json:"dba_business_name"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID field.Field[string] `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName field.Field[string] `json:"legal_business_name,required"`
	// Parent company name (if applicable).
	ParentCompany field.Field[string] `json:"parent_company"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers field.Field[[]string] `json:"phone_numbers,required"`
}

type AccountHolderNewParamsBeneficialOwnerIndividuals struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address field.Field[Address] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob field.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email field.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName field.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID field.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName field.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber field.Field[string] `json:"phone_number,required"`
}

type AccountHolderNewParamsControlPerson struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address field.Field[Address] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob field.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email field.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName field.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID field.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName field.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber field.Field[string] `json:"phone_number,required"`
}

type AccountHolderNewParamsWorkflow string

const (
	AccountHolderNewParamsWorkflowKYBBasic AccountHolderNewParamsWorkflow = "KYB_BASIC"
	AccountHolderNewParamsWorkflowKYBByo   AccountHolderNewParamsWorkflow = "KYB_BYO"
)

type AccountHolderNewParamsIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address field.Field[Address] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob field.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email field.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName field.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID field.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName field.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber field.Field[string] `json:"phone_number,required"`
}

type AccountHolderNewParamsKYCExemptionType string

const (
	AccountHolderNewParamsKYCExemptionTypeAuthorizedUser  AccountHolderNewParamsKYCExemptionType = "AUTHORIZED_USER"
	AccountHolderNewParamsKYCExemptionTypePrepaidCardUser AccountHolderNewParamsKYCExemptionType = "PREPAID_CARD_USER"
)

type AccountHolderUpdateParams struct {
	// Account holder's email address. The primary purpose of this field is for
	// cardholder identification and verification during the digital wallet
	// tokenization process.
	Email field.Field[string] `json:"email"`
	// Account holder's phone number, entered in E.164 format. The primary purpose of
	// this field is for cardholder identification and verification during the digital
	// wallet tokenization process.
	PhoneNumber field.Field[string] `json:"phone_number"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken field.Field[string] `json:"business_account_token"`
}

// MarshalJSON serializes AccountHolderUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountHolderUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type AccountHolderNewWebhookParams struct {
	// URL to receive webhook requests. Must be a valid HTTPS address.
	URL field.Field[string] `json:"url,required"`
}

// MarshalJSON serializes AccountHolderNewWebhookParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *AccountHolderNewWebhookParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type AccountHolderResubmitParams struct {
	Workflow field.Field[AccountHolderResubmitParamsWorkflow] `json:"workflow,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp field.Field[string] `json:"tos_timestamp,required"`
	// Information on individual for whom the account is being opened and KYC is being
	// re-run.
	Individual field.Field[AccountHolderResubmitParamsIndividual] `json:"individual,required"`
}

// MarshalJSON serializes AccountHolderResubmitParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountHolderResubmitParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type AccountHolderResubmitParamsWorkflow string

const (
	AccountHolderResubmitParamsWorkflowKYCAdvanced AccountHolderResubmitParamsWorkflow = "KYC_ADVANCED"
)

type AccountHolderResubmitParamsIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address field.Field[Address] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob field.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email field.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName field.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID field.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName field.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber field.Field[string] `json:"phone_number,required"`
}

type AccountHolderUploadDocumentParams struct {
	// Type of the document to upload.
	DocumentType field.Field[AccountHolderUploadDocumentParamsDocumentType] `json:"document_type,required"`
}

// MarshalJSON serializes AccountHolderUploadDocumentParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *AccountHolderUploadDocumentParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type AccountHolderUploadDocumentParamsDocumentType string

const (
	AccountHolderUploadDocumentParamsDocumentTypeCommercialLicense AccountHolderUploadDocumentParamsDocumentType = "commercial_license"
	AccountHolderUploadDocumentParamsDocumentTypeDriversLicense    AccountHolderUploadDocumentParamsDocumentType = "drivers_license"
	AccountHolderUploadDocumentParamsDocumentTypePassport          AccountHolderUploadDocumentParamsDocumentType = "passport"
	AccountHolderUploadDocumentParamsDocumentTypePassportCard      AccountHolderUploadDocumentParamsDocumentType = "passport_card"
	AccountHolderUploadDocumentParamsDocumentTypeVisa              AccountHolderUploadDocumentParamsDocumentType = "visa"
)
