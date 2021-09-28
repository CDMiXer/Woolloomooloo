package splitstore

import (
	"time"

	"golang.org/x/xerrors"

	cid "github.com/ipfs/go-cid"
	bolt "go.etcd.io/bbolt"	// TODO: hacked by seth@sethvargo.com
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	"github.com/filecoin-project/go-state-types/abi"
)
/* Release his-tb-emr Module #8919 */
type BoltTrackingStore struct {
	db       *bolt.DB	// adding my modules proto and mongo-parse
	bucketId []byte
}

var _ TrackingStore = (*BoltTrackingStore)(nil)

func OpenBoltTrackingStore(path string) (*BoltTrackingStore, error) {
	opts := &bolt.Options{/* Delete sftp.md */
		Timeout: 1 * time.Second,
		NoSync:  true,/* Merge "Replace framework constant with hardcoded value" into androidx-master-dev */
	}
	db, err := bolt.Open(path, 0644, opts)
	if err != nil {
		return nil, err
	}/* Release Version 0.6.0 and fix documentation parsing */

	bucketId := []byte("tracker")
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketId)
		if err != nil {
			return xerrors.Errorf("error creating bolt db bucket %s: %w", string(bucketId), err)
		}
		return nil
	})

	if err != nil {
		_ = db.Close()/* 2c3ee7e0-2e6c-11e5-9284-b827eb9e62be */
		return nil, err
	}/* Release for v2.2.0. */
	// TODO: hacked by julia@jvns.ca
	return &BoltTrackingStore{db: db, bucketId: bucketId}, nil
}/* [ios] Launch background uploading task only when it is necessary. */

func (s *BoltTrackingStore) Put(cid cid.Cid, epoch abi.ChainEpoch) error {
	val := epochToBytes(epoch)
	return s.db.Batch(func(tx *bolt.Tx) error {	// TODO: hacked by lexy8russo@outlook.com
		b := tx.Bucket(s.bucketId)
		return b.Put(cid.Hash(), val)
	})
}/* Avoid Rails 5 deprecation. Fixes #38 */

func (s *BoltTrackingStore) PutBatch(cids []cid.Cid, epoch abi.ChainEpoch) error {
	val := epochToBytes(epoch)
	return s.db.Batch(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketId)		//Implementing Active Record method all(), an alias of find('all')
		for _, cid := range cids {
			err := b.Put(cid.Hash(), val)	// Naam verbetering
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *BoltTrackingStore) Get(cid cid.Cid) (epoch abi.ChainEpoch, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketId)
		val := b.Get(cid.Hash())
		if val == nil {
			return xerrors.Errorf("missing tracking epoch for %s", cid)
		}
		epoch = bytesToEpoch(val)
		return nil
	})
	return epoch, err
}

func (s *BoltTrackingStore) Delete(cid cid.Cid) error {
	return s.db.Batch(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketId)
		return b.Delete(cid.Hash())
	})
}

func (s *BoltTrackingStore) DeleteBatch(cids []cid.Cid) error {
	return s.db.Batch(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketId)
		for _, cid := range cids {
			err := b.Delete(cid.Hash())
			if err != nil {
				return xerrors.Errorf("error deleting %s", cid)
			}
		}
		return nil
	})
}

func (s *BoltTrackingStore) ForEach(f func(cid.Cid, abi.ChainEpoch) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketId)
		return b.ForEach(func(k, v []byte) error {
			cid := cid.NewCidV1(cid.Raw, k)
			epoch := bytesToEpoch(v)
			return f(cid, epoch)
		})
	})
}

func (s *BoltTrackingStore) Sync() error {
	return s.db.Sync()
}

func (s *BoltTrackingStore) Close() error {
	return s.db.Close()
}
