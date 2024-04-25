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
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// PaymentService contains methods and other services that help with interacting
// with the lithic API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPaymentService] method instead.
type PaymentService struct {
	Options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r *PaymentService) {
	r = &PaymentService{}
	r.Options = opts
	return
}

// Initiates a payment between a financial account and an external bank account.
func (r *PaymentService) New(ctx context.Context, body PaymentNewParams, opts ...option.RequestOption) (res *PaymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the payment by token.
func (r *PaymentService) Get(ctx context.Context, paymentToken string, opts ...option.RequestOption) (res *Payment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("payments/%s", paymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all the payments for the provided search criteria.
func (r *PaymentService) List(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Payment], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "payments"
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

// List all the payments for the provided search criteria.
func (r *PaymentService) ListAutoPaging(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Payment] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Retry an origination which has been returned.
func (r *PaymentService) Retry(ctx context.Context, paymentToken string, opts ...option.RequestOption) (res *PaymentRetryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("payments/%s/retry", paymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates a release of a Payment.
func (r *PaymentService) SimulateRelease(ctx context.Context, body PaymentSimulateReleaseParams, opts ...option.RequestOption) (res *PaymentSimulateReleaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/payments/release"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates a return of a Payment.
func (r *PaymentService) SimulateReturn(ctx context.Context, body PaymentSimulateReturnParams, opts ...option.RequestOption) (res *PaymentSimulateReturnResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/payments/return"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Payment struct {
	Direction                PaymentDirection        `json:"direction,required"`
	ExternalBankAccountToken string                  `json:"external_bank_account_token,required,nullable" format:"uuid"`
	FinancialAccountToken    string                  `json:"financial_account_token,required" format:"uuid"`
	Method                   PaymentMethod           `json:"method,required"`
	MethodAttributes         PaymentMethodAttributes `json:"method_attributes,required"`
	Source                   PaymentSource           `json:"source,required"`
	UserDefinedID            string                  `json:"user_defined_id,required,nullable"`
	JSON                     paymentJSON             `json:"-"`
	FinancialTransaction
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	Direction                apijson.Field
	ExternalBankAccountToken apijson.Field
	FinancialAccountToken    apijson.Field
	Method                   apijson.Field
	MethodAttributes         apijson.Field
	Source                   apijson.Field
	UserDefinedID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Payment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentJSON) RawJSON() string {
	return r.raw
}

type PaymentDirection string

const (
	PaymentDirectionCredit PaymentDirection = "CREDIT"
	PaymentDirectionDebit  PaymentDirection = "DEBIT"
)

func (r PaymentDirection) IsKnown() bool {
	switch r {
	case PaymentDirectionCredit, PaymentDirectionDebit:
		return true
	}
	return false
}

type PaymentMethod string

const (
	PaymentMethodACHNextDay PaymentMethod = "ACH_NEXT_DAY"
	PaymentMethodACHSameDay PaymentMethod = "ACH_SAME_DAY"
)

func (r PaymentMethod) IsKnown() bool {
	switch r {
	case PaymentMethodACHNextDay, PaymentMethodACHSameDay:
		return true
	}
	return false
}

type PaymentMethodAttributes struct {
	CompanyID            string                         `json:"company_id,required,nullable"`
	ReceiptRoutingNumber string                         `json:"receipt_routing_number,required,nullable"`
	Retries              int64                          `json:"retries,required,nullable"`
	ReturnReasonCode     string                         `json:"return_reason_code,required,nullable"`
	SecCode              PaymentMethodAttributesSecCode `json:"sec_code,required"`
	JSON                 paymentMethodAttributesJSON    `json:"-"`
}

// paymentMethodAttributesJSON contains the JSON metadata for the struct
// [PaymentMethodAttributes]
type paymentMethodAttributesJSON struct {
	CompanyID            apijson.Field
	ReceiptRoutingNumber apijson.Field
	Retries              apijson.Field
	ReturnReasonCode     apijson.Field
	SecCode              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PaymentMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodAttributesJSON) RawJSON() string {
	return r.raw
}

type PaymentMethodAttributesSecCode string

const (
	PaymentMethodAttributesSecCodeCcd PaymentMethodAttributesSecCode = "CCD"
	PaymentMethodAttributesSecCodePpd PaymentMethodAttributesSecCode = "PPD"
	PaymentMethodAttributesSecCodeWeb PaymentMethodAttributesSecCode = "WEB"
)

func (r PaymentMethodAttributesSecCode) IsKnown() bool {
	switch r {
	case PaymentMethodAttributesSecCodeCcd, PaymentMethodAttributesSecCodePpd, PaymentMethodAttributesSecCodeWeb:
		return true
	}
	return false
}

type PaymentSource string

const (
	PaymentSourceCustomer PaymentSource = "CUSTOMER"
	PaymentSourceLithic   PaymentSource = "LITHIC"
)

func (r PaymentSource) IsKnown() bool {
	switch r {
	case PaymentSourceCustomer, PaymentSourceLithic:
		return true
	}
	return false
}

type PaymentNewResponse struct {
	// Balance
	Balance Balance                `json:"balance"`
	JSON    paymentNewResponseJSON `json:"-"`
	Payment
}

// paymentNewResponseJSON contains the JSON metadata for the struct
// [PaymentNewResponse]
type paymentNewResponseJSON struct {
	Balance     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentNewResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentRetryResponse struct {
	// Balance
	Balance Balance                  `json:"balance"`
	JSON    paymentRetryResponseJSON `json:"-"`
	Payment
}

// paymentRetryResponseJSON contains the JSON metadata for the struct
// [PaymentRetryResponse]
type paymentRetryResponseJSON struct {
	Balance     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentRetryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentRetryResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentSimulateReleaseResponse struct {
	DebuggingRequestID    string                               `json:"debugging_request_id" format:"uuid"`
	Result                PaymentSimulateReleaseResponseResult `json:"result"`
	TransactionEventToken string                               `json:"transaction_event_token" format:"uuid"`
	JSON                  paymentSimulateReleaseResponseJSON   `json:"-"`
}

// paymentSimulateReleaseResponseJSON contains the JSON metadata for the struct
// [PaymentSimulateReleaseResponse]
type paymentSimulateReleaseResponseJSON struct {
	DebuggingRequestID    apijson.Field
	Result                apijson.Field
	TransactionEventToken apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PaymentSimulateReleaseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSimulateReleaseResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentSimulateReleaseResponseResult string

const (
	PaymentSimulateReleaseResponseResultApproved PaymentSimulateReleaseResponseResult = "APPROVED"
	PaymentSimulateReleaseResponseResultDeclined PaymentSimulateReleaseResponseResult = "DECLINED"
)

func (r PaymentSimulateReleaseResponseResult) IsKnown() bool {
	switch r {
	case PaymentSimulateReleaseResponseResultApproved, PaymentSimulateReleaseResponseResultDeclined:
		return true
	}
	return false
}

type PaymentSimulateReturnResponse struct {
	DebuggingRequestID    string                              `json:"debugging_request_id" format:"uuid"`
	Result                PaymentSimulateReturnResponseResult `json:"result"`
	TransactionEventToken string                              `json:"transaction_event_token" format:"uuid"`
	JSON                  paymentSimulateReturnResponseJSON   `json:"-"`
}

// paymentSimulateReturnResponseJSON contains the JSON metadata for the struct
// [PaymentSimulateReturnResponse]
type paymentSimulateReturnResponseJSON struct {
	DebuggingRequestID    apijson.Field
	Result                apijson.Field
	TransactionEventToken apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PaymentSimulateReturnResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSimulateReturnResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentSimulateReturnResponseResult string

const (
	PaymentSimulateReturnResponseResultApproved PaymentSimulateReturnResponseResult = "APPROVED"
	PaymentSimulateReturnResponseResultDeclined PaymentSimulateReturnResponseResult = "DECLINED"
)

func (r PaymentSimulateReturnResponseResult) IsKnown() bool {
	switch r {
	case PaymentSimulateReturnResponseResultApproved, PaymentSimulateReturnResponseResultDeclined:
		return true
	}
	return false
}

type PaymentNewParams struct {
	Amount                   param.Field[int64]                            `json:"amount,required"`
	ExternalBankAccountToken param.Field[string]                           `json:"external_bank_account_token,required" format:"uuid"`
	FinancialAccountToken    param.Field[string]                           `json:"financial_account_token,required" format:"uuid"`
	Method                   param.Field[PaymentNewParamsMethod]           `json:"method,required"`
	MethodAttributes         param.Field[PaymentNewParamsMethodAttributes] `json:"method_attributes,required"`
	Type                     param.Field[PaymentNewParamsType]             `json:"type,required"`
	// Customer-provided token that will serve as an idempotency token. This token will
	// become the transaction token.
	Token         param.Field[string] `json:"token" format:"uuid"`
	Memo          param.Field[string] `json:"memo"`
	UserDefinedID param.Field[string] `json:"user_defined_id"`
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsMethod string

const (
	PaymentNewParamsMethodACHNextDay PaymentNewParamsMethod = "ACH_NEXT_DAY"
	PaymentNewParamsMethodACHSameDay PaymentNewParamsMethod = "ACH_SAME_DAY"
)

func (r PaymentNewParamsMethod) IsKnown() bool {
	switch r {
	case PaymentNewParamsMethodACHNextDay, PaymentNewParamsMethodACHSameDay:
		return true
	}
	return false
}

type PaymentNewParamsMethodAttributes struct {
	CompanyID            param.Field[string]                                  `json:"company_id,required"`
	ReceiptRoutingNumber param.Field[string]                                  `json:"receipt_routing_number,required"`
	Retries              param.Field[int64]                                   `json:"retries,required"`
	ReturnReasonCode     param.Field[string]                                  `json:"return_reason_code,required"`
	SecCode              param.Field[PaymentNewParamsMethodAttributesSecCode] `json:"sec_code,required"`
}

func (r PaymentNewParamsMethodAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsMethodAttributesSecCode string

const (
	PaymentNewParamsMethodAttributesSecCodeCcd PaymentNewParamsMethodAttributesSecCode = "CCD"
	PaymentNewParamsMethodAttributesSecCodePpd PaymentNewParamsMethodAttributesSecCode = "PPD"
	PaymentNewParamsMethodAttributesSecCodeWeb PaymentNewParamsMethodAttributesSecCode = "WEB"
)

func (r PaymentNewParamsMethodAttributesSecCode) IsKnown() bool {
	switch r {
	case PaymentNewParamsMethodAttributesSecCodeCcd, PaymentNewParamsMethodAttributesSecCodePpd, PaymentNewParamsMethodAttributesSecCodeWeb:
		return true
	}
	return false
}

type PaymentNewParamsType string

const (
	PaymentNewParamsTypeCollection PaymentNewParamsType = "COLLECTION"
	PaymentNewParamsTypePayment    PaymentNewParamsType = "PAYMENT"
)

func (r PaymentNewParamsType) IsKnown() bool {
	switch r {
	case PaymentNewParamsTypeCollection, PaymentNewParamsTypePayment:
		return true
	}
	return false
}

type PaymentListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore          param.Field[string] `query:"ending_before"`
	FinancialAccountToken param.Field[string] `query:"financial_account_token" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64]                   `query:"page_size"`
	Result   param.Field[PaymentListParamsResult] `query:"result"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string]                  `query:"starting_after"`
	Status        param.Field[PaymentListParamsStatus] `query:"status"`
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaymentListParamsResult string

const (
	PaymentListParamsResultApproved PaymentListParamsResult = "APPROVED"
	PaymentListParamsResultDeclined PaymentListParamsResult = "DECLINED"
)

func (r PaymentListParamsResult) IsKnown() bool {
	switch r {
	case PaymentListParamsResultApproved, PaymentListParamsResultDeclined:
		return true
	}
	return false
}

type PaymentListParamsStatus string

const (
	PaymentListParamsStatusDeclined PaymentListParamsStatus = "DECLINED"
	PaymentListParamsStatusPending  PaymentListParamsStatus = "PENDING"
	PaymentListParamsStatusReturned PaymentListParamsStatus = "RETURNED"
	PaymentListParamsStatusSettled  PaymentListParamsStatus = "SETTLED"
)

func (r PaymentListParamsStatus) IsKnown() bool {
	switch r {
	case PaymentListParamsStatusDeclined, PaymentListParamsStatusPending, PaymentListParamsStatusReturned, PaymentListParamsStatusSettled:
		return true
	}
	return false
}

type PaymentSimulateReleaseParams struct {
	PaymentToken param.Field[string] `json:"payment_token,required" format:"uuid"`
}

func (r PaymentSimulateReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentSimulateReturnParams struct {
	PaymentToken     param.Field[string] `json:"payment_token,required" format:"uuid"`
	ReturnReasonCode param.Field[string] `json:"return_reason_code"`
}

func (r PaymentSimulateReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
