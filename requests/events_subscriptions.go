package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type SubscriptionNewParams struct {
	// Event subscription description.
	Description field.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled field.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes field.Field[[]SubscriptionNewParamsEventTypes] `json:"event_types"`
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL field.Field[string] `json:"url,required" format:"uri"`
}

// MarshalJSON serializes SubscriptionNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type SubscriptionNewParamsEventTypes string

const (
	SubscriptionNewParamsEventTypesDisputeUpdated                           SubscriptionNewParamsEventTypes = "dispute.updated"
	SubscriptionNewParamsEventTypesDigitalWalletTokenizationApprovalRequest SubscriptionNewParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type SubscriptionUpdateParams struct {
	// Event subscription description.
	Description field.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled field.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes field.Field[[]SubscriptionUpdateParamsEventTypes] `json:"event_types"`
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL field.Field[string] `json:"url,required" format:"uri"`
}

// MarshalJSON serializes SubscriptionUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type SubscriptionUpdateParamsEventTypes string

const (
	SubscriptionUpdateParamsEventTypesDisputeUpdated                           SubscriptionUpdateParamsEventTypes = "dispute.updated"
	SubscriptionUpdateParamsEventTypesDigitalWalletTokenizationApprovalRequest SubscriptionUpdateParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type SubscriptionListParams struct {
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter field.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore field.Field[string] `query:"ending_before"`
}

// URLQuery serializes SubscriptionListParams into a url.Values of the query
// parameters associated with this value
func (r *SubscriptionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type SubscriptionRecoverParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes SubscriptionRecoverParams into a url.Values of the query
// parameters associated with this value
func (r *SubscriptionRecoverParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type SubscriptionReplayMissingParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes SubscriptionReplayMissingParams into a url.Values of the
// query parameters associated with this value
func (r *SubscriptionReplayMissingParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
