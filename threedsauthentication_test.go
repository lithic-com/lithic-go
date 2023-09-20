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

func TestThreeDSAuthenticationGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
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

func TestThreeDSAuthenticationSimulate(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.ThreeDS.Authentication.Simulate(context.TODO(), lithic.ThreeDSAuthenticationSimulateParams{
		Merchant: lithic.F(lithic.ThreeDSAuthenticationSimulateParamsMerchant{
			Country: lithic.F("USA"),
			Mcc:     lithic.F("5812"),
			ID:      lithic.F("OODKZAPJVN4YS7O"),
			Name:    lithic.F("COFFEE SHOP"),
		}),
		Pan: lithic.F("4111111289144142"),
		Transaction: lithic.F(lithic.ThreeDSAuthenticationSimulateParamsTransaction{
			Amount:   lithic.F(int64(0)),
			Currency: lithic.F("GBP"),
		}),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
