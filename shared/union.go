// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsAuthRuleConditionValueUnionParam() {}
func (UnionString) ImplementsAuthRuleConditionValueUnion()      {}

type UnionInt int64

func (UnionInt) ImplementsAuthRuleConditionValueUnionParam() {}
func (UnionInt) ImplementsAuthRuleConditionValueUnion()      {}
func (UnionInt) ImplementsVelocityLimitParamsPeriodUnion()   {}
