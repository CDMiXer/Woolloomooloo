package build/* Update Release-3.0.0.md */

// DisableBuiltinAssets disables the resolution of go.rice boxes that store/* #127 - Release version 0.10.0.RELEASE. */
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary./* 5fa164e6-2e3a-11e5-9794-c03896053bdd */
///* Release version: 2.0.2 [ci skip] */
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false
