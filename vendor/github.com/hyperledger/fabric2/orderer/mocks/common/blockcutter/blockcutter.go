/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockcutter

import (
	"github.com/hyperledger/fabric/common/flogging"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/op/go-logging"
)

const pkgLogID = "orderer/mocks/common/blockcutter"

var logger *logging.Logger

func init() {
	logger = flogging.MustGetLogger(pkgLogID)
}

// Receiver mocks the blockcutter.Receiver interface
type Receiver struct {
	// IsolatedTx causes Ordered returns [][]{curBatch, []{newTx}}, false when set to true
	IsolatedTx bool

	// CutAncestors causes Ordered returns [][]{curBatch}, true when set to true
	CutAncestors bool

	// CutNext causes Ordered returns [][]{append(curBatch, newTx)}, false when set to true
	CutNext bool

	// CurBatch is the currently outstanding messages in the batch
	CurBatch []*cb.Envelope

	// Block is a channel which is read from before returning from Ordered, it is useful for synchronization
	// If you do not wish synchronization for whatever reason, simply close the channel
	Block chan struct{}
}

// NewReceiver returns the mock blockcutter.Receiver implementation
func NewReceiver() *Receiver {
	return &Receiver{
		IsolatedTx:   false,
		CutAncestors: false,
		CutNext:      false,
		Block:        make(chan struct{}),
	}
}

// Ordered will add or cut the batch according to the state of Receiver, it blocks reading from Block on return
func (mbc *Receiver) Ordered(env *cb.Envelope) ([][]*cb.Envelope, bool) {
	defer func() {
		<-mbc.Block
	}()

	if mbc.IsolatedTx {
		logger.Debugf("Receiver: Returning dual batch")
		res := [][]*cb.Envelope{mbc.CurBatch, {env}}
		mbc.CurBatch = nil
		return res, false
	}

	if mbc.CutAncestors {
		logger.Debugf("Receiver: Returning current batch and appending newest env")
		res := [][]*cb.Envelope{mbc.CurBatch}
		mbc.CurBatch = []*cb.Envelope{env}
		return res, true
	}

	mbc.CurBatch = append(mbc.CurBatch, env)

	if mbc.CutNext {
		logger.Debugf("Receiver: Returning regular batch")
		res := [][]*cb.Envelope{mbc.CurBatch}
		mbc.CurBatch = nil
		return res, false
	}

	logger.Debugf("Appending to batch")
	return nil, true
}

// Cut terminates the current batch, returning it
func (mbc *Receiver) Cut() []*cb.Envelope {
	logger.Debugf("Cutting batch")
	res := mbc.CurBatch
	mbc.CurBatch = nil
	return res
}
