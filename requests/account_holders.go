package requests

import (
	"fmt"

	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/core/fields"
	pjson "github.com/lithic-com/lithic-go/core/json"
)

// 🚧 Undiscriminated unions are still being implemented.
type AccountHolderNewParams struct{}

func (r *AccountHolderNewParams) MarshalJSON() ([]byte, error) { return nil, nil }

type KYB struct {
	// Information for business for which the account is being opened and KYB is being
	// run.
	BusinessEntity fields.Field[BusinessEntity] `json:"business_entity,required"`
	// List of all entities with >25% ownership in the company. If no entity or
	// individual owns >25% of the company, and the largest shareholder is an entity,
	// please identify them in this field. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background. If no business owner is an entity, pass in an
	// empty list. However, either this parameter or `beneficial_owner_individuals`
	// must be populated. on entities that should be included.
	BeneficialOwnerEntities fields.Field[[]BusinessEntity] `json:"beneficial_owner_entities,required"`
	// List of all individuals with >25% ownership in the company. If no entity or
	// individual owns >25% of the company, and the largest shareholder is an
	// individual, please identify them in this field. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included. If no
	// individual is an entity, pass in an empty list. However, either this parameter
	// or `beneficial_owner_entities` must be populated.
	BeneficialOwnerIndividuals fields.Field[[]Individual] `json:"beneficial_owner_individuals,required"`
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer, Managing Member, General Partner, President, Vice President, or
	// Treasurer). This can be an executive, or someone who will have program-wide
	// access to the cards that Lithic will provide. In some cases, this individual
	// could also be a beneficial owner listed above. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section II) for more background.
	ControlPerson fields.Field[Individual] `json:"control_person,required"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// business with a pass result.
	//
	// This field is required only if workflow type is `KYB_BYO`.
	KYBPassedTimestamp fields.Field[string] `json:"kyb_passed_timestamp"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness fields.Field[string] `json:"nature_of_business,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp fields.Field[string] `json:"tos_timestamp,required"`
	// Company website URL.
	WebsiteURL fields.Field[string] `json:"website_url,required"`
	// Specifies the type of KYB workflow to run.
	Workflow fields.Field[KYBWorkflow] `json:"workflow,required"`
}

func (r KYB) String() (result string) {
	return fmt.Sprintf("&KYB{BusinessEntity:%s BeneficialOwnerEntities:%s BeneficialOwnerIndividuals:%s ControlPerson:%s KYBPassedTimestamp:%s NatureOfBusiness:%s TosTimestamp:%s WebsiteURL:%s Workflow:%s}", r.BusinessEntity, core.Fmt(r.BeneficialOwnerEntities), core.Fmt(r.BeneficialOwnerIndividuals), r.ControlPerson, r.KYBPassedTimestamp, r.NatureOfBusiness, r.TosTimestamp, r.WebsiteURL, r.Workflow)
}

type BusinessEntity struct {
	// Business's physical address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable.
	Address fields.Field[Address] `json:"address,required"`
	// Any name that the business operates under that is not its legal business name
	// (if applicable).
	DbaBusinessName fields.Field[string] `json:"dba_business_name"`
	// Government-issued identification number. US Federal Employer Identification
	// Numbers (EIN) are currently supported, entered as full nine-digits, with or
	// without hyphens.
	GovernmentID fields.Field[string] `json:"government_id,required"`
	// Legal (formal) business name.
	LegalBusinessName fields.Field[string] `json:"legal_business_name,required"`
	// Parent company name (if applicable).
	ParentCompany fields.Field[string] `json:"parent_company"`
	// One or more of the business's phone number(s), entered as a list in E.164
	// format.
	PhoneNumbers fields.Field[[]string] `json:"phone_numbers,required"`
}

func (r BusinessEntity) String() (result string) {
	return fmt.Sprintf("&BusinessEntity{Address:%s DbaBusinessName:%s GovernmentID:%s LegalBusinessName:%s ParentCompany:%s PhoneNumbers:%s}", r.Address, r.DbaBusinessName, r.GovernmentID, r.LegalBusinessName, r.ParentCompany, core.Fmt(r.PhoneNumbers))
}

type Individual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address fields.Field[Address] `json:"address,required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob fields.Field[string] `json:"dob,required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email fields.Field[string] `json:"email,required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName fields.Field[string] `json:"first_name,required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID fields.Field[string] `json:"government_id,required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName fields.Field[string] `json:"last_name,required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber fields.Field[string] `json:"phone_number,required"`
}

func (r Individual) String() (result string) {
	return fmt.Sprintf("&Individual{Address:%s Dob:%s Email:%s FirstName:%s GovernmentID:%s LastName:%s PhoneNumber:%s}", r.Address, r.Dob, r.Email, r.FirstName, r.GovernmentID, r.LastName, r.PhoneNumber)
}

type KYBWorkflow string

const (
	KYBWorkflowKYBBasic KYBWorkflow = "KYB_BASIC"
	KYBWorkflowKYBByo   KYBWorkflow = "KYB_BYO"
)

type KYC struct {
	// Information on individual for whom the account is being opened and KYC is being
	// run.
	Individual fields.Field[Individual] `json:"individual,required"`
	// An RFC 3339 timestamp indicating when precomputed KYC was completed on the
	// individual with a pass result.
	//
	// This field is required only if workflow type is `KYC_BYO`.
	KYCPassedTimestamp fields.Field[string] `json:"kyc_passed_timestamp"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp fields.Field[string] `json:"tos_timestamp,required"`
	// Specifies the type of KYC workflow to run.
	Workflow fields.Field[KYCWorkflow] `json:"workflow,required"`
}

func (r KYC) String() (result string) {
	return fmt.Sprintf("&KYC{Individual:%s KYCPassedTimestamp:%s TosTimestamp:%s Workflow:%s}", r.Individual, r.KYCPassedTimestamp, r.TosTimestamp, r.Workflow)
}

type KYCWorkflow string

const (
	KYCWorkflowKYCAdvanced KYCWorkflow = "KYC_ADVANCED"
	KYCWorkflowKYCBasic    KYCWorkflow = "KYC_BASIC"
	KYCWorkflowKYCByo      KYCWorkflow = "KYC_BYO"
)

type KYCExempt struct {
	// Specifies the workflow type. This must be 'KYC_EXEMPT'
	Workflow fields.Field[KYCExemptWorkflow] `json:"workflow,required"`
	// Specifies the type of KYC Exempt user
	KYCExemptionType fields.Field[KYCExemptKYCExemptionType] `json:"kyc_exemption_type,required"`
	// The KYC Exempt user's first name
	FirstName fields.Field[string] `json:"first_name,required"`
	// The KYC Exempt user's last name
	LastName fields.Field[string] `json:"last_name,required"`
	// The KYC Exempt user's email
	Email fields.Field[string] `json:"email,required"`
	// The KYC Exempt user's phone number
	PhoneNumber fields.Field[string] `json:"phone_number,required"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken fields.Field[string] `json:"business_account_token"`
	// KYC Exempt user's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address fields.Field[Address] `json:"address"`
}

func (r KYCExempt) String() (result string) {
	return fmt.Sprintf("&KYCExempt{Workflow:%s KYCExemptionType:%s FirstName:%s LastName:%s Email:%s PhoneNumber:%s BusinessAccountToken:%s Address:%s}", r.Workflow, r.KYCExemptionType, r.FirstName, r.LastName, r.Email, r.PhoneNumber, r.BusinessAccountToken, r.Address)
}

type KYCExemptWorkflow string

const (
	KYCExemptWorkflowKYCExempt KYCExemptWorkflow = "KYC_EXEMPT"
)

type KYCExemptKYCExemptionType string

const (
	KYCExemptKYCExemptionTypeAuthorizedUser  KYCExemptKYCExemptionType = "AUTHORIZED_USER"
	KYCExemptKYCExemptionTypePrepaidCardUser KYCExemptKYCExemptionType = "PREPAID_CARD_USER"
)

type AccountHolderUpdateParams struct {
	// Account holder's email address. The primary purpose of this field is for
	// cardholder identification and verification during the digital wallet
	// tokenization process.
	Email fields.Field[string] `json:"email"`
	// Account holder's phone number, entered in E.164 format. The primary purpose of
	// this field is for cardholder identification and verification during the digital
	// wallet tokenization process.
	PhoneNumber fields.Field[string] `json:"phone_number"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Pass the account_token of the enrolled business associated
	// with the AUTHORIZED_USER in this field.
	BusinessAccountToken fields.Field[string] `json:"business_account_token"`
}

// MarshalJSON serializes AccountHolderUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountHolderUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountHolderUpdateParams) String() (result string) {
	return fmt.Sprintf("&AccountHolderUpdateParams{Email:%s PhoneNumber:%s BusinessAccountToken:%s}", r.Email, r.PhoneNumber, r.BusinessAccountToken)
}

type AccountHolderNewWebhookParams struct {
	// URL to receive webhook requests. Must be a valid HTTPS address.
	URL fields.Field[string] `json:"url,required"`
}

// MarshalJSON serializes AccountHolderNewWebhookParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *AccountHolderNewWebhookParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountHolderNewWebhookParams) String() (result string) {
	return fmt.Sprintf("&AccountHolderNewWebhookParams{URL:%s}", r.URL)
}

type AccountHolderResubmitParams struct {
	Workflow fields.Field[AccountHolderResubmitParamsWorkflow] `json:"workflow,required"`
	// An RFC 3339 timestamp indicating when the account holder accepted the applicable
	// legal agreements (e.g., cardholder terms) as agreed upon during API customer's
	// implementation with Lithic.
	TosTimestamp fields.Field[string] `json:"tos_timestamp,required"`
	// Information on individual for whom the account is being opened and KYC is being
	// re-run.
	Individual fields.Field[Individual] `json:"individual,required"`
}

// MarshalJSON serializes AccountHolderResubmitParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountHolderResubmitParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountHolderResubmitParams) String() (result string) {
	return fmt.Sprintf("&AccountHolderResubmitParams{Workflow:%s TosTimestamp:%s Individual:%s}", r.Workflow, r.TosTimestamp, r.Individual)
}

type AccountHolderResubmitParamsWorkflow string

const (
	AccountHolderResubmitParamsWorkflowKYCAdvanced AccountHolderResubmitParamsWorkflow = "KYC_ADVANCED"
)

type AccountHoldersGetDocumentParams struct {
	// Globally unique identifier for the account holder.
	AccountHolderToken string `path:"account_holder_token" json:"-"`
}

type AccountHolderUploadDocumentParams struct {
	// Type of the document to upload.
	DocumentType fields.Field[AccountHolderUploadDocumentParamsDocumentType] `json:"document_type,required"`
}

// MarshalJSON serializes AccountHolderUploadDocumentParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *AccountHolderUploadDocumentParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountHolderUploadDocumentParams) String() (result string) {
	return fmt.Sprintf("&AccountHolderUploadDocumentParams{DocumentType:%s}", r.DocumentType)
}

type AccountHolderUploadDocumentParamsDocumentType string

const (
	AccountHolderUploadDocumentParamsDocumentTypeCommercialLicense AccountHolderUploadDocumentParamsDocumentType = "commercial_license"
	AccountHolderUploadDocumentParamsDocumentTypeDriversLicense    AccountHolderUploadDocumentParamsDocumentType = "drivers_license"
	AccountHolderUploadDocumentParamsDocumentTypePassport          AccountHolderUploadDocumentParamsDocumentType = "passport"
	AccountHolderUploadDocumentParamsDocumentTypePassportCard      AccountHolderUploadDocumentParamsDocumentType = "passport_card"
	AccountHolderUploadDocumentParamsDocumentTypeVisa              AccountHolderUploadDocumentParamsDocumentType = "visa"
)
