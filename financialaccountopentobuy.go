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

// FinancialAccountOpenToBuyService contains methods and other services that help
// with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountOpenToBuyService] method instead.
type FinancialAccountOpenToBuyService struct {
	Options []option.RequestOption
}

// NewFinancialAccountOpenToBuyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFinancialAccountOpenToBuyService(opts ...option.RequestOption) (r *FinancialAccountOpenToBuyService) {
	r = &FinancialAccountOpenToBuyService{}
	r.Options = opts
	return
}

// Get the funds available for card spend backed by a given Security Account, along
// with the balances that amount is derived from.
//
// Open to buy is the amount Lithic authorizes card spend against. It is not a
// stored balance, so it is recalculated on every request from the Security
// Account, the funds held against spend Lithic has already paid out to the
// networks on your behalf, and the spend that has not yet been collected. The
// accounts that feed the calculation depend on your program setup, so
// `summary.settled_funds` is `null` outside Commercial Charge.
//
// Supported for Commercial Charge, Dynamic Reserve, and Secured Charge programs.
// Returns `404` if `financial_account_token` is not a Security Account you own, or
// if your program setup does not use an open to buy calculation.
func (r *FinancialAccountOpenToBuyService) Get(ctx context.Context, financialAccountToken string, opts ...option.RequestOption) (res *OpenToBuy, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/open_to_buy", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Open to Buy
type OpenToBuy struct {
	// Funds available for card spend backed by this Security Account, in the
	// currency's smallest unit (e.g., cents for USD). Equal to the sum of the amounts
	// in `summary`, and reaches zero once outstanding spend has consumed all available
	// funding
	OpenToBuy int64 `json:"open_to_buy" api:"required"`
	// Balances that open to buy is derived from
	Summary OpenToBuySummary `json:"summary" api:"required"`
	JSON    openToBuyJSON    `json:"-"`
}

// openToBuyJSON contains the JSON metadata for the struct [OpenToBuy]
type openToBuyJSON struct {
	OpenToBuy   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OpenToBuy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r openToBuyJSON) RawJSON() string {
	return r.raw
}

// Open to Buy Summary
type OpenToBuySummary struct {
	// Available balance of the Security Account backing card spend, in the currency's
	// smallest unit (e.g., cents for USD)
	Security int64 `json:"security" api:"required"`
	// Funding that has moved out of the Security Account to cover card spend Lithic
	// has already paid out to the networks, in the currency's smallest unit (e.g.,
	// cents for USD). Open to buy counts it alongside `security`, and it clears once
	// collected from your business clients. Only Commercial Charge tracks this
	// separately, so this is `null` for every other program setup
	SettledFunds int64 `json:"settled_funds" api:"required,nullable"`
	// Customer card spend that has not yet been collected, in the currency's smallest
	// unit (e.g., cents for USD). Reported as a negative amount, because it reduces
	// open to buy
	TotalOutstandingSpend int64                `json:"total_outstanding_spend" api:"required"`
	JSON                  openToBuySummaryJSON `json:"-"`
}

// openToBuySummaryJSON contains the JSON metadata for the struct
// [OpenToBuySummary]
type openToBuySummaryJSON struct {
	Security              apijson.Field
	SettledFunds          apijson.Field
	TotalOutstandingSpend apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *OpenToBuySummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r openToBuySummaryJSON) RawJSON() string {
	return r.raw
}
