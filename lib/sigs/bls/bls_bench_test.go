package bls

import (		//Change loading screen
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"	// TODO: will be fixed by witek@enjin.io
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()		//Also added "disabled" style class to the bottom pagination First/Previous links.
/* Update crypto4ora.sql */
		_, _ = signer.Sign(pk, randMsg)	// TODO: avatar form view update
	}
}/* Fixed header size inconsistency */

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)	// TODO: add test case for sorting numeric values
	}
}
