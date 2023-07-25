// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
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
func (r *PaymentService) List(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) (res *shared.CursorPage[Payment], err error) {
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
func (r *PaymentService) ListAutoPaging(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[Payment] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Simulates a release of a Payment.
func (r *PaymentService) SimulateRelease(ctx context.Context, body PaymentSimulateReleaseParams, opts ...option.RequestOption) (res *PaymentSimulateReleaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/payments/release"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Payment struct {
	Direction                PaymentDirection        `json:"direction,required"`
	Method                   PaymentMethod           `json:"method,required"`
	MethodAttributes         PaymentMethodAttributes `json:"method_attributes,required"`
	Source                   PaymentSource           `json:"source,required"`
	ExternalBankAccountToken string                  `json:"external_bank_account_token" format:"uuid"`
	UserDefinedID            string                  `json:"user_defined_id"`
	JSON                     paymentJSON
	FinancialTransaction
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	Direction                apijson.Field
	Method                   apijson.Field
	MethodAttributes         apijson.Field
	Source                   apijson.Field
	ExternalBankAccountToken apijson.Field
	UserDefinedID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Payment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentDirection string

const (
	PaymentDirectionCredit PaymentDirection = "CREDIT"
	PaymentDirectionDebit  PaymentDirection = "DEBIT"
)

type PaymentMethod string

const (
	PaymentMethodACHNextDay PaymentMethod = "ACH_NEXT_DAY"
	PaymentMethodACHSameDay PaymentMethod = "ACH_SAME_DAY"
)

type PaymentMethodAttributes struct {
	SecCode PaymentMethodAttributesSecCode `json:"sec_code,required"`
	JSON    paymentMethodAttributesJSON
}

// paymentMethodAttributesJSON contains the JSON metadata for the struct
// [PaymentMethodAttributes]
type paymentMethodAttributesJSON struct {
	SecCode     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentMethodAttributesSecCode string

const (
	PaymentMethodAttributesSecCodePpd PaymentMethodAttributesSecCode = "PPD"
	PaymentMethodAttributesSecCodeCcd PaymentMethodAttributesSecCode = "CCD"
	PaymentMethodAttributesSecCodeWeb PaymentMethodAttributesSecCode = "WEB"
)

type PaymentSource string

const (
	PaymentSourceLithic   PaymentSource = "LITHIC"
	PaymentSourceCustomer PaymentSource = "CUSTOMER"
)

type PaymentNewResponse struct {
	Direction        PaymentNewResponseDirection        `json:"direction,required"`
	Method           PaymentNewResponseMethod           `json:"method,required"`
	MethodAttributes PaymentNewResponseMethodAttributes `json:"method_attributes,required"`
	Source           PaymentNewResponseSource           `json:"source,required"`
	// Balance of a Financial Account
	Balance                  Balance `json:"balance"`
	ExternalBankAccountToken string  `json:"external_bank_account_token" format:"uuid"`
	UserDefinedID            string  `json:"user_defined_id"`
	JSON                     paymentNewResponseJSON
	FinancialTransaction
}

// paymentNewResponseJSON contains the JSON metadata for the struct
// [PaymentNewResponse]
type paymentNewResponseJSON struct {
	Direction                apijson.Field
	Method                   apijson.Field
	MethodAttributes         apijson.Field
	Source                   apijson.Field
	Balance                  apijson.Field
	ExternalBankAccountToken apijson.Field
	UserDefinedID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *PaymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewResponseDirection string

const (
	PaymentNewResponseDirectionCredit PaymentNewResponseDirection = "CREDIT"
	PaymentNewResponseDirectionDebit  PaymentNewResponseDirection = "DEBIT"
)

type PaymentNewResponseMethod string

const (
	PaymentNewResponseMethodACHNextDay PaymentNewResponseMethod = "ACH_NEXT_DAY"
	PaymentNewResponseMethodACHSameDay PaymentNewResponseMethod = "ACH_SAME_DAY"
)

type PaymentNewResponseMethodAttributes struct {
	SecCode PaymentNewResponseMethodAttributesSecCode `json:"sec_code,required"`
	JSON    paymentNewResponseMethodAttributesJSON
}

// paymentNewResponseMethodAttributesJSON contains the JSON metadata for the struct
// [PaymentNewResponseMethodAttributes]
type paymentNewResponseMethodAttributesJSON struct {
	SecCode     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentNewResponseMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewResponseMethodAttributesSecCode string

const (
	PaymentNewResponseMethodAttributesSecCodePpd PaymentNewResponseMethodAttributesSecCode = "PPD"
	PaymentNewResponseMethodAttributesSecCodeCcd PaymentNewResponseMethodAttributesSecCode = "CCD"
	PaymentNewResponseMethodAttributesSecCodeWeb PaymentNewResponseMethodAttributesSecCode = "WEB"
)

type PaymentNewResponseSource string

const (
	PaymentNewResponseSourceLithic   PaymentNewResponseSource = "LITHIC"
	PaymentNewResponseSourceCustomer PaymentNewResponseSource = "CUSTOMER"
)

type PaymentSimulateReleaseResponse struct {
	DebuggingRequestID    string                               `json:"debugging_request_id" format:"uuid"`
	Result                PaymentSimulateReleaseResponseResult `json:"result"`
	TransactionEventToken string                               `json:"transaction_event_token" format:"uuid"`
	JSON                  paymentSimulateReleaseResponseJSON
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

type PaymentSimulateReleaseResponseResult string

const (
	PaymentSimulateReleaseResponseResultApproved PaymentSimulateReleaseResponseResult = "APPROVED"
	PaymentSimulateReleaseResponseResultDeclined PaymentSimulateReleaseResponseResult = "DECLINED"
)

type PaymentNewParams struct {
	Amount                   param.Field[int64]                            `json:"amount,required"`
	ExternalBankAccountToken param.Field[string]                           `json:"external_bank_account_token,required" format:"uuid"`
	FinancialAccountToken    param.Field[string]                           `json:"financial_account_token,required" format:"uuid"`
	Method                   param.Field[PaymentNewParamsMethod]           `json:"method,required"`
	MethodAttributes         param.Field[PaymentNewParamsMethodAttributes] `json:"method_attributes,required"`
	Type                     param.Field[PaymentNewParamsType]             `json:"type,required"`
	Token                    param.Field[string]                           `json:"token" format:"uuid"`
	Memo                     param.Field[string]                           `json:"memo"`
	UserDefinedID            param.Field[string]                           `json:"user_defined_id"`
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsMethod string

const (
	PaymentNewParamsMethodACHNextDay PaymentNewParamsMethod = "ACH_NEXT_DAY"
	PaymentNewParamsMethodACHSameDay PaymentNewParamsMethod = "ACH_SAME_DAY"
)

type PaymentNewParamsMethodAttributes struct {
	SecCode param.Field[PaymentNewParamsMethodAttributesSecCode] `json:"sec_code,required"`
}

func (r PaymentNewParamsMethodAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsMethodAttributesSecCode string

const (
	PaymentNewParamsMethodAttributesSecCodePpd PaymentNewParamsMethodAttributesSecCode = "PPD"
	PaymentNewParamsMethodAttributesSecCodeCcd PaymentNewParamsMethodAttributesSecCode = "CCD"
	PaymentNewParamsMethodAttributesSecCodeWeb PaymentNewParamsMethodAttributesSecCode = "WEB"
)

type PaymentNewParamsType string

const (
	PaymentNewParamsTypePayment    PaymentNewParamsType = "PAYMENT"
	PaymentNewParamsTypeCollection PaymentNewParamsType = "COLLECTION"
)

type PaymentListParams struct {
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

type PaymentListParamsStatus string

const (
	PaymentListParamsStatusPending  PaymentListParamsStatus = "PENDING"
	PaymentListParamsStatusVoided   PaymentListParamsStatus = "VOIDED"
	PaymentListParamsStatusSettled  PaymentListParamsStatus = "SETTLED"
	PaymentListParamsStatusDeclined PaymentListParamsStatus = "DECLINED"
	PaymentListParamsStatusExpired  PaymentListParamsStatus = "EXPIRED"
)

type PaymentSimulateReleaseParams struct {
	PaymentToken param.Field[string] `json:"payment_token,required" format:"uuid"`
}

func (r PaymentSimulateReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
