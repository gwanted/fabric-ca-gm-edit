/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package common

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func (p *Policy) VariablyOpaqueFields() []string {
	return []string{"value"}
}

func (p *Policy) VariablyOpaqueFieldProto(name string) (proto.Message, error) {
	if name != p.VariablyOpaqueFields()[0] {
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
	switch p.Type {
	case int32(Policy_SIGNATURE):
		return &SignaturePolicyEnvelope{}, nil
	case int32(Policy_IMPLICIT_META):
		return &ImplicitMetaPolicy{}, nil
	default:
		return nil, fmt.Errorf("unable to decode policy type: %v", p.Type)
	}
}
