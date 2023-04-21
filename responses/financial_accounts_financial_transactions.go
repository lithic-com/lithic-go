package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type FinancialTransactionListResponse struct {
	Data []FinancialTransaction `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    FinancialTransactionListResponseJSON
}

type FinancialTransactionListResponseJSON struct {
	Data    pjson.Metadata
	HasMore pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// FinancialTransactionListResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *FinancialTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
