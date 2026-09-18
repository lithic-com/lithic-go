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

// FinancialAccountInstallmentPlanService contains methods and other services that
// help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountInstallmentPlanService] method instead.
type FinancialAccountInstallmentPlanService struct {
	Options []option.RequestOption
}

// NewFinancialAccountInstallmentPlanService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFinancialAccountInstallmentPlanService(opts ...option.RequestOption) (r *FinancialAccountInstallmentPlanService) {
	r = &FinancialAccountInstallmentPlanService{}
	r.Options = opts
	return
}

// Get a specific installment plan for a given financial account.
func (r *FinancialAccountInstallmentPlanService) Get(ctx context.Context, financialAccountToken string, installmentPlanToken string, opts ...option.RequestOption) (res *InstallmentPlan, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	if installmentPlanToken == "" {
		err = errors.New("missing required installment_plan_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/installment_plans/%s", financialAccountToken, installmentPlanToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the installment plans for a given financial account.
func (r *FinancialAccountInstallmentPlanService) List(ctx context.Context, financialAccountToken string, query FinancialAccountInstallmentPlanListParams, opts ...option.RequestOption) (res *pagination.CursorPage[InstallmentPlan], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/installment_plans", financialAccountToken)
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

// List the installment plans for a given financial account.
func (r *FinancialAccountInstallmentPlanService) ListAutoPaging(ctx context.Context, financialAccountToken string, query FinancialAccountInstallmentPlanListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[InstallmentPlan] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, financialAccountToken, query, opts...))
}

type InstallmentPlan struct {
	// Globally unique identifier for an installment plan
	Token string `json:"token" api:"required"`
	// Date the plan was paid off or cancelled, or null while it is still open
	ClosedAt time.Time `json:"closed_at" api:"required,nullable" format:"date"`
	// Timestamp of when the installment plan was created
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Enrollment fee charged when the plan was created in cents
	FeeAmount int64 `json:"fee_amount" api:"required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string `json:"financial_account_token" api:"required" format:"uuid"`
	// Total owed on the plan in cents, the principal amount plus the enrollment fee
	InstallmentPlanTotal int64 `json:"installment_plan_total" api:"required"`
	// Installments that make up the plan, oldest first
	Installments []InstallmentPlanInstallment `json:"installments" api:"required"`
	// Number of installments that still carry a balance
	InstallmentsOutstanding int64 `json:"installments_outstanding" api:"required"`
	// Number of installments that have been paid off
	InstallmentsPaid int64 `json:"installments_paid" api:"required"`
	// Number of installments the plan is broken into
	NumInstallments int64 `json:"num_installments" api:"required"`
	// Balance the plan was opened on in cents, excluding the enrollment fee
	PrincipalAmount int64 `json:"principal_amount" api:"required"`
	// Balance the plan was opened on, broken out by category, or null if it was not
	// recorded
	SourceAmounts TransactionCategoryBalances `json:"source_amounts" api:"required,nullable"`
	// Identifier of the record the plan was opened from, such as the closing statement
	// for an unpaid balance
	SourceID   string                    `json:"source_id" api:"required"`
	SourceType InstallmentPlanSourceType `json:"source_type" api:"required"`
	// Date the plan was created
	StartDate time.Time `json:"start_date" api:"required" format:"date"`
	// State of the installment plan. A plan is REBUILD_IN_PROGRESS while its loan
	// tapes are being rebuilt, during which its payment totals are being recomputed
	// and should not be treated as final
	State InstallmentPlanState `json:"state" api:"required"`
	// Amount paid towards the plan to date in cents
	TotalPaid int64 `json:"total_paid" api:"required"`
	// Timestamp of when the installment plan was updated
	Updated time.Time           `json:"updated" api:"required" format:"date-time"`
	JSON    installmentPlanJSON `json:"-"`
}

// installmentPlanJSON contains the JSON metadata for the struct [InstallmentPlan]
type installmentPlanJSON struct {
	Token                   apijson.Field
	ClosedAt                apijson.Field
	Created                 apijson.Field
	FeeAmount               apijson.Field
	FinancialAccountToken   apijson.Field
	InstallmentPlanTotal    apijson.Field
	Installments            apijson.Field
	InstallmentsOutstanding apijson.Field
	InstallmentsPaid        apijson.Field
	NumInstallments         apijson.Field
	PrincipalAmount         apijson.Field
	SourceAmounts           apijson.Field
	SourceID                apijson.Field
	SourceType              apijson.Field
	StartDate               apijson.Field
	State                   apijson.Field
	TotalPaid               apijson.Field
	Updated                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InstallmentPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r installmentPlanJSON) RawJSON() string {
	return r.raw
}

type InstallmentPlanInstallment struct {
	// Amount the installment was opened for in cents
	AmountDue        int64                       `json:"amount_due" api:"required"`
	AmountDueDetails TransactionCategoryBalances `json:"amount_due_details" api:"required"`
	// Amount still owed on the installment in cents
	AmountOutstanding        int64                       `json:"amount_outstanding" api:"required"`
	AmountOutstandingDetails TransactionCategoryBalances `json:"amount_outstanding_details" api:"required"`
	// Amount paid towards the installment in cents
	AmountPaid        int64                       `json:"amount_paid" api:"required"`
	AmountPaidDetails TransactionCategoryBalances `json:"amount_paid_details" api:"required"`
	// Date the installment was actually assessed onto the account, or null if it has
	// not been assessed yet
	DateAssessed time.Time `json:"date_assessed" api:"required,nullable" format:"date"`
	// Date the installment is scheduled to be assessed onto the account
	DueDate time.Time `json:"due_date" api:"required" format:"date"`
	// Position of this installment within the plan, starting at 0
	InstallmentNum int64 `json:"installment_num" api:"required"`
	// Date the installment must be paid by before it is considered past due
	PaymentDueDate time.Time `json:"payment_due_date" api:"required" format:"date"`
	// Payments applied to this installment, oldest first
	Payments []InstallmentPlanInstallmentsPayment `json:"payments" api:"required"`
	JSON     installmentPlanInstallmentJSON       `json:"-"`
}

// installmentPlanInstallmentJSON contains the JSON metadata for the struct
// [InstallmentPlanInstallment]
type installmentPlanInstallmentJSON struct {
	AmountDue                apijson.Field
	AmountDueDetails         apijson.Field
	AmountOutstanding        apijson.Field
	AmountOutstandingDetails apijson.Field
	AmountPaid               apijson.Field
	AmountPaidDetails        apijson.Field
	DateAssessed             apijson.Field
	DueDate                  apijson.Field
	InstallmentNum           apijson.Field
	PaymentDueDate           apijson.Field
	Payments                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *InstallmentPlanInstallment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r installmentPlanInstallmentJSON) RawJSON() string {
	return r.raw
}

type InstallmentPlanInstallmentsPayment struct {
	// Amount applied to the installment in cents
	Amount        int64                       `json:"amount" api:"required"`
	AmountDetails TransactionCategoryBalances `json:"amount_details" api:"required"`
	// Date the payment was applied to the installment
	Date time.Time                              `json:"date" api:"required" format:"date"`
	JSON installmentPlanInstallmentsPaymentJSON `json:"-"`
}

// installmentPlanInstallmentsPaymentJSON contains the JSON metadata for the struct
// [InstallmentPlanInstallmentsPayment]
type installmentPlanInstallmentsPaymentJSON struct {
	Amount        apijson.Field
	AmountDetails apijson.Field
	Date          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InstallmentPlanInstallmentsPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r installmentPlanInstallmentsPaymentJSON) RawJSON() string {
	return r.raw
}

type InstallmentPlanSourceType string

const (
	InstallmentPlanSourceTypeUnpaidBalance InstallmentPlanSourceType = "UNPAID_BALANCE"
	InstallmentPlanSourceTypeTransaction   InstallmentPlanSourceType = "TRANSACTION"
)

func (r InstallmentPlanSourceType) IsKnown() bool {
	switch r {
	case InstallmentPlanSourceTypeUnpaidBalance, InstallmentPlanSourceTypeTransaction:
		return true
	}
	return false
}

// State of the installment plan. A plan is REBUILD_IN_PROGRESS while its loan
// tapes are being rebuilt, during which its payment totals are being recomputed
// and should not be treated as final
type InstallmentPlanState string

const (
	InstallmentPlanStatePending           InstallmentPlanState = "PENDING"
	InstallmentPlanStateActive            InstallmentPlanState = "ACTIVE"
	InstallmentPlanStateRebuildInProgress InstallmentPlanState = "REBUILD_IN_PROGRESS"
	InstallmentPlanStateFullyPaid         InstallmentPlanState = "FULLY_PAID"
	InstallmentPlanStateCancelled         InstallmentPlanState = "CANCELLED"
)

func (r InstallmentPlanState) IsKnown() bool {
	switch r {
	case InstallmentPlanStatePending, InstallmentPlanStateActive, InstallmentPlanStateRebuildInProgress, InstallmentPlanStateFullyPaid, InstallmentPlanStateCancelled:
		return true
	}
	return false
}

type TransactionCategoryBalances struct {
	// Amounts attributable to balance transfers
	BalanceTransfers CategoryBalances `json:"balance_transfers" api:"required"`
	// Amounts attributable to cash advances
	CashAdvances CategoryBalances `json:"cash_advances" api:"required"`
	// Amounts attributable to purchases
	Purchases CategoryBalances                `json:"purchases" api:"required"`
	JSON      transactionCategoryBalancesJSON `json:"-"`
}

// transactionCategoryBalancesJSON contains the JSON metadata for the struct
// [TransactionCategoryBalances]
type transactionCategoryBalancesJSON struct {
	BalanceTransfers apijson.Field
	CashAdvances     apijson.Field
	Purchases        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionCategoryBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionCategoryBalancesJSON) RawJSON() string {
	return r.raw
}

type FinancialAccountInstallmentPlanListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Only installment plans in this state will be included.
	State param.Field[FinancialAccountInstallmentPlanListParamsState] `query:"state"`
}

// URLQuery serializes [FinancialAccountInstallmentPlanListParams]'s query
// parameters as `url.Values`.
func (r FinancialAccountInstallmentPlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only installment plans in this state will be included.
type FinancialAccountInstallmentPlanListParamsState string

const (
	FinancialAccountInstallmentPlanListParamsStatePending           FinancialAccountInstallmentPlanListParamsState = "PENDING"
	FinancialAccountInstallmentPlanListParamsStateActive            FinancialAccountInstallmentPlanListParamsState = "ACTIVE"
	FinancialAccountInstallmentPlanListParamsStateRebuildInProgress FinancialAccountInstallmentPlanListParamsState = "REBUILD_IN_PROGRESS"
	FinancialAccountInstallmentPlanListParamsStateFullyPaid         FinancialAccountInstallmentPlanListParamsState = "FULLY_PAID"
	FinancialAccountInstallmentPlanListParamsStateCancelled         FinancialAccountInstallmentPlanListParamsState = "CANCELLED"
)

func (r FinancialAccountInstallmentPlanListParamsState) IsKnown() bool {
	switch r {
	case FinancialAccountInstallmentPlanListParamsStatePending, FinancialAccountInstallmentPlanListParamsStateActive, FinancialAccountInstallmentPlanListParamsStateRebuildInProgress, FinancialAccountInstallmentPlanListParamsStateFullyPaid, FinancialAccountInstallmentPlanListParamsStateCancelled:
		return true
	}
	return false
}
