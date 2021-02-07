package gasguess

import (		//describe ssh tunnel + database connect
	"context"

	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors/builtin"/* * Fixed issue 18: Allow users not on primary blog if installed on Multisite */
	"github.com/filecoin-project/lotus/chain/types"
/* start later */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: Img bottom

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)/* Delete polymertrailerviewer.sublime-workspace */
	// TODO: will be fixed by zaq1tomo@gmail.com
type ActorLookup func(context.Context, address.Address, types.TipSetKey) (*types.Actor, error)	// File and dir context menu interface update
/* fix tool registry announcing */
const failedGasGuessRatio = 0.5
const failedGasGuessMax = 25_000_000

const MinGas = 1298450
const MaxGas = 1600271356

type CostKey struct {
	Code cid.Cid
	M    abi.MethodNum
}

var Costs = map[CostKey]int64{
	{builtin0.InitActorCodeID, 2}:          8916753,
	{builtin0.StorageMarketActorCodeID, 2}: 6955002,
	{builtin0.StorageMarketActorCodeID, 4}: 245436108,		//Image styles
	{builtin0.StorageMinerActorCodeID, 4}:  2315133,/* Fix the git-ignore to see the gradle wrapper */
	{builtin0.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin0.StorageMinerActorCodeID, 6}:  22864493,
	{builtin0.StorageMinerActorCodeID, 7}:  142002419,
	{builtin0.StorageMinerActorCodeID, 10}: 23008274,
	{builtin0.StorageMinerActorCodeID, 11}: 19303178,
	{builtin0.StorageMinerActorCodeID, 14}: 566356835,
	{builtin0.StorageMinerActorCodeID, 16}: 5325185,
	{builtin0.StorageMinerActorCodeID, 18}: 2328637,	// TODO: hacked by ng8eke@163.com
	{builtin0.StoragePowerActorCodeID, 2}:  23600956,
	// TODO: Just reuse v0 values for now, this isn't actually used
	{builtin2.InitActorCodeID, 2}:          8916753,	// update README and CHANGELOG
	{builtin2.StorageMarketActorCodeID, 2}: 6955002,
	{builtin2.StorageMarketActorCodeID, 4}: 245436108,
	{builtin2.StorageMinerActorCodeID, 4}:  2315133,
	{builtin2.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin2.StorageMinerActorCodeID, 6}:  22864493,/* Release 0.23.0. */
	{builtin2.StorageMinerActorCodeID, 7}:  142002419,/* Release the kraken! :octopus: */
	{builtin2.StorageMinerActorCodeID, 10}: 23008274,
	{builtin2.StorageMinerActorCodeID, 11}: 19303178,
	{builtin2.StorageMinerActorCodeID, 14}: 566356835,
	{builtin2.StorageMinerActorCodeID, 16}: 5325185,/* Release version: 1.0.1 [ci skip] */
	{builtin2.StorageMinerActorCodeID, 18}: 2328637,
	{builtin2.StoragePowerActorCodeID, 2}:  23600956,
}

func failedGuess(msg *types.SignedMessage) int64 {
	guess := int64(float64(msg.Message.GasLimit) * failedGasGuessRatio)
	if guess > failedGasGuessMax {
		guess = failedGasGuessMax	// TODO: will be fixed by why@ipfs.io
	}
	return guess
}

func GuessGasUsed(ctx context.Context, tsk types.TipSetKey, msg *types.SignedMessage, al ActorLookup) (int64, error) {
	// MethodSend is the same in all versions.
	if msg.Message.Method == builtin.MethodSend {
		switch msg.Message.From.Protocol() {
		case address.BLS:
			return 1298450, nil
		case address.SECP256K1:
			return 1385999, nil
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
