// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// BlockchainRecipientService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainRecipientService] method instead.
type BlockchainRecipientService struct {
	Options []option.RequestOption
}

// NewBlockchainRecipientService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainRecipientService(opts ...option.RequestOption) (r *BlockchainRecipientService) {
	r = &BlockchainRecipientService{}
	r.Options = opts
	return
}

// Register a blockchain address as a withdrawal destination for a financial
// account
//
// The recipient is created with a `PENDING` verification state and cannot receive
// a payout until screening of the address completes. Registering an address that
// is already registered to the same financial account returns the existing
// recipient and its current verification state, rather than creating a second one
func (r *BlockchainRecipientService) New(ctx context.Context, body BlockchainRecipientNewParams, opts ...option.RequestOption) (res *BlockchainRecipient, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/blockchain_recipients"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BlockchainRecipient struct {
	// A globally unique identifier for this blockchain recipient
	Token string `json:"token" api:"required" format:"uuid"`
	// The financial account the blockchain recipient belongs to, or null when the
	// recipient is registered against the program rather than a financial account
	AccountToken string `json:"account_token" api:"required,nullable" format:"uuid"`
	// An optional tag or memo used by some chains to identify the destination of a
	// transfer within a shared address
	AddressTag string `json:"address_tag" api:"required,nullable"`
	// The blockchain network that the address belongs to
	Chain string `json:"chain" api:"required"`
	// An ISO 8601 string representing when this blockchain recipient was created
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The identifier the recipient is registered under with the payment provider
	ExternalID string `json:"external_id" api:"required,nullable"`
	// The nickname for this blockchain recipient
	Name string `json:"name" api:"required,nullable"`
	// Legal name of the business or individual who owns the blockchain address
	Owner string `json:"owner" api:"required"`
	// Owner Type
	OwnerType OwnerType `json:"owner_type" api:"required"`
	// Globally unique identifier for the program the blockchain recipient is
	// associated with
	ProgramID string `json:"program_id" api:"required" format:"uuid"`
	// Account State
	State BlockchainRecipientState `json:"state" api:"required"`
	// An ISO 8601 string representing when this blockchain recipient was last updated
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Verification State
	VerificationState BlockchainRecipientVerificationState `json:"verification_state" api:"required"`
	JSON              blockchainRecipientJSON              `json:"-"`
}

// blockchainRecipientJSON contains the JSON metadata for the struct
// [BlockchainRecipient]
type blockchainRecipientJSON struct {
	Token             apijson.Field
	AccountToken      apijson.Field
	AddressTag        apijson.Field
	Chain             apijson.Field
	Created           apijson.Field
	ExternalID        apijson.Field
	Name              apijson.Field
	Owner             apijson.Field
	OwnerType         apijson.Field
	ProgramID         apijson.Field
	State             apijson.Field
	Updated           apijson.Field
	VerificationState apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BlockchainRecipient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainRecipientJSON) RawJSON() string {
	return r.raw
}

// Account State
type BlockchainRecipientState string

const (
	BlockchainRecipientStateEnabled BlockchainRecipientState = "ENABLED"
	BlockchainRecipientStateClosed  BlockchainRecipientState = "CLOSED"
	BlockchainRecipientStatePaused  BlockchainRecipientState = "PAUSED"
)

func (r BlockchainRecipientState) IsKnown() bool {
	switch r {
	case BlockchainRecipientStateEnabled, BlockchainRecipientStateClosed, BlockchainRecipientStatePaused:
		return true
	}
	return false
}

// Verification State
type BlockchainRecipientVerificationState string

const (
	BlockchainRecipientVerificationStatePending            BlockchainRecipientVerificationState = "PENDING"
	BlockchainRecipientVerificationStateEnabled            BlockchainRecipientVerificationState = "ENABLED"
	BlockchainRecipientVerificationStateFailedVerification BlockchainRecipientVerificationState = "FAILED_VERIFICATION"
	BlockchainRecipientVerificationStateInsufficientFunds  BlockchainRecipientVerificationState = "INSUFFICIENT_FUNDS"
)

func (r BlockchainRecipientVerificationState) IsKnown() bool {
	switch r {
	case BlockchainRecipientVerificationStatePending, BlockchainRecipientVerificationStateEnabled, BlockchainRecipientVerificationStateFailedVerification, BlockchainRecipientVerificationStateInsufficientFunds:
		return true
	}
	return false
}

type BlockchainRecipientNewParams struct {
	// The financial account the blockchain recipient belongs to
	AccountToken param.Field[string] `json:"account_token" api:"required" format:"uuid"`
	// The blockchain address funds will be withdrawn to
	Address param.Field[string] `json:"address" api:"required"`
	// The blockchain network that the address belongs to
	Chain param.Field[string] `json:"chain" api:"required"`
	// Legal name of the business or individual who owns the blockchain address
	Owner param.Field[string] `json:"owner" api:"required"`
	// Owner Type
	OwnerType param.Field[OwnerType] `json:"owner_type" api:"required"`
	// An optional tag or memo used by some chains to identify the destination of a
	// transfer within a shared address
	AddressTag param.Field[string] `json:"address_tag"`
	// The nickname for this blockchain recipient
	Name param.Field[string] `json:"name"`
}

func (r BlockchainRecipientNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
