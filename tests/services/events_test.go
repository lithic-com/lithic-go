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

func TestEventsGet(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Get(
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

func TestEventsListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.List(context.TODO(), &requests.EventListParams{Begin: fields.F(time.Now()), End: fields.F(time.Now()), PageSize: fields.F(int64(1)), StartingAfter: fields.F("string"), EndingBefore: fields.F("string"), EventTypes: fields.F([]requests.EventListParamsEventTypes{requests.EventListParamsEventTypesDisputeUpdated, requests.EventListParamsEventTypesDisputeUpdated, requests.EventListParamsEventTypesDisputeUpdated})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
