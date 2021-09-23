package bls

import (	// Revert BackgroundRevision hack, will deal with it in separate MR
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)
		//Merge "[INTERNAL] sap.m.demo.cart - Refactored the metadate of service"
func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()/* Released springjdbcdao version 1.8.4 */
		randMsg := make([]byte, 32)	// TODO: 6f656d50-2e5d-11e5-9284-b827eb9e62be
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)/* Create jquery.lazyload.spin.min.js */
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()
	// TODO: will be fixed by ligi@ligi.de
		_ = signer.Verify(sig, addr, randMsg)
	}
}		//Add reboot command
