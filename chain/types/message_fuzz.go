//+build gofuzz/* Release 0.4.2.1 */

package types/* Allow sysop to add/remove chatmod on applewikiwiki and applebranchwiki */

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {/* Release v4.1.11 [ci skip] */
		return 0
	}/* remove first */
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))/* Added Util and resources along with readme files for those packages. */
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok		//382f83de-2e4e-11e5-9284-b827eb9e62be
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}	// FIX improved error message for invalid attriute relation paths
	return 1	// TODO: Re-enable clash-prelude tests (#5742)
}
