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
)

func TestExternalBankAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalBankAccounts.New(context.TODO(), lithic.ExternalBankAccountNewParams{
		Body: lithic.ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest{
			VerificationMethod:      lithic.F(lithic.VerificationMethodManual),
			OwnerType:               lithic.F(lithic.OwnerTypeIndividual),
			Owner:                   lithic.F("John Doe"),
			AccountToken:            lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CompanyID:               lithic.F("x"),
			DoingBusinessAs:         lithic.F("x"),
			Dob:                     lithic.F(time.Now()),
			UserDefinedID:           lithic.F("x"),
			Type:                    lithic.F(lithic.ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestTypeChecking),
			RoutingNumber:           lithic.F("011103093"),
			AccountNumber:           lithic.F("13719713158835300"),
			Name:                    lithic.F("John Does Checking"),
			Country:                 lithic.F("USA"),
			Currency:                lithic.F("USD"),
			VerificationEnforcement: lithic.F(true),
			Address: lithic.F(lithic.ExternalBankAccountAddressParam{
				Address1:   lithic.F("5 Broad Street"),
				Address2:   lithic.F("x"),
				City:       lithic.F("New York"),
				State:      lithic.F("NY"),
				PostalCode: lithic.F("10001"),
				Country:    lithic.F("USA"),
			}),
			FinancialAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestExternalBankAccountGet(t *testing.T) {
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
	_, err := client.ExternalBankAccounts.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalBankAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalBankAccounts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.ExternalBankAccountUpdateParams{
			Address: lithic.F(lithic.ExternalBankAccountAddressParam{
				Address1:   lithic.F("x"),
				Address2:   lithic.F("x"),
				City:       lithic.F("x"),
				State:      lithic.F("xx"),
				PostalCode: lithic.F("11201"),
				Country:    lithic.F("USD"),
			}),
			CompanyID:       lithic.F("x"),
			Dob:             lithic.F(time.Now()),
			DoingBusinessAs: lithic.F("x"),
			Name:            lithic.F("x"),
			Owner:           lithic.F("x"),
			OwnerType:       lithic.F(lithic.OwnerTypeIndividual),
			UserDefinedID:   lithic.F("x"),
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

func TestExternalBankAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalBankAccounts.List(context.TODO(), lithic.ExternalBankAccountListParams{
		AccountToken:       lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AccountTypes:       lithic.F([]lithic.ExternalBankAccountListParamsAccountType{lithic.ExternalBankAccountListParamsAccountTypeChecking, lithic.ExternalBankAccountListParamsAccountTypeSavings}),
		Countries:          lithic.F([]string{"string", "string", "string"}),
		EndingBefore:       lithic.F("string"),
		OwnerTypes:         lithic.F([]lithic.OwnerType{lithic.OwnerTypeIndividual, lithic.OwnerTypeBusiness}),
		PageSize:           lithic.F(int64(1)),
		StartingAfter:      lithic.F("string"),
		States:             lithic.F([]lithic.ExternalBankAccountListParamsState{lithic.ExternalBankAccountListParamsStateEnabled, lithic.ExternalBankAccountListParamsStateClosed, lithic.ExternalBankAccountListParamsStatePaused}),
		VerificationStates: lithic.F([]lithic.ExternalBankAccountListParamsVerificationState{lithic.ExternalBankAccountListParamsVerificationStatePending, lithic.ExternalBankAccountListParamsVerificationStateEnabled, lithic.ExternalBankAccountListParamsVerificationStateFailedVerification}),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalBankAccountRetryMicroDepositsWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalBankAccounts.RetryMicroDeposits(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.ExternalBankAccountRetryMicroDepositsParams{
			FinancialAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
