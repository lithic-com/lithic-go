package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/fields"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
)

func TestCardsNewWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.New(context.TODO(), &requests.CardNewParams{AccountToken: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), CardProgramToken: fields.F("00000000-0000-0000-1000-000000000000"), ExpMonth: fields.F("06"), ExpYear: fields.F("2027"), FundingToken: fields.F("ecbd1d58-0299-48b3-84da-6ed7f5bf9ec1"), Memo: fields.F("New Card"), SpendLimit: fields.F(int64(0)), SpendLimitDuration: fields.F(requests.SpendLimitDurationAnnually), State: fields.F(requests.CardNewParamsStateOpen), Type: fields.F(requests.CardNewParamsTypeVirtual), Pin: fields.F("string"), DigitalCardArtToken: fields.F("00000000-0000-0000-1000-000000000000"), ProductID: fields.F("1"), ShippingAddress: fields.F(requests.ShippingAddress{FirstName: fields.F("Michael"), LastName: fields.F("Bluth"), Line2Text: fields.F("The Bluth Company"), Address1: fields.F("5 Broad Street"), Address2: fields.F("Unit 25A"), City: fields.F("NEW YORK"), State: fields.F("NY"), PostalCode: fields.F("10001-1809"), Country: fields.F("USA"), Email: fields.F("johnny@appleseed.com"), PhoneNumber: fields.F("+12124007676")}), ShippingMethod: fields.F(requests.CardNewParamsShippingMethodStandard)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardsGet(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardsUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.CardUpdateParams{FundingToken: fields.F("ecbd1d58-0299-48b3-84da-6ed7f5bf9ec1"), Memo: fields.F("New Card"), SpendLimit: fields.F(int64(0)), SpendLimitDuration: fields.F(requests.SpendLimitDurationAnnually), AuthRuleToken: fields.F("string"), State: fields.F(requests.CardUpdateParamsStateClosed), Pin: fields.F("string"), DigitalCardArtToken: fields.F("00000000-0000-0000-1000-000000000000")},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardsListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	res, err := c.Cards.List(context.TODO(), &requests.CardListParams{AccountToken: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Begin: fields.F(time.Now()), End: fields.F(time.Now()), Page: fields.F(int64(0)), PageSize: fields.F(int64(1))})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}

	for res.Next() {
		items := res.Current()
		if items == nil || len(*items) == 0 {
			t.Fatalf("there should be at least one item")
		}
	}
}

func TestCardsEmbedWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.Embed(context.TODO(), &requests.CardEmbedParams{EmbedRequest: fields.F("string"), Hmac: fields.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardsGetEmbedHTMLWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.GetEmbedHTML(context.TODO(), &requests.EmbedRequestParams{Css: fields.F("string"), Expiration: fields.F(time.Now()), Token: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), TargetOrigin: fields.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardsGetEmbedURLWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.GetEmbedURL(context.TODO(), &requests.EmbedRequestParams{Css: fields.F("string"), Expiration: fields.F(time.Now()), Token: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), TargetOrigin: fields.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardsProvisionWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.Provision(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.CardProvisionParams{DigitalWallet: fields.F(requests.CardProvisionParamsDigitalWalletApplePay), Nonce: fields.F("U3RhaW5sZXNzIHJvY2tz"), NonceSignature: fields.F("U3RhaW5sZXNzIHJvY2tz"), Certificate: fields.F("U3RhaW5sZXNzIHJvY2tz")},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardsReissueWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.Reissue(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.CardReissueParams{ShippingAddress: fields.F(requests.ShippingAddress{FirstName: fields.F("Michael"), LastName: fields.F("Bluth"), Line2Text: fields.F("The Bluth Company"), Address1: fields.F("5 Broad Street"), Address2: fields.F("Unit 25A"), City: fields.F("NEW YORK"), State: fields.F("NY"), PostalCode: fields.F("10001-1809"), Country: fields.F("USA"), Email: fields.F("johnny@appleseed.com"), PhoneNumber: fields.F("+12124007676")}), ShippingMethod: fields.F(requests.CardReissueParamsShippingMethodStandard), ProductID: fields.F("string")},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
