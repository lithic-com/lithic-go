// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"github.com/lithic-com/lithic-go/internal/apijson"
)

type APIStatus struct {
	Message string        `json:"message"`
	JSON    apiStatusJSON `json:"-"`
}

// apiStatusJSON contains the JSON metadata for the struct [APIStatus]
type apiStatusJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiStatusJSON) RawJSON() string {
	return r.raw
}
