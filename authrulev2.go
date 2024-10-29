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

// AuthRuleV2Service contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthRuleV2Service] method instead.
type AuthRuleV2Service struct {
	Options []option.RequestOption
}

// NewAuthRuleV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthRuleV2Service(opts ...option.RequestOption) (r *AuthRuleV2Service) {
	r = &AuthRuleV2Service{}
	r.Options = opts
	return
}

// Creates a new V2 authorization rule in draft mode
func (r *AuthRuleV2Service) New(ctx context.Context, body AuthRuleV2NewParams, opts ...option.RequestOption) (res *AuthRuleV2NewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/auth_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches an authorization rule by its token
func (r *AuthRuleV2Service) Get(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleV2GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an authorization rule's properties
func (r *AuthRuleV2Service) Update(ctx context.Context, authRuleToken string, body AuthRuleV2UpdateParams, opts ...option.RequestOption) (res *AuthRuleV2UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists V2 authorization rules
func (r *AuthRuleV2Service) List(ctx context.Context, query AuthRuleV2ListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AuthRuleV2ListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// Lists V2 authorization rules
func (r *AuthRuleV2Service) ListAutoPaging(ctx context.Context, query AuthRuleV2ListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AuthRuleV2ListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Associates an authorization rules with a card program, the provided account(s)
// or card(s).
//
// This endpoint will replace any existing associations with the provided list of
// entities.
func (r *AuthRuleV2Service) Apply(ctx context.Context, authRuleToken string, body AuthRuleV2ApplyParams, opts ...option.RequestOption) (res *AuthRuleV2ApplyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/apply", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a new draft version of an authorization rules that will be ran in shadow
// mode.
//
// This can also be utilized to reset the draft parameters, causing a draft version
// to no longer be ran in shadow mode.
func (r *AuthRuleV2Service) Draft(ctx context.Context, authRuleToken string, body AuthRuleV2DraftParams, opts ...option.RequestOption) (res *AuthRuleV2DraftResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/draft", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Promotes a draft version of an authorization rule to the currently active
// version such that it is enforced in the authorization stream.
func (r *AuthRuleV2Service) Promote(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleV2PromoteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/promote", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Requests a performance report of an authorization rule to be asynchronously
// generated. Reports can only be run on rules in draft or active mode and will
// included approved and declined statistics as well as examples. The generated
// report will be delivered asynchronously through a webhook with `event_type` =
// `auth_rules.performance_report.created`. See the docs on setting up
// [webhook subscriptions](https://docs.lithic.com/docs/events-api).
//
// Reports are generated based on data collected by Lithic's authorization
// processing system in the trailing week. The performance of the auth rule will be
// assessed on the configuration of the auth rule at the time the report is
// requested. This implies that if a performance report is requested, right after
// updating an auth rule, depending on the number of authorizations processed for a
// card program, it may be the case that no data is available for the report.
// Therefore Lithic recommends to decouple making updates to an Auth Rule, and
// requesting performance reports.
//
// To make this concrete, consider the following example:
//
//  1. At time `t`, a new Auth Rule is created, and applies to all authorizations on
//     a card program. The Auth Rule has not yet been promoted, causing the draft
//     version of the rule to be applied in shadow mode.
//  2. At time `t + 1 hour` a performance report is requested for the Auth Rule.
//     This performance report will _only_ contain data for the Auth Rule being
//     executed in the window between `t` and `t + 1 hour`. This is because Lithic's
//     transaction processing system will only start capturing data for the Auth
//     Rule at the time it is created.
//  3. At time `t + 2 hours` the draft version of the Auth Rule is promoted to the
//     active version of the Auth Rule by calling the
//     `/v2/auth_rules/{auth_rule_token}/promote` endpoint. If a performance report
//     is requested at this moment it will still only contain data for this version
//     of the rule, but the window of available data will now span from `t` to
//     `t + 2 hours`.
//  4. At time `t + 3 hours` a new version of the rule is drafted by calling the
//     `/v2/auth_rules/{auth_rule_token}/draft` endpoint. If a performance report is
//     requested right at this moment, it will only contain data for authorizations
//     to which both the active version and the draft version is applied. Lithic
//     does this to ensure that performance reports represent a fair comparison
//     between rules. Because there may be no authorizations in this window, and
//     because there may be some lag before data is available in a performance
//     report, the requested performance report could contain no to little data.
//  5. At time `t + 4 hours` another performance report is requested: this time the
//     performance report will contain data from the window between `t + 3 hours`
//     and `t + 4 hours`, for any authorizations to which both the current version
//     of the authorization rule (in enforcing mode) and the draft version of the
//     authorization rule (in shadow mode) applied.
//
// Note that generating a report may take up to 15 minutes and that delivery is not
// guaranteed. Customers are required to have created an event subscription to
// receive the webhook. Additionally, there is a delay of approximately 15 minutes
// between when Lithic's transaction processing systems have processed the
// transaction, and when a transaction will be included in the report.
func (r *AuthRuleV2Service) Report(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleV2ReportResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/report", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AuthRuleV2NewResponse struct {
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                            `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2NewResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2NewResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2NewResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2NewResponseType `json:"type,required"`
	JSON authRuleV2NewResponseJSON `json:"-"`
}

// authRuleV2NewResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2NewResponse]
type authRuleV2NewResponseJSON struct {
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

func (r *AuthRuleV2NewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2NewResponseCurrentVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2NewResponseCurrentVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                   `json:"version,required"`
	JSON    authRuleV2NewResponseCurrentVersionJSON `json:"-"`
}

// authRuleV2NewResponseCurrentVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2NewResponseCurrentVersion]
type authRuleV2NewResponseCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2NewResponseCurrentVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                        `json:"conditions,required"`
	Scope      AuthRuleV2NewResponseCurrentVersionParametersScope `json:"scope"`
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
	LimitCount float64                                           `json:"limit_count,nullable"`
	JSON       authRuleV2NewResponseCurrentVersionParametersJSON `json:"-"`
	union      AuthRuleV2NewResponseCurrentVersionParametersUnion
}

// authRuleV2NewResponseCurrentVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2NewResponseCurrentVersionParameters]
type authRuleV2NewResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2NewResponseCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2NewResponseCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2NewResponseCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2NewResponseCurrentVersionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2NewResponseCurrentVersionParameters) AsUnion() AuthRuleV2NewResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2NewResponseCurrentVersionParametersUnion interface {
	ImplementsAuthRuleV2NewResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParameters]
type authRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2NewResponseCurrentVersionParameters() {
}

type AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersCondition]
type authRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2NewResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2NewResponseCurrentVersionParametersScope string

const (
	AuthRuleV2NewResponseCurrentVersionParametersScopeCard    AuthRuleV2NewResponseCurrentVersionParametersScope = "CARD"
	AuthRuleV2NewResponseCurrentVersionParametersScopeAccount AuthRuleV2NewResponseCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersScopeCard, AuthRuleV2NewResponseCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2NewResponseDraftVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2NewResponseDraftVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                 `json:"version,required"`
	JSON    authRuleV2NewResponseDraftVersionJSON `json:"-"`
}

// authRuleV2NewResponseDraftVersionJSON contains the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersion]
type authRuleV2NewResponseDraftVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2NewResponseDraftVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                      `json:"conditions,required"`
	Scope      AuthRuleV2NewResponseDraftVersionParametersScope `json:"scope"`
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
	LimitCount float64                                         `json:"limit_count,nullable"`
	JSON       authRuleV2NewResponseDraftVersionParametersJSON `json:"-"`
	union      AuthRuleV2NewResponseDraftVersionParametersUnion
}

// authRuleV2NewResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2NewResponseDraftVersionParameters]
type authRuleV2NewResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2NewResponseDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2NewResponseDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2NewResponseDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2NewResponseDraftVersionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2NewResponseDraftVersionParameters) AsUnion() AuthRuleV2NewResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2NewResponseDraftVersionParametersUnion interface {
	ImplementsAuthRuleV2NewResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2NewResponseDraftVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2NewResponseDraftVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParameters]
type authRuleV2NewResponseDraftVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2NewResponseDraftVersionParameters() {
}

type AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersCondition]
type authRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2NewResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2NewResponseDraftVersionParametersScope string

const (
	AuthRuleV2NewResponseDraftVersionParametersScopeCard    AuthRuleV2NewResponseDraftVersionParametersScope = "CARD"
	AuthRuleV2NewResponseDraftVersionParametersScopeAccount AuthRuleV2NewResponseDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2NewResponseDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersScopeCard, AuthRuleV2NewResponseDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2NewResponseState string

const (
	AuthRuleV2NewResponseStateActive   AuthRuleV2NewResponseState = "ACTIVE"
	AuthRuleV2NewResponseStateInactive AuthRuleV2NewResponseState = "INACTIVE"
)

func (r AuthRuleV2NewResponseState) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseStateActive, AuthRuleV2NewResponseStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2NewResponseType string

const (
	AuthRuleV2NewResponseTypeConditionalBlock AuthRuleV2NewResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewResponseTypeVelocityLimit    AuthRuleV2NewResponseType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2NewResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseTypeConditionalBlock, AuthRuleV2NewResponseTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2GetResponse struct {
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                            `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2GetResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2GetResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2GetResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2GetResponseType `json:"type,required"`
	JSON authRuleV2GetResponseJSON `json:"-"`
}

// authRuleV2GetResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2GetResponse]
type authRuleV2GetResponseJSON struct {
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

func (r *AuthRuleV2GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetResponseCurrentVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2GetResponseCurrentVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                   `json:"version,required"`
	JSON    authRuleV2GetResponseCurrentVersionJSON `json:"-"`
}

// authRuleV2GetResponseCurrentVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2GetResponseCurrentVersion]
type authRuleV2GetResponseCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2GetResponseCurrentVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                        `json:"conditions,required"`
	Scope      AuthRuleV2GetResponseCurrentVersionParametersScope `json:"scope"`
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
	LimitCount float64                                           `json:"limit_count,nullable"`
	JSON       authRuleV2GetResponseCurrentVersionParametersJSON `json:"-"`
	union      AuthRuleV2GetResponseCurrentVersionParametersUnion
}

// authRuleV2GetResponseCurrentVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2GetResponseCurrentVersionParameters]
type authRuleV2GetResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2GetResponseCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2GetResponseCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2GetResponseCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2GetResponseCurrentVersionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2GetResponseCurrentVersionParameters) AsUnion() AuthRuleV2GetResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2GetResponseCurrentVersionParametersUnion interface {
	ImplementsAuthRuleV2GetResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParameters]
type authRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2GetResponseCurrentVersionParameters() {
}

type AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersCondition]
type authRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2GetResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2GetResponseCurrentVersionParametersScope string

const (
	AuthRuleV2GetResponseCurrentVersionParametersScopeCard    AuthRuleV2GetResponseCurrentVersionParametersScope = "CARD"
	AuthRuleV2GetResponseCurrentVersionParametersScopeAccount AuthRuleV2GetResponseCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersScopeCard, AuthRuleV2GetResponseCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2GetResponseDraftVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2GetResponseDraftVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                 `json:"version,required"`
	JSON    authRuleV2GetResponseDraftVersionJSON `json:"-"`
}

// authRuleV2GetResponseDraftVersionJSON contains the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersion]
type authRuleV2GetResponseDraftVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2GetResponseDraftVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                      `json:"conditions,required"`
	Scope      AuthRuleV2GetResponseDraftVersionParametersScope `json:"scope"`
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
	LimitCount float64                                         `json:"limit_count,nullable"`
	JSON       authRuleV2GetResponseDraftVersionParametersJSON `json:"-"`
	union      AuthRuleV2GetResponseDraftVersionParametersUnion
}

// authRuleV2GetResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2GetResponseDraftVersionParameters]
type authRuleV2GetResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2GetResponseDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2GetResponseDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2GetResponseDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2GetResponseDraftVersionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2GetResponseDraftVersionParameters) AsUnion() AuthRuleV2GetResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2GetResponseDraftVersionParametersUnion interface {
	ImplementsAuthRuleV2GetResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2GetResponseDraftVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2GetResponseDraftVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParameters]
type authRuleV2GetResponseDraftVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2GetResponseDraftVersionParameters() {
}

type AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersCondition]
type authRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2GetResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2GetResponseDraftVersionParametersScope string

const (
	AuthRuleV2GetResponseDraftVersionParametersScopeCard    AuthRuleV2GetResponseDraftVersionParametersScope = "CARD"
	AuthRuleV2GetResponseDraftVersionParametersScopeAccount AuthRuleV2GetResponseDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2GetResponseDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersScopeCard, AuthRuleV2GetResponseDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2GetResponseState string

const (
	AuthRuleV2GetResponseStateActive   AuthRuleV2GetResponseState = "ACTIVE"
	AuthRuleV2GetResponseStateInactive AuthRuleV2GetResponseState = "INACTIVE"
)

func (r AuthRuleV2GetResponseState) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseStateActive, AuthRuleV2GetResponseStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2GetResponseType string

const (
	AuthRuleV2GetResponseTypeConditionalBlock AuthRuleV2GetResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2GetResponseTypeVelocityLimit    AuthRuleV2GetResponseType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2GetResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseTypeConditionalBlock, AuthRuleV2GetResponseTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2UpdateResponse struct {
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                               `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2UpdateResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2UpdateResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2UpdateResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2UpdateResponseType `json:"type,required"`
	JSON authRuleV2UpdateResponseJSON `json:"-"`
}

// authRuleV2UpdateResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponse]
type authRuleV2UpdateResponseJSON struct {
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

func (r *AuthRuleV2UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2UpdateResponseCurrentVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2UpdateResponseCurrentVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                      `json:"version,required"`
	JSON    authRuleV2UpdateResponseCurrentVersionJSON `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2UpdateResponseCurrentVersion]
type authRuleV2UpdateResponseCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2UpdateResponseCurrentVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                           `json:"conditions,required"`
	Scope      AuthRuleV2UpdateResponseCurrentVersionParametersScope `json:"scope"`
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
	LimitCount float64                                              `json:"limit_count,nullable"`
	JSON       authRuleV2UpdateResponseCurrentVersionParametersJSON `json:"-"`
	union      AuthRuleV2UpdateResponseCurrentVersionParametersUnion
}

// authRuleV2UpdateResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2UpdateResponseCurrentVersionParameters]
type authRuleV2UpdateResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2UpdateResponseCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2UpdateResponseCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2UpdateResponseCurrentVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2UpdateResponseCurrentVersionParameters) AsUnion() AuthRuleV2UpdateResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2UpdateResponseCurrentVersionParametersUnion interface {
	ImplementsAuthRuleV2UpdateResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParameters]
type authRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2UpdateResponseCurrentVersionParameters() {
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersCondition]
type authRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2UpdateResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2UpdateResponseCurrentVersionParametersScope string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersScopeCard    AuthRuleV2UpdateResponseCurrentVersionParametersScope = "CARD"
	AuthRuleV2UpdateResponseCurrentVersionParametersScopeAccount AuthRuleV2UpdateResponseCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersScopeCard, AuthRuleV2UpdateResponseCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2UpdateResponseDraftVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2UpdateResponseDraftVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                    `json:"version,required"`
	JSON    authRuleV2UpdateResponseDraftVersionJSON `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2UpdateResponseDraftVersion]
type authRuleV2UpdateResponseDraftVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2UpdateResponseDraftVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                         `json:"conditions,required"`
	Scope      AuthRuleV2UpdateResponseDraftVersionParametersScope `json:"scope"`
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
	LimitCount float64                                            `json:"limit_count,nullable"`
	JSON       authRuleV2UpdateResponseDraftVersionParametersJSON `json:"-"`
	union      AuthRuleV2UpdateResponseDraftVersionParametersUnion
}

// authRuleV2UpdateResponseDraftVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2UpdateResponseDraftVersionParameters]
type authRuleV2UpdateResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2UpdateResponseDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2UpdateResponseDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2UpdateResponseDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2UpdateResponseDraftVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2UpdateResponseDraftVersionParameters) AsUnion() AuthRuleV2UpdateResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2UpdateResponseDraftVersionParametersUnion interface {
	ImplementsAuthRuleV2UpdateResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParameters]
type authRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2UpdateResponseDraftVersionParameters() {
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersCondition]
type authRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2UpdateResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2UpdateResponseDraftVersionParametersScope string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersScopeCard    AuthRuleV2UpdateResponseDraftVersionParametersScope = "CARD"
	AuthRuleV2UpdateResponseDraftVersionParametersScopeAccount AuthRuleV2UpdateResponseDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersScopeCard, AuthRuleV2UpdateResponseDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2UpdateResponseState string

const (
	AuthRuleV2UpdateResponseStateActive   AuthRuleV2UpdateResponseState = "ACTIVE"
	AuthRuleV2UpdateResponseStateInactive AuthRuleV2UpdateResponseState = "INACTIVE"
)

func (r AuthRuleV2UpdateResponseState) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseStateActive, AuthRuleV2UpdateResponseStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2UpdateResponseType string

const (
	AuthRuleV2UpdateResponseTypeConditionalBlock AuthRuleV2UpdateResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2UpdateResponseTypeVelocityLimit    AuthRuleV2UpdateResponseType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2UpdateResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseTypeConditionalBlock, AuthRuleV2UpdateResponseTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2ListResponse struct {
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                             `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2ListResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2ListResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2ListResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2ListResponseType `json:"type,required"`
	JSON authRuleV2ListResponseJSON `json:"-"`
}

// authRuleV2ListResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ListResponse]
type authRuleV2ListResponseJSON struct {
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

func (r *AuthRuleV2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ListResponseCurrentVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2ListResponseCurrentVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                    `json:"version,required"`
	JSON    authRuleV2ListResponseCurrentVersionJSON `json:"-"`
}

// authRuleV2ListResponseCurrentVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2ListResponseCurrentVersion]
type authRuleV2ListResponseCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2ListResponseCurrentVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                         `json:"conditions,required"`
	Scope      AuthRuleV2ListResponseCurrentVersionParametersScope `json:"scope"`
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
	LimitCount float64                                            `json:"limit_count,nullable"`
	JSON       authRuleV2ListResponseCurrentVersionParametersJSON `json:"-"`
	union      AuthRuleV2ListResponseCurrentVersionParametersUnion
}

// authRuleV2ListResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2ListResponseCurrentVersionParameters]
type authRuleV2ListResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2ListResponseCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2ListResponseCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2ListResponseCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2ListResponseCurrentVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2ListResponseCurrentVersionParameters) AsUnion() AuthRuleV2ListResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2ListResponseCurrentVersionParametersUnion interface {
	ImplementsAuthRuleV2ListResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParameters]
type authRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2ListResponseCurrentVersionParameters() {
}

type AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersCondition]
type authRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2ListResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2ListResponseCurrentVersionParametersScope string

const (
	AuthRuleV2ListResponseCurrentVersionParametersScopeCard    AuthRuleV2ListResponseCurrentVersionParametersScope = "CARD"
	AuthRuleV2ListResponseCurrentVersionParametersScopeAccount AuthRuleV2ListResponseCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersScopeCard, AuthRuleV2ListResponseCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2ListResponseDraftVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2ListResponseDraftVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                  `json:"version,required"`
	JSON    authRuleV2ListResponseDraftVersionJSON `json:"-"`
}

// authRuleV2ListResponseDraftVersionJSON contains the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersion]
type authRuleV2ListResponseDraftVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2ListResponseDraftVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                       `json:"conditions,required"`
	Scope      AuthRuleV2ListResponseDraftVersionParametersScope `json:"scope"`
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
	LimitCount float64                                          `json:"limit_count,nullable"`
	JSON       authRuleV2ListResponseDraftVersionParametersJSON `json:"-"`
	union      AuthRuleV2ListResponseDraftVersionParametersUnion
}

// authRuleV2ListResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2ListResponseDraftVersionParameters]
type authRuleV2ListResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2ListResponseDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2ListResponseDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2ListResponseDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2ListResponseDraftVersionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2ListResponseDraftVersionParameters) AsUnion() AuthRuleV2ListResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2ListResponseDraftVersionParametersUnion interface {
	ImplementsAuthRuleV2ListResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ListResponseDraftVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2ListResponseDraftVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParameters]
type authRuleV2ListResponseDraftVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2ListResponseDraftVersionParameters() {
}

type AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersCondition]
type authRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2ListResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2ListResponseDraftVersionParametersScope string

const (
	AuthRuleV2ListResponseDraftVersionParametersScopeCard    AuthRuleV2ListResponseDraftVersionParametersScope = "CARD"
	AuthRuleV2ListResponseDraftVersionParametersScopeAccount AuthRuleV2ListResponseDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2ListResponseDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersScopeCard, AuthRuleV2ListResponseDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2ListResponseState string

const (
	AuthRuleV2ListResponseStateActive   AuthRuleV2ListResponseState = "ACTIVE"
	AuthRuleV2ListResponseStateInactive AuthRuleV2ListResponseState = "INACTIVE"
)

func (r AuthRuleV2ListResponseState) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseStateActive, AuthRuleV2ListResponseStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2ListResponseType string

const (
	AuthRuleV2ListResponseTypeConditionalBlock AuthRuleV2ListResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2ListResponseTypeVelocityLimit    AuthRuleV2ListResponseType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2ListResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseTypeConditionalBlock, AuthRuleV2ListResponseTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2ApplyResponse struct {
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                              `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2ApplyResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2ApplyResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2ApplyResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2ApplyResponseType `json:"type,required"`
	JSON authRuleV2ApplyResponseJSON `json:"-"`
}

// authRuleV2ApplyResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponse]
type authRuleV2ApplyResponseJSON struct {
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

func (r *AuthRuleV2ApplyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ApplyResponseCurrentVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2ApplyResponseCurrentVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                     `json:"version,required"`
	JSON    authRuleV2ApplyResponseCurrentVersionJSON `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2ApplyResponseCurrentVersion]
type authRuleV2ApplyResponseCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2ApplyResponseCurrentVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                          `json:"conditions,required"`
	Scope      AuthRuleV2ApplyResponseCurrentVersionParametersScope `json:"scope"`
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
	LimitCount float64                                             `json:"limit_count,nullable"`
	JSON       authRuleV2ApplyResponseCurrentVersionParametersJSON `json:"-"`
	union      AuthRuleV2ApplyResponseCurrentVersionParametersUnion
}

// authRuleV2ApplyResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2ApplyResponseCurrentVersionParameters]
type authRuleV2ApplyResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2ApplyResponseCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2ApplyResponseCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2ApplyResponseCurrentVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2ApplyResponseCurrentVersionParameters) AsUnion() AuthRuleV2ApplyResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2ApplyResponseCurrentVersionParametersUnion interface {
	ImplementsAuthRuleV2ApplyResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParameters]
type authRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2ApplyResponseCurrentVersionParameters() {
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersCondition]
type authRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2ApplyResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2ApplyResponseCurrentVersionParametersScope string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersScopeCard    AuthRuleV2ApplyResponseCurrentVersionParametersScope = "CARD"
	AuthRuleV2ApplyResponseCurrentVersionParametersScopeAccount AuthRuleV2ApplyResponseCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersScopeCard, AuthRuleV2ApplyResponseCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2ApplyResponseDraftVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2ApplyResponseDraftVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                   `json:"version,required"`
	JSON    authRuleV2ApplyResponseDraftVersionJSON `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2ApplyResponseDraftVersion]
type authRuleV2ApplyResponseDraftVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2ApplyResponseDraftVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                        `json:"conditions,required"`
	Scope      AuthRuleV2ApplyResponseDraftVersionParametersScope `json:"scope"`
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
	LimitCount float64                                           `json:"limit_count,nullable"`
	JSON       authRuleV2ApplyResponseDraftVersionParametersJSON `json:"-"`
	union      AuthRuleV2ApplyResponseDraftVersionParametersUnion
}

// authRuleV2ApplyResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2ApplyResponseDraftVersionParameters]
type authRuleV2ApplyResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2ApplyResponseDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2ApplyResponseDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2ApplyResponseDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2ApplyResponseDraftVersionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2ApplyResponseDraftVersionParameters) AsUnion() AuthRuleV2ApplyResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2ApplyResponseDraftVersionParametersUnion interface {
	ImplementsAuthRuleV2ApplyResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParameters]
type authRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2ApplyResponseDraftVersionParameters() {
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersCondition]
type authRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2ApplyResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2ApplyResponseDraftVersionParametersScope string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersScopeCard    AuthRuleV2ApplyResponseDraftVersionParametersScope = "CARD"
	AuthRuleV2ApplyResponseDraftVersionParametersScopeAccount AuthRuleV2ApplyResponseDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersScopeCard, AuthRuleV2ApplyResponseDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2ApplyResponseState string

const (
	AuthRuleV2ApplyResponseStateActive   AuthRuleV2ApplyResponseState = "ACTIVE"
	AuthRuleV2ApplyResponseStateInactive AuthRuleV2ApplyResponseState = "INACTIVE"
)

func (r AuthRuleV2ApplyResponseState) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseStateActive, AuthRuleV2ApplyResponseStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2ApplyResponseType string

const (
	AuthRuleV2ApplyResponseTypeConditionalBlock AuthRuleV2ApplyResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2ApplyResponseTypeVelocityLimit    AuthRuleV2ApplyResponseType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2ApplyResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseTypeConditionalBlock, AuthRuleV2ApplyResponseTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2DraftResponse struct {
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                              `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2DraftResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2DraftResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2DraftResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2DraftResponseType `json:"type,required"`
	JSON authRuleV2DraftResponseJSON `json:"-"`
}

// authRuleV2DraftResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2DraftResponse]
type authRuleV2DraftResponseJSON struct {
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

func (r *AuthRuleV2DraftResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2DraftResponseCurrentVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2DraftResponseCurrentVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                     `json:"version,required"`
	JSON    authRuleV2DraftResponseCurrentVersionJSON `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2DraftResponseCurrentVersion]
type authRuleV2DraftResponseCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2DraftResponseCurrentVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                          `json:"conditions,required"`
	Scope      AuthRuleV2DraftResponseCurrentVersionParametersScope `json:"scope"`
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
	LimitCount float64                                             `json:"limit_count,nullable"`
	JSON       authRuleV2DraftResponseCurrentVersionParametersJSON `json:"-"`
	union      AuthRuleV2DraftResponseCurrentVersionParametersUnion
}

// authRuleV2DraftResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2DraftResponseCurrentVersionParameters]
type authRuleV2DraftResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2DraftResponseCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2DraftResponseCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2DraftResponseCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2DraftResponseCurrentVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2DraftResponseCurrentVersionParameters) AsUnion() AuthRuleV2DraftResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2DraftResponseCurrentVersionParametersUnion interface {
	ImplementsAuthRuleV2DraftResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParameters]
type authRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2DraftResponseCurrentVersionParameters() {
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersCondition]
type authRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2DraftResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2DraftResponseCurrentVersionParametersScope string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersScopeCard    AuthRuleV2DraftResponseCurrentVersionParametersScope = "CARD"
	AuthRuleV2DraftResponseCurrentVersionParametersScopeAccount AuthRuleV2DraftResponseCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersScopeCard, AuthRuleV2DraftResponseCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2DraftResponseDraftVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2DraftResponseDraftVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                   `json:"version,required"`
	JSON    authRuleV2DraftResponseDraftVersionJSON `json:"-"`
}

// authRuleV2DraftResponseDraftVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2DraftResponseDraftVersion]
type authRuleV2DraftResponseDraftVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2DraftResponseDraftVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                        `json:"conditions,required"`
	Scope      AuthRuleV2DraftResponseDraftVersionParametersScope `json:"scope"`
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
	LimitCount float64                                           `json:"limit_count,nullable"`
	JSON       authRuleV2DraftResponseDraftVersionParametersJSON `json:"-"`
	union      AuthRuleV2DraftResponseDraftVersionParametersUnion
}

// authRuleV2DraftResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2DraftResponseDraftVersionParameters]
type authRuleV2DraftResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2DraftResponseDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2DraftResponseDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2DraftResponseDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2DraftResponseDraftVersionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2DraftResponseDraftVersionParameters) AsUnion() AuthRuleV2DraftResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2DraftResponseDraftVersionParametersUnion interface {
	ImplementsAuthRuleV2DraftResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParameters]
type authRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2DraftResponseDraftVersionParameters() {
}

type AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersCondition]
type authRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2DraftResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2DraftResponseDraftVersionParametersScope string

const (
	AuthRuleV2DraftResponseDraftVersionParametersScopeCard    AuthRuleV2DraftResponseDraftVersionParametersScope = "CARD"
	AuthRuleV2DraftResponseDraftVersionParametersScopeAccount AuthRuleV2DraftResponseDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersScopeCard, AuthRuleV2DraftResponseDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2DraftResponseState string

const (
	AuthRuleV2DraftResponseStateActive   AuthRuleV2DraftResponseState = "ACTIVE"
	AuthRuleV2DraftResponseStateInactive AuthRuleV2DraftResponseState = "INACTIVE"
)

func (r AuthRuleV2DraftResponseState) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseStateActive, AuthRuleV2DraftResponseStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2DraftResponseType string

const (
	AuthRuleV2DraftResponseTypeConditionalBlock AuthRuleV2DraftResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2DraftResponseTypeVelocityLimit    AuthRuleV2DraftResponseType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2DraftResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseTypeConditionalBlock, AuthRuleV2DraftResponseTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2PromoteResponse struct {
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                                `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2PromoteResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2PromoteResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2PromoteResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2PromoteResponseType `json:"type,required"`
	JSON authRuleV2PromoteResponseJSON `json:"-"`
}

// authRuleV2PromoteResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponse]
type authRuleV2PromoteResponseJSON struct {
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

func (r *AuthRuleV2PromoteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2PromoteResponseCurrentVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2PromoteResponseCurrentVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                       `json:"version,required"`
	JSON    authRuleV2PromoteResponseCurrentVersionJSON `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2PromoteResponseCurrentVersion]
type authRuleV2PromoteResponseCurrentVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2PromoteResponseCurrentVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                            `json:"conditions,required"`
	Scope      AuthRuleV2PromoteResponseCurrentVersionParametersScope `json:"scope"`
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
	LimitCount float64                                               `json:"limit_count,nullable"`
	JSON       authRuleV2PromoteResponseCurrentVersionParametersJSON `json:"-"`
	union      AuthRuleV2PromoteResponseCurrentVersionParametersUnion
}

// authRuleV2PromoteResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2PromoteResponseCurrentVersionParameters]
type authRuleV2PromoteResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2PromoteResponseCurrentVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2PromoteResponseCurrentVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2PromoteResponseCurrentVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2PromoteResponseCurrentVersionParameters) AsUnion() AuthRuleV2PromoteResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2PromoteResponseCurrentVersionParametersUnion interface {
	ImplementsAuthRuleV2PromoteResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParameters]
type authRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2PromoteResponseCurrentVersionParameters() {
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersCondition]
type authRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2PromoteResponseCurrentVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2PromoteResponseCurrentVersionParametersScope string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersScopeCard    AuthRuleV2PromoteResponseCurrentVersionParametersScope = "CARD"
	AuthRuleV2PromoteResponseCurrentVersionParametersScopeAccount AuthRuleV2PromoteResponseCurrentVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersScopeCard, AuthRuleV2PromoteResponseCurrentVersionParametersScopeAccount:
		return true
	}
	return false
}

type AuthRuleV2PromoteResponseDraftVersion struct {
	// Parameters for the current version of the Auth Rule
	Parameters AuthRuleV2PromoteResponseDraftVersionParameters `json:"parameters,required"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                     `json:"version,required"`
	JSON    authRuleV2PromoteResponseDraftVersionJSON `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionJSON contains the JSON metadata for the
// struct [AuthRuleV2PromoteResponseDraftVersion]
type authRuleV2PromoteResponseDraftVersionJSON struct {
	Parameters  apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionJSON) RawJSON() string {
	return r.raw
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2PromoteResponseDraftVersionParameters struct {
	// This field can have the runtime type of
	// [[]AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersCondition].
	Conditions interface{}                                          `json:"conditions,required"`
	Scope      AuthRuleV2PromoteResponseDraftVersionParametersScope `json:"scope"`
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
	LimitCount float64                                             `json:"limit_count,nullable"`
	JSON       authRuleV2PromoteResponseDraftVersionParametersJSON `json:"-"`
	union      AuthRuleV2PromoteResponseDraftVersionParametersUnion
}

// authRuleV2PromoteResponseDraftVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2PromoteResponseDraftVersionParameters]
type authRuleV2PromoteResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Scope       apijson.Field
	Period      apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r authRuleV2PromoteResponseDraftVersionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *AuthRuleV2PromoteResponseDraftVersionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = AuthRuleV2PromoteResponseDraftVersionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthRuleV2PromoteResponseDraftVersionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParameters],
// [shared.VelocityLimitParams].
func (r AuthRuleV2PromoteResponseDraftVersionParameters) AsUnion() AuthRuleV2PromoteResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the current version of the Auth Rule
//
// Union satisfied by
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParameters] or
// [shared.VelocityLimitParams].
type AuthRuleV2PromoteResponseDraftVersionParametersUnion interface {
	ImplementsAuthRuleV2PromoteResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.VelocityLimitParams{}),
		},
	)
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParameters struct {
	Conditions []AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersCondition `json:"conditions,required"`
	JSON       authRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersJSON        `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParameters]
type authRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParameters) ImplementsAuthRuleV2PromoteResponseDraftVersionParameters() {
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionJSON        `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersCondition]
type authRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionJSON) RawJSON() string {
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray{}),
		},
	)
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2PromoteResponseDraftVersionParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2PromoteResponseDraftVersionParametersScope string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersScopeCard    AuthRuleV2PromoteResponseDraftVersionParametersScope = "CARD"
	AuthRuleV2PromoteResponseDraftVersionParametersScopeAccount AuthRuleV2PromoteResponseDraftVersionParametersScope = "ACCOUNT"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersScopeCard, AuthRuleV2PromoteResponseDraftVersionParametersScopeAccount:
		return true
	}
	return false
}

// The state of the Auth Rule
type AuthRuleV2PromoteResponseState string

const (
	AuthRuleV2PromoteResponseStateActive   AuthRuleV2PromoteResponseState = "ACTIVE"
	AuthRuleV2PromoteResponseStateInactive AuthRuleV2PromoteResponseState = "INACTIVE"
)

func (r AuthRuleV2PromoteResponseState) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseStateActive, AuthRuleV2PromoteResponseStateInactive:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2PromoteResponseType string

const (
	AuthRuleV2PromoteResponseTypeConditionalBlock AuthRuleV2PromoteResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2PromoteResponseTypeVelocityLimit    AuthRuleV2PromoteResponseType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2PromoteResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseTypeConditionalBlock, AuthRuleV2PromoteResponseTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2ReportResponse struct {
	ReportToken string                       `json:"report_token" format:"uuid"`
	JSON        authRuleV2ReportResponseJSON `json:"-"`
}

// authRuleV2ReportResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ReportResponse]
type authRuleV2ReportResponseJSON struct {
	ReportToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ReportResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2NewParams struct {
	Body AuthRuleV2NewParamsBodyUnion `json:"body,required"`
}

func (r AuthRuleV2NewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AuthRuleV2NewParamsBody struct {
	AccountTokens param.Field[interface{}] `json:"account_tokens,required"`
	// The type of Auth Rule
	Type       param.Field[AuthRuleV2NewParamsBodyType] `json:"type"`
	Parameters param.Field[interface{}]                 `json:"parameters,required"`
	CardTokens param.Field[interface{}]                 `json:"card_tokens,required"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level"`
}

func (r AuthRuleV2NewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBody) implementsAuthRuleV2NewParamsBodyUnion() {}

// Satisfied by [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel],
// [AuthRuleV2NewParamsBody].
type AuthRuleV2NewParamsBodyUnion interface {
	implementsAuthRuleV2NewParamsBodyUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens struct {
	// Account tokens to which the Auth Rule applies.
	AccountTokens param.Field[[]string] `json:"account_tokens,required" format:"uuid"`
	// Parameters for the current version of the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion] `json:"parameters"`
	// The type of Auth Rule
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens) implementsAuthRuleV2NewParamsBodyUnion() {
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters struct {
	Conditions param.Field[interface{}]                                                              `json:"conditions,required"`
	Scope      param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScope] `json:"scope"`
	Period     param.Field[interface{}]                                                              `json:"period,required"`
	Filters    param.Field[interface{}]                                                              `json:"filters,required"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[float64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[float64] `json:"limit_count"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion() {
}

// Parameters for the current version of the Auth Rule
//
// Satisfied by
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParameters],
// [shared.VelocityLimitParams],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParameters struct {
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParameters) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion() {
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionFloat],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScope string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScopeCard    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScope = "CARD"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScopeAccount AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScope = "ACCOUNT"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScopeCard, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScopeAccount:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalBlock AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeVelocityLimit    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens struct {
	// Card tokens to which the Auth Rule applies.
	CardTokens param.Field[[]string] `json:"card_tokens,required" format:"uuid"`
	// Parameters for the current version of the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion] `json:"parameters"`
	// The type of Auth Rule
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens) implementsAuthRuleV2NewParamsBodyUnion() {
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters struct {
	Conditions param.Field[interface{}]                                                           `json:"conditions,required"`
	Scope      param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScope] `json:"scope"`
	Period     param.Field[interface{}]                                                           `json:"period,required"`
	Filters    param.Field[interface{}]                                                           `json:"filters,required"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[float64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[float64] `json:"limit_count"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion() {
}

// Parameters for the current version of the Auth Rule
//
// Satisfied by
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParameters],
// [shared.VelocityLimitParams],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParameters struct {
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParameters) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion() {
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionFloat],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScope string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScopeCard    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScope = "CARD"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScopeAccount AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScope = "ACCOUNT"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScopeCard, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScopeAccount:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditionalBlock AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeVelocityLimit    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel struct {
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level,required"`
	// Parameters for the current version of the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion] `json:"parameters"`
	// The type of Auth Rule
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel) implementsAuthRuleV2NewParamsBodyUnion() {
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters struct {
	Conditions param.Field[interface{}]                                                             `json:"conditions,required"`
	Scope      param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScope] `json:"scope"`
	Period     param.Field[interface{}]                                                             `json:"period,required"`
	Filters    param.Field[interface{}]                                                             `json:"filters,required"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[float64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[float64] `json:"limit_count"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion() {
}

// Parameters for the current version of the Auth Rule
//
// Satisfied by
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParameters],
// [shared.VelocityLimitParams],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParameters struct {
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParameters) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion() {
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionFloat],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalBlockParametersConditionsValueUnion() {
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScope string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScopeCard    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScope = "CARD"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScopeAccount AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScope = "ACCOUNT"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScope) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScopeCard, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScopeAccount:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditionalBlock AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeVelocityLimit    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeVelocityLimit:
		return true
	}
	return false
}

// The type of Auth Rule
type AuthRuleV2NewParamsBodyType string

const (
	AuthRuleV2NewParamsBodyTypeConditionalBlock AuthRuleV2NewParamsBodyType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyTypeVelocityLimit    AuthRuleV2NewParamsBodyType = "VELOCITY_LIMIT"
)

func (r AuthRuleV2NewParamsBodyType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyTypeConditionalBlock, AuthRuleV2NewParamsBodyTypeVelocityLimit:
		return true
	}
	return false
}

type AuthRuleV2UpdateParams struct {
	// The desired state of the Auth Rule.
	//
	// Note that only deactivating an Auth Rule through this endpoint is supported at
	// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
	// should be used to promote a draft to the currently active version.
	State param.Field[AuthRuleV2UpdateParamsState] `json:"state"`
}

func (r AuthRuleV2UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The desired state of the Auth Rule.
//
// Note that only deactivating an Auth Rule through this endpoint is supported at
// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
// should be used to promote a draft to the currently active version.
type AuthRuleV2UpdateParamsState string

const (
	AuthRuleV2UpdateParamsStateInactive AuthRuleV2UpdateParamsState = "INACTIVE"
)

func (r AuthRuleV2UpdateParamsState) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateParamsStateInactive:
		return true
	}
	return false
}

type AuthRuleV2ListParams struct {
	// Only return Authorization Rules that are bound to the provided account token.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Only return Authorization Rules that are bound to the provided card token.
	CardToken param.Field[string] `query:"card_token" format:"uuid"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [AuthRuleV2ListParams]'s query parameters as `url.Values`.
func (r AuthRuleV2ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthRuleV2ApplyParams struct {
	Body AuthRuleV2ApplyParamsBodyUnion `json:"body,required"`
}

func (r AuthRuleV2ApplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AuthRuleV2ApplyParamsBody struct {
	AccountTokens param.Field[interface{}] `json:"account_tokens,required"`
	CardTokens    param.Field[interface{}] `json:"card_tokens,required"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level"`
}

func (r AuthRuleV2ApplyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2ApplyParamsBody) implementsAuthRuleV2ApplyParamsBodyUnion() {}

// Satisfied by [AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestAccountTokens],
// [AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestCardTokens],
// [AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestProgramLevel],
// [AuthRuleV2ApplyParamsBody].
type AuthRuleV2ApplyParamsBodyUnion interface {
	implementsAuthRuleV2ApplyParamsBodyUnion()
}

type AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestAccountTokens struct {
	// Account tokens to which the Auth Rule applies.
	AccountTokens param.Field[[]string] `json:"account_tokens,required" format:"uuid"`
}

func (r AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestAccountTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestAccountTokens) implementsAuthRuleV2ApplyParamsBodyUnion() {
}

type AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestCardTokens struct {
	// Card tokens to which the Auth Rule applies.
	CardTokens param.Field[[]string] `json:"card_tokens,required" format:"uuid"`
}

func (r AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestCardTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestCardTokens) implementsAuthRuleV2ApplyParamsBodyUnion() {
}

type AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestProgramLevel struct {
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level,required"`
}

func (r AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestProgramLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestProgramLevel) implementsAuthRuleV2ApplyParamsBodyUnion() {
}

type AuthRuleV2DraftParams struct {
	// Parameters for the current version of the Auth Rule
	Parameters param.Field[AuthRuleV2DraftParamsParametersUnion] `json:"parameters"`
}

func (r AuthRuleV2DraftParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Parameters for the current version of the Auth Rule
type AuthRuleV2DraftParamsParameters struct {
	Conditions param.Field[interface{}]                          `json:"conditions,required"`
	Scope      param.Field[AuthRuleV2DraftParamsParametersScope] `json:"scope"`
	Period     param.Field[interface{}]                          `json:"period,required"`
	Filters    param.Field[interface{}]                          `json:"filters,required"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[float64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[float64] `json:"limit_count"`
}

func (r AuthRuleV2DraftParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2DraftParamsParameters) ImplementsAuthRuleV2DraftParamsParametersUnion() {}

// Parameters for the current version of the Auth Rule
//
// Satisfied by [AuthRuleV2DraftParamsParametersConditionalBlockParameters],
// [shared.VelocityLimitParams], [AuthRuleV2DraftParamsParameters].
type AuthRuleV2DraftParamsParametersUnion interface {
	ImplementsAuthRuleV2DraftParamsParametersUnion()
}

type AuthRuleV2DraftParamsParametersConditionalBlockParameters struct {
	Conditions param.Field[[]AuthRuleV2DraftParamsParametersConditionalBlockParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2DraftParamsParametersConditionalBlockParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2DraftParamsParametersConditionalBlockParameters) ImplementsAuthRuleV2DraftParamsParametersUnion() {
}

type AuthRuleV2DraftParamsParametersConditionalBlockParametersCondition struct {
	// The attribute to target.
	//
	// The following attributes may be targeted:
	//
	//   - `MCC`: A four-digit number listed in ISO 18245. An MCC is used to classify a
	//     business by the types of goods or services it provides.
	//   - `COUNTRY`: Country of entity of card acceptor. Possible values are: (1) all
	//     ISO 3166-1 alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for
	//     Netherlands Antilles.
	//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
	//     transaction.
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
	Attribute param.Field[AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2DraftParamsParametersConditionalBlockParametersCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
//   - `CURRENCY`: 3-digit alphabetic ISO 4217 code for the merchant currency of the
//     transaction.
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
type AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute string

const (
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeMcc               AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeCountry           AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeCurrency          AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeMerchantID        AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeDescriptor        AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeLiabilityShift    AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributePanEntryMode      AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeTransactionAmount AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeRiskScore         AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute = "RISK_SCORE"
)

func (r AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeMcc, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeCountry, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeCurrency, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeMerchantID, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeDescriptor, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeLiabilityShift, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributePanEntryMode, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeRiskScore:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation string

const (
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsOneOf       AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsNotOneOf    AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationMatches       AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationDoesNotMatch  AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsGreaterThan AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsLessThan    AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsOneOf, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationMatches, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionFloat],
// [AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsValueArray].
type AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsValueUnion()
}

type AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsValueArray []string

func (r AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsValueArray) ImplementsAuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsValueUnion() {
}

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
