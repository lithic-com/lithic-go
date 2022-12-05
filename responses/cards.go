package responses

import (
	"time"

	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/pagination"
)

type Card struct {
	// An RFC 3339 timestamp for when the card was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Three digit cvv printed on the back of the card.
	Cvv     string        `json:"cvv"`
	Funding FundingSource `json:"funding,required"`
	// Two digit (MM) expiry month.
	ExpMonth string `json:"exp_month"`
	// Four digit (yyyy) expiry year.
	ExpYear string `json:"exp_year"`
	// Hostname of card’s locked merchant (will be empty if not applicable).
	Hostname string `json:"hostname"`
	// Last four digits of the card number.
	LastFour string `json:"last_four,required"`
	// Friendly name to identify the card. We recommend against using this field to
	// store JSON data as it can cause unexpected behavior.
	Memo string `json:"memo"`
	// Primary Account Number (PAN) (i.e. the card number). Customers must be PCI
	// compliant to have PAN returned as a field in production. Please contact
	// [support@lithic.com](mailto:support@lithic.com) for questions.
	Pan string `json:"pan"`
	// Amount (in cents) to limit approved authorizations. Transaction requests above
	// the spend limit will be declined.
	SpendLimit int64 `json:"spend_limit,required"`
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
	// List of identifiers for the Auth Rule(s) that are applied on the card.
	AuthRuleTokens []string `json:"auth_rule_tokens"`
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
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
	Type CardType `json:"type,required"`
	// Specifies the digital card art to be displayed in the user’s digital wallet
	// after tokenization. This artwork must be approved by Mastercard and configured
	// by Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken string `json:"digital_card_art_token" format:"uuid"`
	JSON                CardJSON
}

type CardJSON struct {
	Created             pjson.Metadata
	Cvv                 pjson.Metadata
	Funding             pjson.Metadata
	ExpMonth            pjson.Metadata
	ExpYear             pjson.Metadata
	Hostname            pjson.Metadata
	LastFour            pjson.Metadata
	Memo                pjson.Metadata
	Pan                 pjson.Metadata
	SpendLimit          pjson.Metadata
	SpendLimitDuration  pjson.Metadata
	State               pjson.Metadata
	AuthRuleTokens      pjson.Metadata
	Token               pjson.Metadata
	Type                pjson.Metadata
	DigitalCardArtToken pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Card using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SpendLimitDuration string

const (
	SpendLimitDurationAnnually    SpendLimitDuration = "ANNUALLY"
	SpendLimitDurationForever     SpendLimitDuration = "FOREVER"
	SpendLimitDurationMonthly     SpendLimitDuration = "MONTHLY"
	SpendLimitDurationTransaction SpendLimitDuration = "TRANSACTION"
)

type CardState string

const (
	CardStateClosed             CardState = "CLOSED"
	CardStateOpen               CardState = "OPEN"
	CardStatePaused             CardState = "PAUSED"
	CardStatePendingActivation  CardState = "PENDING_ACTIVATION"
	CardStatePendingFulfillment CardState = "PENDING_FULFILLMENT"
)

type CardType string

const (
	CardTypeVirtual        CardType = "VIRTUAL"
	CardTypePhysical       CardType = "PHYSICAL"
	CardTypeMerchantLocked CardType = "MERCHANT_LOCKED"
	CardTypeSingleUse      CardType = "SINGLE_USE"
)

type EmbedRequestParams struct {
	// A publicly available URI, so the white-labeled card element can be styled with
	// the client's branding.
	Css string `json:"css"`
	// An RFC 3339 timestamp for when the request should expire. UTC time zone.
	//
	// If no timezone is specified, UTC will be used. If payload does not contain an
	// expiration, the request will never expire.
	//
	// Using an `expiration` reduces the risk of a
	// [replay attack](https://en.wikipedia.org/wiki/Replay_attack). Without supplying
	// the `expiration`, in the event that a malicious user gets a copy of your request
	// in transit, they will be able to obtain the response data indefinitely.
	Expiration time.Time `json:"expiration" format:"date-time"`
	// Globally unique identifier for the card to be displayed.
	Token string `json:"token,required" format:"uuid"`
	// Required if you want to post the element clicked to the parent iframe.
	//
	// If you supply this param, you can also capture click events in the parent iframe
	// by adding an event listener.
	TargetOrigin string `json:"target_origin"`
	JSON         EmbedRequestParamsJSON
}

type EmbedRequestParamsJSON struct {
	Css          pjson.Metadata
	Expiration   pjson.Metadata
	Token        pjson.Metadata
	TargetOrigin pjson.Metadata
	Raw          []byte
	Extras       map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EmbedRequestParams using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EmbedRequestParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardProvisionResponse struct {
	ProvisioningPayload string `json:"provisioning_payload"`
	JSON                CardProvisionResponseJSON
}

type CardProvisionResponseJSON struct {
	ProvisioningPayload pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardProvisionResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardProvisionResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardListResponse struct {
	Data []Card `json:"data,required"`
	// Page number.
	Page int64 `json:"page,required"`
	// Total number of entries.
	TotalEntries int64 `json:"total_entries,required"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages,required"`
	JSON       CardListResponseJSON
}

type CardListResponseJSON struct {
	Data         pjson.Metadata
	Page         pjson.Metadata
	TotalEntries pjson.Metadata
	TotalPages   pjson.Metadata
	Raw          []byte
	Extras       map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardsPage struct {
	*pagination.Page[Card]
}

func (r *CardsPage) Card() *Card {
	return r.Current()
}

func (r *CardsPage) NextPage() (*CardsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &CardsPage{page}, nil
	}
}
