// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestAccountGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism returns invalid data")
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AccountUpdateParams{
			DailySpendLimit:     lithic.F(int64(0)),
			LifetimeSpendLimit:  lithic.F(int64(0)),
			MonthlySpendLimit:   lithic.F(int64(0)),
			State:               lithic.F(lithic.AccountUpdateParamsStateActive),
			VerificationAddress: lithic.F(lithic.AccountUpdateParamsVerificationAddress{Address1: lithic.F("string"), Address2: lithic.F("string"), City: lithic.F("string"), State: lithic.F("string"), PostalCode: lithic.F("string"), Country: lithic.F("string")}),
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

func TestAccountListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.List(context.TODO(), lithic.AccountListParams{
		Begin:    lithic.F(time.Now()),
		End:      lithic.F(time.Now()),
		Page:     lithic.F(int64(0)),
		PageSize: lithic.F(int64(1)),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
