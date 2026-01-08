// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// FinancialAccountLoanTapeService contains methods and other services that help
// with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountLoanTapeService] method instead.
type FinancialAccountLoanTapeService struct {
	Options []option.RequestOption
}

// NewFinancialAccountLoanTapeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFinancialAccountLoanTapeService(opts ...option.RequestOption) (r *FinancialAccountLoanTapeService) {
	r = &FinancialAccountLoanTapeService{}
	r.Options = opts
	return
}

// Get a specific loan tape for a given financial account.
func (r *FinancialAccountLoanTapeService) Get(ctx context.Context, financialAccountToken string, loanTapeToken string, opts ...option.RequestOption) (res *LoanTape, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	if loanTapeToken == "" {
		err = errors.New("missing required loan_tape_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/loan_tapes/%s", financialAccountToken, loanTapeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the loan tapes for a given financial account.
func (r *FinancialAccountLoanTapeService) List(ctx context.Context, financialAccountToken string, query FinancialAccountLoanTapeListParams, opts ...option.RequestOption) (res *pagination.CursorPage[LoanTape], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/loan_tapes", financialAccountToken)
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

// List the loan tapes for a given financial account.
func (r *FinancialAccountLoanTapeService) ListAutoPaging(ctx context.Context, financialAccountToken string, query FinancialAccountLoanTapeListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[LoanTape] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, financialAccountToken, query, opts...))
}

type CategoryBalances struct {
	Fees      int64                `json:"fees,required"`
	Interest  int64                `json:"interest,required"`
	Principal int64                `json:"principal,required"`
	JSON      categoryBalancesJSON `json:"-"`
}

// categoryBalancesJSON contains the JSON metadata for the struct
// [CategoryBalances]
type categoryBalancesJSON struct {
	Fees        apijson.Field
	Interest    apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CategoryBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r categoryBalancesJSON) RawJSON() string {
	return r.raw
}

type LoanTape struct {
	// Globally unique identifier for a loan tape
	Token           string                  `json:"token,required"`
	AccountStanding LoanTapeAccountStanding `json:"account_standing,required"`
	// Amount of credit available to spend in cents
	AvailableCredit int64            `json:"available_credit,required"`
	Balances        LoanTapeBalances `json:"balances,required"`
	// Timestamp of when the loan tape was created
	Created time.Time `json:"created,required" format:"date-time"`
	// For prepay accounts, this is the minimum prepay balance that must be maintained.
	// For charge card accounts, this is the maximum credit balance extended by a
	// lender
	CreditLimit int64 `json:"credit_limit,required"`
	// Globally unique identifier for a credit product
	CreditProductToken string `json:"credit_product_token,required"`
	// Date of transactions that this loan tape covers
	Date      time.Time       `json:"date,required" format:"date"`
	DayTotals StatementTotals `json:"day_totals,required"`
	// Balance at the end of the day
	EndingBalance int64 `json:"ending_balance,required"`
	// Excess credits in the form of provisional credits, payments, or purchase
	// refunds. If positive, the account is in net credit state with no outstanding
	// balances. An overpayment could land an account in this state
	ExcessCredits int64 `json:"excess_credits,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken    string                           `json:"financial_account_token,required" format:"uuid"`
	InterestDetails          LoanTapeInterestDetails          `json:"interest_details,required,nullable"`
	MinimumPaymentBalance    LoanTapeMinimumPaymentBalance    `json:"minimum_payment_balance,required"`
	PaymentAllocation        LoanTapePaymentAllocation        `json:"payment_allocation,required"`
	PeriodTotals             StatementTotals                  `json:"period_totals,required"`
	PreviousStatementBalance LoanTapePreviousStatementBalance `json:"previous_statement_balance,required"`
	// Balance at the start of the day
	StartingBalance int64 `json:"starting_balance,required"`
	// Timestamp of when the loan tape was updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Version number of the loan tape. This starts at 1
	Version   int64           `json:"version,required"`
	YtdTotals StatementTotals `json:"ytd_totals,required"`
	// Interest tier to which this account belongs to
	Tier string       `json:"tier,nullable"`
	JSON loanTapeJSON `json:"-"`
}

// loanTapeJSON contains the JSON metadata for the struct [LoanTape]
type loanTapeJSON struct {
	Token                    apijson.Field
	AccountStanding          apijson.Field
	AvailableCredit          apijson.Field
	Balances                 apijson.Field
	Created                  apijson.Field
	CreditLimit              apijson.Field
	CreditProductToken       apijson.Field
	Date                     apijson.Field
	DayTotals                apijson.Field
	EndingBalance            apijson.Field
	ExcessCredits            apijson.Field
	FinancialAccountToken    apijson.Field
	InterestDetails          apijson.Field
	MinimumPaymentBalance    apijson.Field
	PaymentAllocation        apijson.Field
	PeriodTotals             apijson.Field
	PreviousStatementBalance apijson.Field
	StartingBalance          apijson.Field
	Updated                  apijson.Field
	Version                  apijson.Field
	YtdTotals                apijson.Field
	Tier                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LoanTape) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeJSON) RawJSON() string {
	return r.raw
}

type LoanTapeAccountStanding struct {
	// Number of consecutive full payments made
	ConsecutiveFullPaymentsMade int64 `json:"consecutive_full_payments_made,required"`
	// Number of consecutive minimum payments made
	ConsecutiveMinimumPaymentsMade int64 `json:"consecutive_minimum_payments_made,required"`
	// Number of consecutive minimum payments missed
	ConsecutiveMinimumPaymentsMissed int64 `json:"consecutive_minimum_payments_missed,required"`
	// Number of days past due
	DaysPastDue int64 `json:"days_past_due,required"`
	// Information about the financial account state
	FinancialAccountState LoanTapeAccountStandingFinancialAccountState `json:"financial_account_state,required"`
	// Whether the account currently has grace or not
	HasGrace bool `json:"has_grace,required"`
	// Current overall period number
	PeriodNumber int64                              `json:"period_number,required"`
	PeriodState  LoanTapeAccountStandingPeriodState `json:"period_state,required"`
	JSON         loanTapeAccountStandingJSON        `json:"-"`
}

// loanTapeAccountStandingJSON contains the JSON metadata for the struct
// [LoanTapeAccountStanding]
type loanTapeAccountStandingJSON struct {
	ConsecutiveFullPaymentsMade      apijson.Field
	ConsecutiveMinimumPaymentsMade   apijson.Field
	ConsecutiveMinimumPaymentsMissed apijson.Field
	DaysPastDue                      apijson.Field
	FinancialAccountState            apijson.Field
	HasGrace                         apijson.Field
	PeriodNumber                     apijson.Field
	PeriodState                      apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *LoanTapeAccountStanding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeAccountStandingJSON) RawJSON() string {
	return r.raw
}

// Information about the financial account state
type LoanTapeAccountStandingFinancialAccountState struct {
	// Status of the financial account
	Status LoanTapeAccountStandingFinancialAccountStateStatus `json:"status,required"`
	// Substatus for the financial account
	Substatus LoanTapeAccountStandingFinancialAccountStateSubstatus `json:"substatus,nullable"`
	JSON      loanTapeAccountStandingFinancialAccountStateJSON      `json:"-"`
}

// loanTapeAccountStandingFinancialAccountStateJSON contains the JSON metadata for
// the struct [LoanTapeAccountStandingFinancialAccountState]
type loanTapeAccountStandingFinancialAccountStateJSON struct {
	Status      apijson.Field
	Substatus   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapeAccountStandingFinancialAccountState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeAccountStandingFinancialAccountStateJSON) RawJSON() string {
	return r.raw
}

// Status of the financial account
type LoanTapeAccountStandingFinancialAccountStateStatus string

const (
	LoanTapeAccountStandingFinancialAccountStateStatusOpen      LoanTapeAccountStandingFinancialAccountStateStatus = "OPEN"
	LoanTapeAccountStandingFinancialAccountStateStatusClosed    LoanTapeAccountStandingFinancialAccountStateStatus = "CLOSED"
	LoanTapeAccountStandingFinancialAccountStateStatusSuspended LoanTapeAccountStandingFinancialAccountStateStatus = "SUSPENDED"
	LoanTapeAccountStandingFinancialAccountStateStatusPending   LoanTapeAccountStandingFinancialAccountStateStatus = "PENDING"
)

func (r LoanTapeAccountStandingFinancialAccountStateStatus) IsKnown() bool {
	switch r {
	case LoanTapeAccountStandingFinancialAccountStateStatusOpen, LoanTapeAccountStandingFinancialAccountStateStatusClosed, LoanTapeAccountStandingFinancialAccountStateStatusSuspended, LoanTapeAccountStandingFinancialAccountStateStatusPending:
		return true
	}
	return false
}

// Substatus for the financial account
type LoanTapeAccountStandingFinancialAccountStateSubstatus string

const (
	LoanTapeAccountStandingFinancialAccountStateSubstatusChargedOffDelinquent LoanTapeAccountStandingFinancialAccountStateSubstatus = "CHARGED_OFF_DELINQUENT"
	LoanTapeAccountStandingFinancialAccountStateSubstatusChargedOffFraud      LoanTapeAccountStandingFinancialAccountStateSubstatus = "CHARGED_OFF_FRAUD"
	LoanTapeAccountStandingFinancialAccountStateSubstatusEndUserRequest       LoanTapeAccountStandingFinancialAccountStateSubstatus = "END_USER_REQUEST"
	LoanTapeAccountStandingFinancialAccountStateSubstatusBankRequest          LoanTapeAccountStandingFinancialAccountStateSubstatus = "BANK_REQUEST"
	LoanTapeAccountStandingFinancialAccountStateSubstatusDelinquent           LoanTapeAccountStandingFinancialAccountStateSubstatus = "DELINQUENT"
)

func (r LoanTapeAccountStandingFinancialAccountStateSubstatus) IsKnown() bool {
	switch r {
	case LoanTapeAccountStandingFinancialAccountStateSubstatusChargedOffDelinquent, LoanTapeAccountStandingFinancialAccountStateSubstatusChargedOffFraud, LoanTapeAccountStandingFinancialAccountStateSubstatusEndUserRequest, LoanTapeAccountStandingFinancialAccountStateSubstatusBankRequest, LoanTapeAccountStandingFinancialAccountStateSubstatusDelinquent:
		return true
	}
	return false
}

type LoanTapeAccountStandingPeriodState string

const (
	LoanTapeAccountStandingPeriodStateStandard LoanTapeAccountStandingPeriodState = "STANDARD"
	LoanTapeAccountStandingPeriodStatePromo    LoanTapeAccountStandingPeriodState = "PROMO"
	LoanTapeAccountStandingPeriodStatePenalty  LoanTapeAccountStandingPeriodState = "PENALTY"
)

func (r LoanTapeAccountStandingPeriodState) IsKnown() bool {
	switch r {
	case LoanTapeAccountStandingPeriodStateStandard, LoanTapeAccountStandingPeriodStatePromo, LoanTapeAccountStandingPeriodStatePenalty:
		return true
	}
	return false
}

type LoanTapeBalances struct {
	// Amount due for the prior billing cycle. Any amounts not fully paid off on this
	// due date will be considered past due the next day
	Due CategoryBalances `json:"due,required"`
	// Amount due for the current billing cycle. Any amounts not paid off by early
	// payments or credits will be considered due at the end of the current billing
	// period
	NextStatementDue CategoryBalances `json:"next_statement_due,required"`
	// Amount not paid off on previous due dates
	PastDue CategoryBalances `json:"past_due,required"`
	// Amount due for the past billing cycles.
	PastStatementsDue CategoryBalances     `json:"past_statements_due,required"`
	JSON              loanTapeBalancesJSON `json:"-"`
}

// loanTapeBalancesJSON contains the JSON metadata for the struct
// [LoanTapeBalances]
type loanTapeBalancesJSON struct {
	Due               apijson.Field
	NextStatementDue  apijson.Field
	PastDue           apijson.Field
	PastStatementsDue apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LoanTapeBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeBalancesJSON) RawJSON() string {
	return r.raw
}

type LoanTapeInterestDetails struct {
	ActualInterestCharged     int64                                            `json:"actual_interest_charged,required,nullable"`
	DailyBalanceAmounts       CategoryDetails                                  `json:"daily_balance_amounts,required"`
	EffectiveApr              CategoryDetails                                  `json:"effective_apr,required"`
	InterestCalculationMethod LoanTapeInterestDetailsInterestCalculationMethod `json:"interest_calculation_method,required"`
	InterestForPeriod         CategoryDetails                                  `json:"interest_for_period,required"`
	PrimeRate                 string                                           `json:"prime_rate,required,nullable"`
	MinimumInterestCharged    int64                                            `json:"minimum_interest_charged,nullable"`
	JSON                      loanTapeInterestDetailsJSON                      `json:"-"`
}

// loanTapeInterestDetailsJSON contains the JSON metadata for the struct
// [LoanTapeInterestDetails]
type loanTapeInterestDetailsJSON struct {
	ActualInterestCharged     apijson.Field
	DailyBalanceAmounts       apijson.Field
	EffectiveApr              apijson.Field
	InterestCalculationMethod apijson.Field
	InterestForPeriod         apijson.Field
	PrimeRate                 apijson.Field
	MinimumInterestCharged    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoanTapeInterestDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeInterestDetailsJSON) RawJSON() string {
	return r.raw
}

type LoanTapeInterestDetailsInterestCalculationMethod string

const (
	LoanTapeInterestDetailsInterestCalculationMethodDaily        LoanTapeInterestDetailsInterestCalculationMethod = "DAILY"
	LoanTapeInterestDetailsInterestCalculationMethodAverageDaily LoanTapeInterestDetailsInterestCalculationMethod = "AVERAGE_DAILY"
)

func (r LoanTapeInterestDetailsInterestCalculationMethod) IsKnown() bool {
	switch r {
	case LoanTapeInterestDetailsInterestCalculationMethodDaily, LoanTapeInterestDetailsInterestCalculationMethodAverageDaily:
		return true
	}
	return false
}

type LoanTapeMinimumPaymentBalance struct {
	Amount    int64                             `json:"amount,required"`
	Remaining int64                             `json:"remaining,required"`
	JSON      loanTapeMinimumPaymentBalanceJSON `json:"-"`
}

// loanTapeMinimumPaymentBalanceJSON contains the JSON metadata for the struct
// [LoanTapeMinimumPaymentBalance]
type loanTapeMinimumPaymentBalanceJSON struct {
	Amount      apijson.Field
	Remaining   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapeMinimumPaymentBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeMinimumPaymentBalanceJSON) RawJSON() string {
	return r.raw
}

type LoanTapePaymentAllocation struct {
	FeeDetails CategoryDetails `json:"fee_details,required,nullable"`
	// Amount allocated to fees in cents
	Fees int64 `json:"fees,required"`
	// Amount allocated to interest in cents
	Interest        int64           `json:"interest,required"`
	InterestDetails CategoryDetails `json:"interest_details,required,nullable"`
	// Amount allocated to principal in cents
	Principal        int64                         `json:"principal,required"`
	PrincipalDetails CategoryDetails               `json:"principal_details,required,nullable"`
	JSON             loanTapePaymentAllocationJSON `json:"-"`
}

// loanTapePaymentAllocationJSON contains the JSON metadata for the struct
// [LoanTapePaymentAllocation]
type loanTapePaymentAllocationJSON struct {
	FeeDetails       apijson.Field
	Fees             apijson.Field
	Interest         apijson.Field
	InterestDetails  apijson.Field
	Principal        apijson.Field
	PrincipalDetails apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoanTapePaymentAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapePaymentAllocationJSON) RawJSON() string {
	return r.raw
}

type LoanTapePreviousStatementBalance struct {
	Amount    int64                                `json:"amount,required"`
	Remaining int64                                `json:"remaining,required"`
	JSON      loanTapePreviousStatementBalanceJSON `json:"-"`
}

// loanTapePreviousStatementBalanceJSON contains the JSON metadata for the struct
// [LoanTapePreviousStatementBalance]
type loanTapePreviousStatementBalanceJSON struct {
	Amount      apijson.Field
	Remaining   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapePreviousStatementBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapePreviousStatementBalanceJSON) RawJSON() string {
	return r.raw
}

type FinancialAccountLoanTapeListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included.
	Begin param.Field[time.Time] `query:"begin" format:"date"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included.
	End param.Field[time.Time] `query:"end" format:"date"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [FinancialAccountLoanTapeListParams]'s query parameters as
// `url.Values`.
func (r FinancialAccountLoanTapeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
