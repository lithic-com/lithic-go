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
			AccountNumber: lithic.F("13719713158835300"),
			AccountToken:  lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Address: lithic.F(lithic.ExternalBankAccountAddressParam{
				Address1:   lithic.F("5 Broad Street"),
				Address2:   lithic.F("x"),
				City:       lithic.F("New York"),
				Country:    lithic.F("USA"),
				PostalCode: lithic.F("10001"),
				State:      lithic.F("NY"),
			}),
			CompanyID:               lithic.F("x"),
			Country:                 lithic.F("USA"),
			Currency:                lithic.F("USD"),
			Dob:                     lithic.F(time.Now()),
			DoingBusinessAs:         lithic.F("string"),
			FinancialAccountToken:   lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:                    lithic.F("John Does Checking"),
			Owner:                   lithic.F("John Doe"),
			OwnerType:               lithic.F(lithic.OwnerTypeBusiness),
			RoutingNumber:           lithic.F("011103093"),
			Type:                    lithic.F(lithic.ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestTypeChecking),
			UserDefinedID:           lithic.F("string"),
			VerificationEnforcement: lithic.F(true),
			VerificationMethod:      lithic.F(lithic.VerificationMethodMicroDeposit),
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
				Country:    lithic.F("USD"),
				PostalCode: lithic.F("11201"),
				State:      lithic.F("xx"),
			}),
			CompanyID:       lithic.F("x"),
			Dob:             lithic.F(time.Now()),
			DoingBusinessAs: lithic.F("string"),
			Name:            lithic.F("x"),
			Owner:           lithic.F("x"),
			OwnerType:       lithic.F(lithic.OwnerTypeBusiness),
			UserDefinedID:   lithic.F("string"),
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
		OwnerTypes:         lithic.F([]lithic.OwnerType{lithic.OwnerTypeBusiness, lithic.OwnerTypeIndividual}),
		PageSize:           lithic.F(int64(1)),
		StartingAfter:      lithic.F("string"),
		States:             lithic.F([]lithic.ExternalBankAccountListParamsState{lithic.ExternalBankAccountListParamsStateClosed, lithic.ExternalBankAccountListParamsStateEnabled, lithic.ExternalBankAccountListParamsStatePaused}),
		VerificationStates: lithic.F([]lithic.ExternalBankAccountListParamsVerificationState{lithic.ExternalBankAccountListParamsVerificationStateEnabled, lithic.ExternalBankAccountListParamsVerificationStateFailedVerification, lithic.ExternalBankAccountListParamsVerificationStateInsufficientFunds}),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalBankAccountRetryMicroDeposits(t *testing.T) {
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
	_, err := client.ExternalBankAccounts.RetryMicroDeposits(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
