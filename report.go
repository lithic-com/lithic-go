// File generated from our OpenAPI spec by Stainless.

package lithic

import (
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/option"
)

// ReportService contains methods and other services that help with interacting
// with the lithic API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewReportService] method instead.
type ReportService struct {
	Options    []option.RequestOption
	Settlement *ReportSettlementService
}

// NewReportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReportService(opts ...option.RequestOption) (r *ReportService) {
	r = &ReportService{}
	r.Options = opts
	r.Settlement = NewReportSettlementService(opts...)
	return
}

type SettlementDetail struct {
	// Globally unique identifier denoting the Settlement Detail.
	Token string `json:"token,required" format:"uuid"`
	// The most granular ID the network settles with (e.g., ICA for Mastercard, FTSRE
	// for Visa).
	AccountToken string `json:"account_token,required" format:"uuid"`
	// Globally unique identifier denoting the card program that the associated
	// Transaction occurred on.
	CardProgramToken string `json:"card_program_token,required" format:"uuid"`
	// Globally unique identifier denoting the card that the associated Transaction
	// occurred on.
	CardToken string `json:"card_token,required" format:"uuid"`
	// Date and time when the transaction first occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Three-digit alphabetic ISO 4217 code.
	Currency string `json:"currency,required"`
	// The total gross amount of disputes settlements.
	DisputesGrossAmount int64 `json:"disputes_gross_amount,required"`
	// Globally unique identifiers denoting the Events associated with this settlement.
	EventTokens []string `json:"event_tokens,required"`
	// The most granular ID the network settles with (e.g., ICA for Mastercard, FTSRE
	// for Visa).
	Institution string `json:"institution,required"`
	// The total amount of interchange.
	InterchangeGrossAmount int64 `json:"interchange_gross_amount,required"`
	// Card network where the transaction took place.
	Network SettlementDetailNetwork `json:"network,required"`
	// The total gross amount of other fees by type.
	OtherFeesDetails SettlementDetailOtherFeesDetails `json:"other_fees_details,required"`
	// Total amount of gross other fees outside of interchange.
	OtherFeesGrossAmount int64 `json:"other_fees_gross_amount,required"`
	// Date of when the report was first generated.
	ReportDate string `json:"report_date,required"`
	// Date of when money movement is triggered for the transaction.
	SettlementDate string `json:"settlement_date,required"`
	// Globally unique identifier denoting the associated Transaction object.
	TransactionToken string `json:"transaction_token,required" format:"uuid"`
	// The total amount of settlement impacting transactions (excluding interchange,
	// fees, and disputes).
	TransactionsGrossAmount int64 `json:"transactions_gross_amount,required"`
	// Date and time when the transaction first occurred. UTC time zone.
	Updated time.Time            `json:"updated,required" format:"date-time"`
	JSON    settlementDetailJSON `json:"-"`
}

// settlementDetailJSON contains the JSON metadata for the struct
// [SettlementDetail]
type settlementDetailJSON struct {
	Token                   apijson.Field
	AccountToken            apijson.Field
	CardProgramToken        apijson.Field
	CardToken               apijson.Field
	Created                 apijson.Field
	Currency                apijson.Field
	DisputesGrossAmount     apijson.Field
	EventTokens             apijson.Field
	Institution             apijson.Field
	InterchangeGrossAmount  apijson.Field
	Network                 apijson.Field
	OtherFeesDetails        apijson.Field
	OtherFeesGrossAmount    apijson.Field
	ReportDate              apijson.Field
	SettlementDate          apijson.Field
	TransactionToken        apijson.Field
	TransactionsGrossAmount apijson.Field
	Updated                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettlementDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Card network where the transaction took place.
type SettlementDetailNetwork string

const (
	SettlementDetailNetworkInterlink  SettlementDetailNetwork = "INTERLINK"
	SettlementDetailNetworkMaestro    SettlementDetailNetwork = "MAESTRO"
	SettlementDetailNetworkMastercard SettlementDetailNetwork = "MASTERCARD"
	SettlementDetailNetworkUnknown    SettlementDetailNetwork = "UNKNOWN"
	SettlementDetailNetworkVisa       SettlementDetailNetwork = "VISA"
)

// The total gross amount of other fees by type.
type SettlementDetailOtherFeesDetails struct {
	Isa  int64                                `json:"ISA"`
	JSON settlementDetailOtherFeesDetailsJSON `json:"-"`
}

// settlementDetailOtherFeesDetailsJSON contains the JSON metadata for the struct
// [SettlementDetailOtherFeesDetails]
type settlementDetailOtherFeesDetailsJSON struct {
	Isa         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettlementDetailOtherFeesDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettlementReport struct {
	// Date and time when the transaction first occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Three-digit alphabetic ISO 4217 code.
	Currency string                     `json:"currency,required"`
	Details  []SettlementSummaryDetails `json:"details,required"`
	// The total gross amount of disputes settlements.
	DisputesGrossAmount int64 `json:"disputes_gross_amount,required"`
	// The total amount of interchange.
	InterchangeGrossAmount int64 `json:"interchange_gross_amount,required"`
	// Total amount of gross other fees outside of interchange.
	OtherFeesGrossAmount int64 `json:"other_fees_gross_amount,required"`
	// Date of when the report was first generated.
	ReportDate string `json:"report_date,required"`
	// The total net amount of cash moved. (net value of settled_gross_amount,
	// interchange, fees).
	SettledNetAmount int64 `json:"settled_net_amount,required"`
	// The total amount of settlement impacting transactions (excluding interchange,
	// fees, and disputes).
	TransactionsGrossAmount int64 `json:"transactions_gross_amount,required"`
	// Date and time when the transaction first occurred. UTC time zone.
	Updated time.Time            `json:"updated,required" format:"date-time"`
	JSON    settlementReportJSON `json:"-"`
}

// settlementReportJSON contains the JSON metadata for the struct
// [SettlementReport]
type settlementReportJSON struct {
	Created                 apijson.Field
	Currency                apijson.Field
	Details                 apijson.Field
	DisputesGrossAmount     apijson.Field
	InterchangeGrossAmount  apijson.Field
	OtherFeesGrossAmount    apijson.Field
	ReportDate              apijson.Field
	SettledNetAmount        apijson.Field
	TransactionsGrossAmount apijson.Field
	Updated                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettlementReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettlementSummaryDetails struct {
	// The total gross amount of disputes settlements.
	DisputesGrossAmount int64 `json:"disputes_gross_amount"`
	// The most granular ID the network settles with (e.g., ICA for Mastercard, FTSRE
	// for Visa).
	Institution string `json:"institution"`
	// The total amount of interchange.
	InterchangeGrossAmount int64 `json:"interchange_gross_amount"`
	// Card network where the transaction took place
	Network SettlementSummaryDetailsNetwork `json:"network"`
	// Total amount of gross other fees outside of interchange.
	OtherFeesGrossAmount int64 `json:"other_fees_gross_amount"`
	// The total net amount of cash moved. (net value of settled_gross_amount,
	// interchange, fees).
	SettledNetAmount int64 `json:"settled_net_amount"`
	// The total amount of settlement impacting transactions (excluding interchange,
	// fees, and disputes).
	TransactionsGrossAmount int64                        `json:"transactions_gross_amount"`
	JSON                    settlementSummaryDetailsJSON `json:"-"`
}

// settlementSummaryDetailsJSON contains the JSON metadata for the struct
// [SettlementSummaryDetails]
type settlementSummaryDetailsJSON struct {
	DisputesGrossAmount     apijson.Field
	Institution             apijson.Field
	InterchangeGrossAmount  apijson.Field
	Network                 apijson.Field
	OtherFeesGrossAmount    apijson.Field
	SettledNetAmount        apijson.Field
	TransactionsGrossAmount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettlementSummaryDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Card network where the transaction took place
type SettlementSummaryDetailsNetwork string

const (
	SettlementSummaryDetailsNetworkInterlink  SettlementSummaryDetailsNetwork = "INTERLINK"
	SettlementSummaryDetailsNetworkMaestro    SettlementSummaryDetailsNetwork = "MAESTRO"
	SettlementSummaryDetailsNetworkMastercard SettlementSummaryDetailsNetwork = "MASTERCARD"
	SettlementSummaryDetailsNetworkUnknown    SettlementSummaryDetailsNetwork = "UNKNOWN"
	SettlementSummaryDetailsNetworkVisa       SettlementSummaryDetailsNetwork = "VISA"
)
