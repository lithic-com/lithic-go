package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	"github.com/lithic-com/lithic-go/core/query"
)

type FinancialAccountBalanceListParams struct {
	// UTC date of the balance to retrieve. Defaults to latest available balance
	BalanceDate field.Field[time.Time] `query:"balance_date" format:"date-time"`
	// Balance after a given financial event occured. For example, passing the
	// event_token of a $5 CARD_CLEARING financial event will return a balance
	// decreased by $5
	LastTransactionEventToken field.Field[string] `query:"last_transaction_event_token" format:"uuid"`
}

// URLQuery serializes FinancialAccountBalanceListParams into a url.Values of the
// query parameters associated with this value
func (r FinancialAccountBalanceListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
