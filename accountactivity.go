// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/tidwall/gjson"
)

// AccountActivityService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountActivityService] method instead.
type AccountActivityService struct {
	Options []option.RequestOption
}

// NewAccountActivityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountActivityService(opts ...option.RequestOption) (r *AccountActivityService) {
	r = &AccountActivityService{}
	r.Options = opts
	return
}

// Retrieve a list of transactions across all public accounts.
func (r *AccountActivityService) List(ctx context.Context, query AccountActivityListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AccountActivityListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/account_activity"
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

// Retrieve a list of transactions across all public accounts.
func (r *AccountActivityService) ListAutoPaging(ctx context.Context, query AccountActivityListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AccountActivityListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a single transaction
func (r *AccountActivityService) GetTransaction(ctx context.Context, transactionToken string, opts ...option.RequestOption) (res *AccountActivityGetTransactionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionToken == "" {
		err = errors.New("missing required transaction_token parameter")
		return
	}
	path := fmt.Sprintf("v1/account_activity/%s", transactionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WirePartyDetails struct {
	// Account number
	AccountNumber string `json:"account_number,nullable"`
	// Routing number or BIC of the financial institution
	AgentID string `json:"agent_id,nullable"`
	// Name of the financial institution
	AgentName string `json:"agent_name,nullable"`
	// Name of the person or company
	Name string               `json:"name,nullable"`
	JSON wirePartyDetailsJSON `json:"-"`
}

// wirePartyDetailsJSON contains the JSON metadata for the struct
// [WirePartyDetails]
type wirePartyDetailsJSON struct {
	AccountNumber apijson.Field
	AgentID       apijson.Field
	AgentName     apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WirePartyDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wirePartyDetailsJSON) RawJSON() string {
	return r.raw
}

// Response containing multiple transaction types. The `family` field determines
// which transaction type is returned: INTERNAL returns FinancialTransaction,
// TRANSFER returns BookTransferTransaction, CARD returns CardTransaction, PAYMENT
// returns PaymentTransaction, EXTERNAL_PAYMENT returns ExternalPaymentResponse,
// and MANAGEMENT_OPERATION returns ManagementOperationTransaction
type AccountActivityListResponse struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// The status of the transaction
	Status AccountActivityListResponseStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// The token for the account associated with this transaction.
	AccountToken string `json:"account_token" format:"uuid"`
	// Fee assessed by the merchant and paid for by the cardholder in the smallest unit
	// of the currency. Will be zero if no fee is assessed. Rebates may be transmitted
	// as a negative value to indicate credited fees.
	AcquirerFee int64 `json:"acquirer_fee,nullable"`
	// Unique identifier assigned to a transaction by the acquirer that can be used in
	// dispute and chargeback filing. This field has been deprecated in favor of the
	// `acquirer_reference_number` that resides in the event-level `network_info`.
	//
	// Deprecated: deprecated
	AcquirerReferenceNumber string `json:"acquirer_reference_number,nullable"`
	// When the transaction is pending, this represents the authorization amount of the
	// transaction in the anticipated settlement currency. Once the transaction has
	// settled, this field represents the settled amount in the settlement currency.
	//
	// Deprecated: deprecated
	Amount int64 `json:"amount"`
	// This field can have the runtime type of [TransactionAmounts].
	Amounts interface{} `json:"amounts"`
	// The authorization amount of the transaction in the anticipated settlement
	// currency.
	//
	// Deprecated: deprecated
	AuthorizationAmount int64 `json:"authorization_amount,nullable"`
	// A fixed-width 6-digit numeric identifier that can be used to identify a
	// transaction with networks.
	AuthorizationCode string `json:"authorization_code,nullable"`
	// This field can have the runtime type of [TransactionAvs].
	Avs interface{} `json:"avs"`
	// Token for the card used in this transaction.
	CardToken string `json:"card_token" format:"uuid"`
	// This field can have the runtime type of [TransactionCardholderAuthentication].
	CardholderAuthentication interface{} `json:"cardholder_authentication"`
	// Transaction category
	Category AccountActivityListResponseCategory `json:"category"`
	// Currency of the transaction, represented in ISO 4217 format
	Currency string `json:"currency"`
	// Transaction descriptor
	Descriptor string `json:"descriptor"`
	// Transfer direction
	Direction AccountActivityListResponseDirection `json:"direction"`
	// This field can have the runtime type of
	// [[]AccountActivityListResponseFinancialTransactionEvent],
	// [[]BookTransferResponseEvent], [[]TransactionEvent], [[]PaymentEvent],
	// [[]ExternalPaymentEvent], [[]ManagementOperationTransactionEvent].
	Events interface{} `json:"events"`
	// Expected release date for the transaction
	ExpectedReleaseDate time.Time `json:"expected_release_date,nullable" format:"date"`
	// External bank account token
	ExternalBankAccountToken string `json:"external_bank_account_token,nullable" format:"uuid"`
	// External ID defined by the customer
	ExternalID string `json:"external_id,nullable"`
	// External resource associated with the management operation
	ExternalResource ExternalResource `json:"external_resource,nullable"`
	// INTERNAL - Financial Transaction
	Family AccountActivityListResponseFamily `json:"family"`
	// Financial account token associated with the transaction
	FinancialAccountToken string `json:"financial_account_token,nullable" format:"uuid"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case
	FromFinancialAccountToken string `json:"from_financial_account_token" format:"uuid"`
	// This field can have the runtime type of [TransactionMerchant].
	Merchant interface{} `json:"merchant"`
	// Analogous to the 'amount', but in the merchant currency.
	//
	// Deprecated: deprecated
	MerchantAmount int64 `json:"merchant_amount,nullable"`
	// Analogous to the 'authorization_amount', but in the merchant currency.
	//
	// Deprecated: deprecated
	MerchantAuthorizationAmount int64 `json:"merchant_authorization_amount,nullable"`
	// 3-character alphabetic ISO 4217 code for the local currency of the transaction.
	//
	// Deprecated: deprecated
	MerchantCurrency string `json:"merchant_currency"`
	// Transfer method
	Method AccountActivityListResponseMethod `json:"method"`
	// This field can have the runtime type of [PaymentMethodAttributes].
	MethodAttributes interface{} `json:"method_attributes"`
	// Card network of the authorization. Value is `UNKNOWN` when Lithic cannot
	// determine the network code from the upstream provider.
	Network AccountActivityListResponseNetwork `json:"network,nullable"`
	// Network-provided score assessing risk level associated with a given
	// authorization. Scores are on a range of 0-999, with 0 representing the lowest
	// risk and 999 representing the highest risk. For Visa transactions, where the raw
	// score has a range of 0-99, Lithic will normalize the score by multiplying the
	// raw score by 10x.
	NetworkRiskScore int64                                  `json:"network_risk_score,nullable"`
	PaymentType      AccountActivityListResponsePaymentType `json:"payment_type"`
	// Pending amount in cents
	PendingAmount int64 `json:"pending_amount"`
	// This field can have the runtime type of [TransactionPos].
	Pos interface{} `json:"pos"`
	// This field can have the runtime type of [PaymentRelatedAccountTokens].
	RelatedAccountTokens interface{} `json:"related_account_tokens"`
	// Transaction result
	Result AccountActivityListResponseResult `json:"result"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount"`
	// Transaction source
	Source AccountActivityListResponseSource `json:"source"`
	// Globally unique identifier for the financial account or card that will receive
	// the funds. Accepted type dependent on the program's use case
	ToFinancialAccountToken string `json:"to_financial_account_token" format:"uuid"`
	// This field can have the runtime type of [TransactionTokenInfo].
	TokenInfo interface{} `json:"token_info"`
	// This field can have the runtime type of [BookTransferResponseTransactionSeries],
	// [ManagementOperationTransactionTransactionSeries].
	TransactionSeries interface{}                     `json:"transaction_series"`
	Type              AccountActivityListResponseType `json:"type"`
	// User-defined identifier
	UserDefinedID string                          `json:"user_defined_id,nullable"`
	JSON          accountActivityListResponseJSON `json:"-"`
	union         AccountActivityListResponseUnion
}

// accountActivityListResponseJSON contains the JSON metadata for the struct
// [AccountActivityListResponse]
type accountActivityListResponseJSON struct {
	Token                       apijson.Field
	Created                     apijson.Field
	Status                      apijson.Field
	Updated                     apijson.Field
	AccountToken                apijson.Field
	AcquirerFee                 apijson.Field
	AcquirerReferenceNumber     apijson.Field
	Amount                      apijson.Field
	Amounts                     apijson.Field
	AuthorizationAmount         apijson.Field
	AuthorizationCode           apijson.Field
	Avs                         apijson.Field
	CardToken                   apijson.Field
	CardholderAuthentication    apijson.Field
	Category                    apijson.Field
	Currency                    apijson.Field
	Descriptor                  apijson.Field
	Direction                   apijson.Field
	Events                      apijson.Field
	ExpectedReleaseDate         apijson.Field
	ExternalBankAccountToken    apijson.Field
	ExternalID                  apijson.Field
	ExternalResource            apijson.Field
	Family                      apijson.Field
	FinancialAccountToken       apijson.Field
	FromFinancialAccountToken   apijson.Field
	Merchant                    apijson.Field
	MerchantAmount              apijson.Field
	MerchantAuthorizationAmount apijson.Field
	MerchantCurrency            apijson.Field
	Method                      apijson.Field
	MethodAttributes            apijson.Field
	Network                     apijson.Field
	NetworkRiskScore            apijson.Field
	PaymentType                 apijson.Field
	PendingAmount               apijson.Field
	Pos                         apijson.Field
	RelatedAccountTokens        apijson.Field
	Result                      apijson.Field
	SettledAmount               apijson.Field
	Source                      apijson.Field
	ToFinancialAccountToken     apijson.Field
	TokenInfo                   apijson.Field
	TransactionSeries           apijson.Field
	Type                        apijson.Field
	UserDefinedID               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r accountActivityListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccountActivityListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccountActivityListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountActivityListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountActivityListResponseFinancialTransaction], [BookTransferResponse],
// [AccountActivityListResponseCardTransaction], [Payment], [ExternalPayment],
// [ManagementOperationTransaction].
func (r AccountActivityListResponse) AsUnion() AccountActivityListResponseUnion {
	return r.union
}

// Response containing multiple transaction types. The `family` field determines
// which transaction type is returned: INTERNAL returns FinancialTransaction,
// TRANSFER returns BookTransferTransaction, CARD returns CardTransaction, PAYMENT
// returns PaymentTransaction, EXTERNAL_PAYMENT returns ExternalPaymentResponse,
// and MANAGEMENT_OPERATION returns ManagementOperationTransaction
//
// Union satisfied by [AccountActivityListResponseFinancialTransaction],
// [BookTransferResponse], [AccountActivityListResponseCardTransaction], [Payment],
// [ExternalPayment] or [ManagementOperationTransaction].
type AccountActivityListResponseUnion interface {
	implementsAccountActivityListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountActivityListResponseUnion)(nil)).Elem(),
		"family",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseFinancialTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BookTransferResponse{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseCardTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(Payment{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
	)
}

// Financial transaction with inheritance from unified base transaction
type AccountActivityListResponseFinancialTransaction struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// Transaction category
	Category AccountActivityListResponseFinancialTransactionCategory `json:"category,required"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Currency of the transaction, represented in ISO 4217 format
	Currency string `json:"currency,required"`
	// Transaction descriptor
	Descriptor string `json:"descriptor,required"`
	// List of transaction events
	Events []AccountActivityListResponseFinancialTransactionEvent `json:"events,required"`
	// INTERNAL - Financial Transaction
	Family AccountActivityListResponseFinancialTransactionFamily `json:"family,required"`
	// Financial account token associated with the transaction
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Pending amount in cents
	PendingAmount int64 `json:"pending_amount,required"`
	// Transaction result
	Result AccountActivityListResponseFinancialTransactionResult `json:"result,required"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount,required"`
	// The status of the transaction
	Status AccountActivityListResponseFinancialTransactionStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time                                           `json:"updated,required" format:"date-time"`
	JSON    accountActivityListResponseFinancialTransactionJSON `json:"-"`
}

// accountActivityListResponseFinancialTransactionJSON contains the JSON metadata
// for the struct [AccountActivityListResponseFinancialTransaction]
type accountActivityListResponseFinancialTransactionJSON struct {
	Token                 apijson.Field
	Category              apijson.Field
	Created               apijson.Field
	Currency              apijson.Field
	Descriptor            apijson.Field
	Events                apijson.Field
	Family                apijson.Field
	FinancialAccountToken apijson.Field
	PendingAmount         apijson.Field
	Result                apijson.Field
	SettledAmount         apijson.Field
	Status                apijson.Field
	Updated               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountActivityListResponseFinancialTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponseFinancialTransactionJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityListResponseFinancialTransaction) implementsAccountActivityListResponse() {}

// Transaction category
type AccountActivityListResponseFinancialTransactionCategory string

const (
	AccountActivityListResponseFinancialTransactionCategoryACH                    AccountActivityListResponseFinancialTransactionCategory = "ACH"
	AccountActivityListResponseFinancialTransactionCategoryBalanceOrFunding       AccountActivityListResponseFinancialTransactionCategory = "BALANCE_OR_FUNDING"
	AccountActivityListResponseFinancialTransactionCategoryFee                    AccountActivityListResponseFinancialTransactionCategory = "FEE"
	AccountActivityListResponseFinancialTransactionCategoryReward                 AccountActivityListResponseFinancialTransactionCategory = "REWARD"
	AccountActivityListResponseFinancialTransactionCategoryAdjustment             AccountActivityListResponseFinancialTransactionCategory = "ADJUSTMENT"
	AccountActivityListResponseFinancialTransactionCategoryDerecognition          AccountActivityListResponseFinancialTransactionCategory = "DERECOGNITION"
	AccountActivityListResponseFinancialTransactionCategoryDispute                AccountActivityListResponseFinancialTransactionCategory = "DISPUTE"
	AccountActivityListResponseFinancialTransactionCategoryCard                   AccountActivityListResponseFinancialTransactionCategory = "CARD"
	AccountActivityListResponseFinancialTransactionCategoryExternalACH            AccountActivityListResponseFinancialTransactionCategory = "EXTERNAL_ACH"
	AccountActivityListResponseFinancialTransactionCategoryExternalCheck          AccountActivityListResponseFinancialTransactionCategory = "EXTERNAL_CHECK"
	AccountActivityListResponseFinancialTransactionCategoryExternalTransfer       AccountActivityListResponseFinancialTransactionCategory = "EXTERNAL_TRANSFER"
	AccountActivityListResponseFinancialTransactionCategoryExternalWire           AccountActivityListResponseFinancialTransactionCategory = "EXTERNAL_WIRE"
	AccountActivityListResponseFinancialTransactionCategoryManagementAdjustment   AccountActivityListResponseFinancialTransactionCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityListResponseFinancialTransactionCategoryManagementDispute      AccountActivityListResponseFinancialTransactionCategory = "MANAGEMENT_DISPUTE"
	AccountActivityListResponseFinancialTransactionCategoryManagementFee          AccountActivityListResponseFinancialTransactionCategory = "MANAGEMENT_FEE"
	AccountActivityListResponseFinancialTransactionCategoryManagementReward       AccountActivityListResponseFinancialTransactionCategory = "MANAGEMENT_REWARD"
	AccountActivityListResponseFinancialTransactionCategoryManagementDisbursement AccountActivityListResponseFinancialTransactionCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityListResponseFinancialTransactionCategoryProgramFunding         AccountActivityListResponseFinancialTransactionCategory = "PROGRAM_FUNDING"
)

func (r AccountActivityListResponseFinancialTransactionCategory) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFinancialTransactionCategoryACH, AccountActivityListResponseFinancialTransactionCategoryBalanceOrFunding, AccountActivityListResponseFinancialTransactionCategoryFee, AccountActivityListResponseFinancialTransactionCategoryReward, AccountActivityListResponseFinancialTransactionCategoryAdjustment, AccountActivityListResponseFinancialTransactionCategoryDerecognition, AccountActivityListResponseFinancialTransactionCategoryDispute, AccountActivityListResponseFinancialTransactionCategoryCard, AccountActivityListResponseFinancialTransactionCategoryExternalACH, AccountActivityListResponseFinancialTransactionCategoryExternalCheck, AccountActivityListResponseFinancialTransactionCategoryExternalTransfer, AccountActivityListResponseFinancialTransactionCategoryExternalWire, AccountActivityListResponseFinancialTransactionCategoryManagementAdjustment, AccountActivityListResponseFinancialTransactionCategoryManagementDispute, AccountActivityListResponseFinancialTransactionCategoryManagementFee, AccountActivityListResponseFinancialTransactionCategoryManagementReward, AccountActivityListResponseFinancialTransactionCategoryManagementDisbursement, AccountActivityListResponseFinancialTransactionCategoryProgramFunding:
		return true
	}
	return false
}

// Financial Event
type AccountActivityListResponseFinancialTransactionEvent struct {
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result AccountActivityListResponseFinancialTransactionEventsResult `json:"result"`
	Type   AccountActivityListResponseFinancialTransactionEventsType   `json:"type"`
	JSON   accountActivityListResponseFinancialTransactionEventJSON    `json:"-"`
}

// accountActivityListResponseFinancialTransactionEventJSON contains the JSON
// metadata for the struct [AccountActivityListResponseFinancialTransactionEvent]
type accountActivityListResponseFinancialTransactionEventJSON struct {
	Token       apijson.Field
	Amount      apijson.Field
	Created     apijson.Field
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountActivityListResponseFinancialTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponseFinancialTransactionEventJSON) RawJSON() string {
	return r.raw
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type AccountActivityListResponseFinancialTransactionEventsResult string

const (
	AccountActivityListResponseFinancialTransactionEventsResultApproved AccountActivityListResponseFinancialTransactionEventsResult = "APPROVED"
	AccountActivityListResponseFinancialTransactionEventsResultDeclined AccountActivityListResponseFinancialTransactionEventsResult = "DECLINED"
)

func (r AccountActivityListResponseFinancialTransactionEventsResult) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFinancialTransactionEventsResultApproved, AccountActivityListResponseFinancialTransactionEventsResultDeclined:
		return true
	}
	return false
}

type AccountActivityListResponseFinancialTransactionEventsType string

const (
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationCancelled      AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_CANCELLED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationInitiated      AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_INITIATED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationProcessed      AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_PROCESSED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationReleased       AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_RELEASED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationRejected       AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_REJECTED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationReviewed       AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationSettled        AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptProcessed          AccountActivityListResponseFinancialTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptReleased           AccountActivityListResponseFinancialTransactionEventsType = "ACH_RECEIPT_RELEASED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptSettled            AccountActivityListResponseFinancialTransactionEventsType = "ACH_RECEIPT_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReturnInitiated           AccountActivityListResponseFinancialTransactionEventsType = "ACH_RETURN_INITIATED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReturnProcessed           AccountActivityListResponseFinancialTransactionEventsType = "ACH_RETURN_PROCESSED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReturnRejected            AccountActivityListResponseFinancialTransactionEventsType = "ACH_RETURN_REJECTED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReturnSettled             AccountActivityListResponseFinancialTransactionEventsType = "ACH_RETURN_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeAuthorization                AccountActivityListResponseFinancialTransactionEventsType = "AUTHORIZATION"
	AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationAdvice          AccountActivityListResponseFinancialTransactionEventsType = "AUTHORIZATION_ADVICE"
	AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationExpiry          AccountActivityListResponseFinancialTransactionEventsType = "AUTHORIZATION_EXPIRY"
	AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationReversal        AccountActivityListResponseFinancialTransactionEventsType = "AUTHORIZATION_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeBalanceInquiry               AccountActivityListResponseFinancialTransactionEventsType = "BALANCE_INQUIRY"
	AccountActivityListResponseFinancialTransactionEventsTypeBillingError                 AccountActivityListResponseFinancialTransactionEventsType = "BILLING_ERROR"
	AccountActivityListResponseFinancialTransactionEventsTypeBillingErrorReversal         AccountActivityListResponseFinancialTransactionEventsType = "BILLING_ERROR_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeCardToCard                   AccountActivityListResponseFinancialTransactionEventsType = "CARD_TO_CARD"
	AccountActivityListResponseFinancialTransactionEventsTypeCashBack                     AccountActivityListResponseFinancialTransactionEventsType = "CASH_BACK"
	AccountActivityListResponseFinancialTransactionEventsTypeCashBackReversal             AccountActivityListResponseFinancialTransactionEventsType = "CASH_BACK_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeClearing                     AccountActivityListResponseFinancialTransactionEventsType = "CLEARING"
	AccountActivityListResponseFinancialTransactionEventsTypeCollection                   AccountActivityListResponseFinancialTransactionEventsType = "COLLECTION"
	AccountActivityListResponseFinancialTransactionEventsTypeCorrectionCredit             AccountActivityListResponseFinancialTransactionEventsType = "CORRECTION_CREDIT"
	AccountActivityListResponseFinancialTransactionEventsTypeCorrectionDebit              AccountActivityListResponseFinancialTransactionEventsType = "CORRECTION_DEBIT"
	AccountActivityListResponseFinancialTransactionEventsTypeCreditAuthorization          AccountActivityListResponseFinancialTransactionEventsType = "CREDIT_AUTHORIZATION"
	AccountActivityListResponseFinancialTransactionEventsTypeCreditAuthorizationAdvice    AccountActivityListResponseFinancialTransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	AccountActivityListResponseFinancialTransactionEventsTypeCurrencyConversion           AccountActivityListResponseFinancialTransactionEventsType = "CURRENCY_CONVERSION"
	AccountActivityListResponseFinancialTransactionEventsTypeCurrencyConversionReversal   AccountActivityListResponseFinancialTransactionEventsType = "CURRENCY_CONVERSION_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeDisputeWon                   AccountActivityListResponseFinancialTransactionEventsType = "DISPUTE_WON"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalACHCanceled          AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_ACH_CANCELED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalACHInitiated         AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_ACH_INITIATED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalACHReleased          AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_ACH_RELEASED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalACHReversed          AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_ACH_REVERSED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalACHSettled           AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_ACH_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckCanceled        AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_CANCELED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckInitiated       AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_INITIATED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckReleased        AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_RELEASED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckReversed        AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_REVERSED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckSettled         AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferCanceled     AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_CANCELED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferInitiated    AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_INITIATED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferReleased     AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_RELEASED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferReversed     AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_REVERSED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferSettled      AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalWireCanceled         AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_CANCELED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalWireInitiated        AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_INITIATED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalWireReleased         AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_RELEASED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalWireReversed         AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_REVERSED"
	AccountActivityListResponseFinancialTransactionEventsTypeExternalWireSettled          AccountActivityListResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeFinancialAuthorization       AccountActivityListResponseFinancialTransactionEventsType = "FINANCIAL_AUTHORIZATION"
	AccountActivityListResponseFinancialTransactionEventsTypeFinancialCreditAuthorization AccountActivityListResponseFinancialTransactionEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	AccountActivityListResponseFinancialTransactionEventsTypeInterest                     AccountActivityListResponseFinancialTransactionEventsType = "INTEREST"
	AccountActivityListResponseFinancialTransactionEventsTypeInterestReversal             AccountActivityListResponseFinancialTransactionEventsType = "INTEREST_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeInternalAdjustment           AccountActivityListResponseFinancialTransactionEventsType = "INTERNAL_ADJUSTMENT"
	AccountActivityListResponseFinancialTransactionEventsTypeLatePayment                  AccountActivityListResponseFinancialTransactionEventsType = "LATE_PAYMENT"
	AccountActivityListResponseFinancialTransactionEventsTypeLatePaymentReversal          AccountActivityListResponseFinancialTransactionEventsType = "LATE_PAYMENT_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeLossWriteOff                 AccountActivityListResponseFinancialTransactionEventsType = "LOSS_WRITE_OFF"
	AccountActivityListResponseFinancialTransactionEventsTypeProvisionalCredit            AccountActivityListResponseFinancialTransactionEventsType = "PROVISIONAL_CREDIT"
	AccountActivityListResponseFinancialTransactionEventsTypeProvisionalCreditReversal    AccountActivityListResponseFinancialTransactionEventsType = "PROVISIONAL_CREDIT_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeService                      AccountActivityListResponseFinancialTransactionEventsType = "SERVICE"
	AccountActivityListResponseFinancialTransactionEventsTypeReturn                       AccountActivityListResponseFinancialTransactionEventsType = "RETURN"
	AccountActivityListResponseFinancialTransactionEventsTypeReturnReversal               AccountActivityListResponseFinancialTransactionEventsType = "RETURN_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeTransfer                     AccountActivityListResponseFinancialTransactionEventsType = "TRANSFER"
	AccountActivityListResponseFinancialTransactionEventsTypeTransferInsufficientFunds    AccountActivityListResponseFinancialTransactionEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
	AccountActivityListResponseFinancialTransactionEventsTypeReturnedPayment              AccountActivityListResponseFinancialTransactionEventsType = "RETURNED_PAYMENT"
	AccountActivityListResponseFinancialTransactionEventsTypeReturnedPaymentReversal      AccountActivityListResponseFinancialTransactionEventsType = "RETURNED_PAYMENT_REVERSAL"
	AccountActivityListResponseFinancialTransactionEventsTypeLithicNetworkPayment         AccountActivityListResponseFinancialTransactionEventsType = "LITHIC_NETWORK_PAYMENT"
)

func (r AccountActivityListResponseFinancialTransactionEventsType) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationCancelled, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationInitiated, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationProcessed, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationReleased, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationRejected, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationReviewed, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationSettled, AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptProcessed, AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptReleased, AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptSettled, AccountActivityListResponseFinancialTransactionEventsTypeACHReturnInitiated, AccountActivityListResponseFinancialTransactionEventsTypeACHReturnProcessed, AccountActivityListResponseFinancialTransactionEventsTypeACHReturnRejected, AccountActivityListResponseFinancialTransactionEventsTypeACHReturnSettled, AccountActivityListResponseFinancialTransactionEventsTypeAuthorization, AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationAdvice, AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationExpiry, AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationReversal, AccountActivityListResponseFinancialTransactionEventsTypeBalanceInquiry, AccountActivityListResponseFinancialTransactionEventsTypeBillingError, AccountActivityListResponseFinancialTransactionEventsTypeBillingErrorReversal, AccountActivityListResponseFinancialTransactionEventsTypeCardToCard, AccountActivityListResponseFinancialTransactionEventsTypeCashBack, AccountActivityListResponseFinancialTransactionEventsTypeCashBackReversal, AccountActivityListResponseFinancialTransactionEventsTypeClearing, AccountActivityListResponseFinancialTransactionEventsTypeCollection, AccountActivityListResponseFinancialTransactionEventsTypeCorrectionCredit, AccountActivityListResponseFinancialTransactionEventsTypeCorrectionDebit, AccountActivityListResponseFinancialTransactionEventsTypeCreditAuthorization, AccountActivityListResponseFinancialTransactionEventsTypeCreditAuthorizationAdvice, AccountActivityListResponseFinancialTransactionEventsTypeCurrencyConversion, AccountActivityListResponseFinancialTransactionEventsTypeCurrencyConversionReversal, AccountActivityListResponseFinancialTransactionEventsTypeDisputeWon, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHCanceled, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHInitiated, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHReleased, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHReversed, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHSettled, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckCanceled, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckInitiated, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckReleased, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckReversed, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckSettled, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferCanceled, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferInitiated, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferReleased, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferReversed, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferSettled, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireCanceled, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireInitiated, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireReleased, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireReversed, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireSettled, AccountActivityListResponseFinancialTransactionEventsTypeFinancialAuthorization, AccountActivityListResponseFinancialTransactionEventsTypeFinancialCreditAuthorization, AccountActivityListResponseFinancialTransactionEventsTypeInterest, AccountActivityListResponseFinancialTransactionEventsTypeInterestReversal, AccountActivityListResponseFinancialTransactionEventsTypeInternalAdjustment, AccountActivityListResponseFinancialTransactionEventsTypeLatePayment, AccountActivityListResponseFinancialTransactionEventsTypeLatePaymentReversal, AccountActivityListResponseFinancialTransactionEventsTypeLossWriteOff, AccountActivityListResponseFinancialTransactionEventsTypeProvisionalCredit, AccountActivityListResponseFinancialTransactionEventsTypeProvisionalCreditReversal, AccountActivityListResponseFinancialTransactionEventsTypeService, AccountActivityListResponseFinancialTransactionEventsTypeReturn, AccountActivityListResponseFinancialTransactionEventsTypeReturnReversal, AccountActivityListResponseFinancialTransactionEventsTypeTransfer, AccountActivityListResponseFinancialTransactionEventsTypeTransferInsufficientFunds, AccountActivityListResponseFinancialTransactionEventsTypeReturnedPayment, AccountActivityListResponseFinancialTransactionEventsTypeReturnedPaymentReversal, AccountActivityListResponseFinancialTransactionEventsTypeLithicNetworkPayment:
		return true
	}
	return false
}

// INTERNAL - Financial Transaction
type AccountActivityListResponseFinancialTransactionFamily string

const (
	AccountActivityListResponseFinancialTransactionFamilyInternal AccountActivityListResponseFinancialTransactionFamily = "INTERNAL"
)

func (r AccountActivityListResponseFinancialTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFinancialTransactionFamilyInternal:
		return true
	}
	return false
}

// Transaction result
type AccountActivityListResponseFinancialTransactionResult string

const (
	AccountActivityListResponseFinancialTransactionResultApproved AccountActivityListResponseFinancialTransactionResult = "APPROVED"
	AccountActivityListResponseFinancialTransactionResultDeclined AccountActivityListResponseFinancialTransactionResult = "DECLINED"
)

func (r AccountActivityListResponseFinancialTransactionResult) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFinancialTransactionResultApproved, AccountActivityListResponseFinancialTransactionResultDeclined:
		return true
	}
	return false
}

// The status of the transaction
type AccountActivityListResponseFinancialTransactionStatus string

const (
	AccountActivityListResponseFinancialTransactionStatusPending  AccountActivityListResponseFinancialTransactionStatus = "PENDING"
	AccountActivityListResponseFinancialTransactionStatusSettled  AccountActivityListResponseFinancialTransactionStatus = "SETTLED"
	AccountActivityListResponseFinancialTransactionStatusDeclined AccountActivityListResponseFinancialTransactionStatus = "DECLINED"
	AccountActivityListResponseFinancialTransactionStatusReversed AccountActivityListResponseFinancialTransactionStatus = "REVERSED"
	AccountActivityListResponseFinancialTransactionStatusCanceled AccountActivityListResponseFinancialTransactionStatus = "CANCELED"
)

func (r AccountActivityListResponseFinancialTransactionStatus) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFinancialTransactionStatusPending, AccountActivityListResponseFinancialTransactionStatusSettled, AccountActivityListResponseFinancialTransactionStatusDeclined, AccountActivityListResponseFinancialTransactionStatusReversed, AccountActivityListResponseFinancialTransactionStatusCanceled:
		return true
	}
	return false
}

// Card transaction with ledger base properties
type AccountActivityListResponseCardTransaction struct {
	JSON accountActivityListResponseCardTransactionJSON `json:"-"`
	Transaction
}

// accountActivityListResponseCardTransactionJSON contains the JSON metadata for
// the struct [AccountActivityListResponseCardTransaction]
type accountActivityListResponseCardTransactionJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountActivityListResponseCardTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponseCardTransactionJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityListResponseCardTransaction) implementsAccountActivityListResponse() {}

// The status of the transaction
type AccountActivityListResponseStatus string

const (
	AccountActivityListResponseStatusPending  AccountActivityListResponseStatus = "PENDING"
	AccountActivityListResponseStatusSettled  AccountActivityListResponseStatus = "SETTLED"
	AccountActivityListResponseStatusDeclined AccountActivityListResponseStatus = "DECLINED"
	AccountActivityListResponseStatusReversed AccountActivityListResponseStatus = "REVERSED"
	AccountActivityListResponseStatusCanceled AccountActivityListResponseStatus = "CANCELED"
	AccountActivityListResponseStatusExpired  AccountActivityListResponseStatus = "EXPIRED"
	AccountActivityListResponseStatusVoided   AccountActivityListResponseStatus = "VOIDED"
)

func (r AccountActivityListResponseStatus) IsKnown() bool {
	switch r {
	case AccountActivityListResponseStatusPending, AccountActivityListResponseStatusSettled, AccountActivityListResponseStatusDeclined, AccountActivityListResponseStatusReversed, AccountActivityListResponseStatusCanceled, AccountActivityListResponseStatusExpired, AccountActivityListResponseStatusVoided:
		return true
	}
	return false
}

// Transaction category
type AccountActivityListResponseCategory string

const (
	AccountActivityListResponseCategoryACH                    AccountActivityListResponseCategory = "ACH"
	AccountActivityListResponseCategoryBalanceOrFunding       AccountActivityListResponseCategory = "BALANCE_OR_FUNDING"
	AccountActivityListResponseCategoryFee                    AccountActivityListResponseCategory = "FEE"
	AccountActivityListResponseCategoryReward                 AccountActivityListResponseCategory = "REWARD"
	AccountActivityListResponseCategoryAdjustment             AccountActivityListResponseCategory = "ADJUSTMENT"
	AccountActivityListResponseCategoryDerecognition          AccountActivityListResponseCategory = "DERECOGNITION"
	AccountActivityListResponseCategoryDispute                AccountActivityListResponseCategory = "DISPUTE"
	AccountActivityListResponseCategoryCard                   AccountActivityListResponseCategory = "CARD"
	AccountActivityListResponseCategoryExternalACH            AccountActivityListResponseCategory = "EXTERNAL_ACH"
	AccountActivityListResponseCategoryExternalCheck          AccountActivityListResponseCategory = "EXTERNAL_CHECK"
	AccountActivityListResponseCategoryExternalTransfer       AccountActivityListResponseCategory = "EXTERNAL_TRANSFER"
	AccountActivityListResponseCategoryExternalWire           AccountActivityListResponseCategory = "EXTERNAL_WIRE"
	AccountActivityListResponseCategoryManagementAdjustment   AccountActivityListResponseCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityListResponseCategoryManagementDispute      AccountActivityListResponseCategory = "MANAGEMENT_DISPUTE"
	AccountActivityListResponseCategoryManagementFee          AccountActivityListResponseCategory = "MANAGEMENT_FEE"
	AccountActivityListResponseCategoryManagementReward       AccountActivityListResponseCategory = "MANAGEMENT_REWARD"
	AccountActivityListResponseCategoryManagementDisbursement AccountActivityListResponseCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityListResponseCategoryProgramFunding         AccountActivityListResponseCategory = "PROGRAM_FUNDING"
	AccountActivityListResponseCategoryInternal               AccountActivityListResponseCategory = "INTERNAL"
	AccountActivityListResponseCategoryTransfer               AccountActivityListResponseCategory = "TRANSFER"
)

func (r AccountActivityListResponseCategory) IsKnown() bool {
	switch r {
	case AccountActivityListResponseCategoryACH, AccountActivityListResponseCategoryBalanceOrFunding, AccountActivityListResponseCategoryFee, AccountActivityListResponseCategoryReward, AccountActivityListResponseCategoryAdjustment, AccountActivityListResponseCategoryDerecognition, AccountActivityListResponseCategoryDispute, AccountActivityListResponseCategoryCard, AccountActivityListResponseCategoryExternalACH, AccountActivityListResponseCategoryExternalCheck, AccountActivityListResponseCategoryExternalTransfer, AccountActivityListResponseCategoryExternalWire, AccountActivityListResponseCategoryManagementAdjustment, AccountActivityListResponseCategoryManagementDispute, AccountActivityListResponseCategoryManagementFee, AccountActivityListResponseCategoryManagementReward, AccountActivityListResponseCategoryManagementDisbursement, AccountActivityListResponseCategoryProgramFunding, AccountActivityListResponseCategoryInternal, AccountActivityListResponseCategoryTransfer:
		return true
	}
	return false
}

// Transfer direction
type AccountActivityListResponseDirection string

const (
	AccountActivityListResponseDirectionCredit AccountActivityListResponseDirection = "CREDIT"
	AccountActivityListResponseDirectionDebit  AccountActivityListResponseDirection = "DEBIT"
)

func (r AccountActivityListResponseDirection) IsKnown() bool {
	switch r {
	case AccountActivityListResponseDirectionCredit, AccountActivityListResponseDirectionDebit:
		return true
	}
	return false
}

// INTERNAL - Financial Transaction
type AccountActivityListResponseFamily string

const (
	AccountActivityListResponseFamilyInternal            AccountActivityListResponseFamily = "INTERNAL"
	AccountActivityListResponseFamilyTransfer            AccountActivityListResponseFamily = "TRANSFER"
	AccountActivityListResponseFamilyCard                AccountActivityListResponseFamily = "CARD"
	AccountActivityListResponseFamilyPayment             AccountActivityListResponseFamily = "PAYMENT"
	AccountActivityListResponseFamilyExternalPayment     AccountActivityListResponseFamily = "EXTERNAL_PAYMENT"
	AccountActivityListResponseFamilyManagementOperation AccountActivityListResponseFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityListResponseFamily) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFamilyInternal, AccountActivityListResponseFamilyTransfer, AccountActivityListResponseFamilyCard, AccountActivityListResponseFamilyPayment, AccountActivityListResponseFamilyExternalPayment, AccountActivityListResponseFamilyManagementOperation:
		return true
	}
	return false
}

// Transfer method
type AccountActivityListResponseMethod string

const (
	AccountActivityListResponseMethodACHNextDay AccountActivityListResponseMethod = "ACH_NEXT_DAY"
	AccountActivityListResponseMethodACHSameDay AccountActivityListResponseMethod = "ACH_SAME_DAY"
	AccountActivityListResponseMethodWire       AccountActivityListResponseMethod = "WIRE"
)

func (r AccountActivityListResponseMethod) IsKnown() bool {
	switch r {
	case AccountActivityListResponseMethodACHNextDay, AccountActivityListResponseMethodACHSameDay, AccountActivityListResponseMethodWire:
		return true
	}
	return false
}

// Card network of the authorization. Value is `UNKNOWN` when Lithic cannot
// determine the network code from the upstream provider.
type AccountActivityListResponseNetwork string

const (
	AccountActivityListResponseNetworkAmex       AccountActivityListResponseNetwork = "AMEX"
	AccountActivityListResponseNetworkInterlink  AccountActivityListResponseNetwork = "INTERLINK"
	AccountActivityListResponseNetworkMaestro    AccountActivityListResponseNetwork = "MAESTRO"
	AccountActivityListResponseNetworkMastercard AccountActivityListResponseNetwork = "MASTERCARD"
	AccountActivityListResponseNetworkUnknown    AccountActivityListResponseNetwork = "UNKNOWN"
	AccountActivityListResponseNetworkVisa       AccountActivityListResponseNetwork = "VISA"
)

func (r AccountActivityListResponseNetwork) IsKnown() bool {
	switch r {
	case AccountActivityListResponseNetworkAmex, AccountActivityListResponseNetworkInterlink, AccountActivityListResponseNetworkMaestro, AccountActivityListResponseNetworkMastercard, AccountActivityListResponseNetworkUnknown, AccountActivityListResponseNetworkVisa:
		return true
	}
	return false
}

type AccountActivityListResponsePaymentType string

const (
	AccountActivityListResponsePaymentTypeDeposit    AccountActivityListResponsePaymentType = "DEPOSIT"
	AccountActivityListResponsePaymentTypeWithdrawal AccountActivityListResponsePaymentType = "WITHDRAWAL"
)

func (r AccountActivityListResponsePaymentType) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTypeDeposit, AccountActivityListResponsePaymentTypeWithdrawal:
		return true
	}
	return false
}

// Transaction result
type AccountActivityListResponseResult string

const (
	AccountActivityListResponseResultApproved                    AccountActivityListResponseResult = "APPROVED"
	AccountActivityListResponseResultDeclined                    AccountActivityListResponseResult = "DECLINED"
	AccountActivityListResponseResultAccountPaused               AccountActivityListResponseResult = "ACCOUNT_PAUSED"
	AccountActivityListResponseResultAccountStateTransactionFail AccountActivityListResponseResult = "ACCOUNT_STATE_TRANSACTION_FAIL"
	AccountActivityListResponseResultBankConnectionError         AccountActivityListResponseResult = "BANK_CONNECTION_ERROR"
	AccountActivityListResponseResultBankNotVerified             AccountActivityListResponseResult = "BANK_NOT_VERIFIED"
	AccountActivityListResponseResultCardClosed                  AccountActivityListResponseResult = "CARD_CLOSED"
	AccountActivityListResponseResultCardPaused                  AccountActivityListResponseResult = "CARD_PAUSED"
	AccountActivityListResponseResultFraudAdvice                 AccountActivityListResponseResult = "FRAUD_ADVICE"
	AccountActivityListResponseResultIgnoredTtlExpiry            AccountActivityListResponseResult = "IGNORED_TTL_EXPIRY"
	AccountActivityListResponseResultSuspectedFraud              AccountActivityListResponseResult = "SUSPECTED_FRAUD"
	AccountActivityListResponseResultInactiveAccount             AccountActivityListResponseResult = "INACTIVE_ACCOUNT"
	AccountActivityListResponseResultIncorrectPin                AccountActivityListResponseResult = "INCORRECT_PIN"
	AccountActivityListResponseResultInvalidCardDetails          AccountActivityListResponseResult = "INVALID_CARD_DETAILS"
	AccountActivityListResponseResultInsufficientFunds           AccountActivityListResponseResult = "INSUFFICIENT_FUNDS"
	AccountActivityListResponseResultInsufficientFundsPreload    AccountActivityListResponseResult = "INSUFFICIENT_FUNDS_PRELOAD"
	AccountActivityListResponseResultInvalidTransaction          AccountActivityListResponseResult = "INVALID_TRANSACTION"
	AccountActivityListResponseResultMerchantBlacklist           AccountActivityListResponseResult = "MERCHANT_BLACKLIST"
	AccountActivityListResponseResultOriginalNotFound            AccountActivityListResponseResult = "ORIGINAL_NOT_FOUND"
	AccountActivityListResponseResultPreviouslyCompleted         AccountActivityListResponseResult = "PREVIOUSLY_COMPLETED"
	AccountActivityListResponseResultSingleUseRecharged          AccountActivityListResponseResult = "SINGLE_USE_RECHARGED"
	AccountActivityListResponseResultSwitchInoperativeAdvice     AccountActivityListResponseResult = "SWITCH_INOPERATIVE_ADVICE"
	AccountActivityListResponseResultUnauthorizedMerchant        AccountActivityListResponseResult = "UNAUTHORIZED_MERCHANT"
	AccountActivityListResponseResultUnknownHostTimeout          AccountActivityListResponseResult = "UNKNOWN_HOST_TIMEOUT"
	AccountActivityListResponseResultUserTransactionLimit        AccountActivityListResponseResult = "USER_TRANSACTION_LIMIT"
)

func (r AccountActivityListResponseResult) IsKnown() bool {
	switch r {
	case AccountActivityListResponseResultApproved, AccountActivityListResponseResultDeclined, AccountActivityListResponseResultAccountPaused, AccountActivityListResponseResultAccountStateTransactionFail, AccountActivityListResponseResultBankConnectionError, AccountActivityListResponseResultBankNotVerified, AccountActivityListResponseResultCardClosed, AccountActivityListResponseResultCardPaused, AccountActivityListResponseResultFraudAdvice, AccountActivityListResponseResultIgnoredTtlExpiry, AccountActivityListResponseResultSuspectedFraud, AccountActivityListResponseResultInactiveAccount, AccountActivityListResponseResultIncorrectPin, AccountActivityListResponseResultInvalidCardDetails, AccountActivityListResponseResultInsufficientFunds, AccountActivityListResponseResultInsufficientFundsPreload, AccountActivityListResponseResultInvalidTransaction, AccountActivityListResponseResultMerchantBlacklist, AccountActivityListResponseResultOriginalNotFound, AccountActivityListResponseResultPreviouslyCompleted, AccountActivityListResponseResultSingleUseRecharged, AccountActivityListResponseResultSwitchInoperativeAdvice, AccountActivityListResponseResultUnauthorizedMerchant, AccountActivityListResponseResultUnknownHostTimeout, AccountActivityListResponseResultUserTransactionLimit:
		return true
	}
	return false
}

// Transaction source
type AccountActivityListResponseSource string

const (
	AccountActivityListResponseSourceLithic   AccountActivityListResponseSource = "LITHIC"
	AccountActivityListResponseSourceExternal AccountActivityListResponseSource = "EXTERNAL"
	AccountActivityListResponseSourceCustomer AccountActivityListResponseSource = "CUSTOMER"
)

func (r AccountActivityListResponseSource) IsKnown() bool {
	switch r {
	case AccountActivityListResponseSourceLithic, AccountActivityListResponseSourceExternal, AccountActivityListResponseSourceCustomer:
		return true
	}
	return false
}

type AccountActivityListResponseType string

const (
	AccountActivityListResponseTypeOriginationCredit   AccountActivityListResponseType = "ORIGINATION_CREDIT"
	AccountActivityListResponseTypeOriginationDebit    AccountActivityListResponseType = "ORIGINATION_DEBIT"
	AccountActivityListResponseTypeReceiptCredit       AccountActivityListResponseType = "RECEIPT_CREDIT"
	AccountActivityListResponseTypeReceiptDebit        AccountActivityListResponseType = "RECEIPT_DEBIT"
	AccountActivityListResponseTypeWireInboundPayment  AccountActivityListResponseType = "WIRE_INBOUND_PAYMENT"
	AccountActivityListResponseTypeWireInboundAdmin    AccountActivityListResponseType = "WIRE_INBOUND_ADMIN"
	AccountActivityListResponseTypeWireOutboundPayment AccountActivityListResponseType = "WIRE_OUTBOUND_PAYMENT"
	AccountActivityListResponseTypeWireOutboundAdmin   AccountActivityListResponseType = "WIRE_OUTBOUND_ADMIN"
)

func (r AccountActivityListResponseType) IsKnown() bool {
	switch r {
	case AccountActivityListResponseTypeOriginationCredit, AccountActivityListResponseTypeOriginationDebit, AccountActivityListResponseTypeReceiptCredit, AccountActivityListResponseTypeReceiptDebit, AccountActivityListResponseTypeWireInboundPayment, AccountActivityListResponseTypeWireInboundAdmin, AccountActivityListResponseTypeWireOutboundPayment, AccountActivityListResponseTypeWireOutboundAdmin:
		return true
	}
	return false
}

// Response containing multiple transaction types. The `family` field determines
// which transaction type is returned: INTERNAL returns FinancialTransaction,
// TRANSFER returns BookTransferTransaction, CARD returns CardTransaction, PAYMENT
// returns PaymentTransaction, EXTERNAL_PAYMENT returns ExternalPaymentResponse,
// and MANAGEMENT_OPERATION returns ManagementOperationTransaction
type AccountActivityGetTransactionResponse struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// The status of the transaction
	Status AccountActivityGetTransactionResponseStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// The token for the account associated with this transaction.
	AccountToken string `json:"account_token" format:"uuid"`
	// Fee assessed by the merchant and paid for by the cardholder in the smallest unit
	// of the currency. Will be zero if no fee is assessed. Rebates may be transmitted
	// as a negative value to indicate credited fees.
	AcquirerFee int64 `json:"acquirer_fee,nullable"`
	// Unique identifier assigned to a transaction by the acquirer that can be used in
	// dispute and chargeback filing. This field has been deprecated in favor of the
	// `acquirer_reference_number` that resides in the event-level `network_info`.
	//
	// Deprecated: deprecated
	AcquirerReferenceNumber string `json:"acquirer_reference_number,nullable"`
	// When the transaction is pending, this represents the authorization amount of the
	// transaction in the anticipated settlement currency. Once the transaction has
	// settled, this field represents the settled amount in the settlement currency.
	//
	// Deprecated: deprecated
	Amount int64 `json:"amount"`
	// This field can have the runtime type of [TransactionAmounts].
	Amounts interface{} `json:"amounts"`
	// The authorization amount of the transaction in the anticipated settlement
	// currency.
	//
	// Deprecated: deprecated
	AuthorizationAmount int64 `json:"authorization_amount,nullable"`
	// A fixed-width 6-digit numeric identifier that can be used to identify a
	// transaction with networks.
	AuthorizationCode string `json:"authorization_code,nullable"`
	// This field can have the runtime type of [TransactionAvs].
	Avs interface{} `json:"avs"`
	// Token for the card used in this transaction.
	CardToken string `json:"card_token" format:"uuid"`
	// This field can have the runtime type of [TransactionCardholderAuthentication].
	CardholderAuthentication interface{} `json:"cardholder_authentication"`
	// Transaction category
	Category AccountActivityGetTransactionResponseCategory `json:"category"`
	// Currency of the transaction, represented in ISO 4217 format
	Currency string `json:"currency"`
	// Transaction descriptor
	Descriptor string `json:"descriptor"`
	// Transfer direction
	Direction AccountActivityGetTransactionResponseDirection `json:"direction"`
	// This field can have the runtime type of
	// [[]AccountActivityGetTransactionResponseFinancialTransactionEvent],
	// [[]BookTransferResponseEvent], [[]TransactionEvent], [[]PaymentEvent],
	// [[]ExternalPaymentEvent], [[]ManagementOperationTransactionEvent].
	Events interface{} `json:"events"`
	// Expected release date for the transaction
	ExpectedReleaseDate time.Time `json:"expected_release_date,nullable" format:"date"`
	// External bank account token
	ExternalBankAccountToken string `json:"external_bank_account_token,nullable" format:"uuid"`
	// External ID defined by the customer
	ExternalID string `json:"external_id,nullable"`
	// External resource associated with the management operation
	ExternalResource ExternalResource `json:"external_resource,nullable"`
	// INTERNAL - Financial Transaction
	Family AccountActivityGetTransactionResponseFamily `json:"family"`
	// Financial account token associated with the transaction
	FinancialAccountToken string `json:"financial_account_token,nullable" format:"uuid"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case
	FromFinancialAccountToken string `json:"from_financial_account_token" format:"uuid"`
	// This field can have the runtime type of [TransactionMerchant].
	Merchant interface{} `json:"merchant"`
	// Analogous to the 'amount', but in the merchant currency.
	//
	// Deprecated: deprecated
	MerchantAmount int64 `json:"merchant_amount,nullable"`
	// Analogous to the 'authorization_amount', but in the merchant currency.
	//
	// Deprecated: deprecated
	MerchantAuthorizationAmount int64 `json:"merchant_authorization_amount,nullable"`
	// 3-character alphabetic ISO 4217 code for the local currency of the transaction.
	//
	// Deprecated: deprecated
	MerchantCurrency string `json:"merchant_currency"`
	// Transfer method
	Method AccountActivityGetTransactionResponseMethod `json:"method"`
	// This field can have the runtime type of [PaymentMethodAttributes].
	MethodAttributes interface{} `json:"method_attributes"`
	// Card network of the authorization. Value is `UNKNOWN` when Lithic cannot
	// determine the network code from the upstream provider.
	Network AccountActivityGetTransactionResponseNetwork `json:"network,nullable"`
	// Network-provided score assessing risk level associated with a given
	// authorization. Scores are on a range of 0-999, with 0 representing the lowest
	// risk and 999 representing the highest risk. For Visa transactions, where the raw
	// score has a range of 0-99, Lithic will normalize the score by multiplying the
	// raw score by 10x.
	NetworkRiskScore int64                                            `json:"network_risk_score,nullable"`
	PaymentType      AccountActivityGetTransactionResponsePaymentType `json:"payment_type"`
	// Pending amount in cents
	PendingAmount int64 `json:"pending_amount"`
	// This field can have the runtime type of [TransactionPos].
	Pos interface{} `json:"pos"`
	// This field can have the runtime type of [PaymentRelatedAccountTokens].
	RelatedAccountTokens interface{} `json:"related_account_tokens"`
	// Transaction result
	Result AccountActivityGetTransactionResponseResult `json:"result"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount"`
	// Transaction source
	Source AccountActivityGetTransactionResponseSource `json:"source"`
	// Globally unique identifier for the financial account or card that will receive
	// the funds. Accepted type dependent on the program's use case
	ToFinancialAccountToken string `json:"to_financial_account_token" format:"uuid"`
	// This field can have the runtime type of [TransactionTokenInfo].
	TokenInfo interface{} `json:"token_info"`
	// This field can have the runtime type of [BookTransferResponseTransactionSeries],
	// [ManagementOperationTransactionTransactionSeries].
	TransactionSeries interface{}                               `json:"transaction_series"`
	Type              AccountActivityGetTransactionResponseType `json:"type"`
	// User-defined identifier
	UserDefinedID string                                    `json:"user_defined_id,nullable"`
	JSON          accountActivityGetTransactionResponseJSON `json:"-"`
	union         AccountActivityGetTransactionResponseUnion
}

// accountActivityGetTransactionResponseJSON contains the JSON metadata for the
// struct [AccountActivityGetTransactionResponse]
type accountActivityGetTransactionResponseJSON struct {
	Token                       apijson.Field
	Created                     apijson.Field
	Status                      apijson.Field
	Updated                     apijson.Field
	AccountToken                apijson.Field
	AcquirerFee                 apijson.Field
	AcquirerReferenceNumber     apijson.Field
	Amount                      apijson.Field
	Amounts                     apijson.Field
	AuthorizationAmount         apijson.Field
	AuthorizationCode           apijson.Field
	Avs                         apijson.Field
	CardToken                   apijson.Field
	CardholderAuthentication    apijson.Field
	Category                    apijson.Field
	Currency                    apijson.Field
	Descriptor                  apijson.Field
	Direction                   apijson.Field
	Events                      apijson.Field
	ExpectedReleaseDate         apijson.Field
	ExternalBankAccountToken    apijson.Field
	ExternalID                  apijson.Field
	ExternalResource            apijson.Field
	Family                      apijson.Field
	FinancialAccountToken       apijson.Field
	FromFinancialAccountToken   apijson.Field
	Merchant                    apijson.Field
	MerchantAmount              apijson.Field
	MerchantAuthorizationAmount apijson.Field
	MerchantCurrency            apijson.Field
	Method                      apijson.Field
	MethodAttributes            apijson.Field
	Network                     apijson.Field
	NetworkRiskScore            apijson.Field
	PaymentType                 apijson.Field
	PendingAmount               apijson.Field
	Pos                         apijson.Field
	RelatedAccountTokens        apijson.Field
	Result                      apijson.Field
	SettledAmount               apijson.Field
	Source                      apijson.Field
	ToFinancialAccountToken     apijson.Field
	TokenInfo                   apijson.Field
	TransactionSeries           apijson.Field
	Type                        apijson.Field
	UserDefinedID               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r accountActivityGetTransactionResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccountActivityGetTransactionResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccountActivityGetTransactionResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountActivityGetTransactionResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountActivityGetTransactionResponseFinancialTransaction],
// [BookTransferResponse], [AccountActivityGetTransactionResponseCardTransaction],
// [Payment], [ExternalPayment], [ManagementOperationTransaction].
func (r AccountActivityGetTransactionResponse) AsUnion() AccountActivityGetTransactionResponseUnion {
	return r.union
}

// Response containing multiple transaction types. The `family` field determines
// which transaction type is returned: INTERNAL returns FinancialTransaction,
// TRANSFER returns BookTransferTransaction, CARD returns CardTransaction, PAYMENT
// returns PaymentTransaction, EXTERNAL_PAYMENT returns ExternalPaymentResponse,
// and MANAGEMENT_OPERATION returns ManagementOperationTransaction
//
// Union satisfied by [AccountActivityGetTransactionResponseFinancialTransaction],
// [BookTransferResponse], [AccountActivityGetTransactionResponseCardTransaction],
// [Payment], [ExternalPayment] or [ManagementOperationTransaction].
type AccountActivityGetTransactionResponseUnion interface {
	implementsAccountActivityGetTransactionResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountActivityGetTransactionResponseUnion)(nil)).Elem(),
		"family",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseFinancialTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BookTransferResponse{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseCardTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(Payment{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
	)
}

// Financial transaction with inheritance from unified base transaction
type AccountActivityGetTransactionResponseFinancialTransaction struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// Transaction category
	Category AccountActivityGetTransactionResponseFinancialTransactionCategory `json:"category,required"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Currency of the transaction, represented in ISO 4217 format
	Currency string `json:"currency,required"`
	// Transaction descriptor
	Descriptor string `json:"descriptor,required"`
	// List of transaction events
	Events []AccountActivityGetTransactionResponseFinancialTransactionEvent `json:"events,required"`
	// INTERNAL - Financial Transaction
	Family AccountActivityGetTransactionResponseFinancialTransactionFamily `json:"family,required"`
	// Financial account token associated with the transaction
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Pending amount in cents
	PendingAmount int64 `json:"pending_amount,required"`
	// Transaction result
	Result AccountActivityGetTransactionResponseFinancialTransactionResult `json:"result,required"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount,required"`
	// The status of the transaction
	Status AccountActivityGetTransactionResponseFinancialTransactionStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time                                                     `json:"updated,required" format:"date-time"`
	JSON    accountActivityGetTransactionResponseFinancialTransactionJSON `json:"-"`
}

// accountActivityGetTransactionResponseFinancialTransactionJSON contains the JSON
// metadata for the struct
// [AccountActivityGetTransactionResponseFinancialTransaction]
type accountActivityGetTransactionResponseFinancialTransactionJSON struct {
	Token                 apijson.Field
	Category              apijson.Field
	Created               apijson.Field
	Currency              apijson.Field
	Descriptor            apijson.Field
	Events                apijson.Field
	Family                apijson.Field
	FinancialAccountToken apijson.Field
	PendingAmount         apijson.Field
	Result                apijson.Field
	SettledAmount         apijson.Field
	Status                apijson.Field
	Updated               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponseFinancialTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponseFinancialTransactionJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityGetTransactionResponseFinancialTransaction) implementsAccountActivityGetTransactionResponse() {
}

// Transaction category
type AccountActivityGetTransactionResponseFinancialTransactionCategory string

const (
	AccountActivityGetTransactionResponseFinancialTransactionCategoryACH                    AccountActivityGetTransactionResponseFinancialTransactionCategory = "ACH"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryBalanceOrFunding       AccountActivityGetTransactionResponseFinancialTransactionCategory = "BALANCE_OR_FUNDING"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryFee                    AccountActivityGetTransactionResponseFinancialTransactionCategory = "FEE"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryReward                 AccountActivityGetTransactionResponseFinancialTransactionCategory = "REWARD"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryAdjustment             AccountActivityGetTransactionResponseFinancialTransactionCategory = "ADJUSTMENT"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryDerecognition          AccountActivityGetTransactionResponseFinancialTransactionCategory = "DERECOGNITION"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryDispute                AccountActivityGetTransactionResponseFinancialTransactionCategory = "DISPUTE"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryCard                   AccountActivityGetTransactionResponseFinancialTransactionCategory = "CARD"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalACH            AccountActivityGetTransactionResponseFinancialTransactionCategory = "EXTERNAL_ACH"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalCheck          AccountActivityGetTransactionResponseFinancialTransactionCategory = "EXTERNAL_CHECK"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalTransfer       AccountActivityGetTransactionResponseFinancialTransactionCategory = "EXTERNAL_TRANSFER"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalWire           AccountActivityGetTransactionResponseFinancialTransactionCategory = "EXTERNAL_WIRE"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementAdjustment   AccountActivityGetTransactionResponseFinancialTransactionCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementDispute      AccountActivityGetTransactionResponseFinancialTransactionCategory = "MANAGEMENT_DISPUTE"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementFee          AccountActivityGetTransactionResponseFinancialTransactionCategory = "MANAGEMENT_FEE"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementReward       AccountActivityGetTransactionResponseFinancialTransactionCategory = "MANAGEMENT_REWARD"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementDisbursement AccountActivityGetTransactionResponseFinancialTransactionCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityGetTransactionResponseFinancialTransactionCategoryProgramFunding         AccountActivityGetTransactionResponseFinancialTransactionCategory = "PROGRAM_FUNDING"
)

func (r AccountActivityGetTransactionResponseFinancialTransactionCategory) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFinancialTransactionCategoryACH, AccountActivityGetTransactionResponseFinancialTransactionCategoryBalanceOrFunding, AccountActivityGetTransactionResponseFinancialTransactionCategoryFee, AccountActivityGetTransactionResponseFinancialTransactionCategoryReward, AccountActivityGetTransactionResponseFinancialTransactionCategoryAdjustment, AccountActivityGetTransactionResponseFinancialTransactionCategoryDerecognition, AccountActivityGetTransactionResponseFinancialTransactionCategoryDispute, AccountActivityGetTransactionResponseFinancialTransactionCategoryCard, AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalACH, AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalCheck, AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalTransfer, AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalWire, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementAdjustment, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementDispute, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementFee, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementReward, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementDisbursement, AccountActivityGetTransactionResponseFinancialTransactionCategoryProgramFunding:
		return true
	}
	return false
}

// Financial Event
type AccountActivityGetTransactionResponseFinancialTransactionEvent struct {
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result AccountActivityGetTransactionResponseFinancialTransactionEventsResult `json:"result"`
	Type   AccountActivityGetTransactionResponseFinancialTransactionEventsType   `json:"type"`
	JSON   accountActivityGetTransactionResponseFinancialTransactionEventJSON    `json:"-"`
}

// accountActivityGetTransactionResponseFinancialTransactionEventJSON contains the
// JSON metadata for the struct
// [AccountActivityGetTransactionResponseFinancialTransactionEvent]
type accountActivityGetTransactionResponseFinancialTransactionEventJSON struct {
	Token       apijson.Field
	Amount      apijson.Field
	Created     apijson.Field
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponseFinancialTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponseFinancialTransactionEventJSON) RawJSON() string {
	return r.raw
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type AccountActivityGetTransactionResponseFinancialTransactionEventsResult string

const (
	AccountActivityGetTransactionResponseFinancialTransactionEventsResultApproved AccountActivityGetTransactionResponseFinancialTransactionEventsResult = "APPROVED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsResultDeclined AccountActivityGetTransactionResponseFinancialTransactionEventsResult = "DECLINED"
)

func (r AccountActivityGetTransactionResponseFinancialTransactionEventsResult) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFinancialTransactionEventsResultApproved, AccountActivityGetTransactionResponseFinancialTransactionEventsResultDeclined:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponseFinancialTransactionEventsType string

const (
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationCancelled      AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_CANCELLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationInitiated      AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_INITIATED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationProcessed      AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_PROCESSED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationReleased       AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_RELEASED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationRejected       AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_REJECTED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationReviewed       AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationSettled        AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptProcessed          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptReleased           AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RECEIPT_RELEASED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptSettled            AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RECEIPT_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnInitiated           AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RETURN_INITIATED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnProcessed           AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RETURN_PROCESSED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnRejected            AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RETURN_REJECTED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnSettled             AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RETURN_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorization                AccountActivityGetTransactionResponseFinancialTransactionEventsType = "AUTHORIZATION"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationAdvice          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "AUTHORIZATION_ADVICE"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationExpiry          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "AUTHORIZATION_EXPIRY"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationReversal        AccountActivityGetTransactionResponseFinancialTransactionEventsType = "AUTHORIZATION_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBalanceInquiry               AccountActivityGetTransactionResponseFinancialTransactionEventsType = "BALANCE_INQUIRY"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBillingError                 AccountActivityGetTransactionResponseFinancialTransactionEventsType = "BILLING_ERROR"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBillingErrorReversal         AccountActivityGetTransactionResponseFinancialTransactionEventsType = "BILLING_ERROR_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCardToCard                   AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CARD_TO_CARD"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCashBack                     AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CASH_BACK"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCashBackReversal             AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CASH_BACK_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeClearing                     AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CLEARING"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCollection                   AccountActivityGetTransactionResponseFinancialTransactionEventsType = "COLLECTION"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCorrectionCredit             AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CORRECTION_CREDIT"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCorrectionDebit              AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CORRECTION_DEBIT"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCreditAuthorization          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CREDIT_AUTHORIZATION"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCreditAuthorizationAdvice    AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCurrencyConversion           AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CURRENCY_CONVERSION"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCurrencyConversionReversal   AccountActivityGetTransactionResponseFinancialTransactionEventsType = "CURRENCY_CONVERSION_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeDisputeWon                   AccountActivityGetTransactionResponseFinancialTransactionEventsType = "DISPUTE_WON"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHCanceled          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_ACH_CANCELED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHInitiated         AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_ACH_INITIATED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHReleased          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_ACH_RELEASED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHReversed          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_ACH_REVERSED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHSettled           AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_ACH_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckCanceled        AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_CANCELED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckInitiated       AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_INITIATED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckReleased        AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_RELEASED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckReversed        AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_REVERSED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckSettled         AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_CHECK_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferCanceled     AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_CANCELED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferInitiated    AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_INITIATED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferReleased     AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_RELEASED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferReversed     AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_REVERSED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferSettled      AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_TRANSFER_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireCanceled         AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_CANCELED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireInitiated        AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_INITIATED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireReleased         AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_RELEASED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireReversed         AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_REVERSED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireSettled          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "EXTERNAL_WIRE_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeFinancialAuthorization       AccountActivityGetTransactionResponseFinancialTransactionEventsType = "FINANCIAL_AUTHORIZATION"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeFinancialCreditAuthorization AccountActivityGetTransactionResponseFinancialTransactionEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInterest                     AccountActivityGetTransactionResponseFinancialTransactionEventsType = "INTEREST"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInterestReversal             AccountActivityGetTransactionResponseFinancialTransactionEventsType = "INTEREST_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInternalAdjustment           AccountActivityGetTransactionResponseFinancialTransactionEventsType = "INTERNAL_ADJUSTMENT"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLatePayment                  AccountActivityGetTransactionResponseFinancialTransactionEventsType = "LATE_PAYMENT"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLatePaymentReversal          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "LATE_PAYMENT_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLossWriteOff                 AccountActivityGetTransactionResponseFinancialTransactionEventsType = "LOSS_WRITE_OFF"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeProvisionalCredit            AccountActivityGetTransactionResponseFinancialTransactionEventsType = "PROVISIONAL_CREDIT"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeProvisionalCreditReversal    AccountActivityGetTransactionResponseFinancialTransactionEventsType = "PROVISIONAL_CREDIT_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeService                      AccountActivityGetTransactionResponseFinancialTransactionEventsType = "SERVICE"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturn                       AccountActivityGetTransactionResponseFinancialTransactionEventsType = "RETURN"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnReversal               AccountActivityGetTransactionResponseFinancialTransactionEventsType = "RETURN_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeTransfer                     AccountActivityGetTransactionResponseFinancialTransactionEventsType = "TRANSFER"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeTransferInsufficientFunds    AccountActivityGetTransactionResponseFinancialTransactionEventsType = "TRANSFER_INSUFFICIENT_FUNDS"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnedPayment              AccountActivityGetTransactionResponseFinancialTransactionEventsType = "RETURNED_PAYMENT"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnedPaymentReversal      AccountActivityGetTransactionResponseFinancialTransactionEventsType = "RETURNED_PAYMENT_REVERSAL"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLithicNetworkPayment         AccountActivityGetTransactionResponseFinancialTransactionEventsType = "LITHIC_NETWORK_PAYMENT"
)

func (r AccountActivityGetTransactionResponseFinancialTransactionEventsType) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationCancelled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationProcessed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationRejected, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationReviewed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptProcessed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnProcessed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnRejected, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorization, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationAdvice, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationExpiry, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBalanceInquiry, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBillingError, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBillingErrorReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCardToCard, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCashBack, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCashBackReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeClearing, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCollection, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCorrectionCredit, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCorrectionDebit, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCreditAuthorization, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCreditAuthorizationAdvice, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCurrencyConversion, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCurrencyConversionReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeDisputeWon, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHCanceled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHReversed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckCanceled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckReversed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferCanceled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferReversed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireCanceled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireReversed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeFinancialAuthorization, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeFinancialCreditAuthorization, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInterest, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInterestReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInternalAdjustment, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLatePayment, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLatePaymentReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLossWriteOff, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeProvisionalCredit, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeProvisionalCreditReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeService, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturn, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeTransfer, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeTransferInsufficientFunds, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnedPayment, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnedPaymentReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLithicNetworkPayment:
		return true
	}
	return false
}

// INTERNAL - Financial Transaction
type AccountActivityGetTransactionResponseFinancialTransactionFamily string

const (
	AccountActivityGetTransactionResponseFinancialTransactionFamilyInternal AccountActivityGetTransactionResponseFinancialTransactionFamily = "INTERNAL"
)

func (r AccountActivityGetTransactionResponseFinancialTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFinancialTransactionFamilyInternal:
		return true
	}
	return false
}

// Transaction result
type AccountActivityGetTransactionResponseFinancialTransactionResult string

const (
	AccountActivityGetTransactionResponseFinancialTransactionResultApproved AccountActivityGetTransactionResponseFinancialTransactionResult = "APPROVED"
	AccountActivityGetTransactionResponseFinancialTransactionResultDeclined AccountActivityGetTransactionResponseFinancialTransactionResult = "DECLINED"
)

func (r AccountActivityGetTransactionResponseFinancialTransactionResult) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFinancialTransactionResultApproved, AccountActivityGetTransactionResponseFinancialTransactionResultDeclined:
		return true
	}
	return false
}

// The status of the transaction
type AccountActivityGetTransactionResponseFinancialTransactionStatus string

const (
	AccountActivityGetTransactionResponseFinancialTransactionStatusPending  AccountActivityGetTransactionResponseFinancialTransactionStatus = "PENDING"
	AccountActivityGetTransactionResponseFinancialTransactionStatusSettled  AccountActivityGetTransactionResponseFinancialTransactionStatus = "SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionStatusDeclined AccountActivityGetTransactionResponseFinancialTransactionStatus = "DECLINED"
	AccountActivityGetTransactionResponseFinancialTransactionStatusReversed AccountActivityGetTransactionResponseFinancialTransactionStatus = "REVERSED"
	AccountActivityGetTransactionResponseFinancialTransactionStatusCanceled AccountActivityGetTransactionResponseFinancialTransactionStatus = "CANCELED"
)

func (r AccountActivityGetTransactionResponseFinancialTransactionStatus) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFinancialTransactionStatusPending, AccountActivityGetTransactionResponseFinancialTransactionStatusSettled, AccountActivityGetTransactionResponseFinancialTransactionStatusDeclined, AccountActivityGetTransactionResponseFinancialTransactionStatusReversed, AccountActivityGetTransactionResponseFinancialTransactionStatusCanceled:
		return true
	}
	return false
}

// Card transaction with ledger base properties
type AccountActivityGetTransactionResponseCardTransaction struct {
	JSON accountActivityGetTransactionResponseCardTransactionJSON `json:"-"`
	Transaction
}

// accountActivityGetTransactionResponseCardTransactionJSON contains the JSON
// metadata for the struct [AccountActivityGetTransactionResponseCardTransaction]
type accountActivityGetTransactionResponseCardTransactionJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponseCardTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponseCardTransactionJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityGetTransactionResponseCardTransaction) implementsAccountActivityGetTransactionResponse() {
}

// The status of the transaction
type AccountActivityGetTransactionResponseStatus string

const (
	AccountActivityGetTransactionResponseStatusPending  AccountActivityGetTransactionResponseStatus = "PENDING"
	AccountActivityGetTransactionResponseStatusSettled  AccountActivityGetTransactionResponseStatus = "SETTLED"
	AccountActivityGetTransactionResponseStatusDeclined AccountActivityGetTransactionResponseStatus = "DECLINED"
	AccountActivityGetTransactionResponseStatusReversed AccountActivityGetTransactionResponseStatus = "REVERSED"
	AccountActivityGetTransactionResponseStatusCanceled AccountActivityGetTransactionResponseStatus = "CANCELED"
	AccountActivityGetTransactionResponseStatusExpired  AccountActivityGetTransactionResponseStatus = "EXPIRED"
	AccountActivityGetTransactionResponseStatusVoided   AccountActivityGetTransactionResponseStatus = "VOIDED"
)

func (r AccountActivityGetTransactionResponseStatus) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseStatusPending, AccountActivityGetTransactionResponseStatusSettled, AccountActivityGetTransactionResponseStatusDeclined, AccountActivityGetTransactionResponseStatusReversed, AccountActivityGetTransactionResponseStatusCanceled, AccountActivityGetTransactionResponseStatusExpired, AccountActivityGetTransactionResponseStatusVoided:
		return true
	}
	return false
}

// Transaction category
type AccountActivityGetTransactionResponseCategory string

const (
	AccountActivityGetTransactionResponseCategoryACH                    AccountActivityGetTransactionResponseCategory = "ACH"
	AccountActivityGetTransactionResponseCategoryBalanceOrFunding       AccountActivityGetTransactionResponseCategory = "BALANCE_OR_FUNDING"
	AccountActivityGetTransactionResponseCategoryFee                    AccountActivityGetTransactionResponseCategory = "FEE"
	AccountActivityGetTransactionResponseCategoryReward                 AccountActivityGetTransactionResponseCategory = "REWARD"
	AccountActivityGetTransactionResponseCategoryAdjustment             AccountActivityGetTransactionResponseCategory = "ADJUSTMENT"
	AccountActivityGetTransactionResponseCategoryDerecognition          AccountActivityGetTransactionResponseCategory = "DERECOGNITION"
	AccountActivityGetTransactionResponseCategoryDispute                AccountActivityGetTransactionResponseCategory = "DISPUTE"
	AccountActivityGetTransactionResponseCategoryCard                   AccountActivityGetTransactionResponseCategory = "CARD"
	AccountActivityGetTransactionResponseCategoryExternalACH            AccountActivityGetTransactionResponseCategory = "EXTERNAL_ACH"
	AccountActivityGetTransactionResponseCategoryExternalCheck          AccountActivityGetTransactionResponseCategory = "EXTERNAL_CHECK"
	AccountActivityGetTransactionResponseCategoryExternalTransfer       AccountActivityGetTransactionResponseCategory = "EXTERNAL_TRANSFER"
	AccountActivityGetTransactionResponseCategoryExternalWire           AccountActivityGetTransactionResponseCategory = "EXTERNAL_WIRE"
	AccountActivityGetTransactionResponseCategoryManagementAdjustment   AccountActivityGetTransactionResponseCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityGetTransactionResponseCategoryManagementDispute      AccountActivityGetTransactionResponseCategory = "MANAGEMENT_DISPUTE"
	AccountActivityGetTransactionResponseCategoryManagementFee          AccountActivityGetTransactionResponseCategory = "MANAGEMENT_FEE"
	AccountActivityGetTransactionResponseCategoryManagementReward       AccountActivityGetTransactionResponseCategory = "MANAGEMENT_REWARD"
	AccountActivityGetTransactionResponseCategoryManagementDisbursement AccountActivityGetTransactionResponseCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityGetTransactionResponseCategoryProgramFunding         AccountActivityGetTransactionResponseCategory = "PROGRAM_FUNDING"
	AccountActivityGetTransactionResponseCategoryInternal               AccountActivityGetTransactionResponseCategory = "INTERNAL"
	AccountActivityGetTransactionResponseCategoryTransfer               AccountActivityGetTransactionResponseCategory = "TRANSFER"
)

func (r AccountActivityGetTransactionResponseCategory) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseCategoryACH, AccountActivityGetTransactionResponseCategoryBalanceOrFunding, AccountActivityGetTransactionResponseCategoryFee, AccountActivityGetTransactionResponseCategoryReward, AccountActivityGetTransactionResponseCategoryAdjustment, AccountActivityGetTransactionResponseCategoryDerecognition, AccountActivityGetTransactionResponseCategoryDispute, AccountActivityGetTransactionResponseCategoryCard, AccountActivityGetTransactionResponseCategoryExternalACH, AccountActivityGetTransactionResponseCategoryExternalCheck, AccountActivityGetTransactionResponseCategoryExternalTransfer, AccountActivityGetTransactionResponseCategoryExternalWire, AccountActivityGetTransactionResponseCategoryManagementAdjustment, AccountActivityGetTransactionResponseCategoryManagementDispute, AccountActivityGetTransactionResponseCategoryManagementFee, AccountActivityGetTransactionResponseCategoryManagementReward, AccountActivityGetTransactionResponseCategoryManagementDisbursement, AccountActivityGetTransactionResponseCategoryProgramFunding, AccountActivityGetTransactionResponseCategoryInternal, AccountActivityGetTransactionResponseCategoryTransfer:
		return true
	}
	return false
}

// Transfer direction
type AccountActivityGetTransactionResponseDirection string

const (
	AccountActivityGetTransactionResponseDirectionCredit AccountActivityGetTransactionResponseDirection = "CREDIT"
	AccountActivityGetTransactionResponseDirectionDebit  AccountActivityGetTransactionResponseDirection = "DEBIT"
)

func (r AccountActivityGetTransactionResponseDirection) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseDirectionCredit, AccountActivityGetTransactionResponseDirectionDebit:
		return true
	}
	return false
}

// INTERNAL - Financial Transaction
type AccountActivityGetTransactionResponseFamily string

const (
	AccountActivityGetTransactionResponseFamilyInternal            AccountActivityGetTransactionResponseFamily = "INTERNAL"
	AccountActivityGetTransactionResponseFamilyTransfer            AccountActivityGetTransactionResponseFamily = "TRANSFER"
	AccountActivityGetTransactionResponseFamilyCard                AccountActivityGetTransactionResponseFamily = "CARD"
	AccountActivityGetTransactionResponseFamilyPayment             AccountActivityGetTransactionResponseFamily = "PAYMENT"
	AccountActivityGetTransactionResponseFamilyExternalPayment     AccountActivityGetTransactionResponseFamily = "EXTERNAL_PAYMENT"
	AccountActivityGetTransactionResponseFamilyManagementOperation AccountActivityGetTransactionResponseFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityGetTransactionResponseFamily) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFamilyInternal, AccountActivityGetTransactionResponseFamilyTransfer, AccountActivityGetTransactionResponseFamilyCard, AccountActivityGetTransactionResponseFamilyPayment, AccountActivityGetTransactionResponseFamilyExternalPayment, AccountActivityGetTransactionResponseFamilyManagementOperation:
		return true
	}
	return false
}

// Transfer method
type AccountActivityGetTransactionResponseMethod string

const (
	AccountActivityGetTransactionResponseMethodACHNextDay AccountActivityGetTransactionResponseMethod = "ACH_NEXT_DAY"
	AccountActivityGetTransactionResponseMethodACHSameDay AccountActivityGetTransactionResponseMethod = "ACH_SAME_DAY"
	AccountActivityGetTransactionResponseMethodWire       AccountActivityGetTransactionResponseMethod = "WIRE"
)

func (r AccountActivityGetTransactionResponseMethod) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseMethodACHNextDay, AccountActivityGetTransactionResponseMethodACHSameDay, AccountActivityGetTransactionResponseMethodWire:
		return true
	}
	return false
}

// Card network of the authorization. Value is `UNKNOWN` when Lithic cannot
// determine the network code from the upstream provider.
type AccountActivityGetTransactionResponseNetwork string

const (
	AccountActivityGetTransactionResponseNetworkAmex       AccountActivityGetTransactionResponseNetwork = "AMEX"
	AccountActivityGetTransactionResponseNetworkInterlink  AccountActivityGetTransactionResponseNetwork = "INTERLINK"
	AccountActivityGetTransactionResponseNetworkMaestro    AccountActivityGetTransactionResponseNetwork = "MAESTRO"
	AccountActivityGetTransactionResponseNetworkMastercard AccountActivityGetTransactionResponseNetwork = "MASTERCARD"
	AccountActivityGetTransactionResponseNetworkUnknown    AccountActivityGetTransactionResponseNetwork = "UNKNOWN"
	AccountActivityGetTransactionResponseNetworkVisa       AccountActivityGetTransactionResponseNetwork = "VISA"
)

func (r AccountActivityGetTransactionResponseNetwork) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseNetworkAmex, AccountActivityGetTransactionResponseNetworkInterlink, AccountActivityGetTransactionResponseNetworkMaestro, AccountActivityGetTransactionResponseNetworkMastercard, AccountActivityGetTransactionResponseNetworkUnknown, AccountActivityGetTransactionResponseNetworkVisa:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponsePaymentType string

const (
	AccountActivityGetTransactionResponsePaymentTypeDeposit    AccountActivityGetTransactionResponsePaymentType = "DEPOSIT"
	AccountActivityGetTransactionResponsePaymentTypeWithdrawal AccountActivityGetTransactionResponsePaymentType = "WITHDRAWAL"
)

func (r AccountActivityGetTransactionResponsePaymentType) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTypeDeposit, AccountActivityGetTransactionResponsePaymentTypeWithdrawal:
		return true
	}
	return false
}

// Transaction result
type AccountActivityGetTransactionResponseResult string

const (
	AccountActivityGetTransactionResponseResultApproved                    AccountActivityGetTransactionResponseResult = "APPROVED"
	AccountActivityGetTransactionResponseResultDeclined                    AccountActivityGetTransactionResponseResult = "DECLINED"
	AccountActivityGetTransactionResponseResultAccountPaused               AccountActivityGetTransactionResponseResult = "ACCOUNT_PAUSED"
	AccountActivityGetTransactionResponseResultAccountStateTransactionFail AccountActivityGetTransactionResponseResult = "ACCOUNT_STATE_TRANSACTION_FAIL"
	AccountActivityGetTransactionResponseResultBankConnectionError         AccountActivityGetTransactionResponseResult = "BANK_CONNECTION_ERROR"
	AccountActivityGetTransactionResponseResultBankNotVerified             AccountActivityGetTransactionResponseResult = "BANK_NOT_VERIFIED"
	AccountActivityGetTransactionResponseResultCardClosed                  AccountActivityGetTransactionResponseResult = "CARD_CLOSED"
	AccountActivityGetTransactionResponseResultCardPaused                  AccountActivityGetTransactionResponseResult = "CARD_PAUSED"
	AccountActivityGetTransactionResponseResultFraudAdvice                 AccountActivityGetTransactionResponseResult = "FRAUD_ADVICE"
	AccountActivityGetTransactionResponseResultIgnoredTtlExpiry            AccountActivityGetTransactionResponseResult = "IGNORED_TTL_EXPIRY"
	AccountActivityGetTransactionResponseResultSuspectedFraud              AccountActivityGetTransactionResponseResult = "SUSPECTED_FRAUD"
	AccountActivityGetTransactionResponseResultInactiveAccount             AccountActivityGetTransactionResponseResult = "INACTIVE_ACCOUNT"
	AccountActivityGetTransactionResponseResultIncorrectPin                AccountActivityGetTransactionResponseResult = "INCORRECT_PIN"
	AccountActivityGetTransactionResponseResultInvalidCardDetails          AccountActivityGetTransactionResponseResult = "INVALID_CARD_DETAILS"
	AccountActivityGetTransactionResponseResultInsufficientFunds           AccountActivityGetTransactionResponseResult = "INSUFFICIENT_FUNDS"
	AccountActivityGetTransactionResponseResultInsufficientFundsPreload    AccountActivityGetTransactionResponseResult = "INSUFFICIENT_FUNDS_PRELOAD"
	AccountActivityGetTransactionResponseResultInvalidTransaction          AccountActivityGetTransactionResponseResult = "INVALID_TRANSACTION"
	AccountActivityGetTransactionResponseResultMerchantBlacklist           AccountActivityGetTransactionResponseResult = "MERCHANT_BLACKLIST"
	AccountActivityGetTransactionResponseResultOriginalNotFound            AccountActivityGetTransactionResponseResult = "ORIGINAL_NOT_FOUND"
	AccountActivityGetTransactionResponseResultPreviouslyCompleted         AccountActivityGetTransactionResponseResult = "PREVIOUSLY_COMPLETED"
	AccountActivityGetTransactionResponseResultSingleUseRecharged          AccountActivityGetTransactionResponseResult = "SINGLE_USE_RECHARGED"
	AccountActivityGetTransactionResponseResultSwitchInoperativeAdvice     AccountActivityGetTransactionResponseResult = "SWITCH_INOPERATIVE_ADVICE"
	AccountActivityGetTransactionResponseResultUnauthorizedMerchant        AccountActivityGetTransactionResponseResult = "UNAUTHORIZED_MERCHANT"
	AccountActivityGetTransactionResponseResultUnknownHostTimeout          AccountActivityGetTransactionResponseResult = "UNKNOWN_HOST_TIMEOUT"
	AccountActivityGetTransactionResponseResultUserTransactionLimit        AccountActivityGetTransactionResponseResult = "USER_TRANSACTION_LIMIT"
)

func (r AccountActivityGetTransactionResponseResult) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseResultApproved, AccountActivityGetTransactionResponseResultDeclined, AccountActivityGetTransactionResponseResultAccountPaused, AccountActivityGetTransactionResponseResultAccountStateTransactionFail, AccountActivityGetTransactionResponseResultBankConnectionError, AccountActivityGetTransactionResponseResultBankNotVerified, AccountActivityGetTransactionResponseResultCardClosed, AccountActivityGetTransactionResponseResultCardPaused, AccountActivityGetTransactionResponseResultFraudAdvice, AccountActivityGetTransactionResponseResultIgnoredTtlExpiry, AccountActivityGetTransactionResponseResultSuspectedFraud, AccountActivityGetTransactionResponseResultInactiveAccount, AccountActivityGetTransactionResponseResultIncorrectPin, AccountActivityGetTransactionResponseResultInvalidCardDetails, AccountActivityGetTransactionResponseResultInsufficientFunds, AccountActivityGetTransactionResponseResultInsufficientFundsPreload, AccountActivityGetTransactionResponseResultInvalidTransaction, AccountActivityGetTransactionResponseResultMerchantBlacklist, AccountActivityGetTransactionResponseResultOriginalNotFound, AccountActivityGetTransactionResponseResultPreviouslyCompleted, AccountActivityGetTransactionResponseResultSingleUseRecharged, AccountActivityGetTransactionResponseResultSwitchInoperativeAdvice, AccountActivityGetTransactionResponseResultUnauthorizedMerchant, AccountActivityGetTransactionResponseResultUnknownHostTimeout, AccountActivityGetTransactionResponseResultUserTransactionLimit:
		return true
	}
	return false
}

// Transaction source
type AccountActivityGetTransactionResponseSource string

const (
	AccountActivityGetTransactionResponseSourceLithic   AccountActivityGetTransactionResponseSource = "LITHIC"
	AccountActivityGetTransactionResponseSourceExternal AccountActivityGetTransactionResponseSource = "EXTERNAL"
	AccountActivityGetTransactionResponseSourceCustomer AccountActivityGetTransactionResponseSource = "CUSTOMER"
)

func (r AccountActivityGetTransactionResponseSource) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseSourceLithic, AccountActivityGetTransactionResponseSourceExternal, AccountActivityGetTransactionResponseSourceCustomer:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponseType string

const (
	AccountActivityGetTransactionResponseTypeOriginationCredit   AccountActivityGetTransactionResponseType = "ORIGINATION_CREDIT"
	AccountActivityGetTransactionResponseTypeOriginationDebit    AccountActivityGetTransactionResponseType = "ORIGINATION_DEBIT"
	AccountActivityGetTransactionResponseTypeReceiptCredit       AccountActivityGetTransactionResponseType = "RECEIPT_CREDIT"
	AccountActivityGetTransactionResponseTypeReceiptDebit        AccountActivityGetTransactionResponseType = "RECEIPT_DEBIT"
	AccountActivityGetTransactionResponseTypeWireInboundPayment  AccountActivityGetTransactionResponseType = "WIRE_INBOUND_PAYMENT"
	AccountActivityGetTransactionResponseTypeWireInboundAdmin    AccountActivityGetTransactionResponseType = "WIRE_INBOUND_ADMIN"
	AccountActivityGetTransactionResponseTypeWireOutboundPayment AccountActivityGetTransactionResponseType = "WIRE_OUTBOUND_PAYMENT"
	AccountActivityGetTransactionResponseTypeWireOutboundAdmin   AccountActivityGetTransactionResponseType = "WIRE_OUTBOUND_ADMIN"
)

func (r AccountActivityGetTransactionResponseType) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseTypeOriginationCredit, AccountActivityGetTransactionResponseTypeOriginationDebit, AccountActivityGetTransactionResponseTypeReceiptCredit, AccountActivityGetTransactionResponseTypeReceiptDebit, AccountActivityGetTransactionResponseTypeWireInboundPayment, AccountActivityGetTransactionResponseTypeWireInboundAdmin, AccountActivityGetTransactionResponseTypeWireOutboundPayment, AccountActivityGetTransactionResponseTypeWireOutboundAdmin:
		return true
	}
	return false
}

type AccountActivityListParams struct {
	// Filter by account token
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Filter by business account token
	BusinessAccountToken param.Field[string] `query:"business_account_token" format:"uuid"`
	// Filter by transaction category
	Category param.Field[AccountActivityListParamsCategory] `query:"category"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Filter by financial account token
	FinancialAccountToken param.Field[string] `query:"financial_account_token" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by transaction result
	Result param.Field[AccountActivityListParamsResult] `query:"result"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Filter by transaction status
	Status param.Field[AccountActivityListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountActivityListParams]'s query parameters as
// `url.Values`.
func (r AccountActivityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by transaction category
type AccountActivityListParamsCategory string

const (
	AccountActivityListParamsCategoryACH                    AccountActivityListParamsCategory = "ACH"
	AccountActivityListParamsCategoryBalanceOrFunding       AccountActivityListParamsCategory = "BALANCE_OR_FUNDING"
	AccountActivityListParamsCategoryFee                    AccountActivityListParamsCategory = "FEE"
	AccountActivityListParamsCategoryReward                 AccountActivityListParamsCategory = "REWARD"
	AccountActivityListParamsCategoryAdjustment             AccountActivityListParamsCategory = "ADJUSTMENT"
	AccountActivityListParamsCategoryDerecognition          AccountActivityListParamsCategory = "DERECOGNITION"
	AccountActivityListParamsCategoryDispute                AccountActivityListParamsCategory = "DISPUTE"
	AccountActivityListParamsCategoryCard                   AccountActivityListParamsCategory = "CARD"
	AccountActivityListParamsCategoryExternalACH            AccountActivityListParamsCategory = "EXTERNAL_ACH"
	AccountActivityListParamsCategoryExternalCheck          AccountActivityListParamsCategory = "EXTERNAL_CHECK"
	AccountActivityListParamsCategoryExternalTransfer       AccountActivityListParamsCategory = "EXTERNAL_TRANSFER"
	AccountActivityListParamsCategoryExternalWire           AccountActivityListParamsCategory = "EXTERNAL_WIRE"
	AccountActivityListParamsCategoryManagementAdjustment   AccountActivityListParamsCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityListParamsCategoryManagementDispute      AccountActivityListParamsCategory = "MANAGEMENT_DISPUTE"
	AccountActivityListParamsCategoryManagementFee          AccountActivityListParamsCategory = "MANAGEMENT_FEE"
	AccountActivityListParamsCategoryManagementReward       AccountActivityListParamsCategory = "MANAGEMENT_REWARD"
	AccountActivityListParamsCategoryManagementDisbursement AccountActivityListParamsCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityListParamsCategoryProgramFunding         AccountActivityListParamsCategory = "PROGRAM_FUNDING"
)

func (r AccountActivityListParamsCategory) IsKnown() bool {
	switch r {
	case AccountActivityListParamsCategoryACH, AccountActivityListParamsCategoryBalanceOrFunding, AccountActivityListParamsCategoryFee, AccountActivityListParamsCategoryReward, AccountActivityListParamsCategoryAdjustment, AccountActivityListParamsCategoryDerecognition, AccountActivityListParamsCategoryDispute, AccountActivityListParamsCategoryCard, AccountActivityListParamsCategoryExternalACH, AccountActivityListParamsCategoryExternalCheck, AccountActivityListParamsCategoryExternalTransfer, AccountActivityListParamsCategoryExternalWire, AccountActivityListParamsCategoryManagementAdjustment, AccountActivityListParamsCategoryManagementDispute, AccountActivityListParamsCategoryManagementFee, AccountActivityListParamsCategoryManagementReward, AccountActivityListParamsCategoryManagementDisbursement, AccountActivityListParamsCategoryProgramFunding:
		return true
	}
	return false
}

// Filter by transaction result
type AccountActivityListParamsResult string

const (
	AccountActivityListParamsResultApproved AccountActivityListParamsResult = "APPROVED"
	AccountActivityListParamsResultDeclined AccountActivityListParamsResult = "DECLINED"
)

func (r AccountActivityListParamsResult) IsKnown() bool {
	switch r {
	case AccountActivityListParamsResultApproved, AccountActivityListParamsResultDeclined:
		return true
	}
	return false
}

// Filter by transaction status
type AccountActivityListParamsStatus string

const (
	AccountActivityListParamsStatusDeclined AccountActivityListParamsStatus = "DECLINED"
	AccountActivityListParamsStatusExpired  AccountActivityListParamsStatus = "EXPIRED"
	AccountActivityListParamsStatusPending  AccountActivityListParamsStatus = "PENDING"
	AccountActivityListParamsStatusReturned AccountActivityListParamsStatus = "RETURNED"
	AccountActivityListParamsStatusReversed AccountActivityListParamsStatus = "REVERSED"
	AccountActivityListParamsStatusSettled  AccountActivityListParamsStatus = "SETTLED"
	AccountActivityListParamsStatusVoided   AccountActivityListParamsStatus = "VOIDED"
)

func (r AccountActivityListParamsStatus) IsKnown() bool {
	switch r {
	case AccountActivityListParamsStatusDeclined, AccountActivityListParamsStatusExpired, AccountActivityListParamsStatusPending, AccountActivityListParamsStatusReturned, AccountActivityListParamsStatusReversed, AccountActivityListParamsStatusSettled, AccountActivityListParamsStatusVoided:
		return true
	}
	return false
}
