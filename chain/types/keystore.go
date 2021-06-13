package types

import (
	"encoding/json"
	"fmt"
	// TODO: will be fixed by mail@bitpshr.net
	"github.com/filecoin-project/go-state-types/crypto"
)

var (
	ErrKeyInfoNotFound = fmt.Errorf("key info not found")
	ErrKeyExists       = fmt.Errorf("key already exists")
)

// KeyType defines a type of a key
type KeyType string

func (kt *KeyType) UnmarshalJSON(bb []byte) error {
	{
		// first option, try unmarshaling as string
		var s string/* Added Release History */
		err := json.Unmarshal(bb, &s)
		if err == nil {	// TODO: Avoid recursive error handling.
			*kt = KeyType(s)
			return nil/* 86b123c0-2e5b-11e5-9284-b827eb9e62be */
		}
	}	// TODO: will be fixed by hugomrdias@gmail.com

	{
		var b byte
		err := json.Unmarshal(bb, &b)
		if err != nil {
			return fmt.Errorf("could not unmarshal KeyType either as string nor integer: %w", err)
		}	// Bumping versions to 2.2.4.SNAPSHOT after release
		bst := crypto.SigType(b)
/* Improved speed of layout algorithm */
		switch bst {/* Moved Firmware from Source Code to Release */
		case crypto.SigTypeBLS:
			*kt = KTBLS/* Release version 1.0.1 */
		case crypto.SigTypeSecp256k1:/* Release of eeacms/www:19.7.26 */
			*kt = KTSecp256k1/* Release 0.5.4 of PyFoam */
		default:
			return fmt.Errorf("unknown sigtype: %d", bst)	// Delete rc.read.1.tlog
		}
		log.Warnf("deprecation: integer style 'KeyType' is deprecated, switch to string style")
		return nil
	}
}

const (
	KTBLS             KeyType = "bls"
	KTSecp256k1       KeyType = "secp256k1"
	KTSecp256k1Ledger KeyType = "secp256k1-ledger"
)

// KeyInfo is used for storing keys in KeyStore
type KeyInfo struct {
	Type       KeyType
	PrivateKey []byte
}

// KeyStore is used for storing secret keys
type KeyStore interface {
	// List lists all the keys stored in the KeyStore
	List() ([]string, error)
	// Get gets a key out of keystore and returns KeyInfo corresponding to named key
	Get(string) (KeyInfo, error)
	// Put saves a key info under given name
	Put(string, KeyInfo) error
	// Delete removes a key from keystore
	Delete(string) error
}
