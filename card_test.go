// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/shared"
)

func TestCardNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.New(context.TODO(), lithic.CardNewParams{
		Type:             lithic.F(lithic.CardNewParamsTypeMerchantLocked),
		AccountToken:     lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CardProgramToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Carrier: lithic.F(shared.CarrierParam{
			QrCodeURL: lithic.F("qr_code_url"),
		}),
		DigitalCardArtToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ExpMonth:            lithic.F("06"),
		ExpYear:             lithic.F("2027"),
		Memo:                lithic.F("New Card"),
		Pin:                 lithic.F("pin"),
		ProductID:           lithic.F("1"),
		ReplacementFor:      lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ShippingAddress: lithic.F(shared.ShippingAddressParam{
			Address1:    lithic.F("5 Broad Street"),
			City:        lithic.F("NEW YORK"),
			Country:     lithic.F("USA"),
			FirstName:   lithic.F("Michael"),
			LastName:    lithic.F("Bluth"),
			PostalCode:  lithic.F("10001-1809"),
			State:       lithic.F("NY"),
			Address2:    lithic.F("Unit 25A"),
			Email:       lithic.F("johnny@appleseed.com"),
			Line2Text:   lithic.F("The Bluth Company"),
			PhoneNumber: lithic.F("+12124007676"),
		}),
		ShippingMethod:     lithic.F(lithic.CardNewParamsShippingMethod2Day),
		SpendLimit:         lithic.F(int64(1000)),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.CardUpdateParams{
			DigitalCardArtToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Memo:                lithic.F("Updated Name"),
			Pin:                 lithic.F("pin"),
			PinStatus:           lithic.F(lithic.CardUpdateParamsPinStatusOk),
			SpendLimit:          lithic.F(int64(100)),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.List(context.TODO(), lithic.CardListParams{
		AccountToken:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Begin:         lithic.F(time.Now()),
		End:           lithic.F(time.Now()),
		EndingBefore:  lithic.F("ending_before"),
		PageSize:      lithic.F(int64(1)),
		StartingAfter: lithic.F("starting_after"),
		State:         lithic.F(lithic.CardListParamsStateClosed),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.Embed(context.TODO(), lithic.CardEmbedParams{
		EmbedRequest: lithic.F("embed_request"),
		Hmac:         lithic.F("hmac"),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.Provision(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.CardProvisionParams{
			Certificate:           lithic.F("U3RhaW5sZXNzIHJvY2tz"),
			ClientDeviceID:        lithic.F("client_device_id"),
			ClientWalletAccountID: lithic.F("client_wallet_account_id"),
			DigitalWallet:         lithic.F(lithic.CardProvisionParamsDigitalWalletApplePay),
			Nonce:                 lithic.F("U3RhaW5sZXNzIHJvY2tz"),
			NonceSignature:        lithic.F("U3RhaW5sZXNzIHJvY2tz"),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.Reissue(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.CardReissueParams{
			Carrier: lithic.F(shared.CarrierParam{
				QrCodeURL: lithic.F("https://lithic.com/activate-card/1"),
			}),
			ProductID: lithic.F("100"),
			ShippingAddress: lithic.F(shared.ShippingAddressParam{
				Address1:    lithic.F("5 Broad Street"),
				City:        lithic.F("NEW YORK"),
				Country:     lithic.F("USA"),
				FirstName:   lithic.F("Janet"),
				LastName:    lithic.F("Yellen"),
				PostalCode:  lithic.F("10001"),
				State:       lithic.F("NY"),
				Address2:    lithic.F("Unit 5A"),
				Email:       lithic.F("johnny@appleseed.com"),
				Line2Text:   lithic.F("The Bluth Company"),
				PhoneNumber: lithic.F("+12124007676"),
			}),
			ShippingMethod: lithic.F(lithic.CardReissueParamsShippingMethod2Day),
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

func TestCardRenewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.Renew(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.CardRenewParams{
			ShippingAddress: lithic.F(shared.ShippingAddressParam{
				Address1:    lithic.F("5 Broad Street"),
				City:        lithic.F("NEW YORK"),
				Country:     lithic.F("USA"),
				FirstName:   lithic.F("Janet"),
				LastName:    lithic.F("Yellen"),
				PostalCode:  lithic.F("10001"),
				State:       lithic.F("NY"),
				Address2:    lithic.F("Unit 5A"),
				Email:       lithic.F("johnny@appleseed.com"),
				Line2Text:   lithic.F("The Bluth Company"),
				PhoneNumber: lithic.F("+12124007676"),
			}),
			Carrier: lithic.F(shared.CarrierParam{
				QrCodeURL: lithic.F("https://lithic.com/activate-card/1"),
			}),
			ExpMonth:       lithic.F("06"),
			ExpYear:        lithic.F("2027"),
			ProductID:      lithic.F("100"),
			ShippingMethod: lithic.F(lithic.CardRenewParamsShippingMethod2Day),
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

func TestCardGetSpendLimits(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.GetSpendLimits(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardSearchByPan(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.Cards.SearchByPan(context.TODO(), lithic.CardSearchByPanParams{
		Pan: lithic.F("4111111289144142"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
