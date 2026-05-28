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
	"github.com/lithic-com/lithic-go/shared"
)

// CardAuthorizationService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardAuthorizationService] method instead.
type CardAuthorizationService struct {
	Options []option.RequestOption
}

// NewCardAuthorizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCardAuthorizationService(opts ...option.RequestOption) (r *CardAuthorizationService) {
	r = &CardAuthorizationService{}
	r.Options = opts
	return
}

// Card program's response to Authorization Challenge. Programs that have
// Authorization Challenges configured as Out of Band receive a
// [card_authorization.challenge](https://docs.lithic.com/reference/post_card-authorization-challenge)
// webhook when an authorization attempt triggers a challenge. The card program
// should respond using this endpoint after the cardholder completes the challenge.
func (r *CardAuthorizationService) ChallengeResponse(ctx context.Context, eventToken string, body CardAuthorizationChallengeResponseParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if eventToken == "" {
		err = errors.New("missing required event_token parameter")
		return err
	}
	path := fmt.Sprintf("v1/card_authorizations/%s/challenge_response", eventToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Card Authorization
type CardAuthorization struct {
	// The provisional transaction group uuid associated with the authorization
	Token string `json:"token" api:"required" format:"uuid"`
	// Fee (in cents) assessed by the merchant and paid for by the cardholder. Will be
	// zero if no fee is assessed. Rebates may be transmitted as a negative value to
	// indicate credited fees.
	AcquirerFee int64 `json:"acquirer_fee" api:"required"`
	// Deprecated, use `amounts`. Authorization amount of the transaction (in cents),
	// including any acquirer fees. The contents of this field are identical to
	// `authorization_amount`.
	//
	// Deprecated: deprecated
	Amount int64 `json:"amount" api:"required"`
	// Structured amounts for this authorization. The `cardholder` and `merchant`
	// amounts reflect the original network authorization values. For programs with
	// hold adjustments enabled (e.g., automated fuel dispensers or tipping MCCs), the
	// `hold` amount may exceed the `cardholder` and `merchant` amounts to account for
	// anticipated final transaction amounts such as tips or fuel fill-ups
	Amounts CardAuthorizationAmounts `json:"amounts" api:"required"`
	// Deprecated, use `amounts`. The base transaction amount (in cents) plus the
	// acquirer fee field. This is the amount the issuer should authorize against
	// unless the issuer is paying the acquirer fee on behalf of the cardholder.
	//
	// Deprecated: deprecated
	AuthorizationAmount int64                `json:"authorization_amount" api:"required"`
	Avs                 CardAuthorizationAvs `json:"avs" api:"required"`
	// Card object in ASA
	Card CardAuthorizationCard `json:"card" api:"required"`
	// Deprecated, use `amounts`. 3-character alphabetic ISO 4217 code for cardholder's
	// billing currency.
	//
	// Deprecated: deprecated
	CardholderCurrency string `json:"cardholder_currency" api:"required"`
	// The portion of the transaction requested as cash back by the cardholder, and
	// does not include any acquirer fees. The amount field includes the purchase
	// amount, the requested cash back amount, and any acquirer fees.
	//
	// If no cash back was requested, the value of this field will be 0, and the field
	// will always be present.
	CashAmount int64 `json:"cash_amount" api:"required"`
	// Date and time when the transaction first occurred in UTC.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Merchant information including full location details.
	Merchant CardAuthorizationMerchant `json:"merchant" api:"required"`
	// Deprecated, use `amounts`. The amount that the merchant will receive,
	// denominated in `merchant_currency` and in the smallest currency unit. Note the
	// amount includes `acquirer_fee`, similar to `authorization_amount`. It will be
	// different from `authorization_amount` if the merchant is taking payment in a
	// different currency.
	//
	// Deprecated: deprecated
	MerchantAmount int64 `json:"merchant_amount" api:"required"`
	// 3-character alphabetic ISO 4217 code for the local currency of the transaction.
	//
	// Deprecated: deprecated
	MerchantCurrency string `json:"merchant_currency" api:"required"`
	// Where the cardholder received the service, when different from the card acceptor
	// location. This is populated from network data elements such as Mastercard DE-122
	// SE1 SF9-14 and Visa F34 DS02.
	ServiceLocation CardAuthorizationServiceLocation `json:"service_location" api:"required,nullable"`
	// Deprecated, use `amounts`. Amount (in cents) of the transaction that has been
	// settled, including any acquirer fees.
	//
	// Deprecated: deprecated
	SettledAmount int64 `json:"settled_amount" api:"required"`
	// The type of authorization request that this request is for. Note that
	// `CREDIT_AUTHORIZATION` and `FINANCIAL_CREDIT_AUTHORIZATION` is only available to
	// users with credit decisioning via ASA enabled.
	Status CardAuthorizationStatus `json:"status" api:"required"`
	// The entity that initiated the transaction.
	TransactionInitiator     CardAuthorizationTransactionInitiator `json:"transaction_initiator" api:"required"`
	AccountType              CardAuthorizationAccountType          `json:"account_type"`
	CardholderAuthentication CardholderAuthentication              `json:"cardholder_authentication"`
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
	FleetInfo CardAuthorizationFleetInfo `json:"fleet_info" api:"nullable"`
	// The latest Authorization Challenge that was issued to the cardholder for this
	// merchant.
	LatestChallenge CardAuthorizationLatestChallenge `json:"latest_challenge"`
	// Card network of the authorization.
	Network CardAuthorizationNetwork `json:"network"`
	// Network-provided score assessing risk level associated with a given
	// authorization. Scores are on a range of 0-999, with 0 representing the lowest
	// risk and 999 representing the highest risk. For Visa transactions, where the raw
	// score has a range of 0-99, Lithic will normalize the score by multiplying the
	// raw score by 10x.
	NetworkRiskScore int64 `json:"network_risk_score" api:"nullable"`
	// Contains raw data provided by the card network, including attributes that
	// provide further context about the authorization. If populated by the network,
	// data is organized by Lithic and passed through without further modification.
	// Please consult the official network documentation for more details about these
	// values and how to use them. This object is only available to certain programs-
	// contact your Customer Success Manager to discuss enabling access.
	NetworkSpecificData CardAuthorizationNetworkSpecificData `json:"network_specific_data" api:"nullable"`
	Pos                 CardAuthorizationPos                 `json:"pos"`
	TokenInfo           TokenInfo                            `json:"token_info" api:"nullable"`
	// Deprecated: approximate time-to-live for the authorization.
	Ttl  time.Time             `json:"ttl" format:"date-time"`
	JSON cardAuthorizationJSON `json:"-"`
}

// cardAuthorizationJSON contains the JSON metadata for the struct
// [CardAuthorization]
type cardAuthorizationJSON struct {
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
	Merchant                 apijson.Field
	MerchantAmount           apijson.Field
	MerchantCurrency         apijson.Field
	ServiceLocation          apijson.Field
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

func (r *CardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Structured amounts for this authorization. The `cardholder` and `merchant`
// amounts reflect the original network authorization values. For programs with
// hold adjustments enabled (e.g., automated fuel dispensers or tipping MCCs), the
// `hold` amount may exceed the `cardholder` and `merchant` amounts to account for
// anticipated final transaction amounts such as tips or fuel fill-ups
type CardAuthorizationAmounts struct {
	Cardholder CardAuthorizationAmountsCardholder `json:"cardholder" api:"required"`
	Hold       CardAuthorizationAmountsHold       `json:"hold" api:"required,nullable"`
	Merchant   CardAuthorizationAmountsMerchant   `json:"merchant" api:"required"`
	Settlement CardAuthorizationAmountsSettlement `json:"settlement" api:"required,nullable"`
	JSON       cardAuthorizationAmountsJSON       `json:"-"`
}

// cardAuthorizationAmountsJSON contains the JSON metadata for the struct
// [CardAuthorizationAmounts]
type cardAuthorizationAmountsJSON struct {
	Cardholder  apijson.Field
	Hold        apijson.Field
	Merchant    apijson.Field
	Settlement  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationAmountsJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationAmountsCardholder struct {
	// Amount in the smallest unit of the applicable currency (e.g., cents)
	Amount int64 `json:"amount" api:"required"`
	// Exchange rate used for currency conversion
	ConversionRate string `json:"conversion_rate" api:"required"`
	// 3-character alphabetic ISO 4217 currency
	Currency shared.Currency                        `json:"currency" api:"required"`
	JSON     cardAuthorizationAmountsCardholderJSON `json:"-"`
}

// cardAuthorizationAmountsCardholderJSON contains the JSON metadata for the struct
// [CardAuthorizationAmountsCardholder]
type cardAuthorizationAmountsCardholderJSON struct {
	Amount         apijson.Field
	ConversionRate apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardAuthorizationAmountsCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationAmountsCardholderJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationAmountsHold struct {
	// Amount in the smallest unit of the applicable currency (e.g., cents)
	Amount int64 `json:"amount" api:"required"`
	// 3-character alphabetic ISO 4217 currency
	Currency shared.Currency                  `json:"currency" api:"required"`
	JSON     cardAuthorizationAmountsHoldJSON `json:"-"`
}

// cardAuthorizationAmountsHoldJSON contains the JSON metadata for the struct
// [CardAuthorizationAmountsHold]
type cardAuthorizationAmountsHoldJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationAmountsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationAmountsHoldJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationAmountsMerchant struct {
	// Amount in the smallest unit of the applicable currency (e.g., cents)
	Amount int64 `json:"amount" api:"required"`
	// 3-character alphabetic ISO 4217 currency
	Currency shared.Currency                      `json:"currency" api:"required"`
	JSON     cardAuthorizationAmountsMerchantJSON `json:"-"`
}

// cardAuthorizationAmountsMerchantJSON contains the JSON metadata for the struct
// [CardAuthorizationAmountsMerchant]
type cardAuthorizationAmountsMerchantJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationAmountsMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationAmountsMerchantJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationAmountsSettlement struct {
	// Amount in the smallest unit of the applicable currency (e.g., cents)
	Amount int64 `json:"amount" api:"required"`
	// 3-character alphabetic ISO 4217 currency
	Currency shared.Currency                        `json:"currency" api:"required"`
	JSON     cardAuthorizationAmountsSettlementJSON `json:"-"`
}

// cardAuthorizationAmountsSettlementJSON contains the JSON metadata for the struct
// [CardAuthorizationAmountsSettlement]
type cardAuthorizationAmountsSettlementJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationAmountsSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationAmountsSettlementJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationAvs struct {
	// Cardholder address
	Address string `json:"address" api:"required"`
	// Lithic's evaluation result comparing the transaction's address data with the
	// cardholder KYC data if it exists. In the event Lithic does not have any
	// Cardholder KYC data, or the transaction does not contain any address data,
	// NOT_PRESENT will be returned
	AddressOnFileMatch CardAuthorizationAvsAddressOnFileMatch `json:"address_on_file_match" api:"required"`
	// Cardholder ZIP code
	Zipcode string                   `json:"zipcode" api:"required"`
	JSON    cardAuthorizationAvsJSON `json:"-"`
}

// cardAuthorizationAvsJSON contains the JSON metadata for the struct
// [CardAuthorizationAvs]
type cardAuthorizationAvsJSON struct {
	Address            apijson.Field
	AddressOnFileMatch apijson.Field
	Zipcode            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationAvs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationAvsJSON) RawJSON() string {
	return r.raw
}

// Lithic's evaluation result comparing the transaction's address data with the
// cardholder KYC data if it exists. In the event Lithic does not have any
// Cardholder KYC data, or the transaction does not contain any address data,
// NOT_PRESENT will be returned
type CardAuthorizationAvsAddressOnFileMatch string

const (
	CardAuthorizationAvsAddressOnFileMatchMatch            CardAuthorizationAvsAddressOnFileMatch = "MATCH"
	CardAuthorizationAvsAddressOnFileMatchMatchAddressOnly CardAuthorizationAvsAddressOnFileMatch = "MATCH_ADDRESS_ONLY"
	CardAuthorizationAvsAddressOnFileMatchMatchZipOnly     CardAuthorizationAvsAddressOnFileMatch = "MATCH_ZIP_ONLY"
	CardAuthorizationAvsAddressOnFileMatchMismatch         CardAuthorizationAvsAddressOnFileMatch = "MISMATCH"
	CardAuthorizationAvsAddressOnFileMatchNotPresent       CardAuthorizationAvsAddressOnFileMatch = "NOT_PRESENT"
)

func (r CardAuthorizationAvsAddressOnFileMatch) IsKnown() bool {
	switch r {
	case CardAuthorizationAvsAddressOnFileMatchMatch, CardAuthorizationAvsAddressOnFileMatchMatchAddressOnly, CardAuthorizationAvsAddressOnFileMatchMatchZipOnly, CardAuthorizationAvsAddressOnFileMatchMismatch, CardAuthorizationAvsAddressOnFileMatchNotPresent:
		return true
	}
	return false
}

// Card object in ASA
type CardAuthorizationCard struct {
	// Globally unique identifier for the card.
	Token string `json:"token" api:"required" format:"uuid"`
	// Last four digits of the card number
	LastFour string `json:"last_four" api:"required"`
	// Customizable name to identify the card
	Memo string `json:"memo" api:"required"`
	// Amount (in cents) to limit approved authorizations. Purchase requests above the
	// spend limit will be declined (refunds and credits will be approved).
	//
	// Note that while spend limits are enforced based on authorized and settled volume
	// on a card, they are not recommended to be used for balance or
	// reconciliation-level accuracy. Spend limits also cannot block force posted
	// charges (i.e., when a merchant sends a clearing message without a prior
	// authorization).
	SpendLimit int64 `json:"spend_limit" api:"required"`
	// Note that to support recurring monthly payments, which can occur on different
	// day every month, the time window we consider for MONTHLY velocity starts 6 days
	// after the current calendar date one month prior.
	SpendLimitDuration CardAuthorizationCardSpendLimitDuration `json:"spend_limit_duration" api:"required"`
	State              CardAuthorizationCardState              `json:"state" api:"required"`
	Type               CardAuthorizationCardType               `json:"type" api:"required"`
	JSON               cardAuthorizationCardJSON               `json:"-"`
}

// cardAuthorizationCardJSON contains the JSON metadata for the struct
// [CardAuthorizationCard]
type cardAuthorizationCardJSON struct {
	Token              apijson.Field
	LastFour           apijson.Field
	Memo               apijson.Field
	SpendLimit         apijson.Field
	SpendLimitDuration apijson.Field
	State              apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationCardJSON) RawJSON() string {
	return r.raw
}

// Note that to support recurring monthly payments, which can occur on different
// day every month, the time window we consider for MONTHLY velocity starts 6 days
// after the current calendar date one month prior.
type CardAuthorizationCardSpendLimitDuration string

const (
	CardAuthorizationCardSpendLimitDurationAnnually    CardAuthorizationCardSpendLimitDuration = "ANNUALLY"
	CardAuthorizationCardSpendLimitDurationForever     CardAuthorizationCardSpendLimitDuration = "FOREVER"
	CardAuthorizationCardSpendLimitDurationMonthly     CardAuthorizationCardSpendLimitDuration = "MONTHLY"
	CardAuthorizationCardSpendLimitDurationTransaction CardAuthorizationCardSpendLimitDuration = "TRANSACTION"
)

func (r CardAuthorizationCardSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardAuthorizationCardSpendLimitDurationAnnually, CardAuthorizationCardSpendLimitDurationForever, CardAuthorizationCardSpendLimitDurationMonthly, CardAuthorizationCardSpendLimitDurationTransaction:
		return true
	}
	return false
}

type CardAuthorizationCardState string

const (
	CardAuthorizationCardStateClosed             CardAuthorizationCardState = "CLOSED"
	CardAuthorizationCardStateOpen               CardAuthorizationCardState = "OPEN"
	CardAuthorizationCardStatePaused             CardAuthorizationCardState = "PAUSED"
	CardAuthorizationCardStatePendingActivation  CardAuthorizationCardState = "PENDING_ACTIVATION"
	CardAuthorizationCardStatePendingFulfillment CardAuthorizationCardState = "PENDING_FULFILLMENT"
)

func (r CardAuthorizationCardState) IsKnown() bool {
	switch r {
	case CardAuthorizationCardStateClosed, CardAuthorizationCardStateOpen, CardAuthorizationCardStatePaused, CardAuthorizationCardStatePendingActivation, CardAuthorizationCardStatePendingFulfillment:
		return true
	}
	return false
}

type CardAuthorizationCardType string

const (
	CardAuthorizationCardTypeSingleUse      CardAuthorizationCardType = "SINGLE_USE"
	CardAuthorizationCardTypeMerchantLocked CardAuthorizationCardType = "MERCHANT_LOCKED"
	CardAuthorizationCardTypeUnlocked       CardAuthorizationCardType = "UNLOCKED"
	CardAuthorizationCardTypePhysical       CardAuthorizationCardType = "PHYSICAL"
	CardAuthorizationCardTypeDigitalWallet  CardAuthorizationCardType = "DIGITAL_WALLET"
	CardAuthorizationCardTypeVirtual        CardAuthorizationCardType = "VIRTUAL"
)

func (r CardAuthorizationCardType) IsKnown() bool {
	switch r {
	case CardAuthorizationCardTypeSingleUse, CardAuthorizationCardTypeMerchantLocked, CardAuthorizationCardTypeUnlocked, CardAuthorizationCardTypePhysical, CardAuthorizationCardTypeDigitalWallet, CardAuthorizationCardTypeVirtual:
		return true
	}
	return false
}

// Merchant information including full location details.
type CardAuthorizationMerchant struct {
	// Phone number of card acceptor.
	PhoneNumber string `json:"phone_number" api:"required,nullable"`
	// Postal code of card acceptor.
	PostalCode string `json:"postal_code" api:"required,nullable"`
	// Street address of card acceptor.
	StreetAddress string                        `json:"street_address" api:"required,nullable"`
	JSON          cardAuthorizationMerchantJSON `json:"-"`
	shared.Merchant
}

// cardAuthorizationMerchantJSON contains the JSON metadata for the struct
// [CardAuthorizationMerchant]
type cardAuthorizationMerchantJSON struct {
	PhoneNumber   apijson.Field
	PostalCode    apijson.Field
	StreetAddress apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardAuthorizationMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationMerchantJSON) RawJSON() string {
	return r.raw
}

// Where the cardholder received the service, when different from the card acceptor
// location. This is populated from network data elements such as Mastercard DE-122
// SE1 SF9-14 and Visa F34 DS02.
type CardAuthorizationServiceLocation struct {
	// City of service location.
	City string `json:"city" api:"required,nullable"`
	// Country code of service location, ISO 3166-1 alpha-3.
	Country string `json:"country" api:"required,nullable"`
	// Postal code of service location.
	PostalCode string `json:"postal_code" api:"required,nullable"`
	// State/province code of service location, ISO 3166-2.
	State string `json:"state" api:"required,nullable"`
	// Street address of service location.
	StreetAddress string                               `json:"street_address" api:"required,nullable"`
	JSON          cardAuthorizationServiceLocationJSON `json:"-"`
}

// cardAuthorizationServiceLocationJSON contains the JSON metadata for the struct
// [CardAuthorizationServiceLocation]
type cardAuthorizationServiceLocationJSON struct {
	City          apijson.Field
	Country       apijson.Field
	PostalCode    apijson.Field
	State         apijson.Field
	StreetAddress apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardAuthorizationServiceLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationServiceLocationJSON) RawJSON() string {
	return r.raw
}

// The type of authorization request that this request is for. Note that
// `CREDIT_AUTHORIZATION` and `FINANCIAL_CREDIT_AUTHORIZATION` is only available to
// users with credit decisioning via ASA enabled.
type CardAuthorizationStatus string

const (
	CardAuthorizationStatusAuthorization                CardAuthorizationStatus = "AUTHORIZATION"
	CardAuthorizationStatusCreditAuthorization          CardAuthorizationStatus = "CREDIT_AUTHORIZATION"
	CardAuthorizationStatusFinancialAuthorization       CardAuthorizationStatus = "FINANCIAL_AUTHORIZATION"
	CardAuthorizationStatusFinancialCreditAuthorization CardAuthorizationStatus = "FINANCIAL_CREDIT_AUTHORIZATION"
	CardAuthorizationStatusBalanceInquiry               CardAuthorizationStatus = "BALANCE_INQUIRY"
)

func (r CardAuthorizationStatus) IsKnown() bool {
	switch r {
	case CardAuthorizationStatusAuthorization, CardAuthorizationStatusCreditAuthorization, CardAuthorizationStatusFinancialAuthorization, CardAuthorizationStatusFinancialCreditAuthorization, CardAuthorizationStatusBalanceInquiry:
		return true
	}
	return false
}

// The entity that initiated the transaction.
type CardAuthorizationTransactionInitiator string

const (
	CardAuthorizationTransactionInitiatorCardholder CardAuthorizationTransactionInitiator = "CARDHOLDER"
	CardAuthorizationTransactionInitiatorMerchant   CardAuthorizationTransactionInitiator = "MERCHANT"
	CardAuthorizationTransactionInitiatorUnknown    CardAuthorizationTransactionInitiator = "UNKNOWN"
)

func (r CardAuthorizationTransactionInitiator) IsKnown() bool {
	switch r {
	case CardAuthorizationTransactionInitiatorCardholder, CardAuthorizationTransactionInitiatorMerchant, CardAuthorizationTransactionInitiatorUnknown:
		return true
	}
	return false
}

type CardAuthorizationAccountType string

const (
	CardAuthorizationAccountTypeChecking CardAuthorizationAccountType = "CHECKING"
	CardAuthorizationAccountTypeSavings  CardAuthorizationAccountType = "SAVINGS"
)

func (r CardAuthorizationAccountType) IsKnown() bool {
	switch r {
	case CardAuthorizationAccountTypeChecking, CardAuthorizationAccountTypeSavings:
		return true
	}
	return false
}

// Optional Object containing information if the Card is a part of a Fleet managed
// program
type CardAuthorizationFleetInfo struct {
	// Code indicating what the driver was prompted to enter at time of purchase. This
	// is configured at a program level and is a static configuration, and does not
	// change on a request to request basis
	FleetPromptCode CardAuthorizationFleetInfoFleetPromptCode `json:"fleet_prompt_code" api:"required"`
	// Code indicating which restrictions, if any, there are on purchase. This is
	// configured at a program level and is a static configuration, and does not change
	// on a request to request basis
	FleetRestrictionCode CardAuthorizationFleetInfoFleetRestrictionCode `json:"fleet_restriction_code" api:"required"`
	// Number representing the driver
	DriverNumber string `json:"driver_number" api:"nullable"`
	// Number associated with the vehicle
	VehicleNumber string                         `json:"vehicle_number" api:"nullable"`
	JSON          cardAuthorizationFleetInfoJSON `json:"-"`
}

// cardAuthorizationFleetInfoJSON contains the JSON metadata for the struct
// [CardAuthorizationFleetInfo]
type cardAuthorizationFleetInfoJSON struct {
	FleetPromptCode      apijson.Field
	FleetRestrictionCode apijson.Field
	DriverNumber         apijson.Field
	VehicleNumber        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationFleetInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationFleetInfoJSON) RawJSON() string {
	return r.raw
}

// Code indicating what the driver was prompted to enter at time of purchase. This
// is configured at a program level and is a static configuration, and does not
// change on a request to request basis
type CardAuthorizationFleetInfoFleetPromptCode string

const (
	CardAuthorizationFleetInfoFleetPromptCodeNoPrompt      CardAuthorizationFleetInfoFleetPromptCode = "NO_PROMPT"
	CardAuthorizationFleetInfoFleetPromptCodeVehicleNumber CardAuthorizationFleetInfoFleetPromptCode = "VEHICLE_NUMBER"
	CardAuthorizationFleetInfoFleetPromptCodeDriverNumber  CardAuthorizationFleetInfoFleetPromptCode = "DRIVER_NUMBER"
)

func (r CardAuthorizationFleetInfoFleetPromptCode) IsKnown() bool {
	switch r {
	case CardAuthorizationFleetInfoFleetPromptCodeNoPrompt, CardAuthorizationFleetInfoFleetPromptCodeVehicleNumber, CardAuthorizationFleetInfoFleetPromptCodeDriverNumber:
		return true
	}
	return false
}

// Code indicating which restrictions, if any, there are on purchase. This is
// configured at a program level and is a static configuration, and does not change
// on a request to request basis
type CardAuthorizationFleetInfoFleetRestrictionCode string

const (
	CardAuthorizationFleetInfoFleetRestrictionCodeNoRestrictions CardAuthorizationFleetInfoFleetRestrictionCode = "NO_RESTRICTIONS"
	CardAuthorizationFleetInfoFleetRestrictionCodeFuelOnly       CardAuthorizationFleetInfoFleetRestrictionCode = "FUEL_ONLY"
)

func (r CardAuthorizationFleetInfoFleetRestrictionCode) IsKnown() bool {
	switch r {
	case CardAuthorizationFleetInfoFleetRestrictionCodeNoRestrictions, CardAuthorizationFleetInfoFleetRestrictionCodeFuelOnly:
		return true
	}
	return false
}

// The latest Authorization Challenge that was issued to the cardholder for this
// merchant.
type CardAuthorizationLatestChallenge struct {
	// The date and time when the Authorization Challenge was completed in UTC. Filled
	// only if the challenge has been completed.
	CompletedAt time.Time `json:"completed_at" api:"required,nullable" format:"date-time"`
	// The date and time when the Authorization Challenge was created in UTC
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The method used to deliver the challenge to the cardholder
	//
	// - `SMS` - Challenge was delivered via SMS
	// - `OUT_OF_BAND` - Challenge was delivered via an out-of-band method
	Method CardAuthorizationLatestChallengeMethod `json:"method" api:"required"`
	// The phone number used for sending the Authorization Challenge. Present only when
	// the challenge method is `SMS`.
	PhoneNumber string `json:"phone_number" api:"required,nullable"`
	// The status of the Authorization Challenge
	//
	// - `COMPLETED` - Challenge was successfully completed by the cardholder
	// - `DECLINED` - Challenge was declined by the cardholder
	// - `PENDING` - Challenge is still open
	// - `EXPIRED` - Challenge has expired without being completed
	// - `ERROR` - There was an error processing the challenge
	Status CardAuthorizationLatestChallengeStatus `json:"status" api:"required"`
	JSON   cardAuthorizationLatestChallengeJSON   `json:"-"`
}

// cardAuthorizationLatestChallengeJSON contains the JSON metadata for the struct
// [CardAuthorizationLatestChallenge]
type cardAuthorizationLatestChallengeJSON struct {
	CompletedAt apijson.Field
	Created     apijson.Field
	Method      apijson.Field
	PhoneNumber apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationLatestChallenge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationLatestChallengeJSON) RawJSON() string {
	return r.raw
}

// The method used to deliver the challenge to the cardholder
//
// - `SMS` - Challenge was delivered via SMS
// - `OUT_OF_BAND` - Challenge was delivered via an out-of-band method
type CardAuthorizationLatestChallengeMethod string

const (
	CardAuthorizationLatestChallengeMethodSMS       CardAuthorizationLatestChallengeMethod = "SMS"
	CardAuthorizationLatestChallengeMethodOutOfBand CardAuthorizationLatestChallengeMethod = "OUT_OF_BAND"
)

func (r CardAuthorizationLatestChallengeMethod) IsKnown() bool {
	switch r {
	case CardAuthorizationLatestChallengeMethodSMS, CardAuthorizationLatestChallengeMethodOutOfBand:
		return true
	}
	return false
}

// The status of the Authorization Challenge
//
// - `COMPLETED` - Challenge was successfully completed by the cardholder
// - `DECLINED` - Challenge was declined by the cardholder
// - `PENDING` - Challenge is still open
// - `EXPIRED` - Challenge has expired without being completed
// - `ERROR` - There was an error processing the challenge
type CardAuthorizationLatestChallengeStatus string

const (
	CardAuthorizationLatestChallengeStatusCompleted CardAuthorizationLatestChallengeStatus = "COMPLETED"
	CardAuthorizationLatestChallengeStatusDeclined  CardAuthorizationLatestChallengeStatus = "DECLINED"
	CardAuthorizationLatestChallengeStatusPending   CardAuthorizationLatestChallengeStatus = "PENDING"
	CardAuthorizationLatestChallengeStatusExpired   CardAuthorizationLatestChallengeStatus = "EXPIRED"
	CardAuthorizationLatestChallengeStatusError     CardAuthorizationLatestChallengeStatus = "ERROR"
)

func (r CardAuthorizationLatestChallengeStatus) IsKnown() bool {
	switch r {
	case CardAuthorizationLatestChallengeStatusCompleted, CardAuthorizationLatestChallengeStatusDeclined, CardAuthorizationLatestChallengeStatusPending, CardAuthorizationLatestChallengeStatusExpired, CardAuthorizationLatestChallengeStatusError:
		return true
	}
	return false
}

// Card network of the authorization.
type CardAuthorizationNetwork string

const (
	CardAuthorizationNetworkAmex       CardAuthorizationNetwork = "AMEX"
	CardAuthorizationNetworkInterlink  CardAuthorizationNetwork = "INTERLINK"
	CardAuthorizationNetworkMaestro    CardAuthorizationNetwork = "MAESTRO"
	CardAuthorizationNetworkMastercard CardAuthorizationNetwork = "MASTERCARD"
	CardAuthorizationNetworkUnknown    CardAuthorizationNetwork = "UNKNOWN"
	CardAuthorizationNetworkVisa       CardAuthorizationNetwork = "VISA"
)

func (r CardAuthorizationNetwork) IsKnown() bool {
	switch r {
	case CardAuthorizationNetworkAmex, CardAuthorizationNetworkInterlink, CardAuthorizationNetworkMaestro, CardAuthorizationNetworkMastercard, CardAuthorizationNetworkUnknown, CardAuthorizationNetworkVisa:
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
type CardAuthorizationNetworkSpecificData struct {
	Mastercard CardAuthorizationNetworkSpecificDataMastercard `json:"mastercard" api:"nullable"`
	Visa       CardAuthorizationNetworkSpecificDataVisa       `json:"visa" api:"nullable"`
	JSON       cardAuthorizationNetworkSpecificDataJSON       `json:"-"`
}

// cardAuthorizationNetworkSpecificDataJSON contains the JSON metadata for the
// struct [CardAuthorizationNetworkSpecificData]
type cardAuthorizationNetworkSpecificDataJSON struct {
	Mastercard  apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationNetworkSpecificData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationNetworkSpecificDataJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationNetworkSpecificDataMastercard struct {
	// Indicates the electronic commerce security level and UCAF collection.
	EcommerceSecurityLevelIndicator string `json:"ecommerce_security_level_indicator" api:"nullable"`
	// The On-behalf Service performed on the transaction and the results. Contains all
	// applicable, on-behalf service results that were performed on a given
	// transaction.
	OnBehalfServiceResult []CardAuthorizationNetworkSpecificDataMastercardOnBehalfServiceResult `json:"on_behalf_service_result" api:"nullable"`
	// Indicates the type of additional transaction purpose.
	TransactionTypeIdentifier string                                             `json:"transaction_type_identifier" api:"nullable"`
	JSON                      cardAuthorizationNetworkSpecificDataMastercardJSON `json:"-"`
}

// cardAuthorizationNetworkSpecificDataMastercardJSON contains the JSON metadata
// for the struct [CardAuthorizationNetworkSpecificDataMastercard]
type cardAuthorizationNetworkSpecificDataMastercardJSON struct {
	EcommerceSecurityLevelIndicator apijson.Field
	OnBehalfServiceResult           apijson.Field
	TransactionTypeIdentifier       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *CardAuthorizationNetworkSpecificDataMastercard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationNetworkSpecificDataMastercardJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationNetworkSpecificDataMastercardOnBehalfServiceResult struct {
	// Indicates the results of the service processing.
	Result1 string `json:"result_1" api:"required"`
	// Identifies the results of the service processing.
	Result2 string `json:"result_2" api:"required"`
	// Indicates the service performed on the transaction.
	Service string                                                                  `json:"service" api:"required"`
	JSON    cardAuthorizationNetworkSpecificDataMastercardOnBehalfServiceResultJSON `json:"-"`
}

// cardAuthorizationNetworkSpecificDataMastercardOnBehalfServiceResultJSON contains
// the JSON metadata for the struct
// [CardAuthorizationNetworkSpecificDataMastercardOnBehalfServiceResult]
type cardAuthorizationNetworkSpecificDataMastercardOnBehalfServiceResultJSON struct {
	Result1     apijson.Field
	Result2     apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationNetworkSpecificDataMastercardOnBehalfServiceResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationNetworkSpecificDataMastercardOnBehalfServiceResultJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationNetworkSpecificDataVisa struct {
	// Identifies the purpose or category of a transaction, used to classify and
	// process transactions according to Visa’s rules.
	BusinessApplicationIdentifier string                                       `json:"business_application_identifier" api:"nullable"`
	JSON                          cardAuthorizationNetworkSpecificDataVisaJSON `json:"-"`
}

// cardAuthorizationNetworkSpecificDataVisaJSON contains the JSON metadata for the
// struct [CardAuthorizationNetworkSpecificDataVisa]
type cardAuthorizationNetworkSpecificDataVisaJSON struct {
	BusinessApplicationIdentifier apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *CardAuthorizationNetworkSpecificDataVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationNetworkSpecificDataVisaJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationPos struct {
	// POS > Entry Mode object in ASA
	EntryMode CardAuthorizationPosEntryMode `json:"entry_mode"`
	Terminal  CardAuthorizationPosTerminal  `json:"terminal"`
	JSON      cardAuthorizationPosJSON      `json:"-"`
}

// cardAuthorizationPosJSON contains the JSON metadata for the struct
// [CardAuthorizationPos]
type cardAuthorizationPosJSON struct {
	EntryMode   apijson.Field
	Terminal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationPos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationPosJSON) RawJSON() string {
	return r.raw
}

// POS > Entry Mode object in ASA
type CardAuthorizationPosEntryMode struct {
	// Card Presence Indicator
	Card CardAuthorizationPosEntryModeCard `json:"card"`
	// Cardholder Presence Indicator
	Cardholder CardAuthorizationPosEntryModeCardholder `json:"cardholder"`
	// Method of entry for the PAN
	Pan CardAuthorizationPosEntryModePan `json:"pan"`
	// Indicates whether the cardholder entered the PIN. True if the PIN was entered.
	PinEntered bool                              `json:"pin_entered"`
	JSON       cardAuthorizationPosEntryModeJSON `json:"-"`
}

// cardAuthorizationPosEntryModeJSON contains the JSON metadata for the struct
// [CardAuthorizationPosEntryMode]
type cardAuthorizationPosEntryModeJSON struct {
	Card        apijson.Field
	Cardholder  apijson.Field
	Pan         apijson.Field
	PinEntered  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationPosEntryMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationPosEntryModeJSON) RawJSON() string {
	return r.raw
}

// Card Presence Indicator
type CardAuthorizationPosEntryModeCard string

const (
	CardAuthorizationPosEntryModeCardPresent    CardAuthorizationPosEntryModeCard = "PRESENT"
	CardAuthorizationPosEntryModeCardNotPresent CardAuthorizationPosEntryModeCard = "NOT_PRESENT"
	CardAuthorizationPosEntryModeCardUnknown    CardAuthorizationPosEntryModeCard = "UNKNOWN"
)

func (r CardAuthorizationPosEntryModeCard) IsKnown() bool {
	switch r {
	case CardAuthorizationPosEntryModeCardPresent, CardAuthorizationPosEntryModeCardNotPresent, CardAuthorizationPosEntryModeCardUnknown:
		return true
	}
	return false
}

// Cardholder Presence Indicator
type CardAuthorizationPosEntryModeCardholder string

const (
	CardAuthorizationPosEntryModeCardholderDeferredBilling CardAuthorizationPosEntryModeCardholder = "DEFERRED_BILLING"
	CardAuthorizationPosEntryModeCardholderElectronicOrder CardAuthorizationPosEntryModeCardholder = "ELECTRONIC_ORDER"
	CardAuthorizationPosEntryModeCardholderInstallment     CardAuthorizationPosEntryModeCardholder = "INSTALLMENT"
	CardAuthorizationPosEntryModeCardholderMailOrder       CardAuthorizationPosEntryModeCardholder = "MAIL_ORDER"
	CardAuthorizationPosEntryModeCardholderNotPresent      CardAuthorizationPosEntryModeCardholder = "NOT_PRESENT"
	CardAuthorizationPosEntryModeCardholderPresent         CardAuthorizationPosEntryModeCardholder = "PRESENT"
	CardAuthorizationPosEntryModeCardholderReoccurring     CardAuthorizationPosEntryModeCardholder = "REOCCURRING"
	CardAuthorizationPosEntryModeCardholderTelephoneOrder  CardAuthorizationPosEntryModeCardholder = "TELEPHONE_ORDER"
	CardAuthorizationPosEntryModeCardholderUnknown         CardAuthorizationPosEntryModeCardholder = "UNKNOWN"
)

func (r CardAuthorizationPosEntryModeCardholder) IsKnown() bool {
	switch r {
	case CardAuthorizationPosEntryModeCardholderDeferredBilling, CardAuthorizationPosEntryModeCardholderElectronicOrder, CardAuthorizationPosEntryModeCardholderInstallment, CardAuthorizationPosEntryModeCardholderMailOrder, CardAuthorizationPosEntryModeCardholderNotPresent, CardAuthorizationPosEntryModeCardholderPresent, CardAuthorizationPosEntryModeCardholderReoccurring, CardAuthorizationPosEntryModeCardholderTelephoneOrder, CardAuthorizationPosEntryModeCardholderUnknown:
		return true
	}
	return false
}

// Method of entry for the PAN
type CardAuthorizationPosEntryModePan string

const (
	CardAuthorizationPosEntryModePanAutoEntry           CardAuthorizationPosEntryModePan = "AUTO_ENTRY"
	CardAuthorizationPosEntryModePanBarCode             CardAuthorizationPosEntryModePan = "BAR_CODE"
	CardAuthorizationPosEntryModePanContactless         CardAuthorizationPosEntryModePan = "CONTACTLESS"
	CardAuthorizationPosEntryModePanEcommerce           CardAuthorizationPosEntryModePan = "ECOMMERCE"
	CardAuthorizationPosEntryModePanErrorKeyed          CardAuthorizationPosEntryModePan = "ERROR_KEYED"
	CardAuthorizationPosEntryModePanErrorMagneticStripe CardAuthorizationPosEntryModePan = "ERROR_MAGNETIC_STRIPE"
	CardAuthorizationPosEntryModePanIcc                 CardAuthorizationPosEntryModePan = "ICC"
	CardAuthorizationPosEntryModePanKeyEntered          CardAuthorizationPosEntryModePan = "KEY_ENTERED"
	CardAuthorizationPosEntryModePanMagneticStripe      CardAuthorizationPosEntryModePan = "MAGNETIC_STRIPE"
	CardAuthorizationPosEntryModePanManual              CardAuthorizationPosEntryModePan = "MANUAL"
	CardAuthorizationPosEntryModePanOcr                 CardAuthorizationPosEntryModePan = "OCR"
	CardAuthorizationPosEntryModePanSecureCardless      CardAuthorizationPosEntryModePan = "SECURE_CARDLESS"
	CardAuthorizationPosEntryModePanUnspecified         CardAuthorizationPosEntryModePan = "UNSPECIFIED"
	CardAuthorizationPosEntryModePanUnknown             CardAuthorizationPosEntryModePan = "UNKNOWN"
	CardAuthorizationPosEntryModePanCredentialOnFile    CardAuthorizationPosEntryModePan = "CREDENTIAL_ON_FILE"
)

func (r CardAuthorizationPosEntryModePan) IsKnown() bool {
	switch r {
	case CardAuthorizationPosEntryModePanAutoEntry, CardAuthorizationPosEntryModePanBarCode, CardAuthorizationPosEntryModePanContactless, CardAuthorizationPosEntryModePanEcommerce, CardAuthorizationPosEntryModePanErrorKeyed, CardAuthorizationPosEntryModePanErrorMagneticStripe, CardAuthorizationPosEntryModePanIcc, CardAuthorizationPosEntryModePanKeyEntered, CardAuthorizationPosEntryModePanMagneticStripe, CardAuthorizationPosEntryModePanManual, CardAuthorizationPosEntryModePanOcr, CardAuthorizationPosEntryModePanSecureCardless, CardAuthorizationPosEntryModePanUnspecified, CardAuthorizationPosEntryModePanUnknown, CardAuthorizationPosEntryModePanCredentialOnFile:
		return true
	}
	return false
}

type CardAuthorizationPosTerminal struct {
	// True if a clerk is present at the sale.
	Attended bool `json:"attended" api:"required"`
	// True if the terminal is capable of retaining the card.
	CardRetentionCapable bool `json:"card_retention_capable" api:"required"`
	// True if the sale was made at the place of business (vs. mobile).
	OnPremise bool `json:"on_premise" api:"required"`
	// The person that is designated to swipe the card
	Operator CardAuthorizationPosTerminalOperator `json:"operator" api:"required"`
	// True if the terminal is capable of partial approval. Partial approval is when
	// part of a transaction is approved and another payment must be used for the
	// remainder. Example scenario: A $40 transaction is attempted on a prepaid card
	// with a $25 balance. If partial approval is enabled, $25 can be authorized, at
	// which point the POS will prompt the user for an additional payment of $15.
	PartialApprovalCapable bool `json:"partial_approval_capable" api:"required"`
	// Status of whether the POS is able to accept PINs
	PinCapability CardAuthorizationPosTerminalPinCapability `json:"pin_capability" api:"required"`
	// POS Type
	Type CardAuthorizationPosTerminalType `json:"type" api:"required"`
	// Uniquely identifies a terminal at the card acceptor location of acquiring
	// institutions or merchant POS Systems. Left justified with trailing spaces.
	AcceptorTerminalID string                           `json:"acceptor_terminal_id" api:"nullable"`
	JSON               cardAuthorizationPosTerminalJSON `json:"-"`
}

// cardAuthorizationPosTerminalJSON contains the JSON metadata for the struct
// [CardAuthorizationPosTerminal]
type cardAuthorizationPosTerminalJSON struct {
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

func (r *CardAuthorizationPosTerminal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationPosTerminalJSON) RawJSON() string {
	return r.raw
}

// The person that is designated to swipe the card
type CardAuthorizationPosTerminalOperator string

const (
	CardAuthorizationPosTerminalOperatorAdministrative CardAuthorizationPosTerminalOperator = "ADMINISTRATIVE"
	CardAuthorizationPosTerminalOperatorCardholder     CardAuthorizationPosTerminalOperator = "CARDHOLDER"
	CardAuthorizationPosTerminalOperatorCardAcceptor   CardAuthorizationPosTerminalOperator = "CARD_ACCEPTOR"
	CardAuthorizationPosTerminalOperatorUnknown        CardAuthorizationPosTerminalOperator = "UNKNOWN"
)

func (r CardAuthorizationPosTerminalOperator) IsKnown() bool {
	switch r {
	case CardAuthorizationPosTerminalOperatorAdministrative, CardAuthorizationPosTerminalOperatorCardholder, CardAuthorizationPosTerminalOperatorCardAcceptor, CardAuthorizationPosTerminalOperatorUnknown:
		return true
	}
	return false
}

// Status of whether the POS is able to accept PINs
type CardAuthorizationPosTerminalPinCapability string

const (
	CardAuthorizationPosTerminalPinCapabilityCapable     CardAuthorizationPosTerminalPinCapability = "CAPABLE"
	CardAuthorizationPosTerminalPinCapabilityInoperative CardAuthorizationPosTerminalPinCapability = "INOPERATIVE"
	CardAuthorizationPosTerminalPinCapabilityNotCapable  CardAuthorizationPosTerminalPinCapability = "NOT_CAPABLE"
	CardAuthorizationPosTerminalPinCapabilityUnspecified CardAuthorizationPosTerminalPinCapability = "UNSPECIFIED"
)

func (r CardAuthorizationPosTerminalPinCapability) IsKnown() bool {
	switch r {
	case CardAuthorizationPosTerminalPinCapabilityCapable, CardAuthorizationPosTerminalPinCapabilityInoperative, CardAuthorizationPosTerminalPinCapabilityNotCapable, CardAuthorizationPosTerminalPinCapabilityUnspecified:
		return true
	}
	return false
}

// POS Type
type CardAuthorizationPosTerminalType string

const (
	CardAuthorizationPosTerminalTypeAdministrative        CardAuthorizationPosTerminalType = "ADMINISTRATIVE"
	CardAuthorizationPosTerminalTypeAtm                   CardAuthorizationPosTerminalType = "ATM"
	CardAuthorizationPosTerminalTypeAuthorization         CardAuthorizationPosTerminalType = "AUTHORIZATION"
	CardAuthorizationPosTerminalTypeCouponMachine         CardAuthorizationPosTerminalType = "COUPON_MACHINE"
	CardAuthorizationPosTerminalTypeDialTerminal          CardAuthorizationPosTerminalType = "DIAL_TERMINAL"
	CardAuthorizationPosTerminalTypeEcommerce             CardAuthorizationPosTerminalType = "ECOMMERCE"
	CardAuthorizationPosTerminalTypeEcr                   CardAuthorizationPosTerminalType = "ECR"
	CardAuthorizationPosTerminalTypeFuelMachine           CardAuthorizationPosTerminalType = "FUEL_MACHINE"
	CardAuthorizationPosTerminalTypeHomeTerminal          CardAuthorizationPosTerminalType = "HOME_TERMINAL"
	CardAuthorizationPosTerminalTypeMicr                  CardAuthorizationPosTerminalType = "MICR"
	CardAuthorizationPosTerminalTypeOffPremise            CardAuthorizationPosTerminalType = "OFF_PREMISE"
	CardAuthorizationPosTerminalTypePayment               CardAuthorizationPosTerminalType = "PAYMENT"
	CardAuthorizationPosTerminalTypePda                   CardAuthorizationPosTerminalType = "PDA"
	CardAuthorizationPosTerminalTypePhone                 CardAuthorizationPosTerminalType = "PHONE"
	CardAuthorizationPosTerminalTypePoint                 CardAuthorizationPosTerminalType = "POINT"
	CardAuthorizationPosTerminalTypePosTerminal           CardAuthorizationPosTerminalType = "POS_TERMINAL"
	CardAuthorizationPosTerminalTypePublicUtility         CardAuthorizationPosTerminalType = "PUBLIC_UTILITY"
	CardAuthorizationPosTerminalTypeSelfService           CardAuthorizationPosTerminalType = "SELF_SERVICE"
	CardAuthorizationPosTerminalTypeTelevision            CardAuthorizationPosTerminalType = "TELEVISION"
	CardAuthorizationPosTerminalTypeTeller                CardAuthorizationPosTerminalType = "TELLER"
	CardAuthorizationPosTerminalTypeTravelersCheckMachine CardAuthorizationPosTerminalType = "TRAVELERS_CHECK_MACHINE"
	CardAuthorizationPosTerminalTypeVending               CardAuthorizationPosTerminalType = "VENDING"
	CardAuthorizationPosTerminalTypeVoice                 CardAuthorizationPosTerminalType = "VOICE"
	CardAuthorizationPosTerminalTypeUnknown               CardAuthorizationPosTerminalType = "UNKNOWN"
)

func (r CardAuthorizationPosTerminalType) IsKnown() bool {
	switch r {
	case CardAuthorizationPosTerminalTypeAdministrative, CardAuthorizationPosTerminalTypeAtm, CardAuthorizationPosTerminalTypeAuthorization, CardAuthorizationPosTerminalTypeCouponMachine, CardAuthorizationPosTerminalTypeDialTerminal, CardAuthorizationPosTerminalTypeEcommerce, CardAuthorizationPosTerminalTypeEcr, CardAuthorizationPosTerminalTypeFuelMachine, CardAuthorizationPosTerminalTypeHomeTerminal, CardAuthorizationPosTerminalTypeMicr, CardAuthorizationPosTerminalTypeOffPremise, CardAuthorizationPosTerminalTypePayment, CardAuthorizationPosTerminalTypePda, CardAuthorizationPosTerminalTypePhone, CardAuthorizationPosTerminalTypePoint, CardAuthorizationPosTerminalTypePosTerminal, CardAuthorizationPosTerminalTypePublicUtility, CardAuthorizationPosTerminalTypeSelfService, CardAuthorizationPosTerminalTypeTelevision, CardAuthorizationPosTerminalTypeTeller, CardAuthorizationPosTerminalTypeTravelersCheckMachine, CardAuthorizationPosTerminalTypeVending, CardAuthorizationPosTerminalTypeVoice, CardAuthorizationPosTerminalTypeUnknown:
		return true
	}
	return false
}

type CardAuthorizationChallengeResponseParams struct {
	// Whether the cardholder has approved or declined the issued challenge
	Response param.Field[CardAuthorizationChallengeResponseParamsResponse] `json:"response" api:"required"`
}

func (r CardAuthorizationChallengeResponseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the cardholder has approved or declined the issued challenge
type CardAuthorizationChallengeResponseParamsResponse string

const (
	CardAuthorizationChallengeResponseParamsResponseApprove CardAuthorizationChallengeResponseParamsResponse = "APPROVE"
	CardAuthorizationChallengeResponseParamsResponseDecline CardAuthorizationChallengeResponseParamsResponse = "DECLINE"
)

func (r CardAuthorizationChallengeResponseParamsResponse) IsKnown() bool {
	switch r {
	case CardAuthorizationChallengeResponseParamsResponseApprove, CardAuthorizationChallengeResponseParamsResponseDecline:
		return true
	}
	return false
}
