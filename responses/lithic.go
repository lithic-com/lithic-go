package responses

import (
	apijson "github.com/lithic-com/lithic-go/internal/json"
)

type APIStatus struct {
	Message string `json:"message"`
	JSON    APIStatusJSON
}

type APIStatusJSON struct {
	Message apijson.Metadata
	raw     string
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into APIStatus using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *APIStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
