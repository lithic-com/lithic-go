package requests

import (
	"net/url"

	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type AuthRuleUpdateParams struct {
	// Array of merchant category codes for which the Auth Rule will permit
	// transactions. Note that only this field or `blocked_mcc` can be used for a given
	// Auth Rule.
	AllowedMcc field.Field[[]string] `json:"allowed_mcc"`
	// Array of merchant category codes for which the Auth Rule will automatically
	// decline transactions. Note that only this field or `allowed_mcc` can be used for
	// a given Auth Rule.
	BlockedMcc field.Field[[]string] `json:"blocked_mcc"`
	// Array of country codes for which the Auth Rule will permit transactions. Note
	// that only this field or `blocked_countries` can be used for a given Auth Rule.
	AllowedCountries field.Field[[]string] `json:"allowed_countries"`
	// Array of country codes for which the Auth Rule will automatically decline
	// transactions. Note that only this field or `allowed_countries` can be used for a
	// given Auth Rule.
	BlockedCountries field.Field[[]string] `json:"blocked_countries"`
	// Address verification to confirm that postal code entered at point of transaction
	// (if applicable) matches the postal code on file for a given card.
	AvsType field.Field[AuthRuleUpdateParamsAvsType] `json:"avs_type"`
}

// MarshalJSON serializes AuthRuleUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r AuthRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type AuthRuleUpdateParamsAvsType string

const (
	AuthRuleUpdateParamsAvsTypeZipOnly AuthRuleUpdateParamsAvsType = "ZIP_ONLY"
)

type AuthRuleListParams struct {
	// Page (for pagination).
	Page field.Field[int64] `query:"page"`
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
}

// URLQuery serializes AuthRuleListParams into a url.Values of the query parameters
// associated with this value
func (r AuthRuleListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type AuthRuleApplyParams struct {
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens field.Field[[]string] `json:"card_tokens"`
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens field.Field[[]string] `json:"account_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel field.Field[bool] `json:"program_level"`
}

// MarshalJSON serializes AuthRuleApplyParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r AuthRuleApplyParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type AuthRuleRemoveParams struct {
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens field.Field[[]string] `json:"card_tokens"`
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens field.Field[[]string] `json:"account_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel field.Field[bool] `json:"program_level"`
}

// MarshalJSON serializes AuthRuleRemoveParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r AuthRuleRemoveParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}
