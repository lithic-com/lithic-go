// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/shared"
)

func TestAuthRuleV2NewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AuthRules.V2.New(context.TODO(), lithic.AuthRuleV2NewParams{
		Body: lithic.AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokens{
			AccountTokens: lithic.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Parameters: lithic.F[lithic.AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersUnion](lithic.AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParameters{
				Conditions: lithic.F([]lithic.AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersCondition{{
					Attribute: lithic.F(lithic.AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsAttributeMcc),
					Operation: lithic.F(lithic.AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsOperationIsOneOf),
					Value:     lithic.F[lithic.AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensParametersConditionalBlockParametersConditionsValueUnion](shared.UnionString("string")),
				}}),
			}),
			Type: lithic.F(lithic.AuthRuleV2NewParamsBodyCreateAuthRuleRequestAccountTokensTypeConditionalBlock),
		},
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleV2Get(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AuthRules.V2.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleV2UpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AuthRules.V2.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleV2UpdateParams{
			State: lithic.F(lithic.AuthRuleV2UpdateParamsStateInactive),
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

func TestAuthRuleV2ListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AuthRules.V2.List(context.TODO(), lithic.AuthRuleV2ListParams{
		AccountToken:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CardToken:     lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		EndingBefore:  lithic.F("ending_before"),
		PageSize:      lithic.F(int64(1)),
		StartingAfter: lithic.F("starting_after"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleV2Apply(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AuthRules.V2.Apply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleV2ApplyParams{
			Body: lithic.AuthRuleV2ApplyParamsBodyApplyAuthRuleRequestAccountTokens{
				AccountTokens: lithic.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			},
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

func TestAuthRuleV2DraftWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AuthRules.V2.Draft(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleV2DraftParams{
			Parameters: lithic.F[lithic.AuthRuleV2DraftParamsParametersUnion](lithic.AuthRuleV2DraftParamsParametersConditionalBlockParameters{
				Conditions: lithic.F([]lithic.AuthRuleV2DraftParamsParametersConditionalBlockParametersCondition{{
					Attribute: lithic.F(lithic.AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsAttributeMcc),
					Operation: lithic.F(lithic.AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsOperationIsOneOf),
					Value:     lithic.F[lithic.AuthRuleV2DraftParamsParametersConditionalBlockParametersConditionsValueUnion](shared.UnionString("string")),
				}}),
			}),
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

func TestAuthRuleV2Promote(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AuthRules.V2.Promote(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleV2Report(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AuthRules.V2.Report(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
