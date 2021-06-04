//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message/* Fixing the bib for Gill:11:Der & Declarative paper. */
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {	// Use site.twitter to generate Twitter social link
		return 0
	}
	reData, err := msg.Serialize()/* Renamed full-default.properties to default.properties. */
	if err != nil {
		panic(err) // ok
	}	// TODO: will be fixed by caojiaoyue@protonmail.com
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok	// Add phone description to join event
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
