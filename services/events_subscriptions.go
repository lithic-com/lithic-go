package services

import (
	"context"
	"fmt"

	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/pagination"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type EventsSubscriptionService struct {
	Options []options.RequestOption
}

func NewEventsSubscriptionService(opts ...options.RequestOption) (r *EventsSubscriptionService) {
	r = &EventsSubscriptionService{}
	r.Options = opts
	return
}

// Create a new event subscription.
func (r *EventsSubscriptionService) New(ctx context.Context, body *requests.SubscriptionNewParams, opts ...options.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := "event_subscriptions"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get an event subscription.
func (r *EventsSubscriptionService) Get(ctx context.Context, event_subscription_token string, opts ...options.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an event subscription.
func (r *EventsSubscriptionService) Update(ctx context.Context, event_subscription_token string, body *requests.SubscriptionUpdateParams, opts ...options.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List all the event subscriptions.
func (r *EventsSubscriptionService) List(ctx context.Context, query *requests.SubscriptionListParams, opts ...options.RequestOption) (res *responses.EventSubscriptionsCursorPage, err error) {
	opts = append(r.Options, opts...)
	path := "event_subscriptions"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.EventSubscriptionsCursorPage{
		CursorPage: &pagination.CursorPage[responses.EventSubscription]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Delete an event subscription.
func (r *EventsSubscriptionService) Delete(ctx context.Context, event_subscription_token string, opts ...options.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = options.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}

// Resend all failed messages since a given time.
func (r *EventsSubscriptionService) Recover(ctx context.Context, event_subscription_token string, query *requests.SubscriptionRecoverParams, opts ...options.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/recover", event_subscription_token)
	err = options.ExecuteNewRequest(ctx, "POST", path, query, nil, opts...)
	return
}

// Replays messages to the endpoint. Only messages that were created after `begin`
// will be sent. Messages that were previously sent to the endpoint are not resent.
func (r *EventsSubscriptionService) ReplayMissing(ctx context.Context, event_subscription_token string, query *requests.SubscriptionReplayMissingParams, opts ...options.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/replay_missing", event_subscription_token)
	err = options.ExecuteNewRequest(ctx, "POST", path, query, nil, opts...)
	return
}

// Get the secret for an event subscription.
func (r *EventsSubscriptionService) GetSecret(ctx context.Context, event_subscription_token string, opts ...options.RequestOption) (res *responses.SubscriptionRetrieveSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret", event_subscription_token)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Rotate the secret for an event subscription. The previous secret will be valid
// for the next 24 hours.
func (r *EventsSubscriptionService) RotateSecret(ctx context.Context, event_subscription_token string, opts ...options.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret/rotate", event_subscription_token)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, nil, opts...)
	return
}
