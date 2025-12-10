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

	payload := []byte(`{"event_type":"account_holder.created","token":"00000000-0000-0000-0000-000000000001","account_token":"00000000-0000-0000-0000-000000000001","created":"2019-12-27T18:11:19.117Z","required_documents":[{"entity_token":"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e","status_reasons":["string"],"valid_documents":["string"]}],"status":"ACCEPTED","status_reason":["string"]}`)

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

func TestWebhookParsed(t *testing.T) {
	client := lithic.NewClient(
		option.WithWebhookSecret("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My Lithic API Key"),
	)
	payload := []byte(`{"event_type":"account_holder.created","token":"00000000-0000-0000-0000-000000000001","account_token":"00000000-0000-0000-0000-000000000001","created":"2019-12-27T18:11:19.117Z","required_documents":[{"entity_token":"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e","status_reasons":["string"],"valid_documents":["string"]}],"status":"ACCEPTED","status_reason":["string"]}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Error("Failed to sign test webhook message")
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Error("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	event, err := client.Webhooks.Parse(payload, headers)
	if err != nil {
		t.Error("Failed to unwrap webhook:", err)
	}

	// Demonstrate type switching on AsUnion() to distinguish event types
	switch e := event.AsUnion().(type) {
	case lithic.AccountHolderCreatedWebhookEvent:
		if e.Token != "00000000-0000-0000-0000-000000000001" {
			t.Errorf("Expected token 00000000-0000-0000-0000-000000000001, got %s", e.Token)
		}
		if e.EventType != lithic.AccountHolderCreatedWebhookEventEventTypeAccountHolderCreated {
			t.Errorf("Expected event type account_holder.created, got %s", e.EventType)
		}
	case lithic.CardCreatedWebhookEvent:
		t.Error("Unexpected CardCreatedWebhookEvent")
	default:
		t.Errorf("Unexpected event type: %T", e)
	}
}

func TestWebhookCardCreatedParsed(t *testing.T) {
	client := lithic.NewClient(
		option.WithWebhookSecret("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My Lithic API Key"),
	)
	payload := []byte(`{"event_type":"card.created","card_token":"00000000-0000-0000-0000-000000000002"}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Fatal("Failed to create webhook signer")
	}
	msgID := "2"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Fatal("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	event, err := client.Webhooks.Parse(payload, headers)
	if err != nil {
		t.Fatal("Failed to unwrap webhook:", err)
	}

	// Demonstrate type switching on AsUnion() for a different event type
	switch e := event.AsUnion().(type) {
	case lithic.CardCreatedWebhookEvent:
		if e.CardToken != "00000000-0000-0000-0000-000000000002" {
			t.Errorf("Expected card token 00000000-0000-0000-0000-000000000002, got %s", e.CardToken)
		}
		if e.EventType != lithic.CardCreatedWebhookEventEventTypeCardCreated {
			t.Errorf("Expected event type card.created, got %s", e.EventType)
		}
	case lithic.AccountHolderCreatedWebhookEvent:
		t.Error("Unexpected AccountHolderCreatedWebhookEvent")
	default:
		t.Errorf("Unexpected event type: %T", e)
	}
}
