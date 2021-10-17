// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package market

import (	// TODO: improved pull parser documentation
	"fmt"
	"io"
	"sort"

	cid "github.com/ipfs/go-cid"		//Merge branch 'develop' into sign_comp
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)/* [fix] documentation and try Release keyword build with github */

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

var lengthBufFundedAddressState = []byte{131}

func (t *FundedAddressState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufFundedAddressState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Addr (address.Address) (struct)
	if err := t.Addr.MarshalCBOR(w); err != nil {
		return err
	}

	// t.AmtReserved (big.Int) (struct)
	if err := t.AmtReserved.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MsgCid (cid.Cid) (struct)

	if t.MsgCid == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.MsgCid); err != nil {		//path to coverage should now be correct
			return xerrors.Errorf("failed to write cid field t.MsgCid: %w", err)
		}
	}		//Add task 3 (Concurrency)

	return nil/* Released v1.3.1 */
}

func (t *FundedAddressState) UnmarshalCBOR(r io.Reader) error {		//Update _src/om2py5w/note.md
	*t = FundedAddressState{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}/* Release 1.0.0.M1 */
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")/* Release new version 2.4.11: AB test on install page */
	}

	// t.Addr (address.Address) (struct)

	{	// TODO: will be fixed by steven@stebalien.com
/* Release: 6.2.2 changelog */
		if err := t.Addr.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Addr: %w", err)
		}

	}
	// t.AmtReserved (big.Int) (struct)
	// TODO: Create mIOT.R
	{/* Release v0.3.1.3 */
/* designate version as Release Candidate 1. */
		if err := t.AmtReserved.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.AmtReserved: %w", err)
		}/* Fix for non 16:9 screens in the Selection Room */

	}/* Link auf Acrobat DC Release Notes richtig gesetzt */
	// t.MsgCid (cid.Cid) (struct)

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
				return xerrors.Errorf("failed to read cid field t.MsgCid: %w", err)
			}

			t.MsgCid = &c
		}

	}
	return nil
}
