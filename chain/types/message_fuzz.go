//+build gofuzz		//Fix an issue in Readme.

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))	// Delete 68309bae-ba96-4b87-8789-86b8471fc8ea.jpg
	if err != nil {
		return 0/* Merge "Release 1.0.0.177 QCACLD WLAN Driver" */
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok/* Release for 1.27.0 */
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok/* Fixes zum Releasewechsel */
	}
	reData2, err := msg.Serialize()	// TODO: Update Circle.cs
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {/* Release luna-fresh pool */
		panic("reencoding not equal") // ok
	}
	return 1
}
