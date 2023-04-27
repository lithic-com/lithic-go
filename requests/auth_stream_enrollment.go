package requests

import (
	"github.com/lithic-com/lithic-go/internal/field"
	apijson "github.com/lithic-com/lithic-go/internal/json"
)

type AuthStreamEnrollmentEnrollParams struct {
	// A user-specified url to receive and respond to ASA request.
	WebhookURL field.Field[string] `json:"webhook_url" format:"uri"`
}

// MarshalJSON serializes AuthStreamEnrollmentEnrollParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r AuthStreamEnrollmentEnrollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
