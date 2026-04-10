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
	payload := []byte(`{"account_token":"00000000-0000-0000-0000-000000000002","card_token":"00000000-0000-0000-0000-000000000001","created":"2023-09-18T12:34:56Z","digital_wallet_token_metadata":{"payment_account_info":{"account_holder_data":{"phone_number":"+15555555555"},"pan_unique_reference":"pan_unique_ref_1234567890123456789012345678","payment_account_reference":"ref_1234567890123456789012","token_unique_reference":"token_unique_ref_1234567890123456789012345678"},"status":"Pending","payment_app_instance_id":"app_instance_123456789012345678901234567890","token_requestor_id":"12345678901","token_requestor_name":"APPLE_PAY"},"event_type":"digital_wallet.tokenization_approval_request","issuer_decision":"APPROVED","tokenization_channel":"DIGITAL_WALLET","tokenization_token":"tok_1234567890abcdef","wallet_decisioning_info":{"account_score":"100","device_score":"100","recommended_decision":"Decision1","recommendation_reasons":["Reason1"]},"customer_tokenization_decision":{"outcome":"APPROVED","responder_url":"https://example.com","latency":"100","response_code":"123456"},"device":{"imei":"123456789012345","ip_address":"1.1.1.1","location":"37.3860517/-122.0838511"},"rule_results":[{"auth_rule_token":"550e8400-e29b-41d4-a716-446655440003","explanation":"Account risk too high","name":"CustomerAccountRule","result":"DECLINED"}],"tokenization_decline_reasons":["ACCOUNT_SCORE_1"],"tokenization_source":"PUSH_PROVISION","tokenization_tfa_reasons":["WALLET_RECOMMENDED_TFA"]}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Fatal("Failed to sign test webhook message", err)
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Fatal("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Webhooks.Parsed(payload, headers)
	if err != nil {
		t.Fatal("Failed to unwrap webhook:", err)
	}
}
