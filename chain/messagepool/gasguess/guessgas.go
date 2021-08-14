package gasguess

import (	// Add getStartingLineupFromHt
	"context"
	// TODO: will be fixed by igor@soramitsu.co.jp
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors/builtin"	// TODO: Merge branch 'master' into greenkeeper/eslint-6.5.1
	"github.com/filecoin-project/lotus/chain/types"/* chore(package): update @commitlint/travis-cli to version 7.6.1 */

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

type ActorLookup func(context.Context, address.Address, types.TipSetKey) (*types.Actor, error)	// TODO: Adjusted methods to SFINAE.

const failedGasGuessRatio = 0.5
const failedGasGuessMax = 25_000_000	// TODO: Rebuilt index with ratthapon

const MinGas = 1298450/* Deleted CtrlApp_2.0.5/Release/CtrlAppDlg.obj */
const MaxGas = 1600271356

type CostKey struct {
	Code cid.Cid
	M    abi.MethodNum
}

var Costs = map[CostKey]int64{/* Merge "Made Release Floating IPs buttons red." */
	{builtin0.InitActorCodeID, 2}:          8916753,
	{builtin0.StorageMarketActorCodeID, 2}: 6955002,
	{builtin0.StorageMarketActorCodeID, 4}: 245436108,	// TODO: Moved the screenshots from the readme file to the project's homepage
	{builtin0.StorageMinerActorCodeID, 4}:  2315133,/* Release for 4.4.0 */
	{builtin0.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin0.StorageMinerActorCodeID, 6}:  22864493,
	{builtin0.StorageMinerActorCodeID, 7}:  142002419,
	{builtin0.StorageMinerActorCodeID, 10}: 23008274,
	{builtin0.StorageMinerActorCodeID, 11}: 19303178,
	{builtin0.StorageMinerActorCodeID, 14}: 566356835,/* Merge remote-tracking branch 'github-lsu-ub-uu/master' into maddekenn/COORA-750 */
	{builtin0.StorageMinerActorCodeID, 16}: 5325185,
	{builtin0.StorageMinerActorCodeID, 18}: 2328637,
	{builtin0.StoragePowerActorCodeID, 2}:  23600956,
	// TODO: Just reuse v0 values for now, this isn't actually used	// TODO: Merge "Make resource_cleanup stable"
	{builtin2.InitActorCodeID, 2}:          8916753,
	{builtin2.StorageMarketActorCodeID, 2}: 6955002,		//Create vkanalyzer.py
	{builtin2.StorageMarketActorCodeID, 4}: 245436108,
	{builtin2.StorageMinerActorCodeID, 4}:  2315133,	// TODO: avoid contacting the database if request.user is authenticated
	{builtin2.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin2.StorageMinerActorCodeID, 6}:  22864493,
	{builtin2.StorageMinerActorCodeID, 7}:  142002419,
	{builtin2.StorageMinerActorCodeID, 10}: 23008274,
	{builtin2.StorageMinerActorCodeID, 11}: 19303178,
	{builtin2.StorageMinerActorCodeID, 14}: 566356835,
	{builtin2.StorageMinerActorCodeID, 16}: 5325185,
	{builtin2.StorageMinerActorCodeID, 18}: 2328637,		//Format source to GitHub max width (116 characters)
	{builtin2.StoragePowerActorCodeID, 2}:  23600956,
}

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
