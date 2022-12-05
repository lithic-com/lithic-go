package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/fields"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
)

func TestFundingSourcesUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.FundingSources.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.FundingSourceUpdateParams{AccountToken: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), State: fields.F(requests.FundingSourceUpdateParamsStateDeleted)},
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

func TestFundingSourcesListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.FundingSources.List(context.TODO(), &requests.FundingSourceListParams{AccountToken: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Page: fields.F(int64(0)), PageSize: fields.F(int64(1))})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFundingSourcesVerifyWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.FundingSources.Verify(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.FundingSourceVerifyParams{AccountToken: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), MicroDeposits: fields.F([]int64{int64(0), int64(0), int64(0)})},
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
