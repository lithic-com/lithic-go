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
	LoanTapes             *FinancialAccountLoanTapeService
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
	r.LoanTapes = NewFinancialAccountLoanTapeService(opts...)
	return
}

// Create a new financial account
func (r *FinancialAccountService) New(ctx context.Context, params FinancialAccountNewParams, opts ...option.RequestOption) (res *FinancialAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/financial_accounts"
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
	path := fmt.Sprintf("v1/financial_accounts/%s", financialAccountToken)
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
	path := fmt.Sprintf("v1/financial_accounts/%s", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve information on your financial accounts including routing and account
// number.
func (r *FinancialAccountService) List(ctx context.Context, query FinancialAccountListParams, opts ...option.RequestOption) (res *pagination.SinglePage[FinancialAccount], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/financial_accounts"
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

type FinancialAccount struct {
	// Globally unique identifier for the account
	Token               string                              `json:"token,required" format:"uuid"`
	AccountToken        string                              `json:"account_token,required,nullable" format:"uuid"`
	Created             time.Time                           `json:"created,required" format:"date-time"`
	CreditConfiguration FinancialAccountCreditConfiguration `json:"credit_configuration,required,nullable"`
	// Whether financial account is for the benefit of another entity
	IsForBenefitOf bool                 `json:"is_for_benefit_of,required"`
	Nickname       string               `json:"nickname,required,nullable"`
	Type           FinancialAccountType `json:"type,required"`
	Updated        time.Time            `json:"updated,required" format:"date-time"`
	AccountNumber  string               `json:"account_number,nullable"`
	RoutingNumber  string               `json:"routing_number,nullable"`
	JSON           financialAccountJSON `json:"-"`
}

// financialAccountJSON contains the JSON metadata for the struct
// [FinancialAccount]
type financialAccountJSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	Created             apijson.Field
	CreditConfiguration apijson.Field
	IsForBenefitOf      apijson.Field
	Nickname            apijson.Field
	Type                apijson.Field
	Updated             apijson.Field
	AccountNumber       apijson.Field
	RoutingNumber       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *FinancialAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountJSON) RawJSON() string {
	return r.raw
}

type FinancialAccountCreditConfiguration struct {
	CreditLimit int64 `json:"credit_limit,required,nullable"`
	// Globally unique identifier for the credit product
	CreditProductToken       string `json:"credit_product_token,required,nullable"`
	ExternalBankAccountToken string `json:"external_bank_account_token,required,nullable" format:"uuid"`
	// Tier assigned to the financial account
	Tier string `json:"tier,required,nullable"`
	// State of the financial account
	FinancialAccountState FinancialAccountCreditConfigurationFinancialAccountState `json:"financial_account_state"`
	JSON                  financialAccountCreditConfigurationJSON                  `json:"-"`
}

// financialAccountCreditConfigurationJSON contains the JSON metadata for the
// struct [FinancialAccountCreditConfiguration]
type financialAccountCreditConfigurationJSON struct {
	CreditLimit              apijson.Field
	CreditProductToken       apijson.Field
	ExternalBankAccountToken apijson.Field
	Tier                     apijson.Field
	FinancialAccountState    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *FinancialAccountCreditConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountCreditConfigurationJSON) RawJSON() string {
	return r.raw
}

// State of the financial account
type FinancialAccountCreditConfigurationFinancialAccountState string

const (
	FinancialAccountCreditConfigurationFinancialAccountStatePending    FinancialAccountCreditConfigurationFinancialAccountState = "PENDING"
	FinancialAccountCreditConfigurationFinancialAccountStateCurrent    FinancialAccountCreditConfigurationFinancialAccountState = "CURRENT"
	FinancialAccountCreditConfigurationFinancialAccountStateDelinquent FinancialAccountCreditConfigurationFinancialAccountState = "DELINQUENT"
)

func (r FinancialAccountCreditConfigurationFinancialAccountState) IsKnown() bool {
	switch r {
	case FinancialAccountCreditConfigurationFinancialAccountStatePending, FinancialAccountCreditConfigurationFinancialAccountStateCurrent, FinancialAccountCreditConfigurationFinancialAccountStateDelinquent:
		return true
	}
	return false
}

type FinancialAccountType string

const (
	FinancialAccountTypeIssuing   FinancialAccountType = "ISSUING"
	FinancialAccountTypeReserve   FinancialAccountType = "RESERVE"
	FinancialAccountTypeOperating FinancialAccountType = "OPERATING"
)

func (r FinancialAccountType) IsKnown() bool {
	switch r {
	case FinancialAccountTypeIssuing, FinancialAccountTypeReserve, FinancialAccountTypeOperating:
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
	Type   FinancialTransactionEventsType   `json:"type"`
	JSON   financialTransactionEventJSON    `json:"-"`
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

type FinancialTransactionEventsType string

const (
	FinancialTransactionEventsTypeACHOriginationCancelled      FinancialTransactionEventsType = "ACH_ORIGINATION_CANCELLED"
	FinancialTransactionEventsTypeACHOriginationInitiated      FinancialTransactionEventsType = "ACH_ORIGINATION_INITIATED"
	FinancialTransactionEventsTypeACHOriginationProcessed      FinancialTransactionEventsType = "ACH_ORIGINATION_PROCESSED"
	FinancialTransactionEventsTypeACHOriginationReleased       FinancialTransactionEventsType = "ACH_ORIGINATION_RELEASED"
	FinancialTransactionEventsTypeACHOriginationReviewed       FinancialTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	FinancialTransactionEventsTypeACHOriginationSettled        FinancialTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	FinancialTransactionEventsTypeACHReceiptProcessed          FinancialTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	FinancialTransactionEventsTypeACHReceiptSettled            FinancialTransactionEventsType = "ACH_RECEIPT_SETTLED"
	FinancialTransactionEventsTypeACHReturnInitiated           FinancialTransactionEventsType = "ACH_RETURN_INITIATED"
	FinancialTransactionEventsTypeACHReturnProcessed           FinancialTransactionEventsType = "ACH_RETURN_PROCESSED"
	FinancialTransactionEventsTypeAuthorization                FinancialTransactionEventsType = "AUTHORIZATION"
	FinancialTransactionEventsTypeAuthorizationAdvice          FinancialTransactionEventsType = "AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeAuthorizationExpiry          FinancialTransactionEventsType = "AUTHORIZATION_EXPIRY"
	FinancialTransactionEventsTypeAuthorizationReversal        FinancialTransactionEventsType = "AUTHORIZATION_REVERSAL"
	FinancialTransactionEventsTypeBalanceInquiry               FinancialTransactionEventsType = "BALANCE_INQUIRY"
	FinancialTransactionEventsTypeBillingError                 FinancialTransactionEventsType = "BILLING_ERROR"
	FinancialTransactionEventsTypeCashBack                     FinancialTransactionEventsType = "CASH_BACK"
	FinancialTransactionEventsTypeClearing                     FinancialTransactionEventsType = "CLEARING"
	FinancialTransactionEventsTypeCorrectionCredit             FinancialTransactionEventsType = "CORRECTION_CREDIT"
	FinancialTransactionEventsTypeCorrectionDebit              FinancialTransactionEventsType = "CORRECTION_DEBIT"
	FinancialTransactionEventsTypeCreditAuthorization          FinancialTransactionEventsType = "CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeCreditAuthorizationAdvice    FinancialTransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeCurrencyConversion           FinancialTransactionEventsType = "CURRENCY_CONVERSION"
	FinancialTransactionEventsTypeDisputeWon                   FinancialTransactionEventsType = "DISPUTE_WON"
	FinancialTransactionEventsTypeExternalACHCanceled          FinancialTransactionEventsType = "EXTERNAL_ACH_CANCELED"
	FinancialTransactionEventsTypeExternalACHInitiated         FinancialTransactionEventsType = "EXTERNAL_ACH_INITIATED"
	FinancialTransactionEventsTypeExternalACHReleased          FinancialTransactionEventsType = "EXTERNAL_ACH_RELEASED"
	FinancialTransactionEventsTypeExternalACHReversed          FinancialTransactionEventsType = "EXTERNAL_ACH_REVERSED"
	FinancialTransactionEventsTypeExternalACHSettled           FinancialTransactionEventsType = "EXTERNAL_ACH_SETTLED"
	FinancialTransactionEventsTypeExternalCheckCanceled        FinancialTransactionEventsType = "EXTERNAL_CHECK_CANCELED"
	FinancialTransactionEventsTypeExternalCheckInitiated       FinancialTransactionEventsType = "EXTERNAL_CHECK_INITIATED"
	FinancialTransactionEventsTypeExternalCheckReleased        FinancialTransactionEventsType = "EXTERNAL_CHECK_RELEASED"
	FinancialTransactionEventsTypeExternalCheckReversed        FinancialTransactionEventsType = "EXTERNAL_CHECK_REVERSED"
	FinancialTransactionEventsTypeExternalCheckSettled         FinancialTransactionEventsType = "EXTERNAL_CHECK_SETTLED"
	FinancialTransactionEventsTypeExternalTransferCanceled     FinancialTransactionEventsType = "EXTERNAL_TRANSFER_CANCELED"
	FinancialTransactionEventsTypeExternalTransferInitiated    FinancialTransactionEventsType = "EXTERNAL_TRANSFER_INITIATED"
	FinancialTransactionEventsTypeExternalTransferReleased     FinancialTransactionEventsType = "EXTERNAL_TRANSFER_RELEASED"
	FinancialTransactionEventsTypeExternalTransferReversed     FinancialTransactionEventsType = "EXTERNAL_TRANSFER_REVERSED"
	FinancialTransactionEventsTypeExternalTransferSettled      FinancialTransactionEventsType = "EXTERNAL_TRANSFER_SETTLED"
	FinancialTransactionEventsTypeExternalWireCanceled         FinancialTransactionEventsType = "EXTERNAL_WIRE_CANCELED"
	FinancialTransactionEventsTypeExternalWireInitiated        FinancialTransactionEventsType = "EXTERNAL_WIRE_INITIATED"
	FinancialTransactionEventsTypeExternalWireReleased         FinancialTransactionEventsType = "EXTERNAL_WIRE_RELEASED"
	FinancialTransactionEventsTypeExternalWireReversed         FinancialTransactionEventsType = "EXTERNAL_WIRE_REVERSED"
	FinancialTransactionEventsTypeExternalWireSettled          FinancialTransactionEventsType = "EXTERNAL_WIRE_SETTLED"
	FinancialTransactionEventsTypeFinancialAuthorization       FinancialTransactionEventsType = "FINANCIAL_AUTHORIZATION"
	FinancialTransactionEventsTypeFinancialCreditAuthorization FinancialTransactionEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeInterest                     FinancialTransactionEventsType = "INTEREST"
	FinancialTransactionEventsTypeLatePayment                  FinancialTransactionEventsType = "LATE_PAYMENT"
	FinancialTransactionEventsTypeProvisionalCredit            FinancialTransactionEventsType = "PROVISIONAL_CREDIT"
	FinancialTransactionEventsTypeReturn                       FinancialTransactionEventsType = "RETURN"
	FinancialTransactionEventsTypeReturnReversal               FinancialTransactionEventsType = "RETURN_REVERSAL"
	FinancialTransactionEventsTypeTransfer                     FinancialTransactionEventsType = "TRANSFER"
	FinancialTransactionEventsTypeTransferInsufficientFunds    FinancialTransactionEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
)

func (r FinancialTransactionEventsType) IsKnown() bool {
	switch r {
	case FinancialTransactionEventsTypeACHOriginationCancelled, FinancialTransactionEventsTypeACHOriginationInitiated, FinancialTransactionEventsTypeACHOriginationProcessed, FinancialTransactionEventsTypeACHOriginationReleased, FinancialTransactionEventsTypeACHOriginationReviewed, FinancialTransactionEventsTypeACHOriginationSettled, FinancialTransactionEventsTypeACHReceiptProcessed, FinancialTransactionEventsTypeACHReceiptSettled, FinancialTransactionEventsTypeACHReturnInitiated, FinancialTransactionEventsTypeACHReturnProcessed, FinancialTransactionEventsTypeAuthorization, FinancialTransactionEventsTypeAuthorizationAdvice, FinancialTransactionEventsTypeAuthorizationExpiry, FinancialTransactionEventsTypeAuthorizationReversal, FinancialTransactionEventsTypeBalanceInquiry, FinancialTransactionEventsTypeBillingError, FinancialTransactionEventsTypeCashBack, FinancialTransactionEventsTypeClearing, FinancialTransactionEventsTypeCorrectionCredit, FinancialTransactionEventsTypeCorrectionDebit, FinancialTransactionEventsTypeCreditAuthorization, FinancialTransactionEventsTypeCreditAuthorizationAdvice, FinancialTransactionEventsTypeCurrencyConversion, FinancialTransactionEventsTypeDisputeWon, FinancialTransactionEventsTypeExternalACHCanceled, FinancialTransactionEventsTypeExternalACHInitiated, FinancialTransactionEventsTypeExternalACHReleased, FinancialTransactionEventsTypeExternalACHReversed, FinancialTransactionEventsTypeExternalACHSettled, FinancialTransactionEventsTypeExternalCheckCanceled, FinancialTransactionEventsTypeExternalCheckInitiated, FinancialTransactionEventsTypeExternalCheckReleased, FinancialTransactionEventsTypeExternalCheckReversed, FinancialTransactionEventsTypeExternalCheckSettled, FinancialTransactionEventsTypeExternalTransferCanceled, FinancialTransactionEventsTypeExternalTransferInitiated, FinancialTransactionEventsTypeExternalTransferReleased, FinancialTransactionEventsTypeExternalTransferReversed, FinancialTransactionEventsTypeExternalTransferSettled, FinancialTransactionEventsTypeExternalWireCanceled, FinancialTransactionEventsTypeExternalWireInitiated, FinancialTransactionEventsTypeExternalWireReleased, FinancialTransactionEventsTypeExternalWireReversed, FinancialTransactionEventsTypeExternalWireSettled, FinancialTransactionEventsTypeFinancialAuthorization, FinancialTransactionEventsTypeFinancialCreditAuthorization, FinancialTransactionEventsTypeInterest, FinancialTransactionEventsTypeLatePayment, FinancialTransactionEventsTypeProvisionalCredit, FinancialTransactionEventsTypeReturn, FinancialTransactionEventsTypeReturnReversal, FinancialTransactionEventsTypeTransfer, FinancialTransactionEventsTypeTransferInsufficientFunds:
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
