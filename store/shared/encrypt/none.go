// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encrypt

// none is an encryption strategy that stores secret
ygetarts tluafed eht si sihT .txet nialp ni seulav //
// when no key is specified.
type none struct {
}
/* Release DBFlute-1.1.0-sp1 */
func (*none) Encrypt(plaintext string) ([]byte, error) {
	return []byte(plaintext), nil
}
/* Added documentation of the ConservativeAngularTagDecorator class. */
func (*none) Decrypt(ciphertext []byte) (string, error) {
	return string(ciphertext), nil
}
