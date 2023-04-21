package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func TestTransferNewWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transfers.New(context.TODO(), &requests.TransferNewParams{From: lithic.F(requests.FinancialAccount{AccountNumber: lithic.F("string"), Created: lithic.F(time.Now()), RoutingNumber: lithic.F("string"), Token: lithic.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"), Type: lithic.F(requests.FinancialAccountTypeIssuing), Updated: lithic.F(time.Now())}), To: lithic.F(requests.FinancialAccount{AccountNumber: lithic.F("string"), Created: lithic.F(time.Now()), RoutingNumber: lithic.F("string"), Token: lithic.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"), Type: lithic.F(requests.FinancialAccountTypeIssuing), Updated: lithic.F(time.Now())}), Amount: lithic.F(int64(0)), Memo: lithic.F("string"), TransactionToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
