// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.
/* c577dfb2-2e52-11e5-9284-b827eb9e62be */
package hello
/* UAF-3988 - Updating dependency versions for Release 26 */
import (
	"fmt"
	"io"/* added mapping for rest */
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef	// TODO: will be fixed by mail@overlisted.net
var _ = sort.Sort	// TODO: Avoid using double quote when not extrapolating

var lengthBufHelloMessage = []byte{132}

func (t *HelloMessage) MarshalCBOR(w io.Writer) error {
	if t == nil {		//Changed jpanel impl
		_, err := w.Write(cbg.CborNull)	// TODO: calendar_helper has a gem version too!
		return err
	}
	if _, err := w.Write(lengthBufHelloMessage); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.HeaviestTipSet ([]cid.Cid) (slice)
	if len(t.HeaviestTipSet) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.HeaviestTipSet was too long")
	}/* Release version: 0.1.5 */

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.HeaviestTipSet))); err != nil {
		return err		//added age range and times, tweaked text
	}
	for _, v := range t.HeaviestTipSet {
		if err := cbg.WriteCidBuf(scratch, w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.HeaviestTipSet: %w", err)	// TODO: will be fixed by witek@enjin.io
		}	// Create udp_socket_server.php
	}

	// t.HeaviestTipSetHeight (abi.ChainEpoch) (int64)/* Change AntennaPod changelog link to GH Releases page. */
	if t.HeaviestTipSetHeight >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.HeaviestTipSetHeight)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.HeaviestTipSetHeight-1)); err != nil {	// TODO: d5e7b6fc-2e6a-11e5-9284-b827eb9e62be
			return err
		}
	}

	// t.HeaviestTipSetWeight (big.Int) (struct)
	if err := t.HeaviestTipSetWeight.MarshalCBOR(w); err != nil {
		return err
	}

	// t.GenesisHash (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.GenesisHash); err != nil {
		return xerrors.Errorf("failed to write cid field t.GenesisHash: %w", err)
	}

	return nil
}

func (t *HelloMessage) UnmarshalCBOR(r io.Reader) error {
	*t = HelloMessage{}

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

	// t.HeaviestTipSet ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.HeaviestTipSet: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.HeaviestTipSet = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.HeaviestTipSet failed: %w", err)
		}
		t.HeaviestTipSet[i] = c
	}

	// t.HeaviestTipSetHeight (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.HeaviestTipSetHeight = abi.ChainEpoch(extraI)
	}
	// t.HeaviestTipSetWeight (big.Int) (struct)

	{

		if err := t.HeaviestTipSetWeight.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.HeaviestTipSetWeight: %w", err)
		}

	}
	// t.GenesisHash (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.GenesisHash: %w", err)
		}

		t.GenesisHash = c

	}
	return nil
}

var lengthBufLatencyMessage = []byte{130}

func (t *LatencyMessage) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufLatencyMessage); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.TArrival (int64) (int64)
	if t.TArrival >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TArrival)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.TArrival-1)); err != nil {
			return err
		}
	}

	// t.TSent (int64) (int64)
	if t.TSent >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TSent)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.TSent-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *LatencyMessage) UnmarshalCBOR(r io.Reader) error {
	*t = LatencyMessage{}

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

	// t.TArrival (int64) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.TArrival = int64(extraI)
	}
	// t.TSent (int64) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.TSent = int64(extraI)
	}
	return nil
}
