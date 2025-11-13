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

// Financial Event
type FinancialEvent struct {
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result FinancialEventResult `json:"result"`
	Type   FinancialEventType   `json:"type"`
	JSON   financialEventJSON   `json:"-"`
}

// financialEventJSON contains the JSON metadata for the struct [FinancialEvent]
type financialEventJSON struct {
	Token       apijson.Field
	Amount      apijson.Field
	Created     apijson.Field
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinancialEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialEventJSON) RawJSON() string {
	return r.raw
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type FinancialEventResult string

const (
	FinancialEventResultApproved FinancialEventResult = "APPROVED"
	FinancialEventResultDeclined FinancialEventResult = "DECLINED"
)

func (r FinancialEventResult) IsKnown() bool {
	switch r {
	case FinancialEventResultApproved, FinancialEventResultDeclined:
		return true
	}
	return false
}

type FinancialEventType string

const (
	FinancialEventTypeACHOriginationCancelled      FinancialEventType = "ACH_ORIGINATION_CANCELLED"
	FinancialEventTypeACHOriginationInitiated      FinancialEventType = "ACH_ORIGINATION_INITIATED"
	FinancialEventTypeACHOriginationProcessed      FinancialEventType = "ACH_ORIGINATION_PROCESSED"
	FinancialEventTypeACHOriginationReleased       FinancialEventType = "ACH_ORIGINATION_RELEASED"
	FinancialEventTypeACHOriginationRejected       FinancialEventType = "ACH_ORIGINATION_REJECTED"
	FinancialEventTypeACHOriginationReviewed       FinancialEventType = "ACH_ORIGINATION_REVIEWED"
	FinancialEventTypeACHOriginationSettled        FinancialEventType = "ACH_ORIGINATION_SETTLED"
	FinancialEventTypeACHReceiptProcessed          FinancialEventType = "ACH_RECEIPT_PROCESSED"
	FinancialEventTypeACHReceiptReleased           FinancialEventType = "ACH_RECEIPT_RELEASED"
	FinancialEventTypeACHReceiptSettled            FinancialEventType = "ACH_RECEIPT_SETTLED"
	FinancialEventTypeACHReturnInitiated           FinancialEventType = "ACH_RETURN_INITIATED"
	FinancialEventTypeACHReturnProcessed           FinancialEventType = "ACH_RETURN_PROCESSED"
	FinancialEventTypeACHReturnRejected            FinancialEventType = "ACH_RETURN_REJECTED"
	FinancialEventTypeACHReturnSettled             FinancialEventType = "ACH_RETURN_SETTLED"
	FinancialEventTypeAuthorization                FinancialEventType = "AUTHORIZATION"
	FinancialEventTypeAuthorizationAdvice          FinancialEventType = "AUTHORIZATION_ADVICE"
	FinancialEventTypeAuthorizationExpiry          FinancialEventType = "AUTHORIZATION_EXPIRY"
	FinancialEventTypeAuthorizationReversal        FinancialEventType = "AUTHORIZATION_REVERSAL"
	FinancialEventTypeBalanceInquiry               FinancialEventType = "BALANCE_INQUIRY"
	FinancialEventTypeBillingError                 FinancialEventType = "BILLING_ERROR"
	FinancialEventTypeBillingErrorReversal         FinancialEventType = "BILLING_ERROR_REVERSAL"
	FinancialEventTypeCardToCard                   FinancialEventType = "CARD_TO_CARD"
	FinancialEventTypeCashBack                     FinancialEventType = "CASH_BACK"
	FinancialEventTypeCashBackReversal             FinancialEventType = "CASH_BACK_REVERSAL"
	FinancialEventTypeClearing                     FinancialEventType = "CLEARING"
	FinancialEventTypeCollection                   FinancialEventType = "COLLECTION"
	FinancialEventTypeCorrectionCredit             FinancialEventType = "CORRECTION_CREDIT"
	FinancialEventTypeCorrectionDebit              FinancialEventType = "CORRECTION_DEBIT"
	FinancialEventTypeCreditAuthorization          FinancialEventType = "CREDIT_AUTHORIZATION"
	FinancialEventTypeCreditAuthorizationAdvice    FinancialEventType = "CREDIT_AUTHORIZATION_ADVICE"
	FinancialEventTypeCurrencyConversion           FinancialEventType = "CURRENCY_CONVERSION"
	FinancialEventTypeCurrencyConversionReversal   FinancialEventType = "CURRENCY_CONVERSION_REVERSAL"
	FinancialEventTypeDisputeWon                   FinancialEventType = "DISPUTE_WON"
	FinancialEventTypeExternalACHCanceled          FinancialEventType = "EXTERNAL_ACH_CANCELED"
	FinancialEventTypeExternalACHInitiated         FinancialEventType = "EXTERNAL_ACH_INITIATED"
	FinancialEventTypeExternalACHReleased          FinancialEventType = "EXTERNAL_ACH_RELEASED"
	FinancialEventTypeExternalACHReversed          FinancialEventType = "EXTERNAL_ACH_REVERSED"
	FinancialEventTypeExternalACHSettled           FinancialEventType = "EXTERNAL_ACH_SETTLED"
	FinancialEventTypeExternalCheckCanceled        FinancialEventType = "EXTERNAL_CHECK_CANCELED"
	FinancialEventTypeExternalCheckInitiated       FinancialEventType = "EXTERNAL_CHECK_INITIATED"
	FinancialEventTypeExternalCheckReleased        FinancialEventType = "EXTERNAL_CHECK_RELEASED"
	FinancialEventTypeExternalCheckReversed        FinancialEventType = "EXTERNAL_CHECK_REVERSED"
	FinancialEventTypeExternalCheckSettled         FinancialEventType = "EXTERNAL_CHECK_SETTLED"
	FinancialEventTypeExternalTransferCanceled     FinancialEventType = "EXTERNAL_TRANSFER_CANCELED"
	FinancialEventTypeExternalTransferInitiated    FinancialEventType = "EXTERNAL_TRANSFER_INITIATED"
	FinancialEventTypeExternalTransferReleased     FinancialEventType = "EXTERNAL_TRANSFER_RELEASED"
	FinancialEventTypeExternalTransferReversed     FinancialEventType = "EXTERNAL_TRANSFER_REVERSED"
	FinancialEventTypeExternalTransferSettled      FinancialEventType = "EXTERNAL_TRANSFER_SETTLED"
	FinancialEventTypeExternalWireCanceled         FinancialEventType = "EXTERNAL_WIRE_CANCELED"
	FinancialEventTypeExternalWireInitiated        FinancialEventType = "EXTERNAL_WIRE_INITIATED"
	FinancialEventTypeExternalWireReleased         FinancialEventType = "EXTERNAL_WIRE_RELEASED"
	FinancialEventTypeExternalWireReversed         FinancialEventType = "EXTERNAL_WIRE_REVERSED"
	FinancialEventTypeExternalWireSettled          FinancialEventType = "EXTERNAL_WIRE_SETTLED"
	FinancialEventTypeFinancialAuthorization       FinancialEventType = "FINANCIAL_AUTHORIZATION"
	FinancialEventTypeFinancialCreditAuthorization FinancialEventType = "FINANCIAL_CREDIT_AUTHORIZATION"
	FinancialEventTypeInterest                     FinancialEventType = "INTEREST"
	FinancialEventTypeInterestReversal             FinancialEventType = "INTEREST_REVERSAL"
	FinancialEventTypeInternalAdjustment           FinancialEventType = "INTERNAL_ADJUSTMENT"
	FinancialEventTypeLatePayment                  FinancialEventType = "LATE_PAYMENT"
	FinancialEventTypeLatePaymentReversal          FinancialEventType = "LATE_PAYMENT_REVERSAL"
	FinancialEventTypeLossWriteOff                 FinancialEventType = "LOSS_WRITE_OFF"
	FinancialEventTypeProvisionalCredit            FinancialEventType = "PROVISIONAL_CREDIT"
	FinancialEventTypeProvisionalCreditReversal    FinancialEventType = "PROVISIONAL_CREDIT_REVERSAL"
	FinancialEventTypeService                      FinancialEventType = "SERVICE"
	FinancialEventTypeReturn                       FinancialEventType = "RETURN"
	FinancialEventTypeReturnReversal               FinancialEventType = "RETURN_REVERSAL"
	FinancialEventTypeTransfer                     FinancialEventType = "TRANSFER"
	FinancialEventTypeTransferInsufficientFunds    FinancialEventType = "TRANSFER_INSUFFICIENT_FUNDS"
	FinancialEventTypeReturnedPayment              FinancialEventType = "RETURNED_PAYMENT"
	FinancialEventTypeReturnedPaymentReversal      FinancialEventType = "RETURNED_PAYMENT_REVERSAL"
	FinancialEventTypeLithicNetworkPayment         FinancialEventType = "LITHIC_NETWORK_PAYMENT"
)

func (r FinancialEventType) IsKnown() bool {
	switch r {
	case FinancialEventTypeACHOriginationCancelled, FinancialEventTypeACHOriginationInitiated, FinancialEventTypeACHOriginationProcessed, FinancialEventTypeACHOriginationReleased, FinancialEventTypeACHOriginationRejected, FinancialEventTypeACHOriginationReviewed, FinancialEventTypeACHOriginationSettled, FinancialEventTypeACHReceiptProcessed, FinancialEventTypeACHReceiptReleased, FinancialEventTypeACHReceiptSettled, FinancialEventTypeACHReturnInitiated, FinancialEventTypeACHReturnProcessed, FinancialEventTypeACHReturnRejected, FinancialEventTypeACHReturnSettled, FinancialEventTypeAuthorization, FinancialEventTypeAuthorizationAdvice, FinancialEventTypeAuthorizationExpiry, FinancialEventTypeAuthorizationReversal, FinancialEventTypeBalanceInquiry, FinancialEventTypeBillingError, FinancialEventTypeBillingErrorReversal, FinancialEventTypeCardToCard, FinancialEventTypeCashBack, FinancialEventTypeCashBackReversal, FinancialEventTypeClearing, FinancialEventTypeCollection, FinancialEventTypeCorrectionCredit, FinancialEventTypeCorrectionDebit, FinancialEventTypeCreditAuthorization, FinancialEventTypeCreditAuthorizationAdvice, FinancialEventTypeCurrencyConversion, FinancialEventTypeCurrencyConversionReversal, FinancialEventTypeDisputeWon, FinancialEventTypeExternalACHCanceled, FinancialEventTypeExternalACHInitiated, FinancialEventTypeExternalACHReleased, FinancialEventTypeExternalACHReversed, FinancialEventTypeExternalACHSettled, FinancialEventTypeExternalCheckCanceled, FinancialEventTypeExternalCheckInitiated, FinancialEventTypeExternalCheckReleased, FinancialEventTypeExternalCheckReversed, FinancialEventTypeExternalCheckSettled, FinancialEventTypeExternalTransferCanceled, FinancialEventTypeExternalTransferInitiated, FinancialEventTypeExternalTransferReleased, FinancialEventTypeExternalTransferReversed, FinancialEventTypeExternalTransferSettled, FinancialEventTypeExternalWireCanceled, FinancialEventTypeExternalWireInitiated, FinancialEventTypeExternalWireReleased, FinancialEventTypeExternalWireReversed, FinancialEventTypeExternalWireSettled, FinancialEventTypeFinancialAuthorization, FinancialEventTypeFinancialCreditAuthorization, FinancialEventTypeInterest, FinancialEventTypeInterestReversal, FinancialEventTypeInternalAdjustment, FinancialEventTypeLatePayment, FinancialEventTypeLatePaymentReversal, FinancialEventTypeLossWriteOff, FinancialEventTypeProvisionalCredit, FinancialEventTypeProvisionalCreditReversal, FinancialEventTypeService, FinancialEventTypeReturn, FinancialEventTypeReturnReversal, FinancialEventTypeTransfer, FinancialEventTypeTransferInsufficientFunds, FinancialEventTypeReturnedPayment, FinancialEventTypeReturnedPaymentReversal, FinancialEventTypeLithicNetworkPayment:
		return true
	}
	return false
}

type Merchant struct {
	// Unique alphanumeric identifier for the payment card acceptor (merchant).
	AcceptorID string `json:"acceptor_id,required"`
	// Unique numeric identifier of the acquiring institution.
	AcquiringInstitutionID string `json:"acquiring_institution_id,required"`
	// City of card acceptor. Note that in many cases, particularly in card-not-present
	// transactions, merchants may send through a phone number or URL in this field.
	City string `json:"city,required"`
	// Country or entity of card acceptor. Possible values are: (1) all ISO 3166-1
	// alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for Netherlands Antilles.
	Country string `json:"country,required"`
	// Short description of card acceptor.
	Descriptor string `json:"descriptor,required"`
	// Merchant category code (MCC). A four-digit number listed in ISO 18245. An MCC is
	// used to classify a business by the types of goods or services it provides.
	Mcc string `json:"mcc,required"`
	// Geographic state of card acceptor.
	State string       `json:"state,required"`
	JSON  merchantJSON `json:"-"`
}

// merchantJSON contains the JSON metadata for the struct [Merchant]
type merchantJSON struct {
	AcceptorID             apijson.Field
	AcquiringInstitutionID apijson.Field
	City                   apijson.Field
	Country                apijson.Field
	Descriptor             apijson.Field
	Mcc                    apijson.Field
	State                  apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *Merchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r merchantJSON) RawJSON() string {
	return r.raw
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
