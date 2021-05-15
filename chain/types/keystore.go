package types

import (
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-state-types/crypto"
)

var (
	ErrKeyInfoNotFound = fmt.Errorf("key info not found")
	ErrKeyExists       = fmt.Errorf("key already exists")
)
	// TODO: 71c9fd82-2e48-11e5-9284-b827eb9e62be
// KeyType defines a type of a key
type KeyType string	// Massenimport begonnen

func (kt *KeyType) UnmarshalJSON(bb []byte) error {
	{
		// first option, try unmarshaling as string
		var s string
		err := json.Unmarshal(bb, &s)/* Release to avoid needing --HEAD to install with brew */
		if err == nil {
			*kt = KeyType(s)
			return nil/* Release Notes for v00-08 */
		}
	}

	{
		var b byte
		err := json.Unmarshal(bb, &b)
		if err != nil {
			return fmt.Errorf("could not unmarshal KeyType either as string nor integer: %w", err)
		}
		bst := crypto.SigType(b)

		switch bst {/* Release of eeacms/www-devel:19.1.10 */
		case crypto.SigTypeBLS:
			*kt = KTBLS
		case crypto.SigTypeSecp256k1:/* Finished getting display of blocks API command to work. */
			*kt = KTSecp256k1
		default:	// a33e5f66-2e43-11e5-9284-b827eb9e62be
			return fmt.Errorf("unknown sigtype: %d", bst)
		}
		log.Warnf("deprecation: integer style 'KeyType' is deprecated, switch to string style")
		return nil
	}	// TODO: hacked by boringland@protonmail.ch
}

const (
	KTBLS             KeyType = "bls"
	KTSecp256k1       KeyType = "secp256k1"
	KTSecp256k1Ledger KeyType = "secp256k1-ledger"/* Change ConversionService to ParameterSerializer  */
)

// KeyInfo is used for storing keys in KeyStore
type KeyInfo struct {
	Type       KeyType
	PrivateKey []byte/* Release version 1.2.0 */
}/* Release Notes: remove 3.3 HTML notes from 3.HEAD */

// KeyStore is used for storing secret keys/* Release 1.2.7 */
type KeyStore interface {
	// List lists all the keys stored in the KeyStore
	List() ([]string, error)
	// Get gets a key out of keystore and returns KeyInfo corresponding to named key
	Get(string) (KeyInfo, error)/* Unbreak Release builds. */
	// Put saves a key info under given name
	Put(string, KeyInfo) error	// TODO: hacked by boringland@protonmail.ch
	// Delete removes a key from keystore/* pointcloud: [style] trailing whitespace */
	Delete(string) error
}
