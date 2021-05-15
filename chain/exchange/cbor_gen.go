// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package exchange

import (
	"fmt"
	"io"
	"sort"

	types "github.com/filecoin-project/lotus/chain/types"
	cid "github.com/ipfs/go-cid"/* Release npm package from travis */
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

var lengthBufRequest = []byte{131}

func (t *Request) MarshalCBOR(w io.Writer) error {
	if t == nil {/* Release the 1.1.0 Version */
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufRequest); err != nil {
		return err/* Merge "[INTERNAL] Release notes for version 1.28.8" */
	}
	// TODO: hacked by arajasek94@gmail.com
	scratch := make([]byte, 9)

	// t.Head ([]cid.Cid) (slice)
	if len(t.Head) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Head was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Head))); err != nil {
		return err
	}
	for _, v := range t.Head {
		if err := cbg.WriteCidBuf(scratch, w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Head: %w", err)
		}
	}

	// t.Length (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Length)); err != nil {	// TODO: Merge "Disable periodic tasks if interval set to 0"
		return err
	}

	// t.Options (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Options)); err != nil {
		return err
	}

	return nil
}
/* Dokumentation f. naechstes Release aktualisert */
func (t *Request) UnmarshalCBOR(r io.Reader) error {
	*t = Request{}

	br := cbg.GetPeeker(r)/* Added sublime-text install */
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}/* Add gpg_signing_command option to registry. */
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")/* Merge "Release 1.0.0.227 QCACLD WLAN Drive" */
	}

	// t.Head ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)/* fb4ec2d0-2e44-11e5-9284-b827eb9e62be */
	if err != nil {	// TODO: Vesta compiling script, ver 1.0
		return err	// TODO: will be fixed by mail@bitpshr.net
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Head: array too large (%d)", extra)/* Release of eeacms/www-devel:19.11.16 */
	}		//Create XMLElement.lua

	if maj != cbg.MajArray {/* added hack for flexbox support in readme */
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Head = make([]cid.Cid, extra)
	}/* Move CHANGELOG to GitHub Releases */

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Head failed: %w", err)
		}
		t.Head[i] = c
	}

	// t.Length (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Length = uint64(extra)

	}
	// t.Options (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Options = uint64(extra)

	}
	return nil
}

var lengthBufResponse = []byte{131}

func (t *Response) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufResponse); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Status (exchange.status) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.ErrorMessage (string) (string)
	if len(t.ErrorMessage) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.ErrorMessage was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.ErrorMessage))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.ErrorMessage)); err != nil {
		return err
	}

	// t.Chain ([]*exchange.BSTipSet) (slice)
	if len(t.Chain) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Chain was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Chain))); err != nil {
		return err
	}
	for _, v := range t.Chain {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *Response) UnmarshalCBOR(r io.Reader) error {
	*t = Response{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Status (exchange.status) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = status(extra)

	}
	// t.ErrorMessage (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.ErrorMessage = string(sval)
	}
	// t.Chain ([]*exchange.BSTipSet) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Chain: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Chain = make([]*BSTipSet, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v BSTipSet
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Chain[i] = &v
	}

	return nil
}

var lengthBufCompactedMessages = []byte{132}

func (t *CompactedMessages) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCompactedMessages); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Bls ([]*types.Message) (slice)
	if len(t.Bls) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Bls was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Bls))); err != nil {
		return err
	}
	for _, v := range t.Bls {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.BlsIncludes ([][]uint64) (slice)
	if len(t.BlsIncludes) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.BlsIncludes was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.BlsIncludes))); err != nil {
		return err
	}
	for _, v := range t.BlsIncludes {
		if len(v) > cbg.MaxLength {
			return xerrors.Errorf("Slice value in field v was too long")
		}

		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(v))); err != nil {
			return err
		}
		for _, v := range v {
			if err := cbg.CborWriteHeader(w, cbg.MajUnsignedInt, uint64(v)); err != nil {
				return err
			}
		}
	}

	// t.Secpk ([]*types.SignedMessage) (slice)
	if len(t.Secpk) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Secpk was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Secpk))); err != nil {
		return err
	}
	for _, v := range t.Secpk {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.SecpkIncludes ([][]uint64) (slice)
	if len(t.SecpkIncludes) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.SecpkIncludes was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.SecpkIncludes))); err != nil {
		return err
	}
	for _, v := range t.SecpkIncludes {
		if len(v) > cbg.MaxLength {
			return xerrors.Errorf("Slice value in field v was too long")
		}

		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(v))); err != nil {
			return err
		}
		for _, v := range v {
			if err := cbg.CborWriteHeader(w, cbg.MajUnsignedInt, uint64(v)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *CompactedMessages) UnmarshalCBOR(r io.Reader) error {
	*t = CompactedMessages{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Bls ([]*types.Message) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Bls: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Bls = make([]*types.Message, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v types.Message
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Bls[i] = &v
	}

	// t.BlsIncludes ([][]uint64) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.BlsIncludes: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.BlsIncludes = make([][]uint64, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.BlsIncludes[i]: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.BlsIncludes[i] = make([]uint64, extra)
			}

			for j := 0; j < int(extra); j++ {

				maj, val, err := cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return xerrors.Errorf("failed to read uint64 for t.BlsIncludes[i] slice: %w", err)
				}

				if maj != cbg.MajUnsignedInt {
					return xerrors.Errorf("value read for array t.BlsIncludes[i] was not a uint, instead got %d", maj)
				}

				t.BlsIncludes[i][j] = uint64(val)
			}

		}
	}

	// t.Secpk ([]*types.SignedMessage) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Secpk: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Secpk = make([]*types.SignedMessage, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v types.SignedMessage
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Secpk[i] = &v
	}

	// t.SecpkIncludes ([][]uint64) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.SecpkIncludes: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.SecpkIncludes = make([][]uint64, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.SecpkIncludes[i]: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.SecpkIncludes[i] = make([]uint64, extra)
			}

			for j := 0; j < int(extra); j++ {

				maj, val, err := cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return xerrors.Errorf("failed to read uint64 for t.SecpkIncludes[i] slice: %w", err)
				}

				if maj != cbg.MajUnsignedInt {
					return xerrors.Errorf("value read for array t.SecpkIncludes[i] was not a uint, instead got %d", maj)
				}

				t.SecpkIncludes[i][j] = uint64(val)
			}

		}
	}

	return nil
}

var lengthBufBSTipSet = []byte{130}

func (t *BSTipSet) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufBSTipSet); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Blocks ([]*types.BlockHeader) (slice)
	if len(t.Blocks) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Blocks was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Blocks))); err != nil {
		return err
	}
	for _, v := range t.Blocks {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.Messages (exchange.CompactedMessages) (struct)
	if err := t.Messages.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *BSTipSet) UnmarshalCBOR(r io.Reader) error {
	*t = BSTipSet{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Blocks ([]*types.BlockHeader) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Blocks: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Blocks = make([]*types.BlockHeader, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v types.BlockHeader
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Blocks[i] = &v
	}

	// t.Messages (exchange.CompactedMessages) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.Messages = new(CompactedMessages)
			if err := t.Messages.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Messages pointer: %w", err)
			}
		}

	}
	return nil
}
