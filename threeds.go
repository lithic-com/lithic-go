// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"github.com/lithic-com/lithic-go/option"
)

// ThreeDSService contains methods and other services that help with interacting
// with the lithic API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewThreeDSService] method instead.
type ThreeDSService struct {
	Options        []option.RequestOption
	Authentication *ThreeDSAuthenticationService
	Descisioning   *ThreeDSDescisioningService
}

// NewThreeDSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewThreeDSService(opts ...option.RequestOption) (r *ThreeDSService) {
	r = &ThreeDSService{}
	r.Options = opts
	r.Authentication = NewThreeDSAuthenticationService(opts...)
	r.Descisioning = NewThreeDSDescisioningService(opts...)
	return
}
