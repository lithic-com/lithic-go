// File generated from our OpenAPI spec by Stainless.

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
func (r *AuthRuleService) New(ctx context.Context, body AuthRuleNewParams, opts ...option.RequestOption) (res *AuthRule, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detail the properties and entities (program, accounts, and cards) associated
// with an existing authorization rule (Auth Rule).
func (r *AuthRuleService) Get(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the properties associated with an existing authorization rule (Auth
// Rule).
func (r *AuthRuleService) Update(ctx context.Context, authRuleToken string, body AuthRuleUpdateParams, opts ...option.RequestOption) (res *AuthRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Return all of the Auth Rules under the program.
func (r *AuthRuleService) List(ctx context.Context, query AuthRuleListParams, opts ...option.RequestOption) (res *shared.CursorPage[AuthRule], err error) {
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
func (r *AuthRuleService) ListAutoPaging(ctx context.Context, query AuthRuleListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[AuthRule] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Applies an existing authorization rule (Auth Rule) to an program, account, or
// card level.
func (r *AuthRuleService) Apply(ctx context.Context, authRuleToken string, body AuthRuleApplyParams, opts ...option.RequestOption) (res *AuthRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s/apply", authRuleToken)
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
	Token string `json:"token,required" format:"uuid"`
	// Indicates whether the Auth Rule is ACTIVE or INACTIVE
	State AuthRuleState `json:"state,required"`
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens []string `json:"account_tokens"`
	// Countries in which the Auth Rule permits transactions. Note that Lithic
	// maintains a list of countries in which all transactions are blocked; "allowing"
	// those countries in an Auth Rule does not override the Lithic-wide restrictions.
	AllowedCountries []string `json:"allowed_countries"`
	// Merchant category codes for which the Auth Rule permits transactions.
	AllowedMcc []string `json:"allowed_mcc"`
	// Countries in which the Auth Rule automatically declines transactions.
	BlockedCountries []string `json:"blocked_countries"`
	// Merchant category codes for which the Auth Rule automatically declines
	// transactions.
	BlockedMcc []string `json:"blocked_mcc"`
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens []string `json:"card_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel bool `json:"program_level"`
	JSON         authRuleJSON
}

// authRuleJSON contains the JSON metadata for the struct [AuthRule]
type authRuleJSON struct {
	Token            apijson.Field
	State            apijson.Field
	AccountTokens    apijson.Field
	AllowedCountries apijson.Field
	AllowedMcc       apijson.Field
	BlockedCountries apijson.Field
	BlockedMcc       apijson.Field
	CardTokens       apijson.Field
	ProgramLevel     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AuthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether the Auth Rule is ACTIVE or INACTIVE
type AuthRuleState string

const (
	AuthRuleStateActive   AuthRuleState = "ACTIVE"
	AuthRuleStateInactive AuthRuleState = "INACTIVE"
)

type AuthRuleGetResponse struct {
	Data []AuthRule `json:"data"`
	JSON authRuleGetResponseJSON
}

// authRuleGetResponseJSON contains the JSON metadata for the struct
// [AuthRuleGetResponse]
type authRuleGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleRemoveResponse struct {
	AccountTokens []string `json:"account_tokens"`
	CardTokens    []string `json:"card_tokens"`
	ProgramLevel  bool     `json:"program_level"`
	JSON          authRuleRemoveResponseJSON
}

// authRuleRemoveResponseJSON contains the JSON metadata for the struct
// [AuthRuleRemoveResponse]
type authRuleRemoveResponseJSON struct {
	AccountTokens apijson.Field
	CardTokens    apijson.Field
	ProgramLevel  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuthRuleRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRuleNewParams struct {
	// Array of account_token(s) identifying the accounts that the Auth Rule applies
	// to. Note that only this field or `card_tokens` can be provided for a given Auth
	// Rule.
	AccountTokens param.Field[[]string] `json:"account_tokens"`
	// Countries in which the Auth Rule permits transactions. Note that Lithic
	// maintains a list of countries in which all transactions are blocked; "allowing"
	// those countries in an Auth Rule does not override the Lithic-wide restrictions.
	AllowedCountries param.Field[[]string] `json:"allowed_countries"`
	// Merchant category codes for which the Auth Rule permits transactions.
	AllowedMcc param.Field[[]string] `json:"allowed_mcc"`
	// Countries in which the Auth Rule automatically declines transactions.
	BlockedCountries param.Field[[]string] `json:"blocked_countries"`
	// Merchant category codes for which the Auth Rule automatically declines
	// transactions.
	BlockedMcc param.Field[[]string] `json:"blocked_mcc"`
	// Array of card_token(s) identifying the cards that the Auth Rule applies to. Note
	// that only this field or `account_tokens` can be provided for a given Auth Rule.
	CardTokens param.Field[[]string] `json:"card_tokens"`
	// Boolean indicating whether the Auth Rule is applied at the program level.
	ProgramLevel param.Field[bool] `json:"program_level"`
}

func (r AuthRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleUpdateParams struct {
	// Array of country codes for which the Auth Rule will permit transactions. Note
	// that only this field or `blocked_countries` can be used for a given Auth Rule.
	AllowedCountries param.Field[[]string] `json:"allowed_countries"`
	// Array of merchant category codes for which the Auth Rule will permit
	// transactions. Note that only this field or `blocked_mcc` can be used for a given
	// Auth Rule.
	AllowedMcc param.Field[[]string] `json:"allowed_mcc"`
	// Array of country codes for which the Auth Rule will automatically decline
	// transactions. Note that only this field or `allowed_countries` can be used for a
	// given Auth Rule.
	BlockedCountries param.Field[[]string] `json:"blocked_countries"`
	// Array of merchant category codes for which the Auth Rule will automatically
	// decline transactions. Note that only this field or `allowed_mcc` can be used for
	// a given Auth Rule.
	BlockedMcc param.Field[[]string] `json:"blocked_mcc"`
}

func (r AuthRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [AuthRuleListParams]'s query parameters as `url.Values`.
func (r AuthRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthRuleApplyParams struct {
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

func (r AuthRuleApplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleRemoveParams struct {
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

func (r AuthRuleRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
