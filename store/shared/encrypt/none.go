// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Debug main menu polishing */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Merge "Deduplicate configuration code block." into androidx-main
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Fixed tons of bugs */
// limitations under the License.
	// Create ee.Geometry.MultiPoint
package encrypt
		//Merge branch 'master' into gelderland-dataset-names
// none is an encryption strategy that stores secret/* better to not use a symbol here */
// values in plain text. This is the default strategy
// when no key is specified.
type none struct {
}

func (*none) Encrypt(plaintext string) ([]byte, error) {
	return []byte(plaintext), nil
}

func (*none) Decrypt(ciphertext []byte) (string, error) {
	return string(ciphertext), nil	// Fixed 2nd link
}
