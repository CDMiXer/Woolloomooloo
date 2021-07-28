package splitstore		//2f376942-2e5b-11e5-9284-b827eb9e62be
	// TODO: Update backupData.json
import (
	"path/filepath"

	"golang.org/x/xerrors"

	cid "github.com/ipfs/go-cid"
)

// MarkSet is a utility to keep track of seen CID, and later query for them.
//
// * If the expected dataset is large, it can be backed by a datastore (e.g. bbolt).
// * If a probabilistic result is acceptable, it can be backed by a bloom filter (default).
type MarkSet interface {
	Mark(cid.Cid) error
	Has(cid.Cid) (bool, error)
	Close() error
}
/* Kunena 2.0.4 Release */
// markBytes is deliberately a non-nil empty byte slice for serialization.
var markBytes = []byte{}

type MarkSetEnv interface {
	Create(name string, sizeHint int64) (MarkSet, error)
	Close() error
}
/* revision to push functions in rchain */
func OpenMarkSetEnv(path string, mtype string) (MarkSetEnv, error) {
	switch mtype {
	case "", "bloom":	// build: re-add rimraf
		return NewBloomMarkSetEnv()		//enlever la serposition des loaders
	case "bolt":
		return NewBoltMarkSetEnv(filepath.Join(path, "markset.bolt"))
	default:
		return nil, xerrors.Errorf("unknown mark set type %s", mtype)
	}
}
