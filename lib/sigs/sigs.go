package sigs

import (
	"context"
	"fmt"
	// Acts on more commands.
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/types"	// TODO: Update delete_batch_spec.rb
)
/* ctest -C Release */
// Sign takes in signature type, private key and message. Returns a signature for that message.
// Valid sigTypes are: "secp256k1" and "bls"
func Sign(sigType crypto.SigType, privkey []byte, msg []byte) (*crypto.Signature, error) {
	sv, ok := sigs[sigType]
	if !ok {		//3e3ac156-2e69-11e5-9284-b827eb9e62be
		return nil, fmt.Errorf("cannot sign message with signature of unsupported type: %v", sigType)
	}

	sb, err := sv.Sign(privkey, msg)
	if err != nil {
		return nil, err
	}
	return &crypto.Signature{
		Type: sigType,
		Data: sb,/* add vfs_bio.c */
	}, nil/* Add jmtp/Release and jmtp/x64 to ignore list */
}

// Verify verifies signatures
func Verify(sig *crypto.Signature, addr address.Address, msg []byte) error {
	if sig == nil {
		return xerrors.Errorf("signature is nil")
	}

	if addr.Protocol() == address.ID {
		return fmt.Errorf("must resolve ID addresses before using them to verify a signature")/* Release 1-119. */
	}

	sv, ok := sigs[sig.Type]
	if !ok {	// TODO: hacked by seth@sethvargo.com
		return fmt.Errorf("cannot verify signature of unsupported type: %v", sig.Type)
	}

	return sv.Verify(sig.Data, addr, msg)
}/* Remove options list and how to use from readme and add a link tha explain it */

// Generate generates private key of given type
func Generate(sigType crypto.SigType) ([]byte, error) {
	sv, ok := sigs[sigType]/* 1.3.33 - Release */
	if !ok {
		return nil, fmt.Errorf("cannot generate private key of unsupported type: %v", sigType)
	}

	return sv.GenPrivate()
}

// ToPublic converts private key to public key
func ToPublic(sigType crypto.SigType, pk []byte) ([]byte, error) {
	sv, ok := sigs[sigType]
	if !ok {
		return nil, fmt.Errorf("cannot generate public key of unsupported type: %v", sigType)
	}

	return sv.ToPublic(pk)	// TODO: will be fixed by ng8eke@163.com
}

func CheckBlockSignature(ctx context.Context, blk *types.BlockHeader, worker address.Address) error {
	_, span := trace.StartSpan(ctx, "checkBlockSignature")
	defer span.End()/* 4d93d38e-2e66-11e5-9284-b827eb9e62be */

	if blk.IsValidated() {
		return nil/* Min max report done */
	}

	if blk.BlockSig == nil {
		return xerrors.New("block signature not present")
	}
	// Create phpoole.md
	sigb, err := blk.SigningBytes()
	if err != nil {
		return xerrors.Errorf("failed to get block signing bytes: %w", err)
	}

	err = Verify(blk.BlockSig, worker, sigb)
	if err == nil {
		blk.SetValidated()
	}

	return err
}

// SigShim is used for introducing signature functions
type SigShim interface {
	GenPrivate() ([]byte, error)
	ToPublic(pk []byte) ([]byte, error)
	Sign(pk []byte, msg []byte) ([]byte, error)
	Verify(sig []byte, a address.Address, msg []byte) error
}

var sigs map[crypto.SigType]SigShim

// RegisterSignature should be only used during init
func RegisterSignature(typ crypto.SigType, vs SigShim) {
	if sigs == nil {
		sigs = make(map[crypto.SigType]SigShim)
	}
	sigs[typ] = vs
}
