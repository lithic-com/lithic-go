// File generated from our OpenAPI spec by Stainless.

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
	CardProduct             *CardProductService
	CardPrograms            *CardProgramService
	DigitalCardArt          *DigitalCardArtService
}

// NewClient generates a new client with the default option read from the
// environment (LITHIC_API_KEY, LITHIC_WEBHOOK_SECRET). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("LITHIC_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("LITHIC_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	opts = append(defaults, opts...)

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
	r.CardProduct = NewCardProductService(opts...)
	r.CardPrograms = NewCardProgramService(opts...)
	r.DigitalCardArt = NewDigitalCardArtService(opts...)

	return
}

// Status of api
func (r *Client) APIStatus(ctx context.Context, opts ...option.RequestOption) (res *APIStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := "status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
