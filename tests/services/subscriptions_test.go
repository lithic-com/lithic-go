package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/fields"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
)

func TestSubscriptionsNewWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.New(context.TODO(), &requests.SubscriptionNewParams{Description: fields.F("string"), Disabled: fields.F(true), EventTypes: fields.F([]requests.SubscriptionNewParamsEventTypes{requests.SubscriptionNewParamsEventTypesDisputeUpdated, requests.SubscriptionNewParamsEventTypesDisputeUpdated, requests.SubscriptionNewParamsEventTypesDisputeUpdated}), URL: fields.F("https://example.com")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionsGet(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.Get(
		context.TODO(),
		"string",
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

func TestSubscriptionsUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.Update(
		context.TODO(),
		"string",
		&requests.SubscriptionUpdateParams{Description: fields.F("string"), Disabled: fields.F(true), EventTypes: fields.F([]requests.SubscriptionUpdateParamsEventTypes{requests.SubscriptionUpdateParamsEventTypesDisputeUpdated, requests.SubscriptionUpdateParamsEventTypesDisputeUpdated, requests.SubscriptionUpdateParamsEventTypesDisputeUpdated}), URL: fields.F("https://example.com")},
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

func TestSubscriptionsListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.List(context.TODO(), &requests.SubscriptionListParams{PageSize: fields.F(int64(1)), StartingAfter: fields.F("string"), EndingBefore: fields.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionsDelete(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.Delete(
		context.TODO(),
		"string",
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

func TestSubscriptionsRecover(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.Recover(
		context.TODO(),
		"string",
		&requests.SubscriptionRecoverParams{Begin: fields.F(time.Now()), End: fields.F(time.Now())},
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

func TestSubscriptionsReplayMissing(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.ReplayMissing(
		context.TODO(),
		"string",
		&requests.SubscriptionReplayMissingParams{Begin: fields.F(time.Now()), End: fields.F(time.Now())},
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

func TestSubscriptionsGetSecret(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.GetSecret(
		context.TODO(),
		"string",
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

func TestSubscriptionsRotateSecret(t *testing.T) {
	t.Skip("Prism Mock server doesnt want Accept header, but server requires it.")
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.RotateSecret(
		context.TODO(),
		"string",
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
