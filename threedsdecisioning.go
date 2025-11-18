// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"slices"

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

// Card program's response to a 3DS Challenge Request. Challenge Request is emitted
// as a webhook
// [three_ds_authentication.challenge](https://docs.lithic.com/reference/post_three-ds-authentication-challenge)
// and your Card Program needs to be configured with Out of Band (OOB) Challenges
// in order to receive it (see https://docs.lithic.com/docs/3ds-challenge-flow for
// more information).
func (r *ThreeDSDecisioningService) ChallengeResponse(ctx context.Context, body ThreeDSDecisioningChallengeResponseParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	opts = slices.Concat(r.Options, opts)
	path := "v1/three_ds_decisioning/secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate a new 3DS Decisioning HMAC secret key. The old secret key will be
// deactivated 24 hours after a successful request to this endpoint. Make a
// [`GET /three_ds_decisioning/secret`](https://docs.lithic.com/reference/getthreedsdecisioningsecret)
// request to retrieve the new secret key.
func (r *ThreeDSDecisioningService) RotateSecret(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/three_ds_decisioning/secret/rotate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Response from Card Program to a 3DS Authentication challenge
type ChallengeResponseParam struct {
	// Globally unique identifier for 3DS Authentication that resulted in
	// PENDING_CHALLENGE authentication result.
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// Whether the Cardholder has approved or declined the issued Challenge
	ChallengeResponse param.Field[ChallengeResult] `json:"challenge_response,required"`
}

func (r ChallengeResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the Cardholder has approved or declined the issued Challenge
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

type ThreeDSDecisioningChallengeResponseParams struct {
	// Response from Card Program to a 3DS Authentication challenge
	ChallengeResponse ChallengeResponseParam `json:"challenge_response,required"`
}

func (r ThreeDSDecisioningChallengeResponseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ChallengeResponse)
}
