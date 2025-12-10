// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/option"
)

// ThreeDSService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreeDSService] method instead.
type ThreeDSService struct {
	Options        []option.RequestOption
	Authentication *ThreeDSAuthenticationService
	Decisioning    *ThreeDSDecisioningService
}

// NewThreeDSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewThreeDSService(opts ...option.RequestOption) (r *ThreeDSService) {
	r = &ThreeDSService{}
	r.Options = opts
	r.Authentication = NewThreeDSAuthenticationService(opts...)
	r.Decisioning = NewThreeDSDecisioningService(opts...)
	return
}

// Represents a 3DS authentication
type ThreeDSAuthentication struct {
	// Globally unique identifier for the 3DS authentication. Permitted values:
	// 36-digit version 4 UUID (including hyphens).
	Token string `json:"token,required" format:"uuid"`
	// Type of account/card that is being used for the transaction. Maps to EMV 3DS
	// field `acctType`.
	AccountType ThreeDSAuthenticationAccountType `json:"account_type,required,nullable"`
	// Indicates the outcome of the 3DS authentication process.
	AuthenticationResult ThreeDSAuthenticationAuthenticationResult `json:"authentication_result,required"`
	// Indicates whether the expiration date provided by the cardholder during checkout
	// matches Lithic's record of the card's expiration date.
	CardExpiryCheck ThreeDSAuthenticationCardExpiryCheck `json:"card_expiry_check,required"`
	// Globally unique identifier for the card on which the 3DS authentication has
	// occurred. Permitted values: 36-digit version 4 UUID (including hyphens).
	CardToken string `json:"card_token,required" format:"uuid"`
	// Object containing data about the cardholder provided during the transaction.
	Cardholder ThreeDSAuthenticationCardholder `json:"cardholder,required"`
	// Channel in which the authentication occurs. Maps to EMV 3DS field
	// `deviceChannel`.
	Channel ThreeDSAuthenticationChannel `json:"channel,required"`
	// Date and time when the authentication was created in Lithic's system. Permitted
	// values: Date string in the ISO 8601 format yyyy-MM-dd'T'hh:mm:ssZ.
	Created time.Time `json:"created,required" format:"date-time"`
	// Object containing data about the merchant involved in the e-commerce
	// transaction.
	Merchant ThreeDSAuthenticationMerchant `json:"merchant,required"`
	// Either PAYMENT_AUTHENTICATION or NON_PAYMENT_AUTHENTICATION. For
	// NON_PAYMENT_AUTHENTICATION, additional_data and transaction fields are not
	// populated.
	MessageCategory ThreeDSAuthenticationMessageCategory `json:"message_category,required"`
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
	ThreeDSRequestorChallengeIndicator ThreeDSAuthenticationThreeDSRequestorChallengeIndicator `json:"three_ds_requestor_challenge_indicator,required"`
	// Object containing additional data about the 3DS request that is beyond the EMV
	// 3DS standard spec (e.g., specific fields that only certain card networks send
	// but are not required across all 3DS requests).
	AdditionalData ThreeDSAuthenticationAdditionalData `json:"additional_data,nullable"`
	// Object containing data about the app used in the e-commerce transaction. Present
	// if the channel is 'APP_BASED'.
	App ThreeDSAuthenticationApp `json:"app,nullable"`
	// Type of authentication request - i.e., the type of transaction or interaction is
	// causing the merchant to request an authentication. Maps to EMV 3DS field
	// `threeDSRequestorAuthenticationInd`.
	AuthenticationRequestType ThreeDSAuthenticationAuthenticationRequestType `json:"authentication_request_type,nullable"`
	// Object containing data about the browser used in the e-commerce transaction.
	// Present if the channel is 'BROWSER'.
	Browser ThreeDSAuthenticationBrowser `json:"browser,nullable"`
	// Metadata about the challenge method and delivery. Only present when a challenge
	// is triggered.
	ChallengeMetadata ThreeDSAuthenticationChallengeMetadata `json:"challenge_metadata,nullable"`
	// Entity that orchestrates the challenge. This won't be set for authentications
	// for which a decision has not yet been made (e.g. in-flight customer decisioning
	// request).
	ChallengeOrchestratedBy ThreeDSAuthenticationChallengeOrchestratedBy `json:"challenge_orchestrated_by,nullable"`
	// Entity that made the authentication decision. This won't be set for
	// authentications for which a decision has not yet been made (e.g. in-flight
	// customer decisioning request).
	DecisionMadeBy ThreeDSAuthenticationDecisionMadeBy `json:"decision_made_by,nullable"`
	// Type of 3DS Requestor Initiated (3RI) request — i.e., a 3DS authentication that
	// takes place at the initiation of the merchant rather than the cardholder. The
	// most common example of this is where a merchant is authenticating before billing
	// for a recurring transaction such as a pay TV subscription or a utility bill.
	// Maps to EMV 3DS field `threeRIInd`.
	ThreeRiRequestType ThreeDSAuthenticationThreeRiRequestType `json:"three_ri_request_type,nullable"`
	// Object containing data about the e-commerce transaction for which the merchant
	// is requesting authentication.
	Transaction ThreeDSAuthenticationTransaction `json:"transaction,nullable"`
	JSON        threeDSAuthenticationJSON        `json:"-"`
}

// threeDSAuthenticationJSON contains the JSON metadata for the struct
// [ThreeDSAuthentication]
type threeDSAuthenticationJSON struct {
	Token                              apijson.Field
	AccountType                        apijson.Field
	AuthenticationResult               apijson.Field
	CardExpiryCheck                    apijson.Field
	CardToken                          apijson.Field
	Cardholder                         apijson.Field
	Channel                            apijson.Field
	Created                            apijson.Field
	Merchant                           apijson.Field
	MessageCategory                    apijson.Field
	ThreeDSRequestorChallengeIndicator apijson.Field
	AdditionalData                     apijson.Field
	App                                apijson.Field
	AuthenticationRequestType          apijson.Field
	Browser                            apijson.Field
	ChallengeMetadata                  apijson.Field
	ChallengeOrchestratedBy            apijson.Field
	DecisionMadeBy                     apijson.Field
	ThreeRiRequestType                 apijson.Field
	Transaction                        apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ThreeDSAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r ThreeDSAuthentication) implementsParsedWebhookEvent() {}

// Type of account/card that is being used for the transaction. Maps to EMV 3DS
// field `acctType`.
type ThreeDSAuthenticationAccountType string

const (
	ThreeDSAuthenticationAccountTypeCredit        ThreeDSAuthenticationAccountType = "CREDIT"
	ThreeDSAuthenticationAccountTypeDebit         ThreeDSAuthenticationAccountType = "DEBIT"
	ThreeDSAuthenticationAccountTypeNotApplicable ThreeDSAuthenticationAccountType = "NOT_APPLICABLE"
)

func (r ThreeDSAuthenticationAccountType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationAccountTypeCredit, ThreeDSAuthenticationAccountTypeDebit, ThreeDSAuthenticationAccountTypeNotApplicable:
		return true
	}
	return false
}

// Indicates the outcome of the 3DS authentication process.
type ThreeDSAuthenticationAuthenticationResult string

const (
	ThreeDSAuthenticationAuthenticationResultDecline          ThreeDSAuthenticationAuthenticationResult = "DECLINE"
	ThreeDSAuthenticationAuthenticationResultSuccess          ThreeDSAuthenticationAuthenticationResult = "SUCCESS"
	ThreeDSAuthenticationAuthenticationResultPendingChallenge ThreeDSAuthenticationAuthenticationResult = "PENDING_CHALLENGE"
	ThreeDSAuthenticationAuthenticationResultPendingDecision  ThreeDSAuthenticationAuthenticationResult = "PENDING_DECISION"
)

func (r ThreeDSAuthenticationAuthenticationResult) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationAuthenticationResultDecline, ThreeDSAuthenticationAuthenticationResultSuccess, ThreeDSAuthenticationAuthenticationResultPendingChallenge, ThreeDSAuthenticationAuthenticationResultPendingDecision:
		return true
	}
	return false
}

// Indicates whether the expiration date provided by the cardholder during checkout
// matches Lithic's record of the card's expiration date.
type ThreeDSAuthenticationCardExpiryCheck string

const (
	ThreeDSAuthenticationCardExpiryCheckMatch      ThreeDSAuthenticationCardExpiryCheck = "MATCH"
	ThreeDSAuthenticationCardExpiryCheckMismatch   ThreeDSAuthenticationCardExpiryCheck = "MISMATCH"
	ThreeDSAuthenticationCardExpiryCheckNotPresent ThreeDSAuthenticationCardExpiryCheck = "NOT_PRESENT"
)

func (r ThreeDSAuthenticationCardExpiryCheck) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationCardExpiryCheckMatch, ThreeDSAuthenticationCardExpiryCheckMismatch, ThreeDSAuthenticationCardExpiryCheckNotPresent:
		return true
	}
	return false
}

// Object containing data about the cardholder provided during the transaction.
type ThreeDSAuthenticationCardholder struct {
	// Indicates whether the shipping address and billing address provided by the
	// cardholder are the same. This value - and assessment of whether the addresses
	// match - is provided directly in the 3DS request and is not determined by Lithic.
	// Maps to EMV 3DS field `addrMatch`.
	AddressMatch bool `json:"address_match,nullable"`
	// Lithic's evaluation result comparing the transaction's address data with the
	// cardholder KYC data if it exists. In the event Lithic does not have any
	// Cardholder KYC data, or the transaction does not contain any address data,
	// NOT_PRESENT will be returned
	AddressOnFileMatch ThreeDSAuthenticationCardholderAddressOnFileMatch `json:"address_on_file_match"`
	// Object containing data on the billing address provided during the transaction.
	BillingAddress ThreeDSAuthenticationCardholderBillingAddress `json:"billing_address"`
	// Email address that is either provided by the cardholder or is on file with the
	// merchant in a 3RI request. Maps to EMV 3DS field `email`.
	Email string `json:"email,nullable"`
	// Name of the cardholder. Maps to EMV 3DS field `cardholderName`.
	Name string `json:"name,nullable"`
	// Home phone number in E.164 format provided by the cardholder. Maps to EMV 3DS
	// fields `homePhone.cc` and `homePhone.subscriber`.
	PhoneNumberHome string `json:"phone_number_home,nullable"`
	// Mobile/cell phone number in E.164 format provided by the cardholder. Maps to EMV
	// 3DS fields `mobilePhone.cc` and `mobilePhone.subscriber`.
	PhoneNumberMobile string `json:"phone_number_mobile,nullable"`
	// Work phone number in E.164 format provided by the cardholder. Maps to EMV 3DS
	// fields `workPhone.cc` and `workPhone.subscriber`.
	PhoneNumberWork string `json:"phone_number_work,nullable"`
	// Object containing data on the shipping address provided during the transaction.
	ShippingAddress ThreeDSAuthenticationCardholderShippingAddress `json:"shipping_address"`
	JSON            threeDSAuthenticationCardholderJSON            `json:"-"`
}

// threeDSAuthenticationCardholderJSON contains the JSON metadata for the struct
// [ThreeDSAuthenticationCardholder]
type threeDSAuthenticationCardholderJSON struct {
	AddressMatch       apijson.Field
	AddressOnFileMatch apijson.Field
	BillingAddress     apijson.Field
	Email              apijson.Field
	Name               apijson.Field
	PhoneNumberHome    apijson.Field
	PhoneNumberMobile  apijson.Field
	PhoneNumberWork    apijson.Field
	ShippingAddress    apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ThreeDSAuthenticationCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationCardholderJSON) RawJSON() string {
	return r.raw
}

// Lithic's evaluation result comparing the transaction's address data with the
// cardholder KYC data if it exists. In the event Lithic does not have any
// Cardholder KYC data, or the transaction does not contain any address data,
// NOT_PRESENT will be returned
type ThreeDSAuthenticationCardholderAddressOnFileMatch string

const (
	ThreeDSAuthenticationCardholderAddressOnFileMatchMatch            ThreeDSAuthenticationCardholderAddressOnFileMatch = "MATCH"
	ThreeDSAuthenticationCardholderAddressOnFileMatchMatchAddressOnly ThreeDSAuthenticationCardholderAddressOnFileMatch = "MATCH_ADDRESS_ONLY"
	ThreeDSAuthenticationCardholderAddressOnFileMatchMatchZipOnly     ThreeDSAuthenticationCardholderAddressOnFileMatch = "MATCH_ZIP_ONLY"
	ThreeDSAuthenticationCardholderAddressOnFileMatchMismatch         ThreeDSAuthenticationCardholderAddressOnFileMatch = "MISMATCH"
	ThreeDSAuthenticationCardholderAddressOnFileMatchNotPresent       ThreeDSAuthenticationCardholderAddressOnFileMatch = "NOT_PRESENT"
)

func (r ThreeDSAuthenticationCardholderAddressOnFileMatch) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationCardholderAddressOnFileMatchMatch, ThreeDSAuthenticationCardholderAddressOnFileMatchMatchAddressOnly, ThreeDSAuthenticationCardholderAddressOnFileMatchMatchZipOnly, ThreeDSAuthenticationCardholderAddressOnFileMatchMismatch, ThreeDSAuthenticationCardholderAddressOnFileMatchNotPresent:
		return true
	}
	return false
}

// Object containing data on the billing address provided during the transaction.
type ThreeDSAuthenticationCardholderBillingAddress struct {
	// First line of the street address provided by the cardholder.
	Address1 string `json:"address1,nullable"`
	// Second line of the street address provided by the cardholder.
	Address2 string `json:"address2,nullable"`
	// Third line of the street address provided by the cardholder.
	Address3 string `json:"address3,nullable"`
	// City of the address provided by the cardholder.
	City string `json:"city,nullable"`
	// Country of the address provided by the cardholder in ISO 3166-1 alpha-3 format
	// (e.g. USA)
	Country string `json:"country,nullable"`
	// Postal code (e.g., ZIP code) of the address provided by the cardholder
	PostalCode string                                            `json:"postal_code,nullable"`
	JSON       threeDSAuthenticationCardholderBillingAddressJSON `json:"-"`
}

// threeDSAuthenticationCardholderBillingAddressJSON contains the JSON metadata for
// the struct [ThreeDSAuthenticationCardholderBillingAddress]
type threeDSAuthenticationCardholderBillingAddressJSON struct {
	Address1    apijson.Field
	Address2    apijson.Field
	Address3    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationCardholderBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationCardholderBillingAddressJSON) RawJSON() string {
	return r.raw
}

// Object containing data on the shipping address provided during the transaction.
type ThreeDSAuthenticationCardholderShippingAddress struct {
	// First line of the street address provided by the cardholder.
	Address1 string `json:"address1,nullable"`
	// Second line of the street address provided by the cardholder.
	Address2 string `json:"address2,nullable"`
	// Third line of the street address provided by the cardholder.
	Address3 string `json:"address3,nullable"`
	// City of the address provided by the cardholder.
	City string `json:"city,nullable"`
	// Country of the address provided by the cardholder in ISO 3166-1 alpha-3 format
	// (e.g. USA)
	Country string `json:"country,nullable"`
	// Postal code (e.g., ZIP code) of the address provided by the cardholder
	PostalCode string                                             `json:"postal_code,nullable"`
	JSON       threeDSAuthenticationCardholderShippingAddressJSON `json:"-"`
}

// threeDSAuthenticationCardholderShippingAddressJSON contains the JSON metadata
// for the struct [ThreeDSAuthenticationCardholderShippingAddress]
type threeDSAuthenticationCardholderShippingAddressJSON struct {
	Address1    apijson.Field
	Address2    apijson.Field
	Address3    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationCardholderShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationCardholderShippingAddressJSON) RawJSON() string {
	return r.raw
}

// Channel in which the authentication occurs. Maps to EMV 3DS field
// `deviceChannel`.
type ThreeDSAuthenticationChannel string

const (
	ThreeDSAuthenticationChannelAppBased                  ThreeDSAuthenticationChannel = "APP_BASED"
	ThreeDSAuthenticationChannelBrowser                   ThreeDSAuthenticationChannel = "BROWSER"
	ThreeDSAuthenticationChannelThreeDSRequestorInitiated ThreeDSAuthenticationChannel = "THREE_DS_REQUESTOR_INITIATED"
)

func (r ThreeDSAuthenticationChannel) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationChannelAppBased, ThreeDSAuthenticationChannelBrowser, ThreeDSAuthenticationChannelThreeDSRequestorInitiated:
		return true
	}
	return false
}

// Object containing data about the merchant involved in the e-commerce
// transaction.
type ThreeDSAuthenticationMerchant struct {
	// Object containing additional data indicating additional risk factors related to
	// the e-commerce transaction.
	RiskIndicator ThreeDSAuthenticationMerchantRiskIndicator `json:"risk_indicator,required"`
	// Merchant identifier as assigned by the acquirer. Maps to EMV 3DS field
	// `acquirerMerchantId`. May not be present for non-payment authentications.
	ID string `json:"id,nullable"`
	// Country code of the merchant requesting 3DS authentication. Maps to EMV 3DS
	// field `merchantCountryCode`. Permitted values: ISO 3166-1 alpha-3 country code
	// (e.g., USA). May not be present for non-payment authentications.
	Country string `json:"country,nullable"`
	// Merchant category code assigned to the merchant that describes its business
	// activity type. Maps to EMV 3DS field `mcc`. May not be present for non-payment
	// authentications.
	Mcc string `json:"mcc,nullable"`
	// Name of the merchant. Maps to EMV 3DS field `merchantName`. May not be present
	// for non-payment authentications.
	Name string                            `json:"name,nullable"`
	JSON threeDSAuthenticationMerchantJSON `json:"-"`
}

// threeDSAuthenticationMerchantJSON contains the JSON metadata for the struct
// [ThreeDSAuthenticationMerchant]
type threeDSAuthenticationMerchantJSON struct {
	RiskIndicator apijson.Field
	ID            apijson.Field
	Country       apijson.Field
	Mcc           apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreeDSAuthenticationMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationMerchantJSON) RawJSON() string {
	return r.raw
}

// Object containing additional data indicating additional risk factors related to
// the e-commerce transaction.
type ThreeDSAuthenticationMerchantRiskIndicator struct {
	// In transactions with electronic delivery, email address to which merchandise is
	// delivered. Maps to EMV 3DS field `deliveryEmailAddress`.
	DeliveryEmailAddress string `json:"delivery_email_address,nullable"`
	// The delivery time frame for the merchandise. Maps to EMV 3DS field
	// `deliveryTimeframe`.
	DeliveryTimeFrame ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrame `json:"delivery_time_frame,nullable"`
	// In prepaid or gift card purchase transactions, purchase amount total in major
	// units (e.g., a purchase of USD $205.10 would be 205). Maps to EMV 3DS field
	// `giftCardAmount`.
	GiftCardAmount int64 `json:"gift_card_amount,nullable"`
	// In prepaid or gift card purchase transactions, count of individual prepaid or
	// gift cards/codes purchased. Maps to EMV 3DS field `giftCardCount`.
	GiftCardCount int64 `json:"gift_card_count,nullable"`
	// In prepaid or gift card purchase transactions, currency code of the gift card.
	// Maps to EMV 3DS field `giftCardCurr`. Permitted values: ISO 4217 three-character
	// currency code (e.g., USD).
	GiftCardCurrency string `json:"gift_card_currency,nullable"`
	// Indicates whether the purchase is for merchandise that is available now or at a
	// future date. Maps to EMV 3DS field `preOrderPurchaseInd`.
	OrderAvailability ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailability `json:"order_availability,nullable"`
	// In pre-order purchase transactions, the expected date that the merchandise will
	// be available. Maps to EMV 3DS field `preOrderDate`. Permitted values: Date
	// string in the ISO 8601 format yyyy-MM-dd'T'hh:mm:ssZ
	PreOrderAvailableDate time.Time `json:"pre_order_available_date,nullable" format:"date-time"`
	// Indicates whether the cardholder is reordering previously purchased merchandise.
	// Maps to EMV 3DS field `reorderItemsInd`.
	ReorderItems ThreeDSAuthenticationMerchantRiskIndicatorReorderItems `json:"reorder_items,nullable"`
	// Shipping method that the cardholder chose for the transaction. If purchase
	// includes one or more item, this indicator is used for the physical goods; if the
	// purchase only includes digital goods, this indicator is used to describe the
	// most expensive item purchased. Maps to EMV 3DS field `shipIndicator`.
	ShippingMethod ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod `json:"shipping_method,nullable"`
	JSON           threeDSAuthenticationMerchantRiskIndicatorJSON           `json:"-"`
}

// threeDSAuthenticationMerchantRiskIndicatorJSON contains the JSON metadata for
// the struct [ThreeDSAuthenticationMerchantRiskIndicator]
type threeDSAuthenticationMerchantRiskIndicatorJSON struct {
	DeliveryEmailAddress  apijson.Field
	DeliveryTimeFrame     apijson.Field
	GiftCardAmount        apijson.Field
	GiftCardCount         apijson.Field
	GiftCardCurrency      apijson.Field
	OrderAvailability     apijson.Field
	PreOrderAvailableDate apijson.Field
	ReorderItems          apijson.Field
	ShippingMethod        apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ThreeDSAuthenticationMerchantRiskIndicator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationMerchantRiskIndicatorJSON) RawJSON() string {
	return r.raw
}

// The delivery time frame for the merchandise. Maps to EMV 3DS field
// `deliveryTimeframe`.
type ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrame string

const (
	ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrameElectronicDelivery   ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrame = "ELECTRONIC_DELIVERY"
	ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrameOvernightShipping    ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrame = "OVERNIGHT_SHIPPING"
	ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrameSameDayShipping      ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrame = "SAME_DAY_SHIPPING"
	ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrameTwoDayOrMoreShipping ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrame = "TWO_DAY_OR_MORE_SHIPPING"
)

func (r ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrame) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrameElectronicDelivery, ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrameOvernightShipping, ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrameSameDayShipping, ThreeDSAuthenticationMerchantRiskIndicatorDeliveryTimeFrameTwoDayOrMoreShipping:
		return true
	}
	return false
}

// Indicates whether the purchase is for merchandise that is available now or at a
// future date. Maps to EMV 3DS field `preOrderPurchaseInd`.
type ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailability string

const (
	ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailabilityFutureAvailability   ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailability = "FUTURE_AVAILABILITY"
	ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailabilityMerchandiseAvailable ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailability = "MERCHANDISE_AVAILABLE"
)

func (r ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailability) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailabilityFutureAvailability, ThreeDSAuthenticationMerchantRiskIndicatorOrderAvailabilityMerchandiseAvailable:
		return true
	}
	return false
}

// Indicates whether the cardholder is reordering previously purchased merchandise.
// Maps to EMV 3DS field `reorderItemsInd`.
type ThreeDSAuthenticationMerchantRiskIndicatorReorderItems string

const (
	ThreeDSAuthenticationMerchantRiskIndicatorReorderItemsFirstTimeOrdered ThreeDSAuthenticationMerchantRiskIndicatorReorderItems = "FIRST_TIME_ORDERED"
	ThreeDSAuthenticationMerchantRiskIndicatorReorderItemsReordered        ThreeDSAuthenticationMerchantRiskIndicatorReorderItems = "REORDERED"
)

func (r ThreeDSAuthenticationMerchantRiskIndicatorReorderItems) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationMerchantRiskIndicatorReorderItemsFirstTimeOrdered, ThreeDSAuthenticationMerchantRiskIndicatorReorderItemsReordered:
		return true
	}
	return false
}

// Shipping method that the cardholder chose for the transaction. If purchase
// includes one or more item, this indicator is used for the physical goods; if the
// purchase only includes digital goods, this indicator is used to describe the
// most expensive item purchased. Maps to EMV 3DS field `shipIndicator`.
type ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod string

const (
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodDigitalGoods               ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "DIGITAL_GOODS"
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodLockerDelivery             ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "LOCKER_DELIVERY"
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodOther                      ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "OTHER"
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodPickUpAndGoDelivery        ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "PICK_UP_AND_GO_DELIVERY"
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodShipToBillingAddress       ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "SHIP_TO_BILLING_ADDRESS"
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodShipToNonBillingAddress    ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "SHIP_TO_NON_BILLING_ADDRESS"
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodShipToOtherVerifiedAddress ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "SHIP_TO_OTHER_VERIFIED_ADDRESS"
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodShipToStore                ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "SHIP_TO_STORE"
	ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodTravelAndEventTickets      ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod = "TRAVEL_AND_EVENT_TICKETS"
)

func (r ThreeDSAuthenticationMerchantRiskIndicatorShippingMethod) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodDigitalGoods, ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodLockerDelivery, ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodOther, ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodPickUpAndGoDelivery, ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodShipToBillingAddress, ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodShipToNonBillingAddress, ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodShipToOtherVerifiedAddress, ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodShipToStore, ThreeDSAuthenticationMerchantRiskIndicatorShippingMethodTravelAndEventTickets:
		return true
	}
	return false
}

// Either PAYMENT_AUTHENTICATION or NON_PAYMENT_AUTHENTICATION. For
// NON_PAYMENT_AUTHENTICATION, additional_data and transaction fields are not
// populated.
type ThreeDSAuthenticationMessageCategory string

const (
	ThreeDSAuthenticationMessageCategoryNonPaymentAuthentication ThreeDSAuthenticationMessageCategory = "NON_PAYMENT_AUTHENTICATION"
	ThreeDSAuthenticationMessageCategoryPaymentAuthentication    ThreeDSAuthenticationMessageCategory = "PAYMENT_AUTHENTICATION"
)

func (r ThreeDSAuthenticationMessageCategory) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationMessageCategoryNonPaymentAuthentication, ThreeDSAuthenticationMessageCategoryPaymentAuthentication:
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
type ThreeDSAuthenticationThreeDSRequestorChallengeIndicator string

const (
	ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorNoPreference                   ThreeDSAuthenticationThreeDSRequestorChallengeIndicator = "NO_PREFERENCE"
	ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorNoChallengeRequested           ThreeDSAuthenticationThreeDSRequestorChallengeIndicator = "NO_CHALLENGE_REQUESTED"
	ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorChallengePreference            ThreeDSAuthenticationThreeDSRequestorChallengeIndicator = "CHALLENGE_PREFERENCE"
	ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorChallengeMandate               ThreeDSAuthenticationThreeDSRequestorChallengeIndicator = "CHALLENGE_MANDATE"
	ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorNoChallengeRiskAlreadyAssessed ThreeDSAuthenticationThreeDSRequestorChallengeIndicator = "NO_CHALLENGE_RISK_ALREADY_ASSESSED"
	ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorDataShareOnly                  ThreeDSAuthenticationThreeDSRequestorChallengeIndicator = "DATA_SHARE_ONLY"
	ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorOther                          ThreeDSAuthenticationThreeDSRequestorChallengeIndicator = "OTHER"
)

func (r ThreeDSAuthenticationThreeDSRequestorChallengeIndicator) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorNoPreference, ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorNoChallengeRequested, ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorChallengePreference, ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorChallengeMandate, ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorNoChallengeRiskAlreadyAssessed, ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorDataShareOnly, ThreeDSAuthenticationThreeDSRequestorChallengeIndicatorOther:
		return true
	}
	return false
}

// Object containing additional data about the 3DS request that is beyond the EMV
// 3DS standard spec (e.g., specific fields that only certain card networks send
// but are not required across all 3DS requests).
type ThreeDSAuthenticationAdditionalData struct {
	// Mastercard only: Indicates whether the network would have considered the
	// authentication request to be low risk or not.
	NetworkDecision ThreeDSAuthenticationAdditionalDataNetworkDecision `json:"network_decision,nullable"`
	// Mastercard only: Assessment by the network of the authentication risk level,
	// with a higher value indicating a higher amount of risk. Permitted values:
	// Integer between 0-950, in increments of 50.
	NetworkRiskScore int64                                   `json:"network_risk_score,nullable"`
	JSON             threeDSAuthenticationAdditionalDataJSON `json:"-"`
}

// threeDSAuthenticationAdditionalDataJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationAdditionalData]
type threeDSAuthenticationAdditionalDataJSON struct {
	NetworkDecision  apijson.Field
	NetworkRiskScore apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreeDSAuthenticationAdditionalData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationAdditionalDataJSON) RawJSON() string {
	return r.raw
}

// Mastercard only: Indicates whether the network would have considered the
// authentication request to be low risk or not.
type ThreeDSAuthenticationAdditionalDataNetworkDecision string

const (
	ThreeDSAuthenticationAdditionalDataNetworkDecisionLowRisk    ThreeDSAuthenticationAdditionalDataNetworkDecision = "LOW_RISK"
	ThreeDSAuthenticationAdditionalDataNetworkDecisionNotLowRisk ThreeDSAuthenticationAdditionalDataNetworkDecision = "NOT_LOW_RISK"
)

func (r ThreeDSAuthenticationAdditionalDataNetworkDecision) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationAdditionalDataNetworkDecisionLowRisk, ThreeDSAuthenticationAdditionalDataNetworkDecisionNotLowRisk:
		return true
	}
	return false
}

// Object containing data about the app used in the e-commerce transaction. Present
// if the channel is 'APP_BASED'.
type ThreeDSAuthenticationApp struct {
	// Device model: e.g. "Apple iPhone 16".
	Device string `json:"device,nullable"`
	// Raw device information - base64-encoded JSON object. Maps to EMV 3DS field
	// `deviceInfo`.
	DeviceInfo string `json:"device_info,nullable"`
	// IP address of the device.
	Ip string `json:"ip"`
	// Latitude coordinate of current device location.
	Latitude float64 `json:"latitude,nullable"`
	// Device locale: e.g. "en-US".
	Locale string `json:"locale,nullable"`
	// Longitude coordinate of current device location.
	Longitude float64 `json:"longitude,nullable"`
	// Operating System: e.g. "Android 12", "iOS 17.1".
	Os string `json:"os,nullable"`
	// Device platform: Android, iOS, Windows, etc.
	Platform string `json:"platform,nullable"`
	// Screen height in pixels.
	ScreenHeight int64 `json:"screen_height,nullable"`
	// Screen width in pixels.
	ScreenWidth int64 `json:"screen_width,nullable"`
	// Time zone offset in minutes between UTC and device local time.
	TimeZone string                       `json:"time_zone,nullable"`
	JSON     threeDSAuthenticationAppJSON `json:"-"`
}

// threeDSAuthenticationAppJSON contains the JSON metadata for the struct
// [ThreeDSAuthenticationApp]
type threeDSAuthenticationAppJSON struct {
	Device       apijson.Field
	DeviceInfo   apijson.Field
	Ip           apijson.Field
	Latitude     apijson.Field
	Locale       apijson.Field
	Longitude    apijson.Field
	Os           apijson.Field
	Platform     apijson.Field
	ScreenHeight apijson.Field
	ScreenWidth  apijson.Field
	TimeZone     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreeDSAuthenticationApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationAppJSON) RawJSON() string {
	return r.raw
}

// Type of authentication request - i.e., the type of transaction or interaction is
// causing the merchant to request an authentication. Maps to EMV 3DS field
// `threeDSRequestorAuthenticationInd`.
type ThreeDSAuthenticationAuthenticationRequestType string

const (
	ThreeDSAuthenticationAuthenticationRequestTypeAddCard                        ThreeDSAuthenticationAuthenticationRequestType = "ADD_CARD"
	ThreeDSAuthenticationAuthenticationRequestTypeBillingAgreement               ThreeDSAuthenticationAuthenticationRequestType = "BILLING_AGREEMENT"
	ThreeDSAuthenticationAuthenticationRequestTypeDelayedShipment                ThreeDSAuthenticationAuthenticationRequestType = "DELAYED_SHIPMENT"
	ThreeDSAuthenticationAuthenticationRequestTypeEmvTokenCardholderVerification ThreeDSAuthenticationAuthenticationRequestType = "EMV_TOKEN_CARDHOLDER_VERIFICATION"
	ThreeDSAuthenticationAuthenticationRequestTypeInstallmentTransaction         ThreeDSAuthenticationAuthenticationRequestType = "INSTALLMENT_TRANSACTION"
	ThreeDSAuthenticationAuthenticationRequestTypeMaintainCard                   ThreeDSAuthenticationAuthenticationRequestType = "MAINTAIN_CARD"
	ThreeDSAuthenticationAuthenticationRequestTypePaymentTransaction             ThreeDSAuthenticationAuthenticationRequestType = "PAYMENT_TRANSACTION"
	ThreeDSAuthenticationAuthenticationRequestTypeRecurringTransaction           ThreeDSAuthenticationAuthenticationRequestType = "RECURRING_TRANSACTION"
	ThreeDSAuthenticationAuthenticationRequestTypeSplitPayment                   ThreeDSAuthenticationAuthenticationRequestType = "SPLIT_PAYMENT"
	ThreeDSAuthenticationAuthenticationRequestTypeSplitShipment                  ThreeDSAuthenticationAuthenticationRequestType = "SPLIT_SHIPMENT"
)

func (r ThreeDSAuthenticationAuthenticationRequestType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationAuthenticationRequestTypeAddCard, ThreeDSAuthenticationAuthenticationRequestTypeBillingAgreement, ThreeDSAuthenticationAuthenticationRequestTypeDelayedShipment, ThreeDSAuthenticationAuthenticationRequestTypeEmvTokenCardholderVerification, ThreeDSAuthenticationAuthenticationRequestTypeInstallmentTransaction, ThreeDSAuthenticationAuthenticationRequestTypeMaintainCard, ThreeDSAuthenticationAuthenticationRequestTypePaymentTransaction, ThreeDSAuthenticationAuthenticationRequestTypeRecurringTransaction, ThreeDSAuthenticationAuthenticationRequestTypeSplitPayment, ThreeDSAuthenticationAuthenticationRequestTypeSplitShipment:
		return true
	}
	return false
}

// Object containing data about the browser used in the e-commerce transaction.
// Present if the channel is 'BROWSER'.
type ThreeDSAuthenticationBrowser struct {
	// Content of the HTTP accept headers as sent from the cardholder's browser to the
	// 3DS requestor (e.g., merchant or digital wallet).
	AcceptHeader string `json:"accept_header,nullable"`
	// IP address of the browser as returned by the HTTP headers to the 3DS requestor
	// (e.g., merchant or digital wallet). Maps to EMV 3DS field `browserIP`.
	Ip string `json:"ip,nullable"`
	// Indicates whether the cardholder's browser has the ability to execute Java. Maps
	// to EMV 3DS field `browserJavaEnabled`.
	JavaEnabled bool `json:"java_enabled,nullable"`
	// Indicates whether the cardholder's browser has the ability to execute
	// JavaScript. Maps to EMV 3DS field `browserJavascriptEnabled`.
	JavascriptEnabled bool `json:"javascript_enabled,nullable"`
	// Language of the cardholder's browser as defined in IETF BCP47. Maps to EMV 3DS
	// field `browserLanguage`.
	Language string `json:"language,nullable"`
	// Time zone offset in minutes between UTC and browser local time. Maps to EMV 3DS
	// field `browserTz`.
	TimeZone string `json:"time_zone,nullable"`
	// Content of the HTTP user-agent header. Maps to EMV 3DS field `browserUserAgent`.
	UserAgent string                           `json:"user_agent,nullable"`
	JSON      threeDSAuthenticationBrowserJSON `json:"-"`
}

// threeDSAuthenticationBrowserJSON contains the JSON metadata for the struct
// [ThreeDSAuthenticationBrowser]
type threeDSAuthenticationBrowserJSON struct {
	AcceptHeader      apijson.Field
	Ip                apijson.Field
	JavaEnabled       apijson.Field
	JavascriptEnabled apijson.Field
	Language          apijson.Field
	TimeZone          apijson.Field
	UserAgent         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ThreeDSAuthenticationBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationBrowserJSON) RawJSON() string {
	return r.raw
}

// Metadata about the challenge method and delivery. Only present when a challenge
// is triggered.
type ThreeDSAuthenticationChallengeMetadata struct {
	// The type of challenge method used for authentication.
	MethodType ThreeDSAuthenticationChallengeMetadataMethodType `json:"method_type,required"`
	// Indicates the status of the challenge
	//
	//   - SUCCESS - Cardholder completed the challenge successfully
	//   - PENDING - Challenge was issued to the cardholder and was not completed yet
	//   - SMS_DELIVERY_FAILED - Lithic confirmed undeliverability of the SMS to the
	//     provided phone number. Relevant only for SMS_OTP method
	//   - CARDHOLDER_TIMEOUT - Cardholder failed to complete the challenge within the
	//     given challenge TTL
	//   - CANCELED_VIA_CHALLENGE_UI - Cardholder canceled the challenge by selecting
	//     "cancel" on the challenge UI
	//   - CANCELED_OOB - Cardholder canceled the challenge out of band
	//   - ATTEMPTS_EXCEEDED - Cardholder failed the challenge by either entering an
	//     incorrect OTP more than the allowed number of times or requesting a new OTP
	//     more than the allowed number of times
	//   - ABORTED - Merchant aborted authentication after a challenge was requested
	//   - ERROR - The challenge failed for a reason different than those documented
	Status ThreeDSAuthenticationChallengeMetadataStatus `json:"status,required"`
	// The phone number used for delivering the OTP. Relevant only for SMS_OTP method.
	PhoneNumber string                                     `json:"phone_number,nullable"`
	JSON        threeDSAuthenticationChallengeMetadataJSON `json:"-"`
}

// threeDSAuthenticationChallengeMetadataJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationChallengeMetadata]
type threeDSAuthenticationChallengeMetadataJSON struct {
	MethodType  apijson.Field
	Status      apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationChallengeMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationChallengeMetadataJSON) RawJSON() string {
	return r.raw
}

// The type of challenge method used for authentication.
type ThreeDSAuthenticationChallengeMetadataMethodType string

const (
	ThreeDSAuthenticationChallengeMetadataMethodTypeSMSOtp    ThreeDSAuthenticationChallengeMetadataMethodType = "SMS_OTP"
	ThreeDSAuthenticationChallengeMetadataMethodTypeOutOfBand ThreeDSAuthenticationChallengeMetadataMethodType = "OUT_OF_BAND"
)

func (r ThreeDSAuthenticationChallengeMetadataMethodType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationChallengeMetadataMethodTypeSMSOtp, ThreeDSAuthenticationChallengeMetadataMethodTypeOutOfBand:
		return true
	}
	return false
}

// Indicates the status of the challenge
//
//   - SUCCESS - Cardholder completed the challenge successfully
//   - PENDING - Challenge was issued to the cardholder and was not completed yet
//   - SMS_DELIVERY_FAILED - Lithic confirmed undeliverability of the SMS to the
//     provided phone number. Relevant only for SMS_OTP method
//   - CARDHOLDER_TIMEOUT - Cardholder failed to complete the challenge within the
//     given challenge TTL
//   - CANCELED_VIA_CHALLENGE_UI - Cardholder canceled the challenge by selecting
//     "cancel" on the challenge UI
//   - CANCELED_OOB - Cardholder canceled the challenge out of band
//   - ATTEMPTS_EXCEEDED - Cardholder failed the challenge by either entering an
//     incorrect OTP more than the allowed number of times or requesting a new OTP
//     more than the allowed number of times
//   - ABORTED - Merchant aborted authentication after a challenge was requested
//   - ERROR - The challenge failed for a reason different than those documented
type ThreeDSAuthenticationChallengeMetadataStatus string

const (
	ThreeDSAuthenticationChallengeMetadataStatusSuccess                ThreeDSAuthenticationChallengeMetadataStatus = "SUCCESS"
	ThreeDSAuthenticationChallengeMetadataStatusPending                ThreeDSAuthenticationChallengeMetadataStatus = "PENDING"
	ThreeDSAuthenticationChallengeMetadataStatusSMSDeliveryFailed      ThreeDSAuthenticationChallengeMetadataStatus = "SMS_DELIVERY_FAILED"
	ThreeDSAuthenticationChallengeMetadataStatusCardholderTimeout      ThreeDSAuthenticationChallengeMetadataStatus = "CARDHOLDER_TIMEOUT"
	ThreeDSAuthenticationChallengeMetadataStatusCanceledViaChallengeUi ThreeDSAuthenticationChallengeMetadataStatus = "CANCELED_VIA_CHALLENGE_UI"
	ThreeDSAuthenticationChallengeMetadataStatusCanceledOob            ThreeDSAuthenticationChallengeMetadataStatus = "CANCELED_OOB"
	ThreeDSAuthenticationChallengeMetadataStatusAttemptsExceeded       ThreeDSAuthenticationChallengeMetadataStatus = "ATTEMPTS_EXCEEDED"
	ThreeDSAuthenticationChallengeMetadataStatusAborted                ThreeDSAuthenticationChallengeMetadataStatus = "ABORTED"
	ThreeDSAuthenticationChallengeMetadataStatusError                  ThreeDSAuthenticationChallengeMetadataStatus = "ERROR"
)

func (r ThreeDSAuthenticationChallengeMetadataStatus) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationChallengeMetadataStatusSuccess, ThreeDSAuthenticationChallengeMetadataStatusPending, ThreeDSAuthenticationChallengeMetadataStatusSMSDeliveryFailed, ThreeDSAuthenticationChallengeMetadataStatusCardholderTimeout, ThreeDSAuthenticationChallengeMetadataStatusCanceledViaChallengeUi, ThreeDSAuthenticationChallengeMetadataStatusCanceledOob, ThreeDSAuthenticationChallengeMetadataStatusAttemptsExceeded, ThreeDSAuthenticationChallengeMetadataStatusAborted, ThreeDSAuthenticationChallengeMetadataStatusError:
		return true
	}
	return false
}

// Entity that orchestrates the challenge. This won't be set for authentications
// for which a decision has not yet been made (e.g. in-flight customer decisioning
// request).
type ThreeDSAuthenticationChallengeOrchestratedBy string

const (
	ThreeDSAuthenticationChallengeOrchestratedByLithic      ThreeDSAuthenticationChallengeOrchestratedBy = "LITHIC"
	ThreeDSAuthenticationChallengeOrchestratedByCustomer    ThreeDSAuthenticationChallengeOrchestratedBy = "CUSTOMER"
	ThreeDSAuthenticationChallengeOrchestratedByNoChallenge ThreeDSAuthenticationChallengeOrchestratedBy = "NO_CHALLENGE"
)

func (r ThreeDSAuthenticationChallengeOrchestratedBy) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationChallengeOrchestratedByLithic, ThreeDSAuthenticationChallengeOrchestratedByCustomer, ThreeDSAuthenticationChallengeOrchestratedByNoChallenge:
		return true
	}
	return false
}

// Entity that made the authentication decision. This won't be set for
// authentications for which a decision has not yet been made (e.g. in-flight
// customer decisioning request).
type ThreeDSAuthenticationDecisionMadeBy string

const (
	ThreeDSAuthenticationDecisionMadeByLithicRules      ThreeDSAuthenticationDecisionMadeBy = "LITHIC_RULES"
	ThreeDSAuthenticationDecisionMadeByLithicDefault    ThreeDSAuthenticationDecisionMadeBy = "LITHIC_DEFAULT"
	ThreeDSAuthenticationDecisionMadeByCustomerRules    ThreeDSAuthenticationDecisionMadeBy = "CUSTOMER_RULES"
	ThreeDSAuthenticationDecisionMadeByCustomerEndpoint ThreeDSAuthenticationDecisionMadeBy = "CUSTOMER_ENDPOINT"
	ThreeDSAuthenticationDecisionMadeByNetwork          ThreeDSAuthenticationDecisionMadeBy = "NETWORK"
	ThreeDSAuthenticationDecisionMadeByUnknown          ThreeDSAuthenticationDecisionMadeBy = "UNKNOWN"
)

func (r ThreeDSAuthenticationDecisionMadeBy) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationDecisionMadeByLithicRules, ThreeDSAuthenticationDecisionMadeByLithicDefault, ThreeDSAuthenticationDecisionMadeByCustomerRules, ThreeDSAuthenticationDecisionMadeByCustomerEndpoint, ThreeDSAuthenticationDecisionMadeByNetwork, ThreeDSAuthenticationDecisionMadeByUnknown:
		return true
	}
	return false
}

// Type of 3DS Requestor Initiated (3RI) request — i.e., a 3DS authentication that
// takes place at the initiation of the merchant rather than the cardholder. The
// most common example of this is where a merchant is authenticating before billing
// for a recurring transaction such as a pay TV subscription or a utility bill.
// Maps to EMV 3DS field `threeRIInd`.
type ThreeDSAuthenticationThreeRiRequestType string

const (
	ThreeDSAuthenticationThreeRiRequestTypeAccountVerification         ThreeDSAuthenticationThreeRiRequestType = "ACCOUNT_VERIFICATION"
	ThreeDSAuthenticationThreeRiRequestTypeAddCard                     ThreeDSAuthenticationThreeRiRequestType = "ADD_CARD"
	ThreeDSAuthenticationThreeRiRequestTypeBillingAgreement            ThreeDSAuthenticationThreeRiRequestType = "BILLING_AGREEMENT"
	ThreeDSAuthenticationThreeRiRequestTypeCardSecurityCodeStatusCheck ThreeDSAuthenticationThreeRiRequestType = "CARD_SECURITY_CODE_STATUS_CHECK"
	ThreeDSAuthenticationThreeRiRequestTypeDelayedShipment             ThreeDSAuthenticationThreeRiRequestType = "DELAYED_SHIPMENT"
	ThreeDSAuthenticationThreeRiRequestTypeDeviceBindingStatusCheck    ThreeDSAuthenticationThreeRiRequestType = "DEVICE_BINDING_STATUS_CHECK"
	ThreeDSAuthenticationThreeRiRequestTypeInstallmentTransaction      ThreeDSAuthenticationThreeRiRequestType = "INSTALLMENT_TRANSACTION"
	ThreeDSAuthenticationThreeRiRequestTypeMailOrder                   ThreeDSAuthenticationThreeRiRequestType = "MAIL_ORDER"
	ThreeDSAuthenticationThreeRiRequestTypeMaintainCardInfo            ThreeDSAuthenticationThreeRiRequestType = "MAINTAIN_CARD_INFO"
	ThreeDSAuthenticationThreeRiRequestTypeOtherPayment                ThreeDSAuthenticationThreeRiRequestType = "OTHER_PAYMENT"
	ThreeDSAuthenticationThreeRiRequestTypeRecurringTransaction        ThreeDSAuthenticationThreeRiRequestType = "RECURRING_TRANSACTION"
	ThreeDSAuthenticationThreeRiRequestTypeSplitPayment                ThreeDSAuthenticationThreeRiRequestType = "SPLIT_PAYMENT"
	ThreeDSAuthenticationThreeRiRequestTypeSplitShipment               ThreeDSAuthenticationThreeRiRequestType = "SPLIT_SHIPMENT"
	ThreeDSAuthenticationThreeRiRequestTypeTelephoneOrder              ThreeDSAuthenticationThreeRiRequestType = "TELEPHONE_ORDER"
	ThreeDSAuthenticationThreeRiRequestTypeTopUp                       ThreeDSAuthenticationThreeRiRequestType = "TOP_UP"
	ThreeDSAuthenticationThreeRiRequestTypeTrustListStatusCheck        ThreeDSAuthenticationThreeRiRequestType = "TRUST_LIST_STATUS_CHECK"
)

func (r ThreeDSAuthenticationThreeRiRequestType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationThreeRiRequestTypeAccountVerification, ThreeDSAuthenticationThreeRiRequestTypeAddCard, ThreeDSAuthenticationThreeRiRequestTypeBillingAgreement, ThreeDSAuthenticationThreeRiRequestTypeCardSecurityCodeStatusCheck, ThreeDSAuthenticationThreeRiRequestTypeDelayedShipment, ThreeDSAuthenticationThreeRiRequestTypeDeviceBindingStatusCheck, ThreeDSAuthenticationThreeRiRequestTypeInstallmentTransaction, ThreeDSAuthenticationThreeRiRequestTypeMailOrder, ThreeDSAuthenticationThreeRiRequestTypeMaintainCardInfo, ThreeDSAuthenticationThreeRiRequestTypeOtherPayment, ThreeDSAuthenticationThreeRiRequestTypeRecurringTransaction, ThreeDSAuthenticationThreeRiRequestTypeSplitPayment, ThreeDSAuthenticationThreeRiRequestTypeSplitShipment, ThreeDSAuthenticationThreeRiRequestTypeTelephoneOrder, ThreeDSAuthenticationThreeRiRequestTypeTopUp, ThreeDSAuthenticationThreeRiRequestTypeTrustListStatusCheck:
		return true
	}
	return false
}

// Object containing data about the e-commerce transaction for which the merchant
// is requesting authentication.
type ThreeDSAuthenticationTransaction struct {
	// Amount of the purchase in minor units of currency with all punctuation removed.
	// Maps to EMV 3DS field `purchaseAmount`.
	Amount float64 `json:"amount,required"`
	// Approximate amount of the purchase in minor units of cardholder currency.
	// Derived from `amount` using a daily conversion rate.
	CardholderAmount float64 `json:"cardholder_amount,required,nullable"`
	// Currency of the purchase. Maps to EMV 3DS field `purchaseCurrency`. Permitted
	// values: ISO 4217 three-character currency code (e.g., USD).
	Currency string `json:"currency,required"`
	// Minor units of currency, as specified in ISO 4217 currency exponent. Maps to EMV
	// 3DS field `purchaseExponent`.
	CurrencyExponent float64 `json:"currency_exponent,required"`
	// Date and time when the authentication was generated by the merchant/acquirer's
	// 3DS server. Maps to EMV 3DS field `purchaseDate`. Permitted values: Date string
	// in the ISO 8601 format yyyy-MM-dd'T'hh:mm:ssZ.
	DateTime time.Time `json:"date_time,required" format:"date-time"`
	// Type of the transaction for which a 3DS authentication request is occurring.
	// Maps to EMV 3DS field `transType`.
	Type ThreeDSAuthenticationTransactionType `json:"type,required,nullable"`
	JSON threeDSAuthenticationTransactionJSON `json:"-"`
}

// threeDSAuthenticationTransactionJSON contains the JSON metadata for the struct
// [ThreeDSAuthenticationTransaction]
type threeDSAuthenticationTransactionJSON struct {
	Amount           apijson.Field
	CardholderAmount apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	DateTime         apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreeDSAuthenticationTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationTransactionJSON) RawJSON() string {
	return r.raw
}

// Type of the transaction for which a 3DS authentication request is occurring.
// Maps to EMV 3DS field `transType`.
type ThreeDSAuthenticationTransactionType string

const (
	ThreeDSAuthenticationTransactionTypeAccountFunding           ThreeDSAuthenticationTransactionType = "ACCOUNT_FUNDING"
	ThreeDSAuthenticationTransactionTypeCheckAcceptance          ThreeDSAuthenticationTransactionType = "CHECK_ACCEPTANCE"
	ThreeDSAuthenticationTransactionTypeGoodsServicePurchase     ThreeDSAuthenticationTransactionType = "GOODS_SERVICE_PURCHASE"
	ThreeDSAuthenticationTransactionTypePrepaidActivationAndLoad ThreeDSAuthenticationTransactionType = "PREPAID_ACTIVATION_AND_LOAD"
	ThreeDSAuthenticationTransactionTypeQuasiCashTransaction     ThreeDSAuthenticationTransactionType = "QUASI_CASH_TRANSACTION"
)

func (r ThreeDSAuthenticationTransactionType) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationTransactionTypeAccountFunding, ThreeDSAuthenticationTransactionTypeCheckAcceptance, ThreeDSAuthenticationTransactionTypeGoodsServicePurchase, ThreeDSAuthenticationTransactionTypePrepaidActivationAndLoad, ThreeDSAuthenticationTransactionTypeQuasiCashTransaction:
		return true
	}
	return false
}
