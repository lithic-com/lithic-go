// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// ThreeDSDecisioningService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreeDSDecisioningService] method instead.
type ThreeDSDecisioningService struct {
	Options []option.RequestOption
}

// NewThreeDSDecisioningService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreeDSDecisioningService(opts ...option.RequestOption) (r *ThreeDSDecisioningService) {
	r = &ThreeDSDecisioningService{}
	r.Options = opts
	return
}

// Card program's response to a 3DS Challenge Request (CReq)
func (r *ThreeDSDecisioningService) ChallengeResponse(ctx context.Context, body ThreeDSDecisioningChallengeResponseParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/three_ds_decisioning/challenge_response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve the 3DS Decisioning HMAC secret key. If one does not exist for your
// program yet, calling this endpoint will create one for you. The headers (which
// you can use to verify 3DS Decisioning requests) will begin appearing shortly
// after calling this endpoint for the first time. See
// [this page](https://docs.lithic.com/docs/3ds-decisioning#3ds-decisioning-hmac-secrets)
// for more detail about verifying 3DS Decisioning requests.
func (r *ThreeDSDecisioningService) GetSecret(ctx context.Context, opts ...option.RequestOption) (res *ThreeDSDecisioningGetSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/three_ds_decisioning/secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate a new 3DS Decisioning HMAC secret key. The old secret key will be
// deactivated 24 hours after a successful request to this endpoint. Make a
// [`GET /three_ds_decisioning/secret`](https://docs.lithic.com/reference/getthreedsdecisioningsecret)
// request to retrieve the new secret key.
func (r *ThreeDSDecisioningService) RotateSecret(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/three_ds_decisioning/secret/rotate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Simulates a 3DS authentication challenge request from the payment network as if
// it came from an ACS. Requires being configured for 3DS Customer Decisioning, and
// enrolled with Lithic's Challenge solution.
func (r *ThreeDSDecisioningService) SimulateChallenge(ctx context.Context, body ThreeDSDecisioningSimulateChallengeParams, opts ...option.RequestOption) (res *ThreeDSDecisioningSimulateChallengeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/three_ds_decisioning/simulate/challenge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Endpoint for responding to a 3DS Challenge initiated by a call to
// /v1/three_ds_decisioning/simulate/challenge
func (r *ThreeDSDecisioningService) SimulateChallengeResponse(ctx context.Context, body ThreeDSDecisioningSimulateChallengeResponseParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/three_ds_decisioning/simulate/challenge_response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ChallengeResponseParam struct {
	// Globally unique identifier for the 3DS authentication. This token is sent as
	// part of the initial 3DS Decisioning Request and as part of the 3DS Challenge
	// Event in the [ThreeDSAuthentication](#/components/schemas/ThreeDSAuthentication)
	// object
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// Whether the Cardholder has Approved or Declined the issued Challenge
	ChallengeResponse param.Field[ChallengeResult] `json:"challenge_response,required"`
}

func (r ChallengeResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the Cardholder has Approved or Declined the issued Challenge
type ChallengeResult string

const (
	ChallengeResultApprove           ChallengeResult = "APPROVE"
	ChallengeResultDeclineByCustomer ChallengeResult = "DECLINE_BY_CUSTOMER"
)

func (r ChallengeResult) IsKnown() bool {
	switch r {
	case ChallengeResultApprove, ChallengeResultDeclineByCustomer:
		return true
	}
	return false
}

type ThreeDSDecisioningGetSecretResponse struct {
	// The 3DS Decisioning HMAC secret
	Secret string                                  `json:"secret"`
	JSON   threeDSDecisioningGetSecretResponseJSON `json:"-"`
}

// threeDSDecisioningGetSecretResponseJSON contains the JSON metadata for the
// struct [ThreeDSDecisioningGetSecretResponse]
type threeDSDecisioningGetSecretResponseJSON struct {
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSDecisioningGetSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSDecisioningGetSecretResponseJSON) RawJSON() string {
	return r.raw
}

type ThreeDSDecisioningSimulateChallengeResponse struct {
	// A unique token to reference this transaction with later calls to void or clear
	// the authorization. This token is used in
	// /v1/three_ds_decisioning/simulate/challenge_response to Approve or Decline the
	// authentication
	Token string                                          `json:"token" format:"uuid"`
	JSON  threeDSDecisioningSimulateChallengeResponseJSON `json:"-"`
}

// threeDSDecisioningSimulateChallengeResponseJSON contains the JSON metadata for
// the struct [ThreeDSDecisioningSimulateChallengeResponse]
type threeDSDecisioningSimulateChallengeResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSDecisioningSimulateChallengeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threeDSDecisioningSimulateChallengeResponseJSON) RawJSON() string {
	return r.raw
}

type ThreeDSDecisioningChallengeResponseParams struct {
	ChallengeResponse ChallengeResponseParam `json:"challenge_response,required"`
}

func (r ThreeDSDecisioningChallengeResponseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ChallengeResponse)
}

type ThreeDSDecisioningSimulateChallengeParams struct {
	// A unique token returned as part of a /v1/three_ds_authentication/simulate call
	// that responded with a CHALLENGE_REQUESTED status.
	Token param.Field[string] `json:"token" format:"uuid"`
}

func (r ThreeDSDecisioningSimulateChallengeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreeDSDecisioningSimulateChallengeResponseParams struct {
	ChallengeResponse ChallengeResponseParam `json:"challenge_response,required"`
}

func (r ThreeDSDecisioningSimulateChallengeResponseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ChallengeResponse)
}
