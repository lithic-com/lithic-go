package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/pagination"
)

type SubscriptionRetrieveSecretResponse struct {
	Key  string `json:"key"`
	JSON SubscriptionRetrieveSecretResponseJSON
}

type SubscriptionRetrieveSecretResponseJSON struct {
	Key    pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// SubscriptionRetrieveSecretResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SubscriptionRetrieveSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SubscriptionListResponse struct {
	Data    []EventSubscription `json:"data,required"`
	HasMore bool                `json:"has_more,required"`
	JSON    SubscriptionListResponseJSON
}

type SubscriptionListResponseJSON struct {
	Data    pjson.Metadata
	HasMore pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SubscriptionListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EventSubscriptionsCursorPage struct {
	*pagination.CursorPage[EventSubscription]
}

func (r *EventSubscriptionsCursorPage) EventSubscription() *EventSubscription {
	return r.Current()
}

func (r *EventSubscriptionsCursorPage) NextPage() (*EventSubscriptionsCursorPage, error) {
	if page, err := r.CursorPage.NextPage(); err != nil {
		return nil, err
	} else {
		return &EventSubscriptionsCursorPage{page}, nil
	}
}
