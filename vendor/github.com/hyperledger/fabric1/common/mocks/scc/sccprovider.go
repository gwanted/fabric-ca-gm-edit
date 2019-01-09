/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package scc

import (
	lm "github.com/hyperledger/fabric/common/mocks/ledger"
	"github.com/hyperledger/fabric/core/common/sysccprovider"
	"github.com/hyperledger/fabric/core/ledger"
)

type MocksccProviderFactory struct {
	Qe   *lm.MockQueryExecutor
	QErr error
}

func (c *MocksccProviderFactory) NewSystemChaincodeProvider() sysccprovider.SystemChaincodeProvider {
	return &mocksccProviderImpl{Qe: c.Qe, QErr: c.QErr}
}

type mocksccProviderImpl struct {
	Qe   *lm.MockQueryExecutor
	QErr error
}

func (c *mocksccProviderImpl) IsSysCC(name string) bool {
	return (name == "lscc") || (name == "escc") || (name == "vscc") || (name == "notext")
}

func (c *mocksccProviderImpl) IsSysCCAndNotInvokableCC2CC(name string) bool {
	return (name == "escc") || (name == "vscc")
}

func (c *mocksccProviderImpl) IsSysCCAndNotInvokableExternal(name string) bool {
	return (name == "escc") || (name == "vscc") || (name == "notext")
}

func (c *mocksccProviderImpl) GetQueryExecutorForLedger(cid string) (ledger.QueryExecutor, error) {
	return c.Qe, c.QErr
}
