package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type AccountUpdateParams struct {
	// Amount (in cents) for the account's new daily spend limit. Note that a spend
	// limit of 0 is effectively no limit, and should only be used to reset or remove a
	// prior limit. Only a limit of 1 or above will result in declined transactions due
	// to checks against the account limit.
	DailySpendLimit field.Field[int64] `json:"daily_spend_limit"`
	// Amount (in cents) for the account's new lifetime limit. Once this limit is
	// reached, no transactions will be accepted on any card created for this account
	// until the limit is updated. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the
	// account limit.
	LifetimeSpendLimit field.Field[int64] `json:"lifetime_spend_limit"`
	// Amount (in cents) for the account's new monthly spend limit. Note that a spend
	// limit of 0 is effectively no limit, and should only be used to reset or remove a
	// prior limit. Only a limit of 1 or above will result in declined transactions due
	// to checks against the account limit.
	MonthlySpendLimit field.Field[int64] `json:"monthly_spend_limit"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	VerificationAddress field.Field[AccountUpdateParamsVerificationAddress] `json:"verification_address"`
	// Account states.
	State field.Field[AccountUpdateParamsState] `json:"state"`
}

// MarshalJSON serializes AccountUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type AccountUpdateParamsVerificationAddress struct {
	Address1   field.Field[string] `json:"address1"`
	Address2   field.Field[string] `json:"address2"`
	City       field.Field[string] `json:"city"`
	State      field.Field[string] `json:"state"`
	PostalCode field.Field[string] `json:"postal_code"`
	Country    field.Field[string] `json:"country"`
}

type AccountUpdateParamsState string

const (
	AccountUpdateParamsStateActive AccountUpdateParamsState = "ACTIVE"
	AccountUpdateParamsStatePaused AccountUpdateParamsState = "PAUSED"
)

type AccountListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
	// Page (for pagination).
	Page field.Field[int64] `query:"page"`
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
}

// URLQuery serializes AccountListParams into a url.Values of the query parameters
// associated with this value
func (r AccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
