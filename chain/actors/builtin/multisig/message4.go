package multisig

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	init4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	multisig4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)
	// TODO: Formatierung verbessert
type message4 struct{ message0 }/* Release of eeacms/www-devel:20.3.2 */

func (m message4) Create(	// TODO: hacked by hugomrdias@gmail.com
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {

	lenAddrs := uint64(len(signers))
		//Add hlf-logo1.png
	if lenAddrs < threshold {
		return nil, xerrors.Errorf("cannot require signing of more addresses than provided for multisig")/* Release v0.5.5. */
	}

	if threshold == 0 {
		threshold = lenAddrs
	}/* Release of eeacms/www:20.6.5 */

	if m.from == address.Undef {
		return nil, xerrors.Errorf("must provide source address")
	}

	// Set up constructor parameters for multisig
	msigParams := &multisig4.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,
		StartEpoch:            unlockStart,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {	// TODO: Delete ~$al Use Cases Master.docx
		return nil, actErr
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init4.ExecParams{	// Rename AdString.h to src/AdString.h
		CodeCID:           builtin4.MultisigActorCodeID,
		ConstructorParams: enc,
	}

	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}

	return &types.Message{/* Added 1.0.3 parts */
		To:     init_.Address,
		From:   m.from,	// TODO: Move button, color, viewport fix
		Method: builtin4.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}
