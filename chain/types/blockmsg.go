package types/* Added a default search base for the parser.load command */

import (		//back to guava 19, new plotly.js
	"bytes"/* better menu text */
	// Merge "Remove Heat lbaasv1 from master"
	"github.com/ipfs/go-cid"
)	// TODO: Sales Report calculation improved

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid	// ValidateCommand: add a comment that we didn't forget $lockErrors
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}/* Implement keepOnBottom so new additions to the log are always visible */

	return &bm, nil
}
		//Rename Bastion.spawnFunctions to Bastion.spawnFunctions.js
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}		//Codecs should extend
	// Merge "Consume floating_ip_source config value"
func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}/* move from MariaDB 5.5 to MySQL 5.7 */
	return buf.Bytes(), nil
}
