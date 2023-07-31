// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// ThreeDSDescisioningService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewThreeDSDescisioningService]
// method instead.
type ThreeDSDescisioningService struct {
	Options []option.RequestOption
}

// NewThreeDSDescisioningService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreeDSDescisioningService(opts ...option.RequestOption) (r *ThreeDSDescisioningService) {
	r = &ThreeDSDescisioningService{}
	r.Options = opts
	return
}

// Retrieve the 3DS Decisioning HMAC secret key. If one does not exist for your
// program yet, calling this endpoint will create one for you. The headers (which
// you can use to verify 3DS Decisioning requests) will begin appearing shortly
// after calling this endpoint for the first time. See
// [this page](https://docs.lithic.com/docs/3ds-decisioning#3ds-decisioning-hmac-secrets)
// for more detail about verifying 3DS Decisioning requests.
func (r *ThreeDSDescisioningService) GetSecret(ctx context.Context, opts ...option.RequestOption) (res *ThreeDSDescisioningGetSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "three_ds_decisioning/secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate a new 3DS Decisioning HMAC secret key. The old secret key will be
// deactivated 24 hours after a successful request to this endpoint. Make a
// [`GET /three_ds_decisioning/secret`](https://docs.lithic.com/reference/getthreedsdecisioningsecret)
// request to retrieve the new secret key.
func (r *ThreeDSDescisioningService) RotateSecret(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "three_ds_decisioning/secret/rotate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type ThreeDSDescisioningGetSecretResponse struct {
	// The 3DS Decisioning HMAC secret
	Secret string `json:"secret"`
	JSON   threeDSDescisioningGetSecretResponseJSON
}

// threeDSDescisioningGetSecretResponseJSON contains the JSON metadata for the
// struct [ThreeDSDescisioningGetSecretResponse]
type threeDSDescisioningGetSecretResponseJSON struct {
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreeDSDescisioningGetSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
