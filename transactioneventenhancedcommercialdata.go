// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// TransactionEventEnhancedCommercialDataService contains methods and other
// services that help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionEventEnhancedCommercialDataService] method instead.
type TransactionEventEnhancedCommercialDataService struct {
	Options []option.RequestOption
}

// NewTransactionEventEnhancedCommercialDataService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewTransactionEventEnhancedCommercialDataService(opts ...option.RequestOption) (r *TransactionEventEnhancedCommercialDataService) {
	r = &TransactionEventEnhancedCommercialDataService{}
	r.Options = opts
	return
}

// Get L2/L3 enhanced commercial data associated with a transaction event. Not
// available in sandbox.
func (r *TransactionEventEnhancedCommercialDataService) Get(ctx context.Context, eventToken string, opts ...option.RequestOption) (res *EnhancedData, err error) {
	opts = append(r.Options[:], opts...)
	if eventToken == "" {
		err = errors.New("missing required event_token parameter")
		return
	}
	path := fmt.Sprintf("transactions/events/%s/enhanced_commercial_data", eventToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EnhancedData struct {
	// A unique identifier for the enhanced commercial data.
	Token  string             `json:"token,required" format:"uuid"`
	Common EnhancedDataCommon `json:"common,required"`
	// The token of the event that the enhanced data is associated with.
	EventToken string              `json:"event_token,required" format:"uuid"`
	Fleet      []EnhancedDataFleet `json:"fleet,required"`
	// The token of the transaction that the enhanced data is associated with.
	TransactionToken string           `json:"transaction_token,required" format:"uuid"`
	JSON             enhancedDataJSON `json:"-"`
}

// enhancedDataJSON contains the JSON metadata for the struct [EnhancedData]
type enhancedDataJSON struct {
	Token            apijson.Field
	Common           apijson.Field
	EventToken       apijson.Field
	Fleet            apijson.Field
	TransactionToken apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EnhancedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enhancedDataJSON) RawJSON() string {
	return r.raw
}

type EnhancedDataCommon struct {
	LineItems []EnhancedDataCommonLineItem `json:"line_items,required"`
	Tax       EnhancedDataCommonTax        `json:"tax,required"`
	// A customer identifier.
	CustomerReferenceNumber string `json:"customer_reference_number"`
	// A merchant identifier.
	MerchantReferenceNumber string `json:"merchant_reference_number"`
	// The date of the order.
	OrderDate time.Time              `json:"order_date" format:"date"`
	JSON      enhancedDataCommonJSON `json:"-"`
}

// enhancedDataCommonJSON contains the JSON metadata for the struct
// [EnhancedDataCommon]
type enhancedDataCommonJSON struct {
	LineItems               apijson.Field
	Tax                     apijson.Field
	CustomerReferenceNumber apijson.Field
	MerchantReferenceNumber apijson.Field
	OrderDate               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EnhancedDataCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enhancedDataCommonJSON) RawJSON() string {
	return r.raw
}

// An L2/L3 enhanced commercial data line item.
type EnhancedDataCommonLineItem struct {
	// The price of the item purchased in merchant currency.
	Amount float64 `json:"amount"`
	// A human-readable description of the item.
	Description string `json:"description"`
	// An identifier for the item purchased.
	ProductCode string `json:"product_code"`
	// The quantity of the item purchased.
	Quantity float64                        `json:"quantity"`
	JSON     enhancedDataCommonLineItemJSON `json:"-"`
}

// enhancedDataCommonLineItemJSON contains the JSON metadata for the struct
// [EnhancedDataCommonLineItem]
type enhancedDataCommonLineItemJSON struct {
	Amount      apijson.Field
	Description apijson.Field
	ProductCode apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnhancedDataCommonLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enhancedDataCommonLineItemJSON) RawJSON() string {
	return r.raw
}

type EnhancedDataCommonTax struct {
	// The amount of tax collected.
	Amount int64 `json:"amount"`
	// A flag indicating whether the transaction is tax exempt or not.
	Exempt EnhancedDataCommonTaxExempt `json:"exempt"`
	// The tax ID of the merchant.
	MerchantTaxID string                    `json:"merchant_tax_id"`
	JSON          enhancedDataCommonTaxJSON `json:"-"`
}

// enhancedDataCommonTaxJSON contains the JSON metadata for the struct
// [EnhancedDataCommonTax]
type enhancedDataCommonTaxJSON struct {
	Amount        apijson.Field
	Exempt        apijson.Field
	MerchantTaxID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnhancedDataCommonTax) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enhancedDataCommonTaxJSON) RawJSON() string {
	return r.raw
}

// A flag indicating whether the transaction is tax exempt or not.
type EnhancedDataCommonTaxExempt string

const (
	EnhancedDataCommonTaxExemptTaxIncluded    EnhancedDataCommonTaxExempt = "TAX_INCLUDED"
	EnhancedDataCommonTaxExemptTaxNotIncluded EnhancedDataCommonTaxExempt = "TAX_NOT_INCLUDED"
	EnhancedDataCommonTaxExemptNotSupported   EnhancedDataCommonTaxExempt = "NOT_SUPPORTED"
)

func (r EnhancedDataCommonTaxExempt) IsKnown() bool {
	switch r {
	case EnhancedDataCommonTaxExemptTaxIncluded, EnhancedDataCommonTaxExemptTaxNotIncluded, EnhancedDataCommonTaxExemptNotSupported:
		return true
	}
	return false
}

type EnhancedDataFleet struct {
	AmountTotals EnhancedDataFleetAmountTotals `json:"amount_totals,required"`
	Fuel         EnhancedDataFleetFuel         `json:"fuel,required"`
	// The driver number entered into the terminal at the time of sale, with leading
	// zeros stripped.
	DriverNumber string `json:"driver_number"`
	// The odometer reading entered into the terminal at the time of sale.
	Odometer int64 `json:"odometer"`
	// The type of fuel service.
	ServiceType EnhancedDataFleetServiceType `json:"service_type"`
	// The vehicle number entered into the terminal at the time of sale, with leading
	// zeros stripped.
	VehicleNumber string                `json:"vehicle_number"`
	JSON          enhancedDataFleetJSON `json:"-"`
}

// enhancedDataFleetJSON contains the JSON metadata for the struct
// [EnhancedDataFleet]
type enhancedDataFleetJSON struct {
	AmountTotals  apijson.Field
	Fuel          apijson.Field
	DriverNumber  apijson.Field
	Odometer      apijson.Field
	ServiceType   apijson.Field
	VehicleNumber apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnhancedDataFleet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enhancedDataFleetJSON) RawJSON() string {
	return r.raw
}

type EnhancedDataFleetAmountTotals struct {
	// The discount applied to the gross sale amount.
	Discount int64 `json:"discount"`
	// The gross sale amount.
	GrossSale int64 `json:"gross_sale"`
	// The amount after discount.
	NetSale int64                             `json:"net_sale"`
	JSON    enhancedDataFleetAmountTotalsJSON `json:"-"`
}

// enhancedDataFleetAmountTotalsJSON contains the JSON metadata for the struct
// [EnhancedDataFleetAmountTotals]
type enhancedDataFleetAmountTotalsJSON struct {
	Discount    apijson.Field
	GrossSale   apijson.Field
	NetSale     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnhancedDataFleetAmountTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enhancedDataFleetAmountTotalsJSON) RawJSON() string {
	return r.raw
}

type EnhancedDataFleetFuel struct {
	// The quantity of fuel purchased.
	Quantity float64 `json:"quantity"`
	// The type of fuel purchased.
	Type EnhancedDataFleetFuelType `json:"type"`
	// Unit of measure for fuel disbursement.
	UnitOfMeasure EnhancedDataFleetFuelUnitOfMeasure `json:"unit_of_measure"`
	// The price per unit of fuel.
	UnitPrice int64                     `json:"unit_price"`
	JSON      enhancedDataFleetFuelJSON `json:"-"`
}

// enhancedDataFleetFuelJSON contains the JSON metadata for the struct
// [EnhancedDataFleetFuel]
type enhancedDataFleetFuelJSON struct {
	Quantity      apijson.Field
	Type          apijson.Field
	UnitOfMeasure apijson.Field
	UnitPrice     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnhancedDataFleetFuel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enhancedDataFleetFuelJSON) RawJSON() string {
	return r.raw
}

// The type of fuel purchased.
type EnhancedDataFleetFuelType string

const (
	EnhancedDataFleetFuelTypeUnknown                                              EnhancedDataFleetFuelType = "UNKNOWN"
	EnhancedDataFleetFuelTypeRegular                                              EnhancedDataFleetFuelType = "REGULAR"
	EnhancedDataFleetFuelTypeMidPlus                                              EnhancedDataFleetFuelType = "MID_PLUS"
	EnhancedDataFleetFuelTypePremiumSuper                                         EnhancedDataFleetFuelType = "PREMIUM_SUPER"
	EnhancedDataFleetFuelTypeMidPlus2                                             EnhancedDataFleetFuelType = "MID_PLUS_2"
	EnhancedDataFleetFuelTypePremiumSuper2                                        EnhancedDataFleetFuelType = "PREMIUM_SUPER_2"
	EnhancedDataFleetFuelTypeEthanol5_7Blend                                      EnhancedDataFleetFuelType = "ETHANOL_5_7_BLEND"
	EnhancedDataFleetFuelTypeMidPlusEthanol5_7PercentBlend                        EnhancedDataFleetFuelType = "MID_PLUS_ETHANOL_5_7_PERCENT_BLEND"
	EnhancedDataFleetFuelTypePremiumSuperEthanol5_7PercentBlend                   EnhancedDataFleetFuelType = "PREMIUM_SUPER_ETHANOL_5_7_PERCENT_BLEND"
	EnhancedDataFleetFuelTypeEthanol7_7PercentBlend                               EnhancedDataFleetFuelType = "ETHANOL_7_7_PERCENT_BLEND"
	EnhancedDataFleetFuelTypeMidPlusEthanol7_7PercentBlend                        EnhancedDataFleetFuelType = "MID_PLUS_ETHANOL_7_7_PERCENT_BLEND"
	EnhancedDataFleetFuelTypeGreenGasolineRegular                                 EnhancedDataFleetFuelType = "GREEN_GASOLINE_REGULAR"
	EnhancedDataFleetFuelTypeGreenGasolineMidPlus                                 EnhancedDataFleetFuelType = "GREEN_GASOLINE_MID_PLUS"
	EnhancedDataFleetFuelTypeGreenGasolinePremiumSuper                            EnhancedDataFleetFuelType = "GREEN_GASOLINE_PREMIUM_SUPER"
	EnhancedDataFleetFuelTypeRegularDiesel2                                       EnhancedDataFleetFuelType = "REGULAR_DIESEL_2"
	EnhancedDataFleetFuelTypePremiumDiesel2                                       EnhancedDataFleetFuelType = "PREMIUM_DIESEL_2"
	EnhancedDataFleetFuelTypeRegularDiesel1                                       EnhancedDataFleetFuelType = "REGULAR_DIESEL_1"
	EnhancedDataFleetFuelTypeCompressedNaturalGas                                 EnhancedDataFleetFuelType = "COMPRESSED_NATURAL_GAS"
	EnhancedDataFleetFuelTypeLiquidPropaneGas                                     EnhancedDataFleetFuelType = "LIQUID_PROPANE_GAS"
	EnhancedDataFleetFuelTypeLiquidNaturalGas                                     EnhancedDataFleetFuelType = "LIQUID_NATURAL_GAS"
	EnhancedDataFleetFuelTypeE85                                                  EnhancedDataFleetFuelType = "E_85"
	EnhancedDataFleetFuelTypeReformulated1                                        EnhancedDataFleetFuelType = "REFORMULATED_1"
	EnhancedDataFleetFuelTypeReformulated2                                        EnhancedDataFleetFuelType = "REFORMULATED_2"
	EnhancedDataFleetFuelTypeReformulated3                                        EnhancedDataFleetFuelType = "REFORMULATED_3"
	EnhancedDataFleetFuelTypeReformulated4                                        EnhancedDataFleetFuelType = "REFORMULATED_4"
	EnhancedDataFleetFuelTypeReformulated5                                        EnhancedDataFleetFuelType = "REFORMULATED_5"
	EnhancedDataFleetFuelTypeDieselOffRoad1And2NonTaxable                         EnhancedDataFleetFuelType = "DIESEL_OFF_ROAD_1_AND_2_NON_TAXABLE"
	EnhancedDataFleetFuelTypeDieselOffRoadNonTaxable                              EnhancedDataFleetFuelType = "DIESEL_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlendOffRoadNonTaxable                      EnhancedDataFleetFuelType = "BIODIESEL_BLEND_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeUndefinedFuel                                        EnhancedDataFleetFuelType = "UNDEFINED_FUEL"
	EnhancedDataFleetFuelTypeRacingFuel                                           EnhancedDataFleetFuelType = "RACING_FUEL"
	EnhancedDataFleetFuelTypeMidPlus2_10PercentBlend                              EnhancedDataFleetFuelType = "MID_PLUS_2_10_PERCENT_BLEND"
	EnhancedDataFleetFuelTypePremiumSuper2_10PercentBlend                         EnhancedDataFleetFuelType = "PREMIUM_SUPER_2_10_PERCENT_BLEND"
	EnhancedDataFleetFuelTypeMidPlusEthanol2_15PercentBlend                       EnhancedDataFleetFuelType = "MID_PLUS_ETHANOL_2_15_PERCENT_BLEND"
	EnhancedDataFleetFuelTypePremiumSuperEthanol2_15PercentBlend                  EnhancedDataFleetFuelType = "PREMIUM_SUPER_ETHANOL_2_15_PERCENT_BLEND"
	EnhancedDataFleetFuelTypePremiumSuperEthanol7_7PercentBlend                   EnhancedDataFleetFuelType = "PREMIUM_SUPER_ETHANOL_7_7_PERCENT_BLEND"
	EnhancedDataFleetFuelTypeRegularEthanol10PercentBlend                         EnhancedDataFleetFuelType = "REGULAR_ETHANOL_10_PERCENT_BLEND"
	EnhancedDataFleetFuelTypeMidPlusEthanol10PercentBlend                         EnhancedDataFleetFuelType = "MID_PLUS_ETHANOL_10_PERCENT_BLEND"
	EnhancedDataFleetFuelTypePremiumSuperEthanol10PercentBlend                    EnhancedDataFleetFuelType = "PREMIUM_SUPER_ETHANOL_10_PERCENT_BLEND"
	EnhancedDataFleetFuelTypeB2DieselBlend2PercentBiodiesel                       EnhancedDataFleetFuelType = "B2_DIESEL_BLEND_2_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB5DieselBlend5PercentBiodiesel                       EnhancedDataFleetFuelType = "B5_DIESEL_BLEND_5_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB10DieselBlend10PercentBiodiesel                     EnhancedDataFleetFuelType = "B10_DIESEL_BLEND_10_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB11DieselBlend11PercentBiodiesel                     EnhancedDataFleetFuelType = "B11_DIESEL_BLEND_11_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB15DieselBlend15PercentBiodiesel                     EnhancedDataFleetFuelType = "B15_DIESEL_BLEND_15_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB20DieselBlend20PercentBiodiesel                     EnhancedDataFleetFuelType = "B20_DIESEL_BLEND_20_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB100DieselBlend100PercentBiodiesel                   EnhancedDataFleetFuelType = "B100_DIESEL_BLEND_100_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB1DieselBlend1PercentBiodiesel                       EnhancedDataFleetFuelType = "B1_DIESEL_BLEND_1_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeAdditizedDiesel2                                     EnhancedDataFleetFuelType = "ADDITIZED_DIESEL_2"
	EnhancedDataFleetFuelTypeAdditizedDiesel3                                     EnhancedDataFleetFuelType = "ADDITIZED_DIESEL_3"
	EnhancedDataFleetFuelTypeRenewableDieselR95                                   EnhancedDataFleetFuelType = "RENEWABLE_DIESEL_R95"
	EnhancedDataFleetFuelTypeRenewableDieselBiodiesel6_20Percent                  EnhancedDataFleetFuelType = "RENEWABLE_DIESEL_BIODIESEL_6_20_PERCENT"
	EnhancedDataFleetFuelTypeDieselExhaustFluid                                   EnhancedDataFleetFuelType = "DIESEL_EXHAUST_FLUID"
	EnhancedDataFleetFuelTypePremiumDiesel1                                       EnhancedDataFleetFuelType = "PREMIUM_DIESEL_1"
	EnhancedDataFleetFuelTypeRegularEthanol15PercentBlend                         EnhancedDataFleetFuelType = "REGULAR_ETHANOL_15_PERCENT_BLEND"
	EnhancedDataFleetFuelTypeMidPlusEthanol15PercentBlend                         EnhancedDataFleetFuelType = "MID_PLUS_ETHANOL_15_PERCENT_BLEND"
	EnhancedDataFleetFuelTypePremiumSuperEthanol15PercentBlend                    EnhancedDataFleetFuelType = "PREMIUM_SUPER_ETHANOL_15_PERCENT_BLEND"
	EnhancedDataFleetFuelTypePremiumDieselBlendLessThan20PercentBiodiesel         EnhancedDataFleetFuelType = "PREMIUM_DIESEL_BLEND_LESS_THAN_20_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypePremiumDieselBlendGreaterThan20PercentBiodiesel      EnhancedDataFleetFuelType = "PREMIUM_DIESEL_BLEND_GREATER_THAN_20_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB75DieselBlend75PercentBiodiesel                     EnhancedDataFleetFuelType = "B75_DIESEL_BLEND_75_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeB99DieselBlend99PercentBiodiesel                     EnhancedDataFleetFuelType = "B99_DIESEL_BLEND_99_PERCENT_BIODIESEL"
	EnhancedDataFleetFuelTypeMiscellaneousFuel                                    EnhancedDataFleetFuelType = "MISCELLANEOUS_FUEL"
	EnhancedDataFleetFuelTypeJetFuel                                              EnhancedDataFleetFuelType = "JET_FUEL"
	EnhancedDataFleetFuelTypeAviationFuelRegular                                  EnhancedDataFleetFuelType = "AVIATION_FUEL_REGULAR"
	EnhancedDataFleetFuelTypeAviationFuelPremium                                  EnhancedDataFleetFuelType = "AVIATION_FUEL_PREMIUM"
	EnhancedDataFleetFuelTypeAviationFuelJp8                                      EnhancedDataFleetFuelType = "AVIATION_FUEL_JP8"
	EnhancedDataFleetFuelTypeAviationFuel4                                        EnhancedDataFleetFuelType = "AVIATION_FUEL_4"
	EnhancedDataFleetFuelTypeAviationFuel5                                        EnhancedDataFleetFuelType = "AVIATION_FUEL_5"
	EnhancedDataFleetFuelTypeBiojetDiesel                                         EnhancedDataFleetFuelType = "BIOJET_DIESEL"
	EnhancedDataFleetFuelTypeAviationBiofuelGasoline                              EnhancedDataFleetFuelType = "AVIATION_BIOFUEL_GASOLINE"
	EnhancedDataFleetFuelTypeMiscellaneousAviationFuel                            EnhancedDataFleetFuelType = "MISCELLANEOUS_AVIATION_FUEL"
	EnhancedDataFleetFuelTypeMarineFuel1                                          EnhancedDataFleetFuelType = "MARINE_FUEL_1"
	EnhancedDataFleetFuelTypeMarineFuel2                                          EnhancedDataFleetFuelType = "MARINE_FUEL_2"
	EnhancedDataFleetFuelTypeMarineFuel3                                          EnhancedDataFleetFuelType = "MARINE_FUEL_3"
	EnhancedDataFleetFuelTypeMarineFuel4                                          EnhancedDataFleetFuelType = "MARINE_FUEL_4"
	EnhancedDataFleetFuelTypeMarineFuel5                                          EnhancedDataFleetFuelType = "MARINE_FUEL_5"
	EnhancedDataFleetFuelTypeMarineOther                                          EnhancedDataFleetFuelType = "MARINE_OTHER"
	EnhancedDataFleetFuelTypeMarineDiesel                                         EnhancedDataFleetFuelType = "MARINE_DIESEL"
	EnhancedDataFleetFuelTypeMiscellaneousMarineFuel                              EnhancedDataFleetFuelType = "MISCELLANEOUS_MARINE_FUEL"
	EnhancedDataFleetFuelTypeKeroseneLowSulfur                                    EnhancedDataFleetFuelType = "KEROSENE_LOW_SULFUR"
	EnhancedDataFleetFuelTypeWhiteGas                                             EnhancedDataFleetFuelType = "WHITE_GAS"
	EnhancedDataFleetFuelTypeHeatingOil                                           EnhancedDataFleetFuelType = "HEATING_OIL"
	EnhancedDataFleetFuelTypeOtherFuelNonTaxable                                  EnhancedDataFleetFuelType = "OTHER_FUEL_NON_TAXABLE"
	EnhancedDataFleetFuelTypeKeroseneUltraLowSulfur                               EnhancedDataFleetFuelType = "KEROSENE_ULTRA_LOW_SULFUR"
	EnhancedDataFleetFuelTypeKeroseneLowSulfurNonTaxable                          EnhancedDataFleetFuelType = "KEROSENE_LOW_SULFUR_NON_TAXABLE"
	EnhancedDataFleetFuelTypeKeroseneUltraLowSulfurNonTaxable                     EnhancedDataFleetFuelType = "KEROSENE_ULTRA_LOW_SULFUR_NON_TAXABLE"
	EnhancedDataFleetFuelTypeEvc1Level1Charge110V15Amp                            EnhancedDataFleetFuelType = "EVC_1_LEVEL_1_CHARGE_110V_15_AMP"
	EnhancedDataFleetFuelTypeEvc2Level2Charge240V15_40Amp                         EnhancedDataFleetFuelType = "EVC_2_LEVEL_2_CHARGE_240V_15_40_AMP"
	EnhancedDataFleetFuelTypeEvc3Level3Charge480V3PhaseCharge                     EnhancedDataFleetFuelType = "EVC_3_LEVEL_3_CHARGE_480V_3_PHASE_CHARGE"
	EnhancedDataFleetFuelTypeBiodieselBlend2PercentOffRoadNonTaxable              EnhancedDataFleetFuelType = "BIODIESEL_BLEND_2_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend5PercentOffRoadNonTaxable              EnhancedDataFleetFuelType = "BIODIESEL_BLEND_5_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend10PercentOffRoadNonTaxable             EnhancedDataFleetFuelType = "BIODIESEL_BLEND_10_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend11PercentOffRoadNonTaxable             EnhancedDataFleetFuelType = "BIODIESEL_BLEND_11_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend15PercentOffRoadNonTaxable             EnhancedDataFleetFuelType = "BIODIESEL_BLEND_15_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend20PercentOffRoadNonTaxable             EnhancedDataFleetFuelType = "BIODIESEL_BLEND_20_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeDiesel1OffRoadNonTaxable                             EnhancedDataFleetFuelType = "DIESEL_1_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeDiesel2OffRoadNonTaxable                             EnhancedDataFleetFuelType = "DIESEL_2_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeDiesel1PremiumOffRoadNonTaxable                      EnhancedDataFleetFuelType = "DIESEL_1_PREMIUM_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeDiesel2PremiumOffRoadNonTaxable                      EnhancedDataFleetFuelType = "DIESEL_2_PREMIUM_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeAdditiveDosage                                       EnhancedDataFleetFuelType = "ADDITIVE_DOSAGE"
	EnhancedDataFleetFuelTypeEthanolBlendsE16E84                                  EnhancedDataFleetFuelType = "ETHANOL_BLENDS_E16_E84"
	EnhancedDataFleetFuelTypeLowOctaneUnl                                         EnhancedDataFleetFuelType = "LOW_OCTANE_UNL"
	EnhancedDataFleetFuelTypeBlendedDiesel1And2                                   EnhancedDataFleetFuelType = "BLENDED_DIESEL_1_AND_2"
	EnhancedDataFleetFuelTypeOffRoadRegularNonTaxable                             EnhancedDataFleetFuelType = "OFF_ROAD_REGULAR_NON_TAXABLE"
	EnhancedDataFleetFuelTypeOffRoadMidPlusNonTaxable                             EnhancedDataFleetFuelType = "OFF_ROAD_MID_PLUS_NON_TAXABLE"
	EnhancedDataFleetFuelTypeOffRoadPremiumSuperNonTaxable                        EnhancedDataFleetFuelType = "OFF_ROAD_PREMIUM_SUPER_NON_TAXABLE"
	EnhancedDataFleetFuelTypeOffRoadMidPlus2NonTaxable                            EnhancedDataFleetFuelType = "OFF_ROAD_MID_PLUS_2_NON_TAXABLE"
	EnhancedDataFleetFuelTypeOffRoadPremiumSuper2NonTaxable                       EnhancedDataFleetFuelType = "OFF_ROAD_PREMIUM_SUPER_2_NON_TAXABLE"
	EnhancedDataFleetFuelTypeRecreationalFuel90Octane                             EnhancedDataFleetFuelType = "RECREATIONAL_FUEL_90_OCTANE"
	EnhancedDataFleetFuelTypeHydrogenH35                                          EnhancedDataFleetFuelType = "HYDROGEN_H35"
	EnhancedDataFleetFuelTypeHydrogenH70                                          EnhancedDataFleetFuelType = "HYDROGEN_H70"
	EnhancedDataFleetFuelTypeRenewableDieselR95OffRoadNonTaxable                  EnhancedDataFleetFuelType = "RENEWABLE_DIESEL_R95_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend1PercentOffRoadNonTaxable              EnhancedDataFleetFuelType = "BIODIESEL_BLEND_1_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend75PercentOffRoadNonTaxable             EnhancedDataFleetFuelType = "BIODIESEL_BLEND_75_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend99PercentOffRoadNonTaxable             EnhancedDataFleetFuelType = "BIODIESEL_BLEND_99_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeBiodieselBlend100PercentOffRoadNonTaxable            EnhancedDataFleetFuelType = "BIODIESEL_BLEND_100_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeRenewableDieselBiodiesel6_20PercentOffRoadNonTaxable EnhancedDataFleetFuelType = "RENEWABLE_DIESEL_BIODIESEL_6_20_PERCENT_OFF_ROAD_NON_TAXABLE"
	EnhancedDataFleetFuelTypeMiscellaneousOtherFuel                               EnhancedDataFleetFuelType = "MISCELLANEOUS_OTHER_FUEL"
)

func (r EnhancedDataFleetFuelType) IsKnown() bool {
	switch r {
	case EnhancedDataFleetFuelTypeUnknown, EnhancedDataFleetFuelTypeRegular, EnhancedDataFleetFuelTypeMidPlus, EnhancedDataFleetFuelTypePremiumSuper, EnhancedDataFleetFuelTypeMidPlus2, EnhancedDataFleetFuelTypePremiumSuper2, EnhancedDataFleetFuelTypeEthanol5_7Blend, EnhancedDataFleetFuelTypeMidPlusEthanol5_7PercentBlend, EnhancedDataFleetFuelTypePremiumSuperEthanol5_7PercentBlend, EnhancedDataFleetFuelTypeEthanol7_7PercentBlend, EnhancedDataFleetFuelTypeMidPlusEthanol7_7PercentBlend, EnhancedDataFleetFuelTypeGreenGasolineRegular, EnhancedDataFleetFuelTypeGreenGasolineMidPlus, EnhancedDataFleetFuelTypeGreenGasolinePremiumSuper, EnhancedDataFleetFuelTypeRegularDiesel2, EnhancedDataFleetFuelTypePremiumDiesel2, EnhancedDataFleetFuelTypeRegularDiesel1, EnhancedDataFleetFuelTypeCompressedNaturalGas, EnhancedDataFleetFuelTypeLiquidPropaneGas, EnhancedDataFleetFuelTypeLiquidNaturalGas, EnhancedDataFleetFuelTypeE85, EnhancedDataFleetFuelTypeReformulated1, EnhancedDataFleetFuelTypeReformulated2, EnhancedDataFleetFuelTypeReformulated3, EnhancedDataFleetFuelTypeReformulated4, EnhancedDataFleetFuelTypeReformulated5, EnhancedDataFleetFuelTypeDieselOffRoad1And2NonTaxable, EnhancedDataFleetFuelTypeDieselOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlendOffRoadNonTaxable, EnhancedDataFleetFuelTypeUndefinedFuel, EnhancedDataFleetFuelTypeRacingFuel, EnhancedDataFleetFuelTypeMidPlus2_10PercentBlend, EnhancedDataFleetFuelTypePremiumSuper2_10PercentBlend, EnhancedDataFleetFuelTypeMidPlusEthanol2_15PercentBlend, EnhancedDataFleetFuelTypePremiumSuperEthanol2_15PercentBlend, EnhancedDataFleetFuelTypePremiumSuperEthanol7_7PercentBlend, EnhancedDataFleetFuelTypeRegularEthanol10PercentBlend, EnhancedDataFleetFuelTypeMidPlusEthanol10PercentBlend, EnhancedDataFleetFuelTypePremiumSuperEthanol10PercentBlend, EnhancedDataFleetFuelTypeB2DieselBlend2PercentBiodiesel, EnhancedDataFleetFuelTypeB5DieselBlend5PercentBiodiesel, EnhancedDataFleetFuelTypeB10DieselBlend10PercentBiodiesel, EnhancedDataFleetFuelTypeB11DieselBlend11PercentBiodiesel, EnhancedDataFleetFuelTypeB15DieselBlend15PercentBiodiesel, EnhancedDataFleetFuelTypeB20DieselBlend20PercentBiodiesel, EnhancedDataFleetFuelTypeB100DieselBlend100PercentBiodiesel, EnhancedDataFleetFuelTypeB1DieselBlend1PercentBiodiesel, EnhancedDataFleetFuelTypeAdditizedDiesel2, EnhancedDataFleetFuelTypeAdditizedDiesel3, EnhancedDataFleetFuelTypeRenewableDieselR95, EnhancedDataFleetFuelTypeRenewableDieselBiodiesel6_20Percent, EnhancedDataFleetFuelTypeDieselExhaustFluid, EnhancedDataFleetFuelTypePremiumDiesel1, EnhancedDataFleetFuelTypeRegularEthanol15PercentBlend, EnhancedDataFleetFuelTypeMidPlusEthanol15PercentBlend, EnhancedDataFleetFuelTypePremiumSuperEthanol15PercentBlend, EnhancedDataFleetFuelTypePremiumDieselBlendLessThan20PercentBiodiesel, EnhancedDataFleetFuelTypePremiumDieselBlendGreaterThan20PercentBiodiesel, EnhancedDataFleetFuelTypeB75DieselBlend75PercentBiodiesel, EnhancedDataFleetFuelTypeB99DieselBlend99PercentBiodiesel, EnhancedDataFleetFuelTypeMiscellaneousFuel, EnhancedDataFleetFuelTypeJetFuel, EnhancedDataFleetFuelTypeAviationFuelRegular, EnhancedDataFleetFuelTypeAviationFuelPremium, EnhancedDataFleetFuelTypeAviationFuelJp8, EnhancedDataFleetFuelTypeAviationFuel4, EnhancedDataFleetFuelTypeAviationFuel5, EnhancedDataFleetFuelTypeBiojetDiesel, EnhancedDataFleetFuelTypeAviationBiofuelGasoline, EnhancedDataFleetFuelTypeMiscellaneousAviationFuel, EnhancedDataFleetFuelTypeMarineFuel1, EnhancedDataFleetFuelTypeMarineFuel2, EnhancedDataFleetFuelTypeMarineFuel3, EnhancedDataFleetFuelTypeMarineFuel4, EnhancedDataFleetFuelTypeMarineFuel5, EnhancedDataFleetFuelTypeMarineOther, EnhancedDataFleetFuelTypeMarineDiesel, EnhancedDataFleetFuelTypeMiscellaneousMarineFuel, EnhancedDataFleetFuelTypeKeroseneLowSulfur, EnhancedDataFleetFuelTypeWhiteGas, EnhancedDataFleetFuelTypeHeatingOil, EnhancedDataFleetFuelTypeOtherFuelNonTaxable, EnhancedDataFleetFuelTypeKeroseneUltraLowSulfur, EnhancedDataFleetFuelTypeKeroseneLowSulfurNonTaxable, EnhancedDataFleetFuelTypeKeroseneUltraLowSulfurNonTaxable, EnhancedDataFleetFuelTypeEvc1Level1Charge110V15Amp, EnhancedDataFleetFuelTypeEvc2Level2Charge240V15_40Amp, EnhancedDataFleetFuelTypeEvc3Level3Charge480V3PhaseCharge, EnhancedDataFleetFuelTypeBiodieselBlend2PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend5PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend10PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend11PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend15PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend20PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeDiesel1OffRoadNonTaxable, EnhancedDataFleetFuelTypeDiesel2OffRoadNonTaxable, EnhancedDataFleetFuelTypeDiesel1PremiumOffRoadNonTaxable, EnhancedDataFleetFuelTypeDiesel2PremiumOffRoadNonTaxable, EnhancedDataFleetFuelTypeAdditiveDosage, EnhancedDataFleetFuelTypeEthanolBlendsE16E84, EnhancedDataFleetFuelTypeLowOctaneUnl, EnhancedDataFleetFuelTypeBlendedDiesel1And2, EnhancedDataFleetFuelTypeOffRoadRegularNonTaxable, EnhancedDataFleetFuelTypeOffRoadMidPlusNonTaxable, EnhancedDataFleetFuelTypeOffRoadPremiumSuperNonTaxable, EnhancedDataFleetFuelTypeOffRoadMidPlus2NonTaxable, EnhancedDataFleetFuelTypeOffRoadPremiumSuper2NonTaxable, EnhancedDataFleetFuelTypeRecreationalFuel90Octane, EnhancedDataFleetFuelTypeHydrogenH35, EnhancedDataFleetFuelTypeHydrogenH70, EnhancedDataFleetFuelTypeRenewableDieselR95OffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend1PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend75PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend99PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeBiodieselBlend100PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeRenewableDieselBiodiesel6_20PercentOffRoadNonTaxable, EnhancedDataFleetFuelTypeMiscellaneousOtherFuel:
		return true
	}
	return false
}

// Unit of measure for fuel disbursement.
type EnhancedDataFleetFuelUnitOfMeasure string

const (
	EnhancedDataFleetFuelUnitOfMeasureGallons         EnhancedDataFleetFuelUnitOfMeasure = "GALLONS"
	EnhancedDataFleetFuelUnitOfMeasureLiters          EnhancedDataFleetFuelUnitOfMeasure = "LITERS"
	EnhancedDataFleetFuelUnitOfMeasurePounds          EnhancedDataFleetFuelUnitOfMeasure = "POUNDS"
	EnhancedDataFleetFuelUnitOfMeasureKilograms       EnhancedDataFleetFuelUnitOfMeasure = "KILOGRAMS"
	EnhancedDataFleetFuelUnitOfMeasureImperialGallons EnhancedDataFleetFuelUnitOfMeasure = "IMPERIAL_GALLONS"
	EnhancedDataFleetFuelUnitOfMeasureNotApplicable   EnhancedDataFleetFuelUnitOfMeasure = "NOT_APPLICABLE"
	EnhancedDataFleetFuelUnitOfMeasureUnknown         EnhancedDataFleetFuelUnitOfMeasure = "UNKNOWN"
)

func (r EnhancedDataFleetFuelUnitOfMeasure) IsKnown() bool {
	switch r {
	case EnhancedDataFleetFuelUnitOfMeasureGallons, EnhancedDataFleetFuelUnitOfMeasureLiters, EnhancedDataFleetFuelUnitOfMeasurePounds, EnhancedDataFleetFuelUnitOfMeasureKilograms, EnhancedDataFleetFuelUnitOfMeasureImperialGallons, EnhancedDataFleetFuelUnitOfMeasureNotApplicable, EnhancedDataFleetFuelUnitOfMeasureUnknown:
		return true
	}
	return false
}

// The type of fuel service.
type EnhancedDataFleetServiceType string

const (
	EnhancedDataFleetServiceTypeUnknown     EnhancedDataFleetServiceType = "UNKNOWN"
	EnhancedDataFleetServiceTypeUndefined   EnhancedDataFleetServiceType = "UNDEFINED"
	EnhancedDataFleetServiceTypeSelfService EnhancedDataFleetServiceType = "SELF_SERVICE"
	EnhancedDataFleetServiceTypeFullService EnhancedDataFleetServiceType = "FULL_SERVICE"
	EnhancedDataFleetServiceTypeNonFuelOnly EnhancedDataFleetServiceType = "NON_FUEL_ONLY"
)

func (r EnhancedDataFleetServiceType) IsKnown() bool {
	switch r {
	case EnhancedDataFleetServiceTypeUnknown, EnhancedDataFleetServiceTypeUndefined, EnhancedDataFleetServiceTypeSelfService, EnhancedDataFleetServiceTypeFullService, EnhancedDataFleetServiceTypeNonFuelOnly:
		return true
	}
	return false
}
