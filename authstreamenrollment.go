// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"slices"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// AuthStreamEnrollmentService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthStreamEnrollmentService] method instead.
type AuthStreamEnrollmentService struct {
	Options []option.RequestOption
}

// NewAuthStreamEnrollmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAuthStreamEnrollmentService(opts ...option.RequestOption) (r *AuthStreamEnrollmentService) {
	r = &AuthStreamEnrollmentService{}
	r.Options = opts
	return
}

// Retrieve the ASA HMAC secret key. If one does not exist for your program yet,
// calling this endpoint will create one for you. The headers (which you can use to
// verify webhooks) will begin appearing shortly after calling this endpoint for
// the first time. See
// [this page](https://docs.lithic.com/docs/auth-stream-access-asa#asa-webhook-verification)
// for more detail about verifying ASA webhooks.
func (r *AuthStreamEnrollmentService) GetSecret(ctx context.Context, opts ...option.RequestOption) (res *AuthStreamSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auth_stream/secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate a new ASA HMAC secret key. The old ASA HMAC secret key will be
// deactivated 24 hours after a successful request to this endpoint. Make a
// [`GET /auth_stream/secret`](https://docs.lithic.com/reference/getauthstreamsecret)
// request to retrieve the new secret key.
func (r *AuthStreamEnrollmentService) RotateSecret(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/auth_stream/secret/rotate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type AuthStreamSecret struct {
	// The shared HMAC ASA secret
	Secret string               `json:"secret"`
	JSON   authStreamSecretJSON `json:"-"`
}

// authStreamSecretJSON contains the JSON metadata for the struct
// [AuthStreamSecret]
type authStreamSecretJSON struct {
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthStreamSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authStreamSecretJSON) RawJSON() string {
	return r.raw
}
