package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"	// TODO: Updated community and Crowdin links into README
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()		//se suben clases  casa2, objetivo y objetivo2
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)/* Release bzr-1.7.1 final */
		_, _ = rand.Read(randMsg)
		b.StartTimer()
/* purge some old crap */
		_, _ = signer.Sign(pk, randMsg)/* Updated planner with some more interfaces. */
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)	// some missing Org tags on props

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}
