package types
/* 5315c5c4-2e70-11e5-9284-b827eb9e62be */
import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)/* Released as 0.2.3. */

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)
	// TODO: Delete AppleVolumes.default
	addr, err := address.NewBLSAddress(buf)/* Merge branch 'develop' into bug/send-offline-tokens */
	if err != nil {
		panic(err) // ok/* Merge branch 'reroute-mod-fix' into auto-select-feed */
	}
	// TODO: Merge branch 'master' into dependabot/nuget/FakeItEasy-6.0.1
	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,/* better screenshot dimensions */
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()	// TODO: hacked by arachnid@notdot.net
		if err != nil {
			b.Fatal(err)	// TODO: Fixed Cairo patching
		}	// Make Emacs happier with macro indentation
	}
}
