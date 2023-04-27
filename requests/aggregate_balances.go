package requests

import (
	"net/url"

	"github.com/lithic-com/lithic-go/internal/field"
	"github.com/lithic-com/lithic-go/internal/query"
)

type AggregateBalanceListParams struct {
	// Get the aggregate balance for a given Financial Account type.
	FinancialAccountType field.Field[AggregateBalanceListParamsFinancialAccountType] `query:"financial_account_type"`
}

// URLQuery serializes AggregateBalanceListParams into a url.Values of the query
// parameters associated with this value
func (r AggregateBalanceListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type AggregateBalanceListParamsFinancialAccountType string

const (
	AggregateBalanceListParamsFinancialAccountTypeIssuing AggregateBalanceListParamsFinancialAccountType = "ISSUING"
	AggregateBalanceListParamsFinancialAccountTypeReserve AggregateBalanceListParamsFinancialAccountType = "RESERVE"
)
