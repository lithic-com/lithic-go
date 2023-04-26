package responses

import (
	"time"

	apijson "github.com/lithic-com/lithic-go/core/json"
)

type AggregateBalance struct {
	// Type of financial account
	FinancialAccountType AggregateBalanceFinancialAccountType `json:"financial_account_type,required"`
	// 3-digit alphabetic ISO 4217 code for the local currency of the balance.
	Currency string `json:"currency,required"`
	// Funds available for spend in the currency's smallest unit (e.g., cents for USD)
	AvailableAmount int64 `json:"available_amount,required"`
	// Funds not available for spend due to card authorizations or pending ACH release.
	// Shown in the currency's smallest unit (e.g., cents for USD)
	PendingAmount int64 `json:"pending_amount,required"`
	// The sum of available and pending balance in the currency's smallest unit (e.g.,
	// cents for USD)
	TotalAmount int64 `json:"total_amount,required"`
	// Date and time for when the balance was first created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Date and time for when the balance was last updated.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Globally unique identifier for the last transaction that impacted this balance
	LastTransactionToken string `json:"last_transaction_token,required" format:"uuid"`
	// Globally unique identifier for the last transaction event that impacted this
	// balance
	LastTransactionEventToken string `json:"last_transaction_event_token,required" format:"uuid"`
	// Globally unique identifier for the financial account that had its balance
	// updated most recently
	LastFinancialAccountToken string `json:"last_financial_account_token,required" format:"uuid"`
	JSON                      AggregateBalanceJSON
}

type AggregateBalanceJSON struct {
	FinancialAccountType      apijson.Metadata
	Currency                  apijson.Metadata
	AvailableAmount           apijson.Metadata
	PendingAmount             apijson.Metadata
	TotalAmount               apijson.Metadata
	Created                   apijson.Metadata
	Updated                   apijson.Metadata
	LastTransactionToken      apijson.Metadata
	LastTransactionEventToken apijson.Metadata
	LastFinancialAccountToken apijson.Metadata
	Raw                       []byte
	Extras                    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AggregateBalance using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AggregateBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AggregateBalanceFinancialAccountType string

const (
	AggregateBalanceFinancialAccountTypeIssuing AggregateBalanceFinancialAccountType = "ISSUING"
	AggregateBalanceFinancialAccountTypeReserve AggregateBalanceFinancialAccountType = "RESERVE"
)

type AggregateBalanceListResponse struct {
	Data []AggregateBalance `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    AggregateBalanceListResponseJSON
}

type AggregateBalanceListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	Raw     []byte
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AggregateBalanceListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AggregateBalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
