// File generated from our OpenAPI spec by Stainless.

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
func (r *EventSubscriptionService) Get(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an event subscription.
func (r *EventSubscriptionService) Update(ctx context.Context, eventSubscriptionToken string, body EventSubscriptionUpdateParams, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", eventSubscriptionToken)
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
func (r *EventSubscriptionService) Delete(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Resend all failed messages since a given time.
func (r *EventSubscriptionService) Recover(ctx context.Context, eventSubscriptionToken string, body EventSubscriptionRecoverParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/recover", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Replays messages to the endpoint. Only messages that were created after `begin`
// will be sent. Messages that were previously sent to the endpoint are not resent.
func (r *EventSubscriptionService) ReplayMissing(ctx context.Context, eventSubscriptionToken string, body EventSubscriptionReplayMissingParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/replay_missing", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get the secret for an event subscription.
func (r *EventSubscriptionService) GetSecret(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (res *EventSubscriptionGetSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Rotate the secret for an event subscription. The previous secret will be valid
// for the next 24 hours.
func (r *EventSubscriptionService) RotateSecret(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("event_subscriptions/%s/secret/rotate", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type EventSubscriptionGetSecretResponse struct {
	Key  string `json:"key"`
	JSON eventSubscriptionGetSecretResponseJSON
}

// eventSubscriptionGetSecretResponseJSON contains the JSON metadata for the struct
// [EventSubscriptionGetSecretResponse]
type eventSubscriptionGetSecretResponseJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventSubscriptionGetSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventSubscriptionNewParams struct {
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Event subscription description.
	Description param.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled param.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes param.Field[[]EventSubscriptionNewParamsEventType] `json:"event_types"`
}

func (r EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventSubscriptionNewParamsEventType string

const (
	EventSubscriptionNewParamsEventTypeCardCreated                                          EventSubscriptionNewParamsEventType = "card.created"
	EventSubscriptionNewParamsEventTypeCardShipped                                          EventSubscriptionNewParamsEventType = "card.shipped"
	EventSubscriptionNewParamsEventTypeCardTransactionUpdated                               EventSubscriptionNewParamsEventType = "card_transaction.updated"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationApprovalRequest             EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationResult                      EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_result"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionNewParamsEventTypeDisputeUpdated                                       EventSubscriptionNewParamsEventType = "dispute.updated"
)

type EventSubscriptionUpdateParams struct {
	// URL to which event webhooks will be sent. URL must be a valid HTTPS address.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Event subscription description.
	Description param.Field[string] `json:"description"`
	// Whether the event subscription is active (false) or inactive (true).
	Disabled param.Field[bool] `json:"disabled"`
	// Indicates types of events that will be sent to this subscription. If left blank,
	// all types will be sent.
	EventTypes param.Field[[]EventSubscriptionUpdateParamsEventType] `json:"event_types"`
}

func (r EventSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventSubscriptionUpdateParamsEventType string

const (
	EventSubscriptionUpdateParamsEventTypeCardCreated                                          EventSubscriptionUpdateParamsEventType = "card.created"
	EventSubscriptionUpdateParamsEventTypeCardShipped                                          EventSubscriptionUpdateParamsEventType = "card.shipped"
	EventSubscriptionUpdateParamsEventTypeCardTransactionUpdated                               EventSubscriptionUpdateParamsEventType = "card_transaction.updated"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationApprovalRequest             EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationResult                      EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_result"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionUpdateParamsEventTypeDisputeUpdated                                       EventSubscriptionUpdateParamsEventType = "dispute.updated"
)

type EventSubscriptionListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [EventSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r EventSubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
