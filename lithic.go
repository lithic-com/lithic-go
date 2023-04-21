package lithic

import (
	"context"
	"os"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/responses"
	"github.com/lithic-com/lithic-go/services"
)

type Lithic struct {
	Options                 []option.RequestOption
	Accounts                *services.AccountService
	AccountHolders          *services.AccountHolderService
	AuthRules               *services.AuthRuleService
	AuthStreamEnrollment    *services.AuthStreamEnrollmentService
	TokenizationDecisioning *services.TokenizationDecisioningService
	Cards                   *services.CardService
	Balances                *services.BalanceService
	AggregateBalances       *services.AggregateBalanceService
	Disputes                *services.DisputeService
	Events                  *services.EventService
	Transfers               *services.TransferService
	FinancialAccounts       *services.FinancialAccountService
	Transactions            *services.TransactionService
	ResponderEndpoints      *services.ResponderEndpointService
	Webhooks                *services.WebhookService
}

// NewLithic generates a new client with the default option read from the
// environment ("LITHIC_API_KEY", "LITHIC_WEBHOOK_SECRET"). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewLithic(opts ...option.RequestOption) (r *Lithic) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("LITHIC_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("LITHIC_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	opts = append(defaults, opts...)

	r = &Lithic{Options: opts}

	r.Accounts = services.NewAccountService(opts...)
	r.AccountHolders = services.NewAccountHolderService(opts...)
	r.AuthRules = services.NewAuthRuleService(opts...)
	r.AuthStreamEnrollment = services.NewAuthStreamEnrollmentService(opts...)
	r.TokenizationDecisioning = services.NewTokenizationDecisioningService(opts...)
	r.Cards = services.NewCardService(opts...)
	r.Balances = services.NewBalanceService(opts...)
	r.AggregateBalances = services.NewAggregateBalanceService(opts...)
	r.Disputes = services.NewDisputeService(opts...)
	r.Events = services.NewEventService(opts...)
	r.Transfers = services.NewTransferService(opts...)
	r.FinancialAccounts = services.NewFinancialAccountService(opts...)
	r.Transactions = services.NewTransactionService(opts...)
	r.ResponderEndpoints = services.NewResponderEndpointService(opts...)
	r.Webhooks = services.NewWebhookService(opts...)

	return
}

// API status check
func (r *Lithic) APIStatus(ctx context.Context, opts ...option.RequestOption) (res *responses.APIStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := "status"
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
