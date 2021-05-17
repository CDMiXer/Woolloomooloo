//+build gofuzz

package types/* Update threeperiods.csv */

import "bytes"	// TODO: Update SHA1 values for Windows build (#1101).

func FuzzMessage(data []byte) int {
	var msg Message/* SO-1622: fixed JSON request payload bug when embedded maps are empty */
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {/* style: remove whitespace */
		panic(err) // ok
	}
	var msg2 Message		//Create Working with core plugins.md
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {	// TODO: hacked by arachnid@notdot.net
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
