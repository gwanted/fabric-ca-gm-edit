/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cauthdsl

import (
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"
)

// TemplatePolicy creates a headerless configuration item representing a policy for a given key
func TemplatePolicy(key string, sigPolicyEnv *cb.SignaturePolicyEnvelope) *cb.ConfigGroup {
	configGroup := cb.NewConfigGroup()
	configGroup.Policies[key] = &cb.ConfigPolicy{
		Policy: &cb.Policy{
			Type:  int32(cb.Policy_SIGNATURE),
			Value: utils.MarshalOrPanic(sigPolicyEnv),
		},
	}
	return configGroup
}
