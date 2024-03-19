// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// FinancialAccountStatementService contains methods and other services that help
// with interacting with the lithic API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewFinancialAccountStatementService] method instead.
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
	path := fmt.Sprintf("financial_accounts/%s/statements/%s", financialAccountToken, statementToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the statements for a given financial account.
func (r *FinancialAccountStatementService) List(ctx context.Context, financialAccountToken string, query FinancialAccountStatementListParams, opts ...option.RequestOption) (res *shared.CursorPage[Statement], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *FinancialAccountStatementService) ListAutoPaging(ctx context.Context, financialAccountToken string, query FinancialAccountStatementListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[Statement] {
	return shared.NewCursorPageAutoPager(r.List(ctx, financialAccountToken, query, opts...))
}

type Statement struct {
	// Globally unique identifier for a statement
	Token string `json:"token,required"`
	// Total payments during this billing period.
	ACHPeriodTotal int64 `json:"ach_period_total,required"`
	// Year-to-date settled payment total
	ACHYtdTotal int64 `json:"ach_ytd_total,required"`
	// Total adjustments during this billing period.
	AdjustmentsPeriodTotal int64 `json:"adjustments_period_total,required"`
	// Year-to-date settled adjustments total
	AdjustmentsYtdTotal int64 `json:"adjustments_ytd_total,required"`
	// Payment due at the end of the billing period. Negative amount indicates
	// something is owed. If the amount owed is positive (e.g., there was a net
	// credit), then payment should be returned to the cardholder via ACH.
	AmountDue int64 `json:"amount_due,required"`
	// Amount of credit available to spend
	AvailableCredit int64 `json:"available_credit,required"`
	// Timestamp of when the statement was created
	Created time.Time `json:"created,required" format:"date-time"`
	// For prepay accounts, this is the minimum prepay balance that must be maintained.
	// For charge card accounts, this is the maximum credit balance extended by a
	// lender.
	CreditLimit int64 `json:"credit_limit,required"`
	// Number of days in the billing cycle
	DaysInBillingCycle int64 `json:"days_in_billing_cycle,required"`
	// Balance at the end of the billing period. For charge cards, this should be the
	// same at the statement amount due.
	EndingBalance int64 `json:"ending_balance,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Date when the payment is due
	PaymentDueDate time.Time `json:"payment_due_date,required" format:"date"`
	// Total settled card transactions during this billing period, determined by
	// liability date.
	PurchasesPeriodTotal int64 `json:"purchases_period_total,required"`
	// Year-to-date settled card transaction total
	PurchasesYtdTotal int64 `json:"purchases_ytd_total,required"`
	// Balance at the start of the billing period
	StartingBalance int64 `json:"starting_balance,required"`
	// Date when the billing period ended
	StatementEndDate time.Time `json:"statement_end_date,required" format:"date"`
	// Date when the billing period began
	StatementStartDate time.Time `json:"statement_start_date,required" format:"date"`
	// Timestamp of when the statement was updated
	Updated time.Time     `json:"updated,required" format:"date-time"`
	JSON    statementJSON `json:"-"`
}

// statementJSON contains the JSON metadata for the struct [Statement]
type statementJSON struct {
	Token                  apijson.Field
	ACHPeriodTotal         apijson.Field
	ACHYtdTotal            apijson.Field
	AdjustmentsPeriodTotal apijson.Field
	AdjustmentsYtdTotal    apijson.Field
	AmountDue              apijson.Field
	AvailableCredit        apijson.Field
	Created                apijson.Field
	CreditLimit            apijson.Field
	DaysInBillingCycle     apijson.Field
	EndingBalance          apijson.Field
	FinancialAccountToken  apijson.Field
	PaymentDueDate         apijson.Field
	PurchasesPeriodTotal   apijson.Field
	PurchasesYtdTotal      apijson.Field
	StartingBalance        apijson.Field
	StatementEndDate       apijson.Field
	StatementStartDate     apijson.Field
	Updated                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *Statement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementJSON) RawJSON() string {
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
