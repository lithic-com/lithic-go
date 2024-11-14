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
)

// CreditProductPrimeRateService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditProductPrimeRateService] method instead.
type CreditProductPrimeRateService struct {
	Options []option.RequestOption
}

// NewCreditProductPrimeRateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCreditProductPrimeRateService(opts ...option.RequestOption) (r *CreditProductPrimeRateService) {
	r = &CreditProductPrimeRateService{}
	r.Options = opts
	return
}

// Post Credit Product Prime Rate
func (r *CreditProductPrimeRateService) New(ctx context.Context, creditProductToken string, body CreditProductPrimeRateNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if creditProductToken == "" {
		err = errors.New("missing required credit_product_token parameter")
		return
	}
	path := fmt.Sprintf("v1/credit_products/%s/prime_rates", creditProductToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get Credit Product Prime Rates
func (r *CreditProductPrimeRateService) Get(ctx context.Context, creditProductToken string, query CreditProductPrimeRateGetParams, opts ...option.RequestOption) (res *CreditProductPrimeRateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if creditProductToken == "" {
		err = errors.New("missing required credit_product_token parameter")
		return
	}
	path := fmt.Sprintf("v1/credit_products/%s/prime_rates", creditProductToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CreditProductPrimeRateGetResponse struct {
	// List of prime rates
	Data []CreditProductPrimeRateGetResponseData `json:"data,required"`
	// Whether there are more prime rates
	HasMore bool                                  `json:"has_more,required"`
	JSON    creditProductPrimeRateGetResponseJSON `json:"-"`
}

// creditProductPrimeRateGetResponseJSON contains the JSON metadata for the struct
// [CreditProductPrimeRateGetResponse]
type creditProductPrimeRateGetResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditProductPrimeRateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditProductPrimeRateGetResponseJSON) RawJSON() string {
	return r.raw
}

type CreditProductPrimeRateGetResponseData struct {
	// Date the rate goes into effect
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// The rate in decimal format
	Rate string                                    `json:"rate,required"`
	JSON creditProductPrimeRateGetResponseDataJSON `json:"-"`
}

// creditProductPrimeRateGetResponseDataJSON contains the JSON metadata for the
// struct [CreditProductPrimeRateGetResponseData]
type creditProductPrimeRateGetResponseDataJSON struct {
	EffectiveDate apijson.Field
	Rate          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CreditProductPrimeRateGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditProductPrimeRateGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type CreditProductPrimeRateNewParams struct {
	// Date the rate goes into effect
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date"`
	// The rate in decimal format
	Rate param.Field[string] `json:"rate,required"`
}

func (r CreditProductPrimeRateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditProductPrimeRateGetParams struct {
	// The effective date that the prime rates ends before
	EndingBefore param.Field[time.Time] `query:"ending_before" format:"date"`
	// The effective date that the prime rate starts after
	StartingAfter param.Field[time.Time] `query:"starting_after" format:"date"`
}

// URLQuery serializes [CreditProductPrimeRateGetParams]'s query parameters as
// `url.Values`.
func (r CreditProductPrimeRateGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
