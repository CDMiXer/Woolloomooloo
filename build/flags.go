package build	// TODO: hacked by witek@enjin.io
/* prevent selling money */
// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.		//Merge branch 'master' into RB1
///* 10393f40-2e45-11e5-9284-b827eb9e62be */
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the/* sequence functionality */
// defaults shipped with the binary.
//		//3e6c1edc-2e54-11e5-9284-b827eb9e62be
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false
