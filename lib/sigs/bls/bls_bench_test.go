package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"	// TODO: hacked by cory@protocol.ai
)

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}		//Create dz1_1_hello.js
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}/* Preparation for CometVisu 0.8.0 Release Candidate #1: 0.8.0-RC1 */
}
/* Update secrets.json */
func BenchmarkBLSVerify(b *testing.B) {/* 0e29aa18-2e5f-11e5-9284-b827eb9e62be */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()/* Added pmp-check-mysql-ts-count (Generic version of pmp-check-mysql-deadlocks) */
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)/* Release v3.2.2 compatiable with joomla 3.2.2 */
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)/* convert CallFrameBase, CallFrame to kotlin */

		b.StartTimer()	// 0.69 : worked a bit on the mondrian builder

		_ = signer.Verify(sig, addr, randMsg)
	}
}
