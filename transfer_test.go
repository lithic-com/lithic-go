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

func TestTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Transfers.New(context.TODO(), lithic.TransferNewParams{
		Amount: lithic.F(int64(0)),
		From:   lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		To:     lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Token:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Memo:   lithic.F("memo"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
