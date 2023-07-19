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

// EventService contains methods and other services that help with interacting with
// the lithic API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewEventService] method instead.
type EventService struct {
	Options       []option.RequestOption
	Subscriptions *EventSubscriptionService
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	r.Subscriptions = NewEventSubscriptionService(opts...)
	return
}

// Get an event.
func (r *EventService) Get(ctx context.Context, eventToken string, opts ...option.RequestOption) (res *Event, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("events/%s", eventToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all events.
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *shared.CursorPage[Event], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events"
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

// List all events.
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[Event] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// List all the message attempts for a given event.
func (r *EventService) ListAttempts(ctx context.Context, eventToken string, query EventListAttemptsParams, opts ...option.RequestOption) (res *shared.CursorPage[MessageAttempt], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("events/%s/attempts", eventToken)
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

// List all the message attempts for a given event.
func (r *EventService) ListAttemptsAutoPaging(ctx context.Context, eventToken string, query EventListAttemptsParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[MessageAttempt] {
	return shared.NewCursorPageAutoPager(r.ListAttempts(ctx, eventToken, query, opts...))
}

// A single event that affects the transaction state and lifecycle.
type Event struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// An RFC 3339 timestamp for when the event was created. UTC time zone.
	//
	// If no timezone is specified, UTC will be used.
	Created time.Time `json:"created,required" format:"date-time"`
	// Event types:
	//
	//   - `card.created` - Notification that a card has been created.
	//   - `card.shipped` - Physical card shipment notification. See
	//     https://docs.lithic.com/docs/cards#physical-card-shipped-webhook.
	//   - `card_transaction.updated` - Transaction Lifecycle webhook. See
	//     https://docs.lithic.com/docs/transaction-webhooks.
	//   - `dispute.updated` - A dispute has been updated.
	//   - `digital_wallet.tokenization_approval_request` - Card network's request to
	//     Lithic to activate a digital wallet token.
	//   - `digital_wallet.tokenization_result` - Notification of the end result of a
	//     tokenization, whether successful or failed.
	//   - `digital_wallet.tokenization_two_factor_authentication_code` - A code to be
	//     passed to an end user to complete digital wallet authentication. See
	//     https://docs.lithic.com/docs/tokenization-control#digital-wallet-tokenization-auth-code.
	EventType EventEventType         `json:"event_type,required"`
	Payload   map[string]interface{} `json:"payload,required"`
	JSON      eventJSON
}

// eventJSON contains the JSON metadata for the struct [Event]
type eventJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	EventType   apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Event types:
//
//   - `card.created` - Notification that a card has been created.
//   - `card.shipped` - Physical card shipment notification. See
//     https://docs.lithic.com/docs/cards#physical-card-shipped-webhook.
//   - `card_transaction.updated` - Transaction Lifecycle webhook. See
//     https://docs.lithic.com/docs/transaction-webhooks.
//   - `dispute.updated` - A dispute has been updated.
//   - `digital_wallet.tokenization_approval_request` - Card network's request to
//     Lithic to activate a digital wallet token.
//   - `digital_wallet.tokenization_result` - Notification of the end result of a
//     tokenization, whether successful or failed.
//   - `digital_wallet.tokenization_two_factor_authentication_code` - A code to be
//     passed to an end user to complete digital wallet authentication. See
//     https://docs.lithic.com/docs/tokenization-control#digital-wallet-tokenization-auth-code.
type EventEventType string

const (
	EventEventTypeCardCreated                                          EventEventType = "card.created"
	EventEventTypeCardShipped                                          EventEventType = "card.shipped"
	EventEventTypeCardTransactionUpdated                               EventEventType = "card_transaction.updated"
	EventEventTypeDigitalWalletTokenizationApprovalRequest             EventEventType = "digital_wallet.tokenization_approval_request"
	EventEventTypeDigitalWalletTokenizationResult                      EventEventType = "digital_wallet.tokenization_result"
	EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode EventEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventEventTypeDisputeUpdated                                       EventEventType = "dispute.updated"
	EventEventTypePaymentTransactionCreated                            EventEventType = "payment_transaction.created"
	EventEventTypePaymentTransactionUpdated                            EventEventType = "payment_transaction.updated"
	EventEventTypeTransferTransactionCreated                           EventEventType = "transfer_transaction.created"
)

// A subscription to specific event types.
type EventSubscription struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// A description of the subscription.
	Description string `json:"description,required"`
	// Whether the subscription is disabled.
	Disabled   bool                         `json:"disabled,required"`
	EventTypes []EventSubscriptionEventType `json:"event_types,required,nullable"`
	URL        string                       `json:"url,required" format:"uri"`
	JSON       eventSubscriptionJSON
}

// eventSubscriptionJSON contains the JSON metadata for the struct
// [EventSubscription]
type eventSubscriptionJSON struct {
	Token       apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	EventTypes  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventSubscriptionEventType string

const (
	EventSubscriptionEventTypeCardCreated                                          EventSubscriptionEventType = "card.created"
	EventSubscriptionEventTypeCardShipped                                          EventSubscriptionEventType = "card.shipped"
	EventSubscriptionEventTypeCardTransactionUpdated                               EventSubscriptionEventType = "card_transaction.updated"
	EventSubscriptionEventTypeDigitalWalletTokenizationApprovalRequest             EventSubscriptionEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionEventTypeDigitalWalletTokenizationResult                      EventSubscriptionEventType = "digital_wallet.tokenization_result"
	EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode EventSubscriptionEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionEventTypeDisputeUpdated                                       EventSubscriptionEventType = "dispute.updated"
	EventSubscriptionEventTypePaymentTransactionCreated                            EventSubscriptionEventType = "payment_transaction.created"
	EventSubscriptionEventTypePaymentTransactionUpdated                            EventSubscriptionEventType = "payment_transaction.updated"
	EventSubscriptionEventTypeTransferTransactionCreated                           EventSubscriptionEventType = "transfer_transaction.created"
)

// A subscription to specific event types.
type MessageAttempt struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// An RFC 3339 timestamp for when the event was created. UTC time zone.
	//
	// If no timezone is specified, UTC will be used.
	Created time.Time `json:"created,required" format:"date-time"`
	// Globally unique identifier.
	EventSubscriptionToken string `json:"event_subscription_token,required"`
	// Globally unique identifier.
	EventToken string `json:"event_token,required"`
	// The response body from the event subscription's URL.
	Response string `json:"response,required"`
	// The response status code from the event subscription's URL.
	ResponseStatusCode int64 `json:"response_status_code,required"`
	// The status of the event attempt.
	Status MessageAttemptStatus `json:"status,required"`
	URL    string               `json:"url,required" format:"uri"`
	JSON   messageAttemptJSON
}

// messageAttemptJSON contains the JSON metadata for the struct [MessageAttempt]
type messageAttemptJSON struct {
	Token                  apijson.Field
	Created                apijson.Field
	EventSubscriptionToken apijson.Field
	EventToken             apijson.Field
	Response               apijson.Field
	ResponseStatusCode     apijson.Field
	Status                 apijson.Field
	URL                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MessageAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the event attempt.
type MessageAttemptStatus string

const (
	MessageAttemptStatusFailed  MessageAttemptStatus = "FAILED"
	MessageAttemptStatusPending MessageAttemptStatus = "PENDING"
	MessageAttemptStatusSending MessageAttemptStatus = "SENDING"
	MessageAttemptStatusSuccess MessageAttemptStatus = "SUCCESS"
)

type EventListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Event types to filter events by.
	EventTypes param.Field[[]EventListParamsEventType] `query:"event_types"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Whether to include the event payload content in the response.
	WithContent param.Field[bool] `query:"with_content"`
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventListParamsEventType string

const (
	EventListParamsEventTypeCardCreated                                          EventListParamsEventType = "card.created"
	EventListParamsEventTypeCardShipped                                          EventListParamsEventType = "card.shipped"
	EventListParamsEventTypeCardTransactionUpdated                               EventListParamsEventType = "card_transaction.updated"
	EventListParamsEventTypeDigitalWalletTokenizationApprovalRequest             EventListParamsEventType = "digital_wallet.tokenization_approval_request"
	EventListParamsEventTypeDigitalWalletTokenizationResult                      EventListParamsEventType = "digital_wallet.tokenization_result"
	EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode EventListParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventListParamsEventTypeDisputeUpdated                                       EventListParamsEventType = "dispute.updated"
	EventListParamsEventTypePaymentTransactionCreated                            EventListParamsEventType = "payment_transaction.created"
	EventListParamsEventTypePaymentTransactionUpdated                            EventListParamsEventType = "payment_transaction.updated"
	EventListParamsEventTypeTransferTransactionCreated                           EventListParamsEventType = "transfer_transaction.created"
)

type EventListAttemptsParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string]                        `query:"starting_after"`
	Status        param.Field[EventListAttemptsParamsStatus] `query:"status"`
}

// URLQuery serializes [EventListAttemptsParams]'s query parameters as
// `url.Values`.
func (r EventListAttemptsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventListAttemptsParamsStatus string

const (
	EventListAttemptsParamsStatusFailed  EventListAttemptsParamsStatus = "FAILED"
	EventListAttemptsParamsStatusPending EventListAttemptsParamsStatus = "PENDING"
	EventListAttemptsParamsStatusSending EventListAttemptsParamsStatus = "SENDING"
	EventListAttemptsParamsStatusSuccess EventListAttemptsParamsStatus = "SUCCESS"
)
