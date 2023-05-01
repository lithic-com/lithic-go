package lithic

import (
	"context"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

type LithicService struct {
	Options []option.RequestOption
}

func NewLithicService(opts ...option.RequestOption) (r *LithicService) {
	r = &LithicService{}
	r.Options = opts
	return
}

// API status check
func (r *LithicService) APIStatus(ctx context.Context, opts ...option.RequestOption) (res *APIStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := "status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIStatus struct {
	Message string `json:"message"`
	JSON    APIStatusJSON
}

type APIStatusJSON struct {
	Message apijson.Metadata
	raw     string
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into APIStatus using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *APIStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
