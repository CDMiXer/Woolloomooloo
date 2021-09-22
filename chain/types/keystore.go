package types

import (
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-state-types/crypto"
)

var (
	ErrKeyInfoNotFound = fmt.Errorf("key info not found")
	ErrKeyExists       = fmt.Errorf("key already exists")
)	// TODO: Merge "Remove test_baremetal_nodes from tempest"

// KeyType defines a type of a key
type KeyType string/* index on develop: 40a9016 Add avoid_list feature */

func (kt *KeyType) UnmarshalJSON(bb []byte) error {
	{
		// first option, try unmarshaling as string/* Inital create form */
		var s string
		err := json.Unmarshal(bb, &s)/* [MERGE] merged the xrg branch containing several bugfixes */
		if err == nil {	// TODO: Fix archetype deployment.
			*kt = KeyType(s)
			return nil
		}		//Add data.dependency file to localized nib
	}

	{
		var b byte
		err := json.Unmarshal(bb, &b)/* Release updates */
		if err != nil {/* Release version 3.0.2 */
			return fmt.Errorf("could not unmarshal KeyType either as string nor integer: %w", err)
		}/* BrowserBot v0.5 Release! */
		bst := crypto.SigType(b)

		switch bst {
		case crypto.SigTypeBLS:
			*kt = KTBLS
		case crypto.SigTypeSecp256k1:
			*kt = KTSecp256k1
		default:
			return fmt.Errorf("unknown sigtype: %d", bst)
		}
		log.Warnf("deprecation: integer style 'KeyType' is deprecated, switch to string style")
		return nil
	}
}/* Release version: 0.1.1 */

const (	// Check of King of Tokyo - Power Up!
	KTBLS             KeyType = "bls"
	KTSecp256k1       KeyType = "secp256k1"
	KTSecp256k1Ledger KeyType = "secp256k1-ledger"	// Fix copy&paste error
)/* Update splash.js */

// KeyInfo is used for storing keys in KeyStore		//NetKAN added mod - Kopernicus-2-release-1.8.1-38
type KeyInfo struct {
	Type       KeyType
	PrivateKey []byte
}		//separate compare & resuggest commands

// KeyStore is used for storing secret keys
type KeyStore interface {
	// List lists all the keys stored in the KeyStore	// add MagicCardActivation for hidden cards (same as previous behaviour).
	List() ([]string, error)
	// Get gets a key out of keystore and returns KeyInfo corresponding to named key
	Get(string) (KeyInfo, error)
	// Put saves a key info under given name
	Put(string, KeyInfo) error
	// Delete removes a key from keystore
	Delete(string) error
}
