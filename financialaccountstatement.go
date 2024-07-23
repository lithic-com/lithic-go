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
	opts = append(r.Options[:], opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	if statementToken == "" {
		err = errors.New("missing required statement_token parameter")
		return
	}
	path := fmt.Sprintf("financial_accounts/%s/statements/%s", financialAccountToken, statementToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the statements for a given financial account.
func (r *FinancialAccountStatementService) List(ctx context.Context, financialAccountToken string, query FinancialAccountStatementListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Statement], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("financial_accounts/%s/statements", financialAccountToken)
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
	// Payment due at the end of the billing period. Negative amount indicates
	// something is owed. If the amount owed is positive (e.g., there was a net
	// credit), then payment should be returned to the cardholder via ACH.
	AmountDue int64 `json:"amount_due,required"`
	// Payment past due at the end of the billing period.
	AmountPastDue int64 `json:"amount_past_due,required"`
	// Amount of credit available to spend
	AvailableCredit int64 `json:"available_credit,required"`
	// Timestamp of when the statement was created
	Created time.Time `json:"created,required" format:"date-time"`
	// For prepay accounts, this is the minimum prepay balance that must be maintained.
	// For charge card accounts, this is the maximum credit balance extended by a
	// lender.
	CreditLimit int64 `json:"credit_limit,required"`
	// Globally unique identifier for a credit product
	CreditProductToken string `json:"credit_product_token,required"`
	// Number of days in the billing cycle
	DaysInBillingCycle int64 `json:"days_in_billing_cycle,required"`
	// Balance at the end of the billing period. For charge cards, this should be the
	// same at the statement amount due.
	EndingBalance int64 `json:"ending_balance,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Date of when the next statement will be created
	NextStatementDate time.Time `json:"next_statement_date,required" format:"date"`
	// Date when the payment is due
	PaymentDueDate time.Time             `json:"payment_due_date,required" format:"date"`
	PeriodTotals   StatementPeriodTotals `json:"period_totals,required"`
	// Balance at the start of the billing period
	StartingBalance int64 `json:"starting_balance,required"`
	// Date when the billing period ended
	StatementEndDate time.Time `json:"statement_end_date,required" format:"date"`
	// Date when the billing period began
	StatementStartDate time.Time `json:"statement_start_date,required" format:"date"`
	// Timestamp of when the statement was updated
	Updated   time.Time          `json:"updated,required" format:"date-time"`
	YtdTotals StatementYtdTotals `json:"ytd_totals,required"`
	JSON      statementJSON      `json:"-"`
}

// statementJSON contains the JSON metadata for the struct [Statement]
type statementJSON struct {
	Token                 apijson.Field
	AccountStanding       apijson.Field
	AmountDue             apijson.Field
	AmountPastDue         apijson.Field
	AvailableCredit       apijson.Field
	Created               apijson.Field
	CreditLimit           apijson.Field
	CreditProductToken    apijson.Field
	DaysInBillingCycle    apijson.Field
	EndingBalance         apijson.Field
	FinancialAccountToken apijson.Field
	NextStatementDate     apijson.Field
	PaymentDueDate        apijson.Field
	PeriodTotals          apijson.Field
	StartingBalance       apijson.Field
	StatementEndDate      apijson.Field
	StatementStartDate    apijson.Field
	Updated               apijson.Field
	YtdTotals             apijson.Field
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
	// Current overall period number
	PeriodNumber int64                         `json:"period_number,required"`
	State        StatementAccountStandingState `json:"state,required"`
	JSON         statementAccountStandingJSON  `json:"-"`
}

// statementAccountStandingJSON contains the JSON metadata for the struct
// [StatementAccountStanding]
type statementAccountStandingJSON struct {
	PeriodNumber apijson.Field
	State        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StatementAccountStanding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementAccountStandingJSON) RawJSON() string {
	return r.raw
}

type StatementAccountStandingState string

const (
	StatementAccountStandingStateStandard StatementAccountStandingState = "STANDARD"
	StatementAccountStandingStatePromo    StatementAccountStandingState = "PROMO"
	StatementAccountStandingStatePenalty  StatementAccountStandingState = "PENALTY"
)

func (r StatementAccountStandingState) IsKnown() bool {
	switch r {
	case StatementAccountStandingStateStandard, StatementAccountStandingStatePromo, StatementAccountStandingStatePenalty:
		return true
	}
	return false
}

type StatementPeriodTotals struct {
	BalanceTransfers int64                     `json:"balance_transfers,required"`
	CashAdvances     int64                     `json:"cash_advances,required"`
	Credits          int64                     `json:"credits,required"`
	Fees             int64                     `json:"fees,required"`
	Interest         int64                     `json:"interest,required"`
	Payments         int64                     `json:"payments,required"`
	Purchases        int64                     `json:"purchases,required"`
	JSON             statementPeriodTotalsJSON `json:"-"`
}

// statementPeriodTotalsJSON contains the JSON metadata for the struct
// [StatementPeriodTotals]
type statementPeriodTotalsJSON struct {
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

func (r *StatementPeriodTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementPeriodTotalsJSON) RawJSON() string {
	return r.raw
}

type StatementYtdTotals struct {
	BalanceTransfers int64                  `json:"balance_transfers,required"`
	CashAdvances     int64                  `json:"cash_advances,required"`
	Credits          int64                  `json:"credits,required"`
	Fees             int64                  `json:"fees,required"`
	Interest         int64                  `json:"interest,required"`
	Payments         int64                  `json:"payments,required"`
	Purchases        int64                  `json:"purchases,required"`
	JSON             statementYtdTotalsJSON `json:"-"`
}

// statementYtdTotalsJSON contains the JSON metadata for the struct
// [StatementYtdTotals]
type statementYtdTotalsJSON struct {
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

func (r *StatementYtdTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementYtdTotalsJSON) RawJSON() string {
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
