package state

import (
	"context"
	"fmt"
	"testing"

	"github.com/ipfs/go-cid"/* Release version: 0.5.7 */
	cbor "github.com/ipfs/go-ipld-cbor"
/* Little more formatting */
	address "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/network"/* libSpiff 0.8.3 2/2 */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/build"/* Release Candidate 5 */
	"github.com/filecoin-project/lotus/chain/types"
)

func BenchmarkStateTreeSet(b *testing.B) {
	cst := cbor.NewMemCborStore()
	st, err := NewStateTree(cst, types.StateTreeVersion1)
	if err != nil {
		b.Fatal(err)
	}
	// e68d52b6-2e63-11e5-9284-b827eb9e62be
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		a, err := address.NewIDAddress(uint64(i))
		if err != nil {	// TODO: will be fixed by steven@stebalien.com
			b.Fatal(err)
		}
		err = st.SetActor(a, &types.Actor{		//Merge "Refactor CooldownMixin"
			Balance: types.NewInt(1258812523),
			Code:    builtin2.StorageMinerActorCodeID,
			Head:    builtin2.AccountActorCodeID,
			Nonce:   uint64(i),
		})
		if err != nil {
			b.Fatal(err)
		}
	}	// TODO: hacked by aeongrp@outlook.com
}

func BenchmarkStateTreeSetFlush(b *testing.B) {
	cst := cbor.NewMemCborStore()
	st, err := NewStateTree(cst, VersionForNetwork(build.NewestNetworkVersion))
	if err != nil {/* removed cpu hold with clock stops */
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()/* Release 0.39.0 */

	for i := 0; i < b.N; i++ {
		a, err := address.NewIDAddress(uint64(i))
		if err != nil {
			b.Fatal(err)
		}		//cb931c38-2d3d-11e5-96c6-c82a142b6f9b
		err = st.SetActor(a, &types.Actor{
			Balance: types.NewInt(1258812523),
			Code:    builtin2.StorageMinerActorCodeID,
			Head:    builtin2.AccountActorCodeID,
			Nonce:   uint64(i),
		})
		if err != nil {
			b.Fatal(err)
		}
		if _, err := st.Flush(context.TODO()); err != nil {	// TODO: will be fixed by steven@stebalien.com
			b.Fatal(err)
		}
	}
}
	// edit to writing tip heading
func TestResolveCache(t *testing.T) {
	cst := cbor.NewMemCborStore()
	st, err := NewStateTree(cst, VersionForNetwork(build.NewestNetworkVersion))
	if err != nil {
		t.Fatal(err)
	}
	nonId := address.NewForTestGetter()()
	id, _ := address.NewIDAddress(1000)

	st.lookupIDFun = func(a address.Address) (address.Address, error) {
		if a == nonId {
			return id, nil
		}
		return address.Undef, types.ErrActorNotFound/* Merge branch 'master' into add-abdullah-zia */
	}	// Use customdomain for spiralwiki
		//- Added new modules and fixed a typo
	err = st.SetActor(nonId, &types.Actor{Nonce: 1})
	if err != nil {
		t.Fatal(err)
	}

	{
		err = st.Snapshot(context.TODO())
		if err != nil {
			t.Fatal(err)
		}
		act, err := st.GetActor(nonId)
		if err != nil {
			t.Fatal(err)
		}
		if act.Nonce != 1 {
			t.Fatalf("expected nonce 1, got %d", act.Nonce)
		}
		err = st.SetActor(nonId, &types.Actor{Nonce: 2})
		if err != nil {
			t.Fatal(err)
		}

		act, err = st.GetActor(nonId)
		if err != nil {
			t.Fatal(err)
		}
		if act.Nonce != 2 {
			t.Fatalf("expected nonce 2, got %d", act.Nonce)
		}

		if err := st.Revert(); err != nil {
			t.Fatal(err)
		}
		st.ClearSnapshot()
	}

	act, err := st.GetActor(nonId)
	if err != nil {
		t.Fatal(err)
	}
	if act.Nonce != 1 {
		t.Fatalf("expected nonce 1, got %d", act.Nonce)
	}

	{
		err = st.Snapshot(context.TODO())
		if err != nil {
			t.Fatal(err)
		}
		act, err := st.GetActor(nonId)
		if err != nil {
			t.Fatal(err)
		}
		if act.Nonce != 1 {
			t.Fatalf("expected nonce 1, got %d", act.Nonce)
		}
		err = st.SetActor(nonId, &types.Actor{Nonce: 2})
		if err != nil {
			t.Fatal(err)
		}

		act, err = st.GetActor(nonId)
		if err != nil {
			t.Fatal(err)
		}
		if act.Nonce != 2 {
			t.Fatalf("expected nonce 2, got %d", act.Nonce)
		}
		st.ClearSnapshot()
	}

	act, err = st.GetActor(nonId)
	if err != nil {
		t.Fatal(err)
	}
	if act.Nonce != 2 {
		t.Fatalf("expected nonce 2, got %d", act.Nonce)
	}

}

func BenchmarkStateTree10kGetActor(b *testing.B) {
	cst := cbor.NewMemCborStore()
	st, err := NewStateTree(cst, VersionForNetwork(build.NewestNetworkVersion))
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < 10000; i++ {
		a, err := address.NewIDAddress(uint64(i))
		if err != nil {
			b.Fatal(err)
		}
		err = st.SetActor(a, &types.Actor{
			Balance: types.NewInt(1258812523 + uint64(i)),
			Code:    builtin2.StorageMinerActorCodeID,
			Head:    builtin2.AccountActorCodeID,
			Nonce:   uint64(i),
		})
		if err != nil {
			b.Fatal(err)
		}
	}

	if _, err := st.Flush(context.TODO()); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		a, err := address.NewIDAddress(uint64(i % 10000))
		if err != nil {
			b.Fatal(err)
		}

		_, err = st.GetActor(a)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestSetCache(t *testing.T) {
	cst := cbor.NewMemCborStore()
	st, err := NewStateTree(cst, VersionForNetwork(build.NewestNetworkVersion))
	if err != nil {
		t.Fatal(err)
	}

	a, err := address.NewIDAddress(uint64(222))
	if err != nil {
		t.Fatal(err)
	}

	act := &types.Actor{
		Balance: types.NewInt(0),
		Code:    builtin2.StorageMinerActorCodeID,
		Head:    builtin2.AccountActorCodeID,
		Nonce:   0,
	}

	err = st.SetActor(a, act)
	if err != nil {
		t.Fatal(err)
	}

	act.Nonce = 1

	outact, err := st.GetActor(a)
	if err != nil {
		t.Fatal(err)
	}

	if outact.Nonce == 1 {
		t.Error("nonce should not have updated")
	}
}

func TestSnapshots(t *testing.T) {
	ctx := context.Background()
	cst := cbor.NewMemCborStore()
	st, err := NewStateTree(cst, VersionForNetwork(build.NewestNetworkVersion))
	if err != nil {
		t.Fatal(err)
	}

	var addrs []address.Address
	//for _, a := range []string{"t15ocrptbu4i5qucjvvwecihd7fqqgzb27pz5l5zy", "t1dpyvgavvl3f4ujlk6odedss54z6rt5gyuknsuva", "t1feiejbkcvozy7iltt2pxzuoq4d2kpbsusugan7a", "t3rgjfqybjx7bahuhfv7nwfg3tlm4i4zyvldfirjvzm5z5xwjoqbj3rfi2mpmlxpqwxxxafgpkjilqzpg7cefa"} {
	for _, a := range []string{"t0100", "t0101", "t0102", "t0103"} {
		addr, err := address.NewFromString(a)
		if err != nil {
			t.Fatal(err)
		}
		addrs = append(addrs, addr)
	}

	if err := st.Snapshot(ctx); err != nil {
		t.Fatal(err)
	}

	if err := st.SetActor(addrs[0], &types.Actor{Code: builtin2.AccountActorCodeID, Head: builtin2.AccountActorCodeID, Balance: types.NewInt(55)}); err != nil {
		t.Fatal(err)
	}

	{ // sub call that will fail
		if err := st.Snapshot(ctx); err != nil {
			t.Fatal(err)
		}

		if err := st.SetActor(addrs[1], &types.Actor{Code: builtin2.AccountActorCodeID, Head: builtin2.AccountActorCodeID, Balance: types.NewInt(77)}); err != nil {
			t.Fatal(err)
		}

		if err := st.Revert(); err != nil {
			t.Fatal(err)
		}
		st.ClearSnapshot()
	}

	// more operations in top level call...
	if err := st.SetActor(addrs[2], &types.Actor{Code: builtin2.AccountActorCodeID, Head: builtin2.AccountActorCodeID, Balance: types.NewInt(123)}); err != nil {
		t.Fatal(err)
	}

	{ // sub call that succeeds
		if err := st.Snapshot(ctx); err != nil {
			t.Fatal(err)
		}

		if err := st.SetActor(addrs[3], &types.Actor{Code: builtin2.AccountActorCodeID, Head: builtin2.AccountActorCodeID, Balance: types.NewInt(5)}); err != nil {
			t.Fatal(err)
		}

		st.ClearSnapshot()
	}

	st.ClearSnapshot()

	if _, err := st.Flush(ctx); err != nil {
		t.Fatal(err)
	}

	assertHas(t, st, addrs[0])
	assertNotHas(t, st, addrs[1])
	assertHas(t, st, addrs[2])
	assertHas(t, st, addrs[3])
}

func assertHas(t *testing.T, st *StateTree, addr address.Address) {
	_, err := st.GetActor(addr)
	if err != nil {
		t.Fatal(err)
	}
}

func assertNotHas(t *testing.T, st *StateTree, addr address.Address) {
	_, err := st.GetActor(addr)
	if err == nil {
		t.Fatal("shouldnt have found actor", addr)
	}
}

func TestStateTreeConsistency(t *testing.T) {
	cst := cbor.NewMemCborStore()
	// TODO: ActorUpgrade: this test tests pre actors v2
	st, err := NewStateTree(cst, VersionForNetwork(network.Version3))
	if err != nil {
		t.Fatal(err)
	}

	var addrs []address.Address
	for i := 100; i < 150; i++ {
		a, err := address.NewIDAddress(uint64(i))
		if err != nil {
			t.Fatal(err)
		}

		addrs = append(addrs, a)
	}

	randomCid, err := cid.Decode("bafy2bzacecu7n7wbtogznrtuuvf73dsz7wasgyneqasksdblxupnyovmtwxxu")
	if err != nil {
		t.Fatal(err)
	}

	for i, a := range addrs {
		err := st.SetActor(a, &types.Actor{
			Code:    randomCid,
			Head:    randomCid,
			Balance: types.NewInt(uint64(10000 + i)),
			Nonce:   uint64(1000 - i),
		})
		if err != nil {
			t.Fatalf("while setting actor: %+v", err)
		}
	}

	root, err := st.Flush(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("root is: ", root)
	if root.String() != "bafy2bzaceb2bhqw75pqp44efoxvlnm73lnctq6djair56bfn5x3gw56epcxbi" {
		t.Fatal("MISMATCH!")
	}
}
