// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"os"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestUsage(t *testing.T) {
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
	card, err := client.Cards.New(context.TODO(), lithic.CardNewParams{
		Type: lithic.F(lithic.CardNewParamsTypeSingleUse),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", card.Token)
}
