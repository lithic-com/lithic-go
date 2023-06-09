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

func TestFinancialTransactionGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.FinancialAccounts.FinancialTransactions.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFinancialTransactionListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.FinancialAccounts.FinancialTransactions.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.FinancialTransactionListParams{
			Begin:         lithic.F(time.Now()),
			Category:      lithic.F(lithic.FinancialTransactionListParamsCategoryACH),
			End:           lithic.F(time.Now()),
			EndingBefore:  lithic.F("string"),
			Result:        lithic.F(lithic.FinancialTransactionListParamsResultApproved),
			StartingAfter: lithic.F("string"),
			Status:        lithic.F(lithic.FinancialTransactionListParamsStatusDeclined),
		},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
