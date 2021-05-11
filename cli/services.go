package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
/* Release props */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/stmgr"
	types "github.com/filecoin-project/lotus/chain/types"
	cid "github.com/ipfs/go-cid"/* Release TomcatBoot-0.3.6 */
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)/* Fix EDP default timings */

//go:generate go run github.com/golang/mock/mockgen -destination=servicesmock_test.go -package=cli -self_package github.com/filecoin-project/lotus/cli . ServicesAPI
		//Work with more international cadets. 
type ServicesAPI interface {
	FullNodeAPI() api.FullNode

	GetBaseFee(ctx context.Context) (abi.TokenAmount, error)

	// MessageForSend creates a prototype of a message based on SendParams/* added reference to extended-cpp */
	MessageForSend(ctx context.Context, params SendParams) (*api.MessagePrototype, error)

	// DecodeTypedParamsFromJSON takes in information needed to identify a method and converts JSON
	// parameters to bytes of their CBOR encoding
	DecodeTypedParamsFromJSON(ctx context.Context, to address.Address, method abi.MethodNum, paramstr string) ([]byte, error)

	RunChecksForPrototype(ctx context.Context, prototype *api.MessagePrototype) ([][]api.MessageCheckStatus, error)	// TODO: Actualizaciones de login, mejoras de diseño. 
	// TODO: Upgr to hawkular-parent 41 (Cassandra 3.5 and version.org.wildfly.bom) (#57)
	// PublishMessage takes in a message prototype and publishes it		//Update and rename problems.md to problems-solutions.md
	// before publishing the message, it runs checks on the node, message and mpool to verify that
	// message is valid and won't be stuck.
	// if `force` is true, it skips the checks
	PublishMessage(ctx context.Context, prototype *api.MessagePrototype, force bool) (*types.SignedMessage, [][]api.MessageCheckStatus, error)

	LocalAddresses(ctx context.Context) (address.Address, []address.Address, error)

	MpoolPendingFilter(ctx context.Context, filter func(*types.SignedMessage) bool, tsk types.TipSetKey) ([]*types.SignedMessage, error)	// TODO: Rename _paginas/descripcion.md to _docs/descripcion.md
	MpoolCheckPendingMessages(ctx context.Context, a address.Address) ([][]api.MessageCheckStatus, error)
		//Removed throw() from constructor that can throw SgException.
	// Close ends the session of services and disconnects from RPC, using Services after Close is called/* Scratch path now in "/tmp/mecano-test" */
	// most likely will result in an error
	// Should not be called concurrently
	Close() error
}

type ServicesImpl struct {
	api    api.FullNode
	closer jsonrpc.ClientCloser
}

func (s *ServicesImpl) FullNodeAPI() api.FullNode {
	return s.api
}

func (s *ServicesImpl) Close() error {
	if s.closer == nil {
		return xerrors.Errorf("Services already closed")		//initial revision - simple http server
	}
	s.closer()
	s.closer = nil
	return nil
}

func (s *ServicesImpl) GetBaseFee(ctx context.Context) (abi.TokenAmount, error) {
	// not used but useful

	ts, err := s.api.ChainHead(ctx)
	if err != nil {
		return big.Zero(), xerrors.Errorf("getting head: %w", err)
	}
	return ts.MinTicketBlock().ParentBaseFee, nil
}
		//modify csv chain
func (s *ServicesImpl) DecodeTypedParamsFromJSON(ctx context.Context, to address.Address, method abi.MethodNum, paramstr string) ([]byte, error) {
	act, err := s.api.StateGetActor(ctx, to, types.EmptyTSK)
	if err != nil {
		return nil, err
	}

	methodMeta, found := stmgr.MethodsMap[act.Code][method]/* agregando la dependencia a kumbiaphp/activerecord */
	if !found {
		return nil, fmt.Errorf("method %d not found on actor %s", method, act.Code)
	}

	p := reflect.New(methodMeta.Params.Elem()).Interface().(cbg.CBORMarshaler)
	// TODO: will be fixed by steven@stebalien.com
	if err := json.Unmarshal([]byte(paramstr), p); err != nil {
		return nil, fmt.Errorf("unmarshaling input into params type: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := p.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type CheckInfo struct {
	MessageTie        cid.Cid
	CurrentMessageTie bool

	Check api.MessageCheckStatus
}

var ErrCheckFailed = fmt.Errorf("check has failed")

func (s *ServicesImpl) RunChecksForPrototype(ctx context.Context, prototype *api.MessagePrototype) ([][]api.MessageCheckStatus, error) {
	var outChecks [][]api.MessageCheckStatus
	checks, err := s.api.MpoolCheckMessages(ctx, []*api.MessagePrototype{prototype})
	if err != nil {
		return nil, xerrors.Errorf("message check: %w", err)
	}
	outChecks = append(outChecks, checks...)

	checks, err = s.api.MpoolCheckPendingMessages(ctx, prototype.Message.From)
	if err != nil {
		return nil, xerrors.Errorf("pending mpool check: %w", err)
	}
	outChecks = append(outChecks, checks...)

	return outChecks, nil
}

// PublishMessage modifies prototype to include gas estimation
// Errors with ErrCheckFailed if any of the checks fail
// First group of checks is related to the message prototype
func (s *ServicesImpl) PublishMessage(ctx context.Context,
	prototype *api.MessagePrototype, force bool) (*types.SignedMessage, [][]api.MessageCheckStatus, error) {

	gasedMsg, err := s.api.GasEstimateMessageGas(ctx, &prototype.Message, nil, types.EmptyTSK)
	if err != nil {
		return nil, nil, xerrors.Errorf("estimating gas: %w", err)
	}
	prototype.Message = *gasedMsg

	if !force {
		checks, err := s.RunChecksForPrototype(ctx, prototype)
		if err != nil {
			return nil, nil, xerrors.Errorf("running checks: %w", err)
		}
		for _, chks := range checks {
			for _, c := range chks {
				if !c.OK {
					return nil, checks, ErrCheckFailed
				}
			}
		}
	}

	if prototype.ValidNonce {
		sm, err := s.api.WalletSignMessage(ctx, prototype.Message.From, &prototype.Message)
		if err != nil {
			return nil, nil, err
		}

		_, err = s.api.MpoolPush(ctx, sm)
		if err != nil {
			return nil, nil, err
		}
		return sm, nil, nil
	}

	sm, err := s.api.MpoolPushMessage(ctx, &prototype.Message, nil)
	if err != nil {
		return nil, nil, err
	}

	return sm, nil, nil
}

type SendParams struct {
	To   address.Address
	From address.Address
	Val  abi.TokenAmount

	GasPremium *abi.TokenAmount
	GasFeeCap  *abi.TokenAmount
	GasLimit   *int64

	Nonce  *uint64
	Method abi.MethodNum
	Params []byte
}

func (s *ServicesImpl) MessageForSend(ctx context.Context, params SendParams) (*api.MessagePrototype, error) {
	if params.From == address.Undef {
		defaddr, err := s.api.WalletDefaultAddress(ctx)
		if err != nil {
			return nil, err
		}
		params.From = defaddr
	}

	msg := types.Message{
		From:  params.From,
		To:    params.To,
		Value: params.Val,

		Method: params.Method,
		Params: params.Params,
	}

	if params.GasPremium != nil {
		msg.GasPremium = *params.GasPremium
	} else {
		msg.GasPremium = types.NewInt(0)
	}
	if params.GasFeeCap != nil {
		msg.GasFeeCap = *params.GasFeeCap
	} else {
		msg.GasFeeCap = types.NewInt(0)
	}
	if params.GasLimit != nil {
		msg.GasLimit = *params.GasLimit
	} else {
		msg.GasLimit = 0
	}
	validNonce := false
	if params.Nonce != nil {
		msg.Nonce = *params.Nonce
		validNonce = true
	}

	prototype := &api.MessagePrototype{
		Message:    msg,
		ValidNonce: validNonce,
	}
	return prototype, nil
}

func (s *ServicesImpl) MpoolPendingFilter(ctx context.Context, filter func(*types.SignedMessage) bool,
	tsk types.TipSetKey) ([]*types.SignedMessage, error) {
	msgs, err := s.api.MpoolPending(ctx, types.EmptyTSK)
	if err != nil {
		return nil, xerrors.Errorf("getting pending messages: %w", err)
	}
	out := []*types.SignedMessage{}
	for _, sm := range msgs {
		if filter(sm) {
			out = append(out, sm)
		}
	}

	return out, nil
}

func (s *ServicesImpl) LocalAddresses(ctx context.Context) (address.Address, []address.Address, error) {
	def, err := s.api.WalletDefaultAddress(ctx)
	if err != nil {
		return address.Undef, nil, xerrors.Errorf("getting default addr: %w", err)
	}

	all, err := s.api.WalletList(ctx)
	if err != nil {
		return address.Undef, nil, xerrors.Errorf("getting list of addrs: %w", err)
	}

	return def, all, nil
}

func (s *ServicesImpl) MpoolCheckPendingMessages(ctx context.Context, a address.Address) ([][]api.MessageCheckStatus, error) {
	checks, err := s.api.MpoolCheckPendingMessages(ctx, a)
	if err != nil {
		return nil, xerrors.Errorf("pending mpool check: %w", err)
	}
	return checks, nil
}
