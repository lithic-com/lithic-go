package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type EventSubscriptionService struct {
	Options []option.RequestOption
}

func NewEventSubscriptionService(opts ...option.RequestOption) (r *EventSubscriptionService) {
	r = &EventSubscriptionService{}
	r.Options = opts
	return
}

// Create a new event subscription.
func (r *EventSubscriptionService) New(ctx context.Context, body *requests.EventSubscriptionNewParams, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := "event_subscriptions"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get an event subscription.
func (r *EventSubscriptionService) Get(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an event subscription.
func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_token string, body *requests.EventSubscriptionUpdateParams, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List all the event subscriptions.
func (r *EventSubscriptionService) List(ctx context.Context, query *requests.EventSubscriptionListParams, opts ...option.RequestOption) (res *responses.CursorPage[responses.EventSubscription], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "event_subscriptions"
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

// List all the event subscriptions.
func (r *EventSubscriptionService) ListAutoPager(ctx context.Context, query *requests.EventSubscriptionListParams, opts ...option.RequestOption) *responses.CursorPageAutoPager[responses.EventSubscription] {
	return responses.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an event subscription.
func (r *EventSubscriptionService) Delete(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}

// Resend all failed messages since a given time.
func (r *EventSubscriptionService) Recover(ctx context.Context, event_subscription_token string, query *requests.EventSubscriptionRecoverParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/recover", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "POST", path, query, nil, opts...)
	return
}

// Replays messages to the endpoint. Only messages that were created after `begin`
// will be sent. Messages that were previously sent to the endpoint are not resent.
func (r *EventSubscriptionService) ReplayMissing(ctx context.Context, event_subscription_token string, query *requests.EventSubscriptionReplayMissingParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/replay_missing", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "POST", path, query, nil, opts...)
	return
}

// Get the secret for an event subscription.
func (r *EventSubscriptionService) GetSecret(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (res *responses.SubscriptionRetrieveSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Rotate the secret for an event subscription. The previous secret will be valid
// for the next 24 hours.
func (r *EventSubscriptionService) RotateSecret(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret/rotate", event_subscription_token)
	err = option.ExecuteNewRequest(ctx, "POST", path, nil, nil, opts...)
	return
}
