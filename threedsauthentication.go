// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// ThreeDSAuthenticationService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewThreeDSAuthenticationService]
// method instead.
type ThreeDSAuthenticationService struct {
	Options []option.RequestOption
}

// NewThreeDSAuthenticationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreeDSAuthenticationService(opts ...option.RequestOption) (r *ThreeDSAuthenticationService) {
	r = &ThreeDSAuthenticationService{}
	r.Options = opts
	return
}

// Get 3DS Authentication by token
func (r *ThreeDSAuthenticationService) Get(ctx context.Context, threeDSAuthenticationToken string, opts ...option.RequestOption) (res *ThreeDSAuthenticationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("three_ds_authentication/%s", threeDSAuthenticationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Simulates a 3DS authentication request from the payment network as if it came
// from an ACS. If you're configured for 3DS Customer Decisioning, simulating
// authentications requires your customer decisioning endpoint to be set up
// properly (respond with a valid JSON).
func (r *ThreeDSAuthenticationService) Simulate(ctx context.Context, body ThreeDSAuthenticationSimulateParams, opts ...option.RequestOption) (res *ThreeDSAuthenticationSimulateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "three_ds_authentication/simulate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ThreeDSAuthenticationGetResponse struct {
	// Globally unique identifier for the 3DS authentication.
	Token string `json:"token,required" format:"uuid"`
	// Type of account/card that is being used for the transaction. Maps to EMV 3DS
	// field acctType.
	AccountType ThreeDSAuthenticationGetResponseAccountType `json:"account_type,required,nullable"`
	// Indicates the outcome of the 3DS authentication process.
	AuthenticationResult ThreeDSAuthenticationGetResponseAuthenticationResult `json:"authentication_result,required,nullable"`
	// Indicates whether the expiration date provided by the cardholder during checkout
	// matches Lithic's record of the card's expiration date.
	CardExpiryCheck ThreeDSAuthenticationGetResponseCardExpiryCheck `json:"card_expiry_check,required"`
	// Globally unique identifier for the card on which the 3DS authentication has
	// occurred.
	CardToken string `json:"card_token,required" format:"uuid"`
	// Object containing data about the cardholder provided during the transaction.
	Cardholder ThreeDSAuthenticationGetResponseCardholder `json:"cardholder,required"`
	// Channel in which the authentication occurs. Maps to EMV 3DS field deviceChannel.
	Channel ThreeDSAuthenticationGetResponseChannel `json:"channel,required"`
	// Date and time when the authentication was created in Lithic's system.
	Created time.Time `json:"created,required" format:"date-time"`
	// Entity that made the authentication decision.
	DecisionMadeBy ThreeDSAuthenticationGetResponseDecisionMadeBy `json:"decision_made_by,required,nullable"`
	// Object containing data about the merchant involved in the e-commerce
	// transaction.
	Merchant ThreeDSAuthenticationGetResponseMerchant `json:"merchant,required"`
	// Either PAYMENT_AUTHENTICATION or NON_PAYMENT_AUTHENTICATION. For
	// NON_PAYMENT_AUTHENTICATION, additional_data and transaction fields are not
	// populated.
	MessageCategory ThreeDSAuthenticationGetResponseMessageCategory `json:"message_category,required"`
	// Object containing additional data about the 3DS request that is beyond the EMV
	// 3DS standard spec (e.g., specific fields that only certain card networks send
	// but are not required across all 3DS requests).
	AdditionalData ThreeDSAuthenticationGetResponseAdditionalData `json:"additional_data,nullable"`
	// Object containing data about the app used in the e-commerce transaction. Present
	// if the channel is 'APP_BASED'.
	App ThreeDSAuthenticationGetResponseApp `json:"app"`
	// Type of authentication request - i.e., the type of transaction or interaction is
	// causing the merchant to request an authentication. Maps to EMV 3DS field
	// threeDSRequestorAuthenticationInd.
	AuthenticationRequestType ThreeDSAuthenticationGetResponseAuthenticationRequestType `json:"authentication_request_type,nullable"`
	// Object containing data about the browser used in the e-commerce transaction.
	// Present if the channel is 'BROWSER'.
	Browser ThreeDSAuthenticationGetResponseBrowser `json:"browser"`
	// Type of 3DS Requestor Initiated (3RI) request i.e., a 3DS authentication that
	// takes place at the initiation of the merchant rather than the cardholder. The
	// most common example of this is where a merchant is authenticating before billing
	// for a recurring transaction such as a pay TV subscription or a utility bill.
	// Maps to EMV 3DS field threeRIInd.
	ThreeRiRequestType ThreeDSAuthenticationGetResponseThreeRiRequestType `json:"three_ri_request_type,nullable"`
	// Object containing data about the e-commerce transaction for which the merchant
	// is requesting authentication.
	Transaction ThreeDSAuthenticationGetResponseTransaction `json:"transaction,nullable"`
	JSON        threeDSAuthenticationGetResponseJSON        `json:"-"`
}

// threeDSAuthenticationGetResponseJSON contains the JSON metadata for the struct
// [ThreeDSAuthenticationGetResponse]
type threeDSAuthenticationGetResponseJSON struct {
	Token                     apijson.Field
	AccountType               apijson.Field
	AuthenticationResult      apijson.Field
	CardExpiryCheck           apijson.Field
	CardToken                 apijson.Field
	Cardholder                apijson.Field
	Channel                   apijson.Field
	Created                   apijson.Field
	DecisionMadeBy            apijson.Field
	Merchant                  apijson.Field
	MessageCategory           apijson.Field
	AdditionalData            apijson.Field
	App                       apijson.Field
	AuthenticationRequestType apijson.Field
	Browser                   apijson.Field
	ThreeRiRequestType        apijson.Field
	Transaction               apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of account/card that is being used for the transaction. Maps to EMV 3DS
// field acctType.
type ThreeDSAuthenticationGetResponseAccountType string

const (
	ThreeDSAuthenticationGetResponseAccountTypeCredit        ThreeDSAuthenticationGetResponseAccountType = "CREDIT"
	ThreeDSAuthenticationGetResponseAccountTypeDebit         ThreeDSAuthenticationGetResponseAccountType = "DEBIT"
	ThreeDSAuthenticationGetResponseAccountTypeNotApplicable ThreeDSAuthenticationGetResponseAccountType = "NOT_APPLICABLE"
)

// Indicates the outcome of the 3DS authentication process.
type ThreeDSAuthenticationGetResponseAuthenticationResult string

const (
	ThreeDSAuthenticationGetResponseAuthenticationResultDecline ThreeDSAuthenticationGetResponseAuthenticationResult = "DECLINE"
	ThreeDSAuthenticationGetResponseAuthenticationResultSuccess ThreeDSAuthenticationGetResponseAuthenticationResult = "SUCCESS"
)

// Indicates whether the expiration date provided by the cardholder during checkout
// matches Lithic's record of the card's expiration date.
type ThreeDSAuthenticationGetResponseCardExpiryCheck string

const (
	ThreeDSAuthenticationGetResponseCardExpiryCheckMatch      ThreeDSAuthenticationGetResponseCardExpiryCheck = "MATCH"
	ThreeDSAuthenticationGetResponseCardExpiryCheckMismatch   ThreeDSAuthenticationGetResponseCardExpiryCheck = "MISMATCH"
	ThreeDSAuthenticationGetResponseCardExpiryCheckNotPresent ThreeDSAuthenticationGetResponseCardExpiryCheck = "NOT_PRESENT"
)

// Object containing data about the cardholder provided during the transaction.
type ThreeDSAuthenticationGetResponseCardholder struct {
	// Indicates whether the shipping address and billing address provided by the
	// cardholder are the same. This value - and assessment of whether the addresses
	// match - is provided directly in the 3DS request and is not determined by Lithic.
	// Maps to EMV 3DS field addrMatch.
	AddressMatch bool `json:"address_match,nullable"`
	// Object containing data on the billing address provided during the transaction.
	BillingAddress ThreeDSAuthenticationGetResponseCardholderBillingAddress `json:"billing_address"`
	// Email address that is either provided by the cardholder or is on file with the
	// merchant in a 3RI request. Maps to EMV 3DS field email.
	Email string `json:"email,nullable"`
	// Name of the cardholder. Maps to EMV 3DS field cardholderName.
	Name string `json:"name,nullable"`
	// Home phone number provided by the cardholder. Maps to EMV 3DS fields
	// homePhone.cc and homePhone.subscriber.
	PhoneNumberHome string `json:"phone_number_home,nullable"`
	// Mobile/cell phone number provided by the cardholder. Maps to EMV 3DS fields
	// mobilePhone.cc and mobilePhone.subscriber.
	PhoneNumberMobile string `json:"phone_number_mobile,nullable"`
	// Work phone number provided by the cardholder. Maps to EMV 3DS fields
	// workPhone.cc and workPhone.subscriber.
	PhoneNumberWork string `json:"phone_number_work,nullable"`
	// Object containing data on the shipping address provided during the transaction.
	ShippingAddress ThreeDSAuthenticationGetResponseCardholderShippingAddress `json:"shipping_address"`
	JSON            threeDSAuthenticationGetResponseCardholderJSON            `json:"-"`
}

// threeDSAuthenticationGetResponseCardholderJSON contains the JSON metadata for
// the struct [ThreeDSAuthenticationGetResponseCardholder]
type threeDSAuthenticationGetResponseCardholderJSON struct {
	AddressMatch      apijson.Field
	BillingAddress    apijson.Field
	Email             apijson.Field
	Name              apijson.Field
	PhoneNumberHome   apijson.Field
	PhoneNumberMobile apijson.Field
	PhoneNumberWork   apijson.Field
	ShippingAddress   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponseCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing data on the billing address provided during the transaction.
type ThreeDSAuthenticationGetResponseCardholderBillingAddress struct {
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
	PostalCode string                                                       `json:"postal_code,nullable"`
	JSON       threeDSAuthenticationGetResponseCardholderBillingAddressJSON `json:"-"`
}

// threeDSAuthenticationGetResponseCardholderBillingAddressJSON contains the JSON
// metadata for the struct
// [ThreeDSAuthenticationGetResponseCardholderBillingAddress]
type threeDSAuthenticationGetResponseCardholderBillingAddressJSON struct {
	Address1    apijson.Field
	Address2    apijson.Field
	Address3    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponseCardholderBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing data on the shipping address provided during the transaction.
type ThreeDSAuthenticationGetResponseCardholderShippingAddress struct {
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
	PostalCode string                                                        `json:"postal_code,nullable"`
	JSON       threeDSAuthenticationGetResponseCardholderShippingAddressJSON `json:"-"`
}

// threeDSAuthenticationGetResponseCardholderShippingAddressJSON contains the JSON
// metadata for the struct
// [ThreeDSAuthenticationGetResponseCardholderShippingAddress]
type threeDSAuthenticationGetResponseCardholderShippingAddressJSON struct {
	Address1    apijson.Field
	Address2    apijson.Field
	Address3    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponseCardholderShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Channel in which the authentication occurs. Maps to EMV 3DS field deviceChannel.
type ThreeDSAuthenticationGetResponseChannel string

const (
	ThreeDSAuthenticationGetResponseChannelAppBased                  ThreeDSAuthenticationGetResponseChannel = "APP_BASED"
	ThreeDSAuthenticationGetResponseChannelBrowser                   ThreeDSAuthenticationGetResponseChannel = "BROWSER"
	ThreeDSAuthenticationGetResponseChannelThreeDSRequestorInitiated ThreeDSAuthenticationGetResponseChannel = "THREE_DS_REQUESTOR_INITIATED"
)

// Entity that made the authentication decision.
type ThreeDSAuthenticationGetResponseDecisionMadeBy string

const (
	ThreeDSAuthenticationGetResponseDecisionMadeByCustomerEndpoint ThreeDSAuthenticationGetResponseDecisionMadeBy = "CUSTOMER_ENDPOINT"
	ThreeDSAuthenticationGetResponseDecisionMadeByLithicDefault    ThreeDSAuthenticationGetResponseDecisionMadeBy = "LITHIC_DEFAULT"
	ThreeDSAuthenticationGetResponseDecisionMadeByLithicRules      ThreeDSAuthenticationGetResponseDecisionMadeBy = "LITHIC_RULES"
	ThreeDSAuthenticationGetResponseDecisionMadeByNetwork          ThreeDSAuthenticationGetResponseDecisionMadeBy = "NETWORK"
	ThreeDSAuthenticationGetResponseDecisionMadeByUnknown          ThreeDSAuthenticationGetResponseDecisionMadeBy = "UNKNOWN"
)

// Object containing data about the merchant involved in the e-commerce
// transaction.
type ThreeDSAuthenticationGetResponseMerchant struct {
	// Merchant identifier as assigned by the acquirer. Maps to EMV 3DS field
	// acquirerMerchantId.
	ID string `json:"id,required"`
	// Country code of the merchant requesting 3DS authentication. Maps to EMV 3DS
	// field merchantCountryCode.
	Country string `json:"country,required"`
	// Merchant category code assigned to the merchant that describes its business
	// activity type. Maps to EMV 3DS field mcc.
	Mcc string `json:"mcc,required"`
	// Name of the merchant. Maps to EMV 3DS field merchantName.
	Name string `json:"name,required"`
	// Object containing additional data indicating additional risk factors related to
	// the e-commerce transaction.
	RiskIndicator ThreeDSAuthenticationGetResponseMerchantRiskIndicator `json:"risk_indicator,required"`
	JSON          threeDSAuthenticationGetResponseMerchantJSON          `json:"-"`
}

// threeDSAuthenticationGetResponseMerchantJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationGetResponseMerchant]
type threeDSAuthenticationGetResponseMerchantJSON struct {
	ID            apijson.Field
	Country       apijson.Field
	Mcc           apijson.Field
	Name          apijson.Field
	RiskIndicator apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponseMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing additional data indicating additional risk factors related to
// the e-commerce transaction.
type ThreeDSAuthenticationGetResponseMerchantRiskIndicator struct {
	// In transactions with electronic delivery, email address to which merchandise is
	// delivered. Maps to EMV 3DS field deliveryEmailAddress.
	DeliveryEmailAddress string `json:"delivery_email_address,nullable"`
	// The delivery time frame for the merchandise. Maps to EMV 3DS field
	// deliveryTimeframe.
	DeliveryTimeFrame ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrame `json:"delivery_time_frame,nullable"`
	// In prepaid or gift card purchase transactions, purchase amount total in major
	// units (e.g., a purchase of USD $205.10 would be 205). Maps to EMV 3DS field
	// giftCardAmount.
	GiftCardAmount float64 `json:"gift_card_amount,nullable"`
	// In prepaid or gift card purchase transactions, count of individual prepaid or
	// gift cards/codes purchased. Maps to EMV 3DS field giftCardCount.
	GiftCardCount float64 `json:"gift_card_count,nullable"`
	// In prepaid or gift card purchase transactions, currency code of the gift card.
	// Maps to EMV 3DS field giftCardCurr.
	GiftCardCurrency string `json:"gift_card_currency,nullable"`
	// Indicates whether the purchase is for merchandise that is available now or at a
	// future date. Maps to EMV 3DS field preOrderPurchaseInd.
	OrderAvailability ThreeDSAuthenticationGetResponseMerchantRiskIndicatorOrderAvailability `json:"order_availability,nullable"`
	// In pre-order purchase transactions, the expected date that the merchandise will
	// be available. Maps to EMV 3DS field preOrderDate.
	PreOrderAvailableDate time.Time `json:"pre_order_available_date,nullable" format:"date-time"`
	// Indicates whether the cardholder is reordering previously purchased merchandise.
	// Maps to EMV 3DS field reorderItemsInd.
	ReorderItems ThreeDSAuthenticationGetResponseMerchantRiskIndicatorReorderItems `json:"reorder_items,nullable"`
	// Shipping method that the cardholder chose for the transaction. If purchase
	// includes one or more item, this indicator is used for the physical goods; if the
	// purchase only includes digital goods, this indicator is used to describe the
	// most expensive item purchased. Maps to EMV 3DS field shipIndicator.
	ShippingMethod ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod `json:"shipping_method,nullable"`
	JSON           threeDSAuthenticationGetResponseMerchantRiskIndicatorJSON           `json:"-"`
}

// threeDSAuthenticationGetResponseMerchantRiskIndicatorJSON contains the JSON
// metadata for the struct [ThreeDSAuthenticationGetResponseMerchantRiskIndicator]
type threeDSAuthenticationGetResponseMerchantRiskIndicatorJSON struct {
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

func (r *ThreeDSAuthenticationGetResponseMerchantRiskIndicator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The delivery time frame for the merchandise. Maps to EMV 3DS field
// deliveryTimeframe.
type ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrame string

const (
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrameElectronicDelivery   ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrame = "ELECTRONIC_DELIVERY"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrameOvernightShipping    ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrame = "OVERNIGHT_SHIPPING"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrameSameDayShipping      ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrame = "SAME_DAY_SHIPPING"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrameTwoDayOrMoreShipping ThreeDSAuthenticationGetResponseMerchantRiskIndicatorDeliveryTimeFrame = "TWO_DAY_OR_MORE_SHIPPING"
)

// Indicates whether the purchase is for merchandise that is available now or at a
// future date. Maps to EMV 3DS field preOrderPurchaseInd.
type ThreeDSAuthenticationGetResponseMerchantRiskIndicatorOrderAvailability string

const (
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorOrderAvailabilityFutureAvailability   ThreeDSAuthenticationGetResponseMerchantRiskIndicatorOrderAvailability = "FUTURE_AVAILABILITY"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorOrderAvailabilityMerchandiseAvailable ThreeDSAuthenticationGetResponseMerchantRiskIndicatorOrderAvailability = "MERCHANDISE_AVAILABLE"
)

// Indicates whether the cardholder is reordering previously purchased merchandise.
// Maps to EMV 3DS field reorderItemsInd.
type ThreeDSAuthenticationGetResponseMerchantRiskIndicatorReorderItems string

const (
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorReorderItemsFirstTimeOrdered ThreeDSAuthenticationGetResponseMerchantRiskIndicatorReorderItems = "FIRST_TIME_ORDERED"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorReorderItemsReordered        ThreeDSAuthenticationGetResponseMerchantRiskIndicatorReorderItems = "REORDERED"
)

// Shipping method that the cardholder chose for the transaction. If purchase
// includes one or more item, this indicator is used for the physical goods; if the
// purchase only includes digital goods, this indicator is used to describe the
// most expensive item purchased. Maps to EMV 3DS field shipIndicator.
type ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod string

const (
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodDigitalGoods               ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "DIGITAL_GOODS"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodLockerDelivery             ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "LOCKER_DELIVERY"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodOther                      ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "OTHER"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodPickUpAndGoDelivery        ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "PICK_UP_AND_GO_DELIVERY"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodShipToBillingAddress       ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "SHIP_TO_BILLING_ADDRESS"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodShipToNonBillingAddress    ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "SHIP_TO_NON_BILLING_ADDRESS"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodShipToOtherVerifiedAddress ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "SHIP_TO_OTHER_VERIFIED_ADDRESS"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodShipToStore                ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "SHIP_TO_STORE"
	ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethodTravelAndEventTickets      ThreeDSAuthenticationGetResponseMerchantRiskIndicatorShippingMethod = "TRAVEL_AND_EVENT_TICKETS"
)

// Either PAYMENT_AUTHENTICATION or NON_PAYMENT_AUTHENTICATION. For
// NON_PAYMENT_AUTHENTICATION, additional_data and transaction fields are not
// populated.
type ThreeDSAuthenticationGetResponseMessageCategory string

const (
	ThreeDSAuthenticationGetResponseMessageCategoryNonPaymentAuthentication ThreeDSAuthenticationGetResponseMessageCategory = "NON_PAYMENT_AUTHENTICATION"
	ThreeDSAuthenticationGetResponseMessageCategoryPaymentAuthentication    ThreeDSAuthenticationGetResponseMessageCategory = "PAYMENT_AUTHENTICATION"
)

// Object containing additional data about the 3DS request that is beyond the EMV
// 3DS standard spec (e.g., specific fields that only certain card networks send
// but are not required across all 3DS requests).
type ThreeDSAuthenticationGetResponseAdditionalData struct {
	// Mastercard only: Indicates whether the network would have considered the
	// authentication request to be low risk or not.
	NetworkDecision ThreeDSAuthenticationGetResponseAdditionalDataNetworkDecision `json:"network_decision,nullable"`
	// Mastercard only: Assessment by the network of the authentication risk level,
	// with a higher value indicating a higher amount of risk.
	NetworkRiskScore float64                                            `json:"network_risk_score,nullable"`
	JSON             threeDSAuthenticationGetResponseAdditionalDataJSON `json:"-"`
}

// threeDSAuthenticationGetResponseAdditionalDataJSON contains the JSON metadata
// for the struct [ThreeDSAuthenticationGetResponseAdditionalData]
type threeDSAuthenticationGetResponseAdditionalDataJSON struct {
	NetworkDecision  apijson.Field
	NetworkRiskScore apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponseAdditionalData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Mastercard only: Indicates whether the network would have considered the
// authentication request to be low risk or not.
type ThreeDSAuthenticationGetResponseAdditionalDataNetworkDecision string

const (
	ThreeDSAuthenticationGetResponseAdditionalDataNetworkDecisionLowRisk    ThreeDSAuthenticationGetResponseAdditionalDataNetworkDecision = "LOW_RISK"
	ThreeDSAuthenticationGetResponseAdditionalDataNetworkDecisionNotLowRisk ThreeDSAuthenticationGetResponseAdditionalDataNetworkDecision = "NOT_LOW_RISK"
)

// Object containing data about the app used in the e-commerce transaction. Present
// if the channel is 'APP_BASED'.
type ThreeDSAuthenticationGetResponseApp struct {
	// Device information gathered from the cardholder's device - JSON name/value pairs
	// that is Base64url encoded. Maps to EMV 3DS field deviceInfo.
	DeviceInfo string `json:"device_info,nullable"`
	// External IP address used by the app generating the 3DS authentication request.
	// Maps to EMV 3DS field appIp.
	Ip   string                                  `json:"ip"`
	JSON threeDSAuthenticationGetResponseAppJSON `json:"-"`
}

// threeDSAuthenticationGetResponseAppJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationGetResponseApp]
type threeDSAuthenticationGetResponseAppJSON struct {
	DeviceInfo  apijson.Field
	Ip          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of authentication request - i.e., the type of transaction or interaction is
// causing the merchant to request an authentication. Maps to EMV 3DS field
// threeDSRequestorAuthenticationInd.
type ThreeDSAuthenticationGetResponseAuthenticationRequestType string

const (
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeAddCard                        ThreeDSAuthenticationGetResponseAuthenticationRequestType = "ADD_CARD"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeBillingAgreement               ThreeDSAuthenticationGetResponseAuthenticationRequestType = "BILLING_AGREEMENT"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeDelayedShipment                ThreeDSAuthenticationGetResponseAuthenticationRequestType = "DELAYED_SHIPMENT"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeEmvTokenCardholderVerification ThreeDSAuthenticationGetResponseAuthenticationRequestType = "EMV_TOKEN_CARDHOLDER_VERIFICATION"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeInstallmentTransaction         ThreeDSAuthenticationGetResponseAuthenticationRequestType = "INSTALLMENT_TRANSACTION"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeMaintainCard                   ThreeDSAuthenticationGetResponseAuthenticationRequestType = "MAINTAIN_CARD"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypePaymentTransaction             ThreeDSAuthenticationGetResponseAuthenticationRequestType = "PAYMENT_TRANSACTION"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeRecurringTransaction           ThreeDSAuthenticationGetResponseAuthenticationRequestType = "RECURRING_TRANSACTION"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeSplitPayment                   ThreeDSAuthenticationGetResponseAuthenticationRequestType = "SPLIT_PAYMENT"
	ThreeDSAuthenticationGetResponseAuthenticationRequestTypeSplitShipment                  ThreeDSAuthenticationGetResponseAuthenticationRequestType = "SPLIT_SHIPMENT"
)

// Object containing data about the browser used in the e-commerce transaction.
// Present if the channel is 'BROWSER'.
type ThreeDSAuthenticationGetResponseBrowser struct {
	// IP address of the browser as returned by the HTTP headers to the 3DS requestor
	// (e.g., merchant or digital wallet). Maps to EMV 3DS field browserIP.
	Ip string `json:"ip,nullable"`
	// Indicates whether the cardholder's browser has the ability to execute Java. Maps
	// to EMV 3DS field browserJavaEnabled.
	JavaEnabled bool `json:"java_enabled,nullable"`
	// Indicates whether the cardholder's browser has the ability to execute
	// JavaScript. Maps to EMV 3DS field browserJavascriptEnabled.
	JavascriptEnabled bool `json:"javascript_enabled,nullable"`
	// Language of the cardholder's browser as defined in IETF BCP47. Maps to EMV 3DS
	// field browserLanguage.
	Language string `json:"language,nullable"`
	// Time zone of the cardholder's browser offset in minutes between UTC and the
	// cardholder browser's local time. The offset is positive if the local time is
	// behind UTC and negative if it is ahead. Maps to EMV 3DS field browserTz.
	TimeZone string `json:"time_zone,nullable"`
	// Content of the HTTP user-agent header. Maps to EMV 3DS field browserUserAgent.
	UserAgent string                                      `json:"user_agent,nullable"`
	JSON      threeDSAuthenticationGetResponseBrowserJSON `json:"-"`
}

// threeDSAuthenticationGetResponseBrowserJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationGetResponseBrowser]
type threeDSAuthenticationGetResponseBrowserJSON struct {
	Ip                apijson.Field
	JavaEnabled       apijson.Field
	JavascriptEnabled apijson.Field
	Language          apijson.Field
	TimeZone          apijson.Field
	UserAgent         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponseBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of 3DS Requestor Initiated (3RI) request i.e., a 3DS authentication that
// takes place at the initiation of the merchant rather than the cardholder. The
// most common example of this is where a merchant is authenticating before billing
// for a recurring transaction such as a pay TV subscription or a utility bill.
// Maps to EMV 3DS field threeRIInd.
type ThreeDSAuthenticationGetResponseThreeRiRequestType string

const (
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeAccountVerification         ThreeDSAuthenticationGetResponseThreeRiRequestType = "ACCOUNT_VERIFICATION"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeAddCard                     ThreeDSAuthenticationGetResponseThreeRiRequestType = "ADD_CARD"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeBillingAgreement            ThreeDSAuthenticationGetResponseThreeRiRequestType = "BILLING_AGREEMENT"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeCardSecurityCodeStatusCheck ThreeDSAuthenticationGetResponseThreeRiRequestType = "CARD_SECURITY_CODE_STATUS_CHECK"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeDelayedShipment             ThreeDSAuthenticationGetResponseThreeRiRequestType = "DELAYED_SHIPMENT"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeDeviceBindingStatusCheck    ThreeDSAuthenticationGetResponseThreeRiRequestType = "DEVICE_BINDING_STATUS_CHECK"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeInstallmentTransaction      ThreeDSAuthenticationGetResponseThreeRiRequestType = "INSTALLMENT_TRANSACTION"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeMailOrder                   ThreeDSAuthenticationGetResponseThreeRiRequestType = "MAIL_ORDER"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeMaintainCardInfo            ThreeDSAuthenticationGetResponseThreeRiRequestType = "MAINTAIN_CARD_INFO"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeOtherPayment                ThreeDSAuthenticationGetResponseThreeRiRequestType = "OTHER_PAYMENT"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeRecurringTransaction        ThreeDSAuthenticationGetResponseThreeRiRequestType = "RECURRING_TRANSACTION"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeSplitPayment                ThreeDSAuthenticationGetResponseThreeRiRequestType = "SPLIT_PAYMENT"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeSplitShipment               ThreeDSAuthenticationGetResponseThreeRiRequestType = "SPLIT_SHIPMENT"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeTelephoneOrder              ThreeDSAuthenticationGetResponseThreeRiRequestType = "TELEPHONE_ORDER"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeTopUp                       ThreeDSAuthenticationGetResponseThreeRiRequestType = "TOP_UP"
	ThreeDSAuthenticationGetResponseThreeRiRequestTypeTrustListStatusCheck        ThreeDSAuthenticationGetResponseThreeRiRequestType = "TRUST_LIST_STATUS_CHECK"
)

// Object containing data about the e-commerce transaction for which the merchant
// is requesting authentication.
type ThreeDSAuthenticationGetResponseTransaction struct {
	// Amount of the purchase in minor units of currency with all punctuation removed.
	// Maps to EMV 3DS field purchaseAmount.
	Amount float64 `json:"amount,required"`
	// Currency of the purchase. Maps to EMV 3DS field purchaseCurrency.
	Currency string `json:"currency,required"`
	// Minor units of currency, as specified in ISO 4217 currency exponent. Maps to EMV
	// 3DS field purchaseExponent.
	CurrencyExponent float64 `json:"currency_exponent,required"`
	// Date and time when the authentication was generated by the merchant/acquirer's
	// 3DS server. Maps to EMV 3DS field purchaseDate.
	DateTime time.Time `json:"date_time,required" format:"date-time"`
	// Type of the transaction for which a 3DS authentication request is occurring.
	// Maps to EMV 3DS field transType.
	Type ThreeDSAuthenticationGetResponseTransactionType `json:"type,required,nullable"`
	JSON threeDSAuthenticationGetResponseTransactionJSON `json:"-"`
}

// threeDSAuthenticationGetResponseTransactionJSON contains the JSON metadata for
// the struct [ThreeDSAuthenticationGetResponseTransaction]
type threeDSAuthenticationGetResponseTransactionJSON struct {
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	DateTime         apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreeDSAuthenticationGetResponseTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the transaction for which a 3DS authentication request is occurring.
// Maps to EMV 3DS field transType.
type ThreeDSAuthenticationGetResponseTransactionType string

const (
	ThreeDSAuthenticationGetResponseTransactionTypeAccountFunding           ThreeDSAuthenticationGetResponseTransactionType = "ACCOUNT_FUNDING"
	ThreeDSAuthenticationGetResponseTransactionTypeCheckAcceptance          ThreeDSAuthenticationGetResponseTransactionType = "CHECK_ACCEPTANCE"
	ThreeDSAuthenticationGetResponseTransactionTypeGoodsServicePurchase     ThreeDSAuthenticationGetResponseTransactionType = "GOODS_SERVICE_PURCHASE"
	ThreeDSAuthenticationGetResponseTransactionTypePrepaidActivationAndLoad ThreeDSAuthenticationGetResponseTransactionType = "PREPAID_ACTIVATION_AND_LOAD"
	ThreeDSAuthenticationGetResponseTransactionTypeQuasiCashTransaction     ThreeDSAuthenticationGetResponseTransactionType = "QUASI_CASH_TRANSACTION"
)

type ThreeDSAuthenticationSimulateResponse struct {
	// A unique token to reference this transaction with later calls to void or clear
	// the authorization.
	Token string `json:"token" format:"uuid"`
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string                                    `json:"debugging_request_id" format:"uuid"`
	JSON               threeDSAuthenticationSimulateResponseJSON `json:"-"`
}

// threeDSAuthenticationSimulateResponseJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationSimulateResponse]
type threeDSAuthenticationSimulateResponseJSON struct {
	Token              apijson.Field
	DebuggingRequestID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ThreeDSAuthenticationSimulateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ThreeDSAuthenticationSimulateParams struct {
	Merchant param.Field[ThreeDSAuthenticationSimulateParamsMerchant] `json:"merchant,required"`
	// Sixteen digit card number.
	Pan         param.Field[string]                                         `json:"pan,required"`
	Transaction param.Field[ThreeDSAuthenticationSimulateParamsTransaction] `json:"transaction,required"`
}

func (r ThreeDSAuthenticationSimulateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreeDSAuthenticationSimulateParamsMerchant struct {
	// Unique identifier to identify the payment card acceptor. Corresponds to
	// `merchant_acceptor_id` in authorization.
	ID param.Field[string] `json:"id,required"`
	// Country of the address provided by the cardholder in ISO 3166-1 alpha-3 format
	// (e.g. USA)
	Country param.Field[string] `json:"country,required"`
	// Merchant category code for the transaction to be simulated. A four-digit number
	// listed in ISO 18245. Supported merchant category codes can be found
	// [here](https://docs.lithic.com/docs/transactions#merchant-category-codes-mccs).
	Mcc param.Field[string] `json:"mcc,required"`
	// Merchant descriptor, corresponds to `descriptor` in authorization.
	Name param.Field[string] `json:"name,required"`
}

func (r ThreeDSAuthenticationSimulateParamsMerchant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreeDSAuthenticationSimulateParamsTransaction struct {
	// Amount (in cents) to authenticate.
	Amount param.Field[int64] `json:"amount,required"`
	// 3-digit alphabetic ISO 4217 currency code.
	Currency param.Field[string] `json:"currency,required"`
}

func (r ThreeDSAuthenticationSimulateParamsTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
