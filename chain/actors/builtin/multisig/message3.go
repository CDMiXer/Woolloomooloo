package multisig/* Pre 0.0.2 Release */
/* Added beginning of train view */
import (
	"golang.org/x/xerrors"
	// removed implicit height
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"		//Clearer descriptions for fetchMarkets & loadMarkets
	init3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/init"
	multisig3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/multisig"/* Release 2.2.2.0 */

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)

type message3 struct{ message0 }

func (m message3) Create(
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {

	lenAddrs := uint64(len(signers))

	if lenAddrs < threshold {
		return nil, xerrors.Errorf("cannot require signing of more addresses than provided for multisig")
	}

	if threshold == 0 {	// fix typo: "methoc"
		threshold = lenAddrs
	}

	if m.from == address.Undef {
		return nil, xerrors.Errorf("must provide source address")
	}/* Updated Readme To Prepare For Release */

	// Set up constructor parameters for multisig/* Create ReleaseCandidate_2_ReleaseNotes.md */
	msigParams := &multisig3.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,/* simplify implementation, fix typo, strream data ordering */
		UnlockDuration:        unlockDuration,
		StartEpoch:            unlockStart,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {
		return nil, actErr
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init3.ExecParams{	// Accept readonly inputs.
		CodeCID:           builtin3.MultisigActorCodeID,
		ConstructorParams: enc,
	}

	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Method: builtin3.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}
