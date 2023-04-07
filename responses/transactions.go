package responses

import (
	"time"

	pjson "github.com/lithic-com/lithic-go/core/json"
)

type Transaction struct {
	// A fixed-width 23-digit numeric identifier for the Transaction that may be set if
	// the transaction originated from the Mastercard network. This number may be used
	// for dispute tracking.
	AcquirerReferenceNumber string `json:"acquirer_reference_number,nullable"`
	// Authorization amount of the transaction (in cents), including any acquirer fees.
	// This may change over time, and will represent the settled amount once the
	// transaction is settled.
	Amount int64 `json:"amount"`
	// Authorization amount (in cents) of the transaction, including any acquirer fees.
	// This amount always represents the amount authorized for the transaction,
	// unaffected by settlement.
	AuthorizationAmount      int64                               `json:"authorization_amount"`
	CardholderAuthentication TransactionCardholderAuthentication `json:"cardholder_authentication,nullable"`
	// Analogous to the "amount" property, but will represent the amount in the
	// transaction's local currency (smallest unit), including any acquirer fees.
	MerchantAmount int64 `json:"merchant_amount"`
	// Analogous to the "authorization_amount" property, but will represent the amount
	// in the transaction's local currency (smallest unit), including any acquirer
	// fees.
	MerchantAuthorizationAmount int64 `json:"merchant_authorization_amount"`
	// 3-digit alphabetic ISO 4217 code for the local currency of the transaction.
	MerchantCurrency string `json:"merchant_currency"`
	// A fixed-width 6-digit numeric identifier that can be used to identify a
	// transaction with networks.
	AuthorizationCode string `json:"authorization_code"`
	// Token for the card used in this transaction.
	CardToken string `json:"card_token" format:"uuid"`
	// Date and time when the transaction first occurred. UTC time zone.
	Created time.Time `json:"created" format:"date-time"`
	// A list of all events that have modified this transaction.
	Events   []TransactionEvents `json:"events"`
	Merchant TransactionMerchant `json:"merchant"`
	// Card network of the authorization. Can be `INTERLINK`, `MAESTRO`, `MASTERCARD`,
	// `VISA`, or `UNKNOWN`. Value is `UNKNOWN` when Lithic cannot determine the
	// network code from the upstream provider.
	Network TransactionNetwork `json:"network,nullable"`
	// `APPROVED` or decline reason. See Event result types
	Result TransactionResult `json:"result"`
	// Amount of the transaction that has been settled (in cents), including any
	// acquirer fees. This may change over time.
	SettledAmount int64 `json:"settled_amount"`
	// Status types:
	//
	//   - `DECLINED` - The transaction was declined.
	//   - `EXPIRED` - Lithic reversed the authorization as it has passed its expiration
	//     time.
	//   - `PENDING` - Authorization is pending completion from the merchant.
	//   - `SETTLED` - The transaction is complete.
	//   - `VOIDED` - The merchant has voided the previously pending authorization.
	Status TransactionStatus `json:"status"`
	// Globally unique identifier.
	Token string `json:"token" format:"uuid"`
	JSON  TransactionJSON
}

type TransactionJSON struct {
	AcquirerReferenceNumber     pjson.Metadata
	Amount                      pjson.Metadata
	AuthorizationAmount         pjson.Metadata
	CardholderAuthentication    pjson.Metadata
	MerchantAmount              pjson.Metadata
	MerchantAuthorizationAmount pjson.Metadata
	MerchantCurrency            pjson.Metadata
	AuthorizationCode           pjson.Metadata
	CardToken                   pjson.Metadata
	Created                     pjson.Metadata
	Events                      pjson.Metadata
	Merchant                    pjson.Metadata
	Network                     pjson.Metadata
	Result                      pjson.Metadata
	SettledAmount               pjson.Metadata
	Status                      pjson.Metadata
	Token                       pjson.Metadata
	Raw                         []byte
	Extras                      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Transaction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionCardholderAuthentication struct {
	// 3-D Secure Protocol version. Possible values:
	//
	// - `1`: 3-D Secure Protocol version 1.x applied to the transaction.
	// - `2`: 3-D Secure Protocol version 2.x applied to the transaction.
	// - `null`: 3-D Secure was not used for the transaction
	_3dsVersion string `json:"3ds_version,required,nullable"`
	// Exemption applied by the ACS to authenticate the transaction without requesting
	// a challenge. Possible values:
	//
	//   - `AUTHENTICATION_OUTAGE_EXCEPTION`: Authentication Outage Exception exemption.
	//   - `LOW_VALUE`: Low Value Payment exemption.
	//   - `MERCHANT_INITIATED_TRANSACTION`: Merchant Initiated Transaction (3RI).
	//   - `NONE`: No exemption applied.
	//   - `RECURRING_PAYMENT`: Recurring Payment exemption.
	//   - `SECURE_CORPORATE_PAYMENT`: Secure Corporate Payment exemption.
	//   - `STRONG_CUSTOMER_AUTHENTICATION_DELEGATION`: Strong Customer Authentication
	//     Delegation exemption.
	//   - `TRANSACTION_RISK_ANALYSIS`: Acquirer Low-Fraud and Transaction Risk Analysis
	//     exemption.
	//
	// Maps to the 3-D Secure `transChallengeExemption` field.
	AcquirerExemption TransactionCardholderAuthenticationAcquirerExemption `json:"acquirer_exemption,required"`
	// Indicates whether chargeback liability shift applies to the transaction.
	// Possible values:
	//
	//   - `3DS_AUTHENTICATED`: The transaction was fully authenticated through a 3-D
	//     Secure flow, chargeback liability shift applies.
	//   - `ACQUIRER_EXEMPTION`: The acquirer utilised an exemption to bypass Strong
	//     Customer Authentication (`transStatus = N`, or `transStatus = I`). Liability
	//     remains with the acquirer and in this case the `acquirer_exemption` field is
	//     expected to be not `NONE`.
	//   - `NONE`: Chargeback liability shift has not shifted to the issuer, i.e. the
	//     merchant is liable.
	//   - `TOKEN_AUTHENTICATED`: The transaction was a tokenized payment with validated
	//     cryptography, possibly recurring. Chargeback liability shift to the issuer
	//     applies.
	LiabilityShift TransactionCardholderAuthenticationLiabilityShift `json:"liability_shift,required"`
	// Verification attempted values:
	//
	//   - `APP_LOGIN`: Out-of-band login verification was attempted by the ACS.
	//   - `BIOMETRIC`: Out-of-band biometric verification was attempted by the ACS.
	//   - `NONE`: No cardholder verification was attempted by the Access Control Server
	//     (e.g. frictionless 3-D Secure flow, no 3-D Secure, or stand-in Risk Based
	//     Analysis).
	//   - `OTHER`: Other method was used by the ACS to verify the cardholder (e.g.
	//     Mastercard Identity Check Express, recurring transactions, etc.)
	//   - `OTP`: One-time password verification was attempted by the ACS.
	VerificationAttempted TransactionCardholderAuthenticationVerificationAttempted `json:"verification_attempted,required"`
	// This field partially maps to the `transStatus` field in the
	// [EMVCo 3-D Secure specification](https://www.emvco.com/emv-technologies/3d-secure/)
	// and Mastercard SPA2 AAV leading indicators.
	//
	// Verification result values:
	//
	//   - `CANCELLED`: Authentication/Account verification could not be performed,
	//     `transStatus = U`.
	//   - `FAILED`: Transaction was not authenticated. `transStatus = N`, note: the
	//     utilization of exemptions could also result in `transStatus = N`, inspect the
	//     `acquirer_exemption` field for more information.
	//   - `FRICTIONLESS`: Attempts processing performed, the transaction was not
	//     authenticated, but a proof of attempted authentication/verification is
	//     provided. `transStatus = A` and the leading AAV indicator was one of {`kE`,
	//     `kF`, `kQ`}.
	//   - `NOT_ATTEMPTED`: A 3-D Secure flow was not applied to this transaction.
	//     Leading AAV indicator was one of {`kN`, `kX`} or no AAV was provided for the
	//     transaction.
	//   - `REJECTED`: Authentication/Account Verification rejected; `transStatus = R`.
	//     Issuer is rejecting authentication/verification and requests that
	//     authorization not be attempted.
	//   - `SUCCESS`: Authentication verification successful. `transStatus = Y` and
	//     leading AAV indicator for the transaction was one of {`kA`, `kB`, `kC`, `kD`,
	//     `kO`, `kP`, `kR`, `kS`}.
	//
	// Note that the following `transStatus` values are not represented by this field:
	//
	// - `C`: Challenge Required
	// - `D`: Challenge Required; decoupled authentication confirmed
	// - `I`: Informational only
	// - `S`: Challenge using Secure Payment Confirmation (SPC)
	VerificationResult TransactionCardholderAuthenticationVerificationResult `json:"verification_result,required"`
	JSON               TransactionCardholderAuthenticationJSON
}

type TransactionCardholderAuthenticationJSON struct {
	_3dsVersion           pjson.Metadata
	AcquirerExemption     pjson.Metadata
	LiabilityShift        pjson.Metadata
	VerificationAttempted pjson.Metadata
	VerificationResult    pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionCardholderAuthentication using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionCardholderAuthentication) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

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

type TransactionCardholderAuthenticationLiabilityShift string

const (
	TransactionCardholderAuthenticationLiabilityShift3DsAuthenticated   TransactionCardholderAuthenticationLiabilityShift = "3DS_AUTHENTICATED"
	TransactionCardholderAuthenticationLiabilityShiftAcquirerExemption  TransactionCardholderAuthenticationLiabilityShift = "ACQUIRER_EXEMPTION"
	TransactionCardholderAuthenticationLiabilityShiftNone               TransactionCardholderAuthenticationLiabilityShift = "NONE"
	TransactionCardholderAuthenticationLiabilityShiftTokenAuthenticated TransactionCardholderAuthenticationLiabilityShift = "TOKEN_AUTHENTICATED"
)

type TransactionCardholderAuthenticationVerificationAttempted string

const (
	TransactionCardholderAuthenticationVerificationAttemptedAppLogin  TransactionCardholderAuthenticationVerificationAttempted = "APP_LOGIN"
	TransactionCardholderAuthenticationVerificationAttemptedBiometric TransactionCardholderAuthenticationVerificationAttempted = "BIOMETRIC"
	TransactionCardholderAuthenticationVerificationAttemptedNone      TransactionCardholderAuthenticationVerificationAttempted = "NONE"
	TransactionCardholderAuthenticationVerificationAttemptedOther     TransactionCardholderAuthenticationVerificationAttempted = "OTHER"
	TransactionCardholderAuthenticationVerificationAttemptedOtp       TransactionCardholderAuthenticationVerificationAttempted = "OTP"
)

type TransactionCardholderAuthenticationVerificationResult string

const (
	TransactionCardholderAuthenticationVerificationResultCancelled    TransactionCardholderAuthenticationVerificationResult = "CANCELLED"
	TransactionCardholderAuthenticationVerificationResultFailed       TransactionCardholderAuthenticationVerificationResult = "FAILED"
	TransactionCardholderAuthenticationVerificationResultFrictionless TransactionCardholderAuthenticationVerificationResult = "FRICTIONLESS"
	TransactionCardholderAuthenticationVerificationResultNotAttempted TransactionCardholderAuthenticationVerificationResult = "NOT_ATTEMPTED"
	TransactionCardholderAuthenticationVerificationResultRejected     TransactionCardholderAuthenticationVerificationResult = "REJECTED"
	TransactionCardholderAuthenticationVerificationResultSuccess      TransactionCardholderAuthenticationVerificationResult = "SUCCESS"
)

type TransactionEvents struct {
	// Amount of the transaction event (in cents), including any acquirer fees.
	Amount int64 `json:"amount,required"`
	// RFC 3339 date and time this event entered the system. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// `APPROVED` or decline reason.
	//
	// Result types:
	//
	//   - `ACCOUNT_STATE_TRANSACTION_FAIL` - Contact
	//     [support@lithic.com](mailto:support@lithic.com).
	//   - `APPROVED` - Transaction is approved.
	//   - `BANK_CONNECTION_ERROR` - Please reconnect a funding source.
	//   - `BANK_NOT_VERIFIED` - Please confirm the funding source.
	//   - `CARD_CLOSED` - Card state was closed at the time of authorization.
	//   - `CARD_PAUSED` - Card state was paused at the time of authorization.
	//   - `FRAUD_ADVICE` - Transaction declined due to risk.
	//   - `GLOBAL_TRANSACTION_LIMIT` - Platform spend limit exceeded, contact
	//     [support@lithic.com](mailto:support@lithic.com).
	//   - `GLOBAL_WEEKLY_LIMIT` - Platform spend limit exceeded, contact
	//     [support@lithic.com](mailto:support@lithic.com).
	//   - `GLOBAL_MONTHLY_LIMIT` - Platform spend limit exceeded, contact
	//     [support@lithic.com](mailto:support@lithic.com).
	//   - `INACTIVE_ACCOUNT` - Account is inactive. Contact
	//     [support@lithic.com](mailto:support@lithic.com).
	//   - `INCORRECT_PIN` - PIN verification failed.
	//   - `INVALID_CARD_DETAILS` - Incorrect CVV or expiry date.
	//   - `INSUFFICIENT_FUNDS` - Please ensure the funding source is connected and up to
	//     date.
	//   - `MERCHANT_BLACKLIST` - This merchant is disallowed on the platform.
	//   - `SINGLE_USE_RECHARGED` - Single use card attempted multiple times.
	//   - `SWITCH_INOPERATIVE_ADVICE` - Network error, re-attempt the transaction.
	//   - `UNAUTHORIZED_MERCHANT` - Merchant locked card attempted at different
	//     merchant.
	//   - `UNKNOWN_HOST_TIMEOUT` - Network error, re-attempt the transaction.
	//   - `USER_TRANSACTION_LIMIT` - User-set spend limit exceeded.
	Result TransactionEventsResult `json:"result,required"`
	// Globally unique identifier.
	Token string `json:"token,required" format:"uuid"`
	// Event types:
	//
	//   - `AUTHORIZATION` - Authorize a transaction.
	//   - `AUTHORIZATION_ADVICE` - Advice on a transaction.
	//   - `AUTHORIZATION_EXPIRY` - Authorization has expired and reversed by Lithic.
	//   - `AUTHORIZATION_REVERSAL` - Authorization was reversed by the merchant.
	//   - `BALANCE_INQUIRY` - A balance inquiry (typically a $0 authorization) has
	//     occurred on a card.
	//   - `CLEARING` - Transaction is settled.
	//   - `CORRECTION_DEBIT` - Manual transaction correction (Debit).
	//   - `CORRECTION_CREDIT` - Manual transaction correction (Credit).
	//   - `CREDIT_AUTHORIZATION` - A refund or credit authorization from a merchant.
	//   - `CREDIT_AUTHORIZATION_ADVICE` - A credit authorization was approved on your
	//     behalf by the network.
	//   - `FINANCIAL_AUTHORIZATION` - A request from a merchant to debit funds without
	//     additional clearing.
	//   - `FINANCIAL_CREDIT_AUTHORIZATION` - A request from a merchant to refund or
	//     credit funds without additional clearing.
	//   - `RETURN` - A refund has been processed on the transaction.
	//   - `RETURN_REVERSAL` - A refund has been reversed (e.g., when a merchant reverses
	//     an incorrect refund).
	Type TransactionEventsType `json:"type,required"`
	JSON TransactionEventsJSON
}

type TransactionEventsJSON struct {
	Amount  pjson.Metadata
	Created pjson.Metadata
	Result  pjson.Metadata
	Token   pjson.Metadata
	Type    pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransactionEvents using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransactionEvents) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionEventsResult string

const (
	TransactionEventsResultAccountStateTransaction TransactionEventsResult = "ACCOUNT_STATE_TRANSACTION"
	TransactionEventsResultApproved                TransactionEventsResult = "APPROVED"
	TransactionEventsResultBankConnectionError     TransactionEventsResult = "BANK_CONNECTION_ERROR"
	TransactionEventsResultBankNotVerified         TransactionEventsResult = "BANK_NOT_VERIFIED"
	TransactionEventsResultCardClosed              TransactionEventsResult = "CARD_CLOSED"
	TransactionEventsResultCardPaused              TransactionEventsResult = "CARD_PAUSED"
	TransactionEventsResultFraudAdvice             TransactionEventsResult = "FRAUD_ADVICE"
	TransactionEventsResultGlobalTransactionLimit  TransactionEventsResult = "GLOBAL_TRANSACTION_LIMIT"
	TransactionEventsResultGlobalWeeklyLimit       TransactionEventsResult = "GLOBAL_WEEKLY_LIMIT"
	TransactionEventsResultGlobalMonthlyLimit      TransactionEventsResult = "GLOBAL_MONTHLY_LIMIT"
	TransactionEventsResultInactiveAccount         TransactionEventsResult = "INACTIVE_ACCOUNT"
	TransactionEventsResultIncorrectPin            TransactionEventsResult = "INCORRECT_PIN"
	TransactionEventsResultInvalidCardDetails      TransactionEventsResult = "INVALID_CARD_DETAILS"
	TransactionEventsResultInsufficientFunds       TransactionEventsResult = "INSUFFICIENT_FUNDS"
	TransactionEventsResultMerchantBlacklist       TransactionEventsResult = "MERCHANT_BLACKLIST"
	TransactionEventsResultSingleUseRecharged      TransactionEventsResult = "SINGLE_USE_RECHARGED"
	TransactionEventsResultSwitchInoperativeAdvice TransactionEventsResult = "SWITCH_INOPERATIVE_ADVICE"
	TransactionEventsResultUnauthorizedMerchant    TransactionEventsResult = "UNAUTHORIZED_MERCHANT"
	TransactionEventsResultUnknownHostTimeout      TransactionEventsResult = "UNKNOWN_HOST_TIMEOUT"
	TransactionEventsResultUserTransactionLimit    TransactionEventsResult = "USER_TRANSACTION_LIMIT"
)

type TransactionEventsType string

const (
	TransactionEventsTypeAuthorization                TransactionEventsType = "AUTHORIZATION"
	TransactionEventsTypeAuthorizationAdvice          TransactionEventsType = "AUTHORIZATION_ADVICE"
	TransactionEventsTypeAuthorizationExpiry          TransactionEventsType = "AUTHORIZATION_EXPIRY"
	TransactionEventsTypeAuthorizationReversal        TransactionEventsType = "AUTHORIZATION_REVERSAL"
	TransactionEventsTypeBalanceInquiry               TransactionEventsType = "BALANCE_INQUIRY"
	TransactionEventsTypeClearing                     TransactionEventsType = "CLEARING"
	TransactionEventsTypeCorrectionDebit              TransactionEventsType = "CORRECTION_DEBIT"
	TransactionEventsTypeCorrectionCredit             TransactionEventsType = "CORRECTION_CREDIT"
	TransactionEventsTypeCreditAuthorization          TransactionEventsType = "CREDIT_AUTHORIZATION"
	TransactionEventsTypeCreditAuthorizationAdvice    TransactionEventsType = "CREDIT_AUTHORIZATION_ADVICE"
	TransactionEventsTypeFinancialAuthorization       TransactionEventsType = "FINANCIAL_AUTHORIZATION"
	TransactionEventsTypeFinancialCreditAuthorization TransactionEventsType = "FINANCIAL_CREDIT_AUTHORIZATION"
	TransactionEventsTypeReturn                       TransactionEventsType = "RETURN"
	TransactionEventsTypeReturnReversal               TransactionEventsType = "RETURN_REVERSAL"
	TransactionEventsTypeVoid                         TransactionEventsType = "VOID"
)

type TransactionMerchant struct {
	// Unique identifier to identify the payment card acceptor.
	AcceptorID string `json:"acceptor_id"`
	// City of card acceptor.
	City string `json:"city"`
	// Uppercase country of card acceptor (see ISO 8583 specs).
	Country string `json:"country"`
	// Short description of card acceptor.
	Descriptor string `json:"descriptor"`
	// Merchant category code (MCC). A four-digit number listed in ISO 18245. An MCC is
	// used to classify a business by the types of goods or services it provides.
	Mcc string `json:"mcc"`
	// Geographic state of card acceptor (see ISO 8583 specs).
	State string `json:"state"`
	JSON  TransactionMerchantJSON
}

type TransactionMerchantJSON struct {
	AcceptorID pjson.Metadata
	City       pjson.Metadata
	Country    pjson.Metadata
	Descriptor pjson.Metadata
	Mcc        pjson.Metadata
	State      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransactionMerchant using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransactionMerchant) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionNetwork string

const (
	TransactionNetworkInterlink  TransactionNetwork = "INTERLINK"
	TransactionNetworkMaestro    TransactionNetwork = "MAESTRO"
	TransactionNetworkMastercard TransactionNetwork = "MASTERCARD"
	TransactionNetworkVisa       TransactionNetwork = "VISA"
	TransactionNetworkUnknown    TransactionNetwork = "UNKNOWN"
)

type TransactionResult string

const (
	TransactionResultAccountStateTransaction TransactionResult = "ACCOUNT_STATE_TRANSACTION"
	TransactionResultApproved                TransactionResult = "APPROVED"
	TransactionResultBankConnectionError     TransactionResult = "BANK_CONNECTION_ERROR"
	TransactionResultBankNotVerified         TransactionResult = "BANK_NOT_VERIFIED"
	TransactionResultCardClosed              TransactionResult = "CARD_CLOSED"
	TransactionResultCardPaused              TransactionResult = "CARD_PAUSED"
	TransactionResultFraudAdvice             TransactionResult = "FRAUD_ADVICE"
	TransactionResultGlobalTransactionLimit  TransactionResult = "GLOBAL_TRANSACTION_LIMIT"
	TransactionResultGlobalWeeklyLimit       TransactionResult = "GLOBAL_WEEKLY_LIMIT"
	TransactionResultGlobalMonthlyLimit      TransactionResult = "GLOBAL_MONTHLY_LIMIT"
	TransactionResultInactiveAccount         TransactionResult = "INACTIVE_ACCOUNT"
	TransactionResultIncorrectPin            TransactionResult = "INCORRECT_PIN"
	TransactionResultInvalidCardDetails      TransactionResult = "INVALID_CARD_DETAILS"
	TransactionResultInsufficientFunds       TransactionResult = "INSUFFICIENT_FUNDS"
	TransactionResultMerchantBlacklist       TransactionResult = "MERCHANT_BLACKLIST"
	TransactionResultSingleUseRecharged      TransactionResult = "SINGLE_USE_RECHARGED"
	TransactionResultSwitchInoperativeAdvice TransactionResult = "SWITCH_INOPERATIVE_ADVICE"
	TransactionResultUnauthorizedMerchant    TransactionResult = "UNAUTHORIZED_MERCHANT"
	TransactionResultUnknownHostTimeout      TransactionResult = "UNKNOWN_HOST_TIMEOUT"
	TransactionResultUserTransactionLimit    TransactionResult = "USER_TRANSACTION_LIMIT"
)

type TransactionStatus string

const (
	TransactionStatusBounced  TransactionStatus = "BOUNCED"
	TransactionStatusDeclined TransactionStatus = "DECLINED"
	TransactionStatusExpired  TransactionStatus = "EXPIRED"
	TransactionStatusPending  TransactionStatus = "PENDING"
	TransactionStatusSettled  TransactionStatus = "SETTLED"
	TransactionStatusSettling TransactionStatus = "SETTLING"
	TransactionStatusVoided   TransactionStatus = "VOIDED"
)

type TransactionSimulateAuthorizationResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string `json:"debugging_request_id" format:"uuid"`
	// A unique token to reference this transaction with later calls to void or clear
	// the authorization.
	Token string `json:"token" format:"uuid"`
	JSON  TransactionSimulateAuthorizationResponseJSON
}

type TransactionSimulateAuthorizationResponseJSON struct {
	DebuggingRequestID pjson.Metadata
	Token              pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSimulateAuthorizationResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSimulateAuthorizationResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSimulateClearingResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string `json:"debugging_request_id" format:"uuid"`
	JSON               TransactionSimulateClearingResponseJSON
}

type TransactionSimulateClearingResponseJSON struct {
	DebuggingRequestID pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSimulateClearingResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSimulateClearingResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSimulateReturnResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string `json:"debugging_request_id" format:"uuid"`
	// A unique token to reference this transaction.
	Token string `json:"token" format:"uuid"`
	JSON  TransactionSimulateReturnResponseJSON
}

type TransactionSimulateReturnResponseJSON struct {
	DebuggingRequestID pjson.Metadata
	Token              pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSimulateReturnResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSimulateReturnResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSimulateReturnReversalResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string `json:"debugging_request_id" format:"uuid"`
	JSON               TransactionSimulateReturnReversalResponseJSON
}

type TransactionSimulateReturnReversalResponseJSON struct {
	DebuggingRequestID pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSimulateReturnReversalResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSimulateReturnReversalResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSimulateVoidResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string `json:"debugging_request_id" format:"uuid"`
	JSON               TransactionSimulateVoidResponseJSON
}

type TransactionSimulateVoidResponseJSON struct {
	DebuggingRequestID pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSimulateVoidResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSimulateVoidResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSimulateCreditAuthorizationResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string `json:"debugging_request_id" format:"uuid"`
	// A unique token to reference this transaction.
	Token string `json:"token" format:"uuid"`
	JSON  TransactionSimulateCreditAuthorizationResponseJSON
}

type TransactionSimulateCreditAuthorizationResponseJSON struct {
	DebuggingRequestID pjson.Metadata
	Token              pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSimulateCreditAuthorizationResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSimulateCreditAuthorizationResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSimulateAuthorizationAdviceResponse struct {
	// Debugging request ID to share with Lithic Support team.
	DebuggingRequestID string `json:"debugging_request_id" format:"uuid"`
	// A unique token to reference this transaction.
	Token string `json:"token" format:"uuid"`
	JSON  TransactionSimulateAuthorizationAdviceResponseJSON
}

type TransactionSimulateAuthorizationAdviceResponseJSON struct {
	DebuggingRequestID pjson.Metadata
	Token              pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSimulateAuthorizationAdviceResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSimulateAuthorizationAdviceResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionListResponse struct {
	Data []Transaction `json:"data,required"`
	// Page of the result.
	Page int64 `json:"page,required"`
	// Number of matched rows.
	TotalEntries int64 `json:"total_entries,required"`
	// Total pages of result.
	TotalPages int64 `json:"total_pages,required"`
	JSON       TransactionListResponseJSON
}

type TransactionListResponseJSON struct {
	Data         pjson.Metadata
	Page         pjson.Metadata
	TotalEntries pjson.Metadata
	TotalPages   pjson.Metadata
	Raw          []byte
	Extras       map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransactionListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
