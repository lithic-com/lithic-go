// File generated from our OpenAPI spec by Stainless.

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

func TestAccountGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Accounts.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	t.Skip("Prism returns invalid data")
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Accounts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AccountUpdateParams{
			DailySpendLimit:    lithic.F(int64(1000)),
			LifetimeSpendLimit: lithic.F(int64(0)),
			MonthlySpendLimit:  lithic.F(int64(0)),
			State:              lithic.F(lithic.AccountUpdateParamsStateActive),
			VerificationAddress: lithic.F(lithic.AccountUpdateParamsVerificationAddress{
				Address1:   lithic.F("string"),
				Address2:   lithic.F("string"),
				City:       lithic.F("string"),
				State:      lithic.F("string"),
				PostalCode: lithic.F("string"),
				Country:    lithic.F("string"),
			}),
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

func TestAccountListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Accounts.List(context.TODO(), lithic.AccountListParams{
		Begin:    lithic.F(time.Now()),
		End:      lithic.F(time.Now()),
		Page:     lithic.F(int64(0)),
		PageSize: lithic.F(int64(1)),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
