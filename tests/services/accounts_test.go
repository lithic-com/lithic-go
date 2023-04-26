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

func TestAccountGet(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Get(
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

func TestAccountUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism returns invalid data")
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AccountUpdateParams{DailySpendLimit: lithic.F(int64(0)), LifetimeSpendLimit: lithic.F(int64(0)), MonthlySpendLimit: lithic.F(int64(0)), VerificationAddress: lithic.F(requests.AccountUpdateParamsVerificationAddress{Address1: lithic.F("string"), Address2: lithic.F("string"), City: lithic.F("string"), State: lithic.F("string"), PostalCode: lithic.F("string"), Country: lithic.F("string")}), State: lithic.F(requests.AccountUpdateParamsStateActive)},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.List(context.TODO(), &requests.AccountListParams{Begin: lithic.F(time.Now()), End: lithic.F(time.Now()), Page: lithic.F(int64(0)), PageSize: lithic.F(int64(1))})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
