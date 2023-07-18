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

func TestEventGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Events.Get(context.TODO(), "string")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Events.List(context.TODO(), lithic.EventListParams{
		Begin:         lithic.F(time.Now()),
		End:           lithic.F(time.Now()),
		EndingBefore:  lithic.F("string"),
		EventTypes:    lithic.F([]lithic.EventListParamsEventType{lithic.EventListParamsEventTypeCardCreated, lithic.EventListParamsEventTypeCardCreated, lithic.EventListParamsEventTypeCardCreated}),
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

func TestEventListAttemptsWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Events.ListAttempts(
		context.TODO(),
		"string",
		lithic.EventListAttemptsParams{
			Begin:         lithic.F(time.Now()),
			End:           lithic.F(time.Now()),
			EndingBefore:  lithic.F("string"),
			PageSize:      lithic.F(int64(1)),
			StartingAfter: lithic.F("string"),
			Status:        lithic.F(lithic.EventListAttemptsParamsStatusFailed),
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
