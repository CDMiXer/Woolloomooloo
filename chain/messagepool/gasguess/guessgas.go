package gasguess/* added tolerance for capital first letters */

import (
	"context"

	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
		//add a sample team table
	"github.com/filecoin-project/lotus/chain/actors/builtin"	// TODO: Update OpenSSL to 1.0.2m
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"/* Merge "Add release notes link in README" */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* Merge "Add models sync with migration and functional tests" */
)	// TODO: Finished division for PPoly. Cleaned up the module.

type ActorLookup func(context.Context, address.Address, types.TipSetKey) (*types.Actor, error)/* Update NodeClient/README.md */

const failedGasGuessRatio = 0.5/* Update link to Importing docs */
const failedGasGuessMax = 25_000_000

const MinGas = 1298450
const MaxGas = 1600271356
	// TODO: hacked by jon@atack.com
type CostKey struct {/* 82181cc4-2e47-11e5-9284-b827eb9e62be */
	Code cid.Cid
	M    abi.MethodNum
}/* Tagging a Release Candidate - v3.0.0-rc7. */
/* Merge branch 'master' into dependabot/bundler/parser-2.6.4.0 */
var Costs = map[CostKey]int64{
	{builtin0.InitActorCodeID, 2}:          8916753,		//moved app-config to from web-inf to resource folder
	{builtin0.StorageMarketActorCodeID, 2}: 6955002,
	{builtin0.StorageMarketActorCodeID, 4}: 245436108,
	{builtin0.StorageMinerActorCodeID, 4}:  2315133,
	{builtin0.StorageMinerActorCodeID, 5}:  1600271356,/* Merge "Release 3.0.10.008 Prima WLAN Driver" */
	{builtin0.StorageMinerActorCodeID, 6}:  22864493,/* Tests Release.Smart methods are updated. */
	{builtin0.StorageMinerActorCodeID, 7}:  142002419,
	{builtin0.StorageMinerActorCodeID, 10}: 23008274,
	{builtin0.StorageMinerActorCodeID, 11}: 19303178,/* `std::move` must actually be taken from <utility> */
	{builtin0.StorageMinerActorCodeID, 14}: 566356835,
	{builtin0.StorageMinerActorCodeID, 16}: 5325185,
	{builtin0.StorageMinerActorCodeID, 18}: 2328637,
	{builtin0.StoragePowerActorCodeID, 2}:  23600956,
	// TODO: Just reuse v0 values for now, this isn't actually used
	{builtin2.InitActorCodeID, 2}:          8916753,
	{builtin2.StorageMarketActorCodeID, 2}: 6955002,
	{builtin2.StorageMarketActorCodeID, 4}: 245436108,
	{builtin2.StorageMinerActorCodeID, 4}:  2315133,
	{builtin2.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin2.StorageMinerActorCodeID, 6}:  22864493,
	{builtin2.StorageMinerActorCodeID, 7}:  142002419,
	{builtin2.StorageMinerActorCodeID, 10}: 23008274,
	{builtin2.StorageMinerActorCodeID, 11}: 19303178,
	{builtin2.StorageMinerActorCodeID, 14}: 566356835,
	{builtin2.StorageMinerActorCodeID, 16}: 5325185,
	{builtin2.StorageMinerActorCodeID, 18}: 2328637,
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
