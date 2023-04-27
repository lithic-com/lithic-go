package services

import (
	"context"
	"errors"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func TestResponderEndpointNewWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ResponderEndpoints.New(context.TODO(), requests.ResponderEndpointNewParams{URL: lithic.F("https://example.com"), Type: lithic.F(requests.ResponderEndpointNewParamsTypeTokenizationDecisioning)})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponderEndpointDelete(t *testing.T) {
	t.Skip("Prism errors when accept header set but no request body is defined")
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.ResponderEndpoints.Delete(context.TODO(), requests.ResponderEndpointDeleteParams{Type: lithic.F(requests.ResponderEndpointDeleteParamsTypeTokenizationDecisioning)})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponderEndpointCheckStatus(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ResponderEndpoints.CheckStatus(context.TODO(), requests.ResponderEndpointCheckStatusParams{Type: lithic.F(requests.ResponderEndpointCheckStatusParamsTypeTokenizationDecisioning)})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
