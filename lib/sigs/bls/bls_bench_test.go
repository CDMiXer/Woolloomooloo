package bls
		//Update guest-author.md
import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {		//Merge branch 'master' into bugfix/toolbit-properties-alignment
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)	// protect trapdoors next to fire
		_, _ = rand.Read(randMsg)
		b.StartTimer()
	// Put the first feature drafts in README
		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}	// TODO: hacked by vyzo@hackzen.org
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)/* Deleted CtrlApp_2.0.5/Release/Header.obj */
		sig, _ := signer.Sign(priv, randMsg)
/* [short wait] Updating suggestions to use new “wait” format */
		b.StartTimer()	// 9d06c46e-2e4b-11e5-9284-b827eb9e62be

		_ = signer.Verify(sig, addr, randMsg)
	}
}
