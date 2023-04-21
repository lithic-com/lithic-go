package services

import (
	"context"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
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
func (r *ResponderEndpointService) New(ctx context.Context, body *requests.ResponderEndpointNewParams, opts ...option.RequestOption) (res *responses.ResponderEndpointCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "responder_endpoints"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Disenroll a responder endpoint
func (r *ResponderEndpointService) Delete(ctx context.Context, query *requests.ResponderEndpointDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "responder_endpoints"
	err = option.ExecuteNewRequest(ctx, "DELETE", path, query, nil, opts...)
	return
}

// Check the status of a responder endpoint
func (r *ResponderEndpointService) CheckStatus(ctx context.Context, query *requests.ResponderEndpointCheckStatusParams, opts ...option.RequestOption) (res *responses.ResponderEndpointStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := "responder_endpoints"
	err = option.ExecuteNewRequest(ctx, "GET", path, query, &res, opts...)
	return
}
