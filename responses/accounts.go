package responses

import (
	apijson "github.com/lithic-com/lithic-go/core/json"
)

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
	JSON                AccountJSON
}

type AccountJSON struct {
	SpendLimit          apijson.Metadata
	State               apijson.Metadata
	Token               apijson.Metadata
	AuthRuleTokens      apijson.Metadata
	VerificationAddress apijson.Metadata
	AccountHolder       apijson.Metadata
	Raw                 []byte
	Extras              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Account using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSpendLimit struct {
	// Daily spend limit (in cents).
	Daily int64 `json:"daily,required"`
	// Monthly spend limit (in cents).
	Monthly int64 `json:"monthly,required"`
	// Total spend limit over account lifetime (in cents).
	Lifetime int64 `json:"lifetime,required"`
	JSON     AccountSpendLimitJSON
}

type AccountSpendLimitJSON struct {
	Daily    apijson.Metadata
	Monthly  apijson.Metadata
	Lifetime apijson.Metadata
	Raw      []byte
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountSpendLimit using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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
	JSON    AccountVerificationAddressJSON
}

type AccountVerificationAddressJSON struct {
	Address1   apijson.Metadata
	Address2   apijson.Metadata
	City       apijson.Metadata
	State      apijson.Metadata
	PostalCode apijson.Metadata
	Country    apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountVerificationAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
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
	JSON                 AccountAccountHolderJSON
}

type AccountAccountHolderJSON struct {
	Token                apijson.Metadata
	PhoneNumber          apijson.Metadata
	Email                apijson.Metadata
	BusinessAccountToken apijson.Metadata
	Raw                  []byte
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountAccountHolder using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountAccountHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponse struct {
	Data []Account `json:"data,required"`
	// Page number.
	Page int64 `json:"page,required"`
	// Total number of entries.
	TotalEntries int64 `json:"total_entries,required"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages,required"`
	JSON       AccountListResponseJSON
}

type AccountListResponseJSON struct {
	Data         apijson.Metadata
	Page         apijson.Metadata
	TotalEntries apijson.Metadata
	TotalPages   apijson.Metadata
	Raw          []byte
	Extras       map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
