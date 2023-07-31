// File generated from our OpenAPI spec by Stainless.

package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/internal/testutil"
	"github.com/lithic-com/lithic-go/option"
)

func TestExternalBankAccountNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.ExternalBankAccounts.New(context.TODO(), lithic.ExternalBankAccountNewParamsPlaidCreateBankAccountAPIRequest{
		Owner:              lithic.F("x"),
		OwnerType:          lithic.F(lithic.OwnerTypeIndividual),
		ProcessorToken:     lithic.F("x"),
		VerificationMethod: lithic.F(lithic.VerificationMethodManual),
		AccountToken:       lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CompanyID:          lithic.F("x"),
		Dob:                lithic.F(time.Now()),
		DoingBusinessAs:    lithic.F("string"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
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
			DoingBusinessAs: lithic.F("string"),
			Name:            lithic.F("x"),
			Owner:           lithic.F("x"),
			OwnerType:       lithic.F(lithic.OwnerTypeIndividual),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := lithic.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.ExternalBankAccounts.List(context.TODO(), lithic.ExternalBankAccountListParams{
		AccountToken:       lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AccountTypes:       lithic.F([]lithic.ExternalBankAccountListParamsAccountType{lithic.ExternalBankAccountListParamsAccountTypeChecking, lithic.ExternalBankAccountListParamsAccountTypeChecking, lithic.ExternalBankAccountListParamsAccountTypeChecking}),
		Countries:          lithic.F([]string{"string", "string", "string"}),
		EndingBefore:       lithic.F("string"),
		OwnerTypes:         lithic.F([]lithic.OwnerType{lithic.OwnerTypeIndividual, lithic.OwnerTypeIndividual, lithic.OwnerTypeIndividual}),
		PageSize:           lithic.F(int64(1)),
		StartingAfter:      lithic.F("string"),
		States:             lithic.F([]lithic.ExternalBankAccountListParamsState{lithic.ExternalBankAccountListParamsStateEnabled, lithic.ExternalBankAccountListParamsStateEnabled, lithic.ExternalBankAccountListParamsStateEnabled}),
		VerificationStates: lithic.F([]lithic.ExternalBankAccountListParamsVerificationState{lithic.ExternalBankAccountListParamsVerificationStatePending, lithic.ExternalBankAccountListParamsVerificationStatePending, lithic.ExternalBankAccountListParamsVerificationStatePending}),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
