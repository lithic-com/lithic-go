package requests

import (
	"fmt"

	"github.com/lithic-com/lithic-go/core/fields"
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type AuthStreamEnrollmentEnrollParams struct {
	// A user-specified url to receive and respond to ASA request.
	WebhookURL fields.Field[string] `json:"webhook_url" format:"uri"`
}

// MarshalJSON serializes AuthStreamEnrollmentEnrollParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *AuthStreamEnrollmentEnrollParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AuthStreamEnrollmentEnrollParams) String() (result string) {
	return fmt.Sprintf("&AuthStreamEnrollmentEnrollParams{WebhookURL:%s}", r.WebhookURL)
}
