package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// EventService contains methods and other services that help with interacting with
// the lithic API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewEventService] method instead.
type EventService struct {
	Options       []option.RequestOption
	Subscriptions *EventSubscriptionService
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
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

// A single event that affects the transaction state and lifecycle.
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
	JSON    eventJSON
}

// eventJSON contains the JSON metadata for the struct [Event]
type eventJSON struct {
	Token       apijson.Field
	EventType   apijson.Field
	Payload     apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventEventType string

const (
	EventEventTypeDisputeUpdated                           EventEventType = "dispute.updated"
	EventEventTypeDigitalWalletTokenizationApprovalRequest EventEventType = "digital_wallet.tokenization_approval_request"
)

// A subscription to specific event types.
type EventSubscription struct {
	// A description of the subscription.
	Description string `json:"description,required"`
	// Whether the subscription is disabled.
	Disabled   bool                          `json:"disabled,required"`
	EventTypes []EventSubscriptionEventTypes `json:"event_types,required,nullable"`
	URL        string                        `json:"url,required" format:"uri"`
	// Globally unique identifier.
	Token string `json:"token,required"`
	JSON  eventSubscriptionJSON
}

// eventSubscriptionJSON contains the JSON metadata for the struct
// [EventSubscription]
type eventSubscriptionJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	EventTypes  apijson.Field
	URL         apijson.Field
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

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
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string]                      `query:"ending_before"`
	EventTypes   param.Field[[]EventListParamsEventTypes] `query:"event_types"`
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventListParamsEventTypes string

const (
	EventListParamsEventTypesDisputeUpdated                           EventListParamsEventTypes = "dispute.updated"
	EventListParamsEventTypesDigitalWalletTokenizationApprovalRequest EventListParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type EventListResponse struct {
	Data    []Event `json:"data,required"`
	HasMore bool    `json:"has_more,required"`
	JSON    eventListResponseJSON
}

// eventListResponseJSON contains the JSON metadata for the struct
// [EventListResponse]
type eventListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
