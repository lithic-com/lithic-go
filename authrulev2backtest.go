// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// AuthRuleV2BacktestService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthRuleV2BacktestService] method instead.
type AuthRuleV2BacktestService struct {
	Options []option.RequestOption
}

// NewAuthRuleV2BacktestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAuthRuleV2BacktestService(opts ...option.RequestOption) (r *AuthRuleV2BacktestService) {
	r = &AuthRuleV2BacktestService{}
	r.Options = opts
	return
}

// Initiates a request to asynchronously generate a backtest for an Auth rule.
// During backtesting, both the active version (if one exists) and the draft
// version of the Auth Rule are evaluated by replaying historical transaction data
// against the rule's conditions. This process allows customers to simulate and
// understand the effects of proposed rule changes before deployment. The generated
// backtest report provides detailed results showing whether the draft version of
// the Auth Rule would have approved or declined historical transactions which were
// processed during the backtest period. These reports help evaluate how changes to
// rule configurations might affect overall transaction approval rates.
//
// The generated backtest report will be delivered asynchronously through a webhook
// with `event_type` = `auth_rules.backtest_report.created`. See the docs on
// setting up [webhook subscriptions](https://docs.lithic.com/docs/events-api). It
// is also possible to request backtest reports on-demand through the
// `/v2/auth_rules/{auth_rule_token}/backtests/{auth_rule_backtest_token}`
// endpoint.
//
// Lithic currently supports backtesting for `CONDITIONAL_BLOCK` rules. Backtesting
// for `VELOCITY_LIMIT` rules is generally not supported. In specific cases (i.e.
// where Lithic has pre-calculated the requested velocity metrics for historical
// transactions), a backtest may be feasible. However, such cases are uncommon and
// customers should not anticipate support for velocity backtests under most
// configurations. If a historical transaction does not feature the required inputs
// to evaluate the rule, then it will not be included in the final backtest report.
func (r *AuthRuleV2BacktestService) New(ctx context.Context, authRuleToken string, body AuthRuleV2BacktestNewParams, opts ...option.RequestOption) (res *AuthRuleV2BacktestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/backtests", authRuleToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the backtest results of an Auth rule (if available).
//
// Backtesting is an asynchronous process that requires time to complete. If a
// customer retrieves the backtest results using this endpoint before the report is
// fully generated, the response will return null for `results.current_version` and
// `results.draft_version`. Customers are advised to wait for the backtest creation
// process to complete (as indicated by the webhook event
// auth_rules.backtest_report.created) before retrieving results from this
// endpoint.
//
// Backtesting is an asynchronous process, while the backtest is being processed,
// results will not be available which will cause `results.current_version` and
// `results.draft_version` objects to contain `null`. The entries in `results` will
// also always represent the configuration of the rule at the time requests are
// made to this endpoint. For example, the results for `current_version` in the
// served backtest report will be consistent with which version of the rule is
// currently activated in the respective event stream, regardless of which version
// of the rule was active in the event stream at the time a backtest is requested.
func (r *AuthRuleV2BacktestService) Get(ctx context.Context, authRuleToken string, authRuleBacktestToken string, opts ...option.RequestOption) (res *BacktestResults, err error) {
	opts = append(r.Options[:], opts...)
	if authRuleToken == "" {
		err = errors.New("missing required auth_rule_token parameter")
		return
	}
	if authRuleBacktestToken == "" {
		err = errors.New("missing required auth_rule_backtest_token parameter")
		return
	}
	path := fmt.Sprintf("v2/auth_rules/%s/backtests/%s", authRuleToken, authRuleBacktestToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BacktestResults struct {
	// Auth Rule Backtest Token
	BacktestToken        string                              `json:"backtest_token,required" format:"uuid"`
	Results              BacktestResultsResults              `json:"results,required"`
	SimulationParameters BacktestResultsSimulationParameters `json:"simulation_parameters,required"`
	JSON                 backtestResultsJSON                 `json:"-"`
}

// backtestResultsJSON contains the JSON metadata for the struct [BacktestResults]
type backtestResultsJSON struct {
	BacktestToken        apijson.Field
	Results              apijson.Field
	SimulationParameters apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BacktestResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestResultsJSON) RawJSON() string {
	return r.raw
}

type BacktestResultsResults struct {
	CurrentVersion BacktestResultsResultsCurrentVersion `json:"current_version,nullable"`
	DraftVersion   BacktestResultsResultsDraftVersion   `json:"draft_version,nullable"`
	JSON           backtestResultsResultsJSON           `json:"-"`
}

// backtestResultsResultsJSON contains the JSON metadata for the struct
// [BacktestResultsResults]
type backtestResultsResultsJSON struct {
	CurrentVersion apijson.Field
	DraftVersion   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BacktestResultsResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestResultsResultsJSON) RawJSON() string {
	return r.raw
}

type BacktestResultsResultsCurrentVersion struct {
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
	Examples []BacktestResultsResultsCurrentVersionExample `json:"examples"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                    `json:"version"`
	JSON    backtestResultsResultsCurrentVersionJSON `json:"-"`
}

// backtestResultsResultsCurrentVersionJSON contains the JSON metadata for the
// struct [BacktestResultsResultsCurrentVersion]
type backtestResultsResultsCurrentVersionJSON struct {
	Approved    apijson.Field
	Challenged  apijson.Field
	Declined    apijson.Field
	Examples    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BacktestResultsResultsCurrentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestResultsResultsCurrentVersionJSON) RawJSON() string {
	return r.raw
}

type BacktestResultsResultsCurrentVersionExample struct {
	// Whether the rule would have approved the request.
	Approved bool `json:"approved"`
	// The decision made by the rule for this event.
	Decision BacktestResultsResultsCurrentVersionExamplesDecision `json:"decision"`
	// The event token.
	EventToken string `json:"event_token" format:"uuid"`
	// The timestamp of the event.
	Timestamp time.Time                                       `json:"timestamp" format:"date-time"`
	JSON      backtestResultsResultsCurrentVersionExampleJSON `json:"-"`
}

// backtestResultsResultsCurrentVersionExampleJSON contains the JSON metadata for
// the struct [BacktestResultsResultsCurrentVersionExample]
type backtestResultsResultsCurrentVersionExampleJSON struct {
	Approved    apijson.Field
	Decision    apijson.Field
	EventToken  apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BacktestResultsResultsCurrentVersionExample) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestResultsResultsCurrentVersionExampleJSON) RawJSON() string {
	return r.raw
}

// The decision made by the rule for this event.
type BacktestResultsResultsCurrentVersionExamplesDecision string

const (
	BacktestResultsResultsCurrentVersionExamplesDecisionApproved   BacktestResultsResultsCurrentVersionExamplesDecision = "APPROVED"
	BacktestResultsResultsCurrentVersionExamplesDecisionDeclined   BacktestResultsResultsCurrentVersionExamplesDecision = "DECLINED"
	BacktestResultsResultsCurrentVersionExamplesDecisionChallenged BacktestResultsResultsCurrentVersionExamplesDecision = "CHALLENGED"
)

func (r BacktestResultsResultsCurrentVersionExamplesDecision) IsKnown() bool {
	switch r {
	case BacktestResultsResultsCurrentVersionExamplesDecisionApproved, BacktestResultsResultsCurrentVersionExamplesDecisionDeclined, BacktestResultsResultsCurrentVersionExamplesDecisionChallenged:
		return true
	}
	return false
}

type BacktestResultsResultsDraftVersion struct {
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
	Examples []BacktestResultsResultsDraftVersionExample `json:"examples"`
	// The version of the rule, this is incremented whenever the rule's parameters
	// change.
	Version int64                                  `json:"version"`
	JSON    backtestResultsResultsDraftVersionJSON `json:"-"`
}

// backtestResultsResultsDraftVersionJSON contains the JSON metadata for the struct
// [BacktestResultsResultsDraftVersion]
type backtestResultsResultsDraftVersionJSON struct {
	Approved    apijson.Field
	Challenged  apijson.Field
	Declined    apijson.Field
	Examples    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BacktestResultsResultsDraftVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestResultsResultsDraftVersionJSON) RawJSON() string {
	return r.raw
}

type BacktestResultsResultsDraftVersionExample struct {
	// Whether the rule would have approved the request.
	Approved bool `json:"approved"`
	// The decision made by the rule for this event.
	Decision BacktestResultsResultsDraftVersionExamplesDecision `json:"decision"`
	// The event token.
	EventToken string `json:"event_token" format:"uuid"`
	// The timestamp of the event.
	Timestamp time.Time                                     `json:"timestamp" format:"date-time"`
	JSON      backtestResultsResultsDraftVersionExampleJSON `json:"-"`
}

// backtestResultsResultsDraftVersionExampleJSON contains the JSON metadata for the
// struct [BacktestResultsResultsDraftVersionExample]
type backtestResultsResultsDraftVersionExampleJSON struct {
	Approved    apijson.Field
	Decision    apijson.Field
	EventToken  apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BacktestResultsResultsDraftVersionExample) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestResultsResultsDraftVersionExampleJSON) RawJSON() string {
	return r.raw
}

// The decision made by the rule for this event.
type BacktestResultsResultsDraftVersionExamplesDecision string

const (
	BacktestResultsResultsDraftVersionExamplesDecisionApproved   BacktestResultsResultsDraftVersionExamplesDecision = "APPROVED"
	BacktestResultsResultsDraftVersionExamplesDecisionDeclined   BacktestResultsResultsDraftVersionExamplesDecision = "DECLINED"
	BacktestResultsResultsDraftVersionExamplesDecisionChallenged BacktestResultsResultsDraftVersionExamplesDecision = "CHALLENGED"
)

func (r BacktestResultsResultsDraftVersionExamplesDecision) IsKnown() bool {
	switch r {
	case BacktestResultsResultsDraftVersionExamplesDecisionApproved, BacktestResultsResultsDraftVersionExamplesDecisionDeclined, BacktestResultsResultsDraftVersionExamplesDecisionChallenged:
		return true
	}
	return false
}

type BacktestResultsSimulationParameters struct {
	// Auth Rule Token
	AuthRuleToken string `json:"auth_rule_token" format:"uuid"`
	// The end time of the simulation.
	End time.Time `json:"end" format:"date-time"`
	// The start time of the simulation.
	Start time.Time                               `json:"start" format:"date-time"`
	JSON  backtestResultsSimulationParametersJSON `json:"-"`
}

// backtestResultsSimulationParametersJSON contains the JSON metadata for the
// struct [BacktestResultsSimulationParameters]
type backtestResultsSimulationParametersJSON struct {
	AuthRuleToken apijson.Field
	End           apijson.Field
	Start         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BacktestResultsSimulationParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backtestResultsSimulationParametersJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2BacktestNewResponse struct {
	// Auth Rule Backtest Token
	BacktestToken string                            `json:"backtest_token" format:"uuid"`
	JSON          authRuleV2BacktestNewResponseJSON `json:"-"`
}

// authRuleV2BacktestNewResponseJSON contains the JSON metadata for the struct
// [AuthRuleV2BacktestNewResponse]
type authRuleV2BacktestNewResponseJSON struct {
	BacktestToken apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuthRuleV2BacktestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authRuleV2BacktestNewResponseJSON) RawJSON() string {
	return r.raw
}

type AuthRuleV2BacktestNewParams struct {
	// The end time of the backtest.
	End param.Field[time.Time] `json:"end" format:"date-time"`
	// The start time of the backtest.
	Start param.Field[time.Time] `json:"start" format:"date-time"`
}

func (r AuthRuleV2BacktestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
