// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// FinancialAccountCreditConfigurationService contains methods and other services
// that help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountCreditConfigurationService] method instead.
type FinancialAccountCreditConfigurationService struct {
	Options []option.RequestOption
}

// NewFinancialAccountCreditConfigurationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewFinancialAccountCreditConfigurationService(opts ...option.RequestOption) (r *FinancialAccountCreditConfigurationService) {
	r = &FinancialAccountCreditConfigurationService{}
	r.Options = opts
	return
}

// Get an Account's credit configuration
func (r *FinancialAccountCreditConfigurationService) Get(ctx context.Context, financialAccountToken string, opts ...option.RequestOption) (res *FinancialAccountCreditConfig, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/credit_configuration", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an account's credit configuration
func (r *FinancialAccountCreditConfigurationService) Update(ctx context.Context, financialAccountToken string, body FinancialAccountCreditConfigurationUpdateParams, opts ...option.RequestOption) (res *FinancialAccountCreditConfig, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/credit_configuration", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type FinancialAccountCreditConfig struct {
	// Globally unique identifier for the account
	AccountToken                string                                                  `json:"account_token,required" format:"uuid"`
	AutoCollectionConfiguration FinancialAccountCreditConfigAutoCollectionConfiguration `json:"auto_collection_configuration,required"`
	CreditLimit                 int64                                                   `json:"credit_limit,required,nullable"`
	// Globally unique identifier for the credit product
	CreditProductToken       string `json:"credit_product_token,required,nullable"`
	ExternalBankAccountToken string `json:"external_bank_account_token,required,nullable" format:"uuid"`
	// Tier assigned to the financial account
	Tier string                           `json:"tier,required,nullable"`
	JSON financialAccountCreditConfigJSON `json:"-"`
}

// financialAccountCreditConfigJSON contains the JSON metadata for the struct
// [FinancialAccountCreditConfig]
type financialAccountCreditConfigJSON struct {
	AccountToken                apijson.Field
	AutoCollectionConfiguration apijson.Field
	CreditLimit                 apijson.Field
	CreditProductToken          apijson.Field
	ExternalBankAccountToken    apijson.Field
	Tier                        apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *FinancialAccountCreditConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountCreditConfigJSON) RawJSON() string {
	return r.raw
}

type FinancialAccountCreditConfigAutoCollectionConfiguration struct {
	// If auto collection is enabled for this account
	AutoCollectionEnabled bool                                                        `json:"auto_collection_enabled,required"`
	JSON                  financialAccountCreditConfigAutoCollectionConfigurationJSON `json:"-"`
}

// financialAccountCreditConfigAutoCollectionConfigurationJSON contains the JSON
// metadata for the struct
// [FinancialAccountCreditConfigAutoCollectionConfiguration]
type financialAccountCreditConfigAutoCollectionConfigurationJSON struct {
	AutoCollectionEnabled apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *FinancialAccountCreditConfigAutoCollectionConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountCreditConfigAutoCollectionConfigurationJSON) RawJSON() string {
	return r.raw
}

type FinancialAccountCreditConfigurationUpdateParams struct {
	AutoCollectionConfiguration param.Field[FinancialAccountCreditConfigurationUpdateParamsAutoCollectionConfiguration] `json:"auto_collection_configuration"`
	CreditLimit                 param.Field[int64]                                                                      `json:"credit_limit"`
	// Globally unique identifier for the credit product
	CreditProductToken       param.Field[string] `json:"credit_product_token"`
	ExternalBankAccountToken param.Field[string] `json:"external_bank_account_token" format:"uuid"`
	// Tier to assign to a financial account
	Tier param.Field[string] `json:"tier"`
}

func (r FinancialAccountCreditConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FinancialAccountCreditConfigurationUpdateParamsAutoCollectionConfiguration struct {
	// If auto collection is enabled for this account
	AutoCollectionEnabled param.Field[bool] `json:"auto_collection_enabled"`
}

func (r FinancialAccountCreditConfigurationUpdateParamsAutoCollectionConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
