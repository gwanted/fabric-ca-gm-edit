/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package crypto

import cb "github.com/hyperledger/fabric/protos/common"

// LocalSigner is a temporary stub interface which will be implemented by the local MSP
type LocalSigner interface {
	// NewSignatureHeader creates a SignatureHeader with the correct signing identity and a valid nonce
	NewSignatureHeader() (*cb.SignatureHeader, error)

	// Sign a message which should embed a signature header created by NewSignatureHeader
	Sign(message []byte) ([]byte, error)
}
