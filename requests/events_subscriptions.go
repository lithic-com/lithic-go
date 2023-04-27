package requests

import (
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/field"
	apijson "github.com/lithic-com/lithic-go/internal/json"
	"github.com/lithic-com/lithic-go/internal/query"
)

type EventSubscriptionNewParams struct {
	// Event subscription description.
	Description field.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled field.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes field.Field[[]EventSubscriptionNewParamsEventTypes] `json:"event_types"`
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL field.Field[string] `json:"url,required" format:"uri"`
}

// MarshalJSON serializes EventSubscriptionNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventSubscriptionNewParamsEventTypes string

const (
	EventSubscriptionNewParamsEventTypesDisputeUpdated                           EventSubscriptionNewParamsEventTypes = "dispute.updated"
	EventSubscriptionNewParamsEventTypesDigitalWalletTokenizationApprovalRequest EventSubscriptionNewParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type EventSubscriptionUpdateParams struct {
	// Event subscription description.
	Description field.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled field.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes field.Field[[]EventSubscriptionUpdateParamsEventTypes] `json:"event_types"`
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL field.Field[string] `json:"url,required" format:"uri"`
}

// MarshalJSON serializes EventSubscriptionUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r EventSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventSubscriptionUpdateParamsEventTypes string

const (
	EventSubscriptionUpdateParamsEventTypesDisputeUpdated                           EventSubscriptionUpdateParamsEventTypes = "dispute.updated"
	EventSubscriptionUpdateParamsEventTypesDigitalWalletTokenizationApprovalRequest EventSubscriptionUpdateParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type EventSubscriptionListParams struct {
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter field.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore field.Field[string] `query:"ending_before"`
}

// URLQuery serializes EventSubscriptionListParams into a url.Values of the query
// parameters associated with this value
func (r EventSubscriptionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type EventSubscriptionRecoverParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes EventSubscriptionRecoverParams into a url.Values of the
// query parameters associated with this value
func (r EventSubscriptionRecoverParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type EventSubscriptionReplayMissingParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes EventSubscriptionReplayMissingParams into a url.Values of
// the query parameters associated with this value
func (r EventSubscriptionReplayMissingParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
