// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"net/http"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
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

// This endpoint is used to simulate a card's tokenization in the Digital Wallet
// and merchant tokenization ecosystem.
func (r *TokenizationService) Simulate(ctx context.Context, body TokenizationSimulateParams, opts ...option.RequestOption) (res *TokenizationSimulateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/tokenizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Tokenization struct {
	// A fixed-width 23-digit numeric identifier for the Transaction that may be set if
	// the transaction originated from the Mastercard network. This number may be used
	// for dispute tracking.
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
	JSON      tokenizationJSON
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
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Tokenization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the tokenization request
type TokenizationStatus string

const (
	TokenizationStatusApproved                        TokenizationStatus = "APPROVED"
	TokenizationStatusDeclined                        TokenizationStatus = "DECLINED"
	TokenizationStatusRequireAdditionalAuthentication TokenizationStatus = "REQUIRE_ADDITIONAL_AUTHENTICATION"
)

// The entity that is requested the tokenization. Represents a Digital Wallet.
type TokenizationTokenRequestorName string

const (
	TokenizationTokenRequestorNameApplePay   TokenizationTokenRequestorName = "APPLE_PAY"
	TokenizationTokenRequestorNameGoogle     TokenizationTokenRequestorName = "GOOGLE"
	TokenizationTokenRequestorNameSamsungPay TokenizationTokenRequestorName = "SAMSUNG_PAY"
)

type TokenizationSimulateResponse struct {
	Data []Tokenization `json:"data"`
	JSON tokenizationSimulateResponseJSON
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
