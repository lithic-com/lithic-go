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

func TestDisputesNewWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.New(context.TODO(), &requests.DisputeNewParams{Amount: fields.F(int64(0)), CustomerFiledDate: fields.F(time.Now()), Reason: fields.F(requests.DisputeNewParamsReasonAtmCashMisdispense), TransactionToken: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), CustomerNote: fields.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputesGet(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestDisputesUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.DisputeUpdateParams{Amount: fields.F(int64(0)), CustomerFiledDate: fields.F(time.Now()), CustomerNote: fields.F("string"), Reason: fields.F(requests.DisputeUpdateParamsReasonAtmCashMisdispense)},
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

func TestDisputesListWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.List(context.TODO(), &requests.DisputeListParams{TransactionToken: fields.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Status: fields.F(requests.DisputeListParamsStatusNew), PageSize: fields.F(int64(1)), Begin: fields.F(time.Now()), End: fields.F(time.Now()), StartingAfter: fields.F("string"), EndingBefore: fields.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputesDelete(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestDisputesDeleteEvidence(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.DeleteEvidence(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.DisputesDeleteEvidenceParams{DisputeToken: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
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

func TestDisputesInitiateEvidenceUpload(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.InitiateEvidenceUpload(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestDisputesListEvidencesWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.ListEvidences(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.DisputeListEvidencesParams{PageSize: fields.F(int64(1)), Begin: fields.F(time.Now()), End: fields.F(time.Now()), StartingAfter: fields.F("string"), EndingBefore: fields.F("string")},
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

func TestDisputesGetEvidence(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.GetEvidence(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.DisputesGetEvidenceParams{DisputeToken: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
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

func TestDisputesUploadEvidence(t *testing.T) {
	t.Skip("TODO")
}
