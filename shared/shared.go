// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
)

type Address struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code, entered in uppercase ISO 3166-1 alpha-3 three-character
	// format. Only USA is currently supported for all workflows. KYC_EXEMPT supports
	// CAN additionally.
	Country string `json:"country,required"`
	// Valid postal code. USA postal codes (ZIP codes) are supported, entered as a
	// five-digit postal code or nine-digit postal code (ZIP+4) using the format
	// 12345-1234. KYC_EXEMPT supports Canadian postal codes.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. USA state codes are supported, entered in uppercase ISO 3166-2
	// two-character format. KYC_EXEMPT supports Canadian province codes.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string      `json:"address2"`
	JSON     addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressParam struct {
	// Valid deliverable address (no PO boxes).
	Address1 param.Field[string] `json:"address1,required"`
	// Name of city.
	City param.Field[string] `json:"city,required"`
	// Valid country code, entered in uppercase ISO 3166-1 alpha-3 three-character
	// format. Only USA is currently supported for all workflows. KYC_EXEMPT supports
	// CAN additionally.
	Country param.Field[string] `json:"country,required"`
	// Valid postal code. USA postal codes (ZIP codes) are supported, entered as a
	// five-digit postal code or nine-digit postal code (ZIP+4) using the format
	// 12345-1234. KYC_EXEMPT supports Canadian postal codes.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Valid state code. USA state codes are supported, entered in uppercase ISO 3166-2
	// two-character format. KYC_EXEMPT supports Canadian province codes.
	State param.Field[string] `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 param.Field[string] `json:"address2"`
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CarrierParam struct {
	// QR code url to display on the card carrier
	QrCodeURL param.Field[string] `json:"qr_code_url"`
}

func (r CarrierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Currency = string

// Describes the document and the required document image uploads required to
// re-run KYC
type Document struct {
	// Globally unique identifier for the document.
	Token string `json:"token,required" format:"uuid"`
	// Globally unique identifier for the account holder.
	AccountHolderToken string `json:"account_holder_token,required" format:"uuid"`
	// Type of documentation to be submitted for verification of an account holder
	DocumentType DocumentDocumentType `json:"document_type,required"`
	// Globally unique identifier for an entity.
	EntityToken string `json:"entity_token,required" format:"uuid"`
	// Represents a single image of the document to upload.
	RequiredDocumentUploads []DocumentRequiredDocumentUpload `json:"required_document_uploads,required"`
	JSON                    documentJSON                     `json:"-"`
}

// documentJSON contains the JSON metadata for the struct [Document]
type documentJSON struct {
	Token                   apijson.Field
	AccountHolderToken      apijson.Field
	DocumentType            apijson.Field
	EntityToken             apijson.Field
	RequiredDocumentUploads apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentJSON) RawJSON() string {
	return r.raw
}

// Type of documentation to be submitted for verification of an account holder
type DocumentDocumentType string

const (
	DocumentDocumentTypeDriversLicense            DocumentDocumentType = "DRIVERS_LICENSE"
	DocumentDocumentTypePassport                  DocumentDocumentType = "PASSPORT"
	DocumentDocumentTypePassportCard              DocumentDocumentType = "PASSPORT_CARD"
	DocumentDocumentTypeEinLetter                 DocumentDocumentType = "EIN_LETTER"
	DocumentDocumentTypeTaxReturn                 DocumentDocumentType = "TAX_RETURN"
	DocumentDocumentTypeOperatingAgreement        DocumentDocumentType = "OPERATING_AGREEMENT"
	DocumentDocumentTypeCertificateOfFormation    DocumentDocumentType = "CERTIFICATE_OF_FORMATION"
	DocumentDocumentTypeCertificateOfGoodStanding DocumentDocumentType = "CERTIFICATE_OF_GOOD_STANDING"
	DocumentDocumentTypeArticlesOfIncorporation   DocumentDocumentType = "ARTICLES_OF_INCORPORATION"
	DocumentDocumentTypeArticlesOfOrganization    DocumentDocumentType = "ARTICLES_OF_ORGANIZATION"
	DocumentDocumentTypeBylaws                    DocumentDocumentType = "BYLAWS"
	DocumentDocumentTypeGovernmentBusinessLicense DocumentDocumentType = "GOVERNMENT_BUSINESS_LICENSE"
	DocumentDocumentTypePartnershipAgreement      DocumentDocumentType = "PARTNERSHIP_AGREEMENT"
	DocumentDocumentTypeSs4Form                   DocumentDocumentType = "SS4_FORM"
	DocumentDocumentTypeBankStatement             DocumentDocumentType = "BANK_STATEMENT"
	DocumentDocumentTypeUtilityBillStatement      DocumentDocumentType = "UTILITY_BILL_STATEMENT"
	DocumentDocumentTypeSsnCard                   DocumentDocumentType = "SSN_CARD"
	DocumentDocumentTypeItinLetter                DocumentDocumentType = "ITIN_LETTER"
	DocumentDocumentTypeFincenBoiReport           DocumentDocumentType = "FINCEN_BOI_REPORT"
)

func (r DocumentDocumentType) IsKnown() bool {
	switch r {
	case DocumentDocumentTypeDriversLicense, DocumentDocumentTypePassport, DocumentDocumentTypePassportCard, DocumentDocumentTypeEinLetter, DocumentDocumentTypeTaxReturn, DocumentDocumentTypeOperatingAgreement, DocumentDocumentTypeCertificateOfFormation, DocumentDocumentTypeCertificateOfGoodStanding, DocumentDocumentTypeArticlesOfIncorporation, DocumentDocumentTypeArticlesOfOrganization, DocumentDocumentTypeBylaws, DocumentDocumentTypeGovernmentBusinessLicense, DocumentDocumentTypePartnershipAgreement, DocumentDocumentTypeSs4Form, DocumentDocumentTypeBankStatement, DocumentDocumentTypeUtilityBillStatement, DocumentDocumentTypeSsnCard, DocumentDocumentTypeItinLetter, DocumentDocumentTypeFincenBoiReport:
		return true
	}
	return false
}

// Represents a single image of the document to upload.
type DocumentRequiredDocumentUpload struct {
	// Globally unique identifier for the document upload.
	Token string `json:"token,required" format:"uuid"`
	// A list of status reasons associated with a KYB account holder that have been
	// satisfied by the document upload
	AcceptedEntityStatusReasons []string `json:"accepted_entity_status_reasons,required"`
	// When the document upload was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Type of image to upload.
	ImageType DocumentRequiredDocumentUploadsImageType `json:"image_type,required"`
	// A list of status reasons associated with a KYB account holder that have not been
	// satisfied by the document upload
	RejectedEntityStatusReasons []string `json:"rejected_entity_status_reasons,required"`
	// Status of an account holder's document upload.
	Status DocumentRequiredDocumentUploadsStatus `json:"status,required"`
	// Reasons for document image upload status.
	StatusReasons []DocumentRequiredDocumentUploadsStatusReason `json:"status_reasons,required"`
	// When the document upload was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// URL to upload document image to.
	//
	// Note that the upload URLs expire after 7 days. If an upload URL expires, you can
	// refresh the URLs by retrieving the document upload from
	// `GET /account_holders/{account_holder_token}/documents`.
	UploadURL string                             `json:"upload_url,required"`
	JSON      documentRequiredDocumentUploadJSON `json:"-"`
}

// documentRequiredDocumentUploadJSON contains the JSON metadata for the struct
// [DocumentRequiredDocumentUpload]
type documentRequiredDocumentUploadJSON struct {
	Token                       apijson.Field
	AcceptedEntityStatusReasons apijson.Field
	Created                     apijson.Field
	ImageType                   apijson.Field
	RejectedEntityStatusReasons apijson.Field
	Status                      apijson.Field
	StatusReasons               apijson.Field
	Updated                     apijson.Field
	UploadURL                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *DocumentRequiredDocumentUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentRequiredDocumentUploadJSON) RawJSON() string {
	return r.raw
}

// Type of image to upload.
type DocumentRequiredDocumentUploadsImageType string

const (
	DocumentRequiredDocumentUploadsImageTypeFront DocumentRequiredDocumentUploadsImageType = "FRONT"
	DocumentRequiredDocumentUploadsImageTypeBack  DocumentRequiredDocumentUploadsImageType = "BACK"
)

func (r DocumentRequiredDocumentUploadsImageType) IsKnown() bool {
	switch r {
	case DocumentRequiredDocumentUploadsImageTypeFront, DocumentRequiredDocumentUploadsImageTypeBack:
		return true
	}
	return false
}

// Status of an account holder's document upload.
type DocumentRequiredDocumentUploadsStatus string

const (
	DocumentRequiredDocumentUploadsStatusAccepted        DocumentRequiredDocumentUploadsStatus = "ACCEPTED"
	DocumentRequiredDocumentUploadsStatusRejected        DocumentRequiredDocumentUploadsStatus = "REJECTED"
	DocumentRequiredDocumentUploadsStatusPendingUpload   DocumentRequiredDocumentUploadsStatus = "PENDING_UPLOAD"
	DocumentRequiredDocumentUploadsStatusUploaded        DocumentRequiredDocumentUploadsStatus = "UPLOADED"
	DocumentRequiredDocumentUploadsStatusPartialApproval DocumentRequiredDocumentUploadsStatus = "PARTIAL_APPROVAL"
)

func (r DocumentRequiredDocumentUploadsStatus) IsKnown() bool {
	switch r {
	case DocumentRequiredDocumentUploadsStatusAccepted, DocumentRequiredDocumentUploadsStatusRejected, DocumentRequiredDocumentUploadsStatusPendingUpload, DocumentRequiredDocumentUploadsStatusUploaded, DocumentRequiredDocumentUploadsStatusPartialApproval:
		return true
	}
	return false
}

// The status reasons for an account holder document upload that is not ACCEPTED
type DocumentRequiredDocumentUploadsStatusReason string

const (
	DocumentRequiredDocumentUploadsStatusReasonDocumentMissingRequiredData     DocumentRequiredDocumentUploadsStatusReason = "DOCUMENT_MISSING_REQUIRED_DATA"
	DocumentRequiredDocumentUploadsStatusReasonDocumentUploadTooBlurry         DocumentRequiredDocumentUploadsStatusReason = "DOCUMENT_UPLOAD_TOO_BLURRY"
	DocumentRequiredDocumentUploadsStatusReasonFileSizeTooLarge                DocumentRequiredDocumentUploadsStatusReason = "FILE_SIZE_TOO_LARGE"
	DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentType             DocumentRequiredDocumentUploadsStatusReason = "INVALID_DOCUMENT_TYPE"
	DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentUpload           DocumentRequiredDocumentUploadsStatusReason = "INVALID_DOCUMENT_UPLOAD"
	DocumentRequiredDocumentUploadsStatusReasonInvalidEntity                   DocumentRequiredDocumentUploadsStatusReason = "INVALID_ENTITY"
	DocumentRequiredDocumentUploadsStatusReasonDocumentExpired                 DocumentRequiredDocumentUploadsStatusReason = "DOCUMENT_EXPIRED"
	DocumentRequiredDocumentUploadsStatusReasonDocumentIssuedGreaterThan30Days DocumentRequiredDocumentUploadsStatusReason = "DOCUMENT_ISSUED_GREATER_THAN_30_DAYS"
	DocumentRequiredDocumentUploadsStatusReasonDocumentTypeNotSupported        DocumentRequiredDocumentUploadsStatusReason = "DOCUMENT_TYPE_NOT_SUPPORTED"
	DocumentRequiredDocumentUploadsStatusReasonUnknownFailureReason            DocumentRequiredDocumentUploadsStatusReason = "UNKNOWN_FAILURE_REASON"
	DocumentRequiredDocumentUploadsStatusReasonUnknownError                    DocumentRequiredDocumentUploadsStatusReason = "UNKNOWN_ERROR"
)

func (r DocumentRequiredDocumentUploadsStatusReason) IsKnown() bool {
	switch r {
	case DocumentRequiredDocumentUploadsStatusReasonDocumentMissingRequiredData, DocumentRequiredDocumentUploadsStatusReasonDocumentUploadTooBlurry, DocumentRequiredDocumentUploadsStatusReasonFileSizeTooLarge, DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentType, DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentUpload, DocumentRequiredDocumentUploadsStatusReasonInvalidEntity, DocumentRequiredDocumentUploadsStatusReasonDocumentExpired, DocumentRequiredDocumentUploadsStatusReasonDocumentIssuedGreaterThan30Days, DocumentRequiredDocumentUploadsStatusReasonDocumentTypeNotSupported, DocumentRequiredDocumentUploadsStatusReasonUnknownFailureReason, DocumentRequiredDocumentUploadsStatusReasonUnknownError:
		return true
	}
	return false
}

type ShippingAddressParam struct {
	// Valid USPS routable address.
	Address1 param.Field[string] `json:"address1,required"`
	// City
	City param.Field[string] `json:"city,required"`
	// Uppercase ISO 3166-1 alpha-3 three character abbreviation.
	Country param.Field[string] `json:"country,required"`
	// Customer's first name. This will be the first name printed on the physical card.
	// The combined length of `first_name` and `last_name` may not exceed 25
	// characters.
	FirstName param.Field[string] `json:"first_name,required"`
	// Customer's surname (family name). This will be the last name printed on the
	// physical card. The combined length of `first_name` and `last_name` may not
	// exceed 25 characters.
	LastName param.Field[string] `json:"last_name,required"`
	// Postal code (formerly zipcode). For US addresses, either five-digit postal code
	// or nine-digit postal code (ZIP+4) using the format 12345-1234.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Uppercase ISO 3166-2 two character abbreviation for US and CA. Optional with a
	// limit of 24 characters for other countries.
	State param.Field[string] `json:"state,required"`
	// Unit number (if applicable).
	Address2 param.Field[string] `json:"address2"`
	// Email address to be contacted for expedited shipping process purposes. Required
	// if `shipping_method` is `EXPEDITED`.
	Email param.Field[string] `json:"email"`
	// Text to be printed on line two of the physical card. Use of this field requires
	// additional permissions.
	Line2Text param.Field[string] `json:"line2_text"`
	// Cardholder's phone number in E.164 format to be contacted for expedited shipping
	// process purposes. Required if `shipping_method` is `EXPEDITED`.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r ShippingAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
