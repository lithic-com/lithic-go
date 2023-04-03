package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/core/fields"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type SubscriptionNewParams struct {
	// Event subscription description.
	Description fields.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled fields.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes fields.Field[[]SubscriptionNewParamsEventTypes] `json:"event_types"`
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL fields.Field[string] `json:"url,required" format:"uri"`
}

// MarshalJSON serializes SubscriptionNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SubscriptionNewParams) String() (result string) {
	return fmt.Sprintf("&SubscriptionNewParams{Description:%s Disabled:%s EventTypes:%s URL:%s}", r.Description, r.Disabled, core.Fmt(r.EventTypes), r.URL)
}

type SubscriptionNewParamsEventTypes string

const (
	SubscriptionNewParamsEventTypesDisputeUpdated                           SubscriptionNewParamsEventTypes = "dispute.updated"
	SubscriptionNewParamsEventTypesDigitalWalletTokenizationApprovalRequest SubscriptionNewParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type SubscriptionUpdateParams struct {
	// Event subscription description.
	Description fields.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled fields.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes fields.Field[[]SubscriptionUpdateParamsEventTypes] `json:"event_types"`
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL fields.Field[string] `json:"url,required" format:"uri"`
}

// MarshalJSON serializes SubscriptionUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SubscriptionUpdateParams) String() (result string) {
	return fmt.Sprintf("&SubscriptionUpdateParams{Description:%s Disabled:%s EventTypes:%s URL:%s}", r.Description, r.Disabled, core.Fmt(r.EventTypes), r.URL)
}

type SubscriptionUpdateParamsEventTypes string

const (
	SubscriptionUpdateParamsEventTypesDisputeUpdated                           SubscriptionUpdateParamsEventTypes = "dispute.updated"
	SubscriptionUpdateParamsEventTypesDigitalWalletTokenizationApprovalRequest SubscriptionUpdateParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type SubscriptionListParams struct {
	// Page size (for pagination).
	PageSize fields.Field[int64] `query:"page_size"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter fields.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore fields.Field[string] `query:"ending_before"`
}

// URLQuery serializes SubscriptionListParams into a url.Values of the query
// parameters associated with this value
func (r *SubscriptionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r SubscriptionListParams) String() (result string) {
	return fmt.Sprintf("&SubscriptionListParams{PageSize:%s StartingAfter:%s EndingBefore:%s}", r.PageSize, r.StartingAfter, r.EndingBefore)
}

type SubscriptionRecoverParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin fields.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End fields.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes SubscriptionRecoverParams into a url.Values of the query
// parameters associated with this value
func (r *SubscriptionRecoverParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r SubscriptionRecoverParams) String() (result string) {
	return fmt.Sprintf("&SubscriptionRecoverParams{Begin:%s End:%s}", r.Begin, r.End)
}

type SubscriptionReplayMissingParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin fields.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End fields.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes SubscriptionReplayMissingParams into a url.Values of the
// query parameters associated with this value
func (r *SubscriptionReplayMissingParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r SubscriptionReplayMissingParams) String() (result string) {
	return fmt.Sprintf("&SubscriptionReplayMissingParams{Begin:%s End:%s}", r.Begin, r.End)
}
