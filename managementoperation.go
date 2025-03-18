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
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// ManagementOperationService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManagementOperationService] method instead.
type ManagementOperationService struct {
	Options []option.RequestOption
}

// NewManagementOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewManagementOperationService(opts ...option.RequestOption) (r *ManagementOperationService) {
	r = &ManagementOperationService{}
	r.Options = opts
	return
}

// Create management operation
func (r *ManagementOperationService) New(ctx context.Context, body ManagementOperationNewParams, opts ...option.RequestOption) (res *ManagementOperationTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/management_operations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get management operation
func (r *ManagementOperationService) Get(ctx context.Context, managementOperationToken string, opts ...option.RequestOption) (res *ManagementOperationTransaction, err error) {
	opts = append(r.Options[:], opts...)
	if managementOperationToken == "" {
		err = errors.New("missing required management_operation_token parameter")
		return
	}
	path := fmt.Sprintf("v1/management_operations/%s", managementOperationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List management operations
func (r *ManagementOperationService) List(ctx context.Context, query ManagementOperationListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ManagementOperationTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/management_operations"
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

// List management operations
func (r *ManagementOperationService) ListAutoPaging(ctx context.Context, query ManagementOperationListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ManagementOperationTransaction] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Reverse a management operation
func (r *ManagementOperationService) Reverse(ctx context.Context, managementOperationToken string, body ManagementOperationReverseParams, opts ...option.RequestOption) (res *ManagementOperationTransaction, err error) {
	opts = append(r.Options[:], opts...)
	if managementOperationToken == "" {
		err = errors.New("missing required management_operation_token parameter")
		return
	}
	path := fmt.Sprintf("v1/management_operations/%s/reverse", managementOperationToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ManagementOperationTransaction struct {
	Token                 string                                  `json:"token,required" format:"uuid"`
	Category              ManagementOperationTransactionCategory  `json:"category,required"`
	Created               time.Time                               `json:"created,required" format:"date-time"`
	Currency              string                                  `json:"currency,required"`
	Direction             ManagementOperationTransactionDirection `json:"direction,required"`
	Events                []ManagementOperationTransactionEvent   `json:"events,required"`
	FinancialAccountToken string                                  `json:"financial_account_token,required" format:"uuid"`
	PendingAmount         int64                                   `json:"pending_amount,required"`
	Result                ManagementOperationTransactionResult    `json:"result,required"`
	SettledAmount         int64                                   `json:"settled_amount,required"`
	Status                ManagementOperationTransactionStatus    `json:"status,required"`
	Updated               time.Time                               `json:"updated,required" format:"date-time"`
	UserDefinedID         string                                  `json:"user_defined_id"`
	JSON                  managementOperationTransactionJSON      `json:"-"`
}

// managementOperationTransactionJSON contains the JSON metadata for the struct
// [ManagementOperationTransaction]
type managementOperationTransactionJSON struct {
	Token                 apijson.Field
	Category              apijson.Field
	Created               apijson.Field
	Currency              apijson.Field
	Direction             apijson.Field
	Events                apijson.Field
	FinancialAccountToken apijson.Field
	PendingAmount         apijson.Field
	Result                apijson.Field
	SettledAmount         apijson.Field
	Status                apijson.Field
	Updated               apijson.Field
	UserDefinedID         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ManagementOperationTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managementOperationTransactionJSON) RawJSON() string {
	return r.raw
}

type ManagementOperationTransactionCategory string

const (
	ManagementOperationTransactionCategoryManagementFee        ManagementOperationTransactionCategory = "MANAGEMENT_FEE"
	ManagementOperationTransactionCategoryManagementDispute    ManagementOperationTransactionCategory = "MANAGEMENT_DISPUTE"
	ManagementOperationTransactionCategoryManagementReward     ManagementOperationTransactionCategory = "MANAGEMENT_REWARD"
	ManagementOperationTransactionCategoryManagementAdjustment ManagementOperationTransactionCategory = "MANAGEMENT_ADJUSTMENT"
)

func (r ManagementOperationTransactionCategory) IsKnown() bool {
	switch r {
	case ManagementOperationTransactionCategoryManagementFee, ManagementOperationTransactionCategoryManagementDispute, ManagementOperationTransactionCategoryManagementReward, ManagementOperationTransactionCategoryManagementAdjustment:
		return true
	}
	return false
}

type ManagementOperationTransactionDirection string

const (
	ManagementOperationTransactionDirectionCredit ManagementOperationTransactionDirection = "CREDIT"
	ManagementOperationTransactionDirectionDebit  ManagementOperationTransactionDirection = "DEBIT"
)

func (r ManagementOperationTransactionDirection) IsKnown() bool {
	switch r {
	case ManagementOperationTransactionDirectionCredit, ManagementOperationTransactionDirectionDebit:
		return true
	}
	return false
}

type ManagementOperationTransactionEvent struct {
	Token           string                                               `json:"token,required" format:"uuid"`
	Amount          int64                                                `json:"amount,required"`
	Created         time.Time                                            `json:"created,required" format:"date-time"`
	DetailedResults []ManagementOperationTransactionEventsDetailedResult `json:"detailed_results,required"`
	EffectiveDate   time.Time                                            `json:"effective_date,required" format:"date"`
	Memo            string                                               `json:"memo,required"`
	Result          ManagementOperationTransactionEventsResult           `json:"result,required"`
	Type            ManagementOperationTransactionEventsType             `json:"type,required"`
	Subtype         string                                               `json:"subtype"`
	JSON            managementOperationTransactionEventJSON              `json:"-"`
}

// managementOperationTransactionEventJSON contains the JSON metadata for the
// struct [ManagementOperationTransactionEvent]
type managementOperationTransactionEventJSON struct {
	Token           apijson.Field
	Amount          apijson.Field
	Created         apijson.Field
	DetailedResults apijson.Field
	EffectiveDate   apijson.Field
	Memo            apijson.Field
	Result          apijson.Field
	Type            apijson.Field
	Subtype         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ManagementOperationTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managementOperationTransactionEventJSON) RawJSON() string {
	return r.raw
}

type ManagementOperationTransactionEventsDetailedResult string

const (
	ManagementOperationTransactionEventsDetailedResultApproved ManagementOperationTransactionEventsDetailedResult = "APPROVED"
)

func (r ManagementOperationTransactionEventsDetailedResult) IsKnown() bool {
	switch r {
	case ManagementOperationTransactionEventsDetailedResultApproved:
		return true
	}
	return false
}

type ManagementOperationTransactionEventsResult string

const (
	ManagementOperationTransactionEventsResultApproved ManagementOperationTransactionEventsResult = "APPROVED"
	ManagementOperationTransactionEventsResultDeclined ManagementOperationTransactionEventsResult = "DECLINED"
)

func (r ManagementOperationTransactionEventsResult) IsKnown() bool {
	switch r {
	case ManagementOperationTransactionEventsResultApproved, ManagementOperationTransactionEventsResultDeclined:
		return true
	}
	return false
}

type ManagementOperationTransactionEventsType string

const (
	ManagementOperationTransactionEventsTypeCashBack                   ManagementOperationTransactionEventsType = "CASH_BACK"
	ManagementOperationTransactionEventsTypeCurrencyConversion         ManagementOperationTransactionEventsType = "CURRENCY_CONVERSION"
	ManagementOperationTransactionEventsTypeInterest                   ManagementOperationTransactionEventsType = "INTEREST"
	ManagementOperationTransactionEventsTypeLatePayment                ManagementOperationTransactionEventsType = "LATE_PAYMENT"
	ManagementOperationTransactionEventsTypeBillingError               ManagementOperationTransactionEventsType = "BILLING_ERROR"
	ManagementOperationTransactionEventsTypeProvisionalCredit          ManagementOperationTransactionEventsType = "PROVISIONAL_CREDIT"
	ManagementOperationTransactionEventsTypeLossWriteOff               ManagementOperationTransactionEventsType = "LOSS_WRITE_OFF"
	ManagementOperationTransactionEventsTypeCashBackReversal           ManagementOperationTransactionEventsType = "CASH_BACK_REVERSAL"
	ManagementOperationTransactionEventsTypeCurrencyConversionReversal ManagementOperationTransactionEventsType = "CURRENCY_CONVERSION_REVERSAL"
	ManagementOperationTransactionEventsTypeInterestReversal           ManagementOperationTransactionEventsType = "INTEREST_REVERSAL"
	ManagementOperationTransactionEventsTypeLatePaymentReversal        ManagementOperationTransactionEventsType = "LATE_PAYMENT_REVERSAL"
	ManagementOperationTransactionEventsTypeBillingErrorReversal       ManagementOperationTransactionEventsType = "BILLING_ERROR_REVERSAL"
	ManagementOperationTransactionEventsTypeProvisionalCreditReversal  ManagementOperationTransactionEventsType = "PROVISIONAL_CREDIT_REVERSAL"
	ManagementOperationTransactionEventsTypeReturnedPayment            ManagementOperationTransactionEventsType = "RETURNED_PAYMENT"
	ManagementOperationTransactionEventsTypeReturnedPaymentReversal    ManagementOperationTransactionEventsType = "RETURNED_PAYMENT_REVERSAL"
)

func (r ManagementOperationTransactionEventsType) IsKnown() bool {
	switch r {
	case ManagementOperationTransactionEventsTypeCashBack, ManagementOperationTransactionEventsTypeCurrencyConversion, ManagementOperationTransactionEventsTypeInterest, ManagementOperationTransactionEventsTypeLatePayment, ManagementOperationTransactionEventsTypeBillingError, ManagementOperationTransactionEventsTypeProvisionalCredit, ManagementOperationTransactionEventsTypeLossWriteOff, ManagementOperationTransactionEventsTypeCashBackReversal, ManagementOperationTransactionEventsTypeCurrencyConversionReversal, ManagementOperationTransactionEventsTypeInterestReversal, ManagementOperationTransactionEventsTypeLatePaymentReversal, ManagementOperationTransactionEventsTypeBillingErrorReversal, ManagementOperationTransactionEventsTypeProvisionalCreditReversal, ManagementOperationTransactionEventsTypeReturnedPayment, ManagementOperationTransactionEventsTypeReturnedPaymentReversal:
		return true
	}
	return false
}

type ManagementOperationTransactionResult string

const (
	ManagementOperationTransactionResultApproved ManagementOperationTransactionResult = "APPROVED"
	ManagementOperationTransactionResultDeclined ManagementOperationTransactionResult = "DECLINED"
)

func (r ManagementOperationTransactionResult) IsKnown() bool {
	switch r {
	case ManagementOperationTransactionResultApproved, ManagementOperationTransactionResultDeclined:
		return true
	}
	return false
}

type ManagementOperationTransactionStatus string

const (
	ManagementOperationTransactionStatusPending  ManagementOperationTransactionStatus = "PENDING"
	ManagementOperationTransactionStatusSettled  ManagementOperationTransactionStatus = "SETTLED"
	ManagementOperationTransactionStatusDeclined ManagementOperationTransactionStatus = "DECLINED"
	ManagementOperationTransactionStatusReversed ManagementOperationTransactionStatus = "REVERSED"
	ManagementOperationTransactionStatusCanceled ManagementOperationTransactionStatus = "CANCELED"
)

func (r ManagementOperationTransactionStatus) IsKnown() bool {
	switch r {
	case ManagementOperationTransactionStatusPending, ManagementOperationTransactionStatusSettled, ManagementOperationTransactionStatusDeclined, ManagementOperationTransactionStatusReversed, ManagementOperationTransactionStatusCanceled:
		return true
	}
	return false
}

type ManagementOperationNewParams struct {
	Amount                param.Field[int64]                                 `json:"amount,required"`
	Category              param.Field[ManagementOperationNewParamsCategory]  `json:"category,required"`
	Direction             param.Field[ManagementOperationNewParamsDirection] `json:"direction,required"`
	EffectiveDate         param.Field[time.Time]                             `json:"effective_date,required" format:"date"`
	EventType             param.Field[ManagementOperationNewParamsEventType] `json:"event_type,required"`
	FinancialAccountToken param.Field[string]                                `json:"financial_account_token,required" format:"uuid"`
	Token                 param.Field[string]                                `json:"token" format:"uuid"`
	Memo                  param.Field[string]                                `json:"memo"`
	Subtype               param.Field[string]                                `json:"subtype"`
	UserDefinedID         param.Field[string]                                `json:"user_defined_id"`
}

func (r ManagementOperationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ManagementOperationNewParamsCategory string

const (
	ManagementOperationNewParamsCategoryManagementFee        ManagementOperationNewParamsCategory = "MANAGEMENT_FEE"
	ManagementOperationNewParamsCategoryManagementDispute    ManagementOperationNewParamsCategory = "MANAGEMENT_DISPUTE"
	ManagementOperationNewParamsCategoryManagementReward     ManagementOperationNewParamsCategory = "MANAGEMENT_REWARD"
	ManagementOperationNewParamsCategoryManagementAdjustment ManagementOperationNewParamsCategory = "MANAGEMENT_ADJUSTMENT"
)

func (r ManagementOperationNewParamsCategory) IsKnown() bool {
	switch r {
	case ManagementOperationNewParamsCategoryManagementFee, ManagementOperationNewParamsCategoryManagementDispute, ManagementOperationNewParamsCategoryManagementReward, ManagementOperationNewParamsCategoryManagementAdjustment:
		return true
	}
	return false
}

type ManagementOperationNewParamsDirection string

const (
	ManagementOperationNewParamsDirectionCredit ManagementOperationNewParamsDirection = "CREDIT"
	ManagementOperationNewParamsDirectionDebit  ManagementOperationNewParamsDirection = "DEBIT"
)

func (r ManagementOperationNewParamsDirection) IsKnown() bool {
	switch r {
	case ManagementOperationNewParamsDirectionCredit, ManagementOperationNewParamsDirectionDebit:
		return true
	}
	return false
}

type ManagementOperationNewParamsEventType string

const (
	ManagementOperationNewParamsEventTypeCashBack                   ManagementOperationNewParamsEventType = "CASH_BACK"
	ManagementOperationNewParamsEventTypeCurrencyConversion         ManagementOperationNewParamsEventType = "CURRENCY_CONVERSION"
	ManagementOperationNewParamsEventTypeInterest                   ManagementOperationNewParamsEventType = "INTEREST"
	ManagementOperationNewParamsEventTypeLatePayment                ManagementOperationNewParamsEventType = "LATE_PAYMENT"
	ManagementOperationNewParamsEventTypeBillingError               ManagementOperationNewParamsEventType = "BILLING_ERROR"
	ManagementOperationNewParamsEventTypeProvisionalCredit          ManagementOperationNewParamsEventType = "PROVISIONAL_CREDIT"
	ManagementOperationNewParamsEventTypeLossWriteOff               ManagementOperationNewParamsEventType = "LOSS_WRITE_OFF"
	ManagementOperationNewParamsEventTypeCashBackReversal           ManagementOperationNewParamsEventType = "CASH_BACK_REVERSAL"
	ManagementOperationNewParamsEventTypeCurrencyConversionReversal ManagementOperationNewParamsEventType = "CURRENCY_CONVERSION_REVERSAL"
	ManagementOperationNewParamsEventTypeInterestReversal           ManagementOperationNewParamsEventType = "INTEREST_REVERSAL"
	ManagementOperationNewParamsEventTypeLatePaymentReversal        ManagementOperationNewParamsEventType = "LATE_PAYMENT_REVERSAL"
	ManagementOperationNewParamsEventTypeBillingErrorReversal       ManagementOperationNewParamsEventType = "BILLING_ERROR_REVERSAL"
	ManagementOperationNewParamsEventTypeProvisionalCreditReversal  ManagementOperationNewParamsEventType = "PROVISIONAL_CREDIT_REVERSAL"
	ManagementOperationNewParamsEventTypeReturnedPayment            ManagementOperationNewParamsEventType = "RETURNED_PAYMENT"
	ManagementOperationNewParamsEventTypeReturnedPaymentReversal    ManagementOperationNewParamsEventType = "RETURNED_PAYMENT_REVERSAL"
)

func (r ManagementOperationNewParamsEventType) IsKnown() bool {
	switch r {
	case ManagementOperationNewParamsEventTypeCashBack, ManagementOperationNewParamsEventTypeCurrencyConversion, ManagementOperationNewParamsEventTypeInterest, ManagementOperationNewParamsEventTypeLatePayment, ManagementOperationNewParamsEventTypeBillingError, ManagementOperationNewParamsEventTypeProvisionalCredit, ManagementOperationNewParamsEventTypeLossWriteOff, ManagementOperationNewParamsEventTypeCashBackReversal, ManagementOperationNewParamsEventTypeCurrencyConversionReversal, ManagementOperationNewParamsEventTypeInterestReversal, ManagementOperationNewParamsEventTypeLatePaymentReversal, ManagementOperationNewParamsEventTypeBillingErrorReversal, ManagementOperationNewParamsEventTypeProvisionalCreditReversal, ManagementOperationNewParamsEventTypeReturnedPayment, ManagementOperationNewParamsEventTypeReturnedPaymentReversal:
		return true
	}
	return false
}

type ManagementOperationListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin                param.Field[time.Time] `query:"begin" format:"date-time"`
	BusinessAccountToken param.Field[string]    `query:"business_account_token" format:"uuid"`
	// Management operation category to be returned.
	Category param.Field[ManagementOperationListParamsCategory] `query:"category"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Globally unique identifier for the financial account. Accepted type dependent on
	// the program's use case.
	FinancialAccountToken param.Field[string] `query:"financial_account_token" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Management operation status to be returned.
	Status param.Field[ManagementOperationListParamsStatus] `query:"status"`
}

// URLQuery serializes [ManagementOperationListParams]'s query parameters as
// `url.Values`.
func (r ManagementOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Management operation category to be returned.
type ManagementOperationListParamsCategory string

const (
	ManagementOperationListParamsCategoryManagementFee        ManagementOperationListParamsCategory = "MANAGEMENT_FEE"
	ManagementOperationListParamsCategoryManagementDispute    ManagementOperationListParamsCategory = "MANAGEMENT_DISPUTE"
	ManagementOperationListParamsCategoryManagementReward     ManagementOperationListParamsCategory = "MANAGEMENT_REWARD"
	ManagementOperationListParamsCategoryManagementAdjustment ManagementOperationListParamsCategory = "MANAGEMENT_ADJUSTMENT"
)

func (r ManagementOperationListParamsCategory) IsKnown() bool {
	switch r {
	case ManagementOperationListParamsCategoryManagementFee, ManagementOperationListParamsCategoryManagementDispute, ManagementOperationListParamsCategoryManagementReward, ManagementOperationListParamsCategoryManagementAdjustment:
		return true
	}
	return false
}

// Management operation status to be returned.
type ManagementOperationListParamsStatus string

const (
	ManagementOperationListParamsStatusPending  ManagementOperationListParamsStatus = "PENDING"
	ManagementOperationListParamsStatusSettled  ManagementOperationListParamsStatus = "SETTLED"
	ManagementOperationListParamsStatusDeclined ManagementOperationListParamsStatus = "DECLINED"
	ManagementOperationListParamsStatusReversed ManagementOperationListParamsStatus = "REVERSED"
	ManagementOperationListParamsStatusCanceled ManagementOperationListParamsStatus = "CANCELED"
)

func (r ManagementOperationListParamsStatus) IsKnown() bool {
	switch r {
	case ManagementOperationListParamsStatusPending, ManagementOperationListParamsStatusSettled, ManagementOperationListParamsStatusDeclined, ManagementOperationListParamsStatusReversed, ManagementOperationListParamsStatusCanceled:
		return true
	}
	return false
}

type ManagementOperationReverseParams struct {
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date"`
	Memo          param.Field[string]    `json:"memo"`
}

func (r ManagementOperationReverseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
