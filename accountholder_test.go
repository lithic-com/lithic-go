// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic_test

import (
	"context"
	"errors"
	"os"
	"testing"

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
		Body: lithic.KYCExemptParam{
			Address: lithic.F(shared.AddressParam{
				Address1:   lithic.F("123 Old Forest Way"),
				City:       lithic.F("Omaha"),
				Country:    lithic.F("USA"),
				PostalCode: lithic.F("68022"),
				State:      lithic.F("NE"),
				Address2:   lithic.F("address2"),
			}),
			Email:                lithic.F("tom@middle-earth.com"),
			FirstName:            lithic.F("Tom"),
			KYCExemptionType:     lithic.F(lithic.KYCExemptKYCExemptionTypeAuthorizedUser),
			LastName:             lithic.F("Bombadil"),
			PhoneNumber:          lithic.F("+12124007676"),
			Workflow:             lithic.F(lithic.KYCExemptWorkflowKYCExempt),
			BusinessAccountToken: lithic.F("e87db14a-4abf-4901-adad-5d5c9f46aff2"),
			ExternalID:           lithic.F("external_id"),
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
		EndingBefore:  lithic.F("ending_before"),
		ExternalID:    lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:         lithic.F(int64(0)),
		StartingAfter: lithic.F("starting_after"),
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
					City:       lithic.F("Omaha"),
					Country:    lithic.F("USA"),
					PostalCode: lithic.F("68022"),
					State:      lithic.F("NE"),
					Address2:   lithic.F("address2"),
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
		DocumentUploadToken: lithic.F("b11cd67b-0a52-4180-8365-314f3def5426"),
		Status:              lithic.F(lithic.AccountHolderSimulateEnrollmentDocumentReviewParamsStatusUploaded),
		StatusReasons:       lithic.F([]lithic.AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReason{lithic.AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentMissingRequiredData, lithic.AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonDocumentUploadTooBlurry, lithic.AccountHolderSimulateEnrollmentDocumentReviewParamsStatusReasonInvalidDocumentType}),
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
		StatusReasons:      lithic.F([]lithic.AccountHolderSimulateEnrollmentReviewParamsStatusReason{lithic.AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityIDVerificationFailure, lithic.AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityAddressVerificationFailure, lithic.AccountHolderSimulateEnrollmentReviewParamsStatusReasonPrimaryBusinessEntityNameVerificationFailure}),
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
