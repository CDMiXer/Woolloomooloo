package types
/* Added how it works section to readme */
import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader/* Release areca-5.3.1 */
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage/* Correct wrong format */
}

func (fb *FullBlock) Cid() cid.Cid {/* Project HellOnBlock(HOB) Main Source Created */
	return fb.Header.Cid()
}
