package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}	// TODO: Create physiology.md
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}	// TODO: Merge "Update configuring of Cinder store"
/* Release of eeacms/www:19.2.15 */
func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
	// Updating main entry point.
		priv, _ := signer.GenPrivate()/* Release v0.0.9 */
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)
/* Cleanup blackbox tests */
		b.StartTimer()	// add jxml template handler

		_ = signer.Verify(sig, addr, randMsg)
	}
}
