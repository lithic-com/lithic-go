// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestPaymentNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Payments.New(context.TODO(), lithic.PaymentNewParams{
		Amount:                   lithic.F(int64(1)),
		ExternalBankAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FinancialAccountToken:    lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Method:                   lithic.F(lithic.PaymentNewParamsMethodACHNextDay),
		MethodAttributes: lithic.F(lithic.PaymentNewParamsMethodAttributes{
			SecCode: lithic.F(lithic.PaymentNewParamsMethodAttributesSecCodePpd),
		}),
		Type:          lithic.F(lithic.PaymentNewParamsTypePayment),
		Token:         lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Memo:          lithic.F("string"),
		UserDefinedID: lithic.F("string"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Payments.List(context.TODO(), lithic.PaymentListParams{
		EndingBefore:          lithic.F("string"),
		FinancialAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:              lithic.F(int64(1)),
		Result:                lithic.F(lithic.PaymentListParamsResultApproved),
		StartingAfter:         lithic.F("string"),
		Status:                lithic.F(lithic.PaymentListParamsStatusPending),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
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