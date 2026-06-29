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
	"github.com/tidwall/gjson"
)

// TransactionMonitoringCaseService contains methods and other services that help
// with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionMonitoringCaseService] method instead.
type TransactionMonitoringCaseService struct {
	Options  []option.RequestOption
	Comments *TransactionMonitoringCaseCommentService
	Files    *TransactionMonitoringCaseFileService
}

// NewTransactionMonitoringCaseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransactionMonitoringCaseService(opts ...option.RequestOption) (r *TransactionMonitoringCaseService) {
	r = &TransactionMonitoringCaseService{}
	r.Options = opts
	r.Comments = NewTransactionMonitoringCaseCommentService(opts...)
	r.Files = NewTransactionMonitoringCaseFileService(opts...)
	return
}

// Retrieves a single transaction monitoring case.
func (r *TransactionMonitoringCaseService) Get(ctx context.Context, caseToken string, opts ...option.RequestOption) (res *MonitoringCase, err error) {
	opts = slices.Concat(r.Options, opts)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s", caseToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a transaction monitoring case.
func (r *TransactionMonitoringCaseService) Update(ctx context.Context, caseToken string, body TransactionMonitoringCaseUpdateParams, opts ...option.RequestOption) (res *MonitoringCase, err error) {
	opts = slices.Concat(r.Options, opts)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s", caseToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists transaction monitoring cases, optionally filtered.
func (r *TransactionMonitoringCaseService) List(ctx context.Context, query TransactionMonitoringCaseListParams, opts ...option.RequestOption) (res *pagination.CursorPage[MonitoringCase], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/transaction_monitoring/cases"
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

// Lists transaction monitoring cases, optionally filtered.
func (r *TransactionMonitoringCaseService) ListAutoPaging(ctx context.Context, query TransactionMonitoringCaseListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[MonitoringCase] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Lists the activity feed for a case.
func (r *TransactionMonitoringCaseService) ListActivity(ctx context.Context, caseToken string, query TransactionMonitoringCaseListActivityParams, opts ...option.RequestOption) (res *pagination.CursorPage[CaseActivityEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/activity", caseToken)
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

// Lists the activity feed for a case.
func (r *TransactionMonitoringCaseService) ListActivityAutoPaging(ctx context.Context, caseToken string, query TransactionMonitoringCaseListActivityParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CaseActivityEntry] {
	return pagination.NewCursorPageAutoPager(r.ListActivity(ctx, caseToken, query, opts...))
}

// Lists the transactions associated with a case.
func (r *TransactionMonitoringCaseService) ListTransactions(ctx context.Context, caseToken string, query TransactionMonitoringCaseListTransactionsParams, opts ...option.RequestOption) (res *pagination.CursorPage[CaseTransaction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/transactions", caseToken)
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

// Lists the transactions associated with a case.
func (r *TransactionMonitoringCaseService) ListTransactionsAutoPaging(ctx context.Context, caseToken string, query TransactionMonitoringCaseListTransactionsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CaseTransaction] {
	return pagination.NewCursorPageAutoPager(r.ListTransactions(ctx, caseToken, query, opts...))
}

// Lists the cards involved in a case, with per-card transaction counts.
func (r *TransactionMonitoringCaseService) GetCards(ctx context.Context, caseToken string, opts ...option.RequestOption) (res *[]CaseCard, err error) {
	opts = slices.Concat(r.Options, opts)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/cards", caseToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A single entry in a case's activity feed
type CaseActivityEntry struct {
	// Globally unique identifier for the activity entry
	Token string `json:"token" api:"required" format:"uuid"`
	// Identifier of the actor that produced the activity entry
	ActorToken string `json:"actor_token" api:"required,nullable"`
	// Date and time at which the activity entry was created
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The case field that changed, or the action that was taken, in an activity entry:
	//
	// - `STATUS` - The case status changed
	// - `TITLE` - The case title changed
	// - `ASSIGNED_TO` - The case assignee changed
	// - `RESOLUTION_OUTCOME` - The resolution outcome was set or changed
	// - `RESOLUTION_NOTES` - The resolution notes were set or changed
	// - `TAGS` - The case tags changed
	// - `PRIORITY` - The case priority changed
	// - `COMMENT` - A comment was added or edited
	// - `FILE` - A file was attached to the case
	EntryType CaseActivityType `json:"entry_type" api:"required"`
	// New value of the changed field, when applicable
	NewValue string `json:"new_value" api:"required,nullable"`
	// Previous value of the changed field, when applicable
	PreviousValue string                `json:"previous_value" api:"required,nullable"`
	JSON          caseActivityEntryJSON `json:"-"`
}

// caseActivityEntryJSON contains the JSON metadata for the struct
// [CaseActivityEntry]
type caseActivityEntryJSON struct {
	Token         apijson.Field
	ActorToken    apijson.Field
	Created       apijson.Field
	EntryType     apijson.Field
	NewValue      apijson.Field
	PreviousValue apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CaseActivityEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseActivityEntryJSON) RawJSON() string {
	return r.raw
}

// The case field that changed, or the action that was taken, in an activity entry:
//
// - `STATUS` - The case status changed
// - `TITLE` - The case title changed
// - `ASSIGNED_TO` - The case assignee changed
// - `RESOLUTION_OUTCOME` - The resolution outcome was set or changed
// - `RESOLUTION_NOTES` - The resolution notes were set or changed
// - `TAGS` - The case tags changed
// - `PRIORITY` - The case priority changed
// - `COMMENT` - A comment was added or edited
// - `FILE` - A file was attached to the case
type CaseActivityType string

const (
	CaseActivityTypeStatus            CaseActivityType = "STATUS"
	CaseActivityTypeTitle             CaseActivityType = "TITLE"
	CaseActivityTypeAssignedTo        CaseActivityType = "ASSIGNED_TO"
	CaseActivityTypeResolutionOutcome CaseActivityType = "RESOLUTION_OUTCOME"
	CaseActivityTypeResolutionNotes   CaseActivityType = "RESOLUTION_NOTES"
	CaseActivityTypeTags              CaseActivityType = "TAGS"
	CaseActivityTypePriority          CaseActivityType = "PRIORITY"
	CaseActivityTypeComment           CaseActivityType = "COMMENT"
	CaseActivityTypeFile              CaseActivityType = "FILE"
)

func (r CaseActivityType) IsKnown() bool {
	switch r {
	case CaseActivityTypeStatus, CaseActivityTypeTitle, CaseActivityTypeAssignedTo, CaseActivityTypeResolutionOutcome, CaseActivityTypeResolutionNotes, CaseActivityTypeTags, CaseActivityTypePriority, CaseActivityTypeComment, CaseActivityTypeFile:
		return true
	}
	return false
}

// Summary of a card's involvement in a case, aggregated across the case's
// transactions
type CaseCard struct {
	// Token of the account the card belongs to
	AccountToken string `json:"account_token" api:"required" format:"uuid"`
	// Token of the card
	CardToken string `json:"card_token" api:"required" format:"uuid"`
	// Number of the card's transactions associated with the case
	TransactionCount int64        `json:"transaction_count" api:"required"`
	JSON             caseCardJSON `json:"-"`
}

// caseCardJSON contains the JSON metadata for the struct [CaseCard]
type caseCardJSON struct {
	AccountToken     apijson.Field
	CardToken        apijson.Field
	TransactionCount apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaseCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseCardJSON) RawJSON() string {
	return r.raw
}

// The entity a case is associated with
type CaseEntity struct {
	// Globally unique identifier for the associated entity
	EntityToken string `json:"entity_token" api:"required" format:"uuid"`
	// The type of entity a case is associated with:
	//
	// - `CARD` - The case is associated with a card
	// - `ACCOUNT` - The case is associated with an account
	EntityType CaseEntityEntityType `json:"entity_type" api:"required"`
	JSON       caseEntityJSON       `json:"-"`
}

// caseEntityJSON contains the JSON metadata for the struct [CaseEntity]
type caseEntityJSON struct {
	EntityToken apijson.Field
	EntityType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaseEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseEntityJSON) RawJSON() string {
	return r.raw
}

// The type of entity a case is associated with:
//
// - `CARD` - The case is associated with a card
// - `ACCOUNT` - The case is associated with an account
type CaseEntityEntityType string

const (
	CaseEntityEntityTypeCard    CaseEntityEntityType = "CARD"
	CaseEntityEntityTypeAccount CaseEntityEntityType = "ACCOUNT"
)

func (r CaseEntityEntityType) IsKnown() bool {
	switch r {
	case CaseEntityEntityTypeCard, CaseEntityEntityTypeAccount:
		return true
	}
	return false
}

// Priority level of a case, controlling queue ordering and SLA urgency
type CasePriority string

const (
	CasePriorityLow      CasePriority = "LOW"
	CasePriorityMedium   CasePriority = "MEDIUM"
	CasePriorityHigh     CasePriority = "HIGH"
	CasePriorityCritical CasePriority = "CRITICAL"
)

func (r CasePriority) IsKnown() bool {
	switch r {
	case CasePriorityLow, CasePriorityMedium, CasePriorityHigh, CasePriorityCritical:
		return true
	}
	return false
}

// Sort order for listing cases. Defaults to `CREATED_DESC` (newest first):
//
// - `CREATED_ASC` - Oldest first
// - `CREATED_DESC` - Newest first
// - `PRIORITY_DESC` - Highest priority first
// - `PRIORITY_ASC` - Lowest priority first
// - `STATUS_DESC` - Furthest workflow stage first
// - `STATUS_ASC` - Earliest workflow stage first
type CaseSortOrder string

const (
	CaseSortOrderCreatedAsc   CaseSortOrder = "CREATED_ASC"
	CaseSortOrderCreatedDesc  CaseSortOrder = "CREATED_DESC"
	CaseSortOrderPriorityDesc CaseSortOrder = "PRIORITY_DESC"
	CaseSortOrderPriorityAsc  CaseSortOrder = "PRIORITY_ASC"
	CaseSortOrderStatusDesc   CaseSortOrder = "STATUS_DESC"
	CaseSortOrderStatusAsc    CaseSortOrder = "STATUS_ASC"
)

func (r CaseSortOrder) IsKnown() bool {
	switch r {
	case CaseSortOrderCreatedAsc, CaseSortOrderCreatedDesc, CaseSortOrderPriorityDesc, CaseSortOrderPriorityAsc, CaseSortOrderStatusDesc, CaseSortOrderStatusAsc:
		return true
	}
	return false
}

// Status of a case as it progresses through the review workflow:
//
//   - `OPEN` - The case has been created and is still collecting matching
//     transactions
//   - `ASSIGNED` - An analyst has been assigned and transaction collection has
//     stopped
//   - `IN_REVIEW` - The case is actively being investigated
//   - `ESCALATED` - The case has been reviewed and requires additional oversight
//   - `RESOLVED` - A determination has been made and a resolution recorded
//   - `CLOSED` - The case is finalized
type CaseStatus string

const (
	CaseStatusOpen      CaseStatus = "OPEN"
	CaseStatusAssigned  CaseStatus = "ASSIGNED"
	CaseStatusInReview  CaseStatus = "IN_REVIEW"
	CaseStatusEscalated CaseStatus = "ESCALATED"
	CaseStatusResolved  CaseStatus = "RESOLVED"
	CaseStatusClosed    CaseStatus = "CLOSED"
)

func (r CaseStatus) IsKnown() bool {
	switch r {
	case CaseStatusOpen, CaseStatusAssigned, CaseStatusInReview, CaseStatusEscalated, CaseStatusResolved, CaseStatusClosed:
		return true
	}
	return false
}

// A single transaction associated with a case. The `category` field identifies
// whether this is a card transaction or a payment transaction.
type CaseTransaction struct {
	// Globally unique identifier for the card transaction
	Token string `json:"token" api:"required" format:"uuid"`
	// Date and time at which the transaction was added to the case
	AddedAt  time.Time               `json:"added_at" api:"required" format:"date-time"`
	Category CaseTransactionCategory `json:"category" api:"required"`
	// Date and time at which the transaction was created
	TransactionCreatedAt time.Time `json:"transaction_created_at" api:"required" format:"date-time"`
	// Token of the account the transaction belongs to
	AccountToken string `json:"account_token" format:"uuid"`
	// Token of the card the transaction was made on
	CardToken string `json:"card_token" format:"uuid"`
	// Token of the financial account the payment belongs to
	FinancialAccountToken string              `json:"financial_account_token" format:"uuid"`
	JSON                  caseTransactionJSON `json:"-"`
	union                 CaseTransactionUnion
}

// caseTransactionJSON contains the JSON metadata for the struct [CaseTransaction]
type caseTransactionJSON struct {
	Token                 apijson.Field
	AddedAt               apijson.Field
	Category              apijson.Field
	TransactionCreatedAt  apijson.Field
	AccountToken          apijson.Field
	CardToken             apijson.Field
	FinancialAccountToken apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r caseTransactionJSON) RawJSON() string {
	return r.raw
}

func (r *CaseTransaction) UnmarshalJSON(data []byte) (err error) {
	*r = CaseTransaction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CaseTransactionUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [CaseTransactionCardCaseTransaction],
// [CaseTransactionPaymentCaseTransaction].
func (r CaseTransaction) AsUnion() CaseTransactionUnion {
	return r.union
}

// A single transaction associated with a case. The `category` field identifies
// whether this is a card transaction or a payment transaction.
//
// Union satisfied by [CaseTransactionCardCaseTransaction] or
// [CaseTransactionPaymentCaseTransaction].
type CaseTransactionUnion interface {
	implementsCaseTransaction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CaseTransactionUnion)(nil)).Elem(),
		"category",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CaseTransactionCardCaseTransaction{}),
			DiscriminatorValue: "CARD",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CaseTransactionPaymentCaseTransaction{}),
			DiscriminatorValue: "PAYMENT",
		},
	)
}

// A card transaction associated with a case
type CaseTransactionCardCaseTransaction struct {
	// Globally unique identifier for the card transaction
	Token string `json:"token" api:"required" format:"uuid"`
	// Token of the account the transaction belongs to
	AccountToken string `json:"account_token" api:"required" format:"uuid"`
	// Date and time at which the transaction was added to the case
	AddedAt time.Time `json:"added_at" api:"required" format:"date-time"`
	// Token of the card the transaction was made on
	CardToken string                                     `json:"card_token" api:"required" format:"uuid"`
	Category  CaseTransactionCardCaseTransactionCategory `json:"category" api:"required"`
	// Date and time at which the transaction was created
	TransactionCreatedAt time.Time                              `json:"transaction_created_at" api:"required" format:"date-time"`
	JSON                 caseTransactionCardCaseTransactionJSON `json:"-"`
}

// caseTransactionCardCaseTransactionJSON contains the JSON metadata for the struct
// [CaseTransactionCardCaseTransaction]
type caseTransactionCardCaseTransactionJSON struct {
	Token                apijson.Field
	AccountToken         apijson.Field
	AddedAt              apijson.Field
	CardToken            apijson.Field
	Category             apijson.Field
	TransactionCreatedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CaseTransactionCardCaseTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseTransactionCardCaseTransactionJSON) RawJSON() string {
	return r.raw
}

func (r CaseTransactionCardCaseTransaction) implementsCaseTransaction() {}

type CaseTransactionCardCaseTransactionCategory string

const (
	CaseTransactionCardCaseTransactionCategoryCard CaseTransactionCardCaseTransactionCategory = "CARD"
)

func (r CaseTransactionCardCaseTransactionCategory) IsKnown() bool {
	switch r {
	case CaseTransactionCardCaseTransactionCategoryCard:
		return true
	}
	return false
}

// A payment (ACH) transaction associated with a case
type CaseTransactionPaymentCaseTransaction struct {
	// Globally unique identifier for the payment transaction
	Token string `json:"token" api:"required" format:"uuid"`
	// Date and time at which the transaction was added to the case
	AddedAt  time.Time                                     `json:"added_at" api:"required" format:"date-time"`
	Category CaseTransactionPaymentCaseTransactionCategory `json:"category" api:"required"`
	// Token of the financial account the payment belongs to
	FinancialAccountToken string `json:"financial_account_token" api:"required" format:"uuid"`
	// Date and time at which the transaction was created
	TransactionCreatedAt time.Time `json:"transaction_created_at" api:"required" format:"date-time"`
	// Token of the account the payment belongs to, if applicable
	AccountToken string                                    `json:"account_token" format:"uuid"`
	JSON         caseTransactionPaymentCaseTransactionJSON `json:"-"`
}

// caseTransactionPaymentCaseTransactionJSON contains the JSON metadata for the
// struct [CaseTransactionPaymentCaseTransaction]
type caseTransactionPaymentCaseTransactionJSON struct {
	Token                 apijson.Field
	AddedAt               apijson.Field
	Category              apijson.Field
	FinancialAccountToken apijson.Field
	TransactionCreatedAt  apijson.Field
	AccountToken          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CaseTransactionPaymentCaseTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseTransactionPaymentCaseTransactionJSON) RawJSON() string {
	return r.raw
}

func (r CaseTransactionPaymentCaseTransaction) implementsCaseTransaction() {}

type CaseTransactionPaymentCaseTransactionCategory string

const (
	CaseTransactionPaymentCaseTransactionCategoryPayment CaseTransactionPaymentCaseTransactionCategory = "PAYMENT"
)

func (r CaseTransactionPaymentCaseTransactionCategory) IsKnown() bool {
	switch r {
	case CaseTransactionPaymentCaseTransactionCategoryPayment:
		return true
	}
	return false
}

type CaseTransactionCategory string

const (
	CaseTransactionCategoryCard    CaseTransactionCategory = "CARD"
	CaseTransactionCategoryPayment CaseTransactionCategory = "PAYMENT"
)

func (r CaseTransactionCategory) IsKnown() bool {
	switch r {
	case CaseTransactionCategoryCard, CaseTransactionCategoryPayment:
		return true
	}
	return false
}

// The type of entity associated with an account holder
type EntityType string

const (
	EntityTypeBeneficialOwnerIndividual EntityType = "BENEFICIAL_OWNER_INDIVIDUAL"
	EntityTypeControlPerson             EntityType = "CONTROL_PERSON"
)

func (r EntityType) IsKnown() bool {
	switch r {
	case EntityTypeBeneficialOwnerIndividual, EntityTypeControlPerson:
		return true
	}
	return false
}

// A transaction monitoring case
type MonitoringCase struct {
	// Globally unique identifier for the case
	Token string `json:"token" api:"required" format:"uuid"`
	// Identifier of the user the case is currently assigned to
	Assignee string `json:"assignee" api:"required,nullable"`
	// Date and time at which transaction collection stopped for the case
	CollectionStopped time.Time `json:"collection_stopped" api:"required,nullable" format:"date-time"`
	// Date and time at which the case was created
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// The entity a case is associated with
	Entity CaseEntity `json:"entity" api:"required,nullable"`
	// Whether the case still has transaction scopes pending resolution
	PendingTransactions bool `json:"pending_transactions" api:"required"`
	// Priority level of a case, controlling queue ordering and SLA urgency
	Priority CasePriority `json:"priority" api:"required"`
	// Token of the queue the case belongs to
	QueueToken string `json:"queue_token" api:"required" format:"uuid"`
	// Outcome recorded when a case is resolved:
	//
	//   - `CONFIRMED_FRAUD` - The reviewed activity was confirmed to be fraudulent
	//   - `SUSPICIOUS_ACTIVITY` - The activity is suspicious but not confirmed fraud
	//   - `FALSE_POSITIVE` - The activity was legitimate and the alert was a false
	//     positive
	//   - `NO_ACTION_REQUIRED` - No further action is required
	//   - `ESCALATED_EXTERNAL` - The case was escalated to an external party
	Resolution ResolutionOutcome `json:"resolution" api:"required,nullable"`
	// Free-form notes describing the resolution
	ResolutionNotes string `json:"resolution_notes" api:"required,nullable"`
	// Date and time at which the case was resolved
	Resolved time.Time `json:"resolved" api:"required,nullable" format:"date-time"`
	// Token of the transaction monitoring rule that triggered the case
	RuleToken string `json:"rule_token" api:"required,nullable" format:"uuid"`
	// Deadline by which the case is expected to be resolved
	SlaDeadline time.Time `json:"sla_deadline" api:"required,nullable" format:"date-time"`
	// Status of a case as it progresses through the review workflow:
	//
	//   - `OPEN` - The case has been created and is still collecting matching
	//     transactions
	//   - `ASSIGNED` - An analyst has been assigned and transaction collection has
	//     stopped
	//   - `IN_REVIEW` - The case is actively being investigated
	//   - `ESCALATED` - The case has been reviewed and requires additional oversight
	//   - `RESOLVED` - A determination has been made and a resolution recorded
	//   - `CLOSED` - The case is finalized
	Status CaseStatus `json:"status" api:"required"`
	// Arbitrary key-value metadata associated with the case
	Tags map[string]string `json:"tags" api:"required"`
	// Short, human-readable summary of the case
	Title string `json:"title" api:"required,nullable"`
	// Date and time at which the case was last updated
	Updated time.Time          `json:"updated" api:"required" format:"date-time"`
	JSON    monitoringCaseJSON `json:"-"`
}

// monitoringCaseJSON contains the JSON metadata for the struct [MonitoringCase]
type monitoringCaseJSON struct {
	Token               apijson.Field
	Assignee            apijson.Field
	CollectionStopped   apijson.Field
	Created             apijson.Field
	Entity              apijson.Field
	PendingTransactions apijson.Field
	Priority            apijson.Field
	QueueToken          apijson.Field
	Resolution          apijson.Field
	ResolutionNotes     apijson.Field
	Resolved            apijson.Field
	RuleToken           apijson.Field
	SlaDeadline         apijson.Field
	Status              apijson.Field
	Tags                apijson.Field
	Title               apijson.Field
	Updated             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MonitoringCase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitoringCaseJSON) RawJSON() string {
	return r.raw
}

// Outcome recorded when a case is resolved:
//
//   - `CONFIRMED_FRAUD` - The reviewed activity was confirmed to be fraudulent
//   - `SUSPICIOUS_ACTIVITY` - The activity is suspicious but not confirmed fraud
//   - `FALSE_POSITIVE` - The activity was legitimate and the alert was a false
//     positive
//   - `NO_ACTION_REQUIRED` - No further action is required
//   - `ESCALATED_EXTERNAL` - The case was escalated to an external party
type ResolutionOutcome string

const (
	ResolutionOutcomeConfirmedFraud     ResolutionOutcome = "CONFIRMED_FRAUD"
	ResolutionOutcomeSuspiciousActivity ResolutionOutcome = "SUSPICIOUS_ACTIVITY"
	ResolutionOutcomeFalsePositive      ResolutionOutcome = "FALSE_POSITIVE"
	ResolutionOutcomeNoActionRequired   ResolutionOutcome = "NO_ACTION_REQUIRED"
	ResolutionOutcomeEscalatedExternal  ResolutionOutcome = "ESCALATED_EXTERNAL"
)

func (r ResolutionOutcome) IsKnown() bool {
	switch r {
	case ResolutionOutcomeConfirmedFraud, ResolutionOutcomeSuspiciousActivity, ResolutionOutcomeFalsePositive, ResolutionOutcomeNoActionRequired, ResolutionOutcomeEscalatedExternal:
		return true
	}
	return false
}

type TransactionMonitoringCaseUpdateParams struct {
	// Optional client-provided identifier for the actor performing this action,
	// recorded on the resulting activity entry. This value is supplied by the client
	// (for example, your own internal user ID) and is not authenticated by Lithic
	ActorToken param.Field[string] `json:"actor_token"`
	// New assignee for the case, or `null` to unassign
	Assignee param.Field[string] `json:"assignee"`
	// Priority level of a case, controlling queue ordering and SLA urgency
	Priority param.Field[CasePriority] `json:"priority"`
	// Outcome recorded when a case is resolved:
	//
	//   - `CONFIRMED_FRAUD` - The reviewed activity was confirmed to be fraudulent
	//   - `SUSPICIOUS_ACTIVITY` - The activity is suspicious but not confirmed fraud
	//   - `FALSE_POSITIVE` - The activity was legitimate and the alert was a false
	//     positive
	//   - `NO_ACTION_REQUIRED` - No further action is required
	//   - `ESCALATED_EXTERNAL` - The case was escalated to an external party
	Resolution param.Field[ResolutionOutcome] `json:"resolution"`
	// Notes describing the resolution
	ResolutionNotes param.Field[string] `json:"resolution_notes"`
	// New SLA deadline for the case, or `null` to clear it
	SlaDeadline param.Field[time.Time] `json:"sla_deadline" format:"date-time"`
	// Status of a case as it progresses through the review workflow:
	//
	//   - `OPEN` - The case has been created and is still collecting matching
	//     transactions
	//   - `ASSIGNED` - An analyst has been assigned and transaction collection has
	//     stopped
	//   - `IN_REVIEW` - The case is actively being investigated
	//   - `ESCALATED` - The case has been reviewed and requires additional oversight
	//   - `RESOLVED` - A determination has been made and a resolution recorded
	//   - `CLOSED` - The case is finalized
	Status param.Field[CaseStatus] `json:"status"`
	// Arbitrary key-value metadata to set on the case
	Tags param.Field[map[string]string] `json:"tags"`
	// New title for the case, or `null` to clear it
	Title param.Field[string] `json:"title"`
}

func (r TransactionMonitoringCaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionMonitoringCaseListParams struct {
	// Only return cases that include transactions on the provided account.
	AccountToken param.Field[string] `query:"account_token" format:"uuid"`
	// Only return cases assigned to the provided value. Pass an empty string to return
	// only unassigned cases.
	Assignee param.Field[string] `query:"assignee"`
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Only return cases that include transactions on the provided card.
	CardToken param.Field[string] `query:"card_token" format:"uuid"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Only return cases associated with the provided entity.
	EntityToken param.Field[string] `query:"entity_token" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// Only return cases belonging to the provided queue.
	QueueToken param.Field[string] `query:"queue_token" format:"uuid"`
	// Only return cases triggered by the provided transaction monitoring rule.
	RuleToken param.Field[string] `query:"rule_token" format:"uuid"`
	// Sort order for the returned cases.
	SortBy param.Field[CaseSortOrder] `query:"sort_by"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
	// Only return cases with the provided status.
	Status param.Field[CaseStatus] `query:"status"`
	// Only return cases that include the provided transaction.
	TransactionToken param.Field[string] `query:"transaction_token" format:"uuid"`
}

// URLQuery serializes [TransactionMonitoringCaseListParams]'s query parameters as
// `url.Values`.
func (r TransactionMonitoringCaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionMonitoringCaseListActivityParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [TransactionMonitoringCaseListActivityParams]'s query
// parameters as `url.Values`.
func (r TransactionMonitoringCaseListActivityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionMonitoringCaseListTransactionsParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [TransactionMonitoringCaseListTransactionsParams]'s query
// parameters as `url.Values`.
func (r TransactionMonitoringCaseListTransactionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
