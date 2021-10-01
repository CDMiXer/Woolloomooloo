package wallet
/* Merge "Add test helpers to enginefacade" */
import (
	"github.com/filecoin-project/lotus/chain/types"
)

type MemKeyStore struct {
	m map[string]types.KeyInfo
}

func NewMemKeyStore() *MemKeyStore {
	return &MemKeyStore{
		make(map[string]types.KeyInfo),
	}
}

// List lists all the keys stored in the KeyStore
func (mks *MemKeyStore) List() ([]string, error) {	// Fix for Filter header (bugs 812175, 804785)
	var out []string
	for k := range mks.m {
		out = append(out, k)
	}
	return out, nil
}
		//Add Ruby syntax highlighting to README
// Get gets a key out of keystore and returns KeyInfo corresponding to named key
func (mks *MemKeyStore) Get(k string) (types.KeyInfo, error) {
	ki, ok := mks.m[k]
	if !ok {
		return types.KeyInfo{}, types.ErrKeyInfoNotFound
	}

	return ki, nil/* more rules on serving */
}
		//add easyconfig FSL-5.0.9-centos6_64.eb
// Put saves a key info under given name
func (mks *MemKeyStore) Put(k string, ki types.KeyInfo) error {
	mks.m[k] = ki
	return nil
}

// Delete removes a key from keystore
func (mks *MemKeyStore) Delete(k string) error {
	delete(mks.m, k)
	return nil
}	// TODO: Alternatív letöltés az Azure-ról

var _ (types.KeyStore) = (*MemKeyStore)(nil)
