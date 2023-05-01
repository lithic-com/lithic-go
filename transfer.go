package lithic

import (
	"context"
	"net/http"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/field"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

type TransferService struct {
	Options []option.RequestOption
}

func NewTransferService(opts ...option.RequestOption) (r *TransferService) {
	r = &TransferService{}
	r.Options = opts
	return
}

// Transfer funds between two financial accounts
func (r *TransferService) New(ctx context.Context, body TransferNewParams, opts ...option.RequestOption) (res *TransferCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Transfer struct {
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
	Events []TransferEvents `json:"events"`
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
	// Globally unique identifier for the transfer event.
	Token string `json:"token" format:"uuid"`
	// Date and time when the financial transaction was last updated. UTC time zone.
	Updated time.Time `json:"updated" format:"date-time"`
	// The updated balance of the sending financial account.
	FromBalance []Balance `json:"from_balance"`
	// The updated balance of the receiving financial account.
	ToBalance []Balance `json:"to_balance"`
	JSON      TransferJSON
}

type TransferJSON struct {
	Category      apijson.Metadata
	Created       apijson.Metadata
	Currency      apijson.Metadata
	Descriptor    apijson.Metadata
	Events        apijson.Metadata
	PendingAmount apijson.Metadata
	Result        apijson.Metadata
	SettledAmount apijson.Metadata
	Status        apijson.Metadata
	Token         apijson.Metadata
	Updated       apijson.Metadata
	FromBalance   apijson.Metadata
	ToBalance     apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Transfer using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Transfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TransferCategory string

const (
	TransferCategoryTransfer TransferCategory = "TRANSFER"
)

type TransferEvents struct {
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result TransferEventsResult `json:"result"`
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Event types:
	//
	//   - `ACH_INSUFFICIENT_FUNDS` - Attempted ACH origination declined due to
	//     insufficient balance.
	//   - `ACH_ORIGINATION_PENDING` - ACH origination pending release from an ACH hold.
	//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
	//     available balance.
	//   - `ACH_RECEIPT_PENDING` - ACH receipt pending release from an ACH holder.
	//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
	//     balance.
	//   - `ACH_RETURN` - ACH origination returned by the Receiving Depository Financial
	//     Institution.
	//   - `AUTHORIZATION` - Authorize a card transaction.
	//   - `AUTHORIZATION_ADVICE` - Advice on a card transaction.
	//   - `AUTHORIZATION_EXPIRY` - Card Authorization has expired and reversed by
	//     Lithic.
	//   - `AUTHORIZATION_REVERSAL` - Card Authorization was reversed by the merchant.
	//   - `BALANCE_INQUIRY` - A card balance inquiry (typically a $0 authorization) has
	//     occurred on a card.
	//   - `CLEARING` - Card Transaction is settled.
	//   - `CORRECTION_DEBIT` - Manual card transaction correction (Debit).
	//   - `CORRECTION_CREDIT` - Manual card transaction correction (Credit).
	//   - `CREDIT_AUTHORIZATION` - A refund or credit card authorization from a
	//     merchant.
	//   - `CREDIT_AUTHORIZATION_ADVICE` - A credit card authorization was approved on
	//     your behalf by the network.
	//   - `FINANCIAL_AUTHORIZATION` - A request from a merchant to debit card funds
	//     without additional clearing.
	//   - `FINANCIAL_CREDIT_AUTHORIZATION` - A request from a merchant to refund or
	//     credit card funds without additional clearing.
	//   - `RETURN` - A card refund has been processed on the transaction.
	//   - `RETURN_REVERSAL` - A card refund has been reversed (e.g., when a merchant
	//     reverses an incorrect refund).
	//   - `TRANSFER` - Successful internal transfer of funds between financial accounts.
	//   - `TRANSFER_INSUFFICIENT_FUNDS` - Declined internl transfer of funds due to
	//     insufficient balance of the sender.
	Type TransferEventsType `json:"type"`
	JSON TransferEventsJSON
}

type TransferEventsJSON struct {
	Amount  apijson.Metadata
	Created apijson.Metadata
	Result  apijson.Metadata
	Token   apijson.Metadata
	Type    apijson.Metadata
	raw     string
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransferEvents using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransferEvents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TransferEventsResult string

const (
	TransferEventsResultApproved TransferEventsResult = "APPROVED"
	TransferEventsResultDeclined TransferEventsResult = "DECLINED"
)

type TransferEventsType string

const (
	TransferEventsTypeACHInsufficientFunds         TransferEventsType = "ACH_INSUFFICIENT_FUNDS"
	TransferEventsTypeACHOriginationPending        TransferEventsType = "ACH_ORIGINATION_PENDING"
	TransferEventsTypeACHOriginationReleased       TransferEventsType = "ACH_ORIGINATION_RELEASED"
	TransferEventsTypeACHReceiptPending            TransferEventsType = "ACH_RECEIPT_PENDING"
	TransferEventsTypeACHReceiptReleased           TransferEventsType = "ACH_RECEIPT_RELEASED"
	TransferEventsTypeACHReturn                    TransferEventsType = "ACH_RETURN"
	TransferEventsTypeAuthorization                TransferEventsType = "AUTHORIZATION"
	TransferEventsTypeAuthorizationAdvice          TransferEventsType = "AUTHORIZATION_ADVICE"
	TransferEventsTypeAuthorizationExpiry          TransferEventsType = "AUTHORIZATION_EXPIRY"
	TransferEventsTypeAuthorizationReversal        TransferEventsType = "AUTHORIZATION_REVERSAL"
	TransferEventsTypeBalanceInquiry               TransferEventsType = "BALANCE_INQUIRY"
	TransferEventsTypeClearing                     TransferEventsType = "CLEARING"
	TransferEventsTypeCorrectionDebit              TransferEventsType = "CORRECTION_DEBIT"
	TransferEventsTypeCorrectionCredit             TransferEventsType = "CORRECTION_CREDIT"
	TransferEventsTypeCreditAuthorization          TransferEventsType = "CREDIT_AUTHORIZATION"
	TransferEventsTypeCreditAuthorizationAdvice    TransferEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	TransferEventsTypeFinancialAuthorization       TransferEventsType = "FINANCIAL_AUTHORIZATION"
	TransferEventsTypeFinancialCreditAuthorization TransferEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	TransferEventsTypeReturn                       TransferEventsType = "RETURN"
	TransferEventsTypeReturnReversal               TransferEventsType = "RETURN_REVERSAL"
	TransferEventsTypeTransfer                     TransferEventsType = "TRANSFER"
	TransferEventsTypeTransferInsufficientFunds    TransferEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
)

type TransferResult string

const (
	TransferResultApproved TransferResult = "APPROVED"
	TransferResultDeclined TransferResult = "DECLINED"
)

type TransferStatus string

const (
	TransferStatusDeclined TransferStatus = "DECLINED"
	TransferStatusExpired  TransferStatus = "EXPIRED"
	TransferStatusPending  TransferStatus = "PENDING"
	TransferStatusSettled  TransferStatus = "SETTLED"
	TransferStatusVoided   TransferStatus = "VOIDED"
)

type TransferCreateResponse struct {
	Data Transfer `json:"data"`
	JSON TransferCreateResponseJSON
}

type TransferCreateResponseJSON struct {
	Data   apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransferCreateResponse using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransferCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewParams struct {
	// Financial Account
	From field.Field[FinancialAccountParam] `json:"from,required"`
	// Financial Account
	To field.Field[FinancialAccountParam] `json:"to,required"`
	// Amount to be transferred in the currency’s smallest unit (e.g., cents for USD).
	// This should always be a positive value.
	Amount field.Field[int64] `json:"amount,required"`
	// Optional descriptor for the transfer.
	Memo field.Field[string] `json:"memo"`
	// Customer-provided transaction_token that will serve as an idempotency token.
	TransactionToken field.Field[string] `json:"transaction_token" format:"uuid"`
}

// MarshalJSON serializes TransferNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r TransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
