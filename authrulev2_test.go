// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

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
			Body: lithic.AuthRuleV2UpdateParamsBodyAccountLevelRule{
				AccountTokens:         lithic.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				BusinessAccountTokens: lithic.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				Name:                  lithic.F("name"),
				State:                 lithic.F(lithic.AuthRuleV2UpdateParamsBodyAccountLevelRuleStateInactive),
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
		AccountToken:         lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		BusinessAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CardToken:            lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		EndingBefore:         lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		EventStream:          lithic.F(lithic.AuthRuleV2ListParamsEventStreamAuthorization),
		EventStreams:         lithic.F([]lithic.AuthRuleV2ListParamsEventStream{lithic.AuthRuleV2ListParamsEventStreamAuthorization}),
		PageSize:             lithic.F(int64(1)),
		Scope:                lithic.F(lithic.AuthRuleV2ListParamsScopeProgram),
		StartingAfter:        lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRuleV2Delete(t *testing.T) {
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
	err := client.AuthRules.V2.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
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

func TestAuthRuleV2GetFeaturesWithOptionalParams(t *testing.T) {
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
	_, err := client.AuthRules.V2.GetFeatures(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleV2GetFeaturesParams{
			AccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CardToken:    lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestAuthRuleV2GetReport(t *testing.T) {
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
	_, err := client.AuthRules.V2.GetReport(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AuthRuleV2GetReportParams{
			Begin: lithic.F(time.Now()),
			End:   lithic.F(time.Now()),
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
