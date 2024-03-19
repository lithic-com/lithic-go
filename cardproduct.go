// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// CardProductService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCardProductService] method
// instead.
type CardProductService struct {
	Options []option.RequestOption
}

// NewCardProductService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardProductService(opts ...option.RequestOption) (r *CardProductService) {
	r = &CardProductService{}
	r.Options = opts
	return
}

// Get the Credit Detail for the card product
func (r *CardProductService) CreditDetail(ctx context.Context, opts ...option.RequestOption) (res *CardProductCreditDetailResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "card_product/credit_detail"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CardProductCreditDetailResponse struct {
	// The amount of credit extended within the program
	CreditExtended int64 `json:"credit_extended,required"`
	// The total credit limit of the program
	CreditLimit int64                               `json:"credit_limit,required"`
	JSON        cardProductCreditDetailResponseJSON `json:"-"`
}

// cardProductCreditDetailResponseJSON contains the JSON metadata for the struct
// [CardProductCreditDetailResponse]
type cardProductCreditDetailResponseJSON struct {
	CreditExtended apijson.Field
	CreditLimit    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardProductCreditDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardProductCreditDetailResponseJSON) RawJSON() string {
	return r.raw
}
