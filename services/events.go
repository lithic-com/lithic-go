package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type EventService struct {
	Options       []option.RequestOption
	Subscriptions *EventsSubscriptionService
}

func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	r.Subscriptions = NewEventsSubscriptionService(opts...)
	return
}

// Get an event.
func (r *EventService) Get(ctx context.Context, event_token string, opts ...option.RequestOption) (res *responses.Event, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("events/%s", event_token)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List all events.
func (r *EventService) List(ctx context.Context, query *requests.EventListParams, opts ...option.RequestOption) (res *responses.CursorPage[responses.Event], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all events.
func (r *EventService) ListAutoPager(ctx context.Context, query *requests.EventListParams, opts ...option.RequestOption) *responses.CursorPageAutoPager[responses.Event] {
	return responses.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}
