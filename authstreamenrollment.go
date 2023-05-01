package lithic

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/field"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

type AuthStreamEnrollmentService struct {
	Options []option.RequestOption
}

func NewAuthStreamEnrollmentService(opts ...option.RequestOption) (r *AuthStreamEnrollmentService) {
	r = &AuthStreamEnrollmentService{}
	r.Options = opts
	return
}

// Check status for whether you have enrolled in Authorization Stream Access (ASA)
// for your program in Sandbox.
func (r *AuthStreamEnrollmentService) Get(ctx context.Context, opts ...option.RequestOption) (res *AuthStreamEnrollment, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth_stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disenroll Authorization Stream Access (ASA) in Sandbox.
func (r *AuthStreamEnrollmentService) Disenroll(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "auth_stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Authorization Stream Access (ASA) provides the ability to make custom
// transaction approval decisions through an HTTP interface to the ISO 8583 message
// stream.
//
// ASA requests are delivered as an HTTP POST during authorization. The ASA request
// body adheres to the Lithic transaction schema, with some additional fields added
// for use in decisioning. A response should be sent with HTTP response code 200
// and the approval decision in the response body. This response is converted by
// Lithic back into ISO 8583 format and forwarded to the network.
//
// In Sandbox, users can self-enroll and disenroll in ASA. In production,
// onboarding requires manual approval and setup.
func (r *AuthStreamEnrollmentService) Enroll(ctx context.Context, body AuthStreamEnrollmentEnrollParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "auth_stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve the ASA HMAC secret key. If one does not exist your program yet,
// calling this endpoint will create one for you. The headers (which you can use to
// verify webhooks) will begin appearing shortly after calling this endpoint for
// the first time. See
// [this page](https://docs.lithic.com/docs/auth-stream-access-asa#asa-webhook-verification)
// for more detail about verifying ASA webhooks.
func (r *AuthStreamEnrollmentService) GetSecret(ctx context.Context, opts ...option.RequestOption) (res *AuthStreamSecret, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth_stream/secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate a new ASA HMAC secret key. The old ASA HMAC secret key will be
// deactivated 24 hours after a successful request to this endpoint. Make a
// [`GET /auth_stream/secret`](https://docs.lithic.com/reference/getauthstreamsecret)
// request to retrieve the new secret key.
func (r *AuthStreamEnrollmentService) RotateSecret(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "auth_stream/secret/rotate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type AuthStreamEnrollment struct {
	// Whether ASA is enrolled.
	Enrolled bool `json:"enrolled"`
	JSON     AuthStreamEnrollmentJSON
}

type AuthStreamEnrollmentJSON struct {
	Enrolled apijson.Metadata
	raw      string
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
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AuthStreamSecret using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AuthStreamSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
