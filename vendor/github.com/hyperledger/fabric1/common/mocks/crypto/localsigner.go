/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package crypto

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

// FakeLocalSigner is a signer which already has identity an nonce set to fake values
var FakeLocalSigner = &LocalSigner{
	Identity: []byte("IdentityBytes"),
	Nonce:    []byte("NonceValue"),
}

// LocalSigner is a mock implmeentation of crypto.LocalSigner
type LocalSigner struct {
	Identity []byte
	Nonce    []byte
}

// Sign returns the msg, nil
func (ls *LocalSigner) Sign(msg []byte) ([]byte, error) {
	return msg, nil
}

// NewSignatureHeader returns a new signature header, nil
func (ls *LocalSigner) NewSignatureHeader() (*cb.SignatureHeader, error) {
	return &cb.SignatureHeader{
		Creator: ls.Identity,
		Nonce:   ls.Nonce,
	}, nil
}
