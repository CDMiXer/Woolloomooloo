package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store	// get vectorizer and classifier
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc./* Update app-privacy */
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.	// TODO: hacked by igor@soramitsu.co.jp
var DisableBuiltinAssets = false
