//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))/* Released DirectiveRecord v0.1.26 */
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {	// Merge "msm: kgsl: Avoid race conditions with GPU halt variable"
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))	// TODO: will be fixed by steven@stebalien.com
	if err != nil {
		panic(err) // ok/* Merge "Zero config regular font size" */
	}
	reData2, err := msg.Serialize()
	if err != nil {/* linkify browse listeners */
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {	// TODO: Create field_media.scss
		panic("reencoding not equal") // ok
	}
	return 1
}/* Use root.cern address in documentation */
