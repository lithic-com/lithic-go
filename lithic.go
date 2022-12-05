package lithic

import (
	"context"
	"os"

	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/responses"
	"github.com/lithic-com/lithic-go/services"
)

type Lithic struct {
	Options              []options.RequestOption
	Accounts             *services.AccountService
	AccountHolders       *services.AccountHolderService
	AuthRules            *services.AuthRuleService
	AuthStreamEnrollment *services.AuthStreamEnrollmentService
	Cards                *services.CardService
	Disputes             *services.DisputeService
	Events               *services.EventService
	FundingSources       *services.FundingSourceService
	Transactions         *services.TransactionService
	Webhooks             *services.WebhookService
}

// NewLithic generates a new client with the default options read from the
// environment ("LITHIC_API_KEY", "LITHIC_WEBHOOK_SECRET"). The options passed in
// as arguments are applied after these default arguments, and all options will be
// passed down to the services and requests that this client makes.
func NewLithic(opts ...options.RequestOption) (r *Lithic) {
	defaults := []options.RequestOption{options.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("LITHIC_API_KEY"); ok {
		defaults = append(defaults, options.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("LITHIC_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, options.WithWebhookSecret(o))
	}
	opts = append(defaults, opts...)

	r = &Lithic{Options: opts}

	r.Accounts = services.NewAccountService(opts...)
	r.AccountHolders = services.NewAccountHolderService(opts...)
	r.AuthRules = services.NewAuthRuleService(opts...)
	r.AuthStreamEnrollment = services.NewAuthStreamEnrollmentService(opts...)
	r.Cards = services.NewCardService(opts...)
	r.Disputes = services.NewDisputeService(opts...)
	r.Events = services.NewEventService(opts...)
	r.FundingSources = services.NewFundingSourceService(opts...)
	r.Transactions = services.NewTransactionService(opts...)
	r.Webhooks = services.NewWebhookService(opts...)

	return
}

// API status check
func (r *Lithic) APIStatus(ctx context.Context, opts ...options.RequestOption) (res *responses.APIStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := "status"
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
