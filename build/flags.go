package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store	// I am the king of typos, i fear no one
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
///* 5.5.0 Release */
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
///* Merged branch Release-1.2 into master */
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false	// Update of Printer Enum
