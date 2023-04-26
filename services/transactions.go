package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type TransactionService struct {
	Options []option.RequestOption
}

func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	return
}

// Get specific transaction.
func (r *TransactionService) Get(ctx context.Context, transaction_token string, opts ...option.RequestOption) (res *responses.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("transactions/%s", transaction_token)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List transactions.
func (r *TransactionService) List(ctx context.Context, query *requests.TransactionListParams, opts ...option.RequestOption) (res *responses.Page[responses.Transaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "transactions"
	cfg, err := option.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List transactions.
func (r *TransactionService) ListAutoPager(ctx context.Context, query *requests.TransactionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Transaction] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Simulates an authorization request from the payment network as if it came from a
// merchant acquirer. If you're configured for ASA, simulating auths requires your
// ASA client to be set up properly (respond with a valid JSON to the ASA request).
// For users that are not configured for ASA, a daily transaction limit of $5000
// USD is applied by default. This limit can be modified via the
// [update account](https://docs.lithic.com/reference/patchaccountbytoken)
// endpoint.
func (r *TransactionService) SimulateAuthorization(ctx context.Context, body *requests.TransactionSimulateAuthorizationParams, opts ...option.RequestOption) (res *responses.TransactionSimulateAuthorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/authorize"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates an authorization advice request from the payment network as if it came
// from a merchant acquirer. An authorization advice request changes the amount of
// the transaction.
func (r *TransactionService) SimulateAuthorizationAdvice(ctx context.Context, body *requests.TransactionSimulateAuthorizationAdviceParams, opts ...option.RequestOption) (res *responses.TransactionSimulateAuthorizationAdviceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/authorization_advice"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Clears an existing authorization. After this event, the transaction is no longer
// pending.
//
// If no `amount` is supplied to this endpoint, the amount of the transaction will
// be captured. Any transaction that has any amount completed at all do not have
// access to this behavior.
func (r *TransactionService) SimulateClearing(ctx context.Context, body *requests.TransactionSimulateClearingParams, opts ...option.RequestOption) (res *responses.TransactionSimulateClearingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/clearing"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates a credit authorization advice message from the payment network. This
// message indicates that a credit authorization was approved on your behalf by the
// network.
func (r *TransactionService) SimulateCreditAuthorization(ctx context.Context, body *requests.TransactionSimulateCreditAuthorizationParams, opts ...option.RequestOption) (res *responses.TransactionSimulateCreditAuthorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/credit_authorization_advice"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns (aka refunds) an amount back to a card. Returns are cleared immediately
// and do not spend time in a `PENDING` state.
func (r *TransactionService) SimulateReturn(ctx context.Context, body *requests.TransactionSimulateReturnParams, opts ...option.RequestOption) (res *responses.TransactionSimulateReturnResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/return"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Voids a settled credit transaction – i.e., a transaction with a negative amount
// and `SETTLED` status. These can be credit authorizations that have already
// cleared or financial credit authorizations.
func (r *TransactionService) SimulateReturnReversal(ctx context.Context, body *requests.TransactionSimulateReturnReversalParams, opts ...option.RequestOption) (res *responses.TransactionSimulateReturnReversalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/return_reversal"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Voids an existing, uncleared (aka pending) authorization. If amount is not sent
// the full amount will be voided. Cannot be used on partially completed
// transactions, but can be used on partially voided transactions. _Note that
// simulating an authorization expiry on credit authorizations or credit
// authorization advice is not currently supported but will be added soon._
func (r *TransactionService) SimulateVoid(ctx context.Context, body *requests.TransactionSimulateVoidParams, opts ...option.RequestOption) (res *responses.TransactionSimulateVoidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulate/void"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}
