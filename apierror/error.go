package apierror

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	apijson "github.com/lithic-com/lithic-go/core/json"
)

type Error struct {
	JSON       ErrorJSON
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

type ErrorJSON struct {
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Error using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	return fmt.Sprintf("%s \"%s\": %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), string(r.JSON.Raw))
}

func (r *Error) DumpRequest(body bool) []byte {
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}
