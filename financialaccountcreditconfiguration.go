// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	AccountToken string `json:"account_token,required" format:"uuid"`
	// Reason for the financial account being marked as Charged Off
	ChargedOffReason FinancialAccountCreditConfigChargedOffReason `json:"charged_off_reason,required,nullable"`
	CreditLimit      int64                                        `json:"credit_limit,required,nullable"`
	// Globally unique identifier for the credit product
	CreditProductToken       string `json:"credit_product_token,required,nullable"`
	ExternalBankAccountToken string `json:"external_bank_account_token,required,nullable" format:"uuid"`
	// State of the financial account
	FinancialAccountState FinancialAccountCreditConfigFinancialAccountState `json:"financial_account_state,required"`
	IsSpendBlocked        bool                                              `json:"is_spend_blocked,required"`
	// Tier assigned to the financial account
	Tier string                           `json:"tier,required,nullable"`
	JSON financialAccountCreditConfigJSON `json:"-"`
}

// financialAccountCreditConfigJSON contains the JSON metadata for the struct
// [FinancialAccountCreditConfig]
type financialAccountCreditConfigJSON struct {
	AccountToken             apijson.Field
	ChargedOffReason         apijson.Field
	CreditLimit              apijson.Field
	CreditProductToken       apijson.Field
	ExternalBankAccountToken apijson.Field
	FinancialAccountState    apijson.Field
	IsSpendBlocked           apijson.Field
	Tier                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *FinancialAccountCreditConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAccountCreditConfigJSON) RawJSON() string {
	return r.raw
}

// Reason for the financial account being marked as Charged Off
type FinancialAccountCreditConfigChargedOffReason string

const (
	FinancialAccountCreditConfigChargedOffReasonDelinquent FinancialAccountCreditConfigChargedOffReason = "DELINQUENT"
	FinancialAccountCreditConfigChargedOffReasonFraud      FinancialAccountCreditConfigChargedOffReason = "FRAUD"
)

func (r FinancialAccountCreditConfigChargedOffReason) IsKnown() bool {
	switch r {
	case FinancialAccountCreditConfigChargedOffReasonDelinquent, FinancialAccountCreditConfigChargedOffReasonFraud:
		return true
	}
	return false
}

// State of the financial account
type FinancialAccountCreditConfigFinancialAccountState string

const (
	FinancialAccountCreditConfigFinancialAccountStatePending    FinancialAccountCreditConfigFinancialAccountState = "PENDING"
	FinancialAccountCreditConfigFinancialAccountStateCurrent    FinancialAccountCreditConfigFinancialAccountState = "CURRENT"
	FinancialAccountCreditConfigFinancialAccountStateDelinquent FinancialAccountCreditConfigFinancialAccountState = "DELINQUENT"
	FinancialAccountCreditConfigFinancialAccountStateChargedOff FinancialAccountCreditConfigFinancialAccountState = "CHARGED_OFF"
)

func (r FinancialAccountCreditConfigFinancialAccountState) IsKnown() bool {
	switch r {
	case FinancialAccountCreditConfigFinancialAccountStatePending, FinancialAccountCreditConfigFinancialAccountStateCurrent, FinancialAccountCreditConfigFinancialAccountStateDelinquent, FinancialAccountCreditConfigFinancialAccountStateChargedOff:
		return true
	}
	return false
}

type FinancialAccountCreditConfigurationUpdateParams struct {
	CreditLimit param.Field[int64] `json:"credit_limit"`
	// Globally unique identifier for the credit product
	CreditProductToken       param.Field[string] `json:"credit_product_token"`
	ExternalBankAccountToken param.Field[string] `json:"external_bank_account_token" format:"uuid"`
	// Tier to assign to a financial account
	Tier param.Field[string] `json:"tier"`
}

func (r FinancialAccountCreditConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
