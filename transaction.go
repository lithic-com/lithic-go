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
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/lithic-com/lithic-go/shared"
)

// TransactionService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options                []option.RequestOption
	EnhancedCommercialData *TransactionEnhancedCommercialDataService
	Events                 *TransactionEventService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	r.EnhancedCommercialData = NewTransactionEnhancedCommercialDataService(opts...)
	r.Events = NewTransactionEventService(opts...)
	return
}

// Get specific card transaction.
func (r *TransactionService) Get(ctx context.Context, transactionToken string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	if transactionToken == "" {
		err = errors.New("missing required transaction_token parameter")
		return
	}
	path := fmt.Sprintf("v1/transactions/%s", transactionToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List card transactions.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Transaction], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/transactions"
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

// List card transactions.
func (r *TransactionService) ListAutoPaging(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Transaction] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Simulates an authorization request from the payment network as if it came from a
// merchant acquirer. If you're configured for ASA, simulating auths requires your
// ASA client to be set up properly (respond with a valid JSON to the ASA request).
// For users that are not configured for ASA, a daily transaction limit of $5000
// USD is applied by default. This limit can be modified via the
// [update account](https://docs.lithic.com/reference/patchaccountbytoken)
// endpoint.
func (r *TransactionService) SimulateAuthorization(ctx context.Context, body TransactionSimulateAuthorizationParams, opts ...option.RequestOption) (res *TransactionSimulateAuthorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/authorize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates an authorization advice request from the payment network as if it came
// from a merchant acquirer. An authorization advice request changes the amount of
// the transaction.
func (r *TransactionService) SimulateAuthorizationAdvice(ctx context.Context, body TransactionSimulateAuthorizationAdviceParams, opts ...option.RequestOption) (res *TransactionSimulateAuthorizationAdviceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/authorization_advice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Clears an existing authorization. After this event, the transaction is no longer
// pending.
//
// If no `amount` is supplied to this endpoint, the amount of the transaction will
// be captured. Any transaction that has any amount completed at all do not have
// access to this behavior.
func (r *TransactionService) SimulateClearing(ctx context.Context, body TransactionSimulateClearingParams, opts ...option.RequestOption) (res *TransactionSimulateClearingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/clearing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates a credit authorization advice message from the payment network. This
// message indicates that a credit authorization was approved on your behalf by the
// network.
func (r *TransactionService) SimulateCreditAuthorization(ctx context.Context, body TransactionSimulateCreditAuthorizationParams, opts ...option.RequestOption) (res *TransactionSimulateCreditAuthorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/credit_authorization_advice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns (aka refunds) an amount back to a card. Returns are cleared immediately
// and do not spend time in a `PENDING` state.
func (r *TransactionService) SimulateReturn(ctx context.Context, body TransactionSimulateReturnParams, opts ...option.RequestOption) (res *TransactionSimulateReturnResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/return"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Voids a settled credit transaction – i.e., a transaction with a negative amount
// and `SETTLED` status. These can be credit authorizations that have already
// cleared or financial credit authorizations.
func (r *TransactionService) SimulateReturnReversal(ctx context.Context, body TransactionSimulateReturnReversalParams, opts ...option.RequestOption) (res *TransactionSimulateReturnReversalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/return_reversal"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Voids an existing, uncleared (aka pending) authorization. If amount is not sent
// the full amount will be voided. Cannot be used on partially completed
// transactions, but can be used on partially voided transactions. _Note that
// simulating an authorization expiry on credit authorizations or credit
// authorization advice is not currently supported but will be added soon._
func (r *TransactionService) SimulateVoid(ctx context.Context, body TransactionSimulateVoidParams, opts ...option.RequestOption) (res *TransactionSimulateVoidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/simulate/void"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Transaction struct {
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// The token for the account associated with this transaction.
	AccountToken string `json:"account_token,required" format:"uuid"`
	// Fee assessed by the merchant and paid for by the cardholder in the smallest unit
	// of the currency. Will be zero if no fee is assessed. Rebates may be transmitted
	// as a negative value to indicate credited fees.
	AcquirerFee int64 `json:"acquirer_fee,required,nullable"`
	// Unique identifier assigned to a transaction by the acquirer that can be used in
	// dispute and chargeback filing.
	AcquirerReferenceNumber string `json:"acquirer_reference_number,required,nullable"`
	// Authorization amount of the transaction (in cents), including any acquirer fees.
	// This may change over time, and will represent the settled amount once the
	// transaction is settled.
	Amount  int64              `json:"amount,required"`
	Amounts TransactionAmounts `json:"amounts,required"`
	// Authorization amount (in cents) of the transaction, including any acquirer fees.
	// This amount always represents the amount authorized for the transaction,
	// unaffected by settlement.
	AuthorizationAmount int64 `json:"authorization_amount,required,nullable"`
	// A fixed-width 6-digit numeric identifier that can be used to identify a
	// transaction with networks.
	AuthorizationCode string         `json:"authorization_code,required,nullable"`
	Avs               TransactionAvs `json:"avs,required,nullable"`
	// Token for the card used in this transaction.
	CardToken                string                              `json:"card_token,required" format:"uuid"`
	CardholderAuthentication TransactionCardholderAuthentication `json:"cardholder_authentication,required,nullable"`
	// Date and time when the transaction first occurred. UTC time zone.
	Created  time.Time           `json:"created,required" format:"date-time"`
	Merchant TransactionMerchant `json:"merchant,required"`
	// Analogous to the 'amount' property, but will represent the amount in the
	// transaction's local currency (smallest unit), including any acquirer fees.
	MerchantAmount int64 `json:"merchant_amount,required,nullable"`
	// Analogous to the 'authorization_amount' property, but will represent the amount
	// in the transaction's local currency (smallest unit), including any acquirer
	// fees.
	MerchantAuthorizationAmount int64 `json:"merchant_authorization_amount,required,nullable"`
	// 3-digit alphabetic ISO 4217 code for the local currency of the transaction.
	MerchantCurrency string `json:"merchant_currency,required"`
	// Card network of the authorization. Can be `INTERLINK`, `MAESTRO`, `MASTERCARD`,
	// `VISA`, or `UNKNOWN`. Value is `UNKNOWN` when Lithic cannot determine the
	// network code from the upstream provider.
	Network TransactionNetwork `json:"network,required,nullable"`
	// Network-provided score assessing risk level associated with a given
	// authorization. Scores are on a range of 0-999, with 0 representing the lowest
	// risk and 999 representing the highest risk. For Visa transactions, where the raw
	// score has a range of 0-99, Lithic will normalize the score by multiplying the
	// raw score by 10x.
	NetworkRiskScore int64             `json:"network_risk_score,required,nullable"`
	Pos              TransactionPos    `json:"pos,required"`
	Result           TransactionResult `json:"result,required"`
	// Amount of the transaction that has been settled (in cents), including any
	// acquirer fees. This may change over time.
	SettledAmount int64 `json:"settled_amount,required"`
	// Status of the transaction.
	Status    TransactionStatus    `json:"status,required"`
	TokenInfo TransactionTokenInfo `json:"token_info,required,nullable"`
	// Date and time when the transaction last updated. UTC time zone.
	Updated time.Time          `json:"updated,required" format:"date-time"`
	Events  []TransactionEvent `json:"events"`
	JSON    transactionJSON    `json:"-"`
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
	Token                       apijson.Field
	AccountToken                apijson.Field
	AcquirerFee                 apijson.Field
	AcquirerReferenceNumber     apijson.Field
	Amount                      apijson.Field
	Amounts                     apijson.Field
	AuthorizationAmount         apijson.Field
	AuthorizationCode           apijson.Field
	Avs                         apijson.Field
	CardToken                   apijson.Field
	CardholderAuthentication    apijson.Field
	Created                     apijson.Field
	Merchant                    apijson.Field
	MerchantAmount              apijson.Field
	MerchantAuthorizationAmount apijson.Field
	MerchantCurrency            apijson.Field
	Network                     apijson.Field
	NetworkRiskScore            apijson.Field
	Pos                         apijson.Field
	Result                      apijson.Field
	SettledAmount               apijson.Field
	Status                      apijson.Field
	TokenInfo                   apijson.Field
	Updated                     apijson.Field
	Events                      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionJSON) RawJSON() string {
	return r.raw
}

type TransactionAmounts struct {
	Cardholder TransactionAmountsCardholder `json:"cardholder,required"`
	Hold       TransactionAmountsHold       `json:"hold,required"`
	Merchant   TransactionAmountsMerchant   `json:"merchant,required"`
	Settlement TransactionAmountsSettlement `json:"settlement,required"`
	JSON       transactionAmountsJSON       `json:"-"`
}

// transactionAmountsJSON contains the JSON metadata for the struct
// [TransactionAmounts]
type transactionAmountsJSON struct {
	Cardholder  apijson.Field
	Hold        apijson.Field
	Merchant    apijson.Field
	Settlement  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAmountsJSON) RawJSON() string {
	return r.raw
}

type TransactionAmountsCardholder struct {
	// The aggregate settled amount in the cardholder's local currency.
	Amount int64 `json:"amount,required"`
	// The conversion rate used to convert the merchant amount to the cardholder
	// amount.
	ConversionRate string `json:"conversion_rate,required"`
	// ISO 4217 currency. Its enumerants are ISO 4217 currencies except for some
	// special currencies like `XXX`. Enumerants names are lowercase currency code e.g.
	// `EUR`, `USD`.
	Currency shared.Currency                  `json:"currency,required"`
	JSON     transactionAmountsCardholderJSON `json:"-"`
}

// transactionAmountsCardholderJSON contains the JSON metadata for the struct
// [TransactionAmountsCardholder]
type transactionAmountsCardholderJSON struct {
	Amount         apijson.Field
	ConversionRate apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionAmountsCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAmountsCardholderJSON) RawJSON() string {
	return r.raw
}

type TransactionAmountsHold struct {
	// The aggregate pending amount in the anticipated settlement currency.
	Amount int64 `json:"amount,required"`
	// ISO 4217 currency. Its enumerants are ISO 4217 currencies except for some
	// special currencies like `XXX`. Enumerants names are lowercase currency code e.g.
	// `EUR`, `USD`.
	Currency shared.Currency            `json:"currency,required"`
	JSON     transactionAmountsHoldJSON `json:"-"`
}

// transactionAmountsHoldJSON contains the JSON metadata for the struct
// [TransactionAmountsHold]
type transactionAmountsHoldJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionAmountsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAmountsHoldJSON) RawJSON() string {
	return r.raw
}

type TransactionAmountsMerchant struct {
	// The aggregate settled amount in the merchant's local currency.
	Amount int64 `json:"amount,required"`
	// ISO 4217 currency. Its enumerants are ISO 4217 currencies except for some
	// special currencies like `XXX`. Enumerants names are lowercase currency code e.g.
	// `EUR`, `USD`.
	Currency shared.Currency                `json:"currency,required"`
	JSON     transactionAmountsMerchantJSON `json:"-"`
}

// transactionAmountsMerchantJSON contains the JSON metadata for the struct
// [TransactionAmountsMerchant]
type transactionAmountsMerchantJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionAmountsMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAmountsMerchantJSON) RawJSON() string {
	return r.raw
}

type TransactionAmountsSettlement struct {
	// The aggregate settled amount in the settlement currency.
	Amount int64 `json:"amount,required"`
	// ISO 4217 currency. Its enumerants are ISO 4217 currencies except for some
	// special currencies like `XXX`. Enumerants names are lowercase currency code e.g.
	// `EUR`, `USD`.
	Currency shared.Currency                  `json:"currency,required"`
	JSON     transactionAmountsSettlementJSON `json:"-"`
}

// transactionAmountsSettlementJSON contains the JSON metadata for the struct
// [TransactionAmountsSettlement]
type transactionAmountsSettlementJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionAmountsSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAmountsSettlementJSON) RawJSON() string {
	return r.raw
}

type TransactionAvs struct {
	// Cardholder address
	Address string `json:"address,required"`
	// Cardholder ZIP code
	Zipcode string             `json:"zipcode,required"`
	JSON    transactionAvsJSON `json:"-"`
}

// transactionAvsJSON contains the JSON metadata for the struct [TransactionAvs]
type transactionAvsJSON struct {
	Address     apijson.Field
	Zipcode     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionAvs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAvsJSON) RawJSON() string {
	return r.raw
}

type TransactionCardholderAuthentication struct {
	// The 3DS version used for the authentication
	ThreeDSVersion string `json:"3ds_version,required,nullable"`
	// Whether an acquirer exemption applied to the transaction.
	AcquirerExemption TransactionCardholderAuthenticationAcquirerExemption `json:"acquirer_exemption,required"`
	// Indicates what the outcome of the 3DS authentication process is.
	AuthenticationResult TransactionCardholderAuthenticationAuthenticationResult `json:"authentication_result,required"`
	// Indicates which party made the 3DS authentication decision.
	DecisionMadeBy TransactionCardholderAuthenticationDecisionMadeBy `json:"decision_made_by,required"`
	// Indicates whether chargeback liability shift applies to the transaction.
	// Possible enum values:
	//
	//   - `3DS_AUTHENTICATED`: The transaction was fully authenticated through a 3-D Secure flow, chargeback liability shift applies.
	//
	//   - `ACQUIRER_EXEMPTION`: The acquirer utilised an exemption to bypass Strong Customer Authentication (`transStatus = N`, or `transStatus = I`). Liability remains with the acquirer and in this case the `acquirer_exemption` field is expected to be not `NONE`.
	//
	//   - `NONE`: Chargeback liability shift has not shifted to the issuer, i.e. the merchant is liable.
	//
	//   - `TOKEN_AUTHENTICATED`: The transaction was a tokenized payment with validated
	//     cryptography, possibly recurring. Chargeback liability shift to the issuer
	//     applies.
	LiabilityShift TransactionCardholderAuthenticationLiabilityShift `json:"liability_shift,required"`
	// Unique identifier you can use to match a given 3DS authentication (available via
	// the three_ds_authentication.created event webhook) and the transaction. Note
	// that in cases where liability shift does not occur, this token is matched to the
	// transaction on a best-effort basis.
	ThreeDSAuthenticationToken string `json:"three_ds_authentication_token,required,nullable" format:"uuid"`
	// Indicates whether a 3DS challenge flow was used, and if so, what the
	// verification method was. (deprecated, use `authentication_result`)
	VerificationAttempted TransactionCardholderAuthenticationVerificationAttempted `json:"verification_attempted,required"`
	// Indicates whether a transaction is considered 3DS authenticated. (deprecated,
	// use `authentication_result`)
	VerificationResult TransactionCardholderAuthenticationVerificationResult `json:"verification_result,required"`
	JSON               transactionCardholderAuthenticationJSON               `json:"-"`
}

// transactionCardholderAuthenticationJSON contains the JSON metadata for the
// struct [TransactionCardholderAuthentication]
type transactionCardholderAuthenticationJSON struct {
	ThreeDSVersion             apijson.Field
	AcquirerExemption          apijson.Field
	AuthenticationResult       apijson.Field
	DecisionMadeBy             apijson.Field
	LiabilityShift             apijson.Field
	ThreeDSAuthenticationToken apijson.Field
	VerificationAttempted      apijson.Field
	VerificationResult         apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *TransactionCardholderAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionCardholderAuthenticationJSON) RawJSON() string {
	return r.raw
}

// Whether an acquirer exemption applied to the transaction.
type TransactionCardholderAuthenticationAcquirerExemption string

const (
	TransactionCardholderAuthenticationAcquirerExemptionAuthenticationOutageException          TransactionCardholderAuthenticationAcquirerExemption = "AUTHENTICATION_OUTAGE_EXCEPTION"
	TransactionCardholderAuthenticationAcquirerExemptionLowValue                               TransactionCardholderAuthenticationAcquirerExemption = "LOW_VALUE"
	TransactionCardholderAuthenticationAcquirerExemptionMerchantInitiatedTransaction           TransactionCardholderAuthenticationAcquirerExemption = "MERCHANT_INITIATED_TRANSACTION"
	TransactionCardholderAuthenticationAcquirerExemptionNone                                   TransactionCardholderAuthenticationAcquirerExemption = "NONE"
	TransactionCardholderAuthenticationAcquirerExemptionRecurringPayment                       TransactionCardholderAuthenticationAcquirerExemption = "RECURRING_PAYMENT"
	TransactionCardholderAuthenticationAcquirerExemptionSecureCorporatePayment                 TransactionCardholderAuthenticationAcquirerExemption = "SECURE_CORPORATE_PAYMENT"
	TransactionCardholderAuthenticationAcquirerExemptionStrongCustomerAuthenticationDelegation TransactionCardholderAuthenticationAcquirerExemption = "STRONG_CUSTOMER_AUTHENTICATION_DELEGATION"
	TransactionCardholderAuthenticationAcquirerExemptionTransactionRiskAnalysis                TransactionCardholderAuthenticationAcquirerExemption = "TRANSACTION_RISK_ANALYSIS"
)

func (r TransactionCardholderAuthenticationAcquirerExemption) IsKnown() bool {
	switch r {
	case TransactionCardholderAuthenticationAcquirerExemptionAuthenticationOutageException, TransactionCardholderAuthenticationAcquirerExemptionLowValue, TransactionCardholderAuthenticationAcquirerExemptionMerchantInitiatedTransaction, TransactionCardholderAuthenticationAcquirerExemptionNone, TransactionCardholderAuthenticationAcquirerExemptionRecurringPayment, TransactionCardholderAuthenticationAcquirerExemptionSecureCorporatePayment, TransactionCardholderAuthenticationAcquirerExemptionStrongCustomerAuthenticationDelegation, TransactionCardholderAuthenticationAcquirerExemptionTransactionRiskAnalysis:
		return true
	}
	return false
}

// Indicates what the outcome of the 3DS authentication process is.
type TransactionCardholderAuthenticationAuthenticationResult string

const (
	TransactionCardholderAuthenticationAuthenticationResultAttempts TransactionCardholderAuthenticationAuthenticationResult = "ATTEMPTS"
	TransactionCardholderAuthenticationAuthenticationResultDecline  TransactionCardholderAuthenticationAuthenticationResult = "DECLINE"
	TransactionCardholderAuthenticationAuthenticationResultNone     TransactionCardholderAuthenticationAuthenticationResult = "NONE"
	TransactionCardholderAuthenticationAuthenticationResultSuccess  TransactionCardholderAuthenticationAuthenticationResult = "SUCCESS"
)

func (r TransactionCardholderAuthenticationAuthenticationResult) IsKnown() bool {
	switch r {
	case TransactionCardholderAuthenticationAuthenticationResultAttempts, TransactionCardholderAuthenticationAuthenticationResultDecline, TransactionCardholderAuthenticationAuthenticationResultNone, TransactionCardholderAuthenticationAuthenticationResultSuccess:
		return true
	}
	return false
}

// Indicates which party made the 3DS authentication decision.
type TransactionCardholderAuthenticationDecisionMadeBy string

const (
	TransactionCardholderAuthenticationDecisionMadeByCustomerEndpoint TransactionCardholderAuthenticationDecisionMadeBy = "CUSTOMER_ENDPOINT"
	TransactionCardholderAuthenticationDecisionMadeByLithicDefault    TransactionCardholderAuthenticationDecisionMadeBy = "LITHIC_DEFAULT"
	TransactionCardholderAuthenticationDecisionMadeByLithicRules      TransactionCardholderAuthenticationDecisionMadeBy = "LITHIC_RULES"
	TransactionCardholderAuthenticationDecisionMadeByNetwork          TransactionCardholderAuthenticationDecisionMadeBy = "NETWORK"
	TransactionCardholderAuthenticationDecisionMadeByUnknown          TransactionCardholderAuthenticationDecisionMadeBy = "UNKNOWN"
)

func (r TransactionCardholderAuthenticationDecisionMadeBy) IsKnown() bool {
	switch r {
	case TransactionCardholderAuthenticationDecisionMadeByCustomerEndpoint, TransactionCardholderAuthenticationDecisionMadeByLithicDefault, TransactionCardholderAuthenticationDecisionMadeByLithicRules, TransactionCardholderAuthenticationDecisionMadeByNetwork, TransactionCardholderAuthenticationDecisionMadeByUnknown:
		return true
	}
	return false
}

// Indicates whether chargeback liability shift applies to the transaction.
// Possible enum values:
//
//   - `3DS_AUTHENTICATED`: The transaction was fully authenticated through a 3-D Secure flow, chargeback liability shift applies.
//
//   - `ACQUIRER_EXEMPTION`: The acquirer utilised an exemption to bypass Strong Customer Authentication (`transStatus = N`, or `transStatus = I`). Liability remains with the acquirer and in this case the `acquirer_exemption` field is expected to be not `NONE`.
//
//   - `NONE`: Chargeback liability shift has not shifted to the issuer, i.e. the merchant is liable.
//
//   - `TOKEN_AUTHENTICATED`: The transaction was a tokenized payment with validated
//     cryptography, possibly recurring. Chargeback liability shift to the issuer
//     applies.
type TransactionCardholderAuthenticationLiabilityShift string

const (
	TransactionCardholderAuthenticationLiabilityShift3DSAuthenticated   TransactionCardholderAuthenticationLiabilityShift = "3DS_AUTHENTICATED"
	TransactionCardholderAuthenticationLiabilityShiftAcquirerExemption  TransactionCardholderAuthenticationLiabilityShift = "ACQUIRER_EXEMPTION"
	TransactionCardholderAuthenticationLiabilityShiftNone               TransactionCardholderAuthenticationLiabilityShift = "NONE"
	TransactionCardholderAuthenticationLiabilityShiftTokenAuthenticated TransactionCardholderAuthenticationLiabilityShift = "TOKEN_AUTHENTICATED"
)

func (r TransactionCardholderAuthenticationLiabilityShift) IsKnown() bool {
	switch r {
	case TransactionCardholderAuthenticationLiabilityShift3DSAuthenticated, TransactionCardholderAuthenticationLiabilityShiftAcquirerExemption, TransactionCardholderAuthenticationLiabilityShiftNone, TransactionCardholderAuthenticationLiabilityShiftTokenAuthenticated:
		return true
	}
	return false
}

// Indicates whether a 3DS challenge flow was used, and if so, what the
// verification method was. (deprecated, use `authentication_result`)
type TransactionCardholderAuthenticationVerificationAttempted string

const (
	TransactionCardholderAuthenticationVerificationAttemptedNone  TransactionCardholderAuthenticationVerificationAttempted = "NONE"
	TransactionCardholderAuthenticationVerificationAttemptedOther TransactionCardholderAuthenticationVerificationAttempted = "OTHER"
)

func (r TransactionCardholderAuthenticationVerificationAttempted) IsKnown() bool {
	switch r {
	case TransactionCardholderAuthenticationVerificationAttemptedNone, TransactionCardholderAuthenticationVerificationAttemptedOther:
		return true
	}
	return false
}

// Indicates whether a transaction is considered 3DS authenticated. (deprecated,
// use `authentication_result`)
type TransactionCardholderAuthenticationVerificationResult string

const (
	TransactionCardholderAuthenticationVerificationResultCancelled    TransactionCardholderAuthenticationVerificationResult = "CANCELLED"
	TransactionCardholderAuthenticationVerificationResultFailed       TransactionCardholderAuthenticationVerificationResult = "FAILED"
	TransactionCardholderAuthenticationVerificationResultFrictionless TransactionCardholderAuthenticationVerificationResult = "FRICTIONLESS"
	TransactionCardholderAuthenticationVerificationResultNotAttempted TransactionCardholderAuthenticationVerificationResult = "NOT_ATTEMPTED"
	TransactionCardholderAuthenticationVerificationResultRejected     TransactionCardholderAuthenticationVerificationResult = "REJECTED"
	TransactionCardholderAuthenticationVerificationResultSuccess      TransactionCardholderAuthenticationVerificationResult = "SUCCESS"
)

func (r TransactionCardholderAuthenticationVerificationResult) IsKnown() bool {
	switch r {
	case TransactionCardholderAuthenticationVerificationResultCancelled, TransactionCardholderAuthenticationVerificationResultFailed, TransactionCardholderAuthenticationVerificationResultFrictionless, TransactionCardholderAuthenticationVerificationResultNotAttempted, TransactionCardholderAuthenticationVerificationResultRejected, TransactionCardholderAuthenticationVerificationResultSuccess:
		return true
	}
	return false
}

type TransactionMerchant struct {
	// Unique alphanumeric identifier for the payment card acceptor (merchant).
	AcceptorID string `json:"acceptor_id,required"`
	// Unique numeric identifier of the acquiring institution.
	AcquiringInstitutionID string `json:"acquiring_institution_id,required"`
	// City of card acceptor. Note that in many cases, particularly in card-not-present
	// transactions, merchants may send through a phone number or URL in this field.
	City string `json:"city,required"`
	// Country or entity of card acceptor. Possible values are: (1) all ISO 3166-1
	// alpha-3 country codes, (2) QZZ for Kosovo, and (3) ANT for Netherlands Antilles.
	Country string `json:"country,required"`
	// Short description of card acceptor.
	Descriptor string `json:"descriptor,required"`
	// Merchant category code (MCC). A four-digit number listed in ISO 18245. An MCC is
	// used to classify a business by the types of goods or services it provides.
	Mcc string `json:"mcc,required"`
	// Geographic state of card acceptor.
	State string                  `json:"state,required"`
	JSON  transactionMerchantJSON `json:"-"`
}

// transactionMerchantJSON contains the JSON metadata for the struct
// [TransactionMerchant]
type transactionMerchantJSON struct {
	AcceptorID             apijson.Field
	AcquiringInstitutionID apijson.Field
	City                   apijson.Field
	Country                apijson.Field
	Descriptor             apijson.Field
	Mcc                    apijson.Field
	State                  apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TransactionMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionMerchantJSON) RawJSON() string {
	return r.raw
}

// Card network of the authorization. Can be `INTERLINK`, `MAESTRO`, `MASTERCARD`,
// `VISA`, or `UNKNOWN`. Value is `UNKNOWN` when Lithic cannot determine the
// network code from the upstream provider.
type TransactionNetwork string

const (
	TransactionNetworkInterlink  TransactionNetwork = "INTERLINK"
	TransactionNetworkMaestro    TransactionNetwork = "MAESTRO"
	TransactionNetworkMastercard TransactionNetwork = "MASTERCARD"
	TransactionNetworkUnknown    TransactionNetwork = "UNKNOWN"
	TransactionNetworkVisa       TransactionNetwork = "VISA"
)

func (r TransactionNetwork) IsKnown() bool {
	switch r {
	case TransactionNetworkInterlink, TransactionNetworkMaestro, TransactionNetworkMastercard, TransactionNetworkUnknown, TransactionNetworkVisa:
		return true
	}
	return false
}

type TransactionPos struct {
	EntryMode TransactionPosEntryMode `json:"entry_mode,required"`
	Terminal  TransactionPosTerminal  `json:"terminal,required"`
	JSON      transactionPosJSON      `json:"-"`
}

// transactionPosJSON contains the JSON metadata for the struct [TransactionPos]
type transactionPosJSON struct {
	EntryMode   apijson.Field
	Terminal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionPos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionPosJSON) RawJSON() string {
	return r.raw
}

type TransactionPosEntryMode struct {
	// Card presence indicator
	Card TransactionPosEntryModeCard `json:"card,required"`
	// Cardholder presence indicator
	Cardholder TransactionPosEntryModeCardholder `json:"cardholder,required"`
	// Method of entry for the PAN
	Pan TransactionPosEntryModePan `json:"pan,required"`
	// Indicates whether the cardholder entered the PIN. True if the PIN was entered.
	PinEntered bool                        `json:"pin_entered,required"`
	JSON       transactionPosEntryModeJSON `json:"-"`
}

// transactionPosEntryModeJSON contains the JSON metadata for the struct
// [TransactionPosEntryMode]
type transactionPosEntryModeJSON struct {
	Card        apijson.Field
	Cardholder  apijson.Field
	Pan         apijson.Field
	PinEntered  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionPosEntryMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionPosEntryModeJSON) RawJSON() string {
	return r.raw
}

// Card presence indicator
type TransactionPosEntryModeCard string

const (
	TransactionPosEntryModeCardNotPresent    TransactionPosEntryModeCard = "NOT_PRESENT"
	TransactionPosEntryModeCardPreauthorized TransactionPosEntryModeCard = "PREAUTHORIZED"
	TransactionPosEntryModeCardPresent       TransactionPosEntryModeCard = "PRESENT"
	TransactionPosEntryModeCardUnknown       TransactionPosEntryModeCard = "UNKNOWN"
)

func (r TransactionPosEntryModeCard) IsKnown() bool {
	switch r {
	case TransactionPosEntryModeCardNotPresent, TransactionPosEntryModeCardPreauthorized, TransactionPosEntryModeCardPresent, TransactionPosEntryModeCardUnknown:
		return true
	}
	return false
}

// Cardholder presence indicator
type TransactionPosEntryModeCardholder string

const (
	TransactionPosEntryModeCardholderDeferredBilling TransactionPosEntryModeCardholder = "DEFERRED_BILLING"
	TransactionPosEntryModeCardholderElectronicOrder TransactionPosEntryModeCardholder = "ELECTRONIC_ORDER"
	TransactionPosEntryModeCardholderInstallment     TransactionPosEntryModeCardholder = "INSTALLMENT"
	TransactionPosEntryModeCardholderMailOrder       TransactionPosEntryModeCardholder = "MAIL_ORDER"
	TransactionPosEntryModeCardholderNotPresent      TransactionPosEntryModeCardholder = "NOT_PRESENT"
	TransactionPosEntryModeCardholderPreauthorized   TransactionPosEntryModeCardholder = "PREAUTHORIZED"
	TransactionPosEntryModeCardholderPresent         TransactionPosEntryModeCardholder = "PRESENT"
	TransactionPosEntryModeCardholderReoccurring     TransactionPosEntryModeCardholder = "REOCCURRING"
	TransactionPosEntryModeCardholderTelephoneOrder  TransactionPosEntryModeCardholder = "TELEPHONE_ORDER"
	TransactionPosEntryModeCardholderUnknown         TransactionPosEntryModeCardholder = "UNKNOWN"
)

func (r TransactionPosEntryModeCardholder) IsKnown() bool {
	switch r {
	case TransactionPosEntryModeCardholderDeferredBilling, TransactionPosEntryModeCardholderElectronicOrder, TransactionPosEntryModeCardholderInstallment, TransactionPosEntryModeCardholderMailOrder, TransactionPosEntryModeCardholderNotPresent, TransactionPosEntryModeCardholderPreauthorized, TransactionPosEntryModeCardholderPresent, TransactionPosEntryModeCardholderReoccurring, TransactionPosEntryModeCardholderTelephoneOrder, TransactionPosEntryModeCardholderUnknown:
		return true
	}
	return false
}

// Method of entry for the PAN
type TransactionPosEntryModePan string

const (
	TransactionPosEntryModePanAutoEntry           TransactionPosEntryModePan = "AUTO_ENTRY"
	TransactionPosEntryModePanBarCode             TransactionPosEntryModePan = "BAR_CODE"
	TransactionPosEntryModePanContactless         TransactionPosEntryModePan = "CONTACTLESS"
	TransactionPosEntryModePanCredentialOnFile    TransactionPosEntryModePan = "CREDENTIAL_ON_FILE"
	TransactionPosEntryModePanEcommerce           TransactionPosEntryModePan = "ECOMMERCE"
	TransactionPosEntryModePanErrorKeyed          TransactionPosEntryModePan = "ERROR_KEYED"
	TransactionPosEntryModePanErrorMagneticStripe TransactionPosEntryModePan = "ERROR_MAGNETIC_STRIPE"
	TransactionPosEntryModePanIcc                 TransactionPosEntryModePan = "ICC"
	TransactionPosEntryModePanKeyEntered          TransactionPosEntryModePan = "KEY_ENTERED"
	TransactionPosEntryModePanMagneticStripe      TransactionPosEntryModePan = "MAGNETIC_STRIPE"
	TransactionPosEntryModePanManual              TransactionPosEntryModePan = "MANUAL"
	TransactionPosEntryModePanOcr                 TransactionPosEntryModePan = "OCR"
	TransactionPosEntryModePanSecureCardless      TransactionPosEntryModePan = "SECURE_CARDLESS"
	TransactionPosEntryModePanUnknown             TransactionPosEntryModePan = "UNKNOWN"
	TransactionPosEntryModePanUnspecified         TransactionPosEntryModePan = "UNSPECIFIED"
)

func (r TransactionPosEntryModePan) IsKnown() bool {
	switch r {
	case TransactionPosEntryModePanAutoEntry, TransactionPosEntryModePanBarCode, TransactionPosEntryModePanContactless, TransactionPosEntryModePanCredentialOnFile, TransactionPosEntryModePanEcommerce, TransactionPosEntryModePanErrorKeyed, TransactionPosEntryModePanErrorMagneticStripe, TransactionPosEntryModePanIcc, TransactionPosEntryModePanKeyEntered, TransactionPosEntryModePanMagneticStripe, TransactionPosEntryModePanManual, TransactionPosEntryModePanOcr, TransactionPosEntryModePanSecureCardless, TransactionPosEntryModePanUnknown, TransactionPosEntryModePanUnspecified:
		return true
	}
	return false
}

type TransactionPosTerminal struct {
	// True if a clerk is present at the sale.
	Attended bool `json:"attended,required"`
	// True if the terminal is capable of retaining the card.
	CardRetentionCapable bool `json:"card_retention_capable,required"`
	// True if the sale was made at the place of business (vs. mobile).
	OnPremise bool `json:"on_premise,required"`
	// The person that is designated to swipe the card
	Operator TransactionPosTerminalOperator `json:"operator,required"`
	// True if the terminal is capable of partial approval. Partial approval is when
	// part of a transaction is approved and another payment must be used for the
	// remainder. Example scenario: A $40 transaction is attempted on a prepaid card
	// with a $25 balance. If partial approval is enabled, $25 can be authorized, at
	// which point the POS will prompt the user for an additional payment of $15.
	PartialApprovalCapable bool `json:"partial_approval_capable,required"`
	// Status of whether the POS is able to accept PINs
	PinCapability TransactionPosTerminalPinCapability `json:"pin_capability,required"`
	// POS Type
	Type TransactionPosTerminalType `json:"type,required"`
	JSON transactionPosTerminalJSON `json:"-"`
}

// transactionPosTerminalJSON contains the JSON metadata for the struct
// [TransactionPosTerminal]
type transactionPosTerminalJSON struct {
	Attended               apijson.Field
	CardRetentionCapable   apijson.Field
	OnPremise              apijson.Field
	Operator               apijson.Field
	PartialApprovalCapable apijson.Field
	PinCapability          apijson.Field
	Type                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TransactionPosTerminal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionPosTerminalJSON) RawJSON() string {
	return r.raw
}

// The person that is designated to swipe the card
type TransactionPosTerminalOperator string

const (
	TransactionPosTerminalOperatorAdministrative TransactionPosTerminalOperator = "ADMINISTRATIVE"
	TransactionPosTerminalOperatorCardholder     TransactionPosTerminalOperator = "CARDHOLDER"
	TransactionPosTerminalOperatorCardAcceptor   TransactionPosTerminalOperator = "CARD_ACCEPTOR"
	TransactionPosTerminalOperatorUnknown        TransactionPosTerminalOperator = "UNKNOWN"
)

func (r TransactionPosTerminalOperator) IsKnown() bool {
	switch r {
	case TransactionPosTerminalOperatorAdministrative, TransactionPosTerminalOperatorCardholder, TransactionPosTerminalOperatorCardAcceptor, TransactionPosTerminalOperatorUnknown:
		return true
	}
	return false
}

// Status of whether the POS is able to accept PINs
type TransactionPosTerminalPinCapability string

const (
	TransactionPosTerminalPinCapabilityCapable     TransactionPosTerminalPinCapability = "CAPABLE"
	TransactionPosTerminalPinCapabilityInoperative TransactionPosTerminalPinCapability = "INOPERATIVE"
	TransactionPosTerminalPinCapabilityNotCapable  TransactionPosTerminalPinCapability = "NOT_CAPABLE"
	TransactionPosTerminalPinCapabilityUnspecified TransactionPosTerminalPinCapability = "UNSPECIFIED"
)

func (r TransactionPosTerminalPinCapability) IsKnown() bool {
	switch r {
	case TransactionPosTerminalPinCapabilityCapable, TransactionPosTerminalPinCapabilityInoperative, TransactionPosTerminalPinCapabilityNotCapable, TransactionPosTerminalPinCapabilityUnspecified:
		return true
	}
	return false
}

// POS Type
type TransactionPosTerminalType string

const (
	TransactionPosTerminalTypeAdministrative        TransactionPosTerminalType = "ADMINISTRATIVE"
	TransactionPosTerminalTypeAtm                   TransactionPosTerminalType = "ATM"
	TransactionPosTerminalTypeAuthorization         TransactionPosTerminalType = "AUTHORIZATION"
	TransactionPosTerminalTypeCouponMachine         TransactionPosTerminalType = "COUPON_MACHINE"
	TransactionPosTerminalTypeDialTerminal          TransactionPosTerminalType = "DIAL_TERMINAL"
	TransactionPosTerminalTypeEcommerce             TransactionPosTerminalType = "ECOMMERCE"
	TransactionPosTerminalTypeEcr                   TransactionPosTerminalType = "ECR"
	TransactionPosTerminalTypeFuelMachine           TransactionPosTerminalType = "FUEL_MACHINE"
	TransactionPosTerminalTypeHomeTerminal          TransactionPosTerminalType = "HOME_TERMINAL"
	TransactionPosTerminalTypeMicr                  TransactionPosTerminalType = "MICR"
	TransactionPosTerminalTypeOffPremise            TransactionPosTerminalType = "OFF_PREMISE"
	TransactionPosTerminalTypePayment               TransactionPosTerminalType = "PAYMENT"
	TransactionPosTerminalTypePda                   TransactionPosTerminalType = "PDA"
	TransactionPosTerminalTypePhone                 TransactionPosTerminalType = "PHONE"
	TransactionPosTerminalTypePoint                 TransactionPosTerminalType = "POINT"
	TransactionPosTerminalTypePosTerminal           TransactionPosTerminalType = "POS_TERMINAL"
	TransactionPosTerminalTypePublicUtility         TransactionPosTerminalType = "PUBLIC_UTILITY"
	TransactionPosTerminalTypeSelfService           TransactionPosTerminalType = "SELF_SERVICE"
	TransactionPosTerminalTypeTelevision            TransactionPosTerminalType = "TELEVISION"
	TransactionPosTerminalTypeTeller                TransactionPosTerminalType = "TELLER"
	TransactionPosTerminalTypeTravelersCheckMachine TransactionPosTerminalType = "TRAVELERS_CHECK_MACHINE"
	TransactionPosTerminalTypeVending               TransactionPosTerminalType = "VENDING"
	TransactionPosTerminalTypeVoice                 TransactionPosTerminalType = "VOICE"
	TransactionPosTerminalTypeUnknown               TransactionPosTerminalType = "UNKNOWN"
)

func (r TransactionPosTerminalType) IsKnown() bool {
	switch r {
	case TransactionPosTerminalTypeAdministrative, TransactionPosTerminalTypeAtm, TransactionPosTerminalTypeAuthorization, TransactionPosTerminalTypeCouponMachine, TransactionPosTerminalTypeDialTerminal, TransactionPosTerminalTypeEcommerce, TransactionPosTerminalTypeEcr, TransactionPosTerminalTypeFuelMachine, TransactionPosTerminalTypeHomeTerminal, TransactionPosTerminalTypeMicr, TransactionPosTerminalTypeOffPremise, TransactionPosTerminalTypePayment, TransactionPosTerminalTypePda, TransactionPosTerminalTypePhone, TransactionPosTerminalTypePoint, TransactionPosTerminalTypePosTerminal, TransactionPosTerminalTypePublicUtility, TransactionPosTerminalTypeSelfService, TransactionPosTerminalTypeTelevision, TransactionPosTerminalTypeTeller, TransactionPosTerminalTypeTravelersCheckMachine, TransactionPosTerminalTypeVending, TransactionPosTerminalTypeVoice, TransactionPosTerminalTypeUnknown:
		return true
	}
	return false
}

type TransactionResult string

const (
	TransactionResultAccountStateTransactionFail TransactionResult = "ACCOUNT_STATE_TRANSACTION_FAIL"
	TransactionResultApproved                    TransactionResult = "APPROVED"
	TransactionResultBankConnectionError         TransactionResult = "BANK_CONNECTION_ERROR"
	TransactionResultBankNotVerified             TransactionResult = "BANK_NOT_VERIFIED"
	TransactionResultCardClosed                  TransactionResult = "CARD_CLOSED"
	TransactionResultCardPaused                  TransactionResult = "CARD_PAUSED"
	TransactionResultDeclined                    TransactionResult = "DECLINED"
	TransactionResultFraudAdvice                 TransactionResult = "FRAUD_ADVICE"
	TransactionResultIgnoredTtlExpiry            TransactionResult = "IGNORED_TTL_EXPIRY"
	TransactionResultInactiveAccount             TransactionResult = "INACTIVE_ACCOUNT"
	TransactionResultIncorrectPin                TransactionResult = "INCORRECT_PIN"
	TransactionResultInvalidCardDetails          TransactionResult = "INVALID_CARD_DETAILS"
	TransactionResultInsufficientFunds           TransactionResult = "INSUFFICIENT_FUNDS"
	TransactionResultInsufficientFundsPreload    TransactionResult = "INSUFFICIENT_FUNDS_PRELOAD"
	TransactionResultInvalidTransaction          TransactionResult = "INVALID_TRANSACTION"
	TransactionResultMerchantBlacklist           TransactionResult = "MERCHANT_BLACKLIST"
	TransactionResultOriginalNotFound            TransactionResult = "ORIGINAL_NOT_FOUND"
	TransactionResultPreviouslyCompleted         TransactionResult = "PREVIOUSLY_COMPLETED"
	TransactionResultSingleUseRecharged          TransactionResult = "SINGLE_USE_RECHARGED"
	TransactionResultSwitchInoperativeAdvice     TransactionResult = "SWITCH_INOPERATIVE_ADVICE"
	TransactionResultUnauthorizedMerchant        TransactionResult = "UNAUTHORIZED_MERCHANT"
	TransactionResultUnknownHostTimeout          TransactionResult = "UNKNOWN_HOST_TIMEOUT"
	TransactionResultUserTransactionLimit        TransactionResult = "USER_TRANSACTION_LIMIT"
)

func (r TransactionResult) IsKnown() bool {
	switch r {
	case TransactionResultAccountStateTransactionFail, TransactionResultApproved, TransactionResultBankConnectionError, TransactionResultBankNotVerified, TransactionResultCardClosed, TransactionResultCardPaused, TransactionResultDeclined, TransactionResultFraudAdvice, TransactionResultIgnoredTtlExpiry, TransactionResultInactiveAccount, TransactionResultIncorrectPin, TransactionResultInvalidCardDetails, TransactionResultInsufficientFunds, TransactionResultInsufficientFundsPreload, TransactionResultInvalidTransaction, TransactionResultMerchantBlacklist, TransactionResultOriginalNotFound, TransactionResultPreviouslyCompleted, TransactionResultSingleUseRecharged, TransactionResultSwitchInoperativeAdvice, TransactionResultUnauthorizedMerchant, TransactionResultUnknownHostTimeout, TransactionResultUserTransactionLimit:
		return true
	}
	return false
}

// Status of the transaction.
type TransactionStatus string

const (
	TransactionStatusDeclined TransactionStatus = "DECLINED"
	TransactionStatusExpired  TransactionStatus = "EXPIRED"
	TransactionStatusPending  TransactionStatus = "PENDING"
	TransactionStatusSettled  TransactionStatus = "SETTLED"
	TransactionStatusVoided   TransactionStatus = "VOIDED"
)

func (r TransactionStatus) IsKnown() bool {
	switch r {
	case TransactionStatusDeclined, TransactionStatusExpired, TransactionStatusPending, TransactionStatusSettled, TransactionStatusVoided:
		return true
	}
	return false
}

type TransactionTokenInfo struct {
	// The wallet_type field will indicate the source of the token. Possible token
	// sources include digital wallets (Apple, Google, or Samsung Pay), merchant
	// tokenization, and “other” sources like in-flight commerce. Masterpass is not
	// currently supported and is included for future use.
	WalletType TransactionTokenInfoWalletType `json:"wallet_type,required"`
	JSON       transactionTokenInfoJSON       `json:"-"`
}

// transactionTokenInfoJSON contains the JSON metadata for the struct
// [TransactionTokenInfo]
type transactionTokenInfoJSON struct {
	WalletType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionTokenInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionTokenInfoJSON) RawJSON() string {
	return r.raw
}

// The wallet_type field will indicate the source of the token. Possible token
// sources include digital wallets (Apple, Google, or Samsung Pay), merchant
// tokenization, and “other” sources like in-flight commerce. Masterpass is not
// currently supported and is included for future use.
type TransactionTokenInfoWalletType string

const (
	TransactionTokenInfoWalletTypeApplePay   TransactionTokenInfoWalletType = "APPLE_PAY"
	TransactionTokenInfoWalletTypeGooglePay  TransactionTokenInfoWalletType = "GOOGLE_PAY"
	TransactionTokenInfoWalletTypeMasterpass TransactionTokenInfoWalletType = "MASTERPASS"
	TransactionTokenInfoWalletTypeMerchant   TransactionTokenInfoWalletType = "MERCHANT"
	TransactionTokenInfoWalletTypeOther      TransactionTokenInfoWalletType = "OTHER"
	TransactionTokenInfoWalletTypeSamsungPay TransactionTokenInfoWalletType = "SAMSUNG_PAY"
)

func (r TransactionTokenInfoWalletType) IsKnown() bool {
	switch r {
	case TransactionTokenInfoWalletTypeApplePay, TransactionTokenInfoWalletTypeGooglePay, TransactionTokenInfoWalletTypeMasterpass, TransactionTokenInfoWalletTypeMerchant, TransactionTokenInfoWalletTypeOther, TransactionTokenInfoWalletTypeSamsungPay:
		return true
	}
	return false
}

type TransactionEvent struct {
	// Transaction event identifier.
	Token string `json:"token,required" format:"uuid"`
	// Amount of the transaction event (in cents), including any acquirer fees.
	Amount  int64                    `json:"amount,required"`
	Amounts TransactionEventsAmounts `json:"amounts,required"`
	// RFC 3339 date and time this event entered the system. UTC time zone.
	Created         time.Time                         `json:"created,required" format:"date-time"`
	DetailedResults []TransactionEventsDetailedResult `json:"detailed_results,required"`
	// Indicates whether the transaction event is a credit or debit to the account.
	EffectivePolarity TransactionEventsEffectivePolarity `json:"effective_polarity,required"`
	Result            TransactionEventsResult            `json:"result,required"`
	// Type of transaction event
	Type TransactionEventsType `json:"type,required"`
	JSON transactionEventJSON  `json:"-"`
}

// transactionEventJSON contains the JSON metadata for the struct
// [TransactionEvent]
type transactionEventJSON struct {
	Token             apijson.Field
	Amount            apijson.Field
	Amounts           apijson.Field
	Created           apijson.Field
	DetailedResults   apijson.Field
	EffectivePolarity apijson.Field
	Result            apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionEventJSON) RawJSON() string {
	return r.raw
}

type TransactionEventsAmounts struct {
	Cardholder TransactionEventsAmountsCardholder `json:"cardholder,required"`
	Merchant   TransactionEventsAmountsMerchant   `json:"merchant,required"`
	Settlement TransactionEventsAmountsSettlement `json:"settlement,required,nullable"`
	JSON       transactionEventsAmountsJSON       `json:"-"`
}

// transactionEventsAmountsJSON contains the JSON metadata for the struct
// [TransactionEventsAmounts]
type transactionEventsAmountsJSON struct {
	Cardholder  apijson.Field
	Merchant    apijson.Field
	Settlement  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionEventsAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionEventsAmountsJSON) RawJSON() string {
	return r.raw
}

type TransactionEventsAmountsCardholder struct {
	// The amount in the cardholder's local currency.
	Amount int64 `json:"amount,required"`
	// The conversion rate used to convert the merchant amount to the cardholder
	// amount.
	ConversionRate string `json:"conversion_rate,required"`
	// ISO 4217 currency. Its enumerants are ISO 4217 currencies except for some
	// special currencies like `XXX`. Enumerants names are lowercase currency code e.g.
	// `EUR`, `USD`.
	Currency shared.Currency                        `json:"currency,required"`
	JSON     transactionEventsAmountsCardholderJSON `json:"-"`
}

// transactionEventsAmountsCardholderJSON contains the JSON metadata for the struct
// [TransactionEventsAmountsCardholder]
type transactionEventsAmountsCardholderJSON struct {
	Amount         apijson.Field
	ConversionRate apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionEventsAmountsCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionEventsAmountsCardholderJSON) RawJSON() string {
	return r.raw
}

type TransactionEventsAmountsMerchant struct {
	// The amount in the merchant's local currency.
	Amount int64 `json:"amount,required"`
	// ISO 4217 currency. Its enumerants are ISO 4217 currencies except for some
	// special currencies like `XXX`. Enumerants names are lowercase currency code e.g.
	// `EUR`, `USD`.
	Currency shared.Currency                      `json:"currency,required"`
	JSON     transactionEventsAmountsMerchantJSON `json:"-"`
}

// transactionEventsAmountsMerchantJSON contains the JSON metadata for the struct
// [TransactionEventsAmountsMerchant]
type transactionEventsAmountsMerchantJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionEventsAmountsMerchant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionEventsAmountsMerchantJSON) RawJSON() string {
	return r.raw
}

type TransactionEventsAmountsSettlement struct {
	// The amount in the settlement currency.
	Amount int64 `json:"amount,required"`
	// Conversion rate used to convert the merchant amount to the settlement amount.
	ConversionRate string `json:"conversion_rate,required"`
	// ISO 4217 currency. Its enumerants are ISO 4217 currencies except for some
	// special currencies like `XXX`. Enumerants names are lowercase currency code e.g.
	// `EUR`, `USD`.
	Currency shared.Currency                        `json:"currency,required"`
	JSON     transactionEventsAmountsSettlementJSON `json:"-"`
}

// transactionEventsAmountsSettlementJSON contains the JSON metadata for the struct
// [TransactionEventsAmountsSettlement]
type transactionEventsAmountsSettlementJSON struct {
	Amount         apijson.Field
	ConversionRate apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionEventsAmountsSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionEventsAmountsSettlementJSON) RawJSON() string {
	return r.raw
}

type TransactionEventsDetailedResult string

const (
	TransactionEventsDetailedResultAccountDailySpendLimitExceeded              TransactionEventsDetailedResult = "ACCOUNT_DAILY_SPEND_LIMIT_EXCEEDED"
	TransactionEventsDetailedResultAccountDelinquent                           TransactionEventsDetailedResult = "ACCOUNT_DELINQUENT"
	TransactionEventsDetailedResultAccountInactive                             TransactionEventsDetailedResult = "ACCOUNT_INACTIVE"
	TransactionEventsDetailedResultAccountLifetimeSpendLimitExceeded           TransactionEventsDetailedResult = "ACCOUNT_LIFETIME_SPEND_LIMIT_EXCEEDED"
	TransactionEventsDetailedResultAccountMonthlySpendLimitExceeded            TransactionEventsDetailedResult = "ACCOUNT_MONTHLY_SPEND_LIMIT_EXCEEDED"
	TransactionEventsDetailedResultAccountUnderReview                          TransactionEventsDetailedResult = "ACCOUNT_UNDER_REVIEW"
	TransactionEventsDetailedResultAddressIncorrect                            TransactionEventsDetailedResult = "ADDRESS_INCORRECT"
	TransactionEventsDetailedResultApproved                                    TransactionEventsDetailedResult = "APPROVED"
	TransactionEventsDetailedResultAuthRuleAllowedCountry                      TransactionEventsDetailedResult = "AUTH_RULE_ALLOWED_COUNTRY"
	TransactionEventsDetailedResultAuthRuleAllowedMcc                          TransactionEventsDetailedResult = "AUTH_RULE_ALLOWED_MCC"
	TransactionEventsDetailedResultAuthRuleBlockedCountry                      TransactionEventsDetailedResult = "AUTH_RULE_BLOCKED_COUNTRY"
	TransactionEventsDetailedResultAuthRuleBlockedMcc                          TransactionEventsDetailedResult = "AUTH_RULE_BLOCKED_MCC"
	TransactionEventsDetailedResultCardClosed                                  TransactionEventsDetailedResult = "CARD_CLOSED"
	TransactionEventsDetailedResultCardCryptogramValidationFailure             TransactionEventsDetailedResult = "CARD_CRYPTOGRAM_VALIDATION_FAILURE"
	TransactionEventsDetailedResultCardExpired                                 TransactionEventsDetailedResult = "CARD_EXPIRED"
	TransactionEventsDetailedResultCardExpiryDateIncorrect                     TransactionEventsDetailedResult = "CARD_EXPIRY_DATE_INCORRECT"
	TransactionEventsDetailedResultCardInvalid                                 TransactionEventsDetailedResult = "CARD_INVALID"
	TransactionEventsDetailedResultCardNotActivated                            TransactionEventsDetailedResult = "CARD_NOT_ACTIVATED"
	TransactionEventsDetailedResultCardPaused                                  TransactionEventsDetailedResult = "CARD_PAUSED"
	TransactionEventsDetailedResultCardPinIncorrect                            TransactionEventsDetailedResult = "CARD_PIN_INCORRECT"
	TransactionEventsDetailedResultCardRestricted                              TransactionEventsDetailedResult = "CARD_RESTRICTED"
	TransactionEventsDetailedResultCardSecurityCodeIncorrect                   TransactionEventsDetailedResult = "CARD_SECURITY_CODE_INCORRECT"
	TransactionEventsDetailedResultCardSpendLimitExceeded                      TransactionEventsDetailedResult = "CARD_SPEND_LIMIT_EXCEEDED"
	TransactionEventsDetailedResultContactCardIssuer                           TransactionEventsDetailedResult = "CONTACT_CARD_ISSUER"
	TransactionEventsDetailedResultCustomerAsaTimeout                          TransactionEventsDetailedResult = "CUSTOMER_ASA_TIMEOUT"
	TransactionEventsDetailedResultCustomAsaResult                             TransactionEventsDetailedResult = "CUSTOM_ASA_RESULT"
	TransactionEventsDetailedResultDeclined                                    TransactionEventsDetailedResult = "DECLINED"
	TransactionEventsDetailedResultDoNotHonor                                  TransactionEventsDetailedResult = "DO_NOT_HONOR"
	TransactionEventsDetailedResultDriverNumberInvalid                         TransactionEventsDetailedResult = "DRIVER_NUMBER_INVALID"
	TransactionEventsDetailedResultFormatError                                 TransactionEventsDetailedResult = "FORMAT_ERROR"
	TransactionEventsDetailedResultInsufficientFundingSourceBalance            TransactionEventsDetailedResult = "INSUFFICIENT_FUNDING_SOURCE_BALANCE"
	TransactionEventsDetailedResultInsufficientFunds                           TransactionEventsDetailedResult = "INSUFFICIENT_FUNDS"
	TransactionEventsDetailedResultLithicSystemError                           TransactionEventsDetailedResult = "LITHIC_SYSTEM_ERROR"
	TransactionEventsDetailedResultLithicSystemRateLimit                       TransactionEventsDetailedResult = "LITHIC_SYSTEM_RATE_LIMIT"
	TransactionEventsDetailedResultMalformedAsaResponse                        TransactionEventsDetailedResult = "MALFORMED_ASA_RESPONSE"
	TransactionEventsDetailedResultMerchantInvalid                             TransactionEventsDetailedResult = "MERCHANT_INVALID"
	TransactionEventsDetailedResultMerchantLockedCardAttemptedElsewhere        TransactionEventsDetailedResult = "MERCHANT_LOCKED_CARD_ATTEMPTED_ELSEWHERE"
	TransactionEventsDetailedResultMerchantNotPermitted                        TransactionEventsDetailedResult = "MERCHANT_NOT_PERMITTED"
	TransactionEventsDetailedResultOverReversalAttempted                       TransactionEventsDetailedResult = "OVER_REVERSAL_ATTEMPTED"
	TransactionEventsDetailedResultPinBlocked                                  TransactionEventsDetailedResult = "PIN_BLOCKED"
	TransactionEventsDetailedResultProgramCardSpendLimitExceeded               TransactionEventsDetailedResult = "PROGRAM_CARD_SPEND_LIMIT_EXCEEDED"
	TransactionEventsDetailedResultProgramSuspended                            TransactionEventsDetailedResult = "PROGRAM_SUSPENDED"
	TransactionEventsDetailedResultProgramUsageRestriction                     TransactionEventsDetailedResult = "PROGRAM_USAGE_RESTRICTION"
	TransactionEventsDetailedResultReversalUnmatched                           TransactionEventsDetailedResult = "REVERSAL_UNMATCHED"
	TransactionEventsDetailedResultSecurityViolation                           TransactionEventsDetailedResult = "SECURITY_VIOLATION"
	TransactionEventsDetailedResultSingleUseCardReattempted                    TransactionEventsDetailedResult = "SINGLE_USE_CARD_REATTEMPTED"
	TransactionEventsDetailedResultTransactionInvalid                          TransactionEventsDetailedResult = "TRANSACTION_INVALID"
	TransactionEventsDetailedResultTransactionNotPermittedToAcquirerOrTerminal TransactionEventsDetailedResult = "TRANSACTION_NOT_PERMITTED_TO_ACQUIRER_OR_TERMINAL"
	TransactionEventsDetailedResultTransactionNotPermittedToIssuerOrCardholder TransactionEventsDetailedResult = "TRANSACTION_NOT_PERMITTED_TO_ISSUER_OR_CARDHOLDER"
	TransactionEventsDetailedResultTransactionPreviouslyCompleted              TransactionEventsDetailedResult = "TRANSACTION_PREVIOUSLY_COMPLETED"
	TransactionEventsDetailedResultUnauthorizedMerchant                        TransactionEventsDetailedResult = "UNAUTHORIZED_MERCHANT"
	TransactionEventsDetailedResultVehicleNumberInvalid                        TransactionEventsDetailedResult = "VEHICLE_NUMBER_INVALID"
)

func (r TransactionEventsDetailedResult) IsKnown() bool {
	switch r {
	case TransactionEventsDetailedResultAccountDailySpendLimitExceeded, TransactionEventsDetailedResultAccountDelinquent, TransactionEventsDetailedResultAccountInactive, TransactionEventsDetailedResultAccountLifetimeSpendLimitExceeded, TransactionEventsDetailedResultAccountMonthlySpendLimitExceeded, TransactionEventsDetailedResultAccountUnderReview, TransactionEventsDetailedResultAddressIncorrect, TransactionEventsDetailedResultApproved, TransactionEventsDetailedResultAuthRuleAllowedCountry, TransactionEventsDetailedResultAuthRuleAllowedMcc, TransactionEventsDetailedResultAuthRuleBlockedCountry, TransactionEventsDetailedResultAuthRuleBlockedMcc, TransactionEventsDetailedResultCardClosed, TransactionEventsDetailedResultCardCryptogramValidationFailure, TransactionEventsDetailedResultCardExpired, TransactionEventsDetailedResultCardExpiryDateIncorrect, TransactionEventsDetailedResultCardInvalid, TransactionEventsDetailedResultCardNotActivated, TransactionEventsDetailedResultCardPaused, TransactionEventsDetailedResultCardPinIncorrect, TransactionEventsDetailedResultCardRestricted, TransactionEventsDetailedResultCardSecurityCodeIncorrect, TransactionEventsDetailedResultCardSpendLimitExceeded, TransactionEventsDetailedResultContactCardIssuer, TransactionEventsDetailedResultCustomerAsaTimeout, TransactionEventsDetailedResultCustomAsaResult, TransactionEventsDetailedResultDeclined, TransactionEventsDetailedResultDoNotHonor, TransactionEventsDetailedResultDriverNumberInvalid, TransactionEventsDetailedResultFormatError, TransactionEventsDetailedResultInsufficientFundingSourceBalance, TransactionEventsDetailedResultInsufficientFunds, TransactionEventsDetailedResultLithicSystemError, TransactionEventsDetailedResultLithicSystemRateLimit, TransactionEventsDetailedResultMalformedAsaResponse, TransactionEventsDetailedResultMerchantInvalid, TransactionEventsDetailedResultMerchantLockedCardAttemptedElsewhere, TransactionEventsDetailedResultMerchantNotPermitted, TransactionEventsDetailedResultOverReversalAttempted, TransactionEventsDetailedResultPinBlocked, TransactionEventsDetailedResultProgramCardSpendLimitExceeded, TransactionEventsDetailedResultProgramSuspended, TransactionEventsDetailedResultProgramUsageRestriction, TransactionEventsDetailedResultReversalUnmatched, TransactionEventsDetailedResultSecurityViolation, TransactionEventsDetailedResultSingleUseCardReattempted, TransactionEventsDetailedResultTransactionInvalid, TransactionEventsDetailedResultTransactionNotPermittedToAcquirerOrTerminal, TransactionEventsDetailedResultTransactionNotPermittedToIssuerOrCardholder, TransactionEventsDetailedResultTransactionPreviouslyCompleted, TransactionEventsDetailedResultUnauthorizedMerchant, TransactionEventsDetailedResultVehicleNumberInvalid:
		return true
	}
	return false
}

// Indicates whether the transaction event is a credit or debit to the account.
type TransactionEventsEffectivePolarity string

const (
	TransactionEventsEffectivePolarityCredit TransactionEventsEffectivePolarity = "CREDIT"
	TransactionEventsEffectivePolarityDebit  TransactionEventsEffectivePolarity = "DEBIT"
)

func (r TransactionEventsEffectivePolarity) IsKnown() bool {
	switch r {
	case TransactionEventsEffectivePolarityCredit, TransactionEventsEffectivePolarityDebit:
		return true
	}
	return false
}

type TransactionEventsResult string

const (
	TransactionEventsResultAccountStateTransactionFail TransactionEventsResult = "ACCOUNT_STATE_TRANSACTION_FAIL"
	TransactionEventsResultApproved                    TransactionEventsResult = "APPROVED"
	TransactionEventsResultBankConnectionError         TransactionEventsResult = "BANK_CONNECTION_ERROR"
	TransactionEventsResultBankNotVerified             TransactionEventsResult = "BANK_NOT_VERIFIED"
	TransactionEventsResultCardClosed                  TransactionEventsResult = "CARD_CLOSED"
	TransactionEventsResultCardPaused                  TransactionEventsResult = "CARD_PAUSED"
	TransactionEventsResultDeclined                    TransactionEventsResult = "DECLINED"
	TransactionEventsResultFraudAdvice                 TransactionEventsResult = "FRAUD_ADVICE"
	TransactionEventsResultIgnoredTtlExpiry            TransactionEventsResult = "IGNORED_TTL_EXPIRY"
	TransactionEventsResultInactiveAccount             TransactionEventsResult = "INACTIVE_ACCOUNT"
	TransactionEventsResultIncorrectPin                TransactionEventsResult = "INCORRECT_PIN"
	TransactionEventsResultInvalidCardDetails          TransactionEventsResult = "INVALID_CARD_DETAILS"
	TransactionEventsResultInsufficientFunds           TransactionEventsResult = "INSUFFICIENT_FUNDS"
	TransactionEventsResultInsufficientFundsPreload    TransactionEventsResult = "INSUFFICIENT_FUNDS_PRELOAD"
	TransactionEventsResultInvalidTransaction          TransactionEventsResult = "INVALID_TRANSACTION"
	TransactionEventsResultMerchantBlacklist           TransactionEventsResult = "MERCHANT_BLACKLIST"
	TransactionEventsResultOriginalNotFound            TransactionEventsResult = "ORIGINAL_NOT_FOUND"
	TransactionEventsResultPreviouslyCompleted         TransactionEventsResult = "PREVIOUSLY_COMPLETED"
	TransactionEventsResultSingleUseRecharged          TransactionEventsResult = "SINGLE_USE_RECHARGED"
	TransactionEventsResultSwitchInoperativeAdvice     TransactionEventsResult = "SWITCH_INOPERATIVE_ADVICE"
	TransactionEventsResultUnauthorizedMerchant        TransactionEventsResult = "UNAUTHORIZED_MERCHANT"
	TransactionEventsResultUnknownHostTimeout          TransactionEventsResult = "UNKNOWN_HOST_TIMEOUT"
	TransactionEventsResultUserTransactionLimit        TransactionEventsResult = "USER_TRANSACTION_LIMIT"
)

func (r TransactionEventsResult) IsKnown() bool {
	switch r {
	case TransactionEventsResultAccountStateTransactionFail, TransactionEventsResultApproved, TransactionEventsResultBankConnectionError, TransactionEventsResultBankNotVerified, TransactionEventsResultCardClosed, TransactionEventsResultCardPaused, TransactionEventsResultDeclined, TransactionEventsResultFraudAdvice, TransactionEventsResultIgnoredTtlExpiry, TransactionEventsResultInactiveAccount, TransactionEventsResultIncorrectPin, TransactionEventsResultInvalidCardDetails, TransactionEventsResultInsufficientFunds, TransactionEventsResultInsufficientFundsPreload, TransactionEventsResultInvalidTransaction, TransactionEventsResultMerchantBlacklist, TransactionEventsResultOriginalNotFound, TransactionEventsResultPreviouslyCompleted, TransactionEventsResultSingleUseRecharged, TransactionEventsResultSwitchInoperativeAdvice, TransactionEventsResultUnauthorizedMerchant, TransactionEventsResultUnknownHostTimeout, TransactionEventsResultUserTransactionLimit:
		return true
	}
	return false
}

// Type of transaction event
type TransactionEventsType string

const (
	TransactionEventsTypeAuthorization                TransactionEventsType = "AUTHORIZATION"
	TransactionEventsTypeAuthorizationAdvice          TransactionEventsType = "AUTHORIZATION_ADVICE"
	TransactionEventsTypeAuthorizationExpiry          TransactionEventsType = "AUTHORIZATION_EXPIRY"
	TransactionEventsTypeAuthorizationReversal        TransactionEventsType = "AUTHORIZATION_REVERSAL"
	TransactionEventsTypeBalanceInquiry               TransactionEventsType = "BALANCE_INQUIRY"
	TransactionEventsTypeClearing                     TransactionEventsType = "CLEARING"
	TransactionEventsTypeCorrectionCredit             TransactionEventsType = "CORRECTION_CREDIT"
	TransactionEventsTypeCorrectionDebit              TransactionEventsType = "CORRECTION_DEBIT"
	TransactionEventsTypeCreditAuthorization          TransactionEventsType = "CREDIT_AUTHORIZATION"
	TransactionEventsTypeCreditAuthorizationAdvice    TransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	TransactionEventsTypeFinancialAuthorization       TransactionEventsType = "FINANCIAL_AUTHORIZATION"
	TransactionEventsTypeFinancialCreditAuthorization TransactionEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	TransactionEventsTypeReturn                       TransactionEventsType = "RETURN"
	TransactionEventsTypeReturnReversal               TransactionEventsType = "RETURN_REVERSAL"
)

func (r TransactionEventsType) IsKnown() bool {
	switch r {
	case TransactionEventsTypeAuthorization, TransactionEventsTypeAuthorizationAdvice, TransactionEventsTypeAuthorizationExpiry, TransactionEventsTypeAuthorizationReversal, TransactionEventsTypeBalanceInquiry, TransactionEventsTypeClearing, TransactionEventsTypeCorrectionCredit, TransactionEventsTypeCorrectionDebit, TransactionEventsTypeCreditAuthorization, TransactionEventsTypeCreditAuthorizationAdvice, TransactionEventsTypeFinancialAuthorization, TransactionEventsTypeFinancialCreditAuthorization, TransactionEventsTypeReturn, TransactionEventsTypeReturnReversal:
		return true
	}
	return false
}

type TransactionSimulateAuthorizationResponse struct {
	// A unique token to reference this transaction with later calls to void or clear
	// the authorization.
	Token string `json:"token" format:"uuid"`
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string                                       `json:"debugging_request_id" format:"uuid"`
	JSON               transactionSimulateAuthorizationResponseJSON `json:"-"`
}

// transactionSimulateAuthorizationResponseJSON contains the JSON metadata for the
// struct [TransactionSimulateAuthorizationResponse]
type transactionSimulateAuthorizationResponseJSON struct {
	Token              apijson.Field
	DebuggingRequestID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSimulateAuthorizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulateAuthorizationResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulateAuthorizationAdviceResponse struct {
	// A unique token to reference this transaction.
	Token string `json:"token" format:"uuid"`
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string                                             `json:"debugging_request_id" format:"uuid"`
	JSON               transactionSimulateAuthorizationAdviceResponseJSON `json:"-"`
}

// transactionSimulateAuthorizationAdviceResponseJSON contains the JSON metadata
// for the struct [TransactionSimulateAuthorizationAdviceResponse]
type transactionSimulateAuthorizationAdviceResponseJSON struct {
	Token              apijson.Field
	DebuggingRequestID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSimulateAuthorizationAdviceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulateAuthorizationAdviceResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulateClearingResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string                                  `json:"debugging_request_id" format:"uuid"`
	JSON               transactionSimulateClearingResponseJSON `json:"-"`
}

// transactionSimulateClearingResponseJSON contains the JSON metadata for the
// struct [TransactionSimulateClearingResponse]
type transactionSimulateClearingResponseJSON struct {
	DebuggingRequestID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSimulateClearingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulateClearingResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulateCreditAuthorizationResponse struct {
	// A unique token to reference this transaction.
	Token string `json:"token" format:"uuid"`
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string                                             `json:"debugging_request_id" format:"uuid"`
	JSON               transactionSimulateCreditAuthorizationResponseJSON `json:"-"`
}

// transactionSimulateCreditAuthorizationResponseJSON contains the JSON metadata
// for the struct [TransactionSimulateCreditAuthorizationResponse]
type transactionSimulateCreditAuthorizationResponseJSON struct {
	Token              apijson.Field
	DebuggingRequestID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSimulateCreditAuthorizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulateCreditAuthorizationResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulateReturnResponse struct {
	// A unique token to reference this transaction.
	Token string `json:"token" format:"uuid"`
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string                                `json:"debugging_request_id" format:"uuid"`
	JSON               transactionSimulateReturnResponseJSON `json:"-"`
}

// transactionSimulateReturnResponseJSON contains the JSON metadata for the struct
// [TransactionSimulateReturnResponse]
type transactionSimulateReturnResponseJSON struct {
	Token              apijson.Field
	DebuggingRequestID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSimulateReturnResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulateReturnResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulateReturnReversalResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string                                        `json:"debugging_request_id" format:"uuid"`
	JSON               transactionSimulateReturnReversalResponseJSON `json:"-"`
}

// transactionSimulateReturnReversalResponseJSON contains the JSON metadata for the
// struct [TransactionSimulateReturnReversalResponse]
type transactionSimulateReturnReversalResponseJSON struct {
	DebuggingRequestID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSimulateReturnReversalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulateReturnReversalResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulateVoidResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string                              `json:"debugging_request_id" format:"uuid"`
	JSON               transactionSimulateVoidResponseJSON `json:"-"`
}

// transactionSimulateVoidResponseJSON contains the JSON metadata for the struct
// [TransactionSimulateVoidResponse]
type transactionSimulateVoidResponseJSON struct {
	DebuggingRequestID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSimulateVoidResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulateVoidResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionListParams struct {
	// Filters for transactions associated with a specific account.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Filters for transactions associated with a specific card.
	CardToken param.Field[string] `query:"card_token" format:"uuid"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// Filters for transactions using transaction result field. Can filter by
	// `APPROVED`, and `DECLINED`.
	Result param.Field[TransactionListParamsResult] `query:"result"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filters for transactions using transaction result field. Can filter by
// `APPROVED`, and `DECLINED`.
type TransactionListParamsResult string

const (
	TransactionListParamsResultApproved TransactionListParamsResult = "APPROVED"
	TransactionListParamsResultDeclined TransactionListParamsResult = "DECLINED"
)

func (r TransactionListParamsResult) IsKnown() bool {
	switch r {
	case TransactionListParamsResultApproved, TransactionListParamsResultDeclined:
		return true
	}
	return false
}

type TransactionSimulateAuthorizationParams struct {
	// Amount (in cents) to authorize. For credit authorizations and financial credit
	// authorizations, any value entered will be converted into a negative amount in
	// the simulated transaction. For example, entering 100 in this field will appear
	// as a -100 amount in the transaction. For balance inquiries, this field must be
	// set to 0.
	Amount param.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor param.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan param.Field[string] `json:"pan,required"`
	// Merchant category code for the transaction to be simulated. A four-digit number
	// listed in ISO 18245. Supported merchant category codes can be found
	// [here](https://docs.lithic.com/docs/transactions#merchant-category-codes-mccs).
	Mcc param.Field[string] `json:"mcc"`
	// Unique identifier to identify the payment card acceptor.
	MerchantAcceptorID param.Field[string] `json:"merchant_acceptor_id"`
	// Amount of the transaction to be simulated in currency specified in
	// merchant_currency, including any acquirer fees.
	MerchantAmount param.Field[int64] `json:"merchant_amount"`
	// 3-digit alphabetic ISO 4217 currency code. Note: Simulator only accepts USD,
	// GBP, EUR and defaults to GBP if another ISO 4217 code is provided
	MerchantCurrency param.Field[string] `json:"merchant_currency"`
	// Set to true if the terminal is capable of partial approval otherwise false.
	// Partial approval is when part of a transaction is approved and another payment
	// must be used for the remainder.
	PartialApprovalCapable param.Field[bool] `json:"partial_approval_capable"`
	// Type of event to simulate.
	//
	//   - `AUTHORIZATION` is a dual message purchase authorization, meaning a subsequent
	//     clearing step is required to settle the transaction.
	//   - `BALANCE_INQUIRY` is a $0 authorization that includes a request for the
	//     balance held on the card, and is most typically seen when a cardholder
	//     requests to view a card's balance at an ATM.
	//   - `CREDIT_AUTHORIZATION` is a dual message request from a merchant to authorize
	//     a refund or credit, meaning a subsequent clearing step is required to settle
	//     the transaction.
	//   - `FINANCIAL_AUTHORIZATION` is a single message request from a merchant to debit
	//     funds immediately (such as an ATM withdrawal), and no subsequent clearing is
	//     required to settle the transaction.
	//   - `FINANCIAL_CREDIT_AUTHORIZATION` is a single message request from a merchant
	//     to credit funds immediately, and no subsequent clearing is required to settle
	//     the transaction.
	Status param.Field[TransactionSimulateAuthorizationParamsStatus] `json:"status"`
}

func (r TransactionSimulateAuthorizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of event to simulate.
//
//   - `AUTHORIZATION` is a dual message purchase authorization, meaning a subsequent
//     clearing step is required to settle the transaction.
//   - `BALANCE_INQUIRY` is a $0 authorization that includes a request for the
//     balance held on the card, and is most typically seen when a cardholder
//     requests to view a card's balance at an ATM.
//   - `CREDIT_AUTHORIZATION` is a dual message request from a merchant to authorize
//     a refund or credit, meaning a subsequent clearing step is required to settle
//     the transaction.
//   - `FINANCIAL_AUTHORIZATION` is a single message request from a merchant to debit
//     funds immediately (such as an ATM withdrawal), and no subsequent clearing is
//     required to settle the transaction.
//   - `FINANCIAL_CREDIT_AUTHORIZATION` is a single message request from a merchant
//     to credit funds immediately, and no subsequent clearing is required to settle
//     the transaction.
type TransactionSimulateAuthorizationParamsStatus string

const (
	TransactionSimulateAuthorizationParamsStatusAuthorization                TransactionSimulateAuthorizationParamsStatus = "AUTHORIZATION"
	TransactionSimulateAuthorizationParamsStatusBalanceInquiry               TransactionSimulateAuthorizationParamsStatus = "BALANCE_INQUIRY"
	TransactionSimulateAuthorizationParamsStatusCreditAuthorization          TransactionSimulateAuthorizationParamsStatus = "CREDIT_AUTHORIZATION"
	TransactionSimulateAuthorizationParamsStatusFinancialAuthorization       TransactionSimulateAuthorizationParamsStatus = "FINANCIAL_AUTHORIZATION"
	TransactionSimulateAuthorizationParamsStatusFinancialCreditAuthorization TransactionSimulateAuthorizationParamsStatus = "FINANCIAL_CREDIT_AUTHORIZATION"
)

func (r TransactionSimulateAuthorizationParamsStatus) IsKnown() bool {
	switch r {
	case TransactionSimulateAuthorizationParamsStatusAuthorization, TransactionSimulateAuthorizationParamsStatusBalanceInquiry, TransactionSimulateAuthorizationParamsStatusCreditAuthorization, TransactionSimulateAuthorizationParamsStatusFinancialAuthorization, TransactionSimulateAuthorizationParamsStatusFinancialCreditAuthorization:
		return true
	}
	return false
}

type TransactionSimulateAuthorizationAdviceParams struct {
	// The transaction token returned from the /v1/simulate/authorize response.
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// Amount (in cents) to authorize. This amount will override the transaction's
	// amount that was originally set by /v1/simulate/authorize.
	Amount param.Field[int64] `json:"amount,required"`
}

func (r TransactionSimulateAuthorizationAdviceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateClearingParams struct {
	// The transaction token returned from the /v1/simulate/authorize response.
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// Amount (in cents) to complete. Typically this will match the original
	// authorization, but may be more or less.
	//
	// If no amount is supplied to this endpoint, the amount of the transaction will be
	// captured. Any transaction that has any amount completed at all do not have
	// access to this behavior.
	Amount param.Field[int64] `json:"amount"`
}

func (r TransactionSimulateClearingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateCreditAuthorizationParams struct {
	// Amount (in cents). Any value entered will be converted into a negative amount in
	// the simulated transaction. For example, entering 100 in this field will appear
	// as a -100 amount in the transaction.
	Amount param.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor param.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan param.Field[string] `json:"pan,required"`
	// Merchant category code for the transaction to be simulated. A four-digit number
	// listed in ISO 18245. Supported merchant category codes can be found
	// [here](https://docs.lithic.com/docs/transactions#merchant-category-codes-mccs).
	Mcc param.Field[string] `json:"mcc"`
	// Unique identifier to identify the payment card acceptor.
	MerchantAcceptorID param.Field[string] `json:"merchant_acceptor_id"`
}

func (r TransactionSimulateCreditAuthorizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateReturnParams struct {
	// Amount (in cents) to authorize.
	Amount param.Field[int64] `json:"amount,required"`
	// Merchant descriptor.
	Descriptor param.Field[string] `json:"descriptor,required"`
	// Sixteen digit card number.
	Pan param.Field[string] `json:"pan,required"`
}

func (r TransactionSimulateReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateReturnReversalParams struct {
	// The transaction token returned from the /v1/simulate/authorize response.
	Token param.Field[string] `json:"token,required" format:"uuid"`
}

func (r TransactionSimulateReturnReversalParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSimulateVoidParams struct {
	// The transaction token returned from the /v1/simulate/authorize response.
	Token param.Field[string] `json:"token,required" format:"uuid"`
	// Amount (in cents) to void. Typically this will match the original authorization,
	// but may be less.
	Amount param.Field[int64] `json:"amount"`
	// Type of event to simulate. Defaults to `AUTHORIZATION_REVERSAL`.
	//
	//   - `AUTHORIZATION_EXPIRY` indicates authorization has expired and been reversed
	//     by Lithic.
	//   - `AUTHORIZATION_REVERSAL` indicates authorization was reversed by the merchant.
	Type param.Field[TransactionSimulateVoidParamsType] `json:"type"`
}

func (r TransactionSimulateVoidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of event to simulate. Defaults to `AUTHORIZATION_REVERSAL`.
//
//   - `AUTHORIZATION_EXPIRY` indicates authorization has expired and been reversed
//     by Lithic.
//   - `AUTHORIZATION_REVERSAL` indicates authorization was reversed by the merchant.
type TransactionSimulateVoidParamsType string

const (
	TransactionSimulateVoidParamsTypeAuthorizationExpiry   TransactionSimulateVoidParamsType = "AUTHORIZATION_EXPIRY"
	TransactionSimulateVoidParamsTypeAuthorizationReversal TransactionSimulateVoidParamsType = "AUTHORIZATION_REVERSAL"
)

func (r TransactionSimulateVoidParamsType) IsKnown() bool {
	switch r {
	case TransactionSimulateVoidParamsTypeAuthorizationExpiry, TransactionSimulateVoidParamsTypeAuthorizationReversal:
		return true
	}
	return false
}
