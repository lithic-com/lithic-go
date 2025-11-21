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

// CardProgramService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardProgramService] method instead.
type CardProgramService struct {
	Options []option.RequestOption
}

// NewCardProgramService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardProgramService(opts ...option.RequestOption) (r *CardProgramService) {
	r = &CardProgramService{}
	r.Options = opts
	return
}

// Get card program.
func (r *CardProgramService) Get(ctx context.Context, cardProgramToken string, opts ...option.RequestOption) (res *CardProgram, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardProgramToken == "" {
		err = errors.New("missing required card_program_token parameter")
		return
	}
	path := fmt.Sprintf("v1/card_programs/%s", cardProgramToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List card programs.
func (r *CardProgramService) List(ctx context.Context, query CardProgramListParams, opts ...option.RequestOption) (res *pagination.CursorPage[CardProgram], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/card_programs"
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

// List card programs.
func (r *CardProgramService) ListAutoPaging(ctx context.Context, query CardProgramListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CardProgram] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type CardProgram struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Whether the card program is participating in Account Level Management. Currently
	// applicable to Visa card programs only.
	AccountLevelManagementEnabled bool `json:"account_level_management_enabled,required"`
	// Timestamp of when the card program was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// The name of the card program.
	Name string `json:"name,required"`
	// The first digits of the card number that this card program ends with.
	PanRangeEnd string `json:"pan_range_end,required"`
	// The first digits of the card number that this card program starts with.
	PanRangeStart string `json:"pan_range_start,required"`
	// 3-character alphabetic ISO 4217 code for the currency of the cardholder.
	CardholderCurrency string `json:"cardholder_currency"`
	// List of 3-character alphabetic ISO 4217 codes for the currencies that the card
	// program supports for settlement.
	SettlementCurrencies []string        `json:"settlement_currencies"`
	JSON                 cardProgramJSON `json:"-"`
}

// cardProgramJSON contains the JSON metadata for the struct [CardProgram]
type cardProgramJSON struct {
	Token                         apijson.Field
	AccountLevelManagementEnabled apijson.Field
	Created                       apijson.Field
	Name                          apijson.Field
	PanRangeEnd                   apijson.Field
	PanRangeStart                 apijson.Field
	CardholderCurrency            apijson.Field
	SettlementCurrencies          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *CardProgram) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardProgramJSON) RawJSON() string {
	return r.raw
}

type CardProgramListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [CardProgramListParams]'s query parameters as `url.Values`.
func (r CardProgramListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
