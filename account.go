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

// AccountService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
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
func (r *AccountService) Get(ctx context.Context, accountToken string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	if accountToken == "" {
		err = errors.New("missing required account_token parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s", accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update account configuration such as spend limits and verification address. Can
// only be run on accounts that are part of the program managed by this API key.
//
// Accounts that are in the `PAUSED` state will not be able to transact or create
// new cards.
func (r *AccountService) Update(ctx context.Context, accountToken string, body AccountUpdateParams, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	if accountToken == "" {
		err = errors.New("missing required account_token parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s", accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List account configurations.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Account], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Account] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Get an Account's available spend limits, which is based on the spend limit
// configured on the Account and the amount already spent over the spend limit's
// duration. For example, if the Account has a daily spend limit of $1000
// configured, and has spent $600 in the last 24 hours, the available spend limit
// returned would be $400.
func (r *AccountService) GetSpendLimits(ctx context.Context, accountToken string, opts ...option.RequestOption) (res *AccountSpendLimits, err error) {
	opts = append(r.Options[:], opts...)
	if accountToken == "" {
		err = errors.New("missing required account_token parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/spend_limits", accountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Account struct {
	// Globally unique identifier for the account. This is the same as the
	// account_token returned by the enroll endpoint. If using this parameter, do not
	// include pagination.
	Token string `json:"token,required" format:"uuid"`
	// Timestamp of when the account was created. For accounts created before
	// 2023-05-11, this field will be null.
	Created time.Time `json:"created,required,nullable" format:"date-time"`
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
	//   - `CLOSED` - Account will not be able to transact or create new cards. `CLOSED`
	//     accounts are also unable to be transitioned to `ACTIVE` or `PAUSED` states.
	//     `CLOSED` accounts result from failing to pass KYB/KYC or Lithic closing for
	//     risk/compliance reasons. Please contact
	//     [support@lithic.com](mailto:support@lithic.com) if you believe this was in
	//     error.
	State         AccountState         `json:"state,required"`
	AccountHolder AccountAccountHolder `json:"account_holder"`
	// List of identifiers for the Auth Rule(s) that are applied on the account. This
	// field is deprecated and will no longer be populated in the `account_holder`
	// object. The key will be removed from the schema in a future release. Use the
	// `/auth_rules` endpoints to fetch Auth Rule information instead.
	AuthRuleTokens      []string                   `json:"auth_rule_tokens"`
	VerificationAddress AccountVerificationAddress `json:"verification_address"`
	JSON                accountJSON                `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	Token               apijson.Field
	Created             apijson.Field
	SpendLimit          apijson.Field
	State               apijson.Field
	AccountHolder       apijson.Field
	AuthRuleTokens      apijson.Field
	VerificationAddress apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountJSON) RawJSON() string {
	return r.raw
}

// Spend limit information for the user containing the daily, monthly, and lifetime
// spend limit of the account. Any charges to a card owned by this account will be
// declined once their transaction volume has surpassed the value in the applicable
// time limit (rolling). A lifetime limit of 0 indicates that the lifetime limit
// feature is disabled.
type AccountSpendLimit struct {
	// Daily spend limit (in cents).
	Daily int64 `json:"daily,required"`
	// Total spend limit over account lifetime (in cents).
	Lifetime int64 `json:"lifetime,required"`
	// Monthly spend limit (in cents).
	Monthly int64                 `json:"monthly,required"`
	JSON    accountSpendLimitJSON `json:"-"`
}

// accountSpendLimitJSON contains the JSON metadata for the struct
// [AccountSpendLimit]
type accountSpendLimitJSON struct {
	Daily       apijson.Field
	Lifetime    apijson.Field
	Monthly     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSpendLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSpendLimitJSON) RawJSON() string {
	return r.raw
}

// Account state:
//
//   - `ACTIVE` - Account is able to transact and create new cards.
//   - `PAUSED` - Account will not be able to transact or create new cards. It can be
//     set back to `ACTIVE`.
//   - `CLOSED` - Account will not be able to transact or create new cards. `CLOSED`
//     accounts are also unable to be transitioned to `ACTIVE` or `PAUSED` states.
//     `CLOSED` accounts result from failing to pass KYB/KYC or Lithic closing for
//     risk/compliance reasons. Please contact
//     [support@lithic.com](mailto:support@lithic.com) if you believe this was in
//     error.
type AccountState string

const (
	AccountStateActive AccountState = "ACTIVE"
	AccountStatePaused AccountState = "PAUSED"
	AccountStateClosed AccountState = "CLOSED"
)

func (r AccountState) IsKnown() bool {
	switch r {
	case AccountStateActive, AccountStatePaused, AccountStateClosed:
		return true
	}
	return false
}

type AccountAccountHolder struct {
	// Globally unique identifier for the account holder.
	Token string `json:"token,required"`
	// Only applicable for customers using the KYC-Exempt workflow to enroll authorized
	// users of businesses. Account_token of the enrolled business associated with an
	// enrolled AUTHORIZED_USER individual.
	BusinessAccountToken string `json:"business_account_token,required"`
	// Email address.
	Email string `json:"email,required"`
	// Phone number of the individual.
	PhoneNumber string                   `json:"phone_number,required"`
	JSON        accountAccountHolderJSON `json:"-"`
}

// accountAccountHolderJSON contains the JSON metadata for the struct
// [AccountAccountHolder]
type accountAccountHolderJSON struct {
	Token                apijson.Field
	BusinessAccountToken apijson.Field
	Email                apijson.Field
	PhoneNumber          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccountHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccountHolderJSON) RawJSON() string {
	return r.raw
}

type AccountVerificationAddress struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// City name.
	City string `json:"city,required"`
	// Country name. Only USA is currently supported.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 string                         `json:"address2"`
	JSON     accountVerificationAddressJSON `json:"-"`
}

// accountVerificationAddressJSON contains the JSON metadata for the struct
// [AccountVerificationAddress]
type accountVerificationAddressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVerificationAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVerificationAddressJSON) RawJSON() string {
	return r.raw
}

type AccountSpendLimits struct {
	AvailableSpendLimit AccountSpendLimitsAvailableSpendLimit `json:"available_spend_limit,required"`
	SpendLimit          AccountSpendLimitsSpendLimit          `json:"spend_limit"`
	SpendVelocity       AccountSpendLimitsSpendVelocity       `json:"spend_velocity"`
	JSON                accountSpendLimitsJSON                `json:"-"`
}

// accountSpendLimitsJSON contains the JSON metadata for the struct
// [AccountSpendLimits]
type accountSpendLimitsJSON struct {
	AvailableSpendLimit apijson.Field
	SpendLimit          apijson.Field
	SpendVelocity       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountSpendLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSpendLimitsJSON) RawJSON() string {
	return r.raw
}

type AccountSpendLimitsAvailableSpendLimit struct {
	// The available spend limit (in cents) relative to the daily limit configured on
	// the Account.
	Daily int64 `json:"daily"`
	// The available spend limit (in cents) relative to the lifetime limit configured
	// on the Account.
	Lifetime int64 `json:"lifetime"`
	// The available spend limit (in cents) relative to the monthly limit configured on
	// the Account.
	Monthly int64                                     `json:"monthly"`
	JSON    accountSpendLimitsAvailableSpendLimitJSON `json:"-"`
}

// accountSpendLimitsAvailableSpendLimitJSON contains the JSON metadata for the
// struct [AccountSpendLimitsAvailableSpendLimit]
type accountSpendLimitsAvailableSpendLimitJSON struct {
	Daily       apijson.Field
	Lifetime    apijson.Field
	Monthly     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSpendLimitsAvailableSpendLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSpendLimitsAvailableSpendLimitJSON) RawJSON() string {
	return r.raw
}

type AccountSpendLimitsSpendLimit struct {
	// The configured daily spend limit (in cents) on the Account.
	Daily int64 `json:"daily"`
	// The configured lifetime spend limit (in cents) on the Account.
	Lifetime int64 `json:"lifetime"`
	// The configured monthly spend limit (in cents) on the Account.
	Monthly int64                            `json:"monthly"`
	JSON    accountSpendLimitsSpendLimitJSON `json:"-"`
}

// accountSpendLimitsSpendLimitJSON contains the JSON metadata for the struct
// [AccountSpendLimitsSpendLimit]
type accountSpendLimitsSpendLimitJSON struct {
	Daily       apijson.Field
	Lifetime    apijson.Field
	Monthly     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSpendLimitsSpendLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSpendLimitsSpendLimitJSON) RawJSON() string {
	return r.raw
}

type AccountSpendLimitsSpendVelocity struct {
	// Current daily spend velocity (in cents) on the Account. Present if daily spend
	// limit is set.
	Daily int64 `json:"daily"`
	// Current lifetime spend velocity (in cents) on the Account. Present if lifetime
	// spend limit is set.
	Lifetime int64 `json:"lifetime"`
	// Current monthly spend velocity (in cents) on the Account. Present if monthly
	// spend limit is set.
	Monthly int64                               `json:"monthly"`
	JSON    accountSpendLimitsSpendVelocityJSON `json:"-"`
}

// accountSpendLimitsSpendVelocityJSON contains the JSON metadata for the struct
// [AccountSpendLimitsSpendVelocity]
type accountSpendLimitsSpendVelocityJSON struct {
	Daily       apijson.Field
	Lifetime    apijson.Field
	Monthly     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSpendLimitsSpendVelocity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSpendLimitsSpendVelocityJSON) RawJSON() string {
	return r.raw
}

type AccountUpdateParams struct {
	// Amount (in cents) for the account's daily spend limit. By default the daily
	// spend limit is set to $1,250.
	DailySpendLimit param.Field[int64] `json:"daily_spend_limit"`
	// Amount (in cents) for the account's lifetime spend limit. Once this limit is
	// reached, no transactions will be accepted on any card created for this account
	// until the limit is updated. Note that a spend limit of 0 is effectively no
	// limit, and should only be used to reset or remove a prior limit. Only a limit of
	// 1 or above will result in declined transactions due to checks against the
	// account limit. This behavior differs from the daily spend limit and the monthly
	// spend limit.
	LifetimeSpendLimit param.Field[int64] `json:"lifetime_spend_limit"`
	// Amount (in cents) for the account's monthly spend limit. By default the monthly
	// spend limit is set to $5,000.
	MonthlySpendLimit param.Field[int64] `json:"monthly_spend_limit"`
	// Account states.
	State param.Field[AccountUpdateParamsState] `json:"state"`
	// Address used during Address Verification Service (AVS) checks during
	// transactions if enabled via Auth Rules.
	VerificationAddress param.Field[AccountUpdateParamsVerificationAddress] `json:"verification_address"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account states.
type AccountUpdateParamsState string

const (
	AccountUpdateParamsStateActive AccountUpdateParamsState = "ACTIVE"
	AccountUpdateParamsStatePaused AccountUpdateParamsState = "PAUSED"
)

func (r AccountUpdateParamsState) IsKnown() bool {
	switch r {
	case AccountUpdateParamsStateActive, AccountUpdateParamsStatePaused:
		return true
	}
	return false
}

// Address used during Address Verification Service (AVS) checks during
// transactions if enabled via Auth Rules.
type AccountUpdateParamsVerificationAddress struct {
	Address1   param.Field[string] `json:"address1"`
	Address2   param.Field[string] `json:"address2"`
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r AccountUpdateParamsVerificationAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
