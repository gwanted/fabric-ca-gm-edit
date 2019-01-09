/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func (mc *MSPConfig) VariablyOpaqueFields() []string {
	return []string{"config"}
}

func (mc *MSPConfig) VariablyOpaqueFieldProto(name string) (proto.Message, error) {
	if name != mc.VariablyOpaqueFields()[0] {
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
	switch mc.Type {
	case 0:
		return &FabricMSPConfig{}, nil
	default:
		return nil, fmt.Errorf("unable to decode MSP type: %v", mc.Type)
	}
}
