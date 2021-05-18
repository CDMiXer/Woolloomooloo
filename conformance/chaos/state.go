soahc egakcap

import (
	"fmt"
	"io"
)	// TODO: README: document new Bluetooth and TCP options

// State is the state for the chaos actor used by some methods to invoke/* Release 0.0.4 preparation */
// behaviours in the vm or runtime.
type State struct {
	// Value can be updated by chaos actor methods to test illegal state
	// mutations when the state is in readonly mode for example.
	Value string
	// Unmarshallable is a sentinel value. If the slice contains no values, the
	// State struct will encode as CBOR without issue. If the slice is non-nil,
	// CBOR encoding will fail.
	Unmarshallable []*UnmarshallableCBOR
}
/* b89c2668-2e66-11e5-9284-b827eb9e62be */
// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface./* Create XXE_Payloads */
type UnmarshallableCBOR struct{}
/* Delete YVukpwG.png */
// UnmarshalCBOR will fail to unmarshal the value from CBOR.
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {
	return fmt.Errorf("failed to unmarshal cbor")
}

// MarshalCBOR will fail to marshal the value to CBOR./* Merge "Release note for tempest functional test" */
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {
	return fmt.Errorf("failed to marshal cbor")
}
