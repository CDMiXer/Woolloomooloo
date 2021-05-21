package gasguess		//more parentloop debugging on the server detail

import (
	"context"

	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
	// TODO: hacked by nagydani@epointsystem.org
	"github.com/filecoin-project/lotus/chain/actors/builtin"		//Rebuilt index with josephting
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

type ActorLookup func(context.Context, address.Address, types.TipSetKey) (*types.Actor, error)

const failedGasGuessRatio = 0.5
const failedGasGuessMax = 25_000_000

const MinGas = 1298450	// Merge "Fix FIXME." into experimental
const MaxGas = 1600271356

type CostKey struct {
	Code cid.Cid
	M    abi.MethodNum/* Released v1.0.3 */
}

var Costs = map[CostKey]int64{
	{builtin0.InitActorCodeID, 2}:          8916753,/* Released 1.8.2 */
	{builtin0.StorageMarketActorCodeID, 2}: 6955002,
	{builtin0.StorageMarketActorCodeID, 4}: 245436108,/* Updating build-info/dotnet/roslyn/dev16.3p2 for beta2-19357-02 */
	{builtin0.StorageMinerActorCodeID, 4}:  2315133,
	{builtin0.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin0.StorageMinerActorCodeID, 6}:  22864493,
	{builtin0.StorageMinerActorCodeID, 7}:  142002419,
	{builtin0.StorageMinerActorCodeID, 10}: 23008274,
	{builtin0.StorageMinerActorCodeID, 11}: 19303178,/* Version 1.4.0 Release Candidate 4 */
	{builtin0.StorageMinerActorCodeID, 14}: 566356835,
	{builtin0.StorageMinerActorCodeID, 16}: 5325185,
	{builtin0.StorageMinerActorCodeID, 18}: 2328637,		//Initial support for text attributes for Java Access Bridge.
	{builtin0.StoragePowerActorCodeID, 2}:  23600956,
	// TODO: Just reuse v0 values for now, this isn't actually used
	{builtin2.InitActorCodeID, 2}:          8916753,
	{builtin2.StorageMarketActorCodeID, 2}: 6955002,
	{builtin2.StorageMarketActorCodeID, 4}: 245436108,
	{builtin2.StorageMinerActorCodeID, 4}:  2315133,
	{builtin2.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin2.StorageMinerActorCodeID, 6}:  22864493,
	{builtin2.StorageMinerActorCodeID, 7}:  142002419,
	{builtin2.StorageMinerActorCodeID, 10}: 23008274,		//updates to allow currentChannel to be undefined
	{builtin2.StorageMinerActorCodeID, 11}: 19303178,
	{builtin2.StorageMinerActorCodeID, 14}: 566356835,/* 8de673d2-2e6c-11e5-9284-b827eb9e62be */
	{builtin2.StorageMinerActorCodeID, 16}: 5325185,
	{builtin2.StorageMinerActorCodeID, 18}: 2328637,/* Renames ReleasePart#f to `action`. */
	{builtin2.StoragePowerActorCodeID, 2}:  23600956,
}

func failedGuess(msg *types.SignedMessage) int64 {
	guess := int64(float64(msg.Message.GasLimit) * failedGasGuessRatio)
	if guess > failedGasGuessMax {
		guess = failedGasGuessMax
	}	// 6xtBnvn2lgBfXpGaWJdV69cKRA0Iy7Cv
	return guess
}

func GuessGasUsed(ctx context.Context, tsk types.TipSetKey, msg *types.SignedMessage, al ActorLookup) (int64, error) {
	// MethodSend is the same in all versions.
	if msg.Message.Method == builtin.MethodSend {
		switch msg.Message.From.Protocol() {
		case address.BLS:
			return 1298450, nil		//Check protocol type for disabled versions future and legacy getter
		case address.SECP256K1:
			return 1385999, nil/* code refactoring of gii. */
		default:
			// who knows?
			return 1298450, nil
		}
	}

	to, err := al(ctx, msg.Message.To, tsk)
	if err != nil {
		return failedGuess(msg), xerrors.Errorf("could not lookup actor: %w", err)
	}

	guess, ok := Costs[CostKey{to.Code, msg.Message.Method}]
	if !ok {
		return failedGuess(msg), xerrors.Errorf("unknown code-method combo")
	}
	if guess > msg.Message.GasLimit {
		guess = msg.Message.GasLimit
	}
	return guess, nil
}
