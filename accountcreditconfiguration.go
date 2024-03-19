// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// AccountCreditConfigurationService contains methods and other services that help
// with interacting with the lithic API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountCreditConfigurationService] method instead.
type AccountCreditConfigurationService struct {
	Options []option.RequestOption
}

// NewAccountCreditConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCreditConfigurationService(opts ...option.RequestOption) (r *AccountCreditConfigurationService) {
	r = &AccountCreditConfigurationService{}
	r.Options = opts
	return
}

// Get an Account's credit configuration
func (r *AccountCreditConfigurationService) Get(ctx context.Context, accountToken string, opts ...option.RequestOption) (res *BusinessAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/credit_configuration", accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Business Accounts credit configuration
func (r *AccountCreditConfigurationService) Update(ctx context.Context, accountToken string, body AccountCreditConfigurationUpdateParams, opts ...option.RequestOption) (res *BusinessAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/credit_configuration", accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountCreditConfigurationUpdateParams struct {
	// Number of days within the billing period
	BillingPeriod param.Field[int64] `json:"billing_period"`
	// Credit limit extended to the Business Account
	CreditLimit param.Field[int64] `json:"credit_limit"`
	// The external bank account token to use for auto-collections
	ExternalBankAccountToken param.Field[string] `json:"external_bank_account_token" format:"uuid"`
	// Number of days after the billing period ends that a payment is required
	PaymentPeriod param.Field[int64] `json:"payment_period"`
}

func (r AccountCreditConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
