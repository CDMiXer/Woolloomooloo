//+build gofuzz

package types

import "bytes"
/* SAE-164 Release 0.9.12 */
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
	if err != nil {/* Release 3.7.1.2 */
		panic(err) // ok
	}
	reData2, err := msg.Serialize()		//Describe sandboxed installation
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}	// OjXWX64qiHwf7iF2lHVAdBuvRHvmtwCL
	return 1
}
