// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

func TestVerifySignature(t *testing.T) {
	secret := "whsec_c2VjcmV0Cg=="

	payload := []byte(`{"account_token":"00000000-0000-0000-0000-000000000002","card_token":"00000000-0000-0000-0000-000000000001","created":"2023-09-18T12:34:56Z","digital_wallet_token_metadata":{"payment_account_info":{"account_holder_data":{"phone_number":"+15555555555"},"pan_unique_reference":"pan_unique_ref_1234567890123456789012345678","payment_account_reference":"ref_1234567890123456789012","token_unique_reference":"token_unique_ref_1234567890123456789012345678"},"status":"Pending","payment_app_instance_id":"app_instance_123456789012345678901234567890","token_requestor_id":"12345678901","token_requestor_name":"APPLE_PAY"},"event_type":"digital_wallet.tokenization_approval_request","issuer_decision":"APPROVED","tokenization_channel":"DIGITAL_WALLET","tokenization_token":"tok_1234567890abcdef","wallet_decisioning_info":{"account_score":"100","device_score":"100","recommended_decision":"Decision1","recommendation_reasons":["Reason1"]},"customer_tokenization_decision":{"outcome":"APPROVED","responder_url":"https://example.com","latency":"100","response_code":"123456"},"device":{"imei":"123456789012345","ip_address":"1.1.1.1","location":"37.3860517/-122.0838511"},"rule_results":[{"auth_rule_token":"550e8400-e29b-41d4-a716-446655440003","explanation":"Account risk too high","name":"CustomerAccountRule","result":"DECLINED"}],"tokenization_decline_reasons":["ACCOUNT_SCORE_1"],"tokenization_source":"PUSH_PROVISION","tokenization_tfa_reasons":["WALLET_RECOMMENDED_TFA"]}`)

	wh, err := standardwebhooks.NewWebhook(secret)
	if err != nil {
		t.Fatalf("Failed to create webhook signer: %s", err.Error())
	}

	msgID := "msg_test"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Fatalf("Failed to sign test webhook message: %s", err.Error())
	}

	header := http.Header{}
	header.Set("webhook-id", msgID)
	header.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	header.Set("webhook-signature", sig)

	client := lithic.NewClient()
	err = client.Webhooks.VerifySignature(payload, header, secret, now)
	if err != nil {
		t.Fatalf("did not expect error %s", err.Error())
	}
}

func TestWebhookParsing(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		check   func(t *testing.T, event *lithic.ParsedWebhookEvent)
	}{
		{
			name:    "account_holder.created with RFC3339 timestamp",
			payload: `{"event_type":"account_holder.created","token":"00000000-0000-0000-0000-000000000001","account_token":"00000000-0000-0000-0000-000000000001","created":"2019-12-27T18:11:19.117Z","required_documents":[{"entity_token":"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e","status_reasons":["string"],"valid_documents":["string"]}],"status":"ACCEPTED","status_reason":["string"]}`,
			check: func(t *testing.T, event *lithic.ParsedWebhookEvent) {
				e, ok := event.AsUnion().(lithic.AccountHolderCreatedWebhookEvent)
				if !ok {
					t.Fatalf("Expected AccountHolderCreatedWebhookEvent, got %T", event.AsUnion())
				}
				if e.Token != "00000000-0000-0000-0000-000000000001" {
					t.Errorf("Expected token 00000000-0000-0000-0000-000000000001, got %s", e.Token)
				}
				if e.EventType != lithic.AccountHolderCreatedWebhookEventEventTypeAccountHolderCreated {
					t.Errorf("Expected event type account_holder.created, got %s", e.EventType)
				}
			},
		},
		{
			name:    "card.created",
			payload: `{"event_type":"card.created","card_token":"00000000-0000-0000-0000-000000000002"}`,
			check: func(t *testing.T, event *lithic.ParsedWebhookEvent) {
				e, ok := event.AsUnion().(lithic.CardCreatedWebhookEvent)
				if !ok {
					t.Fatalf("Expected CardCreatedWebhookEvent, got %T", event.AsUnion())
				}
				if e.CardToken != "00000000-0000-0000-0000-000000000002" {
					t.Errorf("Expected card token 00000000-0000-0000-0000-000000000002, got %s", e.CardToken)
				}
				if e.EventType != lithic.CardCreatedWebhookEventEventTypeCardCreated {
					t.Errorf("Expected event type card.created, got %s", e.EventType)
				}
			},
		},
		{
			name:    "account_holder.verification with space-separated timestamp",
			payload: `{"event_type":"account_holder.verification","token":"00000000-0000-0000-0000-000000000003","account_token":"00000000-0000-0000-0000-000000000004","created":"2025-12-09 16:19:40.228000+00:00","status":"ACCEPTED","status_reasons":["KYC_PASSED"]}`,
			check: func(t *testing.T, event *lithic.ParsedWebhookEvent) {
				e, ok := event.AsUnion().(lithic.AccountHolderVerificationWebhookEvent)
				if !ok {
					t.Fatalf("Expected AccountHolderVerificationWebhookEvent, got %T", event.AsUnion())
				}
				if e.Token != "00000000-0000-0000-0000-000000000003" {
					t.Errorf("Expected token 00000000-0000-0000-0000-000000000003, got %s", e.Token)
				}
				if e.AccountToken != "00000000-0000-0000-0000-000000000004" {
					t.Errorf("Expected account token 00000000-0000-0000-0000-000000000004, got %s", e.AccountToken)
				}
				if e.EventType != lithic.AccountHolderVerificationWebhookEventEventTypeAccountHolderVerification {
					t.Errorf("Expected event type account_holder.verification, got %s", e.EventType)
				}
				if e.Status != lithic.AccountHolderVerificationWebhookEventStatusAccepted {
					t.Errorf("Expected status ACCEPTED, got %s", e.Status)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := parseWebhookPayload(t, []byte(tt.payload))
			tt.check(t, event)
		})
	}
}

func parseWebhookPayload(t *testing.T, payload []byte) *lithic.ParsedWebhookEvent {
	t.Helper()

	secret := "whsec_c2VjcmV0Cg=="
	client := lithic.NewClient(
		option.WithWebhookSecret(secret),
		option.WithAPIKey("test-api-key"),
	)

	wh, err := standardwebhooks.NewWebhook(secret)
	if err != nil {
		t.Fatalf("Failed to create webhook signer: %v", err)
	}

	msgID := "msg_test"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Fatalf("Failed to sign webhook payload: %v", err)
	}

	headers := make(http.Header)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	headers.Set("webhook-signature", sig)

	event, err := client.Webhooks.Parse(payload, headers)
	if err != nil {
		t.Fatalf("Failed to parse webhook: %v", err)
	}

	return event
}
