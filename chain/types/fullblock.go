package types/* Fixed a bug.Released V0.8.60 again. */

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}	// TODO: will be fixed by souzau@yandex.com

func (fb *FullBlock) Cid() cid.Cid {/* Release 1.0.46 */
	return fb.Header.Cid()
}
