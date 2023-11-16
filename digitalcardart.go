// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// DigitalCardArtService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDigitalCardArtService] method
// instead.
type DigitalCardArtService struct {
	Options []option.RequestOption
}

// NewDigitalCardArtService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDigitalCardArtService(opts ...option.RequestOption) (r *DigitalCardArtService) {
	r = &DigitalCardArtService{}
	r.Options = opts
	return
}

// List digital card art.
func (r *DigitalCardArtService) List(ctx context.Context, query DigitalCardArtListParams, opts ...option.RequestOption) (res *shared.CursorPage[DigitalCardArt], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "digital_card_art"
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

// List digital card art.
func (r *DigitalCardArtService) ListAutoPaging(ctx context.Context, query DigitalCardArtListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[DigitalCardArt] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type DigitalCardArt struct {
	// Globally unique identifier for the card art.
	Token string `json:"token,required" format:"uuid"`
	// Globally unique identifier for the card program.
	CardProgramToken string `json:"card_program_token,required" format:"uuid"`
	// Timestamp of when card art was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Description of the card art.
	Description string `json:"description,required"`
	// Whether the card art is enabled.
	IsEnabled bool `json:"is_enabled,required"`
	// Card network.
	Network DigitalCardArtNetwork `json:"network,required"`
	// Whether the card art is the default card art to be added upon tokenization.
	IsCardProgramDefault bool               `json:"is_card_program_default"`
	JSON                 digitalCardArtJSON `json:"-"`
}

// digitalCardArtJSON contains the JSON metadata for the struct [DigitalCardArt]
type digitalCardArtJSON struct {
	Token                apijson.Field
	CardProgramToken     apijson.Field
	Created              apijson.Field
	Description          apijson.Field
	IsEnabled            apijson.Field
	Network              apijson.Field
	IsCardProgramDefault apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DigitalCardArt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Card network.
type DigitalCardArtNetwork string

const (
	DigitalCardArtNetworkMastercard DigitalCardArtNetwork = "MASTERCARD"
	DigitalCardArtNetworkVisa       DigitalCardArtNetwork = "VISA"
)

type DigitalCardArtListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [DigitalCardArtListParams]'s query parameters as
// `url.Values`.
func (r DigitalCardArtListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
