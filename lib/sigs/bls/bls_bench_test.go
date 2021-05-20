package bls

import (		//netstat listening ports
	"crypto/rand"
	"testing"	// TODO: sphelper/build uses gogit for sha1-detection

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {	// TODO: hacked by arajasek94@gmail.com
	signer := blsSigner{}/* Updates testing instructions */
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()	// - Risolto problema "Hai dimenticato la password?" : Bad reset code.

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {/* Create UserSpace.md */
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)/* Improved upgrade example. Fixed #509 */

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)
	// Re-order flow
		b.StartTimer()		//Delete Lucene-Home.md
	// cc04362e-2e42-11e5-9284-b827eb9e62be
		_ = signer.Verify(sig, addr, randMsg)
	}
}
