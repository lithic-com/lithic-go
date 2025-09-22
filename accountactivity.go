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

// Response containing multiple transaction types
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
	// [[]AccountActivityListResponseBookTransferTransactionEvent],
	// [[]TransactionEvent], [[]AccountActivityListResponsePaymentTransactionEvent],
	// [[]ExternalPaymentEvent], [[]ManagementOperationTransactionEvent].
	Events interface{} `json:"events"`
	// Expected release date for the transaction
	ExpectedReleaseDate time.Time `json:"expected_release_date,nullable" format:"date"`
	// External bank account token
	ExternalBankAccountToken string `json:"external_bank_account_token,nullable" format:"uuid"`
	// External identifier for the transaction
	ExternalID string `json:"external_id"`
	// External resource associated with the management operation
	ExternalResource ExternalResource                  `json:"external_resource,nullable"`
	Family           AccountActivityListResponseFamily `json:"family"`
	// Financial account token associated with the transaction
	FinancialAccountToken string `json:"financial_account_token" format:"uuid"`
	// Source account token
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
	// This field can have the runtime type of
	// [AccountActivityListResponsePaymentTransactionMethodAttributes].
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
	// This field can have the runtime type of
	// [AccountActivityListResponsePaymentTransactionRelatedAccountTokens].
	RelatedAccountTokens interface{} `json:"related_account_tokens"`
	// Transaction result
	Result AccountActivityListResponseResult `json:"result"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount"`
	// Transaction source
	Source AccountActivityListResponseSource `json:"source"`
	// Destination account token
	ToFinancialAccountToken string `json:"to_financial_account_token" format:"uuid"`
	// This field can have the runtime type of [TransactionTokenInfo].
	TokenInfo interface{} `json:"token_info"`
	// This field can have the runtime type of
	// [AccountActivityListResponseBookTransferTransactionTransactionSeries],
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
// [AccountActivityListResponseFinancialTransaction],
// [AccountActivityListResponseBookTransferTransaction],
// [AccountActivityListResponseCardTransaction],
// [AccountActivityListResponsePaymentTransaction], [ExternalPayment],
// [ManagementOperationTransaction].
func (r AccountActivityListResponse) AsUnion() AccountActivityListResponseUnion {
	return r.union
}

// Response containing multiple transaction types
//
// Union satisfied by [AccountActivityListResponseFinancialTransaction],
// [AccountActivityListResponseBookTransferTransaction],
// [AccountActivityListResponseCardTransaction],
// [AccountActivityListResponsePaymentTransaction], [ExternalPayment] or
// [ManagementOperationTransaction].
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
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseFinancialTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseFinancialTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseFinancialTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseFinancialTransaction{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseFinancialTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseBookTransferTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseBookTransferTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseBookTransferTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseBookTransferTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseBookTransferTransaction{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseBookTransferTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseCardTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseCardTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseCardTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseCardTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseCardTransaction{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponseCardTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponsePaymentTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponsePaymentTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponsePaymentTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponsePaymentTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponsePaymentTransaction{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityListResponsePaymentTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
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
	Family AccountActivityListResponseFinancialTransactionFamily  `json:"family,required"`
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
	case AccountActivityListResponseFinancialTransactionCategoryACH, AccountActivityListResponseFinancialTransactionCategoryBalanceOrFunding, AccountActivityListResponseFinancialTransactionCategoryCard, AccountActivityListResponseFinancialTransactionCategoryExternalACH, AccountActivityListResponseFinancialTransactionCategoryExternalCheck, AccountActivityListResponseFinancialTransactionCategoryExternalTransfer, AccountActivityListResponseFinancialTransactionCategoryExternalWire, AccountActivityListResponseFinancialTransactionCategoryManagementAdjustment, AccountActivityListResponseFinancialTransactionCategoryManagementDispute, AccountActivityListResponseFinancialTransactionCategoryManagementFee, AccountActivityListResponseFinancialTransactionCategoryManagementReward, AccountActivityListResponseFinancialTransactionCategoryManagementDisbursement, AccountActivityListResponseFinancialTransactionCategoryProgramFunding:
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
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationReviewed       AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationSettled        AccountActivityListResponseFinancialTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptProcessed          AccountActivityListResponseFinancialTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptSettled            AccountActivityListResponseFinancialTransactionEventsType = "ACH_RECEIPT_SETTLED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReturnInitiated           AccountActivityListResponseFinancialTransactionEventsType = "ACH_RETURN_INITIATED"
	AccountActivityListResponseFinancialTransactionEventsTypeACHReturnProcessed           AccountActivityListResponseFinancialTransactionEventsType = "ACH_RETURN_PROCESSED"
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
	case AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationCancelled, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationInitiated, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationProcessed, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationReleased, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationReviewed, AccountActivityListResponseFinancialTransactionEventsTypeACHOriginationSettled, AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptProcessed, AccountActivityListResponseFinancialTransactionEventsTypeACHReceiptSettled, AccountActivityListResponseFinancialTransactionEventsTypeACHReturnInitiated, AccountActivityListResponseFinancialTransactionEventsTypeACHReturnProcessed, AccountActivityListResponseFinancialTransactionEventsTypeACHReturnSettled, AccountActivityListResponseFinancialTransactionEventsTypeAuthorization, AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationAdvice, AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationExpiry, AccountActivityListResponseFinancialTransactionEventsTypeAuthorizationReversal, AccountActivityListResponseFinancialTransactionEventsTypeBalanceInquiry, AccountActivityListResponseFinancialTransactionEventsTypeBillingError, AccountActivityListResponseFinancialTransactionEventsTypeBillingErrorReversal, AccountActivityListResponseFinancialTransactionEventsTypeCardToCard, AccountActivityListResponseFinancialTransactionEventsTypeCashBack, AccountActivityListResponseFinancialTransactionEventsTypeCashBackReversal, AccountActivityListResponseFinancialTransactionEventsTypeClearing, AccountActivityListResponseFinancialTransactionEventsTypeCollection, AccountActivityListResponseFinancialTransactionEventsTypeCorrectionCredit, AccountActivityListResponseFinancialTransactionEventsTypeCorrectionDebit, AccountActivityListResponseFinancialTransactionEventsTypeCreditAuthorization, AccountActivityListResponseFinancialTransactionEventsTypeCreditAuthorizationAdvice, AccountActivityListResponseFinancialTransactionEventsTypeCurrencyConversion, AccountActivityListResponseFinancialTransactionEventsTypeCurrencyConversionReversal, AccountActivityListResponseFinancialTransactionEventsTypeDisputeWon, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHCanceled, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHInitiated, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHReleased, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHReversed, AccountActivityListResponseFinancialTransactionEventsTypeExternalACHSettled, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckCanceled, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckInitiated, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckReleased, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckReversed, AccountActivityListResponseFinancialTransactionEventsTypeExternalCheckSettled, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferCanceled, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferInitiated, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferReleased, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferReversed, AccountActivityListResponseFinancialTransactionEventsTypeExternalTransferSettled, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireCanceled, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireInitiated, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireReleased, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireReversed, AccountActivityListResponseFinancialTransactionEventsTypeExternalWireSettled, AccountActivityListResponseFinancialTransactionEventsTypeFinancialAuthorization, AccountActivityListResponseFinancialTransactionEventsTypeFinancialCreditAuthorization, AccountActivityListResponseFinancialTransactionEventsTypeInterest, AccountActivityListResponseFinancialTransactionEventsTypeInterestReversal, AccountActivityListResponseFinancialTransactionEventsTypeInternalAdjustment, AccountActivityListResponseFinancialTransactionEventsTypeLatePayment, AccountActivityListResponseFinancialTransactionEventsTypeLatePaymentReversal, AccountActivityListResponseFinancialTransactionEventsTypeLossWriteOff, AccountActivityListResponseFinancialTransactionEventsTypeProvisionalCredit, AccountActivityListResponseFinancialTransactionEventsTypeProvisionalCreditReversal, AccountActivityListResponseFinancialTransactionEventsTypeService, AccountActivityListResponseFinancialTransactionEventsTypeReturn, AccountActivityListResponseFinancialTransactionEventsTypeReturnReversal, AccountActivityListResponseFinancialTransactionEventsTypeTransfer, AccountActivityListResponseFinancialTransactionEventsTypeTransferInsufficientFunds, AccountActivityListResponseFinancialTransactionEventsTypeReturnedPayment, AccountActivityListResponseFinancialTransactionEventsTypeReturnedPaymentReversal, AccountActivityListResponseFinancialTransactionEventsTypeLithicNetworkPayment:
		return true
	}
	return false
}

type AccountActivityListResponseFinancialTransactionFamily string

const (
	AccountActivityListResponseFinancialTransactionFamilyCard                AccountActivityListResponseFinancialTransactionFamily = "CARD"
	AccountActivityListResponseFinancialTransactionFamilyPayment             AccountActivityListResponseFinancialTransactionFamily = "PAYMENT"
	AccountActivityListResponseFinancialTransactionFamilyTransfer            AccountActivityListResponseFinancialTransactionFamily = "TRANSFER"
	AccountActivityListResponseFinancialTransactionFamilyInternal            AccountActivityListResponseFinancialTransactionFamily = "INTERNAL"
	AccountActivityListResponseFinancialTransactionFamilyExternalPayment     AccountActivityListResponseFinancialTransactionFamily = "EXTERNAL_PAYMENT"
	AccountActivityListResponseFinancialTransactionFamilyManagementOperation AccountActivityListResponseFinancialTransactionFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityListResponseFinancialTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFinancialTransactionFamilyCard, AccountActivityListResponseFinancialTransactionFamilyPayment, AccountActivityListResponseFinancialTransactionFamilyTransfer, AccountActivityListResponseFinancialTransactionFamilyInternal, AccountActivityListResponseFinancialTransactionFamilyExternalPayment, AccountActivityListResponseFinancialTransactionFamilyManagementOperation:
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

// Book transfer transaction
type AccountActivityListResponseBookTransferTransaction struct {
	// Unique identifier for the transaction
	Token    string                                                     `json:"token,required" format:"uuid"`
	Category AccountActivityListResponseBookTransferTransactionCategory `json:"category,required"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Currency of the transaction in ISO 4217 format
	Currency string `json:"currency,required"`
	// List of events associated with this book transfer
	Events []AccountActivityListResponseBookTransferTransactionEvent `json:"events,required"`
	Family AccountActivityListResponseBookTransferTransactionFamily  `json:"family,required"`
	// Source account token
	FromFinancialAccountToken string `json:"from_financial_account_token,required" format:"uuid"`
	// The pending amount of the transaction in cents
	PendingAmount int64                                                    `json:"pending_amount,required"`
	Result        AccountActivityListResponseBookTransferTransactionResult `json:"result,required"`
	// The settled amount of the transaction in cents
	SettledAmount int64 `json:"settled_amount,required"`
	// The status of the transaction
	Status AccountActivityListResponseBookTransferTransactionStatus `json:"status,required"`
	// Destination account token
	ToFinancialAccountToken string `json:"to_financial_account_token,required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// External identifier for the transaction
	ExternalID string `json:"external_id"`
	// External resource associated with the management operation
	ExternalResource  ExternalResource                                                    `json:"external_resource,nullable"`
	TransactionSeries AccountActivityListResponseBookTransferTransactionTransactionSeries `json:"transaction_series,nullable"`
	JSON              accountActivityListResponseBookTransferTransactionJSON              `json:"-"`
}

// accountActivityListResponseBookTransferTransactionJSON contains the JSON
// metadata for the struct [AccountActivityListResponseBookTransferTransaction]
type accountActivityListResponseBookTransferTransactionJSON struct {
	Token                     apijson.Field
	Category                  apijson.Field
	Created                   apijson.Field
	Currency                  apijson.Field
	Events                    apijson.Field
	Family                    apijson.Field
	FromFinancialAccountToken apijson.Field
	PendingAmount             apijson.Field
	Result                    apijson.Field
	SettledAmount             apijson.Field
	Status                    apijson.Field
	ToFinancialAccountToken   apijson.Field
	Updated                   apijson.Field
	ExternalID                apijson.Field
	ExternalResource          apijson.Field
	TransactionSeries         apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountActivityListResponseBookTransferTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponseBookTransferTransactionJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityListResponseBookTransferTransaction) implementsAccountActivityListResponse() {}

type AccountActivityListResponseBookTransferTransactionCategory string

const (
	AccountActivityListResponseBookTransferTransactionCategoryACH                    AccountActivityListResponseBookTransferTransactionCategory = "ACH"
	AccountActivityListResponseBookTransferTransactionCategoryBalanceOrFunding       AccountActivityListResponseBookTransferTransactionCategory = "BALANCE_OR_FUNDING"
	AccountActivityListResponseBookTransferTransactionCategoryCard                   AccountActivityListResponseBookTransferTransactionCategory = "CARD"
	AccountActivityListResponseBookTransferTransactionCategoryExternalACH            AccountActivityListResponseBookTransferTransactionCategory = "EXTERNAL_ACH"
	AccountActivityListResponseBookTransferTransactionCategoryExternalCheck          AccountActivityListResponseBookTransferTransactionCategory = "EXTERNAL_CHECK"
	AccountActivityListResponseBookTransferTransactionCategoryExternalTransfer       AccountActivityListResponseBookTransferTransactionCategory = "EXTERNAL_TRANSFER"
	AccountActivityListResponseBookTransferTransactionCategoryExternalWire           AccountActivityListResponseBookTransferTransactionCategory = "EXTERNAL_WIRE"
	AccountActivityListResponseBookTransferTransactionCategoryManagementAdjustment   AccountActivityListResponseBookTransferTransactionCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityListResponseBookTransferTransactionCategoryManagementDispute      AccountActivityListResponseBookTransferTransactionCategory = "MANAGEMENT_DISPUTE"
	AccountActivityListResponseBookTransferTransactionCategoryManagementFee          AccountActivityListResponseBookTransferTransactionCategory = "MANAGEMENT_FEE"
	AccountActivityListResponseBookTransferTransactionCategoryManagementReward       AccountActivityListResponseBookTransferTransactionCategory = "MANAGEMENT_REWARD"
	AccountActivityListResponseBookTransferTransactionCategoryManagementDisbursement AccountActivityListResponseBookTransferTransactionCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityListResponseBookTransferTransactionCategoryProgramFunding         AccountActivityListResponseBookTransferTransactionCategory = "PROGRAM_FUNDING"
)

func (r AccountActivityListResponseBookTransferTransactionCategory) IsKnown() bool {
	switch r {
	case AccountActivityListResponseBookTransferTransactionCategoryACH, AccountActivityListResponseBookTransferTransactionCategoryBalanceOrFunding, AccountActivityListResponseBookTransferTransactionCategoryCard, AccountActivityListResponseBookTransferTransactionCategoryExternalACH, AccountActivityListResponseBookTransferTransactionCategoryExternalCheck, AccountActivityListResponseBookTransferTransactionCategoryExternalTransfer, AccountActivityListResponseBookTransferTransactionCategoryExternalWire, AccountActivityListResponseBookTransferTransactionCategoryManagementAdjustment, AccountActivityListResponseBookTransferTransactionCategoryManagementDispute, AccountActivityListResponseBookTransferTransactionCategoryManagementFee, AccountActivityListResponseBookTransferTransactionCategoryManagementReward, AccountActivityListResponseBookTransferTransactionCategoryManagementDisbursement, AccountActivityListResponseBookTransferTransactionCategoryProgramFunding:
		return true
	}
	return false
}

// Book transfer Event
type AccountActivityListResponseBookTransferTransactionEvent struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount,required"`
	// Date and time when the financial event occurred. UTC time zone.
	Created         time.Time                                                               `json:"created,required" format:"date-time"`
	DetailedResults AccountActivityListResponseBookTransferTransactionEventsDetailedResults `json:"detailed_results,required"`
	// Memo for the transfer.
	Memo string `json:"memo,required"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result AccountActivityListResponseBookTransferTransactionEventsResult `json:"result,required"`
	// The program specific subtype code for the specified category/type.
	Subtype string `json:"subtype,required"`
	// Type of the book transfer
	Type AccountActivityListResponseBookTransferTransactionEventsType `json:"type,required"`
	JSON accountActivityListResponseBookTransferTransactionEventJSON  `json:"-"`
}

// accountActivityListResponseBookTransferTransactionEventJSON contains the JSON
// metadata for the struct
// [AccountActivityListResponseBookTransferTransactionEvent]
type accountActivityListResponseBookTransferTransactionEventJSON struct {
	Token           apijson.Field
	Amount          apijson.Field
	Created         apijson.Field
	DetailedResults apijson.Field
	Memo            apijson.Field
	Result          apijson.Field
	Subtype         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountActivityListResponseBookTransferTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponseBookTransferTransactionEventJSON) RawJSON() string {
	return r.raw
}

type AccountActivityListResponseBookTransferTransactionEventsDetailedResults string

const (
	AccountActivityListResponseBookTransferTransactionEventsDetailedResultsApproved          AccountActivityListResponseBookTransferTransactionEventsDetailedResults = "APPROVED"
	AccountActivityListResponseBookTransferTransactionEventsDetailedResultsFundsInsufficient AccountActivityListResponseBookTransferTransactionEventsDetailedResults = "FUNDS_INSUFFICIENT"
)

func (r AccountActivityListResponseBookTransferTransactionEventsDetailedResults) IsKnown() bool {
	switch r {
	case AccountActivityListResponseBookTransferTransactionEventsDetailedResultsApproved, AccountActivityListResponseBookTransferTransactionEventsDetailedResultsFundsInsufficient:
		return true
	}
	return false
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type AccountActivityListResponseBookTransferTransactionEventsResult string

const (
	AccountActivityListResponseBookTransferTransactionEventsResultApproved AccountActivityListResponseBookTransferTransactionEventsResult = "APPROVED"
	AccountActivityListResponseBookTransferTransactionEventsResultDeclined AccountActivityListResponseBookTransferTransactionEventsResult = "DECLINED"
)

func (r AccountActivityListResponseBookTransferTransactionEventsResult) IsKnown() bool {
	switch r {
	case AccountActivityListResponseBookTransferTransactionEventsResultApproved, AccountActivityListResponseBookTransferTransactionEventsResultDeclined:
		return true
	}
	return false
}

// Type of the book transfer
type AccountActivityListResponseBookTransferTransactionEventsType string

const (
	AccountActivityListResponseBookTransferTransactionEventsTypeAtmWithdrawal              AccountActivityListResponseBookTransferTransactionEventsType = "ATM_WITHDRAWAL"
	AccountActivityListResponseBookTransferTransactionEventsTypeAtmDecline                 AccountActivityListResponseBookTransferTransactionEventsType = "ATM_DECLINE"
	AccountActivityListResponseBookTransferTransactionEventsTypeInternationalAtmWithdrawal AccountActivityListResponseBookTransferTransactionEventsType = "INTERNATIONAL_ATM_WITHDRAWAL"
	AccountActivityListResponseBookTransferTransactionEventsTypeInactivity                 AccountActivityListResponseBookTransferTransactionEventsType = "INACTIVITY"
	AccountActivityListResponseBookTransferTransactionEventsTypeStatement                  AccountActivityListResponseBookTransferTransactionEventsType = "STATEMENT"
	AccountActivityListResponseBookTransferTransactionEventsTypeMonthly                    AccountActivityListResponseBookTransferTransactionEventsType = "MONTHLY"
	AccountActivityListResponseBookTransferTransactionEventsTypeQuarterly                  AccountActivityListResponseBookTransferTransactionEventsType = "QUARTERLY"
	AccountActivityListResponseBookTransferTransactionEventsTypeAnnual                     AccountActivityListResponseBookTransferTransactionEventsType = "ANNUAL"
	AccountActivityListResponseBookTransferTransactionEventsTypeCustomerService            AccountActivityListResponseBookTransferTransactionEventsType = "CUSTOMER_SERVICE"
	AccountActivityListResponseBookTransferTransactionEventsTypeAccountMaintenance         AccountActivityListResponseBookTransferTransactionEventsType = "ACCOUNT_MAINTENANCE"
	AccountActivityListResponseBookTransferTransactionEventsTypeAccountActivation          AccountActivityListResponseBookTransferTransactionEventsType = "ACCOUNT_ACTIVATION"
	AccountActivityListResponseBookTransferTransactionEventsTypeAccountClosure             AccountActivityListResponseBookTransferTransactionEventsType = "ACCOUNT_CLOSURE"
	AccountActivityListResponseBookTransferTransactionEventsTypeCardReplacement            AccountActivityListResponseBookTransferTransactionEventsType = "CARD_REPLACEMENT"
	AccountActivityListResponseBookTransferTransactionEventsTypeCardDelivery               AccountActivityListResponseBookTransferTransactionEventsType = "CARD_DELIVERY"
	AccountActivityListResponseBookTransferTransactionEventsTypeCardCreate                 AccountActivityListResponseBookTransferTransactionEventsType = "CARD_CREATE"
	AccountActivityListResponseBookTransferTransactionEventsTypeCurrencyConversion         AccountActivityListResponseBookTransferTransactionEventsType = "CURRENCY_CONVERSION"
	AccountActivityListResponseBookTransferTransactionEventsTypeInterest                   AccountActivityListResponseBookTransferTransactionEventsType = "INTEREST"
	AccountActivityListResponseBookTransferTransactionEventsTypeLatePayment                AccountActivityListResponseBookTransferTransactionEventsType = "LATE_PAYMENT"
	AccountActivityListResponseBookTransferTransactionEventsTypeBillPayment                AccountActivityListResponseBookTransferTransactionEventsType = "BILL_PAYMENT"
	AccountActivityListResponseBookTransferTransactionEventsTypeCashBack                   AccountActivityListResponseBookTransferTransactionEventsType = "CASH_BACK"
	AccountActivityListResponseBookTransferTransactionEventsTypeAccountToAccount           AccountActivityListResponseBookTransferTransactionEventsType = "ACCOUNT_TO_ACCOUNT"
	AccountActivityListResponseBookTransferTransactionEventsTypeCardToCard                 AccountActivityListResponseBookTransferTransactionEventsType = "CARD_TO_CARD"
	AccountActivityListResponseBookTransferTransactionEventsTypeDisburse                   AccountActivityListResponseBookTransferTransactionEventsType = "DISBURSE"
	AccountActivityListResponseBookTransferTransactionEventsTypeBillingError               AccountActivityListResponseBookTransferTransactionEventsType = "BILLING_ERROR"
	AccountActivityListResponseBookTransferTransactionEventsTypeLossWriteOff               AccountActivityListResponseBookTransferTransactionEventsType = "LOSS_WRITE_OFF"
	AccountActivityListResponseBookTransferTransactionEventsTypeExpiredCard                AccountActivityListResponseBookTransferTransactionEventsType = "EXPIRED_CARD"
	AccountActivityListResponseBookTransferTransactionEventsTypeEarlyDerecognition         AccountActivityListResponseBookTransferTransactionEventsType = "EARLY_DERECOGNITION"
	AccountActivityListResponseBookTransferTransactionEventsTypeEscheatment                AccountActivityListResponseBookTransferTransactionEventsType = "ESCHEATMENT"
	AccountActivityListResponseBookTransferTransactionEventsTypeInactivityFeeDown          AccountActivityListResponseBookTransferTransactionEventsType = "INACTIVITY_FEE_DOWN"
	AccountActivityListResponseBookTransferTransactionEventsTypeProvisionalCredit          AccountActivityListResponseBookTransferTransactionEventsType = "PROVISIONAL_CREDIT"
	AccountActivityListResponseBookTransferTransactionEventsTypeDisputeWon                 AccountActivityListResponseBookTransferTransactionEventsType = "DISPUTE_WON"
	AccountActivityListResponseBookTransferTransactionEventsTypeService                    AccountActivityListResponseBookTransferTransactionEventsType = "SERVICE"
	AccountActivityListResponseBookTransferTransactionEventsTypeTransfer                   AccountActivityListResponseBookTransferTransactionEventsType = "TRANSFER"
	AccountActivityListResponseBookTransferTransactionEventsTypeCollection                 AccountActivityListResponseBookTransferTransactionEventsType = "COLLECTION"
)

func (r AccountActivityListResponseBookTransferTransactionEventsType) IsKnown() bool {
	switch r {
	case AccountActivityListResponseBookTransferTransactionEventsTypeAtmWithdrawal, AccountActivityListResponseBookTransferTransactionEventsTypeAtmDecline, AccountActivityListResponseBookTransferTransactionEventsTypeInternationalAtmWithdrawal, AccountActivityListResponseBookTransferTransactionEventsTypeInactivity, AccountActivityListResponseBookTransferTransactionEventsTypeStatement, AccountActivityListResponseBookTransferTransactionEventsTypeMonthly, AccountActivityListResponseBookTransferTransactionEventsTypeQuarterly, AccountActivityListResponseBookTransferTransactionEventsTypeAnnual, AccountActivityListResponseBookTransferTransactionEventsTypeCustomerService, AccountActivityListResponseBookTransferTransactionEventsTypeAccountMaintenance, AccountActivityListResponseBookTransferTransactionEventsTypeAccountActivation, AccountActivityListResponseBookTransferTransactionEventsTypeAccountClosure, AccountActivityListResponseBookTransferTransactionEventsTypeCardReplacement, AccountActivityListResponseBookTransferTransactionEventsTypeCardDelivery, AccountActivityListResponseBookTransferTransactionEventsTypeCardCreate, AccountActivityListResponseBookTransferTransactionEventsTypeCurrencyConversion, AccountActivityListResponseBookTransferTransactionEventsTypeInterest, AccountActivityListResponseBookTransferTransactionEventsTypeLatePayment, AccountActivityListResponseBookTransferTransactionEventsTypeBillPayment, AccountActivityListResponseBookTransferTransactionEventsTypeCashBack, AccountActivityListResponseBookTransferTransactionEventsTypeAccountToAccount, AccountActivityListResponseBookTransferTransactionEventsTypeCardToCard, AccountActivityListResponseBookTransferTransactionEventsTypeDisburse, AccountActivityListResponseBookTransferTransactionEventsTypeBillingError, AccountActivityListResponseBookTransferTransactionEventsTypeLossWriteOff, AccountActivityListResponseBookTransferTransactionEventsTypeExpiredCard, AccountActivityListResponseBookTransferTransactionEventsTypeEarlyDerecognition, AccountActivityListResponseBookTransferTransactionEventsTypeEscheatment, AccountActivityListResponseBookTransferTransactionEventsTypeInactivityFeeDown, AccountActivityListResponseBookTransferTransactionEventsTypeProvisionalCredit, AccountActivityListResponseBookTransferTransactionEventsTypeDisputeWon, AccountActivityListResponseBookTransferTransactionEventsTypeService, AccountActivityListResponseBookTransferTransactionEventsTypeTransfer, AccountActivityListResponseBookTransferTransactionEventsTypeCollection:
		return true
	}
	return false
}

type AccountActivityListResponseBookTransferTransactionFamily string

const (
	AccountActivityListResponseBookTransferTransactionFamilyCard                AccountActivityListResponseBookTransferTransactionFamily = "CARD"
	AccountActivityListResponseBookTransferTransactionFamilyPayment             AccountActivityListResponseBookTransferTransactionFamily = "PAYMENT"
	AccountActivityListResponseBookTransferTransactionFamilyTransfer            AccountActivityListResponseBookTransferTransactionFamily = "TRANSFER"
	AccountActivityListResponseBookTransferTransactionFamilyInternal            AccountActivityListResponseBookTransferTransactionFamily = "INTERNAL"
	AccountActivityListResponseBookTransferTransactionFamilyExternalPayment     AccountActivityListResponseBookTransferTransactionFamily = "EXTERNAL_PAYMENT"
	AccountActivityListResponseBookTransferTransactionFamilyManagementOperation AccountActivityListResponseBookTransferTransactionFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityListResponseBookTransferTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityListResponseBookTransferTransactionFamilyCard, AccountActivityListResponseBookTransferTransactionFamilyPayment, AccountActivityListResponseBookTransferTransactionFamilyTransfer, AccountActivityListResponseBookTransferTransactionFamilyInternal, AccountActivityListResponseBookTransferTransactionFamilyExternalPayment, AccountActivityListResponseBookTransferTransactionFamilyManagementOperation:
		return true
	}
	return false
}

type AccountActivityListResponseBookTransferTransactionResult string

const (
	AccountActivityListResponseBookTransferTransactionResultApproved AccountActivityListResponseBookTransferTransactionResult = "APPROVED"
	AccountActivityListResponseBookTransferTransactionResultDeclined AccountActivityListResponseBookTransferTransactionResult = "DECLINED"
)

func (r AccountActivityListResponseBookTransferTransactionResult) IsKnown() bool {
	switch r {
	case AccountActivityListResponseBookTransferTransactionResultApproved, AccountActivityListResponseBookTransferTransactionResultDeclined:
		return true
	}
	return false
}

// The status of the transaction
type AccountActivityListResponseBookTransferTransactionStatus string

const (
	AccountActivityListResponseBookTransferTransactionStatusPending  AccountActivityListResponseBookTransferTransactionStatus = "PENDING"
	AccountActivityListResponseBookTransferTransactionStatusSettled  AccountActivityListResponseBookTransferTransactionStatus = "SETTLED"
	AccountActivityListResponseBookTransferTransactionStatusDeclined AccountActivityListResponseBookTransferTransactionStatus = "DECLINED"
	AccountActivityListResponseBookTransferTransactionStatusReversed AccountActivityListResponseBookTransferTransactionStatus = "REVERSED"
	AccountActivityListResponseBookTransferTransactionStatusCanceled AccountActivityListResponseBookTransferTransactionStatus = "CANCELED"
)

func (r AccountActivityListResponseBookTransferTransactionStatus) IsKnown() bool {
	switch r {
	case AccountActivityListResponseBookTransferTransactionStatusPending, AccountActivityListResponseBookTransferTransactionStatusSettled, AccountActivityListResponseBookTransferTransactionStatusDeclined, AccountActivityListResponseBookTransferTransactionStatusReversed, AccountActivityListResponseBookTransferTransactionStatusCanceled:
		return true
	}
	return false
}

type AccountActivityListResponseBookTransferTransactionTransactionSeries struct {
	RelatedTransactionEventToken string                                                                  `json:"related_transaction_event_token,required,nullable" format:"uuid"`
	RelatedTransactionToken      string                                                                  `json:"related_transaction_token,required,nullable" format:"uuid"`
	Type                         string                                                                  `json:"type,required"`
	JSON                         accountActivityListResponseBookTransferTransactionTransactionSeriesJSON `json:"-"`
}

// accountActivityListResponseBookTransferTransactionTransactionSeriesJSON contains
// the JSON metadata for the struct
// [AccountActivityListResponseBookTransferTransactionTransactionSeries]
type accountActivityListResponseBookTransferTransactionTransactionSeriesJSON struct {
	RelatedTransactionEventToken apijson.Field
	RelatedTransactionToken      apijson.Field
	Type                         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccountActivityListResponseBookTransferTransactionTransactionSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponseBookTransferTransactionTransactionSeriesJSON) RawJSON() string {
	return r.raw
}

// Base class for all transaction types in the ledger service
type AccountActivityListResponseCardTransaction struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time                                        `json:"created,required" format:"date-time"`
	Family  AccountActivityListResponseCardTransactionFamily `json:"family,required"`
	// The status of the transaction
	Status AccountActivityListResponseCardTransactionStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time                                      `json:"updated,required" format:"date-time"`
	JSON    accountActivityListResponseCardTransactionJSON `json:"-"`
	Transaction
}

// accountActivityListResponseCardTransactionJSON contains the JSON metadata for
// the struct [AccountActivityListResponseCardTransaction]
type accountActivityListResponseCardTransactionJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	Family      apijson.Field
	Status      apijson.Field
	Updated     apijson.Field
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

type AccountActivityListResponseCardTransactionFamily string

const (
	AccountActivityListResponseCardTransactionFamilyCard                AccountActivityListResponseCardTransactionFamily = "CARD"
	AccountActivityListResponseCardTransactionFamilyPayment             AccountActivityListResponseCardTransactionFamily = "PAYMENT"
	AccountActivityListResponseCardTransactionFamilyTransfer            AccountActivityListResponseCardTransactionFamily = "TRANSFER"
	AccountActivityListResponseCardTransactionFamilyInternal            AccountActivityListResponseCardTransactionFamily = "INTERNAL"
	AccountActivityListResponseCardTransactionFamilyExternalPayment     AccountActivityListResponseCardTransactionFamily = "EXTERNAL_PAYMENT"
	AccountActivityListResponseCardTransactionFamilyManagementOperation AccountActivityListResponseCardTransactionFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityListResponseCardTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityListResponseCardTransactionFamilyCard, AccountActivityListResponseCardTransactionFamilyPayment, AccountActivityListResponseCardTransactionFamilyTransfer, AccountActivityListResponseCardTransactionFamilyInternal, AccountActivityListResponseCardTransactionFamilyExternalPayment, AccountActivityListResponseCardTransactionFamilyManagementOperation:
		return true
	}
	return false
}

// The status of the transaction
type AccountActivityListResponseCardTransactionStatus string

const (
	AccountActivityListResponseCardTransactionStatusPending  AccountActivityListResponseCardTransactionStatus = "PENDING"
	AccountActivityListResponseCardTransactionStatusSettled  AccountActivityListResponseCardTransactionStatus = "SETTLED"
	AccountActivityListResponseCardTransactionStatusDeclined AccountActivityListResponseCardTransactionStatus = "DECLINED"
	AccountActivityListResponseCardTransactionStatusReversed AccountActivityListResponseCardTransactionStatus = "REVERSED"
	AccountActivityListResponseCardTransactionStatusCanceled AccountActivityListResponseCardTransactionStatus = "CANCELED"
)

func (r AccountActivityListResponseCardTransactionStatus) IsKnown() bool {
	switch r {
	case AccountActivityListResponseCardTransactionStatusPending, AccountActivityListResponseCardTransactionStatusSettled, AccountActivityListResponseCardTransactionStatusDeclined, AccountActivityListResponseCardTransactionStatusReversed, AccountActivityListResponseCardTransactionStatusCanceled:
		return true
	}
	return false
}

// Payment transaction
type AccountActivityListResponsePaymentTransaction struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// Transaction category
	Category AccountActivityListResponsePaymentTransactionCategory `json:"category,required"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Transaction descriptor
	Descriptor string `json:"descriptor,required"`
	// Transfer direction
	Direction AccountActivityListResponsePaymentTransactionDirection `json:"direction,required"`
	// List of transaction events
	Events []AccountActivityListResponsePaymentTransactionEvent `json:"events,required"`
	Family AccountActivityListResponsePaymentTransactionFamily  `json:"family,required"`
	// Financial account token
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Transfer method
	Method AccountActivityListResponsePaymentTransactionMethod `json:"method,required"`
	// Method-specific attributes
	MethodAttributes AccountActivityListResponsePaymentTransactionMethodAttributes `json:"method_attributes,required"`
	// Pending amount in cents
	PendingAmount int64 `json:"pending_amount,required"`
	// Related account tokens for the transaction
	RelatedAccountTokens AccountActivityListResponsePaymentTransactionRelatedAccountTokens `json:"related_account_tokens,required"`
	// Transaction result
	Result AccountActivityListResponsePaymentTransactionResult `json:"result,required"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount,required"`
	// Transaction source
	Source AccountActivityListResponsePaymentTransactionSource `json:"source,required"`
	// The status of the transaction
	Status AccountActivityListResponsePaymentTransactionStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Currency of the transaction in ISO 4217 format
	Currency string `json:"currency"`
	// Expected release date for the transaction
	ExpectedReleaseDate time.Time `json:"expected_release_date,nullable" format:"date"`
	// External bank account token
	ExternalBankAccountToken string                                            `json:"external_bank_account_token,nullable" format:"uuid"`
	Type                     AccountActivityListResponsePaymentTransactionType `json:"type"`
	// User-defined identifier
	UserDefinedID string                                            `json:"user_defined_id,nullable"`
	JSON          accountActivityListResponsePaymentTransactionJSON `json:"-"`
}

// accountActivityListResponsePaymentTransactionJSON contains the JSON metadata for
// the struct [AccountActivityListResponsePaymentTransaction]
type accountActivityListResponsePaymentTransactionJSON struct {
	Token                    apijson.Field
	Category                 apijson.Field
	Created                  apijson.Field
	Descriptor               apijson.Field
	Direction                apijson.Field
	Events                   apijson.Field
	Family                   apijson.Field
	FinancialAccountToken    apijson.Field
	Method                   apijson.Field
	MethodAttributes         apijson.Field
	PendingAmount            apijson.Field
	RelatedAccountTokens     apijson.Field
	Result                   apijson.Field
	SettledAmount            apijson.Field
	Source                   apijson.Field
	Status                   apijson.Field
	Updated                  apijson.Field
	Currency                 apijson.Field
	ExpectedReleaseDate      apijson.Field
	ExternalBankAccountToken apijson.Field
	Type                     apijson.Field
	UserDefinedID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountActivityListResponsePaymentTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponsePaymentTransactionJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityListResponsePaymentTransaction) implementsAccountActivityListResponse() {}

// Transaction category
type AccountActivityListResponsePaymentTransactionCategory string

const (
	AccountActivityListResponsePaymentTransactionCategoryACH                    AccountActivityListResponsePaymentTransactionCategory = "ACH"
	AccountActivityListResponsePaymentTransactionCategoryBalanceOrFunding       AccountActivityListResponsePaymentTransactionCategory = "BALANCE_OR_FUNDING"
	AccountActivityListResponsePaymentTransactionCategoryCard                   AccountActivityListResponsePaymentTransactionCategory = "CARD"
	AccountActivityListResponsePaymentTransactionCategoryExternalACH            AccountActivityListResponsePaymentTransactionCategory = "EXTERNAL_ACH"
	AccountActivityListResponsePaymentTransactionCategoryExternalCheck          AccountActivityListResponsePaymentTransactionCategory = "EXTERNAL_CHECK"
	AccountActivityListResponsePaymentTransactionCategoryExternalTransfer       AccountActivityListResponsePaymentTransactionCategory = "EXTERNAL_TRANSFER"
	AccountActivityListResponsePaymentTransactionCategoryExternalWire           AccountActivityListResponsePaymentTransactionCategory = "EXTERNAL_WIRE"
	AccountActivityListResponsePaymentTransactionCategoryManagementAdjustment   AccountActivityListResponsePaymentTransactionCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityListResponsePaymentTransactionCategoryManagementDispute      AccountActivityListResponsePaymentTransactionCategory = "MANAGEMENT_DISPUTE"
	AccountActivityListResponsePaymentTransactionCategoryManagementFee          AccountActivityListResponsePaymentTransactionCategory = "MANAGEMENT_FEE"
	AccountActivityListResponsePaymentTransactionCategoryManagementReward       AccountActivityListResponsePaymentTransactionCategory = "MANAGEMENT_REWARD"
	AccountActivityListResponsePaymentTransactionCategoryManagementDisbursement AccountActivityListResponsePaymentTransactionCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityListResponsePaymentTransactionCategoryProgramFunding         AccountActivityListResponsePaymentTransactionCategory = "PROGRAM_FUNDING"
)

func (r AccountActivityListResponsePaymentTransactionCategory) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionCategoryACH, AccountActivityListResponsePaymentTransactionCategoryBalanceOrFunding, AccountActivityListResponsePaymentTransactionCategoryCard, AccountActivityListResponsePaymentTransactionCategoryExternalACH, AccountActivityListResponsePaymentTransactionCategoryExternalCheck, AccountActivityListResponsePaymentTransactionCategoryExternalTransfer, AccountActivityListResponsePaymentTransactionCategoryExternalWire, AccountActivityListResponsePaymentTransactionCategoryManagementAdjustment, AccountActivityListResponsePaymentTransactionCategoryManagementDispute, AccountActivityListResponsePaymentTransactionCategoryManagementFee, AccountActivityListResponsePaymentTransactionCategoryManagementReward, AccountActivityListResponsePaymentTransactionCategoryManagementDisbursement, AccountActivityListResponsePaymentTransactionCategoryProgramFunding:
		return true
	}
	return false
}

// Transfer direction
type AccountActivityListResponsePaymentTransactionDirection string

const (
	AccountActivityListResponsePaymentTransactionDirectionCredit AccountActivityListResponsePaymentTransactionDirection = "CREDIT"
	AccountActivityListResponsePaymentTransactionDirectionDebit  AccountActivityListResponsePaymentTransactionDirection = "DEBIT"
)

func (r AccountActivityListResponsePaymentTransactionDirection) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionDirectionCredit, AccountActivityListResponsePaymentTransactionDirectionDebit:
		return true
	}
	return false
}

// Payment Event
type AccountActivityListResponsePaymentTransactionEvent struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount,required"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result AccountActivityListResponsePaymentTransactionEventsResult `json:"result,required"`
	// Event types:
	//
	//   - `ACH_ORIGINATION_INITIATED` - ACH origination received and pending
	//     approval/release from an ACH hold.
	//   - `ACH_ORIGINATION_REVIEWED` - ACH origination has completed the review process.
	//   - `ACH_ORIGINATION_CANCELLED` - ACH origination has been cancelled.
	//   - `ACH_ORIGINATION_PROCESSED` - ACH origination has been processed and sent to
	//     the Federal Reserve.
	//   - `ACH_ORIGINATION_SETTLED` - ACH origination has settled.
	//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
	//     available balance.
	//   - `ACH_RETURN_PROCESSED` - ACH origination returned by the Receiving Depository
	//     Financial Institution.
	//   - `ACH_RECEIPT_PROCESSED` - ACH receipt pending release from an ACH holder.
	//   - `ACH_RETURN_INITIATED` - ACH initiated return for a ACH receipt.
	//   - `ACH_RECEIPT_SETTLED` - ACH receipt funds have settled.
	//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
	//     balance.
	//   - `ACH_RETURN_SETTLED` - ACH receipt return settled by the Receiving Depository
	//     Financial Institution.
	Type AccountActivityListResponsePaymentTransactionEventsType `json:"type,required"`
	// More detailed reasons for the event
	DetailedResults []AccountActivityListResponsePaymentTransactionEventsDetailedResult `json:"detailed_results"`
	JSON            accountActivityListResponsePaymentTransactionEventJSON              `json:"-"`
}

// accountActivityListResponsePaymentTransactionEventJSON contains the JSON
// metadata for the struct [AccountActivityListResponsePaymentTransactionEvent]
type accountActivityListResponsePaymentTransactionEventJSON struct {
	Token           apijson.Field
	Amount          apijson.Field
	Created         apijson.Field
	Result          apijson.Field
	Type            apijson.Field
	DetailedResults apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountActivityListResponsePaymentTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponsePaymentTransactionEventJSON) RawJSON() string {
	return r.raw
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type AccountActivityListResponsePaymentTransactionEventsResult string

const (
	AccountActivityListResponsePaymentTransactionEventsResultApproved AccountActivityListResponsePaymentTransactionEventsResult = "APPROVED"
	AccountActivityListResponsePaymentTransactionEventsResultDeclined AccountActivityListResponsePaymentTransactionEventsResult = "DECLINED"
)

func (r AccountActivityListResponsePaymentTransactionEventsResult) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionEventsResultApproved, AccountActivityListResponsePaymentTransactionEventsResultDeclined:
		return true
	}
	return false
}

// Event types:
//
//   - `ACH_ORIGINATION_INITIATED` - ACH origination received and pending
//     approval/release from an ACH hold.
//   - `ACH_ORIGINATION_REVIEWED` - ACH origination has completed the review process.
//   - `ACH_ORIGINATION_CANCELLED` - ACH origination has been cancelled.
//   - `ACH_ORIGINATION_PROCESSED` - ACH origination has been processed and sent to
//     the Federal Reserve.
//   - `ACH_ORIGINATION_SETTLED` - ACH origination has settled.
//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
//     available balance.
//   - `ACH_RETURN_PROCESSED` - ACH origination returned by the Receiving Depository
//     Financial Institution.
//   - `ACH_RECEIPT_PROCESSED` - ACH receipt pending release from an ACH holder.
//   - `ACH_RETURN_INITIATED` - ACH initiated return for a ACH receipt.
//   - `ACH_RECEIPT_SETTLED` - ACH receipt funds have settled.
//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
//     balance.
//   - `ACH_RETURN_SETTLED` - ACH receipt return settled by the Receiving Depository
//     Financial Institution.
type AccountActivityListResponsePaymentTransactionEventsType string

const (
	AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationCancelled AccountActivityListResponsePaymentTransactionEventsType = "ACH_ORIGINATION_CANCELLED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationInitiated AccountActivityListResponsePaymentTransactionEventsType = "ACH_ORIGINATION_INITIATED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationProcessed AccountActivityListResponsePaymentTransactionEventsType = "ACH_ORIGINATION_PROCESSED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationSettled   AccountActivityListResponsePaymentTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationReleased  AccountActivityListResponsePaymentTransactionEventsType = "ACH_ORIGINATION_RELEASED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationReviewed  AccountActivityListResponsePaymentTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHReceiptProcessed     AccountActivityListResponsePaymentTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHReceiptSettled       AccountActivityListResponsePaymentTransactionEventsType = "ACH_RECEIPT_SETTLED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHReturnInitiated      AccountActivityListResponsePaymentTransactionEventsType = "ACH_RETURN_INITIATED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHReturnProcessed      AccountActivityListResponsePaymentTransactionEventsType = "ACH_RETURN_PROCESSED"
	AccountActivityListResponsePaymentTransactionEventsTypeACHReturnSettled        AccountActivityListResponsePaymentTransactionEventsType = "ACH_RETURN_SETTLED"
)

func (r AccountActivityListResponsePaymentTransactionEventsType) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationCancelled, AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationInitiated, AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationProcessed, AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationSettled, AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationReleased, AccountActivityListResponsePaymentTransactionEventsTypeACHOriginationReviewed, AccountActivityListResponsePaymentTransactionEventsTypeACHReceiptProcessed, AccountActivityListResponsePaymentTransactionEventsTypeACHReceiptSettled, AccountActivityListResponsePaymentTransactionEventsTypeACHReturnInitiated, AccountActivityListResponsePaymentTransactionEventsTypeACHReturnProcessed, AccountActivityListResponsePaymentTransactionEventsTypeACHReturnSettled:
		return true
	}
	return false
}

type AccountActivityListResponsePaymentTransactionEventsDetailedResult string

const (
	AccountActivityListResponsePaymentTransactionEventsDetailedResultApproved                        AccountActivityListResponsePaymentTransactionEventsDetailedResult = "APPROVED"
	AccountActivityListResponsePaymentTransactionEventsDetailedResultFundsInsufficient               AccountActivityListResponsePaymentTransactionEventsDetailedResult = "FUNDS_INSUFFICIENT"
	AccountActivityListResponsePaymentTransactionEventsDetailedResultAccountInvalid                  AccountActivityListResponsePaymentTransactionEventsDetailedResult = "ACCOUNT_INVALID"
	AccountActivityListResponsePaymentTransactionEventsDetailedResultProgramTransactionLimitExceeded AccountActivityListResponsePaymentTransactionEventsDetailedResult = "PROGRAM_TRANSACTION_LIMIT_EXCEEDED"
	AccountActivityListResponsePaymentTransactionEventsDetailedResultProgramDailyLimitExceeded       AccountActivityListResponsePaymentTransactionEventsDetailedResult = "PROGRAM_DAILY_LIMIT_EXCEEDED"
	AccountActivityListResponsePaymentTransactionEventsDetailedResultProgramMonthlyLimitExceeded     AccountActivityListResponsePaymentTransactionEventsDetailedResult = "PROGRAM_MONTHLY_LIMIT_EXCEEDED"
)

func (r AccountActivityListResponsePaymentTransactionEventsDetailedResult) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionEventsDetailedResultApproved, AccountActivityListResponsePaymentTransactionEventsDetailedResultFundsInsufficient, AccountActivityListResponsePaymentTransactionEventsDetailedResultAccountInvalid, AccountActivityListResponsePaymentTransactionEventsDetailedResultProgramTransactionLimitExceeded, AccountActivityListResponsePaymentTransactionEventsDetailedResultProgramDailyLimitExceeded, AccountActivityListResponsePaymentTransactionEventsDetailedResultProgramMonthlyLimitExceeded:
		return true
	}
	return false
}

type AccountActivityListResponsePaymentTransactionFamily string

const (
	AccountActivityListResponsePaymentTransactionFamilyCard                AccountActivityListResponsePaymentTransactionFamily = "CARD"
	AccountActivityListResponsePaymentTransactionFamilyPayment             AccountActivityListResponsePaymentTransactionFamily = "PAYMENT"
	AccountActivityListResponsePaymentTransactionFamilyTransfer            AccountActivityListResponsePaymentTransactionFamily = "TRANSFER"
	AccountActivityListResponsePaymentTransactionFamilyInternal            AccountActivityListResponsePaymentTransactionFamily = "INTERNAL"
	AccountActivityListResponsePaymentTransactionFamilyExternalPayment     AccountActivityListResponsePaymentTransactionFamily = "EXTERNAL_PAYMENT"
	AccountActivityListResponsePaymentTransactionFamilyManagementOperation AccountActivityListResponsePaymentTransactionFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityListResponsePaymentTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionFamilyCard, AccountActivityListResponsePaymentTransactionFamilyPayment, AccountActivityListResponsePaymentTransactionFamilyTransfer, AccountActivityListResponsePaymentTransactionFamilyInternal, AccountActivityListResponsePaymentTransactionFamilyExternalPayment, AccountActivityListResponsePaymentTransactionFamilyManagementOperation:
		return true
	}
	return false
}

// Transfer method
type AccountActivityListResponsePaymentTransactionMethod string

const (
	AccountActivityListResponsePaymentTransactionMethodACHNextDay AccountActivityListResponsePaymentTransactionMethod = "ACH_NEXT_DAY"
	AccountActivityListResponsePaymentTransactionMethodACHSameDay AccountActivityListResponsePaymentTransactionMethod = "ACH_SAME_DAY"
	AccountActivityListResponsePaymentTransactionMethodWire       AccountActivityListResponsePaymentTransactionMethod = "WIRE"
)

func (r AccountActivityListResponsePaymentTransactionMethod) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionMethodACHNextDay, AccountActivityListResponsePaymentTransactionMethodACHSameDay, AccountActivityListResponsePaymentTransactionMethodWire:
		return true
	}
	return false
}

// Method-specific attributes
type AccountActivityListResponsePaymentTransactionMethodAttributes struct {
	// Addenda information
	Addenda string `json:"addenda,nullable"`
	// Company ID for the ACH transaction
	CompanyID string           `json:"company_id,nullable"`
	Creditor  WirePartyDetails `json:"creditor"`
	Debtor    WirePartyDetails `json:"debtor"`
	// Point to point reference identifier, as assigned by the instructing party, used
	// for tracking the message through the Fedwire system
	MessageID string `json:"message_id,nullable"`
	// Receipt routing number
	ReceiptRoutingNumber string `json:"receipt_routing_number,nullable"`
	// Payment details or invoice reference
	RemittanceInformation string `json:"remittance_information,nullable"`
	// Number of retries attempted
	Retries int64 `json:"retries,nullable"`
	// Return reason code if the transaction was returned
	ReturnReasonCode string `json:"return_reason_code,nullable"`
	// SEC code for ACH transaction
	SecCode AccountActivityListResponsePaymentTransactionMethodAttributesSecCode `json:"sec_code"`
	// This field can have the runtime type of [[]string].
	TraceNumbers interface{} `json:"trace_numbers"`
	// Type of wire message
	WireMessageType string `json:"wire_message_type"`
	// Type of wire transfer
	WireNetwork AccountActivityListResponsePaymentTransactionMethodAttributesWireNetwork `json:"wire_network"`
	JSON        accountActivityListResponsePaymentTransactionMethodAttributesJSON        `json:"-"`
	union       AccountActivityListResponsePaymentTransactionMethodAttributesUnion
}

// accountActivityListResponsePaymentTransactionMethodAttributesJSON contains the
// JSON metadata for the struct
// [AccountActivityListResponsePaymentTransactionMethodAttributes]
type accountActivityListResponsePaymentTransactionMethodAttributesJSON struct {
	Addenda               apijson.Field
	CompanyID             apijson.Field
	Creditor              apijson.Field
	Debtor                apijson.Field
	MessageID             apijson.Field
	ReceiptRoutingNumber  apijson.Field
	RemittanceInformation apijson.Field
	Retries               apijson.Field
	ReturnReasonCode      apijson.Field
	SecCode               apijson.Field
	TraceNumbers          apijson.Field
	WireMessageType       apijson.Field
	WireNetwork           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r accountActivityListResponsePaymentTransactionMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r *AccountActivityListResponsePaymentTransactionMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	*r = AccountActivityListResponsePaymentTransactionMethodAttributes{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccountActivityListResponsePaymentTransactionMethodAttributesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributes],
// [AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributes].
func (r AccountActivityListResponsePaymentTransactionMethodAttributes) AsUnion() AccountActivityListResponsePaymentTransactionMethodAttributesUnion {
	return r.union
}

// Method-specific attributes
//
// Union satisfied by
// [AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributes]
// or
// [AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributes].
type AccountActivityListResponsePaymentTransactionMethodAttributesUnion interface {
	implementsAccountActivityListResponsePaymentTransactionMethodAttributes()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountActivityListResponsePaymentTransactionMethodAttributesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributes{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributes{}),
		},
	)
}

type AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributes struct {
	// SEC code for ACH transaction
	SecCode AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode `json:"sec_code,required"`
	// Addenda information
	Addenda string `json:"addenda,nullable"`
	// Company ID for the ACH transaction
	CompanyID string `json:"company_id,nullable"`
	// Receipt routing number
	ReceiptRoutingNumber string `json:"receipt_routing_number,nullable"`
	// Number of retries attempted
	Retries int64 `json:"retries,nullable"`
	// Return reason code if the transaction was returned
	ReturnReasonCode string `json:"return_reason_code,nullable"`
	// Trace numbers for the ACH transaction
	TraceNumbers []string                                                                             `json:"trace_numbers"`
	JSON         accountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesJSON `json:"-"`
}

// accountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesJSON
// contains the JSON metadata for the struct
// [AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributes]
type accountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesJSON struct {
	SecCode              apijson.Field
	Addenda              apijson.Field
	CompanyID            apijson.Field
	ReceiptRoutingNumber apijson.Field
	Retries              apijson.Field
	ReturnReasonCode     apijson.Field
	TraceNumbers         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributes) implementsAccountActivityListResponsePaymentTransactionMethodAttributes() {
}

// SEC code for ACH transaction
type AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode string

const (
	AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCcd AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "CCD"
	AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodePpd AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "PPD"
	AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeWeb AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "WEB"
	AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeTel AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "TEL"
	AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCie AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "CIE"
	AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCtx AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "CTX"
)

func (r AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCcd, AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodePpd, AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeWeb, AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeTel, AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCie, AccountActivityListResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCtx:
		return true
	}
	return false
}

type AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributes struct {
	// Type of wire transfer
	WireNetwork AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork `json:"wire_network,required"`
	Creditor    WirePartyDetails                                                                             `json:"creditor"`
	Debtor      WirePartyDetails                                                                             `json:"debtor"`
	// Point to point reference identifier, as assigned by the instructing party, used
	// for tracking the message through the Fedwire system
	MessageID string `json:"message_id,nullable"`
	// Payment details or invoice reference
	RemittanceInformation string `json:"remittance_information,nullable"`
	// Type of wire message
	WireMessageType string                                                                                `json:"wire_message_type"`
	JSON            accountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesJSON `json:"-"`
}

// accountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesJSON
// contains the JSON metadata for the struct
// [AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributes]
type accountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesJSON struct {
	WireNetwork           apijson.Field
	Creditor              apijson.Field
	Debtor                apijson.Field
	MessageID             apijson.Field
	RemittanceInformation apijson.Field
	WireMessageType       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributes) implementsAccountActivityListResponsePaymentTransactionMethodAttributes() {
}

// Type of wire transfer
type AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork string

const (
	AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetworkFedwire AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork = "FEDWIRE"
	AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetworkSwift   AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork = "SWIFT"
)

func (r AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetworkFedwire, AccountActivityListResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetworkSwift:
		return true
	}
	return false
}

// SEC code for ACH transaction
type AccountActivityListResponsePaymentTransactionMethodAttributesSecCode string

const (
	AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeCcd AccountActivityListResponsePaymentTransactionMethodAttributesSecCode = "CCD"
	AccountActivityListResponsePaymentTransactionMethodAttributesSecCodePpd AccountActivityListResponsePaymentTransactionMethodAttributesSecCode = "PPD"
	AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeWeb AccountActivityListResponsePaymentTransactionMethodAttributesSecCode = "WEB"
	AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeTel AccountActivityListResponsePaymentTransactionMethodAttributesSecCode = "TEL"
	AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeCie AccountActivityListResponsePaymentTransactionMethodAttributesSecCode = "CIE"
	AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeCtx AccountActivityListResponsePaymentTransactionMethodAttributesSecCode = "CTX"
)

func (r AccountActivityListResponsePaymentTransactionMethodAttributesSecCode) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeCcd, AccountActivityListResponsePaymentTransactionMethodAttributesSecCodePpd, AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeWeb, AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeTel, AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeCie, AccountActivityListResponsePaymentTransactionMethodAttributesSecCodeCtx:
		return true
	}
	return false
}

// Type of wire transfer
type AccountActivityListResponsePaymentTransactionMethodAttributesWireNetwork string

const (
	AccountActivityListResponsePaymentTransactionMethodAttributesWireNetworkFedwire AccountActivityListResponsePaymentTransactionMethodAttributesWireNetwork = "FEDWIRE"
	AccountActivityListResponsePaymentTransactionMethodAttributesWireNetworkSwift   AccountActivityListResponsePaymentTransactionMethodAttributesWireNetwork = "SWIFT"
)

func (r AccountActivityListResponsePaymentTransactionMethodAttributesWireNetwork) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionMethodAttributesWireNetworkFedwire, AccountActivityListResponsePaymentTransactionMethodAttributesWireNetworkSwift:
		return true
	}
	return false
}

// Related account tokens for the transaction
type AccountActivityListResponsePaymentTransactionRelatedAccountTokens struct {
	// Globally unique identifier for the account
	AccountToken string `json:"account_token,required,nullable" format:"uuid"`
	// Globally unique identifier for the business account
	BusinessAccountToken string                                                                `json:"business_account_token,required,nullable" format:"uuid"`
	JSON                 accountActivityListResponsePaymentTransactionRelatedAccountTokensJSON `json:"-"`
}

// accountActivityListResponsePaymentTransactionRelatedAccountTokensJSON contains
// the JSON metadata for the struct
// [AccountActivityListResponsePaymentTransactionRelatedAccountTokens]
type accountActivityListResponsePaymentTransactionRelatedAccountTokensJSON struct {
	AccountToken         apijson.Field
	BusinessAccountToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountActivityListResponsePaymentTransactionRelatedAccountTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityListResponsePaymentTransactionRelatedAccountTokensJSON) RawJSON() string {
	return r.raw
}

// Transaction result
type AccountActivityListResponsePaymentTransactionResult string

const (
	AccountActivityListResponsePaymentTransactionResultApproved AccountActivityListResponsePaymentTransactionResult = "APPROVED"
	AccountActivityListResponsePaymentTransactionResultDeclined AccountActivityListResponsePaymentTransactionResult = "DECLINED"
)

func (r AccountActivityListResponsePaymentTransactionResult) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionResultApproved, AccountActivityListResponsePaymentTransactionResultDeclined:
		return true
	}
	return false
}

// Transaction source
type AccountActivityListResponsePaymentTransactionSource string

const (
	AccountActivityListResponsePaymentTransactionSourceLithic   AccountActivityListResponsePaymentTransactionSource = "LITHIC"
	AccountActivityListResponsePaymentTransactionSourceExternal AccountActivityListResponsePaymentTransactionSource = "EXTERNAL"
	AccountActivityListResponsePaymentTransactionSourceCustomer AccountActivityListResponsePaymentTransactionSource = "CUSTOMER"
)

func (r AccountActivityListResponsePaymentTransactionSource) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionSourceLithic, AccountActivityListResponsePaymentTransactionSourceExternal, AccountActivityListResponsePaymentTransactionSourceCustomer:
		return true
	}
	return false
}

// The status of the transaction
type AccountActivityListResponsePaymentTransactionStatus string

const (
	AccountActivityListResponsePaymentTransactionStatusPending  AccountActivityListResponsePaymentTransactionStatus = "PENDING"
	AccountActivityListResponsePaymentTransactionStatusSettled  AccountActivityListResponsePaymentTransactionStatus = "SETTLED"
	AccountActivityListResponsePaymentTransactionStatusDeclined AccountActivityListResponsePaymentTransactionStatus = "DECLINED"
	AccountActivityListResponsePaymentTransactionStatusReversed AccountActivityListResponsePaymentTransactionStatus = "REVERSED"
	AccountActivityListResponsePaymentTransactionStatusCanceled AccountActivityListResponsePaymentTransactionStatus = "CANCELED"
)

func (r AccountActivityListResponsePaymentTransactionStatus) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionStatusPending, AccountActivityListResponsePaymentTransactionStatusSettled, AccountActivityListResponsePaymentTransactionStatusDeclined, AccountActivityListResponsePaymentTransactionStatusReversed, AccountActivityListResponsePaymentTransactionStatusCanceled:
		return true
	}
	return false
}

type AccountActivityListResponsePaymentTransactionType string

const (
	AccountActivityListResponsePaymentTransactionTypeOriginationCredit   AccountActivityListResponsePaymentTransactionType = "ORIGINATION_CREDIT"
	AccountActivityListResponsePaymentTransactionTypeOriginationDebit    AccountActivityListResponsePaymentTransactionType = "ORIGINATION_DEBIT"
	AccountActivityListResponsePaymentTransactionTypeReceiptCredit       AccountActivityListResponsePaymentTransactionType = "RECEIPT_CREDIT"
	AccountActivityListResponsePaymentTransactionTypeReceiptDebit        AccountActivityListResponsePaymentTransactionType = "RECEIPT_DEBIT"
	AccountActivityListResponsePaymentTransactionTypeWireInboundPayment  AccountActivityListResponsePaymentTransactionType = "WIRE_INBOUND_PAYMENT"
	AccountActivityListResponsePaymentTransactionTypeWireInboundAdmin    AccountActivityListResponsePaymentTransactionType = "WIRE_INBOUND_ADMIN"
	AccountActivityListResponsePaymentTransactionTypeWireOutboundPayment AccountActivityListResponsePaymentTransactionType = "WIRE_OUTBOUND_PAYMENT"
	AccountActivityListResponsePaymentTransactionTypeWireOutboundAdmin   AccountActivityListResponsePaymentTransactionType = "WIRE_OUTBOUND_ADMIN"
)

func (r AccountActivityListResponsePaymentTransactionType) IsKnown() bool {
	switch r {
	case AccountActivityListResponsePaymentTransactionTypeOriginationCredit, AccountActivityListResponsePaymentTransactionTypeOriginationDebit, AccountActivityListResponsePaymentTransactionTypeReceiptCredit, AccountActivityListResponsePaymentTransactionTypeReceiptDebit, AccountActivityListResponsePaymentTransactionTypeWireInboundPayment, AccountActivityListResponsePaymentTransactionTypeWireInboundAdmin, AccountActivityListResponsePaymentTransactionTypeWireOutboundPayment, AccountActivityListResponsePaymentTransactionTypeWireOutboundAdmin:
		return true
	}
	return false
}

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
)

func (r AccountActivityListResponseCategory) IsKnown() bool {
	switch r {
	case AccountActivityListResponseCategoryACH, AccountActivityListResponseCategoryBalanceOrFunding, AccountActivityListResponseCategoryCard, AccountActivityListResponseCategoryExternalACH, AccountActivityListResponseCategoryExternalCheck, AccountActivityListResponseCategoryExternalTransfer, AccountActivityListResponseCategoryExternalWire, AccountActivityListResponseCategoryManagementAdjustment, AccountActivityListResponseCategoryManagementDispute, AccountActivityListResponseCategoryManagementFee, AccountActivityListResponseCategoryManagementReward, AccountActivityListResponseCategoryManagementDisbursement, AccountActivityListResponseCategoryProgramFunding:
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

type AccountActivityListResponseFamily string

const (
	AccountActivityListResponseFamilyCard                AccountActivityListResponseFamily = "CARD"
	AccountActivityListResponseFamilyPayment             AccountActivityListResponseFamily = "PAYMENT"
	AccountActivityListResponseFamilyTransfer            AccountActivityListResponseFamily = "TRANSFER"
	AccountActivityListResponseFamilyInternal            AccountActivityListResponseFamily = "INTERNAL"
	AccountActivityListResponseFamilyExternalPayment     AccountActivityListResponseFamily = "EXTERNAL_PAYMENT"
	AccountActivityListResponseFamilyManagementOperation AccountActivityListResponseFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityListResponseFamily) IsKnown() bool {
	switch r {
	case AccountActivityListResponseFamilyCard, AccountActivityListResponseFamilyPayment, AccountActivityListResponseFamilyTransfer, AccountActivityListResponseFamilyInternal, AccountActivityListResponseFamilyExternalPayment, AccountActivityListResponseFamilyManagementOperation:
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
	case AccountActivityListResponseResultApproved, AccountActivityListResponseResultDeclined, AccountActivityListResponseResultAccountPaused, AccountActivityListResponseResultAccountStateTransactionFail, AccountActivityListResponseResultBankConnectionError, AccountActivityListResponseResultBankNotVerified, AccountActivityListResponseResultCardClosed, AccountActivityListResponseResultCardPaused, AccountActivityListResponseResultFraudAdvice, AccountActivityListResponseResultIgnoredTtlExpiry, AccountActivityListResponseResultInactiveAccount, AccountActivityListResponseResultIncorrectPin, AccountActivityListResponseResultInvalidCardDetails, AccountActivityListResponseResultInsufficientFunds, AccountActivityListResponseResultInsufficientFundsPreload, AccountActivityListResponseResultInvalidTransaction, AccountActivityListResponseResultMerchantBlacklist, AccountActivityListResponseResultOriginalNotFound, AccountActivityListResponseResultPreviouslyCompleted, AccountActivityListResponseResultSingleUseRecharged, AccountActivityListResponseResultSwitchInoperativeAdvice, AccountActivityListResponseResultUnauthorizedMerchant, AccountActivityListResponseResultUnknownHostTimeout, AccountActivityListResponseResultUserTransactionLimit:
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

// Response containing multiple transaction types
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
	// [[]AccountActivityGetTransactionResponseBookTransferTransactionEvent],
	// [[]TransactionEvent],
	// [[]AccountActivityGetTransactionResponsePaymentTransactionEvent],
	// [[]ExternalPaymentEvent], [[]ManagementOperationTransactionEvent].
	Events interface{} `json:"events"`
	// Expected release date for the transaction
	ExpectedReleaseDate time.Time `json:"expected_release_date,nullable" format:"date"`
	// External bank account token
	ExternalBankAccountToken string `json:"external_bank_account_token,nullable" format:"uuid"`
	// External identifier for the transaction
	ExternalID string `json:"external_id"`
	// External resource associated with the management operation
	ExternalResource ExternalResource                            `json:"external_resource,nullable"`
	Family           AccountActivityGetTransactionResponseFamily `json:"family"`
	// Financial account token associated with the transaction
	FinancialAccountToken string `json:"financial_account_token" format:"uuid"`
	// Source account token
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
	// This field can have the runtime type of
	// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributes].
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
	// This field can have the runtime type of
	// [AccountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokens].
	RelatedAccountTokens interface{} `json:"related_account_tokens"`
	// Transaction result
	Result AccountActivityGetTransactionResponseResult `json:"result"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount"`
	// Transaction source
	Source AccountActivityGetTransactionResponseSource `json:"source"`
	// Destination account token
	ToFinancialAccountToken string `json:"to_financial_account_token" format:"uuid"`
	// This field can have the runtime type of [TransactionTokenInfo].
	TokenInfo interface{} `json:"token_info"`
	// This field can have the runtime type of
	// [AccountActivityGetTransactionResponseBookTransferTransactionTransactionSeries],
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
// [AccountActivityGetTransactionResponseBookTransferTransaction],
// [AccountActivityGetTransactionResponseCardTransaction],
// [AccountActivityGetTransactionResponsePaymentTransaction], [ExternalPayment],
// [ManagementOperationTransaction].
func (r AccountActivityGetTransactionResponse) AsUnion() AccountActivityGetTransactionResponseUnion {
	return r.union
}

// Response containing multiple transaction types
//
// Union satisfied by [AccountActivityGetTransactionResponseFinancialTransaction],
// [AccountActivityGetTransactionResponseBookTransferTransaction],
// [AccountActivityGetTransactionResponseCardTransaction],
// [AccountActivityGetTransactionResponsePaymentTransaction], [ExternalPayment] or
// [ManagementOperationTransaction].
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
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseFinancialTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseFinancialTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseFinancialTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseFinancialTransaction{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseFinancialTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseBookTransferTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseBookTransferTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseBookTransferTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseBookTransferTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseBookTransferTransaction{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseBookTransferTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseCardTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseCardTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseCardTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseCardTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseCardTransaction{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponseCardTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponsePaymentTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponsePaymentTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponsePaymentTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponsePaymentTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponsePaymentTransaction{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountActivityGetTransactionResponsePaymentTransaction{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "EXTERNAL_PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExternalPayment{}),
			DiscriminatorValue: "MANAGEMENT_OPERATION",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "TRANSFER",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
			DiscriminatorValue: "INTERNAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagementOperationTransaction{}),
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
	Family AccountActivityGetTransactionResponseFinancialTransactionFamily  `json:"family,required"`
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
	case AccountActivityGetTransactionResponseFinancialTransactionCategoryACH, AccountActivityGetTransactionResponseFinancialTransactionCategoryBalanceOrFunding, AccountActivityGetTransactionResponseFinancialTransactionCategoryCard, AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalACH, AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalCheck, AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalTransfer, AccountActivityGetTransactionResponseFinancialTransactionCategoryExternalWire, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementAdjustment, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementDispute, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementFee, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementReward, AccountActivityGetTransactionResponseFinancialTransactionCategoryManagementDisbursement, AccountActivityGetTransactionResponseFinancialTransactionCategoryProgramFunding:
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
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationReviewed       AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationSettled        AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptProcessed          AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptSettled            AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RECEIPT_SETTLED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnInitiated           AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RETURN_INITIATED"
	AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnProcessed           AccountActivityGetTransactionResponseFinancialTransactionEventsType = "ACH_RETURN_PROCESSED"
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
	case AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationCancelled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationProcessed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationReviewed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHOriginationSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptProcessed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReceiptSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnProcessed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeACHReturnSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorization, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationAdvice, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationExpiry, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeAuthorizationReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBalanceInquiry, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBillingError, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeBillingErrorReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCardToCard, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCashBack, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCashBackReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeClearing, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCollection, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCorrectionCredit, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCorrectionDebit, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCreditAuthorization, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCreditAuthorizationAdvice, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCurrencyConversion, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeCurrencyConversionReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeDisputeWon, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHCanceled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHReversed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalACHSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckCanceled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckReversed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalCheckSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferCanceled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferReversed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalTransferSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireCanceled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireInitiated, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireReleased, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireReversed, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeExternalWireSettled, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeFinancialAuthorization, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeFinancialCreditAuthorization, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInterest, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInterestReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeInternalAdjustment, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLatePayment, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLatePaymentReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLossWriteOff, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeProvisionalCredit, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeProvisionalCreditReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeService, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturn, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeTransfer, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeTransferInsufficientFunds, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnedPayment, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeReturnedPaymentReversal, AccountActivityGetTransactionResponseFinancialTransactionEventsTypeLithicNetworkPayment:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponseFinancialTransactionFamily string

const (
	AccountActivityGetTransactionResponseFinancialTransactionFamilyCard                AccountActivityGetTransactionResponseFinancialTransactionFamily = "CARD"
	AccountActivityGetTransactionResponseFinancialTransactionFamilyPayment             AccountActivityGetTransactionResponseFinancialTransactionFamily = "PAYMENT"
	AccountActivityGetTransactionResponseFinancialTransactionFamilyTransfer            AccountActivityGetTransactionResponseFinancialTransactionFamily = "TRANSFER"
	AccountActivityGetTransactionResponseFinancialTransactionFamilyInternal            AccountActivityGetTransactionResponseFinancialTransactionFamily = "INTERNAL"
	AccountActivityGetTransactionResponseFinancialTransactionFamilyExternalPayment     AccountActivityGetTransactionResponseFinancialTransactionFamily = "EXTERNAL_PAYMENT"
	AccountActivityGetTransactionResponseFinancialTransactionFamilyManagementOperation AccountActivityGetTransactionResponseFinancialTransactionFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityGetTransactionResponseFinancialTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFinancialTransactionFamilyCard, AccountActivityGetTransactionResponseFinancialTransactionFamilyPayment, AccountActivityGetTransactionResponseFinancialTransactionFamilyTransfer, AccountActivityGetTransactionResponseFinancialTransactionFamilyInternal, AccountActivityGetTransactionResponseFinancialTransactionFamilyExternalPayment, AccountActivityGetTransactionResponseFinancialTransactionFamilyManagementOperation:
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

// Book transfer transaction
type AccountActivityGetTransactionResponseBookTransferTransaction struct {
	// Unique identifier for the transaction
	Token    string                                                               `json:"token,required" format:"uuid"`
	Category AccountActivityGetTransactionResponseBookTransferTransactionCategory `json:"category,required"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Currency of the transaction in ISO 4217 format
	Currency string `json:"currency,required"`
	// List of events associated with this book transfer
	Events []AccountActivityGetTransactionResponseBookTransferTransactionEvent `json:"events,required"`
	Family AccountActivityGetTransactionResponseBookTransferTransactionFamily  `json:"family,required"`
	// Source account token
	FromFinancialAccountToken string `json:"from_financial_account_token,required" format:"uuid"`
	// The pending amount of the transaction in cents
	PendingAmount int64                                                              `json:"pending_amount,required"`
	Result        AccountActivityGetTransactionResponseBookTransferTransactionResult `json:"result,required"`
	// The settled amount of the transaction in cents
	SettledAmount int64 `json:"settled_amount,required"`
	// The status of the transaction
	Status AccountActivityGetTransactionResponseBookTransferTransactionStatus `json:"status,required"`
	// Destination account token
	ToFinancialAccountToken string `json:"to_financial_account_token,required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// External identifier for the transaction
	ExternalID string `json:"external_id"`
	// External resource associated with the management operation
	ExternalResource  ExternalResource                                                              `json:"external_resource,nullable"`
	TransactionSeries AccountActivityGetTransactionResponseBookTransferTransactionTransactionSeries `json:"transaction_series,nullable"`
	JSON              accountActivityGetTransactionResponseBookTransferTransactionJSON              `json:"-"`
}

// accountActivityGetTransactionResponseBookTransferTransactionJSON contains the
// JSON metadata for the struct
// [AccountActivityGetTransactionResponseBookTransferTransaction]
type accountActivityGetTransactionResponseBookTransferTransactionJSON struct {
	Token                     apijson.Field
	Category                  apijson.Field
	Created                   apijson.Field
	Currency                  apijson.Field
	Events                    apijson.Field
	Family                    apijson.Field
	FromFinancialAccountToken apijson.Field
	PendingAmount             apijson.Field
	Result                    apijson.Field
	SettledAmount             apijson.Field
	Status                    apijson.Field
	ToFinancialAccountToken   apijson.Field
	Updated                   apijson.Field
	ExternalID                apijson.Field
	ExternalResource          apijson.Field
	TransactionSeries         apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponseBookTransferTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponseBookTransferTransactionJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityGetTransactionResponseBookTransferTransaction) implementsAccountActivityGetTransactionResponse() {
}

type AccountActivityGetTransactionResponseBookTransferTransactionCategory string

const (
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryACH                    AccountActivityGetTransactionResponseBookTransferTransactionCategory = "ACH"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryBalanceOrFunding       AccountActivityGetTransactionResponseBookTransferTransactionCategory = "BALANCE_OR_FUNDING"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryCard                   AccountActivityGetTransactionResponseBookTransferTransactionCategory = "CARD"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryExternalACH            AccountActivityGetTransactionResponseBookTransferTransactionCategory = "EXTERNAL_ACH"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryExternalCheck          AccountActivityGetTransactionResponseBookTransferTransactionCategory = "EXTERNAL_CHECK"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryExternalTransfer       AccountActivityGetTransactionResponseBookTransferTransactionCategory = "EXTERNAL_TRANSFER"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryExternalWire           AccountActivityGetTransactionResponseBookTransferTransactionCategory = "EXTERNAL_WIRE"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementAdjustment   AccountActivityGetTransactionResponseBookTransferTransactionCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementDispute      AccountActivityGetTransactionResponseBookTransferTransactionCategory = "MANAGEMENT_DISPUTE"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementFee          AccountActivityGetTransactionResponseBookTransferTransactionCategory = "MANAGEMENT_FEE"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementReward       AccountActivityGetTransactionResponseBookTransferTransactionCategory = "MANAGEMENT_REWARD"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementDisbursement AccountActivityGetTransactionResponseBookTransferTransactionCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionCategoryProgramFunding         AccountActivityGetTransactionResponseBookTransferTransactionCategory = "PROGRAM_FUNDING"
)

func (r AccountActivityGetTransactionResponseBookTransferTransactionCategory) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseBookTransferTransactionCategoryACH, AccountActivityGetTransactionResponseBookTransferTransactionCategoryBalanceOrFunding, AccountActivityGetTransactionResponseBookTransferTransactionCategoryCard, AccountActivityGetTransactionResponseBookTransferTransactionCategoryExternalACH, AccountActivityGetTransactionResponseBookTransferTransactionCategoryExternalCheck, AccountActivityGetTransactionResponseBookTransferTransactionCategoryExternalTransfer, AccountActivityGetTransactionResponseBookTransferTransactionCategoryExternalWire, AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementAdjustment, AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementDispute, AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementFee, AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementReward, AccountActivityGetTransactionResponseBookTransferTransactionCategoryManagementDisbursement, AccountActivityGetTransactionResponseBookTransferTransactionCategoryProgramFunding:
		return true
	}
	return false
}

// Book transfer Event
type AccountActivityGetTransactionResponseBookTransferTransactionEvent struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount,required"`
	// Date and time when the financial event occurred. UTC time zone.
	Created         time.Time                                                                         `json:"created,required" format:"date-time"`
	DetailedResults AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResults `json:"detailed_results,required"`
	// Memo for the transfer.
	Memo string `json:"memo,required"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result AccountActivityGetTransactionResponseBookTransferTransactionEventsResult `json:"result,required"`
	// The program specific subtype code for the specified category/type.
	Subtype string `json:"subtype,required"`
	// Type of the book transfer
	Type AccountActivityGetTransactionResponseBookTransferTransactionEventsType `json:"type,required"`
	JSON accountActivityGetTransactionResponseBookTransferTransactionEventJSON  `json:"-"`
}

// accountActivityGetTransactionResponseBookTransferTransactionEventJSON contains
// the JSON metadata for the struct
// [AccountActivityGetTransactionResponseBookTransferTransactionEvent]
type accountActivityGetTransactionResponseBookTransferTransactionEventJSON struct {
	Token           apijson.Field
	Amount          apijson.Field
	Created         apijson.Field
	DetailedResults apijson.Field
	Memo            apijson.Field
	Result          apijson.Field
	Subtype         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponseBookTransferTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponseBookTransferTransactionEventJSON) RawJSON() string {
	return r.raw
}

type AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResults string

const (
	AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResultsApproved          AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResults = "APPROVED"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResultsFundsInsufficient AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResults = "FUNDS_INSUFFICIENT"
)

func (r AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResults) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResultsApproved, AccountActivityGetTransactionResponseBookTransferTransactionEventsDetailedResultsFundsInsufficient:
		return true
	}
	return false
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type AccountActivityGetTransactionResponseBookTransferTransactionEventsResult string

const (
	AccountActivityGetTransactionResponseBookTransferTransactionEventsResultApproved AccountActivityGetTransactionResponseBookTransferTransactionEventsResult = "APPROVED"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsResultDeclined AccountActivityGetTransactionResponseBookTransferTransactionEventsResult = "DECLINED"
)

func (r AccountActivityGetTransactionResponseBookTransferTransactionEventsResult) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseBookTransferTransactionEventsResultApproved, AccountActivityGetTransactionResponseBookTransferTransactionEventsResultDeclined:
		return true
	}
	return false
}

// Type of the book transfer
type AccountActivityGetTransactionResponseBookTransferTransactionEventsType string

const (
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAtmWithdrawal              AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "ATM_WITHDRAWAL"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAtmDecline                 AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "ATM_DECLINE"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeInternationalAtmWithdrawal AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "INTERNATIONAL_ATM_WITHDRAWAL"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeInactivity                 AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "INACTIVITY"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeStatement                  AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "STATEMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeMonthly                    AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "MONTHLY"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeQuarterly                  AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "QUARTERLY"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAnnual                     AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "ANNUAL"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCustomerService            AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "CUSTOMER_SERVICE"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAccountMaintenance         AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "ACCOUNT_MAINTENANCE"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAccountActivation          AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "ACCOUNT_ACTIVATION"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAccountClosure             AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "ACCOUNT_CLOSURE"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCardReplacement            AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "CARD_REPLACEMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCardDelivery               AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "CARD_DELIVERY"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCardCreate                 AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "CARD_CREATE"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCurrencyConversion         AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "CURRENCY_CONVERSION"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeInterest                   AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "INTEREST"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeLatePayment                AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "LATE_PAYMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeBillPayment                AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "BILL_PAYMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCashBack                   AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "CASH_BACK"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAccountToAccount           AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "ACCOUNT_TO_ACCOUNT"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCardToCard                 AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "CARD_TO_CARD"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeDisburse                   AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "DISBURSE"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeBillingError               AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "BILLING_ERROR"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeLossWriteOff               AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "LOSS_WRITE_OFF"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeExpiredCard                AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "EXPIRED_CARD"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeEarlyDerecognition         AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "EARLY_DERECOGNITION"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeEscheatment                AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "ESCHEATMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeInactivityFeeDown          AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "INACTIVITY_FEE_DOWN"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeProvisionalCredit          AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "PROVISIONAL_CREDIT"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeDisputeWon                 AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "DISPUTE_WON"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeService                    AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "SERVICE"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeTransfer                   AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "TRANSFER"
	AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCollection                 AccountActivityGetTransactionResponseBookTransferTransactionEventsType = "COLLECTION"
)

func (r AccountActivityGetTransactionResponseBookTransferTransactionEventsType) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAtmWithdrawal, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAtmDecline, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeInternationalAtmWithdrawal, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeInactivity, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeStatement, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeMonthly, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeQuarterly, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAnnual, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCustomerService, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAccountMaintenance, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAccountActivation, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAccountClosure, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCardReplacement, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCardDelivery, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCardCreate, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCurrencyConversion, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeInterest, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeLatePayment, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeBillPayment, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCashBack, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeAccountToAccount, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCardToCard, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeDisburse, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeBillingError, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeLossWriteOff, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeExpiredCard, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeEarlyDerecognition, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeEscheatment, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeInactivityFeeDown, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeProvisionalCredit, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeDisputeWon, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeService, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeTransfer, AccountActivityGetTransactionResponseBookTransferTransactionEventsTypeCollection:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponseBookTransferTransactionFamily string

const (
	AccountActivityGetTransactionResponseBookTransferTransactionFamilyCard                AccountActivityGetTransactionResponseBookTransferTransactionFamily = "CARD"
	AccountActivityGetTransactionResponseBookTransferTransactionFamilyPayment             AccountActivityGetTransactionResponseBookTransferTransactionFamily = "PAYMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionFamilyTransfer            AccountActivityGetTransactionResponseBookTransferTransactionFamily = "TRANSFER"
	AccountActivityGetTransactionResponseBookTransferTransactionFamilyInternal            AccountActivityGetTransactionResponseBookTransferTransactionFamily = "INTERNAL"
	AccountActivityGetTransactionResponseBookTransferTransactionFamilyExternalPayment     AccountActivityGetTransactionResponseBookTransferTransactionFamily = "EXTERNAL_PAYMENT"
	AccountActivityGetTransactionResponseBookTransferTransactionFamilyManagementOperation AccountActivityGetTransactionResponseBookTransferTransactionFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityGetTransactionResponseBookTransferTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseBookTransferTransactionFamilyCard, AccountActivityGetTransactionResponseBookTransferTransactionFamilyPayment, AccountActivityGetTransactionResponseBookTransferTransactionFamilyTransfer, AccountActivityGetTransactionResponseBookTransferTransactionFamilyInternal, AccountActivityGetTransactionResponseBookTransferTransactionFamilyExternalPayment, AccountActivityGetTransactionResponseBookTransferTransactionFamilyManagementOperation:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponseBookTransferTransactionResult string

const (
	AccountActivityGetTransactionResponseBookTransferTransactionResultApproved AccountActivityGetTransactionResponseBookTransferTransactionResult = "APPROVED"
	AccountActivityGetTransactionResponseBookTransferTransactionResultDeclined AccountActivityGetTransactionResponseBookTransferTransactionResult = "DECLINED"
)

func (r AccountActivityGetTransactionResponseBookTransferTransactionResult) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseBookTransferTransactionResultApproved, AccountActivityGetTransactionResponseBookTransferTransactionResultDeclined:
		return true
	}
	return false
}

// The status of the transaction
type AccountActivityGetTransactionResponseBookTransferTransactionStatus string

const (
	AccountActivityGetTransactionResponseBookTransferTransactionStatusPending  AccountActivityGetTransactionResponseBookTransferTransactionStatus = "PENDING"
	AccountActivityGetTransactionResponseBookTransferTransactionStatusSettled  AccountActivityGetTransactionResponseBookTransferTransactionStatus = "SETTLED"
	AccountActivityGetTransactionResponseBookTransferTransactionStatusDeclined AccountActivityGetTransactionResponseBookTransferTransactionStatus = "DECLINED"
	AccountActivityGetTransactionResponseBookTransferTransactionStatusReversed AccountActivityGetTransactionResponseBookTransferTransactionStatus = "REVERSED"
	AccountActivityGetTransactionResponseBookTransferTransactionStatusCanceled AccountActivityGetTransactionResponseBookTransferTransactionStatus = "CANCELED"
)

func (r AccountActivityGetTransactionResponseBookTransferTransactionStatus) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseBookTransferTransactionStatusPending, AccountActivityGetTransactionResponseBookTransferTransactionStatusSettled, AccountActivityGetTransactionResponseBookTransferTransactionStatusDeclined, AccountActivityGetTransactionResponseBookTransferTransactionStatusReversed, AccountActivityGetTransactionResponseBookTransferTransactionStatusCanceled:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponseBookTransferTransactionTransactionSeries struct {
	RelatedTransactionEventToken string                                                                            `json:"related_transaction_event_token,required,nullable" format:"uuid"`
	RelatedTransactionToken      string                                                                            `json:"related_transaction_token,required,nullable" format:"uuid"`
	Type                         string                                                                            `json:"type,required"`
	JSON                         accountActivityGetTransactionResponseBookTransferTransactionTransactionSeriesJSON `json:"-"`
}

// accountActivityGetTransactionResponseBookTransferTransactionTransactionSeriesJSON
// contains the JSON metadata for the struct
// [AccountActivityGetTransactionResponseBookTransferTransactionTransactionSeries]
type accountActivityGetTransactionResponseBookTransferTransactionTransactionSeriesJSON struct {
	RelatedTransactionEventToken apijson.Field
	RelatedTransactionToken      apijson.Field
	Type                         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponseBookTransferTransactionTransactionSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponseBookTransferTransactionTransactionSeriesJSON) RawJSON() string {
	return r.raw
}

// Base class for all transaction types in the ledger service
type AccountActivityGetTransactionResponseCardTransaction struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time                                                  `json:"created,required" format:"date-time"`
	Family  AccountActivityGetTransactionResponseCardTransactionFamily `json:"family,required"`
	// The status of the transaction
	Status AccountActivityGetTransactionResponseCardTransactionStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time                                                `json:"updated,required" format:"date-time"`
	JSON    accountActivityGetTransactionResponseCardTransactionJSON `json:"-"`
	Transaction
}

// accountActivityGetTransactionResponseCardTransactionJSON contains the JSON
// metadata for the struct [AccountActivityGetTransactionResponseCardTransaction]
type accountActivityGetTransactionResponseCardTransactionJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	Family      apijson.Field
	Status      apijson.Field
	Updated     apijson.Field
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

type AccountActivityGetTransactionResponseCardTransactionFamily string

const (
	AccountActivityGetTransactionResponseCardTransactionFamilyCard                AccountActivityGetTransactionResponseCardTransactionFamily = "CARD"
	AccountActivityGetTransactionResponseCardTransactionFamilyPayment             AccountActivityGetTransactionResponseCardTransactionFamily = "PAYMENT"
	AccountActivityGetTransactionResponseCardTransactionFamilyTransfer            AccountActivityGetTransactionResponseCardTransactionFamily = "TRANSFER"
	AccountActivityGetTransactionResponseCardTransactionFamilyInternal            AccountActivityGetTransactionResponseCardTransactionFamily = "INTERNAL"
	AccountActivityGetTransactionResponseCardTransactionFamilyExternalPayment     AccountActivityGetTransactionResponseCardTransactionFamily = "EXTERNAL_PAYMENT"
	AccountActivityGetTransactionResponseCardTransactionFamilyManagementOperation AccountActivityGetTransactionResponseCardTransactionFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityGetTransactionResponseCardTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseCardTransactionFamilyCard, AccountActivityGetTransactionResponseCardTransactionFamilyPayment, AccountActivityGetTransactionResponseCardTransactionFamilyTransfer, AccountActivityGetTransactionResponseCardTransactionFamilyInternal, AccountActivityGetTransactionResponseCardTransactionFamilyExternalPayment, AccountActivityGetTransactionResponseCardTransactionFamilyManagementOperation:
		return true
	}
	return false
}

// The status of the transaction
type AccountActivityGetTransactionResponseCardTransactionStatus string

const (
	AccountActivityGetTransactionResponseCardTransactionStatusPending  AccountActivityGetTransactionResponseCardTransactionStatus = "PENDING"
	AccountActivityGetTransactionResponseCardTransactionStatusSettled  AccountActivityGetTransactionResponseCardTransactionStatus = "SETTLED"
	AccountActivityGetTransactionResponseCardTransactionStatusDeclined AccountActivityGetTransactionResponseCardTransactionStatus = "DECLINED"
	AccountActivityGetTransactionResponseCardTransactionStatusReversed AccountActivityGetTransactionResponseCardTransactionStatus = "REVERSED"
	AccountActivityGetTransactionResponseCardTransactionStatusCanceled AccountActivityGetTransactionResponseCardTransactionStatus = "CANCELED"
)

func (r AccountActivityGetTransactionResponseCardTransactionStatus) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseCardTransactionStatusPending, AccountActivityGetTransactionResponseCardTransactionStatusSettled, AccountActivityGetTransactionResponseCardTransactionStatusDeclined, AccountActivityGetTransactionResponseCardTransactionStatusReversed, AccountActivityGetTransactionResponseCardTransactionStatusCanceled:
		return true
	}
	return false
}

// Payment transaction
type AccountActivityGetTransactionResponsePaymentTransaction struct {
	// Unique identifier for the transaction
	Token string `json:"token,required" format:"uuid"`
	// Transaction category
	Category AccountActivityGetTransactionResponsePaymentTransactionCategory `json:"category,required"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Transaction descriptor
	Descriptor string `json:"descriptor,required"`
	// Transfer direction
	Direction AccountActivityGetTransactionResponsePaymentTransactionDirection `json:"direction,required"`
	// List of transaction events
	Events []AccountActivityGetTransactionResponsePaymentTransactionEvent `json:"events,required"`
	Family AccountActivityGetTransactionResponsePaymentTransactionFamily  `json:"family,required"`
	// Financial account token
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Transfer method
	Method AccountActivityGetTransactionResponsePaymentTransactionMethod `json:"method,required"`
	// Method-specific attributes
	MethodAttributes AccountActivityGetTransactionResponsePaymentTransactionMethodAttributes `json:"method_attributes,required"`
	// Pending amount in cents
	PendingAmount int64 `json:"pending_amount,required"`
	// Related account tokens for the transaction
	RelatedAccountTokens AccountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokens `json:"related_account_tokens,required"`
	// Transaction result
	Result AccountActivityGetTransactionResponsePaymentTransactionResult `json:"result,required"`
	// Settled amount in cents
	SettledAmount int64 `json:"settled_amount,required"`
	// Transaction source
	Source AccountActivityGetTransactionResponsePaymentTransactionSource `json:"source,required"`
	// The status of the transaction
	Status AccountActivityGetTransactionResponsePaymentTransactionStatus `json:"status,required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Currency of the transaction in ISO 4217 format
	Currency string `json:"currency"`
	// Expected release date for the transaction
	ExpectedReleaseDate time.Time `json:"expected_release_date,nullable" format:"date"`
	// External bank account token
	ExternalBankAccountToken string                                                      `json:"external_bank_account_token,nullable" format:"uuid"`
	Type                     AccountActivityGetTransactionResponsePaymentTransactionType `json:"type"`
	// User-defined identifier
	UserDefinedID string                                                      `json:"user_defined_id,nullable"`
	JSON          accountActivityGetTransactionResponsePaymentTransactionJSON `json:"-"`
}

// accountActivityGetTransactionResponsePaymentTransactionJSON contains the JSON
// metadata for the struct
// [AccountActivityGetTransactionResponsePaymentTransaction]
type accountActivityGetTransactionResponsePaymentTransactionJSON struct {
	Token                    apijson.Field
	Category                 apijson.Field
	Created                  apijson.Field
	Descriptor               apijson.Field
	Direction                apijson.Field
	Events                   apijson.Field
	Family                   apijson.Field
	FinancialAccountToken    apijson.Field
	Method                   apijson.Field
	MethodAttributes         apijson.Field
	PendingAmount            apijson.Field
	RelatedAccountTokens     apijson.Field
	Result                   apijson.Field
	SettledAmount            apijson.Field
	Source                   apijson.Field
	Status                   apijson.Field
	Updated                  apijson.Field
	Currency                 apijson.Field
	ExpectedReleaseDate      apijson.Field
	ExternalBankAccountToken apijson.Field
	Type                     apijson.Field
	UserDefinedID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponsePaymentTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponsePaymentTransactionJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityGetTransactionResponsePaymentTransaction) implementsAccountActivityGetTransactionResponse() {
}

// Transaction category
type AccountActivityGetTransactionResponsePaymentTransactionCategory string

const (
	AccountActivityGetTransactionResponsePaymentTransactionCategoryACH                    AccountActivityGetTransactionResponsePaymentTransactionCategory = "ACH"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryBalanceOrFunding       AccountActivityGetTransactionResponsePaymentTransactionCategory = "BALANCE_OR_FUNDING"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryCard                   AccountActivityGetTransactionResponsePaymentTransactionCategory = "CARD"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryExternalACH            AccountActivityGetTransactionResponsePaymentTransactionCategory = "EXTERNAL_ACH"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryExternalCheck          AccountActivityGetTransactionResponsePaymentTransactionCategory = "EXTERNAL_CHECK"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryExternalTransfer       AccountActivityGetTransactionResponsePaymentTransactionCategory = "EXTERNAL_TRANSFER"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryExternalWire           AccountActivityGetTransactionResponsePaymentTransactionCategory = "EXTERNAL_WIRE"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementAdjustment   AccountActivityGetTransactionResponsePaymentTransactionCategory = "MANAGEMENT_ADJUSTMENT"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementDispute      AccountActivityGetTransactionResponsePaymentTransactionCategory = "MANAGEMENT_DISPUTE"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementFee          AccountActivityGetTransactionResponsePaymentTransactionCategory = "MANAGEMENT_FEE"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementReward       AccountActivityGetTransactionResponsePaymentTransactionCategory = "MANAGEMENT_REWARD"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementDisbursement AccountActivityGetTransactionResponsePaymentTransactionCategory = "MANAGEMENT_DISBURSEMENT"
	AccountActivityGetTransactionResponsePaymentTransactionCategoryProgramFunding         AccountActivityGetTransactionResponsePaymentTransactionCategory = "PROGRAM_FUNDING"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionCategory) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionCategoryACH, AccountActivityGetTransactionResponsePaymentTransactionCategoryBalanceOrFunding, AccountActivityGetTransactionResponsePaymentTransactionCategoryCard, AccountActivityGetTransactionResponsePaymentTransactionCategoryExternalACH, AccountActivityGetTransactionResponsePaymentTransactionCategoryExternalCheck, AccountActivityGetTransactionResponsePaymentTransactionCategoryExternalTransfer, AccountActivityGetTransactionResponsePaymentTransactionCategoryExternalWire, AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementAdjustment, AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementDispute, AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementFee, AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementReward, AccountActivityGetTransactionResponsePaymentTransactionCategoryManagementDisbursement, AccountActivityGetTransactionResponsePaymentTransactionCategoryProgramFunding:
		return true
	}
	return false
}

// Transfer direction
type AccountActivityGetTransactionResponsePaymentTransactionDirection string

const (
	AccountActivityGetTransactionResponsePaymentTransactionDirectionCredit AccountActivityGetTransactionResponsePaymentTransactionDirection = "CREDIT"
	AccountActivityGetTransactionResponsePaymentTransactionDirectionDebit  AccountActivityGetTransactionResponsePaymentTransactionDirection = "DEBIT"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionDirection) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionDirectionCredit, AccountActivityGetTransactionResponsePaymentTransactionDirectionDebit:
		return true
	}
	return false
}

// Payment Event
type AccountActivityGetTransactionResponsePaymentTransactionEvent struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount,required"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result AccountActivityGetTransactionResponsePaymentTransactionEventsResult `json:"result,required"`
	// Event types:
	//
	//   - `ACH_ORIGINATION_INITIATED` - ACH origination received and pending
	//     approval/release from an ACH hold.
	//   - `ACH_ORIGINATION_REVIEWED` - ACH origination has completed the review process.
	//   - `ACH_ORIGINATION_CANCELLED` - ACH origination has been cancelled.
	//   - `ACH_ORIGINATION_PROCESSED` - ACH origination has been processed and sent to
	//     the Federal Reserve.
	//   - `ACH_ORIGINATION_SETTLED` - ACH origination has settled.
	//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
	//     available balance.
	//   - `ACH_RETURN_PROCESSED` - ACH origination returned by the Receiving Depository
	//     Financial Institution.
	//   - `ACH_RECEIPT_PROCESSED` - ACH receipt pending release from an ACH holder.
	//   - `ACH_RETURN_INITIATED` - ACH initiated return for a ACH receipt.
	//   - `ACH_RECEIPT_SETTLED` - ACH receipt funds have settled.
	//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
	//     balance.
	//   - `ACH_RETURN_SETTLED` - ACH receipt return settled by the Receiving Depository
	//     Financial Institution.
	Type AccountActivityGetTransactionResponsePaymentTransactionEventsType `json:"type,required"`
	// More detailed reasons for the event
	DetailedResults []AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult `json:"detailed_results"`
	JSON            accountActivityGetTransactionResponsePaymentTransactionEventJSON              `json:"-"`
}

// accountActivityGetTransactionResponsePaymentTransactionEventJSON contains the
// JSON metadata for the struct
// [AccountActivityGetTransactionResponsePaymentTransactionEvent]
type accountActivityGetTransactionResponsePaymentTransactionEventJSON struct {
	Token           apijson.Field
	Amount          apijson.Field
	Created         apijson.Field
	Result          apijson.Field
	Type            apijson.Field
	DetailedResults apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponsePaymentTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponsePaymentTransactionEventJSON) RawJSON() string {
	return r.raw
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type AccountActivityGetTransactionResponsePaymentTransactionEventsResult string

const (
	AccountActivityGetTransactionResponsePaymentTransactionEventsResultApproved AccountActivityGetTransactionResponsePaymentTransactionEventsResult = "APPROVED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsResultDeclined AccountActivityGetTransactionResponsePaymentTransactionEventsResult = "DECLINED"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionEventsResult) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionEventsResultApproved, AccountActivityGetTransactionResponsePaymentTransactionEventsResultDeclined:
		return true
	}
	return false
}

// Event types:
//
//   - `ACH_ORIGINATION_INITIATED` - ACH origination received and pending
//     approval/release from an ACH hold.
//   - `ACH_ORIGINATION_REVIEWED` - ACH origination has completed the review process.
//   - `ACH_ORIGINATION_CANCELLED` - ACH origination has been cancelled.
//   - `ACH_ORIGINATION_PROCESSED` - ACH origination has been processed and sent to
//     the Federal Reserve.
//   - `ACH_ORIGINATION_SETTLED` - ACH origination has settled.
//   - `ACH_ORIGINATION_RELEASED` - ACH origination released from pending to
//     available balance.
//   - `ACH_RETURN_PROCESSED` - ACH origination returned by the Receiving Depository
//     Financial Institution.
//   - `ACH_RECEIPT_PROCESSED` - ACH receipt pending release from an ACH holder.
//   - `ACH_RETURN_INITIATED` - ACH initiated return for a ACH receipt.
//   - `ACH_RECEIPT_SETTLED` - ACH receipt funds have settled.
//   - `ACH_RECEIPT_RELEASED` - ACH receipt released from pending to available
//     balance.
//   - `ACH_RETURN_SETTLED` - ACH receipt return settled by the Receiving Depository
//     Financial Institution.
type AccountActivityGetTransactionResponsePaymentTransactionEventsType string

const (
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationCancelled AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_ORIGINATION_CANCELLED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationInitiated AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_ORIGINATION_INITIATED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationProcessed AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_ORIGINATION_PROCESSED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationSettled   AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_ORIGINATION_SETTLED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationReleased  AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_ORIGINATION_RELEASED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationReviewed  AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_ORIGINATION_REVIEWED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReceiptProcessed     AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_RECEIPT_PROCESSED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReceiptSettled       AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_RECEIPT_SETTLED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReturnInitiated      AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_RETURN_INITIATED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReturnProcessed      AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_RETURN_PROCESSED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReturnSettled        AccountActivityGetTransactionResponsePaymentTransactionEventsType = "ACH_RETURN_SETTLED"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionEventsType) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationCancelled, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationInitiated, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationProcessed, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationSettled, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationReleased, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHOriginationReviewed, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReceiptProcessed, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReceiptSettled, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReturnInitiated, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReturnProcessed, AccountActivityGetTransactionResponsePaymentTransactionEventsTypeACHReturnSettled:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult string

const (
	AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultApproved                        AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult = "APPROVED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultFundsInsufficient               AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult = "FUNDS_INSUFFICIENT"
	AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultAccountInvalid                  AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult = "ACCOUNT_INVALID"
	AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultProgramTransactionLimitExceeded AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult = "PROGRAM_TRANSACTION_LIMIT_EXCEEDED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultProgramDailyLimitExceeded       AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult = "PROGRAM_DAILY_LIMIT_EXCEEDED"
	AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultProgramMonthlyLimitExceeded     AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult = "PROGRAM_MONTHLY_LIMIT_EXCEEDED"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResult) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultApproved, AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultFundsInsufficient, AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultAccountInvalid, AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultProgramTransactionLimitExceeded, AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultProgramDailyLimitExceeded, AccountActivityGetTransactionResponsePaymentTransactionEventsDetailedResultProgramMonthlyLimitExceeded:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponsePaymentTransactionFamily string

const (
	AccountActivityGetTransactionResponsePaymentTransactionFamilyCard                AccountActivityGetTransactionResponsePaymentTransactionFamily = "CARD"
	AccountActivityGetTransactionResponsePaymentTransactionFamilyPayment             AccountActivityGetTransactionResponsePaymentTransactionFamily = "PAYMENT"
	AccountActivityGetTransactionResponsePaymentTransactionFamilyTransfer            AccountActivityGetTransactionResponsePaymentTransactionFamily = "TRANSFER"
	AccountActivityGetTransactionResponsePaymentTransactionFamilyInternal            AccountActivityGetTransactionResponsePaymentTransactionFamily = "INTERNAL"
	AccountActivityGetTransactionResponsePaymentTransactionFamilyExternalPayment     AccountActivityGetTransactionResponsePaymentTransactionFamily = "EXTERNAL_PAYMENT"
	AccountActivityGetTransactionResponsePaymentTransactionFamilyManagementOperation AccountActivityGetTransactionResponsePaymentTransactionFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionFamily) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionFamilyCard, AccountActivityGetTransactionResponsePaymentTransactionFamilyPayment, AccountActivityGetTransactionResponsePaymentTransactionFamilyTransfer, AccountActivityGetTransactionResponsePaymentTransactionFamilyInternal, AccountActivityGetTransactionResponsePaymentTransactionFamilyExternalPayment, AccountActivityGetTransactionResponsePaymentTransactionFamilyManagementOperation:
		return true
	}
	return false
}

// Transfer method
type AccountActivityGetTransactionResponsePaymentTransactionMethod string

const (
	AccountActivityGetTransactionResponsePaymentTransactionMethodACHNextDay AccountActivityGetTransactionResponsePaymentTransactionMethod = "ACH_NEXT_DAY"
	AccountActivityGetTransactionResponsePaymentTransactionMethodACHSameDay AccountActivityGetTransactionResponsePaymentTransactionMethod = "ACH_SAME_DAY"
	AccountActivityGetTransactionResponsePaymentTransactionMethodWire       AccountActivityGetTransactionResponsePaymentTransactionMethod = "WIRE"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionMethod) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionMethodACHNextDay, AccountActivityGetTransactionResponsePaymentTransactionMethodACHSameDay, AccountActivityGetTransactionResponsePaymentTransactionMethodWire:
		return true
	}
	return false
}

// Method-specific attributes
type AccountActivityGetTransactionResponsePaymentTransactionMethodAttributes struct {
	// Addenda information
	Addenda string `json:"addenda,nullable"`
	// Company ID for the ACH transaction
	CompanyID string           `json:"company_id,nullable"`
	Creditor  WirePartyDetails `json:"creditor"`
	Debtor    WirePartyDetails `json:"debtor"`
	// Point to point reference identifier, as assigned by the instructing party, used
	// for tracking the message through the Fedwire system
	MessageID string `json:"message_id,nullable"`
	// Receipt routing number
	ReceiptRoutingNumber string `json:"receipt_routing_number,nullable"`
	// Payment details or invoice reference
	RemittanceInformation string `json:"remittance_information,nullable"`
	// Number of retries attempted
	Retries int64 `json:"retries,nullable"`
	// Return reason code if the transaction was returned
	ReturnReasonCode string `json:"return_reason_code,nullable"`
	// SEC code for ACH transaction
	SecCode AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode `json:"sec_code"`
	// This field can have the runtime type of [[]string].
	TraceNumbers interface{} `json:"trace_numbers"`
	// Type of wire message
	WireMessageType string `json:"wire_message_type"`
	// Type of wire transfer
	WireNetwork AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetwork `json:"wire_network"`
	JSON        accountActivityGetTransactionResponsePaymentTransactionMethodAttributesJSON        `json:"-"`
	union       AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesUnion
}

// accountActivityGetTransactionResponsePaymentTransactionMethodAttributesJSON
// contains the JSON metadata for the struct
// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributes]
type accountActivityGetTransactionResponsePaymentTransactionMethodAttributesJSON struct {
	Addenda               apijson.Field
	CompanyID             apijson.Field
	Creditor              apijson.Field
	Debtor                apijson.Field
	MessageID             apijson.Field
	ReceiptRoutingNumber  apijson.Field
	RemittanceInformation apijson.Field
	Retries               apijson.Field
	ReturnReasonCode      apijson.Field
	SecCode               apijson.Field
	TraceNumbers          apijson.Field
	WireMessageType       apijson.Field
	WireNetwork           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r accountActivityGetTransactionResponsePaymentTransactionMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r *AccountActivityGetTransactionResponsePaymentTransactionMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	*r = AccountActivityGetTransactionResponsePaymentTransactionMethodAttributes{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributes],
// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributes].
func (r AccountActivityGetTransactionResponsePaymentTransactionMethodAttributes) AsUnion() AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesUnion {
	return r.union
}

// Method-specific attributes
//
// Union satisfied by
// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributes]
// or
// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributes].
type AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesUnion interface {
	implementsAccountActivityGetTransactionResponsePaymentTransactionMethodAttributes()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributes{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributes{}),
		},
	)
}

type AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributes struct {
	// SEC code for ACH transaction
	SecCode AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode `json:"sec_code,required"`
	// Addenda information
	Addenda string `json:"addenda,nullable"`
	// Company ID for the ACH transaction
	CompanyID string `json:"company_id,nullable"`
	// Receipt routing number
	ReceiptRoutingNumber string `json:"receipt_routing_number,nullable"`
	// Number of retries attempted
	Retries int64 `json:"retries,nullable"`
	// Return reason code if the transaction was returned
	ReturnReasonCode string `json:"return_reason_code,nullable"`
	// Trace numbers for the ACH transaction
	TraceNumbers []string                                                                                       `json:"trace_numbers"`
	JSON         accountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesJSON `json:"-"`
}

// accountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesJSON
// contains the JSON metadata for the struct
// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributes]
type accountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesJSON struct {
	SecCode              apijson.Field
	Addenda              apijson.Field
	CompanyID            apijson.Field
	ReceiptRoutingNumber apijson.Field
	Retries              apijson.Field
	ReturnReasonCode     apijson.Field
	TraceNumbers         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributes) implementsAccountActivityGetTransactionResponsePaymentTransactionMethodAttributes() {
}

// SEC code for ACH transaction
type AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode string

const (
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCcd AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "CCD"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodePpd AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "PPD"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeWeb AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "WEB"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeTel AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "TEL"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCie AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "CIE"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCtx AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode = "CTX"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCode) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCcd, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodePpd, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeWeb, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeTel, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCie, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesACHMethodAttributesSecCodeCtx:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributes struct {
	// Type of wire transfer
	WireNetwork AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork `json:"wire_network,required"`
	Creditor    WirePartyDetails                                                                                       `json:"creditor"`
	Debtor      WirePartyDetails                                                                                       `json:"debtor"`
	// Point to point reference identifier, as assigned by the instructing party, used
	// for tracking the message through the Fedwire system
	MessageID string `json:"message_id,nullable"`
	// Payment details or invoice reference
	RemittanceInformation string `json:"remittance_information,nullable"`
	// Type of wire message
	WireMessageType string                                                                                          `json:"wire_message_type"`
	JSON            accountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesJSON `json:"-"`
}

// accountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesJSON
// contains the JSON metadata for the struct
// [AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributes]
type accountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesJSON struct {
	WireNetwork           apijson.Field
	Creditor              apijson.Field
	Debtor                apijson.Field
	MessageID             apijson.Field
	RemittanceInformation apijson.Field
	WireMessageType       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesJSON) RawJSON() string {
	return r.raw
}

func (r AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributes) implementsAccountActivityGetTransactionResponsePaymentTransactionMethodAttributes() {
}

// Type of wire transfer
type AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork string

const (
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetworkFedwire AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork = "FEDWIRE"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetworkSwift   AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork = "SWIFT"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetwork) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetworkFedwire, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireMethodAttributesWireNetworkSwift:
		return true
	}
	return false
}

// SEC code for ACH transaction
type AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode string

const (
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeCcd AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode = "CCD"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodePpd AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode = "PPD"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeWeb AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode = "WEB"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeTel AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode = "TEL"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeCie AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode = "CIE"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeCtx AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode = "CTX"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCode) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeCcd, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodePpd, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeWeb, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeTel, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeCie, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesSecCodeCtx:
		return true
	}
	return false
}

// Type of wire transfer
type AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetwork string

const (
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetworkFedwire AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetwork = "FEDWIRE"
	AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetworkSwift   AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetwork = "SWIFT"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetwork) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetworkFedwire, AccountActivityGetTransactionResponsePaymentTransactionMethodAttributesWireNetworkSwift:
		return true
	}
	return false
}

// Related account tokens for the transaction
type AccountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokens struct {
	// Globally unique identifier for the account
	AccountToken string `json:"account_token,required,nullable" format:"uuid"`
	// Globally unique identifier for the business account
	BusinessAccountToken string                                                                          `json:"business_account_token,required,nullable" format:"uuid"`
	JSON                 accountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokensJSON `json:"-"`
}

// accountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokensJSON
// contains the JSON metadata for the struct
// [AccountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokens]
type accountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokensJSON struct {
	AccountToken         apijson.Field
	BusinessAccountToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountActivityGetTransactionResponsePaymentTransactionRelatedAccountTokensJSON) RawJSON() string {
	return r.raw
}

// Transaction result
type AccountActivityGetTransactionResponsePaymentTransactionResult string

const (
	AccountActivityGetTransactionResponsePaymentTransactionResultApproved AccountActivityGetTransactionResponsePaymentTransactionResult = "APPROVED"
	AccountActivityGetTransactionResponsePaymentTransactionResultDeclined AccountActivityGetTransactionResponsePaymentTransactionResult = "DECLINED"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionResult) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionResultApproved, AccountActivityGetTransactionResponsePaymentTransactionResultDeclined:
		return true
	}
	return false
}

// Transaction source
type AccountActivityGetTransactionResponsePaymentTransactionSource string

const (
	AccountActivityGetTransactionResponsePaymentTransactionSourceLithic   AccountActivityGetTransactionResponsePaymentTransactionSource = "LITHIC"
	AccountActivityGetTransactionResponsePaymentTransactionSourceExternal AccountActivityGetTransactionResponsePaymentTransactionSource = "EXTERNAL"
	AccountActivityGetTransactionResponsePaymentTransactionSourceCustomer AccountActivityGetTransactionResponsePaymentTransactionSource = "CUSTOMER"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionSource) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionSourceLithic, AccountActivityGetTransactionResponsePaymentTransactionSourceExternal, AccountActivityGetTransactionResponsePaymentTransactionSourceCustomer:
		return true
	}
	return false
}

// The status of the transaction
type AccountActivityGetTransactionResponsePaymentTransactionStatus string

const (
	AccountActivityGetTransactionResponsePaymentTransactionStatusPending  AccountActivityGetTransactionResponsePaymentTransactionStatus = "PENDING"
	AccountActivityGetTransactionResponsePaymentTransactionStatusSettled  AccountActivityGetTransactionResponsePaymentTransactionStatus = "SETTLED"
	AccountActivityGetTransactionResponsePaymentTransactionStatusDeclined AccountActivityGetTransactionResponsePaymentTransactionStatus = "DECLINED"
	AccountActivityGetTransactionResponsePaymentTransactionStatusReversed AccountActivityGetTransactionResponsePaymentTransactionStatus = "REVERSED"
	AccountActivityGetTransactionResponsePaymentTransactionStatusCanceled AccountActivityGetTransactionResponsePaymentTransactionStatus = "CANCELED"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionStatus) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionStatusPending, AccountActivityGetTransactionResponsePaymentTransactionStatusSettled, AccountActivityGetTransactionResponsePaymentTransactionStatusDeclined, AccountActivityGetTransactionResponsePaymentTransactionStatusReversed, AccountActivityGetTransactionResponsePaymentTransactionStatusCanceled:
		return true
	}
	return false
}

type AccountActivityGetTransactionResponsePaymentTransactionType string

const (
	AccountActivityGetTransactionResponsePaymentTransactionTypeOriginationCredit   AccountActivityGetTransactionResponsePaymentTransactionType = "ORIGINATION_CREDIT"
	AccountActivityGetTransactionResponsePaymentTransactionTypeOriginationDebit    AccountActivityGetTransactionResponsePaymentTransactionType = "ORIGINATION_DEBIT"
	AccountActivityGetTransactionResponsePaymentTransactionTypeReceiptCredit       AccountActivityGetTransactionResponsePaymentTransactionType = "RECEIPT_CREDIT"
	AccountActivityGetTransactionResponsePaymentTransactionTypeReceiptDebit        AccountActivityGetTransactionResponsePaymentTransactionType = "RECEIPT_DEBIT"
	AccountActivityGetTransactionResponsePaymentTransactionTypeWireInboundPayment  AccountActivityGetTransactionResponsePaymentTransactionType = "WIRE_INBOUND_PAYMENT"
	AccountActivityGetTransactionResponsePaymentTransactionTypeWireInboundAdmin    AccountActivityGetTransactionResponsePaymentTransactionType = "WIRE_INBOUND_ADMIN"
	AccountActivityGetTransactionResponsePaymentTransactionTypeWireOutboundPayment AccountActivityGetTransactionResponsePaymentTransactionType = "WIRE_OUTBOUND_PAYMENT"
	AccountActivityGetTransactionResponsePaymentTransactionTypeWireOutboundAdmin   AccountActivityGetTransactionResponsePaymentTransactionType = "WIRE_OUTBOUND_ADMIN"
)

func (r AccountActivityGetTransactionResponsePaymentTransactionType) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponsePaymentTransactionTypeOriginationCredit, AccountActivityGetTransactionResponsePaymentTransactionTypeOriginationDebit, AccountActivityGetTransactionResponsePaymentTransactionTypeReceiptCredit, AccountActivityGetTransactionResponsePaymentTransactionTypeReceiptDebit, AccountActivityGetTransactionResponsePaymentTransactionTypeWireInboundPayment, AccountActivityGetTransactionResponsePaymentTransactionTypeWireInboundAdmin, AccountActivityGetTransactionResponsePaymentTransactionTypeWireOutboundPayment, AccountActivityGetTransactionResponsePaymentTransactionTypeWireOutboundAdmin:
		return true
	}
	return false
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
)

func (r AccountActivityGetTransactionResponseCategory) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseCategoryACH, AccountActivityGetTransactionResponseCategoryBalanceOrFunding, AccountActivityGetTransactionResponseCategoryCard, AccountActivityGetTransactionResponseCategoryExternalACH, AccountActivityGetTransactionResponseCategoryExternalCheck, AccountActivityGetTransactionResponseCategoryExternalTransfer, AccountActivityGetTransactionResponseCategoryExternalWire, AccountActivityGetTransactionResponseCategoryManagementAdjustment, AccountActivityGetTransactionResponseCategoryManagementDispute, AccountActivityGetTransactionResponseCategoryManagementFee, AccountActivityGetTransactionResponseCategoryManagementReward, AccountActivityGetTransactionResponseCategoryManagementDisbursement, AccountActivityGetTransactionResponseCategoryProgramFunding:
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

type AccountActivityGetTransactionResponseFamily string

const (
	AccountActivityGetTransactionResponseFamilyCard                AccountActivityGetTransactionResponseFamily = "CARD"
	AccountActivityGetTransactionResponseFamilyPayment             AccountActivityGetTransactionResponseFamily = "PAYMENT"
	AccountActivityGetTransactionResponseFamilyTransfer            AccountActivityGetTransactionResponseFamily = "TRANSFER"
	AccountActivityGetTransactionResponseFamilyInternal            AccountActivityGetTransactionResponseFamily = "INTERNAL"
	AccountActivityGetTransactionResponseFamilyExternalPayment     AccountActivityGetTransactionResponseFamily = "EXTERNAL_PAYMENT"
	AccountActivityGetTransactionResponseFamilyManagementOperation AccountActivityGetTransactionResponseFamily = "MANAGEMENT_OPERATION"
)

func (r AccountActivityGetTransactionResponseFamily) IsKnown() bool {
	switch r {
	case AccountActivityGetTransactionResponseFamilyCard, AccountActivityGetTransactionResponseFamilyPayment, AccountActivityGetTransactionResponseFamilyTransfer, AccountActivityGetTransactionResponseFamilyInternal, AccountActivityGetTransactionResponseFamilyExternalPayment, AccountActivityGetTransactionResponseFamilyManagementOperation:
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
	case AccountActivityGetTransactionResponseResultApproved, AccountActivityGetTransactionResponseResultDeclined, AccountActivityGetTransactionResponseResultAccountPaused, AccountActivityGetTransactionResponseResultAccountStateTransactionFail, AccountActivityGetTransactionResponseResultBankConnectionError, AccountActivityGetTransactionResponseResultBankNotVerified, AccountActivityGetTransactionResponseResultCardClosed, AccountActivityGetTransactionResponseResultCardPaused, AccountActivityGetTransactionResponseResultFraudAdvice, AccountActivityGetTransactionResponseResultIgnoredTtlExpiry, AccountActivityGetTransactionResponseResultInactiveAccount, AccountActivityGetTransactionResponseResultIncorrectPin, AccountActivityGetTransactionResponseResultInvalidCardDetails, AccountActivityGetTransactionResponseResultInsufficientFunds, AccountActivityGetTransactionResponseResultInsufficientFundsPreload, AccountActivityGetTransactionResponseResultInvalidTransaction, AccountActivityGetTransactionResponseResultMerchantBlacklist, AccountActivityGetTransactionResponseResultOriginalNotFound, AccountActivityGetTransactionResponseResultPreviouslyCompleted, AccountActivityGetTransactionResponseResultSingleUseRecharged, AccountActivityGetTransactionResponseResultSwitchInoperativeAdvice, AccountActivityGetTransactionResponseResultUnauthorizedMerchant, AccountActivityGetTransactionResponseResultUnknownHostTimeout, AccountActivityGetTransactionResponseResultUserTransactionLimit:
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
	case AccountActivityListParamsCategoryACH, AccountActivityListParamsCategoryBalanceOrFunding, AccountActivityListParamsCategoryCard, AccountActivityListParamsCategoryExternalACH, AccountActivityListParamsCategoryExternalCheck, AccountActivityListParamsCategoryExternalTransfer, AccountActivityListParamsCategoryExternalWire, AccountActivityListParamsCategoryManagementAdjustment, AccountActivityListParamsCategoryManagementDispute, AccountActivityListParamsCategoryManagementFee, AccountActivityListParamsCategoryManagementReward, AccountActivityListParamsCategoryManagementDisbursement, AccountActivityListParamsCategoryProgramFunding:
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
