/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bootstrap

import (
	ab "github.com/hyperledger/fabric/protos/common"
)

// Helper defines the functions a bootstrapping implementation should provide.
type Helper interface {
	// GenesisBlock should return the genesis block required to bootstrap
	// the ledger (be it reading from the filesystem, generating it, etc.)
	GenesisBlock() *ab.Block
}
