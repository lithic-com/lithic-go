package lithic

import (
	"context"
	"net/http"
	"net/url"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/field"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

type ResponderEndpointService struct {
	Options []option.RequestOption
}

func NewResponderEndpointService(opts ...option.RequestOption) (r *ResponderEndpointService) {
	r = &ResponderEndpointService{}
	r.Options = opts
	return
}

// Enroll a responder endpoint
func (r *ResponderEndpointService) New(ctx context.Context, body ResponderEndpointNewParams, opts ...option.RequestOption) (res *ResponderEndpointCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "responder_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Disenroll a responder endpoint
func (r *ResponderEndpointService) Delete(ctx context.Context, query ResponderEndpointDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "responder_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, nil, opts...)
	return
}

// Check the status of a responder endpoint
func (r *ResponderEndpointService) CheckStatus(ctx context.Context, query ResponderEndpointCheckStatusParams, opts ...option.RequestOption) (res *ResponderEndpointStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := "responder_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

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

type ResponderEndpointNewParams struct {
	// The URL for the responder endpoint (must be http(s)).
	URL field.Field[string] `json:"url" format:"uri"`
	// The type of the endpoint.
	Type field.Field[ResponderEndpointNewParamsType] `json:"type"`
}

// MarshalJSON serializes ResponderEndpointNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ResponderEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResponderEndpointNewParamsType string

const (
	ResponderEndpointNewParamsTypeTokenizationDecisioning ResponderEndpointNewParamsType = "TOKENIZATION_DECISIONING"
)

type ResponderEndpointDeleteParams struct {
	// The type of the endpoint.
	Type field.Field[ResponderEndpointDeleteParamsType] `query:"type,required"`
}

// URLQuery serializes ResponderEndpointDeleteParams into a url.Values of the query
// parameters associated with this value
func (r ResponderEndpointDeleteParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ResponderEndpointDeleteParamsType string

const (
	ResponderEndpointDeleteParamsTypeTokenizationDecisioning ResponderEndpointDeleteParamsType = "TOKENIZATION_DECISIONING"
)

type ResponderEndpointCheckStatusParams struct {
	// The type of the endpoint.
	Type field.Field[ResponderEndpointCheckStatusParamsType] `query:"type,required"`
}

// URLQuery serializes ResponderEndpointCheckStatusParams into a url.Values of the
// query parameters associated with this value
func (r ResponderEndpointCheckStatusParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ResponderEndpointCheckStatusParamsType string

const (
	ResponderEndpointCheckStatusParamsTypeTokenizationDecisioning ResponderEndpointCheckStatusParamsType = "TOKENIZATION_DECISIONING"
)
