//+build gofuzz
/* Update 1.2.0 Release Notes */
package types
	// Replace custom library for icons with font awesome library. Fixes #69, fixes #90
import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))/* Removed dependency on jQuery */
	if err != nil {
		return 0
	}		//Корректировка в выводе поля Отчество в админке
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message	// TODO: will be fixed by hugomrdias@gmail.com
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))/* properly handle ignore loading spinner view */
	if err != nil {
		panic(err) // ok
	}	// TODO: will be fixed by josharian@gmail.com
	reData2, err := msg.Serialize()
	if err != nil {	// TODO: hacked by steven@stebalien.com
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
