package requests

import (
	"github.com/lithic-com/lithic-go/core/field"
	apijson "github.com/lithic-com/lithic-go/core/json"
)

type KYB struct {
	// Information for business for which the account is being opened and KYB is being
	// run.
	BusinessEntity field.Field[KYBBusinessEntity] `json:"business_entity,required"`
	// List of all entities with >25% ownership in the company. If no entity or
	// individual owns >25% of the company, and the largest shareholder is an entity,
	// please identify them in this field. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background. If no business owner is an entity, pass in an
	// empty list. However, either this parameter or `beneficial_owner_individuals`
	// must be populated. on entities that should be included.
	BeneficialOwnerEntities field.Field[[]KYBBeneficialOwnerEntities] `json:"beneficial_owner_entities,required"`
	// List of all individuals with >25% ownership in the company. If no entity or
	// individual owns >25% of the company, and the largest shareholder is an
	// individual, please identify them in this field. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included. If no
	// individual is an entity, pass in an empty list. However, either this parameter
	// or `beneficial_owner_entities` must be populated.
	BeneficialOwnerIndividuals field.Field[[]KYBBeneficialOwnerIndividuals] `json:"beneficial_owner_individuals,required"`
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer, Managing Member, General Partner, President, Vice President, or
	// Treasurer). This can be an executive, or someone who will have program-wide
	// access to the cards that Lithic will provide. In some cases, this individual
	// could also be a beneficial owner listed above. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section II) for more background.
	ControlPerson field.Field[KYBControlPerson] `json:"control_person,required"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// business with a pass result.
	//
	// This field is required only if workflow type is `KYB_BYO`.
	KYBPassedTimestamp field.Field[string] `json:"kyb_passed_timestamp"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness field.Field[string] `json:"nature_of_business,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp field.Field[string] `json:"tos_timestamp,required"`
	// Company website URL.
	WebsiteURL field.Field[string] `json:"website_url,required"`
	// Specifies the type of KYB workflow to run.
	Workflow field.Field[KYBWorkflow] `json:"workflow,required"`
}

// MarshalJSON serializes KYB into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r KYB) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYB) implementsRequestsAccountHolderNewParams() {}

type KYBBusinessEntity struct {
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

// MarshalJSON serializes KYBBusinessEntity into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r KYBBusinessEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBBeneficialOwnerEntities struct {
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

// MarshalJSON serializes KYBBeneficialOwnerEntities into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r KYBBeneficialOwnerEntities) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBBeneficialOwnerIndividuals struct {
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

// MarshalJSON serializes KYBBeneficialOwnerIndividuals into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r KYBBeneficialOwnerIndividuals) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBControlPerson struct {
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

// MarshalJSON serializes KYBControlPerson into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r KYBControlPerson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYBWorkflow string

const (
	KYBWorkflowKYBBasic KYBWorkflow = "KYB_BASIC"
	KYBWorkflowKYBByo   KYBWorkflow = "KYB_BYO"
)

type KYC struct {
	// Information on individual for whom the account is being opened and KYC is being
	// run.
	Individual field.Field[KYCIndividual] `json:"individual,required"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// individual with a pass result.
	//
	// This field is required only if workflow type is `KYC_BYO`.
	KYCPassedTimestamp field.Field[string] `json:"kyc_passed_timestamp"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp field.Field[string] `json:"tos_timestamp,required"`
	// Specifies the type of KYC workflow to run.
	Workflow field.Field[KYCWorkflow] `json:"workflow,required"`
}

// MarshalJSON serializes KYC into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r KYC) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYC) implementsRequestsAccountHolderNewParams() {}

type KYCIndividual struct {
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

// MarshalJSON serializes KYCIndividual into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r KYCIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KYCWorkflow string

const (
	KYCWorkflowKYCAdvanced KYCWorkflow = "KYC_ADVANCED"
	KYCWorkflowKYCBasic    KYCWorkflow = "KYC_BASIC"
	KYCWorkflowKYCByo      KYCWorkflow = "KYC_BYO"
)

type KYCExempt struct {
	// Specifies the workflow type. This must be 'KYC_EXEMPT'
	Workflow field.Field[KYCExemptWorkflow] `json:"workflow,required"`
	// Specifies the type of KYC Exempt user
	KYCExemptionType field.Field[KYCExemptKYCExemptionType] `json:"kyc_exemption_type,required"`
	// The KYC Exempt user's first name
	FirstName field.Field[string] `json:"first_name,required"`
	// The KYC Exempt user's last name
	LastName field.Field[string] `json:"last_name,required"`
	// The KYC Exempt user's email
	Email field.Field[string] `json:"email,required"`
	// The KYC Exempt user's phone number
	PhoneNumber field.Field[string] `json:"phone_number,required"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken field.Field[string] `json:"business_account_token"`
	// KYC Exempt user's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address field.Field[Address] `json:"address"`
}

// MarshalJSON serializes KYCExempt into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r KYCExempt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KYCExempt) implementsRequestsAccountHolderNewParams() {}

type KYCExemptWorkflow string

const (
	KYCExemptWorkflowKYCExempt KYCExemptWorkflow = "KYC_EXEMPT"
)

type KYCExemptKYCExemptionType string

const (
	KYCExemptKYCExemptionTypeAuthorizedUser  KYCExemptKYCExemptionType = "AUTHORIZED_USER"
	KYCExemptKYCExemptionTypePrepaidCardUser KYCExemptKYCExemptionType = "PREPAID_CARD_USER"
)

// This interface is a union satisfied by one of the following: [KYB], [KYC],
// [KYCExempt].
type AccountHolderNewParams interface {
	implementsRequestsAccountHolderNewParams()
}

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
func (r AccountHolderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHolderNewWebhookParams struct {
	// URL to receive webhook requests. Must be a valid HTTPS address.
	URL field.Field[string] `json:"url,required"`
}

// MarshalJSON serializes AccountHolderNewWebhookParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r AccountHolderNewWebhookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
func (r AccountHolderResubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
