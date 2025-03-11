// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// ExternalBankAccountMicroDepositService contains methods and other services that
// help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalBankAccountMicroDepositService] method instead.
type ExternalBankAccountMicroDepositService struct {
	Options []option.RequestOption
}

// NewExternalBankAccountMicroDepositService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExternalBankAccountMicroDepositService(opts ...option.RequestOption) (r *ExternalBankAccountMicroDepositService) {
	r = &ExternalBankAccountMicroDepositService{}
	r.Options = opts
	return
}

// Verify the external bank account by providing the micro deposit amounts.
func (r *ExternalBankAccountMicroDepositService) New(ctx context.Context, externalBankAccountToken string, body ExternalBankAccountMicroDepositNewParams, opts ...option.RequestOption) (res *ExternalBankAccountMicroDepositNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if externalBankAccountToken == "" {
		err = errors.New("missing required external_bank_account_token parameter")
		return
	}
	path := fmt.Sprintf("v1/external_bank_accounts/%s/micro_deposits", externalBankAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExternalBankAccountMicroDepositNewResponse struct {
	// A globally unique identifier for this record of an external bank account
	// association. If a program links an external bank account to more than one
	// end-user or to both the program and the end-user, then Lithic will return each
	// record of the association
	Token string `json:"token,required" format:"uuid"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept USA bank accounts e.g., USA
	Country string `json:"country,required"`
	// An ISO 8601 string representing when this funding source was added to the Lithic
	// account.
	Created time.Time `json:"created,required" format:"date-time"`
	// currency of the external account 3-character alphabetic ISO 4217 code
	Currency string `json:"currency,required"`
	// The last 4 digits of the bank account. Derived by Lithic from the account number
	// passed
	LastFour string `json:"last_four,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner string `json:"owner,required"`
	// Owner Type
	OwnerType ExternalBankAccountMicroDepositNewResponseOwnerType `json:"owner_type,required"`
	// Routing Number
	RoutingNumber string `json:"routing_number,required"`
	// Account State
	State ExternalBankAccountMicroDepositNewResponseState `json:"state,required"`
	// Account Type
	Type ExternalBankAccountMicroDepositNewResponseType `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64 `json:"verification_attempts,required"`
	// Verification Method
	VerificationMethod ExternalBankAccountMicroDepositNewResponseVerificationMethod `json:"verification_method,required"`
	// Verification State
	VerificationState ExternalBankAccountMicroDepositNewResponseVerificationState `json:"verification_state,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken string `json:"account_token" format:"uuid"`
	// Address
	Address ExternalBankAccountAddress `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID string `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob time.Time `json:"dob" format:"date"`
	// Doing Business As
	DoingBusinessAs string `json:"doing_business_as"`
	// The financial account token of the operating account to fund the micro deposits
	FinancialAccountToken string `json:"financial_account_token" format:"uuid"`
	// The nickname for this External Bank Account
	Name string `json:"name"`
	// User Defined ID
	UserDefinedID string `json:"user_defined_id"`
	// Optional free text description of the reason for the failed verification. For
	// ACH micro-deposits returned, this field will display the reason return code sent
	// by the ACH network
	VerificationFailedReason string                                         `json:"verification_failed_reason"`
	JSON                     externalBankAccountMicroDepositNewResponseJSON `json:"-"`
}

// externalBankAccountMicroDepositNewResponseJSON contains the JSON metadata for
// the struct [ExternalBankAccountMicroDepositNewResponse]
type externalBankAccountMicroDepositNewResponseJSON struct {
	Token                    apijson.Field
	Country                  apijson.Field
	Created                  apijson.Field
	Currency                 apijson.Field
	LastFour                 apijson.Field
	Owner                    apijson.Field
	OwnerType                apijson.Field
	RoutingNumber            apijson.Field
	State                    apijson.Field
	Type                     apijson.Field
	VerificationAttempts     apijson.Field
	VerificationMethod       apijson.Field
	VerificationState        apijson.Field
	AccountToken             apijson.Field
	Address                  apijson.Field
	CompanyID                apijson.Field
	Dob                      apijson.Field
	DoingBusinessAs          apijson.Field
	FinancialAccountToken    apijson.Field
	Name                     apijson.Field
	UserDefinedID            apijson.Field
	VerificationFailedReason apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ExternalBankAccountMicroDepositNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountMicroDepositNewResponseJSON) RawJSON() string {
	return r.raw
}

// Owner Type
type ExternalBankAccountMicroDepositNewResponseOwnerType string

const (
	ExternalBankAccountMicroDepositNewResponseOwnerTypeBusiness   ExternalBankAccountMicroDepositNewResponseOwnerType = "BUSINESS"
	ExternalBankAccountMicroDepositNewResponseOwnerTypeIndividual ExternalBankAccountMicroDepositNewResponseOwnerType = "INDIVIDUAL"
)

func (r ExternalBankAccountMicroDepositNewResponseOwnerType) IsKnown() bool {
	switch r {
	case ExternalBankAccountMicroDepositNewResponseOwnerTypeBusiness, ExternalBankAccountMicroDepositNewResponseOwnerTypeIndividual:
		return true
	}
	return false
}

// Account State
type ExternalBankAccountMicroDepositNewResponseState string

const (
	ExternalBankAccountMicroDepositNewResponseStateEnabled ExternalBankAccountMicroDepositNewResponseState = "ENABLED"
	ExternalBankAccountMicroDepositNewResponseStateClosed  ExternalBankAccountMicroDepositNewResponseState = "CLOSED"
	ExternalBankAccountMicroDepositNewResponseStatePaused  ExternalBankAccountMicroDepositNewResponseState = "PAUSED"
)

func (r ExternalBankAccountMicroDepositNewResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountMicroDepositNewResponseStateEnabled, ExternalBankAccountMicroDepositNewResponseStateClosed, ExternalBankAccountMicroDepositNewResponseStatePaused:
		return true
	}
	return false
}

// Account Type
type ExternalBankAccountMicroDepositNewResponseType string

const (
	ExternalBankAccountMicroDepositNewResponseTypeChecking ExternalBankAccountMicroDepositNewResponseType = "CHECKING"
	ExternalBankAccountMicroDepositNewResponseTypeSavings  ExternalBankAccountMicroDepositNewResponseType = "SAVINGS"
)

func (r ExternalBankAccountMicroDepositNewResponseType) IsKnown() bool {
	switch r {
	case ExternalBankAccountMicroDepositNewResponseTypeChecking, ExternalBankAccountMicroDepositNewResponseTypeSavings:
		return true
	}
	return false
}

// Verification Method
type ExternalBankAccountMicroDepositNewResponseVerificationMethod string

const (
	ExternalBankAccountMicroDepositNewResponseVerificationMethodManual       ExternalBankAccountMicroDepositNewResponseVerificationMethod = "MANUAL"
	ExternalBankAccountMicroDepositNewResponseVerificationMethodMicroDeposit ExternalBankAccountMicroDepositNewResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountMicroDepositNewResponseVerificationMethodPlaid        ExternalBankAccountMicroDepositNewResponseVerificationMethod = "PLAID"
	ExternalBankAccountMicroDepositNewResponseVerificationMethodPrenote      ExternalBankAccountMicroDepositNewResponseVerificationMethod = "PRENOTE"
)

func (r ExternalBankAccountMicroDepositNewResponseVerificationMethod) IsKnown() bool {
	switch r {
	case ExternalBankAccountMicroDepositNewResponseVerificationMethodManual, ExternalBankAccountMicroDepositNewResponseVerificationMethodMicroDeposit, ExternalBankAccountMicroDepositNewResponseVerificationMethodPlaid, ExternalBankAccountMicroDepositNewResponseVerificationMethodPrenote:
		return true
	}
	return false
}

// Verification State
type ExternalBankAccountMicroDepositNewResponseVerificationState string

const (
	ExternalBankAccountMicroDepositNewResponseVerificationStatePending            ExternalBankAccountMicroDepositNewResponseVerificationState = "PENDING"
	ExternalBankAccountMicroDepositNewResponseVerificationStateEnabled            ExternalBankAccountMicroDepositNewResponseVerificationState = "ENABLED"
	ExternalBankAccountMicroDepositNewResponseVerificationStateFailedVerification ExternalBankAccountMicroDepositNewResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountMicroDepositNewResponseVerificationStateInsufficientFunds  ExternalBankAccountMicroDepositNewResponseVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ExternalBankAccountMicroDepositNewResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountMicroDepositNewResponseVerificationStatePending, ExternalBankAccountMicroDepositNewResponseVerificationStateEnabled, ExternalBankAccountMicroDepositNewResponseVerificationStateFailedVerification, ExternalBankAccountMicroDepositNewResponseVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type ExternalBankAccountMicroDepositNewParams struct {
	MicroDeposits param.Field[[]int64] `json:"micro_deposits,required"`
}

func (r ExternalBankAccountMicroDepositNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
