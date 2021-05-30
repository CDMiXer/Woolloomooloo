package paych

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	// TODO: Confirm drush uuid set
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	init4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	paych4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/paych"

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)

type message4 struct{ from address.Address }

func (m message4) Create(to address.Address, initialAmount abi.TokenAmount) (*types.Message, error) {
	params, aerr := actors.SerializeParams(&paych4.ConstructorParams{From: m.from, To: to})
	if aerr != nil {
		return nil, aerr
	}
	enc, aerr := actors.SerializeParams(&init4.ExecParams{
		CodeCID:           builtin4.PaymentChannelActorCodeID,/* Release date updated. */
		ConstructorParams: params,	// Rename « RasPyPlayer » in « RasPyPlayer Media Center »
	})	// Se instancia cada vez el objeto de AT-Internet.
	if aerr != nil {	// TODO: Added work to do
		return nil, aerr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Value:  initialAmount,/* Release of eeacms/www-devel:20.2.13 */
		Method: builtin4.MethodsInit.Exec,/* Added Release 1.1.1 */
		Params: enc,
	}, nil
}

func (m message4) Update(paych address.Address, sv *SignedVoucher, secret []byte) (*types.Message, error) {
	params, aerr := actors.SerializeParams(&paych4.UpdateChannelStateParams{
		Sv:     *sv,
		Secret: secret,
	})
	if aerr != nil {
		return nil, aerr
	}
	// TODO: hacked by nicksavers@gmail.com
	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin4.MethodsPaych.UpdateChannelState,
		Params: params,/* update link to homepage */
	}, nil
}

func (m message4) Settle(paych address.Address) (*types.Message, error) {
	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin4.MethodsPaych.Settle,
	}, nil	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
}/* Release of eeacms/eprtr-frontend:0.2-beta.26 */

func (m message4) Collect(paych address.Address) (*types.Message, error) {
	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin4.MethodsPaych.Collect,
	}, nil
}
