package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type EventsSubscriptionService struct {
	Options []option.RequestOption
}

func NewEventsSubscriptionService(opts ...option.RequestOption) (r *EventsSubscriptionService) {
	r = &EventsSubscriptionService{}
	r.Options = opts
	return
}

// Create a new event subscription.
func (r *EventsSubscriptionService) New(ctx context.Context, body *requests.SubscriptionNewParams, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := "event_subscriptions"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get an event subscription.
func (r *EventsSubscriptionService) Get(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an event subscription.
func (r *EventsSubscriptionService) Update(ctx context.Context, event_subscription_token string, body *requests.SubscriptionUpdateParams, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List all the event subscriptions.
func (r *EventsSubscriptionService) List(ctx context.Context, query *requests.SubscriptionListParams, opts ...option.RequestOption) (res *responses.CursorPage[responses.EventSubscription], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "event_subscriptions"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return
	}
	err = cfg.Execute()
	if err != nil {
		return
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all the event subscriptions.
func (r *EventsSubscriptionService) ListAutoPager(ctx context.Context, query *requests.SubscriptionListParams, opts ...option.RequestOption) *responses.CursorPageAutoPager[responses.EventSubscription] {
	return responses.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an event subscription.
func (r *EventsSubscriptionService) Delete(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}

// Resend all failed messages since a given time.
func (r *EventsSubscriptionService) Recover(ctx context.Context, event_subscription_token string, query *requests.SubscriptionRecoverParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/recover", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "POST", path, query, nil, opts...)
	return
}

// Replays messages to the endpoint. Only messages that were created after `begin`
// will be sent. Messages that were previously sent to the endpoint are not resent.
func (r *EventsSubscriptionService) ReplayMissing(ctx context.Context, event_subscription_token string, query *requests.SubscriptionReplayMissingParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/replay_missing", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "POST", path, query, nil, opts...)
	return
}

// Get the secret for an event subscription.
func (r *EventsSubscriptionService) GetSecret(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (res *responses.SubscriptionRetrieveSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Rotate the secret for an event subscription. The previous secret will be valid
// for the next 24 hours.
func (r *EventsSubscriptionService) RotateSecret(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret/rotate", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "POST", path, nil, nil, opts...)
	return
}
