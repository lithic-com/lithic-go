package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestEventGet(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Get(
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

func TestEventListWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.List(context.TODO(), lithic.EventListParams{Begin: lithic.F(time.Now()), End: lithic.F(time.Now()), PageSize: lithic.F(int64(1)), StartingAfter: lithic.F("string"), EndingBefore: lithic.F("string"), EventTypes: lithic.F([]lithic.EventListParamsEventTypes{lithic.EventListParamsEventTypesDisputeUpdated, lithic.EventListParamsEventTypesDisputeUpdated, lithic.EventListParamsEventTypesDisputeUpdated})})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}