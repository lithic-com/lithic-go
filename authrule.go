// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/lithic-com/lithic-go/shared"
	"github.com/tidwall/gjson"
)

// AuthRuleService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthRuleService] method instead.
type AuthRuleService struct {
	Options []option.RequestOption
	V2      *AuthRuleV2Service
}

// NewAuthRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthRuleService(opts ...option.RequestOption) (r *AuthRuleService) {
	r = &AuthRuleService{}
	r.Options = opts
	r.V2 = NewAuthRuleV2Service(opts...)
	return
}

// Creates an authorization rule (Auth Rule) and applies it at the program,
// account, or card level.
func (r *AuthRuleService) New(ctx context.Context, body AuthRuleNewParams, opts ...option.RequestOption) (res *shared.AuthRule, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/auth_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detail the properties and entities (program, accounts, and cards) associated
// with an existing authorization rule (Auth Rule).
func (r *AuthRuleService) Get(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v1/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the properties associated with an existing authorization rule (Auth
// Rule).
func (r *AuthRuleService) Update(ctx context.Context, authRuleToken string, body AuthRuleUpdateParams, opts ...option.RequestOption) (res *shared.AuthRule, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v1/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Return all of the Auth Rules under the program.
func (r *AuthRuleService) List(ctx context.Context, query AuthRuleListParams, opts ...option.RequestOption) (res *pagination.CursorPage[shared.AuthRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/auth_rules"
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
func (r *AuthRuleService) ListAutoPaging(ctx context.Context, query AuthRuleListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[shared.AuthRule] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Applies an existing authorization rule (Auth Rule) to an program, account, or
// card level.
func (r *AuthRuleService) Apply(ctx context.Context, authRuleToken string, body AuthRuleApplyParams, opts ...option.RequestOption) (res *shared.AuthRule, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v1/auth_rules/%s/apply", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Migrates an existing V1 authorization rule to a V2 authorization rule. Depending
// on the configuration of the V1 Auth Rule, this will yield one or two V2
// authorization rules. This endpoint will alter the internal structure of the Auth
// Rule such that the resulting rules become a V2 Authorization Rule that can be
// operated on through the /v2/auth_rules endpoints.
//
// After a V1 Auth Rule has been migrated, it can no longer be operated on through
// the /v1/auth_rules/\* endpoints. Eventually, Lithic will deprecate the
// /v1/auth_rules endpoints and migrate all existing V1 Auth Rules to V2 Auth
// Rules.
func (r *AuthRuleService) MigrateV1ToV2(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *[]AuthRuleMigrateV1ToV2Response, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v1/auth_rules/%s/migrate", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Remove an existing authorization rule (Auth Rule) from an program, account, or
// card-level.
func (r *AuthRuleService) Remove(ctx context.Context, body AuthRuleRemoveParams, opts ...option.RequestOption) (res *AuthRuleRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/auth_rules/remove"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AuthRuleGetResponse struct {
	Data []shared.AuthRule       `json:"data"`
	JSON authRuleGetResponseJSON `json:"-"`
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

func (r authRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleMigrateV1ToV2Response struct {
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                                    `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleMigrateV1ToV2ResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleMigrateV1ToV2ResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleMigrateV1ToV2ResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleMigrateV1ToV2ResponseType `json:"type,required"`
	JSON authRuleMigrateV1ToV2ResponseJSON `json:"-"`
}

// authRuleMigrateV1ToV2ResponseJSON contains the JSON metadata for the struct
// [AuthRuleMigrateV1ToV2Response]
type authRuleMigrateV1ToV2ResponseJSON struct {
	Token          apijson.Field
	AccountTokens  apijson.Field
	CardTokens     apijson.Field
	CurrentVersion apijson.Field
	DraftVersion   apijson.Field
	ProgramLevel   apijson.Field
	State          apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AuthRuleMigrateV1ToV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleMigrateV1ToV2ResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleMigrateV1ToV2ResponseCurrentVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleMigrateV1ToV2ResponseCurrentVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                           `json:"version,required"`
	JSON    authRuleMigrateV1ToV2ResponseCurrentVersionJSON `json:"-"`
}

// authRuleMigrateV1ToV2ResponseCurrentVersionJSON contains the JSON metadata for
// the struct [AuthRuleMigrateV1ToV2ResponseCurrentVersion]
type authRuleMigrateV1ToV2ResponseCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleMigrateV1ToV2ResponseCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleMigrateV1ToV2ResponseCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleMigrateV1ToV2ResponseCurrentVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                                `json:"conditions,required"`
	Scope      AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScope `json:"scope"`
	// This field can have the runtime type of [shared.VelocityLimitParamsPeriodUnion].
	Period interface{} `json:"period,required"`
	// This field can have the runtime type of [shared.VelocityLimitParamsFilters].
	Filters interface{} `json:"filters,required"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount float64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount float64                                                   `json:"limit_count,nullable"`
	JSON       authRuleMigrateV1ToV2ResponseCurrentVersionParametersJSON `json:"-"`
	union      AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersUnion
}

// authRuleMigrateV1ToV2ResponseCurrentVersionParametersJSON contains the JSON
// metadata for the struct [AuthRuleMigrateV1ToV2ResponseCurrentVersionParameters]
type authRuleMigrateV1ToV2ResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleMigrateV1ToV2ResponseCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleMigrateV1ToV2ResponseCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleMigrateV1ToV2ResponseCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleMigrateV1ToV2ResponseCurrentVersionParameters) AsUnion() AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParameters]
// or [shared.VelocityLimitParams].
type AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersUnion interface {
	ImplementsAuthRuleMigrateV1ToV2ResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParameters]
type authRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParameters) ImplementsAuthRuleMigrateV1ToV2ResponseCurrentVersionParameters() {
}

type AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target
	Attribute AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersCondition]
type authRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
	return r.raw
}

// The attribute to target
type AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleMigrateV1ToV2ResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScope string

const (
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScopeCard    AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScope = "CARD"
	AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScopeAccount AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScopeCard, AuthRuleMigrateV1ToV2ResponseCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleMigrateV1ToV2ResponseDraftVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleMigrateV1ToV2ResponseDraftVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                         `json:"version,required"`
	JSON    authRuleMigrateV1ToV2ResponseDraftVersionJSON `json:"-"`
}

// authRuleMigrateV1ToV2ResponseDraftVersionJSON contains the JSON metadata for the
// struct [AuthRuleMigrateV1ToV2ResponseDraftVersion]
type authRuleMigrateV1ToV2ResponseDraftVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleMigrateV1ToV2ResponseDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleMigrateV1ToV2ResponseDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleMigrateV1ToV2ResponseDraftVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                              `json:"conditions,required"`
	Scope      AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScope `json:"scope"`
	// This field can have the runtime type of [shared.VelocityLimitParamsPeriodUnion].
	Period interface{} `json:"period,required"`
	// This field can have the runtime type of [shared.VelocityLimitParamsFilters].
	Filters interface{} `json:"filters,required"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount float64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount float64                                                 `json:"limit_count,nullable"`
	JSON       authRuleMigrateV1ToV2ResponseDraftVersionParametersJSON `json:"-"`
	union      AuthRuleMigrateV1ToV2ResponseDraftVersionParametersUnion
}

// authRuleMigrateV1ToV2ResponseDraftVersionParametersJSON contains the JSON
// metadata for the struct [AuthRuleMigrateV1ToV2ResponseDraftVersionParameters]
type authRuleMigrateV1ToV2ResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleMigrateV1ToV2ResponseDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleMigrateV1ToV2ResponseDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleMigrateV1ToV2ResponseDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleMigrateV1ToV2ResponseDraftVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleMigrateV1ToV2ResponseDraftVersionParameters) AsUnion() AuthRuleMigrateV1ToV2ResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParameters]
// or [shared.VelocityLimitParams].
type AuthRuleMigrateV1ToV2ResponseDraftVersionParametersUnion interface {
	ImplementsAuthRuleMigrateV1ToV2ResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleMigrateV1ToV2ResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParameters]
type authRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParameters) ImplementsAuthRuleMigrateV1ToV2ResponseDraftVersionParameters() {
}

type AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target
	Attribute AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersCondition]
type authRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
	return r.raw
}

// The attribute to target
type AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleMigrateV1ToV2ResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScope string

const (
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScopeCard    AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScope = "CARD"
	AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScopeAccount AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScopeCard, AuthRuleMigrateV1ToV2ResponseDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleMigrateV1ToV2ResponseState string

const (
	AuthRuleMigrateV1ToV2ResponseStateActive   AuthRuleMigrateV1ToV2ResponseState = "ACTIVE"
	AuthRuleMigrateV1ToV2ResponseStateInactive AuthRuleMigrateV1ToV2ResponseState = "INACTIVE"
)

func (r AuthRuleMigrateV1ToV2ResponseState) IsKnown() bool {
	switch r {
	case AuthRuleMigrateV1ToV2ResponseStateActive, AuthRuleMigrateV1ToV2ResponseStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleMigrateV1ToV2ResponseType string

const (
	AuthRuleMigrateV1ToV2ResponseTypeConditionalBlock AuthRuleMigrateV1ToV2ResponseType = "CONDITIONAL_BLOCK"
	AuthRuleMigrateV1ToV2ResponseTypeVelocityLimit    AuthRuleMigrateV1ToV2ResponseType = "VELOCITY_LIMIT"
)

func (r AuthRuleMigrateV1ToV2ResponseType) IsKnown() bool {
	switch r {
	case AuthRuleMigrateV1ToV2ResponseTypeConditionalBlock, AuthRuleMigrateV1ToV2ResponseTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleRemoveResponse struct {
	AccountTokens []string                   `json:"account_tokens"`
	CardTokens    []string                   `json:"card_tokens"`
	ProgramLevel  bool                       `json:"program_level"`
	JSON          authRuleRemoveResponseJSON `json:"-"`
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

func (r authRuleRemoveResponseJSON) RawJSON() string {
	return r.raw
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
