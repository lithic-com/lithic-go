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

func TestAccountHoldersGet(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.Get(
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

func TestAccountHoldersUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AccountHolderUpdateParams{Email: fields.F("string"), PhoneNumber: fields.F("string"), BusinessAccountToken: fields.F("string")},
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

func TestAccountHoldersNewWebhook(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.NewWebhook(context.TODO(), &requests.AccountHolderNewWebhookParams{URL: fields.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHoldersListDocuments(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.ListDocuments(
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

func TestAccountHoldersResubmit(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.Resubmit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AccountHolderResubmitParams{Workflow: fields.F(requests.AccountHolderResubmitParamsWorkflowKYCAdvanced), TosTimestamp: fields.F("2018-05-29T21:16:05Z"), Individual: fields.F(requests.Individual{Address: fields.F(requests.Address{Address1: fields.F("123 Old Forest Way"), Address2: fields.F("string"), City: fields.F("Omaha"), Country: fields.F("USA"), PostalCode: fields.F("68022"), State: fields.F("NE")}), Dob: fields.F("1991-03-08 08:00:00"), Email: fields.F("tom@middle-earth.com"), FirstName: fields.F("Tom"), GovernmentID: fields.F("111-23-1412"), LastName: fields.F("Bombadil"), PhoneNumber: fields.F("+12124007676")})},
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

func TestAccountHoldersGetDocument(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.GetDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AccountHoldersGetDocumentParams{AccountHolderToken: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
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

func TestAccountHoldersUploadDocument(t *testing.T) {
	c := lithic.NewLithic(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.AccountHolderUploadDocumentParams{DocumentType: fields.F(requests.AccountHolderUploadDocumentParamsDocumentTypeCommercialLicense)},
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
