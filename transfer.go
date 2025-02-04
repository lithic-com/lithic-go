// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// TransferService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferService] method instead.
type TransferService struct {
	Options []option.RequestOption
}

// NewTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransferService(opts ...option.RequestOption) (r *TransferService) {
	r = &TransferService{}
	r.Options = opts
	return
}

// Transfer funds between two financial accounts or between a financial account and
// card
func (r *TransferService) New(ctx context.Context, body TransferNewParams, opts ...option.RequestOption) (res *Transfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Transfer struct {
	// Globally unique identifier for the transfer event.
	Token string `json:"token" format:"uuid"`
	// Status types:
	//
	//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
	//     program.
	Category TransferCategory `json:"category"`
	// Date and time when the transfer occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction.
	Currency string `json:"currency"`
	// A string that provides a description of the transfer; may be useful to display
	// to users.
	Descriptor string `json:"descriptor"`
	// A list of all financial events that have modified this trasnfer.
	Events []TransferEvent `json:"events"`
	// The updated balance of the sending financial account.
	FromBalance []Balance `json:"from_balance"`
	// Pending amount of the transaction in the currency's smallest unit (e.g., cents),
	// including any acquirer fees. The value of this field will go to zero over time
	// once the financial transaction is settled.
	PendingAmount int64 `json:"pending_amount"`
	// APPROVED transactions were successful while DECLINED transactions were declined
	// by user, Lithic, or the network.
	Result TransferResult `json:"result"`
	// Amount of the transaction that has been settled in the currency's smallest unit
	// (e.g., cents).
	SettledAmount int64 `json:"settled_amount"`
	// Status types:
	//
	// - `DECLINED` - The transfer was declined.
	// - `EXPIRED` - The transfer was held in pending for too long and expired.
	// - `PENDING` - The transfer is pending release from a hold.
	// - `SETTLED` - The transfer is completed.
	// - `VOIDED` - The transfer was reversed before it settled.
	Status TransferStatus `json:"status"`
	// The updated balance of the receiving financial account.
	ToBalance []Balance `json:"to_balance"`
	// Date and time when the financial transaction was last updated. UTC time zone.
	Updated time.Time    `json:"updated" format:"date-time"`
	JSON    transferJSON `json:"-"`
}

// transferJSON contains the JSON metadata for the struct [Transfer]
type transferJSON struct {
	Token         apijson.Field
	Category      apijson.Field
	Created       apijson.Field
	Currency      apijson.Field
	Descriptor    apijson.Field
	Events        apijson.Field
	FromBalance   apijson.Field
	PendingAmount apijson.Field
	Result        apijson.Field
	SettledAmount apijson.Field
	Status        apijson.Field
	ToBalance     apijson.Field
	Updated       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Transfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferJSON) RawJSON() string {
	return r.raw
}

// Status types:
//
//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
//     program.
type TransferCategory string

const (
	TransferCategoryTransfer TransferCategory = "TRANSFER"
)

func (r TransferCategory) IsKnown() bool {
	switch r {
	case TransferCategoryTransfer:
		return true
	}
	return false
}

type TransferEvent struct {
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result TransferEventsResult `json:"result"`
	Type   TransferEventsType   `json:"type"`
	JSON   transferEventJSON    `json:"-"`
}

// transferEventJSON contains the JSON metadata for the struct [TransferEvent]
type transferEventJSON struct {
	Token       apijson.Field
	Amount      apijson.Field
	Created     apijson.Field
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferEventJSON) RawJSON() string {
	return r.raw
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type TransferEventsResult string

const (
	TransferEventsResultApproved TransferEventsResult = "APPROVED"
	TransferEventsResultDeclined TransferEventsResult = "DECLINED"
)

func (r TransferEventsResult) IsKnown() bool {
	switch r {
	case TransferEventsResultApproved, TransferEventsResultDeclined:
		return true
	}
	return false
}

type TransferEventsType string

const (
	TransferEventsTypeACHOriginationCancelled      TransferEventsType = "ACH_ORIGINATION_CANCELLED"
	TransferEventsTypeACHOriginationInitiated      TransferEventsType = "ACH_ORIGINATION_INITIATED"
	TransferEventsTypeACHOriginationProcessed      TransferEventsType = "ACH_ORIGINATION_PROCESSED"
	TransferEventsTypeACHOriginationReleased       TransferEventsType = "ACH_ORIGINATION_RELEASED"
	TransferEventsTypeACHOriginationReviewed       TransferEventsType = "ACH_ORIGINATION_REVIEWED"
	TransferEventsTypeACHOriginationSettled        TransferEventsType = "ACH_ORIGINATION_SETTLED"
	TransferEventsTypeACHReceiptProcessed          TransferEventsType = "ACH_RECEIPT_PROCESSED"
	TransferEventsTypeACHReceiptSettled            TransferEventsType = "ACH_RECEIPT_SETTLED"
	TransferEventsTypeACHReturnInitiated           TransferEventsType = "ACH_RETURN_INITIATED"
	TransferEventsTypeACHReturnProcessed           TransferEventsType = "ACH_RETURN_PROCESSED"
	TransferEventsTypeACHReturnSettled             TransferEventsType = "ACH_RETURN_SETTLED"
	TransferEventsTypeAuthorization                TransferEventsType = "AUTHORIZATION"
	TransferEventsTypeAuthorizationAdvice          TransferEventsType = "AUTHORIZATION_ADVICE"
	TransferEventsTypeAuthorizationExpiry          TransferEventsType = "AUTHORIZATION_EXPIRY"
	TransferEventsTypeAuthorizationReversal        TransferEventsType = "AUTHORIZATION_REVERSAL"
	TransferEventsTypeBalanceInquiry               TransferEventsType = "BALANCE_INQUIRY"
	TransferEventsTypeBillingError                 TransferEventsType = "BILLING_ERROR"
	TransferEventsTypeBillingErrorReversal         TransferEventsType = "BILLING_ERROR_REVERSAL"
	TransferEventsTypeCardToCard                   TransferEventsType = "CARD_TO_CARD"
	TransferEventsTypeCashBack                     TransferEventsType = "CASH_BACK"
	TransferEventsTypeCashBackReversal             TransferEventsType = "CASH_BACK_REVERSAL"
	TransferEventsTypeClearing                     TransferEventsType = "CLEARING"
	TransferEventsTypeCorrectionCredit             TransferEventsType = "CORRECTION_CREDIT"
	TransferEventsTypeCorrectionDebit              TransferEventsType = "CORRECTION_DEBIT"
	TransferEventsTypeCreditAuthorization          TransferEventsType = "CREDIT_AUTHORIZATION"
	TransferEventsTypeCreditAuthorizationAdvice    TransferEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	TransferEventsTypeCurrencyConversion           TransferEventsType = "CURRENCY_CONVERSION"
	TransferEventsTypeCurrencyConversionReversal   TransferEventsType = "CURRENCY_CONVERSION_REVERSAL"
	TransferEventsTypeDisputeWon                   TransferEventsType = "DISPUTE_WON"
	TransferEventsTypeExternalACHCanceled          TransferEventsType = "EXTERNAL_ACH_CANCELED"
	TransferEventsTypeExternalACHInitiated         TransferEventsType = "EXTERNAL_ACH_INITIATED"
	TransferEventsTypeExternalACHReleased          TransferEventsType = "EXTERNAL_ACH_RELEASED"
	TransferEventsTypeExternalACHReversed          TransferEventsType = "EXTERNAL_ACH_REVERSED"
	TransferEventsTypeExternalACHSettled           TransferEventsType = "EXTERNAL_ACH_SETTLED"
	TransferEventsTypeExternalCheckCanceled        TransferEventsType = "EXTERNAL_CHECK_CANCELED"
	TransferEventsTypeExternalCheckInitiated       TransferEventsType = "EXTERNAL_CHECK_INITIATED"
	TransferEventsTypeExternalCheckReleased        TransferEventsType = "EXTERNAL_CHECK_RELEASED"
	TransferEventsTypeExternalCheckReversed        TransferEventsType = "EXTERNAL_CHECK_REVERSED"
	TransferEventsTypeExternalCheckSettled         TransferEventsType = "EXTERNAL_CHECK_SETTLED"
	TransferEventsTypeExternalTransferCanceled     TransferEventsType = "EXTERNAL_TRANSFER_CANCELED"
	TransferEventsTypeExternalTransferInitiated    TransferEventsType = "EXTERNAL_TRANSFER_INITIATED"
	TransferEventsTypeExternalTransferReleased     TransferEventsType = "EXTERNAL_TRANSFER_RELEASED"
	TransferEventsTypeExternalTransferReversed     TransferEventsType = "EXTERNAL_TRANSFER_REVERSED"
	TransferEventsTypeExternalTransferSettled      TransferEventsType = "EXTERNAL_TRANSFER_SETTLED"
	TransferEventsTypeExternalWireCanceled         TransferEventsType = "EXTERNAL_WIRE_CANCELED"
	TransferEventsTypeExternalWireInitiated        TransferEventsType = "EXTERNAL_WIRE_INITIATED"
	TransferEventsTypeExternalWireReleased         TransferEventsType = "EXTERNAL_WIRE_RELEASED"
	TransferEventsTypeExternalWireReversed         TransferEventsType = "EXTERNAL_WIRE_REVERSED"
	TransferEventsTypeExternalWireSettled          TransferEventsType = "EXTERNAL_WIRE_SETTLED"
	TransferEventsTypeFinancialAuthorization       TransferEventsType = "FINANCIAL_AUTHORIZATION"
	TransferEventsTypeFinancialCreditAuthorization TransferEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	TransferEventsTypeInterest                     TransferEventsType = "INTEREST"
	TransferEventsTypeInterestReversal             TransferEventsType = "INTEREST_REVERSAL"
	TransferEventsTypeLatePayment                  TransferEventsType = "LATE_PAYMENT"
	TransferEventsTypeLatePaymentReversal          TransferEventsType = "LATE_PAYMENT_REVERSAL"
	TransferEventsTypeProvisionalCredit            TransferEventsType = "PROVISIONAL_CREDIT"
	TransferEventsTypeProvisionalCreditReversal    TransferEventsType = "PROVISIONAL_CREDIT_REVERSAL"
	TransferEventsTypeReturn                       TransferEventsType = "RETURN"
	TransferEventsTypeReturnReversal               TransferEventsType = "RETURN_REVERSAL"
	TransferEventsTypeTransfer                     TransferEventsType = "TRANSFER"
	TransferEventsTypeTransferInsufficientFunds    TransferEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
	TransferEventsTypeReturnedPayment              TransferEventsType = "RETURNED_PAYMENT"
	TransferEventsTypeReturnedPaymentReversal      TransferEventsType = "RETURNED_PAYMENT_REVERSAL"
)

func (r TransferEventsType) IsKnown() bool {
	switch r {
	case TransferEventsTypeACHOriginationCancelled, TransferEventsTypeACHOriginationInitiated, TransferEventsTypeACHOriginationProcessed, TransferEventsTypeACHOriginationReleased, TransferEventsTypeACHOriginationReviewed, TransferEventsTypeACHOriginationSettled, TransferEventsTypeACHReceiptProcessed, TransferEventsTypeACHReceiptSettled, TransferEventsTypeACHReturnInitiated, TransferEventsTypeACHReturnProcessed, TransferEventsTypeACHReturnSettled, TransferEventsTypeAuthorization, TransferEventsTypeAuthorizationAdvice, TransferEventsTypeAuthorizationExpiry, TransferEventsTypeAuthorizationReversal, TransferEventsTypeBalanceInquiry, TransferEventsTypeBillingError, TransferEventsTypeBillingErrorReversal, TransferEventsTypeCardToCard, TransferEventsTypeCashBack, TransferEventsTypeCashBackReversal, TransferEventsTypeClearing, TransferEventsTypeCorrectionCredit, TransferEventsTypeCorrectionDebit, TransferEventsTypeCreditAuthorization, TransferEventsTypeCreditAuthorizationAdvice, TransferEventsTypeCurrencyConversion, TransferEventsTypeCurrencyConversionReversal, TransferEventsTypeDisputeWon, TransferEventsTypeExternalACHCanceled, TransferEventsTypeExternalACHInitiated, TransferEventsTypeExternalACHReleased, TransferEventsTypeExternalACHReversed, TransferEventsTypeExternalACHSettled, TransferEventsTypeExternalCheckCanceled, TransferEventsTypeExternalCheckInitiated, TransferEventsTypeExternalCheckReleased, TransferEventsTypeExternalCheckReversed, TransferEventsTypeExternalCheckSettled, TransferEventsTypeExternalTransferCanceled, TransferEventsTypeExternalTransferInitiated, TransferEventsTypeExternalTransferReleased, TransferEventsTypeExternalTransferReversed, TransferEventsTypeExternalTransferSettled, TransferEventsTypeExternalWireCanceled, TransferEventsTypeExternalWireInitiated, TransferEventsTypeExternalWireReleased, TransferEventsTypeExternalWireReversed, TransferEventsTypeExternalWireSettled, TransferEventsTypeFinancialAuthorization, TransferEventsTypeFinancialCreditAuthorization, TransferEventsTypeInterest, TransferEventsTypeInterestReversal, TransferEventsTypeLatePayment, TransferEventsTypeLatePaymentReversal, TransferEventsTypeProvisionalCredit, TransferEventsTypeProvisionalCreditReversal, TransferEventsTypeReturn, TransferEventsTypeReturnReversal, TransferEventsTypeTransfer, TransferEventsTypeTransferInsufficientFunds, TransferEventsTypeReturnedPayment, TransferEventsTypeReturnedPaymentReversal:
		return true
	}
	return false
}

// APPROVED transactions were successful while DECLINED transactions were declined
// by user, Lithic, or the network.
type TransferResult string

const (
	TransferResultApproved TransferResult = "APPROVED"
	TransferResultDeclined TransferResult = "DECLINED"
)

func (r TransferResult) IsKnown() bool {
	switch r {
	case TransferResultApproved, TransferResultDeclined:
		return true
	}
	return false
}

// Status types:
//
// - `DECLINED` - The transfer was declined.
// - `EXPIRED` - The transfer was held in pending for too long and expired.
// - `PENDING` - The transfer is pending release from a hold.
// - `SETTLED` - The transfer is completed.
// - `VOIDED` - The transfer was reversed before it settled.
type TransferStatus string

const (
	TransferStatusDeclined TransferStatus = "DECLINED"
	TransferStatusExpired  TransferStatus = "EXPIRED"
	TransferStatusPending  TransferStatus = "PENDING"
	TransferStatusSettled  TransferStatus = "SETTLED"
	TransferStatusVoided   TransferStatus = "VOIDED"
)

func (r TransferStatus) IsKnown() bool {
	switch r {
	case TransferStatusDeclined, TransferStatusExpired, TransferStatusPending, TransferStatusSettled, TransferStatusVoided:
		return true
	}
	return false
}

type TransferNewParams struct {
	// Amount to be transferred in the currency’s smallest unit (e.g., cents for USD).
	// This should always be a positive value.
	Amount param.Field[int64] `json:"amount,required"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case.
	From param.Field[string] `json:"from,required" format:"uuid"`
	// Globally unique identifier for the financial account or card that will receive
	// the funds. Accepted type dependent on the program's use case.
	To param.Field[string] `json:"to,required" format:"uuid"`
	// Customer-provided token that will serve as an idempotency token. This token will
	// become the transaction token.
	Token param.Field[string] `json:"token" format:"uuid"`
	// Optional descriptor for the transfer.
	Memo param.Field[string] `json:"memo"`
}

func (r TransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
