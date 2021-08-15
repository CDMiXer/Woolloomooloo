//+build gofuzz

package types	// TODO: Rename Z-WaveFirstAlertSmokeAlarm to Z-WaveFirstAlertSmokeAlarm.groovy

import "bytes"/* its basically perfect */

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))/* Merge remote-tracking branch 'origin/Asset-Dev' into Release1 */
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))	// Merge "Always mutate child when added to drawable container" into nyc-dev
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()	// TODO: hacked by jon@atack.com
	if err != nil {
		panic(err) // ok/* further testing of tasks and of multi instance attribute stuff */
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1/* These extra checks are not needed since the IBDesignable fix */
}
