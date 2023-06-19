// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestUsage(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	card, err := client.Cards.New(context.TODO(), lithic.CardNewParams{
		Type: lithic.F(lithic.CardNewParamsTypeVirtual),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", card.Token)
}
