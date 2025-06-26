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

// PaymentService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
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
	path := "v1/payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the payment by token.
func (r *PaymentService) Get(ctx context.Context, paymentToken string, opts ...option.RequestOption) (res *Payment, err error) {
	opts = append(r.Options[:], opts...)
	if paymentToken == "" {
		err = errors.New("missing required payment_token parameter")
		return
	}
	path := fmt.Sprintf("v1/payments/%s", paymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all the payments for the provided search criteria.
func (r *PaymentService) List(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Payment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/payments"
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
	if paymentToken == "" {
		err = errors.New("missing required payment_token parameter")
		return
	}
	path := fmt.Sprintf("v1/payments/%s/retry", paymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulate payment lifecycle event
func (r *PaymentService) SimulateAction(ctx context.Context, paymentToken string, body PaymentSimulateActionParams, opts ...option.RequestOption) (res *PaymentSimulateActionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if paymentToken == "" {
		err = errors.New("missing required payment_token parameter")
		return
	}
	path := fmt.Sprintf("v1/simulate/payments/%s/action", paymentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates a receipt of a Payment.
func (r *PaymentService) SimulateReceipt(ctx context.Context, body PaymentSimulateReceiptParams, opts ...option.RequestOption) (res *PaymentSimulateReceiptResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/payments/receipt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates a release of a Payment.
func (r *PaymentService) SimulateRelease(ctx context.Context, body PaymentSimulateReleaseParams, opts ...option.RequestOption) (res *PaymentSimulateReleaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/payments/release"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates a return of a Payment.
func (r *PaymentService) SimulateReturn(ctx context.Context, body PaymentSimulateReturnParams, opts ...option.RequestOption) (res *PaymentSimulateReturnResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/payments/return"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Payment struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Payment category
	Category PaymentCategory `json:"category,required"`
	// Date and time when the payment first occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-character alphabetic ISO 4217 code for the settling currency of the payment.
	Currency string `json:"currency,required"`
	// A string that provides a description of the payment; may be useful to display to
	// users.
	Descriptor string           `json:"descriptor,required"`
	Direction  PaymentDirection `json:"direction,required"`
	// A list of all payment events that have modified this payment.
	Events                   []PaymentEvent          `json:"events,required"`
	ExternalBankAccountToken string                  `json:"external_bank_account_token,required,nullable" format:"uuid"`
	FinancialAccountToken    string                  `json:"financial_account_token,required" format:"uuid"`
	Method                   PaymentMethod           `json:"method,required"`
	MethodAttributes         PaymentMethodAttributes `json:"method_attributes,required"`
	// Pending amount of the payment in the currency's smallest unit (e.g., cents). The
	// value of this field will go to zero over time once the payment is settled.
	PendingAmount int64 `json:"pending_amount,required"`
	// APPROVED payments were successful while DECLINED payments were declined by
	// Lithic or returned.
	Result PaymentResult `json:"result,required"`
	// Amount of the payment that has been settled in the currency's smallest unit
	// (e.g., cents).
	SettledAmount int64         `json:"settled_amount,required"`
	Source        PaymentSource `json:"source,required"`
	// Status types:
	//
	//   - `DECLINED` - The payment was declined.
	//   - `PENDING` - The payment is being processed and has yet to settle or release
	//     (origination debit).
	//   - `RETURNED` - The payment has been returned.
	//   - `SETTLED` - The payment is completed.
	Status PaymentStatus `json:"status,required"`
	// Date and time when the financial transaction was last updated. UTC time zone.
	Updated       time.Time `json:"updated,required" format:"date-time"`
	UserDefinedID string    `json:"user_defined_id,required,nullable"`
	// Date when the financial transaction expected to be released after settlement
	ExpectedReleaseDate time.Time   `json:"expected_release_date" format:"date"`
	JSON                paymentJSON `json:"-"`
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	Token                    apijson.Field
	Category                 apijson.Field
	Created                  apijson.Field
	Currency                 apijson.Field
	Descriptor               apijson.Field
	Direction                apijson.Field
	Events                   apijson.Field
	ExternalBankAccountToken apijson.Field
	FinancialAccountToken    apijson.Field
	Method                   apijson.Field
	MethodAttributes         apijson.Field
	PendingAmount            apijson.Field
	Result                   apijson.Field
	SettledAmount            apijson.Field
	Source                   apijson.Field
	Status                   apijson.Field
	Updated                  apijson.Field
	UserDefinedID            apijson.Field
	ExpectedReleaseDate      apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Payment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentJSON) RawJSON() string {
	return r.raw
}

// Payment category
type PaymentCategory string

const (
	PaymentCategoryACH PaymentCategory = "ACH"
)

func (r PaymentCategory) IsKnown() bool {
	switch r {
	case PaymentCategoryACH:
		return true
	}
	return false
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

type PaymentEvent struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount,required"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result PaymentEventsResult `json:"result,required"`
	// Event types:
	//
	//   - `ACH_ORIGINATION_INITIATED` - ACH origination received and pending
	//     approval/release from an ACH hold.
	//   - `ACH_ORIGINATION_REVIEWED` - ACH origination has completed the review process.
	//   - `ACH_ORIGINATION_CANCELLED` - ACH origination has been cancelled.
	//   - `ACH_ORIGINATION_PROCESSED` - ACH origination has been processed and sent to
	//     the Federal Reserve.
	//   - `ACH_ORIGINATION_SETTLED` - ACH origination has settled.
	//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
	//     available balance.
	//   - `ACH_RETURN_PROCESSED` - ACH origination returned by the Receiving Depository
	//     Financial Institution.
	//   - `ACH_RECEIPT_PROCESSED` - ACH receipt pending release from an ACH holder.
	//   - `ACH_RETURN_INITIATED` - ACH initiated return for a ACH receipt.
	//   - `ACH_RECEIPT_SETTLED` - ACH receipt funds have settled.
	//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
	//     balance.
	//   - `ACH_RETURN_SETTLED` - ACH receipt return settled by the Receiving Depository
	//     Financial Institution.
	Type PaymentEventsType `json:"type,required"`
	// More detailed reasons for the event
	DetailedResults []PaymentEventsDetailedResult `json:"detailed_results"`
	JSON            paymentEventJSON              `json:"-"`
}

// paymentEventJSON contains the JSON metadata for the struct [PaymentEvent]
type paymentEventJSON struct {
	Token           apijson.Field
	Amount          apijson.Field
	Created         apijson.Field
	Result          apijson.Field
	Type            apijson.Field
	DetailedResults apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PaymentEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentEventJSON) RawJSON() string {
	return r.raw
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type PaymentEventsResult string

const (
	PaymentEventsResultApproved PaymentEventsResult = "APPROVED"
	PaymentEventsResultDeclined PaymentEventsResult = "DECLINED"
)

func (r PaymentEventsResult) IsKnown() bool {
	switch r {
	case PaymentEventsResultApproved, PaymentEventsResultDeclined:
		return true
	}
	return false
}

// Event types:
//
//   - `ACH_ORIGINATION_INITIATED` - ACH origination received and pending
//     approval/release from an ACH hold.
//   - `ACH_ORIGINATION_REVIEWED` - ACH origination has completed the review process.
//   - `ACH_ORIGINATION_CANCELLED` - ACH origination has been cancelled.
//   - `ACH_ORIGINATION_PROCESSED` - ACH origination has been processed and sent to
//     the Federal Reserve.
//   - `ACH_ORIGINATION_SETTLED` - ACH origination has settled.
//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
//     available balance.
//   - `ACH_RETURN_PROCESSED` - ACH origination returned by the Receiving Depository
//     Financial Institution.
//   - `ACH_RECEIPT_PROCESSED` - ACH receipt pending release from an ACH holder.
//   - `ACH_RETURN_INITIATED` - ACH initiated return for a ACH receipt.
//   - `ACH_RECEIPT_SETTLED` - ACH receipt funds have settled.
//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
//     balance.
//   - `ACH_RETURN_SETTLED` - ACH receipt return settled by the Receiving Depository
//     Financial Institution.
type PaymentEventsType string

const (
	PaymentEventsTypeACHOriginationCancelled PaymentEventsType = "ACH_ORIGINATION_CANCELLED"
	PaymentEventsTypeACHOriginationInitiated PaymentEventsType = "ACH_ORIGINATION_INITIATED"
	PaymentEventsTypeACHOriginationProcessed PaymentEventsType = "ACH_ORIGINATION_PROCESSED"
	PaymentEventsTypeACHOriginationSettled   PaymentEventsType = "ACH_ORIGINATION_SETTLED"
	PaymentEventsTypeACHOriginationReleased  PaymentEventsType = "ACH_ORIGINATION_RELEASED"
	PaymentEventsTypeACHOriginationReviewed  PaymentEventsType = "ACH_ORIGINATION_REVIEWED"
	PaymentEventsTypeACHReceiptProcessed     PaymentEventsType = "ACH_RECEIPT_PROCESSED"
	PaymentEventsTypeACHReceiptSettled       PaymentEventsType = "ACH_RECEIPT_SETTLED"
	PaymentEventsTypeACHReturnInitiated      PaymentEventsType = "ACH_RETURN_INITIATED"
	PaymentEventsTypeACHReturnProcessed      PaymentEventsType = "ACH_RETURN_PROCESSED"
	PaymentEventsTypeACHReturnSettled        PaymentEventsType = "ACH_RETURN_SETTLED"
)

func (r PaymentEventsType) IsKnown() bool {
	switch r {
	case PaymentEventsTypeACHOriginationCancelled, PaymentEventsTypeACHOriginationInitiated, PaymentEventsTypeACHOriginationProcessed, PaymentEventsTypeACHOriginationSettled, PaymentEventsTypeACHOriginationReleased, PaymentEventsTypeACHOriginationReviewed, PaymentEventsTypeACHReceiptProcessed, PaymentEventsTypeACHReceiptSettled, PaymentEventsTypeACHReturnInitiated, PaymentEventsTypeACHReturnProcessed, PaymentEventsTypeACHReturnSettled:
		return true
	}
	return false
}

type PaymentEventsDetailedResult string

const (
	PaymentEventsDetailedResultApproved                        PaymentEventsDetailedResult = "APPROVED"
	PaymentEventsDetailedResultFundsInsufficient               PaymentEventsDetailedResult = "FUNDS_INSUFFICIENT"
	PaymentEventsDetailedResultAccountInvalid                  PaymentEventsDetailedResult = "ACCOUNT_INVALID"
	PaymentEventsDetailedResultProgramTransactionLimitExceeded PaymentEventsDetailedResult = "PROGRAM_TRANSACTION_LIMIT_EXCEEDED"
	PaymentEventsDetailedResultProgramDailyLimitExceeded       PaymentEventsDetailedResult = "PROGRAM_DAILY_LIMIT_EXCEEDED"
	PaymentEventsDetailedResultProgramMonthlyLimitExceeded     PaymentEventsDetailedResult = "PROGRAM_MONTHLY_LIMIT_EXCEEDED"
)

func (r PaymentEventsDetailedResult) IsKnown() bool {
	switch r {
	case PaymentEventsDetailedResultApproved, PaymentEventsDetailedResultFundsInsufficient, PaymentEventsDetailedResultAccountInvalid, PaymentEventsDetailedResultProgramTransactionLimitExceeded, PaymentEventsDetailedResultProgramDailyLimitExceeded, PaymentEventsDetailedResultProgramMonthlyLimitExceeded:
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
	TraceNumbers         []string                       `json:"trace_numbers,required"`
	Addenda              string                         `json:"addenda,nullable"`
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
	TraceNumbers         apijson.Field
	Addenda              apijson.Field
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

// APPROVED payments were successful while DECLINED payments were declined by
// Lithic or returned.
type PaymentResult string

const (
	PaymentResultApproved PaymentResult = "APPROVED"
	PaymentResultDeclined PaymentResult = "DECLINED"
)

func (r PaymentResult) IsKnown() bool {
	switch r {
	case PaymentResultApproved, PaymentResultDeclined:
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

// Status types:
//
//   - `DECLINED` - The payment was declined.
//   - `PENDING` - The payment is being processed and has yet to settle or release
//     (origination debit).
//   - `RETURNED` - The payment has been returned.
//   - `SETTLED` - The payment is completed.
type PaymentStatus string

const (
	PaymentStatusDeclined PaymentStatus = "DECLINED"
	PaymentStatusPending  PaymentStatus = "PENDING"
	PaymentStatusReturned PaymentStatus = "RETURNED"
	PaymentStatusSettled  PaymentStatus = "SETTLED"
)

func (r PaymentStatus) IsKnown() bool {
	switch r {
	case PaymentStatusDeclined, PaymentStatusPending, PaymentStatusReturned, PaymentStatusSettled:
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

type PaymentSimulateActionResponse struct {
	// Debugging Request Id
	DebuggingRequestID string `json:"debugging_request_id,required" format:"uuid"`
	// Request Result
	Result PaymentSimulateActionResponseResult `json:"result,required"`
	// Transaction Event Token
	TransactionEventToken string                            `json:"transaction_event_token,required" format:"uuid"`
	JSON                  paymentSimulateActionResponseJSON `json:"-"`
}

// paymentSimulateActionResponseJSON contains the JSON metadata for the struct
// [PaymentSimulateActionResponse]
type paymentSimulateActionResponseJSON struct {
	DebuggingRequestID    apijson.Field
	Result                apijson.Field
	TransactionEventToken apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PaymentSimulateActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSimulateActionResponseJSON) RawJSON() string {
	return r.raw
}

// Request Result
type PaymentSimulateActionResponseResult string

const (
	PaymentSimulateActionResponseResultApproved PaymentSimulateActionResponseResult = "APPROVED"
	PaymentSimulateActionResponseResultDeclined PaymentSimulateActionResponseResult = "DECLINED"
)

func (r PaymentSimulateActionResponseResult) IsKnown() bool {
	switch r {
	case PaymentSimulateActionResponseResultApproved, PaymentSimulateActionResponseResultDeclined:
		return true
	}
	return false
}

type PaymentSimulateReceiptResponse struct {
	// Debugging Request Id
	DebuggingRequestID string `json:"debugging_request_id,required" format:"uuid"`
	// Request Result
	Result PaymentSimulateReceiptResponseResult `json:"result,required"`
	// Transaction Event Token
	TransactionEventToken string                             `json:"transaction_event_token,required" format:"uuid"`
	JSON                  paymentSimulateReceiptResponseJSON `json:"-"`
}

// paymentSimulateReceiptResponseJSON contains the JSON metadata for the struct
// [PaymentSimulateReceiptResponse]
type paymentSimulateReceiptResponseJSON struct {
	DebuggingRequestID    apijson.Field
	Result                apijson.Field
	TransactionEventToken apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PaymentSimulateReceiptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSimulateReceiptResponseJSON) RawJSON() string {
	return r.raw
}

// Request Result
type PaymentSimulateReceiptResponseResult string

const (
	PaymentSimulateReceiptResponseResultApproved PaymentSimulateReceiptResponseResult = "APPROVED"
	PaymentSimulateReceiptResponseResultDeclined PaymentSimulateReceiptResponseResult = "DECLINED"
)

func (r PaymentSimulateReceiptResponseResult) IsKnown() bool {
	switch r {
	case PaymentSimulateReceiptResponseResultApproved, PaymentSimulateReceiptResponseResultDeclined:
		return true
	}
	return false
}

type PaymentSimulateReleaseResponse struct {
	// Debugging Request Id
	DebuggingRequestID string `json:"debugging_request_id,required" format:"uuid"`
	// Request Result
	Result PaymentSimulateReleaseResponseResult `json:"result,required"`
	// Transaction Event Token
	TransactionEventToken string                             `json:"transaction_event_token,required" format:"uuid"`
	JSON                  paymentSimulateReleaseResponseJSON `json:"-"`
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

// Request Result
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
	// Debugging Request Id
	DebuggingRequestID string `json:"debugging_request_id,required" format:"uuid"`
	// Request Result
	Result PaymentSimulateReturnResponseResult `json:"result,required"`
	// Transaction Event Token
	TransactionEventToken string                            `json:"transaction_event_token,required" format:"uuid"`
	JSON                  paymentSimulateReturnResponseJSON `json:"-"`
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

// Request Result
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
	SecCode param.Field[PaymentNewParamsMethodAttributesSecCode] `json:"sec_code,required"`
	Addenda param.Field[string]                                  `json:"addenda"`
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
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin                param.Field[time.Time]                 `query:"begin" format:"date-time"`
	BusinessAccountToken param.Field[string]                    `query:"business_account_token" format:"uuid"`
	Category             param.Field[PaymentListParamsCategory] `query:"category"`
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

type PaymentListParamsCategory string

const (
	PaymentListParamsCategoryACH PaymentListParamsCategory = "ACH"
)

func (r PaymentListParamsCategory) IsKnown() bool {
	switch r {
	case PaymentListParamsCategoryACH:
		return true
	}
	return false
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

type PaymentSimulateActionParams struct {
	// Event Type
	EventType param.Field[PaymentSimulateActionParamsEventType] `json:"event_type,required"`
	// Decline reason
	DeclineReason param.Field[PaymentSimulateActionParamsDeclineReason] `json:"decline_reason"`
	// Return Reason Code
	ReturnReasonCode param.Field[string] `json:"return_reason_code"`
}

func (r PaymentSimulateActionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Event Type
type PaymentSimulateActionParamsEventType string

const (
	PaymentSimulateActionParamsEventTypeACHOriginationReviewed  PaymentSimulateActionParamsEventType = "ACH_ORIGINATION_REVIEWED"
	PaymentSimulateActionParamsEventTypeACHOriginationReleased  PaymentSimulateActionParamsEventType = "ACH_ORIGINATION_RELEASED"
	PaymentSimulateActionParamsEventTypeACHOriginationProcessed PaymentSimulateActionParamsEventType = "ACH_ORIGINATION_PROCESSED"
	PaymentSimulateActionParamsEventTypeACHOriginationSettled   PaymentSimulateActionParamsEventType = "ACH_ORIGINATION_SETTLED"
	PaymentSimulateActionParamsEventTypeACHReceiptSettled       PaymentSimulateActionParamsEventType = "ACH_RECEIPT_SETTLED"
	PaymentSimulateActionParamsEventTypeACHReturnInitiated      PaymentSimulateActionParamsEventType = "ACH_RETURN_INITIATED"
	PaymentSimulateActionParamsEventTypeACHReturnProcessed      PaymentSimulateActionParamsEventType = "ACH_RETURN_PROCESSED"
	PaymentSimulateActionParamsEventTypeACHReturnSettled        PaymentSimulateActionParamsEventType = "ACH_RETURN_SETTLED"
)

func (r PaymentSimulateActionParamsEventType) IsKnown() bool {
	switch r {
	case PaymentSimulateActionParamsEventTypeACHOriginationReviewed, PaymentSimulateActionParamsEventTypeACHOriginationReleased, PaymentSimulateActionParamsEventTypeACHOriginationProcessed, PaymentSimulateActionParamsEventTypeACHOriginationSettled, PaymentSimulateActionParamsEventTypeACHReceiptSettled, PaymentSimulateActionParamsEventTypeACHReturnInitiated, PaymentSimulateActionParamsEventTypeACHReturnProcessed, PaymentSimulateActionParamsEventTypeACHReturnSettled:
		return true
	}
	return false
}

// Decline reason
type PaymentSimulateActionParamsDeclineReason string

const (
	PaymentSimulateActionParamsDeclineReasonProgramTransactionLimitExceeded PaymentSimulateActionParamsDeclineReason = "PROGRAM_TRANSACTION_LIMIT_EXCEEDED"
	PaymentSimulateActionParamsDeclineReasonProgramDailyLimitExceeded       PaymentSimulateActionParamsDeclineReason = "PROGRAM_DAILY_LIMIT_EXCEEDED"
	PaymentSimulateActionParamsDeclineReasonProgramMonthlyLimitExceeded     PaymentSimulateActionParamsDeclineReason = "PROGRAM_MONTHLY_LIMIT_EXCEEDED"
)

func (r PaymentSimulateActionParamsDeclineReason) IsKnown() bool {
	switch r {
	case PaymentSimulateActionParamsDeclineReasonProgramTransactionLimitExceeded, PaymentSimulateActionParamsDeclineReasonProgramDailyLimitExceeded, PaymentSimulateActionParamsDeclineReasonProgramMonthlyLimitExceeded:
		return true
	}
	return false
}

type PaymentSimulateReceiptParams struct {
	// Customer-generated payment token used to uniquely identify the simulated payment
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// Amount
	Amount param.Field[int64] `json:"amount,required"`
	// Financial Account Token
	FinancialAccountToken param.Field[string] `json:"financial_account_token,required" format:"uuid"`
	// Receipt Type
	ReceiptType param.Field[PaymentSimulateReceiptParamsReceiptType] `json:"receipt_type,required"`
	// Memo
	Memo param.Field[string] `json:"memo"`
}

func (r PaymentSimulateReceiptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Receipt Type
type PaymentSimulateReceiptParamsReceiptType string

const (
	PaymentSimulateReceiptParamsReceiptTypeReceiptCredit PaymentSimulateReceiptParamsReceiptType = "RECEIPT_CREDIT"
	PaymentSimulateReceiptParamsReceiptTypeReceiptDebit  PaymentSimulateReceiptParamsReceiptType = "RECEIPT_DEBIT"
)

func (r PaymentSimulateReceiptParamsReceiptType) IsKnown() bool {
	switch r {
	case PaymentSimulateReceiptParamsReceiptTypeReceiptCredit, PaymentSimulateReceiptParamsReceiptTypeReceiptDebit:
		return true
	}
	return false
}

type PaymentSimulateReleaseParams struct {
	// Payment Token
	PaymentToken param.Field[string] `json:"payment_token,required" format:"uuid"`
}

func (r PaymentSimulateReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentSimulateReturnParams struct {
	// Payment Token
	PaymentToken param.Field[string] `json:"payment_token,required" format:"uuid"`
	// Return Reason Code
	ReturnReasonCode param.Field[string] `json:"return_reason_code"`
}

func (r PaymentSimulateReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
