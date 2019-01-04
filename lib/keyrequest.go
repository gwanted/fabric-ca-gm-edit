// +build !nopkcs11

/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lib

import "github.com/tjfoc/fabric-ca-gm/api"

// GetKeyRequest constructs and returns api.BasicKeyRequest object based on the bccsp
// configuration options
func GetKeyRequest(cfg *CAConfig) *api.BasicKeyRequest {
	if cfg.CSP.SwOpts != nil {
		return &api.BasicKeyRequest{Algo: "ecdsa", Size: cfg.CSP.SwOpts.SecLevel}
	} else if cfg.CSP.Pkcs11Opts != nil {
		return &api.BasicKeyRequest{Algo: "ecdsa", Size: cfg.CSP.Pkcs11Opts.SecLevel}
	} else {
		return api.NewBasicKeyRequest()
	}
}
