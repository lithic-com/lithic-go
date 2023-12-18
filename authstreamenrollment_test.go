// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestAuthStreamEnrollmentGetSecret(t *testing.T) {
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
	_, err := client.AuthStreamEnrollment.GetSecret(context.TODO())
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthStreamEnrollmentRotateSecret(t *testing.T) {
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
	err := client.AuthStreamEnrollment.RotateSecret(context.TODO())
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
