package services

import (
	"context"
	"fmt"

	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/pagination"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type FundingSourceService struct {
	Options []options.RequestOption
}

func NewFundingSourceService(opts ...options.RequestOption) (r *FundingSourceService) {
	r = &FundingSourceService{}
	r.Options = opts
	return
}

// Add a funding source using bank routing and account numbers or via Plaid.
//
// In the production environment, funding accounts will be set to `PENDING` state
// until micro-deposit validation completes while funding accounts in sandbox will
// be set to `ENABLED` state automatically.
func (r *FundingSourceService) New(ctx context.Context, body *requests.FundingSourceNewParams, opts ...options.RequestOption) (res *responses.FundingSource, err error) {
	opts = append(r.Options[:], opts...)
	path := "funding_sources"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Update a funding source.
func (r *FundingSourceService) Update(ctx context.Context, funding_source_token string, body *requests.FundingSourceUpdateParams, opts ...options.RequestOption) (res *responses.FundingSource, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("funding_sources/%s", funding_source_token)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List all the funding sources associated with the Lithic account.
func (r *FundingSourceService) List(ctx context.Context, query *requests.FundingSourceListParams, opts ...options.RequestOption) (res *responses.FundingSourcesPage, err error) {
	opts = append(r.Options, opts...)
	path := "funding_sources"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.FundingSourcesPage{
		Page: &pagination.Page[responses.FundingSource]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Verify a bank account as a funding source by providing received micro-deposit
// amounts.
func (r *FundingSourceService) Verify(ctx context.Context, funding_source_token string, body *requests.FundingSourceVerifyParams, opts ...options.RequestOption) (res *responses.FundingSource, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("funding_sources/%s/verify", funding_source_token)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
