package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// EventSubscriptionService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEventSubscriptionService] method
// instead.
type EventSubscriptionService struct {
	Options []option.RequestOption
}

// NewEventSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEventSubscriptionService(opts ...option.RequestOption) (r *EventSubscriptionService) {
	r = &EventSubscriptionService{}
	r.Options = opts
	return
}

// Create a new event subscription.
func (r *EventSubscriptionService) New(ctx context.Context, body EventSubscriptionNewParams, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := "event_subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an event subscription.
func (r *EventSubscriptionService) Get(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an event subscription.
func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_token string, body EventSubscriptionUpdateParams, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all the event subscriptions.
func (r *EventSubscriptionService) List(ctx context.Context, query EventSubscriptionListParams, opts ...option.RequestOption) (res *shared.CursorPage[EventSubscription], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "event_subscriptions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *EventSubscriptionService) ListAutoPaging(ctx context.Context, query EventSubscriptionListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[EventSubscription] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an event subscription.
func (r *EventSubscriptionService) Delete(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Resend all failed messages since a given time.
func (r *EventSubscriptionService) Recover(ctx context.Context, event_subscription_token string, query EventSubscriptionRecoverParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/recover", event_subscription_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, nil, opts...)
	return
}

// Replays messages to the endpoint. Only messages that were created after `begin`
// will be sent. Messages that were previously sent to the endpoint are not resent.
func (r *EventSubscriptionService) ReplayMissing(ctx context.Context, event_subscription_token string, query EventSubscriptionReplayMissingParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/replay_missing", event_subscription_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, nil, opts...)
	return
}

// Get the secret for an event subscription.
func (r *EventSubscriptionService) GetSecret(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (res *SubscriptionRetrieveSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret", event_subscription_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Rotate the secret for an event subscription. The previous secret will be valid
// for the next 24 hours.
func (r *EventSubscriptionService) RotateSecret(ctx context.Context, event_subscription_token string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret/rotate", event_subscription_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type SubscriptionRetrieveSecretResponse struct {
	Key  string `json:"key"`
	JSON subscriptionRetrieveSecretResponseJSON
}

// subscriptionRetrieveSecretResponseJSON contains the JSON metadata for the struct
// [SubscriptionRetrieveSecretResponse]
type subscriptionRetrieveSecretResponseJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionRetrieveSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventSubscriptionNewParams struct {
	// Event subscription description.
	Description param.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled param.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes param.Field[[]EventSubscriptionNewParamsEventTypes] `json:"event_types"`
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventSubscriptionNewParamsEventTypes string

const (
	EventSubscriptionNewParamsEventTypesDisputeUpdated                           EventSubscriptionNewParamsEventTypes = "dispute.updated"
	EventSubscriptionNewParamsEventTypesDigitalWalletTokenizationApprovalRequest EventSubscriptionNewParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type EventSubscriptionUpdateParams struct {
	// Event subscription description.
	Description param.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled param.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes param.Field[[]EventSubscriptionUpdateParamsEventTypes] `json:"event_types"`
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r EventSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventSubscriptionUpdateParamsEventTypes string

const (
	EventSubscriptionUpdateParamsEventTypesDisputeUpdated                           EventSubscriptionUpdateParamsEventTypes = "dispute.updated"
	EventSubscriptionUpdateParamsEventTypesDigitalWalletTokenizationApprovalRequest EventSubscriptionUpdateParamsEventTypes = "digital_wallet.tokenization_approval_request"
)

type EventSubscriptionListParams struct {
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// The unique identifier of the last item in the previous page. Used to retrieve
	// the next page.
	StartingAfter param.Field[string] `query:"starting_after"`
	// The unique identifier of the first item in the previous page. Used to retrieve
	// the previous page.
	EndingBefore param.Field[string] `query:"ending_before"`
}

// URLQuery serializes [EventSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r EventSubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type EventSubscriptionListResponse struct {
	Data    []EventSubscription `json:"data,required"`
	HasMore bool                `json:"has_more,required"`
	JSON    eventSubscriptionListResponseJSON
}

// eventSubscriptionListResponseJSON contains the JSON metadata for the struct
// [EventSubscriptionListResponse]
type eventSubscriptionListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventSubscriptionRecoverParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes [EventSubscriptionRecoverParams]'s query parameters as
// `url.Values`.
func (r EventSubscriptionRecoverParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type EventSubscriptionReplayMissingParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes [EventSubscriptionReplayMissingParams]'s query parameters as
// `url.Values`.
func (r EventSubscriptionReplayMissingParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
