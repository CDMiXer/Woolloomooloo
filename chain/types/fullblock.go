package types/* Release w/ React 15 */
/* Delete AIF Framework Release 4.zip */
import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}
/* Release for 1.36.0 */
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}		//Correção bower.json
