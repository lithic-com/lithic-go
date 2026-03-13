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

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// FinancialAccountInterestTierScheduleService contains methods and other services
// that help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountInterestTierScheduleService] method instead.
type FinancialAccountInterestTierScheduleService struct {
	Options []option.RequestOption
}

// NewFinancialAccountInterestTierScheduleService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewFinancialAccountInterestTierScheduleService(opts ...option.RequestOption) (r *FinancialAccountInterestTierScheduleService) {
	r = &FinancialAccountInterestTierScheduleService{}
	r.Options = opts
	return
}

// Create a new interest tier schedule entry for a supported financial account
func (r *FinancialAccountInterestTierScheduleService) New(ctx context.Context, financialAccountToken string, body FinancialAccountInterestTierScheduleNewParams, opts ...option.RequestOption) (res *InterestTierSchedule, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/interest_tier_schedule", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a specific interest tier schedule by effective date
func (r *FinancialAccountInterestTierScheduleService) Get(ctx context.Context, financialAccountToken string, effectiveDate time.Time, opts ...option.RequestOption) (res *InterestTierSchedule, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/interest_tier_schedule/%s", financialAccountToken, effectiveDate.Format("2006-01-02"))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing interest tier schedule
func (r *FinancialAccountInterestTierScheduleService) Update(ctx context.Context, financialAccountToken string, effectiveDate time.Time, body FinancialAccountInterestTierScheduleUpdateParams, opts ...option.RequestOption) (res *InterestTierSchedule, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/interest_tier_schedule/%s", financialAccountToken, effectiveDate.Format("2006-01-02"))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List interest tier schedules for a financial account with optional date
// filtering.
//
// If no date parameters are provided, returns all tier schedules. If date
// parameters are provided, uses filtering to return matching schedules (max 100).
//
// - for_date: Returns exact match (takes precedence over other dates)
// - before_date: Returns schedules with effective_date <= before_date
// - after_date: Returns schedules with effective_date >= after_date
// - Both before_date and after_date: Returns schedules in range
func (r *FinancialAccountInterestTierScheduleService) List(ctx context.Context, financialAccountToken string, query FinancialAccountInterestTierScheduleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[InterestTierSchedule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/interest_tier_schedule", financialAccountToken)
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

// List interest tier schedules for a financial account with optional date
// filtering.
//
// If no date parameters are provided, returns all tier schedules. If date
// parameters are provided, uses filtering to return matching schedules (max 100).
//
// - for_date: Returns exact match (takes precedence over other dates)
// - before_date: Returns schedules with effective_date <= before_date
// - after_date: Returns schedules with effective_date >= after_date
// - Both before_date and after_date: Returns schedules in range
func (r *FinancialAccountInterestTierScheduleService) ListAutoPaging(ctx context.Context, financialAccountToken string, query FinancialAccountInterestTierScheduleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[InterestTierSchedule] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, financialAccountToken, query, opts...))
}

// Delete an interest tier schedule entry.
//
// Returns:
//
//   - 400 Bad Request: Invalid effective_date format OR attempting to delete the
//     earliest tier schedule entry for a non-PENDING account
//   - 404 Not Found: Tier schedule entry not found for the given effective_date OR
//     ledger account not found
//
// Note: PENDING accounts can delete the earliest tier schedule entry (account
// hasn't opened yet). Active/non-PENDING accounts cannot delete the earliest entry
// to prevent orphaning the account.
//
// If the deleted tier schedule has a past effective_date and the account is
// ACTIVE, the loan tape rebuild configuration will be updated to trigger rebuilds
// from that date.
func (r *FinancialAccountInterestTierScheduleService) Delete(ctx context.Context, financialAccountToken string, effectiveDate time.Time, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/interest_tier_schedule/%s", financialAccountToken, effectiveDate.Format("2006-01-02"))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Entry in the Tier Schedule of an account
type InterestTierSchedule struct {
	// Globally unique identifier for a credit product
	CreditProductToken string `json:"credit_product_token" api:"required"`
	// Date the tier should be effective in YYYY-MM-DD format
	EffectiveDate time.Time `json:"effective_date" api:"required" format:"date"`
	// Name of a tier contained in the credit product. Mutually exclusive with
	// tier_rates
	TierName string `json:"tier_name"`
	// Custom rates per category. Mutually exclusive with tier_name
	TierRates interface{}              `json:"tier_rates"`
	JSON      interestTierScheduleJSON `json:"-"`
}

// interestTierScheduleJSON contains the JSON metadata for the struct
// [InterestTierSchedule]
type interestTierScheduleJSON struct {
	CreditProductToken apijson.Field
	EffectiveDate      apijson.Field
	TierName           apijson.Field
	TierRates          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InterestTierSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interestTierScheduleJSON) RawJSON() string {
	return r.raw
}

// Entry in the Tier Schedule of an account
type InterestTierScheduleParam struct {
	// Globally unique identifier for a credit product
	CreditProductToken param.Field[string] `json:"credit_product_token" api:"required"`
	// Date the tier should be effective in YYYY-MM-DD format
	EffectiveDate param.Field[time.Time] `json:"effective_date" api:"required" format:"date"`
	// Name of a tier contained in the credit product. Mutually exclusive with
	// tier_rates
	TierName param.Field[string] `json:"tier_name"`
	// Custom rates per category. Mutually exclusive with tier_name
	TierRates param.Field[interface{}] `json:"tier_rates"`
}

func (r InterestTierScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FinancialAccountInterestTierScheduleNewParams struct {
	// Entry in the Tier Schedule of an account
	InterestTierSchedule InterestTierScheduleParam `json:"interest_tier_schedule" api:"required"`
}

func (r FinancialAccountInterestTierScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InterestTierSchedule)
}

type FinancialAccountInterestTierScheduleUpdateParams struct {
	// Name of a tier contained in the credit product. Mutually exclusive with
	// tier_rates
	TierName param.Field[string] `json:"tier_name"`
	// Custom rates per category. Mutually exclusive with tier_name
	TierRates param.Field[interface{}] `json:"tier_rates"`
}

func (r FinancialAccountInterestTierScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FinancialAccountInterestTierScheduleListParams struct {
	// Return schedules with effective_date >= after_date (ISO format YYYY-MM-DD)
	AfterDate param.Field[time.Time] `query:"after_date" format:"date"`
	// Return schedules with effective_date <= before_date (ISO format YYYY-MM-DD)
	BeforeDate param.Field[time.Time] `query:"before_date" format:"date"`
	// Return schedule with effective_date == for_date (ISO format YYYY-MM-DD)
	ForDate param.Field[time.Time] `query:"for_date" format:"date"`
}

// URLQuery serializes [FinancialAccountInterestTierScheduleListParams]'s query
// parameters as `url.Values`.
func (r FinancialAccountInterestTierScheduleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
