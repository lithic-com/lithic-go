// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/lithic-com/lithic-go/shared"
	"github.com/tidwall/gjson"
)

// CardService contains methods and other services that help with interacting with
// the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardService] method instead.
type CardService struct {
	Options               []option.RequestOption
	Balances              *CardBalanceService
	FinancialTransactions *CardFinancialTransactionService
}

// NewCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardService(opts ...option.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	r.Balances = NewCardBalanceService(opts...)
	r.FinancialTransactions = NewCardFinancialTransactionService(opts...)
	return
}

// Create a new virtual or physical card. Parameters `shipping_address` and
// `product_id` only apply to physical cards.
func (r *CardService) New(ctx context.Context, params CardNewParams, opts ...option.RequestOption) (res *Card, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%s", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get card configuration such as spend limit and state.
func (r *CardService) Get(ctx context.Context, cardToken string, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *CardService) Update(ctx context.Context, cardToken string, body CardUpdateParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List cards.
func (r *CardService) List(ctx context.Context, query CardListParams, opts ...option.RequestOption) (res *pagination.CursorPage[NonPCICard], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *CardService) ListAutoPaging(ctx context.Context, query CardListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[NonPCICard] {
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
func (r *CardService) ConvertPhysical(ctx context.Context, cardToken string, body CardConvertPhysicalParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	path := "v1/embed/card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Allow your cardholders to directly add payment cards to the device's digital
// wallet (e.g. Apple Pay) with one touch from your app.
//
// This requires some additional setup and configuration. Please
// [Contact Us](https://lithic.com/contact) or your Customer Success representative
// for more information.
func (r *CardService) Provision(ctx context.Context, cardToken string, body CardProvisionParams, opts ...option.RequestOption) (res *CardProvisionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
// applies to cards of type `PHYSICAL`. A card can be reissued or renewed a total
// of 8 times.
func (r *CardService) Reissue(ctx context.Context, cardToken string, body CardReissueParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
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
// `PHYSICAL` card can be reissued or renewed a total of 8 times. For `VIRTUAL`,
// the card will retain the same card token and PAN and receive an updated expiry
// and CVC2 code. `product_id`, `shipping_method`, `shipping_address`, `carrier`
// are only relevant for renewing `PHYSICAL` cards.
func (r *CardService) Renew(ctx context.Context, cardToken string, body CardRenewParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
func (r *CardService) SearchByPan(ctx context.Context, body CardSearchByPanParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/cards/search_by_pan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allow your cardholders to directly add payment cards to the device's digital
// wallet from a browser on the web.
//
// This requires some additional setup and configuration. Please
// [Contact Us](https://lithic.com/contact) or your Customer Success representative
// for more information.
func (r *CardService) WebProvision(ctx context.Context, cardToken string, body CardWebProvisionParams, opts ...option.RequestOption) (res *CardWebProvisionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardToken == "" {
		err = errors.New("missing required card_token parameter")
		return
	}
	path := fmt.Sprintf("v1/cards/%s/web_provision", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Card details with potentially PCI sensitive information for Enterprise customers
type Card struct {
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// support@lithic.com for questions.
	Pan  string   `json:"pan"`
	JSON cardJSON `json:"-"`
	NonPCICard
}

// cardJSON contains the JSON metadata for the struct [Card]
type cardJSON struct {
	Cvv         apijson.Field
	Pan         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardJSON) RawJSON() string {
	return r.raw
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

// Card details without PCI information
type NonPCICard struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Deprecated: Funding account for the card.
	Funding NonPCICardFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
	// attempts).
	PinStatus NonPCICardPinStatus `json:"pin_status,required"`
	// Amount (in cents) to limit approved authorizations (e.g. 100000 would be a
	// $1,000 limit). Transaction requests above the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
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
	SpendLimitDuration SpendLimitDuration `json:"spend_limit_duration,required"`
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
	State NonPCICardState `json:"state,required"`
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
	Type NonPCICardType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card. This
	// field is deprecated and will no longer be populated in the `Card` object. The
	// key will be removed from the schema in a future release. Use the `/auth_rules`
	// endpoints to fetch Auth Rule information instead.
	//
	// Deprecated: deprecated
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// Globally unique identifier for the bulk order associated with this card. Only
	// applicable to physical cards that are part of a bulk shipment
	BulkOrderToken string `json:"bulk_order_token,nullable" format:"uuid"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// Additional context or information related to the card.
	Comment string `json:"comment"`
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
	// Globally unique identifier for the card's network program. Null if the card is
	// not associated with a network program. Currently applicable to Visa cards
	// participating in Account Level Management only
	NetworkProgramToken string `json:"network_program_token,nullable"`
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
	ReplacementFor string `json:"replacement_for,nullable"`
	// Card state substatus values: _ `LOST` - The physical card is no longer in the
	// cardholder's possession due to being lost or never received by the cardholder. _
	// `COMPROMISED` - Card information has been exposed, potentially leading to
	// unauthorized access. This may involve physical card theft, cloning, or online
	// data breaches. _ `DAMAGED` - The physical card is not functioning properly, such
	// as having chip failures or a demagnetized magnetic stripe. _
	// `END_USER_REQUEST` - The cardholder requested the closure of the card for
	// reasons unrelated to fraud or damage, such as switching to a different product
	// or closing the account. _ `ISSUER_REQUEST` - The issuer closed the card for
	// reasons unrelated to fraud or damage, such as account inactivity, product or
	// policy changes, or technology upgrades. _ `NOT_ACTIVE` - The card hasn’t had any
	// transaction activity for a specified period, applicable to statuses like
	// `PAUSED` or `CLOSED`. _ `SUSPICIOUS_ACTIVITY` - The card has one or more
	// suspicious transactions or activities that require review. This can involve
	// prompting the cardholder to confirm legitimate use or report confirmed fraud. _
	// `INTERNAL_REVIEW` - The card is temporarily paused pending further internal
	// review. _ `EXPIRED` - The card has expired and has been closed without being
	// reissued. _ `UNDELIVERABLE` - The card cannot be delivered to the cardholder and
	// has been returned. \* `OTHER` - The reason for the status does not fall into any
	// of the above categories. A comment can be provided to specify the reason.
	Substatus NonPCICardSubstatus `json:"substatus"`
	JSON      nonPCICardJSON      `json:"-"`
}

// nonPCICardJSON contains the JSON metadata for the struct [NonPCICard]
type nonPCICardJSON struct {
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
	BulkOrderToken      apijson.Field
	CardholderCurrency  apijson.Field
	Comment             apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	NetworkProgramToken apijson.Field
	PendingCommands     apijson.Field
	ProductID           apijson.Field
	ReplacementFor      apijson.Field
	Substatus           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NonPCICard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nonPCICardJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Funding account for the card.
type NonPCICardFunding struct {
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
	State NonPCICardFundingState `json:"state,required"`
	// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
	// `DEPOSITORY_SAVINGS` - Bank savings account.
	Type NonPCICardFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string                `json:"nickname"`
	JSON     nonPCICardFundingJSON `json:"-"`
}

// nonPCICardFundingJSON contains the JSON metadata for the struct
// [NonPCICardFunding]
type nonPCICardFundingJSON struct {
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

func (r *NonPCICardFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nonPCICardFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source. Funding source states: _ `ENABLED` - The funding
// account is available to use for card creation and transactions. _ `PENDING` -
// The funding account is still being verified e.g. bank micro-deposits
// verification. \* `DELETED` - The founding account has been deleted.
type NonPCICardFundingState string

const (
	NonPCICardFundingStateDeleted NonPCICardFundingState = "DELETED"
	NonPCICardFundingStateEnabled NonPCICardFundingState = "ENABLED"
	NonPCICardFundingStatePending NonPCICardFundingState = "PENDING"
)

func (r NonPCICardFundingState) IsKnown() bool {
	switch r {
	case NonPCICardFundingStateDeleted, NonPCICardFundingStateEnabled, NonPCICardFundingStatePending:
		return true
	}
	return false
}

// Types of funding source: _ `DEPOSITORY_CHECKING` - Bank checking account. _
// `DEPOSITORY_SAVINGS` - Bank savings account.
type NonPCICardFundingType string

const (
	NonPCICardFundingTypeDepositoryChecking NonPCICardFundingType = "DEPOSITORY_CHECKING"
	NonPCICardFundingTypeDepositorySavings  NonPCICardFundingType = "DEPOSITORY_SAVINGS"
)

func (r NonPCICardFundingType) IsKnown() bool {
	switch r {
	case NonPCICardFundingTypeDepositoryChecking, NonPCICardFundingTypeDepositorySavings:
		return true
	}
	return false
}

// Indicates if a card is blocked due a PIN status issue (e.g. excessive incorrect
// attempts).
type NonPCICardPinStatus string

const (
	NonPCICardPinStatusOk      NonPCICardPinStatus = "OK"
	NonPCICardPinStatusBlocked NonPCICardPinStatus = "BLOCKED"
	NonPCICardPinStatusNotSet  NonPCICardPinStatus = "NOT_SET"
)

func (r NonPCICardPinStatus) IsKnown() bool {
	switch r {
	case NonPCICardPinStatusOk, NonPCICardPinStatusBlocked, NonPCICardPinStatusNotSet:
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
type NonPCICardState string

const (
	NonPCICardStateClosed             NonPCICardState = "CLOSED"
	NonPCICardStateOpen               NonPCICardState = "OPEN"
	NonPCICardStatePaused             NonPCICardState = "PAUSED"
	NonPCICardStatePendingActivation  NonPCICardState = "PENDING_ACTIVATION"
	NonPCICardStatePendingFulfillment NonPCICardState = "PENDING_FULFILLMENT"
)

func (r NonPCICardState) IsKnown() bool {
	switch r {
	case NonPCICardStateClosed, NonPCICardStateOpen, NonPCICardStatePaused, NonPCICardStatePendingActivation, NonPCICardStatePendingFulfillment:
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
type NonPCICardType string

const (
	NonPCICardTypeMerchantLocked NonPCICardType = "MERCHANT_LOCKED"
	NonPCICardTypePhysical       NonPCICardType = "PHYSICAL"
	NonPCICardTypeSingleUse      NonPCICardType = "SINGLE_USE"
	NonPCICardTypeVirtual        NonPCICardType = "VIRTUAL"
	NonPCICardTypeUnlocked       NonPCICardType = "UNLOCKED"
	NonPCICardTypeDigitalWallet  NonPCICardType = "DIGITAL_WALLET"
)

func (r NonPCICardType) IsKnown() bool {
	switch r {
	case NonPCICardTypeMerchantLocked, NonPCICardTypePhysical, NonPCICardTypeSingleUse, NonPCICardTypeVirtual, NonPCICardTypeUnlocked, NonPCICardTypeDigitalWallet:
		return true
	}
	return false
}

// Card state substatus values: _ `LOST` - The physical card is no longer in the
// cardholder's possession due to being lost or never received by the cardholder. _
// `COMPROMISED` - Card information has been exposed, potentially leading to
// unauthorized access. This may involve physical card theft, cloning, or online
// data breaches. _ `DAMAGED` - The physical card is not functioning properly, such
// as having chip failures or a demagnetized magnetic stripe. _
// `END_USER_REQUEST` - The cardholder requested the closure of the card for
// reasons unrelated to fraud or damage, such as switching to a different product
// or closing the account. _ `ISSUER_REQUEST` - The issuer closed the card for
// reasons unrelated to fraud or damage, such as account inactivity, product or
// policy changes, or technology upgrades. _ `NOT_ACTIVE` - The card hasn’t had any
// transaction activity for a specified period, applicable to statuses like
// `PAUSED` or `CLOSED`. _ `SUSPICIOUS_ACTIVITY` - The card has one or more
// suspicious transactions or activities that require review. This can involve
// prompting the cardholder to confirm legitimate use or report confirmed fraud. _
// `INTERNAL_REVIEW` - The card is temporarily paused pending further internal
// review. _ `EXPIRED` - The card has expired and has been closed without being
// reissued. _ `UNDELIVERABLE` - The card cannot be delivered to the cardholder and
// has been returned. \* `OTHER` - The reason for the status does not fall into any
// of the above categories. A comment can be provided to specify the reason.
type NonPCICardSubstatus string

const (
	NonPCICardSubstatusLost               NonPCICardSubstatus = "LOST"
	NonPCICardSubstatusCompromised        NonPCICardSubstatus = "COMPROMISED"
	NonPCICardSubstatusDamaged            NonPCICardSubstatus = "DAMAGED"
	NonPCICardSubstatusEndUserRequest     NonPCICardSubstatus = "END_USER_REQUEST"
	NonPCICardSubstatusIssuerRequest      NonPCICardSubstatus = "ISSUER_REQUEST"
	NonPCICardSubstatusNotActive          NonPCICardSubstatus = "NOT_ACTIVE"
	NonPCICardSubstatusSuspiciousActivity NonPCICardSubstatus = "SUSPICIOUS_ACTIVITY"
	NonPCICardSubstatusInternalReview     NonPCICardSubstatus = "INTERNAL_REVIEW"
	NonPCICardSubstatusExpired            NonPCICardSubstatus = "EXPIRED"
	NonPCICardSubstatusUndeliverable      NonPCICardSubstatus = "UNDELIVERABLE"
	NonPCICardSubstatusOther              NonPCICardSubstatus = "OTHER"
)

func (r NonPCICardSubstatus) IsKnown() bool {
	switch r {
	case NonPCICardSubstatusLost, NonPCICardSubstatusCompromised, NonPCICardSubstatusDamaged, NonPCICardSubstatusEndUserRequest, NonPCICardSubstatusIssuerRequest, NonPCICardSubstatusNotActive, NonPCICardSubstatusSuspiciousActivity, NonPCICardSubstatusInternalReview, NonPCICardSubstatusExpired, NonPCICardSubstatusUndeliverable, NonPCICardSubstatusOther:
		return true
	}
	return false
}

// Object containing the fields required to add a card to Apple Pay. Applies only
// to Apple Pay wallet.
type ProvisionResponse struct {
	ActivationData     string                `json:"activationData"`
	EncryptedData      string                `json:"encryptedData"`
	EphemeralPublicKey string                `json:"ephemeralPublicKey"`
	JSON               provisionResponseJSON `json:"-"`
}

// provisionResponseJSON contains the JSON metadata for the struct
// [ProvisionResponse]
type provisionResponseJSON struct {
	ActivationData     apijson.Field
	EncryptedData      apijson.Field
	EphemeralPublicKey apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProvisionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r provisionResponseJSON) RawJSON() string {
	return r.raw
}

func (r ProvisionResponse) ImplementsCardProvisionResponseProvisioningPayloadUnion() {}

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

type CardProvisionResponse struct {
	// Base64 encoded JSON payload representing a payment card that can be passed to a
	// device's digital wallet. Applies to Google and Samsung Pay wallets.
	ProvisioningPayload CardProvisionResponseProvisioningPayloadUnion `json:"provisioning_payload"`
	JSON                cardProvisionResponseJSON                     `json:"-"`
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

// Base64 encoded JSON payload representing a payment card that can be passed to a
// device's digital wallet. Applies to Google and Samsung Pay wallets.
//
// Union satisfied by [shared.UnionString] or [ProvisionResponse].
type CardProvisionResponseProvisioningPayloadUnion interface {
	ImplementsCardProvisionResponseProvisioningPayloadUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CardProvisionResponseProvisioningPayloadUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProvisionResponse{}),
		},
	)
}

type CardWebProvisionResponse struct {
	// A base64 encoded and encrypted payload representing card data for the Google Pay
	// UWPP FPAN flow.
	GoogleOpc string `json:"google_opc"`
	// This field can have the runtime type of
	// [CardWebProvisionResponseAppleWebPushProvisioningResponseJws].
	Jws interface{} `json:"jws"`
	// A unique identifier for the JWS object.
	State string `json:"state"`
	// A base64 encoded and encrypted payload representing card data for the Google Pay
	// UWPP tokenization flow.
	TspOpc string                       `json:"tsp_opc"`
	JSON   cardWebProvisionResponseJSON `json:"-"`
	union  CardWebProvisionResponseUnion
}

// cardWebProvisionResponseJSON contains the JSON metadata for the struct
// [CardWebProvisionResponse]
type cardWebProvisionResponseJSON struct {
	GoogleOpc   apijson.Field
	Jws         apijson.Field
	State       apijson.Field
	TspOpc      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r cardWebProvisionResponseJSON) RawJSON() string {
	return r.raw
}

func (r *CardWebProvisionResponse) UnmarshalJSON(data []byte) (err error) {
	*r = CardWebProvisionResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CardWebProvisionResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CardWebProvisionResponseAppleWebPushProvisioningResponse],
// [CardWebProvisionResponseGoogleWebPushProvisioningResponse].
func (r CardWebProvisionResponse) AsUnion() CardWebProvisionResponseUnion {
	return r.union
}

// Union satisfied by [CardWebProvisionResponseAppleWebPushProvisioningResponse] or
// [CardWebProvisionResponseGoogleWebPushProvisioningResponse].
type CardWebProvisionResponseUnion interface {
	implementsCardWebProvisionResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CardWebProvisionResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardWebProvisionResponseAppleWebPushProvisioningResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CardWebProvisionResponseGoogleWebPushProvisioningResponse{}),
		},
	)
}

type CardWebProvisionResponseAppleWebPushProvisioningResponse struct {
	// JWS object required for handoff to Apple's script.
	Jws CardWebProvisionResponseAppleWebPushProvisioningResponseJws `json:"jws,required"`
	// A unique identifier for the JWS object.
	State string                                                       `json:"state,required"`
	JSON  cardWebProvisionResponseAppleWebPushProvisioningResponseJSON `json:"-"`
}

// cardWebProvisionResponseAppleWebPushProvisioningResponseJSON contains the JSON
// metadata for the struct
// [CardWebProvisionResponseAppleWebPushProvisioningResponse]
type cardWebProvisionResponseAppleWebPushProvisioningResponseJSON struct {
	Jws         apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardWebProvisionResponseAppleWebPushProvisioningResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardWebProvisionResponseAppleWebPushProvisioningResponseJSON) RawJSON() string {
	return r.raw
}

func (r CardWebProvisionResponseAppleWebPushProvisioningResponse) implementsCardWebProvisionResponse() {
}

// JWS object required for handoff to Apple's script.
type CardWebProvisionResponseAppleWebPushProvisioningResponseJws struct {
	// JWS unprotected headers containing header parameters that aren't
	// integrity-protected by the JWS signature.
	Header CardWebProvisionResponseAppleWebPushProvisioningResponseJwsHeader `json:"header"`
	// Base64url encoded JSON object containing the provisioning payload.
	Payload string `json:"payload"`
	// Base64url encoded JWS protected headers containing the header parameters that
	// are integrity-protected by the JWS signature.
	Protected string `json:"protected"`
	// Base64url encoded signature of the JWS object.
	Signature string                                                          `json:"signature"`
	JSON      cardWebProvisionResponseAppleWebPushProvisioningResponseJwsJSON `json:"-"`
}

// cardWebProvisionResponseAppleWebPushProvisioningResponseJwsJSON contains the
// JSON metadata for the struct
// [CardWebProvisionResponseAppleWebPushProvisioningResponseJws]
type cardWebProvisionResponseAppleWebPushProvisioningResponseJwsJSON struct {
	Header      apijson.Field
	Payload     apijson.Field
	Protected   apijson.Field
	Signature   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardWebProvisionResponseAppleWebPushProvisioningResponseJws) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardWebProvisionResponseAppleWebPushProvisioningResponseJwsJSON) RawJSON() string {
	return r.raw
}

// JWS unprotected headers containing header parameters that aren't
// integrity-protected by the JWS signature.
type CardWebProvisionResponseAppleWebPushProvisioningResponseJwsHeader struct {
	// The ID for the JWS Public Key of the key pair used to generate the signature.
	Kid  string                                                                `json:"kid"`
	JSON cardWebProvisionResponseAppleWebPushProvisioningResponseJwsHeaderJSON `json:"-"`
}

// cardWebProvisionResponseAppleWebPushProvisioningResponseJwsHeaderJSON contains
// the JSON metadata for the struct
// [CardWebProvisionResponseAppleWebPushProvisioningResponseJwsHeader]
type cardWebProvisionResponseAppleWebPushProvisioningResponseJwsHeaderJSON struct {
	Kid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardWebProvisionResponseAppleWebPushProvisioningResponseJwsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardWebProvisionResponseAppleWebPushProvisioningResponseJwsHeaderJSON) RawJSON() string {
	return r.raw
}

type CardWebProvisionResponseGoogleWebPushProvisioningResponse struct {
	// A base64 encoded and encrypted payload representing card data for the Google Pay
	// UWPP FPAN flow.
	GoogleOpc string `json:"google_opc"`
	// A base64 encoded and encrypted payload representing card data for the Google Pay
	// UWPP tokenization flow.
	TspOpc string                                                        `json:"tsp_opc"`
	JSON   cardWebProvisionResponseGoogleWebPushProvisioningResponseJSON `json:"-"`
}

// cardWebProvisionResponseGoogleWebPushProvisioningResponseJSON contains the JSON
// metadata for the struct
// [CardWebProvisionResponseGoogleWebPushProvisioningResponse]
type cardWebProvisionResponseGoogleWebPushProvisioningResponseJSON struct {
	GoogleOpc   apijson.Field
	TspOpc      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardWebProvisionResponseGoogleWebPushProvisioningResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardWebProvisionResponseGoogleWebPushProvisioningResponseJSON) RawJSON() string {
	return r.raw
}

func (r CardWebProvisionResponseGoogleWebPushProvisioningResponse) implementsCardWebProvisionResponse() {
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
	// Globally unique identifier for an existing bulk order to associate this card
	// with. When specified, the card will be added to the bulk order for batch
	// shipment. Only applicable to cards of type PHYSICAL
	BulkOrderToken param.Field[string] `json:"bulk_order_token" format:"uuid"`
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
	// Additional context or information related to the card that this card will
	// replace.
	ReplacementComment param.Field[string] `json:"replacement_comment"`
	// Globally unique identifier for the card that this card will replace. If the card
	// type is `PHYSICAL` it will be replaced by a `PHYSICAL` card. If the card type is
	// `VIRTUAL` it will be replaced by a `VIRTUAL` card.
	ReplacementFor param.Field[string] `json:"replacement_for" format:"uuid"`
	// Card state substatus values for the card that this card will replace:
	//
	//   - `LOST` - The physical card is no longer in the cardholder's possession due to
	//     being lost or never received by the cardholder.
	//   - `COMPROMISED` - Card information has been exposed, potentially leading to
	//     unauthorized access. This may involve physical card theft, cloning, or online
	//     data breaches.
	//   - `DAMAGED` - The physical card is not functioning properly, such as having chip
	//     failures or a demagnetized magnetic stripe.
	//   - `END_USER_REQUEST` - The cardholder requested the closure of the card for
	//     reasons unrelated to fraud or damage, such as switching to a different product
	//     or closing the account.
	//   - `ISSUER_REQUEST` - The issuer closed the card for reasons unrelated to fraud
	//     or damage, such as account inactivity, product or policy changes, or
	//     technology upgrades.
	//   - `NOT_ACTIVE` - The card hasn’t had any transaction activity for a specified
	//     period, applicable to statuses like `PAUSED` or `CLOSED`.
	//   - `SUSPICIOUS_ACTIVITY` - The card has one or more suspicious transactions or
	//     activities that require review. This can involve prompting the cardholder to
	//     confirm legitimate use or report confirmed fraud.
	//   - `INTERNAL_REVIEW` - The card is temporarily paused pending further internal
	//     review.
	//   - `EXPIRED` - The card has expired and has been closed without being reissued.
	//   - `UNDELIVERABLE` - The card cannot be delivered to the cardholder and has been
	//     returned.
	//   - `OTHER` - The reason for the status does not fall into any of the above
	//     categories. A comment should be provided to specify the reason.
	ReplacementSubstatus param.Field[CardNewParamsReplacementSubstatus] `json:"replacement_substatus"`
	ShippingAddress      param.Field[shared.ShippingAddressParam]       `json:"shipping_address"`
	// Shipping method for the card. Only applies to cards of type PHYSICAL. Use of
	// options besides `STANDARD` require additional permissions.
	//
	//   - `STANDARD` - USPS regular mail or similar international option, with no
	//     tracking
	//   - `STANDARD_WITH_TRACKING` - USPS regular mail or similar international option,
	//     with tracking
	//   - `PRIORITY` - USPS Priority, 1-3 day shipping, with tracking
	//   - `EXPRESS` - FedEx or UPS depending on card manufacturer, Express, 3-day
	//     shipping, with tracking
	//   - `2_DAY` - FedEx or UPS depending on card manufacturer, 2-day shipping, with
	//     tracking
	//   - `EXPEDITED` - FedEx or UPS depending on card manufacturer, Standard Overnight
	//     or similar international option, with tracking
	//   - `BULK_EXPEDITED` - Bulk shipment with Expedited shipping
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
	State          param.Field[CardNewParamsState] `json:"state"`
	IdempotencyKey param.Field[string]             `header:"Idempotency-Key" format:"uuid"`
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

// Card state substatus values for the card that this card will replace:
//
//   - `LOST` - The physical card is no longer in the cardholder's possession due to
//     being lost or never received by the cardholder.
//   - `COMPROMISED` - Card information has been exposed, potentially leading to
//     unauthorized access. This may involve physical card theft, cloning, or online
//     data breaches.
//   - `DAMAGED` - The physical card is not functioning properly, such as having chip
//     failures or a demagnetized magnetic stripe.
//   - `END_USER_REQUEST` - The cardholder requested the closure of the card for
//     reasons unrelated to fraud or damage, such as switching to a different product
//     or closing the account.
//   - `ISSUER_REQUEST` - The issuer closed the card for reasons unrelated to fraud
//     or damage, such as account inactivity, product or policy changes, or
//     technology upgrades.
//   - `NOT_ACTIVE` - The card hasn’t had any transaction activity for a specified
//     period, applicable to statuses like `PAUSED` or `CLOSED`.
//   - `SUSPICIOUS_ACTIVITY` - The card has one or more suspicious transactions or
//     activities that require review. This can involve prompting the cardholder to
//     confirm legitimate use or report confirmed fraud.
//   - `INTERNAL_REVIEW` - The card is temporarily paused pending further internal
//     review.
//   - `EXPIRED` - The card has expired and has been closed without being reissued.
//   - `UNDELIVERABLE` - The card cannot be delivered to the cardholder and has been
//     returned.
//   - `OTHER` - The reason for the status does not fall into any of the above
//     categories. A comment should be provided to specify the reason.
type CardNewParamsReplacementSubstatus string

const (
	CardNewParamsReplacementSubstatusLost               CardNewParamsReplacementSubstatus = "LOST"
	CardNewParamsReplacementSubstatusCompromised        CardNewParamsReplacementSubstatus = "COMPROMISED"
	CardNewParamsReplacementSubstatusDamaged            CardNewParamsReplacementSubstatus = "DAMAGED"
	CardNewParamsReplacementSubstatusEndUserRequest     CardNewParamsReplacementSubstatus = "END_USER_REQUEST"
	CardNewParamsReplacementSubstatusIssuerRequest      CardNewParamsReplacementSubstatus = "ISSUER_REQUEST"
	CardNewParamsReplacementSubstatusNotActive          CardNewParamsReplacementSubstatus = "NOT_ACTIVE"
	CardNewParamsReplacementSubstatusSuspiciousActivity CardNewParamsReplacementSubstatus = "SUSPICIOUS_ACTIVITY"
	CardNewParamsReplacementSubstatusInternalReview     CardNewParamsReplacementSubstatus = "INTERNAL_REVIEW"
	CardNewParamsReplacementSubstatusExpired            CardNewParamsReplacementSubstatus = "EXPIRED"
	CardNewParamsReplacementSubstatusUndeliverable      CardNewParamsReplacementSubstatus = "UNDELIVERABLE"
	CardNewParamsReplacementSubstatusOther              CardNewParamsReplacementSubstatus = "OTHER"
)

func (r CardNewParamsReplacementSubstatus) IsKnown() bool {
	switch r {
	case CardNewParamsReplacementSubstatusLost, CardNewParamsReplacementSubstatusCompromised, CardNewParamsReplacementSubstatusDamaged, CardNewParamsReplacementSubstatusEndUserRequest, CardNewParamsReplacementSubstatusIssuerRequest, CardNewParamsReplacementSubstatusNotActive, CardNewParamsReplacementSubstatusSuspiciousActivity, CardNewParamsReplacementSubstatusInternalReview, CardNewParamsReplacementSubstatusExpired, CardNewParamsReplacementSubstatusUndeliverable, CardNewParamsReplacementSubstatusOther:
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
//   - `EXPRESS` - FedEx or UPS depending on card manufacturer, Express, 3-day
//     shipping, with tracking
//   - `2_DAY` - FedEx or UPS depending on card manufacturer, 2-day shipping, with
//     tracking
//   - `EXPEDITED` - FedEx or UPS depending on card manufacturer, Standard Overnight
//     or similar international option, with tracking
//   - `BULK_EXPEDITED` - Bulk shipment with Expedited shipping
type CardNewParamsShippingMethod string

const (
	CardNewParamsShippingMethod2Day                 CardNewParamsShippingMethod = "2_DAY"
	CardNewParamsShippingMethodBulkExpedited        CardNewParamsShippingMethod = "BULK_EXPEDITED"
	CardNewParamsShippingMethodExpedited            CardNewParamsShippingMethod = "EXPEDITED"
	CardNewParamsShippingMethodExpress              CardNewParamsShippingMethod = "EXPRESS"
	CardNewParamsShippingMethodPriority             CardNewParamsShippingMethod = "PRIORITY"
	CardNewParamsShippingMethodStandard             CardNewParamsShippingMethod = "STANDARD"
	CardNewParamsShippingMethodStandardWithTracking CardNewParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

func (r CardNewParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardNewParamsShippingMethod2Day, CardNewParamsShippingMethodBulkExpedited, CardNewParamsShippingMethodExpedited, CardNewParamsShippingMethodExpress, CardNewParamsShippingMethodPriority, CardNewParamsShippingMethodStandard, CardNewParamsShippingMethodStandardWithTracking:
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
	// Additional context or information related to the card.
	Comment param.Field[string] `json:"comment"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken param.Field[string] `json:"digital_card_art_token" format:"uuid"`
	// Friendly name to identify the card.
	Memo param.Field[string] `json:"memo"`
	// Globally unique identifier for the card's network program. Currently applicable
	// to Visa cards participating in Account Level Management only.
	NetworkProgramToken param.Field[string] `json:"network_program_token" format:"uuid"`
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
	// Card state substatus values:
	//
	//   - `LOST` - The physical card is no longer in the cardholder's possession due to
	//     being lost or never received by the cardholder.
	//   - `COMPROMISED` - Card information has been exposed, potentially leading to
	//     unauthorized access. This may involve physical card theft, cloning, or online
	//     data breaches.
	//   - `DAMAGED` - The physical card is not functioning properly, such as having chip
	//     failures or a demagnetized magnetic stripe.
	//   - `END_USER_REQUEST` - The cardholder requested the closure of the card for
	//     reasons unrelated to fraud or damage, such as switching to a different product
	//     or closing the account.
	//   - `ISSUER_REQUEST` - The issuer closed the card for reasons unrelated to fraud
	//     or damage, such as account inactivity, product or policy changes, or
	//     technology upgrades.
	//   - `NOT_ACTIVE` - The card hasn’t had any transaction activity for a specified
	//     period, applicable to statuses like `PAUSED` or `CLOSED`.
	//   - `SUSPICIOUS_ACTIVITY` - The card has one or more suspicious transactions or
	//     activities that require review. This can involve prompting the cardholder to
	//     confirm legitimate use or report confirmed fraud.
	//   - `INTERNAL_REVIEW` - The card is temporarily paused pending further internal
	//     review.
	//   - `EXPIRED` - The card has expired and has been closed without being reissued.
	//   - `UNDELIVERABLE` - The card cannot be delivered to the cardholder and has been
	//     returned.
	//   - `OTHER` - The reason for the status does not fall into any of the above
	//     categories. A comment should be provided to specify the reason.
	Substatus param.Field[CardUpdateParamsSubstatus] `json:"substatus"`
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

// Card state substatus values:
//
//   - `LOST` - The physical card is no longer in the cardholder's possession due to
//     being lost or never received by the cardholder.
//   - `COMPROMISED` - Card information has been exposed, potentially leading to
//     unauthorized access. This may involve physical card theft, cloning, or online
//     data breaches.
//   - `DAMAGED` - The physical card is not functioning properly, such as having chip
//     failures or a demagnetized magnetic stripe.
//   - `END_USER_REQUEST` - The cardholder requested the closure of the card for
//     reasons unrelated to fraud or damage, such as switching to a different product
//     or closing the account.
//   - `ISSUER_REQUEST` - The issuer closed the card for reasons unrelated to fraud
//     or damage, such as account inactivity, product or policy changes, or
//     technology upgrades.
//   - `NOT_ACTIVE` - The card hasn’t had any transaction activity for a specified
//     period, applicable to statuses like `PAUSED` or `CLOSED`.
//   - `SUSPICIOUS_ACTIVITY` - The card has one or more suspicious transactions or
//     activities that require review. This can involve prompting the cardholder to
//     confirm legitimate use or report confirmed fraud.
//   - `INTERNAL_REVIEW` - The card is temporarily paused pending further internal
//     review.
//   - `EXPIRED` - The card has expired and has been closed without being reissued.
//   - `UNDELIVERABLE` - The card cannot be delivered to the cardholder and has been
//     returned.
//   - `OTHER` - The reason for the status does not fall into any of the above
//     categories. A comment should be provided to specify the reason.
type CardUpdateParamsSubstatus string

const (
	CardUpdateParamsSubstatusLost               CardUpdateParamsSubstatus = "LOST"
	CardUpdateParamsSubstatusCompromised        CardUpdateParamsSubstatus = "COMPROMISED"
	CardUpdateParamsSubstatusDamaged            CardUpdateParamsSubstatus = "DAMAGED"
	CardUpdateParamsSubstatusEndUserRequest     CardUpdateParamsSubstatus = "END_USER_REQUEST"
	CardUpdateParamsSubstatusIssuerRequest      CardUpdateParamsSubstatus = "ISSUER_REQUEST"
	CardUpdateParamsSubstatusNotActive          CardUpdateParamsSubstatus = "NOT_ACTIVE"
	CardUpdateParamsSubstatusSuspiciousActivity CardUpdateParamsSubstatus = "SUSPICIOUS_ACTIVITY"
	CardUpdateParamsSubstatusInternalReview     CardUpdateParamsSubstatus = "INTERNAL_REVIEW"
	CardUpdateParamsSubstatusExpired            CardUpdateParamsSubstatus = "EXPIRED"
	CardUpdateParamsSubstatusUndeliverable      CardUpdateParamsSubstatus = "UNDELIVERABLE"
	CardUpdateParamsSubstatusOther              CardUpdateParamsSubstatus = "OTHER"
)

func (r CardUpdateParamsSubstatus) IsKnown() bool {
	switch r {
	case CardUpdateParamsSubstatusLost, CardUpdateParamsSubstatusCompromised, CardUpdateParamsSubstatusDamaged, CardUpdateParamsSubstatusEndUserRequest, CardUpdateParamsSubstatusIssuerRequest, CardUpdateParamsSubstatusNotActive, CardUpdateParamsSubstatusSuspiciousActivity, CardUpdateParamsSubstatusInternalReview, CardUpdateParamsSubstatusExpired, CardUpdateParamsSubstatusUndeliverable, CardUpdateParamsSubstatusOther:
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
	// Returns cards containing the specified partial or full memo text.
	Memo param.Field[string] `query:"memo"`
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
	//   - `EXPRESS` - FedEx or UPS depending on card manufacturer, Express, 3-day
	//     shipping, with tracking
	//   - `2_DAY` - FedEx or UPS depending on card manufacturer, 2-day shipping, with
	//     tracking
	//   - `EXPEDITED` - FedEx or UPS depending on card manufacturer, Standard Overnight
	//     or similar international option, with tracking
	//   - `BULK_EXPEDITED` - Bulk shipment with Expedited shipping
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
//   - `EXPRESS` - FedEx or UPS depending on card manufacturer, Express, 3-day
//     shipping, with tracking
//   - `2_DAY` - FedEx or UPS depending on card manufacturer, 2-day shipping, with
//     tracking
//   - `EXPEDITED` - FedEx or UPS depending on card manufacturer, Standard Overnight
//     or similar international option, with tracking
//   - `BULK_EXPEDITED` - Bulk shipment with Expedited shipping
type CardConvertPhysicalParamsShippingMethod string

const (
	CardConvertPhysicalParamsShippingMethod2Day                 CardConvertPhysicalParamsShippingMethod = "2_DAY"
	CardConvertPhysicalParamsShippingMethodBulkExpedited        CardConvertPhysicalParamsShippingMethod = "BULK_EXPEDITED"
	CardConvertPhysicalParamsShippingMethodExpedited            CardConvertPhysicalParamsShippingMethod = "EXPEDITED"
	CardConvertPhysicalParamsShippingMethodExpress              CardConvertPhysicalParamsShippingMethod = "EXPRESS"
	CardConvertPhysicalParamsShippingMethodPriority             CardConvertPhysicalParamsShippingMethod = "PRIORITY"
	CardConvertPhysicalParamsShippingMethodStandard             CardConvertPhysicalParamsShippingMethod = "STANDARD"
	CardConvertPhysicalParamsShippingMethodStandardWithTracking CardConvertPhysicalParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

func (r CardConvertPhysicalParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardConvertPhysicalParamsShippingMethod2Day, CardConvertPhysicalParamsShippingMethodBulkExpedited, CardConvertPhysicalParamsShippingMethodExpedited, CardConvertPhysicalParamsShippingMethodExpress, CardConvertPhysicalParamsShippingMethodPriority, CardConvertPhysicalParamsShippingMethodStandard, CardConvertPhysicalParamsShippingMethodStandardWithTracking:
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
	//   - `EXPRESS` - FedEx or UPS depending on card manufacturer, Express, 3-day
	//     shipping, with tracking
	//   - `2_DAY` - FedEx or UPS depending on card manufacturer, 2-day shipping, with
	//     tracking
	//   - `EXPEDITED` - FedEx or UPS depending on card manufacturer, Standard Overnight
	//     or similar international option, with tracking
	//   - `BULK_EXPEDITED` - Bulk shipment with Expedited shipping
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
//   - `EXPRESS` - FedEx or UPS depending on card manufacturer, Express, 3-day
//     shipping, with tracking
//   - `2_DAY` - FedEx or UPS depending on card manufacturer, 2-day shipping, with
//     tracking
//   - `EXPEDITED` - FedEx or UPS depending on card manufacturer, Standard Overnight
//     or similar international option, with tracking
//   - `BULK_EXPEDITED` - Bulk shipment with Expedited shipping
type CardReissueParamsShippingMethod string

const (
	CardReissueParamsShippingMethod2Day                 CardReissueParamsShippingMethod = "2_DAY"
	CardReissueParamsShippingMethodBulkExpedited        CardReissueParamsShippingMethod = "BULK_EXPEDITED"
	CardReissueParamsShippingMethodExpedited            CardReissueParamsShippingMethod = "EXPEDITED"
	CardReissueParamsShippingMethodExpress              CardReissueParamsShippingMethod = "EXPRESS"
	CardReissueParamsShippingMethodPriority             CardReissueParamsShippingMethod = "PRIORITY"
	CardReissueParamsShippingMethodStandard             CardReissueParamsShippingMethod = "STANDARD"
	CardReissueParamsShippingMethodStandardWithTracking CardReissueParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

func (r CardReissueParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardReissueParamsShippingMethod2Day, CardReissueParamsShippingMethodBulkExpedited, CardReissueParamsShippingMethodExpedited, CardReissueParamsShippingMethodExpress, CardReissueParamsShippingMethodPriority, CardReissueParamsShippingMethodStandard, CardReissueParamsShippingMethodStandardWithTracking:
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
	//   - `EXPRESS` - FedEx or UPS depending on card manufacturer, Express, 3-day
	//     shipping, with tracking
	//   - `2_DAY` - FedEx or UPS depending on card manufacturer, 2-day shipping, with
	//     tracking
	//   - `EXPEDITED` - FedEx or UPS depending on card manufacturer, Standard Overnight
	//     or similar international option, with tracking
	//   - `BULK_EXPEDITED` - Bulk shipment with Expedited shipping
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
//   - `EXPRESS` - FedEx or UPS depending on card manufacturer, Express, 3-day
//     shipping, with tracking
//   - `2_DAY` - FedEx or UPS depending on card manufacturer, 2-day shipping, with
//     tracking
//   - `EXPEDITED` - FedEx or UPS depending on card manufacturer, Standard Overnight
//     or similar international option, with tracking
//   - `BULK_EXPEDITED` - Bulk shipment with Expedited shipping
type CardRenewParamsShippingMethod string

const (
	CardRenewParamsShippingMethod2Day                 CardRenewParamsShippingMethod = "2_DAY"
	CardRenewParamsShippingMethodBulkExpedited        CardRenewParamsShippingMethod = "BULK_EXPEDITED"
	CardRenewParamsShippingMethodExpedited            CardRenewParamsShippingMethod = "EXPEDITED"
	CardRenewParamsShippingMethodExpress              CardRenewParamsShippingMethod = "EXPRESS"
	CardRenewParamsShippingMethodPriority             CardRenewParamsShippingMethod = "PRIORITY"
	CardRenewParamsShippingMethodStandard             CardRenewParamsShippingMethod = "STANDARD"
	CardRenewParamsShippingMethodStandardWithTracking CardRenewParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

func (r CardRenewParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardRenewParamsShippingMethod2Day, CardRenewParamsShippingMethodBulkExpedited, CardRenewParamsShippingMethodExpedited, CardRenewParamsShippingMethodExpress, CardRenewParamsShippingMethodPriority, CardRenewParamsShippingMethodStandard, CardRenewParamsShippingMethodStandardWithTracking:
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

type CardWebProvisionParams struct {
	// Only applicable if `digital_wallet` is GOOGLE_PAY. Google Pay Web Push
	// Provisioning device identifier required for the tokenization flow
	ClientDeviceID param.Field[string] `json:"client_device_id" format:"uuid"`
	// Only applicable if `digital_wallet` is GOOGLE_PAY. Google Pay Web Push
	// Provisioning wallet account identifier required for the tokenization flow
	ClientWalletAccountID param.Field[string] `json:"client_wallet_account_id" format:"uuid"`
	// Name of digital wallet provider.
	DigitalWallet param.Field[CardWebProvisionParamsDigitalWallet] `json:"digital_wallet"`
	// Only applicable if `digital_wallet` is GOOGLE_PAY. Google Pay Web Push
	// Provisioning session identifier required for the FPAN flow.
	ServerSessionID param.Field[string] `json:"server_session_id" format:"uuid"`
}

func (r CardWebProvisionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Name of digital wallet provider.
type CardWebProvisionParamsDigitalWallet string

const (
	CardWebProvisionParamsDigitalWalletApplePay  CardWebProvisionParamsDigitalWallet = "APPLE_PAY"
	CardWebProvisionParamsDigitalWalletGooglePay CardWebProvisionParamsDigitalWallet = "GOOGLE_PAY"
)

func (r CardWebProvisionParamsDigitalWallet) IsKnown() bool {
	switch r {
	case CardWebProvisionParamsDigitalWalletApplePay, CardWebProvisionParamsDigitalWalletGooglePay:
		return true
	}
	return false
}
