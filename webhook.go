// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"errors"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/shared"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"github.com/tidwall/gjson"
)

// WebhookService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) Parsed(payload []byte, headers http.Header, opts ...option.RequestOption) (*ParsedWebhookEvent, error) {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.WebhookSecret
	if key == "" {
		return nil, errors.New("The WebhookSecret option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &ParsedWebhookEvent{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type AccountHolderCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType AccountHolderCreatedWebhookEventEventType `json:"event_type,required"`
	// The token of the account_holder that was created.
	Token string `json:"token" format:"uuid"`
	// The token of the account that was created.
	AccountToken string `json:"account_token" format:"uuid"`
	// When the account_holder was created
	Created           time.Time          `json:"created" format:"date-time"`
	RequiredDocuments []RequiredDocument `json:"required_documents"`
	// The status of the account_holder that was created.
	Status       AccountHolderCreatedWebhookEventStatus `json:"status"`
	StatusReason []string                               `json:"status_reason"`
	JSON         accountHolderCreatedWebhookEventJSON   `json:"-"`
}

// accountHolderCreatedWebhookEventJSON contains the JSON metadata for the struct
// [AccountHolderCreatedWebhookEvent]
type accountHolderCreatedWebhookEventJSON struct {
	EventType         apijson.Field
	Token             apijson.Field
	AccountToken      apijson.Field
	Created           apijson.Field
	RequiredDocuments apijson.Field
	Status            apijson.Field
	StatusReason      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountHolderCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r AccountHolderCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type AccountHolderCreatedWebhookEventEventType string

const (
	AccountHolderCreatedWebhookEventEventTypeAccountHolderCreated AccountHolderCreatedWebhookEventEventType = "account_holder.created"
)

func (r AccountHolderCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case AccountHolderCreatedWebhookEventEventTypeAccountHolderCreated:
		return true
	}
	return false
}

// The status of the account_holder that was created.
type AccountHolderCreatedWebhookEventStatus string

const (
	AccountHolderCreatedWebhookEventStatusAccepted      AccountHolderCreatedWebhookEventStatus = "ACCEPTED"
	AccountHolderCreatedWebhookEventStatusPendingReview AccountHolderCreatedWebhookEventStatus = "PENDING_REVIEW"
)

func (r AccountHolderCreatedWebhookEventStatus) IsKnown() bool {
	switch r {
	case AccountHolderCreatedWebhookEventStatusAccepted, AccountHolderCreatedWebhookEventStatusPendingReview:
		return true
	}
	return false
}

// KYB payload for an updated account holder.
type AccountHolderUpdatedWebhookEvent struct {
	// The token of the account_holder that was created.
	Token string `json:"token,required" format:"uuid"`
	// If applicable, represents the business account token associated with the
	// account_holder.
	BusinessAccountToken string `json:"business_account_token,nullable" format:"uuid"`
	// When the account_holder updated event was created
	Created time.Time `json:"created" format:"date-time"`
	// If updated, the newly updated email associated with the account_holder otherwise
	// the existing email is provided.
	Email string `json:"email"`
	// The type of event that occurred.
	EventType AccountHolderUpdatedWebhookEventEventType `json:"event_type"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID string `json:"external_id,nullable"`
	// If applicable, represents the account_holder's first name.
	FirstName string `json:"first_name"`
	// If applicable, represents the account_holder's last name.
	LastName string `json:"last_name"`
	// If applicable, represents the account_holder's business name.
	LegalBusinessName string `json:"legal_business_name"`
	// 6-digit North American Industry Classification System (NAICS) code for the
	// business. Only present if naics_code was included in the update request.
	NaicsCode string `json:"naics_code"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness string `json:"nature_of_business"`
	// If updated, the newly updated phone_number associated with the account_holder
	// otherwise the existing phone_number is provided.
	PhoneNumber string `json:"phone_number"`
	// This field can have the runtime type of
	// [AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequest],
	// [AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequest].
	UpdateRequest interface{} `json:"update_request"`
	// Company website URL.
	WebsiteURL string                               `json:"website_url"`
	JSON       accountHolderUpdatedWebhookEventJSON `json:"-"`
	union      AccountHolderUpdatedWebhookEventUnion
}

// accountHolderUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [AccountHolderUpdatedWebhookEvent]
type accountHolderUpdatedWebhookEventJSON struct {
	Token                apijson.Field
	BusinessAccountToken apijson.Field
	Created              apijson.Field
	Email                apijson.Field
	EventType            apijson.Field
	ExternalID           apijson.Field
	FirstName            apijson.Field
	LastName             apijson.Field
	LegalBusinessName    apijson.Field
	NaicsCode            apijson.Field
	NatureOfBusiness     apijson.Field
	PhoneNumber          apijson.Field
	UpdateRequest        apijson.Field
	WebsiteURL           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r accountHolderUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *AccountHolderUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = AccountHolderUpdatedWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountHolderUpdatedWebhookEventUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountHolderUpdatedWebhookEventKYBPayload],
// [AccountHolderUpdatedWebhookEventKYCPayload],
// [AccountHolderUpdatedWebhookEventLegacyPayload].
func (r AccountHolderUpdatedWebhookEvent) AsUnion() AccountHolderUpdatedWebhookEventUnion {
	return r.union
}

// KYB payload for an updated account holder.
//
// Union satisfied by [AccountHolderUpdatedWebhookEventKYBPayload],
// [AccountHolderUpdatedWebhookEventKYCPayload] or
// [AccountHolderUpdatedWebhookEventLegacyPayload].
type AccountHolderUpdatedWebhookEventUnion interface {
	implementsAccountHolderUpdatedWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountHolderUpdatedWebhookEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountHolderUpdatedWebhookEventKYBPayload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountHolderUpdatedWebhookEventKYCPayload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountHolderUpdatedWebhookEventLegacyPayload{}),
		},
	)
}

// KYB payload for an updated account holder.
type AccountHolderUpdatedWebhookEventKYBPayload struct {
	// The token of the account_holder that was created.
	Token string `json:"token,required" format:"uuid"`
	// Original request to update the account holder.
	UpdateRequest AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequest `json:"update_request,required"`
	// The type of event that occurred.
	EventType AccountHolderUpdatedWebhookEventKYBPayloadEventType `json:"event_type"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID string `json:"external_id"`
	// 6-digit North American Industry Classification System (NAICS) code for the
	// business. Only present if naics_code was included in the update request.
	NaicsCode string `json:"naics_code"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness string `json:"nature_of_business"`
	// Company website URL.
	WebsiteURL string                                         `json:"website_url"`
	JSON       accountHolderUpdatedWebhookEventKYBPayloadJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventKYBPayloadJSON contains the JSON metadata for
// the struct [AccountHolderUpdatedWebhookEventKYBPayload]
type accountHolderUpdatedWebhookEventKYBPayloadJSON struct {
	Token            apijson.Field
	UpdateRequest    apijson.Field
	EventType        apijson.Field
	ExternalID       apijson.Field
	NaicsCode        apijson.Field
	NatureOfBusiness apijson.Field
	WebsiteURL       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYBPayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYBPayloadJSON) RawJSON() string {
	return r.raw
}

func (r AccountHolderUpdatedWebhookEventKYBPayload) implementsAccountHolderUpdatedWebhookEvent() {}

// Original request to update the account holder.
type AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequest struct {
	// Deprecated.
	//
	// Deprecated: deprecated
	BeneficialOwnerEntities []KYBBusinessEntity `json:"beneficial_owner_entities"`
	// You must submit a list of all direct and indirect individuals with 25% or more
	// ownership in the company. A maximum of 4 beneficial owners can be submitted. If
	// no individual owns 25% of the company you do not need to send beneficial owner
	// information. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included.
	BeneficialOwnerIndividuals []AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// Information for business for which the account is being opened and KYB is being
	// run.
	BusinessEntity KYBBusinessEntity `json:"business_entity"`
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer, Managing Member, General Partner, President, Vice President, or
	// Treasurer). This can be an executive, or someone who will have program-wide
	// access to the cards that Lithic will provide. In some cases, this individual
	// could also be a beneficial owner listed above. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section II) for more background.
	ControlPerson AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPerson `json:"control_person"`
	JSON          accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestJSON          `json:"-"`
}

// accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestJSON contains the JSON
// metadata for the struct
// [AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequest]
type accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestJSON struct {
	BeneficialOwnerEntities    apijson.Field
	BeneficialOwnerIndividuals apijson.Field
	BusinessEntity             apijson.Field
	ControlPerson              apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestJSON) RawJSON() string {
	return r.raw
}

type AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddress `json:"address"`
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
	PhoneNumber string                                                                               `json:"phone_number"`
	JSON        accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualJSON
// contains the JSON metadata for the struct
// [AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividual]
type accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddress struct {
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
	Address2 string                                                                                       `json:"address2"`
	JSON     accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddressJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddressJSON
// contains the JSON metadata for the struct
// [AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddress]
type accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddressJSON) RawJSON() string {
	return r.raw
}

// An individual with significant responsibility for managing the legal entity
// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
// Officer, Managing Member, General Partner, President, Vice President, or
// Treasurer). This can be an executive, or someone who will have program-wide
// access to the cards that Lithic will provide. In some cases, this individual
// could also be a beneficial owner listed above. See
// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
// (Section II) for more background.
type AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPerson struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonAddress `json:"address"`
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
	PhoneNumber string                                                                   `json:"phone_number"`
	JSON        accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonJSON
// contains the JSON metadata for the struct
// [AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPerson]
type accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonAddress struct {
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
	JSON     accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonAddressJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonAddressJSON
// contains the JSON metadata for the struct
// [AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonAddress]
type accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYBPayloadUpdateRequestControlPersonAddressJSON) RawJSON() string {
	return r.raw
}

// The type of event that occurred.
type AccountHolderUpdatedWebhookEventKYBPayloadEventType string

const (
	AccountHolderUpdatedWebhookEventKYBPayloadEventTypeAccountHolderUpdated AccountHolderUpdatedWebhookEventKYBPayloadEventType = "account_holder.updated"
)

func (r AccountHolderUpdatedWebhookEventKYBPayloadEventType) IsKnown() bool {
	switch r {
	case AccountHolderUpdatedWebhookEventKYBPayloadEventTypeAccountHolderUpdated:
		return true
	}
	return false
}

// KYC payload for an updated account holder.
type AccountHolderUpdatedWebhookEventKYCPayload struct {
	// The token of the account_holder that was created.
	Token string `json:"token,required" format:"uuid"`
	// Original request to update the account holder.
	UpdateRequest AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequest `json:"update_request,required"`
	// The type of event that occurred.
	EventType AccountHolderUpdatedWebhookEventKYCPayloadEventType `json:"event_type"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID string                                         `json:"external_id"`
	JSON       accountHolderUpdatedWebhookEventKYCPayloadJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventKYCPayloadJSON contains the JSON metadata for
// the struct [AccountHolderUpdatedWebhookEventKYCPayload]
type accountHolderUpdatedWebhookEventKYCPayloadJSON struct {
	Token         apijson.Field
	UpdateRequest apijson.Field
	EventType     apijson.Field
	ExternalID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYCPayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYCPayloadJSON) RawJSON() string {
	return r.raw
}

func (r AccountHolderUpdatedWebhookEventKYCPayload) implementsAccountHolderUpdatedWebhookEvent() {}

// Original request to update the account holder.
type AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequest struct {
	// Information on the individual for whom the account is being opened and KYC is
	// being run.
	Individual AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividual `json:"individual"`
	JSON       accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestJSON       `json:"-"`
}

// accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestJSON contains the JSON
// metadata for the struct
// [AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequest]
type accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestJSON struct {
	Individual  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestJSON) RawJSON() string {
	return r.raw
}

// Information on the individual for whom the account is being opened and KYC is
// being run.
type AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualAddress `json:"address"`
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
	PhoneNumber string                                                                `json:"phone_number"`
	JSON        accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualJSON contains
// the JSON metadata for the struct
// [AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividual]
type accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualAddress struct {
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
	Address2 string                                                                       `json:"address2"`
	JSON     accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualAddressJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualAddressJSON
// contains the JSON metadata for the struct
// [AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualAddress]
type accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventKYCPayloadUpdateRequestIndividualAddressJSON) RawJSON() string {
	return r.raw
}

// The type of event that occurred.
type AccountHolderUpdatedWebhookEventKYCPayloadEventType string

const (
	AccountHolderUpdatedWebhookEventKYCPayloadEventTypeAccountHolderUpdated AccountHolderUpdatedWebhookEventKYCPayloadEventType = "account_holder.updated"
)

func (r AccountHolderUpdatedWebhookEventKYCPayloadEventType) IsKnown() bool {
	switch r {
	case AccountHolderUpdatedWebhookEventKYCPayloadEventTypeAccountHolderUpdated:
		return true
	}
	return false
}

// Legacy payload for an updated account holder.
type AccountHolderUpdatedWebhookEventLegacyPayload struct {
	// The token of the account_holder that was created.
	Token string `json:"token,required" format:"uuid"`
	// If applicable, represents the business account token associated with the
	// account_holder.
	BusinessAccountToken string `json:"business_account_token,nullable" format:"uuid"`
	// When the account_holder updated event was created
	Created time.Time `json:"created" format:"date-time"`
	// If updated, the newly updated email associated with the account_holder otherwise
	// the existing email is provided.
	Email string `json:"email"`
	// The type of event that occurred.
	EventType AccountHolderUpdatedWebhookEventLegacyPayloadEventType `json:"event_type"`
	// If applicable, represents the external_id associated with the account_holder.
	ExternalID string `json:"external_id,nullable"`
	// If applicable, represents the account_holder's first name.
	FirstName string `json:"first_name"`
	// If applicable, represents the account_holder's last name.
	LastName string `json:"last_name"`
	// If applicable, represents the account_holder's business name.
	LegalBusinessName string `json:"legal_business_name"`
	// If updated, the newly updated phone_number associated with the account_holder
	// otherwise the existing phone_number is provided.
	PhoneNumber string                                            `json:"phone_number"`
	JSON        accountHolderUpdatedWebhookEventLegacyPayloadJSON `json:"-"`
}

// accountHolderUpdatedWebhookEventLegacyPayloadJSON contains the JSON metadata for
// the struct [AccountHolderUpdatedWebhookEventLegacyPayload]
type accountHolderUpdatedWebhookEventLegacyPayloadJSON struct {
	Token                apijson.Field
	BusinessAccountToken apijson.Field
	Created              apijson.Field
	Email                apijson.Field
	EventType            apijson.Field
	ExternalID           apijson.Field
	FirstName            apijson.Field
	LastName             apijson.Field
	LegalBusinessName    apijson.Field
	PhoneNumber          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountHolderUpdatedWebhookEventLegacyPayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderUpdatedWebhookEventLegacyPayloadJSON) RawJSON() string {
	return r.raw
}

func (r AccountHolderUpdatedWebhookEventLegacyPayload) implementsAccountHolderUpdatedWebhookEvent() {}

// The type of event that occurred.
type AccountHolderUpdatedWebhookEventLegacyPayloadEventType string

const (
	AccountHolderUpdatedWebhookEventLegacyPayloadEventTypeAccountHolderUpdated AccountHolderUpdatedWebhookEventLegacyPayloadEventType = "account_holder.updated"
)

func (r AccountHolderUpdatedWebhookEventLegacyPayloadEventType) IsKnown() bool {
	switch r {
	case AccountHolderUpdatedWebhookEventLegacyPayloadEventTypeAccountHolderUpdated:
		return true
	}
	return false
}

// The type of event that occurred.
type AccountHolderUpdatedWebhookEventEventType string

const (
	AccountHolderUpdatedWebhookEventEventTypeAccountHolderUpdated AccountHolderUpdatedWebhookEventEventType = "account_holder.updated"
)

func (r AccountHolderUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case AccountHolderUpdatedWebhookEventEventTypeAccountHolderUpdated:
		return true
	}
	return false
}

type AccountHolderVerificationWebhookEvent struct {
	// The type of event that occurred.
	EventType AccountHolderVerificationWebhookEventEventType `json:"event_type,required"`
	// The token of the account_holder being verified.
	Token string `json:"token" format:"uuid"`
	// The token of the account being verified.
	AccountToken string `json:"account_token" format:"uuid"`
	// When the account_holder verification status was updated
	Created time.Time `json:"created" format:"date-time"`
	// The status of the account_holder that was created
	Status        AccountHolderVerificationWebhookEventStatus `json:"status"`
	StatusReasons []string                                    `json:"status_reasons"`
	JSON          accountHolderVerificationWebhookEventJSON   `json:"-"`
}

// accountHolderVerificationWebhookEventJSON contains the JSON metadata for the
// struct [AccountHolderVerificationWebhookEvent]
type accountHolderVerificationWebhookEventJSON struct {
	EventType     apijson.Field
	Token         apijson.Field
	AccountToken  apijson.Field
	Created       apijson.Field
	Status        apijson.Field
	StatusReasons apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountHolderVerificationWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderVerificationWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r AccountHolderVerificationWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type AccountHolderVerificationWebhookEventEventType string

const (
	AccountHolderVerificationWebhookEventEventTypeAccountHolderVerification AccountHolderVerificationWebhookEventEventType = "account_holder.verification"
)

func (r AccountHolderVerificationWebhookEventEventType) IsKnown() bool {
	switch r {
	case AccountHolderVerificationWebhookEventEventTypeAccountHolderVerification:
		return true
	}
	return false
}

// The status of the account_holder that was created
type AccountHolderVerificationWebhookEventStatus string

const (
	AccountHolderVerificationWebhookEventStatusAccepted      AccountHolderVerificationWebhookEventStatus = "ACCEPTED"
	AccountHolderVerificationWebhookEventStatusPendingReview AccountHolderVerificationWebhookEventStatus = "PENDING_REVIEW"
	AccountHolderVerificationWebhookEventStatusRejected      AccountHolderVerificationWebhookEventStatus = "REJECTED"
)

func (r AccountHolderVerificationWebhookEventStatus) IsKnown() bool {
	switch r {
	case AccountHolderVerificationWebhookEventStatusAccepted, AccountHolderVerificationWebhookEventStatusPendingReview, AccountHolderVerificationWebhookEventStatusRejected:
		return true
	}
	return false
}

type AccountHolderDocumentUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType AccountHolderDocumentUpdatedWebhookEventEventType `json:"event_type,required"`
	// The token of the account holder document
	Token string `json:"token" format:"uuid"`
	// The token of the account_holder that the document belongs to
	AccountHolderToken string `json:"account_holder_token" format:"uuid"`
	// When the account_holder was created
	Created time.Time `json:"created" format:"date-time"`
	// Type of documentation to be submitted for verification of an account holder
	DocumentType AccountHolderDocumentUpdatedWebhookEventDocumentType `json:"document_type"`
	// The token of the entity that the document belongs to
	EntityToken             string                                                           `json:"entity_token" format:"uuid"`
	RequiredDocumentUploads []AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUpload `json:"required_document_uploads"`
	JSON                    accountHolderDocumentUpdatedWebhookEventJSON                     `json:"-"`
}

// accountHolderDocumentUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [AccountHolderDocumentUpdatedWebhookEvent]
type accountHolderDocumentUpdatedWebhookEventJSON struct {
	EventType               apijson.Field
	Token                   apijson.Field
	AccountHolderToken      apijson.Field
	Created                 apijson.Field
	DocumentType            apijson.Field
	EntityToken             apijson.Field
	RequiredDocumentUploads apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountHolderDocumentUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderDocumentUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r AccountHolderDocumentUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type AccountHolderDocumentUpdatedWebhookEventEventType string

const (
	AccountHolderDocumentUpdatedWebhookEventEventTypeAccountHolderDocumentUpdated AccountHolderDocumentUpdatedWebhookEventEventType = "account_holder_document.updated"
)

func (r AccountHolderDocumentUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case AccountHolderDocumentUpdatedWebhookEventEventTypeAccountHolderDocumentUpdated:
		return true
	}
	return false
}

// Type of documentation to be submitted for verification of an account holder
type AccountHolderDocumentUpdatedWebhookEventDocumentType string

const (
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeDriversLicense            AccountHolderDocumentUpdatedWebhookEventDocumentType = "DRIVERS_LICENSE"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypePassport                  AccountHolderDocumentUpdatedWebhookEventDocumentType = "PASSPORT"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypePassportCard              AccountHolderDocumentUpdatedWebhookEventDocumentType = "PASSPORT_CARD"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeEinLetter                 AccountHolderDocumentUpdatedWebhookEventDocumentType = "EIN_LETTER"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeTaxReturn                 AccountHolderDocumentUpdatedWebhookEventDocumentType = "TAX_RETURN"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeOperatingAgreement        AccountHolderDocumentUpdatedWebhookEventDocumentType = "OPERATING_AGREEMENT"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeCertificateOfFormation    AccountHolderDocumentUpdatedWebhookEventDocumentType = "CERTIFICATE_OF_FORMATION"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeCertificateOfGoodStanding AccountHolderDocumentUpdatedWebhookEventDocumentType = "CERTIFICATE_OF_GOOD_STANDING"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeArticlesOfIncorporation   AccountHolderDocumentUpdatedWebhookEventDocumentType = "ARTICLES_OF_INCORPORATION"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeArticlesOfOrganization    AccountHolderDocumentUpdatedWebhookEventDocumentType = "ARTICLES_OF_ORGANIZATION"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeBylaws                    AccountHolderDocumentUpdatedWebhookEventDocumentType = "BYLAWS"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeGovernmentBusinessLicense AccountHolderDocumentUpdatedWebhookEventDocumentType = "GOVERNMENT_BUSINESS_LICENSE"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypePartnershipAgreement      AccountHolderDocumentUpdatedWebhookEventDocumentType = "PARTNERSHIP_AGREEMENT"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeSs4Form                   AccountHolderDocumentUpdatedWebhookEventDocumentType = "SS4_FORM"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeBankStatement             AccountHolderDocumentUpdatedWebhookEventDocumentType = "BANK_STATEMENT"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeUtilityBillStatement      AccountHolderDocumentUpdatedWebhookEventDocumentType = "UTILITY_BILL_STATEMENT"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeSsnCard                   AccountHolderDocumentUpdatedWebhookEventDocumentType = "SSN_CARD"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeItinLetter                AccountHolderDocumentUpdatedWebhookEventDocumentType = "ITIN_LETTER"
	AccountHolderDocumentUpdatedWebhookEventDocumentTypeFincenBoiReport           AccountHolderDocumentUpdatedWebhookEventDocumentType = "FINCEN_BOI_REPORT"
)

func (r AccountHolderDocumentUpdatedWebhookEventDocumentType) IsKnown() bool {
	switch r {
	case AccountHolderDocumentUpdatedWebhookEventDocumentTypeDriversLicense, AccountHolderDocumentUpdatedWebhookEventDocumentTypePassport, AccountHolderDocumentUpdatedWebhookEventDocumentTypePassportCard, AccountHolderDocumentUpdatedWebhookEventDocumentTypeEinLetter, AccountHolderDocumentUpdatedWebhookEventDocumentTypeTaxReturn, AccountHolderDocumentUpdatedWebhookEventDocumentTypeOperatingAgreement, AccountHolderDocumentUpdatedWebhookEventDocumentTypeCertificateOfFormation, AccountHolderDocumentUpdatedWebhookEventDocumentTypeCertificateOfGoodStanding, AccountHolderDocumentUpdatedWebhookEventDocumentTypeArticlesOfIncorporation, AccountHolderDocumentUpdatedWebhookEventDocumentTypeArticlesOfOrganization, AccountHolderDocumentUpdatedWebhookEventDocumentTypeBylaws, AccountHolderDocumentUpdatedWebhookEventDocumentTypeGovernmentBusinessLicense, AccountHolderDocumentUpdatedWebhookEventDocumentTypePartnershipAgreement, AccountHolderDocumentUpdatedWebhookEventDocumentTypeSs4Form, AccountHolderDocumentUpdatedWebhookEventDocumentTypeBankStatement, AccountHolderDocumentUpdatedWebhookEventDocumentTypeUtilityBillStatement, AccountHolderDocumentUpdatedWebhookEventDocumentTypeSsnCard, AccountHolderDocumentUpdatedWebhookEventDocumentTypeItinLetter, AccountHolderDocumentUpdatedWebhookEventDocumentTypeFincenBoiReport:
		return true
	}
	return false
}

// A document upload that belongs to the overall account holder document
type AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUpload struct {
	// The token of the document upload
	Token                       string   `json:"token" format:"uuid"`
	AcceptedEntityStatusReasons []string `json:"accepted_entity_status_reasons"`
	// When the document upload was created
	Created time.Time `json:"created" format:"date-time"`
	// The type of image that was uploaded
	ImageType                   AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageType `json:"image_type"`
	RejectedEntityStatusReasons []string                                                                 `json:"rejected_entity_status_reasons"`
	// The status of the document upload
	Status        AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatus `json:"status"`
	StatusReasons []string                                                              `json:"status_reasons"`
	// When the document upload was last updated
	Updated time.Time                                                          `json:"updated" format:"date-time"`
	JSON    accountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadJSON `json:"-"`
}

// accountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadJSON contains the
// JSON metadata for the struct
// [AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUpload]
type accountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadJSON struct {
	Token                       apijson.Field
	AcceptedEntityStatusReasons apijson.Field
	Created                     apijson.Field
	ImageType                   apijson.Field
	RejectedEntityStatusReasons apijson.Field
	Status                      apijson.Field
	StatusReasons               apijson.Field
	Updated                     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadJSON) RawJSON() string {
	return r.raw
}

// The type of image that was uploaded
type AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageType string

const (
	AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageTypeFront AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageType = "FRONT"
	AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageTypeBack  AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageType = "BACK"
)

func (r AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageType) IsKnown() bool {
	switch r {
	case AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageTypeFront, AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsImageTypeBack:
		return true
	}
	return false
}

// The status of the document upload
type AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatus string

const (
	AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusAccepted        AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatus = "ACCEPTED"
	AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusRejected        AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatus = "REJECTED"
	AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusPendingUpload   AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatus = "PENDING_UPLOAD"
	AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusUploaded        AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatus = "UPLOADED"
	AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusPartialApproval AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatus = "PARTIAL_APPROVAL"
)

func (r AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatus) IsKnown() bool {
	switch r {
	case AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusAccepted, AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusRejected, AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusPendingUpload, AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusUploaded, AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUploadsStatusPartialApproval:
		return true
	}
	return false
}

type CardAuthorizationApprovalRequestWebhookEvent struct {
	// The provisional transaction group uuid associated with the authorization
	Token string `json:"token,required" format:"uuid"`
	// Fee (in cents) assessed by the merchant and paid for by the cardholder. Will be
	// zero if no fee is assessed. Rebates may be transmitted as a negative value to
	// indicate credited fees.
	AcquirerFee int64 `json:"acquirer_fee,required"`
	// Deprecated, use `amounts`. Authorization amount of the transaction (in cents),
	// including any acquirer fees. The contents of this field are identical to
	// `authorization_amount`.
	//
	// Deprecated: deprecated
	Amount int64 `json:"amount,required"`
	// Structured amounts for this authorization. The `cardholder` and `merchant`
	// amounts reflect the original network authorization values. For programs with
	// hold adjustments enabled (e.g., automated fuel dispensers or tipping MCCs), the
	// `hold` amount may exceed the `cardholder` and `merchant` amounts to account for
	// anticipated final transaction amounts such as tips or fuel fill-ups
	Amounts CardAuthorizationApprovalRequestWebhookEventAmounts `json:"amounts,required"`
	// Deprecated, use `amounts`. The base transaction amount (in cents) plus the
	// acquirer fee field. This is the amount the issuer should authorize against
	// unless the issuer is paying the acquirer fee on behalf of the cardholder.
	//
	// Deprecated: deprecated
	AuthorizationAmount int64                                           `json:"authorization_amount,required"`
	Avs                 CardAuthorizationApprovalRequestWebhookEventAvs `json:"avs,required"`
	// Card object in ASA
	Card CardAuthorizationApprovalRequestWebhookEventCard `json:"card,required"`
	// Deprecated, use `amounts`. 3-character alphabetic ISO 4217 code for cardholder's
	// billing currency.
	//
	// Deprecated: deprecated
	CardholderCurrency string `json:"cardholder_currency,required"`
	// The portion of the transaction requested as cash back by the cardholder, and
	// does not include any acquirer fees. The amount field includes the purchase
	// amount, the requested cash back amount, and any acquirer fees.
	//
	// If no cash back was requested, the value of this field will be 0, and the field
	// will always be present.
	CashAmount int64 `json:"cash_amount,required"`
	// Date and time when the transaction first occurred in UTC.
	Created   time.Time                                             `json:"created,required" format:"date-time"`
	EventType CardAuthorizationApprovalRequestWebhookEventEventType `json:"event_type,required"`
	Merchant  shared.Merchant                                       `json:"merchant,required"`
	// Deprecated, use `amounts`. The amount that the merchant will receive,
	// denominated in `merchant_currency` and in the smallest currency unit. Note the
	// amount includes `acquirer_fee`, similar to `authorization_amount`. It will be
	// different from `authorization_amount` if the merchant is taking payment in a
	// different currency.
	//
	// Deprecated: deprecated
	MerchantAmount int64 `json:"merchant_amount,required"`
	// 3-character alphabetic ISO 4217 code for the local currency of the transaction.
	//
	// Deprecated: deprecated
	MerchantCurrency string `json:"merchant_currency,required"`
	// Deprecated, use `amounts`. Amount (in cents) of the transaction that has been
	// settled, including any acquirer fees.
	//
	// Deprecated: deprecated
	SettledAmount int64 `json:"settled_amount,required"`
	// The type of authorization request that this request is for. Note that
	// `CREDIT_AUTHORIZATION` and `FINANCIAL_CREDIT_AUTHORIZATION` is only available to
	// users with credit decisioning via ASA enabled.
	Status CardAuthorizationApprovalRequestWebhookEventStatus `json:"status,required"`
	// The entity that initiated the transaction.
	TransactionInitiator     CardAuthorizationApprovalRequestWebhookEventTransactionInitiator `json:"transaction_initiator,required"`
	AccountType              CardAuthorizationApprovalRequestWebhookEventAccountType          `json:"account_type"`
	CardholderAuthentication CardholderAuthentication                                         `json:"cardholder_authentication"`
	// Deprecated, use `cash_amount`.
	Cashback int64 `json:"cashback"`
	// Deprecated, use `amounts`. If the transaction was requested in a currency other
	// than the settlement currency, this field will be populated to indicate the rate
	// used to translate the merchant_amount to the amount (i.e., `merchant_amount` x
	// `conversion_rate` = `amount`). Note that the `merchant_amount` is in the local
	// currency and the amount is in the settlement currency.
	//
	// Deprecated: deprecated
	ConversionRate float64 `json:"conversion_rate"`
	// The event token associated with the authorization. This field is only set for
	// programs enrolled into the beta.
	EventToken string `json:"event_token" format:"uuid"`
	// Optional Object containing information if the Card is a part of a Fleet managed
	// program
	FleetInfo CardAuthorizationApprovalRequestWebhookEventFleetInfo `json:"fleet_info,nullable"`
	// The latest Authorization Challenge that was issued to the cardholder for this
	// merchant.
	LatestChallenge CardAuthorizationApprovalRequestWebhookEventLatestChallenge `json:"latest_challenge"`
	// Card network of the authorization.
	Network CardAuthorizationApprovalRequestWebhookEventNetwork `json:"network"`
	// Network-provided score assessing risk level associated with a given
	// authorization. Scores are on a range of 0-999, with 0 representing the lowest
	// risk and 999 representing the highest risk. For Visa transactions, where the raw
	// score has a range of 0-99, Lithic will normalize the score by multiplying the
	// raw score by 10x.
	NetworkRiskScore int64 `json:"network_risk_score,nullable"`
	// Contains raw data provided by the card network, including attributes that
	// provide further context about the authorization. If populated by the network,
	// data is organized by Lithic and passed through without further modification.
	// Please consult the official network documentation for more details about these
	// values and how to use them. This object is only available to certain programs-
	// contact your Customer Success Manager to discuss enabling access.
	NetworkSpecificData CardAuthorizationApprovalRequestWebhookEventNetworkSpecificData `json:"network_specific_data,nullable"`
	Pos                 CardAuthorizationApprovalRequestWebhookEventPos                 `json:"pos"`
	TokenInfo           TokenInfo                                                       `json:"token_info,nullable"`
	// Deprecated: approximate time-to-live for the authorization.
	Ttl  time.Time                                        `json:"ttl" format:"date-time"`
	JSON cardAuthorizationApprovalRequestWebhookEventJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventJSON contains the JSON metadata for
// the struct [CardAuthorizationApprovalRequestWebhookEvent]
type cardAuthorizationApprovalRequestWebhookEventJSON struct {
	Token                    apijson.Field
	AcquirerFee              apijson.Field
	Amount                   apijson.Field
	Amounts                  apijson.Field
	AuthorizationAmount      apijson.Field
	Avs                      apijson.Field
	Card                     apijson.Field
	CardholderCurrency       apijson.Field
	CashAmount               apijson.Field
	Created                  apijson.Field
	EventType                apijson.Field
	Merchant                 apijson.Field
	MerchantAmount           apijson.Field
	MerchantCurrency         apijson.Field
	SettledAmount            apijson.Field
	Status                   apijson.Field
	TransactionInitiator     apijson.Field
	AccountType              apijson.Field
	CardholderAuthentication apijson.Field
	Cashback                 apijson.Field
	ConversionRate           apijson.Field
	EventToken               apijson.Field
	FleetInfo                apijson.Field
	LatestChallenge          apijson.Field
	Network                  apijson.Field
	NetworkRiskScore         apijson.Field
	NetworkSpecificData      apijson.Field
	Pos                      apijson.Field
	TokenInfo                apijson.Field
	Ttl                      apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardAuthorizationApprovalRequestWebhookEvent) implementsParsedWebhookEvent() {}

// Structured amounts for this authorization. The `cardholder` and `merchant`
// amounts reflect the original network authorization values. For programs with
// hold adjustments enabled (e.g., automated fuel dispensers or tipping MCCs), the
// `hold` amount may exceed the `cardholder` and `merchant` amounts to account for
// anticipated final transaction amounts such as tips or fuel fill-ups
type CardAuthorizationApprovalRequestWebhookEventAmounts struct {
	Cardholder CardAuthorizationApprovalRequestWebhookEventAmountsCardholder `json:"cardholder,required"`
	Hold       CardAuthorizationApprovalRequestWebhookEventAmountsHold       `json:"hold,required,nullable"`
	Merchant   CardAuthorizationApprovalRequestWebhookEventAmountsMerchant   `json:"merchant,required"`
	Settlement CardAuthorizationApprovalRequestWebhookEventAmountsSettlement `json:"settlement,required,nullable"`
	JSON       cardAuthorizationApprovalRequestWebhookEventAmountsJSON       `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventAmountsJSON contains the JSON
// metadata for the struct [CardAuthorizationApprovalRequestWebhookEventAmounts]
type cardAuthorizationApprovalRequestWebhookEventAmountsJSON struct {
	Cardholder  apijson.Field
	Hold        apijson.Field
	Merchant    apijson.Field
	Settlement  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventAmountsJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventAmountsCardholder struct {
	// Amount in the smallest unit of the applicable currency (e.g., cents)
	Amount int64 `json:"amount,required"`
	// Exchange rate used for currency conversion
	ConversionRate string `json:"conversion_rate,required"`
	// 3-character alphabetic ISO 4217 currency
	Currency shared.Currency                                                   `json:"currency,required"`
	JSON     cardAuthorizationApprovalRequestWebhookEventAmountsCardholderJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventAmountsCardholderJSON contains the
// JSON metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventAmountsCardholder]
type cardAuthorizationApprovalRequestWebhookEventAmountsCardholderJSON struct {
	Amount         apijson.Field
	ConversionRate apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventAmountsCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventAmountsCardholderJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventAmountsHold struct {
	// Amount in the smallest unit of the applicable currency (e.g., cents)
	Amount int64 `json:"amount,required"`
	// 3-character alphabetic ISO 4217 currency
	Currency shared.Currency                                             `json:"currency,required"`
	JSON     cardAuthorizationApprovalRequestWebhookEventAmountsHoldJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventAmountsHoldJSON contains the JSON
// metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventAmountsHold]
type cardAuthorizationApprovalRequestWebhookEventAmountsHoldJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventAmountsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventAmountsHoldJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventAmountsMerchant struct {
	// Amount in the smallest unit of the applicable currency (e.g., cents)
	Amount int64 `json:"amount,required"`
	// 3-character alphabetic ISO 4217 currency
	Currency shared.Currency                                                 `json:"currency,required"`
	JSON     cardAuthorizationApprovalRequestWebhookEventAmountsMerchantJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventAmountsMerchantJSON contains the
// JSON metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventAmountsMerchant]
type cardAuthorizationApprovalRequestWebhookEventAmountsMerchantJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventAmountsMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventAmountsMerchantJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventAmountsSettlement struct {
	// Amount in the smallest unit of the applicable currency (e.g., cents)
	Amount int64 `json:"amount,required"`
	// 3-character alphabetic ISO 4217 currency
	Currency shared.Currency                                                   `json:"currency,required"`
	JSON     cardAuthorizationApprovalRequestWebhookEventAmountsSettlementJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventAmountsSettlementJSON contains the
// JSON metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventAmountsSettlement]
type cardAuthorizationApprovalRequestWebhookEventAmountsSettlementJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventAmountsSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventAmountsSettlementJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventAvs struct {
	// Cardholder address
	Address string `json:"address,required"`
	// Lithic's evaluation result comparing the transaction's address data with the
	// cardholder KYC data if it exists. In the event Lithic does not have any
	// Cardholder KYC data, or the transaction does not contain any address data,
	// NOT_PRESENT will be returned
	AddressOnFileMatch CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatch `json:"address_on_file_match,required"`
	// Cardholder ZIP code
	Zipcode string                                              `json:"zipcode,required"`
	JSON    cardAuthorizationApprovalRequestWebhookEventAvsJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventAvsJSON contains the JSON metadata
// for the struct [CardAuthorizationApprovalRequestWebhookEventAvs]
type cardAuthorizationApprovalRequestWebhookEventAvsJSON struct {
	Address            apijson.Field
	AddressOnFileMatch apijson.Field
	Zipcode            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventAvs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventAvsJSON) RawJSON() string {
	return r.raw
}

// Lithic's evaluation result comparing the transaction's address data with the
// cardholder KYC data if it exists. In the event Lithic does not have any
// Cardholder KYC data, or the transaction does not contain any address data,
// NOT_PRESENT will be returned
type CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatch string

const (
	CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchMatch            CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatch = "MATCH"
	CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchMatchAddressOnly CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatch = "MATCH_ADDRESS_ONLY"
	CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchMatchZipOnly     CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatch = "MATCH_ZIP_ONLY"
	CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchMismatch         CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatch = "MISMATCH"
	CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchNotPresent       CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatch = "NOT_PRESENT"
)

func (r CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatch) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchMatch, CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchMatchAddressOnly, CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchMatchZipOnly, CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchMismatch, CardAuthorizationApprovalRequestWebhookEventAvsAddressOnFileMatchNotPresent:
		return true
	}
	return false
}

// Card object in ASA
type CardAuthorizationApprovalRequestWebhookEventCard struct {
	// Globally unique identifier for the card.
	Token string `json:"token" format:"uuid"`
	// Hostname of card’s locked merchant (will be empty if not applicable)
	Hostname string `json:"hostname"`
	// Last four digits of the card number
	LastFour string `json:"last_four"`
	// Customizable name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo string `json:"memo"`
	// Amount (in cents) to limit approved authorizations. Purchase requests above the
	// spend limit will be declined (refunds and credits will be approved).
	//
	// Note that while spend limits are enforced based on authorized and settled volume
	// on a card, they are not recommended to be used for balance or
	// reconciliation-level accuracy. Spend limits also cannot block force posted
	// charges (i.e., when a merchant sends a clearing message without a prior
	// authorization).
	SpendLimit int64 `json:"spend_limit"`
	// Note that to support recurring monthly payments, which can occur on different
	// day every month, the time window we consider for MONTHLY velocity starts 6 days
	// after the current calendar date one month prior.
	SpendLimitDuration CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDuration `json:"spend_limit_duration"`
	State              CardAuthorizationApprovalRequestWebhookEventCardState              `json:"state"`
	Type               CardAuthorizationApprovalRequestWebhookEventCardType               `json:"type"`
	JSON               cardAuthorizationApprovalRequestWebhookEventCardJSON               `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventCardJSON contains the JSON metadata
// for the struct [CardAuthorizationApprovalRequestWebhookEventCard]
type cardAuthorizationApprovalRequestWebhookEventCardJSON struct {
	Token              apijson.Field
	Hostname           apijson.Field
	LastFour           apijson.Field
	Memo               apijson.Field
	SpendLimit         apijson.Field
	SpendLimitDuration apijson.Field
	State              apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventCardJSON) RawJSON() string {
	return r.raw
}

// Note that to support recurring monthly payments, which can occur on different
// day every month, the time window we consider for MONTHLY velocity starts 6 days
// after the current calendar date one month prior.
type CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDuration string

const (
	CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDurationAnnually    CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDuration = "ANNUALLY"
	CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDurationForever     CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDuration = "FOREVER"
	CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDurationMonthly     CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDuration = "MONTHLY"
	CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDurationTransaction CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDuration = "TRANSACTION"
)

func (r CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDurationAnnually, CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDurationForever, CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDurationMonthly, CardAuthorizationApprovalRequestWebhookEventCardSpendLimitDurationTransaction:
		return true
	}
	return false
}

type CardAuthorizationApprovalRequestWebhookEventCardState string

const (
	CardAuthorizationApprovalRequestWebhookEventCardStateClosed             CardAuthorizationApprovalRequestWebhookEventCardState = "CLOSED"
	CardAuthorizationApprovalRequestWebhookEventCardStateOpen               CardAuthorizationApprovalRequestWebhookEventCardState = "OPEN"
	CardAuthorizationApprovalRequestWebhookEventCardStatePaused             CardAuthorizationApprovalRequestWebhookEventCardState = "PAUSED"
	CardAuthorizationApprovalRequestWebhookEventCardStatePendingActivation  CardAuthorizationApprovalRequestWebhookEventCardState = "PENDING_ACTIVATION"
	CardAuthorizationApprovalRequestWebhookEventCardStatePendingFulfillment CardAuthorizationApprovalRequestWebhookEventCardState = "PENDING_FULFILLMENT"
)

func (r CardAuthorizationApprovalRequestWebhookEventCardState) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventCardStateClosed, CardAuthorizationApprovalRequestWebhookEventCardStateOpen, CardAuthorizationApprovalRequestWebhookEventCardStatePaused, CardAuthorizationApprovalRequestWebhookEventCardStatePendingActivation, CardAuthorizationApprovalRequestWebhookEventCardStatePendingFulfillment:
		return true
	}
	return false
}

type CardAuthorizationApprovalRequestWebhookEventCardType string

const (
	CardAuthorizationApprovalRequestWebhookEventCardTypeSingleUse      CardAuthorizationApprovalRequestWebhookEventCardType = "SINGLE_USE"
	CardAuthorizationApprovalRequestWebhookEventCardTypeMerchantLocked CardAuthorizationApprovalRequestWebhookEventCardType = "MERCHANT_LOCKED"
	CardAuthorizationApprovalRequestWebhookEventCardTypeUnlocked       CardAuthorizationApprovalRequestWebhookEventCardType = "UNLOCKED"
	CardAuthorizationApprovalRequestWebhookEventCardTypePhysical       CardAuthorizationApprovalRequestWebhookEventCardType = "PHYSICAL"
	CardAuthorizationApprovalRequestWebhookEventCardTypeDigitalWallet  CardAuthorizationApprovalRequestWebhookEventCardType = "DIGITAL_WALLET"
	CardAuthorizationApprovalRequestWebhookEventCardTypeVirtual        CardAuthorizationApprovalRequestWebhookEventCardType = "VIRTUAL"
)

func (r CardAuthorizationApprovalRequestWebhookEventCardType) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventCardTypeSingleUse, CardAuthorizationApprovalRequestWebhookEventCardTypeMerchantLocked, CardAuthorizationApprovalRequestWebhookEventCardTypeUnlocked, CardAuthorizationApprovalRequestWebhookEventCardTypePhysical, CardAuthorizationApprovalRequestWebhookEventCardTypeDigitalWallet, CardAuthorizationApprovalRequestWebhookEventCardTypeVirtual:
		return true
	}
	return false
}

type CardAuthorizationApprovalRequestWebhookEventEventType string

const (
	CardAuthorizationApprovalRequestWebhookEventEventTypeCardAuthorizationApprovalRequest CardAuthorizationApprovalRequestWebhookEventEventType = "card_authorization.approval_request"
)

func (r CardAuthorizationApprovalRequestWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventEventTypeCardAuthorizationApprovalRequest:
		return true
	}
	return false
}

// The type of authorization request that this request is for. Note that
// `CREDIT_AUTHORIZATION` and `FINANCIAL_CREDIT_AUTHORIZATION` is only available to
// users with credit decisioning via ASA enabled.
type CardAuthorizationApprovalRequestWebhookEventStatus string

const (
	CardAuthorizationApprovalRequestWebhookEventStatusAuthorization                CardAuthorizationApprovalRequestWebhookEventStatus = "AUTHORIZATION"
	CardAuthorizationApprovalRequestWebhookEventStatusCreditAuthorization          CardAuthorizationApprovalRequestWebhookEventStatus = "CREDIT_AUTHORIZATION"
	CardAuthorizationApprovalRequestWebhookEventStatusFinancialAuthorization       CardAuthorizationApprovalRequestWebhookEventStatus = "FINANCIAL_AUTHORIZATION"
	CardAuthorizationApprovalRequestWebhookEventStatusFinancialCreditAuthorization CardAuthorizationApprovalRequestWebhookEventStatus = "FINANCIAL_CREDIT_AUTHORIZATION"
	CardAuthorizationApprovalRequestWebhookEventStatusBalanceInquiry               CardAuthorizationApprovalRequestWebhookEventStatus = "BALANCE_INQUIRY"
)

func (r CardAuthorizationApprovalRequestWebhookEventStatus) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventStatusAuthorization, CardAuthorizationApprovalRequestWebhookEventStatusCreditAuthorization, CardAuthorizationApprovalRequestWebhookEventStatusFinancialAuthorization, CardAuthorizationApprovalRequestWebhookEventStatusFinancialCreditAuthorization, CardAuthorizationApprovalRequestWebhookEventStatusBalanceInquiry:
		return true
	}
	return false
}

// The entity that initiated the transaction.
type CardAuthorizationApprovalRequestWebhookEventTransactionInitiator string

const (
	CardAuthorizationApprovalRequestWebhookEventTransactionInitiatorCardholder CardAuthorizationApprovalRequestWebhookEventTransactionInitiator = "CARDHOLDER"
	CardAuthorizationApprovalRequestWebhookEventTransactionInitiatorMerchant   CardAuthorizationApprovalRequestWebhookEventTransactionInitiator = "MERCHANT"
	CardAuthorizationApprovalRequestWebhookEventTransactionInitiatorUnknown    CardAuthorizationApprovalRequestWebhookEventTransactionInitiator = "UNKNOWN"
)

func (r CardAuthorizationApprovalRequestWebhookEventTransactionInitiator) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventTransactionInitiatorCardholder, CardAuthorizationApprovalRequestWebhookEventTransactionInitiatorMerchant, CardAuthorizationApprovalRequestWebhookEventTransactionInitiatorUnknown:
		return true
	}
	return false
}

type CardAuthorizationApprovalRequestWebhookEventAccountType string

const (
	CardAuthorizationApprovalRequestWebhookEventAccountTypeChecking CardAuthorizationApprovalRequestWebhookEventAccountType = "CHECKING"
	CardAuthorizationApprovalRequestWebhookEventAccountTypeSavings  CardAuthorizationApprovalRequestWebhookEventAccountType = "SAVINGS"
)

func (r CardAuthorizationApprovalRequestWebhookEventAccountType) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventAccountTypeChecking, CardAuthorizationApprovalRequestWebhookEventAccountTypeSavings:
		return true
	}
	return false
}

// Optional Object containing information if the Card is a part of a Fleet managed
// program
type CardAuthorizationApprovalRequestWebhookEventFleetInfo struct {
	// Code indicating what the driver was prompted to enter at time of purchase. This
	// is configured at a program level and is a static configuration, and does not
	// change on a request to request basis
	FleetPromptCode CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCode `json:"fleet_prompt_code,required"`
	// Code indicating which restrictions, if any, there are on purchase. This is
	// configured at a program level and is a static configuration, and does not change
	// on a request to request basis
	FleetRestrictionCode CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCode `json:"fleet_restriction_code,required"`
	// Number representing the driver
	DriverNumber string `json:"driver_number,nullable"`
	// Number associated with the vehicle
	VehicleNumber string                                                    `json:"vehicle_number,nullable"`
	JSON          cardAuthorizationApprovalRequestWebhookEventFleetInfoJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventFleetInfoJSON contains the JSON
// metadata for the struct [CardAuthorizationApprovalRequestWebhookEventFleetInfo]
type cardAuthorizationApprovalRequestWebhookEventFleetInfoJSON struct {
	FleetPromptCode      apijson.Field
	FleetRestrictionCode apijson.Field
	DriverNumber         apijson.Field
	VehicleNumber        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventFleetInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventFleetInfoJSON) RawJSON() string {
	return r.raw
}

// Code indicating what the driver was prompted to enter at time of purchase. This
// is configured at a program level and is a static configuration, and does not
// change on a request to request basis
type CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCode string

const (
	CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCodeNoPrompt      CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCode = "NO_PROMPT"
	CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCodeVehicleNumber CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCode = "VEHICLE_NUMBER"
	CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCodeDriverNumber  CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCode = "DRIVER_NUMBER"
)

func (r CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCode) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCodeNoPrompt, CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCodeVehicleNumber, CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetPromptCodeDriverNumber:
		return true
	}
	return false
}

// Code indicating which restrictions, if any, there are on purchase. This is
// configured at a program level and is a static configuration, and does not change
// on a request to request basis
type CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCode string

const (
	CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCodeNoRestrictions CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCode = "NO_RESTRICTIONS"
	CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCodeFuelOnly       CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCode = "FUEL_ONLY"
)

func (r CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCode) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCodeNoRestrictions, CardAuthorizationApprovalRequestWebhookEventFleetInfoFleetRestrictionCodeFuelOnly:
		return true
	}
	return false
}

// The latest Authorization Challenge that was issued to the cardholder for this
// merchant.
type CardAuthorizationApprovalRequestWebhookEventLatestChallenge struct {
	// The phone number used for sending Authorization Challenge SMS.
	PhoneNumber string `json:"phone_number,required"`
	// The status of the Authorization Challenge
	//
	// - `COMPLETED` - Challenge was successfully completed by the cardholder
	// - `PENDING` - Challenge is still open
	// - `EXPIRED` - Challenge has expired without being completed
	// - `ERROR` - There was an error processing the challenge
	Status CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatus `json:"status,required"`
	// The date and time when the Authorization Challenge was completed in UTC. Present
	// only if the status is `COMPLETED`.
	CompletedAt time.Time                                                       `json:"completed_at" format:"date-time"`
	JSON        cardAuthorizationApprovalRequestWebhookEventLatestChallengeJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventLatestChallengeJSON contains the
// JSON metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventLatestChallenge]
type cardAuthorizationApprovalRequestWebhookEventLatestChallengeJSON struct {
	PhoneNumber apijson.Field
	Status      apijson.Field
	CompletedAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventLatestChallenge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventLatestChallengeJSON) RawJSON() string {
	return r.raw
}

// The status of the Authorization Challenge
//
// - `COMPLETED` - Challenge was successfully completed by the cardholder
// - `PENDING` - Challenge is still open
// - `EXPIRED` - Challenge has expired without being completed
// - `ERROR` - There was an error processing the challenge
type CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatus string

const (
	CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatusCompleted CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatus = "COMPLETED"
	CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatusPending   CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatus = "PENDING"
	CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatusExpired   CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatus = "EXPIRED"
	CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatusError     CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatus = "ERROR"
)

func (r CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatus) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatusCompleted, CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatusPending, CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatusExpired, CardAuthorizationApprovalRequestWebhookEventLatestChallengeStatusError:
		return true
	}
	return false
}

// Card network of the authorization.
type CardAuthorizationApprovalRequestWebhookEventNetwork string

const (
	CardAuthorizationApprovalRequestWebhookEventNetworkAmex       CardAuthorizationApprovalRequestWebhookEventNetwork = "AMEX"
	CardAuthorizationApprovalRequestWebhookEventNetworkInterlink  CardAuthorizationApprovalRequestWebhookEventNetwork = "INTERLINK"
	CardAuthorizationApprovalRequestWebhookEventNetworkMaestro    CardAuthorizationApprovalRequestWebhookEventNetwork = "MAESTRO"
	CardAuthorizationApprovalRequestWebhookEventNetworkMastercard CardAuthorizationApprovalRequestWebhookEventNetwork = "MASTERCARD"
	CardAuthorizationApprovalRequestWebhookEventNetworkUnknown    CardAuthorizationApprovalRequestWebhookEventNetwork = "UNKNOWN"
	CardAuthorizationApprovalRequestWebhookEventNetworkVisa       CardAuthorizationApprovalRequestWebhookEventNetwork = "VISA"
)

func (r CardAuthorizationApprovalRequestWebhookEventNetwork) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventNetworkAmex, CardAuthorizationApprovalRequestWebhookEventNetworkInterlink, CardAuthorizationApprovalRequestWebhookEventNetworkMaestro, CardAuthorizationApprovalRequestWebhookEventNetworkMastercard, CardAuthorizationApprovalRequestWebhookEventNetworkUnknown, CardAuthorizationApprovalRequestWebhookEventNetworkVisa:
		return true
	}
	return false
}

// Contains raw data provided by the card network, including attributes that
// provide further context about the authorization. If populated by the network,
// data is organized by Lithic and passed through without further modification.
// Please consult the official network documentation for more details about these
// values and how to use them. This object is only available to certain programs-
// contact your Customer Success Manager to discuss enabling access.
type CardAuthorizationApprovalRequestWebhookEventNetworkSpecificData struct {
	Mastercard CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercard `json:"mastercard,nullable"`
	Visa       CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataVisa       `json:"visa,nullable"`
	JSON       cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataJSON       `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataJSON contains the
// JSON metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventNetworkSpecificData]
type cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataJSON struct {
	Mastercard  apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventNetworkSpecificData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercard struct {
	// Indicates the electronic commerce security level and UCAF collection.
	EcommerceSecurityLevelIndicator string `json:"ecommerce_security_level_indicator,nullable"`
	// The On-behalf Service performed on the transaction and the results. Contains all
	// applicable, on-behalf service results that were performed on a given
	// transaction.
	OnBehalfServiceResult []CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardOnBehalfServiceResult `json:"on_behalf_service_result,nullable"`
	// Indicates the type of additional transaction purpose.
	TransactionTypeIdentifier string                                                                        `json:"transaction_type_identifier,nullable"`
	JSON                      cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardJSON
// contains the JSON metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercard]
type cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardJSON struct {
	EcommerceSecurityLevelIndicator apijson.Field
	OnBehalfServiceResult           apijson.Field
	TransactionTypeIdentifier       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardOnBehalfServiceResult struct {
	// Indicates the results of the service processing.
	Result1 string `json:"result_1,required"`
	// Identifies the results of the service processing.
	Result2 string `json:"result_2,required"`
	// Indicates the service performed on the transaction.
	Service string                                                                                             `json:"service,required"`
	JSON    cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardOnBehalfServiceResultJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardOnBehalfServiceResultJSON
// contains the JSON metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardOnBehalfServiceResult]
type cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardOnBehalfServiceResultJSON struct {
	Result1     apijson.Field
	Result2     apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardOnBehalfServiceResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataMastercardOnBehalfServiceResultJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataVisa struct {
	// Identifies the purpose or category of a transaction, used to classify and
	// process transactions according to Visa’s rules.
	BusinessApplicationIdentifier string                                                                  `json:"business_application_identifier,nullable"`
	JSON                          cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataVisaJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataVisaJSON contains
// the JSON metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataVisa]
type cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataVisaJSON struct {
	BusinessApplicationIdentifier apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventNetworkSpecificDataVisaJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationApprovalRequestWebhookEventPos struct {
	// POS > Entry Mode object in ASA
	EntryMode CardAuthorizationApprovalRequestWebhookEventPosEntryMode `json:"entry_mode"`
	Terminal  CardAuthorizationApprovalRequestWebhookEventPosTerminal  `json:"terminal"`
	JSON      cardAuthorizationApprovalRequestWebhookEventPosJSON      `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventPosJSON contains the JSON metadata
// for the struct [CardAuthorizationApprovalRequestWebhookEventPos]
type cardAuthorizationApprovalRequestWebhookEventPosJSON struct {
	EntryMode   apijson.Field
	Terminal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventPos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventPosJSON) RawJSON() string {
	return r.raw
}

// POS > Entry Mode object in ASA
type CardAuthorizationApprovalRequestWebhookEventPosEntryMode struct {
	// Card Presence Indicator
	Card CardAuthorizationApprovalRequestWebhookEventPosEntryModeCard `json:"card"`
	// Cardholder Presence Indicator
	Cardholder CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder `json:"cardholder"`
	// Method of entry for the PAN
	Pan CardAuthorizationApprovalRequestWebhookEventPosEntryModePan `json:"pan"`
	// Indicates whether the cardholder entered the PIN. True if the PIN was entered.
	PinEntered bool                                                         `json:"pin_entered"`
	JSON       cardAuthorizationApprovalRequestWebhookEventPosEntryModeJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventPosEntryModeJSON contains the JSON
// metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventPosEntryMode]
type cardAuthorizationApprovalRequestWebhookEventPosEntryModeJSON struct {
	Card        apijson.Field
	Cardholder  apijson.Field
	Pan         apijson.Field
	PinEntered  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventPosEntryMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventPosEntryModeJSON) RawJSON() string {
	return r.raw
}

// Card Presence Indicator
type CardAuthorizationApprovalRequestWebhookEventPosEntryModeCard string

const (
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardPresent    CardAuthorizationApprovalRequestWebhookEventPosEntryModeCard = "PRESENT"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardNotPresent CardAuthorizationApprovalRequestWebhookEventPosEntryModeCard = "NOT_PRESENT"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardUnknown    CardAuthorizationApprovalRequestWebhookEventPosEntryModeCard = "UNKNOWN"
)

func (r CardAuthorizationApprovalRequestWebhookEventPosEntryModeCard) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardPresent, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardNotPresent, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardUnknown:
		return true
	}
	return false
}

// Cardholder Presence Indicator
type CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder string

const (
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderDeferredBilling CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "DEFERRED_BILLING"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderElectronicOrder CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "ELECTRONIC_ORDER"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderInstallment     CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "INSTALLMENT"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderMailOrder       CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "MAIL_ORDER"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderNotPresent      CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "NOT_PRESENT"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderPresent         CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "PRESENT"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderReoccurring     CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "REOCCURRING"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderTelephoneOrder  CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "TELEPHONE_ORDER"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderUnknown         CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder = "UNKNOWN"
)

func (r CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholder) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderDeferredBilling, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderElectronicOrder, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderInstallment, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderMailOrder, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderNotPresent, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderPresent, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderReoccurring, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderTelephoneOrder, CardAuthorizationApprovalRequestWebhookEventPosEntryModeCardholderUnknown:
		return true
	}
	return false
}

// Method of entry for the PAN
type CardAuthorizationApprovalRequestWebhookEventPosEntryModePan string

const (
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanAutoEntry           CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "AUTO_ENTRY"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanBarCode             CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "BAR_CODE"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanContactless         CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "CONTACTLESS"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanEcommerce           CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "ECOMMERCE"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanErrorKeyed          CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "ERROR_KEYED"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanErrorMagneticStripe CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "ERROR_MAGNETIC_STRIPE"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanIcc                 CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "ICC"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanKeyEntered          CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "KEY_ENTERED"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanMagneticStripe      CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "MAGNETIC_STRIPE"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanManual              CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "MANUAL"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanOcr                 CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "OCR"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanSecureCardless      CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "SECURE_CARDLESS"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanUnspecified         CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "UNSPECIFIED"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanUnknown             CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "UNKNOWN"
	CardAuthorizationApprovalRequestWebhookEventPosEntryModePanCredentialOnFile    CardAuthorizationApprovalRequestWebhookEventPosEntryModePan = "CREDENTIAL_ON_FILE"
)

func (r CardAuthorizationApprovalRequestWebhookEventPosEntryModePan) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventPosEntryModePanAutoEntry, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanBarCode, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanContactless, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanEcommerce, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanErrorKeyed, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanErrorMagneticStripe, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanIcc, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanKeyEntered, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanMagneticStripe, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanManual, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanOcr, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanSecureCardless, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanUnspecified, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanUnknown, CardAuthorizationApprovalRequestWebhookEventPosEntryModePanCredentialOnFile:
		return true
	}
	return false
}

type CardAuthorizationApprovalRequestWebhookEventPosTerminal struct {
	// True if a clerk is present at the sale.
	Attended bool `json:"attended,required"`
	// True if the terminal is capable of retaining the card.
	CardRetentionCapable bool `json:"card_retention_capable,required"`
	// True if the sale was made at the place of business (vs. mobile).
	OnPremise bool `json:"on_premise,required"`
	// The person that is designated to swipe the card
	Operator CardAuthorizationApprovalRequestWebhookEventPosTerminalOperator `json:"operator,required"`
	// True if the terminal is capable of partial approval. Partial approval is when
	// part of a transaction is approved and another payment must be used for the
	// remainder. Example scenario: A $40 transaction is attempted on a prepaid card
	// with a $25 balance. If partial approval is enabled, $25 can be authorized, at
	// which point the POS will prompt the user for an additional payment of $15.
	PartialApprovalCapable bool `json:"partial_approval_capable,required"`
	// Status of whether the POS is able to accept PINs
	PinCapability CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapability `json:"pin_capability,required"`
	// POS Type
	Type CardAuthorizationApprovalRequestWebhookEventPosTerminalType `json:"type,required"`
	// Uniquely identifies a terminal at the card acceptor location of acquiring
	// institutions or merchant POS Systems. Left justified with trailing spaces.
	AcceptorTerminalID string                                                      `json:"acceptor_terminal_id,nullable"`
	JSON               cardAuthorizationApprovalRequestWebhookEventPosTerminalJSON `json:"-"`
}

// cardAuthorizationApprovalRequestWebhookEventPosTerminalJSON contains the JSON
// metadata for the struct
// [CardAuthorizationApprovalRequestWebhookEventPosTerminal]
type cardAuthorizationApprovalRequestWebhookEventPosTerminalJSON struct {
	Attended               apijson.Field
	CardRetentionCapable   apijson.Field
	OnPremise              apijson.Field
	Operator               apijson.Field
	PartialApprovalCapable apijson.Field
	PinCapability          apijson.Field
	Type                   apijson.Field
	AcceptorTerminalID     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CardAuthorizationApprovalRequestWebhookEventPosTerminal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationApprovalRequestWebhookEventPosTerminalJSON) RawJSON() string {
	return r.raw
}

// The person that is designated to swipe the card
type CardAuthorizationApprovalRequestWebhookEventPosTerminalOperator string

const (
	CardAuthorizationApprovalRequestWebhookEventPosTerminalOperatorAdministrative CardAuthorizationApprovalRequestWebhookEventPosTerminalOperator = "ADMINISTRATIVE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalOperatorCardholder     CardAuthorizationApprovalRequestWebhookEventPosTerminalOperator = "CARDHOLDER"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalOperatorCardAcceptor   CardAuthorizationApprovalRequestWebhookEventPosTerminalOperator = "CARD_ACCEPTOR"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalOperatorUnknown        CardAuthorizationApprovalRequestWebhookEventPosTerminalOperator = "UNKNOWN"
)

func (r CardAuthorizationApprovalRequestWebhookEventPosTerminalOperator) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventPosTerminalOperatorAdministrative, CardAuthorizationApprovalRequestWebhookEventPosTerminalOperatorCardholder, CardAuthorizationApprovalRequestWebhookEventPosTerminalOperatorCardAcceptor, CardAuthorizationApprovalRequestWebhookEventPosTerminalOperatorUnknown:
		return true
	}
	return false
}

// Status of whether the POS is able to accept PINs
type CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapability string

const (
	CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapabilityCapable     CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapability = "CAPABLE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapabilityInoperative CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapability = "INOPERATIVE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapabilityNotCapable  CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapability = "NOT_CAPABLE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapabilityUnspecified CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapability = "UNSPECIFIED"
)

func (r CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapability) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapabilityCapable, CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapabilityInoperative, CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapabilityNotCapable, CardAuthorizationApprovalRequestWebhookEventPosTerminalPinCapabilityUnspecified:
		return true
	}
	return false
}

// POS Type
type CardAuthorizationApprovalRequestWebhookEventPosTerminalType string

const (
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeAdministrative        CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "ADMINISTRATIVE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeAtm                   CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "ATM"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeAuthorization         CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "AUTHORIZATION"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeCouponMachine         CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "COUPON_MACHINE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeDialTerminal          CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "DIAL_TERMINAL"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeEcommerce             CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "ECOMMERCE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeEcr                   CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "ECR"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeFuelMachine           CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "FUEL_MACHINE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeHomeTerminal          CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "HOME_TERMINAL"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeMicr                  CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "MICR"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeOffPremise            CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "OFF_PREMISE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePayment               CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "PAYMENT"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePda                   CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "PDA"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePhone                 CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "PHONE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePoint                 CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "POINT"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePosTerminal           CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "POS_TERMINAL"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePublicUtility         CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "PUBLIC_UTILITY"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeSelfService           CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "SELF_SERVICE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeTelevision            CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "TELEVISION"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeTeller                CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "TELLER"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeTravelersCheckMachine CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "TRAVELERS_CHECK_MACHINE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeVending               CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "VENDING"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeVoice                 CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "VOICE"
	CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeUnknown               CardAuthorizationApprovalRequestWebhookEventPosTerminalType = "UNKNOWN"
)

func (r CardAuthorizationApprovalRequestWebhookEventPosTerminalType) IsKnown() bool {
	switch r {
	case CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeAdministrative, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeAtm, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeAuthorization, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeCouponMachine, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeDialTerminal, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeEcommerce, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeEcr, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeFuelMachine, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeHomeTerminal, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeMicr, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeOffPremise, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePayment, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePda, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePhone, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePoint, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePosTerminal, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypePublicUtility, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeSelfService, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeTelevision, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeTeller, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeTravelersCheckMachine, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeVending, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeVoice, CardAuthorizationApprovalRequestWebhookEventPosTerminalTypeUnknown:
		return true
	}
	return false
}

// A webhook for tokenization decisioning sent to the customer's responder endpoint
type TokenizationDecisioningRequestWebhookEvent struct {
	// Unique identifier for the user tokenizing a card
	AccountToken string `json:"account_token,required"`
	// Unique identifier for the card being tokenized
	CardToken string `json:"card_token,required"`
	// Indicate when the request was received from Mastercard or Visa
	Created time.Time `json:"created,required" format:"date-time"`
	// The name of this event
	EventType TokenizationDecisioningRequestWebhookEventEventType `json:"event_type,required"`
	// Whether Lithic decisioned on the token, and if so, what the decision was.
	// APPROVED/VERIFICATION_REQUIRED/DENIED.
	IssuerDecision TokenizationDecisioningRequestWebhookEventIssuerDecision `json:"issuer_decision,required"`
	// The channel through which the tokenization was made.
	TokenizationChannel TokenizationDecisioningRequestWebhookEventTokenizationChannel `json:"tokenization_channel,required"`
	// Unique identifier for the digital wallet token attempt
	TokenizationToken     string                `json:"tokenization_token,required"`
	WalletDecisioningInfo WalletDecisioningInfo `json:"wallet_decisioning_info,required"`
	Device                Device                `json:"device"`
	// Contains the metadata for the digital wallet being tokenized.
	DigitalWalletTokenMetadata DigitalWalletTokenMetadata `json:"digital_wallet_token_metadata"`
	// The source of the tokenization.
	TokenizationSource TokenizationDecisioningRequestWebhookEventTokenizationSource `json:"tokenization_source"`
	JSON               tokenizationDecisioningRequestWebhookEventJSON               `json:"-"`
}

// tokenizationDecisioningRequestWebhookEventJSON contains the JSON metadata for
// the struct [TokenizationDecisioningRequestWebhookEvent]
type tokenizationDecisioningRequestWebhookEventJSON struct {
	AccountToken               apijson.Field
	CardToken                  apijson.Field
	Created                    apijson.Field
	EventType                  apijson.Field
	IssuerDecision             apijson.Field
	TokenizationChannel        apijson.Field
	TokenizationToken          apijson.Field
	WalletDecisioningInfo      apijson.Field
	Device                     apijson.Field
	DigitalWalletTokenMetadata apijson.Field
	TokenizationSource         apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *TokenizationDecisioningRequestWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationDecisioningRequestWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r TokenizationDecisioningRequestWebhookEvent) implementsParsedWebhookEvent() {}

// The name of this event
type TokenizationDecisioningRequestWebhookEventEventType string

const (
	TokenizationDecisioningRequestWebhookEventEventTypeDigitalWalletTokenizationApprovalRequest TokenizationDecisioningRequestWebhookEventEventType = "digital_wallet.tokenization_approval_request"
)

func (r TokenizationDecisioningRequestWebhookEventEventType) IsKnown() bool {
	switch r {
	case TokenizationDecisioningRequestWebhookEventEventTypeDigitalWalletTokenizationApprovalRequest:
		return true
	}
	return false
}

// Whether Lithic decisioned on the token, and if so, what the decision was.
// APPROVED/VERIFICATION_REQUIRED/DENIED.
type TokenizationDecisioningRequestWebhookEventIssuerDecision string

const (
	TokenizationDecisioningRequestWebhookEventIssuerDecisionApproved             TokenizationDecisioningRequestWebhookEventIssuerDecision = "APPROVED"
	TokenizationDecisioningRequestWebhookEventIssuerDecisionDenied               TokenizationDecisioningRequestWebhookEventIssuerDecision = "DENIED"
	TokenizationDecisioningRequestWebhookEventIssuerDecisionVerificationRequired TokenizationDecisioningRequestWebhookEventIssuerDecision = "VERIFICATION_REQUIRED"
)

func (r TokenizationDecisioningRequestWebhookEventIssuerDecision) IsKnown() bool {
	switch r {
	case TokenizationDecisioningRequestWebhookEventIssuerDecisionApproved, TokenizationDecisioningRequestWebhookEventIssuerDecisionDenied, TokenizationDecisioningRequestWebhookEventIssuerDecisionVerificationRequired:
		return true
	}
	return false
}

// The channel through which the tokenization was made.
type TokenizationDecisioningRequestWebhookEventTokenizationChannel string

const (
	TokenizationDecisioningRequestWebhookEventTokenizationChannelDigitalWallet TokenizationDecisioningRequestWebhookEventTokenizationChannel = "DIGITAL_WALLET"
	TokenizationDecisioningRequestWebhookEventTokenizationChannelMerchant      TokenizationDecisioningRequestWebhookEventTokenizationChannel = "MERCHANT"
)

func (r TokenizationDecisioningRequestWebhookEventTokenizationChannel) IsKnown() bool {
	switch r {
	case TokenizationDecisioningRequestWebhookEventTokenizationChannelDigitalWallet, TokenizationDecisioningRequestWebhookEventTokenizationChannelMerchant:
		return true
	}
	return false
}

// The source of the tokenization.
type TokenizationDecisioningRequestWebhookEventTokenizationSource string

const (
	TokenizationDecisioningRequestWebhookEventTokenizationSourceAccountOnFile   TokenizationDecisioningRequestWebhookEventTokenizationSource = "ACCOUNT_ON_FILE"
	TokenizationDecisioningRequestWebhookEventTokenizationSourceContactlessTap  TokenizationDecisioningRequestWebhookEventTokenizationSource = "CONTACTLESS_TAP"
	TokenizationDecisioningRequestWebhookEventTokenizationSourceManualProvision TokenizationDecisioningRequestWebhookEventTokenizationSource = "MANUAL_PROVISION"
	TokenizationDecisioningRequestWebhookEventTokenizationSourcePushProvision   TokenizationDecisioningRequestWebhookEventTokenizationSource = "PUSH_PROVISION"
	TokenizationDecisioningRequestWebhookEventTokenizationSourceToken           TokenizationDecisioningRequestWebhookEventTokenizationSource = "TOKEN"
	TokenizationDecisioningRequestWebhookEventTokenizationSourceUnknown         TokenizationDecisioningRequestWebhookEventTokenizationSource = "UNKNOWN"
)

func (r TokenizationDecisioningRequestWebhookEventTokenizationSource) IsKnown() bool {
	switch r {
	case TokenizationDecisioningRequestWebhookEventTokenizationSourceAccountOnFile, TokenizationDecisioningRequestWebhookEventTokenizationSourceContactlessTap, TokenizationDecisioningRequestWebhookEventTokenizationSourceManualProvision, TokenizationDecisioningRequestWebhookEventTokenizationSourcePushProvision, TokenizationDecisioningRequestWebhookEventTokenizationSourceToken, TokenizationDecisioningRequestWebhookEventTokenizationSourceUnknown:
		return true
	}
	return false
}

type AuthRulesBacktestReportCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType AuthRulesBacktestReportCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      authRulesBacktestReportCreatedWebhookEventJSON      `json:"-"`
	BacktestResults
}

// authRulesBacktestReportCreatedWebhookEventJSON contains the JSON metadata for
// the struct [AuthRulesBacktestReportCreatedWebhookEvent]
type authRulesBacktestReportCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRulesBacktestReportCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRulesBacktestReportCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r AuthRulesBacktestReportCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type AuthRulesBacktestReportCreatedWebhookEventEventType string

const (
	AuthRulesBacktestReportCreatedWebhookEventEventTypeAuthRulesBacktestReportCreated AuthRulesBacktestReportCreatedWebhookEventEventType = "auth_rules.backtest_report.created"
)

func (r AuthRulesBacktestReportCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case AuthRulesBacktestReportCreatedWebhookEventEventTypeAuthRulesBacktestReportCreated:
		return true
	}
	return false
}

type BalanceUpdatedWebhookEvent struct {
	Data []FinancialAccountBalance `json:"data,required"`
	// The type of event that occurred.
	EventType BalanceUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      balanceUpdatedWebhookEventJSON      `json:"-"`
}

// balanceUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [BalanceUpdatedWebhookEvent]
type balanceUpdatedWebhookEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BalanceUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r BalanceUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type BalanceUpdatedWebhookEventEventType string

const (
	BalanceUpdatedWebhookEventEventTypeBalanceUpdated BalanceUpdatedWebhookEventEventType = "balance.updated"
)

func (r BalanceUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case BalanceUpdatedWebhookEventEventTypeBalanceUpdated:
		return true
	}
	return false
}

// Book transfer transaction
type BookTransferTransactionCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType BookTransferTransactionCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      bookTransferTransactionCreatedWebhookEventJSON      `json:"-"`
	BookTransferResponse
}

// bookTransferTransactionCreatedWebhookEventJSON contains the JSON metadata for
// the struct [BookTransferTransactionCreatedWebhookEvent]
type bookTransferTransactionCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookTransferTransactionCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookTransferTransactionCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r BookTransferTransactionCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type BookTransferTransactionCreatedWebhookEventEventType string

const (
	BookTransferTransactionCreatedWebhookEventEventTypeBookTransferTransactionCreated BookTransferTransactionCreatedWebhookEventEventType = "book_transfer_transaction.created"
)

func (r BookTransferTransactionCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case BookTransferTransactionCreatedWebhookEventEventTypeBookTransferTransactionCreated:
		return true
	}
	return false
}

// Book transfer transaction
type BookTransferTransactionUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType BookTransferTransactionUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      bookTransferTransactionUpdatedWebhookEventJSON      `json:"-"`
	BookTransferResponse
}

// bookTransferTransactionUpdatedWebhookEventJSON contains the JSON metadata for
// the struct [BookTransferTransactionUpdatedWebhookEvent]
type bookTransferTransactionUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookTransferTransactionUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookTransferTransactionUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r BookTransferTransactionUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type BookTransferTransactionUpdatedWebhookEventEventType string

const (
	BookTransferTransactionUpdatedWebhookEventEventTypeBookTransferTransactionUpdated BookTransferTransactionUpdatedWebhookEventEventType = "book_transfer_transaction.updated"
)

func (r BookTransferTransactionUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case BookTransferTransactionUpdatedWebhookEventEventTypeBookTransferTransactionUpdated:
		return true
	}
	return false
}

type CardCreatedWebhookEvent struct {
	// The token of the card that was created.
	CardToken string `json:"card_token,required" format:"uuid"`
	// The type of event that occurred.
	EventType CardCreatedWebhookEventEventType `json:"event_type,required"`
	// The token of the card that was replaced, if the new card is a replacement card.
	ReplacementFor string                      `json:"replacement_for,nullable" format:"uuid"`
	JSON           cardCreatedWebhookEventJSON `json:"-"`
}

// cardCreatedWebhookEventJSON contains the JSON metadata for the struct
// [CardCreatedWebhookEvent]
type cardCreatedWebhookEventJSON struct {
	CardToken      apijson.Field
	EventType      apijson.Field
	ReplacementFor apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardCreatedWebhookEventEventType string

const (
	CardCreatedWebhookEventEventTypeCardCreated CardCreatedWebhookEventEventType = "card.created"
)

func (r CardCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardCreatedWebhookEventEventTypeCardCreated:
		return true
	}
	return false
}

type CardConvertedWebhookEvent struct {
	// The token of the card that was created.
	CardToken string `json:"card_token,required" format:"uuid"`
	// The type of event that occurred.
	EventType CardConvertedWebhookEventEventType `json:"event_type,required"`
	JSON      cardConvertedWebhookEventJSON      `json:"-"`
}

// cardConvertedWebhookEventJSON contains the JSON metadata for the struct
// [CardConvertedWebhookEvent]
type cardConvertedWebhookEventJSON struct {
	CardToken   apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardConvertedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardConvertedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardConvertedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardConvertedWebhookEventEventType string

const (
	CardConvertedWebhookEventEventTypeCardConverted CardConvertedWebhookEventEventType = "card.converted"
)

func (r CardConvertedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardConvertedWebhookEventEventTypeCardConverted:
		return true
	}
	return false
}

type CardRenewedWebhookEvent struct {
	// The type of event that occurred.
	EventType CardRenewedWebhookEventEventType `json:"event_type,required"`
	// The token of the card that was renewed.
	CardToken string `json:"card_token" format:"uuid"`
	// The new expiration month of the card.
	ExpMonth string `json:"exp_month"`
	// The new expiration year of the card.
	ExpYear string `json:"exp_year"`
	// The previous expiration month of the card.
	PreviousExpMonth string `json:"previous_exp_month"`
	// The previous expiration year of the card.
	PreviousExpYear string                      `json:"previous_exp_year"`
	JSON            cardRenewedWebhookEventJSON `json:"-"`
}

// cardRenewedWebhookEventJSON contains the JSON metadata for the struct
// [CardRenewedWebhookEvent]
type cardRenewedWebhookEventJSON struct {
	EventType        apijson.Field
	CardToken        apijson.Field
	ExpMonth         apijson.Field
	ExpYear          apijson.Field
	PreviousExpMonth apijson.Field
	PreviousExpYear  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CardRenewedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardRenewedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardRenewedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardRenewedWebhookEventEventType string

const (
	CardRenewedWebhookEventEventTypeCardRenewed CardRenewedWebhookEventEventType = "card.renewed"
)

func (r CardRenewedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardRenewedWebhookEventEventTypeCardRenewed:
		return true
	}
	return false
}

type CardReissuedWebhookEvent struct {
	// The type of event that occurred.
	EventType CardReissuedWebhookEventEventType `json:"event_type,required"`
	// The token of the card that was reissued.
	CardToken string                       `json:"card_token" format:"uuid"`
	JSON      cardReissuedWebhookEventJSON `json:"-"`
}

// cardReissuedWebhookEventJSON contains the JSON metadata for the struct
// [CardReissuedWebhookEvent]
type cardReissuedWebhookEventJSON struct {
	EventType   apijson.Field
	CardToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardReissuedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardReissuedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardReissuedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardReissuedWebhookEventEventType string

const (
	CardReissuedWebhookEventEventTypeCardReissued CardReissuedWebhookEventEventType = "card.reissued"
)

func (r CardReissuedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardReissuedWebhookEventEventTypeCardReissued:
		return true
	}
	return false
}

type CardShippedWebhookEvent struct {
	// The token of the bulk order associated with this card shipment, if applicable.
	BulkOrderToken string `json:"bulk_order_token,required,nullable" format:"uuid"`
	// The token of the card that was shipped.
	CardToken string `json:"card_token,required" format:"uuid"`
	// The type of event that occurred.
	EventType CardShippedWebhookEventEventType `json:"event_type,required"`
	// The specific shipping method used to ship the card.
	ShippingMethod CardShippedWebhookEventShippingMethod `json:"shipping_method,required"`
	// The tracking number of the shipment.
	TrackingNumber string                      `json:"tracking_number,required,nullable"`
	JSON           cardShippedWebhookEventJSON `json:"-"`
}

// cardShippedWebhookEventJSON contains the JSON metadata for the struct
// [CardShippedWebhookEvent]
type cardShippedWebhookEventJSON struct {
	BulkOrderToken apijson.Field
	CardToken      apijson.Field
	EventType      apijson.Field
	ShippingMethod apijson.Field
	TrackingNumber apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardShippedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardShippedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardShippedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardShippedWebhookEventEventType string

const (
	CardShippedWebhookEventEventTypeCardShipped CardShippedWebhookEventEventType = "card.shipped"
)

func (r CardShippedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardShippedWebhookEventEventTypeCardShipped:
		return true
	}
	return false
}

// The specific shipping method used to ship the card.
type CardShippedWebhookEventShippingMethod string

const (
	CardShippedWebhookEventShippingMethodExUsExpeditedWithTracking             CardShippedWebhookEventShippingMethod = "Ex-US expedited with tracking"
	CardShippedWebhookEventShippingMethodExUsStandardWithTracking              CardShippedWebhookEventShippingMethod = "Ex-US standard with tracking"
	CardShippedWebhookEventShippingMethodExUsStandardWithoutTracking           CardShippedWebhookEventShippingMethod = "Ex-US standard without tracking"
	CardShippedWebhookEventShippingMethodFedEx2Days                            CardShippedWebhookEventShippingMethod = "FedEx 2 days"
	CardShippedWebhookEventShippingMethodFedExExpress                          CardShippedWebhookEventShippingMethod = "FedEx express"
	CardShippedWebhookEventShippingMethodFedExOvernight                        CardShippedWebhookEventShippingMethod = "FedEx overnight"
	CardShippedWebhookEventShippingMethodUspsPriority                          CardShippedWebhookEventShippingMethod = "USPS priority"
	CardShippedWebhookEventShippingMethodUspsWithTracking                      CardShippedWebhookEventShippingMethod = "USPS with tracking"
	CardShippedWebhookEventShippingMethodUspsWithoutTrackingEnvelope           CardShippedWebhookEventShippingMethod = "USPS without tracking envelope"
	CardShippedWebhookEventShippingMethodUspsWithoutTrackingEnvelopeNonMachine CardShippedWebhookEventShippingMethod = "USPS without tracking envelope non-machine"
	CardShippedWebhookEventShippingMethodUspsWithoutTrackingFlat               CardShippedWebhookEventShippingMethod = "USPS without tracking flat"
)

func (r CardShippedWebhookEventShippingMethod) IsKnown() bool {
	switch r {
	case CardShippedWebhookEventShippingMethodExUsExpeditedWithTracking, CardShippedWebhookEventShippingMethodExUsStandardWithTracking, CardShippedWebhookEventShippingMethodExUsStandardWithoutTracking, CardShippedWebhookEventShippingMethodFedEx2Days, CardShippedWebhookEventShippingMethodFedExExpress, CardShippedWebhookEventShippingMethodFedExOvernight, CardShippedWebhookEventShippingMethodUspsPriority, CardShippedWebhookEventShippingMethodUspsWithTracking, CardShippedWebhookEventShippingMethodUspsWithoutTrackingEnvelope, CardShippedWebhookEventShippingMethodUspsWithoutTrackingEnvelopeNonMachine, CardShippedWebhookEventShippingMethodUspsWithoutTrackingFlat:
		return true
	}
	return false
}

type CardUpdatedWebhookEvent struct {
	// The token of the card that was updated.
	CardToken string `json:"card_token,required" format:"uuid"`
	// The type of event that occurred.
	EventType CardUpdatedWebhookEventEventType `json:"event_type,required"`
	// The previous values of the fields that were updated.
	PreviousFields interface{} `json:"previous_fields,required"`
	// The current state of the card.
	State string                      `json:"state,required"`
	JSON  cardUpdatedWebhookEventJSON `json:"-"`
}

// cardUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [CardUpdatedWebhookEvent]
type cardUpdatedWebhookEventJSON struct {
	CardToken      apijson.Field
	EventType      apijson.Field
	PreviousFields apijson.Field
	State          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardUpdatedWebhookEventEventType string

const (
	CardUpdatedWebhookEventEventTypeCardUpdated CardUpdatedWebhookEventEventType = "card.updated"
)

func (r CardUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardUpdatedWebhookEventEventTypeCardUpdated:
		return true
	}
	return false
}

type CardTransactionUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType CardTransactionUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      cardTransactionUpdatedWebhookEventJSON      `json:"-"`
	Transaction
}

// cardTransactionUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [CardTransactionUpdatedWebhookEvent]
type cardTransactionUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardTransactionUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardTransactionUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardTransactionUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardTransactionUpdatedWebhookEventEventType string

const (
	CardTransactionUpdatedWebhookEventEventTypeCardTransactionUpdated CardTransactionUpdatedWebhookEventEventType = "card_transaction.updated"
)

func (r CardTransactionUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardTransactionUpdatedWebhookEventEventTypeCardTransactionUpdated:
		return true
	}
	return false
}

type CardTransactionEnhancedDataCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType CardTransactionEnhancedDataCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      cardTransactionEnhancedDataCreatedWebhookEventJSON      `json:"-"`
	EnhancedData
}

// cardTransactionEnhancedDataCreatedWebhookEventJSON contains the JSON metadata
// for the struct [CardTransactionEnhancedDataCreatedWebhookEvent]
type cardTransactionEnhancedDataCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardTransactionEnhancedDataCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardTransactionEnhancedDataCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardTransactionEnhancedDataCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardTransactionEnhancedDataCreatedWebhookEventEventType string

const (
	CardTransactionEnhancedDataCreatedWebhookEventEventTypeCardTransactionEnhancedDataCreated CardTransactionEnhancedDataCreatedWebhookEventEventType = "card_transaction.enhanced_data.created"
)

func (r CardTransactionEnhancedDataCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardTransactionEnhancedDataCreatedWebhookEventEventTypeCardTransactionEnhancedDataCreated:
		return true
	}
	return false
}

type CardTransactionEnhancedDataUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType CardTransactionEnhancedDataUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      cardTransactionEnhancedDataUpdatedWebhookEventJSON      `json:"-"`
	EnhancedData
}

// cardTransactionEnhancedDataUpdatedWebhookEventJSON contains the JSON metadata
// for the struct [CardTransactionEnhancedDataUpdatedWebhookEvent]
type cardTransactionEnhancedDataUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardTransactionEnhancedDataUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardTransactionEnhancedDataUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CardTransactionEnhancedDataUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type CardTransactionEnhancedDataUpdatedWebhookEventEventType string

const (
	CardTransactionEnhancedDataUpdatedWebhookEventEventTypeCardTransactionEnhancedDataUpdated CardTransactionEnhancedDataUpdatedWebhookEventEventType = "card_transaction.enhanced_data.updated"
)

func (r CardTransactionEnhancedDataUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case CardTransactionEnhancedDataUpdatedWebhookEventEventTypeCardTransactionEnhancedDataUpdated:
		return true
	}
	return false
}

type DigitalWalletTokenizationApprovalRequestWebhookEvent struct {
	// Unique identifier for the user tokenizing a card
	AccountToken string `json:"account_token,required"`
	// Unique identifier for the card being tokenized
	CardToken string `json:"card_token,required"`
	// Indicate when the request was received from Mastercard or Visa
	Created time.Time `json:"created,required" format:"date-time"`
	// Contains the metadata for the customer tokenization decision.
	CustomerTokenizationDecision DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecision `json:"customer_tokenization_decision,required,nullable"`
	// The name of this event
	EventType DigitalWalletTokenizationApprovalRequestWebhookEventEventType `json:"event_type,required"`
	// Whether Lithic decisioned on the token, and if so, what the decision was.
	// APPROVED/VERIFICATION_REQUIRED/DENIED.
	IssuerDecision DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecision `json:"issuer_decision,required"`
	// The channel through which the tokenization was made.
	TokenizationChannel DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannel `json:"tokenization_channel,required"`
	// Unique identifier for the digital wallet token attempt
	TokenizationToken     string                `json:"tokenization_token,required"`
	WalletDecisioningInfo WalletDecisioningInfo `json:"wallet_decisioning_info,required"`
	Device                Device                `json:"device"`
	// Contains the metadata for the digital wallet being tokenized.
	DigitalWalletTokenMetadata DigitalWalletTokenMetadata `json:"digital_wallet_token_metadata"`
	// Results from rules that were evaluated for this tokenization
	RuleResults []TokenizationRuleResult `json:"rule_results"`
	// List of reasons why the tokenization was declined
	TokenizationDeclineReasons []TokenizationDeclineReason `json:"tokenization_decline_reasons"`
	// The source of the tokenization.
	TokenizationSource DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource `json:"tokenization_source"`
	// List of reasons why two-factor authentication was required
	TokenizationTfaReasons []TokenizationTfaReason                                  `json:"tokenization_tfa_reasons"`
	JSON                   digitalWalletTokenizationApprovalRequestWebhookEventJSON `json:"-"`
}

// digitalWalletTokenizationApprovalRequestWebhookEventJSON contains the JSON
// metadata for the struct [DigitalWalletTokenizationApprovalRequestWebhookEvent]
type digitalWalletTokenizationApprovalRequestWebhookEventJSON struct {
	AccountToken                 apijson.Field
	CardToken                    apijson.Field
	Created                      apijson.Field
	CustomerTokenizationDecision apijson.Field
	EventType                    apijson.Field
	IssuerDecision               apijson.Field
	TokenizationChannel          apijson.Field
	TokenizationToken            apijson.Field
	WalletDecisioningInfo        apijson.Field
	Device                       apijson.Field
	DigitalWalletTokenMetadata   apijson.Field
	RuleResults                  apijson.Field
	TokenizationDeclineReasons   apijson.Field
	TokenizationSource           apijson.Field
	TokenizationTfaReasons       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DigitalWalletTokenizationApprovalRequestWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationApprovalRequestWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DigitalWalletTokenizationApprovalRequestWebhookEvent) implementsParsedWebhookEvent() {}

// Contains the metadata for the customer tokenization decision.
type DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecision struct {
	// The outcome of the customer's decision
	Outcome DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome `json:"outcome,required"`
	// The customer's subscribed URL
	ResponderURL string `json:"responder_url,required"`
	// Time in ms it took for the customer's URL to respond
	Latency string `json:"latency"`
	// The response code that the customer provided
	ResponseCode string                                                                               `json:"response_code"`
	JSON         digitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionJSON `json:"-"`
}

// digitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionJSON
// contains the JSON metadata for the struct
// [DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecision]
type digitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionJSON struct {
	Outcome      apijson.Field
	ResponderURL apijson.Field
	Latency      apijson.Field
	ResponseCode apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionJSON) RawJSON() string {
	return r.raw
}

// The outcome of the customer's decision
type DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome string

const (
	DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeApproved                        DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "APPROVED"
	DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeDeclined                        DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "DECLINED"
	DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeError                           DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "ERROR"
	DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeInvalidResponse                 DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "INVALID_RESPONSE"
	DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeRequireAdditionalAuthentication DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "REQUIRE_ADDITIONAL_AUTHENTICATION"
	DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeTimeout                         DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "TIMEOUT"
)

func (r DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeApproved, DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeDeclined, DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeError, DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeInvalidResponse, DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeRequireAdditionalAuthentication, DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeTimeout:
		return true
	}
	return false
}

// The name of this event
type DigitalWalletTokenizationApprovalRequestWebhookEventEventType string

const (
	DigitalWalletTokenizationApprovalRequestWebhookEventEventTypeDigitalWalletTokenizationApprovalRequest DigitalWalletTokenizationApprovalRequestWebhookEventEventType = "digital_wallet.tokenization_approval_request"
)

func (r DigitalWalletTokenizationApprovalRequestWebhookEventEventType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationApprovalRequestWebhookEventEventTypeDigitalWalletTokenizationApprovalRequest:
		return true
	}
	return false
}

// Whether Lithic decisioned on the token, and if so, what the decision was.
// APPROVED/VERIFICATION_REQUIRED/DENIED.
type DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecision string

const (
	DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecisionApproved             DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecision = "APPROVED"
	DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecisionDenied               DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecision = "DENIED"
	DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecisionVerificationRequired DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecision = "VERIFICATION_REQUIRED"
)

func (r DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecision) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecisionApproved, DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecisionDenied, DigitalWalletTokenizationApprovalRequestWebhookEventIssuerDecisionVerificationRequired:
		return true
	}
	return false
}

// The channel through which the tokenization was made.
type DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannel string

const (
	DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannelDigitalWallet DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannel = "DIGITAL_WALLET"
	DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannelMerchant      DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannel = "MERCHANT"
)

func (r DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannel) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannelDigitalWallet, DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationChannelMerchant:
		return true
	}
	return false
}

// The source of the tokenization.
type DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource string

const (
	DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceAccountOnFile   DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource = "ACCOUNT_ON_FILE"
	DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceContactlessTap  DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource = "CONTACTLESS_TAP"
	DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceManualProvision DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource = "MANUAL_PROVISION"
	DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourcePushProvision   DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource = "PUSH_PROVISION"
	DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceToken           DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource = "TOKEN"
	DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceUnknown         DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource = "UNKNOWN"
)

func (r DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSource) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceAccountOnFile, DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceContactlessTap, DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceManualProvision, DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourcePushProvision, DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceToken, DigitalWalletTokenizationApprovalRequestWebhookEventTokenizationSourceUnknown:
		return true
	}
	return false
}

type DigitalWalletTokenizationResultWebhookEvent struct {
	// Account token
	AccountToken string `json:"account_token,required"`
	// Card token
	CardToken string `json:"card_token,required"`
	// Created date
	Created time.Time `json:"created,required" format:"date-time"`
	// The type of event that occurred.
	EventType DigitalWalletTokenizationResultWebhookEventEventType `json:"event_type,required"`
	// The result of the tokenization request.
	TokenizationResultDetails DigitalWalletTokenizationResultWebhookEventTokenizationResultDetails `json:"tokenization_result_details,required"`
	// Tokenization token
	TokenizationToken string                                          `json:"tokenization_token,required"`
	JSON              digitalWalletTokenizationResultWebhookEventJSON `json:"-"`
}

// digitalWalletTokenizationResultWebhookEventJSON contains the JSON metadata for
// the struct [DigitalWalletTokenizationResultWebhookEvent]
type digitalWalletTokenizationResultWebhookEventJSON struct {
	AccountToken              apijson.Field
	CardToken                 apijson.Field
	Created                   apijson.Field
	EventType                 apijson.Field
	TokenizationResultDetails apijson.Field
	TokenizationToken         apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DigitalWalletTokenizationResultWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationResultWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DigitalWalletTokenizationResultWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type DigitalWalletTokenizationResultWebhookEventEventType string

const (
	DigitalWalletTokenizationResultWebhookEventEventTypeDigitalWalletTokenizationResult DigitalWalletTokenizationResultWebhookEventEventType = "digital_wallet.tokenization_result"
)

func (r DigitalWalletTokenizationResultWebhookEventEventType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationResultWebhookEventEventTypeDigitalWalletTokenizationResult:
		return true
	}
	return false
}

// The result of the tokenization request.
type DigitalWalletTokenizationResultWebhookEventTokenizationResultDetails struct {
	// Lithic's tokenization decision.
	IssuerDecision string `json:"issuer_decision,required"`
	// List of reasons why the tokenization was declined
	TokenizationDeclineReasons []DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason `json:"tokenization_decline_reasons,required"`
	// The customer's tokenization decision if applicable.
	CustomerDecision string `json:"customer_decision,nullable"`
	// Results from rules that were evaluated for this tokenization
	RuleResults []TokenizationRuleResult `json:"rule_results"`
	// An RFC 3339 timestamp indicating when the tokenization succeeded.
	TokenActivatedDateTime time.Time `json:"token_activated_date_time,nullable" format:"date-time"`
	// List of reasons why two-factor authentication was required
	TokenizationTfaReasons []TokenizationTfaReason `json:"tokenization_tfa_reasons"`
	// The wallet's recommended decision.
	WalletDecision string                                                                   `json:"wallet_decision,nullable"`
	JSON           digitalWalletTokenizationResultWebhookEventTokenizationResultDetailsJSON `json:"-"`
}

// digitalWalletTokenizationResultWebhookEventTokenizationResultDetailsJSON
// contains the JSON metadata for the struct
// [DigitalWalletTokenizationResultWebhookEventTokenizationResultDetails]
type digitalWalletTokenizationResultWebhookEventTokenizationResultDetailsJSON struct {
	IssuerDecision             apijson.Field
	TokenizationDeclineReasons apijson.Field
	CustomerDecision           apijson.Field
	RuleResults                apijson.Field
	TokenActivatedDateTime     apijson.Field
	TokenizationTfaReasons     apijson.Field
	WalletDecision             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DigitalWalletTokenizationResultWebhookEventTokenizationResultDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationResultWebhookEventTokenizationResultDetailsJSON) RawJSON() string {
	return r.raw
}

type DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason string

const (
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonAccountScore1                  DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "ACCOUNT_SCORE_1"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonAllWalletDeclineReasonsPresent DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardExpiryMonthMismatch        DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CARD_EXPIRY_MONTH_MISMATCH"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardExpiryYearMismatch         DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CARD_EXPIRY_YEAR_MISMATCH"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardInvalidState               DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CARD_INVALID_STATE"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCustomerRedPath                DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CUSTOMER_RED_PATH"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCvcMismatch                    DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CVC_MISMATCH"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonDeviceScore1                   DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "DEVICE_SCORE_1"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonGenericDecline                 DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "GENERIC_DECLINE"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonInvalidCustomerResponse        DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "INVALID_CUSTOMER_RESPONSE"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonNetworkFailure                 DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "NETWORK_FAILURE"
	DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonWalletRecommendedDecisionRed   DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "WALLET_RECOMMENDED_DECISION_RED"
)

func (r DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonAccountScore1, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonAllWalletDeclineReasonsPresent, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardExpiryMonthMismatch, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardExpiryYearMismatch, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardInvalidState, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCustomerRedPath, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCvcMismatch, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonDeviceScore1, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonGenericDecline, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonInvalidCustomerResponse, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonNetworkFailure, DigitalWalletTokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonWalletRecommendedDecisionRed:
		return true
	}
	return false
}

type DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEvent struct {
	// Unique identifier for the user tokenizing a card
	AccountToken     string                                                                           `json:"account_token,required"`
	ActivationMethod DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod `json:"activation_method,required"`
	// Authentication code to provide to the user tokenizing a card.
	AuthenticationCode string `json:"authentication_code,required"`
	// Unique identifier for the card being tokenized
	CardToken string `json:"card_token,required"`
	// Indicate when the request was received from Mastercard or Visa
	Created time.Time `json:"created,required" format:"date-time"`
	// The type of event that occurred.
	EventType DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventEventType `json:"event_type,required"`
	// Unique identifier for the tokenization
	TokenizationToken string                                                               `json:"tokenization_token,required"`
	JSON              digitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventJSON `json:"-"`
}

// digitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventJSON contains
// the JSON metadata for the struct
// [DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEvent]
type digitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventJSON struct {
	AccountToken       apijson.Field
	ActivationMethod   apijson.Field
	AuthenticationCode apijson.Field
	CardToken          apijson.Field
	Created            apijson.Field
	EventType          apijson.Field
	TokenizationToken  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEvent) implementsParsedWebhookEvent() {
}

type DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod struct {
	// The communication method that the user has selected to use to receive the
	// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
	// = "EMAIL_TO_CARDHOLDER_ADDRESS"
	Type DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType `json:"type,required"`
	// The location where the user wants to receive the authentication code. The format
	// depends on the ActivationMethod.Type field. If Type is Email, the Value will be
	// the email address. If the Type is Sms, the Value will be the phone number.
	Value string                                                                               `json:"value,required"`
	JSON  digitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodJSON `json:"-"`
}

// digitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodJSON
// contains the JSON metadata for the struct
// [DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod]
type digitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodJSON) RawJSON() string {
	return r.raw
}

// The communication method that the user has selected to use to receive the
// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
// = "EMAIL_TO_CARDHOLDER_ADDRESS"
type DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType string

const (
	DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodTypeEmailToCardholderAddress DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType = "EMAIL_TO_CARDHOLDER_ADDRESS"
	DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodTypeTextToCardholderNumber   DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType = "TEXT_TO_CARDHOLDER_NUMBER"
)

func (r DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodTypeEmailToCardholderAddress, DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodTypeTextToCardholderNumber:
		return true
	}
	return false
}

// The type of event that occurred.
type DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventEventType string

const (
	DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventEventType = "digital_wallet.tokenization_two_factor_authentication_code"
)

func (r DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventEventType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode:
		return true
	}
	return false
}

type DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEvent struct {
	// Unique identifier for the user tokenizing a card
	AccountToken     string                                                                               `json:"account_token,required"`
	ActivationMethod DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod `json:"activation_method,required"`
	// Unique identifier for the card being tokenized
	CardToken string `json:"card_token,required"`
	// Indicate when the request was received from Mastercard or Visa
	Created time.Time `json:"created,required" format:"date-time"`
	// The type of event that occurred.
	EventType DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventEventType `json:"event_type,required"`
	// Unique identifier for the tokenization
	TokenizationToken string                                                                   `json:"tokenization_token,required"`
	JSON              digitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventJSON `json:"-"`
}

// digitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventJSON
// contains the JSON metadata for the struct
// [DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEvent]
type digitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventJSON struct {
	AccountToken      apijson.Field
	ActivationMethod  apijson.Field
	CardToken         apijson.Field
	Created           apijson.Field
	EventType         apijson.Field
	TokenizationToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEvent) implementsParsedWebhookEvent() {
}

type DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod struct {
	// The communication method that the user has selected to use to receive the
	// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
	// = "EMAIL_TO_CARDHOLDER_ADDRESS"
	Type DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType `json:"type,required"`
	// The location to which the authentication code was sent. The format depends on
	// the ActivationMethod.Type field. If Type is Email, the Value will be the email
	// address. If the Type is Sms, the Value will be the phone number.
	Value string                                                                                   `json:"value,required"`
	JSON  digitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodJSON `json:"-"`
}

// digitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodJSON
// contains the JSON metadata for the struct
// [DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod]
type digitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodJSON) RawJSON() string {
	return r.raw
}

// The communication method that the user has selected to use to receive the
// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
// = "EMAIL_TO_CARDHOLDER_ADDRESS"
type DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType string

const (
	DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodTypeEmailToCardholderAddress DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType = "EMAIL_TO_CARDHOLDER_ADDRESS"
	DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodTypeTextToCardholderNumber   DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType = "TEXT_TO_CARDHOLDER_NUMBER"
)

func (r DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodTypeEmailToCardholderAddress, DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodTypeTextToCardholderNumber:
		return true
	}
	return false
}

// The type of event that occurred.
type DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventEventType string

const (
	DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
)

func (r DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventEventType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent:
		return true
	}
	return false
}

type DigitalWalletTokenizationUpdatedWebhookEvent struct {
	// Account token
	AccountToken string `json:"account_token,required"`
	// Card token
	CardToken string `json:"card_token,required"`
	// Created date
	Created time.Time `json:"created,required" format:"date-time"`
	// The type of event that occurred.
	EventType    DigitalWalletTokenizationUpdatedWebhookEventEventType `json:"event_type,required"`
	Tokenization Tokenization                                          `json:"tokenization,required"`
	JSON         digitalWalletTokenizationUpdatedWebhookEventJSON      `json:"-"`
}

// digitalWalletTokenizationUpdatedWebhookEventJSON contains the JSON metadata for
// the struct [DigitalWalletTokenizationUpdatedWebhookEvent]
type digitalWalletTokenizationUpdatedWebhookEventJSON struct {
	AccountToken apijson.Field
	CardToken    apijson.Field
	Created      apijson.Field
	EventType    apijson.Field
	Tokenization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DigitalWalletTokenizationUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenizationUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DigitalWalletTokenizationUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type DigitalWalletTokenizationUpdatedWebhookEventEventType string

const (
	DigitalWalletTokenizationUpdatedWebhookEventEventTypeDigitalWalletTokenizationUpdated DigitalWalletTokenizationUpdatedWebhookEventEventType = "digital_wallet.tokenization_updated"
)

func (r DigitalWalletTokenizationUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenizationUpdatedWebhookEventEventTypeDigitalWalletTokenizationUpdated:
		return true
	}
	return false
}

// Dispute.
type DisputeUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType DisputeUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      disputeUpdatedWebhookEventJSON      `json:"-"`
	Dispute
}

// disputeUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [DisputeUpdatedWebhookEvent]
type disputeUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type DisputeUpdatedWebhookEventEventType string

const (
	DisputeUpdatedWebhookEventEventTypeDisputeUpdated DisputeUpdatedWebhookEventEventType = "dispute.updated"
)

func (r DisputeUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case DisputeUpdatedWebhookEventEventTypeDisputeUpdated:
		return true
	}
	return false
}

// Dispute evidence.
type DisputeEvidenceUploadFailedWebhookEvent struct {
	// The type of event that occurred.
	EventType DisputeEvidenceUploadFailedWebhookEventEventType `json:"event_type,required"`
	JSON      disputeEvidenceUploadFailedWebhookEventJSON      `json:"-"`
	DisputeEvidence
}

// disputeEvidenceUploadFailedWebhookEventJSON contains the JSON metadata for the
// struct [DisputeEvidenceUploadFailedWebhookEvent]
type disputeEvidenceUploadFailedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeEvidenceUploadFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeEvidenceUploadFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeEvidenceUploadFailedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type DisputeEvidenceUploadFailedWebhookEventEventType string

const (
	DisputeEvidenceUploadFailedWebhookEventEventTypeDisputeEvidenceUploadFailed DisputeEvidenceUploadFailedWebhookEventEventType = "dispute_evidence.upload_failed"
)

func (r DisputeEvidenceUploadFailedWebhookEventEventType) IsKnown() bool {
	switch r {
	case DisputeEvidenceUploadFailedWebhookEventEventTypeDisputeEvidenceUploadFailed:
		return true
	}
	return false
}

type ExternalBankAccountCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType ExternalBankAccountCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      externalBankAccountCreatedWebhookEventJSON      `json:"-"`
	ExternalBankAccount
}

// externalBankAccountCreatedWebhookEventJSON contains the JSON metadata for the
// struct [ExternalBankAccountCreatedWebhookEvent]
type externalBankAccountCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExternalBankAccountCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ExternalBankAccountCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ExternalBankAccountCreatedWebhookEventEventType string

const (
	ExternalBankAccountCreatedWebhookEventEventTypeExternalBankAccountCreated ExternalBankAccountCreatedWebhookEventEventType = "external_bank_account.created"
)

func (r ExternalBankAccountCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ExternalBankAccountCreatedWebhookEventEventTypeExternalBankAccountCreated:
		return true
	}
	return false
}

type ExternalBankAccountUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType ExternalBankAccountUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      externalBankAccountUpdatedWebhookEventJSON      `json:"-"`
	ExternalBankAccount
}

// externalBankAccountUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [ExternalBankAccountUpdatedWebhookEvent]
type externalBankAccountUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExternalBankAccountUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ExternalBankAccountUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ExternalBankAccountUpdatedWebhookEventEventType string

const (
	ExternalBankAccountUpdatedWebhookEventEventTypeExternalBankAccountUpdated ExternalBankAccountUpdatedWebhookEventEventType = "external_bank_account.updated"
)

func (r ExternalBankAccountUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ExternalBankAccountUpdatedWebhookEventEventTypeExternalBankAccountUpdated:
		return true
	}
	return false
}

type ExternalPaymentCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType ExternalPaymentCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      externalPaymentCreatedWebhookEventJSON      `json:"-"`
	ExternalPayment
}

// externalPaymentCreatedWebhookEventJSON contains the JSON metadata for the struct
// [ExternalPaymentCreatedWebhookEvent]
type externalPaymentCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExternalPaymentCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalPaymentCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ExternalPaymentCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ExternalPaymentCreatedWebhookEventEventType string

const (
	ExternalPaymentCreatedWebhookEventEventTypeExternalPaymentCreated ExternalPaymentCreatedWebhookEventEventType = "external_payment.created"
)

func (r ExternalPaymentCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ExternalPaymentCreatedWebhookEventEventTypeExternalPaymentCreated:
		return true
	}
	return false
}

type ExternalPaymentUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType ExternalPaymentUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      externalPaymentUpdatedWebhookEventJSON      `json:"-"`
	ExternalPayment
}

// externalPaymentUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [ExternalPaymentUpdatedWebhookEvent]
type externalPaymentUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExternalPaymentUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalPaymentUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ExternalPaymentUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ExternalPaymentUpdatedWebhookEventEventType string

const (
	ExternalPaymentUpdatedWebhookEventEventTypeExternalPaymentUpdated ExternalPaymentUpdatedWebhookEventEventType = "external_payment.updated"
)

func (r ExternalPaymentUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ExternalPaymentUpdatedWebhookEventEventTypeExternalPaymentUpdated:
		return true
	}
	return false
}

type FinancialAccountCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType FinancialAccountCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      financialAccountCreatedWebhookEventJSON      `json:"-"`
	FinancialAccount
}

// financialAccountCreatedWebhookEventJSON contains the JSON metadata for the
// struct [FinancialAccountCreatedWebhookEvent]
type financialAccountCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinancialAccountCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r FinancialAccountCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type FinancialAccountCreatedWebhookEventEventType string

const (
	FinancialAccountCreatedWebhookEventEventTypeFinancialAccountCreated FinancialAccountCreatedWebhookEventEventType = "financial_account.created"
)

func (r FinancialAccountCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case FinancialAccountCreatedWebhookEventEventTypeFinancialAccountCreated:
		return true
	}
	return false
}

type FinancialAccountUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType FinancialAccountUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      financialAccountUpdatedWebhookEventJSON      `json:"-"`
	FinancialAccount
}

// financialAccountUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [FinancialAccountUpdatedWebhookEvent]
type financialAccountUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinancialAccountUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r FinancialAccountUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type FinancialAccountUpdatedWebhookEventEventType string

const (
	FinancialAccountUpdatedWebhookEventEventTypeFinancialAccountUpdated FinancialAccountUpdatedWebhookEventEventType = "financial_account.updated"
)

func (r FinancialAccountUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case FinancialAccountUpdatedWebhookEventEventTypeFinancialAccountUpdated:
		return true
	}
	return false
}

type FundingEventCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType FundingEventCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      fundingEventCreatedWebhookEventJSON      `json:"-"`
	FundingEvent
}

// fundingEventCreatedWebhookEventJSON contains the JSON metadata for the struct
// [FundingEventCreatedWebhookEvent]
type fundingEventCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FundingEventCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r FundingEventCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type FundingEventCreatedWebhookEventEventType string

const (
	FundingEventCreatedWebhookEventEventTypeFundingEventCreated FundingEventCreatedWebhookEventEventType = "funding_event.created"
)

func (r FundingEventCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case FundingEventCreatedWebhookEventEventTypeFundingEventCreated:
		return true
	}
	return false
}

type LoanTapeCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType LoanTapeCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      loanTapeCreatedWebhookEventJSON      `json:"-"`
	LoanTape
}

// loanTapeCreatedWebhookEventJSON contains the JSON metadata for the struct
// [LoanTapeCreatedWebhookEvent]
type loanTapeCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapeCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r LoanTapeCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type LoanTapeCreatedWebhookEventEventType string

const (
	LoanTapeCreatedWebhookEventEventTypeLoanTapeCreated LoanTapeCreatedWebhookEventEventType = "loan_tape.created"
)

func (r LoanTapeCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case LoanTapeCreatedWebhookEventEventTypeLoanTapeCreated:
		return true
	}
	return false
}

type LoanTapeUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType LoanTapeUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      loanTapeUpdatedWebhookEventJSON      `json:"-"`
	LoanTape
}

// loanTapeUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [LoanTapeUpdatedWebhookEvent]
type loanTapeUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapeUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r LoanTapeUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type LoanTapeUpdatedWebhookEventEventType string

const (
	LoanTapeUpdatedWebhookEventEventTypeLoanTapeUpdated LoanTapeUpdatedWebhookEventEventType = "loan_tape.updated"
)

func (r LoanTapeUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case LoanTapeUpdatedWebhookEventEventTypeLoanTapeUpdated:
		return true
	}
	return false
}

type ManagementOperationCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType ManagementOperationCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      managementOperationCreatedWebhookEventJSON      `json:"-"`
	ManagementOperationTransaction
}

// managementOperationCreatedWebhookEventJSON contains the JSON metadata for the
// struct [ManagementOperationCreatedWebhookEvent]
type managementOperationCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagementOperationCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managementOperationCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ManagementOperationCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ManagementOperationCreatedWebhookEventEventType string

const (
	ManagementOperationCreatedWebhookEventEventTypeManagementOperationCreated ManagementOperationCreatedWebhookEventEventType = "management_operation.created"
)

func (r ManagementOperationCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ManagementOperationCreatedWebhookEventEventTypeManagementOperationCreated:
		return true
	}
	return false
}

type ManagementOperationUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType ManagementOperationUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      managementOperationUpdatedWebhookEventJSON      `json:"-"`
	ManagementOperationTransaction
}

// managementOperationUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [ManagementOperationUpdatedWebhookEvent]
type managementOperationUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagementOperationUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managementOperationUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ManagementOperationUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ManagementOperationUpdatedWebhookEventEventType string

const (
	ManagementOperationUpdatedWebhookEventEventTypeManagementOperationUpdated ManagementOperationUpdatedWebhookEventEventType = "management_operation.updated"
)

func (r ManagementOperationUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ManagementOperationUpdatedWebhookEventEventTypeManagementOperationUpdated:
		return true
	}
	return false
}

type InternalTransactionCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType InternalTransactionCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      internalTransactionCreatedWebhookEventJSON      `json:"-"`
	InternalTransaction
}

// internalTransactionCreatedWebhookEventJSON contains the JSON metadata for the
// struct [InternalTransactionCreatedWebhookEvent]
type internalTransactionCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InternalTransactionCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r internalTransactionCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InternalTransactionCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type InternalTransactionCreatedWebhookEventEventType string

const (
	InternalTransactionCreatedWebhookEventEventTypeInternalTransactionCreated InternalTransactionCreatedWebhookEventEventType = "internal_transaction.created"
)

func (r InternalTransactionCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case InternalTransactionCreatedWebhookEventEventTypeInternalTransactionCreated:
		return true
	}
	return false
}

type InternalTransactionUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType InternalTransactionUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      internalTransactionUpdatedWebhookEventJSON      `json:"-"`
	InternalTransaction
}

// internalTransactionUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [InternalTransactionUpdatedWebhookEvent]
type internalTransactionUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InternalTransactionUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r internalTransactionUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InternalTransactionUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type InternalTransactionUpdatedWebhookEventEventType string

const (
	InternalTransactionUpdatedWebhookEventEventTypeInternalTransactionUpdated InternalTransactionUpdatedWebhookEventEventType = "internal_transaction.updated"
)

func (r InternalTransactionUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case InternalTransactionUpdatedWebhookEventEventTypeInternalTransactionUpdated:
		return true
	}
	return false
}

type NetworkTotalCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType NetworkTotalCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      networkTotalCreatedWebhookEventJSON      `json:"-"`
	NetworkTotal
}

// networkTotalCreatedWebhookEventJSON contains the JSON metadata for the struct
// [NetworkTotalCreatedWebhookEvent]
type networkTotalCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkTotalCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkTotalCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r NetworkTotalCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type NetworkTotalCreatedWebhookEventEventType string

const (
	NetworkTotalCreatedWebhookEventEventTypeNetworkTotalCreated NetworkTotalCreatedWebhookEventEventType = "network_total.created"
)

func (r NetworkTotalCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case NetworkTotalCreatedWebhookEventEventTypeNetworkTotalCreated:
		return true
	}
	return false
}

type NetworkTotalUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType NetworkTotalUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      networkTotalUpdatedWebhookEventJSON      `json:"-"`
	NetworkTotal
}

// networkTotalUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [NetworkTotalUpdatedWebhookEvent]
type networkTotalUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkTotalUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkTotalUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r NetworkTotalUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type NetworkTotalUpdatedWebhookEventEventType string

const (
	NetworkTotalUpdatedWebhookEventEventTypeNetworkTotalUpdated NetworkTotalUpdatedWebhookEventEventType = "network_total.updated"
)

func (r NetworkTotalUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case NetworkTotalUpdatedWebhookEventEventTypeNetworkTotalUpdated:
		return true
	}
	return false
}

// Payment transaction
type PaymentTransactionCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType PaymentTransactionCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      paymentTransactionCreatedWebhookEventJSON      `json:"-"`
	Payment
}

// paymentTransactionCreatedWebhookEventJSON contains the JSON metadata for the
// struct [PaymentTransactionCreatedWebhookEvent]
type paymentTransactionCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentTransactionCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentTransactionCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PaymentTransactionCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type PaymentTransactionCreatedWebhookEventEventType string

const (
	PaymentTransactionCreatedWebhookEventEventTypePaymentTransactionCreated PaymentTransactionCreatedWebhookEventEventType = "payment_transaction.created"
)

func (r PaymentTransactionCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case PaymentTransactionCreatedWebhookEventEventTypePaymentTransactionCreated:
		return true
	}
	return false
}

// Payment transaction
type PaymentTransactionUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType PaymentTransactionUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      paymentTransactionUpdatedWebhookEventJSON      `json:"-"`
	Payment
}

// paymentTransactionUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [PaymentTransactionUpdatedWebhookEvent]
type paymentTransactionUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentTransactionUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentTransactionUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PaymentTransactionUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type PaymentTransactionUpdatedWebhookEventEventType string

const (
	PaymentTransactionUpdatedWebhookEventEventTypePaymentTransactionUpdated PaymentTransactionUpdatedWebhookEventEventType = "payment_transaction.updated"
)

func (r PaymentTransactionUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case PaymentTransactionUpdatedWebhookEventEventTypePaymentTransactionUpdated:
		return true
	}
	return false
}

type SettlementReportUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType SettlementReportUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      settlementReportUpdatedWebhookEventJSON      `json:"-"`
	SettlementReport
}

// settlementReportUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [SettlementReportUpdatedWebhookEvent]
type settlementReportUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettlementReportUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settlementReportUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SettlementReportUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type SettlementReportUpdatedWebhookEventEventType string

const (
	SettlementReportUpdatedWebhookEventEventTypeSettlementReportUpdated SettlementReportUpdatedWebhookEventEventType = "settlement_report.updated"
)

func (r SettlementReportUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case SettlementReportUpdatedWebhookEventEventTypeSettlementReportUpdated:
		return true
	}
	return false
}

type StatementsCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType StatementsCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      statementsCreatedWebhookEventJSON      `json:"-"`
	Statement
}

// statementsCreatedWebhookEventJSON contains the JSON metadata for the struct
// [StatementsCreatedWebhookEvent]
type statementsCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatementsCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementsCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r StatementsCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type StatementsCreatedWebhookEventEventType string

const (
	StatementsCreatedWebhookEventEventTypeStatementsCreated StatementsCreatedWebhookEventEventType = "statements.created"
)

func (r StatementsCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case StatementsCreatedWebhookEventEventTypeStatementsCreated:
		return true
	}
	return false
}

// Represents a 3DS authentication
type ThreeDSAuthenticationCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType ThreeDSAuthenticationCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      threeDSAuthenticationCreatedWebhookEventJSON      `json:"-"`
	ThreeDSAuthentication
}

// threeDSAuthenticationCreatedWebhookEventJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationCreatedWebhookEvent]
type threeDSAuthenticationCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ThreeDSAuthenticationCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ThreeDSAuthenticationCreatedWebhookEventEventType string

const (
	ThreeDSAuthenticationCreatedWebhookEventEventTypeThreeDSAuthenticationCreated ThreeDSAuthenticationCreatedWebhookEventEventType = "three_ds_authentication.created"
)

func (r ThreeDSAuthenticationCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationCreatedWebhookEventEventTypeThreeDSAuthenticationCreated:
		return true
	}
	return false
}

// Represents a 3DS authentication
type ThreeDSAuthenticationUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType ThreeDSAuthenticationUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      threeDSAuthenticationUpdatedWebhookEventJSON      `json:"-"`
	ThreeDSAuthentication
}

// threeDSAuthenticationUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationUpdatedWebhookEvent]
type threeDSAuthenticationUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ThreeDSAuthenticationUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ThreeDSAuthenticationUpdatedWebhookEventEventType string

const (
	ThreeDSAuthenticationUpdatedWebhookEventEventTypeThreeDSAuthenticationUpdated ThreeDSAuthenticationUpdatedWebhookEventEventType = "three_ds_authentication.updated"
)

func (r ThreeDSAuthenticationUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationUpdatedWebhookEventEventTypeThreeDSAuthenticationUpdated:
		return true
	}
	return false
}

type ThreeDSAuthenticationChallengeWebhookEvent struct {
	// Represents a 3DS authentication
	AuthenticationObject ThreeDSAuthentication `json:"authentication_object,required"`
	// Represents a challenge object for 3DS authentication
	Challenge ThreeDSAuthenticationChallengeWebhookEventChallenge `json:"challenge,required"`
	EventType ThreeDSAuthenticationChallengeWebhookEventEventType `json:"event_type,required"`
	JSON      threeDSAuthenticationChallengeWebhookEventJSON      `json:"-"`
}

// threeDSAuthenticationChallengeWebhookEventJSON contains the JSON metadata for
// the struct [ThreeDSAuthenticationChallengeWebhookEvent]
type threeDSAuthenticationChallengeWebhookEventJSON struct {
	AuthenticationObject apijson.Field
	Challenge            apijson.Field
	EventType            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ThreeDSAuthenticationChallengeWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationChallengeWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ThreeDSAuthenticationChallengeWebhookEvent) implementsParsedWebhookEvent() {}

// Represents a challenge object for 3DS authentication
type ThreeDSAuthenticationChallengeWebhookEventChallenge struct {
	// The type of challenge method issued to the cardholder
	ChallengeMethodType ThreeDSAuthenticationChallengeWebhookEventChallengeChallengeMethodType `json:"challenge_method_type,required"`
	// ISO-8601 time at which the challenge expires
	ExpiryTime time.Time `json:"expiry_time,required" format:"date-time"`
	// ISO-8601 time at which the challenge has started
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Fully qualified app URL of the merchant app. This should be used to redirect the
	// cardholder back to the merchant app after completing an app-based challenge.
	// This URL will only be populated if the 3DS Requestor App is provided to the 3DS
	// SDK.
	AppRequestorURL string                                                  `json:"app_requestor_url,nullable" format:"uri"`
	JSON            threeDSAuthenticationChallengeWebhookEventChallengeJSON `json:"-"`
}

// threeDSAuthenticationChallengeWebhookEventChallengeJSON contains the JSON
// metadata for the struct [ThreeDSAuthenticationChallengeWebhookEventChallenge]
type threeDSAuthenticationChallengeWebhookEventChallengeJSON struct {
	ChallengeMethodType apijson.Field
	ExpiryTime          apijson.Field
	StartTime           apijson.Field
	AppRequestorURL     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ThreeDSAuthenticationChallengeWebhookEventChallenge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationChallengeWebhookEventChallengeJSON) RawJSON() string {
	return r.raw
}

// The type of challenge method issued to the cardholder
type ThreeDSAuthenticationChallengeWebhookEventChallengeChallengeMethodType string

const (
	ThreeDSAuthenticationChallengeWebhookEventChallengeChallengeMethodTypeOutOfBand ThreeDSAuthenticationChallengeWebhookEventChallengeChallengeMethodType = "OUT_OF_BAND"
)

func (r ThreeDSAuthenticationChallengeWebhookEventChallengeChallengeMethodType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationChallengeWebhookEventChallengeChallengeMethodTypeOutOfBand:
		return true
	}
	return false
}

type ThreeDSAuthenticationChallengeWebhookEventEventType string

const (
	ThreeDSAuthenticationChallengeWebhookEventEventTypeThreeDSAuthenticationChallenge ThreeDSAuthenticationChallengeWebhookEventEventType = "three_ds_authentication.challenge"
)

func (r ThreeDSAuthenticationChallengeWebhookEventEventType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationChallengeWebhookEventEventTypeThreeDSAuthenticationChallenge:
		return true
	}
	return false
}

type TokenizationApprovalRequestWebhookEvent struct {
	// Unique identifier for the user tokenizing a card
	AccountToken string `json:"account_token,required"`
	// Unique identifier for the card being tokenized
	CardToken string `json:"card_token,required"`
	// Indicate when the request was received from Mastercard or Visa
	Created time.Time `json:"created,required" format:"date-time"`
	// Contains the metadata for the customer tokenization decision.
	CustomerTokenizationDecision TokenizationApprovalRequestWebhookEventCustomerTokenizationDecision `json:"customer_tokenization_decision,required,nullable"`
	// The name of this event
	EventType TokenizationApprovalRequestWebhookEventEventType `json:"event_type,required"`
	// Whether Lithic decisioned on the token, and if so, what the decision was.
	// APPROVED/VERIFICATION_REQUIRED/DENIED.
	IssuerDecision TokenizationApprovalRequestWebhookEventIssuerDecision `json:"issuer_decision,required"`
	// The channel through which the tokenization was made.
	TokenizationChannel TokenizationApprovalRequestWebhookEventTokenizationChannel `json:"tokenization_channel,required"`
	// Unique identifier for the digital wallet token attempt
	TokenizationToken     string                `json:"tokenization_token,required"`
	WalletDecisioningInfo WalletDecisioningInfo `json:"wallet_decisioning_info,required"`
	Device                Device                `json:"device"`
	// Contains the metadata for the digital wallet being tokenized.
	DigitalWalletTokenMetadata DigitalWalletTokenMetadata `json:"digital_wallet_token_metadata"`
	// Results from rules that were evaluated for this tokenization
	RuleResults []TokenizationRuleResult `json:"rule_results"`
	// List of reasons why the tokenization was declined
	TokenizationDeclineReasons []TokenizationDeclineReason `json:"tokenization_decline_reasons"`
	// The source of the tokenization.
	TokenizationSource TokenizationApprovalRequestWebhookEventTokenizationSource `json:"tokenization_source"`
	// List of reasons why two-factor authentication was required
	TokenizationTfaReasons []TokenizationTfaReason                     `json:"tokenization_tfa_reasons"`
	JSON                   tokenizationApprovalRequestWebhookEventJSON `json:"-"`
}

// tokenizationApprovalRequestWebhookEventJSON contains the JSON metadata for the
// struct [TokenizationApprovalRequestWebhookEvent]
type tokenizationApprovalRequestWebhookEventJSON struct {
	AccountToken                 apijson.Field
	CardToken                    apijson.Field
	Created                      apijson.Field
	CustomerTokenizationDecision apijson.Field
	EventType                    apijson.Field
	IssuerDecision               apijson.Field
	TokenizationChannel          apijson.Field
	TokenizationToken            apijson.Field
	WalletDecisioningInfo        apijson.Field
	Device                       apijson.Field
	DigitalWalletTokenMetadata   apijson.Field
	RuleResults                  apijson.Field
	TokenizationDeclineReasons   apijson.Field
	TokenizationSource           apijson.Field
	TokenizationTfaReasons       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *TokenizationApprovalRequestWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationApprovalRequestWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r TokenizationApprovalRequestWebhookEvent) implementsParsedWebhookEvent() {}

// Contains the metadata for the customer tokenization decision.
type TokenizationApprovalRequestWebhookEventCustomerTokenizationDecision struct {
	// The outcome of the customer's decision
	Outcome TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome `json:"outcome,required"`
	// The customer's subscribed URL
	ResponderURL string `json:"responder_url,required"`
	// Time in ms it took for the customer's URL to respond
	Latency string `json:"latency"`
	// The response code that the customer provided
	ResponseCode string                                                                  `json:"response_code"`
	JSON         tokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionJSON `json:"-"`
}

// tokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionJSON contains
// the JSON metadata for the struct
// [TokenizationApprovalRequestWebhookEventCustomerTokenizationDecision]
type tokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionJSON struct {
	Outcome      apijson.Field
	ResponderURL apijson.Field
	Latency      apijson.Field
	ResponseCode apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenizationApprovalRequestWebhookEventCustomerTokenizationDecision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionJSON) RawJSON() string {
	return r.raw
}

// The outcome of the customer's decision
type TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome string

const (
	TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeApproved                        TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "APPROVED"
	TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeDeclined                        TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "DECLINED"
	TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeError                           TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "ERROR"
	TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeInvalidResponse                 TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "INVALID_RESPONSE"
	TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeRequireAdditionalAuthentication TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "REQUIRE_ADDITIONAL_AUTHENTICATION"
	TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeTimeout                         TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome = "TIMEOUT"
)

func (r TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcome) IsKnown() bool {
	switch r {
	case TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeApproved, TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeDeclined, TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeError, TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeInvalidResponse, TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeRequireAdditionalAuthentication, TokenizationApprovalRequestWebhookEventCustomerTokenizationDecisionOutcomeTimeout:
		return true
	}
	return false
}

// The name of this event
type TokenizationApprovalRequestWebhookEventEventType string

const (
	TokenizationApprovalRequestWebhookEventEventTypeTokenizationApprovalRequest TokenizationApprovalRequestWebhookEventEventType = "tokenization.approval_request"
)

func (r TokenizationApprovalRequestWebhookEventEventType) IsKnown() bool {
	switch r {
	case TokenizationApprovalRequestWebhookEventEventTypeTokenizationApprovalRequest:
		return true
	}
	return false
}

// Whether Lithic decisioned on the token, and if so, what the decision was.
// APPROVED/VERIFICATION_REQUIRED/DENIED.
type TokenizationApprovalRequestWebhookEventIssuerDecision string

const (
	TokenizationApprovalRequestWebhookEventIssuerDecisionApproved             TokenizationApprovalRequestWebhookEventIssuerDecision = "APPROVED"
	TokenizationApprovalRequestWebhookEventIssuerDecisionDenied               TokenizationApprovalRequestWebhookEventIssuerDecision = "DENIED"
	TokenizationApprovalRequestWebhookEventIssuerDecisionVerificationRequired TokenizationApprovalRequestWebhookEventIssuerDecision = "VERIFICATION_REQUIRED"
)

func (r TokenizationApprovalRequestWebhookEventIssuerDecision) IsKnown() bool {
	switch r {
	case TokenizationApprovalRequestWebhookEventIssuerDecisionApproved, TokenizationApprovalRequestWebhookEventIssuerDecisionDenied, TokenizationApprovalRequestWebhookEventIssuerDecisionVerificationRequired:
		return true
	}
	return false
}

// The channel through which the tokenization was made.
type TokenizationApprovalRequestWebhookEventTokenizationChannel string

const (
	TokenizationApprovalRequestWebhookEventTokenizationChannelDigitalWallet TokenizationApprovalRequestWebhookEventTokenizationChannel = "DIGITAL_WALLET"
	TokenizationApprovalRequestWebhookEventTokenizationChannelMerchant      TokenizationApprovalRequestWebhookEventTokenizationChannel = "MERCHANT"
)

func (r TokenizationApprovalRequestWebhookEventTokenizationChannel) IsKnown() bool {
	switch r {
	case TokenizationApprovalRequestWebhookEventTokenizationChannelDigitalWallet, TokenizationApprovalRequestWebhookEventTokenizationChannelMerchant:
		return true
	}
	return false
}

// The source of the tokenization.
type TokenizationApprovalRequestWebhookEventTokenizationSource string

const (
	TokenizationApprovalRequestWebhookEventTokenizationSourceAccountOnFile   TokenizationApprovalRequestWebhookEventTokenizationSource = "ACCOUNT_ON_FILE"
	TokenizationApprovalRequestWebhookEventTokenizationSourceContactlessTap  TokenizationApprovalRequestWebhookEventTokenizationSource = "CONTACTLESS_TAP"
	TokenizationApprovalRequestWebhookEventTokenizationSourceManualProvision TokenizationApprovalRequestWebhookEventTokenizationSource = "MANUAL_PROVISION"
	TokenizationApprovalRequestWebhookEventTokenizationSourcePushProvision   TokenizationApprovalRequestWebhookEventTokenizationSource = "PUSH_PROVISION"
	TokenizationApprovalRequestWebhookEventTokenizationSourceToken           TokenizationApprovalRequestWebhookEventTokenizationSource = "TOKEN"
	TokenizationApprovalRequestWebhookEventTokenizationSourceUnknown         TokenizationApprovalRequestWebhookEventTokenizationSource = "UNKNOWN"
)

func (r TokenizationApprovalRequestWebhookEventTokenizationSource) IsKnown() bool {
	switch r {
	case TokenizationApprovalRequestWebhookEventTokenizationSourceAccountOnFile, TokenizationApprovalRequestWebhookEventTokenizationSourceContactlessTap, TokenizationApprovalRequestWebhookEventTokenizationSourceManualProvision, TokenizationApprovalRequestWebhookEventTokenizationSourcePushProvision, TokenizationApprovalRequestWebhookEventTokenizationSourceToken, TokenizationApprovalRequestWebhookEventTokenizationSourceUnknown:
		return true
	}
	return false
}

type TokenizationResultWebhookEvent struct {
	// Account token
	AccountToken string `json:"account_token,required"`
	// Card token
	CardToken string `json:"card_token,required"`
	// Created date
	Created time.Time `json:"created,required" format:"date-time"`
	// The type of event that occurred.
	EventType TokenizationResultWebhookEventEventType `json:"event_type,required"`
	// The result of the tokenization request.
	TokenizationResultDetails TokenizationResultWebhookEventTokenizationResultDetails `json:"tokenization_result_details,required"`
	// Tokenization token
	TokenizationToken string                             `json:"tokenization_token,required"`
	JSON              tokenizationResultWebhookEventJSON `json:"-"`
}

// tokenizationResultWebhookEventJSON contains the JSON metadata for the struct
// [TokenizationResultWebhookEvent]
type tokenizationResultWebhookEventJSON struct {
	AccountToken              apijson.Field
	CardToken                 apijson.Field
	Created                   apijson.Field
	EventType                 apijson.Field
	TokenizationResultDetails apijson.Field
	TokenizationToken         apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *TokenizationResultWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationResultWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r TokenizationResultWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type TokenizationResultWebhookEventEventType string

const (
	TokenizationResultWebhookEventEventTypeTokenizationResult TokenizationResultWebhookEventEventType = "tokenization.result"
)

func (r TokenizationResultWebhookEventEventType) IsKnown() bool {
	switch r {
	case TokenizationResultWebhookEventEventTypeTokenizationResult:
		return true
	}
	return false
}

// The result of the tokenization request.
type TokenizationResultWebhookEventTokenizationResultDetails struct {
	// Lithic's tokenization decision.
	IssuerDecision string `json:"issuer_decision,required"`
	// List of reasons why the tokenization was declined
	TokenizationDeclineReasons []TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason `json:"tokenization_decline_reasons,required"`
	// The customer's tokenization decision if applicable.
	CustomerDecision string `json:"customer_decision,nullable"`
	// Results from rules that were evaluated for this tokenization
	RuleResults []TokenizationRuleResult `json:"rule_results"`
	// An RFC 3339 timestamp indicating when the tokenization succeeded.
	TokenActivatedDateTime time.Time `json:"token_activated_date_time,nullable" format:"date-time"`
	// List of reasons why two-factor authentication was required
	TokenizationTfaReasons []TokenizationTfaReason `json:"tokenization_tfa_reasons"`
	// The wallet's recommended decision.
	WalletDecision string                                                      `json:"wallet_decision,nullable"`
	JSON           tokenizationResultWebhookEventTokenizationResultDetailsJSON `json:"-"`
}

// tokenizationResultWebhookEventTokenizationResultDetailsJSON contains the JSON
// metadata for the struct
// [TokenizationResultWebhookEventTokenizationResultDetails]
type tokenizationResultWebhookEventTokenizationResultDetailsJSON struct {
	IssuerDecision             apijson.Field
	TokenizationDeclineReasons apijson.Field
	CustomerDecision           apijson.Field
	RuleResults                apijson.Field
	TokenActivatedDateTime     apijson.Field
	TokenizationTfaReasons     apijson.Field
	WalletDecision             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *TokenizationResultWebhookEventTokenizationResultDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationResultWebhookEventTokenizationResultDetailsJSON) RawJSON() string {
	return r.raw
}

type TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason string

const (
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonAccountScore1                  TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "ACCOUNT_SCORE_1"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonAllWalletDeclineReasonsPresent TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardExpiryMonthMismatch        TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CARD_EXPIRY_MONTH_MISMATCH"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardExpiryYearMismatch         TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CARD_EXPIRY_YEAR_MISMATCH"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardInvalidState               TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CARD_INVALID_STATE"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCustomerRedPath                TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CUSTOMER_RED_PATH"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCvcMismatch                    TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "CVC_MISMATCH"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonDeviceScore1                   TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "DEVICE_SCORE_1"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonGenericDecline                 TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "GENERIC_DECLINE"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonInvalidCustomerResponse        TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "INVALID_CUSTOMER_RESPONSE"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonNetworkFailure                 TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "NETWORK_FAILURE"
	TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonWalletRecommendedDecisionRed   TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason = "WALLET_RECOMMENDED_DECISION_RED"
)

func (r TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReason) IsKnown() bool {
	switch r {
	case TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonAccountScore1, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonAllWalletDeclineReasonsPresent, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardExpiryMonthMismatch, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardExpiryYearMismatch, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCardInvalidState, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCustomerRedPath, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonCvcMismatch, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonDeviceScore1, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonGenericDecline, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonInvalidCustomerResponse, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonNetworkFailure, TokenizationResultWebhookEventTokenizationResultDetailsTokenizationDeclineReasonWalletRecommendedDecisionRed:
		return true
	}
	return false
}

type TokenizationTwoFactorAuthenticationCodeWebhookEvent struct {
	// Unique identifier for the user tokenizing a card
	AccountToken     string                                                              `json:"account_token,required"`
	ActivationMethod TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod `json:"activation_method,required"`
	// Authentication code to provide to the user tokenizing a card.
	AuthenticationCode string `json:"authentication_code,required"`
	// Unique identifier for the card being tokenized
	CardToken string `json:"card_token,required"`
	// Indicate when the request was received from Mastercard or Visa
	Created time.Time `json:"created,required" format:"date-time"`
	// The type of event that occurred.
	EventType TokenizationTwoFactorAuthenticationCodeWebhookEventEventType `json:"event_type,required"`
	// Unique identifier for the tokenization
	TokenizationToken string                                                  `json:"tokenization_token,required"`
	JSON              tokenizationTwoFactorAuthenticationCodeWebhookEventJSON `json:"-"`
}

// tokenizationTwoFactorAuthenticationCodeWebhookEventJSON contains the JSON
// metadata for the struct [TokenizationTwoFactorAuthenticationCodeWebhookEvent]
type tokenizationTwoFactorAuthenticationCodeWebhookEventJSON struct {
	AccountToken       apijson.Field
	ActivationMethod   apijson.Field
	AuthenticationCode apijson.Field
	CardToken          apijson.Field
	Created            apijson.Field
	EventType          apijson.Field
	TokenizationToken  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TokenizationTwoFactorAuthenticationCodeWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationTwoFactorAuthenticationCodeWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r TokenizationTwoFactorAuthenticationCodeWebhookEvent) implementsParsedWebhookEvent() {}

type TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod struct {
	// The communication method that the user has selected to use to receive the
	// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
	// = "EMAIL_TO_CARDHOLDER_ADDRESS"
	Type TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType `json:"type,required"`
	// The location where the user wants to receive the authentication code. The format
	// depends on the ActivationMethod.Type field. If Type is Email, the Value will be
	// the email address. If the Type is Sms, the Value will be the phone number.
	Value string                                                                  `json:"value,required"`
	JSON  tokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodJSON `json:"-"`
}

// tokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodJSON contains
// the JSON metadata for the struct
// [TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod]
type tokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodJSON) RawJSON() string {
	return r.raw
}

// The communication method that the user has selected to use to receive the
// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
// = "EMAIL_TO_CARDHOLDER_ADDRESS"
type TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType string

const (
	TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodTypeEmailToCardholderAddress TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType = "EMAIL_TO_CARDHOLDER_ADDRESS"
	TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodTypeTextToCardholderNumber   TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType = "TEXT_TO_CARDHOLDER_NUMBER"
)

func (r TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodType) IsKnown() bool {
	switch r {
	case TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodTypeEmailToCardholderAddress, TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethodTypeTextToCardholderNumber:
		return true
	}
	return false
}

// The type of event that occurred.
type TokenizationTwoFactorAuthenticationCodeWebhookEventEventType string

const (
	TokenizationTwoFactorAuthenticationCodeWebhookEventEventTypeTokenizationTwoFactorAuthenticationCode TokenizationTwoFactorAuthenticationCodeWebhookEventEventType = "tokenization.two_factor_authentication_code"
)

func (r TokenizationTwoFactorAuthenticationCodeWebhookEventEventType) IsKnown() bool {
	switch r {
	case TokenizationTwoFactorAuthenticationCodeWebhookEventEventTypeTokenizationTwoFactorAuthenticationCode:
		return true
	}
	return false
}

type TokenizationTwoFactorAuthenticationCodeSentWebhookEvent struct {
	// Unique identifier for the user tokenizing a card
	AccountToken     string                                                                  `json:"account_token,required"`
	ActivationMethod TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod `json:"activation_method,required"`
	// Unique identifier for the card being tokenized
	CardToken string `json:"card_token,required"`
	// Indicate when the request was received from Mastercard or Visa
	Created time.Time `json:"created,required" format:"date-time"`
	// The type of event that occurred.
	EventType TokenizationTwoFactorAuthenticationCodeSentWebhookEventEventType `json:"event_type,required"`
	// Unique identifier for the tokenization
	TokenizationToken string                                                      `json:"tokenization_token,required"`
	JSON              tokenizationTwoFactorAuthenticationCodeSentWebhookEventJSON `json:"-"`
}

// tokenizationTwoFactorAuthenticationCodeSentWebhookEventJSON contains the JSON
// metadata for the struct
// [TokenizationTwoFactorAuthenticationCodeSentWebhookEvent]
type tokenizationTwoFactorAuthenticationCodeSentWebhookEventJSON struct {
	AccountToken      apijson.Field
	ActivationMethod  apijson.Field
	CardToken         apijson.Field
	Created           apijson.Field
	EventType         apijson.Field
	TokenizationToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenizationTwoFactorAuthenticationCodeSentWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationTwoFactorAuthenticationCodeSentWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r TokenizationTwoFactorAuthenticationCodeSentWebhookEvent) implementsParsedWebhookEvent() {}

type TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod struct {
	// The communication method that the user has selected to use to receive the
	// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
	// = "EMAIL_TO_CARDHOLDER_ADDRESS"
	Type TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType `json:"type,required"`
	// The location to which the authentication code was sent. The format depends on
	// the ActivationMethod.Type field. If Type is Email, the Value will be the email
	// address. If the Type is Sms, the Value will be the phone number.
	Value string                                                                      `json:"value,required"`
	JSON  tokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodJSON `json:"-"`
}

// tokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodJSON
// contains the JSON metadata for the struct
// [TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod]
type tokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodJSON) RawJSON() string {
	return r.raw
}

// The communication method that the user has selected to use to receive the
// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
// = "EMAIL_TO_CARDHOLDER_ADDRESS"
type TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType string

const (
	TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodTypeEmailToCardholderAddress TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType = "EMAIL_TO_CARDHOLDER_ADDRESS"
	TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodTypeTextToCardholderNumber   TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType = "TEXT_TO_CARDHOLDER_NUMBER"
)

func (r TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodType) IsKnown() bool {
	switch r {
	case TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodTypeEmailToCardholderAddress, TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethodTypeTextToCardholderNumber:
		return true
	}
	return false
}

// The type of event that occurred.
type TokenizationTwoFactorAuthenticationCodeSentWebhookEventEventType string

const (
	TokenizationTwoFactorAuthenticationCodeSentWebhookEventEventTypeTokenizationTwoFactorAuthenticationCodeSent TokenizationTwoFactorAuthenticationCodeSentWebhookEventEventType = "tokenization.two_factor_authentication_code_sent"
)

func (r TokenizationTwoFactorAuthenticationCodeSentWebhookEventEventType) IsKnown() bool {
	switch r {
	case TokenizationTwoFactorAuthenticationCodeSentWebhookEventEventTypeTokenizationTwoFactorAuthenticationCodeSent:
		return true
	}
	return false
}

type TokenizationUpdatedWebhookEvent struct {
	// Account token
	AccountToken string `json:"account_token,required"`
	// Card token
	CardToken string `json:"card_token,required"`
	// Created date
	Created time.Time `json:"created,required" format:"date-time"`
	// The type of event that occurred.
	EventType    TokenizationUpdatedWebhookEventEventType `json:"event_type,required"`
	Tokenization Tokenization                             `json:"tokenization,required"`
	JSON         tokenizationUpdatedWebhookEventJSON      `json:"-"`
}

// tokenizationUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [TokenizationUpdatedWebhookEvent]
type tokenizationUpdatedWebhookEventJSON struct {
	AccountToken apijson.Field
	CardToken    apijson.Field
	Created      apijson.Field
	EventType    apijson.Field
	Tokenization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenizationUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r TokenizationUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type TokenizationUpdatedWebhookEventEventType string

const (
	TokenizationUpdatedWebhookEventEventTypeTokenizationUpdated TokenizationUpdatedWebhookEventEventType = "tokenization.updated"
)

func (r TokenizationUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case TokenizationUpdatedWebhookEventEventTypeTokenizationUpdated:
		return true
	}
	return false
}

// Represents a 3DS authentication
type ThreeDSAuthenticationApprovalRequestWebhookEvent struct {
	EventType ThreeDSAuthenticationApprovalRequestWebhookEventEventType `json:"event_type,required"`
	JSON      threeDSAuthenticationApprovalRequestWebhookEventJSON      `json:"-"`
	ThreeDSAuthentication
}

// threeDSAuthenticationApprovalRequestWebhookEventJSON contains the JSON metadata
// for the struct [ThreeDSAuthenticationApprovalRequestWebhookEvent]
type threeDSAuthenticationApprovalRequestWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationApprovalRequestWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationApprovalRequestWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ThreeDSAuthenticationApprovalRequestWebhookEvent) implementsParsedWebhookEvent() {}

type ThreeDSAuthenticationApprovalRequestWebhookEventEventType string

const (
	ThreeDSAuthenticationApprovalRequestWebhookEventEventTypeThreeDSAuthenticationApprovalRequest ThreeDSAuthenticationApprovalRequestWebhookEventEventType = "three_ds_authentication.approval_request"
)

func (r ThreeDSAuthenticationApprovalRequestWebhookEventEventType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationApprovalRequestWebhookEventEventTypeThreeDSAuthenticationApprovalRequest:
		return true
	}
	return false
}

// The Dispute object tracks the progression of a dispute throughout its lifecycle.
type DisputeTransactionCreatedWebhookEvent struct {
	// The type of event that occurred.
	EventType DisputeTransactionCreatedWebhookEventEventType `json:"event_type,required"`
	JSON      disputeTransactionCreatedWebhookEventJSON      `json:"-"`
	DisputeV2
}

// disputeTransactionCreatedWebhookEventJSON contains the JSON metadata for the
// struct [DisputeTransactionCreatedWebhookEvent]
type disputeTransactionCreatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeTransactionCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeTransactionCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeTransactionCreatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type DisputeTransactionCreatedWebhookEventEventType string

const (
	DisputeTransactionCreatedWebhookEventEventTypeDisputeTransactionCreated DisputeTransactionCreatedWebhookEventEventType = "dispute_transaction.created"
)

func (r DisputeTransactionCreatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case DisputeTransactionCreatedWebhookEventEventTypeDisputeTransactionCreated:
		return true
	}
	return false
}

// The Dispute object tracks the progression of a dispute throughout its lifecycle.
type DisputeTransactionUpdatedWebhookEvent struct {
	// The type of event that occurred.
	EventType DisputeTransactionUpdatedWebhookEventEventType `json:"event_type,required"`
	JSON      disputeTransactionUpdatedWebhookEventJSON      `json:"-"`
	DisputeV2
}

// disputeTransactionUpdatedWebhookEventJSON contains the JSON metadata for the
// struct [DisputeTransactionUpdatedWebhookEvent]
type disputeTransactionUpdatedWebhookEventJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeTransactionUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeTransactionUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeTransactionUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type DisputeTransactionUpdatedWebhookEventEventType string

const (
	DisputeTransactionUpdatedWebhookEventEventTypeDisputeTransactionUpdated DisputeTransactionUpdatedWebhookEventEventType = "dispute_transaction.updated"
)

func (r DisputeTransactionUpdatedWebhookEventEventType) IsKnown() bool {
	switch r {
	case DisputeTransactionUpdatedWebhookEventEventTypeDisputeTransactionUpdated:
		return true
	}
	return false
}

// KYB payload for an updated account holder.
type ParsedWebhookEvent struct {
	// The token of the account_holder that was created.
	Token string `json:"token" format:"uuid"`
	// The token of the account_holder that the document belongs to
	AccountHolderToken string `json:"account_holder_token" format:"uuid"`
	AccountNumber      string `json:"account_number,nullable"`
	// This field can have the runtime type of [LoanTapeAccountStanding],
	// [StatementAccountStanding].
	AccountStanding interface{} `json:"account_standing"`
	// The token of the account that was created.
	AccountToken string                        `json:"account_token,nullable" format:"uuid"`
	AccountType  ParsedWebhookEventAccountType `json:"account_type,nullable"`
	// Fee (in cents) assessed by the merchant and paid for by the cardholder. Will be
	// zero if no fee is assessed. Rebates may be transmitted as a negative value to
	// indicate credited fees.
	AcquirerFee int64 `json:"acquirer_fee,nullable"`
	// Unique identifier assigned to a transaction by the acquirer that can be used in
	// dispute and chargeback filing. This field has been deprecated in favor of the
	// `acquirer_reference_number` that resides in the event-level `network_info`.
	//
	// Deprecated: deprecated
	AcquirerReferenceNumber string `json:"acquirer_reference_number,nullable"`
	// This field can have the runtime type of
	// [DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod],
	// [DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod],
	// [TokenizationTwoFactorAuthenticationCodeWebhookEventActivationMethod],
	// [TokenizationTwoFactorAuthenticationCodeSentWebhookEventActivationMethod].
	ActivationMethod interface{} `json:"activation_method"`
	// This field can have the runtime type of [ThreeDSAuthenticationAdditionalData].
	AdditionalData interface{} `json:"additional_data"`
	// Address
	Address ExternalBankAccountAddress `json:"address,nullable"`
	// Deprecated, use `amounts`. Authorization amount of the transaction (in cents),
	// including any acquirer fees. The contents of this field are identical to
	// `authorization_amount`.
	//
	// Deprecated: deprecated
	Amount int64 `json:"amount"`
	// This field can have the runtime type of [StatementAmountDue].
	AmountDue interface{} `json:"amount_due"`
	// This field can have the runtime type of
	// [CardAuthorizationApprovalRequestWebhookEventAmounts], [TransactionAmounts],
	// [NetworkTotalAmounts].
	Amounts interface{} `json:"amounts"`
	// This field can have the runtime type of [ThreeDSAuthenticationApp].
	App interface{} `json:"app"`
	// Date dispute entered arbitration.
	ArbitrationDate time.Time `json:"arbitration_date,nullable" format:"date-time"`
	// Authentication code to provide to the user tokenizing a card.
	AuthenticationCode string `json:"authentication_code"`
	// Represents a 3DS authentication
	AuthenticationObject ThreeDSAuthentication `json:"authentication_object"`
	// Type of authentication request - i.e., the type of transaction or interaction is
	// causing the merchant to request an authentication. Maps to EMV 3DS field
	// `threeDSRequestorAuthenticationInd`.
	AuthenticationRequestType ParsedWebhookEventAuthenticationRequestType `json:"authentication_request_type,nullable"`
	// Indicates the outcome of the 3DS authentication process.
	AuthenticationResult ParsedWebhookEventAuthenticationResult `json:"authentication_result"`
	// Deprecated, use `amounts`. The base transaction amount (in cents) plus the
	// acquirer fee field. This is the amount the issuer should authorize against
	// unless the issuer is paying the acquirer fee on behalf of the cardholder.
	//
	// Deprecated: deprecated
	AuthorizationAmount int64 `json:"authorization_amount,nullable"`
	// A fixed-width 6-digit numeric identifier that can be used to identify a
	// transaction with networks.
	AuthorizationCode string `json:"authorization_code,nullable"`
	// Amount of credit available to spend in cents
	AvailableCredit int64 `json:"available_credit"`
	// This field can have the runtime type of
	// [CardAuthorizationApprovalRequestWebhookEventAvs], [TransactionAvs].
	Avs interface{} `json:"avs"`
	// Auth Rule Backtest Token
	BacktestToken string `json:"backtest_token" format:"uuid"`
	// This field can have the runtime type of [LoanTapeBalances].
	Balances interface{} `json:"balances"`
	// This field can have the runtime type of [ThreeDSAuthenticationBrowser].
	Browser interface{} `json:"browser"`
	// The token of the bulk order associated with this card shipment, if applicable.
	BulkOrderToken string `json:"bulk_order_token,nullable" format:"uuid"`
	// If applicable, represents the business account token associated with the
	// account_holder.
	BusinessAccountToken string `json:"business_account_token,nullable" format:"uuid"`
	// This field can have the runtime type of
	// [CardAuthorizationApprovalRequestWebhookEventCard].
	Card interface{} `json:"card"`
	// Indicates whether the expiration date provided by the cardholder during checkout
	// matches Lithic's record of the card's expiration date.
	CardExpiryCheck ParsedWebhookEventCardExpiryCheck `json:"card_expiry_check"`
	// Unique identifier for the card being tokenized
	CardToken string `json:"card_token"`
	// This field can have the runtime type of [ThreeDSAuthenticationCardholder].
	Cardholder               interface{}              `json:"cardholder"`
	CardholderAuthentication CardholderAuthentication `json:"cardholder_authentication,nullable"`
	// Deprecated, use `amounts`. 3-character alphabetic ISO 4217 code for cardholder's
	// billing currency.
	//
	// Deprecated: deprecated
	CardholderCurrency string `json:"cardholder_currency"`
	// Identifier assigned by the network for this dispute.
	CaseID string `json:"case_id,nullable"`
	// The portion of the transaction requested as cash back by the cardholder, and
	// does not include any acquirer fees. The amount field includes the purchase
	// amount, the requested cash back amount, and any acquirer fees.
	//
	// If no cash back was requested, the value of this field will be 0, and the field
	// will always be present.
	CashAmount int64 `json:"cash_amount"`
	// Deprecated, use `cash_amount`.
	Cashback int64                      `json:"cashback"`
	Category ParsedWebhookEventCategory `json:"category"`
	// This field can have the runtime type of
	// [ThreeDSAuthenticationChallengeWebhookEventChallenge].
	Challenge interface{} `json:"challenge"`
	// This field can have the runtime type of
	// [ThreeDSAuthenticationChallengeMetadata].
	ChallengeMetadata interface{} `json:"challenge_metadata"`
	// Entity that orchestrates the challenge. This won't be set for authentications
	// for which a decision has not yet been made (e.g. in-flight customer decisioning
	// request).
	ChallengeOrchestratedBy ParsedWebhookEventChallengeOrchestratedBy `json:"challenge_orchestrated_by,nullable"`
	// Channel in which the authentication occurs. Maps to EMV 3DS field
	// `deviceChannel`.
	Channel ParsedWebhookEventChannel `json:"channel"`
	// Collection resource type
	CollectionResourceType ParsedWebhookEventCollectionResourceType `json:"collection_resource_type"`
	// This field can have the runtime type of [[]string].
	CollectionTokens interface{} `json:"collection_tokens"`
	// This field can have the runtime type of [EnhancedDataCommon].
	Common interface{} `json:"common"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID string `json:"company_id,nullable"`
	// Deprecated, use `amounts`. If the transaction was requested in a currency other
	// than the settlement currency, this field will be populated to indicate the rate
	// used to translate the merchant_amount to the amount (i.e., `merchant_amount` x
	// `conversion_rate` = `amount`). Note that the `merchant_amount` is in the local
	// currency and the amount is in the settlement currency.
	//
	// Deprecated: deprecated
	ConversionRate float64 `json:"conversion_rate"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept USA bank accounts e.g., USA
	Country string `json:"country"`
	// When the account_holder was created
	Created time.Time `json:"created" format:"date-time"`
	// This field can have the runtime type of [FinancialAccountCreditConfiguration].
	CreditConfiguration interface{} `json:"credit_configuration"`
	// For prepay accounts, this is the minimum prepay balance that must be maintained.
	// For charge card accounts, this is the maximum credit balance extended by a
	// lender
	CreditLimit int64 `json:"credit_limit"`
	// Globally unique identifier for a credit product
	CreditProductToken string `json:"credit_product_token"`
	// 3-character alphabetic ISO 4217 code for the settling currency of the
	// transaction
	Currency string `json:"currency"`
	// Date that the dispute was filed by the customer making the dispute.
	CustomerFiledDate time.Time `json:"customer_filed_date,nullable" format:"date-time"`
	// End customer description of the reason for the dispute.
	CustomerNote string `json:"customer_note,nullable"`
	// This field can have the runtime type of
	// [DigitalWalletTokenizationApprovalRequestWebhookEventCustomerTokenizationDecision],
	// [TokenizationApprovalRequestWebhookEventCustomerTokenizationDecision].
	CustomerTokenizationDecision interface{} `json:"customer_tokenization_decision"`
	// The clearing cycle that the network total record applies to. Mastercard only.
	Cycle int64 `json:"cycle"`
	// This field can have the runtime type of [[]FinancialAccountBalance].
	Data interface{} `json:"data"`
	// Date of transactions that this loan tape covers
	Date      time.Time       `json:"date" format:"date"`
	DayTotals StatementTotals `json:"day_totals"`
	// Number of days in the billing cycle
	DaysInBillingCycle int64 `json:"days_in_billing_cycle"`
	// Entity that made the authentication decision. This won't be set for
	// authentications for which a decision has not yet been made (e.g. in-flight
	// customer decisioning request).
	DecisionMadeBy ParsedWebhookEventDecisionMadeBy `json:"decision_made_by,nullable"`
	Descriptor     string                           `json:"descriptor"`
	// This field can have the runtime type of [[]SettlementSummaryDetails].
	Details interface{} `json:"details"`
	Device  Device      `json:"device"`
	// Contains the metadata for the digital wallet being tokenized.
	DigitalWalletTokenMetadata DigitalWalletTokenMetadata  `json:"digital_wallet_token_metadata"`
	Direction                  ParsedWebhookEventDirection `json:"direction"`
	// Dispute resolution outcome
	Disposition ParsedWebhookEventDisposition `json:"disposition,nullable"`
	// Dispute token evidence is attached to.
	DisputeToken string `json:"dispute_token" format:"uuid"`
	// The total gross amount of disputes settlements. (This field is deprecated and
	// will be removed in a future version of the API. To compute total amounts, Lithic
	// recommends that customers sum the relevant settlement amounts found within
	// `details`.)
	//
	// Deprecated: deprecated
	DisputesGrossAmount int64 `json:"disputes_gross_amount"`
	// Date of Birth of the Individual that owns the external bank account
	Dob time.Time `json:"dob,nullable" format:"date"`
	// Type of documentation to be submitted for verification of an account holder
	DocumentType ParsedWebhookEventDocumentType `json:"document_type"`
	// Doing Business As
	DoingBusinessAs string `json:"doing_business_as,nullable"`
	// URL to download evidence. Only shown when `upload_status` is `UPLOADED`.
	DownloadURL string `json:"download_url"`
	// If updated, the newly updated email associated with the account_holder otherwise
	// the existing email is provided.
	Email string `json:"email"`
	// Balance at the end of the day
	EndingBalance int64 `json:"ending_balance"`
	// The token of the entity that the document belongs to
	EntityToken string `json:"entity_token" format:"uuid"`
	// The event token associated with the authorization. This field is only set for
	// programs enrolled into the beta.
	EventToken string `json:"event_token" format:"uuid"`
	// The type of event that occurred.
	EventType ParsedWebhookEventEventType `json:"event_type"`
	// This field can have the runtime type of [[]BookTransferResponseEvent],
	// [[]TransactionEvent], [[]ExternalPaymentEvent],
	// [[]ManagementOperationTransactionEvent], [[]InternalTransactionEvent],
	// [[]PaymentEvent], [[]DisputeV2Event].
	Events interface{} `json:"events"`
	// Excess credits in the form of provisional credits, payments, or purchase
	// refunds. If positive, the account is in net credit state with no outstanding
	// balances. An overpayment could land an account in this state
	ExcessCredits int64 `json:"excess_credits"`
	// The new expiration month of the card.
	ExpMonth string `json:"exp_month"`
	// The new expiration year of the card.
	ExpYear string `json:"exp_year"`
	// Expected release date for the transaction
	ExpectedReleaseDate time.Time `json:"expected_release_date,nullable" format:"date"`
	// External bank account token
	ExternalBankAccountToken string `json:"external_bank_account_token,nullable" format:"uuid"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID string `json:"external_id,nullable"`
	// External resource associated with the management operation
	ExternalResource ExternalResource `json:"external_resource,nullable"`
	// TRANSFER - Book Transfer Transaction
	Family ParsedWebhookEventFamily `json:"family"`
	// File name of evidence. Recommended to give the dispute evidence a human-readable
	// identifier.
	Filename              string `json:"filename"`
	FinancialAccountToken string `json:"financial_account_token,nullable" format:"uuid"`
	// If applicable, represents the account_holder's first name.
	FirstName string `json:"first_name"`
	// This field can have the runtime type of [[]EnhancedDataFleet].
	Fleet interface{} `json:"fleet"`
	// This field can have the runtime type of
	// [CardAuthorizationApprovalRequestWebhookEventFleetInfo].
	FleetInfo interface{} `json:"fleet_info"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case
	FromFinancialAccountToken string `json:"from_financial_account_token" format:"uuid"`
	// Time of the high watermark
	HighWatermark time.Time `json:"high_watermark" format:"date-time"`
	// The institution that activity occurred on. For Mastercard: ICA (Interbank Card
	// Association). For Maestro: institution ID. For Visa: lowest level SRE
	// (Settlement Reporting Entity).
	InstitutionID string `json:"institution_id"`
	// The total amount of interchange. (This field is deprecated and will be removed
	// in a future version of the API. To compute total amounts, Lithic recommends that
	// customers sum the relevant settlement amounts found within `details`.)
	//
	// Deprecated: deprecated
	InterchangeGrossAmount int64 `json:"interchange_gross_amount"`
	// This field can have the runtime type of [LoanTapeInterestDetails],
	// [StatementInterestDetails].
	InterestDetails interface{} `json:"interest_details"`
	// Indicates that all settlement records related to this Network Total are
	// available in the details endpoint.
	IsComplete bool `json:"is_complete"`
	// Whether financial account is for the benefit of another entity
	IsForBenefitOf bool `json:"is_for_benefit_of"`
	// Whether Lithic decisioned on the token, and if so, what the decision was.
	// APPROVED/VERIFICATION_REQUIRED/DENIED.
	IssuerDecision ParsedWebhookEventIssuerDecision `json:"issuer_decision"`
	// The last 4 digits of the bank account. Derived by Lithic from the account number
	// passed
	LastFour string `json:"last_four"`
	// If applicable, represents the account_holder's last name.
	LastName string `json:"last_name"`
	// This field can have the runtime type of
	// [CardAuthorizationApprovalRequestWebhookEventLatestChallenge].
	LatestChallenge interface{} `json:"latest_challenge"`
	// If applicable, represents the account_holder's business name.
	LegalBusinessName string `json:"legal_business_name"`
	// This field can have the runtime type of [DisputeV2LiabilityAllocation].
	LiabilityAllocation interface{} `json:"liability_allocation"`
	// This field can have the runtime type of [shared.Merchant],
	// [ThreeDSAuthenticationMerchant].
	Merchant interface{} `json:"merchant"`
	// Deprecated, use `amounts`. The amount that the merchant will receive,
	// denominated in `merchant_currency` and in the smallest currency unit. Note the
	// amount includes `acquirer_fee`, similar to `authorization_amount`. It will be
	// different from `authorization_amount` if the merchant is taking payment in a
	// different currency.
	//
	// Deprecated: deprecated
	MerchantAmount int64 `json:"merchant_amount,nullable"`
	// Analogous to the 'authorization_amount', but in the merchant currency.
	//
	// Deprecated: deprecated
	MerchantAuthorizationAmount int64 `json:"merchant_authorization_amount,nullable"`
	// 3-character alphabetic ISO 4217 code for the local currency of the transaction.
	//
	// Deprecated: deprecated
	MerchantCurrency string `json:"merchant_currency"`
	// Either PAYMENT_AUTHENTICATION or NON_PAYMENT_AUTHENTICATION. For
	// NON_PAYMENT_AUTHENTICATION, additional_data and transaction fields are not
	// populated.
	MessageCategory ParsedWebhookEventMessageCategory `json:"message_category"`
	// Transfer method
	Method ParsedWebhookEventMethod `json:"method"`
	// This field can have the runtime type of [PaymentMethodAttributes].
	MethodAttributes interface{} `json:"method_attributes"`
	// This field can have the runtime type of [LoanTapeMinimumPaymentBalance].
	MinimumPaymentBalance interface{} `json:"minimum_payment_balance"`
	// 6-digit North American Industry Classification System (NAICS) code for the
	// business. Only present if naics_code was included in the update request.
	NaicsCode string `json:"naics_code"`
	// The nickname for this External Bank Account
	Name string `json:"name,nullable"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness string `json:"nature_of_business"`
	// Card network of the authorization.
	Network ParsedWebhookEventNetwork `json:"network,nullable"`
	// This field can have the runtime type of [[]string].
	NetworkClaimIDs interface{} `json:"network_claim_ids"`
	// Date that the dispute was submitted to the network.
	NetworkFiledDate time.Time `json:"network_filed_date,nullable" format:"date-time"`
	// Network reason code used to file the dispute.
	NetworkReasonCode string `json:"network_reason_code,nullable"`
	// Network-provided score assessing risk level associated with a given
	// authorization. Scores are on a range of 0-999, with 0 representing the lowest
	// risk and 999 representing the highest risk. For Visa transactions, where the raw
	// score has a range of 0-99, Lithic will normalize the score by multiplying the
	// raw score by 10x.
	NetworkRiskScore int64 `json:"network_risk_score,nullable"`
	// This field can have the runtime type of
	// [[]FundingEventNetworkSettlementSummary].
	NetworkSettlementSummary interface{} `json:"network_settlement_summary"`
	// This field can have the runtime type of
	// [CardAuthorizationApprovalRequestWebhookEventNetworkSpecificData].
	NetworkSpecificData interface{} `json:"network_specific_data"`
	// Date when the next payment is due
	NextPaymentDueDate time.Time `json:"next_payment_due_date" format:"date"`
	// Date when the next billing period will end
	NextStatementEndDate time.Time `json:"next_statement_end_date" format:"date"`
	Nickname             string    `json:"nickname,nullable"`
	// Total amount of gross other fees outside of interchange. (This field is
	// deprecated and will be removed in a future version of the API. To compute total
	// amounts, Lithic recommends that customers sum the relevant settlement amounts
	// found within `details`.)
	//
	// Deprecated: deprecated
	OtherFeesGrossAmount int64 `json:"other_fees_gross_amount"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner string `json:"owner"`
	// Owner Type
	OwnerType OwnerType `json:"owner_type"`
	// This field can have the runtime type of [LoanTapePaymentAllocation].
	PaymentAllocation interface{} `json:"payment_allocation"`
	// Date when the payment is due
	PaymentDueDate time.Time                     `json:"payment_due_date,nullable" format:"date"`
	PaymentType    ParsedWebhookEventPaymentType `json:"payment_type"`
	// This field can have the runtime type of [StatementPayoffDetails].
	PayoffDetails interface{} `json:"payoff_details"`
	// Pending amount of the transaction in the currency's smallest unit (e.g., cents),
	// including any acquirer fees.
	//
	// The value of this field will go to zero over time once the financial transaction
	// is settled.
	PendingAmount int64           `json:"pending_amount"`
	PeriodTotals  StatementTotals `json:"period_totals"`
	// If updated, the newly updated phone_number associated with the account_holder
	// otherwise the existing phone_number is provided.
	PhoneNumber string `json:"phone_number"`
	// This field can have the runtime type of
	// [CardAuthorizationApprovalRequestWebhookEventPos], [TransactionPos].
	Pos interface{} `json:"pos"`
	// Date dispute entered pre-arbitration.
	PrearbitrationDate time.Time `json:"prearbitration_date,nullable" format:"date-time"`
	// The previous expiration month of the card.
	PreviousExpMonth string `json:"previous_exp_month"`
	// The previous expiration year of the card.
	PreviousExpYear string `json:"previous_exp_year"`
	// This field can have the runtime type of [interface{}].
	PreviousFields interface{} `json:"previous_fields"`
	// Time of the previous high watermark
	PreviousHighWatermark time.Time `json:"previous_high_watermark" format:"date-time"`
	// This field can have the runtime type of [LoanTapePreviousStatementBalance].
	PreviousStatementBalance interface{} `json:"previous_statement_balance"`
	// Unique identifier for the dispute from the network. If there are multiple, this
	// will be the first claim id set by the network
	PrimaryClaimID string `json:"primary_claim_id,nullable"`
	// Dispute reason:
	//
	//   - `ATM_CASH_MISDISPENSE`: ATM cash misdispense.
	//   - `CANCELLED`: Transaction was cancelled by the customer.
	//   - `DUPLICATED`: The transaction was a duplicate.
	//   - `FRAUD_CARD_NOT_PRESENT`: Fraudulent transaction, card not present.
	//   - `FRAUD_CARD_PRESENT`: Fraudulent transaction, card present.
	//   - `FRAUD_OTHER`: Fraudulent transaction, other types such as questionable
	//     merchant activity.
	//   - `GOODS_SERVICES_NOT_AS_DESCRIBED`: The goods or services were not as
	//     described.
	//   - `GOODS_SERVICES_NOT_RECEIVED`: The goods or services were not received.
	//   - `INCORRECT_AMOUNT`: The transaction amount was incorrect.
	//   - `MISSING_AUTH`: The transaction was missing authorization.
	//   - `OTHER`: Other reason.
	//   - `PROCESSING_ERROR`: Processing error.
	//   - `REFUND_NOT_PROCESSED`: The refund was not processed.
	//   - `RECURRING_TRANSACTION_NOT_CANCELLED`: The recurring transaction was not
	//     cancelled.
	Reason ParsedWebhookEventReason `json:"reason"`
	// This field can have the runtime type of [PaymentRelatedAccountTokens].
	RelatedAccountTokens interface{} `json:"related_account_tokens"`
	// The token of the card that was replaced, if the new card is a replacement card.
	ReplacementFor string `json:"replacement_for,nullable" format:"uuid"`
	// This field can have the runtime type of [time.Time], [string].
	ReportDate interface{} `json:"report_date"`
	// Date the representment was received.
	RepresentmentDate time.Time `json:"representment_date,nullable" format:"date-time"`
	// This field can have the runtime type of
	// [[]AccountHolderDocumentUpdatedWebhookEventRequiredDocumentUpload].
	RequiredDocumentUploads interface{} `json:"required_document_uploads"`
	// This field can have the runtime type of [[]RequiredDocument].
	RequiredDocuments interface{} `json:"required_documents"`
	// Date that the dispute was resolved.
	ResolutionDate time.Time `json:"resolution_date,nullable" format:"date-time"`
	// Note by Dispute team on the case resolution.
	ResolutionNote string `json:"resolution_note,nullable"`
	// Reason for the dispute resolution:
	//
	// - `CASE_LOST`: This case was lost at final arbitration.
	// - `NETWORK_REJECTED`: Network rejected.
	// - `NO_DISPUTE_RIGHTS_3DS`: No dispute rights, 3DS.
	// - `NO_DISPUTE_RIGHTS_BELOW_THRESHOLD`: No dispute rights, below threshold.
	// - `NO_DISPUTE_RIGHTS_CONTACTLESS`: No dispute rights, contactless.
	// - `NO_DISPUTE_RIGHTS_HYBRID`: No dispute rights, hybrid.
	// - `NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS`: No dispute rights, max chargebacks.
	// - `NO_DISPUTE_RIGHTS_OTHER`: No dispute rights, other.
	// - `PAST_FILING_DATE`: Past filing date.
	// - `PREARBITRATION_REJECTED`: Prearbitration rejected.
	// - `PROCESSOR_REJECTED_OTHER`: Processor rejected, other.
	// - `REFUNDED`: Refunded.
	// - `REFUNDED_AFTER_CHARGEBACK`: Refunded after chargeback.
	// - `WITHDRAWN`: Withdrawn.
	// - `WON_ARBITRATION`: Won arbitration.
	// - `WON_FIRST_CHARGEBACK`: Won first chargeback.
	// - `WON_PREARBITRATION`: Won prearbitration.
	ResolutionReason ParsedWebhookEventResolutionReason `json:"resolution_reason,nullable"`
	Result           ParsedWebhookEventResult           `json:"result"`
	// This field can have the runtime type of [BacktestResultsResults].
	Results interface{} `json:"results"`
	// Routing Number
	RoutingNumber string `json:"routing_number,nullable"`
	// This field can have the runtime type of [[]TokenizationRuleResult].
	RuleResults interface{} `json:"rule_results"`
	// Deprecated, use `amounts`. Amount (in cents) of the transaction that has been
	// settled, including any acquirer fees.
	//
	// Deprecated: deprecated
	SettledAmount int64 `json:"settled_amount"`
	// The total net amount of cash moved. (net value of settled_gross_amount,
	// interchange, fees). (This field is deprecated and will be removed in a future
	// version of the API. To compute total amounts, Lithic recommends that customers
	// sum the relevant settlement amounts found within `details`.)
	//
	// Deprecated: deprecated
	SettledNetAmount int64 `json:"settled_net_amount"`
	// The institution responsible for settlement. For Mastercard: same as
	// `institution_id`. For Maestro: billing ICA. For Visa: Funds Transfer SRE
	// (FTSRE).
	SettlementInstitutionID string `json:"settlement_institution_id"`
	// Settlement service.
	SettlementService string `json:"settlement_service"`
	// The specific shipping method used to ship the card.
	ShippingMethod ParsedWebhookEventShippingMethod `json:"shipping_method"`
	// This field can have the runtime type of [BacktestResultsSimulationParameters].
	SimulationParameters interface{} `json:"simulation_parameters"`
	// Transaction source
	Source ParsedWebhookEventSource `json:"source"`
	// Balance at the start of the day
	StartingBalance int64 `json:"starting_balance"`
	// The current state of the card.
	State string `json:"state"`
	// Date when the billing period ended
	StatementEndDate time.Time `json:"statement_end_date" format:"date"`
	// Date when the billing period began
	StatementStartDate time.Time                       `json:"statement_start_date" format:"date"`
	StatementType      ParsedWebhookEventStatementType `json:"statement_type"`
	// The status of the account_holder that was created.
	Status ParsedWebhookEventStatus `json:"status,nullable"`
	// This field can have the runtime type of [[]string].
	StatusReason interface{} `json:"status_reason"`
	// This field can have the runtime type of [[]string].
	StatusReasons interface{} `json:"status_reasons"`
	// Substatus for the financial account
	Substatus ParsedWebhookEventSubstatus `json:"substatus,nullable"`
	// This field can have the runtime type of [map[string]string].
	Tags interface{} `json:"tags"`
	// Indicates whether a challenge is requested for this transaction
	//
	//   - `NO_PREFERENCE` - No Preference
	//   - `NO_CHALLENGE_REQUESTED` - No Challenge Requested
	//   - `CHALLENGE_PREFERENCE` - Challenge requested (3DS Requestor preference)
	//   - `CHALLENGE_MANDATE` - Challenge requested (Mandate)
	//   - `NO_CHALLENGE_RISK_ALREADY_ASSESSED` - No Challenge requested (Transactional
	//     risk analysis is already performed)
	//   - `DATA_SHARE_ONLY` - No Challenge requested (Data Share Only)
	//   - `OTHER` - Other indicators not captured by above. These are rarely used
	ThreeDSRequestorChallengeIndicator ParsedWebhookEventThreeDSRequestorChallengeIndicator `json:"three_ds_requestor_challenge_indicator"`
	// Type of 3DS Requestor Initiated (3RI) request — i.e., a 3DS authentication that
	// takes place at the initiation of the merchant rather than the cardholder. The
	// most common example of this is where a merchant is authenticating before billing
	// for a recurring transaction such as a pay TV subscription or a utility bill.
	// Maps to EMV 3DS field `threeRIInd`.
	ThreeRiRequestType ParsedWebhookEventThreeRiRequestType `json:"three_ri_request_type,nullable"`
	// Interest tier to which this account belongs to
	Tier string `json:"tier,nullable"`
	// Globally unique identifier for the financial account or card that will receive
	// the funds. Accepted type dependent on the program's use case
	ToFinancialAccountToken string       `json:"to_financial_account_token" format:"uuid"`
	TokenInfo               TokenInfo    `json:"token_info,nullable"`
	Tokenization            Tokenization `json:"tokenization"`
	// The channel through which the tokenization was made.
	TokenizationChannel ParsedWebhookEventTokenizationChannel `json:"tokenization_channel"`
	// This field can have the runtime type of [[]TokenizationDeclineReason].
	TokenizationDeclineReasons interface{} `json:"tokenization_decline_reasons"`
	// This field can have the runtime type of
	// [DigitalWalletTokenizationResultWebhookEventTokenizationResultDetails],
	// [TokenizationResultWebhookEventTokenizationResultDetails].
	TokenizationResultDetails interface{} `json:"tokenization_result_details"`
	// The source of the tokenization.
	TokenizationSource ParsedWebhookEventTokenizationSource `json:"tokenization_source"`
	// This field can have the runtime type of [[]TokenizationTfaReason].
	TokenizationTfaReasons interface{} `json:"tokenization_tfa_reasons"`
	// Unique identifier for the digital wallet token attempt
	TokenizationToken string `json:"tokenization_token"`
	// The tracking number of the shipment.
	TrackingNumber string `json:"tracking_number,nullable"`
	// This field can have the runtime type of [ThreeDSAuthenticationTransaction].
	Transaction interface{} `json:"transaction"`
	// The entity that initiated the transaction.
	TransactionInitiator ParsedWebhookEventTransactionInitiator `json:"transaction_initiator"`
	// This field can have the runtime type of [BookTransferResponseTransactionSeries],
	// [ManagementOperationTransactionTransactionSeries], [DisputeV2TransactionSeries].
	TransactionSeries interface{} `json:"transaction_series"`
	// The token of the transaction that the enhanced data is associated with.
	TransactionToken string `json:"transaction_token" format:"uuid"`
	// The total amount of settlement impacting transactions (excluding interchange,
	// fees, and disputes). (This field is deprecated and will be removed in a future
	// version of the API. To compute total amounts, Lithic recommends that customers
	// sum the relevant settlement amounts found within `details`.)
	//
	// Deprecated: deprecated
	TransactionsGrossAmount int64 `json:"transactions_gross_amount"`
	// Deprecated: approximate time-to-live for the authorization.
	Ttl time.Time `json:"ttl" format:"date-time"`
	// Account Type
	Type ParsedWebhookEventType `json:"type"`
	// This field can have the runtime type of
	// [ParsedWebhookEventKYBPayloadUpdateRequest],
	// [ParsedWebhookEventKYCPayloadUpdateRequest].
	UpdateRequest interface{} `json:"update_request"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated" format:"date-time"`
	// Upload status types:
	//
	// - `DELETED` - Evidence was deleted.
	// - `ERROR` - Evidence upload failed.
	// - `PENDING` - Evidence is pending upload.
	// - `REJECTED` - Evidence was rejected.
	// - `UPLOADED` - Evidence was uploaded.
	UploadStatus ParsedWebhookEventUploadStatus `json:"upload_status"`
	// URL to upload evidence. Only shown when `upload_status` is `PENDING`.
	UploadURL string `json:"upload_url"`
	// User Defined ID
	UserDefinedID string `json:"user_defined_id,nullable"`
	// User-defined status for the financial account
	UserDefinedStatus string `json:"user_defined_status,nullable"`
	// The number of attempts at verification
	VerificationAttempts int64 `json:"verification_attempts"`
	// Optional free text description of the reason for the failed verification. For
	// ACH micro-deposits returned, this field will display the reason return code sent
	// by the ACH network
	VerificationFailedReason string `json:"verification_failed_reason,nullable"`
	// Verification Method
	VerificationMethod VerificationMethod `json:"verification_method"`
	// Verification State
	VerificationState ParsedWebhookEventVerificationState `json:"verification_state"`
	// Version number of the loan tape. This starts at 1
	Version               int64                 `json:"version"`
	WalletDecisioningInfo WalletDecisioningInfo `json:"wallet_decisioning_info"`
	// Company website URL.
	WebsiteURL string                 `json:"website_url"`
	YtdTotals  StatementTotals        `json:"ytd_totals"`
	JSON       parsedWebhookEventJSON `json:"-"`
	union      ParsedWebhookEventUnion
}

// parsedWebhookEventJSON contains the JSON metadata for the struct
// [ParsedWebhookEvent]
type parsedWebhookEventJSON struct {
	Token                              apijson.Field
	AccountHolderToken                 apijson.Field
	AccountNumber                      apijson.Field
	AccountStanding                    apijson.Field
	AccountToken                       apijson.Field
	AccountType                        apijson.Field
	AcquirerFee                        apijson.Field
	AcquirerReferenceNumber            apijson.Field
	ActivationMethod                   apijson.Field
	AdditionalData                     apijson.Field
	Address                            apijson.Field
	Amount                             apijson.Field
	AmountDue                          apijson.Field
	Amounts                            apijson.Field
	App                                apijson.Field
	ArbitrationDate                    apijson.Field
	AuthenticationCode                 apijson.Field
	AuthenticationObject               apijson.Field
	AuthenticationRequestType          apijson.Field
	AuthenticationResult               apijson.Field
	AuthorizationAmount                apijson.Field
	AuthorizationCode                  apijson.Field
	AvailableCredit                    apijson.Field
	Avs                                apijson.Field
	BacktestToken                      apijson.Field
	Balances                           apijson.Field
	Browser                            apijson.Field
	BulkOrderToken                     apijson.Field
	BusinessAccountToken               apijson.Field
	Card                               apijson.Field
	CardExpiryCheck                    apijson.Field
	CardToken                          apijson.Field
	Cardholder                         apijson.Field
	CardholderAuthentication           apijson.Field
	CardholderCurrency                 apijson.Field
	CaseID                             apijson.Field
	CashAmount                         apijson.Field
	Cashback                           apijson.Field
	Category                           apijson.Field
	Challenge                          apijson.Field
	ChallengeMetadata                  apijson.Field
	ChallengeOrchestratedBy            apijson.Field
	Channel                            apijson.Field
	CollectionResourceType             apijson.Field
	CollectionTokens                   apijson.Field
	Common                             apijson.Field
	CompanyID                          apijson.Field
	ConversionRate                     apijson.Field
	Country                            apijson.Field
	Created                            apijson.Field
	CreditConfiguration                apijson.Field
	CreditLimit                        apijson.Field
	CreditProductToken                 apijson.Field
	Currency                           apijson.Field
	CustomerFiledDate                  apijson.Field
	CustomerNote                       apijson.Field
	CustomerTokenizationDecision       apijson.Field
	Cycle                              apijson.Field
	Data                               apijson.Field
	Date                               apijson.Field
	DayTotals                          apijson.Field
	DaysInBillingCycle                 apijson.Field
	DecisionMadeBy                     apijson.Field
	Descriptor                         apijson.Field
	Details                            apijson.Field
	Device                             apijson.Field
	DigitalWalletTokenMetadata         apijson.Field
	Direction                          apijson.Field
	Disposition                        apijson.Field
	DisputeToken                       apijson.Field
	DisputesGrossAmount                apijson.Field
	Dob                                apijson.Field
	DocumentType                       apijson.Field
	DoingBusinessAs                    apijson.Field
	DownloadURL                        apijson.Field
	Email                              apijson.Field
	EndingBalance                      apijson.Field
	EntityToken                        apijson.Field
	EventToken                         apijson.Field
	EventType                          apijson.Field
	Events                             apijson.Field
	ExcessCredits                      apijson.Field
	ExpMonth                           apijson.Field
	ExpYear                            apijson.Field
	ExpectedReleaseDate                apijson.Field
	ExternalBankAccountToken           apijson.Field
	ExternalID                         apijson.Field
	ExternalResource                   apijson.Field
	Family                             apijson.Field
	Filename                           apijson.Field
	FinancialAccountToken              apijson.Field
	FirstName                          apijson.Field
	Fleet                              apijson.Field
	FleetInfo                          apijson.Field
	FromFinancialAccountToken          apijson.Field
	HighWatermark                      apijson.Field
	InstitutionID                      apijson.Field
	InterchangeGrossAmount             apijson.Field
	InterestDetails                    apijson.Field
	IsComplete                         apijson.Field
	IsForBenefitOf                     apijson.Field
	IssuerDecision                     apijson.Field
	LastFour                           apijson.Field
	LastName                           apijson.Field
	LatestChallenge                    apijson.Field
	LegalBusinessName                  apijson.Field
	LiabilityAllocation                apijson.Field
	Merchant                           apijson.Field
	MerchantAmount                     apijson.Field
	MerchantAuthorizationAmount        apijson.Field
	MerchantCurrency                   apijson.Field
	MessageCategory                    apijson.Field
	Method                             apijson.Field
	MethodAttributes                   apijson.Field
	MinimumPaymentBalance              apijson.Field
	NaicsCode                          apijson.Field
	Name                               apijson.Field
	NatureOfBusiness                   apijson.Field
	Network                            apijson.Field
	NetworkClaimIDs                    apijson.Field
	NetworkFiledDate                   apijson.Field
	NetworkReasonCode                  apijson.Field
	NetworkRiskScore                   apijson.Field
	NetworkSettlementSummary           apijson.Field
	NetworkSpecificData                apijson.Field
	NextPaymentDueDate                 apijson.Field
	NextStatementEndDate               apijson.Field
	Nickname                           apijson.Field
	OtherFeesGrossAmount               apijson.Field
	Owner                              apijson.Field
	OwnerType                          apijson.Field
	PaymentAllocation                  apijson.Field
	PaymentDueDate                     apijson.Field
	PaymentType                        apijson.Field
	PayoffDetails                      apijson.Field
	PendingAmount                      apijson.Field
	PeriodTotals                       apijson.Field
	PhoneNumber                        apijson.Field
	Pos                                apijson.Field
	PrearbitrationDate                 apijson.Field
	PreviousExpMonth                   apijson.Field
	PreviousExpYear                    apijson.Field
	PreviousFields                     apijson.Field
	PreviousHighWatermark              apijson.Field
	PreviousStatementBalance           apijson.Field
	PrimaryClaimID                     apijson.Field
	Reason                             apijson.Field
	RelatedAccountTokens               apijson.Field
	ReplacementFor                     apijson.Field
	ReportDate                         apijson.Field
	RepresentmentDate                  apijson.Field
	RequiredDocumentUploads            apijson.Field
	RequiredDocuments                  apijson.Field
	ResolutionDate                     apijson.Field
	ResolutionNote                     apijson.Field
	ResolutionReason                   apijson.Field
	Result                             apijson.Field
	Results                            apijson.Field
	RoutingNumber                      apijson.Field
	RuleResults                        apijson.Field
	SettledAmount                      apijson.Field
	SettledNetAmount                   apijson.Field
	SettlementInstitutionID            apijson.Field
	SettlementService                  apijson.Field
	ShippingMethod                     apijson.Field
	SimulationParameters               apijson.Field
	Source                             apijson.Field
	StartingBalance                    apijson.Field
	State                              apijson.Field
	StatementEndDate                   apijson.Field
	StatementStartDate                 apijson.Field
	StatementType                      apijson.Field
	Status                             apijson.Field
	StatusReason                       apijson.Field
	StatusReasons                      apijson.Field
	Substatus                          apijson.Field
	Tags                               apijson.Field
	ThreeDSRequestorChallengeIndicator apijson.Field
	ThreeRiRequestType                 apijson.Field
	Tier                               apijson.Field
	ToFinancialAccountToken            apijson.Field
	TokenInfo                          apijson.Field
	Tokenization                       apijson.Field
	TokenizationChannel                apijson.Field
	TokenizationDeclineReasons         apijson.Field
	TokenizationResultDetails          apijson.Field
	TokenizationSource                 apijson.Field
	TokenizationTfaReasons             apijson.Field
	TokenizationToken                  apijson.Field
	TrackingNumber                     apijson.Field
	Transaction                        apijson.Field
	TransactionInitiator               apijson.Field
	TransactionSeries                  apijson.Field
	TransactionToken                   apijson.Field
	TransactionsGrossAmount            apijson.Field
	Ttl                                apijson.Field
	Type                               apijson.Field
	UpdateRequest                      apijson.Field
	Updated                            apijson.Field
	UploadStatus                       apijson.Field
	UploadURL                          apijson.Field
	UserDefinedID                      apijson.Field
	UserDefinedStatus                  apijson.Field
	VerificationAttempts               apijson.Field
	VerificationFailedReason           apijson.Field
	VerificationMethod                 apijson.Field
	VerificationState                  apijson.Field
	Version                            apijson.Field
	WalletDecisioningInfo              apijson.Field
	WebsiteURL                         apijson.Field
	YtdTotals                          apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r parsedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *ParsedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = ParsedWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ParsedWebhookEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [AccountHolderCreatedWebhookEvent],
// [ParsedWebhookEventKYBPayload], [ParsedWebhookEventKYCPayload],
// [ParsedWebhookEventLegacyPayload], [AccountHolderVerificationWebhookEvent],
// [AccountHolderDocumentUpdatedWebhookEvent],
// [CardAuthorizationApprovalRequestWebhookEvent],
// [TokenizationDecisioningRequestWebhookEvent],
// [AuthRulesBacktestReportCreatedWebhookEvent], [BalanceUpdatedWebhookEvent],
// [BookTransferTransactionCreatedWebhookEvent],
// [BookTransferTransactionUpdatedWebhookEvent], [CardCreatedWebhookEvent],
// [CardConvertedWebhookEvent], [CardRenewedWebhookEvent],
// [CardReissuedWebhookEvent], [CardShippedWebhookEvent],
// [CardUpdatedWebhookEvent], [CardTransactionUpdatedWebhookEvent],
// [CardTransactionEnhancedDataCreatedWebhookEvent],
// [CardTransactionEnhancedDataUpdatedWebhookEvent],
// [DigitalWalletTokenizationApprovalRequestWebhookEvent],
// [DigitalWalletTokenizationResultWebhookEvent],
// [DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEvent],
// [DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEvent],
// [DigitalWalletTokenizationUpdatedWebhookEvent], [DisputeUpdatedWebhookEvent],
// [DisputeEvidenceUploadFailedWebhookEvent],
// [ExternalBankAccountCreatedWebhookEvent],
// [ExternalBankAccountUpdatedWebhookEvent], [ExternalPaymentCreatedWebhookEvent],
// [ExternalPaymentUpdatedWebhookEvent], [FinancialAccountCreatedWebhookEvent],
// [FinancialAccountUpdatedWebhookEvent], [FundingEventCreatedWebhookEvent],
// [LoanTapeCreatedWebhookEvent], [LoanTapeUpdatedWebhookEvent],
// [ManagementOperationCreatedWebhookEvent],
// [ManagementOperationUpdatedWebhookEvent],
// [InternalTransactionCreatedWebhookEvent],
// [InternalTransactionUpdatedWebhookEvent], [NetworkTotalCreatedWebhookEvent],
// [NetworkTotalUpdatedWebhookEvent], [PaymentTransactionCreatedWebhookEvent],
// [PaymentTransactionUpdatedWebhookEvent], [SettlementReportUpdatedWebhookEvent],
// [StatementsCreatedWebhookEvent], [ThreeDSAuthenticationCreatedWebhookEvent],
// [ThreeDSAuthenticationUpdatedWebhookEvent],
// [ThreeDSAuthenticationChallengeWebhookEvent],
// [TokenizationApprovalRequestWebhookEvent], [TokenizationResultWebhookEvent],
// [TokenizationTwoFactorAuthenticationCodeWebhookEvent],
// [TokenizationTwoFactorAuthenticationCodeSentWebhookEvent],
// [TokenizationUpdatedWebhookEvent],
// [ThreeDSAuthenticationApprovalRequestWebhookEvent],
// [DisputeTransactionCreatedWebhookEvent],
// [DisputeTransactionUpdatedWebhookEvent].
func (r ParsedWebhookEvent) AsUnion() ParsedWebhookEventUnion {
	return r.union
}

// KYB payload for an updated account holder.
//
// Union satisfied by [AccountHolderCreatedWebhookEvent],
// [ParsedWebhookEventKYBPayload], [ParsedWebhookEventKYCPayload],
// [ParsedWebhookEventLegacyPayload], [AccountHolderVerificationWebhookEvent],
// [AccountHolderDocumentUpdatedWebhookEvent],
// [CardAuthorizationApprovalRequestWebhookEvent],
// [TokenizationDecisioningRequestWebhookEvent],
// [AuthRulesBacktestReportCreatedWebhookEvent], [BalanceUpdatedWebhookEvent],
// [BookTransferTransactionCreatedWebhookEvent],
// [BookTransferTransactionUpdatedWebhookEvent], [CardCreatedWebhookEvent],
// [CardConvertedWebhookEvent], [CardRenewedWebhookEvent],
// [CardReissuedWebhookEvent], [CardShippedWebhookEvent],
// [CardUpdatedWebhookEvent], [CardTransactionUpdatedWebhookEvent],
// [CardTransactionEnhancedDataCreatedWebhookEvent],
// [CardTransactionEnhancedDataUpdatedWebhookEvent],
// [DigitalWalletTokenizationApprovalRequestWebhookEvent],
// [DigitalWalletTokenizationResultWebhookEvent],
// [DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEvent],
// [DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEvent],
// [DigitalWalletTokenizationUpdatedWebhookEvent], [DisputeUpdatedWebhookEvent],
// [DisputeEvidenceUploadFailedWebhookEvent],
// [ExternalBankAccountCreatedWebhookEvent],
// [ExternalBankAccountUpdatedWebhookEvent], [ExternalPaymentCreatedWebhookEvent],
// [ExternalPaymentUpdatedWebhookEvent], [FinancialAccountCreatedWebhookEvent],
// [FinancialAccountUpdatedWebhookEvent], [FundingEventCreatedWebhookEvent],
// [LoanTapeCreatedWebhookEvent], [LoanTapeUpdatedWebhookEvent],
// [ManagementOperationCreatedWebhookEvent],
// [ManagementOperationUpdatedWebhookEvent],
// [InternalTransactionCreatedWebhookEvent],
// [InternalTransactionUpdatedWebhookEvent], [NetworkTotalCreatedWebhookEvent],
// [NetworkTotalUpdatedWebhookEvent], [PaymentTransactionCreatedWebhookEvent],
// [PaymentTransactionUpdatedWebhookEvent], [SettlementReportUpdatedWebhookEvent],
// [StatementsCreatedWebhookEvent], [ThreeDSAuthenticationCreatedWebhookEvent],
// [ThreeDSAuthenticationUpdatedWebhookEvent],
// [ThreeDSAuthenticationChallengeWebhookEvent],
// [TokenizationApprovalRequestWebhookEvent], [TokenizationResultWebhookEvent],
// [TokenizationTwoFactorAuthenticationCodeWebhookEvent],
// [TokenizationTwoFactorAuthenticationCodeSentWebhookEvent],
// [TokenizationUpdatedWebhookEvent],
// [ThreeDSAuthenticationApprovalRequestWebhookEvent],
// [DisputeTransactionCreatedWebhookEvent] or
// [DisputeTransactionUpdatedWebhookEvent].
type ParsedWebhookEventUnion interface {
	implementsParsedWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParsedWebhookEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountHolderCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParsedWebhookEventKYBPayload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParsedWebhookEventKYCPayload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParsedWebhookEventLegacyPayload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountHolderVerificationWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountHolderDocumentUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardAuthorizationApprovalRequestWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenizationDecisioningRequestWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRulesBacktestReportCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BalanceUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BookTransferTransactionCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BookTransferTransactionUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardConvertedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardRenewedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardReissuedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardShippedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardTransactionUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardTransactionEnhancedDataCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardTransactionEnhancedDataUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DigitalWalletTokenizationApprovalRequestWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DigitalWalletTokenizationResultWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DigitalWalletTokenizationTwoFactorAuthenticationCodeWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DigitalWalletTokenizationTwoFactorAuthenticationCodeSentWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DigitalWalletTokenizationUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeEvidenceUploadFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalBankAccountCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalBankAccountUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalPaymentCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalPaymentUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FinancialAccountCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FinancialAccountUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FundingEventCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LoanTapeCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LoanTapeUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagementOperationCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagementOperationUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InternalTransactionCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InternalTransactionUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NetworkTotalCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NetworkTotalUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentTransactionCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentTransactionUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettlementReportUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StatementsCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ThreeDSAuthenticationCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ThreeDSAuthenticationUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ThreeDSAuthenticationChallengeWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenizationApprovalRequestWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenizationResultWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenizationTwoFactorAuthenticationCodeWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenizationTwoFactorAuthenticationCodeSentWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenizationUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ThreeDSAuthenticationApprovalRequestWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeTransactionCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeTransactionUpdatedWebhookEvent{}),
		},
	)
}

// KYB payload for an updated account holder.
type ParsedWebhookEventKYBPayload struct {
	// The token of the account_holder that was created.
	Token string `json:"token,required" format:"uuid"`
	// Original request to update the account holder.
	UpdateRequest ParsedWebhookEventKYBPayloadUpdateRequest `json:"update_request,required"`
	// The type of event that occurred.
	EventType ParsedWebhookEventKYBPayloadEventType `json:"event_type"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID string `json:"external_id"`
	// 6-digit North American Industry Classification System (NAICS) code for the
	// business. Only present if naics_code was included in the update request.
	NaicsCode string `json:"naics_code"`
	// Short description of the company's line of business (i.e., what does the company
	// do?).
	NatureOfBusiness string `json:"nature_of_business"`
	// Company website URL.
	WebsiteURL string                           `json:"website_url"`
	JSON       parsedWebhookEventKYBPayloadJSON `json:"-"`
}

// parsedWebhookEventKYBPayloadJSON contains the JSON metadata for the struct
// [ParsedWebhookEventKYBPayload]
type parsedWebhookEventKYBPayloadJSON struct {
	Token            apijson.Field
	UpdateRequest    apijson.Field
	EventType        apijson.Field
	ExternalID       apijson.Field
	NaicsCode        apijson.Field
	NatureOfBusiness apijson.Field
	WebsiteURL       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ParsedWebhookEventKYBPayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYBPayloadJSON) RawJSON() string {
	return r.raw
}

func (r ParsedWebhookEventKYBPayload) implementsParsedWebhookEvent() {}

// Original request to update the account holder.
type ParsedWebhookEventKYBPayloadUpdateRequest struct {
	// Deprecated.
	//
	// Deprecated: deprecated
	BeneficialOwnerEntities []KYBBusinessEntity `json:"beneficial_owner_entities"`
	// You must submit a list of all direct and indirect individuals with 25% or more
	// ownership in the company. A maximum of 4 beneficial owners can be submitted. If
	// no individual owns 25% of the company you do not need to send beneficial owner
	// information. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section I) for more background on individuals that should be included.
	BeneficialOwnerIndividuals []ParsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// Information for business for which the account is being opened and KYB is being
	// run.
	BusinessEntity KYBBusinessEntity `json:"business_entity"`
	// An individual with significant responsibility for managing the legal entity
	// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
	// Officer, Managing Member, General Partner, President, Vice President, or
	// Treasurer). This can be an executive, or someone who will have program-wide
	// access to the cards that Lithic will provide. In some cases, this individual
	// could also be a beneficial owner listed above. See
	// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
	// (Section II) for more background.
	ControlPerson ParsedWebhookEventKYBPayloadUpdateRequestControlPerson `json:"control_person"`
	JSON          parsedWebhookEventKYBPayloadUpdateRequestJSON          `json:"-"`
}

// parsedWebhookEventKYBPayloadUpdateRequestJSON contains the JSON metadata for the
// struct [ParsedWebhookEventKYBPayloadUpdateRequest]
type parsedWebhookEventKYBPayloadUpdateRequestJSON struct {
	BeneficialOwnerEntities    apijson.Field
	BeneficialOwnerIndividuals apijson.Field
	BusinessEntity             apijson.Field
	ControlPerson              apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ParsedWebhookEventKYBPayloadUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYBPayloadUpdateRequestJSON) RawJSON() string {
	return r.raw
}

type ParsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address ParsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddress `json:"address"`
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
	PhoneNumber string                                                                 `json:"phone_number"`
	JSON        parsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualJSON `json:"-"`
}

// parsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualJSON contains
// the JSON metadata for the struct
// [ParsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividual]
type parsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type ParsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddress struct {
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
	Address2 string                                                                         `json:"address2"`
	JSON     parsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddressJSON `json:"-"`
}

// parsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddressJSON
// contains the JSON metadata for the struct
// [ParsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddress]
type parsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYBPayloadUpdateRequestBeneficialOwnerIndividualsAddressJSON) RawJSON() string {
	return r.raw
}

// An individual with significant responsibility for managing the legal entity
// (e.g., a Chief Executive Officer, Chief Financial Officer, Chief Operating
// Officer, Managing Member, General Partner, President, Vice President, or
// Treasurer). This can be an executive, or someone who will have program-wide
// access to the cards that Lithic will provide. In some cases, this individual
// could also be a beneficial owner listed above. See
// [FinCEN requirements](https://www.fincen.gov/sites/default/files/shared/CDD_Rev6.7_Sept_2017_Certificate.pdf)
// (Section II) for more background.
type ParsedWebhookEventKYBPayloadUpdateRequestControlPerson struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address ParsedWebhookEventKYBPayloadUpdateRequestControlPersonAddress `json:"address"`
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
	PhoneNumber string                                                     `json:"phone_number"`
	JSON        parsedWebhookEventKYBPayloadUpdateRequestControlPersonJSON `json:"-"`
}

// parsedWebhookEventKYBPayloadUpdateRequestControlPersonJSON contains the JSON
// metadata for the struct [ParsedWebhookEventKYBPayloadUpdateRequestControlPerson]
type parsedWebhookEventKYBPayloadUpdateRequestControlPersonJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventKYBPayloadUpdateRequestControlPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYBPayloadUpdateRequestControlPersonJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type ParsedWebhookEventKYBPayloadUpdateRequestControlPersonAddress struct {
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
	Address2 string                                                            `json:"address2"`
	JSON     parsedWebhookEventKYBPayloadUpdateRequestControlPersonAddressJSON `json:"-"`
}

// parsedWebhookEventKYBPayloadUpdateRequestControlPersonAddressJSON contains the
// JSON metadata for the struct
// [ParsedWebhookEventKYBPayloadUpdateRequestControlPersonAddress]
type parsedWebhookEventKYBPayloadUpdateRequestControlPersonAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventKYBPayloadUpdateRequestControlPersonAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYBPayloadUpdateRequestControlPersonAddressJSON) RawJSON() string {
	return r.raw
}

// The type of event that occurred.
type ParsedWebhookEventKYBPayloadEventType string

const (
	ParsedWebhookEventKYBPayloadEventTypeAccountHolderUpdated ParsedWebhookEventKYBPayloadEventType = "account_holder.updated"
)

func (r ParsedWebhookEventKYBPayloadEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventKYBPayloadEventTypeAccountHolderUpdated:
		return true
	}
	return false
}

// KYC payload for an updated account holder.
type ParsedWebhookEventKYCPayload struct {
	// The token of the account_holder that was created.
	Token string `json:"token,required" format:"uuid"`
	// Original request to update the account holder.
	UpdateRequest ParsedWebhookEventKYCPayloadUpdateRequest `json:"update_request,required"`
	// The type of event that occurred.
	EventType ParsedWebhookEventKYCPayloadEventType `json:"event_type"`
	// A user provided id that can be used to link an account holder with an external
	// system
	ExternalID string                           `json:"external_id"`
	JSON       parsedWebhookEventKYCPayloadJSON `json:"-"`
}

// parsedWebhookEventKYCPayloadJSON contains the JSON metadata for the struct
// [ParsedWebhookEventKYCPayload]
type parsedWebhookEventKYCPayloadJSON struct {
	Token         apijson.Field
	UpdateRequest apijson.Field
	EventType     apijson.Field
	ExternalID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ParsedWebhookEventKYCPayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYCPayloadJSON) RawJSON() string {
	return r.raw
}

func (r ParsedWebhookEventKYCPayload) implementsParsedWebhookEvent() {}

// Original request to update the account holder.
type ParsedWebhookEventKYCPayloadUpdateRequest struct {
	// Information on the individual for whom the account is being opened and KYC is
	// being run.
	Individual ParsedWebhookEventKYCPayloadUpdateRequestIndividual `json:"individual"`
	JSON       parsedWebhookEventKYCPayloadUpdateRequestJSON       `json:"-"`
}

// parsedWebhookEventKYCPayloadUpdateRequestJSON contains the JSON metadata for the
// struct [ParsedWebhookEventKYCPayloadUpdateRequest]
type parsedWebhookEventKYCPayloadUpdateRequestJSON struct {
	Individual  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventKYCPayloadUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYCPayloadUpdateRequestJSON) RawJSON() string {
	return r.raw
}

// Information on the individual for whom the account is being opened and KYC is
// being run.
type ParsedWebhookEventKYCPayloadUpdateRequestIndividual struct {
	// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
	// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
	Address ParsedWebhookEventKYCPayloadUpdateRequestIndividualAddress `json:"address"`
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
	PhoneNumber string                                                  `json:"phone_number"`
	JSON        parsedWebhookEventKYCPayloadUpdateRequestIndividualJSON `json:"-"`
}

// parsedWebhookEventKYCPayloadUpdateRequestIndividualJSON contains the JSON
// metadata for the struct [ParsedWebhookEventKYCPayloadUpdateRequestIndividual]
type parsedWebhookEventKYCPayloadUpdateRequestIndividualJSON struct {
	Address     apijson.Field
	Dob         apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventKYCPayloadUpdateRequestIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYCPayloadUpdateRequestIndividualJSON) RawJSON() string {
	return r.raw
}

// Individual's current address - PO boxes, UPS drops, and FedEx drops are not
// acceptable; APO/FPO are acceptable. Only USA addresses are currently supported.
type ParsedWebhookEventKYCPayloadUpdateRequestIndividualAddress struct {
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
	Address2 string                                                         `json:"address2"`
	JSON     parsedWebhookEventKYCPayloadUpdateRequestIndividualAddressJSON `json:"-"`
}

// parsedWebhookEventKYCPayloadUpdateRequestIndividualAddressJSON contains the JSON
// metadata for the struct
// [ParsedWebhookEventKYCPayloadUpdateRequestIndividualAddress]
type parsedWebhookEventKYCPayloadUpdateRequestIndividualAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventKYCPayloadUpdateRequestIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventKYCPayloadUpdateRequestIndividualAddressJSON) RawJSON() string {
	return r.raw
}

// The type of event that occurred.
type ParsedWebhookEventKYCPayloadEventType string

const (
	ParsedWebhookEventKYCPayloadEventTypeAccountHolderUpdated ParsedWebhookEventKYCPayloadEventType = "account_holder.updated"
)

func (r ParsedWebhookEventKYCPayloadEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventKYCPayloadEventTypeAccountHolderUpdated:
		return true
	}
	return false
}

// Legacy payload for an updated account holder.
type ParsedWebhookEventLegacyPayload struct {
	// The token of the account_holder that was created.
	Token string `json:"token,required" format:"uuid"`
	// If applicable, represents the business account token associated with the
	// account_holder.
	BusinessAccountToken string `json:"business_account_token,nullable" format:"uuid"`
	// When the account_holder updated event was created
	Created time.Time `json:"created" format:"date-time"`
	// If updated, the newly updated email associated with the account_holder otherwise
	// the existing email is provided.
	Email string `json:"email"`
	// The type of event that occurred.
	EventType ParsedWebhookEventLegacyPayloadEventType `json:"event_type"`
	// If applicable, represents the external_id associated with the account_holder.
	ExternalID string `json:"external_id,nullable"`
	// If applicable, represents the account_holder's first name.
	FirstName string `json:"first_name"`
	// If applicable, represents the account_holder's last name.
	LastName string `json:"last_name"`
	// If applicable, represents the account_holder's business name.
	LegalBusinessName string `json:"legal_business_name"`
	// If updated, the newly updated phone_number associated with the account_holder
	// otherwise the existing phone_number is provided.
	PhoneNumber string                              `json:"phone_number"`
	JSON        parsedWebhookEventLegacyPayloadJSON `json:"-"`
}

// parsedWebhookEventLegacyPayloadJSON contains the JSON metadata for the struct
// [ParsedWebhookEventLegacyPayload]
type parsedWebhookEventLegacyPayloadJSON struct {
	Token                apijson.Field
	BusinessAccountToken apijson.Field
	Created              apijson.Field
	Email                apijson.Field
	EventType            apijson.Field
	ExternalID           apijson.Field
	FirstName            apijson.Field
	LastName             apijson.Field
	LegalBusinessName    apijson.Field
	PhoneNumber          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ParsedWebhookEventLegacyPayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventLegacyPayloadJSON) RawJSON() string {
	return r.raw
}

func (r ParsedWebhookEventLegacyPayload) implementsParsedWebhookEvent() {}

// The type of event that occurred.
type ParsedWebhookEventLegacyPayloadEventType string

const (
	ParsedWebhookEventLegacyPayloadEventTypeAccountHolderUpdated ParsedWebhookEventLegacyPayloadEventType = "account_holder.updated"
)

func (r ParsedWebhookEventLegacyPayloadEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventLegacyPayloadEventTypeAccountHolderUpdated:
		return true
	}
	return false
}

type ParsedWebhookEventAccountType string

const (
	ParsedWebhookEventAccountTypeChecking      ParsedWebhookEventAccountType = "CHECKING"
	ParsedWebhookEventAccountTypeSavings       ParsedWebhookEventAccountType = "SAVINGS"
	ParsedWebhookEventAccountTypeCredit        ParsedWebhookEventAccountType = "CREDIT"
	ParsedWebhookEventAccountTypeDebit         ParsedWebhookEventAccountType = "DEBIT"
	ParsedWebhookEventAccountTypeNotApplicable ParsedWebhookEventAccountType = "NOT_APPLICABLE"
)

func (r ParsedWebhookEventAccountType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventAccountTypeChecking, ParsedWebhookEventAccountTypeSavings, ParsedWebhookEventAccountTypeCredit, ParsedWebhookEventAccountTypeDebit, ParsedWebhookEventAccountTypeNotApplicable:
		return true
	}
	return false
}

// Type of authentication request - i.e., the type of transaction or interaction is
// causing the merchant to request an authentication. Maps to EMV 3DS field
// `threeDSRequestorAuthenticationInd`.
type ParsedWebhookEventAuthenticationRequestType string

const (
	ParsedWebhookEventAuthenticationRequestTypeAddCard                        ParsedWebhookEventAuthenticationRequestType = "ADD_CARD"
	ParsedWebhookEventAuthenticationRequestTypeBillingAgreement               ParsedWebhookEventAuthenticationRequestType = "BILLING_AGREEMENT"
	ParsedWebhookEventAuthenticationRequestTypeDelayedShipment                ParsedWebhookEventAuthenticationRequestType = "DELAYED_SHIPMENT"
	ParsedWebhookEventAuthenticationRequestTypeEmvTokenCardholderVerification ParsedWebhookEventAuthenticationRequestType = "EMV_TOKEN_CARDHOLDER_VERIFICATION"
	ParsedWebhookEventAuthenticationRequestTypeInstallmentTransaction         ParsedWebhookEventAuthenticationRequestType = "INSTALLMENT_TRANSACTION"
	ParsedWebhookEventAuthenticationRequestTypeMaintainCard                   ParsedWebhookEventAuthenticationRequestType = "MAINTAIN_CARD"
	ParsedWebhookEventAuthenticationRequestTypePaymentTransaction             ParsedWebhookEventAuthenticationRequestType = "PAYMENT_TRANSACTION"
	ParsedWebhookEventAuthenticationRequestTypeRecurringTransaction           ParsedWebhookEventAuthenticationRequestType = "RECURRING_TRANSACTION"
	ParsedWebhookEventAuthenticationRequestTypeSplitPayment                   ParsedWebhookEventAuthenticationRequestType = "SPLIT_PAYMENT"
	ParsedWebhookEventAuthenticationRequestTypeSplitShipment                  ParsedWebhookEventAuthenticationRequestType = "SPLIT_SHIPMENT"
)

func (r ParsedWebhookEventAuthenticationRequestType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventAuthenticationRequestTypeAddCard, ParsedWebhookEventAuthenticationRequestTypeBillingAgreement, ParsedWebhookEventAuthenticationRequestTypeDelayedShipment, ParsedWebhookEventAuthenticationRequestTypeEmvTokenCardholderVerification, ParsedWebhookEventAuthenticationRequestTypeInstallmentTransaction, ParsedWebhookEventAuthenticationRequestTypeMaintainCard, ParsedWebhookEventAuthenticationRequestTypePaymentTransaction, ParsedWebhookEventAuthenticationRequestTypeRecurringTransaction, ParsedWebhookEventAuthenticationRequestTypeSplitPayment, ParsedWebhookEventAuthenticationRequestTypeSplitShipment:
		return true
	}
	return false
}

// Indicates the outcome of the 3DS authentication process.
type ParsedWebhookEventAuthenticationResult string

const (
	ParsedWebhookEventAuthenticationResultDecline          ParsedWebhookEventAuthenticationResult = "DECLINE"
	ParsedWebhookEventAuthenticationResultSuccess          ParsedWebhookEventAuthenticationResult = "SUCCESS"
	ParsedWebhookEventAuthenticationResultPendingChallenge ParsedWebhookEventAuthenticationResult = "PENDING_CHALLENGE"
	ParsedWebhookEventAuthenticationResultPendingDecision  ParsedWebhookEventAuthenticationResult = "PENDING_DECISION"
)

func (r ParsedWebhookEventAuthenticationResult) IsKnown() bool {
	switch r {
	case ParsedWebhookEventAuthenticationResultDecline, ParsedWebhookEventAuthenticationResultSuccess, ParsedWebhookEventAuthenticationResultPendingChallenge, ParsedWebhookEventAuthenticationResultPendingDecision:
		return true
	}
	return false
}

// Indicates whether the expiration date provided by the cardholder during checkout
// matches Lithic's record of the card's expiration date.
type ParsedWebhookEventCardExpiryCheck string

const (
	ParsedWebhookEventCardExpiryCheckMatch      ParsedWebhookEventCardExpiryCheck = "MATCH"
	ParsedWebhookEventCardExpiryCheckMismatch   ParsedWebhookEventCardExpiryCheck = "MISMATCH"
	ParsedWebhookEventCardExpiryCheckNotPresent ParsedWebhookEventCardExpiryCheck = "NOT_PRESENT"
)

func (r ParsedWebhookEventCardExpiryCheck) IsKnown() bool {
	switch r {
	case ParsedWebhookEventCardExpiryCheckMatch, ParsedWebhookEventCardExpiryCheckMismatch, ParsedWebhookEventCardExpiryCheckNotPresent:
		return true
	}
	return false
}

type ParsedWebhookEventCategory string

const (
	ParsedWebhookEventCategoryAdjustment             ParsedWebhookEventCategory = "ADJUSTMENT"
	ParsedWebhookEventCategoryBalanceOrFunding       ParsedWebhookEventCategory = "BALANCE_OR_FUNDING"
	ParsedWebhookEventCategoryDerecognition          ParsedWebhookEventCategory = "DERECOGNITION"
	ParsedWebhookEventCategoryDispute                ParsedWebhookEventCategory = "DISPUTE"
	ParsedWebhookEventCategoryFee                    ParsedWebhookEventCategory = "FEE"
	ParsedWebhookEventCategoryInternal               ParsedWebhookEventCategory = "INTERNAL"
	ParsedWebhookEventCategoryReward                 ParsedWebhookEventCategory = "REWARD"
	ParsedWebhookEventCategoryProgramFunding         ParsedWebhookEventCategory = "PROGRAM_FUNDING"
	ParsedWebhookEventCategoryTransfer               ParsedWebhookEventCategory = "TRANSFER"
	ParsedWebhookEventCategoryExternalWire           ParsedWebhookEventCategory = "EXTERNAL_WIRE"
	ParsedWebhookEventCategoryExternalACH            ParsedWebhookEventCategory = "EXTERNAL_ACH"
	ParsedWebhookEventCategoryExternalCheck          ParsedWebhookEventCategory = "EXTERNAL_CHECK"
	ParsedWebhookEventCategoryExternalFednow         ParsedWebhookEventCategory = "EXTERNAL_FEDNOW"
	ParsedWebhookEventCategoryExternalRtp            ParsedWebhookEventCategory = "EXTERNAL_RTP"
	ParsedWebhookEventCategoryExternalTransfer       ParsedWebhookEventCategory = "EXTERNAL_TRANSFER"
	ParsedWebhookEventCategoryManagementFee          ParsedWebhookEventCategory = "MANAGEMENT_FEE"
	ParsedWebhookEventCategoryManagementDispute      ParsedWebhookEventCategory = "MANAGEMENT_DISPUTE"
	ParsedWebhookEventCategoryManagementReward       ParsedWebhookEventCategory = "MANAGEMENT_REWARD"
	ParsedWebhookEventCategoryManagementAdjustment   ParsedWebhookEventCategory = "MANAGEMENT_ADJUSTMENT"
	ParsedWebhookEventCategoryManagementDisbursement ParsedWebhookEventCategory = "MANAGEMENT_DISBURSEMENT"
	ParsedWebhookEventCategoryACH                    ParsedWebhookEventCategory = "ACH"
	ParsedWebhookEventCategoryCard                   ParsedWebhookEventCategory = "CARD"
)

func (r ParsedWebhookEventCategory) IsKnown() bool {
	switch r {
	case ParsedWebhookEventCategoryAdjustment, ParsedWebhookEventCategoryBalanceOrFunding, ParsedWebhookEventCategoryDerecognition, ParsedWebhookEventCategoryDispute, ParsedWebhookEventCategoryFee, ParsedWebhookEventCategoryInternal, ParsedWebhookEventCategoryReward, ParsedWebhookEventCategoryProgramFunding, ParsedWebhookEventCategoryTransfer, ParsedWebhookEventCategoryExternalWire, ParsedWebhookEventCategoryExternalACH, ParsedWebhookEventCategoryExternalCheck, ParsedWebhookEventCategoryExternalFednow, ParsedWebhookEventCategoryExternalRtp, ParsedWebhookEventCategoryExternalTransfer, ParsedWebhookEventCategoryManagementFee, ParsedWebhookEventCategoryManagementDispute, ParsedWebhookEventCategoryManagementReward, ParsedWebhookEventCategoryManagementAdjustment, ParsedWebhookEventCategoryManagementDisbursement, ParsedWebhookEventCategoryACH, ParsedWebhookEventCategoryCard:
		return true
	}
	return false
}

// Entity that orchestrates the challenge. This won't be set for authentications
// for which a decision has not yet been made (e.g. in-flight customer decisioning
// request).
type ParsedWebhookEventChallengeOrchestratedBy string

const (
	ParsedWebhookEventChallengeOrchestratedByLithic      ParsedWebhookEventChallengeOrchestratedBy = "LITHIC"
	ParsedWebhookEventChallengeOrchestratedByCustomer    ParsedWebhookEventChallengeOrchestratedBy = "CUSTOMER"
	ParsedWebhookEventChallengeOrchestratedByNoChallenge ParsedWebhookEventChallengeOrchestratedBy = "NO_CHALLENGE"
)

func (r ParsedWebhookEventChallengeOrchestratedBy) IsKnown() bool {
	switch r {
	case ParsedWebhookEventChallengeOrchestratedByLithic, ParsedWebhookEventChallengeOrchestratedByCustomer, ParsedWebhookEventChallengeOrchestratedByNoChallenge:
		return true
	}
	return false
}

// Channel in which the authentication occurs. Maps to EMV 3DS field
// `deviceChannel`.
type ParsedWebhookEventChannel string

const (
	ParsedWebhookEventChannelAppBased                  ParsedWebhookEventChannel = "APP_BASED"
	ParsedWebhookEventChannelBrowser                   ParsedWebhookEventChannel = "BROWSER"
	ParsedWebhookEventChannelThreeDSRequestorInitiated ParsedWebhookEventChannel = "THREE_DS_REQUESTOR_INITIATED"
)

func (r ParsedWebhookEventChannel) IsKnown() bool {
	switch r {
	case ParsedWebhookEventChannelAppBased, ParsedWebhookEventChannelBrowser, ParsedWebhookEventChannelThreeDSRequestorInitiated:
		return true
	}
	return false
}

// Collection resource type
type ParsedWebhookEventCollectionResourceType string

const (
	ParsedWebhookEventCollectionResourceTypeBookTransfer ParsedWebhookEventCollectionResourceType = "BOOK_TRANSFER"
	ParsedWebhookEventCollectionResourceTypePayment      ParsedWebhookEventCollectionResourceType = "PAYMENT"
)

func (r ParsedWebhookEventCollectionResourceType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventCollectionResourceTypeBookTransfer, ParsedWebhookEventCollectionResourceTypePayment:
		return true
	}
	return false
}

// Entity that made the authentication decision. This won't be set for
// authentications for which a decision has not yet been made (e.g. in-flight
// customer decisioning request).
type ParsedWebhookEventDecisionMadeBy string

const (
	ParsedWebhookEventDecisionMadeByLithicRules      ParsedWebhookEventDecisionMadeBy = "LITHIC_RULES"
	ParsedWebhookEventDecisionMadeByLithicDefault    ParsedWebhookEventDecisionMadeBy = "LITHIC_DEFAULT"
	ParsedWebhookEventDecisionMadeByCustomerRules    ParsedWebhookEventDecisionMadeBy = "CUSTOMER_RULES"
	ParsedWebhookEventDecisionMadeByCustomerEndpoint ParsedWebhookEventDecisionMadeBy = "CUSTOMER_ENDPOINT"
	ParsedWebhookEventDecisionMadeByNetwork          ParsedWebhookEventDecisionMadeBy = "NETWORK"
	ParsedWebhookEventDecisionMadeByUnknown          ParsedWebhookEventDecisionMadeBy = "UNKNOWN"
)

func (r ParsedWebhookEventDecisionMadeBy) IsKnown() bool {
	switch r {
	case ParsedWebhookEventDecisionMadeByLithicRules, ParsedWebhookEventDecisionMadeByLithicDefault, ParsedWebhookEventDecisionMadeByCustomerRules, ParsedWebhookEventDecisionMadeByCustomerEndpoint, ParsedWebhookEventDecisionMadeByNetwork, ParsedWebhookEventDecisionMadeByUnknown:
		return true
	}
	return false
}

type ParsedWebhookEventDirection string

const (
	ParsedWebhookEventDirectionCredit ParsedWebhookEventDirection = "CREDIT"
	ParsedWebhookEventDirectionDebit  ParsedWebhookEventDirection = "DEBIT"
)

func (r ParsedWebhookEventDirection) IsKnown() bool {
	switch r {
	case ParsedWebhookEventDirectionCredit, ParsedWebhookEventDirectionDebit:
		return true
	}
	return false
}

// Dispute resolution outcome
type ParsedWebhookEventDisposition string

const (
	ParsedWebhookEventDispositionWon          ParsedWebhookEventDisposition = "WON"
	ParsedWebhookEventDispositionLost         ParsedWebhookEventDisposition = "LOST"
	ParsedWebhookEventDispositionPartiallyWon ParsedWebhookEventDisposition = "PARTIALLY_WON"
	ParsedWebhookEventDispositionWithdrawn    ParsedWebhookEventDisposition = "WITHDRAWN"
	ParsedWebhookEventDispositionDenied       ParsedWebhookEventDisposition = "DENIED"
)

func (r ParsedWebhookEventDisposition) IsKnown() bool {
	switch r {
	case ParsedWebhookEventDispositionWon, ParsedWebhookEventDispositionLost, ParsedWebhookEventDispositionPartiallyWon, ParsedWebhookEventDispositionWithdrawn, ParsedWebhookEventDispositionDenied:
		return true
	}
	return false
}

// Type of documentation to be submitted for verification of an account holder
type ParsedWebhookEventDocumentType string

const (
	ParsedWebhookEventDocumentTypeDriversLicense            ParsedWebhookEventDocumentType = "DRIVERS_LICENSE"
	ParsedWebhookEventDocumentTypePassport                  ParsedWebhookEventDocumentType = "PASSPORT"
	ParsedWebhookEventDocumentTypePassportCard              ParsedWebhookEventDocumentType = "PASSPORT_CARD"
	ParsedWebhookEventDocumentTypeEinLetter                 ParsedWebhookEventDocumentType = "EIN_LETTER"
	ParsedWebhookEventDocumentTypeTaxReturn                 ParsedWebhookEventDocumentType = "TAX_RETURN"
	ParsedWebhookEventDocumentTypeOperatingAgreement        ParsedWebhookEventDocumentType = "OPERATING_AGREEMENT"
	ParsedWebhookEventDocumentTypeCertificateOfFormation    ParsedWebhookEventDocumentType = "CERTIFICATE_OF_FORMATION"
	ParsedWebhookEventDocumentTypeCertificateOfGoodStanding ParsedWebhookEventDocumentType = "CERTIFICATE_OF_GOOD_STANDING"
	ParsedWebhookEventDocumentTypeArticlesOfIncorporation   ParsedWebhookEventDocumentType = "ARTICLES_OF_INCORPORATION"
	ParsedWebhookEventDocumentTypeArticlesOfOrganization    ParsedWebhookEventDocumentType = "ARTICLES_OF_ORGANIZATION"
	ParsedWebhookEventDocumentTypeBylaws                    ParsedWebhookEventDocumentType = "BYLAWS"
	ParsedWebhookEventDocumentTypeGovernmentBusinessLicense ParsedWebhookEventDocumentType = "GOVERNMENT_BUSINESS_LICENSE"
	ParsedWebhookEventDocumentTypePartnershipAgreement      ParsedWebhookEventDocumentType = "PARTNERSHIP_AGREEMENT"
	ParsedWebhookEventDocumentTypeSs4Form                   ParsedWebhookEventDocumentType = "SS4_FORM"
	ParsedWebhookEventDocumentTypeBankStatement             ParsedWebhookEventDocumentType = "BANK_STATEMENT"
	ParsedWebhookEventDocumentTypeUtilityBillStatement      ParsedWebhookEventDocumentType = "UTILITY_BILL_STATEMENT"
	ParsedWebhookEventDocumentTypeSsnCard                   ParsedWebhookEventDocumentType = "SSN_CARD"
	ParsedWebhookEventDocumentTypeItinLetter                ParsedWebhookEventDocumentType = "ITIN_LETTER"
	ParsedWebhookEventDocumentTypeFincenBoiReport           ParsedWebhookEventDocumentType = "FINCEN_BOI_REPORT"
)

func (r ParsedWebhookEventDocumentType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventDocumentTypeDriversLicense, ParsedWebhookEventDocumentTypePassport, ParsedWebhookEventDocumentTypePassportCard, ParsedWebhookEventDocumentTypeEinLetter, ParsedWebhookEventDocumentTypeTaxReturn, ParsedWebhookEventDocumentTypeOperatingAgreement, ParsedWebhookEventDocumentTypeCertificateOfFormation, ParsedWebhookEventDocumentTypeCertificateOfGoodStanding, ParsedWebhookEventDocumentTypeArticlesOfIncorporation, ParsedWebhookEventDocumentTypeArticlesOfOrganization, ParsedWebhookEventDocumentTypeBylaws, ParsedWebhookEventDocumentTypeGovernmentBusinessLicense, ParsedWebhookEventDocumentTypePartnershipAgreement, ParsedWebhookEventDocumentTypeSs4Form, ParsedWebhookEventDocumentTypeBankStatement, ParsedWebhookEventDocumentTypeUtilityBillStatement, ParsedWebhookEventDocumentTypeSsnCard, ParsedWebhookEventDocumentTypeItinLetter, ParsedWebhookEventDocumentTypeFincenBoiReport:
		return true
	}
	return false
}

// The type of event that occurred.
type ParsedWebhookEventEventType string

const (
	ParsedWebhookEventEventTypeAccountHolderCreated                                     ParsedWebhookEventEventType = "account_holder.created"
	ParsedWebhookEventEventTypeAccountHolderUpdated                                     ParsedWebhookEventEventType = "account_holder.updated"
	ParsedWebhookEventEventTypeAccountHolderVerification                                ParsedWebhookEventEventType = "account_holder.verification"
	ParsedWebhookEventEventTypeAccountHolderDocumentUpdated                             ParsedWebhookEventEventType = "account_holder_document.updated"
	ParsedWebhookEventEventTypeCardAuthorizationApprovalRequest                         ParsedWebhookEventEventType = "card_authorization.approval_request"
	ParsedWebhookEventEventTypeDigitalWalletTokenizationApprovalRequest                 ParsedWebhookEventEventType = "digital_wallet.tokenization_approval_request"
	ParsedWebhookEventEventTypeAuthRulesBacktestReportCreated                           ParsedWebhookEventEventType = "auth_rules.backtest_report.created"
	ParsedWebhookEventEventTypeBalanceUpdated                                           ParsedWebhookEventEventType = "balance.updated"
	ParsedWebhookEventEventTypeBookTransferTransactionCreated                           ParsedWebhookEventEventType = "book_transfer_transaction.created"
	ParsedWebhookEventEventTypeBookTransferTransactionUpdated                           ParsedWebhookEventEventType = "book_transfer_transaction.updated"
	ParsedWebhookEventEventTypeCardCreated                                              ParsedWebhookEventEventType = "card.created"
	ParsedWebhookEventEventTypeCardConverted                                            ParsedWebhookEventEventType = "card.converted"
	ParsedWebhookEventEventTypeCardRenewed                                              ParsedWebhookEventEventType = "card.renewed"
	ParsedWebhookEventEventTypeCardReissued                                             ParsedWebhookEventEventType = "card.reissued"
	ParsedWebhookEventEventTypeCardShipped                                              ParsedWebhookEventEventType = "card.shipped"
	ParsedWebhookEventEventTypeCardUpdated                                              ParsedWebhookEventEventType = "card.updated"
	ParsedWebhookEventEventTypeCardTransactionUpdated                                   ParsedWebhookEventEventType = "card_transaction.updated"
	ParsedWebhookEventEventTypeCardTransactionEnhancedDataCreated                       ParsedWebhookEventEventType = "card_transaction.enhanced_data.created"
	ParsedWebhookEventEventTypeCardTransactionEnhancedDataUpdated                       ParsedWebhookEventEventType = "card_transaction.enhanced_data.updated"
	ParsedWebhookEventEventTypeDigitalWalletTokenizationResult                          ParsedWebhookEventEventType = "digital_wallet.tokenization_result"
	ParsedWebhookEventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     ParsedWebhookEventEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	ParsedWebhookEventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent ParsedWebhookEventEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	ParsedWebhookEventEventTypeDigitalWalletTokenizationUpdated                         ParsedWebhookEventEventType = "digital_wallet.tokenization_updated"
	ParsedWebhookEventEventTypeDisputeUpdated                                           ParsedWebhookEventEventType = "dispute.updated"
	ParsedWebhookEventEventTypeDisputeEvidenceUploadFailed                              ParsedWebhookEventEventType = "dispute_evidence.upload_failed"
	ParsedWebhookEventEventTypeExternalBankAccountCreated                               ParsedWebhookEventEventType = "external_bank_account.created"
	ParsedWebhookEventEventTypeExternalBankAccountUpdated                               ParsedWebhookEventEventType = "external_bank_account.updated"
	ParsedWebhookEventEventTypeExternalPaymentCreated                                   ParsedWebhookEventEventType = "external_payment.created"
	ParsedWebhookEventEventTypeExternalPaymentUpdated                                   ParsedWebhookEventEventType = "external_payment.updated"
	ParsedWebhookEventEventTypeFinancialAccountCreated                                  ParsedWebhookEventEventType = "financial_account.created"
	ParsedWebhookEventEventTypeFinancialAccountUpdated                                  ParsedWebhookEventEventType = "financial_account.updated"
	ParsedWebhookEventEventTypeFundingEventCreated                                      ParsedWebhookEventEventType = "funding_event.created"
	ParsedWebhookEventEventTypeLoanTapeCreated                                          ParsedWebhookEventEventType = "loan_tape.created"
	ParsedWebhookEventEventTypeLoanTapeUpdated                                          ParsedWebhookEventEventType = "loan_tape.updated"
	ParsedWebhookEventEventTypeManagementOperationCreated                               ParsedWebhookEventEventType = "management_operation.created"
	ParsedWebhookEventEventTypeManagementOperationUpdated                               ParsedWebhookEventEventType = "management_operation.updated"
	ParsedWebhookEventEventTypeInternalTransactionCreated                               ParsedWebhookEventEventType = "internal_transaction.created"
	ParsedWebhookEventEventTypeInternalTransactionUpdated                               ParsedWebhookEventEventType = "internal_transaction.updated"
	ParsedWebhookEventEventTypeNetworkTotalCreated                                      ParsedWebhookEventEventType = "network_total.created"
	ParsedWebhookEventEventTypeNetworkTotalUpdated                                      ParsedWebhookEventEventType = "network_total.updated"
	ParsedWebhookEventEventTypePaymentTransactionCreated                                ParsedWebhookEventEventType = "payment_transaction.created"
	ParsedWebhookEventEventTypePaymentTransactionUpdated                                ParsedWebhookEventEventType = "payment_transaction.updated"
	ParsedWebhookEventEventTypeSettlementReportUpdated                                  ParsedWebhookEventEventType = "settlement_report.updated"
	ParsedWebhookEventEventTypeStatementsCreated                                        ParsedWebhookEventEventType = "statements.created"
	ParsedWebhookEventEventTypeThreeDSAuthenticationCreated                             ParsedWebhookEventEventType = "three_ds_authentication.created"
	ParsedWebhookEventEventTypeThreeDSAuthenticationUpdated                             ParsedWebhookEventEventType = "three_ds_authentication.updated"
	ParsedWebhookEventEventTypeThreeDSAuthenticationChallenge                           ParsedWebhookEventEventType = "three_ds_authentication.challenge"
	ParsedWebhookEventEventTypeTokenizationApprovalRequest                              ParsedWebhookEventEventType = "tokenization.approval_request"
	ParsedWebhookEventEventTypeTokenizationResult                                       ParsedWebhookEventEventType = "tokenization.result"
	ParsedWebhookEventEventTypeTokenizationTwoFactorAuthenticationCode                  ParsedWebhookEventEventType = "tokenization.two_factor_authentication_code"
	ParsedWebhookEventEventTypeTokenizationTwoFactorAuthenticationCodeSent              ParsedWebhookEventEventType = "tokenization.two_factor_authentication_code_sent"
	ParsedWebhookEventEventTypeTokenizationUpdated                                      ParsedWebhookEventEventType = "tokenization.updated"
	ParsedWebhookEventEventTypeThreeDSAuthenticationApprovalRequest                     ParsedWebhookEventEventType = "three_ds_authentication.approval_request"
	ParsedWebhookEventEventTypeDisputeTransactionCreated                                ParsedWebhookEventEventType = "dispute_transaction.created"
	ParsedWebhookEventEventTypeDisputeTransactionUpdated                                ParsedWebhookEventEventType = "dispute_transaction.updated"
)

func (r ParsedWebhookEventEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventEventTypeAccountHolderCreated, ParsedWebhookEventEventTypeAccountHolderUpdated, ParsedWebhookEventEventTypeAccountHolderVerification, ParsedWebhookEventEventTypeAccountHolderDocumentUpdated, ParsedWebhookEventEventTypeCardAuthorizationApprovalRequest, ParsedWebhookEventEventTypeDigitalWalletTokenizationApprovalRequest, ParsedWebhookEventEventTypeAuthRulesBacktestReportCreated, ParsedWebhookEventEventTypeBalanceUpdated, ParsedWebhookEventEventTypeBookTransferTransactionCreated, ParsedWebhookEventEventTypeBookTransferTransactionUpdated, ParsedWebhookEventEventTypeCardCreated, ParsedWebhookEventEventTypeCardConverted, ParsedWebhookEventEventTypeCardRenewed, ParsedWebhookEventEventTypeCardReissued, ParsedWebhookEventEventTypeCardShipped, ParsedWebhookEventEventTypeCardUpdated, ParsedWebhookEventEventTypeCardTransactionUpdated, ParsedWebhookEventEventTypeCardTransactionEnhancedDataCreated, ParsedWebhookEventEventTypeCardTransactionEnhancedDataUpdated, ParsedWebhookEventEventTypeDigitalWalletTokenizationResult, ParsedWebhookEventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, ParsedWebhookEventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, ParsedWebhookEventEventTypeDigitalWalletTokenizationUpdated, ParsedWebhookEventEventTypeDisputeUpdated, ParsedWebhookEventEventTypeDisputeEvidenceUploadFailed, ParsedWebhookEventEventTypeExternalBankAccountCreated, ParsedWebhookEventEventTypeExternalBankAccountUpdated, ParsedWebhookEventEventTypeExternalPaymentCreated, ParsedWebhookEventEventTypeExternalPaymentUpdated, ParsedWebhookEventEventTypeFinancialAccountCreated, ParsedWebhookEventEventTypeFinancialAccountUpdated, ParsedWebhookEventEventTypeFundingEventCreated, ParsedWebhookEventEventTypeLoanTapeCreated, ParsedWebhookEventEventTypeLoanTapeUpdated, ParsedWebhookEventEventTypeManagementOperationCreated, ParsedWebhookEventEventTypeManagementOperationUpdated, ParsedWebhookEventEventTypeInternalTransactionCreated, ParsedWebhookEventEventTypeInternalTransactionUpdated, ParsedWebhookEventEventTypeNetworkTotalCreated, ParsedWebhookEventEventTypeNetworkTotalUpdated, ParsedWebhookEventEventTypePaymentTransactionCreated, ParsedWebhookEventEventTypePaymentTransactionUpdated, ParsedWebhookEventEventTypeSettlementReportUpdated, ParsedWebhookEventEventTypeStatementsCreated, ParsedWebhookEventEventTypeThreeDSAuthenticationCreated, ParsedWebhookEventEventTypeThreeDSAuthenticationUpdated, ParsedWebhookEventEventTypeThreeDSAuthenticationChallenge, ParsedWebhookEventEventTypeTokenizationApprovalRequest, ParsedWebhookEventEventTypeTokenizationResult, ParsedWebhookEventEventTypeTokenizationTwoFactorAuthenticationCode, ParsedWebhookEventEventTypeTokenizationTwoFactorAuthenticationCodeSent, ParsedWebhookEventEventTypeTokenizationUpdated, ParsedWebhookEventEventTypeThreeDSAuthenticationApprovalRequest, ParsedWebhookEventEventTypeDisputeTransactionCreated, ParsedWebhookEventEventTypeDisputeTransactionUpdated:
		return true
	}
	return false
}

// TRANSFER - Book Transfer Transaction
type ParsedWebhookEventFamily string

const (
	ParsedWebhookEventFamilyTransfer            ParsedWebhookEventFamily = "TRANSFER"
	ParsedWebhookEventFamilyExternalPayment     ParsedWebhookEventFamily = "EXTERNAL_PAYMENT"
	ParsedWebhookEventFamilyManagementOperation ParsedWebhookEventFamily = "MANAGEMENT_OPERATION"
	ParsedWebhookEventFamilyPayment             ParsedWebhookEventFamily = "PAYMENT"
)

func (r ParsedWebhookEventFamily) IsKnown() bool {
	switch r {
	case ParsedWebhookEventFamilyTransfer, ParsedWebhookEventFamilyExternalPayment, ParsedWebhookEventFamilyManagementOperation, ParsedWebhookEventFamilyPayment:
		return true
	}
	return false
}

// Whether Lithic decisioned on the token, and if so, what the decision was.
// APPROVED/VERIFICATION_REQUIRED/DENIED.
type ParsedWebhookEventIssuerDecision string

const (
	ParsedWebhookEventIssuerDecisionApproved             ParsedWebhookEventIssuerDecision = "APPROVED"
	ParsedWebhookEventIssuerDecisionDenied               ParsedWebhookEventIssuerDecision = "DENIED"
	ParsedWebhookEventIssuerDecisionVerificationRequired ParsedWebhookEventIssuerDecision = "VERIFICATION_REQUIRED"
)

func (r ParsedWebhookEventIssuerDecision) IsKnown() bool {
	switch r {
	case ParsedWebhookEventIssuerDecisionApproved, ParsedWebhookEventIssuerDecisionDenied, ParsedWebhookEventIssuerDecisionVerificationRequired:
		return true
	}
	return false
}

// Either PAYMENT_AUTHENTICATION or NON_PAYMENT_AUTHENTICATION. For
// NON_PAYMENT_AUTHENTICATION, additional_data and transaction fields are not
// populated.
type ParsedWebhookEventMessageCategory string

const (
	ParsedWebhookEventMessageCategoryNonPaymentAuthentication ParsedWebhookEventMessageCategory = "NON_PAYMENT_AUTHENTICATION"
	ParsedWebhookEventMessageCategoryPaymentAuthentication    ParsedWebhookEventMessageCategory = "PAYMENT_AUTHENTICATION"
)

func (r ParsedWebhookEventMessageCategory) IsKnown() bool {
	switch r {
	case ParsedWebhookEventMessageCategoryNonPaymentAuthentication, ParsedWebhookEventMessageCategoryPaymentAuthentication:
		return true
	}
	return false
}

// Transfer method
type ParsedWebhookEventMethod string

const (
	ParsedWebhookEventMethodACHNextDay ParsedWebhookEventMethod = "ACH_NEXT_DAY"
	ParsedWebhookEventMethodACHSameDay ParsedWebhookEventMethod = "ACH_SAME_DAY"
	ParsedWebhookEventMethodWire       ParsedWebhookEventMethod = "WIRE"
)

func (r ParsedWebhookEventMethod) IsKnown() bool {
	switch r {
	case ParsedWebhookEventMethodACHNextDay, ParsedWebhookEventMethodACHSameDay, ParsedWebhookEventMethodWire:
		return true
	}
	return false
}

// Card network of the authorization.
type ParsedWebhookEventNetwork string

const (
	ParsedWebhookEventNetworkAmex       ParsedWebhookEventNetwork = "AMEX"
	ParsedWebhookEventNetworkInterlink  ParsedWebhookEventNetwork = "INTERLINK"
	ParsedWebhookEventNetworkMaestro    ParsedWebhookEventNetwork = "MAESTRO"
	ParsedWebhookEventNetworkMastercard ParsedWebhookEventNetwork = "MASTERCARD"
	ParsedWebhookEventNetworkUnknown    ParsedWebhookEventNetwork = "UNKNOWN"
	ParsedWebhookEventNetworkVisa       ParsedWebhookEventNetwork = "VISA"
)

func (r ParsedWebhookEventNetwork) IsKnown() bool {
	switch r {
	case ParsedWebhookEventNetworkAmex, ParsedWebhookEventNetworkInterlink, ParsedWebhookEventNetworkMaestro, ParsedWebhookEventNetworkMastercard, ParsedWebhookEventNetworkUnknown, ParsedWebhookEventNetworkVisa:
		return true
	}
	return false
}

type ParsedWebhookEventPaymentType string

const (
	ParsedWebhookEventPaymentTypeDeposit    ParsedWebhookEventPaymentType = "DEPOSIT"
	ParsedWebhookEventPaymentTypeWithdrawal ParsedWebhookEventPaymentType = "WITHDRAWAL"
)

func (r ParsedWebhookEventPaymentType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventPaymentTypeDeposit, ParsedWebhookEventPaymentTypeWithdrawal:
		return true
	}
	return false
}

// Dispute reason:
//
//   - `ATM_CASH_MISDISPENSE`: ATM cash misdispense.
//   - `CANCELLED`: Transaction was cancelled by the customer.
//   - `DUPLICATED`: The transaction was a duplicate.
//   - `FRAUD_CARD_NOT_PRESENT`: Fraudulent transaction, card not present.
//   - `FRAUD_CARD_PRESENT`: Fraudulent transaction, card present.
//   - `FRAUD_OTHER`: Fraudulent transaction, other types such as questionable
//     merchant activity.
//   - `GOODS_SERVICES_NOT_AS_DESCRIBED`: The goods or services were not as
//     described.
//   - `GOODS_SERVICES_NOT_RECEIVED`: The goods or services were not received.
//   - `INCORRECT_AMOUNT`: The transaction amount was incorrect.
//   - `MISSING_AUTH`: The transaction was missing authorization.
//   - `OTHER`: Other reason.
//   - `PROCESSING_ERROR`: Processing error.
//   - `REFUND_NOT_PROCESSED`: The refund was not processed.
//   - `RECURRING_TRANSACTION_NOT_CANCELLED`: The recurring transaction was not
//     cancelled.
type ParsedWebhookEventReason string

const (
	ParsedWebhookEventReasonAtmCashMisdispense               ParsedWebhookEventReason = "ATM_CASH_MISDISPENSE"
	ParsedWebhookEventReasonCancelled                        ParsedWebhookEventReason = "CANCELLED"
	ParsedWebhookEventReasonDuplicated                       ParsedWebhookEventReason = "DUPLICATED"
	ParsedWebhookEventReasonFraudCardNotPresent              ParsedWebhookEventReason = "FRAUD_CARD_NOT_PRESENT"
	ParsedWebhookEventReasonFraudCardPresent                 ParsedWebhookEventReason = "FRAUD_CARD_PRESENT"
	ParsedWebhookEventReasonFraudOther                       ParsedWebhookEventReason = "FRAUD_OTHER"
	ParsedWebhookEventReasonGoodsServicesNotAsDescribed      ParsedWebhookEventReason = "GOODS_SERVICES_NOT_AS_DESCRIBED"
	ParsedWebhookEventReasonGoodsServicesNotReceived         ParsedWebhookEventReason = "GOODS_SERVICES_NOT_RECEIVED"
	ParsedWebhookEventReasonIncorrectAmount                  ParsedWebhookEventReason = "INCORRECT_AMOUNT"
	ParsedWebhookEventReasonMissingAuth                      ParsedWebhookEventReason = "MISSING_AUTH"
	ParsedWebhookEventReasonOther                            ParsedWebhookEventReason = "OTHER"
	ParsedWebhookEventReasonProcessingError                  ParsedWebhookEventReason = "PROCESSING_ERROR"
	ParsedWebhookEventReasonRecurringTransactionNotCancelled ParsedWebhookEventReason = "RECURRING_TRANSACTION_NOT_CANCELLED"
	ParsedWebhookEventReasonRefundNotProcessed               ParsedWebhookEventReason = "REFUND_NOT_PROCESSED"
)

func (r ParsedWebhookEventReason) IsKnown() bool {
	switch r {
	case ParsedWebhookEventReasonAtmCashMisdispense, ParsedWebhookEventReasonCancelled, ParsedWebhookEventReasonDuplicated, ParsedWebhookEventReasonFraudCardNotPresent, ParsedWebhookEventReasonFraudCardPresent, ParsedWebhookEventReasonFraudOther, ParsedWebhookEventReasonGoodsServicesNotAsDescribed, ParsedWebhookEventReasonGoodsServicesNotReceived, ParsedWebhookEventReasonIncorrectAmount, ParsedWebhookEventReasonMissingAuth, ParsedWebhookEventReasonOther, ParsedWebhookEventReasonProcessingError, ParsedWebhookEventReasonRecurringTransactionNotCancelled, ParsedWebhookEventReasonRefundNotProcessed:
		return true
	}
	return false
}

// Reason for the dispute resolution:
//
// - `CASE_LOST`: This case was lost at final arbitration.
// - `NETWORK_REJECTED`: Network rejected.
// - `NO_DISPUTE_RIGHTS_3DS`: No dispute rights, 3DS.
// - `NO_DISPUTE_RIGHTS_BELOW_THRESHOLD`: No dispute rights, below threshold.
// - `NO_DISPUTE_RIGHTS_CONTACTLESS`: No dispute rights, contactless.
// - `NO_DISPUTE_RIGHTS_HYBRID`: No dispute rights, hybrid.
// - `NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS`: No dispute rights, max chargebacks.
// - `NO_DISPUTE_RIGHTS_OTHER`: No dispute rights, other.
// - `PAST_FILING_DATE`: Past filing date.
// - `PREARBITRATION_REJECTED`: Prearbitration rejected.
// - `PROCESSOR_REJECTED_OTHER`: Processor rejected, other.
// - `REFUNDED`: Refunded.
// - `REFUNDED_AFTER_CHARGEBACK`: Refunded after chargeback.
// - `WITHDRAWN`: Withdrawn.
// - `WON_ARBITRATION`: Won arbitration.
// - `WON_FIRST_CHARGEBACK`: Won first chargeback.
// - `WON_PREARBITRATION`: Won prearbitration.
type ParsedWebhookEventResolutionReason string

const (
	ParsedWebhookEventResolutionReasonCaseLost                      ParsedWebhookEventResolutionReason = "CASE_LOST"
	ParsedWebhookEventResolutionReasonNetworkRejected               ParsedWebhookEventResolutionReason = "NETWORK_REJECTED"
	ParsedWebhookEventResolutionReasonNoDisputeRights3DS            ParsedWebhookEventResolutionReason = "NO_DISPUTE_RIGHTS_3DS"
	ParsedWebhookEventResolutionReasonNoDisputeRightsBelowThreshold ParsedWebhookEventResolutionReason = "NO_DISPUTE_RIGHTS_BELOW_THRESHOLD"
	ParsedWebhookEventResolutionReasonNoDisputeRightsContactless    ParsedWebhookEventResolutionReason = "NO_DISPUTE_RIGHTS_CONTACTLESS"
	ParsedWebhookEventResolutionReasonNoDisputeRightsHybrid         ParsedWebhookEventResolutionReason = "NO_DISPUTE_RIGHTS_HYBRID"
	ParsedWebhookEventResolutionReasonNoDisputeRightsMaxChargebacks ParsedWebhookEventResolutionReason = "NO_DISPUTE_RIGHTS_MAX_CHARGEBACKS"
	ParsedWebhookEventResolutionReasonNoDisputeRightsOther          ParsedWebhookEventResolutionReason = "NO_DISPUTE_RIGHTS_OTHER"
	ParsedWebhookEventResolutionReasonPastFilingDate                ParsedWebhookEventResolutionReason = "PAST_FILING_DATE"
	ParsedWebhookEventResolutionReasonPrearbitrationRejected        ParsedWebhookEventResolutionReason = "PREARBITRATION_REJECTED"
	ParsedWebhookEventResolutionReasonProcessorRejectedOther        ParsedWebhookEventResolutionReason = "PROCESSOR_REJECTED_OTHER"
	ParsedWebhookEventResolutionReasonRefunded                      ParsedWebhookEventResolutionReason = "REFUNDED"
	ParsedWebhookEventResolutionReasonRefundedAfterChargeback       ParsedWebhookEventResolutionReason = "REFUNDED_AFTER_CHARGEBACK"
	ParsedWebhookEventResolutionReasonWithdrawn                     ParsedWebhookEventResolutionReason = "WITHDRAWN"
	ParsedWebhookEventResolutionReasonWonArbitration                ParsedWebhookEventResolutionReason = "WON_ARBITRATION"
	ParsedWebhookEventResolutionReasonWonFirstChargeback            ParsedWebhookEventResolutionReason = "WON_FIRST_CHARGEBACK"
	ParsedWebhookEventResolutionReasonWonPrearbitration             ParsedWebhookEventResolutionReason = "WON_PREARBITRATION"
)

func (r ParsedWebhookEventResolutionReason) IsKnown() bool {
	switch r {
	case ParsedWebhookEventResolutionReasonCaseLost, ParsedWebhookEventResolutionReasonNetworkRejected, ParsedWebhookEventResolutionReasonNoDisputeRights3DS, ParsedWebhookEventResolutionReasonNoDisputeRightsBelowThreshold, ParsedWebhookEventResolutionReasonNoDisputeRightsContactless, ParsedWebhookEventResolutionReasonNoDisputeRightsHybrid, ParsedWebhookEventResolutionReasonNoDisputeRightsMaxChargebacks, ParsedWebhookEventResolutionReasonNoDisputeRightsOther, ParsedWebhookEventResolutionReasonPastFilingDate, ParsedWebhookEventResolutionReasonPrearbitrationRejected, ParsedWebhookEventResolutionReasonProcessorRejectedOther, ParsedWebhookEventResolutionReasonRefunded, ParsedWebhookEventResolutionReasonRefundedAfterChargeback, ParsedWebhookEventResolutionReasonWithdrawn, ParsedWebhookEventResolutionReasonWonArbitration, ParsedWebhookEventResolutionReasonWonFirstChargeback, ParsedWebhookEventResolutionReasonWonPrearbitration:
		return true
	}
	return false
}

type ParsedWebhookEventResult string

const (
	ParsedWebhookEventResultApproved                    ParsedWebhookEventResult = "APPROVED"
	ParsedWebhookEventResultDeclined                    ParsedWebhookEventResult = "DECLINED"
	ParsedWebhookEventResultAccountPaused               ParsedWebhookEventResult = "ACCOUNT_PAUSED"
	ParsedWebhookEventResultAccountStateTransactionFail ParsedWebhookEventResult = "ACCOUNT_STATE_TRANSACTION_FAIL"
	ParsedWebhookEventResultBankConnectionError         ParsedWebhookEventResult = "BANK_CONNECTION_ERROR"
	ParsedWebhookEventResultBankNotVerified             ParsedWebhookEventResult = "BANK_NOT_VERIFIED"
	ParsedWebhookEventResultCardClosed                  ParsedWebhookEventResult = "CARD_CLOSED"
	ParsedWebhookEventResultCardPaused                  ParsedWebhookEventResult = "CARD_PAUSED"
	ParsedWebhookEventResultFraudAdvice                 ParsedWebhookEventResult = "FRAUD_ADVICE"
	ParsedWebhookEventResultIgnoredTtlExpiry            ParsedWebhookEventResult = "IGNORED_TTL_EXPIRY"
	ParsedWebhookEventResultSuspectedFraud              ParsedWebhookEventResult = "SUSPECTED_FRAUD"
	ParsedWebhookEventResultInactiveAccount             ParsedWebhookEventResult = "INACTIVE_ACCOUNT"
	ParsedWebhookEventResultIncorrectPin                ParsedWebhookEventResult = "INCORRECT_PIN"
	ParsedWebhookEventResultInvalidCardDetails          ParsedWebhookEventResult = "INVALID_CARD_DETAILS"
	ParsedWebhookEventResultInsufficientFunds           ParsedWebhookEventResult = "INSUFFICIENT_FUNDS"
	ParsedWebhookEventResultInsufficientFundsPreload    ParsedWebhookEventResult = "INSUFFICIENT_FUNDS_PRELOAD"
	ParsedWebhookEventResultInvalidTransaction          ParsedWebhookEventResult = "INVALID_TRANSACTION"
	ParsedWebhookEventResultMerchantBlacklist           ParsedWebhookEventResult = "MERCHANT_BLACKLIST"
	ParsedWebhookEventResultOriginalNotFound            ParsedWebhookEventResult = "ORIGINAL_NOT_FOUND"
	ParsedWebhookEventResultPreviouslyCompleted         ParsedWebhookEventResult = "PREVIOUSLY_COMPLETED"
	ParsedWebhookEventResultSingleUseRecharged          ParsedWebhookEventResult = "SINGLE_USE_RECHARGED"
	ParsedWebhookEventResultSwitchInoperativeAdvice     ParsedWebhookEventResult = "SWITCH_INOPERATIVE_ADVICE"
	ParsedWebhookEventResultUnauthorizedMerchant        ParsedWebhookEventResult = "UNAUTHORIZED_MERCHANT"
	ParsedWebhookEventResultUnknownHostTimeout          ParsedWebhookEventResult = "UNKNOWN_HOST_TIMEOUT"
	ParsedWebhookEventResultUserTransactionLimit        ParsedWebhookEventResult = "USER_TRANSACTION_LIMIT"
)

func (r ParsedWebhookEventResult) IsKnown() bool {
	switch r {
	case ParsedWebhookEventResultApproved, ParsedWebhookEventResultDeclined, ParsedWebhookEventResultAccountPaused, ParsedWebhookEventResultAccountStateTransactionFail, ParsedWebhookEventResultBankConnectionError, ParsedWebhookEventResultBankNotVerified, ParsedWebhookEventResultCardClosed, ParsedWebhookEventResultCardPaused, ParsedWebhookEventResultFraudAdvice, ParsedWebhookEventResultIgnoredTtlExpiry, ParsedWebhookEventResultSuspectedFraud, ParsedWebhookEventResultInactiveAccount, ParsedWebhookEventResultIncorrectPin, ParsedWebhookEventResultInvalidCardDetails, ParsedWebhookEventResultInsufficientFunds, ParsedWebhookEventResultInsufficientFundsPreload, ParsedWebhookEventResultInvalidTransaction, ParsedWebhookEventResultMerchantBlacklist, ParsedWebhookEventResultOriginalNotFound, ParsedWebhookEventResultPreviouslyCompleted, ParsedWebhookEventResultSingleUseRecharged, ParsedWebhookEventResultSwitchInoperativeAdvice, ParsedWebhookEventResultUnauthorizedMerchant, ParsedWebhookEventResultUnknownHostTimeout, ParsedWebhookEventResultUserTransactionLimit:
		return true
	}
	return false
}

// The specific shipping method used to ship the card.
type ParsedWebhookEventShippingMethod string

const (
	ParsedWebhookEventShippingMethodExUsExpeditedWithTracking             ParsedWebhookEventShippingMethod = "Ex-US expedited with tracking"
	ParsedWebhookEventShippingMethodExUsStandardWithTracking              ParsedWebhookEventShippingMethod = "Ex-US standard with tracking"
	ParsedWebhookEventShippingMethodExUsStandardWithoutTracking           ParsedWebhookEventShippingMethod = "Ex-US standard without tracking"
	ParsedWebhookEventShippingMethodFedEx2Days                            ParsedWebhookEventShippingMethod = "FedEx 2 days"
	ParsedWebhookEventShippingMethodFedExExpress                          ParsedWebhookEventShippingMethod = "FedEx express"
	ParsedWebhookEventShippingMethodFedExOvernight                        ParsedWebhookEventShippingMethod = "FedEx overnight"
	ParsedWebhookEventShippingMethodUspsPriority                          ParsedWebhookEventShippingMethod = "USPS priority"
	ParsedWebhookEventShippingMethodUspsWithTracking                      ParsedWebhookEventShippingMethod = "USPS with tracking"
	ParsedWebhookEventShippingMethodUspsWithoutTrackingEnvelope           ParsedWebhookEventShippingMethod = "USPS without tracking envelope"
	ParsedWebhookEventShippingMethodUspsWithoutTrackingEnvelopeNonMachine ParsedWebhookEventShippingMethod = "USPS without tracking envelope non-machine"
	ParsedWebhookEventShippingMethodUspsWithoutTrackingFlat               ParsedWebhookEventShippingMethod = "USPS without tracking flat"
)

func (r ParsedWebhookEventShippingMethod) IsKnown() bool {
	switch r {
	case ParsedWebhookEventShippingMethodExUsExpeditedWithTracking, ParsedWebhookEventShippingMethodExUsStandardWithTracking, ParsedWebhookEventShippingMethodExUsStandardWithoutTracking, ParsedWebhookEventShippingMethodFedEx2Days, ParsedWebhookEventShippingMethodFedExExpress, ParsedWebhookEventShippingMethodFedExOvernight, ParsedWebhookEventShippingMethodUspsPriority, ParsedWebhookEventShippingMethodUspsWithTracking, ParsedWebhookEventShippingMethodUspsWithoutTrackingEnvelope, ParsedWebhookEventShippingMethodUspsWithoutTrackingEnvelopeNonMachine, ParsedWebhookEventShippingMethodUspsWithoutTrackingFlat:
		return true
	}
	return false
}

// Transaction source
type ParsedWebhookEventSource string

const (
	ParsedWebhookEventSourceLithic   ParsedWebhookEventSource = "LITHIC"
	ParsedWebhookEventSourceExternal ParsedWebhookEventSource = "EXTERNAL"
	ParsedWebhookEventSourceCustomer ParsedWebhookEventSource = "CUSTOMER"
)

func (r ParsedWebhookEventSource) IsKnown() bool {
	switch r {
	case ParsedWebhookEventSourceLithic, ParsedWebhookEventSourceExternal, ParsedWebhookEventSourceCustomer:
		return true
	}
	return false
}

type ParsedWebhookEventStatementType string

const (
	ParsedWebhookEventStatementTypeInitial   ParsedWebhookEventStatementType = "INITIAL"
	ParsedWebhookEventStatementTypePeriodEnd ParsedWebhookEventStatementType = "PERIOD_END"
	ParsedWebhookEventStatementTypeFinal     ParsedWebhookEventStatementType = "FINAL"
)

func (r ParsedWebhookEventStatementType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventStatementTypeInitial, ParsedWebhookEventStatementTypePeriodEnd, ParsedWebhookEventStatementTypeFinal:
		return true
	}
	return false
}

// The status of the account_holder that was created.
type ParsedWebhookEventStatus string

const (
	ParsedWebhookEventStatusAccepted                     ParsedWebhookEventStatus = "ACCEPTED"
	ParsedWebhookEventStatusPendingReview                ParsedWebhookEventStatus = "PENDING_REVIEW"
	ParsedWebhookEventStatusRejected                     ParsedWebhookEventStatus = "REJECTED"
	ParsedWebhookEventStatusAuthorization                ParsedWebhookEventStatus = "AUTHORIZATION"
	ParsedWebhookEventStatusCreditAuthorization          ParsedWebhookEventStatus = "CREDIT_AUTHORIZATION"
	ParsedWebhookEventStatusFinancialAuthorization       ParsedWebhookEventStatus = "FINANCIAL_AUTHORIZATION"
	ParsedWebhookEventStatusFinancialCreditAuthorization ParsedWebhookEventStatus = "FINANCIAL_CREDIT_AUTHORIZATION"
	ParsedWebhookEventStatusBalanceInquiry               ParsedWebhookEventStatus = "BALANCE_INQUIRY"
	ParsedWebhookEventStatusPending                      ParsedWebhookEventStatus = "PENDING"
	ParsedWebhookEventStatusSettled                      ParsedWebhookEventStatus = "SETTLED"
	ParsedWebhookEventStatusDeclined                     ParsedWebhookEventStatus = "DECLINED"
	ParsedWebhookEventStatusReversed                     ParsedWebhookEventStatus = "REVERSED"
	ParsedWebhookEventStatusCanceled                     ParsedWebhookEventStatus = "CANCELED"
	ParsedWebhookEventStatusReturned                     ParsedWebhookEventStatus = "RETURNED"
	ParsedWebhookEventStatusExpired                      ParsedWebhookEventStatus = "EXPIRED"
	ParsedWebhookEventStatusVoided                       ParsedWebhookEventStatus = "VOIDED"
	ParsedWebhookEventStatusArbitration                  ParsedWebhookEventStatus = "ARBITRATION"
	ParsedWebhookEventStatusCaseClosed                   ParsedWebhookEventStatus = "CASE_CLOSED"
	ParsedWebhookEventStatusCaseWon                      ParsedWebhookEventStatus = "CASE_WON"
	ParsedWebhookEventStatusNew                          ParsedWebhookEventStatus = "NEW"
	ParsedWebhookEventStatusPendingCustomer              ParsedWebhookEventStatus = "PENDING_CUSTOMER"
	ParsedWebhookEventStatusPrearbitration               ParsedWebhookEventStatus = "PREARBITRATION"
	ParsedWebhookEventStatusRepresentment                ParsedWebhookEventStatus = "REPRESENTMENT"
	ParsedWebhookEventStatusSubmitted                    ParsedWebhookEventStatus = "SUBMITTED"
	ParsedWebhookEventStatusOpen                         ParsedWebhookEventStatus = "OPEN"
	ParsedWebhookEventStatusClosed                       ParsedWebhookEventStatus = "CLOSED"
	ParsedWebhookEventStatusSuspended                    ParsedWebhookEventStatus = "SUSPENDED"
)

func (r ParsedWebhookEventStatus) IsKnown() bool {
	switch r {
	case ParsedWebhookEventStatusAccepted, ParsedWebhookEventStatusPendingReview, ParsedWebhookEventStatusRejected, ParsedWebhookEventStatusAuthorization, ParsedWebhookEventStatusCreditAuthorization, ParsedWebhookEventStatusFinancialAuthorization, ParsedWebhookEventStatusFinancialCreditAuthorization, ParsedWebhookEventStatusBalanceInquiry, ParsedWebhookEventStatusPending, ParsedWebhookEventStatusSettled, ParsedWebhookEventStatusDeclined, ParsedWebhookEventStatusReversed, ParsedWebhookEventStatusCanceled, ParsedWebhookEventStatusReturned, ParsedWebhookEventStatusExpired, ParsedWebhookEventStatusVoided, ParsedWebhookEventStatusArbitration, ParsedWebhookEventStatusCaseClosed, ParsedWebhookEventStatusCaseWon, ParsedWebhookEventStatusNew, ParsedWebhookEventStatusPendingCustomer, ParsedWebhookEventStatusPrearbitration, ParsedWebhookEventStatusRepresentment, ParsedWebhookEventStatusSubmitted, ParsedWebhookEventStatusOpen, ParsedWebhookEventStatusClosed, ParsedWebhookEventStatusSuspended:
		return true
	}
	return false
}

// Substatus for the financial account
type ParsedWebhookEventSubstatus string

const (
	ParsedWebhookEventSubstatusChargedOffDelinquent ParsedWebhookEventSubstatus = "CHARGED_OFF_DELINQUENT"
	ParsedWebhookEventSubstatusChargedOffFraud      ParsedWebhookEventSubstatus = "CHARGED_OFF_FRAUD"
	ParsedWebhookEventSubstatusEndUserRequest       ParsedWebhookEventSubstatus = "END_USER_REQUEST"
	ParsedWebhookEventSubstatusBankRequest          ParsedWebhookEventSubstatus = "BANK_REQUEST"
	ParsedWebhookEventSubstatusDelinquent           ParsedWebhookEventSubstatus = "DELINQUENT"
)

func (r ParsedWebhookEventSubstatus) IsKnown() bool {
	switch r {
	case ParsedWebhookEventSubstatusChargedOffDelinquent, ParsedWebhookEventSubstatusChargedOffFraud, ParsedWebhookEventSubstatusEndUserRequest, ParsedWebhookEventSubstatusBankRequest, ParsedWebhookEventSubstatusDelinquent:
		return true
	}
	return false
}

// Indicates whether a challenge is requested for this transaction
//
//   - `NO_PREFERENCE` - No Preference
//   - `NO_CHALLENGE_REQUESTED` - No Challenge Requested
//   - `CHALLENGE_PREFERENCE` - Challenge requested (3DS Requestor preference)
//   - `CHALLENGE_MANDATE` - Challenge requested (Mandate)
//   - `NO_CHALLENGE_RISK_ALREADY_ASSESSED` - No Challenge requested (Transactional
//     risk analysis is already performed)
//   - `DATA_SHARE_ONLY` - No Challenge requested (Data Share Only)
//   - `OTHER` - Other indicators not captured by above. These are rarely used
type ParsedWebhookEventThreeDSRequestorChallengeIndicator string

const (
	ParsedWebhookEventThreeDSRequestorChallengeIndicatorNoPreference                   ParsedWebhookEventThreeDSRequestorChallengeIndicator = "NO_PREFERENCE"
	ParsedWebhookEventThreeDSRequestorChallengeIndicatorNoChallengeRequested           ParsedWebhookEventThreeDSRequestorChallengeIndicator = "NO_CHALLENGE_REQUESTED"
	ParsedWebhookEventThreeDSRequestorChallengeIndicatorChallengePreference            ParsedWebhookEventThreeDSRequestorChallengeIndicator = "CHALLENGE_PREFERENCE"
	ParsedWebhookEventThreeDSRequestorChallengeIndicatorChallengeMandate               ParsedWebhookEventThreeDSRequestorChallengeIndicator = "CHALLENGE_MANDATE"
	ParsedWebhookEventThreeDSRequestorChallengeIndicatorNoChallengeRiskAlreadyAssessed ParsedWebhookEventThreeDSRequestorChallengeIndicator = "NO_CHALLENGE_RISK_ALREADY_ASSESSED"
	ParsedWebhookEventThreeDSRequestorChallengeIndicatorDataShareOnly                  ParsedWebhookEventThreeDSRequestorChallengeIndicator = "DATA_SHARE_ONLY"
	ParsedWebhookEventThreeDSRequestorChallengeIndicatorOther                          ParsedWebhookEventThreeDSRequestorChallengeIndicator = "OTHER"
)

func (r ParsedWebhookEventThreeDSRequestorChallengeIndicator) IsKnown() bool {
	switch r {
	case ParsedWebhookEventThreeDSRequestorChallengeIndicatorNoPreference, ParsedWebhookEventThreeDSRequestorChallengeIndicatorNoChallengeRequested, ParsedWebhookEventThreeDSRequestorChallengeIndicatorChallengePreference, ParsedWebhookEventThreeDSRequestorChallengeIndicatorChallengeMandate, ParsedWebhookEventThreeDSRequestorChallengeIndicatorNoChallengeRiskAlreadyAssessed, ParsedWebhookEventThreeDSRequestorChallengeIndicatorDataShareOnly, ParsedWebhookEventThreeDSRequestorChallengeIndicatorOther:
		return true
	}
	return false
}

// Type of 3DS Requestor Initiated (3RI) request — i.e., a 3DS authentication that
// takes place at the initiation of the merchant rather than the cardholder. The
// most common example of this is where a merchant is authenticating before billing
// for a recurring transaction such as a pay TV subscription or a utility bill.
// Maps to EMV 3DS field `threeRIInd`.
type ParsedWebhookEventThreeRiRequestType string

const (
	ParsedWebhookEventThreeRiRequestTypeAccountVerification         ParsedWebhookEventThreeRiRequestType = "ACCOUNT_VERIFICATION"
	ParsedWebhookEventThreeRiRequestTypeAddCard                     ParsedWebhookEventThreeRiRequestType = "ADD_CARD"
	ParsedWebhookEventThreeRiRequestTypeBillingAgreement            ParsedWebhookEventThreeRiRequestType = "BILLING_AGREEMENT"
	ParsedWebhookEventThreeRiRequestTypeCardSecurityCodeStatusCheck ParsedWebhookEventThreeRiRequestType = "CARD_SECURITY_CODE_STATUS_CHECK"
	ParsedWebhookEventThreeRiRequestTypeDelayedShipment             ParsedWebhookEventThreeRiRequestType = "DELAYED_SHIPMENT"
	ParsedWebhookEventThreeRiRequestTypeDeviceBindingStatusCheck    ParsedWebhookEventThreeRiRequestType = "DEVICE_BINDING_STATUS_CHECK"
	ParsedWebhookEventThreeRiRequestTypeInstallmentTransaction      ParsedWebhookEventThreeRiRequestType = "INSTALLMENT_TRANSACTION"
	ParsedWebhookEventThreeRiRequestTypeMailOrder                   ParsedWebhookEventThreeRiRequestType = "MAIL_ORDER"
	ParsedWebhookEventThreeRiRequestTypeMaintainCardInfo            ParsedWebhookEventThreeRiRequestType = "MAINTAIN_CARD_INFO"
	ParsedWebhookEventThreeRiRequestTypeOtherPayment                ParsedWebhookEventThreeRiRequestType = "OTHER_PAYMENT"
	ParsedWebhookEventThreeRiRequestTypeRecurringTransaction        ParsedWebhookEventThreeRiRequestType = "RECURRING_TRANSACTION"
	ParsedWebhookEventThreeRiRequestTypeSplitPayment                ParsedWebhookEventThreeRiRequestType = "SPLIT_PAYMENT"
	ParsedWebhookEventThreeRiRequestTypeSplitShipment               ParsedWebhookEventThreeRiRequestType = "SPLIT_SHIPMENT"
	ParsedWebhookEventThreeRiRequestTypeTelephoneOrder              ParsedWebhookEventThreeRiRequestType = "TELEPHONE_ORDER"
	ParsedWebhookEventThreeRiRequestTypeTopUp                       ParsedWebhookEventThreeRiRequestType = "TOP_UP"
	ParsedWebhookEventThreeRiRequestTypeTrustListStatusCheck        ParsedWebhookEventThreeRiRequestType = "TRUST_LIST_STATUS_CHECK"
)

func (r ParsedWebhookEventThreeRiRequestType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventThreeRiRequestTypeAccountVerification, ParsedWebhookEventThreeRiRequestTypeAddCard, ParsedWebhookEventThreeRiRequestTypeBillingAgreement, ParsedWebhookEventThreeRiRequestTypeCardSecurityCodeStatusCheck, ParsedWebhookEventThreeRiRequestTypeDelayedShipment, ParsedWebhookEventThreeRiRequestTypeDeviceBindingStatusCheck, ParsedWebhookEventThreeRiRequestTypeInstallmentTransaction, ParsedWebhookEventThreeRiRequestTypeMailOrder, ParsedWebhookEventThreeRiRequestTypeMaintainCardInfo, ParsedWebhookEventThreeRiRequestTypeOtherPayment, ParsedWebhookEventThreeRiRequestTypeRecurringTransaction, ParsedWebhookEventThreeRiRequestTypeSplitPayment, ParsedWebhookEventThreeRiRequestTypeSplitShipment, ParsedWebhookEventThreeRiRequestTypeTelephoneOrder, ParsedWebhookEventThreeRiRequestTypeTopUp, ParsedWebhookEventThreeRiRequestTypeTrustListStatusCheck:
		return true
	}
	return false
}

// The channel through which the tokenization was made.
type ParsedWebhookEventTokenizationChannel string

const (
	ParsedWebhookEventTokenizationChannelDigitalWallet ParsedWebhookEventTokenizationChannel = "DIGITAL_WALLET"
	ParsedWebhookEventTokenizationChannelMerchant      ParsedWebhookEventTokenizationChannel = "MERCHANT"
)

func (r ParsedWebhookEventTokenizationChannel) IsKnown() bool {
	switch r {
	case ParsedWebhookEventTokenizationChannelDigitalWallet, ParsedWebhookEventTokenizationChannelMerchant:
		return true
	}
	return false
}

// The source of the tokenization.
type ParsedWebhookEventTokenizationSource string

const (
	ParsedWebhookEventTokenizationSourceAccountOnFile   ParsedWebhookEventTokenizationSource = "ACCOUNT_ON_FILE"
	ParsedWebhookEventTokenizationSourceContactlessTap  ParsedWebhookEventTokenizationSource = "CONTACTLESS_TAP"
	ParsedWebhookEventTokenizationSourceManualProvision ParsedWebhookEventTokenizationSource = "MANUAL_PROVISION"
	ParsedWebhookEventTokenizationSourcePushProvision   ParsedWebhookEventTokenizationSource = "PUSH_PROVISION"
	ParsedWebhookEventTokenizationSourceToken           ParsedWebhookEventTokenizationSource = "TOKEN"
	ParsedWebhookEventTokenizationSourceUnknown         ParsedWebhookEventTokenizationSource = "UNKNOWN"
)

func (r ParsedWebhookEventTokenizationSource) IsKnown() bool {
	switch r {
	case ParsedWebhookEventTokenizationSourceAccountOnFile, ParsedWebhookEventTokenizationSourceContactlessTap, ParsedWebhookEventTokenizationSourceManualProvision, ParsedWebhookEventTokenizationSourcePushProvision, ParsedWebhookEventTokenizationSourceToken, ParsedWebhookEventTokenizationSourceUnknown:
		return true
	}
	return false
}

// The entity that initiated the transaction.
type ParsedWebhookEventTransactionInitiator string

const (
	ParsedWebhookEventTransactionInitiatorCardholder ParsedWebhookEventTransactionInitiator = "CARDHOLDER"
	ParsedWebhookEventTransactionInitiatorMerchant   ParsedWebhookEventTransactionInitiator = "MERCHANT"
	ParsedWebhookEventTransactionInitiatorUnknown    ParsedWebhookEventTransactionInitiator = "UNKNOWN"
)

func (r ParsedWebhookEventTransactionInitiator) IsKnown() bool {
	switch r {
	case ParsedWebhookEventTransactionInitiatorCardholder, ParsedWebhookEventTransactionInitiatorMerchant, ParsedWebhookEventTransactionInitiatorUnknown:
		return true
	}
	return false
}

// Account Type
type ParsedWebhookEventType string

const (
	ParsedWebhookEventTypeChecking                   ParsedWebhookEventType = "CHECKING"
	ParsedWebhookEventTypeSavings                    ParsedWebhookEventType = "SAVINGS"
	ParsedWebhookEventTypeIssuing                    ParsedWebhookEventType = "ISSUING"
	ParsedWebhookEventTypeReserve                    ParsedWebhookEventType = "RESERVE"
	ParsedWebhookEventTypeOperating                  ParsedWebhookEventType = "OPERATING"
	ParsedWebhookEventTypeChargedOffFees             ParsedWebhookEventType = "CHARGED_OFF_FEES"
	ParsedWebhookEventTypeChargedOffInterest         ParsedWebhookEventType = "CHARGED_OFF_INTEREST"
	ParsedWebhookEventTypeChargedOffPrincipal        ParsedWebhookEventType = "CHARGED_OFF_PRINCIPAL"
	ParsedWebhookEventTypeSecurity                   ParsedWebhookEventType = "SECURITY"
	ParsedWebhookEventTypeProgramReceivables         ParsedWebhookEventType = "PROGRAM_RECEIVABLES"
	ParsedWebhookEventTypeCollection                 ParsedWebhookEventType = "COLLECTION"
	ParsedWebhookEventTypeProgramBankAccountsPayable ParsedWebhookEventType = "PROGRAM_BANK_ACCOUNTS_PAYABLE"
	ParsedWebhookEventTypeOriginationCredit          ParsedWebhookEventType = "ORIGINATION_CREDIT"
	ParsedWebhookEventTypeOriginationDebit           ParsedWebhookEventType = "ORIGINATION_DEBIT"
	ParsedWebhookEventTypeReceiptCredit              ParsedWebhookEventType = "RECEIPT_CREDIT"
	ParsedWebhookEventTypeReceiptDebit               ParsedWebhookEventType = "RECEIPT_DEBIT"
	ParsedWebhookEventTypeWireInboundPayment         ParsedWebhookEventType = "WIRE_INBOUND_PAYMENT"
	ParsedWebhookEventTypeWireInboundAdmin           ParsedWebhookEventType = "WIRE_INBOUND_ADMIN"
	ParsedWebhookEventTypeWireOutboundPayment        ParsedWebhookEventType = "WIRE_OUTBOUND_PAYMENT"
	ParsedWebhookEventTypeWireOutboundAdmin          ParsedWebhookEventType = "WIRE_OUTBOUND_ADMIN"
	ParsedWebhookEventTypeWireInboundDrawdownRequest ParsedWebhookEventType = "WIRE_INBOUND_DRAWDOWN_REQUEST"
)

func (r ParsedWebhookEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventTypeChecking, ParsedWebhookEventTypeSavings, ParsedWebhookEventTypeIssuing, ParsedWebhookEventTypeReserve, ParsedWebhookEventTypeOperating, ParsedWebhookEventTypeChargedOffFees, ParsedWebhookEventTypeChargedOffInterest, ParsedWebhookEventTypeChargedOffPrincipal, ParsedWebhookEventTypeSecurity, ParsedWebhookEventTypeProgramReceivables, ParsedWebhookEventTypeCollection, ParsedWebhookEventTypeProgramBankAccountsPayable, ParsedWebhookEventTypeOriginationCredit, ParsedWebhookEventTypeOriginationDebit, ParsedWebhookEventTypeReceiptCredit, ParsedWebhookEventTypeReceiptDebit, ParsedWebhookEventTypeWireInboundPayment, ParsedWebhookEventTypeWireInboundAdmin, ParsedWebhookEventTypeWireOutboundPayment, ParsedWebhookEventTypeWireOutboundAdmin, ParsedWebhookEventTypeWireInboundDrawdownRequest:
		return true
	}
	return false
}

// Upload status types:
//
// - `DELETED` - Evidence was deleted.
// - `ERROR` - Evidence upload failed.
// - `PENDING` - Evidence is pending upload.
// - `REJECTED` - Evidence was rejected.
// - `UPLOADED` - Evidence was uploaded.
type ParsedWebhookEventUploadStatus string

const (
	ParsedWebhookEventUploadStatusDeleted  ParsedWebhookEventUploadStatus = "DELETED"
	ParsedWebhookEventUploadStatusError    ParsedWebhookEventUploadStatus = "ERROR"
	ParsedWebhookEventUploadStatusPending  ParsedWebhookEventUploadStatus = "PENDING"
	ParsedWebhookEventUploadStatusRejected ParsedWebhookEventUploadStatus = "REJECTED"
	ParsedWebhookEventUploadStatusUploaded ParsedWebhookEventUploadStatus = "UPLOADED"
)

func (r ParsedWebhookEventUploadStatus) IsKnown() bool {
	switch r {
	case ParsedWebhookEventUploadStatusDeleted, ParsedWebhookEventUploadStatusError, ParsedWebhookEventUploadStatusPending, ParsedWebhookEventUploadStatusRejected, ParsedWebhookEventUploadStatusUploaded:
		return true
	}
	return false
}

// Verification State
type ParsedWebhookEventVerificationState string

const (
	ParsedWebhookEventVerificationStatePending            ParsedWebhookEventVerificationState = "PENDING"
	ParsedWebhookEventVerificationStateEnabled            ParsedWebhookEventVerificationState = "ENABLED"
	ParsedWebhookEventVerificationStateFailedVerification ParsedWebhookEventVerificationState = "FAILED_VERIFICATION"
	ParsedWebhookEventVerificationStateInsufficientFunds  ParsedWebhookEventVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ParsedWebhookEventVerificationState) IsKnown() bool {
	switch r {
	case ParsedWebhookEventVerificationStatePending, ParsedWebhookEventVerificationStateEnabled, ParsedWebhookEventVerificationStateFailedVerification, ParsedWebhookEventVerificationStateInsufficientFunds:
		return true
	}
	return false
}
