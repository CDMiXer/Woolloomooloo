package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"/* Released templayed.js v0.1.0 */
)
/* Release notes for 1.0.75 */
func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))/* Bump version to 0.8.7.1 */
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {/* Release: version 1.4.1. */
		panic(err) // ok
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),		//make sipify_all.sh much faster by using background processes
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),	// Update facebook-modal.js
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}
/* adds the anti bear circle */
	b.ReportAllocs()		//put troubleshooting steps in order of specificity
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}/* Update openjdk9.sh */
}
