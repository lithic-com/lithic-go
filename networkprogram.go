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

// NetworkProgramService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkProgramService] method instead.
type NetworkProgramService struct {
	Options []option.RequestOption
}

// NewNetworkProgramService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkProgramService(opts ...option.RequestOption) (r *NetworkProgramService) {
	r = &NetworkProgramService{}
	r.Options = opts
	return
}

// Get network program.
func (r *NetworkProgramService) Get(ctx context.Context, networkProgramToken string, opts ...option.RequestOption) (res *NetworkProgram, err error) {
	opts = append(r.Options[:], opts...)
	if networkProgramToken == "" {
		err = errors.New("missing required network_program_token parameter")
		return
	}
	path := fmt.Sprintf("v1/network_programs/%s", networkProgramToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List network programs.
func (r *NetworkProgramService) List(ctx context.Context, query NetworkProgramListParams, opts ...option.RequestOption) (res *pagination.SinglePage[NetworkProgram], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/network_programs"
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

// List network programs.
func (r *NetworkProgramService) ListAutoPaging(ctx context.Context, query NetworkProgramListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[NetworkProgram] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type NetworkProgram struct {
	// Lithic-generated unique identifier for the program
	Token string `json:"token,required" format:"uuid"`
	// Network product ID associated with this program.
	DefaultProductCode string `json:"default_product_code,required"`
	// The name of the network program.
	Name string `json:"name,required"`
	// RPIN value assigned by the network.
	RegisteredProgramIdentificationNumber string             `json:"registered_program_identification_number,required"`
	JSON                                  networkProgramJSON `json:"-"`
}

// networkProgramJSON contains the JSON metadata for the struct [NetworkProgram]
type networkProgramJSON struct {
	Token                                 apijson.Field
	DefaultProductCode                    apijson.Field
	Name                                  apijson.Field
	RegisteredProgramIdentificationNumber apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *NetworkProgram) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkProgramJSON) RawJSON() string {
	return r.raw
}

type NetworkProgramListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [NetworkProgramListParams]'s query parameters as
// `url.Values`.
func (r NetworkProgramListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
