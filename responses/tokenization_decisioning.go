package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type TokenizationSecret struct {
	// The Tokenization Decisioning HMAC secret
	Secret string `json:"secret"`
	JSON   TokenizationSecretJSON
}

type TokenizationSecretJSON struct {
	Secret pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TokenizationSecret using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TokenizationSecret) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TokenizationDecisioningRotateSecretResponse struct {
	// The new Tokenization Decisioning HMAC secret
	Secret string `json:"secret"`
	JSON   TokenizationDecisioningRotateSecretResponseJSON
}

type TokenizationDecisioningRotateSecretResponseJSON struct {
	Secret pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TokenizationDecisioningRotateSecretResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TokenizationDecisioningRotateSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
