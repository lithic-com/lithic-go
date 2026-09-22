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

// FinancialAccountInstallmentPlanStatementService contains methods and other
// services that help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountInstallmentPlanStatementService] method instead.
type FinancialAccountInstallmentPlanStatementService struct {
	Options []option.RequestOption
}

// NewFinancialAccountInstallmentPlanStatementService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewFinancialAccountInstallmentPlanStatementService(opts ...option.RequestOption) (r *FinancialAccountInstallmentPlanStatementService) {
	r = &FinancialAccountInstallmentPlanStatementService{}
	r.Options = opts
	return
}

// Get a specific statement snapshot for a given installment plan.
func (r *FinancialAccountInstallmentPlanStatementService) Get(ctx context.Context, financialAccountToken string, installmentPlanToken string, statementToken string, opts ...option.RequestOption) (res *InstallmentPlanStatement, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	if installmentPlanToken == "" {
		err = errors.New("missing required installment_plan_token parameter")
		return nil, err
	}
	if statementToken == "" {
		err = errors.New("missing required statement_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/installment_plans/%s/statements/%s", financialAccountToken, installmentPlanToken, statementToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the statement snapshots for a given installment plan.
func (r *FinancialAccountInstallmentPlanStatementService) List(ctx context.Context, financialAccountToken string, installmentPlanToken string, query FinancialAccountInstallmentPlanStatementListParams, opts ...option.RequestOption) (res *pagination.CursorPage[InstallmentPlanStatement], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	if installmentPlanToken == "" {
		err = errors.New("missing required installment_plan_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/installment_plans/%s/statements", financialAccountToken, installmentPlanToken)
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

// List the statement snapshots for a given installment plan.
func (r *FinancialAccountInstallmentPlanStatementService) ListAutoPaging(ctx context.Context, financialAccountToken string, installmentPlanToken string, query FinancialAccountInstallmentPlanStatementListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[InstallmentPlanStatement] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, financialAccountToken, installmentPlanToken, query, opts...))
}

// An immutable snapshot of an installment plan as of the statement it is attached
// to. Lithic cuts one per open plan when a statement is generated and never
// reissues it
type InstallmentPlanStatement struct {
	// Globally unique identifier for this snapshot, which is the token of the
	// statement it is attached to. A plan is snapshotted at most once per statement,
	// so the statement identifies the snapshot within the plan. Pass it as a
	// pagination cursor
	Token string `json:"token" api:"required"`
	// Enrollment fee charged when the plan was created in cents
	FeeAmount int64 `json:"fee_amount" api:"required"`
	// Globally unique identifier for the installment plan this snapshot is of
	InstallmentPlanToken string `json:"installment_plan_token" api:"required"`
	// Total owed on the plan in cents, the principal amount plus the enrollment fee
	InstallmentPlanTotal int64 `json:"installment_plan_total" api:"required"`
	// Installments that make up the plan, oldest first
	Installments []InstallmentPlanStatementInstallment `json:"installments" api:"required"`
	// Number of installments that still carried a balance as of this statement
	InstallmentsOutstanding int64 `json:"installments_outstanding" api:"required"`
	// Number of installments that had been paid off as of this statement
	InstallmentsPaid int64 `json:"installments_paid" api:"required"`
	// Number of installments the plan is broken into
	NumInstallments int64 `json:"num_installments" api:"required"`
	// Balance the plan was opened on in cents, excluding the enrollment fee
	PrincipalAmount int64                              `json:"principal_amount" api:"required"`
	SourceType      InstallmentPlanStatementSourceType `json:"source_type" api:"required"`
	// Date the plan was created
	StartDate time.Time `json:"start_date" api:"required" format:"date"`
	// State of the installment plan. A plan is REBUILD_IN_PROGRESS while its loan
	// tapes are being rebuilt, during which its payment totals are being recomputed
	// and should not be treated as final
	State InstallmentPlanStatementState `json:"state" api:"required"`
	// Amount paid towards the plan as of this statement in cents
	TotalPaid int64                        `json:"total_paid" api:"required"`
	JSON      installmentPlanStatementJSON `json:"-"`
}

// installmentPlanStatementJSON contains the JSON metadata for the struct
// [InstallmentPlanStatement]
type installmentPlanStatementJSON struct {
	Token                   apijson.Field
	FeeAmount               apijson.Field
	InstallmentPlanToken    apijson.Field
	InstallmentPlanTotal    apijson.Field
	Installments            apijson.Field
	InstallmentsOutstanding apijson.Field
	InstallmentsPaid        apijson.Field
	NumInstallments         apijson.Field
	PrincipalAmount         apijson.Field
	SourceType              apijson.Field
	StartDate               apijson.Field
	State                   apijson.Field
	TotalPaid               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InstallmentPlanStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r installmentPlanStatementJSON) RawJSON() string {
	return r.raw
}

// One installment of a plan as of the statement. Amounts are totalled across
// transaction categories rather than broken out by them, which the installment
// plan endpoint does
type InstallmentPlanStatementInstallment struct {
	// Amount the installment was opened for in cents
	AmountDue        int64            `json:"amount_due" api:"required"`
	AmountDueDetails CategoryBalances `json:"amount_due_details" api:"required"`
	// Amount still owed on the installment in cents
	AmountOutstanding        int64            `json:"amount_outstanding" api:"required"`
	AmountOutstandingDetails CategoryBalances `json:"amount_outstanding_details" api:"required"`
	// Amount paid towards the installment in cents
	AmountPaid        int64            `json:"amount_paid" api:"required"`
	AmountPaidDetails CategoryBalances `json:"amount_paid_details" api:"required"`
	// Date the installment was actually assessed onto the account, or null if it had
	// not been assessed as of this statement
	DateAssessed time.Time `json:"date_assessed" api:"required,nullable" format:"date"`
	// Date the installment is scheduled to be assessed onto the account
	DueDate time.Time `json:"due_date" api:"required" format:"date"`
	// Position of this installment within the plan, starting at 0
	InstallmentNum int64 `json:"installment_num" api:"required"`
	// Date the installment must be paid by before it is considered past due
	PaymentDueDate time.Time `json:"payment_due_date" api:"required" format:"date"`
	// Payments applied to this installment, oldest first
	Payments []InstallmentPlanStatementInstallmentsPayment `json:"payments" api:"required"`
	JSON     installmentPlanStatementInstallmentJSON       `json:"-"`
}

// installmentPlanStatementInstallmentJSON contains the JSON metadata for the
// struct [InstallmentPlanStatementInstallment]
type installmentPlanStatementInstallmentJSON struct {
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

func (r *InstallmentPlanStatementInstallment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r installmentPlanStatementInstallmentJSON) RawJSON() string {
	return r.raw
}

type InstallmentPlanStatementInstallmentsPayment struct {
	// Amount applied to the installment in cents
	Amount int64 `json:"amount" api:"required"`
	// Date the payment was applied to the installment
	Date time.Time                                       `json:"date" api:"required" format:"date"`
	JSON installmentPlanStatementInstallmentsPaymentJSON `json:"-"`
}

// installmentPlanStatementInstallmentsPaymentJSON contains the JSON metadata for
// the struct [InstallmentPlanStatementInstallmentsPayment]
type installmentPlanStatementInstallmentsPaymentJSON struct {
	Amount      apijson.Field
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstallmentPlanStatementInstallmentsPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r installmentPlanStatementInstallmentsPaymentJSON) RawJSON() string {
	return r.raw
}

type InstallmentPlanStatementSourceType string

const (
	InstallmentPlanStatementSourceTypeUnpaidBalance InstallmentPlanStatementSourceType = "UNPAID_BALANCE"
	InstallmentPlanStatementSourceTypeTransaction   InstallmentPlanStatementSourceType = "TRANSACTION"
)

func (r InstallmentPlanStatementSourceType) IsKnown() bool {
	switch r {
	case InstallmentPlanStatementSourceTypeUnpaidBalance, InstallmentPlanStatementSourceTypeTransaction:
		return true
	}
	return false
}

// State of the installment plan. A plan is REBUILD_IN_PROGRESS while its loan
// tapes are being rebuilt, during which its payment totals are being recomputed
// and should not be treated as final
type InstallmentPlanStatementState string

const (
	InstallmentPlanStatementStatePending           InstallmentPlanStatementState = "PENDING"
	InstallmentPlanStatementStateActive            InstallmentPlanStatementState = "ACTIVE"
	InstallmentPlanStatementStateRebuildInProgress InstallmentPlanStatementState = "REBUILD_IN_PROGRESS"
	InstallmentPlanStatementStateFullyPaid         InstallmentPlanStatementState = "FULLY_PAID"
	InstallmentPlanStatementStateCancelled         InstallmentPlanStatementState = "CANCELLED"
)

func (r InstallmentPlanStatementState) IsKnown() bool {
	switch r {
	case InstallmentPlanStatementStatePending, InstallmentPlanStatementStateActive, InstallmentPlanStatementStateRebuildInProgress, InstallmentPlanStatementStateFullyPaid, InstallmentPlanStatementStateCancelled:
		return true
	}
	return false
}

type FinancialAccountInstallmentPlanStatementListParams struct {
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
	// Only snapshots in which the plan was in this state will be included.
	State param.Field[FinancialAccountInstallmentPlanStatementListParamsState] `query:"state"`
}

// URLQuery serializes [FinancialAccountInstallmentPlanStatementListParams]'s query
// parameters as `url.Values`.
func (r FinancialAccountInstallmentPlanStatementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only snapshots in which the plan was in this state will be included.
type FinancialAccountInstallmentPlanStatementListParamsState string

const (
	FinancialAccountInstallmentPlanStatementListParamsStatePending           FinancialAccountInstallmentPlanStatementListParamsState = "PENDING"
	FinancialAccountInstallmentPlanStatementListParamsStateActive            FinancialAccountInstallmentPlanStatementListParamsState = "ACTIVE"
	FinancialAccountInstallmentPlanStatementListParamsStateRebuildInProgress FinancialAccountInstallmentPlanStatementListParamsState = "REBUILD_IN_PROGRESS"
	FinancialAccountInstallmentPlanStatementListParamsStateFullyPaid         FinancialAccountInstallmentPlanStatementListParamsState = "FULLY_PAID"
	FinancialAccountInstallmentPlanStatementListParamsStateCancelled         FinancialAccountInstallmentPlanStatementListParamsState = "CANCELLED"
)

func (r FinancialAccountInstallmentPlanStatementListParamsState) IsKnown() bool {
	switch r {
	case FinancialAccountInstallmentPlanStatementListParamsStatePending, FinancialAccountInstallmentPlanStatementListParamsStateActive, FinancialAccountInstallmentPlanStatementListParamsStateRebuildInProgress, FinancialAccountInstallmentPlanStatementListParamsStateFullyPaid, FinancialAccountInstallmentPlanStatementListParamsStateCancelled:
		return true
	}
	return false
}
