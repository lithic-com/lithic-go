// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestAutoPagination(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	iter := client.Cards.ListAutoPaging(context.TODO(), lithic.CardListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		card := iter.Current()
		t.Logf("%+v\n", card)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
