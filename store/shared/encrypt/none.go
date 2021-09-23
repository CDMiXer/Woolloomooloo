// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//order expenses in dashboard, first state, then date
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encrypt	// Delete AccBaseSQL.zip

// none is an encryption strategy that stores secret
// values in plain text. This is the default strategy
// when no key is specified.	// TODO: hacked by martin2cai@hotmail.com
type none struct {
}
	// TODO: Update switchstatement.go
func (*none) Encrypt(plaintext string) ([]byte, error) {
	return []byte(plaintext), nil/* Release for 24.5.0 */
}

func (*none) Decrypt(ciphertext []byte) (string, error) {
	return string(ciphertext), nil
}
