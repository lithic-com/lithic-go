// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// CardBalanceService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCardBalanceService] method
// instead.
type CardBalanceService struct {
	Options []option.RequestOption
}

// NewCardBalanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardBalanceService(opts ...option.RequestOption) (r *CardBalanceService) {
	r = &CardBalanceService{}
	r.Options = opts
	return
}

// Get the balances for a given card.
func (r *CardBalanceService) List(ctx context.Context, cardToken string, query CardBalanceListParams, opts ...option.RequestOption) (res *shared.SinglePage[Balance], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("cards/%s/balances", cardToken)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get the balances for a given card.
func (r *CardBalanceService) ListAutoPaging(ctx context.Context, cardToken string, query CardBalanceListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[Balance] {
	return shared.NewSinglePageAutoPager(r.List(ctx, cardToken, query, opts...))
}

type CardBalanceListParams struct {
	// UTC date of the balance to retrieve. Defaults to latest available balance
	BalanceDate param.Field[time.Time] `query:"balance_date" format:"date-time"`
	// Balance after a given financial event occured. For example, passing the
	// event_token of a $5 CARD_CLEARING financial event will return a balance
	// decreased by $5
	LastTransactionEventToken param.Field[string] `query:"last_transaction_event_token" format:"uuid"`
}

// URLQuery serializes [CardBalanceListParams]'s query parameters as `url.Values`.
func (r CardBalanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
