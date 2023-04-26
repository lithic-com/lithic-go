package responses

import (
	apijson "github.com/lithic-com/lithic-go/core/json"
)

type TokenizationSecret struct {
	// The Tokenization Decisioning HMAC secret
	Secret string `json:"secret"`
	JSON   TokenizationSecretJSON
}

type TokenizationSecretJSON struct {
	Secret apijson.Metadata
	Raw    []byte
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
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TokenizationDecisioningRotateSecretResponse using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TokenizationDecisioningRotateSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
