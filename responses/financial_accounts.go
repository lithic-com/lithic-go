package responses

import (
	"time"

	pjson "github.com/lithic-com/lithic-go/core/json"
)

type FinancialAccount struct {
	// Account number for your Lithic-assigned bank account number, if applicable.
	AccountNumber string `json:"account_number"`
	// Date and time for when the financial account was first created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Routing number for your Lithic-assigned bank account number, if applicable.
	RoutingNumber string `json:"routing_number"`
	// Globally unique identifier for the financial account.
	Token string `json:"token,required" format:"uuid"`
	// Type of financial account
	Type FinancialAccountType `json:"type,required"`
	// Date and time for when the financial account was last updated.
	Updated time.Time `json:"updated,required" format:"date-time"`
	JSON    FinancialAccountJSON
}

type FinancialAccountJSON struct {
	AccountNumber pjson.Metadata
	Created       pjson.Metadata
	RoutingNumber pjson.Metadata
	Token         pjson.Metadata
	Type          pjson.Metadata
	Updated       pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into FinancialAccount using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *FinancialAccount) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type FinancialAccountType string

const (
	FinancialAccountTypeIssuing FinancialAccountType = "ISSUING"
	FinancialAccountTypeReserve FinancialAccountType = "RESERVE"
)

type FinancialTransaction struct {
	// Status types:
	//
	//   - `CARD` - Issuing card transaction.
	//   - `ACH` - Transaction over ACH.
	//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
	//     program.
	Category FinancialTransactionCategory `json:"category"`
	// Date and time when the financial transaction first occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction.
	Currency string `json:"currency"`
	// A string that provides a description of the financial transaction; may be useful
	// to display to users.
	Descriptor string `json:"descriptor"`
	// A list of all financial events that have modified this financial transaction.
	Events []FinancialTransactionEvents `json:"events"`
	// Pending amount of the transaction in the currency's smallest unit (e.g., cents),
	// including any acquirer fees. The value of this field will go to zero over time
	// once the financial transaction is settled.
	PendingAmount int64 `json:"pending_amount"`
	// APPROVED transactions were successful while DECLINED transactions were declined
	// by user, Lithic, or the network.
	Result FinancialTransactionResult `json:"result"`
	// Amount of the transaction that has been settled in the currency's smallest unit
	// (e.g., cents), including any acquirer fees. This may change over time.
	SettledAmount int64 `json:"settled_amount"`
	// Status types:
	//
	//   - `DECLINED` - The card transaction was declined.
	//   - `EXPIRED` - Lithic reversed the card authorization as it has passed its
	//     expiration time.
	//   - `PENDING` - Authorization is pending completion from the merchant or pending
	//     release from ACH hold period
	//   - `SETTLED` - The financial transaction is completed.
	//   - `VOIDED` - The merchant has voided the previously pending card authorization.
	Status FinancialTransactionStatus `json:"status"`
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Date and time when the financial transaction was last updated. UTC time zone.
	Updated time.Time `json:"updated" format:"date-time"`
	JSON    FinancialTransactionJSON
}

type FinancialTransactionJSON struct {
	Category      pjson.Metadata
	Created       pjson.Metadata
	Currency      pjson.Metadata
	Descriptor    pjson.Metadata
	Events        pjson.Metadata
	PendingAmount pjson.Metadata
	Result        pjson.Metadata
	SettledAmount pjson.Metadata
	Status        pjson.Metadata
	Token         pjson.Metadata
	Updated       pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into FinancialTransaction using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *FinancialTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type FinancialTransactionCategory string

const (
	FinancialTransactionCategoryCard     FinancialTransactionCategory = "CARD"
	FinancialTransactionCategoryACH      FinancialTransactionCategory = "ACH"
	FinancialTransactionCategoryTransfer FinancialTransactionCategory = "TRANSFER"
)

type FinancialTransactionEvents struct {
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result FinancialTransactionEventsResult `json:"result"`
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
	Type FinancialTransactionEventsType `json:"type"`
	JSON FinancialTransactionEventsJSON
}

type FinancialTransactionEventsJSON struct {
	Amount  pjson.Metadata
	Created pjson.Metadata
	Result  pjson.Metadata
	Token   pjson.Metadata
	Type    pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into FinancialTransactionEvents
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *FinancialTransactionEvents) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type FinancialTransactionEventsResult string

const (
	FinancialTransactionEventsResultApproved FinancialTransactionEventsResult = "APPROVED"
	FinancialTransactionEventsResultDeclined FinancialTransactionEventsResult = "DECLINED"
)

type FinancialTransactionEventsType string

const (
	FinancialTransactionEventsTypeACHInsufficientFunds         FinancialTransactionEventsType = "ACH_INSUFFICIENT_FUNDS"
	FinancialTransactionEventsTypeACHOriginationPending        FinancialTransactionEventsType = "ACH_ORIGINATION_PENDING"
	FinancialTransactionEventsTypeACHOriginationReleased       FinancialTransactionEventsType = "ACH_ORIGINATION_RELEASED"
	FinancialTransactionEventsTypeACHReceiptPending            FinancialTransactionEventsType = "ACH_RECEIPT_PENDING"
	FinancialTransactionEventsTypeACHReceiptReleased           FinancialTransactionEventsType = "ACH_RECEIPT_RELEASED"
	FinancialTransactionEventsTypeACHReturn                    FinancialTransactionEventsType = "ACH_RETURN"
	FinancialTransactionEventsTypeAuthorization                FinancialTransactionEventsType = "AUTHORIZATION"
	FinancialTransactionEventsTypeAuthorizationAdvice          FinancialTransactionEventsType = "AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeAuthorizationExpiry          FinancialTransactionEventsType = "AUTHORIZATION_EXPIRY"
	FinancialTransactionEventsTypeAuthorizationReversal        FinancialTransactionEventsType = "AUTHORIZATION_REVERSAL"
	FinancialTransactionEventsTypeBalanceInquiry               FinancialTransactionEventsType = "BALANCE_INQUIRY"
	FinancialTransactionEventsTypeClearing                     FinancialTransactionEventsType = "CLEARING"
	FinancialTransactionEventsTypeCorrectionDebit              FinancialTransactionEventsType = "CORRECTION_DEBIT"
	FinancialTransactionEventsTypeCorrectionCredit             FinancialTransactionEventsType = "CORRECTION_CREDIT"
	FinancialTransactionEventsTypeCreditAuthorization          FinancialTransactionEventsType = "CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeCreditAuthorizationAdvice    FinancialTransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeFinancialAuthorization       FinancialTransactionEventsType = "FINANCIAL_AUTHORIZATION"
	FinancialTransactionEventsTypeFinancialCreditAuthorization FinancialTransactionEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeReturn                       FinancialTransactionEventsType = "RETURN"
	FinancialTransactionEventsTypeReturnReversal               FinancialTransactionEventsType = "RETURN_REVERSAL"
	FinancialTransactionEventsTypeTransfer                     FinancialTransactionEventsType = "TRANSFER"
	FinancialTransactionEventsTypeTransferInsufficientFunds    FinancialTransactionEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
)

type FinancialTransactionResult string

const (
	FinancialTransactionResultApproved FinancialTransactionResult = "APPROVED"
	FinancialTransactionResultDeclined FinancialTransactionResult = "DECLINED"
)

type FinancialTransactionStatus string

const (
	FinancialTransactionStatusDeclined FinancialTransactionStatus = "DECLINED"
	FinancialTransactionStatusExpired  FinancialTransactionStatus = "EXPIRED"
	FinancialTransactionStatusPending  FinancialTransactionStatus = "PENDING"
	FinancialTransactionStatusSettled  FinancialTransactionStatus = "SETTLED"
	FinancialTransactionStatusVoided   FinancialTransactionStatus = "VOIDED"
)

type FinancialAccountListResponse struct {
	Data []FinancialAccount `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    FinancialAccountListResponseJSON
}

type FinancialAccountListResponseJSON struct {
	Data    pjson.Metadata
	HasMore pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into FinancialAccountListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *FinancialAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
