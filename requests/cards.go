package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type SpendLimitDuration string

const (
	SpendLimitDurationAnnually    SpendLimitDuration = "ANNUALLY"
	SpendLimitDurationForever     SpendLimitDuration = "FOREVER"
	SpendLimitDurationMonthly     SpendLimitDuration = "MONTHLY"
	SpendLimitDurationTransaction SpendLimitDuration = "TRANSACTION"
)

type EmbedRequestParams struct {
	// A publicly available URI, so the white-labeled card element can be styled with
	// the client's branding.
	Css field.Field[string] `json:"css"`
	// An RFC 3339 timestamp for when the request should expire. UTC time zone.
	//
	// If no timezone is specified, UTC will be used. If payload does not contain an
	// expiration, the request will never expire.
	//
	// Using an `expiration` reduces the risk of a
	// [replay attack](https://en.wikipedia.org/wiki/Replay_attack). Without supplying
	// the `expiration`, in the event that a malicious user gets a copy of your request
	// in transit, they will be able to obtain the response data indefinitely.
	Expiration field.Field[time.Time] `json:"expiration" format:"date-time"`
	// Globally unique identifier for the card to be displayed.
	Token field.Field[string] `json:"token,required" format:"uuid"`
	// Required if you want to post the element clicked to the parent iframe.
	//
	// If you supply this param, you can also capture click events in the parent iframe
	// by adding an event listener.
	TargetOrigin field.Field[string] `json:"target_origin"`
}

// MarshalJSON serializes EmbedRequestParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EmbedRequestParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CardNewParams struct {
	// Globally unique identifier for the account that the card will be associated
	// with. Required for programs enrolling users using the
	// [/account_holders endpoint](https://docs.lithic.com/docs/account-holders-kyc).
	// See [Managing Your Program](doc:managing-your-program) for more information.
	AccountToken field.Field[string] `json:"account_token" format:"uuid"`
	// For physical card programs with more than one BIN range. This must be configured
	// with Lithic before use. Identifies the card program/BIN range under which to
	// create the card. If omitted, will utilize the program's default
	// `card_program_token`. In Sandbox, use 00000000-0000-0000-1000-000000000000 and
	// 00000000-0000-0000-2000-000000000000 to test creating cards on specific card
	// programs.
	CardProgramToken field.Field[string] `json:"card_program_token" format:"uuid"`
	// Two digit (MM) expiry month. If neither `exp_month` nor `exp_year` is provided,
	// an expiration date will be generated.
	ExpMonth field.Field[string] `json:"exp_month"`
	// Four digit (yyyy) expiry year. If neither `exp_month` nor `exp_year` is
	// provided, an expiration date will be generated.
	ExpYear field.Field[string] `json:"exp_year"`
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo field.Field[string] `json:"memo"`
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the card
	// limit.
	SpendLimit field.Field[int64] `json:"spend_limit"`
	// Spend limit duration values:
	//
	//   - `ANNUALLY` - Card will authorize transactions up to spend limit in a calendar
	//     year.
	//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
	//     of the card.
	//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
	//     trailing month. Month is calculated as this calendar date one month prior.
	//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
	//     transaction is under the spend limit.
	SpendLimitDuration field.Field[SpendLimitDuration] `json:"spend_limit_duration"`
	// Card state values:
	//
	//   - `OPEN` - Card will approve authorizations (if they match card and account
	//     parameters).
	//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
	//     time.
	State field.Field[CardNewParamsState] `json:"state"`
	// Card types:
	//
	//   - `VIRTUAL` - Card will authorize at any merchant and can be added to a digital
	//     wallet like Apple Pay or Google Pay (if the card program is digital
	//     wallet-enabled).
	//   - `PHYSICAL` - Manufactured and sent to the cardholder. We offer white label
	//     branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe functionality.
	//     Reach out at [lithic.com/contact](https://lithic.com/contact) for more
	//     information.
	//   - `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first merchant that
	//     successfully authorizes the card.
	//   - `SINGLE_USE` - _[Deprecated]_ Card is closed upon first successful
	//     authorization.
	Type field.Field[CardNewParamsType] `json:"type,required"`
	// Encrypted PIN block (in base64). Only applies to cards of type `PHYSICAL` and
	// `VIRTUAL`. See
	// [Encrypted PIN Block](https://docs.lithic.com/docs/cards#encrypted-pin-block-enterprise).
	Pin field.Field[string] `json:"pin"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken field.Field[string] `json:"digital_card_art_token" format:"uuid"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID       field.Field[string]          `json:"product_id"`
	ShippingAddress field.Field[ShippingAddress] `json:"shipping_address"`
	// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
	// options besides `STANDARD` require additional permissions.
	//
	//   - `STANDARD` - USPS regular mail or similar international option, with no
	//     tracking
	//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
	//     with tracking
	//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
	//     tracking
	ShippingMethod field.Field[CardNewParamsShippingMethod] `json:"shipping_method"`
}

// MarshalJSON serializes CardNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CardNewParamsState string

const (
	CardNewParamsStateOpen   CardNewParamsState = "OPEN"
	CardNewParamsStatePaused CardNewParamsState = "PAUSED"
)

type CardNewParamsType string

const (
	CardNewParamsTypeVirtual        CardNewParamsType = "VIRTUAL"
	CardNewParamsTypePhysical       CardNewParamsType = "PHYSICAL"
	CardNewParamsTypeMerchantLocked CardNewParamsType = "MERCHANT_LOCKED"
	CardNewParamsTypeSingleUse      CardNewParamsType = "SINGLE_USE"
)

type CardNewParamsShippingMethod string

const (
	CardNewParamsShippingMethodStandard             CardNewParamsShippingMethod = "STANDARD"
	CardNewParamsShippingMethodStandardWithTracking CardNewParamsShippingMethod = "STANDARD_WITH_TRACKING"
	CardNewParamsShippingMethodExpedited            CardNewParamsShippingMethod = "EXPEDITED"
)

type CardUpdateParams struct {
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo field.Field[string] `json:"memo"`
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the card
	// limit.
	SpendLimit field.Field[int64] `json:"spend_limit"`
	// Spend limit duration values:
	//
	//   - `ANNUALLY` - Card will authorize transactions up to spend limit in a calendar
	//     year.
	//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
	//     of the card.
	//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
	//     trailing month. Month is calculated as this calendar date one month prior.
	//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
	//     transaction is under the spend limit.
	SpendLimitDuration field.Field[SpendLimitDuration] `json:"spend_limit_duration"`
	// Identifier for any Auth Rules that will be applied to transactions taking place
	// with the card.
	AuthRuleToken field.Field[string] `json:"auth_rule_token"`
	// Card state values:
	//
	//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
	//     be undone.
	//   - `OPEN` - Card will approve authorizations (if they match card and account
	//     parameters).
	//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
	//     time.
	State field.Field[CardUpdateParamsState] `json:"state"`
	// Encrypted PIN block (in base64). Only applies to cards of type `PHYSICAL` and
	// `VIRTUAL`. See
	// [Encrypted PIN Block](https://docs.lithic.com/docs/cards#encrypted-pin-block-enterprise).
	Pin field.Field[string] `json:"pin"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken field.Field[string] `json:"digital_card_art_token" format:"uuid"`
}

// MarshalJSON serializes CardUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CardUpdateParamsState string

const (
	CardUpdateParamsStateClosed CardUpdateParamsState = "CLOSED"
	CardUpdateParamsStateOpen   CardUpdateParamsState = "OPEN"
	CardUpdateParamsStatePaused CardUpdateParamsState = "PAUSED"
)

type CardListParams struct {
	// Returns cards associated with the specified account.
	AccountToken field.Field[string] `query:"account_token" format:"uuid"`
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
	// Page (for pagination).
	Page field.Field[int64] `query:"page"`
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
}

// URLQuery serializes CardListParams into a url.Values of the query parameters
// associated with this value
func (r *CardListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type CardEmbedParams struct {
	// A base64 encoded JSON string of an EmbedRequest to specify which card to load.
	EmbedRequest field.Field[string] `query:"embed_request"`
	// SHA256 HMAC of the embed_request JSON string with base64 digest.
	Hmac field.Field[string] `query:"hmac"`
}

// URLQuery serializes CardEmbedParams into a url.Values of the query parameters
// associated with this value
func (r *CardEmbedParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type CardProvisionParams struct {
	// Name of digital wallet provider.
	DigitalWallet field.Field[CardProvisionParamsDigitalWallet] `json:"digital_wallet"`
	// Required for `APPLE_PAY`. Base64 cryptographic nonce provided by the device's
	// wallet.
	Nonce field.Field[string] `json:"nonce" format:"byte"`
	// Required for `APPLE_PAY`. Base64 cryptographic nonce provided by the device's
	// wallet.
	NonceSignature field.Field[string] `json:"nonce_signature" format:"byte"`
	// Required for `APPLE_PAY`. Apple's public leaf certificate. Base64 encoded in PEM
	// format with headers `(-----BEGIN CERTIFICATE-----)` and trailers omitted.
	// Provided by the device's wallet.
	Certificate field.Field[string] `json:"certificate" format:"byte"`
}

// MarshalJSON serializes CardProvisionParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardProvisionParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CardProvisionParamsDigitalWallet string

const (
	CardProvisionParamsDigitalWalletApplePay   CardProvisionParamsDigitalWallet = "APPLE_PAY"
	CardProvisionParamsDigitalWalletGooglePay  CardProvisionParamsDigitalWallet = "GOOGLE_PAY"
	CardProvisionParamsDigitalWalletSamsungPay CardProvisionParamsDigitalWallet = "SAMSUNG_PAY"
)

type CardReissueParams struct {
	// If omitted, the previous shipping address will be used.
	ShippingAddress field.Field[ShippingAddress] `json:"shipping_address"`
	// Shipping method for the card. Use of options besides `STANDARD` require
	// additional permissions.
	//
	//   - `STANDARD` - USPS regular mail or similar international option, with no
	//     tracking
	//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
	//     with tracking
	//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
	//     tracking
	ShippingMethod field.Field[CardReissueParamsShippingMethod] `json:"shipping_method"`
	// Specifies the configuration (e.g. physical card art) that the card should be
	// manufactured with, and only applies to cards of type `PHYSICAL`. This must be
	// configured with Lithic before use.
	ProductID field.Field[string] `json:"product_id"`
}

// MarshalJSON serializes CardReissueParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardReissueParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CardReissueParamsShippingMethod string

const (
	CardReissueParamsShippingMethodStandard             CardReissueParamsShippingMethod = "STANDARD"
	CardReissueParamsShippingMethodStandardWithTracking CardReissueParamsShippingMethod = "STANDARD_WITH_TRACKING"
	CardReissueParamsShippingMethodExpedited            CardReissueParamsShippingMethod = "EXPEDITED"
)
