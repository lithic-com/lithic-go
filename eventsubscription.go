// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// EventSubscriptionService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventSubscriptionService] method instead.
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
	opts = slices.Concat(r.Options, opts)
	path := "v1/event_subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an event subscription.
func (r *EventSubscriptionService) Get(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/event_subscriptions/%s", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an event subscription.
func (r *EventSubscriptionService) Update(ctx context.Context, eventSubscriptionToken string, body EventSubscriptionUpdateParams, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/event_subscriptions/%s", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all the event subscriptions.
func (r *EventSubscriptionService) List(ctx context.Context, query EventSubscriptionListParams, opts ...option.RequestOption) (res *pagination.CursorPage[EventSubscription], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/event_subscriptions"
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
func (r *EventSubscriptionService) ListAutoPaging(ctx context.Context, query EventSubscriptionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[EventSubscription] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an event subscription.
func (r *EventSubscriptionService) Delete(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/event_subscriptions/%s", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List all the message attempts for a given event subscription.
func (r *EventSubscriptionService) ListAttempts(ctx context.Context, eventSubscriptionToken string, query EventSubscriptionListAttemptsParams, opts ...option.RequestOption) (res *pagination.CursorPage[MessageAttempt], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/event_subscriptions/%s/attempts", eventSubscriptionToken)
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

// List all the message attempts for a given event subscription.
func (r *EventSubscriptionService) ListAttemptsAutoPaging(ctx context.Context, eventSubscriptionToken string, query EventSubscriptionListAttemptsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[MessageAttempt] {
	return pagination.NewCursorPageAutoPager(r.ListAttempts(ctx, eventSubscriptionToken, query, opts...))
}

// Resend all failed messages since a given time.
func (r *EventSubscriptionService) Recover(ctx context.Context, eventSubscriptionToken string, body EventSubscriptionRecoverParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/event_subscriptions/%s/recover", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Replays messages to the endpoint. Only messages that were created after `begin`
// will be sent. Messages that were previously sent to the endpoint are not resent.
// Message will be retried if endpoint responds with a non-2xx status code. See
// [Retry Schedule](https://docs.lithic.com/docs/events-api#retry-schedule) for
// details.
func (r *EventSubscriptionService) ReplayMissing(ctx context.Context, eventSubscriptionToken string, body EventSubscriptionReplayMissingParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/event_subscriptions/%s/replay_missing", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get the secret for an event subscription.
func (r *EventSubscriptionService) GetSecret(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (res *EventSubscriptionGetSecretResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/event_subscriptions/%s/secret", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Rotate the secret for an event subscription. The previous secret will be valid
// for the next 24 hours.
func (r *EventSubscriptionService) RotateSecret(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/event_subscriptions/%s/secret/rotate", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Send an example message for event.
func (r *EventSubscriptionService) SendSimulatedExample(ctx context.Context, eventSubscriptionToken string, body EventSubscriptionSendSimulatedExampleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if eventSubscriptionToken == "" {
		err = errors.New("missing required event_subscription_token parameter")
		return
	}
	path := fmt.Sprintf("v1/simulate/event_subscriptions/%s/send_example", eventSubscriptionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type EventSubscriptionGetSecretResponse struct {
	// The secret for the event subscription.
	Secret string                                 `json:"secret"`
	JSON   eventSubscriptionGetSecretResponseJSON `json:"-"`
}

// eventSubscriptionGetSecretResponseJSON contains the JSON metadata for the struct
// [EventSubscriptionGetSecretResponse]
type eventSubscriptionGetSecretResponseJSON struct {
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventSubscriptionGetSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventSubscriptionGetSecretResponseJSON) RawJSON() string {
	return r.raw
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

// The type of event that occurred. Possible values:
//
//   - account_holder_document.updated: Occurs when an account holder's document
//     upload status has been updated
//   - account_holder.created: Occurs when a new account_holder is created.
//   - account_holder.updated: Occurs when an account_holder is updated.
//   - account_holder.verification: Occurs when an asynchronous account_holder's
//     verification is completed.
//   - auth_rules.backtest_report.created: Auth Rules backtest report created.
//   - balance.updated: Financial Account Balance Update
//   - book_transfer_transaction.created: Occurs when a book transfer transaction is
//     created.
//   - book_transfer_transaction.updated: Occurs when a book transfer transaction is
//     updated.
//   - card_transaction.enhanced_data.created: Occurs when L2/L3 enhanced commercial
//     data is processed for a transaction event.
//   - card_transaction.enhanced_data.updated: Occurs when L2/L3 enhanced commercial
//     data is reprocessed for a transaction event.
//   - card_transaction.updated: Occurs when a card transaction happens.
//   - card.converted: Occurs when a card is converted from virtual to physical
//     cards.
//   - card.created: Occurs when a new card is created.
//   - card.reissued: Occurs when a card is reissued.
//   - card.renewed: Occurs when a card is renewed.
//   - card.shipped: Occurs when a card is shipped.
//   - card.updated: Occurs when a card is updated.
//   - digital_wallet.tokenization_approval_request: Occurs when a tokenization
//     approval request is made. This event will be deprecated in the future. We
//     recommend using `tokenization.approval_request` instead.
//   - digital_wallet.tokenization_result: Occurs when a tokenization request
//     succeeded or failed.
//
// This event will be deprecated in the future. We recommend using
// `tokenization.result` instead.
//
//   - digital_wallet.tokenization_two_factor_authentication_code: Occurs when a
//     tokenization request 2FA code is sent to the Lithic customer for self serve
//     delivery.
//
// This event will be deprecated in the future. We recommend using
// `tokenization.two_factor_authentication_code` instead.
//
//   - digital_wallet.tokenization_two_factor_authentication_code_sent: Occurs when a
//     tokenization request 2FA code is sent to our downstream messaging providers
//     for delivery.
//
// This event will be deprecated in the future. We recommend using
// `tokenization.two_factor_authentication_code_sent` instead.
//
//   - digital_wallet.tokenization_updated: Occurs when a tokenization's status has
//     changed.
//
// This event will be deprecated in the future. We recommend using
// `tokenization.updated` instead.
//
//   - dispute_evidence.upload_failed: Occurs when a dispute evidence upload fails.
//   - dispute_transaction.created: Occurs when a new dispute transaction is created
//   - dispute_transaction.updated: Occurs when a dispute transaction is updated
//   - dispute.updated: Occurs when a dispute is updated.
//   - external_bank_account.created: Occurs when an external bank account is
//     created.
//   - external_bank_account.updated: Occurs when an external bank account is
//     updated.
//   - external_payment.created: Occurs when an external payment is created.
//   - external_payment.updated: Occurs when an external payment is updated.
//   - financial_account.created: Occurs when a financial account is created.
//   - financial_account.updated: Occurs when a financial account is updated.
//   - funding_event.created: Occurs when a funding event is created.
//   - internal_transaction.created: Occurs when an internal adjustment is created.
//   - internal_transaction.updated: Occurs when an internal adjustment is updated.
//   - loan_tape.created: Occurs when a loan tape is created.
//   - loan_tape.updated: Occurs when a loan tape is updated.
//   - management_operation.created: Occurs when an management operation is created.
//   - management_operation.updated: Occurs when an management operation is updated.
//   - network_total.created: Occurs when a network total is created.
//   - network_total.updated: Occurs when a network total is updated.
//   - payment_transaction.created: Occurs when a payment transaction is created.
//   - payment_transaction.updated: Occurs when a payment transaction is updated.
//   - settlement_report.updated: Occurs when a settlement report is created or
//     updated.
//   - statements.created: Occurs when a statement has been created
//   - three_ds_authentication.challenge: The `three_ds_authentication.challenge`
//     event. Upon receiving this request, the Card Program should issue its own
//     challenge to the cardholder. After a cardholder challenge is successfully
//     completed, the Card Program needs to respond back to Lithic by call to
//     [/v1/three_ds_decisioning/challenge_response](https://docs.lithic.com/reference/post_v1-three-ds-decisioning-challenge-response).
//     Then the cardholder must navigate back to the merchant checkout flow to
//     complete the transaction. Some merchants will include an `app_requestor_url`
//     for app-based purchases; Lithic recommends triggering a redirect to that URL
//     after the cardholder completes an app-based challenge.
//   - three_ds_authentication.created: Occurs when a 3DS authentication is created.
//   - three_ds_authentication.updated: Occurs when a 3DS authentication is updated
//     (eg. challenge is completed).
//   - tokenization.approval_request: Occurs when a tokenization approval request is
//     made.
//   - tokenization.result: Occurs when a tokenization request succeeded or failed.
//   - tokenization.two_factor_authentication_code: Occurs when a tokenization
//     request 2FA code is sent to the Lithic customer for self serve delivery.
//   - tokenization.two_factor_authentication_code_sent: Occurs when a tokenization
//     request 2FA code is sent to our downstream messaging providers for delivery.
//   - tokenization.updated: Occurs when a tokenization's status has changed.
type EventSubscriptionNewParamsEventType string

const (
	EventSubscriptionNewParamsEventTypeAccountHolderDocumentUpdated                             EventSubscriptionNewParamsEventType = "account_holder_document.updated"
	EventSubscriptionNewParamsEventTypeAccountHolderCreated                                     EventSubscriptionNewParamsEventType = "account_holder.created"
	EventSubscriptionNewParamsEventTypeAccountHolderUpdated                                     EventSubscriptionNewParamsEventType = "account_holder.updated"
	EventSubscriptionNewParamsEventTypeAccountHolderVerification                                EventSubscriptionNewParamsEventType = "account_holder.verification"
	EventSubscriptionNewParamsEventTypeAuthRulesBacktestReportCreated                           EventSubscriptionNewParamsEventType = "auth_rules.backtest_report.created"
	EventSubscriptionNewParamsEventTypeBalanceUpdated                                           EventSubscriptionNewParamsEventType = "balance.updated"
	EventSubscriptionNewParamsEventTypeBookTransferTransactionCreated                           EventSubscriptionNewParamsEventType = "book_transfer_transaction.created"
	EventSubscriptionNewParamsEventTypeBookTransferTransactionUpdated                           EventSubscriptionNewParamsEventType = "book_transfer_transaction.updated"
	EventSubscriptionNewParamsEventTypeCardTransactionEnhancedDataCreated                       EventSubscriptionNewParamsEventType = "card_transaction.enhanced_data.created"
	EventSubscriptionNewParamsEventTypeCardTransactionEnhancedDataUpdated                       EventSubscriptionNewParamsEventType = "card_transaction.enhanced_data.updated"
	EventSubscriptionNewParamsEventTypeCardTransactionUpdated                                   EventSubscriptionNewParamsEventType = "card_transaction.updated"
	EventSubscriptionNewParamsEventTypeCardConverted                                            EventSubscriptionNewParamsEventType = "card.converted"
	EventSubscriptionNewParamsEventTypeCardCreated                                              EventSubscriptionNewParamsEventType = "card.created"
	EventSubscriptionNewParamsEventTypeCardReissued                                             EventSubscriptionNewParamsEventType = "card.reissued"
	EventSubscriptionNewParamsEventTypeCardRenewed                                              EventSubscriptionNewParamsEventType = "card.renewed"
	EventSubscriptionNewParamsEventTypeCardShipped                                              EventSubscriptionNewParamsEventType = "card.shipped"
	EventSubscriptionNewParamsEventTypeCardUpdated                                              EventSubscriptionNewParamsEventType = "card.updated"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationApprovalRequest                 EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationResult                          EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_result"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationUpdated                         EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_updated"
	EventSubscriptionNewParamsEventTypeDisputeEvidenceUploadFailed                              EventSubscriptionNewParamsEventType = "dispute_evidence.upload_failed"
	EventSubscriptionNewParamsEventTypeDisputeTransactionCreated                                EventSubscriptionNewParamsEventType = "dispute_transaction.created"
	EventSubscriptionNewParamsEventTypeDisputeTransactionUpdated                                EventSubscriptionNewParamsEventType = "dispute_transaction.updated"
	EventSubscriptionNewParamsEventTypeDisputeUpdated                                           EventSubscriptionNewParamsEventType = "dispute.updated"
	EventSubscriptionNewParamsEventTypeExternalBankAccountCreated                               EventSubscriptionNewParamsEventType = "external_bank_account.created"
	EventSubscriptionNewParamsEventTypeExternalBankAccountUpdated                               EventSubscriptionNewParamsEventType = "external_bank_account.updated"
	EventSubscriptionNewParamsEventTypeExternalPaymentCreated                                   EventSubscriptionNewParamsEventType = "external_payment.created"
	EventSubscriptionNewParamsEventTypeExternalPaymentUpdated                                   EventSubscriptionNewParamsEventType = "external_payment.updated"
	EventSubscriptionNewParamsEventTypeFinancialAccountCreated                                  EventSubscriptionNewParamsEventType = "financial_account.created"
	EventSubscriptionNewParamsEventTypeFinancialAccountUpdated                                  EventSubscriptionNewParamsEventType = "financial_account.updated"
	EventSubscriptionNewParamsEventTypeFundingEventCreated                                      EventSubscriptionNewParamsEventType = "funding_event.created"
	EventSubscriptionNewParamsEventTypeInternalTransactionCreated                               EventSubscriptionNewParamsEventType = "internal_transaction.created"
	EventSubscriptionNewParamsEventTypeInternalTransactionUpdated                               EventSubscriptionNewParamsEventType = "internal_transaction.updated"
	EventSubscriptionNewParamsEventTypeLoanTapeCreated                                          EventSubscriptionNewParamsEventType = "loan_tape.created"
	EventSubscriptionNewParamsEventTypeLoanTapeUpdated                                          EventSubscriptionNewParamsEventType = "loan_tape.updated"
	EventSubscriptionNewParamsEventTypeManagementOperationCreated                               EventSubscriptionNewParamsEventType = "management_operation.created"
	EventSubscriptionNewParamsEventTypeManagementOperationUpdated                               EventSubscriptionNewParamsEventType = "management_operation.updated"
	EventSubscriptionNewParamsEventTypeNetworkTotalCreated                                      EventSubscriptionNewParamsEventType = "network_total.created"
	EventSubscriptionNewParamsEventTypeNetworkTotalUpdated                                      EventSubscriptionNewParamsEventType = "network_total.updated"
	EventSubscriptionNewParamsEventTypePaymentTransactionCreated                                EventSubscriptionNewParamsEventType = "payment_transaction.created"
	EventSubscriptionNewParamsEventTypePaymentTransactionUpdated                                EventSubscriptionNewParamsEventType = "payment_transaction.updated"
	EventSubscriptionNewParamsEventTypeSettlementReportUpdated                                  EventSubscriptionNewParamsEventType = "settlement_report.updated"
	EventSubscriptionNewParamsEventTypeStatementsCreated                                        EventSubscriptionNewParamsEventType = "statements.created"
	EventSubscriptionNewParamsEventTypeThreeDSAuthenticationChallenge                           EventSubscriptionNewParamsEventType = "three_ds_authentication.challenge"
	EventSubscriptionNewParamsEventTypeThreeDSAuthenticationCreated                             EventSubscriptionNewParamsEventType = "three_ds_authentication.created"
	EventSubscriptionNewParamsEventTypeThreeDSAuthenticationUpdated                             EventSubscriptionNewParamsEventType = "three_ds_authentication.updated"
	EventSubscriptionNewParamsEventTypeTokenizationApprovalRequest                              EventSubscriptionNewParamsEventType = "tokenization.approval_request"
	EventSubscriptionNewParamsEventTypeTokenizationResult                                       EventSubscriptionNewParamsEventType = "tokenization.result"
	EventSubscriptionNewParamsEventTypeTokenizationTwoFactorAuthenticationCode                  EventSubscriptionNewParamsEventType = "tokenization.two_factor_authentication_code"
	EventSubscriptionNewParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventSubscriptionNewParamsEventType = "tokenization.two_factor_authentication_code_sent"
	EventSubscriptionNewParamsEventTypeTokenizationUpdated                                      EventSubscriptionNewParamsEventType = "tokenization.updated"
)

func (r EventSubscriptionNewParamsEventType) IsKnown() bool {
	switch r {
	case EventSubscriptionNewParamsEventTypeAccountHolderDocumentUpdated, EventSubscriptionNewParamsEventTypeAccountHolderCreated, EventSubscriptionNewParamsEventTypeAccountHolderUpdated, EventSubscriptionNewParamsEventTypeAccountHolderVerification, EventSubscriptionNewParamsEventTypeAuthRulesBacktestReportCreated, EventSubscriptionNewParamsEventTypeBalanceUpdated, EventSubscriptionNewParamsEventTypeBookTransferTransactionCreated, EventSubscriptionNewParamsEventTypeBookTransferTransactionUpdated, EventSubscriptionNewParamsEventTypeCardTransactionEnhancedDataCreated, EventSubscriptionNewParamsEventTypeCardTransactionEnhancedDataUpdated, EventSubscriptionNewParamsEventTypeCardTransactionUpdated, EventSubscriptionNewParamsEventTypeCardConverted, EventSubscriptionNewParamsEventTypeCardCreated, EventSubscriptionNewParamsEventTypeCardReissued, EventSubscriptionNewParamsEventTypeCardRenewed, EventSubscriptionNewParamsEventTypeCardShipped, EventSubscriptionNewParamsEventTypeCardUpdated, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationApprovalRequest, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationResult, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationUpdated, EventSubscriptionNewParamsEventTypeDisputeEvidenceUploadFailed, EventSubscriptionNewParamsEventTypeDisputeTransactionCreated, EventSubscriptionNewParamsEventTypeDisputeTransactionUpdated, EventSubscriptionNewParamsEventTypeDisputeUpdated, EventSubscriptionNewParamsEventTypeExternalBankAccountCreated, EventSubscriptionNewParamsEventTypeExternalBankAccountUpdated, EventSubscriptionNewParamsEventTypeExternalPaymentCreated, EventSubscriptionNewParamsEventTypeExternalPaymentUpdated, EventSubscriptionNewParamsEventTypeFinancialAccountCreated, EventSubscriptionNewParamsEventTypeFinancialAccountUpdated, EventSubscriptionNewParamsEventTypeFundingEventCreated, EventSubscriptionNewParamsEventTypeInternalTransactionCreated, EventSubscriptionNewParamsEventTypeInternalTransactionUpdated, EventSubscriptionNewParamsEventTypeLoanTapeCreated, EventSubscriptionNewParamsEventTypeLoanTapeUpdated, EventSubscriptionNewParamsEventTypeManagementOperationCreated, EventSubscriptionNewParamsEventTypeManagementOperationUpdated, EventSubscriptionNewParamsEventTypeNetworkTotalCreated, EventSubscriptionNewParamsEventTypeNetworkTotalUpdated, EventSubscriptionNewParamsEventTypePaymentTransactionCreated, EventSubscriptionNewParamsEventTypePaymentTransactionUpdated, EventSubscriptionNewParamsEventTypeSettlementReportUpdated, EventSubscriptionNewParamsEventTypeStatementsCreated, EventSubscriptionNewParamsEventTypeThreeDSAuthenticationChallenge, EventSubscriptionNewParamsEventTypeThreeDSAuthenticationCreated, EventSubscriptionNewParamsEventTypeThreeDSAuthenticationUpdated, EventSubscriptionNewParamsEventTypeTokenizationApprovalRequest, EventSubscriptionNewParamsEventTypeTokenizationResult, EventSubscriptionNewParamsEventTypeTokenizationTwoFactorAuthenticationCode, EventSubscriptionNewParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionNewParamsEventTypeTokenizationUpdated:
		return true
	}
	return false
}

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

// The type of event that occurred. Possible values:
//
//   - account_holder_document.updated: Occurs when an account holder's document
//     upload status has been updated
//   - account_holder.created: Occurs when a new account_holder is created.
//   - account_holder.updated: Occurs when an account_holder is updated.
//   - account_holder.verification: Occurs when an asynchronous account_holder's
//     verification is completed.
//   - auth_rules.backtest_report.created: Auth Rules backtest report created.
//   - balance.updated: Financial Account Balance Update
//   - book_transfer_transaction.created: Occurs when a book transfer transaction is
//     created.
//   - book_transfer_transaction.updated: Occurs when a book transfer transaction is
//     updated.
//   - card_transaction.enhanced_data.created: Occurs when L2/L3 enhanced commercial
//     data is processed for a transaction event.
//   - card_transaction.enhanced_data.updated: Occurs when L2/L3 enhanced commercial
//     data is reprocessed for a transaction event.
//   - card_transaction.updated: Occurs when a card transaction happens.
//   - card.converted: Occurs when a card is converted from virtual to physical
//     cards.
//   - card.created: Occurs when a new card is created.
//   - card.reissued: Occurs when a card is reissued.
//   - card.renewed: Occurs when a card is renewed.
//   - card.shipped: Occurs when a card is shipped.
//   - card.updated: Occurs when a card is updated.
//   - digital_wallet.tokenization_approval_request: Occurs when a tokenization
//     approval request is made. This event will be deprecated in the future. We
//     recommend using `tokenization.approval_request` instead.
//   - digital_wallet.tokenization_result: Occurs when a tokenization request
//     succeeded or failed.
//
// This event will be deprecated in the future. We recommend using
// `tokenization.result` instead.
//
//   - digital_wallet.tokenization_two_factor_authentication_code: Occurs when a
//     tokenization request 2FA code is sent to the Lithic customer for self serve
//     delivery.
//
// This event will be deprecated in the future. We recommend using
// `tokenization.two_factor_authentication_code` instead.
//
//   - digital_wallet.tokenization_two_factor_authentication_code_sent: Occurs when a
//     tokenization request 2FA code is sent to our downstream messaging providers
//     for delivery.
//
// This event will be deprecated in the future. We recommend using
// `tokenization.two_factor_authentication_code_sent` instead.
//
//   - digital_wallet.tokenization_updated: Occurs when a tokenization's status has
//     changed.
//
// This event will be deprecated in the future. We recommend using
// `tokenization.updated` instead.
//
//   - dispute_evidence.upload_failed: Occurs when a dispute evidence upload fails.
//   - dispute_transaction.created: Occurs when a new dispute transaction is created
//   - dispute_transaction.updated: Occurs when a dispute transaction is updated
//   - dispute.updated: Occurs when a dispute is updated.
//   - external_bank_account.created: Occurs when an external bank account is
//     created.
//   - external_bank_account.updated: Occurs when an external bank account is
//     updated.
//   - external_payment.created: Occurs when an external payment is created.
//   - external_payment.updated: Occurs when an external payment is updated.
//   - financial_account.created: Occurs when a financial account is created.
//   - financial_account.updated: Occurs when a financial account is updated.
//   - funding_event.created: Occurs when a funding event is created.
//   - internal_transaction.created: Occurs when an internal adjustment is created.
//   - internal_transaction.updated: Occurs when an internal adjustment is updated.
//   - loan_tape.created: Occurs when a loan tape is created.
//   - loan_tape.updated: Occurs when a loan tape is updated.
//   - management_operation.created: Occurs when an management operation is created.
//   - management_operation.updated: Occurs when an management operation is updated.
//   - network_total.created: Occurs when a network total is created.
//   - network_total.updated: Occurs when a network total is updated.
//   - payment_transaction.created: Occurs when a payment transaction is created.
//   - payment_transaction.updated: Occurs when a payment transaction is updated.
//   - settlement_report.updated: Occurs when a settlement report is created or
//     updated.
//   - statements.created: Occurs when a statement has been created
//   - three_ds_authentication.challenge: The `three_ds_authentication.challenge`
//     event. Upon receiving this request, the Card Program should issue its own
//     challenge to the cardholder. After a cardholder challenge is successfully
//     completed, the Card Program needs to respond back to Lithic by call to
//     [/v1/three_ds_decisioning/challenge_response](https://docs.lithic.com/reference/post_v1-three-ds-decisioning-challenge-response).
//     Then the cardholder must navigate back to the merchant checkout flow to
//     complete the transaction. Some merchants will include an `app_requestor_url`
//     for app-based purchases; Lithic recommends triggering a redirect to that URL
//     after the cardholder completes an app-based challenge.
//   - three_ds_authentication.created: Occurs when a 3DS authentication is created.
//   - three_ds_authentication.updated: Occurs when a 3DS authentication is updated
//     (eg. challenge is completed).
//   - tokenization.approval_request: Occurs when a tokenization approval request is
//     made.
//   - tokenization.result: Occurs when a tokenization request succeeded or failed.
//   - tokenization.two_factor_authentication_code: Occurs when a tokenization
//     request 2FA code is sent to the Lithic customer for self serve delivery.
//   - tokenization.two_factor_authentication_code_sent: Occurs when a tokenization
//     request 2FA code is sent to our downstream messaging providers for delivery.
//   - tokenization.updated: Occurs when a tokenization's status has changed.
type EventSubscriptionUpdateParamsEventType string

const (
	EventSubscriptionUpdateParamsEventTypeAccountHolderDocumentUpdated                             EventSubscriptionUpdateParamsEventType = "account_holder_document.updated"
	EventSubscriptionUpdateParamsEventTypeAccountHolderCreated                                     EventSubscriptionUpdateParamsEventType = "account_holder.created"
	EventSubscriptionUpdateParamsEventTypeAccountHolderUpdated                                     EventSubscriptionUpdateParamsEventType = "account_holder.updated"
	EventSubscriptionUpdateParamsEventTypeAccountHolderVerification                                EventSubscriptionUpdateParamsEventType = "account_holder.verification"
	EventSubscriptionUpdateParamsEventTypeAuthRulesBacktestReportCreated                           EventSubscriptionUpdateParamsEventType = "auth_rules.backtest_report.created"
	EventSubscriptionUpdateParamsEventTypeBalanceUpdated                                           EventSubscriptionUpdateParamsEventType = "balance.updated"
	EventSubscriptionUpdateParamsEventTypeBookTransferTransactionCreated                           EventSubscriptionUpdateParamsEventType = "book_transfer_transaction.created"
	EventSubscriptionUpdateParamsEventTypeBookTransferTransactionUpdated                           EventSubscriptionUpdateParamsEventType = "book_transfer_transaction.updated"
	EventSubscriptionUpdateParamsEventTypeCardTransactionEnhancedDataCreated                       EventSubscriptionUpdateParamsEventType = "card_transaction.enhanced_data.created"
	EventSubscriptionUpdateParamsEventTypeCardTransactionEnhancedDataUpdated                       EventSubscriptionUpdateParamsEventType = "card_transaction.enhanced_data.updated"
	EventSubscriptionUpdateParamsEventTypeCardTransactionUpdated                                   EventSubscriptionUpdateParamsEventType = "card_transaction.updated"
	EventSubscriptionUpdateParamsEventTypeCardConverted                                            EventSubscriptionUpdateParamsEventType = "card.converted"
	EventSubscriptionUpdateParamsEventTypeCardCreated                                              EventSubscriptionUpdateParamsEventType = "card.created"
	EventSubscriptionUpdateParamsEventTypeCardReissued                                             EventSubscriptionUpdateParamsEventType = "card.reissued"
	EventSubscriptionUpdateParamsEventTypeCardRenewed                                              EventSubscriptionUpdateParamsEventType = "card.renewed"
	EventSubscriptionUpdateParamsEventTypeCardShipped                                              EventSubscriptionUpdateParamsEventType = "card.shipped"
	EventSubscriptionUpdateParamsEventTypeCardUpdated                                              EventSubscriptionUpdateParamsEventType = "card.updated"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationApprovalRequest                 EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationResult                          EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_result"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationUpdated                         EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_updated"
	EventSubscriptionUpdateParamsEventTypeDisputeEvidenceUploadFailed                              EventSubscriptionUpdateParamsEventType = "dispute_evidence.upload_failed"
	EventSubscriptionUpdateParamsEventTypeDisputeTransactionCreated                                EventSubscriptionUpdateParamsEventType = "dispute_transaction.created"
	EventSubscriptionUpdateParamsEventTypeDisputeTransactionUpdated                                EventSubscriptionUpdateParamsEventType = "dispute_transaction.updated"
	EventSubscriptionUpdateParamsEventTypeDisputeUpdated                                           EventSubscriptionUpdateParamsEventType = "dispute.updated"
	EventSubscriptionUpdateParamsEventTypeExternalBankAccountCreated                               EventSubscriptionUpdateParamsEventType = "external_bank_account.created"
	EventSubscriptionUpdateParamsEventTypeExternalBankAccountUpdated                               EventSubscriptionUpdateParamsEventType = "external_bank_account.updated"
	EventSubscriptionUpdateParamsEventTypeExternalPaymentCreated                                   EventSubscriptionUpdateParamsEventType = "external_payment.created"
	EventSubscriptionUpdateParamsEventTypeExternalPaymentUpdated                                   EventSubscriptionUpdateParamsEventType = "external_payment.updated"
	EventSubscriptionUpdateParamsEventTypeFinancialAccountCreated                                  EventSubscriptionUpdateParamsEventType = "financial_account.created"
	EventSubscriptionUpdateParamsEventTypeFinancialAccountUpdated                                  EventSubscriptionUpdateParamsEventType = "financial_account.updated"
	EventSubscriptionUpdateParamsEventTypeFundingEventCreated                                      EventSubscriptionUpdateParamsEventType = "funding_event.created"
	EventSubscriptionUpdateParamsEventTypeInternalTransactionCreated                               EventSubscriptionUpdateParamsEventType = "internal_transaction.created"
	EventSubscriptionUpdateParamsEventTypeInternalTransactionUpdated                               EventSubscriptionUpdateParamsEventType = "internal_transaction.updated"
	EventSubscriptionUpdateParamsEventTypeLoanTapeCreated                                          EventSubscriptionUpdateParamsEventType = "loan_tape.created"
	EventSubscriptionUpdateParamsEventTypeLoanTapeUpdated                                          EventSubscriptionUpdateParamsEventType = "loan_tape.updated"
	EventSubscriptionUpdateParamsEventTypeManagementOperationCreated                               EventSubscriptionUpdateParamsEventType = "management_operation.created"
	EventSubscriptionUpdateParamsEventTypeManagementOperationUpdated                               EventSubscriptionUpdateParamsEventType = "management_operation.updated"
	EventSubscriptionUpdateParamsEventTypeNetworkTotalCreated                                      EventSubscriptionUpdateParamsEventType = "network_total.created"
	EventSubscriptionUpdateParamsEventTypeNetworkTotalUpdated                                      EventSubscriptionUpdateParamsEventType = "network_total.updated"
	EventSubscriptionUpdateParamsEventTypePaymentTransactionCreated                                EventSubscriptionUpdateParamsEventType = "payment_transaction.created"
	EventSubscriptionUpdateParamsEventTypePaymentTransactionUpdated                                EventSubscriptionUpdateParamsEventType = "payment_transaction.updated"
	EventSubscriptionUpdateParamsEventTypeSettlementReportUpdated                                  EventSubscriptionUpdateParamsEventType = "settlement_report.updated"
	EventSubscriptionUpdateParamsEventTypeStatementsCreated                                        EventSubscriptionUpdateParamsEventType = "statements.created"
	EventSubscriptionUpdateParamsEventTypeThreeDSAuthenticationChallenge                           EventSubscriptionUpdateParamsEventType = "three_ds_authentication.challenge"
	EventSubscriptionUpdateParamsEventTypeThreeDSAuthenticationCreated                             EventSubscriptionUpdateParamsEventType = "three_ds_authentication.created"
	EventSubscriptionUpdateParamsEventTypeThreeDSAuthenticationUpdated                             EventSubscriptionUpdateParamsEventType = "three_ds_authentication.updated"
	EventSubscriptionUpdateParamsEventTypeTokenizationApprovalRequest                              EventSubscriptionUpdateParamsEventType = "tokenization.approval_request"
	EventSubscriptionUpdateParamsEventTypeTokenizationResult                                       EventSubscriptionUpdateParamsEventType = "tokenization.result"
	EventSubscriptionUpdateParamsEventTypeTokenizationTwoFactorAuthenticationCode                  EventSubscriptionUpdateParamsEventType = "tokenization.two_factor_authentication_code"
	EventSubscriptionUpdateParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventSubscriptionUpdateParamsEventType = "tokenization.two_factor_authentication_code_sent"
	EventSubscriptionUpdateParamsEventTypeTokenizationUpdated                                      EventSubscriptionUpdateParamsEventType = "tokenization.updated"
)

func (r EventSubscriptionUpdateParamsEventType) IsKnown() bool {
	switch r {
	case EventSubscriptionUpdateParamsEventTypeAccountHolderDocumentUpdated, EventSubscriptionUpdateParamsEventTypeAccountHolderCreated, EventSubscriptionUpdateParamsEventTypeAccountHolderUpdated, EventSubscriptionUpdateParamsEventTypeAccountHolderVerification, EventSubscriptionUpdateParamsEventTypeAuthRulesBacktestReportCreated, EventSubscriptionUpdateParamsEventTypeBalanceUpdated, EventSubscriptionUpdateParamsEventTypeBookTransferTransactionCreated, EventSubscriptionUpdateParamsEventTypeBookTransferTransactionUpdated, EventSubscriptionUpdateParamsEventTypeCardTransactionEnhancedDataCreated, EventSubscriptionUpdateParamsEventTypeCardTransactionEnhancedDataUpdated, EventSubscriptionUpdateParamsEventTypeCardTransactionUpdated, EventSubscriptionUpdateParamsEventTypeCardConverted, EventSubscriptionUpdateParamsEventTypeCardCreated, EventSubscriptionUpdateParamsEventTypeCardReissued, EventSubscriptionUpdateParamsEventTypeCardRenewed, EventSubscriptionUpdateParamsEventTypeCardShipped, EventSubscriptionUpdateParamsEventTypeCardUpdated, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationApprovalRequest, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationResult, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationUpdated, EventSubscriptionUpdateParamsEventTypeDisputeEvidenceUploadFailed, EventSubscriptionUpdateParamsEventTypeDisputeTransactionCreated, EventSubscriptionUpdateParamsEventTypeDisputeTransactionUpdated, EventSubscriptionUpdateParamsEventTypeDisputeUpdated, EventSubscriptionUpdateParamsEventTypeExternalBankAccountCreated, EventSubscriptionUpdateParamsEventTypeExternalBankAccountUpdated, EventSubscriptionUpdateParamsEventTypeExternalPaymentCreated, EventSubscriptionUpdateParamsEventTypeExternalPaymentUpdated, EventSubscriptionUpdateParamsEventTypeFinancialAccountCreated, EventSubscriptionUpdateParamsEventTypeFinancialAccountUpdated, EventSubscriptionUpdateParamsEventTypeFundingEventCreated, EventSubscriptionUpdateParamsEventTypeInternalTransactionCreated, EventSubscriptionUpdateParamsEventTypeInternalTransactionUpdated, EventSubscriptionUpdateParamsEventTypeLoanTapeCreated, EventSubscriptionUpdateParamsEventTypeLoanTapeUpdated, EventSubscriptionUpdateParamsEventTypeManagementOperationCreated, EventSubscriptionUpdateParamsEventTypeManagementOperationUpdated, EventSubscriptionUpdateParamsEventTypeNetworkTotalCreated, EventSubscriptionUpdateParamsEventTypeNetworkTotalUpdated, EventSubscriptionUpdateParamsEventTypePaymentTransactionCreated, EventSubscriptionUpdateParamsEventTypePaymentTransactionUpdated, EventSubscriptionUpdateParamsEventTypeSettlementReportUpdated, EventSubscriptionUpdateParamsEventTypeStatementsCreated, EventSubscriptionUpdateParamsEventTypeThreeDSAuthenticationChallenge, EventSubscriptionUpdateParamsEventTypeThreeDSAuthenticationCreated, EventSubscriptionUpdateParamsEventTypeThreeDSAuthenticationUpdated, EventSubscriptionUpdateParamsEventTypeTokenizationApprovalRequest, EventSubscriptionUpdateParamsEventTypeTokenizationResult, EventSubscriptionUpdateParamsEventTypeTokenizationTwoFactorAuthenticationCode, EventSubscriptionUpdateParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionUpdateParamsEventTypeTokenizationUpdated:
		return true
	}
	return false
}

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

type EventSubscriptionListAttemptsParams struct {
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
	StartingAfter param.Field[string]                                    `query:"starting_after"`
	Status        param.Field[EventSubscriptionListAttemptsParamsStatus] `query:"status"`
}

// URLQuery serializes [EventSubscriptionListAttemptsParams]'s query parameters as
// `url.Values`.
func (r EventSubscriptionListAttemptsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventSubscriptionListAttemptsParamsStatus string

const (
	EventSubscriptionListAttemptsParamsStatusFailed  EventSubscriptionListAttemptsParamsStatus = "FAILED"
	EventSubscriptionListAttemptsParamsStatusPending EventSubscriptionListAttemptsParamsStatus = "PENDING"
	EventSubscriptionListAttemptsParamsStatusSending EventSubscriptionListAttemptsParamsStatus = "SENDING"
	EventSubscriptionListAttemptsParamsStatusSuccess EventSubscriptionListAttemptsParamsStatus = "SUCCESS"
)

func (r EventSubscriptionListAttemptsParamsStatus) IsKnown() bool {
	switch r {
	case EventSubscriptionListAttemptsParamsStatusFailed, EventSubscriptionListAttemptsParamsStatusPending, EventSubscriptionListAttemptsParamsStatusSending, EventSubscriptionListAttemptsParamsStatusSuccess:
		return true
	}
	return false
}

type EventSubscriptionRecoverParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
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
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
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

type EventSubscriptionSendSimulatedExampleParams struct {
	// Event type to send example message for.
	EventType param.Field[EventSubscriptionSendSimulatedExampleParamsEventType] `json:"event_type"`
}

func (r EventSubscriptionSendSimulatedExampleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Event type to send example message for.
type EventSubscriptionSendSimulatedExampleParamsEventType string

const (
	EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderDocumentUpdated                             EventSubscriptionSendSimulatedExampleParamsEventType = "account_holder_document.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderCreated                                     EventSubscriptionSendSimulatedExampleParamsEventType = "account_holder.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderUpdated                                     EventSubscriptionSendSimulatedExampleParamsEventType = "account_holder.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderVerification                                EventSubscriptionSendSimulatedExampleParamsEventType = "account_holder.verification"
	EventSubscriptionSendSimulatedExampleParamsEventTypeAuthRulesBacktestReportCreated                           EventSubscriptionSendSimulatedExampleParamsEventType = "auth_rules.backtest_report.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeBalanceUpdated                                           EventSubscriptionSendSimulatedExampleParamsEventType = "balance.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeBookTransferTransactionCreated                           EventSubscriptionSendSimulatedExampleParamsEventType = "book_transfer_transaction.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeBookTransferTransactionUpdated                           EventSubscriptionSendSimulatedExampleParamsEventType = "book_transfer_transaction.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardTransactionEnhancedDataCreated                       EventSubscriptionSendSimulatedExampleParamsEventType = "card_transaction.enhanced_data.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardTransactionEnhancedDataUpdated                       EventSubscriptionSendSimulatedExampleParamsEventType = "card_transaction.enhanced_data.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardTransactionUpdated                                   EventSubscriptionSendSimulatedExampleParamsEventType = "card_transaction.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardConverted                                            EventSubscriptionSendSimulatedExampleParamsEventType = "card.converted"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardCreated                                              EventSubscriptionSendSimulatedExampleParamsEventType = "card.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardReissued                                             EventSubscriptionSendSimulatedExampleParamsEventType = "card.reissued"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardRenewed                                              EventSubscriptionSendSimulatedExampleParamsEventType = "card.renewed"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardShipped                                              EventSubscriptionSendSimulatedExampleParamsEventType = "card.shipped"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardUpdated                                              EventSubscriptionSendSimulatedExampleParamsEventType = "card.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationApprovalRequest                 EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationResult                          EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_result"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationUpdated                         EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeEvidenceUploadFailed                              EventSubscriptionSendSimulatedExampleParamsEventType = "dispute_evidence.upload_failed"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeTransactionCreated                                EventSubscriptionSendSimulatedExampleParamsEventType = "dispute_transaction.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeTransactionUpdated                                EventSubscriptionSendSimulatedExampleParamsEventType = "dispute_transaction.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeUpdated                                           EventSubscriptionSendSimulatedExampleParamsEventType = "dispute.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeExternalBankAccountCreated                               EventSubscriptionSendSimulatedExampleParamsEventType = "external_bank_account.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeExternalBankAccountUpdated                               EventSubscriptionSendSimulatedExampleParamsEventType = "external_bank_account.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeExternalPaymentCreated                                   EventSubscriptionSendSimulatedExampleParamsEventType = "external_payment.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeExternalPaymentUpdated                                   EventSubscriptionSendSimulatedExampleParamsEventType = "external_payment.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeFinancialAccountCreated                                  EventSubscriptionSendSimulatedExampleParamsEventType = "financial_account.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeFinancialAccountUpdated                                  EventSubscriptionSendSimulatedExampleParamsEventType = "financial_account.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeFundingEventCreated                                      EventSubscriptionSendSimulatedExampleParamsEventType = "funding_event.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeInternalTransactionCreated                               EventSubscriptionSendSimulatedExampleParamsEventType = "internal_transaction.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeInternalTransactionUpdated                               EventSubscriptionSendSimulatedExampleParamsEventType = "internal_transaction.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeLoanTapeCreated                                          EventSubscriptionSendSimulatedExampleParamsEventType = "loan_tape.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeLoanTapeUpdated                                          EventSubscriptionSendSimulatedExampleParamsEventType = "loan_tape.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeManagementOperationCreated                               EventSubscriptionSendSimulatedExampleParamsEventType = "management_operation.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeManagementOperationUpdated                               EventSubscriptionSendSimulatedExampleParamsEventType = "management_operation.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeNetworkTotalCreated                                      EventSubscriptionSendSimulatedExampleParamsEventType = "network_total.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeNetworkTotalUpdated                                      EventSubscriptionSendSimulatedExampleParamsEventType = "network_total.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypePaymentTransactionCreated                                EventSubscriptionSendSimulatedExampleParamsEventType = "payment_transaction.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypePaymentTransactionUpdated                                EventSubscriptionSendSimulatedExampleParamsEventType = "payment_transaction.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeSettlementReportUpdated                                  EventSubscriptionSendSimulatedExampleParamsEventType = "settlement_report.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeStatementsCreated                                        EventSubscriptionSendSimulatedExampleParamsEventType = "statements.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeThreeDSAuthenticationChallenge                           EventSubscriptionSendSimulatedExampleParamsEventType = "three_ds_authentication.challenge"
	EventSubscriptionSendSimulatedExampleParamsEventTypeThreeDSAuthenticationCreated                             EventSubscriptionSendSimulatedExampleParamsEventType = "three_ds_authentication.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeThreeDSAuthenticationUpdated                             EventSubscriptionSendSimulatedExampleParamsEventType = "three_ds_authentication.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationApprovalRequest                              EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.approval_request"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationResult                                       EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.result"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationTwoFactorAuthenticationCode                  EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.two_factor_authentication_code"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.two_factor_authentication_code_sent"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationUpdated                                      EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.updated"
)

func (r EventSubscriptionSendSimulatedExampleParamsEventType) IsKnown() bool {
	switch r {
	case EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderDocumentUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderVerification, EventSubscriptionSendSimulatedExampleParamsEventTypeAuthRulesBacktestReportCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeBalanceUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeBookTransferTransactionCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeBookTransferTransactionUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeCardTransactionEnhancedDataCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeCardTransactionEnhancedDataUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeCardTransactionUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeCardConverted, EventSubscriptionSendSimulatedExampleParamsEventTypeCardCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeCardReissued, EventSubscriptionSendSimulatedExampleParamsEventTypeCardRenewed, EventSubscriptionSendSimulatedExampleParamsEventTypeCardShipped, EventSubscriptionSendSimulatedExampleParamsEventTypeCardUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationApprovalRequest, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationResult, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeEvidenceUploadFailed, EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeTransactionCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeTransactionUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeExternalBankAccountCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeExternalBankAccountUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeExternalPaymentCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeExternalPaymentUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeFinancialAccountCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeFinancialAccountUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeFundingEventCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeInternalTransactionCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeInternalTransactionUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeLoanTapeCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeLoanTapeUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeManagementOperationCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeManagementOperationUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeNetworkTotalCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeNetworkTotalUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypePaymentTransactionCreated, EventSubscriptionSendSimulatedExampleParamsEventTypePaymentTransactionUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeSettlementReportUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeStatementsCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeThreeDSAuthenticationChallenge, EventSubscriptionSendSimulatedExampleParamsEventTypeThreeDSAuthenticationCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeThreeDSAuthenticationUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationApprovalRequest, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationResult, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationTwoFactorAuthenticationCode, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationUpdated:
		return true
	}
	return false
}
