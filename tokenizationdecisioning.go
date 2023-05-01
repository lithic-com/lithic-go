package lithic

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

type TokenizationDecisioningService struct {
	Options []option.RequestOption
}

func NewTokenizationDecisioningService(opts ...option.RequestOption) (r *TokenizationDecisioningService) {
	r = &TokenizationDecisioningService{}
	r.Options = opts
	return
}

// Retrieve the Tokenization Decisioning secret key. If one does not exist your
// program yet, calling this endpoint will create one for you. The headers of the
// Tokenization Decisioning request will contain a hmac signature which you can use
// to verify requests originate from Lithic. See
// [this page](https://docs.lithic.com/docs/events-api#verifying-webhooks) for more
// detail about verifying Tokenization Decisioning requests.
func (r *TokenizationDecisioningService) GetSecret(ctx context.Context, opts ...option.RequestOption) (res *TokenizationSecret, err error) {
	opts = append(r.Options[:], opts...)
	path := "tokenization_decisioning/secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate a new Tokenization Decisioning secret key. The old Tokenization
// Decisioning secret key will be deactivated 24 hours after a successful request
// to this endpoint.
func (r *TokenizationDecisioningService) RotateSecret(ctx context.Context, opts ...option.RequestOption) (res *TokenizationDecisioningRotateSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "tokenization_decisioning/secret/rotate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type TokenizationSecret struct {
	// The Tokenization Decisioning HMAC secret
	Secret string `json:"secret"`
	JSON   TokenizationSecretJSON
}

type TokenizationSecretJSON struct {
	Secret apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TokenizationSecret using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TokenizationSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TokenizationDecisioningRotateSecretResponse struct {
	// The new Tokenization Decisioning HMAC secret
	Secret string `json:"secret"`
	JSON   TokenizationDecisioningRotateSecretResponseJSON
}

type TokenizationDecisioningRotateSecretResponseJSON struct {
	Secret apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TokenizationDecisioningRotateSecretResponse using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TokenizationDecisioningRotateSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
