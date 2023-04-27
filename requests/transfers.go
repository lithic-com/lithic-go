package requests

import (
	"github.com/lithic-com/lithic-go/internal/field"
	apijson "github.com/lithic-com/lithic-go/internal/json"
)

type TransferNewParams struct {
	// Financial Account
	From field.Field[FinancialAccount] `json:"from,required"`
	// Financial Account
	To field.Field[FinancialAccount] `json:"to,required"`
	// Amount to be transferred in the currency’s smallest unit (e.g., cents for USD).
	// This should always be a positive value.
	Amount field.Field[int64] `json:"amount,required"`
	// Optional descriptor for the transfer.
	Memo field.Field[string] `json:"memo"`
	// Customer-provided transaction_token that will serve as an idempotency token.
	TransactionToken field.Field[string] `json:"transaction_token" format:"uuid"`
}

// MarshalJSON serializes TransferNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r TransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
