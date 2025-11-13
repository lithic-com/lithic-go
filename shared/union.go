// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsConditionalValueUnionParam()                    {}
func (UnionString) ImplementsConditionalValueUnion()                         {}
func (UnionString) ImplementsCardProvisionResponseProvisioningPayloadUnion() {}

type UnionInt int64

func (UnionInt) ImplementsConditionalValueUnionParam() {}
func (UnionInt) ImplementsConditionalValueUnion()      {}
