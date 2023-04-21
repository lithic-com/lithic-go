package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func TestResponderEndpointNewWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ResponderEndpoints.New(context.TODO(), &requests.ResponderEndpointNewParams{URL: lithic.F("https://example.com"), Type: lithic.F(requests.ResponderEndpointNewParamsTypeTokenizationDecisioning)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponderEndpointDelete(t *testing.T) {
	t.Skip("Prism errors when accept header set but no request body is defined")
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	err := c.ResponderEndpoints.Delete(context.TODO(), &requests.ResponderEndpointDeleteParams{Type: lithic.F(requests.ResponderEndpointDeleteParamsTypeTokenizationDecisioning)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponderEndpointCheckStatus(t *testing.T) {
	c := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ResponderEndpoints.CheckStatus(context.TODO(), &requests.ResponderEndpointCheckStatusParams{Type: lithic.F(requests.ResponderEndpointCheckStatusParamsTypeTokenizationDecisioning)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
