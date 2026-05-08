// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/option"
)

// AuthRuleService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthRuleService] method instead.
type AuthRuleService struct {
	Options []option.RequestOption
	V2      *AuthRuleV2Service
}

// NewAuthRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthRuleService(opts ...option.RequestOption) (r *AuthRuleService) {
	r = &AuthRuleService{}
	r.Options = opts
	r.V2 = NewAuthRuleV2Service(opts...)
	return
}

// Behavioral feature state for a card or account derived from its transaction
// history.
//
// Derived statistical features (averages, standard deviations, z-scores) are
// computed using Welford's online algorithm over approved transactions. Average
// fields are null when fewer than 5 approved transactions have been recorded.
// Standard deviation fields are null when fewer than 30 approved transactions have
// been recorded.
//
// 3DS fields (`three_ds_success_rate`, `three_ds_success_count`,
// `three_ds_total_count`) are card-scoped and will be null for account responses.
//
// Raw fields (`seen_countries`, `seen_mccs`, `approved_txn_amount_m2`, etc.) are
// included so clients can compute their own transaction-specific derivations, such
// as checking whether a new transaction's country is in `seen_countries` to
// determine `is_new_country`, or computing a z-score using the raw mean and M2
// values.
type SignalsResponse struct {
	// The Welford M2 accumulator for lifetime approved transaction amounts. Used
	// together with `avg_transaction_amount` and `approved_txn_count` to compute the
	// z-score of a new transaction amount (variance = M2 / (count - 1)).
	ApprovedTxnAmountM2 float64 `json:"approved_txn_amount_m2" api:"required,nullable"`
	// The Welford M2 accumulator for approved transaction amounts over the last 30
	// days.
	ApprovedTxnAmountM2_30d float64 `json:"approved_txn_amount_m2_30d" api:"required,nullable"`
	// The Welford M2 accumulator for approved transaction amounts over the last 7
	// days.
	ApprovedTxnAmountM2_7d float64 `json:"approved_txn_amount_m2_7d" api:"required,nullable"`
	// The Welford M2 accumulator for approved transaction amounts over the last 90
	// days.
	ApprovedTxnAmountM2_90d float64 `json:"approved_txn_amount_m2_90d" api:"required,nullable"`
	// The total number of approved transactions over the entity's lifetime.
	ApprovedTxnCount int64 `json:"approved_txn_count" api:"required,nullable"`
	// The number of approved transactions in the last 30 days.
	ApprovedTxnCount30d int64 `json:"approved_txn_count_30d" api:"required,nullable"`
	// The number of approved transactions in the last 7 days.
	ApprovedTxnCount7d int64 `json:"approved_txn_count_7d" api:"required,nullable"`
	// The number of approved transactions in the last 90 days.
	ApprovedTxnCount90d int64 `json:"approved_txn_count_90d" api:"required,nullable"`
	// The average approved transaction amount over the entity's lifetime, in cents.
	// Null if fewer than 5 approved transactions have been recorded.
	AvgTransactionAmount float64 `json:"avg_transaction_amount" api:"required,nullable"`
	// The average approved transaction amount over the last 30 days, in cents. Null if
	// fewer than 5 approved transactions in window.
	AvgTransactionAmount30d float64 `json:"avg_transaction_amount_30d" api:"required,nullable"`
	// The average approved transaction amount over the last 7 days, in cents. Null if
	// fewer than 5 approved transactions in window.
	AvgTransactionAmount7d float64 `json:"avg_transaction_amount_7d" api:"required,nullable"`
	// The average approved transaction amount over the last 90 days, in cents. Null if
	// fewer than 5 approved transactions in window.
	AvgTransactionAmount90d float64 `json:"avg_transaction_amount_90d" api:"required,nullable"`
	// The number of distinct merchant countries seen in the entity's transaction
	// history.
	DistinctCountryCount int64 `json:"distinct_country_count" api:"required,nullable"`
	// The number of distinct MCCs seen in the entity's transaction history.
	DistinctMccCount int64 `json:"distinct_mcc_count" api:"required,nullable"`
	// The timestamp of the first approved transaction for the entity, in ISO 8601
	// format.
	FirstTxnAt time.Time `json:"first_txn_at" api:"required,nullable" format:"date-time"`
	// Whether the entity has no prior transaction history. Returns true if no history
	// is found. Null if transaction history exists but a first transaction timestamp
	// is unavailable.
	IsFirstTransaction bool `json:"is_first_transaction" api:"required,nullable"`
	// The merchant country of the last card-present transaction.
	LastCpCountry string `json:"last_cp_country" api:"required,nullable"`
	// The merchant postal code of the last card-present transaction.
	LastCpPostalCode string `json:"last_cp_postal_code" api:"required,nullable"`
	// The timestamp of the last card-present transaction, in ISO 8601 format.
	LastCpTimestamp time.Time `json:"last_cp_timestamp" api:"required,nullable" format:"date-time"`
	// The timestamp of the most recent approved transaction for the entity, in ISO
	// 8601 format.
	LastTxnApprovedAt time.Time `json:"last_txn_approved_at" api:"required,nullable" format:"date-time"`
	// The set of merchant countries seen in the entity's transaction history. Clients
	// can use this to determine whether a new transaction's country is novel (i.e.
	// compute `is_new_country`).
	SeenCountries []string `json:"seen_countries" api:"required,nullable"`
	// The set of MCCs seen in the entity's transaction history. Clients can use this
	// to determine whether a new transaction's MCC is novel (i.e. compute
	// `is_new_mcc`).
	SeenMccs []string `json:"seen_mccs" api:"required,nullable"`
	// The set of card acceptor IDs seen in the card's approved transaction history,
	// capped at the 1000 most recently seen. Null for account responses. Clients can
	// use this to determine whether a new transaction's merchant is novel (i.e.
	// compute `is_new_merchant`).
	SeenMerchants []string `json:"seen_merchants" api:"required,nullable"`
	// The standard deviation of approved transaction amounts over the entity's
	// lifetime, in cents. Null if fewer than 30 approved transactions have been
	// recorded.
	StdevTransactionAmount float64 `json:"stdev_transaction_amount" api:"required,nullable"`
	// The standard deviation of approved transaction amounts over the last 30 days, in
	// cents. Null if fewer than 30 approved transactions in window.
	StdevTransactionAmount30d float64 `json:"stdev_transaction_amount_30d" api:"required,nullable"`
	// The standard deviation of approved transaction amounts over the last 7 days, in
	// cents. Null if fewer than 30 approved transactions in window.
	StdevTransactionAmount7d float64 `json:"stdev_transaction_amount_7d" api:"required,nullable"`
	// The standard deviation of approved transaction amounts over the last 90 days, in
	// cents. Null if fewer than 30 approved transactions in window.
	StdevTransactionAmount90d float64 `json:"stdev_transaction_amount_90d" api:"required,nullable"`
	// The number of successful 3DS authentications for the card. Null for account
	// responses.
	ThreeDSSuccessCount int64 `json:"three_ds_success_count" api:"required,nullable"`
	// The 3DS authentication success rate for the card, as a percentage from 0.0 to
	// 100.0. Null for account responses.
	ThreeDSSuccessRate float64 `json:"three_ds_success_rate" api:"required,nullable"`
	// The total number of 3DS authentication attempts for the card. Null for account
	// responses.
	ThreeDSTotalCount int64 `json:"three_ds_total_count" api:"required,nullable"`
	// The number of days since the last approved transaction on the entity.
	TimeSinceLastTransactionDays float64             `json:"time_since_last_transaction_days" api:"required,nullable"`
	JSON                         signalsResponseJSON `json:"-"`
}

// signalsResponseJSON contains the JSON metadata for the struct [SignalsResponse]
type signalsResponseJSON struct {
	ApprovedTxnAmountM2          apijson.Field
	ApprovedTxnAmountM2_30d      apijson.Field
	ApprovedTxnAmountM2_7d       apijson.Field
	ApprovedTxnAmountM2_90d      apijson.Field
	ApprovedTxnCount             apijson.Field
	ApprovedTxnCount30d          apijson.Field
	ApprovedTxnCount7d           apijson.Field
	ApprovedTxnCount90d          apijson.Field
	AvgTransactionAmount         apijson.Field
	AvgTransactionAmount30d      apijson.Field
	AvgTransactionAmount7d       apijson.Field
	AvgTransactionAmount90d      apijson.Field
	DistinctCountryCount         apijson.Field
	DistinctMccCount             apijson.Field
	FirstTxnAt                   apijson.Field
	IsFirstTransaction           apijson.Field
	LastCpCountry                apijson.Field
	LastCpPostalCode             apijson.Field
	LastCpTimestamp              apijson.Field
	LastTxnApprovedAt            apijson.Field
	SeenCountries                apijson.Field
	SeenMccs                     apijson.Field
	SeenMerchants                apijson.Field
	StdevTransactionAmount       apijson.Field
	StdevTransactionAmount30d    apijson.Field
	StdevTransactionAmount7d     apijson.Field
	StdevTransactionAmount90d    apijson.Field
	ThreeDSSuccessCount          apijson.Field
	ThreeDSSuccessRate           apijson.Field
	ThreeDSTotalCount            apijson.Field
	TimeSinceLastTransactionDays apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SignalsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r signalsResponseJSON) RawJSON() string {
	return r.raw
}
