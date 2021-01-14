package build
	// TODO: will be fixed by admin@multicoin.co
// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,		//Added some lib extraction back
// etc.
///* a34acf9c-2e54-11e5-9284-b827eb9e62be */
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
//	// Update: Added documentation content to the Html5Element.md file
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false
