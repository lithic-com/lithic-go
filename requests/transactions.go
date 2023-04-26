package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	apijson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type TransactionListParams struct {
	// Filters for transactions associated with a specific account.
	AccountToken field.Field[string] `query:"account_token" format:"uuid"`
	// Filters for transactions associated with a specific card.
	CardToken field.Field[string] `query:"card_token" format:"uuid"`
	// Filters for transactions using transaction result field. Can filter by
	// `APPROVED`, and `DECLINED`.
	Result field.Field[TransactionListParamsResult] `query:"result"`
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
	// Page (for pagination).
	Page field.Field[int64] `query:"page"`
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
}

// URLQuery serializes TransactionListParams into a url.Values of the query
// parameters associated with this value
func (r TransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
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
	Amount field.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor field.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan field.Field[string] `json:"pan,required"`
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
	Status field.Field[TransactionSimulateAuthorizationParamsStatus] `json:"status"`
	// Unique identifier to identify the payment card acceptor.
	MerchantAcceptorID field.Field[string] `json:"merchant_acceptor_id"`
	// 3-digit alphabetic ISO 4217 currency code.
	MerchantCurrency field.Field[string] `json:"merchant_currency"`
	// Amount of the transaction to be simulated in currency specified in
	// merchant_currency, including any acquirer fees.
	MerchantAmount field.Field[int64] `json:"merchant_amount"`
	// Merchant category code for the transaction to be simulated. A four-digit number
	// listed in ISO 18245. Supported merchant category codes can be found
	// [here](https://docs.lithic.com/docs/transactions#merchant-category-codes-mccs).
	Mcc field.Field[string] `json:"mcc"`
	// Set to true if the terminal is capable of partial approval otherwise false.
	// Partial approval is when part of a transaction is approved and another payment
	// must be used for the remainder.
	PartialApprovalCapable field.Field[bool] `json:"partial_approval_capable"`
}

// MarshalJSON serializes TransactionSimulateAuthorizationParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r TransactionSimulateAuthorizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Amount field.Field[int64] `json:"amount,required"`
	// The transaction token returned from the /v1/simulate/authorize response.
	Token field.Field[string] `json:"token,required" format:"uuid"`
}

// MarshalJSON serializes TransactionSimulateAuthorizationAdviceParams into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r TransactionSimulateAuthorizationAdviceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateClearingParams struct {
	// Amount (in cents) to complete. Typically this will match the original
	// authorization, but may be more or less.
	//
	// If no amount is supplied to this endpoint, the amount of the transaction will be
	// captured. Any transaction that has any amount completed at all do not have
	// access to this behavior.
	Amount field.Field[int64] `json:"amount"`
	// The transaction token returned from the /v1/simulate/authorize response.
	Token field.Field[string] `json:"token,required" format:"uuid"`
}

// MarshalJSON serializes TransactionSimulateClearingParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r TransactionSimulateClearingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateCreditAuthorizationParams struct {
	// Amount (in cents). Any value entered will be converted into a negative amount in
	// the simulated transaction. For example, entering 100 in this field will appear
	// as a -100 amount in the transaction.
	Amount field.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor field.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan field.Field[string] `json:"pan,required"`
	// Unique identifier to identify the payment card acceptor.
	MerchantAcceptorID field.Field[string] `json:"merchant_acceptor_id"`
	// Merchant category code for the transaction to be simulated. A four-digit number
	// listed in ISO 18245. Supported merchant category codes can be found
	// [here](https://docs.lithic.com/docs/transactions#merchant-category-codes-mccs).
	Mcc field.Field[string] `json:"mcc"`
}

// MarshalJSON serializes TransactionSimulateCreditAuthorizationParams into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r TransactionSimulateCreditAuthorizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateReturnParams struct {
	// Amount (in cents) to authorize.
	Amount field.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor field.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan field.Field[string] `json:"pan,required"`
}

// MarshalJSON serializes TransactionSimulateReturnParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r TransactionSimulateReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateReturnReversalParams struct {
	// The transaction token returned from the /v1/simulate/authorize response.
	Token field.Field[string] `json:"token,required" format:"uuid"`
}

// MarshalJSON serializes TransactionSimulateReturnReversalParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r TransactionSimulateReturnReversalParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateVoidParams struct {
	// Amount (in cents) to void. Typically this will match the original authorization,
	// but may be less.
	Amount field.Field[int64] `json:"amount"`
	// The transaction token returned from the /v1/simulate/authorize response.
	Token field.Field[string] `json:"token,required" format:"uuid"`
	// Type of event to simulate. Defaults to `AUTHORIZATION_REVERSAL`.
	//
	//   - `AUTHORIZATION_EXPIRY` indicates authorization has expired and been reversed
	//     by Lithic.
	//   - `AUTHORIZATION_REVERSAL` indicates authorization was reversed by the merchant.
	Type field.Field[TransactionSimulateVoidParamsType] `json:"type"`
}

// MarshalJSON serializes TransactionSimulateVoidParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r TransactionSimulateVoidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateVoidParamsType string

const (
	TransactionSimulateVoidParamsTypeAuthorizationExpiry   TransactionSimulateVoidParamsType = "AUTHORIZATION_EXPIRY"
	TransactionSimulateVoidParamsTypeAuthorizationReversal TransactionSimulateVoidParamsType = "AUTHORIZATION_REVERSAL"
)
