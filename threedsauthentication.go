// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// ThreeDSAuthenticationService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreeDSAuthenticationService] method instead.
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
func (r *ThreeDSAuthenticationService) Get(ctx context.Context, threeDSAuthenticationToken string, opts ...option.RequestOption) (res *ThreeDSAuthentication, err error) {
	opts = slices.Concat(r.Options, opts)
	if threeDSAuthenticationToken == "" {
		err = errors.New("missing required three_ds_authentication_token parameter")
		return
	}
	path := fmt.Sprintf("v1/three_ds_authentication/%s", threeDSAuthenticationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Simulates a 3DS authentication request from the payment network as if it came
// from an ACS. If you're configured for 3DS Customer Decisioning, simulating
// authentications requires your customer decisioning endpoint to be set up
// properly (respond with a valid JSON). If the authentication decision is to
// challenge, ensure that the account holder associated with the card transaction
// has a valid phone number configured to receive the OTP code via SMS.
func (r *ThreeDSAuthenticationService) Simulate(ctx context.Context, body ThreeDSAuthenticationSimulateParams, opts ...option.RequestOption) (res *ThreeDSAuthenticationSimulateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/three_ds_authentication/simulate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Endpoint for simulating entering OTP into 3DS Challenge UI. A call to
// [/v1/three_ds_authentication/simulate](https://docs.lithic.com/reference/postsimulateauthentication)
// that resulted in triggered SMS-OTP challenge must precede. Only a single attempt
// is supported; upon entering OTP, the challenge is either approved or declined.
func (r *ThreeDSAuthenticationService) SimulateOtpEntry(ctx context.Context, body ThreeDSAuthenticationSimulateOtpEntryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/three_ds_decisioning/simulate/enter_otp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ThreeDSAuthenticationSimulateResponse struct {
	// Globally unique identifier for the 3DS authentication.
	Token string                                    `json:"token" format:"uuid"`
	JSON  threeDSAuthenticationSimulateResponseJSON `json:"-"`
}

// threeDSAuthenticationSimulateResponseJSON contains the JSON metadata for the
// struct [ThreeDSAuthenticationSimulateResponse]
type threeDSAuthenticationSimulateResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSAuthenticationSimulateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSAuthenticationSimulateResponseJSON) RawJSON() string {
	return r.raw
}

type ThreeDSAuthenticationSimulateParams struct {
	// Merchant information for the simulated transaction
	Merchant param.Field[ThreeDSAuthenticationSimulateParamsMerchant] `json:"merchant,required"`
	// Sixteen digit card number.
	Pan param.Field[string] `json:"pan,required"`
	// Transaction details for the simulation
	Transaction param.Field[ThreeDSAuthenticationSimulateParamsTransaction] `json:"transaction,required"`
	// When set will use the following values as part of the Simulated Authentication.
	// When not set defaults to MATCH
	CardExpiryCheck param.Field[ThreeDSAuthenticationSimulateParamsCardExpiryCheck] `json:"card_expiry_check"`
}

func (r ThreeDSAuthenticationSimulateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant information for the simulated transaction
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
	// Merchant descriptor, corresponds to `descriptor` in authorization. If CHALLENGE
	// keyword is included, Lithic will trigger a challenge.
	Name param.Field[string] `json:"name,required"`
}

func (r ThreeDSAuthenticationSimulateParamsMerchant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Transaction details for the simulation
type ThreeDSAuthenticationSimulateParamsTransaction struct {
	// Amount (in cents) to authenticate.
	Amount param.Field[int64] `json:"amount,required"`
	// 3-character alphabetic ISO 4217 currency code.
	Currency param.Field[string] `json:"currency,required"`
}

func (r ThreeDSAuthenticationSimulateParamsTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When set will use the following values as part of the Simulated Authentication.
// When not set defaults to MATCH
type ThreeDSAuthenticationSimulateParamsCardExpiryCheck string

const (
	ThreeDSAuthenticationSimulateParamsCardExpiryCheckMatch      ThreeDSAuthenticationSimulateParamsCardExpiryCheck = "MATCH"
	ThreeDSAuthenticationSimulateParamsCardExpiryCheckMismatch   ThreeDSAuthenticationSimulateParamsCardExpiryCheck = "MISMATCH"
	ThreeDSAuthenticationSimulateParamsCardExpiryCheckNotPresent ThreeDSAuthenticationSimulateParamsCardExpiryCheck = "NOT_PRESENT"
)

func (r ThreeDSAuthenticationSimulateParamsCardExpiryCheck) IsKnown() bool {
	switch r {
	case ThreeDSAuthenticationSimulateParamsCardExpiryCheckMatch, ThreeDSAuthenticationSimulateParamsCardExpiryCheckMismatch, ThreeDSAuthenticationSimulateParamsCardExpiryCheckNotPresent:
		return true
	}
	return false
}

type ThreeDSAuthenticationSimulateOtpEntryParams struct {
	// A unique token returned as part of a /v1/three_ds_authentication/simulate call
	// that resulted in PENDING_CHALLENGE authentication result.
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// The OTP entered by the cardholder
	Otp param.Field[string] `json:"otp,required"`
}

func (r ThreeDSAuthenticationSimulateOtpEntryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
