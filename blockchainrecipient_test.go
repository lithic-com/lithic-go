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

func TestBlockchainRecipientNewWithOptionalParams(t *testing.T) {
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
	_, err := client.BlockchainRecipients.New(context.TODO(), lithic.BlockchainRecipientNewParams{
		AccountToken: lithic.F("dabadb3b-700c-41e3-8801-d5dfc84ebea0"),
		Address:      lithic.F("0x45bfcf1a6289a0b77b4d3f7d12005a05949fd8c3"),
		Chain:        lithic.F("ETHEREUM"),
		Owner:        lithic.F("John Doe"),
		OwnerType:    lithic.F(lithic.OwnerTypeIndividual),
		AddressTag:   lithic.F("address_tag"),
		Name:         lithic.F("Cold wallet"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
