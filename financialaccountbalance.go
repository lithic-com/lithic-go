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

type FinancialAccountBalanceService struct {
	Options []option.RequestOption
}

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

// URLQuery serializes FinancialAccountBalanceListParams into a url.Values of the
// query parameters associated with this value
func (r FinancialAccountBalanceListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type FinancialAccountBalanceListResponse struct {
	Data []Balance `json:"data,required"`
	// More data exists.
	HasMore bool `json:"has_more,required"`
	JSON    FinancialAccountBalanceListResponseJSON
}

type FinancialAccountBalanceListResponseJSON struct {
	Data    apijson.Metadata
	HasMore apijson.Metadata
	raw     string
	Extras  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// FinancialAccountBalanceListResponse using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *FinancialAccountBalanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
