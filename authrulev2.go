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

// Creates a new V2 authorization rule in draft mode
func (r *AuthRuleV2Service) New(ctx context.Context, body AuthRuleV2NewParams, opts ...option.RequestOption) (res *AuthRuleV2NewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/auth_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a V2 authorization rule by its token
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

// Updates a V2 authorization rule's properties
//
// If `account_tokens`, `card_tokens`, `program_level`, or `excluded_card_tokens`
// is provided, this will replace existing associations with the provided list of
// entities.
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

// Deletes a V2 authorization rule
func (r *AuthRuleV2Service) Delete(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Associates a V2 authorization rule with a card program, the provided account(s)
// or card(s).
//
// Prefer using the `PATCH` method for this operation.
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

// Creates a new draft version of a rule that will be ran in shadow mode.
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

// Promotes the draft version of an authorization rule to the currently active
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
	//   - `CARD_TRANSACTION_COUNT_1H`: The number of transactions on the card in the
	//     trailing hour up and until the authorization.
	//   - `CARD_TRANSACTION_COUNT_24H`: The number of transactions on the card in the
	//     trailing 24 hours up and until the authorization.
	//   - `CARD_STATE`: The current state of the card associated with the transaction.
	//     Valid values are `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`,
	//     `PENDING_FULFILLMENT`.
	Attribute ConditionalAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleConditionOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleConditionValueUnion `json:"value"`
	JSON  authRuleConditionJSON       `json:"-"`
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

// The operation to apply to the attribute
type AuthRuleConditionOperation string

const (
	AuthRuleConditionOperationIsOneOf       AuthRuleConditionOperation = "IS_ONE_OF"
	AuthRuleConditionOperationIsNotOneOf    AuthRuleConditionOperation = "IS_NOT_ONE_OF"
	AuthRuleConditionOperationMatches       AuthRuleConditionOperation = "MATCHES"
	AuthRuleConditionOperationDoesNotMatch  AuthRuleConditionOperation = "DOES_NOT_MATCH"
	AuthRuleConditionOperationIsGreaterThan AuthRuleConditionOperation = "IS_GREATER_THAN"
	AuthRuleConditionOperationIsLessThan    AuthRuleConditionOperation = "IS_LESS_THAN"
)

func (r AuthRuleConditionOperation) IsKnown() bool {
	switch r {
	case AuthRuleConditionOperationIsOneOf, AuthRuleConditionOperationIsNotOneOf, AuthRuleConditionOperationMatches, AuthRuleConditionOperationDoesNotMatch, AuthRuleConditionOperationIsGreaterThan, AuthRuleConditionOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleConditionValueListOfStrings].
type AuthRuleConditionValueUnion interface {
	ImplementsAuthRuleConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleConditionValueListOfStrings{}),
		},
	)
}

type AuthRuleConditionValueListOfStrings []string

func (r AuthRuleConditionValueListOfStrings) ImplementsAuthRuleConditionValueUnion() {}

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
	//   - `CARD_TRANSACTION_COUNT_1H`: The number of transactions on the card in the
	//     trailing hour up and until the authorization.
	//   - `CARD_TRANSACTION_COUNT_24H`: The number of transactions on the card in the
	//     trailing 24 hours up and until the authorization.
	//   - `CARD_STATE`: The current state of the card associated with the transaction.
	//     Valid values are `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`,
	//     `PENDING_FULFILLMENT`.
	Attribute param.Field[ConditionalAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleConditionOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleConditionValueUnionParam] `json:"value"`
}

func (r AuthRuleConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleConditionValueListOfStringsParam].
type AuthRuleConditionValueUnionParam interface {
	ImplementsAuthRuleConditionValueUnionParam()
}

type AuthRuleConditionValueListOfStringsParam []string

func (r AuthRuleConditionValueListOfStringsParam) ImplementsAuthRuleConditionValueUnionParam() {}

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
//   - `CARD_TRANSACTION_COUNT_1H`: The number of transactions on the card in the
//     trailing hour up and until the authorization.
//   - `CARD_TRANSACTION_COUNT_24H`: The number of transactions on the card in the
//     trailing 24 hours up and until the authorization.
//   - `CARD_STATE`: The current state of the card associated with the transaction.
//     Valid values are `CLOSED`, `OPEN`, `PAUSED`, `PENDING_ACTIVATION`,
//     `PENDING_FULFILLMENT`.
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
	ConditionalAttributeCardTransactionCount1H  ConditionalAttribute = "CARD_TRANSACTION_COUNT_1H"
	ConditionalAttributeCardTransactionCount24H ConditionalAttribute = "CARD_TRANSACTION_COUNT_24H"
	ConditionalAttributeCardState               ConditionalAttribute = "CARD_STATE"
)

func (r ConditionalAttribute) IsKnown() bool {
	switch r {
	case ConditionalAttributeMcc, ConditionalAttributeCountry, ConditionalAttributeCurrency, ConditionalAttributeMerchantID, ConditionalAttributeDescriptor, ConditionalAttributeLiabilityShift, ConditionalAttributePanEntryMode, ConditionalAttributeTransactionAmount, ConditionalAttributeRiskScore, ConditionalAttributeCardTransactionCount1H, ConditionalAttributeCardTransactionCount24H, ConditionalAttributeCardState:
		return true
	}
	return false
}

type ConditionalBlockParameters struct {
	Conditions []AuthRuleCondition            `json:"conditions,required"`
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

func (r ConditionalBlockParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2ApplyResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2ApplyResponseDraftVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {}

type VelocityLimitParams struct {
	Filters VelocityLimitParamsFilters `json:"filters,required"`
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodUnion `json:"period,required"`
	Scope  VelocityLimitParamsScope       `json:"scope,required"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64                   `json:"limit_count,nullable"`
	JSON       velocityLimitParamsJSON `json:"-"`
}

// velocityLimitParamsJSON contains the JSON metadata for the struct
// [VelocityLimitParams]
type velocityLimitParamsJSON struct {
	Filters     apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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

func (r VelocityLimitParams) implementsAuthRuleV2NewResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2NewResponseDraftVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2GetResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2GetResponseDraftVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2ListResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2ListResponseDraftVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2ApplyResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2ApplyResponseDraftVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2DraftResponseDraftVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {}

func (r VelocityLimitParams) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {}

type VelocityLimitParamsFilters struct {
	// ISO-3166-1 alpha-3 Country Codes to include in the velocity calculation.
	// Transactions not matching any of the provided will not be included in the
	// calculated velocity.
	IncludeCountries []string `json:"include_countries,nullable"`
	// Merchant Category Codes to include in the velocity calculation. Transactions not
	// matching this MCC will not be included in the calculated velocity.
	IncludeMccs []string                       `json:"include_mccs,nullable"`
	JSON        velocityLimitParamsFiltersJSON `json:"-"`
}

// velocityLimitParamsFiltersJSON contains the JSON metadata for the struct
// [VelocityLimitParamsFilters]
type velocityLimitParamsFiltersJSON struct {
	IncludeCountries apijson.Field
	IncludeMccs      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VelocityLimitParamsFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitParamsFiltersJSON) RawJSON() string {
	return r.raw
}

// The size of the trailing window to calculate Spend Velocity over in seconds. The
// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
//
// Union satisfied by [shared.UnionInt] or [VelocityLimitParamsPeriodWindow].
type VelocityLimitParamsPeriodUnion interface {
	ImplementsVelocityLimitParamsPeriodUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VelocityLimitParamsPeriodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(VelocityLimitParamsPeriodWindow("")),
		},
	)
}

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

// The window of time to calculate Spend Velocity over.
//
//   - `DAY`: Velocity over the current day since midnight Eastern Time.
//   - `WEEK`: Velocity over the current week since 00:00 / 12 AM on Monday in
//     Eastern Time.
//   - `MONTH`: Velocity over the current month since 00:00 / 12 AM on the first of
//     the month in Eastern Time.
type VelocityLimitParamsPeriodWindow string

const (
	VelocityLimitParamsPeriodWindowDay   VelocityLimitParamsPeriodWindow = "DAY"
	VelocityLimitParamsPeriodWindowWeek  VelocityLimitParamsPeriodWindow = "WEEK"
	VelocityLimitParamsPeriodWindowMonth VelocityLimitParamsPeriodWindow = "MONTH"
)

func (r VelocityLimitParamsPeriodWindow) IsKnown() bool {
	switch r {
	case VelocityLimitParamsPeriodWindowDay, VelocityLimitParamsPeriodWindowWeek, VelocityLimitParamsPeriodWindowMonth:
		return true
	}
	return false
}

func (r VelocityLimitParamsPeriodWindow) ImplementsVelocityLimitParamsPeriodUnion() {}

type AuthRuleV2NewResponse struct {
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                            `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2NewResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2NewResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2NewResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2NewResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                  `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2NewResponseJSON `json:"-"`
}

// authRuleV2NewResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2NewResponse]
type authRuleV2NewResponseJSON struct {
	Token              apijson.Field
	AccountTokens      apijson.Field
	CardTokens         apijson.Field
	CurrentVersion     apijson.Field
	DraftVersion       apijson.Field
	Name               apijson.Field
	ProgramLevel       apijson.Field
	State              apijson.Field
	Type               apijson.Field
	ExcludedCardTokens apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AuthRuleV2NewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2NewResponseCurrentVersion struct {
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2NewResponseCurrentVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                        `json:"period"`
	Scope  AuthRuleV2NewResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2NewResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2NewResponseCurrentVersionParametersUnion
}

// authRuleV2NewResponseCurrentVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2NewResponseCurrentVersionParameters]
type authRuleV2NewResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2NewResponseCurrentVersionParameters) AsUnion() AuthRuleV2NewResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2NewResponseCurrentVersionParametersUnion interface {
	implementsAuthRuleV2NewResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2NewResponseDraftVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                      `json:"period"`
	Scope  AuthRuleV2NewResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2NewResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2NewResponseDraftVersionParametersUnion
}

// authRuleV2NewResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2NewResponseDraftVersionParameters]
type authRuleV2NewResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2NewResponseDraftVersionParameters) AsUnion() AuthRuleV2NewResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2NewResponseDraftVersionParametersUnion interface {
	implementsAuthRuleV2NewResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                            `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2GetResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2GetResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2GetResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2GetResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                  `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2GetResponseJSON `json:"-"`
}

// authRuleV2GetResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2GetResponse]
type authRuleV2GetResponseJSON struct {
	Token              apijson.Field
	AccountTokens      apijson.Field
	CardTokens         apijson.Field
	CurrentVersion     apijson.Field
	DraftVersion       apijson.Field
	Name               apijson.Field
	ProgramLevel       apijson.Field
	State              apijson.Field
	Type               apijson.Field
	ExcludedCardTokens apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AuthRuleV2GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetResponseCurrentVersion struct {
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2GetResponseCurrentVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                        `json:"period"`
	Scope  AuthRuleV2GetResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2GetResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2GetResponseCurrentVersionParametersUnion
}

// authRuleV2GetResponseCurrentVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2GetResponseCurrentVersionParameters]
type authRuleV2GetResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2GetResponseCurrentVersionParameters) AsUnion() AuthRuleV2GetResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2GetResponseCurrentVersionParametersUnion interface {
	implementsAuthRuleV2GetResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2GetResponseDraftVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                      `json:"period"`
	Scope  AuthRuleV2GetResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2GetResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2GetResponseDraftVersionParametersUnion
}

// authRuleV2GetResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2GetResponseDraftVersionParameters]
type authRuleV2GetResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2GetResponseDraftVersionParameters) AsUnion() AuthRuleV2GetResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2GetResponseDraftVersionParametersUnion interface {
	implementsAuthRuleV2GetResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                               `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2UpdateResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2UpdateResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2UpdateResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2UpdateResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                     `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2UpdateResponseJSON `json:"-"`
}

// authRuleV2UpdateResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponse]
type authRuleV2UpdateResponseJSON struct {
	Token              apijson.Field
	AccountTokens      apijson.Field
	CardTokens         apijson.Field
	CurrentVersion     apijson.Field
	DraftVersion       apijson.Field
	Name               apijson.Field
	ProgramLevel       apijson.Field
	State              apijson.Field
	Type               apijson.Field
	ExcludedCardTokens apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2UpdateResponseCurrentVersion struct {
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2UpdateResponseCurrentVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                           `json:"period"`
	Scope  AuthRuleV2UpdateResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2UpdateResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2UpdateResponseCurrentVersionParametersUnion
}

// authRuleV2UpdateResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2UpdateResponseCurrentVersionParameters]
type authRuleV2UpdateResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2UpdateResponseCurrentVersionParameters) AsUnion() AuthRuleV2UpdateResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2UpdateResponseCurrentVersionParametersUnion interface {
	implementsAuthRuleV2UpdateResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2UpdateResponseDraftVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                         `json:"period"`
	Scope  AuthRuleV2UpdateResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2UpdateResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2UpdateResponseDraftVersionParametersUnion
}

// authRuleV2UpdateResponseDraftVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2UpdateResponseDraftVersionParameters]
type authRuleV2UpdateResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2UpdateResponseDraftVersionParameters) AsUnion() AuthRuleV2UpdateResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2UpdateResponseDraftVersionParametersUnion interface {
	implementsAuthRuleV2UpdateResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                             `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2ListResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2ListResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2ListResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2ListResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                   `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2ListResponseJSON `json:"-"`
}

// authRuleV2ListResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ListResponse]
type authRuleV2ListResponseJSON struct {
	Token              apijson.Field
	AccountTokens      apijson.Field
	CardTokens         apijson.Field
	CurrentVersion     apijson.Field
	DraftVersion       apijson.Field
	Name               apijson.Field
	ProgramLevel       apijson.Field
	State              apijson.Field
	Type               apijson.Field
	ExcludedCardTokens apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AuthRuleV2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ListResponseCurrentVersion struct {
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2ListResponseCurrentVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                         `json:"period"`
	Scope  AuthRuleV2ListResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2ListResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2ListResponseCurrentVersionParametersUnion
}

// authRuleV2ListResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2ListResponseCurrentVersionParameters]
type authRuleV2ListResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2ListResponseCurrentVersionParameters) AsUnion() AuthRuleV2ListResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2ListResponseCurrentVersionParametersUnion interface {
	implementsAuthRuleV2ListResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2ListResponseDraftVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                       `json:"period"`
	Scope  AuthRuleV2ListResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2ListResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2ListResponseDraftVersionParametersUnion
}

// authRuleV2ListResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2ListResponseDraftVersionParameters]
type authRuleV2ListResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2ListResponseDraftVersionParameters) AsUnion() AuthRuleV2ListResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2ListResponseDraftVersionParametersUnion interface {
	implementsAuthRuleV2ListResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                              `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2ApplyResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2ApplyResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2ApplyResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2ApplyResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                    `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2ApplyResponseJSON `json:"-"`
}

// authRuleV2ApplyResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponse]
type authRuleV2ApplyResponseJSON struct {
	Token              apijson.Field
	AccountTokens      apijson.Field
	CardTokens         apijson.Field
	CurrentVersion     apijson.Field
	DraftVersion       apijson.Field
	Name               apijson.Field
	ProgramLevel       apijson.Field
	State              apijson.Field
	Type               apijson.Field
	ExcludedCardTokens apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ApplyResponseCurrentVersion struct {
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2ApplyResponseCurrentVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                          `json:"period"`
	Scope  AuthRuleV2ApplyResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2ApplyResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2ApplyResponseCurrentVersionParametersUnion
}

// authRuleV2ApplyResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2ApplyResponseCurrentVersionParameters]
type authRuleV2ApplyResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2ApplyResponseCurrentVersionParameters) AsUnion() AuthRuleV2ApplyResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2ApplyResponseCurrentVersionParametersUnion interface {
	implementsAuthRuleV2ApplyResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2ApplyResponseDraftVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                        `json:"period"`
	Scope  AuthRuleV2ApplyResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2ApplyResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2ApplyResponseDraftVersionParametersUnion
}

// authRuleV2ApplyResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2ApplyResponseDraftVersionParameters]
type authRuleV2ApplyResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2ApplyResponseDraftVersionParameters) AsUnion() AuthRuleV2ApplyResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2ApplyResponseDraftVersionParametersUnion interface {
	implementsAuthRuleV2ApplyResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                              `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2DraftResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2DraftResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2DraftResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2DraftResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                    `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2DraftResponseJSON `json:"-"`
}

// authRuleV2DraftResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2DraftResponse]
type authRuleV2DraftResponseJSON struct {
	Token              apijson.Field
	AccountTokens      apijson.Field
	CardTokens         apijson.Field
	CurrentVersion     apijson.Field
	DraftVersion       apijson.Field
	Name               apijson.Field
	ProgramLevel       apijson.Field
	State              apijson.Field
	Type               apijson.Field
	ExcludedCardTokens apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2DraftResponseCurrentVersion struct {
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2DraftResponseCurrentVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                          `json:"period"`
	Scope  AuthRuleV2DraftResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2DraftResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2DraftResponseCurrentVersionParametersUnion
}

// authRuleV2DraftResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2DraftResponseCurrentVersionParameters]
type authRuleV2DraftResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2DraftResponseCurrentVersionParameters) AsUnion() AuthRuleV2DraftResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2DraftResponseCurrentVersionParametersUnion interface {
	implementsAuthRuleV2DraftResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2DraftResponseDraftVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                        `json:"period"`
	Scope  AuthRuleV2DraftResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2DraftResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2DraftResponseDraftVersionParametersUnion
}

// authRuleV2DraftResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2DraftResponseDraftVersionParameters]
type authRuleV2DraftResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2DraftResponseDraftVersionParameters) AsUnion() AuthRuleV2DraftResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2DraftResponseDraftVersionParametersUnion interface {
	implementsAuthRuleV2DraftResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                                `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2PromoteResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2PromoteResponseDraftVersion   `json:"draft_version,required,nullable"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2PromoteResponseState `json:"state,required"`
	// The type of Auth Rule
	Type AuthRuleV2PromoteResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                      `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2PromoteResponseJSON `json:"-"`
}

// authRuleV2PromoteResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponse]
type authRuleV2PromoteResponseJSON struct {
	Token              apijson.Field
	AccountTokens      apijson.Field
	CardTokens         apijson.Field
	CurrentVersion     apijson.Field
	DraftVersion       apijson.Field
	Name               apijson.Field
	ProgramLevel       apijson.Field
	State              apijson.Field
	Type               apijson.Field
	ExcludedCardTokens apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2PromoteResponseCurrentVersion struct {
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2PromoteResponseCurrentVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                            `json:"period"`
	Scope  AuthRuleV2PromoteResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2PromoteResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2PromoteResponseCurrentVersionParametersUnion
}

// authRuleV2PromoteResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2PromoteResponseCurrentVersionParameters]
type authRuleV2PromoteResponseCurrentVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2PromoteResponseCurrentVersionParameters) AsUnion() AuthRuleV2PromoteResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2PromoteResponseCurrentVersionParametersUnion interface {
	implementsAuthRuleV2PromoteResponseCurrentVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseCurrentVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	// Parameters for the Auth Rule
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

// Parameters for the Auth Rule
type AuthRuleV2PromoteResponseDraftVersionParameters struct {
	// This field can have the runtime type of [[]AuthRuleCondition].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of [VelocityLimitParamsFilters].
	Filters interface{} `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount int64 `json:"limit_amount,nullable"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount int64 `json:"limit_count,nullable"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                          `json:"period"`
	Scope  AuthRuleV2PromoteResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2PromoteResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2PromoteResponseDraftVersionParametersUnion
}

// authRuleV2PromoteResponseDraftVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2PromoteResponseDraftVersionParameters]
type authRuleV2PromoteResponseDraftVersionParametersJSON struct {
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Period      apijson.Field
	Scope       apijson.Field
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
// Possible runtime types of the union are [ConditionalBlockParameters],
// [VelocityLimitParams].
func (r AuthRuleV2PromoteResponseDraftVersionParameters) AsUnion() AuthRuleV2PromoteResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters] or [VelocityLimitParams].
type AuthRuleV2PromoteResponseDraftVersionParametersUnion interface {
	implementsAuthRuleV2PromoteResponseDraftVersionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseDraftVersionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalBlockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParams{}),
		},
	)
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
	AccountTokens      param.Field[interface{}] `json:"account_tokens"`
	CardTokens         param.Field[interface{}] `json:"card_tokens"`
	ExcludedCardTokens param.Field[interface{}] `json:"excluded_card_tokens"`
	// Auth Rule Name
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level"`
	// The type of Auth Rule
	Type param.Field[AuthRuleV2NewParamsBodyType] `json:"type"`
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
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion] `json:"parameters"`
	// The type of Auth Rule
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens) implementsAuthRuleV2NewParamsBodyUnion() {
}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters struct {
	Conditions param.Field[interface{}] `json:"conditions"`
	Filters    param.Field[interface{}] `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[int64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[int64]                                                                    `json:"limit_count"`
	Period     param.Field[interface{}]                                                              `json:"period"`
	Scope      param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion()
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
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion] `json:"parameters"`
	// The type of Auth Rule
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens) implementsAuthRuleV2NewParamsBodyUnion() {
}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters struct {
	Conditions param.Field[interface{}] `json:"conditions"`
	Filters    param.Field[interface{}] `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[int64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[int64]                                                                 `json:"limit_count"`
	Period     param.Field[interface{}]                                                           `json:"period"`
	Scope      param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion()
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
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens param.Field[[]string] `json:"excluded_card_tokens" format:"uuid"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion] `json:"parameters"`
	// The type of Auth Rule
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel) implementsAuthRuleV2NewParamsBodyUnion() {
}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters struct {
	Conditions param.Field[interface{}] `json:"conditions"`
	Filters    param.Field[interface{}] `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[int64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[int64]                                                                   `json:"limit_count"`
	Period     param.Field[interface{}]                                                             `json:"period"`
	Scope      param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion()
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
	Body AuthRuleV2UpdateParamsBodyUnion `json:"body,required"`
}

func (r AuthRuleV2UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AuthRuleV2UpdateParamsBody struct {
	AccountTokens      param.Field[interface{}] `json:"account_tokens"`
	CardTokens         param.Field[interface{}] `json:"card_tokens"`
	ExcludedCardTokens param.Field[interface{}] `json:"excluded_card_tokens"`
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
	AccountTokens      param.Field[interface{}] `json:"account_tokens"`
	CardTokens         param.Field[interface{}] `json:"card_tokens"`
	ExcludedCardTokens param.Field[interface{}] `json:"excluded_card_tokens"`
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
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens param.Field[[]string] `json:"excluded_card_tokens" format:"uuid"`
}

func (r AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestProgramLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestProgramLevel) implementsAuthRuleV2ApplyParamsBodyUnion() {
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
	Conditions param.Field[interface{}] `json:"conditions"`
	Filters    param.Field[interface{}] `json:"filters"`
	// The maximum amount of spend velocity allowed in the period in minor units (the
	// smallest unit of a currency, e.g. cents for USD). Transactions exceeding this
	// limit will be declined.
	LimitAmount param.Field[int64] `json:"limit_amount"`
	// The number of spend velocity impacting transactions may not exceed this limit in
	// the period. Transactions exceeding this limit will be declined. A spend velocity
	// impacting transaction is a transaction that has been authorized, and optionally
	// settled, or a force post (a transaction that settled without prior
	// authorization).
	LimitCount param.Field[int64]                                `json:"limit_count"`
	Period     param.Field[interface{}]                          `json:"period"`
	Scope      param.Field[AuthRuleV2DraftParamsParametersScope] `json:"scope"`
}

func (r AuthRuleV2DraftParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2DraftParamsParameters) implementsAuthRuleV2DraftParamsParametersUnion() {}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2DraftParamsParameters].
type AuthRuleV2DraftParamsParametersUnion interface {
	implementsAuthRuleV2DraftParamsParametersUnion()
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
