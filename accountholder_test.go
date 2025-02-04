// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/shared"
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
					City:       lithic.F("Portland"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("90210"),
					State:      lithic.F("OR"),
					Address2:   lithic.F("address2"),
				}),
				GovernmentID:      lithic.F("98-7654321"),
				LegalBusinessName: lithic.F("Majority Holdings LLC"),
				PhoneNumbers:      lithic.F([]string{"+15555555555"}),
				DbaBusinessName:   lithic.F("dba_business_name"),
				ParentCompany:     lithic.F("parent_company"),
			}}),
			BeneficialOwnerIndividuals: lithic.F([]lithic.KYBBeneficialOwnerIndividualParam{{
				Address: lithic.F(shared.AddressParam{
					Address1:   lithic.F("300 Normal Forest Way"),
					City:       lithic.F("Portland"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("90210"),
					State:      lithic.F("OR"),
					Address2:   lithic.F("address2"),
				}),
				Dob:          lithic.F("1991-03-08T08:00:00Z"),
				Email:        lithic.F("tim@left-earth.com"),
				FirstName:    lithic.F("Timmy"),
				GovernmentID: lithic.F("211-23-1412"),
				LastName:     lithic.F("Turner"),
				PhoneNumber:  lithic.F("+15555555555"),
			}}),
			BusinessEntity: lithic.F(lithic.KYBBusinessEntityParam{
				Address: lithic.F(shared.AddressParam{
					Address1:   lithic.F("123 Old Forest Way"),
					City:       lithic.F("Omaha"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("61022"),
					State:      lithic.F("NE"),
					Address2:   lithic.F("address2"),
				}),
				GovernmentID:      lithic.F("12-3456789"),
				LegalBusinessName: lithic.F("Busy Business, Inc."),
				PhoneNumbers:      lithic.F([]string{"+15555555555"}),
				DbaBusinessName:   lithic.F("dba_business_name"),
				ParentCompany:     lithic.F("parent_company"),
			}),
			ControlPerson: lithic.F(lithic.KYBControlPersonParam{
				Address: lithic.F(shared.AddressParam{
					Address1:   lithic.F("451 New Forest Way"),
					City:       lithic.F("Springfield"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("68022"),
					State:      lithic.F("IL"),
					Address2:   lithic.F("address2"),
				}),
				Dob:          lithic.F("1991-03-08T08:00:00Z"),
				Email:        lithic.F("tom@middle-pluto.com"),
				FirstName:    lithic.F("Tom"),
				GovernmentID: lithic.F("111-23-1412"),
				LastName:     lithic.F("Timothy"),
				PhoneNumber:  lithic.F("+15555555555"),
			}),
			NatureOfBusiness:   lithic.F("Software company selling solutions to the restaurant industry"),
			TosTimestamp:       lithic.F("2022-03-08T08:00:00Z"),
			Workflow:           lithic.F(lithic.KYBWorkflowKYBBasic),
			ExternalID:         lithic.F("external_id"),
			KYBPassedTimestamp: lithic.F("2022-03-08T08:00:00Z"),
			WebsiteURL:         lithic.F("https://www.mybusiness.com"),
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
			BusinessAccountToken: lithic.F("business_account_token"),
			Email:                lithic.F("email"),
			PhoneNumber:          lithic.F("phone_number"),
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
		Begin:             lithic.F(time.Now()),
		Email:             lithic.F("email"),
		End:               lithic.F(time.Now()),
		EndingBefore:      lithic.F("ending_before"),
		ExternalID:        lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FirstName:         lithic.F("first_name"),
		LastName:          lithic.F("last_name"),
		LegalBusinessName: lithic.F("legal_business_name"),
		Limit:             lithic.F(int64(0)),
		PhoneNumber:       lithic.F("phone_number"),
		StartingAfter:     lithic.F("starting_after"),
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

func TestAccountHolderSimulateEnrollmentDocumentReviewWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountHolders.SimulateEnrollmentDocumentReview(context.TODO(), lithic.AccountHolderSimulateEnrollmentDocumentReviewParams{
		DocumentUploadToken:         lithic.F("b11cd67b-0a52-4180-8365-314f3def5426"),
		Status:                      lithic.F(lithic.AccountHolderSimulateEnrollmentDocumentReviewParamsStatusUploaded),
		AcceptedEntityStatusReasons: lithic.F([]string{"string"}),
		StatusReason:                lithic.F(lithic.AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentMissingRequiredData),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHolderSimulateEnrollmentReviewWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountHolders.SimulateEnrollmentReview(context.TODO(), lithic.AccountHolderSimulateEnrollmentReviewParams{
		AccountHolderToken: lithic.F("1415964d-4400-4d79-9fb3-eee0faaee4e4"),
		Status:             lithic.F(lithic.AccountHolderSimulateEnrollmentReviewParamsStatusAccepted),
		StatusReasons:      lithic.F([]lithic.AccountHolderSimulateEnrollmentReviewParamsStatusReason{lithic.AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityIDVerificationFailure}),
	})
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
			DocumentType: lithic.F(lithic.AccountHolderUploadDocumentParamsDocumentTypeEinLetter),
			EntityToken:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
