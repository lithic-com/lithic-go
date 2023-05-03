package lithic

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/field"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// FinancialAccountService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFinancialAccountService] method
// instead.
type FinancialAccountService struct {
	Options               []option.RequestOption
	Balances              *FinancialAccountBalanceService
	FinancialTransactions *FinancialAccountFinancialTransactionService
}

// NewFinancialAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFinancialAccountService(opts ...option.RequestOption) (r *FinancialAccountService) {
	r = &FinancialAccountService{}
	r.Options = opts
	r.Balances = NewFinancialAccountBalanceService(opts...)
	r.FinancialTransactions = NewFinancialAccountFinancialTransactionService(opts...)
	return
}

// Retrieve information on your financial accounts including routing and account
// number.
func (r *FinancialAccountService) List(ctx context.Context, query FinancialAccountListParams, opts ...option.RequestOption) (res *shared.SinglePage[FinancialAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "financial_accounts"
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

// Retrieve information on your financial accounts including routing and account
// number.
func (r *FinancialAccountService) ListAutoPaging(ctx context.Context, query FinancialAccountListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[FinancialAccount] {
	return shared.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Financial Account
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
	JSON    financialAccountJSON
}

// financialAccountJSON contains the JSON metadata for the struct
// [FinancialAccount]
type financialAccountJSON struct {
	AccountNumber apijson.Field
	Created       apijson.Field
	RoutingNumber apijson.Field
	Token         apijson.Field
	Type          apijson.Field
	Updated       apijson.Field
	raw           string
	Extras        map[string]apijson.Field
}

func (r *FinancialAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Financial Account
type FinancialAccountParam struct {
	// Account number for your Lithic-assigned bank account number, if applicable.
	AccountNumber field.Field[string] `json:"account_number"`
	// Date and time for when the financial account was first created.
	Created field.Field[time.Time] `json:"created,required" format:"date-time"`
	// Routing number for your Lithic-assigned bank account number, if applicable.
	RoutingNumber field.Field[string] `json:"routing_number"`
	// Globally unique identifier for the financial account.
	Token field.Field[string] `json:"token,required" format:"uuid"`
	// Type of financial account
	Type field.Field[FinancialAccountType] `json:"type,required"`
	// Date and time for when the financial account was last updated.
	Updated field.Field[time.Time] `json:"updated,required" format:"date-time"`
}

func (r FinancialAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	JSON    financialTransactionJSON
}

// financialTransactionJSON contains the JSON metadata for the struct
// [FinancialTransaction]
type financialTransactionJSON struct {
	Category      apijson.Field
	Created       apijson.Field
	Currency      apijson.Field
	Descriptor    apijson.Field
	Events        apijson.Field
	PendingAmount apijson.Field
	Result        apijson.Field
	SettledAmount apijson.Field
	Status        apijson.Field
	Token         apijson.Field
	Updated       apijson.Field
	raw           string
	Extras        map[string]apijson.Field
}

func (r *FinancialTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	JSON financialTransactionEventsJSON
}

// financialTransactionEventsJSON contains the JSON metadata for the struct
// [FinancialTransactionEvents]
type financialTransactionEventsJSON struct {
	Amount  apijson.Field
	Created apijson.Field
	Result  apijson.Field
	Token   apijson.Field
	Type    apijson.Field
	raw     string
	Extras  map[string]apijson.Field
}

func (r *FinancialTransactionEvents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type FinancialAccountListParams struct {
	// List financial accounts for a given account_token
	AccountToken field.Field[string] `query:"account_token" format:"uuid"`
	// List financial accounts of a given type
	Type field.Field[FinancialAccountListParamsType] `query:"type"`
}

// URLQuery serializes [FinancialAccountListParams]'s query parameters as
// `url.Values`.
func (r FinancialAccountListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type FinancialAccountListParamsType string

const (
	FinancialAccountListParamsTypeIssuing FinancialAccountListParamsType = "ISSUING"
	FinancialAccountListParamsTypeReserve FinancialAccountListParamsType = "RESERVE"
)

type FinancialAccountListResponse struct {
	Data []FinancialAccount `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    financialAccountListResponseJSON
}

// financialAccountListResponseJSON contains the JSON metadata for the struct
// [FinancialAccountListResponse]
type financialAccountListResponseJSON struct {
	Data    apijson.Field
	HasMore apijson.Field
	raw     string
	Extras  map[string]apijson.Field
}

func (r *FinancialAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
