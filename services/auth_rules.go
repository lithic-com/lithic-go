package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type AuthRuleService struct {
	Options []option.RequestOption
}

func NewAuthRuleService(opts ...option.RequestOption) (r *AuthRuleService) {
	r = &AuthRuleService{}
	r.Options = opts
	return
}

// Detail the properties and entities (program, accounts, and cards) associated
// with an existing authorization rule (Auth Rule).
func (r *AuthRuleService) Get(ctx context.Context, auth_rule_token string, opts ...option.RequestOption) (res *responses.AuthRuleRetrieveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s", auth_rule_token)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update the properties associated with an existing authorization rule (Auth
// Rule).
func (r *AuthRuleService) Update(ctx context.Context, auth_rule_token string, body *requests.AuthRuleUpdateParams, opts ...option.RequestOption) (res *responses.AuthRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s", auth_rule_token)
	err = option.ExecuteNewRequest(ctx, "PUT", path, body, &res, opts...)
	return
}

// Return all of the Auth Rules under the program.
func (r *AuthRuleService) List(ctx context.Context, query *requests.AuthRuleListParams, opts ...option.RequestOption) (res *responses.Page[responses.AuthRule], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "auth_rules"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
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

// Return all of the Auth Rules under the program.
func (r *AuthRuleService) ListAutoPager(ctx context.Context, query *requests.AuthRuleListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.AuthRule] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Applies an existing authorization rule (Auth Rule) to an program, account, or
// card level.
func (r *AuthRuleService) Apply(ctx context.Context, auth_rule_token string, body *requests.AuthRuleApplyParams, opts ...option.RequestOption) (res *responses.AuthRuleApplyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("auth_rules/%s/apply", auth_rule_token)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Remove an existing authorization rule (Auth Rule) from an program, account, or
// card-level.
func (r *AuthRuleService) Remove(ctx context.Context, body *requests.AuthRuleRemoveParams, opts ...option.RequestOption) (res *responses.AuthRuleRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth_rules/remove"
	err = option.ExecuteNewRequest(ctx, "DELETE", path, body, &res, opts...)
	return
}
