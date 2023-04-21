package requests

import (
	"net/url"

	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

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
	return pjson.MarshalRoot(r)
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
	return query.Marshal(r)
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
	return query.Marshal(r)
}

type ResponderEndpointCheckStatusParamsType string

const (
	ResponderEndpointCheckStatusParamsTypeTokenizationDecisioning ResponderEndpointCheckStatusParamsType = "TOKENIZATION_DECISIONING"
)
