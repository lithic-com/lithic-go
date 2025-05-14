// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"context"
	"net/http"
	"os"

	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the lithic API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                 []option.RequestOption
	Accounts                *AccountService
	AccountHolders          *AccountHolderService
	AuthRules               *AuthRuleService
	AuthStreamEnrollment    *AuthStreamEnrollmentService
	TokenizationDecisioning *TokenizationDecisioningService
	Tokenizations           *TokenizationService
	Cards                   *CardService
	Balances                *BalanceService
	AggregateBalances       *AggregateBalanceService
	Disputes                *DisputeService
	Events                  *EventService
	Transfers               *TransferService
	FinancialAccounts       *FinancialAccountService
	Transactions            *TransactionService
	ResponderEndpoints      *ResponderEndpointService
	ExternalBankAccounts    *ExternalBankAccountService
	Payments                *PaymentService
	ThreeDS                 *ThreeDSService
	Reports                 *ReportService
	CardPrograms            *CardProgramService
	DigitalCardArt          *DigitalCardArtService
	BookTransfers           *BookTransferService
	CreditProducts          *CreditProductService
	ExternalPayments        *ExternalPaymentService
	ManagementOperations    *ManagementOperationService
	FundingEvents           *FundingEventService
}

// DefaultClientOptions read from the environment (LITHIC_API_KEY,
// LITHIC_WEBHOOK_SECRET, LITHIC_BASE_URL). This should be used to initialize new
// clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("LITHIC_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("LITHIC_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("LITHIC_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (LITHIC_API_KEY, LITHIC_WEBHOOK_SECRET, LITHIC_BASE_URL). The option
// passed in as arguments are applied after these default arguments, and all option
// will be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = &Client{Options: opts}

	r.Accounts = NewAccountService(opts...)
	r.AccountHolders = NewAccountHolderService(opts...)
	r.AuthRules = NewAuthRuleService(opts...)
	r.AuthStreamEnrollment = NewAuthStreamEnrollmentService(opts...)
	r.TokenizationDecisioning = NewTokenizationDecisioningService(opts...)
	r.Tokenizations = NewTokenizationService(opts...)
	r.Cards = NewCardService(opts...)
	r.Balances = NewBalanceService(opts...)
	r.AggregateBalances = NewAggregateBalanceService(opts...)
	r.Disputes = NewDisputeService(opts...)
	r.Events = NewEventService(opts...)
	r.Transfers = NewTransferService(opts...)
	r.FinancialAccounts = NewFinancialAccountService(opts...)
	r.Transactions = NewTransactionService(opts...)
	r.ResponderEndpoints = NewResponderEndpointService(opts...)
	r.ExternalBankAccounts = NewExternalBankAccountService(opts...)
	r.Payments = NewPaymentService(opts...)
	r.ThreeDS = NewThreeDSService(opts...)
	r.Reports = NewReportService(opts...)
	r.CardPrograms = NewCardProgramService(opts...)
	r.DigitalCardArt = NewDigitalCardArtService(opts...)
	r.BookTransfers = NewBookTransferService(opts...)
	r.CreditProducts = NewCreditProductService(opts...)
	r.ExternalPayments = NewExternalPaymentService(opts...)
	r.ManagementOperations = NewManagementOperationService(opts...)
	r.FundingEvents = NewFundingEventService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}

// Status of api
func (r *Client) APIStatus(ctx context.Context, opts ...option.RequestOption) (res *APIStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
