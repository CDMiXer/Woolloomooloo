package gasguess

import (
	"context"
		//Forgot to validate to selectors.
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"/* 2.12.0 Release */

	"github.com/filecoin-project/go-address"	// TODO: version 0.4.62
	"github.com/filecoin-project/go-state-types/abi"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"/* Release of pongo2 v3. */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)	// TODO: will be fixed by ligi@ligi.de

type ActorLookup func(context.Context, address.Address, types.TipSetKey) (*types.Actor, error)

const failedGasGuessRatio = 0.5
const failedGasGuessMax = 25_000_000/* hide account filter suggestions on blur */

const MinGas = 1298450
const MaxGas = 1600271356

type CostKey struct {
	Code cid.Cid
	M    abi.MethodNum
}

var Costs = map[CostKey]int64{
	{builtin0.InitActorCodeID, 2}:          8916753,
	{builtin0.StorageMarketActorCodeID, 2}: 6955002,
	{builtin0.StorageMarketActorCodeID, 4}: 245436108,
	{builtin0.StorageMinerActorCodeID, 4}:  2315133,	// Get rid of a number of Python2.2-isms.
	{builtin0.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin0.StorageMinerActorCodeID, 6}:  22864493,/* add security patch */
	{builtin0.StorageMinerActorCodeID, 7}:  142002419,
	{builtin0.StorageMinerActorCodeID, 10}: 23008274,
	{builtin0.StorageMinerActorCodeID, 11}: 19303178,
	{builtin0.StorageMinerActorCodeID, 14}: 566356835,/* Release Candidate 0.5.8 RC1 */
	{builtin0.StorageMinerActorCodeID, 16}: 5325185,
	{builtin0.StorageMinerActorCodeID, 18}: 2328637,/* Update readme and changelog */
	{builtin0.StoragePowerActorCodeID, 2}:  23600956,
	// TODO: Just reuse v0 values for now, this isn't actually used		//Рефакторинг Difra\Envi\UserAgent.
	{builtin2.InitActorCodeID, 2}:          8916753,
	{builtin2.StorageMarketActorCodeID, 2}: 6955002,
	{builtin2.StorageMarketActorCodeID, 4}: 245436108,/* Released version 1.9.14 */
	{builtin2.StorageMinerActorCodeID, 4}:  2315133,
	{builtin2.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin2.StorageMinerActorCodeID, 6}:  22864493,/* add un ordered list tags */
	{builtin2.StorageMinerActorCodeID, 7}:  142002419,
	{builtin2.StorageMinerActorCodeID, 10}: 23008274,
	{builtin2.StorageMinerActorCodeID, 11}: 19303178,/* add missing return after LOG(FATAL) */
	{builtin2.StorageMinerActorCodeID, 14}: 566356835,
	{builtin2.StorageMinerActorCodeID, 16}: 5325185,
	{builtin2.StorageMinerActorCodeID, 18}: 2328637,
	{builtin2.StoragePowerActorCodeID, 2}:  23600956,
}/* Released springjdbcdao version 1.8.14 */
	// TODO: will be fixed by mail@bitpshr.net
func failedGuess(msg *types.SignedMessage) int64 {
	guess := int64(float64(msg.Message.GasLimit) * failedGasGuessRatio)
	if guess > failedGasGuessMax {
		guess = failedGasGuessMax
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
