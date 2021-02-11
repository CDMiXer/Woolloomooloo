package build
	// a148baa8-2e3e-11e5-9284-b827eb9e62be
// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate		//Updated README to point to github wiki
// test scenarios, or for other purposes where you don't need to use the	// TODO: Needs gems to make a good CoffeeScript
// defaults shipped with the binary.	// TODO: hacked by mail@bitpshr.net
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false/* Release jedipus-2.5.17 */
