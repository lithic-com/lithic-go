package responses

import (
	"time"

	apijson "github.com/lithic-com/lithic-go/core/json"
)

type Balance struct {
	// Funds available for spend in the currency's smallest unit (e.g., cents for USD)
	AvailableAmount int64 `json:"available_amount,required"`
	// Date and time for when the balance was first created.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-digit alphabetic ISO 4217 code for the local currency of the balance.
	Currency string `json:"currency,required"`
	// Globally unique identifier for the last financial transaction event that
	// impacted this balance.
	LastTransactionEventToken string `json:"last_transaction_event_token,required" format:"uuid"`
	// Globally unique identifier for the last financial transaction that impacted this
	// balance.
	LastTransactionToken string `json:"last_transaction_token,required" format:"uuid"`
	// Funds not available for spend due to card authorizations or pending ACH release.
	// Shown in the currency's smallest unit (e.g., cents for USD).
	PendingAmount int64 `json:"pending_amount,required"`
	// Globally unique identifier for the financial account that holds this balance.
	Token string `json:"token,required" format:"uuid"`
	// The sum of available and pending balance in the currency's smallest unit (e.g.,
	// cents for USD).
	TotalAmount int64 `json:"total_amount,required"`
	// Type of financial account.
	Type BalanceType `json:"type,required"`
	// Date and time for when the balance was last updated.
	Updated time.Time `json:"updated,required" format:"date-time"`
	JSON    BalanceJSON
}

type BalanceJSON struct {
	AvailableAmount           apijson.Metadata
	Created                   apijson.Metadata
	Currency                  apijson.Metadata
	LastTransactionEventToken apijson.Metadata
	LastTransactionToken      apijson.Metadata
	PendingAmount             apijson.Metadata
	Token                     apijson.Metadata
	TotalAmount               apijson.Metadata
	Type                      apijson.Metadata
	Updated                   apijson.Metadata
	Raw                       []byte
	Extras                    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Balance using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Balance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceType string

const (
	BalanceTypeIssuing BalanceType = "ISSUING"
	BalanceTypeReserve BalanceType = "RESERVE"
)

type BalanceListResponse struct {
	Data []Balance `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    BalanceListResponseJSON
}

type BalanceListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	Raw     []byte
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BalanceListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
