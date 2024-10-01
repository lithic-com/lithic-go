// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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

type LoanTape struct {
	// Globally unique identifier for a loan tape
	Token           string                  `json:"token,required"`
	AccountStanding LoanTapeAccountStanding `json:"account_standing,required"`
	// Amount due for the prior billing cycle. Any amounts not fully paid off on this
	// due date will be considered past due the next day
	BalanceDue LoanTapeBalanceDue `json:"balance_due,required"`
	// Amount due for the current billing cycle. Any amounts not paid off by early
	// payments or credits will be considered due at the end of the current billing
	// period
	BalanceNextDue LoanTapeBalanceNextDue `json:"balance_next_due,required"`
	// Amount not paid off on previous due dates
	BalancePastDue LoanTapeBalancePastDue `json:"balance_past_due,required"`
	// Timestamp of when the loan tape was created
	Created time.Time `json:"created,required" format:"date-time"`
	// For prepay accounts, this is the minimum prepay balance that must be maintained.
	// For charge card accounts, this is the maximum credit balance extended by a
	// lender
	CreditLimit int64 `json:"credit_limit,required"`
	// Globally unique identifier for a credit product
	CreditProductToken string `json:"credit_product_token,required"`
	// Date of transactions that this loan tape covers
	Date      time.Time         `json:"date,required" format:"date"`
	DayTotals LoanTapeDayTotals `json:"day_totals,required"`
	// Excess credits in the form of provisional credits, payments, or purchase
	// refunds. If positive, the account is in net credit state with no outstanding
	// balances. An overpayment could land an account in this state
	ExcessCredits int64 `json:"excess_credits,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string                        `json:"financial_account_token,required" format:"uuid"`
	MinimumPaymentBalance LoanTapeMinimumPaymentBalance `json:"minimum_payment_balance,required"`
	PaymentAllocation     LoanTapePaymentAllocation     `json:"payment_allocation,required"`
	PeriodTotals          LoanTapePeriodTotals          `json:"period_totals,required"`
	StatementBalance      LoanTapeStatementBalance      `json:"statement_balance,required"`
	// Timestamp of when the loan tape was updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Version number of the loan tape. This starts at 1
	Version   int64             `json:"version,required"`
	YtdTotals LoanTapeYtdTotals `json:"ytd_totals,required"`
	// Interest tier to which this account belongs to
	Tier string       `json:"tier"`
	JSON loanTapeJSON `json:"-"`
}

// loanTapeJSON contains the JSON metadata for the struct [LoanTape]
type loanTapeJSON struct {
	Token                 apijson.Field
	AccountStanding       apijson.Field
	BalanceDue            apijson.Field
	BalanceNextDue        apijson.Field
	BalancePastDue        apijson.Field
	Created               apijson.Field
	CreditLimit           apijson.Field
	CreditProductToken    apijson.Field
	Date                  apijson.Field
	DayTotals             apijson.Field
	ExcessCredits         apijson.Field
	FinancialAccountToken apijson.Field
	MinimumPaymentBalance apijson.Field
	PaymentAllocation     apijson.Field
	PeriodTotals          apijson.Field
	StatementBalance      apijson.Field
	Updated               apijson.Field
	Version               apijson.Field
	YtdTotals             apijson.Field
	Tier                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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

// Amount due for the prior billing cycle. Any amounts not fully paid off on this
// due date will be considered past due the next day
type LoanTapeBalanceDue struct {
	Fees      int64                  `json:"fees,required"`
	Interest  int64                  `json:"interest,required"`
	Principal int64                  `json:"principal,required"`
	JSON      loanTapeBalanceDueJSON `json:"-"`
}

// loanTapeBalanceDueJSON contains the JSON metadata for the struct
// [LoanTapeBalanceDue]
type loanTapeBalanceDueJSON struct {
	Fees        apijson.Field
	Interest    apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapeBalanceDue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeBalanceDueJSON) RawJSON() string {
	return r.raw
}

// Amount due for the current billing cycle. Any amounts not paid off by early
// payments or credits will be considered due at the end of the current billing
// period
type LoanTapeBalanceNextDue struct {
	Fees      int64                      `json:"fees,required"`
	Interest  int64                      `json:"interest,required"`
	Principal int64                      `json:"principal,required"`
	JSON      loanTapeBalanceNextDueJSON `json:"-"`
}

// loanTapeBalanceNextDueJSON contains the JSON metadata for the struct
// [LoanTapeBalanceNextDue]
type loanTapeBalanceNextDueJSON struct {
	Fees        apijson.Field
	Interest    apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapeBalanceNextDue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeBalanceNextDueJSON) RawJSON() string {
	return r.raw
}

// Amount not paid off on previous due dates
type LoanTapeBalancePastDue struct {
	Fees      int64                      `json:"fees,required"`
	Interest  int64                      `json:"interest,required"`
	Principal int64                      `json:"principal,required"`
	JSON      loanTapeBalancePastDueJSON `json:"-"`
}

// loanTapeBalancePastDueJSON contains the JSON metadata for the struct
// [LoanTapeBalancePastDue]
type loanTapeBalancePastDueJSON struct {
	Fees        apijson.Field
	Interest    apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapeBalancePastDue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeBalancePastDueJSON) RawJSON() string {
	return r.raw
}

type LoanTapeDayTotals struct {
	// Opening balance transferred from previous account in cents
	BalanceTransfers int64 `json:"balance_transfers,required"`
	// ATM and cashback transactions in cents
	CashAdvances int64 `json:"cash_advances,required"`
	// Volume of credit management operation transactions less any balance transfers in
	// cents
	Credits int64 `json:"credits,required"`
	// Volume of debit management operation transactions less any interest in cents
	Fees int64 `json:"fees,required"`
	// Interest accrued in cents
	Interest int64 `json:"interest,required"`
	// Any funds transfers which affective the balance in cents
	Payments int64 `json:"payments,required"`
	// Net card transaction volume less any cash advances in cents
	Purchases int64                 `json:"purchases,required"`
	JSON      loanTapeDayTotalsJSON `json:"-"`
}

// loanTapeDayTotalsJSON contains the JSON metadata for the struct
// [LoanTapeDayTotals]
type loanTapeDayTotalsJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Credits          apijson.Field
	Fees             apijson.Field
	Interest         apijson.Field
	Payments         apijson.Field
	Purchases        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoanTapeDayTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeDayTotalsJSON) RawJSON() string {
	return r.raw
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
	Fees      int64                         `json:"fees,required"`
	Interest  int64                         `json:"interest,required"`
	Principal int64                         `json:"principal,required"`
	JSON      loanTapePaymentAllocationJSON `json:"-"`
}

// loanTapePaymentAllocationJSON contains the JSON metadata for the struct
// [LoanTapePaymentAllocation]
type loanTapePaymentAllocationJSON struct {
	Fees        apijson.Field
	Interest    apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapePaymentAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapePaymentAllocationJSON) RawJSON() string {
	return r.raw
}

type LoanTapePeriodTotals struct {
	// Opening balance transferred from previous account in cents
	BalanceTransfers int64 `json:"balance_transfers,required"`
	// ATM and cashback transactions in cents
	CashAdvances int64 `json:"cash_advances,required"`
	// Volume of credit management operation transactions less any balance transfers in
	// cents
	Credits int64 `json:"credits,required"`
	// Volume of debit management operation transactions less any interest in cents
	Fees int64 `json:"fees,required"`
	// Interest accrued in cents
	Interest int64 `json:"interest,required"`
	// Any funds transfers which affective the balance in cents
	Payments int64 `json:"payments,required"`
	// Net card transaction volume less any cash advances in cents
	Purchases int64                    `json:"purchases,required"`
	JSON      loanTapePeriodTotalsJSON `json:"-"`
}

// loanTapePeriodTotalsJSON contains the JSON metadata for the struct
// [LoanTapePeriodTotals]
type loanTapePeriodTotalsJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Credits          apijson.Field
	Fees             apijson.Field
	Interest         apijson.Field
	Payments         apijson.Field
	Purchases        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoanTapePeriodTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapePeriodTotalsJSON) RawJSON() string {
	return r.raw
}

type LoanTapeStatementBalance struct {
	Amount    int64                        `json:"amount,required"`
	Remaining int64                        `json:"remaining,required"`
	JSON      loanTapeStatementBalanceJSON `json:"-"`
}

// loanTapeStatementBalanceJSON contains the JSON metadata for the struct
// [LoanTapeStatementBalance]
type loanTapeStatementBalanceJSON struct {
	Amount      apijson.Field
	Remaining   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoanTapeStatementBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeStatementBalanceJSON) RawJSON() string {
	return r.raw
}

type LoanTapeYtdTotals struct {
	// Opening balance transferred from previous account in cents
	BalanceTransfers int64 `json:"balance_transfers,required"`
	// ATM and cashback transactions in cents
	CashAdvances int64 `json:"cash_advances,required"`
	// Volume of credit management operation transactions less any balance transfers in
	// cents
	Credits int64 `json:"credits,required"`
	// Volume of debit management operation transactions less any interest in cents
	Fees int64 `json:"fees,required"`
	// Interest accrued in cents
	Interest int64 `json:"interest,required"`
	// Any funds transfers which affective the balance in cents
	Payments int64 `json:"payments,required"`
	// Net card transaction volume less any cash advances in cents
	Purchases int64                 `json:"purchases,required"`
	JSON      loanTapeYtdTotalsJSON `json:"-"`
}

// loanTapeYtdTotalsJSON contains the JSON metadata for the struct
// [LoanTapeYtdTotals]
type loanTapeYtdTotalsJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Credits          apijson.Field
	Fees             apijson.Field
	Interest         apijson.Field
	Payments         apijson.Field
	Purchases        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoanTapeYtdTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeYtdTotalsJSON) RawJSON() string {
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
