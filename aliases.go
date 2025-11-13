// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"github.com/lithic-com/lithic-go/internal/apierror"
	"github.com/lithic-com/lithic-go/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type Address = shared.Address

// This is an alias to an internal type.
type AddressParam = shared.AddressParam

// This is an alias to an internal type.
type CarrierParam = shared.CarrierParam

// 3-character alphabetic ISO 4217 currency
//
// This is an alias to an internal type.
type Currency = shared.Currency

// Describes the document and the required document image uploads required to
// re-run KYC
//
// This is an alias to an internal type.
type Document = shared.Document

// Type of documentation to be submitted for verification of an account holder
//
// This is an alias to an internal type.
type DocumentDocumentType = shared.DocumentDocumentType

// This is an alias to an internal value.
const DocumentDocumentTypeDriversLicense = shared.DocumentDocumentTypeDriversLicense

// This is an alias to an internal value.
const DocumentDocumentTypePassport = shared.DocumentDocumentTypePassport

// This is an alias to an internal value.
const DocumentDocumentTypePassportCard = shared.DocumentDocumentTypePassportCard

// This is an alias to an internal value.
const DocumentDocumentTypeEinLetter = shared.DocumentDocumentTypeEinLetter

// This is an alias to an internal value.
const DocumentDocumentTypeTaxReturn = shared.DocumentDocumentTypeTaxReturn

// This is an alias to an internal value.
const DocumentDocumentTypeOperatingAgreement = shared.DocumentDocumentTypeOperatingAgreement

// This is an alias to an internal value.
const DocumentDocumentTypeCertificateOfFormation = shared.DocumentDocumentTypeCertificateOfFormation

// This is an alias to an internal value.
const DocumentDocumentTypeCertificateOfGoodStanding = shared.DocumentDocumentTypeCertificateOfGoodStanding

// This is an alias to an internal value.
const DocumentDocumentTypeArticlesOfIncorporation = shared.DocumentDocumentTypeArticlesOfIncorporation

// This is an alias to an internal value.
const DocumentDocumentTypeArticlesOfOrganization = shared.DocumentDocumentTypeArticlesOfOrganization

// This is an alias to an internal value.
const DocumentDocumentTypeBylaws = shared.DocumentDocumentTypeBylaws

// This is an alias to an internal value.
const DocumentDocumentTypeGovernmentBusinessLicense = shared.DocumentDocumentTypeGovernmentBusinessLicense

// This is an alias to an internal value.
const DocumentDocumentTypePartnershipAgreement = shared.DocumentDocumentTypePartnershipAgreement

// This is an alias to an internal value.
const DocumentDocumentTypeSs4Form = shared.DocumentDocumentTypeSs4Form

// This is an alias to an internal value.
const DocumentDocumentTypeBankStatement = shared.DocumentDocumentTypeBankStatement

// This is an alias to an internal value.
const DocumentDocumentTypeUtilityBillStatement = shared.DocumentDocumentTypeUtilityBillStatement

// This is an alias to an internal value.
const DocumentDocumentTypeSsnCard = shared.DocumentDocumentTypeSsnCard

// This is an alias to an internal value.
const DocumentDocumentTypeItinLetter = shared.DocumentDocumentTypeItinLetter

// This is an alias to an internal value.
const DocumentDocumentTypeFincenBoiReport = shared.DocumentDocumentTypeFincenBoiReport

// Represents a single image of the document to upload.
//
// This is an alias to an internal type.
type DocumentRequiredDocumentUpload = shared.DocumentRequiredDocumentUpload

// Type of image to upload.
//
// This is an alias to an internal type.
type DocumentRequiredDocumentUploadsImageType = shared.DocumentRequiredDocumentUploadsImageType

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsImageTypeFront = shared.DocumentRequiredDocumentUploadsImageTypeFront

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsImageTypeBack = shared.DocumentRequiredDocumentUploadsImageTypeBack

// Status of an account holder's document upload.
//
// This is an alias to an internal type.
type DocumentRequiredDocumentUploadsStatus = shared.DocumentRequiredDocumentUploadsStatus

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusAccepted = shared.DocumentRequiredDocumentUploadsStatusAccepted

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusRejected = shared.DocumentRequiredDocumentUploadsStatusRejected

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusPendingUpload = shared.DocumentRequiredDocumentUploadsStatusPendingUpload

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusUploaded = shared.DocumentRequiredDocumentUploadsStatusUploaded

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusPartialApproval = shared.DocumentRequiredDocumentUploadsStatusPartialApproval

// The status reasons for an account holder document upload that is not ACCEPTED
//
// This is an alias to an internal type.
type DocumentRequiredDocumentUploadsStatusReason = shared.DocumentRequiredDocumentUploadsStatusReason

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonDocumentMissingRequiredData = shared.DocumentRequiredDocumentUploadsStatusReasonDocumentMissingRequiredData

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonDocumentUploadTooBlurry = shared.DocumentRequiredDocumentUploadsStatusReasonDocumentUploadTooBlurry

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonFileSizeTooLarge = shared.DocumentRequiredDocumentUploadsStatusReasonFileSizeTooLarge

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentType = shared.DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentType

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentUpload = shared.DocumentRequiredDocumentUploadsStatusReasonInvalidDocumentUpload

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonInvalidEntity = shared.DocumentRequiredDocumentUploadsStatusReasonInvalidEntity

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonDocumentExpired = shared.DocumentRequiredDocumentUploadsStatusReasonDocumentExpired

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonDocumentIssuedGreaterThan30Days = shared.DocumentRequiredDocumentUploadsStatusReasonDocumentIssuedGreaterThan30Days

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonDocumentTypeNotSupported = shared.DocumentRequiredDocumentUploadsStatusReasonDocumentTypeNotSupported

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonUnknownFailureReason = shared.DocumentRequiredDocumentUploadsStatusReasonUnknownFailureReason

// This is an alias to an internal value.
const DocumentRequiredDocumentUploadsStatusReasonUnknownError = shared.DocumentRequiredDocumentUploadsStatusReasonUnknownError

// Financial Event
//
// This is an alias to an internal type.
type FinancialEvent = shared.FinancialEvent

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
//
// This is an alias to an internal type.
type FinancialEventResult = shared.FinancialEventResult

// This is an alias to an internal value.
const FinancialEventResultApproved = shared.FinancialEventResultApproved

// This is an alias to an internal value.
const FinancialEventResultDeclined = shared.FinancialEventResultDeclined

// This is an alias to an internal type.
type FinancialEventType = shared.FinancialEventType

// This is an alias to an internal value.
const FinancialEventTypeACHOriginationCancelled = shared.FinancialEventTypeACHOriginationCancelled

// This is an alias to an internal value.
const FinancialEventTypeACHOriginationInitiated = shared.FinancialEventTypeACHOriginationInitiated

// This is an alias to an internal value.
const FinancialEventTypeACHOriginationProcessed = shared.FinancialEventTypeACHOriginationProcessed

// This is an alias to an internal value.
const FinancialEventTypeACHOriginationReleased = shared.FinancialEventTypeACHOriginationReleased

// This is an alias to an internal value.
const FinancialEventTypeACHOriginationRejected = shared.FinancialEventTypeACHOriginationRejected

// This is an alias to an internal value.
const FinancialEventTypeACHOriginationReviewed = shared.FinancialEventTypeACHOriginationReviewed

// This is an alias to an internal value.
const FinancialEventTypeACHOriginationSettled = shared.FinancialEventTypeACHOriginationSettled

// This is an alias to an internal value.
const FinancialEventTypeACHReceiptProcessed = shared.FinancialEventTypeACHReceiptProcessed

// This is an alias to an internal value.
const FinancialEventTypeACHReceiptReleased = shared.FinancialEventTypeACHReceiptReleased

// This is an alias to an internal value.
const FinancialEventTypeACHReceiptSettled = shared.FinancialEventTypeACHReceiptSettled

// This is an alias to an internal value.
const FinancialEventTypeACHReturnInitiated = shared.FinancialEventTypeACHReturnInitiated

// This is an alias to an internal value.
const FinancialEventTypeACHReturnProcessed = shared.FinancialEventTypeACHReturnProcessed

// This is an alias to an internal value.
const FinancialEventTypeACHReturnRejected = shared.FinancialEventTypeACHReturnRejected

// This is an alias to an internal value.
const FinancialEventTypeACHReturnSettled = shared.FinancialEventTypeACHReturnSettled

// This is an alias to an internal value.
const FinancialEventTypeAuthorization = shared.FinancialEventTypeAuthorization

// This is an alias to an internal value.
const FinancialEventTypeAuthorizationAdvice = shared.FinancialEventTypeAuthorizationAdvice

// This is an alias to an internal value.
const FinancialEventTypeAuthorizationExpiry = shared.FinancialEventTypeAuthorizationExpiry

// This is an alias to an internal value.
const FinancialEventTypeAuthorizationReversal = shared.FinancialEventTypeAuthorizationReversal

// This is an alias to an internal value.
const FinancialEventTypeBalanceInquiry = shared.FinancialEventTypeBalanceInquiry

// This is an alias to an internal value.
const FinancialEventTypeBillingError = shared.FinancialEventTypeBillingError

// This is an alias to an internal value.
const FinancialEventTypeBillingErrorReversal = shared.FinancialEventTypeBillingErrorReversal

// This is an alias to an internal value.
const FinancialEventTypeCardToCard = shared.FinancialEventTypeCardToCard

// This is an alias to an internal value.
const FinancialEventTypeCashBack = shared.FinancialEventTypeCashBack

// This is an alias to an internal value.
const FinancialEventTypeCashBackReversal = shared.FinancialEventTypeCashBackReversal

// This is an alias to an internal value.
const FinancialEventTypeClearing = shared.FinancialEventTypeClearing

// This is an alias to an internal value.
const FinancialEventTypeCollection = shared.FinancialEventTypeCollection

// This is an alias to an internal value.
const FinancialEventTypeCorrectionCredit = shared.FinancialEventTypeCorrectionCredit

// This is an alias to an internal value.
const FinancialEventTypeCorrectionDebit = shared.FinancialEventTypeCorrectionDebit

// This is an alias to an internal value.
const FinancialEventTypeCreditAuthorization = shared.FinancialEventTypeCreditAuthorization

// This is an alias to an internal value.
const FinancialEventTypeCreditAuthorizationAdvice = shared.FinancialEventTypeCreditAuthorizationAdvice

// This is an alias to an internal value.
const FinancialEventTypeCurrencyConversion = shared.FinancialEventTypeCurrencyConversion

// This is an alias to an internal value.
const FinancialEventTypeCurrencyConversionReversal = shared.FinancialEventTypeCurrencyConversionReversal

// This is an alias to an internal value.
const FinancialEventTypeDisputeWon = shared.FinancialEventTypeDisputeWon

// This is an alias to an internal value.
const FinancialEventTypeExternalACHCanceled = shared.FinancialEventTypeExternalACHCanceled

// This is an alias to an internal value.
const FinancialEventTypeExternalACHInitiated = shared.FinancialEventTypeExternalACHInitiated

// This is an alias to an internal value.
const FinancialEventTypeExternalACHReleased = shared.FinancialEventTypeExternalACHReleased

// This is an alias to an internal value.
const FinancialEventTypeExternalACHReversed = shared.FinancialEventTypeExternalACHReversed

// This is an alias to an internal value.
const FinancialEventTypeExternalACHSettled = shared.FinancialEventTypeExternalACHSettled

// This is an alias to an internal value.
const FinancialEventTypeExternalCheckCanceled = shared.FinancialEventTypeExternalCheckCanceled

// This is an alias to an internal value.
const FinancialEventTypeExternalCheckInitiated = shared.FinancialEventTypeExternalCheckInitiated

// This is an alias to an internal value.
const FinancialEventTypeExternalCheckReleased = shared.FinancialEventTypeExternalCheckReleased

// This is an alias to an internal value.
const FinancialEventTypeExternalCheckReversed = shared.FinancialEventTypeExternalCheckReversed

// This is an alias to an internal value.
const FinancialEventTypeExternalCheckSettled = shared.FinancialEventTypeExternalCheckSettled

// This is an alias to an internal value.
const FinancialEventTypeExternalTransferCanceled = shared.FinancialEventTypeExternalTransferCanceled

// This is an alias to an internal value.
const FinancialEventTypeExternalTransferInitiated = shared.FinancialEventTypeExternalTransferInitiated

// This is an alias to an internal value.
const FinancialEventTypeExternalTransferReleased = shared.FinancialEventTypeExternalTransferReleased

// This is an alias to an internal value.
const FinancialEventTypeExternalTransferReversed = shared.FinancialEventTypeExternalTransferReversed

// This is an alias to an internal value.
const FinancialEventTypeExternalTransferSettled = shared.FinancialEventTypeExternalTransferSettled

// This is an alias to an internal value.
const FinancialEventTypeExternalWireCanceled = shared.FinancialEventTypeExternalWireCanceled

// This is an alias to an internal value.
const FinancialEventTypeExternalWireInitiated = shared.FinancialEventTypeExternalWireInitiated

// This is an alias to an internal value.
const FinancialEventTypeExternalWireReleased = shared.FinancialEventTypeExternalWireReleased

// This is an alias to an internal value.
const FinancialEventTypeExternalWireReversed = shared.FinancialEventTypeExternalWireReversed

// This is an alias to an internal value.
const FinancialEventTypeExternalWireSettled = shared.FinancialEventTypeExternalWireSettled

// This is an alias to an internal value.
const FinancialEventTypeFinancialAuthorization = shared.FinancialEventTypeFinancialAuthorization

// This is an alias to an internal value.
const FinancialEventTypeFinancialCreditAuthorization = shared.FinancialEventTypeFinancialCreditAuthorization

// This is an alias to an internal value.
const FinancialEventTypeInterest = shared.FinancialEventTypeInterest

// This is an alias to an internal value.
const FinancialEventTypeInterestReversal = shared.FinancialEventTypeInterestReversal

// This is an alias to an internal value.
const FinancialEventTypeInternalAdjustment = shared.FinancialEventTypeInternalAdjustment

// This is an alias to an internal value.
const FinancialEventTypeLatePayment = shared.FinancialEventTypeLatePayment

// This is an alias to an internal value.
const FinancialEventTypeLatePaymentReversal = shared.FinancialEventTypeLatePaymentReversal

// This is an alias to an internal value.
const FinancialEventTypeLossWriteOff = shared.FinancialEventTypeLossWriteOff

// This is an alias to an internal value.
const FinancialEventTypeProvisionalCredit = shared.FinancialEventTypeProvisionalCredit

// This is an alias to an internal value.
const FinancialEventTypeProvisionalCreditReversal = shared.FinancialEventTypeProvisionalCreditReversal

// This is an alias to an internal value.
const FinancialEventTypeService = shared.FinancialEventTypeService

// This is an alias to an internal value.
const FinancialEventTypeReturn = shared.FinancialEventTypeReturn

// This is an alias to an internal value.
const FinancialEventTypeReturnReversal = shared.FinancialEventTypeReturnReversal

// This is an alias to an internal value.
const FinancialEventTypeTransfer = shared.FinancialEventTypeTransfer

// This is an alias to an internal value.
const FinancialEventTypeTransferInsufficientFunds = shared.FinancialEventTypeTransferInsufficientFunds

// This is an alias to an internal value.
const FinancialEventTypeReturnedPayment = shared.FinancialEventTypeReturnedPayment

// This is an alias to an internal value.
const FinancialEventTypeReturnedPaymentReversal = shared.FinancialEventTypeReturnedPaymentReversal

// This is an alias to an internal value.
const FinancialEventTypeLithicNetworkPayment = shared.FinancialEventTypeLithicNetworkPayment

// This is an alias to an internal type.
type Merchant = shared.Merchant

// This is an alias to an internal type.
type ShippingAddressParam = shared.ShippingAddressParam
