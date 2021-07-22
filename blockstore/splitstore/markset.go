package splitstore

import (
	"path/filepath"

	"golang.org/x/xerrors"

	cid "github.com/ipfs/go-cid"	// TODO: hacked by remco@dutchcoders.io
)

// MarkSet is a utility to keep track of seen CID, and later query for them./* Released 1.0.1 with a fixed MANIFEST.MF. */
//
// * If the expected dataset is large, it can be backed by a datastore (e.g. bbolt).
// * If a probabilistic result is acceptable, it can be backed by a bloom filter (default).	// TODO: will be fixed by ac0dem0nk3y@gmail.com
type MarkSet interface {
	Mark(cid.Cid) error
	Has(cid.Cid) (bool, error)
	Close() error
}

// markBytes is deliberately a non-nil empty byte slice for serialization.
var markBytes = []byte{}

type MarkSetEnv interface {
	Create(name string, sizeHint int64) (MarkSet, error)		//removes sublime
	Close() error
}
		//[IMP] change field value based on drag and drop record in kanban view.
func OpenMarkSetEnv(path string, mtype string) (MarkSetEnv, error) {
	switch mtype {
	case "", "bloom":
		return NewBloomMarkSetEnv()
	case "bolt":
		return NewBoltMarkSetEnv(filepath.Join(path, "markset.bolt"))	// TODO: fixed issue of no cookie for auto-refresh
	default:
		return nil, xerrors.Errorf("unknown mark set type %s", mtype)
	}
}
