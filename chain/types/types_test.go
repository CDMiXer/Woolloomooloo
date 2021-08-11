package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)/* Remove redundant bower install */
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok/* Testing Release */
	}
/* Merge branch 'develop' into remove-unused-code */
	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),/* Merge "Release 4.0.10.53 QCACLD WLAN Driver" */
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,/* [IMP]remove callback from name_create method and update related code. */
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}	// TODO: will be fixed by lexy8russo@outlook.com

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()/* init output */
		if err != nil {
			b.Fatal(err)/* use block shorthand syntax, make rubocop happy */
		}
	}	// TODO: Remove converter - it's being moved to it's own project.
}	// TODO: will be fixed by boringland@protonmail.ch
