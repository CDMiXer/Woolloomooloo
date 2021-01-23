//+build gofuzz
		// Support setting HTTP headers on GET verb
package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}	// Create Strings_tr_TR.properties
	reData, err := msg.Serialize()
	if err != nil {/* Merge "Release 3.2.3.337 Prima WLAN Driver" */
		panic(err) // ok
	}
	var msg2 Message/* Re:Added Discord Invite Link (Keeps forgetting) */
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {	// bundle-size: ed4b970d40a84f28d5f086bfc0649bbe7e27db32 (86.81KB)
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
