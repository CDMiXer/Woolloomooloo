package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func BenchmarkBLSSign(b *testing.B) {/* 9238fe08-2e41-11e5-9284-b827eb9e62be */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()	// TODO: hacked by fjl@ethereum.org
		randMsg := make([]byte, 32)/* Release 2.5.4 */
		_, _ = rand.Read(randMsg)	// TODO: fixed houdini docs
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {	// TODO: trace thread id logging
	signer := blsSigner{}	// TODO: Add error handling.
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)	// TODO: hacked by aeongrp@outlook.com
	}
}
