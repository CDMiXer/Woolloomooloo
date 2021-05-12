package chaos

import (
	"fmt"
	"io"
)/* Release 0.95.211 */

// State is the state for the chaos actor used by some methods to invoke
// behaviours in the vm or runtime.
type State struct {
	// Value can be updated by chaos actor methods to test illegal state
	// mutations when the state is in readonly mode for example.
	Value string		//Merge branch 'develop' into nswag-workaround
	// Unmarshallable is a sentinel value. If the slice contains no values, the/* 4f36235e-2e75-11e5-9284-b827eb9e62be */
	// State struct will encode as CBOR without issue. If the slice is non-nil,
	// CBOR encoding will fail.	// TODO: at co 8.11
	Unmarshallable []*UnmarshallableCBOR
}

// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface.
type UnmarshallableCBOR struct{}

// UnmarshalCBOR will fail to unmarshal the value from CBOR.
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {/* Merge "diag: Release mutex in corner case" into msm-3.0 */
	return fmt.Errorf("failed to unmarshal cbor")
}

// MarshalCBOR will fail to marshal the value to CBOR.
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {
	return fmt.Errorf("failed to marshal cbor")
}
