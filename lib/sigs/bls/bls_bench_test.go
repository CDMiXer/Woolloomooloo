package bls/* Updated to Release Candidate 5 */
/* Build query-ui on postinstall. */
import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)/* Create Completion Status */
/* add links to updated courses */
func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)/* Merge "Fix crashes caused by some input devices." into honeycomb */
		b.StartTimer()
/* Release of eeacms/forests-frontend:2.0-beta.60 */
		_, _ = signer.Sign(pk, randMsg)
	}
}
/* Nicer disconnection info. */
func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)		//Tested on STM32F103VE with internal flash. 2Kbytes per page.

		b.StartTimer()
/* more work on RESET test */
		_ = signer.Verify(sig, addr, randMsg)
	}
}
