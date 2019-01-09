/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func (mp *MSPPrincipal) VariablyOpaqueFields() []string {
	return []string{"principal"}
}

func (mp *MSPPrincipal) VariablyOpaqueFieldProto(name string) (proto.Message, error) {
	if name != mp.VariablyOpaqueFields()[0] {
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
	switch mp.PrincipalClassification {
	case MSPPrincipal_ROLE:
		return &MSPRole{}, nil
	case MSPPrincipal_ORGANIZATION_UNIT:
		return &OrganizationUnit{}, nil
	case MSPPrincipal_IDENTITY:
		return nil, fmt.Errorf("unable to decode MSP type IDENTITY until the protos are fixed to include the IDENTITY proto in protos/msp")
	default:
		return nil, fmt.Errorf("unable to decode MSP type: %v", mp.PrincipalClassification)
	}
}
