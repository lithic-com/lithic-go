// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
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
	path := fmt.Sprintf("external_bank_accounts/%s", externalBankAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the external bank account by token.
func (r *ExternalBankAccountService) Update(ctx context.Context, externalBankAccountToken string, body ExternalBankAccountUpdateParams, opts ...option.RequestOption) (res *ExternalBankAccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("external_bank_accounts/%s", externalBankAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all the external bank accounts for the provided search criteria.
func (r *ExternalBankAccountService) List(ctx context.Context, query ExternalBankAccountListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ExternalBankAccountListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
	path := fmt.Sprintf("external_bank_accounts/%s/retry_micro_deposits", externalBankAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Address used during Address Verification Service (AVS) checks during
// transactions if enabled via Auth Rules.
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

// Address used during Address Verification Service (AVS) checks during
// transactions if enabled via Auth Rules.
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
	OwnerTypeBusiness   OwnerType = "BUSINESS"
	OwnerTypeIndividual OwnerType = "INDIVIDUAL"
)

func (r OwnerType) IsKnown() bool {
	switch r {
	case OwnerTypeBusiness, OwnerTypeIndividual:
		return true
	}
	return false
}

type VerificationMethod string

const (
	VerificationMethodManual       VerificationMethod = "MANUAL"
	VerificationMethodMicroDeposit VerificationMethod = "MICRO_DEPOSIT"
	VerificationMethodPlaid        VerificationMethod = "PLAID"
	VerificationMethodPrenote      VerificationMethod = "PRENOTE"
)

func (r VerificationMethod) IsKnown() bool {
	switch r {
	case VerificationMethodManual, VerificationMethodMicroDeposit, VerificationMethodPlaid, VerificationMethodPrenote:
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
	Owner         string                                  `json:"owner,required"`
	OwnerType     ExternalBankAccountNewResponseOwnerType `json:"owner_type,required"`
	RoutingNumber string                                  `json:"routing_number,required"`
	State         ExternalBankAccountNewResponseState     `json:"state,required"`
	Type          ExternalBankAccountNewResponseType      `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64                                            `json:"verification_attempts,required"`
	VerificationMethod   ExternalBankAccountNewResponseVerificationMethod `json:"verification_method,required"`
	VerificationState    ExternalBankAccountNewResponseVerificationState  `json:"verification_state,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken string `json:"account_token" format:"uuid"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address ExternalBankAccountAddress `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID string `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             time.Time `json:"dob" format:"date"`
	DoingBusinessAs string    `json:"doing_business_as"`
	// The financial account token of the operating account, which will provide the
	// funds for micro deposits used to verify the account
	FinancialAccountToken string `json:"financial_account_token" format:"uuid"`
	// The nickname given to this record of External Bank Account
	Name          string `json:"name"`
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

type ExternalBankAccountNewResponseState string

const (
	ExternalBankAccountNewResponseStateClosed  ExternalBankAccountNewResponseState = "CLOSED"
	ExternalBankAccountNewResponseStateEnabled ExternalBankAccountNewResponseState = "ENABLED"
	ExternalBankAccountNewResponseStatePaused  ExternalBankAccountNewResponseState = "PAUSED"
)

func (r ExternalBankAccountNewResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewResponseStateClosed, ExternalBankAccountNewResponseStateEnabled, ExternalBankAccountNewResponseStatePaused:
		return true
	}
	return false
}

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

type ExternalBankAccountNewResponseVerificationState string

const (
	ExternalBankAccountNewResponseVerificationStateEnabled            ExternalBankAccountNewResponseVerificationState = "ENABLED"
	ExternalBankAccountNewResponseVerificationStateFailedVerification ExternalBankAccountNewResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountNewResponseVerificationStateInsufficientFunds  ExternalBankAccountNewResponseVerificationState = "INSUFFICIENT_FUNDS"
	ExternalBankAccountNewResponseVerificationStatePending            ExternalBankAccountNewResponseVerificationState = "PENDING"
)

func (r ExternalBankAccountNewResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountNewResponseVerificationStateEnabled, ExternalBankAccountNewResponseVerificationStateFailedVerification, ExternalBankAccountNewResponseVerificationStateInsufficientFunds, ExternalBankAccountNewResponseVerificationStatePending:
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
	Owner         string                                  `json:"owner,required"`
	OwnerType     ExternalBankAccountGetResponseOwnerType `json:"owner_type,required"`
	RoutingNumber string                                  `json:"routing_number,required"`
	State         ExternalBankAccountGetResponseState     `json:"state,required"`
	Type          ExternalBankAccountGetResponseType      `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64                                            `json:"verification_attempts,required"`
	VerificationMethod   ExternalBankAccountGetResponseVerificationMethod `json:"verification_method,required"`
	VerificationState    ExternalBankAccountGetResponseVerificationState  `json:"verification_state,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken string `json:"account_token" format:"uuid"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address ExternalBankAccountAddress `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID string `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             time.Time `json:"dob" format:"date"`
	DoingBusinessAs string    `json:"doing_business_as"`
	// The financial account token of the operating account, which will provide the
	// funds for micro deposits used to verify the account
	FinancialAccountToken string `json:"financial_account_token" format:"uuid"`
	// The nickname given to this record of External Bank Account
	Name          string `json:"name"`
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

type ExternalBankAccountGetResponseState string

const (
	ExternalBankAccountGetResponseStateClosed  ExternalBankAccountGetResponseState = "CLOSED"
	ExternalBankAccountGetResponseStateEnabled ExternalBankAccountGetResponseState = "ENABLED"
	ExternalBankAccountGetResponseStatePaused  ExternalBankAccountGetResponseState = "PAUSED"
)

func (r ExternalBankAccountGetResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountGetResponseStateClosed, ExternalBankAccountGetResponseStateEnabled, ExternalBankAccountGetResponseStatePaused:
		return true
	}
	return false
}

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

type ExternalBankAccountGetResponseVerificationState string

const (
	ExternalBankAccountGetResponseVerificationStateEnabled            ExternalBankAccountGetResponseVerificationState = "ENABLED"
	ExternalBankAccountGetResponseVerificationStateFailedVerification ExternalBankAccountGetResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountGetResponseVerificationStateInsufficientFunds  ExternalBankAccountGetResponseVerificationState = "INSUFFICIENT_FUNDS"
	ExternalBankAccountGetResponseVerificationStatePending            ExternalBankAccountGetResponseVerificationState = "PENDING"
)

func (r ExternalBankAccountGetResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountGetResponseVerificationStateEnabled, ExternalBankAccountGetResponseVerificationStateFailedVerification, ExternalBankAccountGetResponseVerificationStateInsufficientFunds, ExternalBankAccountGetResponseVerificationStatePending:
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
	Owner         string                                     `json:"owner,required"`
	OwnerType     ExternalBankAccountUpdateResponseOwnerType `json:"owner_type,required"`
	RoutingNumber string                                     `json:"routing_number,required"`
	State         ExternalBankAccountUpdateResponseState     `json:"state,required"`
	Type          ExternalBankAccountUpdateResponseType      `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64                                               `json:"verification_attempts,required"`
	VerificationMethod   ExternalBankAccountUpdateResponseVerificationMethod `json:"verification_method,required"`
	VerificationState    ExternalBankAccountUpdateResponseVerificationState  `json:"verification_state,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken string `json:"account_token" format:"uuid"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address ExternalBankAccountAddress `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID string `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             time.Time `json:"dob" format:"date"`
	DoingBusinessAs string    `json:"doing_business_as"`
	// The financial account token of the operating account, which will provide the
	// funds for micro deposits used to verify the account
	FinancialAccountToken string `json:"financial_account_token" format:"uuid"`
	// The nickname given to this record of External Bank Account
	Name          string `json:"name"`
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

type ExternalBankAccountUpdateResponseState string

const (
	ExternalBankAccountUpdateResponseStateClosed  ExternalBankAccountUpdateResponseState = "CLOSED"
	ExternalBankAccountUpdateResponseStateEnabled ExternalBankAccountUpdateResponseState = "ENABLED"
	ExternalBankAccountUpdateResponseStatePaused  ExternalBankAccountUpdateResponseState = "PAUSED"
)

func (r ExternalBankAccountUpdateResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountUpdateResponseStateClosed, ExternalBankAccountUpdateResponseStateEnabled, ExternalBankAccountUpdateResponseStatePaused:
		return true
	}
	return false
}

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

type ExternalBankAccountUpdateResponseVerificationState string

const (
	ExternalBankAccountUpdateResponseVerificationStateEnabled            ExternalBankAccountUpdateResponseVerificationState = "ENABLED"
	ExternalBankAccountUpdateResponseVerificationStateFailedVerification ExternalBankAccountUpdateResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountUpdateResponseVerificationStateInsufficientFunds  ExternalBankAccountUpdateResponseVerificationState = "INSUFFICIENT_FUNDS"
	ExternalBankAccountUpdateResponseVerificationStatePending            ExternalBankAccountUpdateResponseVerificationState = "PENDING"
)

func (r ExternalBankAccountUpdateResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountUpdateResponseVerificationStateEnabled, ExternalBankAccountUpdateResponseVerificationStateFailedVerification, ExternalBankAccountUpdateResponseVerificationStateInsufficientFunds, ExternalBankAccountUpdateResponseVerificationStatePending:
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
	Owner         string                                   `json:"owner,required"`
	OwnerType     ExternalBankAccountListResponseOwnerType `json:"owner_type,required"`
	RoutingNumber string                                   `json:"routing_number,required"`
	State         ExternalBankAccountListResponseState     `json:"state,required"`
	Type          ExternalBankAccountListResponseType      `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64                                             `json:"verification_attempts,required"`
	VerificationMethod   ExternalBankAccountListResponseVerificationMethod `json:"verification_method,required"`
	VerificationState    ExternalBankAccountListResponseVerificationState  `json:"verification_state,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken string `json:"account_token" format:"uuid"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address ExternalBankAccountAddress `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID string `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             time.Time `json:"dob" format:"date"`
	DoingBusinessAs string    `json:"doing_business_as"`
	// The financial account token of the operating account, which will provide the
	// funds for micro deposits used to verify the account
	FinancialAccountToken string `json:"financial_account_token" format:"uuid"`
	// The nickname given to this record of External Bank Account
	Name          string `json:"name"`
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

type ExternalBankAccountListResponseState string

const (
	ExternalBankAccountListResponseStateClosed  ExternalBankAccountListResponseState = "CLOSED"
	ExternalBankAccountListResponseStateEnabled ExternalBankAccountListResponseState = "ENABLED"
	ExternalBankAccountListResponseStatePaused  ExternalBankAccountListResponseState = "PAUSED"
)

func (r ExternalBankAccountListResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountListResponseStateClosed, ExternalBankAccountListResponseStateEnabled, ExternalBankAccountListResponseStatePaused:
		return true
	}
	return false
}

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

type ExternalBankAccountListResponseVerificationState string

const (
	ExternalBankAccountListResponseVerificationStateEnabled            ExternalBankAccountListResponseVerificationState = "ENABLED"
	ExternalBankAccountListResponseVerificationStateFailedVerification ExternalBankAccountListResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountListResponseVerificationStateInsufficientFunds  ExternalBankAccountListResponseVerificationState = "INSUFFICIENT_FUNDS"
	ExternalBankAccountListResponseVerificationStatePending            ExternalBankAccountListResponseVerificationState = "PENDING"
)

func (r ExternalBankAccountListResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountListResponseVerificationStateEnabled, ExternalBankAccountListResponseVerificationStateFailedVerification, ExternalBankAccountListResponseVerificationStateInsufficientFunds, ExternalBankAccountListResponseVerificationStatePending:
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
	Owner         string                                                 `json:"owner,required"`
	OwnerType     ExternalBankAccountRetryMicroDepositsResponseOwnerType `json:"owner_type,required"`
	RoutingNumber string                                                 `json:"routing_number,required"`
	State         ExternalBankAccountRetryMicroDepositsResponseState     `json:"state,required"`
	Type          ExternalBankAccountRetryMicroDepositsResponseType      `json:"type,required"`
	// The number of attempts at verification
	VerificationAttempts int64                                                           `json:"verification_attempts,required"`
	VerificationMethod   ExternalBankAccountRetryMicroDepositsResponseVerificationMethod `json:"verification_method,required"`
	VerificationState    ExternalBankAccountRetryMicroDepositsResponseVerificationState  `json:"verification_state,required"`
	// Indicates which Lithic account the external account is associated with. For
	// external accounts that are associated with the program, account_token field
	// returned will be null
	AccountToken string `json:"account_token" format:"uuid"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address ExternalBankAccountAddress `json:"address"`
	// Optional field that helps identify bank accounts in receipts
	CompanyID string `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             time.Time `json:"dob" format:"date"`
	DoingBusinessAs string    `json:"doing_business_as"`
	// The financial account token of the operating account, which will provide the
	// funds for micro deposits used to verify the account
	FinancialAccountToken string `json:"financial_account_token" format:"uuid"`
	// The nickname given to this record of External Bank Account
	Name          string `json:"name"`
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

type ExternalBankAccountRetryMicroDepositsResponseState string

const (
	ExternalBankAccountRetryMicroDepositsResponseStateClosed  ExternalBankAccountRetryMicroDepositsResponseState = "CLOSED"
	ExternalBankAccountRetryMicroDepositsResponseStateEnabled ExternalBankAccountRetryMicroDepositsResponseState = "ENABLED"
	ExternalBankAccountRetryMicroDepositsResponseStatePaused  ExternalBankAccountRetryMicroDepositsResponseState = "PAUSED"
)

func (r ExternalBankAccountRetryMicroDepositsResponseState) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryMicroDepositsResponseStateClosed, ExternalBankAccountRetryMicroDepositsResponseStateEnabled, ExternalBankAccountRetryMicroDepositsResponseStatePaused:
		return true
	}
	return false
}

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

type ExternalBankAccountRetryMicroDepositsResponseVerificationState string

const (
	ExternalBankAccountRetryMicroDepositsResponseVerificationStateEnabled            ExternalBankAccountRetryMicroDepositsResponseVerificationState = "ENABLED"
	ExternalBankAccountRetryMicroDepositsResponseVerificationStateFailedVerification ExternalBankAccountRetryMicroDepositsResponseVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountRetryMicroDepositsResponseVerificationStateInsufficientFunds  ExternalBankAccountRetryMicroDepositsResponseVerificationState = "INSUFFICIENT_FUNDS"
	ExternalBankAccountRetryMicroDepositsResponseVerificationStatePending            ExternalBankAccountRetryMicroDepositsResponseVerificationState = "PENDING"
)

func (r ExternalBankAccountRetryMicroDepositsResponseVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountRetryMicroDepositsResponseVerificationStateEnabled, ExternalBankAccountRetryMicroDepositsResponseVerificationStateFailedVerification, ExternalBankAccountRetryMicroDepositsResponseVerificationStateInsufficientFunds, ExternalBankAccountRetryMicroDepositsResponseVerificationStatePending:
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
	AccountNumber param.Field[string] `json:"account_number"`
	AccountToken  param.Field[string] `json:"account_token" format:"uuid"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address   param.Field[ExternalBankAccountAddressParam] `json:"address"`
	CompanyID param.Field[string]                          `json:"company_id"`
	Country   param.Field[string]                          `json:"country"`
	Currency  param.Field[string]                          `json:"currency"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             param.Field[time.Time] `json:"dob" format:"date"`
	DoingBusinessAs param.Field[string]    `json:"doing_business_as"`
	// The financial account token of the operating account, which will provide the
	// funds for micro deposits used to verify the account
	FinancialAccountToken param.Field[string]                               `json:"financial_account_token" format:"uuid"`
	Name                  param.Field[string]                               `json:"name"`
	Owner                 param.Field[string]                               `json:"owner,required"`
	OwnerType             param.Field[OwnerType]                            `json:"owner_type,required"`
	RoutingNumber         param.Field[string]                               `json:"routing_number"`
	Type                  param.Field[ExternalBankAccountNewParamsBodyType] `json:"type"`
	UserDefinedID         param.Field[string]                               `json:"user_defined_id"`
	// Indicates whether verification was enforced for a given association record. For
	// MICRO_DEPOSIT, option to disable verification if the external bank account has
	// already been verified before. By default, verification will be required unless
	// users pass in a value of false
	VerificationEnforcement param.Field[bool]               `json:"verification_enforcement"`
	VerificationMethod      param.Field[VerificationMethod] `json:"verification_method,required"`
	ProcessorToken          param.Field[string]             `json:"processor_token"`
}

func (r ExternalBankAccountNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalBankAccountNewParamsBody) implementsExternalBankAccountNewParamsBodyUnion() {}

// Satisfied by
// [ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest],
// [ExternalBankAccountNewParamsBodyPlaidCreateBankAccountAPIRequest],
// [ExternalBankAccountNewParamsBody].
type ExternalBankAccountNewParamsBodyUnion interface {
	implementsExternalBankAccountNewParamsBodyUnion()
}

type ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest struct {
	AccountNumber param.Field[string] `json:"account_number,required"`
	Country       param.Field[string] `json:"country,required"`
	Currency      param.Field[string] `json:"currency,required"`
	// The financial account token of the operating account, which will provide the
	// funds for micro deposits used to verify the account
	FinancialAccountToken param.Field[string]                                                                      `json:"financial_account_token,required" format:"uuid"`
	Owner                 param.Field[string]                                                                      `json:"owner,required"`
	OwnerType             param.Field[OwnerType]                                                                   `json:"owner_type,required"`
	RoutingNumber         param.Field[string]                                                                      `json:"routing_number,required"`
	Type                  param.Field[ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequestType] `json:"type,required"`
	VerificationMethod    param.Field[VerificationMethod]                                                          `json:"verification_method,required"`
	AccountToken          param.Field[string]                                                                      `json:"account_token" format:"uuid"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address   param.Field[ExternalBankAccountAddressParam] `json:"address"`
	CompanyID param.Field[string]                          `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             param.Field[time.Time] `json:"dob" format:"date"`
	DoingBusinessAs param.Field[string]    `json:"doing_business_as"`
	Name            param.Field[string]    `json:"name"`
	UserDefinedID   param.Field[string]    `json:"user_defined_id"`
	// Indicates whether verification was enforced for a given association record. For
	// MICRO_DEPOSIT, option to disable verification if the external bank account has
	// already been verified before. By default, verification will be required unless
	// users pass in a value of false
	VerificationEnforcement param.Field[bool] `json:"verification_enforcement"`
}

func (r ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalBankAccountNewParamsBodyBankVerifiedCreateBankAccountAPIRequest) implementsExternalBankAccountNewParamsBodyUnion() {
}

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
	Owner              param.Field[string]             `json:"owner,required"`
	OwnerType          param.Field[OwnerType]          `json:"owner_type,required"`
	ProcessorToken     param.Field[string]             `json:"processor_token,required"`
	VerificationMethod param.Field[VerificationMethod] `json:"verification_method,required"`
	AccountToken       param.Field[string]             `json:"account_token" format:"uuid"`
	CompanyID          param.Field[string]             `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             param.Field[time.Time] `json:"dob" format:"date"`
	DoingBusinessAs param.Field[string]    `json:"doing_business_as"`
	UserDefinedID   param.Field[string]    `json:"user_defined_id"`
}

func (r ExternalBankAccountNewParamsBodyPlaidCreateBankAccountAPIRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalBankAccountNewParamsBodyPlaidCreateBankAccountAPIRequest) implementsExternalBankAccountNewParamsBodyUnion() {
}

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
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address   param.Field[ExternalBankAccountAddressParam] `json:"address"`
	CompanyID param.Field[string]                          `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             param.Field[time.Time] `json:"dob" format:"date"`
	DoingBusinessAs param.Field[string]    `json:"doing_business_as"`
	Name            param.Field[string]    `json:"name"`
	Owner           param.Field[string]    `json:"owner"`
	OwnerType       param.Field[OwnerType] `json:"owner_type"`
	UserDefinedID   param.Field[string]    `json:"user_defined_id"`
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
	ExternalBankAccountListParamsStateClosed  ExternalBankAccountListParamsState = "CLOSED"
	ExternalBankAccountListParamsStateEnabled ExternalBankAccountListParamsState = "ENABLED"
	ExternalBankAccountListParamsStatePaused  ExternalBankAccountListParamsState = "PAUSED"
)

func (r ExternalBankAccountListParamsState) IsKnown() bool {
	switch r {
	case ExternalBankAccountListParamsStateClosed, ExternalBankAccountListParamsStateEnabled, ExternalBankAccountListParamsStatePaused:
		return true
	}
	return false
}

type ExternalBankAccountListParamsVerificationState string

const (
	ExternalBankAccountListParamsVerificationStateEnabled            ExternalBankAccountListParamsVerificationState = "ENABLED"
	ExternalBankAccountListParamsVerificationStateFailedVerification ExternalBankAccountListParamsVerificationState = "FAILED_VERIFICATION"
	ExternalBankAccountListParamsVerificationStateInsufficientFunds  ExternalBankAccountListParamsVerificationState = "INSUFFICIENT_FUNDS"
	ExternalBankAccountListParamsVerificationStatePending            ExternalBankAccountListParamsVerificationState = "PENDING"
)

func (r ExternalBankAccountListParamsVerificationState) IsKnown() bool {
	switch r {
	case ExternalBankAccountListParamsVerificationStateEnabled, ExternalBankAccountListParamsVerificationStateFailedVerification, ExternalBankAccountListParamsVerificationStateInsufficientFunds, ExternalBankAccountListParamsVerificationStatePending:
		return true
	}
	return false
}

type ExternalBankAccountRetryMicroDepositsParams struct {
	// The financial account token of the operating account, which will provide the
	// funds for micro deposits used to verify the account
	FinancialAccountToken param.Field[string] `json:"financial_account_token" format:"uuid"`
}

func (r ExternalBankAccountRetryMicroDepositsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
