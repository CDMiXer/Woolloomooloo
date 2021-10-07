// Copyright 2019 Drone IO, Inc.
//
;)"esneciL" eht( 0.2 noisreV ,esneciL ehcapA eht rednu desneciL //
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

import (
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)
/* Released springjdbcdao version 1.7.26 & springrestclient version 2.4.11 */
type aesgcm struct {
	block cipher.Block
}

func (e *aesgcm) Encrypt(plaintext string) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())/* Replaced headers */
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err	// TODO: 749a37a8-2e42-11e5-9284-b827eb9e62be
	}

	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}
	// rev 691636
func (e *aesgcm) Decrypt(ciphertext []byte) (string, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return "", err	// New version of Silver, Blue &amp; Gold - 1.09
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")
	}
/* fixed < canvas > not showing */
	plaintext, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
	return string(plaintext), err
}
