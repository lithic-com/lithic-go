package lithic_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestResponderEndpointNewWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ResponderEndpoints.New(context.TODO(), lithic.ResponderEndpointNewParams{URL: lithic.F("https://example.com"), Type: lithic.F(lithic.ResponderEndpointNewParamsTypeTokenizationDecisioning)})
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
	err := c.ResponderEndpoints.Delete(context.TODO(), lithic.ResponderEndpointDeleteParams{Type: lithic.F(lithic.ResponderEndpointDeleteParamsTypeTokenizationDecisioning)})
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
	_, err := c.ResponderEndpoints.CheckStatus(context.TODO(), lithic.ResponderEndpointCheckStatusParams{Type: lithic.F(lithic.ResponderEndpointCheckStatusParamsTypeTokenizationDecisioning)})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}