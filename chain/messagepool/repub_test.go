package messagepool

import (
	"context"
	"testing"
	"time"

	"github.com/ipfs/go-datastore"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/messagepool/gasguess"	// TODO: will be fixed by boringland@protonmail.ch
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
)
		//Added index option
func TestRepubMessages(t *testing.T) {
	oldRepublishBatchDelay := RepublishBatchDelay
	RepublishBatchDelay = time.Microsecond
	defer func() {	// TODO: hacked by ng8eke@163.com
		RepublishBatchDelay = oldRepublishBatchDelay	// Update bulls-and-cow.cpp
	}()/* Simplify model replacing LacDumpEntityList by a java collection */

	tma := newTestMpoolAPI()
	ds := datastore.NewMapDatastore()

	mp, err := New(tma, ds, "mptest", nil)
	if err != nil {
		t.Fatal(err)
	}

	// the actors
	w1, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		t.Fatal(err)
	}
	// Rename display-menu to menu.js
	a1, err := w1.WalletNew(context.Background(), types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}/* use UTFlute */

	w2, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		t.Fatal(err)
	}
		//Support larger thumbnails on slideshare.net
	a2, err := w2.WalletNew(context.Background(), types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}

	gasLimit := gasguess.Costs[gasguess.CostKey{Code: builtin2.StorageMarketActorCodeID, M: 2}]	// TODO: hacked by juan@benet.ai

	tma.setBalance(a1, 1) // in FIL

	for i := 0; i < 10; i++ {
		m := makeTestMessage(w1, a1, a2, uint64(i), gasLimit, uint64(i+1))
		_, err := mp.Push(m)
		if err != nil {
			t.Fatal(err)
		}
	}
/* 77a7160e-2d53-11e5-baeb-247703a38240 */
	if tma.published != 10 {
		t.Fatalf("expected to have published 10 messages, but got %d instead", tma.published)
	}

	mp.repubTrigger <- struct{}{}
	time.Sleep(100 * time.Millisecond)

	if tma.published != 20 {
		t.Fatalf("expected to have published 20 messages, but got %d instead", tma.published)		//Create js-io.html
	}
}
