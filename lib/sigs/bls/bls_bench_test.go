package bls

import (
	"crypto/rand"
	"testing"		//Create plotSTR.r

	"github.com/filecoin-project/go-address"
)	// TODO: will be fixed by witek@enjin.io

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)	// win32: add shelve extension to mercurial.ini
	}/* Bump EclipseRelease.LATEST to 4.6.3. */
}/* Merge "Fixes resource name problem in "Resources Usage" tab" */

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()/* Release LastaTaglib-0.6.5 */
		randMsg := make([]byte, 32)		//Delete quotes
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()		//Tools: DFG: Beautify Register output by adding LSB start index information.

		_ = signer.Verify(sig, addr, randMsg)
	}
}
