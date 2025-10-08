// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
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
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%s", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/financial_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a financial account
func (r *FinancialAccountService) Get(ctx context.Context, financialAccountToken string, opts ...option.RequestOption) (res *FinancialAccount, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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

// Register account number
func (r *FinancialAccountService) RegisterAccountNumber(ctx context.Context, financialAccountToken string, body FinancialAccountRegisterAccountNumberParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/register_account_number", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Update financial account status
func (r *FinancialAccountService) UpdateStatus(ctx context.Context, financialAccountToken string, body FinancialAccountUpdateStatusParams, opts ...option.RequestOption) (res *FinancialAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/update_status", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FinancialAccount struct {
	// Globally unique identifier for the account
	Token               string                              `json:"token,required" format:"uuid"`
	AccountToken        string                              `json:"account_token,required,nullable" format:"uuid"`
	Created             time.Time                           `json:"created,required" format:"date-time"`
	CreditConfiguration FinancialAccountCreditConfiguration `json:"credit_configuration,required,nullable"`
	// Whether financial account is for the benefit of another entity
	IsForBenefitOf bool   `json:"is_for_benefit_of,required"`
	Nickname       string `json:"nickname,required,nullable"`
	// Status of the financial account
	Status        FinancialAccountStatus `json:"status,required"`
	Type          FinancialAccountType   `json:"type,required"`
	Updated       time.Time              `json:"updated,required" format:"date-time"`
	AccountNumber string                 `json:"account_number,nullable"`
	RoutingNumber string                 `json:"routing_number,nullable"`
	// Substatus for the financial account
	Substatus FinancialAccountSubstatus `json:"substatus,nullable"`
	JSON      financialAccountJSON      `json:"-"`
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
	Status              apijson.Field
	Type                apijson.Field
	Updated             apijson.Field
	AccountNumber       apijson.Field
	RoutingNumber       apijson.Field
	Substatus           apijson.Field
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
	// Reason for the financial account being marked as Charged Off
	ChargedOffReason FinancialAccountCreditConfigurationChargedOffReason `json:"charged_off_reason,required,nullable"`
	CreditLimit      int64                                               `json:"credit_limit,required,nullable"`
	// Globally unique identifier for the credit product
	CreditProductToken       string `json:"credit_product_token,required,nullable"`
	ExternalBankAccountToken string `json:"external_bank_account_token,required,nullable" format:"uuid"`
	// State of the financial account
	FinancialAccountState FinancialAccountCreditConfigurationFinancialAccountState `json:"financial_account_state,required,nullable"`
	IsSpendBlocked        bool                                                     `json:"is_spend_blocked,required"`
	// Tier assigned to the financial account
	Tier                        string                                                         `json:"tier,required,nullable"`
	AutoCollectionConfiguration FinancialAccountCreditConfigurationAutoCollectionConfiguration `json:"auto_collection_configuration"`
	JSON                        financialAccountCreditConfigurationJSON                        `json:"-"`
}

// financialAccountCreditConfigurationJSON contains the JSON metadata for the
// struct [FinancialAccountCreditConfiguration]
type financialAccountCreditConfigurationJSON struct {
	ChargedOffReason            apijson.Field
	CreditLimit                 apijson.Field
	CreditProductToken          apijson.Field
	ExternalBankAccountToken    apijson.Field
	FinancialAccountState       apijson.Field
	IsSpendBlocked              apijson.Field
	Tier                        apijson.Field
	AutoCollectionConfiguration apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *FinancialAccountCreditConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountCreditConfigurationJSON) RawJSON() string {
	return r.raw
}

// Reason for the financial account being marked as Charged Off
type FinancialAccountCreditConfigurationChargedOffReason string

const (
	FinancialAccountCreditConfigurationChargedOffReasonDelinquent FinancialAccountCreditConfigurationChargedOffReason = "DELINQUENT"
	FinancialAccountCreditConfigurationChargedOffReasonFraud      FinancialAccountCreditConfigurationChargedOffReason = "FRAUD"
)

func (r FinancialAccountCreditConfigurationChargedOffReason) IsKnown() bool {
	switch r {
	case FinancialAccountCreditConfigurationChargedOffReasonDelinquent, FinancialAccountCreditConfigurationChargedOffReasonFraud:
		return true
	}
	return false
}

// State of the financial account
type FinancialAccountCreditConfigurationFinancialAccountState string

const (
	FinancialAccountCreditConfigurationFinancialAccountStatePending    FinancialAccountCreditConfigurationFinancialAccountState = "PENDING"
	FinancialAccountCreditConfigurationFinancialAccountStateCurrent    FinancialAccountCreditConfigurationFinancialAccountState = "CURRENT"
	FinancialAccountCreditConfigurationFinancialAccountStateDelinquent FinancialAccountCreditConfigurationFinancialAccountState = "DELINQUENT"
	FinancialAccountCreditConfigurationFinancialAccountStateChargedOff FinancialAccountCreditConfigurationFinancialAccountState = "CHARGED_OFF"
)

func (r FinancialAccountCreditConfigurationFinancialAccountState) IsKnown() bool {
	switch r {
	case FinancialAccountCreditConfigurationFinancialAccountStatePending, FinancialAccountCreditConfigurationFinancialAccountStateCurrent, FinancialAccountCreditConfigurationFinancialAccountStateDelinquent, FinancialAccountCreditConfigurationFinancialAccountStateChargedOff:
		return true
	}
	return false
}

type FinancialAccountCreditConfigurationAutoCollectionConfiguration struct {
	// If auto collection is enabled for this account
	AutoCollectionEnabled bool                                                               `json:"auto_collection_enabled,required"`
	JSON                  financialAccountCreditConfigurationAutoCollectionConfigurationJSON `json:"-"`
}

// financialAccountCreditConfigurationAutoCollectionConfigurationJSON contains the
// JSON metadata for the struct
// [FinancialAccountCreditConfigurationAutoCollectionConfiguration]
type financialAccountCreditConfigurationAutoCollectionConfigurationJSON struct {
	AutoCollectionEnabled apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *FinancialAccountCreditConfigurationAutoCollectionConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountCreditConfigurationAutoCollectionConfigurationJSON) RawJSON() string {
	return r.raw
}

// Status of the financial account
type FinancialAccountStatus string

const (
	FinancialAccountStatusOpen      FinancialAccountStatus = "OPEN"
	FinancialAccountStatusClosed    FinancialAccountStatus = "CLOSED"
	FinancialAccountStatusSuspended FinancialAccountStatus = "SUSPENDED"
	FinancialAccountStatusPending   FinancialAccountStatus = "PENDING"
)

func (r FinancialAccountStatus) IsKnown() bool {
	switch r {
	case FinancialAccountStatusOpen, FinancialAccountStatusClosed, FinancialAccountStatusSuspended, FinancialAccountStatusPending:
		return true
	}
	return false
}

type FinancialAccountType string

const (
	FinancialAccountTypeIssuing                    FinancialAccountType = "ISSUING"
	FinancialAccountTypeReserve                    FinancialAccountType = "RESERVE"
	FinancialAccountTypeOperating                  FinancialAccountType = "OPERATING"
	FinancialAccountTypeChargedOffFees             FinancialAccountType = "CHARGED_OFF_FEES"
	FinancialAccountTypeChargedOffInterest         FinancialAccountType = "CHARGED_OFF_INTEREST"
	FinancialAccountTypeChargedOffPrincipal        FinancialAccountType = "CHARGED_OFF_PRINCIPAL"
	FinancialAccountTypeSecurity                   FinancialAccountType = "SECURITY"
	FinancialAccountTypeProgramReceivables         FinancialAccountType = "PROGRAM_RECEIVABLES"
	FinancialAccountTypeCollection                 FinancialAccountType = "COLLECTION"
	FinancialAccountTypeProgramBankAccountsPayable FinancialAccountType = "PROGRAM_BANK_ACCOUNTS_PAYABLE"
)

func (r FinancialAccountType) IsKnown() bool {
	switch r {
	case FinancialAccountTypeIssuing, FinancialAccountTypeReserve, FinancialAccountTypeOperating, FinancialAccountTypeChargedOffFees, FinancialAccountTypeChargedOffInterest, FinancialAccountTypeChargedOffPrincipal, FinancialAccountTypeSecurity, FinancialAccountTypeProgramReceivables, FinancialAccountTypeCollection, FinancialAccountTypeProgramBankAccountsPayable:
		return true
	}
	return false
}

// Substatus for the financial account
type FinancialAccountSubstatus string

const (
	FinancialAccountSubstatusChargedOffDelinquent FinancialAccountSubstatus = "CHARGED_OFF_DELINQUENT"
	FinancialAccountSubstatusChargedOffFraud      FinancialAccountSubstatus = "CHARGED_OFF_FRAUD"
	FinancialAccountSubstatusEndUserRequest       FinancialAccountSubstatus = "END_USER_REQUEST"
	FinancialAccountSubstatusBankRequest          FinancialAccountSubstatus = "BANK_REQUEST"
	FinancialAccountSubstatusDelinquent           FinancialAccountSubstatus = "DELINQUENT"
)

func (r FinancialAccountSubstatus) IsKnown() bool {
	switch r {
	case FinancialAccountSubstatusChargedOffDelinquent, FinancialAccountSubstatusChargedOffFraud, FinancialAccountSubstatusEndUserRequest, FinancialAccountSubstatusBankRequest, FinancialAccountSubstatusDelinquent:
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
	//   - `INTERNAL` - Transaction for internal adjustment.
	//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
	//     program.
	Category FinancialTransactionCategory `json:"category,required"`
	// Date and time when the financial transaction first occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-character alphabetic ISO 4217 code for the settling currency of the
	// transaction.
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
//   - `INTERNAL` - Transaction for internal adjustment.
//   - `TRANSFER` - Internal transfer of funds between financial accounts in your
//     program.
type FinancialTransactionCategory string

const (
	FinancialTransactionCategoryACH      FinancialTransactionCategory = "ACH"
	FinancialTransactionCategoryCard     FinancialTransactionCategory = "CARD"
	FinancialTransactionCategoryInternal FinancialTransactionCategory = "INTERNAL"
	FinancialTransactionCategoryTransfer FinancialTransactionCategory = "TRANSFER"
)

func (r FinancialTransactionCategory) IsKnown() bool {
	switch r {
	case FinancialTransactionCategoryACH, FinancialTransactionCategoryCard, FinancialTransactionCategoryInternal, FinancialTransactionCategoryTransfer:
		return true
	}
	return false
}

// Financial Event
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
	FinancialTransactionEventsTypeACHOriginationRejected       FinancialTransactionEventsType = "ACH_ORIGINATION_REJECTED"
	FinancialTransactionEventsTypeACHOriginationReviewed       FinancialTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	FinancialTransactionEventsTypeACHOriginationSettled        FinancialTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	FinancialTransactionEventsTypeACHReceiptProcessed          FinancialTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	FinancialTransactionEventsTypeACHReceiptSettled            FinancialTransactionEventsType = "ACH_RECEIPT_SETTLED"
	FinancialTransactionEventsTypeACHReturnInitiated           FinancialTransactionEventsType = "ACH_RETURN_INITIATED"
	FinancialTransactionEventsTypeACHReturnProcessed           FinancialTransactionEventsType = "ACH_RETURN_PROCESSED"
	FinancialTransactionEventsTypeACHReturnRejected            FinancialTransactionEventsType = "ACH_RETURN_REJECTED"
	FinancialTransactionEventsTypeACHReturnSettled             FinancialTransactionEventsType = "ACH_RETURN_SETTLED"
	FinancialTransactionEventsTypeAuthorization                FinancialTransactionEventsType = "AUTHORIZATION"
	FinancialTransactionEventsTypeAuthorizationAdvice          FinancialTransactionEventsType = "AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeAuthorizationExpiry          FinancialTransactionEventsType = "AUTHORIZATION_EXPIRY"
	FinancialTransactionEventsTypeAuthorizationReversal        FinancialTransactionEventsType = "AUTHORIZATION_REVERSAL"
	FinancialTransactionEventsTypeBalanceInquiry               FinancialTransactionEventsType = "BALANCE_INQUIRY"
	FinancialTransactionEventsTypeBillingError                 FinancialTransactionEventsType = "BILLING_ERROR"
	FinancialTransactionEventsTypeBillingErrorReversal         FinancialTransactionEventsType = "BILLING_ERROR_REVERSAL"
	FinancialTransactionEventsTypeCardToCard                   FinancialTransactionEventsType = "CARD_TO_CARD"
	FinancialTransactionEventsTypeCashBack                     FinancialTransactionEventsType = "CASH_BACK"
	FinancialTransactionEventsTypeCashBackReversal             FinancialTransactionEventsType = "CASH_BACK_REVERSAL"
	FinancialTransactionEventsTypeClearing                     FinancialTransactionEventsType = "CLEARING"
	FinancialTransactionEventsTypeCollection                   FinancialTransactionEventsType = "COLLECTION"
	FinancialTransactionEventsTypeCorrectionCredit             FinancialTransactionEventsType = "CORRECTION_CREDIT"
	FinancialTransactionEventsTypeCorrectionDebit              FinancialTransactionEventsType = "CORRECTION_DEBIT"
	FinancialTransactionEventsTypeCreditAuthorization          FinancialTransactionEventsType = "CREDIT_AUTHORIZATION"
	FinancialTransactionEventsTypeCreditAuthorizationAdvice    FinancialTransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	FinancialTransactionEventsTypeCurrencyConversion           FinancialTransactionEventsType = "CURRENCY_CONVERSION"
	FinancialTransactionEventsTypeCurrencyConversionReversal   FinancialTransactionEventsType = "CURRENCY_CONVERSION_REVERSAL"
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
	FinancialTransactionEventsTypeInterestReversal             FinancialTransactionEventsType = "INTEREST_REVERSAL"
	FinancialTransactionEventsTypeInternalAdjustment           FinancialTransactionEventsType = "INTERNAL_ADJUSTMENT"
	FinancialTransactionEventsTypeLatePayment                  FinancialTransactionEventsType = "LATE_PAYMENT"
	FinancialTransactionEventsTypeLatePaymentReversal          FinancialTransactionEventsType = "LATE_PAYMENT_REVERSAL"
	FinancialTransactionEventsTypeLossWriteOff                 FinancialTransactionEventsType = "LOSS_WRITE_OFF"
	FinancialTransactionEventsTypeProvisionalCredit            FinancialTransactionEventsType = "PROVISIONAL_CREDIT"
	FinancialTransactionEventsTypeProvisionalCreditReversal    FinancialTransactionEventsType = "PROVISIONAL_CREDIT_REVERSAL"
	FinancialTransactionEventsTypeService                      FinancialTransactionEventsType = "SERVICE"
	FinancialTransactionEventsTypeReturn                       FinancialTransactionEventsType = "RETURN"
	FinancialTransactionEventsTypeReturnReversal               FinancialTransactionEventsType = "RETURN_REVERSAL"
	FinancialTransactionEventsTypeTransfer                     FinancialTransactionEventsType = "TRANSFER"
	FinancialTransactionEventsTypeTransferInsufficientFunds    FinancialTransactionEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
	FinancialTransactionEventsTypeReturnedPayment              FinancialTransactionEventsType = "RETURNED_PAYMENT"
	FinancialTransactionEventsTypeReturnedPaymentReversal      FinancialTransactionEventsType = "RETURNED_PAYMENT_REVERSAL"
	FinancialTransactionEventsTypeLithicNetworkPayment         FinancialTransactionEventsType = "LITHIC_NETWORK_PAYMENT"
)

func (r FinancialTransactionEventsType) IsKnown() bool {
	switch r {
	case FinancialTransactionEventsTypeACHOriginationCancelled, FinancialTransactionEventsTypeACHOriginationInitiated, FinancialTransactionEventsTypeACHOriginationProcessed, FinancialTransactionEventsTypeACHOriginationReleased, FinancialTransactionEventsTypeACHOriginationRejected, FinancialTransactionEventsTypeACHOriginationReviewed, FinancialTransactionEventsTypeACHOriginationSettled, FinancialTransactionEventsTypeACHReceiptProcessed, FinancialTransactionEventsTypeACHReceiptSettled, FinancialTransactionEventsTypeACHReturnInitiated, FinancialTransactionEventsTypeACHReturnProcessed, FinancialTransactionEventsTypeACHReturnRejected, FinancialTransactionEventsTypeACHReturnSettled, FinancialTransactionEventsTypeAuthorization, FinancialTransactionEventsTypeAuthorizationAdvice, FinancialTransactionEventsTypeAuthorizationExpiry, FinancialTransactionEventsTypeAuthorizationReversal, FinancialTransactionEventsTypeBalanceInquiry, FinancialTransactionEventsTypeBillingError, FinancialTransactionEventsTypeBillingErrorReversal, FinancialTransactionEventsTypeCardToCard, FinancialTransactionEventsTypeCashBack, FinancialTransactionEventsTypeCashBackReversal, FinancialTransactionEventsTypeClearing, FinancialTransactionEventsTypeCollection, FinancialTransactionEventsTypeCorrectionCredit, FinancialTransactionEventsTypeCorrectionDebit, FinancialTransactionEventsTypeCreditAuthorization, FinancialTransactionEventsTypeCreditAuthorizationAdvice, FinancialTransactionEventsTypeCurrencyConversion, FinancialTransactionEventsTypeCurrencyConversionReversal, FinancialTransactionEventsTypeDisputeWon, FinancialTransactionEventsTypeExternalACHCanceled, FinancialTransactionEventsTypeExternalACHInitiated, FinancialTransactionEventsTypeExternalACHReleased, FinancialTransactionEventsTypeExternalACHReversed, FinancialTransactionEventsTypeExternalACHSettled, FinancialTransactionEventsTypeExternalCheckCanceled, FinancialTransactionEventsTypeExternalCheckInitiated, FinancialTransactionEventsTypeExternalCheckReleased, FinancialTransactionEventsTypeExternalCheckReversed, FinancialTransactionEventsTypeExternalCheckSettled, FinancialTransactionEventsTypeExternalTransferCanceled, FinancialTransactionEventsTypeExternalTransferInitiated, FinancialTransactionEventsTypeExternalTransferReleased, FinancialTransactionEventsTypeExternalTransferReversed, FinancialTransactionEventsTypeExternalTransferSettled, FinancialTransactionEventsTypeExternalWireCanceled, FinancialTransactionEventsTypeExternalWireInitiated, FinancialTransactionEventsTypeExternalWireReleased, FinancialTransactionEventsTypeExternalWireReversed, FinancialTransactionEventsTypeExternalWireSettled, FinancialTransactionEventsTypeFinancialAuthorization, FinancialTransactionEventsTypeFinancialCreditAuthorization, FinancialTransactionEventsTypeInterest, FinancialTransactionEventsTypeInterestReversal, FinancialTransactionEventsTypeInternalAdjustment, FinancialTransactionEventsTypeLatePayment, FinancialTransactionEventsTypeLatePaymentReversal, FinancialTransactionEventsTypeLossWriteOff, FinancialTransactionEventsTypeProvisionalCredit, FinancialTransactionEventsTypeProvisionalCreditReversal, FinancialTransactionEventsTypeService, FinancialTransactionEventsTypeReturn, FinancialTransactionEventsTypeReturnReversal, FinancialTransactionEventsTypeTransfer, FinancialTransactionEventsTypeTransferInsufficientFunds, FinancialTransactionEventsTypeReturnedPayment, FinancialTransactionEventsTypeReturnedPaymentReversal, FinancialTransactionEventsTypeLithicNetworkPayment:
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
	FinancialAccountListParamsTypeSecurity  FinancialAccountListParamsType = "SECURITY"
)

func (r FinancialAccountListParamsType) IsKnown() bool {
	switch r {
	case FinancialAccountListParamsTypeIssuing, FinancialAccountListParamsTypeOperating, FinancialAccountListParamsTypeReserve, FinancialAccountListParamsTypeSecurity:
		return true
	}
	return false
}

type FinancialAccountRegisterAccountNumberParams struct {
	AccountNumber param.Field[string] `json:"account_number,required"`
}

func (r FinancialAccountRegisterAccountNumberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FinancialAccountUpdateStatusParams struct {
	// Status of the financial account
	Status param.Field[FinancialAccountUpdateStatusParamsStatus] `json:"status,required"`
	// Substatus for the financial account
	Substatus param.Field[FinancialAccountUpdateStatusParamsSubstatus] `json:"substatus,required"`
}

func (r FinancialAccountUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of the financial account
type FinancialAccountUpdateStatusParamsStatus string

const (
	FinancialAccountUpdateStatusParamsStatusOpen      FinancialAccountUpdateStatusParamsStatus = "OPEN"
	FinancialAccountUpdateStatusParamsStatusClosed    FinancialAccountUpdateStatusParamsStatus = "CLOSED"
	FinancialAccountUpdateStatusParamsStatusSuspended FinancialAccountUpdateStatusParamsStatus = "SUSPENDED"
	FinancialAccountUpdateStatusParamsStatusPending   FinancialAccountUpdateStatusParamsStatus = "PENDING"
)

func (r FinancialAccountUpdateStatusParamsStatus) IsKnown() bool {
	switch r {
	case FinancialAccountUpdateStatusParamsStatusOpen, FinancialAccountUpdateStatusParamsStatusClosed, FinancialAccountUpdateStatusParamsStatusSuspended, FinancialAccountUpdateStatusParamsStatusPending:
		return true
	}
	return false
}

// Substatus for the financial account
type FinancialAccountUpdateStatusParamsSubstatus string

const (
	FinancialAccountUpdateStatusParamsSubstatusChargedOffFraud      FinancialAccountUpdateStatusParamsSubstatus = "CHARGED_OFF_FRAUD"
	FinancialAccountUpdateStatusParamsSubstatusEndUserRequest       FinancialAccountUpdateStatusParamsSubstatus = "END_USER_REQUEST"
	FinancialAccountUpdateStatusParamsSubstatusBankRequest          FinancialAccountUpdateStatusParamsSubstatus = "BANK_REQUEST"
	FinancialAccountUpdateStatusParamsSubstatusChargedOffDelinquent FinancialAccountUpdateStatusParamsSubstatus = "CHARGED_OFF_DELINQUENT"
)

func (r FinancialAccountUpdateStatusParamsSubstatus) IsKnown() bool {
	switch r {
	case FinancialAccountUpdateStatusParamsSubstatusChargedOffFraud, FinancialAccountUpdateStatusParamsSubstatusEndUserRequest, FinancialAccountUpdateStatusParamsSubstatusBankRequest, FinancialAccountUpdateStatusParamsSubstatusChargedOffDelinquent:
		return true
	}
	return false
}
