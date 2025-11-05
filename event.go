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

// EventService contains methods and other services that help with interacting with
// the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options            []option.RequestOption
	Subscriptions      *EventSubscriptionService
	EventSubscriptions *EventEventSubscriptionService
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	r.Subscriptions = NewEventSubscriptionService(opts...)
	r.EventSubscriptions = NewEventEventSubscriptionService(opts...)
	return
}

// Get an event.
func (r *EventService) Get(ctx context.Context, eventToken string, opts ...option.RequestOption) (res *Event, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventToken == "" {
		err = errors.New("missing required event_token parameter")
		return
	}
	path := fmt.Sprintf("v1/events/%s", eventToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all events.
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Event], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/events"
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if eventToken == "" {
		err = errors.New("missing required event_token parameter")
		return
	}
	path := fmt.Sprintf("v1/events/%s/attempts", eventToken)
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
type EventEventType string

const (
	EventEventTypeAccountHolderDocumentUpdated                             EventEventType = "account_holder_document.updated"
	EventEventTypeAccountHolderCreated                                     EventEventType = "account_holder.created"
	EventEventTypeAccountHolderUpdated                                     EventEventType = "account_holder.updated"
	EventEventTypeAccountHolderVerification                                EventEventType = "account_holder.verification"
	EventEventTypeAuthRulesBacktestReportCreated                           EventEventType = "auth_rules.backtest_report.created"
	EventEventTypeBalanceUpdated                                           EventEventType = "balance.updated"
	EventEventTypeBookTransferTransactionCreated                           EventEventType = "book_transfer_transaction.created"
	EventEventTypeBookTransferTransactionUpdated                           EventEventType = "book_transfer_transaction.updated"
	EventEventTypeCardTransactionEnhancedDataCreated                       EventEventType = "card_transaction.enhanced_data.created"
	EventEventTypeCardTransactionEnhancedDataUpdated                       EventEventType = "card_transaction.enhanced_data.updated"
	EventEventTypeCardTransactionUpdated                                   EventEventType = "card_transaction.updated"
	EventEventTypeCardConverted                                            EventEventType = "card.converted"
	EventEventTypeCardCreated                                              EventEventType = "card.created"
	EventEventTypeCardReissued                                             EventEventType = "card.reissued"
	EventEventTypeCardRenewed                                              EventEventType = "card.renewed"
	EventEventTypeCardShipped                                              EventEventType = "card.shipped"
	EventEventTypeDigitalWalletTokenizationApprovalRequest                 EventEventType = "digital_wallet.tokenization_approval_request"
	EventEventTypeDigitalWalletTokenizationResult                          EventEventType = "digital_wallet.tokenization_result"
	EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventEventTypeDigitalWalletTokenizationUpdated                         EventEventType = "digital_wallet.tokenization_updated"
	EventEventTypeDisputeEvidenceUploadFailed                              EventEventType = "dispute_evidence.upload_failed"
	EventEventTypeDisputeTransactionCreated                                EventEventType = "dispute_transaction.created"
	EventEventTypeDisputeTransactionUpdated                                EventEventType = "dispute_transaction.updated"
	EventEventTypeDisputeUpdated                                           EventEventType = "dispute.updated"
	EventEventTypeExternalBankAccountCreated                               EventEventType = "external_bank_account.created"
	EventEventTypeExternalBankAccountUpdated                               EventEventType = "external_bank_account.updated"
	EventEventTypeExternalPaymentCreated                                   EventEventType = "external_payment.created"
	EventEventTypeExternalPaymentUpdated                                   EventEventType = "external_payment.updated"
	EventEventTypeFinancialAccountCreated                                  EventEventType = "financial_account.created"
	EventEventTypeFinancialAccountUpdated                                  EventEventType = "financial_account.updated"
	EventEventTypeFundingEventCreated                                      EventEventType = "funding_event.created"
	EventEventTypeInternalTransactionCreated                               EventEventType = "internal_transaction.created"
	EventEventTypeInternalTransactionUpdated                               EventEventType = "internal_transaction.updated"
	EventEventTypeLoanTapeCreated                                          EventEventType = "loan_tape.created"
	EventEventTypeLoanTapeUpdated                                          EventEventType = "loan_tape.updated"
	EventEventTypeManagementOperationCreated                               EventEventType = "management_operation.created"
	EventEventTypeManagementOperationUpdated                               EventEventType = "management_operation.updated"
	EventEventTypeNetworkTotalCreated                                      EventEventType = "network_total.created"
	EventEventTypeNetworkTotalUpdated                                      EventEventType = "network_total.updated"
	EventEventTypePaymentTransactionCreated                                EventEventType = "payment_transaction.created"
	EventEventTypePaymentTransactionUpdated                                EventEventType = "payment_transaction.updated"
	EventEventTypeSettlementReportUpdated                                  EventEventType = "settlement_report.updated"
	EventEventTypeStatementsCreated                                        EventEventType = "statements.created"
	EventEventTypeThreeDSAuthenticationChallenge                           EventEventType = "three_ds_authentication.challenge"
	EventEventTypeThreeDSAuthenticationCreated                             EventEventType = "three_ds_authentication.created"
	EventEventTypeThreeDSAuthenticationUpdated                             EventEventType = "three_ds_authentication.updated"
	EventEventTypeTokenizationApprovalRequest                              EventEventType = "tokenization.approval_request"
	EventEventTypeTokenizationResult                                       EventEventType = "tokenization.result"
	EventEventTypeTokenizationTwoFactorAuthenticationCode                  EventEventType = "tokenization.two_factor_authentication_code"
	EventEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventEventType = "tokenization.two_factor_authentication_code_sent"
	EventEventTypeTokenizationUpdated                                      EventEventType = "tokenization.updated"
)

func (r EventEventType) IsKnown() bool {
	switch r {
	case EventEventTypeAccountHolderDocumentUpdated, EventEventTypeAccountHolderCreated, EventEventTypeAccountHolderUpdated, EventEventTypeAccountHolderVerification, EventEventTypeAuthRulesBacktestReportCreated, EventEventTypeBalanceUpdated, EventEventTypeBookTransferTransactionCreated, EventEventTypeBookTransferTransactionUpdated, EventEventTypeCardTransactionEnhancedDataCreated, EventEventTypeCardTransactionEnhancedDataUpdated, EventEventTypeCardTransactionUpdated, EventEventTypeCardConverted, EventEventTypeCardCreated, EventEventTypeCardReissued, EventEventTypeCardRenewed, EventEventTypeCardShipped, EventEventTypeDigitalWalletTokenizationApprovalRequest, EventEventTypeDigitalWalletTokenizationResult, EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventEventTypeDigitalWalletTokenizationUpdated, EventEventTypeDisputeEvidenceUploadFailed, EventEventTypeDisputeTransactionCreated, EventEventTypeDisputeTransactionUpdated, EventEventTypeDisputeUpdated, EventEventTypeExternalBankAccountCreated, EventEventTypeExternalBankAccountUpdated, EventEventTypeExternalPaymentCreated, EventEventTypeExternalPaymentUpdated, EventEventTypeFinancialAccountCreated, EventEventTypeFinancialAccountUpdated, EventEventTypeFundingEventCreated, EventEventTypeInternalTransactionCreated, EventEventTypeInternalTransactionUpdated, EventEventTypeLoanTapeCreated, EventEventTypeLoanTapeUpdated, EventEventTypeManagementOperationCreated, EventEventTypeManagementOperationUpdated, EventEventTypeNetworkTotalCreated, EventEventTypeNetworkTotalUpdated, EventEventTypePaymentTransactionCreated, EventEventTypePaymentTransactionUpdated, EventEventTypeSettlementReportUpdated, EventEventTypeStatementsCreated, EventEventTypeThreeDSAuthenticationChallenge, EventEventTypeThreeDSAuthenticationCreated, EventEventTypeThreeDSAuthenticationUpdated, EventEventTypeTokenizationApprovalRequest, EventEventTypeTokenizationResult, EventEventTypeTokenizationTwoFactorAuthenticationCode, EventEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventEventTypeTokenizationUpdated:
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
type EventSubscriptionEventType string

const (
	EventSubscriptionEventTypeAccountHolderDocumentUpdated                             EventSubscriptionEventType = "account_holder_document.updated"
	EventSubscriptionEventTypeAccountHolderCreated                                     EventSubscriptionEventType = "account_holder.created"
	EventSubscriptionEventTypeAccountHolderUpdated                                     EventSubscriptionEventType = "account_holder.updated"
	EventSubscriptionEventTypeAccountHolderVerification                                EventSubscriptionEventType = "account_holder.verification"
	EventSubscriptionEventTypeAuthRulesBacktestReportCreated                           EventSubscriptionEventType = "auth_rules.backtest_report.created"
	EventSubscriptionEventTypeBalanceUpdated                                           EventSubscriptionEventType = "balance.updated"
	EventSubscriptionEventTypeBookTransferTransactionCreated                           EventSubscriptionEventType = "book_transfer_transaction.created"
	EventSubscriptionEventTypeBookTransferTransactionUpdated                           EventSubscriptionEventType = "book_transfer_transaction.updated"
	EventSubscriptionEventTypeCardTransactionEnhancedDataCreated                       EventSubscriptionEventType = "card_transaction.enhanced_data.created"
	EventSubscriptionEventTypeCardTransactionEnhancedDataUpdated                       EventSubscriptionEventType = "card_transaction.enhanced_data.updated"
	EventSubscriptionEventTypeCardTransactionUpdated                                   EventSubscriptionEventType = "card_transaction.updated"
	EventSubscriptionEventTypeCardConverted                                            EventSubscriptionEventType = "card.converted"
	EventSubscriptionEventTypeCardCreated                                              EventSubscriptionEventType = "card.created"
	EventSubscriptionEventTypeCardReissued                                             EventSubscriptionEventType = "card.reissued"
	EventSubscriptionEventTypeCardRenewed                                              EventSubscriptionEventType = "card.renewed"
	EventSubscriptionEventTypeCardShipped                                              EventSubscriptionEventType = "card.shipped"
	EventSubscriptionEventTypeDigitalWalletTokenizationApprovalRequest                 EventSubscriptionEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionEventTypeDigitalWalletTokenizationResult                          EventSubscriptionEventType = "digital_wallet.tokenization_result"
	EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventSubscriptionEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventSubscriptionEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventSubscriptionEventTypeDigitalWalletTokenizationUpdated                         EventSubscriptionEventType = "digital_wallet.tokenization_updated"
	EventSubscriptionEventTypeDisputeEvidenceUploadFailed                              EventSubscriptionEventType = "dispute_evidence.upload_failed"
	EventSubscriptionEventTypeDisputeTransactionCreated                                EventSubscriptionEventType = "dispute_transaction.created"
	EventSubscriptionEventTypeDisputeTransactionUpdated                                EventSubscriptionEventType = "dispute_transaction.updated"
	EventSubscriptionEventTypeDisputeUpdated                                           EventSubscriptionEventType = "dispute.updated"
	EventSubscriptionEventTypeExternalBankAccountCreated                               EventSubscriptionEventType = "external_bank_account.created"
	EventSubscriptionEventTypeExternalBankAccountUpdated                               EventSubscriptionEventType = "external_bank_account.updated"
	EventSubscriptionEventTypeExternalPaymentCreated                                   EventSubscriptionEventType = "external_payment.created"
	EventSubscriptionEventTypeExternalPaymentUpdated                                   EventSubscriptionEventType = "external_payment.updated"
	EventSubscriptionEventTypeFinancialAccountCreated                                  EventSubscriptionEventType = "financial_account.created"
	EventSubscriptionEventTypeFinancialAccountUpdated                                  EventSubscriptionEventType = "financial_account.updated"
	EventSubscriptionEventTypeFundingEventCreated                                      EventSubscriptionEventType = "funding_event.created"
	EventSubscriptionEventTypeInternalTransactionCreated                               EventSubscriptionEventType = "internal_transaction.created"
	EventSubscriptionEventTypeInternalTransactionUpdated                               EventSubscriptionEventType = "internal_transaction.updated"
	EventSubscriptionEventTypeLoanTapeCreated                                          EventSubscriptionEventType = "loan_tape.created"
	EventSubscriptionEventTypeLoanTapeUpdated                                          EventSubscriptionEventType = "loan_tape.updated"
	EventSubscriptionEventTypeManagementOperationCreated                               EventSubscriptionEventType = "management_operation.created"
	EventSubscriptionEventTypeManagementOperationUpdated                               EventSubscriptionEventType = "management_operation.updated"
	EventSubscriptionEventTypeNetworkTotalCreated                                      EventSubscriptionEventType = "network_total.created"
	EventSubscriptionEventTypeNetworkTotalUpdated                                      EventSubscriptionEventType = "network_total.updated"
	EventSubscriptionEventTypePaymentTransactionCreated                                EventSubscriptionEventType = "payment_transaction.created"
	EventSubscriptionEventTypePaymentTransactionUpdated                                EventSubscriptionEventType = "payment_transaction.updated"
	EventSubscriptionEventTypeSettlementReportUpdated                                  EventSubscriptionEventType = "settlement_report.updated"
	EventSubscriptionEventTypeStatementsCreated                                        EventSubscriptionEventType = "statements.created"
	EventSubscriptionEventTypeThreeDSAuthenticationChallenge                           EventSubscriptionEventType = "three_ds_authentication.challenge"
	EventSubscriptionEventTypeThreeDSAuthenticationCreated                             EventSubscriptionEventType = "three_ds_authentication.created"
	EventSubscriptionEventTypeThreeDSAuthenticationUpdated                             EventSubscriptionEventType = "three_ds_authentication.updated"
	EventSubscriptionEventTypeTokenizationApprovalRequest                              EventSubscriptionEventType = "tokenization.approval_request"
	EventSubscriptionEventTypeTokenizationResult                                       EventSubscriptionEventType = "tokenization.result"
	EventSubscriptionEventTypeTokenizationTwoFactorAuthenticationCode                  EventSubscriptionEventType = "tokenization.two_factor_authentication_code"
	EventSubscriptionEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventSubscriptionEventType = "tokenization.two_factor_authentication_code_sent"
	EventSubscriptionEventTypeTokenizationUpdated                                      EventSubscriptionEventType = "tokenization.updated"
)

func (r EventSubscriptionEventType) IsKnown() bool {
	switch r {
	case EventSubscriptionEventTypeAccountHolderDocumentUpdated, EventSubscriptionEventTypeAccountHolderCreated, EventSubscriptionEventTypeAccountHolderUpdated, EventSubscriptionEventTypeAccountHolderVerification, EventSubscriptionEventTypeAuthRulesBacktestReportCreated, EventSubscriptionEventTypeBalanceUpdated, EventSubscriptionEventTypeBookTransferTransactionCreated, EventSubscriptionEventTypeBookTransferTransactionUpdated, EventSubscriptionEventTypeCardTransactionEnhancedDataCreated, EventSubscriptionEventTypeCardTransactionEnhancedDataUpdated, EventSubscriptionEventTypeCardTransactionUpdated, EventSubscriptionEventTypeCardConverted, EventSubscriptionEventTypeCardCreated, EventSubscriptionEventTypeCardReissued, EventSubscriptionEventTypeCardRenewed, EventSubscriptionEventTypeCardShipped, EventSubscriptionEventTypeDigitalWalletTokenizationApprovalRequest, EventSubscriptionEventTypeDigitalWalletTokenizationResult, EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventSubscriptionEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionEventTypeDigitalWalletTokenizationUpdated, EventSubscriptionEventTypeDisputeEvidenceUploadFailed, EventSubscriptionEventTypeDisputeTransactionCreated, EventSubscriptionEventTypeDisputeTransactionUpdated, EventSubscriptionEventTypeDisputeUpdated, EventSubscriptionEventTypeExternalBankAccountCreated, EventSubscriptionEventTypeExternalBankAccountUpdated, EventSubscriptionEventTypeExternalPaymentCreated, EventSubscriptionEventTypeExternalPaymentUpdated, EventSubscriptionEventTypeFinancialAccountCreated, EventSubscriptionEventTypeFinancialAccountUpdated, EventSubscriptionEventTypeFundingEventCreated, EventSubscriptionEventTypeInternalTransactionCreated, EventSubscriptionEventTypeInternalTransactionUpdated, EventSubscriptionEventTypeLoanTapeCreated, EventSubscriptionEventTypeLoanTapeUpdated, EventSubscriptionEventTypeManagementOperationCreated, EventSubscriptionEventTypeManagementOperationUpdated, EventSubscriptionEventTypeNetworkTotalCreated, EventSubscriptionEventTypeNetworkTotalUpdated, EventSubscriptionEventTypePaymentTransactionCreated, EventSubscriptionEventTypePaymentTransactionUpdated, EventSubscriptionEventTypeSettlementReportUpdated, EventSubscriptionEventTypeStatementsCreated, EventSubscriptionEventTypeThreeDSAuthenticationChallenge, EventSubscriptionEventTypeThreeDSAuthenticationCreated, EventSubscriptionEventTypeThreeDSAuthenticationUpdated, EventSubscriptionEventTypeTokenizationApprovalRequest, EventSubscriptionEventTypeTokenizationResult, EventSubscriptionEventTypeTokenizationTwoFactorAuthenticationCode, EventSubscriptionEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionEventTypeTokenizationUpdated:
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
type EventListParamsEventType string

const (
	EventListParamsEventTypeAccountHolderDocumentUpdated                             EventListParamsEventType = "account_holder_document.updated"
	EventListParamsEventTypeAccountHolderCreated                                     EventListParamsEventType = "account_holder.created"
	EventListParamsEventTypeAccountHolderUpdated                                     EventListParamsEventType = "account_holder.updated"
	EventListParamsEventTypeAccountHolderVerification                                EventListParamsEventType = "account_holder.verification"
	EventListParamsEventTypeAuthRulesBacktestReportCreated                           EventListParamsEventType = "auth_rules.backtest_report.created"
	EventListParamsEventTypeBalanceUpdated                                           EventListParamsEventType = "balance.updated"
	EventListParamsEventTypeBookTransferTransactionCreated                           EventListParamsEventType = "book_transfer_transaction.created"
	EventListParamsEventTypeBookTransferTransactionUpdated                           EventListParamsEventType = "book_transfer_transaction.updated"
	EventListParamsEventTypeCardTransactionEnhancedDataCreated                       EventListParamsEventType = "card_transaction.enhanced_data.created"
	EventListParamsEventTypeCardTransactionEnhancedDataUpdated                       EventListParamsEventType = "card_transaction.enhanced_data.updated"
	EventListParamsEventTypeCardTransactionUpdated                                   EventListParamsEventType = "card_transaction.updated"
	EventListParamsEventTypeCardConverted                                            EventListParamsEventType = "card.converted"
	EventListParamsEventTypeCardCreated                                              EventListParamsEventType = "card.created"
	EventListParamsEventTypeCardReissued                                             EventListParamsEventType = "card.reissued"
	EventListParamsEventTypeCardRenewed                                              EventListParamsEventType = "card.renewed"
	EventListParamsEventTypeCardShipped                                              EventListParamsEventType = "card.shipped"
	EventListParamsEventTypeDigitalWalletTokenizationApprovalRequest                 EventListParamsEventType = "digital_wallet.tokenization_approval_request"
	EventListParamsEventTypeDigitalWalletTokenizationResult                          EventListParamsEventType = "digital_wallet.tokenization_result"
	EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventListParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventListParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventListParamsEventTypeDigitalWalletTokenizationUpdated                         EventListParamsEventType = "digital_wallet.tokenization_updated"
	EventListParamsEventTypeDisputeEvidenceUploadFailed                              EventListParamsEventType = "dispute_evidence.upload_failed"
	EventListParamsEventTypeDisputeTransactionCreated                                EventListParamsEventType = "dispute_transaction.created"
	EventListParamsEventTypeDisputeTransactionUpdated                                EventListParamsEventType = "dispute_transaction.updated"
	EventListParamsEventTypeDisputeUpdated                                           EventListParamsEventType = "dispute.updated"
	EventListParamsEventTypeExternalBankAccountCreated                               EventListParamsEventType = "external_bank_account.created"
	EventListParamsEventTypeExternalBankAccountUpdated                               EventListParamsEventType = "external_bank_account.updated"
	EventListParamsEventTypeExternalPaymentCreated                                   EventListParamsEventType = "external_payment.created"
	EventListParamsEventTypeExternalPaymentUpdated                                   EventListParamsEventType = "external_payment.updated"
	EventListParamsEventTypeFinancialAccountCreated                                  EventListParamsEventType = "financial_account.created"
	EventListParamsEventTypeFinancialAccountUpdated                                  EventListParamsEventType = "financial_account.updated"
	EventListParamsEventTypeFundingEventCreated                                      EventListParamsEventType = "funding_event.created"
	EventListParamsEventTypeInternalTransactionCreated                               EventListParamsEventType = "internal_transaction.created"
	EventListParamsEventTypeInternalTransactionUpdated                               EventListParamsEventType = "internal_transaction.updated"
	EventListParamsEventTypeLoanTapeCreated                                          EventListParamsEventType = "loan_tape.created"
	EventListParamsEventTypeLoanTapeUpdated                                          EventListParamsEventType = "loan_tape.updated"
	EventListParamsEventTypeManagementOperationCreated                               EventListParamsEventType = "management_operation.created"
	EventListParamsEventTypeManagementOperationUpdated                               EventListParamsEventType = "management_operation.updated"
	EventListParamsEventTypeNetworkTotalCreated                                      EventListParamsEventType = "network_total.created"
	EventListParamsEventTypeNetworkTotalUpdated                                      EventListParamsEventType = "network_total.updated"
	EventListParamsEventTypePaymentTransactionCreated                                EventListParamsEventType = "payment_transaction.created"
	EventListParamsEventTypePaymentTransactionUpdated                                EventListParamsEventType = "payment_transaction.updated"
	EventListParamsEventTypeSettlementReportUpdated                                  EventListParamsEventType = "settlement_report.updated"
	EventListParamsEventTypeStatementsCreated                                        EventListParamsEventType = "statements.created"
	EventListParamsEventTypeThreeDSAuthenticationChallenge                           EventListParamsEventType = "three_ds_authentication.challenge"
	EventListParamsEventTypeThreeDSAuthenticationCreated                             EventListParamsEventType = "three_ds_authentication.created"
	EventListParamsEventTypeThreeDSAuthenticationUpdated                             EventListParamsEventType = "three_ds_authentication.updated"
	EventListParamsEventTypeTokenizationApprovalRequest                              EventListParamsEventType = "tokenization.approval_request"
	EventListParamsEventTypeTokenizationResult                                       EventListParamsEventType = "tokenization.result"
	EventListParamsEventTypeTokenizationTwoFactorAuthenticationCode                  EventListParamsEventType = "tokenization.two_factor_authentication_code"
	EventListParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventListParamsEventType = "tokenization.two_factor_authentication_code_sent"
	EventListParamsEventTypeTokenizationUpdated                                      EventListParamsEventType = "tokenization.updated"
)

func (r EventListParamsEventType) IsKnown() bool {
	switch r {
	case EventListParamsEventTypeAccountHolderDocumentUpdated, EventListParamsEventTypeAccountHolderCreated, EventListParamsEventTypeAccountHolderUpdated, EventListParamsEventTypeAccountHolderVerification, EventListParamsEventTypeAuthRulesBacktestReportCreated, EventListParamsEventTypeBalanceUpdated, EventListParamsEventTypeBookTransferTransactionCreated, EventListParamsEventTypeBookTransferTransactionUpdated, EventListParamsEventTypeCardTransactionEnhancedDataCreated, EventListParamsEventTypeCardTransactionEnhancedDataUpdated, EventListParamsEventTypeCardTransactionUpdated, EventListParamsEventTypeCardConverted, EventListParamsEventTypeCardCreated, EventListParamsEventTypeCardReissued, EventListParamsEventTypeCardRenewed, EventListParamsEventTypeCardShipped, EventListParamsEventTypeDigitalWalletTokenizationApprovalRequest, EventListParamsEventTypeDigitalWalletTokenizationResult, EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventListParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventListParamsEventTypeDigitalWalletTokenizationUpdated, EventListParamsEventTypeDisputeEvidenceUploadFailed, EventListParamsEventTypeDisputeTransactionCreated, EventListParamsEventTypeDisputeTransactionUpdated, EventListParamsEventTypeDisputeUpdated, EventListParamsEventTypeExternalBankAccountCreated, EventListParamsEventTypeExternalBankAccountUpdated, EventListParamsEventTypeExternalPaymentCreated, EventListParamsEventTypeExternalPaymentUpdated, EventListParamsEventTypeFinancialAccountCreated, EventListParamsEventTypeFinancialAccountUpdated, EventListParamsEventTypeFundingEventCreated, EventListParamsEventTypeInternalTransactionCreated, EventListParamsEventTypeInternalTransactionUpdated, EventListParamsEventTypeLoanTapeCreated, EventListParamsEventTypeLoanTapeUpdated, EventListParamsEventTypeManagementOperationCreated, EventListParamsEventTypeManagementOperationUpdated, EventListParamsEventTypeNetworkTotalCreated, EventListParamsEventTypeNetworkTotalUpdated, EventListParamsEventTypePaymentTransactionCreated, EventListParamsEventTypePaymentTransactionUpdated, EventListParamsEventTypeSettlementReportUpdated, EventListParamsEventTypeStatementsCreated, EventListParamsEventTypeThreeDSAuthenticationChallenge, EventListParamsEventTypeThreeDSAuthenticationCreated, EventListParamsEventTypeThreeDSAuthenticationUpdated, EventListParamsEventTypeTokenizationApprovalRequest, EventListParamsEventTypeTokenizationResult, EventListParamsEventTypeTokenizationTwoFactorAuthenticationCode, EventListParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventListParamsEventTypeTokenizationUpdated:
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
