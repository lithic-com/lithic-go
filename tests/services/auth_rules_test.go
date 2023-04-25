package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func TestAuthRuleNewWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.New(context.TODO(), &requests.AuthRuleNewParams{AllowedMcc: lithic.F([]string{"string", "string", "string"}), BlockedMcc: lithic.F([]string{"string", "string", "string"}), AllowedCountries: lithic.F([]string{"string", "string", "string"}), BlockedCountries: lithic.F([]string{"string", "string", "string"}), AvsType: lithic.F(requests.AuthRuleNewParamsAvsTypeZipOnly), AccountTokens: lithic.F([]string{"string", "string", "string"}), CardTokens: lithic.F([]string{"string", "string", "string"}), ProgramLevel: lithic.F(false)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleGet(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Get(
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

func TestAuthRuleUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AuthRuleUpdateParams{AllowedMcc: lithic.F([]string{"string", "string", "string"}), BlockedMcc: lithic.F([]string{"string", "string", "string"}), AllowedCountries: lithic.F([]string{"string", "string", "string"}), BlockedCountries: lithic.F([]string{"string", "string", "string"}), AvsType: lithic.F(requests.AuthRuleUpdateParamsAvsTypeZipOnly)},
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

func TestAuthRuleListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.List(context.TODO(), &requests.AuthRuleListParams{Page: lithic.F(int64(0)), PageSize: lithic.F(int64(1))})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleApplyWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Apply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AuthRuleApplyParams{CardTokens: lithic.F([]string{"string", "string", "string"}), AccountTokens: lithic.F([]string{"string", "string", "string"}), ProgramLevel: lithic.F(true)},
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

func TestAuthRuleRemoveWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Remove(context.TODO(), &requests.AuthRuleRemoveParams{CardTokens: lithic.F([]string{"string", "string", "string"}), AccountTokens: lithic.F([]string{"string", "string", "string"}), ProgramLevel: lithic.F(true)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
