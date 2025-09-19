// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// FraudTransactionService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFraudTransactionService] method instead.
type FraudTransactionService struct {
	Options []option.RequestOption
}

// NewFraudTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFraudTransactionService(opts ...option.RequestOption) (r *FraudTransactionService) {
	r = &FraudTransactionService{}
	r.Options = opts
	return
}

// Retrieve a fraud report for a specific transaction identified by its unique
// transaction token.
func (r *FraudTransactionService) Get(ctx context.Context, transactionToken string, opts ...option.RequestOption) (res *FraudTransactionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionToken == "" {
		err = errors.New("missing required transaction_token parameter")
		return
	}
	path := fmt.Sprintf("v1/fraud/transactions/%s", transactionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Report fraud for a specific transaction token by providing details such as fraud
// type, fraud status, and any additional comments.
func (r *FraudTransactionService) Report(ctx context.Context, transactionToken string, body FraudTransactionReportParams, opts ...option.RequestOption) (res *FraudTransactionReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionToken == "" {
		err = errors.New("missing required transaction_token parameter")
		return
	}
	path := fmt.Sprintf("v1/fraud/transactions/%s", transactionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FraudTransactionGetResponse struct {
	// The fraud status of the transaction, string (enum) supporting the following
	// values:
	//
	//   - `SUSPECTED_FRAUD`: The transaction is suspected to be fraudulent, but this
	//     hasn’t been confirmed.
	//   - `FRAUDULENT`: The transaction is confirmed to be fraudulent. A transaction may
	//     immediately be moved into this state, or be graduated into this state from the
	//     `SUSPECTED_FRAUD` state.
	//   - `NOT_FRAUDULENT`: The transaction is (explicitly) marked as not fraudulent. A
	//     transaction may immediately be moved into this state, or be graduated into
	//     this state from the `SUSPECTED_FRAUD` state.
	//   - `NO_REPORTED_FRAUD`: Indicates that no fraud report exists for the
	//     transaction. It is the default state for transactions that have not been
	//     analyzed or associated with any known fraudulent activity.
	FraudStatus FraudTransactionGetResponseFraudStatus `json:"fraud_status,required"`
	// The universally unique identifier (UUID) associated with the transaction being
	// reported.
	TransactionToken string `json:"transaction_token,required" format:"uuid"`
	// Provides additional context or details about the fraud report.
	Comment string `json:"comment"`
	// Timestamp representing when the fraud report was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies the type or category of fraud that the transaction is suspected or
	// confirmed to involve, string (enum) supporting the following values:
	//
	//   - `FIRST_PARTY_FRAUD`: First-party fraud occurs when a legitimate account or
	//     cardholder intentionally misuses financial services for personal gain. This
	//     includes actions such as disputing legitimate transactions to obtain a refund,
	//     abusing return policies, or defaulting on credit obligations without intent to
	//     repay.
	//   - `ACCOUNT_TAKEOVER`: Account takeover fraud occurs when a fraudster gains
	//     unauthorized access to an existing account, modifies account settings, and
	//     carries out fraudulent transactions.
	//   - `CARD_COMPROMISED`: Card compromised fraud occurs when a fraudster gains
	//     access to card details without taking over the account, such as through
	//     physical card theft, cloning, or online data breaches.
	//   - `IDENTITY_THEFT`: Identity theft fraud occurs when a fraudster uses stolen
	//     personal information, such as Social Security numbers or addresses, to open
	//     accounts, apply for loans, or conduct financial transactions in someone's
	//     name.
	//   - `CARDHOLDER_MANIPULATION`: This type of fraud occurs when a fraudster
	//     manipulates or coerces a legitimate cardholder into unauthorized transactions,
	//     often through social engineering tactics.
	FraudType FraudTransactionGetResponseFraudType `json:"fraud_type"`
	// Timestamp representing the last update to the fraud report.
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      fraudTransactionGetResponseJSON `json:"-"`
}

// fraudTransactionGetResponseJSON contains the JSON metadata for the struct
// [FraudTransactionGetResponse]
type fraudTransactionGetResponseJSON struct {
	FraudStatus      apijson.Field
	TransactionToken apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	FraudType        apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FraudTransactionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudTransactionGetResponseJSON) RawJSON() string {
	return r.raw
}

// The fraud status of the transaction, string (enum) supporting the following
// values:
//
//   - `SUSPECTED_FRAUD`: The transaction is suspected to be fraudulent, but this
//     hasn’t been confirmed.
//   - `FRAUDULENT`: The transaction is confirmed to be fraudulent. A transaction may
//     immediately be moved into this state, or be graduated into this state from the
//     `SUSPECTED_FRAUD` state.
//   - `NOT_FRAUDULENT`: The transaction is (explicitly) marked as not fraudulent. A
//     transaction may immediately be moved into this state, or be graduated into
//     this state from the `SUSPECTED_FRAUD` state.
//   - `NO_REPORTED_FRAUD`: Indicates that no fraud report exists for the
//     transaction. It is the default state for transactions that have not been
//     analyzed or associated with any known fraudulent activity.
type FraudTransactionGetResponseFraudStatus string

const (
	FraudTransactionGetResponseFraudStatusSuspectedFraud  FraudTransactionGetResponseFraudStatus = "SUSPECTED_FRAUD"
	FraudTransactionGetResponseFraudStatusFraudulent      FraudTransactionGetResponseFraudStatus = "FRAUDULENT"
	FraudTransactionGetResponseFraudStatusNotFraudulent   FraudTransactionGetResponseFraudStatus = "NOT_FRAUDULENT"
	FraudTransactionGetResponseFraudStatusNoReportedFraud FraudTransactionGetResponseFraudStatus = "NO_REPORTED_FRAUD"
)

func (r FraudTransactionGetResponseFraudStatus) IsKnown() bool {
	switch r {
	case FraudTransactionGetResponseFraudStatusSuspectedFraud, FraudTransactionGetResponseFraudStatusFraudulent, FraudTransactionGetResponseFraudStatusNotFraudulent, FraudTransactionGetResponseFraudStatusNoReportedFraud:
		return true
	}
	return false
}

// Specifies the type or category of fraud that the transaction is suspected or
// confirmed to involve, string (enum) supporting the following values:
//
//   - `FIRST_PARTY_FRAUD`: First-party fraud occurs when a legitimate account or
//     cardholder intentionally misuses financial services for personal gain. This
//     includes actions such as disputing legitimate transactions to obtain a refund,
//     abusing return policies, or defaulting on credit obligations without intent to
//     repay.
//   - `ACCOUNT_TAKEOVER`: Account takeover fraud occurs when a fraudster gains
//     unauthorized access to an existing account, modifies account settings, and
//     carries out fraudulent transactions.
//   - `CARD_COMPROMISED`: Card compromised fraud occurs when a fraudster gains
//     access to card details without taking over the account, such as through
//     physical card theft, cloning, or online data breaches.
//   - `IDENTITY_THEFT`: Identity theft fraud occurs when a fraudster uses stolen
//     personal information, such as Social Security numbers or addresses, to open
//     accounts, apply for loans, or conduct financial transactions in someone's
//     name.
//   - `CARDHOLDER_MANIPULATION`: This type of fraud occurs when a fraudster
//     manipulates or coerces a legitimate cardholder into unauthorized transactions,
//     often through social engineering tactics.
type FraudTransactionGetResponseFraudType string

const (
	FraudTransactionGetResponseFraudTypeFirstPartyFraud        FraudTransactionGetResponseFraudType = "FIRST_PARTY_FRAUD"
	FraudTransactionGetResponseFraudTypeAccountTakeover        FraudTransactionGetResponseFraudType = "ACCOUNT_TAKEOVER"
	FraudTransactionGetResponseFraudTypeCardCompromised        FraudTransactionGetResponseFraudType = "CARD_COMPROMISED"
	FraudTransactionGetResponseFraudTypeIdentityTheft          FraudTransactionGetResponseFraudType = "IDENTITY_THEFT"
	FraudTransactionGetResponseFraudTypeCardholderManipulation FraudTransactionGetResponseFraudType = "CARDHOLDER_MANIPULATION"
)

func (r FraudTransactionGetResponseFraudType) IsKnown() bool {
	switch r {
	case FraudTransactionGetResponseFraudTypeFirstPartyFraud, FraudTransactionGetResponseFraudTypeAccountTakeover, FraudTransactionGetResponseFraudTypeCardCompromised, FraudTransactionGetResponseFraudTypeIdentityTheft, FraudTransactionGetResponseFraudTypeCardholderManipulation:
		return true
	}
	return false
}

type FraudTransactionReportResponse struct {
	// The fraud status of the transaction, string (enum) supporting the following
	// values:
	//
	//   - `SUSPECTED_FRAUD`: The transaction is suspected to be fraudulent, but this
	//     hasn’t been confirmed.
	//   - `FRAUDULENT`: The transaction is confirmed to be fraudulent. A transaction may
	//     immediately be moved into this state, or be graduated into this state from the
	//     `SUSPECTED_FRAUD` state.
	//   - `NOT_FRAUDULENT`: The transaction is (explicitly) marked as not fraudulent. A
	//     transaction may immediately be moved into this state, or be graduated into
	//     this state from the `SUSPECTED_FRAUD` state.
	//   - `NO_REPORTED_FRAUD`: Indicates that no fraud report exists for the
	//     transaction. It is the default state for transactions that have not been
	//     analyzed or associated with any known fraudulent activity.
	FraudStatus FraudTransactionReportResponseFraudStatus `json:"fraud_status,required"`
	// The universally unique identifier (UUID) associated with the transaction being
	// reported.
	TransactionToken string `json:"transaction_token,required" format:"uuid"`
	// Provides additional context or details about the fraud report.
	Comment string `json:"comment"`
	// Timestamp representing when the fraud report was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies the type or category of fraud that the transaction is suspected or
	// confirmed to involve, string (enum) supporting the following values:
	//
	//   - `FIRST_PARTY_FRAUD`: First-party fraud occurs when a legitimate account or
	//     cardholder intentionally misuses financial services for personal gain. This
	//     includes actions such as disputing legitimate transactions to obtain a refund,
	//     abusing return policies, or defaulting on credit obligations without intent to
	//     repay.
	//   - `ACCOUNT_TAKEOVER`: Account takeover fraud occurs when a fraudster gains
	//     unauthorized access to an existing account, modifies account settings, and
	//     carries out fraudulent transactions.
	//   - `CARD_COMPROMISED`: Card compromised fraud occurs when a fraudster gains
	//     access to card details without taking over the account, such as through
	//     physical card theft, cloning, or online data breaches.
	//   - `IDENTITY_THEFT`: Identity theft fraud occurs when a fraudster uses stolen
	//     personal information, such as Social Security numbers or addresses, to open
	//     accounts, apply for loans, or conduct financial transactions in someone's
	//     name.
	//   - `CARDHOLDER_MANIPULATION`: This type of fraud occurs when a fraudster
	//     manipulates or coerces a legitimate cardholder into unauthorized transactions,
	//     often through social engineering tactics.
	FraudType FraudTransactionReportResponseFraudType `json:"fraud_type"`
	// Timestamp representing the last update to the fraud report.
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      fraudTransactionReportResponseJSON `json:"-"`
}

// fraudTransactionReportResponseJSON contains the JSON metadata for the struct
// [FraudTransactionReportResponse]
type fraudTransactionReportResponseJSON struct {
	FraudStatus      apijson.Field
	TransactionToken apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	FraudType        apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FraudTransactionReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudTransactionReportResponseJSON) RawJSON() string {
	return r.raw
}

// The fraud status of the transaction, string (enum) supporting the following
// values:
//
//   - `SUSPECTED_FRAUD`: The transaction is suspected to be fraudulent, but this
//     hasn’t been confirmed.
//   - `FRAUDULENT`: The transaction is confirmed to be fraudulent. A transaction may
//     immediately be moved into this state, or be graduated into this state from the
//     `SUSPECTED_FRAUD` state.
//   - `NOT_FRAUDULENT`: The transaction is (explicitly) marked as not fraudulent. A
//     transaction may immediately be moved into this state, or be graduated into
//     this state from the `SUSPECTED_FRAUD` state.
//   - `NO_REPORTED_FRAUD`: Indicates that no fraud report exists for the
//     transaction. It is the default state for transactions that have not been
//     analyzed or associated with any known fraudulent activity.
type FraudTransactionReportResponseFraudStatus string

const (
	FraudTransactionReportResponseFraudStatusSuspectedFraud  FraudTransactionReportResponseFraudStatus = "SUSPECTED_FRAUD"
	FraudTransactionReportResponseFraudStatusFraudulent      FraudTransactionReportResponseFraudStatus = "FRAUDULENT"
	FraudTransactionReportResponseFraudStatusNotFraudulent   FraudTransactionReportResponseFraudStatus = "NOT_FRAUDULENT"
	FraudTransactionReportResponseFraudStatusNoReportedFraud FraudTransactionReportResponseFraudStatus = "NO_REPORTED_FRAUD"
)

func (r FraudTransactionReportResponseFraudStatus) IsKnown() bool {
	switch r {
	case FraudTransactionReportResponseFraudStatusSuspectedFraud, FraudTransactionReportResponseFraudStatusFraudulent, FraudTransactionReportResponseFraudStatusNotFraudulent, FraudTransactionReportResponseFraudStatusNoReportedFraud:
		return true
	}
	return false
}

// Specifies the type or category of fraud that the transaction is suspected or
// confirmed to involve, string (enum) supporting the following values:
//
//   - `FIRST_PARTY_FRAUD`: First-party fraud occurs when a legitimate account or
//     cardholder intentionally misuses financial services for personal gain. This
//     includes actions such as disputing legitimate transactions to obtain a refund,
//     abusing return policies, or defaulting on credit obligations without intent to
//     repay.
//   - `ACCOUNT_TAKEOVER`: Account takeover fraud occurs when a fraudster gains
//     unauthorized access to an existing account, modifies account settings, and
//     carries out fraudulent transactions.
//   - `CARD_COMPROMISED`: Card compromised fraud occurs when a fraudster gains
//     access to card details without taking over the account, such as through
//     physical card theft, cloning, or online data breaches.
//   - `IDENTITY_THEFT`: Identity theft fraud occurs when a fraudster uses stolen
//     personal information, such as Social Security numbers or addresses, to open
//     accounts, apply for loans, or conduct financial transactions in someone's
//     name.
//   - `CARDHOLDER_MANIPULATION`: This type of fraud occurs when a fraudster
//     manipulates or coerces a legitimate cardholder into unauthorized transactions,
//     often through social engineering tactics.
type FraudTransactionReportResponseFraudType string

const (
	FraudTransactionReportResponseFraudTypeFirstPartyFraud        FraudTransactionReportResponseFraudType = "FIRST_PARTY_FRAUD"
	FraudTransactionReportResponseFraudTypeAccountTakeover        FraudTransactionReportResponseFraudType = "ACCOUNT_TAKEOVER"
	FraudTransactionReportResponseFraudTypeCardCompromised        FraudTransactionReportResponseFraudType = "CARD_COMPROMISED"
	FraudTransactionReportResponseFraudTypeIdentityTheft          FraudTransactionReportResponseFraudType = "IDENTITY_THEFT"
	FraudTransactionReportResponseFraudTypeCardholderManipulation FraudTransactionReportResponseFraudType = "CARDHOLDER_MANIPULATION"
)

func (r FraudTransactionReportResponseFraudType) IsKnown() bool {
	switch r {
	case FraudTransactionReportResponseFraudTypeFirstPartyFraud, FraudTransactionReportResponseFraudTypeAccountTakeover, FraudTransactionReportResponseFraudTypeCardCompromised, FraudTransactionReportResponseFraudTypeIdentityTheft, FraudTransactionReportResponseFraudTypeCardholderManipulation:
		return true
	}
	return false
}

type FraudTransactionReportParams struct {
	// The fraud status of the transaction, string (enum) supporting the following
	// values:
	//
	//   - `SUSPECTED_FRAUD`: The transaction is suspected to be fraudulent, but this
	//     hasn’t been confirmed.
	//   - `FRAUDULENT`: The transaction is confirmed to be fraudulent. A transaction may
	//     immediately be moved into this state, or be graduated into this state from the
	//     `SUSPECTED_FRAUD` state.
	//   - `NOT_FRAUDULENT`: The transaction is (explicitly) marked as not fraudulent. A
	//     transaction may immediately be moved into this state, or be graduated into
	//     this state from the `SUSPECTED_FRAUD` state.
	FraudStatus param.Field[FraudTransactionReportParamsFraudStatus] `json:"fraud_status,required"`
	// Optional field providing additional information or context about why the
	// transaction is considered fraudulent.
	Comment param.Field[string] `json:"comment"`
	// Specifies the type or category of fraud that the transaction is suspected or
	// confirmed to involve, string (enum) supporting the following values:
	//
	//   - `FIRST_PARTY_FRAUD`: First-party fraud occurs when a legitimate account or
	//     cardholder intentionally misuses financial services for personal gain. This
	//     includes actions such as disputing legitimate transactions to obtain a refund,
	//     abusing return policies, or defaulting on credit obligations without intent to
	//     repay.
	//   - `ACCOUNT_TAKEOVER`: Account takeover fraud occurs when a fraudster gains
	//     unauthorized access to an existing account, modifies account settings, and
	//     carries out fraudulent transactions.
	//   - `CARD_COMPROMISED`: Card compromised fraud occurs when a fraudster gains
	//     access to card details without taking over the account, such as through
	//     physical card theft, cloning, or online data breaches.
	//   - `IDENTITY_THEFT`: Identity theft fraud occurs when a fraudster uses stolen
	//     personal information, such as Social Security numbers or addresses, to open
	//     accounts, apply for loans, or conduct financial transactions in someone's
	//     name.
	//   - `CARDHOLDER_MANIPULATION`: This type of fraud occurs when a fraudster
	//     manipulates or coerces a legitimate cardholder into unauthorized transactions,
	//     often through social engineering tactics.
	FraudType param.Field[FraudTransactionReportParamsFraudType] `json:"fraud_type"`
}

func (r FraudTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fraud status of the transaction, string (enum) supporting the following
// values:
//
//   - `SUSPECTED_FRAUD`: The transaction is suspected to be fraudulent, but this
//     hasn’t been confirmed.
//   - `FRAUDULENT`: The transaction is confirmed to be fraudulent. A transaction may
//     immediately be moved into this state, or be graduated into this state from the
//     `SUSPECTED_FRAUD` state.
//   - `NOT_FRAUDULENT`: The transaction is (explicitly) marked as not fraudulent. A
//     transaction may immediately be moved into this state, or be graduated into
//     this state from the `SUSPECTED_FRAUD` state.
type FraudTransactionReportParamsFraudStatus string

const (
	FraudTransactionReportParamsFraudStatusSuspectedFraud FraudTransactionReportParamsFraudStatus = "SUSPECTED_FRAUD"
	FraudTransactionReportParamsFraudStatusFraudulent     FraudTransactionReportParamsFraudStatus = "FRAUDULENT"
	FraudTransactionReportParamsFraudStatusNotFraudulent  FraudTransactionReportParamsFraudStatus = "NOT_FRAUDULENT"
)

func (r FraudTransactionReportParamsFraudStatus) IsKnown() bool {
	switch r {
	case FraudTransactionReportParamsFraudStatusSuspectedFraud, FraudTransactionReportParamsFraudStatusFraudulent, FraudTransactionReportParamsFraudStatusNotFraudulent:
		return true
	}
	return false
}

// Specifies the type or category of fraud that the transaction is suspected or
// confirmed to involve, string (enum) supporting the following values:
//
//   - `FIRST_PARTY_FRAUD`: First-party fraud occurs when a legitimate account or
//     cardholder intentionally misuses financial services for personal gain. This
//     includes actions such as disputing legitimate transactions to obtain a refund,
//     abusing return policies, or defaulting on credit obligations without intent to
//     repay.
//   - `ACCOUNT_TAKEOVER`: Account takeover fraud occurs when a fraudster gains
//     unauthorized access to an existing account, modifies account settings, and
//     carries out fraudulent transactions.
//   - `CARD_COMPROMISED`: Card compromised fraud occurs when a fraudster gains
//     access to card details without taking over the account, such as through
//     physical card theft, cloning, or online data breaches.
//   - `IDENTITY_THEFT`: Identity theft fraud occurs when a fraudster uses stolen
//     personal information, such as Social Security numbers or addresses, to open
//     accounts, apply for loans, or conduct financial transactions in someone's
//     name.
//   - `CARDHOLDER_MANIPULATION`: This type of fraud occurs when a fraudster
//     manipulates or coerces a legitimate cardholder into unauthorized transactions,
//     often through social engineering tactics.
type FraudTransactionReportParamsFraudType string

const (
	FraudTransactionReportParamsFraudTypeFirstPartyFraud        FraudTransactionReportParamsFraudType = "FIRST_PARTY_FRAUD"
	FraudTransactionReportParamsFraudTypeAccountTakeover        FraudTransactionReportParamsFraudType = "ACCOUNT_TAKEOVER"
	FraudTransactionReportParamsFraudTypeCardCompromised        FraudTransactionReportParamsFraudType = "CARD_COMPROMISED"
	FraudTransactionReportParamsFraudTypeIdentityTheft          FraudTransactionReportParamsFraudType = "IDENTITY_THEFT"
	FraudTransactionReportParamsFraudTypeCardholderManipulation FraudTransactionReportParamsFraudType = "CARDHOLDER_MANIPULATION"
)

func (r FraudTransactionReportParamsFraudType) IsKnown() bool {
	switch r {
	case FraudTransactionReportParamsFraudTypeFirstPartyFraud, FraudTransactionReportParamsFraudTypeAccountTakeover, FraudTransactionReportParamsFraudTypeCardCompromised, FraudTransactionReportParamsFraudTypeIdentityTheft, FraudTransactionReportParamsFraudTypeCardholderManipulation:
		return true
	}
	return false
}
