// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestAccountHolderNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.New(context.TODO(), lithic.AccountHolderNewParamsKYB{
		BeneficialOwnerEntities:    lithic.F([]lithic.AccountHolderNewParamsKYBBeneficialOwnerEntities{{Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), DbaBusinessName: lithic.F("string"), GovernmentID: lithic.F("114-123-1513"), LegalBusinessName: lithic.F("Acme, Inc."), ParentCompany: lithic.F("string"), PhoneNumbers: lithic.F([]string{"+12124007676"})}, {Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), DbaBusinessName: lithic.F("string"), GovernmentID: lithic.F("114-123-1513"), LegalBusinessName: lithic.F("Acme, Inc."), ParentCompany: lithic.F("string"), PhoneNumbers: lithic.F([]string{"+12124007676"})}, {Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), DbaBusinessName: lithic.F("string"), GovernmentID: lithic.F("114-123-1513"), LegalBusinessName: lithic.F("Acme, Inc."), ParentCompany: lithic.F("string"), PhoneNumbers: lithic.F([]string{"+12124007676"})}}),
		BeneficialOwnerIndividuals: lithic.F([]lithic.AccountHolderNewParamsKYBBeneficialOwnerIndividuals{{Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), Dob: lithic.F("1991-03-08 08:00:00"), Email: lithic.F("tom@middle-earth.com"), FirstName: lithic.F("Tom"), GovernmentID: lithic.F("111-23-1412"), LastName: lithic.F("Bombadil"), PhoneNumber: lithic.F("+12124007676")}, {Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), Dob: lithic.F("1991-03-08 08:00:00"), Email: lithic.F("tom@middle-earth.com"), FirstName: lithic.F("Tom"), GovernmentID: lithic.F("111-23-1412"), LastName: lithic.F("Bombadil"), PhoneNumber: lithic.F("+12124007676")}, {Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), Dob: lithic.F("1991-03-08 08:00:00"), Email: lithic.F("tom@middle-earth.com"), FirstName: lithic.F("Tom"), GovernmentID: lithic.F("111-23-1412"), LastName: lithic.F("Bombadil"), PhoneNumber: lithic.F("+12124007676")}}),
		BusinessEntity:             lithic.F(lithic.AccountHolderNewParamsKYBBusinessEntity{Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), DbaBusinessName: lithic.F("string"), GovernmentID: lithic.F("114-123-1513"), LegalBusinessName: lithic.F("Acme, Inc."), ParentCompany: lithic.F("string"), PhoneNumbers: lithic.F([]string{"+12124007676"})}),
		ControlPerson:              lithic.F(lithic.AccountHolderNewParamsKYBControlPerson{Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), Dob: lithic.F("1991-03-08 08:00:00"), Email: lithic.F("tom@middle-earth.com"), FirstName: lithic.F("Tom"), GovernmentID: lithic.F("111-23-1412"), LastName: lithic.F("Bombadil"), PhoneNumber: lithic.F("+12124007676")}),
		NatureOfBusiness:           lithic.F("Software company selling solutions to the restaurant industry"),
		TosTimestamp:               lithic.F("2018-05-29T21:16:05Z"),
		WebsiteURL:                 lithic.F("www.mybusiness.com"),
		Workflow:                   lithic.F(lithic.AccountHolderNewParamsKYBWorkflowKYBBasic),
		KYBPassedTimestamp:         lithic.F("2018-05-29T21:16:05Z"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHolderGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHolderUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AccountHolderUpdateParams{
			BusinessAccountToken: lithic.F("string"),
			Email:                lithic.F("string"),
			PhoneNumber:          lithic.F("string"),
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

func TestAccountHolderNewWebhook(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.NewWebhook(context.TODO(), lithic.AccountHolderNewWebhookParams{
		URL: lithic.F("string"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHolderListDocuments(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.ListDocuments(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHolderResubmit(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.Resubmit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AccountHolderResubmitParams{
			Individual:   lithic.F(lithic.AccountHolderResubmitParamsIndividual{Address: lithic.F(shared.AddressParam{Address1: lithic.F("123 Old Forest Way"), Address2: lithic.F("string"), City: lithic.F("Omaha"), Country: lithic.F("USA"), PostalCode: lithic.F("68022"), State: lithic.F("NE")}), Dob: lithic.F("1991-03-08 08:00:00"), Email: lithic.F("tom@middle-earth.com"), FirstName: lithic.F("Tom"), GovernmentID: lithic.F("111-23-1412"), LastName: lithic.F("Bombadil"), PhoneNumber: lithic.F("+12124007676")}),
			TosTimestamp: lithic.F("2018-05-29T21:16:05Z"),
			Workflow:     lithic.F(lithic.AccountHolderResubmitParamsWorkflowKYCAdvanced),
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

func TestAccountHolderGetDocument(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.GetDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHolderUploadDocument(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountHolders.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AccountHolderUploadDocumentParams{
			DocumentType: lithic.F(lithic.AccountHolderUploadDocumentParamsDocumentTypeCommercialLicense),
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
