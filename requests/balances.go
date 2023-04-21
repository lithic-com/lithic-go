package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	"github.com/lithic-com/lithic-go/core/query"
)

type BalanceListParams struct {
	// List balances for all financial accounts of a given account_token.
	AccountToken field.Field[string] `query:"account_token" format:"uuid"`
	// UTC date and time of the balances to retrieve. Defaults to latest available
	// balances
	BalanceDate field.Field[time.Time] `query:"balance_date" format:"date-time"`
	// List balances for a given Financial Account type.
	FinancialAccountType field.Field[BalanceListParamsFinancialAccountType] `query:"financial_account_type"`
}

// URLQuery serializes BalanceListParams into a url.Values of the query parameters
// associated with this value
func (r BalanceListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type BalanceListParamsFinancialAccountType string

const (
	BalanceListParamsFinancialAccountTypeIssuing BalanceListParamsFinancialAccountType = "ISSUING"
	BalanceListParamsFinancialAccountTypeReserve BalanceListParamsFinancialAccountType = "RESERVE"
)
