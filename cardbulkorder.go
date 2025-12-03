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

// CardBulkOrderService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardBulkOrderService] method instead.
type CardBulkOrderService struct {
	Options []option.RequestOption
}

// NewCardBulkOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardBulkOrderService(opts ...option.RequestOption) (r *CardBulkOrderService) {
	r = &CardBulkOrderService{}
	r.Options = opts
	return
}

// Create a new bulk order for physical card shipments **[BETA]**. Cards can be
// added to the order via the POST /v1/cards endpoint by specifying the
// bulk_order_token. Lock the order via PATCH
// /v1/card_bulk_orders/{bulk_order_token} to prepare for shipment. Please work
// with your Customer Success Manager and card personalization bureau to ensure
// bulk shipping is supported for your program.
func (r *CardBulkOrderService) New(ctx context.Context, body CardBulkOrderNewParams, opts ...option.RequestOption) (res *CardBulkOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/card_bulk_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific bulk order by token **[BETA]**
func (r *CardBulkOrderService) Get(ctx context.Context, bulkOrderToken string, opts ...option.RequestOption) (res *CardBulkOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	if bulkOrderToken == "" {
		err = errors.New("missing required bulk_order_token parameter")
		return
	}
	path := fmt.Sprintf("v1/card_bulk_orders/%s", bulkOrderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a bulk order **[BETA]**. Primarily used to lock the order, preventing
// additional cards from being added
func (r *CardBulkOrderService) Update(ctx context.Context, bulkOrderToken string, body CardBulkOrderUpdateParams, opts ...option.RequestOption) (res *CardBulkOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	if bulkOrderToken == "" {
		err = errors.New("missing required bulk_order_token parameter")
		return
	}
	path := fmt.Sprintf("v1/card_bulk_orders/%s", bulkOrderToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List bulk orders for physical card shipments **[BETA]**
func (r *CardBulkOrderService) List(ctx context.Context, query CardBulkOrderListParams, opts ...option.RequestOption) (res *pagination.CursorPage[CardBulkOrder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/card_bulk_orders"
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

// List bulk orders for physical card shipments **[BETA]**
func (r *CardBulkOrderService) ListAutoPaging(ctx context.Context, query CardBulkOrderListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CardBulkOrder] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Represents a bulk order for physical card shipments
type CardBulkOrder struct {
	// Globally unique identifier for the bulk order
	Token string `json:"token,required" format:"uuid"`
	// List of card tokens associated with this bulk order
	CardTokens []string `json:"card_tokens,required" format:"uuid"`
	// An RFC 3339 timestamp for when the bulk order was created. UTC time zone
	Created time.Time `json:"created,required" format:"date-time"`
	// Customer-specified product configuration for physical card manufacturing. This
	// must be configured with Lithic before use
	CustomerProductID string `json:"customer_product_id,required,nullable"`
	// Shipping address for all cards in this bulk order
	ShippingAddress interface{} `json:"shipping_address,required"`
	// Shipping method for all cards in this bulk order
	ShippingMethod CardBulkOrderShippingMethod `json:"shipping_method,required"`
	// Status of the bulk order. OPEN indicates the order is accepting cards. LOCKED
	// indicates the order is finalized and no more cards can be added
	Status CardBulkOrderStatus `json:"status,required"`
	// An RFC 3339 timestamp for when the bulk order was last updated. UTC time zone
	Updated time.Time         `json:"updated,required" format:"date-time"`
	JSON    cardBulkOrderJSON `json:"-"`
}

// cardBulkOrderJSON contains the JSON metadata for the struct [CardBulkOrder]
type cardBulkOrderJSON struct {
	Token             apijson.Field
	CardTokens        apijson.Field
	Created           apijson.Field
	CustomerProductID apijson.Field
	ShippingAddress   apijson.Field
	ShippingMethod    apijson.Field
	Status            apijson.Field
	Updated           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CardBulkOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardBulkOrderJSON) RawJSON() string {
	return r.raw
}

// Shipping method for all cards in this bulk order
type CardBulkOrderShippingMethod string

const (
	CardBulkOrderShippingMethodBulkExpedited CardBulkOrderShippingMethod = "BULK_EXPEDITED"
)

func (r CardBulkOrderShippingMethod) IsKnown() bool {
	switch r {
	case CardBulkOrderShippingMethodBulkExpedited:
		return true
	}
	return false
}

// Status of the bulk order. OPEN indicates the order is accepting cards. LOCKED
// indicates the order is finalized and no more cards can be added
type CardBulkOrderStatus string

const (
	CardBulkOrderStatusOpen   CardBulkOrderStatus = "OPEN"
	CardBulkOrderStatusLocked CardBulkOrderStatus = "LOCKED"
)

func (r CardBulkOrderStatus) IsKnown() bool {
	switch r {
	case CardBulkOrderStatusOpen, CardBulkOrderStatusLocked:
		return true
	}
	return false
}

type CardBulkOrderNewParams struct {
	// Customer-specified product configuration for physical card manufacturing. This
	// must be configured with Lithic before use
	CustomerProductID param.Field[string] `json:"customer_product_id,required"`
	// Shipping address for all cards in this bulk order
	ShippingAddress param.Field[interface{}] `json:"shipping_address,required"`
	// Shipping method for all cards in this bulk order
	ShippingMethod param.Field[CardBulkOrderNewParamsShippingMethod] `json:"shipping_method,required"`
}

func (r CardBulkOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Shipping method for all cards in this bulk order
type CardBulkOrderNewParamsShippingMethod string

const (
	CardBulkOrderNewParamsShippingMethodBulkExpedited CardBulkOrderNewParamsShippingMethod = "BULK_EXPEDITED"
)

func (r CardBulkOrderNewParamsShippingMethod) IsKnown() bool {
	switch r {
	case CardBulkOrderNewParamsShippingMethodBulkExpedited:
		return true
	}
	return false
}

type CardBulkOrderUpdateParams struct {
	// Status to update the bulk order to. Use LOCKED to finalize the order
	Status param.Field[CardBulkOrderUpdateParamsStatus] `json:"status,required"`
}

func (r CardBulkOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status to update the bulk order to. Use LOCKED to finalize the order
type CardBulkOrderUpdateParamsStatus string

const (
	CardBulkOrderUpdateParamsStatusLocked CardBulkOrderUpdateParamsStatus = "LOCKED"
)

func (r CardBulkOrderUpdateParamsStatus) IsKnown() bool {
	switch r {
	case CardBulkOrderUpdateParamsStatusLocked:
		return true
	}
	return false
}

type CardBulkOrderListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [CardBulkOrderListParams]'s query parameters as
// `url.Values`.
func (r CardBulkOrderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
