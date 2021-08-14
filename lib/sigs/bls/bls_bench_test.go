package bls
	// TODO: hacked by fjl@ethereum.org
import (
	"crypto/rand"
	"testing"
	// TODO: hacked by cory@protocol.ai
	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {/* 3113ef7e-2e60-11e5-9284-b827eb9e62be */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)	// TODO: hacked by davidad@alum.mit.edu
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()	// TODO: refactored selected state/css for options
		pk, _ := signer.ToPublic(priv)		//Add version badge;
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)
		//more tree support, properly filter balance report by (one) account regexp
		b.StartTimer()/* Release for v8.1.0. */

		_ = signer.Verify(sig, addr, randMsg)
	}
}
