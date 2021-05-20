package build/* Merge "Release 3.2.3.371 Prima WLAN Driver" */

// DisableBuiltinAssets disables the resolution of go.rice boxes that store/* Release v.0.0.1 */
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,	// TODO: Merge "ARM: dts: msm: disable charging only on 8994 CDPs"
// etc.
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//		//DBFlute Engine: Oracle emergency option for too many procedures
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
///* Release 5.2.2 prep */
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false
