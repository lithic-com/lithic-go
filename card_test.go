// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestCardNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.New(context.TODO(), lithic.CardNewParams{
		Type:             lithic.F(lithic.CardNewParamsTypeVirtual),
		AccountToken:     lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CardProgramToken: lithic.F("00000000-0000-0000-1000-000000000000"),
		Carrier: lithic.F(shared.CarrierParam{
			QrCodeURL: lithic.F("string"),
		}),
		DigitalCardArtToken: lithic.F("00000000-0000-0000-1000-000000000000"),
		ExpMonth:            lithic.F("06"),
		ExpYear:             lithic.F("2027"),
		Memo:                lithic.F("New Card"),
		Pin:                 lithic.F("string"),
		ProductID:           lithic.F("1"),
		ShippingAddress: lithic.F(shared.ShippingAddressParam{
			FirstName:   lithic.F("Michael"),
			LastName:    lithic.F("Bluth"),
			Line2Text:   lithic.F("The Bluth Company"),
			Address1:    lithic.F("5 Broad Street"),
			Address2:    lithic.F("Unit 25A"),
			City:        lithic.F("NEW YORK"),
			State:       lithic.F("NY"),
			PostalCode:  lithic.F("10001-1809"),
			Country:     lithic.F("USA"),
			Email:       lithic.F("johnny@appleseed.com"),
			PhoneNumber: lithic.F("+12124007676"),
		}),
		ShippingMethod:     lithic.F(lithic.CardNewParamsShippingMethodStandard),
		SpendLimit:         lithic.F(int64(0)),
		SpendLimitDuration: lithic.F(lithic.SpendLimitDurationAnnually),
		State:              lithic.F(lithic.CardNewParamsStateOpen),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.CardUpdateParams{
			AuthRuleToken:       lithic.F("string"),
			DigitalCardArtToken: lithic.F("00000000-0000-0000-1000-000000000000"),
			Memo:                lithic.F("New Card"),
			Pin:                 lithic.F("string"),
			SpendLimit:          lithic.F(int64(0)),
			SpendLimitDuration:  lithic.F(lithic.SpendLimitDurationAnnually),
			State:               lithic.F(lithic.CardUpdateParamsStateClosed),
		},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.List(context.TODO(), lithic.CardListParams{
		AccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Begin:        lithic.F(time.Now()),
		End:          lithic.F(time.Now()),
		Page:         lithic.F(int64(0)),
		PageSize:     lithic.F(int64(1)),
		State:        lithic.F(lithic.CardListParamsStateOpen),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardEmbed(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.Embed(context.TODO(), lithic.CardEmbedParams{
		EmbedRequest: lithic.F("string"),
		Hmac:         lithic.F("string"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGetEmbedHTMLWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.GetEmbedHTML(context.TODO(), lithic.CardGetEmbedHTMLParams{
		Token:        lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Css:          lithic.F("string"),
		Expiration:   lithic.F(time.Now()),
		TargetOrigin: lithic.F("string"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGetEmbedURLWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.GetEmbedURL(context.TODO(), lithic.CardGetEmbedURLParams{
		Token:        lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Css:          lithic.F("string"),
		Expiration:   lithic.F(time.Now()),
		TargetOrigin: lithic.F("string"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProvisionWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.Provision(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.CardProvisionParams{
			Certificate:    lithic.F("U3RhaW5sZXNzIHJvY2tz"),
			DigitalWallet:  lithic.F(lithic.CardProvisionParamsDigitalWalletApplePay),
			Nonce:          lithic.F("U3RhaW5sZXNzIHJvY2tz"),
			NonceSignature: lithic.F("U3RhaW5sZXNzIHJvY2tz"),
		},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardReissueWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.Reissue(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.CardReissueParams{
			Carrier: lithic.F(shared.CarrierParam{
				QrCodeURL: lithic.F("string"),
			}),
			ProductID: lithic.F("string"),
			ShippingAddress: lithic.F(shared.ShippingAddressParam{
				FirstName:   lithic.F("Michael"),
				LastName:    lithic.F("Bluth"),
				Line2Text:   lithic.F("The Bluth Company"),
				Address1:    lithic.F("5 Broad Street"),
				Address2:    lithic.F("Unit 25A"),
				City:        lithic.F("NEW YORK"),
				State:       lithic.F("NY"),
				PostalCode:  lithic.F("10001-1809"),
				Country:     lithic.F("USA"),
				Email:       lithic.F("johnny@appleseed.com"),
				PhoneNumber: lithic.F("+12124007676"),
			}),
			ShippingMethod: lithic.F(lithic.CardReissueParamsShippingMethodStandard),
		},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
