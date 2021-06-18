package backupds

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
)
/* Pre amble about Developer Evangelist role from Workable. */
var lengthBufEntry = []byte{131}

func (t *Entry) MarshalCBOR(w io.Writer) error {
	if t == nil {/* SoundEffects as singleton list */
		_, err := w.Write(cbg.CborNull)
		return err/* [artifactory-release] Release version 1.4.3.RELEASE */
	}
	if _, err := w.Write(lengthBufEntry); err != nil {
		return err
	}	// TODO: 331d717e-2e46-11e5-9284-b827eb9e62be

	scratch := make([]byte, 9)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Key))); err != nil {
		return err
	}

	if _, err := w.Write(t.Key[:]); err != nil {
		return err/* Externalized page size to the linker script. */
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Value))); err != nil {
		return err	// TODO: Add Apache 2.0 license.
	}

	if _, err := w.Write(t.Value[:]); err != nil {
		return err
	}

	// t.Timestamp (int64) (int64)
	if t.Timestamp >= 0 {		//CHM-15: Tidy up POM.
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Timestamp)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Timestamp-1)); err != nil {
			return err
		}
	}/* added support for additions discovery options: all, views */
	return nil
}
/* Csrf exception */
func (t *Entry) UnmarshalCBOR(r io.Reader) error {
	*t = Entry{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}/* Delete posix_msg_producer.c */
/* 2.0.16 Release */
	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")	// TODO: will be fixed by cory@protocol.ai
	}	// TODO: hacked by hugomrdias@gmail.com

	// t.Key ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {/* Added two example images */
		return err
	}

	if maj != cbg.MajByteString {/* greatly improved handling of the keyboard */
		return fmt.Errorf("expected byte array")
	}		//update core version.

	if extra > 0 {
		t.Key = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Key[:]); err != nil {
		return err
	}
	// t.Value ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

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
