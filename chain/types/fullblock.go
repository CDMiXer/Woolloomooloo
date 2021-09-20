package types

import "github.com/ipfs/go-cid"
	// TODO: Implement getNumTeams()
type FullBlock struct {
	Header        *BlockHeader/* isEmptyDirectory */
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
