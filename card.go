// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/lithic-com/lithic-go/shared"
)

// CardService contains methods and other services that help with interacting with
// the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardService] method instead.
type CardService struct {
	Options               []option.RequestOption
	AggregateBalances     *CardAggregateBalanceService
	Balances              *CardBalanceService
	FinancialTransactions *CardFinancialTransactionService
}

// NewCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardService(opts ...option.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	r.AggregateBalances = NewCardAggregateBalanceService(opts...)
	r.Balances = NewCardBalanceService(opts...)
	r.FinancialTransactions = NewCardFinancialTransactionService(opts...)
	return
}

// Create a new virtual or physical card. Parameters `shipping_address` and
// `product_id` only apply to physical cards.
func (r *CardService) New(ctx context.Context, body CardNewParams, opts ...option.RequestOption) (res *CardNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get card configuration such as spend limit and state.
func (r *CardService) Get(ctx context.Context, cardToken string, opts ...option.RequestOption) (res *CardGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the specified properties of the card. Unsupplied properties will remain
// unchanged.
//
// _Note: setting a card to a `CLOSED` state is a final action that cannot be
// undone._
func (r *CardService) Update(ctx context.Context, cardToken string, body CardUpdateParams, opts ...option.RequestOption) (res *CardUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List cards.
func (r *CardService) List(ctx context.Context, query CardListParams, opts ...option.RequestOption) (res *pagination.CursorPage[CardListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/cards"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List cards.
func (r *CardService) ListAutoPaging(ctx context.Context, query CardListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CardListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Convert a virtual card into a physical card and manufacture it. Customer must
// supply relevant fields for physical card creation including `product_id`,
// `carrier`, `shipping_method`, and `shipping_address`. The card token will be
// unchanged. The card's type will be altered to `PHYSICAL`. The card will be set
// to state `PENDING_FULFILLMENT` and fulfilled at next fulfillment cycle. Virtual
// cards created on card programs which do not support physical cards cannot be
// converted. The card program cannot be changed as part of the conversion. Cards
// must be in an `OPEN` state to be converted. Only applies to cards of type
// `VIRTUAL` (or existing cards with deprecated types of `DIGITAL_WALLET` and
// `UNLOCKED`).
func (r *CardService) ConvertPhysical(ctx context.Context, cardToken string, body CardConvertPhysicalParams, opts ...option.RequestOption) (res *CardConvertPhysicalResponse, err error) {
	opts = append(r.Options[:], opts...)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s/convert_physical", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Handling full card PANs and CVV codes requires that you comply with the Payment
// Card Industry Data Security Standards (PCI DSS). Some clients choose to reduce
// their compliance obligations by leveraging our embedded card UI solution
// documented below.
//
// In this setup, PANs and CVV codes are presented to the end-user via a card UI
// that we provide, optionally styled in the customer's branding using a specified
// css stylesheet. A user's browser makes the request directly to api.lithic.com,
// so card PANs and CVVs never touch the API customer's servers while full card
// data is displayed to their end-users. The response contains an HTML document
// (see Embedded Card UI or Changelog for upcoming changes in January). This means
// that the url for the request can be inserted straight into the `src` attribute
// of an iframe.
//
// ```html
// <iframe
//
//	id="card-iframe"
//	src="https://sandbox.lithic.com/v1/embed/card?embed_request=eyJjc3MiO...;hmac=r8tx1..."
//	allow="clipboard-write"
//	class="content"
//
// ></iframe>
// ```
//
// You should compute the request payload on the server side. You can render it (or
// the whole iframe) on the server or make an ajax call from your front end code,
// but **do not ever embed your API key into front end code, as doing so introduces
// a serious security vulnerability**.
func (r *CardService) Embed(ctx context.Context, query CardEmbedParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	path := "v1/embed/card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *CardService) GetEmbedHTML(ctx context.Context, params CardGetEmbedHTMLParams, opts ...option.RequestOption) (res []byte, err error) {
	opts = append(r.Options, opts...)
	buf, err := params.MarshalJSON()
	if err != nil {
		return nil, err
	}
	cfg, err := requestconfig.NewRequestConfig(ctx, "GET", "v1/embed/card", nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, []byte(cfg.APIKey))
	mac.Write(buf)
	sign := mac.Sum(nil)
	err = cfg.Apply(
		option.WithHeader("Accept", "text/html"),
		option.WithQuery("hmac", base64.StdEncoding.EncodeToString(sign)),
		option.WithQuery("embed_request", base64.StdEncoding.EncodeToString(buf)),
	)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	return
}

// Handling full card PANs and CVV codes requires that you comply with the Payment
// Card Industry Data Security Standards (PCI DSS). Some clients choose to reduce
// their compliance obligations by leveraging our embedded card UI solution
// documented below.
//
// In this setup, PANs and CVV codes are presented to the end-user via a card UI
// that we provide, optionally styled in the customer's branding using a specified
// css stylesheet. A user's browser makes the request directly to api.lithic.com,
// so card PANs and CVVs never touch the API customer's servers while full card
// data is displayed to their end-users. The response contains an HTML document.
// This means that the url for the request can be inserted straight into the `src`
// attribute of an iframe.
//
// ```html
// <iframe
//
//	id="card-iframe"
//	src="https://sandbox.lithic.com/v1/embed/card?embed_request=eyJjc3MiO...;hmac=r8tx1..."
//	allow="clipboard-write"
//	class="content"
//
// ></iframe>
// ```
//
// You should compute the request payload on the server side. You can render it (or
// the whole iframe) on the server or make an ajax call from your front end code,
// but **do not ever embed your API key into front end code, as doing so introduces
// a serious security vulnerability**.
func (r *CardService) GetEmbedURL(ctx context.Context, params CardGetEmbedURLParams, opts ...option.RequestOption) (res *url.URL, err error) {
	buf, err := params.MarshalJSON()
	if err != nil {
		return nil, err
	}
	cfg, err := requestconfig.NewRequestConfig(ctx, "GET", "v1/embed/card", nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, []byte(cfg.APIKey))
	mac.Write(buf)
	sign := mac.Sum(nil)
	err = cfg.Apply(
		option.WithQuery("hmac", base64.StdEncoding.EncodeToString(sign)),
		option.WithQuery("embed_request", base64.StdEncoding.EncodeToString(buf)),
	)
	if err != nil {
		return nil, err
	}
	return cfg.Request.URL, nil
}

// Allow your cardholders to directly add payment cards to the device's digital
// wallet (e.g. Apple Pay) with one touch from your app.
//
// This requires some additional setup and configuration. Please
// [Contact Us](https://lithic.com/contact) or your Customer Success representative
// for more information.
func (r *CardService) Provision(ctx context.Context, cardToken string, body CardProvisionParams, opts ...option.RequestOption) (res *CardProvisionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s/provision", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Initiate print and shipment of a duplicate physical card (e.g. card is
// physically damaged). The PAN, expiry, and CVC2 will remain the same and the
// original card can continue to be used until the new card is activated. Only
// applies to cards of type `PHYSICAL`. A card can be replaced or renewed a total
// of 8 times.
func (r *CardService) Reissue(ctx context.Context, cardToken string, body CardReissueParams, opts ...option.RequestOption) (res *CardReissueResponse, err error) {
	opts = append(r.Options[:], opts...)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s/reissue", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Applies to card types `PHYSICAL` and `VIRTUAL`. For `PHYSICAL`, creates a new
// card with the same card token and PAN, but updated expiry and CVC2 code. The
// original card will keep working for card-present transactions until the new card
// is activated. For card-not-present transactions, the original card details
// (expiry, CVC2) will also keep working until the new card is activated. A
// `PHYSICAL` card can be replaced or renewed a total of 8 times. For `VIRTUAL`,
// the card will retain the same card token and PAN and receive an updated expiry
// and CVC2 code. `product_id`, `shipping_method`, `shipping_address`, `carrier`
// are only relevant for renewing `PHYSICAL` cards.
func (r *CardService) Renew(ctx context.Context, cardToken string, body CardRenewParams, opts ...option.RequestOption) (res *CardRenewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s/renew", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Card's available spend limit, which is based on the spend limit configured
// on the Card and the amount already spent over the spend limit's duration. For
// example, if the Card has a monthly spend limit of $1000 configured, and has
// spent $600 in the last month, the available spend limit returned would be $400.
func (r *CardService) GetSpendLimits(ctx context.Context, cardToken string, opts ...option.RequestOption) (res *CardSpendLimits, err error) {
	opts = append(r.Options[:], opts...)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s/spend_limits", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get card configuration such as spend limit and state. Customers must be PCI
// compliant to use this endpoint. Please contact
// [support@lithic.com](mailto:support@lithic.com) for questions. _Note: this is a
// `POST` endpoint because it is more secure to send sensitive data in a request
// body than in a URL._
func (r *CardService) SearchByPan(ctx context.Context, body CardSearchByPanParams, opts ...option.RequestOption) (res *CardSearchByPanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/cards/search_by_pan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CardSpendLimits struct {
	AvailableSpendLimit CardSpendLimitsAvailableSpendLimit `json:"available_spend_limit,required"`
	SpendLimit          CardSpendLimitsSpendLimit          `json:"spend_limit"`
	SpendVelocity       CardSpendLimitsSpendVelocity       `json:"spend_velocity"`
	JSON                cardSpendLimitsJSON                `json:"-"`
}

// cardSpendLimitsJSON contains the JSON metadata for the struct [CardSpendLimits]
type cardSpendLimitsJSON struct {
	AvailableSpendLimit apijson.Field
	SpendLimit          apijson.Field
	SpendVelocity       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardSpendLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardSpendLimitsJSON) RawJSON() string {
	return r.raw
}

type CardSpendLimitsAvailableSpendLimit struct {
	// The available spend limit (in cents) relative to the annual limit configured on
	// the Card (e.g. 100000 would be a $1,000 limit).
	Annually int64 `json:"annually"`
	// The available spend limit (in cents) relative to the forever limit configured on
	// the Card.
	Forever int64 `json:"forever"`
	// The available spend limit (in cents) relative to the monthly limit configured on
	// the Card.
	Monthly int64                                  `json:"monthly"`
	JSON    cardSpendLimitsAvailableSpendLimitJSON `json:"-"`
}

// cardSpendLimitsAvailableSpendLimitJSON contains the JSON metadata for the struct
// [CardSpendLimitsAvailableSpendLimit]
type cardSpendLimitsAvailableSpendLimitJSON struct {
	Annually    apijson.Field
	Forever     apijson.Field
	Monthly     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardSpendLimitsAvailableSpendLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardSpendLimitsAvailableSpendLimitJSON) RawJSON() string {
	return r.raw
}

type CardSpendLimitsSpendLimit struct {
	// The configured annual spend limit (in cents) on the Card.
	Annually int64 `json:"annually"`
	// The configured forever spend limit (in cents) on the Card.
	Forever int64 `json:"forever"`
	// The configured monthly spend limit (in cents) on the Card.
	Monthly int64                         `json:"monthly"`
	JSON    cardSpendLimitsSpendLimitJSON `json:"-"`
}

// cardSpendLimitsSpendLimitJSON contains the JSON metadata for the struct
// [CardSpendLimitsSpendLimit]
type cardSpendLimitsSpendLimitJSON struct {
	Annually    apijson.Field
	Forever     apijson.Field
	Monthly     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardSpendLimitsSpendLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardSpendLimitsSpendLimitJSON) RawJSON() string {
	return r.raw
}

type CardSpendLimitsSpendVelocity struct {
	// Current annual spend velocity (in cents) on the Card. Present if annual spend
	// limit is set.
	Annually int64 `json:"annually"`
	// Current forever spend velocity (in cents) on the Card. Present if forever spend
	// limit is set.
	Forever int64 `json:"forever"`
	// Current monthly spend velocity (in cents) on the Card. Present if monthly spend
	// limit is set.
	Monthly int64                            `json:"monthly"`
	JSON    cardSpendLimitsSpendVelocityJSON `json:"-"`
}

// cardSpendLimitsSpendVelocityJSON contains the JSON metadata for the struct
// [CardSpendLimitsSpendVelocity]
type cardSpendLimitsSpendVelocityJSON struct {
	Annually    apijson.Field
	Forever     apijson.Field
	Monthly     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardSpendLimitsSpendVelocity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardSpendLimitsSpendVelocityJSON) RawJSON() string {
	return r.raw
}

// Spend limit duration values:
//
//   - `ANNUALLY` - Card will authorize transactions up to spend limit for the
//     trailing year.
//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
//     of the card.
//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
//     trailing month. To support recurring monthly payments, which can occur on
//     different day every month, the time window we consider for monthly velocity
//     starts 6 days after the current calendar date one month prior.
//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
//     transaction is under the spend limit.
type SpendLimitDuration string

const (
	SpendLimitDurationAnnually    SpendLimitDuration = "ANNUALLY"
	SpendLimitDurationForever     SpendLimitDuration = "FOREVER"
	SpendLimitDurationMonthly     SpendLimitDuration = "MONTHLY"
	SpendLimitDurationTransaction SpendLimitDuration = "TRANSACTION"
)

func (r SpendLimitDuration) IsKnown() bool {
	switch r {
	case SpendLimitDurationAnnually, SpendLimitDurationForever, SpendLimitDurationMonthly, SpendLimitDurationTransaction:
		return true
	}
	return false
}

// Card details with potentially PCI sensitive information for Enterprise customers
type CardNewResponse struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding CardNewResponseFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus CardNewResponsePinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration
	SpendLimitDuration CardNewResponseSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
	// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
	// they match card and account parameters). _ `PAUSED` - Card will decline
	// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
	// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
	// manufacturing and fulfillment. Cards in this state can accept authorizations for
	// e-commerce purchases, but not for "Card Present" purchases where the physical
	// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
	// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
	// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
	// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	// transactions or can be added to mobile wallets. API clients should update the
	// card's state to `OPEN` only after the cardholder confirms receipt of the card.
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardNewResponseState `json:"state,required"`
	// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
	// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
	// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
	// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
	// functionality. _ `SINGLE_USE` - Card is closed upon first successful
	// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
	// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
	// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
	// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	// VIRTUAL instead.
	Type CardNewResponseType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user's digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use.
	DigitalCardArtToken string `json:"digital_card_art_token"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card's locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// support@lithic.com for questions.
	Pan string `json:"pan"`
	// Indicates if there are offline PIN changes pending card interaction with an
	// offline PIN terminal. Possible commands are: CHANGE_PIN, UNBLOCK_PIN. Applicable
	// only to cards issued in markets supporting offline PINs.
	PendingCommands []string `json:"pending_commands"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string `json:"product_id"`
	// If the card is a replacement for another card, the globally unique identifier
	// for the card that was replaced.
	ReplacementFor string              `json:"replacement_for,nullable"`
	JSON           cardNewResponseJSON `json:"-"`
}

// cardNewResponseJSON contains the JSON metadata for the struct [CardNewResponse]
type cardNewResponseJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	PinStatus           apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	CardholderCurrency  apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardNewResponseJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type CardNewResponseFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source. Funding source states: _ `ENABLED` - The funding
	// account is available to use for card creation and transactions. _ `PENDING` -
	// The funding account is still being verified e.g. bank micro-deposits
	// verification. \* `DELETED` - The founding account has been deleted.
	State CardNewResponseFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardNewResponseFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                     `json:"nickname"`
	JSON     cardNewResponseFundingJSON `json:"-"`
}

// cardNewResponseFundingJSON contains the JSON metadata for the struct
// [CardNewResponseFunding]
type cardNewResponseFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardNewResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardNewResponseFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type CardNewResponseFundingState string

const (
	CardNewResponseFundingStateDeleted CardNewResponseFundingState = "DELETED"
	CardNewResponseFundingStateEnabled CardNewResponseFundingState = "ENABLED"
	CardNewResponseFundingStatePending CardNewResponseFundingState = "PENDING"
)

func (r CardNewResponseFundingState) IsKnown() bool {
	switch r {
	case CardNewResponseFundingStateDeleted, CardNewResponseFundingStateEnabled, CardNewResponseFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type CardNewResponseFundingType string

const (
	CardNewResponseFundingTypeDepositoryChecking CardNewResponseFundingType = "DEPOSITORY_CHECKING"
	CardNewResponseFundingTypeDepositorySavings  CardNewResponseFundingType = "DEPOSITORY_SAVINGS"
)

func (r CardNewResponseFundingType) IsKnown() bool {
	switch r {
	case CardNewResponseFundingTypeDepositoryChecking, CardNewResponseFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type CardNewResponsePinStatus string

const (
	CardNewResponsePinStatusOk      CardNewResponsePinStatus = "OK"
	CardNewResponsePinStatusBlocked CardNewResponsePinStatus = "BLOCKED"
	CardNewResponsePinStatusNotSet  CardNewResponsePinStatus = "NOT_SET"
)

func (r CardNewResponsePinStatus) IsKnown() bool {
	switch r {
	case CardNewResponsePinStatusOk, CardNewResponsePinStatusBlocked, CardNewResponsePinStatusNotSet:
		return true
	}
	return false
}

// Spend limit duration
type CardNewResponseSpendLimitDuration string

const (
	CardNewResponseSpendLimitDurationAnnually    CardNewResponseSpendLimitDuration = "ANNUALLY"
	CardNewResponseSpendLimitDurationForever     CardNewResponseSpendLimitDuration = "FOREVER"
	CardNewResponseSpendLimitDurationMonthly     CardNewResponseSpendLimitDuration = "MONTHLY"
	CardNewResponseSpendLimitDurationTransaction CardNewResponseSpendLimitDuration = "TRANSACTION"
	CardNewResponseSpendLimitDurationDaily       CardNewResponseSpendLimitDuration = "DAILY"
)

func (r CardNewResponseSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardNewResponseSpendLimitDurationAnnually, CardNewResponseSpendLimitDurationForever, CardNewResponseSpendLimitDurationMonthly, CardNewResponseSpendLimitDurationTransaction, CardNewResponseSpendLimitDurationDaily:
		return true
	}
	return false
}

// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
// they match card and account parameters). _ `PAUSED` - Card will decline
// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
// manufacturing and fulfillment. Cards in this state can accept authorizations for
// e-commerce purchases, but not for "Card Present" purchases where the physical
// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
// transactions or can be added to mobile wallets. API clients should update the
// card's state to `OPEN` only after the cardholder confirms receipt of the card.
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardNewResponseState string

const (
	CardNewResponseStateClosed             CardNewResponseState = "CLOSED"
	CardNewResponseStateOpen               CardNewResponseState = "OPEN"
	CardNewResponseStatePaused             CardNewResponseState = "PAUSED"
	CardNewResponseStatePendingActivation  CardNewResponseState = "PENDING_ACTIVATION"
	CardNewResponseStatePendingFulfillment CardNewResponseState = "PENDING_FULFILLMENT"
)

func (r CardNewResponseState) IsKnown() bool {
	switch r {
	case CardNewResponseStateClosed, CardNewResponseStateOpen, CardNewResponseStatePaused, CardNewResponseStatePendingActivation, CardNewResponseStatePendingFulfillment:
		return true
	}
	return false
}

// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
// functionality. _ `SINGLE_USE` - Card is closed upon first successful
// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
// VIRTUAL instead.
type CardNewResponseType string

const (
	CardNewResponseTypeMerchantLocked CardNewResponseType = "MERCHANT_LOCKED"
	CardNewResponseTypePhysical       CardNewResponseType = "PHYSICAL"
	CardNewResponseTypeSingleUse      CardNewResponseType = "SINGLE_USE"
	CardNewResponseTypeVirtual        CardNewResponseType = "VIRTUAL"
	CardNewResponseTypeUnlocked       CardNewResponseType = "UNLOCKED"
	CardNewResponseTypeDigitalWallet  CardNewResponseType = "DIGITAL_WALLET"
)

func (r CardNewResponseType) IsKnown() bool {
	switch r {
	case CardNewResponseTypeMerchantLocked, CardNewResponseTypePhysical, CardNewResponseTypeSingleUse, CardNewResponseTypeVirtual, CardNewResponseTypeUnlocked, CardNewResponseTypeDigitalWallet:
		return true
	}
	return false
}

// Card details with potentially PCI sensitive information for Enterprise customers
type CardGetResponse struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding CardGetResponseFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus CardGetResponsePinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration
	SpendLimitDuration CardGetResponseSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
	// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
	// they match card and account parameters). _ `PAUSED` - Card will decline
	// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
	// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
	// manufacturing and fulfillment. Cards in this state can accept authorizations for
	// e-commerce purchases, but not for "Card Present" purchases where the physical
	// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
	// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
	// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
	// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	// transactions or can be added to mobile wallets. API clients should update the
	// card's state to `OPEN` only after the cardholder confirms receipt of the card.
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardGetResponseState `json:"state,required"`
	// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
	// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
	// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
	// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
	// functionality. _ `SINGLE_USE` - Card is closed upon first successful
	// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
	// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
	// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
	// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	// VIRTUAL instead.
	Type CardGetResponseType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user's digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use.
	DigitalCardArtToken string `json:"digital_card_art_token"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card's locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// support@lithic.com for questions.
	Pan string `json:"pan"`
	// Indicates if there are offline PIN changes pending card interaction with an
	// offline PIN terminal. Possible commands are: CHANGE_PIN, UNBLOCK_PIN. Applicable
	// only to cards issued in markets supporting offline PINs.
	PendingCommands []string `json:"pending_commands"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string `json:"product_id"`
	// If the card is a replacement for another card, the globally unique identifier
	// for the card that was replaced.
	ReplacementFor string              `json:"replacement_for,nullable"`
	JSON           cardGetResponseJSON `json:"-"`
}

// cardGetResponseJSON contains the JSON metadata for the struct [CardGetResponse]
type cardGetResponseJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	PinStatus           apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	CardholderCurrency  apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardGetResponseJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type CardGetResponseFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source. Funding source states: _ `ENABLED` - The funding
	// account is available to use for card creation and transactions. _ `PENDING` -
	// The funding account is still being verified e.g. bank micro-deposits
	// verification. \* `DELETED` - The founding account has been deleted.
	State CardGetResponseFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardGetResponseFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                     `json:"nickname"`
	JSON     cardGetResponseFundingJSON `json:"-"`
}

// cardGetResponseFundingJSON contains the JSON metadata for the struct
// [CardGetResponseFunding]
type cardGetResponseFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardGetResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardGetResponseFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type CardGetResponseFundingState string

const (
	CardGetResponseFundingStateDeleted CardGetResponseFundingState = "DELETED"
	CardGetResponseFundingStateEnabled CardGetResponseFundingState = "ENABLED"
	CardGetResponseFundingStatePending CardGetResponseFundingState = "PENDING"
)

func (r CardGetResponseFundingState) IsKnown() bool {
	switch r {
	case CardGetResponseFundingStateDeleted, CardGetResponseFundingStateEnabled, CardGetResponseFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type CardGetResponseFundingType string

const (
	CardGetResponseFundingTypeDepositoryChecking CardGetResponseFundingType = "DEPOSITORY_CHECKING"
	CardGetResponseFundingTypeDepositorySavings  CardGetResponseFundingType = "DEPOSITORY_SAVINGS"
)

func (r CardGetResponseFundingType) IsKnown() bool {
	switch r {
	case CardGetResponseFundingTypeDepositoryChecking, CardGetResponseFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type CardGetResponsePinStatus string

const (
	CardGetResponsePinStatusOk      CardGetResponsePinStatus = "OK"
	CardGetResponsePinStatusBlocked CardGetResponsePinStatus = "BLOCKED"
	CardGetResponsePinStatusNotSet  CardGetResponsePinStatus = "NOT_SET"
)

func (r CardGetResponsePinStatus) IsKnown() bool {
	switch r {
	case CardGetResponsePinStatusOk, CardGetResponsePinStatusBlocked, CardGetResponsePinStatusNotSet:
		return true
	}
	return false
}

// Spend limit duration
type CardGetResponseSpendLimitDuration string

const (
	CardGetResponseSpendLimitDurationAnnually    CardGetResponseSpendLimitDuration = "ANNUALLY"
	CardGetResponseSpendLimitDurationForever     CardGetResponseSpendLimitDuration = "FOREVER"
	CardGetResponseSpendLimitDurationMonthly     CardGetResponseSpendLimitDuration = "MONTHLY"
	CardGetResponseSpendLimitDurationTransaction CardGetResponseSpendLimitDuration = "TRANSACTION"
	CardGetResponseSpendLimitDurationDaily       CardGetResponseSpendLimitDuration = "DAILY"
)

func (r CardGetResponseSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardGetResponseSpendLimitDurationAnnually, CardGetResponseSpendLimitDurationForever, CardGetResponseSpendLimitDurationMonthly, CardGetResponseSpendLimitDurationTransaction, CardGetResponseSpendLimitDurationDaily:
		return true
	}
	return false
}

// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
// they match card and account parameters). _ `PAUSED` - Card will decline
// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
// manufacturing and fulfillment. Cards in this state can accept authorizations for
// e-commerce purchases, but not for "Card Present" purchases where the physical
// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
// transactions or can be added to mobile wallets. API clients should update the
// card's state to `OPEN` only after the cardholder confirms receipt of the card.
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardGetResponseState string

const (
	CardGetResponseStateClosed             CardGetResponseState = "CLOSED"
	CardGetResponseStateOpen               CardGetResponseState = "OPEN"
	CardGetResponseStatePaused             CardGetResponseState = "PAUSED"
	CardGetResponseStatePendingActivation  CardGetResponseState = "PENDING_ACTIVATION"
	CardGetResponseStatePendingFulfillment CardGetResponseState = "PENDING_FULFILLMENT"
)

func (r CardGetResponseState) IsKnown() bool {
	switch r {
	case CardGetResponseStateClosed, CardGetResponseStateOpen, CardGetResponseStatePaused, CardGetResponseStatePendingActivation, CardGetResponseStatePendingFulfillment:
		return true
	}
	return false
}

// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
// functionality. _ `SINGLE_USE` - Card is closed upon first successful
// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
// VIRTUAL instead.
type CardGetResponseType string

const (
	CardGetResponseTypeMerchantLocked CardGetResponseType = "MERCHANT_LOCKED"
	CardGetResponseTypePhysical       CardGetResponseType = "PHYSICAL"
	CardGetResponseTypeSingleUse      CardGetResponseType = "SINGLE_USE"
	CardGetResponseTypeVirtual        CardGetResponseType = "VIRTUAL"
	CardGetResponseTypeUnlocked       CardGetResponseType = "UNLOCKED"
	CardGetResponseTypeDigitalWallet  CardGetResponseType = "DIGITAL_WALLET"
)

func (r CardGetResponseType) IsKnown() bool {
	switch r {
	case CardGetResponseTypeMerchantLocked, CardGetResponseTypePhysical, CardGetResponseTypeSingleUse, CardGetResponseTypeVirtual, CardGetResponseTypeUnlocked, CardGetResponseTypeDigitalWallet:
		return true
	}
	return false
}

// Card details with potentially PCI sensitive information for Enterprise customers
type CardUpdateResponse struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding CardUpdateResponseFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus CardUpdateResponsePinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration
	SpendLimitDuration CardUpdateResponseSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
	// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
	// they match card and account parameters). _ `PAUSED` - Card will decline
	// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
	// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
	// manufacturing and fulfillment. Cards in this state can accept authorizations for
	// e-commerce purchases, but not for "Card Present" purchases where the physical
	// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
	// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
	// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
	// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	// transactions or can be added to mobile wallets. API clients should update the
	// card's state to `OPEN` only after the cardholder confirms receipt of the card.
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardUpdateResponseState `json:"state,required"`
	// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
	// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
	// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
	// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
	// functionality. _ `SINGLE_USE` - Card is closed upon first successful
	// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
	// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
	// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
	// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	// VIRTUAL instead.
	Type CardUpdateResponseType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user's digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use.
	DigitalCardArtToken string `json:"digital_card_art_token"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card's locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// support@lithic.com for questions.
	Pan string `json:"pan"`
	// Indicates if there are offline PIN changes pending card interaction with an
	// offline PIN terminal. Possible commands are: CHANGE_PIN, UNBLOCK_PIN. Applicable
	// only to cards issued in markets supporting offline PINs.
	PendingCommands []string `json:"pending_commands"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string `json:"product_id"`
	// If the card is a replacement for another card, the globally unique identifier
	// for the card that was replaced.
	ReplacementFor string                 `json:"replacement_for,nullable"`
	JSON           cardUpdateResponseJSON `json:"-"`
}

// cardUpdateResponseJSON contains the JSON metadata for the struct
// [CardUpdateResponse]
type cardUpdateResponseJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	PinStatus           apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	CardholderCurrency  apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type CardUpdateResponseFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source. Funding source states: _ `ENABLED` - The funding
	// account is available to use for card creation and transactions. _ `PENDING` -
	// The funding account is still being verified e.g. bank micro-deposits
	// verification. \* `DELETED` - The founding account has been deleted.
	State CardUpdateResponseFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardUpdateResponseFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                        `json:"nickname"`
	JSON     cardUpdateResponseFundingJSON `json:"-"`
}

// cardUpdateResponseFundingJSON contains the JSON metadata for the struct
// [CardUpdateResponseFunding]
type cardUpdateResponseFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardUpdateResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardUpdateResponseFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type CardUpdateResponseFundingState string

const (
	CardUpdateResponseFundingStateDeleted CardUpdateResponseFundingState = "DELETED"
	CardUpdateResponseFundingStateEnabled CardUpdateResponseFundingState = "ENABLED"
	CardUpdateResponseFundingStatePending CardUpdateResponseFundingState = "PENDING"
)

func (r CardUpdateResponseFundingState) IsKnown() bool {
	switch r {
	case CardUpdateResponseFundingStateDeleted, CardUpdateResponseFundingStateEnabled, CardUpdateResponseFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type CardUpdateResponseFundingType string

const (
	CardUpdateResponseFundingTypeDepositoryChecking CardUpdateResponseFundingType = "DEPOSITORY_CHECKING"
	CardUpdateResponseFundingTypeDepositorySavings  CardUpdateResponseFundingType = "DEPOSITORY_SAVINGS"
)

func (r CardUpdateResponseFundingType) IsKnown() bool {
	switch r {
	case CardUpdateResponseFundingTypeDepositoryChecking, CardUpdateResponseFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type CardUpdateResponsePinStatus string

const (
	CardUpdateResponsePinStatusOk      CardUpdateResponsePinStatus = "OK"
	CardUpdateResponsePinStatusBlocked CardUpdateResponsePinStatus = "BLOCKED"
	CardUpdateResponsePinStatusNotSet  CardUpdateResponsePinStatus = "NOT_SET"
)

func (r CardUpdateResponsePinStatus) IsKnown() bool {
	switch r {
	case CardUpdateResponsePinStatusOk, CardUpdateResponsePinStatusBlocked, CardUpdateResponsePinStatusNotSet:
		return true
	}
	return false
}

// Spend limit duration
type CardUpdateResponseSpendLimitDuration string

const (
	CardUpdateResponseSpendLimitDurationAnnually    CardUpdateResponseSpendLimitDuration = "ANNUALLY"
	CardUpdateResponseSpendLimitDurationForever     CardUpdateResponseSpendLimitDuration = "FOREVER"
	CardUpdateResponseSpendLimitDurationMonthly     CardUpdateResponseSpendLimitDuration = "MONTHLY"
	CardUpdateResponseSpendLimitDurationTransaction CardUpdateResponseSpendLimitDuration = "TRANSACTION"
	CardUpdateResponseSpendLimitDurationDaily       CardUpdateResponseSpendLimitDuration = "DAILY"
)

func (r CardUpdateResponseSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardUpdateResponseSpendLimitDurationAnnually, CardUpdateResponseSpendLimitDurationForever, CardUpdateResponseSpendLimitDurationMonthly, CardUpdateResponseSpendLimitDurationTransaction, CardUpdateResponseSpendLimitDurationDaily:
		return true
	}
	return false
}

// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
// they match card and account parameters). _ `PAUSED` - Card will decline
// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
// manufacturing and fulfillment. Cards in this state can accept authorizations for
// e-commerce purchases, but not for "Card Present" purchases where the physical
// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
// transactions or can be added to mobile wallets. API clients should update the
// card's state to `OPEN` only after the cardholder confirms receipt of the card.
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardUpdateResponseState string

const (
	CardUpdateResponseStateClosed             CardUpdateResponseState = "CLOSED"
	CardUpdateResponseStateOpen               CardUpdateResponseState = "OPEN"
	CardUpdateResponseStatePaused             CardUpdateResponseState = "PAUSED"
	CardUpdateResponseStatePendingActivation  CardUpdateResponseState = "PENDING_ACTIVATION"
	CardUpdateResponseStatePendingFulfillment CardUpdateResponseState = "PENDING_FULFILLMENT"
)

func (r CardUpdateResponseState) IsKnown() bool {
	switch r {
	case CardUpdateResponseStateClosed, CardUpdateResponseStateOpen, CardUpdateResponseStatePaused, CardUpdateResponseStatePendingActivation, CardUpdateResponseStatePendingFulfillment:
		return true
	}
	return false
}

// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
// functionality. _ `SINGLE_USE` - Card is closed upon first successful
// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
// VIRTUAL instead.
type CardUpdateResponseType string

const (
	CardUpdateResponseTypeMerchantLocked CardUpdateResponseType = "MERCHANT_LOCKED"
	CardUpdateResponseTypePhysical       CardUpdateResponseType = "PHYSICAL"
	CardUpdateResponseTypeSingleUse      CardUpdateResponseType = "SINGLE_USE"
	CardUpdateResponseTypeVirtual        CardUpdateResponseType = "VIRTUAL"
	CardUpdateResponseTypeUnlocked       CardUpdateResponseType = "UNLOCKED"
	CardUpdateResponseTypeDigitalWallet  CardUpdateResponseType = "DIGITAL_WALLET"
)

func (r CardUpdateResponseType) IsKnown() bool {
	switch r {
	case CardUpdateResponseTypeMerchantLocked, CardUpdateResponseTypePhysical, CardUpdateResponseTypeSingleUse, CardUpdateResponseTypeVirtual, CardUpdateResponseTypeUnlocked, CardUpdateResponseTypeDigitalWallet:
		return true
	}
	return false
}

// Card details without PCI information
type CardListResponse struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding CardListResponseFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus CardListResponsePinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration
	SpendLimitDuration CardListResponseSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
	// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
	// they match card and account parameters). _ `PAUSED` - Card will decline
	// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
	// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
	// manufacturing and fulfillment. Cards in this state can accept authorizations for
	// e-commerce purchases, but not for "Card Present" purchases where the physical
	// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
	// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
	// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
	// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	// transactions or can be added to mobile wallets. API clients should update the
	// card's state to `OPEN` only after the cardholder confirms receipt of the card.
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardListResponseState `json:"state,required"`
	// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
	// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
	// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
	// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
	// functionality. _ `SINGLE_USE` - Card is closed upon first successful
	// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
	// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
	// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
	// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	// VIRTUAL instead.
	Type CardListResponseType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Specifies the digital card art to be displayed in the user's digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use.
	DigitalCardArtToken string `json:"digital_card_art_token"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card's locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card.
	Memo string `json:"memo"`
	// Indicates if there are offline PIN changes pending card interaction with an
	// offline PIN terminal. Possible commands are: CHANGE_PIN, UNBLOCK_PIN. Applicable
	// only to cards issued in markets supporting offline PINs.
	PendingCommands []string `json:"pending_commands"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string `json:"product_id"`
	// If the card is a replacement for another card, the globally unique identifier
	// for the card that was replaced.
	ReplacementFor string               `json:"replacement_for,nullable"`
	JSON           cardListResponseJSON `json:"-"`
}

// cardListResponseJSON contains the JSON metadata for the struct
// [CardListResponse]
type cardListResponseJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	PinStatus           apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	CardholderCurrency  apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardListResponseJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type CardListResponseFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source. Funding source states: _ `ENABLED` - The funding
	// account is available to use for card creation and transactions. _ `PENDING` -
	// The funding account is still being verified e.g. bank micro-deposits
	// verification. \* `DELETED` - The founding account has been deleted.
	State CardListResponseFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardListResponseFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                      `json:"nickname"`
	JSON     cardListResponseFundingJSON `json:"-"`
}

// cardListResponseFundingJSON contains the JSON metadata for the struct
// [CardListResponseFunding]
type cardListResponseFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardListResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardListResponseFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type CardListResponseFundingState string

const (
	CardListResponseFundingStateDeleted CardListResponseFundingState = "DELETED"
	CardListResponseFundingStateEnabled CardListResponseFundingState = "ENABLED"
	CardListResponseFundingStatePending CardListResponseFundingState = "PENDING"
)

func (r CardListResponseFundingState) IsKnown() bool {
	switch r {
	case CardListResponseFundingStateDeleted, CardListResponseFundingStateEnabled, CardListResponseFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type CardListResponseFundingType string

const (
	CardListResponseFundingTypeDepositoryChecking CardListResponseFundingType = "DEPOSITORY_CHECKING"
	CardListResponseFundingTypeDepositorySavings  CardListResponseFundingType = "DEPOSITORY_SAVINGS"
)

func (r CardListResponseFundingType) IsKnown() bool {
	switch r {
	case CardListResponseFundingTypeDepositoryChecking, CardListResponseFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type CardListResponsePinStatus string

const (
	CardListResponsePinStatusOk      CardListResponsePinStatus = "OK"
	CardListResponsePinStatusBlocked CardListResponsePinStatus = "BLOCKED"
	CardListResponsePinStatusNotSet  CardListResponsePinStatus = "NOT_SET"
)

func (r CardListResponsePinStatus) IsKnown() bool {
	switch r {
	case CardListResponsePinStatusOk, CardListResponsePinStatusBlocked, CardListResponsePinStatusNotSet:
		return true
	}
	return false
}

// Spend limit duration
type CardListResponseSpendLimitDuration string

const (
	CardListResponseSpendLimitDurationAnnually    CardListResponseSpendLimitDuration = "ANNUALLY"
	CardListResponseSpendLimitDurationForever     CardListResponseSpendLimitDuration = "FOREVER"
	CardListResponseSpendLimitDurationMonthly     CardListResponseSpendLimitDuration = "MONTHLY"
	CardListResponseSpendLimitDurationTransaction CardListResponseSpendLimitDuration = "TRANSACTION"
	CardListResponseSpendLimitDurationDaily       CardListResponseSpendLimitDuration = "DAILY"
)

func (r CardListResponseSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardListResponseSpendLimitDurationAnnually, CardListResponseSpendLimitDurationForever, CardListResponseSpendLimitDurationMonthly, CardListResponseSpendLimitDurationTransaction, CardListResponseSpendLimitDurationDaily:
		return true
	}
	return false
}

// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
// they match card and account parameters). _ `PAUSED` - Card will decline
// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
// manufacturing and fulfillment. Cards in this state can accept authorizations for
// e-commerce purchases, but not for "Card Present" purchases where the physical
// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
// transactions or can be added to mobile wallets. API clients should update the
// card's state to `OPEN` only after the cardholder confirms receipt of the card.
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardListResponseState string

const (
	CardListResponseStateClosed             CardListResponseState = "CLOSED"
	CardListResponseStateOpen               CardListResponseState = "OPEN"
	CardListResponseStatePaused             CardListResponseState = "PAUSED"
	CardListResponseStatePendingActivation  CardListResponseState = "PENDING_ACTIVATION"
	CardListResponseStatePendingFulfillment CardListResponseState = "PENDING_FULFILLMENT"
)

func (r CardListResponseState) IsKnown() bool {
	switch r {
	case CardListResponseStateClosed, CardListResponseStateOpen, CardListResponseStatePaused, CardListResponseStatePendingActivation, CardListResponseStatePendingFulfillment:
		return true
	}
	return false
}

// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
// functionality. _ `SINGLE_USE` - Card is closed upon first successful
// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
// VIRTUAL instead.
type CardListResponseType string

const (
	CardListResponseTypeMerchantLocked CardListResponseType = "MERCHANT_LOCKED"
	CardListResponseTypePhysical       CardListResponseType = "PHYSICAL"
	CardListResponseTypeSingleUse      CardListResponseType = "SINGLE_USE"
	CardListResponseTypeVirtual        CardListResponseType = "VIRTUAL"
	CardListResponseTypeUnlocked       CardListResponseType = "UNLOCKED"
	CardListResponseTypeDigitalWallet  CardListResponseType = "DIGITAL_WALLET"
)

func (r CardListResponseType) IsKnown() bool {
	switch r {
	case CardListResponseTypeMerchantLocked, CardListResponseTypePhysical, CardListResponseTypeSingleUse, CardListResponseTypeVirtual, CardListResponseTypeUnlocked, CardListResponseTypeDigitalWallet:
		return true
	}
	return false
}

// Card details with potentially PCI sensitive information for Enterprise customers
type CardConvertPhysicalResponse struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding CardConvertPhysicalResponseFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus CardConvertPhysicalResponsePinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration
	SpendLimitDuration CardConvertPhysicalResponseSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
	// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
	// they match card and account parameters). _ `PAUSED` - Card will decline
	// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
	// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
	// manufacturing and fulfillment. Cards in this state can accept authorizations for
	// e-commerce purchases, but not for "Card Present" purchases where the physical
	// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
	// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
	// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
	// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	// transactions or can be added to mobile wallets. API clients should update the
	// card's state to `OPEN` only after the cardholder confirms receipt of the card.
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardConvertPhysicalResponseState `json:"state,required"`
	// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
	// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
	// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
	// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
	// functionality. _ `SINGLE_USE` - Card is closed upon first successful
	// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
	// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
	// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
	// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	// VIRTUAL instead.
	Type CardConvertPhysicalResponseType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user's digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use.
	DigitalCardArtToken string `json:"digital_card_art_token"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card's locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// support@lithic.com for questions.
	Pan string `json:"pan"`
	// Indicates if there are offline PIN changes pending card interaction with an
	// offline PIN terminal. Possible commands are: CHANGE_PIN, UNBLOCK_PIN. Applicable
	// only to cards issued in markets supporting offline PINs.
	PendingCommands []string `json:"pending_commands"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string `json:"product_id"`
	// If the card is a replacement for another card, the globally unique identifier
	// for the card that was replaced.
	ReplacementFor string                          `json:"replacement_for,nullable"`
	JSON           cardConvertPhysicalResponseJSON `json:"-"`
}

// cardConvertPhysicalResponseJSON contains the JSON metadata for the struct
// [CardConvertPhysicalResponse]
type cardConvertPhysicalResponseJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	PinStatus           apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	CardholderCurrency  apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardConvertPhysicalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardConvertPhysicalResponseJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type CardConvertPhysicalResponseFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source. Funding source states: _ `ENABLED` - The funding
	// account is available to use for card creation and transactions. _ `PENDING` -
	// The funding account is still being verified e.g. bank micro-deposits
	// verification. \* `DELETED` - The founding account has been deleted.
	State CardConvertPhysicalResponseFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardConvertPhysicalResponseFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                                 `json:"nickname"`
	JSON     cardConvertPhysicalResponseFundingJSON `json:"-"`
}

// cardConvertPhysicalResponseFundingJSON contains the JSON metadata for the struct
// [CardConvertPhysicalResponseFunding]
type cardConvertPhysicalResponseFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardConvertPhysicalResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardConvertPhysicalResponseFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type CardConvertPhysicalResponseFundingState string

const (
	CardConvertPhysicalResponseFundingStateDeleted CardConvertPhysicalResponseFundingState = "DELETED"
	CardConvertPhysicalResponseFundingStateEnabled CardConvertPhysicalResponseFundingState = "ENABLED"
	CardConvertPhysicalResponseFundingStatePending CardConvertPhysicalResponseFundingState = "PENDING"
)

func (r CardConvertPhysicalResponseFundingState) IsKnown() bool {
	switch r {
	case CardConvertPhysicalResponseFundingStateDeleted, CardConvertPhysicalResponseFundingStateEnabled, CardConvertPhysicalResponseFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type CardConvertPhysicalResponseFundingType string

const (
	CardConvertPhysicalResponseFundingTypeDepositoryChecking CardConvertPhysicalResponseFundingType = "DEPOSITORY_CHECKING"
	CardConvertPhysicalResponseFundingTypeDepositorySavings  CardConvertPhysicalResponseFundingType = "DEPOSITORY_SAVINGS"
)

func (r CardConvertPhysicalResponseFundingType) IsKnown() bool {
	switch r {
	case CardConvertPhysicalResponseFundingTypeDepositoryChecking, CardConvertPhysicalResponseFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type CardConvertPhysicalResponsePinStatus string

const (
	CardConvertPhysicalResponsePinStatusOk      CardConvertPhysicalResponsePinStatus = "OK"
	CardConvertPhysicalResponsePinStatusBlocked CardConvertPhysicalResponsePinStatus = "BLOCKED"
	CardConvertPhysicalResponsePinStatusNotSet  CardConvertPhysicalResponsePinStatus = "NOT_SET"
)

func (r CardConvertPhysicalResponsePinStatus) IsKnown() bool {
	switch r {
	case CardConvertPhysicalResponsePinStatusOk, CardConvertPhysicalResponsePinStatusBlocked, CardConvertPhysicalResponsePinStatusNotSet:
		return true
	}
	return false
}

// Spend limit duration
type CardConvertPhysicalResponseSpendLimitDuration string

const (
	CardConvertPhysicalResponseSpendLimitDurationAnnually    CardConvertPhysicalResponseSpendLimitDuration = "ANNUALLY"
	CardConvertPhysicalResponseSpendLimitDurationForever     CardConvertPhysicalResponseSpendLimitDuration = "FOREVER"
	CardConvertPhysicalResponseSpendLimitDurationMonthly     CardConvertPhysicalResponseSpendLimitDuration = "MONTHLY"
	CardConvertPhysicalResponseSpendLimitDurationTransaction CardConvertPhysicalResponseSpendLimitDuration = "TRANSACTION"
	CardConvertPhysicalResponseSpendLimitDurationDaily       CardConvertPhysicalResponseSpendLimitDuration = "DAILY"
)

func (r CardConvertPhysicalResponseSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardConvertPhysicalResponseSpendLimitDurationAnnually, CardConvertPhysicalResponseSpendLimitDurationForever, CardConvertPhysicalResponseSpendLimitDurationMonthly, CardConvertPhysicalResponseSpendLimitDurationTransaction, CardConvertPhysicalResponseSpendLimitDurationDaily:
		return true
	}
	return false
}

// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
// they match card and account parameters). _ `PAUSED` - Card will decline
// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
// manufacturing and fulfillment. Cards in this state can accept authorizations for
// e-commerce purchases, but not for "Card Present" purchases where the physical
// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
// transactions or can be added to mobile wallets. API clients should update the
// card's state to `OPEN` only after the cardholder confirms receipt of the card.
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardConvertPhysicalResponseState string

const (
	CardConvertPhysicalResponseStateClosed             CardConvertPhysicalResponseState = "CLOSED"
	CardConvertPhysicalResponseStateOpen               CardConvertPhysicalResponseState = "OPEN"
	CardConvertPhysicalResponseStatePaused             CardConvertPhysicalResponseState = "PAUSED"
	CardConvertPhysicalResponseStatePendingActivation  CardConvertPhysicalResponseState = "PENDING_ACTIVATION"
	CardConvertPhysicalResponseStatePendingFulfillment CardConvertPhysicalResponseState = "PENDING_FULFILLMENT"
)

func (r CardConvertPhysicalResponseState) IsKnown() bool {
	switch r {
	case CardConvertPhysicalResponseStateClosed, CardConvertPhysicalResponseStateOpen, CardConvertPhysicalResponseStatePaused, CardConvertPhysicalResponseStatePendingActivation, CardConvertPhysicalResponseStatePendingFulfillment:
		return true
	}
	return false
}

// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
// functionality. _ `SINGLE_USE` - Card is closed upon first successful
// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
// VIRTUAL instead.
type CardConvertPhysicalResponseType string

const (
	CardConvertPhysicalResponseTypeMerchantLocked CardConvertPhysicalResponseType = "MERCHANT_LOCKED"
	CardConvertPhysicalResponseTypePhysical       CardConvertPhysicalResponseType = "PHYSICAL"
	CardConvertPhysicalResponseTypeSingleUse      CardConvertPhysicalResponseType = "SINGLE_USE"
	CardConvertPhysicalResponseTypeVirtual        CardConvertPhysicalResponseType = "VIRTUAL"
	CardConvertPhysicalResponseTypeUnlocked       CardConvertPhysicalResponseType = "UNLOCKED"
	CardConvertPhysicalResponseTypeDigitalWallet  CardConvertPhysicalResponseType = "DIGITAL_WALLET"
)

func (r CardConvertPhysicalResponseType) IsKnown() bool {
	switch r {
	case CardConvertPhysicalResponseTypeMerchantLocked, CardConvertPhysicalResponseTypePhysical, CardConvertPhysicalResponseTypeSingleUse, CardConvertPhysicalResponseTypeVirtual, CardConvertPhysicalResponseTypeUnlocked, CardConvertPhysicalResponseTypeDigitalWallet:
		return true
	}
	return false
}

type CardProvisionResponse struct {
	ProvisioningPayload string                    `json:"provisioning_payload"`
	JSON                cardProvisionResponseJSON `json:"-"`
}

// cardProvisionResponseJSON contains the JSON metadata for the struct
// [CardProvisionResponse]
type cardProvisionResponseJSON struct {
	ProvisioningPayload apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardProvisionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardProvisionResponseJSON) RawJSON() string {
	return r.raw
}

// Card details with potentially PCI sensitive information for Enterprise customers
type CardReissueResponse struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding CardReissueResponseFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus CardReissueResponsePinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration
	SpendLimitDuration CardReissueResponseSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
	// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
	// they match card and account parameters). _ `PAUSED` - Card will decline
	// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
	// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
	// manufacturing and fulfillment. Cards in this state can accept authorizations for
	// e-commerce purchases, but not for "Card Present" purchases where the physical
	// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
	// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
	// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
	// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	// transactions or can be added to mobile wallets. API clients should update the
	// card's state to `OPEN` only after the cardholder confirms receipt of the card.
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardReissueResponseState `json:"state,required"`
	// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
	// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
	// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
	// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
	// functionality. _ `SINGLE_USE` - Card is closed upon first successful
	// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
	// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
	// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
	// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	// VIRTUAL instead.
	Type CardReissueResponseType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user's digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use.
	DigitalCardArtToken string `json:"digital_card_art_token"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card's locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// support@lithic.com for questions.
	Pan string `json:"pan"`
	// Indicates if there are offline PIN changes pending card interaction with an
	// offline PIN terminal. Possible commands are: CHANGE_PIN, UNBLOCK_PIN. Applicable
	// only to cards issued in markets supporting offline PINs.
	PendingCommands []string `json:"pending_commands"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string `json:"product_id"`
	// If the card is a replacement for another card, the globally unique identifier
	// for the card that was replaced.
	ReplacementFor string                  `json:"replacement_for,nullable"`
	JSON           cardReissueResponseJSON `json:"-"`
}

// cardReissueResponseJSON contains the JSON metadata for the struct
// [CardReissueResponse]
type cardReissueResponseJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	PinStatus           apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	CardholderCurrency  apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardReissueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardReissueResponseJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type CardReissueResponseFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source. Funding source states: _ `ENABLED` - The funding
	// account is available to use for card creation and transactions. _ `PENDING` -
	// The funding account is still being verified e.g. bank micro-deposits
	// verification. \* `DELETED` - The founding account has been deleted.
	State CardReissueResponseFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardReissueResponseFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                         `json:"nickname"`
	JSON     cardReissueResponseFundingJSON `json:"-"`
}

// cardReissueResponseFundingJSON contains the JSON metadata for the struct
// [CardReissueResponseFunding]
type cardReissueResponseFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardReissueResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardReissueResponseFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type CardReissueResponseFundingState string

const (
	CardReissueResponseFundingStateDeleted CardReissueResponseFundingState = "DELETED"
	CardReissueResponseFundingStateEnabled CardReissueResponseFundingState = "ENABLED"
	CardReissueResponseFundingStatePending CardReissueResponseFundingState = "PENDING"
)

func (r CardReissueResponseFundingState) IsKnown() bool {
	switch r {
	case CardReissueResponseFundingStateDeleted, CardReissueResponseFundingStateEnabled, CardReissueResponseFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type CardReissueResponseFundingType string

const (
	CardReissueResponseFundingTypeDepositoryChecking CardReissueResponseFundingType = "DEPOSITORY_CHECKING"
	CardReissueResponseFundingTypeDepositorySavings  CardReissueResponseFundingType = "DEPOSITORY_SAVINGS"
)

func (r CardReissueResponseFundingType) IsKnown() bool {
	switch r {
	case CardReissueResponseFundingTypeDepositoryChecking, CardReissueResponseFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type CardReissueResponsePinStatus string

const (
	CardReissueResponsePinStatusOk      CardReissueResponsePinStatus = "OK"
	CardReissueResponsePinStatusBlocked CardReissueResponsePinStatus = "BLOCKED"
	CardReissueResponsePinStatusNotSet  CardReissueResponsePinStatus = "NOT_SET"
)

func (r CardReissueResponsePinStatus) IsKnown() bool {
	switch r {
	case CardReissueResponsePinStatusOk, CardReissueResponsePinStatusBlocked, CardReissueResponsePinStatusNotSet:
		return true
	}
	return false
}

// Spend limit duration
type CardReissueResponseSpendLimitDuration string

const (
	CardReissueResponseSpendLimitDurationAnnually    CardReissueResponseSpendLimitDuration = "ANNUALLY"
	CardReissueResponseSpendLimitDurationForever     CardReissueResponseSpendLimitDuration = "FOREVER"
	CardReissueResponseSpendLimitDurationMonthly     CardReissueResponseSpendLimitDuration = "MONTHLY"
	CardReissueResponseSpendLimitDurationTransaction CardReissueResponseSpendLimitDuration = "TRANSACTION"
	CardReissueResponseSpendLimitDurationDaily       CardReissueResponseSpendLimitDuration = "DAILY"
)

func (r CardReissueResponseSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardReissueResponseSpendLimitDurationAnnually, CardReissueResponseSpendLimitDurationForever, CardReissueResponseSpendLimitDurationMonthly, CardReissueResponseSpendLimitDurationTransaction, CardReissueResponseSpendLimitDurationDaily:
		return true
	}
	return false
}

// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
// they match card and account parameters). _ `PAUSED` - Card will decline
// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
// manufacturing and fulfillment. Cards in this state can accept authorizations for
// e-commerce purchases, but not for "Card Present" purchases where the physical
// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
// transactions or can be added to mobile wallets. API clients should update the
// card's state to `OPEN` only after the cardholder confirms receipt of the card.
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardReissueResponseState string

const (
	CardReissueResponseStateClosed             CardReissueResponseState = "CLOSED"
	CardReissueResponseStateOpen               CardReissueResponseState = "OPEN"
	CardReissueResponseStatePaused             CardReissueResponseState = "PAUSED"
	CardReissueResponseStatePendingActivation  CardReissueResponseState = "PENDING_ACTIVATION"
	CardReissueResponseStatePendingFulfillment CardReissueResponseState = "PENDING_FULFILLMENT"
)

func (r CardReissueResponseState) IsKnown() bool {
	switch r {
	case CardReissueResponseStateClosed, CardReissueResponseStateOpen, CardReissueResponseStatePaused, CardReissueResponseStatePendingActivation, CardReissueResponseStatePendingFulfillment:
		return true
	}
	return false
}

// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
// functionality. _ `SINGLE_USE` - Card is closed upon first successful
// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
// VIRTUAL instead.
type CardReissueResponseType string

const (
	CardReissueResponseTypeMerchantLocked CardReissueResponseType = "MERCHANT_LOCKED"
	CardReissueResponseTypePhysical       CardReissueResponseType = "PHYSICAL"
	CardReissueResponseTypeSingleUse      CardReissueResponseType = "SINGLE_USE"
	CardReissueResponseTypeVirtual        CardReissueResponseType = "VIRTUAL"
	CardReissueResponseTypeUnlocked       CardReissueResponseType = "UNLOCKED"
	CardReissueResponseTypeDigitalWallet  CardReissueResponseType = "DIGITAL_WALLET"
)

func (r CardReissueResponseType) IsKnown() bool {
	switch r {
	case CardReissueResponseTypeMerchantLocked, CardReissueResponseTypePhysical, CardReissueResponseTypeSingleUse, CardReissueResponseTypeVirtual, CardReissueResponseTypeUnlocked, CardReissueResponseTypeDigitalWallet:
		return true
	}
	return false
}

// Card details with potentially PCI sensitive information for Enterprise customers
type CardRenewResponse struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding CardRenewResponseFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus CardRenewResponsePinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration
	SpendLimitDuration CardRenewResponseSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
	// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
	// they match card and account parameters). _ `PAUSED` - Card will decline
	// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
	// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
	// manufacturing and fulfillment. Cards in this state can accept authorizations for
	// e-commerce purchases, but not for "Card Present" purchases where the physical
	// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
	// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
	// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
	// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	// transactions or can be added to mobile wallets. API clients should update the
	// card's state to `OPEN` only after the cardholder confirms receipt of the card.
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardRenewResponseState `json:"state,required"`
	// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
	// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
	// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
	// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
	// functionality. _ `SINGLE_USE` - Card is closed upon first successful
	// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
	// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
	// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
	// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	// VIRTUAL instead.
	Type CardRenewResponseType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user's digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use.
	DigitalCardArtToken string `json:"digital_card_art_token"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card's locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// support@lithic.com for questions.
	Pan string `json:"pan"`
	// Indicates if there are offline PIN changes pending card interaction with an
	// offline PIN terminal. Possible commands are: CHANGE_PIN, UNBLOCK_PIN. Applicable
	// only to cards issued in markets supporting offline PINs.
	PendingCommands []string `json:"pending_commands"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string `json:"product_id"`
	// If the card is a replacement for another card, the globally unique identifier
	// for the card that was replaced.
	ReplacementFor string                `json:"replacement_for,nullable"`
	JSON           cardRenewResponseJSON `json:"-"`
}

// cardRenewResponseJSON contains the JSON metadata for the struct
// [CardRenewResponse]
type cardRenewResponseJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	PinStatus           apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	CardholderCurrency  apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardRenewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardRenewResponseJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type CardRenewResponseFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source. Funding source states: _ `ENABLED` - The funding
	// account is available to use for card creation and transactions. _ `PENDING` -
	// The funding account is still being verified e.g. bank micro-deposits
	// verification. \* `DELETED` - The founding account has been deleted.
	State CardRenewResponseFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardRenewResponseFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                       `json:"nickname"`
	JSON     cardRenewResponseFundingJSON `json:"-"`
}

// cardRenewResponseFundingJSON contains the JSON metadata for the struct
// [CardRenewResponseFunding]
type cardRenewResponseFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardRenewResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardRenewResponseFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type CardRenewResponseFundingState string

const (
	CardRenewResponseFundingStateDeleted CardRenewResponseFundingState = "DELETED"
	CardRenewResponseFundingStateEnabled CardRenewResponseFundingState = "ENABLED"
	CardRenewResponseFundingStatePending CardRenewResponseFundingState = "PENDING"
)

func (r CardRenewResponseFundingState) IsKnown() bool {
	switch r {
	case CardRenewResponseFundingStateDeleted, CardRenewResponseFundingStateEnabled, CardRenewResponseFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type CardRenewResponseFundingType string

const (
	CardRenewResponseFundingTypeDepositoryChecking CardRenewResponseFundingType = "DEPOSITORY_CHECKING"
	CardRenewResponseFundingTypeDepositorySavings  CardRenewResponseFundingType = "DEPOSITORY_SAVINGS"
)

func (r CardRenewResponseFundingType) IsKnown() bool {
	switch r {
	case CardRenewResponseFundingTypeDepositoryChecking, CardRenewResponseFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type CardRenewResponsePinStatus string

const (
	CardRenewResponsePinStatusOk      CardRenewResponsePinStatus = "OK"
	CardRenewResponsePinStatusBlocked CardRenewResponsePinStatus = "BLOCKED"
	CardRenewResponsePinStatusNotSet  CardRenewResponsePinStatus = "NOT_SET"
)

func (r CardRenewResponsePinStatus) IsKnown() bool {
	switch r {
	case CardRenewResponsePinStatusOk, CardRenewResponsePinStatusBlocked, CardRenewResponsePinStatusNotSet:
		return true
	}
	return false
}

// Spend limit duration
type CardRenewResponseSpendLimitDuration string

const (
	CardRenewResponseSpendLimitDurationAnnually    CardRenewResponseSpendLimitDuration = "ANNUALLY"
	CardRenewResponseSpendLimitDurationForever     CardRenewResponseSpendLimitDuration = "FOREVER"
	CardRenewResponseSpendLimitDurationMonthly     CardRenewResponseSpendLimitDuration = "MONTHLY"
	CardRenewResponseSpendLimitDurationTransaction CardRenewResponseSpendLimitDuration = "TRANSACTION"
	CardRenewResponseSpendLimitDurationDaily       CardRenewResponseSpendLimitDuration = "DAILY"
)

func (r CardRenewResponseSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardRenewResponseSpendLimitDurationAnnually, CardRenewResponseSpendLimitDurationForever, CardRenewResponseSpendLimitDurationMonthly, CardRenewResponseSpendLimitDurationTransaction, CardRenewResponseSpendLimitDurationDaily:
		return true
	}
	return false
}

// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
// they match card and account parameters). _ `PAUSED` - Card will decline
// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
// manufacturing and fulfillment. Cards in this state can accept authorizations for
// e-commerce purchases, but not for "Card Present" purchases where the physical
// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
// transactions or can be added to mobile wallets. API clients should update the
// card's state to `OPEN` only after the cardholder confirms receipt of the card.
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardRenewResponseState string

const (
	CardRenewResponseStateClosed             CardRenewResponseState = "CLOSED"
	CardRenewResponseStateOpen               CardRenewResponseState = "OPEN"
	CardRenewResponseStatePaused             CardRenewResponseState = "PAUSED"
	CardRenewResponseStatePendingActivation  CardRenewResponseState = "PENDING_ACTIVATION"
	CardRenewResponseStatePendingFulfillment CardRenewResponseState = "PENDING_FULFILLMENT"
)

func (r CardRenewResponseState) IsKnown() bool {
	switch r {
	case CardRenewResponseStateClosed, CardRenewResponseStateOpen, CardRenewResponseStatePaused, CardRenewResponseStatePendingActivation, CardRenewResponseStatePendingFulfillment:
		return true
	}
	return false
}

// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
// functionality. _ `SINGLE_USE` - Card is closed upon first successful
// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
// VIRTUAL instead.
type CardRenewResponseType string

const (
	CardRenewResponseTypeMerchantLocked CardRenewResponseType = "MERCHANT_LOCKED"
	CardRenewResponseTypePhysical       CardRenewResponseType = "PHYSICAL"
	CardRenewResponseTypeSingleUse      CardRenewResponseType = "SINGLE_USE"
	CardRenewResponseTypeVirtual        CardRenewResponseType = "VIRTUAL"
	CardRenewResponseTypeUnlocked       CardRenewResponseType = "UNLOCKED"
	CardRenewResponseTypeDigitalWallet  CardRenewResponseType = "DIGITAL_WALLET"
)

func (r CardRenewResponseType) IsKnown() bool {
	switch r {
	case CardRenewResponseTypeMerchantLocked, CardRenewResponseTypePhysical, CardRenewResponseTypeSingleUse, CardRenewResponseTypeVirtual, CardRenewResponseTypeUnlocked, CardRenewResponseTypeDigitalWallet:
		return true
	}
	return false
}

// Card details with potentially PCI sensitive information for Enterprise customers
type CardSearchByPanResponse struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding CardSearchByPanResponseFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus CardSearchByPanResponsePinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
	// Spend limit duration
	SpendLimitDuration CardSearchByPanResponseSpendLimitDuration `json:"spend_limit_duration,required"`
	// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
	// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
	// they match card and account parameters). _ `PAUSED` - Card will decline
	// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
	// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
	// manufacturing and fulfillment. Cards in this state can accept authorizations for
	// e-commerce purchases, but not for "Card Present" purchases where the physical
	// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
	// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
	// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
	// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	// transactions or can be added to mobile wallets. API clients should update the
	// card's state to `OPEN` only after the cardholder confirms receipt of the card.
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardSearchByPanResponseState `json:"state,required"`
	// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
	// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
	// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
	// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
	// functionality. _ `SINGLE_USE` - Card is closed upon first successful
	// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
	// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
	// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
	// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	// VIRTUAL instead.
	Type CardSearchByPanResponseType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user's digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use.
	DigitalCardArtToken string `json:"digital_card_art_token"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card's locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// support@lithic.com for questions.
	Pan string `json:"pan"`
	// Indicates if there are offline PIN changes pending card interaction with an
	// offline PIN terminal. Possible commands are: CHANGE_PIN, UNBLOCK_PIN. Applicable
	// only to cards issued in markets supporting offline PINs.
	PendingCommands []string `json:"pending_commands"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string `json:"product_id"`
	// If the card is a replacement for another card, the globally unique identifier
	// for the card that was replaced.
	ReplacementFor string                      `json:"replacement_for,nullable"`
	JSON           cardSearchByPanResponseJSON `json:"-"`
}

// cardSearchByPanResponseJSON contains the JSON metadata for the struct
// [CardSearchByPanResponse]
type cardSearchByPanResponseJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	PinStatus           apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	CardholderCurrency  apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardSearchByPanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardSearchByPanResponseJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type CardSearchByPanResponseFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source. Funding source states: _ `ENABLED` - The funding
	// account is available to use for card creation and transactions. _ `PENDING` -
	// The funding account is still being verified e.g. bank micro-deposits
	// verification. \* `DELETED` - The founding account has been deleted.
	State CardSearchByPanResponseFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardSearchByPanResponseFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                             `json:"nickname"`
	JSON     cardSearchByPanResponseFundingJSON `json:"-"`
}

// cardSearchByPanResponseFundingJSON contains the JSON metadata for the struct
// [CardSearchByPanResponseFunding]
type cardSearchByPanResponseFundingJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	LastFour    apijson.Field
	State       apijson.Field
	Type        apijson.Field
	AccountName apijson.Field
	Nickname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardSearchByPanResponseFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardSearchByPanResponseFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type CardSearchByPanResponseFundingState string

const (
	CardSearchByPanResponseFundingStateDeleted CardSearchByPanResponseFundingState = "DELETED"
	CardSearchByPanResponseFundingStateEnabled CardSearchByPanResponseFundingState = "ENABLED"
	CardSearchByPanResponseFundingStatePending CardSearchByPanResponseFundingState = "PENDING"
)

func (r CardSearchByPanResponseFundingState) IsKnown() bool {
	switch r {
	case CardSearchByPanResponseFundingStateDeleted, CardSearchByPanResponseFundingStateEnabled, CardSearchByPanResponseFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type CardSearchByPanResponseFundingType string

const (
	CardSearchByPanResponseFundingTypeDepositoryChecking CardSearchByPanResponseFundingType = "DEPOSITORY_CHECKING"
	CardSearchByPanResponseFundingTypeDepositorySavings  CardSearchByPanResponseFundingType = "DEPOSITORY_SAVINGS"
)

func (r CardSearchByPanResponseFundingType) IsKnown() bool {
	switch r {
	case CardSearchByPanResponseFundingTypeDepositoryChecking, CardSearchByPanResponseFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type CardSearchByPanResponsePinStatus string

const (
	CardSearchByPanResponsePinStatusOk      CardSearchByPanResponsePinStatus = "OK"
	CardSearchByPanResponsePinStatusBlocked CardSearchByPanResponsePinStatus = "BLOCKED"
	CardSearchByPanResponsePinStatusNotSet  CardSearchByPanResponsePinStatus = "NOT_SET"
)

func (r CardSearchByPanResponsePinStatus) IsKnown() bool {
	switch r {
	case CardSearchByPanResponsePinStatusOk, CardSearchByPanResponsePinStatusBlocked, CardSearchByPanResponsePinStatusNotSet:
		return true
	}
	return false
}

// Spend limit duration
type CardSearchByPanResponseSpendLimitDuration string

const (
	CardSearchByPanResponseSpendLimitDurationAnnually    CardSearchByPanResponseSpendLimitDuration = "ANNUALLY"
	CardSearchByPanResponseSpendLimitDurationForever     CardSearchByPanResponseSpendLimitDuration = "FOREVER"
	CardSearchByPanResponseSpendLimitDurationMonthly     CardSearchByPanResponseSpendLimitDuration = "MONTHLY"
	CardSearchByPanResponseSpendLimitDurationTransaction CardSearchByPanResponseSpendLimitDuration = "TRANSACTION"
	CardSearchByPanResponseSpendLimitDurationDaily       CardSearchByPanResponseSpendLimitDuration = "DAILY"
)

func (r CardSearchByPanResponseSpendLimitDuration) IsKnown() bool {
	switch r {
	case CardSearchByPanResponseSpendLimitDurationAnnually, CardSearchByPanResponseSpendLimitDurationForever, CardSearchByPanResponseSpendLimitDurationMonthly, CardSearchByPanResponseSpendLimitDurationTransaction, CardSearchByPanResponseSpendLimitDurationDaily:
		return true
	}
	return false
}

// Card state values: _ `CLOSED` - Card will no longer approve authorizations.
// Closing a card cannot be undone. _ `OPEN` - Card will approve authorizations (if
// they match card and account parameters). _ `PAUSED` - Card will decline
// authorizations, but can be resumed at a later time. _ `PENDING_FULFILLMENT` -
// The initial state for cards of type `PHYSICAL`. The card is provisioned pending
// manufacturing and fulfillment. Cards in this state can accept authorizations for
// e-commerce purchases, but not for "Card Present" purchases where the physical
// card itself is present. \* `PENDING_ACTIVATION` - At regular intervals, cards of
// type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card production
// warehouse and updated to state `PENDING_ACTIVATION`. Similar to
// `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
// transactions or can be added to mobile wallets. API clients should update the
// card's state to `OPEN` only after the cardholder confirms receipt of the card.
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardSearchByPanResponseState string

const (
	CardSearchByPanResponseStateClosed             CardSearchByPanResponseState = "CLOSED"
	CardSearchByPanResponseStateOpen               CardSearchByPanResponseState = "OPEN"
	CardSearchByPanResponseStatePaused             CardSearchByPanResponseState = "PAUSED"
	CardSearchByPanResponseStatePendingActivation  CardSearchByPanResponseState = "PENDING_ACTIVATION"
	CardSearchByPanResponseStatePendingFulfillment CardSearchByPanResponseState = "PENDING_FULFILLMENT"
)

func (r CardSearchByPanResponseState) IsKnown() bool {
	switch r {
	case CardSearchByPanResponseStateClosed, CardSearchByPanResponseStateOpen, CardSearchByPanResponseStatePaused, CardSearchByPanResponseStatePendingActivation, CardSearchByPanResponseStatePendingFulfillment:
		return true
	}
	return false
}

// Card types: _ `VIRTUAL` - Card will authorize at any merchant and can be added
// to a digital wallet like Apple Pay or Google Pay (if the card program is digital
// wallet-enabled). _ `PHYSICAL` - Manufactured and sent to the cardholder. We
// offer white label branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe
// functionality. _ `SINGLE_USE` - Card is closed upon first successful
// authorization. _ `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first
// merchant that successfully authorizes the card. _ `UNLOCKED` - _[Deprecated]_
// Similar behavior to VIRTUAL cards, please use VIRTUAL instead. _
// `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
// VIRTUAL instead.
type CardSearchByPanResponseType string

const (
	CardSearchByPanResponseTypeMerchantLocked CardSearchByPanResponseType = "MERCHANT_LOCKED"
	CardSearchByPanResponseTypePhysical       CardSearchByPanResponseType = "PHYSICAL"
	CardSearchByPanResponseTypeSingleUse      CardSearchByPanResponseType = "SINGLE_USE"
	CardSearchByPanResponseTypeVirtual        CardSearchByPanResponseType = "VIRTUAL"
	CardSearchByPanResponseTypeUnlocked       CardSearchByPanResponseType = "UNLOCKED"
	CardSearchByPanResponseTypeDigitalWallet  CardSearchByPanResponseType = "DIGITAL_WALLET"
)

func (r CardSearchByPanResponseType) IsKnown() bool {
	switch r {
	case CardSearchByPanResponseTypeMerchantLocked, CardSearchByPanResponseTypePhysical, CardSearchByPanResponseTypeSingleUse, CardSearchByPanResponseTypeVirtual, CardSearchByPanResponseTypeUnlocked, CardSearchByPanResponseTypeDigitalWallet:
		return true
	}
	return false
}

type CardNewParams struct {
	// Card types:
	//
	//   - `VIRTUAL` - Card will authorize at any merchant and can be added to a digital
	//     wallet like Apple Pay or Google Pay (if the card program is digital
	//     wallet-enabled).
	//   - `PHYSICAL` - Manufactured and sent to the cardholder. We offer white label
	//     branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe functionality.
	//     Reach out at [lithic.com/contact](https://lithic.com/contact) for more
	//     information.
	//   - `SINGLE_USE` - Card is closed upon first successful authorization.
	//   - `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first merchant that
	//     successfully authorizes the card.
	//   - `UNLOCKED` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
	//     VIRTUAL instead.
	//   - `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please
	//     use VIRTUAL instead.
	Type param.Field[CardNewParamsType] `json:"type,required"`
	// Globally unique identifier for the account that the card will be associated
	// with. Required for programs enrolling users using the
	// [/account_holders endpoint](https://docs.lithic.com/docs/account-holders-kyc).
	// See [Managing Your Program](doc:managing-your-program) for more information.
	AccountToken param.Field[string] `json:"account_token" format:"uuid"`
	// For card programs with more than one BIN range. This must be configured with
	// Lithic before use. Identifies the card program/BIN range under which to create
	// the card. If omitted, will utilize the program's default `card_program_token`.
	// In Sandbox, use 00000000-0000-0000-1000-000000000000 and
	// 00000000-0000-0000-2000-000000000000 to test creating cards on specific card
	// programs.
	CardProgramToken param.Field[string]              `json:"card_program_token" format:"uuid"`
	Carrier          param.Field[shared.CarrierParam] `json:"carrier"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken param.Field[string] `json:"digital_card_art_token" format:"uuid"`
	// Two digit (MM) expiry month. If neither `exp_month` nor `exp_year` is provided,
	// an expiration date will be generated.
	ExpMonth param.Field[string] `json:"exp_month"`
	// Four digit (yyyy) expiry year. If neither `exp_month` nor `exp_year` is
	// provided, an expiration date will be generated.
	ExpYear param.Field[string] `json:"exp_year"`
	// Friendly name to identify the card.
	Memo param.Field[string] `json:"memo"`
	// Encrypted PIN block (in base64). Applies to cards of type `PHYSICAL` and
	// `VIRTUAL`. See
	// [Encrypted PIN Block](https://docs.lithic.com/docs/cards#encrypted-pin-block).
	Pin param.Field[string] `json:"pin"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID param.Field[string] `json:"product_id"`
	// Restricted field limited to select use cases. Lithic will reach out directly if
	// this field should be used. Globally unique identifier for the replacement card's
	// account. If this field is specified, `replacement_for` must also be specified.
	// If `replacement_for` is specified and this field is omitted, the replacement
	// card's account will be inferred from the card being replaced.
	ReplacementAccountToken param.Field[string] `json:"replacement_account_token" format:"uuid"`
	// Globally unique identifier for the card that this card will replace. If the card
	// type is `PHYSICAL` it will be replaced by a `PHYSICAL` card. If the card type is
	// `VIRTUAL` it will be replaced by a `VIRTUAL` card.
	ReplacementFor  param.Field[string]                      `json:"replacement_for" format:"uuid"`
	ShippingAddress param.Field[shared.ShippingAddressParam] `json:"shipping_address"`
	// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
	// options besides `STANDARD` require additional permissions.
	//
	//   - `STANDARD` - USPS regular mail or similar international option, with no
	//     tracking
	//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
	//     with tracking
	//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
	//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
	//   - `2_DAY` - FedEx 2-day shipping, with tracking
	//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
	//     tracking
	ShippingMethod param.Field[CardNewParamsShippingMethod] `json:"shipping_method"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined. Note
	// that a spend limit of 0 is effectively no limit, and should only be used to
	// reset or remove a prior limit. Only a limit of 1 or above will result in
	// declined transactions due to checks against the card limit.
	SpendLimit param.Field[int64] `json:"spend_limit"`
	// Spend limit duration values:
	//
	//   - `ANNUALLY` - Card will authorize transactions up to spend limit for the
	//     trailing year.
	//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
	//     of the card.
	//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
	//     trailing month. To support recurring monthly payments, which can occur on
	//     different day every month, the time window we consider for monthly velocity
	//     starts 6 days after the current calendar date one month prior.
	//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
	//     transaction is under the spend limit.
	SpendLimitDuration param.Field[SpendLimitDuration] `json:"spend_limit_duration"`
	// Card state values:
	//
	//   - `OPEN` - Card will approve authorizations (if they match card and account
	//     parameters).
	//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
	//     time.
	State param.Field[CardNewParamsState] `json:"state"`
}

func (r CardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Card types:
//
//   - `VIRTUAL` - Card will authorize at any merchant and can be added to a digital
//     wallet like Apple Pay or Google Pay (if the card program is digital
//     wallet-enabled).
//   - `PHYSICAL` - Manufactured and sent to the cardholder. We offer white label
//     branding, credit, ATM, PIN debit, chip/EMV, NFC and magstripe functionality.
//     Reach out at [lithic.com/contact](https://lithic.com/contact) for more
//     information.
//   - `SINGLE_USE` - Card is closed upon first successful authorization.
//   - `MERCHANT_LOCKED` - _[Deprecated]_ Card is locked to the first merchant that
//     successfully authorizes the card.
//   - `UNLOCKED` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please use
//     VIRTUAL instead.
//   - `DIGITAL_WALLET` - _[Deprecated]_ Similar behavior to VIRTUAL cards, please
//     use VIRTUAL instead.
type CardNewParamsType string

const (
	CardNewParamsTypeMerchantLocked CardNewParamsType = "MERCHANT_LOCKED"
	CardNewParamsTypePhysical       CardNewParamsType = "PHYSICAL"
	CardNewParamsTypeSingleUse      CardNewParamsType = "SINGLE_USE"
	CardNewParamsTypeVirtual        CardNewParamsType = "VIRTUAL"
	CardNewParamsTypeUnlocked       CardNewParamsType = "UNLOCKED"
	CardNewParamsTypeDigitalWallet  CardNewParamsType = "DIGITAL_WALLET"
)

func (r CardNewParamsType) IsKnown() bool {
	switch r {
	case CardNewParamsTypeMerchantLocked, CardNewParamsTypePhysical, CardNewParamsTypeSingleUse, CardNewParamsTypeVirtual, CardNewParamsTypeUnlocked, CardNewParamsTypeDigitalWallet:
		return true
	}
	return false
}

// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
// options besides `STANDARD` require additional permissions.
//
//   - `STANDARD` - USPS regular mail or similar international option, with no
//     tracking
//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
//     with tracking
//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
//   - `2_DAY` - FedEx 2-day shipping, with tracking
//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
//     tracking
type CardNewParamsShippingMethod string

const (
	CardNewParamsShippingMethod2Day                 CardNewParamsShippingMethod = "2_DAY"
	CardNewParamsShippingMethodExpedited            CardNewParamsShippingMethod = "EXPEDITED"
	CardNewParamsShippingMethodExpress              CardNewParamsShippingMethod = "EXPRESS"
	CardNewParamsShippingMethodPriority             CardNewParamsShippingMethod = "PRIORITY"
	CardNewParamsShippingMethodStandard             CardNewParamsShippingMethod = "STANDARD"
	CardNewParamsShippingMethodStandardWithTracking CardNewParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

func (r CardNewParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardNewParamsShippingMethod2Day, CardNewParamsShippingMethodExpedited, CardNewParamsShippingMethodExpress, CardNewParamsShippingMethodPriority, CardNewParamsShippingMethodStandard, CardNewParamsShippingMethodStandardWithTracking:
		return true
	}
	return false
}

// Card state values:
//
//   - `OPEN` - Card will approve authorizations (if they match card and account
//     parameters).
//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
//     time.
type CardNewParamsState string

const (
	CardNewParamsStateOpen   CardNewParamsState = "OPEN"
	CardNewParamsStatePaused CardNewParamsState = "PAUSED"
)

func (r CardNewParamsState) IsKnown() bool {
	switch r {
	case CardNewParamsStateOpen, CardNewParamsStatePaused:
		return true
	}
	return false
}

type CardUpdateParams struct {
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken param.Field[string] `json:"digital_card_art_token" format:"uuid"`
	// Friendly name to identify the card.
	Memo param.Field[string] `json:"memo"`
	// Encrypted PIN block (in base64). Only applies to cards of type `PHYSICAL` and
	// `VIRTUAL`. Changing PIN also resets PIN status to `OK`. See
	// [Encrypted PIN Block](https://docs.lithic.com/docs/cards#encrypted-pin-block).
	Pin param.Field[string] `json:"pin"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts). Can only be set to `OK` to unblock a card.
	PinStatus param.Field[CardUpdateParamsPinStatus] `json:"pin_status"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined. Note
	// that a spend limit of 0 is effectively no limit, and should only be used to
	// reset or remove a prior limit. Only a limit of 1 or above will result in
	// declined transactions due to checks against the card limit.
	SpendLimit param.Field[int64] `json:"spend_limit"`
	// Spend limit duration values:
	//
	//   - `ANNUALLY` - Card will authorize transactions up to spend limit for the
	//     trailing year.
	//   - `FOREVER` - Card will authorize only up to spend limit for the entire lifetime
	//     of the card.
	//   - `MONTHLY` - Card will authorize transactions up to spend limit for the
	//     trailing month. To support recurring monthly payments, which can occur on
	//     different day every month, the time window we consider for monthly velocity
	//     starts 6 days after the current calendar date one month prior.
	//   - `TRANSACTION` - Card will authorize multiple transactions if each individual
	//     transaction is under the spend limit.
	SpendLimitDuration param.Field[SpendLimitDuration] `json:"spend_limit_duration"`
	// Card state values:
	//
	//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
	//     be undone.
	//   - `OPEN` - Card will approve authorizations (if they match card and account
	//     parameters).
	//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
	//     time.
	State param.Field[CardUpdateParamsState] `json:"state"`
}

func (r CardUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts). Can only be set to `OK` to unblock a card.
type CardUpdateParamsPinStatus string

const (
	CardUpdateParamsPinStatusOk CardUpdateParamsPinStatus = "OK"
)

func (r CardUpdateParamsPinStatus) IsKnown() bool {
	switch r {
	case CardUpdateParamsPinStatusOk:
		return true
	}
	return false
}

// Card state values:
//
//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
//     be undone.
//   - `OPEN` - Card will approve authorizations (if they match card and account
//     parameters).
//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
//     time.
type CardUpdateParamsState string

const (
	CardUpdateParamsStateClosed CardUpdateParamsState = "CLOSED"
	CardUpdateParamsStateOpen   CardUpdateParamsState = "OPEN"
	CardUpdateParamsStatePaused CardUpdateParamsState = "PAUSED"
)

func (r CardUpdateParamsState) IsKnown() bool {
	switch r {
	case CardUpdateParamsStateClosed, CardUpdateParamsStateOpen, CardUpdateParamsStatePaused:
		return true
	}
	return false
}

type CardListParams struct {
	// Returns cards associated with the specified account.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Returns cards with the specified state.
	State param.Field[CardListParamsState] `query:"state"`
}

// URLQuery serializes [CardListParams]'s query parameters as `url.Values`.
func (r CardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Returns cards with the specified state.
type CardListParamsState string

const (
	CardListParamsStateClosed             CardListParamsState = "CLOSED"
	CardListParamsStateOpen               CardListParamsState = "OPEN"
	CardListParamsStatePaused             CardListParamsState = "PAUSED"
	CardListParamsStatePendingActivation  CardListParamsState = "PENDING_ACTIVATION"
	CardListParamsStatePendingFulfillment CardListParamsState = "PENDING_FULFILLMENT"
)

func (r CardListParamsState) IsKnown() bool {
	switch r {
	case CardListParamsStateClosed, CardListParamsStateOpen, CardListParamsStatePaused, CardListParamsStatePendingActivation, CardListParamsStatePendingFulfillment:
		return true
	}
	return false
}

type CardConvertPhysicalParams struct {
	// The shipping address this card will be sent to.
	ShippingAddress param.Field[shared.ShippingAddressParam] `json:"shipping_address,required"`
	// If omitted, the previous carrier will be used.
	Carrier param.Field[shared.CarrierParam] `json:"carrier"`
	// Specifies the configuration (e.g. physical card art) that the card should be
	// manufactured with, and only applies to cards of type `PHYSICAL`. This must be
	// configured with Lithic before use.
	ProductID param.Field[string] `json:"product_id"`
	// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
	// options besides `STANDARD` require additional permissions.
	//
	//   - `STANDARD` - USPS regular mail or similar international option, with no
	//     tracking
	//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
	//     with tracking
	//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
	//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
	//   - `2_DAY` - FedEx 2-day shipping, with tracking
	//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
	//     tracking
	ShippingMethod param.Field[CardConvertPhysicalParamsShippingMethod] `json:"shipping_method"`
}

func (r CardConvertPhysicalParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
// options besides `STANDARD` require additional permissions.
//
//   - `STANDARD` - USPS regular mail or similar international option, with no
//     tracking
//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
//     with tracking
//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
//   - `2_DAY` - FedEx 2-day shipping, with tracking
//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
//     tracking
type CardConvertPhysicalParamsShippingMethod string

const (
	CardConvertPhysicalParamsShippingMethod2Day                 CardConvertPhysicalParamsShippingMethod = "2_DAY"
	CardConvertPhysicalParamsShippingMethodExpedited            CardConvertPhysicalParamsShippingMethod = "EXPEDITED"
	CardConvertPhysicalParamsShippingMethodExpress              CardConvertPhysicalParamsShippingMethod = "EXPRESS"
	CardConvertPhysicalParamsShippingMethodPriority             CardConvertPhysicalParamsShippingMethod = "PRIORITY"
	CardConvertPhysicalParamsShippingMethodStandard             CardConvertPhysicalParamsShippingMethod = "STANDARD"
	CardConvertPhysicalParamsShippingMethodStandardWithTracking CardConvertPhysicalParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

func (r CardConvertPhysicalParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardConvertPhysicalParamsShippingMethod2Day, CardConvertPhysicalParamsShippingMethodExpedited, CardConvertPhysicalParamsShippingMethodExpress, CardConvertPhysicalParamsShippingMethodPriority, CardConvertPhysicalParamsShippingMethodStandard, CardConvertPhysicalParamsShippingMethodStandardWithTracking:
		return true
	}
	return false
}

type CardEmbedParams struct {
	// A base64 encoded JSON string of an EmbedRequest to specify which card to load.
	EmbedRequest param.Field[string] `query:"embed_request,required"`
	// SHA256 HMAC of the embed_request JSON string with base64 digest.
	Hmac param.Field[string] `query:"hmac,required"`
}

// URLQuery serializes [CardEmbedParams]'s query parameters as `url.Values`.
func (r CardEmbedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CardGetEmbedHTMLParams struct {
	// Globally unique identifier for the card to be displayed.
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// A publicly available URI, so the white-labeled card element can be styled with
	// the client's branding.
	Css param.Field[string] `json:"css"`
	// An RFC 3339 timestamp for when the request should expire. UTC time zone.
	//
	// If no timezone is specified, UTC will be used. If payload does not contain an
	// expiration, the request will never expire.
	//
	// Using an `expiration` reduces the risk of a
	// [replay attack](https://en.wikipedia.org/wiki/Replay_attack). Without supplying
	// the `expiration`, in the event that a malicious user gets a copy of your request
	// in transit, they will be able to obtain the response data indefinitely.
	Expiration param.Field[time.Time] `json:"expiration" format:"date-time"`
	// Required if you want to post the element clicked to the parent iframe.
	//
	// If you supply this param, you can also capture click events in the parent iframe
	// by adding an event listener.
	TargetOrigin param.Field[string] `json:"target_origin"`
}

func (r CardGetEmbedHTMLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardGetEmbedURLParams struct {
	// Globally unique identifier for the card to be displayed.
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// A publicly available URI, so the white-labeled card element can be styled with
	// the client's branding.
	Css param.Field[string] `json:"css"`
	// An RFC 3339 timestamp for when the request should expire. UTC time zone.
	//
	// If no timezone is specified, UTC will be used. If payload does not contain an
	// expiration, the request will never expire.
	//
	// Using an `expiration` reduces the risk of a
	// [replay attack](https://en.wikipedia.org/wiki/Replay_attack). Without supplying
	// the `expiration`, in the event that a malicious user gets a copy of your request
	// in transit, they will be able to obtain the response data indefinitely.
	Expiration param.Field[time.Time] `json:"expiration" format:"date-time"`
	// Required if you want to post the element clicked to the parent iframe.
	//
	// If you supply this param, you can also capture click events in the parent iframe
	// by adding an event listener.
	TargetOrigin param.Field[string] `json:"target_origin"`
}

func (r CardGetEmbedURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardProvisionParams struct {
	// Only applicable if `digital_wallet` is `APPLE_PAY`. Omit to receive only
	// `activationData` in the response. Apple's public leaf certificate. Base64
	// encoded in PEM format with headers `(-----BEGIN CERTIFICATE-----)` and trailers
	// omitted. Provided by the device's wallet.
	Certificate param.Field[string] `json:"certificate" format:"byte"`
	// Only applicable if `digital_wallet` is `GOOGLE_PAY` or `SAMSUNG_PAY` and the
	// card is on the Visa network. Stable device identification set by the wallet
	// provider.
	ClientDeviceID param.Field[string] `json:"client_device_id"`
	// Only applicable if `digital_wallet` is `GOOGLE_PAY` or `SAMSUNG_PAY` and the
	// card is on the Visa network. Consumer ID that identifies the wallet account
	// holder entity.
	ClientWalletAccountID param.Field[string] `json:"client_wallet_account_id"`
	// Name of digital wallet provider.
	DigitalWallet param.Field[CardProvisionParamsDigitalWallet] `json:"digital_wallet"`
	// Only applicable if `digital_wallet` is `APPLE_PAY`. Omit to receive only
	// `activationData` in the response. Base64 cryptographic nonce provided by the
	// device's wallet.
	Nonce param.Field[string] `json:"nonce" format:"byte"`
	// Only applicable if `digital_wallet` is `APPLE_PAY`. Omit to receive only
	// `activationData` in the response. Base64 cryptographic nonce provided by the
	// device's wallet.
	NonceSignature param.Field[string] `json:"nonce_signature" format:"byte"`
}

func (r CardProvisionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Name of digital wallet provider.
type CardProvisionParamsDigitalWallet string

const (
	CardProvisionParamsDigitalWalletApplePay   CardProvisionParamsDigitalWallet = "APPLE_PAY"
	CardProvisionParamsDigitalWalletGooglePay  CardProvisionParamsDigitalWallet = "GOOGLE_PAY"
	CardProvisionParamsDigitalWalletSamsungPay CardProvisionParamsDigitalWallet = "SAMSUNG_PAY"
)

func (r CardProvisionParamsDigitalWallet) IsKnown() bool {
	switch r {
	case CardProvisionParamsDigitalWalletApplePay, CardProvisionParamsDigitalWalletGooglePay, CardProvisionParamsDigitalWalletSamsungPay:
		return true
	}
	return false
}

type CardReissueParams struct {
	// If omitted, the previous carrier will be used.
	Carrier param.Field[shared.CarrierParam] `json:"carrier"`
	// Specifies the configuration (e.g. physical card art) that the card should be
	// manufactured with, and only applies to cards of type `PHYSICAL`. This must be
	// configured with Lithic before use.
	ProductID param.Field[string] `json:"product_id"`
	// If omitted, the previous shipping address will be used.
	ShippingAddress param.Field[shared.ShippingAddressParam] `json:"shipping_address"`
	// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
	// options besides `STANDARD` require additional permissions.
	//
	//   - `STANDARD` - USPS regular mail or similar international option, with no
	//     tracking
	//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
	//     with tracking
	//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
	//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
	//   - `2_DAY` - FedEx 2-day shipping, with tracking
	//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
	//     tracking
	ShippingMethod param.Field[CardReissueParamsShippingMethod] `json:"shipping_method"`
}

func (r CardReissueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
// options besides `STANDARD` require additional permissions.
//
//   - `STANDARD` - USPS regular mail or similar international option, with no
//     tracking
//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
//     with tracking
//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
//   - `2_DAY` - FedEx 2-day shipping, with tracking
//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
//     tracking
type CardReissueParamsShippingMethod string

const (
	CardReissueParamsShippingMethod2Day                 CardReissueParamsShippingMethod = "2_DAY"
	CardReissueParamsShippingMethodExpedited            CardReissueParamsShippingMethod = "EXPEDITED"
	CardReissueParamsShippingMethodExpress              CardReissueParamsShippingMethod = "EXPRESS"
	CardReissueParamsShippingMethodPriority             CardReissueParamsShippingMethod = "PRIORITY"
	CardReissueParamsShippingMethodStandard             CardReissueParamsShippingMethod = "STANDARD"
	CardReissueParamsShippingMethodStandardWithTracking CardReissueParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

func (r CardReissueParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardReissueParamsShippingMethod2Day, CardReissueParamsShippingMethodExpedited, CardReissueParamsShippingMethodExpress, CardReissueParamsShippingMethodPriority, CardReissueParamsShippingMethodStandard, CardReissueParamsShippingMethodStandardWithTracking:
		return true
	}
	return false
}

type CardRenewParams struct {
	// The shipping address this card will be sent to.
	ShippingAddress param.Field[shared.ShippingAddressParam] `json:"shipping_address,required"`
	// If omitted, the previous carrier will be used.
	Carrier param.Field[shared.CarrierParam] `json:"carrier"`
	// Two digit (MM) expiry month. If neither `exp_month` nor `exp_year` is provided,
	// an expiration date six years in the future will be generated.
	ExpMonth param.Field[string] `json:"exp_month"`
	// Four digit (yyyy) expiry year. If neither `exp_month` nor `exp_year` is
	// provided, an expiration date six years in the future will be generated.
	ExpYear param.Field[string] `json:"exp_year"`
	// Specifies the configuration (e.g. physical card art) that the card should be
	// manufactured with, and only applies to cards of type `PHYSICAL`. This must be
	// configured with Lithic before use.
	ProductID param.Field[string] `json:"product_id"`
	// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
	// options besides `STANDARD` require additional permissions.
	//
	//   - `STANDARD` - USPS regular mail or similar international option, with no
	//     tracking
	//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
	//     with tracking
	//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
	//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
	//   - `2_DAY` - FedEx 2-day shipping, with tracking
	//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
	//     tracking
	ShippingMethod param.Field[CardRenewParamsShippingMethod] `json:"shipping_method"`
}

func (r CardRenewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
// options besides `STANDARD` require additional permissions.
//
//   - `STANDARD` - USPS regular mail or similar international option, with no
//     tracking
//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
//     with tracking
//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
//   - `EXPRESS` - FedEx Express, 3-day shipping, with tracking
//   - `2_DAY` - FedEx 2-day shipping, with tracking
//   - `EXPEDITED` - FedEx Standard Overnight or similar international option, with
//     tracking
type CardRenewParamsShippingMethod string

const (
	CardRenewParamsShippingMethod2Day                 CardRenewParamsShippingMethod = "2_DAY"
	CardRenewParamsShippingMethodExpedited            CardRenewParamsShippingMethod = "EXPEDITED"
	CardRenewParamsShippingMethodExpress              CardRenewParamsShippingMethod = "EXPRESS"
	CardRenewParamsShippingMethodPriority             CardRenewParamsShippingMethod = "PRIORITY"
	CardRenewParamsShippingMethodStandard             CardRenewParamsShippingMethod = "STANDARD"
	CardRenewParamsShippingMethodStandardWithTracking CardRenewParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

func (r CardRenewParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardRenewParamsShippingMethod2Day, CardRenewParamsShippingMethodExpedited, CardRenewParamsShippingMethodExpress, CardRenewParamsShippingMethodPriority, CardRenewParamsShippingMethodStandard, CardRenewParamsShippingMethodStandardWithTracking:
		return true
	}
	return false
}

type CardSearchByPanParams struct {
	// The PAN for the card being retrieved.
	Pan param.Field[string] `json:"pan,required"`
}

func (r CardSearchByPanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
