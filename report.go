// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/option"
)

// ReportService contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportService] method instead.
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

type NetworkTotal struct {
	// Globally unique identifier.
	Token   string              `json:"token,required" format:"uuid"`
	Amounts NetworkTotalAmounts `json:"amounts,required"`
	// RFC 3339 timestamp for when the record was created. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-character alphabetic ISO 4217 code.
	Currency string `json:"currency,required"`
	// The institution that activity occurred on. For Mastercard: ICA (Interbank Card
	// Association). For Maestro: institution ID. For Visa: lowest level SRE
	// (Settlement Reporting Entity).
	InstitutionID string `json:"institution_id,required"`
	// Indicates that all settlement records related to this Network Total are
	// available in the details endpoint.
	IsComplete bool `json:"is_complete,required"`
	// Card network where the transaction took place. AMEX, VISA, MASTERCARD, MAESTRO,
	// or INTERLINK.
	Network NetworkTotalNetwork `json:"network,required"`
	// Date that the network total record applies to. YYYY-MM-DD format.
	ReportDate time.Time `json:"report_date,required" format:"date"`
	// The institution responsible for settlement. For Mastercard: same as
	// `institution_id`. For Maestro: billing ICA. For Visa: Funds Transfer SRE
	// (FTSRE).
	SettlementInstitutionID string `json:"settlement_institution_id,required"`
	// Settlement service.
	SettlementService string `json:"settlement_service,required"`
	// RFC 3339 timestamp for when the record was last updated. UTC time zone.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// The clearing cycle that the network total record applies to. Mastercard only.
	Cycle int64            `json:"cycle"`
	JSON  networkTotalJSON `json:"-"`
}

// networkTotalJSON contains the JSON metadata for the struct [NetworkTotal]
type networkTotalJSON struct {
	Token                   apijson.Field
	Amounts                 apijson.Field
	Created                 apijson.Field
	Currency                apijson.Field
	InstitutionID           apijson.Field
	IsComplete              apijson.Field
	Network                 apijson.Field
	ReportDate              apijson.Field
	SettlementInstitutionID apijson.Field
	SettlementService       apijson.Field
	Updated                 apijson.Field
	Cycle                   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NetworkTotal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkTotalJSON) RawJSON() string {
	return r.raw
}

type NetworkTotalAmounts struct {
	// Total settlement amount excluding interchange, in currency's smallest unit.
	GrossSettlement int64 `json:"gross_settlement,required"`
	// Interchange amount, in currency's smallest unit.
	InterchangeFees int64 `json:"interchange_fees,required"`
	// `gross_settlement` net of `interchange_fees` and `visa_charges` (if applicable),
	// in currency's smallest unit.
	NetSettlement int64 `json:"net_settlement,required"`
	// Charges specific to Visa/Interlink, in currency's smallest unit.
	VisaCharges int64                   `json:"visa_charges"`
	JSON        networkTotalAmountsJSON `json:"-"`
}

// networkTotalAmountsJSON contains the JSON metadata for the struct
// [NetworkTotalAmounts]
type networkTotalAmountsJSON struct {
	GrossSettlement apijson.Field
	InterchangeFees apijson.Field
	NetSettlement   apijson.Field
	VisaCharges     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NetworkTotalAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkTotalAmountsJSON) RawJSON() string {
	return r.raw
}

// Card network where the transaction took place. AMEX, VISA, MASTERCARD, MAESTRO,
// or INTERLINK.
type NetworkTotalNetwork string

const (
	NetworkTotalNetworkAmex       NetworkTotalNetwork = "AMEX"
	NetworkTotalNetworkVisa       NetworkTotalNetwork = "VISA"
	NetworkTotalNetworkMastercard NetworkTotalNetwork = "MASTERCARD"
	NetworkTotalNetworkMaestro    NetworkTotalNetwork = "MAESTRO"
	NetworkTotalNetworkInterlink  NetworkTotalNetwork = "INTERLINK"
)

func (r NetworkTotalNetwork) IsKnown() bool {
	switch r {
	case NetworkTotalNetworkAmex, NetworkTotalNetworkVisa, NetworkTotalNetworkMastercard, NetworkTotalNetworkMaestro, NetworkTotalNetworkInterlink:
		return true
	}
	return false
}

type SettlementDetail struct {
	// Globally unique identifier denoting the Settlement Detail.
	Token string `json:"token,required" format:"uuid"`
	// Globally unique identifier denoting the account that the associated transaction
	// occurred on.
	AccountToken string `json:"account_token,required" format:"uuid"`
	// Globally unique identifier denoting the card program that the associated
	// transaction occurred on.
	CardProgramToken string `json:"card_program_token,required" format:"uuid"`
	// Globally unique identifier denoting the card that the associated transaction
	// occurred on.
	CardToken string `json:"card_token,required" format:"uuid"`
	// Date and time when the transaction first occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// Three-character alphabetic ISO 4217 code.
	Currency string `json:"currency,required"`
	// The total gross amount of disputes settlements.
	DisputesGrossAmount int64 `json:"disputes_gross_amount,required"`
	// Globally unique identifiers denoting the Events associated with this settlement.
	EventTokens []string `json:"event_tokens,required"`
	// The most granular ID the network settles with (e.g., ICA for Mastercard, FTSRE
	// for Visa).
	Institution string `json:"institution,required"`
	// The total amount of interchange in six-digit extended precision.
	InterchangeFeeExtendedPrecision int64 `json:"interchange_fee_extended_precision,required"`
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
	// Date of when money movement is triggered for the transaction. One exception
	// applies - for Mastercard dual message settlement, this is the settlement
	// advisement date, which is distinct from the date of money movement.
	SettlementDate string `json:"settlement_date,required"`
	// Globally unique identifier denoting the associated Transaction object.
	TransactionToken string `json:"transaction_token,required" format:"uuid"`
	// The total amount of settlement impacting transactions (excluding interchange,
	// fees, and disputes).
	TransactionsGrossAmount int64 `json:"transactions_gross_amount,required"`
	// The type of settlement record.
	Type SettlementDetailType `json:"type,required"`
	// Date and time when the transaction first occurred. UTC time zone.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Network's description of a fee, only present on records with type `FEE`.
	FeeDescription string               `json:"fee_description"`
	JSON           settlementDetailJSON `json:"-"`
}

// settlementDetailJSON contains the JSON metadata for the struct
// [SettlementDetail]
type settlementDetailJSON struct {
	Token                           apijson.Field
	AccountToken                    apijson.Field
	CardProgramToken                apijson.Field
	CardToken                       apijson.Field
	Created                         apijson.Field
	Currency                        apijson.Field
	DisputesGrossAmount             apijson.Field
	EventTokens                     apijson.Field
	Institution                     apijson.Field
	InterchangeFeeExtendedPrecision apijson.Field
	InterchangeGrossAmount          apijson.Field
	Network                         apijson.Field
	OtherFeesDetails                apijson.Field
	OtherFeesGrossAmount            apijson.Field
	ReportDate                      apijson.Field
	SettlementDate                  apijson.Field
	TransactionToken                apijson.Field
	TransactionsGrossAmount         apijson.Field
	Type                            apijson.Field
	Updated                         apijson.Field
	FeeDescription                  apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *SettlementDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settlementDetailJSON) RawJSON() string {
	return r.raw
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

func (r SettlementDetailNetwork) IsKnown() bool {
	switch r {
	case SettlementDetailNetworkInterlink, SettlementDetailNetworkMaestro, SettlementDetailNetworkMastercard, SettlementDetailNetworkUnknown, SettlementDetailNetworkVisa:
		return true
	}
	return false
}

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

func (r settlementDetailOtherFeesDetailsJSON) RawJSON() string {
	return r.raw
}

// The type of settlement record.
type SettlementDetailType string

const (
	SettlementDetailTypeAdjustment     SettlementDetailType = "ADJUSTMENT"
	SettlementDetailTypeArbitration    SettlementDetailType = "ARBITRATION"
	SettlementDetailTypeChargeback     SettlementDetailType = "CHARGEBACK"
	SettlementDetailTypeClearing       SettlementDetailType = "CLEARING"
	SettlementDetailTypeCollaboration  SettlementDetailType = "COLLABORATION"
	SettlementDetailTypeFee            SettlementDetailType = "FEE"
	SettlementDetailTypeFinancial      SettlementDetailType = "FINANCIAL"
	SettlementDetailTypeNonFinancial   SettlementDetailType = "NON-FINANCIAL"
	SettlementDetailTypePrearbitration SettlementDetailType = "PREARBITRATION"
	SettlementDetailTypeRepresentment  SettlementDetailType = "REPRESENTMENT"
)

func (r SettlementDetailType) IsKnown() bool {
	switch r {
	case SettlementDetailTypeAdjustment, SettlementDetailTypeArbitration, SettlementDetailTypeChargeback, SettlementDetailTypeClearing, SettlementDetailTypeCollaboration, SettlementDetailTypeFee, SettlementDetailTypeFinancial, SettlementDetailTypeNonFinancial, SettlementDetailTypePrearbitration, SettlementDetailTypeRepresentment:
		return true
	}
	return false
}

type SettlementReport struct {
	// Date and time when the transaction first occurred. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// 3-character alphabetic ISO 4217 code. (This field is deprecated and will be
	// removed in a future version of the API.)
	//
	// Deprecated: deprecated
	Currency string                     `json:"currency,required"`
	Details  []SettlementSummaryDetails `json:"details,required"`
	// The total gross amount of disputes settlements. (This field is deprecated and
	// will be removed in a future version of the API. To compute total amounts, Lithic
	// recommends that customers sum the relevant settlement amounts found within
	// `details`.)
	//
	// Deprecated: deprecated
	DisputesGrossAmount int64 `json:"disputes_gross_amount,required"`
	// The total amount of interchange. (This field is deprecated and will be removed
	// in a future version of the API. To compute total amounts, Lithic recommends that
	// customers sum the relevant settlement amounts found within `details`.)
	//
	// Deprecated: deprecated
	InterchangeGrossAmount int64 `json:"interchange_gross_amount,required"`
	// Indicates that all data expected on the given report date is available.
	IsComplete bool `json:"is_complete,required"`
	// Total amount of gross other fees outside of interchange. (This field is
	// deprecated and will be removed in a future version of the API. To compute total
	// amounts, Lithic recommends that customers sum the relevant settlement amounts
	// found within `details`.)
	//
	// Deprecated: deprecated
	OtherFeesGrossAmount int64 `json:"other_fees_gross_amount,required"`
	// Date of when the report was first generated.
	ReportDate string `json:"report_date,required"`
	// The total net amount of cash moved. (net value of settled_gross_amount,
	// interchange, fees). (This field is deprecated and will be removed in a future
	// version of the API. To compute total amounts, Lithic recommends that customers
	// sum the relevant settlement amounts found within `details`.)
	//
	// Deprecated: deprecated
	SettledNetAmount int64 `json:"settled_net_amount,required"`
	// The total amount of settlement impacting transactions (excluding interchange,
	// fees, and disputes). (This field is deprecated and will be removed in a future
	// version of the API. To compute total amounts, Lithic recommends that customers
	// sum the relevant settlement amounts found within `details`.)
	//
	// Deprecated: deprecated
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
	IsComplete              apijson.Field
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

func (r settlementReportJSON) RawJSON() string {
	return r.raw
}

type SettlementSummaryDetails struct {
	// 3-character alphabetic ISO 4217 code.
	Currency string `json:"currency"`
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
	Currency                apijson.Field
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

func (r settlementSummaryDetailsJSON) RawJSON() string {
	return r.raw
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

func (r SettlementSummaryDetailsNetwork) IsKnown() bool {
	switch r {
	case SettlementSummaryDetailsNetworkInterlink, SettlementSummaryDetailsNetworkMaestro, SettlementSummaryDetailsNetworkMastercard, SettlementSummaryDetailsNetworkUnknown, SettlementSummaryDetailsNetworkVisa:
		return true
	}
	return false
}
