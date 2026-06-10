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

// TransactionMonitoringCaseFileService contains methods and other services that
// help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionMonitoringCaseFileService] method instead.
type TransactionMonitoringCaseFileService struct {
	Options []option.RequestOption
}

// NewTransactionMonitoringCaseFileService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransactionMonitoringCaseFileService(opts ...option.RequestOption) (r *TransactionMonitoringCaseFileService) {
	r = &TransactionMonitoringCaseFileService{}
	r.Options = opts
	return
}

// Creates a file record and returns a presigned URL for uploading the file to the
// case.
func (r *TransactionMonitoringCaseFileService) New(ctx context.Context, caseToken string, body TransactionMonitoringCaseFileNewParams, opts ...option.RequestOption) (res *CaseFile, err error) {
	opts = slices.Concat(r.Options, opts)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/files", caseToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single file attached to a case, including a presigned download URL
// when the file is ready.
func (r *TransactionMonitoringCaseFileService) Get(ctx context.Context, caseToken string, fileToken string, opts ...option.RequestOption) (res *CaseFile, err error) {
	opts = slices.Concat(r.Options, opts)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	if fileToken == "" {
		err = errors.New("missing required file_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/files/%s", caseToken, fileToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the files attached to a case.
func (r *TransactionMonitoringCaseFileService) List(ctx context.Context, caseToken string, query TransactionMonitoringCaseFileListParams, opts ...option.RequestOption) (res *pagination.CursorPage[CaseFile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/files", caseToken)
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

// Lists the files attached to a case.
func (r *TransactionMonitoringCaseFileService) ListAutoPaging(ctx context.Context, caseToken string, query TransactionMonitoringCaseFileListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CaseFile] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, caseToken, query, opts...))
}

// Deletes a file from a case.
func (r *TransactionMonitoringCaseFileService) Delete(ctx context.Context, caseToken string, fileToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return err
	}
	if fileToken == "" {
		err = errors.New("missing required file_token parameter")
		return err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/files/%s", caseToken, fileToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A file attached to a case. Status-dependent fields are always present but may be
// `null`:
//
//   - `upload_url`, `upload_url_expires`, and `upload_constraints` are populated
//     when `status` is `PENDING` or `REJECTED`
//   - `download_url` and `download_url_expires` are populated when `status` is
//     `READY`
//   - `failure_reason` is populated when `status` is `REJECTED`
type CaseFile struct {
	// Globally unique identifier for the file
	Token string `json:"token" api:"required" format:"uuid"`
	// Date and time at which the file record was created
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Presigned URL the client uses to download the file
	DownloadURL string `json:"download_url" api:"required,nullable"`
	// Date and time at which the download URL expires
	DownloadURLExpires time.Time `json:"download_url_expires" api:"required,nullable" format:"date-time"`
	// Reason the file was rejected, when applicable
	FailureReason string `json:"failure_reason" api:"required,nullable"`
	// MIME type of the file, available once the file is ready
	MimeType string `json:"mime_type" api:"required,nullable"`
	// Name of the file
	Name string `json:"name" api:"required"`
	// Size of the file in bytes, available once the file is ready
	SizeBytes int64 `json:"size_bytes" api:"required,nullable"`
	// Lifecycle status of a case file:
	//
	//   - `PENDING` - An upload URL has been issued and the file is awaiting upload
	//   - `READY` - The file has been uploaded and validated; a download URL is
	//     available
	//   - `REJECTED` - File validation failed; see `failure_reason` for details
	Status FileStatus `json:"status" api:"required"`
	// Date and time at which the file record was last updated
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Constraints applied to a file upload, returned alongside the upload URL so
	// clients can validate before uploading
	UploadConstraints UploadConstraints `json:"upload_constraints" api:"required,nullable"`
	// Presigned URL the client uses to upload the file
	UploadURL string `json:"upload_url" api:"required,nullable"`
	// Date and time at which the upload URL expires
	UploadURLExpires time.Time    `json:"upload_url_expires" api:"required,nullable" format:"date-time"`
	JSON             caseFileJSON `json:"-"`
}

// caseFileJSON contains the JSON metadata for the struct [CaseFile]
type caseFileJSON struct {
	Token              apijson.Field
	Created            apijson.Field
	DownloadURL        apijson.Field
	DownloadURLExpires apijson.Field
	FailureReason      apijson.Field
	MimeType           apijson.Field
	Name               apijson.Field
	SizeBytes          apijson.Field
	Status             apijson.Field
	Updated            apijson.Field
	UploadConstraints  apijson.Field
	UploadURL          apijson.Field
	UploadURLExpires   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CaseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseFileJSON) RawJSON() string {
	return r.raw
}

// Lifecycle status of a case file:
//
//   - `PENDING` - An upload URL has been issued and the file is awaiting upload
//   - `READY` - The file has been uploaded and validated; a download URL is
//     available
//   - `REJECTED` - File validation failed; see `failure_reason` for details
type FileStatus string

const (
	FileStatusPending  FileStatus = "PENDING"
	FileStatusReady    FileStatus = "READY"
	FileStatusRejected FileStatus = "REJECTED"
)

func (r FileStatus) IsKnown() bool {
	switch r {
	case FileStatusPending, FileStatusReady, FileStatusRejected:
		return true
	}
	return false
}

// Constraints applied to a file upload, returned alongside the upload URL so
// clients can validate before uploading
type UploadConstraints struct {
	// MIME types accepted for the upload
	AcceptedMimeTypes []string `json:"accepted_mime_types" api:"required"`
	// Maximum accepted file size, in bytes
	MaxSizeBytes int64                 `json:"max_size_bytes" api:"required"`
	JSON         uploadConstraintsJSON `json:"-"`
}

// uploadConstraintsJSON contains the JSON metadata for the struct
// [UploadConstraints]
type uploadConstraintsJSON struct {
	AcceptedMimeTypes apijson.Field
	MaxSizeBytes      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UploadConstraints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadConstraintsJSON) RawJSON() string {
	return r.raw
}

type TransactionMonitoringCaseFileNewParams struct {
	// Name of the file to upload
	Name param.Field[string] `json:"name" api:"required"`
}

func (r TransactionMonitoringCaseFileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionMonitoringCaseFileListParams struct {
	// A cursor representing an item's token before which a page of results should end.
	// Used to retrieve the previous page of results before this item.
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Page size (for pagination).
	PageSize param.Field[int64] `query:"page_size"`
	// A cursor representing an item's token after which a page of results should
	// begin. Used to retrieve the next page of results after this item.
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [TransactionMonitoringCaseFileListParams]'s query parameters
// as `url.Values`.
func (r TransactionMonitoringCaseFileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
