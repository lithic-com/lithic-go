// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/apiquery"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/packages/pagination"
)

// TransactionMonitoringQueueService contains methods and other services that help
// with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionMonitoringQueueService] method instead.
type TransactionMonitoringQueueService struct {
	Options []option.RequestOption
}

// NewTransactionMonitoringQueueService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransactionMonitoringQueueService(opts ...option.RequestOption) (r *TransactionMonitoringQueueService) {
	r = &TransactionMonitoringQueueService{}
	r.Options = opts
	return
}

// Creates a new queue for grouping transaction monitoring cases.
func (r *TransactionMonitoringQueueService) New(ctx context.Context, body TransactionMonitoringQueueNewParams, opts ...option.RequestOption) (res *Queue, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/transaction_monitoring/queues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single transaction monitoring queue.
func (r *TransactionMonitoringQueueService) Get(ctx context.Context, queueToken string, opts ...option.RequestOption) (res *Queue, err error) {
	opts = slices.Concat(r.Options, opts)
	if queueToken == "" {
		err = errors.New("missing required queue_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/queues/%s", queueToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a transaction monitoring queue.
func (r *TransactionMonitoringQueueService) Update(ctx context.Context, queueToken string, body TransactionMonitoringQueueUpdateParams, opts ...option.RequestOption) (res *Queue, err error) {
	opts = slices.Concat(r.Options, opts)
	if queueToken == "" {
		err = errors.New("missing required queue_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/queues/%s", queueToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists transaction monitoring queues.
func (r *TransactionMonitoringQueueService) List(ctx context.Context, query TransactionMonitoringQueueListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Queue], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/transaction_monitoring/queues"
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

// Lists transaction monitoring queues.
func (r *TransactionMonitoringQueueService) ListAutoPaging(ctx context.Context, query TransactionMonitoringQueueListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Queue] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a transaction monitoring queue.
func (r *TransactionMonitoringQueueService) Delete(ctx context.Context, queueToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if queueToken == "" {
		err = errors.New("missing required queue_token parameter")
		return err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/queues/%s", queueToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A queue that groups transaction monitoring cases for review
type Queue struct {
	// Globally unique identifier for the queue
	Token string `json:"token" api:"required" format:"uuid"`
	// Number of cases in the queue, broken down by status. A status is omitted when
	// the queue has no cases in that status
	CaseCounts QueueCaseCounts `json:"case_counts" api:"required"`
	// Date and time at which the queue was created
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Optional description of the queue
	Description string `json:"description" api:"required,nullable"`
	// Human-readable name of the queue
	Name string `json:"name" api:"required"`
	// Date and time at which the queue was last updated
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	JSON    queueJSON `json:"-"`
}

// queueJSON contains the JSON metadata for the struct [Queue]
type queueJSON struct {
	Token       apijson.Field
	CaseCounts  apijson.Field
	Created     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Queue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueJSON) RawJSON() string {
	return r.raw
}

// Number of cases in the queue, broken down by status. A status is omitted when
// the queue has no cases in that status
type QueueCaseCounts struct {
	// Number of cases in the queue with status `ASSIGNED`
	Assigned int64 `json:"ASSIGNED"`
	// Number of cases in the queue with status `CLOSED`
	Closed int64 `json:"CLOSED"`
	// Number of cases in the queue with status `ESCALATED`
	Escalated int64 `json:"ESCALATED"`
	// Number of cases in the queue with status `IN_REVIEW`
	InReview int64 `json:"IN_REVIEW"`
	// Number of cases in the queue with status `OPEN`
	Open int64 `json:"OPEN"`
	// Number of cases in the queue with status `RESOLVED`
	Resolved int64               `json:"RESOLVED"`
	JSON     queueCaseCountsJSON `json:"-"`
}

// queueCaseCountsJSON contains the JSON metadata for the struct [QueueCaseCounts]
type queueCaseCountsJSON struct {
	Assigned    apijson.Field
	Closed      apijson.Field
	Escalated   apijson.Field
	InReview    apijson.Field
	Open        apijson.Field
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueCaseCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queueCaseCountsJSON) RawJSON() string {
	return r.raw
}

type TransactionMonitoringQueueNewParams struct {
	// Human-readable name of the queue
	Name param.Field[string] `json:"name" api:"required"`
	// Optional description of the queue
	Description param.Field[string] `json:"description"`
}

func (r TransactionMonitoringQueueNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionMonitoringQueueUpdateParams struct {
	// New description for the queue, or `null` to clear it
	Description param.Field[string] `json:"description"`
	// New name for the queue
	Name param.Field[string] `json:"name"`
}

func (r TransactionMonitoringQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionMonitoringQueueListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [TransactionMonitoringQueueListParams]'s query parameters as
// `url.Values`.
func (r TransactionMonitoringQueueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
