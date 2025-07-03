// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
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
func (r *ReportSettlementNetworkTotalService) Get(ctx context.Context, token string, opts ...option.RequestOption) (res *ReportSettlementNetworkTotalGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("v1/reports/settlement/network_totals/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List network total records with optional filters. Not available in sandbox.
func (r *ReportSettlementNetworkTotalService) List(ctx context.Context, query ReportSettlementNetworkTotalListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ReportSettlementNetworkTotalListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *ReportSettlementNetworkTotalService) ListAutoPaging(ctx context.Context, query ReportSettlementNetworkTotalListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ReportSettlementNetworkTotalListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type ReportSettlementNetworkTotalGetResponse struct {
	// Globally unique identifier.
	Token   string                                         `json:"token,required" format:"uuid"`
	Amounts ReportSettlementNetworkTotalGetResponseAmounts `json:"amounts,required"`
	// RFC 3339 timestamp for when the record was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-character alphabetic ISO 4217 code.
	Currency string `json:"currency,required"`
	// The institution that activity occurred on. For Mastercard: ICA (Interbank Card
	// Association). For Maestro: institution ID. For Visa: lowest level SRE
	// (Settlement Reporting Entity).
	InstitutionID string `json:"institution_id,required"`
	// Indicates that all settlement records related to this Network Total are
	// available in the details endpoint.
	IsComplete bool `json:"is_complete,required"`
	// Card network where the transaction took place. AMEX, VISA, MASTERCARD, MAESTRO,
	// or INTERLINK.
	Network ReportSettlementNetworkTotalGetResponseNetwork `json:"network,required"`
	// Date that the network total record applies to. YYYY-MM-DD format.
	ReportDate time.Time `json:"report_date,required" format:"date"`
	// The institution responsible for settlement. For Mastercard: same as
	// `institution_id`. For Maestro: billing ICA. For Visa: Funds Transfer SRE
	// (FTSRE).
	SettlementInstitutionID string `json:"settlement_institution_id,required"`
	// Settlement service.
	SettlementService string `json:"settlement_service,required"`
	// RFC 3339 timestamp for when the record was last updated. UTC time zone.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// The clearing cycle that the network total record applies to. Mastercard only.
	Cycle int64                                       `json:"cycle"`
	JSON  reportSettlementNetworkTotalGetResponseJSON `json:"-"`
}

// reportSettlementNetworkTotalGetResponseJSON contains the JSON metadata for the
// struct [ReportSettlementNetworkTotalGetResponse]
type reportSettlementNetworkTotalGetResponseJSON struct {
	Token                   apijson.Field
	Amounts                 apijson.Field
	Created                 apijson.Field
	Currency                apijson.Field
	InstitutionID           apijson.Field
	IsComplete              apijson.Field
	Network                 apijson.Field
	ReportDate              apijson.Field
	SettlementInstitutionID apijson.Field
	SettlementService       apijson.Field
	Updated                 apijson.Field
	Cycle                   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ReportSettlementNetworkTotalGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportSettlementNetworkTotalGetResponseJSON) RawJSON() string {
	return r.raw
}

type ReportSettlementNetworkTotalGetResponseAmounts struct {
	// Total settlement amount excluding interchange, in currency's smallest unit.
	GrossSettlement int64 `json:"gross_settlement,required"`
	// Interchange amount, in currency's smallest unit.
	InterchangeFees int64 `json:"interchange_fees,required"`
	// `gross_settlement` net of `interchange_fees` and `visa_charges` (if applicable),
	// in currency's smallest unit.
	NetSettlement int64 `json:"net_settlement,required"`
	// Charges specific to Visa/Interlink, in currency's smallest unit.
	VisaCharges int64                                              `json:"visa_charges"`
	JSON        reportSettlementNetworkTotalGetResponseAmountsJSON `json:"-"`
}

// reportSettlementNetworkTotalGetResponseAmountsJSON contains the JSON metadata
// for the struct [ReportSettlementNetworkTotalGetResponseAmounts]
type reportSettlementNetworkTotalGetResponseAmountsJSON struct {
	GrossSettlement apijson.Field
	InterchangeFees apijson.Field
	NetSettlement   apijson.Field
	VisaCharges     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ReportSettlementNetworkTotalGetResponseAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportSettlementNetworkTotalGetResponseAmountsJSON) RawJSON() string {
	return r.raw
}

// Card network where the transaction took place. AMEX, VISA, MASTERCARD, MAESTRO,
// or INTERLINK.
type ReportSettlementNetworkTotalGetResponseNetwork string

const (
	ReportSettlementNetworkTotalGetResponseNetworkAmex       ReportSettlementNetworkTotalGetResponseNetwork = "AMEX"
	ReportSettlementNetworkTotalGetResponseNetworkVisa       ReportSettlementNetworkTotalGetResponseNetwork = "VISA"
	ReportSettlementNetworkTotalGetResponseNetworkMastercard ReportSettlementNetworkTotalGetResponseNetwork = "MASTERCARD"
	ReportSettlementNetworkTotalGetResponseNetworkMaestro    ReportSettlementNetworkTotalGetResponseNetwork = "MAESTRO"
	ReportSettlementNetworkTotalGetResponseNetworkInterlink  ReportSettlementNetworkTotalGetResponseNetwork = "INTERLINK"
)

func (r ReportSettlementNetworkTotalGetResponseNetwork) IsKnown() bool {
	switch r {
	case ReportSettlementNetworkTotalGetResponseNetworkAmex, ReportSettlementNetworkTotalGetResponseNetworkVisa, ReportSettlementNetworkTotalGetResponseNetworkMastercard, ReportSettlementNetworkTotalGetResponseNetworkMaestro, ReportSettlementNetworkTotalGetResponseNetworkInterlink:
		return true
	}
	return false
}

type ReportSettlementNetworkTotalListResponse struct {
	// Globally unique identifier.
	Token   string                                          `json:"token,required" format:"uuid"`
	Amounts ReportSettlementNetworkTotalListResponseAmounts `json:"amounts,required"`
	// RFC 3339 timestamp for when the record was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-character alphabetic ISO 4217 code.
	Currency string `json:"currency,required"`
	// The institution that activity occurred on. For Mastercard: ICA (Interbank Card
	// Association). For Maestro: institution ID. For Visa: lowest level SRE
	// (Settlement Reporting Entity).
	InstitutionID string `json:"institution_id,required"`
	// Indicates that all settlement records related to this Network Total are
	// available in the details endpoint.
	IsComplete bool `json:"is_complete,required"`
	// Card network where the transaction took place. AMEX, VISA, MASTERCARD, MAESTRO,
	// or INTERLINK.
	Network ReportSettlementNetworkTotalListResponseNetwork `json:"network,required"`
	// Date that the network total record applies to. YYYY-MM-DD format.
	ReportDate time.Time `json:"report_date,required" format:"date"`
	// The institution responsible for settlement. For Mastercard: same as
	// `institution_id`. For Maestro: billing ICA. For Visa: Funds Transfer SRE
	// (FTSRE).
	SettlementInstitutionID string `json:"settlement_institution_id,required"`
	// Settlement service.
	SettlementService string `json:"settlement_service,required"`
	// RFC 3339 timestamp for when the record was last updated. UTC time zone.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// The clearing cycle that the network total record applies to. Mastercard only.
	Cycle int64                                        `json:"cycle"`
	JSON  reportSettlementNetworkTotalListResponseJSON `json:"-"`
}

// reportSettlementNetworkTotalListResponseJSON contains the JSON metadata for the
// struct [ReportSettlementNetworkTotalListResponse]
type reportSettlementNetworkTotalListResponseJSON struct {
	Token                   apijson.Field
	Amounts                 apijson.Field
	Created                 apijson.Field
	Currency                apijson.Field
	InstitutionID           apijson.Field
	IsComplete              apijson.Field
	Network                 apijson.Field
	ReportDate              apijson.Field
	SettlementInstitutionID apijson.Field
	SettlementService       apijson.Field
	Updated                 apijson.Field
	Cycle                   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ReportSettlementNetworkTotalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportSettlementNetworkTotalListResponseJSON) RawJSON() string {
	return r.raw
}

type ReportSettlementNetworkTotalListResponseAmounts struct {
	// Total settlement amount excluding interchange, in currency's smallest unit.
	GrossSettlement int64 `json:"gross_settlement,required"`
	// Interchange amount, in currency's smallest unit.
	InterchangeFees int64 `json:"interchange_fees,required"`
	// `gross_settlement` net of `interchange_fees` and `visa_charges` (if applicable),
	// in currency's smallest unit.
	NetSettlement int64 `json:"net_settlement,required"`
	// Charges specific to Visa/Interlink, in currency's smallest unit.
	VisaCharges int64                                               `json:"visa_charges"`
	JSON        reportSettlementNetworkTotalListResponseAmountsJSON `json:"-"`
}

// reportSettlementNetworkTotalListResponseAmountsJSON contains the JSON metadata
// for the struct [ReportSettlementNetworkTotalListResponseAmounts]
type reportSettlementNetworkTotalListResponseAmountsJSON struct {
	GrossSettlement apijson.Field
	InterchangeFees apijson.Field
	NetSettlement   apijson.Field
	VisaCharges     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ReportSettlementNetworkTotalListResponseAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportSettlementNetworkTotalListResponseAmountsJSON) RawJSON() string {
	return r.raw
}

// Card network where the transaction took place. AMEX, VISA, MASTERCARD, MAESTRO,
// or INTERLINK.
type ReportSettlementNetworkTotalListResponseNetwork string

const (
	ReportSettlementNetworkTotalListResponseNetworkAmex       ReportSettlementNetworkTotalListResponseNetwork = "AMEX"
	ReportSettlementNetworkTotalListResponseNetworkVisa       ReportSettlementNetworkTotalListResponseNetwork = "VISA"
	ReportSettlementNetworkTotalListResponseNetworkMastercard ReportSettlementNetworkTotalListResponseNetwork = "MASTERCARD"
	ReportSettlementNetworkTotalListResponseNetworkMaestro    ReportSettlementNetworkTotalListResponseNetwork = "MAESTRO"
	ReportSettlementNetworkTotalListResponseNetworkInterlink  ReportSettlementNetworkTotalListResponseNetwork = "INTERLINK"
)

func (r ReportSettlementNetworkTotalListResponseNetwork) IsKnown() bool {
	switch r {
	case ReportSettlementNetworkTotalListResponseNetworkAmex, ReportSettlementNetworkTotalListResponseNetworkVisa, ReportSettlementNetworkTotalListResponseNetworkMastercard, ReportSettlementNetworkTotalListResponseNetworkMaestro, ReportSettlementNetworkTotalListResponseNetworkInterlink:
		return true
	}
	return false
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
