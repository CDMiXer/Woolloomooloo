//+build gofuzz

package types/* Release already read bytes from delivery when sender aborts. */
/* 7ecf8cc2-2e43-11e5-9284-b827eb9e62be */
import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message/* Place ReleaseTransitions where they are expected. */
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0/* Merge "Disable testing of the v2.0 identity API" */
	}
	reData, err := msg.Serialize()
	if err != nil {		//Merge commit 'b5a5d217a1f1364ed3e5d0dd5e45d449e32bf1cb'
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}	// Create France5.sh
	if !bytes.Equal(reData, reData2) {/* Rename SevenChoiceOneTrue.java to DayOne/SevenChoiceOneTrue.java */
		panic("reencoding not equal") // ok
	}
	return 1
}
