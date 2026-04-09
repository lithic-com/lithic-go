// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/lithic-com/lithic-go/shared"
	"github.com/tidwall/gjson"
)

// AuthRuleV2Service contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthRuleV2Service] method instead.
type AuthRuleV2Service struct {
	Options   []option.RequestOption
	Backtests *AuthRuleV2BacktestService
}

// NewAuthRuleV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthRuleV2Service(opts ...option.RequestOption) (r *AuthRuleV2Service) {
	r = &AuthRuleV2Service{}
	r.Options = opts
	r.Backtests = NewAuthRuleV2BacktestService(opts...)
	return
}

// Creates a new V2 Auth rule in draft mode
func (r *AuthRuleV2Service) New(ctx context.Context, body AuthRuleV2NewParams, opts ...option.RequestOption) (res *AuthRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/auth_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetches a V2 Auth rule by its token
func (r *AuthRuleV2Service) Get(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRule, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a V2 Auth rule's properties
//
// If `account_tokens`, `card_tokens`, `program_level`, `excluded_card_tokens`,
// `excluded_account_tokens`, or `excluded_business_account_tokens` is provided,
// this will replace existing associations with the provided list of entities.
func (r *AuthRuleV2Service) Update(ctx context.Context, authRuleToken string, body AuthRuleV2UpdateParams, opts ...option.RequestOption) (res *AuthRule, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists V2 Auth rules
func (r *AuthRuleV2Service) List(ctx context.Context, query AuthRuleV2ListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AuthRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/auth_rules"
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

// Lists V2 Auth rules
func (r *AuthRuleV2Service) ListAutoPaging(ctx context.Context, query AuthRuleV2ListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AuthRule] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a V2 Auth rule
func (r *AuthRuleV2Service) Delete(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return err
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Creates a new draft version of a rule that will be ran in shadow mode.
//
// This can also be utilized to reset the draft parameters, causing a draft version
// to no longer be ran in shadow mode.
func (r *AuthRuleV2Service) Draft(ctx context.Context, authRuleToken string, body AuthRuleV2DraftParams, opts ...option.RequestOption) (res *AuthRule, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/auth_rules/%s/draft", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists Auth Rule evaluation results.
//
// **Limitations:**
//
// - Results are available for the past 3 months only
// - At least one filter (`event_token` or `auth_rule_token`) must be provided
// - When filtering by `event_token`, pagination is not supported
func (r *AuthRuleV2Service) ListResults(ctx context.Context, query AuthRuleV2ListResultsParams, opts ...option.RequestOption) (res *pagination.CursorPage[AuthRuleV2ListResultsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/auth_rules/results"
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

// Lists Auth Rule evaluation results.
//
// **Limitations:**
//
// - Results are available for the past 3 months only
// - At least one filter (`event_token` or `auth_rule_token`) must be provided
// - When filtering by `event_token`, pagination is not supported
func (r *AuthRuleV2Service) ListResultsAutoPaging(ctx context.Context, query AuthRuleV2ListResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AuthRuleV2ListResultsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListResults(ctx, query, opts...))
}

// Returns all versions of an auth rule, sorted by version number descending
// (newest first).
func (r *AuthRuleV2Service) ListVersions(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleV2ListVersionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/auth_rules/%s/versions", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Promotes the draft version of an Auth rule to the currently active version such
// that it is enforced in the respective stream.
func (r *AuthRuleV2Service) Promote(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRule, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/auth_rules/%s/promote", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Fetches the current calculated Feature values for the given Auth Rule
//
// This only calculates the features for the active version.
//
//   - VelocityLimit Rules calculates the current Velocity Feature data. This
//     requires a `card_token` or `account_token` matching what the rule is Scoped
//     to.
//   - ConditionalBlock Rules calculates the CARD*TRANSACTION_COUNT*\* attributes on
//     the rule. This requires a `card_token`
func (r *AuthRuleV2Service) GetFeatures(ctx context.Context, authRuleToken string, query AuthRuleV2GetFeaturesParams, opts ...option.RequestOption) (res *AuthRuleV2GetFeaturesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/auth_rules/%s/features", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves a performance report for an Auth rule containing daily statistics and
// evaluation outcomes.
//
// **Time Range Limitations:**
//
//   - Reports are supported for the past 3 months only
//   - Maximum interval length is 1 month
//   - Report data is available only through the previous day in UTC (current day
//     data is not available)
//
// The report provides daily statistics for both current and draft versions of the
// Auth rule, including approval, decline, and challenge counts along with sample
// events.
func (r *AuthRuleV2Service) GetReport(ctx context.Context, authRuleToken string, query AuthRuleV2GetReportParams, opts ...option.RequestOption) (res *AuthRuleV2GetReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/auth_rules/%s/report", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AuthRule struct {
	// Auth Rule Token
	Token string `json:"token" api:"required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens" api:"required" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens []string `json:"business_account_tokens" api:"required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string               `json:"card_tokens" api:"required" format:"uuid"`
	CurrentVersion AuthRuleCurrentVersion `json:"current_version" api:"required,nullable"`
	DraftVersion   AuthRuleDraftVersion   `json:"draft_version" api:"required,nullable"`
	// The event stream during which the rule will be evaluated.
	EventStream EventStream `json:"event_stream" api:"required"`
	// Indicates whether this auth rule is managed by Lithic. If true, the rule cannot
	// be modified or deleted by the user
	LithicManaged bool `json:"lithic_managed" api:"required"`
	// Auth Rule Name
	Name string `json:"name" api:"required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level" api:"required"`
	// The state of the Auth Rule
	State AuthRuleState `json:"state" api:"required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
	//     AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type AuthRuleType `json:"type" api:"required"`
	// Account tokens to which the Auth Rule does not apply.
	ExcludedAccountTokens []string `json:"excluded_account_tokens" format:"uuid"`
	// Business account tokens to which the Auth Rule does not apply.
	ExcludedBusinessAccountTokens []string `json:"excluded_business_account_tokens" format:"uuid"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string     `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleJSON `json:"-"`
}

// authRuleJSON contains the JSON metadata for the struct [AuthRule]
type authRuleJSON struct {
	Token                         apijson.Field
	AccountTokens                 apijson.Field
	BusinessAccountTokens         apijson.Field
	CardTokens                    apijson.Field
	CurrentVersion                apijson.Field
	DraftVersion                  apijson.Field
	EventStream                   apijson.Field
	LithicManaged                 apijson.Field
	Name                          apijson.Field
	ProgramLevel                  apijson.Field
	State                         apijson.Field
	Type                          apijson.Field
	ExcludedAccountTokens         apijson.Field
	ExcludedBusinessAccountTokens apijson.Field
	ExcludedCardTokens            apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *AuthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleJSON) RawJSON() string {
	return r.raw
}

type AuthRuleCurrentVersion struct {
	// Parameters for the Auth Rule
	Parameters AuthRuleCurrentVersionParameters `json:"parameters" api:"required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                      `json:"version" api:"required"`
	JSON    authRuleCurrentVersionJSON `json:"-"`
}

// authRuleCurrentVersionJSON contains the JSON metadata for the struct
// [AuthRuleCurrentVersion]
type authRuleCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the Auth Rule
type AuthRuleCurrentVersionParameters struct {
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// The TypeScript source code of the rule. Must define a `rule()` function that
	// accepts the declared features as positional arguments (in the same order as the
	// `features` array) and returns an array of actions.
	Code string `json:"code"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [[]RuleFeature].
	Features interface{}          `json:"features"`
	Filters  VelocityLimitFilters `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount" api:"nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count" api:"nullable"`
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleCurrentVersionParametersScope `json:"scope"`
	JSON  authRuleCurrentVersionParametersJSON  `json:"-"`
	union AuthRuleCurrentVersionParametersUnion
}

// authRuleCurrentVersionParametersJSON contains the JSON metadata for the struct
// [AuthRuleCurrentVersionParameters]
type authRuleCurrentVersionParametersJSON struct {
	Action      apijson.Field
	Code        apijson.Field
	Conditions  apijson.Field
	Features    apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleCurrentVersionParametersUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters],
// [TypescriptCodeParameters].
func (r AuthRuleCurrentVersionParameters) AsUnion() AuthRuleCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters],
// [ConditionalTokenizationActionParameters] or [TypescriptCodeParameters].
type AuthRuleCurrentVersionParametersUnion interface {
	implementsAuthRuleCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Conditional3DSActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalAuthorizationActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalACHActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalTokenizationActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TypescriptCodeParameters{}),
		},
	)
}

// The scope the velocity is calculated for
type AuthRuleCurrentVersionParametersScope string

const (
	AuthRuleCurrentVersionParametersScopeCard    AuthRuleCurrentVersionParametersScope = "CARD"
	AuthRuleCurrentVersionParametersScopeAccount AuthRuleCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleCurrentVersionParametersScopeCard, AuthRuleCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleDraftVersion struct {
	// An error message if the draft version failed compilation. Populated when `state`
	// is `ERROR`, `null` otherwise.
	Error string `json:"error" api:"required,nullable"`
	// Parameters for the Auth Rule
	Parameters AuthRuleDraftVersionParameters `json:"parameters" api:"required"`
	// The state of the draft version. Most rules are created synchronously and the
	// state is immediately `SHADOWING`. Rules backed by TypeScript code are compiled
	// asynchronously — the state starts as `PENDING` and transitions to `SHADOWING` on
	// success or `ERROR` on failure.
	//
	//   - `PENDING`: Compilation of the rule is in progress (TypeScript rules only).
	//   - `SHADOWING`: The draft version is ready and evaluating in shadow mode
	//     alongside the current active version. It can be promoted to the active
	//     version.
	//   - `ERROR`: Compilation of the rule failed. Check the `error` field for details.
	State AuthRuleDraftVersionState `json:"state" api:"required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                    `json:"version" api:"required"`
	JSON    authRuleDraftVersionJSON `json:"-"`
}

// authRuleDraftVersionJSON contains the JSON metadata for the struct
// [AuthRuleDraftVersion]
type authRuleDraftVersionJSON struct {
	Error       apijson.Field
	Parameters  apijson.Field
	State       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the Auth Rule
type AuthRuleDraftVersionParameters struct {
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// The TypeScript source code of the rule. Must define a `rule()` function that
	// accepts the declared features as positional arguments (in the same order as the
	// `features` array) and returns an array of actions.
	Code string `json:"code"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [[]RuleFeature].
	Features interface{}          `json:"features"`
	Filters  VelocityLimitFilters `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount" api:"nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count" api:"nullable"`
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleDraftVersionParametersScope `json:"scope"`
	JSON  authRuleDraftVersionParametersJSON  `json:"-"`
	union AuthRuleDraftVersionParametersUnion
}

// authRuleDraftVersionParametersJSON contains the JSON metadata for the struct
// [AuthRuleDraftVersionParameters]
type authRuleDraftVersionParametersJSON struct {
	Action      apijson.Field
	Code        apijson.Field
	Conditions  apijson.Field
	Features    apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleDraftVersionParametersUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters],
// [TypescriptCodeParameters].
func (r AuthRuleDraftVersionParameters) AsUnion() AuthRuleDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters],
// [ConditionalTokenizationActionParameters] or [TypescriptCodeParameters].
type AuthRuleDraftVersionParametersUnion interface {
	implementsAuthRuleDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Conditional3DSActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalAuthorizationActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalACHActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalTokenizationActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TypescriptCodeParameters{}),
		},
	)
}

// The scope the velocity is calculated for
type AuthRuleDraftVersionParametersScope string

const (
	AuthRuleDraftVersionParametersScopeCard    AuthRuleDraftVersionParametersScope = "CARD"
	AuthRuleDraftVersionParametersScopeAccount AuthRuleDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleDraftVersionParametersScopeCard, AuthRuleDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the draft version. Most rules are created synchronously and the
// state is immediately `SHADOWING`. Rules backed by TypeScript code are compiled
// asynchronously — the state starts as `PENDING` and transitions to `SHADOWING` on
// success or `ERROR` on failure.
//
//   - `PENDING`: Compilation of the rule is in progress (TypeScript rules only).
//   - `SHADOWING`: The draft version is ready and evaluating in shadow mode
//     alongside the current active version. It can be promoted to the active
//     version.
//   - `ERROR`: Compilation of the rule failed. Check the `error` field for details.
type AuthRuleDraftVersionState string

const (
	AuthRuleDraftVersionStatePending   AuthRuleDraftVersionState = "PENDING"
	AuthRuleDraftVersionStateShadowing AuthRuleDraftVersionState = "SHADOWING"
	AuthRuleDraftVersionStateError     AuthRuleDraftVersionState = "ERROR"
)

func (r AuthRuleDraftVersionState) IsKnown() bool {
	switch r {
	case AuthRuleDraftVersionStatePending, AuthRuleDraftVersionStateShadowing, AuthRuleDraftVersionStateError:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleState string

const (
	AuthRuleStateActive   AuthRuleState = "ACTIVE"
	AuthRuleStateInactive AuthRuleState = "INACTIVE"
)

func (r AuthRuleState) IsKnown() bool {
	switch r {
	case AuthRuleStateActive, AuthRuleStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
//     AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
type AuthRuleType string

const (
	AuthRuleTypeConditionalBlock  AuthRuleType = "CONDITIONAL_BLOCK"
	AuthRuleTypeVelocityLimit     AuthRuleType = "VELOCITY_LIMIT"
	AuthRuleTypeMerchantLock      AuthRuleType = "MERCHANT_LOCK"
	AuthRuleTypeConditionalAction AuthRuleType = "CONDITIONAL_ACTION"
	AuthRuleTypeTypescriptCode    AuthRuleType = "TYPESCRIPT_CODE"
)

func (r AuthRuleType) IsKnown() bool {
	switch r {
	case AuthRuleTypeConditionalBlock, AuthRuleTypeVelocityLimit, AuthRuleTypeMerchantLock, AuthRuleTypeConditionalAction, AuthRuleTypeTypescriptCode:
		return true
	}
	return false
}

type AuthRuleCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-character alphabetic ISO 4217 code for the merchant currency of
	//     the transaction.
	//   - `MERCHANT_ID`: Unique alphanumeric identifier for the payment card acceptor
	//     (merchant).
	//   - `DESCRIPTOR`: Short description of card acceptor.
	//   - `LIABILITY_SHIFT`: Indicates whether chargeback liability shift to the issuer
	//     applies to the transaction. Valid values are `NONE`, `3DS_AUTHENTICATED`, or
	//     `TOKEN_AUTHENTICATED`.
	//   - `PAN_ENTRY_MODE`: The method by which the cardholder's primary account number
	//     (PAN) was entered. Valid values are `AUTO_ENTRY`, `BAR_CODE`, `CONTACTLESS`,
	//     `ECOMMERCE`, `ERROR_KEYED`, `ERROR_MAGNETIC_STRIPE`, `ICC`, `KEY_ENTERED`,
	//     `MAGNETIC_STRIPE`, `MANUAL`, `OCR`, `SECURE_CARDLESS`, `UNSPECIFIED`,
	//     `UNKNOWN`, `CREDENTIAL_ON_FILE`, or `ECOMMERCE`.
	//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
	//     fee field in the settlement/cardholder billing currency. This is the amount
	//     the issuer should authorize against unless the issuer is paying the acquirer
	//     fee on behalf of the cardholder.
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authorization. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `CARD_TRANSACTION_COUNT_15M`: The number of transactions on the card in the
	//     trailing 15 minutes before the authorization.
	//   - `CARD_TRANSACTION_COUNT_1H`: The number of transactions on the card in the
	//     trailing hour up and until the authorization.
	//   - `CARD_TRANSACTION_COUNT_24H`: The number of transactions on the card in the
	//     trailing 24 hours up and until the authorization.
	//   - `CARD_STATE`: The current state of the card associated with the transaction.
	//     Valid values are `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`,
	//     `PENDING_FULFILLMENT`.
	//   - `PIN_ENTERED`: Indicates whether a PIN was entered during the transaction.
	//     Valid values are `TRUE`, `FALSE`.
	//   - `PIN_STATUS`: The current state of card's PIN. Valid values are `NOT_SET`,
	//     `OK`, `BLOCKED`.
	//   - `WALLET_TYPE`: For transactions using a digital wallet token, indicates the
	//     source of the token. Valid values are `APPLE_PAY`, `GOOGLE_PAY`,
	//     `SAMSUNG_PAY`, `MASTERPASS`, `MERCHANT`, `OTHER`, `NONE`.
	//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
	//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
	//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
	Attribute ConditionalAttribute `json:"attribute" api:"required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion `json:"value" api:"required"`
	JSON  authRuleConditionJSON `json:"-"`
}

// authRuleConditionJSON contains the JSON metadata for the struct
// [AuthRuleCondition]
type authRuleConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleConditionJSON) RawJSON() string {
	return r.raw
}

type AuthRuleConditionParam struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-character alphabetic ISO 4217 code for the merchant currency of
	//     the transaction.
	//   - `MERCHANT_ID`: Unique alphanumeric identifier for the payment card acceptor
	//     (merchant).
	//   - `DESCRIPTOR`: Short description of card acceptor.
	//   - `LIABILITY_SHIFT`: Indicates whether chargeback liability shift to the issuer
	//     applies to the transaction. Valid values are `NONE`, `3DS_AUTHENTICATED`, or
	//     `TOKEN_AUTHENTICATED`.
	//   - `PAN_ENTRY_MODE`: The method by which the cardholder's primary account number
	//     (PAN) was entered. Valid values are `AUTO_ENTRY`, `BAR_CODE`, `CONTACTLESS`,
	//     `ECOMMERCE`, `ERROR_KEYED`, `ERROR_MAGNETIC_STRIPE`, `ICC`, `KEY_ENTERED`,
	//     `MAGNETIC_STRIPE`, `MANUAL`, `OCR`, `SECURE_CARDLESS`, `UNSPECIFIED`,
	//     `UNKNOWN`, `CREDENTIAL_ON_FILE`, or `ECOMMERCE`.
	//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
	//     fee field in the settlement/cardholder billing currency. This is the amount
	//     the issuer should authorize against unless the issuer is paying the acquirer
	//     fee on behalf of the cardholder.
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authorization. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `CARD_TRANSACTION_COUNT_15M`: The number of transactions on the card in the
	//     trailing 15 minutes before the authorization.
	//   - `CARD_TRANSACTION_COUNT_1H`: The number of transactions on the card in the
	//     trailing hour up and until the authorization.
	//   - `CARD_TRANSACTION_COUNT_24H`: The number of transactions on the card in the
	//     trailing 24 hours up and until the authorization.
	//   - `CARD_STATE`: The current state of the card associated with the transaction.
	//     Valid values are `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`,
	//     `PENDING_FULFILLMENT`.
	//   - `PIN_ENTERED`: Indicates whether a PIN was entered during the transaction.
	//     Valid values are `TRUE`, `FALSE`.
	//   - `PIN_STATUS`: The current state of card's PIN. Valid values are `NOT_SET`,
	//     `OK`, `BLOCKED`.
	//   - `WALLET_TYPE`: For transactions using a digital wallet token, indicates the
	//     source of the token. Valid values are `APPLE_PAY`, `GOOGLE_PAY`,
	//     `SAMSUNG_PAY`, `MASTERPASS`, `MERCHANT`, `OTHER`, `NONE`.
	//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
	//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
	//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
	Attribute param.Field[ConditionalAttribute] `json:"attribute" api:"required"`
	// The operation to apply to the attribute
	Operation param.Field[ConditionalOperation] `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[ConditionalValueUnionParam] `json:"value" api:"required"`
}

func (r AuthRuleConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleVersion struct {
	// Timestamp of when this version was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Parameters for the Auth Rule
	Parameters AuthRuleVersionParameters `json:"parameters" api:"required"`
	// The current state of this version.
	State AuthRuleVersionState `json:"state" api:"required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64               `json:"version" api:"required"`
	JSON    authRuleVersionJSON `json:"-"`
}

// authRuleVersionJSON contains the JSON metadata for the struct [AuthRuleVersion]
type authRuleVersionJSON struct {
	Created     apijson.Field
	Parameters  apijson.Field
	State       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the Auth Rule
type AuthRuleVersionParameters struct {
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// The TypeScript source code of the rule. Must define a `rule()` function that
	// accepts the declared features as positional arguments (in the same order as the
	// `features` array) and returns an array of actions.
	Code string `json:"code"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [[]RuleFeature].
	Features interface{}          `json:"features"`
	Filters  VelocityLimitFilters `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount" api:"nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count" api:"nullable"`
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleVersionParametersScope `json:"scope"`
	JSON  authRuleVersionParametersJSON  `json:"-"`
	union AuthRuleVersionParametersUnion
}

// authRuleVersionParametersJSON contains the JSON metadata for the struct
// [AuthRuleVersionParameters]
type authRuleVersionParametersJSON struct {
	Action      apijson.Field
	Code        apijson.Field
	Conditions  apijson.Field
	Features    apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleVersionParametersUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters],
// [TypescriptCodeParameters].
func (r AuthRuleVersionParameters) AsUnion() AuthRuleVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters],
// [ConditionalTokenizationActionParameters] or [TypescriptCodeParameters].
type AuthRuleVersionParametersUnion interface {
	implementsAuthRuleVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Conditional3DSActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalAuthorizationActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalACHActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalTokenizationActionParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TypescriptCodeParameters{}),
		},
	)
}

// The scope the velocity is calculated for
type AuthRuleVersionParametersScope string

const (
	AuthRuleVersionParametersScopeCard    AuthRuleVersionParametersScope = "CARD"
	AuthRuleVersionParametersScopeAccount AuthRuleVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleVersionParametersScopeCard, AuthRuleVersionParametersScopeAccount:
		return true
	}
	return false
}

// The current state of this version.
type AuthRuleVersionState string

const (
	AuthRuleVersionStateActive   AuthRuleVersionState = "ACTIVE"
	AuthRuleVersionStateShadow   AuthRuleVersionState = "SHADOW"
	AuthRuleVersionStateInactive AuthRuleVersionState = "INACTIVE"
)

func (r AuthRuleVersionState) IsKnown() bool {
	switch r {
	case AuthRuleVersionStateActive, AuthRuleVersionStateShadow, AuthRuleVersionStateInactive:
		return true
	}
	return false
}

type BacktestStats struct {
	// The total number of historical transactions approved by this rule during the
	// backtest period, or the number of transactions that would have been approved if
	// the rule was evaluated in shadow mode.
	Approved int64 `json:"approved"`
	// The total number of historical transactions challenged by this rule during the
	// backtest period, or the number of transactions that would have been challenged
	// if the rule was evaluated in shadow mode. Currently applicable only for 3DS Auth
	// Rules.
	Challenged int64 `json:"challenged"`
	// The total number of historical transactions declined by this rule during the
	// backtest period, or the number of transactions that would have been declined if
	// the rule was evaluated in shadow mode.
	Declined int64 `json:"declined"`
	// Example events and their outcomes.
	Examples []BacktestStatsExample `json:"examples"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64             `json:"version"`
	JSON    backtestStatsJSON `json:"-"`
}

// backtestStatsJSON contains the JSON metadata for the struct [BacktestStats]
type backtestStatsJSON struct {
	Approved    apijson.Field
	Challenged  apijson.Field
	Declined    apijson.Field
	Examples    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BacktestStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestStatsJSON) RawJSON() string {
	return r.raw
}

type BacktestStatsExample struct {
	// The decision made by the rule for this event.
	Decision BacktestStatsExamplesDecision `json:"decision"`
	// The event token.
	EventToken string `json:"event_token" format:"uuid"`
	// The timestamp of the event.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// The token of the transaction associated with the event
	TransactionToken string                   `json:"transaction_token" api:"nullable" format:"uuid"`
	JSON             backtestStatsExampleJSON `json:"-"`
}

// backtestStatsExampleJSON contains the JSON metadata for the struct
// [BacktestStatsExample]
type backtestStatsExampleJSON struct {
	Decision         apijson.Field
	EventToken       apijson.Field
	Timestamp        apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BacktestStatsExample) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestStatsExampleJSON) RawJSON() string {
	return r.raw
}

// The decision made by the rule for this event.
type BacktestStatsExamplesDecision string

const (
	BacktestStatsExamplesDecisionApproved   BacktestStatsExamplesDecision = "APPROVED"
	BacktestStatsExamplesDecisionDeclined   BacktestStatsExamplesDecision = "DECLINED"
	BacktestStatsExamplesDecisionChallenged BacktestStatsExamplesDecision = "CHALLENGED"
)

func (r BacktestStatsExamplesDecision) IsKnown() bool {
	switch r {
	case BacktestStatsExamplesDecisionApproved, BacktestStatsExamplesDecisionDeclined, BacktestStatsExamplesDecisionChallenged:
		return true
	}
	return false
}

type Conditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     Conditional3DSActionParametersAction      `json:"action" api:"required"`
	Conditions []Conditional3DsActionParametersCondition `json:"conditions" api:"required"`
	JSON       conditional3DsActionParametersJSON        `json:"-"`
}

// conditional3DsActionParametersJSON contains the JSON metadata for the struct
// [Conditional3DSActionParameters]
type conditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Conditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r Conditional3DSActionParameters) implementsAuthRuleCurrentVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleDraftVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleVersionParameters() {}

// The action to take if the conditions are met.
type Conditional3DSActionParametersAction string

const (
	Conditional3DSActionParametersActionDecline   Conditional3DSActionParametersAction = "DECLINE"
	Conditional3DSActionParametersActionChallenge Conditional3DSActionParametersAction = "CHALLENGE"
)

func (r Conditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case Conditional3DSActionParametersActionDecline, Conditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type Conditional3DsActionParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-character alphabetic ISO 4217 code for the merchant currency of
	//     the transaction.
	//   - `MERCHANT_ID`: Unique alphanumeric identifier for the payment card acceptor
	//     (merchant).
	//   - `DESCRIPTOR`: Short description of card acceptor.
	//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
	//     fee field in the settlement/cardholder billing currency. This is the amount
	//     the issuer should authorize against unless the issuer is paying the acquirer
	//     fee on behalf of the cardholder.
	//   - `RISK_SCORE`: Mastercard only: Assessment by the network of the authentication
	//     risk level, with a higher value indicating a higher amount of risk.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
	//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
	//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
	Attribute Conditional3DSActionParametersConditionsAttribute `json:"attribute" api:"required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion                       `json:"value" api:"required"`
	JSON  conditional3DsActionParametersConditionJSON `json:"-"`
}

// conditional3DsActionParametersConditionJSON contains the JSON metadata for the
// struct [Conditional3DsActionParametersCondition]
type conditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Conditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditional3DsActionParametersConditionJSON) RawJSON() string {
	return r.raw
}

// The attribute to target.
//
// The following attributes may be targeted:
//
//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
//     business by the types of goods or services it provides.
//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
//     Netherlands Antilles.
//   - `CURRENCY`: 3-character alphabetic ISO 4217 code for the merchant currency of
//     the transaction.
//   - `MERCHANT_ID`: Unique alphanumeric identifier for the payment card acceptor
//     (merchant).
//   - `DESCRIPTOR`: Short description of card acceptor.
//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
//     fee field in the settlement/cardholder billing currency. This is the amount
//     the issuer should authorize against unless the issuer is paying the acquirer
//     fee on behalf of the cardholder.
//   - `RISK_SCORE`: Mastercard only: Assessment by the network of the authentication
//     risk level, with a higher value indicating a higher amount of risk.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
type Conditional3DSActionParametersConditionsAttribute string

const (
	Conditional3DSActionParametersConditionsAttributeMcc               Conditional3DSActionParametersConditionsAttribute = "MCC"
	Conditional3DSActionParametersConditionsAttributeCountry           Conditional3DSActionParametersConditionsAttribute = "COUNTRY"
	Conditional3DSActionParametersConditionsAttributeCurrency          Conditional3DSActionParametersConditionsAttribute = "CURRENCY"
	Conditional3DSActionParametersConditionsAttributeMerchantID        Conditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	Conditional3DSActionParametersConditionsAttributeDescriptor        Conditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	Conditional3DSActionParametersConditionsAttributeTransactionAmount Conditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	Conditional3DSActionParametersConditionsAttributeRiskScore         Conditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	Conditional3DSActionParametersConditionsAttributeMessageCategory   Conditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
	Conditional3DSActionParametersConditionsAttributeAddressMatch      Conditional3DSActionParametersConditionsAttribute = "ADDRESS_MATCH"
)

func (r Conditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case Conditional3DSActionParametersConditionsAttributeMcc, Conditional3DSActionParametersConditionsAttributeCountry, Conditional3DSActionParametersConditionsAttributeCurrency, Conditional3DSActionParametersConditionsAttributeMerchantID, Conditional3DSActionParametersConditionsAttributeDescriptor, Conditional3DSActionParametersConditionsAttributeTransactionAmount, Conditional3DSActionParametersConditionsAttributeRiskScore, Conditional3DSActionParametersConditionsAttributeMessageCategory, Conditional3DSActionParametersConditionsAttributeAddressMatch:
		return true
	}
	return false
}

type ConditionalACHActionParameters struct {
	// The action to take if the conditions are met.
	Action     ConditionalACHActionParametersAction      `json:"action" api:"required"`
	Conditions []ConditionalACHActionParametersCondition `json:"conditions" api:"required"`
	JSON       conditionalACHActionParametersJSON        `json:"-"`
}

// conditionalACHActionParametersJSON contains the JSON metadata for the struct
// [ConditionalACHActionParameters]
type conditionalACHActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalACHActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalACHActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalACHActionParameters) implementsAuthRuleCurrentVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleDraftVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleVersionParameters() {}

// The action to take if the conditions are met.
type ConditionalACHActionParametersAction struct {
	// Approve the ACH transaction
	Type ConditionalACHActionParametersActionType `json:"type" api:"required"`
	// NACHA return code to use when returning the transaction. Note that the list of
	// available return codes is subject to an allowlist configured at the program
	// level
	Code  ConditionalACHActionParametersActionCode `json:"code"`
	JSON  conditionalACHActionParametersActionJSON `json:"-"`
	union ConditionalACHActionParametersActionUnion
}

// conditionalACHActionParametersActionJSON contains the JSON metadata for the
// struct [ConditionalACHActionParametersAction]
type conditionalACHActionParametersActionJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r conditionalACHActionParametersActionJSON) RawJSON() string {
	return r.raw
}

func (r *ConditionalACHActionParametersAction) UnmarshalJSON(data []byte) (err error) {
	*r = ConditionalACHActionParametersAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConditionalACHActionParametersActionUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConditionalACHActionParametersActionApproveActionACH],
// [ConditionalACHActionParametersActionReturnAction].
func (r ConditionalACHActionParametersAction) AsUnion() ConditionalACHActionParametersActionUnion {
	return r.union
}

// The action to take if the conditions are met.
//
// Union satisfied by [ConditionalACHActionParametersActionApproveActionACH] or
// [ConditionalACHActionParametersActionReturnAction].
type ConditionalACHActionParametersActionUnion interface {
	implementsConditionalACHActionParametersAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConditionalACHActionParametersActionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalACHActionParametersActionApproveActionACH{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalACHActionParametersActionReturnAction{}),
		},
	)
}

type ConditionalACHActionParametersActionApproveActionACH struct {
	// Approve the ACH transaction
	Type ConditionalACHActionParametersActionApproveActionACHType `json:"type" api:"required"`
	JSON conditionalACHActionParametersActionApproveActionACHJSON `json:"-"`
}

// conditionalACHActionParametersActionApproveActionACHJSON contains the JSON
// metadata for the struct [ConditionalACHActionParametersActionApproveActionACH]
type conditionalACHActionParametersActionApproveActionACHJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalACHActionParametersActionApproveActionACH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalACHActionParametersActionApproveActionACHJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalACHActionParametersActionApproveActionACH) implementsConditionalACHActionParametersAction() {
}

// Approve the ACH transaction
type ConditionalACHActionParametersActionApproveActionACHType string

const (
	ConditionalACHActionParametersActionApproveActionACHTypeApprove ConditionalACHActionParametersActionApproveActionACHType = "APPROVE"
)

func (r ConditionalACHActionParametersActionApproveActionACHType) IsKnown() bool {
	switch r {
	case ConditionalACHActionParametersActionApproveActionACHTypeApprove:
		return true
	}
	return false
}

type ConditionalACHActionParametersActionReturnAction struct {
	// NACHA return code to use when returning the transaction. Note that the list of
	// available return codes is subject to an allowlist configured at the program
	// level
	Code ConditionalACHActionParametersActionReturnActionCode `json:"code" api:"required"`
	// Return the ACH transaction
	Type ConditionalACHActionParametersActionReturnActionType `json:"type" api:"required"`
	JSON conditionalACHActionParametersActionReturnActionJSON `json:"-"`
}

// conditionalACHActionParametersActionReturnActionJSON contains the JSON metadata
// for the struct [ConditionalACHActionParametersActionReturnAction]
type conditionalACHActionParametersActionReturnActionJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalACHActionParametersActionReturnAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalACHActionParametersActionReturnActionJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalACHActionParametersActionReturnAction) implementsConditionalACHActionParametersAction() {
}

// NACHA return code to use when returning the transaction. Note that the list of
// available return codes is subject to an allowlist configured at the program
// level
type ConditionalACHActionParametersActionReturnActionCode string

const (
	ConditionalACHActionParametersActionReturnActionCodeR01 ConditionalACHActionParametersActionReturnActionCode = "R01"
	ConditionalACHActionParametersActionReturnActionCodeR02 ConditionalACHActionParametersActionReturnActionCode = "R02"
	ConditionalACHActionParametersActionReturnActionCodeR03 ConditionalACHActionParametersActionReturnActionCode = "R03"
	ConditionalACHActionParametersActionReturnActionCodeR04 ConditionalACHActionParametersActionReturnActionCode = "R04"
	ConditionalACHActionParametersActionReturnActionCodeR05 ConditionalACHActionParametersActionReturnActionCode = "R05"
	ConditionalACHActionParametersActionReturnActionCodeR06 ConditionalACHActionParametersActionReturnActionCode = "R06"
	ConditionalACHActionParametersActionReturnActionCodeR07 ConditionalACHActionParametersActionReturnActionCode = "R07"
	ConditionalACHActionParametersActionReturnActionCodeR08 ConditionalACHActionParametersActionReturnActionCode = "R08"
	ConditionalACHActionParametersActionReturnActionCodeR09 ConditionalACHActionParametersActionReturnActionCode = "R09"
	ConditionalACHActionParametersActionReturnActionCodeR10 ConditionalACHActionParametersActionReturnActionCode = "R10"
	ConditionalACHActionParametersActionReturnActionCodeR11 ConditionalACHActionParametersActionReturnActionCode = "R11"
	ConditionalACHActionParametersActionReturnActionCodeR12 ConditionalACHActionParametersActionReturnActionCode = "R12"
	ConditionalACHActionParametersActionReturnActionCodeR13 ConditionalACHActionParametersActionReturnActionCode = "R13"
	ConditionalACHActionParametersActionReturnActionCodeR14 ConditionalACHActionParametersActionReturnActionCode = "R14"
	ConditionalACHActionParametersActionReturnActionCodeR15 ConditionalACHActionParametersActionReturnActionCode = "R15"
	ConditionalACHActionParametersActionReturnActionCodeR16 ConditionalACHActionParametersActionReturnActionCode = "R16"
	ConditionalACHActionParametersActionReturnActionCodeR17 ConditionalACHActionParametersActionReturnActionCode = "R17"
	ConditionalACHActionParametersActionReturnActionCodeR18 ConditionalACHActionParametersActionReturnActionCode = "R18"
	ConditionalACHActionParametersActionReturnActionCodeR19 ConditionalACHActionParametersActionReturnActionCode = "R19"
	ConditionalACHActionParametersActionReturnActionCodeR20 ConditionalACHActionParametersActionReturnActionCode = "R20"
	ConditionalACHActionParametersActionReturnActionCodeR21 ConditionalACHActionParametersActionReturnActionCode = "R21"
	ConditionalACHActionParametersActionReturnActionCodeR22 ConditionalACHActionParametersActionReturnActionCode = "R22"
	ConditionalACHActionParametersActionReturnActionCodeR23 ConditionalACHActionParametersActionReturnActionCode = "R23"
	ConditionalACHActionParametersActionReturnActionCodeR24 ConditionalACHActionParametersActionReturnActionCode = "R24"
	ConditionalACHActionParametersActionReturnActionCodeR25 ConditionalACHActionParametersActionReturnActionCode = "R25"
	ConditionalACHActionParametersActionReturnActionCodeR26 ConditionalACHActionParametersActionReturnActionCode = "R26"
	ConditionalACHActionParametersActionReturnActionCodeR27 ConditionalACHActionParametersActionReturnActionCode = "R27"
	ConditionalACHActionParametersActionReturnActionCodeR28 ConditionalACHActionParametersActionReturnActionCode = "R28"
	ConditionalACHActionParametersActionReturnActionCodeR29 ConditionalACHActionParametersActionReturnActionCode = "R29"
	ConditionalACHActionParametersActionReturnActionCodeR30 ConditionalACHActionParametersActionReturnActionCode = "R30"
	ConditionalACHActionParametersActionReturnActionCodeR31 ConditionalACHActionParametersActionReturnActionCode = "R31"
	ConditionalACHActionParametersActionReturnActionCodeR32 ConditionalACHActionParametersActionReturnActionCode = "R32"
	ConditionalACHActionParametersActionReturnActionCodeR33 ConditionalACHActionParametersActionReturnActionCode = "R33"
	ConditionalACHActionParametersActionReturnActionCodeR34 ConditionalACHActionParametersActionReturnActionCode = "R34"
	ConditionalACHActionParametersActionReturnActionCodeR35 ConditionalACHActionParametersActionReturnActionCode = "R35"
	ConditionalACHActionParametersActionReturnActionCodeR36 ConditionalACHActionParametersActionReturnActionCode = "R36"
	ConditionalACHActionParametersActionReturnActionCodeR37 ConditionalACHActionParametersActionReturnActionCode = "R37"
	ConditionalACHActionParametersActionReturnActionCodeR38 ConditionalACHActionParametersActionReturnActionCode = "R38"
	ConditionalACHActionParametersActionReturnActionCodeR39 ConditionalACHActionParametersActionReturnActionCode = "R39"
	ConditionalACHActionParametersActionReturnActionCodeR40 ConditionalACHActionParametersActionReturnActionCode = "R40"
	ConditionalACHActionParametersActionReturnActionCodeR41 ConditionalACHActionParametersActionReturnActionCode = "R41"
	ConditionalACHActionParametersActionReturnActionCodeR42 ConditionalACHActionParametersActionReturnActionCode = "R42"
	ConditionalACHActionParametersActionReturnActionCodeR43 ConditionalACHActionParametersActionReturnActionCode = "R43"
	ConditionalACHActionParametersActionReturnActionCodeR44 ConditionalACHActionParametersActionReturnActionCode = "R44"
	ConditionalACHActionParametersActionReturnActionCodeR45 ConditionalACHActionParametersActionReturnActionCode = "R45"
	ConditionalACHActionParametersActionReturnActionCodeR46 ConditionalACHActionParametersActionReturnActionCode = "R46"
	ConditionalACHActionParametersActionReturnActionCodeR47 ConditionalACHActionParametersActionReturnActionCode = "R47"
	ConditionalACHActionParametersActionReturnActionCodeR50 ConditionalACHActionParametersActionReturnActionCode = "R50"
	ConditionalACHActionParametersActionReturnActionCodeR51 ConditionalACHActionParametersActionReturnActionCode = "R51"
	ConditionalACHActionParametersActionReturnActionCodeR52 ConditionalACHActionParametersActionReturnActionCode = "R52"
	ConditionalACHActionParametersActionReturnActionCodeR53 ConditionalACHActionParametersActionReturnActionCode = "R53"
	ConditionalACHActionParametersActionReturnActionCodeR61 ConditionalACHActionParametersActionReturnActionCode = "R61"
	ConditionalACHActionParametersActionReturnActionCodeR62 ConditionalACHActionParametersActionReturnActionCode = "R62"
	ConditionalACHActionParametersActionReturnActionCodeR67 ConditionalACHActionParametersActionReturnActionCode = "R67"
	ConditionalACHActionParametersActionReturnActionCodeR68 ConditionalACHActionParametersActionReturnActionCode = "R68"
	ConditionalACHActionParametersActionReturnActionCodeR69 ConditionalACHActionParametersActionReturnActionCode = "R69"
	ConditionalACHActionParametersActionReturnActionCodeR70 ConditionalACHActionParametersActionReturnActionCode = "R70"
	ConditionalACHActionParametersActionReturnActionCodeR71 ConditionalACHActionParametersActionReturnActionCode = "R71"
	ConditionalACHActionParametersActionReturnActionCodeR72 ConditionalACHActionParametersActionReturnActionCode = "R72"
	ConditionalACHActionParametersActionReturnActionCodeR73 ConditionalACHActionParametersActionReturnActionCode = "R73"
	ConditionalACHActionParametersActionReturnActionCodeR74 ConditionalACHActionParametersActionReturnActionCode = "R74"
	ConditionalACHActionParametersActionReturnActionCodeR75 ConditionalACHActionParametersActionReturnActionCode = "R75"
	ConditionalACHActionParametersActionReturnActionCodeR76 ConditionalACHActionParametersActionReturnActionCode = "R76"
	ConditionalACHActionParametersActionReturnActionCodeR77 ConditionalACHActionParametersActionReturnActionCode = "R77"
	ConditionalACHActionParametersActionReturnActionCodeR80 ConditionalACHActionParametersActionReturnActionCode = "R80"
	ConditionalACHActionParametersActionReturnActionCodeR81 ConditionalACHActionParametersActionReturnActionCode = "R81"
	ConditionalACHActionParametersActionReturnActionCodeR82 ConditionalACHActionParametersActionReturnActionCode = "R82"
	ConditionalACHActionParametersActionReturnActionCodeR83 ConditionalACHActionParametersActionReturnActionCode = "R83"
	ConditionalACHActionParametersActionReturnActionCodeR84 ConditionalACHActionParametersActionReturnActionCode = "R84"
	ConditionalACHActionParametersActionReturnActionCodeR85 ConditionalACHActionParametersActionReturnActionCode = "R85"
)

func (r ConditionalACHActionParametersActionReturnActionCode) IsKnown() bool {
	switch r {
	case ConditionalACHActionParametersActionReturnActionCodeR01, ConditionalACHActionParametersActionReturnActionCodeR02, ConditionalACHActionParametersActionReturnActionCodeR03, ConditionalACHActionParametersActionReturnActionCodeR04, ConditionalACHActionParametersActionReturnActionCodeR05, ConditionalACHActionParametersActionReturnActionCodeR06, ConditionalACHActionParametersActionReturnActionCodeR07, ConditionalACHActionParametersActionReturnActionCodeR08, ConditionalACHActionParametersActionReturnActionCodeR09, ConditionalACHActionParametersActionReturnActionCodeR10, ConditionalACHActionParametersActionReturnActionCodeR11, ConditionalACHActionParametersActionReturnActionCodeR12, ConditionalACHActionParametersActionReturnActionCodeR13, ConditionalACHActionParametersActionReturnActionCodeR14, ConditionalACHActionParametersActionReturnActionCodeR15, ConditionalACHActionParametersActionReturnActionCodeR16, ConditionalACHActionParametersActionReturnActionCodeR17, ConditionalACHActionParametersActionReturnActionCodeR18, ConditionalACHActionParametersActionReturnActionCodeR19, ConditionalACHActionParametersActionReturnActionCodeR20, ConditionalACHActionParametersActionReturnActionCodeR21, ConditionalACHActionParametersActionReturnActionCodeR22, ConditionalACHActionParametersActionReturnActionCodeR23, ConditionalACHActionParametersActionReturnActionCodeR24, ConditionalACHActionParametersActionReturnActionCodeR25, ConditionalACHActionParametersActionReturnActionCodeR26, ConditionalACHActionParametersActionReturnActionCodeR27, ConditionalACHActionParametersActionReturnActionCodeR28, ConditionalACHActionParametersActionReturnActionCodeR29, ConditionalACHActionParametersActionReturnActionCodeR30, ConditionalACHActionParametersActionReturnActionCodeR31, ConditionalACHActionParametersActionReturnActionCodeR32, ConditionalACHActionParametersActionReturnActionCodeR33, ConditionalACHActionParametersActionReturnActionCodeR34, ConditionalACHActionParametersActionReturnActionCodeR35, ConditionalACHActionParametersActionReturnActionCodeR36, ConditionalACHActionParametersActionReturnActionCodeR37, ConditionalACHActionParametersActionReturnActionCodeR38, ConditionalACHActionParametersActionReturnActionCodeR39, ConditionalACHActionParametersActionReturnActionCodeR40, ConditionalACHActionParametersActionReturnActionCodeR41, ConditionalACHActionParametersActionReturnActionCodeR42, ConditionalACHActionParametersActionReturnActionCodeR43, ConditionalACHActionParametersActionReturnActionCodeR44, ConditionalACHActionParametersActionReturnActionCodeR45, ConditionalACHActionParametersActionReturnActionCodeR46, ConditionalACHActionParametersActionReturnActionCodeR47, ConditionalACHActionParametersActionReturnActionCodeR50, ConditionalACHActionParametersActionReturnActionCodeR51, ConditionalACHActionParametersActionReturnActionCodeR52, ConditionalACHActionParametersActionReturnActionCodeR53, ConditionalACHActionParametersActionReturnActionCodeR61, ConditionalACHActionParametersActionReturnActionCodeR62, ConditionalACHActionParametersActionReturnActionCodeR67, ConditionalACHActionParametersActionReturnActionCodeR68, ConditionalACHActionParametersActionReturnActionCodeR69, ConditionalACHActionParametersActionReturnActionCodeR70, ConditionalACHActionParametersActionReturnActionCodeR71, ConditionalACHActionParametersActionReturnActionCodeR72, ConditionalACHActionParametersActionReturnActionCodeR73, ConditionalACHActionParametersActionReturnActionCodeR74, ConditionalACHActionParametersActionReturnActionCodeR75, ConditionalACHActionParametersActionReturnActionCodeR76, ConditionalACHActionParametersActionReturnActionCodeR77, ConditionalACHActionParametersActionReturnActionCodeR80, ConditionalACHActionParametersActionReturnActionCodeR81, ConditionalACHActionParametersActionReturnActionCodeR82, ConditionalACHActionParametersActionReturnActionCodeR83, ConditionalACHActionParametersActionReturnActionCodeR84, ConditionalACHActionParametersActionReturnActionCodeR85:
		return true
	}
	return false
}

// Return the ACH transaction
type ConditionalACHActionParametersActionReturnActionType string

const (
	ConditionalACHActionParametersActionReturnActionTypeReturn ConditionalACHActionParametersActionReturnActionType = "RETURN"
)

func (r ConditionalACHActionParametersActionReturnActionType) IsKnown() bool {
	switch r {
	case ConditionalACHActionParametersActionReturnActionTypeReturn:
		return true
	}
	return false
}

// Approve the ACH transaction
type ConditionalACHActionParametersActionType string

const (
	ConditionalACHActionParametersActionTypeApprove ConditionalACHActionParametersActionType = "APPROVE"
	ConditionalACHActionParametersActionTypeReturn  ConditionalACHActionParametersActionType = "RETURN"
)

func (r ConditionalACHActionParametersActionType) IsKnown() bool {
	switch r {
	case ConditionalACHActionParametersActionTypeApprove, ConditionalACHActionParametersActionTypeReturn:
		return true
	}
	return false
}

// NACHA return code to use when returning the transaction. Note that the list of
// available return codes is subject to an allowlist configured at the program
// level
type ConditionalACHActionParametersActionCode string

const (
	ConditionalACHActionParametersActionCodeR01 ConditionalACHActionParametersActionCode = "R01"
	ConditionalACHActionParametersActionCodeR02 ConditionalACHActionParametersActionCode = "R02"
	ConditionalACHActionParametersActionCodeR03 ConditionalACHActionParametersActionCode = "R03"
	ConditionalACHActionParametersActionCodeR04 ConditionalACHActionParametersActionCode = "R04"
	ConditionalACHActionParametersActionCodeR05 ConditionalACHActionParametersActionCode = "R05"
	ConditionalACHActionParametersActionCodeR06 ConditionalACHActionParametersActionCode = "R06"
	ConditionalACHActionParametersActionCodeR07 ConditionalACHActionParametersActionCode = "R07"
	ConditionalACHActionParametersActionCodeR08 ConditionalACHActionParametersActionCode = "R08"
	ConditionalACHActionParametersActionCodeR09 ConditionalACHActionParametersActionCode = "R09"
	ConditionalACHActionParametersActionCodeR10 ConditionalACHActionParametersActionCode = "R10"
	ConditionalACHActionParametersActionCodeR11 ConditionalACHActionParametersActionCode = "R11"
	ConditionalACHActionParametersActionCodeR12 ConditionalACHActionParametersActionCode = "R12"
	ConditionalACHActionParametersActionCodeR13 ConditionalACHActionParametersActionCode = "R13"
	ConditionalACHActionParametersActionCodeR14 ConditionalACHActionParametersActionCode = "R14"
	ConditionalACHActionParametersActionCodeR15 ConditionalACHActionParametersActionCode = "R15"
	ConditionalACHActionParametersActionCodeR16 ConditionalACHActionParametersActionCode = "R16"
	ConditionalACHActionParametersActionCodeR17 ConditionalACHActionParametersActionCode = "R17"
	ConditionalACHActionParametersActionCodeR18 ConditionalACHActionParametersActionCode = "R18"
	ConditionalACHActionParametersActionCodeR19 ConditionalACHActionParametersActionCode = "R19"
	ConditionalACHActionParametersActionCodeR20 ConditionalACHActionParametersActionCode = "R20"
	ConditionalACHActionParametersActionCodeR21 ConditionalACHActionParametersActionCode = "R21"
	ConditionalACHActionParametersActionCodeR22 ConditionalACHActionParametersActionCode = "R22"
	ConditionalACHActionParametersActionCodeR23 ConditionalACHActionParametersActionCode = "R23"
	ConditionalACHActionParametersActionCodeR24 ConditionalACHActionParametersActionCode = "R24"
	ConditionalACHActionParametersActionCodeR25 ConditionalACHActionParametersActionCode = "R25"
	ConditionalACHActionParametersActionCodeR26 ConditionalACHActionParametersActionCode = "R26"
	ConditionalACHActionParametersActionCodeR27 ConditionalACHActionParametersActionCode = "R27"
	ConditionalACHActionParametersActionCodeR28 ConditionalACHActionParametersActionCode = "R28"
	ConditionalACHActionParametersActionCodeR29 ConditionalACHActionParametersActionCode = "R29"
	ConditionalACHActionParametersActionCodeR30 ConditionalACHActionParametersActionCode = "R30"
	ConditionalACHActionParametersActionCodeR31 ConditionalACHActionParametersActionCode = "R31"
	ConditionalACHActionParametersActionCodeR32 ConditionalACHActionParametersActionCode = "R32"
	ConditionalACHActionParametersActionCodeR33 ConditionalACHActionParametersActionCode = "R33"
	ConditionalACHActionParametersActionCodeR34 ConditionalACHActionParametersActionCode = "R34"
	ConditionalACHActionParametersActionCodeR35 ConditionalACHActionParametersActionCode = "R35"
	ConditionalACHActionParametersActionCodeR36 ConditionalACHActionParametersActionCode = "R36"
	ConditionalACHActionParametersActionCodeR37 ConditionalACHActionParametersActionCode = "R37"
	ConditionalACHActionParametersActionCodeR38 ConditionalACHActionParametersActionCode = "R38"
	ConditionalACHActionParametersActionCodeR39 ConditionalACHActionParametersActionCode = "R39"
	ConditionalACHActionParametersActionCodeR40 ConditionalACHActionParametersActionCode = "R40"
	ConditionalACHActionParametersActionCodeR41 ConditionalACHActionParametersActionCode = "R41"
	ConditionalACHActionParametersActionCodeR42 ConditionalACHActionParametersActionCode = "R42"
	ConditionalACHActionParametersActionCodeR43 ConditionalACHActionParametersActionCode = "R43"
	ConditionalACHActionParametersActionCodeR44 ConditionalACHActionParametersActionCode = "R44"
	ConditionalACHActionParametersActionCodeR45 ConditionalACHActionParametersActionCode = "R45"
	ConditionalACHActionParametersActionCodeR46 ConditionalACHActionParametersActionCode = "R46"
	ConditionalACHActionParametersActionCodeR47 ConditionalACHActionParametersActionCode = "R47"
	ConditionalACHActionParametersActionCodeR50 ConditionalACHActionParametersActionCode = "R50"
	ConditionalACHActionParametersActionCodeR51 ConditionalACHActionParametersActionCode = "R51"
	ConditionalACHActionParametersActionCodeR52 ConditionalACHActionParametersActionCode = "R52"
	ConditionalACHActionParametersActionCodeR53 ConditionalACHActionParametersActionCode = "R53"
	ConditionalACHActionParametersActionCodeR61 ConditionalACHActionParametersActionCode = "R61"
	ConditionalACHActionParametersActionCodeR62 ConditionalACHActionParametersActionCode = "R62"
	ConditionalACHActionParametersActionCodeR67 ConditionalACHActionParametersActionCode = "R67"
	ConditionalACHActionParametersActionCodeR68 ConditionalACHActionParametersActionCode = "R68"
	ConditionalACHActionParametersActionCodeR69 ConditionalACHActionParametersActionCode = "R69"
	ConditionalACHActionParametersActionCodeR70 ConditionalACHActionParametersActionCode = "R70"
	ConditionalACHActionParametersActionCodeR71 ConditionalACHActionParametersActionCode = "R71"
	ConditionalACHActionParametersActionCodeR72 ConditionalACHActionParametersActionCode = "R72"
	ConditionalACHActionParametersActionCodeR73 ConditionalACHActionParametersActionCode = "R73"
	ConditionalACHActionParametersActionCodeR74 ConditionalACHActionParametersActionCode = "R74"
	ConditionalACHActionParametersActionCodeR75 ConditionalACHActionParametersActionCode = "R75"
	ConditionalACHActionParametersActionCodeR76 ConditionalACHActionParametersActionCode = "R76"
	ConditionalACHActionParametersActionCodeR77 ConditionalACHActionParametersActionCode = "R77"
	ConditionalACHActionParametersActionCodeR80 ConditionalACHActionParametersActionCode = "R80"
	ConditionalACHActionParametersActionCodeR81 ConditionalACHActionParametersActionCode = "R81"
	ConditionalACHActionParametersActionCodeR82 ConditionalACHActionParametersActionCode = "R82"
	ConditionalACHActionParametersActionCodeR83 ConditionalACHActionParametersActionCode = "R83"
	ConditionalACHActionParametersActionCodeR84 ConditionalACHActionParametersActionCode = "R84"
	ConditionalACHActionParametersActionCodeR85 ConditionalACHActionParametersActionCode = "R85"
)

func (r ConditionalACHActionParametersActionCode) IsKnown() bool {
	switch r {
	case ConditionalACHActionParametersActionCodeR01, ConditionalACHActionParametersActionCodeR02, ConditionalACHActionParametersActionCodeR03, ConditionalACHActionParametersActionCodeR04, ConditionalACHActionParametersActionCodeR05, ConditionalACHActionParametersActionCodeR06, ConditionalACHActionParametersActionCodeR07, ConditionalACHActionParametersActionCodeR08, ConditionalACHActionParametersActionCodeR09, ConditionalACHActionParametersActionCodeR10, ConditionalACHActionParametersActionCodeR11, ConditionalACHActionParametersActionCodeR12, ConditionalACHActionParametersActionCodeR13, ConditionalACHActionParametersActionCodeR14, ConditionalACHActionParametersActionCodeR15, ConditionalACHActionParametersActionCodeR16, ConditionalACHActionParametersActionCodeR17, ConditionalACHActionParametersActionCodeR18, ConditionalACHActionParametersActionCodeR19, ConditionalACHActionParametersActionCodeR20, ConditionalACHActionParametersActionCodeR21, ConditionalACHActionParametersActionCodeR22, ConditionalACHActionParametersActionCodeR23, ConditionalACHActionParametersActionCodeR24, ConditionalACHActionParametersActionCodeR25, ConditionalACHActionParametersActionCodeR26, ConditionalACHActionParametersActionCodeR27, ConditionalACHActionParametersActionCodeR28, ConditionalACHActionParametersActionCodeR29, ConditionalACHActionParametersActionCodeR30, ConditionalACHActionParametersActionCodeR31, ConditionalACHActionParametersActionCodeR32, ConditionalACHActionParametersActionCodeR33, ConditionalACHActionParametersActionCodeR34, ConditionalACHActionParametersActionCodeR35, ConditionalACHActionParametersActionCodeR36, ConditionalACHActionParametersActionCodeR37, ConditionalACHActionParametersActionCodeR38, ConditionalACHActionParametersActionCodeR39, ConditionalACHActionParametersActionCodeR40, ConditionalACHActionParametersActionCodeR41, ConditionalACHActionParametersActionCodeR42, ConditionalACHActionParametersActionCodeR43, ConditionalACHActionParametersActionCodeR44, ConditionalACHActionParametersActionCodeR45, ConditionalACHActionParametersActionCodeR46, ConditionalACHActionParametersActionCodeR47, ConditionalACHActionParametersActionCodeR50, ConditionalACHActionParametersActionCodeR51, ConditionalACHActionParametersActionCodeR52, ConditionalACHActionParametersActionCodeR53, ConditionalACHActionParametersActionCodeR61, ConditionalACHActionParametersActionCodeR62, ConditionalACHActionParametersActionCodeR67, ConditionalACHActionParametersActionCodeR68, ConditionalACHActionParametersActionCodeR69, ConditionalACHActionParametersActionCodeR70, ConditionalACHActionParametersActionCodeR71, ConditionalACHActionParametersActionCodeR72, ConditionalACHActionParametersActionCodeR73, ConditionalACHActionParametersActionCodeR74, ConditionalACHActionParametersActionCodeR75, ConditionalACHActionParametersActionCodeR76, ConditionalACHActionParametersActionCodeR77, ConditionalACHActionParametersActionCodeR80, ConditionalACHActionParametersActionCodeR81, ConditionalACHActionParametersActionCodeR82, ConditionalACHActionParametersActionCodeR83, ConditionalACHActionParametersActionCodeR84, ConditionalACHActionParametersActionCodeR85:
		return true
	}
	return false
}

type ConditionalACHActionParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `COMPANY_NAME`: The name of the company initiating the ACH transaction.
	//   - `COMPANY_ID`: The company ID (also known as Standard Entry Class (SEC) Company
	//     ID) of the entity initiating the ACH transaction.
	//   - `TIMESTAMP`: The timestamp of the ACH transaction in ISO 8601 format.
	//   - `TRANSACTION_AMOUNT`: The amount of the ACH transaction in minor units
	//     (cents).
	//   - `SEC_CODE`: Standard Entry Class code indicating the type of ACH transaction.
	//     Valid values include PPD (Prearranged Payment and Deposit Entry), CCD
	//     (Corporate Credit or Debit Entry), WEB (Internet-Initiated/Mobile Entry), TEL
	//     (Telephone-Initiated Entry), and others.
	//   - `MEMO`: Optional memo or description field included with the ACH transaction.
	Attribute ConditionalACHActionParametersConditionsAttribute `json:"attribute" api:"required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion                       `json:"value" api:"required"`
	JSON  conditionalACHActionParametersConditionJSON `json:"-"`
}

// conditionalACHActionParametersConditionJSON contains the JSON metadata for the
// struct [ConditionalACHActionParametersCondition]
type conditionalACHActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalACHActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalACHActionParametersConditionJSON) RawJSON() string {
	return r.raw
}

// The attribute to target.
//
// The following attributes may be targeted:
//
//   - `COMPANY_NAME`: The name of the company initiating the ACH transaction.
//   - `COMPANY_ID`: The company ID (also known as Standard Entry Class (SEC) Company
//     ID) of the entity initiating the ACH transaction.
//   - `TIMESTAMP`: The timestamp of the ACH transaction in ISO 8601 format.
//   - `TRANSACTION_AMOUNT`: The amount of the ACH transaction in minor units
//     (cents).
//   - `SEC_CODE`: Standard Entry Class code indicating the type of ACH transaction.
//     Valid values include PPD (Prearranged Payment and Deposit Entry), CCD
//     (Corporate Credit or Debit Entry), WEB (Internet-Initiated/Mobile Entry), TEL
//     (Telephone-Initiated Entry), and others.
//   - `MEMO`: Optional memo or description field included with the ACH transaction.
type ConditionalACHActionParametersConditionsAttribute string

const (
	ConditionalACHActionParametersConditionsAttributeCompanyName       ConditionalACHActionParametersConditionsAttribute = "COMPANY_NAME"
	ConditionalACHActionParametersConditionsAttributeCompanyID         ConditionalACHActionParametersConditionsAttribute = "COMPANY_ID"
	ConditionalACHActionParametersConditionsAttributeTimestamp         ConditionalACHActionParametersConditionsAttribute = "TIMESTAMP"
	ConditionalACHActionParametersConditionsAttributeTransactionAmount ConditionalACHActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	ConditionalACHActionParametersConditionsAttributeSecCode           ConditionalACHActionParametersConditionsAttribute = "SEC_CODE"
	ConditionalACHActionParametersConditionsAttributeMemo              ConditionalACHActionParametersConditionsAttribute = "MEMO"
)

func (r ConditionalACHActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case ConditionalACHActionParametersConditionsAttributeCompanyName, ConditionalACHActionParametersConditionsAttributeCompanyID, ConditionalACHActionParametersConditionsAttributeTimestamp, ConditionalACHActionParametersConditionsAttributeTransactionAmount, ConditionalACHActionParametersConditionsAttributeSecCode, ConditionalACHActionParametersConditionsAttributeMemo:
		return true
	}
	return false
}

// The attribute to target.
//
// The following attributes may be targeted:
//
//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
//     business by the types of goods or services it provides.
//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
//     Netherlands Antilles.
//   - `CURRENCY`: 3-character alphabetic ISO 4217 code for the merchant currency of
//     the transaction.
//   - `MERCHANT_ID`: Unique alphanumeric identifier for the payment card acceptor
//     (merchant).
//   - `DESCRIPTOR`: Short description of card acceptor.
//   - `LIABILITY_SHIFT`: Indicates whether chargeback liability shift to the issuer
//     applies to the transaction. Valid values are `NONE`, `3DS_AUTHENTICATED`, or
//     `TOKEN_AUTHENTICATED`.
//   - `PAN_ENTRY_MODE`: The method by which the cardholder's primary account number
//     (PAN) was entered. Valid values are `AUTO_ENTRY`, `BAR_CODE`, `CONTACTLESS`,
//     `ECOMMERCE`, `ERROR_KEYED`, `ERROR_MAGNETIC_STRIPE`, `ICC`, `KEY_ENTERED`,
//     `MAGNETIC_STRIPE`, `MANUAL`, `OCR`, `SECURE_CARDLESS`, `UNSPECIFIED`,
//     `UNKNOWN`, `CREDENTIAL_ON_FILE`, or `ECOMMERCE`.
//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
//     fee field in the settlement/cardholder billing currency. This is the amount
//     the issuer should authorize against unless the issuer is paying the acquirer
//     fee on behalf of the cardholder.
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authorization. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `CARD_TRANSACTION_COUNT_15M`: The number of transactions on the card in the
//     trailing 15 minutes before the authorization.
//   - `CARD_TRANSACTION_COUNT_1H`: The number of transactions on the card in the
//     trailing hour up and until the authorization.
//   - `CARD_TRANSACTION_COUNT_24H`: The number of transactions on the card in the
//     trailing 24 hours up and until the authorization.
//   - `CARD_STATE`: The current state of the card associated with the transaction.
//     Valid values are `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`,
//     `PENDING_FULFILLMENT`.
//   - `PIN_ENTERED`: Indicates whether a PIN was entered during the transaction.
//     Valid values are `TRUE`, `FALSE`.
//   - `PIN_STATUS`: The current state of card's PIN. Valid values are `NOT_SET`,
//     `OK`, `BLOCKED`.
//   - `WALLET_TYPE`: For transactions using a digital wallet token, indicates the
//     source of the token. Valid values are `APPLE_PAY`, `GOOGLE_PAY`,
//     `SAMSUNG_PAY`, `MASTERPASS`, `MERCHANT`, `OTHER`, `NONE`.
//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
type ConditionalAttribute string

const (
	ConditionalAttributeMcc                     ConditionalAttribute = "MCC"
	ConditionalAttributeCountry                 ConditionalAttribute = "COUNTRY"
	ConditionalAttributeCurrency                ConditionalAttribute = "CURRENCY"
	ConditionalAttributeMerchantID              ConditionalAttribute = "MERCHANT_ID"
	ConditionalAttributeDescriptor              ConditionalAttribute = "DESCRIPTOR"
	ConditionalAttributeLiabilityShift          ConditionalAttribute = "LIABILITY_SHIFT"
	ConditionalAttributePanEntryMode            ConditionalAttribute = "PAN_ENTRY_MODE"
	ConditionalAttributeTransactionAmount       ConditionalAttribute = "TRANSACTION_AMOUNT"
	ConditionalAttributeRiskScore               ConditionalAttribute = "RISK_SCORE"
	ConditionalAttributeCardTransactionCount15M ConditionalAttribute = "CARD_TRANSACTION_COUNT_15M"
	ConditionalAttributeCardTransactionCount1H  ConditionalAttribute = "CARD_TRANSACTION_COUNT_1H"
	ConditionalAttributeCardTransactionCount24H ConditionalAttribute = "CARD_TRANSACTION_COUNT_24H"
	ConditionalAttributeCardState               ConditionalAttribute = "CARD_STATE"
	ConditionalAttributePinEntered              ConditionalAttribute = "PIN_ENTERED"
	ConditionalAttributePinStatus               ConditionalAttribute = "PIN_STATUS"
	ConditionalAttributeWalletType              ConditionalAttribute = "WALLET_TYPE"
	ConditionalAttributeAddressMatch            ConditionalAttribute = "ADDRESS_MATCH"
)

func (r ConditionalAttribute) IsKnown() bool {
	switch r {
	case ConditionalAttributeMcc, ConditionalAttributeCountry, ConditionalAttributeCurrency, ConditionalAttributeMerchantID, ConditionalAttributeDescriptor, ConditionalAttributeLiabilityShift, ConditionalAttributePanEntryMode, ConditionalAttributeTransactionAmount, ConditionalAttributeRiskScore, ConditionalAttributeCardTransactionCount15M, ConditionalAttributeCardTransactionCount1H, ConditionalAttributeCardTransactionCount24H, ConditionalAttributeCardState, ConditionalAttributePinEntered, ConditionalAttributePinStatus, ConditionalAttributeWalletType, ConditionalAttributeAddressMatch:
		return true
	}
	return false
}

type ConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     ConditionalAuthorizationActionParametersAction      `json:"action" api:"required"`
	Conditions []ConditionalAuthorizationActionParametersCondition `json:"conditions" api:"required"`
	JSON       conditionalAuthorizationActionParametersJSON        `json:"-"`
}

// conditionalAuthorizationActionParametersJSON contains the JSON metadata for the
// struct [ConditionalAuthorizationActionParameters]
type conditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleCurrentVersionParameters() {}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleDraftVersionParameters() {}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleVersionParameters() {}

// The action to take if the conditions are met.
type ConditionalAuthorizationActionParametersAction string

const (
	ConditionalAuthorizationActionParametersActionDecline   ConditionalAuthorizationActionParametersAction = "DECLINE"
	ConditionalAuthorizationActionParametersActionChallenge ConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r ConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case ConditionalAuthorizationActionParametersActionDecline, ConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type ConditionalAuthorizationActionParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-character alphabetic ISO 4217 code for the merchant currency of
	//     the transaction.
	//   - `MERCHANT_ID`: Unique alphanumeric identifier for the payment card acceptor
	//     (merchant).
	//   - `DESCRIPTOR`: Short description of card acceptor.
	//   - `LIABILITY_SHIFT`: Indicates whether chargeback liability shift to the issuer
	//     applies to the transaction. Valid values are `NONE`, `3DS_AUTHENTICATED`, or
	//     `TOKEN_AUTHENTICATED`.
	//   - `PAN_ENTRY_MODE`: The method by which the cardholder's primary account number
	//     (PAN) was entered. Valid values are `AUTO_ENTRY`, `BAR_CODE`, `CONTACTLESS`,
	//     `ECOMMERCE`, `ERROR_KEYED`, `ERROR_MAGNETIC_STRIPE`, `ICC`, `KEY_ENTERED`,
	//     `MAGNETIC_STRIPE`, `MANUAL`, `OCR`, `SECURE_CARDLESS`, `UNSPECIFIED`,
	//     `UNKNOWN`, `CREDENTIAL_ON_FILE`, or `ECOMMERCE`.
	//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
	//     fee field in the settlement/cardholder billing currency. This is the amount
	//     the issuer should authorize against unless the issuer is paying the acquirer
	//     fee on behalf of the cardholder.
	//   - `CASH_AMOUNT`: The cash amount of the transaction in minor units (cents). This
	//     represents the amount of cash being withdrawn or advanced.
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authorization. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `CARD_TRANSACTION_COUNT_15M`: The number of transactions on the card in the
	//     trailing 15 minutes before the authorization.
	//   - `CARD_TRANSACTION_COUNT_1H`: The number of transactions on the card in the
	//     trailing hour up and until the authorization.
	//   - `CARD_TRANSACTION_COUNT_24H`: The number of transactions on the card in the
	//     trailing 24 hours up and until the authorization.
	//   - `CARD_DECLINE_COUNT_15M`: The number of declined transactions on the card in
	//     the trailing 15 minutes before the authorization.
	//   - `CARD_DECLINE_COUNT_1H`: The number of declined transactions on the card in
	//     the trailing hour up and until the authorization.
	//   - `CARD_DECLINE_COUNT_24H`: The number of declined transactions on the card in
	//     the trailing 24 hours up and until the authorization.
	//   - `CARD_STATE`: The current state of the card associated with the transaction.
	//     Valid values are `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`,
	//     `PENDING_FULFILLMENT`.
	//   - `PIN_ENTERED`: Indicates whether a PIN was entered during the transaction.
	//     Valid values are `TRUE`, `FALSE`.
	//   - `PIN_STATUS`: The current state of card's PIN. Valid values are `NOT_SET`,
	//     `OK`, `BLOCKED`.
	//   - `WALLET_TYPE`: For transactions using a digital wallet token, indicates the
	//     source of the token. Valid values are `APPLE_PAY`, `GOOGLE_PAY`,
	//     `SAMSUNG_PAY`, `MASTERPASS`, `MERCHANT`, `OTHER`, `NONE`.
	//   - `TRANSACTION_INITIATOR`: The entity that initiated the transaction indicates
	//     the source of the token. Valid values are `CARDHOLDER`, `MERCHANT`, `UNKNOWN`.
	//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
	//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
	//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
	//   - `SERVICE_LOCATION_STATE`: The state/province code (ISO 3166-2) where the
	//     cardholder received the service, e.g. "NY". When a service location is present
	//     in the network data, the service location state is used. Otherwise, falls back
	//     to the card acceptor state.
	//   - `SERVICE_LOCATION_POSTAL_CODE`: The postal code where the cardholder received
	//     the service, e.g. "10001". When a service location is present in the network
	//     data, the service location postal code is used. Otherwise, falls back to the
	//     card acceptor postal code.
	//   - `CARD_AGE`: The age of the card in seconds at the time of the authorization.
	//   - `ACCOUNT_AGE`: The age of the account holder's account in seconds at the time
	//     of the authorization.
	Attribute ConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute" api:"required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion                                 `json:"value" api:"required"`
	JSON  conditionalAuthorizationActionParametersConditionJSON `json:"-"`
}

// conditionalAuthorizationActionParametersConditionJSON contains the JSON metadata
// for the struct [ConditionalAuthorizationActionParametersCondition]
type conditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
	return r.raw
}

// The attribute to target.
//
// The following attributes may be targeted:
//
//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
//     business by the types of goods or services it provides.
//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
//     Netherlands Antilles.
//   - `CURRENCY`: 3-character alphabetic ISO 4217 code for the merchant currency of
//     the transaction.
//   - `MERCHANT_ID`: Unique alphanumeric identifier for the payment card acceptor
//     (merchant).
//   - `DESCRIPTOR`: Short description of card acceptor.
//   - `LIABILITY_SHIFT`: Indicates whether chargeback liability shift to the issuer
//     applies to the transaction. Valid values are `NONE`, `3DS_AUTHENTICATED`, or
//     `TOKEN_AUTHENTICATED`.
//   - `PAN_ENTRY_MODE`: The method by which the cardholder's primary account number
//     (PAN) was entered. Valid values are `AUTO_ENTRY`, `BAR_CODE`, `CONTACTLESS`,
//     `ECOMMERCE`, `ERROR_KEYED`, `ERROR_MAGNETIC_STRIPE`, `ICC`, `KEY_ENTERED`,
//     `MAGNETIC_STRIPE`, `MANUAL`, `OCR`, `SECURE_CARDLESS`, `UNSPECIFIED`,
//     `UNKNOWN`, `CREDENTIAL_ON_FILE`, or `ECOMMERCE`.
//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
//     fee field in the settlement/cardholder billing currency. This is the amount
//     the issuer should authorize against unless the issuer is paying the acquirer
//     fee on behalf of the cardholder.
//   - `CASH_AMOUNT`: The cash amount of the transaction in minor units (cents). This
//     represents the amount of cash being withdrawn or advanced.
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authorization. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `CARD_TRANSACTION_COUNT_15M`: The number of transactions on the card in the
//     trailing 15 minutes before the authorization.
//   - `CARD_TRANSACTION_COUNT_1H`: The number of transactions on the card in the
//     trailing hour up and until the authorization.
//   - `CARD_TRANSACTION_COUNT_24H`: The number of transactions on the card in the
//     trailing 24 hours up and until the authorization.
//   - `CARD_DECLINE_COUNT_15M`: The number of declined transactions on the card in
//     the trailing 15 minutes before the authorization.
//   - `CARD_DECLINE_COUNT_1H`: The number of declined transactions on the card in
//     the trailing hour up and until the authorization.
//   - `CARD_DECLINE_COUNT_24H`: The number of declined transactions on the card in
//     the trailing 24 hours up and until the authorization.
//   - `CARD_STATE`: The current state of the card associated with the transaction.
//     Valid values are `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`,
//     `PENDING_FULFILLMENT`.
//   - `PIN_ENTERED`: Indicates whether a PIN was entered during the transaction.
//     Valid values are `TRUE`, `FALSE`.
//   - `PIN_STATUS`: The current state of card's PIN. Valid values are `NOT_SET`,
//     `OK`, `BLOCKED`.
//   - `WALLET_TYPE`: For transactions using a digital wallet token, indicates the
//     source of the token. Valid values are `APPLE_PAY`, `GOOGLE_PAY`,
//     `SAMSUNG_PAY`, `MASTERPASS`, `MERCHANT`, `OTHER`, `NONE`.
//   - `TRANSACTION_INITIATOR`: The entity that initiated the transaction indicates
//     the source of the token. Valid values are `CARDHOLDER`, `MERCHANT`, `UNKNOWN`.
//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
//   - `SERVICE_LOCATION_STATE`: The state/province code (ISO 3166-2) where the
//     cardholder received the service, e.g. "NY". When a service location is present
//     in the network data, the service location state is used. Otherwise, falls back
//     to the card acceptor state.
//   - `SERVICE_LOCATION_POSTAL_CODE`: The postal code where the cardholder received
//     the service, e.g. "10001". When a service location is present in the network
//     data, the service location postal code is used. Otherwise, falls back to the
//     card acceptor postal code.
//   - `CARD_AGE`: The age of the card in seconds at the time of the authorization.
//   - `ACCOUNT_AGE`: The age of the account holder's account in seconds at the time
//     of the authorization.
type ConditionalAuthorizationActionParametersConditionsAttribute string

const (
	ConditionalAuthorizationActionParametersConditionsAttributeMcc                       ConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	ConditionalAuthorizationActionParametersConditionsAttributeCountry                   ConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	ConditionalAuthorizationActionParametersConditionsAttributeCurrency                  ConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	ConditionalAuthorizationActionParametersConditionsAttributeMerchantID                ConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	ConditionalAuthorizationActionParametersConditionsAttributeDescriptor                ConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	ConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift            ConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	ConditionalAuthorizationActionParametersConditionsAttributePanEntryMode              ConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	ConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount         ConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	ConditionalAuthorizationActionParametersConditionsAttributeCashAmount                ConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	ConditionalAuthorizationActionParametersConditionsAttributeRiskScore                 ConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M   ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H    ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H   ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	ConditionalAuthorizationActionParametersConditionsAttributeCardDeclineCount15M       ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_DECLINE_COUNT_15M"
	ConditionalAuthorizationActionParametersConditionsAttributeCardDeclineCount1H        ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_DECLINE_COUNT_1H"
	ConditionalAuthorizationActionParametersConditionsAttributeCardDeclineCount24H       ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_DECLINE_COUNT_24H"
	ConditionalAuthorizationActionParametersConditionsAttributeCardState                 ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	ConditionalAuthorizationActionParametersConditionsAttributePinEntered                ConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	ConditionalAuthorizationActionParametersConditionsAttributePinStatus                 ConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	ConditionalAuthorizationActionParametersConditionsAttributeWalletType                ConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	ConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator      ConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
	ConditionalAuthorizationActionParametersConditionsAttributeAddressMatch              ConditionalAuthorizationActionParametersConditionsAttribute = "ADDRESS_MATCH"
	ConditionalAuthorizationActionParametersConditionsAttributeServiceLocationState      ConditionalAuthorizationActionParametersConditionsAttribute = "SERVICE_LOCATION_STATE"
	ConditionalAuthorizationActionParametersConditionsAttributeServiceLocationPostalCode ConditionalAuthorizationActionParametersConditionsAttribute = "SERVICE_LOCATION_POSTAL_CODE"
	ConditionalAuthorizationActionParametersConditionsAttributeCardAge                   ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_AGE"
	ConditionalAuthorizationActionParametersConditionsAttributeAccountAge                ConditionalAuthorizationActionParametersConditionsAttribute = "ACCOUNT_AGE"
)

func (r ConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case ConditionalAuthorizationActionParametersConditionsAttributeMcc, ConditionalAuthorizationActionParametersConditionsAttributeCountry, ConditionalAuthorizationActionParametersConditionsAttributeCurrency, ConditionalAuthorizationActionParametersConditionsAttributeMerchantID, ConditionalAuthorizationActionParametersConditionsAttributeDescriptor, ConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, ConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, ConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, ConditionalAuthorizationActionParametersConditionsAttributeCashAmount, ConditionalAuthorizationActionParametersConditionsAttributeRiskScore, ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, ConditionalAuthorizationActionParametersConditionsAttributeCardDeclineCount15M, ConditionalAuthorizationActionParametersConditionsAttributeCardDeclineCount1H, ConditionalAuthorizationActionParametersConditionsAttributeCardDeclineCount24H, ConditionalAuthorizationActionParametersConditionsAttributeCardState, ConditionalAuthorizationActionParametersConditionsAttributePinEntered, ConditionalAuthorizationActionParametersConditionsAttributePinStatus, ConditionalAuthorizationActionParametersConditionsAttributeWalletType, ConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator, ConditionalAuthorizationActionParametersConditionsAttributeAddressMatch, ConditionalAuthorizationActionParametersConditionsAttributeServiceLocationState, ConditionalAuthorizationActionParametersConditionsAttributeServiceLocationPostalCode, ConditionalAuthorizationActionParametersConditionsAttributeCardAge, ConditionalAuthorizationActionParametersConditionsAttributeAccountAge:
		return true
	}
	return false
}

// Deprecated: Use CONDITIONAL_ACTION instead.
//
// Deprecated: deprecated
type ConditionalBlockParameters struct {
	Conditions []AuthRuleCondition            `json:"conditions" api:"required"`
	JSON       conditionalBlockParametersJSON `json:"-"`
}

// conditionalBlockParametersJSON contains the JSON metadata for the struct
// [ConditionalBlockParameters]
type conditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalBlockParameters) implementsAuthRuleCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleDraftVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleVersionParameters() {}

// The operation to apply to the attribute
type ConditionalOperation string

const (
	ConditionalOperationIsOneOf                ConditionalOperation = "IS_ONE_OF"
	ConditionalOperationIsNotOneOf             ConditionalOperation = "IS_NOT_ONE_OF"
	ConditionalOperationMatches                ConditionalOperation = "MATCHES"
	ConditionalOperationDoesNotMatch           ConditionalOperation = "DOES_NOT_MATCH"
	ConditionalOperationIsEqualTo              ConditionalOperation = "IS_EQUAL_TO"
	ConditionalOperationIsNotEqualTo           ConditionalOperation = "IS_NOT_EQUAL_TO"
	ConditionalOperationIsGreaterThan          ConditionalOperation = "IS_GREATER_THAN"
	ConditionalOperationIsGreaterThanOrEqualTo ConditionalOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	ConditionalOperationIsLessThan             ConditionalOperation = "IS_LESS_THAN"
	ConditionalOperationIsLessThanOrEqualTo    ConditionalOperation = "IS_LESS_THAN_OR_EQUAL_TO"
	ConditionalOperationIsAfter                ConditionalOperation = "IS_AFTER"
	ConditionalOperationIsBefore               ConditionalOperation = "IS_BEFORE"
	ConditionalOperationContainsAny            ConditionalOperation = "CONTAINS_ANY"
	ConditionalOperationContainsAll            ConditionalOperation = "CONTAINS_ALL"
	ConditionalOperationContainsNone           ConditionalOperation = "CONTAINS_NONE"
)

func (r ConditionalOperation) IsKnown() bool {
	switch r {
	case ConditionalOperationIsOneOf, ConditionalOperationIsNotOneOf, ConditionalOperationMatches, ConditionalOperationDoesNotMatch, ConditionalOperationIsEqualTo, ConditionalOperationIsNotEqualTo, ConditionalOperationIsGreaterThan, ConditionalOperationIsGreaterThanOrEqualTo, ConditionalOperationIsLessThan, ConditionalOperationIsLessThanOrEqualTo, ConditionalOperationIsAfter, ConditionalOperationIsBefore, ConditionalOperationContainsAny, ConditionalOperationContainsAll, ConditionalOperationContainsNone:
		return true
	}
	return false
}

type ConditionalTokenizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     ConditionalTokenizationActionParametersAction      `json:"action" api:"required"`
	Conditions []ConditionalTokenizationActionParametersCondition `json:"conditions" api:"required"`
	JSON       conditionalTokenizationActionParametersJSON        `json:"-"`
}

// conditionalTokenizationActionParametersJSON contains the JSON metadata for the
// struct [ConditionalTokenizationActionParameters]
type conditionalTokenizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalTokenizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalTokenizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleCurrentVersionParameters() {}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleDraftVersionParameters() {}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleVersionParameters() {}

// The action to take if the conditions are met.
type ConditionalTokenizationActionParametersAction struct {
	// Decline the tokenization request
	Type ConditionalTokenizationActionParametersActionType `json:"type" api:"required"`
	// Reason code for declining the tokenization request
	Reason ConditionalTokenizationActionParametersActionReason `json:"reason"`
	JSON   conditionalTokenizationActionParametersActionJSON   `json:"-"`
	union  ConditionalTokenizationActionParametersActionUnion
}

// conditionalTokenizationActionParametersActionJSON contains the JSON metadata for
// the struct [ConditionalTokenizationActionParametersAction]
type conditionalTokenizationActionParametersActionJSON struct {
	Type        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r conditionalTokenizationActionParametersActionJSON) RawJSON() string {
	return r.raw
}

func (r *ConditionalTokenizationActionParametersAction) UnmarshalJSON(data []byte) (err error) {
	*r = ConditionalTokenizationActionParametersAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConditionalTokenizationActionParametersActionUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConditionalTokenizationActionParametersActionDeclineActionTokenization],
// [ConditionalTokenizationActionParametersActionRequireTfaAction].
func (r ConditionalTokenizationActionParametersAction) AsUnion() ConditionalTokenizationActionParametersActionUnion {
	return r.union
}

// The action to take if the conditions are met.
//
// Union satisfied by
// [ConditionalTokenizationActionParametersActionDeclineActionTokenization] or
// [ConditionalTokenizationActionParametersActionRequireTfaAction].
type ConditionalTokenizationActionParametersActionUnion interface {
	implementsConditionalTokenizationActionParametersAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConditionalTokenizationActionParametersActionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalTokenizationActionParametersActionDeclineActionTokenization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalTokenizationActionParametersActionRequireTfaAction{}),
		},
	)
}

type ConditionalTokenizationActionParametersActionDeclineActionTokenization struct {
	// Decline the tokenization request
	Type ConditionalTokenizationActionParametersActionDeclineActionTokenizationType `json:"type" api:"required"`
	// Reason code for declining the tokenization request
	Reason ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason `json:"reason"`
	JSON   conditionalTokenizationActionParametersActionDeclineActionTokenizationJSON   `json:"-"`
}

// conditionalTokenizationActionParametersActionDeclineActionTokenizationJSON
// contains the JSON metadata for the struct
// [ConditionalTokenizationActionParametersActionDeclineActionTokenization]
type conditionalTokenizationActionParametersActionDeclineActionTokenizationJSON struct {
	Type        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalTokenizationActionParametersActionDeclineActionTokenization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalTokenizationActionParametersActionDeclineActionTokenizationJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalTokenizationActionParametersActionDeclineActionTokenization) implementsConditionalTokenizationActionParametersAction() {
}

// Decline the tokenization request
type ConditionalTokenizationActionParametersActionDeclineActionTokenizationType string

const (
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationTypeDecline ConditionalTokenizationActionParametersActionDeclineActionTokenizationType = "DECLINE"
)

func (r ConditionalTokenizationActionParametersActionDeclineActionTokenizationType) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersActionDeclineActionTokenizationTypeDecline:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason string

const (
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonAccountScore1                  ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "ACCOUNT_SCORE_1"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonDeviceScore1                   ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "DEVICE_SCORE_1"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonAllWalletDeclineReasonsPresent ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonWalletRecommendedDecisionRed   ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "WALLET_RECOMMENDED_DECISION_RED"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCvcMismatch                    ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "CVC_MISMATCH"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCardExpiryMonthMismatch        ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "CARD_EXPIRY_MONTH_MISMATCH"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCardExpiryYearMismatch         ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "CARD_EXPIRY_YEAR_MISMATCH"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCardInvalidState               ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "CARD_INVALID_STATE"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCustomerRedPath                ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "CUSTOMER_RED_PATH"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonInvalidCustomerResponse        ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "INVALID_CUSTOMER_RESPONSE"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonNetworkFailure                 ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "NETWORK_FAILURE"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonGenericDecline                 ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "GENERIC_DECLINE"
	ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonDigitalCardArtRequired         ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason = "DIGITAL_CARD_ART_REQUIRED"
)

func (r ConditionalTokenizationActionParametersActionDeclineActionTokenizationReason) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonAccountScore1, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonDeviceScore1, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonAllWalletDeclineReasonsPresent, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonWalletRecommendedDecisionRed, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCvcMismatch, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCardExpiryMonthMismatch, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCardExpiryYearMismatch, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCardInvalidState, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonCustomerRedPath, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonInvalidCustomerResponse, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonNetworkFailure, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonGenericDecline, ConditionalTokenizationActionParametersActionDeclineActionTokenizationReasonDigitalCardArtRequired:
		return true
	}
	return false
}

type ConditionalTokenizationActionParametersActionRequireTfaAction struct {
	// Require two-factor authentication for the tokenization request
	Type ConditionalTokenizationActionParametersActionRequireTfaActionType `json:"type" api:"required"`
	// Reason code for requiring two-factor authentication
	Reason ConditionalTokenizationActionParametersActionRequireTfaActionReason `json:"reason"`
	JSON   conditionalTokenizationActionParametersActionRequireTfaActionJSON   `json:"-"`
}

// conditionalTokenizationActionParametersActionRequireTfaActionJSON contains the
// JSON metadata for the struct
// [ConditionalTokenizationActionParametersActionRequireTfaAction]
type conditionalTokenizationActionParametersActionRequireTfaActionJSON struct {
	Type        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalTokenizationActionParametersActionRequireTfaAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalTokenizationActionParametersActionRequireTfaActionJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalTokenizationActionParametersActionRequireTfaAction) implementsConditionalTokenizationActionParametersAction() {
}

// Require two-factor authentication for the tokenization request
type ConditionalTokenizationActionParametersActionRequireTfaActionType string

const (
	ConditionalTokenizationActionParametersActionRequireTfaActionTypeRequireTfa ConditionalTokenizationActionParametersActionRequireTfaActionType = "REQUIRE_TFA"
)

func (r ConditionalTokenizationActionParametersActionRequireTfaActionType) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersActionRequireTfaActionTypeRequireTfa:
		return true
	}
	return false
}

// Reason code for requiring two-factor authentication
type ConditionalTokenizationActionParametersActionRequireTfaActionReason string

const (
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonWalletRecommendedTfa        ConditionalTokenizationActionParametersActionRequireTfaActionReason = "WALLET_RECOMMENDED_TFA"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonSuspiciousActivity          ConditionalTokenizationActionParametersActionRequireTfaActionReason = "SUSPICIOUS_ACTIVITY"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonDeviceRecentlyLost          ConditionalTokenizationActionParametersActionRequireTfaActionReason = "DEVICE_RECENTLY_LOST"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonTooManyRecentAttempts       ConditionalTokenizationActionParametersActionRequireTfaActionReason = "TOO_MANY_RECENT_ATTEMPTS"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonTooManyRecentTokens         ConditionalTokenizationActionParametersActionRequireTfaActionReason = "TOO_MANY_RECENT_TOKENS"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonTooManyDifferentCardholders ConditionalTokenizationActionParametersActionRequireTfaActionReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonOutsideHomeTerritory        ConditionalTokenizationActionParametersActionRequireTfaActionReason = "OUTSIDE_HOME_TERRITORY"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonHasSuspendedTokens          ConditionalTokenizationActionParametersActionRequireTfaActionReason = "HAS_SUSPENDED_TOKENS"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonHighRisk                    ConditionalTokenizationActionParametersActionRequireTfaActionReason = "HIGH_RISK"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonAccountScoreLow             ConditionalTokenizationActionParametersActionRequireTfaActionReason = "ACCOUNT_SCORE_LOW"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonDeviceScoreLow              ConditionalTokenizationActionParametersActionRequireTfaActionReason = "DEVICE_SCORE_LOW"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonCardStateTfa                ConditionalTokenizationActionParametersActionRequireTfaActionReason = "CARD_STATE_TFA"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonHardcodedTfa                ConditionalTokenizationActionParametersActionRequireTfaActionReason = "HARDCODED_TFA"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonCustomerRuleTfa             ConditionalTokenizationActionParametersActionRequireTfaActionReason = "CUSTOMER_RULE_TFA"
	ConditionalTokenizationActionParametersActionRequireTfaActionReasonDeviceHostCardEmulation     ConditionalTokenizationActionParametersActionRequireTfaActionReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r ConditionalTokenizationActionParametersActionRequireTfaActionReason) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersActionRequireTfaActionReasonWalletRecommendedTfa, ConditionalTokenizationActionParametersActionRequireTfaActionReasonSuspiciousActivity, ConditionalTokenizationActionParametersActionRequireTfaActionReasonDeviceRecentlyLost, ConditionalTokenizationActionParametersActionRequireTfaActionReasonTooManyRecentAttempts, ConditionalTokenizationActionParametersActionRequireTfaActionReasonTooManyRecentTokens, ConditionalTokenizationActionParametersActionRequireTfaActionReasonTooManyDifferentCardholders, ConditionalTokenizationActionParametersActionRequireTfaActionReasonOutsideHomeTerritory, ConditionalTokenizationActionParametersActionRequireTfaActionReasonHasSuspendedTokens, ConditionalTokenizationActionParametersActionRequireTfaActionReasonHighRisk, ConditionalTokenizationActionParametersActionRequireTfaActionReasonAccountScoreLow, ConditionalTokenizationActionParametersActionRequireTfaActionReasonDeviceScoreLow, ConditionalTokenizationActionParametersActionRequireTfaActionReasonCardStateTfa, ConditionalTokenizationActionParametersActionRequireTfaActionReasonHardcodedTfa, ConditionalTokenizationActionParametersActionRequireTfaActionReasonCustomerRuleTfa, ConditionalTokenizationActionParametersActionRequireTfaActionReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

// Decline the tokenization request
type ConditionalTokenizationActionParametersActionType string

const (
	ConditionalTokenizationActionParametersActionTypeDecline    ConditionalTokenizationActionParametersActionType = "DECLINE"
	ConditionalTokenizationActionParametersActionTypeRequireTfa ConditionalTokenizationActionParametersActionType = "REQUIRE_TFA"
)

func (r ConditionalTokenizationActionParametersActionType) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersActionTypeDecline, ConditionalTokenizationActionParametersActionTypeRequireTfa:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type ConditionalTokenizationActionParametersActionReason string

const (
	ConditionalTokenizationActionParametersActionReasonAccountScore1                  ConditionalTokenizationActionParametersActionReason = "ACCOUNT_SCORE_1"
	ConditionalTokenizationActionParametersActionReasonDeviceScore1                   ConditionalTokenizationActionParametersActionReason = "DEVICE_SCORE_1"
	ConditionalTokenizationActionParametersActionReasonAllWalletDeclineReasonsPresent ConditionalTokenizationActionParametersActionReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	ConditionalTokenizationActionParametersActionReasonWalletRecommendedDecisionRed   ConditionalTokenizationActionParametersActionReason = "WALLET_RECOMMENDED_DECISION_RED"
	ConditionalTokenizationActionParametersActionReasonCvcMismatch                    ConditionalTokenizationActionParametersActionReason = "CVC_MISMATCH"
	ConditionalTokenizationActionParametersActionReasonCardExpiryMonthMismatch        ConditionalTokenizationActionParametersActionReason = "CARD_EXPIRY_MONTH_MISMATCH"
	ConditionalTokenizationActionParametersActionReasonCardExpiryYearMismatch         ConditionalTokenizationActionParametersActionReason = "CARD_EXPIRY_YEAR_MISMATCH"
	ConditionalTokenizationActionParametersActionReasonCardInvalidState               ConditionalTokenizationActionParametersActionReason = "CARD_INVALID_STATE"
	ConditionalTokenizationActionParametersActionReasonCustomerRedPath                ConditionalTokenizationActionParametersActionReason = "CUSTOMER_RED_PATH"
	ConditionalTokenizationActionParametersActionReasonInvalidCustomerResponse        ConditionalTokenizationActionParametersActionReason = "INVALID_CUSTOMER_RESPONSE"
	ConditionalTokenizationActionParametersActionReasonNetworkFailure                 ConditionalTokenizationActionParametersActionReason = "NETWORK_FAILURE"
	ConditionalTokenizationActionParametersActionReasonGenericDecline                 ConditionalTokenizationActionParametersActionReason = "GENERIC_DECLINE"
	ConditionalTokenizationActionParametersActionReasonDigitalCardArtRequired         ConditionalTokenizationActionParametersActionReason = "DIGITAL_CARD_ART_REQUIRED"
	ConditionalTokenizationActionParametersActionReasonWalletRecommendedTfa           ConditionalTokenizationActionParametersActionReason = "WALLET_RECOMMENDED_TFA"
	ConditionalTokenizationActionParametersActionReasonSuspiciousActivity             ConditionalTokenizationActionParametersActionReason = "SUSPICIOUS_ACTIVITY"
	ConditionalTokenizationActionParametersActionReasonDeviceRecentlyLost             ConditionalTokenizationActionParametersActionReason = "DEVICE_RECENTLY_LOST"
	ConditionalTokenizationActionParametersActionReasonTooManyRecentAttempts          ConditionalTokenizationActionParametersActionReason = "TOO_MANY_RECENT_ATTEMPTS"
	ConditionalTokenizationActionParametersActionReasonTooManyRecentTokens            ConditionalTokenizationActionParametersActionReason = "TOO_MANY_RECENT_TOKENS"
	ConditionalTokenizationActionParametersActionReasonTooManyDifferentCardholders    ConditionalTokenizationActionParametersActionReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	ConditionalTokenizationActionParametersActionReasonOutsideHomeTerritory           ConditionalTokenizationActionParametersActionReason = "OUTSIDE_HOME_TERRITORY"
	ConditionalTokenizationActionParametersActionReasonHasSuspendedTokens             ConditionalTokenizationActionParametersActionReason = "HAS_SUSPENDED_TOKENS"
	ConditionalTokenizationActionParametersActionReasonHighRisk                       ConditionalTokenizationActionParametersActionReason = "HIGH_RISK"
	ConditionalTokenizationActionParametersActionReasonAccountScoreLow                ConditionalTokenizationActionParametersActionReason = "ACCOUNT_SCORE_LOW"
	ConditionalTokenizationActionParametersActionReasonDeviceScoreLow                 ConditionalTokenizationActionParametersActionReason = "DEVICE_SCORE_LOW"
	ConditionalTokenizationActionParametersActionReasonCardStateTfa                   ConditionalTokenizationActionParametersActionReason = "CARD_STATE_TFA"
	ConditionalTokenizationActionParametersActionReasonHardcodedTfa                   ConditionalTokenizationActionParametersActionReason = "HARDCODED_TFA"
	ConditionalTokenizationActionParametersActionReasonCustomerRuleTfa                ConditionalTokenizationActionParametersActionReason = "CUSTOMER_RULE_TFA"
	ConditionalTokenizationActionParametersActionReasonDeviceHostCardEmulation        ConditionalTokenizationActionParametersActionReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r ConditionalTokenizationActionParametersActionReason) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersActionReasonAccountScore1, ConditionalTokenizationActionParametersActionReasonDeviceScore1, ConditionalTokenizationActionParametersActionReasonAllWalletDeclineReasonsPresent, ConditionalTokenizationActionParametersActionReasonWalletRecommendedDecisionRed, ConditionalTokenizationActionParametersActionReasonCvcMismatch, ConditionalTokenizationActionParametersActionReasonCardExpiryMonthMismatch, ConditionalTokenizationActionParametersActionReasonCardExpiryYearMismatch, ConditionalTokenizationActionParametersActionReasonCardInvalidState, ConditionalTokenizationActionParametersActionReasonCustomerRedPath, ConditionalTokenizationActionParametersActionReasonInvalidCustomerResponse, ConditionalTokenizationActionParametersActionReasonNetworkFailure, ConditionalTokenizationActionParametersActionReasonGenericDecline, ConditionalTokenizationActionParametersActionReasonDigitalCardArtRequired, ConditionalTokenizationActionParametersActionReasonWalletRecommendedTfa, ConditionalTokenizationActionParametersActionReasonSuspiciousActivity, ConditionalTokenizationActionParametersActionReasonDeviceRecentlyLost, ConditionalTokenizationActionParametersActionReasonTooManyRecentAttempts, ConditionalTokenizationActionParametersActionReasonTooManyRecentTokens, ConditionalTokenizationActionParametersActionReasonTooManyDifferentCardholders, ConditionalTokenizationActionParametersActionReasonOutsideHomeTerritory, ConditionalTokenizationActionParametersActionReasonHasSuspendedTokens, ConditionalTokenizationActionParametersActionReasonHighRisk, ConditionalTokenizationActionParametersActionReasonAccountScoreLow, ConditionalTokenizationActionParametersActionReasonDeviceScoreLow, ConditionalTokenizationActionParametersActionReasonCardStateTfa, ConditionalTokenizationActionParametersActionReasonHardcodedTfa, ConditionalTokenizationActionParametersActionReasonCustomerRuleTfa, ConditionalTokenizationActionParametersActionReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

type ConditionalTokenizationActionParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `TIMESTAMP`: The timestamp of the tokenization request in ISO 8601 format.
	//   - `TOKENIZATION_CHANNEL`: The channel through which the tokenization request was
	//     initiated. Valid values are `DIGITAL_WALLET`, `MERCHANT`.
	//   - `TOKENIZATION_SOURCE`: The source of the tokenization request. Valid values
	//     are `ACCOUNT_ON_FILE`, `MANUAL_PROVISION`, `PUSH_PROVISION`, `CHIP_DIP`,
	//     `CONTACTLESS_TAP`, `TOKEN`, `UNKNOWN`.
	//   - `TOKEN_REQUESTOR_NAME`: The name of the entity requesting the token. Valid
	//     values are `ALT_ID`, `AMAZON_ONE`, `AMERICAN_EXPRESS_TOKEN_SERVICE`,
	//     `ANDROID_PAY`, `APPLE_PAY`, `FACEBOOK`, `FITBIT_PAY`, `GARMIN_PAY`,
	//     `GOOGLE_PAY`, `GOOGLE_VCN`, `ISSUER_HCE`, `MICROSOFT_PAY`, `NETFLIX`,
	//     `SAMSUNG_PAY`, `UNKNOWN`, `VISA_CHECKOUT`.
	//   - `WALLET_ACCOUNT_SCORE`: Risk score for the account in the digital wallet.
	//     Numeric value where lower numbers indicate higher risk (e.g., 1 = high risk, 2
	//     = medium risk).
	//   - `WALLET_DEVICE_SCORE`: Risk score for the device in the digital wallet.
	//     Numeric value where lower numbers indicate higher risk (e.g., 1 = high risk, 2
	//     = medium risk).
	//   - `WALLET_RECOMMENDED_DECISION`: The decision recommended by the digital wallet
	//     provider. Valid values include APPROVE, DECLINE,
	//     REQUIRE_ADDITIONAL_AUTHENTICATION.
	//   - `WALLET_RECOMMENDATION_REASONS`: List of reasons provided by the digital
	//     wallet provider for the recommended decision. Valid values are:
	//   - Common: `ACCOUNT_CARD_TOO_NEW`, `ACCOUNT_RECENTLY_CHANGED`,
	//     `ACCOUNT_TOO_NEW`, `ACCOUNT_TOO_NEW_SINCE_LAUNCH`, `DEVICE_RECENTLY_LOST`,
	//     `HAS_SUSPENDED_TOKENS`, `HIGH_RISK`, `INACTIVE_ACCOUNT`,
	//     `LOW_ACCOUNT_SCORE`, `LOW_DEVICE_SCORE`, `OUTSIDE_HOME_TERRITORY`,
	//     `SUSPICIOUS_ACTIVITY`, `TOO_MANY_DIFFERENT_CARDHOLDERS`,
	//     `TOO_MANY_RECENT_ATTEMPTS`, `TOO_MANY_RECENT_TOKENS`, `UNABLE_TO_ASSESS`
	//   - Visa only: `ACCOUNT_DATA_RECENTLY_CHANGED`, `ACCOUNT_PAN_PAIRING_TOO_NEW`,
	//     `LOW_TRANSACTION_VOLUME`, `USER_ACCOUNT_DEVICE_TOO_NEW`,
	//     `WALLET_ACCOUNT_TOO_NEW`
	//   - Amex only: `DEVICE_USING_VPN_PROXY`,
	//     `EXCESSIVE_BILLING_NAME_ATTEMPTS_MODERATE`,
	//     `EXCESSIVE_BILLING_NAME_ATTEMPTS_SEVERE`,
	//     `EXCESSIVE_CARD_PROVISION_ATTEMPTS_MODERATE`,
	//     `EXCESSIVE_CARD_PROVISION_ATTEMPTS_SEVERE`, `EXCESSIVE_WALLET_RESETS`,
	//     `EXCESSIVE_ZIP_ATTEMPTS_MODERATE`, `EXCESSIVE_ZIP_ATTEMPTS_SEVERE`,
	//     `USER_ID_CARD_PAIRING_TOO_NEW`, `USER_ID_DEVICE_ID_PAIRING_TOO_NEW`,
	//     `USER_ID_OS_ID_PAIRING_TOO_NEW`, `USER_ID_TOO_NEW`
	//   - `TOKEN_REQUESTOR_ID`: Unique identifier for the entity requesting the token.
	//   - `WALLET_TOKEN_STATUS`: The current status of the wallet token.
	//   - `CARD_STATE`: The state of the card being tokenized. Valid values are
	//     `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`, `PENDING_FULFILLMENT`.
	Attribute ConditionalTokenizationActionParametersConditionsAttribute `json:"attribute" api:"required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion                                `json:"value" api:"required"`
	JSON  conditionalTokenizationActionParametersConditionJSON `json:"-"`
}

// conditionalTokenizationActionParametersConditionJSON contains the JSON metadata
// for the struct [ConditionalTokenizationActionParametersCondition]
type conditionalTokenizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalTokenizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalTokenizationActionParametersConditionJSON) RawJSON() string {
	return r.raw
}

// The attribute to target.
//
// The following attributes may be targeted:
//
//   - `TIMESTAMP`: The timestamp of the tokenization request in ISO 8601 format.
//   - `TOKENIZATION_CHANNEL`: The channel through which the tokenization request was
//     initiated. Valid values are `DIGITAL_WALLET`, `MERCHANT`.
//   - `TOKENIZATION_SOURCE`: The source of the tokenization request. Valid values
//     are `ACCOUNT_ON_FILE`, `MANUAL_PROVISION`, `PUSH_PROVISION`, `CHIP_DIP`,
//     `CONTACTLESS_TAP`, `TOKEN`, `UNKNOWN`.
//   - `TOKEN_REQUESTOR_NAME`: The name of the entity requesting the token. Valid
//     values are `ALT_ID`, `AMAZON_ONE`, `AMERICAN_EXPRESS_TOKEN_SERVICE`,
//     `ANDROID_PAY`, `APPLE_PAY`, `FACEBOOK`, `FITBIT_PAY`, `GARMIN_PAY`,
//     `GOOGLE_PAY`, `GOOGLE_VCN`, `ISSUER_HCE`, `MICROSOFT_PAY`, `NETFLIX`,
//     `SAMSUNG_PAY`, `UNKNOWN`, `VISA_CHECKOUT`.
//   - `WALLET_ACCOUNT_SCORE`: Risk score for the account in the digital wallet.
//     Numeric value where lower numbers indicate higher risk (e.g., 1 = high risk, 2
//     = medium risk).
//   - `WALLET_DEVICE_SCORE`: Risk score for the device in the digital wallet.
//     Numeric value where lower numbers indicate higher risk (e.g., 1 = high risk, 2
//     = medium risk).
//   - `WALLET_RECOMMENDED_DECISION`: The decision recommended by the digital wallet
//     provider. Valid values include APPROVE, DECLINE,
//     REQUIRE_ADDITIONAL_AUTHENTICATION.
//   - `WALLET_RECOMMENDATION_REASONS`: List of reasons provided by the digital
//     wallet provider for the recommended decision. Valid values are:
//   - Common: `ACCOUNT_CARD_TOO_NEW`, `ACCOUNT_RECENTLY_CHANGED`,
//     `ACCOUNT_TOO_NEW`, `ACCOUNT_TOO_NEW_SINCE_LAUNCH`, `DEVICE_RECENTLY_LOST`,
//     `HAS_SUSPENDED_TOKENS`, `HIGH_RISK`, `INACTIVE_ACCOUNT`,
//     `LOW_ACCOUNT_SCORE`, `LOW_DEVICE_SCORE`, `OUTSIDE_HOME_TERRITORY`,
//     `SUSPICIOUS_ACTIVITY`, `TOO_MANY_DIFFERENT_CARDHOLDERS`,
//     `TOO_MANY_RECENT_ATTEMPTS`, `TOO_MANY_RECENT_TOKENS`, `UNABLE_TO_ASSESS`
//   - Visa only: `ACCOUNT_DATA_RECENTLY_CHANGED`, `ACCOUNT_PAN_PAIRING_TOO_NEW`,
//     `LOW_TRANSACTION_VOLUME`, `USER_ACCOUNT_DEVICE_TOO_NEW`,
//     `WALLET_ACCOUNT_TOO_NEW`
//   - Amex only: `DEVICE_USING_VPN_PROXY`,
//     `EXCESSIVE_BILLING_NAME_ATTEMPTS_MODERATE`,
//     `EXCESSIVE_BILLING_NAME_ATTEMPTS_SEVERE`,
//     `EXCESSIVE_CARD_PROVISION_ATTEMPTS_MODERATE`,
//     `EXCESSIVE_CARD_PROVISION_ATTEMPTS_SEVERE`, `EXCESSIVE_WALLET_RESETS`,
//     `EXCESSIVE_ZIP_ATTEMPTS_MODERATE`, `EXCESSIVE_ZIP_ATTEMPTS_SEVERE`,
//     `USER_ID_CARD_PAIRING_TOO_NEW`, `USER_ID_DEVICE_ID_PAIRING_TOO_NEW`,
//     `USER_ID_OS_ID_PAIRING_TOO_NEW`, `USER_ID_TOO_NEW`
//   - `TOKEN_REQUESTOR_ID`: Unique identifier for the entity requesting the token.
//   - `WALLET_TOKEN_STATUS`: The current status of the wallet token.
//   - `CARD_STATE`: The state of the card being tokenized. Valid values are
//     `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`, `PENDING_FULFILLMENT`.
type ConditionalTokenizationActionParametersConditionsAttribute string

const (
	ConditionalTokenizationActionParametersConditionsAttributeTimestamp                   ConditionalTokenizationActionParametersConditionsAttribute = "TIMESTAMP"
	ConditionalTokenizationActionParametersConditionsAttributeTokenizationChannel         ConditionalTokenizationActionParametersConditionsAttribute = "TOKENIZATION_CHANNEL"
	ConditionalTokenizationActionParametersConditionsAttributeTokenizationSource          ConditionalTokenizationActionParametersConditionsAttribute = "TOKENIZATION_SOURCE"
	ConditionalTokenizationActionParametersConditionsAttributeTokenRequestorName          ConditionalTokenizationActionParametersConditionsAttribute = "TOKEN_REQUESTOR_NAME"
	ConditionalTokenizationActionParametersConditionsAttributeWalletAccountScore          ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_ACCOUNT_SCORE"
	ConditionalTokenizationActionParametersConditionsAttributeWalletDeviceScore           ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_DEVICE_SCORE"
	ConditionalTokenizationActionParametersConditionsAttributeWalletRecommendedDecision   ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_RECOMMENDED_DECISION"
	ConditionalTokenizationActionParametersConditionsAttributeWalletRecommendationReasons ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_RECOMMENDATION_REASONS"
	ConditionalTokenizationActionParametersConditionsAttributeTokenRequestorID            ConditionalTokenizationActionParametersConditionsAttribute = "TOKEN_REQUESTOR_ID"
	ConditionalTokenizationActionParametersConditionsAttributeWalletTokenStatus           ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_TOKEN_STATUS"
	ConditionalTokenizationActionParametersConditionsAttributeCardState                   ConditionalTokenizationActionParametersConditionsAttribute = "CARD_STATE"
)

func (r ConditionalTokenizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersConditionsAttributeTimestamp, ConditionalTokenizationActionParametersConditionsAttributeTokenizationChannel, ConditionalTokenizationActionParametersConditionsAttributeTokenizationSource, ConditionalTokenizationActionParametersConditionsAttributeTokenRequestorName, ConditionalTokenizationActionParametersConditionsAttributeWalletAccountScore, ConditionalTokenizationActionParametersConditionsAttributeWalletDeviceScore, ConditionalTokenizationActionParametersConditionsAttributeWalletRecommendedDecision, ConditionalTokenizationActionParametersConditionsAttributeWalletRecommendationReasons, ConditionalTokenizationActionParametersConditionsAttributeTokenRequestorID, ConditionalTokenizationActionParametersConditionsAttributeWalletTokenStatus, ConditionalTokenizationActionParametersConditionsAttributeCardState:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt],
// [ConditionalValueListOfStrings] or [shared.UnionTime].
type ConditionalValueUnion interface {
	ImplementsConditionalValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConditionalValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalValueListOfStrings{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
	)
}

type ConditionalValueListOfStrings []string

func (r ConditionalValueListOfStrings) ImplementsConditionalValueUnion() {}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [ConditionalValueListOfStringsParam], [shared.UnionTime].
type ConditionalValueUnionParam interface {
	ImplementsConditionalValueUnionParam()
}

type ConditionalValueListOfStringsParam []string

func (r ConditionalValueListOfStringsParam) ImplementsConditionalValueUnionParam() {}

// The event stream during which the rule will be evaluated.
type EventStream string

const (
	EventStreamAuthorization         EventStream = "AUTHORIZATION"
	EventStreamThreeDSAuthentication EventStream = "THREE_DS_AUTHENTICATION"
	EventStreamTokenization          EventStream = "TOKENIZATION"
	EventStreamACHCreditReceipt      EventStream = "ACH_CREDIT_RECEIPT"
	EventStreamACHDebitReceipt       EventStream = "ACH_DEBIT_RECEIPT"
)

func (r EventStream) IsKnown() bool {
	switch r {
	case EventStreamAuthorization, EventStreamThreeDSAuthentication, EventStreamTokenization, EventStreamACHCreditReceipt, EventStreamACHDebitReceipt:
		return true
	}
	return false
}

type MerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []MerchantLockParametersMerchant `json:"merchants" api:"required"`
	JSON      merchantLockParametersJSON       `json:"-"`
}

// merchantLockParametersJSON contains the JSON metadata for the struct
// [MerchantLockParameters]
type merchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r merchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r MerchantLockParameters) implementsAuthRuleCurrentVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleDraftVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleVersionParameters() {}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type MerchantLockParametersMerchant struct {
	// A comment or explanation about the merchant, used internally for rule management
	// purposes.
	Comment string `json:"comment"`
	// Short description of the merchant, often used to provide more human-readable
	// context about the transaction merchant. This is typically the name or label
	// shown on transaction summaries.
	Descriptor string `json:"descriptor"`
	// Unique alphanumeric identifier for the payment card acceptor (merchant). This
	// attribute specifies the merchant entity that will be locked or referenced for
	// authorization rules.
	MerchantID string                             `json:"merchant_id"`
	JSON       merchantLockParametersMerchantJSON `json:"-"`
}

// merchantLockParametersMerchantJSON contains the JSON metadata for the struct
// [MerchantLockParametersMerchant]
type merchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r merchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type ReportStats struct {
	// A mapping of action types to the number of times that action was returned by
	// this rule during the relevant period. Actions are the possible outcomes of a
	// rule evaluation, such as DECLINE, CHALLENGE, REQUIRE_TFA, etc. In case rule
	// didn't trigger any action, it's counted under NO_ACTION key.
	ActionCounts map[string]int64 `json:"action_counts"`
	// The total number of historical transactions approved by this rule during the
	// relevant period, or the number of transactions that would have been approved if
	// the rule was evaluated in shadow mode.
	//
	// Deprecated: deprecated
	Approved int64 `json:"approved"`
	// The total number of historical transactions challenged by this rule during the
	// relevant period, or the number of transactions that would have been challenged
	// if the rule was evaluated in shadow mode. Currently applicable only for 3DS Auth
	// Rules.
	//
	// Deprecated: deprecated
	Challenged int64 `json:"challenged"`
	// The total number of historical transactions declined by this rule during the
	// relevant period, or the number of transactions that would have been declined if
	// the rule was evaluated in shadow mode.
	//
	// Deprecated: deprecated
	Declined int64 `json:"declined"`
	// Example events and their outcomes.
	Examples []ReportStatsExample `json:"examples"`
	JSON     reportStatsJSON      `json:"-"`
}

// reportStatsJSON contains the JSON metadata for the struct [ReportStats]
type reportStatsJSON struct {
	ActionCounts apijson.Field
	Approved     apijson.Field
	Challenged   apijson.Field
	Declined     apijson.Field
	Examples     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ReportStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsJSON) RawJSON() string {
	return r.raw
}

type ReportStatsExample struct {
	// The actions taken by the rule for this event.
	Actions []ReportStatsExamplesAction `json:"actions"`
	// Whether the rule would have approved the request.
	//
	// Deprecated: deprecated
	Approved bool `json:"approved"`
	// The decision made by the rule for this event.
	//
	// Deprecated: deprecated
	Decision ReportStatsExamplesDecision `json:"decision"`
	// The event token.
	EventToken string `json:"event_token" format:"uuid"`
	// The timestamp of the event.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// The token of the transaction associated with the event
	TransactionToken string                 `json:"transaction_token" api:"nullable" format:"uuid"`
	JSON             reportStatsExampleJSON `json:"-"`
}

// reportStatsExampleJSON contains the JSON metadata for the struct
// [ReportStatsExample]
type reportStatsExampleJSON struct {
	Actions          apijson.Field
	Approved         apijson.Field
	Decision         apijson.Field
	EventToken       apijson.Field
	Timestamp        apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ReportStatsExample) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsExampleJSON) RawJSON() string {
	return r.raw
}

type ReportStatsExamplesAction struct {
	Type ReportStatsExamplesActionsType `json:"type" api:"required"`
	// The detailed result code explaining the specific reason for the decline
	Code ReportStatsExamplesActionsCode `json:"code"`
	// Reason code for declining the tokenization request
	Reason ReportStatsExamplesActionsReason `json:"reason"`
	JSON   reportStatsExamplesActionJSON    `json:"-"`
	union  ReportStatsExamplesActionsUnion
}

// reportStatsExamplesActionJSON contains the JSON metadata for the struct
// [ReportStatsExamplesAction]
type reportStatsExamplesActionJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r reportStatsExamplesActionJSON) RawJSON() string {
	return r.raw
}

func (r *ReportStatsExamplesAction) UnmarshalJSON(data []byte) (err error) {
	*r = ReportStatsExamplesAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ReportStatsExamplesActionsUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ReportStatsExamplesActionsDeclineActionAuthorization],
// [ReportStatsExamplesActionsChallengeActionAuthorization],
// [ReportStatsExamplesActionsResultAuthentication3DSAction],
// [ReportStatsExamplesActionsDeclineActionTokenization],
// [ReportStatsExamplesActionsRequireTfaAction],
// [ReportStatsExamplesActionsApproveActionACH],
// [ReportStatsExamplesActionsReturnAction].
func (r ReportStatsExamplesAction) AsUnion() ReportStatsExamplesActionsUnion {
	return r.union
}

// Union satisfied by [ReportStatsExamplesActionsDeclineActionAuthorization],
// [ReportStatsExamplesActionsChallengeActionAuthorization],
// [ReportStatsExamplesActionsResultAuthentication3DSAction],
// [ReportStatsExamplesActionsDeclineActionTokenization],
// [ReportStatsExamplesActionsRequireTfaAction],
// [ReportStatsExamplesActionsApproveActionACH] or
// [ReportStatsExamplesActionsReturnAction].
type ReportStatsExamplesActionsUnion interface {
	implementsReportStatsExamplesAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ReportStatsExamplesActionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReportStatsExamplesActionsDeclineActionAuthorization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReportStatsExamplesActionsChallengeActionAuthorization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReportStatsExamplesActionsResultAuthentication3DSAction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReportStatsExamplesActionsDeclineActionTokenization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReportStatsExamplesActionsRequireTfaAction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReportStatsExamplesActionsApproveActionACH{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReportStatsExamplesActionsReturnAction{}),
		},
	)
}

type ReportStatsExamplesActionsDeclineActionAuthorization struct {
	// The detailed result code explaining the specific reason for the decline
	Code ReportStatsExamplesActionsDeclineActionAuthorizationCode `json:"code" api:"required"`
	Type ReportStatsExamplesActionsDeclineActionAuthorizationType `json:"type" api:"required"`
	JSON reportStatsExamplesActionsDeclineActionAuthorizationJSON `json:"-"`
}

// reportStatsExamplesActionsDeclineActionAuthorizationJSON contains the JSON
// metadata for the struct [ReportStatsExamplesActionsDeclineActionAuthorization]
type reportStatsExamplesActionsDeclineActionAuthorizationJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportStatsExamplesActionsDeclineActionAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsExamplesActionsDeclineActionAuthorizationJSON) RawJSON() string {
	return r.raw
}

func (r ReportStatsExamplesActionsDeclineActionAuthorization) implementsReportStatsExamplesAction() {}

// The detailed result code explaining the specific reason for the decline
type ReportStatsExamplesActionsDeclineActionAuthorizationCode string

const (
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountDailySpendLimitExceeded              ReportStatsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_DAILY_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountDelinquent                           ReportStatsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_DELINQUENT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountInactive                             ReportStatsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_INACTIVE"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountLifetimeSpendLimitExceeded           ReportStatsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_LIFETIME_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountMonthlySpendLimitExceeded            ReportStatsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_MONTHLY_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountPaused                               ReportStatsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_PAUSED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountUnderReview                          ReportStatsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_UNDER_REVIEW"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAddressIncorrect                            ReportStatsExamplesActionsDeclineActionAuthorizationCode = "ADDRESS_INCORRECT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeApproved                                    ReportStatsExamplesActionsDeclineActionAuthorizationCode = "APPROVED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRuleAllowedCountry                      ReportStatsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE_ALLOWED_COUNTRY"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRuleAllowedMcc                          ReportStatsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE_ALLOWED_MCC"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRuleBlockedCountry                      ReportStatsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE_BLOCKED_COUNTRY"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRuleBlockedMcc                          ReportStatsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE_BLOCKED_MCC"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRule                                    ReportStatsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardClosed                                  ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_CLOSED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardCryptogramValidationFailure             ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_CRYPTOGRAM_VALIDATION_FAILURE"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardExpired                                 ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_EXPIRED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardExpiryDateIncorrect                     ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_EXPIRY_DATE_INCORRECT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardInvalid                                 ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_INVALID"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardNotActivated                            ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_NOT_ACTIVATED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardPaused                                  ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_PAUSED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardPinIncorrect                            ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_PIN_INCORRECT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardRestricted                              ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_RESTRICTED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardSecurityCodeIncorrect                   ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_SECURITY_CODE_INCORRECT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardSpendLimitExceeded                      ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARD_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeContactCardIssuer                           ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CONTACT_CARD_ISSUER"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCustomerAsaTimeout                          ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CUSTOMER_ASA_TIMEOUT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCustomAsaResult                             ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CUSTOM_ASA_RESULT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeDeclined                                    ReportStatsExamplesActionsDeclineActionAuthorizationCode = "DECLINED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeDoNotHonor                                  ReportStatsExamplesActionsDeclineActionAuthorizationCode = "DO_NOT_HONOR"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeDriverNumberInvalid                         ReportStatsExamplesActionsDeclineActionAuthorizationCode = "DRIVER_NUMBER_INVALID"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeFormatError                                 ReportStatsExamplesActionsDeclineActionAuthorizationCode = "FORMAT_ERROR"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeInsufficientFundingSourceBalance            ReportStatsExamplesActionsDeclineActionAuthorizationCode = "INSUFFICIENT_FUNDING_SOURCE_BALANCE"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeInsufficientFunds                           ReportStatsExamplesActionsDeclineActionAuthorizationCode = "INSUFFICIENT_FUNDS"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeLithicSystemError                           ReportStatsExamplesActionsDeclineActionAuthorizationCode = "LITHIC_SYSTEM_ERROR"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeLithicSystemRateLimit                       ReportStatsExamplesActionsDeclineActionAuthorizationCode = "LITHIC_SYSTEM_RATE_LIMIT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeMalformedAsaResponse                        ReportStatsExamplesActionsDeclineActionAuthorizationCode = "MALFORMED_ASA_RESPONSE"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeMerchantInvalid                             ReportStatsExamplesActionsDeclineActionAuthorizationCode = "MERCHANT_INVALID"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeMerchantLockedCardAttemptedElsewhere        ReportStatsExamplesActionsDeclineActionAuthorizationCode = "MERCHANT_LOCKED_CARD_ATTEMPTED_ELSEWHERE"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeMerchantNotPermitted                        ReportStatsExamplesActionsDeclineActionAuthorizationCode = "MERCHANT_NOT_PERMITTED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeOverReversalAttempted                       ReportStatsExamplesActionsDeclineActionAuthorizationCode = "OVER_REVERSAL_ATTEMPTED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodePinBlocked                                  ReportStatsExamplesActionsDeclineActionAuthorizationCode = "PIN_BLOCKED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeProgramCardSpendLimitExceeded               ReportStatsExamplesActionsDeclineActionAuthorizationCode = "PROGRAM_CARD_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeProgramSuspended                            ReportStatsExamplesActionsDeclineActionAuthorizationCode = "PROGRAM_SUSPENDED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeProgramUsageRestriction                     ReportStatsExamplesActionsDeclineActionAuthorizationCode = "PROGRAM_USAGE_RESTRICTION"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeReversalUnmatched                           ReportStatsExamplesActionsDeclineActionAuthorizationCode = "REVERSAL_UNMATCHED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeSecurityViolation                           ReportStatsExamplesActionsDeclineActionAuthorizationCode = "SECURITY_VIOLATION"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeSingleUseCardReattempted                    ReportStatsExamplesActionsDeclineActionAuthorizationCode = "SINGLE_USE_CARD_REATTEMPTED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeSuspectedFraud                              ReportStatsExamplesActionsDeclineActionAuthorizationCode = "SUSPECTED_FRAUD"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeTransactionInvalid                          ReportStatsExamplesActionsDeclineActionAuthorizationCode = "TRANSACTION_INVALID"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeTransactionNotPermittedToAcquirerOrTerminal ReportStatsExamplesActionsDeclineActionAuthorizationCode = "TRANSACTION_NOT_PERMITTED_TO_ACQUIRER_OR_TERMINAL"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeTransactionNotPermittedToIssuerOrCardholder ReportStatsExamplesActionsDeclineActionAuthorizationCode = "TRANSACTION_NOT_PERMITTED_TO_ISSUER_OR_CARDHOLDER"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeTransactionPreviouslyCompleted              ReportStatsExamplesActionsDeclineActionAuthorizationCode = "TRANSACTION_PREVIOUSLY_COMPLETED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeUnauthorizedMerchant                        ReportStatsExamplesActionsDeclineActionAuthorizationCode = "UNAUTHORIZED_MERCHANT"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeVehicleNumberInvalid                        ReportStatsExamplesActionsDeclineActionAuthorizationCode = "VEHICLE_NUMBER_INVALID"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardholderChallenged                        ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARDHOLDER_CHALLENGED"
	ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardholderChallengeFailed                   ReportStatsExamplesActionsDeclineActionAuthorizationCode = "CARDHOLDER_CHALLENGE_FAILED"
)

func (r ReportStatsExamplesActionsDeclineActionAuthorizationCode) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountDailySpendLimitExceeded, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountDelinquent, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountInactive, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountLifetimeSpendLimitExceeded, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountMonthlySpendLimitExceeded, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountPaused, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAccountUnderReview, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAddressIncorrect, ReportStatsExamplesActionsDeclineActionAuthorizationCodeApproved, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRuleAllowedCountry, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRuleAllowedMcc, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRuleBlockedCountry, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRuleBlockedMcc, ReportStatsExamplesActionsDeclineActionAuthorizationCodeAuthRule, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardClosed, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardCryptogramValidationFailure, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardExpired, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardExpiryDateIncorrect, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardInvalid, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardNotActivated, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardPaused, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardPinIncorrect, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardRestricted, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardSecurityCodeIncorrect, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardSpendLimitExceeded, ReportStatsExamplesActionsDeclineActionAuthorizationCodeContactCardIssuer, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCustomerAsaTimeout, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCustomAsaResult, ReportStatsExamplesActionsDeclineActionAuthorizationCodeDeclined, ReportStatsExamplesActionsDeclineActionAuthorizationCodeDoNotHonor, ReportStatsExamplesActionsDeclineActionAuthorizationCodeDriverNumberInvalid, ReportStatsExamplesActionsDeclineActionAuthorizationCodeFormatError, ReportStatsExamplesActionsDeclineActionAuthorizationCodeInsufficientFundingSourceBalance, ReportStatsExamplesActionsDeclineActionAuthorizationCodeInsufficientFunds, ReportStatsExamplesActionsDeclineActionAuthorizationCodeLithicSystemError, ReportStatsExamplesActionsDeclineActionAuthorizationCodeLithicSystemRateLimit, ReportStatsExamplesActionsDeclineActionAuthorizationCodeMalformedAsaResponse, ReportStatsExamplesActionsDeclineActionAuthorizationCodeMerchantInvalid, ReportStatsExamplesActionsDeclineActionAuthorizationCodeMerchantLockedCardAttemptedElsewhere, ReportStatsExamplesActionsDeclineActionAuthorizationCodeMerchantNotPermitted, ReportStatsExamplesActionsDeclineActionAuthorizationCodeOverReversalAttempted, ReportStatsExamplesActionsDeclineActionAuthorizationCodePinBlocked, ReportStatsExamplesActionsDeclineActionAuthorizationCodeProgramCardSpendLimitExceeded, ReportStatsExamplesActionsDeclineActionAuthorizationCodeProgramSuspended, ReportStatsExamplesActionsDeclineActionAuthorizationCodeProgramUsageRestriction, ReportStatsExamplesActionsDeclineActionAuthorizationCodeReversalUnmatched, ReportStatsExamplesActionsDeclineActionAuthorizationCodeSecurityViolation, ReportStatsExamplesActionsDeclineActionAuthorizationCodeSingleUseCardReattempted, ReportStatsExamplesActionsDeclineActionAuthorizationCodeSuspectedFraud, ReportStatsExamplesActionsDeclineActionAuthorizationCodeTransactionInvalid, ReportStatsExamplesActionsDeclineActionAuthorizationCodeTransactionNotPermittedToAcquirerOrTerminal, ReportStatsExamplesActionsDeclineActionAuthorizationCodeTransactionNotPermittedToIssuerOrCardholder, ReportStatsExamplesActionsDeclineActionAuthorizationCodeTransactionPreviouslyCompleted, ReportStatsExamplesActionsDeclineActionAuthorizationCodeUnauthorizedMerchant, ReportStatsExamplesActionsDeclineActionAuthorizationCodeVehicleNumberInvalid, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardholderChallenged, ReportStatsExamplesActionsDeclineActionAuthorizationCodeCardholderChallengeFailed:
		return true
	}
	return false
}

type ReportStatsExamplesActionsDeclineActionAuthorizationType string

const (
	ReportStatsExamplesActionsDeclineActionAuthorizationTypeDecline ReportStatsExamplesActionsDeclineActionAuthorizationType = "DECLINE"
)

func (r ReportStatsExamplesActionsDeclineActionAuthorizationType) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsDeclineActionAuthorizationTypeDecline:
		return true
	}
	return false
}

type ReportStatsExamplesActionsChallengeActionAuthorization struct {
	Type ReportStatsExamplesActionsChallengeActionAuthorizationType `json:"type" api:"required"`
	JSON reportStatsExamplesActionsChallengeActionAuthorizationJSON `json:"-"`
}

// reportStatsExamplesActionsChallengeActionAuthorizationJSON contains the JSON
// metadata for the struct [ReportStatsExamplesActionsChallengeActionAuthorization]
type reportStatsExamplesActionsChallengeActionAuthorizationJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportStatsExamplesActionsChallengeActionAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsExamplesActionsChallengeActionAuthorizationJSON) RawJSON() string {
	return r.raw
}

func (r ReportStatsExamplesActionsChallengeActionAuthorization) implementsReportStatsExamplesAction() {
}

type ReportStatsExamplesActionsChallengeActionAuthorizationType string

const (
	ReportStatsExamplesActionsChallengeActionAuthorizationTypeChallenge ReportStatsExamplesActionsChallengeActionAuthorizationType = "CHALLENGE"
)

func (r ReportStatsExamplesActionsChallengeActionAuthorizationType) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsChallengeActionAuthorizationTypeChallenge:
		return true
	}
	return false
}

type ReportStatsExamplesActionsResultAuthentication3DSAction struct {
	Type ReportStatsExamplesActionsResultAuthentication3DSActionType `json:"type" api:"required"`
	JSON reportStatsExamplesActionsResultAuthentication3DsActionJSON `json:"-"`
}

// reportStatsExamplesActionsResultAuthentication3DsActionJSON contains the JSON
// metadata for the struct
// [ReportStatsExamplesActionsResultAuthentication3DSAction]
type reportStatsExamplesActionsResultAuthentication3DsActionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportStatsExamplesActionsResultAuthentication3DSAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsExamplesActionsResultAuthentication3DsActionJSON) RawJSON() string {
	return r.raw
}

func (r ReportStatsExamplesActionsResultAuthentication3DSAction) implementsReportStatsExamplesAction() {
}

type ReportStatsExamplesActionsResultAuthentication3DSActionType string

const (
	ReportStatsExamplesActionsResultAuthentication3DSActionTypeDecline   ReportStatsExamplesActionsResultAuthentication3DSActionType = "DECLINE"
	ReportStatsExamplesActionsResultAuthentication3DSActionTypeChallenge ReportStatsExamplesActionsResultAuthentication3DSActionType = "CHALLENGE"
)

func (r ReportStatsExamplesActionsResultAuthentication3DSActionType) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsResultAuthentication3DSActionTypeDecline, ReportStatsExamplesActionsResultAuthentication3DSActionTypeChallenge:
		return true
	}
	return false
}

type ReportStatsExamplesActionsDeclineActionTokenization struct {
	// Decline the tokenization request
	Type ReportStatsExamplesActionsDeclineActionTokenizationType `json:"type" api:"required"`
	// Reason code for declining the tokenization request
	Reason ReportStatsExamplesActionsDeclineActionTokenizationReason `json:"reason"`
	JSON   reportStatsExamplesActionsDeclineActionTokenizationJSON   `json:"-"`
}

// reportStatsExamplesActionsDeclineActionTokenizationJSON contains the JSON
// metadata for the struct [ReportStatsExamplesActionsDeclineActionTokenization]
type reportStatsExamplesActionsDeclineActionTokenizationJSON struct {
	Type        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportStatsExamplesActionsDeclineActionTokenization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsExamplesActionsDeclineActionTokenizationJSON) RawJSON() string {
	return r.raw
}

func (r ReportStatsExamplesActionsDeclineActionTokenization) implementsReportStatsExamplesAction() {}

// Decline the tokenization request
type ReportStatsExamplesActionsDeclineActionTokenizationType string

const (
	ReportStatsExamplesActionsDeclineActionTokenizationTypeDecline ReportStatsExamplesActionsDeclineActionTokenizationType = "DECLINE"
)

func (r ReportStatsExamplesActionsDeclineActionTokenizationType) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsDeclineActionTokenizationTypeDecline:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type ReportStatsExamplesActionsDeclineActionTokenizationReason string

const (
	ReportStatsExamplesActionsDeclineActionTokenizationReasonAccountScore1                  ReportStatsExamplesActionsDeclineActionTokenizationReason = "ACCOUNT_SCORE_1"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonDeviceScore1                   ReportStatsExamplesActionsDeclineActionTokenizationReason = "DEVICE_SCORE_1"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonAllWalletDeclineReasonsPresent ReportStatsExamplesActionsDeclineActionTokenizationReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonWalletRecommendedDecisionRed   ReportStatsExamplesActionsDeclineActionTokenizationReason = "WALLET_RECOMMENDED_DECISION_RED"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonCvcMismatch                    ReportStatsExamplesActionsDeclineActionTokenizationReason = "CVC_MISMATCH"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonCardExpiryMonthMismatch        ReportStatsExamplesActionsDeclineActionTokenizationReason = "CARD_EXPIRY_MONTH_MISMATCH"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonCardExpiryYearMismatch         ReportStatsExamplesActionsDeclineActionTokenizationReason = "CARD_EXPIRY_YEAR_MISMATCH"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonCardInvalidState               ReportStatsExamplesActionsDeclineActionTokenizationReason = "CARD_INVALID_STATE"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonCustomerRedPath                ReportStatsExamplesActionsDeclineActionTokenizationReason = "CUSTOMER_RED_PATH"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonInvalidCustomerResponse        ReportStatsExamplesActionsDeclineActionTokenizationReason = "INVALID_CUSTOMER_RESPONSE"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonNetworkFailure                 ReportStatsExamplesActionsDeclineActionTokenizationReason = "NETWORK_FAILURE"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonGenericDecline                 ReportStatsExamplesActionsDeclineActionTokenizationReason = "GENERIC_DECLINE"
	ReportStatsExamplesActionsDeclineActionTokenizationReasonDigitalCardArtRequired         ReportStatsExamplesActionsDeclineActionTokenizationReason = "DIGITAL_CARD_ART_REQUIRED"
)

func (r ReportStatsExamplesActionsDeclineActionTokenizationReason) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsDeclineActionTokenizationReasonAccountScore1, ReportStatsExamplesActionsDeclineActionTokenizationReasonDeviceScore1, ReportStatsExamplesActionsDeclineActionTokenizationReasonAllWalletDeclineReasonsPresent, ReportStatsExamplesActionsDeclineActionTokenizationReasonWalletRecommendedDecisionRed, ReportStatsExamplesActionsDeclineActionTokenizationReasonCvcMismatch, ReportStatsExamplesActionsDeclineActionTokenizationReasonCardExpiryMonthMismatch, ReportStatsExamplesActionsDeclineActionTokenizationReasonCardExpiryYearMismatch, ReportStatsExamplesActionsDeclineActionTokenizationReasonCardInvalidState, ReportStatsExamplesActionsDeclineActionTokenizationReasonCustomerRedPath, ReportStatsExamplesActionsDeclineActionTokenizationReasonInvalidCustomerResponse, ReportStatsExamplesActionsDeclineActionTokenizationReasonNetworkFailure, ReportStatsExamplesActionsDeclineActionTokenizationReasonGenericDecline, ReportStatsExamplesActionsDeclineActionTokenizationReasonDigitalCardArtRequired:
		return true
	}
	return false
}

type ReportStatsExamplesActionsRequireTfaAction struct {
	// Require two-factor authentication for the tokenization request
	Type ReportStatsExamplesActionsRequireTfaActionType `json:"type" api:"required"`
	// Reason code for requiring two-factor authentication
	Reason ReportStatsExamplesActionsRequireTfaActionReason `json:"reason"`
	JSON   reportStatsExamplesActionsRequireTfaActionJSON   `json:"-"`
}

// reportStatsExamplesActionsRequireTfaActionJSON contains the JSON metadata for
// the struct [ReportStatsExamplesActionsRequireTfaAction]
type reportStatsExamplesActionsRequireTfaActionJSON struct {
	Type        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportStatsExamplesActionsRequireTfaAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsExamplesActionsRequireTfaActionJSON) RawJSON() string {
	return r.raw
}

func (r ReportStatsExamplesActionsRequireTfaAction) implementsReportStatsExamplesAction() {}

// Require two-factor authentication for the tokenization request
type ReportStatsExamplesActionsRequireTfaActionType string

const (
	ReportStatsExamplesActionsRequireTfaActionTypeRequireTfa ReportStatsExamplesActionsRequireTfaActionType = "REQUIRE_TFA"
)

func (r ReportStatsExamplesActionsRequireTfaActionType) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsRequireTfaActionTypeRequireTfa:
		return true
	}
	return false
}

// Reason code for requiring two-factor authentication
type ReportStatsExamplesActionsRequireTfaActionReason string

const (
	ReportStatsExamplesActionsRequireTfaActionReasonWalletRecommendedTfa        ReportStatsExamplesActionsRequireTfaActionReason = "WALLET_RECOMMENDED_TFA"
	ReportStatsExamplesActionsRequireTfaActionReasonSuspiciousActivity          ReportStatsExamplesActionsRequireTfaActionReason = "SUSPICIOUS_ACTIVITY"
	ReportStatsExamplesActionsRequireTfaActionReasonDeviceRecentlyLost          ReportStatsExamplesActionsRequireTfaActionReason = "DEVICE_RECENTLY_LOST"
	ReportStatsExamplesActionsRequireTfaActionReasonTooManyRecentAttempts       ReportStatsExamplesActionsRequireTfaActionReason = "TOO_MANY_RECENT_ATTEMPTS"
	ReportStatsExamplesActionsRequireTfaActionReasonTooManyRecentTokens         ReportStatsExamplesActionsRequireTfaActionReason = "TOO_MANY_RECENT_TOKENS"
	ReportStatsExamplesActionsRequireTfaActionReasonTooManyDifferentCardholders ReportStatsExamplesActionsRequireTfaActionReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	ReportStatsExamplesActionsRequireTfaActionReasonOutsideHomeTerritory        ReportStatsExamplesActionsRequireTfaActionReason = "OUTSIDE_HOME_TERRITORY"
	ReportStatsExamplesActionsRequireTfaActionReasonHasSuspendedTokens          ReportStatsExamplesActionsRequireTfaActionReason = "HAS_SUSPENDED_TOKENS"
	ReportStatsExamplesActionsRequireTfaActionReasonHighRisk                    ReportStatsExamplesActionsRequireTfaActionReason = "HIGH_RISK"
	ReportStatsExamplesActionsRequireTfaActionReasonAccountScoreLow             ReportStatsExamplesActionsRequireTfaActionReason = "ACCOUNT_SCORE_LOW"
	ReportStatsExamplesActionsRequireTfaActionReasonDeviceScoreLow              ReportStatsExamplesActionsRequireTfaActionReason = "DEVICE_SCORE_LOW"
	ReportStatsExamplesActionsRequireTfaActionReasonCardStateTfa                ReportStatsExamplesActionsRequireTfaActionReason = "CARD_STATE_TFA"
	ReportStatsExamplesActionsRequireTfaActionReasonHardcodedTfa                ReportStatsExamplesActionsRequireTfaActionReason = "HARDCODED_TFA"
	ReportStatsExamplesActionsRequireTfaActionReasonCustomerRuleTfa             ReportStatsExamplesActionsRequireTfaActionReason = "CUSTOMER_RULE_TFA"
	ReportStatsExamplesActionsRequireTfaActionReasonDeviceHostCardEmulation     ReportStatsExamplesActionsRequireTfaActionReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r ReportStatsExamplesActionsRequireTfaActionReason) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsRequireTfaActionReasonWalletRecommendedTfa, ReportStatsExamplesActionsRequireTfaActionReasonSuspiciousActivity, ReportStatsExamplesActionsRequireTfaActionReasonDeviceRecentlyLost, ReportStatsExamplesActionsRequireTfaActionReasonTooManyRecentAttempts, ReportStatsExamplesActionsRequireTfaActionReasonTooManyRecentTokens, ReportStatsExamplesActionsRequireTfaActionReasonTooManyDifferentCardholders, ReportStatsExamplesActionsRequireTfaActionReasonOutsideHomeTerritory, ReportStatsExamplesActionsRequireTfaActionReasonHasSuspendedTokens, ReportStatsExamplesActionsRequireTfaActionReasonHighRisk, ReportStatsExamplesActionsRequireTfaActionReasonAccountScoreLow, ReportStatsExamplesActionsRequireTfaActionReasonDeviceScoreLow, ReportStatsExamplesActionsRequireTfaActionReasonCardStateTfa, ReportStatsExamplesActionsRequireTfaActionReasonHardcodedTfa, ReportStatsExamplesActionsRequireTfaActionReasonCustomerRuleTfa, ReportStatsExamplesActionsRequireTfaActionReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

type ReportStatsExamplesActionsApproveActionACH struct {
	// Approve the ACH transaction
	Type ReportStatsExamplesActionsApproveActionACHType `json:"type" api:"required"`
	JSON reportStatsExamplesActionsApproveActionACHJSON `json:"-"`
}

// reportStatsExamplesActionsApproveActionACHJSON contains the JSON metadata for
// the struct [ReportStatsExamplesActionsApproveActionACH]
type reportStatsExamplesActionsApproveActionACHJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportStatsExamplesActionsApproveActionACH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsExamplesActionsApproveActionACHJSON) RawJSON() string {
	return r.raw
}

func (r ReportStatsExamplesActionsApproveActionACH) implementsReportStatsExamplesAction() {}

// Approve the ACH transaction
type ReportStatsExamplesActionsApproveActionACHType string

const (
	ReportStatsExamplesActionsApproveActionACHTypeApprove ReportStatsExamplesActionsApproveActionACHType = "APPROVE"
)

func (r ReportStatsExamplesActionsApproveActionACHType) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsApproveActionACHTypeApprove:
		return true
	}
	return false
}

type ReportStatsExamplesActionsReturnAction struct {
	// NACHA return code to use when returning the transaction. Note that the list of
	// available return codes is subject to an allowlist configured at the program
	// level
	Code ReportStatsExamplesActionsReturnActionCode `json:"code" api:"required"`
	// Return the ACH transaction
	Type ReportStatsExamplesActionsReturnActionType `json:"type" api:"required"`
	JSON reportStatsExamplesActionsReturnActionJSON `json:"-"`
}

// reportStatsExamplesActionsReturnActionJSON contains the JSON metadata for the
// struct [ReportStatsExamplesActionsReturnAction]
type reportStatsExamplesActionsReturnActionJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportStatsExamplesActionsReturnAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportStatsExamplesActionsReturnActionJSON) RawJSON() string {
	return r.raw
}

func (r ReportStatsExamplesActionsReturnAction) implementsReportStatsExamplesAction() {}

// NACHA return code to use when returning the transaction. Note that the list of
// available return codes is subject to an allowlist configured at the program
// level
type ReportStatsExamplesActionsReturnActionCode string

const (
	ReportStatsExamplesActionsReturnActionCodeR01 ReportStatsExamplesActionsReturnActionCode = "R01"
	ReportStatsExamplesActionsReturnActionCodeR02 ReportStatsExamplesActionsReturnActionCode = "R02"
	ReportStatsExamplesActionsReturnActionCodeR03 ReportStatsExamplesActionsReturnActionCode = "R03"
	ReportStatsExamplesActionsReturnActionCodeR04 ReportStatsExamplesActionsReturnActionCode = "R04"
	ReportStatsExamplesActionsReturnActionCodeR05 ReportStatsExamplesActionsReturnActionCode = "R05"
	ReportStatsExamplesActionsReturnActionCodeR06 ReportStatsExamplesActionsReturnActionCode = "R06"
	ReportStatsExamplesActionsReturnActionCodeR07 ReportStatsExamplesActionsReturnActionCode = "R07"
	ReportStatsExamplesActionsReturnActionCodeR08 ReportStatsExamplesActionsReturnActionCode = "R08"
	ReportStatsExamplesActionsReturnActionCodeR09 ReportStatsExamplesActionsReturnActionCode = "R09"
	ReportStatsExamplesActionsReturnActionCodeR10 ReportStatsExamplesActionsReturnActionCode = "R10"
	ReportStatsExamplesActionsReturnActionCodeR11 ReportStatsExamplesActionsReturnActionCode = "R11"
	ReportStatsExamplesActionsReturnActionCodeR12 ReportStatsExamplesActionsReturnActionCode = "R12"
	ReportStatsExamplesActionsReturnActionCodeR13 ReportStatsExamplesActionsReturnActionCode = "R13"
	ReportStatsExamplesActionsReturnActionCodeR14 ReportStatsExamplesActionsReturnActionCode = "R14"
	ReportStatsExamplesActionsReturnActionCodeR15 ReportStatsExamplesActionsReturnActionCode = "R15"
	ReportStatsExamplesActionsReturnActionCodeR16 ReportStatsExamplesActionsReturnActionCode = "R16"
	ReportStatsExamplesActionsReturnActionCodeR17 ReportStatsExamplesActionsReturnActionCode = "R17"
	ReportStatsExamplesActionsReturnActionCodeR18 ReportStatsExamplesActionsReturnActionCode = "R18"
	ReportStatsExamplesActionsReturnActionCodeR19 ReportStatsExamplesActionsReturnActionCode = "R19"
	ReportStatsExamplesActionsReturnActionCodeR20 ReportStatsExamplesActionsReturnActionCode = "R20"
	ReportStatsExamplesActionsReturnActionCodeR21 ReportStatsExamplesActionsReturnActionCode = "R21"
	ReportStatsExamplesActionsReturnActionCodeR22 ReportStatsExamplesActionsReturnActionCode = "R22"
	ReportStatsExamplesActionsReturnActionCodeR23 ReportStatsExamplesActionsReturnActionCode = "R23"
	ReportStatsExamplesActionsReturnActionCodeR24 ReportStatsExamplesActionsReturnActionCode = "R24"
	ReportStatsExamplesActionsReturnActionCodeR25 ReportStatsExamplesActionsReturnActionCode = "R25"
	ReportStatsExamplesActionsReturnActionCodeR26 ReportStatsExamplesActionsReturnActionCode = "R26"
	ReportStatsExamplesActionsReturnActionCodeR27 ReportStatsExamplesActionsReturnActionCode = "R27"
	ReportStatsExamplesActionsReturnActionCodeR28 ReportStatsExamplesActionsReturnActionCode = "R28"
	ReportStatsExamplesActionsReturnActionCodeR29 ReportStatsExamplesActionsReturnActionCode = "R29"
	ReportStatsExamplesActionsReturnActionCodeR30 ReportStatsExamplesActionsReturnActionCode = "R30"
	ReportStatsExamplesActionsReturnActionCodeR31 ReportStatsExamplesActionsReturnActionCode = "R31"
	ReportStatsExamplesActionsReturnActionCodeR32 ReportStatsExamplesActionsReturnActionCode = "R32"
	ReportStatsExamplesActionsReturnActionCodeR33 ReportStatsExamplesActionsReturnActionCode = "R33"
	ReportStatsExamplesActionsReturnActionCodeR34 ReportStatsExamplesActionsReturnActionCode = "R34"
	ReportStatsExamplesActionsReturnActionCodeR35 ReportStatsExamplesActionsReturnActionCode = "R35"
	ReportStatsExamplesActionsReturnActionCodeR36 ReportStatsExamplesActionsReturnActionCode = "R36"
	ReportStatsExamplesActionsReturnActionCodeR37 ReportStatsExamplesActionsReturnActionCode = "R37"
	ReportStatsExamplesActionsReturnActionCodeR38 ReportStatsExamplesActionsReturnActionCode = "R38"
	ReportStatsExamplesActionsReturnActionCodeR39 ReportStatsExamplesActionsReturnActionCode = "R39"
	ReportStatsExamplesActionsReturnActionCodeR40 ReportStatsExamplesActionsReturnActionCode = "R40"
	ReportStatsExamplesActionsReturnActionCodeR41 ReportStatsExamplesActionsReturnActionCode = "R41"
	ReportStatsExamplesActionsReturnActionCodeR42 ReportStatsExamplesActionsReturnActionCode = "R42"
	ReportStatsExamplesActionsReturnActionCodeR43 ReportStatsExamplesActionsReturnActionCode = "R43"
	ReportStatsExamplesActionsReturnActionCodeR44 ReportStatsExamplesActionsReturnActionCode = "R44"
	ReportStatsExamplesActionsReturnActionCodeR45 ReportStatsExamplesActionsReturnActionCode = "R45"
	ReportStatsExamplesActionsReturnActionCodeR46 ReportStatsExamplesActionsReturnActionCode = "R46"
	ReportStatsExamplesActionsReturnActionCodeR47 ReportStatsExamplesActionsReturnActionCode = "R47"
	ReportStatsExamplesActionsReturnActionCodeR50 ReportStatsExamplesActionsReturnActionCode = "R50"
	ReportStatsExamplesActionsReturnActionCodeR51 ReportStatsExamplesActionsReturnActionCode = "R51"
	ReportStatsExamplesActionsReturnActionCodeR52 ReportStatsExamplesActionsReturnActionCode = "R52"
	ReportStatsExamplesActionsReturnActionCodeR53 ReportStatsExamplesActionsReturnActionCode = "R53"
	ReportStatsExamplesActionsReturnActionCodeR61 ReportStatsExamplesActionsReturnActionCode = "R61"
	ReportStatsExamplesActionsReturnActionCodeR62 ReportStatsExamplesActionsReturnActionCode = "R62"
	ReportStatsExamplesActionsReturnActionCodeR67 ReportStatsExamplesActionsReturnActionCode = "R67"
	ReportStatsExamplesActionsReturnActionCodeR68 ReportStatsExamplesActionsReturnActionCode = "R68"
	ReportStatsExamplesActionsReturnActionCodeR69 ReportStatsExamplesActionsReturnActionCode = "R69"
	ReportStatsExamplesActionsReturnActionCodeR70 ReportStatsExamplesActionsReturnActionCode = "R70"
	ReportStatsExamplesActionsReturnActionCodeR71 ReportStatsExamplesActionsReturnActionCode = "R71"
	ReportStatsExamplesActionsReturnActionCodeR72 ReportStatsExamplesActionsReturnActionCode = "R72"
	ReportStatsExamplesActionsReturnActionCodeR73 ReportStatsExamplesActionsReturnActionCode = "R73"
	ReportStatsExamplesActionsReturnActionCodeR74 ReportStatsExamplesActionsReturnActionCode = "R74"
	ReportStatsExamplesActionsReturnActionCodeR75 ReportStatsExamplesActionsReturnActionCode = "R75"
	ReportStatsExamplesActionsReturnActionCodeR76 ReportStatsExamplesActionsReturnActionCode = "R76"
	ReportStatsExamplesActionsReturnActionCodeR77 ReportStatsExamplesActionsReturnActionCode = "R77"
	ReportStatsExamplesActionsReturnActionCodeR80 ReportStatsExamplesActionsReturnActionCode = "R80"
	ReportStatsExamplesActionsReturnActionCodeR81 ReportStatsExamplesActionsReturnActionCode = "R81"
	ReportStatsExamplesActionsReturnActionCodeR82 ReportStatsExamplesActionsReturnActionCode = "R82"
	ReportStatsExamplesActionsReturnActionCodeR83 ReportStatsExamplesActionsReturnActionCode = "R83"
	ReportStatsExamplesActionsReturnActionCodeR84 ReportStatsExamplesActionsReturnActionCode = "R84"
	ReportStatsExamplesActionsReturnActionCodeR85 ReportStatsExamplesActionsReturnActionCode = "R85"
)

func (r ReportStatsExamplesActionsReturnActionCode) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsReturnActionCodeR01, ReportStatsExamplesActionsReturnActionCodeR02, ReportStatsExamplesActionsReturnActionCodeR03, ReportStatsExamplesActionsReturnActionCodeR04, ReportStatsExamplesActionsReturnActionCodeR05, ReportStatsExamplesActionsReturnActionCodeR06, ReportStatsExamplesActionsReturnActionCodeR07, ReportStatsExamplesActionsReturnActionCodeR08, ReportStatsExamplesActionsReturnActionCodeR09, ReportStatsExamplesActionsReturnActionCodeR10, ReportStatsExamplesActionsReturnActionCodeR11, ReportStatsExamplesActionsReturnActionCodeR12, ReportStatsExamplesActionsReturnActionCodeR13, ReportStatsExamplesActionsReturnActionCodeR14, ReportStatsExamplesActionsReturnActionCodeR15, ReportStatsExamplesActionsReturnActionCodeR16, ReportStatsExamplesActionsReturnActionCodeR17, ReportStatsExamplesActionsReturnActionCodeR18, ReportStatsExamplesActionsReturnActionCodeR19, ReportStatsExamplesActionsReturnActionCodeR20, ReportStatsExamplesActionsReturnActionCodeR21, ReportStatsExamplesActionsReturnActionCodeR22, ReportStatsExamplesActionsReturnActionCodeR23, ReportStatsExamplesActionsReturnActionCodeR24, ReportStatsExamplesActionsReturnActionCodeR25, ReportStatsExamplesActionsReturnActionCodeR26, ReportStatsExamplesActionsReturnActionCodeR27, ReportStatsExamplesActionsReturnActionCodeR28, ReportStatsExamplesActionsReturnActionCodeR29, ReportStatsExamplesActionsReturnActionCodeR30, ReportStatsExamplesActionsReturnActionCodeR31, ReportStatsExamplesActionsReturnActionCodeR32, ReportStatsExamplesActionsReturnActionCodeR33, ReportStatsExamplesActionsReturnActionCodeR34, ReportStatsExamplesActionsReturnActionCodeR35, ReportStatsExamplesActionsReturnActionCodeR36, ReportStatsExamplesActionsReturnActionCodeR37, ReportStatsExamplesActionsReturnActionCodeR38, ReportStatsExamplesActionsReturnActionCodeR39, ReportStatsExamplesActionsReturnActionCodeR40, ReportStatsExamplesActionsReturnActionCodeR41, ReportStatsExamplesActionsReturnActionCodeR42, ReportStatsExamplesActionsReturnActionCodeR43, ReportStatsExamplesActionsReturnActionCodeR44, ReportStatsExamplesActionsReturnActionCodeR45, ReportStatsExamplesActionsReturnActionCodeR46, ReportStatsExamplesActionsReturnActionCodeR47, ReportStatsExamplesActionsReturnActionCodeR50, ReportStatsExamplesActionsReturnActionCodeR51, ReportStatsExamplesActionsReturnActionCodeR52, ReportStatsExamplesActionsReturnActionCodeR53, ReportStatsExamplesActionsReturnActionCodeR61, ReportStatsExamplesActionsReturnActionCodeR62, ReportStatsExamplesActionsReturnActionCodeR67, ReportStatsExamplesActionsReturnActionCodeR68, ReportStatsExamplesActionsReturnActionCodeR69, ReportStatsExamplesActionsReturnActionCodeR70, ReportStatsExamplesActionsReturnActionCodeR71, ReportStatsExamplesActionsReturnActionCodeR72, ReportStatsExamplesActionsReturnActionCodeR73, ReportStatsExamplesActionsReturnActionCodeR74, ReportStatsExamplesActionsReturnActionCodeR75, ReportStatsExamplesActionsReturnActionCodeR76, ReportStatsExamplesActionsReturnActionCodeR77, ReportStatsExamplesActionsReturnActionCodeR80, ReportStatsExamplesActionsReturnActionCodeR81, ReportStatsExamplesActionsReturnActionCodeR82, ReportStatsExamplesActionsReturnActionCodeR83, ReportStatsExamplesActionsReturnActionCodeR84, ReportStatsExamplesActionsReturnActionCodeR85:
		return true
	}
	return false
}

// Return the ACH transaction
type ReportStatsExamplesActionsReturnActionType string

const (
	ReportStatsExamplesActionsReturnActionTypeReturn ReportStatsExamplesActionsReturnActionType = "RETURN"
)

func (r ReportStatsExamplesActionsReturnActionType) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsReturnActionTypeReturn:
		return true
	}
	return false
}

type ReportStatsExamplesActionsType string

const (
	ReportStatsExamplesActionsTypeDecline    ReportStatsExamplesActionsType = "DECLINE"
	ReportStatsExamplesActionsTypeChallenge  ReportStatsExamplesActionsType = "CHALLENGE"
	ReportStatsExamplesActionsTypeRequireTfa ReportStatsExamplesActionsType = "REQUIRE_TFA"
	ReportStatsExamplesActionsTypeApprove    ReportStatsExamplesActionsType = "APPROVE"
	ReportStatsExamplesActionsTypeReturn     ReportStatsExamplesActionsType = "RETURN"
)

func (r ReportStatsExamplesActionsType) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsTypeDecline, ReportStatsExamplesActionsTypeChallenge, ReportStatsExamplesActionsTypeRequireTfa, ReportStatsExamplesActionsTypeApprove, ReportStatsExamplesActionsTypeReturn:
		return true
	}
	return false
}

// The detailed result code explaining the specific reason for the decline
type ReportStatsExamplesActionsCode string

const (
	ReportStatsExamplesActionsCodeAccountDailySpendLimitExceeded              ReportStatsExamplesActionsCode = "ACCOUNT_DAILY_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsCodeAccountDelinquent                           ReportStatsExamplesActionsCode = "ACCOUNT_DELINQUENT"
	ReportStatsExamplesActionsCodeAccountInactive                             ReportStatsExamplesActionsCode = "ACCOUNT_INACTIVE"
	ReportStatsExamplesActionsCodeAccountLifetimeSpendLimitExceeded           ReportStatsExamplesActionsCode = "ACCOUNT_LIFETIME_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsCodeAccountMonthlySpendLimitExceeded            ReportStatsExamplesActionsCode = "ACCOUNT_MONTHLY_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsCodeAccountPaused                               ReportStatsExamplesActionsCode = "ACCOUNT_PAUSED"
	ReportStatsExamplesActionsCodeAccountUnderReview                          ReportStatsExamplesActionsCode = "ACCOUNT_UNDER_REVIEW"
	ReportStatsExamplesActionsCodeAddressIncorrect                            ReportStatsExamplesActionsCode = "ADDRESS_INCORRECT"
	ReportStatsExamplesActionsCodeApproved                                    ReportStatsExamplesActionsCode = "APPROVED"
	ReportStatsExamplesActionsCodeAuthRuleAllowedCountry                      ReportStatsExamplesActionsCode = "AUTH_RULE_ALLOWED_COUNTRY"
	ReportStatsExamplesActionsCodeAuthRuleAllowedMcc                          ReportStatsExamplesActionsCode = "AUTH_RULE_ALLOWED_MCC"
	ReportStatsExamplesActionsCodeAuthRuleBlockedCountry                      ReportStatsExamplesActionsCode = "AUTH_RULE_BLOCKED_COUNTRY"
	ReportStatsExamplesActionsCodeAuthRuleBlockedMcc                          ReportStatsExamplesActionsCode = "AUTH_RULE_BLOCKED_MCC"
	ReportStatsExamplesActionsCodeAuthRule                                    ReportStatsExamplesActionsCode = "AUTH_RULE"
	ReportStatsExamplesActionsCodeCardClosed                                  ReportStatsExamplesActionsCode = "CARD_CLOSED"
	ReportStatsExamplesActionsCodeCardCryptogramValidationFailure             ReportStatsExamplesActionsCode = "CARD_CRYPTOGRAM_VALIDATION_FAILURE"
	ReportStatsExamplesActionsCodeCardExpired                                 ReportStatsExamplesActionsCode = "CARD_EXPIRED"
	ReportStatsExamplesActionsCodeCardExpiryDateIncorrect                     ReportStatsExamplesActionsCode = "CARD_EXPIRY_DATE_INCORRECT"
	ReportStatsExamplesActionsCodeCardInvalid                                 ReportStatsExamplesActionsCode = "CARD_INVALID"
	ReportStatsExamplesActionsCodeCardNotActivated                            ReportStatsExamplesActionsCode = "CARD_NOT_ACTIVATED"
	ReportStatsExamplesActionsCodeCardPaused                                  ReportStatsExamplesActionsCode = "CARD_PAUSED"
	ReportStatsExamplesActionsCodeCardPinIncorrect                            ReportStatsExamplesActionsCode = "CARD_PIN_INCORRECT"
	ReportStatsExamplesActionsCodeCardRestricted                              ReportStatsExamplesActionsCode = "CARD_RESTRICTED"
	ReportStatsExamplesActionsCodeCardSecurityCodeIncorrect                   ReportStatsExamplesActionsCode = "CARD_SECURITY_CODE_INCORRECT"
	ReportStatsExamplesActionsCodeCardSpendLimitExceeded                      ReportStatsExamplesActionsCode = "CARD_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsCodeContactCardIssuer                           ReportStatsExamplesActionsCode = "CONTACT_CARD_ISSUER"
	ReportStatsExamplesActionsCodeCustomerAsaTimeout                          ReportStatsExamplesActionsCode = "CUSTOMER_ASA_TIMEOUT"
	ReportStatsExamplesActionsCodeCustomAsaResult                             ReportStatsExamplesActionsCode = "CUSTOM_ASA_RESULT"
	ReportStatsExamplesActionsCodeDeclined                                    ReportStatsExamplesActionsCode = "DECLINED"
	ReportStatsExamplesActionsCodeDoNotHonor                                  ReportStatsExamplesActionsCode = "DO_NOT_HONOR"
	ReportStatsExamplesActionsCodeDriverNumberInvalid                         ReportStatsExamplesActionsCode = "DRIVER_NUMBER_INVALID"
	ReportStatsExamplesActionsCodeFormatError                                 ReportStatsExamplesActionsCode = "FORMAT_ERROR"
	ReportStatsExamplesActionsCodeInsufficientFundingSourceBalance            ReportStatsExamplesActionsCode = "INSUFFICIENT_FUNDING_SOURCE_BALANCE"
	ReportStatsExamplesActionsCodeInsufficientFunds                           ReportStatsExamplesActionsCode = "INSUFFICIENT_FUNDS"
	ReportStatsExamplesActionsCodeLithicSystemError                           ReportStatsExamplesActionsCode = "LITHIC_SYSTEM_ERROR"
	ReportStatsExamplesActionsCodeLithicSystemRateLimit                       ReportStatsExamplesActionsCode = "LITHIC_SYSTEM_RATE_LIMIT"
	ReportStatsExamplesActionsCodeMalformedAsaResponse                        ReportStatsExamplesActionsCode = "MALFORMED_ASA_RESPONSE"
	ReportStatsExamplesActionsCodeMerchantInvalid                             ReportStatsExamplesActionsCode = "MERCHANT_INVALID"
	ReportStatsExamplesActionsCodeMerchantLockedCardAttemptedElsewhere        ReportStatsExamplesActionsCode = "MERCHANT_LOCKED_CARD_ATTEMPTED_ELSEWHERE"
	ReportStatsExamplesActionsCodeMerchantNotPermitted                        ReportStatsExamplesActionsCode = "MERCHANT_NOT_PERMITTED"
	ReportStatsExamplesActionsCodeOverReversalAttempted                       ReportStatsExamplesActionsCode = "OVER_REVERSAL_ATTEMPTED"
	ReportStatsExamplesActionsCodePinBlocked                                  ReportStatsExamplesActionsCode = "PIN_BLOCKED"
	ReportStatsExamplesActionsCodeProgramCardSpendLimitExceeded               ReportStatsExamplesActionsCode = "PROGRAM_CARD_SPEND_LIMIT_EXCEEDED"
	ReportStatsExamplesActionsCodeProgramSuspended                            ReportStatsExamplesActionsCode = "PROGRAM_SUSPENDED"
	ReportStatsExamplesActionsCodeProgramUsageRestriction                     ReportStatsExamplesActionsCode = "PROGRAM_USAGE_RESTRICTION"
	ReportStatsExamplesActionsCodeReversalUnmatched                           ReportStatsExamplesActionsCode = "REVERSAL_UNMATCHED"
	ReportStatsExamplesActionsCodeSecurityViolation                           ReportStatsExamplesActionsCode = "SECURITY_VIOLATION"
	ReportStatsExamplesActionsCodeSingleUseCardReattempted                    ReportStatsExamplesActionsCode = "SINGLE_USE_CARD_REATTEMPTED"
	ReportStatsExamplesActionsCodeSuspectedFraud                              ReportStatsExamplesActionsCode = "SUSPECTED_FRAUD"
	ReportStatsExamplesActionsCodeTransactionInvalid                          ReportStatsExamplesActionsCode = "TRANSACTION_INVALID"
	ReportStatsExamplesActionsCodeTransactionNotPermittedToAcquirerOrTerminal ReportStatsExamplesActionsCode = "TRANSACTION_NOT_PERMITTED_TO_ACQUIRER_OR_TERMINAL"
	ReportStatsExamplesActionsCodeTransactionNotPermittedToIssuerOrCardholder ReportStatsExamplesActionsCode = "TRANSACTION_NOT_PERMITTED_TO_ISSUER_OR_CARDHOLDER"
	ReportStatsExamplesActionsCodeTransactionPreviouslyCompleted              ReportStatsExamplesActionsCode = "TRANSACTION_PREVIOUSLY_COMPLETED"
	ReportStatsExamplesActionsCodeUnauthorizedMerchant                        ReportStatsExamplesActionsCode = "UNAUTHORIZED_MERCHANT"
	ReportStatsExamplesActionsCodeVehicleNumberInvalid                        ReportStatsExamplesActionsCode = "VEHICLE_NUMBER_INVALID"
	ReportStatsExamplesActionsCodeCardholderChallenged                        ReportStatsExamplesActionsCode = "CARDHOLDER_CHALLENGED"
	ReportStatsExamplesActionsCodeCardholderChallengeFailed                   ReportStatsExamplesActionsCode = "CARDHOLDER_CHALLENGE_FAILED"
	ReportStatsExamplesActionsCodeR01                                         ReportStatsExamplesActionsCode = "R01"
	ReportStatsExamplesActionsCodeR02                                         ReportStatsExamplesActionsCode = "R02"
	ReportStatsExamplesActionsCodeR03                                         ReportStatsExamplesActionsCode = "R03"
	ReportStatsExamplesActionsCodeR04                                         ReportStatsExamplesActionsCode = "R04"
	ReportStatsExamplesActionsCodeR05                                         ReportStatsExamplesActionsCode = "R05"
	ReportStatsExamplesActionsCodeR06                                         ReportStatsExamplesActionsCode = "R06"
	ReportStatsExamplesActionsCodeR07                                         ReportStatsExamplesActionsCode = "R07"
	ReportStatsExamplesActionsCodeR08                                         ReportStatsExamplesActionsCode = "R08"
	ReportStatsExamplesActionsCodeR09                                         ReportStatsExamplesActionsCode = "R09"
	ReportStatsExamplesActionsCodeR10                                         ReportStatsExamplesActionsCode = "R10"
	ReportStatsExamplesActionsCodeR11                                         ReportStatsExamplesActionsCode = "R11"
	ReportStatsExamplesActionsCodeR12                                         ReportStatsExamplesActionsCode = "R12"
	ReportStatsExamplesActionsCodeR13                                         ReportStatsExamplesActionsCode = "R13"
	ReportStatsExamplesActionsCodeR14                                         ReportStatsExamplesActionsCode = "R14"
	ReportStatsExamplesActionsCodeR15                                         ReportStatsExamplesActionsCode = "R15"
	ReportStatsExamplesActionsCodeR16                                         ReportStatsExamplesActionsCode = "R16"
	ReportStatsExamplesActionsCodeR17                                         ReportStatsExamplesActionsCode = "R17"
	ReportStatsExamplesActionsCodeR18                                         ReportStatsExamplesActionsCode = "R18"
	ReportStatsExamplesActionsCodeR19                                         ReportStatsExamplesActionsCode = "R19"
	ReportStatsExamplesActionsCodeR20                                         ReportStatsExamplesActionsCode = "R20"
	ReportStatsExamplesActionsCodeR21                                         ReportStatsExamplesActionsCode = "R21"
	ReportStatsExamplesActionsCodeR22                                         ReportStatsExamplesActionsCode = "R22"
	ReportStatsExamplesActionsCodeR23                                         ReportStatsExamplesActionsCode = "R23"
	ReportStatsExamplesActionsCodeR24                                         ReportStatsExamplesActionsCode = "R24"
	ReportStatsExamplesActionsCodeR25                                         ReportStatsExamplesActionsCode = "R25"
	ReportStatsExamplesActionsCodeR26                                         ReportStatsExamplesActionsCode = "R26"
	ReportStatsExamplesActionsCodeR27                                         ReportStatsExamplesActionsCode = "R27"
	ReportStatsExamplesActionsCodeR28                                         ReportStatsExamplesActionsCode = "R28"
	ReportStatsExamplesActionsCodeR29                                         ReportStatsExamplesActionsCode = "R29"
	ReportStatsExamplesActionsCodeR30                                         ReportStatsExamplesActionsCode = "R30"
	ReportStatsExamplesActionsCodeR31                                         ReportStatsExamplesActionsCode = "R31"
	ReportStatsExamplesActionsCodeR32                                         ReportStatsExamplesActionsCode = "R32"
	ReportStatsExamplesActionsCodeR33                                         ReportStatsExamplesActionsCode = "R33"
	ReportStatsExamplesActionsCodeR34                                         ReportStatsExamplesActionsCode = "R34"
	ReportStatsExamplesActionsCodeR35                                         ReportStatsExamplesActionsCode = "R35"
	ReportStatsExamplesActionsCodeR36                                         ReportStatsExamplesActionsCode = "R36"
	ReportStatsExamplesActionsCodeR37                                         ReportStatsExamplesActionsCode = "R37"
	ReportStatsExamplesActionsCodeR38                                         ReportStatsExamplesActionsCode = "R38"
	ReportStatsExamplesActionsCodeR39                                         ReportStatsExamplesActionsCode = "R39"
	ReportStatsExamplesActionsCodeR40                                         ReportStatsExamplesActionsCode = "R40"
	ReportStatsExamplesActionsCodeR41                                         ReportStatsExamplesActionsCode = "R41"
	ReportStatsExamplesActionsCodeR42                                         ReportStatsExamplesActionsCode = "R42"
	ReportStatsExamplesActionsCodeR43                                         ReportStatsExamplesActionsCode = "R43"
	ReportStatsExamplesActionsCodeR44                                         ReportStatsExamplesActionsCode = "R44"
	ReportStatsExamplesActionsCodeR45                                         ReportStatsExamplesActionsCode = "R45"
	ReportStatsExamplesActionsCodeR46                                         ReportStatsExamplesActionsCode = "R46"
	ReportStatsExamplesActionsCodeR47                                         ReportStatsExamplesActionsCode = "R47"
	ReportStatsExamplesActionsCodeR50                                         ReportStatsExamplesActionsCode = "R50"
	ReportStatsExamplesActionsCodeR51                                         ReportStatsExamplesActionsCode = "R51"
	ReportStatsExamplesActionsCodeR52                                         ReportStatsExamplesActionsCode = "R52"
	ReportStatsExamplesActionsCodeR53                                         ReportStatsExamplesActionsCode = "R53"
	ReportStatsExamplesActionsCodeR61                                         ReportStatsExamplesActionsCode = "R61"
	ReportStatsExamplesActionsCodeR62                                         ReportStatsExamplesActionsCode = "R62"
	ReportStatsExamplesActionsCodeR67                                         ReportStatsExamplesActionsCode = "R67"
	ReportStatsExamplesActionsCodeR68                                         ReportStatsExamplesActionsCode = "R68"
	ReportStatsExamplesActionsCodeR69                                         ReportStatsExamplesActionsCode = "R69"
	ReportStatsExamplesActionsCodeR70                                         ReportStatsExamplesActionsCode = "R70"
	ReportStatsExamplesActionsCodeR71                                         ReportStatsExamplesActionsCode = "R71"
	ReportStatsExamplesActionsCodeR72                                         ReportStatsExamplesActionsCode = "R72"
	ReportStatsExamplesActionsCodeR73                                         ReportStatsExamplesActionsCode = "R73"
	ReportStatsExamplesActionsCodeR74                                         ReportStatsExamplesActionsCode = "R74"
	ReportStatsExamplesActionsCodeR75                                         ReportStatsExamplesActionsCode = "R75"
	ReportStatsExamplesActionsCodeR76                                         ReportStatsExamplesActionsCode = "R76"
	ReportStatsExamplesActionsCodeR77                                         ReportStatsExamplesActionsCode = "R77"
	ReportStatsExamplesActionsCodeR80                                         ReportStatsExamplesActionsCode = "R80"
	ReportStatsExamplesActionsCodeR81                                         ReportStatsExamplesActionsCode = "R81"
	ReportStatsExamplesActionsCodeR82                                         ReportStatsExamplesActionsCode = "R82"
	ReportStatsExamplesActionsCodeR83                                         ReportStatsExamplesActionsCode = "R83"
	ReportStatsExamplesActionsCodeR84                                         ReportStatsExamplesActionsCode = "R84"
	ReportStatsExamplesActionsCodeR85                                         ReportStatsExamplesActionsCode = "R85"
)

func (r ReportStatsExamplesActionsCode) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsCodeAccountDailySpendLimitExceeded, ReportStatsExamplesActionsCodeAccountDelinquent, ReportStatsExamplesActionsCodeAccountInactive, ReportStatsExamplesActionsCodeAccountLifetimeSpendLimitExceeded, ReportStatsExamplesActionsCodeAccountMonthlySpendLimitExceeded, ReportStatsExamplesActionsCodeAccountPaused, ReportStatsExamplesActionsCodeAccountUnderReview, ReportStatsExamplesActionsCodeAddressIncorrect, ReportStatsExamplesActionsCodeApproved, ReportStatsExamplesActionsCodeAuthRuleAllowedCountry, ReportStatsExamplesActionsCodeAuthRuleAllowedMcc, ReportStatsExamplesActionsCodeAuthRuleBlockedCountry, ReportStatsExamplesActionsCodeAuthRuleBlockedMcc, ReportStatsExamplesActionsCodeAuthRule, ReportStatsExamplesActionsCodeCardClosed, ReportStatsExamplesActionsCodeCardCryptogramValidationFailure, ReportStatsExamplesActionsCodeCardExpired, ReportStatsExamplesActionsCodeCardExpiryDateIncorrect, ReportStatsExamplesActionsCodeCardInvalid, ReportStatsExamplesActionsCodeCardNotActivated, ReportStatsExamplesActionsCodeCardPaused, ReportStatsExamplesActionsCodeCardPinIncorrect, ReportStatsExamplesActionsCodeCardRestricted, ReportStatsExamplesActionsCodeCardSecurityCodeIncorrect, ReportStatsExamplesActionsCodeCardSpendLimitExceeded, ReportStatsExamplesActionsCodeContactCardIssuer, ReportStatsExamplesActionsCodeCustomerAsaTimeout, ReportStatsExamplesActionsCodeCustomAsaResult, ReportStatsExamplesActionsCodeDeclined, ReportStatsExamplesActionsCodeDoNotHonor, ReportStatsExamplesActionsCodeDriverNumberInvalid, ReportStatsExamplesActionsCodeFormatError, ReportStatsExamplesActionsCodeInsufficientFundingSourceBalance, ReportStatsExamplesActionsCodeInsufficientFunds, ReportStatsExamplesActionsCodeLithicSystemError, ReportStatsExamplesActionsCodeLithicSystemRateLimit, ReportStatsExamplesActionsCodeMalformedAsaResponse, ReportStatsExamplesActionsCodeMerchantInvalid, ReportStatsExamplesActionsCodeMerchantLockedCardAttemptedElsewhere, ReportStatsExamplesActionsCodeMerchantNotPermitted, ReportStatsExamplesActionsCodeOverReversalAttempted, ReportStatsExamplesActionsCodePinBlocked, ReportStatsExamplesActionsCodeProgramCardSpendLimitExceeded, ReportStatsExamplesActionsCodeProgramSuspended, ReportStatsExamplesActionsCodeProgramUsageRestriction, ReportStatsExamplesActionsCodeReversalUnmatched, ReportStatsExamplesActionsCodeSecurityViolation, ReportStatsExamplesActionsCodeSingleUseCardReattempted, ReportStatsExamplesActionsCodeSuspectedFraud, ReportStatsExamplesActionsCodeTransactionInvalid, ReportStatsExamplesActionsCodeTransactionNotPermittedToAcquirerOrTerminal, ReportStatsExamplesActionsCodeTransactionNotPermittedToIssuerOrCardholder, ReportStatsExamplesActionsCodeTransactionPreviouslyCompleted, ReportStatsExamplesActionsCodeUnauthorizedMerchant, ReportStatsExamplesActionsCodeVehicleNumberInvalid, ReportStatsExamplesActionsCodeCardholderChallenged, ReportStatsExamplesActionsCodeCardholderChallengeFailed, ReportStatsExamplesActionsCodeR01, ReportStatsExamplesActionsCodeR02, ReportStatsExamplesActionsCodeR03, ReportStatsExamplesActionsCodeR04, ReportStatsExamplesActionsCodeR05, ReportStatsExamplesActionsCodeR06, ReportStatsExamplesActionsCodeR07, ReportStatsExamplesActionsCodeR08, ReportStatsExamplesActionsCodeR09, ReportStatsExamplesActionsCodeR10, ReportStatsExamplesActionsCodeR11, ReportStatsExamplesActionsCodeR12, ReportStatsExamplesActionsCodeR13, ReportStatsExamplesActionsCodeR14, ReportStatsExamplesActionsCodeR15, ReportStatsExamplesActionsCodeR16, ReportStatsExamplesActionsCodeR17, ReportStatsExamplesActionsCodeR18, ReportStatsExamplesActionsCodeR19, ReportStatsExamplesActionsCodeR20, ReportStatsExamplesActionsCodeR21, ReportStatsExamplesActionsCodeR22, ReportStatsExamplesActionsCodeR23, ReportStatsExamplesActionsCodeR24, ReportStatsExamplesActionsCodeR25, ReportStatsExamplesActionsCodeR26, ReportStatsExamplesActionsCodeR27, ReportStatsExamplesActionsCodeR28, ReportStatsExamplesActionsCodeR29, ReportStatsExamplesActionsCodeR30, ReportStatsExamplesActionsCodeR31, ReportStatsExamplesActionsCodeR32, ReportStatsExamplesActionsCodeR33, ReportStatsExamplesActionsCodeR34, ReportStatsExamplesActionsCodeR35, ReportStatsExamplesActionsCodeR36, ReportStatsExamplesActionsCodeR37, ReportStatsExamplesActionsCodeR38, ReportStatsExamplesActionsCodeR39, ReportStatsExamplesActionsCodeR40, ReportStatsExamplesActionsCodeR41, ReportStatsExamplesActionsCodeR42, ReportStatsExamplesActionsCodeR43, ReportStatsExamplesActionsCodeR44, ReportStatsExamplesActionsCodeR45, ReportStatsExamplesActionsCodeR46, ReportStatsExamplesActionsCodeR47, ReportStatsExamplesActionsCodeR50, ReportStatsExamplesActionsCodeR51, ReportStatsExamplesActionsCodeR52, ReportStatsExamplesActionsCodeR53, ReportStatsExamplesActionsCodeR61, ReportStatsExamplesActionsCodeR62, ReportStatsExamplesActionsCodeR67, ReportStatsExamplesActionsCodeR68, ReportStatsExamplesActionsCodeR69, ReportStatsExamplesActionsCodeR70, ReportStatsExamplesActionsCodeR71, ReportStatsExamplesActionsCodeR72, ReportStatsExamplesActionsCodeR73, ReportStatsExamplesActionsCodeR74, ReportStatsExamplesActionsCodeR75, ReportStatsExamplesActionsCodeR76, ReportStatsExamplesActionsCodeR77, ReportStatsExamplesActionsCodeR80, ReportStatsExamplesActionsCodeR81, ReportStatsExamplesActionsCodeR82, ReportStatsExamplesActionsCodeR83, ReportStatsExamplesActionsCodeR84, ReportStatsExamplesActionsCodeR85:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type ReportStatsExamplesActionsReason string

const (
	ReportStatsExamplesActionsReasonAccountScore1                  ReportStatsExamplesActionsReason = "ACCOUNT_SCORE_1"
	ReportStatsExamplesActionsReasonDeviceScore1                   ReportStatsExamplesActionsReason = "DEVICE_SCORE_1"
	ReportStatsExamplesActionsReasonAllWalletDeclineReasonsPresent ReportStatsExamplesActionsReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	ReportStatsExamplesActionsReasonWalletRecommendedDecisionRed   ReportStatsExamplesActionsReason = "WALLET_RECOMMENDED_DECISION_RED"
	ReportStatsExamplesActionsReasonCvcMismatch                    ReportStatsExamplesActionsReason = "CVC_MISMATCH"
	ReportStatsExamplesActionsReasonCardExpiryMonthMismatch        ReportStatsExamplesActionsReason = "CARD_EXPIRY_MONTH_MISMATCH"
	ReportStatsExamplesActionsReasonCardExpiryYearMismatch         ReportStatsExamplesActionsReason = "CARD_EXPIRY_YEAR_MISMATCH"
	ReportStatsExamplesActionsReasonCardInvalidState               ReportStatsExamplesActionsReason = "CARD_INVALID_STATE"
	ReportStatsExamplesActionsReasonCustomerRedPath                ReportStatsExamplesActionsReason = "CUSTOMER_RED_PATH"
	ReportStatsExamplesActionsReasonInvalidCustomerResponse        ReportStatsExamplesActionsReason = "INVALID_CUSTOMER_RESPONSE"
	ReportStatsExamplesActionsReasonNetworkFailure                 ReportStatsExamplesActionsReason = "NETWORK_FAILURE"
	ReportStatsExamplesActionsReasonGenericDecline                 ReportStatsExamplesActionsReason = "GENERIC_DECLINE"
	ReportStatsExamplesActionsReasonDigitalCardArtRequired         ReportStatsExamplesActionsReason = "DIGITAL_CARD_ART_REQUIRED"
	ReportStatsExamplesActionsReasonWalletRecommendedTfa           ReportStatsExamplesActionsReason = "WALLET_RECOMMENDED_TFA"
	ReportStatsExamplesActionsReasonSuspiciousActivity             ReportStatsExamplesActionsReason = "SUSPICIOUS_ACTIVITY"
	ReportStatsExamplesActionsReasonDeviceRecentlyLost             ReportStatsExamplesActionsReason = "DEVICE_RECENTLY_LOST"
	ReportStatsExamplesActionsReasonTooManyRecentAttempts          ReportStatsExamplesActionsReason = "TOO_MANY_RECENT_ATTEMPTS"
	ReportStatsExamplesActionsReasonTooManyRecentTokens            ReportStatsExamplesActionsReason = "TOO_MANY_RECENT_TOKENS"
	ReportStatsExamplesActionsReasonTooManyDifferentCardholders    ReportStatsExamplesActionsReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	ReportStatsExamplesActionsReasonOutsideHomeTerritory           ReportStatsExamplesActionsReason = "OUTSIDE_HOME_TERRITORY"
	ReportStatsExamplesActionsReasonHasSuspendedTokens             ReportStatsExamplesActionsReason = "HAS_SUSPENDED_TOKENS"
	ReportStatsExamplesActionsReasonHighRisk                       ReportStatsExamplesActionsReason = "HIGH_RISK"
	ReportStatsExamplesActionsReasonAccountScoreLow                ReportStatsExamplesActionsReason = "ACCOUNT_SCORE_LOW"
	ReportStatsExamplesActionsReasonDeviceScoreLow                 ReportStatsExamplesActionsReason = "DEVICE_SCORE_LOW"
	ReportStatsExamplesActionsReasonCardStateTfa                   ReportStatsExamplesActionsReason = "CARD_STATE_TFA"
	ReportStatsExamplesActionsReasonHardcodedTfa                   ReportStatsExamplesActionsReason = "HARDCODED_TFA"
	ReportStatsExamplesActionsReasonCustomerRuleTfa                ReportStatsExamplesActionsReason = "CUSTOMER_RULE_TFA"
	ReportStatsExamplesActionsReasonDeviceHostCardEmulation        ReportStatsExamplesActionsReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r ReportStatsExamplesActionsReason) IsKnown() bool {
	switch r {
	case ReportStatsExamplesActionsReasonAccountScore1, ReportStatsExamplesActionsReasonDeviceScore1, ReportStatsExamplesActionsReasonAllWalletDeclineReasonsPresent, ReportStatsExamplesActionsReasonWalletRecommendedDecisionRed, ReportStatsExamplesActionsReasonCvcMismatch, ReportStatsExamplesActionsReasonCardExpiryMonthMismatch, ReportStatsExamplesActionsReasonCardExpiryYearMismatch, ReportStatsExamplesActionsReasonCardInvalidState, ReportStatsExamplesActionsReasonCustomerRedPath, ReportStatsExamplesActionsReasonInvalidCustomerResponse, ReportStatsExamplesActionsReasonNetworkFailure, ReportStatsExamplesActionsReasonGenericDecline, ReportStatsExamplesActionsReasonDigitalCardArtRequired, ReportStatsExamplesActionsReasonWalletRecommendedTfa, ReportStatsExamplesActionsReasonSuspiciousActivity, ReportStatsExamplesActionsReasonDeviceRecentlyLost, ReportStatsExamplesActionsReasonTooManyRecentAttempts, ReportStatsExamplesActionsReasonTooManyRecentTokens, ReportStatsExamplesActionsReasonTooManyDifferentCardholders, ReportStatsExamplesActionsReasonOutsideHomeTerritory, ReportStatsExamplesActionsReasonHasSuspendedTokens, ReportStatsExamplesActionsReasonHighRisk, ReportStatsExamplesActionsReasonAccountScoreLow, ReportStatsExamplesActionsReasonDeviceScoreLow, ReportStatsExamplesActionsReasonCardStateTfa, ReportStatsExamplesActionsReasonHardcodedTfa, ReportStatsExamplesActionsReasonCustomerRuleTfa, ReportStatsExamplesActionsReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

// The decision made by the rule for this event.
type ReportStatsExamplesDecision string

const (
	ReportStatsExamplesDecisionApproved   ReportStatsExamplesDecision = "APPROVED"
	ReportStatsExamplesDecisionDeclined   ReportStatsExamplesDecision = "DECLINED"
	ReportStatsExamplesDecisionChallenged ReportStatsExamplesDecision = "CHALLENGED"
)

func (r ReportStatsExamplesDecision) IsKnown() bool {
	switch r {
	case ReportStatsExamplesDecisionApproved, ReportStatsExamplesDecisionDeclined, ReportStatsExamplesDecisionChallenged:
		return true
	}
	return false
}

// A feature made available to the rule. The `name` field is the variable name used
// in the rule function signature. The `type` field determines which data the
// feature provides to the rule at evaluation time.
//
//   - `AUTHORIZATION`: The authorization request being evaluated. Only available for
//     AUTHORIZATION event stream rules.
//   - `AUTHENTICATION`: The 3DS authentication request being evaluated. Only
//     available for THREE_DS_AUTHENTICATION event stream rules.
//   - `TOKENIZATION`: The tokenization request being evaluated. Only available for
//     TOKENIZATION event stream rules.
//   - `ACH_RECEIPT`: The ACH receipt being evaluated. Only available for
//     ACH_CREDIT_RECEIPT and ACH_DEBIT_RECEIPT event stream rules.
//   - `CARD`: The card associated with the event. Available for AUTHORIZATION and
//     THREE_DS_AUTHENTICATION event stream rules.
//   - `ACCOUNT_HOLDER`: The account holder associated with the card. Available for
//     AUTHORIZATION and THREE_DS_AUTHENTICATION event stream rules.
//   - `IP_METADATA`: IP address metadata for the request. Available for
//     THREE_DS_AUTHENTICATION event stream rules.
//   - `SPEND_VELOCITY`: Spend velocity data for the card or account. Requires
//     `scope`, `period`, and optionally `filters` to configure the velocity
//     calculation. Available for AUTHORIZATION event stream rules.
type RuleFeature struct {
	Type    RuleFeatureType      `json:"type" api:"required"`
	Filters VelocityLimitFilters `json:"filters"`
	// The variable name for this feature in the rule function signature
	Name string `json:"name"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
	// The scope the velocity is calculated for
	Scope RuleFeatureScope `json:"scope"`
	JSON  ruleFeatureJSON  `json:"-"`
	union RuleFeatureUnion
}

// ruleFeatureJSON contains the JSON metadata for the struct [RuleFeature]
type ruleFeatureJSON struct {
	Type        apijson.Field
	Filters     apijson.Field
	Name        apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r ruleFeatureJSON) RawJSON() string {
	return r.raw
}

func (r *RuleFeature) UnmarshalJSON(data []byte) (err error) {
	*r = RuleFeature{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RuleFeatureUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [RuleFeatureAuthorizationFeature],
// [RuleFeatureAuthenticationFeature], [RuleFeatureTokenizationFeature],
// [RuleFeatureACHReceiptFeature], [RuleFeatureCardFeature],
// [RuleFeatureAccountHolderFeature], [RuleFeatureIPMetadataFeature],
// [RuleFeatureSpendVelocityFeature].
func (r RuleFeature) AsUnion() RuleFeatureUnion {
	return r.union
}

// A feature made available to the rule. The `name` field is the variable name used
// in the rule function signature. The `type` field determines which data the
// feature provides to the rule at evaluation time.
//
//   - `AUTHORIZATION`: The authorization request being evaluated. Only available for
//     AUTHORIZATION event stream rules.
//   - `AUTHENTICATION`: The 3DS authentication request being evaluated. Only
//     available for THREE_DS_AUTHENTICATION event stream rules.
//   - `TOKENIZATION`: The tokenization request being evaluated. Only available for
//     TOKENIZATION event stream rules.
//   - `ACH_RECEIPT`: The ACH receipt being evaluated. Only available for
//     ACH_CREDIT_RECEIPT and ACH_DEBIT_RECEIPT event stream rules.
//   - `CARD`: The card associated with the event. Available for AUTHORIZATION and
//     THREE_DS_AUTHENTICATION event stream rules.
//   - `ACCOUNT_HOLDER`: The account holder associated with the card. Available for
//     AUTHORIZATION and THREE_DS_AUTHENTICATION event stream rules.
//   - `IP_METADATA`: IP address metadata for the request. Available for
//     THREE_DS_AUTHENTICATION event stream rules.
//   - `SPEND_VELOCITY`: Spend velocity data for the card or account. Requires
//     `scope`, `period`, and optionally `filters` to configure the velocity
//     calculation. Available for AUTHORIZATION event stream rules.
//
// Union satisfied by [RuleFeatureAuthorizationFeature],
// [RuleFeatureAuthenticationFeature], [RuleFeatureTokenizationFeature],
// [RuleFeatureACHReceiptFeature], [RuleFeatureCardFeature],
// [RuleFeatureAccountHolderFeature], [RuleFeatureIPMetadataFeature] or
// [RuleFeatureSpendVelocityFeature].
type RuleFeatureUnion interface {
	implementsRuleFeature()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleFeatureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFeatureAuthorizationFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFeatureAuthenticationFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFeatureTokenizationFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFeatureACHReceiptFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFeatureCardFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFeatureAccountHolderFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFeatureIPMetadataFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFeatureSpendVelocityFeature{}),
		},
	)
}

type RuleFeatureAuthorizationFeature struct {
	Type RuleFeatureAuthorizationFeatureType `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name string                              `json:"name"`
	JSON ruleFeatureAuthorizationFeatureJSON `json:"-"`
}

// ruleFeatureAuthorizationFeatureJSON contains the JSON metadata for the struct
// [RuleFeatureAuthorizationFeature]
type ruleFeatureAuthorizationFeatureJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFeatureAuthorizationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFeatureAuthorizationFeatureJSON) RawJSON() string {
	return r.raw
}

func (r RuleFeatureAuthorizationFeature) implementsRuleFeature() {}

type RuleFeatureAuthorizationFeatureType string

const (
	RuleFeatureAuthorizationFeatureTypeAuthorization RuleFeatureAuthorizationFeatureType = "AUTHORIZATION"
)

func (r RuleFeatureAuthorizationFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureAuthorizationFeatureTypeAuthorization:
		return true
	}
	return false
}

type RuleFeatureAuthenticationFeature struct {
	Type RuleFeatureAuthenticationFeatureType `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name string                               `json:"name"`
	JSON ruleFeatureAuthenticationFeatureJSON `json:"-"`
}

// ruleFeatureAuthenticationFeatureJSON contains the JSON metadata for the struct
// [RuleFeatureAuthenticationFeature]
type ruleFeatureAuthenticationFeatureJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFeatureAuthenticationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFeatureAuthenticationFeatureJSON) RawJSON() string {
	return r.raw
}

func (r RuleFeatureAuthenticationFeature) implementsRuleFeature() {}

type RuleFeatureAuthenticationFeatureType string

const (
	RuleFeatureAuthenticationFeatureTypeAuthentication RuleFeatureAuthenticationFeatureType = "AUTHENTICATION"
)

func (r RuleFeatureAuthenticationFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureAuthenticationFeatureTypeAuthentication:
		return true
	}
	return false
}

type RuleFeatureTokenizationFeature struct {
	Type RuleFeatureTokenizationFeatureType `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name string                             `json:"name"`
	JSON ruleFeatureTokenizationFeatureJSON `json:"-"`
}

// ruleFeatureTokenizationFeatureJSON contains the JSON metadata for the struct
// [RuleFeatureTokenizationFeature]
type ruleFeatureTokenizationFeatureJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFeatureTokenizationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFeatureTokenizationFeatureJSON) RawJSON() string {
	return r.raw
}

func (r RuleFeatureTokenizationFeature) implementsRuleFeature() {}

type RuleFeatureTokenizationFeatureType string

const (
	RuleFeatureTokenizationFeatureTypeTokenization RuleFeatureTokenizationFeatureType = "TOKENIZATION"
)

func (r RuleFeatureTokenizationFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureTokenizationFeatureTypeTokenization:
		return true
	}
	return false
}

type RuleFeatureACHReceiptFeature struct {
	Type RuleFeatureACHReceiptFeatureType `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name string                           `json:"name"`
	JSON ruleFeatureACHReceiptFeatureJSON `json:"-"`
}

// ruleFeatureACHReceiptFeatureJSON contains the JSON metadata for the struct
// [RuleFeatureACHReceiptFeature]
type ruleFeatureACHReceiptFeatureJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFeatureACHReceiptFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFeatureACHReceiptFeatureJSON) RawJSON() string {
	return r.raw
}

func (r RuleFeatureACHReceiptFeature) implementsRuleFeature() {}

type RuleFeatureACHReceiptFeatureType string

const (
	RuleFeatureACHReceiptFeatureTypeACHReceipt RuleFeatureACHReceiptFeatureType = "ACH_RECEIPT"
)

func (r RuleFeatureACHReceiptFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureACHReceiptFeatureTypeACHReceipt:
		return true
	}
	return false
}

type RuleFeatureCardFeature struct {
	Type RuleFeatureCardFeatureType `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name string                     `json:"name"`
	JSON ruleFeatureCardFeatureJSON `json:"-"`
}

// ruleFeatureCardFeatureJSON contains the JSON metadata for the struct
// [RuleFeatureCardFeature]
type ruleFeatureCardFeatureJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFeatureCardFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFeatureCardFeatureJSON) RawJSON() string {
	return r.raw
}

func (r RuleFeatureCardFeature) implementsRuleFeature() {}

type RuleFeatureCardFeatureType string

const (
	RuleFeatureCardFeatureTypeCard RuleFeatureCardFeatureType = "CARD"
)

func (r RuleFeatureCardFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureCardFeatureTypeCard:
		return true
	}
	return false
}

type RuleFeatureAccountHolderFeature struct {
	Type RuleFeatureAccountHolderFeatureType `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name string                              `json:"name"`
	JSON ruleFeatureAccountHolderFeatureJSON `json:"-"`
}

// ruleFeatureAccountHolderFeatureJSON contains the JSON metadata for the struct
// [RuleFeatureAccountHolderFeature]
type ruleFeatureAccountHolderFeatureJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFeatureAccountHolderFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFeatureAccountHolderFeatureJSON) RawJSON() string {
	return r.raw
}

func (r RuleFeatureAccountHolderFeature) implementsRuleFeature() {}

type RuleFeatureAccountHolderFeatureType string

const (
	RuleFeatureAccountHolderFeatureTypeAccountHolder RuleFeatureAccountHolderFeatureType = "ACCOUNT_HOLDER"
)

func (r RuleFeatureAccountHolderFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureAccountHolderFeatureTypeAccountHolder:
		return true
	}
	return false
}

type RuleFeatureIPMetadataFeature struct {
	Type RuleFeatureIPMetadataFeatureType `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name string                           `json:"name"`
	JSON ruleFeatureIPMetadataFeatureJSON `json:"-"`
}

// ruleFeatureIPMetadataFeatureJSON contains the JSON metadata for the struct
// [RuleFeatureIPMetadataFeature]
type ruleFeatureIPMetadataFeatureJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFeatureIPMetadataFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFeatureIPMetadataFeatureJSON) RawJSON() string {
	return r.raw
}

func (r RuleFeatureIPMetadataFeature) implementsRuleFeature() {}

type RuleFeatureIPMetadataFeatureType string

const (
	RuleFeatureIPMetadataFeatureTypeIPMetadata RuleFeatureIPMetadataFeatureType = "IP_METADATA"
)

func (r RuleFeatureIPMetadataFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureIPMetadataFeatureTypeIPMetadata:
		return true
	}
	return false
}

type RuleFeatureSpendVelocityFeature struct {
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period" api:"required"`
	// The scope the velocity is calculated for
	Scope   RuleFeatureSpendVelocityFeatureScope `json:"scope" api:"required"`
	Type    RuleFeatureSpendVelocityFeatureType  `json:"type" api:"required"`
	Filters VelocityLimitFilters                 `json:"filters"`
	// The variable name for this feature in the rule function signature
	Name string                              `json:"name"`
	JSON ruleFeatureSpendVelocityFeatureJSON `json:"-"`
}

// ruleFeatureSpendVelocityFeatureJSON contains the JSON metadata for the struct
// [RuleFeatureSpendVelocityFeature]
type ruleFeatureSpendVelocityFeatureJSON struct {
	Period      apijson.Field
	Scope       apijson.Field
	Type        apijson.Field
	Filters     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFeatureSpendVelocityFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFeatureSpendVelocityFeatureJSON) RawJSON() string {
	return r.raw
}

func (r RuleFeatureSpendVelocityFeature) implementsRuleFeature() {}

// The scope the velocity is calculated for
type RuleFeatureSpendVelocityFeatureScope string

const (
	RuleFeatureSpendVelocityFeatureScopeCard    RuleFeatureSpendVelocityFeatureScope = "CARD"
	RuleFeatureSpendVelocityFeatureScopeAccount RuleFeatureSpendVelocityFeatureScope = "ACCOUNT"
)

func (r RuleFeatureSpendVelocityFeatureScope) IsKnown() bool {
	switch r {
	case RuleFeatureSpendVelocityFeatureScopeCard, RuleFeatureSpendVelocityFeatureScopeAccount:
		return true
	}
	return false
}

type RuleFeatureSpendVelocityFeatureType string

const (
	RuleFeatureSpendVelocityFeatureTypeSpendVelocity RuleFeatureSpendVelocityFeatureType = "SPEND_VELOCITY"
)

func (r RuleFeatureSpendVelocityFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureSpendVelocityFeatureTypeSpendVelocity:
		return true
	}
	return false
}

type RuleFeatureType string

const (
	RuleFeatureTypeAuthorization  RuleFeatureType = "AUTHORIZATION"
	RuleFeatureTypeAuthentication RuleFeatureType = "AUTHENTICATION"
	RuleFeatureTypeTokenization   RuleFeatureType = "TOKENIZATION"
	RuleFeatureTypeACHReceipt     RuleFeatureType = "ACH_RECEIPT"
	RuleFeatureTypeCard           RuleFeatureType = "CARD"
	RuleFeatureTypeAccountHolder  RuleFeatureType = "ACCOUNT_HOLDER"
	RuleFeatureTypeIPMetadata     RuleFeatureType = "IP_METADATA"
	RuleFeatureTypeSpendVelocity  RuleFeatureType = "SPEND_VELOCITY"
)

func (r RuleFeatureType) IsKnown() bool {
	switch r {
	case RuleFeatureTypeAuthorization, RuleFeatureTypeAuthentication, RuleFeatureTypeTokenization, RuleFeatureTypeACHReceipt, RuleFeatureTypeCard, RuleFeatureTypeAccountHolder, RuleFeatureTypeIPMetadata, RuleFeatureTypeSpendVelocity:
		return true
	}
	return false
}

// The scope the velocity is calculated for
type RuleFeatureScope string

const (
	RuleFeatureScopeCard    RuleFeatureScope = "CARD"
	RuleFeatureScopeAccount RuleFeatureScope = "ACCOUNT"
)

func (r RuleFeatureScope) IsKnown() bool {
	switch r {
	case RuleFeatureScopeCard, RuleFeatureScopeAccount:
		return true
	}
	return false
}

// A feature made available to the rule. The `name` field is the variable name used
// in the rule function signature. The `type` field determines which data the
// feature provides to the rule at evaluation time.
//
//   - `AUTHORIZATION`: The authorization request being evaluated. Only available for
//     AUTHORIZATION event stream rules.
//   - `AUTHENTICATION`: The 3DS authentication request being evaluated. Only
//     available for THREE_DS_AUTHENTICATION event stream rules.
//   - `TOKENIZATION`: The tokenization request being evaluated. Only available for
//     TOKENIZATION event stream rules.
//   - `ACH_RECEIPT`: The ACH receipt being evaluated. Only available for
//     ACH_CREDIT_RECEIPT and ACH_DEBIT_RECEIPT event stream rules.
//   - `CARD`: The card associated with the event. Available for AUTHORIZATION and
//     THREE_DS_AUTHENTICATION event stream rules.
//   - `ACCOUNT_HOLDER`: The account holder associated with the card. Available for
//     AUTHORIZATION and THREE_DS_AUTHENTICATION event stream rules.
//   - `IP_METADATA`: IP address metadata for the request. Available for
//     THREE_DS_AUTHENTICATION event stream rules.
//   - `SPEND_VELOCITY`: Spend velocity data for the card or account. Requires
//     `scope`, `period`, and optionally `filters` to configure the velocity
//     calculation. Available for AUTHORIZATION event stream rules.
type RuleFeatureParam struct {
	Type    param.Field[RuleFeatureType]           `json:"type" api:"required"`
	Filters param.Field[VelocityLimitFiltersParam] `json:"filters"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period param.Field[VelocityLimitPeriodUnionParam] `json:"period"`
	// The scope the velocity is calculated for
	Scope param.Field[RuleFeatureScope] `json:"scope"`
}

func (r RuleFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureParam) implementsRuleFeatureUnionParam() {}

// A feature made available to the rule. The `name` field is the variable name used
// in the rule function signature. The `type` field determines which data the
// feature provides to the rule at evaluation time.
//
//   - `AUTHORIZATION`: The authorization request being evaluated. Only available for
//     AUTHORIZATION event stream rules.
//   - `AUTHENTICATION`: The 3DS authentication request being evaluated. Only
//     available for THREE_DS_AUTHENTICATION event stream rules.
//   - `TOKENIZATION`: The tokenization request being evaluated. Only available for
//     TOKENIZATION event stream rules.
//   - `ACH_RECEIPT`: The ACH receipt being evaluated. Only available for
//     ACH_CREDIT_RECEIPT and ACH_DEBIT_RECEIPT event stream rules.
//   - `CARD`: The card associated with the event. Available for AUTHORIZATION and
//     THREE_DS_AUTHENTICATION event stream rules.
//   - `ACCOUNT_HOLDER`: The account holder associated with the card. Available for
//     AUTHORIZATION and THREE_DS_AUTHENTICATION event stream rules.
//   - `IP_METADATA`: IP address metadata for the request. Available for
//     THREE_DS_AUTHENTICATION event stream rules.
//   - `SPEND_VELOCITY`: Spend velocity data for the card or account. Requires
//     `scope`, `period`, and optionally `filters` to configure the velocity
//     calculation. Available for AUTHORIZATION event stream rules.
//
// Satisfied by [RuleFeatureAuthorizationFeatureParam],
// [RuleFeatureAuthenticationFeatureParam], [RuleFeatureTokenizationFeatureParam],
// [RuleFeatureACHReceiptFeatureParam], [RuleFeatureCardFeatureParam],
// [RuleFeatureAccountHolderFeatureParam], [RuleFeatureIPMetadataFeatureParam],
// [RuleFeatureSpendVelocityFeatureParam], [RuleFeatureParam].
type RuleFeatureUnionParam interface {
	implementsRuleFeatureUnionParam()
}

type RuleFeatureAuthorizationFeatureParam struct {
	Type param.Field[RuleFeatureAuthorizationFeatureType] `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
}

func (r RuleFeatureAuthorizationFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureAuthorizationFeatureParam) implementsRuleFeatureUnionParam() {}

type RuleFeatureAuthenticationFeatureParam struct {
	Type param.Field[RuleFeatureAuthenticationFeatureType] `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
}

func (r RuleFeatureAuthenticationFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureAuthenticationFeatureParam) implementsRuleFeatureUnionParam() {}

type RuleFeatureTokenizationFeatureParam struct {
	Type param.Field[RuleFeatureTokenizationFeatureType] `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
}

func (r RuleFeatureTokenizationFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureTokenizationFeatureParam) implementsRuleFeatureUnionParam() {}

type RuleFeatureACHReceiptFeatureParam struct {
	Type param.Field[RuleFeatureACHReceiptFeatureType] `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
}

func (r RuleFeatureACHReceiptFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureACHReceiptFeatureParam) implementsRuleFeatureUnionParam() {}

type RuleFeatureCardFeatureParam struct {
	Type param.Field[RuleFeatureCardFeatureType] `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
}

func (r RuleFeatureCardFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureCardFeatureParam) implementsRuleFeatureUnionParam() {}

type RuleFeatureAccountHolderFeatureParam struct {
	Type param.Field[RuleFeatureAccountHolderFeatureType] `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
}

func (r RuleFeatureAccountHolderFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureAccountHolderFeatureParam) implementsRuleFeatureUnionParam() {}

type RuleFeatureIPMetadataFeatureParam struct {
	Type param.Field[RuleFeatureIPMetadataFeatureType] `json:"type" api:"required"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
}

func (r RuleFeatureIPMetadataFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureIPMetadataFeatureParam) implementsRuleFeatureUnionParam() {}

type RuleFeatureSpendVelocityFeatureParam struct {
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period param.Field[VelocityLimitPeriodUnionParam] `json:"period" api:"required"`
	// The scope the velocity is calculated for
	Scope   param.Field[RuleFeatureSpendVelocityFeatureScope] `json:"scope" api:"required"`
	Type    param.Field[RuleFeatureSpendVelocityFeatureType]  `json:"type" api:"required"`
	Filters param.Field[VelocityLimitFiltersParam]            `json:"filters"`
	// The variable name for this feature in the rule function signature
	Name param.Field[string] `json:"name"`
}

func (r RuleFeatureSpendVelocityFeatureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFeatureSpendVelocityFeatureParam) implementsRuleFeatureUnionParam() {}

// Parameters for defining a TypeScript code rule
type TypescriptCodeParameters struct {
	// The TypeScript source code of the rule. Must define a `rule()` function that
	// accepts the declared features as positional arguments (in the same order as the
	// `features` array) and returns an array of actions.
	Code string `json:"code" api:"required"`
	// Features available to the TypeScript code at evaluation time
	Features []RuleFeature                `json:"features" api:"required"`
	JSON     typescriptCodeParametersJSON `json:"-"`
}

// typescriptCodeParametersJSON contains the JSON metadata for the struct
// [TypescriptCodeParameters]
type typescriptCodeParametersJSON struct {
	Code        apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TypescriptCodeParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r typescriptCodeParametersJSON) RawJSON() string {
	return r.raw
}

func (r TypescriptCodeParameters) implementsAuthRuleCurrentVersionParameters() {}

func (r TypescriptCodeParameters) implementsAuthRuleDraftVersionParameters() {}

func (r TypescriptCodeParameters) implementsAuthRuleVersionParameters() {}

type VelocityLimitFilters struct {
	// ISO-3166-1 alpha-3 Country Codes to exclude from the velocity calculation.
	// Transactions matching any of the provided will be excluded from the calculated
	// velocity.
	ExcludeCountries []string `json:"exclude_countries" api:"nullable"`
	// Merchant Category Codes to exclude from the velocity calculation. Transactions
	// matching this MCC will be excluded from the calculated velocity.
	ExcludeMccs []string `json:"exclude_mccs" api:"nullable"`
	// ISO-3166-1 alpha-3 Country Codes to include in the velocity calculation.
	// Transactions not matching any of the provided will not be included in the
	// calculated velocity.
	IncludeCountries []string `json:"include_countries" api:"nullable"`
	// Merchant Category Codes to include in the velocity calculation. Transactions not
	// matching this MCC will not be included in the calculated velocity.
	IncludeMccs []string `json:"include_mccs" api:"nullable"`
	// PAN entry modes to include in the velocity calculation. Transactions not
	// matching any of the provided will not be included in the calculated velocity.
	IncludePanEntryModes []VelocityLimitFiltersIncludePanEntryMode `json:"include_pan_entry_modes" api:"nullable"`
	JSON                 velocityLimitFiltersJSON                  `json:"-"`
}

// velocityLimitFiltersJSON contains the JSON metadata for the struct
// [VelocityLimitFilters]
type velocityLimitFiltersJSON struct {
	ExcludeCountries     apijson.Field
	ExcludeMccs          apijson.Field
	IncludeCountries     apijson.Field
	IncludeMccs          apijson.Field
	IncludePanEntryModes apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *VelocityLimitFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitFiltersJSON) RawJSON() string {
	return r.raw
}

type VelocityLimitFiltersIncludePanEntryMode string

const (
	VelocityLimitFiltersIncludePanEntryModeAutoEntry           VelocityLimitFiltersIncludePanEntryMode = "AUTO_ENTRY"
	VelocityLimitFiltersIncludePanEntryModeBarCode             VelocityLimitFiltersIncludePanEntryMode = "BAR_CODE"
	VelocityLimitFiltersIncludePanEntryModeContactless         VelocityLimitFiltersIncludePanEntryMode = "CONTACTLESS"
	VelocityLimitFiltersIncludePanEntryModeCredentialOnFile    VelocityLimitFiltersIncludePanEntryMode = "CREDENTIAL_ON_FILE"
	VelocityLimitFiltersIncludePanEntryModeEcommerce           VelocityLimitFiltersIncludePanEntryMode = "ECOMMERCE"
	VelocityLimitFiltersIncludePanEntryModeErrorKeyed          VelocityLimitFiltersIncludePanEntryMode = "ERROR_KEYED"
	VelocityLimitFiltersIncludePanEntryModeErrorMagneticStripe VelocityLimitFiltersIncludePanEntryMode = "ERROR_MAGNETIC_STRIPE"
	VelocityLimitFiltersIncludePanEntryModeIcc                 VelocityLimitFiltersIncludePanEntryMode = "ICC"
	VelocityLimitFiltersIncludePanEntryModeKeyEntered          VelocityLimitFiltersIncludePanEntryMode = "KEY_ENTERED"
	VelocityLimitFiltersIncludePanEntryModeMagneticStripe      VelocityLimitFiltersIncludePanEntryMode = "MAGNETIC_STRIPE"
	VelocityLimitFiltersIncludePanEntryModeManual              VelocityLimitFiltersIncludePanEntryMode = "MANUAL"
	VelocityLimitFiltersIncludePanEntryModeOcr                 VelocityLimitFiltersIncludePanEntryMode = "OCR"
	VelocityLimitFiltersIncludePanEntryModeSecureCardless      VelocityLimitFiltersIncludePanEntryMode = "SECURE_CARDLESS"
	VelocityLimitFiltersIncludePanEntryModeUnspecified         VelocityLimitFiltersIncludePanEntryMode = "UNSPECIFIED"
	VelocityLimitFiltersIncludePanEntryModeUnknown             VelocityLimitFiltersIncludePanEntryMode = "UNKNOWN"
)

func (r VelocityLimitFiltersIncludePanEntryMode) IsKnown() bool {
	switch r {
	case VelocityLimitFiltersIncludePanEntryModeAutoEntry, VelocityLimitFiltersIncludePanEntryModeBarCode, VelocityLimitFiltersIncludePanEntryModeContactless, VelocityLimitFiltersIncludePanEntryModeCredentialOnFile, VelocityLimitFiltersIncludePanEntryModeEcommerce, VelocityLimitFiltersIncludePanEntryModeErrorKeyed, VelocityLimitFiltersIncludePanEntryModeErrorMagneticStripe, VelocityLimitFiltersIncludePanEntryModeIcc, VelocityLimitFiltersIncludePanEntryModeKeyEntered, VelocityLimitFiltersIncludePanEntryModeMagneticStripe, VelocityLimitFiltersIncludePanEntryModeManual, VelocityLimitFiltersIncludePanEntryModeOcr, VelocityLimitFiltersIncludePanEntryModeSecureCardless, VelocityLimitFiltersIncludePanEntryModeUnspecified, VelocityLimitFiltersIncludePanEntryModeUnknown:
		return true
	}
	return false
}

type VelocityLimitFiltersParam struct {
	// ISO-3166-1 alpha-3 Country Codes to exclude from the velocity calculation.
	// Transactions matching any of the provided will be excluded from the calculated
	// velocity.
	ExcludeCountries param.Field[[]string] `json:"exclude_countries"`
	// Merchant Category Codes to exclude from the velocity calculation. Transactions
	// matching this MCC will be excluded from the calculated velocity.
	ExcludeMccs param.Field[[]string] `json:"exclude_mccs"`
	// ISO-3166-1 alpha-3 Country Codes to include in the velocity calculation.
	// Transactions not matching any of the provided will not be included in the
	// calculated velocity.
	IncludeCountries param.Field[[]string] `json:"include_countries"`
	// Merchant Category Codes to include in the velocity calculation. Transactions not
	// matching this MCC will not be included in the calculated velocity.
	IncludeMccs param.Field[[]string] `json:"include_mccs"`
	// PAN entry modes to include in the velocity calculation. Transactions not
	// matching any of the provided will not be included in the calculated velocity.
	IncludePanEntryModes param.Field[[]VelocityLimitFiltersIncludePanEntryMode] `json:"include_pan_entry_modes"`
}

func (r VelocityLimitFiltersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VelocityLimitParams struct {
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period" api:"required"`
	// The scope the velocity is calculated for
	Scope   VelocityLimitParamsScope `json:"scope" api:"required"`
	Filters VelocityLimitFilters     `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount" api:"nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64                   `json:"limit_count" api:"nullable"`
	JSON       velocityLimitParamsJSON `json:"-"`
}

// velocityLimitParamsJSON contains the JSON metadata for the struct
// [VelocityLimitParams]
type velocityLimitParamsJSON struct {
	Period      apijson.Field
	Scope       apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VelocityLimitParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitParamsJSON) RawJSON() string {
	return r.raw
}

func (r VelocityLimitParams) implementsAuthRuleCurrentVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleDraftVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleVersionParameters() {}

// The scope the velocity is calculated for
type VelocityLimitParamsScope string

const (
	VelocityLimitParamsScopeCard    VelocityLimitParamsScope = "CARD"
	VelocityLimitParamsScopeAccount VelocityLimitParamsScope = "ACCOUNT"
)

func (r VelocityLimitParamsScope) IsKnown() bool {
	switch r {
	case VelocityLimitParamsScopeCard, VelocityLimitParamsScopeAccount:
		return true
	}
	return false
}

// Velocity over the current day since 00:00 / 12 AM in Eastern Time
type VelocityLimitPeriod struct {
	Type VelocityLimitPeriodType `json:"type" api:"required"`
	// The day of the month to start from. Accepts values from 1 to 31, and will reset
	// at the end of the month if the day exceeds the number of days in the month.
	// Defaults to the 1st of the month if not specified.
	DayOfMonth int64 `json:"day_of_month"`
	// The day of the week to start the week from. Following ISO-8601, 1 is Monday and
	// 7 is Sunday. Defaults to Monday if not specified.
	DayOfWeek int64 `json:"day_of_week"`
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Duration int64 `json:"duration"`
	// The month to start from. 1 is January and 12 is December. Defaults to January if
	// not specified.
	Month int64                   `json:"month"`
	JSON  velocityLimitPeriodJSON `json:"-"`
	union VelocityLimitPeriodUnion
}

// velocityLimitPeriodJSON contains the JSON metadata for the struct
// [VelocityLimitPeriod]
type velocityLimitPeriodJSON struct {
	Type        apijson.Field
	DayOfMonth  apijson.Field
	DayOfWeek   apijson.Field
	Duration    apijson.Field
	Month       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r velocityLimitPeriodJSON) RawJSON() string {
	return r.raw
}

func (r *VelocityLimitPeriod) UnmarshalJSON(data []byte) (err error) {
	*r = VelocityLimitPeriod{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VelocityLimitPeriodUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [VelocityLimitPeriodTrailingWindowObject], [VelocityLimitPeriodFixedWindowDay],
// [VelocityLimitPeriodFixedWindowWeek], [VelocityLimitPeriodFixedWindowMonth],
// [VelocityLimitPeriodFixedWindowYear].
func (r VelocityLimitPeriod) AsUnion() VelocityLimitPeriodUnion {
	return r.union
}

// Velocity over the current day since 00:00 / 12 AM in Eastern Time
//
// Union satisfied by [VelocityLimitPeriodTrailingWindowObject],
// [VelocityLimitPeriodFixedWindowDay], [VelocityLimitPeriodFixedWindowWeek],
// [VelocityLimitPeriodFixedWindowMonth] or [VelocityLimitPeriodFixedWindowYear].
type VelocityLimitPeriodUnion interface {
	implementsVelocityLimitPeriod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VelocityLimitPeriodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitPeriodTrailingWindowObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitPeriodFixedWindowDay{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitPeriodFixedWindowWeek{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitPeriodFixedWindowMonth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitPeriodFixedWindowYear{}),
		},
	)
}

type VelocityLimitPeriodTrailingWindowObject struct {
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Duration int64                                       `json:"duration" api:"required"`
	Type     VelocityLimitPeriodTrailingWindowObjectType `json:"type" api:"required"`
	JSON     velocityLimitPeriodTrailingWindowObjectJSON `json:"-"`
}

// velocityLimitPeriodTrailingWindowObjectJSON contains the JSON metadata for the
// struct [VelocityLimitPeriodTrailingWindowObject]
type velocityLimitPeriodTrailingWindowObjectJSON struct {
	Duration    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VelocityLimitPeriodTrailingWindowObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitPeriodTrailingWindowObjectJSON) RawJSON() string {
	return r.raw
}

func (r VelocityLimitPeriodTrailingWindowObject) implementsVelocityLimitPeriod() {}

type VelocityLimitPeriodTrailingWindowObjectType string

const (
	VelocityLimitPeriodTrailingWindowObjectTypeCustom VelocityLimitPeriodTrailingWindowObjectType = "CUSTOM"
)

func (r VelocityLimitPeriodTrailingWindowObjectType) IsKnown() bool {
	switch r {
	case VelocityLimitPeriodTrailingWindowObjectTypeCustom:
		return true
	}
	return false
}

// Velocity over the current day since 00:00 / 12 AM in Eastern Time
type VelocityLimitPeriodFixedWindowDay struct {
	Type VelocityLimitPeriodFixedWindowDayType `json:"type" api:"required"`
	JSON velocityLimitPeriodFixedWindowDayJSON `json:"-"`
}

// velocityLimitPeriodFixedWindowDayJSON contains the JSON metadata for the struct
// [VelocityLimitPeriodFixedWindowDay]
type velocityLimitPeriodFixedWindowDayJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VelocityLimitPeriodFixedWindowDay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitPeriodFixedWindowDayJSON) RawJSON() string {
	return r.raw
}

func (r VelocityLimitPeriodFixedWindowDay) implementsVelocityLimitPeriod() {}

type VelocityLimitPeriodFixedWindowDayType string

const (
	VelocityLimitPeriodFixedWindowDayTypeDay VelocityLimitPeriodFixedWindowDayType = "DAY"
)

func (r VelocityLimitPeriodFixedWindowDayType) IsKnown() bool {
	switch r {
	case VelocityLimitPeriodFixedWindowDayTypeDay:
		return true
	}
	return false
}

// Velocity over the current week since 00:00 / 12 AM in Eastern Time on specified
// `day_of_week`
type VelocityLimitPeriodFixedWindowWeek struct {
	Type VelocityLimitPeriodFixedWindowWeekType `json:"type" api:"required"`
	// The day of the week to start the week from. Following ISO-8601, 1 is Monday and
	// 7 is Sunday. Defaults to Monday if not specified.
	DayOfWeek int64                                  `json:"day_of_week"`
	JSON      velocityLimitPeriodFixedWindowWeekJSON `json:"-"`
}

// velocityLimitPeriodFixedWindowWeekJSON contains the JSON metadata for the struct
// [VelocityLimitPeriodFixedWindowWeek]
type velocityLimitPeriodFixedWindowWeekJSON struct {
	Type        apijson.Field
	DayOfWeek   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VelocityLimitPeriodFixedWindowWeek) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitPeriodFixedWindowWeekJSON) RawJSON() string {
	return r.raw
}

func (r VelocityLimitPeriodFixedWindowWeek) implementsVelocityLimitPeriod() {}

type VelocityLimitPeriodFixedWindowWeekType string

const (
	VelocityLimitPeriodFixedWindowWeekTypeWeek VelocityLimitPeriodFixedWindowWeekType = "WEEK"
)

func (r VelocityLimitPeriodFixedWindowWeekType) IsKnown() bool {
	switch r {
	case VelocityLimitPeriodFixedWindowWeekTypeWeek:
		return true
	}
	return false
}

// Velocity over the current month since 00:00 / 12 AM in Eastern Time on specified
// `day_of_month`.
type VelocityLimitPeriodFixedWindowMonth struct {
	Type VelocityLimitPeriodFixedWindowMonthType `json:"type" api:"required"`
	// The day of the month to start from. Accepts values from 1 to 31, and will reset
	// at the end of the month if the day exceeds the number of days in the month.
	// Defaults to the 1st of the month if not specified.
	DayOfMonth int64                                   `json:"day_of_month"`
	JSON       velocityLimitPeriodFixedWindowMonthJSON `json:"-"`
}

// velocityLimitPeriodFixedWindowMonthJSON contains the JSON metadata for the
// struct [VelocityLimitPeriodFixedWindowMonth]
type velocityLimitPeriodFixedWindowMonthJSON struct {
	Type        apijson.Field
	DayOfMonth  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VelocityLimitPeriodFixedWindowMonth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitPeriodFixedWindowMonthJSON) RawJSON() string {
	return r.raw
}

func (r VelocityLimitPeriodFixedWindowMonth) implementsVelocityLimitPeriod() {}

type VelocityLimitPeriodFixedWindowMonthType string

const (
	VelocityLimitPeriodFixedWindowMonthTypeMonth VelocityLimitPeriodFixedWindowMonthType = "MONTH"
)

func (r VelocityLimitPeriodFixedWindowMonthType) IsKnown() bool {
	switch r {
	case VelocityLimitPeriodFixedWindowMonthTypeMonth:
		return true
	}
	return false
}

// Velocity over the current year since 00:00 / 12 AM in Eastern Time on specified
// `month` and `day_of_month`. This validates the month and day of the year to
// start from is a real date. In the event that February 29th is selected, in
// non-leap years, the window will start from February 28th.
type VelocityLimitPeriodFixedWindowYear struct {
	Type VelocityLimitPeriodFixedWindowYearType `json:"type" api:"required"`
	// The day of the month to start from. Defaults to the 1st of the month if not
	// specified.
	DayOfMonth int64 `json:"day_of_month"`
	// The month to start from. 1 is January and 12 is December. Defaults to January if
	// not specified.
	Month int64                                  `json:"month"`
	JSON  velocityLimitPeriodFixedWindowYearJSON `json:"-"`
}

// velocityLimitPeriodFixedWindowYearJSON contains the JSON metadata for the struct
// [VelocityLimitPeriodFixedWindowYear]
type velocityLimitPeriodFixedWindowYearJSON struct {
	Type        apijson.Field
	DayOfMonth  apijson.Field
	Month       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VelocityLimitPeriodFixedWindowYear) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitPeriodFixedWindowYearJSON) RawJSON() string {
	return r.raw
}

func (r VelocityLimitPeriodFixedWindowYear) implementsVelocityLimitPeriod() {}

type VelocityLimitPeriodFixedWindowYearType string

const (
	VelocityLimitPeriodFixedWindowYearTypeYear VelocityLimitPeriodFixedWindowYearType = "YEAR"
)

func (r VelocityLimitPeriodFixedWindowYearType) IsKnown() bool {
	switch r {
	case VelocityLimitPeriodFixedWindowYearTypeYear:
		return true
	}
	return false
}

type VelocityLimitPeriodType string

const (
	VelocityLimitPeriodTypeCustom VelocityLimitPeriodType = "CUSTOM"
	VelocityLimitPeriodTypeDay    VelocityLimitPeriodType = "DAY"
	VelocityLimitPeriodTypeWeek   VelocityLimitPeriodType = "WEEK"
	VelocityLimitPeriodTypeMonth  VelocityLimitPeriodType = "MONTH"
	VelocityLimitPeriodTypeYear   VelocityLimitPeriodType = "YEAR"
)

func (r VelocityLimitPeriodType) IsKnown() bool {
	switch r {
	case VelocityLimitPeriodTypeCustom, VelocityLimitPeriodTypeDay, VelocityLimitPeriodTypeWeek, VelocityLimitPeriodTypeMonth, VelocityLimitPeriodTypeYear:
		return true
	}
	return false
}

// Velocity over the current day since 00:00 / 12 AM in Eastern Time
type VelocityLimitPeriodParam struct {
	Type param.Field[VelocityLimitPeriodType] `json:"type" api:"required"`
	// The day of the month to start from. Accepts values from 1 to 31, and will reset
	// at the end of the month if the day exceeds the number of days in the month.
	// Defaults to the 1st of the month if not specified.
	DayOfMonth param.Field[int64] `json:"day_of_month"`
	// The day of the week to start the week from. Following ISO-8601, 1 is Monday and
	// 7 is Sunday. Defaults to Monday if not specified.
	DayOfWeek param.Field[int64] `json:"day_of_week"`
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Duration param.Field[int64] `json:"duration"`
	// The month to start from. 1 is January and 12 is December. Defaults to January if
	// not specified.
	Month param.Field[int64] `json:"month"`
}

func (r VelocityLimitPeriodParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VelocityLimitPeriodParam) implementsVelocityLimitPeriodUnionParam() {}

// Velocity over the current day since 00:00 / 12 AM in Eastern Time
//
// Satisfied by [VelocityLimitPeriodTrailingWindowObjectParam],
// [VelocityLimitPeriodFixedWindowDayParam],
// [VelocityLimitPeriodFixedWindowWeekParam],
// [VelocityLimitPeriodFixedWindowMonthParam],
// [VelocityLimitPeriodFixedWindowYearParam], [VelocityLimitPeriodParam].
type VelocityLimitPeriodUnionParam interface {
	implementsVelocityLimitPeriodUnionParam()
}

type VelocityLimitPeriodTrailingWindowObjectParam struct {
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Duration param.Field[int64]                                       `json:"duration" api:"required"`
	Type     param.Field[VelocityLimitPeriodTrailingWindowObjectType] `json:"type" api:"required"`
}

func (r VelocityLimitPeriodTrailingWindowObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VelocityLimitPeriodTrailingWindowObjectParam) implementsVelocityLimitPeriodUnionParam() {}

// Velocity over the current day since 00:00 / 12 AM in Eastern Time
type VelocityLimitPeriodFixedWindowDayParam struct {
	Type param.Field[VelocityLimitPeriodFixedWindowDayType] `json:"type" api:"required"`
}

func (r VelocityLimitPeriodFixedWindowDayParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VelocityLimitPeriodFixedWindowDayParam) implementsVelocityLimitPeriodUnionParam() {}

// Velocity over the current week since 00:00 / 12 AM in Eastern Time on specified
// `day_of_week`
type VelocityLimitPeriodFixedWindowWeekParam struct {
	Type param.Field[VelocityLimitPeriodFixedWindowWeekType] `json:"type" api:"required"`
	// The day of the week to start the week from. Following ISO-8601, 1 is Monday and
	// 7 is Sunday. Defaults to Monday if not specified.
	DayOfWeek param.Field[int64] `json:"day_of_week"`
}

func (r VelocityLimitPeriodFixedWindowWeekParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VelocityLimitPeriodFixedWindowWeekParam) implementsVelocityLimitPeriodUnionParam() {}

// Velocity over the current month since 00:00 / 12 AM in Eastern Time on specified
// `day_of_month`.
type VelocityLimitPeriodFixedWindowMonthParam struct {
	Type param.Field[VelocityLimitPeriodFixedWindowMonthType] `json:"type" api:"required"`
	// The day of the month to start from. Accepts values from 1 to 31, and will reset
	// at the end of the month if the day exceeds the number of days in the month.
	// Defaults to the 1st of the month if not specified.
	DayOfMonth param.Field[int64] `json:"day_of_month"`
}

func (r VelocityLimitPeriodFixedWindowMonthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VelocityLimitPeriodFixedWindowMonthParam) implementsVelocityLimitPeriodUnionParam() {}

// Velocity over the current year since 00:00 / 12 AM in Eastern Time on specified
// `month` and `day_of_month`. This validates the month and day of the year to
// start from is a real date. In the event that February 29th is selected, in
// non-leap years, the window will start from February 28th.
type VelocityLimitPeriodFixedWindowYearParam struct {
	Type param.Field[VelocityLimitPeriodFixedWindowYearType] `json:"type" api:"required"`
	// The day of the month to start from. Defaults to the 1st of the month if not
	// specified.
	DayOfMonth param.Field[int64] `json:"day_of_month"`
	// The month to start from. 1 is January and 12 is December. Defaults to January if
	// not specified.
	Month param.Field[int64] `json:"month"`
}

func (r VelocityLimitPeriodFixedWindowYearParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VelocityLimitPeriodFixedWindowYearParam) implementsVelocityLimitPeriodUnionParam() {}

// Result of an Auth Rule evaluation
type AuthRuleV2ListResultsResponse struct {
	// Globally unique identifier for the evaluation
	Token string `json:"token" api:"required" format:"uuid"`
	// This field can have the runtime type of
	// [[]AuthRuleV2ListResultsResponseAuthorizationResultAction],
	// [[]AuthRuleV2ListResultsResponseAuthentication3DsResultAction],
	// [[]AuthRuleV2ListResultsResponseTokenizationResultAction],
	// [[]AuthRuleV2ListResultsResponseACHResultAction].
	Actions interface{} `json:"actions" api:"required"`
	// The Auth Rule token
	AuthRuleToken string `json:"auth_rule_token" api:"required" format:"uuid"`
	// Timestamp of the rule evaluation
	EvaluationTime time.Time `json:"evaluation_time" api:"required" format:"date-time"`
	// The event stream during which the rule was evaluated
	EventStream AuthRuleV2ListResultsResponseEventStream `json:"event_stream" api:"required"`
	// Token of the event that triggered the evaluation
	EventToken string `json:"event_token" api:"required" format:"uuid"`
	// The state of the Auth Rule
	Mode AuthRuleV2ListResultsResponseMode `json:"mode" api:"required"`
	// Version of the rule that was evaluated
	RuleVersion int64 `json:"rule_version" api:"required"`
	// The token of the transaction that triggered the rule evaluation
	TransactionToken string                            `json:"transaction_token" api:"required,nullable" format:"uuid"`
	JSON             authRuleV2ListResultsResponseJSON `json:"-"`
	union            AuthRuleV2ListResultsResponseUnion
}

// authRuleV2ListResultsResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ListResultsResponse]
type authRuleV2ListResultsResponseJSON struct {
	Token            apijson.Field
	Actions          apijson.Field
	AuthRuleToken    apijson.Field
	EvaluationTime   apijson.Field
	EventStream      apijson.Field
	EventToken       apijson.Field
	Mode             apijson.Field
	RuleVersion      apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r authRuleV2ListResultsResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2ListResultsResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2ListResultsResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2ListResultsResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2ListResultsResponseAuthorizationResult],
// [AuthRuleV2ListResultsResponseAuthentication3DSResult],
// [AuthRuleV2ListResultsResponseTokenizationResult],
// [AuthRuleV2ListResultsResponseACHResult].
func (r AuthRuleV2ListResultsResponse) AsUnion() AuthRuleV2ListResultsResponseUnion {
	return r.union
}

// Result of an Auth Rule evaluation
//
// Union satisfied by [AuthRuleV2ListResultsResponseAuthorizationResult],
// [AuthRuleV2ListResultsResponseAuthentication3DSResult],
// [AuthRuleV2ListResultsResponseTokenizationResult] or
// [AuthRuleV2ListResultsResponseACHResult].
type AuthRuleV2ListResultsResponseUnion interface {
	implementsAuthRuleV2ListResultsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResultsResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseAuthorizationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseAuthentication3DSResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseTokenizationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseACHResult{}),
		},
	)
}

type AuthRuleV2ListResultsResponseAuthorizationResult struct {
	// Globally unique identifier for the evaluation
	Token string `json:"token" api:"required" format:"uuid"`
	// Actions returned by the rule evaluation
	Actions []AuthRuleV2ListResultsResponseAuthorizationResultAction `json:"actions" api:"required"`
	// The Auth Rule token
	AuthRuleToken string `json:"auth_rule_token" api:"required" format:"uuid"`
	// Timestamp of the rule evaluation
	EvaluationTime time.Time `json:"evaluation_time" api:"required" format:"date-time"`
	// The event stream during which the rule was evaluated
	EventStream AuthRuleV2ListResultsResponseAuthorizationResultEventStream `json:"event_stream" api:"required"`
	// Token of the event that triggered the evaluation
	EventToken string `json:"event_token" api:"required" format:"uuid"`
	// The state of the Auth Rule
	Mode AuthRuleV2ListResultsResponseAuthorizationResultMode `json:"mode" api:"required"`
	// Version of the rule that was evaluated
	RuleVersion int64 `json:"rule_version" api:"required"`
	// The token of the transaction that triggered the rule evaluation
	TransactionToken string                                               `json:"transaction_token" api:"required,nullable" format:"uuid"`
	JSON             authRuleV2ListResultsResponseAuthorizationResultJSON `json:"-"`
}

// authRuleV2ListResultsResponseAuthorizationResultJSON contains the JSON metadata
// for the struct [AuthRuleV2ListResultsResponseAuthorizationResult]
type authRuleV2ListResultsResponseAuthorizationResultJSON struct {
	Token            apijson.Field
	Actions          apijson.Field
	AuthRuleToken    apijson.Field
	EvaluationTime   apijson.Field
	EventStream      apijson.Field
	EventToken       apijson.Field
	Mode             apijson.Field
	RuleVersion      apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseAuthorizationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseAuthorizationResultJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseAuthorizationResult) implementsAuthRuleV2ListResultsResponse() {}

type AuthRuleV2ListResultsResponseAuthorizationResultAction struct {
	Type AuthRuleV2ListResultsResponseAuthorizationResultActionsType `json:"type" api:"required"`
	// The detailed result code explaining the specific reason for the decline
	Code AuthRuleV2ListResultsResponseAuthorizationResultActionsCode `json:"code"`
	// Optional explanation for why this action was taken
	Explanation string                                                     `json:"explanation"`
	JSON        authRuleV2ListResultsResponseAuthorizationResultActionJSON `json:"-"`
	union       AuthRuleV2ListResultsResponseAuthorizationResultActionsUnion
}

// authRuleV2ListResultsResponseAuthorizationResultActionJSON contains the JSON
// metadata for the struct [AuthRuleV2ListResultsResponseAuthorizationResultAction]
type authRuleV2ListResultsResponseAuthorizationResultActionJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2ListResultsResponseAuthorizationResultActionJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2ListResultsResponseAuthorizationResultAction) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2ListResultsResponseAuthorizationResultAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2ListResultsResponseAuthorizationResultActionsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorization],
// [AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorization].
func (r AuthRuleV2ListResultsResponseAuthorizationResultAction) AsUnion() AuthRuleV2ListResultsResponseAuthorizationResultActionsUnion {
	return r.union
}

// Union satisfied by
// [AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorization]
// or
// [AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorization].
type AuthRuleV2ListResultsResponseAuthorizationResultActionsUnion interface {
	implementsAuthRuleV2ListResultsResponseAuthorizationResultAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResultsResponseAuthorizationResultActionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorization{}),
		},
	)
}

type AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorization struct {
	// The detailed result code explaining the specific reason for the decline
	Code AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode `json:"code" api:"required"`
	Type AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationType `json:"type" api:"required"`
	// Optional explanation for why this action was taken
	Explanation string                                                                                `json:"explanation"`
	JSON        authRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationJSON `json:"-"`
}

// authRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorization]
type authRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorization) implementsAuthRuleV2ListResultsResponseAuthorizationResultAction() {
}

// The detailed result code explaining the specific reason for the decline
type AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode string

const (
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountDailySpendLimitExceeded              AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "ACCOUNT_DAILY_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountDelinquent                           AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "ACCOUNT_DELINQUENT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountInactive                             AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "ACCOUNT_INACTIVE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountLifetimeSpendLimitExceeded           AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "ACCOUNT_LIFETIME_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountMonthlySpendLimitExceeded            AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "ACCOUNT_MONTHLY_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountPaused                               AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "ACCOUNT_PAUSED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountUnderReview                          AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "ACCOUNT_UNDER_REVIEW"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAddressIncorrect                            AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "ADDRESS_INCORRECT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeApproved                                    AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "APPROVED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRuleAllowedCountry                      AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "AUTH_RULE_ALLOWED_COUNTRY"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRuleAllowedMcc                          AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "AUTH_RULE_ALLOWED_MCC"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRuleBlockedCountry                      AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "AUTH_RULE_BLOCKED_COUNTRY"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRuleBlockedMcc                          AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "AUTH_RULE_BLOCKED_MCC"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRule                                    AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "AUTH_RULE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardClosed                                  AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_CLOSED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardCryptogramValidationFailure             AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_CRYPTOGRAM_VALIDATION_FAILURE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardExpired                                 AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_EXPIRED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardExpiryDateIncorrect                     AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_EXPIRY_DATE_INCORRECT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardInvalid                                 AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardNotActivated                            AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_NOT_ACTIVATED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardPaused                                  AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_PAUSED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardPinIncorrect                            AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_PIN_INCORRECT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardRestricted                              AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_RESTRICTED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardSecurityCodeIncorrect                   AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_SECURITY_CODE_INCORRECT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardSpendLimitExceeded                      AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARD_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeContactCardIssuer                           AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CONTACT_CARD_ISSUER"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCustomerAsaTimeout                          AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CUSTOMER_ASA_TIMEOUT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCustomAsaResult                             AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CUSTOM_ASA_RESULT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeDeclined                                    AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "DECLINED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeDoNotHonor                                  AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "DO_NOT_HONOR"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeDriverNumberInvalid                         AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "DRIVER_NUMBER_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeFormatError                                 AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "FORMAT_ERROR"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeInsufficientFundingSourceBalance            AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "INSUFFICIENT_FUNDING_SOURCE_BALANCE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeInsufficientFunds                           AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "INSUFFICIENT_FUNDS"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeLithicSystemError                           AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "LITHIC_SYSTEM_ERROR"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeLithicSystemRateLimit                       AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "LITHIC_SYSTEM_RATE_LIMIT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeMalformedAsaResponse                        AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "MALFORMED_ASA_RESPONSE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeMerchantInvalid                             AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "MERCHANT_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeMerchantLockedCardAttemptedElsewhere        AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "MERCHANT_LOCKED_CARD_ATTEMPTED_ELSEWHERE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeMerchantNotPermitted                        AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "MERCHANT_NOT_PERMITTED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeOverReversalAttempted                       AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "OVER_REVERSAL_ATTEMPTED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodePinBlocked                                  AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "PIN_BLOCKED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeProgramCardSpendLimitExceeded               AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "PROGRAM_CARD_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeProgramSuspended                            AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "PROGRAM_SUSPENDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeProgramUsageRestriction                     AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "PROGRAM_USAGE_RESTRICTION"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeReversalUnmatched                           AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "REVERSAL_UNMATCHED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeSecurityViolation                           AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "SECURITY_VIOLATION"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeSingleUseCardReattempted                    AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "SINGLE_USE_CARD_REATTEMPTED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeSuspectedFraud                              AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "SUSPECTED_FRAUD"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeTransactionInvalid                          AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "TRANSACTION_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeTransactionNotPermittedToAcquirerOrTerminal AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "TRANSACTION_NOT_PERMITTED_TO_ACQUIRER_OR_TERMINAL"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeTransactionNotPermittedToIssuerOrCardholder AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "TRANSACTION_NOT_PERMITTED_TO_ISSUER_OR_CARDHOLDER"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeTransactionPreviouslyCompleted              AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "TRANSACTION_PREVIOUSLY_COMPLETED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeUnauthorizedMerchant                        AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "UNAUTHORIZED_MERCHANT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeVehicleNumberInvalid                        AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "VEHICLE_NUMBER_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardholderChallenged                        AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARDHOLDER_CHALLENGED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardholderChallengeFailed                   AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode = "CARDHOLDER_CHALLENGE_FAILED"
)

func (r AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountDailySpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountDelinquent, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountInactive, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountLifetimeSpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountMonthlySpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountPaused, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAccountUnderReview, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAddressIncorrect, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeApproved, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRuleAllowedCountry, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRuleAllowedMcc, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRuleBlockedCountry, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRuleBlockedMcc, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeAuthRule, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardClosed, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardCryptogramValidationFailure, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardExpired, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardExpiryDateIncorrect, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardNotActivated, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardPaused, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardPinIncorrect, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardRestricted, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardSecurityCodeIncorrect, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardSpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeContactCardIssuer, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCustomerAsaTimeout, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCustomAsaResult, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeDeclined, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeDoNotHonor, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeDriverNumberInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeFormatError, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeInsufficientFundingSourceBalance, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeInsufficientFunds, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeLithicSystemError, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeLithicSystemRateLimit, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeMalformedAsaResponse, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeMerchantInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeMerchantLockedCardAttemptedElsewhere, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeMerchantNotPermitted, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeOverReversalAttempted, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodePinBlocked, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeProgramCardSpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeProgramSuspended, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeProgramUsageRestriction, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeReversalUnmatched, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeSecurityViolation, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeSingleUseCardReattempted, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeSuspectedFraud, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeTransactionInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeTransactionNotPermittedToAcquirerOrTerminal, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeTransactionNotPermittedToIssuerOrCardholder, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeTransactionPreviouslyCompleted, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeUnauthorizedMerchant, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeVehicleNumberInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardholderChallenged, AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationCodeCardholderChallengeFailed:
		return true
	}
	return false
}

type AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationType string

const (
	AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationTypeDecline AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationType = "DECLINE"
)

func (r AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthorizationResultActionsDeclineActionAuthorizationTypeDecline:
		return true
	}
	return false
}

type AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorization struct {
	Type AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationType `json:"type" api:"required"`
	// Optional explanation for why this action was taken
	Explanation string                                                                                  `json:"explanation"`
	JSON        authRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationJSON `json:"-"`
}

// authRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorization]
type authRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationJSON struct {
	Type        apijson.Field
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorization) implementsAuthRuleV2ListResultsResponseAuthorizationResultAction() {
}

type AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationType string

const (
	AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationTypeChallenge AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationType = "CHALLENGE"
)

func (r AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthorizationResultActionsChallengeActionAuthorizationTypeChallenge:
		return true
	}
	return false
}

type AuthRuleV2ListResultsResponseAuthorizationResultActionsType string

const (
	AuthRuleV2ListResultsResponseAuthorizationResultActionsTypeDecline   AuthRuleV2ListResultsResponseAuthorizationResultActionsType = "DECLINE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsTypeChallenge AuthRuleV2ListResultsResponseAuthorizationResultActionsType = "CHALLENGE"
)

func (r AuthRuleV2ListResultsResponseAuthorizationResultActionsType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthorizationResultActionsTypeDecline, AuthRuleV2ListResultsResponseAuthorizationResultActionsTypeChallenge:
		return true
	}
	return false
}

// The detailed result code explaining the specific reason for the decline
type AuthRuleV2ListResultsResponseAuthorizationResultActionsCode string

const (
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountDailySpendLimitExceeded              AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "ACCOUNT_DAILY_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountDelinquent                           AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "ACCOUNT_DELINQUENT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountInactive                             AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "ACCOUNT_INACTIVE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountLifetimeSpendLimitExceeded           AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "ACCOUNT_LIFETIME_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountMonthlySpendLimitExceeded            AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "ACCOUNT_MONTHLY_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountPaused                               AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "ACCOUNT_PAUSED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountUnderReview                          AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "ACCOUNT_UNDER_REVIEW"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAddressIncorrect                            AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "ADDRESS_INCORRECT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeApproved                                    AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "APPROVED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRuleAllowedCountry                      AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "AUTH_RULE_ALLOWED_COUNTRY"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRuleAllowedMcc                          AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "AUTH_RULE_ALLOWED_MCC"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRuleBlockedCountry                      AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "AUTH_RULE_BLOCKED_COUNTRY"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRuleBlockedMcc                          AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "AUTH_RULE_BLOCKED_MCC"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRule                                    AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "AUTH_RULE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardClosed                                  AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_CLOSED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardCryptogramValidationFailure             AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_CRYPTOGRAM_VALIDATION_FAILURE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardExpired                                 AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_EXPIRED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardExpiryDateIncorrect                     AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_EXPIRY_DATE_INCORRECT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardInvalid                                 AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardNotActivated                            AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_NOT_ACTIVATED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardPaused                                  AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_PAUSED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardPinIncorrect                            AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_PIN_INCORRECT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardRestricted                              AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_RESTRICTED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardSecurityCodeIncorrect                   AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_SECURITY_CODE_INCORRECT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardSpendLimitExceeded                      AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARD_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeContactCardIssuer                           AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CONTACT_CARD_ISSUER"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCustomerAsaTimeout                          AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CUSTOMER_ASA_TIMEOUT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCustomAsaResult                             AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CUSTOM_ASA_RESULT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeDeclined                                    AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "DECLINED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeDoNotHonor                                  AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "DO_NOT_HONOR"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeDriverNumberInvalid                         AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "DRIVER_NUMBER_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeFormatError                                 AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "FORMAT_ERROR"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeInsufficientFundingSourceBalance            AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "INSUFFICIENT_FUNDING_SOURCE_BALANCE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeInsufficientFunds                           AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "INSUFFICIENT_FUNDS"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeLithicSystemError                           AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "LITHIC_SYSTEM_ERROR"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeLithicSystemRateLimit                       AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "LITHIC_SYSTEM_RATE_LIMIT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeMalformedAsaResponse                        AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "MALFORMED_ASA_RESPONSE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeMerchantInvalid                             AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "MERCHANT_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeMerchantLockedCardAttemptedElsewhere        AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "MERCHANT_LOCKED_CARD_ATTEMPTED_ELSEWHERE"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeMerchantNotPermitted                        AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "MERCHANT_NOT_PERMITTED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeOverReversalAttempted                       AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "OVER_REVERSAL_ATTEMPTED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodePinBlocked                                  AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "PIN_BLOCKED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeProgramCardSpendLimitExceeded               AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "PROGRAM_CARD_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeProgramSuspended                            AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "PROGRAM_SUSPENDED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeProgramUsageRestriction                     AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "PROGRAM_USAGE_RESTRICTION"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeReversalUnmatched                           AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "REVERSAL_UNMATCHED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeSecurityViolation                           AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "SECURITY_VIOLATION"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeSingleUseCardReattempted                    AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "SINGLE_USE_CARD_REATTEMPTED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeSuspectedFraud                              AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "SUSPECTED_FRAUD"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeTransactionInvalid                          AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "TRANSACTION_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeTransactionNotPermittedToAcquirerOrTerminal AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "TRANSACTION_NOT_PERMITTED_TO_ACQUIRER_OR_TERMINAL"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeTransactionNotPermittedToIssuerOrCardholder AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "TRANSACTION_NOT_PERMITTED_TO_ISSUER_OR_CARDHOLDER"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeTransactionPreviouslyCompleted              AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "TRANSACTION_PREVIOUSLY_COMPLETED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeUnauthorizedMerchant                        AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "UNAUTHORIZED_MERCHANT"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeVehicleNumberInvalid                        AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "VEHICLE_NUMBER_INVALID"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardholderChallenged                        AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARDHOLDER_CHALLENGED"
	AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardholderChallengeFailed                   AuthRuleV2ListResultsResponseAuthorizationResultActionsCode = "CARDHOLDER_CHALLENGE_FAILED"
)

func (r AuthRuleV2ListResultsResponseAuthorizationResultActionsCode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountDailySpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountDelinquent, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountInactive, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountLifetimeSpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountMonthlySpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountPaused, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAccountUnderReview, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAddressIncorrect, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeApproved, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRuleAllowedCountry, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRuleAllowedMcc, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRuleBlockedCountry, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRuleBlockedMcc, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeAuthRule, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardClosed, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardCryptogramValidationFailure, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardExpired, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardExpiryDateIncorrect, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardNotActivated, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardPaused, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardPinIncorrect, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardRestricted, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardSecurityCodeIncorrect, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardSpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeContactCardIssuer, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCustomerAsaTimeout, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCustomAsaResult, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeDeclined, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeDoNotHonor, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeDriverNumberInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeFormatError, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeInsufficientFundingSourceBalance, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeInsufficientFunds, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeLithicSystemError, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeLithicSystemRateLimit, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeMalformedAsaResponse, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeMerchantInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeMerchantLockedCardAttemptedElsewhere, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeMerchantNotPermitted, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeOverReversalAttempted, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodePinBlocked, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeProgramCardSpendLimitExceeded, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeProgramSuspended, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeProgramUsageRestriction, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeReversalUnmatched, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeSecurityViolation, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeSingleUseCardReattempted, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeSuspectedFraud, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeTransactionInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeTransactionNotPermittedToAcquirerOrTerminal, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeTransactionNotPermittedToIssuerOrCardholder, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeTransactionPreviouslyCompleted, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeUnauthorizedMerchant, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeVehicleNumberInvalid, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardholderChallenged, AuthRuleV2ListResultsResponseAuthorizationResultActionsCodeCardholderChallengeFailed:
		return true
	}
	return false
}

// The event stream during which the rule was evaluated
type AuthRuleV2ListResultsResponseAuthorizationResultEventStream string

const (
	AuthRuleV2ListResultsResponseAuthorizationResultEventStreamAuthorization AuthRuleV2ListResultsResponseAuthorizationResultEventStream = "AUTHORIZATION"
)

func (r AuthRuleV2ListResultsResponseAuthorizationResultEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthorizationResultEventStreamAuthorization:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2ListResultsResponseAuthorizationResultMode string

const (
	AuthRuleV2ListResultsResponseAuthorizationResultModeActive   AuthRuleV2ListResultsResponseAuthorizationResultMode = "ACTIVE"
	AuthRuleV2ListResultsResponseAuthorizationResultModeInactive AuthRuleV2ListResultsResponseAuthorizationResultMode = "INACTIVE"
)

func (r AuthRuleV2ListResultsResponseAuthorizationResultMode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthorizationResultModeActive, AuthRuleV2ListResultsResponseAuthorizationResultModeInactive:
		return true
	}
	return false
}

type AuthRuleV2ListResultsResponseAuthentication3DSResult struct {
	// Globally unique identifier for the evaluation
	Token string `json:"token" api:"required" format:"uuid"`
	// Actions returned by the rule evaluation
	Actions []AuthRuleV2ListResultsResponseAuthentication3DsResultAction `json:"actions" api:"required"`
	// The Auth Rule token
	AuthRuleToken string `json:"auth_rule_token" api:"required" format:"uuid"`
	// Timestamp of the rule evaluation
	EvaluationTime time.Time `json:"evaluation_time" api:"required" format:"date-time"`
	// The event stream during which the rule was evaluated
	EventStream AuthRuleV2ListResultsResponseAuthentication3DSResultEventStream `json:"event_stream" api:"required"`
	// Token of the event that triggered the evaluation
	EventToken string `json:"event_token" api:"required" format:"uuid"`
	// The state of the Auth Rule
	Mode AuthRuleV2ListResultsResponseAuthentication3DSResultMode `json:"mode" api:"required"`
	// Version of the rule that was evaluated
	RuleVersion int64 `json:"rule_version" api:"required"`
	// The token of the transaction that triggered the rule evaluation
	TransactionToken string                                                   `json:"transaction_token" api:"required,nullable" format:"uuid"`
	JSON             authRuleV2ListResultsResponseAuthentication3DsResultJSON `json:"-"`
}

// authRuleV2ListResultsResponseAuthentication3DsResultJSON contains the JSON
// metadata for the struct [AuthRuleV2ListResultsResponseAuthentication3DSResult]
type authRuleV2ListResultsResponseAuthentication3DsResultJSON struct {
	Token            apijson.Field
	Actions          apijson.Field
	AuthRuleToken    apijson.Field
	EvaluationTime   apijson.Field
	EventStream      apijson.Field
	EventToken       apijson.Field
	Mode             apijson.Field
	RuleVersion      apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseAuthentication3DSResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseAuthentication3DsResultJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseAuthentication3DSResult) implementsAuthRuleV2ListResultsResponse() {
}

type AuthRuleV2ListResultsResponseAuthentication3DsResultAction struct {
	Type AuthRuleV2ListResultsResponseAuthentication3DSResultActionsType `json:"type" api:"required"`
	// Optional explanation for why this action was taken
	Explanation string                                                         `json:"explanation"`
	JSON        authRuleV2ListResultsResponseAuthentication3DsResultActionJSON `json:"-"`
}

// authRuleV2ListResultsResponseAuthentication3DsResultActionJSON contains the JSON
// metadata for the struct
// [AuthRuleV2ListResultsResponseAuthentication3DsResultAction]
type authRuleV2ListResultsResponseAuthentication3DsResultActionJSON struct {
	Type        apijson.Field
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseAuthentication3DsResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseAuthentication3DsResultActionJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ListResultsResponseAuthentication3DSResultActionsType string

const (
	AuthRuleV2ListResultsResponseAuthentication3DSResultActionsTypeDecline   AuthRuleV2ListResultsResponseAuthentication3DSResultActionsType = "DECLINE"
	AuthRuleV2ListResultsResponseAuthentication3DSResultActionsTypeChallenge AuthRuleV2ListResultsResponseAuthentication3DSResultActionsType = "CHALLENGE"
)

func (r AuthRuleV2ListResultsResponseAuthentication3DSResultActionsType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthentication3DSResultActionsTypeDecline, AuthRuleV2ListResultsResponseAuthentication3DSResultActionsTypeChallenge:
		return true
	}
	return false
}

// The event stream during which the rule was evaluated
type AuthRuleV2ListResultsResponseAuthentication3DSResultEventStream string

const (
	AuthRuleV2ListResultsResponseAuthentication3DSResultEventStreamThreeDSAuthentication AuthRuleV2ListResultsResponseAuthentication3DSResultEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2ListResultsResponseAuthentication3DSResultEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthentication3DSResultEventStreamThreeDSAuthentication:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2ListResultsResponseAuthentication3DSResultMode string

const (
	AuthRuleV2ListResultsResponseAuthentication3DSResultModeActive   AuthRuleV2ListResultsResponseAuthentication3DSResultMode = "ACTIVE"
	AuthRuleV2ListResultsResponseAuthentication3DSResultModeInactive AuthRuleV2ListResultsResponseAuthentication3DSResultMode = "INACTIVE"
)

func (r AuthRuleV2ListResultsResponseAuthentication3DSResultMode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseAuthentication3DSResultModeActive, AuthRuleV2ListResultsResponseAuthentication3DSResultModeInactive:
		return true
	}
	return false
}

type AuthRuleV2ListResultsResponseTokenizationResult struct {
	// Globally unique identifier for the evaluation
	Token string `json:"token" api:"required" format:"uuid"`
	// Actions returned by the rule evaluation
	Actions []AuthRuleV2ListResultsResponseTokenizationResultAction `json:"actions" api:"required"`
	// The Auth Rule token
	AuthRuleToken string `json:"auth_rule_token" api:"required" format:"uuid"`
	// Timestamp of the rule evaluation
	EvaluationTime time.Time `json:"evaluation_time" api:"required" format:"date-time"`
	// The event stream during which the rule was evaluated
	EventStream AuthRuleV2ListResultsResponseTokenizationResultEventStream `json:"event_stream" api:"required"`
	// Token of the event that triggered the evaluation
	EventToken string `json:"event_token" api:"required" format:"uuid"`
	// The state of the Auth Rule
	Mode AuthRuleV2ListResultsResponseTokenizationResultMode `json:"mode" api:"required"`
	// Version of the rule that was evaluated
	RuleVersion int64 `json:"rule_version" api:"required"`
	// The token of the transaction that triggered the rule evaluation
	TransactionToken string                                              `json:"transaction_token" api:"required,nullable" format:"uuid"`
	JSON             authRuleV2ListResultsResponseTokenizationResultJSON `json:"-"`
}

// authRuleV2ListResultsResponseTokenizationResultJSON contains the JSON metadata
// for the struct [AuthRuleV2ListResultsResponseTokenizationResult]
type authRuleV2ListResultsResponseTokenizationResultJSON struct {
	Token            apijson.Field
	Actions          apijson.Field
	AuthRuleToken    apijson.Field
	EvaluationTime   apijson.Field
	EventStream      apijson.Field
	EventToken       apijson.Field
	Mode             apijson.Field
	RuleVersion      apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseTokenizationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseTokenizationResultJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseTokenizationResult) implementsAuthRuleV2ListResultsResponse() {}

type AuthRuleV2ListResultsResponseTokenizationResultAction struct {
	// Decline the tokenization request
	Type AuthRuleV2ListResultsResponseTokenizationResultActionsType `json:"type" api:"required"`
	// Optional explanation for why this action was taken
	Explanation string `json:"explanation"`
	// Reason code for declining the tokenization request
	Reason AuthRuleV2ListResultsResponseTokenizationResultActionsReason `json:"reason"`
	JSON   authRuleV2ListResultsResponseTokenizationResultActionJSON    `json:"-"`
	union  AuthRuleV2ListResultsResponseTokenizationResultActionsUnion
}

// authRuleV2ListResultsResponseTokenizationResultActionJSON contains the JSON
// metadata for the struct [AuthRuleV2ListResultsResponseTokenizationResultAction]
type authRuleV2ListResultsResponseTokenizationResultActionJSON struct {
	Type        apijson.Field
	Explanation apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2ListResultsResponseTokenizationResultActionJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2ListResultsResponseTokenizationResultAction) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2ListResultsResponseTokenizationResultAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2ListResultsResponseTokenizationResultActionsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenization],
// [AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaAction].
func (r AuthRuleV2ListResultsResponseTokenizationResultAction) AsUnion() AuthRuleV2ListResultsResponseTokenizationResultActionsUnion {
	return r.union
}

// Union satisfied by
// [AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenization]
// or [AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaAction].
type AuthRuleV2ListResultsResponseTokenizationResultActionsUnion interface {
	implementsAuthRuleV2ListResultsResponseTokenizationResultAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResultsResponseTokenizationResultActionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaAction{}),
		},
	)
}

type AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenization struct {
	// Decline the tokenization request
	Type AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationType `json:"type" api:"required"`
	// Optional explanation for why this action was taken
	Explanation string `json:"explanation"`
	// Reason code for declining the tokenization request
	Reason AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason `json:"reason"`
	JSON   authRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationJSON   `json:"-"`
}

// authRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenization]
type authRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationJSON struct {
	Type        apijson.Field
	Explanation apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenization) implementsAuthRuleV2ListResultsResponseTokenizationResultAction() {
}

// Decline the tokenization request
type AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationType string

const (
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationTypeDecline AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationType = "DECLINE"
)

func (r AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationTypeDecline:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason string

const (
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonAccountScore1                  AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "ACCOUNT_SCORE_1"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonDeviceScore1                   AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "DEVICE_SCORE_1"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonAllWalletDeclineReasonsPresent AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonWalletRecommendedDecisionRed   AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "WALLET_RECOMMENDED_DECISION_RED"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCvcMismatch                    AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "CVC_MISMATCH"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCardExpiryMonthMismatch        AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "CARD_EXPIRY_MONTH_MISMATCH"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCardExpiryYearMismatch         AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "CARD_EXPIRY_YEAR_MISMATCH"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCardInvalidState               AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "CARD_INVALID_STATE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCustomerRedPath                AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "CUSTOMER_RED_PATH"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonInvalidCustomerResponse        AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "INVALID_CUSTOMER_RESPONSE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonNetworkFailure                 AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "NETWORK_FAILURE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonGenericDecline                 AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "GENERIC_DECLINE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonDigitalCardArtRequired         AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason = "DIGITAL_CARD_ART_REQUIRED"
)

func (r AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReason) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonAccountScore1, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonDeviceScore1, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonAllWalletDeclineReasonsPresent, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonWalletRecommendedDecisionRed, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCvcMismatch, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCardExpiryMonthMismatch, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCardExpiryYearMismatch, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCardInvalidState, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonCustomerRedPath, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonInvalidCustomerResponse, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonNetworkFailure, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonGenericDecline, AuthRuleV2ListResultsResponseTokenizationResultActionsDeclineActionTokenizationReasonDigitalCardArtRequired:
		return true
	}
	return false
}

type AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaAction struct {
	// Require two-factor authentication for the tokenization request
	Type AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionType `json:"type" api:"required"`
	// Optional explanation for why this action was taken
	Explanation string `json:"explanation"`
	// Reason code for requiring two-factor authentication
	Reason AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason `json:"reason"`
	JSON   authRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionJSON   `json:"-"`
}

// authRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaAction]
type authRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionJSON struct {
	Type        apijson.Field
	Explanation apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaAction) implementsAuthRuleV2ListResultsResponseTokenizationResultAction() {
}

// Require two-factor authentication for the tokenization request
type AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionType string

const (
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionTypeRequireTfa AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionType = "REQUIRE_TFA"
)

func (r AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionTypeRequireTfa:
		return true
	}
	return false
}

// Reason code for requiring two-factor authentication
type AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason string

const (
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonWalletRecommendedTfa        AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "WALLET_RECOMMENDED_TFA"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonSuspiciousActivity          AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "SUSPICIOUS_ACTIVITY"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonDeviceRecentlyLost          AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "DEVICE_RECENTLY_LOST"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonTooManyRecentAttempts       AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "TOO_MANY_RECENT_ATTEMPTS"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonTooManyRecentTokens         AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "TOO_MANY_RECENT_TOKENS"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonTooManyDifferentCardholders AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonOutsideHomeTerritory        AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "OUTSIDE_HOME_TERRITORY"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonHasSuspendedTokens          AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "HAS_SUSPENDED_TOKENS"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonHighRisk                    AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "HIGH_RISK"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonAccountScoreLow             AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "ACCOUNT_SCORE_LOW"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonDeviceScoreLow              AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "DEVICE_SCORE_LOW"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonCardStateTfa                AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "CARD_STATE_TFA"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonHardcodedTfa                AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "HARDCODED_TFA"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonCustomerRuleTfa             AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "CUSTOMER_RULE_TFA"
	AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonDeviceHostCardEmulation     AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReason) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonWalletRecommendedTfa, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonSuspiciousActivity, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonDeviceRecentlyLost, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonTooManyRecentAttempts, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonTooManyRecentTokens, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonTooManyDifferentCardholders, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonOutsideHomeTerritory, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonHasSuspendedTokens, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonHighRisk, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonAccountScoreLow, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonDeviceScoreLow, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonCardStateTfa, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonHardcodedTfa, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonCustomerRuleTfa, AuthRuleV2ListResultsResponseTokenizationResultActionsRequireTfaActionReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

// Decline the tokenization request
type AuthRuleV2ListResultsResponseTokenizationResultActionsType string

const (
	AuthRuleV2ListResultsResponseTokenizationResultActionsTypeDecline    AuthRuleV2ListResultsResponseTokenizationResultActionsType = "DECLINE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsTypeRequireTfa AuthRuleV2ListResultsResponseTokenizationResultActionsType = "REQUIRE_TFA"
)

func (r AuthRuleV2ListResultsResponseTokenizationResultActionsType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseTokenizationResultActionsTypeDecline, AuthRuleV2ListResultsResponseTokenizationResultActionsTypeRequireTfa:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type AuthRuleV2ListResultsResponseTokenizationResultActionsReason string

const (
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonAccountScore1                  AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "ACCOUNT_SCORE_1"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDeviceScore1                   AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "DEVICE_SCORE_1"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonAllWalletDeclineReasonsPresent AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonWalletRecommendedDecisionRed   AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "WALLET_RECOMMENDED_DECISION_RED"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCvcMismatch                    AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "CVC_MISMATCH"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCardExpiryMonthMismatch        AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "CARD_EXPIRY_MONTH_MISMATCH"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCardExpiryYearMismatch         AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "CARD_EXPIRY_YEAR_MISMATCH"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCardInvalidState               AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "CARD_INVALID_STATE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCustomerRedPath                AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "CUSTOMER_RED_PATH"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonInvalidCustomerResponse        AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "INVALID_CUSTOMER_RESPONSE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonNetworkFailure                 AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "NETWORK_FAILURE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonGenericDecline                 AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "GENERIC_DECLINE"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDigitalCardArtRequired         AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "DIGITAL_CARD_ART_REQUIRED"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonWalletRecommendedTfa           AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "WALLET_RECOMMENDED_TFA"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonSuspiciousActivity             AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "SUSPICIOUS_ACTIVITY"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDeviceRecentlyLost             AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "DEVICE_RECENTLY_LOST"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonTooManyRecentAttempts          AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "TOO_MANY_RECENT_ATTEMPTS"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonTooManyRecentTokens            AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "TOO_MANY_RECENT_TOKENS"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonTooManyDifferentCardholders    AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonOutsideHomeTerritory           AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "OUTSIDE_HOME_TERRITORY"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonHasSuspendedTokens             AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "HAS_SUSPENDED_TOKENS"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonHighRisk                       AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "HIGH_RISK"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonAccountScoreLow                AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "ACCOUNT_SCORE_LOW"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDeviceScoreLow                 AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "DEVICE_SCORE_LOW"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCardStateTfa                   AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "CARD_STATE_TFA"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonHardcodedTfa                   AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "HARDCODED_TFA"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCustomerRuleTfa                AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "CUSTOMER_RULE_TFA"
	AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDeviceHostCardEmulation        AuthRuleV2ListResultsResponseTokenizationResultActionsReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r AuthRuleV2ListResultsResponseTokenizationResultActionsReason) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseTokenizationResultActionsReasonAccountScore1, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDeviceScore1, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonAllWalletDeclineReasonsPresent, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonWalletRecommendedDecisionRed, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCvcMismatch, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCardExpiryMonthMismatch, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCardExpiryYearMismatch, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCardInvalidState, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCustomerRedPath, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonInvalidCustomerResponse, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonNetworkFailure, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonGenericDecline, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDigitalCardArtRequired, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonWalletRecommendedTfa, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonSuspiciousActivity, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDeviceRecentlyLost, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonTooManyRecentAttempts, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonTooManyRecentTokens, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonTooManyDifferentCardholders, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonOutsideHomeTerritory, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonHasSuspendedTokens, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonHighRisk, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonAccountScoreLow, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDeviceScoreLow, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCardStateTfa, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonHardcodedTfa, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonCustomerRuleTfa, AuthRuleV2ListResultsResponseTokenizationResultActionsReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

// The event stream during which the rule was evaluated
type AuthRuleV2ListResultsResponseTokenizationResultEventStream string

const (
	AuthRuleV2ListResultsResponseTokenizationResultEventStreamTokenization AuthRuleV2ListResultsResponseTokenizationResultEventStream = "TOKENIZATION"
)

func (r AuthRuleV2ListResultsResponseTokenizationResultEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseTokenizationResultEventStreamTokenization:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2ListResultsResponseTokenizationResultMode string

const (
	AuthRuleV2ListResultsResponseTokenizationResultModeActive   AuthRuleV2ListResultsResponseTokenizationResultMode = "ACTIVE"
	AuthRuleV2ListResultsResponseTokenizationResultModeInactive AuthRuleV2ListResultsResponseTokenizationResultMode = "INACTIVE"
)

func (r AuthRuleV2ListResultsResponseTokenizationResultMode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseTokenizationResultModeActive, AuthRuleV2ListResultsResponseTokenizationResultModeInactive:
		return true
	}
	return false
}

type AuthRuleV2ListResultsResponseACHResult struct {
	// Globally unique identifier for the evaluation
	Token string `json:"token" api:"required" format:"uuid"`
	// Actions returned by the rule evaluation
	Actions []AuthRuleV2ListResultsResponseACHResultAction `json:"actions" api:"required"`
	// The Auth Rule token
	AuthRuleToken string `json:"auth_rule_token" api:"required" format:"uuid"`
	// Timestamp of the rule evaluation
	EvaluationTime time.Time `json:"evaluation_time" api:"required" format:"date-time"`
	// The event stream during which the rule was evaluated
	EventStream AuthRuleV2ListResultsResponseACHResultEventStream `json:"event_stream" api:"required"`
	// Token of the event that triggered the evaluation
	EventToken string `json:"event_token" api:"required" format:"uuid"`
	// The state of the Auth Rule
	Mode AuthRuleV2ListResultsResponseACHResultMode `json:"mode" api:"required"`
	// Version of the rule that was evaluated
	RuleVersion int64 `json:"rule_version" api:"required"`
	// The token of the transaction that triggered the rule evaluation
	TransactionToken string                                     `json:"transaction_token" api:"required,nullable" format:"uuid"`
	JSON             authRuleV2ListResultsResponseACHResultJSON `json:"-"`
}

// authRuleV2ListResultsResponseACHResultJSON contains the JSON metadata for the
// struct [AuthRuleV2ListResultsResponseACHResult]
type authRuleV2ListResultsResponseACHResultJSON struct {
	Token            apijson.Field
	Actions          apijson.Field
	AuthRuleToken    apijson.Field
	EvaluationTime   apijson.Field
	EventStream      apijson.Field
	EventToken       apijson.Field
	Mode             apijson.Field
	RuleVersion      apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseACHResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseACHResultJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseACHResult) implementsAuthRuleV2ListResultsResponse() {}

type AuthRuleV2ListResultsResponseACHResultAction struct {
	// Approve the ACH transaction
	Type AuthRuleV2ListResultsResponseACHResultActionsType `json:"type" api:"required"`
	// NACHA return code to use when returning the transaction. Note that the list of
	// available return codes is subject to an allowlist configured at the program
	// level
	Code AuthRuleV2ListResultsResponseACHResultActionsCode `json:"code"`
	// Optional explanation for why this action was taken
	Explanation string                                           `json:"explanation"`
	JSON        authRuleV2ListResultsResponseACHResultActionJSON `json:"-"`
	union       AuthRuleV2ListResultsResponseACHResultActionsUnion
}

// authRuleV2ListResultsResponseACHResultActionJSON contains the JSON metadata for
// the struct [AuthRuleV2ListResultsResponseACHResultAction]
type authRuleV2ListResultsResponseACHResultActionJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2ListResultsResponseACHResultActionJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2ListResultsResponseACHResultAction) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2ListResultsResponseACHResultAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2ListResultsResponseACHResultActionsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2ListResultsResponseACHResultActionsApproveActionACH],
// [AuthRuleV2ListResultsResponseACHResultActionsReturnAction].
func (r AuthRuleV2ListResultsResponseACHResultAction) AsUnion() AuthRuleV2ListResultsResponseACHResultActionsUnion {
	return r.union
}

// Union satisfied by
// [AuthRuleV2ListResultsResponseACHResultActionsApproveActionACH] or
// [AuthRuleV2ListResultsResponseACHResultActionsReturnAction].
type AuthRuleV2ListResultsResponseACHResultActionsUnion interface {
	implementsAuthRuleV2ListResultsResponseACHResultAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResultsResponseACHResultActionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseACHResultActionsApproveActionACH{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResultsResponseACHResultActionsReturnAction{}),
		},
	)
}

type AuthRuleV2ListResultsResponseACHResultActionsApproveActionACH struct {
	// Approve the ACH transaction
	Type AuthRuleV2ListResultsResponseACHResultActionsApproveActionACHType `json:"type" api:"required"`
	// Optional explanation for why this action was taken
	Explanation string                                                            `json:"explanation"`
	JSON        authRuleV2ListResultsResponseACHResultActionsApproveActionACHJSON `json:"-"`
}

// authRuleV2ListResultsResponseACHResultActionsApproveActionACHJSON contains the
// JSON metadata for the struct
// [AuthRuleV2ListResultsResponseACHResultActionsApproveActionACH]
type authRuleV2ListResultsResponseACHResultActionsApproveActionACHJSON struct {
	Type        apijson.Field
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseACHResultActionsApproveActionACH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseACHResultActionsApproveActionACHJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseACHResultActionsApproveActionACH) implementsAuthRuleV2ListResultsResponseACHResultAction() {
}

// Approve the ACH transaction
type AuthRuleV2ListResultsResponseACHResultActionsApproveActionACHType string

const (
	AuthRuleV2ListResultsResponseACHResultActionsApproveActionACHTypeApprove AuthRuleV2ListResultsResponseACHResultActionsApproveActionACHType = "APPROVE"
)

func (r AuthRuleV2ListResultsResponseACHResultActionsApproveActionACHType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseACHResultActionsApproveActionACHTypeApprove:
		return true
	}
	return false
}

type AuthRuleV2ListResultsResponseACHResultActionsReturnAction struct {
	// NACHA return code to use when returning the transaction. Note that the list of
	// available return codes is subject to an allowlist configured at the program
	// level
	Code AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode `json:"code" api:"required"`
	// Return the ACH transaction
	Type AuthRuleV2ListResultsResponseACHResultActionsReturnActionType `json:"type" api:"required"`
	// Optional explanation for why this action was taken
	Explanation string                                                        `json:"explanation"`
	JSON        authRuleV2ListResultsResponseACHResultActionsReturnActionJSON `json:"-"`
}

// authRuleV2ListResultsResponseACHResultActionsReturnActionJSON contains the JSON
// metadata for the struct
// [AuthRuleV2ListResultsResponseACHResultActionsReturnAction]
type authRuleV2ListResultsResponseACHResultActionsReturnActionJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResultsResponseACHResultActionsReturnAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResultsResponseACHResultActionsReturnActionJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResultsResponseACHResultActionsReturnAction) implementsAuthRuleV2ListResultsResponseACHResultAction() {
}

// NACHA return code to use when returning the transaction. Note that the list of
// available return codes is subject to an allowlist configured at the program
// level
type AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode string

const (
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR01 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R01"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR02 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R02"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR03 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R03"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR04 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R04"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR05 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R05"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR06 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R06"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR07 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R07"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR08 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R08"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR09 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R09"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR10 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R10"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR11 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R11"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR12 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R12"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR13 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R13"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR14 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R14"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR15 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R15"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR16 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R16"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR17 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R17"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR18 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R18"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR19 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R19"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR20 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R20"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR21 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R21"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR22 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R22"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR23 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R23"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR24 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R24"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR25 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R25"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR26 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R26"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR27 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R27"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR28 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R28"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR29 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R29"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR30 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R30"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR31 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R31"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR32 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R32"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR33 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R33"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR34 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R34"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR35 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R35"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR36 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R36"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR37 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R37"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR38 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R38"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR39 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R39"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR40 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R40"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR41 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R41"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR42 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R42"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR43 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R43"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR44 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R44"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR45 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R45"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR46 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R46"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR47 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R47"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR50 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R50"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR51 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R51"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR52 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R52"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR53 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R53"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR61 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R61"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR62 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R62"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR67 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R67"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR68 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R68"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR69 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R69"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR70 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R70"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR71 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R71"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR72 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R72"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR73 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R73"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR74 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R74"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR75 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R75"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR76 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R76"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR77 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R77"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR80 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R80"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR81 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R81"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR82 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R82"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR83 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R83"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR84 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R84"
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR85 AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode = "R85"
)

func (r AuthRuleV2ListResultsResponseACHResultActionsReturnActionCode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR01, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR02, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR03, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR04, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR05, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR06, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR07, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR08, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR09, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR10, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR11, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR12, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR13, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR14, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR15, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR16, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR17, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR18, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR19, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR20, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR21, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR22, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR23, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR24, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR25, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR26, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR27, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR28, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR29, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR30, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR31, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR32, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR33, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR34, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR35, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR36, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR37, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR38, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR39, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR40, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR41, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR42, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR43, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR44, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR45, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR46, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR47, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR50, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR51, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR52, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR53, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR61, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR62, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR67, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR68, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR69, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR70, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR71, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR72, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR73, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR74, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR75, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR76, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR77, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR80, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR81, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR82, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR83, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR84, AuthRuleV2ListResultsResponseACHResultActionsReturnActionCodeR85:
		return true
	}
	return false
}

// Return the ACH transaction
type AuthRuleV2ListResultsResponseACHResultActionsReturnActionType string

const (
	AuthRuleV2ListResultsResponseACHResultActionsReturnActionTypeReturn AuthRuleV2ListResultsResponseACHResultActionsReturnActionType = "RETURN"
)

func (r AuthRuleV2ListResultsResponseACHResultActionsReturnActionType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseACHResultActionsReturnActionTypeReturn:
		return true
	}
	return false
}

// Approve the ACH transaction
type AuthRuleV2ListResultsResponseACHResultActionsType string

const (
	AuthRuleV2ListResultsResponseACHResultActionsTypeApprove AuthRuleV2ListResultsResponseACHResultActionsType = "APPROVE"
	AuthRuleV2ListResultsResponseACHResultActionsTypeReturn  AuthRuleV2ListResultsResponseACHResultActionsType = "RETURN"
)

func (r AuthRuleV2ListResultsResponseACHResultActionsType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseACHResultActionsTypeApprove, AuthRuleV2ListResultsResponseACHResultActionsTypeReturn:
		return true
	}
	return false
}

// NACHA return code to use when returning the transaction. Note that the list of
// available return codes is subject to an allowlist configured at the program
// level
type AuthRuleV2ListResultsResponseACHResultActionsCode string

const (
	AuthRuleV2ListResultsResponseACHResultActionsCodeR01 AuthRuleV2ListResultsResponseACHResultActionsCode = "R01"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR02 AuthRuleV2ListResultsResponseACHResultActionsCode = "R02"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR03 AuthRuleV2ListResultsResponseACHResultActionsCode = "R03"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR04 AuthRuleV2ListResultsResponseACHResultActionsCode = "R04"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR05 AuthRuleV2ListResultsResponseACHResultActionsCode = "R05"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR06 AuthRuleV2ListResultsResponseACHResultActionsCode = "R06"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR07 AuthRuleV2ListResultsResponseACHResultActionsCode = "R07"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR08 AuthRuleV2ListResultsResponseACHResultActionsCode = "R08"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR09 AuthRuleV2ListResultsResponseACHResultActionsCode = "R09"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR10 AuthRuleV2ListResultsResponseACHResultActionsCode = "R10"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR11 AuthRuleV2ListResultsResponseACHResultActionsCode = "R11"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR12 AuthRuleV2ListResultsResponseACHResultActionsCode = "R12"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR13 AuthRuleV2ListResultsResponseACHResultActionsCode = "R13"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR14 AuthRuleV2ListResultsResponseACHResultActionsCode = "R14"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR15 AuthRuleV2ListResultsResponseACHResultActionsCode = "R15"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR16 AuthRuleV2ListResultsResponseACHResultActionsCode = "R16"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR17 AuthRuleV2ListResultsResponseACHResultActionsCode = "R17"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR18 AuthRuleV2ListResultsResponseACHResultActionsCode = "R18"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR19 AuthRuleV2ListResultsResponseACHResultActionsCode = "R19"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR20 AuthRuleV2ListResultsResponseACHResultActionsCode = "R20"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR21 AuthRuleV2ListResultsResponseACHResultActionsCode = "R21"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR22 AuthRuleV2ListResultsResponseACHResultActionsCode = "R22"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR23 AuthRuleV2ListResultsResponseACHResultActionsCode = "R23"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR24 AuthRuleV2ListResultsResponseACHResultActionsCode = "R24"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR25 AuthRuleV2ListResultsResponseACHResultActionsCode = "R25"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR26 AuthRuleV2ListResultsResponseACHResultActionsCode = "R26"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR27 AuthRuleV2ListResultsResponseACHResultActionsCode = "R27"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR28 AuthRuleV2ListResultsResponseACHResultActionsCode = "R28"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR29 AuthRuleV2ListResultsResponseACHResultActionsCode = "R29"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR30 AuthRuleV2ListResultsResponseACHResultActionsCode = "R30"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR31 AuthRuleV2ListResultsResponseACHResultActionsCode = "R31"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR32 AuthRuleV2ListResultsResponseACHResultActionsCode = "R32"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR33 AuthRuleV2ListResultsResponseACHResultActionsCode = "R33"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR34 AuthRuleV2ListResultsResponseACHResultActionsCode = "R34"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR35 AuthRuleV2ListResultsResponseACHResultActionsCode = "R35"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR36 AuthRuleV2ListResultsResponseACHResultActionsCode = "R36"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR37 AuthRuleV2ListResultsResponseACHResultActionsCode = "R37"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR38 AuthRuleV2ListResultsResponseACHResultActionsCode = "R38"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR39 AuthRuleV2ListResultsResponseACHResultActionsCode = "R39"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR40 AuthRuleV2ListResultsResponseACHResultActionsCode = "R40"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR41 AuthRuleV2ListResultsResponseACHResultActionsCode = "R41"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR42 AuthRuleV2ListResultsResponseACHResultActionsCode = "R42"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR43 AuthRuleV2ListResultsResponseACHResultActionsCode = "R43"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR44 AuthRuleV2ListResultsResponseACHResultActionsCode = "R44"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR45 AuthRuleV2ListResultsResponseACHResultActionsCode = "R45"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR46 AuthRuleV2ListResultsResponseACHResultActionsCode = "R46"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR47 AuthRuleV2ListResultsResponseACHResultActionsCode = "R47"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR50 AuthRuleV2ListResultsResponseACHResultActionsCode = "R50"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR51 AuthRuleV2ListResultsResponseACHResultActionsCode = "R51"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR52 AuthRuleV2ListResultsResponseACHResultActionsCode = "R52"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR53 AuthRuleV2ListResultsResponseACHResultActionsCode = "R53"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR61 AuthRuleV2ListResultsResponseACHResultActionsCode = "R61"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR62 AuthRuleV2ListResultsResponseACHResultActionsCode = "R62"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR67 AuthRuleV2ListResultsResponseACHResultActionsCode = "R67"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR68 AuthRuleV2ListResultsResponseACHResultActionsCode = "R68"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR69 AuthRuleV2ListResultsResponseACHResultActionsCode = "R69"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR70 AuthRuleV2ListResultsResponseACHResultActionsCode = "R70"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR71 AuthRuleV2ListResultsResponseACHResultActionsCode = "R71"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR72 AuthRuleV2ListResultsResponseACHResultActionsCode = "R72"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR73 AuthRuleV2ListResultsResponseACHResultActionsCode = "R73"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR74 AuthRuleV2ListResultsResponseACHResultActionsCode = "R74"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR75 AuthRuleV2ListResultsResponseACHResultActionsCode = "R75"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR76 AuthRuleV2ListResultsResponseACHResultActionsCode = "R76"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR77 AuthRuleV2ListResultsResponseACHResultActionsCode = "R77"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR80 AuthRuleV2ListResultsResponseACHResultActionsCode = "R80"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR81 AuthRuleV2ListResultsResponseACHResultActionsCode = "R81"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR82 AuthRuleV2ListResultsResponseACHResultActionsCode = "R82"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR83 AuthRuleV2ListResultsResponseACHResultActionsCode = "R83"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR84 AuthRuleV2ListResultsResponseACHResultActionsCode = "R84"
	AuthRuleV2ListResultsResponseACHResultActionsCodeR85 AuthRuleV2ListResultsResponseACHResultActionsCode = "R85"
)

func (r AuthRuleV2ListResultsResponseACHResultActionsCode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseACHResultActionsCodeR01, AuthRuleV2ListResultsResponseACHResultActionsCodeR02, AuthRuleV2ListResultsResponseACHResultActionsCodeR03, AuthRuleV2ListResultsResponseACHResultActionsCodeR04, AuthRuleV2ListResultsResponseACHResultActionsCodeR05, AuthRuleV2ListResultsResponseACHResultActionsCodeR06, AuthRuleV2ListResultsResponseACHResultActionsCodeR07, AuthRuleV2ListResultsResponseACHResultActionsCodeR08, AuthRuleV2ListResultsResponseACHResultActionsCodeR09, AuthRuleV2ListResultsResponseACHResultActionsCodeR10, AuthRuleV2ListResultsResponseACHResultActionsCodeR11, AuthRuleV2ListResultsResponseACHResultActionsCodeR12, AuthRuleV2ListResultsResponseACHResultActionsCodeR13, AuthRuleV2ListResultsResponseACHResultActionsCodeR14, AuthRuleV2ListResultsResponseACHResultActionsCodeR15, AuthRuleV2ListResultsResponseACHResultActionsCodeR16, AuthRuleV2ListResultsResponseACHResultActionsCodeR17, AuthRuleV2ListResultsResponseACHResultActionsCodeR18, AuthRuleV2ListResultsResponseACHResultActionsCodeR19, AuthRuleV2ListResultsResponseACHResultActionsCodeR20, AuthRuleV2ListResultsResponseACHResultActionsCodeR21, AuthRuleV2ListResultsResponseACHResultActionsCodeR22, AuthRuleV2ListResultsResponseACHResultActionsCodeR23, AuthRuleV2ListResultsResponseACHResultActionsCodeR24, AuthRuleV2ListResultsResponseACHResultActionsCodeR25, AuthRuleV2ListResultsResponseACHResultActionsCodeR26, AuthRuleV2ListResultsResponseACHResultActionsCodeR27, AuthRuleV2ListResultsResponseACHResultActionsCodeR28, AuthRuleV2ListResultsResponseACHResultActionsCodeR29, AuthRuleV2ListResultsResponseACHResultActionsCodeR30, AuthRuleV2ListResultsResponseACHResultActionsCodeR31, AuthRuleV2ListResultsResponseACHResultActionsCodeR32, AuthRuleV2ListResultsResponseACHResultActionsCodeR33, AuthRuleV2ListResultsResponseACHResultActionsCodeR34, AuthRuleV2ListResultsResponseACHResultActionsCodeR35, AuthRuleV2ListResultsResponseACHResultActionsCodeR36, AuthRuleV2ListResultsResponseACHResultActionsCodeR37, AuthRuleV2ListResultsResponseACHResultActionsCodeR38, AuthRuleV2ListResultsResponseACHResultActionsCodeR39, AuthRuleV2ListResultsResponseACHResultActionsCodeR40, AuthRuleV2ListResultsResponseACHResultActionsCodeR41, AuthRuleV2ListResultsResponseACHResultActionsCodeR42, AuthRuleV2ListResultsResponseACHResultActionsCodeR43, AuthRuleV2ListResultsResponseACHResultActionsCodeR44, AuthRuleV2ListResultsResponseACHResultActionsCodeR45, AuthRuleV2ListResultsResponseACHResultActionsCodeR46, AuthRuleV2ListResultsResponseACHResultActionsCodeR47, AuthRuleV2ListResultsResponseACHResultActionsCodeR50, AuthRuleV2ListResultsResponseACHResultActionsCodeR51, AuthRuleV2ListResultsResponseACHResultActionsCodeR52, AuthRuleV2ListResultsResponseACHResultActionsCodeR53, AuthRuleV2ListResultsResponseACHResultActionsCodeR61, AuthRuleV2ListResultsResponseACHResultActionsCodeR62, AuthRuleV2ListResultsResponseACHResultActionsCodeR67, AuthRuleV2ListResultsResponseACHResultActionsCodeR68, AuthRuleV2ListResultsResponseACHResultActionsCodeR69, AuthRuleV2ListResultsResponseACHResultActionsCodeR70, AuthRuleV2ListResultsResponseACHResultActionsCodeR71, AuthRuleV2ListResultsResponseACHResultActionsCodeR72, AuthRuleV2ListResultsResponseACHResultActionsCodeR73, AuthRuleV2ListResultsResponseACHResultActionsCodeR74, AuthRuleV2ListResultsResponseACHResultActionsCodeR75, AuthRuleV2ListResultsResponseACHResultActionsCodeR76, AuthRuleV2ListResultsResponseACHResultActionsCodeR77, AuthRuleV2ListResultsResponseACHResultActionsCodeR80, AuthRuleV2ListResultsResponseACHResultActionsCodeR81, AuthRuleV2ListResultsResponseACHResultActionsCodeR82, AuthRuleV2ListResultsResponseACHResultActionsCodeR83, AuthRuleV2ListResultsResponseACHResultActionsCodeR84, AuthRuleV2ListResultsResponseACHResultActionsCodeR85:
		return true
	}
	return false
}

// The event stream during which the rule was evaluated
type AuthRuleV2ListResultsResponseACHResultEventStream string

const (
	AuthRuleV2ListResultsResponseACHResultEventStreamACHCreditReceipt AuthRuleV2ListResultsResponseACHResultEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2ListResultsResponseACHResultEventStreamACHDebitReceipt  AuthRuleV2ListResultsResponseACHResultEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2ListResultsResponseACHResultEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseACHResultEventStreamACHCreditReceipt, AuthRuleV2ListResultsResponseACHResultEventStreamACHDebitReceipt:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2ListResultsResponseACHResultMode string

const (
	AuthRuleV2ListResultsResponseACHResultModeActive   AuthRuleV2ListResultsResponseACHResultMode = "ACTIVE"
	AuthRuleV2ListResultsResponseACHResultModeInactive AuthRuleV2ListResultsResponseACHResultMode = "INACTIVE"
)

func (r AuthRuleV2ListResultsResponseACHResultMode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseACHResultModeActive, AuthRuleV2ListResultsResponseACHResultModeInactive:
		return true
	}
	return false
}

// The event stream during which the rule was evaluated
type AuthRuleV2ListResultsResponseEventStream string

const (
	AuthRuleV2ListResultsResponseEventStreamAuthorization         AuthRuleV2ListResultsResponseEventStream = "AUTHORIZATION"
	AuthRuleV2ListResultsResponseEventStreamThreeDSAuthentication AuthRuleV2ListResultsResponseEventStream = "THREE_DS_AUTHENTICATION"
	AuthRuleV2ListResultsResponseEventStreamTokenization          AuthRuleV2ListResultsResponseEventStream = "TOKENIZATION"
	AuthRuleV2ListResultsResponseEventStreamACHCreditReceipt      AuthRuleV2ListResultsResponseEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2ListResultsResponseEventStreamACHDebitReceipt       AuthRuleV2ListResultsResponseEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2ListResultsResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseEventStreamAuthorization, AuthRuleV2ListResultsResponseEventStreamThreeDSAuthentication, AuthRuleV2ListResultsResponseEventStreamTokenization, AuthRuleV2ListResultsResponseEventStreamACHCreditReceipt, AuthRuleV2ListResultsResponseEventStreamACHDebitReceipt:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2ListResultsResponseMode string

const (
	AuthRuleV2ListResultsResponseModeActive   AuthRuleV2ListResultsResponseMode = "ACTIVE"
	AuthRuleV2ListResultsResponseModeInactive AuthRuleV2ListResultsResponseMode = "INACTIVE"
)

func (r AuthRuleV2ListResultsResponseMode) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResultsResponseModeActive, AuthRuleV2ListResultsResponseModeInactive:
		return true
	}
	return false
}

type AuthRuleV2ListVersionsResponse struct {
	Data []AuthRuleVersion                  `json:"data" api:"required"`
	JSON authRuleV2ListVersionsResponseJSON `json:"-"`
}

// authRuleV2ListVersionsResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ListVersionsResponse]
type authRuleV2ListVersionsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListVersionsResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetFeaturesResponse struct {
	// Timestamp at which the Features were evaluated
	Evaluated time.Time `json:"evaluated" api:"required" format:"date-time"`
	// Calculated Features used for evaluation of the provided Auth Rule
	Features []AuthRuleV2GetFeaturesResponseFeature `json:"features" api:"required"`
	JSON     authRuleV2GetFeaturesResponseJSON      `json:"-"`
}

// authRuleV2GetFeaturesResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2GetFeaturesResponse]
type authRuleV2GetFeaturesResponseJSON struct {
	Evaluated   apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetFeaturesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetFeaturesResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetFeaturesResponseFeature struct {
	Filters VelocityLimitFilters `json:"filters" api:"required"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period" api:"required"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2GetFeaturesResponseFeaturesScope `json:"scope" api:"required"`
	Value AuthRuleV2GetFeaturesResponseFeaturesValue `json:"value" api:"required"`
	JSON  authRuleV2GetFeaturesResponseFeatureJSON   `json:"-"`
}

// authRuleV2GetFeaturesResponseFeatureJSON contains the JSON metadata for the
// struct [AuthRuleV2GetFeaturesResponseFeature]
type authRuleV2GetFeaturesResponseFeatureJSON struct {
	Filters     apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetFeaturesResponseFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetFeaturesResponseFeatureJSON) RawJSON() string {
	return r.raw
}

// The scope the velocity is calculated for
type AuthRuleV2GetFeaturesResponseFeaturesScope string

const (
	AuthRuleV2GetFeaturesResponseFeaturesScopeCard    AuthRuleV2GetFeaturesResponseFeaturesScope = "CARD"
	AuthRuleV2GetFeaturesResponseFeaturesScopeAccount AuthRuleV2GetFeaturesResponseFeaturesScope = "ACCOUNT"
)

func (r AuthRuleV2GetFeaturesResponseFeaturesScope) IsKnown() bool {
	switch r {
	case AuthRuleV2GetFeaturesResponseFeaturesScopeCard, AuthRuleV2GetFeaturesResponseFeaturesScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2GetFeaturesResponseFeaturesValue struct {
	// Amount (in cents) for the given Auth Rule that is used as input for calculating
	// the rule. For Velocity Limit rules this would be the calculated Velocity. For
	// Conditional Rules using CARD*TRANSACTION_COUNT*\* this will be 0
	Amount int64 `json:"amount" api:"required"`
	// Number of velocity impacting transactions matching the given scope, period and
	// filters
	Count int64                                          `json:"count" api:"required"`
	JSON  authRuleV2GetFeaturesResponseFeaturesValueJSON `json:"-"`
}

// authRuleV2GetFeaturesResponseFeaturesValueJSON contains the JSON metadata for
// the struct [AuthRuleV2GetFeaturesResponseFeaturesValue]
type authRuleV2GetFeaturesResponseFeaturesValueJSON struct {
	Amount      apijson.Field
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetFeaturesResponseFeaturesValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetFeaturesResponseFeaturesValueJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetReportResponse struct {
	// Auth Rule Token
	AuthRuleToken string `json:"auth_rule_token" api:"required" format:"uuid"`
	// The start date (UTC) of the report.
	Begin time.Time `json:"begin" api:"required" format:"date"`
	// Daily evaluation statistics for the Auth Rule.
	DailyStatistics []AuthRuleV2GetReportResponseDailyStatistic `json:"daily_statistics" api:"required"`
	// The end date (UTC) of the report.
	End  time.Time                       `json:"end" api:"required" format:"date"`
	JSON authRuleV2GetReportResponseJSON `json:"-"`
}

// authRuleV2GetReportResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2GetReportResponse]
type authRuleV2GetReportResponseJSON struct {
	AuthRuleToken   apijson.Field
	Begin           apijson.Field
	DailyStatistics apijson.Field
	End             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetReportResponseDailyStatistic struct {
	// Detailed statistics for the current version of the rule.
	CurrentVersionStatistics ReportStats `json:"current_version_statistics" api:"required,nullable"`
	// The date (UTC) for which the statistics are reported.
	Date time.Time `json:"date" api:"required" format:"date"`
	// Detailed statistics for the draft version of the rule.
	DraftVersionStatistics ReportStats `json:"draft_version_statistics" api:"required,nullable"`
	// Statistics for each version of the rule that was evaluated during the reported
	// day.
	Versions []AuthRuleV2GetReportResponseDailyStatisticsVersion `json:"versions" api:"required"`
	JSON     authRuleV2GetReportResponseDailyStatisticJSON       `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticJSON contains the JSON metadata for the
// struct [AuthRuleV2GetReportResponseDailyStatistic]
type authRuleV2GetReportResponseDailyStatisticJSON struct {
	CurrentVersionStatistics apijson.Field
	Date                     apijson.Field
	DraftVersionStatistics   apijson.Field
	Versions                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatistic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetReportResponseDailyStatisticsVersion struct {
	// A mapping of action types to the number of times that action was returned by
	// this version during the relevant period. Actions are the possible outcomes of a
	// rule evaluation, such as DECLINE, CHALLENGE, REQUIRE_TFA, etc. In case rule
	// didn't trigger any action, it's counted under NO_ACTION key.
	ActionCounts map[string]int64 `json:"action_counts" api:"required"`
	// Example events and their outcomes for this version.
	Examples []AuthRuleV2GetReportResponseDailyStatisticsVersionsExample `json:"examples" api:"required"`
	// The evaluation mode of this version during the reported period.
	State AuthRuleV2GetReportResponseDailyStatisticsVersionsState `json:"state" api:"required"`
	// The rule version number.
	Version int64                                                 `json:"version" api:"required"`
	JSON    authRuleV2GetReportResponseDailyStatisticsVersionJSON `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionJSON contains the JSON metadata
// for the struct [AuthRuleV2GetReportResponseDailyStatisticsVersion]
type authRuleV2GetReportResponseDailyStatisticsVersionJSON struct {
	ActionCounts apijson.Field
	Examples     apijson.Field
	State        apijson.Field
	Version      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExample struct {
	// The actions taken by this version for this event.
	Actions []AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction `json:"actions" api:"required"`
	// The event token.
	EventToken string `json:"event_token" api:"required" format:"uuid"`
	// The timestamp of the event.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The token of the transaction associated with the event
	TransactionToken string                                                        `json:"transaction_token" api:"nullable" format:"uuid"`
	JSON             authRuleV2GetReportResponseDailyStatisticsVersionsExampleJSON `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExampleJSON contains the JSON
// metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExample]
type authRuleV2GetReportResponseDailyStatisticsVersionsExampleJSON struct {
	Actions          apijson.Field
	EventToken       apijson.Field
	Timestamp        apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExample) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExampleJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction struct {
	Type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsType `json:"type" api:"required"`
	// The detailed result code explaining the specific reason for the decline
	Code AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode `json:"code"`
	// Reason code for declining the tokenization request
	Reason AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason `json:"reason"`
	JSON   authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionJSON    `json:"-"`
	union  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsUnion
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionJSON contains
// the JSON metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction]
type authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorization],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorization],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSAction],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenization],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaAction],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACH],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnAction].
func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction) AsUnion() AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsUnion {
	return r.union
}

// Union satisfied by
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorization],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorization],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSAction],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenization],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaAction],
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACH]
// or
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnAction].
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsUnion interface {
	implementsAuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSAction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaAction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACH{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnAction{}),
		},
	)
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorization struct {
	// The detailed result code explaining the specific reason for the decline
	Code AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode `json:"code" api:"required"`
	Type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationType `json:"type" api:"required"`
	JSON authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationJSON `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorization]
type authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorization) implementsAuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction() {
}

// The detailed result code explaining the specific reason for the decline
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountDailySpendLimitExceeded              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_DAILY_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountDelinquent                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_DELINQUENT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountInactive                             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_INACTIVE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountLifetimeSpendLimitExceeded           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_LIFETIME_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountMonthlySpendLimitExceeded            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_MONTHLY_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountPaused                               AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_PAUSED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountUnderReview                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "ACCOUNT_UNDER_REVIEW"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAddressIncorrect                            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "ADDRESS_INCORRECT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeApproved                                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "APPROVED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRuleAllowedCountry                      AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE_ALLOWED_COUNTRY"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRuleAllowedMcc                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE_ALLOWED_MCC"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRuleBlockedCountry                      AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE_BLOCKED_COUNTRY"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRuleBlockedMcc                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE_BLOCKED_MCC"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRule                                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "AUTH_RULE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardClosed                                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_CLOSED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardCryptogramValidationFailure             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_CRYPTOGRAM_VALIDATION_FAILURE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardExpired                                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_EXPIRED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardExpiryDateIncorrect                     AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_EXPIRY_DATE_INCORRECT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardInvalid                                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardNotActivated                            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_NOT_ACTIVATED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardPaused                                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_PAUSED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardPinIncorrect                            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_PIN_INCORRECT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardRestricted                              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_RESTRICTED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardSecurityCodeIncorrect                   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_SECURITY_CODE_INCORRECT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardSpendLimitExceeded                      AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARD_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeContactCardIssuer                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CONTACT_CARD_ISSUER"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCustomerAsaTimeout                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CUSTOMER_ASA_TIMEOUT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCustomAsaResult                             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CUSTOM_ASA_RESULT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeDeclined                                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "DECLINED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeDoNotHonor                                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "DO_NOT_HONOR"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeDriverNumberInvalid                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "DRIVER_NUMBER_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeFormatError                                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "FORMAT_ERROR"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeInsufficientFundingSourceBalance            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "INSUFFICIENT_FUNDING_SOURCE_BALANCE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeInsufficientFunds                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "INSUFFICIENT_FUNDS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeLithicSystemError                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "LITHIC_SYSTEM_ERROR"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeLithicSystemRateLimit                       AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "LITHIC_SYSTEM_RATE_LIMIT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeMalformedAsaResponse                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "MALFORMED_ASA_RESPONSE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeMerchantInvalid                             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "MERCHANT_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeMerchantLockedCardAttemptedElsewhere        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "MERCHANT_LOCKED_CARD_ATTEMPTED_ELSEWHERE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeMerchantNotPermitted                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "MERCHANT_NOT_PERMITTED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeOverReversalAttempted                       AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "OVER_REVERSAL_ATTEMPTED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodePinBlocked                                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "PIN_BLOCKED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeProgramCardSpendLimitExceeded               AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "PROGRAM_CARD_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeProgramSuspended                            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "PROGRAM_SUSPENDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeProgramUsageRestriction                     AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "PROGRAM_USAGE_RESTRICTION"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeReversalUnmatched                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "REVERSAL_UNMATCHED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeSecurityViolation                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "SECURITY_VIOLATION"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeSingleUseCardReattempted                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "SINGLE_USE_CARD_REATTEMPTED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeSuspectedFraud                              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "SUSPECTED_FRAUD"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeTransactionInvalid                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "TRANSACTION_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeTransactionNotPermittedToAcquirerOrTerminal AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "TRANSACTION_NOT_PERMITTED_TO_ACQUIRER_OR_TERMINAL"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeTransactionNotPermittedToIssuerOrCardholder AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "TRANSACTION_NOT_PERMITTED_TO_ISSUER_OR_CARDHOLDER"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeTransactionPreviouslyCompleted              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "TRANSACTION_PREVIOUSLY_COMPLETED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeUnauthorizedMerchant                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "UNAUTHORIZED_MERCHANT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeVehicleNumberInvalid                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "VEHICLE_NUMBER_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardholderChallenged                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARDHOLDER_CHALLENGED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardholderChallengeFailed                   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode = "CARDHOLDER_CHALLENGE_FAILED"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCode) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountDailySpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountDelinquent, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountInactive, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountLifetimeSpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountMonthlySpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountPaused, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAccountUnderReview, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAddressIncorrect, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeApproved, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRuleAllowedCountry, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRuleAllowedMcc, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRuleBlockedCountry, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRuleBlockedMcc, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeAuthRule, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardClosed, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardCryptogramValidationFailure, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardExpired, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardExpiryDateIncorrect, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardNotActivated, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardPaused, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardPinIncorrect, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardRestricted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardSecurityCodeIncorrect, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardSpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeContactCardIssuer, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCustomerAsaTimeout, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCustomAsaResult, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeDeclined, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeDoNotHonor, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeDriverNumberInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeFormatError, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeInsufficientFundingSourceBalance, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeInsufficientFunds, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeLithicSystemError, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeLithicSystemRateLimit, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeMalformedAsaResponse, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeMerchantInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeMerchantLockedCardAttemptedElsewhere, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeMerchantNotPermitted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeOverReversalAttempted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodePinBlocked, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeProgramCardSpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeProgramSuspended, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeProgramUsageRestriction, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeReversalUnmatched, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeSecurityViolation, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeSingleUseCardReattempted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeSuspectedFraud, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeTransactionInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeTransactionNotPermittedToAcquirerOrTerminal, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeTransactionNotPermittedToIssuerOrCardholder, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeTransactionPreviouslyCompleted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeUnauthorizedMerchant, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeVehicleNumberInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardholderChallenged, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationCodeCardholderChallengeFailed:
		return true
	}
	return false
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationType string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationTypeDecline AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationType = "DECLINE"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionAuthorizationTypeDecline:
		return true
	}
	return false
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorization struct {
	Type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationType `json:"type" api:"required"`
	JSON authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationJSON `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorization]
type authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorization) implementsAuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction() {
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationType string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationTypeChallenge AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationType = "CHALLENGE"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsChallengeActionAuthorizationTypeChallenge:
		return true
	}
	return false
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSAction struct {
	Type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionType `json:"type" api:"required"`
	JSON authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DsActionJSON `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DsActionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSAction]
type authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DsActionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DsActionJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSAction) implementsAuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction() {
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionType string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionTypeDecline   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionType = "DECLINE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionTypeChallenge AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionType = "CHALLENGE"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionTypeDecline, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsResultAuthentication3DSActionTypeChallenge:
		return true
	}
	return false
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenization struct {
	// Decline the tokenization request
	Type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationType `json:"type" api:"required"`
	// Reason code for declining the tokenization request
	Reason AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason `json:"reason"`
	JSON   authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationJSON   `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenization]
type authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationJSON struct {
	Type        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenization) implementsAuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction() {
}

// Decline the tokenization request
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationType string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationTypeDecline AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationType = "DECLINE"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationTypeDecline:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonAccountScore1                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "ACCOUNT_SCORE_1"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonDeviceScore1                   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "DEVICE_SCORE_1"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonAllWalletDeclineReasonsPresent AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonWalletRecommendedDecisionRed   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "WALLET_RECOMMENDED_DECISION_RED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCvcMismatch                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "CVC_MISMATCH"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCardExpiryMonthMismatch        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "CARD_EXPIRY_MONTH_MISMATCH"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCardExpiryYearMismatch         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "CARD_EXPIRY_YEAR_MISMATCH"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCardInvalidState               AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "CARD_INVALID_STATE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCustomerRedPath                AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "CUSTOMER_RED_PATH"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonInvalidCustomerResponse        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "INVALID_CUSTOMER_RESPONSE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonNetworkFailure                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "NETWORK_FAILURE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonGenericDecline                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "GENERIC_DECLINE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonDigitalCardArtRequired         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason = "DIGITAL_CARD_ART_REQUIRED"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReason) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonAccountScore1, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonDeviceScore1, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonAllWalletDeclineReasonsPresent, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonWalletRecommendedDecisionRed, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCvcMismatch, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCardExpiryMonthMismatch, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCardExpiryYearMismatch, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCardInvalidState, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonCustomerRedPath, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonInvalidCustomerResponse, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonNetworkFailure, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonGenericDecline, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsDeclineActionTokenizationReasonDigitalCardArtRequired:
		return true
	}
	return false
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaAction struct {
	// Require two-factor authentication for the tokenization request
	Type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionType `json:"type" api:"required"`
	// Reason code for requiring two-factor authentication
	Reason AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason `json:"reason"`
	JSON   authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionJSON   `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaAction]
type authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionJSON struct {
	Type        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaAction) implementsAuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction() {
}

// Require two-factor authentication for the tokenization request
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionType string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionTypeRequireTfa AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionType = "REQUIRE_TFA"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionTypeRequireTfa:
		return true
	}
	return false
}

// Reason code for requiring two-factor authentication
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonWalletRecommendedTfa        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "WALLET_RECOMMENDED_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonSuspiciousActivity          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "SUSPICIOUS_ACTIVITY"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonDeviceRecentlyLost          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "DEVICE_RECENTLY_LOST"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonTooManyRecentAttempts       AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "TOO_MANY_RECENT_ATTEMPTS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonTooManyRecentTokens         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "TOO_MANY_RECENT_TOKENS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonTooManyDifferentCardholders AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonOutsideHomeTerritory        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "OUTSIDE_HOME_TERRITORY"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonHasSuspendedTokens          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "HAS_SUSPENDED_TOKENS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonHighRisk                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "HIGH_RISK"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonAccountScoreLow             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "ACCOUNT_SCORE_LOW"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonDeviceScoreLow              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "DEVICE_SCORE_LOW"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonCardStateTfa                AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "CARD_STATE_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonHardcodedTfa                AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "HARDCODED_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonCustomerRuleTfa             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "CUSTOMER_RULE_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonDeviceHostCardEmulation     AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReason) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonWalletRecommendedTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonSuspiciousActivity, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonDeviceRecentlyLost, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonTooManyRecentAttempts, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonTooManyRecentTokens, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonTooManyDifferentCardholders, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonOutsideHomeTerritory, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonHasSuspendedTokens, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonHighRisk, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonAccountScoreLow, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonDeviceScoreLow, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonCardStateTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonHardcodedTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonCustomerRuleTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsRequireTfaActionReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACH struct {
	// Approve the ACH transaction
	Type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHType `json:"type" api:"required"`
	JSON authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHJSON `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACH]
type authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACH) implementsAuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction() {
}

// Approve the ACH transaction
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHType string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHTypeApprove AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHType = "APPROVE"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsApproveActionACHTypeApprove:
		return true
	}
	return false
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnAction struct {
	// NACHA return code to use when returning the transaction. Note that the list of
	// available return codes is subject to an allowlist configured at the program
	// level
	Code AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode `json:"code" api:"required"`
	// Return the ACH transaction
	Type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionType `json:"type" api:"required"`
	JSON authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionJSON `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnAction]
type authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnAction) implementsAuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesAction() {
}

// NACHA return code to use when returning the transaction. Note that the list of
// available return codes is subject to an allowlist configured at the program
// level
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR01 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R01"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR02 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R02"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR03 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R03"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR04 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R04"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR05 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R05"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR06 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R06"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR07 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R07"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR08 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R08"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR09 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R09"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR10 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R10"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR11 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R11"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR12 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R12"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR13 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R13"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR14 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R14"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR15 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R15"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR16 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R16"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR17 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R17"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR18 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R18"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR19 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R19"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR20 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R20"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR21 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R21"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR22 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R22"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR23 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R23"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR24 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R24"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR25 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R25"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR26 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R26"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR27 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R27"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR28 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R28"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR29 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R29"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR30 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R30"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR31 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R31"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR32 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R32"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR33 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R33"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR34 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R34"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR35 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R35"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR36 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R36"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR37 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R37"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR38 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R38"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR39 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R39"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR40 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R40"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR41 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R41"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR42 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R42"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR43 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R43"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR44 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R44"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR45 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R45"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR46 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R46"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR47 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R47"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR50 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R50"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR51 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R51"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR52 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R52"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR53 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R53"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR61 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R61"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR62 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R62"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR67 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R67"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR68 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R68"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR69 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R69"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR70 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R70"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR71 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R71"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR72 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R72"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR73 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R73"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR74 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R74"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR75 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R75"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR76 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R76"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR77 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R77"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR80 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R80"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR81 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R81"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR82 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R82"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR83 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R83"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR84 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R84"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR85 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode = "R85"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCode) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR01, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR02, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR03, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR04, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR05, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR06, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR07, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR08, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR09, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR10, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR11, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR12, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR13, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR14, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR15, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR16, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR17, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR18, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR19, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR20, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR21, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR22, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR23, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR24, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR25, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR26, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR27, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR28, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR29, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR30, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR31, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR32, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR33, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR34, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR35, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR36, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR37, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR38, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR39, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR40, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR41, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR42, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR43, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR44, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR45, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR46, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR47, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR50, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR51, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR52, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR53, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR61, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR62, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR67, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR68, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR69, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR70, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR71, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR72, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR73, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR74, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR75, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR76, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR77, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR80, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR81, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR82, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR83, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR84, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionCodeR85:
		return true
	}
	return false
}

// Return the ACH transaction
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionType string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionTypeReturn AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionType = "RETURN"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReturnActionTypeReturn:
		return true
	}
	return false
}

type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsType string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeDecline    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsType = "DECLINE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeChallenge  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsType = "CHALLENGE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeRequireTfa AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsType = "REQUIRE_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeApprove    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsType = "APPROVE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeReturn     AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsType = "RETURN"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeDecline, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeChallenge, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeRequireTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeApprove, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsTypeReturn:
		return true
	}
	return false
}

// The detailed result code explaining the specific reason for the decline
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountDailySpendLimitExceeded              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "ACCOUNT_DAILY_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountDelinquent                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "ACCOUNT_DELINQUENT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountInactive                             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "ACCOUNT_INACTIVE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountLifetimeSpendLimitExceeded           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "ACCOUNT_LIFETIME_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountMonthlySpendLimitExceeded            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "ACCOUNT_MONTHLY_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountPaused                               AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "ACCOUNT_PAUSED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountUnderReview                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "ACCOUNT_UNDER_REVIEW"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAddressIncorrect                            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "ADDRESS_INCORRECT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeApproved                                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "APPROVED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRuleAllowedCountry                      AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "AUTH_RULE_ALLOWED_COUNTRY"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRuleAllowedMcc                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "AUTH_RULE_ALLOWED_MCC"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRuleBlockedCountry                      AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "AUTH_RULE_BLOCKED_COUNTRY"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRuleBlockedMcc                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "AUTH_RULE_BLOCKED_MCC"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRule                                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "AUTH_RULE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardClosed                                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_CLOSED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardCryptogramValidationFailure             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_CRYPTOGRAM_VALIDATION_FAILURE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardExpired                                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_EXPIRED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardExpiryDateIncorrect                     AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_EXPIRY_DATE_INCORRECT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardInvalid                                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardNotActivated                            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_NOT_ACTIVATED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardPaused                                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_PAUSED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardPinIncorrect                            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_PIN_INCORRECT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardRestricted                              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_RESTRICTED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardSecurityCodeIncorrect                   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_SECURITY_CODE_INCORRECT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardSpendLimitExceeded                      AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARD_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeContactCardIssuer                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CONTACT_CARD_ISSUER"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCustomerAsaTimeout                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CUSTOMER_ASA_TIMEOUT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCustomAsaResult                             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CUSTOM_ASA_RESULT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeDeclined                                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "DECLINED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeDoNotHonor                                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "DO_NOT_HONOR"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeDriverNumberInvalid                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "DRIVER_NUMBER_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeFormatError                                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "FORMAT_ERROR"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeInsufficientFundingSourceBalance            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "INSUFFICIENT_FUNDING_SOURCE_BALANCE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeInsufficientFunds                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "INSUFFICIENT_FUNDS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeLithicSystemError                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "LITHIC_SYSTEM_ERROR"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeLithicSystemRateLimit                       AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "LITHIC_SYSTEM_RATE_LIMIT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeMalformedAsaResponse                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "MALFORMED_ASA_RESPONSE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeMerchantInvalid                             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "MERCHANT_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeMerchantLockedCardAttemptedElsewhere        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "MERCHANT_LOCKED_CARD_ATTEMPTED_ELSEWHERE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeMerchantNotPermitted                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "MERCHANT_NOT_PERMITTED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeOverReversalAttempted                       AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "OVER_REVERSAL_ATTEMPTED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodePinBlocked                                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "PIN_BLOCKED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeProgramCardSpendLimitExceeded               AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "PROGRAM_CARD_SPEND_LIMIT_EXCEEDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeProgramSuspended                            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "PROGRAM_SUSPENDED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeProgramUsageRestriction                     AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "PROGRAM_USAGE_RESTRICTION"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeReversalUnmatched                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "REVERSAL_UNMATCHED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeSecurityViolation                           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "SECURITY_VIOLATION"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeSingleUseCardReattempted                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "SINGLE_USE_CARD_REATTEMPTED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeSuspectedFraud                              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "SUSPECTED_FRAUD"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeTransactionInvalid                          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "TRANSACTION_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeTransactionNotPermittedToAcquirerOrTerminal AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "TRANSACTION_NOT_PERMITTED_TO_ACQUIRER_OR_TERMINAL"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeTransactionNotPermittedToIssuerOrCardholder AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "TRANSACTION_NOT_PERMITTED_TO_ISSUER_OR_CARDHOLDER"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeTransactionPreviouslyCompleted              AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "TRANSACTION_PREVIOUSLY_COMPLETED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeUnauthorizedMerchant                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "UNAUTHORIZED_MERCHANT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeVehicleNumberInvalid                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "VEHICLE_NUMBER_INVALID"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardholderChallenged                        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARDHOLDER_CHALLENGED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardholderChallengeFailed                   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "CARDHOLDER_CHALLENGE_FAILED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR01                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R01"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR02                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R02"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR03                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R03"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR04                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R04"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR05                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R05"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR06                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R06"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR07                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R07"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR08                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R08"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR09                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R09"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR10                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R10"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR11                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R11"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR12                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R12"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR13                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R13"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR14                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R14"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR15                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R15"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR16                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R16"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR17                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R17"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR18                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R18"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR19                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R19"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR20                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R20"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR21                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R21"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR22                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R22"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR23                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R23"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR24                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R24"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR25                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R25"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR26                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R26"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR27                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R27"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR28                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R28"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR29                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R29"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR30                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R30"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR31                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R31"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR32                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R32"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR33                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R33"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR34                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R34"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR35                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R35"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR36                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R36"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR37                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R37"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR38                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R38"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR39                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R39"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR40                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R40"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR41                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R41"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR42                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R42"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR43                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R43"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR44                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R44"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR45                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R45"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR46                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R46"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR47                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R47"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR50                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R50"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR51                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R51"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR52                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R52"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR53                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R53"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR61                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R61"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR62                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R62"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR67                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R67"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR68                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R68"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR69                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R69"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR70                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R70"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR71                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R71"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR72                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R72"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR73                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R73"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR74                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R74"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR75                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R75"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR76                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R76"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR77                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R77"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR80                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R80"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR81                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R81"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR82                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R82"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR83                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R83"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR84                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R84"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR85                                         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode = "R85"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCode) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountDailySpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountDelinquent, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountInactive, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountLifetimeSpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountMonthlySpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountPaused, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAccountUnderReview, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAddressIncorrect, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeApproved, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRuleAllowedCountry, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRuleAllowedMcc, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRuleBlockedCountry, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRuleBlockedMcc, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeAuthRule, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardClosed, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardCryptogramValidationFailure, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardExpired, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardExpiryDateIncorrect, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardNotActivated, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardPaused, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardPinIncorrect, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardRestricted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardSecurityCodeIncorrect, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardSpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeContactCardIssuer, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCustomerAsaTimeout, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCustomAsaResult, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeDeclined, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeDoNotHonor, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeDriverNumberInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeFormatError, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeInsufficientFundingSourceBalance, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeInsufficientFunds, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeLithicSystemError, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeLithicSystemRateLimit, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeMalformedAsaResponse, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeMerchantInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeMerchantLockedCardAttemptedElsewhere, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeMerchantNotPermitted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeOverReversalAttempted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodePinBlocked, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeProgramCardSpendLimitExceeded, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeProgramSuspended, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeProgramUsageRestriction, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeReversalUnmatched, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeSecurityViolation, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeSingleUseCardReattempted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeSuspectedFraud, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeTransactionInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeTransactionNotPermittedToAcquirerOrTerminal, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeTransactionNotPermittedToIssuerOrCardholder, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeTransactionPreviouslyCompleted, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeUnauthorizedMerchant, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeVehicleNumberInvalid, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardholderChallenged, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeCardholderChallengeFailed, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR01, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR02, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR03, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR04, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR05, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR06, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR07, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR08, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR09, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR10, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR11, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR12, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR13, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR14, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR15, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR16, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR17, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR18, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR19, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR20, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR21, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR22, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR23, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR24, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR25, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR26, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR27, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR28, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR29, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR30, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR31, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR32, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR33, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR34, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR35, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR36, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR37, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR38, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR39, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR40, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR41, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR42, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR43, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR44, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR45, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR46, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR47, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR50, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR51, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR52, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR53, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR61, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR62, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR67, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR68, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR69, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR70, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR71, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR72, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR73, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR74, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR75, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR76, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR77, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR80, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR81, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR82, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR83, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR84, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsCodeR85:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonAccountScore1                  AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "ACCOUNT_SCORE_1"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDeviceScore1                   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "DEVICE_SCORE_1"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonAllWalletDeclineReasonsPresent AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonWalletRecommendedDecisionRed   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "WALLET_RECOMMENDED_DECISION_RED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCvcMismatch                    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "CVC_MISMATCH"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCardExpiryMonthMismatch        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "CARD_EXPIRY_MONTH_MISMATCH"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCardExpiryYearMismatch         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "CARD_EXPIRY_YEAR_MISMATCH"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCardInvalidState               AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "CARD_INVALID_STATE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCustomerRedPath                AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "CUSTOMER_RED_PATH"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonInvalidCustomerResponse        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "INVALID_CUSTOMER_RESPONSE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonNetworkFailure                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "NETWORK_FAILURE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonGenericDecline                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "GENERIC_DECLINE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDigitalCardArtRequired         AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "DIGITAL_CARD_ART_REQUIRED"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonWalletRecommendedTfa           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "WALLET_RECOMMENDED_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonSuspiciousActivity             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "SUSPICIOUS_ACTIVITY"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDeviceRecentlyLost             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "DEVICE_RECENTLY_LOST"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonTooManyRecentAttempts          AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "TOO_MANY_RECENT_ATTEMPTS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonTooManyRecentTokens            AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "TOO_MANY_RECENT_TOKENS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonTooManyDifferentCardholders    AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "TOO_MANY_DIFFERENT_CARDHOLDERS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonOutsideHomeTerritory           AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "OUTSIDE_HOME_TERRITORY"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonHasSuspendedTokens             AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "HAS_SUSPENDED_TOKENS"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonHighRisk                       AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "HIGH_RISK"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonAccountScoreLow                AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "ACCOUNT_SCORE_LOW"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDeviceScoreLow                 AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "DEVICE_SCORE_LOW"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCardStateTfa                   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "CARD_STATE_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonHardcodedTfa                   AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "HARDCODED_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCustomerRuleTfa                AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "CUSTOMER_RULE_TFA"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDeviceHostCardEmulation        AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason = "DEVICE_HOST_CARD_EMULATION"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReason) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonAccountScore1, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDeviceScore1, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonAllWalletDeclineReasonsPresent, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonWalletRecommendedDecisionRed, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCvcMismatch, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCardExpiryMonthMismatch, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCardExpiryYearMismatch, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCardInvalidState, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCustomerRedPath, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonInvalidCustomerResponse, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonNetworkFailure, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonGenericDecline, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDigitalCardArtRequired, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonWalletRecommendedTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonSuspiciousActivity, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDeviceRecentlyLost, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonTooManyRecentAttempts, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonTooManyRecentTokens, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonTooManyDifferentCardholders, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonOutsideHomeTerritory, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonHasSuspendedTokens, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonHighRisk, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonAccountScoreLow, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDeviceScoreLow, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCardStateTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonHardcodedTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonCustomerRuleTfa, AuthRuleV2GetReportResponseDailyStatisticsVersionsExamplesActionsReasonDeviceHostCardEmulation:
		return true
	}
	return false
}

// The evaluation mode of this version during the reported period.
type AuthRuleV2GetReportResponseDailyStatisticsVersionsState string

const (
	AuthRuleV2GetReportResponseDailyStatisticsVersionsStateActive   AuthRuleV2GetReportResponseDailyStatisticsVersionsState = "ACTIVE"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsStateShadow   AuthRuleV2GetReportResponseDailyStatisticsVersionsState = "SHADOW"
	AuthRuleV2GetReportResponseDailyStatisticsVersionsStateInactive AuthRuleV2GetReportResponseDailyStatisticsVersionsState = "INACTIVE"
)

func (r AuthRuleV2GetReportResponseDailyStatisticsVersionsState) IsKnown() bool {
	switch r {
	case AuthRuleV2GetReportResponseDailyStatisticsVersionsStateActive, AuthRuleV2GetReportResponseDailyStatisticsVersionsStateShadow, AuthRuleV2GetReportResponseDailyStatisticsVersionsStateInactive:
		return true
	}
	return false
}

type AuthRuleV2NewParams struct {
	Body AuthRuleV2NewParamsBodyUnion `json:"body" api:"required"`
}

func (r AuthRuleV2NewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AuthRuleV2NewParamsBody struct {
	Parameters param.Field[interface{}] `json:"parameters" api:"required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
	//     AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type                  param.Field[AuthRuleV2NewParamsBodyType] `json:"type" api:"required"`
	AccountTokens         param.Field[interface{}]                 `json:"account_tokens"`
	BusinessAccountTokens param.Field[interface{}]                 `json:"business_account_tokens"`
	CardTokens            param.Field[interface{}]                 `json:"card_tokens"`
	// The event stream during which the rule will be evaluated.
	EventStream                   param.Field[EventStream] `json:"event_stream"`
	ExcludedAccountTokens         param.Field[interface{}] `json:"excluded_account_tokens"`
	ExcludedBusinessAccountTokens param.Field[interface{}] `json:"excluded_business_account_tokens"`
	ExcludedCardTokens            param.Field[interface{}] `json:"excluded_card_tokens"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level"`
}

func (r AuthRuleV2NewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBody) implementsAuthRuleV2NewParamsBodyUnion() {}

// Satisfied by [AuthRuleV2NewParamsBodyAccountLevelRule],
// [AuthRuleV2NewParamsBodyCardLevelRule],
// [AuthRuleV2NewParamsBodyProgramLevelRule], [AuthRuleV2NewParamsBody].
type AuthRuleV2NewParamsBodyUnion interface {
	implementsAuthRuleV2NewParamsBodyUnion()
}

type AuthRuleV2NewParamsBodyAccountLevelRule struct {
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyAccountLevelRuleParametersUnion] `json:"parameters" api:"required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
	//     AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type param.Field[AuthRuleV2NewParamsBodyAccountLevelRuleType] `json:"type" api:"required"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens param.Field[[]string] `json:"account_tokens" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens param.Field[[]string] `json:"business_account_tokens" format:"uuid"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[EventStream] `json:"event_stream"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
}

func (r AuthRuleV2NewParamsBodyAccountLevelRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyAccountLevelRule) implementsAuthRuleV2NewParamsBodyUnion() {}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyAccountLevelRuleParameters struct {
	Action param.Field[interface{}] `json:"action"`
	// The TypeScript source code of the rule. Must define a `rule()` function that
	// accepts the declared features as positional arguments (in the same order as the
	// `features` array) and returns an array of actions.
	Code       param.Field[string]                    `json:"code"`
	Conditions param.Field[interface{}]               `json:"conditions"`
	Features   param.Field[interface{}]               `json:"features"`
	Filters    param.Field[VelocityLimitFiltersParam] `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[int64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[int64]       `json:"limit_count"`
	Merchants  param.Field[interface{}] `json:"merchants"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period param.Field[VelocityLimitPeriodUnionParam] `json:"period"`
	// The scope the velocity is calculated for
	Scope param.Field[AuthRuleV2NewParamsBodyAccountLevelRuleParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyAccountLevelRuleParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyAccountLevelRuleParameters) implementsAuthRuleV2NewParamsBodyAccountLevelRuleParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters],
// [ConditionalTokenizationActionParameters], [TypescriptCodeParameters],
// [AuthRuleV2NewParamsBodyAccountLevelRuleParameters].
type AuthRuleV2NewParamsBodyAccountLevelRuleParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyAccountLevelRuleParametersUnion()
}

// The scope the velocity is calculated for
type AuthRuleV2NewParamsBodyAccountLevelRuleParametersScope string

const (
	AuthRuleV2NewParamsBodyAccountLevelRuleParametersScopeCard    AuthRuleV2NewParamsBodyAccountLevelRuleParametersScope = "CARD"
	AuthRuleV2NewParamsBodyAccountLevelRuleParametersScopeAccount AuthRuleV2NewParamsBodyAccountLevelRuleParametersScope = "ACCOUNT"
)

func (r AuthRuleV2NewParamsBodyAccountLevelRuleParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyAccountLevelRuleParametersScopeCard, AuthRuleV2NewParamsBodyAccountLevelRuleParametersScopeAccount:
		return true
	}
	return false
}

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
//     AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
type AuthRuleV2NewParamsBodyAccountLevelRuleType string

const (
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeConditionalBlock  AuthRuleV2NewParamsBodyAccountLevelRuleType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeVelocityLimit     AuthRuleV2NewParamsBodyAccountLevelRuleType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeMerchantLock      AuthRuleV2NewParamsBodyAccountLevelRuleType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeConditionalAction AuthRuleV2NewParamsBodyAccountLevelRuleType = "CONDITIONAL_ACTION"
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeTypescriptCode    AuthRuleV2NewParamsBodyAccountLevelRuleType = "TYPESCRIPT_CODE"
)

func (r AuthRuleV2NewParamsBodyAccountLevelRuleType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyAccountLevelRuleTypeConditionalBlock, AuthRuleV2NewParamsBodyAccountLevelRuleTypeVelocityLimit, AuthRuleV2NewParamsBodyAccountLevelRuleTypeMerchantLock, AuthRuleV2NewParamsBodyAccountLevelRuleTypeConditionalAction, AuthRuleV2NewParamsBodyAccountLevelRuleTypeTypescriptCode:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCardLevelRule struct {
	// Card tokens to which the Auth Rule applies.
	CardTokens param.Field[[]string] `json:"card_tokens" api:"required" format:"uuid"`
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCardLevelRuleParametersUnion] `json:"parameters" api:"required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
	//     AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type param.Field[AuthRuleV2NewParamsBodyCardLevelRuleType] `json:"type" api:"required"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[EventStream] `json:"event_stream"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
}

func (r AuthRuleV2NewParamsBodyCardLevelRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCardLevelRule) implementsAuthRuleV2NewParamsBodyUnion() {}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyCardLevelRuleParameters struct {
	Action param.Field[interface{}] `json:"action"`
	// The TypeScript source code of the rule. Must define a `rule()` function that
	// accepts the declared features as positional arguments (in the same order as the
	// `features` array) and returns an array of actions.
	Code       param.Field[string]                    `json:"code"`
	Conditions param.Field[interface{}]               `json:"conditions"`
	Features   param.Field[interface{}]               `json:"features"`
	Filters    param.Field[VelocityLimitFiltersParam] `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[int64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[int64]       `json:"limit_count"`
	Merchants  param.Field[interface{}] `json:"merchants"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period param.Field[VelocityLimitPeriodUnionParam] `json:"period"`
	// The scope the velocity is calculated for
	Scope param.Field[AuthRuleV2NewParamsBodyCardLevelRuleParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyCardLevelRuleParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCardLevelRuleParameters) implementsAuthRuleV2NewParamsBodyCardLevelRuleParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters],
// [ConditionalTokenizationActionParameters], [TypescriptCodeParameters],
// [AuthRuleV2NewParamsBodyCardLevelRuleParameters].
type AuthRuleV2NewParamsBodyCardLevelRuleParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCardLevelRuleParametersUnion()
}

// The scope the velocity is calculated for
type AuthRuleV2NewParamsBodyCardLevelRuleParametersScope string

const (
	AuthRuleV2NewParamsBodyCardLevelRuleParametersScopeCard    AuthRuleV2NewParamsBodyCardLevelRuleParametersScope = "CARD"
	AuthRuleV2NewParamsBodyCardLevelRuleParametersScopeAccount AuthRuleV2NewParamsBodyCardLevelRuleParametersScope = "ACCOUNT"
)

func (r AuthRuleV2NewParamsBodyCardLevelRuleParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCardLevelRuleParametersScopeCard, AuthRuleV2NewParamsBodyCardLevelRuleParametersScopeAccount:
		return true
	}
	return false
}

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
//     AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
type AuthRuleV2NewParamsBodyCardLevelRuleType string

const (
	AuthRuleV2NewParamsBodyCardLevelRuleTypeConditionalBlock  AuthRuleV2NewParamsBodyCardLevelRuleType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCardLevelRuleTypeVelocityLimit     AuthRuleV2NewParamsBodyCardLevelRuleType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyCardLevelRuleTypeMerchantLock      AuthRuleV2NewParamsBodyCardLevelRuleType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyCardLevelRuleTypeConditionalAction AuthRuleV2NewParamsBodyCardLevelRuleType = "CONDITIONAL_ACTION"
	AuthRuleV2NewParamsBodyCardLevelRuleTypeTypescriptCode    AuthRuleV2NewParamsBodyCardLevelRuleType = "TYPESCRIPT_CODE"
)

func (r AuthRuleV2NewParamsBodyCardLevelRuleType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCardLevelRuleTypeConditionalBlock, AuthRuleV2NewParamsBodyCardLevelRuleTypeVelocityLimit, AuthRuleV2NewParamsBodyCardLevelRuleTypeMerchantLock, AuthRuleV2NewParamsBodyCardLevelRuleTypeConditionalAction, AuthRuleV2NewParamsBodyCardLevelRuleTypeTypescriptCode:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyProgramLevelRule struct {
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyProgramLevelRuleParametersUnion] `json:"parameters" api:"required"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level" api:"required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
	//     AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type param.Field[AuthRuleV2NewParamsBodyProgramLevelRuleType] `json:"type" api:"required"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[EventStream] `json:"event_stream"`
	// Account tokens to which the Auth Rule does not apply.
	ExcludedAccountTokens param.Field[[]string] `json:"excluded_account_tokens" format:"uuid"`
	// Business account tokens to which the Auth Rule does not apply.
	ExcludedBusinessAccountTokens param.Field[[]string] `json:"excluded_business_account_tokens" format:"uuid"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens param.Field[[]string] `json:"excluded_card_tokens" format:"uuid"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
}

func (r AuthRuleV2NewParamsBodyProgramLevelRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyProgramLevelRule) implementsAuthRuleV2NewParamsBodyUnion() {}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyProgramLevelRuleParameters struct {
	Action param.Field[interface{}] `json:"action"`
	// The TypeScript source code of the rule. Must define a `rule()` function that
	// accepts the declared features as positional arguments (in the same order as the
	// `features` array) and returns an array of actions.
	Code       param.Field[string]                    `json:"code"`
	Conditions param.Field[interface{}]               `json:"conditions"`
	Features   param.Field[interface{}]               `json:"features"`
	Filters    param.Field[VelocityLimitFiltersParam] `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[int64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[int64]       `json:"limit_count"`
	Merchants  param.Field[interface{}] `json:"merchants"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period param.Field[VelocityLimitPeriodUnionParam] `json:"period"`
	// The scope the velocity is calculated for
	Scope param.Field[AuthRuleV2NewParamsBodyProgramLevelRuleParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyProgramLevelRuleParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyProgramLevelRuleParameters) implementsAuthRuleV2NewParamsBodyProgramLevelRuleParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters],
// [ConditionalTokenizationActionParameters], [TypescriptCodeParameters],
// [AuthRuleV2NewParamsBodyProgramLevelRuleParameters].
type AuthRuleV2NewParamsBodyProgramLevelRuleParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyProgramLevelRuleParametersUnion()
}

// The scope the velocity is calculated for
type AuthRuleV2NewParamsBodyProgramLevelRuleParametersScope string

const (
	AuthRuleV2NewParamsBodyProgramLevelRuleParametersScopeCard    AuthRuleV2NewParamsBodyProgramLevelRuleParametersScope = "CARD"
	AuthRuleV2NewParamsBodyProgramLevelRuleParametersScopeAccount AuthRuleV2NewParamsBodyProgramLevelRuleParametersScope = "ACCOUNT"
)

func (r AuthRuleV2NewParamsBodyProgramLevelRuleParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyProgramLevelRuleParametersScopeCard, AuthRuleV2NewParamsBodyProgramLevelRuleParametersScopeAccount:
		return true
	}
	return false
}

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
//     AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
type AuthRuleV2NewParamsBodyProgramLevelRuleType string

const (
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeConditionalBlock  AuthRuleV2NewParamsBodyProgramLevelRuleType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeVelocityLimit     AuthRuleV2NewParamsBodyProgramLevelRuleType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeMerchantLock      AuthRuleV2NewParamsBodyProgramLevelRuleType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeConditionalAction AuthRuleV2NewParamsBodyProgramLevelRuleType = "CONDITIONAL_ACTION"
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeTypescriptCode    AuthRuleV2NewParamsBodyProgramLevelRuleType = "TYPESCRIPT_CODE"
)

func (r AuthRuleV2NewParamsBodyProgramLevelRuleType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyProgramLevelRuleTypeConditionalBlock, AuthRuleV2NewParamsBodyProgramLevelRuleTypeVelocityLimit, AuthRuleV2NewParamsBodyProgramLevelRuleTypeMerchantLock, AuthRuleV2NewParamsBodyProgramLevelRuleTypeConditionalAction, AuthRuleV2NewParamsBodyProgramLevelRuleTypeTypescriptCode:
		return true
	}
	return false
}

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
//   - `CONDITIONAL_BLOCK`: Deprecated. Use `CONDITIONAL_ACTION` instead.
//     AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
//   - `TYPESCRIPT_CODE`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
type AuthRuleV2NewParamsBodyType string

const (
	AuthRuleV2NewParamsBodyTypeConditionalBlock  AuthRuleV2NewParamsBodyType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyTypeVelocityLimit     AuthRuleV2NewParamsBodyType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyTypeMerchantLock      AuthRuleV2NewParamsBodyType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyTypeConditionalAction AuthRuleV2NewParamsBodyType = "CONDITIONAL_ACTION"
	AuthRuleV2NewParamsBodyTypeTypescriptCode    AuthRuleV2NewParamsBodyType = "TYPESCRIPT_CODE"
)

func (r AuthRuleV2NewParamsBodyType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyTypeConditionalBlock, AuthRuleV2NewParamsBodyTypeVelocityLimit, AuthRuleV2NewParamsBodyTypeMerchantLock, AuthRuleV2NewParamsBodyTypeConditionalAction, AuthRuleV2NewParamsBodyTypeTypescriptCode:
		return true
	}
	return false
}

type AuthRuleV2UpdateParams struct {
	Body AuthRuleV2UpdateParamsBodyUnion `json:"body" api:"required"`
}

func (r AuthRuleV2UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AuthRuleV2UpdateParamsBody struct {
	AccountTokens                 param.Field[interface{}] `json:"account_tokens"`
	BusinessAccountTokens         param.Field[interface{}] `json:"business_account_tokens"`
	CardTokens                    param.Field[interface{}] `json:"card_tokens"`
	ExcludedAccountTokens         param.Field[interface{}] `json:"excluded_account_tokens"`
	ExcludedBusinessAccountTokens param.Field[interface{}] `json:"excluded_business_account_tokens"`
	ExcludedCardTokens            param.Field[interface{}] `json:"excluded_card_tokens"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level"`
	// The desired state of the Auth Rule.
	//
	// Note that only deactivating an Auth Rule through this endpoint is supported at
	// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
	// should be used to promote a draft to the currently active version.
	State param.Field[AuthRuleV2UpdateParamsBodyState] `json:"state"`
}

func (r AuthRuleV2UpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2UpdateParamsBody) implementsAuthRuleV2UpdateParamsBodyUnion() {}

// Satisfied by [AuthRuleV2UpdateParamsBodyAccountLevelRule],
// [AuthRuleV2UpdateParamsBodyCardLevelRule],
// [AuthRuleV2UpdateParamsBodyProgramLevelRule], [AuthRuleV2UpdateParamsBody].
type AuthRuleV2UpdateParamsBodyUnion interface {
	implementsAuthRuleV2UpdateParamsBodyUnion()
}

type AuthRuleV2UpdateParamsBodyAccountLevelRule struct {
	// Account tokens to which the Auth Rule applies.
	AccountTokens param.Field[[]string] `json:"account_tokens" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens param.Field[[]string] `json:"business_account_tokens" format:"uuid"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// The desired state of the Auth Rule.
	//
	// Note that only deactivating an Auth Rule through this endpoint is supported at
	// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
	// should be used to promote a draft to the currently active version.
	State param.Field[AuthRuleV2UpdateParamsBodyAccountLevelRuleState] `json:"state"`
}

func (r AuthRuleV2UpdateParamsBodyAccountLevelRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2UpdateParamsBodyAccountLevelRule) implementsAuthRuleV2UpdateParamsBodyUnion() {}

// The desired state of the Auth Rule.
//
// Note that only deactivating an Auth Rule through this endpoint is supported at
// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
// should be used to promote a draft to the currently active version.
type AuthRuleV2UpdateParamsBodyAccountLevelRuleState string

const (
	AuthRuleV2UpdateParamsBodyAccountLevelRuleStateInactive AuthRuleV2UpdateParamsBodyAccountLevelRuleState = "INACTIVE"
)

func (r AuthRuleV2UpdateParamsBodyAccountLevelRuleState) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateParamsBodyAccountLevelRuleStateInactive:
		return true
	}
	return false
}

type AuthRuleV2UpdateParamsBodyCardLevelRule struct {
	// Card tokens to which the Auth Rule applies.
	CardTokens param.Field[[]string] `json:"card_tokens" format:"uuid"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// The desired state of the Auth Rule.
	//
	// Note that only deactivating an Auth Rule through this endpoint is supported at
	// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
	// should be used to promote a draft to the currently active version.
	State param.Field[AuthRuleV2UpdateParamsBodyCardLevelRuleState] `json:"state"`
}

func (r AuthRuleV2UpdateParamsBodyCardLevelRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2UpdateParamsBodyCardLevelRule) implementsAuthRuleV2UpdateParamsBodyUnion() {}

// The desired state of the Auth Rule.
//
// Note that only deactivating an Auth Rule through this endpoint is supported at
// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
// should be used to promote a draft to the currently active version.
type AuthRuleV2UpdateParamsBodyCardLevelRuleState string

const (
	AuthRuleV2UpdateParamsBodyCardLevelRuleStateInactive AuthRuleV2UpdateParamsBodyCardLevelRuleState = "INACTIVE"
)

func (r AuthRuleV2UpdateParamsBodyCardLevelRuleState) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateParamsBodyCardLevelRuleStateInactive:
		return true
	}
	return false
}

type AuthRuleV2UpdateParamsBodyProgramLevelRule struct {
	// Account tokens to which the Auth Rule does not apply.
	ExcludedAccountTokens param.Field[[]string] `json:"excluded_account_tokens" format:"uuid"`
	// Business account tokens to which the Auth Rule does not apply.
	ExcludedBusinessAccountTokens param.Field[[]string] `json:"excluded_business_account_tokens" format:"uuid"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens param.Field[[]string] `json:"excluded_card_tokens" format:"uuid"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level"`
	// The desired state of the Auth Rule.
	//
	// Note that only deactivating an Auth Rule through this endpoint is supported at
	// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
	// should be used to promote a draft to the currently active version.
	State param.Field[AuthRuleV2UpdateParamsBodyProgramLevelRuleState] `json:"state"`
}

func (r AuthRuleV2UpdateParamsBodyProgramLevelRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2UpdateParamsBodyProgramLevelRule) implementsAuthRuleV2UpdateParamsBodyUnion() {}

// The desired state of the Auth Rule.
//
// Note that only deactivating an Auth Rule through this endpoint is supported at
// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
// should be used to promote a draft to the currently active version.
type AuthRuleV2UpdateParamsBodyProgramLevelRuleState string

const (
	AuthRuleV2UpdateParamsBodyProgramLevelRuleStateInactive AuthRuleV2UpdateParamsBodyProgramLevelRuleState = "INACTIVE"
)

func (r AuthRuleV2UpdateParamsBodyProgramLevelRuleState) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateParamsBodyProgramLevelRuleStateInactive:
		return true
	}
	return false
}

// The desired state of the Auth Rule.
//
// Note that only deactivating an Auth Rule through this endpoint is supported at
// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
// should be used to promote a draft to the currently active version.
type AuthRuleV2UpdateParamsBodyState string

const (
	AuthRuleV2UpdateParamsBodyStateInactive AuthRuleV2UpdateParamsBodyState = "INACTIVE"
)

func (r AuthRuleV2UpdateParamsBodyState) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateParamsBodyStateInactive:
		return true
	}
	return false
}

type AuthRuleV2ListParams struct {
	// Only return Auth Rules that are bound to the provided account token.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Only return Auth Rules that are bound to the provided business account token.
	BusinessAccountToken param.Field[string] `query:"business_account_token" format:"uuid"`
	// Only return Auth Rules that are bound to the provided card token.
	CardToken param.Field[string] `query:"card_token" format:"uuid"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Deprecated: Use event_streams instead. Only return Auth rules that are executed
	// during the provided event stream.
	EventStream param.Field[EventStream] `query:"event_stream"`
	// Only return Auth rules that are executed during any of the provided event
	// streams. If event_streams and event_stream are specified, the values will be
	// combined.
	EventStreams param.Field[[]EventStream] `query:"event_streams"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// Only return Auth Rules that are bound to the provided scope.
	Scope param.Field[AuthRuleV2ListParamsScope] `query:"scope"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [AuthRuleV2ListParams]'s query parameters as `url.Values`.
func (r AuthRuleV2ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only return Auth Rules that are bound to the provided scope.
type AuthRuleV2ListParamsScope string

const (
	AuthRuleV2ListParamsScopeProgram         AuthRuleV2ListParamsScope = "PROGRAM"
	AuthRuleV2ListParamsScopeAccount         AuthRuleV2ListParamsScope = "ACCOUNT"
	AuthRuleV2ListParamsScopeBusinessAccount AuthRuleV2ListParamsScope = "BUSINESS_ACCOUNT"
	AuthRuleV2ListParamsScopeCard            AuthRuleV2ListParamsScope = "CARD"
	AuthRuleV2ListParamsScopeAny             AuthRuleV2ListParamsScope = "ANY"
)

func (r AuthRuleV2ListParamsScope) IsKnown() bool {
	switch r {
	case AuthRuleV2ListParamsScopeProgram, AuthRuleV2ListParamsScopeAccount, AuthRuleV2ListParamsScopeBusinessAccount, AuthRuleV2ListParamsScopeCard, AuthRuleV2ListParamsScopeAny:
		return true
	}
	return false
}

type AuthRuleV2DraftParams struct {
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2DraftParamsParametersUnion] `json:"parameters"`
}

func (r AuthRuleV2DraftParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Parameters for the Auth Rule
type AuthRuleV2DraftParamsParameters struct {
	Action param.Field[interface{}] `json:"action"`
	// The TypeScript source code of the rule. Must define a `rule()` function that
	// accepts the declared features as positional arguments (in the same order as the
	// `features` array) and returns an array of actions.
	Code       param.Field[string]                    `json:"code"`
	Conditions param.Field[interface{}]               `json:"conditions"`
	Features   param.Field[interface{}]               `json:"features"`
	Filters    param.Field[VelocityLimitFiltersParam] `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[int64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[int64]       `json:"limit_count"`
	Merchants  param.Field[interface{}] `json:"merchants"`
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period param.Field[VelocityLimitPeriodUnionParam] `json:"period"`
	// The scope the velocity is calculated for
	Scope param.Field[AuthRuleV2DraftParamsParametersScope] `json:"scope"`
}

func (r AuthRuleV2DraftParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2DraftParamsParameters) implementsAuthRuleV2DraftParamsParametersUnion() {}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters],
// [ConditionalTokenizationActionParameters], [TypescriptCodeParameters],
// [AuthRuleV2DraftParamsParameters].
type AuthRuleV2DraftParamsParametersUnion interface {
	implementsAuthRuleV2DraftParamsParametersUnion()
}

// The scope the velocity is calculated for
type AuthRuleV2DraftParamsParametersScope string

const (
	AuthRuleV2DraftParamsParametersScopeCard    AuthRuleV2DraftParamsParametersScope = "CARD"
	AuthRuleV2DraftParamsParametersScopeAccount AuthRuleV2DraftParamsParametersScope = "ACCOUNT"
)

func (r AuthRuleV2DraftParamsParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersScopeCard, AuthRuleV2DraftParamsParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2ListResultsParams struct {
	// Filter by Auth Rule token
	AuthRuleToken param.Field[string] `query:"auth_rule_token" format:"uuid"`
	// Date string in RFC 3339 format. Only events evaluated after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only events evaluated before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter by event token
	EventToken param.Field[string] `query:"event_token" format:"uuid"`
	// Filter by whether the rule evaluation produced any actions. When not provided,
	// all results are returned.
	HasActions param.Field[bool] `query:"has_actions"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [AuthRuleV2ListResultsParams]'s query parameters as
// `url.Values`.
func (r AuthRuleV2ListResultsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthRuleV2GetFeaturesParams struct {
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	CardToken    param.Field[string] `query:"card_token" format:"uuid"`
}

// URLQuery serializes [AuthRuleV2GetFeaturesParams]'s query parameters as
// `url.Values`.
func (r AuthRuleV2GetFeaturesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthRuleV2GetReportParams struct {
	// Start date for the report
	Begin param.Field[time.Time] `query:"begin" api:"required" format:"date"`
	// End date for the report
	End param.Field[time.Time] `query:"end" api:"required" format:"date"`
}

// URLQuery serializes [AuthRuleV2GetReportParams]'s query parameters as
// `url.Values`.
func (r AuthRuleV2GetReportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
