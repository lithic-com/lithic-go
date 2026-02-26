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
)

func TestAccountHolderEntityNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountHolders.Entities.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.AccountHolderEntityNewParams{
			Address: lithic.F(lithic.AccountHolderEntityNewParamsAddress{
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
			Type:         lithic.F(lithic.AccountHolderEntityNewParamsTypeBeneficialOwnerIndividual),
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

func TestAccountHolderEntityDelete(t *testing.T) {
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
	_, err := client.AccountHolders.Entities.Delete(
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
