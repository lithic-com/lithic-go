package lithic_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
)

func TestVerifySignature(t *testing.T) {
	secret := "whsec_zlFsbBZ8Xcodlpcu6NDTdSzZRLSdhkst"

	payload := `{"card_token":"sit Lorem ipsum, accusantium repellendus possimus","created_at":"elit. placeat libero architecto molestias, sit","account_token":"elit.","issuer_decision":"magnam, libero esse Lorem ipsum magnam, magnam,","tokenization_attempt_id":"illum dolor repellendus libero esse accusantium","wallet_decisioning_info":{"device_score":"placeat architecto"},"digital_wallet_token_metadata":{"status":"reprehenderit dolor","token_requestor_id":"possimus","payment_account_info":{"account_holder_data":{"phone_number":"libero","email_address":"nobis molestias, veniam culpa! quas elit. quas libero esse architecto placeat"},"pan_unique_reference":"adipisicing odit magnam, odit"}}}`

	header := http.Header{}
	header.Add("webhook-id", "msg_2Lh9KRb0pzN4LePd3XiA4v12Axj")
	header.Add("webhook-timestamp", "1676312382")
	header.Add("webhook-signature", "v1,Dwa0AHInLL3XFo2sxcHamOQDrJNi7F654S3L6skMAOI=")

	client := lithic.NewClient()
	err := client.Webhooks.VerifySignature([]byte(payload), header, secret, time.Unix(1676312382, 0))
	if err != nil {
		t.Fatalf("did not expect error %s", err.Error())
	}
}
