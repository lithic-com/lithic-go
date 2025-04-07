// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// EventEventSubscriptionService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventEventSubscriptionService] method instead.
type EventEventSubscriptionService struct {
	Options []option.RequestOption
}

// NewEventEventSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEventEventSubscriptionService(opts ...option.RequestOption) (r *EventEventSubscriptionService) {
	r = &EventEventSubscriptionService{}
	r.Options = opts
	return
}

// Resend an event to an event subscription.
func (r *EventEventSubscriptionService) Resend(ctx context.Context, eventToken string, eventSubscriptionToken string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if eventToken == "" {
		err = errors.New("missing required event_token parameter")
		return
	}
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/events/%s/event_subscriptions/%s/resend", eventToken, eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}
