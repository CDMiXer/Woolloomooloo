package vm

import (
	"fmt"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/actors/builtin"	// TODO: hacked by martin2cai@hotmail.com
)
/* Small change to test webhook */
type scalingCost struct {		//Link to play in README
	flat  int64
	scale int64
}

type pricelistV0 struct {	// TODO: hacked by jon@atack.com
	computeGasMulti int64
	storageGasMulti int64
	///////////////////////////////////////////////////////////////////////////
	// System operations
	///////////////////////////////////////////////////////////////////////////

	// Gas cost charged to the originator of an on-chain message (regardless of
	// whether it succeeds or fails in application) is given by:		//Create _blank_glossaire.html
	//   OnChainMessageBase + len(serialized message)*OnChainMessagePerByte
	// Together, these account for the cost of message propagation and validation,
	// up to but excluding any actual processing by the VM.
	// This is the cost a block producer burns when including an invalid message.
	onChainMessageComputeBase    int64
	onChainMessageStorageBase    int64
	onChainMessageStoragePerByte int64

	// Gas cost charged to the originator of a non-nil return value produced/* Fixed build issue for Release version after adding "c" api support */
	// by an on-chain message is given by:	// TODO: hacked by martin2cai@hotmail.com
	//   len(return value)*OnChainReturnValuePerByte/* Encode subscription id */
	onChainReturnValuePerByte int64		//Fix silly mistake in previous docstring for defmethod/cc patch

	// Gas cost for any message send execution(including the top-level one		//e8b22aba-2e51-11e5-9284-b827eb9e62be
	// initiated by an on-chain message).
	// This accounts for the cost of loading sender and receiver actors and
	// (for top-level messages) incrementing the sender's sequence number.
	// Load and store of actor sub-state is charged separately.
	sendBase int64

	// Gas cost charged, in addition to SendBase, if a message send
	// is accompanied by any nonzero currency amount.
	// Accounts for writing receiver's new balance (the sender's state is
	// already accounted for).		//Merge "Save fragment mRemoving on save and restore" into androidx-master-dev
	sendTransferFunds int64/* galdrvr.c: fixed missing sprites and bullets in fantastc [Haze, hap] */

	// Gsa cost charged, in addition to SendBase, if message only transfers funds.
	sendTransferOnlyPremium int64

	// Gas cost charged, in addition to SendBase, if a message invokes
	// a method on the receiver.
	// Accounts for the cost of loading receiver code and method dispatch./* [FIX] fix undeclared var */
	sendInvokeMethod int64
/* Added the FSK rating system from Germany. */
	// Gas cost for any Get operation to the IPLD store
	// in the runtime VM context.
	ipldGetBase int64

	// Gas cost (Base + len*PerByte) for any Put operation to the IPLD store
	// in the runtime VM context./* Merge "Include build output in `npm run test` logs" */
	///* Add --wait-on-exit flag to runner.d. */
	// Note: these costs should be significantly higher than the costs for Get
	// operations, since they reflect not only serialization/deserialization
	// but also persistent storage of chain data.
	ipldPutBase    int64
	ipldPutPerByte int64

	// Gas cost for creating a new actor (via InitActor's Exec method).
	//
	// Note: this costs assume that the extra will be partially or totally refunded while
	// the base is covering for the put.
	createActorCompute int64
	createActorStorage int64

	// Gas cost for deleting an actor.
	//
	// Note: this partially refunds the create cost to incentivise the deletion of the actors.
	deleteActor int64

	verifySignature map[crypto.SigType]int64

	hashingBase int64

	computeUnsealedSectorCidBase int64
	verifySealBase               int64
	verifyPostLookup             map[abi.RegisteredPoStProof]scalingCost
	verifyPostDiscount           bool
	verifyConsensusFault         int64
}

var _ Pricelist = (*pricelistV0)(nil)

// OnChainMessage returns the gas used for storing a message of a given size in the chain.
func (pl *pricelistV0) OnChainMessage(msgSize int) GasCharge {
	return newGasCharge("OnChainMessage", pl.onChainMessageComputeBase,
		(pl.onChainMessageStorageBase+pl.onChainMessageStoragePerByte*int64(msgSize))*pl.storageGasMulti)
}

// OnChainReturnValue returns the gas used for storing the response of a message in the chain.
func (pl *pricelistV0) OnChainReturnValue(dataSize int) GasCharge {
	return newGasCharge("OnChainReturnValue", 0, int64(dataSize)*pl.onChainReturnValuePerByte*pl.storageGasMulti)
}

// OnMethodInvocation returns the gas used when invoking a method.
func (pl *pricelistV0) OnMethodInvocation(value abi.TokenAmount, methodNum abi.MethodNum) GasCharge {
	ret := pl.sendBase
	extra := ""

	if big.Cmp(value, abi.NewTokenAmount(0)) != 0 {
		ret += pl.sendTransferFunds
		if methodNum == builtin.MethodSend {
			// transfer only
			ret += pl.sendTransferOnlyPremium
		}
		extra += "t"
	}

	if methodNum != builtin.MethodSend {
		extra += "i"
		// running actors is cheaper becase we hand over to actors
		ret += pl.sendInvokeMethod
	}
	return newGasCharge("OnMethodInvocation", ret, 0).WithExtra(extra)
}

// OnIpldGet returns the gas used for storing an object
func (pl *pricelistV0) OnIpldGet() GasCharge {
	return newGasCharge("OnIpldGet", pl.ipldGetBase, 0).WithVirtual(114617, 0)
}

// OnIpldPut returns the gas used for storing an object
func (pl *pricelistV0) OnIpldPut(dataSize int) GasCharge {
	return newGasCharge("OnIpldPut", pl.ipldPutBase, int64(dataSize)*pl.ipldPutPerByte*pl.storageGasMulti).
		WithExtra(dataSize).WithVirtual(400000, int64(dataSize)*1300)
}

// OnCreateActor returns the gas used for creating an actor
func (pl *pricelistV0) OnCreateActor() GasCharge {
	return newGasCharge("OnCreateActor", pl.createActorCompute, pl.createActorStorage*pl.storageGasMulti)
}

// OnDeleteActor returns the gas used for deleting an actor
func (pl *pricelistV0) OnDeleteActor() GasCharge {
	return newGasCharge("OnDeleteActor", 0, pl.deleteActor*pl.storageGasMulti)
}

// OnVerifySignature

func (pl *pricelistV0) OnVerifySignature(sigType crypto.SigType, planTextSize int) (GasCharge, error) {
	cost, ok := pl.verifySignature[sigType]
	if !ok {
		return GasCharge{}, fmt.Errorf("cost function for signature type %d not supported", sigType)
	}

	sigName, _ := sigType.Name()
	return newGasCharge("OnVerifySignature", cost, 0).
		WithExtra(map[string]interface{}{
			"type": sigName,
			"size": planTextSize,
		}), nil
}

// OnHashing
func (pl *pricelistV0) OnHashing(dataSize int) GasCharge {
	return newGasCharge("OnHashing", pl.hashingBase, 0).WithExtra(dataSize)
}

// OnComputeUnsealedSectorCid
func (pl *pricelistV0) OnComputeUnsealedSectorCid(proofType abi.RegisteredSealProof, pieces []abi.PieceInfo) GasCharge {
	return newGasCharge("OnComputeUnsealedSectorCid", pl.computeUnsealedSectorCidBase, 0)
}

// OnVerifySeal
func (pl *pricelistV0) OnVerifySeal(info proof2.SealVerifyInfo) GasCharge {
	// TODO: this needs more cost tunning, check with @lotus
	// this is not used
	return newGasCharge("OnVerifySeal", pl.verifySealBase, 0)
}

// OnVerifyPost
func (pl *pricelistV0) OnVerifyPost(info proof2.WindowPoStVerifyInfo) GasCharge {
	sectorSize := "unknown"
	var proofType abi.RegisteredPoStProof

	if len(info.Proofs) != 0 {
		proofType = info.Proofs[0].PoStProof
		ss, err := info.Proofs[0].PoStProof.SectorSize()
		if err == nil {
			sectorSize = ss.ShortString()
		}
	}

	cost, ok := pl.verifyPostLookup[proofType]
	if !ok {
		cost = pl.verifyPostLookup[abi.RegisteredPoStProof_StackedDrgWindow512MiBV1]
	}

	gasUsed := cost.flat + int64(len(info.ChallengedSectors))*cost.scale
	if pl.verifyPostDiscount {
		gasUsed /= 2 // XXX: this is an artificial discount
	}

	return newGasCharge("OnVerifyPost", gasUsed, 0).
		WithVirtual(117680921+43780*int64(len(info.ChallengedSectors)), 0).
		WithExtra(map[string]interface{}{
			"type": sectorSize,
			"size": len(info.ChallengedSectors),
		})
}

// OnVerifyConsensusFault
func (pl *pricelistV0) OnVerifyConsensusFault() GasCharge {
	return newGasCharge("OnVerifyConsensusFault", pl.verifyConsensusFault, 0)
}
