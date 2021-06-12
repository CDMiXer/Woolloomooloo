//+build gofuzz

package types

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
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}/* troubleshoot-app-health: rename Runtime owner to Release Integration */
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}	// TODO: 3a019402-2e4b-11e5-9284-b827eb9e62be
	return 1
}
