// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// IPRangeService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPRangeService] method instead.
type IPRangeService struct {
	Options []option.RequestOption
}

// NewIPRangeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPRangeService(opts ...option.RequestOption) (r *IPRangeService) {
	r = &IPRangeService{}
	r.Options = opts
	return
}

// Returns the IP addresses Lithic sends responder requests from (Auth Stream
// Access, 3DS decisioning, and tokenization decisioning), so you can allowlist
// them in your firewall. Each environment returns its own set of addresses. Lithic
// announces new addresses at least 30 days before sending traffic from them. Event
// webhooks are not sent from these addresses
func (r *IPRangeService) Get(ctx context.Context, opts ...option.RequestOption) (res *IPRangeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/ip_ranges"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type IPRangeGetResponse struct {
	// IP ranges that responder requests are sent from, in CIDR notation
	Responders []string `json:"responders" api:"required"`
	// Date and time the IP ranges were last changed
	Updated time.Time              `json:"updated" api:"required" format:"date-time"`
	JSON    ipRangeGetResponseJSON `json:"-"`
}

// ipRangeGetResponseJSON contains the JSON metadata for the struct
// [IPRangeGetResponse]
type ipRangeGetResponseJSON struct {
	Responders  apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPRangeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipRangeGetResponseJSON) RawJSON() string {
	return r.raw
}
