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
		return nil, err
	}
	path := fmt.Sprintf("v2/disputes/%s", disputeToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
	Token string `json:"token" api:"required" format:"uuid"`
	// Token for the account associated with the dispute, in UUID format.
	AccountToken string `json:"account_token" api:"required" format:"uuid"`
	// Token for the card used in the dispute, in UUID format.
	CardToken string `json:"card_token" api:"required" format:"uuid"`
	// Identifier assigned by the network for this dispute.
	CaseID string `json:"case_id" api:"required,nullable"`
	// Token for the claim this dispute was filed under, in UUID format. Null for
	// disputes not initiated through the Dispute Intake API.
	ClaimToken string `json:"claim_token" api:"required,nullable" format:"uuid"`
	// When the dispute was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Three-letter ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// Dispute resolution outcome
	Disposition DisputeV2Disposition `json:"disposition" api:"required,nullable"`
	// Chronological list of events that have occurred in the dispute lifecycle
	Events []DisputeV2Event `json:"events" api:"required"`
	// Current breakdown of how liability is allocated for the disputed amount
	LiabilityAllocation DisputeV2LiabilityAllocation `json:"liability_allocation" api:"required"`
	Merchant            shared.Merchant              `json:"merchant" api:"required"`
	// Card network handling the dispute.
	Network DisputeV2Network `json:"network" api:"required"`
	// Current status of the dispute.
	Status DisputeV2Status `json:"status" api:"required,nullable"`
	// Contains identifiers for the transaction and specific event within being
	// disputed; null if no transaction can be identified
	TransactionSeries DisputeV2TransactionSeries `json:"transaction_series" api:"required,nullable"`
	// When the dispute was last updated.
	Updated time.Time     `json:"updated" api:"required" format:"date-time"`
	JSON    disputeV2JSON `json:"-"`
}

// disputeV2JSON contains the JSON metadata for the struct [DisputeV2]
type disputeV2JSON struct {
	Token               apijson.Field
	AccountToken        apijson.Field
	CardToken           apijson.Field
	CaseID              apijson.Field
	ClaimToken          apijson.Field
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

// Event that occurred in the dispute lifecycle. The `type` field identifies the
// event variant and determines the shape of `data`
type DisputeV2Event struct {
	// Unique identifier for the event, in UUID format
	Token string `json:"token" api:"required" format:"uuid"`
	// When the event occurred
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// This field can have the runtime type of [DisputeV2EventsWorkflowEventData],
	// [DisputeV2EventsFinancialEventData],
	// [DisputeV2EventsCardholderLiabilityEventData].
	Data interface{} `json:"data" api:"required"`
	// Type of event. Always `WORKFLOW`
	Type  DisputeV2EventsType `json:"type" api:"required"`
	JSON  disputeV2EventJSON  `json:"-"`
	union DisputeV2EventsUnion
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

func (r disputeV2EventJSON) RawJSON() string {
	return r.raw
}

func (r *DisputeV2Event) UnmarshalJSON(data []byte) (err error) {
	*r = DisputeV2Event{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DisputeV2EventsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DisputeV2EventsWorkflowEvent],
// [DisputeV2EventsFinancialEvent], [DisputeV2EventsCardholderLiabilityEvent].
func (r DisputeV2Event) AsUnion() DisputeV2EventsUnion {
	return r.union
}

// Event that occurred in the dispute lifecycle. The `type` field identifies the
// event variant and determines the shape of `data`
//
// Union satisfied by [DisputeV2EventsWorkflowEvent],
// [DisputeV2EventsFinancialEvent] or [DisputeV2EventsCardholderLiabilityEvent].
type DisputeV2EventsUnion interface {
	implementsDisputeV2Event()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DisputeV2EventsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DisputeV2EventsWorkflowEvent{}),
			DiscriminatorValue: "WORKFLOW",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DisputeV2EventsFinancialEvent{}),
			DiscriminatorValue: "FINANCIAL",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DisputeV2EventsCardholderLiabilityEvent{}),
			DiscriminatorValue: "CARDHOLDER_LIABILITY",
		},
	)
}

// Event tracking the dispute's case management workflow
type DisputeV2EventsWorkflowEvent struct {
	// Unique identifier for the event, in UUID format
	Token string `json:"token" api:"required" format:"uuid"`
	// When the event occurred
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Details specific to workflow events
	Data DisputeV2EventsWorkflowEventData `json:"data" api:"required"`
	// Type of event. Always `WORKFLOW`
	Type DisputeV2EventsWorkflowEventType `json:"type" api:"required"`
	JSON disputeV2EventsWorkflowEventJSON `json:"-"`
}

// disputeV2EventsWorkflowEventJSON contains the JSON metadata for the struct
// [DisputeV2EventsWorkflowEvent]
type disputeV2EventsWorkflowEventJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsWorkflowEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsWorkflowEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeV2EventsWorkflowEvent) implementsDisputeV2Event() {}

// Details specific to workflow events
type DisputeV2EventsWorkflowEventData struct {
	// Action taken in this stage
	Action DisputeV2EventsWorkflowEventDataAction `json:"action" api:"required"`
	// Amount in minor units
	Amount int64 `json:"amount" api:"required,nullable"`
	// Dispute resolution outcome
	Disposition DisputeV2EventsWorkflowEventDataDisposition `json:"disposition" api:"required,nullable"`
	// Reason for the action
	Reason string `json:"reason" api:"required,nullable"`
	// Current stage of the dispute workflow
	Stage DisputeV2EventsWorkflowEventDataStage `json:"stage" api:"required"`
	JSON  disputeV2EventsWorkflowEventDataJSON  `json:"-"`
}

// disputeV2EventsWorkflowEventDataJSON contains the JSON metadata for the struct
// [DisputeV2EventsWorkflowEventData]
type disputeV2EventsWorkflowEventDataJSON struct {
	Action      apijson.Field
	Amount      apijson.Field
	Disposition apijson.Field
	Reason      apijson.Field
	Stage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsWorkflowEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsWorkflowEventDataJSON) RawJSON() string {
	return r.raw
}

// Action taken in this stage
type DisputeV2EventsWorkflowEventDataAction string

const (
	DisputeV2EventsWorkflowEventDataActionOpened   DisputeV2EventsWorkflowEventDataAction = "OPENED"
	DisputeV2EventsWorkflowEventDataActionClosed   DisputeV2EventsWorkflowEventDataAction = "CLOSED"
	DisputeV2EventsWorkflowEventDataActionReopened DisputeV2EventsWorkflowEventDataAction = "REOPENED"
)

func (r DisputeV2EventsWorkflowEventDataAction) IsKnown() bool {
	switch r {
	case DisputeV2EventsWorkflowEventDataActionOpened, DisputeV2EventsWorkflowEventDataActionClosed, DisputeV2EventsWorkflowEventDataActionReopened:
		return true
	}
	return false
}

// Dispute resolution outcome
type DisputeV2EventsWorkflowEventDataDisposition string

const (
	DisputeV2EventsWorkflowEventDataDispositionWon          DisputeV2EventsWorkflowEventDataDisposition = "WON"
	DisputeV2EventsWorkflowEventDataDispositionLost         DisputeV2EventsWorkflowEventDataDisposition = "LOST"
	DisputeV2EventsWorkflowEventDataDispositionPartiallyWon DisputeV2EventsWorkflowEventDataDisposition = "PARTIALLY_WON"
	DisputeV2EventsWorkflowEventDataDispositionWithdrawn    DisputeV2EventsWorkflowEventDataDisposition = "WITHDRAWN"
	DisputeV2EventsWorkflowEventDataDispositionDenied       DisputeV2EventsWorkflowEventDataDisposition = "DENIED"
)

func (r DisputeV2EventsWorkflowEventDataDisposition) IsKnown() bool {
	switch r {
	case DisputeV2EventsWorkflowEventDataDispositionWon, DisputeV2EventsWorkflowEventDataDispositionLost, DisputeV2EventsWorkflowEventDataDispositionPartiallyWon, DisputeV2EventsWorkflowEventDataDispositionWithdrawn, DisputeV2EventsWorkflowEventDataDispositionDenied:
		return true
	}
	return false
}

// Current stage of the dispute workflow
type DisputeV2EventsWorkflowEventDataStage string

const (
	DisputeV2EventsWorkflowEventDataStageClaim DisputeV2EventsWorkflowEventDataStage = "CLAIM"
)

func (r DisputeV2EventsWorkflowEventDataStage) IsKnown() bool {
	switch r {
	case DisputeV2EventsWorkflowEventDataStageClaim:
		return true
	}
	return false
}

// Type of event. Always `WORKFLOW`
type DisputeV2EventsWorkflowEventType string

const (
	DisputeV2EventsWorkflowEventTypeWorkflow DisputeV2EventsWorkflowEventType = "WORKFLOW"
)

func (r DisputeV2EventsWorkflowEventType) IsKnown() bool {
	switch r {
	case DisputeV2EventsWorkflowEventTypeWorkflow:
		return true
	}
	return false
}

// Event tracking a funds movement between issuer and acquirer
type DisputeV2EventsFinancialEvent struct {
	// Unique identifier for the event, in UUID format
	Token string `json:"token" api:"required" format:"uuid"`
	// When the event occurred
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Details specific to financial events
	Data DisputeV2EventsFinancialEventData `json:"data" api:"required"`
	// Type of event. Always `FINANCIAL`
	Type DisputeV2EventsFinancialEventType `json:"type" api:"required"`
	JSON disputeV2EventsFinancialEventJSON `json:"-"`
}

// disputeV2EventsFinancialEventJSON contains the JSON metadata for the struct
// [DisputeV2EventsFinancialEvent]
type disputeV2EventsFinancialEventJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsFinancialEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsFinancialEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeV2EventsFinancialEvent) implementsDisputeV2Event() {}

// Details specific to financial events
type DisputeV2EventsFinancialEventData struct {
	// Amount in minor units
	Amount int64 `json:"amount" api:"required"`
	// Direction of funds flow
	Polarity DisputeV2EventsFinancialEventDataPolarity `json:"polarity" api:"required"`
	// Stage at which the financial event occurred
	Stage DisputeV2EventsFinancialEventDataStage `json:"stage" api:"required"`
	JSON  disputeV2EventsFinancialEventDataJSON  `json:"-"`
}

// disputeV2EventsFinancialEventDataJSON contains the JSON metadata for the struct
// [DisputeV2EventsFinancialEventData]
type disputeV2EventsFinancialEventDataJSON struct {
	Amount      apijson.Field
	Polarity    apijson.Field
	Stage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsFinancialEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsFinancialEventDataJSON) RawJSON() string {
	return r.raw
}

// Direction of funds flow
type DisputeV2EventsFinancialEventDataPolarity string

const (
	DisputeV2EventsFinancialEventDataPolarityCredit DisputeV2EventsFinancialEventDataPolarity = "CREDIT"
	DisputeV2EventsFinancialEventDataPolarityDebit  DisputeV2EventsFinancialEventDataPolarity = "DEBIT"
)

func (r DisputeV2EventsFinancialEventDataPolarity) IsKnown() bool {
	switch r {
	case DisputeV2EventsFinancialEventDataPolarityCredit, DisputeV2EventsFinancialEventDataPolarityDebit:
		return true
	}
	return false
}

// Stage at which the financial event occurred
type DisputeV2EventsFinancialEventDataStage string

const (
	DisputeV2EventsFinancialEventDataStageChargeback     DisputeV2EventsFinancialEventDataStage = "CHARGEBACK"
	DisputeV2EventsFinancialEventDataStageRepresentment  DisputeV2EventsFinancialEventDataStage = "REPRESENTMENT"
	DisputeV2EventsFinancialEventDataStagePrearbitration DisputeV2EventsFinancialEventDataStage = "PREARBITRATION"
	DisputeV2EventsFinancialEventDataStageArbitration    DisputeV2EventsFinancialEventDataStage = "ARBITRATION"
	DisputeV2EventsFinancialEventDataStageCollaboration  DisputeV2EventsFinancialEventDataStage = "COLLABORATION"
)

func (r DisputeV2EventsFinancialEventDataStage) IsKnown() bool {
	switch r {
	case DisputeV2EventsFinancialEventDataStageChargeback, DisputeV2EventsFinancialEventDataStageRepresentment, DisputeV2EventsFinancialEventDataStagePrearbitration, DisputeV2EventsFinancialEventDataStageArbitration, DisputeV2EventsFinancialEventDataStageCollaboration:
		return true
	}
	return false
}

// Type of event. Always `FINANCIAL`
type DisputeV2EventsFinancialEventType string

const (
	DisputeV2EventsFinancialEventTypeFinancial DisputeV2EventsFinancialEventType = "FINANCIAL"
)

func (r DisputeV2EventsFinancialEventType) IsKnown() bool {
	switch r {
	case DisputeV2EventsFinancialEventTypeFinancial:
		return true
	}
	return false
}

// Event tracking a change in cardholder liability
type DisputeV2EventsCardholderLiabilityEvent struct {
	// Unique identifier for the event, in UUID format
	Token string `json:"token" api:"required" format:"uuid"`
	// When the event occurred
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Details specific to cardholder liability events
	Data DisputeV2EventsCardholderLiabilityEventData `json:"data" api:"required"`
	// Type of event. Always `CARDHOLDER_LIABILITY`
	Type DisputeV2EventsCardholderLiabilityEventType `json:"type" api:"required"`
	JSON disputeV2EventsCardholderLiabilityEventJSON `json:"-"`
}

// disputeV2EventsCardholderLiabilityEventJSON contains the JSON metadata for the
// struct [DisputeV2EventsCardholderLiabilityEvent]
type disputeV2EventsCardholderLiabilityEventJSON struct {
	Token       apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsCardholderLiabilityEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsCardholderLiabilityEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeV2EventsCardholderLiabilityEvent) implementsDisputeV2Event() {}

// Details specific to cardholder liability events
type DisputeV2EventsCardholderLiabilityEventData struct {
	// Action taken regarding cardholder liability
	Action DisputeV2EventsCardholderLiabilityEventDataAction `json:"action" api:"required"`
	// Amount in minor units
	Amount int64 `json:"amount" api:"required"`
	// Reason for the action
	Reason string                                          `json:"reason" api:"required,nullable"`
	JSON   disputeV2EventsCardholderLiabilityEventDataJSON `json:"-"`
}

// disputeV2EventsCardholderLiabilityEventDataJSON contains the JSON metadata for
// the struct [DisputeV2EventsCardholderLiabilityEventData]
type disputeV2EventsCardholderLiabilityEventDataJSON struct {
	Action      apijson.Field
	Amount      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeV2EventsCardholderLiabilityEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeV2EventsCardholderLiabilityEventDataJSON) RawJSON() string {
	return r.raw
}

// Action taken regarding cardholder liability
type DisputeV2EventsCardholderLiabilityEventDataAction string

const (
	DisputeV2EventsCardholderLiabilityEventDataActionProvisionalCreditGranted  DisputeV2EventsCardholderLiabilityEventDataAction = "PROVISIONAL_CREDIT_GRANTED"
	DisputeV2EventsCardholderLiabilityEventDataActionProvisionalCreditReversed DisputeV2EventsCardholderLiabilityEventDataAction = "PROVISIONAL_CREDIT_REVERSED"
	DisputeV2EventsCardholderLiabilityEventDataActionWrittenOff                DisputeV2EventsCardholderLiabilityEventDataAction = "WRITTEN_OFF"
	DisputeV2EventsCardholderLiabilityEventDataActionWriteOffReversed          DisputeV2EventsCardholderLiabilityEventDataAction = "WRITE_OFF_REVERSED"
)

func (r DisputeV2EventsCardholderLiabilityEventDataAction) IsKnown() bool {
	switch r {
	case DisputeV2EventsCardholderLiabilityEventDataActionProvisionalCreditGranted, DisputeV2EventsCardholderLiabilityEventDataActionProvisionalCreditReversed, DisputeV2EventsCardholderLiabilityEventDataActionWrittenOff, DisputeV2EventsCardholderLiabilityEventDataActionWriteOffReversed:
		return true
	}
	return false
}

// Type of event. Always `CARDHOLDER_LIABILITY`
type DisputeV2EventsCardholderLiabilityEventType string

const (
	DisputeV2EventsCardholderLiabilityEventTypeCardholderLiability DisputeV2EventsCardholderLiabilityEventType = "CARDHOLDER_LIABILITY"
)

func (r DisputeV2EventsCardholderLiabilityEventType) IsKnown() bool {
	switch r {
	case DisputeV2EventsCardholderLiabilityEventTypeCardholderLiability:
		return true
	}
	return false
}

// Type of event. Always `WORKFLOW`
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
	DeniedAmount int64 `json:"denied_amount" api:"required"`
	// The initial amount disputed
	OriginalAmount int64 `json:"original_amount" api:"required"`
	// The amount that has been recovered from the merchant through the dispute process
	RecoveredAmount int64 `json:"recovered_amount" api:"required"`
	// Any disputed amount that is still outstanding, i.e. has not been recovered,
	// written off, or denied
	RemainingAmount int64 `json:"remaining_amount" api:"required"`
	// The amount the issuer has chosen to write off
	WrittenOffAmount int64                            `json:"written_off_amount" api:"required"`
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
	RelatedTransactionEventToken string `json:"related_transaction_event_token" api:"required,nullable" format:"uuid"`
	// Token of the original transaction being disputed, in UUID format
	RelatedTransactionToken string `json:"related_transaction_token" api:"required" format:"uuid"`
	// The type of transaction series associating the dispute and the original
	// transaction. Always set to DISPUTE
	Type DisputeV2TransactionSeriesType `json:"type" api:"required"`
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
	// Filter by the token of the claim the dispute was filed under. Returns the
	// disputes created from that claim's disputed transaction events.
	ClaimToken param.Field[string] `query:"claim_token" format:"uuid"`
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
