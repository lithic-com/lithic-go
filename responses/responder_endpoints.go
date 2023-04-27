package responses

import (
	apijson "github.com/lithic-com/lithic-go/internal/json"
)

type ResponderEndpointStatus struct {
	// True if the instance has an endpoint enrolled.
	Enrolled bool `json:"enrolled"`
	// The URL of the currently enrolled endpoint or null.
	URL  string `json:"url,nullable" format:"uri"`
	JSON ResponderEndpointStatusJSON
}

type ResponderEndpointStatusJSON struct {
	Enrolled apijson.Metadata
	URL      apijson.Metadata
	raw      string
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ResponderEndpointStatus using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ResponderEndpointStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponderEndpointCreateResponse struct {
	// True if the endpoint was enrolled successfully.
	Enrolled bool `json:"enrolled"`
	JSON     ResponderEndpointCreateResponseJSON
}

type ResponderEndpointCreateResponseJSON struct {
	Enrolled apijson.Metadata
	raw      string
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ResponderEndpointCreateResponse using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ResponderEndpointCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
