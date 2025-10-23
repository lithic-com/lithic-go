// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/tidwall/gjson"
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
	opts = slices.Concat(r.Options, opts)
	path := "v1/payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the payment by token.
func (r *PaymentService) Get(ctx context.Context, paymentToken string, opts ...option.RequestOption) (res *Payment, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	path := "v1/simulate/payments/receipt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates a release of a Payment.
func (r *PaymentService) SimulateRelease(ctx context.Context, body PaymentSimulateReleaseParams, opts ...option.RequestOption) (res *PaymentSimulateReleaseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/simulate/payments/release"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates a return of a Payment.
func (r *PaymentService) SimulateReturn(ctx context.Context, body PaymentSimulateReturnParams, opts ...option.RequestOption) (res *PaymentSimulateReturnResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/simulate/payments/return"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Payment transaction
type Payment struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// Transaction category
	Category PaymentCategory `json:"category,required"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Transaction descriptor
	Descriptor string `json:"descriptor,required"`
	// Transfer direction
	Direction PaymentDirection `json:"direction,required"`
	// List of transaction events
	Events []PaymentEvent `json:"events,required"`
	// PAYMENT - Payment Transaction
	Family PaymentFamily `json:"family,required"`
	// Financial account token
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Transfer method
	Method PaymentMethod `json:"method,required"`
	// Method-specific attributes
	MethodAttributes PaymentMethodAttributes `json:"method_attributes,required"`
	// Pending amount in cents
	PendingAmount int64 `json:"pending_amount,required"`
	// Related account tokens for the transaction
	RelatedAccountTokens PaymentRelatedAccountTokens `json:"related_account_tokens,required"`
	// Transaction result
	Result PaymentResult `json:"result,required"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount,required"`
	// Transaction source
	Source PaymentSource `json:"source,required"`
	// The status of the transaction
	Status PaymentStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Currency of the transaction in ISO 4217 format
	Currency string `json:"currency"`
	// Expected release date for the transaction
	ExpectedReleaseDate time.Time `json:"expected_release_date,nullable" format:"date"`
	// External bank account token
	ExternalBankAccountToken string      `json:"external_bank_account_token,nullable" format:"uuid"`
	Type                     PaymentType `json:"type"`
	// User-defined identifier
	UserDefinedID string      `json:"user_defined_id,nullable"`
	JSON          paymentJSON `json:"-"`
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	Token                    apijson.Field
	Category                 apijson.Field
	Created                  apijson.Field
	Descriptor               apijson.Field
	Direction                apijson.Field
	Events                   apijson.Field
	Family                   apijson.Field
	FinancialAccountToken    apijson.Field
	Method                   apijson.Field
	MethodAttributes         apijson.Field
	PendingAmount            apijson.Field
	RelatedAccountTokens     apijson.Field
	Result                   apijson.Field
	SettledAmount            apijson.Field
	Source                   apijson.Field
	Status                   apijson.Field
	Updated                  apijson.Field
	Currency                 apijson.Field
	ExpectedReleaseDate      apijson.Field
	ExternalBankAccountToken apijson.Field
	Type                     apijson.Field
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

func (r Payment) implementsAccountActivityListResponse() {}

func (r Payment) implementsAccountActivityGetTransactionResponse() {}

// Transaction category
type PaymentCategory string

const (
	PaymentCategoryACH                    PaymentCategory = "ACH"
	PaymentCategoryBalanceOrFunding       PaymentCategory = "BALANCE_OR_FUNDING"
	PaymentCategoryFee                    PaymentCategory = "FEE"
	PaymentCategoryReward                 PaymentCategory = "REWARD"
	PaymentCategoryAdjustment             PaymentCategory = "ADJUSTMENT"
	PaymentCategoryDerecognition          PaymentCategory = "DERECOGNITION"
	PaymentCategoryDispute                PaymentCategory = "DISPUTE"
	PaymentCategoryCard                   PaymentCategory = "CARD"
	PaymentCategoryExternalACH            PaymentCategory = "EXTERNAL_ACH"
	PaymentCategoryExternalCheck          PaymentCategory = "EXTERNAL_CHECK"
	PaymentCategoryExternalTransfer       PaymentCategory = "EXTERNAL_TRANSFER"
	PaymentCategoryExternalWire           PaymentCategory = "EXTERNAL_WIRE"
	PaymentCategoryManagementAdjustment   PaymentCategory = "MANAGEMENT_ADJUSTMENT"
	PaymentCategoryManagementDispute      PaymentCategory = "MANAGEMENT_DISPUTE"
	PaymentCategoryManagementFee          PaymentCategory = "MANAGEMENT_FEE"
	PaymentCategoryManagementReward       PaymentCategory = "MANAGEMENT_REWARD"
	PaymentCategoryManagementDisbursement PaymentCategory = "MANAGEMENT_DISBURSEMENT"
	PaymentCategoryProgramFunding         PaymentCategory = "PROGRAM_FUNDING"
)

func (r PaymentCategory) IsKnown() bool {
	switch r {
	case PaymentCategoryACH, PaymentCategoryBalanceOrFunding, PaymentCategoryFee, PaymentCategoryReward, PaymentCategoryAdjustment, PaymentCategoryDerecognition, PaymentCategoryDispute, PaymentCategoryCard, PaymentCategoryExternalACH, PaymentCategoryExternalCheck, PaymentCategoryExternalTransfer, PaymentCategoryExternalWire, PaymentCategoryManagementAdjustment, PaymentCategoryManagementDispute, PaymentCategoryManagementFee, PaymentCategoryManagementReward, PaymentCategoryManagementDisbursement, PaymentCategoryProgramFunding:
		return true
	}
	return false
}

// Transfer direction
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

// Payment Event
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

// PAYMENT - Payment Transaction
type PaymentFamily string

const (
	PaymentFamilyPayment PaymentFamily = "PAYMENT"
)

func (r PaymentFamily) IsKnown() bool {
	switch r {
	case PaymentFamilyPayment:
		return true
	}
	return false
}

// Transfer method
type PaymentMethod string

const (
	PaymentMethodACHNextDay PaymentMethod = "ACH_NEXT_DAY"
	PaymentMethodACHSameDay PaymentMethod = "ACH_SAME_DAY"
	PaymentMethodWire       PaymentMethod = "WIRE"
)

func (r PaymentMethod) IsKnown() bool {
	switch r {
	case PaymentMethodACHNextDay, PaymentMethodACHSameDay, PaymentMethodWire:
		return true
	}
	return false
}

// Method-specific attributes
type PaymentMethodAttributes struct {
	// Addenda information
	Addenda string `json:"addenda,nullable"`
	// Company ID for the ACH transaction
	CompanyID string           `json:"company_id,nullable"`
	Creditor  WirePartyDetails `json:"creditor"`
	Debtor    WirePartyDetails `json:"debtor"`
	// Point to point reference identifier, as assigned by the instructing party, used
	// for tracking the message through the Fedwire system
	MessageID string `json:"message_id,nullable"`
	// Receipt routing number
	ReceiptRoutingNumber string `json:"receipt_routing_number,nullable"`
	// Payment details or invoice reference
	RemittanceInformation string `json:"remittance_information,nullable"`
	// Number of retries attempted
	Retries int64 `json:"retries,nullable"`
	// Return reason code if the transaction was returned
	ReturnReasonCode string `json:"return_reason_code,nullable"`
	// SEC code for ACH transaction
	SecCode PaymentMethodAttributesSecCode `json:"sec_code"`
	// This field can have the runtime type of [[]string].
	TraceNumbers interface{} `json:"trace_numbers"`
	// Type of wire message
	WireMessageType string `json:"wire_message_type"`
	// Type of wire transfer
	WireNetwork PaymentMethodAttributesWireNetwork `json:"wire_network"`
	JSON        paymentMethodAttributesJSON        `json:"-"`
	union       PaymentMethodAttributesUnion
}

// paymentMethodAttributesJSON contains the JSON metadata for the struct
// [PaymentMethodAttributes]
type paymentMethodAttributesJSON struct {
	Addenda               apijson.Field
	CompanyID             apijson.Field
	Creditor              apijson.Field
	Debtor                apijson.Field
	MessageID             apijson.Field
	ReceiptRoutingNumber  apijson.Field
	RemittanceInformation apijson.Field
	Retries               apijson.Field
	ReturnReasonCode      apijson.Field
	SecCode               apijson.Field
	TraceNumbers          apijson.Field
	WireMessageType       apijson.Field
	WireNetwork           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r paymentMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r *PaymentMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	*r = PaymentMethodAttributes{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PaymentMethodAttributesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [PaymentMethodAttributesACHMethodAttributes],
// [PaymentMethodAttributesWireMethodAttributes].
func (r PaymentMethodAttributes) AsUnion() PaymentMethodAttributesUnion {
	return r.union
}

// Method-specific attributes
//
// Union satisfied by [PaymentMethodAttributesACHMethodAttributes] or
// [PaymentMethodAttributesWireMethodAttributes].
type PaymentMethodAttributesUnion interface {
	implementsPaymentMethodAttributes()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PaymentMethodAttributesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentMethodAttributesACHMethodAttributes{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentMethodAttributesWireMethodAttributes{}),
		},
	)
}

type PaymentMethodAttributesACHMethodAttributes struct {
	// SEC code for ACH transaction
	SecCode PaymentMethodAttributesACHMethodAttributesSecCode `json:"sec_code,required"`
	// Addenda information
	Addenda string `json:"addenda,nullable"`
	// Company ID for the ACH transaction
	CompanyID string `json:"company_id,nullable"`
	// Receipt routing number
	ReceiptRoutingNumber string `json:"receipt_routing_number,nullable"`
	// Number of retries attempted
	Retries int64 `json:"retries,nullable"`
	// Return reason code if the transaction was returned
	ReturnReasonCode string `json:"return_reason_code,nullable"`
	// Trace numbers for the ACH transaction
	TraceNumbers []string                                       `json:"trace_numbers"`
	JSON         paymentMethodAttributesACHMethodAttributesJSON `json:"-"`
}

// paymentMethodAttributesACHMethodAttributesJSON contains the JSON metadata for
// the struct [PaymentMethodAttributesACHMethodAttributes]
type paymentMethodAttributesACHMethodAttributesJSON struct {
	SecCode              apijson.Field
	Addenda              apijson.Field
	CompanyID            apijson.Field
	ReceiptRoutingNumber apijson.Field
	Retries              apijson.Field
	ReturnReasonCode     apijson.Field
	TraceNumbers         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PaymentMethodAttributesACHMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodAttributesACHMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r PaymentMethodAttributesACHMethodAttributes) implementsPaymentMethodAttributes() {}

// SEC code for ACH transaction
type PaymentMethodAttributesACHMethodAttributesSecCode string

const (
	PaymentMethodAttributesACHMethodAttributesSecCodeCcd PaymentMethodAttributesACHMethodAttributesSecCode = "CCD"
	PaymentMethodAttributesACHMethodAttributesSecCodePpd PaymentMethodAttributesACHMethodAttributesSecCode = "PPD"
	PaymentMethodAttributesACHMethodAttributesSecCodeWeb PaymentMethodAttributesACHMethodAttributesSecCode = "WEB"
	PaymentMethodAttributesACHMethodAttributesSecCodeTel PaymentMethodAttributesACHMethodAttributesSecCode = "TEL"
	PaymentMethodAttributesACHMethodAttributesSecCodeCie PaymentMethodAttributesACHMethodAttributesSecCode = "CIE"
	PaymentMethodAttributesACHMethodAttributesSecCodeCtx PaymentMethodAttributesACHMethodAttributesSecCode = "CTX"
)

func (r PaymentMethodAttributesACHMethodAttributesSecCode) IsKnown() bool {
	switch r {
	case PaymentMethodAttributesACHMethodAttributesSecCodeCcd, PaymentMethodAttributesACHMethodAttributesSecCodePpd, PaymentMethodAttributesACHMethodAttributesSecCodeWeb, PaymentMethodAttributesACHMethodAttributesSecCodeTel, PaymentMethodAttributesACHMethodAttributesSecCodeCie, PaymentMethodAttributesACHMethodAttributesSecCodeCtx:
		return true
	}
	return false
}

type PaymentMethodAttributesWireMethodAttributes struct {
	// Type of wire transfer
	WireNetwork PaymentMethodAttributesWireMethodAttributesWireNetwork `json:"wire_network,required"`
	Creditor    WirePartyDetails                                       `json:"creditor"`
	Debtor      WirePartyDetails                                       `json:"debtor"`
	// Point to point reference identifier, as assigned by the instructing party, used
	// for tracking the message through the Fedwire system
	MessageID string `json:"message_id,nullable"`
	// Payment details or invoice reference
	RemittanceInformation string `json:"remittance_information,nullable"`
	// Type of wire message
	WireMessageType string                                          `json:"wire_message_type"`
	JSON            paymentMethodAttributesWireMethodAttributesJSON `json:"-"`
}

// paymentMethodAttributesWireMethodAttributesJSON contains the JSON metadata for
// the struct [PaymentMethodAttributesWireMethodAttributes]
type paymentMethodAttributesWireMethodAttributesJSON struct {
	WireNetwork           apijson.Field
	Creditor              apijson.Field
	Debtor                apijson.Field
	MessageID             apijson.Field
	RemittanceInformation apijson.Field
	WireMessageType       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PaymentMethodAttributesWireMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodAttributesWireMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r PaymentMethodAttributesWireMethodAttributes) implementsPaymentMethodAttributes() {}

// Type of wire transfer
type PaymentMethodAttributesWireMethodAttributesWireNetwork string

const (
	PaymentMethodAttributesWireMethodAttributesWireNetworkFedwire PaymentMethodAttributesWireMethodAttributesWireNetwork = "FEDWIRE"
	PaymentMethodAttributesWireMethodAttributesWireNetworkSwift   PaymentMethodAttributesWireMethodAttributesWireNetwork = "SWIFT"
)

func (r PaymentMethodAttributesWireMethodAttributesWireNetwork) IsKnown() bool {
	switch r {
	case PaymentMethodAttributesWireMethodAttributesWireNetworkFedwire, PaymentMethodAttributesWireMethodAttributesWireNetworkSwift:
		return true
	}
	return false
}

// SEC code for ACH transaction
type PaymentMethodAttributesSecCode string

const (
	PaymentMethodAttributesSecCodeCcd PaymentMethodAttributesSecCode = "CCD"
	PaymentMethodAttributesSecCodePpd PaymentMethodAttributesSecCode = "PPD"
	PaymentMethodAttributesSecCodeWeb PaymentMethodAttributesSecCode = "WEB"
	PaymentMethodAttributesSecCodeTel PaymentMethodAttributesSecCode = "TEL"
	PaymentMethodAttributesSecCodeCie PaymentMethodAttributesSecCode = "CIE"
	PaymentMethodAttributesSecCodeCtx PaymentMethodAttributesSecCode = "CTX"
)

func (r PaymentMethodAttributesSecCode) IsKnown() bool {
	switch r {
	case PaymentMethodAttributesSecCodeCcd, PaymentMethodAttributesSecCodePpd, PaymentMethodAttributesSecCodeWeb, PaymentMethodAttributesSecCodeTel, PaymentMethodAttributesSecCodeCie, PaymentMethodAttributesSecCodeCtx:
		return true
	}
	return false
}

// Type of wire transfer
type PaymentMethodAttributesWireNetwork string

const (
	PaymentMethodAttributesWireNetworkFedwire PaymentMethodAttributesWireNetwork = "FEDWIRE"
	PaymentMethodAttributesWireNetworkSwift   PaymentMethodAttributesWireNetwork = "SWIFT"
)

func (r PaymentMethodAttributesWireNetwork) IsKnown() bool {
	switch r {
	case PaymentMethodAttributesWireNetworkFedwire, PaymentMethodAttributesWireNetworkSwift:
		return true
	}
	return false
}

// Related account tokens for the transaction
type PaymentRelatedAccountTokens struct {
	// Globally unique identifier for the account
	AccountToken string `json:"account_token,required,nullable" format:"uuid"`
	// Globally unique identifier for the business account
	BusinessAccountToken string                          `json:"business_account_token,required,nullable" format:"uuid"`
	JSON                 paymentRelatedAccountTokensJSON `json:"-"`
}

// paymentRelatedAccountTokensJSON contains the JSON metadata for the struct
// [PaymentRelatedAccountTokens]
type paymentRelatedAccountTokensJSON struct {
	AccountToken         apijson.Field
	BusinessAccountToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PaymentRelatedAccountTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentRelatedAccountTokensJSON) RawJSON() string {
	return r.raw
}

// Transaction result
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

// Transaction source
type PaymentSource string

const (
	PaymentSourceLithic   PaymentSource = "LITHIC"
	PaymentSourceExternal PaymentSource = "EXTERNAL"
	PaymentSourceCustomer PaymentSource = "CUSTOMER"
)

func (r PaymentSource) IsKnown() bool {
	switch r {
	case PaymentSourceLithic, PaymentSourceExternal, PaymentSourceCustomer:
		return true
	}
	return false
}

// The status of the transaction
type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "PENDING"
	PaymentStatusSettled  PaymentStatus = "SETTLED"
	PaymentStatusDeclined PaymentStatus = "DECLINED"
	PaymentStatusReversed PaymentStatus = "REVERSED"
	PaymentStatusCanceled PaymentStatus = "CANCELED"
)

func (r PaymentStatus) IsKnown() bool {
	switch r {
	case PaymentStatusPending, PaymentStatusSettled, PaymentStatusDeclined, PaymentStatusReversed, PaymentStatusCanceled:
		return true
	}
	return false
}

type PaymentType string

const (
	PaymentTypeOriginationCredit   PaymentType = "ORIGINATION_CREDIT"
	PaymentTypeOriginationDebit    PaymentType = "ORIGINATION_DEBIT"
	PaymentTypeReceiptCredit       PaymentType = "RECEIPT_CREDIT"
	PaymentTypeReceiptDebit        PaymentType = "RECEIPT_DEBIT"
	PaymentTypeWireInboundPayment  PaymentType = "WIRE_INBOUND_PAYMENT"
	PaymentTypeWireInboundAdmin    PaymentType = "WIRE_INBOUND_ADMIN"
	PaymentTypeWireOutboundPayment PaymentType = "WIRE_OUTBOUND_PAYMENT"
	PaymentTypeWireOutboundAdmin   PaymentType = "WIRE_OUTBOUND_ADMIN"
)

func (r PaymentType) IsKnown() bool {
	switch r {
	case PaymentTypeOriginationCredit, PaymentTypeOriginationDebit, PaymentTypeReceiptCredit, PaymentTypeReceiptDebit, PaymentTypeWireInboundPayment, PaymentTypeWireInboundAdmin, PaymentTypeWireOutboundPayment, PaymentTypeWireOutboundAdmin:
		return true
	}
	return false
}

// Payment transaction
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

// Payment transaction
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
	PaymentSimulateActionParamsEventTypeACHReceiptReleased      PaymentSimulateActionParamsEventType = "ACH_RECEIPT_RELEASED"
	PaymentSimulateActionParamsEventTypeACHReturnInitiated      PaymentSimulateActionParamsEventType = "ACH_RETURN_INITIATED"
	PaymentSimulateActionParamsEventTypeACHReturnProcessed      PaymentSimulateActionParamsEventType = "ACH_RETURN_PROCESSED"
	PaymentSimulateActionParamsEventTypeACHReturnSettled        PaymentSimulateActionParamsEventType = "ACH_RETURN_SETTLED"
)

func (r PaymentSimulateActionParamsEventType) IsKnown() bool {
	switch r {
	case PaymentSimulateActionParamsEventTypeACHOriginationReviewed, PaymentSimulateActionParamsEventTypeACHOriginationReleased, PaymentSimulateActionParamsEventTypeACHOriginationProcessed, PaymentSimulateActionParamsEventTypeACHOriginationSettled, PaymentSimulateActionParamsEventTypeACHReceiptSettled, PaymentSimulateActionParamsEventTypeACHReceiptReleased, PaymentSimulateActionParamsEventTypeACHReturnInitiated, PaymentSimulateActionParamsEventTypeACHReturnProcessed, PaymentSimulateActionParamsEventTypeACHReturnSettled:
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
