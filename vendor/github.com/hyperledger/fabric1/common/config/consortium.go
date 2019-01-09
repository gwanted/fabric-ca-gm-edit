/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"

	"github.com/hyperledger/fabric/common/config/msp"
	cb "github.com/hyperledger/fabric/protos/common"
)

// ConsortiumProtos holds the config protos for the consortium config
type ConsortiumProtos struct {
	ChannelCreationPolicy *cb.Policy
}

// ConsortiumGroup stores the set of Consortium
type ConsortiumGroup struct {
	*Proposer
	*ConsortiumConfig

	mspConfig *msp.MSPConfigHandler
}

// NewConsortiumGroup creates a new *ConsortiumGroup
func NewConsortiumGroup(mspConfig *msp.MSPConfigHandler) *ConsortiumGroup {
	cg := &ConsortiumGroup{
		mspConfig: mspConfig,
	}
	cg.Proposer = NewProposer(cg)
	return cg
}

// NewGroup returns a Consortium instance
func (cg *ConsortiumGroup) NewGroup(name string) (ValueProposer, error) {
	return NewOrganizationGroup(name, cg.mspConfig), nil
}

// Allocate returns the resources for a new config proposal
func (cg *ConsortiumGroup) Allocate() Values {
	return NewConsortiumConfig(cg)
}

// BeginValueProposals calls through to Proposer after calling into the MSP config Handler
func (cg *ConsortiumGroup) BeginValueProposals(tx interface{}, groups []string) (ValueDeserializer, []ValueProposer, error) {
	return cg.Proposer.BeginValueProposals(tx, groups)
}

// PreCommit intercepts the precommit request and commits the MSP config handler before calling the underlying proposer
func (cg *ConsortiumGroup) PreCommit(tx interface{}) error {
	return cg.Proposer.PreCommit(tx)
}

// RollbackProposals intercepts the rollback request and commits the MSP config handler before calling the underlying proposer
func (cg *ConsortiumGroup) RollbackProposals(tx interface{}) {
	cg.Proposer.RollbackProposals(tx)
}

// CommitProposals intercepts the commit request and commits the MSP config handler before calling the underlying proposer
func (cg *ConsortiumGroup) CommitProposals(tx interface{}) {
	cg.Proposer.CommitProposals(tx)
}

// ConsortiumConfig holds the consoritums configuration information
type ConsortiumConfig struct {
	*standardValues
	protos *ConsortiumProtos
	orgs   map[string]*OrganizationGroup

	consortiumGroup *ConsortiumGroup
}

// NewConsortiumConfig creates a new instance of the consoritums config
func NewConsortiumConfig(cg *ConsortiumGroup) *ConsortiumConfig {
	cc := &ConsortiumConfig{
		protos:          &ConsortiumProtos{},
		orgs:            make(map[string]*OrganizationGroup),
		consortiumGroup: cg,
	}
	var err error
	cc.standardValues, err = NewStandardValues(cc.protos)
	if err != nil {
		logger.Panicf("Programming error: %s", err)
	}
	return cc
}

// Organizations returns the set of organizations in the consortium
func (cc *ConsortiumConfig) Organizations() map[string]*OrganizationGroup {
	return cc.orgs
}

// CreationPolicy returns the policy structure used to validate
// the channel creation
func (cc *ConsortiumConfig) ChannelCreationPolicy() *cb.Policy {
	return cc.protos.ChannelCreationPolicy
}

// Commit commits the ConsortiumConfig
func (cc *ConsortiumConfig) Commit() {
	cc.consortiumGroup.ConsortiumConfig = cc
}

// Validate builds the Consortium map
func (cc *ConsortiumConfig) Validate(tx interface{}, groups map[string]ValueProposer) error {
	var ok bool
	for key, group := range groups {
		cc.orgs[key], ok = group.(*OrganizationGroup)
		if !ok {
			return fmt.Errorf("Unexpected group type: %T", group)
		}
	}
	return nil
}
