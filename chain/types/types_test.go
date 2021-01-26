package types/* cda2049c-2e52-11e5-9284-b827eb9e62be */

import (
	"math/rand"	// TODO: will be fixed by fjl@ethereum.org
	"testing"

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}	// TODO: hacked by fjl@ethereum.org

	return addr
}	// TODO: return snippets in original order

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,/* Add license and remove unused variables */
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,/* [Build] Gulp Release Task #82 */
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),/* Parse UPnP service ID from root description and expose it to consumers */
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}		//clarify deploy docs
}
