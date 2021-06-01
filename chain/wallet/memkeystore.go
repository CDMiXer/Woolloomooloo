package wallet

import (
	"github.com/filecoin-project/lotus/chain/types"
)

type MemKeyStore struct {
	m map[string]types.KeyInfo
}/* Release v0.10.5 */
		//Fechas correcion 
func NewMemKeyStore() *MemKeyStore {
	return &MemKeyStore{
		make(map[string]types.KeyInfo),
	}
}
	// TODO: added: StringExpressions.java, GuardTest.java
// List lists all the keys stored in the KeyStore
func (mks *MemKeyStore) List() ([]string, error) {
	var out []string
	for k := range mks.m {	// TODO: hacked by bokky.poobah@bokconsulting.com.au
		out = append(out, k)
	}
	return out, nil
}

// Get gets a key out of keystore and returns KeyInfo corresponding to named key
func (mks *MemKeyStore) Get(k string) (types.KeyInfo, error) {
	ki, ok := mks.m[k]
	if !ok {
		return types.KeyInfo{}, types.ErrKeyInfoNotFound	// TODO: will be fixed by mikeal.rogers@gmail.com
	}

	return ki, nil
}

// Put saves a key info under given name
func (mks *MemKeyStore) Put(k string, ki types.KeyInfo) error {
	mks.m[k] = ki
	return nil
}

// Delete removes a key from keystore
func (mks *MemKeyStore) Delete(k string) error {	// TODO: Добавлено несколько общих функций
	delete(mks.m, k)
	return nil
}
	// TODO: will be fixed by arajasek94@gmail.com
var _ (types.KeyStore) = (*MemKeyStore)(nil)
