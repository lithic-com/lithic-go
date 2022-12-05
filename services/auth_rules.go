package services

import (
	"context"
	"fmt"

	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/pagination"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type AuthRuleService struct {
	Options []options.RequestOption
}

func NewAuthRuleService(opts ...options.RequestOption) (r *AuthRuleService) {
	r = &AuthRuleService{}
	r.Options = opts
	return
}

// Creates an authorization rule (Auth Rule) and applies it at the program,
// account, or card level.
func (r *AuthRuleService) New(ctx context.Context, body *requests.AuthRuleRequest, opts ...options.RequestOption) (res *responses.AuthRuleCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth_rules"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Detail the properties and entities (program, accounts, and cards) associated
// with an existing authorization rule (Auth Rule).
func (r *AuthRuleService) Get(ctx context.Context, auth_rule_token string, opts ...options.RequestOption) (res *responses.AuthRuleRetrieveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s", auth_rule_token)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update the properties associated with an existing authorization rule (Auth
// Rule).
func (r *AuthRuleService) Update(ctx context.Context, auth_rule_token string, body *requests.AuthRuleUpdateParams, opts ...options.RequestOption) (res *responses.AuthRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s", auth_rule_token)
	err = options.ExecuteNewRequest(ctx, "PUT", path, body, &res, opts...)
	return
}

// Return all of the Auth Rules under the program.
func (r *AuthRuleService) List(ctx context.Context, query *requests.AuthRuleListParams, opts ...options.RequestOption) (res *responses.AuthRulesPage, err error) {
	opts = append(r.Options, opts...)
	path := "auth_rules"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.AuthRulesPage{
		Page: &pagination.Page[responses.AuthRule]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Applies an existing authorization rule (Auth Rule) to an program, account, or
// card level.
func (r *AuthRuleService) Apply(ctx context.Context, auth_rule_token string, body *requests.AuthRuleApplyParams, opts ...options.RequestOption) (res *responses.AuthRuleApplyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s/apply", auth_rule_token)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Remove an existing authorization rule (Auth Rule) from an program, account, or
// card-level.
func (r *AuthRuleService) Remove(ctx context.Context, body *requests.AuthRuleRemoveParams, opts ...options.RequestOption) (res *responses.AuthRuleRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth_rules/remove"
	err = options.ExecuteNewRequest(ctx, "DELETE", path, body, &res, opts...)
	return
}
