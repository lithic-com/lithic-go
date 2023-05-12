package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// AuthRuleService contains methods and other services that help with interacting
// with the lithic API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAuthRuleService] method instead.
type AuthRuleService struct {
	Options []option.RequestOption
}

// NewAuthRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthRuleService(opts ...option.RequestOption) (r *AuthRuleService) {
	r = &AuthRuleService{}
	r.Options = opts
	return
}

// Creates an authorization rule (Auth Rule) and applies it at the program,
// account, or card level.
func (r *AuthRuleService) New(ctx context.Context, body AuthRuleNewParams, opts ...option.RequestOption) (res *AuthRuleCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detail the properties and entities (program, accounts, and cards) associated
// with an existing authorization rule (Auth Rule).
func (r *AuthRuleService) Get(ctx context.Context, auth_rule_token string, opts ...option.RequestOption) (res *AuthRuleRetrieveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s", auth_rule_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the properties associated with an existing authorization rule (Auth
// Rule).
func (r *AuthRuleService) Update(ctx context.Context, auth_rule_token string, body AuthRuleUpdateParams, opts ...option.RequestOption) (res *AuthRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s", auth_rule_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Return all of the Auth Rules under the program.
func (r *AuthRuleService) List(ctx context.Context, query AuthRuleListParams, opts ...option.RequestOption) (res *shared.Page[AuthRule], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "auth_rules"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Return all of the Auth Rules under the program.
func (r *AuthRuleService) ListAutoPaging(ctx context.Context, query AuthRuleListParams, opts ...option.RequestOption) *shared.PageAutoPager[AuthRule] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Applies an existing authorization rule (Auth Rule) to an program, account, or
// card level.
func (r *AuthRuleService) Apply(ctx context.Context, auth_rule_token string, body AuthRuleApplyParams, opts ...option.RequestOption) (res *AuthRuleApplyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s/apply", auth_rule_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove an existing authorization rule (Auth Rule) from an program, account, or
// card-level.
func (r *AuthRuleService) Remove(ctx context.Context, body AuthRuleRemoveParams, opts ...option.RequestOption) (res *AuthRuleRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth_rules/remove"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AuthRule struct {
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Indicates whether the Auth Rule is ACTIVE or INACTIVE
	State AuthRuleState `json:"state"`
	// Identifier for the Auth Rule(s) that a new Auth Rule replaced; will be returned
	// only if an Auth Rule is applied to entities that previously already had one
	// applied.
	PreviousAuthRuleTokens []string `json:"previous_auth_rule_tokens"`
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
	JSON         authRuleJSON
}

// authRuleJSON contains the JSON metadata for the struct [AuthRule]
type authRuleJSON struct {
	Token                  apijson.Field
	State                  apijson.Field
	PreviousAuthRuleTokens apijson.Field
	AllowedMcc             apijson.Field
	BlockedMcc             apijson.Field
	AllowedCountries       apijson.Field
	BlockedCountries       apijson.Field
	AvsType                apijson.Field
	AccountTokens          apijson.Field
	CardTokens             apijson.Field
	ProgramLevel           apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AuthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleState string

const (
	AuthRuleStateActive   AuthRuleState = "ACTIVE"
	AuthRuleStateInactive AuthRuleState = "INACTIVE"
)

type AuthRuleAvsType string

const (
	AuthRuleAvsTypeZipOnly AuthRuleAvsType = "ZIP_ONLY"
)

type AuthRuleCreateResponse struct {
	Data AuthRule `json:"data"`
	JSON authRuleCreateResponseJSON
}

// authRuleCreateResponseJSON contains the JSON metadata for the struct
// [AuthRuleCreateResponse]
type authRuleCreateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleRetrieveResponse struct {
	Data []AuthRule `json:"data"`
	JSON authRuleRetrieveResponseJSON
}

// authRuleRetrieveResponseJSON contains the JSON metadata for the struct
// [AuthRuleRetrieveResponse]
type authRuleRetrieveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleRetrieveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleUpdateResponse struct {
	Data AuthRule `json:"data"`
	JSON authRuleUpdateResponseJSON
}

// authRuleUpdateResponseJSON contains the JSON metadata for the struct
// [AuthRuleUpdateResponse]
type authRuleUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleApplyResponse struct {
	Data AuthRule `json:"data"`
	JSON authRuleApplyResponseJSON
}

// authRuleApplyResponseJSON contains the JSON metadata for the struct
// [AuthRuleApplyResponse]
type authRuleApplyResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleApplyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleRemoveResponse struct {
	AccountTokens          []string `json:"account_tokens"`
	CardTokens             []string `json:"card_tokens"`
	PreviousAuthRuleTokens []string `json:"previous_auth_rule_tokens"`
	ProgramLevel           bool     `json:"program_level"`
	JSON                   authRuleRemoveResponseJSON
}

// authRuleRemoveResponseJSON contains the JSON metadata for the struct
// [AuthRuleRemoveResponse]
type authRuleRemoveResponseJSON struct {
	AccountTokens          apijson.Field
	CardTokens             apijson.Field
	PreviousAuthRuleTokens apijson.Field
	ProgramLevel           apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AuthRuleRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleNewParams struct {
	// Merchant category codes for which the Auth Rule permits transactions.
	AllowedMcc param.Field[[]string] `json:"allowed_mcc"`
	// Merchant category codes for which the Auth Rule automatically declines
	// transactions.
	BlockedMcc param.Field[[]string] `json:"blocked_mcc"`
	// Countries in which the Auth Rule permits transactions. Note that Lithic
	// maintains a list of countries in which all transactions are blocked; "allowing"
	// those countries in an Auth Rule does not override the Lithic-wide restrictions.
	AllowedCountries param.Field[[]string] `json:"allowed_countries"`
	// Countries in which the Auth Rule automatically declines transactions.
	BlockedCountries param.Field[[]string] `json:"blocked_countries"`
	// Address verification to confirm that postal code entered at point of transaction
	// (if applicable) matches the postal code on file for a given card. Since this
	// check is performed against the address submitted via the Enroll Consumer
	// endpoint, it should only be used in cases where card users are enrolled with
	// their own accounts. Available values:
	//
	//   - `ZIP_ONLY` - AVS check is performed to confirm ZIP code entered at point of
	//     transaction (if applicable) matches address on file.
	AvsType param.Field[AuthRuleNewParamsAvsType] `json:"avs_type"`
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens param.Field[[]string] `json:"account_tokens"`
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens param.Field[[]string] `json:"card_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel param.Field[bool] `json:"program_level"`
}

func (r AuthRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleNewParamsAvsType string

const (
	AuthRuleNewParamsAvsTypeZipOnly AuthRuleNewParamsAvsType = "ZIP_ONLY"
)

type AuthRuleUpdateParams struct {
	// Array of merchant category codes for which the Auth Rule will permit
	// transactions. Note that only this field or `blocked_mcc` can be used for a given
	// Auth Rule.
	AllowedMcc param.Field[[]string] `json:"allowed_mcc"`
	// Array of merchant category codes for which the Auth Rule will automatically
	// decline transactions. Note that only this field or `allowed_mcc` can be used for
	// a given Auth Rule.
	BlockedMcc param.Field[[]string] `json:"blocked_mcc"`
	// Array of country codes for which the Auth Rule will permit transactions. Note
	// that only this field or `blocked_countries` can be used for a given Auth Rule.
	AllowedCountries param.Field[[]string] `json:"allowed_countries"`
	// Array of country codes for which the Auth Rule will automatically decline
	// transactions. Note that only this field or `allowed_countries` can be used for a
	// given Auth Rule.
	BlockedCountries param.Field[[]string] `json:"blocked_countries"`
	// Address verification to confirm that postal code entered at point of transaction
	// (if applicable) matches the postal code on file for a given card.
	AvsType param.Field[AuthRuleUpdateParamsAvsType] `json:"avs_type"`
}

func (r AuthRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleUpdateParamsAvsType string

const (
	AuthRuleUpdateParamsAvsTypeZipOnly AuthRuleUpdateParamsAvsType = "ZIP_ONLY"
)

type AuthRuleListParams struct {
	// Page (for pagination).
	Page param.Field[int64] `query:"page"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [AuthRuleListParams]'s query parameters as `url.Values`.
func (r AuthRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthRuleListResponse struct {
	Data []AuthRule `json:"data"`
	// Total number of entries.
	TotalEntries string `json:"total_entries"`
	// Total number of pages
	TotalPages int64 `json:"total_pages"`
	// Page number.
	Page int64 `json:"page"`
	JSON authRuleListResponseJSON
}

// authRuleListResponseJSON contains the JSON metadata for the struct
// [AuthRuleListResponse]
type authRuleListResponseJSON struct {
	Data         apijson.Field
	TotalEntries apijson.Field
	TotalPages   apijson.Field
	Page         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AuthRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleApplyParams struct {
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens param.Field[[]string] `json:"card_tokens"`
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens param.Field[[]string] `json:"account_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel param.Field[bool] `json:"program_level"`
}

func (r AuthRuleApplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleRemoveParams struct {
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens param.Field[[]string] `json:"card_tokens"`
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens param.Field[[]string] `json:"account_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel param.Field[bool] `json:"program_level"`
}

func (r AuthRuleRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
