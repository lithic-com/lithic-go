// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// CardService contains methods and other services that help with interacting with
// the lithic API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewCardService] method instead.
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

// Create a new virtual or physical card. Parameters `pin`, `shipping_address`, and
// `product_id` only apply to physical cards.
func (r *CardService) New(ctx context.Context, body CardNewParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := "cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get card configuration such as spend limit and state.
func (r *CardService) Get(ctx context.Context, cardToken string, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the specified properties of the card. Unsupplied properties will remain
// unchanged. `pin` parameter only applies to physical cards.
//
// _Note: setting a card to a `CLOSED` state is a final action that cannot be
// undone._
func (r *CardService) Update(ctx context.Context, cardToken string, body CardUpdateParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List cards.
func (r *CardService) List(ctx context.Context, query CardListParams, opts ...option.RequestOption) (res *shared.CursorPage[Card], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cards"
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
func (r *CardService) ListAutoPaging(ctx context.Context, query CardListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[Card] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
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
func (r *CardService) Embed(ctx context.Context, query CardEmbedParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	path := "embed/card"
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
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/provision", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Initiate print and shipment of a duplicate physical card.
//
// Only applies to cards of type `PHYSICAL`.
func (r *CardService) Reissue(ctx context.Context, cardToken string, body CardReissueParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/reissue", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Initiate print and shipment of a renewed physical card.
//
// Only applies to cards of type `PHYSICAL`.
func (r *CardService) Renew(ctx context.Context, cardToken string, body CardRenewParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/renew", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Card's available spend limit, which is based on the spend limit configured
// on the Card and the amount already spent over the spend limit's duration. For
// example, if the Card has a monthly spend limit of $1000 configured, and has
// spent $600 in the last month, the available spend limit returned would be $400.
func (r *CardService) GetSpendLimits(ctx context.Context, cardToken string, opts ...option.RequestOption) (res *CardSpendLimits, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/spend_limits", cardToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get card configuration such as spend limit and state. Customers must be PCI
// compliant to use this endpoint. Please contact
// [support@lithic.com](mailto:support@lithic.com) for questions. _Note: this is a
// `POST` endpoint because it is more secure to send sensitive data in a request
// body than in a URL._
func (r *CardService) SearchByPan(ctx context.Context, body CardSearchByPanParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := "cards/search_by_pan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Card struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Globally unique identifier for the account to which the card belongs.
	AccountToken string `json:"account_token,required" format:"uuid"`
	// Globally unique identifier for the card program on which the card exists.
	CardProgramToken string `json:"card_program_token,required" format:"uuid"`
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time   `json:"created,required" format:"date-time"`
	Funding CardFunding `json:"funding,required"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined.
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
	// Card state values:
	//
	//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
	//     be undone.
	//   - `OPEN` - Card will approve authorizations (if they match card and account
	//     parameters).
	//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
	//     time.
	//   - `PENDING_FULFILLMENT` - The initial state for cards of type `PHYSICAL`. The
	//     card is provisioned pending manufacturing and fulfillment. Cards in this state
	//     can accept authorizations for e-commerce purchases, but not for "Card Present"
	//     purchases where the physical card itself is present.
	//   - `PENDING_ACTIVATION` - Each business day at 2pm Eastern Time Zone (ET), cards
	//     of type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card
	//     production warehouse and updated to state `PENDING_ACTIVATION` . Similar to
	//     `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
	//     transactions. API clients should update the card's state to `OPEN` only after
	//     the cardholder confirms receipt of the card.
	//
	// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
	// manufactured.
	State CardState `json:"state,required"`
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
	Type CardType `json:"type,required"`
	// List of identifiers for the Auth Rule(s) that are applied on the card.
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// Three digit cvv printed on the back of the card.
	Cvv string `json:"cvv"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken string `json:"digital_card_art_token" format:"uuid"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card’s locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// [support@lithic.com](mailto:support@lithic.com) for questions.
	Pan string `json:"pan"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID string   `json:"product_id"`
	JSON      cardJSON `json:"-"`
}

// cardJSON contains the JSON metadata for the struct [Card]
type cardJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardProgramToken    apijson.Field
	Created             apijson.Field
	Funding             apijson.Field
	LastFour            apijson.Field
	SpendLimit          apijson.Field
	SpendLimitDuration  apijson.Field
	State               apijson.Field
	Type                apijson.Field
	AuthRuleTokens      apijson.Field
	Cvv                 apijson.Field
	DigitalCardArtToken apijson.Field
	ExpMonth            apijson.Field
	ExpYear             apijson.Field
	Hostname            apijson.Field
	Memo                apijson.Field
	Pan                 apijson.Field
	ProductID           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardJSON) RawJSON() string {
	return r.raw
}

type CardFunding struct {
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// State of funding source.
	//
	// Funding source states:
	//
	//   - `ENABLED` - The funding account is available to use for card creation and
	//     transactions.
	//   - `PENDING` - The funding account is still being verified e.g. bank
	//     micro-deposits verification.
	//   - `DELETED` - The founding account has been deleted.
	State CardFundingState `json:"state,required"`
	// Types of funding source:
	//
	// - `DEPOSITORY_CHECKING` - Bank checking account.
	// - `DEPOSITORY_SAVINGS` - Bank savings account.
	Type CardFundingType `json:"type,required"`
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string          `json:"nickname"`
	JSON     cardFundingJSON `json:"-"`
}

// cardFundingJSON contains the JSON metadata for the struct [CardFunding]
type cardFundingJSON struct {
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

func (r *CardFunding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardFundingJSON) RawJSON() string {
	return r.raw
}

// State of funding source.
//
// Funding source states:
//
//   - `ENABLED` - The funding account is available to use for card creation and
//     transactions.
//   - `PENDING` - The funding account is still being verified e.g. bank
//     micro-deposits verification.
//   - `DELETED` - The founding account has been deleted.
type CardFundingState string

const (
	CardFundingStateDeleted CardFundingState = "DELETED"
	CardFundingStateEnabled CardFundingState = "ENABLED"
	CardFundingStatePending CardFundingState = "PENDING"
)

// Types of funding source:
//
// - `DEPOSITORY_CHECKING` - Bank checking account.
// - `DEPOSITORY_SAVINGS` - Bank savings account.
type CardFundingType string

const (
	CardFundingTypeDepositoryChecking CardFundingType = "DEPOSITORY_CHECKING"
	CardFundingTypeDepositorySavings  CardFundingType = "DEPOSITORY_SAVINGS"
)

// Card state values:
//
//   - `CLOSED` - Card will no longer approve authorizations. Closing a card cannot
//     be undone.
//   - `OPEN` - Card will approve authorizations (if they match card and account
//     parameters).
//   - `PAUSED` - Card will decline authorizations, but can be resumed at a later
//     time.
//   - `PENDING_FULFILLMENT` - The initial state for cards of type `PHYSICAL`. The
//     card is provisioned pending manufacturing and fulfillment. Cards in this state
//     can accept authorizations for e-commerce purchases, but not for "Card Present"
//     purchases where the physical card itself is present.
//   - `PENDING_ACTIVATION` - Each business day at 2pm Eastern Time Zone (ET), cards
//     of type `PHYSICAL` in state `PENDING_FULFILLMENT` are sent to the card
//     production warehouse and updated to state `PENDING_ACTIVATION` . Similar to
//     `PENDING_FULFILLMENT`, cards in this state can be used for e-commerce
//     transactions. API clients should update the card's state to `OPEN` only after
//     the cardholder confirms receipt of the card.
//
// In sandbox, the same daily batch fulfillment occurs, but no cards are actually
// manufactured.
type CardState string

const (
	CardStateClosed             CardState = "CLOSED"
	CardStateOpen               CardState = "OPEN"
	CardStatePaused             CardState = "PAUSED"
	CardStatePendingActivation  CardState = "PENDING_ACTIVATION"
	CardStatePendingFulfillment CardState = "PENDING_FULFILLMENT"
)

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
type CardType string

const (
	CardTypeMerchantLocked CardType = "MERCHANT_LOCKED"
	CardTypePhysical       CardType = "PHYSICAL"
	CardTypeSingleUse      CardType = "SINGLE_USE"
	CardTypeVirtual        CardType = "VIRTUAL"
)

type CardSpendLimits struct {
	AvailableSpendLimit CardSpendLimitsAvailableSpendLimit `json:"available_spend_limit,required"`
	JSON                cardSpendLimitsJSON                `json:"-"`
}

// cardSpendLimitsJSON contains the JSON metadata for the struct [CardSpendLimits]
type cardSpendLimitsJSON struct {
	AvailableSpendLimit apijson.Field
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
	// The available spend limit relative to the annual limit configured on the Card.
	Annually int64 `json:"annually"`
	// The available spend limit relative to the forever limit configured on the Card.
	Forever int64 `json:"forever"`
	// The available spend limit relative to the monthly limit configured on the Card.
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
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo param.Field[string] `json:"memo"`
	// Encrypted PIN block (in base64). Only applies to cards of type `PHYSICAL` and
	// `VIRTUAL`. See
	// [Encrypted PIN Block](https://docs.lithic.com/docs/cards#encrypted-pin-block-enterprise).
	Pin param.Field[string] `json:"pin"`
	// Only applicable to cards of type `PHYSICAL`. This must be configured with Lithic
	// before use. Specifies the configuration (i.e., physical card art) that the card
	// should be manufactured with.
	ProductID param.Field[string] `json:"product_id"`
	// Only applicable to cards of type `PHYSICAL`. Globally unique identifier for the
	// card that this physical card will replace.
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
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the card
	// limit.
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
type CardNewParamsType string

const (
	CardNewParamsTypeMerchantLocked CardNewParamsType = "MERCHANT_LOCKED"
	CardNewParamsTypePhysical       CardNewParamsType = "PHYSICAL"
	CardNewParamsTypeSingleUse      CardNewParamsType = "SINGLE_USE"
	CardNewParamsTypeVirtual        CardNewParamsType = "VIRTUAL"
)

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

type CardUpdateParams struct {
	// Identifier for any Auth Rules that will be applied to transactions taking place
	// with the card.
	AuthRuleToken param.Field[string] `json:"auth_rule_token"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken param.Field[string] `json:"digital_card_art_token" format:"uuid"`
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo param.Field[string] `json:"memo"`
	// Encrypted PIN block (in base64). Only applies to cards of type `PHYSICAL` and
	// `VIRTUAL`. See
	// [Encrypted PIN Block](https://docs.lithic.com/docs/cards#encrypted-pin-block-enterprise).
	Pin param.Field[string] `json:"pin"`
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the card
	// limit.
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

type CardReissueParams struct {
	// If omitted, the previous carrier will be used.
	Carrier param.Field[shared.CarrierParam] `json:"carrier"`
	// Specifies the configuration (e.g. physical card art) that the card should be
	// manufactured with, and only applies to cards of type `PHYSICAL`. This must be
	// configured with Lithic before use.
	ProductID param.Field[string] `json:"product_id"`
	// If omitted, the previous shipping address will be used.
	ShippingAddress param.Field[shared.ShippingAddressParam] `json:"shipping_address"`
	// Shipping method for the card. Use of options besides `STANDARD` require
	// additional permissions.
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

// Shipping method for the card. Use of options besides `STANDARD` require
// additional permissions.
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
	CardReissueParamsShippingMethod2Day                 CardReissueParamsShippingMethod = "2-DAY"
	CardReissueParamsShippingMethodExpedited            CardReissueParamsShippingMethod = "EXPEDITED"
	CardReissueParamsShippingMethodExpress              CardReissueParamsShippingMethod = "EXPRESS"
	CardReissueParamsShippingMethodPriority             CardReissueParamsShippingMethod = "PRIORITY"
	CardReissueParamsShippingMethodStandard             CardReissueParamsShippingMethod = "STANDARD"
	CardReissueParamsShippingMethodStandardWithTracking CardReissueParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

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
	// Shipping method for the card. Use of options besides `STANDARD` require
	// additional permissions.
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

// Shipping method for the card. Use of options besides `STANDARD` require
// additional permissions.
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
	CardRenewParamsShippingMethod2Day                 CardRenewParamsShippingMethod = "2-DAY"
	CardRenewParamsShippingMethodExpedited            CardRenewParamsShippingMethod = "EXPEDITED"
	CardRenewParamsShippingMethodExpress              CardRenewParamsShippingMethod = "EXPRESS"
	CardRenewParamsShippingMethodPriority             CardRenewParamsShippingMethod = "PRIORITY"
	CardRenewParamsShippingMethodStandard             CardRenewParamsShippingMethod = "STANDARD"
	CardRenewParamsShippingMethodStandardWithTracking CardRenewParamsShippingMethod = "STANDARD_WITH_TRACKING"
)

type CardSearchByPanParams struct {
	// The PAN for the card being retrieved.
	Pan param.Field[string] `json:"pan,required"`
}

func (r CardSearchByPanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
