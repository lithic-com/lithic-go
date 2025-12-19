// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
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

// TransferLimitService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferLimitService] method instead.
type TransferLimitService struct {
	Options []option.RequestOption
}

// NewTransferLimitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransferLimitService(opts ...option.RequestOption) (r *TransferLimitService) {
	r = &TransferLimitService{}
	r.Options = opts
	return
}

// Get transfer limits for a specified date
func (r *TransferLimitService) List(ctx context.Context, query TransferLimitListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TransferLimitsResponseData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/transfer_limits"
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

// Get transfer limits for a specified date
func (r *TransferLimitService) ListAutoPaging(ctx context.Context, query TransferLimitListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TransferLimitsResponseData] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type TransferLimitsResponse struct {
	// List of transfer limits
	Data []TransferLimitsResponseData `json:"data,required"`
	// Whether there are more transfer limits
	HasMore bool                       `json:"has_more,required"`
	JSON    transferLimitsResponseJSON `json:"-"`
}

// transferLimitsResponseJSON contains the JSON metadata for the struct
// [TransferLimitsResponse]
type transferLimitsResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferLimitsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseJSON) RawJSON() string {
	return r.raw
}

type TransferLimitsResponseData struct {
	// Company ID
	CompanyID string `json:"company_id,required"`
	// Daily limits with progress
	DailyLimit TransferLimitsResponseDataDailyLimit `json:"daily_limit,required"`
	// The date for the limit view (ISO format)
	Date time.Time `json:"date,required" format:"date"`
	// Whether the company is a FBO; based on the company ID prefix
	IsFbo bool `json:"is_fbo,required"`
	// Monthly limits with progress
	MonthlyLimit TransferLimitsResponseDataMonthlyLimit `json:"monthly_limit,required"`
	// Program transaction limits
	ProgramLimitPerTransaction TransferLimitsResponseDataProgramLimitPerTransaction `json:"program_limit_per_transaction,required"`
	JSON                       transferLimitsResponseDataJSON                       `json:"-"`
}

// transferLimitsResponseDataJSON contains the JSON metadata for the struct
// [TransferLimitsResponseData]
type transferLimitsResponseDataJSON struct {
	CompanyID                  apijson.Field
	DailyLimit                 apijson.Field
	Date                       apijson.Field
	IsFbo                      apijson.Field
	MonthlyLimit               apijson.Field
	ProgramLimitPerTransaction apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *TransferLimitsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Daily limits with progress
type TransferLimitsResponseDataDailyLimit struct {
	// Credit limits
	Credit TransferLimitsResponseDataDailyLimitCredit `json:"credit,required"`
	// Debit limits
	Debit TransferLimitsResponseDataDailyLimitDebit `json:"debit,required"`
	JSON  transferLimitsResponseDataDailyLimitJSON  `json:"-"`
}

// transferLimitsResponseDataDailyLimitJSON contains the JSON metadata for the
// struct [TransferLimitsResponseDataDailyLimit]
type transferLimitsResponseDataDailyLimitJSON struct {
	Credit      apijson.Field
	Debit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferLimitsResponseDataDailyLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataDailyLimitJSON) RawJSON() string {
	return r.raw
}

// Credit limits
type TransferLimitsResponseDataDailyLimitCredit struct {
	// The limit amount
	Limit int64 `json:"limit,required"`
	// Amount originated towards limit
	AmountOriginated int64                                          `json:"amount_originated"`
	JSON             transferLimitsResponseDataDailyLimitCreditJSON `json:"-"`
}

// transferLimitsResponseDataDailyLimitCreditJSON contains the JSON metadata for
// the struct [TransferLimitsResponseDataDailyLimitCredit]
type transferLimitsResponseDataDailyLimitCreditJSON struct {
	Limit            apijson.Field
	AmountOriginated apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransferLimitsResponseDataDailyLimitCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataDailyLimitCreditJSON) RawJSON() string {
	return r.raw
}

// Debit limits
type TransferLimitsResponseDataDailyLimitDebit struct {
	// The limit amount
	Limit int64 `json:"limit,required"`
	// Amount originated towards limit
	AmountOriginated int64                                         `json:"amount_originated"`
	JSON             transferLimitsResponseDataDailyLimitDebitJSON `json:"-"`
}

// transferLimitsResponseDataDailyLimitDebitJSON contains the JSON metadata for the
// struct [TransferLimitsResponseDataDailyLimitDebit]
type transferLimitsResponseDataDailyLimitDebitJSON struct {
	Limit            apijson.Field
	AmountOriginated apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransferLimitsResponseDataDailyLimitDebit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataDailyLimitDebitJSON) RawJSON() string {
	return r.raw
}

// Monthly limits with progress
type TransferLimitsResponseDataMonthlyLimit struct {
	// Credit limits
	Credit TransferLimitsResponseDataMonthlyLimitCredit `json:"credit,required"`
	// Debit limits
	Debit TransferLimitsResponseDataMonthlyLimitDebit `json:"debit,required"`
	JSON  transferLimitsResponseDataMonthlyLimitJSON  `json:"-"`
}

// transferLimitsResponseDataMonthlyLimitJSON contains the JSON metadata for the
// struct [TransferLimitsResponseDataMonthlyLimit]
type transferLimitsResponseDataMonthlyLimitJSON struct {
	Credit      apijson.Field
	Debit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferLimitsResponseDataMonthlyLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataMonthlyLimitJSON) RawJSON() string {
	return r.raw
}

// Credit limits
type TransferLimitsResponseDataMonthlyLimitCredit struct {
	// The limit amount
	Limit int64 `json:"limit,required"`
	// Amount originated towards limit
	AmountOriginated int64                                            `json:"amount_originated"`
	JSON             transferLimitsResponseDataMonthlyLimitCreditJSON `json:"-"`
}

// transferLimitsResponseDataMonthlyLimitCreditJSON contains the JSON metadata for
// the struct [TransferLimitsResponseDataMonthlyLimitCredit]
type transferLimitsResponseDataMonthlyLimitCreditJSON struct {
	Limit            apijson.Field
	AmountOriginated apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransferLimitsResponseDataMonthlyLimitCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataMonthlyLimitCreditJSON) RawJSON() string {
	return r.raw
}

// Debit limits
type TransferLimitsResponseDataMonthlyLimitDebit struct {
	// The limit amount
	Limit int64 `json:"limit,required"`
	// Amount originated towards limit
	AmountOriginated int64                                           `json:"amount_originated"`
	JSON             transferLimitsResponseDataMonthlyLimitDebitJSON `json:"-"`
}

// transferLimitsResponseDataMonthlyLimitDebitJSON contains the JSON metadata for
// the struct [TransferLimitsResponseDataMonthlyLimitDebit]
type transferLimitsResponseDataMonthlyLimitDebitJSON struct {
	Limit            apijson.Field
	AmountOriginated apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransferLimitsResponseDataMonthlyLimitDebit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataMonthlyLimitDebitJSON) RawJSON() string {
	return r.raw
}

// Program transaction limits
type TransferLimitsResponseDataProgramLimitPerTransaction struct {
	// Credit limits
	Credit TransferLimitsResponseDataProgramLimitPerTransactionCredit `json:"credit,required"`
	// Debit limits
	Debit TransferLimitsResponseDataProgramLimitPerTransactionDebit `json:"debit,required"`
	JSON  transferLimitsResponseDataProgramLimitPerTransactionJSON  `json:"-"`
}

// transferLimitsResponseDataProgramLimitPerTransactionJSON contains the JSON
// metadata for the struct [TransferLimitsResponseDataProgramLimitPerTransaction]
type transferLimitsResponseDataProgramLimitPerTransactionJSON struct {
	Credit      apijson.Field
	Debit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferLimitsResponseDataProgramLimitPerTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataProgramLimitPerTransactionJSON) RawJSON() string {
	return r.raw
}

// Credit limits
type TransferLimitsResponseDataProgramLimitPerTransactionCredit struct {
	// The limit amount
	Limit int64 `json:"limit,required"`
	// Amount originated towards limit
	AmountOriginated int64                                                          `json:"amount_originated"`
	JSON             transferLimitsResponseDataProgramLimitPerTransactionCreditJSON `json:"-"`
}

// transferLimitsResponseDataProgramLimitPerTransactionCreditJSON contains the JSON
// metadata for the struct
// [TransferLimitsResponseDataProgramLimitPerTransactionCredit]
type transferLimitsResponseDataProgramLimitPerTransactionCreditJSON struct {
	Limit            apijson.Field
	AmountOriginated apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransferLimitsResponseDataProgramLimitPerTransactionCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataProgramLimitPerTransactionCreditJSON) RawJSON() string {
	return r.raw
}

// Debit limits
type TransferLimitsResponseDataProgramLimitPerTransactionDebit struct {
	// The limit amount
	Limit int64 `json:"limit,required"`
	// Amount originated towards limit
	AmountOriginated int64                                                         `json:"amount_originated"`
	JSON             transferLimitsResponseDataProgramLimitPerTransactionDebitJSON `json:"-"`
}

// transferLimitsResponseDataProgramLimitPerTransactionDebitJSON contains the JSON
// metadata for the struct
// [TransferLimitsResponseDataProgramLimitPerTransactionDebit]
type transferLimitsResponseDataProgramLimitPerTransactionDebitJSON struct {
	Limit            apijson.Field
	AmountOriginated apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransferLimitsResponseDataProgramLimitPerTransactionDebit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLimitsResponseDataProgramLimitPerTransactionDebitJSON) RawJSON() string {
	return r.raw
}

type TransferLimitListParams struct {
	// Date for which to retrieve transfer limits (ISO 8601 format)
	Date param.Field[time.Time] `query:"date" format:"date"`
}

// URLQuery serializes [TransferLimitListParams]'s query parameters as
// `url.Values`.
func (r TransferLimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
