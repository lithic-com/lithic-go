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

func TestAuthRuleNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.AuthRules.New(context.TODO(), lithic.AuthRuleNewParams{
		AccountTokens:    lithic.F([]string{"3fa85f64-5717-4562-b3fc-2c963f66afa6"}),
		AllowedCountries: lithic.F([]string{"MEX"}),
		AllowedMcc:       lithic.F([]string{"3000"}),
		BlockedCountries: lithic.F([]string{"USA", "CAN"}),
		BlockedMcc:       lithic.F([]string{"5811", "5812"}),
		CardTokens:       lithic.F([]string{"3fa85f64-5717-4562-b3fc-2c963f66afa6"}),
		ProgramLevel:     lithic.F(false),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.AuthRules.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.AuthRules.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleUpdateParams{
			AllowedCountries: lithic.F([]string{"string", "string", "string"}),
			AllowedMcc:       lithic.F([]string{"string", "string", "string"}),
			BlockedCountries: lithic.F([]string{"string", "string", "string"}),
			BlockedMcc:       lithic.F([]string{"string", "string", "string"}),
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

func TestAuthRuleListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.AuthRules.List(context.TODO(), lithic.AuthRuleListParams{
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

func TestAuthRuleApplyWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.AuthRules.Apply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleApplyParams{
			AccountTokens: lithic.F([]string{"string", "string", "string"}),
			CardTokens:    lithic.F([]string{"string", "string", "string"}),
			ProgramLevel:  lithic.F(true),
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

func TestAuthRuleRemoveWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.AuthRules.Remove(context.TODO(), lithic.AuthRuleRemoveParams{
		AccountTokens: lithic.F([]string{"string", "string", "string"}),
		CardTokens:    lithic.F([]string{"string", "string", "string"}),
		ProgramLevel:  lithic.F(true),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
