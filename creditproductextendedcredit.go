// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// CreditProductExtendedCreditService contains methods and other services that help
// with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditProductExtendedCreditService] method instead.
type CreditProductExtendedCreditService struct {
	Options []option.RequestOption
}

// NewCreditProductExtendedCreditService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCreditProductExtendedCreditService(opts ...option.RequestOption) (r *CreditProductExtendedCreditService) {
	r = &CreditProductExtendedCreditService{}
	r.Options = opts
	return
}

// Get the extended credit for a given credit product under a program
func (r *CreditProductExtendedCreditService) Get(ctx context.Context, creditProductToken string, opts ...option.RequestOption) (res *ExtendedCredit, err error) {
	opts = append(r.Options[:], opts...)
	if creditProductToken == "" {
		err = errors.New("missing required credit_product_token parameter")
		return
	}
	path := fmt.Sprintf("v1/credit_products/%s/extended_credit", creditProductToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ExtendedCredit struct {
	CreditExtended int64              `json:"credit_extended,required"`
	JSON           extendedCreditJSON `json:"-"`
}

// extendedCreditJSON contains the JSON metadata for the struct [ExtendedCredit]
type extendedCreditJSON struct {
	CreditExtended apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ExtendedCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extendedCreditJSON) RawJSON() string {
	return r.raw
}
