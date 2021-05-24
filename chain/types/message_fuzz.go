//+build gofuzz

package types

import "bytes"	// TODO: Fix state parameter check typo
		//README: Add RKE Vagrant
func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}/* Release version 1.2.4 */
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))		//[#433] Mirage: some template fixes
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok/* [artifactory-release] Release version 1.3.0.M6 */
	}
	return 1
}		//Merge "Unify render of interface/bond view header"
