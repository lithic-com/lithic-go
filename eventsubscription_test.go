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

func TestEventSubscriptionNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Events.Subscriptions.New(context.TODO(), lithic.EventSubscriptionNewParams{
		URL:         lithic.F("https://example.com"),
		Description: lithic.F("string"),
		Disabled:    lithic.F(true),
		EventTypes:  lithic.F([]lithic.EventSubscriptionNewParamsEventType{lithic.EventSubscriptionNewParamsEventTypeCardCreated, lithic.EventSubscriptionNewParamsEventTypeCardCreated, lithic.EventSubscriptionNewParamsEventTypeCardCreated}),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Events.Subscriptions.Get(context.TODO(), "string")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Events.Subscriptions.Update(
		context.TODO(),
		"string",
		lithic.EventSubscriptionUpdateParams{
			URL:         lithic.F("https://example.com"),
			Description: lithic.F("string"),
			Disabled:    lithic.F(true),
			EventTypes:  lithic.F([]lithic.EventSubscriptionUpdateParamsEventType{lithic.EventSubscriptionUpdateParamsEventTypeCardCreated, lithic.EventSubscriptionUpdateParamsEventTypeCardCreated, lithic.EventSubscriptionUpdateParamsEventTypeCardCreated}),
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

func TestEventSubscriptionListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Events.Subscriptions.List(context.TODO(), lithic.EventSubscriptionListParams{
		EndingBefore:  lithic.F("string"),
		PageSize:      lithic.F(int64(1)),
		StartingAfter: lithic.F("string"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	err := client.Events.Subscriptions.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionListAttemptsWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Events.Subscriptions.ListAttempts(
		context.TODO(),
		"string",
		lithic.EventSubscriptionListAttemptsParams{
			Begin:         lithic.F(time.Now()),
			End:           lithic.F(time.Now()),
			EndingBefore:  lithic.F("string"),
			PageSize:      lithic.F(int64(1)),
			StartingAfter: lithic.F("string"),
			Status:        lithic.F(lithic.EventSubscriptionListAttemptsParamsStatusFailed),
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

func TestEventSubscriptionRecoverWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	err := client.Events.Subscriptions.Recover(
		context.TODO(),
		"string",
		lithic.EventSubscriptionRecoverParams{
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

func TestEventSubscriptionReplayMissingWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	err := client.Events.Subscriptions.ReplayMissing(
		context.TODO(),
		"string",
		lithic.EventSubscriptionReplayMissingParams{
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

func TestEventSubscriptionGetSecret(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Events.Subscriptions.GetSecret(context.TODO(), "string")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionRotateSecret(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	err := client.Events.Subscriptions.RotateSecret(context.TODO(), "string")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
