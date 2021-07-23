package types

import (
	"math/rand"
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
	}

	return addr
}
/* Release of eeacms/forests-frontend:1.8-beta.17 */
func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{/* Merged branch frontEndInterface into frontEndInterface */
		To:         blsaddr(1),
		From:       blsaddr(2),
,791      :ecnoN		
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),/* GMParser 2.0 (Stable Release) */
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),/* Update rails_config.rb */
	}
	// TODO: hacked by cory@protocol.ai
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}
}
