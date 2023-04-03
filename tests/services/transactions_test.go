package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
)

func TestTransactionsGet(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.List(context.TODO(), &requests.TransactionListParams{AccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), CardToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Result: lithic.F(requests.TransactionListParamsResultApproved), Begin: lithic.F(time.Now()), End: lithic.F(time.Now()), Page: lithic.F(int64(0)), PageSize: lithic.F(int64(1))})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsSimulateAuthorizationWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateAuthorization(context.TODO(), &requests.TransactionSimulateAuthorizationParams{Amount: lithic.F(int64(0)), Descriptor: lithic.F("COFFEE SHOP"), Pan: lithic.F("4111111289144142"), Status: lithic.F(requests.TransactionSimulateAuthorizationParamsStatusAuthorization), MerchantAcceptorID: lithic.F("OODKZAPJVN4YS7O"), MerchantCurrency: lithic.F("GBP"), MerchantAmount: lithic.F(int64(0)), Mcc: lithic.F("5812"), PartialApprovalCapable: lithic.F(true)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsSimulateAuthorizationAdvice(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateAuthorizationAdvice(context.TODO(), &requests.TransactionSimulateAuthorizationAdviceParams{Amount: lithic.F(int64(0)), Token: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsSimulateClearingWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateClearing(context.TODO(), &requests.TransactionSimulateClearingParams{Amount: lithic.F(int64(0)), Token: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsSimulateCreditAuthorizationWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateCreditAuthorization(context.TODO(), &requests.TransactionSimulateCreditAuthorizationParams{Amount: lithic.F(int64(0)), Descriptor: lithic.F("COFFEE SHOP"), Pan: lithic.F("4111111289144142"), MerchantAcceptorID: lithic.F("XRKGDPOWEWQRRWU"), Mcc: lithic.F("5812")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsSimulateReturn(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateReturn(context.TODO(), &requests.TransactionSimulateReturnParams{Amount: lithic.F(int64(0)), Descriptor: lithic.F("COFFEE SHOP"), Pan: lithic.F("4111111289144142")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsSimulateReturnReversal(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateReturnReversal(context.TODO(), &requests.TransactionSimulateReturnReversalParams{Token: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsSimulateVoidWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.SimulateVoid(context.TODO(), &requests.TransactionSimulateVoidParams{Amount: lithic.F(int64(0)), Token: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Type: lithic.F(requests.TransactionSimulateVoidParamsTypeAuthorizationExpiry)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
