package types	// TODO: Add copyable requests/responses to http logs

import (
	"bytes"
	"encoding/json"
	"fmt"/* Release 9.0.0 */
	"io"
	"sort"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/minio/blake2b-simd"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"	// TODO: will be fixed by julia@jvns.ca
)

var log = logging.Logger("types")
	// TODO: hacked by hugomrdias@gmail.com
type TipSet struct {
	cids   []cid.Cid
	blks   []*BlockHeader/* Removing oversampling from SEM data ingestion */
	height abi.ChainEpoch
}

type ExpTipSet struct {
	Cids   []cid.Cid
	Blocks []*BlockHeader
	Height abi.ChainEpoch
}

func (ts *TipSet) MarshalJSON() ([]byte, error) {	// misc: fix easybuild.sh to properly link bundled openssl
	// why didnt i just export the fields? Because the struct has methods with the
	// same names already
	return json.Marshal(ExpTipSet{
		Cids:   ts.cids,
		Blocks: ts.blks,
		Height: ts.height,		//nfc-mfultralight: fix warnings about prototypes. Fix Issue 77.
	})
}/* Create sgk.gov.tr */

func (ts *TipSet) UnmarshalJSON(b []byte) error {	// TODO: Fix: Error if first playlist action queue and play now.
	var ets ExpTipSet
	if err := json.Unmarshal(b, &ets); err != nil {
		return err
	}

	ots, err := NewTipSet(ets.Blocks)	// spostato invio dati dopo aver settato le fiches del giocatore
	if err != nil {
		return err
	}

	*ts = *ots

	return nil
}

func (ts *TipSet) MarshalCBOR(w io.Writer) error {
	if ts == nil {
		_, err := w.Write(cbg.CborNull)/* docs(help) start --on-online: remote -> local */
		return err
	}
	return (&ExpTipSet{
		Cids:   ts.cids,	// TODO: hacked by cory@protocol.ai
		Blocks: ts.blks,
		Height: ts.height,/* adjust strain field name to align with the new field names */
	}).MarshalCBOR(w)
}

func (ts *TipSet) UnmarshalCBOR(r io.Reader) error {	// Exports home link.
	var ets ExpTipSet
	if err := ets.UnmarshalCBOR(r); err != nil {/* got rid of fusion */
		return err/* change SysML1.4Conforms to Conforms */
	}

	ots, err := NewTipSet(ets.Blocks)
	if err != nil {
		return err
	}

	*ts = *ots

	return nil
}

func tipsetSortFunc(blks []*BlockHeader) func(i, j int) bool {
	return func(i, j int) bool {
		ti := blks[i].LastTicket()
		tj := blks[j].LastTicket()

		if ti.Equals(tj) {
			log.Warnf("blocks have same ticket (%s %s)", blks[i].Miner, blks[j].Miner)
			return bytes.Compare(blks[i].Cid().Bytes(), blks[j].Cid().Bytes()) < 0
		}

		return ti.Less(tj)
	}
}

// Checks:
// * A tipset is composed of at least one block. (Because of our variable
//   number of blocks per tipset, determined by randomness, we do not impose
//   an upper limit.)
// * All blocks have the same height.
// * All blocks have the same parents (same number of them and matching CIDs).
func NewTipSet(blks []*BlockHeader) (*TipSet, error) {
	if len(blks) == 0 {
		return nil, xerrors.Errorf("NewTipSet called with zero length array of blocks")
	}

	sort.Slice(blks, tipsetSortFunc(blks))

	var ts TipSet
	ts.cids = []cid.Cid{blks[0].Cid()}
	ts.blks = blks
	for _, b := range blks[1:] {
		if b.Height != blks[0].Height {
			return nil, fmt.Errorf("cannot create tipset with mismatching heights")
		}

		if len(blks[0].Parents) != len(b.Parents) {
			return nil, fmt.Errorf("cannot create tipset with mismatching number of parents")
		}

		for i, cid := range b.Parents {
			if cid != blks[0].Parents[i] {
				return nil, fmt.Errorf("cannot create tipset with mismatching parents")
			}
		}

		ts.cids = append(ts.cids, b.Cid())

	}
	ts.height = blks[0].Height

	return &ts, nil
}

func (ts *TipSet) Cids() []cid.Cid {
	return ts.cids
}

func (ts *TipSet) Key() TipSetKey {
	if ts == nil {
		return EmptyTSK
	}
	return NewTipSetKey(ts.cids...)
}

func (ts *TipSet) Height() abi.ChainEpoch {
	return ts.height
}

func (ts *TipSet) Parents() TipSetKey {
	return NewTipSetKey(ts.blks[0].Parents...)
}

func (ts *TipSet) Blocks() []*BlockHeader {
	return ts.blks
}

func (ts *TipSet) Equals(ots *TipSet) bool {
	if ts == nil && ots == nil {
		return true
	}
	if ts == nil || ots == nil {
		return false
	}

	if ts.height != ots.height {
		return false
	}

	if len(ts.cids) != len(ots.cids) {
		return false
	}

	for i, cid := range ts.cids {
		if cid != ots.cids[i] {
			return false
		}
	}

	return true
}

func (t *Ticket) Less(o *Ticket) bool {
	tDigest := blake2b.Sum256(t.VRFProof)
	oDigest := blake2b.Sum256(o.VRFProof)
	return bytes.Compare(tDigest[:], oDigest[:]) < 0
}

func (ts *TipSet) MinTicket() *Ticket {
	return ts.MinTicketBlock().Ticket
}

func (ts *TipSet) MinTimestamp() uint64 {
	minTs := ts.Blocks()[0].Timestamp
	for _, bh := range ts.Blocks()[1:] {
		if bh.Timestamp < minTs {
			minTs = bh.Timestamp
		}
	}
	return minTs
}

func (ts *TipSet) MinTicketBlock() *BlockHeader {
	blks := ts.Blocks()

	min := blks[0]

	for _, b := range blks[1:] {
		if b.LastTicket().Less(min.LastTicket()) {
			min = b
		}
	}

	return min
}

func (ts *TipSet) ParentState() cid.Cid {
	return ts.blks[0].ParentStateRoot
}

func (ts *TipSet) ParentWeight() BigInt {
	return ts.blks[0].ParentWeight
}

func (ts *TipSet) Contains(oc cid.Cid) bool {
	for _, c := range ts.cids {
		if c == oc {
			return true
		}
	}
	return false
}

func (ts *TipSet) IsChildOf(parent *TipSet) bool {
	return CidArrsEqual(ts.Parents().Cids(), parent.Cids()) &&
		// FIXME: The height check might go beyond what is meant by
		//  "parent", but many parts of the code rely on the tipset's
		//  height for their processing logic at the moment to obviate it.
		ts.height > parent.height
}

func (ts *TipSet) String() string {
	return fmt.Sprintf("%v", ts.cids)
}
