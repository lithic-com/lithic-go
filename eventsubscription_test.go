package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestEventSubscriptionNewWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.New(context.TODO(), lithic.EventSubscriptionNewParams{Description: lithic.F("string"), Disabled: lithic.F(true), EventTypes: lithic.F([]lithic.EventSubscriptionNewParamsEventTypes{lithic.EventSubscriptionNewParamsEventTypesDisputeUpdated, lithic.EventSubscriptionNewParamsEventTypesDisputeUpdated, lithic.EventSubscriptionNewParamsEventTypesDisputeUpdated}), URL: lithic.F("https://example.com")})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionGet(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.Get(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.Update(
		context.TODO(),
		"string",
		lithic.EventSubscriptionUpdateParams{Description: lithic.F("string"), Disabled: lithic.F(true), EventTypes: lithic.F([]lithic.EventSubscriptionUpdateParamsEventTypes{lithic.EventSubscriptionUpdateParamsEventTypesDisputeUpdated, lithic.EventSubscriptionUpdateParamsEventTypesDisputeUpdated, lithic.EventSubscriptionUpdateParamsEventTypesDisputeUpdated}), URL: lithic.F("https://example.com")},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionListWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.List(context.TODO(), lithic.EventSubscriptionListParams{PageSize: lithic.F(int64(1)), StartingAfter: lithic.F("string"), EndingBefore: lithic.F("string")})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionDelete(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.Delete(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionRecover(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.Recover(
		context.TODO(),
		"string",
		lithic.EventSubscriptionRecoverParams{Begin: lithic.F(time.Now()), End: lithic.F(time.Now())},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionReplayMissing(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.ReplayMissing(
		context.TODO(),
		"string",
		lithic.EventSubscriptionReplayMissingParams{Begin: lithic.F(time.Now()), End: lithic.F(time.Now())},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionGetSecret(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.GetSecret(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionRotateSecret(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.RotateSecret(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}