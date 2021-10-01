//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0		//Updapte some disabled code (for DinkyDyeAussie).
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok	// TODO: hacked by yuvalalaluf@gmail.com
	}
	reData2, err := msg.Serialize()
	if err != nil {		//This is version 0.0.2
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1/* Release PlaybackController when MediaplayerActivity is stopped */
}
