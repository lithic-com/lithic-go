package requests

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/lithic-com/lithic-go/core"
	"github.com/lithic-com/lithic-go/core/fields"
	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/core/query"
)

type ValidationMethodType BankValidationMethod

const (
	ValidationMethodTypeBank  ValidationMethodType = "BANK"
	ValidationMethodTypePlaid ValidationMethodType = "PLAID"
)

type FundingSourceNewParams struct {
	ValidationMethodType ValidationMethodType
	Bank                 Bank
	Plaid                Plaid
}

func (r *FundingSourceNewParams) MarshalJSON() (data []byte, err error) {
	switch r.ValidationMethodType {
	case "BANK":
		return pjson.Marshal(&r.Bank)
	case "PLAID":
		return pjson.Marshal(&r.Plaid)
	}
	return nil, errors.New("no valid union found for request")
}

type Bank struct {
	ValidationMethod fields.Field[BankValidationMethod] `json:"validation_method,required"`
	// The name associated with the bank account. This property is only for
	// identification purposes, and has no bearing on the external properties of the
	// bank.
	AccountName fields.Field[string] `json:"account_name"`
	// The account number of the bank account.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// Only required for multi-account users. Token identifying the account that the
	// bank account will be associated with. Only applicable if using account holder
	// enrollment. See
	// [Managing Your Program](https://docs.lithic.com/docs/managing-your-program) for
	// more information.
	AccountToken fields.Field[string] `json:"account_token" format:"uuid"`
	// The routing number of the bank account.
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
}

func (r Bank) String() (result string) {
	return fmt.Sprintf("&Bank{ValidationMethod:%s AccountName:%s AccountNumber:%s AccountToken:%s RoutingNumber:%s}", r.ValidationMethod, r.AccountName, r.AccountNumber, r.AccountToken, r.RoutingNumber)
}

type BankValidationMethod string

const (
	BankValidationMethodBank BankValidationMethod = "BANK"
)

type Plaid struct {
	ValidationMethod fields.Field[PlaidValidationMethod] `json:"validation_method,required"`
	// Only required for multi-account users. Token identifying the account associated
	// with the bank account. Only applicable if using account holder enrollment. See
	// [Managing Your Program](https://docs.lithic.com/docs/managing-your-program) for
	// more information.
	AccountToken fields.Field[string] `json:"account_token" format:"uuid"`
	// The processor token associated with the bank account.
	ProcessorToken fields.Field[string] `json:"processor_token,required"`
}

func (r Plaid) String() (result string) {
	return fmt.Sprintf("&Plaid{ValidationMethod:%s AccountToken:%s ProcessorToken:%s}", r.ValidationMethod, r.AccountToken, r.ProcessorToken)
}

type PlaidValidationMethod string

const (
	PlaidValidationMethodPlaid PlaidValidationMethod = "PLAID"
)

type FundingSourceUpdateParams struct {
	// Only required for multi-account users. Token identifying the account that the
	// bank account will be associated with. Only applicable if using account holder
	// enrollment. See
	// [Managing Your Program](https://docs.lithic.com/docs/managing-your-program) for
	// more information.
	AccountToken fields.Field[string] `json:"account_token" format:"uuid"`
	// The desired state of the bank account.
	//
	// If a bank account is set to `DELETED`, all cards linked to this account will no
	// longer be associated with it. If there are no other bank accounts in state
	// `ENABLED` on the account, authorizations will not be accepted on the card until
	// a new funding account is added.
	State fields.Field[FundingSourceUpdateParamsState] `json:"state"`
}

// MarshalJSON serializes FundingSourceUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *FundingSourceUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r FundingSourceUpdateParams) String() (result string) {
	return fmt.Sprintf("&FundingSourceUpdateParams{AccountToken:%s State:%s}", r.AccountToken, r.State)
}

type FundingSourceUpdateParamsState string

const (
	FundingSourceUpdateParamsStateDeleted FundingSourceUpdateParamsState = "DELETED"
	FundingSourceUpdateParamsStateEnabled FundingSourceUpdateParamsState = "ENABLED"
)

type FundingSourceListParams struct {
	AccountToken fields.Field[string] `query:"account_token" format:"uuid"`
	// Page (for pagination).
	Page fields.Field[int64] `query:"page"`
	// Page size (for pagination).
	PageSize fields.Field[int64] `query:"page_size"`
}

// URLQuery serializes FundingSourceListParams into a url.Values of the query
// parameters associated with this value
func (r *FundingSourceListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r FundingSourceListParams) String() (result string) {
	return fmt.Sprintf("&FundingSourceListParams{AccountToken:%s Page:%s PageSize:%s}", r.AccountToken, r.Page, r.PageSize)
}

type FundingSourceVerifyParams struct {
	// Only required for multi-account users. Token identifying the account that the
	// bank account will be associated with. Only applicable if using account holder
	// enrollment. See
	// [Managing Your Program](https://docs.lithic.com/docs/managing-your-program) for
	// more information.
	AccountToken fields.Field[string] `json:"account_token" format:"uuid"`
	// An array of dollar amounts (in cents) received in two credit transactions.
	MicroDeposits fields.Field[[]int64] `json:"micro_deposits,required"`
}

// MarshalJSON serializes FundingSourceVerifyParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *FundingSourceVerifyParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r FundingSourceVerifyParams) String() (result string) {
	return fmt.Sprintf("&FundingSourceVerifyParams{AccountToken:%s MicroDeposits:%s}", r.AccountToken, core.Fmt(r.MicroDeposits))
}
