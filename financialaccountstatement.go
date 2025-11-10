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

// FinancialAccountStatementService contains methods and other services that help
// with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountStatementService] method instead.
type FinancialAccountStatementService struct {
	Options   []option.RequestOption
	LineItems *FinancialAccountStatementLineItemService
}

// NewFinancialAccountStatementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFinancialAccountStatementService(opts ...option.RequestOption) (r *FinancialAccountStatementService) {
	r = &FinancialAccountStatementService{}
	r.Options = opts
	r.LineItems = NewFinancialAccountStatementLineItemService(opts...)
	return
}

// Get a specific statement for a given financial account.
func (r *FinancialAccountStatementService) Get(ctx context.Context, financialAccountToken string, statementToken string, opts ...option.RequestOption) (res *Statement, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	if statementToken == "" {
		err = errors.New("missing required statement_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/statements/%s", financialAccountToken, statementToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the statements for a given financial account.
func (r *FinancialAccountStatementService) List(ctx context.Context, financialAccountToken string, query FinancialAccountStatementListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Statement], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/statements", financialAccountToken)
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

// List the statements for a given financial account.
func (r *FinancialAccountStatementService) ListAutoPaging(ctx context.Context, financialAccountToken string, query FinancialAccountStatementListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Statement] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, financialAccountToken, query, opts...))
}

type Statement struct {
	// Globally unique identifier for a statement
	Token           string                   `json:"token,required"`
	AccountStanding StatementAccountStanding `json:"account_standing,required"`
	AmountDue       StatementAmountDue       `json:"amount_due,required"`
	// Amount of credit available to spend in cents
	AvailableCredit int64 `json:"available_credit,required"`
	// Timestamp of when the statement was created
	Created time.Time `json:"created,required" format:"date-time"`
	// This is the maximum credit balance extended by the lender in cents
	CreditLimit int64 `json:"credit_limit,required"`
	// Globally unique identifier for a credit product
	CreditProductToken string `json:"credit_product_token,required"`
	// Number of days in the billing cycle
	DaysInBillingCycle int64 `json:"days_in_billing_cycle,required"`
	// Balance at the end of the billing period. For charge cards, this should be the
	// same at the statement amount due in cents
	EndingBalance int64 `json:"ending_balance,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Date when the payment is due
	PaymentDueDate time.Time             `json:"payment_due_date,required,nullable" format:"date"`
	PeriodTotals   StatementPeriodTotals `json:"period_totals,required"`
	// Balance at the start of the billing period
	StartingBalance int64 `json:"starting_balance,required"`
	// Date when the billing period ended
	StatementEndDate time.Time `json:"statement_end_date,required" format:"date"`
	// Date when the billing period began
	StatementStartDate time.Time              `json:"statement_start_date,required" format:"date"`
	StatementType      StatementStatementType `json:"statement_type,required"`
	// Timestamp of when the statement was updated
	Updated         time.Time                `json:"updated,required" format:"date-time"`
	YtdTotals       StatementYtdTotals       `json:"ytd_totals,required"`
	InterestDetails StatementInterestDetails `json:"interest_details,nullable"`
	// Date when the next payment is due
	NextPaymentDueDate time.Time `json:"next_payment_due_date" format:"date"`
	// Date when the next billing period will end
	NextStatementEndDate time.Time     `json:"next_statement_end_date" format:"date"`
	JSON                 statementJSON `json:"-"`
}

// statementJSON contains the JSON metadata for the struct [Statement]
type statementJSON struct {
	Token                 apijson.Field
	AccountStanding       apijson.Field
	AmountDue             apijson.Field
	AvailableCredit       apijson.Field
	Created               apijson.Field
	CreditLimit           apijson.Field
	CreditProductToken    apijson.Field
	DaysInBillingCycle    apijson.Field
	EndingBalance         apijson.Field
	FinancialAccountToken apijson.Field
	PaymentDueDate        apijson.Field
	PeriodTotals          apijson.Field
	StartingBalance       apijson.Field
	StatementEndDate      apijson.Field
	StatementStartDate    apijson.Field
	StatementType         apijson.Field
	Updated               apijson.Field
	YtdTotals             apijson.Field
	InterestDetails       apijson.Field
	NextPaymentDueDate    apijson.Field
	NextStatementEndDate  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Statement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementJSON) RawJSON() string {
	return r.raw
}

type StatementAccountStanding struct {
	// Number of consecutive full payments made
	ConsecutiveFullPaymentsMade int64 `json:"consecutive_full_payments_made,required"`
	// Number of consecutive minimum payments made
	ConsecutiveMinimumPaymentsMade int64 `json:"consecutive_minimum_payments_made,required"`
	// Number of consecutive minimum payments missed
	ConsecutiveMinimumPaymentsMissed int64 `json:"consecutive_minimum_payments_missed,required"`
	// Number of days past due
	DaysPastDue int64 `json:"days_past_due,required"`
	// Information about the financial account state
	FinancialAccountState StatementAccountStandingFinancialAccountState `json:"financial_account_state,required"`
	// Whether the account currently has grace or not
	HasGrace bool `json:"has_grace,required"`
	// Current overall period number
	PeriodNumber int64                               `json:"period_number,required"`
	PeriodState  StatementAccountStandingPeriodState `json:"period_state,required"`
	JSON         statementAccountStandingJSON        `json:"-"`
}

// statementAccountStandingJSON contains the JSON metadata for the struct
// [StatementAccountStanding]
type statementAccountStandingJSON struct {
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

func (r *StatementAccountStanding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementAccountStandingJSON) RawJSON() string {
	return r.raw
}

// Information about the financial account state
type StatementAccountStandingFinancialAccountState struct {
	// Status of the financial account
	Status StatementAccountStandingFinancialAccountStateStatus `json:"status,required"`
	// Substatus for the financial account
	Substatus StatementAccountStandingFinancialAccountStateSubstatus `json:"substatus,nullable"`
	JSON      statementAccountStandingFinancialAccountStateJSON      `json:"-"`
}

// statementAccountStandingFinancialAccountStateJSON contains the JSON metadata for
// the struct [StatementAccountStandingFinancialAccountState]
type statementAccountStandingFinancialAccountStateJSON struct {
	Status      apijson.Field
	Substatus   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatementAccountStandingFinancialAccountState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementAccountStandingFinancialAccountStateJSON) RawJSON() string {
	return r.raw
}

// Status of the financial account
type StatementAccountStandingFinancialAccountStateStatus string

const (
	StatementAccountStandingFinancialAccountStateStatusOpen      StatementAccountStandingFinancialAccountStateStatus = "OPEN"
	StatementAccountStandingFinancialAccountStateStatusClosed    StatementAccountStandingFinancialAccountStateStatus = "CLOSED"
	StatementAccountStandingFinancialAccountStateStatusSuspended StatementAccountStandingFinancialAccountStateStatus = "SUSPENDED"
	StatementAccountStandingFinancialAccountStateStatusPending   StatementAccountStandingFinancialAccountStateStatus = "PENDING"
)

func (r StatementAccountStandingFinancialAccountStateStatus) IsKnown() bool {
	switch r {
	case StatementAccountStandingFinancialAccountStateStatusOpen, StatementAccountStandingFinancialAccountStateStatusClosed, StatementAccountStandingFinancialAccountStateStatusSuspended, StatementAccountStandingFinancialAccountStateStatusPending:
		return true
	}
	return false
}

// Substatus for the financial account
type StatementAccountStandingFinancialAccountStateSubstatus string

const (
	StatementAccountStandingFinancialAccountStateSubstatusChargedOffDelinquent StatementAccountStandingFinancialAccountStateSubstatus = "CHARGED_OFF_DELINQUENT"
	StatementAccountStandingFinancialAccountStateSubstatusChargedOffFraud      StatementAccountStandingFinancialAccountStateSubstatus = "CHARGED_OFF_FRAUD"
	StatementAccountStandingFinancialAccountStateSubstatusEndUserRequest       StatementAccountStandingFinancialAccountStateSubstatus = "END_USER_REQUEST"
	StatementAccountStandingFinancialAccountStateSubstatusBankRequest          StatementAccountStandingFinancialAccountStateSubstatus = "BANK_REQUEST"
	StatementAccountStandingFinancialAccountStateSubstatusDelinquent           StatementAccountStandingFinancialAccountStateSubstatus = "DELINQUENT"
)

func (r StatementAccountStandingFinancialAccountStateSubstatus) IsKnown() bool {
	switch r {
	case StatementAccountStandingFinancialAccountStateSubstatusChargedOffDelinquent, StatementAccountStandingFinancialAccountStateSubstatusChargedOffFraud, StatementAccountStandingFinancialAccountStateSubstatusEndUserRequest, StatementAccountStandingFinancialAccountStateSubstatusBankRequest, StatementAccountStandingFinancialAccountStateSubstatusDelinquent:
		return true
	}
	return false
}

type StatementAccountStandingPeriodState string

const (
	StatementAccountStandingPeriodStateStandard StatementAccountStandingPeriodState = "STANDARD"
	StatementAccountStandingPeriodStatePromo    StatementAccountStandingPeriodState = "PROMO"
	StatementAccountStandingPeriodStatePenalty  StatementAccountStandingPeriodState = "PENALTY"
)

func (r StatementAccountStandingPeriodState) IsKnown() bool {
	switch r {
	case StatementAccountStandingPeriodStateStandard, StatementAccountStandingPeriodStatePromo, StatementAccountStandingPeriodStatePenalty:
		return true
	}
	return false
}

type StatementAmountDue struct {
	// Payment due at the end of the billing period in cents. Negative amount indicates
	// something is owed. If the amount owed is positive there was a net credit. If
	// auto-collections are enabled this is the amount that will be requested on the
	// payment due date
	Amount int64 `json:"amount,required"`
	// Amount past due for statement in cents
	PastDue int64                  `json:"past_due,required"`
	JSON    statementAmountDueJSON `json:"-"`
}

// statementAmountDueJSON contains the JSON metadata for the struct
// [StatementAmountDue]
type statementAmountDueJSON struct {
	Amount      apijson.Field
	PastDue     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatementAmountDue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementAmountDueJSON) RawJSON() string {
	return r.raw
}

type StatementPeriodTotals struct {
	// Opening balance transferred from previous account in cents
	BalanceTransfers int64 `json:"balance_transfers,required"`
	// ATM and cashback transactions in cents
	CashAdvances int64 `json:"cash_advances,required"`
	// Volume of credit management operation transactions less any balance transfers in
	// cents
	Credits int64 `json:"credits,required"`
	// Volume of debit management operation transactions less any interest in cents
	Debits int64 `json:"debits,required"`
	// Volume of debit management operation transactions less any interest in cents
	Fees int64 `json:"fees,required"`
	// Interest accrued in cents
	Interest int64 `json:"interest,required"`
	// Any funds transfers which affective the balance in cents
	Payments int64 `json:"payments,required"`
	// Net card transaction volume less any cash advances in cents
	Purchases int64 `json:"purchases,required"`
	// Breakdown of credits
	CreditDetails interface{} `json:"credit_details"`
	// Breakdown of debits
	DebitDetails interface{} `json:"debit_details"`
	// Breakdown of payments
	PaymentDetails interface{}               `json:"payment_details"`
	JSON           statementPeriodTotalsJSON `json:"-"`
}

// statementPeriodTotalsJSON contains the JSON metadata for the struct
// [StatementPeriodTotals]
type statementPeriodTotalsJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Credits          apijson.Field
	Debits           apijson.Field
	Fees             apijson.Field
	Interest         apijson.Field
	Payments         apijson.Field
	Purchases        apijson.Field
	CreditDetails    apijson.Field
	DebitDetails     apijson.Field
	PaymentDetails   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StatementPeriodTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementPeriodTotalsJSON) RawJSON() string {
	return r.raw
}

type StatementStatementType string

const (
	StatementStatementTypeInitial   StatementStatementType = "INITIAL"
	StatementStatementTypePeriodEnd StatementStatementType = "PERIOD_END"
	StatementStatementTypeFinal     StatementStatementType = "FINAL"
)

func (r StatementStatementType) IsKnown() bool {
	switch r {
	case StatementStatementTypeInitial, StatementStatementTypePeriodEnd, StatementStatementTypeFinal:
		return true
	}
	return false
}

type StatementYtdTotals struct {
	// Opening balance transferred from previous account in cents
	BalanceTransfers int64 `json:"balance_transfers,required"`
	// ATM and cashback transactions in cents
	CashAdvances int64 `json:"cash_advances,required"`
	// Volume of credit management operation transactions less any balance transfers in
	// cents
	Credits int64 `json:"credits,required"`
	// Volume of debit management operation transactions less any interest in cents
	Debits int64 `json:"debits,required"`
	// Volume of debit management operation transactions less any interest in cents
	Fees int64 `json:"fees,required"`
	// Interest accrued in cents
	Interest int64 `json:"interest,required"`
	// Any funds transfers which affective the balance in cents
	Payments int64 `json:"payments,required"`
	// Net card transaction volume less any cash advances in cents
	Purchases int64 `json:"purchases,required"`
	// Breakdown of credits
	CreditDetails interface{} `json:"credit_details"`
	// Breakdown of debits
	DebitDetails interface{} `json:"debit_details"`
	// Breakdown of payments
	PaymentDetails interface{}            `json:"payment_details"`
	JSON           statementYtdTotalsJSON `json:"-"`
}

// statementYtdTotalsJSON contains the JSON metadata for the struct
// [StatementYtdTotals]
type statementYtdTotalsJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Credits          apijson.Field
	Debits           apijson.Field
	Fees             apijson.Field
	Interest         apijson.Field
	Payments         apijson.Field
	Purchases        apijson.Field
	CreditDetails    apijson.Field
	DebitDetails     apijson.Field
	PaymentDetails   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StatementYtdTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementYtdTotalsJSON) RawJSON() string {
	return r.raw
}

type StatementInterestDetails struct {
	ActualInterestCharged     int64                                             `json:"actual_interest_charged,required,nullable"`
	DailyBalanceAmounts       StatementInterestDetailsDailyBalanceAmounts       `json:"daily_balance_amounts,required"`
	EffectiveApr              StatementInterestDetailsEffectiveApr              `json:"effective_apr,required"`
	InterestCalculationMethod StatementInterestDetailsInterestCalculationMethod `json:"interest_calculation_method,required"`
	InterestForPeriod         StatementInterestDetailsInterestForPeriod         `json:"interest_for_period,required"`
	PrimeRate                 string                                            `json:"prime_rate,required,nullable"`
	MinimumInterestCharged    int64                                             `json:"minimum_interest_charged,nullable"`
	JSON                      statementInterestDetailsJSON                      `json:"-"`
}

// statementInterestDetailsJSON contains the JSON metadata for the struct
// [StatementInterestDetails]
type statementInterestDetailsJSON struct {
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

func (r *StatementInterestDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementInterestDetailsJSON) RawJSON() string {
	return r.raw
}

type StatementInterestDetailsDailyBalanceAmounts struct {
	BalanceTransfers string                                          `json:"balance_transfers,required"`
	CashAdvances     string                                          `json:"cash_advances,required"`
	Purchases        string                                          `json:"purchases,required"`
	JSON             statementInterestDetailsDailyBalanceAmountsJSON `json:"-"`
}

// statementInterestDetailsDailyBalanceAmountsJSON contains the JSON metadata for
// the struct [StatementInterestDetailsDailyBalanceAmounts]
type statementInterestDetailsDailyBalanceAmountsJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Purchases        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StatementInterestDetailsDailyBalanceAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementInterestDetailsDailyBalanceAmountsJSON) RawJSON() string {
	return r.raw
}

type StatementInterestDetailsEffectiveApr struct {
	BalanceTransfers string                                   `json:"balance_transfers,required"`
	CashAdvances     string                                   `json:"cash_advances,required"`
	Purchases        string                                   `json:"purchases,required"`
	JSON             statementInterestDetailsEffectiveAprJSON `json:"-"`
}

// statementInterestDetailsEffectiveAprJSON contains the JSON metadata for the
// struct [StatementInterestDetailsEffectiveApr]
type statementInterestDetailsEffectiveAprJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Purchases        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StatementInterestDetailsEffectiveApr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementInterestDetailsEffectiveAprJSON) RawJSON() string {
	return r.raw
}

type StatementInterestDetailsInterestCalculationMethod string

const (
	StatementInterestDetailsInterestCalculationMethodDaily        StatementInterestDetailsInterestCalculationMethod = "DAILY"
	StatementInterestDetailsInterestCalculationMethodAverageDaily StatementInterestDetailsInterestCalculationMethod = "AVERAGE_DAILY"
)

func (r StatementInterestDetailsInterestCalculationMethod) IsKnown() bool {
	switch r {
	case StatementInterestDetailsInterestCalculationMethodDaily, StatementInterestDetailsInterestCalculationMethodAverageDaily:
		return true
	}
	return false
}

type StatementInterestDetailsInterestForPeriod struct {
	BalanceTransfers string                                        `json:"balance_transfers,required"`
	CashAdvances     string                                        `json:"cash_advances,required"`
	Purchases        string                                        `json:"purchases,required"`
	JSON             statementInterestDetailsInterestForPeriodJSON `json:"-"`
}

// statementInterestDetailsInterestForPeriodJSON contains the JSON metadata for the
// struct [StatementInterestDetailsInterestForPeriod]
type statementInterestDetailsInterestForPeriodJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Purchases        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StatementInterestDetailsInterestForPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementInterestDetailsInterestForPeriodJSON) RawJSON() string {
	return r.raw
}

type Statements struct {
	Data    []Statement    `json:"data,required"`
	HasMore bool           `json:"has_more,required"`
	JSON    statementsJSON `json:"-"`
}

// statementsJSON contains the JSON metadata for the struct [Statements]
type statementsJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Statements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementsJSON) RawJSON() string {
	return r.raw
}

type FinancialAccountStatementListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included.
	Begin param.Field[time.Time] `query:"begin" format:"date"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included.
	End param.Field[time.Time] `query:"end" format:"date"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Whether to include the initial statement. It is not included by default.
	IncludeInitialStatements param.Field[bool] `query:"include_initial_statements"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [FinancialAccountStatementListParams]'s query parameters as
// `url.Values`.
func (r FinancialAccountStatementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
