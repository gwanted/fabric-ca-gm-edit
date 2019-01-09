/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"

	mspconfig "github.com/hyperledger/fabric/common/config/msp"
	"github.com/hyperledger/fabric/msp"
	mspprotos "github.com/hyperledger/fabric/protos/msp"
)

// Org config keys
const (
	// MSPKey is value key for marshaled *mspconfig.MSPConfig
	MSPKey = "MSP"
)

type OrganizationProtos struct {
	MSP *mspprotos.MSPConfig
}

type OrganizationConfig struct {
	*standardValues
	protos *OrganizationProtos

	organizationGroup *OrganizationGroup

	msp   msp.MSP
	mspID string
}

// Config stores common configuration information for organizations
type OrganizationGroup struct {
	*Proposer
	*OrganizationConfig
	name             string
	mspConfigHandler *mspconfig.MSPConfigHandler
}

// NewConfig creates an instnace of the organization Config
func NewOrganizationGroup(name string, mspConfigHandler *mspconfig.MSPConfigHandler) *OrganizationGroup {
	og := &OrganizationGroup{
		name:             name,
		mspConfigHandler: mspConfigHandler,
	}
	og.Proposer = NewProposer(og)
	return og
}

// Name returns the name this org is referred to in config
func (og *OrganizationGroup) Name() string {
	return og.name
}

// MSPID returns the MSP ID associated with this org
func (og *OrganizationGroup) MSPID() string {
	return og.mspID
}

// NewGroup always errors
func (og *OrganizationGroup) NewGroup(name string) (ValueProposer, error) {
	return nil, fmt.Errorf("Organization does not support subgroups")
}

// Allocate creates the proto resources neeeded for a proposal
func (og *OrganizationGroup) Allocate() Values {
	return NewOrganizationConfig(og)
}

func NewOrganizationConfig(og *OrganizationGroup) *OrganizationConfig {
	oc := &OrganizationConfig{
		protos: &OrganizationProtos{},

		organizationGroup: og,
	}

	var err error
	oc.standardValues, err = NewStandardValues(oc.protos)
	if err != nil {
		logger.Panicf("Programming error: %s", err)
	}
	return oc
}

// Validate returns whether the configuration is valid
func (oc *OrganizationConfig) Validate(tx interface{}, groups map[string]ValueProposer) error {
	return oc.validateMSP(tx)
}

func (oc *OrganizationConfig) Commit() {
	oc.organizationGroup.OrganizationConfig = oc
}

func (oc *OrganizationConfig) validateMSP(tx interface{}) error {
	var err error

	logger.Debugf("Setting up MSP for org %s", oc.organizationGroup.name)
	oc.msp, err = oc.organizationGroup.mspConfigHandler.ProposeMSP(tx, oc.protos.MSP)
	if err != nil {
		return err
	}

	oc.mspID, _ = oc.msp.GetIdentifier()

	if oc.mspID == "" {
		return fmt.Errorf("MSP for org %s has empty MSP ID", oc.organizationGroup.name)
	}

	if oc.organizationGroup.OrganizationConfig != nil && oc.organizationGroup.mspID != oc.mspID {
		return fmt.Errorf("Organization %s attempted to change its MSP ID from %s to %s", oc.organizationGroup.name, oc.organizationGroup.mspID, oc.mspID)
	}

	return nil
}
