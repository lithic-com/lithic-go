// File generated from our OpenAPI spec by Stainless.

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

func TestTokenizationSimulateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Tokenizations.Simulate(context.TODO(), lithic.TokenizationSimulateParams{
		Cvv:                       lithic.F("776"),
		ExpirationDate:            lithic.F("08/29"),
		Pan:                       lithic.F("4111111289144142"),
		TokenizationSource:        lithic.F(lithic.TokenizationSimulateParamsTokenizationSourceApplePay),
		AccountScore:              lithic.F(int64(5)),
		DeviceScore:               lithic.F(int64(5)),
		WalletRecommendedDecision: lithic.F(lithic.TokenizationSimulateParamsWalletRecommendedDecisionApproved),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
