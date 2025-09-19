// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// TransactionEnhancedCommercialDataService contains methods and other services
// that help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionEnhancedCommercialDataService] method instead.
type TransactionEnhancedCommercialDataService struct {
	Options []option.RequestOption
}

// NewTransactionEnhancedCommercialDataService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransactionEnhancedCommercialDataService(opts ...option.RequestOption) (r *TransactionEnhancedCommercialDataService) {
	r = &TransactionEnhancedCommercialDataService{}
	r.Options = opts
	return
}

// Get all L2/L3 enhanced commercial data associated with a transaction. Not
// available in sandbox.
func (r *TransactionEnhancedCommercialDataService) Get(ctx context.Context, transactionToken string, opts ...option.RequestOption) (res *TransactionEnhancedCommercialDataGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionToken == "" {
		err = errors.New("missing required transaction_token parameter")
		return
	}
	path := fmt.Sprintf("v1/transactions/%s/enhanced_commercial_data", transactionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TransactionEnhancedCommercialDataGetResponse struct {
	Data []EnhancedData                                   `json:"data,required"`
	JSON transactionEnhancedCommercialDataGetResponseJSON `json:"-"`
}

// transactionEnhancedCommercialDataGetResponseJSON contains the JSON metadata for
// the struct [TransactionEnhancedCommercialDataGetResponse]
type transactionEnhancedCommercialDataGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionEnhancedCommercialDataGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionEnhancedCommercialDataGetResponseJSON) RawJSON() string {
	return r.raw
}
