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
func (r *AuthRuleV2Service) New(ctx context.Context, body AuthRuleV2NewParams, opts ...option.RequestOption) (res *AuthRuleV2NewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/auth_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a V2 Auth rule by its token
func (r *AuthRuleV2Service) Get(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleV2GetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a V2 Auth rule's properties
//
// If `account_tokens`, `card_tokens`, `program_level`, or `excluded_card_tokens`
// is provided, this will replace existing associations with the provided list of
// entities.
func (r *AuthRuleV2Service) Update(ctx context.Context, authRuleToken string, body AuthRuleV2UpdateParams, opts ...option.RequestOption) (res *AuthRuleV2UpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists V2 Auth rules
func (r *AuthRuleV2Service) List(ctx context.Context, query AuthRuleV2ListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AuthRuleV2ListResponse], err error) {
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
func (r *AuthRuleV2Service) ListAutoPaging(ctx context.Context, query AuthRuleV2ListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AuthRuleV2ListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a V2 Auth rule
func (r *AuthRuleV2Service) Delete(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Associates a V2 Auth rule with a card program, the provided account(s) or
// card(s).
//
// Prefer using the `PATCH` method for this operation.
//
// Deprecated: deprecated
func (r *AuthRuleV2Service) Apply(ctx context.Context, authRuleToken string, body AuthRuleV2ApplyParams, opts ...option.RequestOption) (res *AuthRuleV2ApplyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/draft", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Promotes the draft version of an Auth rule to the currently active version such
// that it is enforced in the respective stream.
func (r *AuthRuleV2Service) Promote(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleV2PromoteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/promote", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint is deprecated and will be removed in the future. Requests a
// performance report of an Auth rule to be asynchronously generated. Reports can
// only be run on rules in draft or active mode and will included approved and
// declined statistics as well as examples. The generated report will be delivered
// asynchronously through a webhook with `event_type` =
// `auth_rules.performance_report.created`. See the docs on setting up
// [webhook subscriptions](https://docs.lithic.com/docs/events-api).
//
// Reports are generated based on data collected by Lithic's processing system in
// the trailing week. The performance of the auth rule will be assessed on the
// configuration of the auth rule at the time the report is requested. This implies
// that if a performance report is requested, right after updating an auth rule,
// depending on the number of events processed for a card program, it may be the
// case that no data is available for the report. Therefore Lithic recommends to
// decouple making updates to an Auth Rule, and requesting performance reports.
//
// To make this concrete, consider the following example:
//
//  1. At time `t`, a new Auth Rule is created, and applies to all auth events on a
//     card program. The Auth Rule has not yet been promoted, causing the draft
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
//     requested right at this moment, it will only contain data for events to which
//     both the active version and the draft version is applied. Lithic does this to
//     ensure that performance reports represent a fair comparison between rules.
//     Because there may be no events in this window, and because there may be some
//     lag before data is available in a performance report, the requested
//     performance report could contain no to little data.
//  5. At time `t + 4 hours` another performance report is requested: this time the
//     performance report will contain data from the window between `t + 3 hours`
//     and `t + 4 hours`, for any events to which both the current version of the
//     Auth rule (in enforcing mode) and the draft version of the Auth rule (in
//     shadow mode) applied.
//
// Note that generating a report may take up to 15 minutes and that delivery is not
// guaranteed. Customers are required to have created an event subscription to
// receive the webhook. Additionally, there is a delay of approximately 15 minutes
// between when Lithic's transaction processing systems have processed the
// transaction, and when a transaction will be included in the report.
//
// Deprecated: deprecated
func (r *AuthRuleV2Service) Report(ctx context.Context, authRuleToken string, opts ...option.RequestOption) (res *AuthRuleV2ReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/report", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
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
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/features", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/report", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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
	AuthRuleConditionOperationIsOneOf                AuthRuleConditionOperation = "IS_ONE_OF"
	AuthRuleConditionOperationIsNotOneOf             AuthRuleConditionOperation = "IS_NOT_ONE_OF"
	AuthRuleConditionOperationMatches                AuthRuleConditionOperation = "MATCHES"
	AuthRuleConditionOperationDoesNotMatch           AuthRuleConditionOperation = "DOES_NOT_MATCH"
	AuthRuleConditionOperationIsEqualTo              AuthRuleConditionOperation = "IS_EQUAL_TO"
	AuthRuleConditionOperationIsNotEqualTo           AuthRuleConditionOperation = "IS_NOT_EQUAL_TO"
	AuthRuleConditionOperationIsGreaterThan          AuthRuleConditionOperation = "IS_GREATER_THAN"
	AuthRuleConditionOperationIsGreaterThanOrEqualTo AuthRuleConditionOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleConditionOperationIsLessThan             AuthRuleConditionOperation = "IS_LESS_THAN"
	AuthRuleConditionOperationIsLessThanOrEqualTo    AuthRuleConditionOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleConditionOperation) IsKnown() bool {
	switch r {
	case AuthRuleConditionOperationIsOneOf, AuthRuleConditionOperationIsNotOneOf, AuthRuleConditionOperationMatches, AuthRuleConditionOperationDoesNotMatch, AuthRuleConditionOperationIsEqualTo, AuthRuleConditionOperationIsNotEqualTo, AuthRuleConditionOperationIsGreaterThan, AuthRuleConditionOperationIsGreaterThanOrEqualTo, AuthRuleConditionOperationIsLessThan, AuthRuleConditionOperationIsLessThanOrEqualTo:
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

type Conditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     Conditional3DSActionParametersAction      `json:"action,required"`
	Conditions []Conditional3DsActionParametersCondition `json:"conditions,required"`
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

func (r Conditional3DSActionParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {
}

func (r Conditional3DSActionParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2ApplyResponseCurrentVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2ApplyResponseDraftVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {}

func (r Conditional3DSActionParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {
}

func (r Conditional3DSActionParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {}

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
	Attribute Conditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation Conditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value Conditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  conditional3DsActionParametersConditionJSON        `json:"-"`
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
)

func (r Conditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case Conditional3DSActionParametersConditionsAttributeMcc, Conditional3DSActionParametersConditionsAttributeCountry, Conditional3DSActionParametersConditionsAttributeCurrency, Conditional3DSActionParametersConditionsAttributeMerchantID, Conditional3DSActionParametersConditionsAttributeDescriptor, Conditional3DSActionParametersConditionsAttributeTransactionAmount, Conditional3DSActionParametersConditionsAttributeRiskScore, Conditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type Conditional3DSActionParametersConditionsOperation string

const (
	Conditional3DSActionParametersConditionsOperationIsOneOf                Conditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	Conditional3DSActionParametersConditionsOperationIsNotOneOf             Conditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	Conditional3DSActionParametersConditionsOperationMatches                Conditional3DSActionParametersConditionsOperation = "MATCHES"
	Conditional3DSActionParametersConditionsOperationDoesNotMatch           Conditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	Conditional3DSActionParametersConditionsOperationIsEqualTo              Conditional3DSActionParametersConditionsOperation = "IS_EQUAL_TO"
	Conditional3DSActionParametersConditionsOperationIsNotEqualTo           Conditional3DSActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	Conditional3DSActionParametersConditionsOperationIsGreaterThan          Conditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	Conditional3DSActionParametersConditionsOperationIsGreaterThanOrEqualTo Conditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	Conditional3DSActionParametersConditionsOperationIsLessThan             Conditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
	Conditional3DSActionParametersConditionsOperationIsLessThanOrEqualTo    Conditional3DSActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r Conditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case Conditional3DSActionParametersConditionsOperationIsOneOf, Conditional3DSActionParametersConditionsOperationIsNotOneOf, Conditional3DSActionParametersConditionsOperationMatches, Conditional3DSActionParametersConditionsOperationDoesNotMatch, Conditional3DSActionParametersConditionsOperationIsEqualTo, Conditional3DSActionParametersConditionsOperationIsNotEqualTo, Conditional3DSActionParametersConditionsOperationIsGreaterThan, Conditional3DSActionParametersConditionsOperationIsGreaterThanOrEqualTo, Conditional3DSActionParametersConditionsOperationIsLessThan, Conditional3DSActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [Conditional3DSActionParametersConditionsValueListOfStrings].
type Conditional3DSActionParametersConditionsValueUnion interface {
	ImplementsConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Conditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(Conditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type Conditional3DSActionParametersConditionsValueListOfStrings []string

func (r Conditional3DSActionParametersConditionsValueListOfStrings) ImplementsConditional3DsActionParametersConditionsValueUnion() {
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
)

func (r ConditionalAttribute) IsKnown() bool {
	switch r {
	case ConditionalAttributeMcc, ConditionalAttributeCountry, ConditionalAttributeCurrency, ConditionalAttributeMerchantID, ConditionalAttributeDescriptor, ConditionalAttributeLiabilityShift, ConditionalAttributePanEntryMode, ConditionalAttributeTransactionAmount, ConditionalAttributeRiskScore, ConditionalAttributeCardTransactionCount15M, ConditionalAttributeCardTransactionCount1H, ConditionalAttributeCardTransactionCount24H, ConditionalAttributeCardState, ConditionalAttributePinEntered, ConditionalAttributePinStatus, ConditionalAttributeWalletType:
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

type MerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []MerchantLockParametersMerchant `json:"merchants,required"`
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

func (r MerchantLockParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2ApplyResponseCurrentVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2ApplyResponseDraftVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {}

func (r MerchantLockParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {}

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

type RuleStats struct {
	// The total number of historical transactions approved by this rule during the
	// relevant period, or the number of transactions that would have been approved if
	// the rule was evaluated in shadow mode.
	Approved int64 `json:"approved"`
	// The total number of historical transactions challenged by this rule during the
	// relevant period, or the number of transactions that would have been challenged
	// if the rule was evaluated in shadow mode. Currently applicable only for 3DS Auth
	// Rules.
	Challenged int64 `json:"challenged"`
	// The total number of historical transactions declined by this rule during the
	// relevant period, or the number of transactions that would have been declined if
	// the rule was evaluated in shadow mode.
	Declined int64 `json:"declined"`
	// Example events and their outcomes.
	Examples []RuleStatsExample `json:"examples"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64         `json:"version"`
	JSON    ruleStatsJSON `json:"-"`
}

// ruleStatsJSON contains the JSON metadata for the struct [RuleStats]
type ruleStatsJSON struct {
	Approved    apijson.Field
	Challenged  apijson.Field
	Declined    apijson.Field
	Examples    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleStatsJSON) RawJSON() string {
	return r.raw
}

type RuleStatsExample struct {
	// Whether the rule would have approved the request.
	Approved bool `json:"approved"`
	// The decision made by the rule for this event.
	Decision RuleStatsExamplesDecision `json:"decision"`
	// The event token.
	EventToken string `json:"event_token" format:"uuid"`
	// The timestamp of the event.
	Timestamp time.Time            `json:"timestamp" format:"date-time"`
	JSON      ruleStatsExampleJSON `json:"-"`
}

// ruleStatsExampleJSON contains the JSON metadata for the struct
// [RuleStatsExample]
type ruleStatsExampleJSON struct {
	Approved    apijson.Field
	Decision    apijson.Field
	EventToken  apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleStatsExample) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleStatsExampleJSON) RawJSON() string {
	return r.raw
}

// The decision made by the rule for this event.
type RuleStatsExamplesDecision string

const (
	RuleStatsExamplesDecisionApproved   RuleStatsExamplesDecision = "APPROVED"
	RuleStatsExamplesDecisionDeclined   RuleStatsExamplesDecision = "DECLINED"
	RuleStatsExamplesDecisionChallenged RuleStatsExamplesDecision = "CHALLENGED"
)

func (r RuleStatsExamplesDecision) IsKnown() bool {
	switch r {
	case RuleStatsExamplesDecisionApproved, RuleStatsExamplesDecisionDeclined, RuleStatsExamplesDecisionChallenged:
		return true
	}
	return false
}

type VelocityLimitParams struct {
	Filters VelocityLimitParamsFilters `json:"filters,required"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period,required"`
	// The scope the velocity is calculated for
	Scope VelocityLimitParamsScope `json:"scope,required"`
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
	// ISO-3166-1 alpha-3 Country Codes to exclude from the velocity calculation.
	// Transactions matching any of the provided will be excluded from the calculated
	// velocity.
	ExcludeCountries []string `json:"exclude_countries,nullable"`
	// Merchant Category Codes to exclude from the velocity calculation. Transactions
	// matching this MCC will be excluded from the calculated velocity.
	ExcludeMccs []string `json:"exclude_mccs,nullable"`
	// ISO-3166-1 alpha-3 Country Codes to include in the velocity calculation.
	// Transactions not matching any of the provided will not be included in the
	// calculated velocity.
	IncludeCountries []string `json:"include_countries,nullable"`
	// Merchant Category Codes to include in the velocity calculation. Transactions not
	// matching this MCC will not be included in the calculated velocity.
	IncludeMccs []string `json:"include_mccs,nullable"`
	// PAN entry modes to include in the velocity calculation. Transactions not
	// matching any of the provided will not be included in the calculated velocity.
	IncludePanEntryModes []VelocityLimitParamsFiltersIncludePanEntryMode `json:"include_pan_entry_modes,nullable"`
	JSON                 velocityLimitParamsFiltersJSON                  `json:"-"`
}

// velocityLimitParamsFiltersJSON contains the JSON metadata for the struct
// [VelocityLimitParamsFilters]
type velocityLimitParamsFiltersJSON struct {
	ExcludeCountries     apijson.Field
	ExcludeMccs          apijson.Field
	IncludeCountries     apijson.Field
	IncludeMccs          apijson.Field
	IncludePanEntryModes apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *VelocityLimitParamsFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitParamsFiltersJSON) RawJSON() string {
	return r.raw
}

type VelocityLimitParamsFiltersIncludePanEntryMode string

const (
	VelocityLimitParamsFiltersIncludePanEntryModeAutoEntry           VelocityLimitParamsFiltersIncludePanEntryMode = "AUTO_ENTRY"
	VelocityLimitParamsFiltersIncludePanEntryModeBarCode             VelocityLimitParamsFiltersIncludePanEntryMode = "BAR_CODE"
	VelocityLimitParamsFiltersIncludePanEntryModeContactless         VelocityLimitParamsFiltersIncludePanEntryMode = "CONTACTLESS"
	VelocityLimitParamsFiltersIncludePanEntryModeCredentialOnFile    VelocityLimitParamsFiltersIncludePanEntryMode = "CREDENTIAL_ON_FILE"
	VelocityLimitParamsFiltersIncludePanEntryModeEcommerce           VelocityLimitParamsFiltersIncludePanEntryMode = "ECOMMERCE"
	VelocityLimitParamsFiltersIncludePanEntryModeErrorKeyed          VelocityLimitParamsFiltersIncludePanEntryMode = "ERROR_KEYED"
	VelocityLimitParamsFiltersIncludePanEntryModeErrorMagneticStripe VelocityLimitParamsFiltersIncludePanEntryMode = "ERROR_MAGNETIC_STRIPE"
	VelocityLimitParamsFiltersIncludePanEntryModeIcc                 VelocityLimitParamsFiltersIncludePanEntryMode = "ICC"
	VelocityLimitParamsFiltersIncludePanEntryModeKeyEntered          VelocityLimitParamsFiltersIncludePanEntryMode = "KEY_ENTERED"
	VelocityLimitParamsFiltersIncludePanEntryModeMagneticStripe      VelocityLimitParamsFiltersIncludePanEntryMode = "MAGNETIC_STRIPE"
	VelocityLimitParamsFiltersIncludePanEntryModeManual              VelocityLimitParamsFiltersIncludePanEntryMode = "MANUAL"
	VelocityLimitParamsFiltersIncludePanEntryModeOcr                 VelocityLimitParamsFiltersIncludePanEntryMode = "OCR"
	VelocityLimitParamsFiltersIncludePanEntryModeSecureCardless      VelocityLimitParamsFiltersIncludePanEntryMode = "SECURE_CARDLESS"
	VelocityLimitParamsFiltersIncludePanEntryModeUnspecified         VelocityLimitParamsFiltersIncludePanEntryMode = "UNSPECIFIED"
	VelocityLimitParamsFiltersIncludePanEntryModeUnknown             VelocityLimitParamsFiltersIncludePanEntryMode = "UNKNOWN"
)

func (r VelocityLimitParamsFiltersIncludePanEntryMode) IsKnown() bool {
	switch r {
	case VelocityLimitParamsFiltersIncludePanEntryModeAutoEntry, VelocityLimitParamsFiltersIncludePanEntryModeBarCode, VelocityLimitParamsFiltersIncludePanEntryModeContactless, VelocityLimitParamsFiltersIncludePanEntryModeCredentialOnFile, VelocityLimitParamsFiltersIncludePanEntryModeEcommerce, VelocityLimitParamsFiltersIncludePanEntryModeErrorKeyed, VelocityLimitParamsFiltersIncludePanEntryModeErrorMagneticStripe, VelocityLimitParamsFiltersIncludePanEntryModeIcc, VelocityLimitParamsFiltersIncludePanEntryModeKeyEntered, VelocityLimitParamsFiltersIncludePanEntryModeMagneticStripe, VelocityLimitParamsFiltersIncludePanEntryModeManual, VelocityLimitParamsFiltersIncludePanEntryModeOcr, VelocityLimitParamsFiltersIncludePanEntryModeSecureCardless, VelocityLimitParamsFiltersIncludePanEntryModeUnspecified, VelocityLimitParamsFiltersIncludePanEntryModeUnknown:
		return true
	}
	return false
}

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

// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
//
// The size of the trailing window to calculate Spend Velocity over in seconds. The
// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
//
// Union satisfied by [shared.UnionInt],
// [VelocityLimitParamsPeriodWindowFixedWindow],
// [VelocityLimitParamsPeriodWindowTrailingWindowObject],
// [VelocityLimitParamsPeriodWindowFixedWindowDay],
// [VelocityLimitParamsPeriodWindowFixedWindowWeek],
// [VelocityLimitParamsPeriodWindowFixedWindowMonth] or
// [VelocityLimitParamsPeriodWindowFixedWindowYear].
type VelocityLimitParamsPeriodWindowUnion interface {
	ImplementsVelocityLimitParamsPeriodWindowUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VelocityLimitParamsPeriodWindowUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(VelocityLimitParamsPeriodWindowFixedWindow("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParamsPeriodWindowTrailingWindowObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParamsPeriodWindowFixedWindowDay{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParamsPeriodWindowFixedWindowWeek{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParamsPeriodWindowFixedWindowMonth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VelocityLimitParamsPeriodWindowFixedWindowYear{}),
		},
	)
}

// DEPRECATED: This has been deprecated in favor of the other Fixed Window Objects
//
// The window of time to calculate Spend Velocity over.
//
//   - `DAY`: Velocity over the current day since midnight Eastern Time.
//   - `WEEK`: Velocity over the current week since 00:00 / 12 AM on Monday in
//     Eastern Time.
//   - `MONTH`: Velocity over the current month since 00:00 / 12 AM on the first of
//     the month in Eastern Time.
//   - `YEAR`: Velocity over the current year since 00:00 / 12 AM on January 1st in
//     Eastern Time.
type VelocityLimitParamsPeriodWindowFixedWindow string

const (
	VelocityLimitParamsPeriodWindowFixedWindowDay   VelocityLimitParamsPeriodWindowFixedWindow = "DAY"
	VelocityLimitParamsPeriodWindowFixedWindowWeek  VelocityLimitParamsPeriodWindowFixedWindow = "WEEK"
	VelocityLimitParamsPeriodWindowFixedWindowMonth VelocityLimitParamsPeriodWindowFixedWindow = "MONTH"
	VelocityLimitParamsPeriodWindowFixedWindowYear  VelocityLimitParamsPeriodWindowFixedWindow = "YEAR"
)

func (r VelocityLimitParamsPeriodWindowFixedWindow) IsKnown() bool {
	switch r {
	case VelocityLimitParamsPeriodWindowFixedWindowDay, VelocityLimitParamsPeriodWindowFixedWindowWeek, VelocityLimitParamsPeriodWindowFixedWindowMonth, VelocityLimitParamsPeriodWindowFixedWindowYear:
		return true
	}
	return false
}

func (r VelocityLimitParamsPeriodWindowFixedWindow) ImplementsVelocityLimitParamsPeriodWindowUnion() {
}

type VelocityLimitParamsPeriodWindowTrailingWindowObject struct {
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Duration int64                                                   `json:"duration"`
	Type     VelocityLimitParamsPeriodWindowTrailingWindowObjectType `json:"type"`
	JSON     velocityLimitParamsPeriodWindowTrailingWindowObjectJSON `json:"-"`
}

// velocityLimitParamsPeriodWindowTrailingWindowObjectJSON contains the JSON
// metadata for the struct [VelocityLimitParamsPeriodWindowTrailingWindowObject]
type velocityLimitParamsPeriodWindowTrailingWindowObjectJSON struct {
	Duration    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VelocityLimitParamsPeriodWindowTrailingWindowObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r velocityLimitParamsPeriodWindowTrailingWindowObjectJSON) RawJSON() string {
	return r.raw
}

func (r VelocityLimitParamsPeriodWindowTrailingWindowObject) ImplementsVelocityLimitParamsPeriodWindowUnion() {
}

type VelocityLimitParamsPeriodWindowTrailingWindowObjectType string

const (
	VelocityLimitParamsPeriodWindowTrailingWindowObjectTypeCustom VelocityLimitParamsPeriodWindowTrailingWindowObjectType = "CUSTOM"
)

func (r VelocityLimitParamsPeriodWindowTrailingWindowObjectType) IsKnown() bool {
	switch r {
	case VelocityLimitParamsPeriodWindowTrailingWindowObjectTypeCustom:
		return true
	}
	return false
}

type AuthRuleV2NewResponse struct {
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens []string `json:"business_account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                            `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2NewResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2NewResponseDraftVersion   `json:"draft_version,required,nullable"`
	// The event stream during which the rule will be evaluated.
	EventStream AuthRuleV2NewResponseEventStream `json:"event_stream,required"`
	// Indicates whether this auth rule is managed by Lithic. If true, the rule cannot
	// be modified or deleted by the user
	LithicManaged bool `json:"lithic_managed,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2NewResponseState `json:"state,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type AuthRuleV2NewResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                  `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2NewResponseJSON `json:"-"`
}

// authRuleV2NewResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2NewResponse]
type authRuleV2NewResponseJSON struct {
	Token                 apijson.Field
	AccountTokens         apijson.Field
	BusinessAccountTokens apijson.Field
	CardTokens            apijson.Field
	CurrentVersion        apijson.Field
	DraftVersion          apijson.Field
	EventStream           apijson.Field
	LithicManaged         apijson.Field
	Name                  apijson.Field
	ProgramLevel          apijson.Field
	State                 apijson.Field
	Type                  apijson.Field
	ExcludedCardTokens    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2NewResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2NewResponseCurrentVersionParametersScope `json:"scope"`
	JSON  authRuleV2NewResponseCurrentVersionParametersJSON  `json:"-"`
	union AuthRuleV2NewResponseCurrentVersionParametersUnion
}

// authRuleV2NewResponseCurrentVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2NewResponseCurrentVersionParameters]
type authRuleV2NewResponseCurrentVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2NewResponseCurrentVersionParameters) AsUnion() AuthRuleV2NewResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewResponseCurrentVersionParametersAction string

const (
	AuthRuleV2NewResponseCurrentVersionParametersActionDecline   AuthRuleV2NewResponseCurrentVersionParametersAction = "DECLINE"
	AuthRuleV2NewResponseCurrentVersionParametersActionChallenge AuthRuleV2NewResponseCurrentVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersActionDecline, AuthRuleV2NewResponseCurrentVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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
	// The action to take if the conditions are met.
	Action AuthRuleV2NewResponseDraftVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2NewResponseDraftVersionParametersScope `json:"scope"`
	JSON  authRuleV2NewResponseDraftVersionParametersJSON  `json:"-"`
	union AuthRuleV2NewResponseDraftVersionParametersUnion
}

// authRuleV2NewResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2NewResponseDraftVersionParameters]
type authRuleV2NewResponseDraftVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2NewResponseDraftVersionParameters) AsUnion() AuthRuleV2NewResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewResponseDraftVersionParametersAction string

const (
	AuthRuleV2NewResponseDraftVersionParametersActionDecline   AuthRuleV2NewResponseDraftVersionParametersAction = "DECLINE"
	AuthRuleV2NewResponseDraftVersionParametersActionChallenge AuthRuleV2NewResponseDraftVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewResponseDraftVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersActionDecline, AuthRuleV2NewResponseDraftVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewResponseEventStream string

const (
	AuthRuleV2NewResponseEventStreamAuthorization         AuthRuleV2NewResponseEventStream = "AUTHORIZATION"
	AuthRuleV2NewResponseEventStreamThreeDSAuthentication AuthRuleV2NewResponseEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2NewResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseEventStreamAuthorization, AuthRuleV2NewResponseEventStreamThreeDSAuthentication:
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewResponseType string

const (
	AuthRuleV2NewResponseTypeConditionalBlock  AuthRuleV2NewResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewResponseTypeVelocityLimit     AuthRuleV2NewResponseType = "VELOCITY_LIMIT"
	AuthRuleV2NewResponseTypeMerchantLock      AuthRuleV2NewResponseType = "MERCHANT_LOCK"
	AuthRuleV2NewResponseTypeConditionalAction AuthRuleV2NewResponseType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2NewResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseTypeConditionalBlock, AuthRuleV2NewResponseTypeVelocityLimit, AuthRuleV2NewResponseTypeMerchantLock, AuthRuleV2NewResponseTypeConditionalAction:
		return true
	}
	return false
}

type AuthRuleV2GetResponse struct {
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens []string `json:"business_account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                            `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2GetResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2GetResponseDraftVersion   `json:"draft_version,required,nullable"`
	// The event stream during which the rule will be evaluated.
	EventStream AuthRuleV2GetResponseEventStream `json:"event_stream,required"`
	// Indicates whether this auth rule is managed by Lithic. If true, the rule cannot
	// be modified or deleted by the user
	LithicManaged bool `json:"lithic_managed,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2GetResponseState `json:"state,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type AuthRuleV2GetResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                  `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2GetResponseJSON `json:"-"`
}

// authRuleV2GetResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2GetResponse]
type authRuleV2GetResponseJSON struct {
	Token                 apijson.Field
	AccountTokens         apijson.Field
	BusinessAccountTokens apijson.Field
	CardTokens            apijson.Field
	CurrentVersion        apijson.Field
	DraftVersion          apijson.Field
	EventStream           apijson.Field
	LithicManaged         apijson.Field
	Name                  apijson.Field
	ProgramLevel          apijson.Field
	State                 apijson.Field
	Type                  apijson.Field
	ExcludedCardTokens    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2GetResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2GetResponseCurrentVersionParametersScope `json:"scope"`
	JSON  authRuleV2GetResponseCurrentVersionParametersJSON  `json:"-"`
	union AuthRuleV2GetResponseCurrentVersionParametersUnion
}

// authRuleV2GetResponseCurrentVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2GetResponseCurrentVersionParameters]
type authRuleV2GetResponseCurrentVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2GetResponseCurrentVersionParameters) AsUnion() AuthRuleV2GetResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2GetResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2GetResponseCurrentVersionParametersAction string

const (
	AuthRuleV2GetResponseCurrentVersionParametersActionDecline   AuthRuleV2GetResponseCurrentVersionParametersAction = "DECLINE"
	AuthRuleV2GetResponseCurrentVersionParametersActionChallenge AuthRuleV2GetResponseCurrentVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersActionDecline, AuthRuleV2GetResponseCurrentVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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
	// The action to take if the conditions are met.
	Action AuthRuleV2GetResponseDraftVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2GetResponseDraftVersionParametersScope `json:"scope"`
	JSON  authRuleV2GetResponseDraftVersionParametersJSON  `json:"-"`
	union AuthRuleV2GetResponseDraftVersionParametersUnion
}

// authRuleV2GetResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2GetResponseDraftVersionParameters]
type authRuleV2GetResponseDraftVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2GetResponseDraftVersionParameters) AsUnion() AuthRuleV2GetResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2GetResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2GetResponseDraftVersionParametersAction string

const (
	AuthRuleV2GetResponseDraftVersionParametersActionDecline   AuthRuleV2GetResponseDraftVersionParametersAction = "DECLINE"
	AuthRuleV2GetResponseDraftVersionParametersActionChallenge AuthRuleV2GetResponseDraftVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2GetResponseDraftVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersActionDecline, AuthRuleV2GetResponseDraftVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The event stream during which the rule will be evaluated.
type AuthRuleV2GetResponseEventStream string

const (
	AuthRuleV2GetResponseEventStreamAuthorization         AuthRuleV2GetResponseEventStream = "AUTHORIZATION"
	AuthRuleV2GetResponseEventStreamThreeDSAuthentication AuthRuleV2GetResponseEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2GetResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseEventStreamAuthorization, AuthRuleV2GetResponseEventStreamThreeDSAuthentication:
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2GetResponseType string

const (
	AuthRuleV2GetResponseTypeConditionalBlock  AuthRuleV2GetResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2GetResponseTypeVelocityLimit     AuthRuleV2GetResponseType = "VELOCITY_LIMIT"
	AuthRuleV2GetResponseTypeMerchantLock      AuthRuleV2GetResponseType = "MERCHANT_LOCK"
	AuthRuleV2GetResponseTypeConditionalAction AuthRuleV2GetResponseType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2GetResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseTypeConditionalBlock, AuthRuleV2GetResponseTypeVelocityLimit, AuthRuleV2GetResponseTypeMerchantLock, AuthRuleV2GetResponseTypeConditionalAction:
		return true
	}
	return false
}

type AuthRuleV2UpdateResponse struct {
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens []string `json:"business_account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                               `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2UpdateResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2UpdateResponseDraftVersion   `json:"draft_version,required,nullable"`
	// The event stream during which the rule will be evaluated.
	EventStream AuthRuleV2UpdateResponseEventStream `json:"event_stream,required"`
	// Indicates whether this auth rule is managed by Lithic. If true, the rule cannot
	// be modified or deleted by the user
	LithicManaged bool `json:"lithic_managed,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2UpdateResponseState `json:"state,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type AuthRuleV2UpdateResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                     `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2UpdateResponseJSON `json:"-"`
}

// authRuleV2UpdateResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponse]
type authRuleV2UpdateResponseJSON struct {
	Token                 apijson.Field
	AccountTokens         apijson.Field
	BusinessAccountTokens apijson.Field
	CardTokens            apijson.Field
	CurrentVersion        apijson.Field
	DraftVersion          apijson.Field
	EventStream           apijson.Field
	LithicManaged         apijson.Field
	Name                  apijson.Field
	ProgramLevel          apijson.Field
	State                 apijson.Field
	Type                  apijson.Field
	ExcludedCardTokens    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2UpdateResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2UpdateResponseCurrentVersionParametersScope `json:"scope"`
	JSON  authRuleV2UpdateResponseCurrentVersionParametersJSON  `json:"-"`
	union AuthRuleV2UpdateResponseCurrentVersionParametersUnion
}

// authRuleV2UpdateResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2UpdateResponseCurrentVersionParameters]
type authRuleV2UpdateResponseCurrentVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2UpdateResponseCurrentVersionParameters) AsUnion() AuthRuleV2UpdateResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2UpdateResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2UpdateResponseCurrentVersionParametersAction string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersActionDecline   AuthRuleV2UpdateResponseCurrentVersionParametersAction = "DECLINE"
	AuthRuleV2UpdateResponseCurrentVersionParametersActionChallenge AuthRuleV2UpdateResponseCurrentVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersActionDecline, AuthRuleV2UpdateResponseCurrentVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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
	// The action to take if the conditions are met.
	Action AuthRuleV2UpdateResponseDraftVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2UpdateResponseDraftVersionParametersScope `json:"scope"`
	JSON  authRuleV2UpdateResponseDraftVersionParametersJSON  `json:"-"`
	union AuthRuleV2UpdateResponseDraftVersionParametersUnion
}

// authRuleV2UpdateResponseDraftVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2UpdateResponseDraftVersionParameters]
type authRuleV2UpdateResponseDraftVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2UpdateResponseDraftVersionParameters) AsUnion() AuthRuleV2UpdateResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2UpdateResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2UpdateResponseDraftVersionParametersAction string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersActionDecline   AuthRuleV2UpdateResponseDraftVersionParametersAction = "DECLINE"
	AuthRuleV2UpdateResponseDraftVersionParametersActionChallenge AuthRuleV2UpdateResponseDraftVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersActionDecline, AuthRuleV2UpdateResponseDraftVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The event stream during which the rule will be evaluated.
type AuthRuleV2UpdateResponseEventStream string

const (
	AuthRuleV2UpdateResponseEventStreamAuthorization         AuthRuleV2UpdateResponseEventStream = "AUTHORIZATION"
	AuthRuleV2UpdateResponseEventStreamThreeDSAuthentication AuthRuleV2UpdateResponseEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2UpdateResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseEventStreamAuthorization, AuthRuleV2UpdateResponseEventStreamThreeDSAuthentication:
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2UpdateResponseType string

const (
	AuthRuleV2UpdateResponseTypeConditionalBlock  AuthRuleV2UpdateResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2UpdateResponseTypeVelocityLimit     AuthRuleV2UpdateResponseType = "VELOCITY_LIMIT"
	AuthRuleV2UpdateResponseTypeMerchantLock      AuthRuleV2UpdateResponseType = "MERCHANT_LOCK"
	AuthRuleV2UpdateResponseTypeConditionalAction AuthRuleV2UpdateResponseType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2UpdateResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseTypeConditionalBlock, AuthRuleV2UpdateResponseTypeVelocityLimit, AuthRuleV2UpdateResponseTypeMerchantLock, AuthRuleV2UpdateResponseTypeConditionalAction:
		return true
	}
	return false
}

type AuthRuleV2ListResponse struct {
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens []string `json:"business_account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                             `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2ListResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2ListResponseDraftVersion   `json:"draft_version,required,nullable"`
	// The event stream during which the rule will be evaluated.
	EventStream AuthRuleV2ListResponseEventStream `json:"event_stream,required"`
	// Indicates whether this auth rule is managed by Lithic. If true, the rule cannot
	// be modified or deleted by the user
	LithicManaged bool `json:"lithic_managed,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2ListResponseState `json:"state,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type AuthRuleV2ListResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                   `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2ListResponseJSON `json:"-"`
}

// authRuleV2ListResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ListResponse]
type authRuleV2ListResponseJSON struct {
	Token                 apijson.Field
	AccountTokens         apijson.Field
	BusinessAccountTokens apijson.Field
	CardTokens            apijson.Field
	CurrentVersion        apijson.Field
	DraftVersion          apijson.Field
	EventStream           apijson.Field
	LithicManaged         apijson.Field
	Name                  apijson.Field
	ProgramLevel          apijson.Field
	State                 apijson.Field
	Type                  apijson.Field
	ExcludedCardTokens    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2ListResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2ListResponseCurrentVersionParametersScope `json:"scope"`
	JSON  authRuleV2ListResponseCurrentVersionParametersJSON  `json:"-"`
	union AuthRuleV2ListResponseCurrentVersionParametersUnion
}

// authRuleV2ListResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2ListResponseCurrentVersionParameters]
type authRuleV2ListResponseCurrentVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2ListResponseCurrentVersionParameters) AsUnion() AuthRuleV2ListResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2ListResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2ListResponseCurrentVersionParametersAction string

const (
	AuthRuleV2ListResponseCurrentVersionParametersActionDecline   AuthRuleV2ListResponseCurrentVersionParametersAction = "DECLINE"
	AuthRuleV2ListResponseCurrentVersionParametersActionChallenge AuthRuleV2ListResponseCurrentVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersActionDecline, AuthRuleV2ListResponseCurrentVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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
	// The action to take if the conditions are met.
	Action AuthRuleV2ListResponseDraftVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2ListResponseDraftVersionParametersScope `json:"scope"`
	JSON  authRuleV2ListResponseDraftVersionParametersJSON  `json:"-"`
	union AuthRuleV2ListResponseDraftVersionParametersUnion
}

// authRuleV2ListResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2ListResponseDraftVersionParameters]
type authRuleV2ListResponseDraftVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2ListResponseDraftVersionParameters) AsUnion() AuthRuleV2ListResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2ListResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2ListResponseDraftVersionParametersAction string

const (
	AuthRuleV2ListResponseDraftVersionParametersActionDecline   AuthRuleV2ListResponseDraftVersionParametersAction = "DECLINE"
	AuthRuleV2ListResponseDraftVersionParametersActionChallenge AuthRuleV2ListResponseDraftVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ListResponseDraftVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersActionDecline, AuthRuleV2ListResponseDraftVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The event stream during which the rule will be evaluated.
type AuthRuleV2ListResponseEventStream string

const (
	AuthRuleV2ListResponseEventStreamAuthorization         AuthRuleV2ListResponseEventStream = "AUTHORIZATION"
	AuthRuleV2ListResponseEventStreamThreeDSAuthentication AuthRuleV2ListResponseEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2ListResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseEventStreamAuthorization, AuthRuleV2ListResponseEventStreamThreeDSAuthentication:
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2ListResponseType string

const (
	AuthRuleV2ListResponseTypeConditionalBlock  AuthRuleV2ListResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2ListResponseTypeVelocityLimit     AuthRuleV2ListResponseType = "VELOCITY_LIMIT"
	AuthRuleV2ListResponseTypeMerchantLock      AuthRuleV2ListResponseType = "MERCHANT_LOCK"
	AuthRuleV2ListResponseTypeConditionalAction AuthRuleV2ListResponseType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2ListResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseTypeConditionalBlock, AuthRuleV2ListResponseTypeVelocityLimit, AuthRuleV2ListResponseTypeMerchantLock, AuthRuleV2ListResponseTypeConditionalAction:
		return true
	}
	return false
}

type AuthRuleV2ApplyResponse struct {
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens []string `json:"business_account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                              `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2ApplyResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2ApplyResponseDraftVersion   `json:"draft_version,required,nullable"`
	// The event stream during which the rule will be evaluated.
	EventStream AuthRuleV2ApplyResponseEventStream `json:"event_stream,required"`
	// Indicates whether this auth rule is managed by Lithic. If true, the rule cannot
	// be modified or deleted by the user
	LithicManaged bool `json:"lithic_managed,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2ApplyResponseState `json:"state,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type AuthRuleV2ApplyResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                    `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2ApplyResponseJSON `json:"-"`
}

// authRuleV2ApplyResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponse]
type authRuleV2ApplyResponseJSON struct {
	Token                 apijson.Field
	AccountTokens         apijson.Field
	BusinessAccountTokens apijson.Field
	CardTokens            apijson.Field
	CurrentVersion        apijson.Field
	DraftVersion          apijson.Field
	EventStream           apijson.Field
	LithicManaged         apijson.Field
	Name                  apijson.Field
	ProgramLevel          apijson.Field
	State                 apijson.Field
	Type                  apijson.Field
	ExcludedCardTokens    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2ApplyResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2ApplyResponseCurrentVersionParametersScope `json:"scope"`
	JSON  authRuleV2ApplyResponseCurrentVersionParametersJSON  `json:"-"`
	union AuthRuleV2ApplyResponseCurrentVersionParametersUnion
}

// authRuleV2ApplyResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2ApplyResponseCurrentVersionParameters]
type authRuleV2ApplyResponseCurrentVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2ApplyResponseCurrentVersionParameters) AsUnion() AuthRuleV2ApplyResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2ApplyResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2ApplyResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2ApplyResponseCurrentVersionParametersAction string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersActionDecline   AuthRuleV2ApplyResponseCurrentVersionParametersAction = "DECLINE"
	AuthRuleV2ApplyResponseCurrentVersionParametersActionChallenge AuthRuleV2ApplyResponseCurrentVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersActionDecline, AuthRuleV2ApplyResponseCurrentVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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
	// The action to take if the conditions are met.
	Action AuthRuleV2ApplyResponseDraftVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2ApplyResponseDraftVersionParametersScope `json:"scope"`
	JSON  authRuleV2ApplyResponseDraftVersionParametersJSON  `json:"-"`
	union AuthRuleV2ApplyResponseDraftVersionParametersUnion
}

// authRuleV2ApplyResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2ApplyResponseDraftVersionParameters]
type authRuleV2ApplyResponseDraftVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2ApplyResponseDraftVersionParameters) AsUnion() AuthRuleV2ApplyResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2ApplyResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2ApplyResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2ApplyResponseDraftVersionParametersAction string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersActionDecline   AuthRuleV2ApplyResponseDraftVersionParametersAction = "DECLINE"
	AuthRuleV2ApplyResponseDraftVersionParametersActionChallenge AuthRuleV2ApplyResponseDraftVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersActionDecline, AuthRuleV2ApplyResponseDraftVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The event stream during which the rule will be evaluated.
type AuthRuleV2ApplyResponseEventStream string

const (
	AuthRuleV2ApplyResponseEventStreamAuthorization         AuthRuleV2ApplyResponseEventStream = "AUTHORIZATION"
	AuthRuleV2ApplyResponseEventStreamThreeDSAuthentication AuthRuleV2ApplyResponseEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2ApplyResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseEventStreamAuthorization, AuthRuleV2ApplyResponseEventStreamThreeDSAuthentication:
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2ApplyResponseType string

const (
	AuthRuleV2ApplyResponseTypeConditionalBlock  AuthRuleV2ApplyResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2ApplyResponseTypeVelocityLimit     AuthRuleV2ApplyResponseType = "VELOCITY_LIMIT"
	AuthRuleV2ApplyResponseTypeMerchantLock      AuthRuleV2ApplyResponseType = "MERCHANT_LOCK"
	AuthRuleV2ApplyResponseTypeConditionalAction AuthRuleV2ApplyResponseType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2ApplyResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseTypeConditionalBlock, AuthRuleV2ApplyResponseTypeVelocityLimit, AuthRuleV2ApplyResponseTypeMerchantLock, AuthRuleV2ApplyResponseTypeConditionalAction:
		return true
	}
	return false
}

type AuthRuleV2DraftResponse struct {
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens []string `json:"business_account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                              `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2DraftResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2DraftResponseDraftVersion   `json:"draft_version,required,nullable"`
	// The event stream during which the rule will be evaluated.
	EventStream AuthRuleV2DraftResponseEventStream `json:"event_stream,required"`
	// Indicates whether this auth rule is managed by Lithic. If true, the rule cannot
	// be modified or deleted by the user
	LithicManaged bool `json:"lithic_managed,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2DraftResponseState `json:"state,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type AuthRuleV2DraftResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                    `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2DraftResponseJSON `json:"-"`
}

// authRuleV2DraftResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2DraftResponse]
type authRuleV2DraftResponseJSON struct {
	Token                 apijson.Field
	AccountTokens         apijson.Field
	BusinessAccountTokens apijson.Field
	CardTokens            apijson.Field
	CurrentVersion        apijson.Field
	DraftVersion          apijson.Field
	EventStream           apijson.Field
	LithicManaged         apijson.Field
	Name                  apijson.Field
	ProgramLevel          apijson.Field
	State                 apijson.Field
	Type                  apijson.Field
	ExcludedCardTokens    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2DraftResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2DraftResponseCurrentVersionParametersScope `json:"scope"`
	JSON  authRuleV2DraftResponseCurrentVersionParametersJSON  `json:"-"`
	union AuthRuleV2DraftResponseCurrentVersionParametersUnion
}

// authRuleV2DraftResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2DraftResponseCurrentVersionParameters]
type authRuleV2DraftResponseCurrentVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2DraftResponseCurrentVersionParameters) AsUnion() AuthRuleV2DraftResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2DraftResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftResponseCurrentVersionParametersAction string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersActionDecline   AuthRuleV2DraftResponseCurrentVersionParametersAction = "DECLINE"
	AuthRuleV2DraftResponseCurrentVersionParametersActionChallenge AuthRuleV2DraftResponseCurrentVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersActionDecline, AuthRuleV2DraftResponseCurrentVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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
	// The action to take if the conditions are met.
	Action AuthRuleV2DraftResponseDraftVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2DraftResponseDraftVersionParametersScope `json:"scope"`
	JSON  authRuleV2DraftResponseDraftVersionParametersJSON  `json:"-"`
	union AuthRuleV2DraftResponseDraftVersionParametersUnion
}

// authRuleV2DraftResponseDraftVersionParametersJSON contains the JSON metadata for
// the struct [AuthRuleV2DraftResponseDraftVersionParameters]
type authRuleV2DraftResponseDraftVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2DraftResponseDraftVersionParameters) AsUnion() AuthRuleV2DraftResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2DraftResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftResponseDraftVersionParametersAction string

const (
	AuthRuleV2DraftResponseDraftVersionParametersActionDecline   AuthRuleV2DraftResponseDraftVersionParametersAction = "DECLINE"
	AuthRuleV2DraftResponseDraftVersionParametersActionChallenge AuthRuleV2DraftResponseDraftVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersActionDecline, AuthRuleV2DraftResponseDraftVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The event stream during which the rule will be evaluated.
type AuthRuleV2DraftResponseEventStream string

const (
	AuthRuleV2DraftResponseEventStreamAuthorization         AuthRuleV2DraftResponseEventStream = "AUTHORIZATION"
	AuthRuleV2DraftResponseEventStreamThreeDSAuthentication AuthRuleV2DraftResponseEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2DraftResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseEventStreamAuthorization, AuthRuleV2DraftResponseEventStreamThreeDSAuthentication:
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2DraftResponseType string

const (
	AuthRuleV2DraftResponseTypeConditionalBlock  AuthRuleV2DraftResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2DraftResponseTypeVelocityLimit     AuthRuleV2DraftResponseType = "VELOCITY_LIMIT"
	AuthRuleV2DraftResponseTypeMerchantLock      AuthRuleV2DraftResponseType = "MERCHANT_LOCK"
	AuthRuleV2DraftResponseTypeConditionalAction AuthRuleV2DraftResponseType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2DraftResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseTypeConditionalBlock, AuthRuleV2DraftResponseTypeVelocityLimit, AuthRuleV2DraftResponseTypeMerchantLock, AuthRuleV2DraftResponseTypeConditionalAction:
		return true
	}
	return false
}

type AuthRuleV2PromoteResponse struct {
	// Auth Rule Token
	Token string `json:"token,required" format:"uuid"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens []string `json:"account_tokens,required" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens []string `json:"business_account_tokens,required" format:"uuid"`
	// Card tokens to which the Auth Rule applies.
	CardTokens     []string                                `json:"card_tokens,required" format:"uuid"`
	CurrentVersion AuthRuleV2PromoteResponseCurrentVersion `json:"current_version,required,nullable"`
	DraftVersion   AuthRuleV2PromoteResponseDraftVersion   `json:"draft_version,required,nullable"`
	// The event stream during which the rule will be evaluated.
	EventStream AuthRuleV2PromoteResponseEventStream `json:"event_stream,required"`
	// Indicates whether this auth rule is managed by Lithic. If true, the rule cannot
	// be modified or deleted by the user
	LithicManaged bool `json:"lithic_managed,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2PromoteResponseState `json:"state,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type AuthRuleV2PromoteResponseType `json:"type,required"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens []string                      `json:"excluded_card_tokens" format:"uuid"`
	JSON               authRuleV2PromoteResponseJSON `json:"-"`
}

// authRuleV2PromoteResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponse]
type authRuleV2PromoteResponseJSON struct {
	Token                 apijson.Field
	AccountTokens         apijson.Field
	BusinessAccountTokens apijson.Field
	CardTokens            apijson.Field
	CurrentVersion        apijson.Field
	DraftVersion          apijson.Field
	EventStream           apijson.Field
	LithicManaged         apijson.Field
	Name                  apijson.Field
	ProgramLevel          apijson.Field
	State                 apijson.Field
	Type                  apijson.Field
	ExcludedCardTokens    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2PromoteResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2PromoteResponseCurrentVersionParametersScope `json:"scope"`
	JSON  authRuleV2PromoteResponseCurrentVersionParametersJSON  `json:"-"`
	union AuthRuleV2PromoteResponseCurrentVersionParametersUnion
}

// authRuleV2PromoteResponseCurrentVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2PromoteResponseCurrentVersionParameters]
type authRuleV2PromoteResponseCurrentVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2PromoteResponseCurrentVersionParameters) AsUnion() AuthRuleV2PromoteResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2PromoteResponseCurrentVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2PromoteResponseCurrentVersionParametersAction string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersActionDecline   AuthRuleV2PromoteResponseCurrentVersionParametersAction = "DECLINE"
	AuthRuleV2PromoteResponseCurrentVersionParametersActionChallenge AuthRuleV2PromoteResponseCurrentVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersActionDecline, AuthRuleV2PromoteResponseCurrentVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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
	// The action to take if the conditions are met.
	Action AuthRuleV2PromoteResponseDraftVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition].
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
	// This field can have the runtime type of [[]MerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2PromoteResponseDraftVersionParametersScope `json:"scope"`
	JSON  authRuleV2PromoteResponseDraftVersionParametersJSON  `json:"-"`
	union AuthRuleV2PromoteResponseDraftVersionParametersUnion
}

// authRuleV2PromoteResponseDraftVersionParametersJSON contains the JSON metadata
// for the struct [AuthRuleV2PromoteResponseDraftVersionParameters]
type authRuleV2PromoteResponseDraftVersionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	Filters     apijson.Field
	LimitAmount apijson.Field
	LimitCount  apijson.Field
	Merchants   apijson.Field
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
// [VelocityLimitParams], [MerchantLockParameters],
// [Conditional3DSActionParameters],
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParameters].
func (r AuthRuleV2PromoteResponseDraftVersionParameters) AsUnion() AuthRuleV2PromoteResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters] or
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParameters{}),
		},
	)
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON        `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParameters]
type authRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON        `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition]
type authRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionJSON) RawJSON() string {
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
type AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2PromoteResponseDraftVersionParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2PromoteResponseDraftVersionParametersAction string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersActionDecline   AuthRuleV2PromoteResponseDraftVersionParametersAction = "DECLINE"
	AuthRuleV2PromoteResponseDraftVersionParametersActionChallenge AuthRuleV2PromoteResponseDraftVersionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersActionDecline, AuthRuleV2PromoteResponseDraftVersionParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The event stream during which the rule will be evaluated.
type AuthRuleV2PromoteResponseEventStream string

const (
	AuthRuleV2PromoteResponseEventStreamAuthorization         AuthRuleV2PromoteResponseEventStream = "AUTHORIZATION"
	AuthRuleV2PromoteResponseEventStreamThreeDSAuthentication AuthRuleV2PromoteResponseEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2PromoteResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseEventStreamAuthorization, AuthRuleV2PromoteResponseEventStreamThreeDSAuthentication:
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2PromoteResponseType string

const (
	AuthRuleV2PromoteResponseTypeConditionalBlock  AuthRuleV2PromoteResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2PromoteResponseTypeVelocityLimit     AuthRuleV2PromoteResponseType = "VELOCITY_LIMIT"
	AuthRuleV2PromoteResponseTypeMerchantLock      AuthRuleV2PromoteResponseType = "MERCHANT_LOCK"
	AuthRuleV2PromoteResponseTypeConditionalAction AuthRuleV2PromoteResponseType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2PromoteResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseTypeConditionalBlock, AuthRuleV2PromoteResponseTypeVelocityLimit, AuthRuleV2PromoteResponseTypeMerchantLock, AuthRuleV2PromoteResponseTypeConditionalAction:
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

type AuthRuleV2GetFeaturesResponse struct {
	// Timestamp at which the Features were evaluated
	Evaluated time.Time `json:"evaluated,required" format:"date-time"`
	// Calculated Features used for evaluation of the provided Auth Rule
	Features []AuthRuleV2GetFeaturesResponseFeature `json:"features,required"`
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
	Filters AuthRuleV2GetFeaturesResponseFeaturesFilters `json:"filters,required"`
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period VelocityLimitParamsPeriodWindowUnion `json:"period,required"`
	// The scope the velocity is calculated for
	Scope AuthRuleV2GetFeaturesResponseFeaturesScope `json:"scope,required"`
	Value AuthRuleV2GetFeaturesResponseFeaturesValue `json:"value,required"`
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

type AuthRuleV2GetFeaturesResponseFeaturesFilters struct {
	// ISO-3166-1 alpha-3 Country Codes to exclude from the velocity calculation.
	// Transactions matching any of the provided will be excluded from the calculated
	// velocity.
	ExcludeCountries []string `json:"exclude_countries,nullable"`
	// Merchant Category Codes to exclude from the velocity calculation. Transactions
	// matching this MCC will be excluded from the calculated velocity.
	ExcludeMccs []string `json:"exclude_mccs,nullable"`
	// ISO-3166-1 alpha-3 Country Codes to include in the velocity calculation.
	// Transactions not matching any of the provided will not be included in the
	// calculated velocity.
	IncludeCountries []string `json:"include_countries,nullable"`
	// Merchant Category Codes to include in the velocity calculation. Transactions not
	// matching this MCC will not be included in the calculated velocity.
	IncludeMccs []string `json:"include_mccs,nullable"`
	// PAN entry modes to include in the velocity calculation. Transactions not
	// matching any of the provided will not be included in the calculated velocity.
	IncludePanEntryModes []AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode `json:"include_pan_entry_modes,nullable"`
	JSON                 authRuleV2GetFeaturesResponseFeaturesFiltersJSON                  `json:"-"`
}

// authRuleV2GetFeaturesResponseFeaturesFiltersJSON contains the JSON metadata for
// the struct [AuthRuleV2GetFeaturesResponseFeaturesFilters]
type authRuleV2GetFeaturesResponseFeaturesFiltersJSON struct {
	ExcludeCountries     apijson.Field
	ExcludeMccs          apijson.Field
	IncludeCountries     apijson.Field
	IncludeMccs          apijson.Field
	IncludePanEntryModes apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AuthRuleV2GetFeaturesResponseFeaturesFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetFeaturesResponseFeaturesFiltersJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode string

const (
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeAutoEntry           AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "AUTO_ENTRY"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeBarCode             AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "BAR_CODE"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeContactless         AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "CONTACTLESS"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeCredentialOnFile    AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "CREDENTIAL_ON_FILE"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeEcommerce           AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "ECOMMERCE"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeErrorKeyed          AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "ERROR_KEYED"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeErrorMagneticStripe AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "ERROR_MAGNETIC_STRIPE"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeIcc                 AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "ICC"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeKeyEntered          AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "KEY_ENTERED"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeMagneticStripe      AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "MAGNETIC_STRIPE"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeManual              AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "MANUAL"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeOcr                 AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "OCR"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeSecureCardless      AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "SECURE_CARDLESS"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeUnspecified         AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "UNSPECIFIED"
	AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeUnknown             AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode = "UNKNOWN"
)

func (r AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryMode) IsKnown() bool {
	switch r {
	case AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeAutoEntry, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeBarCode, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeContactless, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeCredentialOnFile, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeEcommerce, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeErrorKeyed, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeErrorMagneticStripe, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeIcc, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeKeyEntered, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeMagneticStripe, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeManual, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeOcr, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeSecureCardless, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeUnspecified, AuthRuleV2GetFeaturesResponseFeaturesFiltersIncludePanEntryModeUnknown:
		return true
	}
	return false
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
	Amount int64 `json:"amount,required"`
	// Number of velocity impacting transactions matching the given scope, period and
	// filters
	Count int64                                          `json:"count,required"`
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
	AuthRuleToken string `json:"auth_rule_token,required" format:"uuid"`
	// The start date (UTC) of the report.
	Begin time.Time `json:"begin,required" format:"date"`
	// Daily evaluation statistics for the Auth Rule.
	DailyStatistics []AuthRuleV2GetReportResponseDailyStatistic `json:"daily_statistics,required"`
	// The end date (UTC) of the report.
	End  time.Time                       `json:"end,required" format:"date"`
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
	CurrentVersionStatistics RuleStats `json:"current_version_statistics,required,nullable"`
	// The date (UTC) for which the statistics are reported.
	Date time.Time `json:"date,required" format:"date"`
	// Detailed statistics for the draft version of the rule.
	DraftVersionStatistics RuleStats                                     `json:"draft_version_statistics,required,nullable"`
	JSON                   authRuleV2GetReportResponseDailyStatisticJSON `json:"-"`
}

// authRuleV2GetReportResponseDailyStatisticJSON contains the JSON metadata for the
// struct [AuthRuleV2GetReportResponseDailyStatistic]
type authRuleV2GetReportResponseDailyStatisticJSON struct {
	CurrentVersionStatistics apijson.Field
	Date                     apijson.Field
	DraftVersionStatistics   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AuthRuleV2GetReportResponseDailyStatistic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetReportResponseDailyStatisticJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2NewParams struct {
	Body AuthRuleV2NewParamsBodyUnion `json:"body,required"`
}

func (r AuthRuleV2NewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AuthRuleV2NewParamsBody struct {
	AccountTokens         param.Field[interface{}] `json:"account_tokens"`
	BusinessAccountTokens param.Field[interface{}] `json:"business_account_tokens"`
	CardTokens            param.Field[interface{}] `json:"card_tokens"`
	// The event stream during which the rule will be evaluated.
	EventStream        param.Field[AuthRuleV2NewParamsBodyEventStream] `json:"event_stream"`
	ExcludedCardTokens param.Field[interface{}]                        `json:"excluded_card_tokens"`
	// Auth Rule Name
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
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
	AccountTokens param.Field[[]string] `json:"account_tokens" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens param.Field[[]string] `json:"business_account_tokens" format:"uuid"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStream] `json:"event_stream"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion] `json:"parameters"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens) implementsAuthRuleV2NewParamsBodyUnion() {
}

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStream string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStreamAuthorization         AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStream = "AUTHORIZATION"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStreamThreeDSAuthentication AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStreamAuthorization, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensEventStreamThreeDSAuthentication:
		return true
	}
	return false
}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersAction] `json:"action"`
	Conditions param.Field[interface{}]                                                               `json:"conditions"`
	Filters    param.Field[interface{}]                                                               `json:"filters"`
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
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period param.Field[VelocityLimitParamsPeriodWindowUnion] `json:"period"`
	// The scope the velocity is calculated for
	Scope param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersAction]      `json:"action,required"`
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersCondition) MarshalJSON() (data []byte, err error) {
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
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalBlock  AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeVelocityLimit     AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeMerchantLock      AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalAction AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeVelocityLimit, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeMerchantLock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalAction:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens struct {
	// Card tokens to which the Auth Rule applies.
	CardTokens param.Field[[]string] `json:"card_tokens,required" format:"uuid"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStream] `json:"event_stream"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion] `json:"parameters"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens) implementsAuthRuleV2NewParamsBodyUnion() {
}

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStream string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStreamAuthorization         AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStream = "AUTHORIZATION"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStreamThreeDSAuthentication AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStreamAuthorization, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensEventStreamThreeDSAuthentication:
		return true
	}
	return false
}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersAction] `json:"action"`
	Conditions param.Field[interface{}]                                                            `json:"conditions"`
	Filters    param.Field[interface{}]                                                            `json:"filters"`
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
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period param.Field[VelocityLimitParamsPeriodWindowUnion] `json:"period"`
	// The scope the velocity is calculated for
	Scope param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersAction]      `json:"action,required"`
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersCondition) MarshalJSON() (data []byte, err error) {
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
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditionalBlock  AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeVelocityLimit     AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeMerchantLock      AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditionalAction AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeVelocityLimit, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeMerchantLock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditionalAction:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel struct {
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level,required"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStream] `json:"event_stream"`
	// Card tokens to which the Auth Rule does not apply.
	ExcludedCardTokens param.Field[[]string] `json:"excluded_card_tokens" format:"uuid"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion] `json:"parameters"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel) implementsAuthRuleV2NewParamsBodyUnion() {
}

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStream string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStreamAuthorization         AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStream = "AUTHORIZATION"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStreamThreeDSAuthentication AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStreamAuthorization, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelEventStreamThreeDSAuthentication:
		return true
	}
	return false
}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersAction] `json:"action"`
	Conditions param.Field[interface{}]                                                              `json:"conditions"`
	Filters    param.Field[interface{}]                                                              `json:"filters"`
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
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period param.Field[VelocityLimitParamsPeriodWindowUnion] `json:"period"`
	// The scope the velocity is calculated for
	Scope param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersScope] `json:"scope"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion() {
}

// Parameters for the Auth Rule
//
// Satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersAction]      `json:"action,required"`
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersCondition) MarshalJSON() (data []byte, err error) {
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
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersActionChallenge:
		return true
	}
	return false
}

// The scope the velocity is calculated for
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

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditionalBlock  AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeVelocityLimit     AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeMerchantLock      AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditionalAction AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeVelocityLimit, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeMerchantLock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditionalAction:
		return true
	}
	return false
}

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewParamsBodyEventStream string

const (
	AuthRuleV2NewParamsBodyEventStreamAuthorization         AuthRuleV2NewParamsBodyEventStream = "AUTHORIZATION"
	AuthRuleV2NewParamsBodyEventStreamThreeDSAuthentication AuthRuleV2NewParamsBodyEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2NewParamsBodyEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyEventStreamAuthorization, AuthRuleV2NewParamsBodyEventStreamThreeDSAuthentication:
		return true
	}
	return false
}

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_ACTION`: AUTHORIZATION or THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewParamsBodyType string

const (
	AuthRuleV2NewParamsBodyTypeConditionalBlock  AuthRuleV2NewParamsBodyType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyTypeVelocityLimit     AuthRuleV2NewParamsBodyType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyTypeMerchantLock      AuthRuleV2NewParamsBodyType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyTypeConditionalAction AuthRuleV2NewParamsBodyType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2NewParamsBodyType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyTypeConditionalBlock, AuthRuleV2NewParamsBodyTypeVelocityLimit, AuthRuleV2NewParamsBodyTypeMerchantLock, AuthRuleV2NewParamsBodyTypeConditionalAction:
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

// Satisfied by [AuthRuleV2UpdateParamsBodyObject],
// [AuthRuleV2UpdateParamsBodyObject], [AuthRuleV2UpdateParamsBodyCardLevelRule],
// [AuthRuleV2UpdateParamsBodyProgramLevelRule], [AuthRuleV2UpdateParamsBody].
type AuthRuleV2UpdateParamsBodyUnion interface {
	implementsAuthRuleV2UpdateParamsBodyUnion()
}

type AuthRuleV2UpdateParamsBodyObject struct {
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
	// The desired state of the Auth Rule.
	//
	// Note that only deactivating an Auth Rule through this endpoint is supported at
	// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
	// should be used to promote a draft to the currently active version.
	State param.Field[AuthRuleV2UpdateParamsBodyObjectState] `json:"state"`
}

func (r AuthRuleV2UpdateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2UpdateParamsBodyObject) implementsAuthRuleV2UpdateParamsBodyUnion() {}

// The desired state of the Auth Rule.
//
// Note that only deactivating an Auth Rule through this endpoint is supported at
// this time. If you need to (re-)activate an Auth Rule the /promote endpoint
// should be used to promote a draft to the currently active version.
type AuthRuleV2UpdateParamsBodyObjectState string

const (
	AuthRuleV2UpdateParamsBodyObjectStateInactive AuthRuleV2UpdateParamsBodyObjectState = "INACTIVE"
)

func (r AuthRuleV2UpdateParamsBodyObjectState) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateParamsBodyObjectStateInactive:
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
	// Only return Auth Rules that are bound to the provided account token.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Only return Auth Rules that are bound to the provided business account token.
	BusinessAccountToken param.Field[string] `query:"business_account_token" format:"uuid"`
	// Only return Auth Rules that are bound to the provided card token.
	CardToken param.Field[string] `query:"card_token" format:"uuid"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Only return Auth rules that are executed during the provided event stream.
	EventStream param.Field[AuthRuleV2ListParamsEventStream] `query:"event_stream"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// Only return Auth Rules that are bound to the provided scope.
	Scope param.Field[AuthRuleV2ListParamsScope] `query:"scope"`
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

// Only return Auth rules that are executed during the provided event stream.
type AuthRuleV2ListParamsEventStream string

const (
	AuthRuleV2ListParamsEventStreamAuthorization         AuthRuleV2ListParamsEventStream = "AUTHORIZATION"
	AuthRuleV2ListParamsEventStreamThreeDSAuthentication AuthRuleV2ListParamsEventStream = "THREE_DS_AUTHENTICATION"
)

func (r AuthRuleV2ListParamsEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListParamsEventStreamAuthorization, AuthRuleV2ListParamsEventStreamThreeDSAuthentication:
		return true
	}
	return false
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

type AuthRuleV2ApplyParams struct {
	Body AuthRuleV2ApplyParamsBodyUnion `json:"body,required"`
}

func (r AuthRuleV2ApplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AuthRuleV2ApplyParamsBody struct {
	AccountTokens         param.Field[interface{}] `json:"account_tokens"`
	BusinessAccountTokens param.Field[interface{}] `json:"business_account_tokens"`
	CardTokens            param.Field[interface{}] `json:"card_tokens"`
	ExcludedCardTokens    param.Field[interface{}] `json:"excluded_card_tokens"`
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
	AccountTokens param.Field[[]string] `json:"account_tokens" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens param.Field[[]string] `json:"business_account_tokens" format:"uuid"`
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
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2DraftParamsParametersAction] `json:"action"`
	Conditions param.Field[interface{}]                           `json:"conditions"`
	Filters    param.Field[interface{}]                           `json:"filters"`
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
	// DEPRECATED: This has been deprecated in favor of the Trailing Window Objects
	//
	// The size of the trailing window to calculate Spend Velocity over in seconds. The
	// minimum value is 10 seconds, and the maximum value is 2678400 seconds (31 days).
	Period param.Field[VelocityLimitParamsPeriodWindowUnion] `json:"period"`
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
// [AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParameters],
// [AuthRuleV2DraftParamsParameters].
type AuthRuleV2DraftParamsParametersUnion interface {
	implementsAuthRuleV2DraftParamsParametersUnion()
}

type AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersAction]      `json:"action,required"`
	Conditions param.Field[[]AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParameters) implementsAuthRuleV2DraftParamsParametersUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersAction string

const (
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersActionDecline   AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersAction = "DECLINE"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersActionChallenge AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersActionDecline, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersCondition struct {
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
	Attribute param.Field[AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersCondition) MarshalJSON() (data []byte, err error) {
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
type AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute string

const (
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeMcc                     AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCountry                 AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency                AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID              AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor              AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount              AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore               AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCardState               AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered              AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus               AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType              AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
)

func (r AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeMcc, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCountry, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCurrency, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeMerchantID, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeDescriptor, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCashAmount, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeRiskScore, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeCardState, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributePinEntered, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributePinStatus, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeWalletType, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation string

const (
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf                AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf             AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationMatches                AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch           AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo              AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_EQUAL_TO"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo           AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_NOT_EQUAL_TO"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan          AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_GREATER_THAN_OR_EQUAL_TO"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan             AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN"
	AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo    AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation = "IS_LESS_THAN_OR_EQUAL_TO"
)

func (r AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsOneOf, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationMatches, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsEqualTo, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsNotEqualTo, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsGreaterThanOrEqualTo, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThan, AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings].
type AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsValueUnion()
}

type AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2DraftParamsParametersConditionalAuthorizationActionParametersConditionsValueUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftParamsParametersAction string

const (
	AuthRuleV2DraftParamsParametersActionDecline   AuthRuleV2DraftParamsParametersAction = "DECLINE"
	AuthRuleV2DraftParamsParametersActionChallenge AuthRuleV2DraftParamsParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftParamsParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersActionDecline, AuthRuleV2DraftParamsParametersActionChallenge:
		return true
	}
	return false
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
	Begin param.Field[time.Time] `query:"begin,required" format:"date"`
	// End date for the report
	End param.Field[time.Time] `query:"end,required" format:"date"`
}

// URLQuery serializes [AuthRuleV2GetReportParams]'s query parameters as
// `url.Values`.
func (r AuthRuleV2GetReportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
