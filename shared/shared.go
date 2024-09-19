// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/tidwall/gjson"
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

type AuthRule struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Indicates whether the Auth Rule is ACTIVE or INACTIVE
	State AuthRuleState `json:"state,required"`
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens []string `json:"account_tokens"`
	// Countries in which the Auth Rule permits transactions. Note that Lithic
	// maintains a list of countries in which all transactions are blocked; "allowing"
	// those countries in an Auth Rule does not override the Lithic-wide restrictions.
	AllowedCountries []string `json:"allowed_countries"`
	// Merchant category codes for which the Auth Rule permits transactions.
	AllowedMcc []string `json:"allowed_mcc"`
	// Countries in which the Auth Rule automatically declines transactions.
	BlockedCountries []string `json:"blocked_countries"`
	// Merchant category codes for which the Auth Rule automatically declines
	// transactions.
	BlockedMcc []string `json:"blocked_mcc"`
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens []string `json:"card_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel bool         `json:"program_level"`
	JSON         authRuleJSON `json:"-"`
}

// authRuleJSON contains the JSON metadata for the struct [AuthRule]
type authRuleJSON struct {
	Token            apijson.Field
	State            apijson.Field
	AccountTokens    apijson.Field
	AllowedCountries apijson.Field
	AllowedMcc       apijson.Field
	BlockedCountries apijson.Field
	BlockedMcc       apijson.Field
	CardTokens       apijson.Field
	ProgramLevel     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AuthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the Auth Rule is ACTIVE or INACTIVE
type AuthRuleState string

const (
	AuthRuleStateActive   AuthRuleState = "ACTIVE"
	AuthRuleStateInactive AuthRuleState = "INACTIVE"
)

func (r AuthRuleState) IsKnown() bool {
	switch r {
	case AuthRuleStateActive, AuthRuleStateInactive:
		return true
	}
	return false
}

type CarrierParam struct {
	// QR code url to display on the card carrier
	QrCodeURL param.Field[string] `json:"qr_code_url"`
}

func (r CarrierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ISO 4217 currency. Its enumerants are ISO 4217 currencies except for some
// special currencies like “XXX`. Enumerants names are lowercase cureency code
// e.g. :attr:`Currency.eur`, :attr:`Currency.usd`.
type Currency string

const (
	CurrencyAed Currency = "AED"
	CurrencyAfn Currency = "AFN"
	CurrencyAll Currency = "ALL"
	CurrencyAmd Currency = "AMD"
	CurrencyAng Currency = "ANG"
	CurrencyAoa Currency = "AOA"
	CurrencyArs Currency = "ARS"
	CurrencyAud Currency = "AUD"
	CurrencyAwg Currency = "AWG"
	CurrencyAzn Currency = "AZN"
	CurrencyBam Currency = "BAM"
	CurrencyBbd Currency = "BBD"
	CurrencyBdt Currency = "BDT"
	CurrencyBgn Currency = "BGN"
	CurrencyBhd Currency = "BHD"
	CurrencyBif Currency = "BIF"
	CurrencyBmd Currency = "BMD"
	CurrencyBnd Currency = "BND"
	CurrencyBob Currency = "BOB"
	CurrencyBov Currency = "BOV"
	CurrencyBrl Currency = "BRL"
	CurrencyBsd Currency = "BSD"
	CurrencyBtn Currency = "BTN"
	CurrencyBwp Currency = "BWP"
	CurrencyByn Currency = "BYN"
	CurrencyBzd Currency = "BZD"
	CurrencyCad Currency = "CAD"
	CurrencyCdf Currency = "CDF"
	CurrencyChe Currency = "CHE"
	CurrencyChf Currency = "CHF"
	CurrencyChw Currency = "CHW"
	CurrencyClf Currency = "CLF"
	CurrencyClp Currency = "CLP"
	CurrencyCny Currency = "CNY"
	CurrencyCop Currency = "COP"
	CurrencyCou Currency = "COU"
	CurrencyCrc Currency = "CRC"
	CurrencyCuc Currency = "CUC"
	CurrencyCup Currency = "CUP"
	CurrencyCve Currency = "CVE"
	CurrencyCzk Currency = "CZK"
	CurrencyDjf Currency = "DJF"
	CurrencyDkk Currency = "DKK"
	CurrencyDop Currency = "DOP"
	CurrencyDzd Currency = "DZD"
	CurrencyEgp Currency = "EGP"
	CurrencyErn Currency = "ERN"
	CurrencyEtb Currency = "ETB"
	CurrencyEur Currency = "EUR"
	CurrencyFjd Currency = "FJD"
	CurrencyFkp Currency = "FKP"
	CurrencyGbp Currency = "GBP"
	CurrencyGel Currency = "GEL"
	CurrencyGhs Currency = "GHS"
	CurrencyGip Currency = "GIP"
	CurrencyGmd Currency = "GMD"
	CurrencyGnf Currency = "GNF"
	CurrencyGtq Currency = "GTQ"
	CurrencyGyd Currency = "GYD"
	CurrencyHkd Currency = "HKD"
	CurrencyHnl Currency = "HNL"
	CurrencyHrk Currency = "HRK"
	CurrencyHtg Currency = "HTG"
	CurrencyHuf Currency = "HUF"
	CurrencyIdr Currency = "IDR"
	CurrencyIls Currency = "ILS"
	CurrencyInr Currency = "INR"
	CurrencyIqd Currency = "IQD"
	CurrencyIrr Currency = "IRR"
	CurrencyIsk Currency = "ISK"
	CurrencyJmd Currency = "JMD"
	CurrencyJod Currency = "JOD"
	CurrencyJpy Currency = "JPY"
	CurrencyKes Currency = "KES"
	CurrencyKgs Currency = "KGS"
	CurrencyKhr Currency = "KHR"
	CurrencyKmf Currency = "KMF"
	CurrencyKpw Currency = "KPW"
	CurrencyKrw Currency = "KRW"
	CurrencyKwd Currency = "KWD"
	CurrencyKyd Currency = "KYD"
	CurrencyKzt Currency = "KZT"
	CurrencyLak Currency = "LAK"
	CurrencyLbp Currency = "LBP"
	CurrencyLkr Currency = "LKR"
	CurrencyLrd Currency = "LRD"
	CurrencyLsl Currency = "LSL"
	CurrencyLyd Currency = "LYD"
	CurrencyMad Currency = "MAD"
	CurrencyMdl Currency = "MDL"
	CurrencyMga Currency = "MGA"
	CurrencyMkd Currency = "MKD"
	CurrencyMmk Currency = "MMK"
	CurrencyMnt Currency = "MNT"
	CurrencyMop Currency = "MOP"
	CurrencyMru Currency = "MRU"
	CurrencyMur Currency = "MUR"
	CurrencyMvr Currency = "MVR"
	CurrencyMwk Currency = "MWK"
	CurrencyMxn Currency = "MXN"
	CurrencyMxv Currency = "MXV"
	CurrencyMyr Currency = "MYR"
	CurrencyMzn Currency = "MZN"
	CurrencyNad Currency = "NAD"
	CurrencyNgn Currency = "NGN"
	CurrencyNio Currency = "NIO"
	CurrencyNok Currency = "NOK"
	CurrencyNpr Currency = "NPR"
	CurrencyNzd Currency = "NZD"
	CurrencyOmr Currency = "OMR"
	CurrencyPab Currency = "PAB"
	CurrencyPen Currency = "PEN"
	CurrencyPgk Currency = "PGK"
	CurrencyPhp Currency = "PHP"
	CurrencyPkr Currency = "PKR"
	CurrencyPln Currency = "PLN"
	CurrencyPyg Currency = "PYG"
	CurrencyQar Currency = "QAR"
	CurrencyRon Currency = "RON"
	CurrencyRsd Currency = "RSD"
	CurrencyRub Currency = "RUB"
	CurrencyRwf Currency = "RWF"
	CurrencySar Currency = "SAR"
	CurrencySbd Currency = "SBD"
	CurrencyScr Currency = "SCR"
	CurrencySdg Currency = "SDG"
	CurrencySek Currency = "SEK"
	CurrencySgd Currency = "SGD"
	CurrencyShp Currency = "SHP"
	CurrencySle Currency = "SLE"
	CurrencySll Currency = "SLL"
	CurrencySos Currency = "SOS"
	CurrencySrd Currency = "SRD"
	CurrencySsp Currency = "SSP"
	CurrencyStn Currency = "STN"
	CurrencySvc Currency = "SVC"
	CurrencySyp Currency = "SYP"
	CurrencySzl Currency = "SZL"
	CurrencyThb Currency = "THB"
	CurrencyTjs Currency = "TJS"
	CurrencyTmt Currency = "TMT"
	CurrencyTnd Currency = "TND"
	CurrencyTop Currency = "TOP"
	CurrencyTry Currency = "TRY"
	CurrencyTtd Currency = "TTD"
	CurrencyTwd Currency = "TWD"
	CurrencyTzs Currency = "TZS"
	CurrencyUah Currency = "UAH"
	CurrencyUgx Currency = "UGX"
	CurrencyUsd Currency = "USD"
	CurrencyUsn Currency = "USN"
	CurrencyUyi Currency = "UYI"
	CurrencyUyu Currency = "UYU"
	CurrencyUyw Currency = "UYW"
	CurrencyUzs Currency = "UZS"
	CurrencyVed Currency = "VED"
	CurrencyVes Currency = "VES"
	CurrencyVnd Currency = "VND"
	CurrencyVuv Currency = "VUV"
	CurrencyWst Currency = "WST"
	CurrencyXaf Currency = "XAF"
	CurrencyXag Currency = "XAG"
	CurrencyXau Currency = "XAU"
	CurrencyXba Currency = "XBA"
	CurrencyXbb Currency = "XBB"
	CurrencyXbc Currency = "XBC"
	CurrencyXbd Currency = "XBD"
	CurrencyXcd Currency = "XCD"
	CurrencyXdr Currency = "XDR"
	CurrencyXof Currency = "XOF"
	CurrencyXpd Currency = "XPD"
	CurrencyXpf Currency = "XPF"
	CurrencyXpt Currency = "XPT"
	CurrencyXsu Currency = "XSU"
	CurrencyXts Currency = "XTS"
	CurrencyXua Currency = "XUA"
	CurrencyXxx Currency = "XXX"
	CurrencyYer Currency = "YER"
	CurrencyZar Currency = "ZAR"
	CurrencyZmw Currency = "ZMW"
	CurrencyZwl Currency = "ZWL"
)

func (r Currency) IsKnown() bool {
	switch r {
	case CurrencyAed, CurrencyAfn, CurrencyAll, CurrencyAmd, CurrencyAng, CurrencyAoa, CurrencyArs, CurrencyAud, CurrencyAwg, CurrencyAzn, CurrencyBam, CurrencyBbd, CurrencyBdt, CurrencyBgn, CurrencyBhd, CurrencyBif, CurrencyBmd, CurrencyBnd, CurrencyBob, CurrencyBov, CurrencyBrl, CurrencyBsd, CurrencyBtn, CurrencyBwp, CurrencyByn, CurrencyBzd, CurrencyCad, CurrencyCdf, CurrencyChe, CurrencyChf, CurrencyChw, CurrencyClf, CurrencyClp, CurrencyCny, CurrencyCop, CurrencyCou, CurrencyCrc, CurrencyCuc, CurrencyCup, CurrencyCve, CurrencyCzk, CurrencyDjf, CurrencyDkk, CurrencyDop, CurrencyDzd, CurrencyEgp, CurrencyErn, CurrencyEtb, CurrencyEur, CurrencyFjd, CurrencyFkp, CurrencyGbp, CurrencyGel, CurrencyGhs, CurrencyGip, CurrencyGmd, CurrencyGnf, CurrencyGtq, CurrencyGyd, CurrencyHkd, CurrencyHnl, CurrencyHrk, CurrencyHtg, CurrencyHuf, CurrencyIdr, CurrencyIls, CurrencyInr, CurrencyIqd, CurrencyIrr, CurrencyIsk, CurrencyJmd, CurrencyJod, CurrencyJpy, CurrencyKes, CurrencyKgs, CurrencyKhr, CurrencyKmf, CurrencyKpw, CurrencyKrw, CurrencyKwd, CurrencyKyd, CurrencyKzt, CurrencyLak, CurrencyLbp, CurrencyLkr, CurrencyLrd, CurrencyLsl, CurrencyLyd, CurrencyMad, CurrencyMdl, CurrencyMga, CurrencyMkd, CurrencyMmk, CurrencyMnt, CurrencyMop, CurrencyMru, CurrencyMur, CurrencyMvr, CurrencyMwk, CurrencyMxn, CurrencyMxv, CurrencyMyr, CurrencyMzn, CurrencyNad, CurrencyNgn, CurrencyNio, CurrencyNok, CurrencyNpr, CurrencyNzd, CurrencyOmr, CurrencyPab, CurrencyPen, CurrencyPgk, CurrencyPhp, CurrencyPkr, CurrencyPln, CurrencyPyg, CurrencyQar, CurrencyRon, CurrencyRsd, CurrencyRub, CurrencyRwf, CurrencySar, CurrencySbd, CurrencyScr, CurrencySdg, CurrencySek, CurrencySgd, CurrencyShp, CurrencySle, CurrencySll, CurrencySos, CurrencySrd, CurrencySsp, CurrencyStn, CurrencySvc, CurrencySyp, CurrencySzl, CurrencyThb, CurrencyTjs, CurrencyTmt, CurrencyTnd, CurrencyTop, CurrencyTry, CurrencyTtd, CurrencyTwd, CurrencyTzs, CurrencyUah, CurrencyUgx, CurrencyUsd, CurrencyUsn, CurrencyUyi, CurrencyUyu, CurrencyUyw, CurrencyUzs, CurrencyVed, CurrencyVes, CurrencyVnd, CurrencyVuv, CurrencyWst, CurrencyXaf, CurrencyXag, CurrencyXau, CurrencyXba, CurrencyXbb, CurrencyXbc, CurrencyXbd, CurrencyXcd, CurrencyXdr, CurrencyXof, CurrencyXpd, CurrencyXpf, CurrencyXpt, CurrencyXsu, CurrencyXts, CurrencyXua, CurrencyXxx, CurrencyYer, CurrencyZar, CurrencyZmw, CurrencyZwl:
		return true
	}
	return false
}

// Describes the document and the required document image uploads required to
// re-run KYC
type Document struct {
	// Globally unique identifier for the document.
	Token string `json:"token,required" format:"uuid"`
	// Globally unique identifier for the account holder.
	AccountHolderToken string `json:"account_holder_token,required" format:"uuid"`
	// Type of documentation to be submitted for verification.
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

// Type of documentation to be submitted for verification.
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
)

func (r DocumentDocumentType) IsKnown() bool {
	switch r {
	case DocumentDocumentTypeDriversLicense, DocumentDocumentTypePassport, DocumentDocumentTypePassportCard, DocumentDocumentTypeEinLetter, DocumentDocumentTypeTaxReturn, DocumentDocumentTypeOperatingAgreement, DocumentDocumentTypeCertificateOfFormation, DocumentDocumentTypeCertificateOfGoodStanding, DocumentDocumentTypeArticlesOfIncorporation, DocumentDocumentTypeArticlesOfOrganization, DocumentDocumentTypeBylaws, DocumentDocumentTypeGovernmentBusinessLicense, DocumentDocumentTypePartnershipAgreement, DocumentDocumentTypeSs4Form, DocumentDocumentTypeBankStatement, DocumentDocumentTypeUtilityBillStatement, DocumentDocumentTypeSsnCard, DocumentDocumentTypeItinLetter:
		return true
	}
	return false
}

// Represents a single image of the document to upload.
type DocumentRequiredDocumentUpload struct {
	// Globally unique identifier for the document upload.
	Token string `json:"token,required" format:"uuid"`
	// Type of image to upload.
	ImageType DocumentRequiredDocumentUploadsImageType `json:"image_type,required"`
	// Status of document image upload.
	Status DocumentRequiredDocumentUploadsStatus `json:"status,required"`
	// Reasons for document image upload status.
	StatusReasons []DocumentRequiredDocumentUploadsStatusReason `json:"status_reasons,required"`
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
	Token         apijson.Field
	ImageType     apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	UploadURL     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Status of document image upload.
type DocumentRequiredDocumentUploadsStatus string

const (
	DocumentRequiredDocumentUploadsStatusAccepted      DocumentRequiredDocumentUploadsStatus = "ACCEPTED"
	DocumentRequiredDocumentUploadsStatusRejected      DocumentRequiredDocumentUploadsStatus = "REJECTED"
	DocumentRequiredDocumentUploadsStatusPendingUpload DocumentRequiredDocumentUploadsStatus = "PENDING_UPLOAD"
	DocumentRequiredDocumentUploadsStatusUploaded      DocumentRequiredDocumentUploadsStatus = "UPLOADED"
)

func (r DocumentRequiredDocumentUploadsStatus) IsKnown() bool {
	switch r {
	case DocumentRequiredDocumentUploadsStatusAccepted, DocumentRequiredDocumentUploadsStatusRejected, DocumentRequiredDocumentUploadsStatusPendingUpload, DocumentRequiredDocumentUploadsStatusUploaded:
		return true
	}
	return false
}

type DocumentRequiredDocumentUploadsStatusReason string

const (
	DocumentRequiredDocumentUploadsStatusReasonDocumentMissingRequiredData DocumentRequiredDocumentUploadsStatusReason = "DOCUMENT_MISSING_REQUIRED_DATA"
	DocumentRequiredDocumentUploadsStatusReasonDocumentUploadTooBlurry     DocumentRequiredDocumentUploadsStatusReason = "DOCUMENT_UPLOAD_TOO_BLURRY"
	DocumentRequiredDocumentUploadsStatusReasonFileSizeTooLarge            DocumentRequiredDocumentUploadsStatusReason = "FILE_SIZE_TOO_LARGE"
	DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentType         DocumentRequiredDocumentUploadsStatusReason = "INVALID_DOCUMENT_TYPE"
	DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentUpload       DocumentRequiredDocumentUploadsStatusReason = "INVALID_DOCUMENT_UPLOAD"
	DocumentRequiredDocumentUploadsStatusReasonUnknownError                DocumentRequiredDocumentUploadsStatusReason = "UNKNOWN_ERROR"
)

func (r DocumentRequiredDocumentUploadsStatusReason) IsKnown() bool {
	switch r {
	case DocumentRequiredDocumentUploadsStatusReasonDocumentMissingRequiredData, DocumentRequiredDocumentUploadsStatusReasonDocumentUploadTooBlurry, DocumentRequiredDocumentUploadsStatusReasonFileSizeTooLarge, DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentType, DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentUpload, DocumentRequiredDocumentUploadsStatusReasonUnknownError:
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

type VelocityLimitParams struct {
	Filters VelocityLimitParamsFilters `json:"filters,required"`
	// The size of the trailing window to calculate Spend Velocity over in seconds.
	Period VelocityLimitParamsPeriodUnion `json:"period,required"`
	Scope  VelocityLimitParamsScope       `json:"scope,required"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount float64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount float64                 `json:"limit_count,nullable"`
	JSON       velocityLimitParamsJSON `json:"-"`
}

// velocityLimitParamsJSON contains the JSON metadata for the struct
// [VelocityLimitParams]
type velocityLimitParamsJSON struct {
	Filters     apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VelocityLimitParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitParamsJSON) RawJSON() string {
	return r.raw
}

func (r VelocityLimitParams) ImplementsAuthRuleV2NewResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2NewResponseDraftVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2GetResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2GetResponseDraftVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2UpdateResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2UpdateResponseDraftVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2ListResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2ListResponseDraftVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2ApplyResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2ApplyResponseDraftVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2DraftResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2DraftResponseDraftVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2PromoteResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) ImplementsAuthRuleV2PromoteResponseDraftVersionParameters() {}

type VelocityLimitParamsFilters struct {
	// ISO-3166-1 alpha-3 Country Codes to include in the velocity calculation.
	// Transactions not matching any of the provided will not be included in the
	// calculated velocity.
	IncludeCountries []string `json:"include_countries,nullable"`
	// Merchant Category Codes to include in the velocity calculation. Transactions not
	// matching this MCC will not be included in the calculated velocity.
	IncludeMccs []string                       `json:"include_mccs,nullable"`
	JSON        velocityLimitParamsFiltersJSON `json:"-"`
}

// velocityLimitParamsFiltersJSON contains the JSON metadata for the struct
// [VelocityLimitParamsFilters]
type velocityLimitParamsFiltersJSON struct {
	IncludeCountries apijson.Field
	IncludeMccs      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VelocityLimitParamsFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitParamsFiltersJSON) RawJSON() string {
	return r.raw
}

// The size of the trailing window to calculate Spend Velocity over in seconds.
//
// Union satisfied by [shared.UnionFloat] or
// [shared.VelocityLimitParamsPeriodWindow].
type VelocityLimitParamsPeriodUnion interface {
	ImplementsSharedVelocityLimitParamsPeriodUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VelocityLimitParamsPeriodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(VelocityLimitParamsPeriodWindow("")),
		},
	)
}

type VelocityLimitParamsScope string

const (
	VelocityLimitParamsScopeCard    VelocityLimitParamsScope = "CARD"
	VelocityLimitParamsScopeAccount VelocityLimitParamsScope = "ACCOUNT"
)

func (r VelocityLimitParamsScope) IsKnown() bool {
	switch r {
	case VelocityLimitParamsScopeCard, VelocityLimitParamsScopeAccount:
		return true
	}
	return false
}

// The window of time to calculate Spend Velocity over.
//
//   - `DAY`: Velocity over the current day since midnight Eastern Time.
//   - `MONTH`: Velocity over the current month since 00:00 / 12 AM on the first of
//     the month in Eastern Time.
type VelocityLimitParamsPeriodWindow string

const (
	VelocityLimitParamsPeriodWindowDay   VelocityLimitParamsPeriodWindow = "DAY"
	VelocityLimitParamsPeriodWindowMonth VelocityLimitParamsPeriodWindow = "MONTH"
)

func (r VelocityLimitParamsPeriodWindow) IsKnown() bool {
	switch r {
	case VelocityLimitParamsPeriodWindowDay, VelocityLimitParamsPeriodWindowMonth:
		return true
	}
	return false
}

func (r VelocityLimitParamsPeriodWindow) ImplementsSharedVelocityLimitParamsPeriodUnion() {}
