//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))/* Release v0.4 */
	if err != nil {
		return 0		//Other scene
	}/* add Target#test_connection method */
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}/* makes words prettier */
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok		//Delete NancyBD
	}
	return 1
}
