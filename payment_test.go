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

func TestPaymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.New(context.TODO(), lithic.PaymentNewParams{
		Amount:                   lithic.F(int64(1)),
		ExternalBankAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FinancialAccountToken:    lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Method:                   lithic.F(lithic.PaymentNewParamsMethodACHNextDay),
		MethodAttributes: lithic.F(lithic.PaymentNewParamsMethodAttributes{
			SecCode: lithic.F(lithic.PaymentNewParamsMethodAttributesSecCodeCcd),
			Addenda: lithic.F("addenda"),
		}),
		Type:          lithic.F(lithic.PaymentNewParamsTypeCollection),
		Token:         lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Memo:          lithic.F("memo"),
		UserDefinedID: lithic.F("user_defined_id"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentGet(t *testing.T) {
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
	_, err := client.Payments.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.List(context.TODO(), lithic.PaymentListParams{
		AccountToken:          lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Begin:                 lithic.F(time.Now()),
		BusinessAccountToken:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Category:              lithic.F(lithic.PaymentListParamsCategoryACH),
		End:                   lithic.F(time.Now()),
		EndingBefore:          lithic.F("ending_before"),
		FinancialAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:              lithic.F(int64(1)),
		Result:                lithic.F(lithic.PaymentListParamsResultApproved),
		StartingAfter:         lithic.F("starting_after"),
		Status:                lithic.F(lithic.PaymentListParamsStatusDeclined),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentRetry(t *testing.T) {
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
	_, err := client.Payments.Retry(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentSimulateActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.SimulateAction(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.PaymentSimulateActionParams{
			EventType:        lithic.F(lithic.PaymentSimulateActionParamsEventTypeACHOriginationReviewed),
			DateOfDeath:      lithic.F(time.Now()),
			DeclineReason:    lithic.F(lithic.PaymentSimulateActionParamsDeclineReasonProgramTransactionLimitExceeded),
			ReturnAddenda:    lithic.F("return_addenda"),
			ReturnReasonCode: lithic.F("return_reason_code"),
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

func TestPaymentSimulateReceiptWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.SimulateReceipt(context.TODO(), lithic.PaymentSimulateReceiptParams{
		Token:                 lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Amount:                lithic.F(int64(0)),
		FinancialAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ReceiptType:           lithic.F(lithic.PaymentSimulateReceiptParamsReceiptTypeReceiptCredit),
		Memo:                  lithic.F("memo"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentSimulateRelease(t *testing.T) {
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
	_, err := client.Payments.SimulateRelease(context.TODO(), lithic.PaymentSimulateReleaseParams{
		PaymentToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentSimulateReturnWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.SimulateReturn(context.TODO(), lithic.PaymentSimulateReturnParams{
		PaymentToken:     lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ReturnReasonCode: lithic.F("R12"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
