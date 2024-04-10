// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestAccountHolderNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AccountHolders.New(context.TODO(), lithic.AccountHolderNewParams{
		Body: lithic.KYBParam{
			BeneficialOwnerEntities: lithic.F([]lithic.KYBBeneficialOwnerEntityParam{{
				Address: lithic.F(shared.AddressParam{
					Address1:   lithic.F("300 Normal Forest Way"),
					Address2:   lithic.F("string"),
					City:       lithic.F("Portland"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("90210"),
					State:      lithic.F("OR"),
				}),
				DbaBusinessName:   lithic.F("string"),
				GovernmentID:      lithic.F("98-7654321"),
				LegalBusinessName: lithic.F("Majority Holdings LLC"),
				ParentCompany:     lithic.F("string"),
				PhoneNumbers:      lithic.F([]string{"+12124007676"}),
			}}),
			BeneficialOwnerIndividuals: lithic.F([]lithic.KYBBeneficialOwnerIndividualParam{{
				Address: lithic.F(shared.AddressParam{
					Address1:   lithic.F("300 Normal Forest Way"),
					Address2:   lithic.F("string"),
					City:       lithic.F("Portland"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("90210"),
					State:      lithic.F("OR"),
				}),
				Dob:          lithic.F("1991-03-08T08:00:00Z"),
				Email:        lithic.F("tim@left-earth.com"),
				FirstName:    lithic.F("Timmy"),
				GovernmentID: lithic.F("211-23-1412"),
				LastName:     lithic.F("Turner"),
				PhoneNumber:  lithic.F("+12024007611"),
			}}),
			BusinessEntity: lithic.F(lithic.KYBBusinessEntityParam{
				Address: lithic.F(shared.AddressParam{
					Address1:   lithic.F("123 Old Forest Way"),
					Address2:   lithic.F("string"),
					City:       lithic.F("Omaha"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("61022"),
					State:      lithic.F("NE"),
				}),
				DbaBusinessName:   lithic.F("string"),
				GovernmentID:      lithic.F("12-3456789"),
				LegalBusinessName: lithic.F("Busy Business, Inc."),
				ParentCompany:     lithic.F("string"),
				PhoneNumbers:      lithic.F([]string{"+12124007676"}),
			}),
			ControlPerson: lithic.F(lithic.KYBControlPersonParam{
				Address: lithic.F(shared.AddressParam{
					Address1:   lithic.F("451 New Forest Way"),
					Address2:   lithic.F("string"),
					City:       lithic.F("Springfield"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("68022"),
					State:      lithic.F("IL"),
				}),
				Dob:          lithic.F("1991-03-08T08:00:00Z"),
				Email:        lithic.F("tom@middle-pluto.com"),
				FirstName:    lithic.F("Tom"),
				GovernmentID: lithic.F("111-23-1412"),
				LastName:     lithic.F("Timothy"),
				PhoneNumber:  lithic.F("+12024008176"),
			}),
			ExternalID:         lithic.F("string"),
			KYBPassedTimestamp: lithic.F("2022-03-08T08:00:00Z"),
			NatureOfBusiness:   lithic.F("Software company selling solutions to the restaurant industry"),
			TosTimestamp:       lithic.F("2022-03-08T08:00:00Z"),
			WebsiteURL:         lithic.F("https://www.mybusiness.com"),
			Workflow:           lithic.F(lithic.KYBWorkflowKYBBasic),
		},
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AccountHolders.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHolderUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AccountHolders.Update(
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

func TestAccountHolderListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AccountHolders.List(context.TODO(), lithic.AccountHolderListParams{
		EndingBefore:  lithic.F("string"),
		ExternalID:    lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:         lithic.F(int64(0)),
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

func TestAccountHolderListDocuments(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AccountHolders.ListDocuments(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHolderResubmit(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AccountHolders.Resubmit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AccountHolderResubmitParams{
			Individual: lithic.F(lithic.AccountHolderResubmitParamsIndividual{
				Address: lithic.F(shared.AddressParam{
					Address1:   lithic.F("123 Old Forest Way"),
					Address2:   lithic.F("string"),
					City:       lithic.F("Omaha"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("68022"),
					State:      lithic.F("NE"),
				}),
				Dob:          lithic.F("1991-03-08 08:00:00"),
				Email:        lithic.F("tom@middle-earth.com"),
				FirstName:    lithic.F("Tom"),
				GovernmentID: lithic.F("111-23-1412"),
				LastName:     lithic.F("Bombadil"),
				PhoneNumber:  lithic.F("+12124007676"),
			}),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AccountHolders.GetDocument(
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My Lithic API Key"),
	)
	_, err := client.AccountHolders.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AccountHolderUploadDocumentParams{
			DocumentType: lithic.F(lithic.AccountHolderUploadDocumentParamsDocumentTypeDriversLicense),
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
