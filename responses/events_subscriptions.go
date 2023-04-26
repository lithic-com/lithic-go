package responses

import (
	apijson "github.com/lithic-com/lithic-go/core/json"
)

type SubscriptionRetrieveSecretResponse struct {
	Key  string `json:"key"`
	JSON SubscriptionRetrieveSecretResponseJSON
}

type SubscriptionRetrieveSecretResponseJSON struct {
	Key    apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// SubscriptionRetrieveSecretResponse using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *SubscriptionRetrieveSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventSubscriptionListResponse struct {
	Data    []EventSubscription `json:"data,required"`
	HasMore bool                `json:"has_more,required"`
	JSON    EventSubscriptionListResponseJSON
}

type EventSubscriptionListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	Raw     []byte
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EventSubscriptionListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EventSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
