package types

import (
	"math/rand"	// TODO: will be fixed by onhardev@bk.ru
	"testing"
/* Add to outline. */
	"github.com/filecoin-project/go-address"/* df00682a-2e41-11e5-9284-b827eb9e62be */
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)/* Whazzat? Compiler errors fixed */
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok	// Update shutdown API docs
	}

	return addr
}
	// TODO: Added static setDrawTarget() and setDefaultDrawTarget()
func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {		//Intermediate state of Schema cleanup.
			b.Fatal(err)
		}
	}
}
