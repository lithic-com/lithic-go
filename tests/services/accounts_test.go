package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/fields"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
)

func TestAccountsGet(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Get(
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

func TestAccountsUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism returns invalid data")
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AccountUpdateParams{DailySpendLimit: fields.F(int64(0)), LifetimeSpendLimit: fields.F(int64(0)), MonthlySpendLimit: fields.F(int64(0)), VerificationAddress: fields.F(requests.AccountUpdateParamsVerificationAddress{Address1: fields.F("string"), Address2: fields.F("string"), City: fields.F("string"), State: fields.F("string"), PostalCode: fields.F("string"), Country: fields.F("string")}), State: fields.F(requests.AccountUpdateParamsStateActive)},
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

func TestAccountsListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.List(context.TODO(), &requests.AccountListParams{Begin: fields.F(time.Now()), End: fields.F(time.Now()), Page: fields.F(int64(0)), PageSize: fields.F(int64(1))})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
