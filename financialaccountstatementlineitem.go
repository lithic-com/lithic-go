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
func (r *FinancialAccountStatementLineItemService) List(ctx context.Context, financialAccountToken string, statementToken string, query FinancialAccountStatementLineItemListParams, opts ...option.RequestOption) (res *pagination.CursorPage[StatementLineItemsData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *FinancialAccountStatementLineItemService) ListAutoPaging(ctx context.Context, financialAccountToken string, statementToken string, query FinancialAccountStatementLineItemListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[StatementLineItemsData] {
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
	// 3-character alphabetic ISO 4217 code for the settling currency of the
	// transaction
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
	StatementLineItemsDataCategoryACH                    StatementLineItemsDataCategory = "ACH"
	StatementLineItemsDataCategoryBalanceOrFunding       StatementLineItemsDataCategory = "BALANCE_OR_FUNDING"
	StatementLineItemsDataCategoryFee                    StatementLineItemsDataCategory = "FEE"
	StatementLineItemsDataCategoryReward                 StatementLineItemsDataCategory = "REWARD"
	StatementLineItemsDataCategoryAdjustment             StatementLineItemsDataCategory = "ADJUSTMENT"
	StatementLineItemsDataCategoryDerecognition          StatementLineItemsDataCategory = "DERECOGNITION"
	StatementLineItemsDataCategoryDispute                StatementLineItemsDataCategory = "DISPUTE"
	StatementLineItemsDataCategoryCard                   StatementLineItemsDataCategory = "CARD"
	StatementLineItemsDataCategoryExternalACH            StatementLineItemsDataCategory = "EXTERNAL_ACH"
	StatementLineItemsDataCategoryExternalCheck          StatementLineItemsDataCategory = "EXTERNAL_CHECK"
	StatementLineItemsDataCategoryExternalTransfer       StatementLineItemsDataCategory = "EXTERNAL_TRANSFER"
	StatementLineItemsDataCategoryExternalWire           StatementLineItemsDataCategory = "EXTERNAL_WIRE"
	StatementLineItemsDataCategoryManagementAdjustment   StatementLineItemsDataCategory = "MANAGEMENT_ADJUSTMENT"
	StatementLineItemsDataCategoryManagementDispute      StatementLineItemsDataCategory = "MANAGEMENT_DISPUTE"
	StatementLineItemsDataCategoryManagementFee          StatementLineItemsDataCategory = "MANAGEMENT_FEE"
	StatementLineItemsDataCategoryManagementReward       StatementLineItemsDataCategory = "MANAGEMENT_REWARD"
	StatementLineItemsDataCategoryManagementDisbursement StatementLineItemsDataCategory = "MANAGEMENT_DISBURSEMENT"
	StatementLineItemsDataCategoryProgramFunding         StatementLineItemsDataCategory = "PROGRAM_FUNDING"
)

func (r StatementLineItemsDataCategory) IsKnown() bool {
	switch r {
	case StatementLineItemsDataCategoryACH, StatementLineItemsDataCategoryBalanceOrFunding, StatementLineItemsDataCategoryFee, StatementLineItemsDataCategoryReward, StatementLineItemsDataCategoryAdjustment, StatementLineItemsDataCategoryDerecognition, StatementLineItemsDataCategoryDispute, StatementLineItemsDataCategoryCard, StatementLineItemsDataCategoryExternalACH, StatementLineItemsDataCategoryExternalCheck, StatementLineItemsDataCategoryExternalTransfer, StatementLineItemsDataCategoryExternalWire, StatementLineItemsDataCategoryManagementAdjustment, StatementLineItemsDataCategoryManagementDispute, StatementLineItemsDataCategoryManagementFee, StatementLineItemsDataCategoryManagementReward, StatementLineItemsDataCategoryManagementDisbursement, StatementLineItemsDataCategoryProgramFunding:
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
	StatementLineItemsDataEventTypeACHOriginationRejected       StatementLineItemsDataEventType = "ACH_ORIGINATION_REJECTED"
	StatementLineItemsDataEventTypeACHOriginationReviewed       StatementLineItemsDataEventType = "ACH_ORIGINATION_REVIEWED"
	StatementLineItemsDataEventTypeACHOriginationSettled        StatementLineItemsDataEventType = "ACH_ORIGINATION_SETTLED"
	StatementLineItemsDataEventTypeACHReceiptProcessed          StatementLineItemsDataEventType = "ACH_RECEIPT_PROCESSED"
	StatementLineItemsDataEventTypeACHReceiptSettled            StatementLineItemsDataEventType = "ACH_RECEIPT_SETTLED"
	StatementLineItemsDataEventTypeACHReturnInitiated           StatementLineItemsDataEventType = "ACH_RETURN_INITIATED"
	StatementLineItemsDataEventTypeACHReturnProcessed           StatementLineItemsDataEventType = "ACH_RETURN_PROCESSED"
	StatementLineItemsDataEventTypeACHReturnRejected            StatementLineItemsDataEventType = "ACH_RETURN_REJECTED"
	StatementLineItemsDataEventTypeACHReturnSettled             StatementLineItemsDataEventType = "ACH_RETURN_SETTLED"
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
	StatementLineItemsDataEventTypeCollection                   StatementLineItemsDataEventType = "COLLECTION"
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
	StatementLineItemsDataEventTypeInternalAdjustment           StatementLineItemsDataEventType = "INTERNAL_ADJUSTMENT"
	StatementLineItemsDataEventTypeLatePayment                  StatementLineItemsDataEventType = "LATE_PAYMENT"
	StatementLineItemsDataEventTypeLatePaymentReversal          StatementLineItemsDataEventType = "LATE_PAYMENT_REVERSAL"
	StatementLineItemsDataEventTypeLossWriteOff                 StatementLineItemsDataEventType = "LOSS_WRITE_OFF"
	StatementLineItemsDataEventTypeProvisionalCredit            StatementLineItemsDataEventType = "PROVISIONAL_CREDIT"
	StatementLineItemsDataEventTypeProvisionalCreditReversal    StatementLineItemsDataEventType = "PROVISIONAL_CREDIT_REVERSAL"
	StatementLineItemsDataEventTypeService                      StatementLineItemsDataEventType = "SERVICE"
	StatementLineItemsDataEventTypeReturn                       StatementLineItemsDataEventType = "RETURN"
	StatementLineItemsDataEventTypeReturnReversal               StatementLineItemsDataEventType = "RETURN_REVERSAL"
	StatementLineItemsDataEventTypeTransfer                     StatementLineItemsDataEventType = "TRANSFER"
	StatementLineItemsDataEventTypeTransferInsufficientFunds    StatementLineItemsDataEventType = "TRANSFER_INSUFFICIENT_FUNDS"
	StatementLineItemsDataEventTypeReturnedPayment              StatementLineItemsDataEventType = "RETURNED_PAYMENT"
	StatementLineItemsDataEventTypeReturnedPaymentReversal      StatementLineItemsDataEventType = "RETURNED_PAYMENT_REVERSAL"
	StatementLineItemsDataEventTypeLithicNetworkPayment         StatementLineItemsDataEventType = "LITHIC_NETWORK_PAYMENT"
)

func (r StatementLineItemsDataEventType) IsKnown() bool {
	switch r {
	case StatementLineItemsDataEventTypeACHOriginationCancelled, StatementLineItemsDataEventTypeACHOriginationInitiated, StatementLineItemsDataEventTypeACHOriginationProcessed, StatementLineItemsDataEventTypeACHOriginationReleased, StatementLineItemsDataEventTypeACHOriginationRejected, StatementLineItemsDataEventTypeACHOriginationReviewed, StatementLineItemsDataEventTypeACHOriginationSettled, StatementLineItemsDataEventTypeACHReceiptProcessed, StatementLineItemsDataEventTypeACHReceiptSettled, StatementLineItemsDataEventTypeACHReturnInitiated, StatementLineItemsDataEventTypeACHReturnProcessed, StatementLineItemsDataEventTypeACHReturnRejected, StatementLineItemsDataEventTypeACHReturnSettled, StatementLineItemsDataEventTypeAuthorization, StatementLineItemsDataEventTypeAuthorizationAdvice, StatementLineItemsDataEventTypeAuthorizationExpiry, StatementLineItemsDataEventTypeAuthorizationReversal, StatementLineItemsDataEventTypeBalanceInquiry, StatementLineItemsDataEventTypeBillingError, StatementLineItemsDataEventTypeBillingErrorReversal, StatementLineItemsDataEventTypeCardToCard, StatementLineItemsDataEventTypeCashBack, StatementLineItemsDataEventTypeCashBackReversal, StatementLineItemsDataEventTypeClearing, StatementLineItemsDataEventTypeCollection, StatementLineItemsDataEventTypeCorrectionCredit, StatementLineItemsDataEventTypeCorrectionDebit, StatementLineItemsDataEventTypeCreditAuthorization, StatementLineItemsDataEventTypeCreditAuthorizationAdvice, StatementLineItemsDataEventTypeCurrencyConversion, StatementLineItemsDataEventTypeCurrencyConversionReversal, StatementLineItemsDataEventTypeDisputeWon, StatementLineItemsDataEventTypeExternalACHCanceled, StatementLineItemsDataEventTypeExternalACHInitiated, StatementLineItemsDataEventTypeExternalACHReleased, StatementLineItemsDataEventTypeExternalACHReversed, StatementLineItemsDataEventTypeExternalACHSettled, StatementLineItemsDataEventTypeExternalCheckCanceled, StatementLineItemsDataEventTypeExternalCheckInitiated, StatementLineItemsDataEventTypeExternalCheckReleased, StatementLineItemsDataEventTypeExternalCheckReversed, StatementLineItemsDataEventTypeExternalCheckSettled, StatementLineItemsDataEventTypeExternalTransferCanceled, StatementLineItemsDataEventTypeExternalTransferInitiated, StatementLineItemsDataEventTypeExternalTransferReleased, StatementLineItemsDataEventTypeExternalTransferReversed, StatementLineItemsDataEventTypeExternalTransferSettled, StatementLineItemsDataEventTypeExternalWireCanceled, StatementLineItemsDataEventTypeExternalWireInitiated, StatementLineItemsDataEventTypeExternalWireReleased, StatementLineItemsDataEventTypeExternalWireReversed, StatementLineItemsDataEventTypeExternalWireSettled, StatementLineItemsDataEventTypeFinancialAuthorization, StatementLineItemsDataEventTypeFinancialCreditAuthorization, StatementLineItemsDataEventTypeInterest, StatementLineItemsDataEventTypeInterestReversal, StatementLineItemsDataEventTypeInternalAdjustment, StatementLineItemsDataEventTypeLatePayment, StatementLineItemsDataEventTypeLatePaymentReversal, StatementLineItemsDataEventTypeLossWriteOff, StatementLineItemsDataEventTypeProvisionalCredit, StatementLineItemsDataEventTypeProvisionalCreditReversal, StatementLineItemsDataEventTypeService, StatementLineItemsDataEventTypeReturn, StatementLineItemsDataEventTypeReturnReversal, StatementLineItemsDataEventTypeTransfer, StatementLineItemsDataEventTypeTransferInsufficientFunds, StatementLineItemsDataEventTypeReturnedPayment, StatementLineItemsDataEventTypeReturnedPaymentReversal, StatementLineItemsDataEventTypeLithicNetworkPayment:
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
