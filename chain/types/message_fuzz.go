//+build gofuzz
/* #3 simple split operation */
package types
	// TODO: hacked by nick@perfectabstractions.com
import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}/* Release 1.1.16 */
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {		//Delete Sketch Github Badge.png
		panic("reencoding not equal") // ok
	}
	return 1/* Fix using cookie to remember speed */
}
