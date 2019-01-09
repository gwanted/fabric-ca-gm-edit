/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package scc

import (
	"github.com/hyperledger/fabric/common/channelconfig"
	lm "github.com/hyperledger/fabric/common/mocks/ledger"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/core/common/sysccprovider"
	"github.com/hyperledger/fabric/core/ledger"
)

type MocksccProviderFactory struct {
	Qe                    *lm.MockQueryExecutor
	QErr                  error
	ApplicationConfigRv   channelconfig.Application
	ApplicationConfigBool bool
	PolicyManagerRv       policies.Manager
	PolicyManagerBool     bool
}

func (c *MocksccProviderFactory) NewSystemChaincodeProvider() sysccprovider.SystemChaincodeProvider {
	return &MocksccProviderImpl{
		Qe:                    c.Qe,
		QErr:                  c.QErr,
		ApplicationConfigRv:   c.ApplicationConfigRv,
		ApplicationConfigBool: c.ApplicationConfigBool,
		PolicyManagerBool:     c.PolicyManagerBool,
		PolicyManagerRv:       c.PolicyManagerRv,
	}
}

type MocksccProviderImpl struct {
	Qe                    *lm.MockQueryExecutor
	QErr                  error
	ApplicationConfigRv   channelconfig.Application
	ApplicationConfigBool bool
	PolicyManagerRv       policies.Manager
	PolicyManagerBool     bool
	SysCCMap              map[string]bool
}

func (c *MocksccProviderImpl) IsSysCC(name string) bool {
	if c.SysCCMap != nil {
		return c.SysCCMap[name]
	}

	return (name == "lscc") || (name == "escc") || (name == "vscc") || (name == "notext")
}

func (c *MocksccProviderImpl) IsSysCCAndNotInvokableCC2CC(name string) bool {
	return (name == "escc") || (name == "vscc")
}

func (c *MocksccProviderImpl) IsSysCCAndNotInvokableExternal(name string) bool {
	return (name == "escc") || (name == "vscc") || (name == "notext")
}

func (c *MocksccProviderImpl) GetQueryExecutorForLedger(cid string) (ledger.QueryExecutor, error) {
	return c.Qe, c.QErr
}

func (c *MocksccProviderImpl) GetApplicationConfig(cid string) (channelconfig.Application, bool) {
	return c.ApplicationConfigRv, c.ApplicationConfigBool
}

func (c *MocksccProviderImpl) PolicyManager(channelID string) (policies.Manager, bool) {
	return c.PolicyManagerRv, c.PolicyManagerBool
}
