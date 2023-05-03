package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/field"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// FinancialAccountBalanceService contains methods and other services that help
// with interacting with the lithic API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewFinancialAccountBalanceService] method instead.
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
func (r *FinancialAccountBalanceService) List(ctx context.Context, financial_account_token string, query FinancialAccountBalanceListParams, opts ...option.RequestOption) (res *shared.SinglePage[Balance], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("financial_accounts/%s/balances", financial_account_token)
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
func (r *FinancialAccountBalanceService) ListAutoPaging(ctx context.Context, financial_account_token string, query FinancialAccountBalanceListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[Balance] {
	return shared.NewSinglePageAutoPager(r.List(ctx, financial_account_token, query, opts...))
}

type FinancialAccountBalanceListParams struct {
	// UTC date of the balance to retrieve. Defaults to latest available balance
	BalanceDate field.Field[time.Time] `query:"balance_date" format:"date-time"`
	// Balance after a given financial event occured. For example, passing the
	// event_token of a $5 CARD_CLEARING financial event will return a balance
	// decreased by $5
	LastTransactionEventToken field.Field[string] `query:"last_transaction_event_token" format:"uuid"`
}

// URLQuery serializes [FinancialAccountBalanceListParams]'s query parameters as
// `url.Values`.
func (r FinancialAccountBalanceListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type FinancialAccountBalanceListResponse struct {
	Data []Balance `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    financialAccountBalanceListResponseJSON
}

// financialAccountBalanceListResponseJSON contains the JSON metadata for the
// struct [FinancialAccountBalanceListResponse]
type financialAccountBalanceListResponseJSON struct {
	Data    apijson.Field
	HasMore apijson.Field
	raw     string
	Extras  map[string]apijson.Field
}

func (r *FinancialAccountBalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
