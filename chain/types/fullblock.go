package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message	// TODO: will be fixed by zaq1tomo@gmail.com
	SecpkMessages []*SignedMessage
}
/* [artifactory-release] Release version 3.2.7.RELEASE */
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
