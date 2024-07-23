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
	path := fmt.Sprintf("financial_accounts/%s/statements/%s/line_items", financialAccountToken, statementToken)
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
	Token    string                         `json:"token,required"`
	Amount   int64                          `json:"amount,required"`
	Category StatementLineItemsDataCategory `json:"category,required"`
	// Timestamp of when the line item was generated
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction
	Currency string `json:"currency,required"`
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
	EventType StatementLineItemsDataEventType `json:"event_type,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Globally unique identifier for a financial transaction
	FinancialTransactionToken string `json:"financial_transaction_token,required" format:"uuid"`
	// Date that the transaction settled
	SettledDate time.Time `json:"settled_date,required" format:"date"`
	// Globally unique identifier for a card
	CardToken  string                     `json:"card_token" format:"uuid"`
	Descriptor string                     `json:"descriptor"`
	JSON       statementLineItemsDataJSON `json:"-"`
}

// statementLineItemsDataJSON contains the JSON metadata for the struct
// [StatementLineItemsData]
type statementLineItemsDataJSON struct {
	Token                     apijson.Field
	Amount                    apijson.Field
	Category                  apijson.Field
	Created                   apijson.Field
	Currency                  apijson.Field
	EventType                 apijson.Field
	FinancialAccountToken     apijson.Field
	FinancialTransactionToken apijson.Field
	SettledDate               apijson.Field
	CardToken                 apijson.Field
	Descriptor                apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *StatementLineItemsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementLineItemsDataJSON) RawJSON() string {
	return r.raw
}

type StatementLineItemsDataCategory string

const (
	StatementLineItemsDataCategoryACH      StatementLineItemsDataCategory = "ACH"
	StatementLineItemsDataCategoryCard     StatementLineItemsDataCategory = "CARD"
	StatementLineItemsDataCategoryTransfer StatementLineItemsDataCategory = "TRANSFER"
)

func (r StatementLineItemsDataCategory) IsKnown() bool {
	switch r {
	case StatementLineItemsDataCategoryACH, StatementLineItemsDataCategoryCard, StatementLineItemsDataCategoryTransfer:
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
type StatementLineItemsDataEventType string

const (
	StatementLineItemsDataEventTypeACHOriginationCancelled      StatementLineItemsDataEventType = "ACH_ORIGINATION_CANCELLED"
	StatementLineItemsDataEventTypeACHOriginationInitiated      StatementLineItemsDataEventType = "ACH_ORIGINATION_INITIATED"
	StatementLineItemsDataEventTypeACHOriginationProcessed      StatementLineItemsDataEventType = "ACH_ORIGINATION_PROCESSED"
	StatementLineItemsDataEventTypeACHOriginationSettled        StatementLineItemsDataEventType = "ACH_ORIGINATION_SETTLED"
	StatementLineItemsDataEventTypeACHOriginationReleased       StatementLineItemsDataEventType = "ACH_ORIGINATION_RELEASED"
	StatementLineItemsDataEventTypeACHOriginationReviewed       StatementLineItemsDataEventType = "ACH_ORIGINATION_REVIEWED"
	StatementLineItemsDataEventTypeACHReceiptProcessed          StatementLineItemsDataEventType = "ACH_RECEIPT_PROCESSED"
	StatementLineItemsDataEventTypeACHReceiptSettled            StatementLineItemsDataEventType = "ACH_RECEIPT_SETTLED"
	StatementLineItemsDataEventTypeACHReturnInitiated           StatementLineItemsDataEventType = "ACH_RETURN_INITIATED"
	StatementLineItemsDataEventTypeACHReturnProcessed           StatementLineItemsDataEventType = "ACH_RETURN_PROCESSED"
	StatementLineItemsDataEventTypeAuthorization                StatementLineItemsDataEventType = "AUTHORIZATION"
	StatementLineItemsDataEventTypeAuthorizationAdvice          StatementLineItemsDataEventType = "AUTHORIZATION_ADVICE"
	StatementLineItemsDataEventTypeAuthorizationExpiry          StatementLineItemsDataEventType = "AUTHORIZATION_EXPIRY"
	StatementLineItemsDataEventTypeAuthorizationReversal        StatementLineItemsDataEventType = "AUTHORIZATION_REVERSAL"
	StatementLineItemsDataEventTypeBalanceInquiry               StatementLineItemsDataEventType = "BALANCE_INQUIRY"
	StatementLineItemsDataEventTypeClearing                     StatementLineItemsDataEventType = "CLEARING"
	StatementLineItemsDataEventTypeCorrectionCredit             StatementLineItemsDataEventType = "CORRECTION_CREDIT"
	StatementLineItemsDataEventTypeCorrectionDebit              StatementLineItemsDataEventType = "CORRECTION_DEBIT"
	StatementLineItemsDataEventTypeCreditAuthorization          StatementLineItemsDataEventType = "CREDIT_AUTHORIZATION"
	StatementLineItemsDataEventTypeCreditAuthorizationAdvice    StatementLineItemsDataEventType = "CREDIT_AUTHORIZATION_ADVICE"
	StatementLineItemsDataEventTypeFinancialAuthorization       StatementLineItemsDataEventType = "FINANCIAL_AUTHORIZATION"
	StatementLineItemsDataEventTypeFinancialCreditAuthorization StatementLineItemsDataEventType = "FINANCIAL_CREDIT_AUTHORIZATION"
	StatementLineItemsDataEventTypeReturn                       StatementLineItemsDataEventType = "RETURN"
	StatementLineItemsDataEventTypeReturnReversal               StatementLineItemsDataEventType = "RETURN_REVERSAL"
	StatementLineItemsDataEventTypeTransfer                     StatementLineItemsDataEventType = "TRANSFER"
	StatementLineItemsDataEventTypeTransferInsufficientFunds    StatementLineItemsDataEventType = "TRANSFER_INSUFFICIENT_FUNDS"
)

func (r StatementLineItemsDataEventType) IsKnown() bool {
	switch r {
	case StatementLineItemsDataEventTypeACHOriginationCancelled, StatementLineItemsDataEventTypeACHOriginationInitiated, StatementLineItemsDataEventTypeACHOriginationProcessed, StatementLineItemsDataEventTypeACHOriginationSettled, StatementLineItemsDataEventTypeACHOriginationReleased, StatementLineItemsDataEventTypeACHOriginationReviewed, StatementLineItemsDataEventTypeACHReceiptProcessed, StatementLineItemsDataEventTypeACHReceiptSettled, StatementLineItemsDataEventTypeACHReturnInitiated, StatementLineItemsDataEventTypeACHReturnProcessed, StatementLineItemsDataEventTypeAuthorization, StatementLineItemsDataEventTypeAuthorizationAdvice, StatementLineItemsDataEventTypeAuthorizationExpiry, StatementLineItemsDataEventTypeAuthorizationReversal, StatementLineItemsDataEventTypeBalanceInquiry, StatementLineItemsDataEventTypeClearing, StatementLineItemsDataEventTypeCorrectionCredit, StatementLineItemsDataEventTypeCorrectionDebit, StatementLineItemsDataEventTypeCreditAuthorization, StatementLineItemsDataEventTypeCreditAuthorizationAdvice, StatementLineItemsDataEventTypeFinancialAuthorization, StatementLineItemsDataEventTypeFinancialCreditAuthorization, StatementLineItemsDataEventTypeReturn, StatementLineItemsDataEventTypeReturnReversal, StatementLineItemsDataEventTypeTransfer, StatementLineItemsDataEventTypeTransferInsufficientFunds:
		return true
	}
	return false
}

type FinancialAccountStatementLineItemListResponse struct {
	// Globally unique identifier for a Statement Line Item
	Token    string                                                `json:"token,required"`
	Amount   int64                                                 `json:"amount,required"`
	Category FinancialAccountStatementLineItemListResponseCategory `json:"category,required"`
	// Timestamp of when the line item was generated
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction
	Currency string `json:"currency,required"`
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
	EventType FinancialAccountStatementLineItemListResponseEventType `json:"event_type,required"`
	// Globally unique identifier for a financial account
	FinancialAccountToken string `json:"financial_account_token,required" format:"uuid"`
	// Globally unique identifier for a financial transaction
	FinancialTransactionToken string `json:"financial_transaction_token,required" format:"uuid"`
	// Date that the transaction settled
	SettledDate time.Time `json:"settled_date,required" format:"date"`
	// Globally unique identifier for a card
	CardToken  string                                            `json:"card_token" format:"uuid"`
	Descriptor string                                            `json:"descriptor"`
	JSON       financialAccountStatementLineItemListResponseJSON `json:"-"`
}

// financialAccountStatementLineItemListResponseJSON contains the JSON metadata for
// the struct [FinancialAccountStatementLineItemListResponse]
type financialAccountStatementLineItemListResponseJSON struct {
	Token                     apijson.Field
	Amount                    apijson.Field
	Category                  apijson.Field
	Created                   apijson.Field
	Currency                  apijson.Field
	EventType                 apijson.Field
	FinancialAccountToken     apijson.Field
	FinancialTransactionToken apijson.Field
	SettledDate               apijson.Field
	CardToken                 apijson.Field
	Descriptor                apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *FinancialAccountStatementLineItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountStatementLineItemListResponseJSON) RawJSON() string {
	return r.raw
}

type FinancialAccountStatementLineItemListResponseCategory string

const (
	FinancialAccountStatementLineItemListResponseCategoryACH      FinancialAccountStatementLineItemListResponseCategory = "ACH"
	FinancialAccountStatementLineItemListResponseCategoryCard     FinancialAccountStatementLineItemListResponseCategory = "CARD"
	FinancialAccountStatementLineItemListResponseCategoryTransfer FinancialAccountStatementLineItemListResponseCategory = "TRANSFER"
)

func (r FinancialAccountStatementLineItemListResponseCategory) IsKnown() bool {
	switch r {
	case FinancialAccountStatementLineItemListResponseCategoryACH, FinancialAccountStatementLineItemListResponseCategoryCard, FinancialAccountStatementLineItemListResponseCategoryTransfer:
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
type FinancialAccountStatementLineItemListResponseEventType string

const (
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationCancelled      FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_CANCELLED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationInitiated      FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_INITIATED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationProcessed      FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_PROCESSED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationSettled        FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_SETTLED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReleased       FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_RELEASED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReviewed       FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_REVIEWED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReceiptProcessed          FinancialAccountStatementLineItemListResponseEventType = "ACH_RECEIPT_PROCESSED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReceiptSettled            FinancialAccountStatementLineItemListResponseEventType = "ACH_RECEIPT_SETTLED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReturnInitiated           FinancialAccountStatementLineItemListResponseEventType = "ACH_RETURN_INITIATED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReturnProcessed           FinancialAccountStatementLineItemListResponseEventType = "ACH_RETURN_PROCESSED"
	FinancialAccountStatementLineItemListResponseEventTypeAuthorization                FinancialAccountStatementLineItemListResponseEventType = "AUTHORIZATION"
	FinancialAccountStatementLineItemListResponseEventTypeAuthorizationAdvice          FinancialAccountStatementLineItemListResponseEventType = "AUTHORIZATION_ADVICE"
	FinancialAccountStatementLineItemListResponseEventTypeAuthorizationExpiry          FinancialAccountStatementLineItemListResponseEventType = "AUTHORIZATION_EXPIRY"
	FinancialAccountStatementLineItemListResponseEventTypeAuthorizationReversal        FinancialAccountStatementLineItemListResponseEventType = "AUTHORIZATION_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeBalanceInquiry               FinancialAccountStatementLineItemListResponseEventType = "BALANCE_INQUIRY"
	FinancialAccountStatementLineItemListResponseEventTypeClearing                     FinancialAccountStatementLineItemListResponseEventType = "CLEARING"
	FinancialAccountStatementLineItemListResponseEventTypeCorrectionCredit             FinancialAccountStatementLineItemListResponseEventType = "CORRECTION_CREDIT"
	FinancialAccountStatementLineItemListResponseEventTypeCorrectionDebit              FinancialAccountStatementLineItemListResponseEventType = "CORRECTION_DEBIT"
	FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorization          FinancialAccountStatementLineItemListResponseEventType = "CREDIT_AUTHORIZATION"
	FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorizationAdvice    FinancialAccountStatementLineItemListResponseEventType = "CREDIT_AUTHORIZATION_ADVICE"
	FinancialAccountStatementLineItemListResponseEventTypeFinancialAuthorization       FinancialAccountStatementLineItemListResponseEventType = "FINANCIAL_AUTHORIZATION"
	FinancialAccountStatementLineItemListResponseEventTypeFinancialCreditAuthorization FinancialAccountStatementLineItemListResponseEventType = "FINANCIAL_CREDIT_AUTHORIZATION"
	FinancialAccountStatementLineItemListResponseEventTypeReturn                       FinancialAccountStatementLineItemListResponseEventType = "RETURN"
	FinancialAccountStatementLineItemListResponseEventTypeReturnReversal               FinancialAccountStatementLineItemListResponseEventType = "RETURN_REVERSAL"
	FinancialAccountStatementLineItemListResponseEventTypeTransfer                     FinancialAccountStatementLineItemListResponseEventType = "TRANSFER"
	FinancialAccountStatementLineItemListResponseEventTypeTransferInsufficientFunds    FinancialAccountStatementLineItemListResponseEventType = "TRANSFER_INSUFFICIENT_FUNDS"
)

func (r FinancialAccountStatementLineItemListResponseEventType) IsKnown() bool {
	switch r {
	case FinancialAccountStatementLineItemListResponseEventTypeACHOriginationCancelled, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationInitiated, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationProcessed, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationSettled, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReleased, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReviewed, FinancialAccountStatementLineItemListResponseEventTypeACHReceiptProcessed, FinancialAccountStatementLineItemListResponseEventTypeACHReceiptSettled, FinancialAccountStatementLineItemListResponseEventTypeACHReturnInitiated, FinancialAccountStatementLineItemListResponseEventTypeACHReturnProcessed, FinancialAccountStatementLineItemListResponseEventTypeAuthorization, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationAdvice, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationExpiry, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationReversal, FinancialAccountStatementLineItemListResponseEventTypeBalanceInquiry, FinancialAccountStatementLineItemListResponseEventTypeClearing, FinancialAccountStatementLineItemListResponseEventTypeCorrectionCredit, FinancialAccountStatementLineItemListResponseEventTypeCorrectionDebit, FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorization, FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorizationAdvice, FinancialAccountStatementLineItemListResponseEventTypeFinancialAuthorization, FinancialAccountStatementLineItemListResponseEventTypeFinancialCreditAuthorization, FinancialAccountStatementLineItemListResponseEventTypeReturn, FinancialAccountStatementLineItemListResponseEventTypeReturnReversal, FinancialAccountStatementLineItemListResponseEventTypeTransfer, FinancialAccountStatementLineItemListResponseEventTypeTransferInsufficientFunds:
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
