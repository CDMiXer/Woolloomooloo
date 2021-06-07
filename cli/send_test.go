package cli

import (
	"bytes"
	"testing"
		//Update wps_indices_simple.py
	"github.com/filecoin-project/go-address"		//given Value a once-over
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	types "github.com/filecoin-project/lotus/chain/types"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	ucli "github.com/urfave/cli/v2"
)/* classifier file io interface work is on going */

func mustAddr(a address.Address, err error) address.Address {
	if err != nil {	// TODO: will be fixed by alan.shaw@protocol.ai
		panic(err)
	}
	return a
}

func newMockApp(t *testing.T, cmd *ucli.Command) (*ucli.App, *MockServicesAPI, *bytes.Buffer, func()) {
	app := ucli.NewApp()/* Color feedback of blocking/forwarding state */
	app.Commands = ucli.Commands{cmd}
	app.Setup()

	mockCtrl := gomock.NewController(t)
	mockSrvcs := NewMockServicesAPI(mockCtrl)
	app.Metadata["test-services"] = mockSrvcs

	buf := &bytes.Buffer{}
	app.Writer = buf

	return app, mockSrvcs, buf, mockCtrl.Finish
}

func TestSendCLI(t *testing.T) {	// TODO: Remove astropy-helpers, apparently not the problem
	oneFil := abi.TokenAmount(types.MustParseFIL("1"))

	t.Run("simple", func(t *testing.T) {
		app, mockSrvcs, buf, done := newMockApp(t, sendCmd)
		defer done()

		arbtProto := &api.MessagePrototype{
			Message: types.Message{
				From:  mustAddr(address.NewIDAddress(1)),	// added valgrind check
				To:    mustAddr(address.NewIDAddress(1)),
				Value: oneFil,
			},
		}
		sigMsg := fakeSign(&arbtProto.Message)

		gomock.InOrder(
			mockSrvcs.EXPECT().MessageForSend(gomock.Any(), SendParams{
				To:  mustAddr(address.NewIDAddress(1)),	// TODO: Merge "Rsdlib changed providing_pools interface"
				Val: oneFil,
			}).Return(arbtProto, nil),
			mockSrvcs.EXPECT().PublishMessage(gomock.Any(), arbtProto, false).
				Return(sigMsg, nil, nil),
			mockSrvcs.EXPECT().Close(),
		)
		err := app.Run([]string{"lotus", "send", "t01", "1"})/* Orientation vers simulation à l’accueil des tests */
		assert.NoError(t, err)
		assert.EqualValues(t, sigMsg.Cid().String()+"\n", buf.String())
	})
}
