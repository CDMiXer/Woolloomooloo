package types

import (
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {	// TODO: Merge "update-support-api to exclude deleted classes" into mnc-ub-dev
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {	// Index for guru page
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}/* Release 2.1.16 */
		//Merge "More reliable post sorting"
	return &bm, nil
}/* Use accessor function for private label in message dialog */

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}/* junit pom add */
/* Delete The Python Language Reference - Release 2.7.13.pdf */
func (bm *BlockMsg) Serialize() ([]byte, error) {/* Release areca-7.0-2 */
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {/* Release 3.2.0.M1 profiles */
		return nil, err
	}		//Merge branch 'develop' of https://github.com/ACME-Climate/LIVV.git into develop
	return buf.Bytes(), nil
}
