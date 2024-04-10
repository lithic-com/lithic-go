// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
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
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Event], err error) {
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
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Event] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// List all the message attempts for a given event.
func (r *EventService) ListAttempts(ctx context.Context, eventToken string, query EventListAttemptsParams, opts ...option.RequestOption) (res *pagination.CursorPage[MessageAttempt], err error) {
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
func (r *EventService) ListAttemptsAutoPaging(ctx context.Context, eventToken string, query EventListAttemptsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[MessageAttempt] {
	return pagination.NewCursorPageAutoPager(r.ListAttempts(ctx, eventToken, query, opts...))
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
	//   - `account_holder.created` - Notification that a new account holder has been
	//     created and was not rejected.
	//   - `account_holder.updated` - Notification that an account holder was updated.
	//   - `account_holder.verification` - Notification than an account holder's identity
	//     verification is complete.
	//   - `card.created` - Notification that a card has been created.
	//   - `card.renewed` - Notification that a card has been renewed.
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
	//   - `digital_wallet.tokenization_two_factor_authentication_code_sent` -
	//     Notification that a two factor authentication code for activating a digital
	//     wallet has been sent to the end user.
	//   - `digital_wallet.tokenization_updated` - Notification that a digital wallet
	//     tokenization's status has changed.
	EventType EventEventType         `json:"event_type,required"`
	Payload   map[string]interface{} `json:"payload,required"`
	JSON      eventJSON              `json:"-"`
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

func (r eventJSON) RawJSON() string {
	return r.raw
}

// Event types:
//
//   - `account_holder.created` - Notification that a new account holder has been
//     created and was not rejected.
//   - `account_holder.updated` - Notification that an account holder was updated.
//   - `account_holder.verification` - Notification than an account holder's identity
//     verification is complete.
//   - `card.created` - Notification that a card has been created.
//   - `card.renewed` - Notification that a card has been renewed.
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
//   - `digital_wallet.tokenization_two_factor_authentication_code_sent` -
//     Notification that a two factor authentication code for activating a digital
//     wallet has been sent to the end user.
//   - `digital_wallet.tokenization_updated` - Notification that a digital wallet
//     tokenization's status has changed.
type EventEventType string

const (
	EventEventTypeAccountHolderCreated                                     EventEventType = "account_holder.created"
	EventEventTypeAccountHolderUpdated                                     EventEventType = "account_holder.updated"
	EventEventTypeAccountHolderVerification                                EventEventType = "account_holder.verification"
	EventEventTypeBalanceUpdated                                           EventEventType = "balance.updated"
	EventEventTypeCardCreated                                              EventEventType = "card.created"
	EventEventTypeCardRenewed                                              EventEventType = "card.renewed"
	EventEventTypeCardShipped                                              EventEventType = "card.shipped"
	EventEventTypeCardTransactionUpdated                                   EventEventType = "card_transaction.updated"
	EventEventTypeDigitalWalletTokenizationApprovalRequest                 EventEventType = "digital_wallet.tokenization_approval_request"
	EventEventTypeDigitalWalletTokenizationResult                          EventEventType = "digital_wallet.tokenization_result"
	EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventEventTypeDigitalWalletTokenizationUpdated                         EventEventType = "digital_wallet.tokenization_updated"
	EventEventTypeDisputeUpdated                                           EventEventType = "dispute.updated"
	EventEventTypeDisputeEvidenceUploadFailed                              EventEventType = "dispute_evidence.upload_failed"
	EventEventTypeExternalBankAccountCreated                               EventEventType = "external_bank_account.created"
	EventEventTypeExternalBankAccountUpdated                               EventEventType = "external_bank_account.updated"
	EventEventTypeFinancialAccountCreated                                  EventEventType = "financial_account.created"
	EventEventTypePaymentTransactionCreated                                EventEventType = "payment_transaction.created"
	EventEventTypePaymentTransactionUpdated                                EventEventType = "payment_transaction.updated"
	EventEventTypeSettlementReportUpdated                                  EventEventType = "settlement_report.updated"
	EventEventTypeStatementsCreated                                        EventEventType = "statements.created"
	EventEventTypeThreeDSAuthenticationCreated                             EventEventType = "three_ds_authentication.created"
	EventEventTypeTransferTransactionCreated                               EventEventType = "transfer_transaction.created"
)

func (r EventEventType) IsKnown() bool {
	switch r {
	case EventEventTypeAccountHolderCreated, EventEventTypeAccountHolderUpdated, EventEventTypeAccountHolderVerification, EventEventTypeBalanceUpdated, EventEventTypeCardCreated, EventEventTypeCardRenewed, EventEventTypeCardShipped, EventEventTypeCardTransactionUpdated, EventEventTypeDigitalWalletTokenizationApprovalRequest, EventEventTypeDigitalWalletTokenizationResult, EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventEventTypeDigitalWalletTokenizationUpdated, EventEventTypeDisputeUpdated, EventEventTypeDisputeEvidenceUploadFailed, EventEventTypeExternalBankAccountCreated, EventEventTypeExternalBankAccountUpdated, EventEventTypeFinancialAccountCreated, EventEventTypePaymentTransactionCreated, EventEventTypePaymentTransactionUpdated, EventEventTypeSettlementReportUpdated, EventEventTypeStatementsCreated, EventEventTypeThreeDSAuthenticationCreated, EventEventTypeTransferTransactionCreated:
		return true
	}
	return false
}

// A subscription to specific event types.
type EventSubscription struct {
	// Globally unique identifier.
	Token string `json:"token,required"`
	// A description of the subscription.
	Description string `json:"description,required"`
	// Whether the subscription is disabled.
	Disabled   bool                         `json:"disabled,required"`
	URL        string                       `json:"url,required" format:"uri"`
	EventTypes []EventSubscriptionEventType `json:"event_types,nullable"`
	JSON       eventSubscriptionJSON        `json:"-"`
}

// eventSubscriptionJSON contains the JSON metadata for the struct
// [EventSubscription]
type eventSubscriptionJSON struct {
	Token       apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	URL         apijson.Field
	EventTypes  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventSubscriptionJSON) RawJSON() string {
	return r.raw
}

type EventSubscriptionEventType string

const (
	EventSubscriptionEventTypeAccountHolderCreated                                     EventSubscriptionEventType = "account_holder.created"
	EventSubscriptionEventTypeAccountHolderUpdated                                     EventSubscriptionEventType = "account_holder.updated"
	EventSubscriptionEventTypeAccountHolderVerification                                EventSubscriptionEventType = "account_holder.verification"
	EventSubscriptionEventTypeBalanceUpdated                                           EventSubscriptionEventType = "balance.updated"
	EventSubscriptionEventTypeCardCreated                                              EventSubscriptionEventType = "card.created"
	EventSubscriptionEventTypeCardRenewed                                              EventSubscriptionEventType = "card.renewed"
	EventSubscriptionEventTypeCardShipped                                              EventSubscriptionEventType = "card.shipped"
	EventSubscriptionEventTypeCardTransactionUpdated                                   EventSubscriptionEventType = "card_transaction.updated"
	EventSubscriptionEventTypeDigitalWalletTokenizationApprovalRequest                 EventSubscriptionEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionEventTypeDigitalWalletTokenizationResult                          EventSubscriptionEventType = "digital_wallet.tokenization_result"
	EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventSubscriptionEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventSubscriptionEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventSubscriptionEventTypeDigitalWalletTokenizationUpdated                         EventSubscriptionEventType = "digital_wallet.tokenization_updated"
	EventSubscriptionEventTypeDisputeUpdated                                           EventSubscriptionEventType = "dispute.updated"
	EventSubscriptionEventTypeDisputeEvidenceUploadFailed                              EventSubscriptionEventType = "dispute_evidence.upload_failed"
	EventSubscriptionEventTypeExternalBankAccountCreated                               EventSubscriptionEventType = "external_bank_account.created"
	EventSubscriptionEventTypeExternalBankAccountUpdated                               EventSubscriptionEventType = "external_bank_account.updated"
	EventSubscriptionEventTypeFinancialAccountCreated                                  EventSubscriptionEventType = "financial_account.created"
	EventSubscriptionEventTypePaymentTransactionCreated                                EventSubscriptionEventType = "payment_transaction.created"
	EventSubscriptionEventTypePaymentTransactionUpdated                                EventSubscriptionEventType = "payment_transaction.updated"
	EventSubscriptionEventTypeSettlementReportUpdated                                  EventSubscriptionEventType = "settlement_report.updated"
	EventSubscriptionEventTypeStatementsCreated                                        EventSubscriptionEventType = "statements.created"
	EventSubscriptionEventTypeThreeDSAuthenticationCreated                             EventSubscriptionEventType = "three_ds_authentication.created"
	EventSubscriptionEventTypeTransferTransactionCreated                               EventSubscriptionEventType = "transfer_transaction.created"
)

func (r EventSubscriptionEventType) IsKnown() bool {
	switch r {
	case EventSubscriptionEventTypeAccountHolderCreated, EventSubscriptionEventTypeAccountHolderUpdated, EventSubscriptionEventTypeAccountHolderVerification, EventSubscriptionEventTypeBalanceUpdated, EventSubscriptionEventTypeCardCreated, EventSubscriptionEventTypeCardRenewed, EventSubscriptionEventTypeCardShipped, EventSubscriptionEventTypeCardTransactionUpdated, EventSubscriptionEventTypeDigitalWalletTokenizationApprovalRequest, EventSubscriptionEventTypeDigitalWalletTokenizationResult, EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionEventTypeDigitalWalletTokenizationUpdated, EventSubscriptionEventTypeDisputeUpdated, EventSubscriptionEventTypeDisputeEvidenceUploadFailed, EventSubscriptionEventTypeExternalBankAccountCreated, EventSubscriptionEventTypeExternalBankAccountUpdated, EventSubscriptionEventTypeFinancialAccountCreated, EventSubscriptionEventTypePaymentTransactionCreated, EventSubscriptionEventTypePaymentTransactionUpdated, EventSubscriptionEventTypeSettlementReportUpdated, EventSubscriptionEventTypeStatementsCreated, EventSubscriptionEventTypeThreeDSAuthenticationCreated, EventSubscriptionEventTypeTransferTransactionCreated:
		return true
	}
	return false
}

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
	JSON   messageAttemptJSON   `json:"-"`
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

func (r messageAttemptJSON) RawJSON() string {
	return r.raw
}

// The status of the event attempt.
type MessageAttemptStatus string

const (
	MessageAttemptStatusFailed  MessageAttemptStatus = "FAILED"
	MessageAttemptStatusPending MessageAttemptStatus = "PENDING"
	MessageAttemptStatusSending MessageAttemptStatus = "SENDING"
	MessageAttemptStatusSuccess MessageAttemptStatus = "SUCCESS"
)

func (r MessageAttemptStatus) IsKnown() bool {
	switch r {
	case MessageAttemptStatusFailed, MessageAttemptStatusPending, MessageAttemptStatusSending, MessageAttemptStatusSuccess:
		return true
	}
	return false
}

type EventListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
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
	EventListParamsEventTypeAccountHolderCreated                                     EventListParamsEventType = "account_holder.created"
	EventListParamsEventTypeAccountHolderUpdated                                     EventListParamsEventType = "account_holder.updated"
	EventListParamsEventTypeAccountHolderVerification                                EventListParamsEventType = "account_holder.verification"
	EventListParamsEventTypeBalanceUpdated                                           EventListParamsEventType = "balance.updated"
	EventListParamsEventTypeCardCreated                                              EventListParamsEventType = "card.created"
	EventListParamsEventTypeCardRenewed                                              EventListParamsEventType = "card.renewed"
	EventListParamsEventTypeCardShipped                                              EventListParamsEventType = "card.shipped"
	EventListParamsEventTypeCardTransactionUpdated                                   EventListParamsEventType = "card_transaction.updated"
	EventListParamsEventTypeDigitalWalletTokenizationApprovalRequest                 EventListParamsEventType = "digital_wallet.tokenization_approval_request"
	EventListParamsEventTypeDigitalWalletTokenizationResult                          EventListParamsEventType = "digital_wallet.tokenization_result"
	EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventListParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventListParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventListParamsEventTypeDigitalWalletTokenizationUpdated                         EventListParamsEventType = "digital_wallet.tokenization_updated"
	EventListParamsEventTypeDisputeUpdated                                           EventListParamsEventType = "dispute.updated"
	EventListParamsEventTypeDisputeEvidenceUploadFailed                              EventListParamsEventType = "dispute_evidence.upload_failed"
	EventListParamsEventTypeExternalBankAccountCreated                               EventListParamsEventType = "external_bank_account.created"
	EventListParamsEventTypeExternalBankAccountUpdated                               EventListParamsEventType = "external_bank_account.updated"
	EventListParamsEventTypeFinancialAccountCreated                                  EventListParamsEventType = "financial_account.created"
	EventListParamsEventTypePaymentTransactionCreated                                EventListParamsEventType = "payment_transaction.created"
	EventListParamsEventTypePaymentTransactionUpdated                                EventListParamsEventType = "payment_transaction.updated"
	EventListParamsEventTypeSettlementReportUpdated                                  EventListParamsEventType = "settlement_report.updated"
	EventListParamsEventTypeStatementsCreated                                        EventListParamsEventType = "statements.created"
	EventListParamsEventTypeThreeDSAuthenticationCreated                             EventListParamsEventType = "three_ds_authentication.created"
	EventListParamsEventTypeTransferTransactionCreated                               EventListParamsEventType = "transfer_transaction.created"
)

func (r EventListParamsEventType) IsKnown() bool {
	switch r {
	case EventListParamsEventTypeAccountHolderCreated, EventListParamsEventTypeAccountHolderUpdated, EventListParamsEventTypeAccountHolderVerification, EventListParamsEventTypeBalanceUpdated, EventListParamsEventTypeCardCreated, EventListParamsEventTypeCardRenewed, EventListParamsEventTypeCardShipped, EventListParamsEventTypeCardTransactionUpdated, EventListParamsEventTypeDigitalWalletTokenizationApprovalRequest, EventListParamsEventTypeDigitalWalletTokenizationResult, EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventListParamsEventTypeDigitalWalletTokenizationUpdated, EventListParamsEventTypeDisputeUpdated, EventListParamsEventTypeDisputeEvidenceUploadFailed, EventListParamsEventTypeExternalBankAccountCreated, EventListParamsEventTypeExternalBankAccountUpdated, EventListParamsEventTypeFinancialAccountCreated, EventListParamsEventTypePaymentTransactionCreated, EventListParamsEventTypePaymentTransactionUpdated, EventListParamsEventTypeSettlementReportUpdated, EventListParamsEventTypeStatementsCreated, EventListParamsEventTypeThreeDSAuthenticationCreated, EventListParamsEventTypeTransferTransactionCreated:
		return true
	}
	return false
}

type EventListAttemptsParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
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

func (r EventListAttemptsParamsStatus) IsKnown() bool {
	switch r {
	case EventListAttemptsParamsStatusFailed, EventListAttemptsParamsStatusPending, EventListAttemptsParamsStatusSending, EventListAttemptsParamsStatusSuccess:
		return true
	}
	return false
}
