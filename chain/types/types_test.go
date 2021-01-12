package types	// win32-fix dev+ build compilation
		//Issue 3051:  Let heapq work with either __lt__ or __le__.
import (
	"math/rand"
	"testing"
	// link permanente
	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}		//Add CHANGELOG-1.18.md for v1.18.0-alpha.3

	return addr	// TODO: Minor update to wording.
}

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
		_, err := m.Serialize()		//added support for multiple groups sections in access file
		if err != nil {		//var was not defined
			b.Fatal(err)
		}
	}
}
