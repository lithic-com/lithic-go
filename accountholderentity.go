// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// AccountHolderEntityService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountHolderEntityService] method instead.
type AccountHolderEntityService struct {
	Options []option.RequestOption
}

// NewAccountHolderEntityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountHolderEntityService(opts ...option.RequestOption) (r *AccountHolderEntityService) {
	r = &AccountHolderEntityService{}
	r.Options = opts
	return
}

// Create a new beneficial owner or replace the control person entity on an
// existing KYB account holder. This endpoint is only applicable for account
// holders enrolled through a KYB workflow with the Persona KYB provider. A new
// control person can only replace the existing one. A maximum of 4 beneficial
// owners can be associated with an account holder.
func (r *AccountHolderEntityService) New(ctx context.Context, accountHolderToken string, body AccountHolderEntityNewParams, opts ...option.RequestOption) (res *AccountHolderEntityNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	path := fmt.Sprintf("v1/account_holders/%s/entities", accountHolderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deactivate a beneficial owner entity on an existing KYB account holder. Only
// beneficial owner entities can be deactivated.
func (r *AccountHolderEntityService) Delete(ctx context.Context, accountHolderToken string, entityToken string, opts ...option.RequestOption) (res *AccountHolderEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountHolderToken == "" {
		err = errors.New("missing required account_holder_token parameter")
		return
	}
	if entityToken == "" {
		err = errors.New("missing required entity_token parameter")
		return
	}
	path := fmt.Sprintf("v1/account_holders/%s/entities/%s", accountHolderToken, entityToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Information about an entity associated with an account holder
type AccountHolderEntity struct {
	// Globally unique identifier for the entity
	Token string `json:"token" api:"required" format:"uuid"`
	// Globally unique identifier for the account holder
	AccountHolderToken string `json:"account_holder_token" api:"required" format:"uuid"`
	// Individual's current address
	Address AccountHolderEntityAddress `json:"address" api:"required"`
	// Individual's date of birth, as an RFC 3339 date
	Dob string `json:"dob" api:"required,nullable"`
	// Individual's email address
	Email string `json:"email" api:"required,nullable"`
	// Individual's first name, as it appears on government-issued identity documents
	FirstName string `json:"first_name" api:"required,nullable"`
	// Individual's last name, as it appears on government-issued identity documents
	LastName string `json:"last_name" api:"required,nullable"`
	// Individual's phone number, entered in E.164 format
	PhoneNumber string `json:"phone_number" api:"required,nullable"`
	// The status of the entity
	Status AccountHolderEntityStatus `json:"status" api:"required"`
	// The type of entity
	Type AccountHolderEntityType `json:"type" api:"required"`
	JSON accountHolderEntityJSON `json:"-"`
}

// accountHolderEntityJSON contains the JSON metadata for the struct
// [AccountHolderEntity]
type accountHolderEntityJSON struct {
	Token              apijson.Field
	AccountHolderToken apijson.Field
	Address            apijson.Field
	Dob                apijson.Field
	Email              apijson.Field
	FirstName          apijson.Field
	LastName           apijson.Field
	PhoneNumber        apijson.Field
	Status             apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountHolderEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderEntityJSON) RawJSON() string {
	return r.raw
}

// Individual's current address
type AccountHolderEntityAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1" api:"required"`
	// Name of city.
	City string `json:"city" api:"required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country" api:"required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code" api:"required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state" api:"required"`
	// Unit or apartment number (if applicable).
	Address2 string                         `json:"address2"`
	JSON     accountHolderEntityAddressJSON `json:"-"`
}

// accountHolderEntityAddressJSON contains the JSON metadata for the struct
// [AccountHolderEntityAddress]
type accountHolderEntityAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderEntityAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderEntityAddressJSON) RawJSON() string {
	return r.raw
}

// The status of the entity
type AccountHolderEntityStatus string

const (
	AccountHolderEntityStatusAccepted      AccountHolderEntityStatus = "ACCEPTED"
	AccountHolderEntityStatusInactive      AccountHolderEntityStatus = "INACTIVE"
	AccountHolderEntityStatusPendingReview AccountHolderEntityStatus = "PENDING_REVIEW"
	AccountHolderEntityStatusRejected      AccountHolderEntityStatus = "REJECTED"
)

func (r AccountHolderEntityStatus) IsKnown() bool {
	switch r {
	case AccountHolderEntityStatusAccepted, AccountHolderEntityStatusInactive, AccountHolderEntityStatusPendingReview, AccountHolderEntityStatusRejected:
		return true
	}
	return false
}

// The type of entity
type AccountHolderEntityType string

const (
	AccountHolderEntityTypeBeneficialOwnerIndividual AccountHolderEntityType = "BENEFICIAL_OWNER_INDIVIDUAL"
	AccountHolderEntityTypeControlPerson             AccountHolderEntityType = "CONTROL_PERSON"
)

func (r AccountHolderEntityType) IsKnown() bool {
	switch r {
	case AccountHolderEntityTypeBeneficialOwnerIndividual, AccountHolderEntityTypeControlPerson:
		return true
	}
	return false
}

// Response body for creating a new beneficial owner or replacing the control
// person entity on an existing KYB account holder.
type AccountHolderEntityNewResponse struct {
	// Globally unique identifier for the entity
	Token string `json:"token" api:"required" format:"uuid"`
	// Globally unique identifier for the account holder
	AccountHolderToken string `json:"account_holder_token" api:"required" format:"uuid"`
	// Timestamp of when the entity was created
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// A list of documents required for the entity to be approved
	RequiredDocuments []RequiredDocument `json:"required_documents" api:"required"`
	// Entity verification status
	Status AccountHolderEntityNewResponseStatus `json:"status" api:"required"`
	// Reason for the evaluation status
	StatusReasons []AccountHolderEntityNewResponseStatusReason `json:"status_reasons" api:"required"`
	JSON          accountHolderEntityNewResponseJSON           `json:"-"`
}

// accountHolderEntityNewResponseJSON contains the JSON metadata for the struct
// [AccountHolderEntityNewResponse]
type accountHolderEntityNewResponseJSON struct {
	Token              apijson.Field
	AccountHolderToken apijson.Field
	Created            apijson.Field
	RequiredDocuments  apijson.Field
	Status             apijson.Field
	StatusReasons      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountHolderEntityNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderEntityNewResponseJSON) RawJSON() string {
	return r.raw
}

// Entity verification status
type AccountHolderEntityNewResponseStatus string

const (
	AccountHolderEntityNewResponseStatusAccepted      AccountHolderEntityNewResponseStatus = "ACCEPTED"
	AccountHolderEntityNewResponseStatusInactive      AccountHolderEntityNewResponseStatus = "INACTIVE"
	AccountHolderEntityNewResponseStatusPendingReview AccountHolderEntityNewResponseStatus = "PENDING_REVIEW"
	AccountHolderEntityNewResponseStatusRejected      AccountHolderEntityNewResponseStatus = "REJECTED"
)

func (r AccountHolderEntityNewResponseStatus) IsKnown() bool {
	switch r {
	case AccountHolderEntityNewResponseStatusAccepted, AccountHolderEntityNewResponseStatusInactive, AccountHolderEntityNewResponseStatusPendingReview, AccountHolderEntityNewResponseStatusRejected:
		return true
	}
	return false
}

// Status Reasons for KYC/KYB enrollment states
type AccountHolderEntityNewResponseStatusReason string

const (
	AccountHolderEntityNewResponseStatusReasonAddressVerificationFailure                      AccountHolderEntityNewResponseStatusReason = "ADDRESS_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonAgeThresholdFailure                             AccountHolderEntityNewResponseStatusReason = "AGE_THRESHOLD_FAILURE"
	AccountHolderEntityNewResponseStatusReasonCompleteVerificationFailure                     AccountHolderEntityNewResponseStatusReason = "COMPLETE_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonDobVerificationFailure                          AccountHolderEntityNewResponseStatusReason = "DOB_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonIDVerificationFailure                           AccountHolderEntityNewResponseStatusReason = "ID_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonMaxDocumentAttempts                             AccountHolderEntityNewResponseStatusReason = "MAX_DOCUMENT_ATTEMPTS"
	AccountHolderEntityNewResponseStatusReasonMaxResubmissionAttempts                         AccountHolderEntityNewResponseStatusReason = "MAX_RESUBMISSION_ATTEMPTS"
	AccountHolderEntityNewResponseStatusReasonNameVerificationFailure                         AccountHolderEntityNewResponseStatusReason = "NAME_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonOtherVerificationFailure                        AccountHolderEntityNewResponseStatusReason = "OTHER_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonRiskThresholdFailure                            AccountHolderEntityNewResponseStatusReason = "RISK_THRESHOLD_FAILURE"
	AccountHolderEntityNewResponseStatusReasonWatchlistAlertFailure                           AccountHolderEntityNewResponseStatusReason = "WATCHLIST_ALERT_FAILURE"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityIDVerificationFailure      AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_ID_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityAddressVerificationFailure AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_ADDRESS_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityNameVerificationFailure    AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_NAME_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_BUSINESS_OFFICERS_NOT_MATCHED"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntitySosFilingInactive          AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_FILING_INACTIVE"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntitySosNotMatched              AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_SOS_NOT_MATCHED"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityCmraFailure                AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_CMRA_FAILURE"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityWatchlistFailure           AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_WATCHLIST_FAILURE"
	AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityRegisteredAgentFailure     AccountHolderEntityNewResponseStatusReason = "PRIMARY_BUSINESS_ENTITY_REGISTERED_AGENT_FAILURE"
	AccountHolderEntityNewResponseStatusReasonControlPersonBlocklistAlertFailure              AccountHolderEntityNewResponseStatusReason = "CONTROL_PERSON_BLOCKLIST_ALERT_FAILURE"
	AccountHolderEntityNewResponseStatusReasonControlPersonIDVerificationFailure              AccountHolderEntityNewResponseStatusReason = "CONTROL_PERSON_ID_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonControlPersonDobVerificationFailure             AccountHolderEntityNewResponseStatusReason = "CONTROL_PERSON_DOB_VERIFICATION_FAILURE"
	AccountHolderEntityNewResponseStatusReasonControlPersonNameVerificationFailure            AccountHolderEntityNewResponseStatusReason = "CONTROL_PERSON_NAME_VERIFICATION_FAILURE"
)

func (r AccountHolderEntityNewResponseStatusReason) IsKnown() bool {
	switch r {
	case AccountHolderEntityNewResponseStatusReasonAddressVerificationFailure, AccountHolderEntityNewResponseStatusReasonAgeThresholdFailure, AccountHolderEntityNewResponseStatusReasonCompleteVerificationFailure, AccountHolderEntityNewResponseStatusReasonDobVerificationFailure, AccountHolderEntityNewResponseStatusReasonIDVerificationFailure, AccountHolderEntityNewResponseStatusReasonMaxDocumentAttempts, AccountHolderEntityNewResponseStatusReasonMaxResubmissionAttempts, AccountHolderEntityNewResponseStatusReasonNameVerificationFailure, AccountHolderEntityNewResponseStatusReasonOtherVerificationFailure, AccountHolderEntityNewResponseStatusReasonRiskThresholdFailure, AccountHolderEntityNewResponseStatusReasonWatchlistAlertFailure, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityIDVerificationFailure, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityAddressVerificationFailure, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityNameVerificationFailure, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityBusinessOfficersNotMatched, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntitySosFilingInactive, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntitySosNotMatched, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityCmraFailure, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityWatchlistFailure, AccountHolderEntityNewResponseStatusReasonPrimaryBusinessEntityRegisteredAgentFailure, AccountHolderEntityNewResponseStatusReasonControlPersonBlocklistAlertFailure, AccountHolderEntityNewResponseStatusReasonControlPersonIDVerificationFailure, AccountHolderEntityNewResponseStatusReasonControlPersonDobVerificationFailure, AccountHolderEntityNewResponseStatusReasonControlPersonNameVerificationFailure:
		return true
	}
	return false
}

type AccountHolderEntityNewParams struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address param.Field[AccountHolderEntityNewParamsAddress] `json:"address" api:"required"`
	// Individual's date of birth, as an RFC 3339 date.
	Dob param.Field[string] `json:"dob" api:"required"`
	// Individual's email address. If utilizing Lithic for chargeback processing, this
	// customer email address may be used to communicate dispute status and resolution.
	Email param.Field[string] `json:"email" api:"required"`
	// Individual's first name, as it appears on government-issued identity documents.
	FirstName param.Field[string] `json:"first_name" api:"required"`
	// Government-issued identification number (required for identity verification and
	// compliance with banking regulations). Social Security Numbers (SSN) and
	// Individual Taxpayer Identification Numbers (ITIN) are currently supported,
	// entered as full nine-digits, with or without hyphens
	GovernmentID param.Field[string] `json:"government_id" api:"required"`
	// Individual's last name, as it appears on government-issued identity documents.
	LastName param.Field[string] `json:"last_name" api:"required"`
	// Individual's phone number, entered in E.164 format.
	PhoneNumber param.Field[string] `json:"phone_number" api:"required"`
	// The type of entity to create on the account holder
	Type param.Field[AccountHolderEntityNewParamsType] `json:"type" api:"required"`
}

func (r AccountHolderEntityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderEntityNewParamsAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 param.Field[string] `json:"address1" api:"required"`
	// Name of city.
	City param.Field[string] `json:"city" api:"required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country param.Field[string] `json:"country" api:"required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State param.Field[string] `json:"state" api:"required"`
	// Unit or apartment number (if applicable).
	Address2 param.Field[string] `json:"address2"`
}

func (r AccountHolderEntityNewParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of entity to create on the account holder
type AccountHolderEntityNewParamsType string

const (
	AccountHolderEntityNewParamsTypeBeneficialOwnerIndividual AccountHolderEntityNewParamsType = "BENEFICIAL_OWNER_INDIVIDUAL"
	AccountHolderEntityNewParamsTypeControlPerson             AccountHolderEntityNewParamsType = "CONTROL_PERSON"
)

func (r AccountHolderEntityNewParamsType) IsKnown() bool {
	switch r {
	case AccountHolderEntityNewParamsTypeBeneficialOwnerIndividual, AccountHolderEntityNewParamsTypeControlPerson:
		return true
	}
	return false
}
