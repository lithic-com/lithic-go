// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// TransactionMonitoringCaseCommentService contains methods and other services that
// help with interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionMonitoringCaseCommentService] method instead.
type TransactionMonitoringCaseCommentService struct {
	Options []option.RequestOption
}

// NewTransactionMonitoringCaseCommentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransactionMonitoringCaseCommentService(opts ...option.RequestOption) (r *TransactionMonitoringCaseCommentService) {
	r = &TransactionMonitoringCaseCommentService{}
	r.Options = opts
	return
}

// Adds a comment to a case.
func (r *TransactionMonitoringCaseCommentService) New(ctx context.Context, caseToken string, body TransactionMonitoringCaseCommentNewParams, opts ...option.RequestOption) (res *CaseActivityEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/comments", caseToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Edits an existing comment on a case.
func (r *TransactionMonitoringCaseCommentService) Update(ctx context.Context, caseToken string, commentToken string, body TransactionMonitoringCaseCommentUpdateParams, opts ...option.RequestOption) (res *CaseActivityEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return nil, err
	}
	if commentToken == "" {
		err = errors.New("missing required comment_token parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/comments/%s", caseToken, commentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Deletes a comment from a case.
func (r *TransactionMonitoringCaseCommentService) Delete(ctx context.Context, caseToken string, commentToken string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if caseToken == "" {
		err = errors.New("missing required case_token parameter")
		return err
	}
	if commentToken == "" {
		err = errors.New("missing required comment_token parameter")
		return err
	}
	path := fmt.Sprintf("v1/transaction_monitoring/cases/%s/comments/%s", caseToken, commentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type TransactionMonitoringCaseCommentNewParams struct {
	// Text of the comment
	Comment param.Field[string] `json:"comment" api:"required"`
	// Optional client-provided identifier for the actor performing this action,
	// recorded on the resulting activity entry. This value is supplied by the client
	// (for example, your own internal user ID) and is not authenticated by Lithic
	ActorToken param.Field[string] `json:"actor_token"`
}

func (r TransactionMonitoringCaseCommentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionMonitoringCaseCommentUpdateParams struct {
	// New text of the comment
	Comment param.Field[string] `json:"comment" api:"required"`
	// Optional client-provided identifier for the actor performing this action,
	// recorded on the resulting activity entry. This value is supplied by the client
	// (for example, your own internal user ID) and is not authenticated by Lithic
	ActorToken param.Field[string] `json:"actor_token"`
}

func (r TransactionMonitoringCaseCommentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
