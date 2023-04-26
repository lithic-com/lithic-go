package responses

import (
	apijson "github.com/lithic-com/lithic-go/core/json"
)

type AuthStreamEnrollment struct {
	// Whether ASA is enrolled.
	Enrolled bool `json:"enrolled"`
	JSON     AuthStreamEnrollmentJSON
}

type AuthStreamEnrollmentJSON struct {
	Enrolled apijson.Metadata
	Raw      []byte
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthStreamEnrollment using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthStreamEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuthStreamSecret struct {
	// The shared HMAC ASA secret
	Secret string `json:"secret"`
	JSON   AuthStreamSecretJSON
}

type AuthStreamSecretJSON struct {
	Secret apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthStreamSecret using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthStreamSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
