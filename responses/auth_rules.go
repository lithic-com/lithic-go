package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/pagination"
)

type AuthRule struct {
	// Merchant category codes for which the Auth Rule permits transactions.
	AllowedMcc []string `json:"allowed_mcc"`
	// Merchant category codes for which the Auth Rule automatically declines
	// transactions.
	BlockedMcc []string `json:"blocked_mcc"`
	// Countries in which the Auth Rule permits transactions. Note that Lithic
	// maintains a list of countries in which all transactions are blocked; "allowing"
	// those countries in an Auth Rule does not override the Lithic-wide restrictions.
	AllowedCountries []string `json:"allowed_countries"`
	// Countries in which the Auth Rule automatically declines transactions.
	BlockedCountries []string `json:"blocked_countries"`
	// Address verification to confirm that postal code entered at point of transaction
	// (if applicable) matches the postal code on file for a given card. Since this
	// check is performed against the address submitted via the Enroll Consumer
	// endpoint, it should only be used in cases where card users are enrolled with
	// their own accounts. Available values:
	//
	//   - `ZIP_ONLY` - AVS check is performed to confirm ZIP code entered at point of
	//     transaction (if applicable) matches address on file.
	AvsType AuthRuleAvsType `json:"avs_type"`
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens []string `json:"account_tokens"`
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens []string `json:"card_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel bool `json:"program_level"`
	JSON         AuthRuleJSON
}

type AuthRuleJSON struct {
	AllowedMcc       pjson.Metadata
	BlockedMcc       pjson.Metadata
	AllowedCountries pjson.Metadata
	BlockedCountries pjson.Metadata
	AvsType          pjson.Metadata
	AccountTokens    pjson.Metadata
	CardTokens       pjson.Metadata
	ProgramLevel     pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthRule using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *AuthRule) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AuthRuleAvsType string

const (
	AuthRuleAvsTypeZipOnly AuthRuleAvsType = "ZIP_ONLY"
)

type AuthRuleCreateResponse struct {
	Data AuthRule `json:"data"`
	JSON AuthRuleCreateResponseJSON
}

type AuthRuleCreateResponseJSON struct {
	Data   pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthRuleCreateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthRuleCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AuthRuleRetrieveResponse struct {
	Data []AuthRule `json:"data"`
	JSON AuthRuleRetrieveResponseJSON
}

type AuthRuleRetrieveResponseJSON struct {
	Data   pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthRuleRetrieveResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AuthRuleRetrieveResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AuthRuleUpdateResponse struct {
	Data AuthRule `json:"data"`
	JSON AuthRuleUpdateResponseJSON
}

type AuthRuleUpdateResponseJSON struct {
	Data   pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthRuleUpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AuthRuleApplyResponse struct {
	Data AuthRule `json:"data"`
	JSON AuthRuleApplyResponseJSON
}

type AuthRuleApplyResponseJSON struct {
	Data   pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthRuleApplyResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthRuleApplyResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AuthRuleRemoveResponse struct {
	AccountTokens          []string `json:"account_tokens"`
	CardTokens             []string `json:"card_tokens"`
	PreviousAuthRuleTokens []string `json:"previous_auth_rule_tokens"`
	ProgramLevel           bool     `json:"program_level"`
	JSON                   AuthRuleRemoveResponseJSON
}

type AuthRuleRemoveResponseJSON struct {
	AccountTokens          pjson.Metadata
	CardTokens             pjson.Metadata
	PreviousAuthRuleTokens pjson.Metadata
	ProgramLevel           pjson.Metadata
	Raw                    []byte
	Extras                 map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthRuleRemoveResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthRuleRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AuthRuleListResponse struct {
	Data []AuthRule `json:"data"`
	// Total number of entries.
	TotalEntries string `json:"total_entries"`
	// Total number of pages
	TotalPages int64 `json:"total_pages"`
	// Page number.
	Page int64 `json:"page"`
	JSON AuthRuleListResponseJSON
}

type AuthRuleListResponseJSON struct {
	Data         pjson.Metadata
	TotalEntries pjson.Metadata
	TotalPages   pjson.Metadata
	Page         pjson.Metadata
	Raw          []byte
	Extras       map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthRuleListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AuthRulesPage struct {
	*pagination.Page[AuthRule]
}

func (r *AuthRulesPage) AuthRule() *AuthRule {
	return r.Current()
}

func (r *AuthRulesPage) NextPage() (*AuthRulesPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &AuthRulesPage{page}, nil
	}
}
