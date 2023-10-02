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

func TestTransactionGet(t *testing.T) {
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
	_, err := client.Transactions.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), lithic.TransactionListParams{
		AccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Begin:        lithic.F(time.Now()),
		CardToken:    lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		End:          lithic.F(time.Now()),
		Page:         lithic.F(int64(0)),
		PageSize:     lithic.F(int64(1)),
		Result:       lithic.F(lithic.TransactionListParamsResultApproved),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateAuthorizationWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.SimulateAuthorization(context.TODO(), lithic.TransactionSimulateAuthorizationParams{
		Amount:                 lithic.F(int64(0)),
		Descriptor:             lithic.F("COFFEE SHOP"),
		Pan:                    lithic.F("4111111289144142"),
		Mcc:                    lithic.F("5812"),
		MerchantAcceptorID:     lithic.F("OODKZAPJVN4YS7O"),
		MerchantAmount:         lithic.F(int64(0)),
		MerchantCurrency:       lithic.F("GBP"),
		PartialApprovalCapable: lithic.F(true),
		Status:                 lithic.F(lithic.TransactionSimulateAuthorizationParamsStatusAuthorization),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateAuthorizationAdvice(t *testing.T) {
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
	_, err := client.Transactions.SimulateAuthorizationAdvice(context.TODO(), lithic.TransactionSimulateAuthorizationAdviceParams{
		Token:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Amount: lithic.F(int64(0)),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateClearingWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.SimulateClearing(context.TODO(), lithic.TransactionSimulateClearingParams{
		Token:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Amount: lithic.F(int64(0)),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateCreditAuthorizationWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.SimulateCreditAuthorization(context.TODO(), lithic.TransactionSimulateCreditAuthorizationParams{
		Amount:             lithic.F(int64(0)),
		Descriptor:         lithic.F("COFFEE SHOP"),
		Pan:                lithic.F("4111111289144142"),
		Mcc:                lithic.F("5812"),
		MerchantAcceptorID: lithic.F("XRKGDPOWEWQRRWU"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateReturn(t *testing.T) {
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
	_, err := client.Transactions.SimulateReturn(context.TODO(), lithic.TransactionSimulateReturnParams{
		Amount:     lithic.F(int64(0)),
		Descriptor: lithic.F("COFFEE SHOP"),
		Pan:        lithic.F("4111111289144142"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateReturnReversal(t *testing.T) {
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
	_, err := client.Transactions.SimulateReturnReversal(context.TODO(), lithic.TransactionSimulateReturnReversalParams{
		Token: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateVoidWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.SimulateVoid(context.TODO(), lithic.TransactionSimulateVoidParams{
		Token:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Amount: lithic.F(int64(0)),
		Type:   lithic.F(lithic.TransactionSimulateVoidParamsTypeAuthorizationExpiry),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
