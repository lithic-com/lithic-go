package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func TestSubscriptionsNewWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.New(context.TODO(), &requests.SubscriptionNewParams{Description: lithic.F("string"), Disabled: lithic.F(true), EventTypes: lithic.F([]requests.SubscriptionNewParamsEventTypes{requests.SubscriptionNewParamsEventTypesDisputeUpdated, requests.SubscriptionNewParamsEventTypesDisputeUpdated, requests.SubscriptionNewParamsEventTypesDisputeUpdated}), URL: lithic.F("https://example.com")})
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
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.Update(
		context.TODO(),
		"string",
		&requests.SubscriptionUpdateParams{Description: lithic.F("string"), Disabled: lithic.F(true), EventTypes: lithic.F([]requests.SubscriptionUpdateParamsEventTypes{requests.SubscriptionUpdateParamsEventTypesDisputeUpdated, requests.SubscriptionUpdateParamsEventTypesDisputeUpdated, requests.SubscriptionUpdateParamsEventTypesDisputeUpdated}), URL: lithic.F("https://example.com")},
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
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Subscriptions.List(context.TODO(), &requests.SubscriptionListParams{PageSize: lithic.F(int64(1)), StartingAfter: lithic.F("string"), EndingBefore: lithic.F("string")})
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
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.Recover(
		context.TODO(),
		"string",
		&requests.SubscriptionRecoverParams{Begin: lithic.F(time.Now()), End: lithic.F(time.Now())},
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
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.Events.Subscriptions.ReplayMissing(
		context.TODO(),
		"string",
		&requests.SubscriptionReplayMissingParams{Begin: lithic.F(time.Now()), End: lithic.F(time.Now())},
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
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
