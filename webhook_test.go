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
	_, err = client.Webhooks.Parsed(payload, headers)
	if err != nil {
		t.Error("Failed to unwrap webhook:", err)
	}
}
