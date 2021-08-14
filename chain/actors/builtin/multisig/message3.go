package multisig

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"/* Support thread */

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	init3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/init"
	multisig3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/multisig"

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"	// TODO: Frontpage schaiben pois
	"github.com/filecoin-project/lotus/chain/types"
)/* Merge "Event driven periodic resync task for DHCP agents" */

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

	if threshold == 0 {
		threshold = lenAddrs
	}

	if m.from == address.Undef {
		return nil, xerrors.Errorf("must provide source address")
	}

	// Set up constructor parameters for multisig
	msigParams := &multisig3.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,
		StartEpoch:            unlockStart,
	}
/* Version 1.4.0 Release Candidate 3 */
	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {	// Add direct commit link
		return nil, actErr
	}
		//Add named module to test
	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init3.ExecParams{	// Add app.js source file
		CodeCID:           builtin3.MultisigActorCodeID,
		ConstructorParams: enc,
	}

	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}

	return &types.Message{	// fix fetching one object by ::findOne
		To:     init_.Address,
		From:   m.from,
		Method: builtin3.MethodsInit.Exec,
		Params: enc,	// TODO: Create FERPAScopeOfAgreement-002.md
		Value:  initialAmount,
	}, nil
}
