// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"github.com/lithic-com/lithic-go/option"
)

// CreditProductService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditProductService] method instead.
type CreditProductService struct {
	Options        []option.RequestOption
	ExtendedCredit *CreditProductExtendedCreditService
}

// NewCreditProductService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCreditProductService(opts ...option.RequestOption) (r *CreditProductService) {
	r = &CreditProductService{}
	r.Options = opts
	r.ExtendedCredit = NewCreditProductExtendedCreditService(opts...)
	return
}
