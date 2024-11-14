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
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// FinancialAccountStatementLineItemService contains methods and other services
// that help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountStatementLineItemService] method instead.
type FinancialAccountStatementLineItemService struct {
	Options []option.RequestOption
}

// NewFinancialAccountStatementLineItemService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFinancialAccountStatementLineItemService(opts ...option.RequestOption) (r *FinancialAccountStatementLineItemService) {
	r = &FinancialAccountStatementLineItemService{}
	r.Options = opts
	return
}

// List the line items for a given statement within a given financial account.
func (r *FinancialAccountStatementLineItemService) List(ctx context.Context, financialAccountToken string, statementToken string, query FinancialAccountStatementLineItemListParams, opts ...option.RequestOption) (res *pagination.CursorPage[FinancialAccountStatementLineItemListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	if statementToken == "" {
		err = errors.New("missing required statement_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/statements/%s/line_items", financialAccountToken, statementToken)
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

// List the line items for a given statement within a given financial account.
func (r *FinancialAccountStatementLineItemService) ListAutoPaging(ctx context.Context, financialAccountToken string, statementToken string, query FinancialAccountStatementLineItemListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[FinancialAccountStatementLineItemListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, financialAccountToken, statementToken, query, opts...))
}

type StatementLineItems struct {
	Data    []StatementLineItemsData `json:"data,required"`
	HasMore bool                     `json:"has_more,required"`
	JSON    statementLineItemsJSON   `json:"-"`
}

// statementLineItemsJSON contains the JSON metadata for the struct
// [StatementLineItems]
type statementLineItemsJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatementLineItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementLineItemsJSON) RawJSON() string {
	return r.raw
}

type StatementLineItemsData struct {
	// Globally unique identifier for a Statement Line Item
	Token string `json:"token,required"`
	// Transaction amount in cents
	Amount   int64                          `json:"amount,required"`
	Category StatementLineItemsDataCategory `json:"category,required"`
	// Timestamp of when the line item was generated
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction
	Currency string `json:"currency,required"`
	// Date that the transaction effected the account balance
	EffectiveDate time.Time                       `json:"effective_date,required" format:"date"`
	EventType     StatementLineItemsDataEventType `json:"event_type,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Globally unique identifier for a financial transaction event
	FinancialTransactionEventToken string `json:"financial_transaction_event_token,required" format:"uuid"`
	// Globally unique identifier for a financial transaction
	FinancialTransactionToken string `json:"financial_transaction_token,required" format:"uuid"`
	// Globally unique identifier for a card
	CardToken  string                     `json:"card_token" format:"uuid"`
	Descriptor string                     `json:"descriptor"`
	JSON       statementLineItemsDataJSON `json:"-"`
}

// statementLineItemsDataJSON contains the JSON metadata for the struct
// [StatementLineItemsData]
type statementLineItemsDataJSON struct {
	Token                          apijson.Field
	Amount                         apijson.Field
	Category                       apijson.Field
	Created                        apijson.Field
	Currency                       apijson.Field
	EffectiveDate                  apijson.Field
	EventType                      apijson.Field
	FinancialAccountToken          apijson.Field
	FinancialTransactionEventToken apijson.Field
	FinancialTransactionToken      apijson.Field
	CardToken                      apijson.Field
	Descriptor                     apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *StatementLineItemsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementLineItemsDataJSON) RawJSON() string {
	return r.raw
}

type StatementLineItemsDataCategory string

const (
	StatementLineItemsDataCategoryACH                  StatementLineItemsDataCategory = "ACH"
	StatementLineItemsDataCategoryBalanceOrFunding     StatementLineItemsDataCategory = "BALANCE_OR_FUNDING"
	StatementLineItemsDataCategoryCard                 StatementLineItemsDataCategory = "CARD"
	StatementLineItemsDataCategoryExternalACH          StatementLineItemsDataCategory = "EXTERNAL_ACH"
	StatementLineItemsDataCategoryExternalCheck        StatementLineItemsDataCategory = "EXTERNAL_CHECK"
	StatementLineItemsDataCategoryExternalTransfer     StatementLineItemsDataCategory = "EXTERNAL_TRANSFER"
	StatementLineItemsDataCategoryExternalWire         StatementLineItemsDataCategory = "EXTERNAL_WIRE"
	StatementLineItemsDataCategoryManagementAdjustment StatementLineItemsDataCategory = "MANAGEMENT_ADJUSTMENT"
	StatementLineItemsDataCategoryManagementDispute    StatementLineItemsDataCategory = "MANAGEMENT_DISPUTE"
	StatementLineItemsDataCategoryManagementFee        StatementLineItemsDataCategory = "MANAGEMENT_FEE"
	StatementLineItemsDataCategoryManagementReward     StatementLineItemsDataCategory = "MANAGEMENT_REWARD"
)

func (r StatementLineItemsDataCategory) IsKnown() bool {
	switch r {
	case StatementLineItemsDataCategoryACH, StatementLineItemsDataCategoryBalanceOrFunding, StatementLineItemsDataCategoryCard, StatementLineItemsDataCategoryExternalACH, StatementLineItemsDataCategoryExternalCheck, StatementLineItemsDataCategoryExternalTransfer, StatementLineItemsDataCategoryExternalWire, StatementLineItemsDataCategoryManagementAdjustment, StatementLineItemsDataCategoryManagementDispute, StatementLineItemsDataCategoryManagementFee, StatementLineItemsDataCategoryManagementReward:
		return true
	}
	return false
}

type StatementLineItemsDataEventType string

const (
	StatementLineItemsDataEventTypeACHOriginationCancelled      StatementLineItemsDataEventType = "ACH_ORIGINATION_CANCELLED"
	StatementLineItemsDataEventTypeACHOriginationInitiated      StatementLineItemsDataEventType = "ACH_ORIGINATION_INITIATED"
	StatementLineItemsDataEventTypeACHOriginationProcessed      StatementLineItemsDataEventType = "ACH_ORIGINATION_PROCESSED"
	StatementLineItemsDataEventTypeACHOriginationReleased       StatementLineItemsDataEventType = "ACH_ORIGINATION_RELEASED"
	StatementLineItemsDataEventTypeACHOriginationReviewed       StatementLineItemsDataEventType = "ACH_ORIGINATION_REVIEWED"
	StatementLineItemsDataEventTypeACHOriginationSettled        StatementLineItemsDataEventType = "ACH_ORIGINATION_SETTLED"
	StatementLineItemsDataEventTypeACHReceiptProcessed          StatementLineItemsDataEventType = "ACH_RECEIPT_PROCESSED"
	StatementLineItemsDataEventTypeACHReceiptSettled            StatementLineItemsDataEventType = "ACH_RECEIPT_SETTLED"
	StatementLineItemsDataEventTypeACHReturnInitiated           StatementLineItemsDataEventType = "ACH_RETURN_INITIATED"
	StatementLineItemsDataEventTypeACHReturnProcessed           StatementLineItemsDataEventType = "ACH_RETURN_PROCESSED"
	StatementLineItemsDataEventTypeAuthorization                StatementLineItemsDataEventType = "AUTHORIZATION"
	StatementLineItemsDataEventTypeAuthorizationAdvice          StatementLineItemsDataEventType = "AUTHORIZATION_ADVICE"
	StatementLineItemsDataEventTypeAuthorizationExpiry          StatementLineItemsDataEventType = "AUTHORIZATION_EXPIRY"
	StatementLineItemsDataEventTypeAuthorizationReversal        StatementLineItemsDataEventType = "AUTHORIZATION_REVERSAL"
	StatementLineItemsDataEventTypeBalanceInquiry               StatementLineItemsDataEventType = "BALANCE_INQUIRY"
	StatementLineItemsDataEventTypeBillingError                 StatementLineItemsDataEventType = "BILLING_ERROR"
	StatementLineItemsDataEventTypeBillingErrorReversal         StatementLineItemsDataEventType = "BILLING_ERROR_REVERSAL"
	StatementLineItemsDataEventTypeCardToCard                   StatementLineItemsDataEventType = "CARD_TO_CARD"
	StatementLineItemsDataEventTypeCashBack                     StatementLineItemsDataEventType = "CASH_BACK"
	StatementLineItemsDataEventTypeCashBackReversal             StatementLineItemsDataEventType = "CASH_BACK_REVERSAL"
	StatementLineItemsDataEventTypeClearing                     StatementLineItemsDataEventType = "CLEARING"
	StatementLineItemsDataEventTypeCorrectionCredit             StatementLineItemsDataEventType = "CORRECTION_CREDIT"
	StatementLineItemsDataEventTypeCorrectionDebit              StatementLineItemsDataEventType = "CORRECTION_DEBIT"
	StatementLineItemsDataEventTypeCreditAuthorization          StatementLineItemsDataEventType = "CREDIT_AUTHORIZATION"
	StatementLineItemsDataEventTypeCreditAuthorizationAdvice    StatementLineItemsDataEventType = "CREDIT_AUTHORIZATION_ADVICE"
	StatementLineItemsDataEventTypeCurrencyConversion           StatementLineItemsDataEventType = "CURRENCY_CONVERSION"
	StatementLineItemsDataEventTypeCurrencyConversionReversal   StatementLineItemsDataEventType = "CURRENCY_CONVERSION_REVERSAL"
	StatementLineItemsDataEventTypeDisputeWon                   StatementLineItemsDataEventType = "DISPUTE_WON"
	StatementLineItemsDataEventTypeExternalACHCanceled          StatementLineItemsDataEventType = "EXTERNAL_ACH_CANCELED"
	StatementLineItemsDataEventTypeExternalACHInitiated         StatementLineItemsDataEventType = "EXTERNAL_ACH_INITIATED"
	StatementLineItemsDataEventTypeExternalACHReleased          StatementLineItemsDataEventType = "EXTERNAL_ACH_RELEASED"
	StatementLineItemsDataEventTypeExternalACHReversed          StatementLineItemsDataEventType = "EXTERNAL_ACH_REVERSED"
	StatementLineItemsDataEventTypeExternalACHSettled           StatementLineItemsDataEventType = "EXTERNAL_ACH_SETTLED"
	StatementLineItemsDataEventTypeExternalCheckCanceled        StatementLineItemsDataEventType = "EXTERNAL_CHECK_CANCELED"
	StatementLineItemsDataEventTypeExternalCheckInitiated       StatementLineItemsDataEventType = "EXTERNAL_CHECK_INITIATED"
	StatementLineItemsDataEventTypeExternalCheckReleased        StatementLineItemsDataEventType = "EXTERNAL_CHECK_RELEASED"
	StatementLineItemsDataEventTypeExternalCheckReversed        StatementLineItemsDataEventType = "EXTERNAL_CHECK_REVERSED"
	StatementLineItemsDataEventTypeExternalCheckSettled         StatementLineItemsDataEventType = "EXTERNAL_CHECK_SETTLED"
	StatementLineItemsDataEventTypeExternalTransferCanceled     StatementLineItemsDataEventType = "EXTERNAL_TRANSFER_CANCELED"
	StatementLineItemsDataEventTypeExternalTransferInitiated    StatementLineItemsDataEventType = "EXTERNAL_TRANSFER_INITIATED"
	StatementLineItemsDataEventTypeExternalTransferReleased     StatementLineItemsDataEventType = "EXTERNAL_TRANSFER_RELEASED"
	StatementLineItemsDataEventTypeExternalTransferReversed     StatementLineItemsDataEventType = "EXTERNAL_TRANSFER_REVERSED"
	StatementLineItemsDataEventTypeExternalTransferSettled      StatementLineItemsDataEventType = "EXTERNAL_TRANSFER_SETTLED"
	StatementLineItemsDataEventTypeExternalWireCanceled         StatementLineItemsDataEventType = "EXTERNAL_WIRE_CANCELED"
	StatementLineItemsDataEventTypeExternalWireInitiated        StatementLineItemsDataEventType = "EXTERNAL_WIRE_INITIATED"
	StatementLineItemsDataEventTypeExternalWireReleased         StatementLineItemsDataEventType = "EXTERNAL_WIRE_RELEASED"
	StatementLineItemsDataEventTypeExternalWireReversed         StatementLineItemsDataEventType = "EXTERNAL_WIRE_REVERSED"
	StatementLineItemsDataEventTypeExternalWireSettled          StatementLineItemsDataEventType = "EXTERNAL_WIRE_SETTLED"
	StatementLineItemsDataEventTypeFinancialAuthorization       StatementLineItemsDataEventType = "FINANCIAL_AUTHORIZATION"
	StatementLineItemsDataEventTypeFinancialCreditAuthorization StatementLineItemsDataEventType = "FINANCIAL_CREDIT_AUTHORIZATION"
	StatementLineItemsDataEventTypeInterest                     StatementLineItemsDataEventType = "INTEREST"
	StatementLineItemsDataEventTypeInterestReversal             StatementLineItemsDataEventType = "INTEREST_REVERSAL"
	StatementLineItemsDataEventTypeLatePayment                  StatementLineItemsDataEventType = "LATE_PAYMENT"
	StatementLineItemsDataEventTypeLatePaymentReversal          StatementLineItemsDataEventType = "LATE_PAYMENT_REVERSAL"
	StatementLineItemsDataEventTypeProvisionalCredit            StatementLineItemsDataEventType = "PROVISIONAL_CREDIT"
	StatementLineItemsDataEventTypeProvisionalCreditReversal    StatementLineItemsDataEventType = "PROVISIONAL_CREDIT_REVERSAL"
	StatementLineItemsDataEventTypeReturn                       StatementLineItemsDataEventType = "RETURN"
	StatementLineItemsDataEventTypeReturnReversal               StatementLineItemsDataEventType = "RETURN_REVERSAL"
	StatementLineItemsDataEventTypeTransfer                     StatementLineItemsDataEventType = "TRANSFER"
	StatementLineItemsDataEventTypeTransferInsufficientFunds    StatementLineItemsDataEventType = "TRANSFER_INSUFFICIENT_FUNDS"
)

func (r StatementLineItemsDataEventType) IsKnown() bool {
	switch r {
	case StatementLineItemsDataEventTypeACHOriginationCancelled, StatementLineItemsDataEventTypeACHOriginationInitiated, StatementLineItemsDataEventTypeACHOriginationProcessed, StatementLineItemsDataEventTypeACHOriginationReleased, StatementLineItemsDataEventTypeACHOriginationReviewed, StatementLineItemsDataEventTypeACHOriginationSettled, StatementLineItemsDataEventTypeACHReceiptProcessed, StatementLineItemsDataEventTypeACHReceiptSettled, StatementLineItemsDataEventTypeACHReturnInitiated, StatementLineItemsDataEventTypeACHReturnProcessed, StatementLineItemsDataEventTypeAuthorization, StatementLineItemsDataEventTypeAuthorizationAdvice, StatementLineItemsDataEventTypeAuthorizationExpiry, StatementLineItemsDataEventTypeAuthorizationReversal, StatementLineItemsDataEventTypeBalanceInquiry, StatementLineItemsDataEventTypeBillingError, StatementLineItemsDataEventTypeBillingErrorReversal, StatementLineItemsDataEventTypeCardToCard, StatementLineItemsDataEventTypeCashBack, StatementLineItemsDataEventTypeCashBackReversal, StatementLineItemsDataEventTypeClearing, StatementLineItemsDataEventTypeCorrectionCredit, StatementLineItemsDataEventTypeCorrectionDebit, StatementLineItemsDataEventTypeCreditAuthorization, StatementLineItemsDataEventTypeCreditAuthorizationAdvice, StatementLineItemsDataEventTypeCurrencyConversion, StatementLineItemsDataEventTypeCurrencyConversionReversal, StatementLineItemsDataEventTypeDisputeWon, StatementLineItemsDataEventTypeExternalACHCanceled, StatementLineItemsDataEventTypeExternalACHInitiated, StatementLineItemsDataEventTypeExternalACHReleased, StatementLineItemsDataEventTypeExternalACHReversed, StatementLineItemsDataEventTypeExternalACHSettled, StatementLineItemsDataEventTypeExternalCheckCanceled, StatementLineItemsDataEventTypeExternalCheckInitiated, StatementLineItemsDataEventTypeExternalCheckReleased, StatementLineItemsDataEventTypeExternalCheckReversed, StatementLineItemsDataEventTypeExternalCheckSettled, StatementLineItemsDataEventTypeExternalTransferCanceled, StatementLineItemsDataEventTypeExternalTransferInitiated, StatementLineItemsDataEventTypeExternalTransferReleased, StatementLineItemsDataEventTypeExternalTransferReversed, StatementLineItemsDataEventTypeExternalTransferSettled, StatementLineItemsDataEventTypeExternalWireCanceled, StatementLineItemsDataEventTypeExternalWireInitiated, StatementLineItemsDataEventTypeExternalWireReleased, StatementLineItemsDataEventTypeExternalWireReversed, StatementLineItemsDataEventTypeExternalWireSettled, StatementLineItemsDataEventTypeFinancialAuthorization, StatementLineItemsDataEventTypeFinancialCreditAuthorization, StatementLineItemsDataEventTypeInterest, StatementLineItemsDataEventTypeInterestReversal, StatementLineItemsDataEventTypeLatePayment, StatementLineItemsDataEventTypeLatePaymentReversal, StatementLineItemsDataEventTypeProvisionalCredit, StatementLineItemsDataEventTypeProvisionalCreditReversal, StatementLineItemsDataEventTypeReturn, StatementLineItemsDataEventTypeReturnReversal, StatementLineItemsDataEventTypeTransfer, StatementLineItemsDataEventTypeTransferInsufficientFunds:
		return true
	}
	return false
}

type FinancialAccountStatementLineItemListResponse struct {
	// Globally unique identifier for a Statement Line Item
	Token string `json:"token,required"`
	// Transaction amount in cents
	Amount   int64                                                 `json:"amount,required"`
	Category FinancialAccountStatementLineItemListResponseCategory `json:"category,required"`
	// Timestamp of when the line item was generated
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction
	Currency string `json:"currency,required"`
	// Date that the transaction effected the account balance
	EffectiveDate time.Time                                              `json:"effective_date,required" format:"date"`
	EventType     FinancialAccountStatementLineItemListResponseEventType `json:"event_type,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Globally unique identifier for a financial transaction event
	FinancialTransactionEventToken string `json:"financial_transaction_event_token,required" format:"uuid"`
	// Globally unique identifier for a financial transaction
	FinancialTransactionToken string `json:"financial_transaction_token,required" format:"uuid"`
	// Globally unique identifier for a card
	CardToken  string                                            `json:"card_token" format:"uuid"`
	Descriptor string                                            `json:"descriptor"`
	JSON       financialAccountStatementLineItemListResponseJSON `json:"-"`
}

// financialAccountStatementLineItemListResponseJSON contains the JSON metadata for
// the struct [FinancialAccountStatementLineItemListResponse]
type financialAccountStatementLineItemListResponseJSON struct {
	Token                          apijson.Field
	Amount                         apijson.Field
	Category                       apijson.Field
	Created                        apijson.Field
	Currency                       apijson.Field
	EffectiveDate                  apijson.Field
	EventType                      apijson.Field
	FinancialAccountToken          apijson.Field
	FinancialTransactionEventToken apijson.Field
	FinancialTransactionToken      apijson.Field
	CardToken                      apijson.Field
	Descriptor                     apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *FinancialAccountStatementLineItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountStatementLineItemListResponseJSON) RawJSON() string {
	return r.raw
}

type FinancialAccountStatementLineItemListResponseCategory string

const (
	FinancialAccountStatementLineItemListResponseCategoryACH                  FinancialAccountStatementLineItemListResponseCategory = "ACH"
	FinancialAccountStatementLineItemListResponseCategoryBalanceOrFunding     FinancialAccountStatementLineItemListResponseCategory = "BALANCE_OR_FUNDING"
	FinancialAccountStatementLineItemListResponseCategoryCard                 FinancialAccountStatementLineItemListResponseCategory = "CARD"
	FinancialAccountStatementLineItemListResponseCategoryExternalACH          FinancialAccountStatementLineItemListResponseCategory = "EXTERNAL_ACH"
	FinancialAccountStatementLineItemListResponseCategoryExternalCheck        FinancialAccountStatementLineItemListResponseCategory = "EXTERNAL_CHECK"
	FinancialAccountStatementLineItemListResponseCategoryExternalTransfer     FinancialAccountStatementLineItemListResponseCategory = "EXTERNAL_TRANSFER"
	FinancialAccountStatementLineItemListResponseCategoryExternalWire         FinancialAccountStatementLineItemListResponseCategory = "EXTERNAL_WIRE"
	FinancialAccountStatementLineItemListResponseCategoryManagementAdjustment FinancialAccountStatementLineItemListResponseCategory = "MANAGEMENT_ADJUSTMENT"
	FinancialAccountStatementLineItemListResponseCategoryManagementDispute    FinancialAccountStatementLineItemListResponseCategory = "MANAGEMENT_DISPUTE"
	FinancialAccountStatementLineItemListResponseCategoryManagementFee        FinancialAccountStatementLineItemListResponseCategory = "MANAGEMENT_FEE"
	FinancialAccountStatementLineItemListResponseCategoryManagementReward     FinancialAccountStatementLineItemListResponseCategory = "MANAGEMENT_REWARD"
)

func (r FinancialAccountStatementLineItemListResponseCategory) IsKnown() bool {
	switch r {
	case FinancialAccountStatementLineItemListResponseCategoryACH, FinancialAccountStatementLineItemListResponseCategoryBalanceOrFunding, FinancialAccountStatementLineItemListResponseCategoryCard, FinancialAccountStatementLineItemListResponseCategoryExternalACH, FinancialAccountStatementLineItemListResponseCategoryExternalCheck, FinancialAccountStatementLineItemListResponseCategoryExternalTransfer, FinancialAccountStatementLineItemListResponseCategoryExternalWire, FinancialAccountStatementLineItemListResponseCategoryManagementAdjustment, FinancialAccountStatementLineItemListResponseCategoryManagementDispute, FinancialAccountStatementLineItemListResponseCategoryManagementFee, FinancialAccountStatementLineItemListResponseCategoryManagementReward:
		return true
	}
	return false
}

type FinancialAccountStatementLineItemListResponseEventType string

const (
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationCancelled      FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_CANCELLED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationInitiated      FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_INITIATED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationProcessed      FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_PROCESSED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReleased       FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_RELEASED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReviewed       FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_REVIEWED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationSettled        FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_SETTLED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReceiptProcessed          FinancialAccountStatementLineItemListResponseEventType = "ACH_RECEIPT_PROCESSED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReceiptSettled            FinancialAccountStatementLineItemListResponseEventType = "ACH_RECEIPT_SETTLED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReturnInitiated           FinancialAccountStatementLineItemListResponseEventType = "ACH_RETURN_INITIATED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReturnProcessed           FinancialAccountStatementLineItemListResponseEventType = "ACH_RETURN_PROCESSED"
	FinancialAccountStatementLineItemListResponseEventTypeAuthorization                FinancialAccountStatementLineItemListResponseEventType = "AUTHORIZATION"
	FinancialAccountStatementLineItemListResponseEventTypeAuthorizationAdvice          FinancialAccountStatementLineItemListResponseEventType = "AUTHORIZATION_ADVICE"
	FinancialAccountStatementLineItemListResponseEventTypeAuthorizationExpiry          FinancialAccountStatementLineItemListResponseEventType = "AUTHORIZATION_EXPIRY"
	FinancialAccountStatementLineItemListResponseEventTypeAuthorizationReversal        FinancialAccountStatementLineItemListResponseEventType = "AUTHORIZATION_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeBalanceInquiry               FinancialAccountStatementLineItemListResponseEventType = "BALANCE_INQUIRY"
	FinancialAccountStatementLineItemListResponseEventTypeBillingError                 FinancialAccountStatementLineItemListResponseEventType = "BILLING_ERROR"
	FinancialAccountStatementLineItemListResponseEventTypeBillingErrorReversal         FinancialAccountStatementLineItemListResponseEventType = "BILLING_ERROR_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeCardToCard                   FinancialAccountStatementLineItemListResponseEventType = "CARD_TO_CARD"
	FinancialAccountStatementLineItemListResponseEventTypeCashBack                     FinancialAccountStatementLineItemListResponseEventType = "CASH_BACK"
	FinancialAccountStatementLineItemListResponseEventTypeCashBackReversal             FinancialAccountStatementLineItemListResponseEventType = "CASH_BACK_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeClearing                     FinancialAccountStatementLineItemListResponseEventType = "CLEARING"
	FinancialAccountStatementLineItemListResponseEventTypeCorrectionCredit             FinancialAccountStatementLineItemListResponseEventType = "CORRECTION_CREDIT"
	FinancialAccountStatementLineItemListResponseEventTypeCorrectionDebit              FinancialAccountStatementLineItemListResponseEventType = "CORRECTION_DEBIT"
	FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorization          FinancialAccountStatementLineItemListResponseEventType = "CREDIT_AUTHORIZATION"
	FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorizationAdvice    FinancialAccountStatementLineItemListResponseEventType = "CREDIT_AUTHORIZATION_ADVICE"
	FinancialAccountStatementLineItemListResponseEventTypeCurrencyConversion           FinancialAccountStatementLineItemListResponseEventType = "CURRENCY_CONVERSION"
	FinancialAccountStatementLineItemListResponseEventTypeCurrencyConversionReversal   FinancialAccountStatementLineItemListResponseEventType = "CURRENCY_CONVERSION_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeDisputeWon                   FinancialAccountStatementLineItemListResponseEventType = "DISPUTE_WON"
	FinancialAccountStatementLineItemListResponseEventTypeExternalACHCanceled          FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_ACH_CANCELED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalACHInitiated         FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_ACH_INITIATED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalACHReleased          FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_ACH_RELEASED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalACHReversed          FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_ACH_REVERSED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalACHSettled           FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_ACH_SETTLED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalCheckCanceled        FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_CHECK_CANCELED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalCheckInitiated       FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_CHECK_INITIATED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalCheckReleased        FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_CHECK_RELEASED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalCheckReversed        FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_CHECK_REVERSED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalCheckSettled         FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_CHECK_SETTLED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalTransferCanceled     FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_TRANSFER_CANCELED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalTransferInitiated    FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_TRANSFER_INITIATED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalTransferReleased     FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_TRANSFER_RELEASED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalTransferReversed     FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_TRANSFER_REVERSED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalTransferSettled      FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_TRANSFER_SETTLED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalWireCanceled         FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_WIRE_CANCELED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalWireInitiated        FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_WIRE_INITIATED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalWireReleased         FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_WIRE_RELEASED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalWireReversed         FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_WIRE_REVERSED"
	FinancialAccountStatementLineItemListResponseEventTypeExternalWireSettled          FinancialAccountStatementLineItemListResponseEventType = "EXTERNAL_WIRE_SETTLED"
	FinancialAccountStatementLineItemListResponseEventTypeFinancialAuthorization       FinancialAccountStatementLineItemListResponseEventType = "FINANCIAL_AUTHORIZATION"
	FinancialAccountStatementLineItemListResponseEventTypeFinancialCreditAuthorization FinancialAccountStatementLineItemListResponseEventType = "FINANCIAL_CREDIT_AUTHORIZATION"
	FinancialAccountStatementLineItemListResponseEventTypeInterest                     FinancialAccountStatementLineItemListResponseEventType = "INTEREST"
	FinancialAccountStatementLineItemListResponseEventTypeInterestReversal             FinancialAccountStatementLineItemListResponseEventType = "INTEREST_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeLatePayment                  FinancialAccountStatementLineItemListResponseEventType = "LATE_PAYMENT"
	FinancialAccountStatementLineItemListResponseEventTypeLatePaymentReversal          FinancialAccountStatementLineItemListResponseEventType = "LATE_PAYMENT_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeProvisionalCredit            FinancialAccountStatementLineItemListResponseEventType = "PROVISIONAL_CREDIT"
	FinancialAccountStatementLineItemListResponseEventTypeProvisionalCreditReversal    FinancialAccountStatementLineItemListResponseEventType = "PROVISIONAL_CREDIT_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeReturn                       FinancialAccountStatementLineItemListResponseEventType = "RETURN"
	FinancialAccountStatementLineItemListResponseEventTypeReturnReversal               FinancialAccountStatementLineItemListResponseEventType = "RETURN_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeTransfer                     FinancialAccountStatementLineItemListResponseEventType = "TRANSFER"
	FinancialAccountStatementLineItemListResponseEventTypeTransferInsufficientFunds    FinancialAccountStatementLineItemListResponseEventType = "TRANSFER_INSUFFICIENT_FUNDS"
)

func (r FinancialAccountStatementLineItemListResponseEventType) IsKnown() bool {
	switch r {
	case FinancialAccountStatementLineItemListResponseEventTypeACHOriginationCancelled, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationInitiated, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationProcessed, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReleased, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReviewed, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationSettled, FinancialAccountStatementLineItemListResponseEventTypeACHReceiptProcessed, FinancialAccountStatementLineItemListResponseEventTypeACHReceiptSettled, FinancialAccountStatementLineItemListResponseEventTypeACHReturnInitiated, FinancialAccountStatementLineItemListResponseEventTypeACHReturnProcessed, FinancialAccountStatementLineItemListResponseEventTypeAuthorization, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationAdvice, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationExpiry, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationReversal, FinancialAccountStatementLineItemListResponseEventTypeBalanceInquiry, FinancialAccountStatementLineItemListResponseEventTypeBillingError, FinancialAccountStatementLineItemListResponseEventTypeBillingErrorReversal, FinancialAccountStatementLineItemListResponseEventTypeCardToCard, FinancialAccountStatementLineItemListResponseEventTypeCashBack, FinancialAccountStatementLineItemListResponseEventTypeCashBackReversal, FinancialAccountStatementLineItemListResponseEventTypeClearing, FinancialAccountStatementLineItemListResponseEventTypeCorrectionCredit, FinancialAccountStatementLineItemListResponseEventTypeCorrectionDebit, FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorization, FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorizationAdvice, FinancialAccountStatementLineItemListResponseEventTypeCurrencyConversion, FinancialAccountStatementLineItemListResponseEventTypeCurrencyConversionReversal, FinancialAccountStatementLineItemListResponseEventTypeDisputeWon, FinancialAccountStatementLineItemListResponseEventTypeExternalACHCanceled, FinancialAccountStatementLineItemListResponseEventTypeExternalACHInitiated, FinancialAccountStatementLineItemListResponseEventTypeExternalACHReleased, FinancialAccountStatementLineItemListResponseEventTypeExternalACHReversed, FinancialAccountStatementLineItemListResponseEventTypeExternalACHSettled, FinancialAccountStatementLineItemListResponseEventTypeExternalCheckCanceled, FinancialAccountStatementLineItemListResponseEventTypeExternalCheckInitiated, FinancialAccountStatementLineItemListResponseEventTypeExternalCheckReleased, FinancialAccountStatementLineItemListResponseEventTypeExternalCheckReversed, FinancialAccountStatementLineItemListResponseEventTypeExternalCheckSettled, FinancialAccountStatementLineItemListResponseEventTypeExternalTransferCanceled, FinancialAccountStatementLineItemListResponseEventTypeExternalTransferInitiated, FinancialAccountStatementLineItemListResponseEventTypeExternalTransferReleased, FinancialAccountStatementLineItemListResponseEventTypeExternalTransferReversed, FinancialAccountStatementLineItemListResponseEventTypeExternalTransferSettled, FinancialAccountStatementLineItemListResponseEventTypeExternalWireCanceled, FinancialAccountStatementLineItemListResponseEventTypeExternalWireInitiated, FinancialAccountStatementLineItemListResponseEventTypeExternalWireReleased, FinancialAccountStatementLineItemListResponseEventTypeExternalWireReversed, FinancialAccountStatementLineItemListResponseEventTypeExternalWireSettled, FinancialAccountStatementLineItemListResponseEventTypeFinancialAuthorization, FinancialAccountStatementLineItemListResponseEventTypeFinancialCreditAuthorization, FinancialAccountStatementLineItemListResponseEventTypeInterest, FinancialAccountStatementLineItemListResponseEventTypeInterestReversal, FinancialAccountStatementLineItemListResponseEventTypeLatePayment, FinancialAccountStatementLineItemListResponseEventTypeLatePaymentReversal, FinancialAccountStatementLineItemListResponseEventTypeProvisionalCredit, FinancialAccountStatementLineItemListResponseEventTypeProvisionalCreditReversal, FinancialAccountStatementLineItemListResponseEventTypeReturn, FinancialAccountStatementLineItemListResponseEventTypeReturnReversal, FinancialAccountStatementLineItemListResponseEventTypeTransfer, FinancialAccountStatementLineItemListResponseEventTypeTransferInsufficientFunds:
		return true
	}
	return false
}

type FinancialAccountStatementLineItemListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [FinancialAccountStatementLineItemListParams]'s query
// parameters as `url.Values`.
func (r FinancialAccountStatementLineItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
