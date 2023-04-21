package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type ResponderEndpointStatus struct {
	// True if the instance has an endpoint enrolled.
	Enrolled bool `json:"enrolled"`
	// The URL of the currently enrolled endpoint or null.
	URL  string `json:"url,nullable" format:"uri"`
	JSON ResponderEndpointStatusJSON
}

type ResponderEndpointStatusJSON struct {
	Enrolled pjson.Metadata
	URL      pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ResponderEndpointStatus using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ResponderEndpointStatus) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ResponderEndpointCreateResponse struct {
	// True if the endpoint was enrolled successfully.
	Enrolled bool `json:"enrolled"`
	JSON     ResponderEndpointCreateResponseJSON
}

type ResponderEndpointCreateResponseJSON struct {
	Enrolled pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ResponderEndpointCreateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ResponderEndpointCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
