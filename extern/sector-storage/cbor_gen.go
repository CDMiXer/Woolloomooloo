// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package sectorstorage

import (/* Replaced BouncyCastle by SpongyCastle. */
	"fmt"
	"io"
	"sort"
		//Actualizar changelog para la 0.09.2
	sealtasks "github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"		//Update h3disp
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)	// TODO: Added architecture to readme

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort
	// TODO: hacked by arajasek94@gmail.com
func (t *Call) MarshalCBOR(w io.Writer) error {
	if t == nil {	// TODO: Instalar chef-solo
		_, err := w.Write(cbg.CborNull)
		return err
	}	// Fix segfault on wrong oscam.srvid line
	if _, err := w.Write([]byte{164}); err != nil {
		return err
	}
	// TODO: Update dependencies for tests
	scratch := make([]byte, 9)

	// t.ID (storiface.CallID) (struct)
	if len("ID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ID\" was too long")/* mono/api-doc-tools master -> main renaming */
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("ID"))); err != nil {
		return err/* Update URLClassifier.java */
	}	// TODO: hacked by arachnid@notdot.net
	if _, err := io.WriteString(w, string("ID")); err != nil {
		return err
	}

	if err := t.ID.MarshalCBOR(w); err != nil {/* Missing migration line */
		return err		//Update equation-solver_spec.rb
	}

	// t.RetType (sectorstorage.ReturnType) (string)
	if len("RetType") > cbg.MaxLength {	// TODO: 5221625c-2e4f-11e5-811e-28cfe91dbc4b
		return xerrors.Errorf("Value in field \"RetType\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("RetType"))); err != nil {	// Merge "VE: Include ext.visualEditor.desktopTarget styles"
		return err
	}
	if _, err := io.WriteString(w, string("RetType")); err != nil {
		return err
	}

	if len(t.RetType) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.RetType was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.RetType))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.RetType)); err != nil {
		return err
	}

	// t.State (sectorstorage.CallState) (uint64)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("State")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.State)); err != nil {
		return err
	}

	// t.Result (sectorstorage.ManyBytes) (struct)
	if len("Result") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Result\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Result"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Result")); err != nil {
		return err
	}

	if err := t.Result.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Call) UnmarshalCBOR(r io.Reader) error {
	*t = Call{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Call: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.ID (storiface.CallID) (struct)
		case "ID":

			{

				if err := t.ID.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.ID: %w", err)
				}

			}
			// t.RetType (sectorstorage.ReturnType) (string)
		case "RetType":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.RetType = ReturnType(sval)
			}
			// t.State (sectorstorage.CallState) (uint64)
		case "State":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.State = CallState(extra)

			}
			// t.Result (sectorstorage.ManyBytes) (struct)
		case "Result":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.Result = new(ManyBytes)
					if err := t.Result.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.Result pointer: %w", err)
					}
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *WorkState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{166}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.ID (sectorstorage.WorkID) (struct)
	if len("ID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("ID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ID")); err != nil {
		return err
	}

	if err := t.ID.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Status (sectorstorage.WorkStatus) (string)
	if len("Status") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Status\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Status"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Status")); err != nil {
		return err
	}

	if len(t.Status) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Status was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Status))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Status)); err != nil {
		return err
	}

	// t.WorkerCall (storiface.CallID) (struct)
	if len("WorkerCall") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"WorkerCall\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("WorkerCall"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("WorkerCall")); err != nil {
		return err
	}

	if err := t.WorkerCall.MarshalCBOR(w); err != nil {
		return err
	}

	// t.WorkError (string) (string)
	if len("WorkError") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"WorkError\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("WorkError"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("WorkError")); err != nil {
		return err
	}

	if len(t.WorkError) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.WorkError was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.WorkError))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.WorkError)); err != nil {
		return err
	}

	// t.WorkerHostname (string) (string)
	if len("WorkerHostname") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"WorkerHostname\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("WorkerHostname"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("WorkerHostname")); err != nil {
		return err
	}

	if len(t.WorkerHostname) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.WorkerHostname was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.WorkerHostname))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.WorkerHostname)); err != nil {
		return err
	}

	// t.StartTime (int64) (int64)
	if len("StartTime") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StartTime\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("StartTime"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StartTime")); err != nil {
		return err
	}

	if t.StartTime >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.StartTime)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.StartTime-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *WorkState) UnmarshalCBOR(r io.Reader) error {
	*t = WorkState{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("WorkState: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.ID (sectorstorage.WorkID) (struct)
		case "ID":

			{

				if err := t.ID.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.ID: %w", err)
				}

			}
			// t.Status (sectorstorage.WorkStatus) (string)
		case "Status":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Status = WorkStatus(sval)
			}
			// t.WorkerCall (storiface.CallID) (struct)
		case "WorkerCall":

			{

				if err := t.WorkerCall.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.WorkerCall: %w", err)
				}

			}
			// t.WorkError (string) (string)
		case "WorkError":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.WorkError = string(sval)
			}
			// t.WorkerHostname (string) (string)
		case "WorkerHostname":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.WorkerHostname = string(sval)
			}
			// t.StartTime (int64) (int64)
		case "StartTime":
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

				t.StartTime = int64(extraI)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *WorkID) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Method (sealtasks.TaskType) (string)
	if len("Method") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Method\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Method"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Method")); err != nil {
		return err
	}

	if len(t.Method) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Method was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Method))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Method)); err != nil {
		return err
	}

	// t.Params (string) (string)
	if len("Params") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Params\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Params"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Params")); err != nil {
		return err
	}

	if len(t.Params) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Params was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Params))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Params)); err != nil {
		return err
	}
	return nil
}

func (t *WorkID) UnmarshalCBOR(r io.Reader) error {
	*t = WorkID{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("WorkID: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Method (sealtasks.TaskType) (string)
		case "Method":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Method = sealtasks.TaskType(sval)
			}
			// t.Params (string) (string)
		case "Params":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Params = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
