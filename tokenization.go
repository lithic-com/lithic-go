// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// TokenizationService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenizationService] method instead.
type TokenizationService struct {
	Options []option.RequestOption
}

// NewTokenizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenizationService(opts ...option.RequestOption) (r *TokenizationService) {
	r = &TokenizationService{}
	r.Options = opts
	return
}

// Get tokenization
func (r *TokenizationService) Get(ctx context.Context, tokenizationToken string, opts ...option.RequestOption) (res *Tokenization, err error) {
	opts = slices.Concat(r.Options, opts)
	if tokenizationToken == "" {
		err = errors.New("missing required tokenization_token parameter")
		return
	}
	path := fmt.Sprintf("v1/tokenizations/%s", tokenizationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List card tokenizations
func (r *TokenizationService) List(ctx context.Context, query TokenizationListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Tokenization], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/tokenizations"
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

// List card tokenizations
func (r *TokenizationService) ListAutoPaging(ctx context.Context, query TokenizationListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Tokenization] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to ask the card network to activate a tokenization. A
// successful response indicates that the request was successfully delivered to the
// card network. When the card network activates the tokenization, the state will
// be updated and a tokenization.updated event will be sent. The endpoint may only
// be used on digital wallet tokenizations with status `INACTIVE`,
// `PENDING_ACTIVATION`, or `PENDING_2FA`. This will put the tokenization in an
// active state, and transactions will be allowed. Reach out at
// [lithic.com/contact](https://lithic.com/contact) for more information.
func (r *TokenizationService) Activate(ctx context.Context, tokenizationToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tokenizationToken == "" {
		err = errors.New("missing required tokenization_token parameter")
		return
	}
	path := fmt.Sprintf("v1/tokenizations/%s/activate", tokenizationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This endpoint is used to ask the card network to deactivate a tokenization. A
// successful response indicates that the request was successfully delivered to the
// card network. When the card network deactivates the tokenization, the state will
// be updated and a tokenization.updated event will be sent. Authorizations
// attempted with a deactivated tokenization will be blocked and will not be
// forwarded to Lithic from the network. Deactivating the token is a permanent
// operation. If the target is a digital wallet tokenization, it will be removed
// from its device. Reach out at [lithic.com/contact](https://lithic.com/contact)
// for more information.
func (r *TokenizationService) Deactivate(ctx context.Context, tokenizationToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tokenizationToken == "" {
		err = errors.New("missing required tokenization_token parameter")
		return
	}
	path := fmt.Sprintf("v1/tokenizations/%s/deactivate", tokenizationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This endpoint is used to ask the card network to pause a tokenization. A
// successful response indicates that the request was successfully delivered to the
// card network. When the card network pauses the tokenization, the state will be
// updated and a tokenization.updated event will be sent. The endpoint may only be
// used on tokenizations with status `ACTIVE`. A paused token will prevent
// merchants from sending authorizations, and is a temporary status that can be
// changed. Reach out at [lithic.com/contact](https://lithic.com/contact) for more
// information.
func (r *TokenizationService) Pause(ctx context.Context, tokenizationToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tokenizationToken == "" {
		err = errors.New("missing required tokenization_token parameter")
		return
	}
	path := fmt.Sprintf("v1/tokenizations/%s/pause", tokenizationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This endpoint is used to ask the card network to send another activation code to
// a cardholder that has already tried tokenizing a card. A successful response
// indicates that the request was successfully delivered to the card network. The
// endpoint may only be used on Mastercard digital wallet tokenizations with status
// `INACTIVE`, `PENDING_ACTIVATION`, or `PENDING_2FA`. The network will send a new
// activation code to the one of the contact methods provided in the initial
// tokenization flow. If a user fails to enter the code correctly 3 times, the
// contact method will not be eligible for resending the activation code, and the
// cardholder must restart the provision process. Reach out at
// [lithic.com/contact](https://lithic.com/contact) for more information.
func (r *TokenizationService) ResendActivationCode(ctx context.Context, tokenizationToken string, body TokenizationResendActivationCodeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tokenizationToken == "" {
		err = errors.New("missing required tokenization_token parameter")
		return
	}
	path := fmt.Sprintf("v1/tokenizations/%s/resend_activation_code", tokenizationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// This endpoint is used to simulate a card's tokenization in the Digital Wallet
// and merchant tokenization ecosystem.
func (r *TokenizationService) Simulate(ctx context.Context, body TokenizationSimulateParams, opts ...option.RequestOption) (res *Tokenization, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/simulate/tokenizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to ask the card network to unpause a tokenization. A
// successful response indicates that the request was successfully delivered to the
// card network. When the card network unpauses the tokenization, the state will be
// updated and a tokenization.updated event will be sent. The endpoint may only be
// used on tokenizations with status `PAUSED`. This will put the tokenization in an
// active state, and transactions may resume. Reach out at
// [lithic.com/contact](https://lithic.com/contact) for more information.
func (r *TokenizationService) Unpause(ctx context.Context, tokenizationToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tokenizationToken == "" {
		err = errors.New("missing required tokenization_token parameter")
		return
	}
	path := fmt.Sprintf("v1/tokenizations/%s/unpause", tokenizationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This endpoint is used update the digital card art for a digital wallet
// tokenization. A successful response indicates that the card network has updated
// the tokenization's art, and the tokenization's `digital_cart_art_token` field
// was updated. The endpoint may not be used on tokenizations with status
// `DEACTIVATED`. Note that this updates the art for one specific tokenization, not
// all tokenizations for a card. New tokenizations for a card will be created with
// the art referenced in the card object's `digital_card_art_token` field. Reach
// out at [lithic.com/contact](https://lithic.com/contact) for more information.
func (r *TokenizationService) UpdateDigitalCardArt(ctx context.Context, tokenizationToken string, body TokenizationUpdateDigitalCardArtParams, opts ...option.RequestOption) (res *Tokenization, err error) {
	opts = slices.Concat(r.Options, opts)
	if tokenizationToken == "" {
		err = errors.New("missing required tokenization_token parameter")
		return
	}
	path := fmt.Sprintf("v1/tokenizations/%s/update_digital_card_art", tokenizationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Tokenization struct {
	// Globally unique identifier for a Tokenization
	Token string `json:"token,required" format:"uuid"`
	// The account token associated with the card being tokenized.
	AccountToken string `json:"account_token,required" format:"uuid"`
	// The card token associated with the card being tokenized.
	CardToken string `json:"card_token,required" format:"uuid"`
	// Date and time when the tokenization first occurred. UTC time zone.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The dynamic pan assigned to the token by the network.
	Dpan string `json:"dpan,required,nullable"`
	// The status of the tokenization request
	Status TokenizationStatus `json:"status,required"`
	// The entity that requested the tokenization. For digital wallets, this will be
	// one of the defined wallet types. For merchant tokenizations, this will be a
	// free-form merchant name string.
	TokenRequestorName TokenizationTokenRequestorName `json:"token_requestor_name,required"`
	// The network's unique reference for the tokenization.
	TokenUniqueReference string `json:"token_unique_reference,required"`
	// The channel through which the tokenization was made.
	TokenizationChannel TokenizationTokenizationChannel `json:"tokenization_channel,required"`
	// Latest date and time when the tokenization was updated. UTC time zone.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The device identifier associated with the tokenization.
	DeviceID string `json:"device_id,nullable"`
	// Specifies the digital card art displayed in the user's digital wallet after
	// tokenization. This will be null if the tokenization was created without an
	// associated digital card art. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken string `json:"digital_card_art_token,nullable" format:"uuid"`
	// A list of events related to the tokenization.
	Events []TokenizationEvent `json:"events"`
	// The network's unique reference for the card that is tokenized.
	PaymentAccountReferenceID string           `json:"payment_account_reference_id,nullable"`
	JSON                      tokenizationJSON `json:"-"`
}

// tokenizationJSON contains the JSON metadata for the struct [Tokenization]
type tokenizationJSON struct {
	Token                     apijson.Field
	AccountToken              apijson.Field
	CardToken                 apijson.Field
	CreatedAt                 apijson.Field
	Dpan                      apijson.Field
	Status                    apijson.Field
	TokenRequestorName        apijson.Field
	TokenUniqueReference      apijson.Field
	TokenizationChannel       apijson.Field
	UpdatedAt                 apijson.Field
	DeviceID                  apijson.Field
	DigitalCardArtToken       apijson.Field
	Events                    apijson.Field
	PaymentAccountReferenceID apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Tokenization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationJSON) RawJSON() string {
	return r.raw
}

// The status of the tokenization request
type TokenizationStatus string

const (
	TokenizationStatusActive            TokenizationStatus = "ACTIVE"
	TokenizationStatusDeactivated       TokenizationStatus = "DEACTIVATED"
	TokenizationStatusInactive          TokenizationStatus = "INACTIVE"
	TokenizationStatusPaused            TokenizationStatus = "PAUSED"
	TokenizationStatusPending2Fa        TokenizationStatus = "PENDING_2FA"
	TokenizationStatusPendingActivation TokenizationStatus = "PENDING_ACTIVATION"
	TokenizationStatusUnknown           TokenizationStatus = "UNKNOWN"
)

func (r TokenizationStatus) IsKnown() bool {
	switch r {
	case TokenizationStatusActive, TokenizationStatusDeactivated, TokenizationStatusInactive, TokenizationStatusPaused, TokenizationStatusPending2Fa, TokenizationStatusPendingActivation, TokenizationStatusUnknown:
		return true
	}
	return false
}

// Digital wallet type
type TokenizationTokenRequestorName string

const (
	TokenizationTokenRequestorNameAmazonOne    TokenizationTokenRequestorName = "AMAZON_ONE"
	TokenizationTokenRequestorNameAndroidPay   TokenizationTokenRequestorName = "ANDROID_PAY"
	TokenizationTokenRequestorNameApplePay     TokenizationTokenRequestorName = "APPLE_PAY"
	TokenizationTokenRequestorNameFacebook     TokenizationTokenRequestorName = "FACEBOOK"
	TokenizationTokenRequestorNameFitbitPay    TokenizationTokenRequestorName = "FITBIT_PAY"
	TokenizationTokenRequestorNameGarminPay    TokenizationTokenRequestorName = "GARMIN_PAY"
	TokenizationTokenRequestorNameGooglePay    TokenizationTokenRequestorName = "GOOGLE_PAY"
	TokenizationTokenRequestorNameMicrosoftPay TokenizationTokenRequestorName = "MICROSOFT_PAY"
	TokenizationTokenRequestorNameNetflix      TokenizationTokenRequestorName = "NETFLIX"
	TokenizationTokenRequestorNameSamsungPay   TokenizationTokenRequestorName = "SAMSUNG_PAY"
	TokenizationTokenRequestorNameUnknown      TokenizationTokenRequestorName = "UNKNOWN"
	TokenizationTokenRequestorNameVisaCheckout TokenizationTokenRequestorName = "VISA_CHECKOUT"
)

func (r TokenizationTokenRequestorName) IsKnown() bool {
	switch r {
	case TokenizationTokenRequestorNameAmazonOne, TokenizationTokenRequestorNameAndroidPay, TokenizationTokenRequestorNameApplePay, TokenizationTokenRequestorNameFacebook, TokenizationTokenRequestorNameFitbitPay, TokenizationTokenRequestorNameGarminPay, TokenizationTokenRequestorNameGooglePay, TokenizationTokenRequestorNameMicrosoftPay, TokenizationTokenRequestorNameNetflix, TokenizationTokenRequestorNameSamsungPay, TokenizationTokenRequestorNameUnknown, TokenizationTokenRequestorNameVisaCheckout:
		return true
	}
	return false
}

// The channel through which the tokenization was made.
type TokenizationTokenizationChannel string

const (
	TokenizationTokenizationChannelDigitalWallet TokenizationTokenizationChannel = "DIGITAL_WALLET"
	TokenizationTokenizationChannelMerchant      TokenizationTokenizationChannel = "MERCHANT"
)

func (r TokenizationTokenizationChannel) IsKnown() bool {
	switch r {
	case TokenizationTokenizationChannelDigitalWallet, TokenizationTokenizationChannelMerchant:
		return true
	}
	return false
}

type TokenizationEvent struct {
	// Globally unique identifier for a Tokenization Event
	Token string `json:"token" format:"uuid"`
	// Date and time when the tokenization event first occurred. UTC time zone.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Enum representing the result of the tokenization event
	Result TokenizationEventsResult `json:"result"`
	// Results from rules that were evaluated for this tokenization
	RuleResults []TokenizationRuleResult `json:"rule_results"`
	// List of reasons why the tokenization was declined
	TokenizationDeclineReasons []TokenizationDeclineReason `json:"tokenization_decline_reasons"`
	// List of reasons why two-factor authentication was required
	TokenizationTfaReasons []TokenizationTfaReason `json:"tokenization_tfa_reasons"`
	// Enum representing the type of tokenization event that occurred
	Type TokenizationEventsType `json:"type"`
	JSON tokenizationEventJSON  `json:"-"`
}

// tokenizationEventJSON contains the JSON metadata for the struct
// [TokenizationEvent]
type tokenizationEventJSON struct {
	Token                      apijson.Field
	CreatedAt                  apijson.Field
	Result                     apijson.Field
	RuleResults                apijson.Field
	TokenizationDeclineReasons apijson.Field
	TokenizationTfaReasons     apijson.Field
	Type                       apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *TokenizationEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationEventJSON) RawJSON() string {
	return r.raw
}

// Enum representing the result of the tokenization event
type TokenizationEventsResult string

const (
	TokenizationEventsResultApproved                        TokenizationEventsResult = "APPROVED"
	TokenizationEventsResultDeclined                        TokenizationEventsResult = "DECLINED"
	TokenizationEventsResultNotificationDelivered           TokenizationEventsResult = "NOTIFICATION_DELIVERED"
	TokenizationEventsResultRequireAdditionalAuthentication TokenizationEventsResult = "REQUIRE_ADDITIONAL_AUTHENTICATION"
	TokenizationEventsResultTokenActivated                  TokenizationEventsResult = "TOKEN_ACTIVATED"
	TokenizationEventsResultTokenCreated                    TokenizationEventsResult = "TOKEN_CREATED"
	TokenizationEventsResultTokenDeactivated                TokenizationEventsResult = "TOKEN_DEACTIVATED"
	TokenizationEventsResultTokenDeletedFromConsumerApp     TokenizationEventsResult = "TOKEN_DELETED_FROM_CONSUMER_APP"
	TokenizationEventsResultTokenInactive                   TokenizationEventsResult = "TOKEN_INACTIVE"
	TokenizationEventsResultTokenStateUnknown               TokenizationEventsResult = "TOKEN_STATE_UNKNOWN"
	TokenizationEventsResultTokenSuspended                  TokenizationEventsResult = "TOKEN_SUSPENDED"
	TokenizationEventsResultTokenUpdated                    TokenizationEventsResult = "TOKEN_UPDATED"
)

func (r TokenizationEventsResult) IsKnown() bool {
	switch r {
	case TokenizationEventsResultApproved, TokenizationEventsResultDeclined, TokenizationEventsResultNotificationDelivered, TokenizationEventsResultRequireAdditionalAuthentication, TokenizationEventsResultTokenActivated, TokenizationEventsResultTokenCreated, TokenizationEventsResultTokenDeactivated, TokenizationEventsResultTokenDeletedFromConsumerApp, TokenizationEventsResultTokenInactive, TokenizationEventsResultTokenStateUnknown, TokenizationEventsResultTokenSuspended, TokenizationEventsResultTokenUpdated:
		return true
	}
	return false
}

// Enum representing the type of tokenization event that occurred
type TokenizationEventsType string

const (
	TokenizationEventsTypeTokenization2Fa              TokenizationEventsType = "TOKENIZATION_2FA"
	TokenizationEventsTypeTokenizationAuthorization    TokenizationEventsType = "TOKENIZATION_AUTHORIZATION"
	TokenizationEventsTypeTokenizationDecisioning      TokenizationEventsType = "TOKENIZATION_DECISIONING"
	TokenizationEventsTypeTokenizationEligibilityCheck TokenizationEventsType = "TOKENIZATION_ELIGIBILITY_CHECK"
	TokenizationEventsTypeTokenizationUpdated          TokenizationEventsType = "TOKENIZATION_UPDATED"
)

func (r TokenizationEventsType) IsKnown() bool {
	switch r {
	case TokenizationEventsTypeTokenization2Fa, TokenizationEventsTypeTokenizationAuthorization, TokenizationEventsTypeTokenizationDecisioning, TokenizationEventsTypeTokenizationEligibilityCheck, TokenizationEventsTypeTokenizationUpdated:
		return true
	}
	return false
}

// Reason code for why a tokenization was declined
type TokenizationDeclineReason string

const (
	TokenizationDeclineReasonAccountScore1                  TokenizationDeclineReason = "ACCOUNT_SCORE_1"
	TokenizationDeclineReasonDeviceScore1                   TokenizationDeclineReason = "DEVICE_SCORE_1"
	TokenizationDeclineReasonAllWalletDeclineReasonsPresent TokenizationDeclineReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	TokenizationDeclineReasonWalletRecommendedDecisionRed   TokenizationDeclineReason = "WALLET_RECOMMENDED_DECISION_RED"
	TokenizationDeclineReasonCvcMismatch                    TokenizationDeclineReason = "CVC_MISMATCH"
	TokenizationDeclineReasonCardExpiryMonthMismatch        TokenizationDeclineReason = "CARD_EXPIRY_MONTH_MISMATCH"
	TokenizationDeclineReasonCardExpiryYearMismatch         TokenizationDeclineReason = "CARD_EXPIRY_YEAR_MISMATCH"
	TokenizationDeclineReasonCardInvalidState               TokenizationDeclineReason = "CARD_INVALID_STATE"
	TokenizationDeclineReasonCustomerRedPath                TokenizationDeclineReason = "CUSTOMER_RED_PATH"
	TokenizationDeclineReasonInvalidCustomerResponse        TokenizationDeclineReason = "INVALID_CUSTOMER_RESPONSE"
	TokenizationDeclineReasonNetworkFailure                 TokenizationDeclineReason = "NETWORK_FAILURE"
	TokenizationDeclineReasonGenericDecline                 TokenizationDeclineReason = "GENERIC_DECLINE"
	TokenizationDeclineReasonDigitalCardArtRequired         TokenizationDeclineReason = "DIGITAL_CARD_ART_REQUIRED"
)

func (r TokenizationDeclineReason) IsKnown() bool {
	switch r {
	case TokenizationDeclineReasonAccountScore1, TokenizationDeclineReasonDeviceScore1, TokenizationDeclineReasonAllWalletDeclineReasonsPresent, TokenizationDeclineReasonWalletRecommendedDecisionRed, TokenizationDeclineReasonCvcMismatch, TokenizationDeclineReasonCardExpiryMonthMismatch, TokenizationDeclineReasonCardExpiryYearMismatch, TokenizationDeclineReasonCardInvalidState, TokenizationDeclineReasonCustomerRedPath, TokenizationDeclineReasonInvalidCustomerResponse, TokenizationDeclineReasonNetworkFailure, TokenizationDeclineReasonGenericDecline, TokenizationDeclineReasonDigitalCardArtRequired:
		return true
	}
	return false
}

type TokenizationRuleResult struct {
	// The Auth Rule Token associated with the rule. If this is set to null, then the
	// result was not associated with a customer-configured rule. This may happen in
	// cases where a tokenization is declined or requires TFA due to a
	// Lithic-configured security or compliance rule, for example.
	AuthRuleToken string `json:"auth_rule_token,required,nullable" format:"uuid"`
	// A human-readable explanation outlining the motivation for the rule's result
	Explanation string `json:"explanation,required,nullable"`
	// The name for the rule, if any was configured
	Name string `json:"name,required,nullable"`
	// The result associated with this rule
	Result TokenizationRuleResultResult `json:"result,required"`
	JSON   tokenizationRuleResultJSON   `json:"-"`
}

// tokenizationRuleResultJSON contains the JSON metadata for the struct
// [TokenizationRuleResult]
type tokenizationRuleResultJSON struct {
	AuthRuleToken apijson.Field
	Explanation   apijson.Field
	Name          apijson.Field
	Result        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TokenizationRuleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationRuleResultJSON) RawJSON() string {
	return r.raw
}

// The result associated with this rule
type TokenizationRuleResultResult string

const (
	TokenizationRuleResultResultApproved   TokenizationRuleResultResult = "APPROVED"
	TokenizationRuleResultResultDeclined   TokenizationRuleResultResult = "DECLINED"
	TokenizationRuleResultResultRequireTfa TokenizationRuleResultResult = "REQUIRE_TFA"
	TokenizationRuleResultResultError      TokenizationRuleResultResult = "ERROR"
)

func (r TokenizationRuleResultResult) IsKnown() bool {
	switch r {
	case TokenizationRuleResultResultApproved, TokenizationRuleResultResultDeclined, TokenizationRuleResultResultRequireTfa, TokenizationRuleResultResultError:
		return true
	}
	return false
}

// Reason code for why a tokenization required two-factor authentication
type TokenizationTfaReason string

const (
	TokenizationTfaReasonWalletRecommendedTfa        TokenizationTfaReason = "WALLET_RECOMMENDED_TFA"
	TokenizationTfaReasonSuspiciousActivity          TokenizationTfaReason = "SUSPICIOUS_ACTIVITY"
	TokenizationTfaReasonDeviceRecentlyLost          TokenizationTfaReason = "DEVICE_RECENTLY_LOST"
	TokenizationTfaReasonTooManyRecentAttempts       TokenizationTfaReason = "TOO_MANY_RECENT_ATTEMPTS"
	TokenizationTfaReasonTooManyRecentTokens         TokenizationTfaReason = "TOO_MANY_RECENT_TOKENS"
	TokenizationTfaReasonTooManyDifferentCardholders TokenizationTfaReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	TokenizationTfaReasonOutsideHomeTerritory        TokenizationTfaReason = "OUTSIDE_HOME_TERRITORY"
	TokenizationTfaReasonHasSuspendedTokens          TokenizationTfaReason = "HAS_SUSPENDED_TOKENS"
	TokenizationTfaReasonHighRisk                    TokenizationTfaReason = "HIGH_RISK"
	TokenizationTfaReasonAccountScoreLow             TokenizationTfaReason = "ACCOUNT_SCORE_LOW"
	TokenizationTfaReasonDeviceScoreLow              TokenizationTfaReason = "DEVICE_SCORE_LOW"
	TokenizationTfaReasonCardStateTfa                TokenizationTfaReason = "CARD_STATE_TFA"
	TokenizationTfaReasonHardcodedTfa                TokenizationTfaReason = "HARDCODED_TFA"
	TokenizationTfaReasonCustomerRuleTfa             TokenizationTfaReason = "CUSTOMER_RULE_TFA"
	TokenizationTfaReasonDeviceHostCardEmulation     TokenizationTfaReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r TokenizationTfaReason) IsKnown() bool {
	switch r {
	case TokenizationTfaReasonWalletRecommendedTfa, TokenizationTfaReasonSuspiciousActivity, TokenizationTfaReasonDeviceRecentlyLost, TokenizationTfaReasonTooManyRecentAttempts, TokenizationTfaReasonTooManyRecentTokens, TokenizationTfaReasonTooManyDifferentCardholders, TokenizationTfaReasonOutsideHomeTerritory, TokenizationTfaReasonHasSuspendedTokens, TokenizationTfaReasonHighRisk, TokenizationTfaReasonAccountScoreLow, TokenizationTfaReasonDeviceScoreLow, TokenizationTfaReasonCardStateTfa, TokenizationTfaReasonHardcodedTfa, TokenizationTfaReasonCustomerRuleTfa, TokenizationTfaReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

type TokenizationListParams struct {
	// Filters for tokenizations associated with a specific account.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Filter for tokenizations created after this date.
	Begin param.Field[time.Time] `query:"begin" format:"date"`
	// Filters for tokenizations associated with a specific card.
	CardToken param.Field[string] `query:"card_token" format:"uuid"`
	// Filter for tokenizations created before this date.
	End param.Field[time.Time] `query:"end" format:"date"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Filter for tokenizations by tokenization channel. If this is not specified, only
	// DIGITAL_WALLET tokenizations will be returned.
	TokenizationChannel param.Field[TokenizationListParamsTokenizationChannel] `query:"tokenization_channel"`
}

// URLQuery serializes [TokenizationListParams]'s query parameters as `url.Values`.
func (r TokenizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter for tokenizations by tokenization channel. If this is not specified, only
// DIGITAL_WALLET tokenizations will be returned.
type TokenizationListParamsTokenizationChannel string

const (
	TokenizationListParamsTokenizationChannelDigitalWallet TokenizationListParamsTokenizationChannel = "DIGITAL_WALLET"
	TokenizationListParamsTokenizationChannelMerchant      TokenizationListParamsTokenizationChannel = "MERCHANT"
	TokenizationListParamsTokenizationChannelAll           TokenizationListParamsTokenizationChannel = "ALL"
)

func (r TokenizationListParamsTokenizationChannel) IsKnown() bool {
	switch r {
	case TokenizationListParamsTokenizationChannelDigitalWallet, TokenizationListParamsTokenizationChannelMerchant, TokenizationListParamsTokenizationChannelAll:
		return true
	}
	return false
}

type TokenizationResendActivationCodeParams struct {
	// The communication method that the user has selected to use to receive the
	// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
	// = "EMAIL_TO_CARDHOLDER_ADDRESS"
	ActivationMethodType param.Field[TokenizationResendActivationCodeParamsActivationMethodType] `json:"activation_method_type"`
}

func (r TokenizationResendActivationCodeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The communication method that the user has selected to use to receive the
// authentication code. Supported Values: Sms = "TEXT_TO_CARDHOLDER_NUMBER". Email
// = "EMAIL_TO_CARDHOLDER_ADDRESS"
type TokenizationResendActivationCodeParamsActivationMethodType string

const (
	TokenizationResendActivationCodeParamsActivationMethodTypeEmailToCardholderAddress TokenizationResendActivationCodeParamsActivationMethodType = "EMAIL_TO_CARDHOLDER_ADDRESS"
	TokenizationResendActivationCodeParamsActivationMethodTypeTextToCardholderNumber   TokenizationResendActivationCodeParamsActivationMethodType = "TEXT_TO_CARDHOLDER_NUMBER"
)

func (r TokenizationResendActivationCodeParamsActivationMethodType) IsKnown() bool {
	switch r {
	case TokenizationResendActivationCodeParamsActivationMethodTypeEmailToCardholderAddress, TokenizationResendActivationCodeParamsActivationMethodTypeTextToCardholderNumber:
		return true
	}
	return false
}

type TokenizationSimulateParams struct {
	// The three digit cvv for the card.
	Cvv param.Field[string] `json:"cvv,required"`
	// The expiration date of the card in 'MM/YY' format.
	ExpirationDate param.Field[string] `json:"expiration_date,required"`
	// The sixteen digit card number.
	Pan param.Field[string] `json:"pan,required"`
	// The source of the tokenization request.
	TokenizationSource param.Field[TokenizationSimulateParamsTokenizationSource] `json:"tokenization_source,required"`
	// The account score (1-5) that represents how the Digital Wallet's view on how
	// reputable an end user's account is.
	AccountScore param.Field[int64] `json:"account_score"`
	// The device score (1-5) that represents how the Digital Wallet's view on how
	// reputable an end user's device is.
	DeviceScore param.Field[int64] `json:"device_score"`
	// Optional field to specify the token requestor name for a merchant token
	// simulation. Ignored when tokenization_source is not MERCHANT.
	Entity param.Field[string] `json:"entity"`
	// The decision that the Digital Wallet's recommend
	WalletRecommendedDecision param.Field[TokenizationSimulateParamsWalletRecommendedDecision] `json:"wallet_recommended_decision"`
}

func (r TokenizationSimulateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the tokenization request.
type TokenizationSimulateParamsTokenizationSource string

const (
	TokenizationSimulateParamsTokenizationSourceApplePay   TokenizationSimulateParamsTokenizationSource = "APPLE_PAY"
	TokenizationSimulateParamsTokenizationSourceGoogle     TokenizationSimulateParamsTokenizationSource = "GOOGLE"
	TokenizationSimulateParamsTokenizationSourceSamsungPay TokenizationSimulateParamsTokenizationSource = "SAMSUNG_PAY"
	TokenizationSimulateParamsTokenizationSourceMerchant   TokenizationSimulateParamsTokenizationSource = "MERCHANT"
)

func (r TokenizationSimulateParamsTokenizationSource) IsKnown() bool {
	switch r {
	case TokenizationSimulateParamsTokenizationSourceApplePay, TokenizationSimulateParamsTokenizationSourceGoogle, TokenizationSimulateParamsTokenizationSourceSamsungPay, TokenizationSimulateParamsTokenizationSourceMerchant:
		return true
	}
	return false
}

// The decision that the Digital Wallet's recommend
type TokenizationSimulateParamsWalletRecommendedDecision string

const (
	TokenizationSimulateParamsWalletRecommendedDecisionApproved                        TokenizationSimulateParamsWalletRecommendedDecision = "APPROVED"
	TokenizationSimulateParamsWalletRecommendedDecisionDeclined                        TokenizationSimulateParamsWalletRecommendedDecision = "DECLINED"
	TokenizationSimulateParamsWalletRecommendedDecisionRequireAdditionalAuthentication TokenizationSimulateParamsWalletRecommendedDecision = "REQUIRE_ADDITIONAL_AUTHENTICATION"
)

func (r TokenizationSimulateParamsWalletRecommendedDecision) IsKnown() bool {
	switch r {
	case TokenizationSimulateParamsWalletRecommendedDecisionApproved, TokenizationSimulateParamsWalletRecommendedDecisionDeclined, TokenizationSimulateParamsWalletRecommendedDecisionRequireAdditionalAuthentication:
		return true
	}
	return false
}

type TokenizationUpdateDigitalCardArtParams struct {
	// Specifies the digital card art to be displayed in the user’s digital wallet for
	// a tokenization. This artwork must be approved by the network and configured by
	// Lithic to use. See
	// [Flexible Card Art Guide](https://docs.lithic.com/docs/about-digital-wallets#flexible-card-art).
	DigitalCardArtToken param.Field[string] `json:"digital_card_art_token" format:"uuid"`
}

func (r TokenizationUpdateDigitalCardArtParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
