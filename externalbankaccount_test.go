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
			AccountNumber:         lithic.F("13719713158835300"),
			Country:               lithic.F("USA"),
			Currency:              lithic.F("USD"),
			FinancialAccountToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Owner:                 lithic.F("John Doe"),
			OwnerType:             lithic.F(lithic.OwnerTypeIndividual),
			RoutingNumber:         lithic.F("011103093"),
			Type:                  lithic.F(lithic.ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestTypeChecking),
			VerificationMethod:    lithic.F(lithic.VerificationMethodManual),
			AccountToken:          lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Address: lithic.F(lithic.ExternalBankAccountAddressParam{
				Address1:   lithic.F("5 Broad Street"),
				City:       lithic.F("New York"),
				Country:    lithic.F("USA"),
				PostalCode: lithic.F("10001"),
				State:      lithic.F("NY"),
				Address2:   lithic.F("x"),
			}),
			CompanyID:               lithic.F("sq"),
			Dob:                     lithic.F(time.Now()),
			DoingBusinessAs:         lithic.F("x"),
			Name:                    lithic.F("John Does Checking"),
			UserDefinedID:           lithic.F("x"),
			VerificationEnforcement: lithic.F(true),
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
				City:       lithic.F("x"),
				Country:    lithic.F("USD"),
				PostalCode: lithic.F("11201"),
				State:      lithic.F("xx"),
				Address2:   lithic.F("x"),
			}),
			CompanyID:       lithic.F("sq"),
			Dob:             lithic.F(time.Now()),
			DoingBusinessAs: lithic.F("x"),
			Name:            lithic.F("name"),
			Owner:           lithic.F("owner"),
			OwnerType:       lithic.F(lithic.OwnerTypeIndividual),
			Type:            lithic.F(lithic.ExternalBankAccountUpdateParamsTypeChecking),
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
		AccountTypes:       lithic.F([]lithic.ExternalBankAccountListParamsAccountType{lithic.ExternalBankAccountListParamsAccountTypeChecking}),
		Countries:          lithic.F([]string{"string"}),
		EndingBefore:       lithic.F("ending_before"),
		OwnerTypes:         lithic.F([]lithic.OwnerType{lithic.OwnerTypeIndividual}),
		PageSize:           lithic.F(int64(1)),
		StartingAfter:      lithic.F("starting_after"),
		States:             lithic.F([]lithic.ExternalBankAccountListParamsState{lithic.ExternalBankAccountListParamsStateEnabled}),
		VerificationStates: lithic.F([]lithic.ExternalBankAccountListParamsVerificationState{lithic.ExternalBankAccountListParamsVerificationStatePending}),
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

func TestExternalBankAccountRetryPrenoteWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalBankAccounts.RetryPrenote(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.ExternalBankAccountRetryPrenoteParams{
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
