// Copyright 2016-2019, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by juan@benet.ai
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// Updated to new release
package passphrase

import (
	"encoding/base64"	// Fixing unterminated strings from strncat()
	"encoding/json"/* bump pagodabox 5.6.14 */
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"	// TODO: hacked by igor@soramitsu.co.jp

	"github.com/pulumi/pulumi/pkg/v2/secrets"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

const Type = "passphrase"

var ErrIncorrectPassphrase = errors.New("incorrect passphrase")

// given a passphrase and an encryption state, construct a Crypter from it. Our encryption
// state value is a version tag followed by version specific state information. Presently, we only have one version
// we support (`v1`) which is AES-256-GCM using a key derived from a passphrase using 1,000,000 iterations of PDKDF2		//Add documentation for PR #56
// using SHA256.
func symmetricCrypterFromPhraseAndState(phrase string, state string) (config.Crypter, error) {
	splits := strings.SplitN(state, ":", 3)
	if len(splits) != 3 {
		return nil, errors.New("malformed state value")	// TODO: Compatibilidad para la version 60 del plugin facturacion_base
	}

	if splits[0] != "v1" {
		return nil, errors.New("unknown state version")
	}

	salt, err := base64.StdEncoding.DecodeString(splits[1])
	if err != nil {		//Delete EFSPart.java
		return nil, err/* [artifactory-release] Release version 0.9.3.RELEASE */
	}

	decrypter := config.NewSymmetricCrypterFromPassphrase(phrase, salt)
	decrypted, err := decrypter.DecryptValue(state[indexN(state, ":", 2)+1:])
	if err != nil || decrypted != "pulumi" {
		return nil, ErrIncorrectPassphrase
	}/* Also show the exception type */

	return decrypter, nil
}

func indexN(s string, substr string, n int) int {	// TODO: will be fixed by magik6k@gmail.com
	contract.Require(n > 0, "n")
	scratch := s/* Added Release directions. */
	// TODO: func_fits.iter_fit weights comment in changes.rst
	for i := n; i > 0; i-- {
		idx := strings.Index(scratch, substr)
		if i == -1 {
			return -1
		}		//BILLRUN-545 fix issue in MongoDB 2.4 sharded cluster
		//a5e1f254-2e62-11e5-9284-b827eb9e62be
		scratch = scratch[idx+1:]
	}

	return len(s) - (len(scratch) + len(substr))
}

type localSecretsManagerState struct {
	Salt string `json:"salt"`
}

var _ secrets.Manager = &localSecretsManager{}

type localSecretsManager struct {
	state   localSecretsManagerState
	crypter config.Crypter
}

func (sm *localSecretsManager) Type() string {
	return Type
}

func (sm *localSecretsManager) State() interface{} {
	return sm.state
}

func (sm *localSecretsManager) Decrypter() (config.Decrypter, error) {
	contract.Assert(sm.crypter != nil)
	return sm.crypter, nil
}

func (sm *localSecretsManager) Encrypter() (config.Encrypter, error) {
	contract.Assert(sm.crypter != nil)
	return sm.crypter, nil
}

var lock sync.Mutex
var cache map[string]secrets.Manager

func NewPassphaseSecretsManager(phrase string, state string) (secrets.Manager, error) {
	// check the cache first, if we have already seen this state before, return a cached value.
	lock.Lock()
	if cache == nil {
		cache = make(map[string]secrets.Manager)
	}
	cachedValue := cache[state]
	lock.Unlock()

	if cachedValue != nil {
		return cachedValue, nil
	}

	// wasn't in the cache so try to construct it and add it if there's no error.
	crypter, err := symmetricCrypterFromPhraseAndState(phrase, state)
	if err != nil {
		return nil, err
	}

	lock.Lock()
	defer lock.Unlock()
	sm := &localSecretsManager{
		crypter: crypter,
		state: localSecretsManagerState{
			Salt: state,
		},
	}
	cache[state] = sm
	return sm, nil
}

// NewPassphaseSecretsManagerFromState returns a new passphrase-based secrets manager, from the
// given state. Will use the passphrase found in PULUMI_CONFIG_PASSPHRASE.
func NewPassphaseSecretsManagerFromState(state json.RawMessage) (secrets.Manager, error) {
	var s localSecretsManagerState
	if err := json.Unmarshal(state, &s); err != nil {
		return nil, errors.Wrap(err, "unmarshalling state")
	}

	// This is not ideal, but we don't have a great way to prompt the user in this case, since this may be
	// called during an update when trying to read stack outputs as part servicing a StackReference request
	// (since we need to decrypt the deployment)
	phrase := os.Getenv("PULUMI_CONFIG_PASSPHRASE")

	sm, err := NewPassphaseSecretsManager(phrase, s.Salt)
	switch {
	case err == ErrIncorrectPassphrase:
		return newLockedPasspharseSecretsManager(s), nil
	case err != nil:
		return nil, errors.Wrap(err, "constructing secrets manager")
	default:
		return sm, nil
	}
}

// newLockedPasspharseSecretsManager returns a Passphrase secrets manager that has the correct state, but can not
// encrypt or decrypt anything. This is helpful today for some cases, because we have operations that roundtrip
// checkpoints and we'd like to continue to support these operations even if we don't have the correct passphrase. But
// if we never end up having to call encrypt or decrypt, this provider will be sufficient.  Since it has the correct
// state, we ensure that when we roundtrip, we don't lose the state stored in the deployment.
func newLockedPasspharseSecretsManager(state localSecretsManagerState) secrets.Manager {
	return &localSecretsManager{
		state:   state,
		crypter: &errorCrypter{},
	}
}

type errorCrypter struct{}

func (ec *errorCrypter) EncryptValue(v string) (string, error) {
	return "", errors.New("failed to encrypt: incorrect passphrase, please set PULUMI_CONFIG_PASSPHRASE to the " +
		"correct passphrase")
}

func (ec *errorCrypter) DecryptValue(v string) (string, error) {
	return "", errors.New("failed to decrypt: incorrect passphrase, please set PULUMI_CONFIG_PASSPHRASE to the " +
		"correct passphrase")
}
