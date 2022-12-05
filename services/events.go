package services

import (
	"context"
	"fmt"

	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/pagination"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type EventService struct {
	Options       []options.RequestOption
	Subscriptions *EventsSubscriptionService
}

func NewEventService(opts ...options.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	r.Subscriptions = NewEventsSubscriptionService(opts...)
	return
}

// Get an event.
func (r *EventService) Get(ctx context.Context, event_token string, opts ...options.RequestOption) (res *responses.Event, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("events/%s", event_token)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List all events.
func (r *EventService) List(ctx context.Context, query *requests.EventListParams, opts ...options.RequestOption) (res *responses.EventsCursorPage, err error) {
	opts = append(r.Options, opts...)
	path := "events"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.EventsCursorPage{
		CursorPage: &pagination.CursorPage[responses.Event]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
