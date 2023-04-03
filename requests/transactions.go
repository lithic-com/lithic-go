package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/fields"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type TransactionListParams struct {
	// Filters for transactions associated with a specific account.
	AccountToken fields.Field[string] `query:"account_token" format:"uuid"`
	// Filters for transactions associated with a specific card.
	CardToken fields.Field[string] `query:"card_token" format:"uuid"`
	// Filters for transactions using transaction result field. Can filter by
	// `APPROVED`, and `DECLINED`.
	Result fields.Field[TransactionListParamsResult] `query:"result"`
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin fields.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End fields.Field[time.Time] `query:"end" format:"date-time"`
	// Page (for pagination).
	Page fields.Field[int64] `query:"page"`
	// Page size (for pagination).
	PageSize fields.Field[int64] `query:"page_size"`
}

// URLQuery serializes TransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *TransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r TransactionListParams) String() (result string) {
	return fmt.Sprintf("&TransactionListParams{AccountToken:%s CardToken:%s Result:%s Begin:%s End:%s Page:%s PageSize:%s}", r.AccountToken, r.CardToken, r.Result, r.Begin, r.End, r.Page, r.PageSize)
}

type TransactionListParamsResult string

const (
	TransactionListParamsResultApproved TransactionListParamsResult = "APPROVED"
	TransactionListParamsResultDeclined TransactionListParamsResult = "DECLINED"
)

type TransactionSimulateAuthorizationParams struct {
	// Amount (in cents) to authorize. For credit authorizations and financial credit
	// authorizations, any value entered will be converted into a negative amount in
	// the simulated transaction. For example, entering 100 in this field will appear
	// as a -100 amount in the transaction. For balance inquiries, this field must be
	// set to 0.
	Amount fields.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor fields.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan fields.Field[string] `json:"pan,required"`
	// Type of event to simulate.
	//
	//   - `AUTHORIZATION` is a dual message purchase authorization, meaning a subsequent
	//     clearing step is required to settle the transaction.
	//   - `BALANCE_INQUIRY` is a $0 authorization that includes a request for the
	//     balance held on the card, and is most typically seen when a cardholder
	//     requests to view a card's balance at an ATM.
	//   - `CREDIT_AUTHORIZATION` is a dual message request from a merchant to authorize
	//     a refund or credit, meaning a subsequent clearing step is required to settle
	//     the transaction.
	//   - `FINANCIAL_AUTHORIZATION` is a single message request from a merchant to debit
	//     funds immediately (such as an ATM withdrawal), and no subsequent clearing is
	//     required to settle the transaction.
	//   - `FINANCIAL_CREDIT_AUTHORIZATION` is a single message request from a merchant
	//     to credit funds immediately, and no subsequent clearing is required to settle
	//     the transaction.
	Status fields.Field[TransactionSimulateAuthorizationParamsStatus] `json:"status"`
	// Unique identifier to identify the payment card acceptor.
	MerchantAcceptorID fields.Field[string] `json:"merchant_acceptor_id"`
	// 3-digit alphabetic ISO 4217 currency code.
	MerchantCurrency fields.Field[string] `json:"merchant_currency"`
	// Amount of the transaction to be simulated in currency specified in
	// merchant_currency, including any acquirer fees.
	MerchantAmount fields.Field[int64] `json:"merchant_amount"`
	// Merchant category code for the transaction to be simulated. A four-digit number
	// listed in ISO 18245. Supported merchant category codes can be found
	// [here](https://docs.lithic.com/docs/transactions#merchant-category-codes-mccs).
	Mcc fields.Field[string] `json:"mcc"`
	// Set to true if the terminal is capable of partial approval otherwise false.
	// Partial approval is when part of a transaction is approved and another payment
	// must be used for the remainder.
	PartialApprovalCapable fields.Field[bool] `json:"partial_approval_capable"`
}

// MarshalJSON serializes TransactionSimulateAuthorizationParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSimulateAuthorizationParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSimulateAuthorizationParams) String() (result string) {
	return fmt.Sprintf("&TransactionSimulateAuthorizationParams{Amount:%s Descriptor:%s Pan:%s Status:%s MerchantAcceptorID:%s MerchantCurrency:%s MerchantAmount:%s Mcc:%s PartialApprovalCapable:%s}", r.Amount, r.Descriptor, r.Pan, r.Status, r.MerchantAcceptorID, r.MerchantCurrency, r.MerchantAmount, r.Mcc, r.PartialApprovalCapable)
}

type TransactionSimulateAuthorizationParamsStatus string

const (
	TransactionSimulateAuthorizationParamsStatusAuthorization                TransactionSimulateAuthorizationParamsStatus = "AUTHORIZATION"
	TransactionSimulateAuthorizationParamsStatusBalanceInquiry               TransactionSimulateAuthorizationParamsStatus = "BALANCE_INQUIRY"
	TransactionSimulateAuthorizationParamsStatusCreditAuthorization          TransactionSimulateAuthorizationParamsStatus = "CREDIT_AUTHORIZATION"
	TransactionSimulateAuthorizationParamsStatusFinancialAuthorization       TransactionSimulateAuthorizationParamsStatus = "FINANCIAL_AUTHORIZATION"
	TransactionSimulateAuthorizationParamsStatusFinancialCreditAuthorization TransactionSimulateAuthorizationParamsStatus = "FINANCIAL_CREDIT_AUTHORIZATION"
)

type TransactionSimulateAuthorizationAdviceParams struct {
	// Amount (in cents) to authorize. This amount will override the transaction's
	// amount that was originally set by /v1/simulate/authorize.
	Amount fields.Field[int64] `json:"amount,required"`
	// The transaction token returned from the /v1/simulate/authorize response.
	Token fields.Field[string] `json:"token,required" format:"uuid"`
}

// MarshalJSON serializes TransactionSimulateAuthorizationAdviceParams into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSimulateAuthorizationAdviceParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSimulateAuthorizationAdviceParams) String() (result string) {
	return fmt.Sprintf("&TransactionSimulateAuthorizationAdviceParams{Amount:%s Token:%s}", r.Amount, r.Token)
}

type TransactionSimulateClearingParams struct {
	// Amount (in cents) to complete. Typically this will match the original
	// authorization, but may be more or less.
	//
	// If no amount is supplied to this endpoint, the amount of the transaction will be
	// captured. Any transaction that has any amount completed at all do not have
	// access to this behavior.
	Amount fields.Field[int64] `json:"amount"`
	// The transaction token returned from the /v1/simulate/authorize response.
	Token fields.Field[string] `json:"token,required" format:"uuid"`
}

// MarshalJSON serializes TransactionSimulateClearingParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSimulateClearingParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSimulateClearingParams) String() (result string) {
	return fmt.Sprintf("&TransactionSimulateClearingParams{Amount:%s Token:%s}", r.Amount, r.Token)
}

type TransactionSimulateCreditAuthorizationParams struct {
	// Amount (in cents). Any value entered will be converted into a negative amount in
	// the simulated transaction. For example, entering 100 in this field will appear
	// as a -100 amount in the transaction.
	Amount fields.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor fields.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan fields.Field[string] `json:"pan,required"`
	// Unique identifier to identify the payment card acceptor.
	MerchantAcceptorID fields.Field[string] `json:"merchant_acceptor_id"`
	// Merchant category code for the transaction to be simulated. A four-digit number
	// listed in ISO 18245. Supported merchant category codes can be found
	// [here](https://docs.lithic.com/docs/transactions#merchant-category-codes-mccs).
	Mcc fields.Field[string] `json:"mcc"`
}

// MarshalJSON serializes TransactionSimulateCreditAuthorizationParams into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSimulateCreditAuthorizationParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSimulateCreditAuthorizationParams) String() (result string) {
	return fmt.Sprintf("&TransactionSimulateCreditAuthorizationParams{Amount:%s Descriptor:%s Pan:%s MerchantAcceptorID:%s Mcc:%s}", r.Amount, r.Descriptor, r.Pan, r.MerchantAcceptorID, r.Mcc)
}

type TransactionSimulateReturnParams struct {
	// Amount (in cents) to authorize.
	Amount fields.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor fields.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan fields.Field[string] `json:"pan,required"`
}

// MarshalJSON serializes TransactionSimulateReturnParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSimulateReturnParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSimulateReturnParams) String() (result string) {
	return fmt.Sprintf("&TransactionSimulateReturnParams{Amount:%s Descriptor:%s Pan:%s}", r.Amount, r.Descriptor, r.Pan)
}

type TransactionSimulateReturnReversalParams struct {
	// The transaction token returned from the /v1/simulate/authorize response.
	Token fields.Field[string] `json:"token,required" format:"uuid"`
}

// MarshalJSON serializes TransactionSimulateReturnReversalParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSimulateReturnReversalParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSimulateReturnReversalParams) String() (result string) {
	return fmt.Sprintf("&TransactionSimulateReturnReversalParams{Token:%s}", r.Token)
}

type TransactionSimulateVoidParams struct {
	// Amount (in cents) to void. Typically this will match the original authorization,
	// but may be less.
	Amount fields.Field[int64] `json:"amount"`
	// The transaction token returned from the /v1/simulate/authorize response.
	Token fields.Field[string] `json:"token,required" format:"uuid"`
	// Type of event to simulate. Defaults to `AUTHORIZATION_REVERSAL`.
	//
	//   - `AUTHORIZATION_EXPIRY` indicates authorization has expired and been reversed
	//     by Lithic.
	//   - `AUTHORIZATION_REVERSAL` indicates authorization was reversed by the merchant.
	Type fields.Field[TransactionSimulateVoidParamsType] `json:"type"`
}

// MarshalJSON serializes TransactionSimulateVoidParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSimulateVoidParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSimulateVoidParams) String() (result string) {
	return fmt.Sprintf("&TransactionSimulateVoidParams{Amount:%s Token:%s Type:%s}", r.Amount, r.Token, r.Type)
}

type TransactionSimulateVoidParamsType string

const (
	TransactionSimulateVoidParamsTypeAuthorizationExpiry   TransactionSimulateVoidParamsType = "AUTHORIZATION_EXPIRY"
	TransactionSimulateVoidParamsTypeAuthorizationReversal TransactionSimulateVoidParamsType = "AUTHORIZATION_REVERSAL"
)
