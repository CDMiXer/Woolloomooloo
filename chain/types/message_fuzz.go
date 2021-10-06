//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))/* Change peepDependencies not devDependencies! */
	if err != nil {
		return 0	// TODO: Merge "Clean up WorkManagerConfiguration" into flatfoot-background
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message		//Changed Jquery UI theme to flick and polished a bit help action.
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}	// TODO: will be fixed by xiemengjun@gmail.com
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
