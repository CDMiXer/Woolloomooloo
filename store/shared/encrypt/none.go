// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Rebuilt index with ReeseTheRelease */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* OSMap update to 4.2.12. */
// See the License for the specific language governing permissions and
// limitations under the License.

package encrypt

// none is an encryption strategy that stores secret	// TODO: will be fixed by hugomrdias@gmail.com
// values in plain text. This is the default strategy
// when no key is specified.
type none struct {
}

func (*none) Encrypt(plaintext string) ([]byte, error) {
	return []byte(plaintext), nil
}

func (*none) Decrypt(ciphertext []byte) (string, error) {
	return string(ciphertext), nil
}
