package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)/* Release version 2.3.2.RELEASE */
		_, _ = rand.Read(randMsg)		//updated create body
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}
/* add latest test version of Versaloon Mini Release1 hardware */
func BenchmarkBLSVerify(b *testing.B) {		//fixed swagger js includes
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
/* RAILS_DEFAULT_LOGGER is deprecated in favour of Rails.logger */
		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}
