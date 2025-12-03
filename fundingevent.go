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

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// FundingEventService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFundingEventService] method instead.
type FundingEventService struct {
	Options []option.RequestOption
}

// NewFundingEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFundingEventService(opts ...option.RequestOption) (r *FundingEventService) {
	r = &FundingEventService{}
	r.Options = opts
	return
}

// Get funding event for program by id
func (r *FundingEventService) Get(ctx context.Context, fundingEventToken string, opts ...option.RequestOption) (res *FundingEvent, err error) {
	opts = slices.Concat(r.Options, opts)
	if fundingEventToken == "" {
		err = errors.New("missing required funding_event_token parameter")
		return
	}
	path := fmt.Sprintf("v1/funding_events/%s", fundingEventToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all funding events for program
func (r *FundingEventService) List(ctx context.Context, query FundingEventListParams, opts ...option.RequestOption) (res *pagination.CursorPage[FundingEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/funding_events"
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

// Get all funding events for program
func (r *FundingEventService) ListAutoPaging(ctx context.Context, query FundingEventListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[FundingEvent] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Get funding event details by id
func (r *FundingEventService) GetDetails(ctx context.Context, fundingEventToken string, opts ...option.RequestOption) (res *FundingEventGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if fundingEventToken == "" {
		err = errors.New("missing required funding_event_token parameter")
		return
	}
	path := fmt.Sprintf("v1/funding_events/%s/details", fundingEventToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FundingEvent struct {
	// Unique token ID
	Token string `json:"token,required" format:"uuid"`
	// Collection resource type
	CollectionResourceType FundingEventCollectionResourceType `json:"collection_resource_type,required"`
	// IDs of collections, further information can be gathered from the appropriate
	// collection API based on collection_resource_type
	CollectionTokens []string `json:"collection_tokens,required" format:"uuid"`
	// Time of the creation
	Created time.Time `json:"created,required" format:"date-time"`
	// Time of the high watermark
	HighWatermark time.Time `json:"high_watermark,required" format:"date-time"`
	// Network settlement summary breakdown by network settlement date
	NetworkSettlementSummary []FundingEventNetworkSettlementSummary `json:"network_settlement_summary,required"`
	// Time of the previous high watermark
	PreviousHighWatermark time.Time `json:"previous_high_watermark,required" format:"date-time"`
	// Time of the update
	Updated time.Time        `json:"updated,required" format:"date-time"`
	JSON    fundingEventJSON `json:"-"`
}

// fundingEventJSON contains the JSON metadata for the struct [FundingEvent]
type fundingEventJSON struct {
	Token                    apijson.Field
	CollectionResourceType   apijson.Field
	CollectionTokens         apijson.Field
	Created                  apijson.Field
	HighWatermark            apijson.Field
	NetworkSettlementSummary apijson.Field
	PreviousHighWatermark    apijson.Field
	Updated                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *FundingEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventJSON) RawJSON() string {
	return r.raw
}

// Collection resource type
type FundingEventCollectionResourceType string

const (
	FundingEventCollectionResourceTypeBookTransfer FundingEventCollectionResourceType = "BOOK_TRANSFER"
	FundingEventCollectionResourceTypePayment      FundingEventCollectionResourceType = "PAYMENT"
)

func (r FundingEventCollectionResourceType) IsKnown() bool {
	switch r {
	case FundingEventCollectionResourceTypeBookTransfer, FundingEventCollectionResourceTypePayment:
		return true
	}
	return false
}

type FundingEventNetworkSettlementSummary struct {
	NetworkSettlementDate time.Time                                `json:"network_settlement_date,required" format:"date"`
	SettledGrossAmount    int64                                    `json:"settled_gross_amount,required"`
	JSON                  fundingEventNetworkSettlementSummaryJSON `json:"-"`
}

// fundingEventNetworkSettlementSummaryJSON contains the JSON metadata for the
// struct [FundingEventNetworkSettlementSummary]
type fundingEventNetworkSettlementSummaryJSON struct {
	NetworkSettlementDate apijson.Field
	SettledGrossAmount    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *FundingEventNetworkSettlementSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventNetworkSettlementSummaryJSON) RawJSON() string {
	return r.raw
}

type FundingEventGetDetailsResponse struct {
	// Unique token ID
	Token string `json:"token,required" format:"uuid"`
	// URL of the settlement details
	SettlementDetailsURL string `json:"settlement_details_url,required" format:"uri"`
	// URL of the settlement summary
	SettlementSummaryURL string                             `json:"settlement_summary_url,required" format:"uri"`
	JSON                 fundingEventGetDetailsResponseJSON `json:"-"`
}

// fundingEventGetDetailsResponseJSON contains the JSON metadata for the struct
// [FundingEventGetDetailsResponse]
type fundingEventGetDetailsResponseJSON struct {
	Token                apijson.Field
	SettlementDetailsURL apijson.Field
	SettlementSummaryURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FundingEventGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type FundingEventListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [FundingEventListParams]'s query parameters as `url.Values`.
func (r FundingEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
