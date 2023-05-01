package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/field"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

type EventService struct {
	Options       []option.RequestOption
	Subscriptions *EventSubscriptionService
}

func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	r.Subscriptions = NewEventSubscriptionService(opts...)
	return
}

// Get an event.
func (r *EventService) Get(ctx context.Context, event_token string, opts ...option.RequestOption) (res *Event, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("events/%s", event_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all events.
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *shared.CursorPage[Event], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all events.
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[Event] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type Event struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// Event types:
	//
	//   - `dispute.updated` - A dispute has been updated.
	//   - `digital_wallet.tokenization_approval_request` - Card network's request to
	//     Lithic to activate a digital wallet token.
	EventType EventEventType         `json:"event_type,required"`
	Payload   map[string]interface{} `json:"payload,required"`
	// An RFC 3339 timestamp for when the event was created. UTC time zone.
	//
	// If no timezone is specified, UTC will be used.
	Created time.Time `json:"created,required" format:"date-time"`
	JSON    EventJSON
}

type EventJSON struct {
	Token     apijson.Metadata
	EventType apijson.Metadata
	Payload   apijson.Metadata
	Created   apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Event using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventEventType string

const (
	EventEventTypeDisputeUpdated                           EventEventType = "dispute.updated"
	EventEventTypeDigitalWalletTokenizationApprovalRequest EventEventType = "digital_wallet.tokenization_approval_request"
)

type EventSubscription struct {
	// A description of the subscription.
	Description string `json:"description,required"`
	// Whether the subscription is disabled.
	Disabled   bool                          `json:"disabled,required"`
	EventTypes []EventSubscriptionEventTypes `json:"event_types,required,nullable"`
	URL        string                        `json:"url,required" format:"uri"`
	// Globally unique identifier.
	Token string `json:"token,required"`
	JSON  EventSubscriptionJSON
}

type EventSubscriptionJSON struct {
	Description apijson.Metadata
	Disabled    apijson.Metadata
	EventTypes  apijson.Metadata
	URL         apijson.Metadata
	Token       apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EventSubscription using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventSubscriptionEventTypes string

const (
	EventSubscriptionEventTypesDisputeUpdated                           EventSubscriptionEventTypes = "dispute.updated"
	EventSubscriptionEventTypesDigitalWalletTokenizationApprovalRequest EventSubscriptionEventTypes = "digital_wallet.tokenization_approval_request"
)

type EventListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin field.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End field.Field[time.Time] `query:"end" format:"date-time"`
	// Page size (for pagination).
	PageSize field.Field[int64] `query:"page_size"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter field.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore field.Field[string]                      `query:"ending_before"`
	EventTypes   field.Field[[]EventListParamsEventTypes] `query:"event_types[]"`
}

// URLQuery serializes EventListParams into a url.Values of the query parameters
// associated with this value
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type EventListParamsEventTypes string

const (
	EventListParamsEventTypesDisputeUpdated                           EventListParamsEventTypes = "dispute.updated"
	EventListParamsEventTypesDigitalWalletTokenizationApprovalRequest EventListParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type EventListResponse struct {
	Data    []Event `json:"data,required"`
	HasMore bool    `json:"has_more,required"`
	JSON    EventListResponseJSON
}

type EventListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	raw     string
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EventListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
