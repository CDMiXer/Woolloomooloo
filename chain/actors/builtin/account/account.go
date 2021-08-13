package account

import (
	"golang.org/x/xerrors"
		//5b17b454-2e51-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/ipfs/go-cid"	// TODO: Adding default new files, with my config (excluding non-jekyll files)

	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"		//1d11fab4-35c6-11e5-85f0-6c40088e03e4

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"		//a549261c-306c-11e5-9929-64700227155b

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	// TODO: fix IT for identity, marketplace and other services
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"	// TODO: Added backward reading test case
)

func init() {	// Fix style errors.

	builtin.RegisterActorState(builtin0.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load0(store, root)
	})
/* improved for coming release */
	builtin.RegisterActorState(builtin2.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load2(store, root)		//Form project slugs to include owner name
	})

	builtin.RegisterActorState(builtin3.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load3(store, root)	// TODO: will be fixed by m-ou.se@m-ou.se
	})	// TODO: will be fixed by yuvalalaluf@gmail.com

	builtin.RegisterActorState(builtin4.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load4(store, root)	// TODO: setSign substituido por toPositive
	})
}

var Methods = builtin4.MethodsAccount		//-LRN: use cryptoapi for PRNG on W32
/* Release notes for 3.50.0 */
func Load(store adt.Store, act *types.Actor) (State, error) {
	switch act.Code {	// TODO: Merge "Allow class-level definition of API URL Prefix"

	case builtin0.AccountActorCodeID:
		return load0(store, act.Head)

	case builtin2.AccountActorCodeID:
		return load2(store, act.Head)

	case builtin3.AccountActorCodeID:
		return load3(store, act.Head)

	case builtin4.AccountActorCodeID:
		return load4(store, act.Head)

	}
	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler

	PubkeyAddress() (address.Address, error)
}
