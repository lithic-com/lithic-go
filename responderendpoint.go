// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"net/url"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// ResponderEndpointService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponderEndpointService] method instead.
type ResponderEndpointService struct {
	Options []option.RequestOption
}

// NewResponderEndpointService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewResponderEndpointService(opts ...option.RequestOption) (r *ResponderEndpointService) {
	r = &ResponderEndpointService{}
	r.Options = opts
	return
}

// Enroll a responder endpoint
func (r *ResponderEndpointService) New(ctx context.Context, body ResponderEndpointNewParams, opts ...option.RequestOption) (res *ResponderEndpointNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "responder_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Disenroll a responder endpoint
func (r *ResponderEndpointService) Delete(ctx context.Context, body ResponderEndpointDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "responder_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
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
	URL  string                      `json:"url,nullable" format:"uri"`
	JSON responderEndpointStatusJSON `json:"-"`
}

// responderEndpointStatusJSON contains the JSON metadata for the struct
// [ResponderEndpointStatus]
type responderEndpointStatusJSON struct {
	Enrolled    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponderEndpointStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responderEndpointStatusJSON) RawJSON() string {
	return r.raw
}

type ResponderEndpointNewResponse struct {
	// True if the endpoint was enrolled successfully.
	Enrolled bool                             `json:"enrolled"`
	JSON     responderEndpointNewResponseJSON `json:"-"`
}

// responderEndpointNewResponseJSON contains the JSON metadata for the struct
// [ResponderEndpointNewResponse]
type responderEndpointNewResponseJSON struct {
	Enrolled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponderEndpointNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responderEndpointNewResponseJSON) RawJSON() string {
	return r.raw
}

type ResponderEndpointNewParams struct {
	// The type of the endpoint.
	Type param.Field[ResponderEndpointNewParamsType] `json:"type"`
	// The URL for the responder endpoint (must be http(s)).
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r ResponderEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the endpoint.
type ResponderEndpointNewParamsType string

const (
	ResponderEndpointNewParamsTypeAuthStreamAccess        ResponderEndpointNewParamsType = "AUTH_STREAM_ACCESS"
	ResponderEndpointNewParamsTypeThreeDSDecisioning      ResponderEndpointNewParamsType = "THREE_DS_DECISIONING"
	ResponderEndpointNewParamsTypeTokenizationDecisioning ResponderEndpointNewParamsType = "TOKENIZATION_DECISIONING"
)

func (r ResponderEndpointNewParamsType) IsKnown() bool {
	switch r {
	case ResponderEndpointNewParamsTypeAuthStreamAccess, ResponderEndpointNewParamsTypeThreeDSDecisioning, ResponderEndpointNewParamsTypeTokenizationDecisioning:
		return true
	}
	return false
}

type ResponderEndpointDeleteParams struct {
	// The type of the endpoint.
	Type param.Field[ResponderEndpointDeleteParamsType] `query:"type,required"`
}

// URLQuery serializes [ResponderEndpointDeleteParams]'s query parameters as
// `url.Values`.
func (r ResponderEndpointDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of the endpoint.
type ResponderEndpointDeleteParamsType string

const (
	ResponderEndpointDeleteParamsTypeAuthStreamAccess        ResponderEndpointDeleteParamsType = "AUTH_STREAM_ACCESS"
	ResponderEndpointDeleteParamsTypeThreeDSDecisioning      ResponderEndpointDeleteParamsType = "THREE_DS_DECISIONING"
	ResponderEndpointDeleteParamsTypeTokenizationDecisioning ResponderEndpointDeleteParamsType = "TOKENIZATION_DECISIONING"
)

func (r ResponderEndpointDeleteParamsType) IsKnown() bool {
	switch r {
	case ResponderEndpointDeleteParamsTypeAuthStreamAccess, ResponderEndpointDeleteParamsTypeThreeDSDecisioning, ResponderEndpointDeleteParamsTypeTokenizationDecisioning:
		return true
	}
	return false
}

type ResponderEndpointCheckStatusParams struct {
	// The type of the endpoint.
	Type param.Field[ResponderEndpointCheckStatusParamsType] `query:"type,required"`
}

// URLQuery serializes [ResponderEndpointCheckStatusParams]'s query parameters as
// `url.Values`.
func (r ResponderEndpointCheckStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of the endpoint.
type ResponderEndpointCheckStatusParamsType string

const (
	ResponderEndpointCheckStatusParamsTypeAuthStreamAccess        ResponderEndpointCheckStatusParamsType = "AUTH_STREAM_ACCESS"
	ResponderEndpointCheckStatusParamsTypeThreeDSDecisioning      ResponderEndpointCheckStatusParamsType = "THREE_DS_DECISIONING"
	ResponderEndpointCheckStatusParamsTypeTokenizationDecisioning ResponderEndpointCheckStatusParamsType = "TOKENIZATION_DECISIONING"
)

func (r ResponderEndpointCheckStatusParamsType) IsKnown() bool {
	switch r {
	case ResponderEndpointCheckStatusParamsTypeAuthStreamAccess, ResponderEndpointCheckStatusParamsTypeThreeDSDecisioning, ResponderEndpointCheckStatusParamsTypeTokenizationDecisioning:
		return true
	}
	return false
}
