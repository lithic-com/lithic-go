// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestThreeDSAuthenticationGet(t *testing.T) {
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
	_, err := client.ThreeDS.Authentication.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreeDSAuthenticationSimulateWithOptionalParams(t *testing.T) {
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
	_, err := client.ThreeDS.Authentication.Simulate(context.TODO(), lithic.ThreeDSAuthenticationSimulateParams{
		Merchant: lithic.F(lithic.ThreeDSAuthenticationSimulateParamsMerchant{
			ID:      lithic.F("OODKZAPJVN4YS7O"),
			Country: lithic.F("USA"),
			Mcc:     lithic.F("5812"),
			Name:    lithic.F("COFFEE SHOP"),
		}),
		Pan: lithic.F("4111111289144142"),
		Transaction: lithic.F(lithic.ThreeDSAuthenticationSimulateParamsTransaction{
			Amount:   lithic.F(int64(100)),
			Currency: lithic.F("USD"),
		}),
		CardExpiryCheck: lithic.F(lithic.ThreeDSAuthenticationSimulateParamsCardExpiryCheckMatch),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreeDSAuthenticationSimulateOtpEntry(t *testing.T) {
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
	err := client.ThreeDS.Authentication.SimulateOtpEntry(context.TODO(), lithic.ThreeDSAuthenticationSimulateOtpEntryParams{
		Token: lithic.F("fabd829d-7f7b-4432-a8f2-07ea4889aaac"),
		Otp:   lithic.F("123456"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
