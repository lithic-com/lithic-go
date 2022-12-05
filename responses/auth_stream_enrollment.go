package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type AuthStreamEnrollment struct {
	// Whether ASA is enrolled.
	Enrolled bool `json:"enrolled"`
	JSON     AuthStreamEnrollmentJSON
}

type AuthStreamEnrollmentJSON struct {
	Enrolled pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthStreamEnrollment using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthStreamEnrollment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AuthStreamSecret struct {
	// The shared HMAC ASA secret
	Secret string `json:"secret"`
	JSON   AuthStreamSecretJSON
}

type AuthStreamSecretJSON struct {
	Secret pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthStreamSecret using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthStreamSecret) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
