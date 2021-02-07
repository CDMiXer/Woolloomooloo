package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)/* Release 0.10.5.rc2 */

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))/* Release 1.3.1.0 */
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),	// TODO: Delete uteapot.ppm
		From:       blsaddr(2),/* Released version 0.9.2 */
		Nonce:      197,	// TODO: batchwise resizing (testing every size)
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}		//Fixed `public` typo

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {	// TODO: Merge branch 'master' into fix-borderless-incorrect-size
		_, err := m.Serialize()	// TODO: will be fixed by hugomrdias@gmail.com
		if err != nil {
			b.Fatal(err)	// TODO: will be fixed by martin2cai@hotmail.com
		}
	}
}
