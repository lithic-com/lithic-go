package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type FinancialAccount struct {
	// Account number for your Lithic-assigned bank account number, if applicable.
	AccountNumber field.Field[string] `json:"account_number"`
	// Date and time for when the financial account was first created.
	Created field.Field[time.Time] `json:"created,required" format:"date-time"`
	// Routing number for your Lithic-assigned bank account number, if applicable.
	RoutingNumber field.Field[string] `json:"routing_number"`
	// Globally unique identifier for the financial account.
	Token field.Field[string] `json:"token,required" format:"uuid"`
	// Type of financial account
	Type field.Field[FinancialAccountType] `json:"type,required"`
	// Date and time for when the financial account was last updated.
	Updated field.Field[time.Time] `json:"updated,required" format:"date-time"`
}

// MarshalJSON serializes FinancialAccount into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r FinancialAccount) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type FinancialAccountType string

const (
	FinancialAccountTypeIssuing FinancialAccountType = "ISSUING"
	FinancialAccountTypeReserve FinancialAccountType = "RESERVE"
)

type FinancialAccountListParams struct {
	// List financial accounts for a given account_token
	AccountToken field.Field[string] `query:"account_token" format:"uuid"`
	// List financial accounts of a given type
	Type field.Field[FinancialAccountListParamsType] `query:"type"`
}

// URLQuery serializes FinancialAccountListParams into a url.Values of the query
// parameters associated with this value
func (r FinancialAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type FinancialAccountListParamsType string

const (
	FinancialAccountListParamsTypeIssuing FinancialAccountListParamsType = "ISSUING"
	FinancialAccountListParamsTypeReserve FinancialAccountListParamsType = "RESERVE"
)
