package responses

import (
	apijson "github.com/lithic-com/lithic-go/core/json"
)

type FinancialAccountBalanceListResponse struct {
	Data []Balance `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    FinancialAccountBalanceListResponseJSON
}

type FinancialAccountBalanceListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	Raw     []byte
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// FinancialAccountBalanceListResponse using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *FinancialAccountBalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
