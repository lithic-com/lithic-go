// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
	"github.com/lithic-com/lithic-go/shared"
	"github.com/tidwall/gjson"
)

// DisputesV2Service contains methods and other services that help with interacting
// with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDisputesV2Service] method instead.
type DisputesV2Service struct {
	Options []option.RequestOption
}

// NewDisputesV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDisputesV2Service(opts ...option.RequestOption) (r *DisputesV2Service) {
	r = &DisputesV2Service{}
	r.Options = opts
	return
}

// Retrieves a specific dispute by its token.
func (r *DisputesV2Service) Get(ctx context.Context, disputeToken string, opts ...option.RequestOption) (res *DisputeV2, err error) {
	opts = slices.Concat(r.Options, opts)
	if disputeToken == "" {
		err = errors.New("missing required dispute_token parameter")
		return
	}
	path := fmt.Sprintf("v2/disputes/%s", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of disputes.
func (r *DisputesV2Service) List(ctx context.Context, query DisputesV2ListParams, opts ...option.RequestOption) (res *pagination.CursorPage[DisputeV2], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/disputes"
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

// Returns a paginated list of disputes.
func (r *DisputesV2Service) ListAutoPaging(ctx context.Context, query DisputesV2ListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[DisputeV2] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Dispute object tracks the progression of a dispute throughout its lifecycle.
type DisputeV2 struct {
	// Token assigned by Lithic for the dispute, in UUID format.
	Token string `json:"token,required" format:"uuid"`
	// Token for the account associated with the dispute, in UUID format.
	AccountToken string `json:"account_token,required" format:"uuid"`
	// Token for the card used in the dispute, in UUID format.
	CardToken string `json:"card_token,required" format:"uuid"`
	// Identifier assigned by the network for this dispute.
	CaseID string `json:"case_id,required,nullable"`
	// When the dispute was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Three-letter ISO 4217 currency code.
	Currency string `json:"currency,required"`
	// Dispute resolution outcome
	Disposition DisputeV2Disposition `json:"disposition,required,nullable"`
	// Chronological list of events that have occurred in the dispute lifecycle
	Events []DisputeV2Event `json:"events,required"`
	// Current breakdown of how liability is allocated for the disputed amount
	LiabilityAllocation DisputeV2LiabilityAllocation `json:"liability_allocation,required"`
	Merchant            shared.Merchant              `json:"merchant,required"`
	// Card network handling the dispute.
	Network DisputeV2Network `json:"network,required"`
	// Current status of the dispute.
	Status DisputeV2Status `json:"status,required,nullable"`
	// Contains identifiers for the transaction and specific event within being
	// disputed; null if no transaction can be identified
	TransactionSeries DisputeV2TransactionSeries `json:"transaction_series,required,nullable"`
	// When the dispute was last updated.
	Updated time.Time     `json:"updated,required" format:"date-time"`
	JSON    disputeV2JSON `json:"-"`
}

// disputeV2JSON contains the JSON metadata for the struct [DisputeV2]
type disputeV2JSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardToken           apijson.Field
	CaseID              apijson.Field
	Created             apijson.Field
	Currency            apijson.Field
	Disposition         apijson.Field
	Events              apijson.Field
	LiabilityAllocation apijson.Field
	Merchant            apijson.Field
	Network             apijson.Field
	Status              apijson.Field
	TransactionSeries   apijson.Field
	Updated             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DisputeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2JSON) RawJSON() string {
	return r.raw
}

// Dispute resolution outcome
type DisputeV2Disposition string

const (
	DisputeV2DispositionWon          DisputeV2Disposition = "WON"
	DisputeV2DispositionLost         DisputeV2Disposition = "LOST"
	DisputeV2DispositionPartiallyWon DisputeV2Disposition = "PARTIALLY_WON"
	DisputeV2DispositionWithdrawn    DisputeV2Disposition = "WITHDRAWN"
	DisputeV2DispositionDenied       DisputeV2Disposition = "DENIED"
)

func (r DisputeV2Disposition) IsKnown() bool {
	switch r {
	case DisputeV2DispositionWon, DisputeV2DispositionLost, DisputeV2DispositionPartiallyWon, DisputeV2DispositionWithdrawn, DisputeV2DispositionDenied:
		return true
	}
	return false
}

// Event that occurred in the dispute lifecycle
type DisputeV2Event struct {
	// Unique identifier for the event, in UUID format
	Token string `json:"token,required" format:"uuid"`
	// When the event occurred
	Created time.Time `json:"created,required" format:"date-time"`
	// Details specific to the event type
	Data DisputeV2EventsData `json:"data,required"`
	// Type of event
	Type DisputeV2EventsType `json:"type,required"`
	JSON disputeV2EventJSON  `json:"-"`
}

// disputeV2EventJSON contains the JSON metadata for the struct [DisputeV2Event]
type disputeV2EventJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventJSON) RawJSON() string {
	return r.raw
}

// Details specific to the event type
type DisputeV2EventsData struct {
	// Amount in minor units
	Amount int64 `json:"amount,required,nullable"`
	// Event type discriminator
	Type DisputeV2EventsDataType `json:"type,required"`
	// Action taken in this stage
	Action DisputeV2EventsDataAction `json:"action"`
	// Dispute resolution outcome
	Disposition DisputeV2EventsDataDisposition `json:"disposition,nullable"`
	// Direction of funds flow
	Polarity DisputeV2EventsDataPolarity `json:"polarity"`
	// Reason for the action
	Reason string `json:"reason,nullable"`
	// Current stage of the dispute workflow
	Stage DisputeV2EventsDataStage `json:"stage"`
	JSON  disputeV2EventsDataJSON  `json:"-"`
	union DisputeV2EventsDataUnion
}

// disputeV2EventsDataJSON contains the JSON metadata for the struct
// [DisputeV2EventsData]
type disputeV2EventsDataJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	Action      apijson.Field
	Disposition apijson.Field
	Polarity    apijson.Field
	Reason      apijson.Field
	Stage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r disputeV2EventsDataJSON) RawJSON() string {
	return r.raw
}

func (r *DisputeV2EventsData) UnmarshalJSON(data []byte) (err error) {
	*r = DisputeV2EventsData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DisputeV2EventsDataUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DisputeV2EventsDataWorkflow],
// [DisputeV2EventsDataFinancial], [DisputeV2EventsDataCardholderLiability].
func (r DisputeV2EventsData) AsUnion() DisputeV2EventsDataUnion {
	return r.union
}

// Details specific to the event type
//
// Union satisfied by [DisputeV2EventsDataWorkflow], [DisputeV2EventsDataFinancial]
// or [DisputeV2EventsDataCardholderLiability].
type DisputeV2EventsDataUnion interface {
	implementsDisputeV2EventsData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DisputeV2EventsDataUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DisputeV2EventsDataWorkflow{}),
			DiscriminatorValue: "WORKFLOW",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DisputeV2EventsDataFinancial{}),
			DiscriminatorValue: "FINANCIAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DisputeV2EventsDataCardholderLiability{}),
			DiscriminatorValue: "CARDHOLDER_LIABILITY",
		},
	)
}

// Details specific to workflow events
type DisputeV2EventsDataWorkflow struct {
	// Action taken in this stage
	Action DisputeV2EventsDataWorkflowAction `json:"action,required"`
	// Amount in minor units
	Amount int64 `json:"amount,required,nullable"`
	// Dispute resolution outcome
	Disposition DisputeV2EventsDataWorkflowDisposition `json:"disposition,required,nullable"`
	// Reason for the action
	Reason string `json:"reason,required,nullable"`
	// Current stage of the dispute workflow
	Stage DisputeV2EventsDataWorkflowStage `json:"stage,required"`
	// Event type discriminator
	Type DisputeV2EventsDataWorkflowType `json:"type,required"`
	JSON disputeV2EventsDataWorkflowJSON `json:"-"`
}

// disputeV2EventsDataWorkflowJSON contains the JSON metadata for the struct
// [DisputeV2EventsDataWorkflow]
type disputeV2EventsDataWorkflowJSON struct {
	Action      apijson.Field
	Amount      apijson.Field
	Disposition apijson.Field
	Reason      apijson.Field
	Stage       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsDataWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsDataWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r DisputeV2EventsDataWorkflow) implementsDisputeV2EventsData() {}

// Action taken in this stage
type DisputeV2EventsDataWorkflowAction string

const (
	DisputeV2EventsDataWorkflowActionOpened   DisputeV2EventsDataWorkflowAction = "OPENED"
	DisputeV2EventsDataWorkflowActionClosed   DisputeV2EventsDataWorkflowAction = "CLOSED"
	DisputeV2EventsDataWorkflowActionReopened DisputeV2EventsDataWorkflowAction = "REOPENED"
)

func (r DisputeV2EventsDataWorkflowAction) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataWorkflowActionOpened, DisputeV2EventsDataWorkflowActionClosed, DisputeV2EventsDataWorkflowActionReopened:
		return true
	}
	return false
}

// Dispute resolution outcome
type DisputeV2EventsDataWorkflowDisposition string

const (
	DisputeV2EventsDataWorkflowDispositionWon          DisputeV2EventsDataWorkflowDisposition = "WON"
	DisputeV2EventsDataWorkflowDispositionLost         DisputeV2EventsDataWorkflowDisposition = "LOST"
	DisputeV2EventsDataWorkflowDispositionPartiallyWon DisputeV2EventsDataWorkflowDisposition = "PARTIALLY_WON"
	DisputeV2EventsDataWorkflowDispositionWithdrawn    DisputeV2EventsDataWorkflowDisposition = "WITHDRAWN"
	DisputeV2EventsDataWorkflowDispositionDenied       DisputeV2EventsDataWorkflowDisposition = "DENIED"
)

func (r DisputeV2EventsDataWorkflowDisposition) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataWorkflowDispositionWon, DisputeV2EventsDataWorkflowDispositionLost, DisputeV2EventsDataWorkflowDispositionPartiallyWon, DisputeV2EventsDataWorkflowDispositionWithdrawn, DisputeV2EventsDataWorkflowDispositionDenied:
		return true
	}
	return false
}

// Current stage of the dispute workflow
type DisputeV2EventsDataWorkflowStage string

const (
	DisputeV2EventsDataWorkflowStageClaim DisputeV2EventsDataWorkflowStage = "CLAIM"
)

func (r DisputeV2EventsDataWorkflowStage) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataWorkflowStageClaim:
		return true
	}
	return false
}

// Event type discriminator
type DisputeV2EventsDataWorkflowType string

const (
	DisputeV2EventsDataWorkflowTypeWorkflow DisputeV2EventsDataWorkflowType = "WORKFLOW"
)

func (r DisputeV2EventsDataWorkflowType) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataWorkflowTypeWorkflow:
		return true
	}
	return false
}

// Details specific to financial events
type DisputeV2EventsDataFinancial struct {
	// Amount in minor units
	Amount int64 `json:"amount,required"`
	// Direction of funds flow
	Polarity DisputeV2EventsDataFinancialPolarity `json:"polarity,required"`
	// Stage at which the financial event occurred
	Stage DisputeV2EventsDataFinancialStage `json:"stage,required"`
	// Event type discriminator
	Type DisputeV2EventsDataFinancialType `json:"type,required"`
	JSON disputeV2EventsDataFinancialJSON `json:"-"`
}

// disputeV2EventsDataFinancialJSON contains the JSON metadata for the struct
// [DisputeV2EventsDataFinancial]
type disputeV2EventsDataFinancialJSON struct {
	Amount      apijson.Field
	Polarity    apijson.Field
	Stage       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsDataFinancial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsDataFinancialJSON) RawJSON() string {
	return r.raw
}

func (r DisputeV2EventsDataFinancial) implementsDisputeV2EventsData() {}

// Direction of funds flow
type DisputeV2EventsDataFinancialPolarity string

const (
	DisputeV2EventsDataFinancialPolarityCredit DisputeV2EventsDataFinancialPolarity = "CREDIT"
	DisputeV2EventsDataFinancialPolarityDebit  DisputeV2EventsDataFinancialPolarity = "DEBIT"
)

func (r DisputeV2EventsDataFinancialPolarity) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataFinancialPolarityCredit, DisputeV2EventsDataFinancialPolarityDebit:
		return true
	}
	return false
}

// Stage at which the financial event occurred
type DisputeV2EventsDataFinancialStage string

const (
	DisputeV2EventsDataFinancialStageChargeback     DisputeV2EventsDataFinancialStage = "CHARGEBACK"
	DisputeV2EventsDataFinancialStageRepresentment  DisputeV2EventsDataFinancialStage = "REPRESENTMENT"
	DisputeV2EventsDataFinancialStagePrearbitration DisputeV2EventsDataFinancialStage = "PREARBITRATION"
	DisputeV2EventsDataFinancialStageArbitration    DisputeV2EventsDataFinancialStage = "ARBITRATION"
	DisputeV2EventsDataFinancialStageCollaboration  DisputeV2EventsDataFinancialStage = "COLLABORATION"
)

func (r DisputeV2EventsDataFinancialStage) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataFinancialStageChargeback, DisputeV2EventsDataFinancialStageRepresentment, DisputeV2EventsDataFinancialStagePrearbitration, DisputeV2EventsDataFinancialStageArbitration, DisputeV2EventsDataFinancialStageCollaboration:
		return true
	}
	return false
}

// Event type discriminator
type DisputeV2EventsDataFinancialType string

const (
	DisputeV2EventsDataFinancialTypeFinancial DisputeV2EventsDataFinancialType = "FINANCIAL"
)

func (r DisputeV2EventsDataFinancialType) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataFinancialTypeFinancial:
		return true
	}
	return false
}

// Details specific to cardholder liability events
type DisputeV2EventsDataCardholderLiability struct {
	// Action taken regarding cardholder liability
	Action DisputeV2EventsDataCardholderLiabilityAction `json:"action,required"`
	// Amount in minor units
	Amount int64 `json:"amount,required"`
	// Reason for the action
	Reason string `json:"reason,required"`
	// Event type discriminator
	Type DisputeV2EventsDataCardholderLiabilityType `json:"type,required"`
	JSON disputeV2EventsDataCardholderLiabilityJSON `json:"-"`
}

// disputeV2EventsDataCardholderLiabilityJSON contains the JSON metadata for the
// struct [DisputeV2EventsDataCardholderLiability]
type disputeV2EventsDataCardholderLiabilityJSON struct {
	Action      apijson.Field
	Amount      apijson.Field
	Reason      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsDataCardholderLiability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsDataCardholderLiabilityJSON) RawJSON() string {
	return r.raw
}

func (r DisputeV2EventsDataCardholderLiability) implementsDisputeV2EventsData() {}

// Action taken regarding cardholder liability
type DisputeV2EventsDataCardholderLiabilityAction string

const (
	DisputeV2EventsDataCardholderLiabilityActionProvisionalCreditGranted  DisputeV2EventsDataCardholderLiabilityAction = "PROVISIONAL_CREDIT_GRANTED"
	DisputeV2EventsDataCardholderLiabilityActionProvisionalCreditReversed DisputeV2EventsDataCardholderLiabilityAction = "PROVISIONAL_CREDIT_REVERSED"
	DisputeV2EventsDataCardholderLiabilityActionWrittenOff                DisputeV2EventsDataCardholderLiabilityAction = "WRITTEN_OFF"
)

func (r DisputeV2EventsDataCardholderLiabilityAction) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataCardholderLiabilityActionProvisionalCreditGranted, DisputeV2EventsDataCardholderLiabilityActionProvisionalCreditReversed, DisputeV2EventsDataCardholderLiabilityActionWrittenOff:
		return true
	}
	return false
}

// Event type discriminator
type DisputeV2EventsDataCardholderLiabilityType string

const (
	DisputeV2EventsDataCardholderLiabilityTypeCardholderLiability DisputeV2EventsDataCardholderLiabilityType = "CARDHOLDER_LIABILITY"
)

func (r DisputeV2EventsDataCardholderLiabilityType) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataCardholderLiabilityTypeCardholderLiability:
		return true
	}
	return false
}

// Event type discriminator
type DisputeV2EventsDataType string

const (
	DisputeV2EventsDataTypeWorkflow            DisputeV2EventsDataType = "WORKFLOW"
	DisputeV2EventsDataTypeFinancial           DisputeV2EventsDataType = "FINANCIAL"
	DisputeV2EventsDataTypeCardholderLiability DisputeV2EventsDataType = "CARDHOLDER_LIABILITY"
)

func (r DisputeV2EventsDataType) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataTypeWorkflow, DisputeV2EventsDataTypeFinancial, DisputeV2EventsDataTypeCardholderLiability:
		return true
	}
	return false
}

// Action taken in this stage
type DisputeV2EventsDataAction string

const (
	DisputeV2EventsDataActionOpened                    DisputeV2EventsDataAction = "OPENED"
	DisputeV2EventsDataActionClosed                    DisputeV2EventsDataAction = "CLOSED"
	DisputeV2EventsDataActionReopened                  DisputeV2EventsDataAction = "REOPENED"
	DisputeV2EventsDataActionProvisionalCreditGranted  DisputeV2EventsDataAction = "PROVISIONAL_CREDIT_GRANTED"
	DisputeV2EventsDataActionProvisionalCreditReversed DisputeV2EventsDataAction = "PROVISIONAL_CREDIT_REVERSED"
	DisputeV2EventsDataActionWrittenOff                DisputeV2EventsDataAction = "WRITTEN_OFF"
)

func (r DisputeV2EventsDataAction) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataActionOpened, DisputeV2EventsDataActionClosed, DisputeV2EventsDataActionReopened, DisputeV2EventsDataActionProvisionalCreditGranted, DisputeV2EventsDataActionProvisionalCreditReversed, DisputeV2EventsDataActionWrittenOff:
		return true
	}
	return false
}

// Dispute resolution outcome
type DisputeV2EventsDataDisposition string

const (
	DisputeV2EventsDataDispositionWon          DisputeV2EventsDataDisposition = "WON"
	DisputeV2EventsDataDispositionLost         DisputeV2EventsDataDisposition = "LOST"
	DisputeV2EventsDataDispositionPartiallyWon DisputeV2EventsDataDisposition = "PARTIALLY_WON"
	DisputeV2EventsDataDispositionWithdrawn    DisputeV2EventsDataDisposition = "WITHDRAWN"
	DisputeV2EventsDataDispositionDenied       DisputeV2EventsDataDisposition = "DENIED"
)

func (r DisputeV2EventsDataDisposition) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataDispositionWon, DisputeV2EventsDataDispositionLost, DisputeV2EventsDataDispositionPartiallyWon, DisputeV2EventsDataDispositionWithdrawn, DisputeV2EventsDataDispositionDenied:
		return true
	}
	return false
}

// Direction of funds flow
type DisputeV2EventsDataPolarity string

const (
	DisputeV2EventsDataPolarityCredit DisputeV2EventsDataPolarity = "CREDIT"
	DisputeV2EventsDataPolarityDebit  DisputeV2EventsDataPolarity = "DEBIT"
)

func (r DisputeV2EventsDataPolarity) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataPolarityCredit, DisputeV2EventsDataPolarityDebit:
		return true
	}
	return false
}

// Current stage of the dispute workflow
type DisputeV2EventsDataStage string

const (
	DisputeV2EventsDataStageClaim          DisputeV2EventsDataStage = "CLAIM"
	DisputeV2EventsDataStageChargeback     DisputeV2EventsDataStage = "CHARGEBACK"
	DisputeV2EventsDataStageRepresentment  DisputeV2EventsDataStage = "REPRESENTMENT"
	DisputeV2EventsDataStagePrearbitration DisputeV2EventsDataStage = "PREARBITRATION"
	DisputeV2EventsDataStageArbitration    DisputeV2EventsDataStage = "ARBITRATION"
	DisputeV2EventsDataStageCollaboration  DisputeV2EventsDataStage = "COLLABORATION"
)

func (r DisputeV2EventsDataStage) IsKnown() bool {
	switch r {
	case DisputeV2EventsDataStageClaim, DisputeV2EventsDataStageChargeback, DisputeV2EventsDataStageRepresentment, DisputeV2EventsDataStagePrearbitration, DisputeV2EventsDataStageArbitration, DisputeV2EventsDataStageCollaboration:
		return true
	}
	return false
}

// Type of event
type DisputeV2EventsType string

const (
	DisputeV2EventsTypeWorkflow            DisputeV2EventsType = "WORKFLOW"
	DisputeV2EventsTypeFinancial           DisputeV2EventsType = "FINANCIAL"
	DisputeV2EventsTypeCardholderLiability DisputeV2EventsType = "CARDHOLDER_LIABILITY"
)

func (r DisputeV2EventsType) IsKnown() bool {
	switch r {
	case DisputeV2EventsTypeWorkflow, DisputeV2EventsTypeFinancial, DisputeV2EventsTypeCardholderLiability:
		return true
	}
	return false
}

// Current breakdown of how liability is allocated for the disputed amount
type DisputeV2LiabilityAllocation struct {
	// The amount that has been denied to the cardholder
	DeniedAmount int64 `json:"denied_amount,required"`
	// The initial amount disputed
	OriginalAmount int64 `json:"original_amount,required"`
	// The amount that has been recovered from the merchant through the dispute process
	RecoveredAmount int64 `json:"recovered_amount,required"`
	// Any disputed amount that is still outstanding, i.e. has not been recovered,
	// written off, or denied
	RemainingAmount int64 `json:"remaining_amount,required"`
	// The amount the issuer has chosen to write off
	WrittenOffAmount int64                            `json:"written_off_amount,required"`
	JSON             disputeV2LiabilityAllocationJSON `json:"-"`
}

// disputeV2LiabilityAllocationJSON contains the JSON metadata for the struct
// [DisputeV2LiabilityAllocation]
type disputeV2LiabilityAllocationJSON struct {
	DeniedAmount     apijson.Field
	OriginalAmount   apijson.Field
	RecoveredAmount  apijson.Field
	RemainingAmount  apijson.Field
	WrittenOffAmount apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DisputeV2LiabilityAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2LiabilityAllocationJSON) RawJSON() string {
	return r.raw
}

// Card network handling the dispute.
type DisputeV2Network string

const (
	DisputeV2NetworkVisa       DisputeV2Network = "VISA"
	DisputeV2NetworkMastercard DisputeV2Network = "MASTERCARD"
)

func (r DisputeV2Network) IsKnown() bool {
	switch r {
	case DisputeV2NetworkVisa, DisputeV2NetworkMastercard:
		return true
	}
	return false
}

// Current status of the dispute.
type DisputeV2Status string

const (
	DisputeV2StatusOpen   DisputeV2Status = "OPEN"
	DisputeV2StatusClosed DisputeV2Status = "CLOSED"
)

func (r DisputeV2Status) IsKnown() bool {
	switch r {
	case DisputeV2StatusOpen, DisputeV2StatusClosed:
		return true
	}
	return false
}

// Contains identifiers for the transaction and specific event within being
// disputed; null if no transaction can be identified
type DisputeV2TransactionSeries struct {
	// Token of the specific event in the original transaction being disputed, in UUID
	// format; null if no event can be identified
	RelatedTransactionEventToken string `json:"related_transaction_event_token,required,nullable" format:"uuid"`
	// Token of the original transaction being disputed, in UUID format
	RelatedTransactionToken string `json:"related_transaction_token,required" format:"uuid"`
	// The type of transaction series associating the dispute and the original
	// transaction. Always set to DISPUTE
	Type DisputeV2TransactionSeriesType `json:"type,required"`
	JSON disputeV2TransactionSeriesJSON `json:"-"`
}

// disputeV2TransactionSeriesJSON contains the JSON metadata for the struct
// [DisputeV2TransactionSeries]
type disputeV2TransactionSeriesJSON struct {
	RelatedTransactionEventToken apijson.Field
	RelatedTransactionToken      apijson.Field
	Type                         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DisputeV2TransactionSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2TransactionSeriesJSON) RawJSON() string {
	return r.raw
}

// The type of transaction series associating the dispute and the original
// transaction. Always set to DISPUTE
type DisputeV2TransactionSeriesType string

const (
	DisputeV2TransactionSeriesTypeDispute DisputeV2TransactionSeriesType = "DISPUTE"
)

func (r DisputeV2TransactionSeriesType) IsKnown() bool {
	switch r {
	case DisputeV2TransactionSeriesTypeDispute:
		return true
	}
	return false
}

type DisputesV2ListParams struct {
	// Filter by account token.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// RFC 3339 timestamp for filtering by created date, inclusive.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Filter by card token.
	CardToken param.Field[string] `query:"card_token" format:"uuid"`
	// Filter by the token of the transaction being disputed. Corresponds with
	// transaction_series.related_transaction_token in the Dispute.
	DisputedTransactionToken param.Field[string] `query:"disputed_transaction_token" format:"uuid"`
	// RFC 3339 timestamp for filtering by created date, inclusive.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Number of items to return.
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [DisputesV2ListParams]'s query parameters as `url.Values`.
func (r DisputesV2ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
