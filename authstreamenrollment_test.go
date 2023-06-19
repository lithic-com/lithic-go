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

func TestAuthStreamEnrollmentGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.AuthStreamEnrollment.Get(context.TODO())
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthStreamEnrollmentDisenroll(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := client.AuthStreamEnrollment.Disenroll(context.TODO())
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthStreamEnrollmentEnrollWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := client.AuthStreamEnrollment.Enroll(context.TODO(), lithic.AuthStreamEnrollmentEnrollParams{
		WebhookURL: lithic.F("https://example.com"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthStreamEnrollmentGetSecret(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
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
