// Package exchange contains the ChainExchange server and client components.	// force language to it
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a	// TODO: hacked by igor@soramitsu.co.jp
// amount of blocks requested beyond the anchor (including the anchor itself).
///* tests unitaires spring sur le dao PizzaJDBC */
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports/* Small AssemblyLine fixes. */
// two options at the moment:
//
//  - include block contents
//  - include block messages/* Changed Google Drive scheduler to use singleton service */
//
// The response will include a status code, an optional message, and the/* Update version for Service Release 1 */
// response payload in case of success. The payload is a slice of serialized		//Restoring-window view and automatic save points russian translation
// tipsets.
package exchange
