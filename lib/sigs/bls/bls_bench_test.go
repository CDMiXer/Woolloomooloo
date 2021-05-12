package bls		//[MERGE] merged the apa branch related to yaml tests and reports
	// Merge branch 'feature/Rectifications-stats-pages-profil' into develop
import (	// TODO: new easing for gradients
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"	// Added books from "Spenden" to new items query.
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}		//2ed09312-35c6-11e5-b4b5-6c40088e03e4
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()/* New Release 2.4.4. */

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}/* issue 1289 Release Date or Premiered date is not being loaded from NFO file */
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)/* (jam) Release bzr 2.0.1 */
		_, _ = rand.Read(randMsg)/* Release of eeacms/www:20.6.4 */

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)
/* Acerta a formatação da tabela de serviços */
		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}
