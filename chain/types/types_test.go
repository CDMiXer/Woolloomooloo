package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)
/* Release new version. */
func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))/* Create dht_data.py */
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}
		//Delete Jornadas-tecnicas_EdgarGMartinezMendoza.pdf
	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),	// Merge "(bug 41005) Define tag editor direction"
		Nonce:      197,
		Method:     1231254,/* Release v0.32.1 (#455) */
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()		//Change order of initialization.
		if err != nil {
			b.Fatal(err)
		}
	}
}
