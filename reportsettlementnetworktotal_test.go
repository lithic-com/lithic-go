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

func TestReportSettlementNetworkTotalGet(t *testing.T) {
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
	_, err := client.Reports.Settlement.NetworkTotals.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReportSettlementNetworkTotalListWithOptionalParams(t *testing.T) {
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
	_, err := client.Reports.Settlement.NetworkTotals.List(context.TODO(), lithic.ReportSettlementNetworkTotalListParams{
		Begin:                   lithic.F(time.Now()),
		End:                     lithic.F(time.Now()),
		EndingBefore:            lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		InstitutionID:           lithic.F("institution_id"),
		Network:                 lithic.F(lithic.ReportSettlementNetworkTotalListParamsNetworkVisa),
		PageSize:                lithic.F(int64(1)),
		ReportDate:              lithic.F(time.Now()),
		ReportDateBegin:         lithic.F(time.Now()),
		ReportDateEnd:           lithic.F(time.Now()),
		SettlementInstitutionID: lithic.F("settlement_institution_id"),
		StartingAfter:           lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
