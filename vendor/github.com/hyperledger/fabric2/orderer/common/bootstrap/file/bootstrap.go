/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package file

import (
	"fmt"
	"io/ioutil"

	"github.com/hyperledger/fabric/orderer/common/bootstrap"
	cb "github.com/hyperledger/fabric/protos/common"

	"github.com/golang/protobuf/proto"
)

type fileBootstrapper struct {
	GenesisBlockFile string
}

// New returns a new static bootstrap helper
func New(fileName string) bootstrap.Helper {
	return &fileBootstrapper{
		GenesisBlockFile: fileName,
	}
}

// GenesisBlock returns the genesis block to be used for bootstrapping
func (b *fileBootstrapper) GenesisBlock() *cb.Block {
	bootstrapFile, fileErr := ioutil.ReadFile(b.GenesisBlockFile)
	if fileErr != nil {
		panic(fmt.Errorf("Unable to bootstrap orderer. Error reading genesis block file: %v", fileErr))
	}
	genesisBlock := &cb.Block{}
	unmarshallErr := proto.Unmarshal(bootstrapFile, genesisBlock)
	if unmarshallErr != nil {
		panic(fmt.Errorf("Unable to bootstrap orderer. Error unmarshalling genesis block: %v", unmarshallErr))

	}
	return genesisBlock
} // GenesisBlock
