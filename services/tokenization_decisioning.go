package services

import (
	"context"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/responses"
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
func (r *TokenizationDecisioningService) GetSecret(ctx context.Context, opts ...option.RequestOption) (res *responses.TokenizationSecret, err error) {
	opts = append(r.Options[:], opts...)
	path := "tokenization_decisioning/secret"
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Generate a new Tokenization Decisioning secret key. The old Tokenization
// Decisioning secret key will be deactivated 24 hours after a successful request
// to this endpoint.
func (r *TokenizationDecisioningService) RotateSecret(ctx context.Context, opts ...option.RequestOption) (res *responses.TokenizationDecisioningRotateSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "tokenization_decisioning/secret/rotate"
	err = option.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}
