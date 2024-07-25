// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// ReportSettlementService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportSettlementService] method instead.
type ReportSettlementService struct {
	Options []option.RequestOption
}

// NewReportSettlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReportSettlementService(opts ...option.RequestOption) (r *ReportSettlementService) {
	r = &ReportSettlementService{}
	r.Options = opts
	return
}

// List details.
func (r *ReportSettlementService) ListDetails(ctx context.Context, reportDate time.Time, query ReportSettlementListDetailsParams, opts ...option.RequestOption) (res *pagination.CursorPage[SettlementDetail], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("reports/settlement/details/%s", reportDate.Format("2006-01-02"))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List details.
func (r *ReportSettlementService) ListDetailsAutoPaging(ctx context.Context, reportDate time.Time, query ReportSettlementListDetailsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SettlementDetail] {
	return pagination.NewCursorPageAutoPager(r.ListDetails(ctx, reportDate, query, opts...))
}

// Get the settlement report for a specified report date. Not available in sandbox.
func (r *ReportSettlementService) Summary(ctx context.Context, reportDate time.Time, opts ...option.RequestOption) (res *SettlementReport, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("reports/settlement/summary/%s", reportDate.Format("2006-01-02"))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ReportSettlementListDetailsParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [ReportSettlementListDetailsParams]'s query parameters as
// `url.Values`.
func (r ReportSettlementListDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
