// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestBookTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.BookTransfers.New(context.TODO(), lithic.BookTransferNewParams{
		Amount:                    lithic.F(int64(1)),
		Category:                  lithic.F(lithic.BookTransferNewParamsCategoryAdjustment),
		FromFinancialAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Subtype:                   lithic.F("subtype"),
		ToFinancialAccountToken:   lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Type:                      lithic.F(lithic.BookTransferNewParamsTypeAtmWithdrawal),
		Token:                     lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ExternalID:                lithic.F("external_id"),
		Memo:                      lithic.F("memo"),
		OnClosedAccount:           lithic.F(lithic.BookTransferNewParamsOnClosedAccountFail),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookTransferGet(t *testing.T) {
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
	_, err := client.BookTransfers.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.BookTransfers.List(context.TODO(), lithic.BookTransferListParams{
		AccountToken:          lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Begin:                 lithic.F(time.Now()),
		BusinessAccountToken:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Category:              lithic.F(lithic.BookTransferListParamsCategoryBalanceOrFunding),
		End:                   lithic.F(time.Now()),
		EndingBefore:          lithic.F("ending_before"),
		FinancialAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:              lithic.F(int64(1)),
		Result:                lithic.F(lithic.BookTransferListParamsResultApproved),
		StartingAfter:         lithic.F("starting_after"),
		Status:                lithic.F(lithic.BookTransferListParamsStatusDeclined),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookTransferReverseWithOptionalParams(t *testing.T) {
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
	_, err := client.BookTransfers.Reverse(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.BookTransferReverseParams{
			Memo: lithic.F("memo"),
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
