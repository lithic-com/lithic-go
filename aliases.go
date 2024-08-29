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

// Describes the document and the required document image uploads required to
// re-run KYC
//
// This is an alias to an internal type.
type Document = shared.Document

// Type of documentation to be submitted for verification.
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

// Status of document image upload.
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
const DocumentRequiredDocumentUploadsStatusReasonUnknownError = shared.DocumentRequiredDocumentUploadsStatusReasonUnknownError

// This is an alias to an internal type.
type ShippingAddressParam = shared.ShippingAddressParam
