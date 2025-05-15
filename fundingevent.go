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
func (r *FundingEventService) Get(ctx context.Context, fundingEventToken string, opts ...option.RequestOption) (res *FundingEventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fundingEventToken == "" {
		err = errors.New("missing required funding_event_token parameter")
		return
	}
	path := fmt.Sprintf("v1/funding_events/%s", fundingEventToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all funding events for program
func (r *FundingEventService) List(ctx context.Context, query FundingEventListParams, opts ...option.RequestOption) (res *pagination.CursorPage[FundingEventListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *FundingEventService) ListAutoPaging(ctx context.Context, query FundingEventListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[FundingEventListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Get funding event details by id
func (r *FundingEventService) GetDetails(ctx context.Context, fundingEventToken string, opts ...option.RequestOption) (res *FundingEventGetDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fundingEventToken == "" {
		err = errors.New("missing required funding_event_token parameter")
		return
	}
	path := fmt.Sprintf("v1/funding_events/%s/details", fundingEventToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FundingEventGetResponse struct {
	// Unique token ID
	Token string `json:"token,required" format:"uuid"`
	// Collection resource type
	CollectionResourceType FundingEventGetResponseCollectionResourceType `json:"collection_resource_type,required"`
	// IDs of collections
	CollectionTokens []string `json:"collection_tokens,required" format:"uuid"`
	// Time of the creation
	Created time.Time `json:"created,required" format:"date-time"`
	// Time of the high watermark
	HighWatermark time.Time `json:"high_watermark,required" format:"date-time"`
	// Time of the previous high watermark
	PreviousHighWatermark time.Time `json:"previous_high_watermark,required" format:"date-time"`
	// List of settlements
	SettlementBreakdowns []FundingEventGetResponseSettlementBreakdown `json:"settlement_breakdowns,required"`
	// Time of the update
	Updated time.Time                   `json:"updated,required" format:"date-time"`
	JSON    fundingEventGetResponseJSON `json:"-"`
}

// fundingEventGetResponseJSON contains the JSON metadata for the struct
// [FundingEventGetResponse]
type fundingEventGetResponseJSON struct {
	Token                  apijson.Field
	CollectionResourceType apijson.Field
	CollectionTokens       apijson.Field
	Created                apijson.Field
	HighWatermark          apijson.Field
	PreviousHighWatermark  apijson.Field
	SettlementBreakdowns   apijson.Field
	Updated                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *FundingEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventGetResponseJSON) RawJSON() string {
	return r.raw
}

// Collection resource type
type FundingEventGetResponseCollectionResourceType string

const (
	FundingEventGetResponseCollectionResourceTypeBookTransfer FundingEventGetResponseCollectionResourceType = "BOOK_TRANSFER"
	FundingEventGetResponseCollectionResourceTypePayment      FundingEventGetResponseCollectionResourceType = "PAYMENT"
)

func (r FundingEventGetResponseCollectionResourceType) IsKnown() bool {
	switch r {
	case FundingEventGetResponseCollectionResourceTypeBookTransfer, FundingEventGetResponseCollectionResourceTypePayment:
		return true
	}
	return false
}

type FundingEventGetResponseSettlementBreakdown struct {
	Amount         int64                                          `json:"amount,required"`
	SettlementDate time.Time                                      `json:"settlement_date,required" format:"date"`
	JSON           fundingEventGetResponseSettlementBreakdownJSON `json:"-"`
}

// fundingEventGetResponseSettlementBreakdownJSON contains the JSON metadata for
// the struct [FundingEventGetResponseSettlementBreakdown]
type fundingEventGetResponseSettlementBreakdownJSON struct {
	Amount         apijson.Field
	SettlementDate apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FundingEventGetResponseSettlementBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventGetResponseSettlementBreakdownJSON) RawJSON() string {
	return r.raw
}

type FundingEventListResponse struct {
	// Unique token ID
	Token string `json:"token,required" format:"uuid"`
	// Collection resource type
	CollectionResourceType FundingEventListResponseCollectionResourceType `json:"collection_resource_type,required"`
	// IDs of collections
	CollectionTokens []string `json:"collection_tokens,required" format:"uuid"`
	// Time of the creation
	Created time.Time `json:"created,required" format:"date-time"`
	// Time of the high watermark
	HighWatermark time.Time `json:"high_watermark,required" format:"date-time"`
	// Time of the previous high watermark
	PreviousHighWatermark time.Time `json:"previous_high_watermark,required" format:"date-time"`
	// List of settlements
	SettlementBreakdowns []FundingEventListResponseSettlementBreakdown `json:"settlement_breakdowns,required"`
	// Time of the update
	Updated time.Time                    `json:"updated,required" format:"date-time"`
	JSON    fundingEventListResponseJSON `json:"-"`
}

// fundingEventListResponseJSON contains the JSON metadata for the struct
// [FundingEventListResponse]
type fundingEventListResponseJSON struct {
	Token                  apijson.Field
	CollectionResourceType apijson.Field
	CollectionTokens       apijson.Field
	Created                apijson.Field
	HighWatermark          apijson.Field
	PreviousHighWatermark  apijson.Field
	SettlementBreakdowns   apijson.Field
	Updated                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *FundingEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventListResponseJSON) RawJSON() string {
	return r.raw
}

// Collection resource type
type FundingEventListResponseCollectionResourceType string

const (
	FundingEventListResponseCollectionResourceTypeBookTransfer FundingEventListResponseCollectionResourceType = "BOOK_TRANSFER"
	FundingEventListResponseCollectionResourceTypePayment      FundingEventListResponseCollectionResourceType = "PAYMENT"
)

func (r FundingEventListResponseCollectionResourceType) IsKnown() bool {
	switch r {
	case FundingEventListResponseCollectionResourceTypeBookTransfer, FundingEventListResponseCollectionResourceTypePayment:
		return true
	}
	return false
}

type FundingEventListResponseSettlementBreakdown struct {
	Amount         int64                                           `json:"amount,required"`
	SettlementDate time.Time                                       `json:"settlement_date,required" format:"date"`
	JSON           fundingEventListResponseSettlementBreakdownJSON `json:"-"`
}

// fundingEventListResponseSettlementBreakdownJSON contains the JSON metadata for
// the struct [FundingEventListResponseSettlementBreakdown]
type fundingEventListResponseSettlementBreakdownJSON struct {
	Amount         apijson.Field
	SettlementDate apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FundingEventListResponseSettlementBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fundingEventListResponseSettlementBreakdownJSON) RawJSON() string {
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
