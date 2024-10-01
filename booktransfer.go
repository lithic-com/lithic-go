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

// BookTransferService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBookTransferService] method instead.
type BookTransferService struct {
	Options []option.RequestOption
}

// NewBookTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBookTransferService(opts ...option.RequestOption) (r *BookTransferService) {
	r = &BookTransferService{}
	r.Options = opts
	return
}

// Book transfer funds between two financial accounts or between a financial
// account and card
func (r *BookTransferService) New(ctx context.Context, body BookTransferNewParams, opts ...option.RequestOption) (res *BookTransferResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/book_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get book transfer by token
func (r *BookTransferService) Get(ctx context.Context, bookTransferToken string, opts ...option.RequestOption) (res *BookTransferResponse, err error) {
	opts = append(r.Options[:], opts...)
	if bookTransferToken == "" {
		err = errors.New("missing required book_transfer_token parameter")
		return
	}
	path := fmt.Sprintf("v1/book_transfers/%s", bookTransferToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List book transfers
func (r *BookTransferService) List(ctx context.Context, query BookTransferListParams, opts ...option.RequestOption) (res *pagination.CursorPage[BookTransferResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/book_transfers"
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

// List book transfers
func (r *BookTransferService) ListAutoPaging(ctx context.Context, query BookTransferListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[BookTransferResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Reverse a book transfer
func (r *BookTransferService) Reverse(ctx context.Context, bookTransferToken string, body BookTransferReverseParams, opts ...option.RequestOption) (res *BookTransferResponse, err error) {
	opts = append(r.Options[:], opts...)
	if bookTransferToken == "" {
		err = errors.New("missing required book_transfer_token parameter")
		return
	}
	path := fmt.Sprintf("v1/book_transfers/%s/reverse", bookTransferToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BookTransferResponse struct {
	// Customer-provided token that will serve as an idempotency token. This token will
	// become the transaction token.
	Token string `json:"token,required" format:"uuid"`
	// Category of the book transfer
	Category BookTransferResponseCategory `json:"category,required"`
	// Date and time when the transfer occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the settling currency of the transaction.
	Currency string `json:"currency,required"`
	// A list of all financial events that have modified this transfer.
	Events []BookTransferResponseEvent `json:"events,required"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case.
	FromFinancialAccountToken string `json:"from_financial_account_token,required" format:"uuid"`
	// Pending amount of the transaction in the currency's smallest unit (e.g., cents),
	// including any acquirer fees. The value of this field will go to zero over time
	// once the financial transaction is settled.
	PendingAmount int64 `json:"pending_amount,required"`
	// APPROVED transactions were successful while DECLINED transactions were declined
	// by user, Lithic, or the network.
	Result BookTransferResponseResult `json:"result,required"`
	// Amount of the transaction that has been settled in the currency's smallest unit
	// (e.g., cents).
	SettledAmount int64 `json:"settled_amount,required"`
	// Status types: _ `DECLINED` - The transfer was declined. _ `REVERSED` - The
	// transfer was reversed \* `SETTLED` - The transfer is completed.
	Status BookTransferResponseStatus `json:"status,required"`
	// Globally unique identifier for the financial account or card that will receive
	// the funds. Accepted type dependent on the program's use case.
	ToFinancialAccountToken interface{} `json:"to_financial_account_token,required"`
	// Date and time when the financial transaction was last updated. UTC time zone.
	Updated time.Time                `json:"updated,required" format:"date-time"`
	JSON    bookTransferResponseJSON `json:"-"`
}

// bookTransferResponseJSON contains the JSON metadata for the struct
// [BookTransferResponse]
type bookTransferResponseJSON struct {
	Token                     apijson.Field
	Category                  apijson.Field
	Created                   apijson.Field
	Currency                  apijson.Field
	Events                    apijson.Field
	FromFinancialAccountToken apijson.Field
	PendingAmount             apijson.Field
	Result                    apijson.Field
	SettledAmount             apijson.Field
	Status                    apijson.Field
	ToFinancialAccountToken   apijson.Field
	Updated                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *BookTransferResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookTransferResponseJSON) RawJSON() string {
	return r.raw
}

// Category of the book transfer
type BookTransferResponseCategory string

const (
	BookTransferResponseCategoryAdjustment       BookTransferResponseCategory = "ADJUSTMENT"
	BookTransferResponseCategoryBalanceOrFunding BookTransferResponseCategory = "BALANCE_OR_FUNDING"
	BookTransferResponseCategoryDerecognition    BookTransferResponseCategory = "DERECOGNITION"
	BookTransferResponseCategoryDispute          BookTransferResponseCategory = "DISPUTE"
	BookTransferResponseCategoryFee              BookTransferResponseCategory = "FEE"
	BookTransferResponseCategoryReward           BookTransferResponseCategory = "REWARD"
	BookTransferResponseCategoryTransfer         BookTransferResponseCategory = "TRANSFER"
)

func (r BookTransferResponseCategory) IsKnown() bool {
	switch r {
	case BookTransferResponseCategoryAdjustment, BookTransferResponseCategoryBalanceOrFunding, BookTransferResponseCategoryDerecognition, BookTransferResponseCategoryDispute, BookTransferResponseCategoryFee, BookTransferResponseCategoryReward, BookTransferResponseCategoryTransfer:
		return true
	}
	return false
}

type BookTransferResponseEvent struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount,required"`
	// Date and time when the financial event occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Detailed Results
	DetailedResults []BookTransferResponseEventsDetailedResult `json:"detailed_results,required"`
	// Memo for the transfer.
	Memo string `json:"memo,required"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result BookTransferResponseEventsResult `json:"result,required"`
	// The program specific subtype code for the specified category/type.
	Subtype string `json:"subtype,required"`
	// Type of the book transfer
	Type string                        `json:"type,required"`
	JSON bookTransferResponseEventJSON `json:"-"`
}

// bookTransferResponseEventJSON contains the JSON metadata for the struct
// [BookTransferResponseEvent]
type bookTransferResponseEventJSON struct {
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

func (r *BookTransferResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookTransferResponseEventJSON) RawJSON() string {
	return r.raw
}

type BookTransferResponseEventsDetailedResult string

const (
	BookTransferResponseEventsDetailedResultApproved          BookTransferResponseEventsDetailedResult = "APPROVED"
	BookTransferResponseEventsDetailedResultFundsInsufficient BookTransferResponseEventsDetailedResult = "FUNDS_INSUFFICIENT"
)

func (r BookTransferResponseEventsDetailedResult) IsKnown() bool {
	switch r {
	case BookTransferResponseEventsDetailedResultApproved, BookTransferResponseEventsDetailedResultFundsInsufficient:
		return true
	}
	return false
}

// APPROVED financial events were successful while DECLINED financial events were
// declined by user, Lithic, or the network.
type BookTransferResponseEventsResult string

const (
	BookTransferResponseEventsResultApproved BookTransferResponseEventsResult = "APPROVED"
	BookTransferResponseEventsResultDeclined BookTransferResponseEventsResult = "DECLINED"
)

func (r BookTransferResponseEventsResult) IsKnown() bool {
	switch r {
	case BookTransferResponseEventsResultApproved, BookTransferResponseEventsResultDeclined:
		return true
	}
	return false
}

// APPROVED transactions were successful while DECLINED transactions were declined
// by user, Lithic, or the network.
type BookTransferResponseResult string

const (
	BookTransferResponseResultApproved BookTransferResponseResult = "APPROVED"
	BookTransferResponseResultDeclined BookTransferResponseResult = "DECLINED"
)

func (r BookTransferResponseResult) IsKnown() bool {
	switch r {
	case BookTransferResponseResultApproved, BookTransferResponseResultDeclined:
		return true
	}
	return false
}

// Status types: _ `DECLINED` - The transfer was declined. _ `REVERSED` - The
// transfer was reversed \* `SETTLED` - The transfer is completed.
type BookTransferResponseStatus string

const (
	BookTransferResponseStatusDeclined BookTransferResponseStatus = "DECLINED"
	BookTransferResponseStatusReversed BookTransferResponseStatus = "REVERSED"
	BookTransferResponseStatusSettled  BookTransferResponseStatus = "SETTLED"
)

func (r BookTransferResponseStatus) IsKnown() bool {
	switch r {
	case BookTransferResponseStatusDeclined, BookTransferResponseStatusReversed, BookTransferResponseStatusSettled:
		return true
	}
	return false
}

type BookTransferNewParams struct {
	// Amount to be transferred in the currency’s smallest unit (e.g., cents for USD).
	// This should always be a positive value.
	Amount param.Field[int64] `json:"amount,required"`
	// Category of the book transfer
	Category param.Field[BookTransferNewParamsCategory] `json:"category,required"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case.
	FromFinancialAccountToken param.Field[string] `json:"from_financial_account_token,required" format:"uuid"`
	// The program specific subtype code for the specified category/type.
	Subtype param.Field[string] `json:"subtype,required"`
	// Globally unique identifier for the financial account or card that will receive
	// the funds. Accepted type dependent on the program's use case.
	ToFinancialAccountToken param.Field[string] `json:"to_financial_account_token,required" format:"uuid"`
	// Type of book_transfer
	Type param.Field[BookTransferNewParamsType] `json:"type,required"`
	// Customer-provided token that will serve as an idempotency token. This token will
	// become the transaction token.
	Token param.Field[string] `json:"token" format:"uuid"`
	// Optional descriptor for the transfer.
	Memo param.Field[string] `json:"memo"`
}

func (r BookTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Category of the book transfer
type BookTransferNewParamsCategory string

const (
	BookTransferNewParamsCategoryAdjustment       BookTransferNewParamsCategory = "ADJUSTMENT"
	BookTransferNewParamsCategoryBalanceOrFunding BookTransferNewParamsCategory = "BALANCE_OR_FUNDING"
	BookTransferNewParamsCategoryDerecognition    BookTransferNewParamsCategory = "DERECOGNITION"
	BookTransferNewParamsCategoryDispute          BookTransferNewParamsCategory = "DISPUTE"
	BookTransferNewParamsCategoryFee              BookTransferNewParamsCategory = "FEE"
	BookTransferNewParamsCategoryReward           BookTransferNewParamsCategory = "REWARD"
	BookTransferNewParamsCategoryTransfer         BookTransferNewParamsCategory = "TRANSFER"
)

func (r BookTransferNewParamsCategory) IsKnown() bool {
	switch r {
	case BookTransferNewParamsCategoryAdjustment, BookTransferNewParamsCategoryBalanceOrFunding, BookTransferNewParamsCategoryDerecognition, BookTransferNewParamsCategoryDispute, BookTransferNewParamsCategoryFee, BookTransferNewParamsCategoryReward, BookTransferNewParamsCategoryTransfer:
		return true
	}
	return false
}

// Type of book_transfer
type BookTransferNewParamsType string

const (
	BookTransferNewParamsTypeAtmWithdrawal              BookTransferNewParamsType = "ATM_WITHDRAWAL"
	BookTransferNewParamsTypeAtmDecline                 BookTransferNewParamsType = "ATM_DECLINE"
	BookTransferNewParamsTypeInternationalAtmWithdrawal BookTransferNewParamsType = "INTERNATIONAL_ATM_WITHDRAWAL"
	BookTransferNewParamsTypeInactivity                 BookTransferNewParamsType = "INACTIVITY"
	BookTransferNewParamsTypeStatement                  BookTransferNewParamsType = "STATEMENT"
	BookTransferNewParamsTypeMonthly                    BookTransferNewParamsType = "MONTHLY"
	BookTransferNewParamsTypeQuarterly                  BookTransferNewParamsType = "QUARTERLY"
	BookTransferNewParamsTypeAnnual                     BookTransferNewParamsType = "ANNUAL"
	BookTransferNewParamsTypeCustomerService            BookTransferNewParamsType = "CUSTOMER_SERVICE"
	BookTransferNewParamsTypeAccountMaintenance         BookTransferNewParamsType = "ACCOUNT_MAINTENANCE"
	BookTransferNewParamsTypeAccountActivation          BookTransferNewParamsType = "ACCOUNT_ACTIVATION"
	BookTransferNewParamsTypeAccountClosure             BookTransferNewParamsType = "ACCOUNT_CLOSURE"
	BookTransferNewParamsTypeCardReplacement            BookTransferNewParamsType = "CARD_REPLACEMENT"
	BookTransferNewParamsTypeCardDelivery               BookTransferNewParamsType = "CARD_DELIVERY"
	BookTransferNewParamsTypeCardCreate                 BookTransferNewParamsType = "CARD_CREATE"
	BookTransferNewParamsTypeCurrencyConversion         BookTransferNewParamsType = "CURRENCY_CONVERSION"
	BookTransferNewParamsTypeInterest                   BookTransferNewParamsType = "INTEREST"
	BookTransferNewParamsTypeLatePayment                BookTransferNewParamsType = "LATE_PAYMENT"
	BookTransferNewParamsTypeBillPayment                BookTransferNewParamsType = "BILL_PAYMENT"
	BookTransferNewParamsTypeCashBack                   BookTransferNewParamsType = "CASH_BACK"
	BookTransferNewParamsTypeAccountToAccount           BookTransferNewParamsType = "ACCOUNT_TO_ACCOUNT"
	BookTransferNewParamsTypeCardToCard                 BookTransferNewParamsType = "CARD_TO_CARD"
	BookTransferNewParamsTypeDisburse                   BookTransferNewParamsType = "DISBURSE"
	BookTransferNewParamsTypeBillingError               BookTransferNewParamsType = "BILLING_ERROR"
	BookTransferNewParamsTypeLossWriteOff               BookTransferNewParamsType = "LOSS_WRITE_OFF"
	BookTransferNewParamsTypeExpiredCard                BookTransferNewParamsType = "EXPIRED_CARD"
	BookTransferNewParamsTypeEarlyDerecognition         BookTransferNewParamsType = "EARLY_DERECOGNITION"
	BookTransferNewParamsTypeEscheatment                BookTransferNewParamsType = "ESCHEATMENT"
	BookTransferNewParamsTypeInactivityFeeDown          BookTransferNewParamsType = "INACTIVITY_FEE_DOWN"
	BookTransferNewParamsTypeProvisionalCredit          BookTransferNewParamsType = "PROVISIONAL_CREDIT"
	BookTransferNewParamsTypeDisputeWon                 BookTransferNewParamsType = "DISPUTE_WON"
	BookTransferNewParamsTypeTransfer                   BookTransferNewParamsType = "TRANSFER"
)

func (r BookTransferNewParamsType) IsKnown() bool {
	switch r {
	case BookTransferNewParamsTypeAtmWithdrawal, BookTransferNewParamsTypeAtmDecline, BookTransferNewParamsTypeInternationalAtmWithdrawal, BookTransferNewParamsTypeInactivity, BookTransferNewParamsTypeStatement, BookTransferNewParamsTypeMonthly, BookTransferNewParamsTypeQuarterly, BookTransferNewParamsTypeAnnual, BookTransferNewParamsTypeCustomerService, BookTransferNewParamsTypeAccountMaintenance, BookTransferNewParamsTypeAccountActivation, BookTransferNewParamsTypeAccountClosure, BookTransferNewParamsTypeCardReplacement, BookTransferNewParamsTypeCardDelivery, BookTransferNewParamsTypeCardCreate, BookTransferNewParamsTypeCurrencyConversion, BookTransferNewParamsTypeInterest, BookTransferNewParamsTypeLatePayment, BookTransferNewParamsTypeBillPayment, BookTransferNewParamsTypeCashBack, BookTransferNewParamsTypeAccountToAccount, BookTransferNewParamsTypeCardToCard, BookTransferNewParamsTypeDisburse, BookTransferNewParamsTypeBillingError, BookTransferNewParamsTypeLossWriteOff, BookTransferNewParamsTypeExpiredCard, BookTransferNewParamsTypeEarlyDerecognition, BookTransferNewParamsTypeEscheatment, BookTransferNewParamsTypeInactivityFeeDown, BookTransferNewParamsTypeProvisionalCredit, BookTransferNewParamsTypeDisputeWon, BookTransferNewParamsTypeTransfer:
		return true
	}
	return false
}

type BookTransferListParams struct {
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin                param.Field[time.Time] `query:"begin" format:"date-time"`
	BusinessAccountToken param.Field[string]    `query:"business_account_token" format:"uuid"`
	// Book Transfer category to be returned.
	Category param.Field[BookTransferListParamsCategory] `query:"category"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case.
	FinancialAccountToken param.Field[string] `query:"financial_account_token" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// Book transfer result to be returned.
	Result param.Field[BookTransferListParamsResult] `query:"result"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Book transfer status to be returned.
	Status param.Field[BookTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [BookTransferListParams]'s query parameters as `url.Values`.
func (r BookTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Book Transfer category to be returned.
type BookTransferListParamsCategory string

const (
	BookTransferListParamsCategoryBalanceOrFunding BookTransferListParamsCategory = "BALANCE_OR_FUNDING"
	BookTransferListParamsCategoryFee              BookTransferListParamsCategory = "FEE"
	BookTransferListParamsCategoryReward           BookTransferListParamsCategory = "REWARD"
	BookTransferListParamsCategoryAdjustment       BookTransferListParamsCategory = "ADJUSTMENT"
	BookTransferListParamsCategoryDerecognition    BookTransferListParamsCategory = "DERECOGNITION"
	BookTransferListParamsCategoryDispute          BookTransferListParamsCategory = "DISPUTE"
	BookTransferListParamsCategoryInternal         BookTransferListParamsCategory = "INTERNAL"
)

func (r BookTransferListParamsCategory) IsKnown() bool {
	switch r {
	case BookTransferListParamsCategoryBalanceOrFunding, BookTransferListParamsCategoryFee, BookTransferListParamsCategoryReward, BookTransferListParamsCategoryAdjustment, BookTransferListParamsCategoryDerecognition, BookTransferListParamsCategoryDispute, BookTransferListParamsCategoryInternal:
		return true
	}
	return false
}

// Book transfer result to be returned.
type BookTransferListParamsResult string

const (
	BookTransferListParamsResultApproved BookTransferListParamsResult = "APPROVED"
	BookTransferListParamsResultDeclined BookTransferListParamsResult = "DECLINED"
)

func (r BookTransferListParamsResult) IsKnown() bool {
	switch r {
	case BookTransferListParamsResultApproved, BookTransferListParamsResultDeclined:
		return true
	}
	return false
}

// Book transfer status to be returned.
type BookTransferListParamsStatus string

const (
	BookTransferListParamsStatusDeclined BookTransferListParamsStatus = "DECLINED"
	BookTransferListParamsStatusSettled  BookTransferListParamsStatus = "SETTLED"
)

func (r BookTransferListParamsStatus) IsKnown() bool {
	switch r {
	case BookTransferListParamsStatusDeclined, BookTransferListParamsStatusSettled:
		return true
	}
	return false
}

type BookTransferReverseParams struct {
	// Optional descriptor for the reversal.
	Memo param.Field[string] `json:"memo"`
}

func (r BookTransferReverseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
