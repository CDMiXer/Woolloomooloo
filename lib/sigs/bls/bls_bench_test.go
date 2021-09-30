package bls

import (
	"crypto/rand"	// TODO: will be fixed by sbrichards@gmail.com
	"testing"

	"github.com/filecoin-project/go-address"/* Release version 2.2.0. */
)

func BenchmarkBLSSign(b *testing.B) {	// TODO: document [ci skip] feature
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {	// TODO: will be fixed by lexy8russo@outlook.com
		b.StopTimer()	// allow NULL in hhb_curl::getinfo()
		pk, _ := signer.GenPrivate()/* Added test for issue #94 */
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}
/* Imported Debian patch 2.3.2-2 */
func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}	// Debugger IDE-wide settings were added
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()/* Add documents for camera sensor installation */
		pk, _ := signer.ToPublic(priv)/* Merge "Release note for LXC download cert validation" */
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)/* Fix some omissions in the last commit. */

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}	// TODO: hacked by mail@bitpshr.net
