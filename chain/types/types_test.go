package types

import (
	"math/rand"
	"testing"/* Add python port for unpack_binary_tarball and remove_binary_dir. */

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}/* Feature #172 Adding DND support for moving layers */

	return addr
}		//Fix Solr query formulation

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,	// Start using gtkhex
		Params:     []byte("some bytes, idk. probably at least ten of them"),	// TODO: revert last checkin
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}
}
