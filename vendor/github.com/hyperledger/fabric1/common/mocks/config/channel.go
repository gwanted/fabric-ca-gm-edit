/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import "github.com/hyperledger/fabric/common/util"

func nearIdentityHash(input []byte) []byte {
	return util.ConcatenateBytes([]byte("FakeHash("), input, []byte(""))
}

// Channel is a mock implementation of config.Channel
type Channel struct {
	// HashingAlgorithmVal is returned as the result of HashingAlgorithm() if set
	HashingAlgorithmVal func([]byte) []byte
	// BlockDataHashingStructureWidthVal is returned as the result of BlockDataHashingStructureWidth()
	BlockDataHashingStructureWidthVal uint32
	// OrdererAddressesVal is returned as the result of OrdererAddresses()
	OrdererAddressesVal []string
}

// HashingAlgorithm returns the HashingAlgorithmVal if set, otherwise a fake simple hash function
func (scm *Channel) HashingAlgorithm() func([]byte) []byte {
	if scm.HashingAlgorithmVal == nil {
		return nearIdentityHash
	}
	return scm.HashingAlgorithmVal
}

// BlockDataHashingStructureWidth returns the BlockDataHashingStructureWidthVal
func (scm *Channel) BlockDataHashingStructureWidth() uint32 {
	return scm.BlockDataHashingStructureWidthVal
}

// OrdererAddresses returns the OrdererAddressesVal
func (scm *Channel) OrdererAddresses() []string {
	return scm.OrdererAddressesVal
}
