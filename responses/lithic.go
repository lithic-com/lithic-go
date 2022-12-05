package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type APIStatus struct {
	Message string `json:"message"`
	JSON    APIStatusJSON
}

type APIStatusJSON struct {
	Message pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into APIStatus using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *APIStatus) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
