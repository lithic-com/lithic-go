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
	//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
	//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
	//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
	Attribute ConditionalAttribute `json:"attribute,required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation,required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion `json:"value,required"`
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
	Attribute param.Field[ConditionalAttribute] `json:"attribute,required"`
	// The operation to apply to the attribute
	Operation param.Field[ConditionalOperation] `json:"operation,required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value param.Field[ConditionalValueUnionParam] `json:"value,required"`
}

func (r AuthRuleConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	//   - `ADDRESS_MATCH`: Lithic's evaluation result comparing transaction's address
	//     data with the cardholder KYC data if it exists. Valid values are `MATCH`,
	//     `MATCH_ADDRESS_ONLY`, `MATCH_ZIP_ONLY`,`MISMATCH`,`NOT_PRESENT`.
	Attribute Conditional3DSActionParametersConditionsAttribute `json:"attribute,required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation,required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion                       `json:"value,required"`
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
	// The action to take if the conditions are met
	Action     ConditionalACHActionParametersAction      `json:"action,required"`
	Conditions []ConditionalACHActionParametersCondition `json:"conditions,required"`
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

func (r ConditionalACHActionParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {
}

func (r ConditionalACHActionParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {}

func (r ConditionalACHActionParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {
}

func (r ConditionalACHActionParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {}

// The action to take if the conditions are met
type ConditionalACHActionParametersAction struct {
	// Approve the ACH transaction
	Type ConditionalACHActionParametersActionType `json:"type,required"`
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
// [ConditionalACHActionParametersActionApproveAction],
// [ConditionalACHActionParametersActionReturnAction].
func (r ConditionalACHActionParametersAction) AsUnion() ConditionalACHActionParametersActionUnion {
	return r.union
}

// The action to take if the conditions are met
//
// Union satisfied by [ConditionalACHActionParametersActionApproveAction] or
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
			Type:       reflect.TypeOf(ConditionalACHActionParametersActionApproveAction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalACHActionParametersActionReturnAction{}),
		},
	)
}

type ConditionalACHActionParametersActionApproveAction struct {
	// Approve the ACH transaction
	Type ConditionalACHActionParametersActionApproveActionType `json:"type,required"`
	JSON conditionalACHActionParametersActionApproveActionJSON `json:"-"`
}

// conditionalACHActionParametersActionApproveActionJSON contains the JSON metadata
// for the struct [ConditionalACHActionParametersActionApproveAction]
type conditionalACHActionParametersActionApproveActionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalACHActionParametersActionApproveAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalACHActionParametersActionApproveActionJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalACHActionParametersActionApproveAction) implementsConditionalACHActionParametersAction() {
}

// Approve the ACH transaction
type ConditionalACHActionParametersActionApproveActionType string

const (
	ConditionalACHActionParametersActionApproveActionTypeApprove ConditionalACHActionParametersActionApproveActionType = "APPROVE"
)

func (r ConditionalACHActionParametersActionApproveActionType) IsKnown() bool {
	switch r {
	case ConditionalACHActionParametersActionApproveActionTypeApprove:
		return true
	}
	return false
}

type ConditionalACHActionParametersActionReturnAction struct {
	// NACHA return code to use when returning the transaction. Note that the list of
	// available return codes is subject to an allowlist configured at the program
	// level
	Code ConditionalACHActionParametersActionReturnActionCode `json:"code,required"`
	// Return the ACH transaction
	Type ConditionalACHActionParametersActionReturnActionType `json:"type,required"`
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
	Attribute ConditionalACHActionParametersConditionsAttribute `json:"attribute,required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation,required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion                       `json:"value,required"`
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
	Action     ConditionalAuthorizationActionParametersAction      `json:"action,required"`
	Conditions []ConditionalAuthorizationActionParametersCondition `json:"conditions,required"`
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

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {
}

func (r ConditionalAuthorizationActionParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {
}

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
	Attribute ConditionalAuthorizationActionParametersConditionsAttribute `json:"attribute,required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation,required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion                                 `json:"value,required"`
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
type ConditionalAuthorizationActionParametersConditionsAttribute string

const (
	ConditionalAuthorizationActionParametersConditionsAttributeMcc                     ConditionalAuthorizationActionParametersConditionsAttribute = "MCC"
	ConditionalAuthorizationActionParametersConditionsAttributeCountry                 ConditionalAuthorizationActionParametersConditionsAttribute = "COUNTRY"
	ConditionalAuthorizationActionParametersConditionsAttributeCurrency                ConditionalAuthorizationActionParametersConditionsAttribute = "CURRENCY"
	ConditionalAuthorizationActionParametersConditionsAttributeMerchantID              ConditionalAuthorizationActionParametersConditionsAttribute = "MERCHANT_ID"
	ConditionalAuthorizationActionParametersConditionsAttributeDescriptor              ConditionalAuthorizationActionParametersConditionsAttribute = "DESCRIPTOR"
	ConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift          ConditionalAuthorizationActionParametersConditionsAttribute = "LIABILITY_SHIFT"
	ConditionalAuthorizationActionParametersConditionsAttributePanEntryMode            ConditionalAuthorizationActionParametersConditionsAttribute = "PAN_ENTRY_MODE"
	ConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount       ConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_AMOUNT"
	ConditionalAuthorizationActionParametersConditionsAttributeCashAmount              ConditionalAuthorizationActionParametersConditionsAttribute = "CASH_AMOUNT"
	ConditionalAuthorizationActionParametersConditionsAttributeRiskScore               ConditionalAuthorizationActionParametersConditionsAttribute = "RISK_SCORE"
	ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_15M"
	ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H  ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_1H"
	ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_TRANSACTION_COUNT_24H"
	ConditionalAuthorizationActionParametersConditionsAttributeCardState               ConditionalAuthorizationActionParametersConditionsAttribute = "CARD_STATE"
	ConditionalAuthorizationActionParametersConditionsAttributePinEntered              ConditionalAuthorizationActionParametersConditionsAttribute = "PIN_ENTERED"
	ConditionalAuthorizationActionParametersConditionsAttributePinStatus               ConditionalAuthorizationActionParametersConditionsAttribute = "PIN_STATUS"
	ConditionalAuthorizationActionParametersConditionsAttributeWalletType              ConditionalAuthorizationActionParametersConditionsAttribute = "WALLET_TYPE"
	ConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator    ConditionalAuthorizationActionParametersConditionsAttribute = "TRANSACTION_INITIATOR"
	ConditionalAuthorizationActionParametersConditionsAttributeAddressMatch            ConditionalAuthorizationActionParametersConditionsAttribute = "ADDRESS_MATCH"
)

func (r ConditionalAuthorizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case ConditionalAuthorizationActionParametersConditionsAttributeMcc, ConditionalAuthorizationActionParametersConditionsAttributeCountry, ConditionalAuthorizationActionParametersConditionsAttributeCurrency, ConditionalAuthorizationActionParametersConditionsAttributeMerchantID, ConditionalAuthorizationActionParametersConditionsAttributeDescriptor, ConditionalAuthorizationActionParametersConditionsAttributeLiabilityShift, ConditionalAuthorizationActionParametersConditionsAttributePanEntryMode, ConditionalAuthorizationActionParametersConditionsAttributeTransactionAmount, ConditionalAuthorizationActionParametersConditionsAttributeCashAmount, ConditionalAuthorizationActionParametersConditionsAttributeRiskScore, ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount15M, ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount1H, ConditionalAuthorizationActionParametersConditionsAttributeCardTransactionCount24H, ConditionalAuthorizationActionParametersConditionsAttributeCardState, ConditionalAuthorizationActionParametersConditionsAttributePinEntered, ConditionalAuthorizationActionParametersConditionsAttributePinStatus, ConditionalAuthorizationActionParametersConditionsAttributeWalletType, ConditionalAuthorizationActionParametersConditionsAttributeTransactionInitiator, ConditionalAuthorizationActionParametersConditionsAttributeAddressMatch:
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

func (r ConditionalBlockParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {}

func (r ConditionalBlockParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {}

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
)

func (r ConditionalOperation) IsKnown() bool {
	switch r {
	case ConditionalOperationIsOneOf, ConditionalOperationIsNotOneOf, ConditionalOperationMatches, ConditionalOperationDoesNotMatch, ConditionalOperationIsEqualTo, ConditionalOperationIsNotEqualTo, ConditionalOperationIsGreaterThan, ConditionalOperationIsGreaterThanOrEqualTo, ConditionalOperationIsLessThan, ConditionalOperationIsLessThanOrEqualTo:
		return true
	}
	return false
}

type ConditionalTokenizationActionParameters struct {
	// The action to take if the conditions are met
	Action     ConditionalTokenizationActionParametersAction      `json:"action,required"`
	Conditions []ConditionalTokenizationActionParametersCondition `json:"conditions,required"`
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

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2NewResponseCurrentVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2NewResponseDraftVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2GetResponseCurrentVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2GetResponseDraftVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2UpdateResponseCurrentVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2UpdateResponseDraftVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2ListResponseCurrentVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2ListResponseDraftVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2DraftResponseCurrentVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2DraftResponseDraftVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2PromoteResponseCurrentVersionParameters() {
}

func (r ConditionalTokenizationActionParameters) implementsAuthRuleV2PromoteResponseDraftVersionParameters() {
}

// The action to take if the conditions are met
type ConditionalTokenizationActionParametersAction struct {
	// Decline the tokenization request
	Type ConditionalTokenizationActionParametersActionType `json:"type,required"`
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
// [ConditionalTokenizationActionParametersActionDeclineAction],
// [ConditionalTokenizationActionParametersActionRequireTfaAction].
func (r ConditionalTokenizationActionParametersAction) AsUnion() ConditionalTokenizationActionParametersActionUnion {
	return r.union
}

// The action to take if the conditions are met
//
// Union satisfied by [ConditionalTokenizationActionParametersActionDeclineAction]
// or [ConditionalTokenizationActionParametersActionRequireTfaAction].
type ConditionalTokenizationActionParametersActionUnion interface {
	implementsConditionalTokenizationActionParametersAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConditionalTokenizationActionParametersActionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalTokenizationActionParametersActionDeclineAction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionalTokenizationActionParametersActionRequireTfaAction{}),
		},
	)
}

type ConditionalTokenizationActionParametersActionDeclineAction struct {
	// Decline the tokenization request
	Type ConditionalTokenizationActionParametersActionDeclineActionType `json:"type,required"`
	// Reason code for declining the tokenization request
	Reason ConditionalTokenizationActionParametersActionDeclineActionReason `json:"reason"`
	JSON   conditionalTokenizationActionParametersActionDeclineActionJSON   `json:"-"`
}

// conditionalTokenizationActionParametersActionDeclineActionJSON contains the JSON
// metadata for the struct
// [ConditionalTokenizationActionParametersActionDeclineAction]
type conditionalTokenizationActionParametersActionDeclineActionJSON struct {
	Type        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConditionalTokenizationActionParametersActionDeclineAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionalTokenizationActionParametersActionDeclineActionJSON) RawJSON() string {
	return r.raw
}

func (r ConditionalTokenizationActionParametersActionDeclineAction) implementsConditionalTokenizationActionParametersAction() {
}

// Decline the tokenization request
type ConditionalTokenizationActionParametersActionDeclineActionType string

const (
	ConditionalTokenizationActionParametersActionDeclineActionTypeDecline ConditionalTokenizationActionParametersActionDeclineActionType = "DECLINE"
)

func (r ConditionalTokenizationActionParametersActionDeclineActionType) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersActionDeclineActionTypeDecline:
		return true
	}
	return false
}

// Reason code for declining the tokenization request
type ConditionalTokenizationActionParametersActionDeclineActionReason string

const (
	ConditionalTokenizationActionParametersActionDeclineActionReasonAccountScore1                  ConditionalTokenizationActionParametersActionDeclineActionReason = "ACCOUNT_SCORE_1"
	ConditionalTokenizationActionParametersActionDeclineActionReasonDeviceScore1                   ConditionalTokenizationActionParametersActionDeclineActionReason = "DEVICE_SCORE_1"
	ConditionalTokenizationActionParametersActionDeclineActionReasonAllWalletDeclineReasonsPresent ConditionalTokenizationActionParametersActionDeclineActionReason = "ALL_WALLET_DECLINE_REASONS_PRESENT"
	ConditionalTokenizationActionParametersActionDeclineActionReasonWalletRecommendedDecisionRed   ConditionalTokenizationActionParametersActionDeclineActionReason = "WALLET_RECOMMENDED_DECISION_RED"
	ConditionalTokenizationActionParametersActionDeclineActionReasonCvcMismatch                    ConditionalTokenizationActionParametersActionDeclineActionReason = "CVC_MISMATCH"
	ConditionalTokenizationActionParametersActionDeclineActionReasonCardExpiryMonthMismatch        ConditionalTokenizationActionParametersActionDeclineActionReason = "CARD_EXPIRY_MONTH_MISMATCH"
	ConditionalTokenizationActionParametersActionDeclineActionReasonCardExpiryYearMismatch         ConditionalTokenizationActionParametersActionDeclineActionReason = "CARD_EXPIRY_YEAR_MISMATCH"
	ConditionalTokenizationActionParametersActionDeclineActionReasonCardInvalidState               ConditionalTokenizationActionParametersActionDeclineActionReason = "CARD_INVALID_STATE"
	ConditionalTokenizationActionParametersActionDeclineActionReasonCustomerRedPath                ConditionalTokenizationActionParametersActionDeclineActionReason = "CUSTOMER_RED_PATH"
	ConditionalTokenizationActionParametersActionDeclineActionReasonInvalidCustomerResponse        ConditionalTokenizationActionParametersActionDeclineActionReason = "INVALID_CUSTOMER_RESPONSE"
	ConditionalTokenizationActionParametersActionDeclineActionReasonNetworkFailure                 ConditionalTokenizationActionParametersActionDeclineActionReason = "NETWORK_FAILURE"
	ConditionalTokenizationActionParametersActionDeclineActionReasonGenericDecline                 ConditionalTokenizationActionParametersActionDeclineActionReason = "GENERIC_DECLINE"
	ConditionalTokenizationActionParametersActionDeclineActionReasonDigitalCardArtRequired         ConditionalTokenizationActionParametersActionDeclineActionReason = "DIGITAL_CARD_ART_REQUIRED"
)

func (r ConditionalTokenizationActionParametersActionDeclineActionReason) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersActionDeclineActionReasonAccountScore1, ConditionalTokenizationActionParametersActionDeclineActionReasonDeviceScore1, ConditionalTokenizationActionParametersActionDeclineActionReasonAllWalletDeclineReasonsPresent, ConditionalTokenizationActionParametersActionDeclineActionReasonWalletRecommendedDecisionRed, ConditionalTokenizationActionParametersActionDeclineActionReasonCvcMismatch, ConditionalTokenizationActionParametersActionDeclineActionReasonCardExpiryMonthMismatch, ConditionalTokenizationActionParametersActionDeclineActionReasonCardExpiryYearMismatch, ConditionalTokenizationActionParametersActionDeclineActionReasonCardInvalidState, ConditionalTokenizationActionParametersActionDeclineActionReasonCustomerRedPath, ConditionalTokenizationActionParametersActionDeclineActionReasonInvalidCustomerResponse, ConditionalTokenizationActionParametersActionDeclineActionReasonNetworkFailure, ConditionalTokenizationActionParametersActionDeclineActionReasonGenericDecline, ConditionalTokenizationActionParametersActionDeclineActionReasonDigitalCardArtRequired:
		return true
	}
	return false
}

type ConditionalTokenizationActionParametersActionRequireTfaAction struct {
	// Require two-factor authentication for the tokenization request
	Type ConditionalTokenizationActionParametersActionRequireTfaActionType `json:"type,required"`
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
	//     initiated (e.g., DIGITAL_WALLET, ECOMMERCE).
	//   - `TOKENIZATION_SOURCE`: The source of the tokenization request.
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
	//   - `TOKEN_REQUESTOR_ID`: Unique identifier for the entity requesting the token.
	//   - `WALLET_TOKEN_STATUS`: The current status of the wallet token.
	Attribute ConditionalTokenizationActionParametersConditionsAttribute `json:"attribute,required"`
	// The operation to apply to the attribute
	Operation ConditionalOperation `json:"operation,required"`
	// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
	Value ConditionalValueUnion                                `json:"value,required"`
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
//     initiated (e.g., DIGITAL_WALLET, ECOMMERCE).
//   - `TOKENIZATION_SOURCE`: The source of the tokenization request.
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
//   - `TOKEN_REQUESTOR_ID`: Unique identifier for the entity requesting the token.
//   - `WALLET_TOKEN_STATUS`: The current status of the wallet token.
type ConditionalTokenizationActionParametersConditionsAttribute string

const (
	ConditionalTokenizationActionParametersConditionsAttributeTimestamp                 ConditionalTokenizationActionParametersConditionsAttribute = "TIMESTAMP"
	ConditionalTokenizationActionParametersConditionsAttributeTokenizationChannel       ConditionalTokenizationActionParametersConditionsAttribute = "TOKENIZATION_CHANNEL"
	ConditionalTokenizationActionParametersConditionsAttributeTokenizationSource        ConditionalTokenizationActionParametersConditionsAttribute = "TOKENIZATION_SOURCE"
	ConditionalTokenizationActionParametersConditionsAttributeTokenRequestorName        ConditionalTokenizationActionParametersConditionsAttribute = "TOKEN_REQUESTOR_NAME"
	ConditionalTokenizationActionParametersConditionsAttributeWalletAccountScore        ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_ACCOUNT_SCORE"
	ConditionalTokenizationActionParametersConditionsAttributeWalletDeviceScore         ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_DEVICE_SCORE"
	ConditionalTokenizationActionParametersConditionsAttributeWalletRecommendedDecision ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_RECOMMENDED_DECISION"
	ConditionalTokenizationActionParametersConditionsAttributeTokenRequestorID          ConditionalTokenizationActionParametersConditionsAttribute = "TOKEN_REQUESTOR_ID"
	ConditionalTokenizationActionParametersConditionsAttributeWalletTokenStatus         ConditionalTokenizationActionParametersConditionsAttribute = "WALLET_TOKEN_STATUS"
)

func (r ConditionalTokenizationActionParametersConditionsAttribute) IsKnown() bool {
	switch r {
	case ConditionalTokenizationActionParametersConditionsAttributeTimestamp, ConditionalTokenizationActionParametersConditionsAttributeTokenizationChannel, ConditionalTokenizationActionParametersConditionsAttributeTokenizationSource, ConditionalTokenizationActionParametersConditionsAttributeTokenRequestorName, ConditionalTokenizationActionParametersConditionsAttributeWalletAccountScore, ConditionalTokenizationActionParametersConditionsAttributeWalletDeviceScore, ConditionalTokenizationActionParametersConditionsAttributeWalletRecommendedDecision, ConditionalTokenizationActionParametersConditionsAttributeTokenRequestorID, ConditionalTokenizationActionParametersConditionsAttributeWalletTokenStatus:
		return true
	}
	return false
}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt] or
// [ConditionalValueListOfStrings].
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
	)
}

type ConditionalValueListOfStrings []string

func (r ConditionalValueListOfStrings) ImplementsConditionalValueUnion() {}

// A regex string, to be used with `MATCHES` or `DOES_NOT_MATCH`
//
// Satisfied by [shared.UnionString], [shared.UnionInt],
// [ConditionalValueListOfStringsParam].
type ConditionalValueUnionParam interface {
	ImplementsConditionalValueUnionParam()
}

type ConditionalValueListOfStringsParam []string

func (r ConditionalValueListOfStringsParam) ImplementsConditionalValueUnionParam() {}

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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period,required"`
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

// Velocity over the current day since 00:00 / 12 AM in Eastern Time
type VelocityLimitPeriod struct {
	Type VelocityLimitPeriodType `json:"type,required"`
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
	Duration int64                                       `json:"duration,required"`
	Type     VelocityLimitPeriodTrailingWindowObjectType `json:"type,required"`
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
	Type VelocityLimitPeriodFixedWindowDayType `json:"type,required"`
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
	Type VelocityLimitPeriodFixedWindowWeekType `json:"type,required"`
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
	Type VelocityLimitPeriodFixedWindowMonthType `json:"type,required"`
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
	Type VelocityLimitPeriodFixedWindowYearType `json:"type,required"`
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
	Type param.Field[VelocityLimitPeriodType] `json:"type,required"`
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
	Duration param.Field[int64]                                       `json:"duration,required"`
	Type     param.Field[VelocityLimitPeriodTrailingWindowObjectType] `json:"type,required"`
}

func (r VelocityLimitPeriodTrailingWindowObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VelocityLimitPeriodTrailingWindowObjectParam) implementsVelocityLimitPeriodUnionParam() {}

// Velocity over the current day since 00:00 / 12 AM in Eastern Time
type VelocityLimitPeriodFixedWindowDayParam struct {
	Type param.Field[VelocityLimitPeriodFixedWindowDayType] `json:"type,required"`
}

func (r VelocityLimitPeriodFixedWindowDayParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VelocityLimitPeriodFixedWindowDayParam) implementsVelocityLimitPeriodUnionParam() {}

// Velocity over the current week since 00:00 / 12 AM in Eastern Time on specified
// `day_of_week`
type VelocityLimitPeriodFixedWindowWeekParam struct {
	Type param.Field[VelocityLimitPeriodFixedWindowWeekType] `json:"type,required"`
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
	Type param.Field[VelocityLimitPeriodFixedWindowMonthType] `json:"type,required"`
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
	Type param.Field[VelocityLimitPeriodFixedWindowYearType] `json:"type,required"`
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
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2NewResponseCurrentVersionParameters) AsUnion() AuthRuleV2NewResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2NewResponseDraftVersionParameters) AsUnion() AuthRuleV2NewResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	AuthRuleV2NewResponseEventStreamTokenization          AuthRuleV2NewResponseEventStream = "TOKENIZATION"
	AuthRuleV2NewResponseEventStreamACHCreditReceipt      AuthRuleV2NewResponseEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2NewResponseEventStreamACHDebitReceipt       AuthRuleV2NewResponseEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2NewResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewResponseEventStreamAuthorization, AuthRuleV2NewResponseEventStreamThreeDSAuthentication, AuthRuleV2NewResponseEventStreamTokenization, AuthRuleV2NewResponseEventStreamACHCreditReceipt, AuthRuleV2NewResponseEventStreamACHDebitReceipt:
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2GetResponseCurrentVersionParameters) AsUnion() AuthRuleV2GetResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2GetResponseDraftVersionParameters) AsUnion() AuthRuleV2GetResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	AuthRuleV2GetResponseEventStreamTokenization          AuthRuleV2GetResponseEventStream = "TOKENIZATION"
	AuthRuleV2GetResponseEventStreamACHCreditReceipt      AuthRuleV2GetResponseEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2GetResponseEventStreamACHDebitReceipt       AuthRuleV2GetResponseEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2GetResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2GetResponseEventStreamAuthorization, AuthRuleV2GetResponseEventStreamThreeDSAuthentication, AuthRuleV2GetResponseEventStreamTokenization, AuthRuleV2GetResponseEventStreamACHCreditReceipt, AuthRuleV2GetResponseEventStreamACHDebitReceipt:
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2UpdateResponseCurrentVersionParameters) AsUnion() AuthRuleV2UpdateResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2UpdateResponseDraftVersionParameters) AsUnion() AuthRuleV2UpdateResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	AuthRuleV2UpdateResponseEventStreamTokenization          AuthRuleV2UpdateResponseEventStream = "TOKENIZATION"
	AuthRuleV2UpdateResponseEventStreamACHCreditReceipt      AuthRuleV2UpdateResponseEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2UpdateResponseEventStreamACHDebitReceipt       AuthRuleV2UpdateResponseEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2UpdateResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2UpdateResponseEventStreamAuthorization, AuthRuleV2UpdateResponseEventStreamThreeDSAuthentication, AuthRuleV2UpdateResponseEventStreamTokenization, AuthRuleV2UpdateResponseEventStreamACHCreditReceipt, AuthRuleV2UpdateResponseEventStreamACHDebitReceipt:
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2ListResponseCurrentVersionParameters) AsUnion() AuthRuleV2ListResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2ListResponseDraftVersionParameters) AsUnion() AuthRuleV2ListResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	AuthRuleV2ListResponseEventStreamTokenization          AuthRuleV2ListResponseEventStream = "TOKENIZATION"
	AuthRuleV2ListResponseEventStreamACHCreditReceipt      AuthRuleV2ListResponseEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2ListResponseEventStreamACHDebitReceipt       AuthRuleV2ListResponseEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2ListResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListResponseEventStreamAuthorization, AuthRuleV2ListResponseEventStreamThreeDSAuthentication, AuthRuleV2ListResponseEventStreamTokenization, AuthRuleV2ListResponseEventStreamACHCreditReceipt, AuthRuleV2ListResponseEventStreamACHDebitReceipt:
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2DraftResponseCurrentVersionParameters) AsUnion() AuthRuleV2DraftResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2DraftResponseDraftVersionParameters) AsUnion() AuthRuleV2DraftResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	AuthRuleV2DraftResponseEventStreamTokenization          AuthRuleV2DraftResponseEventStream = "TOKENIZATION"
	AuthRuleV2DraftResponseEventStreamACHCreditReceipt      AuthRuleV2DraftResponseEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2DraftResponseEventStreamACHDebitReceipt       AuthRuleV2DraftResponseEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2DraftResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2DraftResponseEventStreamAuthorization, AuthRuleV2DraftResponseEventStreamThreeDSAuthentication, AuthRuleV2DraftResponseEventStreamTokenization, AuthRuleV2DraftResponseEventStreamACHCreditReceipt, AuthRuleV2DraftResponseEventStreamACHDebitReceipt:
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2PromoteResponseCurrentVersionParameters) AsUnion() AuthRuleV2PromoteResponseCurrentVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	// This field can have the runtime type of [Conditional3DSActionParametersAction],
	// [ConditionalAuthorizationActionParametersAction],
	// [ConditionalACHActionParametersAction],
	// [ConditionalTokenizationActionParametersAction].
	Action interface{} `json:"action"`
	// This field can have the runtime type of [[]AuthRuleCondition],
	// [[]Conditional3DsActionParametersCondition],
	// [[]ConditionalAuthorizationActionParametersCondition],
	// [[]ConditionalACHActionParametersCondition],
	// [[]ConditionalTokenizationActionParametersCondition].
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period"`
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
// [Conditional3DSActionParameters], [ConditionalAuthorizationActionParameters],
// [ConditionalACHActionParameters], [ConditionalTokenizationActionParameters].
func (r AuthRuleV2PromoteResponseDraftVersionParameters) AsUnion() AuthRuleV2PromoteResponseDraftVersionParametersUnion {
	return r.union
}

// Parameters for the Auth Rule
//
// Union satisfied by [ConditionalBlockParameters], [VelocityLimitParams],
// [MerchantLockParameters], [Conditional3DSActionParameters],
// [ConditionalAuthorizationActionParameters], [ConditionalACHActionParameters] or
// [ConditionalTokenizationActionParameters].
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
	)
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
	AuthRuleV2PromoteResponseEventStreamTokenization          AuthRuleV2PromoteResponseEventStream = "TOKENIZATION"
	AuthRuleV2PromoteResponseEventStreamACHCreditReceipt      AuthRuleV2PromoteResponseEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2PromoteResponseEventStreamACHDebitReceipt       AuthRuleV2PromoteResponseEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2PromoteResponseEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2PromoteResponseEventStreamAuthorization, AuthRuleV2PromoteResponseEventStreamThreeDSAuthentication, AuthRuleV2PromoteResponseEventStreamTokenization, AuthRuleV2PromoteResponseEventStreamACHCreditReceipt, AuthRuleV2PromoteResponseEventStreamACHDebitReceipt:
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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
	// Velocity over the current day since 00:00 / 12 AM in Eastern Time
	Period VelocityLimitPeriod `json:"period,required"`
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
	Parameters param.Field[interface{}] `json:"parameters,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type                  param.Field[AuthRuleV2NewParamsBodyType] `json:"type,required"`
	AccountTokens         param.Field[interface{}]                 `json:"account_tokens"`
	BusinessAccountTokens param.Field[interface{}]                 `json:"business_account_tokens"`
	CardTokens            param.Field[interface{}]                 `json:"card_tokens"`
	// The event stream during which the rule will be evaluated.
	EventStream        param.Field[AuthRuleV2NewParamsBodyEventStream] `json:"event_stream"`
	ExcludedCardTokens param.Field[interface{}]                        `json:"excluded_card_tokens"`
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
	Parameters param.Field[AuthRuleV2NewParamsBodyAccountLevelRuleParametersUnion] `json:"parameters,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type param.Field[AuthRuleV2NewParamsBodyAccountLevelRuleType] `json:"type,required"`
	// Account tokens to which the Auth Rule applies.
	AccountTokens param.Field[[]string] `json:"account_tokens" format:"uuid"`
	// Business Account tokens to which the Auth Rule applies.
	BusinessAccountTokens param.Field[[]string] `json:"business_account_tokens" format:"uuid"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[AuthRuleV2NewParamsBodyAccountLevelRuleEventStream] `json:"event_stream"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
}

func (r AuthRuleV2NewParamsBodyAccountLevelRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyAccountLevelRule) implementsAuthRuleV2NewParamsBodyUnion() {}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyAccountLevelRuleParameters struct {
	Action     param.Field[interface{}] `json:"action"`
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
// [ConditionalTokenizationActionParameters],
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
type AuthRuleV2NewParamsBodyAccountLevelRuleType string

const (
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeConditionalBlock  AuthRuleV2NewParamsBodyAccountLevelRuleType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeVelocityLimit     AuthRuleV2NewParamsBodyAccountLevelRuleType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeMerchantLock      AuthRuleV2NewParamsBodyAccountLevelRuleType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyAccountLevelRuleTypeConditionalAction AuthRuleV2NewParamsBodyAccountLevelRuleType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2NewParamsBodyAccountLevelRuleType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyAccountLevelRuleTypeConditionalBlock, AuthRuleV2NewParamsBodyAccountLevelRuleTypeVelocityLimit, AuthRuleV2NewParamsBodyAccountLevelRuleTypeMerchantLock, AuthRuleV2NewParamsBodyAccountLevelRuleTypeConditionalAction:
		return true
	}
	return false
}

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewParamsBodyAccountLevelRuleEventStream string

const (
	AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamAuthorization         AuthRuleV2NewParamsBodyAccountLevelRuleEventStream = "AUTHORIZATION"
	AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamThreeDSAuthentication AuthRuleV2NewParamsBodyAccountLevelRuleEventStream = "THREE_DS_AUTHENTICATION"
	AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamTokenization          AuthRuleV2NewParamsBodyAccountLevelRuleEventStream = "TOKENIZATION"
	AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamACHCreditReceipt      AuthRuleV2NewParamsBodyAccountLevelRuleEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamACHDebitReceipt       AuthRuleV2NewParamsBodyAccountLevelRuleEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2NewParamsBodyAccountLevelRuleEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamAuthorization, AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamThreeDSAuthentication, AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamTokenization, AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamACHCreditReceipt, AuthRuleV2NewParamsBodyAccountLevelRuleEventStreamACHDebitReceipt:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyCardLevelRule struct {
	// Card tokens to which the Auth Rule applies.
	CardTokens param.Field[[]string] `json:"card_tokens,required" format:"uuid"`
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyCardLevelRuleParametersUnion] `json:"parameters,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type param.Field[AuthRuleV2NewParamsBodyCardLevelRuleType] `json:"type,required"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[AuthRuleV2NewParamsBodyCardLevelRuleEventStream] `json:"event_stream"`
	// Auth Rule Name
	Name param.Field[string] `json:"name"`
}

func (r AuthRuleV2NewParamsBodyCardLevelRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthRuleV2NewParamsBodyCardLevelRule) implementsAuthRuleV2NewParamsBodyUnion() {}

// Parameters for the Auth Rule
type AuthRuleV2NewParamsBodyCardLevelRuleParameters struct {
	Action     param.Field[interface{}] `json:"action"`
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
// [ConditionalTokenizationActionParameters],
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
type AuthRuleV2NewParamsBodyCardLevelRuleType string

const (
	AuthRuleV2NewParamsBodyCardLevelRuleTypeConditionalBlock  AuthRuleV2NewParamsBodyCardLevelRuleType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyCardLevelRuleTypeVelocityLimit     AuthRuleV2NewParamsBodyCardLevelRuleType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyCardLevelRuleTypeMerchantLock      AuthRuleV2NewParamsBodyCardLevelRuleType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyCardLevelRuleTypeConditionalAction AuthRuleV2NewParamsBodyCardLevelRuleType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2NewParamsBodyCardLevelRuleType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCardLevelRuleTypeConditionalBlock, AuthRuleV2NewParamsBodyCardLevelRuleTypeVelocityLimit, AuthRuleV2NewParamsBodyCardLevelRuleTypeMerchantLock, AuthRuleV2NewParamsBodyCardLevelRuleTypeConditionalAction:
		return true
	}
	return false
}

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewParamsBodyCardLevelRuleEventStream string

const (
	AuthRuleV2NewParamsBodyCardLevelRuleEventStreamAuthorization         AuthRuleV2NewParamsBodyCardLevelRuleEventStream = "AUTHORIZATION"
	AuthRuleV2NewParamsBodyCardLevelRuleEventStreamThreeDSAuthentication AuthRuleV2NewParamsBodyCardLevelRuleEventStream = "THREE_DS_AUTHENTICATION"
	AuthRuleV2NewParamsBodyCardLevelRuleEventStreamTokenization          AuthRuleV2NewParamsBodyCardLevelRuleEventStream = "TOKENIZATION"
	AuthRuleV2NewParamsBodyCardLevelRuleEventStreamACHCreditReceipt      AuthRuleV2NewParamsBodyCardLevelRuleEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2NewParamsBodyCardLevelRuleEventStreamACHDebitReceipt       AuthRuleV2NewParamsBodyCardLevelRuleEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2NewParamsBodyCardLevelRuleEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyCardLevelRuleEventStreamAuthorization, AuthRuleV2NewParamsBodyCardLevelRuleEventStreamThreeDSAuthentication, AuthRuleV2NewParamsBodyCardLevelRuleEventStreamTokenization, AuthRuleV2NewParamsBodyCardLevelRuleEventStreamACHCreditReceipt, AuthRuleV2NewParamsBodyCardLevelRuleEventStreamACHDebitReceipt:
		return true
	}
	return false
}

type AuthRuleV2NewParamsBodyProgramLevelRule struct {
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2NewParamsBodyProgramLevelRuleParametersUnion] `json:"parameters,required"`
	// Whether the Auth Rule applies to all authorizations on the card program.
	ProgramLevel param.Field[bool] `json:"program_level,required"`
	// The type of Auth Rule. For certain rule types, this determines the event stream
	// during which it will be evaluated. For rules that can be applied to one of
	// several event streams, the effective one is defined by the separate
	// `event_stream` field.
	//
	//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
	//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
	//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
	//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
	//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
	Type param.Field[AuthRuleV2NewParamsBodyProgramLevelRuleType] `json:"type,required"`
	// The event stream during which the rule will be evaluated.
	EventStream param.Field[AuthRuleV2NewParamsBodyProgramLevelRuleEventStream] `json:"event_stream"`
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
	Action     param.Field[interface{}] `json:"action"`
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
// [ConditionalTokenizationActionParameters],
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
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
type AuthRuleV2NewParamsBodyProgramLevelRuleType string

const (
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeConditionalBlock  AuthRuleV2NewParamsBodyProgramLevelRuleType = "CONDITIONAL_BLOCK"
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeVelocityLimit     AuthRuleV2NewParamsBodyProgramLevelRuleType = "VELOCITY_LIMIT"
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeMerchantLock      AuthRuleV2NewParamsBodyProgramLevelRuleType = "MERCHANT_LOCK"
	AuthRuleV2NewParamsBodyProgramLevelRuleTypeConditionalAction AuthRuleV2NewParamsBodyProgramLevelRuleType = "CONDITIONAL_ACTION"
)

func (r AuthRuleV2NewParamsBodyProgramLevelRuleType) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyProgramLevelRuleTypeConditionalBlock, AuthRuleV2NewParamsBodyProgramLevelRuleTypeVelocityLimit, AuthRuleV2NewParamsBodyProgramLevelRuleTypeMerchantLock, AuthRuleV2NewParamsBodyProgramLevelRuleTypeConditionalAction:
		return true
	}
	return false
}

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewParamsBodyProgramLevelRuleEventStream string

const (
	AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamAuthorization         AuthRuleV2NewParamsBodyProgramLevelRuleEventStream = "AUTHORIZATION"
	AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamThreeDSAuthentication AuthRuleV2NewParamsBodyProgramLevelRuleEventStream = "THREE_DS_AUTHENTICATION"
	AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamTokenization          AuthRuleV2NewParamsBodyProgramLevelRuleEventStream = "TOKENIZATION"
	AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamACHCreditReceipt      AuthRuleV2NewParamsBodyProgramLevelRuleEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamACHDebitReceipt       AuthRuleV2NewParamsBodyProgramLevelRuleEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2NewParamsBodyProgramLevelRuleEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamAuthorization, AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamThreeDSAuthentication, AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamTokenization, AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamACHCreditReceipt, AuthRuleV2NewParamsBodyProgramLevelRuleEventStreamACHDebitReceipt:
		return true
	}
	return false
}

// The type of Auth Rule. For certain rule types, this determines the event stream
// during which it will be evaluated. For rules that can be applied to one of
// several event streams, the effective one is defined by the separate
// `event_stream` field.
//
//   - `CONDITIONAL_BLOCK`: AUTHORIZATION event stream.
//   - `VELOCITY_LIMIT`: AUTHORIZATION event stream.
//   - `MERCHANT_LOCK`: AUTHORIZATION event stream.
//   - `CONDITIONAL_ACTION`: AUTHORIZATION, THREE_DS_AUTHENTICATION, TOKENIZATION,
//     ACH_CREDIT_RECEIPT, or ACH_DEBIT_RECEIPT event stream.
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

// The event stream during which the rule will be evaluated.
type AuthRuleV2NewParamsBodyEventStream string

const (
	AuthRuleV2NewParamsBodyEventStreamAuthorization         AuthRuleV2NewParamsBodyEventStream = "AUTHORIZATION"
	AuthRuleV2NewParamsBodyEventStreamThreeDSAuthentication AuthRuleV2NewParamsBodyEventStream = "THREE_DS_AUTHENTICATION"
	AuthRuleV2NewParamsBodyEventStreamTokenization          AuthRuleV2NewParamsBodyEventStream = "TOKENIZATION"
	AuthRuleV2NewParamsBodyEventStreamACHCreditReceipt      AuthRuleV2NewParamsBodyEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2NewParamsBodyEventStreamACHDebitReceipt       AuthRuleV2NewParamsBodyEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2NewParamsBodyEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2NewParamsBodyEventStreamAuthorization, AuthRuleV2NewParamsBodyEventStreamThreeDSAuthentication, AuthRuleV2NewParamsBodyEventStreamTokenization, AuthRuleV2NewParamsBodyEventStreamACHCreditReceipt, AuthRuleV2NewParamsBodyEventStreamACHDebitReceipt:
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
	AccountTokens         param.Field[interface{}] `json:"account_tokens"`
	BusinessAccountTokens param.Field[interface{}] `json:"business_account_tokens"`
	CardTokens            param.Field[interface{}] `json:"card_tokens"`
	ExcludedCardTokens    param.Field[interface{}] `json:"excluded_card_tokens"`
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
	// Only return Auth rules that are executed during the provided event stream.
	EventStream param.Field[AuthRuleV2ListParamsEventStream] `query:"event_stream"`
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

// Only return Auth rules that are executed during the provided event stream.
type AuthRuleV2ListParamsEventStream string

const (
	AuthRuleV2ListParamsEventStreamAuthorization         AuthRuleV2ListParamsEventStream = "AUTHORIZATION"
	AuthRuleV2ListParamsEventStreamThreeDSAuthentication AuthRuleV2ListParamsEventStream = "THREE_DS_AUTHENTICATION"
	AuthRuleV2ListParamsEventStreamTokenization          AuthRuleV2ListParamsEventStream = "TOKENIZATION"
	AuthRuleV2ListParamsEventStreamACHCreditReceipt      AuthRuleV2ListParamsEventStream = "ACH_CREDIT_RECEIPT"
	AuthRuleV2ListParamsEventStreamACHDebitReceipt       AuthRuleV2ListParamsEventStream = "ACH_DEBIT_RECEIPT"
)

func (r AuthRuleV2ListParamsEventStream) IsKnown() bool {
	switch r {
	case AuthRuleV2ListParamsEventStreamAuthorization, AuthRuleV2ListParamsEventStreamThreeDSAuthentication, AuthRuleV2ListParamsEventStreamTokenization, AuthRuleV2ListParamsEventStreamACHCreditReceipt, AuthRuleV2ListParamsEventStreamACHDebitReceipt:
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

type AuthRuleV2DraftParams struct {
	// Parameters for the Auth Rule
	Parameters param.Field[AuthRuleV2DraftParamsParametersUnion] `json:"parameters"`
}

func (r AuthRuleV2DraftParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Parameters for the Auth Rule
type AuthRuleV2DraftParamsParameters struct {
	Action     param.Field[interface{}] `json:"action"`
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
// [ConditionalTokenizationActionParameters], [AuthRuleV2DraftParamsParameters].
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
