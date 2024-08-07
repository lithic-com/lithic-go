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
	path := "transfer"
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
	// Event types: _ `ACH_ORIGINATION_INITIATED` - ACH origination received and
	// pending approval/release from an ACH hold. _ `ACH_ORIGINATION_REVIEWED` - ACH
	// origination has completed the review process. _ `ACH_ORIGINATION_CANCELLED` -
	// ACH origination has been cancelled. _ `ACH_ORIGINATION_PROCESSED` - ACH
	// origination has been processed and sent to the fed. _
	// `ACH_ORIGINATION_SETTLED` - ACH origination has settled. _
	// `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to available
	// balance. _ `ACH_RETURN_PROCESSED` - ACH origination returned by the Receiving
	// Depository Financial Institution. _ `ACH_RECEIPT_PROCESSED` - ACH receipt
	// pending release from an ACH holder. _ `ACH_RETURN_INITIATED` - ACH initiated
	// return for a ACH receipt. _ `ACH_RECEIPT_SETTLED` - ACH receipt funds have
	// settled. _ `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to
	// available balance. _ `AUTHORIZATION` - Authorize a card transaction. _
	// `AUTHORIZATION_ADVICE` - Advice on a card transaction. _
	// `AUTHORIZATION_EXPIRY` - Card Authorization has expired and reversed by Lithic.
	// _ `AUTHORIZATION_REVERSAL` - Card Authorization was reversed by the merchant. _
	// `BALANCE_INQUIRY` - A card balance inquiry (typically a $0 authorization) has
	// occurred on a card. _ `CLEARING` - Card Transaction is settled. _
	// `CORRECTION_DEBIT` - Manual card transaction correction (Debit). _
	// `CORRECTION_CREDIT` - Manual card transaction correction (Credit). _
	// `CREDIT_AUTHORIZATION` - A refund or credit card authorization from a merchant.
	// _ `CREDIT_AUTHORIZATION_ADVICE` - A credit card authorization was approved on
	// your behalf by the network. _ `FINANCIAL_AUTHORIZATION` - A request from a
	// merchant to debit card funds without additional clearing. _
	// `FINANCIAL_CREDIT_AUTHORIZATION` - A request from a merchant to refund or credit
	// card funds without additional clearing. _ `RETURN` - A card refund has been
	// processed on the transaction. _ `RETURN_REVERSAL` - A card refund has been
	// reversed (e.g., when a merchant reverses an incorrect refund). _ `TRANSFER` -
	// Successful internal transfer of funds between financial accounts. \*
	// `TRANSFER_INSUFFICIENT_FUNDS` - Declined internal transfer of funds due to
	// insufficient balance of the sender.
	Type TransferEventsType `json:"type"`
	JSON transferEventJSON  `json:"-"`
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

// Event types: _ `ACH_ORIGINATION_INITIATED` - ACH origination received and
// pending approval/release from an ACH hold. _ `ACH_ORIGINATION_REVIEWED` - ACH
// origination has completed the review process. _ `ACH_ORIGINATION_CANCELLED` -
// ACH origination has been cancelled. _ `ACH_ORIGINATION_PROCESSED` - ACH
// origination has been processed and sent to the fed. _
// `ACH_ORIGINATION_SETTLED` - ACH origination has settled. _
// `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to available
// balance. _ `ACH_RETURN_PROCESSED` - ACH origination returned by the Receiving
// Depository Financial Institution. _ `ACH_RECEIPT_PROCESSED` - ACH receipt
// pending release from an ACH holder. _ `ACH_RETURN_INITIATED` - ACH initiated
// return for a ACH receipt. _ `ACH_RECEIPT_SETTLED` - ACH receipt funds have
// settled. _ `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to
// available balance. _ `AUTHORIZATION` - Authorize a card transaction. _
// `AUTHORIZATION_ADVICE` - Advice on a card transaction. _
// `AUTHORIZATION_EXPIRY` - Card Authorization has expired and reversed by Lithic.
// _ `AUTHORIZATION_REVERSAL` - Card Authorization was reversed by the merchant. _
// `BALANCE_INQUIRY` - A card balance inquiry (typically a $0 authorization) has
// occurred on a card. _ `CLEARING` - Card Transaction is settled. _
// `CORRECTION_DEBIT` - Manual card transaction correction (Debit). _
// `CORRECTION_CREDIT` - Manual card transaction correction (Credit). _
// `CREDIT_AUTHORIZATION` - A refund or credit card authorization from a merchant.
// _ `CREDIT_AUTHORIZATION_ADVICE` - A credit card authorization was approved on
// your behalf by the network. _ `FINANCIAL_AUTHORIZATION` - A request from a
// merchant to debit card funds without additional clearing. _
// `FINANCIAL_CREDIT_AUTHORIZATION` - A request from a merchant to refund or credit
// card funds without additional clearing. _ `RETURN` - A card refund has been
// processed on the transaction. _ `RETURN_REVERSAL` - A card refund has been
// reversed (e.g., when a merchant reverses an incorrect refund). _ `TRANSFER` -
// Successful internal transfer of funds between financial accounts. \*
// `TRANSFER_INSUFFICIENT_FUNDS` - Declined internal transfer of funds due to
// insufficient balance of the sender.
type TransferEventsType string

const (
	TransferEventsTypeACHOriginationCancelled      TransferEventsType = "ACH_ORIGINATION_CANCELLED"
	TransferEventsTypeACHOriginationInitiated      TransferEventsType = "ACH_ORIGINATION_INITIATED"
	TransferEventsTypeACHOriginationProcessed      TransferEventsType = "ACH_ORIGINATION_PROCESSED"
	TransferEventsTypeACHOriginationSettled        TransferEventsType = "ACH_ORIGINATION_SETTLED"
	TransferEventsTypeACHOriginationReleased       TransferEventsType = "ACH_ORIGINATION_RELEASED"
	TransferEventsTypeACHOriginationReviewed       TransferEventsType = "ACH_ORIGINATION_REVIEWED"
	TransferEventsTypeACHReceiptProcessed          TransferEventsType = "ACH_RECEIPT_PROCESSED"
	TransferEventsTypeACHReceiptSettled            TransferEventsType = "ACH_RECEIPT_SETTLED"
	TransferEventsTypeACHReturnInitiated           TransferEventsType = "ACH_RETURN_INITIATED"
	TransferEventsTypeACHReturnProcessed           TransferEventsType = "ACH_RETURN_PROCESSED"
	TransferEventsTypeAuthorization                TransferEventsType = "AUTHORIZATION"
	TransferEventsTypeAuthorizationAdvice          TransferEventsType = "AUTHORIZATION_ADVICE"
	TransferEventsTypeAuthorizationExpiry          TransferEventsType = "AUTHORIZATION_EXPIRY"
	TransferEventsTypeAuthorizationReversal        TransferEventsType = "AUTHORIZATION_REVERSAL"
	TransferEventsTypeBalanceInquiry               TransferEventsType = "BALANCE_INQUIRY"
	TransferEventsTypeClearing                     TransferEventsType = "CLEARING"
	TransferEventsTypeCorrectionCredit             TransferEventsType = "CORRECTION_CREDIT"
	TransferEventsTypeCorrectionDebit              TransferEventsType = "CORRECTION_DEBIT"
	TransferEventsTypeCreditAuthorization          TransferEventsType = "CREDIT_AUTHORIZATION"
	TransferEventsTypeCreditAuthorizationAdvice    TransferEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	TransferEventsTypeFinancialAuthorization       TransferEventsType = "FINANCIAL_AUTHORIZATION"
	TransferEventsTypeFinancialCreditAuthorization TransferEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	TransferEventsTypeReturn                       TransferEventsType = "RETURN"
	TransferEventsTypeReturnReversal               TransferEventsType = "RETURN_REVERSAL"
	TransferEventsTypeTransfer                     TransferEventsType = "TRANSFER"
	TransferEventsTypeTransferInsufficientFunds    TransferEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
)

func (r TransferEventsType) IsKnown() bool {
	switch r {
	case TransferEventsTypeACHOriginationCancelled, TransferEventsTypeACHOriginationInitiated, TransferEventsTypeACHOriginationProcessed, TransferEventsTypeACHOriginationSettled, TransferEventsTypeACHOriginationReleased, TransferEventsTypeACHOriginationReviewed, TransferEventsTypeACHReceiptProcessed, TransferEventsTypeACHReceiptSettled, TransferEventsTypeACHReturnInitiated, TransferEventsTypeACHReturnProcessed, TransferEventsTypeAuthorization, TransferEventsTypeAuthorizationAdvice, TransferEventsTypeAuthorizationExpiry, TransferEventsTypeAuthorizationReversal, TransferEventsTypeBalanceInquiry, TransferEventsTypeClearing, TransferEventsTypeCorrectionCredit, TransferEventsTypeCorrectionDebit, TransferEventsTypeCreditAuthorization, TransferEventsTypeCreditAuthorizationAdvice, TransferEventsTypeFinancialAuthorization, TransferEventsTypeFinancialCreditAuthorization, TransferEventsTypeReturn, TransferEventsTypeReturnReversal, TransferEventsTypeTransfer, TransferEventsTypeTransferInsufficientFunds:
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
