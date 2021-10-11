package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)
/* Delete edit-routing.png */
func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()/* Add `gradle-qemu` plugin */
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()/* Release version 2.4.1 */

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {		//[#762] change class name
		b.StopTimer()	// rebuilt with @GPXCAT added!
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()		//Migrate from gitter
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)		//7d61bada-2e62-11e5-9284-b827eb9e62be
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()
	// TODO: hacked by sebastian.tharakan97@gmail.com
		_ = signer.Verify(sig, addr, randMsg)		//Vickers Medium Mk. I is a medium tank
	}
}
