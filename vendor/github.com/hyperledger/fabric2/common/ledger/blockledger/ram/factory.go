/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ramledger

import (
	"sync"

	"github.com/hyperledger/fabric/common/ledger/blockledger"
	cb "github.com/hyperledger/fabric/protos/common"
)

type ramLedgerFactory struct {
	maxSize int
	ledgers map[string]blockledger.ReadWriter
	mutex   sync.Mutex
}

// GetOrCreate gets an existing ledger (if it exists) or creates it if it does not
func (rlf *ramLedgerFactory) GetOrCreate(chainID string) (blockledger.ReadWriter, error) {
	rlf.mutex.Lock()
	defer rlf.mutex.Unlock()

	key := chainID

	l, ok := rlf.ledgers[key]
	if ok {
		return l, nil
	}

	ch := newChain(rlf.maxSize)
	rlf.ledgers[key] = ch
	return ch, nil
}

// newChain creates a new chain backed by a RAM ledger
func newChain(maxSize int) blockledger.ReadWriter {
	preGenesis := &cb.Block{
		Header: &cb.BlockHeader{
			Number: ^uint64(0),
		},
	}

	rl := &ramLedger{
		maxSize: maxSize,
		size:    1,
		oldest: &simpleList{
			signal: make(chan struct{}),
			block:  preGenesis,
		},
	}
	rl.newest = rl.oldest
	return rl
}

// ChainIDs returns the chain IDs the factory is aware of
func (rlf *ramLedgerFactory) ChainIDs() []string {
	rlf.mutex.Lock()
	defer rlf.mutex.Unlock()
	ids := make([]string, len(rlf.ledgers))

	i := 0
	for key := range rlf.ledgers {
		ids[i] = key
		i++
	}

	return ids
}

// Close is a no-op for the RAM ledger
func (rlf *ramLedgerFactory) Close() {
	return // nothing to do
}

// New creates a new ledger factory
func New(maxSize int) blockledger.Factory {
	rlf := &ramLedgerFactory{
		maxSize: maxSize,
		ledgers: make(map[string]blockledger.ReadWriter),
	}

	return rlf
}
