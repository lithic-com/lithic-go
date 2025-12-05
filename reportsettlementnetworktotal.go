// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// ReportSettlementNetworkTotalService contains methods and other services that
// help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportSettlementNetworkTotalService] method instead.
type ReportSettlementNetworkTotalService struct {
	Options []option.RequestOption
}

// NewReportSettlementNetworkTotalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewReportSettlementNetworkTotalService(opts ...option.RequestOption) (r *ReportSettlementNetworkTotalService) {
	r = &ReportSettlementNetworkTotalService{}
	r.Options = opts
	return
}

// Retrieve a specific network total record by token. Not available in sandbox.
func (r *ReportSettlementNetworkTotalService) Get(ctx context.Context, token string, opts ...option.RequestOption) (res *NetworkTotal, err error) {
	opts = slices.Concat(r.Options, opts)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("v1/reports/settlement/network_totals/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List network total records with optional filters. Not available in sandbox.
func (r *ReportSettlementNetworkTotalService) List(ctx context.Context, query ReportSettlementNetworkTotalListParams, opts ...option.RequestOption) (res *pagination.CursorPage[NetworkTotal], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/reports/settlement/network_totals"
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

// List network total records with optional filters. Not available in sandbox.
func (r *ReportSettlementNetworkTotalService) ListAutoPaging(ctx context.Context, query ReportSettlementNetworkTotalListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[NetworkTotal] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type ReportSettlementNetworkTotalListParams struct {
	// Datetime in RFC 3339 format. Only entries created after the specified time will
	// be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Datetime in RFC 3339 format. Only entries created before the specified time will
	// be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Institution ID to filter on.
	InstitutionID param.Field[string] `query:"institution_id"`
	// Network to filter on.
	Network param.Field[ReportSettlementNetworkTotalListParamsNetwork] `query:"network"`
	// Number of records per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Singular report date to filter on (YYYY-MM-DD). Cannot be populated in
	// conjunction with report_date_begin or report_date_end.
	ReportDate param.Field[time.Time] `query:"report_date" format:"date"`
	// Earliest report date to filter on, inclusive (YYYY-MM-DD).
	ReportDateBegin param.Field[time.Time] `query:"report_date_begin" format:"date"`
	// Latest report date to filter on, inclusive (YYYY-MM-DD).
	ReportDateEnd param.Field[time.Time] `query:"report_date_end" format:"date"`
	// Settlement institution ID to filter on.
	SettlementInstitutionID param.Field[string] `query:"settlement_institution_id"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [ReportSettlementNetworkTotalListParams]'s query parameters
// as `url.Values`.
func (r ReportSettlementNetworkTotalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Network to filter on.
type ReportSettlementNetworkTotalListParamsNetwork string

const (
	ReportSettlementNetworkTotalListParamsNetworkVisa       ReportSettlementNetworkTotalListParamsNetwork = "VISA"
	ReportSettlementNetworkTotalListParamsNetworkMastercard ReportSettlementNetworkTotalListParamsNetwork = "MASTERCARD"
	ReportSettlementNetworkTotalListParamsNetworkMaestro    ReportSettlementNetworkTotalListParamsNetwork = "MAESTRO"
	ReportSettlementNetworkTotalListParamsNetworkInterlink  ReportSettlementNetworkTotalListParamsNetwork = "INTERLINK"
)

func (r ReportSettlementNetworkTotalListParamsNetwork) IsKnown() bool {
	switch r {
	case ReportSettlementNetworkTotalListParamsNetworkVisa, ReportSettlementNetworkTotalListParamsNetworkMastercard, ReportSettlementNetworkTotalListParamsNetworkMaestro, ReportSettlementNetworkTotalListParamsNetworkInterlink:
		return true
	}
	return false
}
