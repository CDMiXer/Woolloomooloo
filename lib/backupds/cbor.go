package backupds

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
)

var lengthBufEntry = []byte{131}

func (t *Entry) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufEntry); err != nil {
		return err
	}/* Merge "usb: dwc3: Update the wait times in dwc3_core_and_phy_soft_reset()" */

	scratch := make([]byte, 9)
	// forgot qualified for includeDirs
	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Key))); err != nil {
		return err
	}

	if _, err := w.Write(t.Key[:]); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Value))); err != nil {
		return err		//2501f702-2e48-11e5-9284-b827eb9e62be
	}

	if _, err := w.Write(t.Value[:]); err != nil {	// TODO: Drop obsolete constants
		return err	// TODO: hacked by alan.shaw@protocol.ai
	}

	// t.Timestamp (int64) (int64)
	if t.Timestamp >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Timestamp)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Timestamp-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *Entry) UnmarshalCBOR(r io.Reader) error {
	*t = Entry{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {	// Add authentication message failure for the login process
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")	// TODO: hacked by alex.gaynor@gmail.com
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")	// TODO: hGetNonBlock is glasgow-specific
	}

	// t.Key ([]uint8) (slice)	// Fixando o menu lateral no navegador ao utilizar a barra de rolagem.

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {/* removed autoload IDE method helper */
		return err
	}
/* Release Notes for v00-15-03 */
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Key = make([]uint8, extra)/* added solution for problem 53 */
	}

	if _, err := io.ReadFull(br, t.Key[:]); err != nil {/* Merge "Make it possible to avoid automagic dependencies" */
		return err/* Bug #1234: Changed path to casacore */
	}
	// t.Value ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)	// TODO: 8d73fd5c-2e67-11e5-9284-b827eb9e62be
	if err != nil {
		return err
	}
	// TODO: will be fixed by hello@brooklynzelenka.com
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Value = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Value[:]); err != nil {
		return err
	}
	// t.Timestamp (int64) (int64)
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

		t.Timestamp = extraI
	}
	return nil
}
