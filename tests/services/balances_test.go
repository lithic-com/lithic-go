package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func TestBalanceListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Balances.List(context.TODO(), &requests.BalanceListParams{AccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), BalanceDate: lithic.F(time.Now()), FinancialAccountType: lithic.F(requests.BalanceListParamsFinancialAccountTypeIssuing)})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
