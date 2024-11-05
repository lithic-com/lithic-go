// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"github.com/lithic-com/lithic-go/option"
)

// AuthRuleService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthRuleService] method instead.
type AuthRuleService struct {
	Options []option.RequestOption
	V2      *AuthRuleV2Service
}

// NewAuthRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthRuleService(opts ...option.RequestOption) (r *AuthRuleService) {
	r = &AuthRuleService{}
	r.Options = opts
	r.V2 = NewAuthRuleV2Service(opts...)
	return
}
