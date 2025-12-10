// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"time"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/option"
)

// InternalTransactionService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInternalTransactionService] method instead.
type InternalTransactionService struct {
	Options []option.RequestOption
}

// NewInternalTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInternalTransactionService(opts ...option.RequestOption) (r *InternalTransactionService) {
	r = &InternalTransactionService{}
	r.Options = opts
	return
}

type InternalTransaction struct {
	Token         string                      `json:"token,required" format:"uuid"`
	Category      InternalTransactionCategory `json:"category,required"`
	Created       time.Time                   `json:"created,required" format:"date-time"`
	Currency      string                      `json:"currency,required"`
	Descriptor    string                      `json:"descriptor,required"`
	Events        []InternalTransactionEvent  `json:"events,required"`
	PendingAmount int64                       `json:"pending_amount,required"`
	Result        InternalTransactionResult   `json:"result,required"`
	SettledAmount int64                       `json:"settled_amount,required"`
	Status        InternalTransactionStatus   `json:"status,required"`
	Updated       time.Time                   `json:"updated,required" format:"date-time"`
	JSON          internalTransactionJSON     `json:"-"`
}

// internalTransactionJSON contains the JSON metadata for the struct
// [InternalTransaction]
type internalTransactionJSON struct {
	Token         apijson.Field
	Category      apijson.Field
	Created       apijson.Field
	Currency      apijson.Field
	Descriptor    apijson.Field
	Events        apijson.Field
	PendingAmount apijson.Field
	Result        apijson.Field
	SettledAmount apijson.Field
	Status        apijson.Field
	Updated       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InternalTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r internalTransactionJSON) RawJSON() string {
	return r.raw
}

type InternalTransactionCategory string

const (
	InternalTransactionCategoryInternal InternalTransactionCategory = "INTERNAL"
)

func (r InternalTransactionCategory) IsKnown() bool {
	switch r {
	case InternalTransactionCategoryInternal:
		return true
	}
	return false
}

type InternalTransactionEvent struct {
	Token   string                          `json:"token,required" format:"uuid"`
	Amount  int64                           `json:"amount,required"`
	Created time.Time                       `json:"created,required" format:"date-time"`
	Result  InternalTransactionEventsResult `json:"result,required"`
	Type    InternalTransactionEventsType   `json:"type,required"`
	JSON    internalTransactionEventJSON    `json:"-"`
}

// internalTransactionEventJSON contains the JSON metadata for the struct
// [InternalTransactionEvent]
type internalTransactionEventJSON struct {
	Token       apijson.Field
	Amount      apijson.Field
	Created     apijson.Field
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InternalTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r internalTransactionEventJSON) RawJSON() string {
	return r.raw
}

type InternalTransactionEventsResult string

const (
	InternalTransactionEventsResultApproved InternalTransactionEventsResult = "APPROVED"
	InternalTransactionEventsResultDeclined InternalTransactionEventsResult = "DECLINED"
)

func (r InternalTransactionEventsResult) IsKnown() bool {
	switch r {
	case InternalTransactionEventsResultApproved, InternalTransactionEventsResultDeclined:
		return true
	}
	return false
}

type InternalTransactionEventsType string

const (
	InternalTransactionEventsTypeInternalAdjustment InternalTransactionEventsType = "INTERNAL_ADJUSTMENT"
)

func (r InternalTransactionEventsType) IsKnown() bool {
	switch r {
	case InternalTransactionEventsTypeInternalAdjustment:
		return true
	}
	return false
}

type InternalTransactionResult string

const (
	InternalTransactionResultApproved InternalTransactionResult = "APPROVED"
	InternalTransactionResultDeclined InternalTransactionResult = "DECLINED"
)

func (r InternalTransactionResult) IsKnown() bool {
	switch r {
	case InternalTransactionResultApproved, InternalTransactionResultDeclined:
		return true
	}
	return false
}

type InternalTransactionStatus string

const (
	InternalTransactionStatusPending  InternalTransactionStatus = "PENDING"
	InternalTransactionStatusSettled  InternalTransactionStatus = "SETTLED"
	InternalTransactionStatusDeclined InternalTransactionStatus = "DECLINED"
	InternalTransactionStatusReversed InternalTransactionStatus = "REVERSED"
	InternalTransactionStatusCanceled InternalTransactionStatus = "CANCELED"
	InternalTransactionStatusReturned InternalTransactionStatus = "RETURNED"
)

func (r InternalTransactionStatus) IsKnown() bool {
	switch r {
	case InternalTransactionStatusPending, InternalTransactionStatusSettled, InternalTransactionStatusDeclined, InternalTransactionStatusReversed, InternalTransactionStatusCanceled, InternalTransactionStatusReturned:
		return true
	}
	return false
}
