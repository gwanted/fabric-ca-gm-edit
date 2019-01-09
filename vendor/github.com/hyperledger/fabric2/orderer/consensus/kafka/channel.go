/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kafka

import "fmt"

const defaultPartition = 0

// channel identifies the Kafka partition the Kafka-based orderer interacts
// with.
type channel interface {
	topic() string
	partition() int32
	fmt.Stringer
}

type channelImpl struct {
	tpc string
	prt int32
}

// Returns a new channel for a given topic name and partition number.
func newChannel(topic string, partition int32) channel {
	return &channelImpl{
		tpc: fmt.Sprintf("%s", topic),
		prt: partition,
	}
}

// topic returns the Kafka topic this channel belongs to.
func (chn *channelImpl) topic() string {
	return chn.tpc
}

// partition returns the Kafka partition where this channel resides.
func (chn *channelImpl) partition() int32 {
	return chn.prt
}

// String returns a string identifying the Kafka topic/partition corresponding
// to this channel.
func (chn *channelImpl) String() string {
	return fmt.Sprintf("%s/%d", chn.tpc, chn.prt)
}
