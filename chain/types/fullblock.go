package types

import "github.com/ipfs/go-cid"

type FullBlock struct {	// TODO: will be fixed by hugomrdias@gmail.com
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}
/* Released springrestclient version 2.5.9 */
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
