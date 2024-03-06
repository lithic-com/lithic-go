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

// TokenizationService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTokenizationService] method
// instead.
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
func (r *TokenizationService) Get(ctx context.Context, tokenizationToken string, opts ...option.RequestOption) (res *TokenizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("tokenizations/%s", tokenizationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List card tokenizations
func (r *TokenizationService) List(ctx context.Context, query TokenizationListParams, opts ...option.RequestOption) (res *shared.CursorPage[Tokenization], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "tokenizations"
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
func (r *TokenizationService) ListAutoPaging(ctx context.Context, query TokenizationListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[Tokenization] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to simulate a card's tokenization in the Digital Wallet
// and merchant tokenization ecosystem.
func (r *TokenizationService) Simulate(ctx context.Context, body TokenizationSimulateParams, opts ...option.RequestOption) (res *TokenizationSimulateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/tokenizations"
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
	// The status of the tokenization request
	Status TokenizationStatus `json:"status,required"`
	// The entity that is requested the tokenization. Represents a Digital Wallet.
	TokenRequestorName TokenizationTokenRequestorName `json:"token_requestor_name,required"`
	// The network's unique reference for the tokenization.
	TokenUniqueReference string `json:"token_unique_reference,required"`
	// Latest date and time when the tokenization was updated. UTC time zone.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// A list of events related to the tokenization.
	Events []TokenizationEvent `json:"events"`
	JSON   tokenizationJSON    `json:"-"`
}

// tokenizationJSON contains the JSON metadata for the struct [Tokenization]
type tokenizationJSON struct {
	Token                apijson.Field
	AccountToken         apijson.Field
	CardToken            apijson.Field
	CreatedAt            apijson.Field
	Status               apijson.Field
	TokenRequestorName   apijson.Field
	TokenUniqueReference apijson.Field
	UpdatedAt            apijson.Field
	Events               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
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

// The entity that is requested the tokenization. Represents a Digital Wallet.
type TokenizationTokenRequestorName string

const (
	TokenizationTokenRequestorNameAmazonOne    TokenizationTokenRequestorName = "AMAZON_ONE"
	TokenizationTokenRequestorNameAndroidPay   TokenizationTokenRequestorName = "ANDROID_PAY"
	TokenizationTokenRequestorNameApplePay     TokenizationTokenRequestorName = "APPLE_PAY"
	TokenizationTokenRequestorNameFitbitPay    TokenizationTokenRequestorName = "FITBIT_PAY"
	TokenizationTokenRequestorNameGarminPay    TokenizationTokenRequestorName = "GARMIN_PAY"
	TokenizationTokenRequestorNameMicrosoftPay TokenizationTokenRequestorName = "MICROSOFT_PAY"
	TokenizationTokenRequestorNameSamsungPay   TokenizationTokenRequestorName = "SAMSUNG_PAY"
	TokenizationTokenRequestorNameUnknown      TokenizationTokenRequestorName = "UNKNOWN"
	TokenizationTokenRequestorNameVisaCheckout TokenizationTokenRequestorName = "VISA_CHECKOUT"
)

type TokenizationEvent struct {
	// Globally unique identifier for a Tokenization Event
	Token string `json:"token" format:"uuid"`
	// Date and time when the tokenization event first occurred. UTC time zone.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Enum representing the result of the tokenization event
	Result TokenizationEventsResult `json:"result"`
	// Enum representing the type of tokenization event that occurred
	Type TokenizationEventsType `json:"type"`
	JSON tokenizationEventJSON  `json:"-"`
}

// tokenizationEventJSON contains the JSON metadata for the struct
// [TokenizationEvent]
type tokenizationEventJSON struct {
	Token       apijson.Field
	CreatedAt   apijson.Field
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	TokenizationEventsResultTokenInactive                   TokenizationEventsResult = "TOKEN_INACTIVE"
	TokenizationEventsResultTokenStateUnknown               TokenizationEventsResult = "TOKEN_STATE_UNKNOWN"
	TokenizationEventsResultTokenSuspended                  TokenizationEventsResult = "TOKEN_SUSPENDED"
	TokenizationEventsResultTokenUpdated                    TokenizationEventsResult = "TOKEN_UPDATED"
)

// Enum representing the type of tokenization event that occurred
type TokenizationEventsType string

const (
	TokenizationEventsTypeTokenization2Fa              TokenizationEventsType = "TOKENIZATION_2FA"
	TokenizationEventsTypeTokenizationAuthorization    TokenizationEventsType = "TOKENIZATION_AUTHORIZATION"
	TokenizationEventsTypeTokenizationDecisioning      TokenizationEventsType = "TOKENIZATION_DECISIONING"
	TokenizationEventsTypeTokenizationEligibilityCheck TokenizationEventsType = "TOKENIZATION_ELIGIBILITY_CHECK"
	TokenizationEventsTypeTokenizationUpdated          TokenizationEventsType = "TOKENIZATION_UPDATED"
)

type TokenizationGetResponse struct {
	Data Tokenization                `json:"data"`
	JSON tokenizationGetResponseJSON `json:"-"`
}

// tokenizationGetResponseJSON contains the JSON metadata for the struct
// [TokenizationGetResponse]
type tokenizationGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationGetResponseJSON) RawJSON() string {
	return r.raw
}

type TokenizationSimulateResponse struct {
	Data []Tokenization                   `json:"data"`
	JSON tokenizationSimulateResponseJSON `json:"-"`
}

// tokenizationSimulateResponseJSON contains the JSON metadata for the struct
// [TokenizationSimulateResponse]
type tokenizationSimulateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenizationSimulateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenizationSimulateResponseJSON) RawJSON() string {
	return r.raw
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
}

// URLQuery serializes [TokenizationListParams]'s query parameters as `url.Values`.
func (r TokenizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
)

// The decision that the Digital Wallet's recommend
type TokenizationSimulateParamsWalletRecommendedDecision string

const (
	TokenizationSimulateParamsWalletRecommendedDecisionApproved                        TokenizationSimulateParamsWalletRecommendedDecision = "APPROVED"
	TokenizationSimulateParamsWalletRecommendedDecisionDeclined                        TokenizationSimulateParamsWalletRecommendedDecision = "DECLINED"
	TokenizationSimulateParamsWalletRecommendedDecisionRequireAdditionalAuthentication TokenizationSimulateParamsWalletRecommendedDecision = "REQUIRE_ADDITIONAL_AUTHENTICATION"
)
