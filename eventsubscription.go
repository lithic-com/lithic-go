// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	opts = append(r.Options[:], opts...)
	path := "v1/event_subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an event subscription.
func (r *EventSubscriptionService) Get(ctx context.Context, eventSubscriptionToken string, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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

type EventSubscriptionNewParamsEventType string

const (
	EventSubscriptionNewParamsEventTypeAccountHolderCreated                                     EventSubscriptionNewParamsEventType = "account_holder.created"
	EventSubscriptionNewParamsEventTypeAccountHolderUpdated                                     EventSubscriptionNewParamsEventType = "account_holder.updated"
	EventSubscriptionNewParamsEventTypeAccountHolderVerification                                EventSubscriptionNewParamsEventType = "account_holder.verification"
	EventSubscriptionNewParamsEventTypeAuthRulesPerformanceReportCreated                        EventSubscriptionNewParamsEventType = "auth_rules.performance_report.created"
	EventSubscriptionNewParamsEventTypeBalanceUpdated                                           EventSubscriptionNewParamsEventType = "balance.updated"
	EventSubscriptionNewParamsEventTypeBookTransferTransactionCreated                           EventSubscriptionNewParamsEventType = "book_transfer_transaction.created"
	EventSubscriptionNewParamsEventTypeCardCreated                                              EventSubscriptionNewParamsEventType = "card.created"
	EventSubscriptionNewParamsEventTypeCardRenewed                                              EventSubscriptionNewParamsEventType = "card.renewed"
	EventSubscriptionNewParamsEventTypeCardReissued                                             EventSubscriptionNewParamsEventType = "card.reissued"
	EventSubscriptionNewParamsEventTypeCardConverted                                            EventSubscriptionNewParamsEventType = "card.converted"
	EventSubscriptionNewParamsEventTypeCardShipped                                              EventSubscriptionNewParamsEventType = "card.shipped"
	EventSubscriptionNewParamsEventTypeCardTransactionUpdated                                   EventSubscriptionNewParamsEventType = "card_transaction.updated"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationApprovalRequest                 EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationResult                          EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_result"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationUpdated                         EventSubscriptionNewParamsEventType = "digital_wallet.tokenization_updated"
	EventSubscriptionNewParamsEventTypeDisputeUpdated                                           EventSubscriptionNewParamsEventType = "dispute.updated"
	EventSubscriptionNewParamsEventTypeDisputeEvidenceUploadFailed                              EventSubscriptionNewParamsEventType = "dispute_evidence.upload_failed"
	EventSubscriptionNewParamsEventTypeExternalBankAccountCreated                               EventSubscriptionNewParamsEventType = "external_bank_account.created"
	EventSubscriptionNewParamsEventTypeExternalBankAccountUpdated                               EventSubscriptionNewParamsEventType = "external_bank_account.updated"
	EventSubscriptionNewParamsEventTypeExternalPaymentCreated                                   EventSubscriptionNewParamsEventType = "external_payment.created"
	EventSubscriptionNewParamsEventTypeExternalPaymentUpdated                                   EventSubscriptionNewParamsEventType = "external_payment.updated"
	EventSubscriptionNewParamsEventTypeFinancialAccountCreated                                  EventSubscriptionNewParamsEventType = "financial_account.created"
	EventSubscriptionNewParamsEventTypeFinancialAccountUpdated                                  EventSubscriptionNewParamsEventType = "financial_account.updated"
	EventSubscriptionNewParamsEventTypeLoanTapeCreated                                          EventSubscriptionNewParamsEventType = "loan_tape.created"
	EventSubscriptionNewParamsEventTypeLoanTapeUpdated                                          EventSubscriptionNewParamsEventType = "loan_tape.updated"
	EventSubscriptionNewParamsEventTypeManagementOperationCreated                               EventSubscriptionNewParamsEventType = "management_operation.created"
	EventSubscriptionNewParamsEventTypeManagementOperationUpdated                               EventSubscriptionNewParamsEventType = "management_operation.updated"
	EventSubscriptionNewParamsEventTypePaymentTransactionCreated                                EventSubscriptionNewParamsEventType = "payment_transaction.created"
	EventSubscriptionNewParamsEventTypePaymentTransactionUpdated                                EventSubscriptionNewParamsEventType = "payment_transaction.updated"
	EventSubscriptionNewParamsEventTypeSettlementReportUpdated                                  EventSubscriptionNewParamsEventType = "settlement_report.updated"
	EventSubscriptionNewParamsEventTypeStatementsCreated                                        EventSubscriptionNewParamsEventType = "statements.created"
	EventSubscriptionNewParamsEventTypeThreeDSAuthenticationCreated                             EventSubscriptionNewParamsEventType = "three_ds_authentication.created"
	EventSubscriptionNewParamsEventTypeTokenizationApprovalRequest                              EventSubscriptionNewParamsEventType = "tokenization.approval_request"
	EventSubscriptionNewParamsEventTypeTokenizationResult                                       EventSubscriptionNewParamsEventType = "tokenization.result"
	EventSubscriptionNewParamsEventTypeTokenizationTwoFactorAuthenticationCode                  EventSubscriptionNewParamsEventType = "tokenization.two_factor_authentication_code"
	EventSubscriptionNewParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventSubscriptionNewParamsEventType = "tokenization.two_factor_authentication_code_sent"
	EventSubscriptionNewParamsEventTypeTokenizationUpdated                                      EventSubscriptionNewParamsEventType = "tokenization.updated"
)

func (r EventSubscriptionNewParamsEventType) IsKnown() bool {
	switch r {
	case EventSubscriptionNewParamsEventTypeAccountHolderCreated, EventSubscriptionNewParamsEventTypeAccountHolderUpdated, EventSubscriptionNewParamsEventTypeAccountHolderVerification, EventSubscriptionNewParamsEventTypeAuthRulesPerformanceReportCreated, EventSubscriptionNewParamsEventTypeBalanceUpdated, EventSubscriptionNewParamsEventTypeBookTransferTransactionCreated, EventSubscriptionNewParamsEventTypeCardCreated, EventSubscriptionNewParamsEventTypeCardRenewed, EventSubscriptionNewParamsEventTypeCardReissued, EventSubscriptionNewParamsEventTypeCardConverted, EventSubscriptionNewParamsEventTypeCardShipped, EventSubscriptionNewParamsEventTypeCardTransactionUpdated, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationApprovalRequest, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationResult, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionNewParamsEventTypeDigitalWalletTokenizationUpdated, EventSubscriptionNewParamsEventTypeDisputeUpdated, EventSubscriptionNewParamsEventTypeDisputeEvidenceUploadFailed, EventSubscriptionNewParamsEventTypeExternalBankAccountCreated, EventSubscriptionNewParamsEventTypeExternalBankAccountUpdated, EventSubscriptionNewParamsEventTypeExternalPaymentCreated, EventSubscriptionNewParamsEventTypeExternalPaymentUpdated, EventSubscriptionNewParamsEventTypeFinancialAccountCreated, EventSubscriptionNewParamsEventTypeFinancialAccountUpdated, EventSubscriptionNewParamsEventTypeLoanTapeCreated, EventSubscriptionNewParamsEventTypeLoanTapeUpdated, EventSubscriptionNewParamsEventTypeManagementOperationCreated, EventSubscriptionNewParamsEventTypeManagementOperationUpdated, EventSubscriptionNewParamsEventTypePaymentTransactionCreated, EventSubscriptionNewParamsEventTypePaymentTransactionUpdated, EventSubscriptionNewParamsEventTypeSettlementReportUpdated, EventSubscriptionNewParamsEventTypeStatementsCreated, EventSubscriptionNewParamsEventTypeThreeDSAuthenticationCreated, EventSubscriptionNewParamsEventTypeTokenizationApprovalRequest, EventSubscriptionNewParamsEventTypeTokenizationResult, EventSubscriptionNewParamsEventTypeTokenizationTwoFactorAuthenticationCode, EventSubscriptionNewParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionNewParamsEventTypeTokenizationUpdated:
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

type EventSubscriptionUpdateParamsEventType string

const (
	EventSubscriptionUpdateParamsEventTypeAccountHolderCreated                                     EventSubscriptionUpdateParamsEventType = "account_holder.created"
	EventSubscriptionUpdateParamsEventTypeAccountHolderUpdated                                     EventSubscriptionUpdateParamsEventType = "account_holder.updated"
	EventSubscriptionUpdateParamsEventTypeAccountHolderVerification                                EventSubscriptionUpdateParamsEventType = "account_holder.verification"
	EventSubscriptionUpdateParamsEventTypeAuthRulesPerformanceReportCreated                        EventSubscriptionUpdateParamsEventType = "auth_rules.performance_report.created"
	EventSubscriptionUpdateParamsEventTypeBalanceUpdated                                           EventSubscriptionUpdateParamsEventType = "balance.updated"
	EventSubscriptionUpdateParamsEventTypeBookTransferTransactionCreated                           EventSubscriptionUpdateParamsEventType = "book_transfer_transaction.created"
	EventSubscriptionUpdateParamsEventTypeCardCreated                                              EventSubscriptionUpdateParamsEventType = "card.created"
	EventSubscriptionUpdateParamsEventTypeCardRenewed                                              EventSubscriptionUpdateParamsEventType = "card.renewed"
	EventSubscriptionUpdateParamsEventTypeCardReissued                                             EventSubscriptionUpdateParamsEventType = "card.reissued"
	EventSubscriptionUpdateParamsEventTypeCardConverted                                            EventSubscriptionUpdateParamsEventType = "card.converted"
	EventSubscriptionUpdateParamsEventTypeCardShipped                                              EventSubscriptionUpdateParamsEventType = "card.shipped"
	EventSubscriptionUpdateParamsEventTypeCardTransactionUpdated                                   EventSubscriptionUpdateParamsEventType = "card_transaction.updated"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationApprovalRequest                 EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationResult                          EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_result"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationUpdated                         EventSubscriptionUpdateParamsEventType = "digital_wallet.tokenization_updated"
	EventSubscriptionUpdateParamsEventTypeDisputeUpdated                                           EventSubscriptionUpdateParamsEventType = "dispute.updated"
	EventSubscriptionUpdateParamsEventTypeDisputeEvidenceUploadFailed                              EventSubscriptionUpdateParamsEventType = "dispute_evidence.upload_failed"
	EventSubscriptionUpdateParamsEventTypeExternalBankAccountCreated                               EventSubscriptionUpdateParamsEventType = "external_bank_account.created"
	EventSubscriptionUpdateParamsEventTypeExternalBankAccountUpdated                               EventSubscriptionUpdateParamsEventType = "external_bank_account.updated"
	EventSubscriptionUpdateParamsEventTypeExternalPaymentCreated                                   EventSubscriptionUpdateParamsEventType = "external_payment.created"
	EventSubscriptionUpdateParamsEventTypeExternalPaymentUpdated                                   EventSubscriptionUpdateParamsEventType = "external_payment.updated"
	EventSubscriptionUpdateParamsEventTypeFinancialAccountCreated                                  EventSubscriptionUpdateParamsEventType = "financial_account.created"
	EventSubscriptionUpdateParamsEventTypeFinancialAccountUpdated                                  EventSubscriptionUpdateParamsEventType = "financial_account.updated"
	EventSubscriptionUpdateParamsEventTypeLoanTapeCreated                                          EventSubscriptionUpdateParamsEventType = "loan_tape.created"
	EventSubscriptionUpdateParamsEventTypeLoanTapeUpdated                                          EventSubscriptionUpdateParamsEventType = "loan_tape.updated"
	EventSubscriptionUpdateParamsEventTypeManagementOperationCreated                               EventSubscriptionUpdateParamsEventType = "management_operation.created"
	EventSubscriptionUpdateParamsEventTypeManagementOperationUpdated                               EventSubscriptionUpdateParamsEventType = "management_operation.updated"
	EventSubscriptionUpdateParamsEventTypePaymentTransactionCreated                                EventSubscriptionUpdateParamsEventType = "payment_transaction.created"
	EventSubscriptionUpdateParamsEventTypePaymentTransactionUpdated                                EventSubscriptionUpdateParamsEventType = "payment_transaction.updated"
	EventSubscriptionUpdateParamsEventTypeSettlementReportUpdated                                  EventSubscriptionUpdateParamsEventType = "settlement_report.updated"
	EventSubscriptionUpdateParamsEventTypeStatementsCreated                                        EventSubscriptionUpdateParamsEventType = "statements.created"
	EventSubscriptionUpdateParamsEventTypeThreeDSAuthenticationCreated                             EventSubscriptionUpdateParamsEventType = "three_ds_authentication.created"
	EventSubscriptionUpdateParamsEventTypeTokenizationApprovalRequest                              EventSubscriptionUpdateParamsEventType = "tokenization.approval_request"
	EventSubscriptionUpdateParamsEventTypeTokenizationResult                                       EventSubscriptionUpdateParamsEventType = "tokenization.result"
	EventSubscriptionUpdateParamsEventTypeTokenizationTwoFactorAuthenticationCode                  EventSubscriptionUpdateParamsEventType = "tokenization.two_factor_authentication_code"
	EventSubscriptionUpdateParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventSubscriptionUpdateParamsEventType = "tokenization.two_factor_authentication_code_sent"
	EventSubscriptionUpdateParamsEventTypeTokenizationUpdated                                      EventSubscriptionUpdateParamsEventType = "tokenization.updated"
)

func (r EventSubscriptionUpdateParamsEventType) IsKnown() bool {
	switch r {
	case EventSubscriptionUpdateParamsEventTypeAccountHolderCreated, EventSubscriptionUpdateParamsEventTypeAccountHolderUpdated, EventSubscriptionUpdateParamsEventTypeAccountHolderVerification, EventSubscriptionUpdateParamsEventTypeAuthRulesPerformanceReportCreated, EventSubscriptionUpdateParamsEventTypeBalanceUpdated, EventSubscriptionUpdateParamsEventTypeBookTransferTransactionCreated, EventSubscriptionUpdateParamsEventTypeCardCreated, EventSubscriptionUpdateParamsEventTypeCardRenewed, EventSubscriptionUpdateParamsEventTypeCardReissued, EventSubscriptionUpdateParamsEventTypeCardConverted, EventSubscriptionUpdateParamsEventTypeCardShipped, EventSubscriptionUpdateParamsEventTypeCardTransactionUpdated, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationApprovalRequest, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationResult, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionUpdateParamsEventTypeDigitalWalletTokenizationUpdated, EventSubscriptionUpdateParamsEventTypeDisputeUpdated, EventSubscriptionUpdateParamsEventTypeDisputeEvidenceUploadFailed, EventSubscriptionUpdateParamsEventTypeExternalBankAccountCreated, EventSubscriptionUpdateParamsEventTypeExternalBankAccountUpdated, EventSubscriptionUpdateParamsEventTypeExternalPaymentCreated, EventSubscriptionUpdateParamsEventTypeExternalPaymentUpdated, EventSubscriptionUpdateParamsEventTypeFinancialAccountCreated, EventSubscriptionUpdateParamsEventTypeFinancialAccountUpdated, EventSubscriptionUpdateParamsEventTypeLoanTapeCreated, EventSubscriptionUpdateParamsEventTypeLoanTapeUpdated, EventSubscriptionUpdateParamsEventTypeManagementOperationCreated, EventSubscriptionUpdateParamsEventTypeManagementOperationUpdated, EventSubscriptionUpdateParamsEventTypePaymentTransactionCreated, EventSubscriptionUpdateParamsEventTypePaymentTransactionUpdated, EventSubscriptionUpdateParamsEventTypeSettlementReportUpdated, EventSubscriptionUpdateParamsEventTypeStatementsCreated, EventSubscriptionUpdateParamsEventTypeThreeDSAuthenticationCreated, EventSubscriptionUpdateParamsEventTypeTokenizationApprovalRequest, EventSubscriptionUpdateParamsEventTypeTokenizationResult, EventSubscriptionUpdateParamsEventTypeTokenizationTwoFactorAuthenticationCode, EventSubscriptionUpdateParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionUpdateParamsEventTypeTokenizationUpdated:
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
	EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderCreated                                     EventSubscriptionSendSimulatedExampleParamsEventType = "account_holder.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderUpdated                                     EventSubscriptionSendSimulatedExampleParamsEventType = "account_holder.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderVerification                                EventSubscriptionSendSimulatedExampleParamsEventType = "account_holder.verification"
	EventSubscriptionSendSimulatedExampleParamsEventTypeAuthRulesPerformanceReportCreated                        EventSubscriptionSendSimulatedExampleParamsEventType = "auth_rules.performance_report.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeBalanceUpdated                                           EventSubscriptionSendSimulatedExampleParamsEventType = "balance.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeBookTransferTransactionCreated                           EventSubscriptionSendSimulatedExampleParamsEventType = "book_transfer_transaction.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardCreated                                              EventSubscriptionSendSimulatedExampleParamsEventType = "card.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardRenewed                                              EventSubscriptionSendSimulatedExampleParamsEventType = "card.renewed"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardReissued                                             EventSubscriptionSendSimulatedExampleParamsEventType = "card.reissued"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardConverted                                            EventSubscriptionSendSimulatedExampleParamsEventType = "card.converted"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardShipped                                              EventSubscriptionSendSimulatedExampleParamsEventType = "card.shipped"
	EventSubscriptionSendSimulatedExampleParamsEventTypeCardTransactionUpdated                                   EventSubscriptionSendSimulatedExampleParamsEventType = "card_transaction.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationApprovalRequest                 EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_approval_request"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationResult                          EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_result"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode     EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_two_factor_authentication_code_sent"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationUpdated                         EventSubscriptionSendSimulatedExampleParamsEventType = "digital_wallet.tokenization_updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeUpdated                                           EventSubscriptionSendSimulatedExampleParamsEventType = "dispute.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeEvidenceUploadFailed                              EventSubscriptionSendSimulatedExampleParamsEventType = "dispute_evidence.upload_failed"
	EventSubscriptionSendSimulatedExampleParamsEventTypeExternalBankAccountCreated                               EventSubscriptionSendSimulatedExampleParamsEventType = "external_bank_account.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeExternalBankAccountUpdated                               EventSubscriptionSendSimulatedExampleParamsEventType = "external_bank_account.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeExternalPaymentCreated                                   EventSubscriptionSendSimulatedExampleParamsEventType = "external_payment.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeExternalPaymentUpdated                                   EventSubscriptionSendSimulatedExampleParamsEventType = "external_payment.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeFinancialAccountCreated                                  EventSubscriptionSendSimulatedExampleParamsEventType = "financial_account.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeFinancialAccountUpdated                                  EventSubscriptionSendSimulatedExampleParamsEventType = "financial_account.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeLoanTapeCreated                                          EventSubscriptionSendSimulatedExampleParamsEventType = "loan_tape.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeLoanTapeUpdated                                          EventSubscriptionSendSimulatedExampleParamsEventType = "loan_tape.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeManagementOperationCreated                               EventSubscriptionSendSimulatedExampleParamsEventType = "management_operation.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeManagementOperationUpdated                               EventSubscriptionSendSimulatedExampleParamsEventType = "management_operation.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypePaymentTransactionCreated                                EventSubscriptionSendSimulatedExampleParamsEventType = "payment_transaction.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypePaymentTransactionUpdated                                EventSubscriptionSendSimulatedExampleParamsEventType = "payment_transaction.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeSettlementReportUpdated                                  EventSubscriptionSendSimulatedExampleParamsEventType = "settlement_report.updated"
	EventSubscriptionSendSimulatedExampleParamsEventTypeStatementsCreated                                        EventSubscriptionSendSimulatedExampleParamsEventType = "statements.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeThreeDSAuthenticationCreated                             EventSubscriptionSendSimulatedExampleParamsEventType = "three_ds_authentication.created"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationApprovalRequest                              EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.approval_request"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationResult                                       EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.result"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationTwoFactorAuthenticationCode                  EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.two_factor_authentication_code"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent              EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.two_factor_authentication_code_sent"
	EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationUpdated                                      EventSubscriptionSendSimulatedExampleParamsEventType = "tokenization.updated"
)

func (r EventSubscriptionSendSimulatedExampleParamsEventType) IsKnown() bool {
	switch r {
	case EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeAccountHolderVerification, EventSubscriptionSendSimulatedExampleParamsEventTypeAuthRulesPerformanceReportCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeBalanceUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeBookTransferTransactionCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeCardCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeCardRenewed, EventSubscriptionSendSimulatedExampleParamsEventTypeCardReissued, EventSubscriptionSendSimulatedExampleParamsEventTypeCardConverted, EventSubscriptionSendSimulatedExampleParamsEventTypeCardShipped, EventSubscriptionSendSimulatedExampleParamsEventTypeCardTransactionUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationApprovalRequest, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationResult, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCode, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionSendSimulatedExampleParamsEventTypeDigitalWalletTokenizationUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeDisputeEvidenceUploadFailed, EventSubscriptionSendSimulatedExampleParamsEventTypeExternalBankAccountCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeExternalBankAccountUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeExternalPaymentCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeExternalPaymentUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeFinancialAccountCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeFinancialAccountUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeLoanTapeCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeLoanTapeUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeManagementOperationCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeManagementOperationUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypePaymentTransactionCreated, EventSubscriptionSendSimulatedExampleParamsEventTypePaymentTransactionUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeSettlementReportUpdated, EventSubscriptionSendSimulatedExampleParamsEventTypeStatementsCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeThreeDSAuthenticationCreated, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationApprovalRequest, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationResult, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationTwoFactorAuthenticationCode, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationTwoFactorAuthenticationCodeSent, EventSubscriptionSendSimulatedExampleParamsEventTypeTokenizationUpdated:
		return true
	}
	return false
}
