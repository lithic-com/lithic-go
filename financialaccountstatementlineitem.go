// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// FinancialAccountStatementLineItemService contains methods and other services
// that help with interacting with the lithic API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewFinancialAccountStatementLineItemService] method instead.
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
func (r *FinancialAccountStatementLineItemService) List(ctx context.Context, financialAccountToken string, statementToken string, query FinancialAccountStatementLineItemListParams, opts ...option.RequestOption) (res *shared.CursorPage[FinancialAccountStatementLineItemListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *FinancialAccountStatementLineItemService) ListAutoPaging(ctx context.Context, financialAccountToken string, statementToken string, query FinancialAccountStatementLineItemListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[FinancialAccountStatementLineItemListResponse] {
	return shared.NewCursorPageAutoPager(r.List(ctx, financialAccountToken, statementToken, query, opts...))
}

type FinancialAccountStatementLineItemListResponse struct {
	// Globally unique identifier for a Statement Line Item
	Token    string                                                `json:"token,required" format:"uuid"`
	Amount   int64                                                 `json:"amount,required"`
	Category FinancialAccountStatementLineItemListResponseCategory `json:"category,required"`
	// Timestamp of when the line item was generated
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction
	Currency string `json:"currency,required"`
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
type FinancialAccountStatementLineItemListResponseEventType string

const (
	FinancialAccountStatementLineItemListResponseEventTypeACHExceededThreshold         FinancialAccountStatementLineItemListResponseEventType = "ACH_EXCEEDED_THRESHOLD"
	FinancialAccountStatementLineItemListResponseEventTypeACHInsufficientFunds         FinancialAccountStatementLineItemListResponseEventType = "ACH_INSUFFICIENT_FUNDS"
	FinancialAccountStatementLineItemListResponseEventTypeACHInvalidAccount            FinancialAccountStatementLineItemListResponseEventType = "ACH_INVALID_ACCOUNT"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationPending        FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_PENDING"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationProcessed      FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_PROCESSED"
	FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReleased       FinancialAccountStatementLineItemListResponseEventType = "ACH_ORIGINATION_RELEASED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReceiptPending            FinancialAccountStatementLineItemListResponseEventType = "ACH_RECEIPT_PENDING"
	FinancialAccountStatementLineItemListResponseEventTypeACHReceiptReleased           FinancialAccountStatementLineItemListResponseEventType = "ACH_RECEIPT_RELEASED"
	FinancialAccountStatementLineItemListResponseEventTypeACHReturn                    FinancialAccountStatementLineItemListResponseEventType = "ACH_RETURN"
	FinancialAccountStatementLineItemListResponseEventTypeACHReturnPending             FinancialAccountStatementLineItemListResponseEventType = "ACH_RETURN_PENDING"
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
	case FinancialAccountStatementLineItemListResponseEventTypeACHExceededThreshold, FinancialAccountStatementLineItemListResponseEventTypeACHInsufficientFunds, FinancialAccountStatementLineItemListResponseEventTypeACHInvalidAccount, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationPending, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationProcessed, FinancialAccountStatementLineItemListResponseEventTypeACHOriginationReleased, FinancialAccountStatementLineItemListResponseEventTypeACHReceiptPending, FinancialAccountStatementLineItemListResponseEventTypeACHReceiptReleased, FinancialAccountStatementLineItemListResponseEventTypeACHReturn, FinancialAccountStatementLineItemListResponseEventTypeACHReturnPending, FinancialAccountStatementLineItemListResponseEventTypeAuthorization, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationAdvice, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationExpiry, FinancialAccountStatementLineItemListResponseEventTypeAuthorizationReversal, FinancialAccountStatementLineItemListResponseEventTypeBalanceInquiry, FinancialAccountStatementLineItemListResponseEventTypeClearing, FinancialAccountStatementLineItemListResponseEventTypeCorrectionCredit, FinancialAccountStatementLineItemListResponseEventTypeCorrectionDebit, FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorization, FinancialAccountStatementLineItemListResponseEventTypeCreditAuthorizationAdvice, FinancialAccountStatementLineItemListResponseEventTypeFinancialAuthorization, FinancialAccountStatementLineItemListResponseEventTypeFinancialCreditAuthorization, FinancialAccountStatementLineItemListResponseEventTypeReturn, FinancialAccountStatementLineItemListResponseEventTypeReturnReversal, FinancialAccountStatementLineItemListResponseEventTypeTransfer, FinancialAccountStatementLineItemListResponseEventTypeTransferInsufficientFunds:
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
