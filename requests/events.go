package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/core/fields"
	"github.com/lithic-com/lithic-go/core/query"
)

type EventListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin fields.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End fields.Field[time.Time] `query:"end" format:"date-time"`
	// Page size (for pagination).
	PageSize fields.Field[int64] `query:"page_size"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter fields.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore fields.Field[string]                      `query:"ending_before"`
	EventTypes   fields.Field[[]EventListParamsEventTypes] `query:"event_types[]"`
}

// URLQuery serializes EventListParams into a url.Values of the query parameters
// associated with this value
func (r *EventListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r EventListParams) String() (result string) {
	return fmt.Sprintf("&EventListParams{Begin:%s End:%s PageSize:%s StartingAfter:%s EndingBefore:%s EventTypes:%s}", r.Begin, r.End, r.PageSize, r.StartingAfter, r.EndingBefore, core.Fmt(r.EventTypes))
}

type EventListParamsEventTypes string

const (
	EventListParamsEventTypesDisputeUpdated                           EventListParamsEventTypes = "dispute.updated"
	EventListParamsEventTypesDigitalWalletTokenizationApprovalRequest EventListParamsEventTypes = "digital_wallet.tokenization_approval_request"
)
