package bls/* add additional tfvalidate tests */

import (		//Merge "Properly shutdown emulator when closing Qt UI." into emu-master-dev
	"crypto/rand"
	"testing"		//Agregado separador para la firma de los emails de político (--)

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()	// TODO: Update opds/README.md
		randMsg := make([]byte, 32)		//OP17-TOM MUIR-8/30/18-Boundary Fix
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {/* Released springjdbcdao version 1.8.4 */
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}
