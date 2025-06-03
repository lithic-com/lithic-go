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

// Creates a new V2 Auth rule in draft mode
func (r *AuthRuleV2Service) New(ctx context.Context, body AuthRuleV2NewParams, opts ...option.RequestOption) (res *AuthRuleV2NewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/auth_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a V2 Auth rule by its token
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

// Updates a V2 Auth rule's properties
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

// Lists V2 Auth rules
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

// Lists V2 Auth rules
func (r *AuthRuleV2Service) ListAutoPaging(ctx context.Context, query AuthRuleV2ListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AuthRuleV2ListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a V2 Auth rule
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

// Associates a V2 Auth rule with a card program, the provided account(s) or
// card(s).
//
// Prefer using the `PATCH` method for this operation.
//
// Deprecated: deprecated
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

// Promotes the draft version of an Auth rule to the currently active version such
// that it is enforced in the respective stream.
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

// Requests a performance report of an Auth rule to be asynchronously generated.
// Reports can only be run on rules in draft or active mode and will included
// approved and declined statistics as well as examples. The generated report will
// be delivered asynchronously through a webhook with `event_type` =
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
	IncludeMccs []string                       `json:"include_mccs,nullable"`
	JSON        velocityLimitParamsFiltersJSON `json:"-"`
}

// velocityLimitParamsFiltersJSON contains the JSON metadata for the struct
// [VelocityLimitParamsFilters]
type velocityLimitParamsFiltersJSON struct {
	ExcludeCountries apijson.Field
	ExcludeMccs      apijson.Field
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
	// The type of event stream the Auth rule applies to.
	EventStream AuthRuleV2NewResponseEventStream `json:"event_stream,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2NewResponseState `json:"state,required"`
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
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
	EventStream        apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2NewResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]AuthRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                        `json:"period"`
	Scope  AuthRuleV2NewResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2NewResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2NewResponseCurrentVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParameters],
// [AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2NewResponseCurrentVersionParameters) AsUnion() AuthRuleV2NewResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParameters] or
// [AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2NewResponseCurrentVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2NewResponseCurrentVersionParametersMerchantLockParametersJSON contains
// the JSON metadata for the struct
// [AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParameters]
type authRuleV2NewResponseCurrentVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                          `json:"merchant_id"`
	JSON       authRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchant]
type authRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParameters]
type authRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersCondition]
type authRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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
	// [[]AuthRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                      `json:"period"`
	Scope  AuthRuleV2NewResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2NewResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2NewResponseDraftVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2NewResponseDraftVersionParametersMerchantLockParameters],
// [AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2NewResponseDraftVersionParameters) AsUnion() AuthRuleV2NewResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2NewResponseDraftVersionParametersMerchantLockParameters] or
// [AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseDraftVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2NewResponseDraftVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2NewResponseDraftVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2NewResponseDraftVersionParametersMerchantLockParametersJSON contains
// the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersionParametersMerchantLockParameters]
type authRuleV2NewResponseDraftVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2NewResponseDraftVersionParametersMerchantLockParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                        `json:"merchant_id"`
	JSON       authRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchant]
type authRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParameters]
type authRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersCondition]
type authRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of event stream the Auth rule applies to.
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewResponseType string

const (
	AuthRuleV2NewResponseTypeConditionalBlock     AuthRuleV2NewResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewResponseTypeVelocityLimit        AuthRuleV2NewResponseType = "VELOCITY_LIMIT"
	AuthRuleV2NewResponseTypeMerchantLock         AuthRuleV2NewResponseType = "MERCHANT_LOCK"
	AuthRuleV2NewResponseTypeConditional3DSAction AuthRuleV2NewResponseType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2NewResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseTypeConditionalBlock, AuthRuleV2NewResponseTypeVelocityLimit, AuthRuleV2NewResponseTypeMerchantLock, AuthRuleV2NewResponseTypeConditional3DSAction:
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
	// The type of event stream the Auth rule applies to.
	EventStream AuthRuleV2GetResponseEventStream `json:"event_stream,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2GetResponseState `json:"state,required"`
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
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
	EventStream        apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2GetResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]AuthRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                        `json:"period"`
	Scope  AuthRuleV2GetResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2GetResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2GetResponseCurrentVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParameters],
// [AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2GetResponseCurrentVersionParameters) AsUnion() AuthRuleV2GetResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParameters] or
// [AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2GetResponseCurrentVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2GetResponseCurrentVersionParametersMerchantLockParametersJSON contains
// the JSON metadata for the struct
// [AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParameters]
type authRuleV2GetResponseCurrentVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                          `json:"merchant_id"`
	JSON       authRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchant]
type authRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParameters]
type authRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersCondition]
type authRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2GetResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2GetResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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
	// [[]AuthRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                      `json:"period"`
	Scope  AuthRuleV2GetResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2GetResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2GetResponseDraftVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2GetResponseDraftVersionParametersMerchantLockParameters],
// [AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2GetResponseDraftVersionParameters) AsUnion() AuthRuleV2GetResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2GetResponseDraftVersionParametersMerchantLockParameters] or
// [AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseDraftVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2GetResponseDraftVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2GetResponseDraftVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2GetResponseDraftVersionParametersMerchantLockParametersJSON contains
// the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersionParametersMerchantLockParameters]
type authRuleV2GetResponseDraftVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetResponseDraftVersionParametersMerchantLockParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                        `json:"merchant_id"`
	JSON       authRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchant]
type authRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParameters]
type authRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersCondition]
type authRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2GetResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2GetResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of event stream the Auth rule applies to.
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2GetResponseType string

const (
	AuthRuleV2GetResponseTypeConditionalBlock     AuthRuleV2GetResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2GetResponseTypeVelocityLimit        AuthRuleV2GetResponseType = "VELOCITY_LIMIT"
	AuthRuleV2GetResponseTypeMerchantLock         AuthRuleV2GetResponseType = "MERCHANT_LOCK"
	AuthRuleV2GetResponseTypeConditional3DSAction AuthRuleV2GetResponseType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2GetResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseTypeConditionalBlock, AuthRuleV2GetResponseTypeVelocityLimit, AuthRuleV2GetResponseTypeMerchantLock, AuthRuleV2GetResponseTypeConditional3DSAction:
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
	// The type of event stream the Auth rule applies to.
	EventStream AuthRuleV2UpdateResponseEventStream `json:"event_stream,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2UpdateResponseState `json:"state,required"`
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
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
	EventStream        apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2UpdateResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                           `json:"period"`
	Scope  AuthRuleV2UpdateResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2UpdateResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2UpdateResponseCurrentVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParameters],
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2UpdateResponseCurrentVersionParameters) AsUnion() AuthRuleV2UpdateResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParameters] or
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParameters]
type authRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                             `json:"merchant_id"`
	JSON       authRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchant]
type authRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParameters]
type authRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersCondition]
type authRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2UpdateResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2UpdateResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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
	// [[]AuthRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                         `json:"period"`
	Scope  AuthRuleV2UpdateResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2UpdateResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2UpdateResponseDraftVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParameters],
// [AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2UpdateResponseDraftVersionParameters) AsUnion() AuthRuleV2UpdateResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParameters] or
// [AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParameters]
type authRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                           `json:"merchant_id"`
	JSON       authRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchant]
type authRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParameters]
type authRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersCondition]
type authRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2UpdateResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2UpdateResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of event stream the Auth rule applies to.
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2UpdateResponseType string

const (
	AuthRuleV2UpdateResponseTypeConditionalBlock     AuthRuleV2UpdateResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2UpdateResponseTypeVelocityLimit        AuthRuleV2UpdateResponseType = "VELOCITY_LIMIT"
	AuthRuleV2UpdateResponseTypeMerchantLock         AuthRuleV2UpdateResponseType = "MERCHANT_LOCK"
	AuthRuleV2UpdateResponseTypeConditional3DSAction AuthRuleV2UpdateResponseType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2UpdateResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseTypeConditionalBlock, AuthRuleV2UpdateResponseTypeVelocityLimit, AuthRuleV2UpdateResponseTypeMerchantLock, AuthRuleV2UpdateResponseTypeConditional3DSAction:
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
	// The type of event stream the Auth rule applies to.
	EventStream AuthRuleV2ListResponseEventStream `json:"event_stream,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2ListResponseState `json:"state,required"`
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
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
	EventStream        apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2ListResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]AuthRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                         `json:"period"`
	Scope  AuthRuleV2ListResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2ListResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2ListResponseCurrentVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParameters],
// [AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2ListResponseCurrentVersionParameters) AsUnion() AuthRuleV2ListResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParameters] or
// [AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2ListResponseCurrentVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2ListResponseCurrentVersionParametersMerchantLockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParameters]
type authRuleV2ListResponseCurrentVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                           `json:"merchant_id"`
	JSON       authRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchant]
type authRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParameters]
type authRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersCondition]
type authRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2ListResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2ListResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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
	// [[]AuthRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                       `json:"period"`
	Scope  AuthRuleV2ListResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2ListResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2ListResponseDraftVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2ListResponseDraftVersionParametersMerchantLockParameters],
// [AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2ListResponseDraftVersionParameters) AsUnion() AuthRuleV2ListResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2ListResponseDraftVersionParametersMerchantLockParameters] or
// [AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseDraftVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2ListResponseDraftVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2ListResponseDraftVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2ListResponseDraftVersionParametersMerchantLockParametersJSON contains
// the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersionParametersMerchantLockParameters]
type authRuleV2ListResponseDraftVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResponseDraftVersionParametersMerchantLockParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                         `json:"merchant_id"`
	JSON       authRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchant]
type authRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParameters]
type authRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersCondition]
type authRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2ListResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2ListResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of event stream the Auth rule applies to.
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2ListResponseType string

const (
	AuthRuleV2ListResponseTypeConditionalBlock     AuthRuleV2ListResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2ListResponseTypeVelocityLimit        AuthRuleV2ListResponseType = "VELOCITY_LIMIT"
	AuthRuleV2ListResponseTypeMerchantLock         AuthRuleV2ListResponseType = "MERCHANT_LOCK"
	AuthRuleV2ListResponseTypeConditional3DSAction AuthRuleV2ListResponseType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2ListResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseTypeConditionalBlock, AuthRuleV2ListResponseTypeVelocityLimit, AuthRuleV2ListResponseTypeMerchantLock, AuthRuleV2ListResponseTypeConditional3DSAction:
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
	// The type of event stream the Auth rule applies to.
	EventStream AuthRuleV2ApplyResponseEventStream `json:"event_stream,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2ApplyResponseState `json:"state,required"`
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
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
	EventStream        apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2ApplyResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                          `json:"period"`
	Scope  AuthRuleV2ApplyResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2ApplyResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2ApplyResponseCurrentVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParameters],
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2ApplyResponseCurrentVersionParameters) AsUnion() AuthRuleV2ApplyResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParameters] or
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParameters]
type authRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParameters) implementsAuthRuleV2ApplyResponseCurrentVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                            `json:"merchant_id"`
	JSON       authRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchant]
type authRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParameters]
type authRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParameters) implementsAuthRuleV2ApplyResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersCondition]
type authRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2ApplyResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2ApplyResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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
	// [[]AuthRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                        `json:"period"`
	Scope  AuthRuleV2ApplyResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2ApplyResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2ApplyResponseDraftVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParameters],
// [AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2ApplyResponseDraftVersionParameters) AsUnion() AuthRuleV2ApplyResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParameters] or
// [AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersJSON contains
// the JSON metadata for the struct
// [AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParameters]
type authRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParameters) implementsAuthRuleV2ApplyResponseDraftVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                          `json:"merchant_id"`
	JSON       authRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchant]
type authRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParameters]
type authRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParameters) implementsAuthRuleV2ApplyResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersCondition]
type authRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2ApplyResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2ApplyResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of event stream the Auth rule applies to.
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2ApplyResponseType string

const (
	AuthRuleV2ApplyResponseTypeConditionalBlock     AuthRuleV2ApplyResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2ApplyResponseTypeVelocityLimit        AuthRuleV2ApplyResponseType = "VELOCITY_LIMIT"
	AuthRuleV2ApplyResponseTypeMerchantLock         AuthRuleV2ApplyResponseType = "MERCHANT_LOCK"
	AuthRuleV2ApplyResponseTypeConditional3DSAction AuthRuleV2ApplyResponseType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2ApplyResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2ApplyResponseTypeConditionalBlock, AuthRuleV2ApplyResponseTypeVelocityLimit, AuthRuleV2ApplyResponseTypeMerchantLock, AuthRuleV2ApplyResponseTypeConditional3DSAction:
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
	// The type of event stream the Auth rule applies to.
	EventStream AuthRuleV2DraftResponseEventStream `json:"event_stream,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2DraftResponseState `json:"state,required"`
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
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
	EventStream        apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2DraftResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]AuthRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                          `json:"period"`
	Scope  AuthRuleV2DraftResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2DraftResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2DraftResponseCurrentVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParameters],
// [AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2DraftResponseCurrentVersionParameters) AsUnion() AuthRuleV2DraftResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParameters] or
// [AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParameters]
type authRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                            `json:"merchant_id"`
	JSON       authRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchant]
type authRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParameters]
type authRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersCondition]
type authRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2DraftResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2DraftResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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
	// [[]AuthRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                        `json:"period"`
	Scope  AuthRuleV2DraftResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2DraftResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2DraftResponseDraftVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParameters],
// [AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2DraftResponseDraftVersionParameters) AsUnion() AuthRuleV2DraftResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParameters] or
// [AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2DraftResponseDraftVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2DraftResponseDraftVersionParametersMerchantLockParametersJSON contains
// the JSON metadata for the struct
// [AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParameters]
type authRuleV2DraftResponseDraftVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                          `json:"merchant_id"`
	JSON       authRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchant]
type authRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParameters]
type authRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersCondition]
type authRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2DraftResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2DraftResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of event stream the Auth rule applies to.
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2DraftResponseType string

const (
	AuthRuleV2DraftResponseTypeConditionalBlock     AuthRuleV2DraftResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2DraftResponseTypeVelocityLimit        AuthRuleV2DraftResponseType = "VELOCITY_LIMIT"
	AuthRuleV2DraftResponseTypeMerchantLock         AuthRuleV2DraftResponseType = "MERCHANT_LOCK"
	AuthRuleV2DraftResponseTypeConditional3DSAction AuthRuleV2DraftResponseType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2DraftResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseTypeConditionalBlock, AuthRuleV2DraftResponseTypeVelocityLimit, AuthRuleV2DraftResponseTypeMerchantLock, AuthRuleV2DraftResponseTypeConditional3DSAction:
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
	// The type of event stream the Auth rule applies to.
	EventStream AuthRuleV2PromoteResponseEventStream `json:"event_stream,required"`
	// Auth Rule Name
	Name string `json:"name,required,nullable"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel bool `json:"program_level,required"`
	// The state of the Auth Rule
	State AuthRuleV2PromoteResponseState `json:"state,required"`
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
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
	EventStream        apijson.Field
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
	// The action to take if the conditions are met.
	Action AuthRuleV2PromoteResponseCurrentVersionParametersAction `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                            `json:"period"`
	Scope  AuthRuleV2PromoteResponseCurrentVersionParametersScope `json:"scope"`
	JSON   authRuleV2PromoteResponseCurrentVersionParametersJSON  `json:"-"`
	union  AuthRuleV2PromoteResponseCurrentVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParameters],
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2PromoteResponseCurrentVersionParameters) AsUnion() AuthRuleV2PromoteResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParameters] or
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParameters]
type authRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                              `json:"merchant_id"`
	JSON       authRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchant]
type authRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParameters]
type authRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersCondition]
type authRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2PromoteResponseCurrentVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2PromoteResponseCurrentVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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
	// [[]AuthRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersCondition].
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
	// This field can have the runtime type of
	// [[]AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchant].
	Merchants interface{} `json:"merchants"`
	// This field can have the runtime type of [VelocityLimitParamsPeriodUnion].
	Period interface{}                                          `json:"period"`
	Scope  AuthRuleV2PromoteResponseDraftVersionParametersScope `json:"scope"`
	JSON   authRuleV2PromoteResponseDraftVersionParametersJSON  `json:"-"`
	union  AuthRuleV2PromoteResponseDraftVersionParametersUnion
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
// [VelocityLimitParams],
// [AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParameters],
// [AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParameters].
func (r AuthRuleV2PromoteResponseDraftVersionParameters) AsUnion() AuthRuleV2PromoteResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParameters] or
// [AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParameters].
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParameters{}),
		},
	)
}

type AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants []AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchant `json:"merchants,required"`
	JSON      authRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersJSON       `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParameters]
type authRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersJSON struct {
	Merchants   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchant struct {
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
	MerchantID string                                                                            `json:"merchant_id"`
	JSON       authRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchantJSON `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchantJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchant]
type authRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchantJSON struct {
	Comment     apijson.Field
	Descriptor  apijson.Field
	MerchantID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionParametersMerchantLockParametersMerchantJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersAction      `json:"action,required"`
	Conditions []AuthRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersCondition `json:"conditions,required"`
	JSON       authRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersJSON        `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParameters]
type authRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {
}

// The action to take if the conditions are met.
type AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersActionDecline   AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersActionChallenge AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersActionDecline, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute `json:"attribute"`
	// The operation to apply to the attribute
	Operation AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion `json:"value"`
	JSON  authRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersConditionJSON        `json:"-"`
}

// authRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersConditionJSON
// contains the JSON metadata for the struct
// [AuthRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersCondition]
type authRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersConditionJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersConditionJSON) RawJSON() string {
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
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings{}),
		},
	)
}

type AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2PromoteResponseDraftVersionParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2PromoteResponseDraftVersionParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of event stream the Auth rule applies to.
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2PromoteResponseType string

const (
	AuthRuleV2PromoteResponseTypeConditionalBlock     AuthRuleV2PromoteResponseType = "CONDITIONAL_BLOCK"
	AuthRuleV2PromoteResponseTypeVelocityLimit        AuthRuleV2PromoteResponseType = "VELOCITY_LIMIT"
	AuthRuleV2PromoteResponseTypeMerchantLock         AuthRuleV2PromoteResponseType = "MERCHANT_LOCK"
	AuthRuleV2PromoteResponseTypeConditional3DSAction AuthRuleV2PromoteResponseType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2PromoteResponseType) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseTypeConditionalBlock, AuthRuleV2PromoteResponseTypeVelocityLimit, AuthRuleV2PromoteResponseTypeMerchantLock, AuthRuleV2PromoteResponseTypeConditional3DSAction:
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
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
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
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens) implementsAuthRuleV2NewParamsBodyUnion() {
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
	LimitCount param.Field[int64]                                                                    `json:"limit_count"`
	Merchants  param.Field[interface{}]                                                              `json:"merchants"`
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
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersMerchantLockParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersMerchantLockParametersMerchant] `json:"merchants,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersMerchantLockParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersMerchantLockParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersMerchantLockParametersMerchant struct {
	// A comment or explanation about the merchant, used internally for rule management
	// purposes.
	Comment param.Field[string] `json:"comment"`
	// Short description of the merchant, often used to provide more human-readable
	// context about the transaction merchant. This is typically the name or label
	// shown on transaction summaries.
	Descriptor param.Field[string] `json:"descriptor"`
	// Unique alphanumeric identifier for the payment card acceptor (merchant). This
	// attribute specifies the merchant entity that will be locked or referenced for
	// authorization rules.
	MerchantID param.Field[string] `json:"merchant_id"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersMerchantLockParametersMerchant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersAction]      `json:"action,required"`
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DsActionParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DsActionParametersCondition) MarshalJSON() (data []byte, err error) {
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
//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
//     fee field in the settlement/cardholder billing currency. This is the amount
//     the issuer should authorize against unless the issuer is paying the acquirer
//     fee on behalf of the cardholder.
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DsActionParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalBlock     AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeVelocityLimit        AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeMerchantLock         AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditional3DSAction AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeVelocityLimit, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeMerchantLock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditional3DSAction:
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
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokens) implementsAuthRuleV2NewParamsBodyUnion() {
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
	LimitCount param.Field[int64]                                                                 `json:"limit_count"`
	Merchants  param.Field[interface{}]                                                           `json:"merchants"`
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
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersMerchantLockParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersMerchantLockParametersMerchant] `json:"merchants,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersMerchantLockParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersMerchantLockParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersMerchantLockParametersMerchant struct {
	// A comment or explanation about the merchant, used internally for rule management
	// purposes.
	Comment param.Field[string] `json:"comment"`
	// Short description of the merchant, often used to provide more human-readable
	// context about the transaction merchant. This is typically the name or label
	// shown on transaction summaries.
	Descriptor param.Field[string] `json:"descriptor"`
	// Unique alphanumeric identifier for the payment card acceptor (merchant). This
	// attribute specifies the merchant entity that will be locked or referenced for
	// authorization rules.
	MerchantID param.Field[string] `json:"merchant_id"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersMerchantLockParametersMerchant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersAction]      `json:"action,required"`
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DsActionParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DsActionParametersCondition) MarshalJSON() (data []byte, err error) {
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
//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
//     fee field in the settlement/cardholder billing currency. This is the amount
//     the issuer should authorize against unless the issuer is paying the acquirer
//     fee on behalf of the cardholder.
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DsActionParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditionalBlock     AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeVelocityLimit        AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeMerchantLock         AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditional3DSAction AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeVelocityLimit, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeMerchantLock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestCardTokensTypeConditional3DSAction:
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
	// The type of Auth Rule. Effectively determines the event stream during which it
	// will be evaluated.
	//
	// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
	Type param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType] `json:"type"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevel) implementsAuthRuleV2NewParamsBodyUnion() {
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
	LimitCount param.Field[int64]                                                                   `json:"limit_count"`
	Merchants  param.Field[interface{}]                                                             `json:"merchants"`
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
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersMerchantLockParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParameters],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParameters].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion interface {
	implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersMerchantLockParametersMerchant] `json:"merchants,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersMerchantLockParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersMerchantLockParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersMerchantLockParametersMerchant struct {
	// A comment or explanation about the merchant, used internally for rule management
	// purposes.
	Comment param.Field[string] `json:"comment"`
	// Short description of the merchant, often used to provide more human-readable
	// context about the transaction merchant. This is typically the name or label
	// shown on transaction summaries.
	Descriptor param.Field[string] `json:"descriptor"`
	// Unique alphanumeric identifier for the payment card acceptor (merchant). This
	// attribute specifies the merchant entity that will be locked or referenced for
	// authorization rules.
	MerchantID param.Field[string] `json:"merchant_id"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersMerchantLockParametersMerchant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersAction]      `json:"action,required"`
	Conditions param.Field[[]AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DsActionParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParameters) implementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersActionDecline   AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersActionChallenge AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersActionDecline, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DsActionParametersCondition) MarshalJSON() (data []byte, err error) {
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
//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
//     fee field in the settlement/cardholder billing currency. This is the amount
//     the issuer should authorize against unless the issuer is paying the acquirer
//     fee on behalf of the cardholder.
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DsActionParametersConditionsValueUnion()
}

type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelParametersConditional3DsActionParametersConditionsValueUnion() {
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

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType string

const (
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditionalBlock     AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeVelocityLimit        AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeMerchantLock         AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditional3DSAction AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditionalBlock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeVelocityLimit, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeMerchantLock, AuthRuleV2NewParamsBodyCreateAuthRuleRequestProgramLevelTypeConditional3DSAction:
		return true
	}
	return false
}

// The type of Auth Rule. Effectively determines the event stream during which it
// will be evaluated.
//
// - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
// - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
// - `MERCHANT_LOCK`: AUTHORIZATION event stream.
// - `CONDITIONAL_3DS_ACTION`: THREE_DS_AUTHENTICATION event stream.
type AuthRuleV2NewParamsBodyType string

const (
	AuthRuleV2NewParamsBodyTypeConditionalBlock     AuthRuleV2NewParamsBodyType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyTypeVelocityLimit        AuthRuleV2NewParamsBodyType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyTypeMerchantLock         AuthRuleV2NewParamsBodyType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyTypeConditional3DSAction AuthRuleV2NewParamsBodyType = "CONDITIONAL_3DS_ACTION"
)

func (r AuthRuleV2NewParamsBodyType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyTypeConditionalBlock, AuthRuleV2NewParamsBodyTypeVelocityLimit, AuthRuleV2NewParamsBodyTypeMerchantLock, AuthRuleV2NewParamsBodyTypeConditional3DSAction:
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
	// Only return Auth Rules that are bound to the provided account token.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
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
	AuthRuleV2ListParamsScopeProgram AuthRuleV2ListParamsScope = "PROGRAM"
	AuthRuleV2ListParamsScopeAccount AuthRuleV2ListParamsScope = "ACCOUNT"
	AuthRuleV2ListParamsScopeCard    AuthRuleV2ListParamsScope = "CARD"
)

func (r AuthRuleV2ListParamsScope) IsKnown() bool {
	switch r {
	case AuthRuleV2ListParamsScopeProgram, AuthRuleV2ListParamsScopeAccount, AuthRuleV2ListParamsScopeCard:
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
	LimitCount param.Field[int64]                                `json:"limit_count"`
	Merchants  param.Field[interface{}]                          `json:"merchants"`
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
// [AuthRuleV2DraftParamsParametersMerchantLockParameters],
// [AuthRuleV2DraftParamsParametersConditional3DSActionParameters],
// [AuthRuleV2DraftParamsParameters].
type AuthRuleV2DraftParamsParametersUnion interface {
	implementsAuthRuleV2DraftParamsParametersUnion()
}

type AuthRuleV2DraftParamsParametersMerchantLockParameters struct {
	// A list of merchant locks defining specific merchants or groups of merchants
	// (based on descriptors or IDs) that the lock applies to.
	Merchants param.Field[[]AuthRuleV2DraftParamsParametersMerchantLockParametersMerchant] `json:"merchants,required"`
}

func (r AuthRuleV2DraftParamsParametersMerchantLockParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2DraftParamsParametersMerchantLockParameters) implementsAuthRuleV2DraftParamsParametersUnion() {
}

// Represents a specific merchant lock based on their ID or descriptor. Each
// merchant object allows transaction rules to work at a granular level and
// requires at least one of merchant_id or descriptor.
type AuthRuleV2DraftParamsParametersMerchantLockParametersMerchant struct {
	// A comment or explanation about the merchant, used internally for rule management
	// purposes.
	Comment param.Field[string] `json:"comment"`
	// Short description of the merchant, often used to provide more human-readable
	// context about the transaction merchant. This is typically the name or label
	// shown on transaction summaries.
	Descriptor param.Field[string] `json:"descriptor"`
	// Unique alphanumeric identifier for the payment card acceptor (merchant). This
	// attribute specifies the merchant entity that will be locked or referenced for
	// authorization rules.
	MerchantID param.Field[string] `json:"merchant_id"`
}

func (r AuthRuleV2DraftParamsParametersMerchantLockParametersMerchant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRuleV2DraftParamsParametersConditional3DSActionParameters struct {
	// The action to take if the conditions are met.
	Action     param.Field[AuthRuleV2DraftParamsParametersConditional3DSActionParametersAction]      `json:"action,required"`
	Conditions param.Field[[]AuthRuleV2DraftParamsParametersConditional3DsActionParametersCondition] `json:"conditions,required"`
}

func (r AuthRuleV2DraftParamsParametersConditional3DSActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2DraftParamsParametersConditional3DSActionParameters) implementsAuthRuleV2DraftParamsParametersUnion() {
}

// The action to take if the conditions are met.
type AuthRuleV2DraftParamsParametersConditional3DSActionParametersAction string

const (
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersActionDecline   AuthRuleV2DraftParamsParametersConditional3DSActionParametersAction = "DECLINE"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersActionChallenge AuthRuleV2DraftParamsParametersConditional3DSActionParametersAction = "CHALLENGE"
)

func (r AuthRuleV2DraftParamsParametersConditional3DSActionParametersAction) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersConditional3DSActionParametersActionDecline, AuthRuleV2DraftParamsParametersConditional3DSActionParametersActionChallenge:
		return true
	}
	return false
}

type AuthRuleV2DraftParamsParametersConditional3DsActionParametersCondition struct {
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
	//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
	//     given authentication. Scores are on a range of 0-999, with 0 representing the
	//     lowest risk and 999 representing the highest risk. For Visa transactions,
	//     where the raw score has a range of 0-99, Lithic will normalize the score by
	//     multiplying the raw score by 10x.
	//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
	Attribute param.Field[AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute] `json:"attribute"`
	// The operation to apply to the attribute
	Operation param.Field[AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation] `json:"operation"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsValueUnion] `json:"value"`
}

func (r AuthRuleV2DraftParamsParametersConditional3DsActionParametersCondition) MarshalJSON() (data []byte, err error) {
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
//   - `TRANSACTION_AMOUNT`: The base transaction amount (in cents) plus the acquirer
//     fee field in the settlement/cardholder billing currency. This is the amount
//     the issuer should authorize against unless the issuer is paying the acquirer
//     fee on behalf of the cardholder.
//   - `RISK_SCORE`: Network-provided score assessing risk level associated with a
//     given authentication. Scores are on a range of 0-999, with 0 representing the
//     lowest risk and 999 representing the highest risk. For Visa transactions,
//     where the raw score has a range of 0-99, Lithic will normalize the score by
//     multiplying the raw score by 10x.
//   - `MESSAGE_CATEGORY`: The category of the authentication being processed.
type AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute string

const (
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeMcc               AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute = "MCC"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeCountry           AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute = "COUNTRY"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeCurrency          AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute = "CURRENCY"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeMerchantID        AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute = "MERCHANT_ID"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeDescriptor        AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute = "DESCRIPTOR"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeTransactionAmount AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeRiskScore         AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute = "RISK_SCORE"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeMessageCategory   AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute = "MESSAGE_CATEGORY"
)

func (r AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeMcc, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeCountry, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeCurrency, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeMerchantID, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeDescriptor, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeTransactionAmount, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeRiskScore, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsAttributeMessageCategory:
		return true
	}
	return false
}

// The operation to apply to the attribute
type AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation string

const (
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationIsOneOf       AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation = "IS_ONE_OF"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationIsNotOneOf    AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation = "IS_NOT_ONE_OF"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationMatches       AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation = "MATCHES"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationDoesNotMatch  AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation = "DOES_NOT_MATCH"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationIsGreaterThan AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation = "IS_GREATER_THAN"
	AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationIsLessThan    AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation = "IS_LESS_THAN"
)

func (r AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperation) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationIsOneOf, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationIsNotOneOf, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationMatches, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationDoesNotMatch, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationIsGreaterThan, AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsOperationIsLessThan:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsValueListOfStrings].
type AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsValueUnion interface {
	ImplementsAuthRuleV2DraftParamsParametersConditional3DsActionParametersConditionsValueUnion()
}

type AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsValueListOfStrings []string

func (r AuthRuleV2DraftParamsParametersConditional3DSActionParametersConditionsValueListOfStrings) ImplementsAuthRuleV2DraftParamsParametersConditional3DsActionParametersConditionsValueUnion() {
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
