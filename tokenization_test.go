// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestTokenizationSimulateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Tokenizations.Simulate(context.TODO(), lithic.TokenizationSimulateParams{
		Cvv:                       lithic.F("776"),
		ExpirationDate:            lithic.F("xxxxx"),
		Pan:                       lithic.F("4111111289144142"),
		TokenizationSource:        lithic.F(lithic.TokenizationSimulateParamsTokenizationSourceApplePay),
		AccountScore:              lithic.F(int64(0)),
		DeviceScore:               lithic.F(int64(0)),
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
