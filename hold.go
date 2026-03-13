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

// HoldService contains methods and other services that help with interacting with
// the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHoldService] method instead.
type HoldService struct {
	Options []option.RequestOption
}

// NewHoldService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHoldService(opts ...option.RequestOption) (r *HoldService) {
	r = &HoldService{}
	r.Options = opts
	return
}

// Create a hold on a financial account. Holds reserve funds by moving them from
// available to pending balance. They can be resolved via settlement (linked to a
// payment or book transfer), voiding, or expiration.
func (r *HoldService) New(ctx context.Context, financialAccountToken string, body HoldNewParams, opts ...option.RequestOption) (res *Hold, err error) {
	opts = slices.Concat(r.Options, opts)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/holds", financialAccountToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get hold by token.
func (r *HoldService) Get(ctx context.Context, holdToken string, opts ...option.RequestOption) (res *Hold, err error) {
	opts = slices.Concat(r.Options, opts)
	if holdToken == "" {
		err = errors.New("missing required hold_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/holds/%s", holdToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List holds for a financial account.
func (r *HoldService) List(ctx context.Context, financialAccountToken string, query HoldListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Hold], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if financialAccountToken == "" {
		err = errors.New("missing required financial_account_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/financial_accounts/%s/holds", financialAccountToken)
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

// List holds for a financial account.
func (r *HoldService) ListAutoPaging(ctx context.Context, financialAccountToken string, query HoldListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Hold] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, financialAccountToken, query, opts...))
}

// Void an active hold. This returns the held funds from pending back to available
// balance. Only holds in PENDING status can be voided.
func (r *HoldService) Void(ctx context.Context, holdToken string, body HoldVoidParams, opts ...option.RequestOption) (res *Hold, err error) {
	opts = slices.Concat(r.Options, opts)
	if holdToken == "" {
		err = errors.New("missing required hold_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/holds/%s/void", holdToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A hold transaction representing reserved funds on a financial account. Holds
// move funds from available to pending balance in anticipation of future payments.
// They can be resolved via settlement (linked to payment), manual release, or
// expiration.
type Hold struct {
	// Unique identifier for the transaction
	Token string `json:"token" api:"required" format:"uuid"`
	// ISO 8601 timestamp of when the transaction was created
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Status of a hold transaction
	Status HoldStatus `json:"status" api:"required"`
	// ISO 8601 timestamp of when the transaction was last updated
	Updated  time.Time   `json:"updated" api:"required" format:"date-time"`
	Currency string      `json:"currency"`
	Events   []HoldEvent `json:"events"`
	// When the hold will auto-expire if not resolved
	ExpirationDatetime time.Time `json:"expiration_datetime" api:"nullable" format:"date-time"`
	// HOLD - Hold Transaction
	Family                HoldFamily `json:"family"`
	FinancialAccountToken string     `json:"financial_account_token" format:"uuid"`
	// Current pending amount (0 when resolved)
	PendingAmount int64      `json:"pending_amount"`
	Result        HoldResult `json:"result"`
	UserDefinedID string     `json:"user_defined_id" api:"nullable"`
	JSON          holdJSON   `json:"-"`
}

// holdJSON contains the JSON metadata for the struct [Hold]
type holdJSON struct {
	Token                 apijson.Field
	Created               apijson.Field
	Status                apijson.Field
	Updated               apijson.Field
	Currency              apijson.Field
	Events                apijson.Field
	ExpirationDatetime    apijson.Field
	Family                apijson.Field
	FinancialAccountToken apijson.Field
	PendingAmount         apijson.Field
	Result                apijson.Field
	UserDefinedID         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Hold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdJSON) RawJSON() string {
	return r.raw
}

func (r Hold) implementsAccountActivityListResponse() {}

func (r Hold) implementsAccountActivityGetTransactionResponse() {}

// Status of a hold transaction
type HoldStatus string

const (
	HoldStatusPending  HoldStatus = "PENDING"
	HoldStatusSettled  HoldStatus = "SETTLED"
	HoldStatusExpired  HoldStatus = "EXPIRED"
	HoldStatusVoided   HoldStatus = "VOIDED"
	HoldStatusDeclined HoldStatus = "DECLINED"
	HoldStatusReversed HoldStatus = "REVERSED"
	HoldStatusCanceled HoldStatus = "CANCELED"
	HoldStatusReturned HoldStatus = "RETURNED"
)

func (r HoldStatus) IsKnown() bool {
	switch r {
	case HoldStatusPending, HoldStatusSettled, HoldStatusExpired, HoldStatusVoided, HoldStatusDeclined, HoldStatusReversed, HoldStatusCanceled, HoldStatusReturned:
		return true
	}
	return false
}

// HOLD - Hold Transaction
type HoldFamily string

const (
	HoldFamilyHold HoldFamily = "HOLD"
)

func (r HoldFamily) IsKnown() bool {
	switch r {
	case HoldFamilyHold:
		return true
	}
	return false
}

type HoldResult string

const (
	HoldResultApproved HoldResult = "APPROVED"
	HoldResultDeclined HoldResult = "DECLINED"
)

func (r HoldResult) IsKnown() bool {
	switch r {
	case HoldResultApproved, HoldResultDeclined:
		return true
	}
	return false
}

// Event representing a lifecycle change to a hold
type HoldEvent struct {
	Token string `json:"token" api:"required" format:"uuid"`
	// Amount in cents
	Amount          int64                     `json:"amount" api:"required"`
	Created         time.Time                 `json:"created" api:"required" format:"date-time"`
	DetailedResults []HoldEventDetailedResult `json:"detailed_results" api:"required"`
	Memo            string                    `json:"memo" api:"required,nullable"`
	Result          HoldEventResult           `json:"result" api:"required"`
	// Transaction token of the payment that settled this hold (only populated for
	// HOLD_SETTLED events)
	SettlingTransactionToken string `json:"settling_transaction_token" api:"required,nullable" format:"uuid"`
	// Type of hold lifecycle event
	Type HoldEventType `json:"type" api:"required"`
	JSON holdEventJSON `json:"-"`
}

// holdEventJSON contains the JSON metadata for the struct [HoldEvent]
type holdEventJSON struct {
	Token                    apijson.Field
	Amount                   apijson.Field
	Created                  apijson.Field
	DetailedResults          apijson.Field
	Memo                     apijson.Field
	Result                   apijson.Field
	SettlingTransactionToken apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *HoldEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdEventJSON) RawJSON() string {
	return r.raw
}

type HoldEventDetailedResult string

const (
	HoldEventDetailedResultApproved          HoldEventDetailedResult = "APPROVED"
	HoldEventDetailedResultInsufficientFunds HoldEventDetailedResult = "INSUFFICIENT_FUNDS"
)

func (r HoldEventDetailedResult) IsKnown() bool {
	switch r {
	case HoldEventDetailedResultApproved, HoldEventDetailedResultInsufficientFunds:
		return true
	}
	return false
}

type HoldEventResult string

const (
	HoldEventResultApproved HoldEventResult = "APPROVED"
	HoldEventResultDeclined HoldEventResult = "DECLINED"
)

func (r HoldEventResult) IsKnown() bool {
	switch r {
	case HoldEventResultApproved, HoldEventResultDeclined:
		return true
	}
	return false
}

// Type of hold lifecycle event
type HoldEventType string

const (
	HoldEventTypeHoldInitiated HoldEventType = "HOLD_INITIATED"
	HoldEventTypeHoldVoided    HoldEventType = "HOLD_VOIDED"
	HoldEventTypeHoldExpired   HoldEventType = "HOLD_EXPIRED"
	HoldEventTypeHoldSettled   HoldEventType = "HOLD_SETTLED"
)

func (r HoldEventType) IsKnown() bool {
	switch r {
	case HoldEventTypeHoldInitiated, HoldEventTypeHoldVoided, HoldEventTypeHoldExpired, HoldEventTypeHoldSettled:
		return true
	}
	return false
}

type HoldNewParams struct {
	// Amount to hold in cents
	Amount param.Field[int64] `json:"amount" api:"required"`
	// Customer-provided token for idempotency. Becomes the hold token.
	Token param.Field[string] `json:"token" format:"uuid"`
	// When the hold should auto-expire
	ExpirationDatetime param.Field[time.Time] `json:"expiration_datetime" format:"date-time"`
	// Reason for the hold
	Memo param.Field[string] `json:"memo"`
	// User-provided identifier for the hold
	UserDefinedID param.Field[string] `json:"user_defined_id"`
}

func (r HoldNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HoldListParams struct {
	// Date string in RFC 3339 format. Only entries created after the specified time
	// will be included. UTC time zone.
	Begin param.Field[time.Time] `query:"begin" format:"date-time"`
	// Date string in RFC 3339 format. Only entries created before the specified time
	// will be included. UTC time zone.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after"`
	// Hold status to filter by.
	Status param.Field[HoldListParamsStatus] `query:"status"`
}

// URLQuery serializes [HoldListParams]'s query parameters as `url.Values`.
func (r HoldListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Hold status to filter by.
type HoldListParamsStatus string

const (
	HoldListParamsStatusPending HoldListParamsStatus = "PENDING"
	HoldListParamsStatusSettled HoldListParamsStatus = "SETTLED"
	HoldListParamsStatusExpired HoldListParamsStatus = "EXPIRED"
	HoldListParamsStatusVoided  HoldListParamsStatus = "VOIDED"
)

func (r HoldListParamsStatus) IsKnown() bool {
	switch r {
	case HoldListParamsStatusPending, HoldListParamsStatusSettled, HoldListParamsStatusExpired, HoldListParamsStatusVoided:
		return true
	}
	return false
}

type HoldVoidParams struct {
	// Reason for voiding the hold
	Memo param.Field[string] `json:"memo"`
}

func (r HoldVoidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
