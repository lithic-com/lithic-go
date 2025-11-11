// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsAuthRuleConditionValueUnionParam()                             {}
func (UnionString) ImplementsAuthRuleConditionValueUnion()                                  {}
func (UnionString) ImplementsConditional3DsActionParametersConditionsValueUnion()           {}
func (UnionString) ImplementsConditionalAuthorizationActionParametersConditionsValueUnion() {}
func (UnionString) ImplementsCardProvisionResponseProvisioningPayloadUnion()                {}

type UnionInt int64

func (UnionInt) ImplementsAuthRuleConditionValueUnionParam()                             {}
func (UnionInt) ImplementsAuthRuleConditionValueUnion()                                  {}
func (UnionInt) ImplementsConditional3DsActionParametersConditionsValueUnion()           {}
func (UnionInt) ImplementsConditionalAuthorizationActionParametersConditionsValueUnion() {}
