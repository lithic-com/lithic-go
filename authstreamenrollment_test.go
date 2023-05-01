package lithic_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestAuthStreamEnrollmentGet(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthStreamEnrollment.Get(
		context.TODO(),
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthStreamEnrollmentDisenroll(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.AuthStreamEnrollment.Disenroll(
		context.TODO(),
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthStreamEnrollmentEnrollWithOptionalParams(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.AuthStreamEnrollment.Enroll(context.TODO(), lithic.AuthStreamEnrollmentEnrollParams{WebhookURL: lithic.F("https://example.com")})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthStreamEnrollmentGetSecret(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AuthStreamEnrollment.GetSecret(
		context.TODO(),
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthStreamEnrollmentRotateSecret(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.AuthStreamEnrollment.RotateSecret(
		context.TODO(),
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
