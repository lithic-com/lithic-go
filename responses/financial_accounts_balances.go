package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type FinancialAccountBalanceListResponse struct {
	Data []Balance `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    FinancialAccountBalanceListResponseJSON
}

type FinancialAccountBalanceListResponseJSON struct {
	Data    pjson.Metadata
	HasMore pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// FinancialAccountBalanceListResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *FinancialAccountBalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
