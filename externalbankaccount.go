// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/internal/shared"
	"github.com/lithic-com/lithic-go/option"
)

// ExternalBankAccountService contains methods and other services that help with
// interacting with the lithic API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewExternalBankAccountService]
// method instead.
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
func (r *ExternalBankAccountService) List(ctx context.Context, query ExternalBankAccountListParams, opts ...option.RequestOption) (res *shared.CursorPage[ExternalBankAccountListResponse], err error) {
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
func (r *ExternalBankAccountService) ListAutoPaging(ctx context.Context, query ExternalBankAccountListParams, opts ...option.RequestOption) *shared.CursorPageAutoPager[ExternalBankAccountListResponse] {
	return shared.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Address used during Address Verification Service (AVS) checks during
// transactions if enabled via Auth Rules.
type ExternalBankAccountAddress struct {
	Address1   string `json:"address1,required"`
	City       string `json:"city,required"`
	Country    string `json:"country,required"`
	PostalCode string `json:"postal_code,required"`
	State      string `json:"state,required"`
	Address2   string `json:"address2"`
	JSON       externalBankAccountAddressJSON
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
	OwnerTypeIndividual OwnerType = "INDIVIDUAL"
	OwnerTypeBusiness   OwnerType = "BUSINESS"
)

type VerificationMethod string

const (
	VerificationMethodManual       VerificationMethod = "MANUAL"
	VerificationMethodMicroDeposit VerificationMethod = "MICRO_DEPOSIT"
	VerificationMethodPlaid        VerificationMethod = "PLAID"
)

type ExternalBankAccountNewResponse struct {
	// A globally unique identifier for this record of an external bank account
	// association. If a program links an external bank account to more than one
	// end-user or to both the program and the end-user, then Lithic will return each
	// record of the association
	Token string `json:"token,required" format:"uuid"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept US bank accounts e.g., US
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
	Owner              string                                           `json:"owner,required"`
	OwnerType          ExternalBankAccountNewResponseOwnerType          `json:"owner_type,required"`
	RoutingNumber      string                                           `json:"routing_number,required"`
	State              ExternalBankAccountNewResponseState              `json:"state,required"`
	Type               ExternalBankAccountNewResponseType               `json:"type,required"`
	VerificationMethod ExternalBankAccountNewResponseVerificationMethod `json:"verification_method,required"`
	VerificationState  ExternalBankAccountNewResponseVerificationState  `json:"verification_state,required"`
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
	// The nickname given to this record of External Bank Account
	Name string `json:"name"`
	JSON externalBankAccountNewResponseJSON
}

// externalBankAccountNewResponseJSON contains the JSON metadata for the struct
// [ExternalBankAccountNewResponse]
type externalBankAccountNewResponseJSON struct {
	Token              apijson.Field
	Country            apijson.Field
	Created            apijson.Field
	Currency           apijson.Field
	LastFour           apijson.Field
	Owner              apijson.Field
	OwnerType          apijson.Field
	RoutingNumber      apijson.Field
	State              apijson.Field
	Type               apijson.Field
	VerificationMethod apijson.Field
	VerificationState  apijson.Field
	AccountToken       apijson.Field
	Address            apijson.Field
	CompanyID          apijson.Field
	Dob                apijson.Field
	DoingBusinessAs    apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExternalBankAccountNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBankAccountNewResponseOwnerType string

const (
	ExternalBankAccountNewResponseOwnerTypeIndividual ExternalBankAccountNewResponseOwnerType = "INDIVIDUAL"
	ExternalBankAccountNewResponseOwnerTypeBusiness   ExternalBankAccountNewResponseOwnerType = "BUSINESS"
)

type ExternalBankAccountNewResponseState string

const (
	ExternalBankAccountNewResponseStateEnabled ExternalBankAccountNewResponseState = "ENABLED"
	ExternalBankAccountNewResponseStateClosed  ExternalBankAccountNewResponseState = "CLOSED"
	ExternalBankAccountNewResponseStatePaused  ExternalBankAccountNewResponseState = "PAUSED"
)

type ExternalBankAccountNewResponseType string

const (
	ExternalBankAccountNewResponseTypeChecking ExternalBankAccountNewResponseType = "CHECKING"
	ExternalBankAccountNewResponseTypeSavings  ExternalBankAccountNewResponseType = "SAVINGS"
)

type ExternalBankAccountNewResponseVerificationMethod string

const (
	ExternalBankAccountNewResponseVerificationMethodManual       ExternalBankAccountNewResponseVerificationMethod = "MANUAL"
	ExternalBankAccountNewResponseVerificationMethodMicroDeposit ExternalBankAccountNewResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountNewResponseVerificationMethodPlaid        ExternalBankAccountNewResponseVerificationMethod = "PLAID"
)

type ExternalBankAccountNewResponseVerificationState string

const (
	ExternalBankAccountNewResponseVerificationStatePending            ExternalBankAccountNewResponseVerificationState = "PENDING"
	ExternalBankAccountNewResponseVerificationStateEnabled            ExternalBankAccountNewResponseVerificationState = "ENABLED"
	ExternalBankAccountNewResponseVerificationStateFailedVerification ExternalBankAccountNewResponseVerificationState = "FAILED_VERIFICATION"
)

type ExternalBankAccountGetResponse struct {
	// A globally unique identifier for this record of an external bank account
	// association. If a program links an external bank account to more than one
	// end-user or to both the program and the end-user, then Lithic will return each
	// record of the association
	Token string `json:"token,required" format:"uuid"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept US bank accounts e.g., US
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
	Owner              string                                           `json:"owner,required"`
	OwnerType          ExternalBankAccountGetResponseOwnerType          `json:"owner_type,required"`
	RoutingNumber      string                                           `json:"routing_number,required"`
	State              ExternalBankAccountGetResponseState              `json:"state,required"`
	Type               ExternalBankAccountGetResponseType               `json:"type,required"`
	VerificationMethod ExternalBankAccountGetResponseVerificationMethod `json:"verification_method,required"`
	VerificationState  ExternalBankAccountGetResponseVerificationState  `json:"verification_state,required"`
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
	// The nickname given to this record of External Bank Account
	Name string `json:"name"`
	JSON externalBankAccountGetResponseJSON
}

// externalBankAccountGetResponseJSON contains the JSON metadata for the struct
// [ExternalBankAccountGetResponse]
type externalBankAccountGetResponseJSON struct {
	Token              apijson.Field
	Country            apijson.Field
	Created            apijson.Field
	Currency           apijson.Field
	LastFour           apijson.Field
	Owner              apijson.Field
	OwnerType          apijson.Field
	RoutingNumber      apijson.Field
	State              apijson.Field
	Type               apijson.Field
	VerificationMethod apijson.Field
	VerificationState  apijson.Field
	AccountToken       apijson.Field
	Address            apijson.Field
	CompanyID          apijson.Field
	Dob                apijson.Field
	DoingBusinessAs    apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExternalBankAccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBankAccountGetResponseOwnerType string

const (
	ExternalBankAccountGetResponseOwnerTypeIndividual ExternalBankAccountGetResponseOwnerType = "INDIVIDUAL"
	ExternalBankAccountGetResponseOwnerTypeBusiness   ExternalBankAccountGetResponseOwnerType = "BUSINESS"
)

type ExternalBankAccountGetResponseState string

const (
	ExternalBankAccountGetResponseStateEnabled ExternalBankAccountGetResponseState = "ENABLED"
	ExternalBankAccountGetResponseStateClosed  ExternalBankAccountGetResponseState = "CLOSED"
	ExternalBankAccountGetResponseStatePaused  ExternalBankAccountGetResponseState = "PAUSED"
)

type ExternalBankAccountGetResponseType string

const (
	ExternalBankAccountGetResponseTypeChecking ExternalBankAccountGetResponseType = "CHECKING"
	ExternalBankAccountGetResponseTypeSavings  ExternalBankAccountGetResponseType = "SAVINGS"
)

type ExternalBankAccountGetResponseVerificationMethod string

const (
	ExternalBankAccountGetResponseVerificationMethodManual       ExternalBankAccountGetResponseVerificationMethod = "MANUAL"
	ExternalBankAccountGetResponseVerificationMethodMicroDeposit ExternalBankAccountGetResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountGetResponseVerificationMethodPlaid        ExternalBankAccountGetResponseVerificationMethod = "PLAID"
)

type ExternalBankAccountGetResponseVerificationState string

const (
	ExternalBankAccountGetResponseVerificationStatePending            ExternalBankAccountGetResponseVerificationState = "PENDING"
	ExternalBankAccountGetResponseVerificationStateEnabled            ExternalBankAccountGetResponseVerificationState = "ENABLED"
	ExternalBankAccountGetResponseVerificationStateFailedVerification ExternalBankAccountGetResponseVerificationState = "FAILED_VERIFICATION"
)

type ExternalBankAccountUpdateResponse struct {
	// A globally unique identifier for this record of an external bank account
	// association. If a program links an external bank account to more than one
	// end-user or to both the program and the end-user, then Lithic will return each
	// record of the association
	Token string `json:"token,required" format:"uuid"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept US bank accounts e.g., US
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
	Owner              string                                              `json:"owner,required"`
	OwnerType          ExternalBankAccountUpdateResponseOwnerType          `json:"owner_type,required"`
	RoutingNumber      string                                              `json:"routing_number,required"`
	State              ExternalBankAccountUpdateResponseState              `json:"state,required"`
	Type               ExternalBankAccountUpdateResponseType               `json:"type,required"`
	VerificationMethod ExternalBankAccountUpdateResponseVerificationMethod `json:"verification_method,required"`
	VerificationState  ExternalBankAccountUpdateResponseVerificationState  `json:"verification_state,required"`
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
	// The nickname given to this record of External Bank Account
	Name string `json:"name"`
	JSON externalBankAccountUpdateResponseJSON
}

// externalBankAccountUpdateResponseJSON contains the JSON metadata for the struct
// [ExternalBankAccountUpdateResponse]
type externalBankAccountUpdateResponseJSON struct {
	Token              apijson.Field
	Country            apijson.Field
	Created            apijson.Field
	Currency           apijson.Field
	LastFour           apijson.Field
	Owner              apijson.Field
	OwnerType          apijson.Field
	RoutingNumber      apijson.Field
	State              apijson.Field
	Type               apijson.Field
	VerificationMethod apijson.Field
	VerificationState  apijson.Field
	AccountToken       apijson.Field
	Address            apijson.Field
	CompanyID          apijson.Field
	Dob                apijson.Field
	DoingBusinessAs    apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExternalBankAccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBankAccountUpdateResponseOwnerType string

const (
	ExternalBankAccountUpdateResponseOwnerTypeIndividual ExternalBankAccountUpdateResponseOwnerType = "INDIVIDUAL"
	ExternalBankAccountUpdateResponseOwnerTypeBusiness   ExternalBankAccountUpdateResponseOwnerType = "BUSINESS"
)

type ExternalBankAccountUpdateResponseState string

const (
	ExternalBankAccountUpdateResponseStateEnabled ExternalBankAccountUpdateResponseState = "ENABLED"
	ExternalBankAccountUpdateResponseStateClosed  ExternalBankAccountUpdateResponseState = "CLOSED"
	ExternalBankAccountUpdateResponseStatePaused  ExternalBankAccountUpdateResponseState = "PAUSED"
)

type ExternalBankAccountUpdateResponseType string

const (
	ExternalBankAccountUpdateResponseTypeChecking ExternalBankAccountUpdateResponseType = "CHECKING"
	ExternalBankAccountUpdateResponseTypeSavings  ExternalBankAccountUpdateResponseType = "SAVINGS"
)

type ExternalBankAccountUpdateResponseVerificationMethod string

const (
	ExternalBankAccountUpdateResponseVerificationMethodManual       ExternalBankAccountUpdateResponseVerificationMethod = "MANUAL"
	ExternalBankAccountUpdateResponseVerificationMethodMicroDeposit ExternalBankAccountUpdateResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountUpdateResponseVerificationMethodPlaid        ExternalBankAccountUpdateResponseVerificationMethod = "PLAID"
)

type ExternalBankAccountUpdateResponseVerificationState string

const (
	ExternalBankAccountUpdateResponseVerificationStatePending            ExternalBankAccountUpdateResponseVerificationState = "PENDING"
	ExternalBankAccountUpdateResponseVerificationStateEnabled            ExternalBankAccountUpdateResponseVerificationState = "ENABLED"
	ExternalBankAccountUpdateResponseVerificationStateFailedVerification ExternalBankAccountUpdateResponseVerificationState = "FAILED_VERIFICATION"
)

type ExternalBankAccountListResponse struct {
	// A globally unique identifier for this record of an external bank account
	// association. If a program links an external bank account to more than one
	// end-user or to both the program and the end-user, then Lithic will return each
	// record of the association
	Token string `json:"token,required" format:"uuid"`
	// The country that the bank account is located in using ISO 3166-1. We will only
	// accept US bank accounts e.g., US
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
	Owner              string                                            `json:"owner,required"`
	OwnerType          ExternalBankAccountListResponseOwnerType          `json:"owner_type,required"`
	RoutingNumber      string                                            `json:"routing_number,required"`
	State              ExternalBankAccountListResponseState              `json:"state,required"`
	Type               ExternalBankAccountListResponseType               `json:"type,required"`
	VerificationMethod ExternalBankAccountListResponseVerificationMethod `json:"verification_method,required"`
	VerificationState  ExternalBankAccountListResponseVerificationState  `json:"verification_state,required"`
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
	// The nickname given to this record of External Bank Account
	Name string `json:"name"`
	JSON externalBankAccountListResponseJSON
}

// externalBankAccountListResponseJSON contains the JSON metadata for the struct
// [ExternalBankAccountListResponse]
type externalBankAccountListResponseJSON struct {
	Token              apijson.Field
	Country            apijson.Field
	Created            apijson.Field
	Currency           apijson.Field
	LastFour           apijson.Field
	Owner              apijson.Field
	OwnerType          apijson.Field
	RoutingNumber      apijson.Field
	State              apijson.Field
	Type               apijson.Field
	VerificationMethod apijson.Field
	VerificationState  apijson.Field
	AccountToken       apijson.Field
	Address            apijson.Field
	CompanyID          apijson.Field
	Dob                apijson.Field
	DoingBusinessAs    apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExternalBankAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalBankAccountListResponseOwnerType string

const (
	ExternalBankAccountListResponseOwnerTypeIndividual ExternalBankAccountListResponseOwnerType = "INDIVIDUAL"
	ExternalBankAccountListResponseOwnerTypeBusiness   ExternalBankAccountListResponseOwnerType = "BUSINESS"
)

type ExternalBankAccountListResponseState string

const (
	ExternalBankAccountListResponseStateEnabled ExternalBankAccountListResponseState = "ENABLED"
	ExternalBankAccountListResponseStateClosed  ExternalBankAccountListResponseState = "CLOSED"
	ExternalBankAccountListResponseStatePaused  ExternalBankAccountListResponseState = "PAUSED"
)

type ExternalBankAccountListResponseType string

const (
	ExternalBankAccountListResponseTypeChecking ExternalBankAccountListResponseType = "CHECKING"
	ExternalBankAccountListResponseTypeSavings  ExternalBankAccountListResponseType = "SAVINGS"
)

type ExternalBankAccountListResponseVerificationMethod string

const (
	ExternalBankAccountListResponseVerificationMethodManual       ExternalBankAccountListResponseVerificationMethod = "MANUAL"
	ExternalBankAccountListResponseVerificationMethodMicroDeposit ExternalBankAccountListResponseVerificationMethod = "MICRO_DEPOSIT"
	ExternalBankAccountListResponseVerificationMethodPlaid        ExternalBankAccountListResponseVerificationMethod = "PLAID"
)

type ExternalBankAccountListResponseVerificationState string

const (
	ExternalBankAccountListResponseVerificationStatePending            ExternalBankAccountListResponseVerificationState = "PENDING"
	ExternalBankAccountListResponseVerificationStateEnabled            ExternalBankAccountListResponseVerificationState = "ENABLED"
	ExternalBankAccountListResponseVerificationStateFailedVerification ExternalBankAccountListResponseVerificationState = "FAILED_VERIFICATION"
)

// This interface is a union satisfied by one of the following:
// [ExternalBankAccountNewParamsPlaidCreateBankAccountAPIRequest],
// [ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequest].
type ExternalBankAccountNewParams interface {
	ImplementsExternalBankAccountNewParams()
}

type ExternalBankAccountNewParamsPlaidCreateBankAccountAPIRequest struct {
	Owner              param.Field[string]             `json:"owner,required"`
	OwnerType          param.Field[OwnerType]          `json:"owner_type,required"`
	ProcessorToken     param.Field[string]             `json:"processor_token,required"`
	VerificationMethod param.Field[VerificationMethod] `json:"verification_method,required"`
	AccountToken       param.Field[string]             `json:"account_token" format:"uuid"`
	CompanyID          param.Field[string]             `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             param.Field[time.Time] `json:"dob" format:"date"`
	DoingBusinessAs param.Field[string]    `json:"doing_business_as"`
}

func (r ExternalBankAccountNewParamsPlaidCreateBankAccountAPIRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ExternalBankAccountNewParamsPlaidCreateBankAccountAPIRequest) ImplementsExternalBankAccountNewParams() {

}

type ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequest struct {
	AccountNumber      param.Field[string]                                                                  `json:"account_number,required"`
	Country            param.Field[string]                                                                  `json:"country,required"`
	Currency           param.Field[string]                                                                  `json:"currency,required"`
	Owner              param.Field[string]                                                                  `json:"owner,required"`
	OwnerType          param.Field[OwnerType]                                                               `json:"owner_type,required"`
	RoutingNumber      param.Field[string]                                                                  `json:"routing_number,required"`
	Type               param.Field[ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequestType] `json:"type,required"`
	VerificationMethod param.Field[VerificationMethod]                                                      `json:"verification_method,required"`
	AccountToken       param.Field[string]                                                                  `json:"account_token" format:"uuid"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	Address   param.Field[ExternalBankAccountAddressParam] `json:"address"`
	CompanyID param.Field[string]                          `json:"company_id"`
	// Date of Birth of the Individual that owns the external bank account
	Dob             param.Field[time.Time] `json:"dob" format:"date"`
	DoingBusinessAs param.Field[string]    `json:"doing_business_as"`
	Name            param.Field[string]    `json:"name"`
	// Indicates whether verification was enforced for a given association record. For
	// MICRO_DEPOSIT, option to disable verification if the external bank account has
	// already been verified before. By default, verification will be required unless
	// users pass in a value of false
	VerificationEnforcement param.Field[bool] `json:"verification_enforcement"`
}

func (r ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequest) ImplementsExternalBankAccountNewParams() {

}

type ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequestType string

const (
	ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequestTypeChecking ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequestType = "CHECKING"
	ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequestTypeSavings  ExternalBankAccountNewParamsBankVerifiedCreateBankAccountAPIRequestType = "SAVINGS"
)

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

type ExternalBankAccountListParamsState string

const (
	ExternalBankAccountListParamsStateEnabled ExternalBankAccountListParamsState = "ENABLED"
	ExternalBankAccountListParamsStateClosed  ExternalBankAccountListParamsState = "CLOSED"
	ExternalBankAccountListParamsStatePaused  ExternalBankAccountListParamsState = "PAUSED"
)

type ExternalBankAccountListParamsVerificationState string

const (
	ExternalBankAccountListParamsVerificationStatePending            ExternalBankAccountListParamsVerificationState = "PENDING"
	ExternalBankAccountListParamsVerificationStateEnabled            ExternalBankAccountListParamsVerificationState = "ENABLED"
	ExternalBankAccountListParamsVerificationStateFailedVerification ExternalBankAccountListParamsVerificationState = "FAILED_VERIFICATION"
)
