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
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// FinancialAccountService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountService] method instead.
type FinancialAccountService struct {
	Options               []option.RequestOption
	Balances              *FinancialAccountBalanceService
	FinancialTransactions *FinancialAccountFinancialTransactionService
	CreditConfiguration   *FinancialAccountCreditConfigurationService
	Statements            *FinancialAccountStatementService
}

// NewFinancialAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFinancialAccountService(opts ...option.RequestOption) (r *FinancialAccountService) {
	r = &FinancialAccountService{}
	r.Options = opts
	r.Balances = NewFinancialAccountBalanceService(opts...)
	r.FinancialTransactions = NewFinancialAccountFinancialTransactionService(opts...)
	r.CreditConfiguration = NewFinancialAccountCreditConfigurationService(opts...)
	r.Statements = NewFinancialAccountStatementService(opts...)
	return
}

// Create a new financial account
func (r *FinancialAccountService) New(ctx context.Context, params FinancialAccountNewParams, opts ...option.RequestOption) (res *FinancialAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "financial_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a financial account
func (r *FinancialAccountService) Get(ctx context.Context, financialAccountToken string, opts ...option.RequestOption) (res *FinancialAccount, err error) {
	opts = append(r.Options[:], opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("financial_accounts/%s", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a financial account
func (r *FinancialAccountService) Update(ctx context.Context, financialAccountToken string, body FinancialAccountUpdateParams, opts ...option.RequestOption) (res *FinancialAccount, err error) {
	opts = append(r.Options[:], opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("financial_accounts/%s", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve information on your financial accounts including routing and account
// number.
func (r *FinancialAccountService) List(ctx context.Context, query FinancialAccountListParams, opts ...option.RequestOption) (res *pagination.SinglePage[FinancialAccount], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *FinancialAccountService) ListAutoPaging(ctx context.Context, query FinancialAccountListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FinancialAccount] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Financial Account
type FinancialAccount struct {
	// Globally unique identifier for the financial account.
	Token string `json:"token,required" format:"uuid"`
	// Date and time for when the financial account was first created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Whether the financial account holds funds for benefit of another party.
	IsForBenefitOf bool `json:"is_for_benefit_of,required"`
	// Type of financial account
	Type FinancialAccountType `json:"type,required"`
	// Date and time for when the financial account was last updated.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Account number for your Lithic-assigned bank account number, if applicable.
	AccountNumber string `json:"account_number"`
	// Account token of the financial account if applicable.
	AccountToken string `json:"account_token" format:"uuid"`
	// User-defined nickname for the financial account.
	Nickname string `json:"nickname"`
	// Routing number for your Lithic-assigned bank account number, if applicable.
	RoutingNumber string               `json:"routing_number"`
	JSON          financialAccountJSON `json:"-"`
}

// financialAccountJSON contains the JSON metadata for the struct
// [FinancialAccount]
type financialAccountJSON struct {
	Token          apijson.Field
	Created        apijson.Field
	IsForBenefitOf apijson.Field
	Type           apijson.Field
	Updated        apijson.Field
	AccountNumber  apijson.Field
	AccountToken   apijson.Field
	Nickname       apijson.Field
	RoutingNumber  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FinancialAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountJSON) RawJSON() string {
	return r.raw
}

// Type of financial account
type FinancialAccountType string

const (
	FinancialAccountTypeIssuing   FinancialAccountType = "ISSUING"
	FinancialAccountTypeOperating FinancialAccountType = "OPERATING"
	FinancialAccountTypeReserve   FinancialAccountType = "RESERVE"
)

func (r FinancialAccountType) IsKnown() bool {
	switch r {
	case FinancialAccountTypeIssuing, FinancialAccountTypeOperating, FinancialAccountTypeReserve:
		return true
	}
	return false
}

type FinancialTransaction struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Status types:
	//
	//   - `CARD` - Issuing card transaction.
	//   - `ACH` - Transaction over ACH.
	//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
	//     program.
	Category FinancialTransactionCategory `json:"category,required"`
	// Date and time when the financial transaction first occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction.
	Currency string `json:"currency,required"`
	// A string that provides a description of the financial transaction; may be useful
	// to display to users.
	Descriptor string `json:"descriptor,required"`
	// A list of all financial events that have modified this financial transaction.
	Events []FinancialTransactionEvent `json:"events,required"`
	// Pending amount of the transaction in the currency's smallest unit (e.g., cents),
	// including any acquirer fees. The value of this field will go to zero over time
	// once the financial transaction is settled.
	PendingAmount int64 `json:"pending_amount,required"`
	// APPROVED transactions were successful while DECLINED transactions were declined
	// by user, Lithic, or the network.
	Result FinancialTransactionResult `json:"result,required"`
	// Amount of the transaction that has been settled in the currency's smallest unit
	// (e.g., cents), including any acquirer fees. This may change over time.
	SettledAmount int64 `json:"settled_amount,required"`
	// Status types:
	//
	//   - `DECLINED` - The transaction was declined.
	//   - `EXPIRED` - The authorization as it has passed its expiration time. Card
	//     transaction only.
	//   - `PENDING` - The transaction is expected to settle.
	//   - `RETURNED` - The transaction has been returned.
	//   - `SETTLED` - The transaction is completed.
	//   - `VOIDED` - The transaction was voided. Card transaction only.
	Status FinancialTransactionStatus `json:"status,required"`
	// Date and time when the financial transaction was last updated. UTC time zone.
	Updated time.Time                `json:"updated,required" format:"date-time"`
	JSON    financialTransactionJSON `json:"-"`
}

// financialTransactionJSON contains the JSON metadata for the struct
// [FinancialTransaction]
type financialTransactionJSON struct {
	Token         apijson.Field
	Category      apijson.Field
	Created       apijson.Field
	Currency      apijson.Field
	Descriptor    apijson.Field
	Events        apijson.Field
	PendingAmount apijson.Field
	Result        apijson.Field
	SettledAmount apijson.Field
	Status        apijson.Field
	Updated       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FinancialTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialTransactionJSON) RawJSON() string {
	return r.raw
}

// Status types:
//
//   - `CARD` - Issuing card transaction.
//   - `ACH` - Transaction over ACH.
//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
//     program.
type FinancialTransactionCategory string

const (
	FinancialTransactionCategoryACH      FinancialTransactionCategory = "ACH"
	FinancialTransactionCategoryCard     FinancialTransactionCategory = "CARD"
	FinancialTransactionCategoryTransfer FinancialTransactionCategory = "TRANSFER"
)

func (r FinancialTransactionCategory) IsKnown() bool {
	switch r {
	case FinancialTransactionCategoryACH, FinancialTransactionCategoryCard, FinancialTransactionCategoryTransfer:
		return true
	}
	return false
}

type FinancialTransactionEvent struct {
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result FinancialTransactionEventsResult `json:"result"`
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
	Type FinancialTransactionEventsType `json:"type"`
	JSON financialTransactionEventJSON  `json:"-"`
}

// financialTransactionEventJSON contains the JSON metadata for the struct
// [FinancialTransactionEvent]
type financialTransactionEventJSON struct {
	Token       apijson.Field
	Amount      apijson.Field
	Created     apijson.Field
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinancialTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialTransactionEventJSON) RawJSON() string {
	return r.raw
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type FinancialTransactionEventsResult string

const (
	FinancialTransactionEventsResultApproved FinancialTransactionEventsResult = "APPROVED"
	FinancialTransactionEventsResultDeclined FinancialTransactionEventsResult = "DECLINED"
)

func (r FinancialTransactionEventsResult) IsKnown() bool {
	switch r {
	case FinancialTransactionEventsResultApproved, FinancialTransactionEventsResultDeclined:
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
type FinancialTransactionEventsType string

const (
	FinancialTransactionEventsTypeACHOriginationCancelled      FinancialTransactionEventsType = "ACH_ORIGINATION_CANCELLED"
	FinancialTransactionEventsTypeACHOriginationInitiated      FinancialTransactionEventsType = "ACH_ORIGINATION_INITIATED"
	FinancialTransactionEventsTypeACHOriginationProcessed      FinancialTransactionEventsType = "ACH_ORIGINATION_PROCESSED"
	FinancialTransactionEventsTypeACHOriginationSettled        FinancialTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	FinancialTransactionEventsTypeACHOriginationReleased       FinancialTransactionEventsType = "ACH_ORIGINATION_RELEASED"
	FinancialTransactionEventsTypeACHOriginationReviewed       FinancialTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	FinancialTransactionEventsTypeACHReceiptProcessed          FinancialTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	FinancialTransactionEventsTypeACHReceiptSettled            FinancialTransactionEventsType = "ACH_RECEIPT_SETTLED"
	FinancialTransactionEventsTypeACHReturnInitiated           FinancialTransactionEventsType = "ACH_RETURN_INITIATED"
	FinancialTransactionEventsTypeACHReturnProcessed           FinancialTransactionEventsType = "ACH_RETURN_PROCESSED"
	FinancialTransactionEventsTypeAuthorization                FinancialTransactionEventsType = "AUTHORIZATION"
	FinancialTransactionEventsTypeAuthorizationAdvice          FinancialTransactionEventsType = "AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeAuthorizationExpiry          FinancialTransactionEventsType = "AUTHORIZATION_EXPIRY"
	FinancialTransactionEventsTypeAuthorizationReversal        FinancialTransactionEventsType = "AUTHORIZATION_REVERSAL"
	FinancialTransactionEventsTypeBalanceInquiry               FinancialTransactionEventsType = "BALANCE_INQUIRY"
	FinancialTransactionEventsTypeClearing                     FinancialTransactionEventsType = "CLEARING"
	FinancialTransactionEventsTypeCorrectionCredit             FinancialTransactionEventsType = "CORRECTION_CREDIT"
	FinancialTransactionEventsTypeCorrectionDebit              FinancialTransactionEventsType = "CORRECTION_DEBIT"
	FinancialTransactionEventsTypeCreditAuthorization          FinancialTransactionEventsType = "CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeCreditAuthorizationAdvice    FinancialTransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeFinancialAuthorization       FinancialTransactionEventsType = "FINANCIAL_AUTHORIZATION"
	FinancialTransactionEventsTypeFinancialCreditAuthorization FinancialTransactionEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeReturn                       FinancialTransactionEventsType = "RETURN"
	FinancialTransactionEventsTypeReturnReversal               FinancialTransactionEventsType = "RETURN_REVERSAL"
	FinancialTransactionEventsTypeTransfer                     FinancialTransactionEventsType = "TRANSFER"
	FinancialTransactionEventsTypeTransferInsufficientFunds    FinancialTransactionEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
)

func (r FinancialTransactionEventsType) IsKnown() bool {
	switch r {
	case FinancialTransactionEventsTypeACHOriginationCancelled, FinancialTransactionEventsTypeACHOriginationInitiated, FinancialTransactionEventsTypeACHOriginationProcessed, FinancialTransactionEventsTypeACHOriginationSettled, FinancialTransactionEventsTypeACHOriginationReleased, FinancialTransactionEventsTypeACHOriginationReviewed, FinancialTransactionEventsTypeACHReceiptProcessed, FinancialTransactionEventsTypeACHReceiptSettled, FinancialTransactionEventsTypeACHReturnInitiated, FinancialTransactionEventsTypeACHReturnProcessed, FinancialTransactionEventsTypeAuthorization, FinancialTransactionEventsTypeAuthorizationAdvice, FinancialTransactionEventsTypeAuthorizationExpiry, FinancialTransactionEventsTypeAuthorizationReversal, FinancialTransactionEventsTypeBalanceInquiry, FinancialTransactionEventsTypeClearing, FinancialTransactionEventsTypeCorrectionCredit, FinancialTransactionEventsTypeCorrectionDebit, FinancialTransactionEventsTypeCreditAuthorization, FinancialTransactionEventsTypeCreditAuthorizationAdvice, FinancialTransactionEventsTypeFinancialAuthorization, FinancialTransactionEventsTypeFinancialCreditAuthorization, FinancialTransactionEventsTypeReturn, FinancialTransactionEventsTypeReturnReversal, FinancialTransactionEventsTypeTransfer, FinancialTransactionEventsTypeTransferInsufficientFunds:
		return true
	}
	return false
}

// APPROVED transactions were successful while DECLINED transactions were declined
// by user, Lithic, or the network.
type FinancialTransactionResult string

const (
	FinancialTransactionResultApproved FinancialTransactionResult = "APPROVED"
	FinancialTransactionResultDeclined FinancialTransactionResult = "DECLINED"
)

func (r FinancialTransactionResult) IsKnown() bool {
	switch r {
	case FinancialTransactionResultApproved, FinancialTransactionResultDeclined:
		return true
	}
	return false
}

// Status types:
//
//   - `DECLINED` - The transaction was declined.
//   - `EXPIRED` - The authorization as it has passed its expiration time. Card
//     transaction only.
//   - `PENDING` - The transaction is expected to settle.
//   - `RETURNED` - The transaction has been returned.
//   - `SETTLED` - The transaction is completed.
//   - `VOIDED` - The transaction was voided. Card transaction only.
type FinancialTransactionStatus string

const (
	FinancialTransactionStatusDeclined FinancialTransactionStatus = "DECLINED"
	FinancialTransactionStatusExpired  FinancialTransactionStatus = "EXPIRED"
	FinancialTransactionStatusPending  FinancialTransactionStatus = "PENDING"
	FinancialTransactionStatusReturned FinancialTransactionStatus = "RETURNED"
	FinancialTransactionStatusSettled  FinancialTransactionStatus = "SETTLED"
	FinancialTransactionStatusVoided   FinancialTransactionStatus = "VOIDED"
)

func (r FinancialTransactionStatus) IsKnown() bool {
	switch r {
	case FinancialTransactionStatusDeclined, FinancialTransactionStatusExpired, FinancialTransactionStatusPending, FinancialTransactionStatusReturned, FinancialTransactionStatusSettled, FinancialTransactionStatusVoided:
		return true
	}
	return false
}

type FinancialAccountNewParams struct {
	Nickname       param.Field[string]                        `json:"nickname,required"`
	Type           param.Field[FinancialAccountNewParamsType] `json:"type,required"`
	AccountToken   param.Field[string]                        `json:"account_token" format:"uuid"`
	IsForBenefitOf param.Field[bool]                          `json:"is_for_benefit_of"`
	IdempotencyKey param.Field[string]                        `header:"Idempotency-Key" format:"uuid"`
}

func (r FinancialAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FinancialAccountNewParamsType string

const (
	FinancialAccountNewParamsTypeOperating FinancialAccountNewParamsType = "OPERATING"
)

func (r FinancialAccountNewParamsType) IsKnown() bool {
	switch r {
	case FinancialAccountNewParamsTypeOperating:
		return true
	}
	return false
}

type FinancialAccountUpdateParams struct {
	Nickname param.Field[string] `json:"nickname"`
}

func (r FinancialAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FinancialAccountListParams struct {
	// List financial accounts for a given account_token or business_account_token
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// List financial accounts for a given business_account_token
	BusinessAccountToken param.Field[string] `query:"business_account_token" format:"uuid"`
	// List financial accounts of a given type
	Type param.Field[FinancialAccountListParamsType] `query:"type"`
}

// URLQuery serializes [FinancialAccountListParams]'s query parameters as
// `url.Values`.
func (r FinancialAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// List financial accounts of a given type
type FinancialAccountListParamsType string

const (
	FinancialAccountListParamsTypeIssuing   FinancialAccountListParamsType = "ISSUING"
	FinancialAccountListParamsTypeOperating FinancialAccountListParamsType = "OPERATING"
	FinancialAccountListParamsTypeReserve   FinancialAccountListParamsType = "RESERVE"
)

func (r FinancialAccountListParamsType) IsKnown() bool {
	switch r {
	case FinancialAccountListParamsTypeIssuing, FinancialAccountListParamsTypeOperating, FinancialAccountListParamsTypeReserve:
		return true
	}
	return false
}
