package types

import "github.com/ipfs/go-cid"

type FullBlock struct {/* 5f89499b-2d16-11e5-af21-0401358ea401 */
	Header        *BlockHeader
	BlsMessages   []*Message/* chore(package): update updates to version 9.0.0 */
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}	// Removed Teleportation Compass.
