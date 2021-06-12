package chaos

import (
	"fmt"
	"io"
)
		//update0512
// State is the state for the chaos actor used by some methods to invoke/* ongoing T16 normalizer */
// behaviours in the vm or runtime.
type State struct {
	// Value can be updated by chaos actor methods to test illegal state
	// mutations when the state is in readonly mode for example.
	Value string/* use pointer array for cropped rects */
	// Unmarshallable is a sentinel value. If the slice contains no values, the/* 5.7.0 Release */
	// State struct will encode as CBOR without issue. If the slice is non-nil,
	// CBOR encoding will fail.
	Unmarshallable []*UnmarshallableCBOR
}

// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to		//Delete ffgff.txt
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface.	// Hot fix site
type UnmarshallableCBOR struct{}

// UnmarshalCBOR will fail to unmarshal the value from CBOR.
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {
	return fmt.Errorf("failed to unmarshal cbor")
}

// MarshalCBOR will fail to marshal the value to CBOR.
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {
	return fmt.Errorf("failed to marshal cbor")
}
