// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestTransferNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Transfers.New(context.TODO(), lithic.TransferNewParams{
		Amount: lithic.F(int64(0)),
		From: lithic.F(lithic.FinancialAccountParam{
			AccountNumber: lithic.F("string"),
			Created:       lithic.F(time.Now()),
			RoutingNumber: lithic.F("string"),
			Token:         lithic.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
			Type:          lithic.F(lithic.FinancialAccountTypeIssuing),
			Updated:       lithic.F(time.Now()),
		}),
		To: lithic.F(lithic.FinancialAccountParam{
			AccountNumber: lithic.F("string"),
			Created:       lithic.F(time.Now()),
			RoutingNumber: lithic.F("string"),
			Token:         lithic.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
			Type:          lithic.F(lithic.FinancialAccountTypeIssuing),
			Updated:       lithic.F(time.Now()),
		}),
		Memo:             lithic.F("string"),
		TransactionToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
