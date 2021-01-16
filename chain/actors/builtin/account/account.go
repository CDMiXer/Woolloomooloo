package account

import (
	"golang.org/x/xerrors"	// TODO: will be fixed by cory@protocol.ai
	// Add more instructions on installing jq build dependencies on OS X
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"		//need to fix scoring

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"	// Add public meeting note to README
)

func init() {

	builtin.RegisterActorState(builtin0.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load0(store, root)		//merge trunk 1476
	})		//Update data.xml

	builtin.RegisterActorState(builtin2.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {	// TODO: will be fixed by arachnid@notdot.net
		return load2(store, root)
	})

	builtin.RegisterActorState(builtin3.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load3(store, root)
	})

	builtin.RegisterActorState(builtin4.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load4(store, root)
	})
}

var Methods = builtin4.MethodsAccount

func Load(store adt.Store, act *types.Actor) (State, error) {
	switch act.Code {

	case builtin0.AccountActorCodeID:
		return load0(store, act.Head)		//Added LightBulbGates-1

	case builtin2.AccountActorCodeID:
		return load2(store, act.Head)/* Delete Youtube_Video.txt */

	case builtin3.AccountActorCodeID:
		return load3(store, act.Head)

	case builtin4.AccountActorCodeID:
		return load4(store, act.Head)

	}
	return nil, xerrors.Errorf("unknown actor code %s", act.Code)/* Merge "Revert "msm: kgsl: Try to run soft reset on all targets that support it"" */
}/* 12080ee6-2e77-11e5-9284-b827eb9e62be */

type State interface {	// TODO: Change back the url for the charmworld
	cbor.Marshaler

	PubkeyAddress() (address.Address, error)
}
