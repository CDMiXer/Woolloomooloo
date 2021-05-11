package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)	// TODO: url test for non-cached pages was not being run if only one page was returned
/* Updating header locations for latest locm3. */
	addr, err := address.NewBLSAddress(buf)
	if err != nil {		//Create .test.basic.vim
		panic(err) // ok
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{/* Create search-component.js */
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,		//Ongoing rendertarget work.
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)		//[tbsl] final changes
		}
	}
}
