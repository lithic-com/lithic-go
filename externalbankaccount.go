// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/pagination"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// ExternalBankAccountService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalBankAccountService] method instead.
type ExternalBankAccountService struct {
	Options       []option.RequestOption
	MicroDeposits *ExternalBankAccountMicroDepositService
}

// NewExternalBankAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExternalBankAccountService(opts ...option.RequestOption) (r *ExternalBankAccountService) {
	r = &ExternalBankAccountService{}
	r.Options = opts
	r.MicroDeposits = NewExternalBankAccountMicroDepositService(opts...)
	return
}

// Creates an external bank account within a program or Lithic account.
func (r *ExternalBankAccountService) New(ctx context.Context, body ExternalBankAccountNewParams, opts ...option.RequestOption) (res *ExternalBankAccountNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "external_bank_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the external bank account by token.
func (r *ExternalBankAccountService) Get(ctx context.Context, externalBankAccountToken string, opts ...option.RequestOption) (res *ExternalBankAccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if externalBankAccountToken == "" {
		err = errors.New("missing required external_bank_account_token parameter")
		return
	}
	path := fmt.Sprintf("external_bank_accounts/%s", externalBankAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the external bank account by token.
func (r *ExternalBankAccountService) Update(ctx context.Context, externalBankAccountToken string, body ExternalBankAccountUpdateParams, opts ...option.RequestOption) (res *ExternalBankAccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if externalBankAccountToken == "" {
		err = errors.New("missing required external_bank_account_token parameter")
		return
	}
	path := fmt.Sprintf("external_bank_accounts/%s", externalBankAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all the external bank accounts for the provided search criteria.
func (r *ExternalBankAccountService) List(ctx context.Context, query ExternalBankAccountListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ExternalBankAccountListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "external_bank_accounts"
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

// List all the external bank accounts for the provided search criteria.
func (r *ExternalBankAccountService) ListAutoPaging(ctx context.Context, query ExternalBankAccountListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ExternalBankAccountListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Retry external bank account micro deposit verification.
func (r *ExternalBankAccountService) RetryMicroDeposits(ctx context.Context, externalBankAccountToken string, body ExternalBankAccountRetryMicroDepositsParams, opts ...option.RequestOption) (res *ExternalBankAccountRetryMicroDepositsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if externalBankAccountToken == "" {
		err = errors.New("missing required external_bank_account_token parameter")
		return
	}
	path := fmt.Sprintf("external_bank_accounts/%s/retry_micro_deposits", externalBankAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retry external bank account prenote verification.
func (r *ExternalBankAccountService) RetryPrenote(ctx context.Context, externalBankAccountToken string, body ExternalBankAccountRetryPrenoteParams, opts ...option.RequestOption) (res *ExternalBankAccountRetryPrenoteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if externalBankAccountToken == "" {
		err = errors.New("missing required external_bank_account_token parameter")
		return
	}
	path := fmt.Sprintf("external_bank_accounts/%s/retry_prenote", externalBankAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExternalBankAccountAddress struct {
	Address1   string                         `json:"address1,required"`
	City       string                         `json:"city,required"`
	Country    string                         `json:"country,required"`
	PostalCode string                         `json:"postal_code,required"`
	State      string                         `json:"state,required"`
	Address2   string                         `json:"address2"`
	JSON       externalBankAccountAddressJSON `json:"-"`
}

// externalBankAccountAddressJSON contains the JSON metadata for the struct
// [ExternalBankAccountAddress]
type externalBankAccountAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExternalBankAccountAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountAddressJSON) RawJSON() string {
	return r.raw
}

type ExternalBankAccountAddressParam struct {
	Address1   param.Field[string] `json:"address1,required"`
	City       param.Field[string] `json:"city,required"`
	Country    param.Field[string] `json:"country,required"`
	PostalCode param.Field[string] `json:"postal_code,required"`
	State      param.Field[string] `json:"state,required"`
	Address2   param.Field[string] `json:"address2"`
}

func (r ExternalBankAccountAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnerType string

const (
	OwnerTypeIndividual OwnerType = "INDIVIDUAL"
	OwnerTypeBusiness   OwnerType = "BUSINESS"
)

func (r OwnerType) IsKnown() bool {
	switch r {
	case OwnerTypeIndividual, OwnerTypeBusiness:
		return true
	}
	return false
}

type VerificationMethod string

const (
	VerificationMethodManual             VerificationMethod = "MANUAL"
	VerificationMethodMicroDeposit       VerificationMethod = "MICRO_DEPOSIT"
	VerificationMethodPlaid              VerificationMethod = "PLAID"
	VerificationMethodPrenote            VerificationMethod = "PRENOTE"
	VerificationMethodExternallyVerified VerificationMethod = "EXTERNALLY_VERIFIED"
)

func (r VerificationMethod) IsKnown() bool {
	switch r {
	case VerificationMethodManual, VerificationMethodMicroDeposit, VerificationMethodPlaid, VerificationMethodPrenote, VerificationMethodExternallyVerified:
		return true
	}
	return false
}

type ExternalBankAccountNewResponse struct {
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
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency string `json:"currency,required"`
	// The last 4 digits of the bank account. Derived by Lithic from the account number
	// passed
	LastFour string `json:"last_four,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner string `json:"owner,required"`
	// Owner Type
	OwnerType ExternalBankAccountNewResponseOwnerType `json:"owner_type,required"`
	// Routing Number
	RoutingNumber string `json:"routing_number,required"`
	// Account State
	State ExternalBankAccountNewResponseState `json:"state,required"`
	// Account Type
	Type ExternalBankAccountNewResponseType `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64 `json:"verification_attempts,required"`
	// Verification Method
	VerificationMethod ExternalBankAccountNewResponseVerificationMethod `json:"verification_method,required"`
	// Verification State
	VerificationState ExternalBankAccountNewResponseVerificationState `json:"verification_state,required"`
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
	VerificationFailedReason string                             `json:"verification_failed_reason"`
	JSON                     externalBankAccountNewResponseJSON `json:"-"`
}

// externalBankAccountNewResponseJSON contains the JSON metadata for the struct
// [ExternalBankAccountNewResponse]
type externalBankAccountNewResponseJSON struct {
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

func (r *ExternalBankAccountNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountNewResponseJSON) RawJSON() string {
	return r.raw
}

// Owner Type
type ExternalBankAccountNewResponseOwnerType string

const (
	ExternalBankAccountNewResponseOwnerTypeBusiness   ExternalBankAccountNewResponseOwnerType = "BUSINESS"
	ExternalBankAccountNewResponseOwnerTypeIndividual ExternalBankAccountNewResponseOwnerType = "INDIVIDUAL"
)

func (r ExternalBankAccountNewResponseOwnerType) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewResponseOwnerTypeBusiness, ExternalBankAccountNewResponseOwnerTypeIndividual:
		return true
	}
	return false
}

// Account State
type ExternalBankAccountNewResponseState string

const (
	ExternalBankAccountNewResponseStateEnabled ExternalBankAccountNewResponseState = "ENABLED"
	ExternalBankAccountNewResponseStateClosed  ExternalBankAccountNewResponseState = "CLOSED"
	ExternalBankAccountNewResponseStatePaused  ExternalBankAccountNewResponseState = "PAUSED"
)

func (r ExternalBankAccountNewResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewResponseStateEnabled, ExternalBankAccountNewResponseStateClosed, ExternalBankAccountNewResponseStatePaused:
		return true
	}
	return false
}

// Account Type
type ExternalBankAccountNewResponseType string

const (
	ExternalBankAccountNewResponseTypeChecking ExternalBankAccountNewResponseType = "CHECKING"
	ExternalBankAccountNewResponseTypeSavings  ExternalBankAccountNewResponseType = "SAVINGS"
)

func (r ExternalBankAccountNewResponseType) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewResponseTypeChecking, ExternalBankAccountNewResponseTypeSavings:
		return true
	}
	return false
}

// Verification Method
type ExternalBankAccountNewResponseVerificationMethod string

const (
	ExternalBankAccountNewResponseVerificationMethodManual       ExternalBankAccountNewResponseVerificationMethod = "MANUAL"
	ExternalBankAccountNewResponseVerificationMethodMicroDeposit ExternalBankAccountNewResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountNewResponseVerificationMethodPlaid        ExternalBankAccountNewResponseVerificationMethod = "PLAID"
	ExternalBankAccountNewResponseVerificationMethodPrenote      ExternalBankAccountNewResponseVerificationMethod = "PRENOTE"
)

func (r ExternalBankAccountNewResponseVerificationMethod) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewResponseVerificationMethodManual, ExternalBankAccountNewResponseVerificationMethodMicroDeposit, ExternalBankAccountNewResponseVerificationMethodPlaid, ExternalBankAccountNewResponseVerificationMethodPrenote:
		return true
	}
	return false
}

// Verification State
type ExternalBankAccountNewResponseVerificationState string

const (
	ExternalBankAccountNewResponseVerificationStatePending            ExternalBankAccountNewResponseVerificationState = "PENDING"
	ExternalBankAccountNewResponseVerificationStateEnabled            ExternalBankAccountNewResponseVerificationState = "ENABLED"
	ExternalBankAccountNewResponseVerificationStateFailedVerification ExternalBankAccountNewResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountNewResponseVerificationStateInsufficientFunds  ExternalBankAccountNewResponseVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ExternalBankAccountNewResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewResponseVerificationStatePending, ExternalBankAccountNewResponseVerificationStateEnabled, ExternalBankAccountNewResponseVerificationStateFailedVerification, ExternalBankAccountNewResponseVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type ExternalBankAccountGetResponse struct {
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
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency string `json:"currency,required"`
	// The last 4 digits of the bank account. Derived by Lithic from the account number
	// passed
	LastFour string `json:"last_four,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner string `json:"owner,required"`
	// Owner Type
	OwnerType ExternalBankAccountGetResponseOwnerType `json:"owner_type,required"`
	// Routing Number
	RoutingNumber string `json:"routing_number,required"`
	// Account State
	State ExternalBankAccountGetResponseState `json:"state,required"`
	// Account Type
	Type ExternalBankAccountGetResponseType `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64 `json:"verification_attempts,required"`
	// Verification Method
	VerificationMethod ExternalBankAccountGetResponseVerificationMethod `json:"verification_method,required"`
	// Verification State
	VerificationState ExternalBankAccountGetResponseVerificationState `json:"verification_state,required"`
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
	VerificationFailedReason string                             `json:"verification_failed_reason"`
	JSON                     externalBankAccountGetResponseJSON `json:"-"`
}

// externalBankAccountGetResponseJSON contains the JSON metadata for the struct
// [ExternalBankAccountGetResponse]
type externalBankAccountGetResponseJSON struct {
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

func (r *ExternalBankAccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountGetResponseJSON) RawJSON() string {
	return r.raw
}

// Owner Type
type ExternalBankAccountGetResponseOwnerType string

const (
	ExternalBankAccountGetResponseOwnerTypeBusiness   ExternalBankAccountGetResponseOwnerType = "BUSINESS"
	ExternalBankAccountGetResponseOwnerTypeIndividual ExternalBankAccountGetResponseOwnerType = "INDIVIDUAL"
)

func (r ExternalBankAccountGetResponseOwnerType) IsKnown() bool {
	switch r {
	case ExternalBankAccountGetResponseOwnerTypeBusiness, ExternalBankAccountGetResponseOwnerTypeIndividual:
		return true
	}
	return false
}

// Account State
type ExternalBankAccountGetResponseState string

const (
	ExternalBankAccountGetResponseStateEnabled ExternalBankAccountGetResponseState = "ENABLED"
	ExternalBankAccountGetResponseStateClosed  ExternalBankAccountGetResponseState = "CLOSED"
	ExternalBankAccountGetResponseStatePaused  ExternalBankAccountGetResponseState = "PAUSED"
)

func (r ExternalBankAccountGetResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountGetResponseStateEnabled, ExternalBankAccountGetResponseStateClosed, ExternalBankAccountGetResponseStatePaused:
		return true
	}
	return false
}

// Account Type
type ExternalBankAccountGetResponseType string

const (
	ExternalBankAccountGetResponseTypeChecking ExternalBankAccountGetResponseType = "CHECKING"
	ExternalBankAccountGetResponseTypeSavings  ExternalBankAccountGetResponseType = "SAVINGS"
)

func (r ExternalBankAccountGetResponseType) IsKnown() bool {
	switch r {
	case ExternalBankAccountGetResponseTypeChecking, ExternalBankAccountGetResponseTypeSavings:
		return true
	}
	return false
}

// Verification Method
type ExternalBankAccountGetResponseVerificationMethod string

const (
	ExternalBankAccountGetResponseVerificationMethodManual       ExternalBankAccountGetResponseVerificationMethod = "MANUAL"
	ExternalBankAccountGetResponseVerificationMethodMicroDeposit ExternalBankAccountGetResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountGetResponseVerificationMethodPlaid        ExternalBankAccountGetResponseVerificationMethod = "PLAID"
	ExternalBankAccountGetResponseVerificationMethodPrenote      ExternalBankAccountGetResponseVerificationMethod = "PRENOTE"
)

func (r ExternalBankAccountGetResponseVerificationMethod) IsKnown() bool {
	switch r {
	case ExternalBankAccountGetResponseVerificationMethodManual, ExternalBankAccountGetResponseVerificationMethodMicroDeposit, ExternalBankAccountGetResponseVerificationMethodPlaid, ExternalBankAccountGetResponseVerificationMethodPrenote:
		return true
	}
	return false
}

// Verification State
type ExternalBankAccountGetResponseVerificationState string

const (
	ExternalBankAccountGetResponseVerificationStatePending            ExternalBankAccountGetResponseVerificationState = "PENDING"
	ExternalBankAccountGetResponseVerificationStateEnabled            ExternalBankAccountGetResponseVerificationState = "ENABLED"
	ExternalBankAccountGetResponseVerificationStateFailedVerification ExternalBankAccountGetResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountGetResponseVerificationStateInsufficientFunds  ExternalBankAccountGetResponseVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ExternalBankAccountGetResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountGetResponseVerificationStatePending, ExternalBankAccountGetResponseVerificationStateEnabled, ExternalBankAccountGetResponseVerificationStateFailedVerification, ExternalBankAccountGetResponseVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type ExternalBankAccountUpdateResponse struct {
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
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency string `json:"currency,required"`
	// The last 4 digits of the bank account. Derived by Lithic from the account number
	// passed
	LastFour string `json:"last_four,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner string `json:"owner,required"`
	// Owner Type
	OwnerType ExternalBankAccountUpdateResponseOwnerType `json:"owner_type,required"`
	// Routing Number
	RoutingNumber string `json:"routing_number,required"`
	// Account State
	State ExternalBankAccountUpdateResponseState `json:"state,required"`
	// Account Type
	Type ExternalBankAccountUpdateResponseType `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64 `json:"verification_attempts,required"`
	// Verification Method
	VerificationMethod ExternalBankAccountUpdateResponseVerificationMethod `json:"verification_method,required"`
	// Verification State
	VerificationState ExternalBankAccountUpdateResponseVerificationState `json:"verification_state,required"`
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
	VerificationFailedReason string                                `json:"verification_failed_reason"`
	JSON                     externalBankAccountUpdateResponseJSON `json:"-"`
}

// externalBankAccountUpdateResponseJSON contains the JSON metadata for the struct
// [ExternalBankAccountUpdateResponse]
type externalBankAccountUpdateResponseJSON struct {
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

func (r *ExternalBankAccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Owner Type
type ExternalBankAccountUpdateResponseOwnerType string

const (
	ExternalBankAccountUpdateResponseOwnerTypeBusiness   ExternalBankAccountUpdateResponseOwnerType = "BUSINESS"
	ExternalBankAccountUpdateResponseOwnerTypeIndividual ExternalBankAccountUpdateResponseOwnerType = "INDIVIDUAL"
)

func (r ExternalBankAccountUpdateResponseOwnerType) IsKnown() bool {
	switch r {
	case ExternalBankAccountUpdateResponseOwnerTypeBusiness, ExternalBankAccountUpdateResponseOwnerTypeIndividual:
		return true
	}
	return false
}

// Account State
type ExternalBankAccountUpdateResponseState string

const (
	ExternalBankAccountUpdateResponseStateEnabled ExternalBankAccountUpdateResponseState = "ENABLED"
	ExternalBankAccountUpdateResponseStateClosed  ExternalBankAccountUpdateResponseState = "CLOSED"
	ExternalBankAccountUpdateResponseStatePaused  ExternalBankAccountUpdateResponseState = "PAUSED"
)

func (r ExternalBankAccountUpdateResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountUpdateResponseStateEnabled, ExternalBankAccountUpdateResponseStateClosed, ExternalBankAccountUpdateResponseStatePaused:
		return true
	}
	return false
}

// Account Type
type ExternalBankAccountUpdateResponseType string

const (
	ExternalBankAccountUpdateResponseTypeChecking ExternalBankAccountUpdateResponseType = "CHECKING"
	ExternalBankAccountUpdateResponseTypeSavings  ExternalBankAccountUpdateResponseType = "SAVINGS"
)

func (r ExternalBankAccountUpdateResponseType) IsKnown() bool {
	switch r {
	case ExternalBankAccountUpdateResponseTypeChecking, ExternalBankAccountUpdateResponseTypeSavings:
		return true
	}
	return false
}

// Verification Method
type ExternalBankAccountUpdateResponseVerificationMethod string

const (
	ExternalBankAccountUpdateResponseVerificationMethodManual       ExternalBankAccountUpdateResponseVerificationMethod = "MANUAL"
	ExternalBankAccountUpdateResponseVerificationMethodMicroDeposit ExternalBankAccountUpdateResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountUpdateResponseVerificationMethodPlaid        ExternalBankAccountUpdateResponseVerificationMethod = "PLAID"
	ExternalBankAccountUpdateResponseVerificationMethodPrenote      ExternalBankAccountUpdateResponseVerificationMethod = "PRENOTE"
)

func (r ExternalBankAccountUpdateResponseVerificationMethod) IsKnown() bool {
	switch r {
	case ExternalBankAccountUpdateResponseVerificationMethodManual, ExternalBankAccountUpdateResponseVerificationMethodMicroDeposit, ExternalBankAccountUpdateResponseVerificationMethodPlaid, ExternalBankAccountUpdateResponseVerificationMethodPrenote:
		return true
	}
	return false
}

// Verification State
type ExternalBankAccountUpdateResponseVerificationState string

const (
	ExternalBankAccountUpdateResponseVerificationStatePending            ExternalBankAccountUpdateResponseVerificationState = "PENDING"
	ExternalBankAccountUpdateResponseVerificationStateEnabled            ExternalBankAccountUpdateResponseVerificationState = "ENABLED"
	ExternalBankAccountUpdateResponseVerificationStateFailedVerification ExternalBankAccountUpdateResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountUpdateResponseVerificationStateInsufficientFunds  ExternalBankAccountUpdateResponseVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ExternalBankAccountUpdateResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountUpdateResponseVerificationStatePending, ExternalBankAccountUpdateResponseVerificationStateEnabled, ExternalBankAccountUpdateResponseVerificationStateFailedVerification, ExternalBankAccountUpdateResponseVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type ExternalBankAccountListResponse struct {
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
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency string `json:"currency,required"`
	// The last 4 digits of the bank account. Derived by Lithic from the account number
	// passed
	LastFour string `json:"last_four,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner string `json:"owner,required"`
	// Owner Type
	OwnerType ExternalBankAccountListResponseOwnerType `json:"owner_type,required"`
	// Routing Number
	RoutingNumber string `json:"routing_number,required"`
	// Account State
	State ExternalBankAccountListResponseState `json:"state,required"`
	// Account Type
	Type ExternalBankAccountListResponseType `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64 `json:"verification_attempts,required"`
	// Verification Method
	VerificationMethod ExternalBankAccountListResponseVerificationMethod `json:"verification_method,required"`
	// Verification State
	VerificationState ExternalBankAccountListResponseVerificationState `json:"verification_state,required"`
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
	VerificationFailedReason string                              `json:"verification_failed_reason"`
	JSON                     externalBankAccountListResponseJSON `json:"-"`
}

// externalBankAccountListResponseJSON contains the JSON metadata for the struct
// [ExternalBankAccountListResponse]
type externalBankAccountListResponseJSON struct {
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

func (r *ExternalBankAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountListResponseJSON) RawJSON() string {
	return r.raw
}

// Owner Type
type ExternalBankAccountListResponseOwnerType string

const (
	ExternalBankAccountListResponseOwnerTypeBusiness   ExternalBankAccountListResponseOwnerType = "BUSINESS"
	ExternalBankAccountListResponseOwnerTypeIndividual ExternalBankAccountListResponseOwnerType = "INDIVIDUAL"
)

func (r ExternalBankAccountListResponseOwnerType) IsKnown() bool {
	switch r {
	case ExternalBankAccountListResponseOwnerTypeBusiness, ExternalBankAccountListResponseOwnerTypeIndividual:
		return true
	}
	return false
}

// Account State
type ExternalBankAccountListResponseState string

const (
	ExternalBankAccountListResponseStateEnabled ExternalBankAccountListResponseState = "ENABLED"
	ExternalBankAccountListResponseStateClosed  ExternalBankAccountListResponseState = "CLOSED"
	ExternalBankAccountListResponseStatePaused  ExternalBankAccountListResponseState = "PAUSED"
)

func (r ExternalBankAccountListResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountListResponseStateEnabled, ExternalBankAccountListResponseStateClosed, ExternalBankAccountListResponseStatePaused:
		return true
	}
	return false
}

// Account Type
type ExternalBankAccountListResponseType string

const (
	ExternalBankAccountListResponseTypeChecking ExternalBankAccountListResponseType = "CHECKING"
	ExternalBankAccountListResponseTypeSavings  ExternalBankAccountListResponseType = "SAVINGS"
)

func (r ExternalBankAccountListResponseType) IsKnown() bool {
	switch r {
	case ExternalBankAccountListResponseTypeChecking, ExternalBankAccountListResponseTypeSavings:
		return true
	}
	return false
}

// Verification Method
type ExternalBankAccountListResponseVerificationMethod string

const (
	ExternalBankAccountListResponseVerificationMethodManual       ExternalBankAccountListResponseVerificationMethod = "MANUAL"
	ExternalBankAccountListResponseVerificationMethodMicroDeposit ExternalBankAccountListResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountListResponseVerificationMethodPlaid        ExternalBankAccountListResponseVerificationMethod = "PLAID"
	ExternalBankAccountListResponseVerificationMethodPrenote      ExternalBankAccountListResponseVerificationMethod = "PRENOTE"
)

func (r ExternalBankAccountListResponseVerificationMethod) IsKnown() bool {
	switch r {
	case ExternalBankAccountListResponseVerificationMethodManual, ExternalBankAccountListResponseVerificationMethodMicroDeposit, ExternalBankAccountListResponseVerificationMethodPlaid, ExternalBankAccountListResponseVerificationMethodPrenote:
		return true
	}
	return false
}

// Verification State
type ExternalBankAccountListResponseVerificationState string

const (
	ExternalBankAccountListResponseVerificationStatePending            ExternalBankAccountListResponseVerificationState = "PENDING"
	ExternalBankAccountListResponseVerificationStateEnabled            ExternalBankAccountListResponseVerificationState = "ENABLED"
	ExternalBankAccountListResponseVerificationStateFailedVerification ExternalBankAccountListResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountListResponseVerificationStateInsufficientFunds  ExternalBankAccountListResponseVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ExternalBankAccountListResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountListResponseVerificationStatePending, ExternalBankAccountListResponseVerificationStateEnabled, ExternalBankAccountListResponseVerificationStateFailedVerification, ExternalBankAccountListResponseVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type ExternalBankAccountRetryMicroDepositsResponse struct {
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
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency string `json:"currency,required"`
	// The last 4 digits of the bank account. Derived by Lithic from the account number
	// passed
	LastFour string `json:"last_four,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner string `json:"owner,required"`
	// Owner Type
	OwnerType ExternalBankAccountRetryMicroDepositsResponseOwnerType `json:"owner_type,required"`
	// Routing Number
	RoutingNumber string `json:"routing_number,required"`
	// Account State
	State ExternalBankAccountRetryMicroDepositsResponseState `json:"state,required"`
	// Account Type
	Type ExternalBankAccountRetryMicroDepositsResponseType `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64 `json:"verification_attempts,required"`
	// Verification Method
	VerificationMethod ExternalBankAccountRetryMicroDepositsResponseVerificationMethod `json:"verification_method,required"`
	// Verification State
	VerificationState ExternalBankAccountRetryMicroDepositsResponseVerificationState `json:"verification_state,required"`
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
	VerificationFailedReason string                                            `json:"verification_failed_reason"`
	JSON                     externalBankAccountRetryMicroDepositsResponseJSON `json:"-"`
}

// externalBankAccountRetryMicroDepositsResponseJSON contains the JSON metadata for
// the struct [ExternalBankAccountRetryMicroDepositsResponse]
type externalBankAccountRetryMicroDepositsResponseJSON struct {
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

func (r *ExternalBankAccountRetryMicroDepositsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountRetryMicroDepositsResponseJSON) RawJSON() string {
	return r.raw
}

// Owner Type
type ExternalBankAccountRetryMicroDepositsResponseOwnerType string

const (
	ExternalBankAccountRetryMicroDepositsResponseOwnerTypeBusiness   ExternalBankAccountRetryMicroDepositsResponseOwnerType = "BUSINESS"
	ExternalBankAccountRetryMicroDepositsResponseOwnerTypeIndividual ExternalBankAccountRetryMicroDepositsResponseOwnerType = "INDIVIDUAL"
)

func (r ExternalBankAccountRetryMicroDepositsResponseOwnerType) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryMicroDepositsResponseOwnerTypeBusiness, ExternalBankAccountRetryMicroDepositsResponseOwnerTypeIndividual:
		return true
	}
	return false
}

// Account State
type ExternalBankAccountRetryMicroDepositsResponseState string

const (
	ExternalBankAccountRetryMicroDepositsResponseStateEnabled ExternalBankAccountRetryMicroDepositsResponseState = "ENABLED"
	ExternalBankAccountRetryMicroDepositsResponseStateClosed  ExternalBankAccountRetryMicroDepositsResponseState = "CLOSED"
	ExternalBankAccountRetryMicroDepositsResponseStatePaused  ExternalBankAccountRetryMicroDepositsResponseState = "PAUSED"
)

func (r ExternalBankAccountRetryMicroDepositsResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryMicroDepositsResponseStateEnabled, ExternalBankAccountRetryMicroDepositsResponseStateClosed, ExternalBankAccountRetryMicroDepositsResponseStatePaused:
		return true
	}
	return false
}

// Account Type
type ExternalBankAccountRetryMicroDepositsResponseType string

const (
	ExternalBankAccountRetryMicroDepositsResponseTypeChecking ExternalBankAccountRetryMicroDepositsResponseType = "CHECKING"
	ExternalBankAccountRetryMicroDepositsResponseTypeSavings  ExternalBankAccountRetryMicroDepositsResponseType = "SAVINGS"
)

func (r ExternalBankAccountRetryMicroDepositsResponseType) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryMicroDepositsResponseTypeChecking, ExternalBankAccountRetryMicroDepositsResponseTypeSavings:
		return true
	}
	return false
}

// Verification Method
type ExternalBankAccountRetryMicroDepositsResponseVerificationMethod string

const (
	ExternalBankAccountRetryMicroDepositsResponseVerificationMethodManual       ExternalBankAccountRetryMicroDepositsResponseVerificationMethod = "MANUAL"
	ExternalBankAccountRetryMicroDepositsResponseVerificationMethodMicroDeposit ExternalBankAccountRetryMicroDepositsResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountRetryMicroDepositsResponseVerificationMethodPlaid        ExternalBankAccountRetryMicroDepositsResponseVerificationMethod = "PLAID"
	ExternalBankAccountRetryMicroDepositsResponseVerificationMethodPrenote      ExternalBankAccountRetryMicroDepositsResponseVerificationMethod = "PRENOTE"
)

func (r ExternalBankAccountRetryMicroDepositsResponseVerificationMethod) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryMicroDepositsResponseVerificationMethodManual, ExternalBankAccountRetryMicroDepositsResponseVerificationMethodMicroDeposit, ExternalBankAccountRetryMicroDepositsResponseVerificationMethodPlaid, ExternalBankAccountRetryMicroDepositsResponseVerificationMethodPrenote:
		return true
	}
	return false
}

// Verification State
type ExternalBankAccountRetryMicroDepositsResponseVerificationState string

const (
	ExternalBankAccountRetryMicroDepositsResponseVerificationStatePending            ExternalBankAccountRetryMicroDepositsResponseVerificationState = "PENDING"
	ExternalBankAccountRetryMicroDepositsResponseVerificationStateEnabled            ExternalBankAccountRetryMicroDepositsResponseVerificationState = "ENABLED"
	ExternalBankAccountRetryMicroDepositsResponseVerificationStateFailedVerification ExternalBankAccountRetryMicroDepositsResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountRetryMicroDepositsResponseVerificationStateInsufficientFunds  ExternalBankAccountRetryMicroDepositsResponseVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ExternalBankAccountRetryMicroDepositsResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryMicroDepositsResponseVerificationStatePending, ExternalBankAccountRetryMicroDepositsResponseVerificationStateEnabled, ExternalBankAccountRetryMicroDepositsResponseVerificationStateFailedVerification, ExternalBankAccountRetryMicroDepositsResponseVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type ExternalBankAccountRetryPrenoteResponse struct {
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
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency string `json:"currency,required"`
	// The last 4 digits of the bank account. Derived by Lithic from the account number
	// passed
	LastFour string `json:"last_four,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner string `json:"owner,required"`
	// Owner Type
	OwnerType OwnerType `json:"owner_type,required"`
	// Routing Number
	RoutingNumber string `json:"routing_number,required"`
	// Account State
	State ExternalBankAccountRetryPrenoteResponseState `json:"state,required"`
	// Account Type
	Type ExternalBankAccountRetryPrenoteResponseType `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64 `json:"verification_attempts,required"`
	// Verification Method
	VerificationMethod VerificationMethod `json:"verification_method,required"`
	// Verification State
	VerificationState ExternalBankAccountRetryPrenoteResponseVerificationState `json:"verification_state,required"`
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
	VerificationFailedReason string                                      `json:"verification_failed_reason"`
	JSON                     externalBankAccountRetryPrenoteResponseJSON `json:"-"`
}

// externalBankAccountRetryPrenoteResponseJSON contains the JSON metadata for the
// struct [ExternalBankAccountRetryPrenoteResponse]
type externalBankAccountRetryPrenoteResponseJSON struct {
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

func (r *ExternalBankAccountRetryPrenoteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalBankAccountRetryPrenoteResponseJSON) RawJSON() string {
	return r.raw
}

// Account State
type ExternalBankAccountRetryPrenoteResponseState string

const (
	ExternalBankAccountRetryPrenoteResponseStateEnabled ExternalBankAccountRetryPrenoteResponseState = "ENABLED"
	ExternalBankAccountRetryPrenoteResponseStateClosed  ExternalBankAccountRetryPrenoteResponseState = "CLOSED"
	ExternalBankAccountRetryPrenoteResponseStatePaused  ExternalBankAccountRetryPrenoteResponseState = "PAUSED"
)

func (r ExternalBankAccountRetryPrenoteResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryPrenoteResponseStateEnabled, ExternalBankAccountRetryPrenoteResponseStateClosed, ExternalBankAccountRetryPrenoteResponseStatePaused:
		return true
	}
	return false
}

// Account Type
type ExternalBankAccountRetryPrenoteResponseType string

const (
	ExternalBankAccountRetryPrenoteResponseTypeChecking ExternalBankAccountRetryPrenoteResponseType = "CHECKING"
	ExternalBankAccountRetryPrenoteResponseTypeSavings  ExternalBankAccountRetryPrenoteResponseType = "SAVINGS"
)

func (r ExternalBankAccountRetryPrenoteResponseType) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryPrenoteResponseTypeChecking, ExternalBankAccountRetryPrenoteResponseTypeSavings:
		return true
	}
	return false
}

// Verification State
type ExternalBankAccountRetryPrenoteResponseVerificationState string

const (
	ExternalBankAccountRetryPrenoteResponseVerificationStatePending            ExternalBankAccountRetryPrenoteResponseVerificationState = "PENDING"
	ExternalBankAccountRetryPrenoteResponseVerificationStateEnabled            ExternalBankAccountRetryPrenoteResponseVerificationState = "ENABLED"
	ExternalBankAccountRetryPrenoteResponseVerificationStateFailedVerification ExternalBankAccountRetryPrenoteResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountRetryPrenoteResponseVerificationStateInsufficientFunds  ExternalBankAccountRetryPrenoteResponseVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ExternalBankAccountRetryPrenoteResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryPrenoteResponseVerificationStatePending, ExternalBankAccountRetryPrenoteResponseVerificationStateEnabled, ExternalBankAccountRetryPrenoteResponseVerificationStateFailedVerification, ExternalBankAccountRetryPrenoteResponseVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type ExternalBankAccountNewParams struct {
	Body ExternalBankAccountNewParamsBodyUnion `json:"body,required"`
}

func (r ExternalBankAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ExternalBankAccountNewParamsBody struct {
	// Verification Method
	VerificationMethod param.Field[VerificationMethod] `json:"verification_method,required"`
	// Owner Type
	OwnerType param.Field[OwnerType] `json:"owner_type,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner param.Field[string] `json:"owner,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken param.Field[string] `json:"account_token" format:"uuid"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID param.Field[string] `json:"company_id"`
	// Doing Business As
	DoingBusinessAs param.Field[string] `json:"doing_business_as"`
	// Date of Birth of the Individual that owns the external bank account
	Dob param.Field[time.Time] `json:"dob" format:"date"`
	// User Defined ID
	UserDefinedID param.Field[string] `json:"user_defined_id"`
	// Account Type
	Type param.Field[ExternalBankAccountNewParamsBodyType] `json:"type"`
	// Routing Number
	RoutingNumber param.Field[string] `json:"routing_number"`
	// Account Number
	AccountNumber param.Field[string] `json:"account_number"`
	// The nickname for this External Bank Account
	Name param.Field[string] `json:"name"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept USA bank accounts e.g., USA
	Country param.Field[string] `json:"country"`
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency                param.Field[string] `json:"currency"`
	VerificationEnforcement param.Field[bool]   `json:"verification_enforcement"`
	// Address
	Address param.Field[ExternalBankAccountAddressParam] `json:"address"`
	// The financial account token of the operating account to fund the micro deposits
	FinancialAccountToken param.Field[string] `json:"financial_account_token" format:"uuid"`
	ProcessorToken        param.Field[string] `json:"processor_token"`
}

func (r ExternalBankAccountNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalBankAccountNewParamsBody) implementsExternalBankAccountNewParamsBodyUnion() {}

// Satisfied by
// [ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest],
// [ExternalBankAccountNewParamsBodyPlaidCreateBankAccountAPIRequest],
// [ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequest],
// [ExternalBankAccountNewParamsBody].
type ExternalBankAccountNewParamsBodyUnion interface {
	implementsExternalBankAccountNewParamsBodyUnion()
}

type ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest struct {
	// Account Number
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept USA bank accounts e.g., USA
	Country param.Field[string] `json:"country,required"`
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency param.Field[string] `json:"currency,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner param.Field[string] `json:"owner,required"`
	// Owner Type
	OwnerType param.Field[OwnerType] `json:"owner_type,required"`
	// Routing Number
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// Account Type
	Type param.Field[ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestType] `json:"type,required"`
	// Verification Method
	VerificationMethod param.Field[VerificationMethod] `json:"verification_method,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken param.Field[string] `json:"account_token" format:"uuid"`
	// Address
	Address param.Field[ExternalBankAccountAddressParam] `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID param.Field[string] `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob param.Field[time.Time] `json:"dob" format:"date"`
	// Doing Business As
	DoingBusinessAs param.Field[string] `json:"doing_business_as"`
	// The financial account token of the operating account to fund the micro deposits
	FinancialAccountToken param.Field[string] `json:"financial_account_token" format:"uuid"`
	// The nickname for this External Bank Account
	Name param.Field[string] `json:"name"`
	// User Defined ID
	UserDefinedID           param.Field[string] `json:"user_defined_id"`
	VerificationEnforcement param.Field[bool]   `json:"verification_enforcement"`
}

func (r ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest) implementsExternalBankAccountNewParamsBodyUnion() {
}

// Account Type
type ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestType string

const (
	ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestTypeChecking ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestType = "CHECKING"
	ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestTypeSavings  ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestType = "SAVINGS"
)

func (r ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestType) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestTypeChecking, ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestTypeSavings:
		return true
	}
	return false
}

type ExternalBankAccountNewParamsBodyPlaidCreateBankAccountAPIRequest struct {
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner param.Field[string] `json:"owner,required"`
	// Owner Type
	OwnerType      param.Field[OwnerType] `json:"owner_type,required"`
	ProcessorToken param.Field[string]    `json:"processor_token,required"`
	// Verification Method
	VerificationMethod param.Field[VerificationMethod] `json:"verification_method,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken param.Field[string] `json:"account_token" format:"uuid"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID param.Field[string] `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob param.Field[time.Time] `json:"dob" format:"date"`
	// Doing Business As
	DoingBusinessAs param.Field[string] `json:"doing_business_as"`
	// User Defined ID
	UserDefinedID param.Field[string] `json:"user_defined_id"`
}

func (r ExternalBankAccountNewParamsBodyPlaidCreateBankAccountAPIRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalBankAccountNewParamsBodyPlaidCreateBankAccountAPIRequest) implementsExternalBankAccountNewParamsBodyUnion() {
}

type ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequest struct {
	// Account Number
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept USA bank accounts e.g., USA
	Country param.Field[string] `json:"country,required"`
	// currency of the external account 3-digit alphabetic ISO 4217 code
	Currency param.Field[string] `json:"currency,required"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner param.Field[string] `json:"owner,required"`
	// Owner Type
	OwnerType param.Field[OwnerType] `json:"owner_type,required"`
	// Routing Number
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// Account Type
	Type param.Field[ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestType] `json:"type,required"`
	// Verification Method
	VerificationMethod param.Field[ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestVerificationMethod] `json:"verification_method,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken param.Field[string] `json:"account_token" format:"uuid"`
	// Address
	Address param.Field[ExternalBankAccountAddressParam] `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID param.Field[string] `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob param.Field[time.Time] `json:"dob" format:"date"`
	// Doing Business As
	DoingBusinessAs param.Field[string] `json:"doing_business_as"`
	// The nickname for this External Bank Account
	Name param.Field[string] `json:"name"`
	// User Defined ID
	UserDefinedID param.Field[string] `json:"user_defined_id"`
}

func (r ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequest) implementsExternalBankAccountNewParamsBodyUnion() {
}

// Account Type
type ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestType string

const (
	ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestTypeChecking ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestType = "CHECKING"
	ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestTypeSavings  ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestType = "SAVINGS"
)

func (r ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestType) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestTypeChecking, ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestTypeSavings:
		return true
	}
	return false
}

// Verification Method
type ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestVerificationMethod string

const (
	ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestVerificationMethodExternallyVerified ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestVerificationMethod = "EXTERNALLY_VERIFIED"
)

func (r ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestVerificationMethod) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewParamsBodyExternallyVerifiedCreateBankAccountAPIRequestVerificationMethodExternallyVerified:
		return true
	}
	return false
}

// Account Type
type ExternalBankAccountNewParamsBodyType string

const (
	ExternalBankAccountNewParamsBodyTypeChecking ExternalBankAccountNewParamsBodyType = "CHECKING"
	ExternalBankAccountNewParamsBodyTypeSavings  ExternalBankAccountNewParamsBodyType = "SAVINGS"
)

func (r ExternalBankAccountNewParamsBodyType) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewParamsBodyTypeChecking, ExternalBankAccountNewParamsBodyTypeSavings:
		return true
	}
	return false
}

type ExternalBankAccountUpdateParams struct {
	// Address
	Address param.Field[ExternalBankAccountAddressParam] `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID param.Field[string] `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob param.Field[time.Time] `json:"dob" format:"date"`
	// Doing Business As
	DoingBusinessAs param.Field[string] `json:"doing_business_as"`
	// The nickname for this External Bank Account
	Name param.Field[string] `json:"name"`
	// Legal Name of the business or individual who owns the external account. This
	// will appear in statements
	Owner param.Field[string] `json:"owner"`
	// Owner Type
	OwnerType param.Field[OwnerType] `json:"owner_type"`
	// User Defined ID
	UserDefinedID param.Field[string] `json:"user_defined_id"`
}

func (r ExternalBankAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalBankAccountListParams struct {
	AccountToken param.Field[string]                                     `query:"account_token" format:"uuid"`
	AccountTypes param.Field[[]ExternalBankAccountListParamsAccountType] `query:"account_types"`
	Countries    param.Field[[]string]                                   `query:"countries"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string]      `query:"ending_before"`
	OwnerTypes   param.Field[[]OwnerType] `query:"owner_types"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter      param.Field[string]                                           `query:"starting_after"`
	States             param.Field[[]ExternalBankAccountListParamsState]             `query:"states"`
	VerificationStates param.Field[[]ExternalBankAccountListParamsVerificationState] `query:"verification_states"`
}

// URLQuery serializes [ExternalBankAccountListParams]'s query parameters as
// `url.Values`.
func (r ExternalBankAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalBankAccountListParamsAccountType string

const (
	ExternalBankAccountListParamsAccountTypeChecking ExternalBankAccountListParamsAccountType = "CHECKING"
	ExternalBankAccountListParamsAccountTypeSavings  ExternalBankAccountListParamsAccountType = "SAVINGS"
)

func (r ExternalBankAccountListParamsAccountType) IsKnown() bool {
	switch r {
	case ExternalBankAccountListParamsAccountTypeChecking, ExternalBankAccountListParamsAccountTypeSavings:
		return true
	}
	return false
}

type ExternalBankAccountListParamsState string

const (
	ExternalBankAccountListParamsStateEnabled ExternalBankAccountListParamsState = "ENABLED"
	ExternalBankAccountListParamsStateClosed  ExternalBankAccountListParamsState = "CLOSED"
	ExternalBankAccountListParamsStatePaused  ExternalBankAccountListParamsState = "PAUSED"
)

func (r ExternalBankAccountListParamsState) IsKnown() bool {
	switch r {
	case ExternalBankAccountListParamsStateEnabled, ExternalBankAccountListParamsStateClosed, ExternalBankAccountListParamsStatePaused:
		return true
	}
	return false
}

type ExternalBankAccountListParamsVerificationState string

const (
	ExternalBankAccountListParamsVerificationStatePending            ExternalBankAccountListParamsVerificationState = "PENDING"
	ExternalBankAccountListParamsVerificationStateEnabled            ExternalBankAccountListParamsVerificationState = "ENABLED"
	ExternalBankAccountListParamsVerificationStateFailedVerification ExternalBankAccountListParamsVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountListParamsVerificationStateInsufficientFunds  ExternalBankAccountListParamsVerificationState = "INSUFFICIENT_FUNDS"
)

func (r ExternalBankAccountListParamsVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountListParamsVerificationStatePending, ExternalBankAccountListParamsVerificationStateEnabled, ExternalBankAccountListParamsVerificationStateFailedVerification, ExternalBankAccountListParamsVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type ExternalBankAccountRetryMicroDepositsParams struct {
	FinancialAccountToken param.Field[string] `json:"financial_account_token" format:"uuid"`
}

func (r ExternalBankAccountRetryMicroDepositsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalBankAccountRetryPrenoteParams struct {
	FinancialAccountToken param.Field[string] `json:"financial_account_token" format:"uuid"`
}

func (r ExternalBankAccountRetryPrenoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
