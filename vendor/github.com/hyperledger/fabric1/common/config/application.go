/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"

	"github.com/hyperledger/fabric/common/config/msp"
)

const (
	// ApplicationGroupKey is the group name for the Application config
	ApplicationGroupKey = "Application"
)

// ApplicationGroup represents the application config group
type ApplicationGroup struct {
	*Proposer
	*ApplicationConfig
	mspConfig *msp.MSPConfigHandler
}

type ApplicationConfig struct {
	*standardValues

	applicationGroup *ApplicationGroup
	applicationOrgs  map[string]ApplicationOrg
}

// NewSharedConfigImpl creates a new SharedConfigImpl with the given CryptoHelper
func NewApplicationGroup(mspConfig *msp.MSPConfigHandler) *ApplicationGroup {
	ag := &ApplicationGroup{
		mspConfig: mspConfig,
	}
	ag.Proposer = NewProposer(ag)

	return ag
}

func (ag *ApplicationGroup) NewGroup(name string) (ValueProposer, error) {
	return NewApplicationOrgGroup(name, ag.mspConfig), nil
}

// Allocate returns a new instance of the ApplicationConfig
func (ag *ApplicationGroup) Allocate() Values {
	return NewApplicationConfig(ag)
}

func NewApplicationConfig(ag *ApplicationGroup) *ApplicationConfig {
	sv, err := NewStandardValues(&(struct{}{}))
	if err != nil {
		logger.Panicf("Programming error: %s", err)
	}

	return &ApplicationConfig{
		applicationGroup: ag,

		// Currently there are no config values
		standardValues: sv,
	}
}

func (ac *ApplicationConfig) Validate(tx interface{}, groups map[string]ValueProposer) error {
	ac.applicationOrgs = make(map[string]ApplicationOrg)
	var ok bool
	for key, value := range groups {
		ac.applicationOrgs[key], ok = value.(*ApplicationOrgGroup)
		if !ok {
			return fmt.Errorf("Application sub-group %s was not an ApplicationOrgGroup, actually %T", key, value)
		}
	}
	return nil
}

func (ac *ApplicationConfig) Commit() {
	ac.applicationGroup.ApplicationConfig = ac
}

// Organizations returns a map of org ID to ApplicationOrg
func (ac *ApplicationConfig) Organizations() map[string]ApplicationOrg {
	return ac.applicationOrgs
}
