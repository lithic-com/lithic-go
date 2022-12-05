package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/fields"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
)

func TestAuthRulesNewWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.New(context.TODO(), &requests.AuthRuleRequest{AllowedMcc: fields.F([]string{"string", "string", "string"}), BlockedMcc: fields.F([]string{"string", "string", "string"}), AllowedCountries: fields.F([]string{"string", "string", "string"}), BlockedCountries: fields.F([]string{"string", "string", "string"}), AvsType: fields.F(requests.AuthRuleRequestAvsTypeZipOnly), AccountTokens: fields.F([]string{"string", "string", "string"}), CardTokens: fields.F([]string{"string", "string", "string"}), ProgramLevel: fields.F(false)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRulesGet(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
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

func TestAuthRulesUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AuthRuleUpdateParams{AllowedMcc: fields.F([]string{"string", "string", "string"}), BlockedMcc: fields.F([]string{"string", "string", "string"}), AllowedCountries: fields.F([]string{"string", "string", "string"}), BlockedCountries: fields.F([]string{"string", "string", "string"}), AvsType: fields.F(requests.AuthRuleUpdateParamsAvsTypeZipOnly)},
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

func TestAuthRulesListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.List(context.TODO(), &requests.AuthRuleListParams{Page: fields.F(int64(0)), PageSize: fields.F(int64(1))})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRulesApplyWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Apply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AuthRuleApplyParams{CardTokens: fields.F([]string{"string", "string", "string"}), AccountTokens: fields.F([]string{"string", "string", "string"}), ProgramLevel: fields.F(true)},
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

func TestAuthRulesRemoveWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthRules.Remove(context.TODO(), &requests.AuthRuleRemoveParams{CardTokens: fields.F([]string{"string", "string", "string"}), AccountTokens: fields.F([]string{"string", "string", "string"}), ProgramLevel: fields.F(true)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
