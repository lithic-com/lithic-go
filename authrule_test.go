package lithic_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestAuthRuleNewWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.New(context.TODO(), lithic.AuthRuleNewParams{AllowedMcc: lithic.F([]string{"string", "string", "string"}), BlockedMcc: lithic.F([]string{"string", "string", "string"}), AllowedCountries: lithic.F([]string{"string", "string", "string"}), BlockedCountries: lithic.F([]string{"string", "string", "string"}), AvsType: lithic.F(lithic.AuthRuleNewParamsAvsTypeZipOnly), AccountTokens: lithic.F([]string{"string", "string", "string"}), CardTokens: lithic.F([]string{"string", "string", "string"}), ProgramLevel: lithic.F(false)})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleGet(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Get(
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

func TestAuthRuleUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleUpdateParams{AllowedMcc: lithic.F([]string{"string", "string", "string"}), BlockedMcc: lithic.F([]string{"string", "string", "string"}), AllowedCountries: lithic.F([]string{"string", "string", "string"}), BlockedCountries: lithic.F([]string{"string", "string", "string"}), AvsType: lithic.F(lithic.AuthRuleUpdateParamsAvsTypeZipOnly)},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleListWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.List(context.TODO(), lithic.AuthRuleListParams{Page: lithic.F(int64(0)), PageSize: lithic.F(int64(1))})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleApplyWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Apply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleApplyParams{CardTokens: lithic.F([]string{"string", "string", "string"}), AccountTokens: lithic.F([]string{"string", "string", "string"}), ProgramLevel: lithic.F(true)},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleRemoveWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Remove(context.TODO(), lithic.AuthRuleRemoveParams{CardTokens: lithic.F([]string{"string", "string", "string"}), AccountTokens: lithic.F([]string{"string", "string", "string"}), ProgramLevel: lithic.F(true)})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}