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
	opts = slices.Concat(r.Options, opts)
	path := "v1/book_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get book transfer by token
func (r *BookTransferService) Get(ctx context.Context, bookTransferToken string, opts ...option.RequestOption) (res *BookTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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

// Retry a book transfer that has been declined
func (r *BookTransferService) Retry(ctx context.Context, bookTransferToken string, body BookTransferRetryParams, opts ...option.RequestOption) (res *BookTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bookTransferToken == "" {
		err = errors.New("missing required book_transfer_token parameter")
		return
	}
	path := fmt.Sprintf("v1/book_transfers/%s/retry", bookTransferToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reverse a book transfer
func (r *BookTransferService) Reverse(ctx context.Context, bookTransferToken string, body BookTransferReverseParams, opts ...option.RequestOption) (res *BookTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bookTransferToken == "" {
		err = errors.New("missing required book_transfer_token parameter")
		return
	}
	path := fmt.Sprintf("v1/book_transfers/%s/reverse", bookTransferToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Book transfer transaction
type BookTransferResponse struct {
	// Unique identifier for the transaction
	Token    string                       `json:"token,required" format:"uuid"`
	Category BookTransferResponseCategory `json:"category,required"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-character alphabetic ISO 4217 code for the settling currency of the
	// transaction
	Currency string `json:"currency,required"`
	// A list of all financial events that have modified this transfer
	Events []BookTransferResponseEvent `json:"events,required"`
	// TRANSFER - Book Transfer Transaction
	Family BookTransferResponseFamily `json:"family,required"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case
	FromFinancialAccountToken string `json:"from_financial_account_token,required" format:"uuid"`
	// Pending amount of the transaction in the currency's smallest unit (e.g., cents),
	// including any acquirer fees.
	//
	// The value of this field will go to zero over time once the financial transaction
	// is settled.
	PendingAmount int64                      `json:"pending_amount,required"`
	Result        BookTransferResponseResult `json:"result,required"`
	// Amount of the transaction that has been settled in the currency's smallest unit
	// (e.g., cents)
	SettledAmount int64 `json:"settled_amount,required"`
	// The status of the transaction
	Status BookTransferResponseStatus `json:"status,required"`
	// Globally unique identifier for the financial account or card that will receive
	// the funds. Accepted type dependent on the program's use case
	ToFinancialAccountToken string `json:"to_financial_account_token,required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated time.Time `json:"updated,required" format:"date-time"`
	// External ID defined by the customer
	ExternalID string `json:"external_id,nullable"`
	// External resource associated with the management operation
	ExternalResource ExternalResource `json:"external_resource,nullable"`
	// A series of transactions that are grouped together
	TransactionSeries BookTransferResponseTransactionSeries `json:"transaction_series,nullable"`
	JSON              bookTransferResponseJSON              `json:"-"`
}

// bookTransferResponseJSON contains the JSON metadata for the struct
// [BookTransferResponse]
type bookTransferResponseJSON struct {
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

func (r *BookTransferResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookTransferResponseJSON) RawJSON() string {
	return r.raw
}

func (r BookTransferResponse) implementsAccountActivityListResponse() {}

func (r BookTransferResponse) implementsAccountActivityGetTransactionResponse() {}

type BookTransferResponseCategory string

const (
	BookTransferResponseCategoryAdjustment       BookTransferResponseCategory = "ADJUSTMENT"
	BookTransferResponseCategoryBalanceOrFunding BookTransferResponseCategory = "BALANCE_OR_FUNDING"
	BookTransferResponseCategoryDerecognition    BookTransferResponseCategory = "DERECOGNITION"
	BookTransferResponseCategoryDispute          BookTransferResponseCategory = "DISPUTE"
	BookTransferResponseCategoryFee              BookTransferResponseCategory = "FEE"
	BookTransferResponseCategoryInternal         BookTransferResponseCategory = "INTERNAL"
	BookTransferResponseCategoryReward           BookTransferResponseCategory = "REWARD"
	BookTransferResponseCategoryProgramFunding   BookTransferResponseCategory = "PROGRAM_FUNDING"
	BookTransferResponseCategoryTransfer         BookTransferResponseCategory = "TRANSFER"
)

func (r BookTransferResponseCategory) IsKnown() bool {
	switch r {
	case BookTransferResponseCategoryAdjustment, BookTransferResponseCategoryBalanceOrFunding, BookTransferResponseCategoryDerecognition, BookTransferResponseCategoryDispute, BookTransferResponseCategoryFee, BookTransferResponseCategoryInternal, BookTransferResponseCategoryReward, BookTransferResponseCategoryProgramFunding, BookTransferResponseCategoryTransfer:
		return true
	}
	return false
}

// Book transfer Event
type BookTransferResponseEvent struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount of the financial event that has been settled in the currency's smallest
	// unit (e.g., cents).
	Amount int64 `json:"amount,required"`
	// Date and time when the financial event occurred. UTC time zone.
	Created         time.Time                                  `json:"created,required" format:"date-time"`
	DetailedResults []BookTransferResponseEventsDetailedResult `json:"detailed_results,required"`
	// Memo for the transfer.
	Memo string `json:"memo,required"`
	// APPROVED financial events were successful while DECLINED financial events were
	// declined by user, Lithic, or the network.
	Result BookTransferResponseEventsResult `json:"result,required"`
	// The program specific subtype code for the specified category/type.
	Subtype string `json:"subtype,required"`
	// Type of the book transfer
	Type BookTransferResponseEventsType `json:"type,required"`
	JSON bookTransferResponseEventJSON  `json:"-"`
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

// Type of the book transfer
type BookTransferResponseEventsType string

const (
	BookTransferResponseEventsTypeAtmBalanceInquiry          BookTransferResponseEventsType = "ATM_BALANCE_INQUIRY"
	BookTransferResponseEventsTypeAtmWithdrawal              BookTransferResponseEventsType = "ATM_WITHDRAWAL"
	BookTransferResponseEventsTypeAtmDecline                 BookTransferResponseEventsType = "ATM_DECLINE"
	BookTransferResponseEventsTypeInternationalAtmWithdrawal BookTransferResponseEventsType = "INTERNATIONAL_ATM_WITHDRAWAL"
	BookTransferResponseEventsTypeInactivity                 BookTransferResponseEventsType = "INACTIVITY"
	BookTransferResponseEventsTypeStatement                  BookTransferResponseEventsType = "STATEMENT"
	BookTransferResponseEventsTypeMonthly                    BookTransferResponseEventsType = "MONTHLY"
	BookTransferResponseEventsTypeQuarterly                  BookTransferResponseEventsType = "QUARTERLY"
	BookTransferResponseEventsTypeAnnual                     BookTransferResponseEventsType = "ANNUAL"
	BookTransferResponseEventsTypeCustomerService            BookTransferResponseEventsType = "CUSTOMER_SERVICE"
	BookTransferResponseEventsTypeAccountMaintenance         BookTransferResponseEventsType = "ACCOUNT_MAINTENANCE"
	BookTransferResponseEventsTypeAccountActivation          BookTransferResponseEventsType = "ACCOUNT_ACTIVATION"
	BookTransferResponseEventsTypeAccountClosure             BookTransferResponseEventsType = "ACCOUNT_CLOSURE"
	BookTransferResponseEventsTypeCardReplacement            BookTransferResponseEventsType = "CARD_REPLACEMENT"
	BookTransferResponseEventsTypeCardDelivery               BookTransferResponseEventsType = "CARD_DELIVERY"
	BookTransferResponseEventsTypeCardCreate                 BookTransferResponseEventsType = "CARD_CREATE"
	BookTransferResponseEventsTypeCurrencyConversion         BookTransferResponseEventsType = "CURRENCY_CONVERSION"
	BookTransferResponseEventsTypeInterest                   BookTransferResponseEventsType = "INTEREST"
	BookTransferResponseEventsTypeLatePayment                BookTransferResponseEventsType = "LATE_PAYMENT"
	BookTransferResponseEventsTypeBillPayment                BookTransferResponseEventsType = "BILL_PAYMENT"
	BookTransferResponseEventsTypeCashBack                   BookTransferResponseEventsType = "CASH_BACK"
	BookTransferResponseEventsTypeAccountToAccount           BookTransferResponseEventsType = "ACCOUNT_TO_ACCOUNT"
	BookTransferResponseEventsTypeCardToCard                 BookTransferResponseEventsType = "CARD_TO_CARD"
	BookTransferResponseEventsTypeDisburse                   BookTransferResponseEventsType = "DISBURSE"
	BookTransferResponseEventsTypeBillingError               BookTransferResponseEventsType = "BILLING_ERROR"
	BookTransferResponseEventsTypeLossWriteOff               BookTransferResponseEventsType = "LOSS_WRITE_OFF"
	BookTransferResponseEventsTypeExpiredCard                BookTransferResponseEventsType = "EXPIRED_CARD"
	BookTransferResponseEventsTypeEarlyDerecognition         BookTransferResponseEventsType = "EARLY_DERECOGNITION"
	BookTransferResponseEventsTypeEscheatment                BookTransferResponseEventsType = "ESCHEATMENT"
	BookTransferResponseEventsTypeInactivityFeeDown          BookTransferResponseEventsType = "INACTIVITY_FEE_DOWN"
	BookTransferResponseEventsTypeProvisionalCredit          BookTransferResponseEventsType = "PROVISIONAL_CREDIT"
	BookTransferResponseEventsTypeDisputeWon                 BookTransferResponseEventsType = "DISPUTE_WON"
	BookTransferResponseEventsTypeService                    BookTransferResponseEventsType = "SERVICE"
	BookTransferResponseEventsTypeTransfer                   BookTransferResponseEventsType = "TRANSFER"
	BookTransferResponseEventsTypeCollection                 BookTransferResponseEventsType = "COLLECTION"
)

func (r BookTransferResponseEventsType) IsKnown() bool {
	switch r {
	case BookTransferResponseEventsTypeAtmBalanceInquiry, BookTransferResponseEventsTypeAtmWithdrawal, BookTransferResponseEventsTypeAtmDecline, BookTransferResponseEventsTypeInternationalAtmWithdrawal, BookTransferResponseEventsTypeInactivity, BookTransferResponseEventsTypeStatement, BookTransferResponseEventsTypeMonthly, BookTransferResponseEventsTypeQuarterly, BookTransferResponseEventsTypeAnnual, BookTransferResponseEventsTypeCustomerService, BookTransferResponseEventsTypeAccountMaintenance, BookTransferResponseEventsTypeAccountActivation, BookTransferResponseEventsTypeAccountClosure, BookTransferResponseEventsTypeCardReplacement, BookTransferResponseEventsTypeCardDelivery, BookTransferResponseEventsTypeCardCreate, BookTransferResponseEventsTypeCurrencyConversion, BookTransferResponseEventsTypeInterest, BookTransferResponseEventsTypeLatePayment, BookTransferResponseEventsTypeBillPayment, BookTransferResponseEventsTypeCashBack, BookTransferResponseEventsTypeAccountToAccount, BookTransferResponseEventsTypeCardToCard, BookTransferResponseEventsTypeDisburse, BookTransferResponseEventsTypeBillingError, BookTransferResponseEventsTypeLossWriteOff, BookTransferResponseEventsTypeExpiredCard, BookTransferResponseEventsTypeEarlyDerecognition, BookTransferResponseEventsTypeEscheatment, BookTransferResponseEventsTypeInactivityFeeDown, BookTransferResponseEventsTypeProvisionalCredit, BookTransferResponseEventsTypeDisputeWon, BookTransferResponseEventsTypeService, BookTransferResponseEventsTypeTransfer, BookTransferResponseEventsTypeCollection:
		return true
	}
	return false
}

// TRANSFER - Book Transfer Transaction
type BookTransferResponseFamily string

const (
	BookTransferResponseFamilyTransfer BookTransferResponseFamily = "TRANSFER"
)

func (r BookTransferResponseFamily) IsKnown() bool {
	switch r {
	case BookTransferResponseFamilyTransfer:
		return true
	}
	return false
}

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

// The status of the transaction
type BookTransferResponseStatus string

const (
	BookTransferResponseStatusPending  BookTransferResponseStatus = "PENDING"
	BookTransferResponseStatusSettled  BookTransferResponseStatus = "SETTLED"
	BookTransferResponseStatusDeclined BookTransferResponseStatus = "DECLINED"
	BookTransferResponseStatusReversed BookTransferResponseStatus = "REVERSED"
	BookTransferResponseStatusCanceled BookTransferResponseStatus = "CANCELED"
	BookTransferResponseStatusReturned BookTransferResponseStatus = "RETURNED"
)

func (r BookTransferResponseStatus) IsKnown() bool {
	switch r {
	case BookTransferResponseStatusPending, BookTransferResponseStatusSettled, BookTransferResponseStatusDeclined, BookTransferResponseStatusReversed, BookTransferResponseStatusCanceled, BookTransferResponseStatusReturned:
		return true
	}
	return false
}

// A series of transactions that are grouped together
type BookTransferResponseTransactionSeries struct {
	RelatedTransactionEventToken string                                    `json:"related_transaction_event_token,required,nullable" format:"uuid"`
	RelatedTransactionToken      string                                    `json:"related_transaction_token,required,nullable" format:"uuid"`
	Type                         string                                    `json:"type,required"`
	JSON                         bookTransferResponseTransactionSeriesJSON `json:"-"`
}

// bookTransferResponseTransactionSeriesJSON contains the JSON metadata for the
// struct [BookTransferResponseTransactionSeries]
type bookTransferResponseTransactionSeriesJSON struct {
	RelatedTransactionEventToken apijson.Field
	RelatedTransactionToken      apijson.Field
	Type                         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BookTransferResponseTransactionSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookTransferResponseTransactionSeriesJSON) RawJSON() string {
	return r.raw
}

type BookTransferNewParams struct {
	// Amount to be transferred in the currency's smallest unit (e.g., cents for USD).
	// This should always be a positive value.
	Amount   param.Field[int64]                         `json:"amount,required"`
	Category param.Field[BookTransferNewParamsCategory] `json:"category,required"`
	// Globally unique identifier for the financial account or card that will send the
	// funds. Accepted type dependent on the program's use case.
	FromFinancialAccountToken param.Field[string] `json:"from_financial_account_token,required" format:"uuid"`
	// The program specific subtype code for the specified category/type.
	Subtype param.Field[string] `json:"subtype,required"`
	// Globally unique identifier for the financial account or card that will receive
	// the funds. Accepted type dependent on the program's use case.
	ToFinancialAccountToken param.Field[string] `json:"to_financial_account_token,required" format:"uuid"`
	// Type of the book transfer
	Type param.Field[BookTransferNewParamsType] `json:"type,required"`
	// Customer-provided token that will serve as an idempotency token. This token will
	// become the transaction token.
	Token param.Field[string] `json:"token" format:"uuid"`
	// External ID defined by the customer
	ExternalID param.Field[string] `json:"external_id"`
	// Optional descriptor for the transfer.
	Memo param.Field[string] `json:"memo"`
	// What to do if the financial account is closed when posting an operation
	OnClosedAccount param.Field[BookTransferNewParamsOnClosedAccount] `json:"on_closed_account"`
}

func (r BookTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookTransferNewParamsCategory string

const (
	BookTransferNewParamsCategoryAdjustment       BookTransferNewParamsCategory = "ADJUSTMENT"
	BookTransferNewParamsCategoryBalanceOrFunding BookTransferNewParamsCategory = "BALANCE_OR_FUNDING"
	BookTransferNewParamsCategoryDerecognition    BookTransferNewParamsCategory = "DERECOGNITION"
	BookTransferNewParamsCategoryDispute          BookTransferNewParamsCategory = "DISPUTE"
	BookTransferNewParamsCategoryFee              BookTransferNewParamsCategory = "FEE"
	BookTransferNewParamsCategoryInternal         BookTransferNewParamsCategory = "INTERNAL"
	BookTransferNewParamsCategoryReward           BookTransferNewParamsCategory = "REWARD"
	BookTransferNewParamsCategoryProgramFunding   BookTransferNewParamsCategory = "PROGRAM_FUNDING"
	BookTransferNewParamsCategoryTransfer         BookTransferNewParamsCategory = "TRANSFER"
)

func (r BookTransferNewParamsCategory) IsKnown() bool {
	switch r {
	case BookTransferNewParamsCategoryAdjustment, BookTransferNewParamsCategoryBalanceOrFunding, BookTransferNewParamsCategoryDerecognition, BookTransferNewParamsCategoryDispute, BookTransferNewParamsCategoryFee, BookTransferNewParamsCategoryInternal, BookTransferNewParamsCategoryReward, BookTransferNewParamsCategoryProgramFunding, BookTransferNewParamsCategoryTransfer:
		return true
	}
	return false
}

// Type of the book transfer
type BookTransferNewParamsType string

const (
	BookTransferNewParamsTypeAtmBalanceInquiry          BookTransferNewParamsType = "ATM_BALANCE_INQUIRY"
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
	BookTransferNewParamsTypeService                    BookTransferNewParamsType = "SERVICE"
	BookTransferNewParamsTypeTransfer                   BookTransferNewParamsType = "TRANSFER"
	BookTransferNewParamsTypeCollection                 BookTransferNewParamsType = "COLLECTION"
)

func (r BookTransferNewParamsType) IsKnown() bool {
	switch r {
	case BookTransferNewParamsTypeAtmBalanceInquiry, BookTransferNewParamsTypeAtmWithdrawal, BookTransferNewParamsTypeAtmDecline, BookTransferNewParamsTypeInternationalAtmWithdrawal, BookTransferNewParamsTypeInactivity, BookTransferNewParamsTypeStatement, BookTransferNewParamsTypeMonthly, BookTransferNewParamsTypeQuarterly, BookTransferNewParamsTypeAnnual, BookTransferNewParamsTypeCustomerService, BookTransferNewParamsTypeAccountMaintenance, BookTransferNewParamsTypeAccountActivation, BookTransferNewParamsTypeAccountClosure, BookTransferNewParamsTypeCardReplacement, BookTransferNewParamsTypeCardDelivery, BookTransferNewParamsTypeCardCreate, BookTransferNewParamsTypeCurrencyConversion, BookTransferNewParamsTypeInterest, BookTransferNewParamsTypeLatePayment, BookTransferNewParamsTypeBillPayment, BookTransferNewParamsTypeCashBack, BookTransferNewParamsTypeAccountToAccount, BookTransferNewParamsTypeCardToCard, BookTransferNewParamsTypeDisburse, BookTransferNewParamsTypeBillingError, BookTransferNewParamsTypeLossWriteOff, BookTransferNewParamsTypeExpiredCard, BookTransferNewParamsTypeEarlyDerecognition, BookTransferNewParamsTypeEscheatment, BookTransferNewParamsTypeInactivityFeeDown, BookTransferNewParamsTypeProvisionalCredit, BookTransferNewParamsTypeDisputeWon, BookTransferNewParamsTypeService, BookTransferNewParamsTypeTransfer, BookTransferNewParamsTypeCollection:
		return true
	}
	return false
}

// What to do if the financial account is closed when posting an operation
type BookTransferNewParamsOnClosedAccount string

const (
	BookTransferNewParamsOnClosedAccountFail        BookTransferNewParamsOnClosedAccount = "FAIL"
	BookTransferNewParamsOnClosedAccountUseSuspense BookTransferNewParamsOnClosedAccount = "USE_SUSPENSE"
)

func (r BookTransferNewParamsOnClosedAccount) IsKnown() bool {
	switch r {
	case BookTransferNewParamsOnClosedAccountFail, BookTransferNewParamsOnClosedAccountUseSuspense:
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
	BookTransferListParamsCategoryAdjustment       BookTransferListParamsCategory = "ADJUSTMENT"
	BookTransferListParamsCategoryBalanceOrFunding BookTransferListParamsCategory = "BALANCE_OR_FUNDING"
	BookTransferListParamsCategoryDerecognition    BookTransferListParamsCategory = "DERECOGNITION"
	BookTransferListParamsCategoryDispute          BookTransferListParamsCategory = "DISPUTE"
	BookTransferListParamsCategoryFee              BookTransferListParamsCategory = "FEE"
	BookTransferListParamsCategoryInternal         BookTransferListParamsCategory = "INTERNAL"
	BookTransferListParamsCategoryReward           BookTransferListParamsCategory = "REWARD"
	BookTransferListParamsCategoryProgramFunding   BookTransferListParamsCategory = "PROGRAM_FUNDING"
	BookTransferListParamsCategoryTransfer         BookTransferListParamsCategory = "TRANSFER"
)

func (r BookTransferListParamsCategory) IsKnown() bool {
	switch r {
	case BookTransferListParamsCategoryAdjustment, BookTransferListParamsCategoryBalanceOrFunding, BookTransferListParamsCategoryDerecognition, BookTransferListParamsCategoryDispute, BookTransferListParamsCategoryFee, BookTransferListParamsCategoryInternal, BookTransferListParamsCategoryReward, BookTransferListParamsCategoryProgramFunding, BookTransferListParamsCategoryTransfer:
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

type BookTransferRetryParams struct {
	// Customer-provided token that will serve as an idempotency token. This token will
	// become the transaction token.
	RetryToken param.Field[string] `json:"retry_token,required" format:"uuid"`
}

func (r BookTransferRetryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookTransferReverseParams struct {
	// Optional descriptor for the reversal.
	Memo param.Field[string] `json:"memo"`
}

func (r BookTransferReverseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
