package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestTransferNewWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transfers.New(context.TODO(), lithic.TransferNewParams{From: lithic.F(lithic.FinancialAccountParam{AccountNumber: lithic.F("string"), Created: lithic.F(time.Now()), RoutingNumber: lithic.F("string"), Token: lithic.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"), Type: lithic.F(lithic.FinancialAccountTypeIssuing), Updated: lithic.F(time.Now())}), To: lithic.F(lithic.FinancialAccountParam{AccountNumber: lithic.F("string"), Created: lithic.F(time.Now()), RoutingNumber: lithic.F("string"), Token: lithic.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"), Type: lithic.F(lithic.FinancialAccountTypeIssuing), Updated: lithic.F(time.Now())}), Amount: lithic.F(int64(0)), Memo: lithic.F("string"), TransactionToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}