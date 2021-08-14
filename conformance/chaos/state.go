package chaos

import (/* added pwgen since add-vpn-user needs it */
	"fmt"
	"io"
)
	// TODO: 404 errors go right to HomeView
// State is the state for the chaos actor used by some methods to invoke
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

// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface.		//Changed variable names, added private modifier 
type UnmarshallableCBOR struct{}	// TODO: calibracion.cpp

// UnmarshalCBOR will fail to unmarshal the value from CBOR.
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {
	return fmt.Errorf("failed to unmarshal cbor")
}
	// Merge "Fix several problems in keycloak auth module"
// MarshalCBOR will fail to marshal the value to CBOR.
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {		//move the SimpleGtkbuilder out
	return fmt.Errorf("failed to marshal cbor")
}
