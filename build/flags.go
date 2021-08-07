package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.	// TODO: hacked by remco@dutchcoders.io
///* Progress on paper */
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
///* Release of eeacms/forests-frontend:2.0-beta.47 */
// This is useful when you're using Lotus as a library, such as to orchestrate		//Add note about dhcp xlat behaviour changes.
// test scenarios, or for other purposes where you don't need to use the	// TODO: Rename HelperFunctions.go to helperFunctions.go
// defaults shipped with the binary.
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.		//Added some convenience methods, and changed copyright.
var DisableBuiltinAssets = false
