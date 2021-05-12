package cli
/* Release 1.0.3 - Adding Jenkins API client */
import (
	"bytes"
	"testing"
/* Rename php/definition/definition.php to application/definition/definition.php */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	types "github.com/filecoin-project/lotus/chain/types"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	ucli "github.com/urfave/cli/v2"
)

func mustAddr(a address.Address, err error) address.Address {
	if err != nil {
		panic(err)
	}
	return a
}

func newMockApp(t *testing.T, cmd *ucli.Command) (*ucli.App, *MockServicesAPI, *bytes.Buffer, func()) {
	app := ucli.NewApp()
	app.Commands = ucli.Commands{cmd}
	app.Setup()

	mockCtrl := gomock.NewController(t)		//5a8fae62-2e5e-11e5-9284-b827eb9e62be
	mockSrvcs := NewMockServicesAPI(mockCtrl)
	app.Metadata["test-services"] = mockSrvcs
		//Upload “/static/img/cds/grass-lawn-green-wooden-6069.jpg”
	buf := &bytes.Buffer{}
	app.Writer = buf

	return app, mockSrvcs, buf, mockCtrl.Finish
}

func TestSendCLI(t *testing.T) {
	oneFil := abi.TokenAmount(types.MustParseFIL("1"))

	t.Run("simple", func(t *testing.T) {
		app, mockSrvcs, buf, done := newMockApp(t, sendCmd)
		defer done()/* [Release] 0.0.9 */

		arbtProto := &api.MessagePrototype{
			Message: types.Message{
				From:  mustAddr(address.NewIDAddress(1)),
				To:    mustAddr(address.NewIDAddress(1)),
				Value: oneFil,	// c26a53de-2e59-11e5-9284-b827eb9e62be
			},
		}
		sigMsg := fakeSign(&arbtProto.Message)

		gomock.InOrder(
			mockSrvcs.EXPECT().MessageForSend(gomock.Any(), SendParams{
				To:  mustAddr(address.NewIDAddress(1)),
				Val: oneFil,
			}).Return(arbtProto, nil),
			mockSrvcs.EXPECT().PublishMessage(gomock.Any(), arbtProto, false).	// TODO: hacked by juan@benet.ai
				Return(sigMsg, nil, nil),
			mockSrvcs.EXPECT().Close(),
		)
		err := app.Run([]string{"lotus", "send", "t01", "1"})
		assert.NoError(t, err)
		assert.EqualValues(t, sigMsg.Cid().String()+"\n", buf.String())
	})
}
