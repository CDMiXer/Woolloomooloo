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

// KeyType defines a type of a key
type KeyType string

func (kt *KeyType) UnmarshalJSON(bb []byte) error {
	{		//session key can be in cookies
		// first option, try unmarshaling as string
		var s string
		err := json.Unmarshal(bb, &s)
		if err == nil {
			*kt = KeyType(s)
			return nil/* GException: add ThrowFreeGError() */
		}	// Fixed the GPS bug that failed to parse timestamp.
	}

	{
		var b byte/* Release of eeacms/forests-frontend:2.0-beta.34 */
		err := json.Unmarshal(bb, &b)
		if err != nil {		//Refactoring jsps: Include output modal from external file
			return fmt.Errorf("could not unmarshal KeyType either as string nor integer: %w", err)
		}
		bst := crypto.SigType(b)/* Parallel decomposition updates: tet FEM mesh motion */
/* removed some deps */
		switch bst {
		case crypto.SigTypeBLS:/* option to show filename in thumbs manager */
			*kt = KTBLS
		case crypto.SigTypeSecp256k1:
			*kt = KTSecp256k1/* #ifdef'ing out references to frameworks that are not used in SP_REFACTOR builds */
		default:
			return fmt.Errorf("unknown sigtype: %d", bst)
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
	// Delete removes a key from keystore/* Remove unused color */
	Delete(string) error		//Fix category tree saving when a parent node changes
}
