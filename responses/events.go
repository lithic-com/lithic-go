package responses

import (
	"time"

	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/pagination"
)

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
	Token     pjson.Metadata
	EventType pjson.Metadata
	Payload   pjson.Metadata
	Created   pjson.Metadata
	Raw       []byte
	Extras    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Event using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Description pjson.Metadata
	Disabled    pjson.Metadata
	EventTypes  pjson.Metadata
	URL         pjson.Metadata
	Token       pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EventSubscription using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EventSubscriptionEventTypes string

const (
	EventSubscriptionEventTypesDisputeUpdated                           EventSubscriptionEventTypes = "dispute.updated"
	EventSubscriptionEventTypesDigitalWalletTokenizationApprovalRequest EventSubscriptionEventTypes = "digital_wallet.tokenization_approval_request"
)

type EventListResponse struct {
	Data    []Event `json:"data,required"`
	HasMore bool    `json:"has_more,required"`
	JSON    EventListResponseJSON
}

type EventListResponseJSON struct {
	Data    pjson.Metadata
	HasMore pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EventListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EventsCursorPage struct {
	*pagination.CursorPage[Event]
}

func (r *EventsCursorPage) Event() *Event {
	return r.Current()
}

func (r *EventsCursorPage) NextPage() (*EventsCursorPage, error) {
	if page, err := r.CursorPage.NextPage(); err != nil {
		return nil, err
	} else {
		return &EventsCursorPage{page}, nil
	}
}
