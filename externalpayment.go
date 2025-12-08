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

// ExternalPaymentService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalPaymentService] method instead.
type ExternalPaymentService struct {
	Options []option.RequestOption
}

// NewExternalPaymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExternalPaymentService(opts ...option.RequestOption) (r *ExternalPaymentService) {
	r = &ExternalPaymentService{}
	r.Options = opts
	return
}

// Create external payment
func (r *ExternalPaymentService) New(ctx context.Context, body ExternalPaymentNewParams, opts ...option.RequestOption) (res *ExternalPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/external_payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get external payment
func (r *ExternalPaymentService) Get(ctx context.Context, externalPaymentToken string, opts ...option.RequestOption) (res *ExternalPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPaymentToken == "" {
		err = errors.New("missing required external_payment_token parameter")
		return
	}
	path := fmt.Sprintf("v1/external_payments/%s", externalPaymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List external payments
func (r *ExternalPaymentService) List(ctx context.Context, query ExternalPaymentListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ExternalPayment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/external_payments"
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

// List external payments
func (r *ExternalPaymentService) ListAutoPaging(ctx context.Context, query ExternalPaymentListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ExternalPayment] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Cancel external payment
func (r *ExternalPaymentService) Cancel(ctx context.Context, externalPaymentToken string, body ExternalPaymentCancelParams, opts ...option.RequestOption) (res *ExternalPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPaymentToken == "" {
		err = errors.New("missing required external_payment_token parameter")
		return
	}
	path := fmt.Sprintf("v1/external_payments/%s/cancel", externalPaymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Release external payment
func (r *ExternalPaymentService) Release(ctx context.Context, externalPaymentToken string, body ExternalPaymentReleaseParams, opts ...option.RequestOption) (res *ExternalPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPaymentToken == "" {
		err = errors.New("missing required external_payment_token parameter")
		return
	}
	path := fmt.Sprintf("v1/external_payments/%s/release", externalPaymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reverse external payment
func (r *ExternalPaymentService) Reverse(ctx context.Context, externalPaymentToken string, body ExternalPaymentReverseParams, opts ...option.RequestOption) (res *ExternalPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPaymentToken == "" {
		err = errors.New("missing required external_payment_token parameter")
		return
	}
	path := fmt.Sprintf("v1/external_payments/%s/reverse", externalPaymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Settle external payment
func (r *ExternalPaymentService) Settle(ctx context.Context, externalPaymentToken string, body ExternalPaymentSettleParams, opts ...option.RequestOption) (res *ExternalPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPaymentToken == "" {
		err = errors.New("missing required external_payment_token parameter")
		return
	}
	path := fmt.Sprintf("v1/external_payments/%s/settle", externalPaymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExternalPayment struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// The status of the transaction
	Status ExternalPaymentStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated  time.Time               `json:"updated,required" format:"date-time"`
	Category ExternalPaymentCategory `json:"category"`
	Currency string                  `json:"currency"`
	Events   []ExternalPaymentEvent  `json:"events"`
	// EXTERNAL_PAYMENT - External Payment Response
	Family                ExternalPaymentFamily      `json:"family"`
	FinancialAccountToken string                     `json:"financial_account_token" format:"uuid"`
	PaymentType           ExternalPaymentPaymentType `json:"payment_type"`
	PendingAmount         int64                      `json:"pending_amount"`
	Result                ExternalPaymentResult      `json:"result"`
	SettledAmount         int64                      `json:"settled_amount"`
	UserDefinedID         string                     `json:"user_defined_id,nullable"`
	JSON                  externalPaymentJSON        `json:"-"`
}

// externalPaymentJSON contains the JSON metadata for the struct [ExternalPayment]
type externalPaymentJSON struct {
	Token                 apijson.Field
	Created               apijson.Field
	Status                apijson.Field
	Updated               apijson.Field
	Category              apijson.Field
	Currency              apijson.Field
	Events                apijson.Field
	Family                apijson.Field
	FinancialAccountToken apijson.Field
	PaymentType           apijson.Field
	PendingAmount         apijson.Field
	Result                apijson.Field
	SettledAmount         apijson.Field
	UserDefinedID         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ExternalPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalPaymentJSON) RawJSON() string {
	return r.raw
}

func (r ExternalPayment) implementsAccountActivityListResponse() {}

func (r ExternalPayment) implementsAccountActivityGetTransactionResponse() {}

// The status of the transaction
type ExternalPaymentStatus string

const (
	ExternalPaymentStatusPending  ExternalPaymentStatus = "PENDING"
	ExternalPaymentStatusSettled  ExternalPaymentStatus = "SETTLED"
	ExternalPaymentStatusDeclined ExternalPaymentStatus = "DECLINED"
	ExternalPaymentStatusReversed ExternalPaymentStatus = "REVERSED"
	ExternalPaymentStatusCanceled ExternalPaymentStatus = "CANCELED"
	ExternalPaymentStatusReturned ExternalPaymentStatus = "RETURNED"
)

func (r ExternalPaymentStatus) IsKnown() bool {
	switch r {
	case ExternalPaymentStatusPending, ExternalPaymentStatusSettled, ExternalPaymentStatusDeclined, ExternalPaymentStatusReversed, ExternalPaymentStatusCanceled, ExternalPaymentStatusReturned:
		return true
	}
	return false
}

type ExternalPaymentCategory string

const (
	ExternalPaymentCategoryExternalWire     ExternalPaymentCategory = "EXTERNAL_WIRE"
	ExternalPaymentCategoryExternalACH      ExternalPaymentCategory = "EXTERNAL_ACH"
	ExternalPaymentCategoryExternalCheck    ExternalPaymentCategory = "EXTERNAL_CHECK"
	ExternalPaymentCategoryExternalFednow   ExternalPaymentCategory = "EXTERNAL_FEDNOW"
	ExternalPaymentCategoryExternalRtp      ExternalPaymentCategory = "EXTERNAL_RTP"
	ExternalPaymentCategoryExternalTransfer ExternalPaymentCategory = "EXTERNAL_TRANSFER"
)

func (r ExternalPaymentCategory) IsKnown() bool {
	switch r {
	case ExternalPaymentCategoryExternalWire, ExternalPaymentCategoryExternalACH, ExternalPaymentCategoryExternalCheck, ExternalPaymentCategoryExternalFednow, ExternalPaymentCategoryExternalRtp, ExternalPaymentCategoryExternalTransfer:
		return true
	}
	return false
}

type ExternalPaymentEvent struct {
	Token           string                                `json:"token,required" format:"uuid"`
	Amount          int64                                 `json:"amount,required"`
	Created         time.Time                             `json:"created,required" format:"date-time"`
	DetailedResults []ExternalPaymentEventsDetailedResult `json:"detailed_results,required"`
	EffectiveDate   time.Time                             `json:"effective_date,required" format:"date"`
	Memo            string                                `json:"memo,required"`
	Result          ExternalPaymentEventsResult           `json:"result,required"`
	Type            ExternalPaymentEventsType             `json:"type,required"`
	JSON            externalPaymentEventJSON              `json:"-"`
}

// externalPaymentEventJSON contains the JSON metadata for the struct
// [ExternalPaymentEvent]
type externalPaymentEventJSON struct {
	Token           apijson.Field
	Amount          apijson.Field
	Created         apijson.Field
	DetailedResults apijson.Field
	EffectiveDate   apijson.Field
	Memo            apijson.Field
	Result          apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ExternalPaymentEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalPaymentEventJSON) RawJSON() string {
	return r.raw
}

type ExternalPaymentEventsDetailedResult string

const (
	ExternalPaymentEventsDetailedResultApproved          ExternalPaymentEventsDetailedResult = "APPROVED"
	ExternalPaymentEventsDetailedResultInsufficientFunds ExternalPaymentEventsDetailedResult = "INSUFFICIENT_FUNDS"
)

func (r ExternalPaymentEventsDetailedResult) IsKnown() bool {
	switch r {
	case ExternalPaymentEventsDetailedResultApproved, ExternalPaymentEventsDetailedResultInsufficientFunds:
		return true
	}
	return false
}

type ExternalPaymentEventsResult string

const (
	ExternalPaymentEventsResultApproved ExternalPaymentEventsResult = "APPROVED"
	ExternalPaymentEventsResultDeclined ExternalPaymentEventsResult = "DECLINED"
)

func (r ExternalPaymentEventsResult) IsKnown() bool {
	switch r {
	case ExternalPaymentEventsResultApproved, ExternalPaymentEventsResultDeclined:
		return true
	}
	return false
}

type ExternalPaymentEventsType string

const (
	ExternalPaymentEventsTypeExternalWireInitiated     ExternalPaymentEventsType = "EXTERNAL_WIRE_INITIATED"
	ExternalPaymentEventsTypeExternalWireCanceled      ExternalPaymentEventsType = "EXTERNAL_WIRE_CANCELED"
	ExternalPaymentEventsTypeExternalWireSettled       ExternalPaymentEventsType = "EXTERNAL_WIRE_SETTLED"
	ExternalPaymentEventsTypeExternalWireReversed      ExternalPaymentEventsType = "EXTERNAL_WIRE_REVERSED"
	ExternalPaymentEventsTypeExternalWireReleased      ExternalPaymentEventsType = "EXTERNAL_WIRE_RELEASED"
	ExternalPaymentEventsTypeExternalACHInitiated      ExternalPaymentEventsType = "EXTERNAL_ACH_INITIATED"
	ExternalPaymentEventsTypeExternalACHCanceled       ExternalPaymentEventsType = "EXTERNAL_ACH_CANCELED"
	ExternalPaymentEventsTypeExternalACHSettled        ExternalPaymentEventsType = "EXTERNAL_ACH_SETTLED"
	ExternalPaymentEventsTypeExternalACHReversed       ExternalPaymentEventsType = "EXTERNAL_ACH_REVERSED"
	ExternalPaymentEventsTypeExternalACHReleased       ExternalPaymentEventsType = "EXTERNAL_ACH_RELEASED"
	ExternalPaymentEventsTypeExternalTransferInitiated ExternalPaymentEventsType = "EXTERNAL_TRANSFER_INITIATED"
	ExternalPaymentEventsTypeExternalTransferCanceled  ExternalPaymentEventsType = "EXTERNAL_TRANSFER_CANCELED"
	ExternalPaymentEventsTypeExternalTransferSettled   ExternalPaymentEventsType = "EXTERNAL_TRANSFER_SETTLED"
	ExternalPaymentEventsTypeExternalTransferReversed  ExternalPaymentEventsType = "EXTERNAL_TRANSFER_REVERSED"
	ExternalPaymentEventsTypeExternalTransferReleased  ExternalPaymentEventsType = "EXTERNAL_TRANSFER_RELEASED"
	ExternalPaymentEventsTypeExternalCheckInitiated    ExternalPaymentEventsType = "EXTERNAL_CHECK_INITIATED"
	ExternalPaymentEventsTypeExternalCheckCanceled     ExternalPaymentEventsType = "EXTERNAL_CHECK_CANCELED"
	ExternalPaymentEventsTypeExternalCheckSettled      ExternalPaymentEventsType = "EXTERNAL_CHECK_SETTLED"
	ExternalPaymentEventsTypeExternalCheckReversed     ExternalPaymentEventsType = "EXTERNAL_CHECK_REVERSED"
	ExternalPaymentEventsTypeExternalCheckReleased     ExternalPaymentEventsType = "EXTERNAL_CHECK_RELEASED"
	ExternalPaymentEventsTypeExternalFednowInitiated   ExternalPaymentEventsType = "EXTERNAL_FEDNOW_INITIATED"
	ExternalPaymentEventsTypeExternalFednowCanceled    ExternalPaymentEventsType = "EXTERNAL_FEDNOW_CANCELED"
	ExternalPaymentEventsTypeExternalFednowSettled     ExternalPaymentEventsType = "EXTERNAL_FEDNOW_SETTLED"
	ExternalPaymentEventsTypeExternalFednowReversed    ExternalPaymentEventsType = "EXTERNAL_FEDNOW_REVERSED"
	ExternalPaymentEventsTypeExternalFednowReleased    ExternalPaymentEventsType = "EXTERNAL_FEDNOW_RELEASED"
	ExternalPaymentEventsTypeExternalRtpInitiated      ExternalPaymentEventsType = "EXTERNAL_RTP_INITIATED"
	ExternalPaymentEventsTypeExternalRtpCanceled       ExternalPaymentEventsType = "EXTERNAL_RTP_CANCELED"
	ExternalPaymentEventsTypeExternalRtpSettled        ExternalPaymentEventsType = "EXTERNAL_RTP_SETTLED"
	ExternalPaymentEventsTypeExternalRtpReversed       ExternalPaymentEventsType = "EXTERNAL_RTP_REVERSED"
	ExternalPaymentEventsTypeExternalRtpReleased       ExternalPaymentEventsType = "EXTERNAL_RTP_RELEASED"
)

func (r ExternalPaymentEventsType) IsKnown() bool {
	switch r {
	case ExternalPaymentEventsTypeExternalWireInitiated, ExternalPaymentEventsTypeExternalWireCanceled, ExternalPaymentEventsTypeExternalWireSettled, ExternalPaymentEventsTypeExternalWireReversed, ExternalPaymentEventsTypeExternalWireReleased, ExternalPaymentEventsTypeExternalACHInitiated, ExternalPaymentEventsTypeExternalACHCanceled, ExternalPaymentEventsTypeExternalACHSettled, ExternalPaymentEventsTypeExternalACHReversed, ExternalPaymentEventsTypeExternalACHReleased, ExternalPaymentEventsTypeExternalTransferInitiated, ExternalPaymentEventsTypeExternalTransferCanceled, ExternalPaymentEventsTypeExternalTransferSettled, ExternalPaymentEventsTypeExternalTransferReversed, ExternalPaymentEventsTypeExternalTransferReleased, ExternalPaymentEventsTypeExternalCheckInitiated, ExternalPaymentEventsTypeExternalCheckCanceled, ExternalPaymentEventsTypeExternalCheckSettled, ExternalPaymentEventsTypeExternalCheckReversed, ExternalPaymentEventsTypeExternalCheckReleased, ExternalPaymentEventsTypeExternalFednowInitiated, ExternalPaymentEventsTypeExternalFednowCanceled, ExternalPaymentEventsTypeExternalFednowSettled, ExternalPaymentEventsTypeExternalFednowReversed, ExternalPaymentEventsTypeExternalFednowReleased, ExternalPaymentEventsTypeExternalRtpInitiated, ExternalPaymentEventsTypeExternalRtpCanceled, ExternalPaymentEventsTypeExternalRtpSettled, ExternalPaymentEventsTypeExternalRtpReversed, ExternalPaymentEventsTypeExternalRtpReleased:
		return true
	}
	return false
}

// EXTERNAL_PAYMENT - External Payment Response
type ExternalPaymentFamily string

const (
	ExternalPaymentFamilyExternalPayment ExternalPaymentFamily = "EXTERNAL_PAYMENT"
)

func (r ExternalPaymentFamily) IsKnown() bool {
	switch r {
	case ExternalPaymentFamilyExternalPayment:
		return true
	}
	return false
}

type ExternalPaymentPaymentType string

const (
	ExternalPaymentPaymentTypeDeposit    ExternalPaymentPaymentType = "DEPOSIT"
	ExternalPaymentPaymentTypeWithdrawal ExternalPaymentPaymentType = "WITHDRAWAL"
)

func (r ExternalPaymentPaymentType) IsKnown() bool {
	switch r {
	case ExternalPaymentPaymentTypeDeposit, ExternalPaymentPaymentTypeWithdrawal:
		return true
	}
	return false
}

type ExternalPaymentResult string

const (
	ExternalPaymentResultApproved ExternalPaymentResult = "APPROVED"
	ExternalPaymentResultDeclined ExternalPaymentResult = "DECLINED"
)

func (r ExternalPaymentResult) IsKnown() bool {
	switch r {
	case ExternalPaymentResultApproved, ExternalPaymentResultDeclined:
		return true
	}
	return false
}

type ExternalPaymentNewParams struct {
	Amount                param.Field[int64]                               `json:"amount,required"`
	Category              param.Field[ExternalPaymentNewParamsCategory]    `json:"category,required"`
	EffectiveDate         param.Field[time.Time]                           `json:"effective_date,required" format:"date"`
	FinancialAccountToken param.Field[string]                              `json:"financial_account_token,required" format:"uuid"`
	PaymentType           param.Field[ExternalPaymentNewParamsPaymentType] `json:"payment_type,required"`
	Token                 param.Field[string]                              `json:"token" format:"uuid"`
	Memo                  param.Field[string]                              `json:"memo"`
	ProgressTo            param.Field[ExternalPaymentNewParamsProgressTo]  `json:"progress_to"`
	UserDefinedID         param.Field[string]                              `json:"user_defined_id"`
}

func (r ExternalPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalPaymentNewParamsCategory string

const (
	ExternalPaymentNewParamsCategoryExternalWire     ExternalPaymentNewParamsCategory = "EXTERNAL_WIRE"
	ExternalPaymentNewParamsCategoryExternalACH      ExternalPaymentNewParamsCategory = "EXTERNAL_ACH"
	ExternalPaymentNewParamsCategoryExternalCheck    ExternalPaymentNewParamsCategory = "EXTERNAL_CHECK"
	ExternalPaymentNewParamsCategoryExternalFednow   ExternalPaymentNewParamsCategory = "EXTERNAL_FEDNOW"
	ExternalPaymentNewParamsCategoryExternalRtp      ExternalPaymentNewParamsCategory = "EXTERNAL_RTP"
	ExternalPaymentNewParamsCategoryExternalTransfer ExternalPaymentNewParamsCategory = "EXTERNAL_TRANSFER"
)

func (r ExternalPaymentNewParamsCategory) IsKnown() bool {
	switch r {
	case ExternalPaymentNewParamsCategoryExternalWire, ExternalPaymentNewParamsCategoryExternalACH, ExternalPaymentNewParamsCategoryExternalCheck, ExternalPaymentNewParamsCategoryExternalFednow, ExternalPaymentNewParamsCategoryExternalRtp, ExternalPaymentNewParamsCategoryExternalTransfer:
		return true
	}
	return false
}

type ExternalPaymentNewParamsPaymentType string

const (
	ExternalPaymentNewParamsPaymentTypeDeposit    ExternalPaymentNewParamsPaymentType = "DEPOSIT"
	ExternalPaymentNewParamsPaymentTypeWithdrawal ExternalPaymentNewParamsPaymentType = "WITHDRAWAL"
)

func (r ExternalPaymentNewParamsPaymentType) IsKnown() bool {
	switch r {
	case ExternalPaymentNewParamsPaymentTypeDeposit, ExternalPaymentNewParamsPaymentTypeWithdrawal:
		return true
	}
	return false
}

type ExternalPaymentNewParamsProgressTo string

const (
	ExternalPaymentNewParamsProgressToSettled  ExternalPaymentNewParamsProgressTo = "SETTLED"
	ExternalPaymentNewParamsProgressToReleased ExternalPaymentNewParamsProgressTo = "RELEASED"
)

func (r ExternalPaymentNewParamsProgressTo) IsKnown() bool {
	switch r {
	case ExternalPaymentNewParamsProgressToSettled, ExternalPaymentNewParamsProgressToReleased:
		return true
	}
	return false
}

type ExternalPaymentListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin                param.Field[time.Time] `query:"begin" format:"date-time"`
	BusinessAccountToken param.Field[string]    `query:"business_account_token" format:"uuid"`
	// External Payment category to be returned.
	Category param.Field[ExternalPaymentListParamsCategory] `query:"category"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case.
	FinancialAccountToken param.Field[string] `query:"financial_account_token" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// External Payment result to be returned.
	Result param.Field[ExternalPaymentListParamsResult] `query:"result"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Book transfer status to be returned.
	Status param.Field[ExternalPaymentListParamsStatus] `query:"status"`
}

// URLQuery serializes [ExternalPaymentListParams]'s query parameters as
// `url.Values`.
func (r ExternalPaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// External Payment category to be returned.
type ExternalPaymentListParamsCategory string

const (
	ExternalPaymentListParamsCategoryExternalWire     ExternalPaymentListParamsCategory = "EXTERNAL_WIRE"
	ExternalPaymentListParamsCategoryExternalACH      ExternalPaymentListParamsCategory = "EXTERNAL_ACH"
	ExternalPaymentListParamsCategoryExternalCheck    ExternalPaymentListParamsCategory = "EXTERNAL_CHECK"
	ExternalPaymentListParamsCategoryExternalFednow   ExternalPaymentListParamsCategory = "EXTERNAL_FEDNOW"
	ExternalPaymentListParamsCategoryExternalRtp      ExternalPaymentListParamsCategory = "EXTERNAL_RTP"
	ExternalPaymentListParamsCategoryExternalTransfer ExternalPaymentListParamsCategory = "EXTERNAL_TRANSFER"
)

func (r ExternalPaymentListParamsCategory) IsKnown() bool {
	switch r {
	case ExternalPaymentListParamsCategoryExternalWire, ExternalPaymentListParamsCategoryExternalACH, ExternalPaymentListParamsCategoryExternalCheck, ExternalPaymentListParamsCategoryExternalFednow, ExternalPaymentListParamsCategoryExternalRtp, ExternalPaymentListParamsCategoryExternalTransfer:
		return true
	}
	return false
}

// External Payment result to be returned.
type ExternalPaymentListParamsResult string

const (
	ExternalPaymentListParamsResultApproved ExternalPaymentListParamsResult = "APPROVED"
	ExternalPaymentListParamsResultDeclined ExternalPaymentListParamsResult = "DECLINED"
)

func (r ExternalPaymentListParamsResult) IsKnown() bool {
	switch r {
	case ExternalPaymentListParamsResultApproved, ExternalPaymentListParamsResultDeclined:
		return true
	}
	return false
}

// Book transfer status to be returned.
type ExternalPaymentListParamsStatus string

const (
	ExternalPaymentListParamsStatusPending  ExternalPaymentListParamsStatus = "PENDING"
	ExternalPaymentListParamsStatusSettled  ExternalPaymentListParamsStatus = "SETTLED"
	ExternalPaymentListParamsStatusDeclined ExternalPaymentListParamsStatus = "DECLINED"
	ExternalPaymentListParamsStatusReversed ExternalPaymentListParamsStatus = "REVERSED"
	ExternalPaymentListParamsStatusCanceled ExternalPaymentListParamsStatus = "CANCELED"
	ExternalPaymentListParamsStatusReturned ExternalPaymentListParamsStatus = "RETURNED"
)

func (r ExternalPaymentListParamsStatus) IsKnown() bool {
	switch r {
	case ExternalPaymentListParamsStatusPending, ExternalPaymentListParamsStatusSettled, ExternalPaymentListParamsStatusDeclined, ExternalPaymentListParamsStatusReversed, ExternalPaymentListParamsStatusCanceled, ExternalPaymentListParamsStatusReturned:
		return true
	}
	return false
}

type ExternalPaymentCancelParams struct {
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date"`
	Memo          param.Field[string]    `json:"memo"`
}

func (r ExternalPaymentCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalPaymentReleaseParams struct {
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date"`
	Memo          param.Field[string]    `json:"memo"`
}

func (r ExternalPaymentReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalPaymentReverseParams struct {
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date"`
	Memo          param.Field[string]    `json:"memo"`
}

func (r ExternalPaymentReverseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalPaymentSettleParams struct {
	EffectiveDate param.Field[time.Time]                             `json:"effective_date,required" format:"date"`
	Memo          param.Field[string]                                `json:"memo"`
	ProgressTo    param.Field[ExternalPaymentSettleParamsProgressTo] `json:"progress_to"`
}

func (r ExternalPaymentSettleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalPaymentSettleParamsProgressTo string

const (
	ExternalPaymentSettleParamsProgressToSettled  ExternalPaymentSettleParamsProgressTo = "SETTLED"
	ExternalPaymentSettleParamsProgressToReleased ExternalPaymentSettleParamsProgressTo = "RELEASED"
)

func (r ExternalPaymentSettleParamsProgressTo) IsKnown() bool {
	switch r {
	case ExternalPaymentSettleParamsProgressToSettled, ExternalPaymentSettleParamsProgressToReleased:
		return true
	}
	return false
}
