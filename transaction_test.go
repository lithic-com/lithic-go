package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestTransactionGet(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionListWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.List(context.TODO(), lithic.TransactionListParams{
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
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateAuthorizationWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateAuthorization(context.TODO(), lithic.TransactionSimulateAuthorizationParams{
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
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateAuthorizationAdvice(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateAuthorizationAdvice(context.TODO(), lithic.TransactionSimulateAuthorizationAdviceParams{
		Amount: lithic.F(int64(0)),
		Token:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateClearingWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateClearing(context.TODO(), lithic.TransactionSimulateClearingParams{
		Token:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Amount: lithic.F(int64(0)),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateCreditAuthorizationWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateCreditAuthorization(context.TODO(), lithic.TransactionSimulateCreditAuthorizationParams{
		Amount:             lithic.F(int64(0)),
		Descriptor:         lithic.F("COFFEE SHOP"),
		Pan:                lithic.F("4111111289144142"),
		Mcc:                lithic.F("5812"),
		MerchantAcceptorID: lithic.F("XRKGDPOWEWQRRWU"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateReturn(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateReturn(context.TODO(), lithic.TransactionSimulateReturnParams{
		Amount:     lithic.F(int64(0)),
		Descriptor: lithic.F("COFFEE SHOP"),
		Pan:        lithic.F("4111111289144142"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateReturnReversal(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateReturnReversal(context.TODO(), lithic.TransactionSimulateReturnReversalParams{
		Token: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSimulateVoidWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateVoid(context.TODO(), lithic.TransactionSimulateVoidParams{
		Token:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Amount: lithic.F(int64(0)),
		Type:   lithic.F(lithic.TransactionSimulateVoidParamsTypeAuthorizationExpiry),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
