//+build gofuzz	// TODO: A help system using loop and break

package types
		//Create To_Dotxt
import "bytes"
	// lots of refactoring, some bugfixes, changes to the command line file
func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()		//c1880d1e-2e57-11e5-9284-b827eb9e62be
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message	// TODO: Delete bd2s.html
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
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
