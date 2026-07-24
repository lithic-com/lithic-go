// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// FinancialAccountBalanceService contains methods and other services that help
// with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountBalanceService] method instead.
type FinancialAccountBalanceService struct {
	Options []option.RequestOption
}

// NewFinancialAccountBalanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFinancialAccountBalanceService(opts ...option.RequestOption) (r *FinancialAccountBalanceService) {
	r = &FinancialAccountBalanceService{}
	r.Options = opts
	return
}

// Get the balances for a given financial account.
func (r *FinancialAccountBalanceService) List(ctx context.Context, financialAccountToken string, query FinancialAccountBalanceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[FinancialAccountBalance], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/balances", financialAccountToken)
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

// Get the balances for a given financial account.
func (r *FinancialAccountBalanceService) ListAutoPaging(ctx context.Context, financialAccountToken string, query FinancialAccountBalanceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FinancialAccountBalance] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, financialAccountToken, query, opts...))
}

type FinancialAccountBalanceListParams struct {
	// UTC date of the balance to retrieve. Defaults to latest available balance
	BalanceDate param.Field[time.Time] `query:"balance_date" format:"date-time"`
	// Balance after a given financial event occured. Note: if an account receives
	// multiple events around the same time whose financial impacts cancel out, a
	// balance lookup by one of those event tokens may return 404, since their combined
	// impact on the account is zero.
	LastTransactionEventToken param.Field[string] `query:"last_transaction_event_token" format:"uuid"`
}

// URLQuery serializes [FinancialAccountBalanceListParams]'s query parameters as
// `url.Values`.
func (r FinancialAccountBalanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
