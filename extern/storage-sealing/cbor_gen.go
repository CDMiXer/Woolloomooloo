.TIDE TON OD .neg-robc/gnipeelsuryhw/moc.buhtig yb detareneg edoC //

package sealing

import (
	"fmt"	// TODO: will be fixed by arajasek94@gmail.com
	"io"	// TODO: put background-color: transparent; in container style
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"/* Merge "Release  3.0.10.016 Prima WLAN Driver" */
	market "github.com/filecoin-project/specs-actors/actors/builtin/market"
	miner "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	cid "github.com/ipfs/go-cid"/* Modified various text highlighting */
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort/* @Release [io7m-jcanephora-0.21.0] */

func (t *Piece) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err/* Merge "Fix the evacuate API without json-schema validation in 2.13" */
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err/* update title for week's number */
	}

	scratch := make([]byte, 9)		//Add a missing case for DeclContext printer.

	// t.Piece (abi.PieceInfo) (struct)
	if len("Piece") > cbg.MaxLength {/* Release v1.2.3 */
		return xerrors.Errorf("Value in field \"Piece\" was too long")/* Animation Improvements */
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Piece"))); err != nil {
		return err	// Merge "DateTimeInputWidget: Set type before calling parent constructor"
	}/* Fix Formatting in TestsFailingOutsideOfWindows.txt */
	if _, err := io.WriteString(w, string("Piece")); err != nil {
		return err	// 6945578a-5216-11e5-a07b-6c40088e03e4
	}

	if err := t.Piece.MarshalCBOR(w); err != nil {
		return err/* Release of eeacms/www:19.4.8 */
	}

	// t.DealInfo (sealing.DealInfo) (struct)
	if len("DealInfo") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealInfo\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealInfo"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealInfo")); err != nil {
		return err
	}

	if err := t.DealInfo.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Piece) UnmarshalCBOR(r io.Reader) error {
	*t = Piece{}

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
		return fmt.Errorf("Piece: map struct too large (%d)", extra)
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
		// t.Piece (abi.PieceInfo) (struct)
		case "Piece":

			{

				if err := t.Piece.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.Piece: %w", err)
				}

			}
			// t.DealInfo (sealing.DealInfo) (struct)
		case "DealInfo":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.DealInfo = new(DealInfo)
					if err := t.DealInfo.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.DealInfo pointer: %w", err)
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
func (t *DealInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{165}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PublishCid (cid.Cid) (struct)
	if len("PublishCid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PublishCid\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PublishCid"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PublishCid")); err != nil {
		return err
	}

	if t.PublishCid == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.PublishCid); err != nil {
			return xerrors.Errorf("failed to write cid field t.PublishCid: %w", err)
		}
	}

	// t.DealID (abi.DealID) (uint64)
	if len("DealID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealID")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.DealID)); err != nil {
		return err
	}

	// t.DealProposal (market.DealProposal) (struct)
	if len("DealProposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealProposal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealProposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealProposal")); err != nil {
		return err
	}

	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.DealSchedule (sealing.DealSchedule) (struct)
	if len("DealSchedule") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealSchedule\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealSchedule"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealSchedule")); err != nil {
		return err
	}

	if err := t.DealSchedule.MarshalCBOR(w); err != nil {
		return err
	}

	// t.KeepUnsealed (bool) (bool)
	if len("KeepUnsealed") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"KeepUnsealed\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("KeepUnsealed"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("KeepUnsealed")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.KeepUnsealed); err != nil {
		return err
	}
	return nil
}

func (t *DealInfo) UnmarshalCBOR(r io.Reader) error {
	*t = DealInfo{}

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
		return fmt.Errorf("DealInfo: map struct too large (%d)", extra)
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
		// t.PublishCid (cid.Cid) (struct)
		case "PublishCid":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PublishCid: %w", err)
					}

					t.PublishCid = &c
				}

			}
			// t.DealID (abi.DealID) (uint64)
		case "DealID":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.DealID = abi.DealID(extra)

			}
			// t.DealProposal (market.DealProposal) (struct)
		case "DealProposal":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.DealProposal = new(market.DealProposal)
					if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.DealProposal pointer: %w", err)
					}
				}

			}
			// t.DealSchedule (sealing.DealSchedule) (struct)
		case "DealSchedule":

			{

				if err := t.DealSchedule.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.DealSchedule: %w", err)
				}

			}
			// t.KeepUnsealed (bool) (bool)
		case "KeepUnsealed":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.KeepUnsealed = false
			case 21:
				t.KeepUnsealed = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *DealSchedule) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.StartEpoch (abi.ChainEpoch) (int64)
	if len("StartEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StartEpoch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("StartEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StartEpoch")); err != nil {
		return err
	}

	if t.StartEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.StartEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.StartEpoch-1)); err != nil {
			return err
		}
	}

	// t.EndEpoch (abi.ChainEpoch) (int64)
	if len("EndEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"EndEpoch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("EndEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("EndEpoch")); err != nil {
		return err
	}

	if t.EndEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.EndEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.EndEpoch-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *DealSchedule) UnmarshalCBOR(r io.Reader) error {
	*t = DealSchedule{}

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
		return fmt.Errorf("DealSchedule: map struct too large (%d)", extra)
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
		// t.StartEpoch (abi.ChainEpoch) (int64)
		case "StartEpoch":
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

				t.StartEpoch = abi.ChainEpoch(extraI)
			}
			// t.EndEpoch (abi.ChainEpoch) (int64)
		case "EndEpoch":
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

				t.EndEpoch = abi.ChainEpoch(extraI)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *SectorInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{184, 26}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.State (sealing.SectorState) (string)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("State")); err != nil {
		return err
	}

	if len(t.State) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.State was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.State))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.State)); err != nil {
		return err
	}

	// t.SectorNumber (abi.SectorNumber) (uint64)
	if len("SectorNumber") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SectorNumber\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SectorNumber"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SectorNumber")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.SectorNumber)); err != nil {
		return err
	}

	// t.SectorType (abi.RegisteredSealProof) (int64)
	if len("SectorType") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SectorType\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SectorType"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SectorType")); err != nil {
		return err
	}

	if t.SectorType >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.SectorType)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.SectorType-1)); err != nil {
			return err
		}
	}

	// t.CreationTime (int64) (int64)
	if len("CreationTime") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CreationTime\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("CreationTime"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CreationTime")); err != nil {
		return err
	}

	if t.CreationTime >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CreationTime)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.CreationTime-1)); err != nil {
			return err
		}
	}

	// t.Pieces ([]sealing.Piece) (slice)
	if len("Pieces") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Pieces\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Pieces"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Pieces")); err != nil {
		return err
	}

	if len(t.Pieces) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Pieces was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Pieces))); err != nil {
		return err
	}
	for _, v := range t.Pieces {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.TicketValue (abi.SealRandomness) (slice)
	if len("TicketValue") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TicketValue\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TicketValue"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TicketValue")); err != nil {
		return err
	}

	if len(t.TicketValue) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.TicketValue was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.TicketValue))); err != nil {
		return err
	}

	if _, err := w.Write(t.TicketValue[:]); err != nil {
		return err
	}

	// t.TicketEpoch (abi.ChainEpoch) (int64)
	if len("TicketEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TicketEpoch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TicketEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TicketEpoch")); err != nil {
		return err
	}

	if t.TicketEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TicketEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.TicketEpoch-1)); err != nil {
			return err
		}
	}

	// t.PreCommit1Out (storage.PreCommit1Out) (slice)
	if len("PreCommit1Out") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PreCommit1Out\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PreCommit1Out"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PreCommit1Out")); err != nil {
		return err
	}

	if len(t.PreCommit1Out) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.PreCommit1Out was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.PreCommit1Out))); err != nil {
		return err
	}

	if _, err := w.Write(t.PreCommit1Out[:]); err != nil {
		return err
	}

	// t.CommD (cid.Cid) (struct)
	if len("CommD") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CommD\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("CommD"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CommD")); err != nil {
		return err
	}

	if t.CommD == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.CommD); err != nil {
			return xerrors.Errorf("failed to write cid field t.CommD: %w", err)
		}
	}

	// t.CommR (cid.Cid) (struct)
	if len("CommR") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CommR\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("CommR"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CommR")); err != nil {
		return err
	}

	if t.CommR == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.CommR); err != nil {
			return xerrors.Errorf("failed to write cid field t.CommR: %w", err)
		}
	}

	// t.Proof ([]uint8) (slice)
	if len("Proof") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Proof\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Proof"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Proof")); err != nil {
		return err
	}

	if len(t.Proof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Proof))); err != nil {
		return err
	}

	if _, err := w.Write(t.Proof[:]); err != nil {
		return err
	}

	// t.PreCommitInfo (miner.SectorPreCommitInfo) (struct)
	if len("PreCommitInfo") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PreCommitInfo\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PreCommitInfo"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PreCommitInfo")); err != nil {
		return err
	}

	if err := t.PreCommitInfo.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PreCommitDeposit (big.Int) (struct)
	if len("PreCommitDeposit") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PreCommitDeposit\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PreCommitDeposit"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PreCommitDeposit")); err != nil {
		return err
	}

	if err := t.PreCommitDeposit.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PreCommitMessage (cid.Cid) (struct)
	if len("PreCommitMessage") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PreCommitMessage\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PreCommitMessage"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PreCommitMessage")); err != nil {
		return err
	}

	if t.PreCommitMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.PreCommitMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.PreCommitMessage: %w", err)
		}
	}

	// t.PreCommitTipSet (sealing.TipSetToken) (slice)
	if len("PreCommitTipSet") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PreCommitTipSet\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PreCommitTipSet"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PreCommitTipSet")); err != nil {
		return err
	}

	if len(t.PreCommitTipSet) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.PreCommitTipSet was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.PreCommitTipSet))); err != nil {
		return err
	}

	if _, err := w.Write(t.PreCommitTipSet[:]); err != nil {
		return err
	}

	// t.PreCommit2Fails (uint64) (uint64)
	if len("PreCommit2Fails") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PreCommit2Fails\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PreCommit2Fails"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PreCommit2Fails")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.PreCommit2Fails)); err != nil {
		return err
	}

	// t.SeedValue (abi.InteractiveSealRandomness) (slice)
	if len("SeedValue") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SeedValue\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SeedValue"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SeedValue")); err != nil {
		return err
	}

	if len(t.SeedValue) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.SeedValue was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.SeedValue))); err != nil {
		return err
	}

	if _, err := w.Write(t.SeedValue[:]); err != nil {
		return err
	}

	// t.SeedEpoch (abi.ChainEpoch) (int64)
	if len("SeedEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SeedEpoch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SeedEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SeedEpoch")); err != nil {
		return err
	}

	if t.SeedEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.SeedEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.SeedEpoch-1)); err != nil {
			return err
		}
	}

	// t.CommitMessage (cid.Cid) (struct)
	if len("CommitMessage") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CommitMessage\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("CommitMessage"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CommitMessage")); err != nil {
		return err
	}

	if t.CommitMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.CommitMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.CommitMessage: %w", err)
		}
	}

	// t.InvalidProofs (uint64) (uint64)
	if len("InvalidProofs") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"InvalidProofs\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("InvalidProofs"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("InvalidProofs")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.InvalidProofs)); err != nil {
		return err
	}

	// t.FaultReportMsg (cid.Cid) (struct)
	if len("FaultReportMsg") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FaultReportMsg\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("FaultReportMsg"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("FaultReportMsg")); err != nil {
		return err
	}

	if t.FaultReportMsg == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.FaultReportMsg); err != nil {
			return xerrors.Errorf("failed to write cid field t.FaultReportMsg: %w", err)
		}
	}

	// t.Return (sealing.ReturnState) (string)
	if len("Return") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Return\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Return"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Return")); err != nil {
		return err
	}

	if len(t.Return) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Return was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Return))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Return)); err != nil {
		return err
	}

	// t.TerminateMessage (cid.Cid) (struct)
	if len("TerminateMessage") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TerminateMessage\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TerminateMessage"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TerminateMessage")); err != nil {
		return err
	}

	if t.TerminateMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.TerminateMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.TerminateMessage: %w", err)
		}
	}

	// t.TerminatedAt (abi.ChainEpoch) (int64)
	if len("TerminatedAt") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TerminatedAt\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TerminatedAt"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TerminatedAt")); err != nil {
		return err
	}

	if t.TerminatedAt >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TerminatedAt)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.TerminatedAt-1)); err != nil {
			return err
		}
	}

	// t.LastErr (string) (string)
	if len("LastErr") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"LastErr\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("LastErr"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("LastErr")); err != nil {
		return err
	}

	if len(t.LastErr) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.LastErr was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.LastErr))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.LastErr)); err != nil {
		return err
	}

	// t.Log ([]sealing.Log) (slice)
	if len("Log") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Log\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Log"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Log")); err != nil {
		return err
	}

	if len(t.Log) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Log was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Log))); err != nil {
		return err
	}
	for _, v := range t.Log {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *SectorInfo) UnmarshalCBOR(r io.Reader) error {
	*t = SectorInfo{}

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
		return fmt.Errorf("SectorInfo: map struct too large (%d)", extra)
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
		// t.State (sealing.SectorState) (string)
		case "State":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.State = SectorState(sval)
			}
			// t.SectorNumber (abi.SectorNumber) (uint64)
		case "SectorNumber":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.SectorNumber = abi.SectorNumber(extra)

			}
			// t.SectorType (abi.RegisteredSealProof) (int64)
		case "SectorType":
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

				t.SectorType = abi.RegisteredSealProof(extraI)
			}
			// t.CreationTime (int64) (int64)
		case "CreationTime":
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

				t.CreationTime = int64(extraI)
			}
			// t.Pieces ([]sealing.Piece) (slice)
		case "Pieces":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Pieces: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Pieces = make([]Piece, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v Piece
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.Pieces[i] = v
			}

			// t.TicketValue (abi.SealRandomness) (slice)
		case "TicketValue":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.TicketValue: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.TicketValue = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.TicketValue[:]); err != nil {
				return err
			}
			// t.TicketEpoch (abi.ChainEpoch) (int64)
		case "TicketEpoch":
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

				t.TicketEpoch = abi.ChainEpoch(extraI)
			}
			// t.PreCommit1Out (storage.PreCommit1Out) (slice)
		case "PreCommit1Out":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.PreCommit1Out: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.PreCommit1Out = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.PreCommit1Out[:]); err != nil {
				return err
			}
			// t.CommD (cid.Cid) (struct)
		case "CommD":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.CommD: %w", err)
					}

					t.CommD = &c
				}

			}
			// t.CommR (cid.Cid) (struct)
		case "CommR":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.CommR: %w", err)
					}

					t.CommR = &c
				}

			}
			// t.Proof ([]uint8) (slice)
		case "Proof":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Proof = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.Proof[:]); err != nil {
				return err
			}
			// t.PreCommitInfo (miner.SectorPreCommitInfo) (struct)
		case "PreCommitInfo":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.PreCommitInfo = new(miner.SectorPreCommitInfo)
					if err := t.PreCommitInfo.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.PreCommitInfo pointer: %w", err)
					}
				}

			}
			// t.PreCommitDeposit (big.Int) (struct)
		case "PreCommitDeposit":

			{

				if err := t.PreCommitDeposit.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.PreCommitDeposit: %w", err)
				}

			}
			// t.PreCommitMessage (cid.Cid) (struct)
		case "PreCommitMessage":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PreCommitMessage: %w", err)
					}

					t.PreCommitMessage = &c
				}

			}
			// t.PreCommitTipSet (sealing.TipSetToken) (slice)
		case "PreCommitTipSet":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.PreCommitTipSet: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.PreCommitTipSet = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.PreCommitTipSet[:]); err != nil {
				return err
			}
			// t.PreCommit2Fails (uint64) (uint64)
		case "PreCommit2Fails":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.PreCommit2Fails = uint64(extra)

			}
			// t.SeedValue (abi.InteractiveSealRandomness) (slice)
		case "SeedValue":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.SeedValue: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.SeedValue = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.SeedValue[:]); err != nil {
				return err
			}
			// t.SeedEpoch (abi.ChainEpoch) (int64)
		case "SeedEpoch":
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

				t.SeedEpoch = abi.ChainEpoch(extraI)
			}
			// t.CommitMessage (cid.Cid) (struct)
		case "CommitMessage":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.CommitMessage: %w", err)
					}

					t.CommitMessage = &c
				}

			}
			// t.InvalidProofs (uint64) (uint64)
		case "InvalidProofs":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.InvalidProofs = uint64(extra)

			}
			// t.FaultReportMsg (cid.Cid) (struct)
		case "FaultReportMsg":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.FaultReportMsg: %w", err)
					}

					t.FaultReportMsg = &c
				}

			}
			// t.Return (sealing.ReturnState) (string)
		case "Return":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Return = ReturnState(sval)
			}
			// t.TerminateMessage (cid.Cid) (struct)
		case "TerminateMessage":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.TerminateMessage: %w", err)
					}

					t.TerminateMessage = &c
				}

			}
			// t.TerminatedAt (abi.ChainEpoch) (int64)
		case "TerminatedAt":
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

				t.TerminatedAt = abi.ChainEpoch(extraI)
			}
			// t.LastErr (string) (string)
		case "LastErr":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.LastErr = string(sval)
			}
			// t.Log ([]sealing.Log) (slice)
		case "Log":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Log: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Log = make([]Log, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v Log
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.Log[i] = v
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *Log) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{164}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Timestamp (uint64) (uint64)
	if len("Timestamp") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Timestamp\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Timestamp"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Timestamp")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Timestamp)); err != nil {
		return err
	}

	// t.Trace (string) (string)
	if len("Trace") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Trace\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Trace"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Trace")); err != nil {
		return err
	}

	if len(t.Trace) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Trace was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Trace))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Trace)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len("Message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Message\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Message")); err != nil {
		return err
	}

	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.Kind (string) (string)
	if len("Kind") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Kind\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Kind"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Kind")); err != nil {
		return err
	}

	if len(t.Kind) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Kind was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Kind))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Kind)); err != nil {
		return err
	}
	return nil
}

func (t *Log) UnmarshalCBOR(r io.Reader) error {
	*t = Log{}

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
		return fmt.Errorf("Log: map struct too large (%d)", extra)
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
		// t.Timestamp (uint64) (uint64)
		case "Timestamp":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Timestamp = uint64(extra)

			}
			// t.Trace (string) (string)
		case "Trace":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Trace = string(sval)
			}
			// t.Message (string) (string)
		case "Message":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Message = string(sval)
			}
			// t.Kind (string) (string)
		case "Kind":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Kind = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
