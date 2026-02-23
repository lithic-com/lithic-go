// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// FinancialAccountLoanTapeConfigurationService contains methods and other services
// that help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinancialAccountLoanTapeConfigurationService] method instead.
type FinancialAccountLoanTapeConfigurationService struct {
	Options []option.RequestOption
}

// NewFinancialAccountLoanTapeConfigurationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewFinancialAccountLoanTapeConfigurationService(opts ...option.RequestOption) (r *FinancialAccountLoanTapeConfigurationService) {
	r = &FinancialAccountLoanTapeConfigurationService{}
	r.Options = opts
	return
}

// Get the loan tape configuration for a given financial account.
func (r *FinancialAccountLoanTapeConfigurationService) Get(ctx context.Context, financialAccountToken string, opts ...option.RequestOption) (res *LoanTapeConfiguration, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/loan_tape_configuration", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Configuration for loan tapes
type LoanTapeConfiguration struct {
	CreatedAt             time.Time `json:"created_at,required" format:"date-time"`
	FinancialAccountToken string    `json:"financial_account_token,required" format:"uuid"`
	InstanceToken         string    `json:"instance_token,required" format:"uuid"`
	UpdatedAt             time.Time `json:"updated_at,required" format:"date-time"`
	CreditProductToken    string    `json:"credit_product_token"`
	// Configuration for building loan tapes
	LoanTapeRebuildConfiguration LoanTapeRebuildConfiguration `json:"loan_tape_rebuild_configuration"`
	TierScheduleChangedAt        time.Time                    `json:"tier_schedule_changed_at" format:"date-time"`
	JSON                         loanTapeConfigurationJSON    `json:"-"`
}

// loanTapeConfigurationJSON contains the JSON metadata for the struct
// [LoanTapeConfiguration]
type loanTapeConfigurationJSON struct {
	CreatedAt                    apijson.Field
	FinancialAccountToken        apijson.Field
	InstanceToken                apijson.Field
	UpdatedAt                    apijson.Field
	CreditProductToken           apijson.Field
	LoanTapeRebuildConfiguration apijson.Field
	TierScheduleChangedAt        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *LoanTapeConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeConfigurationJSON) RawJSON() string {
	return r.raw
}

// Configuration for building loan tapes
type LoanTapeRebuildConfiguration struct {
	// Whether the account's loan tapes need to be rebuilt or not
	RebuildNeeded bool `json:"rebuild_needed,required"`
	// The date for which the account's loan tapes were last rebuilt
	LastRebuild time.Time `json:"last_rebuild" format:"date"`
	// Date from which to start rebuilding from if the account requires a rebuild
	RebuildFrom time.Time                        `json:"rebuild_from" format:"date"`
	JSON        loanTapeRebuildConfigurationJSON `json:"-"`
}

// loanTapeRebuildConfigurationJSON contains the JSON metadata for the struct
// [LoanTapeRebuildConfiguration]
type loanTapeRebuildConfigurationJSON struct {
	RebuildNeeded apijson.Field
	LastRebuild   apijson.Field
	RebuildFrom   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoanTapeRebuildConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanTapeRebuildConfigurationJSON) RawJSON() string {
	return r.raw
}
