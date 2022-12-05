package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lithic-com/lithic-go/options"
)

type WebhookService struct {
	Options []options.RequestOption
}

func NewWebhookService(opts ...options.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// Validates whether or not the webhook payload was sent by Lithic.
//
// An error will be raised if the webhook payload was not sent by Lithic.
func (r *WebhookService) VerifySignature(payload []byte, headers http.Header, secret string, now time.Time) (err error) {
	whsecret, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(secret, "whsec_"))
	if err != nil {
		return fmt.Errorf("invalid webhook secret: %s", err)
	}

	id := headers.Get("webhook-id")
	if len(id) == 0 {
		return errors.New("could not find webhook-id header")
	}
	sign := headers.Values("webhook-signature")
	if len(sign) == 0 {
		return errors.New("could not find webhook-signature header")
	}
	unixtime := headers.Get("webhook-timestamp")
	if len(unixtime) == 0 {
		return errors.New("could not find webhook-timestamp header")
	}

	timestamp, err := strconv.ParseInt(unixtime, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid signature headers: %s", err)
	}

	if timestamp < now.UnixMilli()-300 {
		return errors.New("webhook timestamp too old")
	}
	if timestamp > now.UnixMilli()+300 {
		return errors.New("webhook timestamp too new")
	}

	mac := hmac.New(sha256.New, whsecret)
	mac.Write([]byte(id))
	mac.Write([]byte("."))
	mac.Write([]byte(unixtime))
	mac.Write([]byte("."))
	mac.Write(payload)
	expected := mac.Sum(nil)

	for _, part := range sign {
		parts := strings.Split(part, ",")
		if len(parts) != 2 {
			continue
		}
		if parts[0] != "v1" {
			continue
		}
		signature, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			continue
		}
		if hmac.Equal(signature, expected) {
			return nil
		}
	}

	return errors.New("None of the given webhook signatures match the expected signature")

}
