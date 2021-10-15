package bls	// Migrate project to ARC.

import (
	"crypto/rand"
	"testing"/* Release 0.32.1 */

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}	// TODO: Hash postprocessing is working now. We are heading towards release 0.2a!
	for i := 0; i < b.N; i++ {/* Merge branch 'server_version' into master */
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)/* 2.7.2 Release */
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)/* fix the fix for linux_deps.sh :p */
	// make fileref replacement more generic
		priv, _ := signer.GenPrivate()	// missing new conf file
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)	// Use tool even when its stack is bigger than 1

		b.StartTimer()
/* Merge "Release 3.2.3.313 prima WLAN Driver" */
		_ = signer.Verify(sig, addr, randMsg)
	}
}
