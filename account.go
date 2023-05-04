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

// AccountService contains methods and other services that help with interacting
// with the lithic API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Get account configuration such as spend limits.
func (r *AccountService) Get(ctx context.Context, account_token string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update account configuration such as spend limits and verification address. Can
// only be run on accounts that are part of the program managed by this API key.
//
// Accounts that are in the `PAUSED` state will not be able to transact or create
// new cards.
func (r *AccountService) Update(ctx context.Context, account_token string, body AccountUpdateParams, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List account configurations.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *shared.Page[Account], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
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

// List account configurations.
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[Account] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Account struct {
	// Spend limit information for the user containing the daily, monthly, and lifetime
	// spend limit of the account. Any charges to a card owned by this account will be
	// declined once their transaction volume has surpassed the value in the applicable
	// time limit (rolling). A lifetime limit of 0 indicates that the lifetime limit
	// feature is disabled.
	SpendLimit AccountSpendLimit `json:"spend_limit,required"`
	// Account state:
	//
	//   - `ACTIVE` - Account is able to transact and create new cards.
	//   - `PAUSED` - Account will not be able to transact or create new cards. It can be
	//     set back to `ACTIVE`.
	//   - `CLOSED` - Account will permanently not be able to transact or create new
	//     cards.
	State AccountState `json:"state,required"`
	// Globally unique identifier for the account. This is the same as the
	// account_token returned by the enroll endpoint. If using this parameter, do not
	// include pagination.
	Token string `json:"token,required" format:"uuid"`
	// List of identifiers for the Auth Rule(s) that are applied on the account.
	AuthRuleTokens      []string                   `json:"auth_rule_tokens"`
	VerificationAddress AccountVerificationAddress `json:"verification_address"`
	AccountHolder       AccountAccountHolder       `json:"account_holder"`
	JSON                accountJSON
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	SpendLimit          apijson.Field
	State               apijson.Field
	Token               apijson.Field
	AuthRuleTokens      apijson.Field
	VerificationAddress apijson.Field
	AccountHolder       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Spend limit information for the user containing the daily, monthly, and lifetime
// spend limit of the account. Any charges to a card owned by this account will be
// declined once their transaction volume has surpassed the value in the applicable
// time limit (rolling). A lifetime limit of 0 indicates that the lifetime limit
// feature is disabled.
type AccountSpendLimit struct {
	// Daily spend limit (in cents).
	Daily int64 `json:"daily,required"`
	// Monthly spend limit (in cents).
	Monthly int64 `json:"monthly,required"`
	// Total spend limit over account lifetime (in cents).
	Lifetime int64 `json:"lifetime,required"`
	JSON     accountSpendLimitJSON
}

// accountSpendLimitJSON contains the JSON metadata for the struct
// [AccountSpendLimit]
type accountSpendLimitJSON struct {
	Daily       apijson.Field
	Monthly     apijson.Field
	Lifetime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSpendLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountState string

const (
	AccountStateActive AccountState = "ACTIVE"
	AccountStatePaused AccountState = "PAUSED"
	AccountStateClosed AccountState = "CLOSED"
)

type AccountVerificationAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Unit or apartment number (if applicable).
	Address2 string `json:"address2"`
	// City name.
	City string `json:"city,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Country name. Only USA is currently supported.
	Country string `json:"country,required"`
	JSON    accountVerificationAddressJSON
}

// accountVerificationAddressJSON contains the JSON metadata for the struct
// [AccountVerificationAddress]
type accountVerificationAddressJSON struct {
	Address1    apijson.Field
	Address2    apijson.Field
	City        apijson.Field
	State       apijson.Field
	PostalCode  apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVerificationAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccountHolder struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token,required"`
	// Phone number of the individual.
	PhoneNumber string `json:"phone_number,required"`
	// Email address.
	Email string `json:"email,required"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Account_token of the enrolled business associated with an
	// enrolled AUTHORIZED_USER individual.
	BusinessAccountToken string `json:"business_account_token,required"`
	JSON                 accountAccountHolderJSON
}

// accountAccountHolderJSON contains the JSON metadata for the struct
// [AccountAccountHolder]
type accountAccountHolderJSON struct {
	Token                apijson.Field
	PhoneNumber          apijson.Field
	Email                apijson.Field
	BusinessAccountToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccountHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateParams struct {
	// Amount (in cents) for the account's new daily spend limit. Note that a spend
	// limit of 0 is effectively no limit, and should only be used to reset or remove a
	// prior limit. Only a limit of 1 or above will result in declined transactions due
	// to checks against the account limit.
	DailySpendLimit param.Field[int64] `json:"daily_spend_limit"`
	// Amount (in cents) for the account's new lifetime limit. Once this limit is
	// reached, no transactions will be accepted on any card created for this account
	// until the limit is updated. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the
	// account limit.
	LifetimeSpendLimit param.Field[int64] `json:"lifetime_spend_limit"`
	// Amount (in cents) for the account's new monthly spend limit. Note that a spend
	// limit of 0 is effectively no limit, and should only be used to reset or remove a
	// prior limit. Only a limit of 1 or above will result in declined transactions due
	// to checks against the account limit.
	MonthlySpendLimit param.Field[int64] `json:"monthly_spend_limit"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	VerificationAddress param.Field[AccountUpdateParamsVerificationAddress] `json:"verification_address"`
	// Account states.
	State param.Field[AccountUpdateParamsState] `json:"state"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Address used during Address Verification Service (AVS) checks during
// transactions if enabled via Auth Rules.
type AccountUpdateParamsVerificationAddress struct {
	Address1   param.Field[string] `json:"address1"`
	Address2   param.Field[string] `json:"address2"`
	City       param.Field[string] `json:"city"`
	State      param.Field[string] `json:"state"`
	PostalCode param.Field[string] `json:"postal_code"`
	Country    param.Field[string] `json:"country"`
}

type AccountUpdateParamsState string

const (
	AccountUpdateParamsStateActive AccountUpdateParamsState = "ACTIVE"
	AccountUpdateParamsStatePaused AccountUpdateParamsState = "PAUSED"
)

type AccountListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified date
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified date
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Page (for pagination).
	Page param.Field[int64] `query:"page"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type AccountListResponse struct {
	Data []Account `json:"data,required"`
	// Page number.
	Page int64 `json:"page,required"`
	// Total number of entries.
	TotalEntries int64 `json:"total_entries,required"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages,required"`
	JSON       accountListResponseJSON
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	Data         apijson.Field
	Page         apijson.Field
	TotalEntries apijson.Field
	TotalPages   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
