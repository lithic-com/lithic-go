package responses

import (
	apijson "github.com/lithic-com/lithic-go/core/json"
)

type FinancialTransactionListResponse struct {
	Data []FinancialTransaction `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    FinancialTransactionListResponseJSON
}

type FinancialTransactionListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	Raw     []byte
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// FinancialTransactionListResponse using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *FinancialTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
