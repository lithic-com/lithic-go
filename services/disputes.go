package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/pagination"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type DisputeService struct {
	Options []options.RequestOption
}

func NewDisputeService(opts ...options.RequestOption) (r *DisputeService) {
	r = &DisputeService{}
	r.Options = opts
	return
}

// Initiate a dispute.
func (r *DisputeService) New(ctx context.Context, body *requests.DisputeNewParams, opts ...options.RequestOption) (res *responses.Dispute, err error) {
	opts = append(r.Options[:], opts...)
	path := "disputes"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get dispute.
func (r *DisputeService) Get(ctx context.Context, dispute_token string, opts ...options.RequestOption) (res *responses.Dispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s", dispute_token)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update dispute. Can only be modified if status is `NEW`.
func (r *DisputeService) Update(ctx context.Context, dispute_token string, body *requests.DisputeUpdateParams, opts ...options.RequestOption) (res *responses.Dispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s", dispute_token)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List disputes.
func (r *DisputeService) List(ctx context.Context, query *requests.DisputeListParams, opts ...options.RequestOption) (res *responses.DisputesCursorPage, err error) {
	opts = append(r.Options, opts...)
	path := "disputes"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.DisputesCursorPage{
		CursorPage: &pagination.CursorPage[responses.Dispute]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Withdraw dispute.
func (r *DisputeService) Delete(ctx context.Context, dispute_token string, opts ...options.RequestOption) (res *responses.Dispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s", dispute_token)
	err = options.ExecuteNewRequest(ctx, "DELETE", path, nil, &res, opts...)
	return
}

// Soft delete evidence for a dispute. Evidence will not be reviewed or submitted
// by Lithic after it is withdrawn.
func (r *DisputeService) DeleteEvidence(ctx context.Context, evidence_token string, params *requests.DisputesDeleteEvidenceParams, opts ...options.RequestOption) (res *responses.DisputeEvidence, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s/evidences/%s", params.DisputeToken, evidence_token)
	err = options.ExecuteNewRequest(ctx, "DELETE", path, params, &res, opts...)
	return
}

// Use this endpoint to upload evidences for the dispute. It will return a URL to
// upload your documents to. The URL will expire in 30 minutes.
//
// Uploaded documents must either be a `jpg`, `png` or `pdf` file, and each must be
// less than 5 GiB.
func (r *DisputeService) InitiateEvidenceUpload(ctx context.Context, dispute_token string, opts ...options.RequestOption) (res *responses.DisputeInitiateEvidenceUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s/evidences", dispute_token)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// List evidence metadata for a dispute.
func (r *DisputeService) ListEvidences(ctx context.Context, dispute_token string, query *requests.DisputeListEvidencesParams, opts ...options.RequestOption) (res *responses.DisputeEvidencesCursorPage, err error) {
	opts = append(r.Options, opts...)
	path := fmt.Sprintf("disputes/%s/evidences", dispute_token)
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.DisputeEvidencesCursorPage{
		CursorPage: &pagination.CursorPage[responses.DisputeEvidence]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Get a dispute's evidence metadata.
func (r *DisputeService) GetEvidence(ctx context.Context, evidence_token string, params *requests.DisputesGetEvidenceParams, opts ...options.RequestOption) (res *responses.DisputeEvidence, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("disputes/%s/evidences/%s", params.DisputeToken, evidence_token)
	err = options.ExecuteNewRequest(ctx, "GET", path, params, &res, opts...)
	return
}

func (r *DisputeService) UploadEvidence(ctx context.Context, disputeToken string, file io.Reader, opts ...options.RequestOption) (err error) {
	payload, err := r.InitiateEvidenceUpload(ctx, disputeToken, opts...)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	name := "anonymous_file"
	if nameable, ok := file.(interface{ Name() string }); ok {
		name = nameable.Name()
	}
	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		return err
	}
	io.Copy(part, file)

	req, err := http.NewRequestWithContext(ctx, "PUT", payload.UploadURL, body)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return

}
