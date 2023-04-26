package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type AccountHolderService struct {
	Options []option.RequestOption
}

func NewAccountHolderService(opts ...option.RequestOption) (r *AccountHolderService) {
	r = &AccountHolderService{}
	r.Options = opts
	return
}

// Run an individual or business's information through the Customer Identification
// Program (CIP) and return an `account_token` if the status is accepted or pending
// (i.e., further action required). All calls to this endpoint will return an
// immediate response - though in some cases, the response may indicate the
// workflow is under review or further action will be needed to complete the
// account creation process. This endpoint can only be used on accounts that are
// part of the program the calling API key manages.
//
// Note: If you choose to set a timeout for this request, we recommend 5 minutes.
func (r *AccountHolderService) New(ctx context.Context, body requests.AccountHolderNewParams, opts ...option.RequestOption) (res *responses.AccountHolder, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_holders"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an Individual or Business Account Holder and/or their KYC or KYB evaluation
// status.
func (r *AccountHolderService) Get(ctx context.Context, account_holder_token string, opts ...option.RequestOption) (res *responses.AccountHolder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s", account_holder_token)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the information associated with a particular account holder.
func (r *AccountHolderService) Update(ctx context.Context, account_holder_token string, body *requests.AccountHolderUpdateParams, opts ...option.RequestOption) (res *responses.AccountHolderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s", account_holder_token)
	err = option.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Create a webhook to receive KYC or KYB evaluation events.
//
// There are two types of account holder webhooks:
//
//   - `verification`: Webhook sent when the status of a KYC or KYB evaluation
//     changes from `PENDING_DOCUMENT` (KYC) or `PENDING` (KYB) to `ACCEPTED` or
//     `REJECTED`.
//   - `document_upload_front`/`document_upload_back`: Webhook sent when a document
//     upload fails.
//
// After a webhook has been created, this endpoint can be used to rotate a webhooks
// HMAC token or modify the registered URL. Only a single webhook is allowed per
// program. Since HMAC verification is available, the IP addresses from which
// KYC/KYB webhooks are sent are subject to change.
func (r *AccountHolderService) NewWebhook(ctx context.Context, body *requests.AccountHolderNewWebhookParams, opts ...option.RequestOption) (res *responses.AccountHolderCreateWebhookResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "webhooks/account_holders"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the status of account holder document uploads, or retrieve the upload
// URLs to process your image uploads.
//
// Note that this is not equivalent to checking the status of the KYC evaluation
// overall (a document may be successfully uploaded but not be sufficient for KYC
// to pass).
//
// In the event your upload URLs have expired, calling this endpoint will refresh
// them. Similarly, in the event a previous account holder document upload has
// failed, you can use this endpoint to get a new upload URL for the failed image
// upload.
//
// When a new document upload is generated for a failed attempt, the response will
// show an additional entry in the `required_document_uploads` list in a `PENDING`
// state for the corresponding `image_type`.
func (r *AccountHolderService) ListDocuments(ctx context.Context, account_holder_token string, opts ...option.RequestOption) (res *responses.AccountHolderListDocumentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s/documents", account_holder_token)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resubmit a KYC submission. This endpoint should be used in cases where a KYC
// submission returned a `PENDING_RESUBMIT` result, meaning one or more critical
// KYC fields may have been mis-entered and the individual's identity has not yet
// been successfully verified. This step must be completed in order to proceed with
// the KYC evaluation.
//
// Two resubmission attempts are permitted via this endpoint before a `REJECTED`
// status is returned and the account creation process is ended.
func (r *AccountHolderService) Resubmit(ctx context.Context, account_holder_token string, body *requests.AccountHolderResubmitParams, opts ...option.RequestOption) (res *responses.AccountHolder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s/resubmit", account_holder_token)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check the status of an account holder document upload, or retrieve the upload
// URLs to process your image uploads.
//
// Note that this is not equivalent to checking the status of the KYC evaluation
// overall (a document may be successfully uploaded but not be sufficient for KYC
// to pass).
//
// In the event your upload URLs have expired, calling this endpoint will refresh
// them. Similarly, in the event a document upload has failed, you can use this
// endpoint to get a new upload URL for the failed image upload.
//
// When a new account holder document upload is generated for a failed attempt, the
// response will show an additional entry in the `required_document_uploads` array
// in a `PENDING` state for the corresponding `image_type`.
func (r *AccountHolderService) GetDocument(ctx context.Context, account_holder_token string, document_token string, opts ...option.RequestOption) (res *responses.AccountHolderDocument, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s/documents/%s", account_holder_token, document_token)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to identify which type of supported government-issued
// documentation you will upload for further verification. It will return two URLs
// to upload your document images to - one for the front image and one for the back
// image.
//
// This endpoint is only valid for evaluations in a `PENDING_DOCUMENT` state.
//
// Uploaded images must either be a `jpg` or `png` file, and each must be less than
// 15 MiB. Once both required uploads have been successfully completed, your
// document will be run through KYC verification.
//
// If you have registered a webhook, you will receive evaluation updates for any
// document submission evaluations, as well as for any failed document uploads.
//
// Two document submission attempts are permitted via this endpoint before a
// `REJECTED` status is returned and the account creation process is ended.
// Currently only one type of account holder document is supported per KYC
// verification.
func (r *AccountHolderService) UploadDocument(ctx context.Context, account_holder_token string, body *requests.AccountHolderUploadDocumentParams, opts ...option.RequestOption) (res *responses.AccountHolderDocument, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_holders/%s/documents", account_holder_token)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}
