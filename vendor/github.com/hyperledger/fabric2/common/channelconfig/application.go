/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package channelconfig

import (
	"github.com/hyperledger/fabric/common/capabilities"
	cb "github.com/hyperledger/fabric/protos/common"

	"github.com/pkg/errors"
)

const (
	// ApplicationGroupKey is the group name for the Application config
	ApplicationGroupKey = "Application"
)

// ApplicationProtos is used as the source of the ApplicationConfig
type ApplicationProtos struct {
	Capabilities *cb.Capabilities
}

// ApplicationConfig implements the Application interface
type ApplicationConfig struct {
	applicationOrgs map[string]ApplicationOrg
	protos          *ApplicationProtos
}

// NewApplicationConfig creates config from an Application config group
func NewApplicationConfig(appGroup *cb.ConfigGroup, mspConfig *MSPConfigHandler) (*ApplicationConfig, error) {
	ac := &ApplicationConfig{
		applicationOrgs: make(map[string]ApplicationOrg),
		protos:          &ApplicationProtos{},
	}

	if err := DeserializeProtoValuesFromGroup(appGroup, ac.protos); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize values")
	}

	var err error
	for orgName, orgGroup := range appGroup.Groups {
		ac.applicationOrgs[orgName], err = NewApplicationOrgConfig(orgName, orgGroup, mspConfig)
		if err != nil {
			return nil, err
		}
	}

	return ac, nil
}

// Organizations returns a map of org ID to ApplicationOrg
func (ac *ApplicationConfig) Organizations() map[string]ApplicationOrg {
	return ac.applicationOrgs
}

// Capabilities returns a map of capability name to Capability
func (ac *ApplicationConfig) Capabilities() ApplicationCapabilities {
	return capabilities.NewApplicationProvider(ac.protos.Capabilities.Capabilities)
}
