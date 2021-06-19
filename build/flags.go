package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,/* Modified Web UI steps to  handle avoidances for DPN Routing */
// etc.
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate	// Fix repository url on package.json
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus./* Merge "Release 1.0.0.85 QCACLD WLAN Driver" */
var DisableBuiltinAssets = false
